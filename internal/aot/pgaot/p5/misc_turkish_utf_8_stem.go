package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_turkish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
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
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
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
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
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
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1832 int32
	_ = v1832
	var v1840 int32
	_ = v1840
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
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
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
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1968 int32
	_ = v1968
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2050 int32
	_ = v2050
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2115 int32
	_ = v2115
	var v2122 int32
	_ = v2122
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2218 int32
	_ = v2218
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2283 int32
	_ = v2283
	var v2290 int32
	_ = v2290
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2350 int32
	_ = v2350
	var v2354 int32
	_ = v2354
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
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
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2508 int32
	_ = v2508
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2559 int32
	_ = v2559
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2609 int32
	_ = v2609
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2643 int32
	_ = v2643
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2702 int32
	_ = v2702
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2767 int32
	_ = v2767
	var v2774 int32
	_ = v2774
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2812 int32
	_ = v2812
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2831 int32
	_ = v2831
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2849 int32
	_ = v2849
	var v2854 int32
	_ = v2854
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2869 int32
	_ = v2869
	var v2871 int32
	_ = v2871
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2884 int32
	_ = v2884
	var v2898 int32
	_ = v2898
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2934 int32
	_ = v2934
	var v2936 int32
	_ = v2936
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2953 int32
	_ = v2953
	var v2971 int32
	_ = v2971
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2984 int32
	_ = v2984
	var v2990 int32
	_ = v2990
	var v3006 int32
	_ = v3006
	var v3013 int32
	_ = v3013
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3029 int32
	_ = v3029
	var v3033 int32
	_ = v3033
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3056 int32
	_ = v3056
	var v3060 int32
	_ = v3060
	var v3062 int32
	_ = v3062
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3099 int32
	_ = v3099
	var v3100 int32
	_ = v3100
	var v3102 int32
	_ = v3102
	var v3106 int32
	_ = v3106
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3129 int32
	_ = v3129
	var v3133 int32
	_ = v3133
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3143 int32
	_ = v3143
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3148 int32
	_ = v3148
	var v3149 int32
	_ = v3149
	var v3156 int32
	_ = v3156
	var v3160 int32
	_ = v3160
	var v3162 int32
	_ = v3162
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3181 int32
	_ = v3181
	var v3185 int32
	_ = v3185
	var v3191 int32
	_ = v3191
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3202 int32
	_ = v3202
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3260 int32
	_ = v3260
	var v3262 int32
	_ = v3262
	var v3268 int32
	_ = v3268
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3285 int32
	_ = v3285
	var v3289 int32
	_ = v3289
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3311 int32
	_ = v3311
	var v3315 int32
	_ = v3315
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3333 int32
	_ = v3333
	var v3340 int32
	_ = v3340
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3378 int32
	_ = v3378
	var v3383 int32
	_ = v3383
	var v3386 int32
	_ = v3386
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3416 int32
	_ = v3416
	var v3418 int32
	_ = v3418
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3447 int32
	_ = v3447
	var v3450 int32
	_ = v3450
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3458 int32
	_ = v3458
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3472 int32
	_ = v3472
	var v3476 int32
	_ = v3476
	var v3480 int32
	_ = v3480
	var v3495 int32
	_ = v3495
	var v3499 int32
	_ = v3499
	var v3508 int32
	_ = v3508
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3519 int32
	_ = v3519
	var v3521 int32
	_ = v3521
	var v3527 int32
	_ = v3527
	var v3529 int32
	_ = v3529
	var v3531 int32
	_ = v3531
	var v3533 int32
	_ = v3533
	var v3546 int32
	_ = v3546
	var v3548 int32
	_ = v3548
	var v3550 int32
	_ = v3550
	var v3568 int32
	_ = v3568
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3581 int32
	_ = v3581
	var v3587 int32
	_ = v3587
	var v3596 int32
	_ = v3596
	var v3611 int32
	_ = v3611
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3628 int32
	_ = v3628
	var v3630 int32
	_ = v3630
	var v3632 int32
	_ = v3632
	var v3633 int32
	_ = v3633
	var v3636 int32
	_ = v3636
	var v3639 int32
	_ = v3639
	var v3643 int32
	_ = v3643
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3652 int32
	_ = v3652
	var v3653 int32
	_ = v3653
	var v3654 int32
	_ = v3654
	var v3656 int32
	_ = v3656
	var v3658 int32
	_ = v3658
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3664 int32
	_ = v3664
	var v3674 int32
	_ = v3674
	var v3679 int32
	_ = v3679
	var v3680 int32
	_ = v3680
	var v3684 int32
	_ = v3684
	var v3686 int32
	_ = v3686
	var v3688 int32
	_ = v3688
	var v3689 int32
	_ = v3689
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3699 int32
	_ = v3699
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3707 int32
	_ = v3707
	var v3710 int32
	_ = v3710
	var v3713 int32
	_ = v3713
	var v3716 int32
	_ = v3716
	var v3720 int32
	_ = v3720
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3729 int32
	_ = v3729
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3741 int32
	_ = v3741
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3750 int32
	_ = v3750
	var v3756 int32
	_ = v3756
	var v3757 int32
	_ = v3757
	var v3762 int32
	_ = v3762
	var v3763 int32
	_ = v3763
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3779 int32
	_ = v3779
	var v3782 int32
	_ = v3782
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L3
L1:
	;
	return v3782
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = v7
	goto L36
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 != v9 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v79 = F_slice_del(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	goto L4
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v14))))
	if v18 == int32(39) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L12
L9:
	;
	goto L8
L10:
	;
	if v72 < int32(0) {
		goto L2
	} else {
		goto L30
	}
L12:
	;
	goto L13
L13:
	;
	goto L14
L14:
	;
	v27 = v9
	v29 = int32(1)
	goto L17
L16:
	;
	v72 = v57
	goto L10
L17:
	;
	if v15 <= v27 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v72 = int32(-1)
	goto L10
L20:
	;
	goto L21
L21:
	;
	v34 = v27 + int32(1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v27))))
	if base.Ui32(v36) < base.Ui32(int32(192)) {
		v57 = v34
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(1)
	if v58 < v29 {
		v27 = v57
		v29 = v29 - v58
		goto L17
	} else {
		goto L29
	}
L23:
	;
	if v15 <= v34 {
		v57 = v34
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v43 = v34
	goto L25
L25:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+v43))))
	if int32(-65) < v46 {
		v57 = v43
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v57 = v15
	goto L22
L27:
	;
	v50 = v43 + int32(1)
	if v50 != v15 {
		v43 = v50
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L18
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	v9 = v72
	goto L3
L31:
	;
	return int32(0)
L32:
	;
	if v79 < int32(0) {
		v3782 = v79
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v204 < int32(0) {
		v3782 = int32(0)
		goto L1
	} else {
		goto L59
	}
L35:
	;
	v204 = v176
	goto L34
L36:
	;
	if v100 <= v109 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v204 = int32(-1)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v116 = int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v101))))
	if base.Ui32(v118) < base.Ui32(int32(192)) {
		v175 = v118
		v176 = v116
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if int32(305) < v175 {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v122 = v109 + int32(1)
	if v122 == v100 {
		v175 = v118
		v176 = v116
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v101))))
	v127 = v125 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v118) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v101))))
	v143 = v141 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v118) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v131 = v109 + int32(2)
	if v131 != v100 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
	v176 = int32(2)
	goto L41
L48:
	;
	goto L47
L49:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
	v175 = v160&int32(63) | (v118<<(uint(int32(18))%32)&int32(1835008) | v127<<(uint(int32(12))%32) | v143<<(uint(int32(6))%32))
	v176 = int32(4)
	goto L41
L50:
	;
	v147 = v109 + int32(3)
	if v147 != v100 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v175 = v118<<(uint(int32(12))%32)&int32(61440) | v127<<(uint(int32(6))%32) | v143
	v176 = int32(3)
	goto L41
L53:
	;
	goto L52
L54:
	;
	v193 = v176 + v109
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v193
	v109 = v193
	goto L36
L55:
	;
	v180 = v175 - int32(97)
	if v180 < int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_consts[1448]))))
	if int32(base.Ui32(v186)>>(uint(v180&int32(7))%32))&int32(1) != 0 {
		goto L35
	} else {
		goto L57
	}
L57:
	;
	goto L54
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v208 = v207 + v204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v208
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = v208
	goto L62
L60:
	;
	if v326 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L61:
	;
	v326 = v298
	goto L60
L62:
	;
	if v222 <= v231 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v326 = int32(-1)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v238 = int32(1)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v223))))
	if base.Ui32(v240) < base.Ui32(int32(192)) {
		v297 = v240
		v298 = v238
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if int32(305) < v297 {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v244 = v231 + int32(1)
	if v244 == v222 {
		v297 = v240
		v298 = v238
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v223))))
	v249 = v247 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v240) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v223))))
	v265 = v263 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v240) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v253 = v231 + int32(2)
	if v253 != v222 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v297 = v240<<(uint(int32(6))%32)&int32(1984) | v249
	v298 = int32(2)
	goto L67
L74:
	;
	goto L73
L75:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223+v269))))
	v297 = v282&int32(63) | (v240<<(uint(int32(18))%32)&int32(1835008) | v249<<(uint(int32(12))%32) | v265<<(uint(int32(6))%32))
	v298 = int32(4)
	goto L67
L76:
	;
	v269 = v231 + int32(3)
	if v269 != v222 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v297 = v240<<(uint(int32(12))%32)&int32(61440) | v249<<(uint(int32(6))%32) | v265
	v298 = int32(3)
	goto L67
L79:
	;
	goto L78
L80:
	;
	v315 = v298 + v231
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v315
	v231 = v315
	goto L62
L81:
	;
	v302 = v297 - int32(97)
	if v302 < int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v302)>>(uint(int32(3))%32)))+uint32(_consts[1448]))))
	if int32(base.Ui32(v308)>>(uint(v302&int32(7))%32))&int32(1) != 0 {
		goto L61
	} else {
		goto L83
	}
L83:
	;
	goto L80
L85:
	;
	return int32(0)
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v340 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v340 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1638
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	if v1641 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L89:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1627
	v1629 = F_slice_del(m, l0)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L31
	} else {
		goto L403
	}
L90:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v438 = v339 - v338
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437 - v438
	v441 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v441 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L91:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v344-int32(3) <= v343 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348+v344-int32(1)))))
	if v352 != int32(159) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v357 = F_find_among_b(m, l0, int32(4356112), int32(4))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L31
	} else {
		goto L94
	}
L94:
	;
	if v357 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v369 <= v370 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	if v431 != 0 {
		goto L89
	} else {
		goto L112
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v427
	v431 = int32(1)
	goto L97
L99:
	;
	v427 = v379 - v367 + v386
	goto L98
L100:
	;
	v394 = v369 - v367
	v395 = v393 + v394
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v395
	if v392 < v395 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v391 = v368
	v392 = v370
	v393 = v367
	goto L100
L102:
	;
	goto L103
L103:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369+v368-int32(1)))))
	if v375 != int32(121) {
		v391 = v368
		v392 = v370
		v393 = v367
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v379 = v369 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v384 = int32(0)
	v385 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v384)
	mBase = m.M
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v385 == v384 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v391 = v389
	v392 = v390
	v393 = v386
	goto L100
