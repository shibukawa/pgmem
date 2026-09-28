package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateTriggerFiringOn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1500 int32
	_ = v1500
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1562 int64
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1570 int64
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1647 int32
	_ = v1647
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int64
	_ = v1676
	var v1682 int64
	_ = v1682
	var v1689 int64
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1707 int64
	_ = v1707
	var v1709 int64
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1727 int32
	_ = v1727
	var v1736 int32
	_ = v1736
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1765 int32
	_ = v1765
	var v1777 int32
	_ = v1777
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1903 int32
	_ = v1903
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1949 int32
	_ = v1949
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1972 int32
	_ = v1972
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2010 int32
	_ = v2010
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2052 int32
	_ = v2052
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2089 int32
	_ = v2089
	var v2097 int32
	_ = v2097
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2191 int32
	_ = v2191
	var v2206 int32
	_ = v2206
	var v2210 int64
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2229 int32
	_ = v2229
	var v2239 int32
	_ = v2239
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2330 int32
	_ = v2330
	var v2339 int32
	_ = v2339
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2375 int32
	_ = v2375
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2389 int32
	_ = v2389
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2448 int64
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2456 int64
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2504 int64
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2562 int32
	_ = v2562
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2578 int32
	_ = v2578
	var v2588 int32
	_ = v2588
	var v2600 int32
	_ = v2600
	var v2603 int32
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2614 int32
	_ = v2614
	var v2621 int32
	_ = v2621
	var v2629 int32
	_ = v2629
	var v2661 int32
	_ = v2661
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2751 int32
	_ = v2751
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2844 int32
	_ = v2844
	var v2879 int32
	_ = v2879
	var v2886 int32
	_ = v2886
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2899 int32
	_ = v2899
	var v2904 int32
	_ = v2904
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2915 int32
	_ = v2915
	var v2920 int32
	_ = v2920
	v14 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(608)
	m.G0 = v35
	if l3 != 0 {
		goto L28
	} else {
		goto L29
	}
L1:
	;
	if l9 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L31
	} else {
		goto L268
	}
L3:
	;
	v1082 = int32(0)
	if base.B2i32(v1076 == v1082)|base.B2i32(v1074 == v1082) != 0 {
		v1205 = v1074
		v1207 = v1076
		goto L1
	} else {
		goto L255
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L31
	} else {
		goto L250
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L31
	} else {
		goto L245
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L31
	} else {
		goto L241
	}
L7:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+131)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L31
	} else {
		goto L232
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L31
	} else {
		goto L228
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L31
	} else {
		goto L224
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L31
	} else {
		goto L220
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L31
	} else {
		goto L216
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L31
	} else {
		goto L212
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L31
	} else {
		goto L208
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L31
	} else {
		goto L204
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L31
	} else {
		goto L200
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L31
	} else {
		goto L196
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L31
	} else {
		goto L192
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L31
	} else {
		goto L188
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L31
	} else {
		goto L184
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L31
	} else {
		goto L179
	}
L21:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[0])))
	if v255 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L31
	} else {
		goto L77
	}
L23:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v177 {
	case 0, 2:
		goto L64
	default:
		goto L65
	}
L24:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	if v141 != int32(64) {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v84 {
	case 0, 2:
		goto L40
	default:
		goto L41
	}
L26:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v58 {
	case 0, 2:
		goto L21
	default:
		goto L34
	}
L27:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
	v48 = v46 - int32(102)
	switch (v48<<(uint(int32(7))%32) | int32(base.Ui32(v48&int32(254))>>(uint(int32(1))%32))) & int32(255) {
	case 0:
		goto L23
	default:
		goto L22
	case 5:
		goto L25
	case 6:
		goto L26
	case 8:
		goto L24
	}
L28:
	;
	v38 = F_table_open(m, l3, int32(6))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = F_table_openrv(m, v40, int32(6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L31
	} else {
		goto L33
	}
L31:
	;
	return
L32:
	;
	v44 = v38
	goto L27
L33:
	;
	v44 = v42
	goto L27
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v66 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_0), v35+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v77 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(231), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v110 != int32(1) {
		goto L21
	} else {
		goto L47
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v92 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_0), v35+int32(240))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	v103 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_1), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(242), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v113 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L31
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+256)) = v123 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_4), v35+int32(256))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v134 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_5), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(265), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v144 == int32(1) {
		goto L20
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v147&int32(32) == int32(0) {
		goto L21
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L31
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v159 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v35+int32(272))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L31
	} else {
		goto L61
	}
L61:
	;
	v170 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_7), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L31
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(286), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L31
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v203 != int32(1) {
		goto L21
	} else {
		goto L71
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+304)) = v185 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v35+int32(304))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	v196 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_9), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(296), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L31
	} else {
		goto L73
	}
L73:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+320)) = v213 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v35+int32(320))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L31
	} else {
		goto L74
	}
L74:
	;
	v224 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_10), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L31
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(308), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L31
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v238 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_11), v35)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L31
	} else {
		goto L79
	}
L79:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v246 = int32(*(*int8)(unsafe.Add(mBase, uint32(v245)+119)))
	F_errdetail_relkind_not_supported(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(315), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L31
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v259 = int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	if base.Ui32(v260) < base.Ui32(int32(_a_F_CreateTriggerFiringOn_12)) {
		v269 = v259
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L84
L84:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v270 != int32(1) {
		v286 = v14
		goto L90
	} else {
		goto L91
	}
L85:
	;
	if v269 != 0 {
		goto L19
	} else {
		goto L89
	}
L86:
	;
	goto L85
L87:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+68))
	if v264 == int32(99) {
		v269 = v259
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v267 = F_isTempToastNamespace(m, v264)
	mBase = m.M
	v269 = v267
	goto L86
L89:
	;
	goto L84
L90:
	;
	if l10 != 0 {
		v356 = v14
		goto L98
	} else {
		goto L99
	}
L91:
	;
	if l4 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_LockRelationOid(m, l4, int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L31
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v276 == int32(0) {
		v286 = v14
		goto L90
	} else {
		goto L96
	}
L95:
	;
	v286 = l4
	goto L90
L96:
	;
	v280 = int32(0)
	v283 = F_RangeVarGetRelidExtended(m, v276, int32(1), v280, v280, v280)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L31
	} else {
		goto L97
	}
L97:
	;
	v286 = v283
	goto L90
L98:
	;
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v361 = v357 | (v358 | v359)
	v362 = int32(33)
	if v361&v362 == v362 {
		goto L18
	} else {
		goto L130
	}
L99:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v291 = F_pg_class_aclcheck(m, v287, v289, int64(64))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L31
	} else {
		goto L100
	}
