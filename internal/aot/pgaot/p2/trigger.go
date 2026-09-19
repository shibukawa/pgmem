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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1486 int32
	_ = v1486
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1632 int32
	_ = v1632
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int64
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1700 int32
	_ = v1700
	var v1709 int32
	_ = v1709
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1737 int32
	_ = v1737
	var v1749 int32
	_ = v1749
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1775 int32
	_ = v1775
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1920 int32
	_ = v1920
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1943 int32
	_ = v1943
	var v1948 int32
	_ = v1948
	var v1960 int32
	_ = v1960
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v2001 int32
	_ = v2001
	var v2026 int32
	_ = v2026
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2045 int32
	_ = v2045
	var v2065 int32
	_ = v2065
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2136 int32
	_ = v2136
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2182 int32
	_ = v2182
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2232 int32
	_ = v2232
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2272 int32
	_ = v2272
	var v2281 int32
	_ = v2281
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2330 int32
	_ = v2330
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2498 int32
	_ = v2498
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2514 int32
	_ = v2514
	var v2524 int32
	_ = v2524
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2550 int32
	_ = v2550
	var v2557 int32
	_ = v2557
	var v2565 int32
	_ = v2565
	var v2596 int32
	_ = v2596
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2679 int32
	_ = v2679
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2726 int32
	_ = v2726
	var v2729 int32
	_ = v2729
	var v2731 int32
	_ = v2731
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2770 int32
	_ = v2770
	var v2804 int32
	_ = v2804
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2824 int32
	_ = v2824
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2840 int32
	_ = v2840
	var v2845 int32
	_ = v2845
	v14 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(528)
	m.G0 = v34
	if l3 != 0 {
		goto L28
	} else {
		goto L29
	}
L1:
	;
	if l9 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L31
	} else {
		goto L268
	}
L3:
	;
	v1072 = int32(0)
	if base.B2i32(v1067 == v1072)|base.B2i32(v1065 == v1072) != 0 {
		v1194 = v1065
		v1196 = v1067
		goto L1
	} else {
		goto L255
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L31
	} else {
		goto L250
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L31
	} else {
		goto L245
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L31
	} else {
		goto L241
	}
L7:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853)+131)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L31
	} else {
		goto L232
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L31
	} else {
		goto L228
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L31
	} else {
		goto L224
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L31
	} else {
		goto L220
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L31
	} else {
		goto L216
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L31
	} else {
		goto L212
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L31
	} else {
		goto L208
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L31
	} else {
		goto L204
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L31
	} else {
		goto L200
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L31
	} else {
		goto L196
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L31
	} else {
		goto L192
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L31
	} else {
		goto L188
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L31
	} else {
		goto L184
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L31
	} else {
		goto L179
	}
L21:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[0])))
	if v254 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L31
	} else {
		goto L77
	}
L23:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v176 {
	case 0, 2:
		goto L64
	default:
		goto L65
	}
L24:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	if v140 != int32(64) {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v83 {
	case 0, 2:
		goto L40
	default:
		goto L41
	}
L26:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	switch v57 {
	case 0, 2:
		goto L21
	default:
		goto L34
	}
L27:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	v47 = v45 - int32(102)
	switch (v47<<(uint(int32(7))%32) | int32(base.Ui32(v47&int32(254))>>(uint(int32(1))%32))) & int32(255) {
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
	v37 = F_table_open(m, l3, int32(6))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = F_table_openrv(m, v39, int32(6))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L31
	} else {
		goto L33
	}
L31:
	;
	return
L32:
	;
	v43 = v37
	goto L27
L33:
	;
	v43 = v41
	goto L27
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v65 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_0), v34+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_1), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(230), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v109 != int32(1) {
		goto L21
	} else {
		goto L47
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+224)) = v91 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_0), v34+int32(224))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_1), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(241), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v112 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L31
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+240)) = v122 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_4), v34+int32(240))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_5), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(264), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v143 == int32(1) {
		goto L20
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v146&int32(32) == int32(0) {
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
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+256)) = v158 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v34+int32(256))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L31
	} else {
		goto L61
	}
L61:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_7), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L31
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(285), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v202 != int32(1) {
		goto L21
	} else {
		goto L71
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+288)) = v184 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v34+int32(288))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_9), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(295), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L31
	} else {
		goto L73
	}
L73:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+304)) = v212 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v34+int32(304))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L31
	} else {
		goto L74
	}
L74:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_10), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L31
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(307), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v237 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_11), v34)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L31
	} else {
		goto L79
	}
