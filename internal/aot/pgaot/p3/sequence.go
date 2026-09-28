package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SequenceSyncWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v688 int32
	_ = v688
	var v691 int64
	_ = v691
	var v694 int64
	_ = v694
	var v697 int64
	_ = v697
	var v700 int64
	_ = v700
	var v703 int64
	_ = v703
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v875 int32
	_ = v875
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int64
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1064 int64
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int64
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int64
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int64
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int64
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int64
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1252 int32
	_ = v1252
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int64
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1272 int64
	_ = v1272
	var v1274 int64
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1440 int64
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1458 int32
	_ = v1458
	var v1459 int64
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1471 int32
	_ = v1471
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int64
	_ = v1484
	var v1498 int32
	_ = v1498
	var v1508 int32
	_ = v1508
	var v1521 int32
	_ = v1521
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int64
	_ = v1533
	var v1546 int32
	_ = v1546
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1704 int32
	_ = v1704
	var v1711 int32
	_ = v1711
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1770 int32
	_ = v1770
	var v1779 int32
	_ = v1779
	var v1789 int32
	_ = v1789
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
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
	var v1877 int32
	_ = v1877
	var v1885 int32
	_ = v1885
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1978 int32
	_ = v1978
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2004 int32
	_ = v2004
	var v2011 int32
	_ = v2011
	var v2021 int32
	_ = v2021
	var v2033 int32
	_ = v2033
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2059 int32
	_ = v2059
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2082 int32
	_ = v2082
	var v2093 int32
	_ = v2093
	var v2105 int32
	_ = v2105
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2149 int32
	_ = v2149
	var v2159 int32
	_ = v2159
	var v2171 int32
	_ = v2171
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2197 int32
	_ = v2197
	var v2204 int32
	_ = v2204
	var v2214 int32
	_ = v2214
	var v2224 int32
	_ = v2224
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2246 int32
	_ = v2246
	var v2256 int32
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2272 int32
	_ = v2272
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2338 int32
	_ = v2338
	var v2345 int32
	_ = v2345
	var v2388 int32
	_ = v2388
	var v2389 int64
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
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
	var v2410 int32
	_ = v2410
	var v2455 int32
	_ = v2455
	v2 = int32(0)
	F_SetConfigOption(m, int32(_a_F_SequenceSyncWorkerMain_0), int32(_a_F_SequenceSyncWorkerMain_1), int32(5), int32(10))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_SetupApplyOrSyncWorker(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v52 = m.G0
	v54 = v52 - int32(720)
	m.G0 = v54
	v71 = v2
	v72 = v2
	v73 = v2
	v74 = v2
	v75 = v2
	v76 = v2
	v78 = int32(-1)
	v97 = v2
	v98 = v2
	goto L5
L4:
	;
	F_FinishSyncWorker(m)
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L1
	} else {
		goto L356
	}
L5:
	;
	if v78 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L7:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[0]))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[1]))
	v119 = v54 + int32(288)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v54 + int32(284)
	goto L10
L8:
	;
	v125 = v76
	v126 = v97
	v127 = v98
	goto L9
L9:
	;
	goto L12
L10:
	;
	v125 = int32(0)
	v126 = v115
	v127 = v117
	goto L9
L11:
	;
	goto L6
L12:
	;
	if v125 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L11
L14:
	;
	v2388 = int32(m.ExcTag)
	v2389 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2388 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L14
	} else {
		goto L343
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[0])) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[1])) = v127
	m.G0 = v54 + int32(720)
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[1])) = v54 + int32(288)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[2]))
	v136 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v135)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_StartTransactionCommand(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[0])) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[1])) = v127
	v2262 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2262)+37)))
	if v2263 != int32(1) {
		goto L15
	} else {
		goto L341
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_maybe_reread_subscription(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	v158 = F_table_open(m, int32(_a_F_SequenceSyncWorkerMain_2), int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	v166 = v54 + int32(464)
	F_ScanKeyInit(m, v166, int32(1), int32(3), int32(184), v136)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	v177 = int32(3)
	F_ScanKeyInit(m, v54+int32(520), v177, v177, int32(61), int64(105))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	v188 = int32(0)
	v192 = F_systable_beginscan(m, v158, v188, v188, v188, int32(2), v166)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	v199 = F_systable_getnext(m, v192)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	if v199 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v206 = v74
	v208 = v199
	goto L30
L28:
	;
	v344 = v74
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_systable_endscan(m, v192)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L50
	}
L30:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[4]))
	if v244 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v344 = v337
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	F_ProcessInterrupts(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L14
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v208)+16))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+22)))
	v254 = v252 + v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v262 = F_try_table_open(m, v255, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L14
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	if v262 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v262)+48))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+119)))
	if v265 == int32(83) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v337 = F_systable_getnext(m, v192)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L14
	} else {
		goto L48
	}
