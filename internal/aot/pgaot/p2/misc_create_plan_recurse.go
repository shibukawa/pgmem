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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
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
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 float64
	_ = v301
	var v302 float64
	_ = v302
	var v303 float64
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 float64
	_ = v310
	var v312 float64
	_ = v312
	var v314 float64
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 float64
	_ = v368
	var v369 float64
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 float64
	_ = v383
	var v386 int32
	_ = v386
	var v387 float64
	_ = v387
	var v393 float64
	_ = v393
	var v399 float64
	_ = v399
	var v401 float64
	_ = v401
	var v403 float64
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 float64
	_ = v473
	var v474 float64
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 float64
	_ = v488
	var v491 int32
	_ = v491
	var v492 float64
	_ = v492
	var v498 float64
	_ = v498
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v508 float64
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 float64
	_ = v547
	var v549 float64
	_ = v549
	var v551 float64
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 float64
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
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
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v800 int32
	_ = v800
	var v835 int32
	_ = v835
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1020 int32
	_ = v1020
	var v1022 float64
	_ = v1022
	var v1024 float64
	_ = v1024
	var v1026 float64
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1145 int32
	_ = v1145
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
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1230 float64
	_ = v1230
	var v1232 float64
	_ = v1232
	var v1234 float64
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1297 int32
	_ = v1297
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1406 int64
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1428 float64
	_ = v1428
	var v1430 float64
	_ = v1430
	var v1432 float64
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1523 int64
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1540 int32
	_ = v1540
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1611 int32
	_ = v1611
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1717 int32
	_ = v1717
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
	var v1728 int32
	_ = v1728
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1822 int32
	_ = v1822
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
	var v1844 int32
	_ = v1844
	var v1855 int32
	_ = v1855
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1920 int32
	_ = v1920
	var v1925 int32
	_ = v1925
	var v1935 int32
	_ = v1935
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1995 int32
	_ = v1995
	var v2019 int32
	_ = v2019
	var v2026 int32
	_ = v2026
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2119 int32
	_ = v2119
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2163 int32
	_ = v2163
	var v2178 int32
	_ = v2178
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2395 int32
	_ = v2395
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2446 int32
	_ = v2446
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2482 int32
	_ = v2482
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2498 int32
	_ = v2498
	var v2503 int32
	_ = v2503
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2622 int32
	_ = v2622
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2688 int32
	_ = v2688
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
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
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2862 int32
	_ = v2862
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2874 int32
	_ = v2874
	var v2879 int32
	_ = v2879
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2942 int32
	_ = v2942
	var v2945 int32
	_ = v2945
	var v2948 int32
	_ = v2948
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2970 int32
	_ = v2970
	var v2974 int32
	_ = v2974
	var v2979 int32
	_ = v2979
	var v2983 int32
	_ = v2983
	var v2992 int32
	_ = v2992
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3004 int32
	_ = v3004
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3021 int32
	_ = v3021
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3036 int32
	_ = v3036
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3073 int32
	_ = v3073
	var v3082 int32
	_ = v3082
	var v3111 int32
	_ = v3111
	var v3113 float64
	_ = v3113
	var v3115 float64
	_ = v3115
	var v3117 float64
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3122 int32
	_ = v3122
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int64
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3137 int32
	_ = v3137
	var v3143 int32
	_ = v3143
	var v3145 float64
	_ = v3145
	var v3147 float64
	_ = v3147
	var v3149 float64
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3175 int32
	_ = v3175
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3222 int32
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3293 float64
	_ = v3293
	var v3303 float64
	_ = v3303
	var v3306 float64
	_ = v3306
	var v3310 int32
	_ = v3310
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3327 int32
	_ = v3327
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3348 int32
	_ = v3348
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3406 int32
	_ = v3406
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3509 int32
	_ = v3509
	var v3511 float64
	_ = v3511
	var v3513 float64
	_ = v3513
	var v3515 float64
	_ = v3515
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3538 int32
	_ = v3538
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3579 int32
	_ = v3579
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3589 int32
	_ = v3589
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3596 int32
	_ = v3596
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3612 int32
	_ = v3612
	var v3650 int32
	_ = v3650
	var v3652 int32
	_ = v3652
	var v3653 int32
	_ = v3653
	var v3654 int32
	_ = v3654
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3658 float64
	_ = v3658
	var v3668 float64
	_ = v3668
	var v3671 float64
	_ = v3671
	var v3675 int32
	_ = v3675
	var v3678 int32
	_ = v3678
	var v3679 int32
	_ = v3679
	var v3680 int32
	_ = v3680
	var v3681 int32
	_ = v3681
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3708 int32
	_ = v3708
	var v3715 int32
	_ = v3715
	var v3717 int32
	_ = v3717
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3784 int32
	_ = v3784
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3842 int32
	_ = v3842
	var v3844 float64
	_ = v3844
	var v3846 float64
	_ = v3846
	var v3848 float64
	_ = v3848
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3877 int32
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3886 int32
	_ = v3886
	var v3924 int32
	_ = v3924
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3930 int32
	_ = v3930
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3934 int32
	_ = v3934
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3941 int32
	_ = v3941
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3948 int32
	_ = v3948
	var v3949 int32
	_ = v3949
	var v3957 int32
	_ = v3957
	var v3997 int32
	_ = v3997
	var v3998 int32
	_ = v3998
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4005 int32
	_ = v4005
	var v4008 int32
	_ = v4008
	var v4015 int32
	_ = v4015
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4064 int32
	_ = v4064
	var v4065 int32
	_ = v4065
	var v4066 int32
	_ = v4066
	var v4069 int32
	_ = v4069
	var v4072 int32
	_ = v4072
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4084 int32
	_ = v4084
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4129 int32
	_ = v4129
	var v4130 int32
	_ = v4130
	var v4131 int32
	_ = v4131
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4135 int32
	_ = v4135
	var v4138 int32
	_ = v4138
	var v4141 int32
	_ = v4141
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4199 int32
	_ = v4199
	var v4202 int32
	_ = v4202
	var v4203 int32
	_ = v4203
	var v4204 int32
	_ = v4204
	var v4207 int32
	_ = v4207
	var v4208 int32
	_ = v4208
	var v4210 int32
	_ = v4210
	var v4254 int32
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4256 int32
	_ = v4256
	var v4258 int32
	_ = v4258
	var v4259 int32
	_ = v4259
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4274 int32
	_ = v4274
	var v4276 int32
	_ = v4276
	var v4278 int32
	_ = v4278
	var v4282 int32
	_ = v4282
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4288 int32
	_ = v4288
	var v4290 int32
	_ = v4290
	var v4298 int32
	_ = v4298
	var v4300 float64
	_ = v4300
	var v4302 float64
	_ = v4302
	var v4304 float64
	_ = v4304
	var v4306 int32
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4313 int32
	_ = v4313
	var v4316 int32
	_ = v4316
	var v4317 int32
	_ = v4317
	var v4319 int32
	_ = v4319
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4322 int32
	_ = v4322
	var v4325 int32
	_ = v4325
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4339 int32
	_ = v4339
	var v4343 int32
	_ = v4343
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4387 int32
	_ = v4387
	var v4388 int32
	_ = v4388
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4391 int32
	_ = v4391
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4397 int32
	_ = v4397
	var v4399 int32
	_ = v4399
	var v4401 int32
	_ = v4401
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4407 int32
	_ = v4407
	var v4411 int32
	_ = v4411
	var v4414 int32
	_ = v4414
	var v4453 int32
	_ = v4453
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4460 int32
	_ = v4460
	var v4502 int32
	_ = v4502
	var v4503 int32
	_ = v4503
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4509 int32
	_ = v4509
	var v4513 int32
	_ = v4513
	var v4561 int32
	_ = v4561
	var v4603 int32
	_ = v4603
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4608 int32
	_ = v4608
	var v4609 int32
	_ = v4609
	var v4614 int32
	_ = v4614
	var v4656 int32
	_ = v4656
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4662 int32
	_ = v4662
	var v4663 int32
	_ = v4663
	var v4664 int32
	_ = v4664
	var v4665 int32
	_ = v4665
	var v4668 int32
	_ = v4668
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4721 int32
	_ = v4721
	var v4724 int32
	_ = v4724
	var v4725 int32
	_ = v4725
	var v4726 int32
	_ = v4726
	var v4744 int32
	_ = v4744
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4772 int32
	_ = v4772
	var v4776 int32
	_ = v4776
	var v4777 int32
	_ = v4777
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4783 int32
	_ = v4783
	var v4784 int32
	_ = v4784
	var v4787 int32
	_ = v4787
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4790 int32
	_ = v4790
	var v4795 int32
	_ = v4795
	var v4837 int32
	_ = v4837
	var v4840 int32
	_ = v4840
	var v4844 int32
	_ = v4844
	var v4845 int32
	_ = v4845
	var v4849 int32
	_ = v4849
	var v4852 int32
	_ = v4852
	var v4853 int32
	_ = v4853
	var v4858 int32
	_ = v4858
	var v4899 int32
	_ = v4899
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4909 int32
	_ = v4909
	var v4910 int32
	_ = v4910
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4918 int32
	_ = v4918
	var v4919 int32
	_ = v4919
	var v4920 int32
	_ = v4920
	var v4921 int32
	_ = v4921
	var v4922 int32
	_ = v4922
	var v4925 int32
	_ = v4925
	var v4930 int32
	_ = v4930
	var v4973 int32
	_ = v4973
	var v4974 int32
	_ = v4974
	var v4976 int32
	_ = v4976
	var v4978 int32
	_ = v4978
	var v4980 int32
	_ = v4980
	var v4984 int32
	_ = v4984
	var v4987 int32
	_ = v4987
	var v4990 int32
	_ = v4990
	var v4991 int32
	_ = v4991
	var v4995 int32
	_ = v4995
	var v5003 int32
	_ = v5003
	var v5004 int32
	_ = v5004
	var v5007 int32
	_ = v5007
	var v5018 int32
	_ = v5018
	var v5023 int32
	_ = v5023
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5030 int32
	_ = v5030
	var v5031 int32
	_ = v5031
	var v5034 int32
	_ = v5034
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5042 int32
	_ = v5042
	var v5085 int32
	_ = v5085
	var v5086 int32
	_ = v5086
	var v5089 int32
	_ = v5089
	var v5091 int32
	_ = v5091
	var v5093 int32
	_ = v5093
	var v5099 int32
	_ = v5099
	var v5104 int32
	_ = v5104
	var v5108 int32
	_ = v5108
	var v5111 int32
	_ = v5111
	var v5121 int32
	_ = v5121
	var v5160 int32
	_ = v5160
	var v5161 int32
	_ = v5161
	var v5162 int32
	_ = v5162
	var v5167 int32
	_ = v5167
	var v5171 int32
	_ = v5171
	var v5176 int32
	_ = v5176
	var v5182 int32
	_ = v5182
	var v5221 int32
	_ = v5221
	var v5222 int32
	_ = v5222
	var v5223 int32
	_ = v5223
	var v5224 int32
	_ = v5224
	var v5229 int32
	_ = v5229
	var v5232 int32
	_ = v5232
	var v5245 int32
	_ = v5245
	var v5271 int32
	_ = v5271
	var v5272 int32
	_ = v5272
	var v5274 int32
	_ = v5274
	var v5275 int32
	_ = v5275
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5278 int32
	_ = v5278
	var v5279 int32
	_ = v5279
	var v5280 int32
	_ = v5280
	var v5281 int32
	_ = v5281
	var v5282 int64
	_ = v5282
	var v5283 int32
	_ = v5283
	var v5284 float64
	_ = v5284
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5299 float64
	_ = v5299
	var v5302 float64
	_ = v5302
	var v5306 int32
	_ = v5306
	var v5309 int32
	_ = v5309
	var v5310 int32
	_ = v5310
	var v5329 int32
	_ = v5329
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5336 int32
	_ = v5336
	var v5337 int32
	_ = v5337
	var v5339 int32
	_ = v5339
	var v5345 int32
	_ = v5345
	var v5357 int32
	_ = v5357
	var v5384 int32
	_ = v5384
	var v5385 int32
	_ = v5385
	var v5386 int32
	_ = v5386
	var v5390 int32
	_ = v5390
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5395 int32
	_ = v5395
	var v5396 int32
	_ = v5396
	var v5397 int32
	_ = v5397
	var v5403 int32
	_ = v5403
	var v5445 int32
	_ = v5445
	var v5448 int32
	_ = v5448
	var v5452 int32
	_ = v5452
	var v5453 int32
	_ = v5453
	var v5457 int32
	_ = v5457
	var v5460 int32
	_ = v5460
	var v5461 int32
	_ = v5461
	var v5473 int32
	_ = v5473
	var v5507 int32
	_ = v5507
	var v5509 int32
	_ = v5509
	var v5510 int32
	_ = v5510
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
	var v5513 int32
	_ = v5513
	var v5514 int32
	_ = v5514
	var v5515 int32
	_ = v5515
	var v5519 int32
	_ = v5519
	var v5520 int32
	_ = v5520
	var v5523 int32
	_ = v5523
	var v5528 int32
	_ = v5528
	var v5529 int32
	_ = v5529
	var v5530 int32
	_ = v5530
	var v5570 int32
	_ = v5570
	var v5574 int32
	_ = v5574
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5577 int32
	_ = v5577
	var v5578 int32
	_ = v5578
	var v5580 int32
	_ = v5580
	var v5582 int32
	_ = v5582
	var v5583 int32
	_ = v5583
	var v5587 int32
	_ = v5587
	var v5591 int32
	_ = v5591
	var v5592 int32
	_ = v5592
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5601 int32
	_ = v5601
	var v5641 int32
	_ = v5641
	var v5642 int32
	_ = v5642
	var v5643 int32
	_ = v5643
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5646 int32
	_ = v5646
	var v5647 int32
	_ = v5647
	var v5648 int32
	_ = v5648
	var v5649 int32
	_ = v5649
	var v5650 int64
	_ = v5650
	var v5651 int32
	_ = v5651
	var v5652 float64
	_ = v5652
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5667 float64
	_ = v5667
	var v5670 float64
	_ = v5670
	var v5674 int32
	_ = v5674
	var v5677 int32
	_ = v5677
	var v5680 int32
	_ = v5680
	var v5696 int32
	_ = v5696
	var v5698 float64
	_ = v5698
	var v5700 float64
	_ = v5700
	var v5702 float64
	_ = v5702
	var v5704 int32
	_ = v5704
	var v5705 int32
	_ = v5705
	var v5707 int32
	_ = v5707
	var v5709 int32
	_ = v5709
	var v5711 int32
	_ = v5711
	var v5713 int32
	_ = v5713
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5716 int32
	_ = v5716
	var v5720 int32
	_ = v5720
	var v5723 int32
	_ = v5723
	var v5728 int32
	_ = v5728
	var v5730 int32
	_ = v5730
	var v5736 int32
	_ = v5736
	var v5770 int32
	_ = v5770
	var v5774 int32
	_ = v5774
	var v5775 int32
	_ = v5775
	var v5776 int32
	_ = v5776
	var v5777 int32
	_ = v5777
	var v5778 int32
	_ = v5778
	var v5780 int32
	_ = v5780
	var v5782 int32
	_ = v5782
	var v5783 int32
	_ = v5783
	var v5787 int32
	_ = v5787
	var v5791 int32
	_ = v5791
	var v5792 int32
	_ = v5792
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5807 int32
	_ = v5807
	var v5841 int32
	_ = v5841
	var v5842 int32
	_ = v5842
	var v5843 int32
	_ = v5843
	var v5844 int32
	_ = v5844
	var v5845 int32
	_ = v5845
	var v5846 int32
	_ = v5846
	var v5847 int32
	_ = v5847
	var v5848 int32
	_ = v5848
	var v5849 int32
	_ = v5849
	var v5850 int32
	_ = v5850
	var v5851 int32
	_ = v5851
	var v5852 int32
	_ = v5852
	var v5853 int32
	_ = v5853
	var v5854 int32
	_ = v5854
	var v5855 int32
	_ = v5855
	var v5856 int32
	_ = v5856
	var v5857 int32
	_ = v5857
	var v5858 int32
	_ = v5858
	var v5859 int64
	_ = v5859
	var v5860 float64
	_ = v5860
	var v5862 int32
	_ = v5862
	var v5863 int32
	_ = v5863
	var v5875 float64
	_ = v5875
	var v5878 float64
	_ = v5878
	var v5882 int32
	_ = v5882
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5903 int32
	_ = v5903
	var v5905 float64
	_ = v5905
	var v5907 float64
	_ = v5907
	var v5909 float64
	_ = v5909
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5914 int32
	_ = v5914
	var v5916 int32
	_ = v5916
	var v5918 int32
	_ = v5918
	var v5920 int32
	_ = v5920
	var v5921 int32
	_ = v5921
	var v5922 int32
	_ = v5922
	var v5923 int32
	_ = v5923
	var v5924 int32
	_ = v5924
	var v5928 int32
	_ = v5928
	var v5931 int32
	_ = v5931
	var v5936 int32
	_ = v5936
	var v5937 int32
	_ = v5937
	var v5939 int32
	_ = v5939
	var v5978 int32
	_ = v5978
	var v5982 int32
	_ = v5982
	var v5983 int32
	_ = v5983
	var v5984 int32
	_ = v5984
	var v5985 int32
	_ = v5985
	var v5986 int32
	_ = v5986
	var v5988 int32
	_ = v5988
	var v5990 int32
	_ = v5990
	var v5991 int32
	_ = v5991
	var v5995 int32
	_ = v5995
	var v5999 int32
	_ = v5999
	var v6000 int32
	_ = v6000
	var v6002 int32
	_ = v6002
	var v6003 int32
	_ = v6003
	var v6007 int32
	_ = v6007
	var v6049 int32
	_ = v6049
	var v6050 int32
	_ = v6050
	var v6051 int32
	_ = v6051
	var v6052 int32
	_ = v6052
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6055 int32
	_ = v6055
	var v6056 int32
	_ = v6056
	var v6057 int32
	_ = v6057
	var v6058 int32
	_ = v6058
	var v6059 int32
	_ = v6059
	var v6060 int32
	_ = v6060
	var v6061 int32
	_ = v6061
	var v6062 int32
	_ = v6062
	var v6063 int32
	_ = v6063
	var v6064 int32
	_ = v6064
	var v6066 int32
	_ = v6066
	var v6067 int32
	_ = v6067
	var v6079 int32
	_ = v6079
	var v6081 float64
	_ = v6081
	var v6083 float64
	_ = v6083
	var v6085 float64
	_ = v6085
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6090 int32
	_ = v6090
	var v6092 int32
	_ = v6092
	var v6094 int32
	_ = v6094
	var v6097 int32
	_ = v6097
	var v6098 int32
	_ = v6098
	var v6099 int32
	_ = v6099
	var v6101 int32
	_ = v6101
	var v6102 int32
	_ = v6102
	var v6103 int32
	_ = v6103
	var v6106 int32
	_ = v6106
	var v6113 int32
	_ = v6113
	var v6114 int32
	_ = v6114
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6117 int32
	_ = v6117
	var v6129 int32
	_ = v6129
	var v6130 int32
	_ = v6130
	var v6131 int32
	_ = v6131
	var v6132 int32
	_ = v6132
	var v6133 int32
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6137 int32
	_ = v6137
	var v6138 int32
	_ = v6138
	var v6141 int32
	_ = v6141
	var v6143 int32
	_ = v6143
	var v6154 int32
	_ = v6154
	var v6156 float64
	_ = v6156
	var v6158 float64
	_ = v6158
	var v6160 float64
	_ = v6160
	var v6162 int32
	_ = v6162
	var v6163 int32
	_ = v6163
	var v6165 int32
	_ = v6165
	var v6167 int32
	_ = v6167
	var v6169 int32
	_ = v6169
	var v6172 int32
	_ = v6172
	var v6173 int32
	_ = v6173
	var v6174 int32
	_ = v6174
	var v6176 int32
	_ = v6176
	var v6177 int32
	_ = v6177
	var v6178 int32
	_ = v6178
	var v6181 int32
	_ = v6181
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6190 int32
	_ = v6190
	var v6191 int32
	_ = v6191
	var v6203 int32
	_ = v6203
	var v6204 int32
	_ = v6204
	var v6205 int32
	_ = v6205
	var v6206 int32
	_ = v6206
	var v6207 int32
	_ = v6207
	var v6208 int32
	_ = v6208
	var v6209 int32
	_ = v6209
	var v6211 int32
	_ = v6211
	var v6212 int32
	_ = v6212
	var v6215 int32
	_ = v6215
	var v6217 int32
	_ = v6217
	var v6219 int32
	_ = v6219
	var v6225 int32
	_ = v6225
	var v6234 int32
	_ = v6234
	var v6236 float64
	_ = v6236
	var v6238 float64
	_ = v6238
	var v6240 float64
	_ = v6240
	var v6242 int32
	_ = v6242
	var v6243 int32
	_ = v6243
	var v6245 int32
	_ = v6245
	var v6247 int32
	_ = v6247
	var v6249 int32
	_ = v6249
	var v6250 int32
	_ = v6250
	var v6252 int32
	_ = v6252
	var v6253 int32
	_ = v6253
	var v6254 int32
	_ = v6254
	var v6255 int32
	_ = v6255
	var v6256 int32
	_ = v6256
	var v6259 int32
	_ = v6259
	var v6262 int32
	_ = v6262
	var v6267 int32
	_ = v6267
	var v6268 int32
	_ = v6268
	var v6270 int32
	_ = v6270
	var v6309 int32
	_ = v6309
	var v6313 int32
	_ = v6313
	var v6314 int32
	_ = v6314
	var v6315 int32
	_ = v6315
	var v6316 int32
	_ = v6316
	var v6317 int32
	_ = v6317
	var v6319 int32
	_ = v6319
	var v6321 int32
	_ = v6321
	var v6322 int32
	_ = v6322
	var v6326 int32
	_ = v6326
	var v6330 int32
	_ = v6330
	var v6331 int32
	_ = v6331
	var v6333 int32
	_ = v6333
	var v6334 int32
	_ = v6334
	var v6338 int32
	_ = v6338
	var v6380 int32
	_ = v6380
	var v6381 int32
	_ = v6381
	var v6382 int32
	_ = v6382
	var v6383 int32
	_ = v6383
	var v6385 int32
	_ = v6385
	var v6386 int32
	_ = v6386
	var v6387 int32
	_ = v6387
	var v6402 int32
	_ = v6402
	var v6404 float64
	_ = v6404
	var v6406 float64
	_ = v6406
	var v6408 float64
	_ = v6408
	var v6410 int32
	_ = v6410
	var v6411 int32
	_ = v6411
	var v6413 int32
	_ = v6413
	var v6415 int32
	_ = v6415
	var v6417 int32
	_ = v6417
	var v6418 int32
	_ = v6418
	var v6420 int32
	_ = v6420
	var v6423 int32
	_ = v6423
	var v6425 int32
	_ = v6425
	var v6427 int32
	_ = v6427
	var v6430 int32
	_ = v6430
	var v6431 int32
	_ = v6431
	var v6432 int32
	_ = v6432
	var v6433 int32
	_ = v6433
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6439 int32
	_ = v6439
	var v6440 int32
	_ = v6440
	var v6448 int32
	_ = v6448
	var v6449 int32
	_ = v6449
	var v6451 int32
	_ = v6451
	var v6452 int32
	_ = v6452
	var v6453 int32
	_ = v6453
	var v6454 int32
	_ = v6454
	var v6455 int32
	_ = v6455
	var v6458 int32
	_ = v6458
	var v6461 int32
	_ = v6461
	var v6464 int32
	_ = v6464
	var v6474 int32
	_ = v6474
	var v6511 int32
	_ = v6511
	var v6512 int32
	_ = v6512
	var v6514 int32
	_ = v6514
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6519 int32
	_ = v6519
	var v6523 int32
	_ = v6523
	var v6527 int32
	_ = v6527
	var v6532 int32
	_ = v6532
	var v6533 int32
	_ = v6533
	var v6536 int32
	_ = v6536
	var v6537 int32
	_ = v6537
	var v6542 int32
	_ = v6542
	var v6584 int32
	_ = v6584
	var v6588 int32
	_ = v6588
	var v6589 int32
	_ = v6589
	var v6591 int32
	_ = v6591
	var v6592 int32
	_ = v6592
	var v6594 int32
	_ = v6594
	var v6595 int32
	_ = v6595
	var v6597 int32
	_ = v6597
	var v6598 int32
	_ = v6598
	var v6599 int32
	_ = v6599
	var v6602 int32
	_ = v6602
	var v6603 int32
	_ = v6603
	var v6604 int32
	_ = v6604
	var v6619 int32
	_ = v6619
	var v6621 int32
	_ = v6621
	var v6649 int32
	_ = v6649
	var v6650 int32
	_ = v6650
	var v6652 int32
	_ = v6652
	var v6653 int32
	_ = v6653
	var v6656 int32
	_ = v6656
	var v6659 int32
	_ = v6659
	var v6664 int32
	_ = v6664
	var v6667 int32
	_ = v6667
	var v6668 int32
	_ = v6668
	var v6718 int32
	_ = v6718
	var v6720 float64
	_ = v6720
	var v6722 float64
	_ = v6722
	var v6724 float64
	_ = v6724
	var v6726 int32
	_ = v6726
	var v6727 int32
	_ = v6727
	var v6729 int32
	_ = v6729
	var v6731 int32
	_ = v6731
	var v6739 int32
	_ = v6739
	var v6740 int32
	_ = v6740
	var v6748 int32
	_ = v6748
	var v6753 int32
	_ = v6753
	var v6801 int32
	_ = v6801
	var v6805 int32
	_ = v6805
	var v6810 int32
	_ = v6810
	var v6811 int32
	_ = v6811
	var v6813 int32
	_ = v6813
	var v6815 int32
	_ = v6815
	var v6816 int32
	_ = v6816
	var v6817 int32
	_ = v6817
	var v6818 int32
	_ = v6818
	var v6821 int32
	_ = v6821
	var v6822 int32
	_ = v6822
	var v6823 int32
	_ = v6823
	var v6824 int32
	_ = v6824
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6832 int32
	_ = v6832
	var v6838 int32
	_ = v6838
	var v6840 int32
	_ = v6840
	var v6846 int32
	_ = v6846
	var v6880 int32
	_ = v6880
	var v6884 int32
	_ = v6884
	var v6885 int32
	_ = v6885
	var v6886 int32
	_ = v6886
	var v6887 int32
	_ = v6887
	var v6888 int32
	_ = v6888
	var v6890 int32
	_ = v6890
	var v6892 int32
	_ = v6892
	var v6893 int32
	_ = v6893
	var v6897 int32
	_ = v6897
	var v6901 int32
	_ = v6901
	var v6902 int32
	_ = v6902
	var v6904 int32
	_ = v6904
	var v6905 int32
	_ = v6905
	var v6909 int32
	_ = v6909
	var v6960 int32
	_ = v6960
	var v6962 int32
	_ = v6962
	var v7002 int32
	_ = v7002
	var v7003 int32
	_ = v7003
	var v7010 int32
	_ = v7010
	var v7011 int32
	_ = v7011
	var v7012 int32
	_ = v7012
	var v7014 int32
	_ = v7014
	var v7052 int32
	_ = v7052
	var v7056 int32
	_ = v7056
	var v7057 int32
	_ = v7057
	var v7058 int32
	_ = v7058
	var v7063 int32
	_ = v7063
	var v7065 int32
	_ = v7065
	var v7066 int32
	_ = v7066
	var v7067 int32
	_ = v7067
	var v7068 int32
	_ = v7068
	var v7071 int32
	_ = v7071
	var v7072 int32
	_ = v7072
	var v7073 int32
	_ = v7073
	var v7075 int32
	_ = v7075
	var v7076 int32
	_ = v7076
	var v7082 int32
	_ = v7082
	var v7084 int32
	_ = v7084
	var v7128 int32
	_ = v7128
	var v7168 int32
	_ = v7168
	var v7175 int32
	_ = v7175
	var v7215 int32
	_ = v7215
	var v7216 int32
	_ = v7216
	var v7217 int32
	_ = v7217
	var v7218 int32
	_ = v7218
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
	var v7271 int32
	_ = v7271
	var v7272 int32
	_ = v7272
	var v7273 int32
	_ = v7273
	var v7276 int32
	_ = v7276
	var v7277 int32
	_ = v7277
	var v7282 int32
	_ = v7282
	var v7325 int32
	_ = v7325
	var v7326 int32
	_ = v7326
	var v7328 int32
	_ = v7328
	var v7329 int32
	_ = v7329
	var v7330 int32
	_ = v7330
	var v7336 int32
	_ = v7336
	var v7339 int32
	_ = v7339
	var v7340 int32
	_ = v7340
	var v7341 int32
	_ = v7341
	var v7344 int32
	_ = v7344
	var v7345 int32
	_ = v7345
	var v7391 int32
	_ = v7391
	var v7394 int32
	_ = v7394
	var v7398 int32
	_ = v7398
	var v7406 int32
	_ = v7406
	var v7411 int32
	_ = v7411
	var v7446 int32
	_ = v7446
	var v7450 int32
	_ = v7450
	var v7451 int32
	_ = v7451
	var v7452 int32
	_ = v7452
	var v7456 int32
	_ = v7456
	var v7457 int32
	_ = v7457
	var v7460 int32
	_ = v7460
	var v7464 int32
	_ = v7464
	var v7468 int32
	_ = v7468
	var v7471 int32
	_ = v7471
	var v7474 int32
	_ = v7474
	var v7475 int32
	_ = v7475
	var v7479 int32
	_ = v7479
	var v7487 int32
	_ = v7487
	var v7488 int32
	_ = v7488
	var v7491 int32
	_ = v7491
	var v7502 int32
	_ = v7502
	var v7505 int32
	_ = v7505
	var v7506 int32
	_ = v7506
	var v7509 int32
	_ = v7509
	var v7510 int32
	_ = v7510
	var v7519 int32
	_ = v7519
	var v7526 int32
	_ = v7526
	var v7529 int32
	_ = v7529
	var v7532 int32
	_ = v7532
	var v7534 int32
	_ = v7534
	var v7535 int32
	_ = v7535
	var v7540 int32
	_ = v7540
	var v7544 int32
	_ = v7544
	var v7545 int32
	_ = v7545
	var v7546 int32
	_ = v7546
	var v7555 int32
	_ = v7555
	var v7556 int32
	_ = v7556
	var v7557 int32
	_ = v7557
	var v7558 int32
	_ = v7558
	var v7559 int32
	_ = v7559
	var v7560 int32
	_ = v7560
	var v7561 int32
	_ = v7561
	var v7562 int32
	_ = v7562
	var v7563 int32
	_ = v7563
	var v7565 int32
	_ = v7565
	var v7567 int32
	_ = v7567
	var v7569 int32
	_ = v7569
	var v7571 int32
	_ = v7571
	var v7572 int32
	_ = v7572
	var v7573 int32
	_ = v7573
	var v7575 int32
	_ = v7575
	var v7579 int32
	_ = v7579
	var v7580 int32
	_ = v7580
	var v7588 int32
	_ = v7588
	var v7590 int32
	_ = v7590
	var v7592 int32
	_ = v7592
	var v7593 int32
	_ = v7593
	var v7604 int32
	_ = v7604
	var v7605 int32
	_ = v7605
	var v7607 int32
	_ = v7607
	var v7608 int32
	_ = v7608
	var v7611 int32
	_ = v7611
	var v7615 int32
	_ = v7615
	var v7637 int32
	_ = v7637
	var v7641 int32
	_ = v7641
	var v7650 int32
	_ = v7650
	var v7657 int32
	_ = v7657
	var v7658 int32
	_ = v7658
	var v7660 int32
	_ = v7660
	var v7661 int32
	_ = v7661
	var v7672 int32
	_ = v7672
	var v7707 int32
	_ = v7707
	var v7708 int32
	_ = v7708
	var v7709 int32
	_ = v7709
	var v7710 int32
	_ = v7710
	var v7713 int32
	_ = v7713
	var v7714 int32
	_ = v7714
	var v7716 int32
	_ = v7716
	var v7717 int32
	_ = v7717
	var v7718 int32
	_ = v7718
	var v7719 int32
	_ = v7719
	var v7720 int32
	_ = v7720
	var v7721 int32
	_ = v7721
	var v7722 int32
	_ = v7722
	var v7725 int32
	_ = v7725
	var v7734 int32
	_ = v7734
	var v7776 int32
	_ = v7776
	var v7777 int32
	_ = v7777
	var v7779 int32
	_ = v7779
	var v7780 int32
	_ = v7780
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7785 int32
	_ = v7785
	var v7788 int32
	_ = v7788
	var v7789 int32
	_ = v7789
	var v7790 int32
	_ = v7790
	var v7793 int32
	_ = v7793
	var v7796 int32
	_ = v7796
	var v7797 int32
	_ = v7797
	var v7805 int32
	_ = v7805
	var v7844 int32
	_ = v7844
	var v7845 int32
	_ = v7845
	var v7848 int32
	_ = v7848
	var v7850 int32
	_ = v7850
	var v7851 int32
	_ = v7851
	var v7852 int32
	_ = v7852
	var v7858 int32
	_ = v7858
	var v7863 int32
	_ = v7863
	var v7865 int32
	_ = v7865
	var v7868 int32
	_ = v7868
	var v7871 float64
	_ = v7871
	var v7872 float64
	_ = v7872
	var v7873 int32
	_ = v7873
	var v7876 int32
	_ = v7876
	var v7879 int32
	_ = v7879
	var v7880 int32
	_ = v7880
	var v7881 int32
	_ = v7881
	var v7886 float64
	_ = v7886
	var v7889 int32
	_ = v7889
	var v7890 float64
	_ = v7890
	var v7896 float64
	_ = v7896
	var v7902 float64
	_ = v7902
	var v7904 float64
	_ = v7904
	var v7906 float64
	_ = v7906
	var v7908 int32
	_ = v7908
	var v7909 int32
	_ = v7909
	var v7912 int32
	_ = v7912
	var v7915 int32
	_ = v7915
	var v7916 int32
	_ = v7916
	var v7919 int32
	_ = v7919
	var v7920 int32
	_ = v7920
	var v7921 int32
	_ = v7921
	var v7922 int32
	_ = v7922
	var v7931 int32
	_ = v7931
	var v7932 int32
	_ = v7932
	var v7934 int32
	_ = v7934
	var v7935 int32
	_ = v7935
	var v7936 int32
	_ = v7936
	var v7937 int32
	_ = v7937
	var v7938 int32
	_ = v7938
	var v7941 int32
	_ = v7941
	var v7948 int32
	_ = v7948
	var v7992 int32
	_ = v7992
	var v7993 int32
	_ = v7993
	var v7995 int32
	_ = v7995
	var v7996 int32
	_ = v7996
	var v7997 int32
	_ = v7997
	var v7998 int32
	_ = v7998
	var v7999 int32
	_ = v7999
	var v8002 int32
	_ = v8002
	var v8005 int32
	_ = v8005
	var v8006 int32
	_ = v8006
	var v8007 int32
	_ = v8007
	var v8010 int32
	_ = v8010
	var v8011 int32
	_ = v8011
	var v8061 int32
	_ = v8061
	var v8062 int32
	_ = v8062
	var v8065 int32
	_ = v8065
	var v8066 int32
	_ = v8066
	var v8071 int32
	_ = v8071
	var v8114 int32
	_ = v8114
	var v8115 int32
	_ = v8115
	var v8117 int32
	_ = v8117
	var v8120 int32
	_ = v8120
	var v8121 int32
	_ = v8121
	var v8125 int32
	_ = v8125
	var v8128 int32
	_ = v8128
	var v8129 int32
	_ = v8129
	var v8175 int32
	_ = v8175
	var v8176 int32
	_ = v8176
	var v8177 int32
	_ = v8177
	var v8181 int32
	_ = v8181
	var v8182 int32
	_ = v8182
	var v8185 int32
	_ = v8185
	var v8190 int32
	_ = v8190
	var v8191 int32
	_ = v8191
	var v8192 int32
	_ = v8192
	var v8232 int32
	_ = v8232
	var v8236 int32
	_ = v8236
	var v8237 int32
	_ = v8237
	var v8238 int32
	_ = v8238
	var v8239 int32
	_ = v8239
	var v8240 int32
	_ = v8240
	var v8242 int32
	_ = v8242
	var v8244 int32
	_ = v8244
	var v8245 int32
	_ = v8245
	var v8249 int32
	_ = v8249
	var v8253 int32
	_ = v8253
	var v8254 int32
	_ = v8254
	var v8256 int32
	_ = v8256
	var v8257 int32
	_ = v8257
	var v8262 int32
	_ = v8262
	var v8303 float64
	_ = v8303
	var v8305 int32
	_ = v8305
	var v8306 int32
	_ = v8306
	var v8318 float64
	_ = v8318
	var v8321 float64
	_ = v8321
	var v8325 int32
	_ = v8325
	var v8328 int32
	_ = v8328
	var v8336 int64
	_ = v8336
	var v8338 int32
	_ = v8338
	var v8350 int32
	_ = v8350
	var v8392 int32
	_ = v8392
	var v8394 float64
	_ = v8394
	var v8396 float64
	_ = v8396
	var v8398 float64
	_ = v8398
	var v8400 int32
	_ = v8400
	var v8401 int32
	_ = v8401
	var v8403 int32
	_ = v8403
	var v8405 int32
	_ = v8405
	var v8409 int32
	_ = v8409
	var v8457 int32
	_ = v8457
	var v8461 int32
	_ = v8461
	var v8466 int32
	_ = v8466
	var v8470 int32
	_ = v8470
	var v8474 int32
	_ = v8474
	var v8479 int32
	_ = v8479
	var v8483 int32
	_ = v8483
	var v8489 int32
	_ = v8489
	var v8494 int32
	_ = v8494
	var v8498 int32
	_ = v8498
	var v8504 int32
	_ = v8504
	var v8509 int32
	_ = v8509
	var v8510 int32
	_ = v8510
	var v8513 int32
	_ = v8513
	var v8514 int32
	_ = v8514
	var v8515 int32
	_ = v8515
	var v8516 int32
	_ = v8516
	var v8517 int32
	_ = v8517
	var v8518 int32
	_ = v8518
	var v8521 int32
	_ = v8521
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
	var v8530 int32
	_ = v8530
	var v8571 int32
	_ = v8571
	var v8574 int32
	_ = v8574
	var v8576 int32
	_ = v8576
	var v8580 int32
	_ = v8580
	var v8583 int32
	_ = v8583
	var v8588 int32
	_ = v8588
	var v8589 int32
	_ = v8589
	var v8590 int32
	_ = v8590
	var v8593 int32
	_ = v8593
	var v8595 int32
	_ = v8595
	var v8598 int32
	_ = v8598
	var v8599 int32
	_ = v8599
	var v8605 int32
	_ = v8605
	var v8607 int32
	_ = v8607
	var v8609 int32
	_ = v8609
	var v8614 int32
	_ = v8614
	var v8618 int32
	_ = v8618
	var v8619 int32
	_ = v8619
	var v8620 int32
	_ = v8620
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8626 int32
	_ = v8626
	var v8627 int32
	_ = v8627
	var v8631 int32
	_ = v8631
	var v8632 int32
	_ = v8632
	var v8633 int32
	_ = v8633
	var v8635 int32
	_ = v8635
	var v8636 int32
	_ = v8636
	var v8639 int32
	_ = v8639
	var v8640 int32
	_ = v8640
	var v8647 int32
	_ = v8647
	var v8648 int32
	_ = v8648
	var v8657 int32
	_ = v8657
	var v8659 float64
	_ = v8659
	var v8661 float64
	_ = v8661
	var v8663 float64
	_ = v8663
	var v8665 int32
	_ = v8665
	var v8666 int32
	_ = v8666
	var v8668 int32
	_ = v8668
	var v8670 int32
	_ = v8670
	var v8672 int32
	_ = v8672
	var v8675 int32
	_ = v8675
	var v8676 int32
	_ = v8676
	var v8678 int32
	_ = v8678
	var v8679 int32
	_ = v8679
	var v8682 int32
	_ = v8682
	var v8683 int32
	_ = v8683
	var v8689 int32
	_ = v8689
	var v8691 float64
	_ = v8691
	var v8693 float64
	_ = v8693
	var v8695 float64
	_ = v8695
	var v8697 int32
	_ = v8697
	var v8698 int32
	_ = v8698
	var v8700 int32
	_ = v8700
	var v8702 int32
	_ = v8702
	var v8704 int32
	_ = v8704
	var v8706 int32
	_ = v8706
	var v8707 int32
	_ = v8707
	var v8708 int32
	_ = v8708
	var v8709 int32
	_ = v8709
	var v8710 int32
	_ = v8710
	var v8714 int32
	_ = v8714
	var v8717 int32
	_ = v8717
	var v8722 int32
	_ = v8722
	var v8723 int32
	_ = v8723
	var v8725 int32
	_ = v8725
	var v8764 int32
	_ = v8764
	var v8768 int32
	_ = v8768
	var v8769 int32
	_ = v8769
	var v8770 int32
	_ = v8770
	var v8771 int32
	_ = v8771
	var v8772 int32
	_ = v8772
	var v8774 int32
	_ = v8774
	var v8776 int32
	_ = v8776
	var v8777 int32
	_ = v8777
	var v8781 int32
	_ = v8781
	var v8785 int32
	_ = v8785
	var v8786 int32
	_ = v8786
	var v8788 int32
	_ = v8788
	var v8789 int32
	_ = v8789
	var v8793 int32
	_ = v8793
	var v8836 int32
	_ = v8836
	var v8837 int32
	_ = v8837
	var v8838 int32
	_ = v8838
	var v8846 int32
	_ = v8846
	var v8848 float64
	_ = v8848
	var v8850 float64
	_ = v8850
	var v8852 float64
	_ = v8852
	var v8854 int32
	_ = v8854
	var v8855 int32
	_ = v8855
	var v8857 int32
	_ = v8857
	var v8859 int32
	_ = v8859
	var v8861 int32
	_ = v8861
	var v8864 int32
	_ = v8864
	var v8865 int32
	_ = v8865
	var v8866 int32
	_ = v8866
	var v8868 int32
	_ = v8868
	var v8869 int32
	_ = v8869
	var v8870 int32
	_ = v8870
	var v8875 int32
	_ = v8875
	var v8877 int32
	_ = v8877
	var v8878 int32
	_ = v8878
	var v8879 int32
	_ = v8879
	var v8882 int32
	_ = v8882
	var v8887 int32
	_ = v8887
	var v8890 int32
	_ = v8890
	var v8894 int32
	_ = v8894
	var v8896 int32
	_ = v8896
	var v8898 int32
	_ = v8898
	var v8899 int32
	_ = v8899
	var v8900 int32
	_ = v8900
	var v8901 int32
	_ = v8901
	var v8906 int32
	_ = v8906
	var v8907 int32
	_ = v8907
	var v8910 int32
	_ = v8910
	var v8915 int32
	_ = v8915
	var v8916 int32
	_ = v8916
	var v8917 int32
	_ = v8917
	var v8957 int32
	_ = v8957
	var v8961 int32
	_ = v8961
	var v8962 int32
	_ = v8962
	var v8963 int32
	_ = v8963
	var v8964 int32
	_ = v8964
	var v8965 int32
	_ = v8965
	var v8967 int32
	_ = v8967
	var v8969 int32
	_ = v8969
	var v8970 int32
	_ = v8970
	var v8974 int32
	_ = v8974
	var v8978 int32
	_ = v8978
	var v8979 int32
	_ = v8979
	var v8981 int32
	_ = v8981
	var v8982 int32
	_ = v8982
	var v8984 int32
	_ = v8984
	var v8986 int32
	_ = v8986
	var v8987 int32
	_ = v8987
	var v8988 int32
	_ = v8988
	var v8989 int32
	_ = v8989
	var v8993 int32
	_ = v8993
	var v8996 int32
	_ = v8996
	var v9001 int32
	_ = v9001
	var v9002 int32
	_ = v9002
	var v9004 int32
	_ = v9004
	var v9043 int32
	_ = v9043
	var v9047 int32
	_ = v9047
	var v9048 int32
	_ = v9048
	var v9049 int32
	_ = v9049
	var v9050 int32
	_ = v9050
	var v9051 int32
	_ = v9051
	var v9053 int32
	_ = v9053
	var v9055 int32
	_ = v9055
	var v9056 int32
	_ = v9056
	var v9060 int32
	_ = v9060
	var v9064 int32
	_ = v9064
	var v9065 int32
	_ = v9065
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9073 int32
	_ = v9073
	var v9114 int32
	_ = v9114
	var v9115 int32
	_ = v9115
	var v9116 int32
	_ = v9116
	var v9122 int32
	_ = v9122
	var v9124 int32
	_ = v9124
	var v9164 float64
	_ = v9164
	var v9166 float64
	_ = v9166
	var v9168 float64
	_ = v9168
	var v9170 int32
	_ = v9170
	var v9171 int32
	_ = v9171
	var v9174 int32
	_ = v9174
	var v9175 int32
	_ = v9175
	var v9176 int32
	_ = v9176
	var v9186 int32
	_ = v9186
	var v9188 float64
	_ = v9188
	var v9190 float64
	_ = v9190
	var v9192 float64
	_ = v9192
	var v9194 int32
	_ = v9194
	var v9195 int32
	_ = v9195
	var v9197 int32
	_ = v9197
	var v9204 int32
	_ = v9204
	var v9243 int32
	_ = v9243
	var v9245 int32
	_ = v9245
	var v9248 int32
	_ = v9248
	var v9260 int32
	_ = v9260
	var v9295 int32
	_ = v9295
	var v9299 int32
	_ = v9299
	var v9300 int32
	_ = v9300
	var v9301 int32
	_ = v9301
	var v9302 int32
	_ = v9302
	var v9303 int32
	_ = v9303
	var v9304 int32
	_ = v9304
	var v9305 int64
	_ = v9305
	var v9306 int32
	_ = v9306
	var v9308 int32
	_ = v9308
	var v9309 int32
	_ = v9309
	var v9312 int32
	_ = v9312
	var v9313 int64
	_ = v9313
	var v9317 int32
	_ = v9317
	var v9325 int32
	_ = v9325
	var v9326 int32
	_ = v9326
	var v9328 int32
	_ = v9328
	var v9329 float64
	_ = v9329
	var v9331 float64
	_ = v9331
	var v9335 int32
	_ = v9335
	var v9336 int32
	_ = v9336
	var v9337 int32
	_ = v9337
	var v9341 int32
	_ = v9341
	var v9342 int32
	_ = v9342
	var v9344 int32
	_ = v9344
	var v9346 int32
	_ = v9346
	var v9348 int32
	_ = v9348
	var v9350 int32
	_ = v9350
	var v9351 int32
	_ = v9351
	var v9352 int32
	_ = v9352
	var v9353 int32
	_ = v9353
	var v9354 int32
	_ = v9354
	var v9356 int32
	_ = v9356
	var v9357 int32
	_ = v9357
	var v9359 int32
	_ = v9359
	var v9360 int32
	_ = v9360
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
	var v9370 int32
	_ = v9370
	var v9371 int32
	_ = v9371
	var v9374 int32
	_ = v9374
	var v9375 int32
	_ = v9375
	var v9376 int32
	_ = v9376
	var v9377 int32
	_ = v9377
	var v9383 int32
	_ = v9383
	var v9384 int32
	_ = v9384
	var v9386 int32
	_ = v9386
	var v9389 int32
	_ = v9389
	var v9390 int32
	_ = v9390
	var v9391 int32
	_ = v9391
	var v9392 int32
	_ = v9392
	var v9393 int32
	_ = v9393
	var v9394 int32
	_ = v9394
	var v9396 int32
	_ = v9396
	var v9397 int32
	_ = v9397
	var v9398 int32
	_ = v9398
	var v9400 int32
	_ = v9400
	var v9401 int32
	_ = v9401
	var v9402 int32
	_ = v9402
	var v9408 int32
	_ = v9408
	var v9410 int32
	_ = v9410
	var v9412 int32
	_ = v9412
	var v9418 int32
	_ = v9418
	var v9419 int32
	_ = v9419
	var v9421 int32
	_ = v9421
	var v9422 int32
	_ = v9422
	var v9423 int32
	_ = v9423
	var v9426 int32
	_ = v9426
	var v9431 int32
	_ = v9431
	var v9432 int32
	_ = v9432
	var v9478 int32
	_ = v9478
	var v9479 int32
	_ = v9479
	var v9480 int32
	_ = v9480
	var v9484 int32
	_ = v9484
	var v9485 int32
	_ = v9485
	var v9488 int32
	_ = v9488
	var v9493 int32
	_ = v9493
	var v9495 int32
	_ = v9495
	var v9500 int32
	_ = v9500
	var v9535 int32
	_ = v9535
	var v9539 int32
	_ = v9539
	var v9540 int32
	_ = v9540
	var v9541 int32
	_ = v9541
	var v9542 int32
	_ = v9542
	var v9543 int32
	_ = v9543
	var v9545 int32
	_ = v9545
	var v9547 int32
	_ = v9547
	var v9548 int32
	_ = v9548
	var v9552 int32
	_ = v9552
	var v9556 int32
	_ = v9556
	var v9557 int32
	_ = v9557
	var v9559 int32
	_ = v9559
	var v9560 int32
	_ = v9560
	var v9566 int32
	_ = v9566
	var v9606 int32
	_ = v9606
	var v9608 int32
	_ = v9608
	var v9609 int32
	_ = v9609
	var v9618 int32
	_ = v9618
	var v9620 float64
	_ = v9620
	var v9622 float64
	_ = v9622
	var v9624 float64
	_ = v9624
	var v9626 int32
	_ = v9626
	var v9627 int32
	_ = v9627
	var v9629 int32
	_ = v9629
	var v9631 int32
	_ = v9631
	var v9633 int32
	_ = v9633
	var v9635 int32
	_ = v9635
	var v9636 int32
	_ = v9636
	var v9640 int32
	_ = v9640
	var v9643 int32
	_ = v9643
	var v9648 int32
	_ = v9648
	var v9649 int32
	_ = v9649
	var v9650 int32
	_ = v9650
	var v9690 int32
	_ = v9690
	var v9694 int32
	_ = v9694
	var v9695 int32
	_ = v9695
	var v9696 int32
	_ = v9696
	var v9697 int32
	_ = v9697
	var v9698 int32
	_ = v9698
	var v9700 int32
	_ = v9700
	var v9702 int32
	_ = v9702
	var v9703 int32
	_ = v9703
	var v9707 int32
	_ = v9707
	var v9711 int32
	_ = v9711
	var v9712 int32
	_ = v9712
	var v9714 int32
	_ = v9714
	var v9715 int32
	_ = v9715
	var v9720 int32
	_ = v9720
	var v9761 int32
	_ = v9761
	var v9762 int32
	_ = v9762
	var v9763 int32
	_ = v9763
	var v9765 int32
	_ = v9765
	var v9766 int32
	_ = v9766
	var v9775 int32
	_ = v9775
	var v9777 float64
	_ = v9777
	var v9779 float64
	_ = v9779
	var v9781 float64
	_ = v9781
	var v9783 int32
	_ = v9783
	var v9784 int32
	_ = v9784
	var v9786 int32
	_ = v9786
	var v9788 int32
	_ = v9788
	var v9790 int32
	_ = v9790
	var v9791 int32
	_ = v9791
	var v9793 int32
	_ = v9793
	var v9794 int32
	_ = v9794
	var v9797 int32
	_ = v9797
	var v9798 int32
	_ = v9798
	var v9802 int32
	_ = v9802
	var v9805 int32
	_ = v9805
	var v9812 int32
	_ = v9812
	var v9814 int32
	_ = v9814
	var v9816 int32
	_ = v9816
	var v9852 int32
	_ = v9852
	var v9856 int32
	_ = v9856
	var v9857 int32
	_ = v9857
	var v9858 int32
	_ = v9858
	var v9859 int32
	_ = v9859
	var v9860 int32
	_ = v9860
	var v9862 int32
	_ = v9862
	var v9864 int32
	_ = v9864
	var v9865 int32
	_ = v9865
	var v9869 int32
	_ = v9869
	var v9873 int32
	_ = v9873
	var v9874 int32
	_ = v9874
	var v9876 int32
	_ = v9876
	var v9877 int32
	_ = v9877
	var v9881 int32
	_ = v9881
	var v9926 int32
	_ = v9926
	var v9934 int32
	_ = v9934
	var v9936 int32
	_ = v9936
	var v9972 int32
	_ = v9972
	var v9973 int32
	_ = v9973
	var v9974 int32
	_ = v9974
	var v9976 float64
	_ = v9976
	var v9978 float64
	_ = v9978
	var v9980 float64
	_ = v9980
	var v9982 int32
	_ = v9982
	var v9983 int32
	_ = v9983
	var v9985 int32
	_ = v9985
	var v9987 int32
	_ = v9987
	var v9988 int32
	_ = v9988
	var v9994 int32
	_ = v9994
	var v9996 int32
	_ = v9996
	var v9997 int32
	_ = v9997
	var v10003 int32
	_ = v10003
	var v10010 int32
	_ = v10010
	var v10011 int32
	_ = v10011
	var v10012 int32
	_ = v10012
	var v10013 int32
	_ = v10013
	var v10014 int32
	_ = v10014
	var v10015 int32
	_ = v10015
	var v10016 int32
	_ = v10016
	var v10019 int32
	_ = v10019
	var v10020 int32
	_ = v10020
	var v10031 int32
	_ = v10031
	var v10034 int32
	_ = v10034
	var v10067 int32
	_ = v10067
	var v10071 int32
	_ = v10071
	var v10073 int32
	_ = v10073
	var v10074 int32
	_ = v10074
	var v10075 int32
	_ = v10075
	var v10076 int32
	_ = v10076
	var v10077 int32
	_ = v10077
	var v10089 int32
	_ = v10089
	var v10090 int32
	_ = v10090
	var v10091 int32
	_ = v10091
	var v10092 int32
	_ = v10092
	var v10093 int32
	_ = v10093
	var v10095 int32
	_ = v10095
	var v10103 int32
	_ = v10103
	var v10104 int32
	_ = v10104
	var v10105 int32
	_ = v10105
	var v10108 int32
	_ = v10108
	var v10109 int32
	_ = v10109
	var v10111 int32
	_ = v10111
	var v10112 int32
	_ = v10112
	var v10114 int32
	_ = v10114
	var v10116 int32
	_ = v10116
	var v10119 int32
	_ = v10119
	var v10120 int32
	_ = v10120
	var v10121 int32
	_ = v10121
	var v10126 int32
	_ = v10126
	var v10127 int32
	_ = v10127
	var v10128 int32
	_ = v10128
	var v10131 int32
	_ = v10131
	var v10132 int32
	_ = v10132
	var v10133 int32
	_ = v10133
	var v10136 int32
	_ = v10136
	var v10137 int32
	_ = v10137
	var v10139 int32
	_ = v10139
	var v10144 int32
	_ = v10144
	var v10157 int32
	_ = v10157
	var v10158 int32
	_ = v10158
	var v10167 int32
	_ = v10167
	var v10171 int32
	_ = v10171
	var v10175 int32
	_ = v10175
	var v10177 int32
	_ = v10177
	var v10181 int32
	_ = v10181
	var v10182 int32
	_ = v10182
	var v10187 int32
	_ = v10187
	var v10190 int32
	_ = v10190
	var v10193 int32
	_ = v10193
	var v10198 int32
	_ = v10198
	var v10199 int32
	_ = v10199
	var v10203 int32
	_ = v10203
	var v10211 int32
	_ = v10211
	var v10212 int32
	_ = v10212
	var v10213 int32
	_ = v10213
	var v10214 int32
	_ = v10214
	var v10216 int32
	_ = v10216
	var v10217 int32
	_ = v10217
	var v10220 int32
	_ = v10220
	var v10222 int32
	_ = v10222
	var v10223 int32
	_ = v10223
	var v10224 int32
	_ = v10224
	var v10230 int32
	_ = v10230
	var v10235 int32
	_ = v10235
	var v10237 int32
	_ = v10237
	var v10240 int32
	_ = v10240
	var v10241 float64
	_ = v10241
	var v10242 float64
	_ = v10242
	var v10243 int32
	_ = v10243
	var v10246 int32
	_ = v10246
	var v10247 float64
	_ = v10247
	var v10249 int32
	_ = v10249
	var v10250 int32
	_ = v10250
	var v10251 int32
	_ = v10251
	var v10256 float64
	_ = v10256
	var v10259 int32
	_ = v10259
	var v10260 float64
	_ = v10260
	var v10266 float64
	_ = v10266
	var v10272 float64
	_ = v10272
	var v10274 float64
	_ = v10274
	var v10276 float64
	_ = v10276
	var v10278 int32
	_ = v10278
	var v10279 int32
	_ = v10279
	var v10282 int32
	_ = v10282
	var v10284 int32
	_ = v10284
	var v10291 int32
	_ = v10291
	var v10292 int32
	_ = v10292
	var v10294 int32
	_ = v10294
	var v10295 int32
	_ = v10295
	var v10300 int32
	_ = v10300
	var v10304 int32
	_ = v10304
	var v10309 int32
	_ = v10309
	var v10321 int32
	_ = v10321
	var v10357 int32
	_ = v10357
	var v10360 int32
	_ = v10360
	var v10362 int32
	_ = v10362
	var v10363 int32
	_ = v10363
	var v10366 int32
	_ = v10366
	var v10367 int32
	_ = v10367
	var v10368 int32
	_ = v10368
	var v10377 int32
	_ = v10377
	var v10378 int32
	_ = v10378
	var v10379 int32
	_ = v10379
	var v10380 int32
	_ = v10380
	var v10381 int32
	_ = v10381
	var v10382 int32
	_ = v10382
	var v10383 int32
	_ = v10383
	var v10384 int32
	_ = v10384
	var v10388 int32
	_ = v10388
	var v10391 int32
	_ = v10391
	var v10397 int32
	_ = v10397
	var v10399 int32
	_ = v10399
	var v10400 int32
	_ = v10400
	var v10438 int32
	_ = v10438
	var v10442 int32
	_ = v10442
	var v10443 int32
	_ = v10443
	var v10444 int32
	_ = v10444
	var v10445 int32
	_ = v10445
	var v10446 int32
	_ = v10446
	var v10448 int32
	_ = v10448
	var v10450 int32
	_ = v10450
	var v10451 int32
	_ = v10451
	var v10455 int32
	_ = v10455
	var v10459 int32
	_ = v10459
	var v10460 int32
	_ = v10460
	var v10462 int32
	_ = v10462
	var v10463 int32
	_ = v10463
	var v10467 int32
	_ = v10467
	var v10512 int32
	_ = v10512
	var v10519 int32
	_ = v10519
	var v10520 int32
	_ = v10520
	var v10558 int32
	_ = v10558
	var v10559 int32
	_ = v10559
	var v10560 int32
	_ = v10560
	var v10570 int32
	_ = v10570
	var v10573 int32
	_ = v10573
	var v10575 int32
	_ = v10575
	var v10576 int32
	_ = v10576
	var v10582 int32
	_ = v10582
	var v10583 int32
	_ = v10583
	var v10585 int32
	_ = v10585
	var v10586 int32
	_ = v10586
	var v10595 int32
	_ = v10595
	var v10597 float64
	_ = v10597
	var v10599 float64
	_ = v10599
	var v10601 float64
	_ = v10601
	var v10603 int32
	_ = v10603
	var v10604 int32
	_ = v10604
	var v10606 int32
	_ = v10606
	var v10608 int32
	_ = v10608
	var v10611 int32
	_ = v10611
	var v10612 int32
	_ = v10612
	var v10620 int32
	_ = v10620
	var v10622 int32
	_ = v10622
	var v10623 int32
	_ = v10623
	var v10636 int32
	_ = v10636
	var v10637 int32
	_ = v10637
	var v10638 int32
	_ = v10638
	var v10639 int32
	_ = v10639
	var v10641 int32
	_ = v10641
	var v10643 int32
	_ = v10643
	var v10645 int32
	_ = v10645
	var v10648 int32
	_ = v10648
	var v10649 int32
	_ = v10649
	var v10652 int32
	_ = v10652
	var v10657 int32
	_ = v10657
	var v10658 int32
	_ = v10658
	var v10659 int32
	_ = v10659
	var v10662 int32
	_ = v10662
	var v10674 int32
	_ = v10674
	var v10682 int32
	_ = v10682
	var v10683 int32
	_ = v10683
	var v10710 int32
	_ = v10710
	var v10714 int32
	_ = v10714
	var v10716 int32
	_ = v10716
	var v10717 int32
	_ = v10717
	var v10720 int32
	_ = v10720
	var v10721 int32
	_ = v10721
	var v10722 int32
	_ = v10722
	var v10734 int32
	_ = v10734
	var v10735 int32
	_ = v10735
	var v10736 int32
	_ = v10736
	var v10737 int32
	_ = v10737
	var v10739 int32
	_ = v10739
	var v10747 int32
	_ = v10747
	var v10748 int32
	_ = v10748
	var v10749 int32
	_ = v10749
	var v10752 int32
	_ = v10752
	var v10753 int32
	_ = v10753
	var v10755 int32
	_ = v10755
	var v10756 int32
	_ = v10756
	var v10758 int32
	_ = v10758
	var v10760 int32
	_ = v10760
	var v10763 int32
	_ = v10763
	var v10764 int32
	_ = v10764
	var v10765 int32
	_ = v10765
	var v10770 int32
	_ = v10770
	var v10771 int32
	_ = v10771
	var v10772 int32
	_ = v10772
	var v10775 int32
	_ = v10775
	var v10776 int32
	_ = v10776
	var v10777 int32
	_ = v10777
	var v10780 int32
	_ = v10780
	var v10781 int32
	_ = v10781
	var v10783 int32
	_ = v10783
	var v10788 int32
	_ = v10788
	var v10801 int32
	_ = v10801
	var v10802 int32
	_ = v10802
	var v10811 int32
	_ = v10811
	var v10815 int32
	_ = v10815
	var v10819 int32
	_ = v10819
	var v10821 int32
	_ = v10821
	var v10825 int32
	_ = v10825
	var v10826 int32
	_ = v10826
	var v10831 int32
	_ = v10831
	var v10834 int32
	_ = v10834
	var v10837 int32
	_ = v10837
	var v10842 int32
	_ = v10842
	var v10843 int32
	_ = v10843
	var v10847 int32
	_ = v10847
	var v10855 int32
	_ = v10855
	var v10856 int32
	_ = v10856
	var v10857 int32
	_ = v10857
	var v10858 int32
	_ = v10858
	var v10860 int32
	_ = v10860
	var v10861 int32
	_ = v10861
	var v10864 int32
	_ = v10864
	var v10866 int32
	_ = v10866
	var v10867 int32
	_ = v10867
	var v10868 int32
	_ = v10868
	var v10874 int32
	_ = v10874
	var v10879 int32
	_ = v10879
	var v10881 int32
	_ = v10881
	var v10884 int32
	_ = v10884
	var v10885 float64
	_ = v10885
	var v10886 float64
	_ = v10886
	var v10887 int32
	_ = v10887
	var v10890 int32
	_ = v10890
	var v10891 float64
	_ = v10891
	var v10893 int32
	_ = v10893
	var v10894 int32
	_ = v10894
	var v10895 int32
	_ = v10895
	var v10900 float64
	_ = v10900
	var v10903 int32
	_ = v10903
	var v10904 float64
	_ = v10904
	var v10910 float64
	_ = v10910
	var v10916 float64
	_ = v10916
	var v10918 float64
	_ = v10918
	var v10920 float64
	_ = v10920
	var v10922 int32
	_ = v10922
	var v10923 int32
	_ = v10923
	var v10926 int32
	_ = v10926
	var v10928 int32
	_ = v10928
	var v10937 int32
	_ = v10937
	var v10938 int32
	_ = v10938
	var v10940 int32
	_ = v10940
	var v10941 int32
	_ = v10941
	var v10942 int32
	_ = v10942
	var v10944 int32
	_ = v10944
	var v10945 int32
	_ = v10945
	var v10950 int32
	_ = v10950
	var v10954 int32
	_ = v10954
	var v10959 int32
	_ = v10959
	var v10976 int32
	_ = v10976
	var v10977 int32
	_ = v10977
	var v10987 int32
	_ = v10987
	var v11007 int32
	_ = v11007
	var v11010 int32
	_ = v11010
	var v11012 int32
	_ = v11012
	var v11013 int32
	_ = v11013
	var v11014 int32
	_ = v11014
	var v11015 int32
	_ = v11015
	var v11017 int32
	_ = v11017
	var v11018 int32
	_ = v11018
	var v11019 int32
	_ = v11019
	var v11020 int32
	_ = v11020
	var v11021 int32
	_ = v11021
	var v11022 int32
	_ = v11022
	var v11023 int32
	_ = v11023
	var v11026 int32
	_ = v11026
	var v11027 int32
	_ = v11027
	var v11028 int32
	_ = v11028
	var v11034 int32
	_ = v11034
	var v11036 int32
	_ = v11036
	var v11038 float64
	_ = v11038
	var v11040 float64
	_ = v11040
	var v11042 float64
	_ = v11042
	var v11044 int32
	_ = v11044
	var v11045 int32
	_ = v11045
	var v11047 int32
	_ = v11047
	var v11049 int32
	_ = v11049
	var v11056 int32
	_ = v11056
	var v11057 int32
	_ = v11057
	var v11058 int32
	_ = v11058
	var v11059 int32
	_ = v11059
	var v11060 int32
	_ = v11060
	var v11061 int32
	_ = v11061
	var v11062 int32
	_ = v11062
	var v11063 int32
	_ = v11063
	var v11064 int32
	_ = v11064
	var v11068 int32
	_ = v11068
	var v11071 int32
	_ = v11071
	var v11076 int32
	_ = v11076
	var v11077 int32
	_ = v11077
	var v11079 int32
	_ = v11079
	var v11118 int32
	_ = v11118
	var v11122 int32
	_ = v11122
	var v11123 int32
	_ = v11123
	var v11124 int32
	_ = v11124
	var v11125 int32
	_ = v11125
	var v11126 int32
	_ = v11126
	var v11128 int32
	_ = v11128
	var v11130 int32
	_ = v11130
	var v11131 int32
	_ = v11131
	var v11135 int32
	_ = v11135
	var v11139 int32
	_ = v11139
	var v11140 int32
	_ = v11140
	var v11142 int32
	_ = v11142
	var v11143 int32
	_ = v11143
	var v11147 int32
	_ = v11147
	var v11189 int32
	_ = v11189
	var v11190 int32
	_ = v11190
	var v11191 int32
	_ = v11191
	var v11192 int32
	_ = v11192
	var v11193 int32
	_ = v11193
	var v11194 int32
	_ = v11194
	var v11195 int32
	_ = v11195
	var v11197 int32
	_ = v11197
	var v11199 int32
	_ = v11199
	var v11200 int32
	_ = v11200
	var v11201 int32
	_ = v11201
	var v11202 int32
	_ = v11202
	var v11203 int32
	_ = v11203
	var v11204 int32
	_ = v11204
	var v11205 int32
	_ = v11205
	var v11206 int32
	_ = v11206
	var v11208 int32
	_ = v11208
	var v11210 int32
	_ = v11210
	var v11211 int32
	_ = v11211
	var v11212 int32
	_ = v11212
	var v11214 int32
	_ = v11214
	var v11216 int32
	_ = v11216
	var v11217 int32
	_ = v11217
	var v11219 int32
	_ = v11219
	var v11223 int32
	_ = v11223
	var v11224 int32
	_ = v11224
	var v11230 int32
	_ = v11230
	var v11232 int32
	_ = v11232
	var v11233 int32
	_ = v11233
	var v11238 int32
	_ = v11238
	var v11239 int32
	_ = v11239
	var v11242 int32
	_ = v11242
	var v11243 int32
	_ = v11243
	var v11244 int32
	_ = v11244
	var v11246 int32
	_ = v11246
	var v11247 int32
	_ = v11247
	var v11248 int32
	_ = v11248
	var v11250 int32
	_ = v11250
	var v11253 int32
	_ = v11253
	var v11254 int32
	_ = v11254
	var v11256 int32
	_ = v11256
	var v11257 int32
	_ = v11257
	var v11258 int32
	_ = v11258
	var v11259 int32
	_ = v11259
	var v11260 int32
	_ = v11260
	var v11267 int32
	_ = v11267
	var v11270 int32
	_ = v11270
	var v11279 int32
	_ = v11279
	var v11308 int32
	_ = v11308
	var v11310 int32
	_ = v11310
	var v11314 int32
	_ = v11314
	var v11315 int32
	_ = v11315
	var v11316 int32
	_ = v11316
	var v11319 int32
	_ = v11319
	var v11320 int32
	_ = v11320
	var v11321 int32
	_ = v11321
	var v11322 int32
	_ = v11322
	var v11323 int32
	_ = v11323
	var v11324 int32
	_ = v11324
	var v11326 int32
	_ = v11326
	var v11329 int32
	_ = v11329
	var v11330 int32
	_ = v11330
	var v11331 int32
	_ = v11331
	var v11332 int32
	_ = v11332
	var v11341 int32
	_ = v11341
	var v11342 int32
	_ = v11342
	var v11344 int32
	_ = v11344
	var v11347 int32
	_ = v11347
	var v11348 int32
	_ = v11348
	var v11353 int32
	_ = v11353
	var v11360 int32
	_ = v11360
	var v11362 int32
	_ = v11362
	var v11364 int32
	_ = v11364
	var v11367 int32
	_ = v11367
	var v11369 int32
	_ = v11369
	var v11371 int32
	_ = v11371
	var v11376 int32
	_ = v11376
	var v11385 int32
	_ = v11385
	var v11388 int32
	_ = v11388
	var v11397 int32
	_ = v11397
	var v11398 int32
	_ = v11398
	var v11400 int32
	_ = v11400
	var v11403 int32
	_ = v11403
	var v11404 int32
	_ = v11404
	var v11409 int32
	_ = v11409
	var v11416 int32
	_ = v11416
	var v11418 int32
	_ = v11418
	var v11420 int32
	_ = v11420
	var v11421 int32
	_ = v11421
	var v11423 int32
	_ = v11423
	var v11425 int32
	_ = v11425
	var v11429 int32
	_ = v11429
	var v11435 int32
	_ = v11435
	var v11436 int32
	_ = v11436
	var v11437 int32
	_ = v11437
	var v11439 int32
	_ = v11439
	var v11440 int32
	_ = v11440
	var v11443 int32
	_ = v11443
	var v11444 int32
	_ = v11444
	var v11445 int32
	_ = v11445
	var v11447 int32
	_ = v11447
	var v11448 int32
	_ = v11448
	var v11449 int32
	_ = v11449
	var v11459 int32
	_ = v11459
	var v11460 int32
	_ = v11460
	var v11463 int32
	_ = v11463
	var v11467 int32
	_ = v11467
	var v11470 int32
	_ = v11470
	var v11472 int32
	_ = v11472
	var v11475 int32
	_ = v11475
	var v11482 int32
	_ = v11482
	var v11484 int32
	_ = v11484
	var v11492 int32
	_ = v11492
	var v11493 int32
	_ = v11493
	var v11506 int32
	_ = v11506
	var v11518 int32
	_ = v11518
	var v11534 int32
	_ = v11534
	var v11553 int32
	_ = v11553
	var v11555 int32
	_ = v11555
	var v11559 int32
	_ = v11559
	var v11562 int32
	_ = v11562
	var v11563 int32
	_ = v11563
	var v11564 int32
	_ = v11564
	var v11566 int32
	_ = v11566
	var v11567 int32
	_ = v11567
	var v11574 int32
	_ = v11574
	var v11576 int32
	_ = v11576
	var v11577 int32
	_ = v11577
	var v11580 int32
	_ = v11580
	var v11584 int32
	_ = v11584
	var v11587 int32
	_ = v11587
	var v11589 int32
	_ = v11589
	var v11592 int32
	_ = v11592
	var v11599 int32
	_ = v11599
	var v11601 int32
	_ = v11601
	var v11609 int32
	_ = v11609
	var v11610 int32
	_ = v11610
	var v11623 int32
	_ = v11623
	var v11635 int32
	_ = v11635
	var v11670 int32
	_ = v11670
	var v11671 int32
	_ = v11671
	var v11672 int32
	_ = v11672
	var v11673 int32
	_ = v11673
	var v11674 int32
	_ = v11674
	var v11676 int32
	_ = v11676
	var v11677 int32
	_ = v11677
	var v11681 int32
	_ = v11681
	var v11682 int32
	_ = v11682
	var v11683 int32
	_ = v11683
	var v11684 int32
	_ = v11684
	var v11686 int32
	_ = v11686
	var v11687 int32
	_ = v11687
	var v11688 int32
	_ = v11688
	var v11696 int32
	_ = v11696
	var v11736 int32
	_ = v11736
	var v11737 int32
	_ = v11737
	var v11741 int32
	_ = v11741
	var v11744 int32
	_ = v11744
	var v11753 int32
	_ = v11753
	var v11799 int32
	_ = v11799
	var v11872 int32
	_ = v11872
	var v11875 int32
	_ = v11875
	var v11876 int32
	_ = v11876
	var v11877 int32
	_ = v11877
	var v11878 int32
	_ = v11878
	var v11884 int32
	_ = v11884
	var v11887 int32
	_ = v11887
	var v11892 int32
	_ = v11892
	var v11925 int32
	_ = v11925
	var v11929 int32
	_ = v11929
	var v11930 int32
	_ = v11930
	var v11931 int32
	_ = v11931
	var v11934 int32
	_ = v11934
	var v11935 int32
	_ = v11935
	var v11936 int32
	_ = v11936
	var v11937 int32
	_ = v11937
	var v11938 int32
	_ = v11938
	var v11940 int32
	_ = v11940
	var v11942 int32
	_ = v11942
	var v11943 int32
	_ = v11943
	var v11944 int32
	_ = v11944
	var v11945 int32
	_ = v11945
	var v11946 int32
	_ = v11946
	var v11950 int32
	_ = v11950
	var v11954 int32
	_ = v11954
	var v11958 int32
	_ = v11958
	var v11959 int32
	_ = v11959
	var v11960 int32
	_ = v11960
	var v11961 int32
	_ = v11961
	var v11964 int32
	_ = v11964
	var v11965 int32
	_ = v11965
	var v11966 int32
	_ = v11966
	var v11968 int32
	_ = v11968
	var v11971 int32
	_ = v11971
	var v11972 int32
	_ = v11972
	var v11980 int32
	_ = v11980
	var v11985 int32
	_ = v11985
	var v12018 int32
	_ = v12018
	var v12022 int32
	_ = v12022
	var v12023 int32
	_ = v12023
	var v12032 int32
	_ = v12032
	var v12068 int32
	_ = v12068
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
	var v12085 int32
	_ = v12085
	var v12087 float64
	_ = v12087
	var v12089 float64
	_ = v12089
	var v12091 float64
	_ = v12091
	var v12093 int32
	_ = v12093
	var v12094 int32
	_ = v12094
	var v12096 int32
	_ = v12096
	var v12098 int32
	_ = v12098
	var v12100 int32
	_ = v12100
	var v12101 int32
	_ = v12101
	var v12106 int32
	_ = v12106
	var v12110 int32
	_ = v12110
	var v12116 int32
	_ = v12116
	var v12117 int32
	_ = v12117
	var v12119 int32
	_ = v12119
	var v12158 int32
	_ = v12158
	var v12162 int32
	_ = v12162
	var v12163 int32
	_ = v12163
	var v12164 int32
	_ = v12164
	var v12165 int32
	_ = v12165
	var v12166 int32
	_ = v12166
	var v12168 int32
	_ = v12168
	var v12170 int32
	_ = v12170
	var v12171 int32
	_ = v12171
	var v12175 int32
	_ = v12175
	var v12179 int32
	_ = v12179
	var v12180 int32
	_ = v12180
	var v12182 int32
	_ = v12182
	var v12183 int32
	_ = v12183
	var v12187 int32
	_ = v12187
	var v12229 int32
	_ = v12229
	var v12230 int32
	_ = v12230
	var v12231 int32
	_ = v12231
	var v12235 int32
	_ = v12235
	var v12236 int32
	_ = v12236
	var v12237 int32
	_ = v12237
	var v12239 int32
	_ = v12239
	var v12240 int32
	_ = v12240
	var v12241 int32
	_ = v12241
	var v12242 int32
	_ = v12242
	var v12243 int32
	_ = v12243
	var v12246 int32
	_ = v12246
	var v12250 int32
	_ = v12250
	var v12251 int32
	_ = v12251
	var v12257 int32
	_ = v12257
	var v12259 int32
	_ = v12259
	var v12260 int32
	_ = v12260
	var v12265 int32
	_ = v12265
	var v12266 int32
	_ = v12266
	var v12267 int32
	_ = v12267
	var v12268 int32
	_ = v12268
	var v12269 int32
	_ = v12269
	var v12270 int32
	_ = v12270
	var v12272 int32
	_ = v12272
	var v12273 int32
	_ = v12273
	var v12274 int32
	_ = v12274
	var v12276 int32
	_ = v12276
	var v12277 int32
	_ = v12277
	var v12278 int32
	_ = v12278
	var v12280 int32
	_ = v12280
	var v12281 int32
	_ = v12281
	var v12282 int32
	_ = v12282
	var v12283 int32
	_ = v12283
	var v12284 int32
	_ = v12284
	var v12285 int32
	_ = v12285
	var v12289 int32
	_ = v12289
	var v12290 int32
	_ = v12290
	var v12293 int32
	_ = v12293
	var v12294 int32
	_ = v12294
	var v12295 int32
	_ = v12295
	var v12296 int32
	_ = v12296
	var v12297 int32
	_ = v12297
	var v12298 int32
	_ = v12298
	var v12301 int32
	_ = v12301
	var v12302 int32
	_ = v12302
	var v12303 int32
	_ = v12303
	var v12304 int32
	_ = v12304
	var v12307 int32
	_ = v12307
	var v12308 int32
	_ = v12308
	var v12312 int32
	_ = v12312
	var v12313 int32
	_ = v12313
	var v12314 int32
	_ = v12314
	var v12315 int32
	_ = v12315
	var v12316 int32
	_ = v12316
	var v12322 int32
	_ = v12322
	var v12323 int32
	_ = v12323
	var v12324 int32
	_ = v12324
	var v12325 int32
	_ = v12325
	var v12330 int32
	_ = v12330
	var v12331 int32
	_ = v12331
	var v12332 int32
	_ = v12332
	var v12333 int32
	_ = v12333
	var v12334 int32
	_ = v12334
	var v12336 int32
	_ = v12336
	var v12343 int32
	_ = v12343
	var v12345 int32
	_ = v12345
	var v12347 int32
	_ = v12347
	var v12350 int32
	_ = v12350
	var v12353 int32
	_ = v12353
	var v12383 int32
	_ = v12383
	var v12387 int32
	_ = v12387
	var v12388 int32
	_ = v12388
	var v12389 int32
	_ = v12389
	var v12390 int32
	_ = v12390
	var v12391 int32
	_ = v12391
	var v12392 int32
	_ = v12392
	var v12393 int32
	_ = v12393
	var v12394 int32
	_ = v12394
	var v12395 int32
	_ = v12395
	var v12396 int32
	_ = v12396
	var v12397 int32
	_ = v12397
	var v12398 int32
	_ = v12398
	var v12399 int32
	_ = v12399
	var v12400 int32
	_ = v12400
	var v12401 int32
	_ = v12401
	var v12402 int32
	_ = v12402
	var v12403 int32
	_ = v12403
	var v12405 int32
	_ = v12405
	var v12406 int32
	_ = v12406
	var v12414 int32
	_ = v12414
	var v12416 int32
	_ = v12416
	var v12419 int32
	_ = v12419
	var v12421 int32
	_ = v12421
	var v12422 int32
	_ = v12422
	var v12425 int32
	_ = v12425
	var v12427 int32
	_ = v12427
	var v12453 int32
	_ = v12453
	var v12454 int32
	_ = v12454
	var v12457 int32
	_ = v12457
	var v12459 int32
	_ = v12459
	var v12464 int32
	_ = v12464
	var v12470 int32
	_ = v12470
	var v12472 float64
	_ = v12472
	var v12474 float64
	_ = v12474
	var v12476 float64
	_ = v12476
	var v12478 int32
	_ = v12478
	var v12482 int32
	_ = v12482
	var v12485 int32
	_ = v12485
	var v12486 int32
	_ = v12486
	var v12488 float64
	_ = v12488
	var v12490 int32
	_ = v12490
	var v12491 int32
	_ = v12491
	var v12492 int32
	_ = v12492
	var v12493 int32
	_ = v12493
	var v12495 int32
	_ = v12495
	var v12496 int32
	_ = v12496
	var v12510 int32
	_ = v12510
	var v12512 float64
	_ = v12512
	var v12514 float64
	_ = v12514
	var v12516 float64
	_ = v12516
	var v12518 int32
	_ = v12518
	var v12519 int32
	_ = v12519
	var v12521 int32
	_ = v12521
	var v12523 int32
	_ = v12523
	var v12528 int32
	_ = v12528
	var v12569 int32
	_ = v12569
	var v12572 int32
	_ = v12572
	var v12573 int32
	_ = v12573
	var v12574 int32
	_ = v12574
	var v12576 int32
	_ = v12576
	var v12577 int32
	_ = v12577
	var v12580 int32
	_ = v12580
	var v12581 int32
	_ = v12581
	var v12585 int32
	_ = v12585
	var v12589 int32
	_ = v12589
	v4 = int32(0)
	v45 = m.G0
	v47 = v45 - int32(144)
	m.G0 = v47
	F_check_stack_depth(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v53 - int32(331) {
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
	m.G0 = v12589 + int32(144)
	return v12585
L4:
	;
	v12569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)))
	if v12569 != int32(1) {
		v12585 = v12528
		v12589 = v47
		goto L3
	} else {
		goto L1775
	}
L5:
	;
	v12100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12101 = *(*int32)(unsafe.Add(mBase, uint32(v12100)+4))
	if v12101 == int32(0) {
		goto L1713
	} else {
		goto L1714
	}
L6:
	;
	v11062 = int32(0)
	v11063 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11064 = *(*int32)(unsafe.Add(mBase, uint32(v11063)+4))
	if v11064 == v11062 {
		v11147 = v11062
		goto L1558
	} else {
		goto L1559
	}
L7:
	;
	v10383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10384 = *(*int32)(unsafe.Add(mBase, uint32(v10383)+4))
	if v10384 == int32(0) {
		goto L1449
	} else {
		goto L1450
	}
L8:
	;
	v9793 = F_palloc0(m, int32(104))
	mBase = m.M
	v9794 = m.ExcPending
	if v9794 != 0 {
		goto L1
	} else {
		goto L1361
	}
L9:
	;
	v8861 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v8861 - int32(292) {
	case 0:
		goto L1243
	default:
		goto L1242
	case 9:
		goto L1245
	case 19:
		goto L1244
	}
L10:
	;
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8706 = F_create_plan_recurse(m, l0, v8704, int32(0))
	mBase = m.M
	v8707 = m.ExcPending
	if v8707 != 0 {
		goto L1
	} else {
		goto L1225
	}
L11:
	;
	v8672 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8675 = F_create_plan_recurse(m, l0, v8672, l2|int32(2))
	mBase = m.M
	v8676 = m.ExcPending
	if v8676 != 0 {
		goto L1
	} else {
		goto L1223
	}
L12:
	;
	v8510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8513 = F_create_plan_recurse(m, l0, v8510, l2|int32(2))
	mBase = m.M
	v8514 = m.ExcPending
	if v8514 != 0 {
		goto L1
	} else {
		goto L1194
	}
L13:
	;
	v6420 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6420 == int32(306) {
		goto L935
	} else {
		goto L936
	}
L14:
	;
	v6249 = int32(1)
	v6250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6252 = F_create_plan_recurse(m, l0, v6250, v6249)
	mBase = m.M
	v6253 = m.ExcPending
	if v6253 != 0 {
		goto L1
	} else {
		goto L917
	}
L15:
	;
	v6169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6172 = F_create_plan_recurse(m, l0, v6169, l2|int32(2))
	mBase = m.M
	v6173 = m.ExcPending
	if v6173 != 0 {
		goto L1
	} else {
		goto L911
	}
L16:
	;
	v6094 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6097 = F_create_plan_recurse(m, l0, v6094, l2|int32(2))
	mBase = m.M
	v6098 = m.ExcPending
	if v6098 != 0 {
		goto L1
	} else {
		goto L905
	}
L17:
	;
	v5918 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v5920 = F_create_plan_recurse(m, l0, v5918, int32(4))
	mBase = m.M
	v5921 = m.ExcPending
	if v5921 != 0 {
		goto L1
	} else {
		goto L881
	}
L18:
	;
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4313 == int32(310) {
		goto L675
	} else {
		goto L676
	}
L19:
	;
	v3857 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+12))
	if v3858 != 0 {
		goto L630
	} else {
		goto L631
	}
L20:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v3524)+4))
	if v3525 == int32(0) {
		v3612 = v4
		goto L581
	} else {
		goto L582
	}
L21:
	;
	v3158 = int32(1)
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3161 = F_create_plan_recurse(m, l0, v3159, v3158)
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		goto L1
	} else {
		goto L533
	}
L22:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3127 = F_create_plan_recurse(m, l0, v3126, l2)
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L1
	} else {
		goto L531
	}
L23:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1443 = F_create_plan_recurse(m, l0, v1441, int32(1))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L209
	}
L24:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1268 = F_create_plan_recurse(m, l0, v1267, l2)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L192
	}
L25:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1089 = int32(0)
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1091 == v1089 {
		v1174 = v1089
		goto L173
	} else {
		goto L174
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L170
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v59 == int32(0) {
		v158 = v4
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v56 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v12585 = v56
	v12589 = v47
	goto L3
L30:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v191 = F_create_plan_recurse(m, l0, v185, base.B2i32(v186 != int32(0))<<(uint(int32(1))%32))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L45
	}
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= int32(0) {
		v158 = v4
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v72 = int32(1)
	v74 = v4
	v87 = v4
	goto L33
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v74<<(uint(int32(2))%32))))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v118 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v158 = v134
	goto L30
L35:
	;
	v119 = F_replace_nestloop_params_mutator(m, v117, l0)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v121 = v117
	goto L37
L37:
	;
	v123 = int32(0)
	v125 = F_makeTargetEntry(m, v121, base.I32_extend16_s(v72), v123, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v121 = v119
	goto L37
L39:
	;
	if v66 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v66-int32(4)+v72<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v130
	goto L42
L41:
	;
	goto L42
L42:
	;
	v134 = F_lappend(m, v87, v125)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v137 = v74 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v137 < v138 {
		v72 = v72 + int32(1)
		v74 = v137
		v87 = v134
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v199 = F_create_plan_recurse(m, l0, v193, base.B2i32(v194 != int32(0))<<(uint(int32(1))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v202 = F_order_qual_clauses(m, l0, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v202
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v206)%32)&int32(174) != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v226 = F_get_actual_clauses(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	F_extract_actual_join_clauses(m, v202, v211, v47+int32(52), v47+int32(48))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v219 = F_extract_actual_clauses(m, v202, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L48
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v219
	goto L48
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v229 = F_list_difference(m, v228, v226)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v229
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v232 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v233 = F_replace_nestloop_params_mutator(m, v229, l0)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v241 = l1 + int32(104)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v246 = F_get_switched_clauses(m, v242, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v233
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v237 = F_replace_nestloop_params_mutator(m, v236, l0)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v237
	goto L58
L61:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v248 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v425 != 0 {
		goto L76
	} else {
		goto L77
	}
L63:
	;
	v250 = l1 + int32(100)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _consts[491])))
	if v254 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v414 = v411 - int32(-64)
	v419 = v191
	goto L62
L66:
	;
	v323 = int32(0)
	v335 = F_prepare_sort_from_pathkeys(m, v191, v248, v252, v323, v323, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L72
	}
L67:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v257 <= int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v260 = int32(0)
	v272 = F_prepare_sort_from_pathkeys(m, v191, v248, v252, v260, v260, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v280 = F_palloc0(m, int32(104))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(363)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v272)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+96)) = v257
	v286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+56)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v280)+52)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v280)+48)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v280)+44)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v280)+88)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v280)+84)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v280)+80)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v280)+76)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v280)+72)) = v274
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v301 = *(*float64)(unsafe.Add(mBase, uint32(v272)+8))
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v272)+16))
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v272)+24))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v272)+32))
	v306 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	F_cost_incremental_sort(m, v47+int32(56), l0, v299, v257, v300, v301, v302, v303, v304, v306, float64(-1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v310 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v280)+8)) = v310
	v312 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v280)+16)) = v312
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v272)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v280)+24)) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v272)+32))
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+36)) = uint8(v317)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+32)) = v316
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+37)) = uint8(v320)
	v414 = v250
	v419 = v280
	goto L62
L72:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v343 = F_palloc0(m, int32(96))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = int32(362)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v335)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+44)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v350 = int32(4126313)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v343)+88)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v343)+84)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v343)+80)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v343)+76)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v343)+72)) = v337
	v357 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+56)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v343)+52)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v343)+48)) = v357
	v362 = int32(1)
	v364 = v349 + (v351 ^ v362)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+4)) = v364
	v367 = v47 + int32(56)
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v335)+16))
	v369 = *(*float64)(unsafe.Add(mBase, uint32(v335)+24))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v335)+32))
	v373 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v376 = m.G0
	v377 = int32(16)
	v378 = v376 - v377
	m.G0 = v378
	F_cost_tuplesort(m, v378+int32(8), v378, v369, v370, float64(0), v373, float64(-1))
	mBase = m.M
	v383 = *(*float64)(unsafe.Add(mBase, uint32(v378)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v367)+32)) = v369
	v386 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	v387 = base.F64_add(v368, v383)
	*(*float64)(unsafe.Add(mBase, uint32(v367)+48)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v367)+40)) = v364 + (v386 ^ v362)
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v378)))
	*(*float64)(unsafe.Add(mBase, uint32(v367)+56)) = base.F64_add(v387, v393)
	m.G0 = v378 + v377
	goto L74
L74:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v343)+8)) = v399
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v343)+16)) = v401
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v335)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v343)+24)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v335)+32))
	v406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v343)+36)) = uint8(v406)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+32)) = v405
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v343)+37)) = uint8(v409)
	v414 = v250
	v419 = v343
	goto L62
L75:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+113)))
	if v530 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	v428 = int32(0)
	v440 = F_prepare_sort_from_pathkeys(m, v199, v425, v427, v428, v428, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v520 = v199
	v522 = v516 - int32(-64)
	goto L75
L79:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v448 = F_palloc0(m, int32(96))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = int32(362)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v440)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+44)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v455 = int32(4126313)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+88)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v448)+84)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v448)+80)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v448)+76)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v448)+72)) = v442
	v462 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v448)+56)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v448)+52)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v448)+48)) = v462
	v467 = int32(1)
	v469 = v454 + (v456 ^ v467)
	*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v469
	v472 = v47 + int32(56)
	v473 = *(*float64)(unsafe.Add(mBase, uint32(v440)+16))
	v474 = *(*float64)(unsafe.Add(mBase, uint32(v440)+24))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v440)+32))
	v478 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v481 = m.G0
	v482 = int32(16)
	v483 = v481 - v482
	m.G0 = v483
	F_cost_tuplesort(m, v483+int32(8), v483, v474, v475, float64(0), v478, float64(-1))
	mBase = m.M
	v488 = *(*float64)(unsafe.Add(mBase, uint32(v483)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v472)+32)) = v474
	v491 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	v492 = base.F64_add(v473, v488)
	*(*float64)(unsafe.Add(mBase, uint32(v472)+48)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v472)+40)) = v469 + (v491 ^ v467)
	v498 = *(*float64)(unsafe.Add(mBase, uint32(v483)))
	*(*float64)(unsafe.Add(mBase, uint32(v472)+56)) = base.F64_add(v492, v498)
	m.G0 = v483 + v482
	goto L81
L81:
	;
	v504 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v448)+8)) = v504
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v448)+16)) = v506
	v508 = *(*float64)(unsafe.Add(mBase, uint32(v440)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v448)+24)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v440)+32))
	v511 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+36)) = uint8(v511)
	*(*int32)(unsafe.Add(mBase, uint32(v448)+32)) = v510
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+37)) = uint8(v514)
	v520 = v448
	v522 = v241
	goto L75
L82:
	;
	v568 = int32(0)
	if v246 != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v565 = v520
	goto L82
L84:
	;
	goto L85
L85:
	;
	v534 = F_palloc0(m, int32(72))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(360)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v520)+44))
	v539 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+56)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v534)+52)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v534)+48)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v534)+44)) = v538
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = v545
	v547 = *(*float64)(unsafe.Add(mBase, uint32(v520)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+8)) = v547
	v549 = *(*float64)(unsafe.Add(mBase, uint32(v520)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+16)) = v549
	v551 = *(*float64)(unsafe.Add(mBase, uint32(v520)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+24)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v520)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+36)) = uint8(v539)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+32)) = v553
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+37)) = uint8(v557)
	v560 = *(*float64)(unsafe.Add(mBase, _consts[484]))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+16)) = base.F64_add(v549, base.F64_mul(v551, v560))
	v565 = v534
	goto L82
L87:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v571 = v570
	goto L89
L88:
	;
	v571 = v568
	goto L89
L89:
	;
	v573 = v571 << (uint(int32(2)) % 32)
	v574 = F_palloc(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v576 = F_palloc(m, v573)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v578 = F_palloc(m, v571)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v580 = F_palloc(m, v571)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v424 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v583 = v582
	goto L96
L95:
	;
	v583 = v568
	goto L96
L96:
	;
	if v529 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v529)+12))
	v586 = v584
	goto L99
L98:
	;
	v586 = int32(0)
	goto L99
L99:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	if v587 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L167
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L164
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L161
	}
L103:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+112)))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v1003 = F_palloc0(m, int32(112))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L160
	}
L104:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v590 <= int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v593 = int32(0)
	v600 = v593
	v603 = v586
	v606 = v593
	v610 = v4
	v612 = v583
	goto L106
L106:
	;
	v640 = v606 << (uint(int32(2)) % 32)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v640+v641)))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+120)))
	if v646 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L103
L108:
	;
	v647 = int32(104)
	goto L110
L109:
	;
	v647 = int32(100)
	goto L110
L110:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v643+v647)))
	if v646 != 0 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v672)+8))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v880)+8))
	if v921 != v922 {
		goto L101
	} else {
		goto L152
	}
L112:
	;
	if v529 == int32(0) {
		v790 = v719
		v800 = v720
		goto L140
	} else {
		goto L141
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L136
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	v652 = int32(100)
	goto L117
L116:
	;
	v652 = int32(104)
	goto L117
L117:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v643+v652)))
	if v610 != v654 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v612 == int32(0) {
		goto L114
	} else {
		goto L121
	}
L119:
	;
	v672 = v600
	v673 = v610
	v674 = v612
	goto L120
L120:
	;
	if v603 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	if v654 != v659 {
		goto L113
	} else {
		goto L122
	}
L122:
	;
	v662 = v612 + int32(4)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if base.Ui32(v662) < base.Ui32(v664+v665<<(uint(int32(2))%32)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v670 = v662
	goto L125
L124:
	;
	v670 = int32(0)
	goto L125
L125:
	;
	v672 = v658
	v673 = v654
	v674 = v670
	goto L120
L126:
	;
	v677 = int32(0)
	v719 = v677
	v720 = v677
	goto L112
L127:
	;
	goto L128
L128:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	if v649 != v680 {
		v719 = v679
		v720 = v680
		goto L112
	} else {
		goto L129
	}
L129:
	;
	v683 = v603 + int32(4)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v529)+12))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	if base.Ui32(v683) < base.Ui32(v685+v686<<(uint(int32(2))%32)) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v691 = v683
	goto L132
L131:
	;
	v691 = int32(0)
	goto L132
L132:
	;
	v880 = v679
	v884 = v691
	v920 = int32(1)
	goto L111
L133:
	;
	F_errmsg_internal(m, int32(161509), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(498183), int32(4735), int32(283791))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(161509), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(498183), int32(4740), int32(283791))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v880 = v835
	v884 = v603
	v920 = int32(0)
	goto L111
L140:
	;
	if v649 != v800 {
		goto L102
	} else {
		goto L151
	}
L141:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	if v723 <= int32(0) {
		v790 = v719
		v800 = v720
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v726 = int32(0)
	if v726 < v723 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v729 = v723
	goto L145
L144:
	;
	v729 = v726
	goto L145
L145:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v529)+12))
	v735 = int32(0)
	v736 = v719
	v746 = v720
	goto L146
L146:
	;
	v778 = v730 + v735<<(uint(int32(2))%32)
	if v778 == v603 {
		v790 = v736
		v800 = v746
		goto L140
	} else {
		goto L148
	}
L147:
	;
	v790 = v780
	v800 = v781
	goto L140
L148:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	if v649 == v781 {
		v835 = v780
		goto L139
	} else {
		goto L149
	}
L149:
	;
	v784 = v735 + int32(1)
	if v784 != v729 {
		v735 = v784
		v736 = v780
		v746 = v781
		goto L146
	} else {
		goto L150
	}
L150:
	;
	goto L147
L151:
	;
	v835 = v790
	goto L139
L152:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+8))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v880)+4))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v926)+8))
	if v925 != v927 {
		goto L101
	} else {
		goto L153
	}
L153:
	;
	if v920 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v880)+12))
	if v929 != v930 {
		goto L100
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640+v574))) = v921
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v938)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v640+v576))) = v939
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v606+v578))) = uint8(base.B2i32(v942 == int32(5)))
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v606+v580))) = uint8(v947)
	v950 = v606 + int32(1)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v950 < v951 {
		v600 = v672
		v603 = v884
		v606 = v950
		v610 = v673
		v612 = v674
		goto L106
	} else {
		goto L159
	}
L157:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+16)))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+16)))
	if v932 != v933 {
		goto L100
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	goto L107
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+108)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+104)) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+100)) = v576
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+96)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+92)) = v246
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+88)) = uint8(v999)
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+56)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+52)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+48)) = v1001
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+44)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v1003))) = int32(358)
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+80)) = v1000
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+76)) = uint8(v998)
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+72)) = v997
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+4)) = v1020
	v1022 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1003)+8)) = v1022
	v1024 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1003)+16)) = v1024
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1003)+24)) = v1026
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+32)) = v1029
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+36)) = uint8(v1031)
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+37)) = uint8(v1033)
	v12528 = v1003
	goto L4
L161:
	;
	F_errmsg_internal(m, int32(161550), int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(498183), int32(4784), int32(283791))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errmsg_internal(m, int32(276400), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(498183), int32(4803), int32(283791))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errmsg_internal(m, int32(276400), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(498183), int32(4807), int32(283791))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
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
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v1078
	F_errmsg_internal(m, int32(487371), v47)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(498183), int32(546), int32(361551))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
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
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1218 = F_create_plan_recurse(m, l0, v1216, int32(1))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L188
	}
L174:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1095 <= int32(0) {
		v1174 = v1089
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+8))
	v1103 = v1089
	v1104 = int32(1)
	v1106 = v4
	goto L176
L176:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+12))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1145+v1106<<(uint(int32(2))%32))))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1150 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1174 = v1166
	goto L173
L178:
	;
	v1151 = F_replace_nestloop_params_mutator(m, v1149, l0)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	v1153 = v1149
	goto L180
L180:
	;
	v1155 = int32(0)
	v1157 = F_makeTargetEntry(m, v1153, base.I32_extend16_s(v1104), v1155, v1155)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L182
	}
L181:
	;
	v1153 = v1151
	goto L180
L182:
	;
	if v1098 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1098-int32(4)+v1104<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+16)) = v1162
	goto L185
L184:
	;
	goto L185
L185:
	;
	v1166 = F_lappend(m, v1103, v1157)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1169 = v1106 + int32(1)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1169 < v1170 {
		v1103 = v1166
		v1104 = v1104 + int32(1)
		v1106 = v1169
		goto L176
	} else {
		goto L187
	}
L187:
	;
	goto L177
L188:
	;
	v1221 = F_palloc0(m, int32(104))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+44)) = v1174
	*(*int32)(unsafe.Add(mBase, uint32(v1221))) = int32(369)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+72)) = v1226
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+4)) = v1228
	v1230 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1221)+8)) = v1230
	v1232 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1221)+16)) = v1232
	v1234 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1221)+24)) = v1234
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1236)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+32)) = v1237
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+36)) = uint8(v1239)
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+37)) = uint8(v1241)
	v1243 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+76)) = v1243
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+8))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+8))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+84))
	v1261 = F_prepare_sort_from_pathkeys(m, v1218, v1088, v1248, v1249, int32(0), v1221+int32(80), v1221+int32(84), v1221+int32(88), v1221+int32(92), v1221+int32(96))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+52)) = v1261
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1264)+83)) = uint8(v1265)
	v12585 = v1221
	v12589 = v47
	goto L3
L192:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v1270 != int32(1) {
		v1366 = v4
		v1368 = v4
		v1370 = v4
		v1373 = v4
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1406 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1409 = F_palloc0(m, int32(104))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L208
	}
L194:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+124))
	if v1274 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	v1276 = v1275
	goto L197
L196:
	;
	v1276 = v4
	goto L197
L197:
	;
	v1279 = F_palloc(m, v1276<<(uint(int32(1))%32))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v1282 = v1276 << (uint(int32(2)) % 32)
	v1283 = F_palloc(m, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v1285 = F_palloc(m, v1282)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+124))
	if v1287 == int32(0) {
		v1366 = v4
		v1368 = v1285
		v1370 = v1283
		v1373 = v1279
		goto L193
	} else {
		goto L201
	}
L201:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+4))
	if v1290 <= int32(0) {
		v1366 = v4
		v1368 = v1285
		v1370 = v1283
		v1373 = v1279
		goto L193
	} else {
		goto L202
	}
L202:
	;
	v1297 = v4
	goto L203
L203:
	;
	v1341 = v1297 << (uint(int32(2)) % 32)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+12))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1341+v1342)))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+76))
	v1346 = F_get_sortgroupclause_tle(m, v1344, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L205
	}
L204:
	;
	v1366 = v1359
	v1368 = v1285
	v1370 = v1283
	v1373 = v1279
	goto L193
L205:
	;
	v1348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1346)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1279+v1297<<(uint(int32(1))%32)))) = uint16(v1348)
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1341+v1283))) = v1351
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+4))
	v1355 = F_exprCollation(m, v1354)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1341+v1285))) = v1355
	v1359 = v1297 + int32(1)
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+4))
	if v1359 < v1360 {
		v1297 = v1359
		goto L203
	} else {
		goto L207
	}
L207:
	;
	goto L204
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1409))) = int32(373)
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+96)) = v1368
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+92)) = v1370
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+88)) = v1373
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+84)) = v1366
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+80)) = v1407
	*(*int64)(unsafe.Add(mBase, uint32(v1409)+72)) = v1406
	v1420 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+56)) = v1420
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+52)) = v1268
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+48)) = v1420
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+44)) = v1413
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+4)) = v1426
	v1428 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1409)+8)) = v1428
	v1430 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1409)+16)) = v1430
	v1432 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1409)+24)) = v1432
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+32)) = v1435
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1409)+36)) = uint8(v1437)
	v1439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1409)+37)) = uint8(v1439)
	v12585 = v1409
	v12589 = v47
	goto L3
L209:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+44))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v1453 = int32(0)
	goto L211
L210:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+92)))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v1506 = F_palloc0(m, int32(168))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L222
	}
L211:
	;
	v1455 = int32(0)
	if v1445 == v1455 {
		v1465 = v1455
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if v1446 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+4))
	if v1459 <= v1453 {
		v1465 = int32(0)
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+12))
	v1465 = v1461 + v1453<<(uint(int32(2))%32)
	goto L213
L216:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1465)))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1475)))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1477)+12)) = v1479
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1477)+16)) = v1481
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1477)+20)) = v1483
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1477)+24)) = uint16(v1485)
	v1487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1477)+26)) = uint8(v1487)
	v1453 = v1453 + int32(1)
	goto L211
L217:
	;
	goto L210
L218:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+4))
	if v1468 <= v1453 {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	if v1465 == int32(0) {
		goto L217
	} else {
		goto L220
	}
L220:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+12))
	v1475 = v1472 + v1453<<(uint(int32(2))%32)
	if v1475 != 0 {
		goto L216
	} else {
		goto L221
	}
L221:
	;
	goto L217
L222:
	;
	v1508 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+56)) = v1508
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+52)) = v1443
	*(*int32)(unsafe.Add(mBase, uint32(v1506))) = int32(333)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+92)) = v1504
	*(*uint8)(unsafe.Add(mBase, uint32(v1506)+88)) = uint8(v1503)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+84)) = v1502
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+80)) = v1501
	*(*uint8)(unsafe.Add(mBase, uint32(v1506)+76)) = uint8(v1500)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+72)) = v1499
	*(*int64)(unsafe.Add(mBase, uint32(v1506)+44)) = int64(0)
	if v1494 == v1508 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+100)) = v1497
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+96)) = v1498
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+156)) = v2558
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2600)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+104)) = v2601
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+164)) = v1492
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+160)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+124)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+112)) = v1496
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+108)) = v2604
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+128)) = v1491
	if v1504 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L224:
	;
	v1523 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1506)+132)) = v1523
	*(*int64)(unsafe.Add(mBase, uint32(v1506)+148)) = v1523
	*(*int64)(unsafe.Add(mBase, uint32(v1506)+140)) = v1523
	v2556 = v1506
	v2558 = v4
	v2560 = v47
	goto L223
L225:
	;
	goto L226
L226:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+132)) = v1529
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+140)) = v1531
	if v1531 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+144)) = v1611
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+148)) = v1646
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+84))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+8))
	if v1650 != 0 {
		goto L241
	} else {
		goto L242
	}
L228:
	;
	v1540 = int32(1)
	v1548 = v4
	v1553 = int32(0)
	goto L233
L229:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+4))
	if int32(0) < v1534 {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1611 = v4
	goto L227
L232:
	;
	goto L231
L233:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+12))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1582+v1553<<(uint(int32(2))%32))))
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+26)))
	if v1587 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1611 = v1593
	goto L227
L235:
	;
	v1590 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1586)+8)))
	v1591 = F_lappend_int(m, v1548, v1590)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	v1593 = v1548
	goto L237
L237:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1586)+8)) = uint16(v1540)
	v1595 = int32(1)
	v1598 = v1553 + v1595
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+4))
	if v1598 < v1599 {
		v1540 = v1540 + v1595
		v1548 = v1593
		v1553 = v1598
		goto L233
	} else {
		goto L239
	}
L238:
	;
	v1593 = v1591
	goto L237
L239:
	;
	goto L234
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+136)) = v2548
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+152)) = v2550
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+32))
	v2556 = v1506
	v2558 = v2552
	v2560 = v47
	goto L223
L241:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+52))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1653)+12))
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+32))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1654+v1655<<(uint(int32(2))%32)-int32(4))))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+16))
	v1664 = F_table_open(m, v1662, int32(0))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+16))
	if v1651 != 0 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v2548 = int32(0)
	goto L240
L244:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+8))
	if v1666 == int32(0) {
		v1791 = v4
		v1792 = v4
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+16))
	if v1802 == int32(0) {
		v1823 = v4
		goto L263
	} else {
		goto L264
	}
L246:
	;
	v1669 = int32(0)
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+4))
	if v1670 <= v1669 {
		v1791 = v4
		v1792 = v4
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1675 = v1669
	v1706 = v4
	v1707 = v4
	goto L248
L248:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+12))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1717+v1675<<(uint(int32(2))%32))))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1721)+4))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1722)))
	if v1723 != int32(6) {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L259
	}
L250:
	;
	goto L249
L251:
	;
	v1739 = v1675 + int32(1)
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+4))
	if v1739 < v1740 {
		v1675 = v1739
		v1706 = v1736
		v1707 = v1737
		goto L248
	} else {
		goto L258
	}
L252:
	;
	v1726 = F_lappend(m, v1706, v1722)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v1728 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1722)+8)))
	if v1728 == int32(0) {
		goto L250
	} else {
		goto L256
	}
L255:
	;
	v1736 = v1726
	v1737 = v1707
	goto L251
L256:
	;
	v1733 = F_bms_add_member(m, v1707, v1728+int32(7))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1736 = v1706
	v1737 = v1733
	goto L251
L258:
	;
	v1791 = v1736
	v1792 = v1737
	goto L245
L259:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errmsg(m, int32(444800), int32(0))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(494969), int32(776), int32(157856))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	v1824 = F_RelationGetIndexList(m, v1664)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L1
	} else {
		goto L273
	}
L264:
	;
	v1805 = F_get_constraint_index(m, v1802)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	if v1805 != 0 {
		v1823 = v1805
		goto L263
	} else {
		goto L266
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(29023), int32(0))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(494969), int32(793), int32(157856))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L1
	} else {
		goto L385
	}
L272:
	;
	F_list_free(m, v1824)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L378
	}
L273:
	;
	if v1824 == int32(0) {
		v2446 = v4
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v1828 = int32(0)
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+4))
	if v1829 <= v1828 {
		v2446 = v4
		goto L272
	} else {
		goto L275
	}
L275:
	;
	v1844 = v1828
	v1855 = v4
	goto L276
L276:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+12))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1876+v1844<<(uint(int32(2))%32))))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+24))
	v1882 = F_index_open(m, v1880, v1881)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	v2446 = v2395
	goto L272
L278:
	;
	F_relation_close(m, v1882, int32(0))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L1
	} else {
		goto L376
	}
L279:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+192))
	v1885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+18)))
	if v1885 != int32(1) {
		v2395 = v1855
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1884)))
	if v1888 == v1823 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+15)))
	if v1890 == int32(1) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	if v1823 != 0 {
		v2395 = v1855
		goto L278
	} else {
		goto L292
	}
L284:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+4))
	if v1893 == int32(2) {
		goto L271
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1896 = F_lappend_oid(m, v1855, v1823)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	F_list_free(m, v1824)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	F_relation_close(m, v1882, int32(0))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_sequence_close(m, v1664, int32(0))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v2548 = v1896
	goto L240
L292:
	;
	v1906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+12)))
	if v1906 != int32(1) {
		v2395 = v1855
		goto L278
	} else {
		goto L293
	}
L293:
	;
	v1909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+15)))
	if v1909 != 0 {
		v2395 = v1855
		goto L278
	} else {
		goto L294
	}
L294:
	;
	v1910 = int32(0)
	v1912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1884)+10)))
	if v1910 < v1912 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1920 = v1910
	v1925 = v1912
	v1935 = v1910
	goto L298
L296:
	;
	v1995 = v1910
	goto L297
L297:
	;
	v2019 = int32(0)
	v2026 = base.B2i32(v1995|v1792 == v2019)
	if v1995 == v2019 {
		v2065 = v2026
		goto L306
	} else {
		goto L307
	}
L298:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+192))
	v1963 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1959+v1920<<(uint(int32(1))%32))+48)))
	if v1963 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1995 = v1970
	goto L297
L300:
	;
	v1966 = F_bms_add_member(m, v1935, v1963+int32(7))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	v1969 = v1925
	v1970 = v1935
	goto L302
L302:
	;
	v1972 = v1920 + int32(1)
	if v1972 < base.I32_extend16_s(v1969) {
		v1920 = v1972
		v1925 = v1969
		v1935 = v1970
		goto L298
	} else {
		goto L304
	}
L303:
	;
	v1968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1884)+10)))
	v1969 = v1968
	v1970 = v1966
	goto L302
L304:
	;
	goto L299
L305:
	;
	if v2065 == int32(0) {
		v2395 = v1855
		goto L278
	} else {
		goto L317
	}
L306:
	;
	goto L305
L307:
	;
	if v1792 == int32(0) {
		v2065 = v2026
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v1995)+4))
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v1792)+4))
	if v2032 != v2033 {
		v2065 = int32(0)
		goto L306
	} else {
		goto L309
	}
L309:
	;
	v2035 = int32(1)
	if v2032 <= v2035 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v2038 = v2035
	goto L312
L311:
	;
	v2038 = v2032
	goto L312
L312:
	;
	v2039 = int32(8)
	v2044 = int32(0)
	goto L313
L313:
	;
	v2052 = v2044 << (uint(int32(2)) % 32)
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v1995+v2039+v2052)))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2052+(v1792+v2039))))
	v2057 = base.B2i32(v2054 == v2056)
	if v2056 != v2054 {
		v2065 = v2057
		goto L306
	} else {
		goto L315
	}
L314:
	;
	v2065 = v2057
	goto L306
L315:
	;
	v2060 = v2044 + int32(1)
	if v2060 != v2038 {
		v2044 = v2060
		goto L313
	} else {
		goto L316
	}
L316:
	;
	goto L314
L317:
	;
	v2071 = F_RelationGetIndexExpressions(m, v1882)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v2074 = base.B2i32(v1655 == int32(1))
	if v1655 == int32(1) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+8))
	if v2080 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	if v2071 == int32(0) {
		goto L319
	} else {
		goto L321
	}
L321:
	;
	F_ChangeVarNodes(m, v2071, int32(1), v1655)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	goto L319
L323:
	;
	v2354 = F_list_difference(m, v2071, v1791)
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L1
	} else {
		goto L366
	}
L324:
	;
	v2083 = int32(0)
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+4))
	if v2084 <= v2083 {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v2119 = v2083
	goto L326
L326:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+12))
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2131+v2119<<(uint(int32(2))%32))))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2137 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L327:
	;
	goto L323
L328:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+4))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	if v2297 == int32(6) {
		goto L359
	} else {
		goto L360
	}
L329:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+52))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2151)))
	if v2152 <= int32(0) {
		v2395 = v1855
		goto L278
	} else {
		goto L338
	}
L330:
	;
	v2144 = F_get_opclass_family(m, v2136)
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L1
	} else {
		goto L336
	}
L331:
	;
	if v2136 == int32(0) {
		goto L328
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	if v2136 != 0 {
		goto L330
	} else {
		goto L335
	}
L334:
	;
	goto L330
L335:
	;
	v2142 = int32(0)
	v2149 = v2142
	v2150 = v2142
	goto L329
L336:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	v2147 = F_get_opclass_input_type(m, v2146)
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v2149 = v2144
	v2150 = v2147
	goto L329
L338:
	;
	v2156 = int32(1)
	v2163 = v2156
	v2178 = int32(0)
	v2194 = v2152
	v2195 = v2156
	goto L339
L339:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+192))
	v2203 = int32(1)
	v2204 = v2163 - v2203
	v2208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2202+v2204<<(uint(v2203)%32))+48)))
	v2210 = v2204 << (uint(int32(2)) % 32)
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+248))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2210+v2211)))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	if v2214 != 0 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	v2395 = v1855
	goto L278
L341:
	;
	v2249 = v2195 + int32(1)
	v2250 = base.I32_extend16_s(v2249)
	if v2250 <= v2244 {
		v2163 = v2250
		v2178 = v2178 + base.B2i32(v2208 != int32(0))
		v2194 = v2244
		v2195 = v2249
		goto L339
	} else {
		goto L358
	}
L342:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+208))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2215+v2210)))
	if v2149 != v2217 {
		v2244 = v2194
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2223 != v2213 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+212))
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2219+v2210)))
	if v2150 != v2221 {
		v2244 = v2194
		goto L341
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v2226 = v2223
	goto L349
L348:
	;
	v2226 = int32(0)
	goto L349
L349:
	;
	if v2226 != 0 {
		v2244 = v2194
		goto L341
	} else {
		goto L350
	}
L350:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+4))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	if v2228 == int32(6) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v2231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2227)+8)))
	if v2231 != v2208 {
		v2244 = v2194
		goto L341
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	if v2208 != 0 {
		v2244 = v2194
		goto L341
	} else {
		goto L355
	}
L354:
	;
	goto L328
L355:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v2071)+12))
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2233+(v2204-v2178)<<(uint(int32(2))%32))))
	v2239 = F_equal(m, v2227, v2238)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v2239 != 0 {
		goto L328
	} else {
		goto L357
	}
L357:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+52))
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2241)))
	v2244 = v2242
	goto L341
L358:
	;
	goto L340
L359:
	;
	v2307 = v2119 + int32(1)
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+4))
	if v2307 < v2308 {
		v2119 = v2307
		goto L326
	} else {
		goto L365
	}
L360:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2300 != 0 {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	if v2301 != 0 {
		goto L359
	} else {
		goto L362
	}
L362:
	;
	v2302 = F_list_member(m, v2071, v2296)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	if v2302 == int32(0) {
		v2395 = v1855
		goto L278
	} else {
		goto L364
	}
L364:
	;
	goto L359
L365:
	;
	goto L327
L366:
	;
	if v2354 != 0 {
		v2395 = v1855
		goto L278
	} else {
		goto L367
	}
L367:
	;
	v2356 = F_RelationGetIndexPredicate(m, v1882)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	if v1655 == int32(1) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+12))
	v2365 = F_predicate_implied_by(m, v2356, v2363, int32(0))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L1
	} else {
		goto L373
	}
L370:
	;
	if v2356 == int32(0) {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	F_ChangeVarNodes(m, v2356, int32(1), v1655)
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	goto L369
L373:
	;
	if v2365 == int32(0) {
		v2395 = v1855
		goto L278
	} else {
		goto L374
	}
L374:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v1884)))
	v2370 = F_lappend_oid(m, v1855, v2369)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v2395 = v2370
	goto L278
L376:
	;
	v2420 = v1844 + int32(1)
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+4))
	if v2420 < v2421 {
		v1844 = v2420
		v1855 = v2395
		goto L276
	} else {
		goto L377
	}
L377:
	;
	goto L277
L378:
	;
	F_sequence_close(m, v1664, int32(0))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	if v2446 != 0 {
		v2548 = v2446
		goto L240
	} else {
		goto L380
	}
L380:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	F_errmsg(m, int32(268729), int32(0))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(494969), int32(960), int32(157856))
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	F_errmsg(m, int32(119944), int32(0))
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(494969), int32(843), int32(157856))
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+120)) = v3082
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+116)) = v3073
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+4)) = v3111
	v3113 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2556)+8)) = v3113
	v3115 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2556)+16)) = v3115
	v3117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2556)+24)) = v3117
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v3119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2556)+32)) = v3120
	v3122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2556)+36)) = uint8(v3122)
	v3124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2556)+37)) = uint8(v3124)
	v12585 = v2556
	v12589 = v2560
	goto L3
L390:
	;
	v2613 = int32(0)
	v3073 = v2613
	v3082 = v2613
	goto L389
L391:
	;
	goto L392
L392:
	;
	v2615 = int32(0)
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+4))
	if v2616 <= v2615 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v3073 = int32(0)
	v3082 = v2615
	goto L389
L394:
	;
	goto L395
L395:
	;
	v2622 = int32(0)
	v2632 = v2622
	v2636 = v2622
	v2643 = v2622
	v2645 = v2615
	v2646 = v2622
	v2647 = v2622
	v2649 = v2622
	goto L396
L396:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+12))
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2672+v2632<<(uint(int32(2))%32))))
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v2677) <= base.Ui32(v2676) {
		goto L402
	} else {
		goto L403
	}
L397:
	;
	v3073 = v3059
	v3082 = v3053
	goto L389
L398:
	;
	v3059 = F_lappend(m, v2636, v3048)
	mBase = m.M
	v3060 = m.ExcPending
	if v3060 != 0 {
		goto L1
	} else {
		goto L529
	}
L399:
	;
	v3044 = F_bms_add_member(m, v2645, v2632)
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L1
	} else {
		goto L528
	}
L400:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L1
	} else {
		goto L524
	}
L401:
	;
	if v1499 != int32(5) {
		goto L413
	} else {
		goto L414
	}
L402:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2688 != 0 {
		goto L406
	} else {
		goto L407
	}
L403:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2679+v2676<<(uint(int32(2))%32))))
	if v2683 == int32(0) {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+168))
	v2716 = v2686
	goto L401
L405:
	;
	v2701 = int32(0)
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2700)))
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+12))
	if v2703 != 0 {
		v3048 = v2701
		v3052 = v2643
		v3053 = v2645
		v3054 = v2646
		v3055 = v2647
		v3057 = v2649
		goto L398
	} else {
		goto L409
	}
L406:
	;
	v2700 = v2688 + v2676<<(uint(int32(2))%32)
	goto L405
L407:
	;
	goto L408
L408:
	;
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+52))
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+12))
	v2700 = v2694 + v2676<<(uint(int32(2))%32) - int32(4)
	goto L405
L409:
	;
	v2704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2702)+21)))
	if v2704 != int32(102) {
		v3048 = v2701
		v3052 = v2643
		v3053 = v2645
		v3054 = v2646
		v3055 = v2647
		v3057 = v2649
		goto L398
	} else {
		goto L410
	}
L410:
	;
	v2708 = int32(*(*uint8)(unsafe.Add(mBase, _consts[492])))
	if v2708&int32(2) != 0 {
		goto L400
	} else {
		goto L411
	}
L411:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+16))
	v2712 = F_GetFdwRoutineByRelId(m, v2711)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v2716 = v2712
	goto L401
L413:
	;
	v2757 = int32(0)
	if v2716 == v2757 {
		v3048 = v2757
		v3052 = v2643
		v3053 = v2645
		v3054 = v2646
		v3055 = v2647
		v3057 = v2649
		goto L398
	} else {
		goto L426
	}
L414:
	;
	if v2716 == int32(0) {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2719 != 0 {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2731)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L1
	} else {
		goto L420
	}
L417:
	;
	v2731 = v2719 + v2676<<(uint(int32(2))%32)
	goto L416
L418:
	;
	goto L419
L419:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2723)+52))
	v2725 = *(*int32)(unsafe.Add(mBase, uint32(v2724)+12))
	v2731 = v2725 + v2676<<(uint(int32(2))%32) - int32(4)
	goto L416
L420:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2732)+16))
	v2741 = F_get_rel_name(m, v2740)
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2560)+16)) = v2741
	F_errmsg(m, int32(707846), v2560+int32(16))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	v2749 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2732)+21)))
	F_errdetail_relkind_not_supported(m, v2749)
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(498183), int32(7323), int32(392626))
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L1
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
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+88))
	if v2760 == int32(0) {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+44))
	if v3021 == int32(0) {
		v3048 = v2757
		v3052 = v3015
		v3053 = v2645
		v3054 = v3016
		v3055 = v3017
		v3057 = v3019
		goto L398
	} else {
		goto L522
	}
L428:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+92))
	if v2763 == int32(0) {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+96))
	if v2766 == int32(0) {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L430
	}
L430:
	;
	if v1497 != 0 {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L431
	}
L431:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+100))
	if v2769 == int32(0) {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L432
	}
L432:
	;
	v2772 = m.G0
	v2774 = v2772 - int32(16)
	m.G0 = v2774
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2776 != 0 {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	if v2817 != 0 {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L458
	}
L434:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2788)))
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2789)+16))
	v2792 = F_table_open(m, v2790, int32(0))
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L438
	}
L435:
	;
	v2788 = v2776 + v2676<<(uint(int32(2))%32)
	goto L434
L436:
	;
	goto L437
L437:
	;
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v2780)+52))
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+12))
	v2788 = v2782 + v2676<<(uint(int32(2))%32) - int32(4)
	goto L434
L438:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2792)+76))
	v2795 = int32(0)
	switch v1499 - int32(2) {
	case 0:
		goto L443
	case 1:
		goto L444
	case 2:
		goto L442
	case 3:
		v2817 = v2795
		goto L440
	default:
		goto L439
	}
L439:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L1
	} else {
		goto L455
	}
L440:
	;
	F_sequence_close(m, v2792, int32(0))
	mBase = m.M
	v2820 = m.ExcPending
	if v2820 != 0 {
		goto L1
	} else {
		goto L454
	}
L441:
	;
	v2817 = int32(1)
	goto L440
L442:
	;
	if v2794 == int32(0) {
		v2817 = v2795
		goto L440
	} else {
		goto L451
	}
L443:
	;
	if v2794 == int32(0) {
		v2817 = v2795
		goto L440
	} else {
		goto L448
	}
L444:
	;
	if v2794 == int32(0) {
		v2817 = v2795
		goto L440
	} else {
		goto L445
	}
L445:
	;
	v2800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+9)))
	if v2800 != 0 {
		goto L441
	} else {
		goto L446
	}
L446:
	;
	v2801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+8)))
	if v2801 == int32(1) {
		goto L441
	} else {
		goto L447
	}
L447:
	;
	v2817 = v2795
	goto L440
L448:
	;
	v2806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+14)))
	if v2806 != 0 {
		goto L441
	} else {
		goto L449
	}
L449:
	;
	v2807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+13)))
	if v2807 == int32(1) {
		goto L441
	} else {
		goto L450
	}
L450:
	;
	v2817 = v2795
	goto L440
L451:
	;
	v2812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+19)))
	if v2812 != 0 {
		goto L441
	} else {
		goto L452
	}
L452:
	;
	v2813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2794)+18)))
	if v2813 != int32(1) {
		v2817 = v2795
		goto L440
	} else {
		goto L453
	}
L453:
	;
	goto L441
L454:
	;
	m.G0 = v2774 + int32(16)
	goto L433
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2774))) = v1499
	F_errmsg_internal(m, int32(487717), v2774)
	mBase = m.M
	v2831 = m.ExcPending
	if v2831 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(494969), int32(2312), int32(134856))
	mBase = m.M
	v2836 = m.ExcPending
	if v2836 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2837 != 0 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v2850 = int32(0)
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v2849)))
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v2851)+16))
	v2854 = F_table_open(m, v2852, v2850)
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L463
	}
L460:
	;
	v2849 = v2837 + v2676<<(uint(int32(2))%32)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v2841)+52))
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2842)+12))
	v2849 = v2843 + v2676<<(uint(int32(2))%32) - int32(4)
	goto L459
L463:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2854)+52))
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(v2856)+16))
	if v2857 != 0 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2857)+17)))
	v2859 = v2858
	goto L466
L465:
	;
	v2859 = v2850
	goto L466
L466:
	;
	F_sequence_close(m, v2854, int32(0))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	if v2859&int32(1) != 0 {
		v3015 = v2643
		v3016 = v2646
		v3017 = v2647
		v3019 = v2649
		goto L427
	} else {
		goto L468
	}
L468:
	;
	if v2647&int32(1) == int32(0) {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	if v2649&int32(1) == int32(0) {
		goto L490
	} else {
		goto L491
	}
L470:
	;
	v2869 = int32(0)
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2870)+96))
	if v2871 == v2869 {
		v2893 = v2869
		goto L473
	} else {
		goto L474
	}
L471:
	;
	goto L472
L472:
	;
	v2898 = int32(1)
	if v2646&v2898 == int32(0) {
		goto L469
	} else {
		goto L486
	}
L473:
	;
	if v2893 == int32(0) {
		goto L469
	} else {
		goto L485
	}
L474:
	;
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v2871)))
	if v2874 != int32(61) {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	v2890 = F_expression_tree_walker_impl(m, v2871, int32(901), int32(0))
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L1
	} else {
		goto L484
	}
L476:
	;
	if v2874 != int32(6) {
		goto L475
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+4))
	v2893 = base.B2i32(v2885 == int32(0))
	goto L473
L479:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+28))
	if v2879 == int32(0) {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2871)+32))
	if v2883 != 0 {
		v2893 = int32(1)
		goto L473
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	v2893 = int32(0)
	goto L473
L483:
	;
	goto L482
L484:
	;
	v2893 = v2890
	goto L473
L485:
	;
	v2896 = int32(1)
	v3015 = v2643
	v3016 = v2896
	v3017 = v2896
	v3019 = v2649
	goto L427
L486:
	;
	v3015 = v2643
	v3016 = int32(1)
	v3017 = v2898
	v3019 = v2649
	goto L427
L487:
	;
	v3015 = v3011
	v3016 = v3009
	v3017 = v3010
	v3019 = int32(1)
	goto L427
L488:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+88))
	v3001 = m.T0[v3000].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v2556, v2676, v2632)
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L1
	} else {
		goto L520
	}
L489:
	;
	v3009 = int32(0)
	v3010 = v2992
	v3011 = int32(1)
	goto L487
L490:
	;
	v2911 = m.G0
	v2913 = v2911 - int32(16)
	m.G0 = v2913
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2915 != 0 {
		goto L495
	} else {
		goto L496
	}
L491:
	;
	goto L492
L492:
	;
	v2983 = int32(1)
	if v2643&v2983 == int32(0) {
		goto L488
	} else {
		goto L519
	}
L493:
	;
	if v2960&int32(1) == int32(0) {
		goto L488
	} else {
		goto L518
	}
L494:
	;
	v2928 = int32(0)
	v2929 = *(*int32)(unsafe.Add(mBase, uint32(v2927)))
	v2930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2929)+21)))
	if v2930 != int32(102) {
		goto L499
	} else {
		goto L500
	}
L495:
	;
	v2927 = v2915 + v1501<<(uint(int32(2))%32)
	goto L494
L496:
	;
	goto L497
L497:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2919)+52))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+12))
	v2927 = v2921 + v1501<<(uint(int32(2))%32) - int32(4)
	goto L494
L498:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L1
	} else {
		goto L515
	}
L499:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2929)+16))
	v2935 = F_table_open(m, v2933, int32(0))
	mBase = m.M
	v2936 = m.ExcPending
	if v2936 != 0 {
		goto L1
	} else {
		goto L502
	}
L500:
	;
	v2960 = v2928
	goto L501
L501:
	;
	m.G0 = v2913 + int32(16)
	goto L493
L502:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2935)+76))
	switch v1499 - int32(2) {
	case 0:
		goto L505
	case 1:
		goto L506
	case 2:
		goto L504
	case 3:
		v2955 = v2928
		goto L503
	default:
		goto L498
	}
L503:
	;
	F_sequence_close(m, v2935, int32(0))
	mBase = m.M
	v2958 = m.ExcPending
	if v2958 != 0 {
		goto L1
	} else {
		goto L514
	}
L504:
	;
	if v2937 == int32(0) {
		v2955 = v2928
		goto L503
	} else {
		goto L513
	}
L505:
	;
	if v2937 == int32(0) {
		v2955 = v2928
		goto L503
	} else {
		goto L508
	}
L506:
	;
	if v2937 == int32(0) {
		v2955 = v2928
		goto L503
	} else {
		goto L507
	}
L507:
	;
	v2942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2937)+25)))
	v2955 = v2942
	goto L503
L508:
	;
	v2945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2937)+26)))
	if v2945 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v2948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2937)+27)))
	if v2948 != int32(1) {
		v2955 = v2928
		goto L503
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v2955 = int32(1)
	goto L503
L512:
	;
	goto L511
L513:
	;
	v2954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2937)+28)))
	v2955 = v2954
	goto L503
L514:
	;
	v2960 = v2955
	goto L501
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2913))) = v1499
	F_errmsg_internal(m, int32(487717), v2913)
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_errfinish(m, int32(494969), int32(2366), int32(165768))
	mBase = m.M
	v2979 = m.ExcPending
	if v2979 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L518:
	;
	v2992 = int32(1)
	goto L489
L519:
	;
	v2992 = v2983
	goto L489
L520:
	;
	if v3001 != 0 {
		goto L399
	} else {
		goto L521
	}
L521:
	;
	v3004 = int32(0)
	v3009 = v3004
	v3010 = int32(1)
	v3011 = v3004
	goto L487
L522:
	;
	v3024 = m.T0[v3021].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v2556, v2676, v2632)
	mBase = m.M
	v3025 = m.ExcPending
	if v3025 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v3048 = v3024
	v3052 = v3015
	v3053 = v2645
	v3054 = v3016
	v3055 = v3017
	v3057 = v3019
	goto L398
L524:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3032 = m.ExcPending
	if v3032 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	F_errmsg(m, int32(448161), int32(0))
	mBase = m.M
	v3036 = m.ExcPending
	if v3036 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(498183), int32(7299), int32(392626))
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	v3048 = v2757
	v3052 = int32(0)
	v3053 = v3044
	v3054 = int32(0)
	v3055 = int32(1)
	v3057 = int32(1)
	goto L398
L529:
	;
	v3062 = v2632 + int32(1)
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+4))
	if v3062 < v3063 {
		v2632 = v3062
		v2636 = v3059
		v2643 = v3052
		v2645 = v3053
		v2646 = v3054
		v2647 = v3055
		v2649 = v3057
		goto L396
	} else {
		goto L530
	}
L530:
	;
	goto L397
L531:
	;
	v3129 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v3131 = F_palloc0(m, int32(80))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3131))) = int32(372)
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3127)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v3131)+72)) = v3129
	v3137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+56)) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+52)) = v3127
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+48)) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+44)) = v3135
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+4)) = v3143
	v3145 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3131)+8)) = v3145
	v3147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3131)+16)) = v3147
	v3149 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3131)+24)) = v3149
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v3151)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3131)+32)) = v3152
	v3154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3131)+36)) = uint8(v3154)
	v3156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3131)+37)) = uint8(v3156)
	v12585 = v3131
	v12589 = v47
	goto L3
L533:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3165 = F_create_plan_recurse(m, l0, v3163, int32(1))
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v3167 = int32(0)
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3168)+4))
	if v3169 == v3167 {
		v3251 = v3167
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v3293 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v3293)&int64(9223372036854775807)) {
		goto L551
	} else {
		goto L552
	}
L536:
	;
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3169)+4))
	if v3172 <= int32(0) {
		v3251 = v3167
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v3168)+8))
	v3180 = v3167
	v3181 = v3158
	v3183 = v4
	goto L538
L538:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3169)+12))
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v3222+v3183<<(uint(int32(2))%32))))
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3227 != 0 {
		goto L540
	} else {
		goto L541
	}
L539:
	;
	v3251 = v3243
	goto L535
L540:
	;
	v3228 = F_replace_nestloop_params_mutator(m, v3226, l0)
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		goto L1
	} else {
		goto L543
	}
L541:
	;
	v3230 = v3226
	goto L542
L542:
	;
	v3232 = int32(0)
	v3234 = F_makeTargetEntry(m, v3230, base.I32_extend16_s(v3181), v3232, v3232)
	mBase = m.M
	v3235 = m.ExcPending
	if v3235 != 0 {
		goto L1
	} else {
		goto L544
	}
L543:
	;
	v3230 = v3228
	goto L542
L544:
	;
	if v3175 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(v3175-int32(4)+v3181<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3234)+16)) = v3239
	goto L547
L546:
	;
	goto L547
L547:
	;
	v3243 = F_lappend(m, v3180, v3234)
	mBase = m.M
	v3244 = m.ExcPending
	if v3244 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	v3246 = v3183 + int32(1)
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3169)+4))
	if v3246 < v3247 {
		v3180 = v3243
		v3181 = v3181 + int32(1)
		v3183 = v3246
		goto L538
	} else {
		goto L549
	}
L549:
	;
	goto L539
L550:
	;
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v3317 = F_palloc0(m, int32(96))
	mBase = m.M
	v3318 = m.ExcPending
	if v3318 != 0 {
		goto L1
	} else {
		goto L563
	}
L551:
	;
	v3313 = int32(2147483647)
	goto L550
L552:
	;
	goto L553
L553:
	;
	if base.F64_le(v3293, float64(0)) != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3313 = int32(0)
	goto L550
L555:
	;
	goto L556
L556:
	;
	v3303 = float64(2.147483647e+09)
	if base.F64_lt(v3293, v3303) != 0 {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v3306 = v3293
	goto L559
L558:
	;
	v3306 = v3303
	goto L559
L559:
	;
	if base.F64_lt(base.F64_abs(v3306), float64(2.147483648e+09)) != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3310 = base.I32_trunc_f64_s(v3306)
	v3313 = v3310
	goto L550
L561:
	;
	goto L562
L562:
	;
	v3313 = int32(-2147483648)
	goto L550
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3317))) = int32(336)
	if v3314 != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+4))
	v3322 = v3321
	goto L566
L565:
	;
	v3322 = v4
	goto L566
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+76)) = v3322
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+72)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+56)) = v3165
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+52)) = v3161
	v3327 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+48)) = v3327
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+44)) = v3251
	if v3327 < v3322 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3334 = F_palloc(m, v3322<<(uint(int32(1))%32))
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+92)) = v3313
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+4)) = v3509
	v3511 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3317)+8)) = v3511
	v3513 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3317)+16)) = v3513
	v3515 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3317)+24)) = v3515
	v3517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3517)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+32)) = v3518
	v3520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3317)+36)) = uint8(v3520)
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3317)+37)) = uint8(v3522)
	v12585 = v3317
	v12589 = v47
	goto L3
L570:
	;
	v3337 = v3322 << (uint(int32(2)) % 32)
	v3338 = F_palloc(m, v3337)
	mBase = m.M
	v3339 = m.ExcPending
	if v3339 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	v3340 = F_palloc(m, v3337)
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	if v3314 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+88)) = v3340
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+84)) = v3338
	*(*int32)(unsafe.Add(mBase, uint32(v3317)+80)) = v3334
	goto L569
L574:
	;
	v3344 = int32(0)
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+4))
	if v3345 <= v3344 {
		goto L573
	} else {
		goto L575
	}
L575:
	;
	v3348 = v3344
	goto L576
L576:
	;
	v3396 = v3348 << (uint(int32(2)) % 32)
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+12))
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(v3396+v3397)))
	v3400 = *(*int32)(unsafe.Add(mBase, uint32(v3317)+44))
	v3401 = F_get_sortgroupclause_tle(m, v3399, v3400)
	mBase = m.M
	v3402 = m.ExcPending
	if v3402 != 0 {
		goto L1
	} else {
		goto L578
	}
L577:
	;
	goto L573
L578:
	;
	v3403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3401)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3334+v3348<<(uint(int32(1))%32)))) = uint16(v3403)
	v3406 = *(*int32)(unsafe.Add(mBase, uint32(v3399)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3396+v3338))) = v3406
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(v3401)+4))
	v3410 = F_exprCollation(m, v3409)
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3396+v3340))) = v3410
	v3414 = v3348 + int32(1)
	v3415 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+4))
	if v3414 < v3415 {
		v3348 = v3414
		goto L576
	} else {
		goto L580
	}
L580:
	;
	goto L577
L581:
	;
	v3650 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3652 = l2 | int32(4)
	v3653 = F_create_plan_recurse(m, l0, v3650, v3652)
	mBase = m.M
	v3654 = m.ExcPending
	if v3654 != 0 {
		goto L1
	} else {
		goto L596
	}
L582:
	;
	v3529 = *(*int32)(unsafe.Add(mBase, uint32(v3525)+4))
	if v3529 <= int32(0) {
		v3612 = v4
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v3532 = *(*int32)(unsafe.Add(mBase, uint32(v3524)+8))
	v3538 = int32(1)
	v3540 = v4
	v3541 = v4
	goto L584
L584:
	;
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v3525)+12))
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(v3579+v3540<<(uint(int32(2))%32))))
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3584 != 0 {
		goto L586
	} else {
		goto L587
	}
L585:
	;
	v3612 = v3600
	goto L581
L586:
	;
	v3585 = F_replace_nestloop_params_mutator(m, v3583, l0)
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L1
	} else {
		goto L589
	}
L587:
	;
	v3587 = v3583
	goto L588
L588:
	;
	v3589 = int32(0)
	v3591 = F_makeTargetEntry(m, v3587, base.I32_extend16_s(v3538), v3589, v3589)
	mBase = m.M
	v3592 = m.ExcPending
	if v3592 != 0 {
		goto L1
	} else {
		goto L590
	}
L589:
	;
	v3587 = v3585
	goto L588
L590:
	;
	if v3532 != 0 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v3532-int32(4)+v3538<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3591)+16)) = v3596
	goto L593
L592:
	;
	goto L593
L593:
	;
	v3600 = F_lappend(m, v3541, v3591)
	mBase = m.M
	v3601 = m.ExcPending
	if v3601 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	v3603 = v3540 + int32(1)
	v3604 = *(*int32)(unsafe.Add(mBase, uint32(v3525)+4))
	if v3603 < v3604 {
		v3538 = v3538 + int32(1)
		v3540 = v3603
		v3541 = v3600
		goto L584
	} else {
		goto L595
	}
L595:
	;
	goto L585
L596:
	;
	v3655 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3656 = F_create_plan_recurse(m, l0, v3655, v3652)
	mBase = m.M
	v3657 = m.ExcPending
	if v3657 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v3658 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v3658)&int64(9223372036854775807)) {
		goto L599
	} else {
		goto L600
	}
L598:
	;
	v3679 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v3680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v3683 = F_palloc0(m, int32(104))
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L1
	} else {
		goto L611
	}
L599:
	;
	v3678 = int32(2147483647)
	goto L598
L600:
	;
	goto L601
L601:
	;
	if base.F64_le(v3658, float64(0)) != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v3678 = int32(0)
	goto L598
L603:
	;
	goto L604
L604:
	;
	v3668 = float64(2.147483647e+09)
	if base.F64_lt(v3658, v3668) != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v3671 = v3658
	goto L607
L606:
	;
	v3671 = v3668
	goto L607
L607:
	;
	if base.F64_lt(base.F64_abs(v3671), float64(2.147483648e+09)) != 0 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v3675 = base.I32_trunc_f64_s(v3671)
	v3678 = v3675
	goto L598
L609:
	;
	goto L610
L610:
	;
	v3678 = int32(-2147483648)
	goto L598
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3683))) = int32(371)
	if v3679 != 0 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(v3679)+4))
	v3688 = v3687
	goto L614
L613:
	;
	v3688 = v4
	goto L614
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+56)) = v3656
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+52)) = v3653
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+44)) = v3612
	v3696 = F_palloc(m, v3688<<(uint(int32(1))%32))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	v3699 = v3688 << (uint(int32(2)) % 32)
	v3700 = F_palloc(m, v3699)
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v3702 = F_palloc(m, v3699)
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	v3704 = F_palloc(m, v3688)
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	if v3679 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+100)) = v3678
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+96)) = v3704
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+92)) = v3702
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+88)) = v3700
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+84)) = v3696
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+80)) = v3688
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+76)) = v3680
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+72)) = v3681
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+4)) = v3842
	v3844 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3683)+8)) = v3844
	v3846 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3683)+16)) = v3846
	v3848 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3683)+24)) = v3848
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v3850)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3683)+32)) = v3851
	v3853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3683)+36)) = uint8(v3853)
	v3855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3683)+37)) = uint8(v3855)
	v12585 = v3683
	v12589 = v47
	goto L3
L620:
	;
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3679)+4))
	if v3708 <= int32(0) {
		goto L619
	} else {
		goto L621
	}
L621:
	;
	if v3680 == int32(1) {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v3715 = int32(8)
	goto L624
L623:
	;
	v3715 = int32(12)
	goto L624
L624:
	;
	v3717 = int32(0)
	goto L625
L625:
	;
	v3765 = v3717 << (uint(int32(2)) % 32)
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v3679)+12))
	v3768 = *(*int32)(unsafe.Add(mBase, uint32(v3765+v3766)))
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v3683)+44))
	v3770 = F_get_sortgroupclause_tle(m, v3768, v3769)
	mBase = m.M
	v3771 = m.ExcPending
	if v3771 != 0 {
		goto L1
	} else {
		goto L627
	}
L626:
	;
	goto L619
L627:
	;
	v3772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3770)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3696+v3717<<(uint(int32(1))%32)))) = uint16(v3772)
	v3776 = *(*int32)(unsafe.Add(mBase, uint32(v3768+v3715)))
	*(*int32)(unsafe.Add(mBase, uint32(v3765+v3700))) = v3776
	v3779 = *(*int32)(unsafe.Add(mBase, uint32(v3770)+4))
	v3780 = F_exprCollation(m, v3779)
	mBase = m.M
	v3781 = m.ExcPending
	if v3781 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3765+v3702))) = v3780
	v3784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3768)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3717+v3704))) = uint8(v3784)
	v3787 = v3717 + int32(1)
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v3679)+4))
	if v3787 < v3788 {
		v3717 = v3787
		goto L625
	} else {
		goto L629
	}
L629:
	;
	goto L626
L630:
	;
	v3859 = *(*int32)(unsafe.Add(mBase, uint32(v3858)+4))
	v3860 = v3859
	goto L632
L631:
	;
	v3860 = v4
	goto L632
L632:
	;
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+16))
	if v3861 != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3862 = *(*int32)(unsafe.Add(mBase, uint32(v3861)+4))
	v3863 = v3862
	goto L635
L634:
	;
	v3863 = v4
	goto L635
L635:
	;
	v3864 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3866 = F_create_plan_recurse(m, l0, v3864, int32(6))
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	v3868 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(v3868)+4))
	if v3869 == int32(0) {
		v3957 = v4
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v3997 = F_palloc(m, v3860<<(uint(int32(1))%32))
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L1
	} else {
		goto L652
	}
L638:
	;
	v3873 = int32(0)
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+4))
	if v3874 <= v3873 {
		v3957 = v4
		goto L637
	} else {
		goto L639
	}
L639:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3868)+8))
	v3882 = v3873
	v3883 = int32(1)
	v3886 = v4
	goto L640
L640:
	;
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+12))
	v3928 = *(*int32)(unsafe.Add(mBase, uint32(v3924+v3882<<(uint(int32(2))%32))))
	v3929 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3929 != 0 {
		goto L642
	} else {
		goto L643
	}
L641:
	;
	v3957 = v3945
	goto L637
L642:
	;
	v3930 = F_replace_nestloop_params_mutator(m, v3928, l0)
	mBase = m.M
	v3931 = m.ExcPending
	if v3931 != 0 {
		goto L1
	} else {
		goto L645
	}
L643:
	;
	v3932 = v3928
	goto L644
L644:
	;
	v3934 = int32(0)
	v3936 = F_makeTargetEntry(m, v3932, base.I32_extend16_s(v3883), v3934, v3934)
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L1
	} else {
		goto L646
	}
L645:
	;
	v3932 = v3930
	goto L644
L646:
	;
	if v3877 != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v3941 = *(*int32)(unsafe.Add(mBase, uint32(v3877-int32(4)+v3883<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3936)+16)) = v3941
	goto L649
L648:
	;
	goto L649
L649:
	;
	v3945 = F_lappend(m, v3886, v3936)
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	v3948 = v3882 + int32(1)
	v3949 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+4))
	if v3948 < v3949 {
		v3882 = v3948
		v3883 = v3883 + int32(1)
		v3886 = v3945
		goto L640
	} else {
		goto L651
	}
L651:
	;
	goto L641
L652:
	;
	v4000 = v3860 << (uint(int32(2)) % 32)
	v4001 = F_palloc(m, v4000)
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	v4003 = F_palloc(m, v4000)
	mBase = m.M
	v4004 = m.ExcPending
	if v4004 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	v4005 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+12))
	if v4005 == int32(0) {
		v4084 = v4
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v4126 = F_palloc(m, v3863<<(uint(int32(1))%32))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L1
	} else {
		goto L663
	}
L656:
	;
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v4005)+4))
	if v4008 <= int32(0) {
		v4084 = v4
		goto L655
	} else {
		goto L657
	}
L657:
	;
	v4015 = v4
	goto L658
L658:
	;
	v4059 = v4015 << (uint(int32(2)) % 32)
	v4060 = *(*int32)(unsafe.Add(mBase, uint32(v4005)+12))
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4059+v4060)))
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v3866)+44))
	v4064 = F_get_sortgroupclause_tle(m, v4062, v4063)
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L1
	} else {
		goto L660
	}
L659:
	;
	v4084 = v4077
	goto L655
L660:
	;
	v4066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4064)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3997+v4015<<(uint(int32(1))%32)))) = uint16(v4066)
	v4069 = *(*int32)(unsafe.Add(mBase, uint32(v4062)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4059+v4001))) = v4069
	v4072 = *(*int32)(unsafe.Add(mBase, uint32(v4064)+4))
	v4073 = F_exprCollation(m, v4072)
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4059+v4003))) = v4073
	v4077 = v4015 + int32(1)
	v4078 = *(*int32)(unsafe.Add(mBase, uint32(v4005)+4))
	if v4077 < v4078 {
		v4015 = v4077
		goto L658
	} else {
		goto L662
	}
L662:
	;
	goto L659
L663:
	;
	v4129 = v3863 << (uint(int32(2)) % 32)
	v4130 = F_palloc(m, v4129)
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	v4132 = F_palloc(m, v4129)
	mBase = m.M
	v4133 = m.ExcPending
	if v4133 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v4134 = int32(0)
	v4135 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+16))
	if v4135 == v4134 {
		v4210 = v4134
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v4254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+88)))
	v4256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v4258 = F_palloc0(m, int32(152))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L1
	} else {
		goto L674
	}
L667:
	;
	v4138 = *(*int32)(unsafe.Add(mBase, uint32(v4135)+4))
	if v4138 <= int32(0) {
		v4210 = v4134
		goto L666
	} else {
		goto L668
	}
L668:
	;
	v4141 = v4134
	goto L669
L669:
	;
	v4189 = v4141 << (uint(int32(2)) % 32)
	v4190 = *(*int32)(unsafe.Add(mBase, uint32(v4135)+12))
	v4192 = *(*int32)(unsafe.Add(mBase, uint32(v4189+v4190)))
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v3866)+44))
	v4194 = F_get_sortgroupclause_tle(m, v4192, v4193)
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L1
	} else {
		goto L671
	}
L670:
	;
	v4210 = v4207
	goto L666
L671:
	;
	v4196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4194)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4126+v4141<<(uint(int32(1))%32)))) = uint16(v4196)
	v4199 = *(*int32)(unsafe.Add(mBase, uint32(v4192)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4189+v4130))) = v4199
	v4202 = *(*int32)(unsafe.Add(mBase, uint32(v4194)+4))
	v4203 = F_exprCollation(m, v4202)
	mBase = m.M
	v4204 = m.ExcPending
	if v4204 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4189+v4132))) = v4203
	v4207 = v4141 + int32(1)
	v4208 = *(*int32)(unsafe.Add(mBase, uint32(v4135)+4))
	if v4207 < v4208 {
		v4141 = v4207
		goto L669
	} else {
		goto L673
	}
L673:
	;
	goto L670
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4258))) = int32(366)
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+72)) = v4262
	v4264 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+108)) = v4132
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+104)) = v4130
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+100)) = v4126
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+96)) = v4210
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+92)) = v4003
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+88)) = v4001
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+84)) = v3997
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+80)) = v4084
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+76)) = v4264
	v4274 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+112)) = v4274
	v4276 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+116)) = v4276
	v4278 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+128)) = v4256
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+124)) = v4256
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+120)) = v4278
	v4282 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+132)) = v4282
	v4284 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+136)) = v4284
	v4286 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+140)) = v4286
	v4288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3857)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4258)+144)) = uint8(v4288)
	v4290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3857)+45)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4258)+146)) = uint8(v4255)
	*(*uint8)(unsafe.Add(mBase, uint32(v4258)+145)) = uint8(v4290)
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+52)) = v3866
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+44)) = v3957
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+48)) = v4254
	v4298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+4)) = v4298
	v4300 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4258)+8)) = v4300
	v4302 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v4258)+16)) = v4302
	v4304 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v4258)+24)) = v4304
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4307 = *(*int32)(unsafe.Add(mBase, uint32(v4306)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4258)+32)) = v4307
	v4309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4258)+36)) = uint8(v4309)
	v4311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4258)+37)) = uint8(v4311)
	v12585 = v4258
	v12589 = v47
	goto L3
L675:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v4319 = F_create_plan_recurse(m, l0, v4317, int32(4))
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L1
	} else {
		goto L678
	}
L676:
	;
	goto L677
L677:
	;
	v5711 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v5713 = F_create_plan_recurse(m, l0, v5711, int32(4))
	mBase = m.M
	v5714 = m.ExcPending
	if v5714 != 0 {
		goto L1
	} else {
		goto L844
	}
L678:
	;
	v4321 = int32(2)
	v4322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v4322 == int32(0) {
		v4561 = v4321
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v4603 = F_palloc0(m, v4561)
	mBase = m.M
	v4604 = m.ExcPending
	if v4604 != 0 {
		goto L1
	} else {
		goto L712
	}
L680:
	;
	v4325 = *(*int32)(unsafe.Add(mBase, uint32(v4322)+4))
	if v4325 <= int32(0) {
		v4561 = v4321
		goto L679
	} else {
		goto L681
	}
L681:
	;
	v4328 = int32(0)
	if v4328 < v4325 {
		goto L682
	} else {
		goto L683
	}
L682:
	;
	v4331 = v4325
	goto L684
L683:
	;
	v4331 = v4328
	goto L684
L684:
	;
	v4333 = v4331 & int32(3)
	v4334 = int32(0)
	if int32(4) <= v4325 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4322)+12))
	v4343 = v4334
	v4345 = int32(0)
	v4346 = v4
	goto L688
L686:
	;
	v4411 = v4334
	v4414 = v4
	goto L687
L687:
	;
	if v4333 != 0 {
		goto L703
	} else {
		goto L704
	}
L688:
	;
	v4387 = v4339 + v4346<<(uint(int32(2))%32)
	v4388 = *(*int32)(unsafe.Add(mBase, uint32(v4387)+12))
	v4389 = *(*int32)(unsafe.Add(mBase, uint32(v4388)+4))
	v4390 = *(*int32)(unsafe.Add(mBase, uint32(v4387)+8))
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v4390)+4))
	v4392 = *(*int32)(unsafe.Add(mBase, uint32(v4387)+4))
	v4393 = *(*int32)(unsafe.Add(mBase, uint32(v4392)+4))
	v4394 = *(*int32)(unsafe.Add(mBase, uint32(v4387)))
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v4394)+4))
	if base.Ui32(v4343) < base.Ui32(v4395) {
		goto L690
	} else {
		goto L691
	}
L689:
	;
	v4411 = v4403
	v4414 = v4405
	goto L687
L690:
	;
	v4397 = v4395
	goto L692
L691:
	;
	v4397 = v4343
	goto L692
L692:
	;
	if base.Ui32(v4397) < base.Ui32(v4393) {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v4399 = v4393
	goto L695
L694:
	;
	v4399 = v4397
	goto L695
L695:
	;
	if base.Ui32(v4399) < base.Ui32(v4391) {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	v4401 = v4391
	goto L698
L697:
	;
	v4401 = v4399
	goto L698
L698:
	;
	if base.Ui32(v4401) < base.Ui32(v4389) {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v4403 = v4389
	goto L701
L700:
	;
	v4403 = v4401
	goto L701
L701:
	;
	v4404 = int32(4)
	v4405 = v4346 + v4404
	v4407 = v4345 + v4404
	if v4407 != v4331&int32(2147483644) {
		v4343 = v4403
		v4345 = v4407
		v4346 = v4405
		goto L688
	} else {
		goto L702
	}
L702:
	;
	goto L689
L703:
	;
	v4453 = *(*int32)(unsafe.Add(mBase, uint32(v4322)+12))
	v4457 = v4411
	v4458 = int32(0)
	v4460 = v4414
	goto L706
L704:
	;
	v4513 = v4411
	goto L705
L705:
	;
	v4561 = v4513<<(uint(int32(1))%32) + int32(2)
	goto L679
L706:
	;
	v4502 = *(*int32)(unsafe.Add(mBase, uint32(v4453+v4460<<(uint(int32(2))%32))))
	v4503 = *(*int32)(unsafe.Add(mBase, uint32(v4502)+4))
	if base.Ui32(v4457) < base.Ui32(v4503) {
		goto L708
	} else {
		goto L709
	}
L707:
	;
	v4513 = v4505
	goto L705
L708:
	;
	v4505 = v4503
	goto L710
L709:
	;
	v4505 = v4457
	goto L710
L710:
	;
	v4506 = int32(1)
	v4509 = v4458 + v4506
	if v4509 != v4333 {
		v4457 = v4505
		v4458 = v4509
		v4460 = v4460 + v4506
		goto L706
	} else {
		goto L711
	}
L711:
	;
	goto L707
L712:
	;
	v4605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v4605 == int32(0) {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v4603
	if v4316 == int32(0) {
		v5345 = v4603
		v5357 = v4
		goto L720
	} else {
		goto L721
	}
L714:
	;
	v4608 = int32(0)
	v4609 = *(*int32)(unsafe.Add(mBase, uint32(v4605)+4))
	if v4609 <= v4608 {
		goto L713
	} else {
		goto L715
	}
L715:
	;
	v4614 = v4608
	goto L716
L716:
	;
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v4605)+12))
	v4660 = *(*int32)(unsafe.Add(mBase, uint32(v4656+v4614<<(uint(int32(2))%32))))
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+44))
	v4662 = F_get_sortgroupclause_tle(m, v4660, v4661)
	mBase = m.M
	v4663 = m.ExcPending
	if v4663 != 0 {
		goto L1
	} else {
		goto L718
	}
L717:
	;
	goto L713
L718:
	;
	v4664 = *(*int32)(unsafe.Add(mBase, uint32(v4660)+4))
	v4665 = int32(1)
	v4668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4662)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4603+v4664<<(uint(v4665)%32)))) = uint16(v4668)
	v4671 = v4614 + v4665
	v4672 = *(*int32)(unsafe.Add(mBase, uint32(v4605)+4))
	if v4671 < v4672 {
		v4614 = v4671
		goto L716
	} else {
		goto L719
	}
L719:
	;
	goto L717
L720:
	;
	v5384 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v5385 = *(*int32)(unsafe.Add(mBase, uint32(v5384)))
	v5386 = *(*int32)(unsafe.Add(mBase, uint32(v5385)+4))
	if v5386 == int32(0) {
		goto L801
	} else {
		goto L802
	}
L721:
	;
	v4721 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v4721 <= int32(1) {
		v5345 = v4603
		v5357 = v4
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v4724 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v4725 = *(*int32)(unsafe.Add(mBase, uint32(v4724)))
	v4726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4725)+25)))
	v4744 = int32(1)
	v4745 = v4
	v4746 = v4726
	goto L723
L723:
	;
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v4776 = *(*int32)(unsafe.Add(mBase, uint32(v4772+v4744<<(uint(int32(2))%32))))
	v4777 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+4))
	if v4777 == int32(0) {
		goto L726
	} else {
		goto L727
	}
L724:
	;
	v5339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v5345 = v5339
	v5357 = v5333
	goto L720
L725:
	;
	v4899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4776)+25)))
	if (v4899|v4746)&int32(1) == int32(0) {
		goto L739
	} else {
		goto L740
	}
L726:
	;
	v4781 = F_palloc0(m, int32(0))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L1
	} else {
		goto L729
	}
L727:
	;
	goto L728
L728:
	;
	v4783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v4784 = *(*int32)(unsafe.Add(mBase, uint32(v4777)+4))
	v4787 = F_palloc0(m, v4784<<(uint(int32(1))%32))
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L1
	} else {
		goto L730
	}
L729:
	;
	v4858 = v4781
	goto L725
L730:
	;
	v4789 = int32(0)
	v4790 = *(*int32)(unsafe.Add(mBase, uint32(v4777)+4))
	if v4790 <= v4789 {
		v4858 = v4787
		goto L725
	} else {
		goto L731
	}
L731:
	;
	v4795 = v4789
	goto L732
L732:
	;
	v4837 = int32(1)
	v4840 = *(*int32)(unsafe.Add(mBase, uint32(v4777)+12))
	v4844 = *(*int32)(unsafe.Add(mBase, uint32(v4840+v4795<<(uint(int32(2))%32))))
	v4845 = *(*int32)(unsafe.Add(mBase, uint32(v4844)+4))
	v4849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4783+v4845<<(uint(v4837)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4787+v4795<<(uint(v4837)%32)))) = uint16(v4849)
	v4852 = v4795 + v4837
	v4853 = *(*int32)(unsafe.Add(mBase, uint32(v4777)+4))
	if v4852 < v4853 {
		v4795 = v4852
		goto L732
	} else {
		goto L734
	}
L733:
	;
	v4858 = v4787
	goto L725
L734:
	;
	goto L733
L735:
	;
	if v5229 != 0 {
		goto L776
	} else {
		goto L777
	}
L736:
	;
	v5221 = int32(0)
	v5222 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+8))
	v5223 = *(*int32)(unsafe.Add(mBase, uint32(v5222)+12))
	v5224 = *(*int32)(unsafe.Add(mBase, uint32(v5223)))
	v5229 = v5224
	v5232 = v5182
	v5245 = v5221
	v5271 = base.B2i32(v5224 != v5221)
	goto L735
L737:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5167 = m.ExcPending
	if v5167 != 0 {
		goto L1
	} else {
		goto L773
	}
L738:
	;
	v5160 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+8))
	v5161 = *(*int32)(unsafe.Add(mBase, uint32(v5160)+12))
	v5162 = *(*int32)(unsafe.Add(mBase, uint32(v5161)))
	v5229 = v5162
	v5232 = v5121
	v5245 = v4746
	v5271 = int32(2)
	goto L735
L739:
	;
	v4905 = int32(0)
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+44))
	v4908 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+4))
	if v4908 != 0 {
		goto L742
	} else {
		goto L743
	}
L740:
	;
	goto L741
L741:
	;
	v5111 = int32(0)
	if v4899&int32(1) == v5111 {
		v5182 = v5111
		goto L736
	} else {
		goto L772
	}
L742:
	;
	v4909 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+4))
	v4910 = v4909
	goto L744
L743:
	;
	v4910 = v4905
	goto L744
L744:
	;
	v4913 = F_palloc(m, v4910<<(uint(int32(1))%32))
	mBase = m.M
	v4914 = m.ExcPending
	if v4914 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	v4916 = v4910 << (uint(int32(2)) % 32)
	v4917 = F_palloc(m, v4916)
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	v4919 = F_palloc(m, v4916)
	mBase = m.M
	v4920 = m.ExcPending
	if v4920 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	v4921 = F_palloc(m, v4910)
	mBase = m.M
	v4922 = m.ExcPending
	if v4922 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	if v4908 == int32(0) {
		v5042 = v4905
		goto L749
	} else {
		goto L750
	}
L749:
	;
	v5085 = F_palloc0(m, int32(96))
	mBase = m.M
	v5086 = m.ExcPending
	if v5086 != 0 {
		goto L1
	} else {
		goto L770
	}
L750:
	;
	v4925 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+4))
	if v4925 <= int32(0) {
		v5042 = v4905
		goto L749
	} else {
		goto L751
	}
L751:
	;
	v4930 = v4905
	goto L752
L752:
	;
	v4973 = v4930 << (uint(int32(2)) % 32)
	v4974 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+12))
	v4976 = *(*int32)(unsafe.Add(mBase, uint32(v4973+v4974)))
	v4978 = v4930 << (uint(int32(1)) % 32)
	v4980 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4858+v4978))))
	if v4907 != 0 {
		goto L756
	} else {
		goto L757
	}
L753:
	;
	v5042 = v5037
	goto L749
L754:
	;
	if v5018 == int32(0) {
		goto L737
	} else {
		goto L767
	}
L755:
	;
	goto L754
L756:
	;
	v4984 = *(*int32)(unsafe.Add(mBase, uint32(v4907)+4))
	if v4984 <= int32(0) {
		v5018 = int32(0)
		goto L755
	} else {
		goto L759
	}
L757:
	;
	goto L758
L758:
	;
	v5018 = int32(0)
	goto L755
L759:
	;
	v4987 = int32(0)
	if v4987 < v4984 {
		goto L760
	} else {
		goto L761
	}
L760:
	;
	v4990 = v4984
	goto L762
L761:
	;
	v4990 = v4987
	goto L762
L762:
	;
	v4991 = *(*int32)(unsafe.Add(mBase, uint32(v4907)+12))
	v4995 = int32(0)
	goto L763
L763:
	;
	v5003 = *(*int32)(unsafe.Add(mBase, uint32(v4991+v4995<<(uint(int32(2))%32))))
	v5004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5003)+8)))
	if v5004 == v4980&int32(65535) {
		v5018 = v5003
		goto L755
	} else {
		goto L765
	}
L764:
	;
	goto L758
L765:
	;
	v5007 = v4995 + int32(1)
	if v5007 != v4990 {
		v4995 = v5007
		goto L763
	} else {
		goto L766
	}
L766:
	;
	goto L764
L767:
	;
	v5023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5018)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4913+v4978))) = uint16(v5023)
	v5026 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v4973+v4917))) = v5026
	v5029 = *(*int32)(unsafe.Add(mBase, uint32(v5018)+4))
	v5030 = F_exprCollation(m, v5029)
	mBase = m.M
	v5031 = m.ExcPending
	if v5031 != 0 {
		goto L1
	} else {
		goto L768
	}
L768:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4973+v4919))) = v5030
	v5034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4976)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4930+v4921))) = uint8(v5034)
	v5037 = v4930 + int32(1)
	v5038 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+4))
	if v5037 < v5038 {
		v4930 = v5037
		goto L752
	} else {
		goto L769
	}
L769:
	;
	goto L753
L770:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5085))) = int32(362)
	v5089 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+44)) = v5089
	v5091 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+4))
	v5093 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+88)) = v4921
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+84)) = v4919
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+80)) = v4917
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+76)) = v4913
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+72)) = v5042
	v5099 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+56)) = v5099
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+52)) = v4319
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+48)) = v5099
	v5104 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5085)+4)) = v5091 + (v5093 ^ v5104)
	v5108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4776)+25)))
	if v5108&v5104 != 0 {
		v5121 = v5085
		goto L738
	} else {
		goto L771
	}
L771:
	;
	v5182 = v5085
	goto L736
L772:
	;
	v5121 = v5111
	goto L738
L773:
	;
	F_errmsg_internal(m, int32(152285), int32(0))
	mBase = m.M
	v5171 = m.ExcPending
	if v5171 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	F_errfinish(m, int32(498183), int32(6626), int32(152260))
	mBase = m.M
	v5176 = m.ExcPending
	if v5176 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L776:
	;
	v5272 = *(*int32)(unsafe.Add(mBase, uint32(v5229)+4))
	v5274 = v5272
	goto L778
L777:
	;
	v5274 = int32(0)
	goto L778
L778:
	;
	v5275 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+4))
	v5276 = F_extract_grouping_ops(m, v5275)
	mBase = m.M
	v5277 = m.ExcPending
	if v5277 != 0 {
		goto L1
	} else {
		goto L779
	}
L779:
	;
	v5278 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+4))
	v5279 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+44))
	v5280 = F_extract_grouping_collations(m, v5278, v5279)
	mBase = m.M
	v5281 = m.ExcPending
	if v5281 != 0 {
		goto L1
	} else {
		goto L780
	}
L780:
	;
	v5282 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v5283 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+8))
	v5284 = *(*float64)(unsafe.Add(mBase, uint32(v4776)+16))
	v5286 = F_palloc0(m, int32(128))
	mBase = m.M
	v5287 = m.ExcPending
	if v5287 != 0 {
		goto L1
	} else {
		goto L781
	}
L781:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5286))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5284)&int64(9223372036854775807)) {
		goto L783
	} else {
		goto L784
	}
L782:
	;
	v5310 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+120)) = v5310
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+116)) = v5283
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+112)) = v5310
	*(*int64)(unsafe.Add(mBase, uint32(v5286)+104)) = v5282
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+96)) = v5309
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+92)) = v5280
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+88)) = v5276
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+84)) = v4858
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+80)) = v5274
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+76)) = v5310
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+72)) = v5271
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+56)) = v5310
	*(*int32)(unsafe.Add(mBase, uint32(v5286)+52)) = v5232
	*(*int64)(unsafe.Add(mBase, uint32(v5286)+44)) = int64(0)
	if v5232 != 0 {
		goto L795
	} else {
		goto L796
	}
L783:
	;
	v5309 = int32(2147483647)
	goto L782
L784:
	;
	goto L785
L785:
	;
	if base.F64_le(v5284, float64(0)) != 0 {
		goto L786
	} else {
		goto L787
	}
L786:
	;
	v5309 = int32(0)
	goto L782
L787:
	;
	goto L788
L788:
	;
	v5299 = float64(2.147483647e+09)
	if base.F64_lt(v5284, v5299) != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v5302 = v5284
	goto L791
L790:
	;
	v5302 = v5299
	goto L791
L791:
	;
	if base.F64_lt(base.F64_abs(v5302), float64(2.147483648e+09)) != 0 {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	v5306 = base.I32_trunc_f64_s(v5302)
	v5309 = v5306
	goto L782
L793:
	;
	goto L794
L794:
	;
	v5309 = int32(-2147483648)
	goto L782
L795:
	;
	v5329 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5232)+52)) = v5329
	*(*int32)(unsafe.Add(mBase, uint32(v5232)+44)) = v5329
	goto L797
L796:
	;
	goto L797
L797:
	;
	v5333 = F_lappend(m, v4745, v5286)
	mBase = m.M
	v5334 = m.ExcPending
	if v5334 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	v5336 = v4744 + int32(1)
	v5337 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v5336 < v5337 {
		v4744 = v5336
		v4745 = v5333
		v4746 = v5245
		goto L723
	} else {
		goto L799
	}
L799:
	;
	goto L724
L800:
	;
	v5507 = int32(0)
	v5509 = *(*int32)(unsafe.Add(mBase, uint32(v5385)+8))
	v5510 = *(*int32)(unsafe.Add(mBase, uint32(v5509)+12))
	v5511 = *(*int32)(unsafe.Add(mBase, uint32(v5510)))
	if v5511 != 0 {
		goto L810
	} else {
		goto L811
	}
L801:
	;
	v5390 = F_palloc0(m, int32(0))
	mBase = m.M
	v5391 = m.ExcPending
	if v5391 != 0 {
		goto L1
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v5392 = *(*int32)(unsafe.Add(mBase, uint32(v5386)+4))
	v5395 = F_palloc0(m, v5392<<(uint(int32(1))%32))
	mBase = m.M
	v5396 = m.ExcPending
	if v5396 != 0 {
		goto L1
	} else {
		goto L805
	}
L804:
	;
	v5473 = v5390
	goto L800
L805:
	;
	v5397 = *(*int32)(unsafe.Add(mBase, uint32(v5386)+4))
	if v5397 <= int32(0) {
		v5473 = v5395
		goto L800
	} else {
		goto L806
	}
L806:
	;
	v5403 = int32(0)
	goto L807
L807:
	;
	v5445 = int32(1)
	v5448 = *(*int32)(unsafe.Add(mBase, uint32(v5386)+12))
	v5452 = *(*int32)(unsafe.Add(mBase, uint32(v5448+v5403<<(uint(int32(2))%32))))
	v5453 = *(*int32)(unsafe.Add(mBase, uint32(v5452)+4))
	v5457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5345+v5453<<(uint(v5445)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5395+v5403<<(uint(v5445)%32)))) = uint16(v5457)
	v5460 = v5403 + v5445
	v5461 = *(*int32)(unsafe.Add(mBase, uint32(v5386)+4))
	if v5460 < v5461 {
		v5403 = v5460
		goto L807
	} else {
		goto L809
	}
L808:
	;
	v5473 = v5395
	goto L800
L809:
	;
	goto L808
L810:
	;
	v5512 = *(*int32)(unsafe.Add(mBase, uint32(v5511)+4))
	v5513 = v5512
	goto L812
L811:
	;
	v5513 = v5507
	goto L812
L812:
	;
	v5514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5515 = *(*int32)(unsafe.Add(mBase, uint32(v5514)+4))
	if v5515 == int32(0) {
		v5601 = v5507
		goto L813
	} else {
		goto L814
	}
L813:
	;
	v5641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v5642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v5643 = *(*int32)(unsafe.Add(mBase, uint32(v5385)+4))
	v5644 = F_extract_grouping_ops(m, v5643)
	mBase = m.M
	v5645 = m.ExcPending
	if v5645 != 0 {
		goto L1
	} else {
		goto L828
	}
L814:
	;
	v5519 = int32(0)
	v5520 = *(*int32)(unsafe.Add(mBase, uint32(v5515)+4))
	if v5520 <= v5519 {
		v5601 = v5507
		goto L813
	} else {
		goto L815
	}
L815:
	;
	v5523 = *(*int32)(unsafe.Add(mBase, uint32(v5514)+8))
	v5528 = int32(1)
	v5529 = v5519
	v5530 = v5507
	goto L816
L816:
	;
	v5570 = *(*int32)(unsafe.Add(mBase, uint32(v5515)+12))
	v5574 = *(*int32)(unsafe.Add(mBase, uint32(v5570+v5529<<(uint(int32(2))%32))))
	v5575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5575 != 0 {
		goto L818
	} else {
		goto L819
	}
L817:
	;
	v5601 = v5591
	goto L813
L818:
	;
	v5576 = F_replace_nestloop_params_mutator(m, v5574, l0)
	mBase = m.M
	v5577 = m.ExcPending
	if v5577 != 0 {
		goto L1
	} else {
		goto L821
	}
L819:
	;
	v5578 = v5574
	goto L820
L820:
	;
	v5580 = int32(0)
	v5582 = F_makeTargetEntry(m, v5578, base.I32_extend16_s(v5528), v5580, v5580)
	mBase = m.M
	v5583 = m.ExcPending
	if v5583 != 0 {
		goto L1
	} else {
		goto L822
	}
L821:
	;
	v5578 = v5576
	goto L820
L822:
	;
	if v5523 != 0 {
		goto L823
	} else {
		goto L824
	}
L823:
	;
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v5523-int32(4)+v5528<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v5582)+16)) = v5587
	goto L825
L824:
	;
	goto L825
L825:
	;
	v5591 = F_lappend(m, v5530, v5582)
	mBase = m.M
	v5592 = m.ExcPending
	if v5592 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	v5594 = v5529 + int32(1)
	v5595 = *(*int32)(unsafe.Add(mBase, uint32(v5515)+4))
	if v5594 < v5595 {
		v5528 = v5528 + int32(1)
		v5529 = v5594
		v5530 = v5591
		goto L816
	} else {
		goto L827
	}
L827:
	;
	goto L817
L828:
	;
	v5646 = *(*int32)(unsafe.Add(mBase, uint32(v5385)+4))
	v5647 = *(*int32)(unsafe.Add(mBase, uint32(v4319)+44))
	v5648 = F_extract_grouping_collations(m, v5646, v5647)
	mBase = m.M
	v5649 = m.ExcPending
	if v5649 != 0 {
		goto L1
	} else {
		goto L829
	}
L829:
	;
	v5650 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v5651 = *(*int32)(unsafe.Add(mBase, uint32(v5385)+8))
	v5652 = *(*float64)(unsafe.Add(mBase, uint32(v5385)+16))
	v5654 = F_palloc0(m, int32(128))
	mBase = m.M
	v5655 = m.ExcPending
	if v5655 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5654))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5652)&int64(9223372036854775807)) {
		goto L832
	} else {
		goto L833
	}
L831:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+120)) = v5357
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+116)) = v5651
	v5680 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+112)) = v5680
	*(*int64)(unsafe.Add(mBase, uint32(v5654)+104)) = v5650
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+96)) = v5677
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+92)) = v5648
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+88)) = v5644
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+84)) = v5473
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+80)) = v5513
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+76)) = v5680
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+72)) = v5642
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+48)) = v5641
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+56)) = v5680
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+52)) = v4319
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+44)) = v5601
	v5696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+4)) = v5696
	v5698 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v5654)+8)) = v5698
	v5700 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v5654)+16)) = v5700
	v5702 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v5654)+24)) = v5702
	v5704 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5705 = *(*int32)(unsafe.Add(mBase, uint32(v5704)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+32)) = v5705
	v5707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5654)+36)) = uint8(v5707)
	v5709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5654)+37)) = uint8(v5709)
	v12585 = v5654
	v12589 = v47
	goto L3
L832:
	;
	v5677 = int32(2147483647)
	goto L831
L833:
	;
	goto L834
L834:
	;
	if base.F64_le(v5652, float64(0)) != 0 {
		goto L835
	} else {
		goto L836
	}
L835:
	;
	v5677 = int32(0)
	goto L831
L836:
	;
	goto L837
L837:
	;
	v5667 = float64(2.147483647e+09)
	if base.F64_lt(v5652, v5667) != 0 {
		goto L838
	} else {
		goto L839
	}
L838:
	;
	v5670 = v5652
	goto L840
L839:
	;
	v5670 = v5667
	goto L840
L840:
	;
	if base.F64_lt(base.F64_abs(v5670), float64(2.147483648e+09)) != 0 {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v5674 = base.I32_trunc_f64_s(v5670)
	v5677 = v5674
	goto L831
L842:
	;
	goto L843
L843:
	;
	v5677 = int32(-2147483648)
	goto L831
L844:
	;
	v5715 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5716 = *(*int32)(unsafe.Add(mBase, uint32(v5715)+4))
	if v5716 == int32(0) {
		v5807 = v4
		goto L845
	} else {
		goto L846
	}
L845:
	;
	v5841 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v5842 = F_order_qual_clauses(m, l0, v5841)
	mBase = m.M
	v5843 = m.ExcPending
	if v5843 != 0 {
		goto L1
	} else {
		goto L860
	}
L846:
	;
	v5720 = *(*int32)(unsafe.Add(mBase, uint32(v5716)+4))
	if v5720 <= int32(0) {
		v5807 = v4
		goto L845
	} else {
		goto L847
	}
L847:
	;
	v5723 = *(*int32)(unsafe.Add(mBase, uint32(v5715)+8))
	v5728 = int32(1)
	v5730 = v4
	v5736 = v4
	goto L848
L848:
	;
	v5770 = *(*int32)(unsafe.Add(mBase, uint32(v5716)+12))
	v5774 = *(*int32)(unsafe.Add(mBase, uint32(v5770+v5730<<(uint(int32(2))%32))))
	v5775 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5775 != 0 {
		goto L850
	} else {
		goto L851
	}
L849:
	;
	v5807 = v5791
	goto L845
L850:
	;
	v5776 = F_replace_nestloop_params_mutator(m, v5774, l0)
	mBase = m.M
	v5777 = m.ExcPending
	if v5777 != 0 {
		goto L1
	} else {
		goto L853
	}
L851:
	;
	v5778 = v5774
	goto L852
L852:
	;
	v5780 = int32(0)
	v5782 = F_makeTargetEntry(m, v5778, base.I32_extend16_s(v5728), v5780, v5780)
	mBase = m.M
	v5783 = m.ExcPending
	if v5783 != 0 {
		goto L1
	} else {
		goto L854
	}
L853:
	;
	v5778 = v5776
	goto L852
L854:
	;
	if v5723 != 0 {
		goto L855
	} else {
		goto L856
	}
L855:
	;
	v5787 = *(*int32)(unsafe.Add(mBase, uint32(v5723-int32(4)+v5728<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v5782)+16)) = v5787
	goto L857
L856:
	;
	goto L857
L857:
	;
	v5791 = F_lappend(m, v5736, v5782)
	mBase = m.M
	v5792 = m.ExcPending
	if v5792 != 0 {
		goto L1
	} else {
		goto L858
	}
L858:
	;
	v5794 = v5730 + int32(1)
	v5795 = *(*int32)(unsafe.Add(mBase, uint32(v5716)+4))
	if v5794 < v5795 {
		v5728 = v5728 + int32(1)
		v5730 = v5794
		v5736 = v5791
		goto L848
	} else {
		goto L859
	}
L859:
	;
	goto L849
L860:
	;
	v5844 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v5845 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v5846 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v5846 != 0 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	v5847 = *(*int32)(unsafe.Add(mBase, uint32(v5846)+4))
	v5848 = v5847
	goto L863
L862:
	;
	v5848 = v4
	goto L863
L863:
	;
	v5849 = *(*int32)(unsafe.Add(mBase, uint32(v5713)+44))
	v5850 = F_extract_grouping_cols(m, v5846, v5849)
	mBase = m.M
	v5851 = m.ExcPending
	if v5851 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	v5852 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v5853 = F_extract_grouping_ops(m, v5852)
	mBase = m.M
	v5854 = m.ExcPending
	if v5854 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	v5855 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v5856 = *(*int32)(unsafe.Add(mBase, uint32(v5713)+44))
	v5857 = F_extract_grouping_collations(m, v5855, v5856)
	mBase = m.M
	v5858 = m.ExcPending
	if v5858 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	v5859 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+96)))
	v5860 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v5862 = F_palloc0(m, int32(128))
	mBase = m.M
	v5863 = m.ExcPending
	if v5863 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5862))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5860)&int64(9223372036854775807)) {
		goto L869
	} else {
		goto L870
	}
L868:
	;
	v5886 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+120)) = v5886
	*(*int64)(unsafe.Add(mBase, uint32(v5862)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5862)+104)) = v5859
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+96)) = v5885
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+92)) = v5857
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+88)) = v5853
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+84)) = v5850
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+80)) = v5848
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+76)) = v5844
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+72)) = v5845
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+48)) = v5842
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+56)) = v5886
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+52)) = v5713
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+44)) = v5807
	v5903 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+4)) = v5903
	v5905 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v5862)+8)) = v5905
	v5907 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v5862)+16)) = v5907
	v5909 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v5862)+24)) = v5909
	v5911 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5912 = *(*int32)(unsafe.Add(mBase, uint32(v5911)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+32)) = v5912
	v5914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5862)+36)) = uint8(v5914)
	v5916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5862)+37)) = uint8(v5916)
	v12585 = v5862
	v12589 = v47
	goto L3
L869:
	;
	v5885 = int32(2147483647)
	goto L868
L870:
	;
	goto L871
L871:
	;
	if base.F64_le(v5860, float64(0)) != 0 {
		goto L872
	} else {
		goto L873
	}
L872:
	;
	v5885 = int32(0)
	goto L868
L873:
	;
	goto L874
L874:
	;
	v5875 = float64(2.147483647e+09)
	if base.F64_lt(v5860, v5875) != 0 {
		goto L875
	} else {
		goto L876
	}
L875:
	;
	v5878 = v5860
	goto L877
L876:
	;
	v5878 = v5875
	goto L877
L877:
	;
	if base.F64_lt(base.F64_abs(v5878), float64(2.147483648e+09)) != 0 {
		goto L878
	} else {
		goto L879
	}
L878:
	;
	v5882 = base.I32_trunc_f64_s(v5878)
	v5885 = v5882
	goto L868
L879:
	;
	goto L880
L880:
	;
	v5885 = int32(-2147483648)
	goto L868
L881:
	;
	v5922 = int32(0)
	v5923 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5924 = *(*int32)(unsafe.Add(mBase, uint32(v5923)+4))
	if v5924 == v5922 {
		v6007 = v5922
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v6049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6050 = F_order_qual_clauses(m, l0, v6049)
	mBase = m.M
	v6051 = m.ExcPending
	if v6051 != 0 {
		goto L1
	} else {
		goto L897
	}
L883:
	;
	v5928 = *(*int32)(unsafe.Add(mBase, uint32(v5924)+4))
	if v5928 <= int32(0) {
		v6007 = v5922
		goto L882
	} else {
		goto L884
	}
L884:
	;
	v5931 = *(*int32)(unsafe.Add(mBase, uint32(v5923)+8))
	v5936 = v5922
	v5937 = int32(1)
	v5939 = v4
	goto L885
L885:
	;
	v5978 = *(*int32)(unsafe.Add(mBase, uint32(v5924)+12))
	v5982 = *(*int32)(unsafe.Add(mBase, uint32(v5978+v5939<<(uint(int32(2))%32))))
	v5983 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5983 != 0 {
		goto L887
	} else {
		goto L888
	}
L886:
	;
	v6007 = v5999
	goto L882
L887:
	;
	v5984 = F_replace_nestloop_params_mutator(m, v5982, l0)
	mBase = m.M
	v5985 = m.ExcPending
	if v5985 != 0 {
		goto L1
	} else {
		goto L890
	}
L888:
	;
	v5986 = v5982
	goto L889
L889:
	;
	v5988 = int32(0)
	v5990 = F_makeTargetEntry(m, v5986, base.I32_extend16_s(v5937), v5988, v5988)
	mBase = m.M
	v5991 = m.ExcPending
	if v5991 != 0 {
		goto L1
	} else {
		goto L891
	}
L890:
	;
	v5986 = v5984
	goto L889
L891:
	;
	if v5931 != 0 {
		goto L892
	} else {
		goto L893
	}
L892:
	;
	v5995 = *(*int32)(unsafe.Add(mBase, uint32(v5931-int32(4)+v5937<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v5990)+16)) = v5995
	goto L894
L893:
	;
	goto L894
L894:
	;
	v5999 = F_lappend(m, v5936, v5990)
	mBase = m.M
	v6000 = m.ExcPending
	if v6000 != 0 {
		goto L1
	} else {
		goto L895
	}
L895:
	;
	v6002 = v5939 + int32(1)
	v6003 = *(*int32)(unsafe.Add(mBase, uint32(v5924)+4))
	if v6002 < v6003 {
		v5936 = v5999
		v5937 = v5937 + int32(1)
		v5939 = v6002
		goto L885
	} else {
		goto L896
	}
L896:
	;
	goto L886
L897:
	;
	v6052 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v6052 != 0 {
		goto L898
	} else {
		goto L899
	}
L898:
	;
	v6053 = *(*int32)(unsafe.Add(mBase, uint32(v6052)+4))
	v6054 = v6053
	goto L900
L899:
	;
	v6054 = v4
	goto L900
L900:
	;
	v6055 = *(*int32)(unsafe.Add(mBase, uint32(v5920)+44))
	v6056 = F_extract_grouping_cols(m, v6052, v6055)
	mBase = m.M
	v6057 = m.ExcPending
	if v6057 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	v6058 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6059 = F_extract_grouping_ops(m, v6058)
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	v6061 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6062 = *(*int32)(unsafe.Add(mBase, uint32(v5920)+44))
	v6063 = F_extract_grouping_collations(m, v6061, v6062)
	mBase = m.M
	v6064 = m.ExcPending
	if v6064 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	v6066 = F_palloc0(m, int32(88))
	mBase = m.M
	v6067 = m.ExcPending
	if v6067 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+84)) = v6063
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+80)) = v6059
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+76)) = v6056
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+72)) = v6054
	*(*int32)(unsafe.Add(mBase, uint32(v6066))) = int32(364)
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+48)) = v6050
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+52)) = v5920
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+44)) = v6007
	v6079 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+4)) = v6079
	v6081 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6066)+8)) = v6081
	v6083 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6066)+16)) = v6083
	v6085 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6066)+24)) = v6085
	v6087 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6088 = *(*int32)(unsafe.Add(mBase, uint32(v6087)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6066)+32)) = v6088
	v6090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6066)+36)) = uint8(v6090)
	v6092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6066)+37)) = uint8(v6092)
	v12585 = v6066
	v12589 = v47
	goto L3
L905:
	;
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6102 = *(*int32)(unsafe.Add(mBase, uint32(v6101)+8))
	v6103 = *(*int32)(unsafe.Add(mBase, uint32(v6102)+4))
	if base.Ui32(int32(5)) < base.Ui32(v6103) {
		v6115 = int32(0)
		goto L906
	} else {
		goto L907
	}
L906:
	;
	v6116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6117 = int32(0)
	v6129 = F_prepare_sort_from_pathkeys(m, v6097, v6099, v6115, v6117, v6117, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v6130 = m.ExcPending
	if v6130 != 0 {
		goto L1
	} else {
		goto L909
	}
L907:
	;
	v6106 = int32(0)
	if int32(1)<<(uint(v6103)%32)&int32(44) == v6106 {
		v6115 = v6106
		goto L906
	} else {
		goto L908
	}
L908:
	;
	v6113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6114 = *(*int32)(unsafe.Add(mBase, uint32(v6113)+8))
	v6115 = v6114
	goto L906
L909:
	;
	v6131 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v6132 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v6133 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v6134 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v6135 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v6137 = F_palloc0(m, int32(104))
	mBase = m.M
	v6138 = m.ExcPending
	if v6138 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6137))) = int32(363)
	v6141 = *(*int32)(unsafe.Add(mBase, uint32(v6129)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+96)) = v6116
	v6143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+56)) = v6143
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+52)) = v6129
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+48)) = v6143
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+44)) = v6141
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+88)) = v6135
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+84)) = v6134
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+80)) = v6133
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+76)) = v6132
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+72)) = v6131
	v6154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+4)) = v6154
	v6156 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6137)+8)) = v6156
	v6158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6137)+16)) = v6158
	v6160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6137)+24)) = v6160
	v6162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6163 = *(*int32)(unsafe.Add(mBase, uint32(v6162)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6137)+32)) = v6163
	v6165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6137)+36)) = uint8(v6165)
	v6167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6137)+37)) = uint8(v6167)
	v12585 = v6137
	v12589 = v47
	goto L3
L911:
	;
	v6174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6177 = *(*int32)(unsafe.Add(mBase, uint32(v6176)+8))
	v6178 = *(*int32)(unsafe.Add(mBase, uint32(v6177)+4))
	if base.Ui32(int32(5)) < base.Ui32(v6178) {
		v6190 = int32(0)
		goto L912
	} else {
		goto L913
	}
L912:
	;
	v6191 = int32(0)
	v6203 = F_prepare_sort_from_pathkeys(m, v6172, v6174, v6190, v6191, v6191, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v6204 = m.ExcPending
	if v6204 != 0 {
		goto L1
	} else {
		goto L915
	}
L913:
	;
	v6181 = int32(0)
	if int32(1)<<(uint(v6178)%32)&int32(44) == v6181 {
		v6190 = v6181
		goto L912
	} else {
		goto L914
	}
L914:
	;
	v6188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6189 = *(*int32)(unsafe.Add(mBase, uint32(v6188)+8))
	v6190 = v6189
	goto L912
L915:
	;
	v6205 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v6206 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v6207 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v6208 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v6209 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v6211 = F_palloc0(m, int32(96))
	mBase = m.M
	v6212 = m.ExcPending
	if v6212 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6211))) = int32(362)
	v6215 = *(*int32)(unsafe.Add(mBase, uint32(v6203)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+44)) = v6215
	v6217 = *(*int32)(unsafe.Add(mBase, uint32(v6203)+4))
	v6219 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+88)) = v6209
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+84)) = v6208
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+80)) = v6207
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+76)) = v6206
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+72)) = v6205
	v6225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+56)) = v6225
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+52)) = v6203
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+48)) = v6225
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+4)) = v6217 + (v6219 ^ int32(1))
	v6234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+4)) = v6234
	v6236 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6211)+8)) = v6236
	v6238 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6211)+16)) = v6238
	v6240 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6211)+24)) = v6240
	v6242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6243 = *(*int32)(unsafe.Add(mBase, uint32(v6242)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6211)+32)) = v6243
	v6245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6211)+36)) = uint8(v6245)
	v6247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6211)+37)) = uint8(v6247)
	v12585 = v6211
	v12589 = v47
	goto L3
L917:
	;
	v6254 = int32(0)
	v6255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6256 = *(*int32)(unsafe.Add(mBase, uint32(v6255)+4))
	if v6256 == v6254 {
		v6338 = v6254
		goto L918
	} else {
		goto L919
	}
L918:
	;
	v6380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6381 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v6382 = m.ExcPending
	if v6382 != 0 {
		goto L1
	} else {
		goto L933
	}
L919:
	;
	v6259 = *(*int32)(unsafe.Add(mBase, uint32(v6256)+4))
	if v6259 <= int32(0) {
		v6338 = v6254
		goto L918
	} else {
		goto L920
	}
L920:
	;
	v6262 = *(*int32)(unsafe.Add(mBase, uint32(v6255)+8))
	v6267 = v6254
	v6268 = v6249
	v6270 = v4
	goto L921
L921:
	;
	v6309 = *(*int32)(unsafe.Add(mBase, uint32(v6256)+12))
	v6313 = *(*int32)(unsafe.Add(mBase, uint32(v6309+v6270<<(uint(int32(2))%32))))
	v6314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6314 != 0 {
		goto L923
	} else {
		goto L924
	}
L922:
	;
	v6338 = v6330
	goto L918
L923:
	;
	v6315 = F_replace_nestloop_params_mutator(m, v6313, l0)
	mBase = m.M
	v6316 = m.ExcPending
	if v6316 != 0 {
		goto L1
	} else {
		goto L926
	}
L924:
	;
	v6317 = v6313
	goto L925
L925:
	;
	v6319 = int32(0)
	v6321 = F_makeTargetEntry(m, v6317, base.I32_extend16_s(v6268), v6319, v6319)
	mBase = m.M
	v6322 = m.ExcPending
	if v6322 != 0 {
		goto L1
	} else {
		goto L927
	}
L926:
	;
	v6317 = v6315
	goto L925
L927:
	;
	if v6262 != 0 {
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v6326 = *(*int32)(unsafe.Add(mBase, uint32(v6262-int32(4)+v6268<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v6321)+16)) = v6326
	goto L930
L929:
	;
	goto L930
L930:
	;
	v6330 = F_lappend(m, v6267, v6321)
	mBase = m.M
	v6331 = m.ExcPending
	if v6331 != 0 {
		goto L1
	} else {
		goto L931
	}
L931:
	;
	v6333 = v6270 + int32(1)
	v6334 = *(*int32)(unsafe.Add(mBase, uint32(v6256)+4))
	if v6333 < v6334 {
		v6267 = v6330
		v6268 = v6268 + int32(1)
		v6270 = v6333
		goto L921
	} else {
		goto L932
	}
L932:
	;
	goto L922
L933:
	;
	v6383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v6385 = F_palloc0(m, int32(88))
	mBase = m.M
	v6386 = m.ExcPending
	if v6386 != 0 {
		goto L1
	} else {
		goto L934
	}
L934:
	;
	v6387 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+84)) = v6387
	*(*uint8)(unsafe.Add(mBase, uint32(v6385)+81)) = uint8(v6387)
	*(*uint8)(unsafe.Add(mBase, uint32(v6385)+80)) = uint8(v6383)
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+76)) = v6381
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+72)) = v6380
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+56)) = v6387
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+52)) = v6252
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+48)) = v6387
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+44)) = v6338
	*(*int32)(unsafe.Add(mBase, uint32(v6385))) = int32(368)
	v6402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+4)) = v6402
	v6404 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6385)+8)) = v6404
	v6406 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6385)+16)) = v6406
	v6408 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6385)+24)) = v6408
	v6410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6411 = *(*int32)(unsafe.Add(mBase, uint32(v6410)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6385)+32)) = v6411
	v6413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6385)+36)) = uint8(v6413)
	v6415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6385)+37)) = uint8(v6415)
	v6417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6417)+83)) = uint8(v6418)
	v12585 = v6385
	v12589 = v47
	goto L3
L935:
	;
	v6423 = m.G0
	v6425 = v6423 - int32(16)
	m.G0 = v6425
	v6427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6430 = F_create_plan_recurse(m, l0, v6427, l2|int32(4))
	mBase = m.M
	v6431 = m.ExcPending
	if v6431 != 0 {
		goto L1
	} else {
		goto L939
	}
L936:
	;
	goto L937
L937:
	;
	v6811 = m.G0
	v6813 = v6811 - int32(112)
	m.G0 = v6813
	v6815 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6816 = F_create_plan_recurse(m, l0, v6815, l2)
	mBase = m.M
	v6817 = m.ExcPending
	if v6817 != 0 {
		goto L1
	} else {
		goto L983
	}
L938:
	;
	v12585 = v6435
	v12589 = v47
	goto L3
L939:
	;
	v6432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6435 = F_palloc0(m, int32(88))
	mBase = m.M
	v6436 = m.ExcPending
	if v6436 != 0 {
		goto L1
	} else {
		goto L940
	}
L940:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435))) = int32(367)
	v6439 = *(*int32)(unsafe.Add(mBase, uint32(v6430)+44))
	v6440 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+56)) = v6440
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+52)) = v6430
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+48)) = v6440
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+44)) = v6439
	v6448 = F_palloc(m, v6433<<(uint(int32(1))%32))
	mBase = m.M
	v6449 = m.ExcPending
	if v6449 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	v6451 = v6433 << (uint(int32(2)) % 32)
	v6452 = F_palloc(m, v6451)
	mBase = m.M
	v6453 = m.ExcPending
	if v6453 != 0 {
		goto L1
	} else {
		goto L942
	}
L942:
	;
	v6454 = F_palloc(m, v6451)
	mBase = m.M
	v6455 = m.ExcPending
	if v6455 != 0 {
		goto L1
	} else {
		goto L943
	}
L943:
	;
	if v6432 == int32(0) {
		goto L946
	} else {
		goto L947
	}
L944:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6801 = m.ExcPending
	if v6801 != 0 {
		goto L1
	} else {
		goto L979
	}
L945:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6739 = m.ExcPending
	if v6739 != 0 {
		goto L1
	} else {
		goto L976
	}
L946:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+84)) = v6454
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+80)) = v6452
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+76)) = v6448
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+72)) = v6433
	v6718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+4)) = v6718
	v6720 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6435)+8)) = v6720
	v6722 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6435)+16)) = v6722
	v6724 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6435)+24)) = v6724
	v6726 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6727 = *(*int32)(unsafe.Add(mBase, uint32(v6726)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+32)) = v6727
	v6729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6435)+36)) = uint8(v6729)
	v6731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6435)+37)) = uint8(v6731)
	m.G0 = v6425 + int32(16)
	goto L938
L947:
	;
	v6458 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+4))
	if v6458 <= int32(0) {
		goto L946
	} else {
		goto L948
	}
L948:
	;
	v6461 = int32(0)
	if v6461 < v6433 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	v6464 = v6433
	goto L951
L950:
	;
	v6464 = v6461
	goto L951
L951:
	;
	v6474 = v4
	goto L952
L952:
	;
	if v6474 == v6464 {
		goto L946
	} else {
		goto L954
	}
L953:
	;
	goto L946
L954:
	;
	v6511 = v6474 << (uint(int32(2)) % 32)
	v6512 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+12))
	v6514 = *(*int32)(unsafe.Add(mBase, uint32(v6511+v6512)))
	v6515 = *(*int32)(unsafe.Add(mBase, uint32(v6514)+4))
	v6516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6515)+41)))
	if v6516 == int32(1) {
		goto L957
	} else {
		goto L958
	}
L955:
	;
	v6649 = *(*int32)(unsafe.Add(mBase, uint32(v6514)+8))
	v6650 = *(*int32)(unsafe.Add(mBase, uint32(v6621)+16))
	v6652 = F_get_opfamily_member_for_cmptype(m, v6649, v6650, v6650, int32(3))
	mBase = m.M
	v6653 = m.ExcPending
	if v6653 != 0 {
		goto L1
	} else {
		goto L973
	}
L956:
	;
	v6597 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+44))
	v6598 = F_get_sortgroupref_tle(m, v6519, v6597)
	mBase = m.M
	v6599 = m.ExcPending
	if v6599 != 0 {
		goto L1
	} else {
		goto L971
	}
L957:
	;
	v6519 = *(*int32)(unsafe.Add(mBase, uint32(v6515)+44))
	if v6519 != 0 {
		goto L956
	} else {
		goto L960
	}
L958:
	;
	goto L959
L959:
	;
	v6533 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+44))
	if v6533 == int32(0) {
		goto L944
	} else {
		goto L964
	}
L960:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6523 = m.ExcPending
	if v6523 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	F_errmsg_internal(m, int32(340097), int32(0))
	mBase = m.M
	v6527 = m.ExcPending
	if v6527 != 0 {
		goto L1
	} else {
		goto L962
	}
L962:
	;
	F_errfinish(m, int32(498183), int32(6931), int32(113104))
	mBase = m.M
	v6532 = m.ExcPending
	if v6532 != 0 {
		goto L1
	} else {
		goto L963
	}
L963:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L964:
	;
	v6536 = int32(0)
	v6537 = *(*int32)(unsafe.Add(mBase, uint32(v6533)+4))
	if v6537 <= v6536 {
		goto L944
	} else {
		goto L965
	}
L965:
	;
	v6542 = v6536
	goto L966
L966:
	;
	v6584 = *(*int32)(unsafe.Add(mBase, uint32(v6533)+12))
	v6588 = *(*int32)(unsafe.Add(mBase, uint32(v6584+v6542<<(uint(int32(2))%32))))
	v6589 = *(*int32)(unsafe.Add(mBase, uint32(v6588)+4))
	v6591 = F_find_ec_member_matching_expr(m, v6515, v6589, int32(0))
	mBase = m.M
	v6592 = m.ExcPending
	if v6592 != 0 {
		goto L1
	} else {
		goto L968
	}
L967:
	;
	goto L944
L968:
	;
	if v6591 != 0 {
		v6619 = v6588
		v6621 = v6591
		goto L955
	} else {
		goto L969
	}
L969:
	;
	v6594 = v6542 + int32(1)
	v6595 = *(*int32)(unsafe.Add(mBase, uint32(v6533)+4))
	if v6594 < v6595 {
		v6542 = v6594
		goto L966
	} else {
		goto L970
	}
L970:
	;
	goto L967
L971:
	;
	if v6598 == int32(0) {
		goto L944
	} else {
		goto L972
	}
L972:
	;
	v6602 = *(*int32)(unsafe.Add(mBase, uint32(v6515)+16))
	v6603 = *(*int32)(unsafe.Add(mBase, uint32(v6602)+12))
	v6604 = *(*int32)(unsafe.Add(mBase, uint32(v6603)))
	v6619 = v6598
	v6621 = v6604
	goto L955
L973:
	;
	if v6652 == int32(0) {
		goto L945
	} else {
		goto L974
	}
L974:
	;
	v6656 = int32(1)
	v6659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6619)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v6448+v6474<<(uint(v6656)%32)))) = uint16(v6659)
	*(*int32)(unsafe.Add(mBase, uint32(v6511+v6452))) = v6652
	v6664 = *(*int32)(unsafe.Add(mBase, uint32(v6515)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6511+v6454))) = v6664
	v6667 = v6474 + v6656
	v6668 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+4))
	if v6667 < v6668 {
		v6474 = v6667
		goto L952
	} else {
		goto L975
	}
L975:
	;
	goto L953
L976:
	;
	v6740 = *(*int32)(unsafe.Add(mBase, uint32(v6514)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6425)+12)) = v6740
	*(*int32)(unsafe.Add(mBase, uint32(v6425)+8)) = v6650
	*(*int32)(unsafe.Add(mBase, uint32(v6425)+4)) = v6650
	*(*int32)(unsafe.Add(mBase, uint32(v6425))) = int32(3)
	F_errmsg_internal(m, int32(39996), v6425)
	mBase = m.M
	v6748 = m.ExcPending
	if v6748 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	F_errfinish(m, int32(498183), int32(6972), int32(113104))
	mBase = m.M
	v6753 = m.ExcPending
	if v6753 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L979:
	;
	F_errmsg_internal(m, int32(78942), int32(0))
	mBase = m.M
	v6805 = m.ExcPending
	if v6805 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	F_errfinish(m, int32(498183), int32(6959), int32(113104))
	mBase = m.M
	v6810 = m.ExcPending
	if v6810 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L982:
	;
	v12585 = v8409
	v12589 = v47
	goto L3
L983:
	;
	v6818 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v6818 == int32(0) {
		goto L989
	} else {
		goto L990
	}
L984:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8498 = m.ExcPending
	if v8498 != 0 {
		goto L1
	} else {
		goto L1191
	}
L985:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8483 = m.ExcPending
	if v8483 != 0 {
		goto L1
	} else {
		goto L1188
	}
L986:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8470 = m.ExcPending
	if v8470 != 0 {
		goto L1
	} else {
		goto L1185
	}
L987:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8457 = m.ExcPending
	if v8457 != 0 {
		goto L1
	} else {
		goto L1182
	}
L988:
	;
	m.G0 = v6813 + int32(112)
	goto L982
L989:
	;
	v8409 = v6816
	goto L988
L990:
	;
	goto L991
L991:
	;
	v6821 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v6822 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6823 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6824 = *(*int32)(unsafe.Add(mBase, uint32(v6823)+4))
	if v6824 == int32(0) {
		goto L993
	} else {
		goto L994
	}
L992:
	;
	if v6821 != 0 {
		goto L1013
	} else {
		goto L1014
	}
L993:
	;
	v6960 = int32(1)
	v6962 = int32(0)
	goto L992
L994:
	;
	v6827 = int32(1)
	v6828 = *(*int32)(unsafe.Add(mBase, uint32(v6824)+4))
	if v6828 <= int32(0) {
		goto L995
	} else {
		goto L996
	}
L995:
	;
	v6960 = v6827
	v6962 = int32(0)
	goto L992
L996:
	;
	goto L997
L997:
	;
	v6832 = *(*int32)(unsafe.Add(mBase, uint32(v6823)+8))
	v6838 = v6827
	v6840 = int32(0)
	v6846 = v4
	goto L998
L998:
	;
	v6880 = *(*int32)(unsafe.Add(mBase, uint32(v6824)+12))
	v6884 = *(*int32)(unsafe.Add(mBase, uint32(v6880+v6846<<(uint(int32(2))%32))))
	v6885 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6885 != 0 {
		goto L1000
	} else {
		goto L1001
	}
L999:
	;
	if v6901 == int32(0) {
		goto L993
	} else {
		goto L1010
	}
L1000:
	;
	v6886 = F_replace_nestloop_params_mutator(m, v6884, l0)
	mBase = m.M
	v6887 = m.ExcPending
	if v6887 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1001:
	;
	v6888 = v6884
	goto L1002
L1002:
	;
	v6890 = int32(0)
	v6892 = F_makeTargetEntry(m, v6888, base.I32_extend16_s(v6838), v6890, v6890)
	mBase = m.M
	v6893 = m.ExcPending
	if v6893 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1003:
	;
	v6888 = v6886
	goto L1002
L1004:
	;
	if v6832 != 0 {
		goto L1005
	} else {
		goto L1006
	}
L1005:
	;
	v6897 = *(*int32)(unsafe.Add(mBase, uint32(v6832-int32(4)+v6838<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v6892)+16)) = v6897
	goto L1007
L1006:
	;
	goto L1007
L1007:
	;
	v6901 = F_lappend(m, v6840, v6892)
	mBase = m.M
	v6902 = m.ExcPending
	if v6902 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	v6904 = v6846 + int32(1)
	v6905 = *(*int32)(unsafe.Add(mBase, uint32(v6824)+4))
	if v6904 < v6905 {
		v6838 = v6838 + int32(1)
		v6840 = v6901
		v6846 = v6904
		goto L998
	} else {
		goto L1009
	}
L1009:
	;
	goto L999
L1010:
	;
	v6909 = *(*int32)(unsafe.Add(mBase, uint32(v6901)+4))
	v6960 = v6909 + int32(1)
	v6962 = v6901
	goto L992
L1011:
	;
	v7262 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+44))
	if v6821 != 0 {
		goto L1032
	} else {
		goto L1033
	}
L1012:
	;
	v7215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v7216 = F_change_plan_targetlist(m, v6816, v7175, v7215)
	mBase = m.M
	v7217 = m.ExcPending
	if v7217 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1013:
	;
	v7002 = int32(0)
	v7003 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+4))
	if v7003 <= v7002 {
		goto L1017
	} else {
		goto L1018
	}
L1014:
	;
	v7128 = v6962
	goto L1015
L1015:
	;
	v7168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v7168 != int32(2) {
		v7218 = v6816
		goto L1011
	} else {
		goto L1030
	}
L1016:
	;
	if v7084&int32(1) != 0 {
		v7175 = v7082
		goto L1012
	} else {
		goto L1029
	}
L1017:
	;
	v7082 = v6962
	v7084 = int32(0)
	goto L1016
L1018:
	;
	goto L1019
L1019:
	;
	v7010 = v6960
	v7011 = v7002
	v7012 = v6962
	v7014 = int32(0)
	goto L1020
L1020:
	;
	v7052 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+12))
	v7056 = *(*int32)(unsafe.Add(mBase, uint32(v7052+v7011<<(uint(int32(2))%32))))
	v7057 = F_tlist_member(m, v7056, v7012)
	mBase = m.M
	v7058 = m.ExcPending
	if v7058 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1021:
	;
	v7082 = v7072
	v7084 = v7073
	goto L1016
L1022:
	;
	if v7057 == int32(0) {
		goto L1023
	} else {
		goto L1024
	}
L1023:
	;
	v7063 = int32(0)
	v7065 = F_makeTargetEntry(m, v7056, base.I32_extend16_s(v7010), v7063, v7063)
	mBase = m.M
	v7066 = m.ExcPending
	if v7066 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1024:
	;
	v7071 = v7010
	v7072 = v7012
	v7073 = v7014
	goto L1025
L1025:
	;
	v7075 = v7011 + int32(1)
	v7076 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+4))
	if v7075 < v7076 {
		v7010 = v7071
		v7011 = v7075
		v7012 = v7072
		v7014 = v7073
		goto L1020
	} else {
		goto L1028
	}
L1026:
	;
	v7067 = F_lappend(m, v7012, v7065)
	mBase = m.M
	v7068 = m.ExcPending
	if v7068 != 0 {
		goto L1
	} else {
		goto L1027
	}
L1027:
	;
	v7071 = v7010 + int32(1)
	v7072 = v7067
	v7073 = int32(1)
	goto L1025
L1028:
	;
	goto L1021
L1029:
	;
	v7128 = v7082
	goto L1015
L1030:
	;
	v7175 = v7128
	goto L1012
L1031:
	;
	v7218 = v7216
	goto L1011
L1032:
	;
	v7263 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+4))
	v7265 = v7263
	goto L1034
L1033:
	;
	v7265 = int32(0)
	goto L1034
L1034:
	;
	v7268 = F_palloc(m, v7265<<(uint(int32(1))%32))
	mBase = m.M
	v7269 = m.ExcPending
	if v7269 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	v7271 = v7265 << (uint(int32(2)) % 32)
	v7272 = F_palloc(m, v7271)
	mBase = m.M
	v7273 = m.ExcPending
	if v7273 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1036:
	;
	if v6821 == int32(0) {
		goto L1037
	} else {
		goto L1038
	}
L1037:
	;
	v7391 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v7391 != int32(1) {
		goto L1047
	} else {
		goto L1048
	}
L1038:
	;
	v7276 = int32(0)
	v7277 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+4))
	if v7277 <= v7276 {
		goto L1037
	} else {
		goto L1039
	}
L1039:
	;
	v7282 = v7276
	goto L1040
L1040:
	;
	v7325 = v7282 << (uint(int32(2)) % 32)
	v7326 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+12))
	v7328 = *(*int32)(unsafe.Add(mBase, uint32(v7325+v7326)))
	v7329 = F_tlist_member(m, v7328, v7262)
	mBase = m.M
	v7330 = m.ExcPending
	if v7330 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1041:
	;
	goto L1037
L1042:
	;
	if v7329 == int32(0) {
		goto L987
	} else {
		goto L1043
	}
L1043:
	;
	v7336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7329)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7268+v7282<<(uint(int32(1))%32)))) = uint16(v7336)
	v7339 = *(*int32)(unsafe.Add(mBase, uint32(v7329)+4))
	v7340 = F_exprCollation(m, v7339)
	mBase = m.M
	v7341 = m.ExcPending
	if v7341 != 0 {
		goto L1
	} else {
		goto L1044
	}
L1044:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7325+v7272))) = v7340
	v7344 = v7282 + int32(1)
	v7345 = *(*int32)(unsafe.Add(mBase, uint32(v6821)+4))
	if v7344 < v7345 {
		v7282 = v7344
		goto L1040
	} else {
		goto L1045
	}
L1045:
	;
	goto L1041
L1046:
	;
	v8392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8350)+4)) = v8392
	v8394 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8350)+8)) = v8394
	v8396 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8350)+16)) = v8396
	v8398 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8350)+24)) = v8398
	v8400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8401 = *(*int32)(unsafe.Add(mBase, uint32(v8400)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8350)+32)) = v8401
	v8403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8350)+36)) = uint8(v8403)
	v8405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8350)+37)) = uint8(v8405)
	v8409 = v8350
	goto L988
L1047:
	;
	v7394 = int32(0)
	if v6822 == v7394 {
		v7672 = v7394
		goto L1050
	} else {
		goto L1051
	}
L1048:
	;
	goto L1049
L1049:
	;
	v8061 = F_palloc(m, v7271)
	mBase = m.M
	v8062 = m.ExcPending
	if v8062 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1050:
	;
	v7707 = int32(0)
	v7708 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+44))
	if v7672 != 0 {
		goto L1112
	} else {
		goto L1113
	}
L1051:
	;
	v7398 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+4))
	if v7398 <= int32(0) {
		v7672 = v7394
		goto L1050
	} else {
		goto L1052
	}
L1052:
	;
	v7406 = int32(0)
	v7411 = v7394
	goto L1053
L1053:
	;
	v7446 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+12))
	v7450 = *(*int32)(unsafe.Add(mBase, uint32(v7446+v7406<<(uint(int32(2))%32))))
	v7451 = F_get_ordering_op_for_equality_op(m, v7450)
	mBase = m.M
	v7452 = m.ExcPending
	if v7452 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1054:
	;
	v7672 = v7657
	goto L1050
L1055:
	;
	if v7451 == int32(0) {
		goto L985
	} else {
		goto L1056
	}
L1056:
	;
	v7456 = F_get_equality_op_for_ordering_op(m, v7451, int32(0))
	mBase = m.M
	v7457 = m.ExcPending
	if v7457 != 0 {
		goto L1
	} else {
		goto L1057
	}
L1057:
	;
	if v7456 == int32(0) {
		goto L984
	} else {
		goto L1058
	}
L1058:
	;
	v7460 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+44))
	v7464 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7268+v7406<<(uint(int32(1))%32)))))
	if v7460 != 0 {
		goto L1061
	} else {
		goto L1062
	}
L1059:
	;
	v7505 = F_palloc0(m, int32(20))
	mBase = m.M
	v7506 = m.ExcPending
	if v7506 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1060:
	;
	goto L1059
L1061:
	;
	v7468 = *(*int32)(unsafe.Add(mBase, uint32(v7460)+4))
	if v7468 <= int32(0) {
		v7502 = int32(0)
		goto L1060
	} else {
		goto L1064
	}
L1062:
	;
	goto L1063
L1063:
	;
	v7502 = int32(0)
	goto L1060
L1064:
	;
	v7471 = int32(0)
	if v7471 < v7468 {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	v7474 = v7468
	goto L1067
L1066:
	;
	v7474 = v7471
	goto L1067
L1067:
	;
	v7475 = *(*int32)(unsafe.Add(mBase, uint32(v7460)+12))
	v7479 = int32(0)
	goto L1068
L1068:
	;
	v7487 = *(*int32)(unsafe.Add(mBase, uint32(v7475+v7479<<(uint(int32(2))%32))))
	v7488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7487)+8)))
	if v7488 == v7464&int32(65535) {
		v7502 = v7487
		goto L1060
	} else {
		goto L1070
	}
L1069:
	;
	goto L1063
L1070:
	;
	v7491 = v7479 + int32(1)
	if v7491 != v7474 {
		v7479 = v7491
		goto L1068
	} else {
		goto L1071
	}
L1071:
	;
	goto L1069
L1072:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7505))) = int32(106)
	v7509 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+44))
	v7510 = int32(0)
	v7519 = *(*int32)(unsafe.Add(mBase, uint32(v7502)+16))
	if v7519 == v7510 {
		goto L1074
	} else {
		goto L1075
	}
L1073:
	;
	v7650 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7505)+18)) = uint8(v7650)
	*(*uint16)(unsafe.Add(mBase, uint32(v7505)+16)) = uint16(v7650)
	*(*int32)(unsafe.Add(mBase, uint32(v7505)+12)) = v7451
	*(*int32)(unsafe.Add(mBase, uint32(v7505)+8)) = v7456
	*(*int32)(unsafe.Add(mBase, uint32(v7505)+4)) = v7641
	v7657 = F_lappend(m, v7411, v7505)
	mBase = m.M
	v7658 = m.ExcPending
	if v7658 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1074:
	;
	if v7509 == int32(0) {
		v7637 = int32(1)
		goto L1077
	} else {
		goto L1078
	}
L1075:
	;
	v7641 = v7519
	goto L1076
L1076:
	;
	goto L1073
L1077:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7502)+16)) = v7637
	v7641 = v7637
	goto L1076
L1078:
	;
	v7526 = *(*int32)(unsafe.Add(mBase, uint32(v7509)+4))
	if v7526 <= int32(0) {
		v7637 = int32(1)
		goto L1077
	} else {
		goto L1079
	}
L1079:
	;
	v7529 = int32(0)
	if v7529 < v7526 {
		goto L1080
	} else {
		goto L1081
	}
L1080:
	;
	v7532 = v7526
	goto L1082
L1081:
	;
	v7532 = v7529
	goto L1082
L1082:
	;
	v7534 = v7532 & int32(3)
	v7535 = int32(0)
	if int32(4) <= v7526 {
		goto L1083
	} else {
		goto L1084
	}
L1083:
	;
	v7540 = *(*int32)(unsafe.Add(mBase, uint32(v7509)+12))
	v7544 = v7535
	v7545 = v7510
	v7546 = int32(0)
	goto L1086
L1084:
	;
	v7579 = v7535
	v7580 = v7510
	goto L1085
L1085:
	;
	if v7534 != 0 {
		goto L1101
	} else {
		goto L1102
	}
L1086:
	;
	v7555 = v7540 + v7545<<(uint(int32(2))%32)
	v7556 = *(*int32)(unsafe.Add(mBase, uint32(v7555)+12))
	v7557 = *(*int32)(unsafe.Add(mBase, uint32(v7556)+16))
	v7558 = *(*int32)(unsafe.Add(mBase, uint32(v7555)+8))
	v7559 = *(*int32)(unsafe.Add(mBase, uint32(v7558)+16))
	v7560 = *(*int32)(unsafe.Add(mBase, uint32(v7555)+4))
	v7561 = *(*int32)(unsafe.Add(mBase, uint32(v7560)+16))
	v7562 = *(*int32)(unsafe.Add(mBase, uint32(v7555)))
	v7563 = *(*int32)(unsafe.Add(mBase, uint32(v7562)+16))
	if base.Ui32(v7544) < base.Ui32(v7563) {
		goto L1088
	} else {
		goto L1089
	}
L1087:
	;
	v7579 = v7571
	v7580 = v7573
	goto L1085
L1088:
	;
	v7565 = v7563
	goto L1090
L1089:
	;
	v7565 = v7544
	goto L1090
L1090:
	;
	if base.Ui32(v7565) < base.Ui32(v7561) {
		goto L1091
	} else {
		goto L1092
	}
L1091:
	;
	v7567 = v7561
	goto L1093
L1092:
	;
	v7567 = v7565
	goto L1093
L1093:
	;
	if base.Ui32(v7567) < base.Ui32(v7559) {
		goto L1094
	} else {
		goto L1095
	}
L1094:
	;
	v7569 = v7559
	goto L1096
L1095:
	;
	v7569 = v7567
	goto L1096
L1096:
	;
	if base.Ui32(v7569) < base.Ui32(v7557) {
		goto L1097
	} else {
		goto L1098
	}
L1097:
	;
	v7571 = v7557
	goto L1099
L1098:
	;
	v7571 = v7569
	goto L1099
L1099:
	;
	v7572 = int32(4)
	v7573 = v7545 + v7572
	v7575 = v7546 + v7572
	if v7575 != v7532&int32(2147483644) {
		v7544 = v7571
		v7545 = v7573
		v7546 = v7575
		goto L1086
	} else {
		goto L1100
	}
L1100:
	;
	goto L1087
L1101:
	;
	v7588 = *(*int32)(unsafe.Add(mBase, uint32(v7509)+12))
	v7590 = int32(0)
	v7592 = v7579
	v7593 = v7580
	goto L1104
L1102:
	;
	v7615 = v7579
	goto L1103
L1103:
	;
	v7637 = v7615 + int32(1)
	goto L1077
L1104:
	;
	v7604 = *(*int32)(unsafe.Add(mBase, uint32(v7588+v7593<<(uint(int32(2))%32))))
	v7605 = *(*int32)(unsafe.Add(mBase, uint32(v7604)+16))
	if base.Ui32(v7592) < base.Ui32(v7605) {
		goto L1106
	} else {
		goto L1107
	}
L1105:
	;
	v7615 = v7607
	goto L1103
L1106:
	;
	v7607 = v7605
	goto L1108
L1107:
	;
	v7607 = v7592
	goto L1108
L1108:
	;
	v7608 = int32(1)
	v7611 = v7590 + v7608
	if v7611 != v7534 {
		v7590 = v7611
		v7592 = v7607
		v7593 = v7593 + v7608
		goto L1104
	} else {
		goto L1109
	}
L1109:
	;
	goto L1105
L1110:
	;
	v7660 = v7406 + int32(1)
	v7661 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+4))
	if v7660 < v7661 {
		v7406 = v7660
		v7411 = v7657
		goto L1053
	} else {
		goto L1111
	}
L1111:
	;
	goto L1054
L1112:
	;
	v7709 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	v7710 = v7709
	goto L1114
L1113:
	;
	v7710 = v4
	goto L1114
L1114:
	;
	v7713 = F_palloc(m, v7710<<(uint(int32(1))%32))
	mBase = m.M
	v7714 = m.ExcPending
	if v7714 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	v7716 = v7710 << (uint(int32(2)) % 32)
	v7717 = F_palloc(m, v7716)
	mBase = m.M
	v7718 = m.ExcPending
	if v7718 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1116:
	;
	v7719 = F_palloc(m, v7716)
	mBase = m.M
	v7720 = m.ExcPending
	if v7720 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	v7721 = F_palloc(m, v7710)
	mBase = m.M
	v7722 = m.ExcPending
	if v7722 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	if v7672 == int32(0) {
		v7805 = v7707
		goto L1119
	} else {
		goto L1120
	}
L1119:
	;
	v7844 = F_palloc0(m, int32(96))
	mBase = m.M
	v7845 = m.ExcPending
	if v7845 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1120:
	;
	v7725 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	if v7725 <= int32(0) {
		v7805 = v7707
		goto L1119
	} else {
		goto L1121
	}
L1121:
	;
	v7734 = v7707
	goto L1122
L1122:
	;
	v7776 = v7734 << (uint(int32(2)) % 32)
	v7777 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+12))
	v7779 = *(*int32)(unsafe.Add(mBase, uint32(v7776+v7777)))
	v7780 = F_get_sortgroupclause_tle(m, v7779, v7708)
	mBase = m.M
	v7781 = m.ExcPending
	if v7781 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1123:
	;
	v7805 = v7796
	goto L1119
L1124:
	;
	v7782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7780)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7713+v7734<<(uint(int32(1))%32)))) = uint16(v7782)
	v7785 = *(*int32)(unsafe.Add(mBase, uint32(v7779)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7717+v7776))) = v7785
	v7788 = *(*int32)(unsafe.Add(mBase, uint32(v7780)+4))
	v7789 = F_exprCollation(m, v7788)
	mBase = m.M
	v7790 = m.ExcPending
	if v7790 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7719+v7776))) = v7789
	v7793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7779)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7734+v7721))) = uint8(v7793)
	v7796 = v7734 + int32(1)
	v7797 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	if v7796 < v7797 {
		v7734 = v7796
		goto L1122
	} else {
		goto L1126
	}
L1126:
	;
	goto L1123
L1127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7844))) = int32(362)
	v7848 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+44)) = v7848
	v7850 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+4))
	v7851 = int32(4126313)
	v7852 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+88)) = v7721
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+84)) = v7719
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+80)) = v7717
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+76)) = v7713
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+72)) = v7805
	v7858 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+56)) = v7858
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+52)) = v7218
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+48)) = v7858
	v7863 = int32(1)
	v7865 = v7850 + (v7852 ^ v7863)
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+4)) = v7865
	v7868 = v6813 + int32(40)
	v7871 = *(*float64)(unsafe.Add(mBase, uint32(v7218)+16))
	v7872 = *(*float64)(unsafe.Add(mBase, uint32(v7218)+24))
	v7873 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+32))
	v7876 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v7879 = m.G0
	v7880 = int32(16)
	v7881 = v7879 - v7880
	m.G0 = v7881
	F_cost_tuplesort(m, v7881+int32(8), v7881, v7872, v7873, float64(0), v7876, float64(-1))
	mBase = m.M
	v7886 = *(*float64)(unsafe.Add(mBase, uint32(v7881)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v7868)+32)) = v7872
	v7889 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	v7890 = base.F64_add(v7871, v7886)
	*(*float64)(unsafe.Add(mBase, uint32(v7868)+48)) = v7890
	*(*int32)(unsafe.Add(mBase, uint32(v7868)+40)) = v7865 + (v7889 ^ v7863)
	v7896 = *(*float64)(unsafe.Add(mBase, uint32(v7881)))
	*(*float64)(unsafe.Add(mBase, uint32(v7868)+56)) = base.F64_add(v7890, v7896)
	m.G0 = v7881 + v7880
	goto L1128
L1128:
	;
	v7902 = *(*float64)(unsafe.Add(mBase, uint32(v6813)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v7844)+8)) = v7902
	v7904 = *(*float64)(unsafe.Add(mBase, uint32(v6813)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v7844)+16)) = v7904
	v7906 = *(*float64)(unsafe.Add(mBase, uint32(v7218)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v7844)+24)) = v7906
	v7908 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+32))
	v7909 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7844)+36)) = uint8(v7909)
	*(*int32)(unsafe.Add(mBase, uint32(v7844)+32)) = v7908
	v7912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7218)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7844)+37)) = uint8(v7912)
	v7915 = F_palloc0(m, int32(88))
	mBase = m.M
	v7916 = m.ExcPending
	if v7916 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7915))) = int32(367)
	if v7672 != 0 {
		goto L1130
	} else {
		goto L1131
	}
L1130:
	;
	v7919 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	v7920 = v7919
	goto L1132
L1131:
	;
	v7920 = v7394
	goto L1132
L1132:
	;
	v7921 = *(*int32)(unsafe.Add(mBase, uint32(v7844)+44))
	v7922 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+56)) = v7922
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+52)) = v7844
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+48)) = v7922
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+44)) = v7921
	v7931 = F_palloc(m, v7920<<(uint(int32(1))%32))
	mBase = m.M
	v7932 = m.ExcPending
	if v7932 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	v7934 = v7920 << (uint(int32(2)) % 32)
	v7935 = F_palloc(m, v7934)
	mBase = m.M
	v7936 = m.ExcPending
	if v7936 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	v7937 = F_palloc(m, v7934)
	mBase = m.M
	v7938 = m.ExcPending
	if v7938 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1135:
	;
	if v7672 == int32(0) {
		goto L1136
	} else {
		goto L1137
	}
L1136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+84)) = v7937
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+80)) = v7935
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+76)) = v7931
	*(*int32)(unsafe.Add(mBase, uint32(v7915)+72)) = v7920
	v8350 = v7915
	goto L1046
L1137:
	;
	v7941 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	if v7941 <= int32(0) {
		goto L1136
	} else {
		goto L1138
	}
L1138:
	;
	v7948 = v7922
	goto L1139
L1139:
	;
	v7992 = v7948 << (uint(int32(2)) % 32)
	v7993 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+12))
	v7995 = *(*int32)(unsafe.Add(mBase, uint32(v7992+v7993)))
	v7996 = *(*int32)(unsafe.Add(mBase, uint32(v7915)+44))
	v7997 = F_get_sortgroupclause_tle(m, v7995, v7996)
	mBase = m.M
	v7998 = m.ExcPending
	if v7998 != 0 {
		goto L1
	} else {
		goto L1141
	}
L1140:
	;
	goto L1136
L1141:
	;
	v7999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7997)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7931+v7948<<(uint(int32(1))%32)))) = uint16(v7999)
	v8002 = *(*int32)(unsafe.Add(mBase, uint32(v7995)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7935+v7992))) = v8002
	v8005 = *(*int32)(unsafe.Add(mBase, uint32(v7997)+4))
	v8006 = F_exprCollation(m, v8005)
	mBase = m.M
	v8007 = m.ExcPending
	if v8007 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7992+v7937))) = v8006
	v8010 = v7948 + int32(1)
	v8011 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	if v8010 < v8011 {
		v7948 = v8010
		goto L1139
	} else {
		goto L1143
	}
L1143:
	;
	goto L1140
L1144:
	;
	if v6822 == int32(0) {
		goto L1145
	} else {
		goto L1146
	}
L1145:
	;
	v8175 = int32(0)
	v8176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8177 = *(*int32)(unsafe.Add(mBase, uint32(v8176)+4))
	if v8177 == v8175 {
		v8262 = v8175
		goto L1153
	} else {
		goto L1154
	}
L1146:
	;
	v8065 = int32(0)
	v8066 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+4))
	if v8066 <= v8065 {
		goto L1145
	} else {
		goto L1147
	}
L1147:
	;
	v8071 = v8065
	goto L1148
L1148:
	;
	v8114 = v8071 << (uint(int32(2)) % 32)
	v8115 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+12))
	v8117 = *(*int32)(unsafe.Add(mBase, uint32(v8114+v8115)))
	v8120 = F_get_compatible_hash_operators(m, v8117, v6813+int32(40))
	mBase = m.M
	v8121 = m.ExcPending
	if v8121 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1149:
	;
	goto L1145
L1150:
	;
	if v8120 == int32(0) {
		goto L986
	} else {
		goto L1151
	}
L1151:
	;
	v8125 = *(*int32)(unsafe.Add(mBase, uint32(v6813)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8061+v8114))) = v8125
	v8128 = v8071 + int32(1)
	v8129 = *(*int32)(unsafe.Add(mBase, uint32(v6822)+4))
	if v8128 < v8129 {
		v8071 = v8128
		goto L1148
	} else {
		goto L1152
	}
L1152:
	;
	goto L1149
L1153:
	;
	v8303 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v8305 = F_palloc0(m, int32(128))
	mBase = m.M
	v8306 = m.ExcPending
	if v8306 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1154:
	;
	v8181 = int32(0)
	v8182 = *(*int32)(unsafe.Add(mBase, uint32(v8177)+4))
	if v8182 <= v8181 {
		v8262 = v8175
		goto L1153
	} else {
		goto L1155
	}
L1155:
	;
	v8185 = *(*int32)(unsafe.Add(mBase, uint32(v8176)+8))
	v8190 = int32(1)
	v8191 = v8175
	v8192 = v8181
	goto L1156
L1156:
	;
	v8232 = *(*int32)(unsafe.Add(mBase, uint32(v8177)+12))
	v8236 = *(*int32)(unsafe.Add(mBase, uint32(v8232+v8192<<(uint(int32(2))%32))))
	v8237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8237 != 0 {
		goto L1158
	} else {
		goto L1159
	}
L1157:
	;
	v8262 = v8253
	goto L1153
L1158:
	;
	v8238 = F_replace_nestloop_params_mutator(m, v8236, l0)
	mBase = m.M
	v8239 = m.ExcPending
	if v8239 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1159:
	;
	v8240 = v8236
	goto L1160
L1160:
	;
	v8242 = int32(0)
	v8244 = F_makeTargetEntry(m, v8240, base.I32_extend16_s(v8190), v8242, v8242)
	mBase = m.M
	v8245 = m.ExcPending
	if v8245 != 0 {
		goto L1
	} else {
		goto L1162
	}
L1161:
	;
	v8240 = v8238
	goto L1160
L1162:
	;
	if v8185 != 0 {
		goto L1163
	} else {
		goto L1164
	}
L1163:
	;
	v8249 = *(*int32)(unsafe.Add(mBase, uint32(v8185-int32(4)+v8190<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8244)+16)) = v8249
	goto L1165
L1164:
	;
	goto L1165
L1165:
	;
	v8253 = F_lappend(m, v8191, v8244)
	mBase = m.M
	v8254 = m.ExcPending
	if v8254 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	v8256 = v8192 + int32(1)
	v8257 = *(*int32)(unsafe.Add(mBase, uint32(v8177)+4))
	if v8256 < v8257 {
		v8190 = v8190 + int32(1)
		v8191 = v8253
		v8192 = v8256
		goto L1156
	} else {
		goto L1167
	}
L1167:
	;
	goto L1157
L1168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8305))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v8303)&int64(9223372036854775807)) {
		goto L1170
	} else {
		goto L1171
	}
L1169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+96)) = v8328
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+92)) = v7272
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+88)) = v8061
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+84)) = v7268
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+80)) = v7265
	*(*int64)(unsafe.Add(mBase, uint32(v8305)+72)) = int64(2)
	v8336 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8305)+104)) = v8336
	v8338 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+48)) = v8338
	*(*int64)(unsafe.Add(mBase, uint32(v8305)+112)) = v8336
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+120)) = v8338
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+56)) = v8338
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+52)) = v7218
	*(*int32)(unsafe.Add(mBase, uint32(v8305)+44)) = v8262
	v8350 = v8305
	goto L1046
L1170:
	;
	v8328 = int32(2147483647)
	goto L1169
L1171:
	;
	goto L1172
L1172:
	;
	if base.F64_le(v8303, float64(0)) != 0 {
		goto L1173
	} else {
		goto L1174
	}
L1173:
	;
	v8328 = int32(0)
	goto L1169
L1174:
	;
	goto L1175
L1175:
	;
	v8318 = float64(2.147483647e+09)
	if base.F64_lt(v8303, v8318) != 0 {
		goto L1176
	} else {
		goto L1177
	}
L1176:
	;
	v8321 = v8303
	goto L1178
L1177:
	;
	v8321 = v8318
	goto L1178
L1178:
	;
	if base.F64_lt(base.F64_abs(v8321), float64(2.147483648e+09)) != 0 {
		goto L1179
	} else {
		goto L1180
	}
L1179:
	;
	v8325 = base.I32_trunc_f64_s(v8321)
	v8328 = v8325
	goto L1169
L1180:
	;
	goto L1181
L1181:
	;
	v8328 = int32(-2147483648)
	goto L1169
L1182:
	;
	F_errmsg_internal(m, int32(73907), int32(0))
	mBase = m.M
	v8461 = m.ExcPending
	if v8461 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1183:
	;
	F_errfinish(m, int32(498183), int32(1809), int32(283930))
	mBase = m.M
	v8466 = m.ExcPending
	if v8466 != 0 {
		goto L1
	} else {
		goto L1184
	}
L1184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6813))) = v8117
	F_errmsg_internal(m, int32(43287), v6813)
	mBase = m.M
	v8474 = m.ExcPending
	if v8474 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	F_errfinish(m, int32(498183), int32(1834), int32(283930))
	mBase = m.M
	v8479 = m.ExcPending
	if v8479 != 0 {
		goto L1
	} else {
		goto L1187
	}
L1187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6813)+16)) = v7450
	F_errmsg_internal(m, int32(43139), v6813+int32(16))
	mBase = m.M
	v8489 = m.ExcPending
	if v8489 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	F_errfinish(m, int32(498183), int32(1875), int32(283930))
	mBase = m.M
	v8494 = m.ExcPending
	if v8494 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6813)+32)) = v7451
	F_errmsg_internal(m, int32(43545), v6813+int32(32))
	mBase = m.M
	v8504 = m.ExcPending
	if v8504 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	F_errfinish(m, int32(498183), int32(1886), int32(283930))
	mBase = m.M
	v8509 = m.ExcPending
	if v8509 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1194:
	;
	v8515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v8516 = F_replace_nestloop_params_mutator(m, v8515, l0)
	mBase = m.M
	v8517 = m.ExcPending
	if v8517 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	if v8516 != 0 {
		goto L1196
	} else {
		goto L1197
	}
L1196:
	;
	v8518 = *(*int32)(unsafe.Add(mBase, uint32(v8516)+4))
	v8521 = v8518 << (uint(int32(2)) % 32)
	goto L1198
L1197:
	;
	v8521 = v4
	goto L1198
L1198:
	;
	v8522 = F_palloc(m, v8521)
	mBase = m.M
	v8523 = m.ExcPending
	if v8523 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	v8524 = F_palloc(m, v8521)
	mBase = m.M
	v8525 = m.ExcPending
	if v8525 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1200:
	;
	v8526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v8530 = v4
	goto L1201
L1201:
	;
	v8571 = int32(0)
	if v8516 == v8571 {
		v8580 = v8571
		goto L1203
	} else {
		goto L1204
	}
L1202:
	;
	v8605 = m.G0
	v8607 = v8605 - int32(16)
	m.G0 = v8607
	v8609 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+12)) = v8609
	if v8516 == v8609 {
		v8627 = v8609
		goto L1212
	} else {
		goto L1213
	}
L1203:
	;
	if v8526 == int32(0) {
		goto L1206
	} else {
		goto L1207
	}
L1204:
	;
	v8574 = *(*int32)(unsafe.Add(mBase, uint32(v8516)+4))
	if v8574 <= v8530 {
		v8580 = v8571
		goto L1203
	} else {
		goto L1205
	}
L1205:
	;
	v8576 = *(*int32)(unsafe.Add(mBase, uint32(v8516)+12))
	v8580 = v8576 + v8530<<(uint(int32(2))%32)
	goto L1203
L1206:
	;
	goto L1202
L1207:
	;
	v8583 = *(*int32)(unsafe.Add(mBase, uint32(v8526)+4))
	if v8583 <= v8530 {
		goto L1206
	} else {
		goto L1208
	}
L1208:
	;
	if v8580 == int32(0) {
		goto L1206
	} else {
		goto L1209
	}
L1209:
	;
	v8588 = v8530 << (uint(int32(2)) % 32)
	v8589 = *(*int32)(unsafe.Add(mBase, uint32(v8526)+12))
	v8590 = v8588 + v8589
	if v8590 == int32(0) {
		goto L1206
	} else {
		goto L1210
	}
L1210:
	;
	v8593 = *(*int32)(unsafe.Add(mBase, uint32(v8580)))
	v8595 = *(*int32)(unsafe.Add(mBase, uint32(v8590)))
	*(*int32)(unsafe.Add(mBase, uint32(v8588+v8522))) = v8595
	v8598 = F_exprCollation(m, v8593)
	mBase = m.M
	v8599 = m.ExcPending
	if v8599 != 0 {
		goto L1
	} else {
		goto L1211
	}
L1211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8588+v8524))) = v8598
	v8530 = v8530 + int32(1)
	goto L1201
L1212:
	;
	m.G0 = v8607 + int32(16)
	v8631 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v8632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v8633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v8635 = F_palloc0(m, int32(104))
	mBase = m.M
	v8636 = m.ExcPending
	if v8636 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1213:
	;
	v8614 = *(*int32)(unsafe.Add(mBase, uint32(v8516)))
	if v8614 == int32(8) {
		goto L1214
	} else {
		goto L1215
	}
L1214:
	;
	v8618 = *(*int32)(unsafe.Add(mBase, uint32(v8516)+8))
	v8619 = F_bms_add_member(m, int32(0), v8618)
	mBase = m.M
	v8620 = m.ExcPending
	if v8620 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1215:
	;
	goto L1216
L1216:
	;
	v8624 = F_expression_tree_walker_impl(m, v8516, int32(876), v8607+int32(12))
	mBase = m.M
	v8625 = m.ExcPending
	if v8625 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1217:
	;
	v8627 = v8619
	goto L1212
L1218:
	;
	v8626 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+12))
	v8627 = v8626
	goto L1212
L1219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8635))) = int32(361)
	v8639 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+44))
	v8640 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+56)) = v8640
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+52)) = v8513
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+48)) = v8640
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+44)) = v8639
	if v8516 != 0 {
		goto L1220
	} else {
		goto L1221
	}
L1220:
	;
	v8647 = *(*int32)(unsafe.Add(mBase, uint32(v8516)+4))
	v8648 = v8647
	goto L1222
L1221:
	;
	v8648 = v8640
	goto L1222
L1222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+96)) = v8627
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+92)) = v8631
	*(*uint8)(unsafe.Add(mBase, uint32(v8635)+89)) = uint8(v8632)
	*(*uint8)(unsafe.Add(mBase, uint32(v8635)+88)) = uint8(v8633)
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+84)) = v8516
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+80)) = v8524
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+76)) = v8522
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+72)) = v8648
	v8657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+4)) = v8657
	v8659 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8635)+8)) = v8659
	v8661 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8635)+16)) = v8661
	v8663 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8635)+24)) = v8663
	v8665 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8666 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8635)+32)) = v8666
	v8668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8635)+36)) = uint8(v8668)
	v8670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8635)+37)) = uint8(v8670)
	v12585 = v8635
	v12589 = v47
	goto L3
L1223:
	;
	v8678 = F_palloc0(m, int32(72))
	mBase = m.M
	v8679 = m.ExcPending
	if v8679 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8678))) = int32(360)
	v8682 = *(*int32)(unsafe.Add(mBase, uint32(v8675)+44))
	v8683 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+56)) = v8683
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+52)) = v8675
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+48)) = v8683
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+44)) = v8682
	v8689 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+4)) = v8689
	v8691 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8678)+8)) = v8691
	v8693 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8678)+16)) = v8693
	v8695 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8678)+24)) = v8695
	v8697 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8698 = *(*int32)(unsafe.Add(mBase, uint32(v8697)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8678)+32)) = v8698
	v8700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8678)+36)) = uint8(v8700)
	v8702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8678)+37)) = uint8(v8702)
	v12585 = v8678
	v12589 = v47
	goto L3
L1225:
	;
	v8708 = int32(0)
	v8709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8710 = *(*int32)(unsafe.Add(mBase, uint32(v8709)+4))
	if v8710 == v8708 {
		v8793 = v8708
		goto L1226
	} else {
		goto L1227
	}
L1226:
	;
	v8836 = F_palloc0(m, int32(72))
	mBase = m.M
	v8837 = m.ExcPending
	if v8837 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1227:
	;
	v8714 = *(*int32)(unsafe.Add(mBase, uint32(v8710)+4))
	if v8714 <= int32(0) {
		v8793 = v8708
		goto L1226
	} else {
		goto L1228
	}
L1228:
	;
	v8717 = *(*int32)(unsafe.Add(mBase, uint32(v8709)+8))
	v8722 = v8708
	v8723 = int32(1)
	v8725 = v4
	goto L1229
L1229:
	;
	v8764 = *(*int32)(unsafe.Add(mBase, uint32(v8710)+12))
	v8768 = *(*int32)(unsafe.Add(mBase, uint32(v8764+v8725<<(uint(int32(2))%32))))
	v8769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8769 != 0 {
		goto L1231
	} else {
		goto L1232
	}
L1230:
	;
	v8793 = v8785
	goto L1226
L1231:
	;
	v8770 = F_replace_nestloop_params_mutator(m, v8768, l0)
	mBase = m.M
	v8771 = m.ExcPending
	if v8771 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1232:
	;
	v8772 = v8768
	goto L1233
L1233:
	;
	v8774 = int32(0)
	v8776 = F_makeTargetEntry(m, v8772, base.I32_extend16_s(v8723), v8774, v8774)
	mBase = m.M
	v8777 = m.ExcPending
	if v8777 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1234:
	;
	v8772 = v8770
	goto L1233
L1235:
	;
	if v8717 != 0 {
		goto L1236
	} else {
		goto L1237
	}
L1236:
	;
	v8781 = *(*int32)(unsafe.Add(mBase, uint32(v8717-int32(4)+v8723<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+16)) = v8781
	goto L1238
L1237:
	;
	goto L1238
L1238:
	;
	v8785 = F_lappend(m, v8722, v8776)
	mBase = m.M
	v8786 = m.ExcPending
	if v8786 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	v8788 = v8725 + int32(1)
	v8789 = *(*int32)(unsafe.Add(mBase, uint32(v8710)+4))
	if v8788 < v8789 {
		v8722 = v8785
		v8723 = v8723 + int32(1)
		v8725 = v8788
		goto L1229
	} else {
		goto L1240
	}
L1240:
	;
	goto L1230
L1241:
	;
	v8838 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+56)) = v8838
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+52)) = v8706
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+48)) = v8838
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+44)) = v8793
	*(*int32)(unsafe.Add(mBase, uint32(v8836))) = int32(332)
	v8846 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+4)) = v8846
	v8848 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8836)+8)) = v8848
	v8850 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8836)+16)) = v8850
	v8852 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8836)+24)) = v8852
	v8854 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8855 = *(*int32)(unsafe.Add(mBase, uint32(v8854)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8836)+32)) = v8855
	v8857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8836)+36)) = uint8(v8857)
	v8859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8836)+37)) = uint8(v8859)
	v12585 = v8836
	v12589 = v47
	goto L3
L1242:
	;
	v9790 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v9791 = m.ExcPending
	if v9791 != 0 {
		goto L1
	} else {
		goto L1360
	}
L1243:
	;
	v9635 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9636 = *(*int32)(unsafe.Add(mBase, uint32(v9635)+4))
	if v9636 == int32(0) {
		v9720 = v4
		goto L1343
	} else {
		goto L1344
	}
L1244:
	;
	v9245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v9245 == int32(0) {
		goto L1301
	} else {
		goto L1302
	}
L1245:
	;
	v8864 = F_use_physical_tlist(m, l0, l1, l2)
	mBase = m.M
	v8865 = m.ExcPending
	if v8865 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	v8866 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v8864 != 0 {
		goto L1250
	} else {
		goto L1251
	}
L1247:
	;
	v9243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9204)+37)) = uint8(v9243)
	v12585 = v9204
	v12589 = v47
	goto L3
L1248:
	;
	v9174 = F_palloc0(m, int32(80))
	mBase = m.M
	v9175 = m.ExcPending
	if v9175 != 0 {
		goto L1
	} else {
		goto L1300
	}
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9124)+44)) = v9122
	v9164 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9124)+8)) = v9164
	v9166 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9124)+16)) = v9166
	v9168 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9124)+24)) = v9168
	v9170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9171 = *(*int32)(unsafe.Add(mBase, uint32(v9170)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9124)+32)) = v9171
	v9204 = v9124
	goto L1247
L1250:
	;
	v8868 = F_create_plan_recurse(m, l0, v8866, int32(0))
	mBase = m.M
	v8869 = m.ExcPending
	if v8869 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1251:
	;
	goto L1252
L1252:
	;
	v8878 = int32(0)
	v8879 = *(*int32)(unsafe.Add(mBase, uint32(v8866)+4))
	switch v8879 - int32(332) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v8894 = v8878
		goto L1257
	case 2:
		goto L1259
	default:
		goto L1258
	case 23:
		goto L1260
	}
L1253:
	;
	v8870 = *(*int32)(unsafe.Add(mBase, uint32(v8868)+44))
	if l2&int32(4) == int32(0) {
		v9122 = v8870
		v9124 = v8868
		goto L1249
	} else {
		goto L1254
	}
L1254:
	;
	v8875 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v8870, v8875)
	mBase = m.M
	v8877 = m.ExcPending
	if v8877 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	v9122 = v8870
	v9124 = v8868
	goto L1249
L1256:
	;
	if v8896 != 0 {
		goto L1262
	} else {
		goto L1263
	}
L1257:
	;
	v8896 = v8894
	goto L1256
L1258:
	;
	v8894 = int32(1)
	goto L1257
L1259:
	;
	v8887 = *(*int32)(unsafe.Add(mBase, uint32(v8866)))
	if v8887 != int32(290) {
		v8894 = v8878
		goto L1257
	} else {
		goto L1261
	}
L1260:
	;
	v8882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8866)+72)))
	v8896 = int32(base.Ui32(v8882&int32(4)) >> (uint(int32(2)) % 32))
	goto L1256
L1261:
	;
	v8890 = *(*int32)(unsafe.Add(mBase, uint32(v8866)+72))
	v8896 = base.B2i32(v8890 == int32(0))
	goto L1256
L1262:
	;
	v8898 = F_create_plan_recurse(m, l0, v8866, int32(8))
	mBase = m.M
	v8899 = m.ExcPending
	if v8899 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1263:
	;
	goto L1264
L1264:
	;
	v8984 = int32(0)
	v8986 = F_create_plan_recurse(m, l0, v8866, v8984)
	mBase = m.M
	v8987 = m.ExcPending
	if v8987 != 0 {
		goto L1
	} else {
		goto L1282
	}
L1265:
	;
	v8900 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8901 = *(*int32)(unsafe.Add(mBase, uint32(v8900)+4))
	if v8901 == int32(0) {
		goto L1266
	} else {
		goto L1267
	}
L1266:
	;
	v9122 = int32(0)
	v9124 = v8898
	goto L1249
L1267:
	;
	goto L1268
L1268:
	;
	v8906 = int32(0)
	v8907 = *(*int32)(unsafe.Add(mBase, uint32(v8901)+4))
	if v8907 <= v8906 {
		v9122 = v8906
		v9124 = v8898
		goto L1249
	} else {
		goto L1269
	}
L1269:
	;
	v8910 = *(*int32)(unsafe.Add(mBase, uint32(v8900)+8))
	v8915 = int32(1)
	v8916 = v8906
	v8917 = v4
	goto L1270
L1270:
	;
	v8957 = *(*int32)(unsafe.Add(mBase, uint32(v8901)+12))
	v8961 = *(*int32)(unsafe.Add(mBase, uint32(v8957+v8916<<(uint(int32(2))%32))))
	v8962 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8962 != 0 {
		goto L1272
	} else {
		goto L1273
	}
L1271:
	;
	v9122 = v8978
	v9124 = v8898
	goto L1249
L1272:
	;
	v8963 = F_replace_nestloop_params_mutator(m, v8961, l0)
	mBase = m.M
	v8964 = m.ExcPending
	if v8964 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1273:
	;
	v8965 = v8961
	goto L1274
L1274:
	;
	v8967 = int32(0)
	v8969 = F_makeTargetEntry(m, v8965, base.I32_extend16_s(v8915), v8967, v8967)
	mBase = m.M
	v8970 = m.ExcPending
	if v8970 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1275:
	;
	v8965 = v8963
	goto L1274
L1276:
	;
	if v8910 != 0 {
		goto L1277
	} else {
		goto L1278
	}
L1277:
	;
	v8974 = *(*int32)(unsafe.Add(mBase, uint32(v8910-int32(4)+v8915<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8969)+16)) = v8974
	goto L1279
L1278:
	;
	goto L1279
L1279:
	;
	v8978 = F_lappend(m, v8917, v8969)
	mBase = m.M
	v8979 = m.ExcPending
	if v8979 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	v8981 = v8916 + int32(1)
	v8982 = *(*int32)(unsafe.Add(mBase, uint32(v8901)+4))
	if v8981 < v8982 {
		v8915 = v8915 + int32(1)
		v8916 = v8981
		v8917 = v8978
		goto L1270
	} else {
		goto L1281
	}
L1281:
	;
	goto L1271
L1282:
	;
	v8988 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8989 = *(*int32)(unsafe.Add(mBase, uint32(v8988)+4))
	if v8989 == int32(0) {
		v9073 = v8984
		goto L1283
	} else {
		goto L1284
	}
L1283:
	;
	v9114 = *(*int32)(unsafe.Add(mBase, uint32(v8986)+44))
	v9115 = F_tlist_same_exprs(m, v9073, v9114)
	mBase = m.M
	v9116 = m.ExcPending
	if v9116 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1284:
	;
	v8993 = *(*int32)(unsafe.Add(mBase, uint32(v8989)+4))
	if v8993 <= int32(0) {
		v9073 = v8984
		goto L1283
	} else {
		goto L1285
	}
L1285:
	;
	v8996 = *(*int32)(unsafe.Add(mBase, uint32(v8988)+8))
	v9001 = int32(1)
	v9002 = v8984
	v9004 = v4
	goto L1286
L1286:
	;
	v9043 = *(*int32)(unsafe.Add(mBase, uint32(v8989)+12))
	v9047 = *(*int32)(unsafe.Add(mBase, uint32(v9043+v9004<<(uint(int32(2))%32))))
	v9048 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9048 != 0 {
		goto L1288
	} else {
		goto L1289
	}
L1287:
	;
	v9073 = v9064
	goto L1283
L1288:
	;
	v9049 = F_replace_nestloop_params_mutator(m, v9047, l0)
	mBase = m.M
	v9050 = m.ExcPending
	if v9050 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1289:
	;
	v9051 = v9047
	goto L1290
L1290:
	;
	v9053 = int32(0)
	v9055 = F_makeTargetEntry(m, v9051, base.I32_extend16_s(v9001), v9053, v9053)
	mBase = m.M
	v9056 = m.ExcPending
	if v9056 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1291:
	;
	v9051 = v9049
	goto L1290
L1292:
	;
	if v8996 != 0 {
		goto L1293
	} else {
		goto L1294
	}
L1293:
	;
	v9060 = *(*int32)(unsafe.Add(mBase, uint32(v8996-int32(4)+v9001<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9055)+16)) = v9060
	goto L1295
L1294:
	;
	goto L1295
L1295:
	;
	v9064 = F_lappend(m, v9002, v9055)
	mBase = m.M
	v9065 = m.ExcPending
	if v9065 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	v9067 = v9004 + int32(1)
	v9068 = *(*int32)(unsafe.Add(mBase, uint32(v8989)+4))
	if v9067 < v9068 {
		v9001 = v9001 + int32(1)
		v9002 = v9064
		v9004 = v9067
		goto L1286
	} else {
		goto L1297
	}
L1297:
	;
	goto L1287
L1298:
	;
	if v9115 == int32(0) {
		goto L1248
	} else {
		goto L1299
	}
L1299:
	;
	v9122 = v9073
	v9124 = v8986
	goto L1249
L1300:
	;
	v9176 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+72)) = v9176
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+56)) = v9176
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+52)) = v8986
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+48)) = v9176
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+44)) = v9073
	*(*int32)(unsafe.Add(mBase, uint32(v9174))) = int32(331)
	v9186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+4)) = v9186
	v9188 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9174)+8)) = v9188
	v9190 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9174)+16)) = v9190
	v9192 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9174)+24)) = v9192
	v9194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9195 = *(*int32)(unsafe.Add(mBase, uint32(v9194)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9174)+32)) = v9195
	v9197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9174)+36)) = uint8(v9197)
	v9204 = v9174
	goto L1247
L1301:
	;
	v9478 = int32(0)
	v9479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9480 = *(*int32)(unsafe.Add(mBase, uint32(v9479)+4))
	if v9480 == v9478 {
		v9566 = v9478
		goto L1327
	} else {
		goto L1328
	}
L1302:
	;
	v9248 = *(*int32)(unsafe.Add(mBase, uint32(v9245)+4))
	if v9248 <= int32(0) {
		goto L1301
	} else {
		goto L1303
	}
L1303:
	;
	v9260 = v4
	goto L1304
L1304:
	;
	v9295 = *(*int32)(unsafe.Add(mBase, uint32(v9245)+12))
	v9299 = *(*int32)(unsafe.Add(mBase, uint32(v9295+v9260<<(uint(int32(2))%32))))
	v9300 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+16))
	v9301 = *(*int32)(unsafe.Add(mBase, uint32(v9300)+4))
	v9302 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+20))
	v9303 = F_create_plan(m, v9300, v9302)
	mBase = m.M
	v9304 = m.ExcPending
	if v9304 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1305:
	;
	goto L1301
L1306:
	;
	v9305 = *(*int64)(unsafe.Add(mBase, uint32(v9301)+128))
	v9306 = *(*int32)(unsafe.Add(mBase, uint32(v9301)+136))
	v9308 = F_palloc0(m, int32(104))
	mBase = m.M
	v9309 = m.ExcPending
	if v9309 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9308))) = int32(373)
	v9312 = *(*int32)(unsafe.Add(mBase, uint32(v9303)+44))
	v9313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9308)+84)) = v9313
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+80)) = v9306
	*(*int64)(unsafe.Add(mBase, uint32(v9308)+72)) = v9305
	v9317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+56)) = v9317
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+52)) = v9303
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+48)) = v9317
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+44)) = v9312
	*(*int64)(unsafe.Add(mBase, uint32(v9308)+92)) = v9313
	v9325 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+20))
	v9326 = *(*int32)(unsafe.Add(mBase, uint32(v9325)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+4)) = v9326
	v9328 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+20))
	v9329 = *(*float64)(unsafe.Add(mBase, uint32(v9328)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9308)+8)) = v9329
	v9331 = *(*float64)(unsafe.Add(mBase, uint32(v9299)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9308)+24)) = int64(4607182418800017408)
	*(*float64)(unsafe.Add(mBase, uint32(v9308)+16)) = v9331
	v9335 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+20))
	v9336 = *(*int32)(unsafe.Add(mBase, uint32(v9335)+12))
	v9337 = *(*int32)(unsafe.Add(mBase, uint32(v9336)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v9308)+36)) = uint8(v9317)
	*(*int32)(unsafe.Add(mBase, uint32(v9308)+32)) = v9337
	v9341 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+20))
	v9342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9341)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9308)+37)) = uint8(v9342)
	v9344 = *(*int32)(unsafe.Add(mBase, uint32(v9299)+32))
	v9346 = m.G0
	v9348 = v9346 - int32(32)
	m.G0 = v9348
	v9350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9351 = *(*int32)(unsafe.Add(mBase, uint32(v9350)+8))
	v9352 = F_lappend(m, v9351, v9308)
	mBase = m.M
	v9353 = m.ExcPending
	if v9353 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1308:
	;
	v9354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9354)+8)) = v9352
	v9356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9357 = *(*int32)(unsafe.Add(mBase, uint32(v9356)+12))
	v9359 = F_lappend(m, v9357, int32(0))
	mBase = m.M
	v9360 = m.ExcPending
	if v9360 != 0 {
		goto L1
	} else {
		goto L1309
	}
L1309:
	;
	v9361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9361)+12)) = v9359
	v9363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9364 = *(*int32)(unsafe.Add(mBase, uint32(v9363)+16))
	v9365 = F_lappend(m, v9364, v9300)
	mBase = m.M
	v9366 = m.ExcPending
	if v9366 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1310:
	;
	v9367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9367)+16)) = v9365
	v9370 = F_palloc0(m, int32(72))
	mBase = m.M
	v9371 = m.ExcPending
	if v9371 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1311:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9370))) = int64(17179869207)
	v9374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9375 = *(*int32)(unsafe.Add(mBase, uint32(v9374)+8))
	if v9375 != 0 {
		goto L1312
	} else {
		goto L1313
	}
L1312:
	;
	v9376 = *(*int32)(unsafe.Add(mBase, uint32(v9375)+4))
	v9377 = v9376
	goto L1314
L1313:
	;
	v9377 = v9317
	goto L1314
L1314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+16)) = v9377
	*(*int32)(unsafe.Add(mBase, uint32(v9348)+16)) = v9377
	v9383 = F_psprintf(m, int32(474979), v9348+int32(16))
	mBase = m.M
	v9384 = m.ExcPending
	if v9384 != 0 {
		goto L1
	} else {
		goto L1315
	}
L1315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+20)) = v9383
	v9386 = *(*int32)(unsafe.Add(mBase, uint32(v9308)+44))
	if v9386 == int32(0) {
		goto L1317
	} else {
		goto L1318
	}
L1316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+32)) = v9408
	v9410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9308)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9370)+38)) = uint8(v9410)
	v9412 = *(*int32)(unsafe.Add(mBase, uint32(v9344)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9348)+12)) = v9412
	*(*int32)(unsafe.Add(mBase, uint32(v9348)+28)) = v9412
	v9418 = F_list_make1_impl(m, int32(471), v9348+int32(12))
	mBase = m.M
	v9419 = m.ExcPending
	if v9419 != 0 {
		goto L1
	} else {
		goto L1323
	}
L1317:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9370)+24)) = int64(-4294965018)
	v9408 = int32(0)
	goto L1316
L1318:
	;
	v9389 = *(*int32)(unsafe.Add(mBase, uint32(v9386)+12))
	v9390 = *(*int32)(unsafe.Add(mBase, uint32(v9389)))
	v9391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9390)+26)))
	if v9391 != 0 {
		goto L1317
	} else {
		goto L1319
	}
L1319:
	;
	v9392 = *(*int32)(unsafe.Add(mBase, uint32(v9390)+4))
	v9393 = F_exprType(m, v9392)
	mBase = m.M
	v9394 = m.ExcPending
	if v9394 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+24)) = v9393
	v9396 = *(*int32)(unsafe.Add(mBase, uint32(v9390)+4))
	v9397 = F_exprTypmod(m, v9396)
	mBase = m.M
	v9398 = m.ExcPending
	if v9398 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+28)) = v9397
	v9400 = *(*int32)(unsafe.Add(mBase, uint32(v9390)+4))
	v9401 = F_exprCollation(m, v9400)
	mBase = m.M
	v9402 = m.ExcPending
	if v9402 != 0 {
		goto L1
	} else {
		goto L1322
	}
L1322:
	;
	v9408 = v9401
	goto L1316
L1323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+40)) = v9418
	v9421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9422 = F_lappend(m, v9421, v9370)
	mBase = m.M
	v9423 = m.ExcPending
	if v9423 != 0 {
		goto L1
	} else {
		goto L1324
	}
L1324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v9422
	F_cost_subplan(m, v9370, v9308)
	mBase = m.M
	v9426 = m.ExcPending
	if v9426 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	m.G0 = v9348 + int32(32)
	v9431 = v9260 + int32(1)
	v9432 = *(*int32)(unsafe.Add(mBase, uint32(v9245)+4))
	if v9431 < v9432 {
		v9260 = v9431
		goto L1304
	} else {
		goto L1326
	}
L1326:
	;
	goto L1305
L1327:
	;
	v9606 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v9608 = F_palloc0(m, int32(80))
	mBase = m.M
	v9609 = m.ExcPending
	if v9609 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1328:
	;
	v9484 = int32(0)
	v9485 = *(*int32)(unsafe.Add(mBase, uint32(v9480)+4))
	if v9485 <= v9484 {
		v9566 = v9478
		goto L1327
	} else {
		goto L1329
	}
L1329:
	;
	v9488 = *(*int32)(unsafe.Add(mBase, uint32(v9479)+8))
	v9493 = int32(1)
	v9495 = v9478
	v9500 = v9484
	goto L1330
L1330:
	;
	v9535 = *(*int32)(unsafe.Add(mBase, uint32(v9480)+12))
	v9539 = *(*int32)(unsafe.Add(mBase, uint32(v9535+v9500<<(uint(int32(2))%32))))
	v9540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9540 != 0 {
		goto L1332
	} else {
		goto L1333
	}
L1331:
	;
	v9566 = v9556
	goto L1327
L1332:
	;
	v9541 = F_replace_nestloop_params_mutator(m, v9539, l0)
	mBase = m.M
	v9542 = m.ExcPending
	if v9542 != 0 {
		goto L1
	} else {
		goto L1335
	}
L1333:
	;
	v9543 = v9539
	goto L1334
L1334:
	;
	v9545 = int32(0)
	v9547 = F_makeTargetEntry(m, v9543, base.I32_extend16_s(v9493), v9545, v9545)
	mBase = m.M
	v9548 = m.ExcPending
	if v9548 != 0 {
		goto L1
	} else {
		goto L1336
	}
L1335:
	;
	v9543 = v9541
	goto L1334
L1336:
	;
	if v9488 != 0 {
		goto L1337
	} else {
		goto L1338
	}
L1337:
	;
	v9552 = *(*int32)(unsafe.Add(mBase, uint32(v9488-int32(4)+v9493<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+16)) = v9552
	goto L1339
L1338:
	;
	goto L1339
L1339:
	;
	v9556 = F_lappend(m, v9495, v9547)
	mBase = m.M
	v9557 = m.ExcPending
	if v9557 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1340:
	;
	v9559 = v9500 + int32(1)
	v9560 = *(*int32)(unsafe.Add(mBase, uint32(v9480)+4))
	if v9559 < v9560 {
		v9493 = v9493 + int32(1)
		v9495 = v9556
		v9500 = v9559
		goto L1330
	} else {
		goto L1341
	}
L1341:
	;
	goto L1331
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9608)+72)) = v9606
	*(*int32)(unsafe.Add(mBase, uint32(v9608)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9608)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9608)+44)) = v9566
	*(*int32)(unsafe.Add(mBase, uint32(v9608))) = int32(331)
	v9618 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9608)+4)) = v9618
	v9620 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9608)+8)) = v9620
	v9622 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9608)+16)) = v9622
	v9624 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9608)+24)) = v9624
	v9626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9627 = *(*int32)(unsafe.Add(mBase, uint32(v9626)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9608)+32)) = v9627
	v9629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9608)+36)) = uint8(v9629)
	v9631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9608)+37)) = uint8(v9631)
	v9633 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v9633
	v12585 = v9608
	v12589 = v47
	goto L3
L1343:
	;
	v9761 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v9762 = F_order_qual_clauses(m, l0, v9761)
	mBase = m.M
	v9763 = m.ExcPending
	if v9763 != 0 {
		goto L1
	} else {
		goto L1358
	}
L1344:
	;
	v9640 = *(*int32)(unsafe.Add(mBase, uint32(v9636)+4))
	if v9640 <= int32(0) {
		v9720 = v4
		goto L1343
	} else {
		goto L1345
	}
L1345:
	;
	v9643 = *(*int32)(unsafe.Add(mBase, uint32(v9635)+8))
	v9648 = int32(1)
	v9649 = v4
	v9650 = v4
	goto L1346
L1346:
	;
	v9690 = *(*int32)(unsafe.Add(mBase, uint32(v9636)+12))
	v9694 = *(*int32)(unsafe.Add(mBase, uint32(v9690+v9650<<(uint(int32(2))%32))))
	v9695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9695 != 0 {
		goto L1348
	} else {
		goto L1349
	}
L1347:
	;
	v9720 = v9711
	goto L1343
L1348:
	;
	v9696 = F_replace_nestloop_params_mutator(m, v9694, l0)
	mBase = m.M
	v9697 = m.ExcPending
	if v9697 != 0 {
		goto L1
	} else {
		goto L1351
	}
L1349:
	;
	v9698 = v9694
	goto L1350
L1350:
	;
	v9700 = int32(0)
	v9702 = F_makeTargetEntry(m, v9698, base.I32_extend16_s(v9648), v9700, v9700)
	mBase = m.M
	v9703 = m.ExcPending
	if v9703 != 0 {
		goto L1
	} else {
		goto L1352
	}
L1351:
	;
	v9698 = v9696
	goto L1350
L1352:
	;
	if v9643 != 0 {
		goto L1353
	} else {
		goto L1354
	}
L1353:
	;
	v9707 = *(*int32)(unsafe.Add(mBase, uint32(v9643-int32(4)+v9648<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9702)+16)) = v9707
	goto L1355
L1354:
	;
	goto L1355
L1355:
	;
	v9711 = F_lappend(m, v9649, v9702)
	mBase = m.M
	v9712 = m.ExcPending
	if v9712 != 0 {
		goto L1
	} else {
		goto L1356
	}
L1356:
	;
	v9714 = v9650 + int32(1)
	v9715 = *(*int32)(unsafe.Add(mBase, uint32(v9636)+4))
	if v9714 < v9715 {
		v9648 = v9648 + int32(1)
		v9649 = v9711
		v9650 = v9714
		goto L1346
	} else {
		goto L1357
	}
L1357:
	;
	goto L1347
L1358:
	;
	v9765 = F_palloc0(m, int32(80))
	mBase = m.M
	v9766 = m.ExcPending
	if v9766 != 0 {
		goto L1
	} else {
		goto L1359
	}
L1359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9765)+72)) = v9762
	*(*int32)(unsafe.Add(mBase, uint32(v9765)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9765)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9765)+44)) = v9720
	*(*int32)(unsafe.Add(mBase, uint32(v9765))) = int32(331)
	v9775 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9765)+4)) = v9775
	v9777 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9765)+8)) = v9777
	v9779 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9765)+16)) = v9779
	v9781 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9765)+24)) = v9781
	v9783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9784 = *(*int32)(unsafe.Add(mBase, uint32(v9783)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9765)+32)) = v9784
	v9786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9765)+36)) = uint8(v9786)
	v9788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9765)+37)) = uint8(v9788)
	v12585 = v9765
	v12589 = v47
	goto L3
L1360:
	;
	v12585 = v9790
	v12589 = v47
	goto L3
L1361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9793))) = int32(335)
	v9797 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9798 = *(*int32)(unsafe.Add(mBase, uint32(v9797)+4))
	if v9798 == int32(0) {
		goto L1363
	} else {
		goto L1364
	}
L1362:
	;
	v9972 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v9973 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9974 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+4)) = v9974
	v9976 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9793)+8)) = v9976
	v9978 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9793)+16)) = v9978
	v9980 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9793)+24)) = v9980
	v9982 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9983 = *(*int32)(unsafe.Add(mBase, uint32(v9982)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+32)) = v9983
	v9985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9793)+36)) = uint8(v9985)
	v9987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v9988 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+56)) = v9988
	*(*int64)(unsafe.Add(mBase, uint32(v9793)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+44)) = v9936
	*(*uint8)(unsafe.Add(mBase, uint32(v9793)+37)) = uint8(v9987)
	v9994 = *(*int32)(unsafe.Add(mBase, uint32(v9973)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+72)) = v9994
	v9996 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9997 = *(*int32)(unsafe.Add(mBase, uint32(v9996)+8))
	v10003 = v9793 + int32(84)
	v10010 = F_prepare_sort_from_pathkeys(m, v9793, v9972, v9997, v9988, int32(1), v9793+int32(80), v10003, v9793+int32(88), v9793+int32(92), v9793+int32(96))
	mBase = m.M
	v10011 = m.ExcPending
	if v10011 != 0 {
		goto L1
	} else {
		goto L1379
	}
L1363:
	;
	v9926 = int32(0)
	v9934 = v9926
	v9936 = v9926
	goto L1362
L1364:
	;
	v9802 = *(*int32)(unsafe.Add(mBase, uint32(v9798)+4))
	if v9802 <= int32(0) {
		v9934 = v4
		v9936 = v4
		goto L1362
	} else {
		goto L1365
	}
L1365:
	;
	v9805 = *(*int32)(unsafe.Add(mBase, uint32(v9797)+8))
	v9812 = int32(1)
	v9814 = v4
	v9816 = v4
	goto L1366
L1366:
	;
	v9852 = *(*int32)(unsafe.Add(mBase, uint32(v9798)+12))
	v9856 = *(*int32)(unsafe.Add(mBase, uint32(v9852+v9814<<(uint(int32(2))%32))))
	v9857 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9857 != 0 {
		goto L1368
	} else {
		goto L1369
	}
L1367:
	;
	if v9873 == int32(0) {
		goto L1363
	} else {
		goto L1378
	}
L1368:
	;
	v9858 = F_replace_nestloop_params_mutator(m, v9856, l0)
	mBase = m.M
	v9859 = m.ExcPending
	if v9859 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1369:
	;
	v9860 = v9856
	goto L1370
L1370:
	;
	v9862 = int32(0)
	v9864 = F_makeTargetEntry(m, v9860, base.I32_extend16_s(v9812), v9862, v9862)
	mBase = m.M
	v9865 = m.ExcPending
	if v9865 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1371:
	;
	v9860 = v9858
	goto L1370
L1372:
	;
	if v9805 != 0 {
		goto L1373
	} else {
		goto L1374
	}
L1373:
	;
	v9869 = *(*int32)(unsafe.Add(mBase, uint32(v9805-int32(4)+v9812<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9864)+16)) = v9869
	goto L1375
L1374:
	;
	goto L1375
L1375:
	;
	v9873 = F_lappend(m, v9816, v9864)
	mBase = m.M
	v9874 = m.ExcPending
	if v9874 != 0 {
		goto L1
	} else {
		goto L1376
	}
L1376:
	;
	v9876 = v9814 + int32(1)
	v9877 = *(*int32)(unsafe.Add(mBase, uint32(v9798)+4))
	if v9876 < v9877 {
		v9812 = v9812 + int32(1)
		v9814 = v9876
		v9816 = v9873
		goto L1366
	} else {
		goto L1377
	}
L1377:
	;
	goto L1367
L1378:
	;
	v9881 = *(*int32)(unsafe.Add(mBase, uint32(v9873)+4))
	v9934 = v9881
	v9936 = v9873
	goto L1362
L1379:
	;
	v10012 = *(*int32)(unsafe.Add(mBase, uint32(v9793)+44))
	if v10012 != 0 {
		goto L1380
	} else {
		goto L1381
	}
L1380:
	;
	v10013 = *(*int32)(unsafe.Add(mBase, uint32(v10012)+4))
	v10014 = v10013
	goto L1382
L1381:
	;
	v10014 = v4
	goto L1382
L1382:
	;
	v10015 = int32(0)
	v10016 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10016 == v10015 {
		v10321 = v10015
		goto L1383
	} else {
		goto L1384
	}
L1383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+100)) = int32(-1)
	v10357 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
	if v10357 == int32(0) {
		goto L1439
	} else {
		goto L1440
	}
L1384:
	;
	v10019 = int32(0)
	v10020 = *(*int32)(unsafe.Add(mBase, uint32(v10016)+4))
	if v10020 <= v10019 {
		v10321 = v10015
		goto L1383
	} else {
		goto L1385
	}
L1385:
	;
	v10031 = v10019
	v10034 = v10015
	goto L1386
L1386:
	;
	v10067 = *(*int32)(unsafe.Add(mBase, uint32(v10016)+12))
	v10071 = *(*int32)(unsafe.Add(mBase, uint32(v10067+v10031<<(uint(int32(2))%32))))
	v10073 = F_create_plan_recurse(m, l0, v10071, int32(1))
	mBase = m.M
	v10074 = m.ExcPending
	if v10074 != 0 {
		goto L1
	} else {
		goto L1389
	}
L1387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10300 = m.ExcPending
	if v10300 != 0 {
		goto L1
	} else {
		goto L1436
	}
L1388:
	;
	goto L1387
L1389:
	;
	v10075 = *(*int32)(unsafe.Add(mBase, uint32(v10071)+8))
	v10076 = *(*int32)(unsafe.Add(mBase, uint32(v10075)+8))
	v10077 = *(*int32)(unsafe.Add(mBase, uint32(v10003)))
	v10089 = F_prepare_sort_from_pathkeys(m, v10073, v9972, v10076, v10077, int32(0), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128), v47+int32(52))
	mBase = m.M
	v10090 = m.ExcPending
	if v10090 != 0 {
		goto L1
	} else {
		goto L1390
	}
L1390:
	;
	v10091 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v10092 = *(*int32)(unsafe.Add(mBase, uint32(v10003)))
	v10093 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v10095 = v10093 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v10095) {
		goto L1394
	} else {
		goto L1395
	}
L1391:
	;
	if v10157 != 0 {
		goto L1388
	} else {
		goto L1409
	}
L1392:
	;
	v10157 = int32(0)
	goto L1391
L1393:
	;
	v10131 = v10126
	v10132 = v10127
	v10133 = v10128
	goto L1403
L1394:
	;
	if (v10091|v10092)&int32(3) != 0 {
		v10126 = v10091
		v10127 = v10092
		v10128 = v10095
		goto L1393
	} else {
		goto L1397
	}
L1395:
	;
	v10119 = v10091
	v10120 = v10092
	v10121 = v10095
	goto L1396
L1396:
	;
	if v10121 == int32(0) {
		goto L1392
	} else {
		goto L1402
	}
L1397:
	;
	v10103 = v10091
	v10104 = v10092
	v10105 = v10095
	goto L1398
L1398:
	;
	v10108 = *(*int32)(unsafe.Add(mBase, uint32(v10103)))
	v10109 = *(*int32)(unsafe.Add(mBase, uint32(v10104)))
	if v10108 != v10109 {
		v10126 = v10103
		v10127 = v10104
		v10128 = v10105
		goto L1393
	} else {
		goto L1400
	}
L1399:
	;
	v10119 = v10114
	v10120 = v10112
	v10121 = v10116
	goto L1396
L1400:
	;
	v10111 = int32(4)
	v10112 = v10104 + v10111
	v10114 = v10103 + v10111
	v10116 = v10105 - v10111
	if base.Ui32(int32(3)) < base.Ui32(v10116) {
		v10103 = v10114
		v10104 = v10112
		v10105 = v10116
		goto L1398
	} else {
		goto L1401
	}
L1401:
	;
	goto L1399
L1402:
	;
	v10126 = v10119
	v10127 = v10120
	v10128 = v10121
	goto L1393
L1403:
	;
	v10136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10131))))
	v10137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10132))))
	if v10136 == v10137 {
		goto L1405
	} else {
		goto L1406
	}
L1404:
	;
	v10157 = v10136 - v10137
	goto L1391
L1405:
	;
	v10139 = int32(1)
	v10144 = v10133 - v10139
	if v10144 != 0 {
		v10131 = v10131 + v10139
		v10132 = v10132 + v10139
		v10133 = v10144
		goto L1403
	} else {
		goto L1408
	}
L1406:
	;
	goto L1407
L1407:
	;
	goto L1404
L1408:
	;
	goto L1392
L1409:
	;
	v10158 = *(*int32)(unsafe.Add(mBase, uint32(v10071)+64))
	if v9972 == v10158 {
		goto L1412
	} else {
		goto L1413
	}
L1410:
	;
	v10291 = F_lappend(m, v10034, v10284)
	mBase = m.M
	v10292 = m.ExcPending
	if v10292 != 0 {
		goto L1
	} else {
		goto L1434
	}
L1411:
	;
	if v10211 != 0 {
		goto L1429
	} else {
		goto L1430
	}
L1412:
	;
	v10211 = int32(1)
	goto L1411
L1413:
	;
	goto L1414
L1414:
	;
	v10167 = int32(0)
	goto L1416
L1415:
	;
	v10211 = v10203
	goto L1411
L1416:
	;
	v10171 = int32(0)
	if v9972 == v10171 {
		v10181 = v10171
		goto L1418
	} else {
		goto L1419
	}
L1417:
	;
	v10203 = int32(0)
	goto L1415
L1418:
	;
	if v10158 != 0 {
		goto L1422
	} else {
		goto L1423
	}
L1419:
	;
	v10175 = *(*int32)(unsafe.Add(mBase, uint32(v9972)+4))
	if v10175 <= v10167 {
		v10181 = int32(0)
		goto L1418
	} else {
		goto L1420
	}
L1420:
	;
	v10177 = *(*int32)(unsafe.Add(mBase, uint32(v9972)+12))
	v10181 = v10177 + v10167<<(uint(int32(2))%32)
	goto L1418
L1421:
	;
	v10187 = base.B2i32(v10181 == int32(0))
	if v10181 == int32(0) {
		v10203 = v10187
		goto L1415
	} else {
		goto L1426
	}
L1422:
	;
	v10182 = *(*int32)(unsafe.Add(mBase, uint32(v10158)+4))
	if v10167 < v10182 {
		goto L1421
	} else {
		goto L1425
	}
L1423:
	;
	goto L1424
L1424:
	;
	v10211 = base.B2i32(v10181 == int32(0))
	goto L1411
L1425:
	;
	goto L1424
L1426:
	;
	v10190 = *(*int32)(unsafe.Add(mBase, uint32(v10158)+12))
	v10193 = v10190 + v10167<<(uint(int32(2))%32)
	if v10193 == int32(0) {
		v10203 = v10187
		goto L1415
	} else {
		goto L1427
	}
L1427:
	;
	v10198 = *(*int32)(unsafe.Add(mBase, uint32(v10181)))
	v10199 = *(*int32)(unsafe.Add(mBase, uint32(v10193)))
	if v10198 == v10199 {
		v10167 = v10167 + int32(1)
		goto L1416
	} else {
		goto L1428
	}
L1428:
	;
	goto L1417
L1429:
	;
	v10284 = v10089
	goto L1410
L1430:
	;
	goto L1431
L1431:
	;
	v10212 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v10213 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v10214 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v10216 = F_palloc0(m, int32(96))
	mBase = m.M
	v10217 = m.ExcPending
	if v10217 != 0 {
		goto L1
	} else {
		goto L1432
	}
L1432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10216))) = int32(362)
	v10220 = *(*int32)(unsafe.Add(mBase, uint32(v10089)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+44)) = v10220
	v10222 = *(*int32)(unsafe.Add(mBase, uint32(v10089)+4))
	v10223 = int32(4126313)
	v10224 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+88)) = v10214
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+84)) = v10213
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+80)) = v10212
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+76)) = v10091
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+72)) = v10093
	v10230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+56)) = v10230
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+52)) = v10089
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+48)) = v10230
	v10235 = int32(1)
	v10237 = v10222 + (v10224 ^ v10235)
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+4)) = v10237
	v10240 = v47 + int32(56)
	v10241 = *(*float64)(unsafe.Add(mBase, uint32(v10089)+16))
	v10242 = *(*float64)(unsafe.Add(mBase, uint32(v10089)+24))
	v10243 = *(*int32)(unsafe.Add(mBase, uint32(v10089)+32))
	v10246 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v10247 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v10249 = m.G0
	v10250 = int32(16)
	v10251 = v10249 - v10250
	m.G0 = v10251
	F_cost_tuplesort(m, v10251+int32(8), v10251, v10242, v10243, float64(0), v10246, v10247)
	mBase = m.M
	v10256 = *(*float64)(unsafe.Add(mBase, uint32(v10251)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v10240)+32)) = v10242
	v10259 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	v10260 = base.F64_add(v10241, v10256)
	*(*float64)(unsafe.Add(mBase, uint32(v10240)+48)) = v10260
	*(*int32)(unsafe.Add(mBase, uint32(v10240)+40)) = v10237 + (v10259 ^ v10235)
	v10266 = *(*float64)(unsafe.Add(mBase, uint32(v10251)))
	*(*float64)(unsafe.Add(mBase, uint32(v10240)+56)) = base.F64_add(v10260, v10266)
	m.G0 = v10251 + v10250
	goto L1433
L1433:
	;
	v10272 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v10216)+8)) = v10272
	v10274 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v10216)+16)) = v10274
	v10276 = *(*float64)(unsafe.Add(mBase, uint32(v10089)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10216)+24)) = v10276
	v10278 = *(*int32)(unsafe.Add(mBase, uint32(v10089)+32))
	v10279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10216)+36)) = uint8(v10279)
	*(*int32)(unsafe.Add(mBase, uint32(v10216)+32)) = v10278
	v10282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10089)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10216)+37)) = uint8(v10282)
	v10284 = v10216
	goto L1410
L1434:
	;
	v10294 = v10031 + int32(1)
	v10295 = *(*int32)(unsafe.Add(mBase, uint32(v10016)+4))
	if v10294 < v10295 {
		v10031 = v10294
		v10034 = v10291
		goto L1386
	} else {
		goto L1435
	}
L1435:
	;
	v10321 = v10291
	goto L1383
L1436:
	;
	F_errmsg_internal(m, int32(428052), int32(0))
	mBase = m.M
	v10304 = m.ExcPending
	if v10304 != 0 {
		goto L1
	} else {
		goto L1437
	}
L1437:
	;
	F_errfinish(m, int32(498183), int32(1519), int32(284016))
	mBase = m.M
	v10309 = m.ExcPending
	if v10309 != 0 {
		goto L1
	} else {
		goto L1438
	}
L1438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+76)) = v10321
	if l2&int32(3) == int32(0) {
		v12585 = v9793
		v12589 = v47
		goto L3
	} else {
		goto L1444
	}
L1440:
	;
	v10360 = *(*int32)(unsafe.Add(mBase, uint32(v9973)+184))
	v10362 = F_extract_actual_clauses(m, v10360, int32(0))
	mBase = m.M
	v10363 = m.ExcPending
	if v10363 != 0 {
		goto L1
	} else {
		goto L1441
	}
L1441:
	;
	if v10362 == int32(0) {
		goto L1439
	} else {
		goto L1442
	}
L1442:
	;
	v10366 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v10367 = F_make_partition_pruneinfo(m, l0, v9973, v10366, v10362)
	mBase = m.M
	v10368 = m.ExcPending
	if v10368 != 0 {
		goto L1
	} else {
		goto L1443
	}
L1443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9793)+100)) = v10367
	goto L1439
L1444:
	;
	if v9934 == v10014 {
		v12585 = v9793
		v12589 = v47
		goto L3
	} else {
		goto L1445
	}
L1445:
	;
	v10377 = *(*int32)(unsafe.Add(mBase, uint32(v9793)+44))
	v10378 = F_list_copy_head(m, v10377, v9934)
	mBase = m.M
	v10379 = m.ExcPending
	if v10379 != 0 {
		goto L1
	} else {
		goto L1446
	}
L1446:
	;
	v10380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9793)+37)))
	v10381 = F_inject_projection_plan(m, v9793, v10378, v10380)
	mBase = m.M
	v10382 = m.ExcPending
	if v10382 != 0 {
		goto L1
	} else {
		goto L1447
	}
L1447:
	;
	v12585 = v10381
	v12589 = v47
	goto L3
L1448:
	;
	v10558 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10559 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v10560 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v10560
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v10560
	*(*int32)(unsafe.Add(mBase, uint32(v47)+132)) = v10560
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v10560
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v10560
	v10570 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10570 == v10560 {
		goto L1465
	} else {
		goto L1466
	}
L1449:
	;
	v10512 = int32(0)
	v10519 = v10512
	v10520 = v10512
	goto L1448
L1450:
	;
	v10388 = *(*int32)(unsafe.Add(mBase, uint32(v10384)+4))
	if v10388 <= int32(0) {
		v10519 = v4
		v10520 = v4
		goto L1448
	} else {
		goto L1451
	}
L1451:
	;
	v10391 = *(*int32)(unsafe.Add(mBase, uint32(v10383)+8))
	v10397 = int32(1)
	v10399 = v4
	v10400 = v4
	goto L1452
L1452:
	;
	v10438 = *(*int32)(unsafe.Add(mBase, uint32(v10384)+12))
	v10442 = *(*int32)(unsafe.Add(mBase, uint32(v10438+v10399<<(uint(int32(2))%32))))
	v10443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10443 != 0 {
		goto L1454
	} else {
		goto L1455
	}
L1453:
	;
	if v10459 == int32(0) {
		goto L1449
	} else {
		goto L1464
	}
L1454:
	;
	v10444 = F_replace_nestloop_params_mutator(m, v10442, l0)
	mBase = m.M
	v10445 = m.ExcPending
	if v10445 != 0 {
		goto L1
	} else {
		goto L1457
	}
L1455:
	;
	v10446 = v10442
	goto L1456
L1456:
	;
	v10448 = int32(0)
	v10450 = F_makeTargetEntry(m, v10446, base.I32_extend16_s(v10397), v10448, v10448)
	mBase = m.M
	v10451 = m.ExcPending
	if v10451 != 0 {
		goto L1
	} else {
		goto L1458
	}
L1457:
	;
	v10446 = v10444
	goto L1456
L1458:
	;
	if v10391 != 0 {
		goto L1459
	} else {
		goto L1460
	}
L1459:
	;
	v10455 = *(*int32)(unsafe.Add(mBase, uint32(v10391-int32(4)+v10397<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v10450)+16)) = v10455
	goto L1461
L1460:
	;
	goto L1461
L1461:
	;
	v10459 = F_lappend(m, v10400, v10450)
	mBase = m.M
	v10460 = m.ExcPending
	if v10460 != 0 {
		goto L1
	} else {
		goto L1462
	}
L1462:
	;
	v10462 = v10399 + int32(1)
	v10463 = *(*int32)(unsafe.Add(mBase, uint32(v10384)+4))
	if v10462 < v10463 {
		v10397 = v10397 + int32(1)
		v10399 = v10462
		v10400 = v10459
		goto L1452
	} else {
		goto L1463
	}
L1463:
	;
	goto L1453
L1464:
	;
	v10467 = *(*int32)(unsafe.Add(mBase, uint32(v10459)+4))
	v10519 = v10467
	v10520 = v10459
	goto L1448
L1465:
	;
	v10573 = int32(0)
	v10575 = F_makeBoolConst(m, v10573, v10573)
	mBase = m.M
	v10576 = m.ExcPending
	if v10576 != 0 {
		goto L1
	} else {
		goto L1468
	}
L1466:
	;
	goto L1467
L1467:
	;
	v10611 = F_palloc0(m, int32(96))
	mBase = m.M
	v10612 = m.ExcPending
	if v10612 != 0 {
		goto L1
	} else {
		goto L1471
	}
L1468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v10575
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v10575
	v10582 = F_list_make1_impl(m, int32(1), v47+int32(28))
	mBase = m.M
	v10583 = m.ExcPending
	if v10583 != 0 {
		goto L1
	} else {
		goto L1469
	}
L1469:
	;
	v10585 = F_palloc0(m, int32(80))
	mBase = m.M
	v10586 = m.ExcPending
	if v10586 != 0 {
		goto L1
	} else {
		goto L1470
	}
L1470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10585)+72)) = v10582
	*(*int32)(unsafe.Add(mBase, uint32(v10585)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10585)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10585)+44)) = v10520
	*(*int32)(unsafe.Add(mBase, uint32(v10585))) = int32(331)
	v10595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10585)+4)) = v10595
	v10597 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v10585)+8)) = v10597
	v10599 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v10585)+16)) = v10599
	v10601 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v10585)+24)) = v10601
	v10603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10604 = *(*int32)(unsafe.Add(mBase, uint32(v10603)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10585)+32)) = v10604
	v10606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10585)+36)) = uint8(v10606)
	v10608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10585)+37)) = uint8(v10608)
	v12585 = v10585
	v12589 = v47
	goto L3
L1471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10611)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+44)) = v10520
	*(*int32)(unsafe.Add(mBase, uint32(v10611))) = int32(334)
	v10620 = *(*int32)(unsafe.Add(mBase, uint32(v10558)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+72)) = v10620
	if v10559 != 0 {
		goto L1474
	} else {
		goto L1475
	}
L1472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+88)) = int32(-1)
	v11007 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
	if v11007 == int32(0) {
		goto L1544
	} else {
		goto L1545
	}
L1473:
	;
	v10659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10659 == int32(0) {
		v10976 = v4
		v10977 = v4
		v10987 = v10658
		goto L1472
	} else {
		goto L1484
	}
L1474:
	;
	v10622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10623 = *(*int32)(unsafe.Add(mBase, uint32(v10622)+8))
	v10636 = F_prepare_sort_from_pathkeys(m, v10611, v10559, v10623, int32(0), int32(1), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128), v47+int32(52))
	mBase = m.M
	v10637 = m.ExcPending
	if v10637 != 0 {
		goto L1
	} else {
		goto L1477
	}
L1475:
	;
	goto L1476
L1476:
	;
	v10643 = int32(1)
	v10645 = int32(*(*uint8)(unsafe.Add(mBase, _consts[494])))
	if v10645 != v10643 {
		v10657 = v4
		v10658 = v10643
		goto L1473
	} else {
		goto L1481
	}
L1477:
	;
	v10638 = *(*int32)(unsafe.Add(mBase, uint32(v10611)+44))
	if v10638 != 0 {
		goto L1478
	} else {
		goto L1479
	}
L1478:
	;
	v10639 = *(*int32)(unsafe.Add(mBase, uint32(v10638)+4))
	v10641 = v10639
	goto L1480
L1479:
	;
	v10641 = int32(0)
	goto L1480
L1480:
	;
	v10657 = v4
	v10658 = base.B2i32(v10641 == v10519)
	goto L1473
L1481:
	;
	v10648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v10648 != 0 {
		v10657 = v4
		v10658 = v10643
		goto L1473
	} else {
		goto L1482
	}
L1482:
	;
	v10649 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10649 == int32(0) {
		v10976 = v4
		v10977 = v4
		v10987 = v10643
		goto L1472
	} else {
		goto L1483
	}
L1483:
	;
	v10652 = *(*int32)(unsafe.Add(mBase, uint32(v10649)+4))
	v10657 = base.B2i32(int32(1) < v10652)
	v10658 = v10643
	goto L1473
L1484:
	;
	v10662 = *(*int32)(unsafe.Add(mBase, uint32(v10659)+4))
	if v10662 <= int32(0) {
		v10976 = v4
		v10977 = v4
		v10987 = v10658
		goto L1472
	} else {
		goto L1485
	}
L1485:
	;
	v10674 = int32(0)
	v10682 = v4
	v10683 = v4
	goto L1486
L1486:
	;
	v10710 = *(*int32)(unsafe.Add(mBase, uint32(v10659)+12))
	v10714 = *(*int32)(unsafe.Add(mBase, uint32(v10710+v10674<<(uint(int32(2))%32))))
	v10716 = F_create_plan_recurse(m, l0, v10714, int32(1))
	mBase = m.M
	v10717 = m.ExcPending
	if v10717 != 0 {
		goto L1
	} else {
		goto L1489
	}
L1487:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10950 = m.ExcPending
	if v10950 != 0 {
		goto L1
	} else {
		goto L1541
	}
L1488:
	;
	goto L1487
L1489:
	;
	if v10559 == int32(0) {
		v10928 = v10716
		goto L1490
	} else {
		goto L1491
	}
L1490:
	;
	if v10657 != 0 {
		goto L1535
	} else {
		goto L1536
	}
L1491:
	;
	v10720 = *(*int32)(unsafe.Add(mBase, uint32(v10714)+8))
	v10721 = *(*int32)(unsafe.Add(mBase, uint32(v10720)+8))
	v10722 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v10734 = F_prepare_sort_from_pathkeys(m, v10716, v10559, v10721, v10722, int32(0), v47+int32(48), v47+int32(44), v47+int32(40), v47+int32(36), v47+int32(32))
	mBase = m.M
	v10735 = m.ExcPending
	if v10735 != 0 {
		goto L1
	} else {
		goto L1492
	}
L1492:
	;
	v10736 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v10737 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v10739 = v10737 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v10739) {
		goto L1496
	} else {
		goto L1497
	}
L1493:
	;
	if v10801 != 0 {
		goto L1488
	} else {
		goto L1511
	}
L1494:
	;
	v10801 = int32(0)
	goto L1493
L1495:
	;
	v10775 = v10770
	v10776 = v10771
	v10777 = v10772
	goto L1505
L1496:
	;
	if (v10736|v10722)&int32(3) != 0 {
		v10770 = v10736
		v10771 = v10722
		v10772 = v10739
		goto L1495
	} else {
		goto L1499
	}
L1497:
	;
	v10763 = v10736
	v10764 = v10722
	v10765 = v10739
	goto L1498
L1498:
	;
	if v10765 == int32(0) {
		goto L1494
	} else {
		goto L1504
	}
L1499:
	;
	v10747 = v10736
	v10748 = v10722
	v10749 = v10739
	goto L1500
L1500:
	;
	v10752 = *(*int32)(unsafe.Add(mBase, uint32(v10747)))
	v10753 = *(*int32)(unsafe.Add(mBase, uint32(v10748)))
	if v10752 != v10753 {
		v10770 = v10747
		v10771 = v10748
		v10772 = v10749
		goto L1495
	} else {
		goto L1502
	}
L1501:
	;
	v10763 = v10758
	v10764 = v10756
	v10765 = v10760
	goto L1498
L1502:
	;
	v10755 = int32(4)
	v10756 = v10748 + v10755
	v10758 = v10747 + v10755
	v10760 = v10749 - v10755
	if base.Ui32(int32(3)) < base.Ui32(v10760) {
		v10747 = v10758
		v10748 = v10756
		v10749 = v10760
		goto L1500
	} else {
		goto L1503
	}
L1503:
	;
	goto L1501
L1504:
	;
	v10770 = v10763
	v10771 = v10764
	v10772 = v10765
	goto L1495
L1505:
	;
	v10780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10775))))
	v10781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10776))))
	if v10780 == v10781 {
		goto L1507
	} else {
		goto L1508
	}
L1506:
	;
	v10801 = v10780 - v10781
	goto L1493
L1507:
	;
	v10783 = int32(1)
	v10788 = v10777 - v10783
	if v10788 != 0 {
		v10775 = v10775 + v10783
		v10776 = v10776 + v10783
		v10777 = v10788
		goto L1505
	} else {
		goto L1510
	}
L1508:
	;
	goto L1509
L1509:
	;
	goto L1506
L1510:
	;
	goto L1494
L1511:
	;
	v10802 = *(*int32)(unsafe.Add(mBase, uint32(v10714)+64))
	if v10559 == v10802 {
		goto L1513
	} else {
		goto L1514
	}
L1512:
	;
	if v10855 != 0 {
		goto L1530
	} else {
		goto L1531
	}
L1513:
	;
	v10855 = int32(1)
	goto L1512
L1514:
	;
	goto L1515
L1515:
	;
	v10811 = int32(0)
	goto L1517
L1516:
	;
	v10855 = v10847
	goto L1512
L1517:
	;
	v10815 = int32(0)
	if v10559 == v10815 {
		v10825 = v10815
		goto L1519
	} else {
		goto L1520
	}
L1518:
	;
	v10847 = int32(0)
	goto L1516
L1519:
	;
	if v10802 != 0 {
		goto L1523
	} else {
		goto L1524
	}
L1520:
	;
	v10819 = *(*int32)(unsafe.Add(mBase, uint32(v10559)+4))
	if v10819 <= v10811 {
		v10825 = int32(0)
		goto L1519
	} else {
		goto L1521
	}
L1521:
	;
	v10821 = *(*int32)(unsafe.Add(mBase, uint32(v10559)+12))
	v10825 = v10821 + v10811<<(uint(int32(2))%32)
	goto L1519
L1522:
	;
	v10831 = base.B2i32(v10825 == int32(0))
	if v10825 == int32(0) {
		v10847 = v10831
		goto L1516
	} else {
		goto L1527
	}
L1523:
	;
	v10826 = *(*int32)(unsafe.Add(mBase, uint32(v10802)+4))
	if v10811 < v10826 {
		goto L1522
	} else {
		goto L1526
	}
L1524:
	;
	goto L1525
L1525:
	;
	v10855 = base.B2i32(v10825 == int32(0))
	goto L1512
L1526:
	;
	goto L1525
L1527:
	;
	v10834 = *(*int32)(unsafe.Add(mBase, uint32(v10802)+12))
	v10837 = v10834 + v10811<<(uint(int32(2))%32)
	if v10837 == int32(0) {
		v10847 = v10831
		goto L1516
	} else {
		goto L1528
	}
L1528:
	;
	v10842 = *(*int32)(unsafe.Add(mBase, uint32(v10825)))
	v10843 = *(*int32)(unsafe.Add(mBase, uint32(v10837)))
	if v10842 == v10843 {
		v10811 = v10811 + int32(1)
		goto L1517
	} else {
		goto L1529
	}
L1529:
	;
	goto L1518
L1530:
	;
	v10928 = v10734
	goto L1490
L1531:
	;
	goto L1532
L1532:
	;
	v10856 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	v10857 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v10858 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v10860 = F_palloc0(m, int32(96))
	mBase = m.M
	v10861 = m.ExcPending
	if v10861 != 0 {
		goto L1
	} else {
		goto L1533
	}
L1533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10860))) = int32(362)
	v10864 = *(*int32)(unsafe.Add(mBase, uint32(v10734)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+44)) = v10864
	v10866 = *(*int32)(unsafe.Add(mBase, uint32(v10734)+4))
	v10867 = int32(4126313)
	v10868 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+88)) = v10858
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+84)) = v10857
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+80)) = v10856
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+76)) = v10736
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+72)) = v10737
	v10874 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+56)) = v10874
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+52)) = v10734
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+48)) = v10874
	v10879 = int32(1)
	v10881 = v10866 + (v10868 ^ v10879)
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+4)) = v10881
	v10884 = v47 + int32(56)
	v10885 = *(*float64)(unsafe.Add(mBase, uint32(v10734)+16))
	v10886 = *(*float64)(unsafe.Add(mBase, uint32(v10734)+24))
	v10887 = *(*int32)(unsafe.Add(mBase, uint32(v10734)+32))
	v10890 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v10891 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v10893 = m.G0
	v10894 = int32(16)
	v10895 = v10893 - v10894
	m.G0 = v10895
	F_cost_tuplesort(m, v10895+int32(8), v10895, v10886, v10887, float64(0), v10890, v10891)
	mBase = m.M
	v10900 = *(*float64)(unsafe.Add(mBase, uint32(v10895)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v10884)+32)) = v10886
	v10903 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	v10904 = base.F64_add(v10885, v10900)
	*(*float64)(unsafe.Add(mBase, uint32(v10884)+48)) = v10904
	*(*int32)(unsafe.Add(mBase, uint32(v10884)+40)) = v10881 + (v10903 ^ v10879)
	v10910 = *(*float64)(unsafe.Add(mBase, uint32(v10895)))
	*(*float64)(unsafe.Add(mBase, uint32(v10884)+56)) = base.F64_add(v10904, v10910)
	m.G0 = v10895 + v10894
	goto L1534
L1534:
	;
	v10916 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v10860)+8)) = v10916
	v10918 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v10860)+16)) = v10918
	v10920 = *(*float64)(unsafe.Add(mBase, uint32(v10734)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10860)+24)) = v10920
	v10922 = *(*int32)(unsafe.Add(mBase, uint32(v10734)+32))
	v10923 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10860)+36)) = uint8(v10923)
	*(*int32)(unsafe.Add(mBase, uint32(v10860)+32)) = v10922
	v10926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10734)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10860)+37)) = uint8(v10926)
	v10928 = v10860
	goto L1490
L1535:
	;
	v10937 = F_mark_async_capable_plan(m, v10928, v10714)
	mBase = m.M
	v10938 = m.ExcPending
	if v10938 != 0 {
		goto L1
	} else {
		goto L1538
	}
L1536:
	;
	v10940 = v10683
	goto L1537
L1537:
	;
	v10941 = F_lappend(m, v10682, v10928)
	mBase = m.M
	v10942 = m.ExcPending
	if v10942 != 0 {
		goto L1
	} else {
		goto L1539
	}
L1538:
	;
	v10940 = v10937 + v10683
	goto L1537
L1539:
	;
	v10944 = v10674 + int32(1)
	v10945 = *(*int32)(unsafe.Add(mBase, uint32(v10659)+4))
	if v10944 < v10945 {
		v10674 = v10944
		v10682 = v10941
		v10683 = v10940
		goto L1486
	} else {
		goto L1540
	}
L1540:
	;
	v10976 = v10941
	v10977 = v10940
	v10987 = v10658
	goto L1472
L1541:
	;
	F_errmsg_internal(m, int32(428109), int32(0))
	mBase = m.M
	v10954 = m.ExcPending
	if v10954 != 0 {
		goto L1
	} else {
		goto L1542
	}
L1542:
	;
	F_errfinish(m, int32(498183), int32(1347), int32(283997))
	mBase = m.M
	v10959 = m.ExcPending
	if v10959 != 0 {
		goto L1
	} else {
		goto L1543
	}
L1543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+80)) = v10977
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+76)) = v10976
	v11034 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+84)) = v11034
	v11036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+4)) = v11036
	v11038 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v10611)+8)) = v11038
	v11040 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v10611)+16)) = v11040
	v11042 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v10611)+24)) = v11042
	v11044 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11045 = *(*int32)(unsafe.Add(mBase, uint32(v11044)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+32)) = v11045
	v11047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10611)+36)) = uint8(v11047)
	v11049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10611)+37)) = uint8(v11049)
	if base.B2i32(l2&int32(3) == int32(0))|v10987 != 0 {
		v12585 = v10611
		v12589 = v47
		goto L3
	} else {
		goto L1555
	}
L1545:
	;
	v11010 = *(*int32)(unsafe.Add(mBase, uint32(v10558)+184))
	v11012 = F_extract_actual_clauses(m, v11010, int32(0))
	mBase = m.M
	v11013 = m.ExcPending
	if v11013 != 0 {
		goto L1
	} else {
		goto L1546
	}
L1546:
	;
	v11014 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11014 != 0 {
		goto L1547
	} else {
		goto L1548
	}
L1547:
	;
	v11015 = *(*int32)(unsafe.Add(mBase, uint32(v11014)+16))
	v11017 = F_extract_actual_clauses(m, v11015, int32(0))
	mBase = m.M
	v11018 = m.ExcPending
	if v11018 != 0 {
		goto L1
	} else {
		goto L1550
	}
L1548:
	;
	v11023 = v11012
	goto L1549
L1549:
	;
	if v11023 == int32(0) {
		goto L1544
	} else {
		goto L1553
	}
L1550:
	;
	v11019 = F_replace_nestloop_params_mutator(m, v11017, l0)
	mBase = m.M
	v11020 = m.ExcPending
	if v11020 != 0 {
		goto L1
	} else {
		goto L1551
	}
L1551:
	;
	v11021 = F_list_concat(m, v11012, v11019)
	mBase = m.M
	v11022 = m.ExcPending
	if v11022 != 0 {
		goto L1
	} else {
		goto L1552
	}
L1552:
	;
	v11023 = v11021
	goto L1549
L1553:
	;
	v11026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v11027 = F_make_partition_pruneinfo(m, l0, v10558, v11026, v11023)
	mBase = m.M
	v11028 = m.ExcPending
	if v11028 != 0 {
		goto L1
	} else {
		goto L1554
	}
L1554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10611)+88)) = v11027
	goto L1544
L1555:
	;
	v11056 = *(*int32)(unsafe.Add(mBase, uint32(v10611)+44))
	v11057 = F_list_copy_head(m, v11056, v10519)
	mBase = m.M
	v11058 = m.ExcPending
	if v11058 != 0 {
		goto L1
	} else {
		goto L1556
	}
L1556:
	;
	v11059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10611)+37)))
	v11060 = F_inject_projection_plan(m, v10611, v11057, v11059)
	mBase = m.M
	v11061 = m.ExcPending
	if v11061 != 0 {
		goto L1
	} else {
		goto L1557
	}
L1557:
	;
	v12585 = v11060
	v12589 = v47
	goto L3
L1558:
	;
	v11189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v11190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v11191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v11192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11193 = *(*int32)(unsafe.Add(mBase, uint32(v11192)+8))
	v11194 = F_reparameterize_path_by_child(m, l0, v11191, v11193)
	mBase = m.M
	v11195 = m.ExcPending
	if v11195 != 0 {
		goto L1
	} else {
		goto L1573
	}
L1559:
	;
	v11068 = *(*int32)(unsafe.Add(mBase, uint32(v11064)+4))
	if v11068 <= int32(0) {
		v11147 = v11062
		goto L1558
	} else {
		goto L1560
	}
L1560:
	;
	v11071 = *(*int32)(unsafe.Add(mBase, uint32(v11063)+8))
	v11076 = v11062
	v11077 = int32(1)
	v11079 = v4
	goto L1561
L1561:
	;
	v11118 = *(*int32)(unsafe.Add(mBase, uint32(v11064)+12))
	v11122 = *(*int32)(unsafe.Add(mBase, uint32(v11118+v11079<<(uint(int32(2))%32))))
	v11123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11123 != 0 {
		goto L1563
	} else {
		goto L1564
	}
L1562:
	;
	v11147 = v11139
	goto L1558
L1563:
	;
	v11124 = F_replace_nestloop_params_mutator(m, v11122, l0)
	mBase = m.M
	v11125 = m.ExcPending
	if v11125 != 0 {
		goto L1
	} else {
		goto L1566
	}
L1564:
	;
	v11126 = v11122
	goto L1565
L1565:
	;
	v11128 = int32(0)
	v11130 = F_makeTargetEntry(m, v11126, base.I32_extend16_s(v11077), v11128, v11128)
	mBase = m.M
	v11131 = m.ExcPending
	if v11131 != 0 {
		goto L1
	} else {
		goto L1567
	}
L1566:
	;
	v11126 = v11124
	goto L1565
L1567:
	;
	if v11071 != 0 {
		goto L1568
	} else {
		goto L1569
	}
L1568:
	;
	v11135 = *(*int32)(unsafe.Add(mBase, uint32(v11071-int32(4)+v11077<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11130)+16)) = v11135
	goto L1570
L1569:
	;
	goto L1570
L1570:
	;
	v11139 = F_lappend(m, v11076, v11130)
	mBase = m.M
	v11140 = m.ExcPending
	if v11140 != 0 {
		goto L1
	} else {
		goto L1571
	}
L1571:
	;
	v11142 = v11079 + int32(1)
	v11143 = *(*int32)(unsafe.Add(mBase, uint32(v11064)+4))
	if v11142 < v11143 {
		v11076 = v11139
		v11077 = v11077 + int32(1)
		v11079 = v11142
		goto L1561
	} else {
		goto L1572
	}
L1572:
	;
	goto L1562
L1573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v11194
	v11197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11199 = F_create_plan_recurse(m, l0, v11197, int32(0))
	mBase = m.M
	v11200 = m.ExcPending
	if v11200 != 0 {
		goto L1
	} else {
		goto L1574
	}
L1574:
	;
	v11201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v11202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11203 = *(*int32)(unsafe.Add(mBase, uint32(v11202)+8))
	v11204 = *(*int32)(unsafe.Add(mBase, uint32(v11203)+8))
	v11205 = F_bms_union(m, v11201, v11204)
	mBase = m.M
	v11206 = m.ExcPending
	if v11206 != 0 {
		goto L1
	} else {
		goto L1575
	}
L1575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v11205
	v11208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v11210 = F_create_plan_recurse(m, l0, v11208, int32(0))
	mBase = m.M
	v11211 = m.ExcPending
	if v11211 != 0 {
		goto L1
	} else {
		goto L1576
	}
L1576:
	;
	v11212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	F_bms_free(m, v11212)
	mBase = m.M
	v11214 = m.ExcPending
	if v11214 != 0 {
		goto L1
	} else {
		goto L1577
	}
L1577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v11190
	v11216 = F_order_qual_clauses(m, l0, v11189)
	mBase = m.M
	v11217 = m.ExcPending
	if v11217 != 0 {
		goto L1
	} else {
		goto L1578
	}
L1578:
	;
	v11219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v11219)%32)&int32(174) != 0 {
		goto L1580
	} else {
		goto L1581
	}
L1579:
	;
	v11238 = int32(0)
	v11239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11239 == v11238 {
		v11254 = v11238
		goto L1585
	} else {
		goto L1586
	}
L1580:
	;
	v11223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11224 = *(*int32)(unsafe.Add(mBase, uint32(v11223)+8))
	F_extract_actual_join_clauses(m, v11216, v11224, v47+int32(56), v47+int32(140))
	mBase = m.M
	v11230 = m.ExcPending
	if v11230 != 0 {
		goto L1
	} else {
		goto L1583
	}
L1581:
	;
	goto L1582
L1582:
	;
	v11232 = F_extract_actual_clauses(m, v11216, int32(0))
	mBase = m.M
	v11233 = m.ExcPending
	if v11233 != 0 {
		goto L1
	} else {
		goto L1584
	}
L1583:
	;
	goto L1579
L1584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v11232
	goto L1579
L1585:
	;
	if v11254 != 0 {
		goto L1592
	} else {
		goto L1593
	}
L1586:
	;
	v11242 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v11243 = F_replace_nestloop_params_mutator(m, v11242, l0)
	mBase = m.M
	v11244 = m.ExcPending
	if v11244 != 0 {
		goto L1
	} else {
		goto L1587
	}
L1587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v11243
	v11246 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v11247 = F_replace_nestloop_params_mutator(m, v11246, l0)
	mBase = m.M
	v11248 = m.ExcPending
	if v11248 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v11247
	v11250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11250 == int32(0) {
		v11254 = v11238
		goto L1585
	} else {
		goto L1589
	}
L1589:
	;
	v11253 = *(*int32)(unsafe.Add(mBase, uint32(v11250)+4))
	v11254 = v11253
	goto L1585
L1590:
	;
	v12068 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v12069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v12070 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12071 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12073 = F_palloc0(m, int32(96))
	mBase = m.M
	v12074 = m.ExcPending
	if v12074 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1591:
	;
	if v11872 == int32(0) {
		v12032 = v11199
		goto L1590
	} else {
		goto L1685
	}
L1592:
	;
	v11256 = F_bms_union(m, v11204, v11254)
	mBase = m.M
	v11257 = m.ExcPending
	if v11257 != 0 {
		goto L1
	} else {
		goto L1595
	}
L1593:
	;
	v11258 = v11204
	goto L1594
L1594:
	;
	v11259 = int32(0)
	v11260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v11260 == v11259 {
		v11872 = v11259
		goto L1591
	} else {
		goto L1596
	}
L1595:
	;
	v11258 = v11256
	goto L1594
L1596:
	;
	v11267 = int32(0)
	v11270 = v11260
	v11279 = v4
	goto L1597
L1597:
	;
	v11308 = *(*int32)(unsafe.Add(mBase, uint32(v11270)+4))
	if v11267 < v11308 {
		goto L1599
	} else {
		goto L1600
	}
L1598:
	;
	v11872 = v11799
	goto L1591
L1599:
	;
	v11310 = *(*int32)(unsafe.Add(mBase, uint32(v11270)+12))
	v11314 = *(*int32)(unsafe.Add(mBase, uint32(v11310+v11267<<(uint(int32(2))%32))))
	v11315 = *(*int32)(unsafe.Add(mBase, uint32(v11314)+8))
	v11316 = *(*int32)(unsafe.Add(mBase, uint32(v11315)))
	if v11316 == int32(6) {
		goto L1605
	} else {
		goto L1606
	}
L1600:
	;
	v11799 = v11279
	goto L1601
L1601:
	;
	goto L1598
L1602:
	;
	if v11744 != 0 {
		v11267 = v11741 + int32(1)
		v11270 = v11744
		v11279 = v11753
		goto L1597
	} else {
		goto L1684
	}
L1603:
	;
	v11736 = F_lappend(m, v11279, v11314)
	mBase = m.M
	v11737 = m.ExcPending
	if v11737 != 0 {
		goto L1
	} else {
		goto L1683
	}
L1604:
	;
	v11676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11677 = *(*int32)(unsafe.Add(mBase, uint32(v11322)+4))
	v11681 = *(*int32)(unsafe.Add(mBase, uint32(v11676+v11677<<(uint(int32(2))%32))))
	v11682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v11683 = F_list_delete_nth_cell(m, v11682, v11267)
	mBase = m.M
	v11684 = m.ExcPending
	if v11684 != 0 {
		goto L1
	} else {
		goto L1681
	}
L1605:
	;
	v11319 = *(*int32)(unsafe.Add(mBase, uint32(v11315)+4))
	v11320 = F_bms_is_member(m, v11319, v11204)
	mBase = m.M
	v11321 = m.ExcPending
	if v11321 != 0 {
		goto L1
	} else {
		goto L1608
	}
L1606:
	;
	v11324 = v11315
	v11326 = v11316
	goto L1607
L1607:
	;
	if v11326 != int32(319) {
		v11741 = v11267
		v11744 = v11270
		v11753 = v11279
		goto L1602
	} else {
		goto L1610
	}
L1608:
	;
	v11322 = *(*int32)(unsafe.Add(mBase, uint32(v11314)+8))
	if v11320 != 0 {
		goto L1604
	} else {
		goto L1609
	}
L1609:
	;
	v11323 = *(*int32)(unsafe.Add(mBase, uint32(v11322)))
	v11324 = v11322
	v11326 = v11323
	goto L1607
L1610:
	;
	v11329 = F_find_placeholder_info(m, l0, v11324)
	mBase = m.M
	v11330 = m.ExcPending
	if v11330 != 0 {
		goto L1
	} else {
		goto L1611
	}
L1611:
	;
	v11331 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+12))
	v11332 = int32(0)
	if v11331 == v11332 {
		goto L1613
	} else {
		goto L1614
	}
L1612:
	;
	if v11385 == int32(0) {
		v11741 = v11267
		v11744 = v11270
		v11753 = v11279
		goto L1602
	} else {
		goto L1626
	}
L1613:
	;
	v11385 = int32(1)
	goto L1612
L1614:
	;
	goto L1615
L1615:
	;
	if v11258 == int32(0) {
		v11376 = v11332
		goto L1616
	} else {
		goto L1617
	}
L1616:
	;
	v11385 = v11376
	goto L1612
L1617:
	;
	v11341 = *(*int32)(unsafe.Add(mBase, uint32(v11331)+4))
	v11342 = *(*int32)(unsafe.Add(mBase, uint32(v11258)+4))
	if v11342 < v11341 {
		v11376 = v11332
		goto L1616
	} else {
		goto L1618
	}
L1618:
	;
	v11344 = int32(1)
	if v11341 <= v11344 {
		goto L1619
	} else {
		goto L1620
	}
L1619:
	;
	v11347 = v11344
	goto L1621
L1620:
	;
	v11347 = v11341
	goto L1621
L1621:
	;
	v11348 = int32(8)
	v11353 = int32(0)
	goto L1622
L1622:
	;
	v11360 = v11353 << (uint(int32(2)) % 32)
	v11362 = *(*int32)(unsafe.Add(mBase, uint32(v11331+v11348+v11360)))
	v11364 = *(*int32)(unsafe.Add(mBase, uint32(v11360+(v11258+v11348))))
	v11367 = v11362 & (v11364 ^ int32(-1))
	v11369 = base.B2i32(v11367 == int32(0))
	if v11367 != 0 {
		v11376 = v11369
		goto L1616
	} else {
		goto L1624
	}
L1623:
	;
	v11376 = v11369
	goto L1616
L1624:
	;
	v11371 = v11353 + int32(1)
	if v11371 != v11347 {
		v11353 = v11371
		goto L1622
	} else {
		goto L1625
	}
L1625:
	;
	goto L1623
L1626:
	;
	v11388 = int32(0)
	if v11331 == v11388 {
		v11429 = v11388
		goto L1628
	} else {
		goto L1629
	}
L1627:
	;
	if v11429 == int32(0) {
		v11741 = v11267
		v11744 = v11270
		v11753 = v11279
		goto L1602
	} else {
		goto L1641
	}
L1628:
	;
	goto L1627
L1629:
	;
	if v11204 == int32(0) {
		v11429 = v11388
		goto L1628
	} else {
		goto L1630
	}
L1630:
	;
	v11397 = *(*int32)(unsafe.Add(mBase, uint32(v11331)+4))
	v11398 = *(*int32)(unsafe.Add(mBase, uint32(v11204)+4))
	if v11397 < v11398 {
		goto L1631
	} else {
		goto L1632
	}
L1631:
	;
	v11400 = v11397
	goto L1633
L1632:
	;
	v11400 = v11398
	goto L1633
L1633:
	;
	if v11400 <= int32(1) {
		goto L1634
	} else {
		goto L1635
	}
L1634:
	;
	v11403 = int32(1)
	goto L1636
L1635:
	;
	v11403 = v11400
	goto L1636
L1636:
	;
	v11404 = int32(8)
	v11409 = int32(0)
	goto L1637
L1637:
	;
	v11416 = v11409 << (uint(int32(2)) % 32)
	v11418 = *(*int32)(unsafe.Add(mBase, uint32(v11204+v11404+v11416)))
	v11420 = *(*int32)(unsafe.Add(mBase, uint32(v11416+(v11331+v11404))))
	v11421 = v11418 & v11420
	v11423 = base.B2i32(v11421 != int32(0))
	if v11421 != 0 {
		v11429 = v11423
		goto L1628
	} else {
		goto L1639
	}
L1638:
	;
	v11429 = v11423
	goto L1628
L1639:
	;
	v11425 = v11409 + int32(1)
	if v11425 != v11403 {
		v11409 = v11425
		goto L1637
	} else {
		goto L1640
	}
L1640:
	;
	goto L1638
L1641:
	;
	v11435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v11436 = F_list_delete_nth_cell(m, v11435, v11267)
	mBase = m.M
	v11437 = m.ExcPending
	if v11437 != 0 {
		goto L1
	} else {
		goto L1642
	}
L1642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v11436
	v11439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11439)+39)))
	if v11440 == int32(1) {
		goto L1643
	} else {
		goto L1644
	}
L1643:
	;
	v11443 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+8))
	v11444 = F_copyObjectImpl(m, v11443)
	mBase = m.M
	v11445 = m.ExcPending
	if v11445 != 0 {
		goto L1
	} else {
		goto L1646
	}
L1644:
	;
	v11447 = v11324
	goto L1645
L1645:
	;
	v11448 = int32(0)
	v11449 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+12))
	if v11449 == v11448 {
		goto L1649
	} else {
		goto L1650
	}
L1646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11314)+8)) = v11444
	v11447 = v11444
	goto L1645
L1647:
	;
	if int32(0) < v11506 {
		goto L1658
	} else {
		goto L1659
	}
L1648:
	;
	v11506 = base.I32_ctz(v11492) | v11493<<(uint(int32(5))%32)
	goto L1647
L1649:
	;
	v11506 = int32(-2)
	goto L1647
L1650:
	;
	v11459 = base.I32_div_s(int32(0), int32(32))
	v11460 = *(*int32)(unsafe.Add(mBase, uint32(v11449)+4))
	if v11460 <= v11459 {
		goto L1649
	} else {
		goto L1651
	}
L1651:
	;
	v11463 = v11449 + int32(8)
	v11467 = *(*int32)(unsafe.Add(mBase, uint32(v11463+v11459<<(uint(int32(2))%32))))
	v11470 = v11467 & int32(-1)
	if v11470 != 0 {
		v11492 = v11470
		v11493 = v11459
		goto L1648
	} else {
		goto L1652
	}
L1652:
	;
	v11472 = v11459 + int32(1)
	if v11472 == v11460 {
		goto L1649
	} else {
		goto L1653
	}
L1653:
	;
	v11475 = v11472
	goto L1654
L1654:
	;
	v11482 = *(*int32)(unsafe.Add(mBase, uint32(v11463+v11475<<(uint(int32(2))%32))))
	if v11482 != 0 {
		v11492 = v11482
		v11493 = v11475
		goto L1648
	} else {
		goto L1656
	}
L1655:
	;
	goto L1649
L1656:
	;
	v11484 = v11475 + int32(1)
	if v11484 != v11460 {
		v11475 = v11484
		goto L1654
	} else {
		goto L1657
	}
L1657:
	;
	goto L1655
L1658:
	;
	v11518 = v11448
	v11534 = v11506
	goto L1661
L1659:
	;
	v11635 = v11448
	goto L1660
L1660:
	;
	v11670 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+12))
	v11671 = F_bms_del_members(m, v11635, v11670)
	mBase = m.M
	v11672 = m.ExcPending
	if v11672 != 0 {
		goto L1
	} else {
		goto L1679
	}
L1661:
	;
	v11553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v11534 == v11553 {
		v11566 = v11518
		goto L1663
	} else {
		goto L1664
	}
L1662:
	;
	v11635 = v11566
	goto L1660
L1663:
	;
	v11567 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+12))
	if v11567 == int32(0) {
		goto L1669
	} else {
		goto L1670
	}
L1664:
	;
	v11555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11559 = *(*int32)(unsafe.Add(mBase, uint32(v11555+v11534<<(uint(int32(2))%32))))
	if v11559 == int32(0) {
		v11566 = v11518
		goto L1663
	} else {
		goto L1665
	}
L1665:
	;
	v11562 = *(*int32)(unsafe.Add(mBase, uint32(v11559)+96))
	v11563 = F_bms_add_members(m, v11518, v11562)
	mBase = m.M
	v11564 = m.ExcPending
	if v11564 != 0 {
		goto L1
	} else {
		goto L1666
	}
L1666:
	;
	v11566 = v11563
	goto L1663
L1667:
	;
	if int32(0) < v11623 {
		v11518 = v11566
		v11534 = v11623
		goto L1661
	} else {
		goto L1678
	}
L1668:
	;
	v11623 = base.I32_ctz(v11609) | v11610<<(uint(int32(5))%32)
	goto L1667
L1669:
	;
	v11623 = int32(-2)
	goto L1667
L1670:
	;
	v11574 = v11534 + int32(1)
	v11576 = base.I32_div_s(v11574, int32(32))
	v11577 = *(*int32)(unsafe.Add(mBase, uint32(v11567)+4))
	if v11577 <= v11576 {
		goto L1669
	} else {
		goto L1671
	}
L1671:
	;
	v11580 = v11567 + int32(8)
	v11584 = *(*int32)(unsafe.Add(mBase, uint32(v11580+v11576<<(uint(int32(2))%32))))
	v11587 = v11584 & (int32(-1) << (uint(v11574) % 32))
	if v11587 != 0 {
		v11609 = v11587
		v11610 = v11576
		goto L1668
	} else {
		goto L1672
	}
L1672:
	;
	v11589 = v11576 + int32(1)
	if v11589 == v11577 {
		goto L1669
	} else {
		goto L1673
	}
L1673:
	;
	v11592 = v11589
	goto L1674
L1674:
	;
	v11599 = *(*int32)(unsafe.Add(mBase, uint32(v11580+v11592<<(uint(int32(2))%32))))
	if v11599 != 0 {
		v11609 = v11599
		v11610 = v11592
		goto L1668
	} else {
		goto L1676
	}
L1675:
	;
	goto L1669
L1676:
	;
	v11601 = v11592 + int32(1)
	if v11601 != v11577 {
		v11592 = v11601
		goto L1674
	} else {
		goto L1677
	}
L1677:
	;
	goto L1675
L1678:
	;
	goto L1662
L1679:
	;
	v11673 = F_bms_intersect(m, v11671, v11204)
	mBase = m.M
	v11674 = m.ExcPending
	if v11674 != 0 {
		goto L1
	} else {
		goto L1680
	}
L1680:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11447)+12)) = v11673
	v11696 = v11436
	goto L1603
L1681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v11683
	v11686 = *(*int32)(unsafe.Add(mBase, uint32(v11681)+96))
	v11687 = F_bms_intersect(m, v11686, v11204)
	mBase = m.M
	v11688 = m.ExcPending
	if v11688 != 0 {
		goto L1
	} else {
		goto L1682
	}
L1682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11322)+24)) = v11687
	v11696 = v11683
	goto L1603
L1683:
	;
	v11741 = v11267 - int32(1)
	v11744 = v11696
	v11753 = v11736
	goto L1602
L1684:
	;
	v11799 = v11753
	goto L1601
L1685:
	;
	v11875 = int32(0)
	v11876 = *(*int32)(unsafe.Add(mBase, uint32(v11199)+44))
	v11877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11199)+37)))
	v11878 = *(*int32)(unsafe.Add(mBase, uint32(v11872)+4))
	if v11875 < v11878 {
		goto L1686
	} else {
		goto L1687
	}
L1686:
	;
	v11884 = v11875
	v11887 = v11876
	v11892 = v11877
	goto L1689
L1687:
	;
	v11980 = v11876
	v11985 = v11877
	goto L1688
L1688:
	;
	v12018 = *(*int32)(unsafe.Add(mBase, uint32(v11199)+44))
	if v11980 == v12018 {
		v12032 = v11199
		goto L1590
	} else {
		goto L1709
	}
L1689:
	;
	v11925 = *(*int32)(unsafe.Add(mBase, uint32(v11872)+12))
	v11929 = *(*int32)(unsafe.Add(mBase, uint32(v11925+v11884<<(uint(int32(2))%32))))
	v11930 = *(*int32)(unsafe.Add(mBase, uint32(v11929)+8))
	v11931 = *(*int32)(unsafe.Add(mBase, uint32(v11930)))
	if v11931 == int32(6) {
		v11966 = v11887
		v11968 = v11892
		goto L1691
	} else {
		goto L1692
	}
L1690:
	;
	v11980 = v11966
	v11985 = v11968
	goto L1688
L1691:
	;
	v11971 = v11884 + int32(1)
	v11972 = *(*int32)(unsafe.Add(mBase, uint32(v11872)+4))
	if v11971 < v11972 {
		v11884 = v11971
		v11887 = v11966
		v11892 = v11968
		goto L1689
	} else {
		goto L1708
	}
L1692:
	;
	v11934 = F_tlist_member(m, v11930, v11887)
	mBase = m.M
	v11935 = m.ExcPending
	if v11935 != 0 {
		goto L1
	} else {
		goto L1693
	}
L1693:
	;
	if v11934 != 0 {
		v11966 = v11887
		v11968 = v11892
		goto L1691
	} else {
		goto L1694
	}
L1694:
	;
	v11936 = *(*int32)(unsafe.Add(mBase, uint32(v11930)+4))
	v11937 = F_replace_nestloop_params_mutator(m, v11936, l0)
	mBase = m.M
	v11938 = m.ExcPending
	if v11938 != 0 {
		goto L1
	} else {
		goto L1695
	}
L1695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11930)+4)) = v11937
	v11940 = *(*int32)(unsafe.Add(mBase, uint32(v11199)+44))
	if v11940 == v11887 {
		goto L1696
	} else {
		goto L1697
	}
L1696:
	;
	v11942 = F_list_copy(m, v11887)
	mBase = m.M
	v11943 = m.ExcPending
	if v11943 != 0 {
		goto L1
	} else {
		goto L1699
	}
L1697:
	;
	v11944 = v11887
	goto L1698
L1698:
	;
	v11945 = F_copyObjectImpl(m, v11930)
	mBase = m.M
	v11946 = m.ExcPending
	if v11946 != 0 {
		goto L1
	} else {
		goto L1700
	}
L1699:
	;
	v11944 = v11942
	goto L1698
L1700:
	;
	if v11944 != 0 {
		goto L1701
	} else {
		goto L1702
	}
L1701:
	;
	v11950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11944)+4)))
	v11954 = v11950 + int32(1)
	goto L1703
L1702:
	;
	v11954 = int32(1)
	goto L1703
L1703:
	;
	v11958 = F_makeTargetEntry(m, v11945, base.I32_extend16_s(v11954), int32(0), int32(1))
	mBase = m.M
	v11959 = m.ExcPending
	if v11959 != 0 {
		goto L1
	} else {
		goto L1704
	}
L1704:
	;
	v11960 = F_lappend(m, v11944, v11958)
	mBase = m.M
	v11961 = m.ExcPending
	if v11961 != 0 {
		goto L1
	} else {
		goto L1705
	}
L1705:
	;
	if v11892&int32(1) == int32(0) {
		v11966 = v11960
		v11968 = int32(0)
		goto L1691
	} else {
		goto L1706
	}
L1706:
	;
	v11964 = F_is_parallel_safe(m, l0, v11930)
	mBase = m.M
	v11965 = m.ExcPending
	if v11965 != 0 {
		goto L1
	} else {
		goto L1707
	}
L1707:
	;
	v11966 = v11960
	v11968 = v11964
	goto L1691
L1708:
	;
	goto L1690
L1709:
	;
	v12022 = F_change_plan_targetlist(m, v11199, v11980, v11985&int32(1))
	mBase = m.M
	v12023 = m.ExcPending
	if v12023 != 0 {
		goto L1
	} else {
		goto L1710
	}
L1710:
	;
	v12032 = v12022
	goto L1590
L1711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+88)) = v11872
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+80)) = v12071
	*(*uint8)(unsafe.Add(mBase, uint32(v12073)+76)) = uint8(v12069)
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+72)) = v12068
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+56)) = v11210
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+52)) = v12032
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+48)) = v12070
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+44)) = v11147
	*(*int32)(unsafe.Add(mBase, uint32(v12073))) = int32(356)
	v12085 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+4)) = v12085
	v12087 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v12073)+8)) = v12087
	v12089 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v12073)+16)) = v12089
	v12091 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12073)+24)) = v12091
	v12093 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12094 = *(*int32)(unsafe.Add(mBase, uint32(v12093)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12073)+32)) = v12094
	v12096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12073)+36)) = uint8(v12096)
	v12098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12073)+37)) = uint8(v12098)
	v12528 = v12073
	goto L4
L1712:
	;
	v12229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v12231 = int32(1)
	v12235 = F_create_plan_recurse(m, l0, v12229, base.B2i32(v12231 < v12230)<<(uint(v12231)%32))
	mBase = m.M
	v12236 = m.ExcPending
	if v12236 != 0 {
		goto L1
	} else {
		goto L1731
	}
L1713:
	;
	v12187 = int32(0)
	goto L1712
L1714:
	;
	goto L1715
L1715:
	;
	v12106 = *(*int32)(unsafe.Add(mBase, uint32(v12101)+4))
	if v12106 <= int32(0) {
		goto L1716
	} else {
		goto L1717
	}
L1716:
	;
	v12187 = int32(0)
	goto L1712
L1717:
	;
	goto L1718
L1718:
	;
	v12110 = *(*int32)(unsafe.Add(mBase, uint32(v12100)+8))
	v12116 = int32(0)
	v12117 = int32(1)
	v12119 = v4
	goto L1719
L1719:
	;
	v12158 = *(*int32)(unsafe.Add(mBase, uint32(v12101)+12))
	v12162 = *(*int32)(unsafe.Add(mBase, uint32(v12158+v12119<<(uint(int32(2))%32))))
	v12163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12163 != 0 {
		goto L1721
	} else {
		goto L1722
	}
L1720:
	;
	v12187 = v12179
	goto L1712
L1721:
	;
	v12164 = F_replace_nestloop_params_mutator(m, v12162, l0)
	mBase = m.M
	v12165 = m.ExcPending
	if v12165 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1722:
	;
	v12166 = v12162
	goto L1723
L1723:
	;
	v12168 = int32(0)
	v12170 = F_makeTargetEntry(m, v12166, base.I32_extend16_s(v12117), v12168, v12168)
	mBase = m.M
	v12171 = m.ExcPending
	if v12171 != 0 {
		goto L1
	} else {
		goto L1725
	}
L1724:
	;
	v12166 = v12164
	goto L1723
L1725:
	;
	if v12110 != 0 {
		goto L1726
	} else {
		goto L1727
	}
L1726:
	;
	v12175 = *(*int32)(unsafe.Add(mBase, uint32(v12110-int32(4)+v12117<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12170)+16)) = v12175
	goto L1728
L1727:
	;
	goto L1728
L1728:
	;
	v12179 = F_lappend(m, v12116, v12170)
	mBase = m.M
	v12180 = m.ExcPending
	if v12180 != 0 {
		goto L1
	} else {
		goto L1729
	}
L1729:
	;
	v12182 = v12119 + int32(1)
	v12183 = *(*int32)(unsafe.Add(mBase, uint32(v12101)+4))
	if v12182 < v12183 {
		v12116 = v12179
		v12117 = v12117 + int32(1)
		v12119 = v12182
		goto L1719
	} else {
		goto L1730
	}
L1730:
	;
	goto L1720
L1731:
	;
	v12237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v12239 = F_create_plan_recurse(m, l0, v12237, int32(2))
	mBase = m.M
	v12240 = m.ExcPending
	if v12240 != 0 {
		goto L1
	} else {
		goto L1732
	}
L1732:
	;
	v12241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12242 = F_order_qual_clauses(m, l0, v12241)
	mBase = m.M
	v12243 = m.ExcPending
	if v12243 != 0 {
		goto L1
	} else {
		goto L1733
	}
L1733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12242
	v12246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v12246)%32)&int32(174) != 0 {
		goto L1735
	} else {
		goto L1736
	}
L1734:
	;
	v12265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12266 = F_get_actual_clauses(m, v12265)
	mBase = m.M
	v12267 = m.ExcPending
	if v12267 != 0 {
		goto L1
	} else {
		goto L1740
	}
L1735:
	;
	v12250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12251 = *(*int32)(unsafe.Add(mBase, uint32(v12250)+8))
	F_extract_actual_join_clauses(m, v12242, v12251, v47+int32(56), v47+int32(140))
	mBase = m.M
	v12257 = m.ExcPending
	if v12257 != 0 {
		goto L1
	} else {
		goto L1738
	}
L1736:
	;
	goto L1737
L1737:
	;
	v12259 = F_extract_actual_clauses(m, v12242, int32(0))
	mBase = m.M
	v12260 = m.ExcPending
	if v12260 != 0 {
		goto L1
	} else {
		goto L1739
	}
L1738:
	;
	goto L1734
L1739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12259
	goto L1734
L1740:
	;
	v12268 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12269 = F_list_difference(m, v12268, v12266)
	mBase = m.M
	v12270 = m.ExcPending
	if v12270 != 0 {
		goto L1
	} else {
		goto L1741
	}
L1741:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12269
	v12272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12272 != 0 {
		goto L1742
	} else {
		goto L1743
	}
L1742:
	;
	v12273 = F_replace_nestloop_params_mutator(m, v12269, l0)
	mBase = m.M
	v12274 = m.ExcPending
	if v12274 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1743:
	;
	goto L1744
L1744:
	;
	v12280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12282 = *(*int32)(unsafe.Add(mBase, uint32(v12281)+8))
	v12283 = *(*int32)(unsafe.Add(mBase, uint32(v12282)+8))
	v12284 = F_get_switched_clauses(m, v12280, v12283)
	mBase = m.M
	v12285 = m.ExcPending
	if v12285 != 0 {
		goto L1
	} else {
		goto L1750
	}
L1745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12273
	v12276 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12277 = F_replace_nestloop_params_mutator(m, v12276, l0)
	mBase = m.M
	v12278 = m.ExcPending
	if v12278 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v12277
	goto L1744
L1747:
	;
	v12453 = F_palloc0(m, int32(96))
	mBase = m.M
	v12454 = m.ExcPending
	if v12454 != 0 {
		goto L1
	} else {
		goto L1770
	}
L1748:
	;
	v12336 = int32(0)
	v12343 = v12336
	v12345 = v12289
	v12347 = v12336
	v12350 = v12336
	v12353 = v4
	goto L1763
L1749:
	;
	v12334 = int32(0)
	v12414 = v12330
	v12416 = v12334
	v12419 = v12334
	v12421 = v12331
	v12422 = v4
	v12425 = v12332
	v12427 = v12333
	goto L1747
L1750:
	;
	if v12284 == int32(0) {
		goto L1751
	} else {
		goto L1752
	}
L1751:
	;
	v12330 = int32(0)
	v12331 = v4
	v12332 = v4
	v12333 = v4
	goto L1749
L1752:
	;
	goto L1753
L1753:
	;
	v12289 = int32(0)
	v12290 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+4))
	if v12290 != int32(1) {
		goto L1755
	} else {
		goto L1756
	}
L1754:
	;
	v12325 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+4))
	if int32(0) < v12325 {
		goto L1748
	} else {
		goto L1762
	}
L1755:
	;
	v12322 = v4
	v12323 = v4
	v12324 = int32(0)
	goto L1754
L1756:
	;
	v12293 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+12))
	v12294 = *(*int32)(unsafe.Add(mBase, uint32(v12293)))
	v12295 = *(*int32)(unsafe.Add(mBase, uint32(v12294)+28))
	v12296 = *(*int32)(unsafe.Add(mBase, uint32(v12295)+12))
	v12297 = *(*int32)(unsafe.Add(mBase, uint32(v12296)))
	v12298 = *(*int32)(unsafe.Add(mBase, uint32(v12297)))
	if v12298 == int32(27) {
		goto L1757
	} else {
		goto L1758
	}
L1757:
	;
	v12301 = *(*int32)(unsafe.Add(mBase, uint32(v12297)+4))
	v12302 = *(*int32)(unsafe.Add(mBase, uint32(v12301)))
	v12303 = v12301
	v12304 = v12302
	goto L1759
L1758:
	;
	v12303 = v12297
	v12304 = v12298
	goto L1759
L1759:
	;
	if v12304 != int32(6) {
		goto L1755
	} else {
		goto L1760
	}
L1760:
	;
	v12307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12308 = *(*int32)(unsafe.Add(mBase, uint32(v12303)+4))
	v12312 = *(*int32)(unsafe.Add(mBase, uint32(v12307+v12308<<(uint(int32(2))%32))))
	v12313 = *(*int32)(unsafe.Add(mBase, uint32(v12312)+12))
	if v12313 != 0 {
		goto L1755
	} else {
		goto L1761
	}
L1761:
	;
	v12314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12303)+8)))
	v12315 = *(*int32)(unsafe.Add(mBase, uint32(v12312)+16))
	v12316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12312)+20)))
	v12322 = v12314
	v12323 = v12315
	v12324 = v12316
	goto L1754
L1762:
	;
	v12330 = v12289
	v12331 = v12324
	v12332 = v12322
	v12333 = v12323
	goto L1749
L1763:
	;
	v12383 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+12))
	v12387 = *(*int32)(unsafe.Add(mBase, uint32(v12383+v12343<<(uint(int32(2))%32))))
	v12388 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+4))
	v12389 = F_lappend_oid(m, v12353, v12388)
	mBase = m.M
	v12390 = m.ExcPending
	if v12390 != 0 {
		goto L1
	} else {
		goto L1765
	}
L1764:
	;
	v12414 = v12397
	v12416 = v12402
	v12419 = v12392
	v12421 = v12324
	v12422 = v12389
	v12425 = v12322
	v12427 = v12323
	goto L1747
L1765:
	;
	v12391 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+24))
	v12392 = F_lappend_oid(m, v12350, v12391)
	mBase = m.M
	v12393 = m.ExcPending
	if v12393 != 0 {
		goto L1
	} else {
		goto L1766
	}
L1766:
	;
	v12394 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+28))
	v12395 = *(*int32)(unsafe.Add(mBase, uint32(v12394)+12))
	v12396 = *(*int32)(unsafe.Add(mBase, uint32(v12395)))
	v12397 = F_lappend(m, v12345, v12396)
	mBase = m.M
	v12398 = m.ExcPending
	if v12398 != 0 {
		goto L1
	} else {
		goto L1767
	}
L1767:
	;
	v12399 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+28))
	v12400 = *(*int32)(unsafe.Add(mBase, uint32(v12399)+12))
	v12401 = *(*int32)(unsafe.Add(mBase, uint32(v12400)+4))
	v12402 = F_lappend(m, v12347, v12401)
	mBase = m.M
	v12403 = m.ExcPending
	if v12403 != 0 {
		goto L1
	} else {
		goto L1768
	}
L1768:
	;
	v12405 = v12343 + int32(1)
	v12406 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+4))
	if v12405 < v12406 {
		v12343 = v12405
		v12345 = v12397
		v12347 = v12402
		v12350 = v12392
		v12353 = v12389
		goto L1763
	} else {
		goto L1769
	}
L1769:
	;
	goto L1764
L1770:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12453))) = int32(370)
	v12457 = *(*int32)(unsafe.Add(mBase, uint32(v12239)+44))
	v12459 = v12421 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12453)+82)) = uint8(v12459)
	*(*uint16)(unsafe.Add(mBase, uint32(v12453)+80)) = uint16(v12425)
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+76)) = v12427
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+72)) = v12416
	v12464 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+56)) = v12464
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+52)) = v12239
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+48)) = v12464
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+44)) = v12457
	v12470 = *(*int32)(unsafe.Add(mBase, uint32(v12239)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+4)) = v12470
	v12472 = *(*float64)(unsafe.Add(mBase, uint32(v12239)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12453)+8)) = v12472
	v12474 = *(*float64)(unsafe.Add(mBase, uint32(v12239)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12453)+16)) = v12474
	v12476 = *(*float64)(unsafe.Add(mBase, uint32(v12239)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12453)+24)) = v12476
	v12478 = *(*int32)(unsafe.Add(mBase, uint32(v12239)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v12453)+36)) = uint8(v12464)
	*(*int32)(unsafe.Add(mBase, uint32(v12453)+32)) = v12478
	v12482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12239)+37)))
	*(*float64)(unsafe.Add(mBase, uint32(v12453)+8)) = v12474
	*(*uint8)(unsafe.Add(mBase, uint32(v12453)+37)) = uint8(v12482)
	v12485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v12485 != 0 {
		goto L1771
	} else {
		goto L1772
	}
L1771:
	;
	v12486 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12453)+36)) = uint8(v12486)
	v12488 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12453)+88)) = v12488
	goto L1773
L1772:
	;
	goto L1773
L1773:
	;
	v12490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v12491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v12492 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12493 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12495 = F_palloc0(m, int32(104))
	mBase = m.M
	v12496 = m.ExcPending
	if v12496 != 0 {
		goto L1
	} else {
		goto L1774
	}
L1774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+100)) = v12414
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+96)) = v12419
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+92)) = v12422
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+88)) = v12284
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+56)) = v12453
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+52)) = v12235
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+48)) = v12493
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+44)) = v12187
	*(*int32)(unsafe.Add(mBase, uint32(v12495))) = int32(359)
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+80)) = v12492
	*(*uint8)(unsafe.Add(mBase, uint32(v12495)+76)) = uint8(v12491)
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+72)) = v12490
	v12510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+4)) = v12510
	v12512 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v12495)+8)) = v12512
	v12514 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v12495)+16)) = v12514
	v12516 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12495)+24)) = v12516
	v12518 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12519 = *(*int32)(unsafe.Add(mBase, uint32(v12518)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12495)+32)) = v12519
	v12521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12495)+36)) = uint8(v12521)
	v12523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12495)+37)) = uint8(v12523)
	v12528 = v12495
	goto L4
L1775:
	;
	v12572 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12573 = F_order_qual_clauses(m, l0, v12572)
	mBase = m.M
	v12574 = m.ExcPending
	if v12574 != 0 {
		goto L1
	} else {
		goto L1776
	}
L1776:
	;
	v12576 = F_extract_actual_clauses(m, v12573, int32(1))
	mBase = m.M
	v12577 = m.ExcPending
	if v12577 != 0 {
		goto L1
	} else {
		goto L1777
	}
L1777:
	;
	if v12576 == int32(0) {
		v12585 = v12528
		v12589 = v47
		goto L3
	} else {
		goto L1778
	}
L1778:
	;
	v12580 = F_create_gating_plan(m, l0, l1, v12528, v12576)
	mBase = m.M
	v12581 = m.ExcPending
	if v12581 != 0 {
		goto L1
	} else {
		goto L1779
	}
L1779:
	;
	v12585 = v12580
	v12589 = v47
	goto L3
}
