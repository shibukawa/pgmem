package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bt_index_check_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 float32
	_ = v239
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var __phi470 int32
	_ = __phi470
	var v471 int32
	_ = v471
	var __phi471 int32
	_ = __phi471
	var v474 int32
	_ = v474
	var __phi474 int32
	_ = __phi474
	var v475 int32
	_ = v475
	var __phi475 int32
	_ = __phi475
	var v483 int32
	_ = v483
	var __phi483 int32
	_ = __phi483
	var v484 int32
	_ = v484
	var __phi484 int32
	_ = __phi484
	var v488 int32
	_ = v488
	var __phi488 int32
	_ = __phi488
	var v490 int32
	_ = v490
	var __phi490 int32
	_ = __phi490
	var v493 int32
	_ = v493
	var __phi493 int32
	_ = __phi493
	var v494 int32
	_ = v494
	var __phi494 int32
	_ = __phi494
	var v495 int32
	_ = v495
	var __phi495 int32
	_ = __phi495
	var v496 int32
	_ = v496
	var __phi496 int32
	_ = __phi496
	var v497 int32
	_ = v497
	var __phi497 int32
	_ = __phi497
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int64
	_ = v511
	var v512 int64
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
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
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v871 int64
	_ = v871
	var v874 int64
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
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
	var v924 int32
	_ = v924
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int64
	_ = v1259
	var v1262 int64
	_ = v1262
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1422 int32
	_ = v1422
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
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
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1608 int32
	_ = v1608
	var v1609 int64
	_ = v1609
	var v1614 int64
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1700 int32
	_ = v1700
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1824 int32
	_ = v1824
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1962 int32
	_ = v1962
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2050 int32
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int64
	_ = v2101
	var v2104 int64
	_ = v2104
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2150 int32
	_ = v2150
	var v2151 int64
	_ = v2151
	var v2156 int64
	_ = v2156
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
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
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int64
	_ = v2225
	var v2230 int64
	_ = v2230
	var v2236 int32
	_ = v2236
	var v2242 int32
	_ = v2242
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2310 int32
	_ = v2310
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2353 int64
	_ = v2353
	var v2357 int64
	_ = v2357
	var v2363 int32
	_ = v2363
	var v2373 int32
	_ = v2373
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2393 int32
	_ = v2393
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2413 int32
	_ = v2413
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2494 int64
	_ = v2494
	var v2499 int64
	_ = v2499
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2517 int32
	_ = v2517
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2552 int32
	_ = v2552
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2572 int32
	_ = v2572
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
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2594 int32
	_ = v2594
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2633 int32
	_ = v2633
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2700 int32
	_ = v2700
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2737 int32
	_ = v2737
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2788 int32
	_ = v2788
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2817 int32
	_ = v2817
	var v2855 int32
	_ = v2855
	var v2866 int32
	_ = v2866
	var v2890 int32
	_ = v2890
	var v2897 int32
	_ = v2897
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2912 int64
	_ = v2912
	var v2916 int64
	_ = v2916
	var v2922 int32
	_ = v2922
	var v2927 int32
	_ = v2927
	var v2931 int32
	_ = v2931
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int64
	_ = v2946
	var v2951 int64
	_ = v2951
	var v2957 int32
	_ = v2957
	var v2962 int32
	_ = v2962
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2990 int32
	_ = v2990
	var v2995 int32
	_ = v2995
	var v3006 int32
	_ = v3006
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3040 int32
	_ = v3040
	var v3054 int32
	_ = v3054
	var v3064 int32
	_ = v3064
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3098 int32
	_ = v3098
	var v3099 int32
	_ = v3099
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3112 int32
	_ = v3112
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3117 int32
	_ = v3117
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3141 int32
	_ = v3141
	var v3146 int32
	_ = v3146
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3164 int32
	_ = v3164
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3185 int32
	_ = v3185
	var v3189 int32
	_ = v3189
	var v3194 int32
	_ = v3194
	var v3198 int32
	_ = v3198
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3210 int32
	_ = v3210
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3258 int32
	_ = v3258
	var v3263 int32
	_ = v3263
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3268 int32
	_ = v3268
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3275 float64
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3281 int64
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3283 int64
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3287 int64
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3296 int32
	_ = v3296
	var v3302 int32
	_ = v3302
	var v3304 int32
	_ = v3304
	var v3331 int64
	_ = v3331
	var v3335 int32
	_ = v3335
	var v3338 int64
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3342 int64
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3346 int64
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3350 int64
	_ = v3350
	var v3354 int64
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3389 int64
	_ = v3389
	var v3396 int32
	_ = v3396
	var v3398 int32
	_ = v3398
	var v3425 int64
	_ = v3425
	var v3429 int32
	_ = v3429
	var v3432 int64
	_ = v3432
	var v3433 int64
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3437 int32
	_ = v3437
	var v3443 int64
	_ = v3443
	var v3452 int32
	_ = v3452
	var v3456 int32
	_ = v3456
	var v3458 int32
	_ = v3458
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3470 int64
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3474 int32
	_ = v3474
	var v3477 int32
	_ = v3477
	var v3480 int32
	_ = v3480
	var v3486 int64
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3490 int32
	_ = v3490
	var v3492 int32
	_ = v3492
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3500 int64
	_ = v3500
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3507 int32
	_ = v3507
	var v3510 int64
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3516 int64
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3528 int64
	_ = v3528
	var v3532 int32
	_ = v3532
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3542 int64
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3548 int64
	_ = v3548
	var v3549 int64
	_ = v3549
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3557 int32
	_ = v3557
	var v3561 int64
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3566 int32
	_ = v3566
	var v3570 int64
	_ = v3570
	var v3571 int32
	_ = v3571
	var v3574 int64
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3578 int64
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3582 int64
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3586 int64
	_ = v3586
	var v3590 int64
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3594 int32
	_ = v3594
	var v3601 int64
	_ = v3601
	var v3602 int64
	_ = v3602
	var v3632 int64
	_ = v3632
	var v3633 int64
	_ = v3633
	var v3648 int32
	_ = v3648
	var v3653 int32
	_ = v3653
	var v3687 int32
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3723 int32
	_ = v3723
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3728 int32
	_ = v3728
	var v3740 int32
	_ = v3740
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3755 int64
	_ = v3755
	var v3763 int64
	_ = v3763
	var v3769 int32
	_ = v3769
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3780 int32
	_ = v3780
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3801 int32
	_ = v3801
	var v3802 int32
	_ = v3802
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3819 int32
	_ = v3819
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3831 int64
	_ = v3831
	var v3836 int64
	_ = v3836
	var v3842 int32
	_ = v3842
	var v3848 int32
	_ = v3848
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3859 int32
	_ = v3859
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3864 int32
	_ = v3864
	var v3872 int32
	_ = v3872
	var v3873 int32
	_ = v3873
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3894 int64
	_ = v3894
	var v3899 int64
	_ = v3899
	var v3905 int32
	_ = v3905
	var v3911 int32
	_ = v3911
	var v3916 int32
	_ = v3916
	var v3923 int32
	_ = v3923
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3935 int32
	_ = v3935
	var v3940 int32
	_ = v3940
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3948 int32
	_ = v3948
	var v3956 int32
	_ = v3956
	var v3961 int32
	_ = v3961
	v4 = l3
	v34 = m.G0
	v36 = v34 + int32(-64)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L5
	} else {
		goto L866
	}
L2:
	;
	v64 = v38
	goto L4
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
	v46 = F_smgropen(m, v34+int32(-16), v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v66 = F_smgrexists(m, v64, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L11
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v50 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = v62
	goto L4
L8:
	;
	v58 = v50
	goto L10
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	v58 = v56
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+72)) = v58 + int32(1)
	goto L7
L11:
	;
	if v66 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F__bt_metaversion(m, l0, v34+int32(-1), v34+int32(-2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L5
	} else {
		goto L862
	}
L15:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+62)))
	if v74 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+63)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+3)))
	v191 = m.G0
	v193 = v191 - int32(1168)
	m.G0 = v193
	v197 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L33
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+63)))
	if v77 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v81 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	if v81 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+10)))
	if v84 <= int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v102 = int32(0)
	goto L22
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v87+v102<<(uint(int32(2))%32))))
	if v125 != int32(1982) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L28
	}
L24:
	;
	v129 = v102 + int32(1)
	if v84 != v129 {
		v102 = v129
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L16
L28:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v138 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_0), v36)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_1), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(345), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	if v197 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+1104)) = v199 + int32(4)
	if v4 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v218 = F_palloc0(m, int32(72))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L45
	}
L37:
	;
	v205 = int32(_a_F_bt_index_check_callback_4)
	goto L39
L38:
	;
	v205 = int32(_a_F_bt_index_check_callback_5)
	goto L39
L39:
	;
	F_errmsg_internal(m, v205, v193+int32(1104))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if v4 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v213 = int32(393)
	goto L43
L42:
	;
	v213 = int32(390)
	goto L43
L43:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), v213, int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+28)) = v220
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+12)) = uint8(v190)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+11)) = uint8(v189)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)) = uint8(v188)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = l0
	if v188 == v220 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	m.G0 = v416 - int32(-64)
	return
L47:
	;
	v3854 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3859 = m.ExcPending
	if v3859 != 0 {
		goto L5
	} else {
		goto L850
	}
L48:
	;
	v3775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v3775&int32(32) == int32(0) {
		v3791 = v1082
		goto L837
	} else {
		goto L838
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
		goto L5
	} else {
		goto L831
	}
L50:
	;
	v3216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+10)))
	if v3216 == int32(1) {
		goto L764
	} else {
		goto L765
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L5
	} else {
		goto L760
	}
L52:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+12)))
	if v323 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v232 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v236 = base.I64_extend_i32_u(v232) * int64(452)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+48))
	v239 = *(*float32)(unsafe.Add(mBase, uint32(v238)+100))
	if base.F32_lt(base.F32_abs(v239), float32(9.223372e+18)) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v246 = int32(_a_F_bt_index_check_callback_7)
	v247 = int32(_a_F_bt_index_check_callback_8)
	v248 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0]))
	v250 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1]))
	v251 = v248 ^ v250
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[1])) = base.I64_rotl(v251, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[0])) = v251<<(uint(int64(16))%64) ^ base.I64_rotl(v248, int64(24)) ^ v251
	if v245 < v236 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v243 = base.I64_trunc_f32_s(v239)
	v245 = v243
	goto L55
L57:
	;
	goto L58
L58:
	;
	v245 = int64(-9223372036854775807 - 1)
	goto L55
L59:
	;
	v270 = v236
	goto L61
L60:
	;
	v270 = v245
	goto L61
L61:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[2]))
	v273 = F_bloom_create(m, v270, v272, base.I64_rotl(v248*int64(5), int64(7))*int64(9))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v218)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+60)) = v273
	v278 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v280 = F_RegisterSnapshot(m, v278)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+28)) = v280
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[3]))
	if v284 < int32(2) {
		goto L52
	} else {
		goto L65
	}
L65:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+19)))
	if v288 != int32(1) {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292)+20)))
	v294 = int32(768)
	if v293&v294 != v294 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v300 = v298
	goto L69
L68:
	;
	v300 = int32(2)
	goto L69
L69:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v301))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v300)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v313 == int32(0) {
		goto L51
	} else {
		goto L74
	}
L71:
	;
	v313 = base.B2i32(base.Ui32(v300) < base.Ui32(v301))
	goto L70
L72:
	;
	goto L73