L40:
	;
	v268 = int32(_a_F_SequenceSyncWorkerMain_3)
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5]))
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v280 = F_palloc0(m, int32(40))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L14
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	F_relation_close(m, v262, int32(1))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L14
	} else {
		goto L47
	}
L43:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v262)+48))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v291 = F_get_namespace_name(m, v285)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v262)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v302 = F_pstrdup(m, v294+int32(4))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v206
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	v312 = F_lappend(m, v311, v280)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v269
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7])) = v312
	goto L42
L47:
	;
	goto L39
L48:
	;
	if v337 != 0 {
		v206 = v337
		v208 = v337
		goto L30
	} else {
		goto L49
	}
L49:
	;
	goto L31
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_relation_close(m, v158, int32(1))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_CommitTransactionCommand(m)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	if v404 == int32(0) {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+38)))
	if v410 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+32)))
	v416 = v413 ^ int32(1)
	goto L56
L55:
	;
	v416 = int32(0)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v423 = v54 + int32(448)
	F_initStringInfo(m, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[8]))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v435)))
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v54)+264)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v54)+256)) = v428
	F_appendStringInfo(m, v423, int32(_a_F_SequenceSyncWorkerMain_4), v54+int32(256))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[9]))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[10]))
	v460 = int32(1)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v54)+448))
	v467 = m.T0[v451].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v459, v460, v460, v416&v460, v464, v54+int32(588))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[11])) = v467
	if v467 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L14
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v54)+448))
	F_pfree(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L14
	} else {
		goto L68
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(100663808))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v54)+588))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v491
	F_errmsg(m, int32(_a_F_SequenceSyncWorkerMain_5), v54)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(825), int32(_a_F_SequenceSyncWorkerMain_7))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	goto L11
L68:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[11]))
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	if v525 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v527 = v526
	goto L71
L70:
	;
	v527 = int32(0)
	goto L71
L71:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[9]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v536 = m.T0[v530].(func(*base.Module, int32) int32)(m, v523)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	if v536 <= int32(_a_F_SequenceSyncWorkerMain_8) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L14
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_initStringInfo(m, v54+int32(656))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L14
	} else {
		goto L80
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errmsg(m, int32(_a_F_SequenceSyncWorkerMain_9), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(471), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	goto L11
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_initStringInfo(m, v54+int32(640))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L14
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v601 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	if v601 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54)+244)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v54)+240)) = v605
	F_errmsg_internal(m, int32(_a_F_SequenceSyncWorkerMain_11), v54+int32(240))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L14
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v527 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(480), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	if v1927|(v1926|v1928)|v1931 == int32(0) {
		goto L16
	} else {
		goto L298
	}
L89:
	;
	v631 = int32(0)
	v1916 = v71
	v1917 = v72
	v1918 = v73
	v1920 = v75
	v1926 = v631
	v1927 = v631
	v1928 = v631
	v1931 = v631
	goto L88
L90:
	;
	goto L91
L91:
	;
	v635 = int32(0)
	v642 = v71
	v643 = v72
	v644 = v73
	v646 = v75
	v649 = v635
	v652 = v635
	v653 = v635
	v654 = v635
	v657 = v635
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+632)) = v688
	v691 = *(*int64)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+624)) = v691
	v694 = *(*int64)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+616)) = v694
	v697 = *(*int64)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+608)) = v697
	v700 = *(*int64)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+600)) = v700
	v703 = *(*int64)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[17]))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+592)) = v703
	F_StartTransactionCommand(m)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L14
	} else {
		goto L94
	}