L100:
	;
	if v291 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v294 = int32(*(*int8)(unsafe.Add(mBase, uint32(v293)+119)))
	switch v294 - int32(73) {
	case 0, 32:
		v304 = int32(20)
		goto L105
	default:
		goto L106
	case 10:
		goto L110
	case 29:
		goto L107
	case 36:
		goto L108
	case 45:
		goto L109
	}
L102:
	;
	goto L103
L103:
	;
	if v286 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L104:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	F_aclcheck_error(m, v291, v306, v307+int32(4))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L31
	} else {
		goto L111
	}
L105:
	;
	v306 = v304
	goto L104
L106:
	;
	v304 = int32(42)
	goto L105
L107:
	;
	v306 = int32(18)
	goto L104
L108:
	;
	v306 = int32(23)
	goto L104
L109:
	;
	v306 = int32(52)
	goto L104
L110:
	;
	v306 = int32(38)
	goto L104
L111:
	;
	goto L103
L112:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v340 != int32(1) {
		v356 = v14
		goto L98
	} else {
		goto L126
	}
L113:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v317 = F_pg_class_aclcheck(m, v286, v315, int64(64))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L31
	} else {
		goto L114
	}
L114:
	;
	if v317 == int32(0) {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v321 = F_get_rel_relkind(m, v286)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L31
	} else {
		goto L116
	}
L116:
	;
	switch v321 - int32(73) {
	case 0, 32:
		v332 = int32(20)
		goto L118
	default:
		goto L119
	case 10:
		goto L123
	case 29:
		goto L120
	case 36:
		goto L121
	case 45:
		goto L122
	}
L117:
	;
	v335 = F_get_rel_name(m, v286)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L31
	} else {
		goto L124
	}
L118:
	;
	v334 = v332
	goto L117
L119:
	;
	v332 = int32(42)
	goto L118
L120:
	;
	v334 = int32(18)
	goto L117
L121:
	;
	v334 = int32(23)
	goto L117
L122:
	;
	v334 = int32(52)
	goto L117
L123:
	;
	v334 = int32(38)
	goto L117
L124:
	;
	F_aclcheck_error(m, v317, v334, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L31
	} else {
		goto L125
	}
L125:
	;
	goto L112
L126:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+119)))
	if v344 != int32(112) {
		v356 = v14
		goto L98
	} else {
		goto L127
	}
L127:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v350 = F_find_all_inheritors(m, v347, int32(6), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L31
	} else {
		goto L128
	}
L128:
	;
	F_list_free(m, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L31
	} else {
		goto L129
	}
L129:
	;
	v356 = int32(1)
	goto L98
L130:
	;
	v367 = v361 & int32(1)
	v369 = v361 & int32(66)
	if v369 == int32(64) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v367 == int32(0) {
		goto L17
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v376 == int32(0) {
		v1205 = v14
		v1207 = v14
		goto L1
	} else {
		goto L137
	}
L134:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v374 != 0 {
		goto L16
	} else {
		goto L135
	}
L135:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v375 != 0 {
		goto L15
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v379 <= int32(0) {
		v1205 = v14
		v1207 = v14
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+9)))
	if v384 != int32(1) {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+119)))
	v390 = v388 - int32(102)
	if v390 == int32(0) {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	if v390 == int32(16) {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	if v367 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v396 = F_has_superclass(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L31
	} else {
		goto L145
	}
L143:
	;
	v399 = v357
	goto L144
L144:
	;
	if v399&int32(_a_F_CreateTriggerFiringOn_13) != 0 {
		goto L6
	} else {
		goto L147
	}
L145:
	;
	if v396 != 0 {
		goto L7
	} else {
		goto L146
	}
L146:
	;
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v399 = v398
	goto L144
L147:
	;
	if v361&int32(32) != 0 {
		goto L8
	} else {
		goto L148
	}
L148:
	;
	v406 = int32(1)
	if int32(base.Ui32(v361)>>(uint(int32(2))%32))&v406+int32(base.Ui32(v361)>>(uint(int32(4))%32))&v406+int32(base.Ui32(v361)>>(uint(int32(3))%32))&v406 != v406 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v420 != 0 {
		goto L14
	} else {
		goto L150
	}
L150:
	;
	v422 = v361 & int32(20)
	v424 = v361 & int32(24)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+8)))
	if v425 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v437 < int32(2) {
		v1074 = v436
		v1076 = v435
		goto L3
	} else {
		goto L157
	}
L152:
	;
	if v424 == int32(0) {
		goto L11
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v422 == int32(0) {
		goto L13
	} else {
		goto L156
	}
L155:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	v435 = v14
	v436 = v430
	goto L151
L156:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	v435 = v433
	v436 = int32(0)
	goto L151
L157:
	;
	v444 = int32(1)
	v465 = v436
	v467 = v435
	goto L158
L158:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473+v444<<(uint(int32(2))%32))))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+9)))
	if v478 == int32(0) {
		goto L2
	} else {
		goto L160
	}
L159:
	;
	v1074 = v503
	v1076 = v504
	goto L3
