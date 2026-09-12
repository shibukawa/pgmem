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
	var v82 int32
	_ = v82
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
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
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 float64
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v478 int32
	_ = v478
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v515 int32
	_ = v515
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v582 int32
	_ = v582
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
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
	var v768 int32
	_ = v768
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
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
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
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
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
	var v950 int32
	_ = v950
	var v972 float64
	_ = v972
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
	var v1119 float64
	_ = v1119
	var v1121 float64
	_ = v1121
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
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 float64
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int64
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1203 int32
	_ = v1203
	var v1230 float64
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1236 float64
	_ = v1236
	var v1237 float64
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 float64
	_ = v1244
	var v1245 float64
	_ = v1245
	var v1246 float64
	_ = v1246
	var v1249 float64
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1255 float64
	_ = v1255
	var v1256 float64
	_ = v1256
	var v1259 float64
	_ = v1259
	var v1260 float64
	_ = v1260
	var v1263 float64
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1318 float64
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 float64
	_ = v1329
	var v1330 float64
	_ = v1330
	var v1331 float64
	_ = v1331
	var v1333 float64
	_ = v1333
	var v1341 float64
	_ = v1341
	var v1343 float64
	_ = v1343
	var v1345 float64
	_ = v1345
	var v1346 float64
	_ = v1346
	var v1354 float64
	_ = v1354
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
	var v1361 float64
	_ = v1361
	var v1369 float64
	_ = v1369
	var v1371 float64
	_ = v1371
	var v1373 float64
	_ = v1373
	var v1374 float64
	_ = v1374
	var v1382 float64
	_ = v1382
	var v1384 float64
	_ = v1384
	var v1392 float64
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1431 float64
	_ = v1431
	var v1441 int32
	_ = v1441
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
	var v1460 int32
	_ = v1460
	var v1476 float64
	_ = v1476
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 float64
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1582 int32
	_ = v1582
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1909 int32
	_ = v1909
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1960 int32
	_ = v1960
	var v1969 int32
	_ = v1969
	var v1979 int32
	_ = v1979
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2004 int32
	_ = v2004
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2164 float64
	_ = v2164
	var v2165 float64
	_ = v2165
	var v2170 int32
	_ = v2170
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2240 int32
	_ = v2240
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2349 int32
	_ = v2349
	var v2361 int32
	_ = v2361
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2410 int32
	_ = v2410
	var v2418 int32
	_ = v2418
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2490 int32
	_ = v2490
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2515 int32
	_ = v2515
	var v2545 int32
	_ = v2545
	var v2559 int32
	_ = v2559
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
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
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2614 int32
	_ = v2614
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2642 int32
	_ = v2642
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2721 int32
	_ = v2721
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2763 int32
	_ = v2763
	var v2773 int32
	_ = v2773
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2815 int32
	_ = v2815
	var v2819 int32
	_ = v2819
	var v2825 int32
	_ = v2825
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2922 int32
	_ = v2922
	var v2939 int32
	_ = v2939
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2984 int32
	_ = v2984
	var v3026 int32
	_ = v3026
	var v3062 int32
	_ = v3062
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3075 int32
	_ = v3075
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3099 float64
	_ = v3099
	var v3100 float64
	_ = v3100
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3146 float64
	_ = v3146
	var v3147 float64
	_ = v3147
	var v3159 int32
	_ = v3159
	var v3163 int32
	_ = v3163
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3177 int32
	_ = v3177
	var v3187 float64
	_ = v3187
	var v3188 float64
	_ = v3188
	var v3194 int32
	_ = v3194
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3236 int32
	_ = v3236
	var v3251 int32
	_ = v3251
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3281 int32
	_ = v3281
	var v3284 int32
	_ = v3284
	var v3292 int32
	_ = v3292
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3331 int32
	_ = v3331
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3379 int32
	_ = v3379
	var v3382 float64
	_ = v3382
	var v3383 float64
	_ = v3383
	var v3399 int32
	_ = v3399
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3448 int32
	_ = v3448
	var v3455 int32
	_ = v3455
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3493 int32
	_ = v3493
	var v3534 int32
	_ = v3534
	var v3540 int32
	_ = v3540
	var v3542 int32
	_ = v3542
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3560 int32
	_ = v3560
	var v3564 int32
	_ = v3564
	var v3566 int32
	_ = v3566
	var v3572 int32
	_ = v3572
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3607 int32
	_ = v3607
	var v3630 int32
	_ = v3630
	var v3633 int32
	_ = v3633
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3638 int32
	_ = v3638
	var v3646 int32
	_ = v3646
	var v3650 int32
	_ = v3650
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3690 int32
	_ = v3690
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3775 int32
	_ = v3775
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3786 int32
	_ = v3786
	var v3790 int32
	_ = v3790
	var v3792 int32
	_ = v3792
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3803 int32
	_ = v3803
	var v3810 int32
	_ = v3810
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3828 int32
	_ = v3828
	var v3832 int32
	_ = v3832
	var v3835 int32
	_ = v3835
	var v3837 int32
	_ = v3837
	var v3840 int32
	_ = v3840
	var v3847 int32
	_ = v3847
	var v3849 int32
	_ = v3849
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3871 int32
	_ = v3871
	var v3890 int32
	_ = v3890
	var v3893 int32
	_ = v3893
	var v3919 int32
	_ = v3919
	var v3922 int32
	_ = v3922
	var v3933 int32
	_ = v3933
	var v3935 int32
	_ = v3935
	var v3937 int32
	_ = v3937
	var v3945 int32
	_ = v3945
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3971 int32
	_ = v3971
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3988 int32
	_ = v3988
	var v4010 int32
	_ = v4010
	var v4035 float64
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4052 int32
	_ = v4052
	var v4056 int32
	_ = v4056
	var v4059 int32
	_ = v4059
	var v4061 int32
	_ = v4061
	var v4064 int32
	_ = v4064
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4081 int32
	_ = v4081
	var v4082 int32
	_ = v4082
	var v4095 int32
	_ = v4095
	var v4140 int32
	_ = v4140
	var v4163 int32
	_ = v4163
	var v4184 float64
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4204 int32
	_ = v4204
	var v4232 float64
	_ = v4232
	var v4240 int32
	_ = v4240
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4246 float64
	_ = v4246
	var v4247 float64
	_ = v4247
	var v4248 int32
	_ = v4248
	var v4249 int32
	_ = v4249
	var v4250 int32
	_ = v4250
	var v4266 int32
	_ = v4266
	var v4290 float64
	_ = v4290
	var v4298 int32
	_ = v4298
	var v4299 int32
	_ = v4299
	var v4300 int32
	_ = v4300
	var v4301 float64
	_ = v4301
	var v4304 int32
	_ = v4304
	var v4305 float64
	_ = v4305
	var v4311 float64
	_ = v4311
	var v4365 int32
	_ = v4365
	var v4374 int32
	_ = v4374
	var v4384 int32
	_ = v4384
	var v4394 int32
	_ = v4394
	var v4410 float64
	_ = v4410
	var v4419 int32
	_ = v4419
	var v4420 float64
	_ = v4420
	var v4422 float64
	_ = v4422
	var v4424 float64
	_ = v4424
	var v4426 float64
	_ = v4426
	var v4427 float64
	_ = v4427
	var v4428 int32
	_ = v4428
	var v4429 int32
	_ = v4429
	var v4431 int32
	_ = v4431
	var v4441 int32
	_ = v4441
	var v4467 float64
	_ = v4467
	var v4482 int32
	_ = v4482
	var v4484 int32
	_ = v4484
	var v4508 float64
	_ = v4508
	var v4518 float64
	_ = v4518
	var v4519 float64
	_ = v4519
	var v4520 int32
	_ = v4520
	var v4523 int32
	_ = v4523
	var v4559 float64
	_ = v4559
	var v4566 float64
	_ = v4566
	var v4656 float64
	_ = v4656
	var v4658 int32
	_ = v4658
	var v4660 int32
	_ = v4660
	var v4697 float64
	_ = v4697
	var v4711 int32
	_ = v4711
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4751 int32
	_ = v4751
	var v4764 int32
	_ = v4764
	var v4789 float64
	_ = v4789
	var v4795 int32
	_ = v4795
	var v4797 int32
	_ = v4797
	var v4799 int32
	_ = v4799
	var v4801 int32
	_ = v4801
	var v4809 int32
	_ = v4809
	var v4821 int32
	_ = v4821
	var v4823 int32
	_ = v4823
	var v4837 float64
	_ = v4837
	var v4838 float64
	_ = v4838
	var v4844 int32
	_ = v4844
	var v4864 int32
	_ = v4864
	var v4866 int32
	_ = v4866
	var v4880 float64
	_ = v4880
	var v4881 float64
	_ = v4881
	var v4909 int32
	_ = v4909
	var v4925 float64
	_ = v4925
	v42 = m.G0
	v44 = v42 - int32(48)
	m.G0 = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l5)+112))
	if v64 == int32(0) {
		v1441 = l0
		v1442 = l1
		v1443 = l2
		v1444 = l3
		v1445 = l4
		v1446 = l5
		v1447 = l6
		v1448 = l7
		v1460 = v44
		v1476 = v63
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v1448 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 <= int32(0) {
		v1441 = l0
		v1442 = l1
		v1443 = l2
		v1444 = l3
		v1445 = l4
		v1446 = l5
		v1447 = l6
		v1448 = l7
		v1460 = v44
		v1476 = v63
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v82 = int32(0)
	goto L11
L11:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v71+v82<<(uint(int32(2))%32))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+16)))
	if v118 != int32(109) {
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
	v122 = v82 + int32(1)
	if v122 != v67 {
		v82 = v122
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
	v1441 = l0
	v1442 = l1
	v1443 = l2
	v1444 = l3
	v1445 = l4
	v1446 = l5
	v1447 = l6
	v1448 = l7
	v1460 = v44
	v1476 = v63
	goto L8
L17:
	;
	v269 = l0
	v270 = l1
	v271 = l2
	v272 = l3
	v273 = l4
	v274 = l5
	v275 = l6
	v276 = l7
	v288 = v44
	v293 = v252
	v295 = v254
	v296 = v255
	v298 = v70
	v304 = v63
	goto L36
L18:
	;
	v128 = F_palloc(m, int32(0))
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
	v136 = l1 + int32(4)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v140 = F_palloc(m, v137<<(uint(int32(2))%32))
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
	v133 = F_palloc(m, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v252 = v128
	v254 = v133
	v255 = int32(4)
	goto L17
L24:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v145 = F_palloc(m, v142<<(uint(int32(2))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v147 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v148 <= v147 {
		v252 = v140
		v254 = v145
		v255 = v136
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v160 = v147
	goto L27
L27:
	;
	v193 = v160 << (uint(int32(2)) % 32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193+v194)))
	v197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v203 = F_bms_is_member(m, v160, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v252 = v140
	v254 = v145
	v255 = v136
	goto L17
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193+v145))) = v222
	v225 = v160 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v225 < v226 {
		v160 = v225
		goto L27
	} else {
		goto L35
	}
L30:
	;
	v219 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v193+v140))) = v219
	v222 = v219
	goto L29