L79:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v245 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244)+119)))
	F_errdetail_relkind_not_supported(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(314), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
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
	v258 = int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	if base.Ui32(v259) < base.Ui32(int32(_a_F_CreateTriggerFiringOn_12)) {
		v268 = v258
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L84
L84:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v269 != int32(1) {
		v285 = v14
		goto L90
	} else {
		goto L91
	}
L85:
	;
	if v268 != 0 {
		goto L19
	} else {
		goto L89
	}
L86:
	;
	goto L85
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+68))
	if v263 == int32(99) {
		v268 = v258
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v266 = F_isTempToastNamespace(m, v263)
	mBase = m.M
	v268 = v266
	goto L86
L89:
	;
	goto L84
L90:
	;
	if l10 != 0 {
		v355 = v14
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
	v274 = m.ExcPending
	if v274 != 0 {
		goto L31
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v275 == int32(0) {
		v285 = v14
		goto L90
	} else {
		goto L96
	}
L95:
	;
	v285 = l4
	goto L90
L96:
	;
	v279 = int32(0)
	v282 = F_RangeVarGetRelidExtended(m, v275, int32(1), v279, v279, v279)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L31
	} else {
		goto L97
	}
L97:
	;
	v285 = v282
	goto L90
L98:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v360 = v356 | (v357 | v358)
	v361 = int32(33)
	if v360&v361 == v361 {
		goto L18
	} else {
		goto L130
	}
L99:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v290 = F_pg_class_aclcheck(m, v286, v288, int64(64))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L31
	} else {
		goto L100
	}
L100:
	;
	if v290 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v293 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292)+119)))
	switch v293 - int32(73) {
	case 0, 32:
		v303 = int32(20)
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
	if v285 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L104:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	F_aclcheck_error(m, v290, v305, v306+int32(4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L31
	} else {
		goto L111
	}
L105:
	;
	v305 = v303
	goto L104
L106:
	;
	v303 = int32(41)
	goto L105
L107:
	;
	v305 = int32(18)
	goto L104
L108:
	;
	v305 = int32(23)
	goto L104
L109:
	;
	v305 = int32(51)
	goto L104
L110:
	;
	v305 = int32(37)
	goto L104
L111:
	;
	goto L103
L112:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v339 != int32(1) {
		v355 = v14
		goto L98
	} else {
		goto L126
	}
L113:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v316 = F_pg_class_aclcheck(m, v285, v314, int64(64))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L31
	} else {
		goto L114
	}
L114:
	;
	if v316 == int32(0) {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v320 = F_get_rel_relkind(m, v285)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L31
	} else {
		goto L116
	}
L116:
	;
	switch v320 - int32(73) {
	case 0, 32:
		v331 = int32(20)
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
	v334 = F_get_rel_name(m, v285)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L31
	} else {
		goto L124
	}
L118:
	;
	v333 = v331
	goto L117
L119:
	;
	v331 = int32(41)
	goto L118
L120:
	;
	v333 = int32(18)
	goto L117
L121:
	;
	v333 = int32(23)
	goto L117
L122:
	;
	v333 = int32(51)
	goto L117
L123:
	;
	v333 = int32(37)
	goto L117
L124:
	;
	F_aclcheck_error(m, v316, v333, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L31
	} else {
		goto L125
	}
L125:
	;
	goto L112
L126:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+119)))
	if v343 != int32(112) {
		v355 = v14
		goto L98
	} else {
		goto L127
	}
L127:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v349 = F_find_all_inheritors(m, v346, int32(6), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L31
	} else {
		goto L128
	}
L128:
	;
	F_list_free(m, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L31
	} else {
		goto L129
	}
L129:
	;
	v355 = int32(1)
	goto L98
L130:
	;
	v366 = v360 & int32(1)
	v368 = v360 & int32(66)
	if v368 == int32(64) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v366 == int32(0) {
		goto L17
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v375 == int32(0) {
		v1194 = v14
		v1196 = v14
		goto L1
	} else {
		goto L137
	}
L134:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v373 != 0 {
		goto L16
	} else {
		goto L135
	}
L135:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v374 != 0 {
		goto L15
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v378 <= int32(0) {
		v1194 = v14
		v1196 = v14
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+9)))
	if v383 != int32(1) {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+119)))
	v389 = v387 - int32(102)
	if v389 == int32(0) {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	if v389 == int32(16) {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	if v366 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v395 = F_has_superclass(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L31
	} else {
		goto L145
	}
L143:
	;
	v398 = v356
	goto L144
L144:
	;
	if v398&int32(_a_F_CreateTriggerFiringOn_13) != 0 {
		goto L6
	} else {
		goto L147
	}
L145:
	;
	if v395 != 0 {
		goto L7
	} else {
		goto L146
	}
L146:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v398 = v397
	goto L144
L147:
	;
	if v360&int32(32) != 0 {
		goto L8
	} else {
		goto L148
	}
L148:
	;
	v405 = int32(1)
	if int32(base.Ui32(v360)>>(uint(int32(2))%32))&v405+int32(base.Ui32(v360)>>(uint(int32(4))%32))&v405+int32(base.Ui32(v360)>>(uint(int32(3))%32))&v405 != v405 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v419 != 0 {
		goto L14
	} else {
		goto L150
	}
L150:
	;
	v421 = v360 & int32(20)
	v423 = v360 & int32(24)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+8)))
	if v424 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v436 < int32(2) {
		v1065 = v435
		v1067 = v434
		goto L3
	} else {
		goto L157
	}
L152:
	;
	if v423 == int32(0) {
		goto L11
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v421 == int32(0) {
		goto L13
	} else {
		goto L156
	}
L155:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v434 = v14
	v435 = v429
	goto L151
L156:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v434 = v432
	v435 = int32(0)
	goto L151
L157:
	;
	v443 = int32(1)
	v464 = v435
	v466 = v434
	goto L158
L158:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471+v443<<(uint(int32(2))%32))))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v476 == int32(0) {
		goto L2
	} else {
		goto L160
	}
L159:
	;
	v1065 = v501
	v1067 = v502
	goto L3
L160:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+119)))
	v482 = v480 - int32(102)
	if v482 == int32(0) {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v482 == int32(16) {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	if v366 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v488 = F_has_superclass(m, v487)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L31
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v491 != 0 {
		goto L14
	} else {
		goto L169
	}
L166:
	;
	if v488 != 0 {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	if v490 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
	if v492 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v504 = v443 + int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v504 < v505 {
		v443 = v504
		v464 = v501
		v466 = v502
		goto L158
	} else {
		goto L178
	}
L171:
	;
	if v421 == int32(0) {
		goto L13
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v423 == int32(0) {
		goto L11
	} else {
		goto L176
	}
L174:
	;
	if v466 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v501 = v464
	v502 = v497
	goto L170
L176:
	;
	if v464 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v501 = v500
	v502 = v466
	goto L170
L178:
	;
	goto L159
L179:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L31
	} else {
		goto L180
	}
L180:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+272)) = v514 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v34+int32(272))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L31
	} else {
		goto L181
	}