L106:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v391-int32(1)))))
	if v402 == int32(121) {
		v431 = int32(0)
		goto L97
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v406 = int32(0)
	v408 = F_skip_b_utf8(m, v391, v395, v392, int32(1))
	mBase = m.M
	if v408 < v406 {
		v431 = v406
		goto L97
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v408
	v416 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v416 != 0 {
		v431 = v406
		goto L97
	} else {
		goto L111
	}
L111:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v427 = v417 + v394
	goto L98
L112:
	;
	goto L90
L113:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v526 = v525 - v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v526
	v529 = v526 - int32(1)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v529 <= v530 {
		goto L134
	} else {
		goto L135
	}
L114:
	;
	v446 = F_find_among_b(m, l0, int32(4356192), int32(32))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L31
	} else {
		goto L115
	}
L115:
	;
	if v446 == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v458 <= v459 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	if v520 != 0 {
		goto L89
	} else {
		goto L133
	}
L118:
	;
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v516
	v520 = int32(1)
	goto L118
L120:
	;
	v516 = v468 - v456 + v475
	goto L119
L121:
	;
	v483 = v458 - v456
	v484 = v482 + v483
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v484
	if v481 < v484 {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v480 = v457
	v481 = v459
	v482 = v456
	goto L121
L123:
	;
	goto L124
L124:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v457-int32(1)))))
	if v464 != int32(121) {
		v480 = v457
		v481 = v459
		v482 = v456
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v468 = v458 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468
	v473 = int32(0)
	v474 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v473)
	mBase = m.M
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v474 == v473 {
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v480 = v478
	v481 = v479
	v482 = v475
	goto L121
L127:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v480-int32(1)))))
	if v491 == int32(121) {
		v520 = int32(0)
		goto L118
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v495 = int32(0)
	v497 = F_skip_b_utf8(m, v480, v484, v481, int32(1))
	mBase = m.M
	if v497 < v495 {
		v520 = v495
		goto L118
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v497
	v505 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v505 != 0 {
		v520 = v495
		goto L118
	} else {
		goto L132
	}
L132:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v516 = v506 + v483
	goto L119
L133:
	;
	goto L113
L134:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v628 = v627 - v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v628
	v630 = int32(3)
	v632 = int32(0)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v628-v635 < v630 {
		v645 = v632
		goto L158
	} else {
		goto L159
	}
L135:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532+v529))))
	if v534&int32(224) != int32(96) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	if int32(1)<<(uint(v534)%32)&int32(26658) == int32(0) {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v547 = F_find_among_b(m, l0, int32(4356832), int32(8))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L31
	} else {
		goto L138
	}
L138:
	;
	if v547 == int32(0) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v559 <= v560 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	if v621 != 0 {
		goto L89
	} else {
		goto L156
	}
L141:
	;
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v617
	v621 = int32(1)
	goto L141
L143:
	;
	v617 = v569 - v557 + v576
	goto L142
L144:
	;
	v584 = v559 - v557
	v585 = v583 + v584
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v585
	if v582 < v585 {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v581 = v558
	v582 = v560
	v583 = v557
	goto L144
L146:
	;
	goto L147
L147:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559+v558-int32(1)))))
	if v565 != int32(121) {
		v581 = v558
		v582 = v560
		v583 = v557
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v569 = v559 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v569
	v574 = int32(0)
	v575 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v574)
	mBase = m.M
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v575 == v574 {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v581 = v579
	v582 = v580
	v583 = v576
	goto L144
L150:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585+v581-int32(1)))))
	if v592 == int32(121) {
		v621 = int32(0)
		goto L141
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v596 = int32(0)
	v598 = F_skip_b_utf8(m, v581, v585, v582, int32(1))
	mBase = m.M
	if v598 < v596 {
		v621 = v596
		goto L141
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598
	v606 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v606 != 0 {
		v621 = v596
		goto L141
	} else {
		goto L155
	}
L155:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v617 = v607 + v584
	goto L142
L156:
	;
	goto L134
L157:
	;
	if v645 != 0 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L157
L159:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v641 = F_memcmp(m, v638+v628-v630, int32(2204114), v630)
	mBase = m.M
	if v641 != 0 {
		v645 = v632
		goto L158
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v628 - v630
	v645 = int32(1)
	goto L158
L161:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v654 <= v655 {
		goto L169
	} else {
		goto L170
	}
L162:
	;
	goto L163
L163:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v722 = v721 - v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v722-int32(5) <= v724 {
		goto L181
	} else {
		goto L182
	}
L164:
	;
	if v716 != 0 {
		goto L89
	} else {
		goto L180
	}
L165:
	;
	goto L164
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v712
	v716 = int32(1)
	goto L165
L167:
	;
	v712 = v664 - v652 + v671
	goto L166
L168:
	;
	v679 = v654 - v652
	v680 = v678 + v679
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v680
	if v677 < v680 {
		goto L174
	} else {
		goto L175
	}
L169:
	;
	v676 = v653
	v677 = v655
	v678 = v652
	goto L168
L170:
	;
	goto L171
L171:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+v653-int32(1)))))
	if v660 != int32(121) {
		v676 = v653
		v677 = v655
		v678 = v652
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v664 = v654 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
	v669 = int32(0)
	v670 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v669)
	mBase = m.M
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v670 == v669 {
		goto L167
	} else {
		goto L173
	}
L173:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v676 = v674
	v677 = v675
	v678 = v671
	goto L168
L174:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680+v676-int32(1)))))
	if v687 == int32(121) {
		v716 = int32(0)
		goto L165
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v691 = int32(0)
	v693 = F_skip_b_utf8(m, v676, v680, v677, int32(1))
	mBase = m.M
	if v693 < v691 {
		v716 = v691
		goto L165
	} else {
		goto L178
	}
L177:
	;
	goto L176
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v693
	v701 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v701 != 0 {
		v716 = v691
		goto L165
	} else {
		goto L179
	}
L179:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v712 = v702 + v679
	goto L166
L180:
	;
	goto L163
L181:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v995 - v438
	v998 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v998 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L182:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v722-int32(1)))))
	switch v732 - int32(97) {
	case 0, 4:
		goto L183
	default:
		goto L181
	}
L183:
	;
	v737 = F_find_among_b(m, l0, int32(4356992), int32(2))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L31
	} else {
		goto L184
	}
L184:
	;
	if v737 == int32(0) {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v743-int32(4) <= v742 {
		v759 = v741
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v896 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v896 == int32(0) {
		goto L181
	} else {
		goto L225
	}
L187:
	;
	v760 = v741 - v743
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v759 - v760
	v763 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v763 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747+v743-int32(1)))))
	if v751 != int32(122) {
		v759 = v741
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v756 = F_find_among_b(m, l0, int32(4357040), int32(4))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L31
	} else {
		goto L190
	}
L190:
	;
	if v756 != 0 {
		goto L186
	} else {
		goto L191
	}
L191:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v759 = v758
	goto L187
L192:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783 - v760
	v786 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v786 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v767-int32(2) <= v766 {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771+v767-int32(1)))))
	if v775 != int32(114) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	v780 = F_find_among_b(m, l0, int32(4357120), int32(2))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L31
	} else {
		goto L196
	}
L196:
	;
	if v780 != 0 {
		goto L186
	} else {
		goto L197
	}
L197:
	;
	goto L192
L198:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v881 - v760
	v884 = F_r_mark_sUn(m, l0)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L31
	} else {
		goto L221
	}
L199:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v791 = v789 - int32(1)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v791 <= v792 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794+v791))))
	if v796 != int32(109) {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v801 = F_find_among_b(m, l0, int32(4357168), int32(4))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L31
	} else {
		goto L202
	}
L202:
	;
	if v801 == int32(0) {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v813 <= v814 {
		goto L209
	} else {
		goto L210
	}
L204:
	;
	if v875 != 0 {
		goto L186
	} else {
		goto L220
	}
L205:
	;
	goto L204
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v871
	v875 = int32(1)
	goto L205
L207:
	;
	v871 = v823 - v811 + v830
	goto L206
L208:
	;
	v838 = v813 - v811
	v839 = v837 + v838
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v839
	if v836 < v839 {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v835 = v812
	v836 = v814
	v837 = v811
	goto L208
L210:
	;
	goto L211
L211:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+v812-int32(1)))))
	if v819 != int32(121) {
		v835 = v812
		v836 = v814
		v837 = v811
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v823 = v813 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v823
	v828 = int32(0)
	v829 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v828)
	mBase = m.M
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v829 == v828 {
		goto L207
	} else {
		goto L213
	}
L213:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v835 = v833
	v836 = v834
	v837 = v830
	goto L208
L214:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839+v835-int32(1)))))
	if v846 == int32(121) {
		v875 = int32(0)
		goto L205
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v850 = int32(0)
	v852 = F_skip_b_utf8(m, v835, v839, v836, int32(1))
	mBase = m.M
	if v852 < v850 {
		v875 = v850
		goto L205
	} else {
		goto L218
	}
L217:
	;
	goto L216
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v852
	v860 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v860 != 0 {
		v875 = v850
		goto L205
	} else {
		goto L219
	}
L219:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v871 = v861 + v838
	goto L206
L220:
	;
	goto L198
L221:
	;
	if v884 != 0 {
		goto L186
	} else {
		goto L222
	}
L222:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v886 - v760
	v889 = F_r_mark_yUz(m, l0)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L31
	} else {
		goto L223
	}
L223:
	;
	if v889 != 0 {
		goto L186
	} else {
		goto L224
	}
L224:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v891 - v760
	goto L186
L225:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v900-int32(3) <= v899 {
		goto L181
	} else {
		goto L226
	}
L226:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904+v900-int32(1)))))
	if v908 != int32(159) {
		goto L181
	} else {
		goto L227
	}
L227:
	;
	v913 = F_find_among_b(m, l0, int32(4356112), int32(4))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L31
	} else {
		goto L228
	}
L228:
	;
	if v913 == int32(0) {
		goto L181
	} else {
		goto L229
	}
L229:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v925 <= v926 {
		goto L235
	} else {
		goto L236
	}
L230:
	;
	if v987 != 0 {
		goto L89
	} else {
		goto L246
	}
L231:
	;
	goto L230
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v983
	v987 = int32(1)
	goto L231
L233:
	;
	v983 = v935 - v923 + v942
	goto L232
L234:
	;
	v950 = v925 - v923
	v951 = v949 + v950
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951
	if v948 < v951 {
		goto L240
	} else {
		goto L241
	}
L235:
	;
	v947 = v924
	v948 = v926
	v949 = v923
	goto L234
L236:
	;
	goto L237
L237:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925+v924-int32(1)))))
	if v931 != int32(121) {
		v947 = v924
		v948 = v926
		v949 = v923
		goto L234
	} else {
		goto L238
	}
L238:
	;
	v935 = v925 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935
	v940 = int32(0)
	v941 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v940)
	mBase = m.M
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v941 == v940 {
		goto L233
	} else {
		goto L239
	}
L239:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v947 = v945
	v948 = v946
	v949 = v942
	goto L234
L240:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951+v947-int32(1)))))
	if v958 == int32(121) {
		v987 = int32(0)
		goto L231
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v962 = int32(0)
	v964 = F_skip_b_utf8(m, v947, v951, v948, int32(1))
	mBase = m.M
	if v964 < v962 {
		v987 = v962
		goto L231
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v964
	v972 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v972 != 0 {
		v987 = v962
		goto L231
	} else {
		goto L245
	}
L245:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v983 = v973 + v950
	goto L232
L246:
	;
	goto L181
L247:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1158 - v438
	v1161 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1161 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L248:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1002-int32(2) <= v1001 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006+v1002-int32(1)))))
	if v1010 != int32(114) {
		goto L247
	} else {
		goto L250
	}
L250:
	;
	v1015 = F_find_among_b(m, l0, int32(4357120), int32(2))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L31
	} else {
		goto L251
	}
L251:
	;
	if v1015 == int32(0) {
		goto L247
	} else {
		goto L252
	}
L252:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1019
	v1021 = F_slice_del(m, l0)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L31
	} else {
		goto L253
	}
L253:
	;
	if v1021 < int32(0) {
		v3782 = v1021
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1025
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1028 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1028 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1154))) = int32(0)
	goto L89
L256:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1049 = v1027 - v1025
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1048 - v1049
	v1052 = int32(0)
	v1053 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1053 == v1052 {
		v1137 = v1052
		goto L262
	} else {
		goto L263
	}
L257:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1032-int32(2) <= v1031 {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036+v1032-int32(1)))))
	if v1040 != int32(114) {
		goto L256
	} else {
		goto L259
	}
L259:
	;
	v1045 = F_find_among_b(m, l0, int32(4357408), int32(8))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L31
	} else {
		goto L260
	}
L260:
	;
	if v1045 != 0 {
		goto L255
	} else {
		goto L261
	}
L261:
	;
	goto L256
L262:
	;
	if v1137 != 0 {
		goto L255
	} else {
		goto L282
	}
L263:
	;
	v1058 = F_find_among_b(m, l0, int32(4356192), int32(32))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L31
	} else {
		goto L264
	}
L264:
	;
	if v1058 == int32(0) {
		v1137 = v1052
		goto L262
	} else {
		goto L265
	}
L265:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1070 <= v1071 {
		goto L271
	} else {
		goto L272
	}
L266:
	;
	v1137 = v1132
	goto L262
L267:
	;
	goto L266
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1128
	v1132 = int32(1)
	goto L267
