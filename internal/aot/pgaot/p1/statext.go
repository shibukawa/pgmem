package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
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
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v298 float64
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v468 int32
	_ = v468
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v509 int32
	_ = v509
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v568 int32
	_ = v568
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v764 int32
	_ = v764
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v890 float64
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v918 int32
	_ = v918
	var v934 int32
	_ = v934
	var v957 float64
	_ = v957
	var v965 int32
	_ = v965
	var v966 float64
	_ = v966
	var v967 float64
	_ = v967
	var v971 int32
	_ = v971
	var v974 float64
	_ = v974
	var v975 float64
	_ = v975
	var v978 float64
	_ = v978
	var v980 float64
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v1025 float64
	_ = v1025
	var v1026 float64
	_ = v1026
	var v1027 float64
	_ = v1027
	var v1028 float64
	_ = v1028
	var v1030 float64
	_ = v1030
	var v1038 float64
	_ = v1038
	var v1040 float64
	_ = v1040
	var v1042 float64
	_ = v1042
	var v1043 float64
	_ = v1043
	var v1051 float64
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 float64
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1078 int32
	_ = v1078
	var v1102 float64
	_ = v1102
	var v1104 float64
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 float64
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 float64
	_ = v1118
	var v1120 float64
	_ = v1120
	var v1128 float64
	_ = v1128
	var v1129 float64
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1172 int32
	_ = v1172
	var v1186 int32
	_ = v1186
	var v1212 float64
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1218 float64
	_ = v1218
	var v1219 float64
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 float64
	_ = v1226
	var v1227 float64
	_ = v1227
	var v1228 float64
	_ = v1228
	var v1231 float64
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 float64
	_ = v1237
	var v1238 float64
	_ = v1238
	var v1241 float64
	_ = v1241
	var v1242 float64
	_ = v1242
	var v1245 float64
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1299 float64
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 float64
	_ = v1310
	var v1311 float64
	_ = v1311
	var v1312 float64
	_ = v1312
	var v1314 float64
	_ = v1314
	var v1322 float64
	_ = v1322
	var v1324 float64
	_ = v1324
	var v1326 float64
	_ = v1326
	var v1327 float64
	_ = v1327
	var v1335 float64
	_ = v1335
	var v1336 float64
	_ = v1336
	var v1337 float64
	_ = v1337
	var v1338 float64
	_ = v1338
	var v1339 float64
	_ = v1339
	var v1340 float64
	_ = v1340
	var v1342 float64
	_ = v1342
	var v1350 float64
	_ = v1350
	var v1352 float64
	_ = v1352
	var v1354 float64
	_ = v1354
	var v1355 float64
	_ = v1355
	var v1363 float64
	_ = v1363
	var v1365 float64
	_ = v1365
	var v1373 float64
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1411 float64
	_ = v1411
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1437 int32
	_ = v1437
	var v1455 float64
	_ = v1455
	var v1463 int32
	_ = v1463
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 float64
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1563 int32
	_ = v1563
	var v1635 int32
	_ = v1635
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1670 int32
	_ = v1670
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1836 int32
	_ = v1836
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1884 int32
	_ = v1884
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1967 int32
	_ = v1967
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1989 int32
	_ = v1989
	var v1996 int32
	_ = v1996
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2076 int32
	_ = v2076
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
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
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
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
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2173 float64
	_ = v2173
	var v2174 float64
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2249 int32
	_ = v2249
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2308 int32
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2357 int32
	_ = v2357
	var v2369 int32
	_ = v2369
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2459 int32
	_ = v2459
	var v2465 int32
	_ = v2465
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2518 int32
	_ = v2518
	var v2545 int32
	_ = v2545
	var v2561 int32
	_ = v2561
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2641 int32
	_ = v2641
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2721 int32
	_ = v2721
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2767 int32
	_ = v2767
	var v2777 int32
	_ = v2777
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2818 int32
	_ = v2818
	var v2822 int32
	_ = v2822
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2916 int32
	_ = v2916
	var v2931 int32
	_ = v2931
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2975 int32
	_ = v2975
	var v3016 int32
	_ = v3016
	var v3053 int32
	_ = v3053
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3065 int32
	_ = v3065
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3082 int32
	_ = v3082
	var v3089 float64
	_ = v3089
	var v3090 float64
	_ = v3090
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3121 int32
	_ = v3121
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3135 float64
	_ = v3135
	var v3136 float64
	_ = v3136
	var v3141 int32
	_ = v3141
	var v3149 int32
	_ = v3149
	var v3152 int32
	_ = v3152
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3160 int32
	_ = v3160
	var v3166 int32
	_ = v3166
	var v3176 float64
	_ = v3176
	var v3177 float64
	_ = v3177
	var v3183 int32
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3201 int32
	_ = v3201
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3213 int32
	_ = v3213
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3225 int32
	_ = v3225
	var v3234 int32
	_ = v3234
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3269 int32
	_ = v3269
	var v3272 int32
	_ = v3272
	var v3280 int32
	_ = v3280
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3318 int32
	_ = v3318
	var v3325 int32
	_ = v3325
	var v3330 int32
	_ = v3330
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3365 int32
	_ = v3365
	var v3368 float64
	_ = v3368
	var v3369 float64
	_ = v3369
	var v3385 int32
	_ = v3385
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3429 int32
	_ = v3429
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3440 int32
	_ = v3440
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3477 int32
	_ = v3477
	var v3517 int32
	_ = v3517
	var v3523 int32
	_ = v3523
	var v3527 int32
	_ = v3527
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3545 int32
	_ = v3545
	var v3549 int32
	_ = v3549
	var v3551 int32
	_ = v3551
	var v3557 int32
	_ = v3557
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3574 int32
	_ = v3574
	var v3579 int32
	_ = v3579
	var v3614 int32
	_ = v3614
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3630 int32
	_ = v3630
	var v3636 int32
	_ = v3636
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3673 int32
	_ = v3673
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3756 int32
	_ = v3756
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3771 int32
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3779 int32
	_ = v3779
	var v3782 int32
	_ = v3782
	var v3784 int32
	_ = v3784
	var v3791 int32
	_ = v3791
	var v3794 int32
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3813 int32
	_ = v3813
	var v3816 int32
	_ = v3816
	var v3818 int32
	_ = v3818
	var v3821 int32
	_ = v3821
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3852 int32
	_ = v3852
	var v3861 int32
	_ = v3861
	var v3869 int32
	_ = v3869
	var v3899 int32
	_ = v3899
	var v3902 int32
	_ = v3902
	var v3906 int32
	_ = v3906
	var v3913 int32
	_ = v3913
	var v3916 int32
	_ = v3916
	var v3919 int32
	_ = v3919
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3950 int32
	_ = v3950
	var v3952 int32
	_ = v3952
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3967 int32
	_ = v3967
	var v3980 int32
	_ = v3980
	var v4013 float64
	_ = v4013
	var v4014 int32
	_ = v4014
	var v4024 int32
	_ = v4024
	var v4026 int32
	_ = v4026
	var v4027 int32
	_ = v4027
	var v4030 int32
	_ = v4030
	var v4034 int32
	_ = v4034
	var v4037 int32
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4042 int32
	_ = v4042
	var v4049 int32
	_ = v4049
	var v4051 int32
	_ = v4051
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4073 int32
	_ = v4073
	var v4117 int32
	_ = v4117
	var v4125 int32
	_ = v4125
	var v4160 float64
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4179 int32
	_ = v4179
	var v4206 float64
	_ = v4206
	var v4214 int32
	_ = v4214
	var v4215 int32
	_ = v4215
	var v4216 int32
	_ = v4216
	var v4220 float64
	_ = v4220
	var v4221 float64
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4236 int32
	_ = v4236
	var v4263 float64
	_ = v4263
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4274 float64
	_ = v4274
	var v4277 int32
	_ = v4277
	var v4278 float64
	_ = v4278
	var v4284 float64
	_ = v4284
	var v4337 int32
	_ = v4337
	var v4338 float64
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4351 int32
	_ = v4351
	var v4354 int32
	_ = v4354
	var v4379 float64
	_ = v4379
	var v4388 int32
	_ = v4388
	var v4389 float64
	_ = v4389
	var v4391 float64
	_ = v4391
	var v4393 float64
	_ = v4393
	var v4395 float64
	_ = v4395
	var v4396 float64
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4398 int32
	_ = v4398
	var v4400 int32
	_ = v4400
	var v4412 int32
	_ = v4412
	var v4437 float64
	_ = v4437
	var v4445 int32
	_ = v4445
	var v4452 int32
	_ = v4452
	var v4477 float64
	_ = v4477
	var v4487 float64
	_ = v4487
	var v4488 float64
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4492 int32
	_ = v4492
	var v4527 float64
	_ = v4527
	var v4534 float64
	_ = v4534
	var v4622 float64
	_ = v4622
	var v4624 int32
	_ = v4624
	var v4626 int32
	_ = v4626
	var v4662 float64
	_ = v4662
	var v4676 int32
	_ = v4676
	var v4711 int32
	_ = v4711
	var v4713 int32
	_ = v4713
	var v4715 int32
	_ = v4715
	var v4732 int32
	_ = v4732
	var v4752 float64
	_ = v4752
	var v4758 int32
	_ = v4758
	var v4760 int32
	_ = v4760
	var v4762 int32
	_ = v4762
	var v4764 int32
	_ = v4764
	var v4772 int32
	_ = v4772
	var v4781 int32
	_ = v4781
	var v4783 int32
	_ = v4783
	var v4799 float64
	_ = v4799
	var v4800 float64
	_ = v4800
	var v4806 int32
	_ = v4806
	var v4823 int32
	_ = v4823
	var v4825 int32
	_ = v4825
	var v4841 float64
	_ = v4841
	var v4842 float64
	_ = v4842
	var v4867 int32
	_ = v4867
	var v4885 float64
	_ = v4885
	v41 = m.G0
	v43 = v41 - int32(48)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v45 != 0 {
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v59 = v45 + v46<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v59 = v52 + v53<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	v62 = float64(0)
	goto L7
L6:
	;
	v62 = float64(1)
	goto L7
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l5)+112))
	if v63 == int32(0) {
		v1421 = l0
		v1422 = l1
		v1423 = l2
		v1424 = l3
		v1425 = l4
		v1426 = l5
		v1427 = l6
		v1428 = l7
		v1437 = v43
		v1455 = v62
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v1428 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v66 <= int32(0) {
		v1421 = l0
		v1422 = l1
		v1423 = l2
		v1424 = l3
		v1425 = l4
		v1426 = l5
		v1427 = l6
		v1428 = l7
		v1437 = v43
		v1455 = v62
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v80 = int32(0)
	goto L11
L11:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v70+v80<<(uint(int32(2))%32))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+16)))
	if v115 != int32(109) {
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
	v119 = v80 + int32(1)
	if v119 != v66 {
		v80 = v119
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
	v1421 = l0
	v1422 = l1
	v1423 = l2
	v1424 = l3
	v1425 = l4
	v1426 = l5
	v1427 = l6
	v1428 = l7
	v1437 = v43
	v1455 = v62
	goto L8
L17:
	;
	v264 = l0
	v265 = l1
	v266 = l2
	v267 = l3
	v268 = l4
	v269 = l5
	v270 = l6
	v271 = l7
	v272 = v232
	v280 = v43
	v286 = v246
	v287 = v247
	v289 = v69
	v298 = v62
	goto L36
L18:
	;
	v125 = F_palloc(m, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v133 = l1 + int32(4)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v137 = F_palloc(m, v134<<(uint(int32(2))%32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L21
	} else {
		goto L24
	}
L21:
	;
	return float64(0)
L22:
	;
	v130 = F_palloc(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v232 = int32(4)
	v246 = v125
	v247 = v130
	goto L17
L24:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v142 = F_palloc(m, v139<<(uint(int32(2))%32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v144 <= int32(0) {
		v232 = v133
		v246 = v137
		v247 = v142
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v157 = int32(0)
	goto L27
L27:
	;
	v189 = v157 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190)))
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v193
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v199 = F_bms_is_member(m, v157, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v232 = v133
	v246 = v137
	v247 = v142
	goto L17
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189+v142))) = v218
	v221 = v157 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v221 < v222 {
		v157 = v221
		goto L27
	} else {
		goto L35
	}
L30:
	;
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v189+v137))) = v215
	v218 = v215
	goto L29