L31:
	;
	if v203 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v210 = F_statext_is_compatible_clause(m, l0, v196, v205, v44+int32(32), v44+int32(24))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v210 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v140))) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v222 = v217
	goto L29
L35:
	;
	goto L28
L36:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+20)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v274)+112))
	if v270 == int32(0) {
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
	v314 = int32(0)
	v319 = F_choose_best_statistics(m, v311, v310&int32(1), v293, v295, v314)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v324 = F_choose_best_statistics(m, v311, v310&int32(1), v293, v295, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L21
	} else {
		goto L44
	}
L42:
	;
	if v319 != 0 {
		v860 = v269
		v861 = v270
		v862 = v271
		v863 = v272
		v864 = v273
		v865 = v274
		v866 = v275
		v867 = v276
		v870 = v319
		v875 = v314
		v876 = v314
		v879 = v288
		v884 = v293
		v886 = v295
		v887 = v296
		v889 = v298
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v1441 = v269
	v1442 = v270
	v1443 = v271
	v1444 = v272
	v1445 = v273
	v1446 = v274
	v1447 = v275
	v1448 = v276
	v1460 = v288
	v1476 = v304
	goto L8
L44:
	;
	if v324 == int32(0) {
		v1441 = v269
		v1442 = v270
		v1443 = v271
		v1444 = v272
		v1445 = v273
		v1446 = v274
		v1447 = v275
		v1448 = v276
		v1460 = v288
		v1476 = v304
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v328 = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v329 <= v328 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v860 = v269
	v861 = v270
	v862 = v271
	v863 = v272
	v864 = v273
	v865 = v274
	v866 = v275
	v867 = v276
	v870 = v324
	v875 = int32(0)
	v876 = v328
	v879 = v288
	v884 = v293
	v886 = v295
	v887 = v296
	v889 = v298
	goto L38
L47:
	;
	goto L48
L48:
	;
	v334 = int32(0)
	v350 = v334
	v351 = v334
	v352 = v328
	v357 = int32(-1)
	goto L49
L49:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v379 = v357 + int32(1)
	v381 = v379 << (uint(int32(2)) % 32)
	v382 = v293 + v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v383 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v860 = v269
	v861 = v270
	v862 = v271
	v863 = v272
	v864 = v273
	v865 = v274
	v866 = v275
	v867 = v276
	v870 = v324
	v875 = v830
	v876 = v829
	v879 = v288
	v884 = v293
	v886 = v295
	v887 = v296
	v889 = v298
	goto L38
L51:
	;
	v857 = v352 + int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v857 < v858 {
		v350 = v829
		v351 = v830
		v352 = v857
		v357 = v379
		goto L49
	} else {
		goto L125
	}
L52:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v381+v295)))
	if v387 == int32(0) {
		v829 = v350
		v830 = v351
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v324)+20))
	v391 = int32(0)
	if v383 == v391 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	if v444 == int32(0) {
		v829 = v350
		v830 = v351
		goto L51
	} else {
		goto L70
	}
L57:
	;
	v444 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v390 == int32(0) {
		v435 = v391
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v444 = v435
	goto L56
L61:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v401 < v400 {
		v435 = v391
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v403 = int32(1)
	if v400 <= v403 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v406 = v403
	goto L65
L64:
	;
	v406 = v400
	goto L65
L65:
	;
	v407 = int32(8)
	v412 = int32(0)
	goto L66
L66:
	;
	v419 = v412 << (uint(int32(2)) % 32)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v383+v407+v419)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419+(v390+v407))))
	v426 = v421 & (v423 ^ int32(-1))
	v428 = base.B2i32(v426 == int32(0))
	if v426 != 0 {
		v435 = v428
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v435 = v428
	goto L60
L68:
	;
	v430 = v412 + int32(1)
	if v430 != v406 {
		v412 = v430
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v447 = v381 + v295
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if v448 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v377+v352<<(uint(int32(2))%32))))
	v799 = F_lappend(m, v351, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L21
	} else {
		goto L121
	}
L72:
	;
	if v351 != 0 {
		goto L117
	} else {
		goto L118
	}
L73:
	;
	if v627 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L74:
	;
	v449 = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v449 < v450 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v627 = v617
	goto L73
L77:
	;
	v478 = v449
	goto L80
L78:
	;
	v582 = v448
	goto L79
L79:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v606 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v324)+24))
	if v494 == int32(0) {
		v829 = v350
		v830 = v351
		goto L51
	} else {
		goto L82
	}
L81:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v582 = v564
	goto L79
L82:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v497 <= int32(0) {
		v829 = v350
		v830 = v351
		goto L51
	} else {
		goto L83
	}
L83:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v478<<(uint(int32(2))%32))))
	v515 = int32(0)
	goto L84
L84:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v547+v515<<(uint(int32(2))%32))))
	v552 = F_equal(m, v551, v504)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L21
	} else {
		goto L86
	}
L85:
	;
	v561 = v478 + int32(1)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v561 < v562 {
		v478 = v561
		goto L80
	} else {
		goto L91
	}
L86:
	;
	if v552 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v557 = v515 + int32(1)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v557 < v558 {
		v515 = v557
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
	v829 = v350
	v830 = v351
	goto L51
L91:
	;
	goto L81
L92:
	;
	if v582 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v582 == int32(0) {
		v627 = v606
		goto L73
	} else {
		goto L99
	}
L95:
	;
	v627 = int32(0)
	goto L73
L96:
	;
	goto L97
L97:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v612 == int32(1) {
		goto L72
	} else {
		goto L98
	}
L98:
	;
	v768 = v350
	goto L71
L99:
	;
	v768 = v350
	goto L71
L100:
	;
	if v705 != int32(1) {
		v768 = v350
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
	v667 = int32(1)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	if v668 <= v667 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v671 = v667
	goto L106
L105:
	;
	v671 = v668
	goto L106
L106:
	;
	v674 = int32(0)
	v676 = v674
	v677 = v674
	goto L107
L107:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v627+int32(8)+v676<<(uint(int32(2))%32))))
	if v685 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v705 = v698
	goto L100
L109:
	;
	goto L108
L110:
	;
	v686 = int32(2)
	if v677 != 0 {
		v698 = v686
		goto L109
	} else {
		goto L113
	}
L111:
	;
	v691 = v677
	goto L112
L112:
	;
	v694 = v676 + int32(1)
	if v694 != v671 {
		v676 = v694
		v677 = v691
		goto L107
	} else {
		goto L115
	}
L113:
	;
	v687 = int32(1)
	if base.Ui32(v687) < base.Ui32(base.I32_popcnt(v685)) {
		v698 = v686
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v691 = v687
	goto L112
L115:
	;
	v698 = v691
	goto L109
L116:
	;
	goto L72
L117:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v751 = v749
	goto L119
L118:
	;
	v751 = int32(0)
	goto L119
L119:
	;
	v752 = F_bms_add_member(m, v350, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	v768 = v752
	goto L71
L121:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v802 = F_bms_add_member(m, v801, v379)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v802
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v382))) = int32(0)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = int32(0)
	v829 = v768
	v830 = v799
	goto L51
L125:
	;
	goto L50
L126:
	;
	v904 = F_clauselist_selectivity_ext(m, v860, v875, v862, v863, v864, int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v879)+44)) = int32(0)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889)+20)))
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
	v907 = v879 + int32(32)
	v909 = v879 + int32(24)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v860)+36))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v865)+68))
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
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v879)+32))
	v1043 = *(*float64)(unsafe.Add(mBase, uint32(v879)+24))
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
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v870)+20))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v870)+24))
	v925 = F_mcv_get_match_bitmap(m, v875, v922, v923, v919, v921)
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
	v950 = v921
	v972 = float64(0)
	goto L134
L134:
	;
	v980 = v919 + int32(48) + v950*int32(24)
	v981 = *(*float64)(unsafe.Add(mBase, uint32(v980)))
	v982 = *(*float64)(unsafe.Add(mBase, uint32(v909)))
	*(*float64)(unsafe.Add(mBase, uint32(v909))) = base.F64_add(v981, v982)
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950+v925))))
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
	v995 = base.F64_add(v972, v993)
	goto L138
L137:
	;
	v995 = v972
	goto L138