L269:
	;
	v1128 = v1080 - v1068 + v1087
	goto L268
L270:
	;
	v1095 = v1070 - v1068
	v1096 = v1094 + v1095
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1096
	if v1093 < v1096 {
		goto L276
	} else {
		goto L277
	}
L271:
	;
	v1092 = v1069
	v1093 = v1071
	v1094 = v1068
	goto L270
L272:
	;
	goto L273
L273:
	;
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070+v1069-int32(1)))))
	if v1076 != int32(121) {
		v1092 = v1069
		v1093 = v1071
		v1094 = v1068
		goto L270
	} else {
		goto L274
	}
L274:
	;
	v1080 = v1070 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1080
	v1085 = int32(0)
	v1086 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1085)
	mBase = m.M
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1086 == v1085 {
		goto L269
	} else {
		goto L275
	}
L275:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1092 = v1090
	v1093 = v1091
	v1094 = v1087
	goto L270
L276:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096+v1092-int32(1)))))
	if v1103 == int32(121) {
		v1132 = int32(0)
		goto L267
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1107 = int32(0)
	v1109 = F_skip_b_utf8(m, v1092, v1096, v1093, int32(1))
	mBase = m.M
	if v1109 < v1107 {
		v1132 = v1107
		goto L267
	} else {
		goto L280
	}
L279:
	;
	goto L278
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1109
	v1117 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1117 != 0 {
		v1132 = v1107
		goto L267
	} else {
		goto L281
	}
L281:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1128 = v1118 + v1095
	goto L268
L282:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1138 - v1049
	v1141 = F_r_mark_ysA(m, l0)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L31
	} else {
		goto L283
	}
L283:
	;
	if v1141 != 0 {
		goto L255
	} else {
		goto L284
	}
L284:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1143 - v1049
	v1146 = F_r_mark_ymUs_(m, l0)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L31
	} else {
		goto L285
	}
L285:
	;
	if v1146 != 0 {
		goto L255
	} else {
		goto L286
	}
L286:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1148 + (v1025 - v1027)
	goto L255
L287:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1277 = v1276 - v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1277
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1277-int32(4) <= v1279 {
		v1296 = v1277
		goto L318
	} else {
		goto L319
	}
L288:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1165-int32(2) <= v1164 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169+v1165-int32(1)))))
	if v1173 != int32(122) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	v1178 = F_find_among_b(m, l0, int32(4357568), int32(4))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L31
	} else {
		goto L291
	}
L291:
	;
	if v1178 == int32(0) {
		goto L287
	} else {
		goto L292
	}
L292:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1184 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1184 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1268 + (v1183 - v1182)
	v1272 = F_r_mark_ysA(m, l0)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L31
	} else {
		goto L314
	}
L294:
	;
	v1189 = F_find_among_b(m, l0, int32(4356192), int32(32))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L31
	} else {
		goto L295
	}
L295:
	;
	if v1189 == int32(0) {
		goto L293
	} else {
		goto L296
	}
L296:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1201 <= v1202 {
		goto L302
	} else {
		goto L303
	}
L297:
	;
	if v1263 != 0 {
		goto L89
	} else {
		goto L313
	}
L298:
	;
	goto L297
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259
	v1263 = int32(1)
	goto L298
L300:
	;
	v1259 = v1211 - v1199 + v1218
	goto L299
L301:
	;
	v1226 = v1201 - v1199
	v1227 = v1225 + v1226
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1227
	if v1224 < v1227 {
		goto L307
	} else {
		goto L308
	}
L302:
	;
	v1223 = v1200
	v1224 = v1202
	v1225 = v1199
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201+v1200-int32(1)))))
	if v1207 != int32(121) {
		v1223 = v1200
		v1224 = v1202
		v1225 = v1199
		goto L301
	} else {
		goto L305
	}
L305:
	;
	v1211 = v1201 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1211
	v1216 = int32(0)
	v1217 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1216)
	mBase = m.M
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1217 == v1216 {
		goto L300
	} else {
		goto L306
	}
L306:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1223 = v1221
	v1224 = v1222
	v1225 = v1218
	goto L301
L307:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227+v1223-int32(1)))))
	if v1234 == int32(121) {
		v1263 = int32(0)
		goto L298
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1238 = int32(0)
	v1240 = F_skip_b_utf8(m, v1223, v1227, v1224, int32(1))
	mBase = m.M
	if v1240 < v1238 {
		v1263 = v1238
		goto L298
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1240
	v1248 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1248 != 0 {
		v1263 = v1238
		goto L298
	} else {
		goto L312
	}
L312:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1259 = v1249 + v1226
	goto L299
L313:
	;
	goto L293
L314:
	;
	if v1272 != 0 {
		goto L89
	} else {
		goto L315
	}
L315:
	;
	goto L287
L316:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1534 - v438
	v1537 = int32(0)
	v1538 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1538 == v1537 {
		v1559 = v1537
		goto L379
	} else {
		goto L380
	}
L317:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1424
	v1426 = F_slice_del(m, l0)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L31
	} else {
		goto L354
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1296
	v1298 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1298 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L319:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1283+v1277-int32(1)))))
	if v1287 != int32(122) {
		v1296 = v1277
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1292 = F_find_among_b(m, l0, int32(4357040), int32(4))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L31
	} else {
		goto L321
	}
L321:
	;
	if v1292 != 0 {
		goto L317
	} else {
		goto L322
	}
L322:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1296 = v1294 - v438
	goto L318
L323:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1393 - v438
	v1396 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1396 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L324:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1303 = v1301 - int32(1)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1303 <= v1304 {
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306+v1303))))
	if v1308 != int32(122) {
		goto L323
	} else {
		goto L326
	}
L326:
	;
	v1313 = F_find_among_b(m, l0, int32(4357328), int32(4))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L31
	} else {
		goto L327
	}
L327:
	;
	if v1313 == int32(0) {
		goto L323
	} else {
		goto L328
	}
L328:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1325 <= v1326 {
		goto L334
	} else {
		goto L335
	}
L329:
	;
	if v1387 != 0 {
		goto L317
	} else {
		goto L345
	}
L330:
	;
	goto L329
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1383
	v1387 = int32(1)
	goto L330
L332:
	;
	v1383 = v1335 - v1323 + v1342
	goto L331
L333:
	;
	v1350 = v1325 - v1323
	v1351 = v1349 + v1350
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1351
	if v1348 < v1351 {
		goto L339
	} else {
		goto L340
	}
L334:
	;
	v1347 = v1324
	v1348 = v1326
	v1349 = v1323
	goto L333
L335:
	;
	goto L336
L336:
	;
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1325+v1324-int32(1)))))
	if v1331 != int32(121) {
		v1347 = v1324
		v1348 = v1326
		v1349 = v1323
		goto L333
	} else {
		goto L337
	}
L337:
	;
	v1335 = v1325 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1335
	v1340 = int32(0)
	v1341 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1340)
	mBase = m.M
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1341 == v1340 {
		goto L332
	} else {
		goto L338
	}
L338:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1347 = v1345
	v1348 = v1346
	v1349 = v1342
	goto L333
L339:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1347-int32(1)))))
	if v1358 == int32(121) {
		v1387 = int32(0)
		goto L330
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1362 = int32(0)
	v1364 = F_skip_b_utf8(m, v1347, v1351, v1348, int32(1))
	mBase = m.M
	if v1364 < v1362 {
		v1387 = v1362
		goto L330
	} else {
		goto L343
	}
L342:
	;
	goto L341
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1364
	v1372 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1372 != 0 {
		v1387 = v1362
		goto L330
	} else {
		goto L344
	}
L344:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1383 = v1373 + v1350
	goto L331
L345:
	;
	goto L323
L346:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1416 - v438
	v1419 = F_r_mark_yUm(m, l0)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L31
	} else {
		goto L352
	}
L347:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1400-int32(2) <= v1399 {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404+v1400-int32(1)))))
	if v1408 != int32(110) {
		goto L346
	} else {
		goto L349
	}
L349:
	;
	v1413 = F_find_among_b(m, l0, int32(4357248), int32(4))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L31
	} else {
		goto L350
	}
L350:
	;
	if v1413 != 0 {
		goto L317
	} else {
		goto L351
	}
L351:
	;
	goto L346
L352:
	;
	if v1419 == int32(0) {
		goto L316
	} else {
		goto L353
	}
L353:
	;
	goto L317
L354:
	;
	if v1426 < int32(0) {
		v3782 = v1426
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1430
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1433 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1433 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1530 + (v1430 - v1432)
	goto L89
L357:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1437-int32(3) <= v1436 {
		goto L356
	} else {
		goto L358
	}
L358:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441+v1437-int32(1)))))
	if v1445 != int32(159) {
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v1450 = F_find_among_b(m, l0, int32(4356112), int32(4))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L31
	} else {
		goto L360
	}
L360:
	;
	if v1450 == int32(0) {
		goto L356
	} else {
		goto L361
	}
L361:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1462 <= v1463 {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	if v1524 != 0 {
		goto L89
	} else {
		goto L378
	}
L363:
	;
	goto L362
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1520
	v1524 = int32(1)
	goto L363
L365:
	;
	v1520 = v1472 - v1460 + v1479
	goto L364
L366:
	;
	v1487 = v1462 - v1460
	v1488 = v1486 + v1487
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1488
	if v1485 < v1488 {
		goto L372
	} else {
		goto L373
	}
L367:
	;
	v1484 = v1461
	v1485 = v1463
	v1486 = v1460
	goto L366
L368:
	;
	goto L369
L369:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1462+v1461-int32(1)))))
	if v1468 != int32(121) {
		v1484 = v1461
		v1485 = v1463
		v1486 = v1460
		goto L366
	} else {
		goto L370
	}
L370:
	;
	v1472 = v1462 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1472
	v1477 = int32(0)
	v1478 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1477)
	mBase = m.M
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1478 == v1477 {
		goto L365
	} else {
		goto L371
	}
L371:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1484 = v1482
	v1485 = v1483
	v1486 = v1479
	goto L366
L372:
	;
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1488+v1484-int32(1)))))
	if v1495 == int32(121) {
		v1524 = int32(0)
		goto L363
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v1499 = int32(0)
	v1501 = F_skip_b_utf8(m, v1484, v1488, v1485, int32(1))
	mBase = m.M
	if v1501 < v1499 {
		v1524 = v1499
		goto L363
	} else {
		goto L376
	}
L375:
	;
	goto L374
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1501
	v1509 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1509 != 0 {
		v1524 = v1499
		goto L363
	} else {
		goto L377
	}
L377:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1520 = v1510 + v1487
	goto L364
L378:
	;
	goto L356
L379:
	;
	if v1559 == int32(0) {
		goto L88
	} else {
		goto L384
	}
L380:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1542-int32(2) <= v1541 {
		v1559 = v1537
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546+v1542-int32(1)))))
	if v1550 != int32(114) {
		v1559 = v1537
		goto L379
	} else {
		goto L382
	}
L382:
	;
	v1555 = F_find_among_b(m, l0, int32(4357408), int32(8))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L31
	} else {
		goto L383
	}
L383:
	;
	v1559 = base.B2i32(v1555 != int32(0))
	goto L379
L384:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1563
	v1565 = F_slice_del(m, l0)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L31
	} else {
		goto L385
	}
L385:
	;
	if v1565 < int32(0) {
		v3782 = v1565
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1569
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1569-int32(4) <= v1572 {
		v1590 = v2
		goto L387
	} else {
		goto L388
	}
L387:
	;
	if v1590 != 0 {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577+v1569-int32(1)))))
	if v1581 != int32(122) {
		v1590 = v2
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1586 = F_find_among_b(m, l0, int32(4357040), int32(4))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L31
	} else {
		goto L390
	}
L390:
	;
	v1590 = base.B2i32(v1586 != int32(0))
	goto L387
L391:
	;
	v1616 = F_r_mark_ymUs_(m, l0)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L31
	} else {
		goto L401
	}
L392:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1592 = v1571 - v1569
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1591 - v1592
	v1595 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L31
	} else {
		goto L393
	}
L393:
	;
	if v1595 != 0 {
		goto L391
	} else {
		goto L394
	}
L394:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1597 - v1592
	v1600 = F_r_mark_yUm(m, l0)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L31
	} else {
		goto L395
	}
L395:
	;
	if v1600 != 0 {
		goto L391
	} else {
		goto L396
	}