L31:
	;
	if v199 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l5)+68))
	v206 = F_statext_is_compatible_clause(m, l0, v192, v201, v43+int32(32), v43+int32(24))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v206 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v189+v137))) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v218 = v213
	goto L29
L35:
	;
	goto L28
L36:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+20)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v269)+112))
	if v265 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v854 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L39:
	;
	v308 = int32(0)
	v313 = F_choose_best_statistics(m, v305, v304&int32(1), v286, v287, v308)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v320 = F_choose_best_statistics(m, v305, v304&int32(1), v286, v287, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L21
	} else {
		goto L44
	}
L42:
	;
	if v313 == int32(0) {
		v1421 = v264
		v1422 = v265
		v1423 = v266
		v1424 = v267
		v1425 = v268
		v1426 = v269
		v1427 = v270
		v1428 = v271
		v1437 = v280
		v1455 = v298
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v847 = v264
	v848 = v265
	v849 = v266
	v850 = v267
	v851 = v268
	v852 = v269
	v853 = v270
	v854 = v271
	v855 = v272
	v859 = v313
	v860 = v308
	v863 = v280
	v868 = v308
	v869 = v286
	v870 = v287
	v872 = v289
	goto L38
L44:
	;
	if v320 == int32(0) {
		v1421 = v264
		v1422 = v265
		v1423 = v266
		v1424 = v267
		v1425 = v268
		v1426 = v269
		v1427 = v270
		v1428 = v271
		v1437 = v280
		v1455 = v298
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v324 = int32(0)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v325 <= v324 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v847 = v264
	v848 = v265
	v849 = v266
	v850 = v267
	v851 = v268
	v852 = v269
	v853 = v270
	v854 = v271
	v855 = v272
	v859 = v320
	v860 = int32(0)
	v863 = v280
	v868 = v324
	v869 = v286
	v870 = v287
	v872 = v289
	goto L38
L47:
	;
	goto L48
L48:
	;
	v330 = int32(0)
	v345 = v330
	v346 = v330
	v349 = int32(-1)
	v353 = v324
	goto L49
L49:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v374 = v349 + int32(1)
	v376 = v374 << (uint(int32(2)) % 32)
	v377 = v286 + v376
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	if v378 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v847 = v264
	v848 = v265
	v849 = v266
	v850 = v267
	v851 = v268
	v852 = v269
	v853 = v270
	v854 = v271
	v855 = v272
	v859 = v320
	v860 = v816
	v863 = v280
	v868 = v824
	v869 = v286
	v870 = v287
	v872 = v289
	goto L38
L51:
	;
	v844 = v346 + int32(1)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v844 < v845 {
		v345 = v816
		v346 = v844
		v349 = v374
		v353 = v824
		goto L49
	} else {
		goto L125
	}
L52:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+v287)))
	if v382 == int32(0) {
		v816 = v345
		v824 = v353
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	v386 = int32(0)
	if v378 == v386 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	if v439 == int32(0) {
		v816 = v345
		v824 = v353
		goto L51
	} else {
		goto L70
	}
L57:
	;
	v439 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v385 == int32(0) {
		v432 = v386
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v439 = v432
	goto L56
L61:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v396 < v395 {
		v432 = v386
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v398 = int32(1)
	if v395 <= v398 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v401 = v398
	goto L65
L64:
	;
	v401 = v395
	goto L65
L65:
	;
	v402 = int32(8)
	v407 = int32(0)
	goto L66
L66:
	;
	v414 = v407 << (uint(int32(2)) % 32)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v378+v402+v414)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v385+v402+v414)))
	v421 = v416 & (v418 ^ int32(-1))
	v423 = base.B2i32(v421 == int32(0))
	if v421 != 0 {
		v432 = v423
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v432 = v423
	goto L60
L68:
	;
	v425 = v407 + int32(1)
	if v425 != v401 {
		v407 = v425
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v442 = v376 + v287
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v443 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v372+v346<<(uint(int32(2))%32))))
	v787 = F_lappend(m, v345, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L21
	} else {
		goto L121
	}
L72:
	;
	if v345 != 0 {
		goto L117
	} else {
		goto L118
	}
L73:
	;
	v650 = int32(0)
	if v619 == v650 {
		goto L101
	} else {
		goto L102
	}
L74:
	;
	v444 = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v444 < v445 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v619 = v609
	goto L73
L77:
	;
	v468 = v444
	goto L80
L78:
	;
	v568 = v443
	goto L79
L79:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	if v598 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v320)+24))
	if v488 == int32(0) {
		v816 = v345
		v824 = v353
		goto L51
	} else {
		goto L82
	}
L81:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v568 = v557
	goto L79
L82:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v491 <= int32(0) {
		v816 = v345
		v824 = v353
		goto L51
	} else {
		goto L83
	}
L83:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v494+v468<<(uint(int32(2))%32))))
	v509 = int32(0)
	goto L84
L84:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v509<<(uint(int32(2))%32))))
	v545 = F_equal(m, v544, v498)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L21
	} else {
		goto L86
	}
L85:
	;
	v554 = v468 + int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v554 < v555 {
		v468 = v554
		goto L80
	} else {
		goto L91
	}
L86:
	;
	if v545 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v550 = v509 + int32(1)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v550 < v551 {
		v509 = v550
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
	v816 = v345
	v824 = v353
	goto L51
L91:
	;
	goto L81
L92:
	;
	if v568 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v568 == int32(0) {
		v619 = v598
		goto L73
	} else {
		goto L99
	}
L95:
	;
	v619 = int32(0)
	goto L73
L96:
	;
	goto L97
L97:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	if v604 == int32(1) {
		goto L72
	} else {
		goto L98
	}
L98:
	;
	v764 = v353
	goto L71
L99:
	;
	v764 = v353
	goto L71
L100:
	;
	if v695 != int32(1) {
		v764 = v353
		goto L71
	} else {
		goto L116
	}
L101:
	;
	v695 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v658 = int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	if v659 <= v658 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v662 = v658
	goto L106
L105:
	;
	v662 = v659
	goto L106
L106:
	;
	v666 = int32(0)
	v668 = v650
	goto L107
L107:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v619+int32(8)+v666<<(uint(int32(2))%32))))
	if v675 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v695 = v687
	goto L100
L109:
	;
	goto L108
L110:
	;
	v676 = int32(2)
	if v668 != 0 {
		v687 = v676
		goto L109
	} else {
		goto L113
	}
L111:
	;
	v682 = v668
	goto L112
L112:
	;
	v684 = v666 + int32(1)
	if v684 != v662 {
		v666 = v684
		v668 = v682
		goto L107
	} else {
		goto L115
	}
L113:
	;
	v677 = int32(1)
	if base.Ui32(v677) < base.Ui32(base.I32_popcnt(v675)) {
		v687 = v676
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v682 = v677
	goto L112
L115:
	;
	v687 = v682
	goto L109
L116:
	;
	goto L72
L117:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v740 = v738
	goto L119
L118:
	;
	v740 = int32(0)
	goto L119
L119:
	;
	v741 = F_bms_add_member(m, v353, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	v764 = v741
	goto L71
L121:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v790 = F_bms_add_member(m, v789, v374)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v790
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	F_bms_free(m, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = int32(0)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	F_list_free(m, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = int32(0)
	v816 = v787
	v824 = v764
	goto L51
L125:
	;
	goto L50
L126:
	;
	v890 = F_clauselist_selectivity_ext(m, v847, v860, v849, v850, v851, int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L21
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v863)+44)) = int32(0)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v859)+4))
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872)+20)))
	v1057 = F_statext_mcv_load(m, v1055, v1056)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L21
	} else {
		goto L150
	}