L73:
	;
	v313 = int32(base.Ui32(v300-v301) >> (uint(int32(31)) % 32))
	goto L70
L74:
	;
	goto L52
L75:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+11)))
	if v340 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v327 = F_BuildIndexInfo(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+24)) = v327
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+116)))
	if v330 != int32(1) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	if v333 != 0 {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v334 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v336 = F_RegisterSnapshot(m, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+28)) = v336
	goto L75
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3173 = m.ExcPending
	if v3173 != 0 {
		goto L5
	} else {
		goto L755
	}
L83:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
	if v343 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v352 = F_AllocSetContextCreateInternal(m, v347, int32(_a_F_bt_index_check_callback_9), int32(0), int32(_a_F_bt_index_check_callback_10), int32(_a_F_bt_index_check_callback_11))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+16)) = v352
	v356 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+20)) = v356
	v360 = F_palloc_btree_page(m, v218, int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L90
	}
L89:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v360)+32))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v360)+36))
	v403 = v398
	v406 = v193
	v407 = v218
	v415 = int32(-1)
	v416 = v36
	v420 = l0
	v422 = v399
	v426 = int32(1)
	v427 = l1
	goto L99
L90:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v360)+40))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v360)+32))
	if v362 == v363 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v367 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	if v367 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+1056)) = v374 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_12), v193+int32(1056))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v360)+40))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v360)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v193)+1048)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v193)+1040)) = v383
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_13), v193+int32(1040))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(512), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	goto L89
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L5
	} else {
		goto L751
	}
L99:
	;
	if v403 == int32(0) {
		goto L50
	} else {
		goto L101
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L5
	} else {
		goto L747
	}
L101:
	;
	v437 = int32(_a_F_bt_index_check_callback_14)
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4]))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v440
	v444 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	if v444 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+1024)) = v422
	if v422 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v463 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v407)+56)) = uint8(v463)
	v466 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+52)) = v466
	__phi470 = v463
	__phi471 = v403
	__phi474 = v406
	__phi475 = v407
	__phi483 = v466
	__phi484 = v416
	__phi488 = v420
	__phi490 = v422
	__phi493 = v466
	__phi494 = v426
	__phi495 = v427
	__phi496 = v438
	__phi497 = v415
	v470 = __phi470
	v471 = __phi471
	v474 = __phi474
	v475 = __phi475
	v483 = __phi483
	v484 = __phi484
	v488 = __phi488
	v490 = __phi490
	v493 = __phi493
	v494 = __phi494
	v495 = __phi495
	v496 = __phi496
	v497 = __phi497
	goto L114
L106:
	;
	v450 = int32(_a_F_bt_index_check_callback_15)
	goto L108
L107:
	;
	v450 = int32(_a_F_bt_index_check_callback_16)
	goto L108
L108:
	;
	if v426 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v451 = int32(_a_F_bt_index_check_callback_17)
	goto L111
L110:
	;
	v451 = v450
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+1028)) = v451
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_18), v406+int32(1024))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(645), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	goto L105
L114:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v504 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v475)+48))
	if v3115 != 0 {
		goto L742
	} else {
		goto L743
	}
L116:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L5
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+36)) = v471
	v508 = F_palloc_btree_page(m, v475, v471)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+32)) = v508
	v511 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v508)+4)))
	v512 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v508))))
	*(*int64)(unsafe.Add(mBase, uint32(v475)+40)) = v511 | v512<<(uint(int64(32))%64)
	v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+16)))
	v518 = v508 + v517
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v518)+12)))
	if v519&int32(20) != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	if v471 == v470 {
		goto L98
	} else {
		goto L725
	}
L122:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	if v866 == v490 {
		goto L215
	} else {
		goto L216
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L5
	} else {
		goto L210
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L206
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L5
	} else {
		goto L202
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L5
	} else {
		goto L197
	}
L127:
	;
	if v519&int32(4) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	if v493 != int32(-1) {
		v597 = v483
		v598 = v493
		goto L140
	} else {
		goto L141
	}
L130:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v524&int32(1) != 0 {
		goto L126
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v527 == int32(0) {
		goto L125
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v532 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	if v532 == int32(0) {
		v3054 = v483
		v3064 = v493
		goto L121
	} else {
		goto L136
	}
L136:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1008)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1012)) = v540 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_20), v474+int32(1008))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(692), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	v3054 = v483
	v3064 = v493
	goto L121
L140:
	;
	if v470 == int32(0) {
		goto L122
	} else {
		goto L155
	}
L141:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v557 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v560 = F_bt_leftmost_ignoring_half_dead(m, v475, v471, v518)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	v570 = v519
	goto L144
L144:
	;
	if v570&int32(1) != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	if v560 == int32(0) {
		goto L124
	} else {
		goto L146
	}
L146:
	;
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v518)+12)))
	if base.B2i32(v564&int32(130) == int32(0))&v494 != 0 {
		goto L123
	} else {
		goto L147
	}
L147:
	;
	v570 = v564
	goto L144
L148:
	;
	v597 = int32(-1)
	v598 = int32(0)
	goto L140
L149:
	;
	goto L150
L150:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v579 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v580 = int32(2)
	goto L153
L152:
	;
	v580 = int32(1)
	goto L153
L153:
	;
	v581 = F_PageGetItemIdCareful_2(m, v475, v575, v576, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v587 = v583 + v584&int32(_a_F_bt_index_check_callback_21)
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587))))
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587)+2)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v597 = v593 - int32(1)
	v598 = v588<<(uint(int32(16))%32) | v591
	goto L140
L155:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v601 == v470 {
		goto L122
	} else {
		goto L156
	}
L156:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v603 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v607 = int32(0)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v475)+20))
	v610 = F_ReadBufferExtended(m, v606, v607, v470, v607, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L5
	} else {
		goto L160
	}
L158:
	;
	v729 = v601
	goto L159
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L5
	} else {
		goto L192
	}
L160:
	;
	F_LockBuffer(m, v610, int32(1))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	F__bt_checkpage(m, v615, v610)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	if v610 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v635)+16)))
	v637 = v636 + v635
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+12)))
	if v638&int32(4) != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v621+(v610^int32(-1))<<(uint(int32(2))%32))))
	v635 = v627
	goto L163
L165:
	;
	goto L166
L166:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v635 = v629 + v610<<(uint(int32(13))%32) + int32(-8192)
	goto L163
L167:
	;
	F_UnlockReleaseBuffer(m, v610)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L5
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	if v644 == v470 {
		v688 = int32(-1)
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L122
L171:
	;
	F_UnlockReleaseBuffer(m, v610)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L5
	} else {
		goto L182
	}
L172:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v647 = int32(0)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v475)+20))
	v650 = F_ReadBufferExtended(m, v646, v647, v644, v647, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	F_LockBuffer(m, v650, int32(1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	F__bt_checkpage(m, v655, v650)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	if v650 < int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	F_UnlockReleaseBuffer(m, v650)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L5
	} else {
		goto L181
	}
L177:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661+(v650^int32(-1))<<(uint(int32(2))%32))))
	v668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v667)+16)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v667+v668)))
	v685 = v670
	goto L176
L178:
	;
	goto L179
L179:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v675 = v672 + v650<<(uint(int32(13))%32)
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v675-int32(_a_F_bt_index_check_callback_22)))))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v675+v678)+uint32(_c_F_bt_index_check_callback[8])))
	if v650 == int32(0) {
		v688 = v682
		goto L171
	} else {
		goto L180
	}
L180:
	;
	v685 = v682
	goto L176
L181:
	;
	v688 = v685
	goto L171
L182:
	;
	if v688 == v470 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v695 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L5
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+36)) = v644
	v729 = v688
	goto L159
L186:
	;
	if v695 == int32(0) {
		goto L122
	} else {
		goto L187
	}
L187:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+928)) = v703 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_23), v474+int32(928))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+920)) = v712
	*(*int32)(unsafe.Add(mBase, uint32(v474)+916)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v474)+912)) = v470
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_24), v474+int32(912))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1177), int32(_a_F_bt_index_check_callback_25))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	goto L122
L192:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+80)) = v739 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_26), v474+int32(80))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+72)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v474)+68)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v474)+64)) = v748
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_27), v474-int32(-64))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1197), int32(_a_F_bt_index_check_callback_25))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+976)) = v770 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_28), v474+int32(976))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+968)) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v474)+964)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v474)+960)) = v471
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_27), v474+int32(960))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(681), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v800)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+992)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v474)+996)) = v801 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_29), v474+int32(992))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(687), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+944)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v474)+948)) = v824 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_30), v474+int32(944))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(710), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L211
	}
L211:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+48)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v474)+52)) = v847 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_31), v474+int32(48))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(716), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	v3029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3006)+12)))
	if v3029&int32(1) != 0 {
		v3054 = v597
		v3064 = v598
		goto L121
	} else {
		goto L721
	}
L215:
	;
	v871 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v474+int32(1128)))) = v871
	v874 = *(*int64)(unsafe.Add(mBase, _c_F_bt_index_check_callback[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v474)+1120)) = v874
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v877) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L217
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L5
	} else {
		goto L716
	}
L218:
	;
	v885 = int32(base.Ui32(v877+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L220
L219:
	;
	v885 = int32(0)
	goto L220
L220:
	;
	v886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876)+16)))
	v887 = v876 + v886
	v890 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	if v890 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v892 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v887)+12)))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+872)) = v893
	*(*int32)(unsafe.Add(mBase, uint32(v474)+864)) = v885 & int32(_a_F_bt_index_check_callback_33)
	if v892&int32(1) != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	if v915 != 0 {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v902 = int32(_a_F_bt_index_check_callback_34)
	goto L227
L226:
	;
	v902 = int32(_a_F_bt_index_check_callback_35)
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+868)) = v902
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_36), v474+int32(864))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L5
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1252), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	goto L224
L230:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v919 = F_PageGetItemIdCareful_2(m, v475, v916, v917, int32(1))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L5
	} else {
		goto L233
	}
L231:
	;
	v1029 = int32(1)
	goto L232
L232:
	;
	v1031 = v885 & int32(_a_F_bt_index_check_callback_33)
	if base.Ui32(v1031) < base.Ui32(v1029) {
		v3006 = v887
		goto L214
	} else {
		goto L275
	}