L396:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1602 - v1592
	v1605 = F_r_mark_sUn(m, l0)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L31
	} else {
		goto L397
	}
L397:
	;
	if v1605 != 0 {
		goto L391
	} else {
		goto L398
	}
L398:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1607 - v1592
	v1610 = F_r_mark_yUz(m, l0)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L31
	} else {
		goto L399
	}
L399:
	;
	if v1610 != 0 {
		goto L391
	} else {
		goto L400
	}
L400:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1612 - v1592
	goto L391
L401:
	;
	if v1616 != 0 {
		goto L89
	} else {
		goto L402
	}
L402:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1618 + (v1569 - v1571)
	goto L89
L403:
	;
	if v1629 < int32(0) {
		v3782 = v1629
		goto L1
	} else {
		goto L404
	}
L404:
	;
	goto L88
L405:
	;
	return int32(0)
L406:
	;
	goto L407
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1638
	v1647 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1647 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3416
	v3418 = int32(2)
	v3420 = int32(0)
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3416-v3423 < v3418 {
		v3433 = v3420
		goto L927
	} else {
		goto L928
	}
L409:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1687
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1687
	v1690 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1690 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L410:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1651-int32(2) <= v1650 {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1655+v1651-int32(1)))))
	if v1659 != int32(114) {
		goto L409
	} else {
		goto L412
	}
L412:
	;
	v1664 = F_find_among_b(m, l0, int32(4357120), int32(2))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L31
	} else {
		goto L413
	}
L413:
	;
	if v1664 == int32(0) {
		goto L409
	} else {
		goto L414
	}
L414:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1668
	v1670 = F_slice_del(m, l0)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L31
	} else {
		goto L415
	}
L415:
	;
	if v1670 < int32(0) {
		v3782 = v1670
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1676 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L31
	} else {
		goto L417
	}
L417:
	;
	if v1676 == int32(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1680 + (v1674 - v1675)
	goto L408
L419:
	;
	goto L420
L420:
	;
	if int32(0) <= v1676 {
		goto L408
	} else {
		goto L421
	}
L421:
	;
	v3782 = v1676
	goto L1
L422:
	;
	v3405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3405
	v3407 = F_slice_del(m, l0)
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L31
	} else {
		goto L924
	}
L423:
	;
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3399
	v3401 = F_slice_del(m, l0)
	mBase = m.M
	v3402 = m.ExcPending
	if v3402 != 0 {
		goto L31
	} else {
		goto L922
	}
L424:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1974
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1974
	v1977 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1977 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L425:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1695 = v1693 - int32(1)
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1695 <= v1696 {
		goto L424
	} else {
		goto L426
	}
L426:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1698+v1695))))
	switch v1700 - int32(97) {
	case 0, 4:
		goto L427
	default:
		goto L424
	}
L427:
	;
	v1705 = F_find_among_b(m, l0, int32(4357648), int32(2))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L31
	} else {
		goto L428
	}
L428:
	;
	if v1705 == int32(0) {
		goto L424
	} else {
		goto L429
	}
L429:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1717 <= v1718 {
		goto L435
	} else {
		goto L436
	}
L430:
	;
	if v1779 == int32(0) {
		goto L424
	} else {
		goto L446
	}
L431:
	;
	goto L430
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1775
	v1779 = int32(1)
	goto L431
L433:
	;
	v1775 = v1727 - v1715 + v1734
	goto L432
L434:
	;
	v1742 = v1717 - v1715
	v1743 = v1741 + v1742
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1743
	if v1740 < v1743 {
		goto L440
	} else {
		goto L441
	}
L435:
	;
	v1739 = v1716
	v1740 = v1718
	v1741 = v1715
	goto L434
L436:
	;
	goto L437
L437:
	;
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1717+v1716-int32(1)))))
	if v1723 != int32(110) {
		v1739 = v1716
		v1740 = v1718
		v1741 = v1715
		goto L434
	} else {
		goto L438
	}
L438:
	;
	v1727 = v1717 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1727
	v1732 = int32(0)
	v1733 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1732)
	mBase = m.M
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1733 == v1732 {
		goto L433
	} else {
		goto L439
	}
L439:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1739 = v1737
	v1740 = v1738
	v1741 = v1734
	goto L434
L440:
	;
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1739-int32(1)))))
	if v1750 == int32(110) {
		v1779 = int32(0)
		goto L431
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	v1754 = int32(0)
	v1756 = F_skip_b_utf8(m, v1739, v1743, v1740, int32(1))
	mBase = m.M
	if v1756 < v1754 {
		v1779 = v1754
		goto L431
	} else {
		goto L444
	}
L443:
	;
	goto L442
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1756
	v1764 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1764 != 0 {
		v1779 = v1754
		goto L431
	} else {
		goto L445
	}
L445:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1775 = v1765 + v1742
	goto L432
L446:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1786
	v1788 = F_slice_del(m, l0)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L31
	} else {
		goto L447
	}
L447:
	;
	if v1788 < int32(0) {
		v3782 = v1788
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1792
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1792-int32(3) <= v1795 {
		v1814 = v1794
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1815 = v1794 - v1792
	v1816 = v1814 - v1815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1816
	v1819 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L31
	} else {
		goto L458
	}
L450:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1799+v1792-int32(1)))))
	if v1803 != int32(177) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	if v1803 != int32(105) {
		v1814 = v1794
		goto L449
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1810 = F_find_among_b(m, l0, int32(4357696), int32(2))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L31
	} else {
		goto L455
	}
L454:
	;
	goto L453
L455:
	;
	if v1810 != 0 {
		goto L423
	} else {
		goto L456
	}
L456:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1814 = v1812
	goto L449
L457:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1948 = v1947 - v1815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1948
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1948
	v1951 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L31
	} else {
		goto L496
	}
L458:
	;
	if v1819 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1823 - v1815
	v1826 = int32(0)
	v1832 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1832 == v1826 {
		goto L463
	} else {
		goto L464
	}
L460:
	;
	goto L461
L461:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1915
	v1917 = F_slice_del(m, l0)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L31
	} else {
		goto L483
	}
L462:
	;
	if v1912 == int32(0) {
		goto L457
	} else {
		goto L482
	}
L463:
	;
	v1912 = int32(0)
	goto L462
L464:
	;
	goto L465
L465:
	;
	v1840 = F_in_grouping_b_U(m, l0, int32(2204288), int32(105), int32(305), int32(0))
	mBase = m.M
	if v1840 != 0 {
		v1904 = v1826
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v1912 = v1904
	goto L462
L467:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1843 <= v1844 {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1897
	v1904 = int32(1)
	goto L466
L469:
	;
	v1897 = v1853 - v1841 + v1860
	goto L468
L470:
	;
	v1868 = v1843 - v1841
	v1869 = v1867 + v1868
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1869
	if v1866 < v1869 {
		goto L476
	} else {
		goto L477
	}
L471:
	;
	v1865 = v1842
	v1866 = v1844
	v1867 = v1841
	goto L470
L472:
	;
	goto L473
L473:
	;
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842+v1843-int32(1)))))
	if v1849 != int32(115) {
		v1865 = v1842
		v1866 = v1844
		v1867 = v1841
		goto L470
	} else {
		goto L474
	}
L474:
	;
	v1853 = v1843 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1853
	v1858 = int32(0)
	v1859 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v1858)
	mBase = m.M
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1859 == v1858 {
		goto L469
	} else {
		goto L475
	}
L475:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1865 = v1863
	v1866 = v1864
	v1867 = v1860
	goto L470
L476:
	;
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1869+v1865-int32(1)))))
	if v1875 == int32(115) {
		v1904 = v1826
		goto L466
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v1879 = F_skip_b_utf8(m, v1865, v1869, v1866, int32(1))
	mBase = m.M
	if v1879 < int32(0) {
		v1904 = v1826
		goto L466
	} else {
		goto L480
	}
L479:
	;
	goto L478
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1879
	v1887 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1887 != 0 {
		v1904 = v1826
		goto L466
	} else {
		goto L481
	}
L481:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1897 = v1888 + v1868
	goto L468
L482:
	;
	goto L461
L483:
	;
	if v1917 < int32(0) {
		v3782 = v1917
		goto L1
	} else {
		goto L484
	}
L484:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1921
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1924 = v1923 - v1921
	v1925 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L31
	} else {
		goto L485
	}
L485:
	;
	if v1925 == int32(0) {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1929 - v1924
	goto L408
L487:
	;
	goto L488
L488:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1932
	v1934 = F_slice_del(m, l0)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L31
	} else {
		goto L489
	}
L489:
	;
	if v1934 < int32(0) {
		v3782 = v1934
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v1938 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L31
	} else {
		goto L491
	}
L491:
	;
	if v1938 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1942 - v1924
	goto L408
L493:
	;
	goto L494
L494:
	;
	if int32(0) <= v1938 {
		goto L408
	} else {
		goto L495
	}
L495:
	;
	v3782 = v1938
	goto L1
L496:
	;
	if v1951 == int32(0) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1955 - v1815
	goto L408
L498:
	;
	goto L499
L499:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1958
	v1960 = F_slice_del(m, l0)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L31
	} else {
		goto L500
	}
L500:
	;
	if v1960 < int32(0) {
		v3782 = v1960
		goto L1
	} else {
		goto L501
	}
L501:
	;
	v1964 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L31
	} else {
		goto L502
	}
L502:
	;
	if v1964 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1968 - v1815
	goto L408
L504:
	;
	goto L505
L505:
	;
	if int32(0) <= v1964 {
		goto L408
	} else {
		goto L506
	}
L506:
	;
	v3782 = v1964
	goto L1
L507:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2175
	v2178 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2178 == int32(0) {
		goto L569
	} else {
		goto L570
	}
L508:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2021-int32(3) <= v2020 {
		v2040 = v2019
		goto L520
	} else {
		goto L521
	}
L509:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1997
	v1999 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1999 == int32(0) {
		goto L507
	} else {
		goto L515
	}
L510:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1981-int32(2) <= v1980 {
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1985+v1981-int32(1)))))
	switch v1989 - int32(97) {
	case 0, 4:
		goto L512
	default:
		goto L509
	}
L512:
	;
	v1994 = F_find_among_b(m, l0, int32(4357952), int32(2))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L31
	} else {
		goto L513
	}
L513:
	;
	if v1994 != 0 {
		goto L508
	} else {
		goto L514
	}
L514:
	;
	goto L509
L515:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2004 = v2002 - int32(1)
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2004 <= v2005 {
		goto L507
	} else {
		goto L516
	}
L516:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2007+v2004))))
	switch v2009 - int32(97) {
	case 0, 4:
		goto L517
	default:
		goto L507
	}
L517:
	;
	v2014 = F_find_among_b(m, l0, int32(4358000), int32(2))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L31
	} else {
		goto L518
	}
L518:
	;
	if v2014 == int32(0) {
		goto L507
	} else {
		goto L519
	}
L519:
	;
	goto L508
L520:
	;
	v2041 = v2019 - v2021
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2040 - v2041
	v2044 = int32(0)
	v2050 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2050 == v2044 {
		goto L529
	} else {
		goto L530
	}
L521:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2025+v2021-int32(1)))))
	if v2029 != int32(177) {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	if v2029 != int32(105) {
		v2040 = v2019
		goto L520
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	v2036 = F_find_among_b(m, l0, int32(4357696), int32(2))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L31
	} else {
		goto L526
	}
L525:
	;
	goto L524
L526:
	;
	if v2036 != 0 {
		goto L422
	} else {
		goto L527
	}
L527:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2040 = v2038
	goto L520
L528:
	;
	if v2130 != 0 {
		goto L548
	} else {
		goto L549
	}
L529:
	;
	v2130 = int32(0)
	goto L528
L530:
	;
	goto L531
L531:
	;
	v2058 = F_in_grouping_b_U(m, l0, int32(2204288), int32(105), int32(305), int32(0))
	mBase = m.M
	if v2058 != 0 {
		v2122 = v2044
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v2130 = v2122
	goto L528
L533:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2061 <= v2062 {
		goto L537
	} else {
		goto L538
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2115
	v2122 = int32(1)
	goto L532
L535:
	;
	v2115 = v2071 - v2059 + v2078
	goto L534
L536:
	;
	v2086 = v2061 - v2059
	v2087 = v2085 + v2086
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2087
	if v2084 < v2087 {
		goto L542
	} else {
		goto L543
	}
L537:
	;
	v2083 = v2060
	v2084 = v2062
	v2085 = v2059
	goto L536
L538:
	;
	goto L539
L539:
	;
	v2067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2060+v2061-int32(1)))))
	if v2067 != int32(115) {
		v2083 = v2060
		v2084 = v2062
		v2085 = v2059
		goto L536
	} else {
		goto L540
	}