L138:
	;
	v997 = v950 + int32(1)
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v919)+8))
	if base.Ui32(v997) < base.Ui32(v998) {
		v950 = v997
		v972 = v995
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	v269 = v860
	v270 = v861
	v271 = v862
	v272 = v863
	v273 = v864
	v274 = v865
	v275 = v866
	v276 = v867
	v288 = v879
	v293 = v884
	v295 = v886
	v296 = v887
	v298 = v889
	v304 = base.F64_mul(v304, v1067)
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
	if v875 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v269 = v860
	v270 = v861
	v271 = v862
	v272 = v863
	v273 = v864
	v274 = v865
	v275 = v866
	v276 = v867
	v288 = v879
	v293 = v884
	v295 = v886
	v296 = v887
	v298 = v889
	v304 = base.F64_sub(base.F64_add(v304, v1431), base.F64_mul(v304, v1431))
	goto L36
L152:
	;
	v1094 = v1075
	v1119 = v1076
	v1121 = v1076
	goto L157
L153:
	;
	v1075 = int32(0)
	v1076 = float64(0)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
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
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v875)+12))
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
	v1431 = v1392
	goto L151
L159:
	;
	v1147 = v879 + int32(32)
	v1149 = v879 + int32(24)
	v1150 = int32(16)
	v1151 = v879 + v1150
	v1153 = v879 + int32(8)
	v1154 = float64(0)
	v1155 = m.G0
	v1157 = v1155 - v1150
	m.G0 = v1157
	v1160 = v879 + int32(44)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1160)))
	if v1161 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v1135 = base.F64_mul(v1121, v1133)
	v1137 = base.F64_add(v1121, base.F64_sub(v1133, v1135))
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
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	v1165 = F_palloc0(m, v1164)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L21
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+8)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+12)) = v1131
	v1174 = F_list_make1_impl(m, int32(1), v1157+int32(8))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L21
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1160))) = v1165
	goto L165
L167:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v870)+20))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v870)+24))
	v1179 = F_mcv_get_match_bitmap(m, v1174, v1176, v1177, v1073, int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L21
	} else {
		goto L168
	}
L168:
	;
	v1181 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1147))) = v1181
	*(*int64)(unsafe.Add(mBase, uint32(v1149))) = v1181
	*(*int64)(unsafe.Add(mBase, uint32(v1151))) = v1181
	*(*int64)(unsafe.Add(mBase, uint32(v1153))) = v1181
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	if v1189 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1203 = int32(0)
	v1230 = v1154
	goto L172
L170:
	;
	v1318 = v1154
	goto L171
L171:
	;
	F_pfree(m, v1179)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L21
	} else {
		goto L181
	}
L172:
	;
	v1235 = v1073 + int32(48) + v1203*int32(24)
	v1236 = *(*float64)(unsafe.Add(mBase, uint32(v1235)))
	v1237 = *(*float64)(unsafe.Add(mBase, uint32(v1153)))
	*(*float64)(unsafe.Add(mBase, uint32(v1153))) = base.F64_add(v1236, v1237)
	v1240 = v1203 + v1179
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240))))
	if v1241 != int32(1) {
		v1263 = v1230
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1318 = v1263
	goto L171
L174:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1160)))
	v1267 = v1266 + v1203
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267))))
	if v1268 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v1244 = *(*float64)(unsafe.Add(mBase, uint32(v1235)))
	v1245 = *(*float64)(unsafe.Add(mBase, uint32(v1235)+8))
	v1246 = *(*float64)(unsafe.Add(mBase, uint32(v1147)))
	*(*float64)(unsafe.Add(mBase, uint32(v1147))) = base.F64_add(v1245, v1246)
	v1249 = base.F64_add(v1230, v1244)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1160)))
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250+v1203))))
	if v1252 != int32(1) {
		v1263 = v1249
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v1255 = *(*float64)(unsafe.Add(mBase, uint32(v1235)))
	v1256 = *(*float64)(unsafe.Add(mBase, uint32(v1149)))
	*(*float64)(unsafe.Add(mBase, uint32(v1149))) = base.F64_add(v1255, v1256)
	v1259 = *(*float64)(unsafe.Add(mBase, uint32(v1235)+8))
	v1260 = *(*float64)(unsafe.Add(mBase, uint32(v1151)))
	*(*float64)(unsafe.Add(mBase, uint32(v1151))) = base.F64_add(v1259, v1260)
	v1263 = v1249
	goto L174
L177:
	;
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240))))
	v1272 = v1271
	goto L179
L178:
	;
	v1272 = int32(1)
	goto L179
L179:
	;
	v1273 = int32(1)
	v1274 = v1272 & v1273
	*(*uint8)(unsafe.Add(mBase, uint32(v1267))) = uint8(v1274)
	v1277 = v1203 + v1273
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	if base.Ui32(v1277) < base.Ui32(v1278) {
		v1203 = v1277
		v1230 = v1263
		goto L172
	} else {
		goto L180
	}
L180:
	;
	goto L173
L181:
	;
	m.G0 = v1157 + int32(16)
	v1327 = F_bms_is_member(m, v1094, v876)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L21
	} else {
		goto L183
	}
L182:
	;
	v1394 = v1094 + int32(1)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v1394 < v1395 {
		v1094 = v1394
		v1119 = v1392
		v1121 = v1145
		goto L157
	} else {
		goto L209
	}
L183:
	;
	if v1327 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1355 = v1133
	goto L186
L185:
	;
	v1329 = *(*float64)(unsafe.Add(mBase, uint32(v879)+32))
	v1330 = *(*float64)(unsafe.Add(mBase, uint32(v879)+8))
	v1331 = float64(0)
	v1333 = base.F64_sub(v1133, v1329)
	if base.F64_lt(v1333, v1331) != 0 {
		v1341 = v1331
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1356 = *(*float64)(unsafe.Add(mBase, uint32(v879)+24))
	v1357 = *(*float64)(unsafe.Add(mBase, uint32(v879)+16))
	v1358 = *(*float64)(unsafe.Add(mBase, uint32(v879)+8))
	v1359 = float64(0)
	v1361 = base.F64_sub(v1135, v1357)
	if base.F64_lt(v1361, v1359) != 0 {
		v1369 = v1359
		goto L198
	} else {
		goto L199
	}
L187:
	;
	v1355 = v1354
	goto L186
L188:
	;
	v1343 = base.F64_sub(float64(1), v1330)
	if base.F64_lt(v1343, v1341) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if base.F64_gt(v1333, float64(1)) == int32(0) {
		v1341 = v1333
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1341 = float64(1)
	goto L188
L191:
	;
	goto L187
L192:
	;
	v1345 = v1343
	goto L194
L193:
	;
	v1345 = v1341
	goto L194
L194:
	;
	v1346 = base.F64_add(v1318, v1345)
	if base.F64_lt(v1346, float64(0)) != 0 {
		v1354 = v1331
		goto L191
	} else {
		goto L195
	}
L195:
	;
	if base.F64_gt(v1346, float64(1)) == int32(0) {
		v1354 = v1346
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v1354 = float64(1)
	goto L191
L197:
	;
	v1384 = base.F64_add(v1119, base.F64_sub(v1355, v1382))
	if base.F64_lt(v1384, float64(0)) != 0 {
		v1392 = float64(0)
		goto L182
	} else {
		goto L207
	}
L198:
	;
	v1371 = base.F64_sub(float64(1), v1358)
	if base.F64_lt(v1371, v1369) != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	if base.F64_gt(v1361, float64(1)) == int32(0) {
		v1369 = v1361
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1369 = float64(1)
	goto L198
L201:
	;
	goto L197
L202:
	;
	v1373 = v1371
	goto L204
L203:
	;
	v1373 = v1369
	goto L204
L204:
	;
	v1374 = base.F64_add(v1356, v1373)
	if base.F64_lt(v1374, float64(0)) != 0 {
		v1382 = v1359
		goto L201
	} else {
		goto L205
	}
L205:
	;
	if base.F64_gt(v1374, float64(1)) == int32(0) {
		v1382 = v1374
		goto L201
	} else {
		goto L206
	}
L206:
	;
	v1382 = float64(1)
	goto L201
L207:
	;
	if base.F64_gt(v1384, float64(1)) == int32(0) {
		v1392 = v1384
		goto L182
	} else {
		goto L208
	}
L208:
	;
	v1392 = float64(1)
	goto L182
L209:
	;
	goto L158
L210:
	;
	v1484 = int32(0)
	v1488 = m.G0
	v1490 = v1488 - int32(16)
	m.G0 = v1490
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+36))
	if v1492 != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v4909 = v1460
	v4925 = v1476
	goto L212
L212:
	;
	m.G0 = v4909 + int32(48)
	return v4925
L213:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1506)))
	v1508 = float64(1)
	v1509 = int32(0)
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+112))
	if v1511 == v1509 {
		v1656 = v1509
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+68))
	v1506 = v1492 + v1493<<(uint(int32(2))%32)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+52))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1498)+12))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+68))
	v1506 = v1499 + v1500<<(uint(int32(2))%32) - int32(4)
	goto L213
L217:
	;
	if v1656 != 0 {
		goto L228
	} else {
		goto L229
	}
L218:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+4))
	if v1514 <= int32(0) {
		v1582 = v1509
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1656 = v1582
	goto L217
L220:
	;
	v1517 = int32(0)
	if v1517 < v1514 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1520 = v1514
	goto L223
L222:
	;
	v1520 = v1517
	goto L223
L223:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+12))
	v1523 = int32(0)
	goto L224
L224:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1521+v1523<<(uint(int32(2))%32))))
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1567)+16)))
	v1570 = base.B2i32(v1568 == int32(102))
	if v1568 == int32(102) {
		v1582 = v1570
		goto L219
	} else {
		goto L226
	}
L225:
	;
	v1582 = v1570
	goto L219