L233:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v924 = int32(1)
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v923)+16)))
	v933 = v923 + v932
	v934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v933)+12)))
	if v934&int32(20) != 0 {
		v1011 = v924
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v1020 == int32(0) {
		goto L47
	} else {
		goto L271
	}
L235:
	;
	v1020 = v1011
	goto L234
L236:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v921)+192))
	v938 = int32(*(*int16)(unsafe.Add(mBase, uint32(v937)+10)))
	v939 = int32(*(*int16)(unsafe.Add(mBase, uint32(v937)+8)))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(int32(4)+v923)+20))
	v946 = v923 + v943&int32(_a_F_bt_index_check_callback_21)
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+6)))
	v949 = v947 & int32(_a_F_bt_index_check_callback_10)
	if v949 == int32(0) {
		v965 = v939
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	if v968 != 0 {
		goto L245
	} else {
		goto L246
	}
L238:
	;
	v952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	if v952&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v955 = int32(0)
	if v922 == v955 {
		v1011 = v955
		goto L235
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v965 = v952 & int32(4095)
	goto L237
L242:
	;
	if v952&int32(_a_F_bt_index_check_callback_38) != 0 {
		v1011 = v955
		goto L235
	} else {
		goto L243
	}
L243:
	;
	if v939 != v938 {
		v1011 = v955
		goto L235
	} else {
		goto L244
	}
L244:
	;
	v965 = v939
	goto L237
L245:
	;
	v969 = int32(2)
	goto L247
L246:
	;
	v969 = int32(1)
	goto L247
L247:
	;
	if v934&int32(1) != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v994 = int32(0)
	if v949 == v994 {
		v1011 = v994
		goto L235
	} else {
		goto L265
	}
L249:
	;
	if base.Ui32(v969) <= base.Ui32(v924) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	if v924 == v969 {
		goto L259
	} else {
		goto L260
	}
L252:
	;
	if v949 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	if v922 != 0 {
		goto L248
	} else {
		goto L258
	}
L255:
	;
	v1020 = base.B2i32(v939 == v965)
	goto L234
L256:
	;
	goto L257
L257:
	;
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	v1020 = int32(base.Ui32(v976)>>(uint(int32(13))%32)) & base.B2i32(v939 == v965)
	goto L234
L258:
	;
	v1020 = base.B2i32(v965 == v938)
	goto L234
L259:
	;
	v987 = base.B2i32(v965 == int32(0)) | (v922 ^ int32(1))
	if v922 != 0 {
		v1011 = v987
		goto L235
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	if v922 != 0 {
		goto L248
	} else {
		goto L264
	}
L262:
	;
	if v965 == int32(0) {
		v1011 = v987
		goto L235
	} else {
		goto L263
	}
L263:
	;
	v990 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+4)))
	v1020 = base.B2i32(v990 == int32(1))
	goto L234
L264:
	;
	v1020 = base.B2i32(v965 == v938)
	goto L234
L265:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+5)))
	if v997&int32(32) != 0 {
		v1011 = v994
		goto L235
	} else {
		goto L266
	}
L266:
	;
	v1000 = F_BTreeTupleGetHeapTID(m, v946)
	mBase = m.M
	if v965 != v938 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1003 = v1000
	goto L269
L268:
	;
	v1003 = int32(0)
	goto L269
L269:
	;
	if v1003 != 0 {
		v1011 = v994
		goto L235
	} else {
		goto L270
	}
L270:
	;
	v1011 = base.B2i32(v965 <= v938) & base.B2i32(int32(0) < v965)
	goto L235
L271:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	if v1025 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1026 = int32(2)
	goto L274
L273:
	;
	v1026 = int32(1)
	goto L274
L274:
	;
	v1029 = v1026
	goto L232
L275:
	;
	v1043 = v887
	v1049 = v1029
	goto L278
L276:
	;
	F_pfree(m, v2600)
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L5
	} else {
		goto L715
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L5
	} else {
		goto L710
	}
L278:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v1067 != 0 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L5
	} else {
		goto L705
	}
L280:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v1073 = v1049 & int32(_a_F_bt_index_check_callback_33)
	v1074 = F_PageGetItemIdCareful_2(m, v475, v1070, v1071, v1073)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L5
	} else {
		goto L294
	}
L283:
	;
	goto L282
L284:
	;
	goto L279
L285:
	;
	v2890 = v1049 + int32(1)
	if base.Ui32(v2890&int32(_a_F_bt_index_check_callback_33)) <= base.Ui32(v1031) {
		v1043 = v2866
		v1049 = v2890
		goto L278
	} else {
		goto L704
	}
L286:
	;
	v2656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2633)+12)))
	if v2656&int32(1) != 0 {
		v2866 = v2633
		goto L285
	} else {
		goto L641
	}
L287:
	;
	v2544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+12)))
	if v2544 != int32(1) {
		v2633 = v1043
		goto L286
	} else {
		goto L612
	}
L288:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v2440)))
	v2449 = F__bt_mkscankey(m, v2444, v1996+v2445&int32(_a_F_bt_index_check_callback_21))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L5
	} else {
		goto L599
	}
L289:
	;
	v2438 = int32(0)
	v2517 = v2438
	v2519 = v2413
	v2522 = v2438
	goto L287
L290:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L5
	} else {
		goto L593
	}
L291:
	;
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+484)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+480)) = v2248
	v2254 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(480))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L5
	} else {
		goto L572
	}
L292:
	;
	v2168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v2168&int32(32) == int32(0) {
		v2184 = v1082
		goto L559
	} else {
		goto L560
	}
L293:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+612)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+608)) = v2126
	v2132 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(608))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L5
	} else {
		goto L552
	}
L294:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1074)))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v1082 = v1079 + v1076&int32(_a_F_bt_index_check_callback_21)
	v1083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+6)))
	v1085 = v1083 & int32(_a_F_bt_index_check_callback_40)
	if int32(base.Ui32(v1076)>>(uint(int32(17))%32)) == v1085 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1079)+16)))
	v1097 = v1079 + v1096
	v1098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1097)+12)))
	if v1098&int32(20) != 0 {
		v1175 = int32(1)
		goto L299
	} else {
		goto L300
	}
L296:
	;
	goto L297
L297:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L5
	} else {
		goto L546
	}
L298:
	;
	if v1184 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L299:
	;
	v1184 = v1175
	goto L298
L300:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+192))
	v1102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1101)+10)))
	v1103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1101)+8)))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1073<<(uint(int32(2))%32)+v1079)+20))
	v1110 = v1079 + v1107&int32(_a_F_bt_index_check_callback_21)
	v1111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+6)))
	v1113 = v1111 & int32(_a_F_bt_index_check_callback_10)
	if v1113 == int32(0) {
		v1129 = v1103
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1132 != 0 {
		goto L309
	} else {
		goto L310
	}
L302:
	;
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+4)))
	if v1116&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1119 = int32(0)
	if v1088 == v1119 {
		v1175 = v1119
		goto L299
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1129 = v1116 & int32(4095)
	goto L301
L306:
	;
	if v1116&int32(_a_F_bt_index_check_callback_38) != 0 {
		v1175 = v1119
		goto L299
	} else {
		goto L307
	}
L307:
	;
	if v1103 != v1102 {
		v1175 = v1119
		goto L299
	} else {
		goto L308
	}
L308:
	;
	v1129 = v1103
	goto L301
L309:
	;
	v1133 = int32(2)
	goto L311
L310:
	;
	v1133 = int32(1)
	goto L311
L311:
	;
	if v1098&int32(1) != 0 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1158 = int32(0)
	if v1113 == v1158 {
		v1175 = v1158
		goto L299
	} else {
		goto L329
	}
L313:
	;
	if base.Ui32(v1133) <= base.Ui32(v1073) {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L315
L315:
	;
	if v1073 == v1133 {
		goto L323
	} else {
		goto L324
	}
L316:
	;
	if v1113 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	goto L318
L318:
	;
	if v1088 != 0 {
		goto L312
	} else {
		goto L322
	}
L319:
	;
	v1184 = base.B2i32(v1103 == v1129)
	goto L298
L320:
	;
	goto L321
L321:
	;
	v1140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+4)))
	v1184 = int32(base.Ui32(v1140)>>(uint(int32(13))%32)) & base.B2i32(v1103 == v1129)
	goto L298
L322:
	;
	v1184 = base.B2i32(v1129 == v1102)
	goto L298
L323:
	;
	v1151 = base.B2i32(v1129 == int32(0)) | (v1088 ^ int32(1))
	if v1088 != 0 {
		v1175 = v1151
		goto L299
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	if v1088 != 0 {
		goto L312
	} else {
		goto L328
	}
L326:
	;
	if v1129 == int32(0) {
		v1175 = v1151
		goto L299
	} else {
		goto L327
	}
L327:
	;
	v1154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+4)))
	v1184 = base.B2i32(v1154 == int32(1))
	goto L298
L328:
	;
	v1184 = base.B2i32(v1129 == v1102)
	goto L298
L329:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+5)))
	if v1161&int32(32) != 0 {
		v1175 = v1158
		goto L299
	} else {
		goto L330
	}
L330:
	;
	v1164 = F_BTreeTupleGetHeapTID(m, v1110)
	mBase = m.M
	if v1129 != v1102 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1167 = v1164
	goto L333
L332:
	;
	v1167 = int32(0)
	goto L333
L333:
	;
	if v1167 != 0 {
		v1175 = v1158
		goto L299
	} else {
		goto L334
	}
L334:
	;
	v1175 = base.B2i32(v1129 <= v1102) & base.B2i32(int32(0) < v1129)
	goto L299
L335:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+756)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+752)) = v1187
	v1193 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(752))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L5
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v1283&int32(1) == int32(0) {
		goto L361
	} else {
		goto L362
	}
L338:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1195&int32(32) == int32(0) {
		v1211 = v1082
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211)+2)))
	v1213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211))))
	v1214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+740)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v474)+736)) = v1212 | v1213<<(uint(int32(16))%32)
	v1223 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(736))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L342
	}
L340:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v1200&int32(32) == int32(0) {
		v1211 = v1082
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1211 = v1082 + (v1205 | v1206<<(uint(int32(16))%32))
	goto L339
L342:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L343
	}
L343:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L5
	} else {
		goto L344
	}
L344:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+720)) = v1233 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_41), v474+int32(720))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1242&int32(32) == int32(0) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1043)+12)))
	v1259 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+708)) = uint32(v1259)
	v1262 = int64(base.Ui64(v1259) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+704)) = uint32(v1262)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+700)) = v1223
	*(*int32)(unsafe.Add(mBase, uint32(v474)+692)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v474)+688)) = v1193
	if v1258&int32(1) != 0 {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+192))
	v1255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1254)+8)))
	v1257 = v1255
	goto L346
L348:
	;
	v1247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1247&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1257 = v1247 & int32(4095)
	goto L346
L350:
	;
	v1271 = int32(_a_F_bt_index_check_callback_42)
	goto L352
L351:
	;
	v1271 = int32(_a_F_bt_index_check_callback_43)
	goto L352
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+696)) = v1271
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_44), v474+int32(688))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1352), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	if base.Ui32(v1663) < base.Ui32(v1085) {
		goto L292
	} else {
		goto L428
	}
L356:
	;
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v1627&int32(1) != 0 {
		v1663 = int32(2704)
		goto L355
	} else {
		goto L418
	}
L357:
	;
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1554&int32(32) == int32(0) {
		v1570 = v1082
		goto L408
	} else {
		goto L409
	}
L358:
	;
	F__bt_freestack(m, v1311)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L5
	} else {
		goto L405
	}
L359:
	;
	F__bt_relbuf(m, v1337)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L5
	} else {
		goto L404
	}
L360:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1388&int32(32) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L361:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	if v1290 != 0 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	goto L363
L363:
	;
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+11)))
	if v1300 != int32(1) {
		goto L360
	} else {
		goto L370
	}
L364:
	;
	v1291 = int32(2)
	goto L366
L365:
	;
	v1291 = int32(1)
	goto L366
L366:
	;
	if v1291 != v1073 {
		goto L360
	} else {
		goto L367
	}
L367:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v1293 != int32(1) {
		v2866 = v1043
		goto L285
	} else {
		goto L368
	}
L368:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+8))
	F_bt_child_highkey_check(m, v475, v1073, int32(0), v1297)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L5
	} else {
		goto L369
	}
L369:
	;
	v2866 = v1043
	goto L285
L370:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1304 = F__bt_mkscankey(m, v1303, v1082)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L5
	} else {
		goto L371
	}
L371:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1311 = F__bt_search(m, v1306, int32(0), v1304, v474+int32(1164), int32(1))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L5
	} else {
		goto L372
	}