L540:
	;
	v2071 = v2061 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2071
	v2076 = int32(0)
	v2077 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v2076)
	mBase = m.M
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2077 == v2076 {
		goto L535
	} else {
		goto L541
	}
L541:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2083 = v2081
	v2084 = v2082
	v2085 = v2078
	goto L536
L542:
	;
	v2093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2087+v2083-int32(1)))))
	if v2093 == int32(115) {
		v2122 = v2044
		goto L532
	} else {
		goto L545
	}
L543:
	;
	goto L544
L544:
	;
	v2097 = F_skip_b_utf8(m, v2083, v2087, v2084, int32(1))
	mBase = m.M
	if v2097 < int32(0) {
		v2122 = v2044
		goto L532
	} else {
		goto L546
	}
L545:
	;
	goto L544
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2097
	v2105 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2105 != 0 {
		v2122 = v2044
		goto L532
	} else {
		goto L547
	}
L547:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2115 = v2106 + v2086
	goto L534
L548:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2131
	v2133 = F_slice_del(m, l0)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L31
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2163 - v2041
	v2166 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L31
	} else {
		goto L564
	}
L551:
	;
	if v2133 < int32(0) {
		v3782 = v2133
		goto L1
	} else {
		goto L552
	}
L552:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2137
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2140 = v2139 - v2137
	v2141 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L31
	} else {
		goto L553
	}
L553:
	;
	if v2141 == int32(0) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2145 - v2140
	goto L408
L555:
	;
	goto L556
L556:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2148
	v2150 = F_slice_del(m, l0)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L31
	} else {
		goto L557
	}
L557:
	;
	if v2150 < int32(0) {
		v3782 = v2150
		goto L1
	} else {
		goto L558
	}
L558:
	;
	v2154 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L31
	} else {
		goto L559
	}
L559:
	;
	if v2154 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2158 - v2140
	goto L408
L561:
	;
	goto L562
L562:
	;
	if int32(0) <= v2154 {
		goto L408
	} else {
		goto L563
	}
L563:
	;
	v3782 = v2154
	goto L1
L564:
	;
	if v2166 == int32(0) {
		goto L507
	} else {
		goto L565
	}
L565:
	;
	if int32(0) <= v2166 {
		goto L408
	} else {
		goto L566
	}
L566:
	;
	v3782 = v2166
	goto L1
L567:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2339
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2339
	v2342 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2342 == int32(0) {
		goto L616
	} else {
		goto L617
	}
L568:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2212 = int32(0)
	v2218 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2218 == v2212 {
		goto L579
	} else {
		goto L580
	}
L569:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2198
	v2200 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2200 == int32(0) {
		goto L567
	} else {
		goto L575
	}
L570:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2182-int32(3) <= v2181 {
		goto L569
	} else {
		goto L571
	}
L571:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2186+v2182-int32(1)))))
	if v2190 != int32(110) {
		goto L569
	} else {
		goto L572
	}
L572:
	;
	v2195 = F_find_among_b(m, l0, int32(4358048), int32(2))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L31
	} else {
		goto L573
	}
L573:
	;
	if v2195 != 0 {
		goto L568
	} else {
		goto L574
	}
L574:
	;
	goto L569
L575:
	;
	v2205 = F_find_among_b(m, l0, int32(4358096), int32(4))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L31
	} else {
		goto L576
	}
L576:
	;
	if v2205 == int32(0) {
		goto L567
	} else {
		goto L577
	}
L577:
	;
	goto L568
L578:
	;
	if v2298 != 0 {
		goto L598
	} else {
		goto L599
	}
L579:
	;
	v2298 = int32(0)
	goto L578
L580:
	;
	goto L581
L581:
	;
	v2226 = F_in_grouping_b_U(m, l0, int32(2204288), int32(105), int32(305), int32(0))
	mBase = m.M
	if v2226 != 0 {
		v2290 = v2212
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v2298 = v2290
	goto L578
L583:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2229 <= v2230 {
		goto L587
	} else {
		goto L588
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2283
	v2290 = int32(1)
	goto L582
L585:
	;
	v2283 = v2239 - v2227 + v2246
	goto L584
L586:
	;
	v2254 = v2229 - v2227
	v2255 = v2253 + v2254
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2255
	if v2252 < v2255 {
		goto L592
	} else {
		goto L593
	}
L587:
	;
	v2251 = v2228
	v2252 = v2230
	v2253 = v2227
	goto L586
L588:
	;
	goto L589
L589:
	;
	v2235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228+v2229-int32(1)))))
	if v2235 != int32(115) {
		v2251 = v2228
		v2252 = v2230
		v2253 = v2227
		goto L586
	} else {
		goto L590
	}
L590:
	;
	v2239 = v2229 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2239
	v2244 = int32(0)
	v2245 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v2244)
	mBase = m.M
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2245 == v2244 {
		goto L585
	} else {
		goto L591
	}
L591:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2251 = v2249
	v2252 = v2250
	v2253 = v2246
	goto L586
L592:
	;
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255+v2251-int32(1)))))
	if v2261 == int32(115) {
		v2290 = v2212
		goto L582
	} else {
		goto L595
	}
L593:
	;
	goto L594
L594:
	;
	v2265 = F_skip_b_utf8(m, v2251, v2255, v2252, int32(1))
	mBase = m.M
	if v2265 < int32(0) {
		v2290 = v2212
		goto L582
	} else {
		goto L596
	}
L595:
	;
	goto L594
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2265
	v2273 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2273 != 0 {
		v2290 = v2212
		goto L582
	} else {
		goto L597
	}
L597:
	;
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2283 = v2274 + v2254
	goto L584
L598:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2299
	v2301 = F_slice_del(m, l0)
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L31
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2331 + (v2210 - v2211)
	v2335 = F_r_mark_lArI(m, l0)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L31
	} else {
		goto L614
	}
L601:
	;
	if v2301 < int32(0) {
		v3782 = v2301
		goto L1
	} else {
		goto L602
	}
L602:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2305
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2308 = v2307 - v2305
	v2309 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L31
	} else {
		goto L603
	}
L603:
	;
	if v2309 == int32(0) {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2313 - v2308
	goto L408
L605:
	;
	goto L606
L606:
	;
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2316
	v2318 = F_slice_del(m, l0)
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L31
	} else {
		goto L607
	}
L607:
	;
	if v2318 < int32(0) {
		v3782 = v2318
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v2322 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L31
	} else {
		goto L609
	}
L609:
	;
	if v2322 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2326 - v2308
	goto L408
L611:
	;
	goto L612
L612:
	;
	if int32(0) <= v2322 {
		goto L408
	} else {
		goto L613
	}
L613:
	;
	v3782 = v2322
	goto L1
L614:
	;
	if v2335 != 0 {
		goto L408
	} else {
		goto L615
	}
L615:
	;
	goto L567
L616:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2444
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2444
	v2447 = int32(0)
	v2448 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2448 == v2447 {
		v2542 = v2447
		goto L657
	} else {
		goto L658
	}
L617:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2346-int32(2) <= v2345 {
		goto L616
	} else {
		goto L618
	}
L618:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2350+v2346-int32(1)))))
	if v2354 != int32(110) {
		goto L616
	} else {
		goto L619
	}
L619:
	;
	v2359 = F_find_among_b(m, l0, int32(4358176), int32(4))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L31
	} else {
		goto L620
	}
L620:
	;
	if v2359 == int32(0) {
		goto L616
	} else {
		goto L621
	}
L621:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2363
	v2365 = F_slice_del(m, l0)
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L31
	} else {
		goto L622
	}
L622:
	;
	if v2365 < int32(0) {
		v3782 = v2365
		goto L1
	} else {
		goto L623
	}
L623:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2369
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2372 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L31
	} else {
		goto L624
	}
L624:
	;
	if v2372 != 0 {
		goto L625
	} else {
		goto L626
	}
L625:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2374
	v2376 = F_slice_del(m, l0)
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L31
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2407 = v2371 - v2369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2406 - v2407
	v2410 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L31
	} else {
		goto L641
	}
L628:
	;
	if v2376 < int32(0) {
		v3782 = v2376
		goto L1
	} else {
		goto L629
	}
L629:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2380
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2383 = v2382 - v2380
	v2384 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L31
	} else {
		goto L630
	}
L630:
	;
	if v2384 == int32(0) {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2388 - v2383
	goto L408
L632:
	;
	goto L633
L633:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2391
	v2393 = F_slice_del(m, l0)
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L31
	} else {
		goto L634
	}
L634:
	;
	if v2393 < int32(0) {
		v3782 = v2393
		goto L1
	} else {
		goto L635
	}
L635:
	;
	v2397 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L31
	} else {
		goto L636
	}
L636:
	;
	if v2397 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2401 - v2383
	goto L408
L638:
	;
	goto L639
L639:
	;
	if int32(0) <= v2397 {
		goto L408
	} else {
		goto L640
	}
L640:
	;
	v3782 = v2397
	goto L1
L641:
	;
	if v2410 != 0 {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2412
	v2414 = F_slice_del(m, l0)
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L31
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2430 - v2407
	v2433 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L31
	} else {
		goto L652
	}
L645:
	;
	if v2414 < int32(0) {
		v3782 = v2414
		goto L1
	} else {
		goto L646
	}
L646:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2420 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L31
	} else {
		goto L647
	}
L647:
	;
	if v2420 == int32(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2424 + (v2418 - v2419)
	goto L408
L649:
	;
	goto L650
L650:
	;
	if int32(0) <= v2420 {
		goto L408
	} else {
		goto L651
	}
L651:
	;
	v3782 = v2420
	goto L1
L652:
	;
	if v2433 == int32(0) {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2437 + (v2369 - v2371)
	goto L408
L654:
	;
	goto L655
L655:
	;
	if int32(0) <= v2433 {
		goto L408
	} else {
		goto L656
	}
L656:
	;
	v3782 = v2433
	goto L1
L657:
	;
	if v2542 == int32(0) {
		goto L681
	} else {
		goto L682
	}
L658:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2453 = v2451 - int32(1)
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2453 <= v2454 {
		v2542 = v2447
		goto L657
	} else {
		goto L659
	}
L659:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2456+v2453))))
	if v2458 != int32(110) {
		v2542 = v2447
		goto L657
	} else {
		goto L660
	}
L660:
	;
	v2463 = F_find_among_b(m, l0, int32(4358256), int32(4))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L31
	} else {
		goto L661
	}
L661:
	;
	if v2463 == int32(0) {
		v2542 = v2447
		goto L657
	} else {
		goto L662
	}
L662:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2475 <= v2476 {
		goto L668
	} else {
		goto L669
	}
L663:
	;
	v2542 = v2537
	goto L657
L664:
	;
	goto L663
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2533
	v2537 = int32(1)
	goto L664
L666:
	;
	v2533 = v2485 - v2473 + v2492
	goto L665
L667:
	;
	v2500 = v2475 - v2473
	v2501 = v2499 + v2500
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2501
	if v2498 < v2501 {
		goto L673
	} else {
		goto L674
	}
L668:
	;
	v2497 = v2474
	v2498 = v2476
	v2499 = v2473
	goto L667
L669:
	;
	goto L670
L670:
	;
	v2481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2475+v2474-int32(1)))))
	if v2481 != int32(110) {
		v2497 = v2474
		v2498 = v2476
		v2499 = v2473
		goto L667
	} else {
		goto L671
	}
L671:
	;
	v2485 = v2475 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2485
	v2490 = int32(0)
	v2491 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v2490)
	mBase = m.M
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2491 == v2490 {
		goto L666
	} else {
		goto L672
	}
L672:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2497 = v2495
	v2498 = v2496
	v2499 = v2492
	goto L667
L673:
	;
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2501+v2497-int32(1)))))
	if v2508 == int32(110) {
		v2537 = int32(0)
		goto L664
	} else {
		goto L676
	}
L674:
	;
	goto L675