L129:
	;
	v893 = v863 + int32(32)
	v895 = v863 + int32(24)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v859)+4))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v847)+36))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v852)+68))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v898+v899<<(uint(int32(2))%32))))
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903)+20)))
	v905 = F_statext_mcv_load(m, v897, v904)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L21
	} else {
		goto L131
	}
L130:
	;
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v863)+32))
	v1027 = *(*float64)(unsafe.Add(mBase, uint32(v863)+24))
	v1028 = float64(0)
	v1030 = base.F64_sub(v890, v1026)
	if base.F64_lt(v1030, v1028) != 0 {
		v1038 = v1028
		goto L141
	} else {
		goto L142
	}
L131:
	;
	v907 = int32(0)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v859)+20))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v859)+24))
	v911 = F_mcv_get_match_bitmap(m, v860, v908, v909, v905, v907)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	v913 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v893))) = v913
	*(*int64)(unsafe.Add(mBase, uint32(v895))) = v913
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v905)+8))
	if v918 == int32(0) {
		v1025 = float64(0)
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v934 = v907
	v957 = float64(0)
	goto L134
L134:
	;
	v965 = v905 + int32(48) + v934*int32(24)
	v966 = *(*float64)(unsafe.Add(mBase, uint32(v965)))
	v967 = *(*float64)(unsafe.Add(mBase, uint32(v895)))
	*(*float64)(unsafe.Add(mBase, uint32(v895))) = base.F64_add(v966, v967)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934+v911))))
	if v971 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v1025 = v980
	goto L130
L136:
	;
	v974 = *(*float64)(unsafe.Add(mBase, uint32(v965)+8))
	v975 = *(*float64)(unsafe.Add(mBase, uint32(v893)))
	*(*float64)(unsafe.Add(mBase, uint32(v893))) = base.F64_add(v974, v975)
	v978 = *(*float64)(unsafe.Add(mBase, uint32(v965)))
	v980 = base.F64_add(v957, v978)
	goto L138
L137:
	;
	v980 = v957
	goto L138
L138:
	;
	v982 = v934 + int32(1)
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v905)+8))
	if base.Ui32(v982) < base.Ui32(v983) {
		v934 = v982
		v957 = v980
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	v264 = v847
	v265 = v848
	v266 = v849
	v267 = v850
	v268 = v851
	v269 = v852
	v270 = v853
	v271 = v854
	v272 = v855
	v280 = v863
	v286 = v869
	v287 = v870
	v289 = v872
	v298 = base.F64_mul(v298, v1051)
	goto L36
L141:
	;
	v1040 = base.F64_sub(float64(1), v1027)
	if base.F64_lt(v1040, v1038) != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	if base.F64_gt(v1030, float64(1)) == int32(0) {
		v1038 = v1030
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v1038 = float64(1)
	goto L141
L144:
	;
	goto L140
L145:
	;
	v1042 = v1040
	goto L147
L146:
	;
	v1042 = v1038
	goto L147
L147:
	;
	v1043 = base.F64_add(v1025, v1042)
	if base.F64_lt(v1043, float64(0)) != 0 {
		v1051 = v1028
		goto L144
	} else {
		goto L148
	}
L148:
	;
	if base.F64_gt(v1043, float64(1)) == int32(0) {
		v1051 = v1043
		goto L144
	} else {
		goto L149
	}
L149:
	;
	v1051 = float64(1)
	goto L144
L150:
	;
	if v860 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v264 = v847
	v265 = v848
	v266 = v849
	v267 = v850
	v268 = v851
	v269 = v852
	v270 = v853
	v271 = v854
	v272 = v855
	v280 = v863
	v286 = v869
	v287 = v870
	v289 = v872
	v298 = base.F64_sub(base.F64_add(v298, v1411), base.F64_mul(v298, v1411))
	goto L36
L152:
	;
	v1078 = v1059
	v1102 = v1060
	v1104 = v1060
	goto L157
L153:
	;
	v1059 = int32(0)
	v1060 = float64(0)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	if v1059 < v1062 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v1411 = float64(0)
	goto L151
L156:
	;
	goto L155
L157:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1110+v1078<<(uint(int32(2))%32))))
	v1116 = F_clause_selectivity_ext(m, v847, v1114, v849, v850, v851, int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L21
	} else {
		goto L160
	}
L158:
	;
	v1411 = v1373
	goto L151
L159:
	;
	v1129 = float64(0)
	v1130 = m.G0
	v1132 = v1130 - int32(16)
	m.G0 = v1132
	v1135 = v863 + int32(44)
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	if v1136 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v1118 = base.F64_mul(v1104, v1116)
	v1120 = base.F64_add(v1104, base.F64_sub(v1116, v1118))
	if base.F64_lt(v1120, float64(0)) != 0 {
		v1128 = float64(0)
		goto L159
	} else {
		goto L161
	}
L161:
	;
	if base.F64_gt(v1120, float64(1)) == int32(0) {
		v1128 = v1120
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v1128 = float64(1)
	goto L159
L163:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+8))
	v1140 = F_palloc0(m, v1139)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L21
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v1144 = v863 + int32(32)
	v1146 = v863 + int32(24)
	v1148 = v863 + int32(16)
	v1149 = int32(8)
	v1150 = v863 + v1149
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+8)) = v1114
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+12)) = v1114
	v1157 = F_list_make1_impl(m, int32(1), v1132+v1149)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L21
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1135))) = v1140
	goto L165
L167:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v859)+20))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v859)+24))
	v1162 = F_mcv_get_match_bitmap(m, v1157, v1159, v1160, v1057, int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L21
	} else {
		goto L168
	}
L168:
	;
	v1164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1144))) = v1164
	*(*int64)(unsafe.Add(mBase, uint32(v1146))) = v1164
	*(*int64)(unsafe.Add(mBase, uint32(v1148))) = v1164
	*(*int64)(unsafe.Add(mBase, uint32(v1150))) = v1164
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+8))
	if v1172 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1186 = int32(0)
	v1212 = v1129
	goto L172
L170:
	;
	v1299 = v1129
	goto L171
L171:
	;
	F_pfree(m, v1162)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L21
	} else {
		goto L181
	}
L172:
	;
	v1217 = v1057 + int32(48) + v1186*int32(24)
	v1218 = *(*float64)(unsafe.Add(mBase, uint32(v1217)))
	v1219 = *(*float64)(unsafe.Add(mBase, uint32(v1150)))
	*(*float64)(unsafe.Add(mBase, uint32(v1150))) = base.F64_add(v1218, v1219)
	v1222 = v1186 + v1162
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1222))))
	if v1223 != int32(1) {
		v1245 = v1212
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1299 = v1245
	goto L171
L174:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	v1249 = v1248 + v1186
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249))))
	if v1250 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v1226 = *(*float64)(unsafe.Add(mBase, uint32(v1217)))
	v1227 = *(*float64)(unsafe.Add(mBase, uint32(v1217)+8))
	v1228 = *(*float64)(unsafe.Add(mBase, uint32(v1144)))
	*(*float64)(unsafe.Add(mBase, uint32(v1144))) = base.F64_add(v1227, v1228)
	v1231 = base.F64_add(v1212, v1226)
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1186))))
	if v1234 != int32(1) {
		v1245 = v1231
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v1237 = *(*float64)(unsafe.Add(mBase, uint32(v1217)))
	v1238 = *(*float64)(unsafe.Add(mBase, uint32(v1146)))
	*(*float64)(unsafe.Add(mBase, uint32(v1146))) = base.F64_add(v1237, v1238)
	v1241 = *(*float64)(unsafe.Add(mBase, uint32(v1217)+8))
	v1242 = *(*float64)(unsafe.Add(mBase, uint32(v1148)))
	*(*float64)(unsafe.Add(mBase, uint32(v1148))) = base.F64_add(v1241, v1242)
	v1245 = v1231
	goto L174
L177:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1222))))
	v1254 = v1253
	goto L179
L178:
	;
	v1254 = int32(1)
	goto L179
L179:
	;
	v1255 = int32(1)
	v1256 = v1254 & v1255
	*(*uint8)(unsafe.Add(mBase, uint32(v1249))) = uint8(v1256)
	v1259 = v1186 + v1255
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+8))
	if base.Ui32(v1259) < base.Ui32(v1260) {
		v1186 = v1259
		v1212 = v1245
		goto L172
	} else {
		goto L180
	}
L180:
	;
	goto L173
L181:
	;
	m.G0 = v1132 + int32(16)
	v1308 = F_bms_is_member(m, v1078, v868)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L21
	} else {
		goto L183
	}
L182:
	;
	v1375 = v1078 + int32(1)
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	if v1375 < v1376 {
		v1078 = v1375
		v1102 = v1373
		v1104 = v1128
		goto L157
	} else {
		goto L209
	}
L183:
	;
	if v1308 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1336 = v1116
	goto L186