L226:
	;
	v1572 = v1523 + int32(1)
	if v1572 != v1520 {
		v1523 = v1572
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	if v1442 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v4864 = v1460
	v4866 = v1490
	v4880 = v1476
	v4881 = v1508
	goto L230
L230:
	;
	m.G0 = v4866 + int32(16)
	v4909 = v4864
	v4925 = base.F64_mul(v4880, v4881)
	goto L212
L231:
	;
	v1994 = int32(0)
	v1995 = int32(16)
	if v1994 < v1954 {
		goto L263
	} else {
		goto L264
	}
L232:
	;
	v1662 = F_palloc(m, int32(0))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L21
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1668 = v1442 + int32(4)
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	v1672 = F_palloc(m, v1669<<(uint(int32(1))%32))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L21
	} else {
		goto L237
	}
L235:
	;
	v1665 = F_palloc(m, int32(0))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L21
	} else {
		goto L236
	}
L236:
	;
	v1954 = v1484
	v1960 = v1665
	v1969 = v1662
	v1979 = v1442 + int32(4)
	goto L231
L237:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	v1677 = F_palloc(m, v1674<<(uint(int32(2))%32))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L21
	} else {
		goto L238
	}
L238:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	if v1679 <= int32(0) {
		v1954 = v1484
		v1960 = v1677
		v1969 = v1672
		v1979 = v1668
		goto L231
	} else {
		goto L239
	}
L239:
	;
	v1684 = v1484
	v1691 = int32(0)
	goto L240
L240:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+12))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1724+v1691<<(uint(int32(2))%32))))
	v1729 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+8)) = v1729
	v1733 = v1672 + v1691<<(uint(int32(1))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1733))) = uint16(v1729)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1447)))
	v1737 = F_bms_is_member(m, v1691, v1736)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L21
	} else {
		goto L243
	}
L241:
	;
	v1954 = v1909
	v1960 = v1677
	v1969 = v1672
	v1979 = v1668
	goto L231
L242:
	;
	v1950 = v1691 + int32(1)
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1668)))
	if v1950 < v1951 {
		v1684 = v1909
		v1691 = v1950
		goto L240
	} else {
		goto L262
	}
L243:
	;
	if v1737 != 0 {
		v1909 = v1684
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+68))
	v1742 = F_dependency_is_compatible_clause(m, v1728, v1739, v1490+int32(14))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L21
	} else {
		goto L246
	}
L245:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1733))) = uint16(v1866)
	v1909 = v1867
	goto L242
L246:
	;
	if v1742 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1744 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1490)+14)))
	v1866 = v1744
	v1867 = v1684
	goto L245
L248:
	;
	goto L249
L249:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+112))
	v1748 = F_dependency_is_compatible_expression(m, v1728, v1745, v1490+int32(8))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L21
	} else {
		goto L250
	}
L250:
	;
	if v1748 == int32(0) {
		v1909 = v1684
		goto L242
	} else {
		goto L251
	}
L251:
	;
	v1752 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1490)+14)) = uint16(v1752)
	if v1684 <= v1752 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1677+v1684<<(uint(int32(2))%32)))) = v1860
	v1866 = v1684 ^ int32(-1)
	v1867 = v1684 + int32(1)
	goto L245
L253:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+8))
	v1758 = int32(0)
	goto L254
L254:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1677+v1758<<(uint(int32(2))%32))))
	v1803 = F_equal(m, v1802, v1757)
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L21
	} else {
		goto L256
	}
L255:
	;
	v1810 = int32(65535)
	if v1758&v1810 == v1810 {
		goto L252
	} else {
		goto L261
	}
L256:
	;
	if v1803 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1808 = v1758 + int32(1)
	if v1808 != v1684 {
		v1758 = v1808
		goto L254
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	goto L255
L260:
	;
	goto L252
L261:
	;
	v1866 = v1758 ^ int32(-1)
	v1867 = v1684
	goto L245
L262:
	;
	goto L241
L263:
	;
	v2004 = (v1954<<(uint(v1995)%32) + int32(65536)) >> (uint(v1995) % 32)
	goto L265
L264:
	;
	v2004 = v1994
	goto L265
L265:
	;
	v2014 = v1994
	v2017 = int32(0)
	goto L266
L266:
	;
	if v1442 != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	if v2017 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L268:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1979)))
	v2049 = v2047
	goto L270
L269:
	;
	v2049 = int32(0)
	goto L270
L270:
	;
	if v2014 < v2049 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v2053 = v1969 + v2014<<(uint(int32(1))%32)
	v2054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2053))))
	if v2054 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	goto L273
L273:
	;
	goto L267
L274:
	;
	v2055 = v2054 + v2004
	*(*uint16)(unsafe.Add(mBase, uint32(v2053))) = uint16(v2055)
	v2058 = F_bms_add_member(m, v2017, base.I32_extend16_s(v2055))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L21
	} else {
		goto L277
	}
L275:
	;
	v2060 = v2017
	goto L276
L276:
	;
	v2014 = v2014 + int32(1)
	v2017 = v2060
	goto L266
L277:
	;
	v2060 = v2058
	goto L276
L278:
	;
	F_pfree(m, v4809)
	mBase = m.M
	v4844 = m.ExcPending
	if v4844 != 0 {
		goto L21
	} else {
		goto L599
	}
L279:
	;
	if v2110 != int32(2) {
		goto L295
	} else {
		goto L296
	}
L280:
	;
	v2110 = int32(0)
	goto L279
L281:
	;
	goto L282
L282:
	;
	v2072 = int32(1)
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2073 <= v2072 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v2076 = v2072
	goto L285
L284:
	;
	v2076 = v2073
	goto L285
L285:
	;
	v2079 = int32(0)
	v2081 = v2079
	v2082 = v2079
	goto L286
L286:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2017+int32(8)+v2081<<(uint(int32(2))%32))))
	if v2090 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v2110 = v2103
	goto L279
L288:
	;
	goto L287
L289:
	;
	v2091 = int32(2)
	if v2082 != 0 {
		v2103 = v2091
		goto L288
	} else {
		goto L292
	}
L290:
	;
	v2096 = v2082
	goto L291
L291:
	;
	v2099 = v2081 + int32(1)
	if v2099 != v2076 {
		v2081 = v2099
		v2082 = v2096
		goto L286
	} else {
		goto L294
	}
L292:
	;
	v2092 = int32(1)
	if base.Ui32(v2092) < base.Ui32(base.I32_popcnt(v2090)) {
		v2103 = v2091
		goto L288
	} else {
		goto L293
	}
L293:
	;
	v2096 = v2092
	goto L291
L294:
	;
	v2103 = v2096
	goto L288
L295:
	;
	F_bms_free(m, v2017)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L21
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+112))
	if v2115 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v4809 = v1969
	v4821 = v1460
	v4823 = v1490
	v4837 = v1476
	v4838 = v1508
	goto L278
L299:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+4))
	v2120 = v2116 << (uint(int32(2)) % 32)
	goto L301
L300:
	;
	v2120 = int32(0)
	goto L301
L301:
	;
	v2121 = F_palloc(m, v2120)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L21
	} else {
		goto L302
	}
L302:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+112))
	if v2123 != 0 {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	v3199 = F_palloc(m, v3135)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L21
	} else {
		goto L420
	}
L304:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2123)+4))
	if int32(0) < v2124 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	v3159 = v1960
	v3163 = v2017
	v3168 = v1969
	v3171 = v1460
	v3173 = v1490
	v3177 = v2121
	v3187 = v1476
	v3188 = v1508
	goto L306
L306:
	;
	F_pfree(m, v3177)
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L21
	} else {
		goto L417
	}
L307:
	;
	v2130 = v1954
	v2131 = v1443
	v2132 = v1444
	v2133 = v1445
	v2134 = v1441
	v2135 = v1447
	v2136 = v1960
	v2140 = v2017
	v2143 = v1442
	v2144 = v2004
	v2145 = v1969
	v2148 = v1460
	v2150 = v1490
	v2152 = v1484
	v2153 = v1484
	v2154 = v2121
	v2155 = v1979
	v2156 = v2123
	v2157 = v1507
	v2158 = v1484
	v2160 = base.B2i32(int32(0) < v1954)
	v2164 = v1476
	v2165 = v1508
	goto L310
L308:
	;
	v3113 = v1443
	v3114 = v1444
	v3115 = v1445
	v3116 = v1441
	v3117 = v1447
	v3118 = v1960
	v3122 = v2017
	v3125 = v1442
	v3127 = v1969
	v3130 = v1460
	v3132 = v1490
	v3134 = v1484
	v3135 = v1484
	v3136 = v2121
	v3137 = v1979
	v3146 = v1476
	v3147 = v1508
	goto L309
L309:
	;
	if v3134 != 0 {
		goto L303
	} else {
		goto L416
	}
L310:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2156)+12))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2170+v2153<<(uint(int32(2))%32))))
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174)+16)))
	if v2175 != int32(102) {
		v3065 = v2130
		v3066 = v2131
		v3067 = v2132
		v3068 = v2133
		v3069 = v2134
		v3070 = v2135
		v3071 = v2136
		v3075 = v2140
		v3078 = v2143
		v3079 = v2144
		v3080 = v2145
		v3083 = v2148
		v3085 = v2150
		v3087 = v2152
		v3089 = v2154
		v3090 = v2155
		v3091 = v2156
		v3092 = v2157
		v3093 = v2158
		v3095 = v2160
		v3099 = v2164
		v3100 = v2165
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v3113 = v3066
	v3114 = v3067
	v3115 = v3068
	v3116 = v3069
	v3117 = v3070
	v3118 = v3071
	v3122 = v3075
	v3125 = v3078
	v3127 = v3080
	v3130 = v3083
	v3132 = v3085
	v3134 = v3087
	v3135 = v3093 << (uint(int32(2)) % 32)
	v3136 = v3089
	v3137 = v3090
	v3146 = v3099
	v3147 = v3100
	goto L309
L312:
	;
	v3106 = v2153 + int32(1)
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v3091)+4))
	if v3106 < v3107 {
		v2130 = v3065
		v2131 = v3066
		v2132 = v3067
		v2133 = v3068
		v2134 = v3069
		v2135 = v3070
		v2136 = v3071
		v2140 = v3075
		v2143 = v3078
		v2144 = v3079
		v2145 = v3080
		v2148 = v3083
		v2150 = v3085
		v2152 = v3087
		v2153 = v3106
		v2154 = v3089
		v2155 = v3090
		v2156 = v3091
		v2157 = v3092
		v2158 = v3093
		v2160 = v3095
		v2164 = v3099
		v2165 = v3100
		goto L310
	} else {
		goto L415
	}