L160:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+119)))
	v484 = v482 - int32(102)
	if v484 == int32(0) {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v484 == int32(16) {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	if v367 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v490 = F_has_superclass(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L31
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v493 != 0 {
		goto L14
	} else {
		goto L169
	}
L166:
	;
	if v490 != 0 {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	if v492 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+8)))
	if v494 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v506 = v444 + int32(1)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v506 < v507 {
		v444 = v506
		v465 = v503
		v467 = v504
		goto L158
	} else {
		goto L178
	}
L171:
	;
	if v422 == int32(0) {
		goto L13
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v424 == int32(0) {
		goto L11
	} else {
		goto L176
	}
L174:
	;
	if v467 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v503 = v465
	v504 = v499
	goto L170
L176:
	;
	if v465 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v503 = v502
	v504 = v467
	goto L170
L178:
	;
	goto L159
L179:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L31
	} else {
		goto L180
	}
L180:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+288)) = v516 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v35+int32(288))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L31
	} else {
		goto L181
	}
L181:
	;
	v527 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_14), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L31
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(279), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L31
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L31
	} else {
		goto L185
	}
L185:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+224)) = v541 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_15), v35+int32(224))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L31
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(321), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L31
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L31
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_16), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L31
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(384), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L31
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L31
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_17), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L31
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(392), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L31
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L31
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_18), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L31
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(396), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L31
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L31
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_19), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L31
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(400), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L31
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L31
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_20), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L31
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(509), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L31
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L31
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L31
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(526), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L31
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L31
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_22), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L31
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(531), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L31
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L31
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_23), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L31
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(541), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L31
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L31
	} else {
		goto L221
	}
L221:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_24), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L31
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(546), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L31
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L31
	} else {
		goto L225
	}
L225:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_25), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L31
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(498), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L31
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L31
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_26), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L31
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(481), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L31
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L31
	} else {
		goto L233
	}
L233:
	;
	if v860 != int32(1) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_27), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L31
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_28), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L31
	} else {
		goto L239
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(470), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L31
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(466), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L31
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L31
	} else {
		goto L242
	}
L242:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_29), int32(0))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L31
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(476), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L31
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L31
	} else {
		goto L246
	}
L246:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+208)) = v975 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v35+int32(208))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L31
	} else {
		goto L247
	}
L247:
	;
	v986 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_30), int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L31
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(450), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L31
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L31
	} else {
		goto L251
	}
L251:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+192)) = v1032 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v35+int32(192))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L31
	} else {
		goto L252
	}
L252:
	;
	v1043 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_31), int32(0))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L31
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(443), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L31
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076))))
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1074))))
	if base.B2i32(v1089 == int32(0))|base.B2i32(v1089 != v1092) != 0 {
		v1110 = v1089
		v1111 = v1092
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1110-v1111 != 0 {
		v1205 = v1074
		v1207 = v1076
		goto L1
	} else {
		goto L263
	}
L257:
	;
	goto L256
L258:
	;
	v1095 = v1076
	v1096 = v1074
	goto L259
L259:
	;
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+1)))
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1095)+1)))
	if v1100 == int32(0) {
		v1110 = v1100
		v1111 = v1099
		goto L257
	} else {
		goto L261
	}
L260:
	;
	v1110 = v1100
	v1111 = v1099
	goto L257
L261:
	;
	v1103 = int32(1)
	if v1100 == v1099 {
		v1095 = v1095 + v1103
		v1096 = v1096 + v1103
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L31
	} else {
		goto L264
	}
L264:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L31
	} else {
		goto L265
	}
L265:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_32), int32(0))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L31
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(556), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L31
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L31
	} else {
		goto L269
	}
L269:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_33), int32(0))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L31
	} else {
		goto L270
	}
L270:
	;
	F_errhint(m, int32(_a_F_CreateTriggerFiringOn_34), int32(0))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L31
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(430), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L31
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v2206 = int32(0)
	v2210 = F_DirectFunctionCall1Coll(m, int32(624), v2206, base.I64_extend_i32_u(v2191))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L31
	} else {
		goto L463
	}
L274:
	;
	v2032 = F_palloc(m, v2010)
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L31
	} else {
		goto L449
	}
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L31
	} else {
		goto L445
	}
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L31
	} else {
		goto L441
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L31
	} else {
		goto L437
	}
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L31
	} else {
		goto L433
	}
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L31
	} else {
		goto L428
	}
L280:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L31
	} else {
		goto L423
	}
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L31
	} else {
		goto L418
	}
L282:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L31
	} else {
		goto L413
	}
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L31
	} else {
		goto L408
	}
L284:
	;
	if l7 == int32(0) {
		goto L354
	} else {
		goto L355
	}
L285:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v1215 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	goto L287
L287:
	;
	v1489 = F_nodeToString(m, l9)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L31
	} else {
		goto L353
	}
L288:
	;
	v1500 = int32(0)
	v1514 = v14
	v1519 = v14
	goto L284
L289:
	;
	goto L290
L290:
	;
	v1220 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L31
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+4)) = l2
	v1226 = F_makeAlias(m, int32(_a_F_CreateTriggerFiringOn_35), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L31
	} else {
		goto L292
	}
L292:
	;
	v1228 = int32(0)
	v1230 = F_addRangeTableEntryForRelation(m, v1220, v44, int32(1), v1226, v1228, v1228)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L31
	} else {
		goto L293
	}
L293:
	;
	v1233 = int32(1)
	F_addNSItemToQuery(m, v1220, v1230, int32(0), v1233, v1233)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L31
	} else {
		goto L294
	}
L294:
	;
	v1240 = F_makeAlias(m, int32(_a_F_CreateTriggerFiringOn_36), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L31
	} else {
		goto L295
	}
L295:
	;
	v1242 = int32(0)
	v1244 = F_addRangeTableEntryForRelation(m, v1220, v44, int32(1), v1240, v1242, v1242)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L31
	} else {
		goto L296
	}