L185:
	;
	v1310 = *(*float64)(unsafe.Add(mBase, uint32(v863)+32))
	v1311 = *(*float64)(unsafe.Add(mBase, uint32(v863)+8))
	v1312 = float64(0)
	v1314 = base.F64_sub(v1116, v1310)
	if base.F64_lt(v1314, v1312) != 0 {
		v1322 = v1312
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1337 = *(*float64)(unsafe.Add(mBase, uint32(v863)+24))
	v1338 = *(*float64)(unsafe.Add(mBase, uint32(v863)+16))
	v1339 = *(*float64)(unsafe.Add(mBase, uint32(v863)+8))
	v1340 = float64(0)
	v1342 = base.F64_sub(v1118, v1338)
	if base.F64_lt(v1342, v1340) != 0 {
		v1350 = v1340
		goto L198
	} else {
		goto L199
	}
L187:
	;
	v1336 = v1335
	goto L186
L188:
	;
	v1324 = base.F64_sub(float64(1), v1311)
	if base.F64_lt(v1324, v1322) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if base.F64_gt(v1314, float64(1)) == int32(0) {
		v1322 = v1314
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1322 = float64(1)
	goto L188
L191:
	;
	goto L187
L192:
	;
	v1326 = v1324
	goto L194
L193:
	;
	v1326 = v1322
	goto L194
L194:
	;
	v1327 = base.F64_add(v1299, v1326)
	if base.F64_lt(v1327, float64(0)) != 0 {
		v1335 = v1312
		goto L191
	} else {
		goto L195
	}
L195:
	;
	if base.F64_gt(v1327, float64(1)) == int32(0) {
		v1335 = v1327
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v1335 = float64(1)
	goto L191
L197:
	;
	v1365 = base.F64_add(v1102, base.F64_sub(v1336, v1363))
	if base.F64_lt(v1365, float64(0)) != 0 {
		v1373 = float64(0)
		goto L182
	} else {
		goto L207
	}
L198:
	;
	v1352 = base.F64_sub(float64(1), v1339)
	if base.F64_lt(v1352, v1350) != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	if base.F64_gt(v1342, float64(1)) == int32(0) {
		v1350 = v1342
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1350 = float64(1)
	goto L198
L201:
	;
	goto L197
L202:
	;
	v1354 = v1352
	goto L204
L203:
	;
	v1354 = v1350
	goto L204
L204:
	;
	v1355 = base.F64_add(v1337, v1354)
	if base.F64_lt(v1355, float64(0)) != 0 {
		v1363 = v1340
		goto L201
	} else {
		goto L205
	}
L205:
	;
	if base.F64_gt(v1355, float64(1)) == int32(0) {
		v1363 = v1355
		goto L201
	} else {
		goto L206
	}
L206:
	;
	v1363 = float64(1)
	goto L201
L207:
	;
	if base.F64_gt(v1365, float64(1)) == int32(0) {
		v1373 = v1365
		goto L182
	} else {
		goto L208
	}
L208:
	;
	v1373 = float64(1)
	goto L182
L209:
	;
	goto L158
L210:
	;
	v1463 = int32(0)
	v1470 = m.G0
	v1472 = v1470 - int32(16)
	m.G0 = v1472
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+36))
	if v1474 != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v4867 = v1437
	v4885 = v1455
	goto L212
L212:
	;
	m.G0 = v4867 + int32(48)
	return v4885
L213:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	v1490 = float64(1)
	v1491 = int32(0)
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+112))
	if v1493 == v1491 {
		v1635 = v1491
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+68))
	v1488 = v1474 + v1475<<(uint(int32(2))%32)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+4))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+52))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+12))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+68))
	v1488 = v1481 + v1482<<(uint(int32(2))%32) - int32(4)
	goto L213
L217:
	;
	if v1635 != 0 {
		goto L228
	} else {
		goto L229
	}
L218:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1493)+4))
	if v1496 <= int32(0) {
		v1563 = v1491
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1635 = v1563
	goto L217
L220:
	;
	v1499 = int32(0)
	if v1499 < v1496 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1502 = v1496
	goto L223
L222:
	;
	v1502 = v1499
	goto L223
L223:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1493)+12))
	v1505 = int32(0)
	goto L224
L224:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1503+v1505<<(uint(int32(2))%32))))
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+16)))
	v1551 = base.B2i32(v1549 == int32(102))
	if v1549 == int32(102) {
		v1563 = v1551
		goto L219
	} else {
		goto L226
	}
L225:
	;
	v1563 = v1551
	goto L219
L226:
	;
	v1553 = v1505 + int32(1)
	if v1553 != v1502 {
		v1505 = v1553
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	if v1422 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v4823 = v1437
	v4825 = v1472
	v4841 = v1455
	v4842 = v1490
	goto L230
L230:
	;
	m.G0 = v4825 + int32(16)
	v4867 = v4823
	v4885 = base.F64_mul(v4841, v4842)
	goto L212
L231:
	;
	v2076 = int32(0)
	if v2051 == v2076 {
		goto L278
	} else {
		goto L279
	}
L232:
	;
	v1641 = F_palloc(m, int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L21
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1647 = v1422 + int32(4)
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+4))
	v1651 = F_palloc(m, v1648<<(uint(int32(1))%32))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L21
	} else {
		goto L237
	}
L235:
	;
	v1644 = F_palloc(m, int32(0))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L21
	} else {
		goto L236
	}
L236:
	;
	v2037 = v1463
	v2043 = v1644
	v2046 = v1641
	v2049 = v1463
	v2051 = v1463
	v2057 = v1463
	v2058 = v1422 + int32(4)
	goto L231
L237:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+4))
	v1656 = F_palloc(m, v1653<<(uint(int32(2))%32))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L21
	} else {
		goto L238
	}
L238:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+4))
	if int32(0) < v1658 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1663 = v1463
	v1670 = int32(0)
	goto L242
L240:
	;
	v1928 = v1463
	goto L241
L241:
	;
	v1967 = int32(16)
	v1973 = int32(0)
	v1975 = base.B2i32(v1973 < v1928)
	if v1973 < v1928 {
		goto L265
	} else {
		goto L266
	}
L242:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+12))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1702+v1670<<(uint(int32(2))%32))))
	v1707 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1472)+8)) = v1707
	v1711 = v1651 + v1670<<(uint(int32(1))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1711))) = uint16(v1707)
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1427)))
	v1715 = F_bms_is_member(m, v1670, v1714)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L21
	} else {
		goto L245
	}
L243:
	;
	v1928 = v1884
	goto L241
L244:
	;
	v1924 = v1670 + int32(1)
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1647)))
	if v1924 < v1925 {
		v1663 = v1884
		v1670 = v1924
		goto L242
	} else {
		goto L264
	}
L245:
	;
	if v1715 != 0 {
		v1884 = v1663
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+68))
	v1720 = F_dependency_is_compatible_clause(m, v1706, v1717, v1472+int32(14))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L21
	} else {
		goto L248
	}
L247:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1711))) = uint16(v1842)
	v1884 = v1843
	goto L244
L248:
	;
	if v1720 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1472)+14)))
	v1842 = v1722
	v1843 = v1663
	goto L247
L250:
	;
	goto L251
L251:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+112))
	v1726 = F_dependency_is_compatible_expression(m, v1706, v1723, v1472+int32(8))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L21
	} else {
		goto L252
	}
L252:
	;
	if v1726 == int32(0) {
		v1884 = v1663
		goto L244
	} else {
		goto L253
	}
L253:
	;
	v1730 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1472)+14)) = uint16(v1730)
	if v1663 <= v1730 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1656+v1663<<(uint(int32(2))%32)))) = v1836
	v1842 = v1663 ^ int32(-1)
	v1843 = v1663 + int32(1)
	goto L247
L255:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+8))
	v1736 = int32(0)
	goto L256
L256:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1656+v1736<<(uint(int32(2))%32))))
	v1780 = F_equal(m, v1779, v1735)
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L21
	} else {
		goto L258
	}
L257:
	;
	v1787 = int32(_a_F_statext_clauselist_selectivity_0)
	if v1736&v1787 == v1787 {
		goto L254
	} else {
		goto L263
	}
L258:
	;
	if v1780 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1785 = v1736 + int32(1)
	if v1785 != v1663 {
		v1736 = v1785
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
	v1842 = v1736 ^ int32(-1)
	v1843 = v1663
	goto L247
L264:
	;
	goto L243
L265:
	;
	v1976 = (v1928<<(uint(v1967)%32) + int32(_a_F_statext_clauselist_selectivity_1)) >> (uint(v1967) % 32)
	goto L267
L266:
	;
	v1976 = v1973
	goto L267
L267:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1647)))
	if v1977 <= int32(0) {
		v2037 = v1928
		v2043 = v1656
		v2046 = v1651
		v2049 = v1976
		v2051 = v1463
		v2057 = v1975
		v2058 = v1647
		goto L231
	} else {
		goto L268
	}
L268:
	;
	v1989 = int32(0)
	v1996 = v1463
	goto L269
L269:
	;
	v2023 = v1651 + v1989<<(uint(int32(1))%32)
	v2024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2023))))
	if v2024 != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v2037 = v1928
	v2043 = v1656
	v2046 = v1651
	v2049 = v1976
	v2051 = v2031
	v2057 = v1975
	v2058 = v1647
	goto L231
L271:
	;
	v2025 = v2024 + v1976
	*(*uint16)(unsafe.Add(mBase, uint32(v2023))) = uint16(v2025)
	v2028 = F_bms_add_member(m, v1996, base.I32_extend16_s(v2025))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L21
	} else {
		goto L274
	}
L272:
	;
	v2031 = v1996
	goto L273
L273:
	;
	v2033 = v1989 + int32(1)
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1647)))
	if v2033 < v2034 {
		v1989 = v2033
		v1996 = v2031
		goto L269
	} else {
		goto L275
	}
L274:
	;
	v2031 = v2028
	goto L273
L275:
	;
	goto L270
L276:
	;
	F_pfree(m, v4772)
	mBase = m.M
	v4806 = m.ExcPending
	if v4806 != 0 {
		goto L21
	} else {
		goto L595
	}
L277:
	;
	if v2121 != int32(2) {
		goto L293
	} else {
		goto L294
	}
L278:
	;
	v2121 = int32(0)
	goto L277
L279:
	;
	goto L280
L280:
	;
	v2084 = int32(1)
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2051)+4))
	if v2085 <= v2084 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v2088 = v2084
	goto L283
