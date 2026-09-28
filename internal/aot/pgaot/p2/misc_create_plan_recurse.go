package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_plan_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v322 float64
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 float64
	_ = v374
	var v375 float64
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
	var v389 float64
	_ = v389
	var v392 int32
	_ = v392
	var v393 float64
	_ = v393
	var v399 float64
	_ = v399
	var v405 float64
	_ = v405
	var v407 float64
	_ = v407
	var v409 float64
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 float64
	_ = v493
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v503 float64
	_ = v503
	var v509 float64
	_ = v509
	var v511 float64
	_ = v511
	var v513 float64
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v553 float64
	_ = v553
	var v555 float64
	_ = v555
	var v557 float64
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 float64
	_ = v566
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v812 int32
	_ = v812
	var v853 int32
	_ = v853
	var v902 int32
	_ = v902
	var v909 int32
	_ = v909
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 float64
	_ = v1253
	var v1255 float64
	_ = v1255
	var v1257 float64
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1320 int32
	_ = v1320
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1437 int64
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1451 int32
	_ = v1451
	var v1457 int32
	_ = v1457
	var v1459 float64
	_ = v1459
	var v1461 float64
	_ = v1461
	var v1463 float64
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1555 int64
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1648 int32
	_ = v1648
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1783 int32
	_ = v1783
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1989 int32
	_ = v1989
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2049 int32
	_ = v2049
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2111 int32
	_ = v2111
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2166 int32
	_ = v2166
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2218 int32
	_ = v2218
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2257 int32
	_ = v2257
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
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
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2314 int32
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2409 int32
	_ = v2409
	var v2443 int32
	_ = v2443
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2469 int32
	_ = v2469
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2540 int32
	_ = v2540
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2573 int32
	_ = v2573
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2587 int32
	_ = v2587
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2930 int32
	_ = v2930
	var v2958 int32
	_ = v2958
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v3026 int32
	_ = v3026
	var v3029 int32
	_ = v3029
	var v3033 int32
	_ = v3033
	var v3038 int32
	_ = v3038
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3067 int32
	_ = v3067
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
	var v3081 int32
	_ = v3081
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3101 int32
	_ = v3101
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3249 int32
	_ = v3249
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3256 int32
	_ = v3256
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3267 int32
	_ = v3267
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3288 int32
	_ = v3288
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3300 int32
	_ = v3300
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3325 int32
	_ = v3325
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3331 int32
	_ = v3331
	var v3333 int32
	_ = v3333
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3345 int32
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3355 int32
	_ = v3355
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3382 int32
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3392 int32
	_ = v3392
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3401 int32
	_ = v3401
	var v3411 int32
	_ = v3411
	var v3414 int32
	_ = v3414
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3455 int32
	_ = v3455
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3466 int32
	_ = v3466
	var v3469 int32
	_ = v3469
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3491 int32
	_ = v3491
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3505 int32
	_ = v3505
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3522 int32
	_ = v3522
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3537 int32
	_ = v3537
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3549 int32
	_ = v3549
	var v3552 int32
	_ = v3552
	var v3556 int32
	_ = v3556
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3599 int32
	_ = v3599
	var v3600 int32
	_ = v3600
	var v3604 int32
	_ = v3604
	var v3607 int32
	_ = v3607
	var v3614 int32
	_ = v3614
	var v3618 int32
	_ = v3618
	var v3623 int32
	_ = v3623
	var v3625 int32
	_ = v3625
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3649 int32
	_ = v3649
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3658 int32
	_ = v3658
	var v3661 int32
	_ = v3661
	var v3666 int32
	_ = v3666
	var v3670 int32
	_ = v3670
	var v3672 int32
	_ = v3672
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3683 int32
	_ = v3683
	var v3685 int32
	_ = v3685
	var v3698 int32
	_ = v3698
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3715 int32
	_ = v3715
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3729 int32
	_ = v3729
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3741 int32
	_ = v3741
	var v3742 int32
	_ = v3742
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3757 int32
	_ = v3757
	var v3761 int32
	_ = v3761
	var v3766 int32
	_ = v3766
	var v3770 int32
	_ = v3770
	var v3779 int32
	_ = v3779
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3791 int32
	_ = v3791
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3802 int32
	_ = v3802
	var v3804 int32
	_ = v3804
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3812 int32
	_ = v3812
	var v3813 int32
	_ = v3813
	var v3817 int32
	_ = v3817
	var v3820 int32
	_ = v3820
	var v3824 int32
	_ = v3824
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3838 int32
	_ = v3838
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3903 int32
	_ = v3903
	var v3905 float64
	_ = v3905
	var v3907 float64
	_ = v3907
	var v3909 float64
	_ = v3909
	var v3911 int32
	_ = v3911
	var v3912 int32
	_ = v3912
	var v3914 int32
	_ = v3914
	var v3916 int32
	_ = v3916
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3921 int64
	_ = v3921
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3929 int32
	_ = v3929
	var v3935 int32
	_ = v3935
	var v3937 float64
	_ = v3937
	var v3939 float64
	_ = v3939
	var v3941 float64
	_ = v3941
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3946 int32
	_ = v3946
	var v3948 int32
	_ = v3948
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
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
	var v3964 int32
	_ = v3964
	var v3967 int32
	_ = v3967
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3973 int32
	_ = v3973
	var v4016 int32
	_ = v4016
	var v4020 int32
	_ = v4020
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4035 int32
	_ = v4035
	var v4039 int32
	_ = v4039
	var v4040 int32
	_ = v4040
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4047 int32
	_ = v4047
	var v4093 float64
	_ = v4093
	var v4094 int32
	_ = v4094
	var v4095 int32
	_ = v4095
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4101 int32
	_ = v4101
	var v4102 int32
	_ = v4102
	var v4107 int32
	_ = v4107
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4119 int32
	_ = v4119
	var v4120 int32
	_ = v4120
	var v4123 int32
	_ = v4123
	var v4127 int32
	_ = v4127
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4300 int32
	_ = v4300
	var v4302 float64
	_ = v4302
	var v4304 float64
	_ = v4304
	var v4306 float64
	_ = v4306
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4313 int32
	_ = v4313
	var v4315 int32
	_ = v4315
	var v4316 int32
	_ = v4316
	var v4320 int32
	_ = v4320
	var v4323 int32
	_ = v4323
	var v4327 int32
	_ = v4327
	var v4329 int32
	_ = v4329
	var v4330 int32
	_ = v4330
	var v4372 int32
	_ = v4372
	var v4376 int32
	_ = v4376
	var v4377 int32
	_ = v4377
	var v4378 int32
	_ = v4378
	var v4379 int32
	_ = v4379
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4391 int32
	_ = v4391
	var v4395 int32
	_ = v4395
	var v4396 int32
	_ = v4396
	var v4398 int32
	_ = v4398
	var v4399 int32
	_ = v4399
	var v4407 int32
	_ = v4407
	var v4449 int32
	_ = v4449
	var v4451 int32
	_ = v4451
	var v4452 int32
	_ = v4452
	var v4453 int32
	_ = v4453
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4457 float64
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4460 int32
	_ = v4460
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4474 int32
	_ = v4474
	var v4475 int32
	_ = v4475
	var v4477 int32
	_ = v4477
	var v4478 int32
	_ = v4478
	var v4480 int32
	_ = v4480
	var v4481 int32
	_ = v4481
	var v4483 int32
	_ = v4483
	var v4484 int32
	_ = v4484
	var v4487 int32
	_ = v4487
	var v4494 int32
	_ = v4494
	var v4496 int32
	_ = v4496
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4551 int32
	_ = v4551
	var v4552 int32
	_ = v4552
	var v4553 int32
	_ = v4553
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4559 int32
	_ = v4559
	var v4562 int32
	_ = v4562
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4567 int32
	_ = v4567
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4629 int32
	_ = v4629
	var v4631 float64
	_ = v4631
	var v4633 float64
	_ = v4633
	var v4635 float64
	_ = v4635
	var v4637 int32
	_ = v4637
	var v4638 int32
	_ = v4638
	var v4640 int32
	_ = v4640
	var v4642 int32
	_ = v4642
	var v4644 int32
	_ = v4644
	var v4645 int32
	_ = v4645
	var v4646 int32
	_ = v4646
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4649 int32
	_ = v4649
	var v4650 int32
	_ = v4650
	var v4651 int32
	_ = v4651
	var v4653 int32
	_ = v4653
	var v4654 int32
	_ = v4654
	var v4655 int32
	_ = v4655
	var v4656 int32
	_ = v4656
	var v4660 int32
	_ = v4660
	var v4663 int32
	_ = v4663
	var v4667 int32
	_ = v4667
	var v4668 int32
	_ = v4668
	var v4671 int32
	_ = v4671
	var v4713 int32
	_ = v4713
	var v4717 int32
	_ = v4717
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4720 int32
	_ = v4720
	var v4721 int32
	_ = v4721
	var v4723 int32
	_ = v4723
	var v4725 int32
	_ = v4725
	var v4726 int32
	_ = v4726
	var v4732 int32
	_ = v4732
	var v4736 int32
	_ = v4736
	var v4737 int32
	_ = v4737
	var v4739 int32
	_ = v4739
	var v4740 int32
	_ = v4740
	var v4748 int32
	_ = v4748
	var v4791 int32
	_ = v4791
	var v4792 int32
	_ = v4792
	var v4794 int32
	_ = v4794
	var v4795 int32
	_ = v4795
	var v4797 int32
	_ = v4797
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4802 int32
	_ = v4802
	var v4809 int32
	_ = v4809
	var v4857 int32
	_ = v4857
	var v4858 int32
	_ = v4858
	var v4860 int32
	_ = v4860
	var v4861 int32
	_ = v4861
	var v4862 int32
	_ = v4862
	var v4863 int32
	_ = v4863
	var v4864 int32
	_ = v4864
	var v4867 int32
	_ = v4867
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4872 int32
	_ = v4872
	var v4875 int32
	_ = v4875
	var v4876 int32
	_ = v4876
	var v4882 int32
	_ = v4882
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4930 int32
	_ = v4930
	var v4931 int32
	_ = v4931
	var v4933 int32
	_ = v4933
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4936 int32
	_ = v4936
	var v4939 int32
	_ = v4939
	var v4942 int32
	_ = v4942
	var v4994 int32
	_ = v4994
	var v4995 int32
	_ = v4995
	var v4997 int32
	_ = v4997
	var v4998 int32
	_ = v4998
	var v4999 int32
	_ = v4999
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5004 int32
	_ = v5004
	var v5007 int32
	_ = v5007
	var v5008 int32
	_ = v5008
	var v5009 int32
	_ = v5009
	var v5012 int32
	_ = v5012
	var v5013 int32
	_ = v5013
	var v5015 int32
	_ = v5015
	var v5063 int32
	_ = v5063
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5071 int32
	_ = v5071
	var v5073 int32
	_ = v5073
	var v5083 int32
	_ = v5083
	var v5085 int32
	_ = v5085
	var v5087 int32
	_ = v5087
	var v5091 int32
	_ = v5091
	var v5093 int32
	_ = v5093
	var v5095 int32
	_ = v5095
	var v5097 int32
	_ = v5097
	var v5099 int32
	_ = v5099
	var v5107 int32
	_ = v5107
	var v5109 float64
	_ = v5109
	var v5111 float64
	_ = v5111
	var v5113 float64
	_ = v5113
	var v5115 int32
	_ = v5115
	var v5116 int32
	_ = v5116
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5122 int32
	_ = v5122
	var v5125 int32
	_ = v5125
	var v5126 int32
	_ = v5126
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5130 int32
	_ = v5130
	var v5131 int32
	_ = v5131
	var v5134 int32
	_ = v5134
	var v5137 int32
	_ = v5137
	var v5140 int32
	_ = v5140
	var v5142 int32
	_ = v5142
	var v5143 int32
	_ = v5143
	var v5148 int32
	_ = v5148
	var v5152 int32
	_ = v5152
	var v5155 int32
	_ = v5155
	var v5156 int32
	_ = v5156
	var v5200 int32
	_ = v5200
	var v5201 int32
	_ = v5201
	var v5202 int32
	_ = v5202
	var v5203 int32
	_ = v5203
	var v5204 int32
	_ = v5204
	var v5205 int32
	_ = v5205
	var v5206 int32
	_ = v5206
	var v5207 int32
	_ = v5207
	var v5208 int32
	_ = v5208
	var v5210 int32
	_ = v5210
	var v5212 int32
	_ = v5212
	var v5214 int32
	_ = v5214
	var v5216 int32
	_ = v5216
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5220 int32
	_ = v5220
	var v5226 int32
	_ = v5226
	var v5229 int32
	_ = v5229
	var v5272 int32
	_ = v5272
	var v5276 int32
	_ = v5276
	var v5278 int32
	_ = v5278
	var v5279 int32
	_ = v5279
	var v5325 int32
	_ = v5325
	var v5326 int32
	_ = v5326
	var v5328 int32
	_ = v5328
	var v5329 int32
	_ = v5329
	var v5332 int32
	_ = v5332
	var v5336 int32
	_ = v5336
	var v5388 int32
	_ = v5388
	var v5434 int32
	_ = v5434
	var v5435 int32
	_ = v5435
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5445 int32
	_ = v5445
	var v5491 int32
	_ = v5491
	var v5495 int32
	_ = v5495
	var v5496 int32
	_ = v5496
	var v5497 int32
	_ = v5497
	var v5498 int32
	_ = v5498
	var v5499 int32
	_ = v5499
	var v5500 int32
	_ = v5500
	var v5503 int32
	_ = v5503
	var v5506 int32
	_ = v5506
	var v5507 int32
	_ = v5507
	var v5560 int32
	_ = v5560
	var v5563 int32
	_ = v5563
	var v5564 int32
	_ = v5564
	var v5565 int32
	_ = v5565
	var v5576 int32
	_ = v5576
	var v5580 int32
	_ = v5580
	var v5582 int32
	_ = v5582
	var v5615 int32
	_ = v5615
	var v5619 int32
	_ = v5619
	var v5620 int32
	_ = v5620
	var v5625 int32
	_ = v5625
	var v5626 int32
	_ = v5626
	var v5627 int32
	_ = v5627
	var v5629 int32
	_ = v5629
	var v5630 int32
	_ = v5630
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5633 int32
	_ = v5633
	var v5638 int32
	_ = v5638
	var v5684 int32
	_ = v5684
	var v5687 int32
	_ = v5687
	var v5691 int32
	_ = v5691
	var v5692 int32
	_ = v5692
	var v5696 int32
	_ = v5696
	var v5699 int32
	_ = v5699
	var v5700 int32
	_ = v5700
	var v5706 int32
	_ = v5706
	var v5750 int32
	_ = v5750
	var v5756 int32
	_ = v5756
	var v5758 int32
	_ = v5758
	var v5759 int32
	_ = v5759
	var v5760 int32
	_ = v5760
	var v5761 int32
	_ = v5761
	var v5764 int32
	_ = v5764
	var v5765 int32
	_ = v5765
	var v5767 int32
	_ = v5767
	var v5768 int32
	_ = v5768
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5771 int32
	_ = v5771
	var v5772 int32
	_ = v5772
	var v5773 int32
	_ = v5773
	var v5776 int32
	_ = v5776
	var v5781 int32
	_ = v5781
	var v5828 int32
	_ = v5828
	var v5829 int32
	_ = v5829
	var v5831 int32
	_ = v5831
	var v5833 int32
	_ = v5833
	var v5835 int32
	_ = v5835
	var v5839 int32
	_ = v5839
	var v5842 int32
	_ = v5842
	var v5845 int32
	_ = v5845
	var v5846 int32
	_ = v5846
	var v5850 int32
	_ = v5850
	var v5858 int32
	_ = v5858
	var v5859 int32
	_ = v5859
	var v5862 int32
	_ = v5862
	var v5873 int32
	_ = v5873
	var v5878 int32
	_ = v5878
	var v5881 int32
	_ = v5881
	var v5884 int32
	_ = v5884
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5889 int32
	_ = v5889
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5897 int32
	_ = v5897
	var v5944 int32
	_ = v5944
	var v5945 int32
	_ = v5945
	var v5948 int32
	_ = v5948
	var v5950 int32
	_ = v5950
	var v5952 int32
	_ = v5952
	var v5958 int32
	_ = v5958
	var v5967 int32
	_ = v5967
	var v5968 int32
	_ = v5968
	var v5978 int32
	_ = v5978
	var v6021 int32
	_ = v6021
	var v6022 int32
	_ = v6022
	var v6023 int32
	_ = v6023
	var v6028 int32
	_ = v6028
	var v6032 int32
	_ = v6032
	var v6037 int32
	_ = v6037
	var v6043 int32
	_ = v6043
	var v6086 int32
	_ = v6086
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6089 int32
	_ = v6089
	var v6094 int32
	_ = v6094
	var v6097 int32
	_ = v6097
	var v6105 int32
	_ = v6105
	var v6140 int32
	_ = v6140
	var v6141 int32
	_ = v6141
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6151 float64
	_ = v6151
	var v6152 int64
	_ = v6152
	var v6153 int32
	_ = v6153
	var v6155 int32
	_ = v6155
	var v6156 int32
	_ = v6156
	var v6157 int32
	_ = v6157
	var v6178 int32
	_ = v6178
	var v6182 int32
	_ = v6182
	var v6183 int32
	_ = v6183
	var v6185 int32
	_ = v6185
	var v6186 int32
	_ = v6186
	var v6188 int32
	_ = v6188
	var v6194 int32
	_ = v6194
	var v6204 int32
	_ = v6204
	var v6237 int32
	_ = v6237
	var v6238 int32
	_ = v6238
	var v6239 int32
	_ = v6239
	var v6244 int32
	_ = v6244
	var v6245 int32
	_ = v6245
	var v6247 int32
	_ = v6247
	var v6248 int32
	_ = v6248
	var v6249 int32
	_ = v6249
	var v6250 int32
	_ = v6250
	var v6256 int32
	_ = v6256
	var v6302 int32
	_ = v6302
	var v6305 int32
	_ = v6305
	var v6309 int32
	_ = v6309
	var v6310 int32
	_ = v6310
	var v6314 int32
	_ = v6314
	var v6317 int32
	_ = v6317
	var v6318 int32
	_ = v6318
	var v6329 int32
	_ = v6329
	var v6368 int32
	_ = v6368
	var v6370 int32
	_ = v6370
	var v6371 int32
	_ = v6371
	var v6372 int32
	_ = v6372
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
	var v6383 int32
	_ = v6383
	var v6387 int32
	_ = v6387
	var v6389 int32
	_ = v6389
	var v6391 int32
	_ = v6391
	var v6433 int32
	_ = v6433
	var v6437 int32
	_ = v6437
	var v6438 int32
	_ = v6438
	var v6439 int32
	_ = v6439
	var v6440 int32
	_ = v6440
	var v6441 int32
	_ = v6441
	var v6443 int32
	_ = v6443
	var v6445 int32
	_ = v6445
	var v6446 int32
	_ = v6446
	var v6452 int32
	_ = v6452
	var v6456 int32
	_ = v6456
	var v6457 int32
	_ = v6457
	var v6459 int32
	_ = v6459
	var v6460 int32
	_ = v6460
	var v6468 int32
	_ = v6468
	var v6510 int32
	_ = v6510
	var v6511 int32
	_ = v6511
	var v6512 int32
	_ = v6512
	var v6513 int32
	_ = v6513
	var v6514 int32
	_ = v6514
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6517 int32
	_ = v6517
	var v6518 int32
	_ = v6518
	var v6519 float64
	_ = v6519
	var v6520 int64
	_ = v6520
	var v6521 int32
	_ = v6521
	var v6523 int32
	_ = v6523
	var v6524 int32
	_ = v6524
	var v6527 int32
	_ = v6527
	var v6545 int32
	_ = v6545
	var v6547 float64
	_ = v6547
	var v6549 float64
	_ = v6549
	var v6551 float64
	_ = v6551
	var v6553 int32
	_ = v6553
	var v6554 int32
	_ = v6554
	var v6556 int32
	_ = v6556
	var v6558 int32
	_ = v6558
	var v6560 int32
	_ = v6560
	var v6562 int32
	_ = v6562
	var v6563 int32
	_ = v6563
	var v6564 int32
	_ = v6564
	var v6565 int32
	_ = v6565
	var v6569 int32
	_ = v6569
	var v6572 int32
	_ = v6572
	var v6575 int32
	_ = v6575
	var v6579 int32
	_ = v6579
	var v6584 int32
	_ = v6584
	var v6621 int32
	_ = v6621
	var v6625 int32
	_ = v6625
	var v6626 int32
	_ = v6626
	var v6627 int32
	_ = v6627
	var v6628 int32
	_ = v6628
	var v6629 int32
	_ = v6629
	var v6631 int32
	_ = v6631
	var v6633 int32
	_ = v6633
	var v6634 int32
	_ = v6634
	var v6640 int32
	_ = v6640
	var v6644 int32
	_ = v6644
	var v6645 int32
	_ = v6645
	var v6647 int32
	_ = v6647
	var v6648 int32
	_ = v6648
	var v6661 int32
	_ = v6661
	var v6698 int32
	_ = v6698
	var v6699 int32
	_ = v6699
	var v6700 int32
	_ = v6700
	var v6701 int32
	_ = v6701
	var v6702 int32
	_ = v6702
	var v6703 int32
	_ = v6703
	var v6704 int32
	_ = v6704
	var v6705 int32
	_ = v6705
	var v6706 int32
	_ = v6706
	var v6707 int32
	_ = v6707
	var v6708 int32
	_ = v6708
	var v6709 int32
	_ = v6709
	var v6710 int32
	_ = v6710
	var v6711 int32
	_ = v6711
	var v6712 int32
	_ = v6712
	var v6713 int32
	_ = v6713
	var v6714 int32
	_ = v6714
	var v6715 int32
	_ = v6715
	var v6716 float64
	_ = v6716
	var v6717 int64
	_ = v6717
	var v6719 int32
	_ = v6719
	var v6720 int32
	_ = v6720
	var v6721 int32
	_ = v6721
	var v6740 int32
	_ = v6740
	var v6742 float64
	_ = v6742
	var v6744 float64
	_ = v6744
	var v6746 float64
	_ = v6746
	var v6748 int32
	_ = v6748
	var v6749 int32
	_ = v6749
	var v6751 int32
	_ = v6751
	var v6753 int32
	_ = v6753
	var v6755 int32
	_ = v6755
	var v6757 int32
	_ = v6757
	var v6758 int32
	_ = v6758
	var v6759 int32
	_ = v6759
	var v6760 int32
	_ = v6760
	var v6761 int32
	_ = v6761
	var v6765 int32
	_ = v6765
	var v6768 int32
	_ = v6768
	var v6771 int32
	_ = v6771
	var v6772 int32
	_ = v6772
	var v6774 int32
	_ = v6774
	var v6817 int32
	_ = v6817
	var v6821 int32
	_ = v6821
	var v6822 int32
	_ = v6822
	var v6823 int32
	_ = v6823
	var v6824 int32
	_ = v6824
	var v6825 int32
	_ = v6825
	var v6827 int32
	_ = v6827
	var v6829 int32
	_ = v6829
	var v6830 int32
	_ = v6830
	var v6836 int32
	_ = v6836
	var v6840 int32
	_ = v6840
	var v6841 int32
	_ = v6841
	var v6843 int32
	_ = v6843
	var v6844 int32
	_ = v6844
	var v6848 int32
	_ = v6848
	var v6894 int32
	_ = v6894
	var v6895 int32
	_ = v6895
	var v6896 int32
	_ = v6896
	var v6897 int32
	_ = v6897
	var v6898 int32
	_ = v6898
	var v6899 int32
	_ = v6899
	var v6900 int32
	_ = v6900
	var v6901 int32
	_ = v6901
	var v6902 int32
	_ = v6902
	var v6903 int32
	_ = v6903
	var v6904 int32
	_ = v6904
	var v6905 int32
	_ = v6905
	var v6906 int32
	_ = v6906
	var v6907 int32
	_ = v6907
	var v6908 int32
	_ = v6908
	var v6909 int32
	_ = v6909
	var v6911 int32
	_ = v6911
	var v6912 int32
	_ = v6912
	var v6924 int32
	_ = v6924
	var v6926 float64
	_ = v6926
	var v6928 float64
	_ = v6928
	var v6930 float64
	_ = v6930
	var v6932 int32
	_ = v6932
	var v6933 int32
	_ = v6933
	var v6935 int32
	_ = v6935
	var v6937 int32
	_ = v6937
	var v6939 int32
	_ = v6939
	var v6942 int32
	_ = v6942
	var v6943 int32
	_ = v6943
	var v6944 int32
	_ = v6944
	var v6945 int32
	_ = v6945
	var v6947 int32
	_ = v6947
	var v6948 int32
	_ = v6948
	var v6949 int32
	_ = v6949
	var v6952 int32
	_ = v6952
	var v6959 int32
	_ = v6959
	var v6960 int32
	_ = v6960
	var v6961 int32
	_ = v6961
	var v6962 int32
	_ = v6962
	var v6974 int32
	_ = v6974
	var v6975 int32
	_ = v6975
	var v6976 int32
	_ = v6976
	var v6977 int32
	_ = v6977
	var v6978 int32
	_ = v6978
	var v6979 int32
	_ = v6979
	var v6980 int32
	_ = v6980
	var v6982 int32
	_ = v6982
	var v6983 int32
	_ = v6983
	var v6986 int32
	_ = v6986
	var v6988 int32
	_ = v6988
	var v6999 int32
	_ = v6999
	var v7001 float64
	_ = v7001
	var v7003 float64
	_ = v7003
	var v7005 float64
	_ = v7005
	var v7007 int32
	_ = v7007
	var v7008 int32
	_ = v7008
	var v7010 int32
	_ = v7010
	var v7012 int32
	_ = v7012
	var v7014 int32
	_ = v7014
	var v7017 int32
	_ = v7017
	var v7018 int32
	_ = v7018
	var v7019 int32
	_ = v7019
	var v7021 int32
	_ = v7021
	var v7022 int32
	_ = v7022
	var v7023 int32
	_ = v7023
	var v7026 int32
	_ = v7026
	var v7033 int32
	_ = v7033
	var v7034 int32
	_ = v7034
	var v7035 int32
	_ = v7035
	var v7036 int32
	_ = v7036
	var v7048 int32
	_ = v7048
	var v7049 int32
	_ = v7049
	var v7050 int32
	_ = v7050
	var v7051 int32
	_ = v7051
	var v7052 int32
	_ = v7052
	var v7053 int32
	_ = v7053
	var v7054 int32
	_ = v7054
	var v7056 int32
	_ = v7056
	var v7057 int32
	_ = v7057
	var v7060 int32
	_ = v7060
	var v7062 int32
	_ = v7062
	var v7064 int32
	_ = v7064
	var v7070 int32
	_ = v7070
	var v7079 int32
	_ = v7079
	var v7081 float64
	_ = v7081
	var v7083 float64
	_ = v7083
	var v7085 float64
	_ = v7085
	var v7087 int32
	_ = v7087
	var v7088 int32
	_ = v7088
	var v7090 int32
	_ = v7090
	var v7092 int32
	_ = v7092
	var v7094 int32
	_ = v7094
	var v7095 int32
	_ = v7095
	var v7097 int32
	_ = v7097
	var v7098 int32
	_ = v7098
	var v7099 int32
	_ = v7099
	var v7100 int32
	_ = v7100
	var v7101 int32
	_ = v7101
	var v7104 int32
	_ = v7104
	var v7107 int32
	_ = v7107
	var v7110 int32
	_ = v7110
	var v7111 int32
	_ = v7111
	var v7113 int32
	_ = v7113
	var v7156 int32
	_ = v7156
	var v7160 int32
	_ = v7160
	var v7161 int32
	_ = v7161
	var v7162 int32
	_ = v7162
	var v7163 int32
	_ = v7163
	var v7164 int32
	_ = v7164
	var v7166 int32
	_ = v7166
	var v7168 int32
	_ = v7168
	var v7169 int32
	_ = v7169
	var v7175 int32
	_ = v7175
	var v7179 int32
	_ = v7179
	var v7180 int32
	_ = v7180
	var v7182 int32
	_ = v7182
	var v7183 int32
	_ = v7183
	var v7187 int32
	_ = v7187
	var v7233 int32
	_ = v7233
	var v7234 int32
	_ = v7234
	var v7235 int32
	_ = v7235
	var v7236 int32
	_ = v7236
	var v7238 int32
	_ = v7238
	var v7239 int32
	_ = v7239
	var v7240 int32
	_ = v7240
	var v7255 int32
	_ = v7255
	var v7257 float64
	_ = v7257
	var v7259 float64
	_ = v7259
	var v7261 float64
	_ = v7261
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7266 int32
	_ = v7266
	var v7268 int32
	_ = v7268
	var v7270 int32
	_ = v7270
	var v7271 int32
	_ = v7271
	var v7273 int32
	_ = v7273
	var v7276 int32
	_ = v7276
	var v7277 int32
	_ = v7277
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7285 int32
	_ = v7285
	var v7292 int32
	_ = v7292
	var v7293 int32
	_ = v7293
	var v7295 int32
	_ = v7295
	var v7296 int32
	_ = v7296
	var v7299 int32
	_ = v7299
	var v7300 int32
	_ = v7300
	var v7307 int32
	_ = v7307
	var v7308 int32
	_ = v7308
	var v7310 int32
	_ = v7310
	var v7311 int32
	_ = v7311
	var v7313 int32
	_ = v7313
	var v7314 int32
	_ = v7314
	var v7317 int32
	_ = v7317
	var v7320 int32
	_ = v7320
	var v7323 int32
	_ = v7323
	var v7334 int32
	_ = v7334
	var v7374 int32
	_ = v7374
	var v7375 int32
	_ = v7375
	var v7377 int32
	_ = v7377
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7382 int32
	_ = v7382
	var v7386 int32
	_ = v7386
	var v7390 int32
	_ = v7390
	var v7395 int32
	_ = v7395
	var v7396 int32
	_ = v7396
	var v7399 int32
	_ = v7399
	var v7400 int32
	_ = v7400
	var v7403 int32
	_ = v7403
	var v7451 int32
	_ = v7451
	var v7455 int32
	_ = v7455
	var v7456 int32
	_ = v7456
	var v7457 int32
	_ = v7457
	var v7458 int32
	_ = v7458
	var v7460 int32
	_ = v7460
	var v7461 int32
	_ = v7461
	var v7463 int32
	_ = v7463
	var v7464 int32
	_ = v7464
	var v7465 int32
	_ = v7465
	var v7468 int32
	_ = v7468
	var v7469 int32
	_ = v7469
	var v7470 int32
	_ = v7470
	var v7473 int32
	_ = v7473
	var v7477 int32
	_ = v7477
	var v7519 int32
	_ = v7519
	var v7520 int32
	_ = v7520
	var v7522 int32
	_ = v7522
	var v7523 int32
	_ = v7523
	var v7526 int32
	_ = v7526
	var v7529 int32
	_ = v7529
	var v7534 int32
	_ = v7534
	var v7537 int32
	_ = v7537
	var v7538 int32
	_ = v7538
	var v7592 int32
	_ = v7592
	var v7594 float64
	_ = v7594
	var v7596 float64
	_ = v7596
	var v7598 float64
	_ = v7598
	var v7600 int32
	_ = v7600
	var v7601 int32
	_ = v7601
	var v7603 int32
	_ = v7603
	var v7605 int32
	_ = v7605
	var v7610 int32
	_ = v7610
	var v7611 int32
	_ = v7611
	var v7621 int32
	_ = v7621
	var v7626 int32
	_ = v7626
	var v7678 int32
	_ = v7678
	var v7682 int32
	_ = v7682
	var v7687 int32
	_ = v7687
	var v7688 int32
	_ = v7688
	var v7691 int32
	_ = v7691
	var v7692 int32
	_ = v7692
	var v7693 int32
	_ = v7693
	var v7694 int32
	_ = v7694
	var v7695 int32
	_ = v7695
	var v7696 int32
	_ = v7696
	var v7699 int32
	_ = v7699
	var v7700 int32
	_ = v7700
	var v7701 int32
	_ = v7701
	var v7702 int32
	_ = v7702
	var v7703 int32
	_ = v7703
	var v7704 int32
	_ = v7704
	var v7708 int32
	_ = v7708
	var v7753 int32
	_ = v7753
	var v7756 int32
	_ = v7756
	var v7758 int32
	_ = v7758
	var v7762 int32
	_ = v7762
	var v7767 int32
	_ = v7767
	var v7770 int32
	_ = v7770
	var v7773 int32
	_ = v7773
	var v7775 int32
	_ = v7775
	var v7778 int32
	_ = v7778
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7787 int32
	_ = v7787
	var v7789 int32
	_ = v7789
	var v7791 int32
	_ = v7791
	var v7796 int32
	_ = v7796
	var v7800 int32
	_ = v7800
	var v7801 int32
	_ = v7801
	var v7802 int32
	_ = v7802
	var v7806 int32
	_ = v7806
	var v7807 int32
	_ = v7807
	var v7808 int32
	_ = v7808
	var v7809 int32
	_ = v7809
	var v7813 float64
	_ = v7813
	var v7814 float64
	_ = v7814
	var v7815 float64
	_ = v7815
	var v7816 int32
	_ = v7816
	var v7817 int32
	_ = v7817
	var v7818 int32
	_ = v7818
	var v7820 int32
	_ = v7820
	var v7821 int32
	_ = v7821
	var v7824 int32
	_ = v7824
	var v7825 int32
	_ = v7825
	var v7832 int32
	_ = v7832
	var v7833 int32
	_ = v7833
	var v7845 int32
	_ = v7845
	var v7847 float64
	_ = v7847
	var v7849 float64
	_ = v7849
	var v7851 float64
	_ = v7851
	var v7853 int32
	_ = v7853
	var v7854 int32
	_ = v7854
	var v7856 int32
	_ = v7856
	var v7858 int32
	_ = v7858
	var v7860 int32
	_ = v7860
	var v7863 int32
	_ = v7863
	var v7864 int32
	_ = v7864
	var v7866 int32
	_ = v7866
	var v7867 int32
	_ = v7867
	var v7870 int32
	_ = v7870
	var v7871 int32
	_ = v7871
	var v7877 int32
	_ = v7877
	var v7879 float64
	_ = v7879
	var v7881 float64
	_ = v7881
	var v7883 float64
	_ = v7883
	var v7885 int32
	_ = v7885
	var v7886 int32
	_ = v7886
	var v7888 int32
	_ = v7888
	var v7890 int32
	_ = v7890
	var v7892 int32
	_ = v7892
	var v7894 int32
	_ = v7894
	var v7895 int32
	_ = v7895
	var v7896 int32
	_ = v7896
	var v7897 int32
	_ = v7897
	var v7898 int32
	_ = v7898
	var v7902 int32
	_ = v7902
	var v7905 int32
	_ = v7905
	var v7908 int32
	_ = v7908
	var v7909 int32
	_ = v7909
	var v7911 int32
	_ = v7911
	var v7954 int32
	_ = v7954
	var v7958 int32
	_ = v7958
	var v7959 int32
	_ = v7959
	var v7960 int32
	_ = v7960
	var v7961 int32
	_ = v7961
	var v7962 int32
	_ = v7962
	var v7964 int32
	_ = v7964
	var v7966 int32
	_ = v7966
	var v7967 int32
	_ = v7967
	var v7973 int32
	_ = v7973
	var v7977 int32
	_ = v7977
	var v7978 int32
	_ = v7978
	var v7980 int32
	_ = v7980
	var v7981 int32
	_ = v7981
	var v7985 int32
	_ = v7985
	var v8032 int32
	_ = v8032
	var v8033 int32
	_ = v8033
	var v8034 int32
	_ = v8034
	var v8042 int32
	_ = v8042
	var v8044 float64
	_ = v8044
	var v8046 float64
	_ = v8046
	var v8048 float64
	_ = v8048
	var v8050 int32
	_ = v8050
	var v8051 int32
	_ = v8051
	var v8053 int32
	_ = v8053
	var v8055 int32
	_ = v8055
	var v8057 int32
	_ = v8057
	var v8060 int32
	_ = v8060
	var v8061 int32
	_ = v8061
	var v8062 int32
	_ = v8062
	var v8064 int32
	_ = v8064
	var v8065 int32
	_ = v8065
	var v8066 int32
	_ = v8066
	var v8071 int32
	_ = v8071
	var v8073 int32
	_ = v8073
	var v8074 int32
	_ = v8074
	var v8075 int32
	_ = v8075
	var v8078 int32
	_ = v8078
	var v8083 int32
	_ = v8083
	var v8086 int32
	_ = v8086
	var v8090 int32
	_ = v8090
	var v8092 int32
	_ = v8092
	var v8094 int32
	_ = v8094
	var v8095 int32
	_ = v8095
	var v8096 int32
	_ = v8096
	var v8097 int32
	_ = v8097
	var v8102 int32
	_ = v8102
	var v8103 int32
	_ = v8103
	var v8106 int32
	_ = v8106
	var v8109 int32
	_ = v8109
	var v8110 int32
	_ = v8110
	var v8118 int32
	_ = v8118
	var v8155 int32
	_ = v8155
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
	var v8167 int32
	_ = v8167
	var v8168 int32
	_ = v8168
	var v8174 int32
	_ = v8174
	var v8178 int32
	_ = v8178
	var v8179 int32
	_ = v8179
	var v8181 int32
	_ = v8181
	var v8182 int32
	_ = v8182
	var v8184 int32
	_ = v8184
	var v8186 int32
	_ = v8186
	var v8187 int32
	_ = v8187
	var v8188 int32
	_ = v8188
	var v8189 int32
	_ = v8189
	var v8193 int32
	_ = v8193
	var v8196 int32
	_ = v8196
	var v8199 int32
	_ = v8199
	var v8200 int32
	_ = v8200
	var v8208 int32
	_ = v8208
	var v8245 int32
	_ = v8245
	var v8249 int32
	_ = v8249
	var v8250 int32
	_ = v8250
	var v8251 int32
	_ = v8251
	var v8252 int32
	_ = v8252
	var v8253 int32
	_ = v8253
	var v8255 int32
	_ = v8255
	var v8257 int32
	_ = v8257
	var v8258 int32
	_ = v8258
	var v8264 int32
	_ = v8264
	var v8268 int32
	_ = v8268
	var v8269 int32
	_ = v8269
	var v8271 int32
	_ = v8271
	var v8272 int32
	_ = v8272
	var v8277 int32
	_ = v8277
	var v8322 int32
	_ = v8322
	var v8323 int32
	_ = v8323
	var v8324 int32
	_ = v8324
	var v8327 int32
	_ = v8327
	var v8330 int32
	_ = v8330
	var v8376 float64
	_ = v8376
	var v8378 float64
	_ = v8378
	var v8380 float64
	_ = v8380
	var v8382 int32
	_ = v8382
	var v8383 int32
	_ = v8383
	var v8386 int32
	_ = v8386
	var v8387 int32
	_ = v8387
	var v8388 int32
	_ = v8388
	var v8400 int32
	_ = v8400
	var v8402 float64
	_ = v8402
	var v8404 float64
	_ = v8404
	var v8406 float64
	_ = v8406
	var v8408 int32
	_ = v8408
	var v8409 int32
	_ = v8409
	var v8411 int32
	_ = v8411
	var v8413 int32
	_ = v8413
	var v8461 int32
	_ = v8461
	var v8463 int32
	_ = v8463
	var v8464 int32
	_ = v8464
	var v8467 int32
	_ = v8467
	var v8470 int32
	_ = v8470
	var v8518 int32
	_ = v8518
	var v8522 int32
	_ = v8522
	var v8523 int32
	_ = v8523
	var v8524 int32
	_ = v8524
	var v8525 int32
	_ = v8525
	var v8526 int32
	_ = v8526
	var v8527 int32
	_ = v8527
	var v8528 int64
	_ = v8528
	var v8529 int32
	_ = v8529
	var v8531 int32
	_ = v8531
	var v8532 int32
	_ = v8532
	var v8535 int32
	_ = v8535
	var v8536 int64
	_ = v8536
	var v8540 int32
	_ = v8540
	var v8548 int32
	_ = v8548
	var v8549 int32
	_ = v8549
	var v8551 int32
	_ = v8551
	var v8552 float64
	_ = v8552
	var v8554 float64
	_ = v8554
	var v8558 int32
	_ = v8558
	var v8559 int32
	_ = v8559
	var v8560 int32
	_ = v8560
	var v8564 int32
	_ = v8564
	var v8565 int32
	_ = v8565
	var v8567 int32
	_ = v8567
	var v8568 int32
	_ = v8568
	var v8570 int32
	_ = v8570
	var v8572 int32
	_ = v8572
	var v8573 int32
	_ = v8573
	var v8574 int32
	_ = v8574
	var v8575 int32
	_ = v8575
	var v8576 int32
	_ = v8576
	var v8578 int32
	_ = v8578
	var v8579 int32
	_ = v8579
	var v8581 int32
	_ = v8581
	var v8582 int32
	_ = v8582
	var v8583 int32
	_ = v8583
	var v8585 int32
	_ = v8585
	var v8586 int32
	_ = v8586
	var v8587 int32
	_ = v8587
	var v8588 int32
	_ = v8588
	var v8589 int32
	_ = v8589
	var v8592 int32
	_ = v8592
	var v8593 int32
	_ = v8593
	var v8596 int32
	_ = v8596
	var v8597 int32
	_ = v8597
	var v8598 int32
	_ = v8598
	var v8600 int32
	_ = v8600
	var v8602 int32
	_ = v8602
	var v8603 int32
	_ = v8603
	var v8606 int32
	_ = v8606
	var v8609 int32
	_ = v8609
	var v8610 int32
	_ = v8610
	var v8611 int32
	_ = v8611
	var v8612 int32
	_ = v8612
	var v8613 int32
	_ = v8613
	var v8614 int32
	_ = v8614
	var v8616 int32
	_ = v8616
	var v8617 int32
	_ = v8617
	var v8618 int32
	_ = v8618
	var v8620 int32
	_ = v8620
	var v8621 int32
	_ = v8621
	var v8622 int32
	_ = v8622
	var v8628 int32
	_ = v8628
	var v8630 int32
	_ = v8630
	var v8632 int32
	_ = v8632
	var v8638 int32
	_ = v8638
	var v8639 int32
	_ = v8639
	var v8641 int32
	_ = v8641
	var v8642 int32
	_ = v8642
	var v8643 int32
	_ = v8643
	var v8646 int32
	_ = v8646
	var v8651 int32
	_ = v8651
	var v8652 int32
	_ = v8652
	var v8702 int32
	_ = v8702
	var v8703 int32
	_ = v8703
	var v8704 int32
	_ = v8704
	var v8708 int32
	_ = v8708
	var v8711 int32
	_ = v8711
	var v8713 int32
	_ = v8713
	var v8715 int32
	_ = v8715
	var v8717 int32
	_ = v8717
	var v8761 int32
	_ = v8761
	var v8765 int32
	_ = v8765
	var v8766 int32
	_ = v8766
	var v8767 int32
	_ = v8767
	var v8768 int32
	_ = v8768
	var v8769 int32
	_ = v8769
	var v8771 int32
	_ = v8771
	var v8773 int32
	_ = v8773
	var v8774 int32
	_ = v8774
	var v8780 int32
	_ = v8780
	var v8784 int32
	_ = v8784
	var v8785 int32
	_ = v8785
	var v8787 int32
	_ = v8787
	var v8788 int32
	_ = v8788
	var v8792 int32
	_ = v8792
	var v8838 int32
	_ = v8838
	var v8839 int32
	_ = v8839
	var v8841 int32
	_ = v8841
	var v8842 int32
	_ = v8842
	var v8852 int32
	_ = v8852
	var v8856 int32
	_ = v8856
	var v8857 int32
	_ = v8857
	var v8861 int32
	_ = v8861
	var v8865 int32
	_ = v8865
	var v8867 float64
	_ = v8867
	var v8869 float64
	_ = v8869
	var v8871 float64
	_ = v8871
	var v8873 int32
	_ = v8873
	var v8874 int32
	_ = v8874
	var v8876 int32
	_ = v8876
	var v8878 int32
	_ = v8878
	var v8880 int32
	_ = v8880
	var v8882 int32
	_ = v8882
	var v8884 int32
	_ = v8884
	var v8885 int32
	_ = v8885
	var v8889 int32
	_ = v8889
	var v8892 int32
	_ = v8892
	var v8893 int32
	_ = v8893
	var v8895 int32
	_ = v8895
	var v8897 int32
	_ = v8897
	var v8941 int32
	_ = v8941
	var v8945 int32
	_ = v8945
	var v8946 int32
	_ = v8946
	var v8947 int32
	_ = v8947
	var v8948 int32
	_ = v8948
	var v8949 int32
	_ = v8949
	var v8951 int32
	_ = v8951
	var v8953 int32
	_ = v8953
	var v8954 int32
	_ = v8954
	var v8960 int32
	_ = v8960
	var v8964 int32
	_ = v8964
	var v8965 int32
	_ = v8965
	var v8967 int32
	_ = v8967
	var v8968 int32
	_ = v8968
	var v8970 int32
	_ = v8970
	var v9018 int32
	_ = v9018
	var v9019 int32
	_ = v9019
	var v9020 int32
	_ = v9020
	var v9021 int32
	_ = v9021
	var v9023 int32
	_ = v9023
	var v9024 int32
	_ = v9024
	var v9034 int32
	_ = v9034
	var v9038 int32
	_ = v9038
	var v9040 int32
	_ = v9040
	var v9043 int32
	_ = v9043
	var v9045 int32
	_ = v9045
	var v9047 float64
	_ = v9047
	var v9049 float64
	_ = v9049
	var v9051 float64
	_ = v9051
	var v9053 int32
	_ = v9053
	var v9054 int32
	_ = v9054
	var v9056 int32
	_ = v9056
	var v9058 int32
	_ = v9058
	var v9060 int32
	_ = v9060
	var v9061 int32
	_ = v9061
	var v9063 int32
	_ = v9063
	var v9064 int32
	_ = v9064
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9072 int32
	_ = v9072
	var v9073 int32
	_ = v9073
	var v9076 int32
	_ = v9076
	var v9081 int32
	_ = v9081
	var v9083 int32
	_ = v9083
	var v9087 int32
	_ = v9087
	var v9125 int32
	_ = v9125
	var v9129 int32
	_ = v9129
	var v9130 int32
	_ = v9130
	var v9131 int32
	_ = v9131
	var v9132 int32
	_ = v9132
	var v9133 int32
	_ = v9133
	var v9135 int32
	_ = v9135
	var v9137 int32
	_ = v9137
	var v9138 int32
	_ = v9138
	var v9144 int32
	_ = v9144
	var v9148 int32
	_ = v9148
	var v9149 int32
	_ = v9149
	var v9151 int32
	_ = v9151
	var v9152 int32
	_ = v9152
	var v9156 int32
	_ = v9156
	var v9205 int32
	_ = v9205
	var v9217 int32
	_ = v9217
	var v9255 int32
	_ = v9255
	var v9256 int32
	_ = v9256
	var v9257 int32
	_ = v9257
	var v9258 int32
	_ = v9258
	var v9260 float64
	_ = v9260
	var v9262 float64
	_ = v9262
	var v9264 float64
	_ = v9264
	var v9266 int32
	_ = v9266
	var v9267 int32
	_ = v9267
	var v9269 int32
	_ = v9269
	var v9271 int32
	_ = v9271
	var v9272 int32
	_ = v9272
	var v9278 int32
	_ = v9278
	var v9280 int32
	_ = v9280
	var v9282 int32
	_ = v9282
	var v9283 int32
	_ = v9283
	var v9289 int32
	_ = v9289
	var v9296 int32
	_ = v9296
	var v9297 int32
	_ = v9297
	var v9298 int32
	_ = v9298
	var v9299 int32
	_ = v9299
	var v9300 int32
	_ = v9300
	var v9301 int32
	_ = v9301
	var v9302 int32
	_ = v9302
	var v9305 int32
	_ = v9305
	var v9315 int32
	_ = v9315
	var v9319 int32
	_ = v9319
	var v9357 int32
	_ = v9357
	var v9361 int32
	_ = v9361
	var v9363 int32
	_ = v9363
	var v9364 int32
	_ = v9364
	var v9365 int32
	_ = v9365
	var v9366 int32
	_ = v9366
	var v9367 int32
	_ = v9367
	var v9379 int32
	_ = v9379
	var v9380 int32
	_ = v9380
	var v9381 int32
	_ = v9381
	var v9382 int32
	_ = v9382
	var v9383 int32
	_ = v9383
	var v9385 int32
	_ = v9385
	var v9393 int32
	_ = v9393
	var v9394 int32
	_ = v9394
	var v9395 int32
	_ = v9395
	var v9398 int32
	_ = v9398
	var v9399 int32
	_ = v9399
	var v9401 int32
	_ = v9401
	var v9402 int32
	_ = v9402
	var v9404 int32
	_ = v9404
	var v9406 int32
	_ = v9406
	var v9409 int32
	_ = v9409
	var v9410 int32
	_ = v9410
	var v9411 int32
	_ = v9411
	var v9416 int32
	_ = v9416
	var v9417 int32
	_ = v9417
	var v9418 int32
	_ = v9418
	var v9421 int32
	_ = v9421
	var v9422 int32
	_ = v9422
	var v9423 int32
	_ = v9423
	var v9426 int32
	_ = v9426
	var v9427 int32
	_ = v9427
	var v9429 int32
	_ = v9429
	var v9434 int32
	_ = v9434
	var v9447 int32
	_ = v9447
	var v9448 int32
	_ = v9448
	var v9450 int32
	_ = v9450
	var v9468 int32
	_ = v9468
	var v9471 int32
	_ = v9471
	var v9472 int32
	_ = v9472
	var v9475 int32
	_ = v9475
	var v9476 int32
	_ = v9476
	var v9481 int32
	_ = v9481
	var v9488 int32
	_ = v9488
	var v9492 int32
	_ = v9492
	var v9498 int32
	_ = v9498
	var v9502 int32
	_ = v9502
	var v9506 int32
	_ = v9506
	var v9510 int32
	_ = v9510
	var v9516 int32
	_ = v9516
	var v9528 int32
	_ = v9528
	var v9529 int32
	_ = v9529
	var v9530 int32
	_ = v9530
	var v9531 int32
	_ = v9531
	var v9533 int32
	_ = v9533
	var v9536 int32
	_ = v9536
	var v9540 int32
	_ = v9540
	var v9541 int32
	_ = v9541
	var v9544 int32
	_ = v9544
	var v9546 int32
	_ = v9546
	var v9559 int32
	_ = v9559
	var v9560 float64
	_ = v9560
	var v9561 float64
	_ = v9561
	var v9562 float64
	_ = v9562
	var v9563 int32
	_ = v9563
	var v9565 int32
	_ = v9565
	var v9566 float64
	_ = v9566
	var v9568 int32
	_ = v9568
	var v9569 float64
	_ = v9569
	var v9571 float64
	_ = v9571
	var v9573 float64
	_ = v9573
	var v9575 int32
	_ = v9575
	var v9576 int32
	_ = v9576
	var v9579 int32
	_ = v9579
	var v9583 int32
	_ = v9583
	var v9584 int32
	_ = v9584
	var v9587 int32
	_ = v9587
	var v9589 int32
	_ = v9589
	var v9590 int32
	_ = v9590
	var v9591 int32
	_ = v9591
	var v9597 int32
	_ = v9597
	var v9602 int32
	_ = v9602
	var v9604 int32
	_ = v9604
	var v9607 int32
	_ = v9607
	var v9608 float64
	_ = v9608
	var v9609 float64
	_ = v9609
	var v9610 int32
	_ = v9610
	var v9613 int32
	_ = v9613
	var v9614 float64
	_ = v9614
	var v9616 int32
	_ = v9616
	var v9617 int32
	_ = v9617
	var v9618 int32
	_ = v9618
	var v9623 float64
	_ = v9623
	var v9626 int32
	_ = v9626
	var v9627 float64
	_ = v9627
	var v9633 float64
	_ = v9633
	var v9639 float64
	_ = v9639
	var v9641 float64
	_ = v9641
	var v9643 float64
	_ = v9643
	var v9645 int32
	_ = v9645
	var v9646 int32
	_ = v9646
	var v9649 int32
	_ = v9649
	var v9651 int32
	_ = v9651
	var v9658 int32
	_ = v9658
	var v9659 int32
	_ = v9659
	var v9661 int32
	_ = v9661
	var v9662 int32
	_ = v9662
	var v9674 int32
	_ = v9674
	var v9715 int32
	_ = v9715
	var v9718 int32
	_ = v9718
	var v9720 int32
	_ = v9720
	var v9721 int32
	_ = v9721
	var v9724 int32
	_ = v9724
	var v9725 int32
	_ = v9725
	var v9726 int32
	_ = v9726
	var v9736 int32
	_ = v9736
	var v9737 int32
	_ = v9737
	var v9738 int32
	_ = v9738
	var v9739 int32
	_ = v9739
	var v9740 int32
	_ = v9740
	var v9741 int32
	_ = v9741
	var v9745 int32
	_ = v9745
	var v9749 int32
	_ = v9749
	var v9754 int32
	_ = v9754
	var v9755 int32
	_ = v9755
	var v9756 int32
	_ = v9756
	var v9760 int32
	_ = v9760
	var v9761 int32
	_ = v9761
	var v9764 int32
	_ = v9764
	var v9768 int32
	_ = v9768
	var v9770 int32
	_ = v9770
	var v9771 int32
	_ = v9771
	var v9813 int32
	_ = v9813
	var v9817 int32
	_ = v9817
	var v9818 int32
	_ = v9818
	var v9819 int32
	_ = v9819
	var v9820 int32
	_ = v9820
	var v9821 int32
	_ = v9821
	var v9823 int32
	_ = v9823
	var v9825 int32
	_ = v9825
	var v9826 int32
	_ = v9826
	var v9832 int32
	_ = v9832
	var v9836 int32
	_ = v9836
	var v9837 int32
	_ = v9837
	var v9839 int32
	_ = v9839
	var v9840 int32
	_ = v9840
	var v9844 int32
	_ = v9844
	var v9893 int32
	_ = v9893
	var v9901 int32
	_ = v9901
	var v9943 int32
	_ = v9943
	var v9944 int32
	_ = v9944
	var v9945 int32
	_ = v9945
	var v9946 int32
	_ = v9946
	var v9956 int32
	_ = v9956
	var v9959 int32
	_ = v9959
	var v9961 int32
	_ = v9961
	var v9962 int32
	_ = v9962
	var v9968 int32
	_ = v9968
	var v9969 int32
	_ = v9969
	var v9970 int32
	_ = v9970
	var v9972 int32
	_ = v9972
	var v9973 int32
	_ = v9973
	var v9983 int32
	_ = v9983
	var v9987 int32
	_ = v9987
	var v9989 int32
	_ = v9989
	var v9992 int32
	_ = v9992
	var v9994 int32
	_ = v9994
	var v9996 float64
	_ = v9996
	var v9998 float64
	_ = v9998
	var v10000 float64
	_ = v10000
	var v10002 int32
	_ = v10002
	var v10003 int32
	_ = v10003
	var v10005 int32
	_ = v10005
	var v10007 int32
	_ = v10007
	var v10010 int32
	_ = v10010
	var v10011 int32
	_ = v10011
	var v10019 int32
	_ = v10019
	var v10021 int32
	_ = v10021
	var v10023 int32
	_ = v10023
	var v10024 int32
	_ = v10024
	var v10037 int32
	_ = v10037
	var v10038 int32
	_ = v10038
	var v10039 int32
	_ = v10039
	var v10040 int32
	_ = v10040
	var v10042 int32
	_ = v10042
	var v10044 int32
	_ = v10044
	var v10046 int32
	_ = v10046
	var v10049 int32
	_ = v10049
	var v10050 int32
	_ = v10050
	var v10053 int32
	_ = v10053
	var v10057 int32
	_ = v10057
	var v10058 int32
	_ = v10058
	var v10059 int32
	_ = v10059
	var v10062 int32
	_ = v10062
	var v10072 int32
	_ = v10072
	var v10075 int32
	_ = v10075
	var v10078 int32
	_ = v10078
	var v10114 int32
	_ = v10114
	var v10118 int32
	_ = v10118
	var v10120 int32
	_ = v10120
	var v10121 int32
	_ = v10121
	var v10124 int32
	_ = v10124
	var v10125 int32
	_ = v10125
	var v10126 int32
	_ = v10126
	var v10138 int32
	_ = v10138
	var v10139 int32
	_ = v10139
	var v10140 int32
	_ = v10140
	var v10141 int32
	_ = v10141
	var v10143 int32
	_ = v10143
	var v10151 int32
	_ = v10151
	var v10152 int32
	_ = v10152
	var v10153 int32
	_ = v10153
	var v10156 int32
	_ = v10156
	var v10157 int32
	_ = v10157
	var v10159 int32
	_ = v10159
	var v10160 int32
	_ = v10160
	var v10162 int32
	_ = v10162
	var v10164 int32
	_ = v10164
	var v10167 int32
	_ = v10167
	var v10168 int32
	_ = v10168
	var v10169 int32
	_ = v10169
	var v10174 int32
	_ = v10174
	var v10175 int32
	_ = v10175
	var v10176 int32
	_ = v10176
	var v10179 int32
	_ = v10179
	var v10180 int32
	_ = v10180
	var v10181 int32
	_ = v10181
	var v10184 int32
	_ = v10184
	var v10185 int32
	_ = v10185
	var v10187 int32
	_ = v10187
	var v10192 int32
	_ = v10192
	var v10205 int32
	_ = v10205
	var v10206 int32
	_ = v10206
	var v10208 int32
	_ = v10208
	var v10226 int32
	_ = v10226
	var v10229 int32
	_ = v10229
	var v10230 int32
	_ = v10230
	var v10233 int32
	_ = v10233
	var v10234 int32
	_ = v10234
	var v10239 int32
	_ = v10239
	var v10246 int32
	_ = v10246
	var v10250 int32
	_ = v10250
	var v10256 int32
	_ = v10256
	var v10260 int32
	_ = v10260
	var v10264 int32
	_ = v10264
	var v10268 int32
	_ = v10268
	var v10274 int32
	_ = v10274
	var v10286 int32
	_ = v10286
	var v10287 int32
	_ = v10287
	var v10288 int32
	_ = v10288
	var v10289 int32
	_ = v10289
	var v10291 int32
	_ = v10291
	var v10294 int32
	_ = v10294
	var v10298 int32
	_ = v10298
	var v10299 int32
	_ = v10299
	var v10302 int32
	_ = v10302
	var v10304 int32
	_ = v10304
	var v10317 int32
	_ = v10317
	var v10318 float64
	_ = v10318
	var v10319 float64
	_ = v10319
	var v10320 float64
	_ = v10320
	var v10321 int32
	_ = v10321
	var v10323 int32
	_ = v10323
	var v10324 float64
	_ = v10324
	var v10326 int32
	_ = v10326
	var v10327 float64
	_ = v10327
	var v10329 float64
	_ = v10329
	var v10331 float64
	_ = v10331
	var v10333 int32
	_ = v10333
	var v10334 int32
	_ = v10334
	var v10337 int32
	_ = v10337
	var v10341 int32
	_ = v10341
	var v10342 int32
	_ = v10342
	var v10345 int32
	_ = v10345
	var v10347 int32
	_ = v10347
	var v10348 int32
	_ = v10348
	var v10349 int32
	_ = v10349
	var v10355 int32
	_ = v10355
	var v10360 int32
	_ = v10360
	var v10362 int32
	_ = v10362
	var v10365 int32
	_ = v10365
	var v10366 float64
	_ = v10366
	var v10367 float64
	_ = v10367
	var v10368 int32
	_ = v10368
	var v10371 int32
	_ = v10371
	var v10372 float64
	_ = v10372
	var v10374 int32
	_ = v10374
	var v10375 int32
	_ = v10375
	var v10376 int32
	_ = v10376
	var v10381 float64
	_ = v10381
	var v10384 int32
	_ = v10384
	var v10385 float64
	_ = v10385
	var v10391 float64
	_ = v10391
	var v10397 float64
	_ = v10397
	var v10399 float64
	_ = v10399
	var v10401 float64
	_ = v10401
	var v10403 int32
	_ = v10403
	var v10404 int32
	_ = v10404
	var v10407 int32
	_ = v10407
	var v10409 int32
	_ = v10409
	var v10418 int32
	_ = v10418
	var v10419 int32
	_ = v10419
	var v10421 int32
	_ = v10421
	var v10422 int32
	_ = v10422
	var v10423 int32
	_ = v10423
	var v10425 int32
	_ = v10425
	var v10426 int32
	_ = v10426
	var v10431 int32
	_ = v10431
	var v10435 int32
	_ = v10435
	var v10440 int32
	_ = v10440
	var v10450 int32
	_ = v10450
	var v10453 int32
	_ = v10453
	var v10457 int32
	_ = v10457
	var v10492 int32
	_ = v10492
	var v10495 int32
	_ = v10495
	var v10497 int32
	_ = v10497
	var v10498 int32
	_ = v10498
	var v10499 int32
	_ = v10499
	var v10500 int32
	_ = v10500
	var v10502 int32
	_ = v10502
	var v10503 int32
	_ = v10503
	var v10504 int32
	_ = v10504
	var v10505 int32
	_ = v10505
	var v10506 int32
	_ = v10506
	var v10507 int32
	_ = v10507
	var v10508 int32
	_ = v10508
	var v10511 int32
	_ = v10511
	var v10512 int32
	_ = v10512
	var v10513 int32
	_ = v10513
	var v10519 int32
	_ = v10519
	var v10521 int32
	_ = v10521
	var v10523 float64
	_ = v10523
	var v10525 float64
	_ = v10525
	var v10527 float64
	_ = v10527
	var v10529 int32
	_ = v10529
	var v10530 int32
	_ = v10530
	var v10532 int32
	_ = v10532
	var v10534 int32
	_ = v10534
	var v10541 int32
	_ = v10541
	var v10542 int32
	_ = v10542
	var v10543 int32
	_ = v10543
	var v10544 int32
	_ = v10544
	var v10545 int32
	_ = v10545
	var v10546 int32
	_ = v10546
	var v10547 int32
	_ = v10547
	var v10548 int32
	_ = v10548
	var v10549 int32
	_ = v10549
	var v10553 int32
	_ = v10553
	var v10556 int32
	_ = v10556
	var v10559 int32
	_ = v10559
	var v10560 int32
	_ = v10560
	var v10562 int32
	_ = v10562
	var v10605 int32
	_ = v10605
	var v10609 int32
	_ = v10609
	var v10610 int32
	_ = v10610
	var v10611 int32
	_ = v10611
	var v10612 int32
	_ = v10612
	var v10613 int32
	_ = v10613
	var v10615 int32
	_ = v10615
	var v10617 int32
	_ = v10617
	var v10618 int32
	_ = v10618
	var v10624 int32
	_ = v10624
	var v10628 int32
	_ = v10628
	var v10629 int32
	_ = v10629
	var v10631 int32
	_ = v10631
	var v10632 int32
	_ = v10632
	var v10636 int32
	_ = v10636
	var v10682 int32
	_ = v10682
	var v10683 int32
	_ = v10683
	var v10684 int32
	_ = v10684
	var v10685 int32
	_ = v10685
	var v10686 int32
	_ = v10686
	var v10687 int32
	_ = v10687
	var v10688 int32
	_ = v10688
	var v10690 int32
	_ = v10690
	var v10692 int32
	_ = v10692
	var v10693 int32
	_ = v10693
	var v10694 int32
	_ = v10694
	var v10695 int32
	_ = v10695
	var v10696 int32
	_ = v10696
	var v10697 int32
	_ = v10697
	var v10698 int32
	_ = v10698
	var v10699 int32
	_ = v10699
	var v10700 int32
	_ = v10700
	var v10701 int32
	_ = v10701
	var v10702 int32
	_ = v10702
	var v10703 int32
	_ = v10703
	var v10705 int32
	_ = v10705
	var v10707 int32
	_ = v10707
	var v10708 int32
	_ = v10708
	var v10709 int32
	_ = v10709
	var v10711 int32
	_ = v10711
	var v10713 int32
	_ = v10713
	var v10714 int32
	_ = v10714
	var v10716 int32
	_ = v10716
	var v10720 int32
	_ = v10720
	var v10721 int32
	_ = v10721
	var v10727 int32
	_ = v10727
	var v10729 int32
	_ = v10729
	var v10730 int32
	_ = v10730
	var v10735 int32
	_ = v10735
	var v10736 int32
	_ = v10736
	var v10739 int32
	_ = v10739
	var v10740 int32
	_ = v10740
	var v10741 int32
	_ = v10741
	var v10743 int32
	_ = v10743
	var v10744 int32
	_ = v10744
	var v10745 int32
	_ = v10745
	var v10747 int32
	_ = v10747
	var v10750 int32
	_ = v10750
	var v10760 int32
	_ = v10760
	var v10763 int32
	_ = v10763
	var v10766 int32
	_ = v10766
	var v10770 int32
	_ = v10770
	var v10773 int32
	_ = v10773
	var v10774 int32
	_ = v10774
	var v10778 int32
	_ = v10778
	var v10785 int32
	_ = v10785
	var v10787 int32
	_ = v10787
	var v10795 int32
	_ = v10795
	var v10796 int32
	_ = v10796
	var v10809 int32
	_ = v10809
	var v10815 int32
	_ = v10815
	var v10816 int32
	_ = v10816
	var v10860 int32
	_ = v10860
	var v10861 int32
	_ = v10861
	var v10864 int32
	_ = v10864
	var v10867 int32
	_ = v10867
	var v10868 int32
	_ = v10868
	var v10869 int32
	_ = v10869
	var v10877 int32
	_ = v10877
	var v10879 int32
	_ = v10879
	var v10880 int32
	_ = v10880
	var v10883 int32
	_ = v10883
	var v10887 int32
	_ = v10887
	var v10890 int32
	_ = v10890
	var v10892 int32
	_ = v10892
	var v10895 int32
	_ = v10895
	var v10902 int32
	_ = v10902
	var v10904 int32
	_ = v10904
	var v10912 int32
	_ = v10912
	var v10913 int32
	_ = v10913
	var v10926 int32
	_ = v10926
	var v10933 int32
	_ = v10933
	var v10977 int32
	_ = v10977
	var v10978 int32
	_ = v10978
	var v10979 int32
	_ = v10979
	var v10980 int32
	_ = v10980
	var v10981 int32
	_ = v10981
	var v10989 int32
	_ = v10989
	var v10992 int32
	_ = v10992
	var v10994 int32
	_ = v10994
	var v11033 int32
	_ = v11033
	var v11035 int32
	_ = v11035
	var v11039 int32
	_ = v11039
	var v11040 int32
	_ = v11040
	var v11041 int32
	_ = v11041
	var v11044 int32
	_ = v11044
	var v11045 int32
	_ = v11045
	var v11046 int32
	_ = v11046
	var v11047 int32
	_ = v11047
	var v11048 int32
	_ = v11048
	var v11049 int32
	_ = v11049
	var v11050 int32
	_ = v11050
	var v11053 int32
	_ = v11053
	var v11054 int32
	_ = v11054
	var v11055 int32
	_ = v11055
	var v11056 int32
	_ = v11056
	var v11065 int32
	_ = v11065
	var v11066 int32
	_ = v11066
	var v11068 int32
	_ = v11068
	var v11071 int32
	_ = v11071
	var v11072 int32
	_ = v11072
	var v11077 int32
	_ = v11077
	var v11084 int32
	_ = v11084
	var v11086 int32
	_ = v11086
	var v11088 int32
	_ = v11088
	var v11091 int32
	_ = v11091
	var v11093 int32
	_ = v11093
	var v11095 int32
	_ = v11095
	var v11102 int32
	_ = v11102
	var v11109 int32
	_ = v11109
	var v11112 int32
	_ = v11112
	var v11122 int32
	_ = v11122
	var v11123 int32
	_ = v11123
	var v11125 int32
	_ = v11125
	var v11128 int32
	_ = v11128
	var v11129 int32
	_ = v11129
	var v11134 int32
	_ = v11134
	var v11141 int32
	_ = v11141
	var v11143 int32
	_ = v11143
	var v11145 int32
	_ = v11145
	var v11146 int32
	_ = v11146
	var v11148 int32
	_ = v11148
	var v11150 int32
	_ = v11150
	var v11157 int32
	_ = v11157
	var v11160 int32
	_ = v11160
	var v11161 int32
	_ = v11161
	var v11162 int32
	_ = v11162
	var v11164 int32
	_ = v11164
	var v11165 int32
	_ = v11165
	var v11168 int32
	_ = v11168
	var v11169 int32
	_ = v11169
	var v11170 int32
	_ = v11170
	var v11172 int32
	_ = v11172
	var v11173 int32
	_ = v11173
	var v11174 int32
	_ = v11174
	var v11182 int32
	_ = v11182
	var v11185 int32
	_ = v11185
	var v11188 int32
	_ = v11188
	var v11192 int32
	_ = v11192
	var v11195 int32
	_ = v11195
	var v11196 int32
	_ = v11196
	var v11200 int32
	_ = v11200
	var v11207 int32
	_ = v11207
	var v11209 int32
	_ = v11209
	var v11217 int32
	_ = v11217
	var v11218 int32
	_ = v11218
	var v11231 int32
	_ = v11231
	var v11237 int32
	_ = v11237
	var v11239 int32
	_ = v11239
	var v11282 int32
	_ = v11282
	var v11284 int32
	_ = v11284
	var v11288 int32
	_ = v11288
	var v11291 int32
	_ = v11291
	var v11292 int32
	_ = v11292
	var v11293 int32
	_ = v11293
	var v11294 int32
	_ = v11294
	var v11296 int32
	_ = v11296
	var v11303 int32
	_ = v11303
	var v11305 int32
	_ = v11305
	var v11306 int32
	_ = v11306
	var v11309 int32
	_ = v11309
	var v11313 int32
	_ = v11313
	var v11316 int32
	_ = v11316
	var v11318 int32
	_ = v11318
	var v11321 int32
	_ = v11321
	var v11328 int32
	_ = v11328
	var v11330 int32
	_ = v11330
	var v11338 int32
	_ = v11338
	var v11339 int32
	_ = v11339
	var v11352 int32
	_ = v11352
	var v11358 int32
	_ = v11358
	var v11403 int32
	_ = v11403
	var v11404 int32
	_ = v11404
	var v11405 int32
	_ = v11405
	var v11406 int32
	_ = v11406
	var v11407 int32
	_ = v11407
	var v11409 int32
	_ = v11409
	var v11410 int32
	_ = v11410
	var v11414 int32
	_ = v11414
	var v11415 int32
	_ = v11415
	var v11416 int32
	_ = v11416
	var v11417 int32
	_ = v11417
	var v11419 int32
	_ = v11419
	var v11420 int32
	_ = v11420
	var v11421 int32
	_ = v11421
	var v11432 int32
	_ = v11432
	var v11473 int32
	_ = v11473
	var v11474 int32
	_ = v11474
	var v11479 int32
	_ = v11479
	var v11482 int32
	_ = v11482
	var v11484 int32
	_ = v11484
	var v11532 int32
	_ = v11532
	var v11621 int32
	_ = v11621
	var v11624 int32
	_ = v11624
	var v11625 int32
	_ = v11625
	var v11626 int32
	_ = v11626
	var v11633 int32
	_ = v11633
	var v11636 int32
	_ = v11636
	var v11640 int32
	_ = v11640
	var v11678 int32
	_ = v11678
	var v11682 int32
	_ = v11682
	var v11683 int32
	_ = v11683
	var v11684 int32
	_ = v11684
	var v11687 int32
	_ = v11687
	var v11688 int32
	_ = v11688
	var v11689 int32
	_ = v11689
	var v11690 int32
	_ = v11690
	var v11691 int32
	_ = v11691
	var v11693 int32
	_ = v11693
	var v11695 int32
	_ = v11695
	var v11696 int32
	_ = v11696
	var v11697 int32
	_ = v11697
	var v11698 int32
	_ = v11698
	var v11699 int32
	_ = v11699
	var v11703 int32
	_ = v11703
	var v11707 int32
	_ = v11707
	var v11711 int32
	_ = v11711
	var v11712 int32
	_ = v11712
	var v11713 int32
	_ = v11713
	var v11714 int32
	_ = v11714
	var v11717 int32
	_ = v11717
	var v11718 int32
	_ = v11718
	var v11720 int32
	_ = v11720
	var v11721 int32
	_ = v11721
	var v11723 int32
	_ = v11723
	var v11724 int32
	_ = v11724
	var v11732 int32
	_ = v11732
	var v11736 int32
	_ = v11736
	var v11774 int32
	_ = v11774
	var v11777 int32
	_ = v11777
	var v11778 int32
	_ = v11778
	var v11781 int32
	_ = v11781
	var v11784 int32
	_ = v11784
	var v11785 int32
	_ = v11785
	var v11786 int32
	_ = v11786
	var v11787 int32
	_ = v11787
	var v11789 int32
	_ = v11789
	var v11790 int32
	_ = v11790
	var v11791 int32
	_ = v11791
	var v11803 int32
	_ = v11803
	var v11805 float64
	_ = v11805
	var v11807 float64
	_ = v11807
	var v11809 float64
	_ = v11809
	var v11811 int32
	_ = v11811
	var v11812 int32
	_ = v11812
	var v11818 int32
	_ = v11818
	var v11819 int32
	_ = v11819
	var v11824 int32
	_ = v11824
	var v11836 int32
	_ = v11836
	var v11873 int32
	_ = v11873
	var v11874 int32
	_ = v11874
	var v11875 int32
	_ = v11875
	var v11876 int32
	_ = v11876
	var v11878 int32
	_ = v11878
	var v11879 int32
	_ = v11879
	var v11890 int32
	_ = v11890
	var v11891 int32
	_ = v11891
	var v11892 int32
	_ = v11892
	var v11896 int32
	_ = v11896
	var v11899 int32
	_ = v11899
	var v11902 int32
	_ = v11902
	var v11903 int32
	_ = v11903
	var v11905 int32
	_ = v11905
	var v11948 int32
	_ = v11948
	var v11952 int32
	_ = v11952
	var v11953 int32
	_ = v11953
	var v11954 int32
	_ = v11954
	var v11955 int32
	_ = v11955
	var v11956 int32
	_ = v11956
	var v11958 int32
	_ = v11958
	var v11960 int32
	_ = v11960
	var v11961 int32
	_ = v11961
	var v11967 int32
	_ = v11967
	var v11971 int32
	_ = v11971
	var v11972 int32
	_ = v11972
	var v11974 int32
	_ = v11974
	var v11975 int32
	_ = v11975
	var v11979 int32
	_ = v11979
	var v12025 int32
	_ = v12025
	var v12026 int32
	_ = v12026
	var v12028 int32
	_ = v12028
	var v12031 int32
	_ = v12031
	var v12032 int32
	_ = v12032
	var v12033 int32
	_ = v12033
	var v12034 int32
	_ = v12034
	var v12036 int32
	_ = v12036
	var v12037 int32
	_ = v12037
	var v12038 int32
	_ = v12038
	var v12039 int32
	_ = v12039
	var v12040 int32
	_ = v12040
	var v12043 int32
	_ = v12043
	var v12047 int32
	_ = v12047
	var v12048 int32
	_ = v12048
	var v12054 int32
	_ = v12054
	var v12056 int32
	_ = v12056
	var v12057 int32
	_ = v12057
	var v12062 int32
	_ = v12062
	var v12063 int32
	_ = v12063
	var v12064 int32
	_ = v12064
	var v12065 int32
	_ = v12065
	var v12066 int32
	_ = v12066
	var v12067 int32
	_ = v12067
	var v12069 int32
	_ = v12069
	var v12070 int32
	_ = v12070
	var v12071 int32
	_ = v12071
	var v12073 int32
	_ = v12073
	var v12074 int32
	_ = v12074
	var v12075 int32
	_ = v12075
	var v12077 int32
	_ = v12077
	var v12078 int32
	_ = v12078
	var v12079 int32
	_ = v12079
	var v12080 int32
	_ = v12080
	var v12081 int32
	_ = v12081
	var v12082 int32
	_ = v12082
	var v12086 int32
	_ = v12086
	var v12087 int32
	_ = v12087
	var v12090 int32
	_ = v12090
	var v12091 int32
	_ = v12091
	var v12092 int32
	_ = v12092
	var v12093 int32
	_ = v12093
	var v12094 int32
	_ = v12094
	var v12095 int32
	_ = v12095
	var v12098 int32
	_ = v12098
	var v12099 int32
	_ = v12099
	var v12100 int32
	_ = v12100
	var v12101 int32
	_ = v12101
	var v12104 int32
	_ = v12104
	var v12105 int32
	_ = v12105
	var v12109 int32
	_ = v12109
	var v12110 int32
	_ = v12110
	var v12111 int32
	_ = v12111
	var v12112 int32
	_ = v12112
	var v12113 int32
	_ = v12113
	var v12119 int32
	_ = v12119
	var v12120 int32
	_ = v12120
	var v12121 int32
	_ = v12121
	var v12122 int32
	_ = v12122
	var v12127 int32
	_ = v12127
	var v12128 int32
	_ = v12128
	var v12129 int32
	_ = v12129
	var v12130 int32
	_ = v12130
	var v12136 int32
	_ = v12136
	var v12138 int32
	_ = v12138
	var v12142 int32
	_ = v12142
	var v12143 int32
	_ = v12143
	var v12144 int32
	_ = v12144
	var v12180 int32
	_ = v12180
	var v12184 int32
	_ = v12184
	var v12185 int32
	_ = v12185
	var v12186 int32
	_ = v12186
	var v12187 int32
	_ = v12187
	var v12188 int32
	_ = v12188
	var v12189 int32
	_ = v12189
	var v12190 int32
	_ = v12190
	var v12191 int32
	_ = v12191
	var v12192 int32
	_ = v12192
	var v12193 int32
	_ = v12193
	var v12194 int32
	_ = v12194
	var v12195 int32
	_ = v12195
	var v12196 int32
	_ = v12196
	var v12197 int32
	_ = v12197
	var v12198 int32
	_ = v12198
	var v12199 int32
	_ = v12199
	var v12200 int32
	_ = v12200
	var v12202 int32
	_ = v12202
	var v12203 int32
	_ = v12203
	var v12211 int32
	_ = v12211
	var v12214 int32
	_ = v12214
	var v12215 int32
	_ = v12215
	var v12216 int32
	_ = v12216
	var v12217 int32
	_ = v12217
	var v12218 int32
	_ = v12218
	var v12220 int32
	_ = v12220
	var v12254 int32
	_ = v12254
	var v12255 int32
	_ = v12255
	var v12258 int32
	_ = v12258
	var v12259 int32
	_ = v12259
	var v12260 int32
	_ = v12260
	var v12265 int32
	_ = v12265
	var v12271 int32
	_ = v12271
	var v12273 float64
	_ = v12273
	var v12275 float64
	_ = v12275
	var v12277 float64
	_ = v12277
	var v12279 int32
	_ = v12279
	var v12283 int32
	_ = v12283
	var v12286 int32
	_ = v12286
	var v12289 int32
	_ = v12289
	var v12291 float64
	_ = v12291
	var v12293 int32
	_ = v12293
	var v12294 int32
	_ = v12294
	var v12295 int32
	_ = v12295
	var v12296 int32
	_ = v12296
	var v12298 int32
	_ = v12298
	var v12299 int32
	_ = v12299
	var v12316 int32
	_ = v12316
	var v12361 int32
	_ = v12361
	var v12363 float64
	_ = v12363
	var v12365 float64
	_ = v12365
	var v12367 float64
	_ = v12367
	var v12369 int32
	_ = v12369
	var v12370 int32
	_ = v12370
	var v12372 int32
	_ = v12372
	var v12374 int32
	_ = v12374
	var v12376 int32
	_ = v12376
	var v12379 int32
	_ = v12379
	var v12380 int32
	_ = v12380
	var v12381 int32
	_ = v12381
	var v12383 int32
	_ = v12383
	var v12384 int32
	_ = v12384
	var v12387 int32
	_ = v12387
	var v12388 int32
	_ = v12388
	var v12392 int32
	_ = v12392
	var v12397 int32
	_ = v12397
	v4 = int32(0)
	v49 = m.G0
	v51 = v49 - int32(176)
	m.G0 = v51
	F_check_stack_depth(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v57 - int32(335) {
	case 0:
		goto L9
	case 1:
		goto L10
	case 2:
		goto L23
	case 3:
		goto L7
	case 4:
		goto L8
	case 5:
		goto L21
	default:
		goto L26
	case 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24:
		goto L28
	case 25:
		goto L6
	case 27:
		goto L27
	case 28:
		goto L5
	case 29:
		goto L11
	case 30:
		goto L12
	case 31:
		goto L15
	case 32:
		goto L16
	case 33:
		goto L17
	case 34:
		goto L18
	case 35:
		goto L19
	case 36:
		goto L13
	case 37:
		goto L14
	case 38:
		goto L25
	case 40:
		goto L20
	case 41:
		goto L22
	case 42:
		goto L24
	}
L3:
	;
	m.G0 = v12397 + int32(176)
	return v12392
L4:
	;
	v12361 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12316)+4)) = v12361
	v12363 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v12316)+8)) = v12363
	v12365 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v12316)+16)) = v12365
	v12367 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12316)+24)) = v12367
	v12369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12370 = *(*int32)(unsafe.Add(mBase, uint32(v12369)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12316)+32)) = v12370
	v12372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12316)+36)) = uint8(v12372)
	v12374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12316)+37)) = uint8(v12374)
	v12376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+335)))
	if v12376 != int32(1) {
		v12392 = v12316
		v12397 = v51
		goto L3
	} else {
		goto L1641
	}
L5:
	;
	v11890 = int32(0)
	v11891 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11892 = *(*int32)(unsafe.Add(mBase, uint32(v11891)+4))
	if v11892 == v11890 {
		v11979 = v11890
		goto L1579
	} else {
		goto L1580
	}
L6:
	;
	v10547 = int32(0)
	v10548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10549 = *(*int32)(unsafe.Add(mBase, uint32(v10548)+4))
	if v10549 == v10547 {
		v10636 = v10547
		goto L1381
	} else {
		goto L1382
	}
L7:
	;
	v9755 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9756 = *(*int32)(unsafe.Add(mBase, uint32(v9755)+4))
	if v9756 == int32(0) {
		goto L1250
	} else {
		goto L1251
	}
L8:
	;
	v9063 = F_palloc0(m, int32(112))
	mBase = m.M
	v9064 = m.ExcPending
	if v9064 != 0 {
		goto L1
	} else {
		goto L1144
	}
L9:
	;
	v8057 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v8057 - int32(295) {
	case 0:
		goto L1021
	default:
		goto L1020
	case 8:
		goto L1023
	case 18:
		goto L1022
	}
L10:
	;
	v7892 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7894 = F_create_plan_recurse(m, l0, v7892, int32(0))
	mBase = m.M
	v7895 = m.ExcPending
	if v7895 != 0 {
		goto L1
	} else {
		goto L1003
	}
L11:
	;
	v7860 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7863 = F_create_plan_recurse(m, l0, v7860, l2|int32(2))
	mBase = m.M
	v7864 = m.ExcPending
	if v7864 != 0 {
		goto L1
	} else {
		goto L1001
	}
L12:
	;
	v7688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7691 = F_create_plan_recurse(m, l0, v7688, l2|int32(2))
	mBase = m.M
	v7692 = m.ExcPending
	if v7692 != 0 {
		goto L1
	} else {
		goto L973
	}
L13:
	;
	v7273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7276 = F_create_plan_recurse(m, l0, v7273, l2|int32(4))
	mBase = m.M
	v7277 = m.ExcPending
	if v7277 != 0 {
		goto L1
	} else {
		goto L927
	}
L14:
	;
	v7094 = int32(1)
	v7095 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7097 = F_create_plan_recurse(m, l0, v7095, v7094)
	mBase = m.M
	v7098 = m.ExcPending
	if v7098 != 0 {
		goto L1
	} else {
		goto L909
	}
L15:
	;
	v7014 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7017 = F_create_plan_recurse(m, l0, v7014, l2|int32(2))
	mBase = m.M
	v7018 = m.ExcPending
	if v7018 != 0 {
		goto L1
	} else {
		goto L903
	}
L16:
	;
	v6939 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6942 = F_create_plan_recurse(m, l0, v6939, l2|int32(2))
	mBase = m.M
	v6943 = m.ExcPending
	if v6943 != 0 {
		goto L1
	} else {
		goto L897
	}
L17:
	;
	v6755 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6757 = F_create_plan_recurse(m, l0, v6755, int32(4))
	mBase = m.M
	v6758 = m.ExcPending
	if v6758 != 0 {
		goto L1
	} else {
		goto L873
	}
L18:
	;
	v5122 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5122 == int32(312) {
		goto L707
	} else {
		goto L708
	}
L19:
	;
	v4644 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v4645 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+12))
	if v4645 != 0 {
		goto L662
	} else {
		goto L663
	}
L20:
	;
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4315)+4))
	if v4316 == int32(0) {
		v4407 = v4
		goto L626
	} else {
		goto L627
	}
L21:
	;
	v3950 = int32(1)
	v3951 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3953 = F_create_plan_recurse(m, l0, v3951, v3950)
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L591
	}
L22:
	;
	v3918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3919 = F_create_plan_recurse(m, l0, v3918, l2)
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L1
	} else {
		goto L589
	}
L23:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1474 = F_create_plan_recurse(m, l0, v1472, int32(1))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L1
	} else {
		goto L215
	}
L24:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1291 = F_create_plan_recurse(m, l0, v1290, l2)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L198
	}
L25:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1104 = int32(0)
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+4))
	if v1106 == v1104 {
		v1193 = v1104
		goto L179
	} else {
		goto L180
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L176
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v63 == int32(0) {
		v162 = v4
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v60 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v12392 = v60
	v12397 = v51
	goto L3
L30:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v200 != 0 {
		goto L45
	} else {
		goto L46
	}
L31:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v67 <= int32(0) {
		v162 = v4
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v74 = int32(1)
	v76 = v4
	v85 = v4
	goto L33
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v76<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v124 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v162 = v142
	goto L30
L35:
	;
	v125 = F_replace_nestloop_params_mutator(m, v123, l0)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v127 = v123
	goto L37
L37:
	;
	v129 = int32(0)
	v131 = F_makeTargetEntry(m, v127, base.I32_extend16_s(v74), v129, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v127 = v125
	goto L37
L39:
	;
	if v70 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v70+v74<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v138
	goto L42
L41:
	;
	goto L42
L42:
	;
	v142 = F_lappend(m, v85, v131)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v145 = v76 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v145 < v146 {
		v74 = v74 + int32(1)
		v76 = v145
		v85 = v142
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	v201 = int32(2)
	goto L47
L46:
	;
	v201 = int32(0)
	goto L47
L47:
	;
	v202 = F_create_plan_recurse(m, l0, v197, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v207 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v208 = int32(2)
	goto L51
L50:
	;
	v208 = int32(0)
	goto L51
L51:
	;
	v209 = F_create_plan_recurse(m, l0, v204, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v212 = F_order_qual_clauses(m, l0, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = v212
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v216)%32)&int32(174) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v236 = F_get_actual_clauses(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	F_extract_actual_join_clauses(m, v212, v221, v51+int32(84), v51+int32(80))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v229 = F_extract_actual_clauses(m, v212, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+80)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = v229
	goto L54
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	v239 = F_list_difference(m, v238, v236)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = v239
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v242 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v243 = F_replace_nestloop_params_mutator(m, v239, l0)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v251 = l1 + int32(104)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	v256 = F_get_switched_clauses(m, v252, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v247 = F_replace_nestloop_params_mutator(m, v246, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+80)) = v247
	goto L64
L67:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v258 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v432 != 0 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v260 = l1 + int32(100)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[0])))
	if v264 != int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v420 = v417 - int32(-64)
	v423 = v202
	goto L68
L72:
	;
	v331 = int32(0)
	v334 = v51 + int32(88)
	v343 = F_prepare_sort_from_pathkeys(m, v202, v258, v262, v331, v331, v334, v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L78
	}
L73:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v267 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v270 = int32(0)
	v273 = v51 + int32(88)
	v282 = F_prepare_sort_from_pathkeys(m, v202, v258, v262, v270, v270, v273, v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v290 = F_palloc0(m, int32(104))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = int32(367)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v282)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+96)) = v267
	v296 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v290)+56)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v290)+52)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v290)+48)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v290)+44)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v290)+88)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v290)+84)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v290)+80)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v290)+76)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v290)+72)) = v284
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v282)+8))
	v310 = *(*float64)(unsafe.Add(mBase, uint32(v282)+16))
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v282)+24))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v282)+32))
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	F_cost_incremental_sort(m, v273, l0, v307, v267, v308, v309, v310, v311, v312, v314, float64(-1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v318 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v290)+8)) = v318
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v290)+16)) = v320
	v322 = *(*float64)(unsafe.Add(mBase, uint32(v282)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v290)+24)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v282)+32))
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+36)) = uint8(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v290)+32)) = v324
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+37)) = uint8(v328)
	v420 = v260
	v423 = v290
	goto L68
L78:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v351 = F_palloc0(m, int32(96))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = int32(366)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v343)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+44)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v358 = int32(_a_F_create_plan_recurse_0)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+88)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v351)+84)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v351)+80)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v351)+76)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v351)+72)) = v345
	v365 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+56)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v351)+52)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v351)+48)) = v365
	v370 = int32(1)
	v372 = v357 + (v359 ^ v370)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+4)) = v372
	v374 = *(*float64)(unsafe.Add(mBase, uint32(v343)+16))
	v375 = *(*float64)(unsafe.Add(mBase, uint32(v343)+24))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v343)+32))
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v382 = m.G0
	v383 = int32(16)
	v384 = v382 - v383
	m.G0 = v384
	F_cost_tuplesort(m, v384+int32(8), v384, v375, v376, float64(0), v379, float64(-1))
	mBase = m.M
	v389 = *(*float64)(unsafe.Add(mBase, uint32(v384)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v334)+32)) = v375
	v392 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v393 = base.F64_add(v374, v389)
	*(*float64)(unsafe.Add(mBase, uint32(v334)+48)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v334)+40)) = v372 + (v392 ^ v370)
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v384)))
	*(*float64)(unsafe.Add(mBase, uint32(v334)+56)) = base.F64_add(v393, v399)
	m.G0 = v384 + v383
	goto L80
L80:
	;
	v405 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v351)+8)) = v405
	v407 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v351)+16)) = v407
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v343)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v351)+24)) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v343)+32))
	v412 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+36)) = uint8(v412)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+32)) = v411
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+37)) = uint8(v415)
	v420 = v260
	v423 = v351
	goto L68
L81:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+113)))
	if v536 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+8))
	v435 = int32(0)
	v438 = v51 + int32(88)
	v447 = F_prepare_sort_from_pathkeys(m, v209, v432, v434, v435, v435, v438, v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v525 = v209
	v527 = v521 - int32(-64)
	goto L81
L85:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v455 = F_palloc0(m, int32(96))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(366)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v447)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+44)) = v459
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v462 = int32(_a_F_create_plan_recurse_0)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+88)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v455)+84)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v455)+80)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v455)+76)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v455)+72)) = v449
	v469 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+56)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v455)+52)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v455)+48)) = v469
	v474 = int32(1)
	v476 = v461 + (v463 ^ v474)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v476
	v478 = *(*float64)(unsafe.Add(mBase, uint32(v447)+16))
	v479 = *(*float64)(unsafe.Add(mBase, uint32(v447)+24))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v447)+32))
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v486 = m.G0
	v487 = int32(16)
	v488 = v486 - v487
	m.G0 = v488
	F_cost_tuplesort(m, v488+int32(8), v488, v479, v480, float64(0), v483, float64(-1))
	mBase = m.M
	v493 = *(*float64)(unsafe.Add(mBase, uint32(v488)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v438)+32)) = v479
	v496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v497 = base.F64_add(v478, v493)
	*(*float64)(unsafe.Add(mBase, uint32(v438)+48)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v438)+40)) = v476 + (v496 ^ v474)
	v503 = *(*float64)(unsafe.Add(mBase, uint32(v488)))
	*(*float64)(unsafe.Add(mBase, uint32(v438)+56)) = base.F64_add(v497, v503)
	m.G0 = v488 + v487
	goto L87
L87:
	;
	v509 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v455)+8)) = v509
	v511 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v455)+16)) = v511
	v513 = *(*float64)(unsafe.Add(mBase, uint32(v447)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v455)+24)) = v513
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v447)+32))
	v516 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+36)) = uint8(v516)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+32)) = v515
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+37)) = uint8(v519)
	v525 = v455
	v527 = v251
	goto L81
L88:
	;
	v574 = int32(0)
	if v256 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v571 = v525
	goto L88
L90:
	;
	goto L91
L91:
	;
	v540 = F_palloc0(m, int32(72))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = int32(364)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v525)+44))
	v545 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v540)+56)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v540)+52)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v540)+48)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v540)+44)) = v544
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+4)) = v551
	v553 = *(*float64)(unsafe.Add(mBase, uint32(v525)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v540)+8)) = v553
	v555 = *(*float64)(unsafe.Add(mBase, uint32(v525)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v540)+16)) = v555
	v557 = *(*float64)(unsafe.Add(mBase, uint32(v525)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v540)+24)) = v557
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v525)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+36)) = uint8(v545)
	*(*int32)(unsafe.Add(mBase, uint32(v540)+32)) = v559
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+37)) = uint8(v563)
	v566 = *(*float64)(unsafe.Add(mBase, _c_F_create_plan_recurse[3]))
	*(*float64)(unsafe.Add(mBase, uint32(v540)+16)) = base.F64_add(v555, base.F64_mul(v557, v566))
	v571 = v540
	goto L88
L93:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v577 = v576
	goto L95
L94:
	;
	v577 = v574
	goto L95
L95:
	;
	v579 = v577 << (uint(int32(2)) % 32)
	v580 = F_palloc(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v582 = F_palloc(m, v579)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v584 = F_palloc(m, v577)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v586 = F_palloc(m, v577)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v431 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v589 = v588
	goto L102
L101:
	;
	v589 = v574
	goto L102
L102:
	;
	if v535 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v592 = v590
	goto L105
L104:
	;
	v592 = int32(0)
	goto L105
L105:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	if v593 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L173
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L170
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L167
	}
L109:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+112)))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v1033 = F_palloc0(m, int32(112))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L166
	}
L110:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v596 <= int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v599 = int32(0)
	v606 = v599
	v608 = v4
	v611 = v599
	v612 = v592
	v614 = v589
	goto L112
L112:
	;
	v650 = v611 << (uint(int32(2)) % 32)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v650+v651)))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+120)))
	if v656 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L109
L114:
	;
	v657 = int32(104)
	goto L116
L115:
	;
	v657 = int32(100)
	goto L116
L116:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v653+v657)))
	if v656 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v682)+8))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v902)+8))
	if v947 != v948 {
		goto L107
	} else {
		goto L158
	}
L118:
	;
	if v535 == int32(0) {
		v804 = v729
		v812 = v730
		goto L146
	} else {
		goto L147
	}
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L142
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L139
	}
L121:
	;
	v662 = int32(100)
	goto L123
L122:
	;
	v662 = int32(104)
	goto L123
L123:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v653+v662)))
	if v608 != v664 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if v614 == int32(0) {
		goto L120
	} else {
		goto L127
	}
L125:
	;
	v682 = v606
	v683 = v608
	v684 = v614
	goto L126
L126:
	;
	if v612 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	if v664 != v669 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	v672 = v614 + int32(4)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if base.Ui32(v672) < base.Ui32(v674+v675<<(uint(int32(2))%32)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v680 = v672
	goto L131
L130:
	;
	v680 = int32(0)
	goto L131
L131:
	;
	v682 = v668
	v683 = v664
	v684 = v680
	goto L126
L132:
	;
	v687 = int32(0)
	v729 = v687
	v730 = v687
	goto L118
L133:
	;
	goto L134
L134:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+4))
	if v659 != v690 {
		v729 = v689
		v730 = v690
		goto L118
	} else {
		goto L135
	}
L135:
	;
	v693 = v612 + int32(4)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if base.Ui32(v693) < base.Ui32(v695+v696<<(uint(int32(2))%32)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v701 = v693
	goto L138
L137:
	;
	v701 = int32(0)
	goto L138
L138:
	;
	v902 = v689
	v909 = v701
	v946 = int32(1)
	goto L117
L139:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_1), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_3), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_1), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_5), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v902 = v853
	v909 = v612
	v946 = int32(0)
	goto L117
L146:
	;
	if v812 != v659 {
		goto L108
	} else {
		goto L157
	}
L147:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if v733 <= int32(0) {
		v804 = v729
		v812 = v730
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v736 = int32(0)
	if v736 < v733 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v739 = v733
	goto L151
L150:
	;
	v739 = v736
	goto L151
L151:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v745 = int32(0)
	v746 = v729
	v754 = v730
	goto L152
L152:
	;
	v792 = v740 + v745<<(uint(int32(2))%32)
	if v792 == v612 {
		v804 = v746
		v812 = v754
		goto L146
	} else {
		goto L154
	}
L153:
	;
	v804 = v794
	v812 = v795
	goto L146
L154:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v792)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v659 == v795 {
		v853 = v794
		goto L145
	} else {
		goto L155
	}
L155:
	;
	v798 = v745 + int32(1)
	if v798 != v739 {
		v745 = v798
		v746 = v794
		v754 = v795
		goto L152
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	v853 = v804
	goto L145
L158:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+8))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)+8))
	if v951 != v953 {
		goto L107
	} else {
		goto L159
	}
L159:
	;
	if v946 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v682)+12))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v902)+12))
	if v955 != v956 {
		goto L106
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580+v650))) = v947
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v582+v650))) = v965
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v682)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v611+v584))) = uint8(base.B2i32(v968 == int32(5)))
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v611+v586))) = uint8(v973)
	v976 = v611 + int32(1)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v976 < v977 {
		v606 = v682
		v608 = v683
		v611 = v976
		v612 = v909
		v614 = v684
		goto L112
	} else {
		goto L165
	}
L163:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+16)))
	if v958 != v959 {
		goto L106
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	goto L113
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+108)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+104)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+100)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+96)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+92)) = v256
	*(*uint8)(unsafe.Add(mBase, uint32(v1033)+88)) = uint8(v1029)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+56)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+52)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+48)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+44)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = int32(362)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+80)) = v1030
	*(*uint8)(unsafe.Add(mBase, uint32(v1033)+76)) = uint8(v1028)
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+72)) = v1027
	v12316 = v1033
	goto L4
L167:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_6), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_7), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_8), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_9), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_8), int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_10), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v1093
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_11), v51)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(538), int32(_a_F_create_plan_recurse_12))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1241 = F_create_plan_recurse(m, l0, v1239, int32(1))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L194
	}
L180:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1110 <= int32(0) {
		v1193 = v1104
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+8))
	v1116 = v1104
	v1117 = int32(1)
	v1119 = v4
	goto L182
L182:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+12))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1162+v1119<<(uint(int32(2))%32))))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1167 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1193 = v1185
	goto L179
L184:
	;
	v1168 = F_replace_nestloop_params_mutator(m, v1166, l0)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	v1170 = v1166
	goto L186
L186:
	;
	v1172 = int32(0)
	v1174 = F_makeTargetEntry(m, v1170, base.I32_extend16_s(v1117), v1172, v1172)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L188
	}
L187:
	;
	v1170 = v1168
	goto L186
L188:
	;
	if v1113 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1113+v1117<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+16)) = v1181
	goto L191
L190:
	;
	goto L191
L191:
	;
	v1185 = F_lappend(m, v1116, v1174)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v1188 = v1119 + int32(1)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1188 < v1189 {
		v1116 = v1185
		v1117 = v1117 + int32(1)
		v1119 = v1188
		goto L182
	} else {
		goto L193
	}
L193:
	;
	goto L183
L194:
	;
	v1244 = F_palloc0(m, int32(104))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+44)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1244))) = int32(373)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+72)) = v1249
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+4)) = v1251
	v1253 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1244)+8)) = v1253
	v1255 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1244)+16)) = v1255
	v1257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1244)+24)) = v1257
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+32)) = v1260
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1244)+36)) = uint8(v1262)
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1244)+37)) = uint8(v1264)
	v1266 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+76)) = v1266
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+8))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+8))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+84))
	v1284 = F_prepare_sort_from_pathkeys(m, v1241, v1103, v1271, v1272, int32(0), v1244+int32(80), v1244+int32(84), v1244+int32(88), v1244+int32(92), v1244+int32(96))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+52)) = v1284
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1287)+95)) = uint8(v1288)
	v12392 = v1244
	v12397 = v51
	goto L3
L198:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v1293 != int32(1) {
		v1393 = v4
		v1395 = v4
		v1399 = v4
		v1400 = v4
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1437 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1440 = F_palloc0(m, int32(104))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L214
	}
L200:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+124))
	if v1297 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+4))
	v1299 = v1298
	goto L203
L202:
	;
	v1299 = v4
	goto L203
L203:
	;
	v1302 = F_palloc(m, v1299<<(uint(int32(1))%32))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1305 = v1299 << (uint(int32(2)) % 32)
	v1306 = F_palloc(m, v1305)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1308 = F_palloc(m, v1305)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+124))
	if v1310 == int32(0) {
		v1393 = v4
		v1395 = v1308
		v1399 = v1302
		v1400 = v1306
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+4))
	if v1313 <= int32(0) {
		v1393 = v4
		v1395 = v1308
		v1399 = v1302
		v1400 = v1306
		goto L199
	} else {
		goto L208
	}
L208:
	;
	v1320 = v4
	goto L209
L209:
	;
	v1368 = v1320 << (uint(int32(2)) % 32)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+12))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1368+v1369)))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+76))
	v1373 = F_get_sortgroupclause_tle(m, v1371, v1372)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L211
	}
L210:
	;
	v1393 = v1386
	v1395 = v1308
	v1399 = v1302
	v1400 = v1306
	goto L199
L211:
	;
	v1375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1373)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1302+v1320<<(uint(int32(1))%32)))) = uint16(v1375)
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1368+v1306))) = v1378
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	v1382 = F_exprCollation(m, v1381)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1368+v1308))) = v1382
	v1386 = v1320 + int32(1)
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+4))
	if v1386 < v1387 {
		v1320 = v1386
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = int32(377)
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+96)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+92)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+88)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+84)) = v1393
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+80)) = v1438
	*(*int64)(unsafe.Add(mBase, uint32(v1440)+72)) = v1437
	v1451 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+56)) = v1451
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+52)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+48)) = v1451
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+44)) = v1444
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+4)) = v1457
	v1459 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1440)+8)) = v1459
	v1461 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1440)+16)) = v1461
	v1463 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1440)+24)) = v1463
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+32)) = v1466
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1440)+36)) = uint8(v1468)
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1440)+37)) = uint8(v1470)
	v12392 = v1440
	v12397 = v51
	goto L3
L215:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+44))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v1485 = int32(0)
	goto L217
L216:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v1537 = F_palloc0(m, int32(168))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L227
	}
L217:
	;
	v1486 = int32(0)
	if v1476 == v1486 {
		v1496 = v1486
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v1477 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+4))
	if v1490 <= v1485 {
		v1496 = int32(0)
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+12))
	v1496 = v1492 + v1485<<(uint(int32(2))%32)
	goto L219
L222:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1496)))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1504+v1485<<(uint(int32(2))%32))))
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+12)) = v1511
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+16)) = v1513
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+20)) = v1515
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1506)+24)) = uint16(v1517)
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1506)+26)) = uint8(v1519)
	v1485 = v1485 + int32(1)
	goto L217
L223:
	;
	goto L216
L224:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+4))
	if base.B2i32(v1496 == int32(0))|base.B2i32(v1501 <= v1485) != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+12))
	if v1504 != 0 {
		goto L222
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	v1539 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+56)) = v1539
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+52)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(v1537))) = int32(337)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+88)) = v1535
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+84)) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+80)) = v1533
	*(*uint8)(unsafe.Add(mBase, uint32(v1537)+76)) = uint8(v1532)
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+72)) = v1531
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+44)) = int64(0)
	if v1526 == v1539 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+96)) = v3355
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+92)) = v3369
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+156)) = v3333
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3380 = *(*int32)(unsafe.Add(mBase, uint32(v3379)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+100)) = v3380
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v3382)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+164)) = v3365
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+160)) = v3366
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+120)) = v3367
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+108)) = v3368
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+104)) = v3383
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+124)) = v3364
	if v3347 == int32(0) {
		goto L449
	} else {
		goto L450
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+152)) = int32(0)
	v1555 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+144)) = v1555
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+136)) = v1555
	*(*int64)(unsafe.Add(mBase, uint32(v1537)+128)) = v1555
	v3328 = l0
	v3329 = l1
	v3331 = v1537
	v3333 = v4
	v3336 = v51
	v3338 = v1533
	v3345 = v1531
	v3347 = v1535
	v3355 = v1529
	v3364 = v1523
	v3365 = v1524
	v3366 = v1525
	v3367 = v1527
	v3368 = v1528
	v3369 = v1530
	goto L228
L230:
	;
	goto L231
L231:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+128)) = v1561
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+136)) = v1563
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+140)) = v1565
	v1567 = int32(0)
	if v1565 == v1567 {
		v1738 = v1567
		goto L232
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+144)) = v1738
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+148)) = v1740
	v1742 = int32(0)
	v1744 = m.G0
	v1746 = v1744 - int32(32)
	m.G0 = v1746
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+84))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+8))
	if v1750 != 0 {
		goto L245
	} else {
		goto L246
	}
L233:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	if int32(0) < v1572 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1577 = int32(1)
	v1581 = v4
	v1582 = v1567
	goto L237
L235:
	;
	v1648 = v4
	goto L236
L236:
	;
	v1738 = v1648
	goto L232
L237:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+12))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1623+v1582<<(uint(int32(2))%32))))
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627)+26)))
	if v1628 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1648 = v1634
	goto L236
L239:
	;
	v1631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1627)+8)))
	v1632 = F_lappend_int(m, v1581, v1631)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	v1634 = v1581
	goto L241
L241:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1627)+8)) = uint16(v1577)
	v1636 = int32(1)
	v1639 = v1582 + v1636
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	if v1639 < v1640 {
		v1577 = v1577 + v1636
		v1581 = v1634
		v1582 = v1639
		goto L237
	} else {
		goto L243
	}
L242:
	;
	v1634 = v1632
	goto L241
L243:
	;
	goto L238
L244:
	;
	m.G0 = v3291 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v3276)+132)) = v3279
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3276)+152)) = v3325
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+36))
	v3328 = v3273
	v3329 = v3274
	v3331 = v3276
	v3333 = v3327
	v3336 = v3281
	v3338 = v3283
	v3345 = v3290
	v3347 = v3292
	v3355 = v3300
	v3364 = v3309
	v3365 = v3310
	v3366 = v3311
	v3367 = v3312
	v3368 = v3313
	v3369 = v3314
	goto L228
L245:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+52))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+12))
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+32))
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1753+v1754<<(uint(int32(2))%32)-int32(4))))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1760)+16))
	v1763 = F_table_open(m, v1761, int32(0))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+16))
	if v1751 != 0 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v3273 = l0
	v3274 = l1
	v3276 = v1537
	v3279 = v1742
	v3281 = v51
	v3283 = v1533
	v3288 = v1526
	v3290 = v1531
	v3291 = v1746
	v3292 = v1535
	v3300 = v1529
	v3309 = v1523
	v3310 = v1524
	v3311 = v1525
	v3312 = v1527
	v3313 = v1528
	v3314 = v1530
	goto L244
L248:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+8))
	if v1765 == int32(0) {
		v1888 = v4
		v1889 = v4
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1908 = F_RelationGetIndexList(m, v1763)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L274
	}
L250:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+4))
	if v1768 <= int32(0) {
		v1888 = v4
		v1889 = v4
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v1783 = v4
	v1799 = v4
	v1800 = v4
	goto L252
L252:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+12))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1819+v1783<<(uint(int32(2))%32))))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+4))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	if v1825 != int32(6) {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L263
	}
L254:
	;
	goto L253
L255:
	;
	v1841 = v1783 + int32(1)
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+4))
	if v1841 < v1842 {
		v1783 = v1841
		v1799 = v1838
		v1800 = v1839
		goto L252
	} else {
		goto L262
	}
L256:
	;
	v1828 = F_lappend(m, v1799, v1824)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1830 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1824)+8)))
	if v1830 == int32(0) {
		goto L254
	} else {
		goto L260
	}
L259:
	;
	v1838 = v1828
	v1839 = v1800
	goto L255
L260:
	;
	v1835 = F_bms_add_member(m, v1800, v1830+int32(7))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	v1838 = v1799
	v1839 = v1835
	goto L255
L262:
	;
	v1888 = v1838
	v1889 = v1839
	goto L249
L263:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_13), int32(0))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(877), int32(_a_F_create_plan_recurse_15))
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_list_free(m, v3221)
	mBase = m.M
	v3249 = m.ExcPending
	if v3249 != 0 {
		goto L1
	} else {
		goto L437
	}
L268:
	;
	v3088 = int32(0)
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3052)+4))
	if v3088 < v3089 {
		goto L430
	} else {
		goto L431
	}
L269:
	;
	v3200 = l0
	v3201 = l1
	v3203 = v1537
	v3206 = v1742
	v3208 = v51
	v3210 = v1533
	v3212 = int32(0)
	v3215 = v1526
	v3217 = v1531
	v3218 = v1746
	v3219 = v1535
	v3221 = v1908
	v3223 = v1763
	v3227 = v1529
	v3235 = v4
	v3236 = v1523
	v3237 = v1524
	v3238 = v1525
	v3239 = v1527
	v3240 = v1528
	v3241 = v1530
	goto L267
L270:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L1
	} else {
		goto L426
	}
L271:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+4))
	if v2218 <= int32(0) {
		v3040 = l0
		v3041 = l1
		v3043 = v1537
		v3047 = v1742
		v3048 = v51
		v3050 = v1533
		v3052 = v1989
		v3055 = v1526
		v3057 = v1531
		v3058 = v1746
		v3059 = v1535
		v3061 = v1908
		v3063 = v1763
		v3067 = v1529
		v3075 = v4
		v3076 = v1523
		v3077 = v1524
		v3078 = v1525
		v3079 = v1527
		v3080 = v1528
		v3081 = v1530
		goto L268
	} else {
		goto L305
	}
L272:
	;
	if v1989 == int32(0) {
		goto L269
	} else {
		goto L304
	}
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L1
	} else {
		goto L301
	}
L274:
	;
	if v1908 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1908)+4))
	if v1910 <= int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	goto L277
L277:
	;
	v2098 = int32(0)
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+16))
	if v2099 == v2098 {
		v3200 = l0
		v3201 = l1
		v3203 = v1537
		v3206 = v1742
		v3208 = v51
		v3210 = v1533
		v3212 = v2098
		v3215 = v1526
		v3217 = v1531
		v3218 = v1746
		v3219 = v1535
		v3221 = v1908
		v3223 = v1763
		v3227 = v1529
		v3235 = v4
		v3236 = v1523
		v3237 = v1524
		v3238 = v1525
		v3239 = v1527
		v3240 = v1528
		v3241 = v1530
		goto L267
	} else {
		goto L298
	}
L278:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+16))
	if v2025 == int32(0) {
		goto L272
	} else {
		goto L287
	}
L279:
	;
	v1989 = int32(0)
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1924 = v4
	v1927 = int32(0)
	goto L282
L282:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1908)+12))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1963+v1924<<(uint(int32(2))%32))))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1760)+24))
	v1969 = F_index_open(m, v1967, v1968)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L284
	}
L283:
	;
	v1989 = v1971
	goto L278
L284:
	;
	v1971 = F_lappend(m, v1927, v1969)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1974 = v1924 + int32(1)
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1908)+4))
	if v1974 < v1975 {
		v1924 = v1974
		v1927 = v1971
		goto L282
	} else {
		goto L286
	}
L286:
	;
	goto L283
L287:
	;
	v2028 = F_get_constraint_index(m, v2025)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	if v2028 == int32(0) {
		goto L270
	} else {
		goto L289
	}
L289:
	;
	if v1989 == int32(0) {
		v2111 = v2028
		goto L273
	} else {
		goto L290
	}
L290:
	;
	v2034 = int32(0)
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+4))
	if v2034 < v2035 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v2039 = v2035
	goto L293
L292:
	;
	v2039 = v2034
	goto L293
L293:
	;
	v2049 = v2034
	goto L294
L294:
	;
	if v2039 == v2049 {
		v2111 = v2028
		goto L273
	} else {
		goto L296
	}
L295:
	;
	v2175 = v2028
	v2176 = v2095
	goto L271
L296:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+12))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2049<<(uint(int32(2))%32)+v2093)))
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+56))
	if v2028 != v2096 {
		v2049 = v2049 + int32(1)
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	v2102 = F_get_constraint_index(m, v2099)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	if v2102 == int32(0) {
		goto L270
	} else {
		goto L300
	}
L300:
	;
	v2111 = v2102
	goto L273
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1746))) = v2111
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_16), v1746)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(933), int32(_a_F_create_plan_recurse_15))
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	v2175 = int32(0)
	v2176 = v1742
	goto L271
L305:
	;
	v2229 = v1742
	v2231 = int32(0)
	v2257 = v4
	goto L306
L306:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+12))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v2270+v2231<<(uint(int32(2))%32))))
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+192))
	v2276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+20)))
	if v2276 == int32(0) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v3040 = l0
	v3041 = l1
	v3043 = v1537
	v3047 = v2930
	v3048 = v51
	v3050 = v1533
	v3052 = v1989
	v3055 = v1526
	v3057 = v1531
	v3058 = v1746
	v3059 = v1535
	v3061 = v1908
	v3063 = v1763
	v3067 = v1529
	v3075 = v2958
	v3076 = v1523
	v3077 = v1524
	v3078 = v1525
	v3079 = v1527
	v3080 = v1528
	v3081 = v1530
	goto L268
L308:
	;
	v2972 = v2231 + int32(1)
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+4))
	if v2972 < v2973 {
		v2229 = v2930
		v2231 = v2972
		v2257 = v2958
		goto L306
	} else {
		goto L425
	}
L309:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v1763)+48))
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2279)+119)))
	if v2280 == int32(112) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v2283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+18)))
	if v2283 != int32(1) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2275)))
	if v2286 == v2175 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L312
L314:
	;
	v2917 = F_lappend_oid(m, v2229, v2871)
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L1
	} else {
		goto L424
	}
L315:
	;
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+15)))
	if v2288 != int32(1) {
		v2871 = v2175
		goto L314
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	if v2025 != 0 {
		goto L328
	} else {
		goto L329
	}
L318:
	;
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+4))
	if v2291&int32(-2) != int32(2) {
		v2871 = v2175
		goto L314
	} else {
		goto L319
	}
L319:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+4))
	if v2305 == int32(2) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v2308 = int32(_a_F_create_plan_recurse_17)
	goto L324
L323:
	;
	v2308 = int32(_a_F_create_plan_recurse_18)
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1746)+16)) = v2308
	F_errmsg(m, int32(_a_F_create_plan_recurse_19), v1746+int32(16))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(1010), int32(_a_F_create_plan_recurse_15))
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2275)))
	v2871 = v2868
	goto L314
L328:
	;
	v2320 = F_IsIndexCompatibleAsArbiter(m, v2176, v2274)
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L1
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v2322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+12)))
	if v2322 != int32(1) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L333
	}
L331:
	;
	if v2320 != 0 {
		goto L327
	} else {
		goto L332
	}
L332:
	;
	v2930 = v2229
	v2958 = v2257
	goto L308
L333:
	;
	v2325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+15)))
	if v2325 != 0 {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L334
	}
L334:
	;
	v2326 = int32(0)
	v2328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2275)+10)))
	if v2326 < v2328 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2344 = v2328
	v2345 = v2326
	v2347 = v2326
	goto L338
L336:
	;
	v2409 = v2326
	goto L337
L337:
	;
	v2443 = int32(0)
	if base.B2i32(v2409 == v2443)|base.B2i32(v1889 == v2443) != 0 {
		v2489 = base.B2i32(v2409|v1889 == v2443)
		goto L346
	} else {
		goto L347
	}
L338:
	;
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+192))
	v2383 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2379+v2347<<(uint(int32(1))%32))+48)))
	if v2383 != 0 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v2409 = v2390
	goto L337
L340:
	;
	v2386 = F_bms_add_member(m, v2345, v2383+int32(7))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L1
	} else {
		goto L343
	}
L341:
	;
	v2389 = v2344
	v2390 = v2345
	goto L342
L342:
	;
	v2392 = v2347 + int32(1)
	if v2392 < base.I32_extend16_s(v2389) {
		v2344 = v2389
		v2345 = v2390
		v2347 = v2392
		goto L338
	} else {
		goto L344
	}
L343:
	;
	v2388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2275)+10)))
	v2389 = v2388
	v2390 = v2386
	goto L342
L344:
	;
	goto L339
L345:
	;
	if v2489 == int32(0) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L356
	}
L346:
	;
	goto L345
L347:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2409)+4))
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+4))
	if v2457 != v2458 {
		v2489 = int32(0)
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v2460 = int32(1)
	if v2457 <= v2460 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v2463 = v2460
	goto L351
L350:
	;
	v2463 = v2457
	goto L351
L351:
	;
	v2464 = int32(8)
	v2469 = int32(0)
	goto L352
L352:
	;
	v2477 = v2469 << (uint(int32(2)) % 32)
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2409+v2464+v2477)))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v1889+v2464+v2477)))
	v2482 = base.B2i32(v2479 == v2481)
	if v2479 != v2481 {
		v2489 = v2482
		goto L346
	} else {
		goto L354
	}
L353:
	;
	v2489 = v2482
	goto L346
L354:
	;
	v2485 = v2469 + int32(1)
	if v2485 != v2463 {
		v2469 = v2485
		goto L352
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	v2496 = F_RelationGetIndexExpressions(m, v2274)
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	if v2496 != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if v1754 != int32(1) {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	v2506 = int32(0)
	goto L360
L360:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+8))
	if v2507 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L361:
	;
	F_ChangeVarNodes(m, v2496, int32(1), v1754)
	mBase = m.M
	v2502 = m.ExcPending
	if v2502 != 0 {
		goto L1
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v2503 = F_eval_const_expressions(m, l0, v2496)
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	v2506 = v2503
	goto L360
L366:
	;
	v2797 = F_list_difference(m, v2506, v1888)
	mBase = m.M
	v2798 = m.ExcPending
	if v2798 != 0 {
		goto L1
	} else {
		goto L409
	}
L367:
	;
	v2510 = int32(0)
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+4))
	if v2511 <= v2510 {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v2540 = v2510
	goto L369
L369:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+12))
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v2562+v2540<<(uint(int32(2))%32))))
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+12))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+8))
	if v2568 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L370:
	;
	goto L366
L371:
	;
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+4))
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2735)))
	if v2736 == int32(6) {
		goto L402
	} else {
		goto L403
	}
L372:
	;
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+52))
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2582)))
	if v2583 <= int32(0) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L381
	}
L373:
	;
	v2575 = F_get_opclass_family(m, v2567)
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L1
	} else {
		goto L379
	}
L374:
	;
	if v2567 == int32(0) {
		goto L371
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	if v2567 != 0 {
		goto L373
	} else {
		goto L378
	}
L377:
	;
	goto L373
L378:
	;
	v2573 = int32(0)
	v2580 = v2573
	v2581 = v2573
	goto L372
L379:
	;
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+12))
	v2578 = F_get_opclass_input_type(m, v2577)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	v2580 = v2575
	v2581 = v2578
	goto L372
L381:
	;
	v2587 = int32(1)
	v2603 = int32(0)
	v2605 = v2587
	v2621 = v2583
	v2622 = v2587
	goto L382
L382:
	;
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+192))
	v2638 = int32(1)
	v2639 = v2605 - v2638
	v2643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2637+v2639<<(uint(v2638)%32))+48)))
	v2645 = v2639 << (uint(int32(2)) % 32)
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+248))
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2645+v2646)))
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+12))
	if v2649 != 0 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	v2930 = v2229
	v2958 = v2257
	goto L308
L384:
	;
	v2684 = v2622 + int32(1)
	v2685 = base.I32_extend16_s(v2684)
	if v2685 <= v2679 {
		v2603 = v2603 + base.B2i32(v2643 != int32(0))
		v2605 = v2685
		v2621 = v2679
		v2622 = v2684
		goto L382
	} else {
		goto L401
	}
L385:
	;
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+208))
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v2650+v2645)))
	if v2580 != v2652 {
		v2679 = v2621
		goto L384
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+8))
	if v2658 != v2648 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+212))
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v2654+v2645)))
	if v2581 != v2656 {
		v2679 = v2621
		goto L384
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	v2661 = v2658
	goto L392
L391:
	;
	v2661 = int32(0)
	goto L392
L392:
	;
	if v2661 != 0 {
		v2679 = v2621
		goto L384
	} else {
		goto L393
	}
L393:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+4))
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v2662)))
	if v2663 == int32(6) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2662)+8)))
	if v2666 != v2643 {
		v2679 = v2621
		goto L384
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	if v2643 != 0 {
		v2679 = v2621
		goto L384
	} else {
		goto L398
	}
L397:
	;
	goto L371
L398:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+12))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v2668+(v2639-v2603)<<(uint(int32(2))%32))))
	v2674 = F_equal(m, v2662, v2673)
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	if v2674 != 0 {
		goto L371
	} else {
		goto L400
	}
L400:
	;
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+52))
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v2676)))
	v2679 = v2677
	goto L384
L401:
	;
	goto L383
L402:
	;
	v2746 = v2540 + int32(1)
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+4))
	if v2746 < v2747 {
		v2540 = v2746
		goto L369
	} else {
		goto L408
	}
L403:
	;
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+8))
	if v2739 != 0 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+12))
	if v2740 != 0 {
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2741 = F_list_member(m, v2506, v2735)
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	if v2741 == int32(0) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L407
	}
L407:
	;
	goto L402
L408:
	;
	goto L370
L409:
	;
	if v2797 != 0 {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L410
	}
L410:
	;
	v2799 = F_RelationGetIndexPredicate(m, v2274)
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	if v2799 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	if v1754 != int32(1) {
		goto L415
	} else {
		goto L416
	}
L413:
	;
	v2813 = int32(0)
	goto L414
L414:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+12))
	v2816 = F_predicate_implied_by(m, v2813, v2814, int32(0))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L1
	} else {
		goto L422
	}
L415:
	;
	F_ChangeVarNodes(m, v2799, int32(1), v1754)
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	v2806 = F_make_ands_explicit(m, v2799)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L1
	} else {
		goto L419
	}
L418:
	;
	goto L417
L419:
	;
	v2808 = F_eval_const_expressions(m, l0, v2806)
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	v2810 = F_make_ands_implicit(m, v2808)
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v2813 = v2810
	goto L414
L422:
	;
	if v2816 == int32(0) {
		v2930 = v2229
		v2958 = v2257
		goto L308
	} else {
		goto L423
	}
L423:
	;
	goto L327
L424:
	;
	v2919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+18)))
	v2930 = v2917
	v2958 = base.B2i32(v2919|v2257 != int32(0))
	goto L308
L425:
	;
	goto L307
L426:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_20), int32(0))
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(916), int32(_a_F_create_plan_recurse_15))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L430:
	;
	v3101 = v3088
	goto L433
L431:
	;
	goto L432
L432:
	;
	v3200 = v3040
	v3201 = v3041
	v3203 = v3043
	v3206 = v3047
	v3208 = v3048
	v3210 = v3050
	v3212 = v3052
	v3215 = v3055
	v3217 = v3057
	v3218 = v3058
	v3219 = v3059
	v3221 = v3061
	v3223 = v3063
	v3227 = v3067
	v3235 = v3075
	v3236 = v3076
	v3237 = v3077
	v3238 = v3078
	v3239 = v3079
	v3240 = v3080
	v3241 = v3081
	goto L267
L433:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(v3052)+12))
	v3144 = *(*int32)(unsafe.Add(mBase, uint32(v3140+v3101<<(uint(int32(2))%32))))
	F_relation_close(m, v3144, int32(0))
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L1
	} else {
		goto L435
	}
L434:
	;
	goto L432
L435:
	;
	v3149 = v3101 + int32(1)
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3052)+4))
	if v3149 < v3150 {
		v3101 = v3149
		goto L433
	} else {
		goto L436
	}
L436:
	;
	goto L434
L437:
	;
	F_list_free(m, v3212)
	mBase = m.M
	v3251 = m.ExcPending
	if v3251 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	F_relation_close(m, v3223, int32(0))
	mBase = m.M
	v3254 = m.ExcPending
	if v3254 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	if v3235 != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v3256 = v3206
	goto L442
L441:
	;
	v3256 = int32(0)
	goto L442
L442:
	;
	if v3256 != 0 {
		v3273 = v3200
		v3274 = v3201
		v3276 = v3203
		v3279 = v3206
		v3281 = v3208
		v3283 = v3210
		v3288 = v3215
		v3290 = v3217
		v3291 = v3218
		v3292 = v3219
		v3300 = v3227
		v3309 = v3236
		v3310 = v3237
		v3311 = v3238
		v3312 = v3239
		v3313 = v3240
		v3314 = v3241
		goto L244
	} else {
		goto L443
	}
L443:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3260 = m.ExcPending
	if v3260 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errcode(m, int32(_a_F_create_plan_recurse_21))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_22), int32(0))
	mBase = m.M
	v3267 = m.ExcPending
	if v3267 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(1165), int32(_a_F_create_plan_recurse_15))
	mBase = m.M
	v3272 = m.ExcPending
	if v3272 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+116)) = v3866
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+112)) = v3864
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v3329)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+4)) = v3903
	v3905 = *(*float64)(unsafe.Add(mBase, uint32(v3329)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3331)+8)) = v3905
	v3907 = *(*float64)(unsafe.Add(mBase, uint32(v3329)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3331)+16)) = v3907
	v3909 = *(*float64)(unsafe.Add(mBase, uint32(v3329)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3331)+24)) = v3909
	v3911 = *(*int32)(unsafe.Add(mBase, uint32(v3329)+12))
	v3912 = *(*int32)(unsafe.Add(mBase, uint32(v3911)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3331)+32)) = v3912
	v3914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3329)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3331)+36)) = uint8(v3914)
	v3916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3329)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3331)+37)) = uint8(v3916)
	v12392 = v3331
	v12397 = v3336
	goto L3
L449:
	;
	v3392 = int32(0)
	v3864 = v3392
	v3866 = v3392
	goto L448
L450:
	;
	goto L451
L451:
	;
	v3394 = int32(0)
	v3395 = *(*int32)(unsafe.Add(mBase, uint32(v3347)+4))
	if v3395 <= v3394 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v3864 = int32(0)
	v3866 = v3394
	goto L448
L453:
	;
	goto L454
L454:
	;
	v3401 = int32(0)
	v3411 = v3401
	v3414 = v3401
	v3418 = v3401
	v3419 = v3401
	v3420 = v3394
	v3421 = v3401
	v3422 = v3401
	goto L455
L455:
	;
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(v3347)+12))
	v3459 = *(*int32)(unsafe.Add(mBase, uint32(v3455+v3411<<(uint(int32(2))%32))))
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+40))
	if base.Ui32(v3460) <= base.Ui32(v3459) {
		goto L461
	} else {
		goto L462
	}
L456:
	;
	v3864 = v3847
	v3866 = v3841
	goto L448
L457:
	;
	v3847 = F_lappend(m, v3418, v3836)
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L1
	} else {
		goto L587
	}
L458:
	;
	v3832 = F_bms_add_member(m, v3420, v3411)
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L586
	}
L459:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3817 = m.ExcPending
	if v3817 != 0 {
		goto L1
	} else {
		goto L582
	}
L460:
	;
	v3500 = int32(0)
	if base.B2i32(v3345 != int32(5))|base.B2i32(v3499 == v3500) == v3500 {
		goto L472
	} else {
		goto L473
	}
L461:
	;
	v3471 = int32(0)
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+44))
	if v3472 != 0 {
		goto L465
	} else {
		goto L466
	}
L462:
	;
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+36))
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(v3462+v3459<<(uint(int32(2))%32))))
	if v3466 == int32(0) {
		goto L461
	} else {
		goto L463
	}
L463:
	;
	v3469 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+176))
	v3499 = v3469
	goto L460
L464:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3484)))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v3485)+12))
	if v3486 != 0 {
		v3836 = v3471
		v3838 = v3414
		v3840 = v3419
		v3841 = v3420
		v3842 = v3421
		v3843 = v3422
		goto L457
	} else {
		goto L468
	}
L465:
	;
	v3484 = v3472 + v3459<<(uint(int32(2))%32)
	goto L464
L466:
	;
	goto L467
L467:
	;
	v3476 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3476)+52))
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+12))
	v3484 = v3478 + v3459<<(uint(int32(2))%32) - int32(4)
	goto L464
L468:
	;
	v3487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3485)+21)))
	if v3487 != int32(102) {
		v3836 = v3471
		v3838 = v3414
		v3840 = v3419
		v3841 = v3420
		v3842 = v3421
		v3843 = v3422
		goto L457
	} else {
		goto L469
	}
L469:
	;
	v3491 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[4])))
	if v3491&int32(2) != 0 {
		goto L459
	} else {
		goto L470
	}
L470:
	;
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(v3485)+16))
	v3495 = F_GetFdwRoutineByRelId(m, v3494)
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v3499 = v3495
	goto L460
L472:
	;
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+44))
	if v3505 != 0 {
		goto L476
	} else {
		goto L477
	}
L473:
	;
	goto L474
L474:
	;
	v3543 = int32(0)
	if v3499 == v3543 {
		v3836 = v3543
		v3838 = v3414
		v3840 = v3419
		v3841 = v3420
		v3842 = v3421
		v3843 = v3422
		goto L457
	} else {
		goto L485
	}
L475:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3517)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L1
	} else {
		goto L479
	}
L476:
	;
	v3517 = v3505 + v3459<<(uint(int32(2))%32)
	goto L475
L477:
	;
	goto L478
L478:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3510 = *(*int32)(unsafe.Add(mBase, uint32(v3509)+52))
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3510)+12))
	v3517 = v3511 + v3459<<(uint(int32(2))%32) - int32(4)
	goto L475
L479:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(v3518)+16))
	v3527 = F_get_rel_name(m, v3526)
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3336)+16)) = v3527
	F_errmsg(m, int32(_a_F_create_plan_recurse_23), v3336+int32(16))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v3535 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3518)+21)))
	F_errdetail_relkind_not_supported(m, v3535)
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_24), int32(_a_F_create_plan_recurse_25))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L485:
	;
	v3546 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+88))
	if v3546 == int32(0) {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v3809 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+44))
	if v3809 == int32(0) {
		v3836 = v3543
		v3838 = v3802
		v3840 = v3804
		v3841 = v3420
		v3842 = v3805
		v3843 = v3806
		goto L457
	} else {
		goto L580
	}
L487:
	;
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+92))
	if v3549 == int32(0) {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v3552 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+96))
	if base.B2i32(v3552 == int32(0))|v3355 != 0 {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L489
	}
L489:
	;
	v3556 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+100))
	if v3556 == int32(0) {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L490
	}
L490:
	;
	v3559 = m.G0
	v3561 = v3559 - int32(16)
	m.G0 = v3561
	v3563 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+44))
	if v3563 != 0 {
		goto L493
	} else {
		goto L494
	}
L491:
	;
	if v3604 != 0 {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L516
	}
L492:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v3575)))
	v3577 = *(*int32)(unsafe.Add(mBase, uint32(v3576)+16))
	v3579 = F_table_open(m, v3577, int32(0))
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L1
	} else {
		goto L496
	}
L493:
	;
	v3575 = v3563 + v3459<<(uint(int32(2))%32)
	goto L492
L494:
	;
	goto L495
L495:
	;
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3568 = *(*int32)(unsafe.Add(mBase, uint32(v3567)+52))
	v3569 = *(*int32)(unsafe.Add(mBase, uint32(v3568)+12))
	v3575 = v3569 + v3459<<(uint(int32(2))%32) - int32(4)
	goto L492
L496:
	;
	v3581 = *(*int32)(unsafe.Add(mBase, uint32(v3579)+76))
	v3582 = int32(0)
	switch v3345 - int32(2) {
	case 0:
		goto L501
	case 1:
		goto L502
	case 2:
		goto L500
	case 3:
		v3604 = v3582
		goto L498
	default:
		goto L497
	}
L497:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3614 = m.ExcPending
	if v3614 != 0 {
		goto L1
	} else {
		goto L513
	}
L498:
	;
	F_relation_close(m, v3579, int32(0))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L1
	} else {
		goto L512
	}
L499:
	;
	v3604 = int32(1)
	goto L498
L500:
	;
	if v3581 == int32(0) {
		v3604 = v3582
		goto L498
	} else {
		goto L509
	}
L501:
	;
	if v3581 == int32(0) {
		v3604 = v3582
		goto L498
	} else {
		goto L506
	}
L502:
	;
	if v3581 == int32(0) {
		v3604 = v3582
		goto L498
	} else {
		goto L503
	}
L503:
	;
	v3587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+9)))
	if v3587 != 0 {
		goto L499
	} else {
		goto L504
	}
L504:
	;
	v3588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+8)))
	if v3588 == int32(1) {
		goto L499
	} else {
		goto L505
	}
L505:
	;
	v3604 = v3582
	goto L498
L506:
	;
	v3593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+14)))
	if v3593 != 0 {
		goto L499
	} else {
		goto L507
	}
L507:
	;
	v3594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+13)))
	if v3594 == int32(1) {
		goto L499
	} else {
		goto L508
	}
L508:
	;
	v3604 = v3582
	goto L498
L509:
	;
	v3599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+19)))
	if v3599 != 0 {
		goto L499
	} else {
		goto L510
	}
L510:
	;
	v3600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+18)))
	if v3600 != int32(1) {
		v3604 = v3582
		goto L498
	} else {
		goto L511
	}
L511:
	;
	goto L499
L512:
	;
	m.G0 = v3561 + int32(16)
	goto L491
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3561))) = v3345
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_26), v3561)
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(2522), int32(_a_F_create_plan_recurse_27))
	mBase = m.M
	v3623 = m.ExcPending
	if v3623 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	v3625 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+44))
	if v3625 != 0 {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	v3638 = *(*int32)(unsafe.Add(mBase, uint32(v3637)))
	v3639 = *(*int32)(unsafe.Add(mBase, uint32(v3638)+16))
	v3641 = F_table_open(m, v3639, int32(0))
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		goto L1
	} else {
		goto L521
	}
L518:
	;
	v3637 = v3625 + v3459<<(uint(int32(2))%32)
	goto L517
L519:
	;
	goto L520
L520:
	;
	v3629 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3630 = *(*int32)(unsafe.Add(mBase, uint32(v3629)+52))
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+12))
	v3637 = v3631 + v3459<<(uint(int32(2))%32) - int32(4)
	goto L517
L521:
	;
	v3643 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+52))
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v3643)+24))
	if v3644 != 0 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v3645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3644)+17)))
	v3646 = v3645
	goto L524
L523:
	;
	v3646 = int32(0)
	goto L524
L524:
	;
	F_relation_close(m, v3641, int32(0))
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	if v3646&int32(1) != 0 {
		v3802 = v3414
		v3804 = v3419
		v3805 = v3421
		v3806 = v3422
		goto L486
	} else {
		goto L526
	}
L526:
	;
	if v3422&int32(1) == int32(0) {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	if v3419&int32(1) == int32(0) {
		goto L548
	} else {
		goto L549
	}
L528:
	;
	v3656 = int32(0)
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(v3657)+96))
	if v3658 == v3656 {
		v3680 = v3656
		goto L531
	} else {
		goto L532
	}
L529:
	;
	goto L530
L530:
	;
	v3685 = int32(1)
	if v3421&v3685 == int32(0) {
		goto L527
	} else {
		goto L544
	}
L531:
	;
	if v3680 == int32(0) {
		goto L527
	} else {
		goto L543
	}
L532:
	;
	v3661 = *(*int32)(unsafe.Add(mBase, uint32(v3658)))
	if v3661 != int32(61) {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v3677 = F_expression_tree_walker_impl(m, v3658, int32(952), int32(0))
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L1
	} else {
		goto L542
	}
L534:
	;
	if v3661 != int32(6) {
		goto L533
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3658)+4))
	v3680 = base.B2i32(v3672 == int32(0))
	goto L531
L537:
	;
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3658)+28))
	if v3666 == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v3670 = *(*int32)(unsafe.Add(mBase, uint32(v3658)+32))
	if v3670 != 0 {
		v3680 = int32(1)
		goto L531
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v3680 = int32(0)
	goto L531
L541:
	;
	goto L540
L542:
	;
	v3680 = v3677
	goto L531
L543:
	;
	v3683 = int32(1)
	v3802 = v3414
	v3804 = v3419
	v3805 = v3683
	v3806 = v3683
	goto L486
L544:
	;
	v3802 = v3414
	v3804 = v3419
	v3805 = int32(1)
	v3806 = v3685
	goto L486
L545:
	;
	v3802 = v3799
	v3804 = int32(1)
	v3805 = v3797
	v3806 = v3798
	goto L486
L546:
	;
	v3787 = *(*int32)(unsafe.Add(mBase, uint32(v3499)+88))
	v3788 = m.T0[v3787].(func(*base.Module, int32, int32, int32, int32) int32)(m, v3328, v3331, v3459, v3411)
	mBase = m.M
	v3789 = m.ExcPending
	if v3789 != 0 {
		goto L1
	} else {
		goto L578
	}
L547:
	;
	v3797 = int32(0)
	v3798 = v3779
	v3799 = int32(1)
	goto L545
L548:
	;
	v3698 = m.G0
	v3700 = v3698 - int32(16)
	m.G0 = v3700
	v3702 = int32(0)
	v3703 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+44))
	if v3703 != 0 {
		goto L554
	} else {
		goto L555
	}
L549:
	;
	goto L550
L550:
	;
	v3770 = int32(1)
	if v3414&v3770 == int32(0) {
		goto L546
	} else {
		goto L577
	}
L551:
	;
	if v3747&int32(1) == int32(0) {
		goto L546
	} else {
		goto L576
	}
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L1
	} else {
		goto L573
	}
L553:
	;
	v3716 = *(*int32)(unsafe.Add(mBase, uint32(v3715)))
	v3717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3716)+21)))
	if v3717 != int32(102) {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	v3715 = v3703 + v3338<<(uint(int32(2))%32)
	goto L553
L555:
	;
	goto L556
L556:
	;
	v3707 = *(*int32)(unsafe.Add(mBase, uint32(v3328)+4))
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3707)+52))
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+12))
	v3715 = v3709 + v3338<<(uint(int32(2))%32) - int32(4)
	goto L553
L557:
	;
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+16))
	v3722 = F_table_open(m, v3720, int32(0))
	mBase = m.M
	v3723 = m.ExcPending
	if v3723 != 0 {
		goto L1
	} else {
		goto L560
	}
L558:
	;
	v3747 = v3702
	goto L559
L559:
	;
	m.G0 = v3700 + int32(16)
	goto L551
L560:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v3722)+76))
	switch v3345 - int32(2) {
	case 0:
		goto L563
	case 1:
		goto L564
	case 2:
		goto L562
	case 3:
		v3742 = v3702
		goto L561
	default:
		goto L552
	}
L561:
	;
	F_relation_close(m, v3722, int32(0))
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		goto L1
	} else {
		goto L572
	}
L562:
	;
	if v3724 == int32(0) {
		v3742 = v3702
		goto L561
	} else {
		goto L571
	}
L563:
	;
	if v3724 == int32(0) {
		v3742 = v3702
		goto L561
	} else {
		goto L566
	}
L564:
	;
	if v3724 == int32(0) {
		v3742 = v3702
		goto L561
	} else {
		goto L565
	}
L565:
	;
	v3729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3724)+25)))
	v3742 = v3729
	goto L561
L566:
	;
	v3732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3724)+26)))
	if v3732 == int32(0) {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3724)+27)))
	if v3735 != int32(1) {
		v3742 = v3702
		goto L561
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	v3742 = int32(1)
	goto L561
L570:
	;
	goto L569
L571:
	;
	v3741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3724)+28)))
	v3742 = v3741
	goto L561
L572:
	;
	v3747 = v3742
	goto L559
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3700))) = v3345
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_26), v3700)
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_14), int32(2576), int32(_a_F_create_plan_recurse_28))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L576:
	;
	v3779 = int32(1)
	goto L547
L577:
	;
	v3779 = v3770
	goto L547
L578:
	;
	if v3788 != 0 {
		goto L458
	} else {
		goto L579
	}
L579:
	;
	v3791 = int32(0)
	v3797 = v3791
	v3798 = int32(1)
	v3799 = v3791
	goto L545
L580:
	;
	v3812 = m.T0[v3809].(func(*base.Module, int32, int32, int32, int32) int32)(m, v3328, v3331, v3459, v3411)
	mBase = m.M
	v3813 = m.ExcPending
	if v3813 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	v3836 = v3812
	v3838 = v3802
	v3840 = v3804
	v3841 = v3420
	v3842 = v3805
	v3843 = v3806
	goto L457
L582:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3820 = m.ExcPending
	if v3820 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_29), int32(0))
	mBase = m.M
	v3824 = m.ExcPending
	if v3824 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_30), int32(_a_F_create_plan_recurse_25))
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L586:
	;
	v3836 = v3543
	v3838 = int32(0)
	v3840 = int32(1)
	v3841 = v3832
	v3842 = int32(0)
	v3843 = int32(1)
	goto L457
L587:
	;
	v3850 = v3411 + int32(1)
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v3347)+4))
	if v3850 < v3851 {
		v3411 = v3850
		v3414 = v3838
		v3418 = v3847
		v3419 = v3840
		v3420 = v3841
		v3421 = v3842
		v3422 = v3843
		goto L455
	} else {
		goto L588
	}
L588:
	;
	goto L456
L589:
	;
	v3921 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v3923 = F_palloc0(m, int32(80))
	mBase = m.M
	v3924 = m.ExcPending
	if v3924 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3923))) = int32(376)
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(v3919)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v3923)+72)) = v3921
	v3929 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+56)) = v3929
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+52)) = v3919
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+48)) = v3929
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+44)) = v3927
	v3935 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+4)) = v3935
	v3937 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3923)+8)) = v3937
	v3939 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3923)+16)) = v3939
	v3941 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3923)+24)) = v3941
	v3943 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v3943)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3923)+32)) = v3944
	v3946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3923)+36)) = uint8(v3946)
	v3948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3923)+37)) = uint8(v3948)
	v12392 = v3923
	v12397 = v51
	goto L3
L591:
	;
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3957 = F_create_plan_recurse(m, l0, v3955, int32(1))
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	v3959 = int32(0)
	v3960 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3961 = *(*int32)(unsafe.Add(mBase, uint32(v3960)+4))
	if v3961 == v3959 {
		v4047 = v3959
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v4093 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v4094 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4095 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v4097 = F_palloc0(m, int32(104))
	mBase = m.M
	v4098 = m.ExcPending
	if v4098 != 0 {
		goto L1
	} else {
		goto L608
	}
L594:
	;
	v3964 = *(*int32)(unsafe.Add(mBase, uint32(v3961)+4))
	if v3964 <= int32(0) {
		v4047 = v3959
		goto L593
	} else {
		goto L595
	}
L595:
	;
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(v3960)+8))
	v3970 = v3959
	v3971 = v3950
	v3973 = v4
	goto L596
L596:
	;
	v4016 = *(*int32)(unsafe.Add(mBase, uint32(v3961)+12))
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v4016+v3973<<(uint(int32(2))%32))))
	v4021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v4021 != 0 {
		goto L598
	} else {
		goto L599
	}
L597:
	;
	v4047 = v4039
	goto L593
L598:
	;
	v4022 = F_replace_nestloop_params_mutator(m, v4020, l0)
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L1
	} else {
		goto L601
	}
L599:
	;
	v4024 = v4020
	goto L600
L600:
	;
	v4026 = int32(0)
	v4028 = F_makeTargetEntry(m, v4024, base.I32_extend16_s(v3971), v4026, v4026)
	mBase = m.M
	v4029 = m.ExcPending
	if v4029 != 0 {
		goto L1
	} else {
		goto L602
	}
L601:
	;
	v4024 = v4022
	goto L600
L602:
	;
	if v3967 != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v4035 = *(*int32)(unsafe.Add(mBase, uint32(v3967+v3971<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4028)+16)) = v4035
	goto L605
L604:
	;
	goto L605
L605:
	;
	v4039 = F_lappend(m, v3970, v4028)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	v4042 = v3973 + int32(1)
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v3961)+4))
	if v4042 < v4043 {
		v3970 = v4039
		v3971 = v3971 + int32(1)
		v3973 = v4042
		goto L596
	} else {
		goto L607
	}
L607:
	;
	goto L597
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4097))) = int32(340)
	if v4094 != 0 {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v4101 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+4))
	v4102 = v4101
	goto L611
L610:
	;
	v4102 = v4
	goto L611
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+76)) = v4102
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+72)) = v4095
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+56)) = v3957
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+52)) = v3953
	v4107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+48)) = v4107
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+44)) = v4047
	if v4107 < v4102 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v4113 = F_palloc_mul(m, int32(2), v4102)
	mBase = m.M
	v4114 = m.ExcPending
	if v4114 != 0 {
		goto L1
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v4097)+96)) = v4093
	v4300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+4)) = v4300
	v4302 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4097)+8)) = v4302
	v4304 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v4097)+16)) = v4304
	v4306 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v4097)+24)) = v4306
	v4308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4309 = *(*int32)(unsafe.Add(mBase, uint32(v4308)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+32)) = v4309
	v4311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4097)+36)) = uint8(v4311)
	v4313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4097)+37)) = uint8(v4313)
	v12392 = v4097
	v12397 = v51
	goto L3
L615:
	;
	v4116 = F_palloc_mul(m, int32(4), v4102)
	mBase = m.M
	v4117 = m.ExcPending
	if v4117 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v4119 = F_palloc_mul(m, int32(4), v4102)
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	if v4094 == int32(0) {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+88)) = v4119
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+84)) = v4116
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+80)) = v4113
	goto L614
L619:
	;
	v4123 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+4))
	if v4123 <= int32(0) {
		goto L618
	} else {
		goto L620
	}
L620:
	;
	v4127 = int32(0)
	goto L621
L621:
	;
	v4179 = v4127 << (uint(int32(2)) % 32)
	v4180 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+12))
	v4182 = *(*int32)(unsafe.Add(mBase, uint32(v4179+v4180)))
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+44))
	v4184 = F_get_sortgroupclause_tle(m, v4182, v4183)
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
		goto L1
	} else {
		goto L623
	}
L622:
	;
	goto L618
L623:
	;
	v4186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4113+v4127<<(uint(int32(1))%32)))) = uint16(v4186)
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4179+v4116))) = v4189
	v4192 = *(*int32)(unsafe.Add(mBase, uint32(v4184)+4))
	v4193 = F_exprCollation(m, v4192)
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4179+v4119))) = v4193
	v4197 = v4127 + int32(1)
	v4198 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+4))
	if v4197 < v4198 {
		v4127 = v4197
		goto L621
	} else {
		goto L625
	}
L625:
	;
	goto L622
L626:
	;
	v4449 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v4451 = l2 | int32(4)
	v4452 = F_create_plan_recurse(m, l0, v4449, v4451)
	mBase = m.M
	v4453 = m.ExcPending
	if v4453 != 0 {
		goto L1
	} else {
		goto L641
	}
L627:
	;
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v4320 <= int32(0) {
		v4407 = v4
		goto L626
	} else {
		goto L628
	}
L628:
	;
	v4323 = *(*int32)(unsafe.Add(mBase, uint32(v4315)+8))
	v4327 = int32(1)
	v4329 = v4
	v4330 = v4
	goto L629
L629:
	;
	v4372 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v4372+v4329<<(uint(int32(2))%32))))
	v4377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v4377 != 0 {
		goto L631
	} else {
		goto L632
	}
L630:
	;
	v4407 = v4395
	goto L626
L631:
	;
	v4378 = F_replace_nestloop_params_mutator(m, v4376, l0)
	mBase = m.M
	v4379 = m.ExcPending
	if v4379 != 0 {
		goto L1
	} else {
		goto L634
	}
L632:
	;
	v4380 = v4376
	goto L633
L633:
	;
	v4382 = int32(0)
	v4384 = F_makeTargetEntry(m, v4380, base.I32_extend16_s(v4327), v4382, v4382)
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L1
	} else {
		goto L635
	}
L634:
	;
	v4380 = v4378
	goto L633
L635:
	;
	if v4323 != 0 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v4323+v4327<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4384)+16)) = v4391
	goto L638
L637:
	;
	goto L638
L638:
	;
	v4395 = F_lappend(m, v4330, v4384)
	mBase = m.M
	v4396 = m.ExcPending
	if v4396 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	v4398 = v4329 + int32(1)
	v4399 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v4398 < v4399 {
		v4327 = v4327 + int32(1)
		v4329 = v4398
		v4330 = v4395
		goto L629
	} else {
		goto L640
	}
L640:
	;
	goto L630
L641:
	;
	v4454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v4455 = F_create_plan_recurse(m, l0, v4454, v4451)
	mBase = m.M
	v4456 = m.ExcPending
	if v4456 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	v4457 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	v4458 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v4460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4462 = F_palloc0(m, int32(112))
	mBase = m.M
	v4463 = m.ExcPending
	if v4463 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4462))) = int32(375)
	if v4458 != 0 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v4466 = *(*int32)(unsafe.Add(mBase, uint32(v4458)+4))
	v4467 = v4466
	goto L646
L645:
	;
	v4467 = v4
	goto L646
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+56)) = v4455
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+52)) = v4452
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+44)) = v4407
	v4474 = F_palloc_mul(m, int32(2), v4467)
	mBase = m.M
	v4475 = m.ExcPending
	if v4475 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	v4477 = F_palloc_mul(m, int32(4), v4467)
	mBase = m.M
	v4478 = m.ExcPending
	if v4478 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	v4480 = F_palloc_mul(m, int32(4), v4467)
	mBase = m.M
	v4481 = m.ExcPending
	if v4481 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	v4483 = F_palloc_mul(m, int32(1), v4467)
	mBase = m.M
	v4484 = m.ExcPending
	if v4484 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	if v4458 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v4462)+104)) = v4457
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+96)) = v4483
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+92)) = v4480
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+88)) = v4477
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+84)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+80)) = v4467
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+76)) = v4459
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+72)) = v4460
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+4)) = v4629
	v4631 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4462)+8)) = v4631
	v4633 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v4462)+16)) = v4633
	v4635 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v4462)+24)) = v4635
	v4637 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4638 = *(*int32)(unsafe.Add(mBase, uint32(v4637)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4462)+32)) = v4638
	v4640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4462)+36)) = uint8(v4640)
	v4642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4462)+37)) = uint8(v4642)
	v12392 = v4462
	v12397 = v51
	goto L3
L652:
	;
	v4487 = *(*int32)(unsafe.Add(mBase, uint32(v4458)+4))
	if v4487 <= int32(0) {
		goto L651
	} else {
		goto L653
	}
L653:
	;
	if v4459 == int32(1) {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	v4494 = int32(8)
	goto L656
L655:
	;
	v4494 = int32(12)
	goto L656
L656:
	;
	v4496 = int32(0)
	goto L657
L657:
	;
	v4548 = v4496 << (uint(int32(2)) % 32)
	v4549 = *(*int32)(unsafe.Add(mBase, uint32(v4458)+12))
	v4551 = *(*int32)(unsafe.Add(mBase, uint32(v4548+v4549)))
	v4552 = *(*int32)(unsafe.Add(mBase, uint32(v4462)+44))
	v4553 = F_get_sortgroupclause_tle(m, v4551, v4552)
	mBase = m.M
	v4554 = m.ExcPending
	if v4554 != 0 {
		goto L1
	} else {
		goto L659
	}
L658:
	;
	goto L651
L659:
	;
	v4555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4553)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4474+v4496<<(uint(int32(1))%32)))) = uint16(v4555)
	v4559 = *(*int32)(unsafe.Add(mBase, uint32(v4551+v4494)))
	*(*int32)(unsafe.Add(mBase, uint32(v4477+v4548))) = v4559
	v4562 = *(*int32)(unsafe.Add(mBase, uint32(v4553)+4))
	v4563 = F_exprCollation(m, v4562)
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4480+v4548))) = v4563
	v4567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4551)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4496+v4483))) = uint8(v4567)
	v4570 = v4496 + int32(1)
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(v4458)+4))
	if v4570 < v4571 {
		v4496 = v4570
		goto L657
	} else {
		goto L661
	}
L661:
	;
	goto L658
L662:
	;
	v4646 = *(*int32)(unsafe.Add(mBase, uint32(v4645)+4))
	v4647 = v4646
	goto L664
L663:
	;
	v4647 = v4
	goto L664
L664:
	;
	v4648 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+16))
	if v4648 != 0 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v4649 = *(*int32)(unsafe.Add(mBase, uint32(v4648)+4))
	v4650 = v4649
	goto L667
L666:
	;
	v4650 = v4
	goto L667
L667:
	;
	v4651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v4653 = F_create_plan_recurse(m, l0, v4651, int32(6))
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	v4655 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v4655)+4))
	if v4656 == int32(0) {
		v4748 = v4
		goto L669
	} else {
		goto L670
	}
L669:
	;
	v4791 = F_palloc_mul(m, int32(2), v4647)
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L1
	} else {
		goto L684
	}
L670:
	;
	v4660 = *(*int32)(unsafe.Add(mBase, uint32(v4656)+4))
	if v4660 <= int32(0) {
		v4748 = v4
		goto L669
	} else {
		goto L671
	}
L671:
	;
	v4663 = *(*int32)(unsafe.Add(mBase, uint32(v4655)+8))
	v4667 = int32(0)
	v4668 = int32(1)
	v4671 = v4
	goto L672
L672:
	;
	v4713 = *(*int32)(unsafe.Add(mBase, uint32(v4656)+12))
	v4717 = *(*int32)(unsafe.Add(mBase, uint32(v4713+v4667<<(uint(int32(2))%32))))
	v4718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v4718 != 0 {
		goto L674
	} else {
		goto L675
	}
L673:
	;
	v4748 = v4736
	goto L669
L674:
	;
	v4719 = F_replace_nestloop_params_mutator(m, v4717, l0)
	mBase = m.M
	v4720 = m.ExcPending
	if v4720 != 0 {
		goto L1
	} else {
		goto L677
	}
L675:
	;
	v4721 = v4717
	goto L676
L676:
	;
	v4723 = int32(0)
	v4725 = F_makeTargetEntry(m, v4721, base.I32_extend16_s(v4668), v4723, v4723)
	mBase = m.M
	v4726 = m.ExcPending
	if v4726 != 0 {
		goto L1
	} else {
		goto L678
	}
L677:
	;
	v4721 = v4719
	goto L676
L678:
	;
	if v4663 != 0 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v4732 = *(*int32)(unsafe.Add(mBase, uint32(v4663+v4668<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4725)+16)) = v4732
	goto L681
L680:
	;
	goto L681
L681:
	;
	v4736 = F_lappend(m, v4671, v4725)
	mBase = m.M
	v4737 = m.ExcPending
	if v4737 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	v4739 = v4667 + int32(1)
	v4740 = *(*int32)(unsafe.Add(mBase, uint32(v4656)+4))
	if v4739 < v4740 {
		v4667 = v4739
		v4668 = v4668 + int32(1)
		v4671 = v4736
		goto L672
	} else {
		goto L683
	}
L683:
	;
	goto L673
L684:
	;
	v4794 = F_palloc_mul(m, int32(4), v4647)
	mBase = m.M
	v4795 = m.ExcPending
	if v4795 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	v4797 = F_palloc_mul(m, int32(4), v4647)
	mBase = m.M
	v4798 = m.ExcPending
	if v4798 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+12))
	if v4799 == int32(0) {
		v4882 = v4
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v4927 = F_palloc_mul(m, int32(2), v4650)
	mBase = m.M
	v4928 = m.ExcPending
	if v4928 != 0 {
		goto L1
	} else {
		goto L695
	}
L688:
	;
	v4802 = *(*int32)(unsafe.Add(mBase, uint32(v4799)+4))
	if v4802 <= int32(0) {
		v4882 = v4
		goto L687
	} else {
		goto L689
	}
L689:
	;
	v4809 = v4
	goto L690
L690:
	;
	v4857 = v4809 << (uint(int32(2)) % 32)
	v4858 = *(*int32)(unsafe.Add(mBase, uint32(v4799)+12))
	v4860 = *(*int32)(unsafe.Add(mBase, uint32(v4857+v4858)))
	v4861 = *(*int32)(unsafe.Add(mBase, uint32(v4653)+44))
	v4862 = F_get_sortgroupclause_tle(m, v4860, v4861)
	mBase = m.M
	v4863 = m.ExcPending
	if v4863 != 0 {
		goto L1
	} else {
		goto L692
	}
L691:
	;
	v4882 = v4875
	goto L687
L692:
	;
	v4864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4862)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4791+v4809<<(uint(int32(1))%32)))) = uint16(v4864)
	v4867 = *(*int32)(unsafe.Add(mBase, uint32(v4860)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4857+v4794))) = v4867
	v4870 = *(*int32)(unsafe.Add(mBase, uint32(v4862)+4))
	v4871 = F_exprCollation(m, v4870)
	mBase = m.M
	v4872 = m.ExcPending
	if v4872 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4857+v4797))) = v4871
	v4875 = v4809 + int32(1)
	v4876 = *(*int32)(unsafe.Add(mBase, uint32(v4799)+4))
	if v4875 < v4876 {
		v4809 = v4875
		goto L690
	} else {
		goto L694
	}
L694:
	;
	goto L691
L695:
	;
	v4930 = F_palloc_mul(m, int32(4), v4650)
	mBase = m.M
	v4931 = m.ExcPending
	if v4931 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	v4933 = F_palloc_mul(m, int32(4), v4650)
	mBase = m.M
	v4934 = m.ExcPending
	if v4934 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	v4935 = int32(0)
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+16))
	if v4936 == v4935 {
		v5015 = v4935
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v5063 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v5064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+88)))
	v5065 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v5067 = F_palloc0(m, int32(152))
	mBase = m.M
	v5068 = m.ExcPending
	if v5068 != 0 {
		goto L1
	} else {
		goto L706
	}
L699:
	;
	v4939 = *(*int32)(unsafe.Add(mBase, uint32(v4936)+4))
	if v4939 <= int32(0) {
		v5015 = v4935
		goto L698
	} else {
		goto L700
	}
L700:
	;
	v4942 = v4935
	goto L701
L701:
	;
	v4994 = v4942 << (uint(int32(2)) % 32)
	v4995 = *(*int32)(unsafe.Add(mBase, uint32(v4936)+12))
	v4997 = *(*int32)(unsafe.Add(mBase, uint32(v4994+v4995)))
	v4998 = *(*int32)(unsafe.Add(mBase, uint32(v4653)+44))
	v4999 = F_get_sortgroupclause_tle(m, v4997, v4998)
	mBase = m.M
	v5000 = m.ExcPending
	if v5000 != 0 {
		goto L1
	} else {
		goto L703
	}
L702:
	;
	v5015 = v5012
	goto L698
L703:
	;
	v5001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4999)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4927+v4942<<(uint(int32(1))%32)))) = uint16(v5001)
	v5004 = *(*int32)(unsafe.Add(mBase, uint32(v4997)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4994+v4930))) = v5004
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v4999)+4))
	v5008 = F_exprCollation(m, v5007)
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4994+v4933))) = v5008
	v5012 = v4942 + int32(1)
	v5013 = *(*int32)(unsafe.Add(mBase, uint32(v4936)+4))
	if v5012 < v5013 {
		v4942 = v5012
		goto L701
	} else {
		goto L705
	}
L705:
	;
	goto L702
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5067))) = int32(370)
	v5071 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+72)) = v5071
	v5073 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+108)) = v4933
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+104)) = v4930
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+100)) = v4927
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+96)) = v5015
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+92)) = v4797
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+88)) = v4794
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+84)) = v4791
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+80)) = v4882
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+76)) = v5073
	v5083 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+112)) = v5083
	v5085 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+116)) = v5085
	v5087 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+128)) = v5065
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+124)) = v5065
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+120)) = v5087
	v5091 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+132)) = v5091
	v5093 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+136)) = v5093
	v5095 = *(*int32)(unsafe.Add(mBase, uint32(v4644)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+140)) = v5095
	v5097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4644)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5067)+144)) = uint8(v5097)
	v5099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4644)+45)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5067)+146)) = uint8(v5064)
	*(*uint8)(unsafe.Add(mBase, uint32(v5067)+145)) = uint8(v5099)
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+52)) = v4653
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+44)) = v4748
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+48)) = v5063
	v5107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+4)) = v5107
	v5109 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v5067)+8)) = v5109
	v5111 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v5067)+16)) = v5111
	v5113 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v5067)+24)) = v5113
	v5115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5116 = *(*int32)(unsafe.Add(mBase, uint32(v5115)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5067)+32)) = v5116
	v5118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5067)+36)) = uint8(v5118)
	v5120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5067)+37)) = uint8(v5120)
	v12392 = v5067
	v12397 = v51
	goto L3
L707:
	;
	v5125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v5126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v5128 = F_create_plan_recurse(m, l0, v5126, int32(4))
	mBase = m.M
	v5129 = m.ExcPending
	if v5129 != 0 {
		goto L1
	} else {
		goto L710
	}
L708:
	;
	goto L709
L709:
	;
	v6560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6562 = F_create_plan_recurse(m, l0, v6560, int32(4))
	mBase = m.M
	v6563 = m.ExcPending
	if v6563 != 0 {
		goto L1
	} else {
		goto L849
	}
L710:
	;
	v5130 = int32(2)
	v5131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v5131 == int32(0) {
		v5388 = v5130
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v5434 = F_palloc0(m, v5388)
	mBase = m.M
	v5435 = m.ExcPending
	if v5435 != 0 {
		goto L1
	} else {
		goto L743
	}
L712:
	;
	v5134 = *(*int32)(unsafe.Add(mBase, uint32(v5131)+4))
	if v5134 <= int32(0) {
		v5388 = v5130
		goto L711
	} else {
		goto L713
	}
L713:
	;
	v5137 = int32(0)
	if v5137 < v5134 {
		goto L714
	} else {
		goto L715
	}
L714:
	;
	v5140 = v5134
	goto L716
L715:
	;
	v5140 = v5137
	goto L716
L716:
	;
	v5142 = v5140 & int32(3)
	v5143 = int32(0)
	if int32(4) <= v5134 {
		goto L718
	} else {
		goto L719
	}
L717:
	;
	v5388 = v5336<<(uint(int32(1))%32) + int32(2)
	goto L711
L718:
	;
	v5148 = *(*int32)(unsafe.Add(mBase, uint32(v5131)+12))
	v5152 = v5143
	v5155 = v4
	v5156 = int32(0)
	goto L721
L719:
	;
	v5226 = v5143
	v5229 = v4
	goto L720
L720:
	;
	v5272 = *(*int32)(unsafe.Add(mBase, uint32(v5131)+12))
	v5276 = v5226
	v5278 = int32(0)
	v5279 = v5229
	goto L737
L721:
	;
	v5200 = v5148 + v5155<<(uint(int32(2))%32)
	v5201 = *(*int32)(unsafe.Add(mBase, uint32(v5200)+12))
	v5202 = *(*int32)(unsafe.Add(mBase, uint32(v5201)+4))
	v5203 = *(*int32)(unsafe.Add(mBase, uint32(v5200)+8))
	v5204 = *(*int32)(unsafe.Add(mBase, uint32(v5203)+4))
	v5205 = *(*int32)(unsafe.Add(mBase, uint32(v5200)+4))
	v5206 = *(*int32)(unsafe.Add(mBase, uint32(v5205)+4))
	v5207 = *(*int32)(unsafe.Add(mBase, uint32(v5200)))
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(v5207)+4))
	if base.Ui32(v5152) < base.Ui32(v5208) {
		goto L723
	} else {
		goto L724
	}
L722:
	;
	if v5142 == int32(0) {
		v5336 = v5216
		goto L717
	} else {
		goto L736
	}
L723:
	;
	v5210 = v5208
	goto L725
L724:
	;
	v5210 = v5152
	goto L725
L725:
	;
	if base.Ui32(v5210) < base.Ui32(v5206) {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	v5212 = v5206
	goto L728
L727:
	;
	v5212 = v5210
	goto L728
L728:
	;
	if base.Ui32(v5212) < base.Ui32(v5204) {
		goto L729
	} else {
		goto L730
	}
L729:
	;
	v5214 = v5204
	goto L731
L730:
	;
	v5214 = v5212
	goto L731
L731:
	;
	if base.Ui32(v5214) < base.Ui32(v5202) {
		goto L732
	} else {
		goto L733
	}
L732:
	;
	v5216 = v5202
	goto L734
L733:
	;
	v5216 = v5214
	goto L734
L734:
	;
	v5217 = int32(4)
	v5218 = v5155 + v5217
	v5220 = v5156 + v5217
	if v5220 != v5140&int32(2147483644) {
		v5152 = v5216
		v5155 = v5218
		v5156 = v5220
		goto L721
	} else {
		goto L735
	}
L735:
	;
	goto L722
L736:
	;
	v5226 = v5216
	v5229 = v5218
	goto L720
L737:
	;
	v5325 = *(*int32)(unsafe.Add(mBase, uint32(v5272+v5279<<(uint(int32(2))%32))))
	v5326 = *(*int32)(unsafe.Add(mBase, uint32(v5325)+4))
	if base.Ui32(v5276) < base.Ui32(v5326) {
		goto L739
	} else {
		goto L740
	}
L738:
	;
	v5336 = v5328
	goto L717
L739:
	;
	v5328 = v5326
	goto L741
L740:
	;
	v5328 = v5276
	goto L741
L741:
	;
	v5329 = int32(1)
	v5332 = v5278 + v5329
	if v5332 != v5142 {
		v5276 = v5328
		v5278 = v5332
		v5279 = v5279 + v5329
		goto L737
	} else {
		goto L742
	}
L742:
	;
	goto L738
L743:
	;
	v5436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v5436 == int32(0) {
		goto L744
	} else {
		goto L745
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v5434
	if v5125 == int32(0) {
		v6194 = v5434
		v6204 = v4
		goto L751
	} else {
		goto L752
	}
L745:
	;
	v5439 = *(*int32)(unsafe.Add(mBase, uint32(v5436)+4))
	if v5439 <= int32(0) {
		goto L744
	} else {
		goto L746
	}
L746:
	;
	v5445 = int32(0)
	goto L747
L747:
	;
	v5491 = *(*int32)(unsafe.Add(mBase, uint32(v5436)+12))
	v5495 = *(*int32)(unsafe.Add(mBase, uint32(v5491+v5445<<(uint(int32(2))%32))))
	v5496 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+44))
	v5497 = F_get_sortgroupclause_tle(m, v5495, v5496)
	mBase = m.M
	v5498 = m.ExcPending
	if v5498 != 0 {
		goto L1
	} else {
		goto L749
	}
L748:
	;
	goto L744
L749:
	;
	v5499 = *(*int32)(unsafe.Add(mBase, uint32(v5495)+4))
	v5500 = int32(1)
	v5503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5497)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5434+v5499<<(uint(v5500)%32)))) = uint16(v5503)
	v5506 = v5445 + v5500
	v5507 = *(*int32)(unsafe.Add(mBase, uint32(v5436)+4))
	if v5506 < v5507 {
		v5445 = v5506
		goto L747
	} else {
		goto L750
	}
L750:
	;
	goto L748
L751:
	;
	v6237 = *(*int32)(unsafe.Add(mBase, uint32(v5125)+12))
	v6238 = *(*int32)(unsafe.Add(mBase, uint32(v6237)))
	v6239 = *(*int32)(unsafe.Add(mBase, uint32(v6238)+4))
	if v6239 == int32(0) {
		goto L819
	} else {
		goto L820
	}
L752:
	;
	v5560 = *(*int32)(unsafe.Add(mBase, uint32(v5125)+4))
	if v5560 < int32(2) {
		v6194 = v5434
		v6204 = v4
		goto L751
	} else {
		goto L753
	}
L753:
	;
	v5563 = *(*int32)(unsafe.Add(mBase, uint32(v5125)+12))
	v5564 = *(*int32)(unsafe.Add(mBase, uint32(v5563)))
	v5565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5564)+25)))
	v5576 = int32(1)
	v5580 = v5565
	v5582 = v4
	goto L754
L754:
	;
	v5615 = *(*int32)(unsafe.Add(mBase, uint32(v5125)+12))
	v5619 = *(*int32)(unsafe.Add(mBase, uint32(v5615+v5576<<(uint(int32(2))%32))))
	v5620 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+4))
	if v5620 == int32(0) {
		goto L757
	} else {
		goto L758
	}
L755:
	;
	v6188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v6194 = v6188
	v6204 = v6182
	goto L751
L756:
	;
	v5750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5619)+25)))
	if (v5750|v5580)&int32(1) == int32(0) {
		goto L770
	} else {
		goto L771
	}
L757:
	;
	v5625 = F_palloc0_mul(m, int32(2), int32(0))
	mBase = m.M
	v5626 = m.ExcPending
	if v5626 != 0 {
		goto L1
	} else {
		goto L760
	}
L758:
	;
	goto L759
L759:
	;
	v5627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v5629 = *(*int32)(unsafe.Add(mBase, uint32(v5620)+4))
	v5630 = F_palloc0_mul(m, int32(2), v5629)
	mBase = m.M
	v5631 = m.ExcPending
	if v5631 != 0 {
		goto L1
	} else {
		goto L761
	}
L760:
	;
	v5706 = v5625
	goto L756
L761:
	;
	v5632 = int32(0)
	v5633 = *(*int32)(unsafe.Add(mBase, uint32(v5620)+4))
	if v5633 <= v5632 {
		v5706 = v5630
		goto L756
	} else {
		goto L762
	}
L762:
	;
	v5638 = v5632
	goto L763
L763:
	;
	v5684 = int32(1)
	v5687 = *(*int32)(unsafe.Add(mBase, uint32(v5620)+12))
	v5691 = *(*int32)(unsafe.Add(mBase, uint32(v5687+v5638<<(uint(int32(2))%32))))
	v5692 = *(*int32)(unsafe.Add(mBase, uint32(v5691)+4))
	v5696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5627+v5692<<(uint(v5684)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5630+v5638<<(uint(v5684)%32)))) = uint16(v5696)
	v5699 = v5638 + v5684
	v5700 = *(*int32)(unsafe.Add(mBase, uint32(v5620)+4))
	if v5699 < v5700 {
		v5638 = v5699
		goto L763
	} else {
		goto L765
	}
L764:
	;
	v5706 = v5630
	goto L756
L765:
	;
	goto L764
L766:
	;
	if v6094 != 0 {
		goto L807
	} else {
		goto L808
	}
L767:
	;
	v6086 = int32(0)
	v6087 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+8))
	v6088 = *(*int32)(unsafe.Add(mBase, uint32(v6087)+12))
	v6089 = *(*int32)(unsafe.Add(mBase, uint32(v6088)))
	v6094 = v6089
	v6097 = v6043
	v6105 = v6086
	v6140 = base.B2i32(v6089 != v6086)
	goto L766
L768:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6028 = m.ExcPending
	if v6028 != 0 {
		goto L1
	} else {
		goto L804
	}
L769:
	;
	v6021 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+8))
	v6022 = *(*int32)(unsafe.Add(mBase, uint32(v6021)+12))
	v6023 = *(*int32)(unsafe.Add(mBase, uint32(v6022)))
	v6094 = v6023
	v6097 = v5978
	v6105 = v5580
	v6140 = int32(2)
	goto L766
L770:
	;
	v5756 = int32(0)
	v5758 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+44))
	v5759 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+4))
	if v5759 != 0 {
		goto L773
	} else {
		goto L774
	}
L771:
	;
	goto L772
L772:
	;
	v5968 = int32(0)
	if v5750&int32(1) == v5968 {
		v6043 = v5968
		goto L767
	} else {
		goto L803
	}
L773:
	;
	v5760 = *(*int32)(unsafe.Add(mBase, uint32(v5759)+4))
	v5761 = v5760
	goto L775
L774:
	;
	v5761 = v5756
	goto L775
L775:
	;
	v5764 = F_palloc(m, v5761<<(uint(int32(1))%32))
	mBase = m.M
	v5765 = m.ExcPending
	if v5765 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	v5767 = v5761 << (uint(int32(2)) % 32)
	v5768 = F_palloc(m, v5767)
	mBase = m.M
	v5769 = m.ExcPending
	if v5769 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v5770 = F_palloc(m, v5767)
	mBase = m.M
	v5771 = m.ExcPending
	if v5771 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	v5772 = F_palloc(m, v5761)
	mBase = m.M
	v5773 = m.ExcPending
	if v5773 != 0 {
		goto L1
	} else {
		goto L779
	}
L779:
	;
	if v5759 == int32(0) {
		v5897 = v5756
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v5944 = F_palloc0(m, int32(96))
	mBase = m.M
	v5945 = m.ExcPending
	if v5945 != 0 {
		goto L1
	} else {
		goto L801
	}
L781:
	;
	v5776 = *(*int32)(unsafe.Add(mBase, uint32(v5759)+4))
	if v5776 <= int32(0) {
		v5897 = v5756
		goto L780
	} else {
		goto L782
	}
L782:
	;
	v5781 = v5756
	goto L783
L783:
	;
	v5828 = v5781 << (uint(int32(2)) % 32)
	v5829 = *(*int32)(unsafe.Add(mBase, uint32(v5759)+12))
	v5831 = *(*int32)(unsafe.Add(mBase, uint32(v5828+v5829)))
	v5833 = v5781 << (uint(int32(1)) % 32)
	v5835 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5706+v5833))))
	if v5758 != 0 {
		goto L787
	} else {
		goto L788
	}
L784:
	;
	v5897 = v5892
	goto L780
L785:
	;
	if v5873 == int32(0) {
		goto L768
	} else {
		goto L798
	}
L786:
	;
	goto L785
L787:
	;
	v5839 = *(*int32)(unsafe.Add(mBase, uint32(v5758)+4))
	if v5839 <= int32(0) {
		v5873 = int32(0)
		goto L786
	} else {
		goto L790
	}
L788:
	;
	goto L789
L789:
	;
	v5873 = int32(0)
	goto L786
L790:
	;
	v5842 = int32(0)
	if v5842 < v5839 {
		goto L791
	} else {
		goto L792
	}
L791:
	;
	v5845 = v5839
	goto L793
L792:
	;
	v5845 = v5842
	goto L793
L793:
	;
	v5846 = *(*int32)(unsafe.Add(mBase, uint32(v5758)+12))
	v5850 = int32(0)
	goto L794
L794:
	;
	v5858 = *(*int32)(unsafe.Add(mBase, uint32(v5846+v5850<<(uint(int32(2))%32))))
	v5859 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5858)+8)))
	if v5859 == v5835&int32(_a_F_create_plan_recurse_31) {
		v5873 = v5858
		goto L786
	} else {
		goto L796
	}
L795:
	;
	goto L789
L796:
	;
	v5862 = v5850 + int32(1)
	if v5862 != v5845 {
		v5850 = v5862
		goto L794
	} else {
		goto L797
	}
L797:
	;
	goto L795
L798:
	;
	v5878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5873)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5764+v5833))) = uint16(v5878)
	v5881 = *(*int32)(unsafe.Add(mBase, uint32(v5831)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5828+v5768))) = v5881
	v5884 = *(*int32)(unsafe.Add(mBase, uint32(v5873)+4))
	v5885 = F_exprCollation(m, v5884)
	mBase = m.M
	v5886 = m.ExcPending
	if v5886 != 0 {
		goto L1
	} else {
		goto L799
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5828+v5770))) = v5885
	v5889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5831)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5781+v5772))) = uint8(v5889)
	v5892 = v5781 + int32(1)
	v5893 = *(*int32)(unsafe.Add(mBase, uint32(v5759)+4))
	if v5892 < v5893 {
		v5781 = v5892
		goto L783
	} else {
		goto L800
	}
L800:
	;
	goto L784
L801:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5944))) = int32(366)
	v5948 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+44)) = v5948
	v5950 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+4))
	v5952 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+88)) = v5772
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+84)) = v5770
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+80)) = v5768
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+76)) = v5764
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+72)) = v5897
	v5958 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+56)) = v5958
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+52)) = v5128
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+48)) = v5958
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+4)) = v5950 + (v5952 ^ int32(1))
	v5967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5619)+25)))
	if v5967 != 0 {
		v5978 = v5944
		goto L769
	} else {
		goto L802
	}
L802:
	;
	v6043 = v5944
	goto L767
L803:
	;
	v5978 = v5968
	goto L769
L804:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_32), int32(0))
	mBase = m.M
	v6032 = m.ExcPending
	if v6032 != 0 {
		goto L1
	} else {
		goto L805
	}
L805:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_33), int32(_a_F_create_plan_recurse_34))
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L807:
	;
	v6141 = *(*int32)(unsafe.Add(mBase, uint32(v6094)+4))
	v6143 = v6141
	goto L809
L808:
	;
	v6143 = int32(0)
	goto L809
L809:
	;
	v6144 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+4))
	v6145 = F_extract_grouping_ops(m, v6144)
	mBase = m.M
	v6146 = m.ExcPending
	if v6146 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	v6147 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+4))
	v6148 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+44))
	v6149 = F_extract_grouping_collations(m, v6147, v6148)
	mBase = m.M
	v6150 = m.ExcPending
	if v6150 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	v6151 = *(*float64)(unsafe.Add(mBase, uint32(v5619)+16))
	v6152 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v6153 = *(*int32)(unsafe.Add(mBase, uint32(v5619)+8))
	v6155 = F_palloc0(m, int32(128))
	mBase = m.M
	v6156 = m.ExcPending
	if v6156 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	v6157 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+120)) = v6157
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+116)) = v6153
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+112)) = v6157
	*(*int64)(unsafe.Add(mBase, uint32(v6155)+104)) = v6152
	*(*float64)(unsafe.Add(mBase, uint32(v6155)+96)) = v6151
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+92)) = v6149
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+88)) = v6145
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+84)) = v5706
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+80)) = v6143
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+76)) = v6157
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+72)) = v6140
	*(*int32)(unsafe.Add(mBase, uint32(v6155))) = int32(369)
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+56)) = v6157
	*(*int32)(unsafe.Add(mBase, uint32(v6155)+52)) = v6097
	*(*int64)(unsafe.Add(mBase, uint32(v6155)+44)) = int64(0)
	if v6097 != 0 {
		goto L813
	} else {
		goto L814
	}
L813:
	;
	v6178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6097)+52)) = v6178
	*(*int32)(unsafe.Add(mBase, uint32(v6097)+44)) = v6178
	goto L815
L814:
	;
	goto L815
L815:
	;
	v6182 = F_lappend(m, v5582, v6155)
	mBase = m.M
	v6183 = m.ExcPending
	if v6183 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	v6185 = v5576 + int32(1)
	v6186 = *(*int32)(unsafe.Add(mBase, uint32(v5125)+4))
	if v6185 < v6186 {
		v5576 = v6185
		v5580 = v6105
		v5582 = v6182
		goto L754
	} else {
		goto L817
	}
L817:
	;
	goto L755
L818:
	;
	v6368 = int32(0)
	v6370 = *(*int32)(unsafe.Add(mBase, uint32(v6238)+8))
	v6371 = *(*int32)(unsafe.Add(mBase, uint32(v6370)+12))
	v6372 = *(*int32)(unsafe.Add(mBase, uint32(v6371)))
	if v6372 != 0 {
		goto L828
	} else {
		goto L829
	}
L819:
	;
	v6244 = F_palloc0_mul(m, int32(2), int32(0))
	mBase = m.M
	v6245 = m.ExcPending
	if v6245 != 0 {
		goto L1
	} else {
		goto L822
	}
L820:
	;
	goto L821
L821:
	;
	v6247 = *(*int32)(unsafe.Add(mBase, uint32(v6239)+4))
	v6248 = F_palloc0_mul(m, int32(2), v6247)
	mBase = m.M
	v6249 = m.ExcPending
	if v6249 != 0 {
		goto L1
	} else {
		goto L823
	}
L822:
	;
	v6329 = v6244
	goto L818
L823:
	;
	v6250 = *(*int32)(unsafe.Add(mBase, uint32(v6239)+4))
	if v6250 <= int32(0) {
		v6329 = v6248
		goto L818
	} else {
		goto L824
	}
L824:
	;
	v6256 = int32(0)
	goto L825
L825:
	;
	v6302 = int32(1)
	v6305 = *(*int32)(unsafe.Add(mBase, uint32(v6239)+12))
	v6309 = *(*int32)(unsafe.Add(mBase, uint32(v6305+v6256<<(uint(int32(2))%32))))
	v6310 = *(*int32)(unsafe.Add(mBase, uint32(v6309)+4))
	v6314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6194+v6310<<(uint(v6302)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v6248+v6256<<(uint(v6302)%32)))) = uint16(v6314)
	v6317 = v6256 + v6302
	v6318 = *(*int32)(unsafe.Add(mBase, uint32(v6239)+4))
	if v6317 < v6318 {
		v6256 = v6317
		goto L825
	} else {
		goto L827
	}
L826:
	;
	v6329 = v6248
	goto L818
L827:
	;
	goto L826
L828:
	;
	v6373 = *(*int32)(unsafe.Add(mBase, uint32(v6372)+4))
	v6374 = v6373
	goto L830
L829:
	;
	v6374 = v6368
	goto L830
L830:
	;
	v6375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6376 = *(*int32)(unsafe.Add(mBase, uint32(v6375)+4))
	if v6376 == int32(0) {
		v6468 = v6368
		goto L831
	} else {
		goto L832
	}
L831:
	;
	v6510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v6511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6512 = *(*int32)(unsafe.Add(mBase, uint32(v6238)+4))
	v6513 = F_extract_grouping_ops(m, v6512)
	mBase = m.M
	v6514 = m.ExcPending
	if v6514 != 0 {
		goto L1
	} else {
		goto L846
	}
L832:
	;
	v6380 = *(*int32)(unsafe.Add(mBase, uint32(v6376)+4))
	if v6380 <= int32(0) {
		v6468 = v6368
		goto L831
	} else {
		goto L833
	}
L833:
	;
	v6383 = *(*int32)(unsafe.Add(mBase, uint32(v6375)+8))
	v6387 = int32(1)
	v6389 = int32(0)
	v6391 = v6368
	goto L834
L834:
	;
	v6433 = *(*int32)(unsafe.Add(mBase, uint32(v6376)+12))
	v6437 = *(*int32)(unsafe.Add(mBase, uint32(v6433+v6389<<(uint(int32(2))%32))))
	v6438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6438 != 0 {
		goto L836
	} else {
		goto L837
	}
L835:
	;
	v6468 = v6456
	goto L831
L836:
	;
	v6439 = F_replace_nestloop_params_mutator(m, v6437, l0)
	mBase = m.M
	v6440 = m.ExcPending
	if v6440 != 0 {
		goto L1
	} else {
		goto L839
	}
L837:
	;
	v6441 = v6437
	goto L838
L838:
	;
	v6443 = int32(0)
	v6445 = F_makeTargetEntry(m, v6441, base.I32_extend16_s(v6387), v6443, v6443)
	mBase = m.M
	v6446 = m.ExcPending
	if v6446 != 0 {
		goto L1
	} else {
		goto L840
	}
L839:
	;
	v6441 = v6439
	goto L838
L840:
	;
	if v6383 != 0 {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v6452 = *(*int32)(unsafe.Add(mBase, uint32(v6383+v6387<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6445)+16)) = v6452
	goto L843
L842:
	;
	goto L843
L843:
	;
	v6456 = F_lappend(m, v6391, v6445)
	mBase = m.M
	v6457 = m.ExcPending
	if v6457 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	v6459 = v6389 + int32(1)
	v6460 = *(*int32)(unsafe.Add(mBase, uint32(v6376)+4))
	if v6459 < v6460 {
		v6387 = v6387 + int32(1)
		v6389 = v6459
		v6391 = v6456
		goto L834
	} else {
		goto L845
	}
L845:
	;
	goto L835
L846:
	;
	v6515 = *(*int32)(unsafe.Add(mBase, uint32(v6238)+4))
	v6516 = *(*int32)(unsafe.Add(mBase, uint32(v5128)+44))
	v6517 = F_extract_grouping_collations(m, v6515, v6516)
	mBase = m.M
	v6518 = m.ExcPending
	if v6518 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	v6519 = *(*float64)(unsafe.Add(mBase, uint32(v6238)+16))
	v6520 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v6521 = *(*int32)(unsafe.Add(mBase, uint32(v6238)+8))
	v6523 = F_palloc0(m, int32(128))
	mBase = m.M
	v6524 = m.ExcPending
	if v6524 != 0 {
		goto L1
	} else {
		goto L848
	}
L848:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+120)) = v6204
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+116)) = v6521
	v6527 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+112)) = v6527
	*(*int64)(unsafe.Add(mBase, uint32(v6523)+104)) = v6520
	*(*float64)(unsafe.Add(mBase, uint32(v6523)+96)) = v6519
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+92)) = v6517
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+88)) = v6513
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+84)) = v6329
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+80)) = v6374
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+76)) = v6527
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+72)) = v6511
	*(*int32)(unsafe.Add(mBase, uint32(v6523))) = int32(369)
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+48)) = v6510
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+56)) = v6527
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+52)) = v5128
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+44)) = v6468
	v6545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+4)) = v6545
	v6547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6523)+8)) = v6547
	v6549 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6523)+16)) = v6549
	v6551 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6523)+24)) = v6551
	v6553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6554 = *(*int32)(unsafe.Add(mBase, uint32(v6553)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6523)+32)) = v6554
	v6556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6523)+36)) = uint8(v6556)
	v6558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6523)+37)) = uint8(v6558)
	v12392 = v6523
	v12397 = v51
	goto L3
L849:
	;
	v6564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6565 = *(*int32)(unsafe.Add(mBase, uint32(v6564)+4))
	if v6565 == int32(0) {
		v6661 = v4
		goto L850
	} else {
		goto L851
	}
L850:
	;
	v6698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v6699 = F_order_qual_clauses(m, l0, v6698)
	mBase = m.M
	v6700 = m.ExcPending
	if v6700 != 0 {
		goto L1
	} else {
		goto L865
	}
L851:
	;
	v6569 = *(*int32)(unsafe.Add(mBase, uint32(v6565)+4))
	if v6569 <= int32(0) {
		v6661 = v4
		goto L850
	} else {
		goto L852
	}
L852:
	;
	v6572 = *(*int32)(unsafe.Add(mBase, uint32(v6564)+8))
	v6575 = int32(1)
	v6579 = v4
	v6584 = v4
	goto L853
L853:
	;
	v6621 = *(*int32)(unsafe.Add(mBase, uint32(v6565)+12))
	v6625 = *(*int32)(unsafe.Add(mBase, uint32(v6621+v6579<<(uint(int32(2))%32))))
	v6626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6626 != 0 {
		goto L855
	} else {
		goto L856
	}
L854:
	;
	v6661 = v6644
	goto L850
L855:
	;
	v6627 = F_replace_nestloop_params_mutator(m, v6625, l0)
	mBase = m.M
	v6628 = m.ExcPending
	if v6628 != 0 {
		goto L1
	} else {
		goto L858
	}
L856:
	;
	v6629 = v6625
	goto L857
L857:
	;
	v6631 = int32(0)
	v6633 = F_makeTargetEntry(m, v6629, base.I32_extend16_s(v6575), v6631, v6631)
	mBase = m.M
	v6634 = m.ExcPending
	if v6634 != 0 {
		goto L1
	} else {
		goto L859
	}
L858:
	;
	v6629 = v6627
	goto L857
L859:
	;
	if v6572 != 0 {
		goto L860
	} else {
		goto L861
	}
L860:
	;
	v6640 = *(*int32)(unsafe.Add(mBase, uint32(v6572+v6575<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6633)+16)) = v6640
	goto L862
L861:
	;
	goto L862
L862:
	;
	v6644 = F_lappend(m, v6584, v6633)
	mBase = m.M
	v6645 = m.ExcPending
	if v6645 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	v6647 = v6579 + int32(1)
	v6648 = *(*int32)(unsafe.Add(mBase, uint32(v6565)+4))
	if v6647 < v6648 {
		v6575 = v6575 + int32(1)
		v6579 = v6647
		v6584 = v6644
		goto L853
	} else {
		goto L864
	}
L864:
	;
	goto L854
L865:
	;
	v6701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6702 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v6703 != 0 {
		goto L866
	} else {
		goto L867
	}
L866:
	;
	v6704 = *(*int32)(unsafe.Add(mBase, uint32(v6703)+4))
	v6705 = v6704
	goto L868
L867:
	;
	v6705 = v4
	goto L868
L868:
	;
	v6706 = *(*int32)(unsafe.Add(mBase, uint32(v6562)+44))
	v6707 = F_extract_grouping_cols(m, v6703, v6706)
	mBase = m.M
	v6708 = m.ExcPending
	if v6708 != 0 {
		goto L1
	} else {
		goto L869
	}
L869:
	;
	v6709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v6710 = F_extract_grouping_ops(m, v6709)
	mBase = m.M
	v6711 = m.ExcPending
	if v6711 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	v6712 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v6713 = *(*int32)(unsafe.Add(mBase, uint32(v6562)+44))
	v6714 = F_extract_grouping_collations(m, v6712, v6713)
	mBase = m.M
	v6715 = m.ExcPending
	if v6715 != 0 {
		goto L1
	} else {
		goto L871
	}
L871:
	;
	v6716 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v6717 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+96)))
	v6719 = F_palloc0(m, int32(128))
	mBase = m.M
	v6720 = m.ExcPending
	if v6720 != 0 {
		goto L1
	} else {
		goto L872
	}
L872:
	;
	v6721 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+120)) = v6721
	*(*int64)(unsafe.Add(mBase, uint32(v6719)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6719)+104)) = v6717
	*(*float64)(unsafe.Add(mBase, uint32(v6719)+96)) = v6716
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+92)) = v6714
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+88)) = v6710
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+84)) = v6707
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+80)) = v6705
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+76)) = v6701
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+72)) = v6702
	*(*int32)(unsafe.Add(mBase, uint32(v6719))) = int32(369)
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+48)) = v6699
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+56)) = v6721
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+52)) = v6562
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+44)) = v6661
	v6740 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+4)) = v6740
	v6742 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6719)+8)) = v6742
	v6744 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6719)+16)) = v6744
	v6746 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6719)+24)) = v6746
	v6748 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6749 = *(*int32)(unsafe.Add(mBase, uint32(v6748)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6719)+32)) = v6749
	v6751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6719)+36)) = uint8(v6751)
	v6753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6719)+37)) = uint8(v6753)
	v12392 = v6719
	v12397 = v51
	goto L3
L873:
	;
	v6759 = int32(0)
	v6760 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6761 = *(*int32)(unsafe.Add(mBase, uint32(v6760)+4))
	if v6761 == v6759 {
		v6848 = v6759
		goto L874
	} else {
		goto L875
	}
L874:
	;
	v6894 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6895 = F_order_qual_clauses(m, l0, v6894)
	mBase = m.M
	v6896 = m.ExcPending
	if v6896 != 0 {
		goto L1
	} else {
		goto L889
	}
L875:
	;
	v6765 = *(*int32)(unsafe.Add(mBase, uint32(v6761)+4))
	if v6765 <= int32(0) {
		v6848 = v6759
		goto L874
	} else {
		goto L876
	}
L876:
	;
	v6768 = *(*int32)(unsafe.Add(mBase, uint32(v6760)+8))
	v6771 = v6759
	v6772 = int32(1)
	v6774 = v4
	goto L877
L877:
	;
	v6817 = *(*int32)(unsafe.Add(mBase, uint32(v6761)+12))
	v6821 = *(*int32)(unsafe.Add(mBase, uint32(v6817+v6774<<(uint(int32(2))%32))))
	v6822 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6822 != 0 {
		goto L879
	} else {
		goto L880
	}
L878:
	;
	v6848 = v6840
	goto L874
L879:
	;
	v6823 = F_replace_nestloop_params_mutator(m, v6821, l0)
	mBase = m.M
	v6824 = m.ExcPending
	if v6824 != 0 {
		goto L1
	} else {
		goto L882
	}
L880:
	;
	v6825 = v6821
	goto L881
L881:
	;
	v6827 = int32(0)
	v6829 = F_makeTargetEntry(m, v6825, base.I32_extend16_s(v6772), v6827, v6827)
	mBase = m.M
	v6830 = m.ExcPending
	if v6830 != 0 {
		goto L1
	} else {
		goto L883
	}
L882:
	;
	v6825 = v6823
	goto L881
L883:
	;
	if v6768 != 0 {
		goto L884
	} else {
		goto L885
	}
L884:
	;
	v6836 = *(*int32)(unsafe.Add(mBase, uint32(v6768+v6772<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6829)+16)) = v6836
	goto L886
L885:
	;
	goto L886
L886:
	;
	v6840 = F_lappend(m, v6771, v6829)
	mBase = m.M
	v6841 = m.ExcPending
	if v6841 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	v6843 = v6774 + int32(1)
	v6844 = *(*int32)(unsafe.Add(mBase, uint32(v6761)+4))
	if v6843 < v6844 {
		v6771 = v6840
		v6772 = v6772 + int32(1)
		v6774 = v6843
		goto L877
	} else {
		goto L888
	}
L888:
	;
	goto L878
L889:
	;
	v6897 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v6897 != 0 {
		goto L890
	} else {
		goto L891
	}
L890:
	;
	v6898 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+4))
	v6899 = v6898
	goto L892
L891:
	;
	v6899 = v4
	goto L892
L892:
	;
	v6900 = *(*int32)(unsafe.Add(mBase, uint32(v6757)+44))
	v6901 = F_extract_grouping_cols(m, v6897, v6900)
	mBase = m.M
	v6902 = m.ExcPending
	if v6902 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	v6903 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6904 = F_extract_grouping_ops(m, v6903)
	mBase = m.M
	v6905 = m.ExcPending
	if v6905 != 0 {
		goto L1
	} else {
		goto L894
	}
L894:
	;
	v6906 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6907 = *(*int32)(unsafe.Add(mBase, uint32(v6757)+44))
	v6908 = F_extract_grouping_collations(m, v6906, v6907)
	mBase = m.M
	v6909 = m.ExcPending
	if v6909 != 0 {
		goto L1
	} else {
		goto L895
	}
L895:
	;
	v6911 = F_palloc0(m, int32(88))
	mBase = m.M
	v6912 = m.ExcPending
	if v6912 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+84)) = v6908
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+80)) = v6904
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+76)) = v6901
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+72)) = v6899
	*(*int32)(unsafe.Add(mBase, uint32(v6911))) = int32(368)
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+48)) = v6895
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+52)) = v6757
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+44)) = v6848
	v6924 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+4)) = v6924
	v6926 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6911)+8)) = v6926
	v6928 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6911)+16)) = v6928
	v6930 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6911)+24)) = v6930
	v6932 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6933 = *(*int32)(unsafe.Add(mBase, uint32(v6932)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6911)+32)) = v6933
	v6935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6911)+36)) = uint8(v6935)
	v6937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6911)+37)) = uint8(v6937)
	v12392 = v6911
	v12397 = v51
	goto L3
L897:
	;
	v6944 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6945 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6947 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6948 = *(*int32)(unsafe.Add(mBase, uint32(v6947)+8))
	v6949 = *(*int32)(unsafe.Add(mBase, uint32(v6948)+4))
	if base.Ui32(int32(5)) < base.Ui32(v6949) {
		v6961 = int32(0)
		goto L898
	} else {
		goto L899
	}
L898:
	;
	v6962 = int32(0)
	v6974 = F_prepare_sort_from_pathkeys(m, v6942, v6945, v6961, v6962, v6962, v51+int32(88), v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160))
	mBase = m.M
	v6975 = m.ExcPending
	if v6975 != 0 {
		goto L1
	} else {
		goto L901
	}
L899:
	;
	v6952 = int32(0)
	if int32(1)<<(uint(v6949)%32)&int32(44) == v6952 {
		v6961 = v6952
		goto L898
	} else {
		goto L900
	}
L900:
	;
	v6959 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6960 = *(*int32)(unsafe.Add(mBase, uint32(v6959)+8))
	v6961 = v6960
	goto L898
L901:
	;
	v6976 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v6977 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v6978 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v6979 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v6980 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v6982 = F_palloc0(m, int32(104))
	mBase = m.M
	v6983 = m.ExcPending
	if v6983 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6982))) = int32(367)
	v6986 = *(*int32)(unsafe.Add(mBase, uint32(v6974)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+96)) = v6944
	v6988 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+56)) = v6988
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+52)) = v6974
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+48)) = v6988
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+44)) = v6986
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+88)) = v6980
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+84)) = v6979
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+80)) = v6978
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+76)) = v6977
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+72)) = v6976
	v6999 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+4)) = v6999
	v7001 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6982)+8)) = v7001
	v7003 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6982)+16)) = v7003
	v7005 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6982)+24)) = v7005
	v7007 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7008 = *(*int32)(unsafe.Add(mBase, uint32(v7007)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6982)+32)) = v7008
	v7010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6982)+36)) = uint8(v7010)
	v7012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6982)+37)) = uint8(v7012)
	v12392 = v6982
	v12397 = v51
	goto L3
L903:
	;
	v7019 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v7021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v7022 = *(*int32)(unsafe.Add(mBase, uint32(v7021)+8))
	v7023 = *(*int32)(unsafe.Add(mBase, uint32(v7022)+4))
	if base.Ui32(int32(5)) < base.Ui32(v7023) {
		v7035 = int32(0)
		goto L904
	} else {
		goto L905
	}
L904:
	;
	v7036 = int32(0)
	v7048 = F_prepare_sort_from_pathkeys(m, v7017, v7019, v7035, v7036, v7036, v51+int32(88), v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160))
	mBase = m.M
	v7049 = m.ExcPending
	if v7049 != 0 {
		goto L1
	} else {
		goto L907
	}
L905:
	;
	v7026 = int32(0)
	if int32(1)<<(uint(v7023)%32)&int32(44) == v7026 {
		v7035 = v7026
		goto L904
	} else {
		goto L906
	}
L906:
	;
	v7033 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7034 = *(*int32)(unsafe.Add(mBase, uint32(v7033)+8))
	v7035 = v7034
	goto L904
L907:
	;
	v7050 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v7051 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v7052 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v7053 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v7054 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v7056 = F_palloc0(m, int32(96))
	mBase = m.M
	v7057 = m.ExcPending
	if v7057 != 0 {
		goto L1
	} else {
		goto L908
	}
L908:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7056))) = int32(366)
	v7060 = *(*int32)(unsafe.Add(mBase, uint32(v7048)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+44)) = v7060
	v7062 = *(*int32)(unsafe.Add(mBase, uint32(v7048)+4))
	v7064 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+88)) = v7054
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+84)) = v7053
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+80)) = v7052
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+76)) = v7051
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+72)) = v7050
	v7070 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+56)) = v7070
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+52)) = v7048
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+48)) = v7070
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+4)) = v7062 + (v7064 ^ int32(1))
	v7079 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+4)) = v7079
	v7081 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v7056)+8)) = v7081
	v7083 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v7056)+16)) = v7083
	v7085 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v7056)+24)) = v7085
	v7087 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7088 = *(*int32)(unsafe.Add(mBase, uint32(v7087)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7056)+32)) = v7088
	v7090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7056)+36)) = uint8(v7090)
	v7092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7056)+37)) = uint8(v7092)
	v12392 = v7056
	v12397 = v51
	goto L3
L909:
	;
	v7099 = int32(0)
	v7100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7101 = *(*int32)(unsafe.Add(mBase, uint32(v7100)+4))
	if v7101 == v7099 {
		v7187 = v7099
		goto L910
	} else {
		goto L911
	}
L910:
	;
	v7233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v7234 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v7235 = m.ExcPending
	if v7235 != 0 {
		goto L1
	} else {
		goto L925
	}
L911:
	;
	v7104 = *(*int32)(unsafe.Add(mBase, uint32(v7101)+4))
	if v7104 <= int32(0) {
		v7187 = v7099
		goto L910
	} else {
		goto L912
	}
L912:
	;
	v7107 = *(*int32)(unsafe.Add(mBase, uint32(v7100)+8))
	v7110 = v7099
	v7111 = v7094
	v7113 = v4
	goto L913
L913:
	;
	v7156 = *(*int32)(unsafe.Add(mBase, uint32(v7101)+12))
	v7160 = *(*int32)(unsafe.Add(mBase, uint32(v7156+v7113<<(uint(int32(2))%32))))
	v7161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7161 != 0 {
		goto L915
	} else {
		goto L916
	}
L914:
	;
	v7187 = v7179
	goto L910
L915:
	;
	v7162 = F_replace_nestloop_params_mutator(m, v7160, l0)
	mBase = m.M
	v7163 = m.ExcPending
	if v7163 != 0 {
		goto L1
	} else {
		goto L918
	}
L916:
	;
	v7164 = v7160
	goto L917
L917:
	;
	v7166 = int32(0)
	v7168 = F_makeTargetEntry(m, v7164, base.I32_extend16_s(v7111), v7166, v7166)
	mBase = m.M
	v7169 = m.ExcPending
	if v7169 != 0 {
		goto L1
	} else {
		goto L919
	}
L918:
	;
	v7164 = v7162
	goto L917
L919:
	;
	if v7107 != 0 {
		goto L920
	} else {
		goto L921
	}
L920:
	;
	v7175 = *(*int32)(unsafe.Add(mBase, uint32(v7107+v7111<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v7168)+16)) = v7175
	goto L922
L921:
	;
	goto L922
L922:
	;
	v7179 = F_lappend(m, v7110, v7168)
	mBase = m.M
	v7180 = m.ExcPending
	if v7180 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v7182 = v7113 + int32(1)
	v7183 = *(*int32)(unsafe.Add(mBase, uint32(v7101)+4))
	if v7182 < v7183 {
		v7110 = v7179
		v7111 = v7111 + int32(1)
		v7113 = v7182
		goto L913
	} else {
		goto L924
	}
L924:
	;
	goto L914
L925:
	;
	v7236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v7238 = F_palloc0(m, int32(88))
	mBase = m.M
	v7239 = m.ExcPending
	if v7239 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	v7240 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+84)) = v7240
	*(*uint8)(unsafe.Add(mBase, uint32(v7238)+81)) = uint8(v7240)
	*(*uint8)(unsafe.Add(mBase, uint32(v7238)+80)) = uint8(v7236)
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+76)) = v7234
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+72)) = v7233
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+56)) = v7240
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+52)) = v7097
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+48)) = v7240
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+44)) = v7187
	*(*int32)(unsafe.Add(mBase, uint32(v7238))) = int32(372)
	v7255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+4)) = v7255
	v7257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v7238)+8)) = v7257
	v7259 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v7238)+16)) = v7259
	v7261 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v7238)+24)) = v7261
	v7263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7264 = *(*int32)(unsafe.Add(mBase, uint32(v7263)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7238)+32)) = v7264
	v7266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7238)+36)) = uint8(v7266)
	v7268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7238)+37)) = uint8(v7268)
	v7270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7271 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7270)+95)) = uint8(v7271)
	v12392 = v7238
	v12397 = v51
	goto L3
L927:
	;
	v7278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v7279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v7281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7282 = *(*int32)(unsafe.Add(mBase, uint32(v7281)+4))
	if base.Ui32(int32(5)) < base.Ui32(v7282) {
		v7293 = int32(0)
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v7295 = F_palloc0(m, int32(88))
	mBase = m.M
	v7296 = m.ExcPending
	if v7296 != 0 {
		goto L1
	} else {
		goto L931
	}
L929:
	;
	v7285 = int32(0)
	if int32(1)<<(uint(v7282)%32)&int32(44) == v7285 {
		v7293 = v7285
		goto L928
	} else {
		goto L930
	}
L930:
	;
	v7292 = *(*int32)(unsafe.Add(mBase, uint32(v7281)+8))
	v7293 = v7292
	goto L928
L931:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7295))) = int32(371)
	v7299 = *(*int32)(unsafe.Add(mBase, uint32(v7276)+44))
	v7300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+56)) = v7300
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+52)) = v7276
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+48)) = v7300
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+44)) = v7299
	v7307 = F_palloc_mul(m, int32(2), v7278)
	mBase = m.M
	v7308 = m.ExcPending
	if v7308 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	v7310 = F_palloc_mul(m, int32(4), v7278)
	mBase = m.M
	v7311 = m.ExcPending
	if v7311 != 0 {
		goto L1
	} else {
		goto L933
	}
L933:
	;
	v7313 = F_palloc_mul(m, int32(4), v7278)
	mBase = m.M
	v7314 = m.ExcPending
	if v7314 != 0 {
		goto L1
	} else {
		goto L934
	}
L934:
	;
	if v7279 == int32(0) {
		goto L937
	} else {
		goto L938
	}
L935:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7678 = m.ExcPending
	if v7678 != 0 {
		goto L1
	} else {
		goto L970
	}
L936:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7610 = m.ExcPending
	if v7610 != 0 {
		goto L1
	} else {
		goto L967
	}
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+84)) = v7313
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+80)) = v7310
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+76)) = v7307
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+72)) = v7278
	v7592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+4)) = v7592
	v7594 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v7295)+8)) = v7594
	v7596 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v7295)+16)) = v7596
	v7598 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v7295)+24)) = v7598
	v7600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7601 = *(*int32)(unsafe.Add(mBase, uint32(v7600)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7295)+32)) = v7601
	v7603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7295)+36)) = uint8(v7603)
	v7605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7295)+37)) = uint8(v7605)
	v12392 = v7295
	v12397 = v51
	goto L3
L938:
	;
	v7317 = *(*int32)(unsafe.Add(mBase, uint32(v7279)+4))
	if v7317 <= int32(0) {
		goto L937
	} else {
		goto L939
	}
L939:
	;
	v7320 = int32(0)
	if v7320 < v7278 {
		goto L940
	} else {
		goto L941
	}
L940:
	;
	v7323 = v7278
	goto L942
L941:
	;
	v7323 = v7320
	goto L942
L942:
	;
	v7334 = v4
	goto L943
L943:
	;
	if v7334 == v7323 {
		goto L937
	} else {
		goto L945
	}
L944:
	;
	goto L937
L945:
	;
	v7374 = v7334 << (uint(int32(2)) % 32)
	v7375 = *(*int32)(unsafe.Add(mBase, uint32(v7279)+12))
	v7377 = *(*int32)(unsafe.Add(mBase, uint32(v7374+v7375)))
	v7378 = *(*int32)(unsafe.Add(mBase, uint32(v7377)+4))
	v7379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7378)+41)))
	if v7379 == int32(1) {
		goto L948
	} else {
		goto L949
	}
L946:
	;
	v7519 = *(*int32)(unsafe.Add(mBase, uint32(v7377)+8))
	v7520 = *(*int32)(unsafe.Add(mBase, uint32(v7477)+16))
	v7522 = F_get_opfamily_member_for_cmptype(m, v7519, v7520, v7520, int32(3))
	mBase = m.M
	v7523 = m.ExcPending
	if v7523 != 0 {
		goto L1
	} else {
		goto L964
	}
L947:
	;
	v7463 = *(*int32)(unsafe.Add(mBase, uint32(v7295)+44))
	v7464 = F_get_sortgroupref_tle(m, v7382, v7463)
	mBase = m.M
	v7465 = m.ExcPending
	if v7465 != 0 {
		goto L1
	} else {
		goto L962
	}
L948:
	;
	v7382 = *(*int32)(unsafe.Add(mBase, uint32(v7378)+44))
	if v7382 != 0 {
		goto L947
	} else {
		goto L951
	}
L949:
	;
	goto L950
L950:
	;
	v7396 = *(*int32)(unsafe.Add(mBase, uint32(v7295)+44))
	if v7396 == int32(0) {
		goto L935
	} else {
		goto L955
	}
L951:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7386 = m.ExcPending
	if v7386 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_35), int32(0))
	mBase = m.M
	v7390 = m.ExcPending
	if v7390 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_36), int32(_a_F_create_plan_recurse_37))
	mBase = m.M
	v7395 = m.ExcPending
	if v7395 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L955:
	;
	v7399 = int32(0)
	v7400 = *(*int32)(unsafe.Add(mBase, uint32(v7396)+4))
	if v7400 <= v7399 {
		goto L935
	} else {
		goto L956
	}
L956:
	;
	v7403 = v7399
	goto L957
L957:
	;
	v7451 = *(*int32)(unsafe.Add(mBase, uint32(v7396)+12))
	v7455 = *(*int32)(unsafe.Add(mBase, uint32(v7451+v7403<<(uint(int32(2))%32))))
	v7456 = *(*int32)(unsafe.Add(mBase, uint32(v7455)+4))
	v7457 = F_find_ec_member_matching_expr(m, v7378, v7456, v7293)
	mBase = m.M
	v7458 = m.ExcPending
	if v7458 != 0 {
		goto L1
	} else {
		goto L959
	}
L958:
	;
	goto L935
L959:
	;
	if v7457 != 0 {
		v7473 = v7455
		v7477 = v7457
		goto L946
	} else {
		goto L960
	}
L960:
	;
	v7460 = v7403 + int32(1)
	v7461 = *(*int32)(unsafe.Add(mBase, uint32(v7396)+4))
	if v7460 < v7461 {
		v7403 = v7460
		goto L957
	} else {
		goto L961
	}
L961:
	;
	goto L958
L962:
	;
	if v7464 == int32(0) {
		goto L935
	} else {
		goto L963
	}
L963:
	;
	v7468 = *(*int32)(unsafe.Add(mBase, uint32(v7378)+16))
	v7469 = *(*int32)(unsafe.Add(mBase, uint32(v7468)+12))
	v7470 = *(*int32)(unsafe.Add(mBase, uint32(v7469)))
	v7473 = v7464
	v7477 = v7470
	goto L946
L964:
	;
	if v7522 == int32(0) {
		goto L936
	} else {
		goto L965
	}
L965:
	;
	v7526 = int32(1)
	v7529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7473)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7307+v7334<<(uint(v7526)%32)))) = uint16(v7529)
	*(*int32)(unsafe.Add(mBase, uint32(v7310+v7374))) = v7522
	v7534 = *(*int32)(unsafe.Add(mBase, uint32(v7378)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7313+v7374))) = v7534
	v7537 = v7334 + v7526
	v7538 = *(*int32)(unsafe.Add(mBase, uint32(v7279)+4))
	if v7537 < v7538 {
		v7334 = v7537
		goto L943
	} else {
		goto L966
	}
L966:
	;
	goto L944
L967:
	;
	v7611 = *(*int32)(unsafe.Add(mBase, uint32(v7377)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+44)) = v7611
	*(*int32)(unsafe.Add(mBase, uint32(v51)+40)) = v7520
	*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v7520
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_38), v51+int32(32))
	mBase = m.M
	v7621 = m.ExcPending
	if v7621 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_39), int32(_a_F_create_plan_recurse_37))
	mBase = m.M
	v7626 = m.ExcPending
	if v7626 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L970:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_40), int32(0))
	mBase = m.M
	v7682 = m.ExcPending
	if v7682 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_41), int32(_a_F_create_plan_recurse_37))
	mBase = m.M
	v7687 = m.ExcPending
	if v7687 != 0 {
		goto L1
	} else {
		goto L972
	}
L972:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L973:
	;
	v7693 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v7694 = F_replace_nestloop_params_mutator(m, v7693, l0)
	mBase = m.M
	v7695 = m.ExcPending
	if v7695 != 0 {
		goto L1
	} else {
		goto L974
	}
L974:
	;
	if v7694 != 0 {
		goto L975
	} else {
		goto L976
	}
L975:
	;
	v7696 = *(*int32)(unsafe.Add(mBase, uint32(v7694)+4))
	v7699 = v7696 << (uint(int32(2)) % 32)
	goto L977
L976:
	;
	v7699 = v4
	goto L977
L977:
	;
	v7700 = F_palloc(m, v7699)
	mBase = m.M
	v7701 = m.ExcPending
	if v7701 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	v7702 = F_palloc(m, v7699)
	mBase = m.M
	v7703 = m.ExcPending
	if v7703 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	v7704 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v7708 = v4
	goto L980
L980:
	;
	v7753 = int32(0)
	if v7694 == v7753 {
		v7762 = v7753
		goto L982
	} else {
		goto L983
	}
L981:
	;
	v7787 = m.G0
	v7789 = v7787 - int32(16)
	m.G0 = v7789
	v7791 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7789)+12)) = v7791
	if v7694 == v7791 {
		v7809 = v7791
		goto L990
	} else {
		goto L991
	}
L982:
	;
	if v7704 == int32(0) {
		goto L985
	} else {
		goto L986
	}
L983:
	;
	v7756 = *(*int32)(unsafe.Add(mBase, uint32(v7694)+4))
	if v7756 <= v7708 {
		v7762 = v7753
		goto L982
	} else {
		goto L984
	}
L984:
	;
	v7758 = *(*int32)(unsafe.Add(mBase, uint32(v7694)+12))
	v7762 = v7758 + v7708<<(uint(int32(2))%32)
	goto L982
L985:
	;
	goto L981
L986:
	;
	v7767 = *(*int32)(unsafe.Add(mBase, uint32(v7704)+4))
	if base.B2i32(v7762 == int32(0))|base.B2i32(v7767 <= v7708) != 0 {
		goto L985
	} else {
		goto L987
	}
L987:
	;
	v7770 = *(*int32)(unsafe.Add(mBase, uint32(v7704)+12))
	if v7770 == int32(0) {
		goto L985
	} else {
		goto L988
	}
L988:
	;
	v7773 = *(*int32)(unsafe.Add(mBase, uint32(v7762)))
	v7775 = v7708 << (uint(int32(2)) % 32)
	v7778 = *(*int32)(unsafe.Add(mBase, uint32(v7775+v7770)))
	*(*int32)(unsafe.Add(mBase, uint32(v7700+v7775))) = v7778
	v7781 = F_exprCollation(m, v7773)
	mBase = m.M
	v7782 = m.ExcPending
	if v7782 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7775+v7702))) = v7781
	v7708 = v7708 + int32(1)
	goto L980
L990:
	;
	m.G0 = v7789 + int32(16)
	v7813 = *(*float64)(unsafe.Add(mBase, uint32(l1)+112))
	v7814 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	v7815 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	v7816 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v7817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v7818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v7820 = F_palloc0(m, int32(128))
	mBase = m.M
	v7821 = m.ExcPending
	if v7821 != 0 {
		goto L1
	} else {
		goto L997
	}
L991:
	;
	v7796 = *(*int32)(unsafe.Add(mBase, uint32(v7694)))
	if v7796 == int32(8) {
		goto L992
	} else {
		goto L993
	}
L992:
	;
	v7800 = *(*int32)(unsafe.Add(mBase, uint32(v7694)+8))
	v7801 = F_bms_add_member(m, int32(0), v7800)
	mBase = m.M
	v7802 = m.ExcPending
	if v7802 != 0 {
		goto L1
	} else {
		goto L995
	}
L993:
	;
	goto L994
L994:
	;
	v7806 = F_expression_tree_walker_impl(m, v7694, int32(924), v7789+int32(12))
	mBase = m.M
	v7807 = m.ExcPending
	if v7807 != 0 {
		goto L1
	} else {
		goto L996
	}
L995:
	;
	v7809 = v7801
	goto L990
L996:
	;
	v7808 = *(*int32)(unsafe.Add(mBase, uint32(v7789)+12))
	v7809 = v7808
	goto L990
L997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7820))) = int32(365)
	v7824 = *(*int32)(unsafe.Add(mBase, uint32(v7691)+44))
	v7825 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+56)) = v7825
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+52)) = v7691
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+48)) = v7825
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+44)) = v7824
	if v7694 != 0 {
		goto L998
	} else {
		goto L999
	}
L998:
	;
	v7832 = *(*int32)(unsafe.Add(mBase, uint32(v7694)+4))
	v7833 = v7832
	goto L1000
L999:
	;
	v7833 = v7825
	goto L1000
L1000:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+120)) = v7813
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+112)) = v7814
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+104)) = v7815
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+96)) = v7809
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+92)) = v7816
	*(*uint8)(unsafe.Add(mBase, uint32(v7820)+89)) = uint8(v7817)
	*(*uint8)(unsafe.Add(mBase, uint32(v7820)+88)) = uint8(v7818)
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+84)) = v7694
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+80)) = v7702
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+76)) = v7700
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+72)) = v7833
	v7845 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+4)) = v7845
	v7847 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+8)) = v7847
	v7849 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+16)) = v7849
	v7851 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v7820)+24)) = v7851
	v7853 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7854 = *(*int32)(unsafe.Add(mBase, uint32(v7853)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7820)+32)) = v7854
	v7856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7820)+36)) = uint8(v7856)
	v7858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7820)+37)) = uint8(v7858)
	v12392 = v7820
	v12397 = v51
	goto L3
L1001:
	;
	v7866 = F_palloc0(m, int32(72))
	mBase = m.M
	v7867 = m.ExcPending
	if v7867 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7866))) = int32(364)
	v7870 = *(*int32)(unsafe.Add(mBase, uint32(v7863)+44))
	v7871 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+56)) = v7871
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+52)) = v7863
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+48)) = v7871
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+44)) = v7870
	v7877 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+4)) = v7877
	v7879 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v7866)+8)) = v7879
	v7881 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v7866)+16)) = v7881
	v7883 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v7866)+24)) = v7883
	v7885 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7886 = *(*int32)(unsafe.Add(mBase, uint32(v7885)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7866)+32)) = v7886
	v7888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7866)+36)) = uint8(v7888)
	v7890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7866)+37)) = uint8(v7890)
	v12392 = v7866
	v12397 = v51
	goto L3
L1003:
	;
	v7896 = int32(0)
	v7897 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7898 = *(*int32)(unsafe.Add(mBase, uint32(v7897)+4))
	if v7898 == v7896 {
		v7985 = v7896
		goto L1004
	} else {
		goto L1005
	}
L1004:
	;
	v8032 = F_palloc0(m, int32(72))
	mBase = m.M
	v8033 = m.ExcPending
	if v8033 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1005:
	;
	v7902 = *(*int32)(unsafe.Add(mBase, uint32(v7898)+4))
	if v7902 <= int32(0) {
		v7985 = v7896
		goto L1004
	} else {
		goto L1006
	}
L1006:
	;
	v7905 = *(*int32)(unsafe.Add(mBase, uint32(v7897)+8))
	v7908 = v7896
	v7909 = int32(1)
	v7911 = v4
	goto L1007
L1007:
	;
	v7954 = *(*int32)(unsafe.Add(mBase, uint32(v7898)+12))
	v7958 = *(*int32)(unsafe.Add(mBase, uint32(v7954+v7911<<(uint(int32(2))%32))))
	v7959 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7959 != 0 {
		goto L1009
	} else {
		goto L1010
	}
L1008:
	;
	v7985 = v7977
	goto L1004
L1009:
	;
	v7960 = F_replace_nestloop_params_mutator(m, v7958, l0)
	mBase = m.M
	v7961 = m.ExcPending
	if v7961 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1010:
	;
	v7962 = v7958
	goto L1011
L1011:
	;
	v7964 = int32(0)
	v7966 = F_makeTargetEntry(m, v7962, base.I32_extend16_s(v7909), v7964, v7964)
	mBase = m.M
	v7967 = m.ExcPending
	if v7967 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1012:
	;
	v7962 = v7960
	goto L1011
L1013:
	;
	if v7905 != 0 {
		goto L1014
	} else {
		goto L1015
	}
L1014:
	;
	v7973 = *(*int32)(unsafe.Add(mBase, uint32(v7905+v7909<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v7966)+16)) = v7973
	goto L1016
L1015:
	;
	goto L1016
L1016:
	;
	v7977 = F_lappend(m, v7908, v7966)
	mBase = m.M
	v7978 = m.ExcPending
	if v7978 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	v7980 = v7911 + int32(1)
	v7981 = *(*int32)(unsafe.Add(mBase, uint32(v7898)+4))
	if v7980 < v7981 {
		v7908 = v7977
		v7909 = v7909 + int32(1)
		v7911 = v7980
		goto L1007
	} else {
		goto L1018
	}
L1018:
	;
	goto L1008
L1019:
	;
	v8034 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+56)) = v8034
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+52)) = v7894
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+48)) = v8034
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+44)) = v7985
	*(*int32)(unsafe.Add(mBase, uint32(v8032))) = int32(336)
	v8042 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+4)) = v8042
	v8044 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8032)+8)) = v8044
	v8046 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8032)+16)) = v8046
	v8048 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8032)+24)) = v8048
	v8050 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8051 = *(*int32)(unsafe.Add(mBase, uint32(v8050)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8032)+32)) = v8051
	v8053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8032)+36)) = uint8(v8053)
	v8055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8032)+37)) = uint8(v8055)
	v12392 = v8032
	v12397 = v51
	goto L3
L1020:
	;
	v9060 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v9061 = m.ExcPending
	if v9061 != 0 {
		goto L1
	} else {
		goto L1143
	}
L1021:
	;
	v8882 = int32(0)
	v8884 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8885 = *(*int32)(unsafe.Add(mBase, uint32(v8884)+4))
	if v8885 == v8882 {
		v8970 = v8882
		goto L1123
	} else {
		goto L1124
	}
L1022:
	;
	v8463 = int32(0)
	v8464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v8464 == v8463 {
		goto L1079
	} else {
		goto L1080
	}
L1023:
	;
	v8060 = F_use_physical_tlist(m, l0, l1, l2)
	mBase = m.M
	v8061 = m.ExcPending
	if v8061 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	v8062 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v8060 != 0 {
		goto L1028
	} else {
		goto L1029
	}
L1025:
	;
	v8461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8413)+37)) = uint8(v8461)
	v12392 = v8413
	v12397 = v51
	goto L3
L1026:
	;
	v8386 = F_palloc0(m, int32(88))
	mBase = m.M
	v8387 = m.ExcPending
	if v8387 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1027:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8327)+44)) = v8330
	v8376 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8327)+8)) = v8376
	v8378 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8327)+16)) = v8378
	v8380 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8327)+24)) = v8380
	v8382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8383 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8327)+32)) = v8383
	v8413 = v8327
	goto L1025
L1028:
	;
	v8064 = F_create_plan_recurse(m, l0, v8062, int32(0))
	mBase = m.M
	v8065 = m.ExcPending
	if v8065 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1029:
	;
	goto L1030
L1030:
	;
	v8074 = int32(0)
	v8075 = *(*int32)(unsafe.Add(mBase, uint32(v8062)+4))
	switch v8075 - int32(336) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v8090 = v8074
		goto L1035
	case 2:
		goto L1037
	default:
		goto L1036
	case 23:
		goto L1038
	}
L1031:
	;
	v8066 = *(*int32)(unsafe.Add(mBase, uint32(v8064)+44))
	if l2&int32(4) == int32(0) {
		v8327 = v8064
		v8330 = v8066
		goto L1027
	} else {
		goto L1032
	}
L1032:
	;
	v8071 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v8066, v8071)
	mBase = m.M
	v8073 = m.ExcPending
	if v8073 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	v8327 = v8064
	v8330 = v8066
	goto L1027
L1034:
	;
	if v8092 != 0 {
		goto L1040
	} else {
		goto L1041
	}
L1035:
	;
	v8092 = v8090
	goto L1034
L1036:
	;
	v8090 = int32(1)
	goto L1035
L1037:
	;
	v8083 = *(*int32)(unsafe.Add(mBase, uint32(v8062)))
	if v8083 != int32(293) {
		v8090 = v8074
		goto L1035
	} else {
		goto L1039
	}
L1038:
	;
	v8078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8062)+72)))
	v8092 = int32(base.Ui32(v8078&int32(4)) >> (uint(int32(2)) % 32))
	goto L1034
L1039:
	;
	v8086 = *(*int32)(unsafe.Add(mBase, uint32(v8062)+72))
	v8092 = base.B2i32(v8086 == int32(0))
	goto L1034
L1040:
	;
	v8094 = F_create_plan_recurse(m, l0, v8062, int32(8))
	mBase = m.M
	v8095 = m.ExcPending
	if v8095 != 0 {
		goto L1
	} else {
		goto L1043
	}
L1041:
	;
	goto L1042
L1042:
	;
	v8184 = int32(0)
	v8186 = F_create_plan_recurse(m, l0, v8062, v8184)
	mBase = m.M
	v8187 = m.ExcPending
	if v8187 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1043:
	;
	v8096 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8097 = *(*int32)(unsafe.Add(mBase, uint32(v8096)+4))
	if v8097 == int32(0) {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	v8327 = v8094
	v8330 = int32(0)
	goto L1027
L1045:
	;
	goto L1046
L1046:
	;
	v8102 = int32(0)
	v8103 = *(*int32)(unsafe.Add(mBase, uint32(v8097)+4))
	if v8103 <= v8102 {
		v8327 = v8094
		v8330 = v8102
		goto L1027
	} else {
		goto L1047
	}
L1047:
	;
	v8106 = *(*int32)(unsafe.Add(mBase, uint32(v8096)+8))
	v8109 = int32(1)
	v8110 = v8102
	v8118 = v4
	goto L1048
L1048:
	;
	v8155 = *(*int32)(unsafe.Add(mBase, uint32(v8097)+12))
	v8159 = *(*int32)(unsafe.Add(mBase, uint32(v8155+v8118<<(uint(int32(2))%32))))
	v8160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8160 != 0 {
		goto L1050
	} else {
		goto L1051
	}
L1049:
	;
	v8327 = v8094
	v8330 = v8178
	goto L1027
L1050:
	;
	v8161 = F_replace_nestloop_params_mutator(m, v8159, l0)
	mBase = m.M
	v8162 = m.ExcPending
	if v8162 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1051:
	;
	v8163 = v8159
	goto L1052
L1052:
	;
	v8165 = int32(0)
	v8167 = F_makeTargetEntry(m, v8163, base.I32_extend16_s(v8109), v8165, v8165)
	mBase = m.M
	v8168 = m.ExcPending
	if v8168 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1053:
	;
	v8163 = v8161
	goto L1052
L1054:
	;
	if v8106 != 0 {
		goto L1055
	} else {
		goto L1056
	}
L1055:
	;
	v8174 = *(*int32)(unsafe.Add(mBase, uint32(v8106+v8109<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8167)+16)) = v8174
	goto L1057
L1056:
	;
	goto L1057
L1057:
	;
	v8178 = F_lappend(m, v8110, v8167)
	mBase = m.M
	v8179 = m.ExcPending
	if v8179 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1058:
	;
	v8181 = v8118 + int32(1)
	v8182 = *(*int32)(unsafe.Add(mBase, uint32(v8097)+4))
	if v8181 < v8182 {
		v8109 = v8109 + int32(1)
		v8110 = v8178
		v8118 = v8181
		goto L1048
	} else {
		goto L1059
	}
L1059:
	;
	goto L1049
L1060:
	;
	v8188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8189 = *(*int32)(unsafe.Add(mBase, uint32(v8188)+4))
	if v8189 == int32(0) {
		v8277 = v8184
		goto L1061
	} else {
		goto L1062
	}
L1061:
	;
	v8322 = *(*int32)(unsafe.Add(mBase, uint32(v8186)+44))
	v8323 = F_tlist_same_exprs(m, v8277, v8322)
	mBase = m.M
	v8324 = m.ExcPending
	if v8324 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1062:
	;
	v8193 = *(*int32)(unsafe.Add(mBase, uint32(v8189)+4))
	if v8193 <= int32(0) {
		v8277 = v8184
		goto L1061
	} else {
		goto L1063
	}
L1063:
	;
	v8196 = *(*int32)(unsafe.Add(mBase, uint32(v8188)+8))
	v8199 = int32(1)
	v8200 = v8184
	v8208 = v4
	goto L1064
L1064:
	;
	v8245 = *(*int32)(unsafe.Add(mBase, uint32(v8189)+12))
	v8249 = *(*int32)(unsafe.Add(mBase, uint32(v8245+v8208<<(uint(int32(2))%32))))
	v8250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8250 != 0 {
		goto L1066
	} else {
		goto L1067
	}
L1065:
	;
	v8277 = v8268
	goto L1061
L1066:
	;
	v8251 = F_replace_nestloop_params_mutator(m, v8249, l0)
	mBase = m.M
	v8252 = m.ExcPending
	if v8252 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1067:
	;
	v8253 = v8249
	goto L1068
L1068:
	;
	v8255 = int32(0)
	v8257 = F_makeTargetEntry(m, v8253, base.I32_extend16_s(v8199), v8255, v8255)
	mBase = m.M
	v8258 = m.ExcPending
	if v8258 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1069:
	;
	v8253 = v8251
	goto L1068
L1070:
	;
	if v8196 != 0 {
		goto L1071
	} else {
		goto L1072
	}
L1071:
	;
	v8264 = *(*int32)(unsafe.Add(mBase, uint32(v8196+v8199<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8257)+16)) = v8264
	goto L1073
L1072:
	;
	goto L1073
L1073:
	;
	v8268 = F_lappend(m, v8200, v8257)
	mBase = m.M
	v8269 = m.ExcPending
	if v8269 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	v8271 = v8208 + int32(1)
	v8272 = *(*int32)(unsafe.Add(mBase, uint32(v8189)+4))
	if v8271 < v8272 {
		v8199 = v8199 + int32(1)
		v8200 = v8268
		v8208 = v8271
		goto L1064
	} else {
		goto L1075
	}
L1075:
	;
	goto L1065
L1076:
	;
	if v8323 == int32(0) {
		goto L1026
	} else {
		goto L1077
	}
L1077:
	;
	v8327 = v8186
	v8330 = v8277
	goto L1027
L1078:
	;
	v8388 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+80)) = v8388
	*(*int64)(unsafe.Add(mBase, uint32(v8386)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+56)) = v8388
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+52)) = v8186
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+48)) = v8388
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+44)) = v8277
	*(*int32)(unsafe.Add(mBase, uint32(v8386))) = int32(335)
	v8400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+4)) = v8400
	v8402 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8386)+8)) = v8402
	v8404 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8386)+16)) = v8404
	v8406 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8386)+24)) = v8406
	v8408 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8409 = *(*int32)(unsafe.Add(mBase, uint32(v8408)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8386)+32)) = v8409
	v8411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8386)+36)) = uint8(v8411)
	v8413 = v8386
	goto L1025
L1079:
	;
	v8702 = int32(0)
	v8703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v8703)+4))
	if v8704 == v8702 {
		v8792 = v8702
		goto L1104
	} else {
		goto L1105
	}
L1080:
	;
	v8467 = *(*int32)(unsafe.Add(mBase, uint32(v8464)+4))
	if v8467 <= int32(0) {
		goto L1079
	} else {
		goto L1081
	}
L1081:
	;
	v8470 = v8463
	goto L1082
L1082:
	;
	v8518 = *(*int32)(unsafe.Add(mBase, uint32(v8464)+12))
	v8522 = *(*int32)(unsafe.Add(mBase, uint32(v8518+v8470<<(uint(int32(2))%32))))
	v8523 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+16))
	v8524 = *(*int32)(unsafe.Add(mBase, uint32(v8523)+4))
	v8525 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+20))
	v8526 = F_create_plan(m, v8523, v8525)
	mBase = m.M
	v8527 = m.ExcPending
	if v8527 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1083:
	;
	goto L1079
L1084:
	;
	v8528 = *(*int64)(unsafe.Add(mBase, uint32(v8524)+128))
	v8529 = *(*int32)(unsafe.Add(mBase, uint32(v8524)+136))
	v8531 = F_palloc0(m, int32(104))
	mBase = m.M
	v8532 = m.ExcPending
	if v8532 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8531))) = int32(377)
	v8535 = *(*int32)(unsafe.Add(mBase, uint32(v8526)+44))
	v8536 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8531)+84)) = v8536
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+80)) = v8529
	*(*int64)(unsafe.Add(mBase, uint32(v8531)+72)) = v8528
	v8540 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+56)) = v8540
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+52)) = v8526
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+48)) = v8540
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+44)) = v8535
	*(*int64)(unsafe.Add(mBase, uint32(v8531)+92)) = v8536
	v8548 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+20))
	v8549 = *(*int32)(unsafe.Add(mBase, uint32(v8548)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+4)) = v8549
	v8551 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+20))
	v8552 = *(*float64)(unsafe.Add(mBase, uint32(v8551)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8531)+8)) = v8552
	v8554 = *(*float64)(unsafe.Add(mBase, uint32(v8522)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v8531)+24)) = int64(4607182418800017408)
	*(*float64)(unsafe.Add(mBase, uint32(v8531)+16)) = v8554
	v8558 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+20))
	v8559 = *(*int32)(unsafe.Add(mBase, uint32(v8558)+12))
	v8560 = *(*int32)(unsafe.Add(mBase, uint32(v8559)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v8531)+36)) = uint8(v8540)
	*(*int32)(unsafe.Add(mBase, uint32(v8531)+32)) = v8560
	v8564 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+20))
	v8565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8564)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8531)+37)) = uint8(v8565)
	v8567 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+32))
	v8568 = m.G0
	v8570 = v8568 - int32(16)
	m.G0 = v8570
	v8572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8573 = *(*int32)(unsafe.Add(mBase, uint32(v8572)+8))
	v8574 = F_lappend(m, v8573, v8531)
	mBase = m.M
	v8575 = m.ExcPending
	if v8575 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	v8576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8576)+8)) = v8574
	v8578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8579 = *(*int32)(unsafe.Add(mBase, uint32(v8578)+12))
	v8581 = F_lappend(m, v8579, int32(0))
	mBase = m.M
	v8582 = m.ExcPending
	if v8582 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	v8583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8583)+12)) = v8581
	v8585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8586 = *(*int32)(unsafe.Add(mBase, uint32(v8585)+16))
	v8587 = F_lappend(m, v8586, v8523)
	mBase = m.M
	v8588 = m.ExcPending
	if v8588 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	v8589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8589)+16)) = v8587
	v8592 = F_palloc0(m, int32(72))
	mBase = m.M
	v8593 = m.ExcPending
	if v8593 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8592))) = int64(17179869207)
	v8596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8597 = *(*int32)(unsafe.Add(mBase, uint32(v8596)+8))
	if v8597 != 0 {
		goto L1090
	} else {
		goto L1091
	}
L1090:
	;
	v8598 = *(*int32)(unsafe.Add(mBase, uint32(v8597)+4))
	v8600 = v8598
	goto L1092
L1091:
	;
	v8600 = int32(0)
	goto L1092
L1092:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+16)) = v8600
	v8602 = *(*int32)(unsafe.Add(mBase, uint32(v8523)+20))
	v8603 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8592)+36)) = uint8(v8603)
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+20)) = v8602
	v8606 = *(*int32)(unsafe.Add(mBase, uint32(v8531)+44))
	if v8606 == int32(0) {
		goto L1094
	} else {
		goto L1095
	}
L1093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+32)) = v8628
	v8630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8531)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8592)+39)) = uint8(v8630)
	v8632 = *(*int32)(unsafe.Add(mBase, uint32(v8567)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8570)+8)) = v8632
	*(*int32)(unsafe.Add(mBase, uint32(v8570)+12)) = v8632
	v8638 = F_list_make1_impl(m, int32(479), v8570+int32(8))
	mBase = m.M
	v8639 = m.ExcPending
	if v8639 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1094:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8592)+24)) = int64(-4294965018)
	v8628 = int32(0)
	goto L1093
L1095:
	;
	v8609 = *(*int32)(unsafe.Add(mBase, uint32(v8606)+12))
	v8610 = *(*int32)(unsafe.Add(mBase, uint32(v8609)))
	v8611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8610)+26)))
	if v8611 != 0 {
		goto L1094
	} else {
		goto L1096
	}
L1096:
	;
	v8612 = *(*int32)(unsafe.Add(mBase, uint32(v8610)+4))
	v8613 = F_exprType(m, v8612)
	mBase = m.M
	v8614 = m.ExcPending
	if v8614 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+24)) = v8613
	v8616 = *(*int32)(unsafe.Add(mBase, uint32(v8610)+4))
	v8617 = F_exprTypmod(m, v8616)
	mBase = m.M
	v8618 = m.ExcPending
	if v8618 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+28)) = v8617
	v8620 = *(*int32)(unsafe.Add(mBase, uint32(v8610)+4))
	v8621 = F_exprCollation(m, v8620)
	mBase = m.M
	v8622 = m.ExcPending
	if v8622 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1099:
	;
	v8628 = v8621
	goto L1093
L1100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8592)+40)) = v8638
	v8641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v8642 = F_lappend(m, v8641, v8592)
	mBase = m.M
	v8643 = m.ExcPending
	if v8643 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v8642
	F_cost_subplan(m, v8592, v8531)
	mBase = m.M
	v8646 = m.ExcPending
	if v8646 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1102:
	;
	m.G0 = v8570 + int32(16)
	v8651 = v8470 + int32(1)
	v8652 = *(*int32)(unsafe.Add(mBase, uint32(v8464)+4))
	if v8651 < v8652 {
		v8470 = v8651
		goto L1082
	} else {
		goto L1103
	}
L1103:
	;
	goto L1083
L1104:
	;
	v8838 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v8839 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8841 = F_palloc0(m, int32(88))
	mBase = m.M
	v8842 = m.ExcPending
	if v8842 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1105:
	;
	v8708 = *(*int32)(unsafe.Add(mBase, uint32(v8704)+4))
	if v8708 <= int32(0) {
		v8792 = v8702
		goto L1104
	} else {
		goto L1106
	}
L1106:
	;
	v8711 = *(*int32)(unsafe.Add(mBase, uint32(v8703)+8))
	v8713 = int32(0)
	v8715 = v8702
	v8717 = int32(1)
	goto L1107
L1107:
	;
	v8761 = *(*int32)(unsafe.Add(mBase, uint32(v8704)+12))
	v8765 = *(*int32)(unsafe.Add(mBase, uint32(v8761+v8713<<(uint(int32(2))%32))))
	v8766 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8766 != 0 {
		goto L1109
	} else {
		goto L1110
	}
L1108:
	;
	v8792 = v8784
	goto L1104
L1109:
	;
	v8767 = F_replace_nestloop_params_mutator(m, v8765, l0)
	mBase = m.M
	v8768 = m.ExcPending
	if v8768 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1110:
	;
	v8769 = v8765
	goto L1111
L1111:
	;
	v8771 = int32(0)
	v8773 = F_makeTargetEntry(m, v8769, base.I32_extend16_s(v8717), v8771, v8771)
	mBase = m.M
	v8774 = m.ExcPending
	if v8774 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1112:
	;
	v8769 = v8767
	goto L1111
L1113:
	;
	if v8711 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	v8780 = *(*int32)(unsafe.Add(mBase, uint32(v8711+v8717<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8773)+16)) = v8780
	goto L1116
L1115:
	;
	goto L1116
L1116:
	;
	v8784 = F_lappend(m, v8715, v8773)
	mBase = m.M
	v8785 = m.ExcPending
	if v8785 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	v8787 = v8713 + int32(1)
	v8788 = *(*int32)(unsafe.Add(mBase, uint32(v8704)+4))
	if v8787 < v8788 {
		v8713 = v8787
		v8715 = v8784
		v8717 = v8717 + int32(1)
		goto L1107
	} else {
		goto L1118
	}
L1118:
	;
	goto L1108
L1119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8841)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+44)) = v8792
	*(*int32)(unsafe.Add(mBase, uint32(v8841))) = int32(335)
	v8852 = *(*int32)(unsafe.Add(mBase, uint32(v8839)+4))
	switch v8852 - int32(1) {
	case 0, 2:
		v8856 = int32(2)
		goto L1121
	default:
		goto L1122
	case 3, 4:
		v8857 = int32(3)
		goto L1120
	}
L1120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+76)) = v8838
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+72)) = v8857
	v8861 = *(*int32)(unsafe.Add(mBase, uint32(v8839)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+72)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+80)) = v8861
	v8865 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+4)) = v8865
	v8867 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8841)+8)) = v8867
	v8869 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8841)+16)) = v8869
	v8871 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8841)+24)) = v8871
	v8873 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8874 = *(*int32)(unsafe.Add(mBase, uint32(v8873)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8841)+32)) = v8874
	v8876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8841)+36)) = uint8(v8876)
	v8878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8841)+37)) = uint8(v8878)
	v8880 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v8880
	v12392 = v8841
	v12397 = v51
	goto L3
L1121:
	;
	v8857 = v8856
	goto L1120
L1122:
	;
	v8856 = int32(1)
	goto L1121
L1123:
	;
	v9018 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v9019 = F_order_qual_clauses(m, l0, v9018)
	mBase = m.M
	v9020 = m.ExcPending
	if v9020 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1124:
	;
	v8889 = *(*int32)(unsafe.Add(mBase, uint32(v8885)+4))
	if v8889 <= int32(0) {
		v8970 = v8882
		goto L1123
	} else {
		goto L1125
	}
L1125:
	;
	v8892 = *(*int32)(unsafe.Add(mBase, uint32(v8884)+8))
	v8893 = v8882
	v8895 = v8882
	v8897 = int32(1)
	goto L1126
L1126:
	;
	v8941 = *(*int32)(unsafe.Add(mBase, uint32(v8885)+12))
	v8945 = *(*int32)(unsafe.Add(mBase, uint32(v8941+v8895<<(uint(int32(2))%32))))
	v8946 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8946 != 0 {
		goto L1128
	} else {
		goto L1129
	}
L1127:
	;
	v8970 = v8964
	goto L1123
L1128:
	;
	v8947 = F_replace_nestloop_params_mutator(m, v8945, l0)
	mBase = m.M
	v8948 = m.ExcPending
	if v8948 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1129:
	;
	v8949 = v8945
	goto L1130
L1130:
	;
	v8951 = int32(0)
	v8953 = F_makeTargetEntry(m, v8949, base.I32_extend16_s(v8897), v8951, v8951)
	mBase = m.M
	v8954 = m.ExcPending
	if v8954 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1131:
	;
	v8949 = v8947
	goto L1130
L1132:
	;
	if v8892 != 0 {
		goto L1133
	} else {
		goto L1134
	}
L1133:
	;
	v8960 = *(*int32)(unsafe.Add(mBase, uint32(v8892+v8897<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8953)+16)) = v8960
	goto L1135
L1134:
	;
	goto L1135
L1135:
	;
	v8964 = F_lappend(m, v8893, v8953)
	mBase = m.M
	v8965 = m.ExcPending
	if v8965 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1136:
	;
	v8967 = v8895 + int32(1)
	v8968 = *(*int32)(unsafe.Add(mBase, uint32(v8885)+4))
	if v8967 < v8968 {
		v8893 = v8964
		v8895 = v8967
		v8897 = v8897 + int32(1)
		goto L1126
	} else {
		goto L1137
	}
L1137:
	;
	goto L1127
L1138:
	;
	v9021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9023 = F_palloc0(m, int32(88))
	mBase = m.M
	v9024 = m.ExcPending
	if v9024 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9023)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+44)) = v8970
	*(*int32)(unsafe.Add(mBase, uint32(v9023))) = int32(335)
	v9034 = *(*int32)(unsafe.Add(mBase, uint32(v9021)+4))
	switch v9034 - int32(1) {
	case 0, 2:
		v9038 = int32(2)
		goto L1141
	default:
		goto L1142
	case 3, 4:
		v9040 = int32(3)
		goto L1140
	}
L1140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+76)) = v9019
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+72)) = v9040
	v9043 = *(*int32)(unsafe.Add(mBase, uint32(v9021)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+80)) = v9043
	v9045 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+4)) = v9045
	v9047 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9023)+8)) = v9047
	v9049 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9023)+16)) = v9049
	v9051 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9023)+24)) = v9051
	v9053 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9054 = *(*int32)(unsafe.Add(mBase, uint32(v9053)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+32)) = v9054
	v9056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9023)+36)) = uint8(v9056)
	v9058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9023)+37)) = uint8(v9058)
	v12392 = v9023
	v12397 = v51
	goto L3
L1141:
	;
	v9040 = v9038
	goto L1140
L1142:
	;
	v9038 = int32(1)
	goto L1141
L1143:
	;
	v12392 = v9060
	v12397 = v51
	goto L3
L1144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9063))) = int32(339)
	v9067 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9068 = *(*int32)(unsafe.Add(mBase, uint32(v9067)+4))
	if v9068 == int32(0) {
		goto L1146
	} else {
		goto L1147
	}
L1145:
	;
	v9256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v9257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+4)) = v9258
	v9260 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9063)+8)) = v9260
	v9262 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9063)+16)) = v9262
	v9264 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9063)+24)) = v9264
	v9266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9267 = *(*int32)(unsafe.Add(mBase, uint32(v9266)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+32)) = v9267
	v9269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9063)+36)) = uint8(v9269)
	v9271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v9272 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+56)) = v9272
	*(*int64)(unsafe.Add(mBase, uint32(v9063)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+44)) = v9217
	*(*uint8)(unsafe.Add(mBase, uint32(v9063)+37)) = uint8(v9271)
	v9278 = *(*int32)(unsafe.Add(mBase, uint32(v9257)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+72)) = v9278
	v9280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+76)) = v9280
	v9282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9283 = *(*int32)(unsafe.Add(mBase, uint32(v9282)+8))
	v9289 = v9063 + int32(88)
	v9296 = F_prepare_sort_from_pathkeys(m, v9063, v9256, v9283, v9272, int32(1), v9063+int32(84), v9289, v9063+int32(92), v9063+int32(96), v9063+int32(100))
	mBase = m.M
	v9297 = m.ExcPending
	if v9297 != 0 {
		goto L1
	} else {
		goto L1162
	}
L1146:
	;
	v9205 = int32(0)
	v9217 = v9205
	v9255 = v9205
	goto L1145
L1147:
	;
	v9072 = int32(0)
	v9073 = *(*int32)(unsafe.Add(mBase, uint32(v9068)+4))
	if v9073 <= v9072 {
		v9217 = v4
		v9255 = v9072
		goto L1145
	} else {
		goto L1148
	}
L1148:
	;
	v9076 = *(*int32)(unsafe.Add(mBase, uint32(v9067)+8))
	v9081 = int32(1)
	v9083 = v4
	v9087 = v4
	goto L1149
L1149:
	;
	v9125 = *(*int32)(unsafe.Add(mBase, uint32(v9068)+12))
	v9129 = *(*int32)(unsafe.Add(mBase, uint32(v9125+v9083<<(uint(int32(2))%32))))
	v9130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9130 != 0 {
		goto L1151
	} else {
		goto L1152
	}
L1150:
	;
	if v9148 == int32(0) {
		goto L1146
	} else {
		goto L1161
	}
L1151:
	;
	v9131 = F_replace_nestloop_params_mutator(m, v9129, l0)
	mBase = m.M
	v9132 = m.ExcPending
	if v9132 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1152:
	;
	v9133 = v9129
	goto L1153
L1153:
	;
	v9135 = int32(0)
	v9137 = F_makeTargetEntry(m, v9133, base.I32_extend16_s(v9081), v9135, v9135)
	mBase = m.M
	v9138 = m.ExcPending
	if v9138 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1154:
	;
	v9133 = v9131
	goto L1153
L1155:
	;
	if v9076 != 0 {
		goto L1156
	} else {
		goto L1157
	}
L1156:
	;
	v9144 = *(*int32)(unsafe.Add(mBase, uint32(v9076+v9081<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9137)+16)) = v9144
	goto L1158
L1157:
	;
	goto L1158
L1158:
	;
	v9148 = F_lappend(m, v9087, v9137)
	mBase = m.M
	v9149 = m.ExcPending
	if v9149 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	v9151 = v9083 + int32(1)
	v9152 = *(*int32)(unsafe.Add(mBase, uint32(v9068)+4))
	if v9151 < v9152 {
		v9081 = v9081 + int32(1)
		v9083 = v9151
		v9087 = v9148
		goto L1149
	} else {
		goto L1160
	}
L1160:
	;
	goto L1150
L1161:
	;
	v9156 = *(*int32)(unsafe.Add(mBase, uint32(v9148)+4))
	v9217 = v9148
	v9255 = v9156
	goto L1145
L1162:
	;
	v9298 = *(*int32)(unsafe.Add(mBase, uint32(v9063)+44))
	if v9298 != 0 {
		goto L1163
	} else {
		goto L1164
	}
L1163:
	;
	v9299 = *(*int32)(unsafe.Add(mBase, uint32(v9298)+4))
	v9300 = v9299
	goto L1165
L1164:
	;
	v9300 = v4
	goto L1165
L1165:
	;
	v9301 = int32(0)
	v9302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v9302 == v9301 {
		v9674 = v9301
		goto L1167
	} else {
		goto L1168
	}
L1166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9745 = m.ExcPending
	if v9745 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+104)) = int32(-1)
	v9715 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[5])))
	if v9715 == int32(0) {
		goto L1238
	} else {
		goto L1239
	}
L1168:
	;
	v9305 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+4))
	if v9305 <= int32(0) {
		v9674 = v9301
		goto L1167
	} else {
		goto L1169
	}
L1169:
	;
	v9315 = int32(0)
	v9319 = v9301
	goto L1170
L1170:
	;
	v9357 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+12))
	v9361 = *(*int32)(unsafe.Add(mBase, uint32(v9357+v9315<<(uint(int32(2))%32))))
	v9363 = F_create_plan_recurse(m, l0, v9361, int32(1))
	mBase = m.M
	v9364 = m.ExcPending
	if v9364 != 0 {
		goto L1
	} else {
		goto L1172
	}
L1171:
	;
	v9674 = v9658
	goto L1167
L1172:
	;
	v9365 = *(*int32)(unsafe.Add(mBase, uint32(v9361)+8))
	v9366 = *(*int32)(unsafe.Add(mBase, uint32(v9365)+8))
	v9367 = *(*int32)(unsafe.Add(mBase, uint32(v9289)))
	v9379 = F_prepare_sort_from_pathkeys(m, v9363, v9256, v9366, v9367, int32(0), v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160), v51+int32(84))
	mBase = m.M
	v9380 = m.ExcPending
	if v9380 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1173:
	;
	v9381 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v9382 = *(*int32)(unsafe.Add(mBase, uint32(v9289)))
	v9383 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v9385 = v9383 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v9385) {
		goto L1177
	} else {
		goto L1178
	}
L1174:
	;
	if v9447 != 0 {
		goto L1166
	} else {
		goto L1192
	}
L1175:
	;
	v9447 = int32(0)
	goto L1174
L1176:
	;
	v9421 = v9416
	v9422 = v9417
	v9423 = v9418
	goto L1186
L1177:
	;
	if (v9381|v9382)&int32(3) != 0 {
		v9416 = v9381
		v9417 = v9382
		v9418 = v9385
		goto L1176
	} else {
		goto L1180
	}
L1178:
	;
	v9409 = v9381
	v9410 = v9382
	v9411 = v9385
	goto L1179
L1179:
	;
	if v9411 == int32(0) {
		goto L1175
	} else {
		goto L1185
	}
L1180:
	;
	v9393 = v9381
	v9394 = v9382
	v9395 = v9385
	goto L1181
L1181:
	;
	v9398 = *(*int32)(unsafe.Add(mBase, uint32(v9393)))
	v9399 = *(*int32)(unsafe.Add(mBase, uint32(v9394)))
	if v9398 != v9399 {
		v9416 = v9393
		v9417 = v9394
		v9418 = v9395
		goto L1176
	} else {
		goto L1183
	}
L1182:
	;
	v9409 = v9404
	v9410 = v9402
	v9411 = v9406
	goto L1179
L1183:
	;
	v9401 = int32(4)
	v9402 = v9394 + v9401
	v9404 = v9393 + v9401
	v9406 = v9395 - v9401
	if base.Ui32(int32(3)) < base.Ui32(v9406) {
		v9393 = v9404
		v9394 = v9402
		v9395 = v9406
		goto L1181
	} else {
		goto L1184
	}
L1184:
	;
	goto L1182
L1185:
	;
	v9416 = v9409
	v9417 = v9410
	v9418 = v9411
	goto L1176
L1186:
	;
	v9426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9421))))
	v9427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9422))))
	if v9426 == v9427 {
		goto L1188
	} else {
		goto L1189
	}
L1187:
	;
	v9447 = v9426 - v9427
	goto L1174
L1188:
	;
	v9429 = int32(1)
	v9434 = v9423 - v9429
	if v9434 != 0 {
		v9421 = v9421 + v9429
		v9422 = v9422 + v9429
		v9423 = v9434
		goto L1186
	} else {
		goto L1191
	}
L1189:
	;
	goto L1190
L1190:
	;
	goto L1187
L1191:
	;
	goto L1175
L1192:
	;
	v9448 = *(*int32)(unsafe.Add(mBase, uint32(v9361)+64))
	v9450 = v51 + int32(80)
	if v9256 == v9448 {
		goto L1196
	} else {
		goto L1197
	}
L1193:
	;
	v9658 = F_lappend(m, v9319, v9651)
	mBase = m.M
	v9659 = m.ExcPending
	if v9659 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1194:
	;
	if v9528 != 0 {
		goto L1226
	} else {
		goto L1227
	}
L1195:
	;
	v9516 = *(*int32)(unsafe.Add(mBase, uint32(v9256)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9516
	v9528 = int32(1)
	goto L1194
L1196:
	;
	if v9256 != 0 {
		goto L1195
	} else {
		goto L1199
	}
L1197:
	;
	goto L1198
L1198:
	;
	if v9256 == int32(0) {
		goto L1200
	} else {
		goto L1201
	}
L1199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = int32(0)
	v9528 = int32(1)
	goto L1194
L1200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = int32(0)
	v9528 = int32(1)
	goto L1194
L1201:
	;
	goto L1202
L1202:
	;
	if v9448 == int32(0) {
		goto L1203
	} else {
		goto L1204
	}
L1203:
	;
	v9468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9468
	v9528 = v9468
	goto L1194
L1204:
	;
	goto L1205
L1205:
	;
	v9471 = *(*int32)(unsafe.Add(mBase, uint32(v9448)+4))
	v9472 = int32(0)
	if v9472 < v9471 {
		goto L1206
	} else {
		goto L1207
	}
L1206:
	;
	v9475 = v9471
	goto L1208
L1207:
	;
	v9475 = v9472
	goto L1208
L1208:
	;
	v9476 = *(*int32)(unsafe.Add(mBase, uint32(v9256)+4))
	v9481 = int32(0)
	goto L1209
L1209:
	;
	if v9481 < v9476 {
		goto L1211
	} else {
		goto L1212
	}
L1211:
	;
	v9488 = *(*int32)(unsafe.Add(mBase, uint32(v9256)+12))
	v9492 = v9488 + v9481<<(uint(int32(2))%32)
	goto L1213
L1212:
	;
	v9492 = int32(0)
	goto L1213
L1213:
	;
	if v9481 == v9475 {
		goto L1214
	} else {
		goto L1215
	}
L1214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9475
	v9528 = base.B2i32(v9492 == int32(0))
	goto L1194
L1215:
	;
	goto L1216
L1216:
	;
	v9498 = base.B2i32(v9492 == int32(0))
	if v9492 == int32(0) {
		goto L1217
	} else {
		goto L1218
	}
L1217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9481
	v9528 = v9498
	goto L1194
L1218:
	;
	goto L1219
L1219:
	;
	v9502 = *(*int32)(unsafe.Add(mBase, uint32(v9448)+12))
	if v9502 == int32(0) {
		goto L1220
	} else {
		goto L1221
	}
L1220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9481
	v9528 = v9498
	goto L1194
L1221:
	;
	goto L1222
L1222:
	;
	v9506 = *(*int32)(unsafe.Add(mBase, uint32(v9492)))
	v9510 = *(*int32)(unsafe.Add(mBase, uint32(v9502+v9481<<(uint(int32(2))%32))))
	if v9506 != v9510 {
		goto L1223
	} else {
		goto L1224
	}
L1223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9450))) = v9481
	v9528 = int32(0)
	goto L1194
L1224:
	;
	v9481 = v9481 + int32(1)
	goto L1209
L1226:
	;
	v9651 = v9379
	goto L1193
L1227:
	;
	goto L1228
L1228:
	;
	v9529 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	v9530 = *(*int32)(unsafe.Add(mBase, uint32(v51)+160))
	v9531 = *(*int32)(unsafe.Add(mBase, uint32(v51)+164))
	v9533 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[0])))
	if v9533 != int32(1) {
		goto L1229
	} else {
		goto L1230
	}
L1229:
	;
	v9583 = F_palloc0(m, int32(96))
	mBase = m.M
	v9584 = m.ExcPending
	if v9584 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1230:
	;
	v9536 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if v9536 <= int32(0) {
		goto L1229
	} else {
		goto L1231
	}
L1231:
	;
	v9540 = F_palloc0(m, int32(104))
	mBase = m.M
	v9541 = m.ExcPending
	if v9541 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9540))) = int32(367)
	v9544 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+96)) = v9536
	v9546 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+56)) = v9546
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+52)) = v9379
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+48)) = v9546
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+44)) = v9544
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+88)) = v9529
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+84)) = v9530
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+80)) = v9531
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+76)) = v9381
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+72)) = v9383
	v9559 = *(*int32)(unsafe.Add(mBase, uint32(v9540)+4))
	v9560 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+8))
	v9561 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+16))
	v9562 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+24))
	v9563 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+32))
	v9565 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v9566 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	F_cost_incremental_sort(m, v51+int32(88), l0, v9256, v9536, v9559, v9560, v9561, v9562, v9563, v9565, v9566)
	mBase = m.M
	v9568 = m.ExcPending
	if v9568 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1233:
	;
	v9569 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v9540)+8)) = v9569
	v9571 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v9540)+16)) = v9571
	v9573 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v9540)+24)) = v9573
	v9575 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+32))
	v9576 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9540)+36)) = uint8(v9576)
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+32)) = v9575
	v9579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9379)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9540)+37)) = uint8(v9579)
	v9651 = v9540
	goto L1193
L1234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9583))) = int32(366)
	v9587 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+44)) = v9587
	v9589 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+4))
	v9590 = int32(_a_F_create_plan_recurse_0)
	v9591 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+88)) = v9529
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+84)) = v9530
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+80)) = v9531
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+76)) = v9381
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+72)) = v9383
	v9597 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+56)) = v9597
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+52)) = v9379
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+48)) = v9597
	v9602 = int32(1)
	v9604 = v9589 + (v9591 ^ v9602)
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+4)) = v9604
	v9607 = v51 + int32(88)
	v9608 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+16))
	v9609 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+24))
	v9610 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+32))
	v9613 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v9614 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v9616 = m.G0
	v9617 = int32(16)
	v9618 = v9616 - v9617
	m.G0 = v9618
	F_cost_tuplesort(m, v9618+int32(8), v9618, v9609, v9610, float64(0), v9613, v9614)
	mBase = m.M
	v9623 = *(*float64)(unsafe.Add(mBase, uint32(v9618)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v9607)+32)) = v9609
	v9626 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v9627 = base.F64_add(v9608, v9623)
	*(*float64)(unsafe.Add(mBase, uint32(v9607)+48)) = v9627
	*(*int32)(unsafe.Add(mBase, uint32(v9607)+40)) = v9604 + (v9626 ^ v9602)
	v9633 = *(*float64)(unsafe.Add(mBase, uint32(v9618)))
	*(*float64)(unsafe.Add(mBase, uint32(v9607)+56)) = base.F64_add(v9627, v9633)
	m.G0 = v9618 + v9617
	goto L1235
L1235:
	;
	v9639 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v9583)+8)) = v9639
	v9641 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v9583)+16)) = v9641
	v9643 = *(*float64)(unsafe.Add(mBase, uint32(v9379)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v9583)+24)) = v9643
	v9645 = *(*int32)(unsafe.Add(mBase, uint32(v9379)+32))
	v9646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9583)+36)) = uint8(v9646)
	*(*int32)(unsafe.Add(mBase, uint32(v9583)+32)) = v9645
	v9649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9379)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9583)+37)) = uint8(v9649)
	v9651 = v9583
	goto L1193
L1236:
	;
	v9661 = v9315 + int32(1)
	v9662 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+4))
	if v9661 < v9662 {
		v9315 = v9661
		v9319 = v9658
		goto L1170
	} else {
		goto L1237
	}
L1237:
	;
	goto L1171
L1238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+80)) = v9674
	if base.B2i32(l2&int32(3) == int32(0))|base.B2i32(v9255 == v9300) != 0 {
		v12392 = v9063
		v12397 = v51
		goto L3
	} else {
		goto L1243
	}
L1239:
	;
	v9718 = *(*int32)(unsafe.Add(mBase, uint32(v9257)+204))
	v9720 = F_extract_actual_clauses(m, v9718, int32(0))
	mBase = m.M
	v9721 = m.ExcPending
	if v9721 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	if v9720 == int32(0) {
		goto L1238
	} else {
		goto L1241
	}
L1241:
	;
	v9724 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v9725 = F_make_partition_pruneinfo(m, l0, v9257, v9724, v9720)
	mBase = m.M
	v9726 = m.ExcPending
	if v9726 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9063)+104)) = v9725
	goto L1238
L1243:
	;
	v9736 = *(*int32)(unsafe.Add(mBase, uint32(v9063)+44))
	v9737 = F_list_copy_head(m, v9736, v9255)
	mBase = m.M
	v9738 = m.ExcPending
	if v9738 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	v9739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063)+37)))
	v9740 = F_inject_projection_plan(m, v9063, v9737, v9739)
	mBase = m.M
	v9741 = m.ExcPending
	if v9741 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	v12392 = v9740
	v12397 = v51
	goto L3
L1246:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_42), int32(0))
	mBase = m.M
	v9749 = m.ExcPending
	if v9749 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1540), int32(_a_F_create_plan_recurse_43))
	mBase = m.M
	v9754 = m.ExcPending
	if v9754 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1249:
	;
	v9944 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9945 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v9946 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+172)) = v9946
	*(*int32)(unsafe.Add(mBase, uint32(v51)+168)) = v9946
	*(*int32)(unsafe.Add(mBase, uint32(v51)+164)) = v9946
	*(*int32)(unsafe.Add(mBase, uint32(v51)+160)) = v9946
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = v9946
	v9956 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v9956 == v9946 {
		goto L1266
	} else {
		goto L1267
	}
L1250:
	;
	v9893 = int32(0)
	v9901 = v9893
	v9943 = v9893
	goto L1249
L1251:
	;
	v9760 = int32(0)
	v9761 = *(*int32)(unsafe.Add(mBase, uint32(v9756)+4))
	if v9761 <= v9760 {
		v9901 = v4
		v9943 = v9760
		goto L1249
	} else {
		goto L1252
	}
L1252:
	;
	v9764 = *(*int32)(unsafe.Add(mBase, uint32(v9755)+8))
	v9768 = int32(1)
	v9770 = v4
	v9771 = v4
	goto L1253
L1253:
	;
	v9813 = *(*int32)(unsafe.Add(mBase, uint32(v9756)+12))
	v9817 = *(*int32)(unsafe.Add(mBase, uint32(v9813+v9770<<(uint(int32(2))%32))))
	v9818 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9818 != 0 {
		goto L1255
	} else {
		goto L1256
	}
L1254:
	;
	if v9836 == int32(0) {
		goto L1250
	} else {
		goto L1265
	}
L1255:
	;
	v9819 = F_replace_nestloop_params_mutator(m, v9817, l0)
	mBase = m.M
	v9820 = m.ExcPending
	if v9820 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1256:
	;
	v9821 = v9817
	goto L1257
L1257:
	;
	v9823 = int32(0)
	v9825 = F_makeTargetEntry(m, v9821, base.I32_extend16_s(v9768), v9823, v9823)
	mBase = m.M
	v9826 = m.ExcPending
	if v9826 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1258:
	;
	v9821 = v9819
	goto L1257
L1259:
	;
	if v9764 != 0 {
		goto L1260
	} else {
		goto L1261
	}
L1260:
	;
	v9832 = *(*int32)(unsafe.Add(mBase, uint32(v9764+v9768<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9825)+16)) = v9832
	goto L1262
L1261:
	;
	goto L1262
L1262:
	;
	v9836 = F_lappend(m, v9771, v9825)
	mBase = m.M
	v9837 = m.ExcPending
	if v9837 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	v9839 = v9770 + int32(1)
	v9840 = *(*int32)(unsafe.Add(mBase, uint32(v9756)+4))
	if v9839 < v9840 {
		v9768 = v9768 + int32(1)
		v9770 = v9839
		v9771 = v9836
		goto L1253
	} else {
		goto L1264
	}
L1264:
	;
	goto L1254
L1265:
	;
	v9844 = *(*int32)(unsafe.Add(mBase, uint32(v9836)+4))
	v9901 = v9836
	v9943 = v9844
	goto L1249
L1266:
	;
	v9959 = int32(0)
	v9961 = F_makeBoolConst(m, v9959, v9959)
	mBase = m.M
	v9962 = m.ExcPending
	if v9962 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1267:
	;
	goto L1268
L1268:
	;
	v10010 = F_palloc0(m, int32(96))
	mBase = m.M
	v10011 = m.ExcPending
	if v10011 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v9961
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v9961
	v9968 = F_list_make1_impl(m, int32(1), v51+int32(28))
	mBase = m.M
	v9969 = m.ExcPending
	if v9969 != 0 {
		goto L1
	} else {
		goto L1270
	}
L1270:
	;
	v9970 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9972 = F_palloc0(m, int32(88))
	mBase = m.M
	v9973 = m.ExcPending
	if v9973 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9972)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+44)) = v9901
	*(*int32)(unsafe.Add(mBase, uint32(v9972))) = int32(335)
	v9983 = *(*int32)(unsafe.Add(mBase, uint32(v9970)+4))
	switch v9983 - int32(1) {
	case 0, 2:
		v9987 = int32(2)
		goto L1273
	default:
		goto L1274
	case 3, 4:
		v9989 = int32(3)
		goto L1272
	}
L1272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+76)) = v9968
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+72)) = v9989
	v9992 = *(*int32)(unsafe.Add(mBase, uint32(v9970)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+80)) = v9992
	v9994 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+4)) = v9994
	v9996 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9972)+8)) = v9996
	v9998 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9972)+16)) = v9998
	v10000 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9972)+24)) = v10000
	v10002 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10003 = *(*int32)(unsafe.Add(mBase, uint32(v10002)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9972)+32)) = v10003
	v10005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9972)+36)) = uint8(v10005)
	v10007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9972)+37)) = uint8(v10007)
	v12392 = v9972
	v12397 = v51
	goto L3
L1273:
	;
	v9989 = v9987
	goto L1272
L1274:
	;
	v9987 = int32(1)
	goto L1273
L1275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10010)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+44)) = v9901
	*(*int32)(unsafe.Add(mBase, uint32(v10010))) = int32(338)
	v10019 = *(*int32)(unsafe.Add(mBase, uint32(v9944)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+72)) = v10019
	v10021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+76)) = v10021
	if v9945 != 0 {
		goto L1278
	} else {
		goto L1279
	}
L1276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+92)) = int32(-1)
	v10492 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[5])))
	if v10492 == int32(0) {
		goto L1367
	} else {
		goto L1368
	}
L1277:
	;
	v10059 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10059 == int32(0) {
		v10450 = v4
		v10453 = v4
		v10457 = v10058
		goto L1276
	} else {
		goto L1288
	}
L1278:
	;
	v10023 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10024 = *(*int32)(unsafe.Add(mBase, uint32(v10023)+8))
	v10037 = F_prepare_sort_from_pathkeys(m, v10010, v9945, v10024, int32(0), int32(1), v51+int32(172), v51+int32(168), v51+int32(164), v51+int32(160), v51+int32(84))
	mBase = m.M
	v10038 = m.ExcPending
	if v10038 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1279:
	;
	goto L1280
L1280:
	;
	v10044 = int32(1)
	v10046 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[6])))
	if v10046 != v10044 {
		v10057 = v4
		v10058 = v10044
		goto L1277
	} else {
		goto L1285
	}
L1281:
	;
	v10039 = *(*int32)(unsafe.Add(mBase, uint32(v10010)+44))
	if v10039 != 0 {
		goto L1282
	} else {
		goto L1283
	}
L1282:
	;
	v10040 = *(*int32)(unsafe.Add(mBase, uint32(v10039)+4))
	v10042 = v10040
	goto L1284
L1283:
	;
	v10042 = int32(0)
	goto L1284
L1284:
	;
	v10057 = v4
	v10058 = base.B2i32(v10042 == v9943)
	goto L1277
L1285:
	;
	v10049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v10049 != 0 {
		v10057 = v4
		v10058 = v10044
		goto L1277
	} else {
		goto L1286
	}
L1286:
	;
	v10050 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10050 == int32(0) {
		v10450 = v4
		v10453 = v4
		v10457 = v10044
		goto L1276
	} else {
		goto L1287
	}
L1287:
	;
	v10053 = *(*int32)(unsafe.Add(mBase, uint32(v10050)+4))
	v10057 = base.B2i32(int32(1) < v10053)
	v10058 = v10044
	goto L1277
L1288:
	;
	v10062 = *(*int32)(unsafe.Add(mBase, uint32(v10059)+4))
	if v10062 <= int32(0) {
		v10450 = v4
		v10453 = v4
		v10457 = v10058
		goto L1276
	} else {
		goto L1289
	}
L1289:
	;
	v10072 = int32(0)
	v10075 = v4
	v10078 = v4
	goto L1290
L1290:
	;
	v10114 = *(*int32)(unsafe.Add(mBase, uint32(v10059)+12))
	v10118 = *(*int32)(unsafe.Add(mBase, uint32(v10114+v10072<<(uint(int32(2))%32))))
	v10120 = F_create_plan_recurse(m, l0, v10118, int32(1))
	mBase = m.M
	v10121 = m.ExcPending
	if v10121 != 0 {
		goto L1
	} else {
		goto L1293
	}
L1291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10431 = m.ExcPending
	if v10431 != 0 {
		goto L1
	} else {
		goto L1364
	}
L1292:
	;
	goto L1291
L1293:
	;
	if v9945 == int32(0) {
		v10409 = v10120
		goto L1294
	} else {
		goto L1295
	}
L1294:
	;
	if v10057 != 0 {
		goto L1358
	} else {
		goto L1359
	}
L1295:
	;
	v10124 = *(*int32)(unsafe.Add(mBase, uint32(v10118)+8))
	v10125 = *(*int32)(unsafe.Add(mBase, uint32(v10124)+8))
	v10126 = *(*int32)(unsafe.Add(mBase, uint32(v51)+168))
	v10138 = F_prepare_sort_from_pathkeys(m, v10120, v9945, v10125, v10126, int32(0), v51+int32(80), v51+int32(76), v51+int32(72), v51+int32(68), v51-int32(-64))
	mBase = m.M
	v10139 = m.ExcPending
	if v10139 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	v10140 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	v10141 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v10143 = v10141 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v10143) {
		goto L1300
	} else {
		goto L1301
	}
L1297:
	;
	if v10205 != 0 {
		goto L1292
	} else {
		goto L1315
	}
L1298:
	;
	v10205 = int32(0)
	goto L1297
L1299:
	;
	v10179 = v10174
	v10180 = v10175
	v10181 = v10176
	goto L1309
L1300:
	;
	if (v10140|v10126)&int32(3) != 0 {
		v10174 = v10140
		v10175 = v10126
		v10176 = v10143
		goto L1299
	} else {
		goto L1303
	}
L1301:
	;
	v10167 = v10140
	v10168 = v10126
	v10169 = v10143
	goto L1302
L1302:
	;
	if v10169 == int32(0) {
		goto L1298
	} else {
		goto L1308
	}
L1303:
	;
	v10151 = v10140
	v10152 = v10126
	v10153 = v10143
	goto L1304
L1304:
	;
	v10156 = *(*int32)(unsafe.Add(mBase, uint32(v10151)))
	v10157 = *(*int32)(unsafe.Add(mBase, uint32(v10152)))
	if v10156 != v10157 {
		v10174 = v10151
		v10175 = v10152
		v10176 = v10153
		goto L1299
	} else {
		goto L1306
	}
L1305:
	;
	v10167 = v10162
	v10168 = v10160
	v10169 = v10164
	goto L1302
L1306:
	;
	v10159 = int32(4)
	v10160 = v10152 + v10159
	v10162 = v10151 + v10159
	v10164 = v10153 - v10159
	if base.Ui32(int32(3)) < base.Ui32(v10164) {
		v10151 = v10162
		v10152 = v10160
		v10153 = v10164
		goto L1304
	} else {
		goto L1307
	}
L1307:
	;
	goto L1305
L1308:
	;
	v10174 = v10167
	v10175 = v10168
	v10176 = v10169
	goto L1299
L1309:
	;
	v10184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10179))))
	v10185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10180))))
	if v10184 == v10185 {
		goto L1311
	} else {
		goto L1312
	}
L1310:
	;
	v10205 = v10184 - v10185
	goto L1297
L1311:
	;
	v10187 = int32(1)
	v10192 = v10181 - v10187
	if v10192 != 0 {
		v10179 = v10179 + v10187
		v10180 = v10180 + v10187
		v10181 = v10192
		goto L1309
	} else {
		goto L1314
	}
L1312:
	;
	goto L1313
L1313:
	;
	goto L1310
L1314:
	;
	goto L1298
L1315:
	;
	v10206 = *(*int32)(unsafe.Add(mBase, uint32(v10118)+64))
	v10208 = v51 + int32(60)
	if v9945 == v10206 {
		goto L1318
	} else {
		goto L1319
	}
L1316:
	;
	if v10286 != 0 {
		goto L1348
	} else {
		goto L1349
	}
L1317:
	;
	v10274 = *(*int32)(unsafe.Add(mBase, uint32(v9945)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10274
	v10286 = int32(1)
	goto L1316
L1318:
	;
	if v9945 != 0 {
		goto L1317
	} else {
		goto L1321
	}
L1319:
	;
	goto L1320
L1320:
	;
	if v9945 == int32(0) {
		goto L1322
	} else {
		goto L1323
	}
L1321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = int32(0)
	v10286 = int32(1)
	goto L1316
L1322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = int32(0)
	v10286 = int32(1)
	goto L1316
L1323:
	;
	goto L1324
L1324:
	;
	if v10206 == int32(0) {
		goto L1325
	} else {
		goto L1326
	}
L1325:
	;
	v10226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10226
	v10286 = v10226
	goto L1316
L1326:
	;
	goto L1327
L1327:
	;
	v10229 = *(*int32)(unsafe.Add(mBase, uint32(v10206)+4))
	v10230 = int32(0)
	if v10230 < v10229 {
		goto L1328
	} else {
		goto L1329
	}
L1328:
	;
	v10233 = v10229
	goto L1330
L1329:
	;
	v10233 = v10230
	goto L1330
L1330:
	;
	v10234 = *(*int32)(unsafe.Add(mBase, uint32(v9945)+4))
	v10239 = int32(0)
	goto L1331
L1331:
	;
	if v10239 < v10234 {
		goto L1333
	} else {
		goto L1334
	}
L1333:
	;
	v10246 = *(*int32)(unsafe.Add(mBase, uint32(v9945)+12))
	v10250 = v10246 + v10239<<(uint(int32(2))%32)
	goto L1335
L1334:
	;
	v10250 = int32(0)
	goto L1335
L1335:
	;
	if v10239 == v10233 {
		goto L1336
	} else {
		goto L1337
	}
L1336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10233
	v10286 = base.B2i32(v10250 == int32(0))
	goto L1316
L1337:
	;
	goto L1338
L1338:
	;
	v10256 = base.B2i32(v10250 == int32(0))
	if v10250 == int32(0) {
		goto L1339
	} else {
		goto L1340
	}
L1339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10239
	v10286 = v10256
	goto L1316
L1340:
	;
	goto L1341
L1341:
	;
	v10260 = *(*int32)(unsafe.Add(mBase, uint32(v10206)+12))
	if v10260 == int32(0) {
		goto L1342
	} else {
		goto L1343
	}
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10239
	v10286 = v10256
	goto L1316
L1343:
	;
	goto L1344
L1344:
	;
	v10264 = *(*int32)(unsafe.Add(mBase, uint32(v10250)))
	v10268 = *(*int32)(unsafe.Add(mBase, uint32(v10260+v10239<<(uint(int32(2))%32))))
	if v10264 != v10268 {
		goto L1345
	} else {
		goto L1346
	}
L1345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10208))) = v10239
	v10286 = int32(0)
	goto L1316
L1346:
	;
	v10239 = v10239 + int32(1)
	goto L1331
L1348:
	;
	v10409 = v10138
	goto L1294
L1349:
	;
	goto L1350
L1350:
	;
	v10287 = *(*int32)(unsafe.Add(mBase, uint32(v51)+64))
	v10288 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
	v10289 = *(*int32)(unsafe.Add(mBase, uint32(v51)+72))
	v10291 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[0])))
	if v10291 != int32(1) {
		goto L1351
	} else {
		goto L1352
	}
L1351:
	;
	v10341 = F_palloc0(m, int32(96))
	mBase = m.M
	v10342 = m.ExcPending
	if v10342 != 0 {
		goto L1
	} else {
		goto L1356
	}
L1352:
	;
	v10294 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	if v10294 <= int32(0) {
		goto L1351
	} else {
		goto L1353
	}
L1353:
	;
	v10298 = F_palloc0(m, int32(104))
	mBase = m.M
	v10299 = m.ExcPending
	if v10299 != 0 {
		goto L1
	} else {
		goto L1354
	}
L1354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10298))) = int32(367)
	v10302 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+96)) = v10294
	v10304 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+56)) = v10304
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+52)) = v10138
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+48)) = v10304
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+44)) = v10302
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+88)) = v10287
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+84)) = v10288
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+80)) = v10289
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+76)) = v10140
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+72)) = v10141
	v10317 = *(*int32)(unsafe.Add(mBase, uint32(v10298)+4))
	v10318 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+8))
	v10319 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+16))
	v10320 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+24))
	v10321 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+32))
	v10323 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v10324 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	F_cost_incremental_sort(m, v51+int32(88), l0, v9945, v10294, v10317, v10318, v10319, v10320, v10321, v10323, v10324)
	mBase = m.M
	v10326 = m.ExcPending
	if v10326 != 0 {
		goto L1
	} else {
		goto L1355
	}
L1355:
	;
	v10327 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v10298)+8)) = v10327
	v10329 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v10298)+16)) = v10329
	v10331 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10298)+24)) = v10331
	v10333 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+32))
	v10334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10298)+36)) = uint8(v10334)
	*(*int32)(unsafe.Add(mBase, uint32(v10298)+32)) = v10333
	v10337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10138)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10298)+37)) = uint8(v10337)
	v10409 = v10298
	goto L1294
L1356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10341))) = int32(366)
	v10345 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+44)) = v10345
	v10347 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+4))
	v10348 = int32(_a_F_create_plan_recurse_0)
	v10349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+88)) = v10287
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+84)) = v10288
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+80)) = v10289
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+76)) = v10140
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+72)) = v10141
	v10355 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+56)) = v10355
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+52)) = v10138
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+48)) = v10355
	v10360 = int32(1)
	v10362 = v10347 + (v10349 ^ v10360)
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+4)) = v10362
	v10365 = v51 + int32(88)
	v10366 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+16))
	v10367 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+24))
	v10368 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+32))
	v10371 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v10372 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v10374 = m.G0
	v10375 = int32(16)
	v10376 = v10374 - v10375
	m.G0 = v10376
	F_cost_tuplesort(m, v10376+int32(8), v10376, v10367, v10368, float64(0), v10371, v10372)
	mBase = m.M
	v10381 = *(*float64)(unsafe.Add(mBase, uint32(v10376)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v10365)+32)) = v10367
	v10384 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v10385 = base.F64_add(v10366, v10381)
	*(*float64)(unsafe.Add(mBase, uint32(v10365)+48)) = v10385
	*(*int32)(unsafe.Add(mBase, uint32(v10365)+40)) = v10362 + (v10384 ^ v10360)
	v10391 = *(*float64)(unsafe.Add(mBase, uint32(v10376)))
	*(*float64)(unsafe.Add(mBase, uint32(v10365)+56)) = base.F64_add(v10385, v10391)
	m.G0 = v10376 + v10375
	goto L1357
L1357:
	;
	v10397 = *(*float64)(unsafe.Add(mBase, uint32(v51)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v10341)+8)) = v10397
	v10399 = *(*float64)(unsafe.Add(mBase, uint32(v51)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v10341)+16)) = v10399
	v10401 = *(*float64)(unsafe.Add(mBase, uint32(v10138)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10341)+24)) = v10401
	v10403 = *(*int32)(unsafe.Add(mBase, uint32(v10138)+32))
	v10404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10341)+36)) = uint8(v10404)
	*(*int32)(unsafe.Add(mBase, uint32(v10341)+32)) = v10403
	v10407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10138)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10341)+37)) = uint8(v10407)
	v10409 = v10341
	goto L1294
L1358:
	;
	v10418 = F_mark_async_capable_plan(m, v10409, v10118)
	mBase = m.M
	v10419 = m.ExcPending
	if v10419 != 0 {
		goto L1
	} else {
		goto L1361
	}
L1359:
	;
	v10421 = v10075
	goto L1360
L1360:
	;
	v10422 = F_lappend(m, v10078, v10409)
	mBase = m.M
	v10423 = m.ExcPending
	if v10423 != 0 {
		goto L1
	} else {
		goto L1362
	}
L1361:
	;
	v10421 = v10418 + v10075
	goto L1360
L1362:
	;
	v10425 = v10072 + int32(1)
	v10426 = *(*int32)(unsafe.Add(mBase, uint32(v10059)+4))
	if v10425 < v10426 {
		v10072 = v10425
		v10075 = v10421
		v10078 = v10422
		goto L1290
	} else {
		goto L1363
	}
L1363:
	;
	v10450 = v10421
	v10453 = v10422
	v10457 = v10058
	goto L1276
L1364:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_44), int32(0))
	mBase = m.M
	v10435 = m.ExcPending
	if v10435 != 0 {
		goto L1
	} else {
		goto L1365
	}
L1365:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1342), int32(_a_F_create_plan_recurse_45))
	mBase = m.M
	v10440 = m.ExcPending
	if v10440 != 0 {
		goto L1
	} else {
		goto L1366
	}
L1366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+84)) = v10450
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+80)) = v10453
	v10519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+88)) = v10519
	v10521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+4)) = v10521
	v10523 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v10010)+8)) = v10523
	v10525 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v10010)+16)) = v10525
	v10527 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v10010)+24)) = v10527
	v10529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10530 = *(*int32)(unsafe.Add(mBase, uint32(v10529)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+32)) = v10530
	v10532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10010)+36)) = uint8(v10532)
	v10534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10010)+37)) = uint8(v10534)
	if base.B2i32(l2&int32(3) == int32(0))|v10457 != 0 {
		v12392 = v10010
		v12397 = v51
		goto L3
	} else {
		goto L1378
	}
L1368:
	;
	v10495 = *(*int32)(unsafe.Add(mBase, uint32(v9944)+204))
	v10497 = F_extract_actual_clauses(m, v10495, int32(0))
	mBase = m.M
	v10498 = m.ExcPending
	if v10498 != 0 {
		goto L1
	} else {
		goto L1369
	}
L1369:
	;
	v10499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10499 != 0 {
		goto L1370
	} else {
		goto L1371
	}
L1370:
	;
	v10500 = *(*int32)(unsafe.Add(mBase, uint32(v10499)+16))
	v10502 = F_extract_actual_clauses(m, v10500, int32(0))
	mBase = m.M
	v10503 = m.ExcPending
	if v10503 != 0 {
		goto L1
	} else {
		goto L1373
	}
L1371:
	;
	v10508 = v10497
	goto L1372
L1372:
	;
	if v10508 == int32(0) {
		goto L1367
	} else {
		goto L1376
	}
L1373:
	;
	v10504 = F_replace_nestloop_params_mutator(m, v10502, l0)
	mBase = m.M
	v10505 = m.ExcPending
	if v10505 != 0 {
		goto L1
	} else {
		goto L1374
	}
L1374:
	;
	v10506 = F_list_concat(m, v10497, v10504)
	mBase = m.M
	v10507 = m.ExcPending
	if v10507 != 0 {
		goto L1
	} else {
		goto L1375
	}
L1375:
	;
	v10508 = v10506
	goto L1372
L1376:
	;
	v10511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v10512 = F_make_partition_pruneinfo(m, l0, v9944, v10511, v10508)
	mBase = m.M
	v10513 = m.ExcPending
	if v10513 != 0 {
		goto L1
	} else {
		goto L1377
	}
L1377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10010)+92)) = v10512
	goto L1367
L1378:
	;
	v10541 = *(*int32)(unsafe.Add(mBase, uint32(v10010)+44))
	v10542 = F_list_copy_head(m, v10541, v9943)
	mBase = m.M
	v10543 = m.ExcPending
	if v10543 != 0 {
		goto L1
	} else {
		goto L1379
	}
L1379:
	;
	v10544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10010)+37)))
	v10545 = F_inject_projection_plan(m, v10010, v10542, v10544)
	mBase = m.M
	v10546 = m.ExcPending
	if v10546 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1380:
	;
	v12392 = v10545
	v12397 = v51
	goto L3
L1381:
	;
	v10682 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v10683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v10684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v10685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v10686 = *(*int32)(unsafe.Add(mBase, uint32(v10685)+8))
	v10687 = F_reparameterize_path_by_child(m, l0, v10684, v10686)
	mBase = m.M
	v10688 = m.ExcPending
	if v10688 != 0 {
		goto L1
	} else {
		goto L1396
	}
L1382:
	;
	v10553 = *(*int32)(unsafe.Add(mBase, uint32(v10549)+4))
	if v10553 <= int32(0) {
		v10636 = v10547
		goto L1381
	} else {
		goto L1383
	}
L1383:
	;
	v10556 = *(*int32)(unsafe.Add(mBase, uint32(v10548)+8))
	v10559 = v10547
	v10560 = int32(1)
	v10562 = v4
	goto L1384
L1384:
	;
	v10605 = *(*int32)(unsafe.Add(mBase, uint32(v10549)+12))
	v10609 = *(*int32)(unsafe.Add(mBase, uint32(v10605+v10562<<(uint(int32(2))%32))))
	v10610 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10610 != 0 {
		goto L1386
	} else {
		goto L1387
	}
L1385:
	;
	v10636 = v10628
	goto L1381
L1386:
	;
	v10611 = F_replace_nestloop_params_mutator(m, v10609, l0)
	mBase = m.M
	v10612 = m.ExcPending
	if v10612 != 0 {
		goto L1
	} else {
		goto L1389
	}
L1387:
	;
	v10613 = v10609
	goto L1388
L1388:
	;
	v10615 = int32(0)
	v10617 = F_makeTargetEntry(m, v10613, base.I32_extend16_s(v10560), v10615, v10615)
	mBase = m.M
	v10618 = m.ExcPending
	if v10618 != 0 {
		goto L1
	} else {
		goto L1390
	}
L1389:
	;
	v10613 = v10611
	goto L1388
L1390:
	;
	if v10556 != 0 {
		goto L1391
	} else {
		goto L1392
	}
L1391:
	;
	v10624 = *(*int32)(unsafe.Add(mBase, uint32(v10556+v10560<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v10617)+16)) = v10624
	goto L1393
L1392:
	;
	goto L1393
L1393:
	;
	v10628 = F_lappend(m, v10559, v10617)
	mBase = m.M
	v10629 = m.ExcPending
	if v10629 != 0 {
		goto L1
	} else {
		goto L1394
	}
L1394:
	;
	v10631 = v10562 + int32(1)
	v10632 = *(*int32)(unsafe.Add(mBase, uint32(v10549)+4))
	if v10631 < v10632 {
		v10559 = v10628
		v10560 = v10560 + int32(1)
		v10562 = v10631
		goto L1384
	} else {
		goto L1395
	}
L1395:
	;
	goto L1385
L1396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v10687
	v10690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v10692 = F_create_plan_recurse(m, l0, v10690, int32(0))
	mBase = m.M
	v10693 = m.ExcPending
	if v10693 != 0 {
		goto L1
	} else {
		goto L1397
	}
L1397:
	;
	v10694 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v10695 = *(*int32)(unsafe.Add(mBase, uint32(v10694)+8))
	v10696 = *(*int32)(unsafe.Add(mBase, uint32(v10695)+8))
	v10697 = *(*int32)(unsafe.Add(mBase, uint32(v10695)+252))
	if v10697 != 0 {
		goto L1398
	} else {
		goto L1399
	}
L1398:
	;
	v10698 = F_bms_union(m, v10696, v10697)
	mBase = m.M
	v10699 = m.ExcPending
	if v10699 != 0 {
		goto L1
	} else {
		goto L1401
	}
L1399:
	;
	v10700 = v10696
	goto L1400
L1400:
	;
	v10701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v10702 = F_bms_union(m, v10701, v10700)
	mBase = m.M
	v10703 = m.ExcPending
	if v10703 != 0 {
		goto L1
	} else {
		goto L1402
	}
L1401:
	;
	v10700 = v10698
	goto L1400
L1402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v10702
	v10705 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v10707 = F_create_plan_recurse(m, l0, v10705, int32(0))
	mBase = m.M
	v10708 = m.ExcPending
	if v10708 != 0 {
		goto L1
	} else {
		goto L1403
	}
L1403:
	;
	v10709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	F_bms_free(m, v10709)
	mBase = m.M
	v10711 = m.ExcPending
	if v10711 != 0 {
		goto L1
	} else {
		goto L1404
	}
L1404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v10683
	v10713 = F_order_qual_clauses(m, l0, v10682)
	mBase = m.M
	v10714 = m.ExcPending
	if v10714 != 0 {
		goto L1
	} else {
		goto L1405
	}
L1405:
	;
	v10716 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v10716)%32)&int32(174) != 0 {
		goto L1407
	} else {
		goto L1408
	}
L1406:
	;
	v10735 = int32(0)
	v10736 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10736 == v10735 {
		v10933 = v10735
		goto L1412
	} else {
		goto L1413
	}
L1407:
	;
	v10720 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10721 = *(*int32)(unsafe.Add(mBase, uint32(v10720)+8))
	F_extract_actual_join_clauses(m, v10713, v10721, v51+int32(88), v51+int32(172))
	mBase = m.M
	v10727 = m.ExcPending
	if v10727 != 0 {
		goto L1
	} else {
		goto L1410
	}
L1408:
	;
	goto L1409
L1409:
	;
	v10729 = F_extract_actual_clauses(m, v10713, int32(0))
	mBase = m.M
	v10730 = m.ExcPending
	if v10730 != 0 {
		goto L1
	} else {
		goto L1411
	}
L1410:
	;
	goto L1406
L1411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+172)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v10729
	goto L1406
L1412:
	;
	if v10933 != 0 {
		goto L1453
	} else {
		goto L1454
	}
L1413:
	;
	v10739 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v10740 = F_replace_nestloop_params_mutator(m, v10739, l0)
	mBase = m.M
	v10741 = m.ExcPending
	if v10741 != 0 {
		goto L1
	} else {
		goto L1414
	}
L1414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v10740
	v10743 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v10744 = F_replace_nestloop_params_mutator(m, v10743, l0)
	mBase = m.M
	v10745 = m.ExcPending
	if v10745 != 0 {
		goto L1
	} else {
		goto L1415
	}
L1415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+172)) = v10744
	v10747 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10747 == int32(0) {
		v10933 = v10735
		goto L1412
	} else {
		goto L1416
	}
L1416:
	;
	v10750 = *(*int32)(unsafe.Add(mBase, uint32(v10747)+4))
	if v10750 == int32(0) {
		v10933 = v10735
		goto L1412
	} else {
		goto L1417
	}
L1417:
	;
	if v10750 == int32(0) {
		goto L1420
	} else {
		goto L1421
	}
L1418:
	;
	if v10809 < int32(0) {
		goto L1429
	} else {
		goto L1430
	}
L1419:
	;
	v10809 = base.I32_ctz(v10795) | v10796<<(uint(int32(5))%32)
	goto L1418
L1420:
	;
	v10809 = int32(-2)
	goto L1418
L1421:
	;
	v10760 = int32(0)
	v10763 = *(*int32)(unsafe.Add(mBase, uint32(v10750)+4))
	if v10763 <= v10760 {
		goto L1420
	} else {
		goto L1422
	}
L1422:
	;
	v10766 = v10750 + int32(8)
	v10770 = *(*int32)(unsafe.Add(mBase, uint32(v10766)))
	v10773 = v10770 & int32(-1)
	if v10773 != 0 {
		v10795 = v10773
		v10796 = v10760
		goto L1419
	} else {
		goto L1423
	}
L1423:
	;
	v10774 = int32(1)
	if v10774 == v10763 {
		goto L1420
	} else {
		goto L1424
	}
L1424:
	;
	v10778 = v10774
	goto L1425
L1425:
	;
	v10785 = *(*int32)(unsafe.Add(mBase, uint32(v10766+v10778<<(uint(int32(2))%32))))
	if v10785 != 0 {
		v10795 = v10785
		v10796 = v10778
		goto L1419
	} else {
		goto L1427
	}
L1426:
	;
	goto L1420
L1427:
	;
	v10787 = v10778 + int32(1)
	if v10787 != v10763 {
		v10778 = v10787
		goto L1425
	} else {
		goto L1428
	}
L1428:
	;
	goto L1426
L1429:
	;
	v10933 = v10750
	goto L1412
L1430:
	;
	goto L1431
L1431:
	;
	v10815 = v10809
	v10816 = v10750
	goto L1432
L1432:
	;
	v10860 = F_find_base_rel_ignore_join(m, l0, v10815)
	mBase = m.M
	v10861 = m.ExcPending
	if v10861 != 0 {
		goto L1
	} else {
		goto L1435
	}
L1433:
	;
	v10933 = v10869
	goto L1412
L1434:
	;
	if v10869 == int32(0) {
		goto L1441
	} else {
		goto L1442
	}
L1435:
	;
	if v10860 == int32(0) {
		v10869 = v10816
		goto L1434
	} else {
		goto L1436
	}
L1436:
	;
	v10864 = *(*int32)(unsafe.Add(mBase, uint32(v10860)+252))
	if v10864 == int32(0) {
		v10869 = v10816
		goto L1434
	} else {
		goto L1437
	}
L1437:
	;
	v10867 = F_bms_union(m, v10816, v10864)
	mBase = m.M
	v10868 = m.ExcPending
	if v10868 != 0 {
		goto L1
	} else {
		goto L1438
	}
L1438:
	;
	v10869 = v10867
	goto L1434
L1439:
	;
	if int32(0) <= v10926 {
		v10815 = v10926
		v10816 = v10869
		goto L1432
	} else {
		goto L1450
	}
L1440:
	;
	v10926 = base.I32_ctz(v10912) | v10913<<(uint(int32(5))%32)
	goto L1439
L1441:
	;
	v10926 = int32(-2)
	goto L1439
L1442:
	;
	v10877 = v10815 + int32(1)
	v10879 = int32(base.Ui32(v10877) >> (uint(int32(5)) % 32))
	v10880 = *(*int32)(unsafe.Add(mBase, uint32(v10869)+4))
	if v10880 <= v10879 {
		goto L1441
	} else {
		goto L1443
	}
L1443:
	;
	v10883 = v10869 + int32(8)
	v10887 = *(*int32)(unsafe.Add(mBase, uint32(v10883+v10879<<(uint(int32(2))%32))))
	v10890 = v10887 & (int32(-1) << (uint(v10877) % 32))
	if v10890 != 0 {
		v10912 = v10890
		v10913 = v10879
		goto L1440
	} else {
		goto L1444
	}
L1444:
	;
	v10892 = v10879 + int32(1)
	if v10892 == v10880 {
		goto L1441
	} else {
		goto L1445
	}
L1445:
	;
	v10895 = v10892
	goto L1446
L1446:
	;
	v10902 = *(*int32)(unsafe.Add(mBase, uint32(v10883+v10895<<(uint(int32(2))%32))))
	if v10902 != 0 {
		v10912 = v10902
		v10913 = v10895
		goto L1440
	} else {
		goto L1448
	}
L1447:
	;
	goto L1441
L1448:
	;
	v10904 = v10895 + int32(1)
	if v10904 != v10880 {
		v10895 = v10904
		goto L1446
	} else {
		goto L1449
	}
L1449:
	;
	goto L1447
L1450:
	;
	goto L1433
L1451:
	;
	v11873 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v11874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v11875 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v11876 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v11878 = F_palloc0(m, int32(96))
	mBase = m.M
	v11879 = m.ExcPending
	if v11879 != 0 {
		goto L1
	} else {
		goto L1578
	}
L1452:
	;
	if v11621 == int32(0) {
		v11836 = v10692
		goto L1451
	} else {
		goto L1545
	}
L1453:
	;
	v10977 = F_bms_union(m, v10700, v10933)
	mBase = m.M
	v10978 = m.ExcPending
	if v10978 != 0 {
		goto L1
	} else {
		goto L1456
	}
L1454:
	;
	v10979 = v10700
	goto L1455
L1455:
	;
	v10980 = int32(0)
	v10981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v10981 == v10980 {
		v11621 = v10980
		goto L1452
	} else {
		goto L1457
	}
L1456:
	;
	v10979 = v10977
	goto L1455
L1457:
	;
	v10989 = int32(0)
	v10992 = v4
	v10994 = v10981
	goto L1458
L1458:
	;
	v11033 = *(*int32)(unsafe.Add(mBase, uint32(v10994)+4))
	if v10989 < v11033 {
		goto L1460
	} else {
		goto L1461
	}
L1459:
	;
	v11621 = v11532
	goto L1452
L1460:
	;
	v11035 = *(*int32)(unsafe.Add(mBase, uint32(v10994)+12))
	v11039 = *(*int32)(unsafe.Add(mBase, uint32(v11035+v10989<<(uint(int32(2))%32))))
	v11040 = *(*int32)(unsafe.Add(mBase, uint32(v11039)+8))
	v11041 = *(*int32)(unsafe.Add(mBase, uint32(v11040)))
	if v11041 == int32(6) {
		goto L1466
	} else {
		goto L1467
	}
L1461:
	;
	v11532 = v10992
	goto L1462
L1462:
	;
	goto L1459
L1463:
	;
	if v11484 != 0 {
		v10989 = v11479 + int32(1)
		v10992 = v11482
		v10994 = v11484
		goto L1458
	} else {
		goto L1544
	}
L1464:
	;
	v11473 = F_lappend(m, v10992, v11039)
	mBase = m.M
	v11474 = m.ExcPending
	if v11474 != 0 {
		goto L1
	} else {
		goto L1543
	}
L1465:
	;
	v11409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11410 = *(*int32)(unsafe.Add(mBase, uint32(v11047)+4))
	v11414 = *(*int32)(unsafe.Add(mBase, uint32(v11409+v11410<<(uint(int32(2))%32))))
	v11415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v11416 = F_list_delete_nth_cell(m, v11415, v10989)
	mBase = m.M
	v11417 = m.ExcPending
	if v11417 != 0 {
		goto L1
	} else {
		goto L1541
	}
L1466:
	;
	v11044 = *(*int32)(unsafe.Add(mBase, uint32(v11040)+4))
	v11045 = F_bms_is_member(m, v11044, v10700)
	mBase = m.M
	v11046 = m.ExcPending
	if v11046 != 0 {
		goto L1
	} else {
		goto L1469
	}
L1467:
	;
	v11049 = v11040
	v11050 = v11041
	goto L1468
L1468:
	;
	if v11050 != int32(321) {
		v11479 = v10989
		v11482 = v10992
		v11484 = v10994
		goto L1463
	} else {
		goto L1471
	}
L1469:
	;
	v11047 = *(*int32)(unsafe.Add(mBase, uint32(v11039)+8))
	if v11045 != 0 {
		goto L1465
	} else {
		goto L1470
	}
L1470:
	;
	v11048 = *(*int32)(unsafe.Add(mBase, uint32(v11047)))
	v11049 = v11047
	v11050 = v11048
	goto L1468
L1471:
	;
	v11053 = F_find_placeholder_info(m, l0, v11049)
	mBase = m.M
	v11054 = m.ExcPending
	if v11054 != 0 {
		goto L1
	} else {
		goto L1472
	}
L1472:
	;
	v11055 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+12))
	v11056 = int32(0)
	if v11055 == v11056 {
		goto L1474
	} else {
		goto L1475
	}
L1473:
	;
	if v11109 == int32(0) {
		v11479 = v10989
		v11482 = v10992
		v11484 = v10994
		goto L1463
	} else {
		goto L1487
	}
L1474:
	;
	v11109 = int32(1)
	goto L1473
L1475:
	;
	goto L1476
L1476:
	;
	if v10979 == int32(0) {
		v11102 = v11056
		goto L1477
	} else {
		goto L1478
	}
L1477:
	;
	v11109 = v11102
	goto L1473
L1478:
	;
	v11065 = *(*int32)(unsafe.Add(mBase, uint32(v11055)+4))
	v11066 = *(*int32)(unsafe.Add(mBase, uint32(v10979)+4))
	if v11066 < v11065 {
		v11102 = v11056
		goto L1477
	} else {
		goto L1479
	}
L1479:
	;
	v11068 = int32(1)
	if v11065 <= v11068 {
		goto L1480
	} else {
		goto L1481
	}
L1480:
	;
	v11071 = v11068
	goto L1482
L1481:
	;
	v11071 = v11065
	goto L1482
L1482:
	;
	v11072 = int32(8)
	v11077 = int32(0)
	goto L1483
L1483:
	;
	v11084 = v11077 << (uint(int32(2)) % 32)
	v11086 = *(*int32)(unsafe.Add(mBase, uint32(v11055+v11072+v11084)))
	v11088 = *(*int32)(unsafe.Add(mBase, uint32(v10979+v11072+v11084)))
	v11091 = v11086 & (v11088 ^ int32(-1))
	v11093 = base.B2i32(v11091 == int32(0))
	if v11091 != 0 {
		v11102 = v11093
		goto L1477
	} else {
		goto L1485
	}
L1484:
	;
	v11102 = v11093
	goto L1477
L1485:
	;
	v11095 = v11077 + int32(1)
	if v11095 != v11071 {
		v11077 = v11095
		goto L1483
	} else {
		goto L1486
	}
L1486:
	;
	goto L1484
L1487:
	;
	v11112 = int32(0)
	if base.B2i32(v11055 == v11112)|base.B2i32(v10700 == v11112) != 0 {
		v11157 = v11112
		goto L1489
	} else {
		goto L1490
	}
L1488:
	;
	if v11157 == int32(0) {
		v11479 = v10989
		v11482 = v10992
		v11484 = v10994
		goto L1463
	} else {
		goto L1501
	}
L1489:
	;
	goto L1488
L1490:
	;
	v11122 = *(*int32)(unsafe.Add(mBase, uint32(v11055)+4))
	v11123 = *(*int32)(unsafe.Add(mBase, uint32(v10700)+4))
	if v11122 < v11123 {
		goto L1491
	} else {
		goto L1492
	}
L1491:
	;
	v11125 = v11122
	goto L1493
L1492:
	;
	v11125 = v11123
	goto L1493
L1493:
	;
	if v11125 <= int32(1) {
		goto L1494
	} else {
		goto L1495
	}
L1494:
	;
	v11128 = int32(1)
	goto L1496
L1495:
	;
	v11128 = v11125
	goto L1496
L1496:
	;
	v11129 = int32(8)
	v11134 = int32(0)
	goto L1497
L1497:
	;
	v11141 = v11134 << (uint(int32(2)) % 32)
	v11143 = *(*int32)(unsafe.Add(mBase, uint32(v10700+v11129+v11141)))
	v11145 = *(*int32)(unsafe.Add(mBase, uint32(v11055+v11129+v11141)))
	v11146 = v11143 & v11145
	v11148 = base.B2i32(v11146 != int32(0))
	if v11146 != 0 {
		v11157 = v11148
		goto L1489
	} else {
		goto L1499
	}
L1498:
	;
	v11157 = v11148
	goto L1489
L1499:
	;
	v11150 = v11134 + int32(1)
	if v11150 != v11128 {
		v11134 = v11150
		goto L1497
	} else {
		goto L1500
	}
L1500:
	;
	goto L1498
L1501:
	;
	v11160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v11161 = F_list_delete_nth_cell(m, v11160, v10989)
	mBase = m.M
	v11162 = m.ExcPending
	if v11162 != 0 {
		goto L1
	} else {
		goto L1502
	}
L1502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v11161
	v11164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11164)+39)))
	if v11165 == int32(1) {
		goto L1503
	} else {
		goto L1504
	}
L1503:
	;
	v11168 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+8))
	v11169 = F_copyObjectImpl(m, v11168)
	mBase = m.M
	v11170 = m.ExcPending
	if v11170 != 0 {
		goto L1
	} else {
		goto L1506
	}
L1504:
	;
	v11172 = v11049
	goto L1505
L1505:
	;
	v11173 = int32(0)
	v11174 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+12))
	if v11174 == v11173 {
		goto L1509
	} else {
		goto L1510
	}
L1506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11039)+8)) = v11169
	v11172 = v11169
	goto L1505
L1507:
	;
	if int32(0) < v11231 {
		goto L1518
	} else {
		goto L1519
	}
L1508:
	;
	v11231 = base.I32_ctz(v11217) | v11218<<(uint(int32(5))%32)
	goto L1507
L1509:
	;
	v11231 = int32(-2)
	goto L1507
L1510:
	;
	v11182 = int32(0)
	v11185 = *(*int32)(unsafe.Add(mBase, uint32(v11174)+4))
	if v11185 <= v11182 {
		goto L1509
	} else {
		goto L1511
	}
L1511:
	;
	v11188 = v11174 + int32(8)
	v11192 = *(*int32)(unsafe.Add(mBase, uint32(v11188)))
	v11195 = v11192 & int32(-1)
	if v11195 != 0 {
		v11217 = v11195
		v11218 = v11182
		goto L1508
	} else {
		goto L1512
	}
L1512:
	;
	v11196 = int32(1)
	if v11196 == v11185 {
		goto L1509
	} else {
		goto L1513
	}
L1513:
	;
	v11200 = v11196
	goto L1514
L1514:
	;
	v11207 = *(*int32)(unsafe.Add(mBase, uint32(v11188+v11200<<(uint(int32(2))%32))))
	if v11207 != 0 {
		v11217 = v11207
		v11218 = v11200
		goto L1508
	} else {
		goto L1516
	}
L1515:
	;
	goto L1509
L1516:
	;
	v11209 = v11200 + int32(1)
	if v11209 != v11185 {
		v11200 = v11209
		goto L1514
	} else {
		goto L1517
	}
L1517:
	;
	goto L1515
L1518:
	;
	v11237 = v11173
	v11239 = v11231
	goto L1521
L1519:
	;
	v11358 = v11173
	goto L1520
L1520:
	;
	v11403 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+12))
	v11404 = F_bms_del_members(m, v11358, v11403)
	mBase = m.M
	v11405 = m.ExcPending
	if v11405 != 0 {
		goto L1
	} else {
		goto L1539
	}
L1521:
	;
	v11282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	if v11239 == v11282 {
		v11294 = v11237
		goto L1523
	} else {
		goto L1524
	}
L1522:
	;
	v11358 = v11294
	goto L1520
L1523:
	;
	v11296 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+12))
	if v11296 == int32(0) {
		goto L1529
	} else {
		goto L1530
	}
L1524:
	;
	v11284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11288 = *(*int32)(unsafe.Add(mBase, uint32(v11284+v11239<<(uint(int32(2))%32))))
	if v11288 == int32(0) {
		v11294 = v11237
		goto L1523
	} else {
		goto L1525
	}
L1525:
	;
	v11291 = *(*int32)(unsafe.Add(mBase, uint32(v11288)+104))
	v11292 = F_bms_add_members(m, v11237, v11291)
	mBase = m.M
	v11293 = m.ExcPending
	if v11293 != 0 {
		goto L1
	} else {
		goto L1526
	}
L1526:
	;
	v11294 = v11292
	goto L1523
L1527:
	;
	if int32(0) < v11352 {
		v11237 = v11294
		v11239 = v11352
		goto L1521
	} else {
		goto L1538
	}
L1528:
	;
	v11352 = base.I32_ctz(v11338) | v11339<<(uint(int32(5))%32)
	goto L1527
L1529:
	;
	v11352 = int32(-2)
	goto L1527
L1530:
	;
	v11303 = v11239 + int32(1)
	v11305 = int32(base.Ui32(v11303) >> (uint(int32(5)) % 32))
	v11306 = *(*int32)(unsafe.Add(mBase, uint32(v11296)+4))
	if v11306 <= v11305 {
		goto L1529
	} else {
		goto L1531
	}
L1531:
	;
	v11309 = v11296 + int32(8)
	v11313 = *(*int32)(unsafe.Add(mBase, uint32(v11309+v11305<<(uint(int32(2))%32))))
	v11316 = v11313 & (int32(-1) << (uint(v11303) % 32))
	if v11316 != 0 {
		v11338 = v11316
		v11339 = v11305
		goto L1528
	} else {
		goto L1532
	}
L1532:
	;
	v11318 = v11305 + int32(1)
	if v11318 == v11306 {
		goto L1529
	} else {
		goto L1533
	}
L1533:
	;
	v11321 = v11318
	goto L1534
L1534:
	;
	v11328 = *(*int32)(unsafe.Add(mBase, uint32(v11309+v11321<<(uint(int32(2))%32))))
	if v11328 != 0 {
		v11338 = v11328
		v11339 = v11321
		goto L1528
	} else {
		goto L1536
	}
L1535:
	;
	goto L1529
L1536:
	;
	v11330 = v11321 + int32(1)
	if v11330 != v11306 {
		v11321 = v11330
		goto L1534
	} else {
		goto L1537
	}
L1537:
	;
	goto L1535
L1538:
	;
	goto L1522
L1539:
	;
	v11406 = F_bms_intersect(m, v11404, v10700)
	mBase = m.M
	v11407 = m.ExcPending
	if v11407 != 0 {
		goto L1
	} else {
		goto L1540
	}
L1540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11172)+12)) = v11406
	v11432 = v11161
	goto L1464
L1541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v11416
	v11419 = *(*int32)(unsafe.Add(mBase, uint32(v11414)+104))
	v11420 = F_bms_intersect(m, v11419, v10700)
	mBase = m.M
	v11421 = m.ExcPending
	if v11421 != 0 {
		goto L1
	} else {
		goto L1542
	}
L1542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11047)+24)) = v11420
	v11432 = v11416
	goto L1464
L1543:
	;
	v11479 = v10989 - int32(1)
	v11482 = v11473
	v11484 = v11432
	goto L1463
L1544:
	;
	v11532 = v11482
	goto L1462
L1545:
	;
	v11624 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+44))
	v11625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10692)+37)))
	v11626 = *(*int32)(unsafe.Add(mBase, uint32(v11621)+4))
	if int32(0) < v11626 {
		goto L1546
	} else {
		goto L1547
	}
L1546:
	;
	v11633 = int32(0)
	v11636 = v11624
	v11640 = v11625
	goto L1549
L1547:
	;
	v11732 = v11624
	v11736 = v11625
	goto L1548
L1548:
	;
	v11774 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+44))
	if v11732 == v11774 {
		v11836 = v10692
		goto L1451
	} else {
		goto L1569
	}
L1549:
	;
	v11678 = *(*int32)(unsafe.Add(mBase, uint32(v11621)+12))
	v11682 = *(*int32)(unsafe.Add(mBase, uint32(v11678+v11633<<(uint(int32(2))%32))))
	v11683 = *(*int32)(unsafe.Add(mBase, uint32(v11682)+8))
	v11684 = *(*int32)(unsafe.Add(mBase, uint32(v11683)))
	if v11684 == int32(6) {
		v11720 = v11636
		v11721 = v11640
		goto L1551
	} else {
		goto L1552
	}
L1550:
	;
	v11732 = v11720
	v11736 = v11721
	goto L1548
L1551:
	;
	v11723 = v11633 + int32(1)
	v11724 = *(*int32)(unsafe.Add(mBase, uint32(v11621)+4))
	if v11723 < v11724 {
		v11633 = v11723
		v11636 = v11720
		v11640 = v11721
		goto L1549
	} else {
		goto L1568
	}
L1552:
	;
	v11687 = F_tlist_member(m, v11683, v11636)
	mBase = m.M
	v11688 = m.ExcPending
	if v11688 != 0 {
		goto L1
	} else {
		goto L1553
	}
L1553:
	;
	if v11687 != 0 {
		v11720 = v11636
		v11721 = v11640
		goto L1551
	} else {
		goto L1554
	}
L1554:
	;
	v11689 = *(*int32)(unsafe.Add(mBase, uint32(v11683)+4))
	v11690 = F_replace_nestloop_params_mutator(m, v11689, l0)
	mBase = m.M
	v11691 = m.ExcPending
	if v11691 != 0 {
		goto L1
	} else {
		goto L1555
	}
L1555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11683)+4)) = v11690
	v11693 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+44))
	if v11693 == v11636 {
		goto L1556
	} else {
		goto L1557
	}
L1556:
	;
	v11695 = F_list_copy(m, v11636)
	mBase = m.M
	v11696 = m.ExcPending
	if v11696 != 0 {
		goto L1
	} else {
		goto L1559
	}
L1557:
	;
	v11697 = v11636
	goto L1558
L1558:
	;
	v11698 = F_copyObjectImpl(m, v11683)
	mBase = m.M
	v11699 = m.ExcPending
	if v11699 != 0 {
		goto L1
	} else {
		goto L1560
	}
L1559:
	;
	v11697 = v11695
	goto L1558
L1560:
	;
	if v11697 != 0 {
		goto L1561
	} else {
		goto L1562
	}
L1561:
	;
	v11703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11697)+4)))
	v11707 = v11703 + int32(1)
	goto L1563
L1562:
	;
	v11707 = int32(1)
	goto L1563
L1563:
	;
	v11711 = F_makeTargetEntry(m, v11698, base.I32_extend16_s(v11707), int32(0), int32(1))
	mBase = m.M
	v11712 = m.ExcPending
	if v11712 != 0 {
		goto L1
	} else {
		goto L1564
	}
L1564:
	;
	v11713 = F_lappend(m, v11697, v11711)
	mBase = m.M
	v11714 = m.ExcPending
	if v11714 != 0 {
		goto L1
	} else {
		goto L1565
	}
L1565:
	;
	if v11640&int32(1) == int32(0) {
		v11720 = v11713
		v11721 = int32(0)
		goto L1551
	} else {
		goto L1566
	}
L1566:
	;
	v11717 = F_is_parallel_safe(m, l0, v11683)
	mBase = m.M
	v11718 = m.ExcPending
	if v11718 != 0 {
		goto L1
	} else {
		goto L1567
	}
L1567:
	;
	v11720 = v11713
	v11721 = v11717
	goto L1551
L1568:
	;
	goto L1550
L1569:
	;
	v11777 = v11736 & int32(1)
	v11778 = *(*int32)(unsafe.Add(mBase, uint32(v10692)))
	switch v11778 - int32(336) {
	case 0, 1, 2, 3, 4, 28, 29, 30, 35, 38, 39, 40, 41:
		goto L1572
	default:
		goto L1571
	case 23:
		goto L1573
	}
L1570:
	;
	v11836 = v11824
	goto L1451
L1571:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10692)+44)) = v11732
	v11818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10692)+37)))
	v11819 = v11777 & v11818
	*(*uint8)(unsafe.Add(mBase, uint32(v10692)+37)) = uint8(v11819)
	v11824 = v10692
	goto L1570
L1572:
	;
	v11784 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+44))
	v11785 = F_tlist_same_exprs(m, v11732, v11784)
	mBase = m.M
	v11786 = m.ExcPending
	if v11786 != 0 {
		goto L1
	} else {
		goto L1575
	}
L1573:
	;
	v11781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10692)+80)))
	if v11781&int32(4) != 0 {
		goto L1571
	} else {
		goto L1574
	}
L1574:
	;
	goto L1572
L1575:
	;
	if v11785 != 0 {
		goto L1571
	} else {
		goto L1576
	}
L1576:
	;
	v11787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10692)+37)))
	v11789 = F_palloc0(m, int32(88))
	mBase = m.M
	v11790 = m.ExcPending
	if v11790 != 0 {
		goto L1
	} else {
		goto L1577
	}
L1577:
	;
	v11791 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+80)) = v11791
	*(*int64)(unsafe.Add(mBase, uint32(v11789)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+56)) = v11791
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+52)) = v10692
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+48)) = v11791
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+44)) = v11732
	*(*int32)(unsafe.Add(mBase, uint32(v11789))) = int32(335)
	v11803 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+4)) = v11803
	v11805 = *(*float64)(unsafe.Add(mBase, uint32(v10692)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11789)+8)) = v11805
	v11807 = *(*float64)(unsafe.Add(mBase, uint32(v10692)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v11789)+16)) = v11807
	v11809 = *(*float64)(unsafe.Add(mBase, uint32(v10692)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v11789)+24)) = v11809
	v11811 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+32))
	v11812 = v11777 & v11787
	*(*uint8)(unsafe.Add(mBase, uint32(v11789)+37)) = uint8(v11812)
	*(*uint8)(unsafe.Add(mBase, uint32(v11789)+36)) = uint8(v11791)
	*(*int32)(unsafe.Add(mBase, uint32(v11789)+32)) = v11811
	v11824 = v11789
	goto L1570
L1578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+88)) = v11621
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+80)) = v11876
	*(*uint8)(unsafe.Add(mBase, uint32(v11878)+76)) = uint8(v11874)
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+72)) = v11873
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+56)) = v10707
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+52)) = v11836
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+48)) = v11875
	*(*int32)(unsafe.Add(mBase, uint32(v11878)+44)) = v10636
	*(*int32)(unsafe.Add(mBase, uint32(v11878))) = int32(360)
	v12316 = v11878
	goto L4
L1579:
	;
	v12025 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12026 = int32(2)
	v12028 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v12026 <= v12028 {
		goto L1594
	} else {
		goto L1595
	}
L1580:
	;
	v11896 = *(*int32)(unsafe.Add(mBase, uint32(v11892)+4))
	if v11896 <= int32(0) {
		v11979 = v11890
		goto L1579
	} else {
		goto L1581
	}
L1581:
	;
	v11899 = *(*int32)(unsafe.Add(mBase, uint32(v11891)+8))
	v11902 = v11890
	v11903 = int32(1)
	v11905 = v4
	goto L1582
L1582:
	;
	v11948 = *(*int32)(unsafe.Add(mBase, uint32(v11892)+12))
	v11952 = *(*int32)(unsafe.Add(mBase, uint32(v11948+v11905<<(uint(int32(2))%32))))
	v11953 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11953 != 0 {
		goto L1584
	} else {
		goto L1585
	}
L1583:
	;
	v11979 = v11971
	goto L1579
L1584:
	;
	v11954 = F_replace_nestloop_params_mutator(m, v11952, l0)
	mBase = m.M
	v11955 = m.ExcPending
	if v11955 != 0 {
		goto L1
	} else {
		goto L1587
	}
L1585:
	;
	v11956 = v11952
	goto L1586
L1586:
	;
	v11958 = int32(0)
	v11960 = F_makeTargetEntry(m, v11956, base.I32_extend16_s(v11903), v11958, v11958)
	mBase = m.M
	v11961 = m.ExcPending
	if v11961 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1587:
	;
	v11956 = v11954
	goto L1586
L1588:
	;
	if v11899 != 0 {
		goto L1589
	} else {
		goto L1590
	}
L1589:
	;
	v11967 = *(*int32)(unsafe.Add(mBase, uint32(v11899+v11903<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11960)+16)) = v11967
	goto L1591
L1590:
	;
	goto L1591
L1591:
	;
	v11971 = F_lappend(m, v11902, v11960)
	mBase = m.M
	v11972 = m.ExcPending
	if v11972 != 0 {
		goto L1
	} else {
		goto L1592
	}
L1592:
	;
	v11974 = v11905 + int32(1)
	v11975 = *(*int32)(unsafe.Add(mBase, uint32(v11892)+4))
	if v11974 < v11975 {
		v11902 = v11971
		v11903 = v11903 + int32(1)
		v11905 = v11974
		goto L1582
	} else {
		goto L1593
	}
L1593:
	;
	goto L1583
L1594:
	;
	v12031 = v12026
	goto L1596
L1595:
	;
	v12031 = int32(0)
	goto L1596
L1596:
	;
	v12032 = F_create_plan_recurse(m, l0, v12025, v12031)
	mBase = m.M
	v12033 = m.ExcPending
	if v12033 != 0 {
		goto L1
	} else {
		goto L1597
	}
L1597:
	;
	v12034 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v12036 = F_create_plan_recurse(m, l0, v12034, int32(2))
	mBase = m.M
	v12037 = m.ExcPending
	if v12037 != 0 {
		goto L1
	} else {
		goto L1598
	}
L1598:
	;
	v12038 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12039 = F_order_qual_clauses(m, l0, v12038)
	mBase = m.M
	v12040 = m.ExcPending
	if v12040 != 0 {
		goto L1
	} else {
		goto L1599
	}
L1599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v12039
	v12043 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v12043)%32)&int32(174) != 0 {
		goto L1601
	} else {
		goto L1602
	}
L1600:
	;
	v12062 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12063 = F_get_actual_clauses(m, v12062)
	mBase = m.M
	v12064 = m.ExcPending
	if v12064 != 0 {
		goto L1
	} else {
		goto L1606
	}
L1601:
	;
	v12047 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12048 = *(*int32)(unsafe.Add(mBase, uint32(v12047)+8))
	F_extract_actual_join_clauses(m, v12039, v12048, v51+int32(88), v51+int32(172))
	mBase = m.M
	v12054 = m.ExcPending
	if v12054 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1602:
	;
	goto L1603
L1603:
	;
	v12056 = F_extract_actual_clauses(m, v12039, int32(0))
	mBase = m.M
	v12057 = m.ExcPending
	if v12057 != 0 {
		goto L1
	} else {
		goto L1605
	}
L1604:
	;
	goto L1600
L1605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+172)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v12056
	goto L1600
L1606:
	;
	v12065 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v12066 = F_list_difference(m, v12065, v12063)
	mBase = m.M
	v12067 = m.ExcPending
	if v12067 != 0 {
		goto L1
	} else {
		goto L1607
	}
L1607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v12066
	v12069 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12069 != 0 {
		goto L1608
	} else {
		goto L1609
	}
L1608:
	;
	v12070 = F_replace_nestloop_params_mutator(m, v12066, l0)
	mBase = m.M
	v12071 = m.ExcPending
	if v12071 != 0 {
		goto L1
	} else {
		goto L1611
	}
L1609:
	;
	goto L1610
L1610:
	;
	v12077 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12078 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12079 = *(*int32)(unsafe.Add(mBase, uint32(v12078)+8))
	v12080 = *(*int32)(unsafe.Add(mBase, uint32(v12079)+8))
	v12081 = F_get_switched_clauses(m, v12077, v12080)
	mBase = m.M
	v12082 = m.ExcPending
	if v12082 != 0 {
		goto L1
	} else {
		goto L1616
	}
L1611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v12070
	v12073 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v12074 = F_replace_nestloop_params_mutator(m, v12073, l0)
	mBase = m.M
	v12075 = m.ExcPending
	if v12075 != 0 {
		goto L1
	} else {
		goto L1612
	}
L1612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+172)) = v12074
	goto L1610
L1613:
	;
	v12254 = F_palloc0(m, int32(96))
	mBase = m.M
	v12255 = m.ExcPending
	if v12255 != 0 {
		goto L1
	} else {
		goto L1636
	}
L1614:
	;
	v12136 = int32(0)
	v12138 = v12086
	v12142 = v4
	v12143 = v4
	v12144 = v4
	goto L1629
L1615:
	;
	v12211 = v12127
	v12214 = v12128
	v12215 = v4
	v12216 = v4
	v12217 = v4
	v12218 = v12129
	v12220 = v12130
	goto L1613
L1616:
	;
	if v12081 == int32(0) {
		goto L1617
	} else {
		goto L1618
	}
L1617:
	;
	v12127 = int32(0)
	v12128 = v4
	v12129 = v4
	v12130 = v4
	goto L1615
L1618:
	;
	goto L1619
L1619:
	;
	v12086 = int32(0)
	v12087 = *(*int32)(unsafe.Add(mBase, uint32(v12081)+4))
	if v12087 != int32(1) {
		goto L1621
	} else {
		goto L1622
	}
L1620:
	;
	v12122 = *(*int32)(unsafe.Add(mBase, uint32(v12081)+4))
	if int32(0) < v12122 {
		goto L1614
	} else {
		goto L1628
	}
L1621:
	;
	v12119 = v4
	v12120 = v4
	v12121 = int32(0)
	goto L1620
L1622:
	;
	v12090 = *(*int32)(unsafe.Add(mBase, uint32(v12081)+12))
	v12091 = *(*int32)(unsafe.Add(mBase, uint32(v12090)))
	v12092 = *(*int32)(unsafe.Add(mBase, uint32(v12091)+28))
	v12093 = *(*int32)(unsafe.Add(mBase, uint32(v12092)+12))
	v12094 = *(*int32)(unsafe.Add(mBase, uint32(v12093)))
	v12095 = *(*int32)(unsafe.Add(mBase, uint32(v12094)))
	if v12095 == int32(27) {
		goto L1623
	} else {
		goto L1624
	}
L1623:
	;
	v12098 = *(*int32)(unsafe.Add(mBase, uint32(v12094)+4))
	v12099 = *(*int32)(unsafe.Add(mBase, uint32(v12098)))
	v12100 = v12098
	v12101 = v12099
	goto L1625
L1624:
	;
	v12100 = v12094
	v12101 = v12095
	goto L1625
L1625:
	;
	if v12101 != int32(6) {
		goto L1621
	} else {
		goto L1626
	}
L1626:
	;
	v12104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12105 = *(*int32)(unsafe.Add(mBase, uint32(v12100)+4))
	v12109 = *(*int32)(unsafe.Add(mBase, uint32(v12104+v12105<<(uint(int32(2))%32))))
	v12110 = *(*int32)(unsafe.Add(mBase, uint32(v12109)+12))
	if v12110 != 0 {
		goto L1621
	} else {
		goto L1627
	}
L1627:
	;
	v12111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12109)+20)))
	v12112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12100)+8)))
	v12113 = *(*int32)(unsafe.Add(mBase, uint32(v12109)+16))
	v12119 = v12111
	v12120 = v12112
	v12121 = v12113
	goto L1620
L1628:
	;
	v12127 = v12086
	v12128 = v12119
	v12129 = v12120
	v12130 = v12121
	goto L1615
L1629:
	;
	v12180 = *(*int32)(unsafe.Add(mBase, uint32(v12081)+12))
	v12184 = *(*int32)(unsafe.Add(mBase, uint32(v12180+v12136<<(uint(int32(2))%32))))
	v12185 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+4))
	v12186 = F_lappend_oid(m, v12144, v12185)
	mBase = m.M
	v12187 = m.ExcPending
	if v12187 != 0 {
		goto L1
	} else {
		goto L1631
	}
L1630:
	;
	v12211 = v12194
	v12214 = v12119
	v12215 = v12189
	v12216 = v12199
	v12217 = v12186
	v12218 = v12120
	v12220 = v12121
	goto L1613
L1631:
	;
	v12188 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+24))
	v12189 = F_lappend_oid(m, v12142, v12188)
	mBase = m.M
	v12190 = m.ExcPending
	if v12190 != 0 {
		goto L1
	} else {
		goto L1632
	}
L1632:
	;
	v12191 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+28))
	v12192 = *(*int32)(unsafe.Add(mBase, uint32(v12191)+12))
	v12193 = *(*int32)(unsafe.Add(mBase, uint32(v12192)))
	v12194 = F_lappend(m, v12138, v12193)
	mBase = m.M
	v12195 = m.ExcPending
	if v12195 != 0 {
		goto L1
	} else {
		goto L1633
	}
L1633:
	;
	v12196 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+28))
	v12197 = *(*int32)(unsafe.Add(mBase, uint32(v12196)+12))
	v12198 = *(*int32)(unsafe.Add(mBase, uint32(v12197)+4))
	v12199 = F_lappend(m, v12143, v12198)
	mBase = m.M
	v12200 = m.ExcPending
	if v12200 != 0 {
		goto L1
	} else {
		goto L1634
	}
L1634:
	;
	v12202 = v12136 + int32(1)
	v12203 = *(*int32)(unsafe.Add(mBase, uint32(v12081)+4))
	if v12202 < v12203 {
		v12136 = v12202
		v12138 = v12194
		v12142 = v12189
		v12143 = v12199
		v12144 = v12186
		goto L1629
	} else {
		goto L1635
	}
L1635:
	;
	goto L1630
L1636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12254))) = int32(374)
	v12258 = *(*int32)(unsafe.Add(mBase, uint32(v12036)+44))
	v12259 = int32(1)
	v12260 = v12214 & v12259
	*(*uint8)(unsafe.Add(mBase, uint32(v12254)+82)) = uint8(v12260)
	*(*uint16)(unsafe.Add(mBase, uint32(v12254)+80)) = uint16(v12218)
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+76)) = v12220
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+72)) = v12216
	v12265 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+56)) = v12265
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+52)) = v12036
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+48)) = v12265
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+44)) = v12258
	v12271 = *(*int32)(unsafe.Add(mBase, uint32(v12036)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+4)) = v12271
	v12273 = *(*float64)(unsafe.Add(mBase, uint32(v12036)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12254)+8)) = v12273
	v12275 = *(*float64)(unsafe.Add(mBase, uint32(v12036)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12254)+16)) = v12275
	v12277 = *(*float64)(unsafe.Add(mBase, uint32(v12036)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12254)+24)) = v12277
	v12279 = *(*int32)(unsafe.Add(mBase, uint32(v12036)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v12254)+36)) = uint8(v12265)
	*(*int32)(unsafe.Add(mBase, uint32(v12254)+32)) = v12279
	v12283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12036)+37)))
	*(*float64)(unsafe.Add(mBase, uint32(v12254)+8)) = v12275
	*(*uint8)(unsafe.Add(mBase, uint32(v12254)+37)) = uint8(v12283)
	v12286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v12286 == v12259 {
		goto L1637
	} else {
		goto L1638
	}
L1637:
	;
	v12289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12254)+36)) = uint8(v12289)
	v12291 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12254)+88)) = v12291
	goto L1639
L1638:
	;
	goto L1639
L1639:
	;
	v12293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v12294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v12295 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	v12296 = *(*int32)(unsafe.Add(mBase, uint32(v51)+172))
	v12298 = F_palloc0(m, int32(104))
	mBase = m.M
	v12299 = m.ExcPending
	if v12299 != 0 {
		goto L1
	} else {
		goto L1640
	}
L1640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+100)) = v12211
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+96)) = v12215
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+92)) = v12217
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+88)) = v12081
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+56)) = v12254
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+52)) = v12032
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+48)) = v12296
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+44)) = v11979
	*(*int32)(unsafe.Add(mBase, uint32(v12298))) = int32(363)
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+80)) = v12295
	*(*uint8)(unsafe.Add(mBase, uint32(v12298)+76)) = uint8(v12294)
	*(*int32)(unsafe.Add(mBase, uint32(v12298)+72)) = v12293
	v12316 = v12298
	goto L4
L1641:
	;
	v12379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12380 = F_order_qual_clauses(m, l0, v12379)
	mBase = m.M
	v12381 = m.ExcPending
	if v12381 != 0 {
		goto L1
	} else {
		goto L1642
	}
L1642:
	;
	v12383 = F_extract_actual_clauses(m, v12380, int32(1))
	mBase = m.M
	v12384 = m.ExcPending
	if v12384 != 0 {
		goto L1
	} else {
		goto L1643
	}
L1643:
	;
	if v12383 == int32(0) {
		v12392 = v12316
		v12397 = v51
		goto L3
	} else {
		goto L1644
	}
L1644:
	;
	v12387 = F_create_gating_plan(m, l0, l1, v12316, v12383)
	mBase = m.M
	v12388 = m.ExcPending
	if v12388 != 0 {
		goto L1
	} else {
		goto L1645
	}
L1645:
	;
	v12392 = v12387
	v12397 = v51
	goto L3
}