L296:
	;
	v1247 = int32(1)
	F_addNSItemToQuery(m, v1220, v1244, int32(0), v1247, v1247)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L31
	} else {
		goto L297
	}
L297:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v1252 = F_copyObjectImpl(m, v1251)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L31
	} else {
		goto L298
	}
L298:
	;
	v1256 = F_transformWhereClause(m, v1220, v1252, int32(37), int32(_a_F_CreateTriggerFiringOn_37))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L31
	} else {
		goto L299
	}
L299:
	;
	F_assign_expr_collations(m, v1220, v1256)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L31
	} else {
		goto L300
	}
L300:
	;
	v1261 = F_pull_var_clause(m, v1256, int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L31
	} else {
		goto L302
	}
L301:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+8))
	v1485 = F_nodeToString(m, v1256)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L31
	} else {
		goto L351
	}
L302:
	;
	if v1261 == int32(0) {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+4))
	if v1265 <= int32(0) {
		goto L301
	} else {
		goto L304
	}
L304:
	;
	v1268 = int32(0)
	if v1268 < v1265 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1271 = v1265
	goto L307
L306:
	;
	v1271 = v1268
	goto L307
L307:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+12))
	v1281 = int32(0)
	goto L308
L308:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1276+v1281<<(uint(int32(2))%32))))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	switch v1314 - int32(1) {
	case 0:
		goto L313
	case 1:
		goto L312
	default:
		goto L311
	}
L309:
	;
	goto L301
L310:
	;
	v1450 = v1281 + int32(1)
	if v1450 != v1271 {
		v1281 = v1450
		goto L308
	} else {
		goto L350
	}
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L31
	} else {
		goto L347
	}
L312:
	;
	if v367 == int32(0) {
		goto L282
	} else {
		goto L321
	}
L313:
	;
	if v367 == int32(0) {
		goto L283
	} else {
		goto L314
	}
L314:
	;
	if v361&int32(4) == int32(0) {
		goto L310
	} else {
		goto L315
	}
L315:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L31
	} else {
		goto L316
	}
L316:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L31
	} else {
		goto L317
	}
L317:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_38), int32(0))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L31
	} else {
		goto L318
	}
L318:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1332)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L31
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(626), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L31
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	if v361&int32(8) != 0 {
		goto L281
	} else {
		goto L322
	}
L322:
	;
	v1343 = base.B2i32(v369 != int32(2))
	v1344 = int32(0)
	v1346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1313)+8)))
	if base.B2i32(v1343 == v1344)&base.B2i32(v1346 < v1344) != 0 {
		goto L280
	} else {
		goto L323
	}
L323:
	;
	if v369 != int32(2) {
		goto L310
	} else {
		goto L324
	}
L324:
	;
	if v1346 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+24))
	if v1353 == int32(0) {
		goto L310
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	if v1346 <= int32(0) {
		goto L310
	} else {
		goto L339
	}
L328:
	;
	v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353)+17)))
	if v1356 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353)+18)))
	if v1359 != int32(1) {
		goto L310
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L31
	} else {
		goto L333
	}
L332:
	;
	goto L331
L333:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L31
	} else {
		goto L334
	}
L334:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_39), int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L31
	} else {
		goto L335
	}
L335:
	;
	v1375 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_40), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L31
	} else {
		goto L336
	}
L336:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1377)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L31
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(654), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L31
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387+v1388<<(uint(int32(3))%32)+v1346*int32(100))+18)))
	if v1395 == int32(0) {
		goto L310
	} else {
		goto L340
	}
L340:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L31
	} else {
		goto L341
	}
L341:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L31
	} else {
		goto L342
	}
L342:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_39), int32(0))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L31
	} else {
		goto L343
	}
L343:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	v1414 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1313)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+176)) = v1409 + v1410<<(uint(int32(3))%32) + v1414*int32(100) - int32(68)
	v1424 = F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_41), v35+int32(176))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L31
	} else {
		goto L344
	}
L344:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1426)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L31
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(663), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L31
	} else {
		goto L346
	}
L346:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L347:
	;
	F_errmsg_internal(m, int32(_a_F_CreateTriggerFiringOn_42), int32(0))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L31
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(667), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L31
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	goto L309
L351:
	;
	F_free_parsestate(m, v1220)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L31
	} else {
		goto L352
	}
L352:
	;
	v1500 = v1256
	v1514 = v1484
	v1519 = v1485
	goto L284
L353:
	;
	v1500 = l9
	v1514 = v14
	v1519 = v1489
	goto L284
L354:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1526 = int32(0)
	v1529 = F_LookupFuncName(m, v1525, v1526, v1526, v1526)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L31
	} else {
		goto L357
	}
L355:
	;
	v1531 = l7
	goto L356
L356:
	;
	if l10 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1531 = v1529
	goto L356
L358:
	;
	v1547 = F_get_func_rettype(m, v1531)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L31
	} else {
		goto L364
	}
L359:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v1536 = F_object_aclcheck(m, int32(1255), v1531, v1534, int64(128))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L31
	} else {
		goto L360
	}
L360:
	;
	if v1536 == int32(0) {
		goto L358
	} else {
		goto L361
	}
L361:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1542 = F_NameListToString(m, v1541)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L31
	} else {
		goto L362
	}
L362:
	;
	F_aclcheck_error(m, v1536, int32(19), v1542)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L31
	} else {
		goto L363
	}
L363:
	;
	goto L358
L364:
	;
	if v1547 != int32(2279) {
		goto L279
	} else {
		goto L365
	}
L365:
	;
	v1553 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L31
	} else {
		goto L366
	}
L366:
	;
	if l10 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	if l5 != 0 {
		v1661 = l5
		goto L384
	} else {
		goto L385
	}
L368:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+16))
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1591)+22)))
	v1593 = v1591 + v1592
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+8))
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593)+83)))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+92))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1593)))
	v1598 = F_heap_copytuple(m, v1579)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L31
	} else {
		goto L379
	}