L282:
	;
	v2088 = v2085
	goto L283
L283:
	;
	v2092 = int32(0)
	v2094 = v2076
	goto L284
L284:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2051+int32(8)+v2092<<(uint(int32(2))%32))))
	if v2101 != 0 {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v2121 = v2113
	goto L277
L286:
	;
	goto L285
L287:
	;
	v2102 = int32(2)
	if v2094 != 0 {
		v2113 = v2102
		goto L286
	} else {
		goto L290
	}
L288:
	;
	v2108 = v2094
	goto L289
L289:
	;
	v2110 = v2092 + int32(1)
	if v2110 != v2088 {
		v2092 = v2110
		v2094 = v2108
		goto L284
	} else {
		goto L292
	}
L290:
	;
	v2103 = int32(1)
	if base.Ui32(v2103) < base.Ui32(base.I32_popcnt(v2101)) {
		v2113 = v2102
		goto L286
	} else {
		goto L291
	}
L291:
	;
	v2108 = v2103
	goto L289
L292:
	;
	v2113 = v2108
	goto L286
L293:
	;
	F_bms_free(m, v2051)
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L21
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+112))
	if v2126 != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v4772 = v2046
	v4781 = v1437
	v4783 = v1472
	v4799 = v1455
	v4800 = v1490
	goto L276
L297:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2126)+4))
	v2131 = v2127 << (uint(int32(2)) % 32)
	goto L299
L298:
	;
	v2131 = int32(0)
	goto L299
L299:
	;
	v2132 = F_palloc(m, v2131)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L21
	} else {
		goto L300
	}
L300:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+112))
	if v2134 != 0 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v3188 = F_palloc(m, v3141)
	mBase = m.M
	v3189 = m.ExcPending
	if v3189 != 0 {
		goto L21
	} else {
		goto L418
	}
L302:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2134)+4))
	if v2135 <= int32(0) {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	v3149 = v2043
	v3152 = v2046
	v3157 = v2051
	v3158 = v1437
	v3160 = v1472
	v3166 = v2132
	v3176 = v1455
	v3177 = v1490
	goto L304
L304:
	;
	F_pfree(m, v3166)
	mBase = m.M
	v3183 = m.ExcPending
	if v3183 != 0 {
		goto L21
	} else {
		goto L415
	}
L305:
	;
	v3103 = v1423
	v3104 = v1424
	v3105 = v1425
	v3107 = v1427
	v3108 = v2043
	v3111 = v2046
	v3113 = v1422
	v3116 = v2051
	v3117 = v1437
	v3119 = v1472
	v3121 = v1463
	v3123 = v2058
	v3124 = v1421
	v3125 = v2132
	v3135 = v1455
	v3136 = v1490
	v3141 = int32(0)
	goto L307
L306:
	;
	v2140 = v2037
	v2141 = v1423
	v2142 = v1424
	v2143 = v1425
	v2145 = v1427
	v2146 = v2043
	v2149 = v2046
	v2151 = v1422
	v2152 = v2049
	v2154 = v2051
	v2155 = v1437
	v2156 = v1463
	v2157 = v1472
	v2158 = v1463
	v2159 = v1463
	v2160 = v2057
	v2161 = v2058
	v2162 = v1421
	v2163 = v2132
	v2164 = v2134
	v2166 = v1489
	v2173 = v1455
	v2174 = v1490
	goto L308
L307:
	;
	if v3121 != 0 {
		goto L301
	} else {
		goto L414
	}
L308:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v2164)+12))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2179+v2156<<(uint(int32(2))%32))))
	v2184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2183)+16)))
	if v2184 != int32(102) {
		v3056 = v2140
		v3057 = v2141
		v3058 = v2142
		v3059 = v2143
		v3061 = v2145
		v3062 = v2146
		v3065 = v2149
		v3067 = v2151
		v3068 = v2152
		v3070 = v2154
		v3071 = v2155
		v3073 = v2157
		v3074 = v2158
		v3075 = v2159
		v3076 = v2160
		v3077 = v2161
		v3078 = v2162
		v3079 = v2163
		v3080 = v2164
		v3082 = v2166
		v3089 = v2173
		v3090 = v2174
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v3103 = v3057
	v3104 = v3058
	v3105 = v3059
	v3107 = v3061
	v3108 = v3062
	v3111 = v3065
	v3113 = v3067
	v3116 = v3070
	v3117 = v3071
	v3119 = v3073
	v3121 = v3075
	v3123 = v3077
	v3124 = v3078
	v3125 = v3079
	v3135 = v3089
	v3136 = v3090
	v3141 = v3074 << (uint(int32(2)) % 32)
	goto L307
L310:
	;
	v3096 = v2156 + int32(1)
	v3097 = *(*int32)(unsafe.Add(mBase, uint32(v3080)+4))
	if v3096 < v3097 {
		v2140 = v3056
		v2141 = v3057
		v2142 = v3058
		v2143 = v3059
		v2145 = v3061
		v2146 = v3062
		v2149 = v3065
		v2151 = v3067
		v2152 = v3068
		v2154 = v3070
		v2155 = v3071
		v2156 = v3096
		v2157 = v3073
		v2158 = v3074
		v2159 = v3075
		v2160 = v3076
		v2161 = v3077
		v2162 = v3078
		v2163 = v3079
		v2164 = v3080
		v2166 = v3082
		v2173 = v3089
		v2174 = v3090
		goto L308
	} else {
		goto L413
	}
L311:
	;
	v2187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2183)+8)))
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+20)))
	if v2187 != v2188 {
		v3056 = v2140
		v3057 = v2141
		v3058 = v2142
		v3059 = v2143
		v3061 = v2145
		v3062 = v2146
		v3065 = v2149
		v3067 = v2151
		v3068 = v2152
		v3070 = v2154
		v3071 = v2155
		v3073 = v2157
		v3074 = v2158
		v3075 = v2159
		v3076 = v2160
		v3077 = v2161
		v3078 = v2162
		v3079 = v2163
		v3080 = v2164
		v3082 = v2166
		v3089 = v2173
		v3090 = v2174
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v2190 = int32(0)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+20))
	if v2192 == v2190 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	if int32(0) <= v2249 {
		goto L324
	} else {
		goto L325
	}
L314:
	;
	v2249 = base.I32_ctz(v2235) | v2236<<(uint(int32(5))%32)
	goto L313
L315:
	;
	v2249 = int32(-2)
	goto L313
L316:
	;
	v2202 = base.I32_div_s(int32(0), int32(32))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+4))
	if v2203 <= v2202 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v2206 = v2192 + int32(8)
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2206+v2202<<(uint(int32(2))%32))))
	v2213 = v2210 & int32(-1)
	if v2213 != 0 {
		v2235 = v2213
		v2236 = v2202
		goto L314
	} else {
		goto L318
	}
L318:
	;
	v2215 = v2202 + int32(1)
	if v2215 == v2203 {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	v2218 = v2215
	goto L320
L320:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2206+v2218<<(uint(int32(2))%32))))
	if v2225 != 0 {
		v2235 = v2225
		v2236 = v2218
		goto L314
	} else {
		goto L322
	}
L321:
	;
	goto L315
L322:
	;
	v2227 = v2218 + int32(1)
	if v2227 != v2203 {
		v2218 = v2227
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v2260 = v2249
	v2261 = v2190
	goto L327
L325:
	;
	v2369 = v2190
	goto L326
L326:
	;
	if v2160 != 0 {
		goto L345
	} else {
		goto L346
	}
L327:
	;
	if int32(0) < base.I32_extend16_s(v2260) {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v2369 = v2300
	goto L326
L329:
	;
	v2297 = F_bms_is_member(m, base.I32_extend16_s(v2260+v2152), v2154)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L21
	} else {
		goto L332
	}
L330:
	;
	v2300 = v2261
	goto L331
L331:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+20))
	if v2301 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	v2300 = v2297 + v2261
	goto L331
L333:
	;
	if int32(0) <= v2357 {
		v2260 = v2357
		v2261 = v2300
		goto L327
	} else {
		goto L344
	}
L334:
	;
	v2357 = base.I32_ctz(v2343) | v2344<<(uint(int32(5))%32)
	goto L333
L335:
	;
	v2357 = int32(-2)
	goto L333
L336:
	;
	v2308 = v2260 + int32(1)
	v2310 = base.I32_div_s(v2308, int32(32))
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v2301)+4))
	if v2311 <= v2310 {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v2314 = v2301 + int32(8)
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v2314+v2310<<(uint(int32(2))%32))))
	v2321 = v2318 & (int32(-1) << (uint(v2308) % 32))
	if v2321 != 0 {
		v2343 = v2321
		v2344 = v2310
		goto L334
	} else {
		goto L338
	}
L338:
	;
	v2323 = v2310 + int32(1)
	if v2323 == v2311 {
		goto L335
	} else {
		goto L339
	}
L339:
	;
	v2326 = v2323
	goto L340
L340:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2314+v2326<<(uint(int32(2))%32))))
	if v2333 != 0 {
		v2343 = v2333
		v2344 = v2326
		goto L334
	} else {
		goto L342
	}
L341:
	;
	goto L335