L181:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_14), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L31
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(278), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
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
	v538 = m.ExcPending
	if v538 != 0 {
		goto L31
	} else {
		goto L185
	}
L185:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v539 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_15), v34+int32(208))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L31
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(320), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
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
	v559 = m.ExcPending
	if v559 != 0 {
		goto L31
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_16), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L31
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(383), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
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
	v575 = m.ExcPending
	if v575 != 0 {
		goto L31
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_17), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L31
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(391), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
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
	v591 = m.ExcPending
	if v591 != 0 {
		goto L31
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_18), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L31
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(395), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
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
	v607 = m.ExcPending
	if v607 != 0 {
		goto L31
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_19), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L31
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(399), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
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
	v654 = m.ExcPending
	if v654 != 0 {
		goto L31
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_20), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L31
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(508), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
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
	v701 = m.ExcPending
	if v701 != 0 {
		goto L31
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L31
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(525), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
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
	v717 = m.ExcPending
	if v717 != 0 {
		goto L31
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_22), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L31
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(530), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
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
	v764 = m.ExcPending
	if v764 != 0 {
		goto L31
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_23), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L31
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(540), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
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
	v780 = m.ExcPending
	if v780 != 0 {
		goto L31
	} else {
		goto L221
	}
L221:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_24), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L31
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(545), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
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
	v796 = m.ExcPending
	if v796 != 0 {
		goto L31
	} else {
		goto L225
	}
L225:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_25), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L31
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(497), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
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
	v812 = m.ExcPending
	if v812 != 0 {
		goto L31
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_26), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L31
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(480), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
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
	v861 = m.ExcPending
	if v861 != 0 {
		goto L31
	} else {
		goto L233
	}
L233:
	;
	if v854 != int32(1) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_27), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
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
	v876 = m.ExcPending
	if v876 != 0 {
		goto L31
	} else {
		goto L239
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(469), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
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
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(465), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
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
	v919 = m.ExcPending
	if v919 != 0 {
		goto L31
	} else {
		goto L242
	}
L242:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_29), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L31
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(475), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
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
	v966 = m.ExcPending
	if v966 != 0 {
		goto L31
	} else {
		goto L246
	}
L246:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v967 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_6), v34+int32(192))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L31
	} else {
		goto L247
	}
L247:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_30), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L31
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(449), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
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
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L31
	} else {
		goto L251
	}
L251:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+176)) = v1023 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_8), v34+int32(176))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L31
	} else {
		goto L252
	}
L252:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_31), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L31
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(442), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
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
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067))))
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	if base.B2i32(v1079 == int32(0))|base.B2i32(v1079 != v1082) != 0 {
		v1100 = v1079
		v1101 = v1082
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1100-v1101 != 0 {
		v1194 = v1065
		v1196 = v1067
		goto L1
	} else {
		goto L263
	}
L257:
	;
	goto L256
L258:
	;
	v1085 = v1067
	v1086 = v1065
	goto L259
L259:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+1)))
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085)+1)))
	if v1090 == int32(0) {
		v1100 = v1090
		v1101 = v1089
		goto L257
	} else {
		goto L261
	}
L260:
	;
	v1100 = v1090
	v1101 = v1089
	goto L257
L261:
	;
	v1093 = int32(1)
	if v1090 == v1089 {
		v1085 = v1085 + v1093
		v1086 = v1086 + v1093
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
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L31
	} else {
		goto L264
	}
L264:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L31
	} else {
		goto L265
	}
L265:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_32), int32(0))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L31
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(555), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
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
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L31
	} else {
		goto L269
	}
L269:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_33), int32(0))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L31
	} else {
		goto L270
	}
L270:
	;
	F_errhint(m, int32(_a_F_CreateTriggerFiringOn_34), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L31
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(429), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
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
	v2150 = int32(0)
	v2153 = F_DirectFunctionCall1Coll(m, int32(579), v2150, v2136)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L31
	} else {
		goto L457
	}
L274:
	;
	v1981 = F_palloc(m, v1960)
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L31
	} else {
		goto L443
	}
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L31
	} else {
		goto L439
	}
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L31
	} else {
		goto L435
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L31
	} else {
		goto L431
	}
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L31
	} else {
		goto L426
	}
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L31
	} else {
		goto L421
	}
L280:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L31
	} else {
		goto L416
	}
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L31
	} else {
		goto L411
	}
L282:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L31
	} else {
		goto L406
	}
L283:
	;
	if l7 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L284:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v1203 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	goto L286
L286:
	;
	v1475 = F_nodeToString(m, l9)
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L31
	} else {
		goto L352
	}
L287:
	;
	v1486 = int32(0)
	v1499 = v14
	v1506 = v14
	goto L283
L288:
	;
	goto L289
L289:
	;
	v1208 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L31
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1208)+4)) = l2
	v1214 = F_makeAlias(m, int32(_a_F_CreateTriggerFiringOn_35), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L31
	} else {
		goto L291
	}
L291:
	;
	v1216 = int32(0)
	v1218 = F_addRangeTableEntryForRelation(m, v1208, v43, int32(1), v1214, v1216, v1216)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L31
	} else {
		goto L292
	}
L292:
	;
	v1221 = int32(1)
	F_addNSItemToQuery(m, v1208, v1218, int32(0), v1221, v1221)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L31
	} else {
		goto L293
	}