L93:
	;
	v1916 = v1634
	v1917 = v1635
	v1918 = v1636
	v1920 = v1877
	v1926 = v1644
	v1927 = v1645
	v1928 = v1885
	v1931 = v1649
	goto L88
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_maybe_reread_subscription(m)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	v722 = v649
	v735 = int32(0)
	goto L96
L96:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+12))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v759+v722<<(uint(int32(2))%32))))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v54)+660))
	if int32(0) < v764 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v54)+656))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+208)) = v821
	F_appendStringInfo(m, v54+int32(640), int32(_a_F_SequenceSyncWorkerMain_12), v54+int32(208))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L14
	} else {
		goto L109
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_appendStringInfoString(m, v54+int32(656), int32(_a_F_SequenceSyncWorkerMain_13))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L14
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v783 = F_quote_literal_cstr(m, v777)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L14
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v791 = F_quote_literal_cstr(m, v785)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54)+232)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v54)+228)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v54)+224)) = v783
	F_appendStringInfo(m, v54+int32(656), int32(_a_F_SequenceSyncWorkerMain_14), v54+int32(224))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	v809 = v735 + int32(1)
	if v809 != int32(100) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v813 = v722 + int32(1)
	if v813 < v527 {
		v722 = v813
		v735 = v809
		goto L96
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L97
L108:
	;
	goto L107
L109:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[9]))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v54)+640))
	v842 = m.T0[v832].(func(*base.Module, int32, int32, int32, int32) int32)(m, v523, v838, int32(11), v54+int32(592))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	if v844 != int32(2) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L14
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v842)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v893 = F_MakeSingleTupleTableSlot(m, v886, int32(_a_F_SequenceSyncWorkerMain_15))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L14
	} else {
		goto L118
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(100663808))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v842)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54)+192)) = v864
	F_errmsg(m, int32(_a_F_SequenceSyncWorkerMain_16), v54+int32(192))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L14
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(570), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L14
	} else {
		goto L117
	}
L117:
	;
	goto L11
L118:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v903 = F_tuplestore_gettupleslot(m, v895, int32(1), int32(0), v893)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L14
	} else {
		goto L119
	}
L119:
	;
	v905 = int32(0)
	if v903 == v905 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_ExecDropSingleTupleTableSlot(m, v893)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L14
	} else {
		goto L264
	}
L121:
	;
	v908 = int32(0)
	v1634 = v642
	v1635 = v643
	v1636 = v644
	v1644 = v652
	v1645 = v653
	v1649 = v657
	v1654 = v908
	v1655 = v908
	v1656 = v908
	v1657 = v908
	v1659 = v905
	goto L120
L122:
	;
	goto L123
L123:
	;
	v912 = int32(0)
	v918 = v642
	v919 = v643
	v920 = v644
	v928 = v652
	v929 = v653
	v933 = v657
	v938 = v912
	v939 = v912
	v940 = v912
	v941 = v912
	v943 = v905
	goto L124
L124:
	;
	v959 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[4]))
	if v959 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v1634 = v1599
	v1635 = v1600
	v1636 = v1601
	v1644 = v1605
	v1645 = v1606
	v1649 = v1608
	v1654 = v1611
	v1655 = v1612
	v1656 = v1613
	v1657 = v1614
	v1659 = v1616
	goto L120
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_ProcessInterrupts(m)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L14
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v968 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[18]))
	if v968 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[18])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L14
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v980 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	if v980 <= int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v984].(func(*base.Module, int32, int32))(m, v893, int32(1))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L14
	} else {
		goto L137
	}
L135:
	;
	v994 = v980
	goto L136
L136:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v998+v1000<<(uint(int32(2))%32))))
	if base.I32_extend16_s(v994) <= int32(1) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v993 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v994 = v993
	goto L136
L138:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1009].(func(*base.Module, int32, int32))(m, v893, int32(2))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L14
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v893)+20))
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+1)))
	if v1020 != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	goto L140