L369:
	;
	v1558 = v35 + int32(448)
	v1562 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+56)))
	F_ScanKeyInit(m, v1558, int32(2), int32(3), int32(184), v1562)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L31
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	v1588 = F_GetNewOidWithIndex(m, v1553, int32(2702), int32(1))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L31
	} else {
		goto L378
	}
L372:
	;
	v1570 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)))
	F_ScanKeyInit(m, v35+int32(504), int32(4), int32(3), int32(62), v1570)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L31
	} else {
		goto L373
	}
L373:
	;
	v1577 = F_systable_beginscan(m, v1553, int32(2701), int32(1), int32(0), int32(2), v1558)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L31
	} else {
		goto L374
	}
L374:
	;
	v1579 = F_systable_getnext(m, v1577)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L31
	} else {
		goto L375
	}
L375:
	;
	if v1579 != 0 {
		goto L368
	} else {
		goto L376
	}
L376:
	;
	F_systable_endscan(m, v1577)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L31
	} else {
		goto L377
	}
L377:
	;
	goto L371
L378:
	;
	v1616 = v1588
	v1617 = int32(0)
	v1620 = v14
	goto L367
L379:
	;
	F_systable_endscan(m, v1577)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L31
	} else {
		goto L380
	}
L380:
	;
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v1602 == int32(0) {
		goto L278
	} else {
		goto L381
	}
L381:
	;
	v1605 = int32(0)
	if base.B2i32(v1594 == v1605)&(v1595^int32(-1))|l11 == v1605 {
		goto L277
	} else {
		goto L382
	}
L382:
	;
	if v1596 != 0 {
		goto L276
	} else {
		goto L383
	}
L383:
	;
	v1616 = v1597
	v1617 = int32(1)
	v1620 = v1598
	goto L367
L384:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if l10 != 0 {
		goto L388
	} else {
		goto L389
	}
L385:
	;
	v1621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v1621&int32(1) == int32(0) {
		v1661 = l5
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+68))
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	v1632 = int32(1)
	v1634 = int32(0)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v1647 = int32(32)
	v1659 = F_CreateConstraintEntry(m, v1626, v1628, int32(116), v1630, v1631, v1632, v1632, v1634, v1635, v1634, v1634, v1634, v1634, v1634, v1634, v1634, v1634, v1634, v1634, v1634, v1647, v1647, v1634, v1634, v1647, v1634, v1634, v1634, v1632, v1634, v1632, v1634, l10)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L31
	} else {
		goto L387
	}
L387:
	;
	v1661 = v1659
	goto L384
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+100)) = v1616
	*(*int32)(unsafe.Add(mBase, uint32(v35)+96)) = v1662
	v1666 = v35 + int32(352)
	v1671 = F_pg_snprintf(m, v1666, int32(64), int32(_a_F_CreateTriggerFiringOn_43), v35+int32(96))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L31
	} else {
		goto L391
	}
L389:
	;
	v1673 = v1662
	goto L390
L390:
	;
	v1674 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+431)) = v1674
	v1676 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+424)) = v1676
	*(*int64)(unsafe.Add(mBase, uint32(v35)+416)) = v1676
	*(*int64)(unsafe.Add(mBase, uint32(v35)+448)) = base.I64_extend_i32_u(v1616)
	v1682 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+56)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+464)) = base.I64_extend_i32_u(l8)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+456)) = v1682
	v1689 = F_DirectFunctionCall1Coll(m, int32(534), v1674, base.I64_extend_i32_u(v1673))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L31
	} else {
		goto L392
	}
L391:
	;
	v1673 = v1666
	goto L390
L392:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+528)) = base.I64_extend_i32_u(v1661)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+520)) = base.I64_extend_i32_u(l6)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+512)) = base.I64_extend_i32_u(v286)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+504)) = base.I64_extend_i32_u(l10)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+496)) = base.I64_extend_i32_s(l12)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+488)) = base.I64_extend16_s(base.I64_extend_i32_u(v361))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+480)) = base.I64_extend_i32_u(v1531)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+472)) = v1689
	v1707 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+536)) = v1707
	v1709 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+544)) = v1709
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1711 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+4))
	if int32(_a_F_CreateTriggerFiringOn_44) < v1712 {
		goto L275
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+552)) = int64(0)
	v2191 = int32(_a_F_CreateTriggerFiringOn_45)
	goto L273
L396:
	;
	if v1712 <= int32(0) {
		v2010 = int32(1)
		goto L274
	} else {
		goto L397
	}
L397:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+12))
	v1719 = int32(0)
	v1727 = v1719
	v1736 = v1719
	goto L398
L398:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1718+v1727<<(uint(int32(2))%32))))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	v1758 = F_strlen(m, v1757)
	mBase = m.M
	v1765 = v1757
	v1777 = v1758 + v1736 + int32(4)
	goto L400
L400:
	;
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1765))))
	if v1794 != int32(92) {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v1765 = v1765 + int32(1)
	v1777 = v1804
	goto L400
L403:
	;
	if v1794 != 0 {
		v1804 = v1777
		goto L402
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1804 = v1777 + int32(1)
	goto L402
L406:
	;
	v1798 = v1727 + int32(1)
	if v1798 != v1712 {
		v1727 = v1798
		v1736 = v1777
		goto L398
	} else {
		goto L407
	}
L407:
	;
	v2010 = v1777 + int32(1)
	goto L274
L408:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L31
	} else {
		goto L409
	}
L409:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_46), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L31
	} else {
		goto L410
	}
L410:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1821)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L31
	} else {
		goto L411
	}
L411:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(621), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L31
	} else {
		goto L412
	}
L412:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L413:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L31
	} else {
		goto L414
	}
L414:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_46), int32(0))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L31
	} else {
		goto L415
	}
L415:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1840)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L31
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(634), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L31
	} else {
		goto L417
	}
L417:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L418:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L31
	} else {
		goto L419
	}
L419:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_47), int32(0))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L31
	} else {
		goto L420
	}