L675:
	;
	v2512 = int32(0)
	v2514 = F_skip_b_utf8(m, v2497, v2501, v2498, int32(1))
	mBase = m.M
	if v2514 < v2512 {
		v2537 = v2512
		goto L664
	} else {
		goto L677
	}
L676:
	;
	goto L675
L677:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2514
	v2522 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2522 != 0 {
		v2537 = v2512
		goto L664
	} else {
		goto L678
	}
L678:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2533 = v2523 + v2500
	goto L665
L679:
	;
	if v3392 < int32(0) {
		v3782 = v3392
		goto L1
	} else {
		goto L921
	}
L680:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2831
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2831
	v2834 = F_r_mark_lArI(m, l0)
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L31
	} else {
		goto L766
	}
L681:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2546
	v2548 = int32(0)
	v2549 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2549 == v2548 {
		v2643 = v2548
		goto L684
	} else {
		goto L685
	}
L682:
	;
	goto L683
L683:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2649
	v2651 = F_slice_del(m, l0)
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L31
	} else {
		goto L707
	}
L684:
	;
	if v2643 == int32(0) {
		goto L680
	} else {
		goto L706
	}
L685:
	;
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2554 = v2552 - int32(1)
	v2555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2554 <= v2555 {
		v2643 = v2548
		goto L684
	} else {
		goto L686
	}
L686:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2557+v2554))))
	switch v2559 - int32(97) {
	case 0, 4:
		goto L687
	default:
		v2643 = v2548
		goto L684
	}
L687:
	;
	v2564 = F_find_among_b(m, l0, int32(4358336), int32(2))
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L31
	} else {
		goto L688
	}
L688:
	;
	if v2564 == int32(0) {
		v2643 = v2548
		goto L684
	} else {
		goto L689
	}
L689:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2576 <= v2577 {
		goto L695
	} else {
		goto L696
	}
L690:
	;
	v2643 = v2638
	goto L684
L691:
	;
	goto L690
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2634
	v2638 = int32(1)
	goto L691
L693:
	;
	v2634 = v2586 - v2574 + v2593
	goto L692
L694:
	;
	v2601 = v2576 - v2574
	v2602 = v2600 + v2601
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2602
	if v2599 < v2602 {
		goto L700
	} else {
		goto L701
	}
L695:
	;
	v2598 = v2575
	v2599 = v2577
	v2600 = v2574
	goto L694
L696:
	;
	goto L697
L697:
	;
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2576+v2575-int32(1)))))
	if v2582 != int32(121) {
		v2598 = v2575
		v2599 = v2577
		v2600 = v2574
		goto L694
	} else {
		goto L698
	}
L698:
	;
	v2586 = v2576 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2586
	v2591 = int32(0)
	v2592 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v2591)
	mBase = m.M
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2592 == v2591 {
		goto L693
	} else {
		goto L699
	}
L699:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2598 = v2596
	v2599 = v2597
	v2600 = v2593
	goto L694
L700:
	;
	v2609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2602+v2598-int32(1)))))
	if v2609 == int32(121) {
		v2638 = int32(0)
		goto L691
	} else {
		goto L703
	}
L701:
	;
	goto L702
L702:
	;
	v2613 = int32(0)
	v2615 = F_skip_b_utf8(m, v2598, v2602, v2599, int32(1))
	mBase = m.M
	if v2615 < v2613 {
		v2638 = v2613
		goto L691
	} else {
		goto L704
	}
L703:
	;
	goto L702
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2615
	v2623 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2623 != 0 {
		v2638 = v2613
		goto L691
	} else {
		goto L705
	}
L705:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2634 = v2624 + v2601
	goto L692
L706:
	;
	goto L683
L707:
	;
	if v2651 < int32(0) {
		v3782 = v2651
		goto L1
	} else {
		goto L708
	}
L708:
	;
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2655
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2658 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2659 = m.ExcPending
	if v2659 != 0 {
		goto L31
	} else {
		goto L710
	}
L709:
	;
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2685 = v2657 - v2655
	v2686 = v2684 - v2685
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2686
	v2689 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L31
	} else {
		goto L722
	}
L710:
	;
	if v2658 == int32(0) {
		goto L709
	} else {
		goto L711
	}
L711:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2662
	v2664 = F_slice_del(m, l0)
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L31
	} else {
		goto L712
	}
L712:
	;
	if v2664 < int32(0) {
		v3782 = v2664
		goto L1
	} else {
		goto L713
	}
L713:
	;
	v2668 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L31
	} else {
		goto L714
	}
L714:
	;
	v2671 = int32(base.Ui32(v2668) >> (uint(int32(31)) % 32))
	if v2668 != 0 {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v2673 = v2671
	goto L717
L716:
	;
	v2673 = int32(36)
	goto L717
L717:
	;
	if v2673 == int32(0) {
		goto L408
	} else {
		goto L718
	}
L718:
	;
	if v2673 == int32(36) {
		goto L709
	} else {
		goto L719
	}
L719:
	;
	if v2671 != 0 {
		v3392 = v2668 >> (uint(int32(31)) % 32) & v2668
		goto L679
	} else {
		goto L720
	}
L720:
	;
	goto L408
L721:
	;
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2817 - v2685
	v2820 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L31
	} else {
		goto L760
	}
L722:
	;
	if v2689 == int32(0) {
		goto L723
	} else {
		goto L724
	}
L723:
	;
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2693 - v2685
	v2696 = int32(0)
	v2702 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2702 == v2696 {
		goto L727
	} else {
		goto L728
	}
L724:
	;
	goto L725
L725:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2785
	v2787 = F_slice_del(m, l0)
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L31
	} else {
		goto L747
	}
L726:
	;
	if v2782 == int32(0) {
		goto L721
	} else {
		goto L746
	}
L727:
	;
	v2782 = int32(0)
	goto L726
L728:
	;
	goto L729
L729:
	;
	v2710 = F_in_grouping_b_U(m, l0, int32(2204288), int32(105), int32(305), int32(0))
	mBase = m.M
	if v2710 != 0 {
		v2774 = v2696
		goto L730
	} else {
		goto L731
	}
L730:
	;
	v2782 = v2774
	goto L726
L731:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2713 <= v2714 {
		goto L735
	} else {
		goto L736
	}
L732:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2767
	v2774 = int32(1)
	goto L730
L733:
	;
	v2767 = v2723 - v2711 + v2730
	goto L732
L734:
	;
	v2738 = v2713 - v2711
	v2739 = v2737 + v2738
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2739
	if v2736 < v2739 {
		goto L740
	} else {
		goto L741
	}
L735:
	;
	v2735 = v2712
	v2736 = v2714
	v2737 = v2711
	goto L734
L736:
	;
	goto L737
L737:
	;
	v2719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2712+v2713-int32(1)))))
	if v2719 != int32(115) {
		v2735 = v2712
		v2736 = v2714
		v2737 = v2711
		goto L734
	} else {
		goto L738
	}
L738:
	;
	v2723 = v2713 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2723
	v2728 = int32(0)
	v2729 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v2728)
	mBase = m.M
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2729 == v2728 {
		goto L733
	} else {
		goto L739
	}
L739:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2735 = v2733
	v2736 = v2734
	v2737 = v2730
	goto L734
L740:
	;
	v2745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2739+v2735-int32(1)))))
	if v2745 == int32(115) {
		v2774 = v2696
		goto L730
	} else {
		goto L743
	}
L741:
	;
	goto L742
L742:
	;
	v2749 = F_skip_b_utf8(m, v2735, v2739, v2736, int32(1))
	mBase = m.M
	if v2749 < int32(0) {
		v2774 = v2696
		goto L730
	} else {
		goto L744
	}
L743:
	;
	goto L742
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2749
	v2757 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2757 != 0 {
		v2774 = v2696
		goto L730
	} else {
		goto L745
	}
L745:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2767 = v2758 + v2738
	goto L732
L746:
	;
	goto L725
L747:
	;
	if v2787 < int32(0) {
		v3782 = v2787
		goto L1
	} else {
		goto L748
	}
L748:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2791
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2794 = v2793 - v2791
	v2795 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L31
	} else {
		goto L749
	}
L749:
	;
	if v2795 == int32(0) {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2799 - v2794
	goto L408
L751:
	;
	goto L752
L752:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2802
	v2804 = F_slice_del(m, l0)
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L31
	} else {
		goto L753
	}
L753:
	;
	if v2804 < int32(0) {
		v3782 = v2804
		goto L1
	} else {
		goto L754
	}
L754:
	;
	v2808 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L31
	} else {
		goto L755
	}
L755:
	;
	if v2808 == int32(0) {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2812 - v2794
	goto L408
L757:
	;
	goto L758
L758:
	;
	if int32(0) <= v2808 {
		goto L408
	} else {
		goto L759
	}
L759:
	;
	v3782 = v2808
	goto L1
L760:
	;
	if v2820 == int32(0) {
		goto L761
	} else {
		goto L762
	}
L761:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2824 - v2685
	goto L408
L762:
	;
	goto L763
L763:
	;
	if int32(0) <= v2820 {
		goto L408
	} else {
		goto L764
	}
L764:
	;
	if int32(base.Ui32(v2820)>>(uint(int32(31))%32)) != 0 {
		v3392 = v2820
		goto L679
	} else {
		goto L765
	}
L765:
	;
	goto L408
L766:
	;
	if v2834 != 0 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2836
	v2838 = F_slice_del(m, l0)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L31
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2842
	v2844 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L31
	} else {
		goto L772
	}
L770:
	;
	if int32(0) <= v2838 {
		goto L408
	} else {
		goto L771
	}
L771:
	;
	v3782 = v2838
	goto L1
L772:
	;
	v2847 = int32(base.Ui32(v2844) >> (uint(int32(31)) % 32))
	if v2844 != 0 {
		goto L773
	} else {
		goto L774
	}
L773:
	;
	v2849 = v2847
	goto L775
L774:
	;
	v2849 = int32(44)
	goto L775
L775:
	;
	if v2849 == int32(0) {
		goto L408
	} else {
		goto L776
	}
L776:
	;
	v2854 = v2844 >> (uint(int32(31)) % 32) & v2844
	if v2849 != int32(44) {
		goto L778
	} else {
		goto L779
	}
L777:
	;
	if v3388 == int32(0) {
		goto L408
	} else {
		goto L920
	}
L778:
	;
	v3386 = v2854
	v3388 = v2847
	goto L777
L779:
	;
	goto L780
L780:
	;
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2857
	v2860 = int32(0)
	v2861 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2861 == v2860 {
		v2881 = v2860
		goto L781
	} else {
		goto L782
	}
L781:
	;
	if v2881 != 0 {
		goto L787
	} else {
		goto L788
	}
L782:
	;
	v2864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2866 = v2864 - int32(1)
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2866 <= v2867 {
		v2881 = v2860
		goto L781
	} else {
		goto L783
	}
L783:
	;
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2869+v2866))))
	switch v2871 - int32(97) {
	case 0, 4:
		goto L784
	default:
		v2881 = v2860
		goto L781
	}
L784:
	;
	v2876 = F_find_among_b(m, l0, int32(4358384), int32(4))
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L31
	} else {
		goto L785
	}
L785:
	;
	v2881 = base.B2i32(v2876 != int32(0))
	goto L781
L786:
	;
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3253
	v3256 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v3257 = m.ExcPending
	if v3257 != 0 {
		goto L31
	} else {
		goto L880
	}
L787:
	;
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3196
	v3198 = F_slice_del(m, l0)
	mBase = m.M
	v3199 = m.ExcPending
	if v3199 != 0 {
		goto L31
	} else {
		goto L859
	}
L788:
	;
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2882
	v2884 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2884 != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L794
L790:
	;
	v3092 = int32(0)
	goto L791
L791:
	;
	if v3092 != 0 {
		goto L787
	} else {
		goto L835
	}
L792:
	;
	if v3013 != 0 {
		goto L816
	} else {
		goto L817
	}
L793:
	;
	v3013 = v3006
	goto L792
L794:
	;
	if v2901 <= v2902 {
		v3006 = int32(-1)
		goto L793
	} else {
		goto L796
	}
L795:
	;
	v3006 = int32(0)
	goto L793
L796:
	;
	v2919 = int32(1)
	v2920 = v2901 - v2919
	v2922 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2898+v2920))))
	v2924 = v2922 & int32(255)
	if v2920 == v2902 {
		v2979 = v2924
		v2980 = v2919
		goto L797
	} else {
		goto L798
	}