L293:
	;
	v1228 = F_makeAlias(m, int32(_a_F_CreateTriggerFiringOn_36), int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L31
	} else {
		goto L294
	}
L294:
	;
	v1230 = int32(0)
	v1232 = F_addRangeTableEntryForRelation(m, v1208, v43, int32(1), v1228, v1230, v1230)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L31
	} else {
		goto L295
	}
L295:
	;
	v1235 = int32(1)
	F_addNSItemToQuery(m, v1208, v1232, int32(0), v1235, v1235)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L31
	} else {
		goto L296
	}
L296:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v1240 = F_copyObjectImpl(m, v1239)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L31
	} else {
		goto L297
	}
L297:
	;
	v1244 = F_transformWhereClause(m, v1208, v1240, int32(37), int32(_a_F_CreateTriggerFiringOn_37))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L31
	} else {
		goto L298
	}
L298:
	;
	F_assign_expr_collations(m, v1208, v1244)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L31
	} else {
		goto L299
	}
L299:
	;
	v1249 = F_pull_var_clause(m, v1244, int32(0))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L31
	} else {
		goto L301
	}
L300:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+8))
	v1471 = F_nodeToString(m, v1244)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L31
	} else {
		goto L350
	}
L301:
	;
	if v1249 == int32(0) {
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+4))
	if v1253 <= int32(0) {
		goto L300
	} else {
		goto L303
	}
L303:
	;
	v1256 = int32(0)
	if v1256 < v1253 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1259 = v1253
	goto L306
L305:
	;
	v1259 = v1256
	goto L306
L306:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+12))
	v1269 = int32(0)
	goto L307
L307:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1264+v1269<<(uint(int32(2))%32))))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+4))
	switch v1301 - int32(1) {
	case 0:
		goto L312
	case 1:
		goto L311
	default:
		goto L310
	}
L308:
	;
	goto L300
L309:
	;
	v1437 = v1269 + int32(1)
	if v1437 != v1259 {
		v1269 = v1437
		goto L307
	} else {
		goto L349
	}
L310:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L31
	} else {
		goto L346
	}
L311:
	;
	if v366 == int32(0) {
		goto L281
	} else {
		goto L320
	}
L312:
	;
	if v366 == int32(0) {
		goto L282
	} else {
		goto L313
	}
L313:
	;
	if v360&int32(4) == int32(0) {
		goto L309
	} else {
		goto L314
	}
L314:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L31
	} else {
		goto L315
	}
L315:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L31
	} else {
		goto L316
	}
L316:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_38), int32(0))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L31
	} else {
		goto L317
	}
L317:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1319)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L31
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(625), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L31
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	if v360&int32(8) != 0 {
		goto L280
	} else {
		goto L321
	}
L321:
	;
	v1330 = base.B2i32(v368 != int32(2))
	v1331 = int32(0)
	v1333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1300)+8)))
	if base.B2i32(v1330 == v1331)&base.B2i32(v1333 < v1331) != 0 {
		goto L279
	} else {
		goto L322
	}
L322:
	;
	if v368 != int32(2) {
		goto L309
	} else {
		goto L323
	}
L323:
	;
	if v1333 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+16))
	if v1340 == int32(0) {
		goto L309
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if v1333 <= int32(0) {
		goto L309
	} else {
		goto L338
	}
L327:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1340)+17)))
	if v1343 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1340)+18)))
	if v1346 != int32(1) {
		goto L309
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L31
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L31
	} else {
		goto L333
	}
L333:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_39), int32(0))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L31
	} else {
		goto L334
	}
L334:
	;
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_40), int32(0))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L31
	} else {
		goto L335
	}
L335:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1364)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L31
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(653), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L31
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)))
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374+v1375<<(uint(int32(4))%32)+v1333*int32(100))+10)))
	if v1382 == int32(0) {
		goto L309
	} else {
		goto L339
	}
L339:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L31
	} else {
		goto L340
	}
L340:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L31
	} else {
		goto L341
	}
L341:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_39), int32(0))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L31
	} else {
		goto L342
	}
L342:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1396)))
	v1401 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1300)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v1396 + v1397<<(uint(int32(4))%32) + v1401*int32(100) - int32(76)
	F_errdetail(m, int32(_a_F_CreateTriggerFiringOn_41), v34+int32(160))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L31
	} else {
		goto L343
	}
L343:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1413)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L31
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(662), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L31
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	F_errmsg_internal(m, int32(_a_F_CreateTriggerFiringOn_42), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L31
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(666), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L31
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	goto L308
L350:
	;
	F_free_parsestate(m, v1208)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L31
	} else {
		goto L351
	}
L351:
	;
	v1486 = v1244
	v1499 = v1471
	v1506 = v1470
	goto L283
L352:
	;
	v1486 = l9
	v1499 = v1475
	v1506 = v14
	goto L283
L353:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1511 = int32(0)
	v1514 = F_LookupFuncName(m, v1510, v1511, v1511, v1511)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L31
	} else {
		goto L356
	}
L354:
	;
	v1516 = l7
	goto L355
L355:
	;
	if l10 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1516 = v1514
	goto L355
L357:
	;
	v1532 = F_get_func_rettype(m, v1516)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L31
	} else {
		goto L363
	}
L358:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[1]))
	v1521 = F_object_aclcheck(m, int32(1255), v1516, v1519, int64(128))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L31
	} else {
		goto L359
	}
L359:
	;
	if v1521 == int32(0) {
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1527 = F_NameListToString(m, v1526)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L31
	} else {
		goto L361
	}
L361:
	;
	F_aclcheck_error(m, v1521, int32(19), v1527)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L31
	} else {
		goto L362
	}