L420:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1859)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L31
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(639), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L31
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L31
	} else {
		goto L424
	}
L424:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_48), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L31
	} else {
		goto L425
	}
L425:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	F_parser_errposition(m, v1220, v1878)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L31
	} else {
		goto L426
	}
L426:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(644), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L31
	} else {
		goto L427
	}
L427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L428:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L31
	} else {
		goto L429
	}
L429:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1894 = F_NameListToString(m, v1893)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L31
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+164)) = int32(_a_F_CreateTriggerFiringOn_49)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+160)) = v1894
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_50), v35+int32(160))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L31
	} else {
		goto L431
	}
L431:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(708), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L31
	} else {
		goto L432
	}
L432:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L433:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_51))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L31
	} else {
		goto L434
	}
L434:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+144)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v35)+148)) = v1916 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_52), v35+int32(144))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L31
	} else {
		goto L435
	}
L435:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(769), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L31
	} else {
		goto L436
	}
L436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L437:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_51))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L31
	} else {
		goto L438
	}
L438:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+128)) = v1940
	*(*int32)(unsafe.Add(mBase, uint32(v35)+132)) = v1939 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_53), v35+int32(128))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L31
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(782), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L31
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_51))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L31
	} else {
		goto L442
	}
L442:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v1963
	*(*int32)(unsafe.Add(mBase, uint32(v35)+116)) = v1962 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_54), v35+int32(112))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L31
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(801), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L31
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L31
	} else {
		goto L446
	}
L446:
	;
	v1985 = int32(_a_F_CreateTriggerFiringOn_44)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+80)) = v1985
	F_errmsg_plural(m, int32(_a_F_CreateTriggerFiringOn_55), int32(_a_F_CreateTriggerFiringOn_56), v1985, v35+int32(80))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L31
	} else {
		goto L447
	}
L447:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(898), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L31
	} else {
		goto L448
	}
L448:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L449:
	;
	v2034 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2032))) = uint8(v2034)
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v2036 == v2034 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+552)) = base.I64_extend16_s(base.I64_extend_i32_u(v1712))
	v2191 = v2032
	goto L273
L451:
	;
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v2036)+4))
	if v2039 <= int32(0) {
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v2043 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[2])))
	v2045 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[3]))
	v2052 = int32(0)
	goto L453
L453:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2036)+12))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2078+v2052<<(uint(int32(2))%32))))
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+4))
	v2084 = F_strlen(m, v2032)
	mBase = m.M
	v2089 = v2084 + v2032
	v2097 = v2083
	goto L455
L455:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2097))))
	if v2118 != int32(92) {
		goto L458
	} else {
		goto L459
	}
L457:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2132))) = uint8(v2133)
	v2135 = int32(1)
	v2089 = v2132 + v2135
	v2097 = v2097 + v2135
	goto L455
L458:
	;
	if v2118 != 0 {
		v2132 = v2089
		v2133 = v2118
		goto L457
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v2127 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v2089))) = uint8(v2127)
	v2131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2097))))
	v2132 = v2089 + int32(1)
	v2133 = v2131
	goto L457
L461:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2089)+4)) = uint8(v2043)
	*(*int32)(unsafe.Add(mBase, uint32(v2089))) = v2045
	v2124 = v2052 + int32(1)
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2036)+4))
	if v2125 <= v2124 {
		goto L450
	} else {
		goto L462
	}
L462:
	;
	v2052 = v2124
	goto L453
L463:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+568)) = v2210
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v2213 == int32(0) {
		goto L467
	} else {
		goto L468
	}
L464:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L31
	} else {
		goto L621
	}
L465:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2886 = m.ExcPending
	if v2886 != 0 {
		goto L31
	} else {
		goto L617
	}
L466:
	;
	v2435 = F_buildint2vector(m, v2414, v2418)
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L31
	} else {
		goto L506
	}
L467:
	;
	v2414 = int32(0)
	v2418 = v2206
	goto L466
L468:
	;
	goto L469
L469:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2213)+4))
	if v2217 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2220 = int32(0)
	v2414 = v2220
	v2418 = v2220
	goto L466
L471:
	;
	goto L472
L472:
	;
	v2224 = F_palloc(m, v2217<<(uint(int32(1))%32))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L31
	} else {
		goto L473
	}
L473:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v2226 == int32(0) {
		v2414 = v2224
		v2418 = v2217
		goto L466
	} else {
		goto L474
	}
L474:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+4))
	if v2229 <= int32(0) {
		v2414 = v2224
		v2418 = v2217
		goto L466
	} else {
		goto L475
	}
L475:
	;
	v2239 = int32(0)
	goto L476
L476:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+12))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2265+v2239<<(uint(int32(2))%32))))
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	v2271 = int32(0)
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v2275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2274)+120)))
	if v2271 < v2275 {
		goto L481
	} else {
		goto L482
	}
L477:
	;
	v2414 = v2224
	v2418 = v2217
	goto L466
L478:
	;
	if v2330<<(uint(int32(16))%32) == int32(0) {
		goto L465
	} else {
		goto L495
	}
L479:
	;
	v2330 = v2281 + int32(1)
	goto L478
L480:
	;
	goto L479
L481:
	;
	v2281 = v2271
	goto L484
L482:
	;
	goto L483
L483:
	;
	goto L491
L484:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2283)))
	v2290 = v2283 + v2284<<(uint(int32(3))%32) + v2281*int32(100)
	v2293 = F_namestrcmp(m, v2290+int32(32), v2270)
	mBase = m.M
	if v2293 == int32(0) {
		goto L486
	} else {
		goto L487
	}
L485:
	;
	goto L483
L486:
	;
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290)+119)))
	if v2296 != int32(1) {
		goto L480
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	v2300 = v2281 + int32(1)
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v2302 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2301)+120)))
	if v2300 < v2302 {
		v2281 = v2300
		goto L484
	} else {
		goto L490
	}