L313:
	;
	v2178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174)+8)))
	v2179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2157)+20)))
	if v2178 != v2179 {
		v3065 = v2130
		v3066 = v2131
		v3067 = v2132
		v3068 = v2133
		v3069 = v2134
		v3070 = v2135
		v3071 = v2136
		v3075 = v2140
		v3078 = v2143
		v3079 = v2144
		v3080 = v2145
		v3083 = v2148
		v3085 = v2150
		v3087 = v2152
		v3089 = v2154
		v3090 = v2155
		v3091 = v2156
		v3092 = v2157
		v3093 = v2158
		v3095 = v2160
		v3099 = v2164
		v3100 = v2165
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v2181 = int32(0)
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+20))
	if v2183 == v2181 {
		goto L317
	} else {
		goto L318
	}
L315:
	;
	if int32(0) <= v2240 {
		goto L326
	} else {
		goto L327
	}
L316:
	;
	v2240 = base.I32_ctz(v2226) | v2227<<(uint(int32(5))%32)
	goto L315
L317:
	;
	v2240 = int32(-2)
	goto L315
L318:
	;
	v2193 = base.I32_div_s(int32(0), int32(32))
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+4))
	if v2194 <= v2193 {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v2197 = v2183 + int32(8)
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v2197+v2193<<(uint(int32(2))%32))))
	v2204 = v2201 & int32(-1)
	if v2204 != 0 {
		v2226 = v2204
		v2227 = v2193
		goto L316
	} else {
		goto L320
	}
L320:
	;
	v2206 = v2193 + int32(1)
	if v2206 == v2194 {
		goto L317
	} else {
		goto L321
	}
L321:
	;
	v2209 = v2206
	goto L322
L322:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2197+v2209<<(uint(int32(2))%32))))
	if v2216 != 0 {
		v2226 = v2216
		v2227 = v2209
		goto L316
	} else {
		goto L324
	}
L323:
	;
	goto L317
L324:
	;
	v2218 = v2209 + int32(1)
	if v2218 != v2194 {
		v2209 = v2218
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	v2251 = v2240
	v2252 = v2181
	goto L329
L327:
	;
	v2361 = v2181
	goto L328
L328:
	;
	v2393 = int32(0)
	v2394 = base.B2i32(v2130 <= v2393)
	if v2394 == v2393 {
		goto L347
	} else {
		goto L348
	}
L329:
	;
	if int32(0) < base.I32_extend16_s(v2251) {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v2361 = v2292
	goto L328
L331:
	;
	v2289 = F_bms_is_member(m, base.I32_extend16_s(v2251+v2144), v2140)
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L21
	} else {
		goto L334
	}
L332:
	;
	v2292 = v2252
	goto L333
L333:
	;
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+20))
	if v2293 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L334:
	;
	v2292 = v2289 + v2252
	goto L333
L335:
	;
	if int32(0) <= v2349 {
		v2251 = v2349
		v2252 = v2292
		goto L329
	} else {
		goto L346
	}
L336:
	;
	v2349 = base.I32_ctz(v2335) | v2336<<(uint(int32(5))%32)
	goto L335
L337:
	;
	v2349 = int32(-2)
	goto L335
L338:
	;
	v2300 = v2251 + int32(1)
	v2302 = base.I32_div_s(v2300, int32(32))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2293)+4))
	if v2303 <= v2302 {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v2306 = v2293 + int32(8)
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2306+v2302<<(uint(int32(2))%32))))
	v2313 = v2310 & (int32(-1) << (uint(v2300) % 32))
	if v2313 != 0 {
		v2335 = v2313
		v2336 = v2302
		goto L336
	} else {
		goto L340
	}
L340:
	;
	v2315 = v2302 + int32(1)
	if v2315 == v2303 {
		goto L337
	} else {
		goto L341
	}
L341:
	;
	v2318 = v2315
	goto L342
L342:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2306+v2318<<(uint(int32(2))%32))))
	if v2325 != 0 {
		v2335 = v2325
		v2336 = v2318
		goto L336
	} else {
		goto L344
	}
L343:
	;
	goto L337
L344:
	;
	v2327 = v2318 + int32(1)
	if v2327 != v2303 {
		v2318 = v2327
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	goto L330
L347:
	;
	v2410 = v2181
	v2418 = int32(0)
	goto L350
L348:
	;
	v2559 = v2181
	goto L349
L349:
	;
	if v2361+v2559 < int32(2) {
		v3065 = v2130
		v3066 = v2131
		v3067 = v2132
		v3068 = v2133
		v3069 = v2134
		v3070 = v2135
		v3071 = v2136
		v3075 = v2140
		v3078 = v2143
		v3079 = v2144
		v3080 = v2145
		v3083 = v2148
		v3085 = v2150
		v3087 = v2152
		v3089 = v2154
		v3090 = v2155
		v3091 = v2156
		v3092 = v2157
		v3093 = v2158
		v3095 = v2160
		v3099 = v2164
		v3100 = v2165
		goto L312
	} else {
		goto L360
	}
L350:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+24))
	if v2439 == int32(0) {
		v2515 = v2410
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v2559 = v2515
	goto L349
L352:
	;
	v2545 = v2418 + int32(1)
	if v2545 != v2130 {
		v2410 = v2515
		v2418 = v2545
		goto L350
	} else {
		goto L359
	}
L353:
	;
	v2442 = int32(0)
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+4))
	if v2443 <= v2442 {
		v2515 = v2410
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v2457 = v2442
	v2461 = v2410
	goto L355
L355:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+12))
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v2490+v2457<<(uint(int32(2))%32))))
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2136+v2418<<(uint(int32(2))%32))))
	v2496 = F_equal(m, v2494, v2495)
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L21
	} else {
		goto L357
	}
L356:
	;
	v2515 = v2498
	goto L352
L357:
	;
	v2498 = v2496 + v2461
	v2500 = v2457 + int32(1)
	v2501 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+4))
	if v2500 < v2501 {
		v2457 = v2500
		v2461 = v2498
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	goto L351
L360:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+4))
	v2592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2157)+20)))
	v2593 = m.G0
	v2595 = v2593 - int32(32)
	m.G0 = v2595
	v2598 = F_SearchSysCache2(m, int32(62), v2591, v2592)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L21
	} else {
		goto L363
	}
L361:
	;
	if v2160 != 0 {
		goto L379
	} else {
		goto L380
	}
L362:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2634 = m.ExcPending
	if v2634 != 0 {
		goto L21
	} else {
		goto L375
	}
L363:
	;
	if v2598 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v2604 = F_SysCacheGetAttr(m, int32(62), v2598, int32(4), v2595+int32(31))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L21
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L21
	} else {
		goto L372
	}
L367:
	;
	v2606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2595)+31)))
	if v2606 == int32(1) {
		goto L362
	} else {
		goto L368
	}
L368:
	;
	v2609 = F_pg_detoast_datum_packed(m, v2604)
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L21
	} else {
		goto L369
	}
L369:
	;
	v2611 = F_statext_dependencies_deserialize(m, v2609)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L21
	} else {
		goto L370
	}
L370:
	;
	F_ReleaseCatCache(m, v2598)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L21
	} else {
		goto L371
	}
L371:
	;
	m.G0 = v2595 + int32(32)
	goto L361
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2595))) = v2591
	F_errmsg_internal(m, int32(41609), v2595)
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L21
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(489220), int32(630), int32(459148))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L21
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2595)+20)) = v2591
	*(*int32)(unsafe.Add(mBase, uint32(v2595)+16)) = int32(102)
	F_errmsg_internal(m, int32(41490), v2595+int32(16))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L21
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(489220), int32(637), int32(459148))
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L21
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	if v3026 == int32(0) {
		v3065 = v2130
		v3066 = v2131
		v3067 = v2132
		v3068 = v2133
		v3069 = v2134
		v3070 = v2135
		v3071 = v2136
		v3075 = v2140
		v3078 = v2143
		v3079 = v2144
		v3080 = v2145
		v3083 = v2148
		v3085 = v2150
		v3087 = v2152
		v3089 = v2154
		v3090 = v2155
		v3091 = v2156
		v3092 = v2157
		v3093 = v2158
		v3095 = v2160
		v3099 = v2164
		v3100 = v2165
		goto L312
	} else {
		goto L414
	}
L379:
	;
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+8))
	if v2650 == int32(0) {
		goto L383
	} else {
		goto L384
	}
L380:
	;
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+24))
	if v2648 != 0 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+8))
	v3026 = v2649
	goto L378
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2611)+8)) = v2984
	v3026 = v2984
	goto L378
L383:
	;
	v2984 = int32(0)
	goto L382
L384:
	;
	goto L385
L385:
	;
	v2655 = v2611 + int32(12)
	v2656 = int32(0)
	v2667 = v2656
	v2671 = v2656
	goto L386