L372:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v474)+1164))
	if v1313 == int32(0) {
		goto L358
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1136)) = v1082
	v1317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+6)))
	v1318 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1160)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1144)) = v1304
	*(*uint8)(unsafe.Add(mBase, uint32(v474)+1152)) = uint8(v1318)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1148)) = v1313
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1140)) = (v1317&int32(_a_F_bt_index_check_callback_40) + int32(7)) & int32(_a_F_bt_index_check_callback_45)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1334 = F__bt_binsrch_insert(m, v1331, v474+int32(1136))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L5
	} else {
		goto L374
	}
L374:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v474)+1164))
	if v1337 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1355)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1356) {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[6]))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1341+(v1337^int32(-1))<<(uint(int32(2))%32))))
	v1355 = v1347
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[7]))
	v1355 = v1349 + v1337<<(uint(int32(13))%32) + int32(-8192)
	goto L375
L379:
	;
	v1364 = int32(base.Ui32(v1356+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L381
L380:
	;
	v1364 = int32(0)
	goto L381
L381:
	;
	if base.Ui32(v1364&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v1334) {
		goto L359
	} else {
		goto L382
	}
L382:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v474)+1160))
	if int32(0) < v1368 {
		goto L359
	} else {
		goto L383
	}
L383:
	;
	v1371 = F__bt_compare(m, v1336, v1304, v1355, v1334)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L5
	} else {
		goto L384
	}
L384:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v474)+1164))
	F__bt_relbuf(m, v1374)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L5
	} else {
		goto L385
	}
L385:
	;
	F__bt_freestack(m, v1311)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L5
	} else {
		goto L386
	}
L386:
	;
	F_pfree(m, v1304)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L5
	} else {
		goto L387
	}
L387:
	;
	if v1371 != 0 {
		goto L357
	} else {
		goto L388
	}
L388:
	;
	goto L360
L389:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1530 = F__bt_mkscankey(m, v1529, v1082)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L5
	} else {
		goto L402
	}
L390:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v1393&int32(32) == int32(0) {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1399 = v474 + int32(1140)
	v1400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1405 = v1082 + (v1400 | v1401<<(uint(int32(16))%32))
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1405)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1399))) = uint16(v1406)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1405)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1136)) = v1408
	v1411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1411&int32(4094) == int32(0) {
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v1422 = int32(1)
	goto L393
L393:
	;
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1451 = int32(16)
	v1457 = v1082 + (v1449 | v1450<<(uint(v1451)%32)) + v1422*int32(6)
	v1459 = v474 + int32(1136)
	v1463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1457)+2)))
	v1464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1457))))
	v1467 = v1463 | v1464<<(uint(v1451)%32)
	v1468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1459)+2)))
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1459))))
	v1472 = v1468 | v1469<<(uint(v1451)%32)
	if base.Ui32(v1467) < base.Ui32(v1472) {
		v1483 = int32(-1)
		goto L396
	} else {
		goto L397
	}
L394:
	;
	goto L389
L395:
	;
	if v1483 <= int32(0) {
		goto L293
	} else {
		goto L400
	}
L396:
	;
	goto L395
L397:
	;
	if base.Ui32(v1472) < base.Ui32(v1467) {
		v1483 = int32(1)
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1457)+4)))
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1459)+4)))
	if base.Ui32(v1477) < base.Ui32(v1478) {
		v1483 = int32(-1)
		goto L396
	} else {
		goto L399
	}
L399:
	;
	v1483 = base.B2i32(base.Ui32(v1478) < base.Ui32(v1477))
	goto L396
L400:
	;
	v1486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1457)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1399))) = uint16(v1486)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1457)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1136)) = v1488
	v1491 = v1422 + int32(1)
	v1492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if base.Ui32(v1491) < base.Ui32(v1492&int32(4095)) {
		v1422 = v1491
		goto L393
	} else {
		goto L401
	}
L401:
	;
	goto L394
L402:
	;
	v1532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1530)+4)) = uint8(v1532)
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530))))
	if v1534 == v1532 {
		goto L356
	} else {
		goto L403
	}
L403:
	;
	v1663 = int32(2712)
	goto L355
L404:
	;
	goto L358
L405:
	;
	F_pfree(m, v1304)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	goto L357
L407:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+676)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+672)) = v1571
	v1577 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(672))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L5
	} else {
		goto L411
	}
L408:
	;
	goto L407
L409:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v1559&int32(32) == int32(0) {
		v1570 = v1082
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v1564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1570 = v1082 + (v1564 | v1565<<(uint(int32(16))%32))
	goto L408
L411:
	;
	v1579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1570)+2)))
	v1580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1570))))
	v1581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1570)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+660)) = v1581
	*(*int32)(unsafe.Add(mBase, uint32(v474)+656)) = v1579 | v1580<<(uint(int32(16))%32)
	v1590 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(656))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L5
	} else {
		goto L412
	}
L412:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L5
	} else {
		goto L413
	}
L413:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L5
	} else {
		goto L414
	}
L414:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+640)) = v1600 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_46), v474+int32(640))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	v1609 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+636)) = uint32(v1609)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+628)) = v1590
	*(*int32)(unsafe.Add(mBase, uint32(v474)+624)) = v1577
	v1614 = int64(base.Ui64(v1609) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+632)) = uint32(v1614)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_47), v474+int32(624))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L5
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1399), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L5
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
	v1632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+6)))
	if v1632&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		v1659 = v1082
		goto L419
	} else {
		goto L420
	}
L419:
	;
	if v1659 != 0 {
		goto L425
	} else {
		goto L426
	}
L420:
	;
	v1637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1637&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1642 = int32(0)
	if v1637&int32(_a_F_bt_index_check_callback_38) == v1642 {
		v1659 = v1642
		goto L419
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v1652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1659 = v1082 + (v1652 | v1653<<(uint(int32(16))%32))
	goto L419
L424:
	;
	v1659 = v1082 + v1632&int32(_a_F_bt_index_check_callback_40) - int32(6)
	goto L419
L425:
	;
	v1660 = int32(2712)
	goto L427
L426:
	;
	v1660 = int32(2704)
	goto L427
L427:
	;
	v1663 = v1660
	goto L355
L428:
	;
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+10)))
	if v1665 != int32(1) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+8))
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
	if v1803 != int32(1) {
		goto L452
	} else {
		goto L453
	}
L430:
	;
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v1668&int32(1) == int32(0) {
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1074)))
	v1674 = int32(_a_F_bt_index_check_callback_48)
	if v1673&v1674 == v1674 {
		goto L429
	} else {
		goto L432
	}
L432:
	;
	v1678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1678&int32(32) == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1758 = F_bt_normalize_tuple(m, v475, v1082)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L5
	} else {
		goto L448
	}
L434:
	;
	v1683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1683&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	if v1683&int32(4095) == int32(0) {
		goto L429
	} else {
		goto L436
	}
L436:
	;
	v1700 = int32(0)
	goto L437
L437:
	;
	v1726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1736 = F__bt_form_posting(m, v1082, v1082+(v1726|v1727<<(uint(int32(16))%32))+v1700*int32(6), int32(1))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L5
	} else {
		goto L439
	}
L438:
	;
	goto L429
L439:
	;
	v1738 = F_bt_normalize_tuple(m, v475, v1736)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L5
	} else {
		goto L440
	}
L440:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v475)+60))
	v1741 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1738)+6)))
	F_bloom_add_element(m, v1740, v1738, v1741&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L5
	} else {
		goto L441
	}
L441:
	;
	if v1736 != v1738 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	F_pfree(m, v1738)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L5
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	F_pfree(m, v1736)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L5
	} else {
		goto L446
	}
L445:
	;
	goto L444
L446:
	;
	v1752 = v1700 + int32(1)
	v1753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if base.Ui32(v1752) < base.Ui32(v1753&int32(4095)) {
		v1700 = v1752
		goto L437
	} else {
		goto L447
	}
L447:
	;
	goto L438
L448:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v475)+60))
	v1761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1758)+6)))
	F_bloom_add_element(m, v1760, v1758, v1761&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L5
	} else {
		goto L449
	}
L449:
	;
	if v1082 == v1758 {
		goto L429
	} else {
		goto L450
	}
L450:
	;
	F_pfree(m, v1758)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L5
	} else {
		goto L451
	}
L451:
	;
	goto L429
L452:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	if v1831 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L453:
	;
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1806&int32(32) == int32(0) {
		goto L452
	} else {
		goto L454
	}
L454:
	;
	v1811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+4)))
	if v1811&int32(_a_F_bt_index_check_callback_10) == int32(0) {
		goto L452
	} else {
		goto L455
	}
L455:
	;
	v1816 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v1817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v1824 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v1530)+8)) = v1082 + (v1816 | v1817<<(uint(int32(16))%32)) + v1811&int32(4095)*v1824 - v1824
	goto L452
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1530)+8)) = v1802
	v1851 = v1049 + int32(1)
	v1853 = v1851 & int32(_a_F_bt_index_check_callback_33)
	v1854 = base.B2i32(base.Ui32(v1031) < base.Ui32(v1853))
	if v1854 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L457:
	;
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v1834&int32(1) != 0 {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v1840 = F__bt_compare(m, v1837, v1530, v1838, int32(1))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L5
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v1845 = F_invariant_l_offset(m, v475, v1530, int32(1))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L5
	} else {
		goto L463
	}
L461:
	;
	if v1840 <= int32(0) {
		goto L456
	} else {
		goto L462
	}
L462:
	;
	goto L48
L463:
	;
	if v1845 == int32(0) {
		goto L48
	} else {
		goto L464
	}
L464:
	;
	goto L456
L465:
	;
	v1857 = F_invariant_l_offset(m, v475, v1530, v1853)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L5
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v1861 = int32(0)
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+12)))
	if v1862 != int32(1) {
		v1946 = v1861
		goto L470
	} else {
		goto L471
	}
L468:
	;
	if v1857 == int32(0) {
		goto L291
	} else {
		goto L469
	}
L469:
	;
	goto L467
L470:
	;
	if v1073 != v1031 {
		v2633 = v1043
		goto L286
	} else {
		goto L501
	}
L471:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v475)+24))
	v1866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865)+116)))
	if v1866 != int32(1) {
		v1946 = v1861
		goto L470
	} else {
		goto L472
	}
L472:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v1869&int32(1) == int32(0) {
		v1946 = v1861
		goto L470
	} else {
		goto L473
	}
L473:
	;
	v1874 = int32(0)
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530)+2)))
	if v1876 != 0 {
		v1916 = v1874
		v1918 = v1874
		goto L474
	} else {
		goto L475
	}
L474:
	;
	if base.Ui32(v1031) < base.Ui32(v1853) {
		goto L486
	} else {
		goto L487
	}
L475:
	;
	v1877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v1877&int32(32) != 0 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	F_bt_entry_unique_check(m, v475, v1082, v1892, v1073, v474+int32(1120))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L5
	} else {
		goto L483
	}
L477:
	;
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v1880&int32(32) != 0 {
		goto L476
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v1883 = int32(0)
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v474)+1132))
	if v1884 == v1883 {
		v1916 = v1874
		v1918 = v1883
		goto L474
	} else {
		goto L481
	}
L480:
	;
	goto L479
L481:
	;
	v1887 = int32(0)
	v1888 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1884)+4)))
	if v1888 == v1887 {
		v1916 = v1874
		v1918 = v1887
		goto L474
	} else {
		goto L482
	}
L482:
	;
	goto L476
L483:
	;
	v1897 = int32(1)
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+12)))
	if v1898&v1897 == int32(0) {
		v1946 = v1897
		goto L470
	} else {
		goto L484
	}