L142:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1630 = F_tuplestore_gettupleslot(m, v1622, int32(1), int32(0), v893)
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L14
	} else {
		goto L262
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1578
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1579
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_relation_close(m, v1217, int32(0))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L14
	} else {
		goto L261
	}
L144:
	;
	v1560 = int32(_a_F_SequenceSyncWorkerMain_3)
	v1561 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5]))
	v1564 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1564
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1571 = F_lappend_int(m, v933, v1000)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L14
	} else {
		goto L260
	}
L145:
	;
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004)+33)))
	if v1521 != int32(1) {
		v1599 = v918
		v1600 = v919
		v1601 = v920
		v1605 = v928
		v1606 = v929
		v1608 = v933
		v1611 = v938
		v1612 = v939
		v1613 = v940
		v1614 = v941
		v1616 = v943
		goto L142
	} else {
		goto L253
	}
L146:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1022 = *(*int64)(unsafe.Add(mBase, uint32(v1021)+8))
	v1023 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	if v1023 <= int32(2) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1027].(func(*base.Module, int32, int32))(m, v893, int32(3))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L14
	} else {
		goto L150
	}
L148:
	;
	v1038 = v1019
	goto L149
L149:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038)+2)))
	if v1039 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v893)+20))
	v1038 = v1036
	goto L149
L151:
	;
	if v1022 != int64(0) {
		goto L145
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v1063)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1004)+24)) = v1064
	v1066 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	if v1066 <= int32(3) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v1044 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1004)+33)) = uint8(v1044)
	v1046 = int32(_a_F_SequenceSyncWorkerMain_3)
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5]))
	v1050 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1050
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1057 = F_lappend_int(m, v928, v1000)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L14
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1047
	v1599 = v1057
	v1600 = v919
	v1601 = v920
	v1605 = v1057
	v1606 = v929
	v1608 = v933
	v1611 = v938
	v1612 = v939
	v1613 = v940
	v1614 = v941 + int32(1)
	v1616 = v943
	goto L142
L156:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1070].(func(*base.Module, int32, int32))(m, v893, int32(4))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L14
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1080)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v1004)+32)) = uint8(base.B2i32(v1081 != int64(0)))
	v1085 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	if v1085 <= int32(4) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1089].(func(*base.Module, int32, int32))(m, v893, int32(5))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L14
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1099)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1004)+16)) = v1100
	v1102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	if v1102 <= int32(5) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1106].(func(*base.Module, int32, int32))(m, v893, int32(6))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L14
	} else {
		goto L167
	}
L165:
	;
	v1116 = v1102
	goto L166
L166:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+40))
	if base.I32_extend16_s(v1116) <= int32(6) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v1115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v1116 = v1115
	goto L166
L168:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1123].(func(*base.Module, int32, int32))(m, v893, int32(7))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L14
	} else {
		goto L171
	}
L169:
	;
	v1134 = v1116
	v1135 = v1117
	goto L170
L170:
	;
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1135)+48))
	if base.I32_extend16_s(v1134) <= int32(7) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v1134 = v1133
	v1135 = v1132
	goto L170
L172:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1141].(func(*base.Module, int32, int32))(m, v893, int32(8))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L14
	} else {
		goto L175
	}
L173:
	;
	v1152 = v1134
	v1153 = v1135
	goto L174
L174:
	;
	v1154 = *(*int64)(unsafe.Add(mBase, uint32(v1153)+56))
	if base.I32_extend16_s(v1152) <= int32(8) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v1152 = v1151
	v1153 = v1150
	goto L174
L176:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1159].(func(*base.Module, int32, int32))(m, v893, int32(9))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L14
	} else {
		goto L179
	}
L177:
	;
	v1170 = v1152
	v1171 = v1153
	goto L178
L178:
	;
	v1172 = *(*int64)(unsafe.Add(mBase, uint32(v1171)+64))
	if base.I32_extend16_s(v1170) <= int32(9) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v1170 = v1169
	v1171 = v1168
	goto L178
L180:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1177].(func(*base.Module, int32, int32))(m, v893, int32(10))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L14
	} else {
		goto L183
	}