L489:
	;
	goto L488
L490:
	;
	goto L485
L491:
	;
	v2330 = int32(0)
	goto L478
L495:
	;
	v2339 = v2239
	goto L497
L496:
	;
	v2395 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2224+v2239<<(uint(v2395)%32)))) = uint16(v2330)
	v2400 = v2239 + v2395
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+4))
	if v2400 < v2401 {
		v2239 = v2400
		goto L476
	} else {
		goto L505
	}
L497:
	;
	if v2339 <= int32(0) {
		goto L496
	} else {
		goto L499
	}
L498:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L31
	} else {
		goto L501
	}
L499:
	;
	v2370 = int32(1)
	v2371 = v2339 - v2370
	v2375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2224+v2371<<(uint(v2370)%32)))))
	if base.I32_extend16_s(v2330) != v2375 {
		v2339 = v2371
		goto L497
	} else {
		goto L500
	}
L500:
	;
	goto L498
L501:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L31
	} else {
		goto L502
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+64)) = v2270
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_57), v35-int32(-64))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L31
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(968), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L31
	} else {
		goto L504
	}
L504:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L505:
	;
	goto L477
L506:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+560)) = base.I64_extend_i32_u(v2435)
	if v1519 != 0 {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	if v1205 != 0 {
		goto L513
	} else {
		goto L514
	}
L508:
	;
	v2439 = F_cstring_to_text(m, v1519)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L31
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	v2443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+432)) = uint8(v2443)
	goto L507
L511:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+576)) = base.I64_extend_i32_u(v2439)
	goto L507
L512:
	;
	if v1207 != 0 {
		goto L518
	} else {
		goto L519
	}
L513:
	;
	v2448 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v1205))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L31
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v2451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+433)) = uint8(v2451)
	goto L512
L516:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+584)) = v2448
	goto L512
L517:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+52))
	v2466 = F_heap_form_tuple(m, v2461, v35+int32(448), v35+int32(416))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L31
	} else {
		goto L522
	}
L518:
	;
	v2456 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v1207))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L31
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v2459 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+434)) = uint8(v2459)
	goto L517
L521:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+592)) = v2456
	goto L517
L522:
	;
	if v1617 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L523:
	;
	F_pfree(m, v2478)
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L31
	} else {
		goto L530
	}
L524:
	;
	F_CatalogTupleInsert(m, v1553, v2466)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L31
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	F_CatalogTupleUpdate(m, v1553, v1620+int32(4), v2466)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L31
	} else {
		goto L528
	}
L527:
	;
	v2478 = v2466
	goto L523
L528:
	;
	F_pfree(m, v2466)
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L31
	} else {
		goto L529
	}
L529:
	;
	v2478 = v1620
	goto L523
L530:
	;
	F_relation_close(m, v1553, int32(3))
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L31
	} else {
		goto L531
	}
L531:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v35)+472))
	F_pfree(m, v2484)
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L31
	} else {
		goto L532
	}
L532:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v35)+568))
	F_pfree(m, v2487)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L31
	} else {
		goto L533
	}
L533:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v35)+560))
	F_pfree(m, v2490)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L31
	} else {
		goto L534
	}
L534:
	;
	if v1205 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v35)+584))
	F_pfree(m, v2493)
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L31
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	if v1207 != 0 {
		goto L539
	} else {
		goto L540
	}
L538:
	;
	goto L537
L539:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v35)+592))
	F_pfree(m, v2496)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L31
	} else {
		goto L542
	}
L540:
	;
	goto L541
L541:
	;
	v2501 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v2502 = m.ExcPending
	if v2502 != 0 {
		goto L31
	} else {
		goto L543
	}
L542:
	;
	goto L541
L543:
	;
	v2504 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+56)))
	v2506 = F_SearchSysCacheCopy(m, int32(57), v2504, int64(0))
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L31
	} else {
		goto L544
	}
L544:
	;
	if v2506 == int32(0) {
		goto L464
	} else {
		goto L545
	}
L545:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+16))
	v2511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2510)+22)))
	v2512 = v2510 + v2511
	v2513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2512)+125)))
	if v2513 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L546:
	;
	F_pfree(m, v2506)
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L31
	} else {
		goto L553
	}
L547:
	;
	v2516 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2512)+125)) = uint8(v2516)
	F_CatalogTupleUpdate(m, v2501, v2506+int32(4), v2506)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L31
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	F_CacheInvalidateRelcacheByTuple(m, v2506)
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L31
	} else {
		goto L552
	}
L550:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L31
	} else {
		goto L551
	}
L551:
	;
	goto L546
L552:
	;
	goto L546
L553:
	;
	F_relation_close(m, v2501, int32(3))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L31
	} else {
		goto L554
	}
L554:
	;
	if v1617 != 0 {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v2533 = F_deleteDependencyRecordsFor(m, int32(2620), v1616, int32(1))
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L31
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	v2535 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1616
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2620)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v1531
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(1255)
	F_recordDependencyOn(m, l0, v35+int32(340), int32(110))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L31
	} else {
		goto L559
	}
L558:
	;
	goto L557
L559:
	;
	v2550 = int32(0)
	if base.B2i32(l10 == v2550)|base.B2i32(v1661 == v2550) == v2550 {
		goto L562
	} else {
		goto L563
	}
L560:
	;
	if v2414 == int32(0) {
		goto L577
	} else {
		goto L578
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v2606
	F_recordDependencyOn(m, l0, v35+int32(340), v2607)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L31
	} else {
		goto L576
	}
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(2606)
	v2606 = v1661
	v2607 = int32(105)
	goto L561
L563:
	;
	goto L564
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(1259)
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v2562
	v2567 = v35 + int32(340)
	F_recordDependencyOn(m, l0, v2567, int32(97))
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
		goto L31
	} else {
		goto L565
	}
L565:
	;
	if v286 != 0 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(1259)
	F_recordDependencyOn(m, l0, v2567, int32(97))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L31
	} else {
		goto L569
	}