L342:
	;
	v2335 = v2326 + int32(1)
	if v2335 != v2311 {
		v2326 = v2335
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
	v2412 = int32(0)
	v2415 = v2190
	goto L348
L346:
	;
	v2561 = v2190
	goto L347
L347:
	;
	if v2369+v2561 < int32(2) {
		v3056 = v2140
		v3057 = v2141
		v3058 = v2142
		v3059 = v2143
		v3061 = v2145
		v3062 = v2146
		v3065 = v2149
		v3067 = v2151
		v3068 = v2152
		v3070 = v2154
		v3071 = v2155
		v3073 = v2157
		v3074 = v2158
		v3075 = v2159
		v3076 = v2160
		v3077 = v2161
		v3078 = v2162
		v3079 = v2163
		v3080 = v2164
		v3082 = v2166
		v3089 = v2173
		v3090 = v2174
		goto L310
	} else {
		goto L358
	}
L348:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+24))
	if v2441 == int32(0) {
		v2518 = v2415
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v2561 = v2518
	goto L347
L350:
	;
	v2545 = v2412 + int32(1)
	if v2545 != v2140 {
		v2412 = v2545
		v2415 = v2518
		goto L348
	} else {
		goto L357
	}
L351:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+4))
	if v2444 <= int32(0) {
		v2518 = v2415
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v2459 = int32(0)
	v2465 = v2415
	goto L353
L353:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+12))
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2491+v2459<<(uint(int32(2))%32))))
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v2146+v2412<<(uint(int32(2))%32))))
	v2497 = F_equal(m, v2495, v2496)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L21
	} else {
		goto L355
	}
L354:
	;
	v2518 = v2499
	goto L350
L355:
	;
	v2499 = v2497 + v2465
	v2501 = v2459 + int32(1)
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+4))
	if v2501 < v2502 {
		v2459 = v2501
		v2465 = v2499
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
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+4))
	v2591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+20)))
	v2592 = m.G0
	v2594 = v2592 - int32(32)
	m.G0 = v2594
	v2597 = F_SearchSysCache2(m, int32(62), v2590, v2591)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L21
	} else {
		goto L361
	}
L359:
	;
	if v2160 != 0 {
		goto L377
	} else {
		goto L378
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L21
	} else {
		goto L373
	}
L361:
	;
	if v2597 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2603 = F_SysCacheGetAttr(m, int32(62), v2597, int32(4), v2594+int32(31))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
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
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L21
	} else {
		goto L370
	}
L365:
	;
	v2605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2594)+31)))
	if v2605 == int32(1) {
		goto L360
	} else {
		goto L366
	}
L366:
	;
	v2608 = F_pg_detoast_datum_packed(m, v2603)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L21
	} else {
		goto L367
	}
L367:
	;
	v2610 = F_statext_dependencies_deserialize(m, v2608)
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L21
	} else {
		goto L368
	}
L368:
	;
	F_ReleaseCatCache(m, v2597)
	mBase = m.M
	v2613 = m.ExcPending
	if v2613 != 0 {
		goto L21
	} else {
		goto L369
	}
L369:
	;
	m.G0 = v2594 + int32(32)
	goto L359
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2594))) = v2590
	F_errmsg_internal(m, int32(_a_F_statext_clauselist_selectivity_2), v2594)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L21
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(_a_F_statext_clauselist_selectivity_3), int32(630), int32(_a_F_statext_clauselist_selectivity_4))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+20)) = v2590
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+16)) = int32(102)
	F_errmsg_internal(m, int32(_a_F_statext_clauselist_selectivity_5), v2594+int32(16))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L21
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_statext_clauselist_selectivity_3), int32(637), int32(_a_F_statext_clauselist_selectivity_4))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
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
	if v3016 == int32(0) {
		v3056 = v2140
		v3057 = v2141
		v3058 = v2142
		v3059 = v2143
		v3061 = v2145
		v3062 = v2146
		v3065 = v2149
		v3067 = v2151
		v3068 = v2152
		v3070 = v2154
		v3071 = v2155
		v3073 = v2157
		v3074 = v2158
		v3075 = v2159
		v3076 = v2160
		v3077 = v2161
		v3078 = v2162
		v3079 = v2163
		v3080 = v2164
		v3082 = v2166
		v3089 = v2173
		v3090 = v2174
		goto L310
	} else {
		goto L412
	}
L377:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2610)+8))
	if v2649 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+24))
	if v2647 != 0 {
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2610)+8))
	v3016 = v2648
	goto L376
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2610)+8)) = v2975
	v3016 = v2975
	goto L376
L381:
	;
	v2975 = int32(0)
	goto L380
L382:
	;
	goto L383
L383:
	;
	v2654 = v2610 + int32(12)
	v2655 = int32(0)
	v2666 = v2655
	v2668 = v2655
	goto L384
L384:
	;
	v2699 = v2654 + v2666<<(uint(int32(2))%32)
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v2699)))
	v2701 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2700)+8)))
	if int32(0) < v2701 {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	v2975 = v2931
	goto L380
L386:
	;
	v2961 = v2666 + int32(1)
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v2610)+8))
	if base.Ui32(v2961) < base.Ui32(v2962) {
		v2666 = v2961
		v2668 = v2931
		goto L384
	} else {
		goto L411
	}
L387:
	;
	v2721 = int32(0)
	goto L390
L388:
	;
	goto L389
L389:
	;
	if v2666 != v2668 {
		goto L408
	} else {
		goto L409
	}
L390:
	;
	v2749 = v2700 + int32(10) + v2721<<(uint(int32(1))%32)
	v2750 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2749))))
	if int32(0) < v2750 {
		goto L393
	} else {
		goto L394
	}
L391:
	;
	goto L389
L392:
	;
	v2869 = v2721 + int32(1)
	v2870 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2700)+8)))
	if v2869 < v2870 {
		v2721 = v2869
		goto L390
	} else {
		goto L407
	}
L393:
	;
	v2753 = v2750 + v2152
	*(*uint16)(unsafe.Add(mBase, uint32(v2749))) = uint16(v2753)
	v2756 = F_bms_is_member(m, base.I32_extend16_s(v2753), v2154)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L21
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	if v2160 == int32(0) {
		v2931 = v2668
		goto L386
	} else {
		goto L398
	}
L396:
	;
	if v2756 != 0 {
		goto L392
	} else {
		goto L397
	}
L397:
	;
	v2931 = v2668
	goto L386
L398:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+24))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2760)+12))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2761+(v2750^int32(-1))<<(uint(int32(2))%32))))
	v2777 = int32(0)
	goto L399
L399:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2146+v2777<<(uint(int32(2))%32))))
	v2813 = F_equal(m, v2812, v2767)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L21
	} else {
		goto L401
	}
L400:
	;
	v2822 = v2152 + (v2777 ^ int32(-1))
	if v2822&int32(_a_F_statext_clauselist_selectivity_0) == int32(0) {
		v2931 = v2668
		goto L386
	} else {
		goto L406
	}
L401:
	;
	if v2813 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v2818 = v2777 + int32(1)
	if v2818 != v2140 {
		v2777 = v2818
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
	v2931 = v2668
	goto L386
L406:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2749))) = uint16(v2822)
	goto L392
L407:
	;
	goto L391
L408:
	;
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v2699)))
	*(*int32)(unsafe.Add(mBase, uint32(v2654+v2668<<(uint(int32(2))%32)))) = v2916
	goto L410
L409:
	;
	goto L410
L410:
	;
	v2931 = v2668 + int32(1)
	goto L386
L411:
	;
	goto L385
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2163+v2159<<(uint(int32(2))%32)))) = v2610
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(v2610)+8))
	v3056 = v2140
	v3057 = v2141
	v3058 = v2142
	v3059 = v2143
	v3061 = v2145
	v3062 = v2146
	v3065 = v2149
	v3067 = v2151
	v3068 = v2152
	v3070 = v2154
	v3071 = v2155
	v3073 = v2157
	v3074 = v3053 + v2158
	v3075 = v2159 + int32(1)
	v3076 = v2160
	v3077 = v2161
	v3078 = v2162
	v3079 = v2163
	v3080 = v2164
	v3082 = v2166
	v3089 = v2173
	v3090 = v2174
	goto L310
L413:
	;
	goto L309
L414:
	;
	v3149 = v3108
	v3152 = v3111
	v3157 = v3116
	v3158 = v3117
	v3160 = v3119
	v3166 = v3125
	v3176 = v3135
	v3177 = v3136
	goto L304
L415:
	;
	F_bms_free(m, v3157)
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L21
	} else {
		goto L416
	}
L416:
	;
	F_pfree(m, v3152)
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L21
	} else {
		goto L417
	}
L417:
	;
	v4772 = v3149
	v4781 = v3158
	v4783 = v3160
	v4799 = v3176
	v4800 = v3177
	goto L276
L418:
	;
	v3190 = int32(0)
	if v3116 == v3190 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if int32(0) < v3121 {
		goto L432
	} else {
		goto L433
	}
L420:
	;
	v3225 = int32(0)
	goto L419
L421:
	;
	goto L422
L422:
	;
	v3197 = int32(1)
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3116)+4))
	if v3198 <= v3197 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v3201 = v3197
	goto L425
L424:
	;
	v3201 = v3198
	goto L425
L425:
	;
	v3205 = int32(0)
	v3207 = v3190
	goto L426
L426:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3116+int32(8)+v3205<<(uint(int32(2))%32))))
	if v3213 != 0 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v3225 = v3216
	goto L419
L428:
	;
	v3216 = v3207 + base.I32_popcnt(v3213)
	goto L430
L429:
	;
	v3216 = v3207
	goto L430
L430:
	;
	v3218 = v3205 + int32(1)
	if v3218 != v3201 {
		v3205 = v3218
		v3207 = v3216
		goto L426
	} else {
		goto L431
	}
L431:
	;
	goto L427
L432:
	;
	v3234 = v3225
	v3244 = v3116
	v3246 = int32(0)
	goto L435
L433:
	;
	v4732 = v3116
	v4752 = v3136
	goto L434
L434:
	;
	F_pfree(m, v3188)
	mBase = m.M
	v4758 = m.ExcPending
	if v4758 != 0 {
		goto L21
	} else {
		goto L591
	}
L435:
	;
	v3269 = int32(0)
	v3272 = v3269
	v3280 = v3269
	goto L437
L436:
	;
	if v3246 != 0 {
		goto L479
	} else {
		goto L480
	}