L362:
	;
	goto L357
L363:
	;
	if v1532 != int32(2279) {
		goto L278
	} else {
		goto L364
	}
L364:
	;
	v1538 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L31
	} else {
		goto L365
	}
L365:
	;
	if l10 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	if l5 != 0 {
		v1646 = l5
		goto L383
	} else {
		goto L384
	}
L367:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1564)+16))
	v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1576)+22)))
	v1578 = v1576 + v1577
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+8))
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578)+83)))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+92))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1578)))
	v1583 = F_heap_copytuple(m, v1564)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L31
	} else {
		goto L378
	}
L368:
	;
	v1543 = v34 + int32(320)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	F_ScanKeyInit(m, v1543, int32(2), int32(3), int32(184), v1547)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L31
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v1573 = F_GetNewOidWithIndex(m, v1538, int32(2702), int32(1))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L31
	} else {
		goto L377
	}
L371:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_ScanKeyInit(m, v34+int32(368), int32(4), int32(3), int32(62), v1555)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L31
	} else {
		goto L372
	}
L372:
	;
	v1562 = F_systable_beginscan(m, v1538, int32(2701), int32(1), int32(0), int32(2), v1543)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L31
	} else {
		goto L373
	}
L373:
	;
	v1564 = F_systable_getnext(m, v1562)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L31
	} else {
		goto L374
	}
L374:
	;
	if v1564 != 0 {
		goto L367
	} else {
		goto L375
	}
L375:
	;
	F_systable_endscan(m, v1562)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L31
	} else {
		goto L376
	}
L376:
	;
	goto L370
L377:
	;
	v1601 = v1573
	v1602 = int32(0)
	v1605 = v14
	goto L366
L378:
	;
	F_systable_endscan(m, v1562)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L31
	} else {
		goto L379
	}
L379:
	;
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v1587 == int32(0) {
		goto L277
	} else {
		goto L380
	}
L380:
	;
	v1590 = int32(0)
	if base.B2i32(v1579 == v1590)&(v1580^int32(-1))|l11 == v1590 {
		goto L276
	} else {
		goto L381
	}
L381:
	;
	if v1581 != 0 {
		goto L275
	} else {
		goto L382
	}
L382:
	;
	v1601 = v1582
	v1602 = int32(1)
	v1605 = v1583
	goto L366
L383:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if l10 != 0 {
		goto L387
	} else {
		goto L388
	}
L384:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v1606&int32(1) == int32(0) {
		v1646 = l5
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+68))
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	v1617 = int32(1)
	v1619 = int32(0)
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v1632 = int32(32)
	v1644 = F_CreateConstraintEntry(m, v1611, v1613, int32(116), v1615, v1616, v1617, v1617, v1619, v1620, v1619, v1619, v1619, v1619, v1619, v1619, v1619, v1619, v1619, v1619, v1619, v1632, v1632, v1619, v1619, v1632, v1619, v1619, v1619, v1617, v1619, v1617, v1619, l10)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L31
	} else {
		goto L386
	}
L386:
	;
	v1646 = v1644
	goto L383
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+84)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v1647
	v1651 = v34 + int32(432)
	v1656 = F_pg_snprintf(m, v1651, int32(64), int32(_a_F_CreateTriggerFiringOn_43), v34+int32(80))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L31
	} else {
		goto L390
	}
L388:
	;
	v1658 = v1647
	goto L389
L389:
	;
	v1659 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+511)) = v1659
	v1661 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+504)) = v1661
	*(*int64)(unsafe.Add(mBase, uint32(v34)+496)) = v1661
	*(*int32)(unsafe.Add(mBase, uint32(v34)+320)) = v1601
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+328)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v34)+324)) = v1666
	v1671 = F_DirectFunctionCall1Coll(m, int32(500), v1659, v1658)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L31
	} else {
		goto L391
	}
L390:
	;
	v1658 = v1651
	goto L389
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+360)) = v1646
	*(*int32)(unsafe.Add(mBase, uint32(v34)+356)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v34)+352)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v34)+348)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v34)+344)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v34)+340)) = base.I32_extend16_s(v360)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+336)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+332)) = v1671
	v1682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+364)) = v1682
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+368)) = v1684
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1686 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+4))
	if v1688 <= int32(0) {
		v1960 = int32(1)
		goto L274
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+372)) = int32(0)
	v2136 = int32(_a_F_CreateTriggerFiringOn_44)
	goto L273
L395:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+12))
	v1692 = int32(0)
	v1700 = v1692
	v1709 = v1692
	goto L396
L396:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v1700<<(uint(int32(2))%32))))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+4))
	v1730 = F_strlen(m, v1729)
	mBase = m.M
	v1737 = v1729
	v1749 = v1730 + v1709 + int32(4)
	goto L398
L398:
	;
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1737))))
	if v1765 != int32(92) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1737 = v1737 + int32(1)
	v1749 = v1775
	goto L398
L401:
	;
	if v1765 != 0 {
		v1775 = v1749
		goto L400
	} else {
		goto L404
	}
L402:
	;
	goto L403
L403:
	;
	v1775 = v1749 + int32(1)
	goto L400
L404:
	;
	v1769 = v1700 + int32(1)
	if v1769 != v1688 {
		v1700 = v1769
		v1709 = v1749
		goto L396
	} else {
		goto L405
	}
L405:
	;
	v1960 = v1749 + int32(1)
	goto L274
L406:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L31
	} else {
		goto L407
	}
L407:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_45), int32(0))
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L31
	} else {
		goto L408
	}
L408:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1792)
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L31
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(620), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L31
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L31
	} else {
		goto L412
	}