L567:
	;
	goto L568
L568:
	;
	if v1661 != 0 {
		goto L570
	} else {
		goto L571
	}
L569:
	;
	goto L568
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v1661
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(2606)
	F_recordDependencyOn(m, v35+int32(340), l0, int32(105))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L31
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	if l8 == int32(0) {
		goto L560
	} else {
		goto L574
	}
L573:
	;
	goto L572
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(2620)
	F_recordDependencyOn(m, l0, v35+int32(340), int32(80))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L31
	} else {
		goto L575
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(1259)
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	v2606 = v2603
	v2607 = int32(83)
	goto L561
L576:
	;
	goto L560
L577:
	;
	if v1514 != 0 {
		goto L584
	} else {
		goto L585
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+340)) = int32(1259)
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+344)) = v2621
	if v2418 <= int32(0) {
		goto L577
	} else {
		goto L579
	}
L579:
	;
	v2629 = int32(0)
	goto L580
L580:
	;
	v2661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2414+v2629<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+348)) = v2661
	F_recordDependencyOn(m, l0, v35+int32(340), int32(110))
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L31
	} else {
		goto L582
	}
L581:
	;
	goto L577
L582:
	;
	v2669 = v2629 + int32(1)
	if v2669 != v2418 {
		v2629 = v2669
		goto L580
	} else {
		goto L583
	}
L583:
	;
	goto L581
L584:
	;
	if l10 == int32(0) {
		goto L587
	} else {
		goto L588
	}
L585:
	;
	goto L586
L586:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[4]))
	if v2712 != 0 {
		goto L592
	} else {
		goto L593
	}
L587:
	;
	v2706 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	F_CheckUsageOnTypesInExpr(m, v1500, v1514, v2706)
	mBase = m.M
	v2708 = m.ExcPending
	if v2708 != 0 {
		goto L31
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	F_recordDependencyOnExpr(m, l0, v1500, v1514)
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L31
	} else {
		goto L591
	}
L590:
	;
	goto L589
L591:
	;
	goto L586
L592:
	;
	F_RunObjectPostCreateHook(m, int32(2620), v1616, int32(0), l10)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L31
	} else {
		goto L595
	}
L593:
	;
	goto L594
L594:
	;
	if v356 != 0 {
		goto L596
	} else {
		goto L597
	}
L595:
	;
	goto L594
L596:
	;
	v2718 = F_RelationGetPartitionDesc(m, v44, int32(1))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L31
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	F_relation_close(m, v44, int32(0))
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L31
	} else {
		goto L616
	}
L599:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5]))
	v2726 = F_AllocSetContextCreateInternal(m, v2721, int32(_a_F_CreateTriggerFiringOn_58), int32(0), int32(1024), int32(_a_F_CreateTriggerFiringOn_59))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L31
	} else {
		goto L600
	}
L600:
	;
	v2728 = int32(_a_F_CreateTriggerFiringOn_60)
	v2729 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5])) = v2726
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2718)))
	if int32(0) < v2732 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v2751 = int32(0)
	goto L604
L602:
	;
	goto L603
L603:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5])) = v2729
	F_MemoryContextDelete(m, v2726)
	mBase = m.M
	v2844 = m.ExcPending
	if v2844 != 0 {
		goto L31
	} else {
		goto L615
	}
L604:
	;
	v2769 = v2751 << (uint(int32(2)) % 32)
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2718)+8))
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2769+v2770)))
	v2774 = F_table_open(m, v2772, int32(6))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L31
	} else {
		goto L606
	}
L605:
	;
	goto L603
L606:
	;
	v2776 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L31
	} else {
		goto L607
	}
L607:
	;
	v2778 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2776)+36)) = v2778
	*(*int32)(unsafe.Add(mBase, uint32(v2776)+16)) = v2778
	v2782 = F_copyObjectImpl(m, v1500)
	mBase = m.M
	v2783 = m.ExcPending
	if v2783 != 0 {
		goto L31
	} else {
		goto L608
	}
L608:
	;
	v2785 = F_map_partition_varattnos(m, v2782, int32(1), v2774, v44)
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L31
	} else {
		goto L609
	}
L609:
	;
	v2788 = F_map_partition_varattnos(m, v2785, int32(2), v2774, v44)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L31
	} else {
		goto L610
	}
L610:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2718)+8))
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2792+v2769)))
	v2795 = int32(0)
	F_CreateTriggerFiringOn(m, v35+int32(328), v2776, l2, v2794, l4, v2795, v2795, v1531, v1616, v2788, l10, int32(1), l12)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L31
	} else {
		goto L611
	}
L611:
	;
	F_relation_close(m, v2774, int32(0))
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L31
	} else {
		goto L612
	}
L612:
	;
	F_MemoryContextReset(m, v2726)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L31
	} else {
		goto L613
	}
L613:
	;
	v2806 = v2751 + int32(1)
	v2807 = *(*int32)(unsafe.Add(mBase, uint32(v2718)))
	if v2806 < v2807 {
		v2751 = v2806
		goto L604
	} else {
		goto L614
	}
L614:
	;
	goto L605
L615:
	;
	goto L598
L616:
	;
	m.G0 = v35 + int32(608)
	return
L617:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L31
	} else {
		goto L618
	}
L618:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v2270
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v2890 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_61), v35+int32(48))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L31
	} else {
		goto L619
	}
L619:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(959), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L31
	} else {
		goto L620
	}
L620:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L621:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v2909
	F_errmsg_internal(m, int32(_a_F_CreateTriggerFiringOn_62), v35+int32(32))
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L31
	} else {
		goto L622
	}
L622:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(1031), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L31
	} else {
		goto L623
	}
L623:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_trigger_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_trigger_in_0), int32(366), int32(_a_F_trigger_in_1), int32(_a_F_trigger_in_2), int32(_a_F_trigger_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