L386:
	;
	v2701 = v2655 + v2667<<(uint(int32(2))%32)
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2701)))
	v2703 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2702)+8)))
	if int32(0) < v2703 {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	v2984 = v2939
	goto L382
L388:
	;
	v2968 = v2667 + int32(1)
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+8))
	if base.Ui32(v2968) < base.Ui32(v2969) {
		v2667 = v2968
		v2671 = v2939
		goto L386
	} else {
		goto L413
	}
L389:
	;
	v2721 = int32(0)
	goto L392
L390:
	;
	goto L391
L391:
	;
	if v2667 != v2671 {
		goto L410
	} else {
		goto L411
	}
L392:
	;
	v2752 = v2702 + int32(10) + v2721<<(uint(int32(1))%32)
	v2753 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2752))))
	if v2753 <= int32(0) {
		goto L395
	} else {
		goto L396
	}
L393:
	;
	goto L391
L394:
	;
	v2874 = v2721 + int32(1)
	v2875 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2702)+8)))
	if v2874 < v2875 {
		v2721 = v2874
		goto L392
	} else {
		goto L409
	}
L395:
	;
	if v2130 <= v2393 {
		v2939 = v2671
		goto L388
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v2825 = v2753 + v2144
	*(*uint16)(unsafe.Add(mBase, uint32(v2752))) = uint16(v2825)
	v2828 = F_bms_is_member(m, base.I32_extend16_s(v2825), v2140)
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		goto L21
	} else {
		goto L407
	}
L398:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+24))
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v2756)+12))
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2757+(v2753^int32(-1))<<(uint(int32(2))%32))))
	v2773 = int32(0)
	goto L399
L399:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v2136+v2773<<(uint(int32(2))%32))))
	v2810 = F_equal(m, v2809, v2763)
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L21
	} else {
		goto L401
	}
L400:
	;
	v2819 = v2144 + (v2773 ^ int32(-1))
	if v2819&int32(65535) == int32(0) {
		v2939 = v2671
		goto L388
	} else {
		goto L406
	}
L401:
	;
	if v2810 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v2815 = v2773 + int32(1)
	if v2815 != v2130 {
		v2773 = v2815
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
	v2939 = v2671
	goto L388
L406:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2752))) = uint16(v2819)
	goto L394
L407:
	;
	if v2828 == int32(0) {
		v2939 = v2671
		goto L388
	} else {
		goto L408
	}
L408:
	;
	goto L394
L409:
	;
	goto L393
L410:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2701)))
	*(*int32)(unsafe.Add(mBase, uint32(v2655+v2671<<(uint(int32(2))%32)))) = v2922
	goto L412
L411:
	;
	goto L412
L412:
	;
	v2939 = v2671 + int32(1)
	goto L388
L413:
	;
	goto L387
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2154+v2152<<(uint(int32(2))%32)))) = v2611
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v2611)+8))
	v3065 = v2130
	v3066 = v2131
	v3067 = v2132
	v3068 = v2133
	v3069 = v2134
	v3070 = v2135
	v3071 = v2136
	v3075 = v2140
	v3078 = v2143
	v3079 = v2144
	v3080 = v2145
	v3083 = v2148
	v3085 = v2150
	v3087 = v2152 + int32(1)
	v3089 = v2154
	v3090 = v2155
	v3091 = v2156
	v3092 = v2157
	v3093 = v3062 + v2158
	v3095 = v2160
	v3099 = v2164
	v3100 = v2165
	goto L312
L415:
	;
	goto L311
L416:
	;
	v3159 = v3118
	v3163 = v3122
	v3168 = v3127
	v3171 = v3130
	v3173 = v3132
	v3177 = v3136
	v3187 = v3146
	v3188 = v3147
	goto L306
L417:
	;
	F_bms_free(m, v3163)
	mBase = m.M
	v3196 = m.ExcPending
	if v3196 != 0 {
		goto L21
	} else {
		goto L418
	}
L418:
	;
	F_pfree(m, v3168)
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L21
	} else {
		goto L419
	}
L419:
	;
	v4809 = v3159
	v4821 = v3171
	v4823 = v3173
	v4837 = v3187
	v4838 = v3188
	goto L278
L420:
	;
	v3201 = int32(0)
	if v3122 == v3201 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if int32(0) < v3134 {
		goto L434
	} else {
		goto L435
	}
L422:
	;
	v3236 = int32(0)
	goto L421
L423:
	;
	goto L424
L424:
	;
	v3208 = int32(1)
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3122)+4))
	if v3209 <= v3208 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v3212 = v3208
	goto L427
L426:
	;
	v3212 = v3209
	goto L427
L427:
	;
	v3216 = int32(0)
	v3218 = v3201
	goto L428
L428:
	;
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(v3122+int32(8)+v3216<<(uint(int32(2))%32))))
	if v3224 != 0 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	v3236 = v3227
	goto L421
L430:
	;
	v3227 = v3218 + base.I32_popcnt(v3224)
	goto L432
L431:
	;
	v3227 = v3218
	goto L432
L432:
	;
	v3229 = v3216 + int32(1)
	if v3229 != v3212 {
		v3216 = v3229
		v3218 = v3227
		goto L428
	} else {
		goto L433
	}
L433:
	;
	goto L429
L434:
	;
	v3251 = v3122
	v3257 = int32(0)
	v3258 = v3236
	goto L437
L435:
	;
	v4764 = v3122
	v4789 = v3147
	goto L436
L436:
	;
	F_pfree(m, v3199)
	mBase = m.M
	v4795 = m.ExcPending
	if v4795 != 0 {
		goto L21
	} else {
		goto L595
	}
L437:
	;
	v3281 = int32(0)
	v3284 = v3281
	v3292 = v3281
	goto L439
L438:
	;
	if v3257 != 0 {
		goto L481
	} else {
		goto L482
	}
L439:
	;
	v3327 = v3136 + v3292<<(uint(int32(2))%32)
	v3328 = *(*int32)(unsafe.Add(mBase, uint32(v3327)))
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+8))
	if v3329 != 0 {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	if v3493 != 0 {
		goto L464
	} else {
		goto L465
	}
L441:
	;
	v3331 = v3284
	v3338 = v3328
	v3340 = int32(0)
	goto L444
L442:
	;
	v3493 = v3284
	goto L443
L443:
	;
	v3534 = v3292 + int32(1)
	if v3534 != v3134 {
		v3284 = v3493
		v3292 = v3534
		goto L439
	} else {
		goto L463
	}
L444:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v3338+v3340<<(uint(int32(2))%32))+12))
	v3375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3374)+8)))
	if v3258 < v3375 {
		v3448 = v3331
		v3455 = v3338
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v3493 = v3448
	goto L443
L446:
	;
	v3489 = v3340 + int32(1)
	v3490 = *(*int32)(unsafe.Add(mBase, uint32(v3455)+8))
	if base.Ui32(v3489) < base.Ui32(v3490) {
		v3331 = v3448
		v3338 = v3455
		v3340 = v3489
		goto L444
	} else {
		goto L462
	}
L447:
	;
	if v3331 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	if v3375 <= int32(0) {
		goto L453
	} else {
		goto L454
	}
L449:
	;
	v3379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3331)+8)))
	if v3375 < v3379 {
		v3448 = v3331
		v3455 = v3338
		goto L446
	} else {
		goto L450
	}
L450:
	;
	if v3375 != v3379 {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	v3382 = *(*float64)(unsafe.Add(mBase, uint32(v3331)))
	v3383 = *(*float64)(unsafe.Add(mBase, uint32(v3374)))
	if base.F64_gt(v3382, v3383) != 0 {
		v3448 = v3331
		v3455 = v3338
		goto L446
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	v3448 = v3374
	v3455 = v3338
	goto L446
L454:
	;
	goto L455
L455:
	;
	v3399 = int32(0)
	goto L457
L456:
	;
	v3446 = *(*int32)(unsafe.Add(mBase, uint32(v3327)))
	v3448 = v3444
	v3455 = v3446
	goto L446
L457:
	;
	v3435 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3374+int32(10)+v3399<<(uint(int32(1))%32)))))
	v3436 = F_bms_is_member(m, v3435, v3251)
	mBase = m.M
	v3437 = m.ExcPending
	if v3437 != 0 {
		goto L21
	} else {
		goto L459
	}
L458:
	;
	v3444 = v3374
	goto L456
L459:
	;
	if v3436 == int32(0) {
		v3444 = v3331
		goto L456
	} else {
		goto L460
	}
L460:
	;
	v3441 = v3399 + int32(1)
	v3442 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3374)+8)))
	if v3441 < v3442 {
		v3399 = v3441
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
	*(*int32)(unsafe.Add(mBase, uint32(v3199+v3257<<(uint(int32(2))%32)))) = v3493
	v3540 = int32(1)
	v3542 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3493)+8)))
	v3546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3542<<(uint(v3540)%32)+v3493)+8)))
	v3547 = F_bms_del_member(m, v3251, v3546)
	mBase = m.M
	v3548 = m.ExcPending
	if v3548 != 0 {
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
	v3549 = int32(0)
	if v3547 == v3549 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v3251 = v3547
	v3257 = v3257 + v3540
	v3258 = v3584
	goto L437
L469:
	;
	v3584 = int32(0)
	goto L468
L470:
	;
	goto L471
L471:
	;
	v3556 = int32(1)
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v3547)+4))
	if v3557 <= v3556 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v3560 = v3556
	goto L474
L473:
	;
	v3560 = v3557
	goto L474
L474:
	;
	v3564 = int32(0)
	v3566 = v3549
	goto L475
L475:
	;
	v3572 = *(*int32)(unsafe.Add(mBase, uint32(v3547+int32(8)+v3564<<(uint(int32(2))%32))))
	if v3572 != 0 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v3584 = v3575
	goto L468