L412:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_45), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L31
	} else {
		goto L413
	}
L413:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1811)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L31
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(633), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L31
	} else {
		goto L415
	}
L415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L416:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L31
	} else {
		goto L417
	}
L417:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_46), int32(0))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L31
	} else {
		goto L418
	}
L418:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1830)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L31
	} else {
		goto L419
	}
L419:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(638), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L31
	} else {
		goto L420
	}
L420:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L421:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L31
	} else {
		goto L422
	}
L422:
	;
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_47), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L31
	} else {
		goto L423
	}
L423:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+44))
	F_parser_errposition(m, v1208, v1849)
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L31
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(643), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L31
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L426:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L31
	} else {
		goto L427
	}
L427:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1865 = F_NameListToString(m, v1864)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L31
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+148)) = int32(_a_F_CreateTriggerFiringOn_48)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+144)) = v1865
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_49), v34+int32(144))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L31
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(707), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L31
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_50))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L31
	} else {
		goto L432
	}
L432:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v1888
	*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v1887 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_51), v34+int32(128))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L31
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(768), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L31
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_50))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L31
	} else {
		goto L436
	}
L436:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = v1910 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_52), v34+int32(112))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L31
	} else {
		goto L437
	}
L437:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(781), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L31
	} else {
		goto L438
	}
L438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L439:
	;
	F_errcode(m, int32(_a_F_CreateTriggerFiringOn_50))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L31
	} else {
		goto L440
	}
L440:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v1934
	*(*int32)(unsafe.Add(mBase, uint32(v34)+100)) = v1933 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_53), v34+int32(96))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L31
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(800), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L31
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L443:
	;
	v1983 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1981))) = uint8(v1983)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1985 == v1983 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+372)) = base.I32_extend16_s(v1688)
	v2136 = v1981
	goto L273
L445:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+4))
	if v1988 <= int32(0) {
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[2])))
	v1994 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[3]))
	v2001 = int32(0)
	goto L447
L447:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+12))
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v2026+v2001<<(uint(int32(2))%32))))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+4))
	v2032 = F_strlen(m, v1981)
	mBase = m.M
	v2037 = v2032 + v1981
	v2045 = v2031
	goto L449
L449:
	;
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2045))))
	if v2065 != int32(92) {
		goto L452
	} else {
		goto L453
	}
L451:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2079))) = uint8(v2080)
	v2082 = int32(1)
	v2037 = v2079 + v2082
	v2045 = v2045 + v2082
	goto L449
L452:
	;
	if v2065 != 0 {
		v2079 = v2037
		v2080 = v2065
		goto L451
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v2074 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v2037))) = uint8(v2074)
	v2078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2045))))
	v2079 = v2037 + int32(1)
	v2080 = v2078
	goto L451
L455:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2037)+4)) = uint8(v1992)
	*(*int32)(unsafe.Add(mBase, uint32(v2037))) = v1994
	v2071 = v2001 + int32(1)
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+4))
	if v2072 <= v2071 {
		goto L444
	} else {
		goto L456
	}
L456:
	;
	v2001 = v2071
	goto L447
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+380)) = v2153
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v2156 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L458:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L31
	} else {
		goto L611
	}
L459:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L31
	} else {
		goto L607
	}
L460:
	;
	v2375 = F_buildint2vector(m, v2355, v2359)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L31
	} else {
		goto L500
	}
L461:
	;
	v2355 = int32(0)
	v2359 = v2150
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2156)+4))
	if v2160 == int32(0) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2163 = int32(0)
	v2355 = v2163
	v2359 = v2163
	goto L460
L465:
	;
	goto L466
L466:
	;
	v2167 = F_palloc(m, v2160<<(uint(int32(1))%32))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L31
	} else {
		goto L467
	}
L467:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v2169 == int32(0) {
		v2355 = v2167
		v2359 = v2160
		goto L460
	} else {
		goto L468
	}
L468:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2172 <= int32(0) {
		v2355 = v2167
		v2359 = v2160
		goto L460
	} else {
		goto L469
	}
L469:
	;
	v2182 = int32(0)
	goto L470
L470:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+12))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2207+v2182<<(uint(int32(2))%32))))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2211)+4))
	v2213 = int32(0)
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v2217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2216)+120)))
	if v2213 < v2217 {
		goto L475
	} else {
		goto L476
	}
L471:
	;
	v2355 = v2167
	v2359 = v2160
	goto L460
L472:
	;
	if v2272<<(uint(int32(16))%32) == int32(0) {
		goto L459
	} else {
		goto L489
	}
L473:
	;
	v2272 = v2223 + int32(1)
	goto L472
L474:
	;
	goto L473
L475:
	;
	v2223 = v2213
	goto L478
L476:
	;
	goto L477
L477:
	;
	goto L485
L478:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2225)))
	v2232 = v2225 + v2226<<(uint(int32(4))%32) + v2223*int32(100)
	v2235 = F_namestrcmp(m, v2232+int32(24), v2212)
	mBase = m.M
	if v2235 == int32(0) {
		goto L480
	} else {
		goto L481
	}
L479:
	;
	goto L477
L480:
	;
	v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2232)+111)))
	if v2238 != int32(1) {
		goto L474
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	v2242 = v2223 + int32(1)
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v2244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2243)+120)))
	if v2242 < v2244 {
		v2223 = v2242
		goto L478
	} else {
		goto L484
	}
L483:
	;
	goto L482
L484:
	;
	goto L479
L485:
	;
	v2272 = int32(0)
	goto L472
L489:
	;
	v2281 = v2182
	goto L491