L181:
	;
	v1188 = v1170
	v1189 = v1171
	goto L182
L182:
	;
	v1190 = *(*int64)(unsafe.Add(mBase, uint32(v1189)+72))
	if base.I32_extend16_s(v1188) <= int32(10) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+6)))
	v1188 = v1187
	v1189 = v1186
	goto L182
L184:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	m.T0[v1195].(func(*base.Module, int32, int32))(m, v893, int32(11))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L14
	} else {
		goto L187
	}
L185:
	;
	v1206 = v1189
	goto L186
L186:
	;
	v1207 = *(*int64)(unsafe.Add(mBase, uint32(v1206)+80))
	v1208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1004)+33)) = uint8(v1208)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1217 = F_try_table_open(m, v1210, int32(3))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L14
	} else {
		goto L188
	}
L187:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v893)+16))
	v1206 = v1204
	goto L186
L188:
	;
	if v1217 == int32(0) {
		goto L145
	} else {
		goto L189
	}
L189:
	;
	v1221 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1004)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1228 = F_SearchSysCache1(m, int32(61), v1221)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	if v1228 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L14
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1228)+16))
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1263)+22)))
	v1265 = v1263 + v1264
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+4))
	if v1266 != v1118 {
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54)+160)) = v1241
	F_errmsg_internal(m, int32(_a_F_SequenceSyncWorkerMain_17), v54+int32(160))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L14
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(372), int32(_a_F_SequenceSyncWorkerMain_18))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L14
	} else {
		goto L196
	}
L196:
	;
	goto L11
L197:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+48))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1291 = F_get_namespace_name(m, v1285)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L14
	} else {
		goto L205
	}
L198:
	;
	v1282 = int32(0)
	goto L197
L199:
	;
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v1265)+8))
	if v1268 != v1136 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1270 = *(*int64)(unsafe.Add(mBase, uint32(v1265)+16))
	if v1270 != v1154 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v1272 = *(*int64)(unsafe.Add(mBase, uint32(v1265)+32))
	if v1272 != v1172 {
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v1274 = *(*int64)(unsafe.Add(mBase, uint32(v1265)+24))
	if v1274 != v1190 {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1265)+48)))
	if v1277 == base.B2i32(v1207 != int64(0)) {
		v1282 = int32(1)
		goto L197
	} else {
		goto L204
	}
L204:
	;
	goto L198
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1283))))
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291))))
	if base.B2i32(v1300 == int32(0))|base.B2i32(v1300 != v1303) != 0 {
		v1321 = v1300
		v1322 = v1303
		goto L208
	} else {
		goto L209
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_ReleaseCatCache(m, v1228)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L14
	} else {
		goto L226
	}
L207:
	;
	if v1321-v1322 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	goto L207
L209:
	;
	v1306 = v1283
	v1307 = v1291
	goto L210
L210:
	;
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307)+1)))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306)+1)))
	if v1311 == int32(0) {
		v1321 = v1311
		v1322 = v1310
		goto L208
	} else {
		goto L212
	}
L211:
	;
	v1321 = v1311
	v1322 = v1310
	goto L208
L212:
	;
	v1314 = int32(1)
	if v1311 == v1310 {
		v1306 = v1306 + v1314
		v1307 = v1307 + v1314
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1334 = v1327 + int32(4)
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326))))
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334))))
	if base.B2i32(v1337 == int32(0))|base.B2i32(v1337 != v1340) != 0 {
		v1358 = v1337
		v1359 = v1340
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_ReleaseCatCache(m, v1228)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L14
	} else {
		goto L225
	}
L217:
	;
	if v1358-v1359 == int32(0) {
		goto L206
	} else {
		goto L224
	}
L218:
	;
	goto L217
L219:
	;
	v1343 = v1326
	v1344 = v1334
	goto L220
L220:
	;
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344)+1)))
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343)+1)))
	if v1348 == int32(0) {
		v1358 = v1348
		v1359 = v1347
		goto L218
	} else {
		goto L222
	}
L221:
	;
	v1358 = v1348
	v1359 = v1347
	goto L218