L797:
	;
	if int32(305) < v2979 {
		goto L806
	} else {
		goto L807
	}
L798:
	;
	if int32(0) <= v2922 {
		v2979 = v2924
		v2980 = v2919
		goto L797
	} else {
		goto L799
	}
L799:
	;
	v2930 = v2924 & int32(63)
	v2932 = v2901 - int32(2)
	v2934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2898+v2932))))
	v2936 = v2934 << (uint(int32(6)) % 32)
	if base.B2i32(v2932 != v2902)&base.B2i32(base.Ui32(v2934) < base.Ui32(int32(192))) == int32(0) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	v2979 = v2936&int32(1984) | v2930
	v2980 = int32(2)
	goto L797
L801:
	;
	goto L802
L802:
	;
	v2949 = v2936&int32(4032) | v2930
	v2951 = v2901 - int32(3)
	v2953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2898+v2951))))
	if base.B2i32(v2951 != v2902)&base.B2i32(base.Ui32(v2953) < base.Ui32(int32(224))) == int32(0) {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v2979 = v2953<<(uint(int32(12))%32)&int32(61440) | v2949
	v2980 = int32(3)
	goto L797
L804:
	;
	goto L805
L805:
	;
	v2971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2901+(v2898-int32(4))))))
	v2979 = v2953<<(uint(int32(12))%32)&int32(258048) | v2971&int32(7)<<(uint(int32(18))%32) | v2949
	v2980 = int32(4)
	goto L797
L806:
	;
	v3013 = v2980
	goto L792
L807:
	;
	goto L808
L808:
	;
	v2984 = v2979 - int32(105)
	if v2984 < int32(0) {
		goto L809
	} else {
		goto L810
	}
L809:
	;
	v3013 = v2980
	goto L792
L810:
	;
	goto L811
L811:
	;
	v2990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2984)>>(uint(int32(3))%32)))+uint32(_consts[1449]))))
	if int32(base.Ui32(v2990)>>(uint(v2984&int32(7))%32))&int32(1) == int32(0) {
		goto L812
	} else {
		goto L813
	}
L812:
	;
	v3013 = v2980
	goto L792
L813:
	;
	goto L814
L814:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2901 - v2980
	goto L815
L815:
	;
	goto L795
L816:
	;
	v3090 = int32(0)
	goto L818
L817:
	;
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3023 <= v3024 {
		goto L824
	} else {
		goto L825
	}
L818:
	;
	v3092 = v3090
	goto L791
L819:
	;
	v3090 = v3085
	goto L818
L820:
	;
	goto L819
L821:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3081
	v3085 = int32(1)
	goto L820
L822:
	;
	v3081 = v3033 - v3021 + v3040
	goto L821
L823:
	;
	v3048 = v3023 - v3021
	v3049 = v3047 + v3048
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3049
	if v3046 < v3049 {
		goto L829
	} else {
		goto L830
	}
L824:
	;
	v3045 = v3022
	v3046 = v3024
	v3047 = v3021
	goto L823
L825:
	;
	goto L826
L826:
	;
	v3029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023+v3022-int32(1)))))
	if v3029 != int32(121) {
		v3045 = v3022
		v3046 = v3024
		v3047 = v3021
		goto L823
	} else {
		goto L827
	}
L827:
	;
	v3033 = v3023 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3033
	v3038 = int32(0)
	v3039 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v3038)
	mBase = m.M
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3039 == v3038 {
		goto L822
	} else {
		goto L828
	}
L828:
	;
	v3043 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3045 = v3043
	v3046 = v3044
	v3047 = v3040
	goto L823
L829:
	;
	v3056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3049+v3045-int32(1)))))
	if v3056 == int32(121) {
		v3085 = int32(0)
		goto L820
	} else {
		goto L832
	}
L830:
	;
	goto L831
L831:
	;
	v3060 = int32(0)
	v3062 = F_skip_b_utf8(m, v3045, v3049, v3046, int32(1))
	mBase = m.M
	if v3062 < v3060 {
		v3085 = v3060
		goto L820
	} else {
		goto L833
	}
L832:
	;
	goto L831
L833:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3062
	v3070 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v3070 != 0 {
		v3085 = v3060
		goto L820
	} else {
		goto L834
	}
L834:
	;
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3081 = v3071 + v3048
	goto L821
L835:
	;
	v3093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3093
	v3095 = int32(0)
	v3096 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v3096 == v3095 {
		v3191 = v3095
		goto L836
	} else {
		goto L837
	}
L836:
	;
	if v3191 == int32(0) {
		goto L786
	} else {
		goto L858
	}
L837:
	;
	v3099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3099 <= v3100 {
		v3191 = v3095
		goto L836
	} else {
		goto L838
	}
L838:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3102+v3099-int32(1)))))
	switch v3106 - int32(97) {
	case 0, 4:
		goto L839
	default:
		v3191 = v3095
		goto L836
	}
L839:
	;
	v3111 = F_find_among_b(m, l0, int32(4358464), int32(2))
	mBase = m.M
	v3112 = m.ExcPending
	if v3112 != 0 {
		goto L31
	} else {
		goto L840
	}
L840:
	;
	if v3111 == int32(0) {
		v3191 = v3095
		goto L836
	} else {
		goto L841
	}
L841:
	;
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3123 <= v3124 {
		goto L847
	} else {
		goto L848
	}
L842:
	;
	v3191 = v3185
	goto L836
L843:
	;
	goto L842
L844:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3181
	v3185 = int32(1)
	goto L843
L845:
	;
	v3181 = v3133 - v3121 + v3140
	goto L844
L846:
	;
	v3148 = v3123 - v3121
	v3149 = v3147 + v3148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3149
	if v3146 < v3149 {
		goto L852
	} else {
		goto L853
	}
L847:
	;
	v3145 = v3122
	v3146 = v3124
	v3147 = v3121
	goto L846
L848:
	;
	goto L849
L849:
	;
	v3129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3123+v3122-int32(1)))))
	if v3129 != int32(121) {
		v3145 = v3122
		v3146 = v3124
		v3147 = v3121
		goto L846
	} else {
		goto L850
	}
L850:
	;
	v3133 = v3123 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3133
	v3138 = int32(0)
	v3139 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v3138)
	mBase = m.M
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3139 == v3138 {
		goto L845
	} else {
		goto L851
	}
L851:
	;
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3145 = v3143
	v3146 = v3144
	v3147 = v3140
	goto L846
L852:
	;
	v3156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3149+v3145-int32(1)))))
	if v3156 == int32(121) {
		v3185 = int32(0)
		goto L843
	} else {
		goto L855
	}
L853:
	;
	goto L854
L854:
	;
	v3160 = int32(0)
	v3162 = F_skip_b_utf8(m, v3145, v3149, v3146, int32(1))
	mBase = m.M
	if v3162 < v3160 {
		v3185 = v3160
		goto L843
	} else {
		goto L856
	}
L855:
	;
	goto L854
L856:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3162
	v3170 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v3170 != 0 {
		v3185 = v3160
		goto L843
	} else {
		goto L857
	}
L857:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3181 = v3171 + v3148
	goto L844
L858:
	;
	goto L787
L859:
	;
	if v3198 < int32(0) {
		v3782 = v3198
		goto L1
	} else {
		goto L860
	}
L860:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3202
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3205 = v3204 - v3202
	v3206 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v3207 = m.ExcPending
	if v3207 != 0 {
		goto L31
	} else {
		goto L863
	}
L861:
	;
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3250 - v3205
	goto L408
L862:
	;
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3233
	v3235 = F_slice_del(m, l0)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L31
	} else {
		goto L873
	}
L863:
	;
	if v3206 != 0 {
		goto L864
	} else {
		goto L865
	}
L864:
	;
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3208
	v3210 = F_slice_del(m, l0)
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L31
	} else {
		goto L867
	}
L865:
	;
	goto L866
L866:
	;
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3223 + (v3202 - v3204)
	v3227 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v3228 = m.ExcPending
	if v3228 != 0 {
		goto L31
	} else {
		goto L871
	}
L867:
	;
	if v3210 < int32(0) {
		v3782 = v3210
		goto L1
	} else {
		goto L868
	}
L868:
	;
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3214
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3217 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L31
	} else {
		goto L869
	}
L869:
	;
	if v3217 != 0 {
		goto L862
	} else {
		goto L870
	}
L870:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3219 + (v3214 - v3216)
	goto L862
L871:
	;
	if v3227 == int32(0) {
		goto L861
	} else {
		goto L872
	}
L872:
	;
	goto L862
L873:
	;
	if v3235 < int32(0) {
		v3782 = v3235
		goto L1
	} else {
		goto L874
	}
L874:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3239
	v3241 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L31
	} else {
		goto L875
	}
L875:
	;
	if v3241 == int32(0) {
		goto L876
	} else {
		goto L877
	}
L876:
	;
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3245 - v3205
	goto L408
L877:
	;
	goto L878
L878:
	;
	if int32(0) <= v3241 {
		goto L408
	} else {
		goto L879
	}
L879:
	;
	v3782 = v3241
	goto L1
L880:
	;
	if v3256 == int32(0) {
		goto L881
	} else {
		goto L882
	}
L881:
	;
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3260
	v3262 = int32(0)
	v3268 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v3268 == v3262 {
		goto L885
	} else {
		goto L886
	}
L882:
	;
	goto L883
L883:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3351
	v3353 = F_slice_del(m, l0)
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L31
	} else {
		goto L905
	}
L884:
	;
	if v3348 == int32(0) {
		goto L408
	} else {
		goto L904
	}
L885:
	;
	v3348 = int32(0)
	goto L884
L886:
	;
	goto L887
L887:
	;
	v3276 = F_in_grouping_b_U(m, l0, int32(2204288), int32(105), int32(305), int32(0))
	mBase = m.M
	if v3276 != 0 {
		v3340 = v3262
		goto L888
	} else {
		goto L889
	}
L888:
	;
	v3348 = v3340
	goto L884
L889:
	;
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3279 <= v3280 {
		goto L893
	} else {
		goto L894
	}
L890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3333
	v3340 = int32(1)
	goto L888
L891:
	;
	v3333 = v3289 - v3277 + v3296
	goto L890
L892:
	;
	v3304 = v3279 - v3277
	v3305 = v3303 + v3304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3305
	if v3302 < v3305 {
		goto L898
	} else {
		goto L899
	}
L893:
	;
	v3301 = v3278
	v3302 = v3280
	v3303 = v3277
	goto L892
L894:
	;
	goto L895
L895:
	;
	v3285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3278+v3279-int32(1)))))
	if v3285 != int32(115) {
		v3301 = v3278
		v3302 = v3280
		v3303 = v3277
		goto L892
	} else {
		goto L896
	}
L896:
	;
	v3289 = v3279 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3289
	v3294 = int32(0)
	v3295 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), v3294)
	mBase = m.M
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3295 == v3294 {
		goto L891
	} else {
		goto L897
	}
L897:
	;
	v3299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3301 = v3299
	v3302 = v3300
	v3303 = v3296
	goto L892
L898:
	;
	v3311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3305+v3301-int32(1)))))
	if v3311 == int32(115) {
		v3340 = v3262
		goto L888
	} else {
		goto L901
	}
L899:
	;
	goto L900
L900:
	;
	v3315 = F_skip_b_utf8(m, v3301, v3305, v3302, int32(1))
	mBase = m.M
	if v3315 < int32(0) {
		v3340 = v3262
		goto L888
	} else {
		goto L902
	}
L901:
	;
	goto L900
L902:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3315
	v3323 = F_in_grouping_b_U(m, l0, int32(2203840), int32(97), int32(305), int32(0))
	mBase = m.M
	if v3323 != 0 {
		v3340 = v3262
		goto L888
	} else {
		goto L903
	}
L903:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3333 = v3324 + v3304
	goto L890
L904:
	;
	goto L883
L905:
	;
	if v3353 < int32(0) {
		v3782 = v3353
		goto L1
	} else {
		goto L906
	}
L906:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3357
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3360 = v3359 - v3357
	v3361 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L31
	} else {
		goto L907
	}