L490:
	;
	v2336 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2167+v2182<<(uint(v2336)%32)))) = uint16(v2272)
	v2341 = v2182 + v2336
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2341 < v2342 {
		v2182 = v2341
		goto L470
	} else {
		goto L499
	}
L491:
	;
	if v2281 <= int32(0) {
		goto L490
	} else {
		goto L493
	}
L492:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L31
	} else {
		goto L495
	}
L493:
	;
	v2311 = int32(1)
	v2312 = v2281 - v2311
	v2316 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2167+v2312<<(uint(v2311)%32)))))
	if base.I32_extend16_s(v2272) != v2316 {
		v2281 = v2312
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L31
	} else {
		goto L496
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v2212
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_54), v34-int32(-64))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L31
	} else {
		goto L497
	}
L497:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(958), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L31
	} else {
		goto L498
	}
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L499:
	;
	goto L471
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+376)) = v2375
	if v1499 != 0 {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v1194 != 0 {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2378 = F_cstring_to_text(m, v1499)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L31
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2381 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+512)) = uint8(v2381)
	goto L501
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+384)) = v2378
	goto L501
L506:
	;
	if v1196 != 0 {
		goto L512
	} else {
		goto L513
	}
L507:
	;
	v2385 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v1194)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L31
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v2388 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+513)) = uint8(v2388)
	goto L506
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+388)) = v2385
	goto L506
L511:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+52))
	v2402 = F_heap_form_tuple(m, v2397, v34+int32(320), v34+int32(496))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L31
	} else {
		goto L516
	}
L512:
	;
	v2392 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v1196)
	mBase = m.M
	v2393 = m.ExcPending
	if v2393 != 0 {
		goto L31
	} else {
		goto L515
	}
L513:
	;
	goto L514
L514:
	;
	v2395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+514)) = uint8(v2395)
	goto L511
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+392)) = v2392
	goto L511
L516:
	;
	if v1602 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	F_pfree(m, v2414)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L31
	} else {
		goto L524
	}
L518:
	;
	F_CatalogTupleInsert(m, v1538, v2402)
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L31
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	F_CatalogTupleUpdate(m, v1538, v1605+int32(4), v2402)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L31
	} else {
		goto L522
	}
L521:
	;
	v2414 = v2402
	goto L517
L522:
	;
	F_pfree(m, v2402)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L31
	} else {
		goto L523
	}
L523:
	;
	v2414 = v1605
	goto L517
L524:
	;
	F_relation_close(m, v1538, int32(3))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L31
	} else {
		goto L525
	}
L525:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v34)+332))
	F_pfree(m, v2420)
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L31
	} else {
		goto L526
	}
L526:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v34)+380))
	F_pfree(m, v2423)
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L31
	} else {
		goto L527
	}
L527:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v34)+376))
	F_pfree(m, v2426)
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L31
	} else {
		goto L528
	}
L528:
	;
	if v1194 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v34)+388))
	F_pfree(m, v2429)
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L31
	} else {
		goto L532
	}
L530:
	;
	goto L531
L531:
	;
	if v1196 != 0 {
		goto L533
	} else {
		goto L534
	}
L532:
	;
	goto L531
L533:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v34)+392))
	F_pfree(m, v2432)
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L31
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	v2437 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L31
	} else {
		goto L537
	}
L536:
	;
	goto L535
L537:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v2442 = F_SearchSysCacheCopy(m, int32(57), v2440, int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L31
	} else {
		goto L538
	}
L538:
	;
	if v2442 == int32(0) {
		goto L458
	} else {
		goto L539
	}
L539:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v2442)+16))
	v2447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2446)+22)))
	v2448 = v2446 + v2447
	v2449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2448)+125)))
	if v2449 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	F_pfree(m, v2442)
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L31
	} else {
		goto L547
	}
L541:
	;
	v2452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2448)+125)) = uint8(v2452)
	F_CatalogTupleUpdate(m, v2437, v2442+int32(4), v2442)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L31
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	F_CacheInvalidateRelcacheByTuple(m, v2442)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L31
	} else {
		goto L546
	}
L544:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L31
	} else {
		goto L545
	}
L545:
	;
	goto L540
L546:
	;
	goto L540
L547:
	;
	F_relation_close(m, v2437, int32(3))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L31
	} else {
		goto L548
	}
L548:
	;
	if v1602 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v2469 = F_deleteDependencyRecordsFor(m, int32(2620), v1601, int32(1))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L31
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	v2471 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2471
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2620)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = v2471
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(1255)
	F_recordDependencyOn(m, l0, v34+int32(420), int32(110))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L31
	} else {
		goto L553
	}
L552:
	;
	goto L551
L553:
	;
	v2486 = int32(0)
	if base.B2i32(l10 == v2486)|base.B2i32(v1646 == v2486) == v2486 {
		goto L556
	} else {
		goto L557
	}
L554:
	;
	if v2355 == int32(0) {
		goto L571
	} else {
		goto L572
	}
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v2542
	F_recordDependencyOn(m, l0, v34+int32(420), v2543)
	mBase = m.M
	v2550 = m.ExcPending
	if v2550 != 0 {
		goto L31
	} else {
		goto L570
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(2606)
	v2542 = v1646
	v2543 = int32(105)
	goto L555
L557:
	;
	goto L558
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(1259)
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v2498
	v2503 = v34 + int32(420)
	F_recordDependencyOn(m, l0, v2503, int32(97))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L31
	} else {
		goto L559
	}
L559:
	;
	if v285 != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(1259)
	F_recordDependencyOn(m, l0, v2503, int32(97))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L31
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	if v1646 != 0 {
		goto L564
	} else {
		goto L565
	}