L477:
	;
	v3575 = v3566 + base.I32_popcnt(v3572)
	goto L479
L478:
	;
	v3575 = v3566
	goto L479
L479:
	;
	v3577 = v3564 + int32(1)
	if v3577 != v3560 {
		v3564 = v3577
		v3566 = v3575
		goto L475
	} else {
		goto L480
	}
L480:
	;
	goto L476
L481:
	;
	v3585 = int32(0)
	if v3585 < v3257 {
		goto L484
	} else {
		goto L485
	}
L482:
	;
	v4697 = v3147
	goto L483
L483:
	;
	v4711 = int32(0)
	goto L591
L484:
	;
	v3589 = v3585
	v3607 = int32(0)
	goto L487
L485:
	;
	v3734 = v3585
	goto L486
L486:
	;
	v3775 = int32(0)
	if v3734 == v3775 {
		goto L498
	} else {
		goto L499
	}
L487:
	;
	v3630 = int32(0)
	v3633 = v3199 + v3607<<(uint(int32(2))%32)
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3633)))
	v3635 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3634)+8)))
	if v3630 < v3635 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	v3734 = v3690
	goto L486
L489:
	;
	v3638 = v3589
	v3646 = v3630
	v3650 = v3634
	goto L492
L490:
	;
	v3690 = v3589
	goto L491
L491:
	;
	v3732 = v3607 + int32(1)
	if v3732 != v3257 {
		v3589 = v3690
		v3607 = v3732
		goto L487
	} else {
		goto L496
	}
L492:
	;
	v3682 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3650+v3646<<(uint(int32(1))%32))+10)))
	v3683 = F_bms_add_member(m, v3638, v3682)
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L21
	} else {
		goto L494
	}
L493:
	;
	v3690 = v3683
	goto L491
L494:
	;
	v3686 = v3646 + int32(1)
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(v3633)))
	v3688 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3687)+8)))
	if v3686 < v3688 {
		v3638 = v3683
		v3646 = v3686
		v3650 = v3687
		goto L492
	} else {
		goto L495
	}
L495:
	;
	goto L493
L496:
	;
	goto L488
L497:
	;
	v3813 = F_palloc(m, v3810<<(uint(int32(3))%32))
	mBase = m.M
	v3814 = m.ExcPending
	if v3814 != 0 {
		goto L21
	} else {
		goto L510
	}
L498:
	;
	v3810 = int32(0)
	goto L497
L499:
	;
	goto L500
L500:
	;
	v3782 = int32(1)
	v3783 = *(*int32)(unsafe.Add(mBase, uint32(v3734)+4))
	if v3783 <= v3782 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v3786 = v3782
	goto L503
L502:
	;
	v3786 = v3783
	goto L503
L503:
	;
	v3790 = int32(0)
	v3792 = v3775
	goto L504
L504:
	;
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(v3734+int32(8)+v3790<<(uint(int32(2))%32))))
	if v3798 != 0 {
		goto L506
	} else {
		goto L507
	}
L505:
	;
	v3810 = v3801
	goto L497
L506:
	;
	v3801 = v3792 + base.I32_popcnt(v3798)
	goto L508
L507:
	;
	v3801 = v3792
	goto L508
L508:
	;
	v3803 = v3790 + int32(1)
	if v3803 != v3786 {
		v3790 = v3803
		v3792 = v3801
		goto L504
	} else {
		goto L509
	}
L509:
	;
	goto L505
L510:
	;
	if v3734 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L511:
	;
	if int32(0) <= v3871 {
		goto L522
	} else {
		goto L523
	}
L512:
	;
	v3871 = base.I32_ctz(v3857) | v3858<<(uint(int32(5))%32)
	goto L511
L513:
	;
	v3871 = int32(-2)
	goto L511
L514:
	;
	v3824 = base.I32_div_s(int32(0), int32(32))
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v3734)+4))
	if v3825 <= v3824 {
		goto L513
	} else {
		goto L515
	}
L515:
	;
	v3828 = v3734 + int32(8)
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3828+v3824<<(uint(int32(2))%32))))
	v3835 = v3832 & int32(-1)
	if v3835 != 0 {
		v3857 = v3835
		v3858 = v3824
		goto L512
	} else {
		goto L516
	}
L516:
	;
	v3837 = v3824 + int32(1)
	if v3837 == v3825 {
		goto L513
	} else {
		goto L517
	}
L517:
	;
	v3840 = v3837
	goto L518
L518:
	;
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v3828+v3840<<(uint(int32(2))%32))))
	if v3847 != 0 {
		v3857 = v3847
		v3858 = v3840
		goto L512
	} else {
		goto L520
	}
L519:
	;
	goto L513
L520:
	;
	v3849 = v3840 + int32(1)
	if v3849 != v3825 {
		v3840 = v3849
		goto L518
	} else {
		goto L521
	}
L521:
	;
	goto L519
L522:
	;
	v3890 = int32(0)
	v3893 = v3871
	goto L525
L523:
	;
	goto L524
L524:
	;
	v4140 = v3257 - int32(1)
	if int32(0) <= v4140 {
		goto L553
	} else {
		goto L554
	}
L525:
	;
	if v3125 == int32(0) {
		goto L528
	} else {
		goto L529
	}
L526:
	;
	goto L524
L527:
	;
	v4035 = F_clauselist_selectivity_ext(m, v3116, v4010, v3113, v3114, v3115, int32(0))
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		goto L21
	} else {
		goto L540
	}
L528:
	;
	v4010 = int32(0)
	goto L527
L529:
	;
	goto L530
L530:
	;
	v3919 = int32(0)
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v3137)))
	if v3922 <= v3919 {
		v4010 = v3919
		goto L527
	} else {
		goto L531
	}
L531:
	;
	v3933 = int32(-1)
	v3935 = v3922
	v3937 = v3919
	v3945 = v3919
	goto L532
L532:
	;
	v3966 = int32(1)
	v3967 = v3933 + v3966
	v3971 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3127+v3967<<(uint(v3966)%32)))))
	if v3971 == v3893 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v4010 = v3986
	goto L527
L534:
	;
	v3973 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+12))
	v3977 = *(*int32)(unsafe.Add(mBase, uint32(v3973+v3937<<(uint(int32(2))%32))))
	v3978 = F_lappend(m, v3945, v3977)
	mBase = m.M
	v3979 = m.ExcPending
	if v3979 != 0 {
		goto L21
	} else {
		goto L537
	}
L535:
	;
	v3985 = v3935
	v3986 = v3945
	goto L536
L536:
	;
	v3988 = v3937 + int32(1)
	if v3988 < v3985 {
		v3933 = v3967
		v3935 = v3985
		v3937 = v3988
		v3945 = v3986
		goto L532
	} else {
		goto L539
	}
L537:
	;
	v3980 = *(*int32)(unsafe.Add(mBase, uint32(v3117)))
	v3981 = F_bms_add_member(m, v3980, v3967)
	mBase = m.M
	v3982 = m.ExcPending
	if v3982 != 0 {
		goto L21
	} else {
		goto L538
	}
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3117))) = v3981
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+4))
	v3985 = v3984
	v3986 = v3978
	goto L536
L539:
	;
	goto L533
L540:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v3813+v3890<<(uint(int32(3))%32)))) = v4035
	if v3734 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L541:
	;
	if int32(0) <= v4095 {
		v3890 = v3890 + int32(1)
		v3893 = v4095
		goto L525
	} else {
		goto L552
	}
L542:
	;
	v4095 = base.I32_ctz(v4081) | v4082<<(uint(int32(5))%32)
	goto L541
L543:
	;
	v4095 = int32(-2)
	goto L541
L544:
	;
	v4046 = v3893 + int32(1)
	v4048 = base.I32_div_s(v4046, int32(32))
	v4049 = *(*int32)(unsafe.Add(mBase, uint32(v3734)+4))
	if v4049 <= v4048 {
		goto L543
	} else {
		goto L545
	}
L545:
	;
	v4052 = v3734 + int32(8)
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v4052+v4048<<(uint(int32(2))%32))))
	v4059 = v4056 & (int32(-1) << (uint(v4046) % 32))
	if v4059 != 0 {
		v4081 = v4059
		v4082 = v4048
		goto L542
	} else {
		goto L546
	}
L546:
	;
	v4061 = v4048 + int32(1)
	if v4061 == v4049 {
		goto L543
	} else {
		goto L547
	}
L547:
	;
	v4064 = v4061
	goto L548
L548:
	;
	v4071 = *(*int32)(unsafe.Add(mBase, uint32(v4052+v4064<<(uint(int32(2))%32))))
	if v4071 != 0 {
		v4081 = v4071
		v4082 = v4064
		goto L542
	} else {
		goto L550
	}
L549:
	;
	goto L543
L550:
	;
	v4073 = v4064 + int32(1)
	if v4073 != v4049 {
		v4064 = v4073
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	goto L526
L553:
	;
	v4163 = v4140
	goto L556
L554:
	;
	goto L555
L555:
	;
	if int32(0) < v3810 {
		goto L571
	} else {
		goto L572
	}
L556:
	;
	v4184 = float64(1)
	v4186 = int32(2)
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v3199+v4163<<(uint(v4186)%32))))
	v4190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4189)+8)))
	if v4186 <= v4190 {
		goto L558
	} else {
		goto L559
	}
L557:
	;
	goto L555
L558:
	;
	v4204 = int32(0)
	v4232 = v4184
	goto L561
L559:
	;
	v4266 = int32(0)
	v4290 = v4184
	goto L560
L560:
	;
	v4298 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4189+v4266<<(uint(int32(1))%32))+10)))
	v4299 = F_bms_member_index(m, v3734, v4298)
	mBase = m.M
	v4300 = m.ExcPending
	if v4300 != 0 {
		goto L21
	} else {
		goto L565
	}