L907:
	;
	if v3361 == int32(0) {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v3365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3365 - v3360
	goto L408
L909:
	;
	goto L910
L910:
	;
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3368
	v3370 = F_slice_del(m, l0)
	mBase = m.M
	v3371 = m.ExcPending
	if v3371 != 0 {
		goto L31
	} else {
		goto L911
	}
L911:
	;
	if v3370 < int32(0) {
		v3782 = v3370
		goto L1
	} else {
		goto L912
	}
L912:
	;
	v3374 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L31
	} else {
		goto L913
	}
L913:
	;
	if v3374 == int32(0) {
		goto L914
	} else {
		goto L915
	}
L914:
	;
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3378 - v3360
	goto L408
L915:
	;
	goto L916
L916:
	;
	if v3374 < int32(0) {
		goto L917
	} else {
		goto L918
	}
L917:
	;
	v3383 = v3374
	goto L919
L918:
	;
	v3383 = v2854
	goto L919
L919:
	;
	v3386 = v3383
	v3388 = int32(base.Ui32(v3374) >> (uint(int32(31)) % 32))
	goto L777
L920:
	;
	v3392 = v3386
	goto L679
L921:
	;
	goto L408
L922:
	;
	if int32(0) <= v3401 {
		goto L408
	} else {
		goto L923
	}
L923:
	;
	v3782 = v3401
	goto L1
L924:
	;
	if v3407 < int32(0) {
		v3782 = v3407
		goto L1
	} else {
		goto L925
	}
L925:
	;
	goto L408
L926:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3433 == int32(0) {
		goto L931
	} else {
		goto L932
	}
L927:
	;
	goto L926
L928:
	;
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3429 = F_memcmp(m, v3426+v3416-v3418, int32(2204378), v3418)
	mBase = m.M
	if v3429 != 0 {
		v3433 = v3420
		goto L927
	} else {
		goto L929
	}
L929:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3416 - v3418
	v3433 = int32(1)
	goto L927
L930:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3465
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3465
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3465
	if v3465 <= v3466 {
		goto L943
	} else {
		goto L944
	}
L931:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3465 = v3434
	v3466 = v3437
	goto L930
L932:
	;
	goto L933
L933:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3439 = int32(3)
	v3441 = int32(0)
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3438-v3444 < v3439 {
		v3454 = v3441
		goto L936
	} else {
		goto L937
	}
L934:
	;
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3460 <= v3462 {
		v3782 = int32(0)
		goto L1
	} else {
		goto L942
	}
L935:
	;
	if v3454 != 0 {
		goto L939
	} else {
		goto L940
	}
L936:
	;
	goto L935
L937:
	;
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3450 = F_memcmp(m, v3447+v3438-v3439, int32(2204380), v3439)
	mBase = m.M
	if v3450 != 0 {
		v3454 = v3441
		goto L936
	} else {
		goto L938
	}
L938:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3438 - v3439
	v3454 = int32(1)
	goto L936
L939:
	;
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3460 = v3455
	goto L934
L940:
	;
	goto L941
L941:
	;
	v3456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3458 = v3456 + (v3438 - v3434)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3458
	v3460 = v3458
	goto L934
L942:
	;
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3465 = v3464
	v3466 = v3462
	goto L930
L943:
	;
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3741
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3741
	v3746 = F_find_among_b(m, l0, int32(4358512), int32(4))
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L31
	} else {
		goto L1001
	}
L944:
	;
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3472+v3465-int32(1)))))
	switch v3476 - int32(100) {
	case 0, 3:
		goto L945
	default:
		goto L943
	}
L945:
	;
	v3480 = v3465 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3480
	v3495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3508 = v3480
	goto L948
L946:
	;
	if v3611 < int32(0) {
		goto L943
	} else {
		goto L965
	}
L947:
	;
	v3611 = int32(-1)
	goto L946
L948:
	;
	if v3508 <= v3499 {
		goto L947
	} else {
		goto L950
	}
L950:
	;
	v3516 = int32(1)
	v3517 = v3508 - v3516
	v3519 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3495+v3517))))
	v3521 = v3519 & int32(255)
	if v3517 == v3499 {
		v3576 = v3521
		v3577 = v3516
		goto L951
	} else {
		goto L952
	}
L951:
	;
	if int32(305) < v3576 {
		goto L960
	} else {
		goto L961
	}
L952:
	;
	if int32(0) <= v3519 {
		v3576 = v3521
		v3577 = v3516
		goto L951
	} else {
		goto L953
	}
L953:
	;
	v3527 = v3521 & int32(63)
	v3529 = v3508 - int32(2)
	v3531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3495+v3529))))
	v3533 = v3531 << (uint(int32(6)) % 32)
	if base.B2i32(v3529 != v3499)&base.B2i32(base.Ui32(v3531) < base.Ui32(int32(192))) == int32(0) {
		goto L954
	} else {
		goto L955
	}
L954:
	;
	v3576 = v3533&int32(1984) | v3527
	v3577 = int32(2)
	goto L951
L955:
	;
	goto L956
L956:
	;
	v3546 = v3533&int32(4032) | v3527
	v3548 = v3508 - int32(3)
	v3550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3495+v3548))))
	if base.B2i32(v3548 != v3499)&base.B2i32(base.Ui32(v3550) < base.Ui32(int32(224))) == int32(0) {
		goto L957
	} else {
		goto L958
	}
L957:
	;
	v3576 = v3550<<(uint(int32(12))%32)&int32(61440) | v3546
	v3577 = int32(3)
	goto L951
L958:
	;
	goto L959
L959:
	;
	v3568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3508+(v3495-int32(4))))))
	v3576 = v3550<<(uint(int32(12))%32)&int32(258048) | v3568&int32(7)<<(uint(int32(18))%32) | v3546
	v3577 = int32(4)
	goto L951
L960:
	;
	v3596 = v3508 - v3577
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3596
	v3508 = v3596
	goto L948
L961:
	;
	v3581 = v3576 - int32(97)
	if v3581 < int32(0) {
		goto L960
	} else {
		goto L962
	}
L962:
	;
	v3587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3581)>>(uint(int32(3))%32)))+uint32(_consts[1448]))))
	if int32(base.Ui32(v3587)>>(uint(v3581&int32(7))%32))&int32(1) == int32(0) {
		goto L960
	} else {
		goto L963
	}
L963:
	;
	v3611 = v3577
	goto L946
L965:
	;
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3615 <= v3616 {
		goto L968
	} else {
		goto L969
	}
L966:
	;
	v3652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3653 = v3614 - v3615
	v3654 = v3652 - v3653
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3654
	v3656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3654 <= v3656 {
		goto L979
	} else {
		goto L980
	}
L967:
	;
	v3648 = F_slice_from_s(m, l0, int32(2), int32(2204385))
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L31
	} else {
		goto L976
	}
L968:
	;
	v3628 = int32(2)
	v3630 = int32(0)
	v3632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3632-v3633 < v3628 {
		v3643 = v3630
		goto L972
	} else {
		goto L973
	}
L969:
	;
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3618+v3615-int32(1)))))
	if v3622 != int32(97) {
		goto L968
	} else {
		goto L970
	}
L970:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3615 - int32(1)
	goto L967
L971:
	;
	if v3643 == int32(0) {
		goto L966
	} else {
		goto L975
	}
L972:
	;
	goto L971
L973:
	;
	v3636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3639 = F_memcmp(m, v3636+v3632-v3628, int32(2204383), v3628)
	mBase = m.M
	if v3639 != 0 {
		v3643 = v3630
		goto L972
	} else {
		goto L974
	}
L974:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3632 - v3628
	v3643 = int32(1)
	goto L972
L975:
	;
	goto L967
L976:
	;
	if int32(0) <= v3648 {
		goto L943
	} else {
		goto L977
	}
L977:
	;
	v3782 = v3648
	goto L1
L978:
	;
	v3729 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3654 - v3729
	v3734 = F_slice_from_s(m, l0, v3729, int32(2204387))
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
		goto L31
	} else {
		goto L998
	}
L979:
	;
	v3684 = int32(2)
	v3686 = int32(0)
	v3688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3688-v3689 < v3684 {
		v3699 = v3686
		goto L985
	} else {
		goto L986
	}
L980:
	;
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3660 = int32(1)
	v3662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3658+v3654-v3660))))
	v3664 = v3662 - int32(101)
	switch (v3664<<(uint(int32(7))%32) | int32(base.Ui32(v3664&int32(254))>>(uint(v3660)%32))) & int32(255) {
	case 0, 2:
		goto L978
	default:
		goto L979
	case 5, 8:
		goto L981
	}
L981:
	;
	v3674 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3654 - v3674
	v3679 = F_slice_from_s(m, l0, v3674, int32(2204388))
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L31
	} else {
		goto L982
	}
L982:
	;
	if int32(0) <= v3679 {
		goto L943
	} else {
		goto L983
	}
L983:
	;
	v3782 = v3679
	goto L1
L984:
	;
	if v3699 == int32(0) {
		goto L988
	} else {
		goto L989
	}
L985:
	;
	goto L984
L986:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3695 = F_memcmp(m, v3692+v3688-v3684, int32(2204389), v3684)
	mBase = m.M
	if v3695 != 0 {
		v3699 = v3686
		goto L985
	} else {
		goto L987
	}
L987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3688 - v3684
	v3699 = int32(1)
	goto L985
L988:
	;
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3703 = v3702 - v3653
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3703
	v3705 = int32(2)
	v3707 = int32(0)
	v3710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3703-v3710 < v3705 {
		v3720 = v3707
		goto L992
	} else {
		goto L993
	}
L989:
	;
	goto L990
L990:
	;
	v3725 = F_slice_from_s(m, l0, int32(2), int32(2204393))
	mBase = m.M
	v3726 = m.ExcPending
	if v3726 != 0 {
		goto L31
	} else {
		goto L996
	}
L991:
	;
	if v3720 == int32(0) {
		goto L943
	} else {
		goto L995
	}
L992:
	;
	goto L991
L993:
	;
	v3713 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3716 = F_memcmp(m, v3713+v3703-v3705, int32(2204391), v3705)
	mBase = m.M
	if v3716 != 0 {
		v3720 = v3707
		goto L992
	} else {
		goto L994
	}
L994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3703 - v3705
	v3720 = int32(1)
	goto L992
L995:
	;
	goto L990
L996:
	;
	if int32(0) <= v3725 {
		goto L943
	} else {
		goto L997
	}
L997:
	;
	v3782 = v3725
	goto L1
L998:
	;
	if v3734 < int32(0) {
		v3782 = v3734
		goto L1
	} else {
		goto L999
	}
L999:
	;
	goto L943
L1000:
	;
	v3779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3779
	v3782 = int32(1)
	goto L1
L1001:
	;
	if v3746 == int32(0) {
		goto L1000
	} else {
		goto L1002
	}
L1002:
	;
	v3750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3750
	switch v3746 - int32(1) {
	case 0:
		goto L1006
	case 1:
		goto L1005
	case 2:
		goto L1004
	case 3:
		goto L1003
	default:
		goto L1000
	}
L1003:
	;
	v3774 = F_slice_from_s(m, l0, int32(1), int32(2204399))
	mBase = m.M
	v3775 = m.ExcPending
	if v3775 != 0 {
		goto L31
	} else {
		goto L1013
	}
L1004:
	;
	v3768 = F_slice_from_s(m, l0, int32(1), int32(2204398))
	mBase = m.M
	v3769 = m.ExcPending
	if v3769 != 0 {
		goto L31
	} else {
		goto L1011
	}
L1005:
	;
	v3762 = F_slice_from_s(m, l0, int32(2), int32(2204396))
	mBase = m.M
	v3763 = m.ExcPending
	if v3763 != 0 {
		goto L31
	} else {
		goto L1009
	}
L1006:
	;
	v3756 = F_slice_from_s(m, l0, int32(1), int32(2204395))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L31
	} else {
		goto L1007
	}
L1007:
	;
	if int32(0) <= v3756 {
		goto L1000
	} else {
		goto L1008
	}
L1008:
	;
	v3782 = v3756
	goto L1
L1009:
	;
	if int32(0) <= v3762 {
		goto L1000
	} else {
		goto L1010
	}
L1010:
	;
	v3782 = v3762
	goto L1
L1011:
	;
	if int32(0) <= v3768 {
		goto L1000
	} else {
		goto L1012
	}
L1012:
	;
	v3782 = v3768
	goto L1
L1013:
	;
	if v3774 < int32(0) {
		v3782 = v3774
		goto L1
	} else {
		goto L1014
	}
L1014:
	;
	goto L1000
}