L484:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v475)+24))
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1903)+116)))
	if v1904&int32(1) == int32(0) {
		v1946 = v1897
		goto L470
	} else {
		goto L485
	}
L485:
	;
	v1909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	v1910 = int32(1)
	v1916 = base.B2i32(v1909&v1910 == int32(0))
	v1918 = v1910
	goto L474
L486:
	;
	v1946 = v1918
	goto L470
L487:
	;
	goto L488
L488:
	;
	if v1916 != 0 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v1946 = v1918
	goto L470
L490:
	;
	goto L491
L491:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1530)+8)) = int32(0)
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v1924 = F__bt_compare(m, v1922, v1530, v1923, v1853)
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L5
	} else {
		goto L494
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1530)+8)) = v1919
	v1946 = v1944
	goto L470
L493:
	;
	if v1918 != 0 {
		v1944 = int32(1)
		goto L492
	} else {
		goto L499
	}
L494:
	;
	if v1924 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530)+2)))
	if v1928 != int32(1) {
		goto L493
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v474)+1128)) = int64(4294967295)
	v1933 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v474)+1124)) = uint16(v1933)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+1120)) = int32(-1)
	v1944 = v1918
	goto L492
L498:
	;
	goto L497
L499:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	F_bt_entry_unique_check(m, v475, v1082, v1938, v1073, v474+int32(1120))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L5
	} else {
		goto L500
	}
L500:
	;
	v1944 = int32(0)
	goto L492
L501:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v1953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1952)+16)))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1952+v1953)+4))
	if v1955 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2413 = int32(0)
	goto L289
L503:
	;
	goto L504
L504:
	;
	v1962 = v1955
	goto L505
L505:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[5]))
	if v1993 != 0 {
		goto L507
	} else {
		goto L508
	}
L506:
	;
	v2042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1996)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2042) {
		goto L524
	} else {
		goto L525
	}
L507:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L5
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v1996 = F_palloc_btree_page(m, v475, v1962)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L5
	} else {
		goto L512
	}
L510:
	;
	goto L509
L511:
	;
	goto L506
L512:
	;
	v1998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1996)+16)))
	v1999 = v1996 + v1998
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1999)+12)))
	if v2000&int32(20) == int32(0) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2005 == int32(0) {
		goto L511
	} else {
		goto L514
	}
L514:
	;
	v2010 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L5
	} else {
		goto L515
	}
L515:
	;
	if v2010 != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L5
	} else {
		goto L519
	}
L517:
	;
	goto L518
L518:
	;
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	F_pfree(m, v1996)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L5
	} else {
		goto L523
	}
L519:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+48))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+372)) = v1962
	*(*int32)(unsafe.Add(mBase, uint32(v474)+368)) = v2017
	*(*int32)(unsafe.Add(mBase, uint32(v474)+376)) = v2016 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_49), v474+int32(368))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L5
	} else {
		goto L520
	}
L520:
	;
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_50), int32(0))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L5
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1926), int32(_a_F_bt_index_check_callback_51))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L5
	} else {
		goto L522
	}
L522:
	;
	goto L518
L523:
	;
	v1962 = v2039
	goto L505
L524:
	;
	v2050 = int32(base.Ui32(v2042+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L526
L525:
	;
	v2050 = int32(0)
	goto L526
L526:
	;
	if v2000&int32(1) != 0 {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	v2077 = int32(0)
	v2080 = F_errstart(m, int32(13), v2077)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L5
	} else {
		goto L544
	}
L528:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2055 != 0 {
		goto L531
	} else {
		goto L532
	}
L529:
	;
	goto L530
L530:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2068 != 0 {
		goto L539
	} else {
		goto L540
	}
L531:
	;
	v2056 = int32(2)
	goto L533
L532:
	;
	v2056 = int32(1)
	goto L533
L533:
	;
	if base.Ui32(v2050&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v2056) {
		goto L527
	} else {
		goto L534
	}
L534:
	;
	v2060 = F_PageGetItemIdCareful_2(m, v475, v1962, v1996, v2056)
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L5
	} else {
		goto L535
	}
L535:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2064 != 0 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	v2065 = int32(2)
	goto L538
L537:
	;
	v2065 = int32(1)
	goto L538
L538:
	;
	v2440 = v2060
	v2442 = v2065
	goto L288
L539:
	;
	v2069 = int32(3)
	goto L541
L540:
	;
	v2069 = int32(2)
	goto L541
L541:
	;
	if base.Ui32(v2050&int32(_a_F_bt_index_check_callback_33)) < base.Ui32(v2069) {
		goto L527
	} else {
		goto L542
	}
L542:
	;
	v2074 = F_PageGetItemIdCareful_2(m, v475, v1962, v1996, v2069)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L5
	} else {
		goto L543
	}
L543:
	;
	v2440 = v2074
	v2442 = int32(0)
	goto L288
L544:
	;
	if v2080 != 0 {
		goto L290
	} else {
		goto L545
	}
L545:
	;
	v2413 = v2077
	goto L289
L546:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L5
	} else {
		goto L547
	}
L547:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2089)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+800)) = v2090 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_52), v474+int32(800))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L5
	} else {
		goto L548
	}
L548:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1074)))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2101 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+788)) = uint32(v2101)
	v2104 = int64(base.Ui64(v2101) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+784)) = uint32(v2104)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+776)) = v1085
	*(*int32)(unsafe.Add(mBase, uint32(v474)+772)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+768)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v474)+780)) = int32(base.Ui32(v2099) >> (uint(int32(17)) % 32))
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_53), v474+int32(768))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L5
	} else {
		goto L549
	}
L549:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_54), int32(0))
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L5
	} else {
		goto L550
	}
L550:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1327), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L5
	} else {
		goto L551
	}
L551:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L5
	} else {
		goto L553
	}
L553:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L5
	} else {
		goto L554
	}
L554:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+592)) = v2142 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_55), v474+int32(592))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L5
	} else {
		goto L555
	}
L555:
	;
	v2151 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+588)) = uint32(v2151)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+580)) = v1422
	*(*int32)(unsafe.Add(mBase, uint32(v474)+576)) = v2132
	v2156 = int64(base.Ui64(v2151) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+584)) = uint32(v2156)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_56), v474+int32(576))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L5
	} else {
		goto L556
	}
L556:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1428), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L5
	} else {
		goto L557
	}
L557:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L558:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+196)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+192)) = v2185
	v2191 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(192))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L5
	} else {
		goto L562
	}
L559:
	;
	goto L558
L560:
	;
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v2173&int32(32) == int32(0) {
		v2184 = v1082
		goto L559
	} else {
		goto L561
	}
L561:
	;
	v2178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v2179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v2184 = v1082 + (v2178 | v2179<<(uint(int32(16))%32))
	goto L559
L562:
	;
	v2193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2184)+2)))
	v2194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2184))))
	v2195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2184)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+180)) = v2195
	*(*int32)(unsafe.Add(mBase, uint32(v474)+176)) = v2193 | v2194<<(uint(int32(16))%32)
	v2204 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(176))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L5
	} else {
		goto L563
	}
L563:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L5
	} else {
		goto L564
	}
L564:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L5
	} else {
		goto L565
	}
L565:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2213)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+160)) = v1085
	*(*int32)(unsafe.Add(mBase, uint32(v474)+164)) = v2214 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_57), v474+int32(160))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L5
	} else {
		goto L566
	}
L566:
	;
	v2224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1043)+12)))
	v2225 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+144)) = uint32(v2225)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+136)) = v2204
	*(*int32)(unsafe.Add(mBase, uint32(v474)+128)) = v2191
	v2230 = int64(base.Ui64(v2225) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+140)) = uint32(v2230)
	if v2224&int32(1) != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v2236 = int32(_a_F_bt_index_check_callback_42)
	goto L569
L568:
	;
	v2236 = int32(_a_F_bt_index_check_callback_43)
	goto L569
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+132)) = v2236
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_58), v474+int32(128))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L5
	} else {
		goto L570
	}
L570:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1483), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L5
	} else {
		goto L571
	}
L571:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L572:
	;
	v2256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+7)))
	if v2256&int32(32) == int32(0) {
		v2272 = v1082
		goto L574
	} else {
		goto L575
	}
L573:
	;
	v2273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2272)+2)))
	v2274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2272))))
	v2275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2272)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+468)) = v2275
	*(*int32)(unsafe.Add(mBase, uint32(v474)+464)) = v2273 | v2274<<(uint(int32(16))%32)
	v2284 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(464))
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L5
	} else {
		goto L577
	}
L574:
	;
	goto L573
L575:
	;
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v2261&int32(32) == int32(0) {
		v2272 = v1082
		goto L574
	} else {
		goto L576
	}
L576:
	;
	v2266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v2267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v2272 = v1082 + (v2266 | v2267<<(uint(int32(16))%32))
	goto L574
L577:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2288 = v1851 & int32(_a_F_bt_index_check_callback_33)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+452)) = v2288
	*(*int32)(unsafe.Add(mBase, uint32(v474)+448)) = v2286
	v2294 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(448))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L5
	} else {
		goto L578
	}
L578:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2298 = F_PageGetItemIdCareful_2(m, v475, v2296, v2297, v2288)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L5
	} else {
		goto L579
	}
L579:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2298)))
	v2304 = v2300 + v2301&int32(_a_F_bt_index_check_callback_21)
	v2305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2304)+7)))
	if v2305&int32(32) == int32(0) {
		v2321 = v2304
		goto L581
	} else {
		goto L582
	}
L580:
	;
	v2322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321)+2)))
	v2323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321))))
	v2324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2321)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+436)) = v2324
	*(*int32)(unsafe.Add(mBase, uint32(v474)+432)) = v2322 | v2323<<(uint(int32(16))%32)
	v2333 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(432))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L5
	} else {
		goto L584
	}
L581:
	;
	goto L580
L582:
	;
	v2310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2304)+5)))
	if v2310&int32(32) == int32(0) {
		v2321 = v2304
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v2315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2304)+2)))
	v2316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2304))))
	v2321 = v2304 + (v2315 | v2316<<(uint(int32(16))%32))
	goto L581
L584:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L5
	} else {
		goto L585
	}
L585:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L5
	} else {
		goto L586
	}
L586:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+416)) = v2343 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_59), v474+int32(416))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L5
	} else {
		goto L587
	}
L587:
	;
	v2352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1043)+12)))
	v2353 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+412)) = uint32(v2353)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+404)) = v2333
	v2357 = int64(base.Ui64(v2353) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+408)) = uint32(v2357)
	if v2352&int32(1) != 0 {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v2363 = int32(_a_F_bt_index_check_callback_42)
	goto L590
L589:
	;
	v2363 = int32(_a_F_bt_index_check_callback_43)
	goto L590
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+400)) = v2363
	*(*int32)(unsafe.Add(mBase, uint32(v474)+396)) = v2294
	*(*int32)(unsafe.Add(mBase, uint32(v474)+392)) = v2284
	*(*int32)(unsafe.Add(mBase, uint32(v474)+384)) = v2254
	*(*int32)(unsafe.Add(mBase, uint32(v474)+388)) = v2363
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_60), v474+int32(384))
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L5
	} else {
		goto L591
	}
L591:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1641), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L5
	} else {
		goto L592
	}
L592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	v2382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1999)+12)))
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+324)) = v1962
	*(*int32)(unsafe.Add(mBase, uint32(v474)+328)) = v2384 + int32(4)
	if v2382&int32(1) != 0 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v2393 = int32(_a_F_bt_index_check_callback_34)
	goto L596