L222:
	;
	v1351 = int32(1)
	if v1348 == v1347 {
		v1343 = v1343 + v1351
		v1344 = v1344 + v1351
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	goto L216
L225:
	;
	goto L144
L226:
	;
	if v1282 == int32(0) {
		goto L144
	} else {
		goto L227
	}
L227:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	v1381 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+39)))
	if v1382 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+48))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_SwitchToUntrustedUser(m, v1386, v54+int32(672))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L14
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1402 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1409 = F_pg_class_aclcheck(m, v1379, v1402, int64(4))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L14
	} else {
		goto L232
	}
L231:
	;
	goto L230
L232:
	;
	if v1409 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v1382 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	v1439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004)+32)))
	v1440 = *(*int64)(unsafe.Add(mBase, uint32(v1004)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_SetSequence(m, v1379, v1440, v1439)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L14
	} else {
		goto L241
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_RestoreUserContext(m, v54+int32(672))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L14
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1422 = int32(_a_F_SequenceSyncWorkerMain_3)
	v1423 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5]))
	v1426 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1426
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1433 = F_lappend_int(m, v929, v1000)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L14
	} else {
		goto L240
	}
L239:
	;
	goto L238
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1423
	v1578 = v1433
	v1579 = v920
	v1582 = v1433
	v1584 = v933
	v1587 = v938
	v1588 = v940 + int32(1)
	v1589 = v943
	goto L143
L241:
	;
	if v1382 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_RestoreUserContext(m, v54+int32(672))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L14
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1459 = *(*int64)(unsafe.Add(mBase, uint32(v1004)+16))
	v1461 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_UpdateSubscriptionRelState(m, v1462, v1379, int32(114), v1459, int32(0))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L14
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1479 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L14
	} else {
		goto L247
	}
L247:
	;
	if v1479 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+24))
	v1484 = *(*int64)(unsafe.Add(mBase, uint32(v1004)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v54)+180)) = base.I64_rotl(v1484, int64(32))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+176)) = v1483
	F_errmsg_internal(m, int32(_a_F_SequenceSyncWorkerMain_19), v54+int32(176))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L14
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1578 = v919
	v1579 = v920
	v1582 = v929
	v1584 = v933
	v1587 = v938
	v1588 = v940
	v1589 = v943 + int32(1)
	goto L143
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(600), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L14
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1531 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L14
	} else {
		goto L254
	}
L254:
	;
	if v1531 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1533 = *(*int64)(unsafe.Add(mBase, uint32(v1004)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v54)+144)) = base.I64_rotl(v1533, int64(32))
	F_errmsg(m, int32(_a_F_SequenceSyncWorkerMain_20), v54+int32(144))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L14
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1599 = v918
	v1600 = v919
	v1601 = v920
	v1605 = v928
	v1606 = v929
	v1608 = v933
	v1611 = v938
	v1612 = v939 + int32(1)
	v1613 = v940
	v1614 = v941
	v1616 = v943
	goto L142
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(654), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L14
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[5])) = v1561
	v1578 = v919
	v1579 = v1571
	v1582 = v929
	v1584 = v1571
	v1587 = v938 + int32(1)
	v1588 = v940
	v1589 = v943
	goto L143
L261:
	;
	v1599 = v918
	v1600 = v1578
	v1601 = v1579
	v1605 = v928
	v1606 = v1582
	v1608 = v1584
	v1611 = v1587
	v1612 = v939
	v1613 = v1588
	v1614 = v941
	v1616 = v1589
	goto L142
L262:
	;
	if v1630 != 0 {
		v918 = v1599
		v919 = v1600
		v920 = v1601
		v928 = v1605
		v929 = v1606
		v933 = v1608
		v938 = v1611
		v939 = v1612
		v940 = v1613
		v941 = v1614
		v943 = v1616
		goto L124
	} else {
		goto L263
	}
L263:
	;
	goto L125