L563:
	;
	goto L562
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v1646
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(2606)
	F_recordDependencyOn(m, v34+int32(420), l0, int32(105))
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L31
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	if l8 == int32(0) {
		goto L554
	} else {
		goto L568
	}
L567:
	;
	goto L566
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(2620)
	F_recordDependencyOn(m, l0, v34+int32(420), int32(80))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L31
	} else {
		goto L569
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(1259)
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v2542 = v2539
	v2543 = int32(83)
	goto L555
L570:
	;
	goto L554
L571:
	;
	if v1506 != 0 {
		goto L578
	} else {
		goto L579
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+420)) = int32(1259)
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+424)) = v2557
	if v2359 <= int32(0) {
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v2565 = int32(0)
	goto L574
L574:
	;
	v2596 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2355+v2565<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+428)) = v2596
	F_recordDependencyOn(m, l0, v34+int32(420), int32(110))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L31
	} else {
		goto L576
	}
L575:
	;
	goto L571
L576:
	;
	v2604 = v2565 + int32(1)
	if v2604 != v2359 {
		v2565 = v2604
		goto L574
	} else {
		goto L577
	}
L577:
	;
	goto L575
L578:
	;
	F_recordDependencyOnExpr(m, l0, v1486, v1506)
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L31
	} else {
		goto L581
	}
L579:
	;
	goto L580
L580:
	;
	v2640 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[4]))
	if v2640 != 0 {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	goto L580
L582:
	;
	F_RunObjectPostCreateHook(m, int32(2620), v1601, int32(0), l10)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L31
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	if v355 != 0 {
		goto L586
	} else {
		goto L587
	}
L585:
	;
	goto L584
L586:
	;
	v2646 = F_RelationGetPartitionDesc(m, v43, int32(1))
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L31
	} else {
		goto L589
	}
L587:
	;
	goto L588
L588:
	;
	F_relation_close(m, v43, int32(0))
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L31
	} else {
		goto L606
	}
L589:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5]))
	v2654 = F_AllocSetContextCreateInternal(m, v2649, int32(_a_F_CreateTriggerFiringOn_55), int32(0), int32(1024), int32(_a_F_CreateTriggerFiringOn_56))
	mBase = m.M
	v2655 = m.ExcPending
	if v2655 != 0 {
		goto L31
	} else {
		goto L590
	}
L590:
	;
	v2656 = int32(_a_F_CreateTriggerFiringOn_57)
	v2657 = *(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5])) = v2654
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2646)))
	if int32(0) < v2660 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v2679 = int32(0)
	goto L594
L592:
	;
	goto L593
L593:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateTriggerFiringOn[5])) = v2657
	F_MemoryContextDelete(m, v2654)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L31
	} else {
		goto L605
	}
L594:
	;
	v2696 = v2679 << (uint(int32(2)) % 32)
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+8))
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v2696+v2697)))
	v2701 = F_table_open(m, v2699, int32(6))
	mBase = m.M
	v2702 = m.ExcPending
	if v2702 != 0 {
		goto L31
	} else {
		goto L596
	}
L595:
	;
	goto L593
L596:
	;
	v2703 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v2704 = m.ExcPending
	if v2704 != 0 {
		goto L31
	} else {
		goto L597
	}
L597:
	;
	v2705 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2703)+36)) = v2705
	*(*int32)(unsafe.Add(mBase, uint32(v2703)+16)) = v2705
	v2709 = F_copyObjectImpl(m, v1486)
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L31
	} else {
		goto L598
	}
L598:
	;
	v2712 = F_map_partition_varattnos(m, v2709, int32(1), v2701, v43)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L31
	} else {
		goto L599
	}
L599:
	;
	v2715 = F_map_partition_varattnos(m, v2712, int32(2), v2701, v43)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L31
	} else {
		goto L600
	}
L600:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+8))
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2719+v2696)))
	v2722 = int32(0)
	F_CreateTriggerFiringOn(m, v34+int32(308), v2703, l2, v2721, l4, v2722, v2722, v1516, v1601, v2715, l10, int32(1), l12)
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L31
	} else {
		goto L601
	}
L601:
	;
	F_relation_close(m, v2701, int32(0))
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L31
	} else {
		goto L602
	}
L602:
	;
	F_MemoryContextReset(m, v2654)
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L31
	} else {
		goto L603
	}
L603:
	;
	v2733 = v2679 + int32(1)
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2646)))
	if v2733 < v2734 {
		v2679 = v2733
		goto L594
	} else {
		goto L604
	}
L604:
	;
	goto L595
L605:
	;
	goto L588
L606:
	;
	m.G0 = v34 + int32(528)
	return
L607:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L31
	} else {
		goto L608
	}
L608:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v2815 + int32(4)
	F_errmsg(m, int32(_a_F_CreateTriggerFiringOn_58), v34+int32(48))
	mBase = m.M
	v2824 = m.ExcPending
	if v2824 != 0 {
		goto L31
	} else {
		goto L609
	}
L609:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(949), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		goto L31
	} else {
		goto L610
	}
L610:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L611:
	;
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v2834
	F_errmsg_internal(m, int32(_a_F_CreateTriggerFiringOn_59), v34+int32(32))
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L31
	} else {
		goto L612
	}
L612:
	;
	F_errfinish(m, int32(_a_F_CreateTriggerFiringOn_2), int32(1021), int32(_a_F_CreateTriggerFiringOn_3))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L31
	} else {
		goto L613
	}
L613:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_trigger_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_trigger_in_0), int32(366), int32(_a_F_trigger_in_1), int32(_a_F_trigger_in_2), int32(_a_F_trigger_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