L561:
	;
	v4240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4189+int32(10)+v4204<<(uint(int32(1))%32)))))
	v4241 = F_bms_member_index(m, v3734, v4240)
	mBase = m.M
	v4242 = m.ExcPending
	if v4242 != 0 {
		goto L21
	} else {
		goto L563
	}
L562:
	;
	v4266 = v4249
	v4290 = v4247
	goto L560
L563:
	;
	v4246 = *(*float64)(unsafe.Add(mBase, uint32(v3813+v4241<<(uint(int32(3))%32))))
	v4247 = base.F64_mul(v4232, v4246)
	v4248 = int32(1)
	v4249 = v4204 + v4248
	v4250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4189)+8)))
	if v4249 < v4250-v4248 {
		v4204 = v4249
		v4232 = v4247
		goto L561
	} else {
		goto L564
	}
L564:
	;
	goto L562
L565:
	;
	v4301 = *(*float64)(unsafe.Add(mBase, uint32(v4189)))
	v4304 = v3813 + v4299<<(uint(int32(3))%32)
	v4305 = *(*float64)(unsafe.Add(mBase, uint32(v4304)))
	if base.F64_le(v4290, v4305) == int32(0) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v4311 = base.F64_div(base.F64_mul(v4305, v4301), v4290)
	goto L568
L567:
	;
	v4311 = v4301
	goto L568
L568:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v4304))) = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v4301), v4305), v4311)
	if int32(0) < v4163 {
		v4163 = v4163 - int32(1)
		goto L556
	} else {
		goto L569
	}
L569:
	;
	goto L557
L570:
	;
	F_pfree(m, v3813)
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		goto L21
	} else {
		goto L589
	}
L571:
	;
	v4365 = v3810 & int32(3)
	if base.Ui32(v3810) < base.Ui32(int32(4)) {
		goto L575
	} else {
		goto L576
	}
L572:
	;
	goto L573
L573:
	;
	v4656 = float64(1)
	goto L570
L574:
	;
	if v4365 != 0 {
		goto L581
	} else {
		goto L582
	}
L575:
	;
	v4441 = int32(0)
	v4467 = float64(1)
	goto L574
L576:
	;
	goto L577
L577:
	;
	v4374 = int32(0)
	v4384 = v4374
	v4394 = v4374
	v4410 = float64(1)
	goto L578
L578:
	;
	v4419 = v3813 + v4384<<(uint(int32(3))%32)
	v4420 = *(*float64)(unsafe.Add(mBase, uint32(v4419)))
	v4422 = *(*float64)(unsafe.Add(mBase, uint32(v4419)+8))
	v4424 = *(*float64)(unsafe.Add(mBase, uint32(v4419)+16))
	v4426 = *(*float64)(unsafe.Add(mBase, uint32(v4419)+24))
	v4427 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(v4410, v4420), v4422), v4424), v4426)
	v4428 = int32(4)
	v4429 = v4384 + v4428
	v4431 = v4394 + v4428
	if v4431 != v3810&int32(2147483644) {
		v4384 = v4429
		v4394 = v4431
		v4410 = v4427
		goto L578
	} else {
		goto L580
	}
L579:
	;
	v4441 = v4429
	v4467 = v4427
	goto L574
L580:
	;
	goto L579
L581:
	;
	v4482 = v4441
	v4484 = int32(0)
	v4508 = v4467
	goto L584
L582:
	;
	v4559 = v4467
	goto L583
L583:
	;
	v4566 = float64(0)
	if base.F64_lt(v4559, v4566) != 0 {
		v4656 = v4566
		goto L570
	} else {
		goto L587
	}
L584:
	;
	v4518 = *(*float64)(unsafe.Add(mBase, uint32(v3813+v4482<<(uint(int32(3))%32))))
	v4519 = base.F64_mul(v4508, v4518)
	v4520 = int32(1)
	v4523 = v4484 + v4520
	if v4523 != v4365 {
		v4482 = v4482 + v4520
		v4484 = v4523
		v4508 = v4519
		goto L584
	} else {
		goto L586
	}
L585:
	;
	v4559 = v4519
	goto L583
L586:
	;
	goto L585
L587:
	;
	if base.F64_gt(v4559, float64(1)) == int32(0) {
		v4656 = v4559
		goto L570
	} else {
		goto L588
	}
L588:
	;
	goto L573
L589:
	;
	F_bms_free(m, v3734)
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L21
	} else {
		goto L590
	}
L590:
	;
	v4697 = v4656
	goto L483
L591:
	;
	v4747 = *(*int32)(unsafe.Add(mBase, uint32(v3136+v4711<<(uint(int32(2))%32))))
	F_pfree(m, v4747)
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		goto L21
	} else {
		goto L593
	}
L592:
	;
	v4764 = v3251
	v4789 = v4697
	goto L436
L593:
	;
	v4751 = v4711 + int32(1)
	if v4751 != v3134 {
		v4711 = v4751
		goto L591
	} else {
		goto L594
	}
L594:
	;
	goto L592
L595:
	;
	F_pfree(m, v3136)
	mBase = m.M
	v4797 = m.ExcPending
	if v4797 != 0 {
		goto L21
	} else {
		goto L596
	}
L596:
	;
	F_bms_free(m, v4764)
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L21
	} else {
		goto L597
	}
L597:
	;
	F_pfree(m, v3127)
	mBase = m.M
	v4801 = m.ExcPending
	if v4801 != 0 {
		goto L21
	} else {
		goto L598
	}
L598:
	;
	v4809 = v3118
	v4821 = v3130
	v4823 = v3132
	v4837 = v3146
	v4838 = v4789
	goto L278
L599:
	;
	v4864 = v4821
	v4866 = v4823
	v4880 = v4837
	v4881 = v4838
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
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
	var v294 int32
	_ = v294
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 != int32(318) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v294
L2:
	;
	if v13 != int32(21) {
		v294 = v6
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
		v294 = v6
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 != 0 {
		v294 = v6
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 == int32(0) {
		v294 = v19
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= v23 {
		v294 = v19
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
	v294 = v40
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	if v40 == int32(0) {
		v294 = v40
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
	if v107 == int32(0) {
		v294 = v6
		goto L1
	} else {
		goto L32
	}
L17:
	;
	v107 = int32(0)
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
	v74 = int32(-1)
	goto L24
L23:
	;
	v107 = v99
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
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v91
	v99 = int32(1)
	goto L23
L26:
	;
	if int32(0) <= v74 {
		v99 = v54
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v91 = v74
	goto L28
L28:
	;
	v93 = v71 + int32(1)
	if v93 != v66 {
		v71 = v93
		v74 = v91
		goto L24
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v81)) {
		v99 = v54
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v91 = base.I32_ctz(v81) | v71<<(uint(int32(5))%32)
	goto L28
L31:
	;
	goto L25
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v110 != l2 {
		v294 = v6
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)) = uint8(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v117 = F_statext_is_compatible_clause_internal(m, v114, l2, l3, l4, v11+int32(11))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	if v117 == int32(0) {
		v294 = v6
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v121 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v127 == v124 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	goto L38
L38:
	;
	v294 = int32(1)
	goto L1
L39:
	;
	if int32(0) <= v184 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v184 = base.I32_ctz(v170) | v171<<(uint(int32(5))%32)
	goto L39
L41:
	;
	v184 = int32(-2)
	goto L39
L42:
	;
	v137 = base.I32_div_s(int32(0), int32(32))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v138 <= v137 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v141 = v127 + int32(8)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v137<<(uint(int32(2))%32))))
	v148 = v145 & int32(-1)
	if v148 != 0 {
		v170 = v148
		v171 = v137
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v150 = v137 + int32(1)
	if v150 == v138 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v153 = v150
	goto L46
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v141+v153<<(uint(int32(2))%32))))
	if v160 != 0 {
		v170 = v160
		v171 = v153
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v162 = v153 + int32(1)
	if v162 != v138 {
		v153 = v162
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v188 = v184
	v192 = v124
	goto L53
L51:
	;
	v264 = v124
	goto L52
L52:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v267 != 0 {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	v197 = F_bms_add_member(m, v192, v188+int32(7))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L55
	}
L54:
	;
	v264 = v197
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v197
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v200 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if int32(0) <= v256 {
		v188 = v256
		v192 = v197
		goto L53
	} else {
		goto L67
	}
L57:
	;
	v256 = base.I32_ctz(v242) | v243<<(uint(int32(5))%32)
	goto L56
L58:
	;
	v256 = int32(-2)
	goto L56
L59:
	;
	v207 = v188 + int32(1)
	v209 = base.I32_div_s(v207, int32(32))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v210 <= v209 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v213 = v200 + int32(8)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v209<<(uint(int32(2))%32))))
	v220 = v217 & (int32(-1) << (uint(v207) % 32))
	if v220 != 0 {
		v242 = v220
		v243 = v209
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v222 = v209 + int32(1)
	if v222 == v210 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v225 = v222
	goto L63
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v213+v225<<(uint(int32(2))%32))))
	if v232 != 0 {
		v242 = v232
		v243 = v225
		goto L57
	} else {
		goto L65
	}
L64:
	;
	goto L58
L65:
	;
	v234 = v225 + int32(1)
	if v234 != v210 {
		v225 = v234
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L54
L68:
	;
	F_pull_varattnos(m, v267, l2, v11+int32(4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	v273 = v264
	goto L70
L70:
	;
	v274 = F_all_rows_selectable(m, l0, l2, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L11
	} else {
		goto L72
	}
L71:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v273 = v272
	goto L70
L72:
	;
	if v274 == int32(0) {
		v294 = v6
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L38
}