L264:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v842)+8))
	if v1681 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_pfree(m, v1681)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L14
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	if v1689 != 0 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	goto L267
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_tuplestore_end(m, v1689)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L14
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v842)+16))
	if v1697 != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	goto L271
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_FreeTupleDesc(m, v1697)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L14
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_pfree(m, v842)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L14
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1718 = v54 + int32(656)
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1718)))
	v1720 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1719))) = uint8(v1720)
	*(*int32)(unsafe.Add(mBase, uint32(v1718)+12)) = v1720
	*(*int32)(unsafe.Add(mBase, uint32(v1718)+4)) = v1720
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1732 = v54 + int32(640)
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1732)))
	v1734 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1733))) = uint8(v1734)
	*(*int32)(unsafe.Add(mBase, uint32(v1732)+12)) = v1734
	*(*int32)(unsafe.Add(mBase, uint32(v1732)+4)) = v1734
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1747 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L14
	} else {
		goto L280
	}
L280:
	;
	v1752 = v1654 + v1659 + v1655 + v1656 + v1657
	if v1747 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1754)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54+int32(128)))) = v1655
	*(*int32)(unsafe.Add(mBase, uint32(v54+int32(120)))) = v1657
	*(*int32)(unsafe.Add(mBase, uint32(v54+int32(116)))) = v1656
	*(*int32)(unsafe.Add(mBase, uint32(v54+int32(112)))) = v1654
	*(*int32)(unsafe.Add(mBase, uint32(v54)+104)) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v54+int32(124)))) = v809 - v1752
	*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = v1659
	v1770 = base.I32_div_s(v649, int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+100)) = v1770 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+96)) = v1755
	F_errmsg_internal(m, int32(_a_F_SequenceSyncWorkerMain_21), v54+int32(96))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L14
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L14
	} else {
		goto L286
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(680), int32(_a_F_SequenceSyncWorkerMain_10))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L14
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	v1799 = v649 + v809
	if base.B2i32(v1752 == v809)|base.B2i32(v1799 <= v649) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	v1812 = v646
	v1813 = v1805
	v1816 = v649
	v1820 = v654
	goto L290
L288:
	;
	v1877 = v646
	v1885 = v654
	goto L289
L289:
	;
	if v1799 < v527 {
		v642 = v1634
		v643 = v1635
		v644 = v1636
		v646 = v1877
		v649 = v1799
		v652 = v1644
		v653 = v1645
		v654 = v1885
		v657 = v1649
		goto L92
	} else {
		goto L297
	}
L290:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+12))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1848+v1816<<(uint(int32(2))%32))))
	v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852)+33)))
	if v1853 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1877 = v1865
	v1885 = v1867
	goto L289
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1634
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1861 = F_lappend_int(m, v1820, v1816)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L14
	} else {
		goto L295
	}
L293:
	;
	v1865 = v1812
	v1866 = v1813
	v1867 = v1820
	goto L294
L294:
	;
	v1869 = v1816 + int32(1)
	if v1869 != v1799 {
		v1812 = v1865
		v1813 = v1866
		v1816 = v1869
		v1820 = v1867
		goto L290
	} else {
		goto L296
	}
L295:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[7]))
	v1865 = v1861
	v1866 = v1864
	v1867 = v1861
	goto L294
L296:
	;
	goto L291
L297:
	;
	goto L93
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1967 = v54 + int32(684)
	F_initStringInfo(m, v1967)
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L14
	} else {
		goto L299
	}
L299:
	;
	if v1931 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	if v1927 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_get_sequences_string(m, v1931, v1967)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L14
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v1986 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L14
	} else {
		goto L303
	}
L303:
	;
	if v1986 == int32(0) {
		goto L300
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L14
	} else {
		goto L305
	}
L305:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v54)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+80)) = v2004
	F_errmsg_plural(m, int32(_a_F_SequenceSyncWorkerMain_22), int32(_a_F_SequenceSyncWorkerMain_23), v1998, v54+int32(80))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L14
	} else {
		goto L306
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(198), int32(_a_F_SequenceSyncWorkerMain_24))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L14
	} else {
		goto L307
	}
L307:
	;
	goto L300
L308:
	;
	if v1926 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_get_sequences_string(m, v1927, v54+int32(684))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L14
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2041 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L14
	} else {
		goto L311
	}