L595:
	;
	v2393 = int32(_a_F_bt_index_check_callback_35)
	goto L596
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+320)) = v2393
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_61), v474+int32(320))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L5
	} else {
		goto L597
	}
L597:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2057), int32(_a_F_bt_index_check_callback_51))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L5
	} else {
		goto L598
	}
L598:
	;
	v2413 = v2077
	goto L289
L599:
	;
	v2451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2449)+4)) = uint8(v2451)
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2455 = F__bt_compare(m, v2453, v2449, v2454, v1031)
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L5
	} else {
		goto L600
	}
L600:
	;
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2449))))
	if int32(0)-(v2458^int32(1)) < v2455 {
		v2517 = v2449
		v2519 = int32(1)
		v2522 = v2442
		goto L287
	} else {
		goto L601
	}
L601:
	;
	v2463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v2463 == int32(0) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2467 = F_palloc_btree_page(m, v475, v2466)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L5
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L5
	} else {
		goto L607
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+32)) = v2467
	v2470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2467)+16)))
	v2472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2467+v2470)+12)))
	if v2472&int32(20) != 0 {
		v3054 = v597
		v3064 = v598
		goto L121
	} else {
		goto L606
	}
L606:
	;
	goto L604
L607:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L5
	} else {
		goto L608
	}
L608:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2483)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+352)) = v2484 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_62), v474+int32(352))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L5
	} else {
		goto L609
	}
L609:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2494 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+348)) = uint32(v2494)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+340)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+336)) = v2493
	v2499 = int64(base.Ui64(v2494) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+344)) = uint32(v2499)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_63), v474+int32(336))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L5
	} else {
		goto L610
	}
L610:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1753), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L5
	} else {
		goto L611
	}
L611:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L612:
	;
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v475)+24))
	v2548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+116)))
	if v2519&v2548 != int32(1) {
		v2633 = v1043
		goto L286
	} else {
		goto L613
	}
L613:
	;
	v2552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+12)))
	if v2552&int32(1) == int32(0) {
		v2633 = v1043
		goto L286
	} else {
		goto L614
	}
L614:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	if v2557 == int32(0) {
		v2633 = v1043
		goto L286
	} else {
		goto L615
	}
L615:
	;
	v2562 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L5
	} else {
		goto L616
	}
L616:
	;
	if v2562 != 0 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_64), int32(0))
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L5
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2517)+8)) = int32(0)
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2577 = F__bt_compare(m, v2575, v2517, v2576, v1031)
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L5
	} else {
		goto L622
	}
L620:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1765), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L5
	} else {
		goto L621
	}
L621:
	;
	goto L619
L622:
	;
	if v2577 != 0 {
		v2633 = v1043
		goto L286
	} else {
		goto L623
	}
L623:
	;
	v2579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2517)+2)))
	if v2579 != 0 {
		v2633 = v1043
		goto L286
	} else {
		goto L624
	}
L624:
	;
	if v1946 == int32(0) {
		goto L625
	} else {
		goto L626
	}
L625:
	;
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	F_bt_entry_unique_check(m, v475, v1082, v2582, v1031, v474+int32(1120))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L5
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v2589 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L5
	} else {
		goto L629
	}
L628:
	;
	goto L627
L629:
	;
	if v2589 != 0 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_65), int32(0))
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L5
	} else {
		goto L633
	}
L631:
	;
	goto L632
L632:
	;
	v2600 = F_palloc_btree_page(m, v475, v2557)
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L5
	} else {
		goto L635
	}
L633:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1788), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L5
	} else {
		goto L634
	}
L634:
	;
	goto L632
L635:
	;
	v2602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2600)+16)))
	v2603 = v2600 + v2602
	v2604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2603)+12)))
	if v2604&int32(20) != 0 {
		goto L276
	} else {
		goto L636
	}
L636:
	;
	if v2604&int32(1) == int32(0) {
		goto L284
	} else {
		goto L637
	}
L637:
	;
	v2611 = F_PageGetItemIdCareful_2(m, v475, v2557, v2600, v2522)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L5
	} else {
		goto L638
	}
L638:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2611)))
	F_bt_entry_unique_check(m, v475, v2600+v2613&int32(_a_F_bt_index_check_callback_21), v2557, v2522, v474+int32(1120))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L5
	} else {
		goto L639
	}
L639:
	;
	F_pfree(m, v2600)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L5
	} else {
		goto L640
	}
L640:
	;
	v2633 = v2603
	goto L286
L641:
	;
	v2659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v2659 != int32(1) {
		v2866 = v2633
		goto L285
	} else {
		goto L642
	}
L642:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2664 = F_PageGetItemIdCareful_2(m, v475, v2662, v2663, v1073)
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L5
	} else {
		goto L643
	}
L643:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v2667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2666)+16)))
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2664)))
	v2671 = v2666 + v2668&int32(_a_F_bt_index_check_callback_21)
	v2672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2671))))
	v2675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2671)+2)))
	v2676 = v2672<<(uint(int32(16))%32) | v2675
	v2677 = F_palloc_btree_page(m, v475, v2676)
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L5
	} else {
		goto L644
	}
L644:
	;
	v2679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2677)+16)))
	v2680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2677)+12)))
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v2666+v2667)+8))
	F_bt_child_highkey_check(m, v475, v1073, v2677, v2682)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L5
	} else {
		goto L645
	}
L645:
	;
	v2685 = v2677 + v2679
	v2686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2685)+12)))
	if v2686&int32(4) != 0 {
		goto L277
	} else {
		goto L646
	}
L646:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2685)+4))
	if v2691 != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v2692 = int32(2)
	goto L649
L648:
	;
	v2692 = int32(1)
	goto L649
L649:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v2680) {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v2700 = int32(base.Ui32(v2680+int32(_a_F_bt_index_check_callback_32)) >> (uint(int32(2)) % 32))
	goto L652
L651:
	;
	v2700 = int32(0)
	goto L652
L652:
	;
	v2702 = v2700 & int32(_a_F_bt_index_check_callback_33)
	if base.Ui32(v2692) <= base.Ui32(v2702) {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v2705 = v2692
	goto L656
L654:
	;
	goto L655
L655:
	;
	F_pfree(m, v2677)
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L5
	} else {
		goto L703
	}
L656:
	;
	v2737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2685)+12)))
	if v2737&int32(1) == int32(0) {
		goto L659
	} else {
		goto L660
	}
L657:
	;
	goto L655
L658:
	;
	v2817 = v2705 + int32(1)
	if base.Ui32(v2817&int32(_a_F_bt_index_check_callback_33)) <= base.Ui32(v2702) {
		v2705 = v2817
		goto L656
	} else {
		goto L702
	}
L659:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2685)+4))
	if v2746 != 0 {
		goto L662
	} else {
		goto L663
	}
L660:
	;
	goto L661
L661:
	;
	v2750 = v2705 & int32(_a_F_bt_index_check_callback_33)
	v2751 = F_PageGetItemIdCareful_2(m, v475, v2676, v2677, v2750)
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L5
	} else {
		goto L666
	}
L662:
	;
	v2747 = int32(2)
	goto L664
L663:
	;
	v2747 = int32(1)
	goto L664
L664:
	;
	if v2705&int32(_a_F_bt_index_check_callback_33) == v2747 {
		goto L658
	} else {
		goto L665
	}
L665:
	;
	goto L661
L666:
	;
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2754 = F__bt_compare(m, v2753, v1530, v2677, v2750)
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L5
	} else {
		goto L667
	}
L667:
	;
	v2756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1530))))
	if v2756 == int32(0) {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	if int32(0) < v2754 {
		goto L49
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	if v2754 == int32(0) {
		goto L672
	} else {
		goto L673
	}
L671:
	;
	goto L658
L672:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2751)))
	v2766 = v2677 + v2763&int32(_a_F_bt_index_check_callback_21)
	v2768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2677)+16)))
	v2769 = v2677 + v2768
	v2770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2769)+12)))
	if v2770&int32(1) != 0 {
		goto L675
	} else {
		goto L676
	}
L673:
	;
	goto L674
L674:
	;
	if int32(0) <= v2754 {
		goto L49
	} else {
		goto L701
	}
L675:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(v2769)+4))
	if v2775 != 0 {
		goto L678
	} else {
		goto L679
	}
L676:
	;
	v2778 = int32(0)
	goto L677
L677:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2779)+192))
	v2781 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2780)+10)))
	v2782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2766)+7)))
	if v2782&int32(32) != 0 {
		goto L683
	} else {
		goto L684
	}
L678:
	;
	v2776 = int32(2)
	goto L680
L679:
	;
	v2776 = int32(1)
	goto L680
L680:
	;
	v2778 = base.B2i32(base.Ui32(v2776) <= base.Ui32(v2750))
	goto L677
L681:
	;
	v2801 = F_BTreeTupleGetHeapTIDCareful(m, v475, v2766, v2778)
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L5
	} else {
		goto L694
	}
L682:
	;
	v2798 = v2796
	goto L681
L683:
	;
	v2785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2766)+4)))
	if v2785&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L686
	} else {
		goto L687
	}
L684:
	;
	goto L685
L685:
	;
	v2794 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2780)+8)))
	if v2781 < v2794 {
		v2798 = v2781
		goto L681
	} else {
		goto L693
	}
L686:
	;
	v2788 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2780)+8)))
	if v2788 <= v2781 {
		v2796 = v2788
		goto L682
	} else {
		goto L689
	}
L687:
	;
	goto L688
L688:
	;
	v2791 = v2785 & int32(4095)
	if v2781 < v2791 {
		goto L690
	} else {
		goto L691
	}
L689:
	;
	v2798 = v2781
	goto L681
L690:
	;
	v2793 = v2781
	goto L692
L691:
	;
	v2793 = v2791
	goto L692
L692:
	;
	v2798 = v2793
	goto L681
L693:
	;
	v2796 = v2794
	goto L682
L694:
	;
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+12))
	if v2798 == v2803 {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+8))
	if v2805 != 0 {
		goto L49
	} else {
		goto L698
	}
L696:
	;
	goto L697
L697:
	;
	if v2803 < v2798 {
		goto L658
	} else {
		goto L700
	}
L698:
	;
	if v2801 == int32(0) {
		goto L49
	} else {
		goto L699
	}
L699:
	;
	goto L658
L700:
	;
	goto L49
L701:
	;
	goto L658
L702:
	;
	goto L657
L703:
	;
	v2866 = v2633
	goto L285
L704:
	;
	v3006 = v2866
	goto L214
L705:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L5
	} else {
		goto L706
	}
L706:
	;
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2901)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+304)) = v2902 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_66), v474+int32(304))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L5
	} else {
		goto L707
	}
L707:
	;
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2912 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+296)) = uint32(v2912)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+288)) = v2911
	v2916 = int64(base.Ui64(v2912) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+292)) = uint32(v2916)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_67), v474+int32(288))
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L5
	} else {
		goto L708
	}
L708:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1806), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L5
	} else {
		goto L709
	}
L709:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L710:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L5
	} else {
		goto L711
	}
L711:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2935)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+272)) = v2936 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_68), v474+int32(272))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L5
	} else {
		goto L712
	}
L712:
	;
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v2946 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+268)) = uint32(v2946)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+260)) = v2676
	*(*int32)(unsafe.Add(mBase, uint32(v474)+256)) = v2945
	v2951 = int64(base.Ui64(v2946) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+264)) = uint32(v2951)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_69), v474+int32(256))
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L5
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2498), int32(_a_F_bt_index_check_callback_70))
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L5
	} else {
		goto L714
	}