L437:
	;
	v3314 = v3125 + v3280<<(uint(int32(2))%32)
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v3314)))
	v3316 = *(*int32)(unsafe.Add(mBase, uint32(v3315)+8))
	if v3316 != 0 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	if v3477 != 0 {
		goto L462
	} else {
		goto L463
	}
L439:
	;
	v3318 = v3272
	v3325 = v3315
	v3330 = int32(0)
	goto L442
L440:
	;
	v3477 = v3272
	goto L441
L441:
	;
	v3517 = v3280 + int32(1)
	if v3517 != v3121 {
		v3272 = v3477
		v3280 = v3517
		goto L437
	} else {
		goto L461
	}
L442:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v3325+v3330<<(uint(int32(2))%32))+12))
	v3361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3360)+8)))
	if v3234 < v3361 {
		v3433 = v3318
		v3440 = v3325
		goto L444
	} else {
		goto L445
	}
L443:
	;
	v3477 = v3433
	goto L441
L444:
	;
	v3473 = v3330 + int32(1)
	v3474 = *(*int32)(unsafe.Add(mBase, uint32(v3440)+8))
	if base.Ui32(v3473) < base.Ui32(v3474) {
		v3318 = v3433
		v3325 = v3440
		v3330 = v3473
		goto L442
	} else {
		goto L460
	}
L445:
	;
	if v3318 == int32(0) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	if v3361 <= int32(0) {
		goto L451
	} else {
		goto L452
	}
L447:
	;
	v3365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3318)+8)))
	if v3361 < v3365 {
		v3433 = v3318
		v3440 = v3325
		goto L444
	} else {
		goto L448
	}
L448:
	;
	if v3361 != v3365 {
		goto L446
	} else {
		goto L449
	}
L449:
	;
	v3368 = *(*float64)(unsafe.Add(mBase, uint32(v3318)))
	v3369 = *(*float64)(unsafe.Add(mBase, uint32(v3360)))
	if base.F64_gt(v3368, v3369) != 0 {
		v3433 = v3318
		v3440 = v3325
		goto L444
	} else {
		goto L450
	}
L450:
	;
	goto L446
L451:
	;
	v3433 = v3360
	v3440 = v3325
	goto L444
L452:
	;
	goto L453
L453:
	;
	v3385 = int32(0)
	goto L455
L454:
	;
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3314)))
	v3433 = v3429
	v3440 = v3431
	goto L444
L455:
	;
	v3420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3360+int32(10)+v3385<<(uint(int32(1))%32)))))
	v3421 = F_bms_is_member(m, v3420, v3244)
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L21
	} else {
		goto L457
	}
L456:
	;
	v3429 = v3360
	goto L454
L457:
	;
	if v3421 == int32(0) {
		v3429 = v3318
		goto L454
	} else {
		goto L458
	}
L458:
	;
	v3426 = v3385 + int32(1)
	v3427 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3360)+8)))
	if v3426 < v3427 {
		v3385 = v3426
		goto L455
	} else {
		goto L459
	}
L459:
	;
	goto L456
L460:
	;
	goto L443
L461:
	;
	goto L438
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3188+v3246<<(uint(int32(2))%32)))) = v3477
	v3523 = int32(1)
	v3527 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3477)+8)))
	v3531 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3477+int32(8)+v3527<<(uint(v3523)%32)))))
	v3532 = F_bms_del_member(m, v3244, v3531)
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L21
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	goto L436
L465:
	;
	v3534 = int32(0)
	if v3532 == v3534 {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	v3234 = v3569
	v3244 = v3532
	v3246 = v3246 + v3523
	goto L435
L467:
	;
	v3569 = int32(0)
	goto L466
L468:
	;
	goto L469
L469:
	;
	v3541 = int32(1)
	v3542 = *(*int32)(unsafe.Add(mBase, uint32(v3532)+4))
	if v3542 <= v3541 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v3545 = v3541
	goto L472
L471:
	;
	v3545 = v3542
	goto L472
L472:
	;
	v3549 = int32(0)
	v3551 = v3534
	goto L473
L473:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v3532+int32(8)+v3549<<(uint(int32(2))%32))))
	if v3557 != 0 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	v3569 = v3560
	goto L466
L475:
	;
	v3560 = v3551 + base.I32_popcnt(v3557)
	goto L477
L476:
	;
	v3560 = v3551
	goto L477
L477:
	;
	v3562 = v3549 + int32(1)
	if v3562 != v3545 {
		v3549 = v3562
		v3551 = v3560
		goto L473
	} else {
		goto L478
	}
L478:
	;
	goto L474
L479:
	;
	v3570 = int32(0)
	if v3570 < v3246 {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	v4662 = v3136
	goto L481
L481:
	;
	v4676 = int32(0)
	goto L587
L482:
	;
	v3574 = v3570
	v3579 = int32(0)
	goto L485
L483:
	;
	v3716 = v3570
	goto L484
L484:
	;
	v3756 = int32(0)
	if v3716 == v3756 {
		goto L496
	} else {
		goto L497
	}
L485:
	;
	v3614 = int32(0)
	v3617 = v3188 + v3579<<(uint(int32(2))%32)
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v3617)))
	v3619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3618)+8)))
	if v3614 < v3619 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	v3716 = v3673
	goto L484
L487:
	;
	v3622 = v3574
	v3630 = v3614
	v3636 = v3618
	goto L490
L488:
	;
	v3673 = v3574
	goto L489
L489:
	;
	v3714 = v3579 + int32(1)
	if v3714 != v3246 {
		v3574 = v3673
		v3579 = v3714
		goto L485
	} else {
		goto L494
	}
L490:
	;
	v3665 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3636+v3630<<(uint(int32(1))%32))+10)))
	v3666 = F_bms_add_member(m, v3622, v3665)
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L21
	} else {
		goto L492
	}
L491:
	;
	v3673 = v3666
	goto L489
L492:
	;
	v3669 = v3630 + int32(1)
	v3670 = *(*int32)(unsafe.Add(mBase, uint32(v3617)))
	v3671 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3670)+8)))
	if v3669 < v3671 {
		v3622 = v3666
		v3630 = v3669
		v3636 = v3670
		goto L490
	} else {
		goto L493
	}
L493:
	;
	goto L491
L494:
	;
	goto L486
L495:
	;
	v3794 = F_palloc(m, v3791<<(uint(int32(3))%32))
	mBase = m.M
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L21
	} else {
		goto L508
	}
L496:
	;
	v3791 = int32(0)
	goto L495
L497:
	;
	goto L498
L498:
	;
	v3763 = int32(1)
	v3764 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+4))
	if v3764 <= v3763 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v3767 = v3763
	goto L501
L500:
	;
	v3767 = v3764
	goto L501
L501:
	;
	v3771 = int32(0)
	v3773 = v3756
	goto L502
L502:
	;
	v3779 = *(*int32)(unsafe.Add(mBase, uint32(v3716+int32(8)+v3771<<(uint(int32(2))%32))))
	if v3779 != 0 {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	v3791 = v3782
	goto L495
L504:
	;
	v3782 = v3773 + base.I32_popcnt(v3779)
	goto L506
L505:
	;
	v3782 = v3773
	goto L506
L506:
	;
	v3784 = v3771 + int32(1)
	if v3784 != v3767 {
		v3771 = v3784
		v3773 = v3782
		goto L502
	} else {
		goto L507
	}
L507:
	;
	goto L503
L508:
	;
	if v3716 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L509:
	;
	if int32(0) <= v3852 {
		goto L520
	} else {
		goto L521
	}
L510:
	;
	v3852 = base.I32_ctz(v3838) | v3839<<(uint(int32(5))%32)
	goto L509
L511:
	;
	v3852 = int32(-2)
	goto L509
L512:
	;
	v3805 = base.I32_div_s(int32(0), int32(32))
	v3806 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+4))
	if v3806 <= v3805 {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v3809 = v3716 + int32(8)
	v3813 = *(*int32)(unsafe.Add(mBase, uint32(v3809+v3805<<(uint(int32(2))%32))))
	v3816 = v3813 & int32(-1)
	if v3816 != 0 {
		v3838 = v3816
		v3839 = v3805
		goto L510
	} else {
		goto L514
	}
L514:
	;
	v3818 = v3805 + int32(1)
	if v3818 == v3806 {
		goto L511
	} else {
		goto L515
	}
L515:
	;
	v3821 = v3818
	goto L516
L516:
	;
	v3828 = *(*int32)(unsafe.Add(mBase, uint32(v3809+v3821<<(uint(int32(2))%32))))
	if v3828 != 0 {
		v3838 = v3828
		v3839 = v3821
		goto L510
	} else {
		goto L518
	}
L517:
	;
	goto L511
L518:
	;
	v3830 = v3821 + int32(1)
	if v3830 != v3806 {
		v3821 = v3830
		goto L516
	} else {
		goto L519
	}
L519:
	;
	goto L517
L520:
	;
	v3861 = v3852
	v3869 = int32(0)
	goto L523
L521:
	;
	goto L522
L522:
	;
	v4117 = v3246 - int32(1)
	if int32(0) <= v4117 {
		goto L551
	} else {
		goto L552
	}
L523:
	;
	if v3113 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	goto L522
L525:
	;
	v4013 = F_clauselist_selectivity_ext(m, v3124, v3980, v3103, v3104, v3105, int32(0))
	mBase = m.M
	v4014 = m.ExcPending
	if v4014 != 0 {
		goto L21
	} else {
		goto L538
	}
L526:
	;
	v3980 = int32(0)
	goto L525
L527:
	;
	goto L528
L528:
	;
	v3899 = int32(0)
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3123)))
	if v3902 <= v3899 {
		v3980 = v3899
		goto L525
	} else {
		goto L529
	}
L529:
	;
	v3906 = v3902
	v3913 = v3899
	v3916 = v3899
	v3919 = int32(-1)
	goto L530