L311:
	;
	if v2041 == int32(0) {
		goto L308
	} else {
		goto L312
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L14
	} else {
		goto L313
	}
L313:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v54)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v2059
	F_errmsg_plural(m, int32(_a_F_SequenceSyncWorkerMain_25), int32(_a_F_SequenceSyncWorkerMain_26), v2053, v54-int32(-64))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L14
	} else {
		goto L314
	}
L314:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068)+39)))
	if v2069 == int32(1) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errhint_plural(m, int32(_a_F_SequenceSyncWorkerMain_27), int32(_a_F_SequenceSyncWorkerMain_28), v2072, int32(0))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L14
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(223), int32(_a_F_SequenceSyncWorkerMain_24))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L14
	} else {
		goto L319
	}
L318:
	;
	goto L317
L319:
	;
	goto L308
L320:
	;
	if v1928 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_get_sequences_string(m, v1926, v54+int32(684))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L14
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2113 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L14
	} else {
		goto L323
	}
L323:
	;
	if v2113 == int32(0) {
		goto L320
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L14
	} else {
		goto L325
	}
L325:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v54)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+48)) = v2131
	F_errmsg_plural(m, int32(_a_F_SequenceSyncWorkerMain_29), int32(_a_F_SequenceSyncWorkerMain_30), v2125, v54+int32(48))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L14
	} else {
		goto L326
	}
L326:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errhint_plural(m, int32(_a_F_SequenceSyncWorkerMain_31), int32(_a_F_SequenceSyncWorkerMain_32), v2139, int32(0))
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L14
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(239), int32(_a_F_SequenceSyncWorkerMain_24))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L14
	} else {
		goto L328
	}
L328:
	;
	goto L320
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L14
	} else {
		goto L337
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_get_sequences_string(m, v1928, v54+int32(684))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L14
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2179 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L14
	} else {
		goto L332
	}
L332:
	;
	if v2179 == int32(0) {
		goto L329
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L14
	} else {
		goto L334
	}
L334:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v54)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = v2197
	F_errmsg_plural(m, int32(_a_F_SequenceSyncWorkerMain_33), int32(_a_F_SequenceSyncWorkerMain_34), v2191, v54+int32(32))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L14
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(250), int32(_a_F_SequenceSyncWorkerMain_24))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L14
	} else {
		goto L336
	}
L336:
	;
	goto L329
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errcode(m, int32(325))
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L14
	} else {
		goto L338
	}
L338:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2234)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v2235
	F_errmsg(m, int32(_a_F_SequenceSyncWorkerMain_35), v54+int32(16))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L14
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v344
	F_errfinish(m, int32(_a_F_SequenceSyncWorkerMain_6), int32(256), int32(_a_F_SequenceSyncWorkerMain_24))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L14
	} else {
		goto L340
	}
L340:
	;
	goto L11
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L14
	} else {
		goto L342
	}
L342:
	;
	goto L16
L343:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, _c_F_SequenceSyncWorkerMain[3]))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2330)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_pgstat_report_subscription_error(m, v2331)
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L14
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+704)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v54)+700)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v54)+708)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v54)+712)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v54)+716)) = v74
	F_pg_re_throw(m)
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L14
	} else {
		goto L345
	}
L345:
	;
	goto L13
L346:
	;
	v2393 = int32(v2389)
	m.G0 = v54
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v2396)))
	if v54+int32(284) == v2399 {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	m.ExcPending = 1
	goto L1
L348:
	;
	if v2403 != 0 {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+4))
	v2403 = v2401
	goto L351
L350:
	;
	v2403 = int32(0)
	goto L351
L351:
	;
	goto L348
L352:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v54)+716))
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v54)+712))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v54)+708))
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v54)+704))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v54)+700))
	v71 = v2407
	v72 = v2406
	v73 = v2405
	v74 = v2404
	v75 = v2408
	v76 = v2395
	v78 = v2403
	v97 = v126
	v98 = v127
	goto L5
L353:
	;
	goto L354
L354:
	;
	F___wasm_longjmp(m, v2396, v2395)
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