L714:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L715:
	;
	v3006 = v2603
	goto L214
L716:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		goto L5
	} else {
		goto L717
	}
L717:
	;
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+896)) = v2973 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_71), v474+int32(896))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L5
	} else {
		goto L718
	}
L718:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+888)) = v2982
	*(*int32)(unsafe.Add(mBase, uint32(v474)+884)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v474)+880)) = v471
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_72), v474+int32(880))
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L5
	} else {
		goto L719
	}
L719:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(779), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L5
	} else {
		goto L720
	}
L720:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L721:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v3006)+4))
	if v3032 != 0 {
		v3054 = v597
		v3064 = v598
		goto L121
	} else {
		goto L722
	}
L722:
	;
	v3033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v3033 != int32(1) {
		v3054 = v597
		v3064 = v598
		goto L121
	} else {
		goto L723
	}
L723:
	;
	v3036 = int32(0)
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(v3006)+8))
	F_bt_child_highkey_check(m, v475, v3036, v3036, v3038)
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L5
	} else {
		goto L724
	}
L724:
	;
	v3054 = v597
	v3064 = v598
	goto L121
L725:
	;
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	if v471 == v3075 {
		goto L98
	} else {
		goto L726
	}
L726:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v475)+48))
	if v3078 != 0 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	F_pfree(m, v3078)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L5
	} else {
		goto L730
	}
L728:
	;
	goto L729
L729:
	;
	v3083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
	if v3083 != int32(1) {
		goto L731
	} else {
		goto L732
	}
L730:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+48)) = int32(0)
	goto L729
L731:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v475)+16))
	F_MemoryContextReset(m, v3112)
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L5
	} else {
		goto L740
	}
L732:
	;
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v3086 == int32(0) {
		goto L731
	} else {
		goto L733
	}
L733:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v3092 = F_PageGetItemIdCareful_2(m, v475, v3089, v3090, int32(1))
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L5
	} else {
		goto L734
	}
L734:
	;
	v3094 = *(*int32)(unsafe.Add(mBase, uint32(v475)+32))
	v3095 = *(*int32)(unsafe.Add(mBase, uint32(v3092)))
	v3098 = v3094 + v3095&int32(_a_F_bt_index_check_callback_21)
	v3099 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3098)+6)))
	v3102 = F_MemoryContextAlloc(m, v496, v3099&int32(_a_F_bt_index_check_callback_40))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L5
	} else {
		goto L735
	}
L735:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+48)) = v3102
	v3105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3098)+6)))
	v3107 = v3105 & int32(_a_F_bt_index_check_callback_40)
	if v3107 != 0 {
		goto L737
	} else {
		goto L738
	}
L736:
	;
	goto L731
L737:
	;
	v3108 = F__emscripten_memcpy_bulkmem(m, v3102, v3098, v3107)
	mBase = m.M
	goto L739
L738:
	;
	goto L739
L739:
	;
	goto L736
L740:
	;
	if v3077 != 0 {
		__phi470 = v471
		__phi471 = v3077
		__phi483 = v3054
		__phi493 = v3064
		v470 = __phi470
		v471 = __phi471
		v483 = __phi483
		v493 = __phi493
		goto L114
	} else {
		goto L741
	}
L741:
	;
	goto L115
L742:
	;
	F_pfree(m, v3115)
	mBase = m.M
	v3117 = m.ExcPending
	if v3117 != 0 {
		goto L5
	} else {
		goto L745
	}
L743:
	;
	goto L744
L744:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bt_index_check_callback[4])) = v496
	if v3064 != int32(-1) {
		v403 = v3064
		v406 = v474
		v407 = v475
		v415 = v3054
		v416 = v484
		v420 = v488
		v422 = v3054
		v426 = int32(0)
		v427 = v495
		goto L99
	} else {
		goto L746
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475)+48)) = int32(0)
	goto L744
L746:
	;
	goto L100
L747:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L5
	} else {
		goto L748
	}
L748:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v488)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+116)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v474)+112)) = v3132 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_73), v474+int32(112))
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L5
	} else {
		goto L749
	}
L749:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(535), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3146 = m.ExcPending
	if v3146 != 0 {
		goto L5
	} else {
		goto L750
	}
L750:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L751:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L5
	} else {
		goto L752
	}
L752:
	;
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+96)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v474)+100)) = v3155 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_74), v474+int32(96))
	mBase = m.M
	v3164 = m.ExcPending
	if v3164 != 0 {
		goto L5
	} else {
		goto L753
	}
L753:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(791), int32(_a_F_bt_index_check_callback_19))
	mBase = m.M
	v3169 = m.ExcPending
	if v3169 != 0 {
		goto L5
	} else {
		goto L754
	}
L754:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L755:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3176 = m.ExcPending
	if v3176 != 0 {
		goto L5
	} else {
		goto L756
	}
L756:
	;
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+1072)) = v3177 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_75), v193+int32(1072))
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L5
	} else {
		goto L757
	}
L757:
	;
	F_errhint(m, int32(_a_F_bt_index_check_callback_76), int32(0))
	mBase = m.M
	v3189 = m.ExcPending
	if v3189 != 0 {
		goto L5
	} else {
		goto L758
	}
L758:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(485), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L5
	} else {
		goto L759
	}
L759:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L760:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L5
	} else {
		goto L761
	}
L761:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+1088)) = v3202 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_77), v193+int32(1088))
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L5
	} else {
		goto L762
	}
L762:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(463), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L5
	} else {
		goto L763
	}
L763:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L764:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v3220 = F_BuildIndexInfo(m, v3219)
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		goto L5
	} else {
		goto L767
	}
L765:
	;
	goto L766
L766:
	;
	v3723 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	if v3723 != 0 {
		goto L826
	} else {
		goto L827
	}
L767:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	v3224 = int32(0)
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+188))
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v3230 = m.T0[v3229].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v3222, v3223, v3224, v3224, v3224, int32(449))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L5
	} else {
		goto L768
	}
L768:
	;
	v3232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3220)+116)) = uint8(v3232)
	v3234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3220)+121)) = uint8(v3234)
	*(*int32)(unsafe.Add(mBase, uint32(v3220)+100)) = v3232
	*(*int64)(unsafe.Add(mBase, uint32(v3220)+92)) = int64(0)
	v3242 = F_errstart(m, int32(14), v3232)
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L5
	} else {
		goto L769
	}
L769:
	;
	if v3242 != 0 {
		goto L770
	} else {
		goto L771
	}
L770:
	;
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(v3244)+48))
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3246)+48))
	v3248 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v406)+36)) = v3247 + v3248
	*(*int32)(unsafe.Add(mBase, uint32(v406)+32)) = v3245 + v3248
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_78), v406+int32(32))
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L5
	} else {
		goto L773
	}
L771:
	;
	goto L772
L772:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v3268 = int32(0)
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(v3265)+188))
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3273)+140))
	v3275 = m.T0[v3274].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v3265, v3266, v3220, int32(1), v3268, v3268, v3268, int32(-1), int32(_a_F_bt_index_check_callback_79), v407, v3230)
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L5
	} else {
		goto L775
	}
L773:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(586), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L5
	} else {
		goto L774
	}
L774:
	;
	goto L772
L775:
	;
	v3279 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L5
	} else {
		goto L776
	}
L776:
	;
	if v3279 != 0 {
		goto L777
	} else {
		goto L778
	}
L777:
	;
	v3281 = *(*int64)(unsafe.Add(mBase, uint32(v407)+64))
	v3282 = *(*int32)(unsafe.Add(mBase, uint32(v427)+48))
	v3283 = int64(0)
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v407)+60))
	v3286 = v3284 + int32(24)
	v3287 = *(*int64)(unsafe.Add(mBase, uint32(v3284)+16))
	v3290 = base.I32_wrap_i64(int64(base.Ui64(v3287) >> (uint(int64(3)) % 64)))
	if v3290 <= int32(3) {
		goto L781
	} else {
		goto L782
	}
L778:
	;
	goto L779
L779:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(v407)+60))
	F_pfree(m, v3687)
	mBase = m.M
	v3689 = m.ExcPending
	if v3689 != 0 {
		goto L5
	} else {
		goto L825
	}
L780:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v406)+16)) = base.F64_mul(base.F64_div(base.F64_convert_i64_u(v3632), base.F64_convert_i64_u(v3633)), float64(100))
	*(*int32)(unsafe.Add(mBase, uint32(v406)+8)) = v3282 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v406))) = v3281
	F_errmsg_internal(m, int32(_a_F_bt_index_check_callback_80), v406)
	mBase = m.M
	v3648 = m.ExcPending
	if v3648 != 0 {
		goto L5
	} else {
		goto L823
	}
L781:
	;
	if v3290 == int32(0) {
		v3632 = v3283
		v3633 = v3287
		goto L780
	} else {
		goto L784
	}
L782:
	;
	goto L783
L783:
	;
	v3443 = int64(0)
	if v3290 < int32(4) {
		v3522 = v3286
		v3523 = v3290
		v3528 = v3443
		goto L796
	} else {
		goto L797
	}
L784:
	;
	v3296 = v3290 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3290) {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v3302 = v3286
	v3304 = int32(0)
	v3331 = v3283
	goto L788
L786:
	;
	v3360 = v3286
	v3389 = v3283
	goto L787
L787:
	;
	if v3296 == int32(0) {
		v3632 = v3389
		v3633 = v3287
		goto L780
	} else {
		goto L791
	}
L788:
	;
	v3335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3302)+3)))
	v3338 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3335)+uint32(_c_F_bt_index_check_callback[11]))))
	v3339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3302)+2)))
	v3342 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3339)+uint32(_c_F_bt_index_check_callback[11]))))
	v3343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3302)+1)))
	v3346 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3343)+uint32(_c_F_bt_index_check_callback[11]))))
	v3347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3302))))
	v3350 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3347)+uint32(_c_F_bt_index_check_callback[11]))))
	v3354 = v3338 + (v3342 + (v3346 + (v3331 + v3350)))
	v3355 = int32(4)
	v3356 = v3302 + v3355
	v3358 = v3304 + v3355
	if v3358 != v3290&int32(-4) {
		v3302 = v3356
		v3304 = v3358
		v3331 = v3354
		goto L788
	} else {
		goto L790
	}
L789:
	;
	v3360 = v3356
	v3389 = v3354
	goto L787
L790:
	;
	goto L789
L791:
	;
	v3396 = v3360
	v3398 = int32(0)
	v3425 = v3389
	goto L792
L792:
	;
	v3429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3396))))
	v3432 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3429)+uint32(_c_F_bt_index_check_callback[11]))))
	v3433 = v3425 + v3432
	v3434 = int32(1)
	v3437 = v3398 + v3434
	if v3437 != v3296 {
		v3396 = v3396 + v3434
		v3398 = v3437
		v3425 = v3433
		goto L792
	} else {
		goto L794
	}
L793:
	;
	v3632 = v3433
	v3633 = v3287
	goto L780
L794:
	;
	goto L793
L795:
	;
	v3602 = *(*int64)(unsafe.Add(mBase, uint32(v3284)+16))
	v3632 = v3601
	v3633 = v3602
	goto L780
L796:
	;
	if v3523 == int32(0) {
		v3601 = v3528
		goto L810
	} else {
		goto L811
	}
L797:
	;
	if v3286 != (v3284+int32(27))&int32(-4) {
		v3522 = v3286
		v3523 = v3290
		v3528 = v3443
		goto L796
	} else {
		goto L798
	}