L530:
	;
	v3945 = int32(1)
	v3946 = v3919 + v3945
	v3950 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3111+v3946<<(uint(v3945)%32)))))
	if v3950 == v3861 {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	v3980 = v3965
	goto L525
L532:
	;
	v3952 = *(*int32)(unsafe.Add(mBase, uint32(v3113)+12))
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v3952+v3913<<(uint(int32(2))%32))))
	v3957 = F_lappend(m, v3916, v3956)
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		goto L21
	} else {
		goto L535
	}
L533:
	;
	v3964 = v3906
	v3965 = v3916
	goto L534
L534:
	;
	v3967 = v3913 + int32(1)
	if v3967 < v3964 {
		v3906 = v3964
		v3913 = v3967
		v3916 = v3965
		v3919 = v3946
		goto L530
	} else {
		goto L537
	}
L535:
	;
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v3107)))
	v3960 = F_bms_add_member(m, v3959, v3946)
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L21
	} else {
		goto L536
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3107))) = v3960
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v3113)+4))
	v3964 = v3963
	v3965 = v3957
	goto L534
L537:
	;
	goto L531
L538:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v3794+v3869<<(uint(int32(3))%32)))) = v4013
	if v3716 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L539:
	;
	if int32(0) <= v4073 {
		v3861 = v4073
		v3869 = v3869 + int32(1)
		goto L523
	} else {
		goto L550
	}
L540:
	;
	v4073 = base.I32_ctz(v4059) | v4060<<(uint(int32(5))%32)
	goto L539
L541:
	;
	v4073 = int32(-2)
	goto L539
L542:
	;
	v4024 = v3861 + int32(1)
	v4026 = base.I32_div_s(v4024, int32(32))
	v4027 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+4))
	if v4027 <= v4026 {
		goto L541
	} else {
		goto L543
	}
L543:
	;
	v4030 = v3716 + int32(8)
	v4034 = *(*int32)(unsafe.Add(mBase, uint32(v4030+v4026<<(uint(int32(2))%32))))
	v4037 = v4034 & (int32(-1) << (uint(v4024) % 32))
	if v4037 != 0 {
		v4059 = v4037
		v4060 = v4026
		goto L540
	} else {
		goto L544
	}
L544:
	;
	v4039 = v4026 + int32(1)
	if v4039 == v4027 {
		goto L541
	} else {
		goto L545
	}
L545:
	;
	v4042 = v4039
	goto L546
L546:
	;
	v4049 = *(*int32)(unsafe.Add(mBase, uint32(v4030+v4042<<(uint(int32(2))%32))))
	if v4049 != 0 {
		v4059 = v4049
		v4060 = v4042
		goto L540
	} else {
		goto L548
	}
L547:
	;
	goto L541
L548:
	;
	v4051 = v4042 + int32(1)
	if v4051 != v4027 {
		v4042 = v4051
		goto L546
	} else {
		goto L549
	}
L549:
	;
	goto L547
L550:
	;
	goto L524
L551:
	;
	v4125 = v4117
	goto L554
L552:
	;
	goto L553
L553:
	;
	if int32(0) < v3791 {
		goto L569
	} else {
		goto L570
	}
L554:
	;
	v4160 = float64(1)
	v4161 = int32(0)
	v4162 = int32(2)
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v3188+v4125<<(uint(v4162)%32))))
	v4166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4165)+8)))
	if v4162 <= v4166 {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	goto L553
L556:
	;
	v4179 = v4161
	v4206 = v4160
	goto L559
L557:
	;
	v4236 = v4161
	v4263 = v4160
	goto L558
L558:
	;
	v4271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4165+v4236<<(uint(int32(1))%32))+10)))
	v4272 = F_bms_member_index(m, v3716, v4271)
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L21
	} else {
		goto L563
	}
L559:
	;
	v4214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4165+int32(10)+v4179<<(uint(int32(1))%32)))))
	v4215 = F_bms_member_index(m, v3716, v4214)
	mBase = m.M
	v4216 = m.ExcPending
	if v4216 != 0 {
		goto L21
	} else {
		goto L561
	}
L560:
	;
	v4236 = v4223
	v4263 = v4221
	goto L558
L561:
	;
	v4220 = *(*float64)(unsafe.Add(mBase, uint32(v3794+v4215<<(uint(int32(3))%32))))
	v4221 = base.F64_mul(v4206, v4220)
	v4222 = int32(1)
	v4223 = v4179 + v4222
	v4224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4165)+8)))
	if v4223 < v4224-v4222 {
		v4179 = v4223
		v4206 = v4221
		goto L559
	} else {
		goto L562
	}
L562:
	;
	goto L560
L563:
	;
	v4274 = *(*float64)(unsafe.Add(mBase, uint32(v4165)))
	v4277 = v3794 + v4272<<(uint(int32(3))%32)
	v4278 = *(*float64)(unsafe.Add(mBase, uint32(v4277)))
	if base.F64_le(v4263, v4278) == int32(0) {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v4284 = base.F64_div(base.F64_mul(v4278, v4274), v4263)
	goto L566
L565:
	;
	v4284 = v4274
	goto L566
L566:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v4277))) = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v4274), v4278), v4284)
	if int32(0) < v4125 {
		v4125 = v4125 - int32(1)
		goto L554
	} else {
		goto L567
	}
L567:
	;
	goto L555
L568:
	;
	F_pfree(m, v3794)
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L21
	} else {
		goto L585
	}
L569:
	;
	v4337 = v3791 & int32(3)
	v4338 = float64(1)
	v4339 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3791) {
		goto L573
	} else {
		goto L574
	}
L570:
	;
	goto L571
L571:
	;
	v4622 = float64(1)
	goto L568
L572:
	;
	v4534 = float64(0)
	if base.F64_lt(v4527, v4534) != 0 {
		v4622 = v4534
		goto L568
	} else {
		goto L583
	}
L573:
	;
	v4351 = int32(0)
	v4354 = v4339
	v4379 = v4338
	goto L576
L574:
	;
	v4412 = v4339
	v4437 = v4338
	goto L575
L575:
	;
	v4445 = v4339
	v4452 = v4412
	v4477 = v4437
	goto L580
L576:
	;
	v4388 = v3794 + v4354<<(uint(int32(3))%32)
	v4389 = *(*float64)(unsafe.Add(mBase, uint32(v4388)))
	v4391 = *(*float64)(unsafe.Add(mBase, uint32(v4388)+8))
	v4393 = *(*float64)(unsafe.Add(mBase, uint32(v4388)+16))
	v4395 = *(*float64)(unsafe.Add(mBase, uint32(v4388)+24))
	v4396 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(v4379, v4389), v4391), v4393), v4395)
	v4397 = int32(4)
	v4398 = v4354 + v4397
	v4400 = v4351 + v4397
	if v4400 != v3791&int32(2147483644) {
		v4351 = v4400
		v4354 = v4398
		v4379 = v4396
		goto L576
	} else {
		goto L578
	}
L577:
	;
	if v4337 == int32(0) {
		v4527 = v4396
		goto L572
	} else {
		goto L579
	}
L578:
	;
	goto L577
L579:
	;
	v4412 = v4398
	v4437 = v4396
	goto L575
L580:
	;
	v4487 = *(*float64)(unsafe.Add(mBase, uint32(v3794+v4452<<(uint(int32(3))%32))))
	v4488 = base.F64_mul(v4477, v4487)
	v4489 = int32(1)
	v4492 = v4445 + v4489
	if v4492 != v4337 {
		v4445 = v4492
		v4452 = v4452 + v4489
		v4477 = v4488
		goto L580
	} else {
		goto L582
	}
L581:
	;
	v4527 = v4488
	goto L572
L582:
	;
	goto L581
L583:
	;
	if base.F64_gt(v4527, float64(1)) == int32(0) {
		v4622 = v4527
		goto L568
	} else {
		goto L584
	}
L584:
	;
	goto L571
L585:
	;
	F_bms_free(m, v3716)
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		goto L21
	} else {
		goto L586
	}
L586:
	;
	v4662 = v4622
	goto L481
L587:
	;
	v4711 = *(*int32)(unsafe.Add(mBase, uint32(v3125+v4676<<(uint(int32(2))%32))))
	F_pfree(m, v4711)
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L21
	} else {
		goto L589
	}
L588:
	;
	v4732 = v3244
	v4752 = v4662
	goto L434
L589:
	;
	v4715 = v4676 + int32(1)
	if v4715 != v3121 {
		v4676 = v4715
		goto L587
	} else {
		goto L590
	}
L590:
	;
	goto L588
L591:
	;
	F_pfree(m, v3125)
	mBase = m.M
	v4760 = m.ExcPending
	if v4760 != 0 {
		goto L21
	} else {
		goto L592
	}
L592:
	;
	F_bms_free(m, v4732)
	mBase = m.M
	v4762 = m.ExcPending
	if v4762 != 0 {
		goto L21
	} else {
		goto L593
	}
L593:
	;
	F_pfree(m, v3111)
	mBase = m.M
	v4764 = m.ExcPending
	if v4764 != 0 {
		goto L21
	} else {
		goto L594
	}
L594:
	;
	v4772 = v3108
	v4781 = v3117
	v4783 = v3119
	v4799 = v3135
	v4800 = v4752
	goto L276
L595:
	;
	v4823 = v4781
	v4825 = v4783
	v4841 = v4799
	v4842 = v4800
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
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
	if v13 != int32(318) {
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
	v138 = base.I32_div_s(int32(0), int32(32))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v139 <= v138 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v142 = v128 + int32(8)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v138<<(uint(int32(2))%32))))
	v149 = v146 & int32(-1)
	if v149 != 0 {
		v171 = v149
		v172 = v138
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v151 = v138 + int32(1)
	if v151 == v139 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v154 = v151
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
	v210 = base.I32_div_s(v208, int32(32))
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