L798:
	;
	v3452 = v3290 - int32(4)
	v3456 = int32(base.Ui32(v3452)>>(uint(int32(2))%32)) + int32(1)
	v3458 = v3456 & int32(3)
	if base.Ui32(v3452) < base.Ui32(int32(12)) {
		goto L800
	} else {
		goto L801
	}
L799:
	;
	if v3458 == int32(0) {
		v3522 = v3494
		v3523 = v3495
		v3528 = v3500
		goto L796
	} else {
		goto L806
	}
L800:
	;
	v3494 = v3286
	v3495 = v3290
	v3500 = v3443
	goto L799
L801:
	;
	goto L802
L802:
	;
	v3464 = v3286
	v3465 = v3290
	v3466 = int32(0)
	v3470 = v3443
	goto L803
L803:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3464)+12))
	v3474 = *(*int32)(unsafe.Add(mBase, uint32(v3464)+8))
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3464)+4))
	v3480 = *(*int32)(unsafe.Add(mBase, uint32(v3464)))
	v3486 = base.I64_extend_i32_u(base.I32_popcnt(v3471)) + (base.I64_extend_i32_u(base.I32_popcnt(v3474)) + (base.I64_extend_i32_u(base.I32_popcnt(v3477)) + (v3470 + base.I64_extend_i32_u(base.I32_popcnt(v3480)))))
	v3487 = int32(16)
	v3488 = v3465 - v3487
	v3490 = v3464 + v3487
	v3492 = v3466 + int32(4)
	if v3492 != v3456&int32(2147483644) {
		v3464 = v3490
		v3465 = v3488
		v3466 = v3492
		v3470 = v3486
		goto L803
	} else {
		goto L805
	}
L804:
	;
	v3494 = v3490
	v3495 = v3488
	v3500 = v3486
	goto L799
L805:
	;
	goto L804
L806:
	;
	v3505 = v3495
	v3506 = v3494
	v3507 = int32(0)
	v3510 = v3500
	goto L807
L807:
	;
	v3511 = int32(4)
	v3512 = v3505 - v3511
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v3506)))
	v3516 = v3510 + base.I64_extend_i32_u(base.I32_popcnt(v3513))
	v3518 = v3506 + v3511
	v3520 = v3507 + int32(1)
	if v3520 != v3458 {
		v3505 = v3512
		v3506 = v3518
		v3507 = v3520
		v3510 = v3516
		goto L807
	} else {
		goto L809
	}
L808:
	;
	v3522 = v3518
	v3523 = v3512
	v3528 = v3516
	goto L796
L809:
	;
	goto L808
L810:
	;
	goto L795
L811:
	;
	v3532 = v3523 & int32(3)
	if v3532 == int32(0) {
		goto L813
	} else {
		goto L814
	}
L812:
	;
	if base.Ui32(v3523) < base.Ui32(int32(4)) {
		v3601 = v3561
		goto L810
	} else {
		goto L819
	}
L813:
	;
	v3555 = v3522
	v3557 = v3523
	v3561 = v3528
	goto L812
L814:
	;
	goto L815
L815:
	;
	v3538 = v3523
	v3539 = v3522
	v3540 = int32(0)
	v3542 = v3528
	goto L816
L816:
	;
	v3543 = int32(1)
	v3544 = v3538 - v3543
	v3545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3539))))
	v3548 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3545)+uint32(_c_F_bt_index_check_callback[11]))))
	v3549 = v3542 + v3548
	v3551 = v3539 + v3543
	v3553 = v3540 + v3543
	if v3553 != v3532 {
		v3538 = v3544
		v3539 = v3551
		v3540 = v3553
		v3542 = v3549
		goto L816
	} else {
		goto L818
	}
L817:
	;
	v3555 = v3551
	v3557 = v3544
	v3561 = v3549
	goto L812
L818:
	;
	goto L817
L819:
	;
	v3564 = v3555
	v3566 = v3557
	v3570 = v3561
	goto L820
L820:
	;
	v3571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3564)+3)))
	v3574 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3571)+uint32(_c_F_bt_index_check_callback[11]))))
	v3575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3564)+2)))
	v3578 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3575)+uint32(_c_F_bt_index_check_callback[11]))))
	v3579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3564)+1)))
	v3582 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3579)+uint32(_c_F_bt_index_check_callback[11]))))
	v3583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3564))))
	v3586 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3583)+uint32(_c_F_bt_index_check_callback[11]))))
	v3590 = v3574 + (v3578 + (v3582 + (v3570 + v3586)))
	v3591 = int32(4)
	v3594 = v3566 - v3591
	if v3594 != 0 {
		v3564 = v3564 + v3591
		v3566 = v3594
		v3570 = v3590
		goto L820
	} else {
		goto L822
	}
L821:
	;
	v3601 = v3590
	goto L810
L822:
	;
	goto L821
L823:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(594), int32(_a_F_bt_index_check_callback_6))
	mBase = m.M
	v3653 = m.ExcPending
	if v3653 != 0 {
		goto L5
	} else {
		goto L824
	}
L824:
	;
	goto L779
L825:
	;
	goto L766
L826:
	;
	F_UnregisterSnapshot(m, v3723)
	mBase = m.M
	v3725 = m.ExcPending
	if v3725 != 0 {
		goto L5
	} else {
		goto L829
	}
L827:
	;
	goto L828
L828:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	F_MemoryContextDelete(m, v3726)
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L5
	} else {
		goto L830
	}
L829:
	;
	goto L828
L830:
	;
	m.G0 = v406 + int32(1168)
	goto L46
L831:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3743 = m.ExcPending
	if v3743 != 0 {
		goto L5
	} else {
		goto L832
	}
L832:
	;
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3744)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+240)) = v3745 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_81), v474+int32(240))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L5
	} else {
		goto L833
	}
L833:
	;
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v3755 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+224)) = uint32(v3755)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+216)) = v2705 & int32(_a_F_bt_index_check_callback_33)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+212)) = v2676
	*(*int32)(unsafe.Add(mBase, uint32(v474)+208)) = v3754
	v3763 = int64(base.Ui64(v3755) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+220)) = uint32(v3763)
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_82), v474+int32(208))
	mBase = m.M
	v3769 = m.ExcPending
	if v3769 != 0 {
		goto L5
	} else {
		goto L834
	}
L834:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(2539), int32(_a_F_bt_index_check_callback_70))
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L5
	} else {
		goto L835
	}
L835:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L836:
	;
	v3792 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+564)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v474)+560)) = v3792
	v3798 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(560))
	mBase = m.M
	v3799 = m.ExcPending
	if v3799 != 0 {
		goto L5
	} else {
		goto L840
	}
L837:
	;
	goto L836
L838:
	;
	v3780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+5)))
	if v3780&int32(32) == int32(0) {
		v3791 = v1082
		goto L837
	} else {
		goto L839
	}
L839:
	;
	v3785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082)+2)))
	v3786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082))))
	v3791 = v1082 + (v3785 | v3786<<(uint(int32(16))%32))
	goto L837
L840:
	;
	v3800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3791)+2)))
	v3801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3791))))
	v3802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3791)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+548)) = v3802
	*(*int32)(unsafe.Add(mBase, uint32(v474)+544)) = v3800 | v3801<<(uint(int32(16))%32)
	v3811 = F_psprintf(m, int32(_a_F_bt_index_check_callback_39), v474+int32(544))
	mBase = m.M
	v3812 = m.ExcPending
	if v3812 != 0 {
		goto L5
	} else {
		goto L841
	}
L841:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L5
	} else {
		goto L842
	}
L842:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3819 = m.ExcPending
	if v3819 != 0 {
		goto L5
	} else {
		goto L843
	}
L843:
	;
	v3820 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v3821 = *(*int32)(unsafe.Add(mBase, uint32(v3820)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+528)) = v3821 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_83), v474+int32(528))
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L5
	} else {
		goto L844
	}
L844:
	;
	v3830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1043)+12)))
	v3831 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+512)) = uint32(v3831)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+504)) = v3811
	*(*int32)(unsafe.Add(mBase, uint32(v474)+496)) = v3798
	v3836 = int64(base.Ui64(v3831) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+508)) = uint32(v3836)
	if v3830&int32(1) != 0 {
		goto L845
	} else {
		goto L846
	}
L845:
	;
	v3842 = int32(_a_F_bt_index_check_callback_42)
	goto L847
L846:
	;
	v3842 = int32(_a_F_bt_index_check_callback_43)
	goto L847
L847:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+500)) = v3842
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_58), v474+int32(496))
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L5
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1590), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L5
	} else {
		goto L849
	}
L849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L850:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L5
	} else {
		goto L851
	}
L851:
	;
	v3863 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v3864 = *(*int32)(unsafe.Add(mBase, uint32(v3863)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+848)) = v3864 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_84), v474+int32(848))
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L5
	} else {
		goto L852
	}
L852:
	;
	v3873 = *(*int32)(unsafe.Add(mBase, uint32(v475)+36))
	v3876 = v3854 + v3855&int32(_a_F_bt_index_check_callback_21)
	v3877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3876)+7)))
	if v3877&int32(32) == int32(0) {
		goto L854
	} else {
		goto L855
	}
L853:
	;
	v3893 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v887)+12)))
	v3894 = *(*int64)(unsafe.Add(mBase, uint32(v475)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+832)) = uint32(v3894)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+820)) = v3892
	*(*int32)(unsafe.Add(mBase, uint32(v474)+816)) = v3873
	v3899 = int64(base.Ui64(v3894) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v474)+828)) = uint32(v3899)
	if v3893&int32(1) != 0 {
		goto L857
	} else {
		goto L858
	}
L854:
	;
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3888)+192))
	v3890 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3889)+8)))
	v3892 = v3890
	goto L853
L855:
	;
	v3882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3876)+4)))
	if v3882&int32(_a_F_bt_index_check_callback_10) != 0 {
		goto L854
	} else {
		goto L856
	}
L856:
	;
	v3892 = v3882 & int32(4095)
	goto L853
L857:
	;
	v3905 = int32(_a_F_bt_index_check_callback_42)
	goto L859
L858:
	;
	v3905 = int32(_a_F_bt_index_check_callback_43)
	goto L859
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+824)) = v3905
	F_errdetail_internal(m, int32(_a_F_bt_index_check_callback_85), v474+int32(816))
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L5
	} else {
		goto L860
	}
L860:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(1278), int32(_a_F_bt_index_check_callback_37))
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L5
	} else {
		goto L861
	}
L861:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L862:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L5
	} else {
		goto L863
	}
L863:
	;
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v3927 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_86), v34+int32(-32))
	mBase = m.M
	v3935 = m.ExcPending
	if v3935 != 0 {
		goto L5
	} else {
		goto L864
	}
L864:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(322), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		goto L5
	} else {
		goto L865
	}
L865:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L866:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3947 = m.ExcPending
	if v3947 != 0 {
		goto L5
	} else {
		goto L867
	}
L867:
	;
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v3948 + int32(4)
	F_errmsg(m, int32(_a_F_bt_index_check_callback_87), v34+int32(-48))
	mBase = m.M
	v3956 = m.ExcPending
	if v3956 != 0 {
		goto L5
	} else {
		goto L868
	}
L868:
	;
	F_errfinish(m, int32(_a_F_bt_index_check_callback_2), int32(330), int32(_a_F_bt_index_check_callback_3))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L5
	} else {
		goto L869
	}
L869:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
