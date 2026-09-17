package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PostgresMainLoopOnce(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int64
	_ = v256
	var v264 int64
	_ = v264
	var v272 int64
	_ = v272
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
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
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v671 int64
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int64
	_ = v806
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v843 int32
	_ = v843
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v909 int32
	_ = v909
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
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
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1314 int64
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1560 int64
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1577 int64
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1841 int32
	_ = v1841
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1922 int64
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1950 int64
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int64
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int64
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int64
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int64
	_ = v1969
	var v1974 int32
	_ = v1974
	var v1975 int64
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1984 int64
	_ = v1984
	var v1988 int64
	_ = v1988
	var v1991 int64
	_ = v1991
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 int64
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2086 int32
	_ = v2086
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2126 int32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2135 int64
	_ = v2135
	var v2137 int64
	_ = v2137
	var v2148 int32
	_ = v2148
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2163 int64
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int64
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2194 int64
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int64
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2498 int32
	_ = v2498
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2568 int32
	_ = v2568
	var v2571 int32
	_ = v2571
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2604 int32
	_ = v2604
	var v2610 int32
	_ = v2610
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2644 int32
	_ = v2644
	var v2651 int32
	_ = v2651
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2722 int32
	_ = v2722
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2746 int32
	_ = v2746
	var v2751 int32
	_ = v2751
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2768 int32
	_ = v2768
	var v2772 int32
	_ = v2772
	var v2776 int32
	_ = v2776
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2805 int64
	_ = v2805
	var v2808 int64
	_ = v2808
	var v2811 int32
	_ = v2811
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2828 int32
	_ = v2828
	var v2833 int32
	_ = v2833
	var v2838 int32
	_ = v2838
	var v2843 int32
	_ = v2843
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2900 int32
	_ = v2900
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2930 int32
	_ = v2930
	var v2937 int32
	_ = v2937
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2949 int32
	_ = v2949
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2960 int32
	_ = v2960
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2980 int32
	_ = v2980
	var v2983 int32
	_ = v2983
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3022 int32
	_ = v3022
	var v3028 int32
	_ = v3028
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3059 int32
	_ = v3059
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3073 int32
	_ = v3073
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3084 int32
	_ = v3084
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3101 int32
	_ = v3101
	var v3107 int32
	_ = v3107
	var v3110 int32
	_ = v3110
	var v3114 int32
	_ = v3114
	var v3119 int32
	_ = v3119
	var v3122 int32
	_ = v3122
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3156 int32
	_ = v3156
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3168 int32
	_ = v3168
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3204 int32
	_ = v3204
	var v3209 int32
	_ = v3209
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3219 int32
	_ = v3219
	var v3223 int32
	_ = v3223
	var v3230 int32
	_ = v3230
	var v3233 int32
	_ = v3233
	var v3238 int32
	_ = v3238
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3256 int32
	_ = v3256
	var v3260 int32
	_ = v3260
	var v3265 int32
	_ = v3265
	var v3269 int32
	_ = v3269
	var v3272 int32
	_ = v3272
	var v3276 int32
	_ = v3276
	var v3281 int32
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3292 int32
	_ = v3292
	var v3298 int32
	_ = v3298
	var v3302 int32
	_ = v3302
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3320 int32
	_ = v3320
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3330 int32
	_ = v3330
	var v3335 int64
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3339 int64
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int64
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3347 int64
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3354 int64
	_ = v3354
	var v3359 int32
	_ = v3359
	var v3360 int64
	_ = v3360
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3369 int64
	_ = v3369
	var v3373 int64
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3377 int32
	_ = v3377
	var v3380 int32
	_ = v3380
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3389 int64
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3395 int64
	_ = v3395
	var v3400 int32
	_ = v3400
	var v3401 int64
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3406 int64
	_ = v3406
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3417 int64
	_ = v3417
	var v3423 int32
	_ = v3423
	var v3428 int32
	_ = v3428
	var v3430 int32
	_ = v3430
	var v3436 int32
	_ = v3436
	var v3447 int32
	_ = v3447
	var v3450 int32
	_ = v3450
	var v3454 int32
	_ = v3454
	var v3458 int32
	_ = v3458
	var v3463 int32
	_ = v3463
	var v3467 int32
	_ = v3467
	var v3470 int32
	_ = v3470
	var v3474 int32
	_ = v3474
	var v3479 int32
	_ = v3479
	var v3484 int64
	_ = v3484
	var v3488 int32
	_ = v3488
	var v3494 int32
	_ = v3494
	var v3497 int64
	_ = v3497
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3506 int32
	_ = v3506
	var v3515 int32
	_ = v3515
	var v3521 int32
	_ = v3521
	var v3524 int32
	_ = v3524
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3531 int32
	_ = v3531
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3542 int32
	_ = v3542
	var v3548 int32
	_ = v3548
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3554 int64
	_ = v3554
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3564 int32
	_ = v3564
	var v3571 int32
	_ = v3571
	var v3573 int64
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3580 int32
	_ = v3580
	var v3586 int32
	_ = v3586
	var v3591 int32
	_ = v3591
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3597 int32
	_ = v3597
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3616 int32
	_ = v3616
	var v3618 int32
	_ = v3618
	var v3620 int32
	_ = v3620
	var v3624 int64
	_ = v3624
	var v3626 int32
	_ = v3626
	var v3629 int64
	_ = v3629
	var v3632 int32
	_ = v3632
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3649 int32
	_ = v3649
	var v3654 int32
	_ = v3654
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3659 int64
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3661 int32
	_ = v3661
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3671 int32
	_ = v3671
	var v3673 int32
	_ = v3673
	var v3680 int32
	_ = v3680
	var v3687 int32
	_ = v3687
	var v3688 int64
	_ = v3688
	var v3690 int64
	_ = v3690
	var v3691 int64
	_ = v3691
	var v3695 int64
	_ = v3695
	var v3699 int32
	_ = v3699
	var v3704 int32
	_ = v3704
	var v3707 int32
	_ = v3707
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3721 int32
	_ = v3721
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3729 int32
	_ = v3729
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3748 int32
	_ = v3748
	var v3749 int64
	_ = v3749
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3783 int32
	_ = v3783
	var v3789 int32
	_ = v3789
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3799 int32
	_ = v3799
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3808 int32
	_ = v3808
	var v3810 int32
	_ = v3810
	var v3816 int32
	_ = v3816
	var v3818 int32
	_ = v3818
	var v3819 int32
	_ = v3819
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3827 int64
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3833 int64
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3841 int32
	_ = v3841
	var v3848 int32
	_ = v3848
	var v3850 int32
	_ = v3850
	var v3851 int64
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3868 int32
	_ = v3868
	var v3870 int32
	_ = v3870
	var v3875 int32
	_ = v3875
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3901 int32
	_ = v3901
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3909 int32
	_ = v3909
	var v3911 int32
	_ = v3911
	var v3912 int32
	_ = v3912
	var v3917 int32
	_ = v3917
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3926 int32
	_ = v3926
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3936 int32
	_ = v3936
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3944 int32
	_ = v3944
	var v3946 int32
	_ = v3946
	var v3948 int32
	_ = v3948
	var v3951 int32
	_ = v3951
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3958 int32
	_ = v3958
	var v3963 int32
	_ = v3963
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3972 int32
	_ = v3972
	var v3984 int32
	_ = v3984
	var v3986 int32
	_ = v3986
	var v3987 int32
	_ = v3987
	var v3990 int64
	_ = v3990
	var v3992 int64
	_ = v3992
	var v3995 int64
	_ = v3995
	var v3997 int64
	_ = v3997
	var v4002 int32
	_ = v4002
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4054 int64
	_ = v4054
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4064 int32
	_ = v4064
	var v4066 int32
	_ = v4066
	var v4068 int32
	_ = v4068
	var v4069 int32
	_ = v4069
	var v4078 int32
	_ = v4078
	var v4080 int64
	_ = v4080
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4125 int32
	_ = v4125
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4136 int32
	_ = v4136
	var v4141 int32
	_ = v4141
	var v4142 int32
	_ = v4142
	var v4143 int32
	_ = v4143
	var v4144 int32
	_ = v4144
	var v4148 int32
	_ = v4148
	var v4149 int32
	_ = v4149
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4163 int32
	_ = v4163
	var v4164 int32
	_ = v4164
	var v4172 int32
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4191 int32
	_ = v4191
	var v4199 int32
	_ = v4199
	var v4200 int32
	_ = v4200
	var v4203 int32
	_ = v4203
	var v4210 int32
	_ = v4210
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4226 int32
	_ = v4226
	var v4227 int32
	_ = v4227
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4235 int32
	_ = v4235
	var v4239 int32
	_ = v4239
	var v4241 int32
	_ = v4241
	var v4243 int32
	_ = v4243
	var v4245 int32
	_ = v4245
	var v4246 int32
	_ = v4246
	var v4249 int32
	_ = v4249
	var v4253 int32
	_ = v4253
	var v4254 int32
	_ = v4254
	var v4256 int32
	_ = v4256
	var v4283 int32
	_ = v4283
	var v4284 int32
	_ = v4284
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4297 int32
	_ = v4297
	var v4299 int32
	_ = v4299
	var v4305 int32
	_ = v4305
	var v4308 int32
	_ = v4308
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4315 int32
	_ = v4315
	var v4322 int32
	_ = v4322
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4326 int32
	_ = v4326
	var v4332 int32
	_ = v4332
	var v4334 int32
	_ = v4334
	var v4335 int32
	_ = v4335
	var v4336 int32
	_ = v4336
	var v4337 int32
	_ = v4337
	var v4375 int32
	_ = v4375
	var v4377 int32
	_ = v4377
	var v4382 int32
	_ = v4382
	var v4383 int32
	_ = v4383
	var v4384 int32
	_ = v4384
	var v4386 int32
	_ = v4386
	var v4399 int32
	_ = v4399
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4406 int32
	_ = v4406
	var v4410 int32
	_ = v4410
	var v4411 int32
	_ = v4411
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4415 int32
	_ = v4415
	var v4417 int32
	_ = v4417
	var v4426 int32
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4432 int32
	_ = v4432
	var v4433 int32
	_ = v4433
	var v4434 int32
	_ = v4434
	var v4436 int32
	_ = v4436
	var v4446 int32
	_ = v4446
	var v4452 int32
	_ = v4452
	var v4455 int32
	_ = v4455
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4465 int32
	_ = v4465
	var v4470 int32
	_ = v4470
	var v4471 int32
	_ = v4471
	var v4472 int32
	_ = v4472
	var v4474 int32
	_ = v4474
	var v4476 int32
	_ = v4476
	var v4477 int32
	_ = v4477
	var v4478 int32
	_ = v4478
	var v4481 int32
	_ = v4481
	var v4482 int32
	_ = v4482
	var v4484 int32
	_ = v4484
	var v4487 int32
	_ = v4487
	var v4488 int32
	_ = v4488
	var v4490 int32
	_ = v4490
	var v4492 int32
	_ = v4492
	var v4494 int32
	_ = v4494
	var v4498 int32
	_ = v4498
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4506 int32
	_ = v4506
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4516 int32
	_ = v4516
	var v4523 int32
	_ = v4523
	var v4541 int32
	_ = v4541
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4550 int32
	_ = v4550
	var v4554 int32
	_ = v4554
	var v4559 int32
	_ = v4559
	var v4560 int32
	_ = v4560
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4567 int32
	_ = v4567
	var v4569 int32
	_ = v4569
	var v4573 int32
	_ = v4573
	var v4576 int32
	_ = v4576
	var v4580 int32
	_ = v4580
	var v4585 int32
	_ = v4585
	var v4589 int32
	_ = v4589
	var v4592 int32
	_ = v4592
	var v4598 int32
	_ = v4598
	var v4603 int32
	_ = v4603
	var v4607 int32
	_ = v4607
	var v4610 int32
	_ = v4610
	var v4614 int32
	_ = v4614
	var v4619 int32
	_ = v4619
	var v4623 int32
	_ = v4623
	var v4626 int32
	_ = v4626
	var v4633 int32
	_ = v4633
	var v4638 int32
	_ = v4638
	var v4642 int32
	_ = v4642
	var v4645 int32
	_ = v4645
	var v4649 int32
	_ = v4649
	var v4654 int32
	_ = v4654
	var v4658 int32
	_ = v4658
	var v4661 int32
	_ = v4661
	var v4665 int32
	_ = v4665
	var v4670 int32
	_ = v4670
	var v4674 int32
	_ = v4674
	var v4677 int32
	_ = v4677
	var v4681 int32
	_ = v4681
	var v4686 int32
	_ = v4686
	var v4690 int32
	_ = v4690
	var v4693 int32
	_ = v4693
	var v4697 int32
	_ = v4697
	var v4702 int32
	_ = v4702
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4713 int32
	_ = v4713
	var v4718 int32
	_ = v4718
	var v4722 int32
	_ = v4722
	var v4729 int32
	_ = v4729
	var v4734 int32
	_ = v4734
	var v4738 int32
	_ = v4738
	var v4745 int32
	_ = v4745
	var v4750 int32
	_ = v4750
	var v4754 int32
	_ = v4754
	var v4761 int32
	_ = v4761
	var v4766 int32
	_ = v4766
	var v4770 int32
	_ = v4770
	var v4777 int32
	_ = v4777
	var v4782 int32
	_ = v4782
	var v4786 int32
	_ = v4786
	var v4793 int32
	_ = v4793
	var v4798 int32
	_ = v4798
	var v4802 int32
	_ = v4802
	var v4805 int32
	_ = v4805
	var v4809 int32
	_ = v4809
	var v4814 int32
	_ = v4814
	var v4818 int32
	_ = v4818
	var v4821 int32
	_ = v4821
	var v4825 int32
	_ = v4825
	var v4830 int32
	_ = v4830
	var v4834 int32
	_ = v4834
	var v4835 int32
	_ = v4835
	var v4841 int32
	_ = v4841
	var v4846 int32
	_ = v4846
	var v4849 int32
	_ = v4849
	var v4853 int32
	_ = v4853
	var v4855 int32
	_ = v4855
	var v4863 int32
	_ = v4863
	var v4868 int32
	_ = v4868
	var v4872 int32
	_ = v4872
	var v4874 int32
	_ = v4874
	var v4882 int32
	_ = v4882
	var v4887 int32
	_ = v4887
	var v4891 int32
	_ = v4891
	var v4893 int32
	_ = v4893
	var v4901 int32
	_ = v4901
	var v4906 int32
	_ = v4906
	var v4910 int32
	_ = v4910
	var v4912 int32
	_ = v4912
	var v4920 int32
	_ = v4920
	var v4925 int32
	_ = v4925
	var v4929 int32
	_ = v4929
	var v4932 int32
	_ = v4932
	var v4943 int32
	_ = v4943
	var v4948 int32
	_ = v4948
	var v4952 int32
	_ = v4952
	var v4954 int32
	_ = v4954
	var v4962 int32
	_ = v4962
	var v4967 int32
	_ = v4967
	var v4971 int32
	_ = v4971
	var v4974 int32
	_ = v4974
	var v4978 int32
	_ = v4978
	var v4983 int32
	_ = v4983
	var v4990 int32
	_ = v4990
	var v4993 int32
	_ = v4993
	var v4997 int32
	_ = v4997
	var v5002 int32
	_ = v5002
	var v5006 int32
	_ = v5006
	var v5009 int32
	_ = v5009
	var v5015 int32
	_ = v5015
	var v5020 int32
	_ = v5020
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5027 int32
	_ = v5027
	var v5028 int32
	_ = v5028
	var v5030 int32
	_ = v5030
	var v5032 int32
	_ = v5032
	var v5051 int32
	_ = v5051
	var v5071 int32
	_ = v5071
	var v5075 int32
	_ = v5075
	var v5077 int32
	_ = v5077
	var v5159 int32
	_ = v5159
	var v5161 int32
	_ = v5161
	var v5166 int32
	_ = v5166
	var v5168 int32
	_ = v5168
	var v5173 int32
	_ = v5173
	var v5183 int32
	_ = v5183
	var v5187 int32
	_ = v5187
	var v5189 int32
	_ = v5189
	var v5194 int32
	_ = v5194
	var v5195 int32
	_ = v5195
	var v5196 int32
	_ = v5196
	var v5199 int32
	_ = v5199
	var v5202 int32
	_ = v5202
	var v5205 int32
	_ = v5205
	var v5215 int32
	_ = v5215
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5222 int32
	_ = v5222
	var v5227 int32
	_ = v5227
	var v5229 int32
	_ = v5229
	var v5232 int32
	_ = v5232
	var v5235 int32
	_ = v5235
	var v5272 int32
	_ = v5272
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5278 int32
	_ = v5278
	var v5280 int32
	_ = v5280
	var v5283 int32
	_ = v5283
	var v5284 int32
	_ = v5284
	var v5307 int32
	_ = v5307
	var v5362 int32
	_ = v5362
	var v5365 int32
	_ = v5365
	var v5366 int32
	_ = v5366
	var v5374 int32
	_ = v5374
	var v5376 int32
	_ = v5376
	var v5379 int32
	_ = v5379
	var v5383 int32
	_ = v5383
	var v5392 int32
	_ = v5392
	var v5420 int32
	_ = v5420
	var v5424 int32
	_ = v5424
	var v5425 int32
	_ = v5425
	var v5426 int32
	_ = v5426
	var v5429 int32
	_ = v5429
	var v5431 int32
	_ = v5431
	var v5432 int32
	_ = v5432
	var v5433 int32
	_ = v5433
	var v5435 int32
	_ = v5435
	var v5437 int32
	_ = v5437
	var v5441 int32
	_ = v5441
	var v5442 int32
	_ = v5442
	var v5448 int32
	_ = v5448
	var v5490 int32
	_ = v5490
	var v5520 int32
	_ = v5520
	var v5532 int32
	_ = v5532
	var v5554 int32
	_ = v5554
	var v5562 int32
	_ = v5562
	var v5570 int32
	_ = v5570
	var v5581 int32
	_ = v5581
	var v5610 int32
	_ = v5610
	var v5613 int32
	_ = v5613
	var v5614 int32
	_ = v5614
	var v5619 int32
	_ = v5619
	var v5623 int32
	_ = v5623
	var v5630 int64
	_ = v5630
	var v5634 int32
	_ = v5634
	var v5636 int32
	_ = v5636
	var v5637 int32
	_ = v5637
	var v5640 int32
	_ = v5640
	var v5651 int32
	_ = v5651
	var v5659 int32
	_ = v5659
	var v5663 int32
	_ = v5663
	var v5670 int64
	_ = v5670
	var v5674 int32
	_ = v5674
	var v5676 int32
	_ = v5676
	var v5677 int32
	_ = v5677
	var v5680 int32
	_ = v5680
	var v5691 int32
	_ = v5691
	var v5695 int32
	_ = v5695
	var v5696 int32
	_ = v5696
	var v5697 int32
	_ = v5697
	var v5702 int32
	_ = v5702
	var v5707 int32
	_ = v5707
	var v5708 int32
	_ = v5708
	var v5717 int32
	_ = v5717
	var v5720 int32
	_ = v5720
	var v5723 int32
	_ = v5723
	var v5733 int32
	_ = v5733
	var v5736 int32
	_ = v5736
	var v5740 int32
	_ = v5740
	var v5742 int32
	_ = v5742
	var v5747 int32
	_ = v5747
	var v5750 int32
	_ = v5750
	var v5752 int32
	_ = v5752
	var v5757 int32
	_ = v5757
	var v5758 int32
	_ = v5758
	var v5764 int32
	_ = v5764
	var v5766 int32
	_ = v5766
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5774 int32
	_ = v5774
	var v5775 int32
	_ = v5775
	var v5776 int32
	_ = v5776
	var v5778 int32
	_ = v5778
	var v5781 int32
	_ = v5781
	var v5783 int32
	_ = v5783
	var v5784 int32
	_ = v5784
	var v5785 int32
	_ = v5785
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5796 int32
	_ = v5796
	var v5797 int32
	_ = v5797
	var v5798 int32
	_ = v5798
	var v5799 int32
	_ = v5799
	var v5803 int32
	_ = v5803
	var v5806 int32
	_ = v5806
	var v5816 int32
	_ = v5816
	var v5819 int32
	_ = v5819
	var v5822 int32
	_ = v5822
	var v5823 int32
	_ = v5823
	var v5825 int32
	_ = v5825
	var v5830 int32
	_ = v5830
	var v5831 int32
	_ = v5831
	var v5832 int32
	_ = v5832
	var v5835 int32
	_ = v5835
	var v5836 int32
	_ = v5836
	var v5838 int32
	_ = v5838
	var v5840 int32
	_ = v5840
	var v5842 int32
	_ = v5842
	var v5844 int32
	_ = v5844
	var v5846 int32
	_ = v5846
	var v5847 int32
	_ = v5847
	var v5848 int32
	_ = v5848
	var v5862 int32
	_ = v5862
	var v5866 int32
	_ = v5866
	var v5867 int32
	_ = v5867
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5873 int32
	_ = v5873
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5876 int32
	_ = v5876
	var v5879 int32
	_ = v5879
	var v5884 int32
	_ = v5884
	var v5891 int32
	_ = v5891
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5903 int32
	_ = v5903
	var v5904 int32
	_ = v5904
	var v5905 int32
	_ = v5905
	var v5907 int32
	_ = v5907
	var v5910 int32
	_ = v5910
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5921 int32
	_ = v5921
	var v5922 int32
	_ = v5922
	var v5930 int32
	_ = v5930
	var v5933 int32
	_ = v5933
	var v5935 int32
	_ = v5935
	var v5938 int32
	_ = v5938
	var v5939 int32
	_ = v5939
	var v5945 int32
	_ = v5945
	var v5948 int32
	_ = v5948
	var v5950 int32
	_ = v5950
	var v5952 int32
	_ = v5952
	var v5956 int32
	_ = v5956
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5965 int32
	_ = v5965
	var v5970 int32
	_ = v5970
	var v5972 int32
	_ = v5972
	var v5974 int32
	_ = v5974
	var v5975 int32
	_ = v5975
	var v6006 int32
	_ = v6006
	var v6017 int32
	_ = v6017
	var v6020 int32
	_ = v6020
	var v6022 int32
	_ = v6022
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6029 int32
	_ = v6029
	var v6071 int32
	_ = v6071
	var v6074 int32
	_ = v6074
	var v6075 int32
	_ = v6075
	var v6077 int32
	_ = v6077
	var v6081 int32
	_ = v6081
	var v6083 int32
	_ = v6083
	var v6115 int32
	_ = v6115
	var v6123 int32
	_ = v6123
	var v6126 int32
	_ = v6126
	var v6127 int32
	_ = v6127
	var v6132 int32
	_ = v6132
	var v6133 int32
	_ = v6133
	var v6141 int32
	_ = v6141
	var v6143 int32
	_ = v6143
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6159 int32
	_ = v6159
	var v6161 int32
	_ = v6161
	var v6162 int32
	_ = v6162
	var v6163 int32
	_ = v6163
	var v6167 int32
	_ = v6167
	var v6176 int32
	_ = v6176
	var v6204 int32
	_ = v6204
	var v6208 int32
	_ = v6208
	var v6209 int32
	_ = v6209
	var v6210 int32
	_ = v6210
	var v6213 int32
	_ = v6213
	var v6215 int32
	_ = v6215
	var v6216 int32
	_ = v6216
	var v6217 int32
	_ = v6217
	var v6219 int32
	_ = v6219
	var v6221 int32
	_ = v6221
	var v6223 int32
	_ = v6223
	var v6224 int32
	_ = v6224
	var v6230 int32
	_ = v6230
	var v6249 int32
	_ = v6249
	var v6271 int32
	_ = v6271
	var v6311 int32
	_ = v6311
	var v6356 int32
	_ = v6356
	var v6359 int32
	_ = v6359
	var v6363 int32
	_ = v6363
	var v6367 int64
	_ = v6367
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
	var v6376 int32
	_ = v6376
	var v6377 int32
	_ = v6377
	var v6383 int32
	_ = v6383
	var v6384 int32
	_ = v6384
	var v6391 int32
	_ = v6391
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6432 int32
	_ = v6432
	var v6444 int32
	_ = v6444
	var v6474 int32
	_ = v6474
	var v6480 int32
	_ = v6480
	var v6485 int32
	_ = v6485
	var v6495 int32
	_ = v6495
	var v6500 int32
	_ = v6500
	var v6501 int32
	_ = v6501
	var v6502 int32
	_ = v6502
	var v6505 int32
	_ = v6505
	var v6511 int32
	_ = v6511
	var v6516 int32
	_ = v6516
	var v6519 int32
	_ = v6519
	var v6520 int32
	_ = v6520
	var v6522 int32
	_ = v6522
	var v6525 int32
	_ = v6525
	var v6530 int32
	_ = v6530
	var v6532 int32
	_ = v6532
	var v6537 int32
	_ = v6537
	var v6538 int32
	_ = v6538
	var v6540 int32
	_ = v6540
	var v6541 int32
	_ = v6541
	var v6542 int32
	_ = v6542
	var v6543 int32
	_ = v6543
	var v6547 int32
	_ = v6547
	var v6550 int32
	_ = v6550
	var v6560 int32
	_ = v6560
	var v6564 int32
	_ = v6564
	var v6565 int32
	_ = v6565
	var v6567 int32
	_ = v6567
	var v6572 int32
	_ = v6572
	var v6573 int32
	_ = v6573
	var v6576 int32
	_ = v6576
	var v6577 int32
	_ = v6577
	var v6579 int32
	_ = v6579
	var v6580 int32
	_ = v6580
	var v6587 int32
	_ = v6587
	var v6590 int32
	_ = v6590
	var v6593 int32
	_ = v6593
	var v6598 int32
	_ = v6598
	var v6599 int32
	_ = v6599
	var v6600 int32
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6604 int32
	_ = v6604
	var v6605 int32
	_ = v6605
	var v6609 int32
	_ = v6609
	var v6616 int32
	_ = v6616
	var v6617 int32
	_ = v6617
	var v6618 int32
	_ = v6618
	var v6619 int32
	_ = v6619
	var v6621 int32
	_ = v6621
	var v6626 int32
	_ = v6626
	var v6627 int32
	_ = v6627
	var v6629 int32
	_ = v6629
	var v6630 int32
	_ = v6630
	var v6633 int32
	_ = v6633
	var v6634 int32
	_ = v6634
	var v6635 int32
	_ = v6635
	var v6638 int32
	_ = v6638
	var v6639 int32
	_ = v6639
	var v6641 int32
	_ = v6641
	var v6645 int32
	_ = v6645
	var v6649 int32
	_ = v6649
	var v6650 int32
	_ = v6650
	var v6655 int32
	_ = v6655
	var v6662 int32
	_ = v6662
	var v6674 int32
	_ = v6674
	var v6675 int32
	_ = v6675
	var v6676 int32
	_ = v6676
	var v6681 int32
	_ = v6681
	var v6683 int32
	_ = v6683
	var v6685 int32
	_ = v6685
	var v6688 int32
	_ = v6688
	var v6690 int32
	_ = v6690
	var v6696 int32
	_ = v6696
	var v6698 int32
	_ = v6698
	var v6703 int32
	_ = v6703
	var v6707 int32
	_ = v6707
	var v6708 int32
	_ = v6708
	var v6713 int32
	_ = v6713
	var v6714 int32
	_ = v6714
	var v6724 int32
	_ = v6724
	var v6728 int32
	_ = v6728
	var v6729 int32
	_ = v6729
	var v6732 int32
	_ = v6732
	var v6735 int32
	_ = v6735
	var v6744 int32
	_ = v6744
	var v6747 int32
	_ = v6747
	var v6749 int32
	_ = v6749
	var v6753 int32
	_ = v6753
	var v6757 int32
	_ = v6757
	var v6762 int32
	_ = v6762
	var v6766 int32
	_ = v6766
	var v6770 int64
	_ = v6770
	var v6773 int32
	_ = v6773
	var v6775 int32
	_ = v6775
	var v6776 int32
	_ = v6776
	var v6777 int32
	_ = v6777
	var v6778 int32
	_ = v6778
	var v6779 int32
	_ = v6779
	var v6782 int32
	_ = v6782
	var v6783 int32
	_ = v6783
	var v6784 int32
	_ = v6784
	var v6786 int32
	_ = v6786
	var v6787 int32
	_ = v6787
	var v6790 int32
	_ = v6790
	var v6796 int32
	_ = v6796
	var v6801 int32
	_ = v6801
	var v6803 int32
	_ = v6803
	var v6805 int32
	_ = v6805
	var v6806 int32
	_ = v6806
	var v6807 int32
	_ = v6807
	var v6809 int32
	_ = v6809
	var v6812 int32
	_ = v6812
	var v6814 int32
	_ = v6814
	var v6818 int32
	_ = v6818
	var v6821 int32
	_ = v6821
	var v6824 int32
	_ = v6824
	var v6832 int32
	_ = v6832
	var v6866 int32
	_ = v6866
	var v6867 int64
	_ = v6867
	var v6871 int32
	_ = v6871
	var v6876 int32
	_ = v6876
	var v6880 int32
	_ = v6880
	var v6887 int64
	_ = v6887
	var v6891 int32
	_ = v6891
	var v6893 int32
	_ = v6893
	var v6894 int32
	_ = v6894
	var v6897 int32
	_ = v6897
	var v6908 int32
	_ = v6908
	var v6951 int32
	_ = v6951
	var v6961 int32
	_ = v6961
	var v6965 int32
	_ = v6965
	var v6968 int32
	_ = v6968
	var v6973 int32
	_ = v6973
	var v6974 int32
	_ = v6974
	var v6980 int32
	_ = v6980
	var v6981 int32
	_ = v6981
	var v6988 int32
	_ = v6988
	var v7025 int32
	_ = v7025
	var v7026 int32
	_ = v7026
	var v7029 int32
	_ = v7029
	var v7032 int32
	_ = v7032
	var v7071 int32
	_ = v7071
	var v7072 int32
	_ = v7072
	var v7077 int32
	_ = v7077
	var v7080 int32
	_ = v7080
	var v7081 int32
	_ = v7081
	var v7088 int32
	_ = v7088
	var v7091 int32
	_ = v7091
	var v7094 int32
	_ = v7094
	var v7097 int32
	_ = v7097
	var v7104 int32
	_ = v7104
	var v7107 int32
	_ = v7107
	var v7109 int32
	_ = v7109
	var v7110 int32
	_ = v7110
	var v7111 int32
	_ = v7111
	var v7113 int32
	_ = v7113
	var v7114 int32
	_ = v7114
	var v7115 int32
	_ = v7115
	var v7116 int32
	_ = v7116
	var v7117 int32
	_ = v7117
	var v7119 int32
	_ = v7119
	var v7121 int32
	_ = v7121
	var v7122 int32
	_ = v7122
	var v7123 int32
	_ = v7123
	var v7124 int32
	_ = v7124
	var v7125 int32
	_ = v7125
	var v7126 int32
	_ = v7126
	var v7127 int32
	_ = v7127
	var v7130 int32
	_ = v7130
	var v7131 int32
	_ = v7131
	var v7137 int32
	_ = v7137
	var v7138 int32
	_ = v7138
	var v7142 int32
	_ = v7142
	var v7145 int32
	_ = v7145
	var v7146 int32
	_ = v7146
	var v7148 int32
	_ = v7148
	var v7149 int32
	_ = v7149
	var v7150 int32
	_ = v7150
	var v7152 int32
	_ = v7152
	var v7153 int32
	_ = v7153
	var v7157 int32
	_ = v7157
	var v7158 int32
	_ = v7158
	var v7171 int32
	_ = v7171
	var v7172 int32
	_ = v7172
	var v7175 int32
	_ = v7175
	var v7184 int32
	_ = v7184
	var v7202 int32
	_ = v7202
	var v7216 int32
	_ = v7216
	var v7217 int32
	_ = v7217
	var v7219 int32
	_ = v7219
	var v7226 int32
	_ = v7226
	var v7227 int32
	_ = v7227
	var v7229 int32
	_ = v7229
	var v7230 int32
	_ = v7230
	var v7234 int32
	_ = v7234
	var v7235 int32
	_ = v7235
	var v7236 int32
	_ = v7236
	var v7237 int32
	_ = v7237
	var v7238 int32
	_ = v7238
	var v7243 int32
	_ = v7243
	var v7244 int32
	_ = v7244
	var v7252 int32
	_ = v7252
	var v7253 int32
	_ = v7253
	var v7260 int32
	_ = v7260
	var v7262 int32
	_ = v7262
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7265 int32
	_ = v7265
	var v7267 int32
	_ = v7267
	var v7268 int32
	_ = v7268
	var v7270 int32
	_ = v7270
	var v7271 int32
	_ = v7271
	var v7272 int32
	_ = v7272
	var v7277 int32
	_ = v7277
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7282 int32
	_ = v7282
	var v7286 int32
	_ = v7286
	var v7287 int32
	_ = v7287
	var v7289 int32
	_ = v7289
	var v7290 int32
	_ = v7290
	var v7291 int32
	_ = v7291
	var v7294 int32
	_ = v7294
	var v7295 int32
	_ = v7295
	var v7298 int32
	_ = v7298
	var v7299 int32
	_ = v7299
	var v7300 int32
	_ = v7300
	var v7308 int32
	_ = v7308
	var v7309 int32
	_ = v7309
	var v7312 int32
	_ = v7312
	var v7318 int32
	_ = v7318
	var v7319 int32
	_ = v7319
	var v7323 int32
	_ = v7323
	var v7324 int32
	_ = v7324
	var v7326 int32
	_ = v7326
	var v7327 int32
	_ = v7327
	var v7328 int32
	_ = v7328
	var v7329 int32
	_ = v7329
	var v7335 int32
	_ = v7335
	var v7336 int32
	_ = v7336
	var v7339 int32
	_ = v7339
	var v7344 int32
	_ = v7344
	var v7346 int32
	_ = v7346
	var v7351 int32
	_ = v7351
	var v7353 int32
	_ = v7353
	var v7355 int32
	_ = v7355
	var v7356 int32
	_ = v7356
	var v7359 int32
	_ = v7359
	var v7362 int32
	_ = v7362
	var v7363 int32
	_ = v7363
	var v7394 int32
	_ = v7394
	var v7432 int32
	_ = v7432
	var v7440 int32
	_ = v7440
	var v7443 int32
	_ = v7443
	var v7446 int32
	_ = v7446
	var v7447 int32
	_ = v7447
	var v7462 int32
	_ = v7462
	var v7463 int32
	_ = v7463
	var v7469 int32
	_ = v7469
	var v7470 int32
	_ = v7470
	var v7477 int32
	_ = v7477
	var v7514 int32
	_ = v7514
	var v7515 int32
	_ = v7515
	var v7518 int32
	_ = v7518
	var v7530 int32
	_ = v7530
	var v7560 int32
	_ = v7560
	var v7561 int32
	_ = v7561
	var v7563 int32
	_ = v7563
	var v7564 int32
	_ = v7564
	var v7565 int32
	_ = v7565
	var v7566 int32
	_ = v7566
	var v7577 int32
	_ = v7577
	var v7580 int32
	_ = v7580
	var v7583 int32
	_ = v7583
	var v7591 int32
	_ = v7591
	var v7625 int32
	_ = v7625
	var v7626 int64
	_ = v7626
	var v7630 int32
	_ = v7630
	var v7635 int32
	_ = v7635
	var v7639 int32
	_ = v7639
	var v7646 int64
	_ = v7646
	var v7650 int32
	_ = v7650
	var v7652 int32
	_ = v7652
	var v7653 int32
	_ = v7653
	var v7656 int32
	_ = v7656
	var v7667 int32
	_ = v7667
	var v7709 int32
	_ = v7709
	var v7710 int32
	_ = v7710
	var v7713 int32
	_ = v7713
	var v7715 int32
	_ = v7715
	var v7716 int32
	_ = v7716
	var v7718 int32
	_ = v7718
	var v7719 int32
	_ = v7719
	var v7722 int32
	_ = v7722
	var v7727 int32
	_ = v7727
	var v7731 int32
	_ = v7731
	var v7732 int32
	_ = v7732
	var v7737 int32
	_ = v7737
	var v7738 int32
	_ = v7738
	var v7748 int32
	_ = v7748
	var v7750 int32
	_ = v7750
	var v7754 int32
	_ = v7754
	var v7755 int32
	_ = v7755
	var v7758 int32
	_ = v7758
	var v7759 int32
	_ = v7759
	var v7760 int32
	_ = v7760
	var v7763 int32
	_ = v7763
	var v7767 int32
	_ = v7767
	var v7770 int32
	_ = v7770
	var v7779 int32
	_ = v7779
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7785 int32
	_ = v7785
	var v7789 int32
	_ = v7789
	var v7793 int32
	_ = v7793
	var v7794 int32
	_ = v7794
	var v7797 int32
	_ = v7797
	var v7805 int32
	_ = v7805
	var v7808 int32
	_ = v7808
	var v7812 int32
	_ = v7812
	var v7820 int32
	_ = v7820
	var v7825 int32
	_ = v7825
	var v7829 int32
	_ = v7829
	var v7833 int64
	_ = v7833
	var v7836 int32
	_ = v7836
	var v7837 int32
	_ = v7837
	var v7838 int32
	_ = v7838
	var v7840 int32
	_ = v7840
	var v7841 int32
	_ = v7841
	var v7843 int32
	_ = v7843
	var v7845 int32
	_ = v7845
	var v7847 int32
	_ = v7847
	var v7848 int32
	_ = v7848
	var v7849 int32
	_ = v7849
	var v7855 int32
	_ = v7855
	var v7856 int32
	_ = v7856
	var v7860 int32
	_ = v7860
	var v7861 int32
	_ = v7861
	var v7864 int32
	_ = v7864
	var v7867 int32
	_ = v7867
	var v7868 int32
	_ = v7868
	var v7869 int32
	_ = v7869
	var v7873 int32
	_ = v7873
	var v7874 int32
	_ = v7874
	var v7880 int32
	_ = v7880
	var v7881 int32
	_ = v7881
	var v7882 int32
	_ = v7882
	var v7883 int32
	_ = v7883
	var v7884 int32
	_ = v7884
	var v7885 int32
	_ = v7885
	var v7886 int32
	_ = v7886
	var v7888 int32
	_ = v7888
	var v7889 int32
	_ = v7889
	var v7894 int32
	_ = v7894
	var v7897 int32
	_ = v7897
	var v7900 int32
	_ = v7900
	var v7903 int32
	_ = v7903
	var v7904 int32
	_ = v7904
	var v7912 int32
	_ = v7912
	var v7946 int32
	_ = v7946
	var v7947 int64
	_ = v7947
	var v7951 int32
	_ = v7951
	var v7956 int32
	_ = v7956
	var v7960 int32
	_ = v7960
	var v7967 int64
	_ = v7967
	var v7971 int32
	_ = v7971
	var v7973 int32
	_ = v7973
	var v7974 int32
	_ = v7974
	var v7977 int32
	_ = v7977
	var v7988 int32
	_ = v7988
	var v7992 int32
	_ = v7992
	var v7995 int32
	_ = v7995
	var v7996 int32
	_ = v7996
	var v7999 int32
	_ = v7999
	var v8000 int32
	_ = v8000
	var v8001 int32
	_ = v8001
	var v8003 int32
	_ = v8003
	var v8006 int32
	_ = v8006
	var v8014 int32
	_ = v8014
	var v8048 int32
	_ = v8048
	var v8049 int64
	_ = v8049
	var v8053 int32
	_ = v8053
	var v8058 int32
	_ = v8058
	var v8062 int32
	_ = v8062
	var v8069 int64
	_ = v8069
	var v8073 int32
	_ = v8073
	var v8075 int32
	_ = v8075
	var v8076 int32
	_ = v8076
	var v8079 int32
	_ = v8079
	var v8090 int32
	_ = v8090
	var v8131 int32
	_ = v8131
	var v8136 int32
	_ = v8136
	var v8142 int32
	_ = v8142
	var v8152 int32
	_ = v8152
	var v8156 int32
	_ = v8156
	var v8157 int32
	_ = v8157
	var v8162 int32
	_ = v8162
	var v8163 int32
	_ = v8163
	var v8164 int32
	_ = v8164
	var v8166 int32
	_ = v8166
	var v8167 int32
	_ = v8167
	var v8170 int32
	_ = v8170
	var v8179 int32
	_ = v8179
	var v8210 int32
	_ = v8210
	var v8214 int32
	_ = v8214
	var v8215 int32
	_ = v8215
	var v8216 int32
	_ = v8216
	var v8218 int32
	_ = v8218
	var v8221 int32
	_ = v8221
	var v8222 int32
	_ = v8222
	var v8264 int32
	_ = v8264
	var v8265 int32
	_ = v8265
	var v8269 int32
	_ = v8269
	var v8273 int32
	_ = v8273
	var v8277 int32
	_ = v8277
	var v8279 int32
	_ = v8279
	var v8284 int32
	_ = v8284
	var v8290 int32
	_ = v8290
	var v8292 int32
	_ = v8292
	var v8295 int32
	_ = v8295
	var v8299 int32
	_ = v8299
	var v8303 int32
	_ = v8303
	var v8304 int32
	_ = v8304
	var v8307 int32
	_ = v8307
	var v8315 int32
	_ = v8315
	var v8321 int32
	_ = v8321
	var v8328 int32
	_ = v8328
	var v8360 int32
	_ = v8360
	var v8361 int32
	_ = v8361
	var v8368 int32
	_ = v8368
	var v8371 int32
	_ = v8371
	var v8374 int32
	_ = v8374
	var v8375 int32
	_ = v8375
	var v8376 int32
	_ = v8376
	var v8379 int32
	_ = v8379
	var v8382 int32
	_ = v8382
	var v8385 int32
	_ = v8385
	var v8392 int32
	_ = v8392
	var v8394 int32
	_ = v8394
	var v8395 int32
	_ = v8395
	var v8398 int32
	_ = v8398
	var v8399 int32
	_ = v8399
	var v8413 int32
	_ = v8413
	var v8417 int32
	_ = v8417
	var v8418 int32
	_ = v8418
	var v8419 int32
	_ = v8419
	var v8421 int32
	_ = v8421
	var v8422 int32
	_ = v8422
	var v8424 int32
	_ = v8424
	var v8425 int32
	_ = v8425
	var v8430 int32
	_ = v8430
	var v8438 int32
	_ = v8438
	var v8441 int32
	_ = v8441
	var v8442 int32
	_ = v8442
	var v8444 int32
	_ = v8444
	var v8448 int32
	_ = v8448
	var v8450 int32
	_ = v8450
	var v8453 int32
	_ = v8453
	var v8454 int32
	_ = v8454
	var v8456 int32
	_ = v8456
	var v8463 int32
	_ = v8463
	var v8468 int32
	_ = v8468
	var v8469 int32
	_ = v8469
	var v8473 int32
	_ = v8473
	var v8475 int32
	_ = v8475
	var v8480 int32
	_ = v8480
	var v8481 int32
	_ = v8481
	var v8483 int32
	_ = v8483
	var v8487 int32
	_ = v8487
	var v8490 int32
	_ = v8490
	var v8491 int32
	_ = v8491
	var v8496 int32
	_ = v8496
	var v8497 int32
	_ = v8497
	var v8507 int32
	_ = v8507
	var v8509 int32
	_ = v8509
	var v8513 int32
	_ = v8513
	var v8514 int32
	_ = v8514
	var v8517 int32
	_ = v8517
	var v8520 int32
	_ = v8520
	var v8525 int32
	_ = v8525
	var v8531 int32
	_ = v8531
	var v8540 int32
	_ = v8540
	var v8542 int32
	_ = v8542
	var v8543 int32
	_ = v8543
	var v8546 int32
	_ = v8546
	var v8550 int32
	_ = v8550
	var v8554 int32
	_ = v8554
	var v8555 int32
	_ = v8555
	var v8558 int32
	_ = v8558
	var v8566 int32
	_ = v8566
	var v8568 int32
	_ = v8568
	var v8572 int32
	_ = v8572
	var v8579 int32
	_ = v8579
	var v8584 int32
	_ = v8584
	var v8588 int32
	_ = v8588
	var v8592 int64
	_ = v8592
	var v8598 int32
	_ = v8598
	var v8601 int32
	_ = v8601
	var v8604 int32
	_ = v8604
	var v8605 int32
	_ = v8605
	var v8607 int32
	_ = v8607
	var v8610 int32
	_ = v8610
	var v8611 int32
	_ = v8611
	var v8620 int32
	_ = v8620
	var v8621 int32
	_ = v8621
	var v8623 int32
	_ = v8623
	var v8625 int32
	_ = v8625
	var v8626 int32
	_ = v8626
	var v8633 int32
	_ = v8633
	var v8634 int32
	_ = v8634
	var v8637 int32
	_ = v8637
	var v8638 int32
	_ = v8638
	var v8639 int32
	_ = v8639
	var v8640 int32
	_ = v8640
	var v8643 int32
	_ = v8643
	var v8646 int32
	_ = v8646
	var v8649 int32
	_ = v8649
	var v8651 int32
	_ = v8651
	var v8654 int32
	_ = v8654
	var v8655 int32
	_ = v8655
	var v8657 int32
	_ = v8657
	var v8662 int32
	_ = v8662
	var v8664 int32
	_ = v8664
	var v8666 int32
	_ = v8666
	var v8673 int32
	_ = v8673
	var v8677 int32
	_ = v8677
	var v8689 int32
	_ = v8689
	var v8690 int32
	_ = v8690
	var v8691 int32
	_ = v8691
	var v8693 int32
	_ = v8693
	var v8697 int32
	_ = v8697
	var v8698 int32
	_ = v8698
	var v8700 int32
	_ = v8700
	var v8701 int32
	_ = v8701
	var v8702 int32
	_ = v8702
	var v8704 int32
	_ = v8704
	var v8710 int32
	_ = v8710
	var v8711 int32
	_ = v8711
	var v8712 int32
	_ = v8712
	var v8713 int32
	_ = v8713
	var v8716 int32
	_ = v8716
	var v8723 int32
	_ = v8723
	var v8724 int32
	_ = v8724
	var v8725 int32
	_ = v8725
	var v8728 int32
	_ = v8728
	var v8731 int32
	_ = v8731
	var v8736 int32
	_ = v8736
	var v8737 int32
	_ = v8737
	var v8739 int32
	_ = v8739
	var v8741 int32
	_ = v8741
	var v8745 int32
	_ = v8745
	var v8746 int32
	_ = v8746
	var v8747 int32
	_ = v8747
	var v8752 int32
	_ = v8752
	var v8753 int32
	_ = v8753
	var v8754 int32
	_ = v8754
	var v8757 int32
	_ = v8757
	var v8758 int32
	_ = v8758
	var v8759 int32
	_ = v8759
	var v8761 int32
	_ = v8761
	var v8765 int32
	_ = v8765
	var v8766 int32
	_ = v8766
	var v8768 int32
	_ = v8768
	var v8770 int32
	_ = v8770
	var v8772 int32
	_ = v8772
	var v8773 int32
	_ = v8773
	var v8776 int32
	_ = v8776
	var v8783 int32
	_ = v8783
	var v8787 int32
	_ = v8787
	var v8789 int32
	_ = v8789
	var v8792 int32
	_ = v8792
	var v8797 int32
	_ = v8797
	var v8798 int32
	_ = v8798
	var v8807 int32
	_ = v8807
	var v8812 int32
	_ = v8812
	var v8814 int32
	_ = v8814
	var v8816 int32
	_ = v8816
	var v8818 int32
	_ = v8818
	var v8819 int32
	_ = v8819
	var v8821 int32
	_ = v8821
	var v8822 int32
	_ = v8822
	var v8823 int32
	_ = v8823
	var v8825 int32
	_ = v8825
	var v8827 int32
	_ = v8827
	var v8828 int32
	_ = v8828
	var v8830 int32
	_ = v8830
	var v8831 int32
	_ = v8831
	var v8834 int32
	_ = v8834
	var v8836 int32
	_ = v8836
	var v8837 int32
	_ = v8837
	var v8839 int32
	_ = v8839
	var v8840 int32
	_ = v8840
	var v8842 int32
	_ = v8842
	var v8845 int32
	_ = v8845
	var v8847 int32
	_ = v8847
	var v8848 int64
	_ = v8848
	var v8854 int32
	_ = v8854
	var v8855 int32
	_ = v8855
	var v8860 int32
	_ = v8860
	var v8861 int32
	_ = v8861
	var v8872 int32
	_ = v8872
	var v8903 int32
	_ = v8903
	var v8904 int32
	_ = v8904
	var v8907 int32
	_ = v8907
	var v8909 int32
	_ = v8909
	var v8947 int32
	_ = v8947
	var v8948 int32
	_ = v8948
	var v8951 int32
	_ = v8951
	var v8961 int32
	_ = v8961
	var v8965 int32
	_ = v8965
	var v8979 int32
	_ = v8979
	var v9007 int32
	_ = v9007
	var v9010 int32
	_ = v9010
	var v9012 int32
	_ = v9012
	var v9013 int32
	_ = v9013
	var v9016 int32
	_ = v9016
	var v9018 int32
	_ = v9018
	var v9023 int32
	_ = v9023
	var v9024 int32
	_ = v9024
	var v9025 int32
	_ = v9025
	var v9031 int32
	_ = v9031
	var v9032 int32
	_ = v9032
	var v9034 int32
	_ = v9034
	var v9043 int32
	_ = v9043
	var v9044 int32
	_ = v9044
	var v9049 int32
	_ = v9049
	var v9055 int32
	_ = v9055
	var v9059 int32
	_ = v9059
	var v9060 int32
	_ = v9060
	var v9061 int32
	_ = v9061
	var v9062 int32
	_ = v9062
	var v9064 int32
	_ = v9064
	var v9065 int32
	_ = v9065
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9072 int32
	_ = v9072
	var v9075 int32
	_ = v9075
	var v9079 int32
	_ = v9079
	var v9085 int32
	_ = v9085
	var v9087 int32
	_ = v9087
	var v9092 int32
	_ = v9092
	var v9093 int32
	_ = v9093
	var v9094 int32
	_ = v9094
	var v9096 int32
	_ = v9096
	var v9097 int32
	_ = v9097
	var v9099 int32
	_ = v9099
	var v9100 int32
	_ = v9100
	var v9104 int32
	_ = v9104
	var v9144 int32
	_ = v9144
	var v9145 int32
	_ = v9145
	var v9147 int32
	_ = v9147
	var v9148 int32
	_ = v9148
	var v9151 int32
	_ = v9151
	var v9165 int32
	_ = v9165
	var v9197 int32
	_ = v9197
	var v9201 int32
	_ = v9201
	var v9203 int32
	_ = v9203
	var v9209 int32
	_ = v9209
	var v9212 int32
	_ = v9212
	var v9216 int32
	_ = v9216
	var v9221 int32
	_ = v9221
	var v9225 int32
	_ = v9225
	var v9228 int32
	_ = v9228
	var v9232 int32
	_ = v9232
	var v9237 int32
	_ = v9237
	var v9241 int32
	_ = v9241
	var v9244 int32
	_ = v9244
	var v9252 int32
	_ = v9252
	var v9257 int32
	_ = v9257
	var v9261 int32
	_ = v9261
	var v9271 int32
	_ = v9271
	var v9276 int32
	_ = v9276
	var v9280 int32
	_ = v9280
	var v9283 int32
	_ = v9283
	var v9285 int32
	_ = v9285
	var v9291 int32
	_ = v9291
	var v9296 int32
	_ = v9296
	var v9300 int32
	_ = v9300
	var v9303 int32
	_ = v9303
	var v9310 int32
	_ = v9310
	var v9315 int32
	_ = v9315
	var v9319 int32
	_ = v9319
	var v9322 int32
	_ = v9322
	var v9328 int32
	_ = v9328
	var v9333 int32
	_ = v9333
	var v9337 int32
	_ = v9337
	var v9340 int32
	_ = v9340
	var v9348 int32
	_ = v9348
	var v9353 int32
	_ = v9353
	var v9357 int32
	_ = v9357
	var v9360 int32
	_ = v9360
	var v9367 int32
	_ = v9367
	var v9372 int32
	_ = v9372
	var v9412 int32
	_ = v9412
	var v9413 int32
	_ = v9413
	var v9414 int32
	_ = v9414
	var v9452 int32
	_ = v9452
	var v9454 int32
	_ = v9454
	var v9456 int32
	_ = v9456
	var v9457 int32
	_ = v9457
	var v9458 int32
	_ = v9458
	var v9460 int32
	_ = v9460
	var v9463 int32
	_ = v9463
	var v9468 int32
	_ = v9468
	var v9469 int32
	_ = v9469
	var v9470 int32
	_ = v9470
	var v9484 int32
	_ = v9484
	var v9485 int32
	_ = v9485
	var v9486 int32
	_ = v9486
	var v9488 int32
	_ = v9488
	var v9491 int32
	_ = v9491
	var v9493 int32
	_ = v9493
	var v9496 int32
	_ = v9496
	var v9497 int64
	_ = v9497
	var v9501 int32
	_ = v9501
	var v9505 int32
	_ = v9505
	var v9509 int32
	_ = v9509
	var v9510 int32
	_ = v9510
	var v9511 int32
	_ = v9511
	var v9512 int32
	_ = v9512
	var v9515 int32
	_ = v9515
	var v9518 int32
	_ = v9518
	var v9519 int32
	_ = v9519
	var v9520 int32
	_ = v9520
	var v9527 int32
	_ = v9527
	var v9528 int32
	_ = v9528
	var v9532 int32
	_ = v9532
	var v9537 int32
	_ = v9537
	var v9538 int32
	_ = v9538
	var v9540 int32
	_ = v9540
	var v9543 int32
	_ = v9543
	var v9544 int32
	_ = v9544
	var v9545 int32
	_ = v9545
	var v9547 int32
	_ = v9547
	var v9549 int32
	_ = v9549
	var v9550 int32
	_ = v9550
	var v9551 int32
	_ = v9551
	var v9566 int32
	_ = v9566
	var v9572 int32
	_ = v9572
	var v9574 int32
	_ = v9574
	var v9578 int32
	_ = v9578
	var v9581 int32
	_ = v9581
	var v9588 int32
	_ = v9588
	var v9593 int32
	_ = v9593
	var v9599 int32
	_ = v9599
	var v9602 int32
	_ = v9602
	var v9603 int32
	_ = v9603
	var v9604 int32
	_ = v9604
	var v9605 int32
	_ = v9605
	var v9607 int32
	_ = v9607
	var v9609 int32
	_ = v9609
	var v9615 int32
	_ = v9615
	var v9617 int32
	_ = v9617
	var v9619 int32
	_ = v9619
	var v9623 int32
	_ = v9623
	var v9624 int32
	_ = v9624
	var v9629 int32
	_ = v9629
	var v9630 int32
	_ = v9630
	var v9640 int32
	_ = v9640
	var v9644 int32
	_ = v9644
	var v9645 int32
	_ = v9645
	var v9657 int32
	_ = v9657
	var v9659 int32
	_ = v9659
	var v9662 int32
	_ = v9662
	var v9669 int32
	_ = v9669
	var v9672 int32
	_ = v9672
	var v9674 int32
	_ = v9674
	var v9676 int32
	_ = v9676
	var v9678 int32
	_ = v9678
	var v9681 int32
	_ = v9681
	var v9684 int32
	_ = v9684
	var v9688 int32
	_ = v9688
	var v9689 int32
	_ = v9689
	var v9690 int32
	_ = v9690
	var v9691 int32
	_ = v9691
	var v9692 int32
	_ = v9692
	var v9694 int32
	_ = v9694
	var v9697 int32
	_ = v9697
	var v9700 int32
	_ = v9700
	var v9702 int32
	_ = v9702
	var v9709 int32
	_ = v9709
	var v9710 int32
	_ = v9710
	var v9711 int32
	_ = v9711
	var v9716 int32
	_ = v9716
	var v9719 int32
	_ = v9719
	var v9724 int32
	_ = v9724
	var v9726 int32
	_ = v9726
	var v9730 int32
	_ = v9730
	var v9734 int64
	_ = v9734
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
	var v9743 int32
	_ = v9743
	var v9747 int32
	_ = v9747
	var v9750 int32
	_ = v9750
	var v9752 int32
	_ = v9752
	var v9754 int32
	_ = v9754
	var v9755 int32
	_ = v9755
	var v9756 int32
	_ = v9756
	var v9758 int32
	_ = v9758
	var v9761 int32
	_ = v9761
	var v9763 int32
	_ = v9763
	var v9764 int32
	_ = v9764
	var v9771 int32
	_ = v9771
	var v9773 int32
	_ = v9773
	var v9776 int32
	_ = v9776
	var v9780 int32
	_ = v9780
	var v9784 int32
	_ = v9784
	var v9786 int32
	_ = v9786
	var v9787 int32
	_ = v9787
	var v9788 int32
	_ = v9788
	var v9790 int32
	_ = v9790
	var v9794 int32
	_ = v9794
	var v9798 int32
	_ = v9798
	var v9802 int32
	_ = v9802
	var v9812 int32
	_ = v9812
	var v9843 int32
	_ = v9843
	var v9847 int32
	_ = v9847
	var v9851 int32
	_ = v9851
	var v9852 int32
	_ = v9852
	var v9853 int32
	_ = v9853
	var v9855 int32
	_ = v9855
	var v9859 int32
	_ = v9859
	var v9872 int32
	_ = v9872
	var v9873 int32
	_ = v9873
	var v9914 int32
	_ = v9914
	var v9915 int32
	_ = v9915
	var v9918 int32
	_ = v9918
	var v9919 int32
	_ = v9919
	var v9921 int32
	_ = v9921
	var v9924 int32
	_ = v9924
	var v9926 int32
	_ = v9926
	var v9929 int32
	_ = v9929
	var v9931 int32
	_ = v9931
	var v9932 int32
	_ = v9932
	var v9936 int32
	_ = v9936
	var v9937 int32
	_ = v9937
	var v9944 int32
	_ = v9944
	var v9946 int32
	_ = v9946
	var v9949 int32
	_ = v9949
	var v9951 int32
	_ = v9951
	var v9952 int32
	_ = v9952
	var v9953 int32
	_ = v9953
	var v9955 int32
	_ = v9955
	var v9958 int32
	_ = v9958
	var v9962 int32
	_ = v9962
	var v9965 int32
	_ = v9965
	var v9971 int32
	_ = v9971
	var v9976 int32
	_ = v9976
	var v9980 int32
	_ = v9980
	var v9982 int32
	_ = v9982
	var v9986 int32
	_ = v9986
	var v9987 int32
	_ = v9987
	var v9988 int32
	_ = v9988
	var v9989 int32
	_ = v9989
	var v9993 int32
	_ = v9993
	var v9996 int32
	_ = v9996
	var v9997 int32
	_ = v9997
	var v10005 int32
	_ = v10005
	var v10008 int32
	_ = v10008
	var v10010 int32
	_ = v10010
	var v10012 int32
	_ = v10012
	var v10014 int32
	_ = v10014
	var v10017 int32
	_ = v10017
	var v10023 int32
	_ = v10023
	var v10030 int32
	_ = v10030
	var v10033 int32
	_ = v10033
	var v10037 int32
	_ = v10037
	var v10040 int32
	_ = v10040
	var v10046 int32
	_ = v10046
	var v10051 int32
	_ = v10051
	var v10054 int32
	_ = v10054
	var v10098 int32
	_ = v10098
	var v10101 int32
	_ = v10101
	var v10105 int32
	_ = v10105
	var v10110 int32
	_ = v10110
	var v10114 int32
	_ = v10114
	var v10117 int32
	_ = v10117
	var v10121 int32
	_ = v10121
	var v10123 int32
	_ = v10123
	var v10128 int32
	_ = v10128
	var v10132 int32
	_ = v10132
	var v10135 int32
	_ = v10135
	var v10139 int32
	_ = v10139
	var v10144 int32
	_ = v10144
	var v10148 int32
	_ = v10148
	var v10151 int32
	_ = v10151
	var v10158 int32
	_ = v10158
	var v10163 int32
	_ = v10163
	var v10167 int32
	_ = v10167
	var v10170 int32
	_ = v10170
	var v10171 int32
	_ = v10171
	var v10179 int32
	_ = v10179
	var v10184 int32
	_ = v10184
	var v10189 int32
	_ = v10189
	var v10192 int32
	_ = v10192
	var v10196 int32
	_ = v10196
	var v10198 int32
	_ = v10198
	var v10203 int32
	_ = v10203
	var v10207 int32
	_ = v10207
	var v10210 int32
	_ = v10210
	var v10218 int32
	_ = v10218
	var v10223 int32
	_ = v10223
	var v10227 int32
	_ = v10227
	var v10230 int32
	_ = v10230
	var v10237 int32
	_ = v10237
	var v10242 int32
	_ = v10242
	var v10246 int32
	_ = v10246
	var v10249 int32
	_ = v10249
	var v10255 int32
	_ = v10255
	var v10260 int32
	_ = v10260
	var v10265 int32
	_ = v10265
	var v10268 int32
	_ = v10268
	var v10272 int32
	_ = v10272
	var v10274 int32
	_ = v10274
	var v10279 int32
	_ = v10279
	var v10283 int32
	_ = v10283
	var v10286 int32
	_ = v10286
	var v10290 int32
	_ = v10290
	var v10295 int32
	_ = v10295
	var v10299 int32
	_ = v10299
	var v10302 int32
	_ = v10302
	var v10308 int32
	_ = v10308
	var v10313 int32
	_ = v10313
	var v10317 int32
	_ = v10317
	var v10320 int32
	_ = v10320
	var v10324 int32
	_ = v10324
	var v10329 int32
	_ = v10329
	var v10333 int32
	_ = v10333
	var v10336 int32
	_ = v10336
	var v10340 int32
	_ = v10340
	var v10342 int32
	_ = v10342
	var v10347 int32
	_ = v10347
	var v10351 int32
	_ = v10351
	var v10354 int32
	_ = v10354
	var v10360 int32
	_ = v10360
	var v10365 int32
	_ = v10365
	var v10369 int32
	_ = v10369
	var v10372 int32
	_ = v10372
	var v10376 int32
	_ = v10376
	var v10378 int32
	_ = v10378
	var v10383 int32
	_ = v10383
	var v10394 int32
	_ = v10394
	var v10397 int32
	_ = v10397
	var v10401 int32
	_ = v10401
	var v10406 int32
	_ = v10406
	v1 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(528)
	m.G0 = v40
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v44
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v1)
	F_MemoryContextReset(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_initStringInfo(m, v40+int32(440))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[3]))
	if v56 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])))
	if v111 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[5]))
	if v60 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[6]))
	if v62 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v65 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_pairingheap_remove(m, int32(_a_F_PostgresMainLoopOnce_0), v56+int32(52))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[3])) = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[5]))
	if v75 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[6]))
	if v77 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+40))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[6]))
	v84 = v82 - int32(48)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v85))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v80)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v102 = v1
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[8])) = v102
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+40)) = v102
	goto L4
L14:
	;
	if v97 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L15:
	;
	v97 = base.B2i32(base.Ui32(v80) < base.Ui32(v85))
	goto L14
L16:
	;
	goto L17
L17:
	;
	v97 = int32(base.Ui32(v80-v85) >> (uint(int32(31)) % 32))
	goto L14
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v102 = v100
	goto L13
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+24))
	goto L23
L20:
	;
	goto L21
L21:
	;
	v309 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])) = uint8(v309)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v312 == int32(2) {
		goto L80
	} else {
		goto L81
	}
L22:
	;
	F_ReportChangedGUCOptions(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L59
	}
L23:
	;
	if (v116-int32(7))&int32(-9) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v124 = int32(0)
	F_pgstat_report_activity(m, int32(6), v124)
	mBase = m.M
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[12]))
	if v127 <= v124 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+24))
	goto L33
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[13]))
	if v131 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v134 = base.B2i32(v131 <= v127)
	goto L30
L29:
	;
	v134 = int32(0)
	goto L30
L30:
	;
	if v134 != 0 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v136)
	F_enable_timeout_after(m, int32(7), v127)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	if v143 != int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v147 = int32(0)
	F_pgstat_report_activity(m, int32(4), v147)
	mBase = m.M
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[12]))
	if v150 <= v147 {
		goto L22
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[15]))
	if v165 != 0 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[13]))
	if v154 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v157 = base.B2i32(v154 <= v150)
	goto L40
L39:
	;
	v157 = int32(0)
	goto L40
L40:
	;
	if v157 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v159)
	F_enable_timeout_after(m, int32(7), v150)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L22
L43:
	;
	F_ProcessNotifyInterrupt(m, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v170 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[16])))
	goto L48
L48:
	;
	if int32(0) < v170 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v187 = int32(0)
	F_pgstat_report_activity(m, int32(2), v187)
	mBase = m.M
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[17]))
	if v190 <= v187 {
		goto L22
	} else {
		goto L57
	}
L50:
	;
	if v175 != 0 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v175 == int32(0) {
		goto L49
	} else {
		goto L55
	}
L53:
	;
	F_enable_timeout_after(m, int32(10), v170)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	F_disable_timeout(m, int32(10))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])) = uint8(v194)
	F_enable_timeout_after(m, int32(9), v190)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L22
L59:
	;
	v204 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[19]))
	if v204 != int64(-9223372036854775807-1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	F_ReadyForQuery(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L78
	}
L61:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[20])))
	if v208&int32(8) == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[21]))
	switch v214 - int32(1) {
	case 0, 5:
		goto L63
	default:
		goto L60
	}
L63:
	;
	v221 = m.G0
	v222 = int32(16)
	v223 = v221 - v222
	m.G0 = v223
	F_gettimeofday(m, v223)
	mBase = m.M
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v223)))
	v227 = int64(*(*int32)(unsafe.Add(mBase, uint32(v223)+8)))
	m.G0 = v223 + v222
	v235 = v227 + v226*int64(1000000) - int64(946684800000000)
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[19])) = v235
	v238 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[22]))
	v240 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[23]))
	v242 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[24]))
	v244 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[25]))
	v246 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[26]))
	v249 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v249 == int32(0) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	if v240 < v238 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v256 = v238 - v240
	goto L69
L68:
	;
	v256 = int64(0)
	goto L69
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v40)+432)) = base.F64_div(base.F64_convert_i64_u(v256), float64(1000))
	if v244 < v242 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v264 = v242 - v244
	goto L72
L71:
	;
	v264 = int64(0)
	goto L72
L72:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v40)+424)) = base.F64_div(base.F64_convert_i64_u(v264), float64(1000))
	if v246 < v235 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v272 = v235 - v246
	goto L75
L74:
	;
	v272 = int64(0)
	goto L75
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v40)+416)) = base.F64_div(base.F64_convert_i64_u(v272), float64(1000))
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_1), v40+int32(416))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_3), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L60
L78:
	;
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v298)
	goto L21
L79:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])))
	if v624 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L80:
	;
	v315 = int32(_a_F_PostgresMainLoopOnce_5)
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v317 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_pg_printf(m, int32(_a_F_PostgresMainLoopOnce_6), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L117
	}
L83:
	;
	v325 = F_pq_getbyte(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L94
	}
L84:
	;
	v417 = F_pq_getmessage(m, v40+int32(440), v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L113
	}
L85:
	;
	v416 = int32(1073741822)
	goto L84
L86:
	;
	v413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v413)
	goto L85
L87:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L109
	}
L88:
	;
	v393 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v393)
	v416 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L84
L89:
	;
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[28])) = uint8(v390)
	goto L88
L90:
	;
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v387)
	goto L85
L91:
	;
	v380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v380)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[28])) = uint8(v380)
	v416 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L84
L92:
	;
	v376 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v376)
	v416 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L84
L93:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	goto L95
L94:
	;
	switch v325 + int32(1) {
	case 0:
		goto L93
	default:
		goto L87
	case 67, 81:
		goto L86
	case 68, 69, 70, 73:
		goto L92
	case 71, 82, 101:
		goto L90
	case 84:
		goto L91
	case 89:
		goto L89
	case 100, 103:
		goto L88
	}
L95:
	;
	if v331 == int32(2) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v334 = int32(-1)
	v337 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v354 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11])) = v354
	v356 = int32(-1)
	v359 = F_errstart(m, int32(14), v354)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L104
	}
L99:
	;
	if v337 == int32(0) {
		v592 = v334
		goto L79
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(476), int32(_a_F_PostgresMainLoopOnce_9))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v592 = v334
	goto L79
L104:
	;
	if v359 == int32(0) {
		v592 = v356
		goto L79
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_10), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(487), int32(_a_F_PostgresMainLoopOnce_9))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v592 = v356
	goto L79
L109:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v325
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_11), v40)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(562), int32(_a_F_PostgresMainLoopOnce_9))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	if v417 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v592 = int32(-1)
	goto L79
L115:
	;
	goto L116
L116:
	;
	v420 = int32(_a_F_PostgresMainLoopOnce_5)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v422 - int32(1)
	v592 = v325
	goto L79
L117:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[29]))
	v432 = F_fflush(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v435 = v40 + int32(440)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v437 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v436))) = uint8(v437)
	*(*int32)(unsafe.Add(mBase, uint32(v435)+12)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v437
	goto L119
L119:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[30]))
	goto L122
L120:
	;
	F_appendStringInfoChar(m, v40+int32(440), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L158
	}
L121:
	;
	F_appendStringInfoChar(m, v40+int32(440), int32(10))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L157
	}
L122:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v483 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v40)+444))
	if v559 != 0 {
		goto L120
	} else {
		goto L156
	}
L124:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v486 = F_do_getc(m, v444)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32]))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])))
	if v491 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32])) = v489
	switch v486 + int32(1) {
	case 0:
		goto L145
	default:
		goto L147
	case 11:
		goto L148
	}
L130:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v493 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[33]))
	if v508 == int32(0) {
		goto L129
	} else {
		goto L143
	}
L133:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[34]))
	if v497 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[15]))
	if v501 == int32(0) {
		goto L129
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	F_ProcessNotifyInterrupt(m, int32(1))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L129
L143:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[35]))
	F_SetLatch(m, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	goto L129
L145:
	;
	goto L123
L146:
	;
	if v519 <= int32(0) {
		goto L121
	} else {
		goto L154
	}
L147:
	;
	F_appendStringInfoChar(m, v40+int32(440), base.I32_extend8_s(v486))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L153
	}
L148:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v40)+444))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[36])))
	if v521 == int32(0) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	if v519 < int32(2) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	v527 = v526 + v519
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527-int32(1)))))
	if v530 != int32(10) {
		goto L147
	} else {
		goto L151
	}
L151:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527-int32(2)))))
	if v535 == int32(59) {
		goto L120
	} else {
		goto L152
	}
L152:
	;
	goto L147
L153:
	;
	goto L122
L154:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v519-int32(1)))))
	if v550 != int32(92) {
		goto L121
	} else {
		goto L155
	}
L155:
	;
	v554 = v519 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+444)) = v554
	v557 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v554+v546))) = uint8(v557)
	goto L122
L156:
	;
	v592 = int32(-1)
	goto L79
L157:
	;
	goto L120
L158:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[37])))
	if v575 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+400)) = v576
	F_pg_printf(m, int32(_a_F_PostgresMainLoopOnce_12), v40+int32(400))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v583 = F_fflush(m, v431)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v592 = int32(81)
	goto L79
L164:
	;
	F_disable_timeout(m, int32(7))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])))
	if v634 == int32(1) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v631 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v631)
	goto L166
L168:
	;
	F_disable_timeout(m, int32(9))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v644 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v641 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])) = uint8(v641)
	goto L170
L172:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v648 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])) = uint8(v648)
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[38]))
	if v651 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L174
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[38])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v592 != int32(-1) {
		goto L199
	} else {
		goto L200
	}
L179:
	;
	goto L178
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10394 = m.ExcPending
	if v10394 != 0 {
		goto L1
	} else {
		goto L2478
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[39])) = int32(99)
	goto L2477
L182:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10369 = m.ExcPending
	if v10369 != 0 {
		goto L1
	} else {
		goto L2472
	}
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10351 = m.ExcPending
	if v10351 != 0 {
		goto L1
	} else {
		goto L2468
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10333 = m.ExcPending
	if v10333 != 0 {
		goto L1
	} else {
		goto L2463
	}
L185:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10317 = m.ExcPending
	if v10317 != 0 {
		goto L1
	} else {
		goto L2459
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10299 = m.ExcPending
	if v10299 != 0 {
		goto L1
	} else {
		goto L2455
	}
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10283 = m.ExcPending
	if v10283 != 0 {
		goto L1
	} else {
		goto L2451
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10265 = m.ExcPending
	if v10265 != 0 {
		goto L1
	} else {
		goto L2446
	}
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10246 = m.ExcPending
	if v10246 != 0 {
		goto L1
	} else {
		goto L2442
	}
L190:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10227 = m.ExcPending
	if v10227 != 0 {
		goto L1
	} else {
		goto L2438
	}
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10207 = m.ExcPending
	if v10207 != 0 {
		goto L1
	} else {
		goto L2434
	}
L192:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10189 = m.ExcPending
	if v10189 != 0 {
		goto L1
	} else {
		goto L2429
	}
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10167 = m.ExcPending
	if v10167 != 0 {
		goto L1
	} else {
		goto L2425
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10148 = m.ExcPending
	if v10148 != 0 {
		goto L1
	} else {
		goto L2421
	}
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10132 = m.ExcPending
	if v10132 != 0 {
		goto L1
	} else {
		goto L2417
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10114 = m.ExcPending
	if v10114 != 0 {
		goto L1
	} else {
		goto L2412
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10098 = m.ExcPending
	if v10098 != 0 {
		goto L1
	} else {
		goto L2408
	}
L198:
	;
	m.G0 = v40 + int32(528)
	return
L199:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[28])))
	if v661&int32(1) != 0 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	switch v592 + int32(1) {
	case 0:
		goto L206
	default:
		goto L204
	case 67:
		goto L213
	case 68:
		goto L210
	case 69:
		goto L209
	case 70:
		goto L212
	case 71:
		goto L211
	case 73:
		goto L208
	case 81:
		goto L214
	case 82:
		goto L215
	case 84:
		goto L207
	case 89:
		goto L205
	case 100, 101, 103:
		goto L198
	}
L202:
	;
	goto L201
L203:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v10054 = m.ExcPending
	if v10054 != 0 {
		goto L1
	} else {
		goto L2407
	}
L204:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v10037 = m.ExcPending
	if v10037 != 0 {
		goto L1
	} else {
		goto L2403
	}
L205:
	;
	v10023 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10023 == int32(2) {
		goto L2398
	} else {
		goto L2399
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[40])) = int32(2)
	goto L205
L207:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v9993 = m.ExcPending
	if v9993 != 0 {
		goto L1
	} else {
		goto L2384
	}
L208:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v9980 = m.ExcPending
	if v9980 != 0 {
		goto L1
	} else {
		goto L2381
	}
L209:
	;
	v9726 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v9726 == int32(1) {
		goto L180
	} else {
		goto L2325
	}
L210:
	;
	v9684 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v9684 == int32(1) {
		goto L180
	} else {
		goto L2307
	}
L211:
	;
	v8584 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v8584 == int32(1) {
		goto L187
	} else {
		goto L2057
	}
L212:
	;
	v7825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v7825 == int32(1) {
		goto L180
	} else {
		goto L1871
	}
L213:
	;
	v6762 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v6762 == int32(1) {
		goto L180
	} else {
		goto L1660
	}
L214:
	;
	v6359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v6359 == int32(1) {
		goto L180
	} else {
		goto L1529
	}
L215:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v667 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v674 = v40 + int32(440)
	v675 = F_pq_getmsgstring(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L220
	}
L217:
	;
	v671 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v671
	goto L219
L218:
	;
	goto L219
L219:
	;
	goto L216
L220:
	;
	F_pq_getmsgend(m, v674)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v680 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v6356 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v6356)
	goto L198
L223:
	;
	v683 = m.G0
	v685 = v683 - int32(_a_F_PostgresMainLoopOnce_13)
	m.G0 = v685
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v692 == int32(0) {
		v714 = v688
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v5159 = m.G0
	v5161 = v5159 - int32(112)
	m.G0 = v5161
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v675
	v5166 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v5168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	F_pgstat_report_activity(m, int32(3), v675)
	mBase = m.M
	if v5168 == int32(1) {
		goto L1303
	} else {
		goto L1304
	}
L226:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v715 != int32(4) {
		goto L261
	} else {
		goto L262
	}
L227:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v695 == int32(4) {
		v714 = v688
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v688)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v688)+76)) = int32(1)
	if v698 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	F_s_lock(m, v688+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v688)+4)) = int32(4)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v714 = v713
	goto L226
L232:
	;
	goto L231
L233:
	;
	m.G0 = v685 + int32(_a_F_PostgresMainLoopOnce_13)
	if v1141 != 0 {
		goto L222
	} else {
		goto L1302
	}
L234:
	;
	F_EndReplicationCommand(m, v5051)
	mBase = m.M
	v5071 = m.ExcPending
	if v5071 != 0 {
		goto L1
	} else {
		goto L1300
	}
L235:
	;
	v5022 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v5023 = m.ExcPending
	if v5023 != 0 {
		goto L1
	} else {
		goto L1296
	}
L236:
	;
	if v4383 == int32(-1) {
		goto L1285
	} else {
		goto L1286
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4971 = m.ExcPending
	if v4971 != 0 {
		goto L1
	} else {
		goto L1281
	}
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4952 = m.ExcPending
	if v4952 != 0 {
		goto L1
	} else {
		goto L1277
	}
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4929 = m.ExcPending
	if v4929 != 0 {
		goto L1
	} else {
		goto L1273
	}
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4910 = m.ExcPending
	if v4910 != 0 {
		goto L1
	} else {
		goto L1269
	}
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4891 = m.ExcPending
	if v4891 != 0 {
		goto L1
	} else {
		goto L1265
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4872 = m.ExcPending
	if v4872 != 0 {
		goto L1
	} else {
		goto L1261
	}
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4853 = m.ExcPending
	if v4853 != 0 {
		goto L1
	} else {
		goto L1257
	}
L244:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v4849 = m.ExcPending
	if v4849 != 0 {
		goto L1
	} else {
		goto L1256
	}
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L1
	} else {
		goto L1253
	}
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L1
	} else {
		goto L1249
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4802 = m.ExcPending
	if v4802 != 0 {
		goto L1
	} else {
		goto L1245
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L1
	} else {
		goto L1242
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L1
	} else {
		goto L1239
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		goto L1
	} else {
		goto L1236
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
		goto L1
	} else {
		goto L1233
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		goto L1
	} else {
		goto L1230
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4706 = m.ExcPending
	if v4706 != 0 {
		goto L1
	} else {
		goto L1227
	}
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4690 = m.ExcPending
	if v4690 != 0 {
		goto L1
	} else {
		goto L1223
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4674 = m.ExcPending
	if v4674 != 0 {
		goto L1
	} else {
		goto L1219
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		goto L1
	} else {
		goto L1215
	}
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		goto L1
	} else {
		goto L1211
	}
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4623 = m.ExcPending
	if v4623 != 0 {
		goto L1
	} else {
		goto L1207
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		goto L1
	} else {
		goto L1203
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4589 = m.ExcPending
	if v4589 != 0 {
		goto L1
	} else {
		goto L1199
	}
L261:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48])))
	if v719 != 0 {
		goto L266
	} else {
		goto L267
	}
L262:
	;
	goto L263
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4573 = m.ExcPending
	if v4573 != 0 {
		goto L1
	} else {
		goto L1195
	}
L264:
	;
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v748 != 0 {
		goto L275
	} else {
		goto L276
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L272
	}
L266:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+20))
	goto L269
L267:
	;
	goto L268
L268:
	;
	goto L264
L269:
	;
	if base.B2i32(v722 == int32(2)) == int32(0) {
		goto L265
	} else {
		goto L270
	}
L270:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49]))
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])) = v728
	goto L268
L272:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_16), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(609), int32(_a_F_PostgresMainLoopOnce_18))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	if v752 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L277
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v769
	v773 = v685 + int32(428)
	v775 = F_palloc0(m, int32(20))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L285
	}
L280:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[52]))
	v762 = F_AllocSetContextCreateInternal(m, v757, int32(_a_F_PostgresMainLoopOnce_19), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	F_MemoryContextReset(m, v752)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L284
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51])) = v762
	v769 = v762
	goto L279
L284:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	v769 = v768
	goto L279
L285:
	;
	if v773 != 0 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v800 = int32(0)
	base.MemoryFill(m, v779, v800, int32(96))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+60)) = v800
	v806 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v803)+52)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v803)+44)) = v800
	*(*int64)(unsafe.Add(mBase, uint32(v803)+36)) = v806
	*(*int64)(unsafe.Add(mBase, uint32(v803)+4)) = v806
	*(*int64)(unsafe.Add(mBase, uint32(v803)+12)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v803)+20)) = v800
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	*(*int32)(unsafe.Add(mBase, uint32(v818))) = v775
	v820 = F_strlen(m, v675)
	mBase = m.M
	v822 = v820 + int32(2)
	v823 = F_palloc(m, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L297
	}
L287:
	;
	v779 = F_palloc(m, int32(96))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	v785 = int32(28)
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32])) = v785
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L292
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v773))) = v779
	if v779 != 0 {
		goto L286
	} else {
		goto L291
	}
L291:
	;
	v785 = int32(48)
	goto L289
L292:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_22), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_23), int32(275), int32(_a_F_PostgresMainLoopOnce_24))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	v1123 = m.G0
	v1125 = v1123 - int32(16)
	m.G0 = v1125
	v1129 = F_replication_yylex(m, v1125+int32(8), v1122)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L331
	}
L296:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_25))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L329
	}
L297:
	;
	if v823 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if v820 <= int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	goto L300
L300:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_26))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L328
	}
L301:
	;
	v1020 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v820+v823))) = uint16(v1020)
	if base.Ui32(v822) < base.Ui32(int32(2)) {
		v1108 = v1020
		goto L315
	} else {
		goto L316
	}
L302:
	;
	v828 = v820 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v820) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v834 = v1
	v843 = v1
	goto L306
L304:
	;
	v909 = v1
	goto L305
L305:
	;
	v946 = v909
	v948 = v1
	goto L310
L306:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v843+v823))) = uint8(v872)
	v875 = v843 | int32(1)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v823+v875))) = uint8(v878)
	v881 = v843 | int32(2)
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v823+v881))) = uint8(v884)
	v887 = v843 | int32(3)
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v823+v887))) = uint8(v890)
	v892 = int32(4)
	v893 = v843 + v892
	v895 = v834 + v892
	if v895 != v820&int32(2147483644) {
		v834 = v895
		v843 = v893
		goto L306
	} else {
		goto L308
	}
L307:
	;
	if v828 == int32(0) {
		goto L301
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v909 = v893
	goto L305
L310:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946+v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v946+v823))) = uint8(v975)
	v977 = int32(1)
	v980 = v948 + v977
	if v980 != v828 {
		v946 = v946 + v977
		v948 = v980
		goto L310
	} else {
		goto L312
	}
L311:
	;
	goto L301
L312:
	;
	goto L311
L313:
	;
	if v1108 == int32(0) {
		goto L296
	} else {
		goto L327
	}
L314:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_27))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L326
	}
L315:
	;
	goto L313
L316:
	;
	v1026 = v822 - int32(2)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823+v1026))))
	if v1028 != 0 {
		v1108 = v1020
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822+v823-int32(1)))))
	if v1032 != 0 {
		v1108 = v1020
		goto L315
	} else {
		goto L318
	}
L318:
	;
	v1034 = F_palloc(m, int32(48))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	if v1034 == int32(0) {
		goto L314
	} else {
		goto L320
	}
L320:
	;
	v1038 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+20)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+8)) = v823
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+4)) = v823
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+12)) = v1026
	*(*int64)(unsafe.Add(mBase, uint32(v1034)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1034)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+16)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v1034))) = v1038
	F_replication_yyensure_buffer_stack(m, v818)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1052+v1053<<(uint(int32(2))%32))))
	if v1057 == v1034 {
		v1108 = v1034
		goto L315
	} else {
		goto L322
	}
L322:
	;
	if v1057 != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v818)+36))
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1059))) = uint8(v1060)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v1064 = int32(2)
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1062+v1063<<(uint(v1064)%32))))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v818)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+8)) = v1068
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1070+v1071<<(uint(v1064)%32))))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v818)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+16)) = v1076
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v1080 = v1078
	v1081 = v1079
	goto L325
L324:
	;
	v1080 = v1052
	v1081 = v1053
	goto L325
L325:
	;
	v1082 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1081<<(uint(v1082)%32)+v1080))) = v1034
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v1090 = v1086 + v1087<<(uint(v1082)%32)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+28)) = v1092
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+36)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v818)+80)) = v1095
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+4)) = v1099
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1095))))
	*(*uint8)(unsafe.Add(mBase, uint32(v818)+24)) = uint8(v1101)
	*(*int32)(unsafe.Add(mBase, uint32(v818)+48)) = int32(1)
	v1108 = v1034
	goto L315
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+20)) = int32(1)
	goto L295
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	m.G0 = v1125 + int32(16)
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	if v1141 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L331:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v1129-int32(262)) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	if v1129 != int32(282) {
		v1141 = int32(0)
		goto L330
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	*(*int32)(unsafe.Add(mBase, uint32(v1138))) = v1129
	v1141 = int32(1)
	goto L330
L335:
	;
	goto L334
L336:
	;
	F_replication_scanner_finish(m, v1145)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v1176 = m.G0
	v1178 = v1176 - int32(1872)
	m.G0 = v1178
	*(*int64)(unsafe.Add(mBase, uint32(v1178)+1864)) = int64(0)
	v1184 = v1178 - int32(-64)
	v1186 = v1178 + int32(1664)
	v1188 = v1186
	v1193 = v1
	v1197 = v1184
	v1202 = v1184
	v1206 = v1186
	v1211 = int32(200)
	v1212 = int32(-2)
	goto L352
L339:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v690
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	F_MemoryContextReset(m, v1153)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	if v1157 != 0 {
		goto L233
	} else {
		goto L341
	}
L341:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_28), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2064), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L1
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
	if v1819 != 0 {
		goto L260
	} else {
		goto L519
	}
L347:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L1
	} else {
		goto L515
	}
L348:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L1
	} else {
		goto L511
	}
L349:
	;
	if v1178+int32(1664) != v1810 {
		goto L507
	} else {
		goto L508
	}
L350:
	;
	v1810 = v1242
	v1819 = int32(1)
	goto L349
L351:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L1
	} else {
		goto L506
	}
L352:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1206))) = uint8(v1193)
	if base.Ui32(v1188+v1211-int32(1)) <= base.Ui32(v1206) {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_31))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L505
	}
L354:
	;
	if int32(_a_F_PostgresMainLoopOnce_32) < v1211 {
		goto L351
	} else {
		goto L357
	}
L355:
	;
	v1272 = v1188
	v1275 = v1197
	v1277 = v1202
	v1279 = v1206
	v1280 = v1211
	goto L356
L356:
	;
	if v1193 == int32(34) {
		goto L374
	} else {
		goto L375
	}
L357:
	;
	v1232 = int32(_a_F_PostgresMainLoopOnce_7)
	v1234 = v1211 << (uint(int32(1)) % 32)
	if v1232 <= v1234 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1237 = v1232
	goto L360
L359:
	;
	v1237 = v1234
	goto L360
L360:
	;
	v1242 = F_palloc(m, v1237*int32(9)+int32(7))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	if v1242 == int32(0) {
		goto L351
	} else {
		goto L362
	}
L362:
	;
	v1246 = v1206 - v1188
	v1248 = v1246 + int32(1)
	if v1248 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	base.MemoryCopy(m, v1242, v1188, v1248)
	goto L365
L364:
	;
	goto L365
L365:
	;
	v1253 = base.I32_div_s(v1237+int32(7), int32(8))
	v1254 = int32(3)
	v1256 = v1242 + v1253<<(uint(v1254)%32)
	v1258 = v1248 << (uint(v1254) % 32)
	if v1258 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	base.MemoryCopy(m, v1256, v1197, v1258)
	goto L368
L367:
	;
	goto L368
L368:
	;
	if v1178+int32(1664) != v1188 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	F_pfree(m, v1188)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	if v1237-int32(1) <= v1246 {
		goto L350
	} else {
		goto L373
	}
L372:
	;
	goto L371
L373:
	;
	v1272 = v1242
	v1275 = v1256
	v1277 = v1256 + v1258 - int32(8)
	v1279 = v1242 + v1246
	v1280 = v1237
	goto L356
L374:
	;
	v1810 = v1272
	v1819 = int32(0)
	goto L349
L375:
	;
	goto L376
L376:
	;
	v1286 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1193)+uint32(_c_F_PostgresMainLoopOnce[54]))))
	if v1286 == int32(-36) {
		v1322 = v1212
		goto L379
	} else {
		goto L380
	}
L377:
	;
	goto L353
L378:
	;
	v1188 = v1272
	v1193 = base.I32_extend8_s(v1796)
	v1197 = v1275
	v1202 = v1791
	v1206 = v1793 + int32(1)
	v1211 = v1280
	v1212 = v1794
	goto L352
L379:
	;
	v1325 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1193)+uint32(_c_F_PostgresMainLoopOnce[55]))))
	v1327 = v1325 & int32(255)
	if v1327 == int32(0) {
		goto L377
	} else {
		goto L395
	}
L380:
	;
	if v1212 == int32(-2) {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v1309 = v1286 + v1308
	if base.Ui32(int32(80)) < base.Ui32(v1309) {
		v1322 = v1307
		goto L379
	} else {
		goto L393
	}
L382:
	;
	v1293 = F_replication_yylex(m, v1178+int32(1864), v1145)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	v1295 = v1212
	goto L384
L384:
	;
	if v1295 <= int32(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	v1295 = v1293
	goto L384
L386:
	;
	v1298 = int32(0)
	v1307 = v1298
	v1308 = v1298
	goto L381
L387:
	;
	goto L388
L388:
	;
	if v1295 == int32(256) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1810 = v1272
	v1819 = int32(1)
	goto L349
L390:
	;
	goto L391
L391:
	;
	if base.Ui32(int32(282)) < base.Ui32(v1295) {
		v1307 = v1295
		v1308 = int32(2)
		goto L381
	} else {
		goto L392
	}
L392:
	;
	v1306 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1295)+uint32(_c_F_PostgresMainLoopOnce[56]))))
	v1307 = v1295
	v1308 = v1306
	goto L381
L393:
	;
	v1312 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1309)+uint32(_c_F_PostgresMainLoopOnce[57]))))
	if v1308 != v1312 {
		v1322 = v1307
		goto L379
	} else {
		goto L394
	}
L394:
	;
	v1314 = *(*int64)(unsafe.Add(mBase, uint32(v1178)+1864))
	*(*int64)(unsafe.Add(mBase, uint32(v1277)+8)) = v1314
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+uint32(_c_F_PostgresMainLoopOnce[58]))))
	v1791 = v1277 + int32(8)
	v1793 = v1279
	v1794 = int32(-2)
	v1796 = v1319
	goto L378
L395:
	;
	v1333 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1325)+uint32(_c_F_PostgresMainLoopOnce[59]))))
	v1337 = v1277 + (int32(1)-v1333)<<(uint(int32(3))%32)
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1337)))
	v1340 = int32(base.Ui32(v1338) >> (uint(int32(8)) % 32))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+4))
	switch v1327 - int32(2) {
	case 0:
		goto L458
	default:
		v1751 = v1338
		v1752 = v1340
		goto L396
	case 14:
		goto L457
	case 15:
		goto L456
	case 16:
		goto L455
	case 17:
		goto L454
	case 18:
		goto L453
	case 19:
		goto L452
	case 20:
		goto L451
	case 21:
		goto L450
	case 22:
		goto L449
	case 23:
		goto L448
	case 24:
		goto L447
	case 25:
		goto L446
	case 26, 44, 46, 48, 53:
		goto L445
	case 27:
		goto L444
	case 28:
		goto L443
	case 29:
		goto L442
	case 30:
		goto L441
	case 31:
		goto L440
	case 32:
		goto L439
	case 33:
		goto L438
	case 34:
		goto L437
	case 35:
		goto L436
	case 36:
		goto L435
	case 37:
		goto L434
	case 38:
		goto L433
	case 41:
		goto L432
	case 42:
		goto L431
	case 43:
		goto L430
	case 45:
		goto L429
	case 47:
		goto L428
	case 49:
		goto L427
	case 50:
		goto L426
	case 51:
		goto L425
	case 52:
		goto L424
	case 54:
		goto L423
	case 55:
		goto L422
	case 56:
		goto L421
	case 57:
		goto L420
	case 58:
		goto L419
	case 59:
		goto L418
	case 60:
		goto L417
	case 61:
		goto L416
	case 62:
		goto L415
	case 63:
		goto L414
	case 64:
		goto L413
	case 65:
		goto L412
	case 66:
		goto L411
	case 67:
		goto L410
	case 68:
		goto L409
	case 69:
		goto L408
	case 70:
		goto L407
	case 71:
		goto L406
	case 72:
		goto L405
	case 73:
		goto L404
	case 74:
		goto L403
	case 75:
		goto L402
	case 76:
		goto L401
	case 77:
		goto L400
	case 78:
		goto L399
	case 79:
		goto L398
	case 80:
		goto L397
	}
L396:
	;
	v1755 = v1277 - v1333<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1755)+12)) = v1341
	v1759 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1755)+8)) = v1751&int32(255) | v1752<<(uint(v1759)%32)
	v1764 = v1755 + v1759
	v1765 = v1279 - v1333
	v1766 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1765))))
	v1769 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1325)+uint32(_c_F_PostgresMainLoopOnce[60]))))
	v1772 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1769)+uint32(_c_F_PostgresMainLoopOnce[61]))))
	v1773 = v1766 + v1772
	if base.Ui32(v1773) <= base.Ui32(int32(80)) {
		goto L501
	} else {
		goto L502
	}
L397:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_33)
	v1752 = int32(326)
	goto L396
L398:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_34)
	v1752 = int32(363)
	goto L396
L399:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_35)
	v1752 = int32(362)
	goto L396
L400:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_36)
	v1752 = int32(362)
	goto L396
L401:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_37)
	v1752 = int32(1482)
	goto L396
L402:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_38)
	v1752 = int32(69)
	goto L396
L403:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_39)
	v1752 = int32(1264)
	goto L396
L404:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_40)
	v1752 = int32(361)
	goto L396
L405:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_41)
	v1752 = int32(1287)
	goto L396
L406:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_42)
	v1752 = int32(1287)
	goto L396
L407:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_43)
	v1752 = int32(1530)
	goto L396
L408:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_44)
	v1752 = int32(431)
	goto L396
L409:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_45)
	v1752 = int32(50)
	goto L396
L410:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_46)
	v1752 = int32(356)
	goto L396
L411:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_47)
	v1752 = int32(356)
	goto L396
L412:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_48)
	v1752 = int32(357)
	goto L396
L413:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_49)
	v1752 = int32(1091)
	goto L396
L414:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_50)
	v1752 = int32(130)
	goto L396
L415:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_51)
	v1752 = int32(1186)
	goto L396
L416:
	;
	v1751 = int32(_a_F_PostgresMainLoopOnce_52)
	v1752 = int32(959)
	goto L396
L417:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1751 = v1707
	v1752 = int32(base.Ui32(v1707) >> (uint(int32(8)) % 32))
	goto L396
L418:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1700 = F_makeInteger(m, v1699)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L498
	}
L419:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1689 = F_makeString(m, v1688)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L496
	}
L420:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1678 = F_makeString(m, v1677)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L494
	}
L421:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1670 = F_makeDefElem(m, v1667, int32(0), int32(-1))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L493
	}
L422:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+52)) = v1657
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+56)) = v1657
	v1663 = F_list_make1_impl(m, int32(1), v1178+int32(52))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L492
	}
L423:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(16))))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1653 = F_lappend(m, v1651, v1652)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L1
	} else {
		goto L491
	}
L424:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1645 = F_makeString(m, v1644)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L490
	}
L425:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1640 = F_makeDefElem(m, v1637, v1638, int32(-1))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L489
	}
L426:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(16))))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1631 = F_lappend(m, v1629, v1630)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L1
	} else {
		goto L488
	}
L427:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+48)) = v1617
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+60)) = v1617
	v1623 = F_list_make1_impl(m, int32(1), v1178+int32(48))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L487
	}
L428:
	;
	v1612 = int32(8)
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1612)))
	v1751 = v1614
	v1752 = int32(base.Ui32(v1614) >> (uint(v1612) % 32))
	goto L396
L429:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	if v1607 == int32(0) {
		goto L347
	} else {
		goto L486
	}
L430:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1751 = v1604
	v1752 = int32(base.Ui32(v1604) >> (uint(int32(8)) % 32))
	goto L396
L431:
	;
	v1751 = int32(0)
	v1752 = v1340
	goto L396
L432:
	;
	v1751 = int32(1)
	v1752 = v1340
	goto L396
L433:
	;
	v1596 = F_palloc0(m, int32(4))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L1
	} else {
		goto L485
	}
L434:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	if v1583 == int32(0) {
		goto L348
	} else {
		goto L483
	}
L435:
	;
	v1567 = F_palloc0(m, int32(32))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L1
	} else {
		goto L482
	}
L436:
	;
	v1550 = F_palloc0(m, int32(32))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L481
	}
L437:
	;
	v1535 = F_palloc0(m, int32(12))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L480
	}
L438:
	;
	v1522 = F_palloc0(m, int32(12))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L1
	} else {
		goto L479
	}
L439:
	;
	v1511 = F_palloc0(m, int32(12))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L478
	}
L440:
	;
	v1503 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L476
	}
L441:
	;
	v1494 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L1
	} else {
		goto L474
	}
L442:
	;
	v1485 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_53))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L472
	}
L443:
	;
	v1476 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_54))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L470
	}
L444:
	;
	v1467 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_55))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L468
	}
L445:
	;
	v1463 = int32(0)
	v1751 = v1463
	v1752 = v1463
	goto L396
L446:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1459 = F_lappend(m, v1457, v1458)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L467
	}
L447:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1751 = v1452
	v1752 = int32(base.Ui32(v1452) >> (uint(int32(8)) % 32))
	goto L396
L448:
	;
	v1447 = int32(8)
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1447)))
	v1751 = v1449
	v1752 = int32(base.Ui32(v1449) >> (uint(v1447) % 32))
	goto L396
L449:
	;
	v1425 = F_palloc0(m, int32(24))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L1
	} else {
		goto L466
	}
L450:
	;
	v1406 = F_palloc0(m, int32(24))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L465
	}
L451:
	;
	v1399 = F_palloc0(m, int32(8))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L464
	}
L452:
	;
	v1388 = F_palloc0(m, int32(8))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L463
	}
L453:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(16))))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+4)) = v1379
	*(*int32)(unsafe.Add(mBase, uint32(v1178))) = v1378
	v1383 = F_psprintf(m, int32(_a_F_PostgresMainLoopOnce_56), v1178)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L462
	}
L454:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1751 = v1373
	v1752 = int32(base.Ui32(v1373) >> (uint(int32(8)) % 32))
	goto L396
L455:
	;
	v1365 = F_palloc0(m, int32(8))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L461
	}
L456:
	;
	v1356 = F_palloc0(m, int32(8))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L460
	}
L457:
	;
	v1349 = F_palloc0(m, int32(4))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L459
	}
L458:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v685+int32(424)))) = v1346
	v1751 = v1338
	v1752 = v1340
	goto L396
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349))) = int32(448)
	v1751 = v1349
	v1752 = int32(base.Ui32(v1349) >> (uint(int32(8)) % 32))
	goto L396
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1356))) = int32(454)
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1356)+4)) = v1360
	v1751 = v1356
	v1752 = int32(base.Ui32(v1356) >> (uint(int32(8)) % 32))
	goto L396
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1365))) = int32(159)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1365)+4)) = v1369
	v1751 = v1365
	v1752 = int32(base.Ui32(v1365) >> (uint(int32(8)) % 32))
	goto L396
L462:
	;
	v1751 = v1383
	v1752 = int32(base.Ui32(v1383) >> (uint(int32(8)) % 32))
	goto L396
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1388))) = int32(449)
	v1392 = int32(8)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1392)))
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+4)) = v1394
	v1751 = v1388
	v1752 = int32(base.Ui32(v1388) >> (uint(v1392) % 32))
	goto L396
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = int32(449)
	v1751 = v1399
	v1752 = int32(base.Ui32(v1399) >> (uint(int32(8)) % 32))
	goto L396
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1406)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1406))) = int32(450)
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1406)+4)) = v1414
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1406)+16)) = uint8(v1418)
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1406)+20)) = v1420
	v1751 = v1406
	v1752 = int32(base.Ui32(v1406) >> (uint(int32(8)) % 32))
	goto L396
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1425))) = int32(450)
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+4)) = v1433
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277-int32(24)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1425)+16)) = uint8(v1437)
	v1439 = int32(8)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1439)))
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+12)) = v1441
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1425)+20)) = v1443
	v1751 = v1425
	v1752 = int32(base.Ui32(v1425) >> (uint(v1439) % 32))
	goto L396
L467:
	;
	v1751 = v1459
	v1752 = int32(base.Ui32(v1459) >> (uint(int32(8)) % 32))
	goto L396
L468:
	;
	v1470 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1467, int32(-1))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1751 = v1470
	v1752 = int32(base.Ui32(v1470) >> (uint(int32(8)) % 32))
	goto L396
L470:
	;
	v1479 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1476, int32(-1))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v1751 = v1479
	v1752 = int32(base.Ui32(v1479) >> (uint(int32(8)) % 32))
	goto L396
L472:
	;
	v1488 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1485, int32(-1))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	v1751 = v1488
	v1752 = int32(base.Ui32(v1488) >> (uint(int32(8)) % 32))
	goto L396
L474:
	;
	v1497 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_39), v1494, int32(-1))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	v1751 = v1497
	v1752 = int32(base.Ui32(v1497) >> (uint(int32(8)) % 32))
	goto L396
L476:
	;
	v1506 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_37), v1503, int32(-1))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	v1751 = v1506
	v1752 = int32(base.Ui32(v1506) >> (uint(int32(8)) % 32))
	goto L396
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1511))) = int32(451)
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1516 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1511)+8)) = uint8(v1516)
	*(*int32)(unsafe.Add(mBase, uint32(v1511)+4)) = v1515
	v1751 = v1511
	v1752 = int32(base.Ui32(v1511) >> (uint(int32(8)) % 32))
	goto L396
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1522))) = int32(451)
	v1526 = int32(8)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1526)))
	v1529 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1522)+8)) = uint8(v1529)
	*(*int32)(unsafe.Add(mBase, uint32(v1522)+4)) = v1528
	v1751 = v1522
	v1752 = int32(base.Ui32(v1522) >> (uint(v1526) % 32))
	goto L396
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1535))) = int32(452)
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1535)+4)) = v1541
	v1543 = int32(8)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1277-v1543)))
	*(*int32)(unsafe.Add(mBase, uint32(v1535)+8)) = v1545
	v1751 = v1535
	v1752 = int32(base.Ui32(v1535) >> (uint(v1543) % 32))
	goto L396
L481:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1550))) = int64(453)
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+8)) = v1556
	v1558 = int32(8)
	v1560 = *(*int64)(unsafe.Add(mBase, uint32(v1277-v1558)))
	*(*int64)(unsafe.Add(mBase, uint32(v1550)+16)) = v1560
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1550)+12)) = v1562
	v1751 = v1550
	v1752 = int32(base.Ui32(v1550) >> (uint(v1558) % 32))
	goto L396
L482:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1567))) = int64(4294967749)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1277-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+8)) = v1573
	v1575 = int32(8)
	v1577 = *(*int64)(unsafe.Add(mBase, uint32(v1277-v1575)))
	*(*int64)(unsafe.Add(mBase, uint32(v1567)+16)) = v1577
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+24)) = v1579
	v1751 = v1567
	v1752 = int32(base.Ui32(v1567) >> (uint(v1575) % 32))
	goto L396
L483:
	;
	v1587 = F_palloc0(m, int32(8))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1587))) = int32(455)
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1587)+4)) = v1591
	v1751 = v1587
	v1752 = int32(base.Ui32(v1587) >> (uint(int32(8)) % 32))
	goto L396
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1596))) = int32(456)
	v1751 = v1596
	v1752 = int32(base.Ui32(v1596) >> (uint(int32(8)) % 32))
	goto L396
L486:
	;
	v1751 = v1607
	v1752 = int32(base.Ui32(v1607) >> (uint(int32(8)) % 32))
	goto L396
L487:
	;
	v1751 = v1623
	v1752 = int32(base.Ui32(v1623) >> (uint(int32(8)) % 32))
	goto L396
L488:
	;
	v1751 = v1631
	v1752 = int32(base.Ui32(v1631) >> (uint(int32(8)) % 32))
	goto L396
L489:
	;
	v1751 = v1640
	v1752 = int32(base.Ui32(v1640) >> (uint(int32(8)) % 32))
	goto L396
L490:
	;
	v1751 = v1645
	v1752 = int32(base.Ui32(v1645) >> (uint(int32(8)) % 32))
	goto L396
L491:
	;
	v1751 = v1653
	v1752 = int32(base.Ui32(v1653) >> (uint(int32(8)) % 32))
	goto L396
L492:
	;
	v1751 = v1663
	v1752 = int32(base.Ui32(v1663) >> (uint(int32(8)) % 32))
	goto L396
L493:
	;
	v1751 = v1670
	v1752 = int32(base.Ui32(v1670) >> (uint(int32(8)) % 32))
	goto L396
L494:
	;
	v1681 = F_makeDefElem(m, v1676, v1678, int32(-1))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	v1751 = v1681
	v1752 = int32(base.Ui32(v1681) >> (uint(int32(8)) % 32))
	goto L396
L496:
	;
	v1692 = F_makeDefElem(m, v1687, v1689, int32(-1))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1751 = v1692
	v1752 = int32(base.Ui32(v1692) >> (uint(int32(8)) % 32))
	goto L396
L498:
	;
	v1703 = F_makeDefElem(m, v1698, v1700, int32(-1))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1751 = v1703
	v1752 = int32(base.Ui32(v1703) >> (uint(int32(8)) % 32))
	goto L396
L500:
	;
	v1785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773)+uint32(_c_F_PostgresMainLoopOnce[58]))))
	v1791 = v1764
	v1793 = v1765
	v1794 = v1322
	v1796 = v1785
	goto L378
L501:
	;
	v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773)+uint32(_c_F_PostgresMainLoopOnce[57]))))
	if v1776 == v1766&int32(255) {
		goto L500
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v1782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1769)+uint32(_c_F_PostgresMainLoopOnce[62]))))
	v1791 = v1764
	v1793 = v1765
	v1794 = v1322
	v1796 = v1782
	goto L378
L504:
	;
	goto L503
L505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	F_pfree(m, v1810)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L1
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	m.G0 = v1178 + int32(1872)
	goto L346
L510:
	;
	goto L509
L511:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+16)) = v1835
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_58), v1178+int32(16))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_59), int32(322), int32(_a_F_PostgresMainLoopOnce_60))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L515:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+32)) = v1854
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_58), v1178+int32(32))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_59), int32(363), int32(_a_F_PostgresMainLoopOnce_60))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	F_replication_scanner_finish(m, v1866)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v675
	F_pgstat_report_activity(m, int32(3), v675)
	mBase = m.M
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[63])))
	if v1876 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v1877 = int32(15)
	goto L523
L522:
	;
	v1877 = int32(14)
	goto L523
L523:
	;
	v1879 = F_errstart(m, v1877, int32(0))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	if v1879 != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+400)) = v675
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_61), v685+int32(400))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L1
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+24))
	goto L530
L528:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2095), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	goto L527
L530:
	;
	if (v1894-int32(7))&int32(-9) == int32(0) {
		goto L259
	} else {
		goto L531
	}
L531:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v1902 != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_62))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L536
	}
L535:
	;
	goto L534
L536:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_63))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_64))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L538
	}
L538:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1914)))
	switch v1915 - int32(448) {
	case 0:
		goto L548
	case 1:
		goto L539
	case 2:
		goto L546
	case 3:
		goto L545
	case 4:
		goto L544
	case 5:
		goto L543
	case 6:
		goto L547
	case 7:
		goto L542
	case 8:
		goto L541
	default:
		goto L540
	}
L539:
	;
	v4560 = int32(_a_F_PostgresMainLoopOnce_65)
	F_PreventInTransactionBlock(m, int32(1), v4560)
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L1
	} else {
		goto L1193
	}
L540:
	;
	if v1915 == int32(159) {
		goto L235
	} else {
		goto L1189
	}
L541:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_66))
	mBase = m.M
	v4130 = m.ExcPending
	if v4130 != 0 {
		goto L1
	} else {
		goto L1100
	}
L542:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_67))
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
		goto L1
	} else {
		goto L1071
	}
L543:
	;
	v3282 = int32(_a_F_PostgresMainLoopOnce_68)
	F_PreventInTransactionBlock(m, int32(1), v3282)
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L1
	} else {
		goto L912
	}
L544:
	;
	v2890 = int32(0)
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+8))
	if v2892 == v2890 {
		v3054 = v1
		v3055 = v1
		v3059 = v1
		v3066 = v2890
		goto L814
	} else {
		goto L815
	}
L545:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+4))
	v2882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+8)))
	F_ReplicationSlotDrop(m, v2881, (v2882^int32(-1))&int32(1))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		goto L1
	} else {
		goto L813
	}
L546:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2232 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = v2232
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+20))
	if v2234 == v2232 {
		v2568 = v1
		v2571 = v1
		v2573 = v1
		v2574 = v1
		goto L637
	} else {
		goto L638
	}
L547:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[65]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = int64(0)
	v2080 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L1
	} else {
		goto L596
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = int32(0)
	v1921 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[66]))
	v1922 = *(*int64)(unsafe.Add(mBase, uint32(v1921)))
	goto L549
L549:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v685)+32)) = v1922
	v1926 = int32(32)
	v1930 = F_pg_snprintf(m, v685+int32(_a_F_PostgresMainLoopOnce_69), v1926, int32(_a_F_PostgresMainLoopOnce_70), v685+v1926)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v1935 == int32(1) {
		goto L552
	} else {
		goto L553
	}
L551:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])) = uint8(v1945)
	if v1945 != 0 {
		goto L556
	} else {
		goto L557
	}
L552:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+316))
	v1943 = base.B2i32(v1941 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v1943)
	v1945 = v1943
	goto L554
L553:
	;
	v1945 = int32(0)
	goto L554
L554:
	;
	goto L551
L555:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+20)) = uint32(v1988)
	v1991 = int64(base.Ui64(v1988) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+16)) = uint32(v1991)
	v1999 = F_pg_snprintf(m, v685+int32(464), int32(64), int32(_a_F_PostgresMainLoopOnce_71), v685+int32(16))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L571
	}
L556:
	;
	v1950 = F_GetWalRcvFlushRecPtr(m, int32(0), v685+int32(_a_F_PostgresMainLoopOnce_72))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v1964 = v685 + int32(440)
	v1967 = int32(_a_F_PostgresMainLoopOnce_73)
	v1968 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1969 = *(*int64)(unsafe.Add(mBase, uint32(v1968)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v1968)+280)) = v1969
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = v1969
	v1974 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1975 = *(*int64)(unsafe.Add(mBase, uint32(v1974)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v1974)+272)) = v1975
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = v1975
	if v1964 != 0 {
		goto L568
	} else {
		goto L569
	}
L559:
	;
	v1954 = F_GetXLogReplayRecPtr(m, v685+int32(464))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L560
	}
L560:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+440)) = v1956
	if base.Ui64(v1954) < base.Ui64(v1950) {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v1959 = v1950
	goto L563
L562:
	;
	v1959 = v1954
	goto L563
L563:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72])))
	if v1956 == v1960 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v1962 = v1959
	goto L566
L565:
	;
	v1962 = v1954
	goto L566
L566:
	;
	v1988 = v1962
	goto L555
L567:
	;
	v1988 = v1984
	goto L555
L568:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v1964))) = v1981
	goto L570
L569:
	;
	goto L570
L570:
	;
	v1984 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70]))
	goto L567
L571:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	if v2003 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	v2016 = int32(0)
	goto L574
L574:
	;
	v2018 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L579
	}
L575:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	v2010 = F_get_database_name(m, v2009)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	v2012 = F_MemoryContextStrdup(m, v2005, v2010)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v2016 = v2012
	goto L574
L579:
	;
	v2021 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	F_TupleDescInitBuiltinEntry(m, v2021, int32(1), int32(_a_F_PostgresMainLoopOnce_74), int32(25))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	F_TupleDescInitBuiltinEntry(m, v2021, int32(2), int32(_a_F_PostgresMainLoopOnce_43), int32(20))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	F_TupleDescInitBuiltinEntry(m, v2021, int32(3), int32(_a_F_PostgresMainLoopOnce_75), int32(25))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	F_TupleDescInitBuiltinEntry(m, v2021, int32(4), int32(_a_F_PostgresMainLoopOnce_76), int32(25))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	v2044 = F_begin_tup_output_tupdesc(m, v2018, v2021, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	v2048 = F_cstring_to_text(m, v685+int32(_a_F_PostgresMainLoopOnce_69))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = v2048
	v2051 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v685)+440)))
	v2052 = F_Int64GetDatum(m, v2051)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = v2052
	v2057 = F_cstring_to_text(m, v685+int32(464))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = v2057
	if v2016 != 0 {
		goto L590
	} else {
		goto L591
	}
L589:
	;
	F_do_tup_output(m, v2044, v685+int32(_a_F_PostgresMainLoopOnce_72), v685+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L594
	}
L590:
	;
	v2060 = F_cstring_to_text(m, v2016)
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L1
	} else {
		goto L593
	}
L591:
	;
	goto L592
L592:
	;
	v2063 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[75]))) = uint8(v2063)
	goto L589
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[76]))) = v2060
	goto L589
L594:
	;
	F_end_tup_output(m, v2044)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_79)
	goto L234
L596:
	;
	F_TupleDescInitBuiltinEntry(m, v2080, int32(1), int32(_a_F_PostgresMainLoopOnce_80), int32(25))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	F_TupleDescInitBuiltinEntry(m, v2080, int32(2), int32(_a_F_PostgresMainLoopOnce_81), int32(25))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	F_TupleDescInitBuiltinEntry(m, v2080, int32(3), int32(_a_F_PostgresMainLoopOnce_82), int32(20))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	v2097 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+462)) = uint8(v2097)
	v2099 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+460)) = uint16(v2099)
	v2102 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	v2106 = F_LWLockAcquire(m, v2102+int32(_a_F_PostgresMainLoopOnce_83), v2097)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2074)+4))
	v2110 = F_SearchNamedReplicationSlot(m, v2108, int32(0))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L603
	}
L601:
	;
	v2217 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L1
	} else {
		goto L633
	}
L602:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	*(*int32)(unsafe.Add(mBase, uint32(v2110))) = int32(1)
	if v2119 != 0 {
		goto L609
	} else {
		goto L610
	}
L603:
	;
	if v2110 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v2112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2110)+4)))
	if v2112 != 0 {
		goto L602
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	F_LWLockRelease(m, v2114+int32(_a_F_PostgresMainLoopOnce_83))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L1
	} else {
		goto L608
	}
L607:
	;
	goto L606
L608:
	;
	goto L601
L609:
	;
	F_s_lock(m, v2110, int32(_a_F_PostgresMainLoopOnce_14), int32(511), int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L1
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	base.MemoryCopy(m, v685+int32(_a_F_PostgresMainLoopOnce_69), v2110, int32(88))
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+456)) = v2131
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+448)) = v2133
	v2135 = *(*int64)(unsafe.Add(mBase, uint32(v2110)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v685)+440)) = v2135
	v2137 = *(*int64)(unsafe.Add(mBase, uint32(v2110)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v685)+432)) = v2137
	base.MemoryCopy(m, v685+int32(464), v2110+int32(112), int32(176))
	*(*int32)(unsafe.Add(mBase, uint32(v2110))) = int32(0)
	v2148 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	F_LWLockRelease(m, v2148+int32(_a_F_PostgresMainLoopOnce_83))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L613
	}
L612:
	;
	goto L611
L613:
	;
	if v2131 != 0 {
		goto L258
	} else {
		goto L614
	}
L614:
	;
	v2154 = F_cstring_to_text(m, int32(_a_F_PostgresMainLoopOnce_42))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	v2156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+460)) = uint8(v2156)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = v2154
	if v2137 == int64(0) {
		goto L601
	} else {
		goto L616
	}
L616:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+52)) = uint32(v2137)
	v2163 = int64(base.Ui64(v2137) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+48)) = uint32(v2163)
	v2166 = v685 + int32(_a_F_PostgresMainLoopOnce_72)
	v2171 = F_pg_snprintf(m, v2166, int32(64), int32(_a_F_PostgresMainLoopOnce_71), v685+int32(48))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	v2173 = F_cstring_to_text(m, v2166)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	v2175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+461)) = uint8(v2175)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v2173
	v2178 = *(*int64)(unsafe.Add(mBase, uint32(v685)+432))
	if v2178 == int64(0) {
		goto L601
	} else {
		goto L619
	}
L619:
	;
	v2183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v2183 == int32(1) {
		goto L622
	} else {
		goto L623
	}
L620:
	;
	v2202 = F_readTimeLineHistory(m, v2201)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L1
	} else {
		goto L630
	}
L621:
	;
	if v2193 != 0 {
		goto L625
	} else {
		goto L626
	}
L622:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+316))
	v2191 = base.B2i32(v2189 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v2191)
	v2193 = v2191
	goto L624
L623:
	;
	v2193 = int32(0)
	goto L624
L624:
	;
	goto L621
L625:
	;
	v2194 = F_GetXLogReplayRecPtr(m, v2166)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L1
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2198)+308))
	goto L629
L628:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72])))
	v2201 = v2196
	goto L620
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = v2199
	v2201 = v2199
	goto L620
L630:
	;
	v2204 = *(*int64)(unsafe.Add(mBase, uint32(v685)+432))
	v2205 = F_tliOfPointInHistory(m, v2204, v2202)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	v2208 = F_Int64GetDatum(m, base.I64_extend_i32_u(v2205))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	v2210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+462)) = uint8(v2210)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[65]))) = v2208
	goto L601
L633:
	;
	v2220 = F_begin_tup_output_tupdesc(m, v2217, v2080, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	F_do_tup_output(m, v2220, v685+int32(_a_F_PostgresMainLoopOnce_78), v685+int32(460))
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	F_end_tup_output(m, v2220)
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_85)
	goto L234
L637:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+8))
	if v2589 == int32(0) {
		goto L729
	} else {
		goto L730
	}
L638:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2234)+4))
	if v2237 <= int32(0) {
		v2568 = v1
		v2571 = v1
		v2573 = v1
		v2574 = v1
		goto L637
	} else {
		goto L639
	}
L639:
	;
	v2241 = int32(0)
	v2257 = v1
	v2258 = v1
	v2260 = v1
	v2262 = v1
	v2263 = v1
	v2266 = v1
	v2267 = v1
	v2271 = v1
	goto L640
L640:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2234)+12))
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2278+v2241<<(uint(int32(2))%32))))
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+8))
	v2284 = int32(_a_F_PostgresMainLoopOnce_57)
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283))))
	v2290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[79])))
	if base.B2i32(v2287 == int32(0))|base.B2i32(v2287 != v2290) != 0 {
		v2308 = v2287
		v2309 = v2290
		goto L644
	} else {
		goto L645
	}
L641:
	;
	v2568 = v2540
	v2571 = v2542
	v2573 = v2543
	v2574 = v2544
	goto L637
L642:
	;
	v2549 = v2241 + int32(1)
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2234)+4))
	if v2549 < v2550 {
		v2241 = v2549
		v2257 = v2540
		v2258 = v2541
		v2260 = v2542
		v2262 = v2543
		v2263 = v2544
		v2266 = v2545
		v2267 = v2546
		v2271 = v2547
		goto L640
	} else {
		goto L727
	}
L643:
	;
	if v2308-v2309 == int32(0) {
		goto L650
	} else {
		goto L651
	}
L644:
	;
	goto L643
L645:
	;
	v2293 = v2283
	v2294 = v2284
	goto L646
L646:
	;
	v2297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2294)+1)))
	v2298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2293)+1)))
	if v2298 == int32(0) {
		v2308 = v2298
		v2309 = v2297
		goto L644
	} else {
		goto L648
	}
L647:
	;
	v2308 = v2298
	v2309 = v2297
	goto L644
L648:
	;
	v2301 = int32(1)
	if v2298 == v2297 {
		v2293 = v2293 + v2301
		v2294 = v2294 + v2301
		goto L646
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	if v2266&int32(1) != 0 {
		goto L257
	} else {
		goto L653
	}
L651:
	;
	goto L652
L652:
	;
	v2434 = int32(_a_F_PostgresMainLoopOnce_39)
	v2437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283))))
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[80])))
	if base.B2i32(v2437 == int32(0))|base.B2i32(v2437 != v2440) != 0 {
		v2458 = v2437
		v2459 = v2440
		goto L691
	} else {
		goto L692
	}
L653:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+8))
	if v2315 != int32(1) {
		goto L257
	} else {
		goto L654
	}
L654:
	;
	v2318 = F_defGetString(m, v2282)
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	v2320 = int32(_a_F_PostgresMainLoopOnce_55)
	v2323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2318))))
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[81])))
	if base.B2i32(v2323 == int32(0))|base.B2i32(v2323 != v2326) != 0 {
		v2344 = v2323
		v2345 = v2326
		goto L657
	} else {
		goto L658
	}
L656:
	;
	if v2344-v2345 == int32(0) {
		goto L663
	} else {
		goto L664
	}
L657:
	;
	goto L656
L658:
	;
	v2329 = v2318
	v2330 = v2320
	goto L659
L659:
	;
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2330)+1)))
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2329)+1)))
	if v2334 == int32(0) {
		v2344 = v2334
		v2345 = v2333
		goto L657
	} else {
		goto L661
	}
L660:
	;
	v2344 = v2334
	v2345 = v2333
	goto L657
L661:
	;
	v2337 = int32(1)
	if v2334 == v2333 {
		v2329 = v2329 + v2337
		v2330 = v2330 + v2337
		goto L659
	} else {
		goto L662
	}
L662:
	;
	goto L660
L663:
	;
	v2540 = v2257
	v2541 = v2258
	v2542 = v2260
	v2543 = v2262
	v2544 = int32(0)
	v2545 = int32(1)
	v2546 = v2267
	v2547 = v2271
	goto L642
L664:
	;
	goto L665
L665:
	;
	v2351 = int32(1)
	v2352 = int32(_a_F_PostgresMainLoopOnce_54)
	v2355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2318))))
	v2358 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[82])))
	if base.B2i32(v2355 == int32(0))|base.B2i32(v2355 != v2358) != 0 {
		v2376 = v2355
		v2377 = v2358
		goto L667
	} else {
		goto L668
	}
L666:
	;
	if v2376-v2377 == int32(0) {
		goto L673
	} else {
		goto L674
	}
L667:
	;
	goto L666
L668:
	;
	v2361 = v2318
	v2362 = v2352
	goto L669
L669:
	;
	v2365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2362)+1)))
	v2366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2361)+1)))
	if v2366 == int32(0) {
		v2376 = v2366
		v2377 = v2365
		goto L667
	} else {
		goto L671
	}
L670:
	;
	v2376 = v2366
	v2377 = v2365
	goto L667
L671:
	;
	v2369 = int32(1)
	if v2366 == v2365 {
		v2361 = v2361 + v2369
		v2362 = v2362 + v2369
		goto L669
	} else {
		goto L672
	}
L672:
	;
	goto L670
L673:
	;
	v2540 = v2257
	v2541 = v2258
	v2542 = v2260
	v2543 = v2262
	v2544 = int32(1)
	v2545 = v2351
	v2546 = v2267
	v2547 = v2271
	goto L642
L674:
	;
	goto L675
L675:
	;
	v2382 = int32(_a_F_PostgresMainLoopOnce_53)
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2318))))
	v2388 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[83])))
	if base.B2i32(v2385 == int32(0))|base.B2i32(v2385 != v2388) != 0 {
		v2406 = v2385
		v2407 = v2388
		goto L677
	} else {
		goto L678
	}
L676:
	;
	if v2406-v2407 == int32(0) {
		goto L683
	} else {
		goto L684
	}
L677:
	;
	goto L676
L678:
	;
	v2391 = v2318
	v2392 = v2382
	goto L679
L679:
	;
	v2395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2392)+1)))
	v2396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2391)+1)))
	if v2396 == int32(0) {
		v2406 = v2396
		v2407 = v2395
		goto L677
	} else {
		goto L681
	}
L680:
	;
	v2406 = v2396
	v2407 = v2395
	goto L677
L681:
	;
	v2399 = int32(1)
	if v2396 == v2395 {
		v2391 = v2391 + v2399
		v2392 = v2392 + v2399
		goto L679
	} else {
		goto L682
	}
L682:
	;
	goto L680
L683:
	;
	v2540 = v2257
	v2541 = v2258
	v2542 = v2260
	v2543 = v2262
	v2544 = int32(2)
	v2545 = v2351
	v2546 = v2267
	v2547 = v2271
	goto L642
L684:
	;
	goto L685
L685:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+200)) = v2318
	*(*int32)(unsafe.Add(mBase, uint32(v685)+196)) = v2419
	*(*int32)(unsafe.Add(mBase, uint32(v685)+192)) = int32(_a_F_PostgresMainLoopOnce_86)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_87), v685+int32(192))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L688
	}
L688:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1152), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L690:
	;
	if v2458-v2459 == int32(0) {
		goto L697
	} else {
		goto L698
	}
L691:
	;
	goto L690
L692:
	;
	v2443 = v2283
	v2444 = v2434
	goto L693
L693:
	;
	v2447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2444)+1)))
	v2448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2443)+1)))
	if v2448 == int32(0) {
		v2458 = v2448
		v2459 = v2447
		goto L691
	} else {
		goto L695
	}
L694:
	;
	v2458 = v2448
	v2459 = v2447
	goto L691
L695:
	;
	v2451 = int32(1)
	if v2448 == v2447 {
		v2443 = v2443 + v2451
		v2444 = v2444 + v2451
		goto L693
	} else {
		goto L696
	}
L696:
	;
	goto L694
L697:
	;
	if v2271&int32(1) != 0 {
		goto L256
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	v2469 = int32(_a_F_PostgresMainLoopOnce_37)
	v2472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283))))
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[84])))
	if base.B2i32(v2472 == int32(0))|base.B2i32(v2472 != v2475) != 0 {
		v2493 = v2472
		v2494 = v2475
		goto L704
	} else {
		goto L705
	}
L700:
	;
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+8))
	if v2465 != 0 {
		goto L256
	} else {
		goto L701
	}
L701:
	;
	v2467 = F_defGetBoolean(m, v2282)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	v2540 = v2257
	v2541 = v2258
	v2542 = v2260
	v2543 = v2467
	v2544 = v2263
	v2545 = v2266
	v2546 = v2267
	v2547 = int32(1)
	goto L642
L703:
	;
	if v2493-v2494 == int32(0) {
		goto L710
	} else {
		goto L711
	}
L704:
	;
	goto L703
L705:
	;
	v2478 = v2283
	v2479 = v2469
	goto L706
L706:
	;
	v2482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2479)+1)))
	v2483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478)+1)))
	if v2483 == int32(0) {
		v2493 = v2483
		v2494 = v2482
		goto L704
	} else {
		goto L708
	}
L707:
	;
	v2493 = v2483
	v2494 = v2482
	goto L704
L708:
	;
	v2486 = int32(1)
	if v2483 == v2482 {
		v2478 = v2478 + v2486
		v2479 = v2479 + v2486
		goto L706
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	if v2258 != 0 {
		goto L255
	} else {
		goto L713
	}
L711:
	;
	goto L712
L712:
	;
	v2504 = int32(_a_F_PostgresMainLoopOnce_89)
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283))))
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[85])))
	if base.B2i32(v2507 == int32(0))|base.B2i32(v2507 != v2510) != 0 {
		v2528 = v2507
		v2529 = v2510
		goto L717
	} else {
		goto L718
	}
L713:
	;
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+8))
	if v2498 != int32(1) {
		goto L255
	} else {
		goto L714
	}
L714:
	;
	v2502 = F_defGetBoolean(m, v2282)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	v2540 = v2257
	v2541 = int32(1)
	v2542 = v2502
	v2543 = v2262
	v2544 = v2263
	v2545 = v2266
	v2546 = v2267
	v2547 = v2271
	goto L642
L716:
	;
	if v2528-v2529 != 0 {
		goto L253
	} else {
		goto L723
	}
L717:
	;
	goto L716
L718:
	;
	v2513 = v2283
	v2514 = v2504
	goto L719
L719:
	;
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2514)+1)))
	v2518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2513)+1)))
	if v2518 == int32(0) {
		v2528 = v2518
		v2529 = v2517
		goto L717
	} else {
		goto L721
	}
L720:
	;
	v2528 = v2518
	v2529 = v2517
	goto L717
L721:
	;
	v2521 = int32(1)
	if v2518 == v2517 {
		v2513 = v2513 + v2521
		v2514 = v2514 + v2521
		goto L719
	} else {
		goto L722
	}
L722:
	;
	goto L720
L723:
	;
	if v2267&int32(1) != 0 {
		goto L254
	} else {
		goto L724
	}
L724:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+8))
	if v2533 != int32(1) {
		goto L254
	} else {
		goto L725
	}
L725:
	;
	v2537 = F_defGetBoolean(m, v2282)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	v2540 = v2537
	v2541 = v2258
	v2542 = v2260
	v2543 = v2262
	v2544 = v2263
	v2545 = v2266
	v2546 = int32(1)
	v2547 = v2271
	goto L642
L727:
	;
	goto L641
L728:
	;
	v2804 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v2805 = *(*int64)(unsafe.Add(mBase, uint32(v2804)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+84)) = uint32(v2805)
	v2808 = int64(base.Ui64(v2805) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+80)) = uint32(v2808)
	v2811 = v685 + int32(464)
	v2816 = F_pg_snprintf(m, v2811, int32(64), int32(_a_F_PostgresMainLoopOnce_71), v685+int32(80))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L1
	} else {
		goto L790
	}
L729:
	;
	v2592 = int32(0)
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+4))
	v2595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231)+16)))
	F_ReplicationSlotCreate(m, v2593, v2592, v2595<<(uint(int32(1))%32)&int32(2), v2592, v2592, v2592)
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L732
	}
L730:
	;
	goto L731
L731:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L1
	} else {
		goto L738
	}
L732:
	;
	if v2573&int32(1) == int32(0) {
		v2802 = v2592
		goto L728
	} else {
		goto L733
	}
L733:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	v2613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231)+16)))
	if v2613 != 0 {
		v2802 = v2592
		goto L728
	} else {
		goto L736
	}
L736:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	v2802 = v2592
	goto L728
L738:
	;
	v2618 = int32(0)
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+4))
	v2620 = int32(1)
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231)+16)))
	if v2623 != 0 {
		goto L739
	} else {
		goto L740
	}
L739:
	;
	v2624 = int32(2)
	goto L741
L740:
	;
	v2624 = v2620
	goto L741
L741:
	;
	v2625 = int32(1)
	F_ReplicationSlotCreate(m, v2619, v2620, v2624, v2571&v2625, v2568&v2625, int32(0))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	switch v2574 {
	case 0:
		goto L745
	default:
		v2682 = int32(0)
		goto L743
	case 2:
		goto L744
	}
L743:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = int32(1032)
	v2696 = F_CreateInitDecodingContext(m, v2683, v2682, int64(0), v685+int32(_a_F_PostgresMainLoopOnce_72), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L1
	} else {
		goto L758
	}
L744:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2658)+24))
	goto L751
L745:
	;
	v2633 = int32(1)
	v2635 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+24))
	goto L746
L746:
	;
	if base.B2i32(base.Ui32(v2633) < base.Ui32(v2636)) == int32(0) {
		v2682 = v2633
		goto L743
	} else {
		goto L747
	}
L747:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+96)) = int32(_a_F_PostgresMainLoopOnce_90)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_91), v685+int32(96))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1258), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L1
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
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2659)) == int32(0) {
		goto L252
	} else {
		goto L752
	}
L752:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	if v2665 != int32(2) {
		goto L251
	} else {
		goto L753
	}
L753:
	;
	v2669 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[88])))
	if v2669 == int32(0) {
		goto L250
	} else {
		goto L754
	}
L754:
	;
	v2672 = int32(1)
	v2674 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89])))
	if v2674 == v2672 {
		goto L249
	} else {
		goto L755
	}
L755:
	;
	v2678 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v2678)+28))
	goto L756
L756:
	;
	if int32(1) < v2679 {
		goto L248
	} else {
		goto L757
	}
L757:
	;
	v2682 = v2672
	goto L743
L758:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90])) = int64(0)
	F_DecodingContextFindStartpoint(m, v2696)
	mBase = m.M
	v2702 = m.ExcPending
	if v2702 != 0 {
		goto L1
	} else {
		goto L759
	}
L759:
	;
	switch v2574 {
	case 0:
		goto L762
	default:
		v2792 = v2618
		goto L760
	case 2:
		goto L761
	}
L760:
	;
	F_FreeDecodingContext(m, v2696)
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
		goto L1
	} else {
		goto L787
	}
L761:
	;
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+16))
	v2783 = F_SnapBuildInitialSnapshot(m, v2782)
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L785
	}
L762:
	;
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+16))
	v2704 = m.G0
	v2706 = v2704 - int32(16)
	m.G0 = v2706
	v2709 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2709)+24))
	goto L765
L763:
	;
	v2792 = v2734
	goto L760
L764:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2772 = m.ExcPending
	if v2772 != 0 {
		goto L1
	} else {
		goto L782
	}
L765:
	;
	if base.B2i32(v2710 != int32(0)) == int32(0) {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49]))
	if v2716 != 0 {
		goto L764
	} else {
		goto L769
	}
L767:
	;
	goto L768
L768:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L1
	} else {
		goto L779
	}
L769:
	;
	v2718 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48])) = uint8(v2718)
	v2722 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v2722
	F_StartTransactionCommand(m)
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L1
	} else {
		goto L770
	}
L770:
	;
	v2727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[88])) = uint8(v2727)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87])) = int32(2)
	v2732 = F_SnapBuildInitialSnapshot(m, v2703)
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	v2734 = F_ExportSnapshot(m, v2732)
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	v2738 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	if v2738 != 0 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2732)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2706)+4)) = v2740
	*(*int32)(unsafe.Add(mBase, uint32(v2706))) = v2734
	F_errmsg_plural(m, int32(_a_F_PostgresMainLoopOnce_93), int32(_a_F_PostgresMainLoopOnce_94), v2740, v2706)
	mBase = m.M
	v2746 = m.ExcPending
	if v2746 != 0 {
		goto L1
	} else {
		goto L777
	}
L775:
	;
	goto L776
L776:
	;
	m.G0 = v2706 + int32(16)
	goto L763
L777:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(571), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	goto L776
L779:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_96), int32(0))
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L1
	} else {
		goto L780
	}
L780:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(545), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L1
	} else {
		goto L781
	}
L781:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L782:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_97), int32(0))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L1
	} else {
		goto L783
	}
L783:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(548), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2781 = m.ExcPending
	if v2781 != 0 {
		goto L1
	} else {
		goto L784
	}
L784:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L785:
	;
	v2786 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	F_RestoreTransactionSnapshot(m, v2783, v2786)
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L1
	} else {
		goto L786
	}
L786:
	;
	v2792 = v2618
	goto L760
L787:
	;
	v2795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231)+16)))
	if v2795 != 0 {
		v2802 = v2792
		goto L728
	} else {
		goto L788
	}
L788:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		goto L1
	} else {
		goto L789
	}
L789:
	;
	v2802 = v2792
	goto L728
L790:
	;
	v2819 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2820 = m.ExcPending
	if v2820 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	v2822 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	F_TupleDescInitBuiltinEntry(m, v2822, int32(1), int32(_a_F_PostgresMainLoopOnce_98), int32(25))
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	F_TupleDescInitBuiltinEntry(m, v2822, int32(2), int32(_a_F_PostgresMainLoopOnce_99), int32(25))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L1
	} else {
		goto L794
	}
L794:
	;
	F_TupleDescInitBuiltinEntry(m, v2822, int32(3), int32(_a_F_PostgresMainLoopOnce_100), int32(25))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L795
	}
L795:
	;
	F_TupleDescInitBuiltinEntry(m, v2822, int32(4), int32(_a_F_PostgresMainLoopOnce_101), int32(25))
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	v2845 = F_begin_tup_output_tupdesc(m, v2819, v2822, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2846 = m.ExcPending
	if v2846 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v2851 = F_cstring_to_text(m, v2848+int32(24))
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[91]))) = v2851
	v2854 = F_cstring_to_text(m, v2811)
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L799
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[92]))) = v2854
	if v2802 != 0 {
		goto L801
	} else {
		goto L802
	}
L800:
	;
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+12))
	if v2862 != 0 {
		goto L806
	} else {
		goto L807
	}
L801:
	;
	v2857 = F_cstring_to_text(m, v2802)
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L1
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v2860 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[93]))) = uint8(v2860)
	goto L800
L804:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[94]))) = v2857
	goto L800
L805:
	;
	F_do_tup_output(m, v2845, v685+int32(_a_F_PostgresMainLoopOnce_69), v685+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L1
	} else {
		goto L810
	}
L806:
	;
	v2863 = F_cstring_to_text(m, v2862)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L1
	} else {
		goto L809
	}
L807:
	;
	goto L808
L808:
	;
	v2866 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[75]))) = uint8(v2866)
	goto L805
L809:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[95]))) = v2863
	goto L805
L810:
	;
	F_end_tup_output(m, v2845)
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_86)
	goto L234
L813:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_102)
	goto L234
L814:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+464)) = uint8(v3059)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[91]))) = uint8(v3055)
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+4))
	v3070 = int32(0)
	v3071 = m.G0
	v3073 = v3071 - int32(1072)
	m.G0 = v3073
	F_ReplicationSlotAcquire(m, v3069, v3070, int32(1))
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L1
	} else {
		goto L849
	}
L815:
	;
	v2895 = int32(0)
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+4))
	if v2896 <= v2895 {
		v3054 = v1
		v3055 = v1
		v3059 = v1
		v3066 = v2895
		goto L814
	} else {
		goto L816
	}
L816:
	;
	v2900 = int32(0)
	v2922 = v1
	v2925 = v1
	v2926 = v1
	v2930 = v1
	goto L817
L817:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+12))
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v2937+v2900<<(uint(int32(2))%32))))
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2941)+8))
	v2943 = int32(_a_F_PostgresMainLoopOnce_89)
	v2946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2942))))
	v2949 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[85])))
	if base.B2i32(v2946 == int32(0))|base.B2i32(v2946 != v2949) != 0 {
		v2967 = v2946
		v2968 = v2949
		goto L821
	} else {
		goto L822
	}
L818:
	;
	if v3010&int32(1) != 0 {
		goto L843
	} else {
		goto L844
	}
L819:
	;
	v3014 = v2900 + int32(1)
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+4))
	if v3014 < v3015 {
		v2900 = v3014
		v2922 = v3009
		v2925 = v3010
		v2926 = v3011
		v2930 = v3012
		goto L817
	} else {
		goto L842
	}
L820:
	;
	if v2967-v2968 == int32(0) {
		goto L827
	} else {
		goto L828
	}
L821:
	;
	goto L820
L822:
	;
	v2952 = v2942
	v2953 = v2943
	goto L823
L823:
	;
	v2956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2953)+1)))
	v2957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2952)+1)))
	if v2957 == int32(0) {
		v2967 = v2957
		v2968 = v2956
		goto L821
	} else {
		goto L825
	}
L824:
	;
	v2967 = v2957
	v2968 = v2956
	goto L821
L825:
	;
	v2960 = int32(1)
	if v2957 == v2956 {
		v2952 = v2952 + v2960
		v2953 = v2953 + v2960
		goto L823
	} else {
		goto L826
	}
L826:
	;
	goto L824
L827:
	;
	if v2922&int32(1) != 0 {
		goto L247
	} else {
		goto L830
	}
L828:
	;
	goto L829
L829:
	;
	v2977 = int32(_a_F_PostgresMainLoopOnce_37)
	v2980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2942))))
	v2983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[84])))
	if base.B2i32(v2980 == int32(0))|base.B2i32(v2980 != v2983) != 0 {
		v3001 = v2980
		v3002 = v2983
		goto L833
	} else {
		goto L834
	}
L830:
	;
	v2975 = F_defGetBoolean(m, v2941)
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	v3009 = int32(1)
	v3010 = v2925
	v3011 = v2926
	v3012 = v2975
	goto L819
L832:
	;
	if v3001-v3002 != 0 {
		goto L245
	} else {
		goto L839
	}
L833:
	;
	goto L832
L834:
	;
	v2986 = v2942
	v2987 = v2977
	goto L835
L835:
	;
	v2990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+1)))
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2986)+1)))
	if v2991 == int32(0) {
		v3001 = v2991
		v3002 = v2990
		goto L833
	} else {
		goto L837
	}
L836:
	;
	v3001 = v2991
	v3002 = v2990
	goto L833
L837:
	;
	v2994 = int32(1)
	if v2991 == v2990 {
		v2986 = v2986 + v2994
		v2987 = v2987 + v2994
		goto L835
	} else {
		goto L838
	}
L838:
	;
	goto L836
L839:
	;
	if v2925&int32(1) != 0 {
		goto L246
	} else {
		goto L840
	}
L840:
	;
	v3007 = F_defGetBoolean(m, v2941)
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L841
	}
L841:
	;
	v3009 = v2922
	v3010 = int32(1)
	v3011 = v3007
	v3012 = v2930
	goto L819
L842:
	;
	goto L818
L843:
	;
	v3022 = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	goto L845
L844:
	;
	v3022 = int32(0)
	goto L845
L845:
	;
	if v3009&int32(1) != 0 {
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v3028 = v685 + int32(464)
	goto L848
L847:
	;
	v3028 = int32(0)
	goto L848
L848:
	;
	v3054 = v3022
	v3055 = v3011
	v3059 = v3012
	v3066 = v3028
	goto L814
L849:
	;
	v3080 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v3080)+88))
	if v3081 != 0 {
		goto L853
	} else {
		goto L854
	}
L850:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_103)
	goto L234
L851:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3269 = m.ExcPending
	if v3269 != 0 {
		goto L1
	} else {
		goto L908
	}
L852:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L1
	} else {
		goto L903
	}
L853:
	;
	v3084 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v3084 == int32(1) {
		goto L859
	} else {
		goto L860
	}
L854:
	;
	goto L855
L855:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3230 = m.ExcPending
	if v3230 != 0 {
		goto L1
	} else {
		goto L899
	}
L856:
	;
	if v3054 == int32(0) {
		goto L884
	} else {
		goto L885
	}
L857:
	;
	v3134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3133)+202)))
	if v3132 == v3134 {
		v3156 = v3070
		goto L856
	} else {
		goto L877
	}
L858:
	;
	if v3094 != 0 {
		goto L862
	} else {
		goto L863
	}
L859:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3089)+316))
	v3092 = base.B2i32(v3090 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v3092)
	v3094 = v3092
	goto L861
L860:
	;
	v3094 = int32(0)
	goto L861
L861:
	;
	goto L858
L862:
	;
	v3096 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3096)+201)))
	if v3097 != 0 {
		goto L852
	} else {
		goto L865
	}
L863:
	;
	goto L864
L864:
	;
	if v3066 == int32(0) {
		v3156 = v3070
		goto L856
	} else {
		goto L872
	}
L865:
	;
	if v3066 == int32(0) {
		v3156 = v3070
		goto L856
	} else {
		goto L866
	}
L866:
	;
	v3101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3066))))
	if v3101 != int32(1) {
		v3132 = int32(0)
		v3133 = v3096
		goto L857
	} else {
		goto L867
	}
L867:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		goto L1
	} else {
		goto L868
	}
L868:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L1
	} else {
		goto L869
	}
L869:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_104), int32(0))
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(913), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L1
	} else {
		goto L871
	}
L871:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L872:
	;
	v3122 = int32(1)
	v3124 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3066))))
	if v3125 != v3122 {
		goto L873
	} else {
		goto L874
	}
L873:
	;
	v3132 = int32(0)
	v3133 = v3124
	goto L857
L874:
	;
	goto L875
L875:
	;
	v3129 = *(*int32)(unsafe.Add(mBase, uint32(v3124)+92))
	if v3129 == int32(2) {
		goto L851
	} else {
		goto L876
	}
L876:
	;
	v3132 = v3122
	v3133 = v3124
	goto L857
L877:
	;
	v3136 = *(*int32)(unsafe.Add(mBase, uint32(v3133)))
	v3137 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3133))) = v3137
	if v3136 != 0 {
		goto L878
	} else {
		goto L879
	}
L878:
	;
	v3141 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	F_s_lock(m, v3141, int32(_a_F_PostgresMainLoopOnce_105), int32(929), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3146 = m.ExcPending
	if v3146 != 0 {
		goto L1
	} else {
		goto L881
	}
L879:
	;
	goto L880
L880:
	;
	v3147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3066))))
	v3149 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	*(*int32)(unsafe.Add(mBase, uint32(v3149))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3149)+202)) = uint8(v3147)
	v3156 = v3137
	goto L856
L881:
	;
	goto L880
L882:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L898
	}
L883:
	;
	v3186 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3186)))
	*(*int32)(unsafe.Add(mBase, uint32(v3186))) = int32(1)
	if v3187 != 0 {
		goto L892
	} else {
		goto L893
	}
L884:
	;
	if v3156 == int32(0) {
		goto L882
	} else {
		goto L891
	}
L885:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3160)+136)))
	v3162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3054))))
	if v3161 == v3162 {
		goto L884
	} else {
		goto L886
	}
L886:
	;
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v3160)))
	*(*int32)(unsafe.Add(mBase, uint32(v3160))) = int32(1)
	if v3164 != 0 {
		goto L887
	} else {
		goto L888
	}
L887:
	;
	v3168 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	F_s_lock(m, v3168, int32(_a_F_PostgresMainLoopOnce_105), int32(939), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3173 = m.ExcPending
	if v3173 != 0 {
		goto L1
	} else {
		goto L890
	}
L888:
	;
	goto L889
L889:
	;
	v3174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3054))))
	v3176 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	*(*int32)(unsafe.Add(mBase, uint32(v3176))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3176)+136)) = uint8(v3174)
	goto L883
L890:
	;
	goto L889
L891:
	;
	goto L883
L892:
	;
	F_s_lock(m, v3186, int32(_a_F_PostgresMainLoopOnce_105), int32(1107), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L1
	} else {
		goto L895
	}
L893:
	;
	goto L894
L894:
	;
	v3195 = int32(_a_F_PostgresMainLoopOnce_108)
	v3196 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3197 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3196)+12)) = uint16(v3197)
	*(*int32)(unsafe.Add(mBase, uint32(v3186))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3073)+16)) = int32(_a_F_PostgresMainLoopOnce_109)
	v3204 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	*(*int32)(unsafe.Add(mBase, uint32(v3073)+20)) = v3204 + int32(24)
	v3209 = v3073 + int32(48)
	v3213 = F_pg_sprintf(m, v3209, int32(_a_F_PostgresMainLoopOnce_110), v3073+int32(16))
	mBase = m.M
	v3214 = m.ExcPending
	if v3214 != 0 {
		goto L1
	} else {
		goto L896
	}
L895:
	;
	goto L894
L896:
	;
	v3216 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	F_SaveSlotToPath(m, v3216, v3209, int32(21))
	mBase = m.M
	v3219 = m.ExcPending
	if v3219 != 0 {
		goto L1
	} else {
		goto L897
	}
L897:
	;
	goto L882
L898:
	;
	m.G0 = v3073 + int32(1072)
	goto L850
L899:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3073))) = int32(_a_F_PostgresMainLoopOnce_103)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_111), v3073)
	mBase = m.M
	v3238 = m.ExcPending
	if v3238 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(891), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L903:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3250 = m.ExcPending
	if v3250 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3073)+32)) = v3069
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_112), v3073+int32(32))
	mBase = m.M
	v3256 = m.ExcPending
	if v3256 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_113), int32(0))
	mBase = m.M
	v3260 = m.ExcPending
	if v3260 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(903), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3265 = m.ExcPending
	if v3265 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L908:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3272 = m.ExcPending
	if v3272 != 0 {
		goto L1
	} else {
		goto L909
	}
L909:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_114), int32(0))
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(925), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L1
	} else {
		goto L911
	}
L911:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L912:
	;
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+4))
	if v3287 == int32(0) {
		goto L913
	} else {
		goto L914
	}
L913:
	;
	v3290 = m.G0
	v3292 = v3290 - int32(144)
	m.G0 = v3292
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+116)) = int32(1031)
	v3298 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+112)) = v3298
	v3302 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[96]))
	v3306 = F_XLogReaderAllocate(m, v3302, v3292+int32(112), v3298)
	mBase = m.M
	v3307 = m.ExcPending
	if v3307 != 0 {
		goto L1
	} else {
		goto L916
	}
L914:
	;
	goto L915
L915:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L1
	} else {
		goto L1025
	}
L916:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = v3306
	if v3306 != 0 {
		goto L924
	} else {
		goto L925
	}
L917:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L1
	} else {
		goto L1024
	}
L918:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L1
	} else {
		goto L1021
	}
L919:
	;
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+8))
	if v3616 != 0 {
		goto L1003
	} else {
		goto L1004
	}
L920:
	;
	v3502 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(v3502)+4))
	if v3503 != int32(2) {
		goto L978
	} else {
		goto L979
	}
L921:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[98])) = v3484
	v3488 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99])) = uint8(v3488)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[100])) = uint8(v3488)
	v3494 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])))
	if v3494 != int32(1) {
		goto L920
	} else {
		goto L976
	}
L922:
	;
	v3484 = int64(0)
	goto L921
L923:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L1
	} else {
		goto L972
	}
L924:
	;
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+8))
	if v3309 != 0 {
		goto L927
	} else {
		goto L928
	}
L925:
	;
	goto L926
L926:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L967
	}
L927:
	;
	v3310 = int32(1)
	F_ReplicationSlotAcquire(m, v3309, v3310, v3310)
	mBase = m.M
	v3313 = m.ExcPending
	if v3313 != 0 {
		goto L1
	} else {
		goto L930
	}
L928:
	;
	goto L929
L929:
	;
	v3320 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v3320 == int32(1) {
		goto L933
	} else {
		goto L934
	}
L930:
	;
	v3315 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3316 = *(*int32)(unsafe.Add(mBase, uint32(v3315)+88))
	if v3316 != 0 {
		goto L923
	} else {
		goto L931
	}
L931:
	;
	goto L929
L932:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])) = uint8(v3330)
	if v3330 != 0 {
		goto L937
	} else {
		goto L938
	}
L933:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(v3325)+316))
	v3328 = base.B2i32(v3326 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v3328)
	v3330 = v3328
	goto L935
L934:
	;
	v3330 = int32(0)
	goto L935
L935:
	;
	goto L932
L936:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+12))
	if v3374 != 0 {
		goto L952
	} else {
		goto L953
	}
L937:
	;
	v3335 = F_GetWalRcvFlushRecPtr(m, int32(0), v3292+int32(128))
	mBase = m.M
	v3336 = m.ExcPending
	if v3336 != 0 {
		goto L1
	} else {
		goto L940
	}
L938:
	;
	goto L939
L939:
	;
	v3349 = v3292 + int32(124)
	v3352 = int32(_a_F_PostgresMainLoopOnce_73)
	v3353 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3354 = *(*int64)(unsafe.Add(mBase, uint32(v3353)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v3353)+280)) = v3354
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = v3354
	v3359 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3360 = *(*int64)(unsafe.Add(mBase, uint32(v3359)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v3359)+272)) = v3360
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = v3360
	if v3349 != 0 {
		goto L949
	} else {
		goto L950
	}
L940:
	;
	v3339 = F_GetXLogReplayRecPtr(m, v3292+int32(80))
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+124)) = v3341
	if base.Ui64(v3339) < base.Ui64(v3335) {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v3344 = v3335
	goto L944
L943:
	;
	v3344 = v3339
	goto L944
L944:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+128))
	if v3341 == v3345 {
		goto L945
	} else {
		goto L946
	}
L945:
	;
	v3347 = v3344
	goto L947
L946:
	;
	v3347 = v3339
	goto L947
L947:
	;
	v3373 = v3347
	goto L936
L948:
	;
	v3373 = v3369
	goto L936
L949:
	;
	v3365 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3366 = *(*int32)(unsafe.Add(mBase, uint32(v3365)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v3349))) = v3366
	goto L951
L950:
	;
	goto L951
L951:
	;
	v3369 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70]))
	goto L948
L952:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[102])) = v3374
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+124))
	if v3374 == v3377 {
		goto L955
	} else {
		goto L956
	}
L953:
	;
	goto L954
L954:
	;
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[102])) = v3430
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[98])) = int64(0)
	v3436 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])) = uint8(v3436)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99])) = uint8(v3436)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[100])) = uint8(v3436)
	goto L920
L955:
	;
	v3380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])) = uint8(v3380)
	goto L922
L956:
	;
	goto L957
L957:
	;
	v3383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])) = uint8(v3383)
	v3385 = F_readTimeLineHistory(m, v3377)
	mBase = m.M
	v3386 = m.ExcPending
	if v3386 != 0 {
		goto L1
	} else {
		goto L958
	}
L958:
	;
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+12))
	v3389 = F_tliSwitchPoint(m, v3387, v3385, int32(_a_F_PostgresMainLoopOnce_115))
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	F_list_free_deep(m, v3385)
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	if v3389 == int64(0) {
		goto L922
	} else {
		goto L961
	}
L961:
	;
	v3395 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	if base.Ui64(v3395) <= base.Ui64(v3389) {
		v3484 = v3389
		goto L921
	} else {
		goto L962
	}
L962:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3400 = m.ExcPending
	if v3400 != 0 {
		goto L1
	} else {
		goto L963
	}
L963:
	;
	v3401 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+56)) = v3402
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+52)) = uint32(v3401)
	v3406 = int64(base.Ui64(v3401) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+48)) = uint32(v3406)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_116), v3292+int32(48))
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L1
	} else {
		goto L964
	}
L964:
	;
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+32)) = v3413
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+40)) = uint32(v3389)
	v3417 = int64(base.Ui64(v3389) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+36)) = uint32(v3417)
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_117), v3292+int32(32))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L1
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(914), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3428 = m.ExcPending
	if v3428 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	F_errcode(m, int32(_a_F_PostgresMainLoopOnce_119))
	mBase = m.M
	v3450 = m.ExcPending
	if v3450 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_120), int32(0))
	mBase = m.M
	v3454 = m.ExcPending
	if v3454 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_121), int32(0))
	mBase = m.M
	v3458 = m.ExcPending
	if v3458 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(826), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L972:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_122), int32(0))
	mBase = m.M
	v3474 = m.ExcPending
	if v3474 != 0 {
		goto L1
	} else {
		goto L974
	}
L974:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(843), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L1
	} else {
		goto L975
	}
L975:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L976:
	;
	v3497 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	if base.Ui64(v3484) <= base.Ui64(v3497) {
		goto L919
	} else {
		goto L977
	}
L977:
	;
	goto L920
L978:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v3502)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3502)+76)) = int32(1)
	if v3506 != 0 {
		goto L981
	} else {
		goto L982
	}
L979:
	;
	goto L980
L980:
	;
	v3521 = v3292 + int32(128)
	F_pq_beginmessage(m, v3521, int32(87))
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L1
	} else {
		goto L985
	}
L981:
	;
	F_s_lock(m, v3502+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		goto L1
	} else {
		goto L984
	}
L982:
	;
	goto L983
L983:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3502)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3502)+4)) = int32(2)
	goto L980
L984:
	;
	goto L983
L985:
	;
	F_enlargeStringInfo(m, v3521, int32(1))
	mBase = m.M
	v3527 = m.ExcPending
	if v3527 != 0 {
		goto L1
	} else {
		goto L986
	}
L986:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+132))
	v3529 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+128))
	v3531 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3528+v3529))) = uint8(v3531)
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+132)) = v3528 + int32(1)
	F_enlargeStringInfo(m, v3521, int32(2))
	mBase = m.M
	v3538 = m.ExcPending
	if v3538 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+132))
	v3540 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+128))
	v3542 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3539+v3540))) = uint16(v3542)
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+132)) = v3539 + int32(2)
	F_pq_endmessage(m, v3521)
	mBase = m.M
	v3548 = m.ExcPending
	if v3548 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	v3550 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103]))
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3550)+4))
	v3552 = m.T0[v3551].(func(*base.Module) int32)(m)
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	v3554 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	if base.Ui64(v3373) < base.Ui64(v3554) {
		goto L918
	} else {
		goto L990
	}
L990:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])) = v3554
	v3559 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3559)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3559)+76)) = int32(1)
	if v3560 != 0 {
		goto L991
	} else {
		goto L992
	}
L991:
	;
	v3564 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	F_s_lock(m, v3564+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(965), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L994
	}
L992:
	;
	goto L993
L993:
	;
	v3573 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104]))
	v3575 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v3575)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3575)+8)) = v3573
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L1
	} else {
		goto L995
	}
L994:
	;
	goto L993
L995:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = int32(1)
	F_WalSndLoop(m, int32(1037))
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L1
	} else {
		goto L996
	}
L996:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = int32(0)
	v3591 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v3591 != 0 {
		goto L917
	} else {
		goto L997
	}
L997:
	;
	v3593 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v3593)+4))
	if v3594 == int32(0) {
		goto L919
	} else {
		goto L998
	}
L998:
	;
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v3593)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3593)+76)) = int32(1)
	if v3597 != 0 {
		goto L999
	} else {
		goto L1000
	}
L999:
	;
	F_s_lock(m, v3593+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3606 = m.ExcPending
	if v3606 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1000:
	;
	goto L1001
L1001:
	;
	v3607 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3593)+76)) = v3607
	*(*int32)(unsafe.Add(mBase, uint32(v3593)+4)) = v3607
	goto L919
L1002:
	;
	goto L1001
L1003:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1004:
	;
	goto L1005
L1005:
	;
	v3620 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])))
	if v3620 == int32(1) {
		goto L1007
	} else {
		goto L1008
	}
L1006:
	;
	goto L1005
L1007:
	;
	v3624 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[98]))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+20)) = uint32(v3624)
	v3626 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3292)+70)) = uint16(v3626)
	v3629 = int64(base.Ui64(v3624) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+16)) = uint32(v3629)
	v3632 = v3292 + int32(80)
	v3637 = F_pg_snprintf(m, v3632, int32(18), int32(_a_F_PostgresMainLoopOnce_71), v3292+int32(16))
	mBase = m.M
	v3638 = m.ExcPending
	if v3638 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1008:
	;
	goto L1009
L1009:
	;
	F_EndReplicationCommand(m, int32(_a_F_PostgresMainLoopOnce_123))
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L1
	} else {
		goto L1020
	}
L1010:
	;
	v3640 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3641 = m.ExcPending
	if v3641 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	v3643 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3644 = m.ExcPending
	if v3644 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1012:
	;
	F_TupleDescInitBuiltinEntry(m, v3643, int32(1), int32(_a_F_PostgresMainLoopOnce_124), int32(20))
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1013:
	;
	F_TupleDescInitBuiltinEntry(m, v3643, int32(2), int32(_a_F_PostgresMainLoopOnce_125), int32(25))
	mBase = m.M
	v3654 = m.ExcPending
	if v3654 != 0 {
		goto L1
	} else {
		goto L1014
	}
L1014:
	;
	v3656 = F_begin_tup_output_tupdesc(m, v3640, v3643, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v3657 = m.ExcPending
	if v3657 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	v3659 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106])))
	v3660 = F_Int64GetDatum(m, v3659)
	mBase = m.M
	v3661 = m.ExcPending
	if v3661 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+72)) = v3660
	v3663 = F_cstring_to_text(m, v3632)
	mBase = m.M
	v3664 = m.ExcPending
	if v3664 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3292)+76)) = v3663
	F_do_tup_output(m, v3656, v3292+int32(72), v3292+int32(70))
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L1
	} else {
		goto L1018
	}
L1018:
	;
	F_end_tup_output(m, v3656)
	mBase = m.M
	v3673 = m.ExcPending
	if v3673 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1019:
	;
	goto L1009
L1020:
	;
	m.G0 = v3292 + int32(144)
	v5051 = v3282
	goto L234
L1021:
	;
	v3688 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+4)) = uint32(v3688)
	v3690 = int64(32)
	v3691 = int64(base.Ui64(v3688) >> (uint(v3690) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292))) = uint32(v3691)
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+12)) = uint32(v3373)
	v3695 = int64(base.Ui64(v3373) >> (uint(v3690) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3292)+8)) = uint32(v3695)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_126), v3292)
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1022:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(958), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1024:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1025:
	;
	v3710 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+8))
	v3711 = int32(1)
	F_ReplicationSlotAcquire(m, v3710, v3711, v3711)
	mBase = m.M
	v3714 = m.ExcPending
	if v3714 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	v3716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])))
	if v3716 != int32(1) {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+24))
	v3749 = *(*int64)(unsafe.Add(mBase, uint32(v1914)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = int32(1032)
	v3763 = F_CreateDecodingContext(m, v3749, v3748, int32(0), v685+int32(_a_F_PostgresMainLoopOnce_72), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1028:
	;
	v3721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v3721 == int32(1) {
		goto L1030
	} else {
		goto L1031
	}
L1029:
	;
	if v3731 != 0 {
		goto L1027
	} else {
		goto L1033
	}
L1030:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3727 = *(*int32)(unsafe.Add(mBase, uint32(v3726)+316))
	v3729 = base.B2i32(v3727 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v3729)
	v3731 = v3729
	goto L1032
L1031:
	;
	v3731 = int32(0)
	goto L1032
L1032:
	;
	goto L1029
L1033:
	;
	v3734 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	if v3734 != 0 {
		goto L1035
	} else {
		goto L1036
	}
L1035:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_127), int32(0))
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1036:
	;
	goto L1037
L1037:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45])) = int32(1)
	goto L1027
L1038:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1467), int32(_a_F_PostgresMainLoopOnce_128))
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	goto L1037
L1040:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107])) = v3763
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = v3767
	v3770 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3771 = *(*int32)(unsafe.Add(mBase, uint32(v3770)+4))
	if v3771 != int32(2) {
		goto L1041
	} else {
		goto L1042
	}
L1041:
	;
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(v3770)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3770)+76)) = int32(1)
	if v3774 != 0 {
		goto L1044
	} else {
		goto L1045
	}
L1042:
	;
	goto L1043
L1043:
	;
	v3789 = v685 + int32(464)
	F_pq_beginmessage(m, v3789, int32(87))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1044:
	;
	F_s_lock(m, v3770+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3783 = m.ExcPending
	if v3783 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1045:
	;
	goto L1046
L1046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3770)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3770)+4)) = int32(2)
	goto L1043
L1047:
	;
	goto L1046
L1048:
	;
	F_enlargeStringInfo(m, v3789, int32(1))
	mBase = m.M
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1049:
	;
	v3796 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v3799 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3796+v3797))) = uint8(v3799)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v3796 + int32(1)
	F_enlargeStringInfo(m, v3789, int32(2))
	mBase = m.M
	v3806 = m.ExcPending
	if v3806 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	v3807 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v3808 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v3810 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3807+v3808))) = uint16(v3810)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v3807 + int32(2)
	F_pq_endmessage(m, v3789)
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1051:
	;
	v3818 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103]))
	v3819 = *(*int32)(unsafe.Add(mBase, uint32(v3818)+4))
	v3820 = m.T0[v3819].(func(*base.Module) int32)(m)
	mBase = m.M
	v3821 = m.ExcPending
	if v3821 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	v3823 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107]))
	v3824 = *(*int32)(unsafe.Add(mBase, uint32(v3823)+8))
	v3826 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3827 = *(*int64)(unsafe.Add(mBase, uint32(v3826)+104))
	F_XLogBeginRead(m, v3824, v3827)
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1053:
	;
	v3832 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3833 = *(*int64)(unsafe.Add(mBase, uint32(v3832)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])) = v3833
	v3836 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(v3836)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+76)) = int32(1)
	if v3837 != 0 {
		goto L1054
	} else {
		goto L1055
	}
L1054:
	;
	v3841 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	F_s_lock(m, v3841+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(1507), int32(_a_F_PostgresMainLoopOnce_128))
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L1
	} else {
		goto L1057
	}
L1055:
	;
	goto L1056
L1056:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86]))
	v3851 = *(*int64)(unsafe.Add(mBase, uint32(v3850)+104))
	v3853 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v3853)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3853)+8)) = v3851
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = int32(1)
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1057:
	;
	goto L1056
L1058:
	;
	F_WalSndLoop(m, int32(1036))
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1059:
	;
	v3866 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107]))
	F_FreeDecodingContext(m, v3866)
	mBase = m.M
	v3868 = m.ExcPending
	if v3868 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1061:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = int32(0)
	v3875 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v3875 != 0 {
		goto L244
	} else {
		goto L1062
	}
L1062:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3877)+4))
	if v3878 != 0 {
		goto L1063
	} else {
		goto L1064
	}
L1063:
	;
	v3879 = *(*int32)(unsafe.Add(mBase, uint32(v3877)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3877)+76)) = int32(1)
	if v3879 != 0 {
		goto L1066
	} else {
		goto L1067
	}
L1064:
	;
	goto L1065
L1065:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[94]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[91]))) = int32(56)
	F_EndCommand(m, v685+int32(_a_F_PostgresMainLoopOnce_69), int32(2))
	mBase = m.M
	v3901 = m.ExcPending
	if v3901 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1066:
	;
	F_s_lock(m, v3877+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3888 = m.ExcPending
	if v3888 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1067:
	;
	goto L1068
L1068:
	;
	v3889 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3877)+76)) = v3889
	*(*int32)(unsafe.Add(mBase, uint32(v3877)+4)) = v3889
	goto L1065
L1069:
	;
	goto L1068
L1070:
	;
	v5051 = v3282
	goto L234
L1071:
	;
	v3906 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v3908 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	v3911 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L1
	} else {
		goto L1073
	}
L1073:
	;
	F_TupleDescInitBuiltinEntry(m, v3911, int32(1), int32(_a_F_PostgresMainLoopOnce_129), int32(25))
	mBase = m.M
	v3917 = m.ExcPending
	if v3917 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	F_TupleDescInitBuiltinEntry(m, v3911, int32(2), int32(_a_F_PostgresMainLoopOnce_130), int32(25))
	mBase = m.M
	v3922 = m.ExcPending
	if v3922 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1075:
	;
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v3906)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+352)) = v3923
	v3926 = v685 + int32(_a_F_PostgresMainLoopOnce_72)
	v3931 = F_pg_snprintf(m, v3926, int32(64), int32(_a_F_PostgresMainLoopOnce_131), v685+int32(352))
	mBase = m.M
	v3932 = m.ExcPending
	if v3932 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1076:
	;
	v3933 = *(*int32)(unsafe.Add(mBase, uint32(v3906)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+336)) = v3933
	v3936 = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	v3941 = F_pg_snprintf(m, v3936, int32(1024), int32(_a_F_PostgresMainLoopOnce_132), v685+int32(336))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v3908)+4))
	m.T0[v3944].(func(*base.Module, int32, int32, int32))(m, v3908, int32(1), v3911)
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	v3948 = v685 + int32(_a_F_PostgresMainLoopOnce_78)
	F_pq_beginmessage(m, v3948, int32(68))
	mBase = m.M
	v3951 = m.ExcPending
	if v3951 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	F_enlargeStringInfo(m, v3948, int32(2))
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1080:
	;
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v3958 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3955+v3956))) = uint16(v3958)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v3955 + int32(2)
	v3963 = F_strlen(m, v3926)
	mBase = m.M
	F_enlargeStringInfo(m, v3948, int32(4))
	mBase = m.M
	v3966 = m.ExcPending
	if v3966 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v3968 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v3972 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v3967+v3968))) = base.I32_rotr(v3963, int32(24))&v3972 | base.I32_rotr(v3963&v3972, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v3967 + int32(4)
	F_appendBinaryStringInfo(m, v3948, v3926, v3963)
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	v3986 = F_OpenTransientFile(m, v3936, int32(0))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	if v3986 < int32(0) {
		goto L243
	} else {
		goto L1084
	}
L1084:
	;
	v3990 = int64(0)
	v3992 = F___lseek(m, v3986, v3990, int32(2))
	mBase = m.M
	if v3992 < v3990 {
		goto L242
	} else {
		goto L1085
	}
L1085:
	;
	v3995 = int64(0)
	v3997 = F___lseek(m, v3986, v3995, int32(0))
	mBase = m.M
	if v3997 != v3995 {
		goto L241
	} else {
		goto L1086
	}
L1086:
	;
	F_enlargeStringInfo(m, v3948, int32(4))
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	v4003 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v4004 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v4006 = base.I32_wrap_i64(v3992)
	v4007 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v4003+v4004))) = base.I32_rotr(v4006&v4007, int32(8)) | base.I32_rotr(v4006, int32(24))&v4007
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v4003 + int32(4)
	if v3992 != int64(0) {
		goto L1088
	} else {
		goto L1089
	}
L1088:
	;
	v4054 = v3992
	goto L1091
L1089:
	;
	goto L1090
L1090:
	;
	v4120 = F_CloseTransientFile(m, v3986)
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1091:
	;
	v4059 = int32(_a_F_PostgresMainLoopOnce_133)
	v4060 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108]))
	*(*int32)(unsafe.Add(mBase, uint32(v4060))) = int32(167772227)
	v4064 = v685 + int32(464)
	v4066 = F_read(m, v3986, v4064, int32(_a_F_PostgresMainLoopOnce_20))
	mBase = m.M
	v4068 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108]))
	v4069 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4068))) = v4069
	if v4066 < v4069 {
		goto L240
	} else {
		goto L1093
	}
L1092:
	;
	goto L1090
L1093:
	;
	if v4066 == int32(0) {
		goto L239
	} else {
		goto L1094
	}
L1094:
	;
	F_appendBinaryStringInfo(m, v685+int32(_a_F_PostgresMainLoopOnce_78), v4064, v4066)
	mBase = m.M
	v4078 = m.ExcPending
	if v4078 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	v4080 = v4054 - base.I64_extend_i32_u(v4066)
	if int64(0) < v4080 {
		v4054 = v4080
		goto L1091
	} else {
		goto L1096
	}
L1096:
	;
	goto L1092
L1097:
	;
	if v4120 != 0 {
		goto L238
	} else {
		goto L1098
	}
L1098:
	;
	F_pq_endmessage(m, v685+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v4125 = m.ExcPending
	if v4125 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1099:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_67)
	goto L234
L1100:
	;
	v4133 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[109]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])) = v4133
	v4136 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4141 = F_AllocSetContextCreateInternal(m, v4136, int32(_a_F_PostgresMainLoopOnce_134), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	v4143 = int32(_a_F_PostgresMainLoopOnce_135)
	v4144 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4141
	v4148 = F_palloc0(m, int32(36))
	mBase = m.M
	v4149 = m.ExcPending
	if v4149 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4148))) = v4141
	F_initStringInfo(m, v4148+int32(4))
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L1
	} else {
		goto L1103
	}
L1103:
	;
	v4156 = F_MemoryContextAllocZero(m, v4141, int32(32))
	mBase = m.M
	v4157 = m.ExcPending
	if v4157 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4156)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4156)+24)) = v4141
	v4163 = F_MemoryContextAllocExtended(m, v4141, int32(_a_F_PostgresMainLoopOnce_136), int32(5))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L1105
	}
L1105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4156)+12)) = int64(63329292795903)
	*(*int64)(unsafe.Add(mBase, uint32(v4156))) = int64(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v4156)+20)) = v4163
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+24)) = v4156
	v4172 = F_palloc0(m, int32(24))
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4172)+20)) = int32(427)
	*(*int32)(unsafe.Add(mBase, uint32(v4172)+16)) = int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v4172)+12)) = int32(429)
	*(*int32)(unsafe.Add(mBase, uint32(v4172)+8)) = int32(430)
	*(*int32)(unsafe.Add(mBase, uint32(v4172)+4)) = int32(431)
	*(*int32)(unsafe.Add(mBase, uint32(v4172))) = v4148
	v4186 = F_palloc(m, int32(112))
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1107:
	;
	v4189 = F_palloc(m, int32(68))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L1
	} else {
		goto L1108
	}
L1108:
	;
	v4191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4189)+52)) = uint8(v4191)
	*(*int32)(unsafe.Add(mBase, uint32(v4189)+4)) = v4191
	*(*int32)(unsafe.Add(mBase, uint32(v4189))) = v4172
	if v4186 == v4191 {
		goto L1111
	} else {
		goto L1112
	}
L1109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+104)) = int32(_a_F_PostgresMainLoopOnce_137)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+100)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4186)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+88)) = int32(_a_F_PostgresMainLoopOnce_138)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+84)) = int32(_a_F_PostgresMainLoopOnce_139)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+80)) = int32(_a_F_PostgresMainLoopOnce_140)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+76)) = int32(_a_F_PostgresMainLoopOnce_141)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+72)) = int32(_a_F_PostgresMainLoopOnce_142)
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+68)) = v4189
	v4283 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v4284 = m.ExcPending
	if v4284 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+40)) = int32(1)
	v4216 = F_palloc0(m, int32(20))
	mBase = m.M
	v4217 = m.ExcPending
	if v4217 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1111:
	;
	v4199 = F_palloc0(m, int32(68))
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1112:
	;
	goto L1113
L1113:
	;
	base.MemoryFill(m, v4186, int32(0), int32(68))
	v4210 = v4186
	goto L1110
L1114:
	;
	if v4199 == int32(0) {
		goto L1109
	} else {
		goto L1115
	}
L1115:
	;
	v4203 = *(*int32)(unsafe.Add(mBase, uint32(v4199)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4199)+36)) = v4203 | int32(1)
	v4210 = v4199
	goto L1110
L1116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+52)) = v4216
	v4220 = F_palloc0(m, int32(28))
	mBase = m.M
	v4221 = m.ExcPending
	if v4221 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	v4223 = F_palloc(m, int32(640))
	mBase = m.M
	v4224 = m.ExcPending
	if v4224 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	v4226 = F_palloc(m, int32(256))
	mBase = m.M
	v4227 = m.ExcPending
	if v4227 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1119:
	;
	v4229 = F_palloc(m, int32(64))
	mBase = m.M
	v4230 = m.ExcPending
	if v4230 != 0 {
		goto L1
	} else {
		goto L1120
	}
L1120:
	;
	v4231 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+52))
	F_initStringInfo(m, v4231+int32(4))
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		goto L1
	} else {
		goto L1121
	}
L1121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+48)) = v4220
	*(*int32)(unsafe.Add(mBase, uint32(v4220))) = int32(64)
	v4239 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4239)+4)) = v4223
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4241)+12)) = v4226
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4243)+16)) = v4229
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+48))
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4245)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v4246))) = int32(0)
	v4249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4210)+56)) = uint8(v4249)
	*(*uint8)(unsafe.Add(mBase, uint32(v4210)+24)) = uint8(v4249)
	v4253 = F_makeStringInfo(m)
	mBase = m.M
	v4254 = m.ExcPending
	if v4254 != 0 {
		goto L1
	} else {
		goto L1122
	}
L1122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+60)) = v4253
	v4256 = *(*int32)(unsafe.Add(mBase, uint32(v4210)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4210)+36)) = v4256 | int32(2)
	goto L1109
L1123:
	;
	if v4283 == int32(0) {
		goto L1124
	} else {
		goto L1125
	}
L1124:
	;
	v4289 = *(*int32)(unsafe.Add(mBase, uint32(v4172)+20))
	m.T0[v4289].(func(*base.Module, int32, int32, int32))(m, v4172, int32(_a_F_PostgresMainLoopOnce_120), int32(0))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1125:
	;
	goto L1126
L1126:
	;
	v4292 = F_pg_cryptohash_init(m, v4283)
	mBase = m.M
	if v4292 < int32(0) {
		goto L1128
	} else {
		goto L1129
	}
L1127:
	;
	goto L1126
L1128:
	;
	v4297 = *(*int32)(unsafe.Add(mBase, uint32(v4172)+20))
	m.T0[v4297].(func(*base.Module, int32, int32, int32))(m, v4172, int32(_a_F_PostgresMainLoopOnce_143), int32(0))
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1129:
	;
	goto L1130
L1130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4186)+108)) = v4283
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+32)) = v4186
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4144
	v4305 = v685 + int32(464)
	F_pq_beginmessage(m, v4305, int32(71))
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1131:
	;
	goto L1130
L1132:
	;
	F_enlargeStringInfo(m, v4305, int32(1))
	mBase = m.M
	v4311 = m.ExcPending
	if v4311 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	v4312 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4312+v4313))) = uint8(v4315)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v4312 + int32(1)
	F_enlargeStringInfo(m, v4305, int32(2))
	mBase = m.M
	v4322 = m.ExcPending
	if v4322 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	v4323 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4326 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4323+v4324))) = uint16(v4326)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v4323 + int32(2)
	F_pq_endmessage_reuse(m, v4305)
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1135:
	;
	v4334 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103]))
	v4335 = *(*int32)(unsafe.Add(mBase, uint32(v4334)+4))
	v4336 = m.T0[v4335].(func(*base.Module) int32)(m)
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1136:
	;
	goto L1138
L1137:
	;
	v4471 = int32(_a_F_PostgresMainLoopOnce_135)
	v4472 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4474 = *(*int32)(unsafe.Add(mBase, uint32(v4148)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4474
	v4476 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+32))
	v4477 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+4))
	v4478 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+8))
	F_json_parse_manifest_incremental_chunk(m, v4476, v4477, v4478, int32(1))
	mBase = m.M
	v4481 = m.ExcPending
	if v4481 != 0 {
		goto L1
	} else {
		goto L1162
	}
L1138:
	;
	v4375 = int32(_a_F_PostgresMainLoopOnce_5)
	v4377 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4377 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v4382 = m.ExcPending
	if v4382 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4452 = m.ExcPending
	if v4452 != 0 {
		goto L1
	} else {
		goto L1157
	}
L1140:
	;
	v4383 = F_pq_getbyte(m)
	mBase = m.M
	v4384 = m.ExcPending
	if v4384 != 0 {
		goto L1
	} else {
		goto L1141
	}
L1141:
	;
	v4386 = v4383 - int32(72)
	if base.Ui32(int32(30)) < base.Ui32(v4386) {
		goto L236
	} else {
		goto L1142
	}
L1142:
	;
	if int32(1)<<(uint(v4386)%32)&int32(1207961601) == int32(0) {
		goto L1144
	} else {
		goto L1145
	}
L1143:
	;
	v4402 = F_pq_getmessage(m, v685+int32(464), v4399)
	mBase = m.M
	v4403 = m.ExcPending
	if v4403 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1144:
	;
	if v4386 != int32(28) {
		goto L236
	} else {
		goto L1147
	}
L1145:
	;
	goto L1146
L1146:
	;
	v4399 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L1143
L1147:
	;
	v4399 = int32(1073741822)
	goto L1143
L1148:
	;
	if v4402 != 0 {
		goto L237
	} else {
		goto L1149
	}
L1149:
	;
	v4404 = int32(_a_F_PostgresMainLoopOnce_5)
	v4406 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4406 - int32(1)
	switch v4386 {
	case 0, 11:
		goto L1138
	default:
		goto L1137
	case 28:
		goto L1151
	case 30:
		goto L1150
	}
L1150:
	;
	goto L1139
L1151:
	;
	v4410 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4411 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4412 = int32(_a_F_PostgresMainLoopOnce_135)
	v4413 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4415 = *(*int32)(unsafe.Add(mBase, uint32(v4148)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4415
	v4417 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+8))
	if base.B2i32(v4417 < int32(1025))|base.B2i32(v4417+v4411 < int32(_a_F_PostgresMainLoopOnce_144)) == int32(0) {
		goto L1152
	} else {
		goto L1153
	}
L1152:
	;
	v4426 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+32))
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+4))
	F_json_parse_manifest_incremental_chunk(m, v4426, v4427, v4417-int32(1024), int32(0))
	mBase = m.M
	v4432 = m.ExcPending
	if v4432 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1153:
	;
	goto L1154
L1154:
	;
	F_appendBinaryStringInfo(m, v4148+int32(4), v4410, v4411)
	mBase = m.M
	v4446 = m.ExcPending
	if v4446 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1155:
	;
	v4433 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+4))
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+8))
	v4436 = int32(1024)
	base.MemoryCopy(m, v4433, v4433+v4434-v4436, int32(1025))
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+8)) = v4436
	goto L1154
L1156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4413
	goto L1138
L1157:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1158:
	;
	v4458 = F_pq_getmsgstring(m, v685+int32(464))
	mBase = m.M
	v4459 = m.ExcPending
	if v4459 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+384)) = v4458
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_145), v685+int32(384))
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1160:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(794), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1162:
	;
	v4482 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+4))
	F_pfree(m, v4482)
	mBase = m.M
	v4484 = m.ExcPending
	if v4484 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4148)+4)) = int32(0)
	v4487 = *(*int32)(unsafe.Add(mBase, uint32(v4148)+32))
	v4488 = *(*int32)(unsafe.Add(mBase, uint32(v4487)+68))
	F_pfree(m, v4488)
	mBase = m.M
	v4490 = m.ExcPending
	if v4490 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	F_freeJsonLexContext(m, v4487)
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	F_pfree(m, v4487)
	mBase = m.M
	v4494 = m.ExcPending
	if v4494 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4472
	v4498 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[110]))
	if v4498 != 0 {
		goto L1167
	} else {
		goto L1168
	}
L1167:
	;
	F_MemoryContextDelete(m, v4498)
	mBase = m.M
	v4500 = m.ExcPending
	if v4500 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1168:
	;
	goto L1169
L1169:
	;
	v4502 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[111]))
	v4506 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+16))
	if v4506 != v4502 {
		goto L1172
	} else {
		goto L1173
	}
L1170:
	;
	goto L1169
L1171:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[110])) = v4141
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[112])) = v4148
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v4541 = m.ExcPending
	if v4541 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1172:
	;
	if v4506 == int32(0) {
		goto L1175
	} else {
		goto L1176
	}
L1173:
	;
	goto L1174
L1174:
	;
	goto L1171
L1175:
	;
	if v4502 != 0 {
		goto L1182
	} else {
		goto L1183
	}
L1176:
	;
	v4510 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+28))
	v4511 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+24))
	if v4511 != 0 {
		goto L1178
	} else {
		goto L1179
	}
L1177:
	;
	if v4510 == int32(0) {
		goto L1175
	} else {
		goto L1181
	}
L1178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4511)+28)) = v4510
	goto L1177
L1179:
	;
	goto L1180
L1180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4506)+20)) = v4510
	goto L1177
L1181:
	;
	v4516 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+24)) = v4516
	goto L1175
L1182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+16)) = v4502
	v4523 = *(*int32)(unsafe.Add(mBase, uint32(v4502)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+28)) = v4523
	if v4523 != 0 {
		goto L1185
	} else {
		goto L1186
	}
L1183:
	;
	goto L1184
L1184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4141)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+16)) = int32(0)
	goto L1174
L1185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+24)) = v4141
	goto L1187
L1186:
	;
	goto L1187
L1187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4502)+20)) = v4141
	goto L1171
L1188:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_66)
	goto L234
L1189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4548 = m.ExcPending
	if v4548 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1190:
	;
	v4549 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v4550 = *(*int32)(unsafe.Add(mBase, uint32(v4549)))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v4550
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_147), v685)
	mBase = m.M
	v4554 = m.ExcPending
	if v4554 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2215), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4559 = m.ExcPending
	if v4559 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1193:
	;
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v4567 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[112]))
	F_SendBaseBackup(m, v4565, v4567)
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	v5051 = v4560
	goto L234
L1195:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L1
	} else {
		goto L1196
	}
L1196:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_148), int32(0))
	mBase = m.M
	v4580 = m.ExcPending
	if v4580 != 0 {
		goto L1
	} else {
		goto L1197
	}
L1197:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2010), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1199:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+416)) = v1819
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_149), v685+int32(416))
	mBase = m.M
	v4598 = m.ExcPending
	if v4598 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2078), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4603 = m.ExcPending
	if v4603 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1203:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v4610 = m.ExcPending
	if v4610 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v4614 = m.ExcPending
	if v4614 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2104), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4619 = m.ExcPending
	if v4619 != 0 {
		goto L1
	} else {
		goto L1206
	}
L1206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1207:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+64)) = int32(_a_F_PostgresMainLoopOnce_85)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_151), v685-int32(-64))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1209:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(520), int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v4638 = m.ExcPending
	if v4638 != 0 {
		goto L1
	} else {
		goto L1210
	}
L1210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1211:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4645 = m.ExcPending
	if v4645 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1212:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4649 = m.ExcPending
	if v4649 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1137), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1215:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4661 = m.ExcPending
	if v4661 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1216:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1217:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1159), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4670 = m.ExcPending
	if v4670 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1219:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4677 = m.ExcPending
	if v4677 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1220:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1169), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4686 = m.ExcPending
	if v4686 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1223:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4693 = m.ExcPending
	if v4693 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1178), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		goto L1
	} else {
		goto L1226
	}
L1226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1227:
	;
	v4707 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+208)) = v4707
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_153), v685+int32(208))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1228:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1183), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4718 = m.ExcPending
	if v4718 != 0 {
		goto L1
	} else {
		goto L1229
	}
L1229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+176)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_155), v685+int32(176))
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
		goto L1
	} else {
		goto L1231
	}
L1231:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1268), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4734 = m.ExcPending
	if v4734 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+160)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_156), v685+int32(160))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1234:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1274), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4750 = m.ExcPending
	if v4750 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+144)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_157), v685+int32(144))
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1279), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4766 = m.ExcPending
	if v4766 != 0 {
		goto L1
	} else {
		goto L1238
	}
L1238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+112)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_158), v685+int32(112))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1285), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+128)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_159), v685+int32(128))
	mBase = m.M
	v4793 = m.ExcPending
	if v4793 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1243:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1291), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4798 = m.ExcPending
	if v4798 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1245:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4805 = m.ExcPending
	if v4805 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4809 = m.ExcPending
	if v4809 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1420), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4814 = m.ExcPending
	if v4814 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4821 = m.ExcPending
	if v4821 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4825 = m.ExcPending
	if v4825 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1429), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4830 = m.ExcPending
	if v4830 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1253:
	;
	v4835 = *(*int32)(unsafe.Add(mBase, uint32(v2941)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+224)) = v4835
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_153), v685+int32(224))
	mBase = m.M
	v4841 = m.ExcPending
	if v4841 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1434), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4846 = m.ExcPending
	if v4846 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1257:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4855 = m.ExcPending
	if v4855 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+240)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_161), v685+int32(240))
	mBase = m.M
	v4863 = m.ExcPending
	if v4863 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1259:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(616), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4868 = m.ExcPending
	if v4868 != 0 {
		goto L1
	} else {
		goto L1260
	}
L1260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1261:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4874 = m.ExcPending
	if v4874 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+256)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_163), v685+int32(256))
	mBase = m.M
	v4882 = m.ExcPending
	if v4882 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(623), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4887 = m.ExcPending
	if v4887 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1265:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4893 = m.ExcPending
	if v4893 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+320)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_164), v685+int32(320))
	mBase = m.M
	v4901 = m.ExcPending
	if v4901 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(627), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4906 = m.ExcPending
	if v4906 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1269:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4912 = m.ExcPending
	if v4912 != 0 {
		goto L1
	} else {
		goto L1270
	}
L1270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+288)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_165), v685+int32(288))
	mBase = m.M
	v4920 = m.ExcPending
	if v4920 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(644), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4925 = m.ExcPending
	if v4925 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1273:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4932 = m.ExcPending
	if v4932 != 0 {
		goto L1
	} else {
		goto L1274
	}
L1274:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+312)) = uint32(v4054)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+308)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+304)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_166), v685+int32(304))
	mBase = m.M
	v4943 = m.ExcPending
	if v4943 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1275:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(649), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4948 = m.ExcPending
	if v4948 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1277:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4954 = m.ExcPending
	if v4954 != 0 {
		goto L1
	} else {
		goto L1278
	}
L1278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+272)) = v685 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_167), v685+int32(272))
	mBase = m.M
	v4962 = m.ExcPending
	if v4962 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1279:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(658), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4967 = m.ExcPending
	if v4967 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1281:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v4974 = m.ExcPending
	if v4974 != 0 {
		goto L1
	} else {
		goto L1282
	}
L1282:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v4978 = m.ExcPending
	if v4978 != 0 {
		goto L1
	} else {
		goto L1283
	}
L1283:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(772), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4983 = m.ExcPending
	if v4983 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4990 = m.ExcPending
	if v4990 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1286:
	;
	goto L1287
L1287:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5006 = m.ExcPending
	if v5006 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1288:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v4993 = m.ExcPending
	if v4993 != 0 {
		goto L1
	} else {
		goto L1289
	}
L1289:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v4997 = m.ExcPending
	if v4997 != 0 {
		goto L1
	} else {
		goto L1290
	}
L1290:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(746), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v5002 = m.ExcPending
	if v5002 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1292:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L1
	} else {
		goto L1293
	}
L1293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+368)) = v4383
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_168), v685+int32(368))
	mBase = m.M
	v5015 = m.ExcPending
	if v5015 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1294:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(763), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v5020 = m.ExcPending
	if v5020 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1296:
	;
	v5024 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	F_StartTransactionCommand(m)
	mBase = m.M
	v5027 = m.ExcPending
	if v5027 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	v5028 = *(*int32)(unsafe.Add(mBase, uint32(v5024)+4))
	F_GetPGVariable(m, v5028, v5022)
	mBase = m.M
	v5030 = m.ExcPending
	if v5030 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1298:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5032 = m.ExcPending
	if v5032 != 0 {
		goto L1
	} else {
		goto L1299
	}
L1299:
	;
	v5051 = int32(_a_F_PostgresMainLoopOnce_169)
	goto L234
L1300:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v690
	v5075 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	F_MemoryContextReset(m, v5075)
	mBase = m.M
	v5077 = m.ExcPending
	if v5077 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1301:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L233
L1302:
	;
	goto L225
L1303:
	;
	v5173 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v5183 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1306
L1304:
	;
	goto L1305
L1305:
	;
	F_start_xact_command(m)
	mBase = m.M
	v5187 = m.ExcPending
	if v5187 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1306:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1305
L1307:
	;
	v5189 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117]))
	if v5189 != 0 {
		goto L1308
	} else {
		goto L1309
	}
L1308:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int32(0)
	F_DropCachedPlan(m, v5189)
	mBase = m.M
	v5194 = m.ExcPending
	if v5194 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1309:
	;
	goto L1310
L1310:
	;
	v5195 = int32(_a_F_PostgresMainLoopOnce_135)
	v5196 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v5199 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5199
	v5202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v5202 == int32(1) {
		goto L1312
	} else {
		goto L1313
	}
L1311:
	;
	goto L1310
L1312:
	;
	v5205 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v5215 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1315
L1313:
	;
	goto L1314
L1314:
	;
	v5219 = F_raw_parser(m, v675, int32(0))
	mBase = m.M
	v5220 = m.ExcPending
	if v5220 != 0 {
		goto L1
	} else {
		goto L1316
	}
L1315:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1314
L1316:
	;
	v5222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v5222 == int32(1) {
		goto L1317
	} else {
		goto L1318
	}
L1317:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_172))
	mBase = m.M
	v5227 = m.ExcPending
	if v5227 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1318:
	;
	goto L1319
L1319:
	;
	v5229 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119]))
	switch v5229 {
	case 0:
		v5520 = v1
		goto L1325
	default:
		goto L1330
	case 3:
		goto L1329
	}
L1320:
	;
	goto L1319
L1321:
	;
	v6126 = F_check_log_duration(m, v5161+int32(80), v6115)
	mBase = m.M
	v6127 = m.ExcPending
	if v6127 != 0 {
		goto L1
	} else {
		goto L1503
	}
L1322:
	;
	v6071 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L1492
L1323:
	;
	v6017 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L1482
L1324:
	;
	v5570 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if v5570 <= int32(0) {
		goto L1322
	} else {
		goto L1357
	}
L1325:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5196
	if v5219 == int32(0) {
		v6006 = v5520
		goto L1323
	} else {
		goto L1356
	}
L1326:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1179), int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v5490 = m.ExcPending
	if v5490 != 0 {
		goto L1
	} else {
		goto L1355
	}
L1327:
	;
	v5441 = *(*int32)(unsafe.Add(mBase, uint32(v5431)+64))
	v5442 = *(*int32)(unsafe.Add(mBase, uint32(v5441)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5161)+48)) = v5442
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_174), v5161+int32(48))
	mBase = m.M
	v5448 = m.ExcPending
	if v5448 != 0 {
		goto L1
	} else {
		goto L1354
	}
L1328:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5196
	v6006 = v1
	goto L1323
L1329:
	;
	v5362 = int32(1)
	v5365 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5366 = m.ExcPending
	if v5366 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1330:
	;
	if v5219 == int32(0) {
		goto L1328
	} else {
		goto L1331
	}
L1331:
	;
	v5232 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if int32(0) < v5232 {
		goto L1332
	} else {
		goto L1333
	}
L1332:
	;
	v5235 = int32(0)
	goto L1335
L1333:
	;
	v5307 = v5232
	goto L1334
L1334:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5196
	v5554 = v5307
	v5562 = v1
	goto L1324
L1335:
	;
	v5272 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v5272+v5235<<(uint(int32(2))%32))))
	v5277 = F_GetCommandLogLevel(m, v5276)
	mBase = m.M
	v5278 = m.ExcPending
	if v5278 != 0 {
		goto L1
	} else {
		goto L1337
	}
L1336:
	;
	v5307 = v5284
	goto L1334
L1337:
	;
	v5280 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119]))
	if base.Ui32(v5277) <= base.Ui32(v5280) {
		goto L1329
	} else {
		goto L1338
	}
L1338:
	;
	v5283 = v5235 + int32(1)
	v5284 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if v5283 < v5284 {
		v5235 = v5283
		goto L1335
	} else {
		goto L1339
	}
L1339:
	;
	goto L1336
L1340:
	;
	if v5365 == int32(0) {
		v5520 = v5362
		goto L1325
	} else {
		goto L1341
	}
L1341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5161)+64)) = v675
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_175), v5161-int32(-64))
	mBase = m.M
	v5374 = m.ExcPending
	if v5374 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1342:
	;
	F_errhidestmt(m)
	mBase = m.M
	v5376 = m.ExcPending
	if v5376 != 0 {
		goto L1
	} else {
		goto L1343
	}
L1343:
	;
	if v5219 == int32(0) {
		goto L1326
	} else {
		goto L1344
	}
L1344:
	;
	v5379 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if v5379 <= int32(0) {
		goto L1326
	} else {
		goto L1345
	}
L1345:
	;
	v5383 = int32(0)
	v5392 = v5379
	goto L1346
L1346:
	;
	v5420 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v5424 = *(*int32)(unsafe.Add(mBase, uint32(v5420+v5383<<(uint(int32(2))%32))))
	v5425 = *(*int32)(unsafe.Add(mBase, uint32(v5424)+4))
	v5426 = *(*int32)(unsafe.Add(mBase, uint32(v5425)))
	if v5426 == int32(253) {
		goto L1348
	} else {
		goto L1349
	}
L1347:
	;
	goto L1326
L1348:
	;
	v5429 = *(*int32)(unsafe.Add(mBase, uint32(v5425)+4))
	v5431 = F_FetchPreparedStatement(m, v5429, int32(0))
	mBase = m.M
	v5432 = m.ExcPending
	if v5432 != 0 {
		goto L1
	} else {
		goto L1351
	}
L1349:
	;
	v5435 = v5392
	goto L1350
L1350:
	;
	v5437 = v5383 + int32(1)
	if v5437 < v5435 {
		v5383 = v5437
		v5392 = v5435
		goto L1346
	} else {
		goto L1353
	}
L1351:
	;
	if v5431 != 0 {
		goto L1327
	} else {
		goto L1352
	}
L1352:
	;
	v5433 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	v5435 = v5433
	goto L1350
L1353:
	;
	goto L1347
L1354:
	;
	goto L1326
L1355:
	;
	v5520 = v5362
	goto L1325
L1356:
	;
	v5532 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	v5554 = v5532
	v5562 = v5520
	goto L1324
L1357:
	;
	v5581 = v1
	goto L1358
L1358:
	;
	v5610 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v5613 = v5610 + v5581<<(uint(int32(2))%32)
	v5614 = *(*int32)(unsafe.Add(mBase, uint32(v5613)))
	v5619 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v5619 == int32(0) {
		goto L1361
	} else {
		goto L1362
	}
L1359:
	;
	goto L1322
L1360:
	;
	v5659 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v5659 == int32(0) {
		goto L1366
	} else {
		goto L1367
	}
L1361:
	;
	goto L1360
L1362:
	;
	v5623 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v5623&int32(1) == int32(0) {
		goto L1361
	} else {
		goto L1363
	}
L1363:
	;
	v5630 = *(*int64)(unsafe.Add(mBase, uint32(v5619)+392))
	if int32(0)&base.B2i32(v5630 != int64(0)) != 0 {
		goto L1361
	} else {
		goto L1364
	}
L1364:
	;
	v5634 = int32(_a_F_PostgresMainLoopOnce_176)
	v5636 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v5637 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v5636 + v5637
	v5640 = *(*int32)(unsafe.Add(mBase, uint32(v5619)))
	*(*int32)(unsafe.Add(mBase, uint32(v5619))) = v5640 + v5637
	*(*int64)(unsafe.Add(mBase, uint32(v5619)+392)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5619))) = v5640 + int32(2)
	v5651 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v5651 - v5637
	goto L1361
L1365:
	;
	v5695 = *(*int32)(unsafe.Add(mBase, uint32(v5614)+4))
	v5696 = F_CreateCommandTag(m, v5695)
	mBase = m.M
	v5697 = m.ExcPending
	if v5697 != 0 {
		goto L1
	} else {
		goto L1370
	}
L1366:
	;
	goto L1365
L1367:
	;
	v5663 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v5663&int32(1) == int32(0) {
		goto L1366
	} else {
		goto L1368
	}
L1368:
	;
	v5670 = *(*int64)(unsafe.Add(mBase, uint32(v5659)+400))
	if int32(0)&base.B2i32(v5670 != int64(0)) != 0 {
		goto L1366
	} else {
		goto L1369
	}
L1369:
	;
	v5674 = int32(_a_F_PostgresMainLoopOnce_176)
	v5676 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v5677 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v5676 + v5677
	v5680 = *(*int32)(unsafe.Add(mBase, uint32(v5659)))
	*(*int32)(unsafe.Add(mBase, uint32(v5659))) = v5680 + v5677
	*(*int64)(unsafe.Add(mBase, uint32(v5659)+400)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5659))) = v5680 + int32(2)
	v5691 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v5691 - v5677
	goto L1366
L1370:
	;
	v5702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5696<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[124]))))
	*(*int32)(unsafe.Add(mBase, uint32(v5161+int32(72)))) = v5702
	goto L1371
L1371:
	;
	v5707 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5708 = *(*int32)(unsafe.Add(mBase, uint32(v5707)+24))
	goto L1373
L1372:
	;
	F_start_xact_command(m)
	mBase = m.M
	v5750 = m.ExcPending
	if v5750 != 0 {
		goto L1
	} else {
		goto L1384
	}
L1373:
	;
	if base.B2i32((v5708-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L1372
	} else {
		goto L1374
	}
L1374:
	;
	v5717 = *(*int32)(unsafe.Add(mBase, uint32(v5614)+4))
	if v5717 == int32(0) {
		goto L1375
	} else {
		goto L1376
	}
L1375:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5733 = m.ExcPending
	if v5733 != 0 {
		goto L1
	} else {
		goto L1379
	}
L1376:
	;
	v5720 = *(*int32)(unsafe.Add(mBase, uint32(v5717)))
	if v5720 != int32(225) {
		goto L1375
	} else {
		goto L1377
	}
L1377:
	;
	v5723 = *(*int32)(unsafe.Add(mBase, uint32(v5717)+4))
	if (v5723-int32(2))&int32(-6) == int32(0) {
		goto L1372
	} else {
		goto L1378
	}
L1378:
	;
	goto L1375
L1379:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v5736 = m.ExcPending
	if v5736 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1380:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v5740 = m.ExcPending
	if v5740 != 0 {
		goto L1
	} else {
		goto L1381
	}
L1381:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v5742 = m.ExcPending
	if v5742 != 0 {
		goto L1
	} else {
		goto L1382
	}
L1382:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1246), int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v5747 = m.ExcPending
	if v5747 != 0 {
		goto L1
	} else {
		goto L1383
	}
L1383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1384:
	;
	v5752 = base.B2i32(v5554 < int32(2))
	if v5752 == int32(0) {
		goto L1385
	} else {
		goto L1386
	}
L1385:
	;
	v5757 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5758 = *(*int32)(unsafe.Add(mBase, uint32(v5757)+24))
	if v5758 == int32(1) {
		goto L1389
	} else {
		goto L1390
	}
L1386:
	;
	goto L1387
L1387:
	;
	v5764 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v5764 != 0 {
		goto L1392
	} else {
		goto L1393
	}
L1388:
	;
	goto L1387
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5757)+24)) = int32(4)
	goto L1391
L1390:
	;
	goto L1391
L1391:
	;
	goto L1388
L1392:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5766 = m.ExcPending
	if v5766 != 0 {
		goto L1
	} else {
		goto L1395
	}
L1393:
	;
	goto L1394
L1394:
	;
	v5769 = *(*int32)(unsafe.Add(mBase, uint32(v5614)+4))
	v5770 = *(*int32)(unsafe.Add(mBase, uint32(v5769)))
	switch v5770 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v5774 = int32(1)
		goto L1397
	default:
		goto L1398
	}
L1395:
	;
	goto L1394
L1396:
	;
	if v5774 != 0 {
		goto L1399
	} else {
		goto L1400
	}
L1397:
	;
	goto L1396
L1398:
	;
	v5774 = int32(0)
	goto L1397
L1399:
	;
	v5775 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v5776 = m.ExcPending
	if v5776 != 0 {
		goto L1
	} else {
		goto L1402
	}
L1400:
	;
	goto L1401
L1401:
	;
	v5781 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v5783 = v5613 + int32(4)
	v5784 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v5785 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if base.Ui32(v5783) < base.Ui32(v5784+v5785<<(uint(int32(2))%32)) {
		goto L1404
	} else {
		goto L1405
	}
L1402:
	;
	F_PushActiveSnapshot(m, v5775)
	mBase = m.M
	v5778 = m.ExcPending
	if v5778 != 0 {
		goto L1
	} else {
		goto L1403
	}
L1403:
	;
	goto L1401
L1404:
	;
	v5794 = F_AllocSetContextCreateInternal(m, v5781, int32(_a_F_PostgresMainLoopOnce_177), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v5795 = m.ExcPending
	if v5795 != 0 {
		goto L1
	} else {
		goto L1407
	}
L1405:
	;
	v5796 = v5781
	v5797 = int32(0)
	goto L1406
L1406:
	;
	v5798 = int32(_a_F_PostgresMainLoopOnce_135)
	v5799 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5796
	v5803 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v5803 == int32(1) {
		goto L1408
	} else {
		goto L1409
	}
L1407:
	;
	v5796 = v5794
	v5797 = v5794
	goto L1406
L1408:
	;
	v5806 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v5816 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1411
L1409:
	;
	goto L1410
L1410:
	;
	v5819 = int32(0)
	v5822 = F_parse_analyze_fixedparams(m, v5614, v675, v5819, v5819, v5819)
	mBase = m.M
	v5823 = m.ExcPending
	if v5823 != 0 {
		goto L1
	} else {
		goto L1412
	}
L1411:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1410
L1412:
	;
	v5825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v5825 == int32(1) {
		goto L1413
	} else {
		goto L1414
	}
L1413:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_178))
	mBase = m.M
	v5830 = m.ExcPending
	if v5830 != 0 {
		goto L1
	} else {
		goto L1416
	}
L1414:
	;
	goto L1415
L1415:
	;
	v5831 = F_pg_rewrite_query(m, v5822)
	mBase = m.M
	v5832 = m.ExcPending
	if v5832 != 0 {
		goto L1
	} else {
		goto L1417
	}
L1416:
	;
	goto L1415
L1417:
	;
	v5835 = F_pg_plan_queries(m, v5831, v675, int32(2048), int32(0))
	mBase = m.M
	v5836 = m.ExcPending
	if v5836 != 0 {
		goto L1
	} else {
		goto L1418
	}
L1418:
	;
	if v5774 != 0 {
		goto L1419
	} else {
		goto L1420
	}
L1419:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v5838 = m.ExcPending
	if v5838 != 0 {
		goto L1
	} else {
		goto L1422
	}
L1420:
	;
	goto L1421
L1421:
	;
	v5840 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v5840 != 0 {
		goto L1423
	} else {
		goto L1424
	}
L1422:
	;
	goto L1421
L1423:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5842 = m.ExcPending
	if v5842 != 0 {
		goto L1
	} else {
		goto L1426
	}
L1424:
	;
	goto L1425
L1425:
	;
	v5844 = int32(1)
	v5846 = F_CreatePortal(m, int32(_a_F_PostgresMainLoopOnce_179), v5844, v5844)
	mBase = m.M
	v5847 = m.ExcPending
	if v5847 != 0 {
		goto L1
	} else {
		goto L1427
	}
L1426:
	;
	goto L1425
L1427:
	;
	v5848 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5846)+136)) = uint8(v5848)
	*(*int64)(unsafe.Add(mBase, uint32(v5846)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+40)) = v5696
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+32)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+4)) = v5848
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+60)) = v5848
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+56)) = v5835
	*(*int32)(unsafe.Add(mBase, uint32(v5846)+36)) = v5696
	goto L1428
L1428:
	;
	v5862 = int32(0)
	F_PortalStart(m, v5846, v5862, v5862, v5862)
	mBase = m.M
	v5866 = m.ExcPending
	if v5866 != 0 {
		goto L1
	} else {
		goto L1429
	}
L1429:
	;
	v5867 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5161)+78)) = uint16(v5867)
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(v5614)+4))
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v5869)))
	if v5870 != int32(203) {
		goto L1430
	} else {
		goto L1431
	}
L1430:
	;
	F_PortalSetResultFormat(m, v5846, int32(1), v5161+int32(78))
	mBase = m.M
	v5891 = m.ExcPending
	if v5891 != 0 {
		goto L1
	} else {
		goto L1436
	}
L1431:
	;
	v5873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5869)+16)))
	if v5873 != 0 {
		goto L1430
	} else {
		goto L1432
	}
L1432:
	;
	v5874 = *(*int32)(unsafe.Add(mBase, uint32(v5869)+12))
	v5875 = F_GetPortalByName(m, v5874)
	mBase = m.M
	v5876 = m.ExcPending
	if v5876 != 0 {
		goto L1
	} else {
		goto L1433
	}
L1433:
	;
	if v5875 == int32(0) {
		goto L1430
	} else {
		goto L1434
	}
L1434:
	;
	v5879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5875)+76)))
	if v5879&int32(1) == int32(0) {
		goto L1430
	} else {
		goto L1435
	}
L1435:
	;
	v5884 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5161)+78)) = uint16(v5884)
	goto L1430
L1436:
	;
	v5892 = F_CreateDestReceiver(m, v5166)
	mBase = m.M
	v5893 = m.ExcPending
	if v5893 != 0 {
		goto L1
	} else {
		goto L1437
	}
L1437:
	;
	if v5166 == int32(2) {
		goto L1438
	} else {
		goto L1439
	}
L1438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5892)+20)) = v5846
	goto L1440
L1439:
	;
	goto L1440
L1440:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5799
	v5903 = F_PortalRun(m, v5846, int32(2147483647), int32(1), v5892, v5892, v5161+int32(80))
	mBase = m.M
	v5904 = m.ExcPending
	if v5904 != 0 {
		goto L1
	} else {
		goto L1441
	}
L1441:
	;
	v5905 = *(*int32)(unsafe.Add(mBase, uint32(v5892)+12))
	m.T0[v5905].(func(*base.Module, int32))(m, v5892)
	mBase = m.M
	v5907 = m.ExcPending
	if v5907 != 0 {
		goto L1
	} else {
		goto L1442
	}
L1442:
	;
	F_PortalDrop(m, v5846, int32(0))
	mBase = m.M
	v5910 = m.ExcPending
	if v5910 != 0 {
		goto L1
	} else {
		goto L1443
	}
L1443:
	;
	v5911 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v5912 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if base.Ui32(v5911+v5912<<(uint(int32(2))%32)) <= base.Ui32(v5783) {
		goto L1446
	} else {
		goto L1447
	}
L1444:
	;
	F_EndCommand(m, v5161+int32(80), v5166)
	mBase = m.M
	v5970 = m.ExcPending
	if v5970 != 0 {
		goto L1
	} else {
		goto L1476
	}
L1445:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5963 = m.ExcPending
	if v5963 != 0 {
		goto L1
	} else {
		goto L1475
	}
L1446:
	;
	if v5752 == int32(0) {
		goto L1449
	} else {
		goto L1450
	}
L1447:
	;
	goto L1448
L1448:
	;
	v5938 = *(*int32)(unsafe.Add(mBase, uint32(v5614)+4))
	v5939 = *(*int32)(unsafe.Add(mBase, uint32(v5938)))
	if v5939 == int32(225) {
		goto L1462
	} else {
		goto L1463
	}
L1449:
	;
	v5921 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5922 = *(*int32)(unsafe.Add(mBase, uint32(v5921)+24))
	if v5922 == int32(4) {
		goto L1453
	} else {
		goto L1454
	}
L1450:
	;
	goto L1451
L1451:
	;
	v5930 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L1456
L1452:
	;
	goto L1451
L1453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5921)+24)) = int32(1)
	goto L1455
L1454:
	;
	goto L1455
L1455:
	;
	goto L1452
L1456:
	;
	if v5930 != 0 {
		goto L1457
	} else {
		goto L1458
	}
L1457:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v5933 = m.ExcPending
	if v5933 != 0 {
		goto L1
	} else {
		goto L1460
	}
L1458:
	;
	goto L1459
L1459:
	;
	v5935 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v5935 == int32(0) {
		goto L1444
	} else {
		goto L1461
	}
L1460:
	;
	goto L1459
L1461:
	;
	goto L1445
L1462:
	;
	v5945 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L1465
L1463:
	;
	goto L1464
L1464:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v5952 = m.ExcPending
	if v5952 != 0 {
		goto L1
	} else {
		goto L1471
	}
L1465:
	;
	if v5945 != 0 {
		goto L1466
	} else {
		goto L1467
	}
L1466:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v5948 = m.ExcPending
	if v5948 != 0 {
		goto L1
	} else {
		goto L1469
	}
L1467:
	;
	goto L1468
L1468:
	;
	v5950 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v5950 != 0 {
		goto L1445
	} else {
		goto L1470
	}
L1469:
	;
	goto L1468
L1470:
	;
	goto L1444
L1471:
	;
	v5956 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L1472
L1472:
	;
	if v5956 == int32(0) {
		goto L1444
	} else {
		goto L1473
	}
L1473:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v5961 = m.ExcPending
	if v5961 != 0 {
		goto L1
	} else {
		goto L1474
	}
L1474:
	;
	goto L1444
L1475:
	;
	v5965 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v5965)
	goto L1444
L1476:
	;
	if v5797 != 0 {
		goto L1477
	} else {
		goto L1478
	}
L1477:
	;
	F_MemoryContextDelete(m, v5797)
	mBase = m.M
	v5972 = m.ExcPending
	if v5972 != 0 {
		goto L1
	} else {
		goto L1480
	}
L1478:
	;
	goto L1479
L1479:
	;
	v5974 = v5581 + int32(1)
	v5975 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if v5974 < v5975 {
		v5581 = v5974
		goto L1358
	} else {
		goto L1481
	}
L1480:
	;
	goto L1479
L1481:
	;
	goto L1359
L1482:
	;
	if v6017 != 0 {
		goto L1483
	} else {
		goto L1484
	}
L1483:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6020 = m.ExcPending
	if v6020 != 0 {
		goto L1
	} else {
		goto L1486
	}
L1484:
	;
	goto L1485
L1485:
	;
	v6022 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v6022 != 0 {
		goto L1487
	} else {
		goto L1488
	}
L1486:
	;
	goto L1485
L1487:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L1
	} else {
		goto L1490
	}
L1488:
	;
	goto L1489
L1489:
	;
	F_NullCommand(m, v5166)
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		goto L1
	} else {
		goto L1491
	}
L1490:
	;
	v6026 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v6026)
	goto L1489
L1491:
	;
	v6115 = v6006
	v6123 = int32(1)
	goto L1321
L1492:
	;
	if v6071 != 0 {
		goto L1493
	} else {
		goto L1494
	}
L1493:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6074 = m.ExcPending
	if v6074 != 0 {
		goto L1
	} else {
		goto L1496
	}
L1494:
	;
	goto L1495
L1495:
	;
	v6075 = int32(0)
	v6077 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v6077 == v6075 {
		v6115 = v5562
		v6123 = v6075
		goto L1321
	} else {
		goto L1497
	}
L1496:
	;
	goto L1495
L1497:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6081 = m.ExcPending
	if v6081 != 0 {
		goto L1
	} else {
		goto L1498
	}
L1498:
	;
	v6083 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v6083)
	v6115 = v5562
	v6123 = v6083
	goto L1321
L1499:
	;
	if v5168 != 0 {
		goto L1525
	} else {
		goto L1526
	}
L1500:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v6249, int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v6271 = m.ExcPending
	if v6271 != 0 {
		goto L1
	} else {
		goto L1524
	}
L1501:
	;
	v6147 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6148 = m.ExcPending
	if v6148 != 0 {
		goto L1
	} else {
		goto L1508
	}
L1502:
	;
	v6132 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6133 = m.ExcPending
	if v6133 != 0 {
		goto L1
	} else {
		goto L1504
	}
L1503:
	;
	switch v6126 - int32(1) {
	case 0:
		goto L1502
	case 1:
		goto L1501
	default:
		goto L1499
	}
L1504:
	;
	if v6132 == int32(0) {
		goto L1499
	} else {
		goto L1505
	}
L1505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5161))) = v5161 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v5161)
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L1
	} else {
		goto L1506
	}
L1506:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6143 = m.ExcPending
	if v6143 != 0 {
		goto L1
	} else {
		goto L1507
	}
L1507:
	;
	v6249 = int32(1471)
	goto L1500
L1508:
	;
	if v6147 == int32(0) {
		goto L1499
	} else {
		goto L1509
	}
L1509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5161)+36)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v5161)+32)) = v5161 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_181), v5161+int32(32))
	mBase = m.M
	v6159 = m.ExcPending
	if v6159 != 0 {
		goto L1
	} else {
		goto L1510
	}
L1510:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6161 = m.ExcPending
	if v6161 != 0 {
		goto L1
	} else {
		goto L1511
	}
L1511:
	;
	v6162 = int32(1478)
	if v6123 != 0 {
		v6249 = v6162
		goto L1500
	} else {
		goto L1512
	}
L1512:
	;
	v6163 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	if v6163 <= int32(0) {
		v6249 = v6162
		goto L1500
	} else {
		goto L1513
	}
L1513:
	;
	v6167 = int32(0)
	v6176 = v6163
	goto L1514
L1514:
	;
	v6204 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+12))
	v6208 = *(*int32)(unsafe.Add(mBase, uint32(v6204+v6167<<(uint(int32(2))%32))))
	v6209 = *(*int32)(unsafe.Add(mBase, uint32(v6208)+4))
	v6210 = *(*int32)(unsafe.Add(mBase, uint32(v6209)))
	if v6210 == int32(253) {
		goto L1517
	} else {
		goto L1518
	}
L1515:
	;
	v6223 = *(*int32)(unsafe.Add(mBase, uint32(v6215)+64))
	v6224 = *(*int32)(unsafe.Add(mBase, uint32(v6223)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5161)+16)) = v6224
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_174), v5161+int32(16))
	mBase = m.M
	v6230 = m.ExcPending
	if v6230 != 0 {
		goto L1
	} else {
		goto L1523
	}
L1516:
	;
	goto L1515
L1517:
	;
	v6213 = *(*int32)(unsafe.Add(mBase, uint32(v6209)+4))
	v6215 = F_FetchPreparedStatement(m, v6213, int32(0))
	mBase = m.M
	v6216 = m.ExcPending
	if v6216 != 0 {
		goto L1
	} else {
		goto L1520
	}
L1518:
	;
	v6219 = v6176
	goto L1519
L1519:
	;
	v6221 = v6167 + int32(1)
	if v6221 < v6219 {
		v6167 = v6221
		v6176 = v6219
		goto L1514
	} else {
		goto L1522
	}
L1520:
	;
	if v6215 != 0 {
		goto L1516
	} else {
		goto L1521
	}
L1521:
	;
	v6217 = *(*int32)(unsafe.Add(mBase, uint32(v5219)+4))
	v6219 = v6217
	goto L1519
L1522:
	;
	v6249 = v6162
	goto L1500
L1523:
	;
	v6249 = v6162
	goto L1500
L1524:
	;
	goto L1499
L1525:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_182))
	mBase = m.M
	v6311 = m.ExcPending
	if v6311 != 0 {
		goto L1
	} else {
		goto L1528
	}
L1526:
	;
	goto L1527
L1527:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	m.G0 = v5161 + int32(112)
	goto L222
L1528:
	;
	goto L1527
L1529:
	;
	v6363 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v6363 < int32(0) {
		goto L1531
	} else {
		goto L1532
	}
L1530:
	;
	v6370 = v40 + int32(440)
	v6371 = F_pq_getmsgstring(m, v6370)
	mBase = m.M
	v6372 = m.ExcPending
	if v6372 != 0 {
		goto L1
	} else {
		goto L1534
	}
L1531:
	;
	v6367 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v6367
	goto L1533
L1532:
	;
	goto L1533
L1533:
	;
	goto L1530
L1534:
	;
	v6373 = F_pq_getmsgstring(m, v6370)
	mBase = m.M
	v6374 = m.ExcPending
	if v6374 != 0 {
		goto L1
	} else {
		goto L1535
	}
L1535:
	;
	v6376 = F_pq_getmsgint(m, v6370, int32(2))
	mBase = m.M
	v6377 = m.ExcPending
	if v6377 != 0 {
		goto L1
	} else {
		goto L1536
	}
L1536:
	;
	if int32(0) < v6376 {
		goto L1537
	} else {
		goto L1538
	}
L1537:
	;
	v6383 = F_palloc(m, v6376<<(uint(int32(2))%32))
	mBase = m.M
	v6384 = m.ExcPending
	if v6384 != 0 {
		goto L1
	} else {
		goto L1540
	}
L1538:
	;
	v6444 = v1
	goto L1539
L1539:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v6474 = m.ExcPending
	if v6474 != 0 {
		goto L1
	} else {
		goto L1545
	}
L1540:
	;
	v6391 = int32(0)
	goto L1541
L1541:
	;
	v6428 = F_pq_getmsgint(m, v40+int32(440), int32(4))
	mBase = m.M
	v6429 = m.ExcPending
	if v6429 != 0 {
		goto L1
	} else {
		goto L1543
	}
L1542:
	;
	v6444 = v6383
	goto L1539
L1543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6383+v6391<<(uint(int32(2))%32)))) = v6428
	v6432 = v6391 + int32(1)
	if v6432 != v6376 {
		v6391 = v6432
		goto L1541
	} else {
		goto L1544
	}
L1544:
	;
	goto L1542
L1545:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v6373
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v6444
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v6376
	v6480 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	F_pgstat_report_activity(m, int32(3), v6373)
	mBase = m.M
	if v6480 == int32(1) {
		goto L1546
	} else {
		goto L1547
	}
L1546:
	;
	v6485 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v6495 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1549
L1547:
	;
	goto L1548
L1548:
	;
	v6500 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6501 = m.ExcPending
	if v6501 != 0 {
		goto L1
	} else {
		goto L1550
	}
L1549:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1548
L1550:
	;
	if v6500 != 0 {
		goto L1551
	} else {
		goto L1552
	}
L1551:
	;
	v6502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6371))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v6373
	if v6502 != 0 {
		goto L1554
	} else {
		goto L1555
	}
L1552:
	;
	goto L1553
L1553:
	;
	F_start_xact_command(m)
	mBase = m.M
	v6519 = m.ExcPending
	if v6519 != 0 {
		goto L1
	} else {
		goto L1559
	}
L1554:
	;
	v6505 = v6371
	goto L1556
L1555:
	;
	v6505 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1556
L1556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v6505
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_184), v40-int32(-64))
	mBase = m.M
	v6511 = m.ExcPending
	if v6511 != 0 {
		goto L1
	} else {
		goto L1557
	}
L1557:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1526), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L1
	} else {
		goto L1558
	}
L1558:
	;
	goto L1553
L1559:
	;
	v6520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6371))))
	if v6520 != 0 {
		goto L1561
	} else {
		goto L1562
	}
L1560:
	;
	v6542 = int32(_a_F_PostgresMainLoopOnce_135)
	v6543 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6540
	v6547 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v6547 == int32(1) {
		goto L1569
	} else {
		goto L1570
	}
L1561:
	;
	v6522 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6540 = v6522
	v6541 = int32(0)
	goto L1560
L1562:
	;
	goto L1563
L1563:
	;
	v6525 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117]))
	if v6525 != 0 {
		goto L1564
	} else {
		goto L1565
	}
L1564:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int32(0)
	F_DropCachedPlan(m, v6525)
	mBase = m.M
	v6530 = m.ExcPending
	if v6530 != 0 {
		goto L1
	} else {
		goto L1567
	}
L1565:
	;
	goto L1566
L1566:
	;
	v6532 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6537 = F_AllocSetContextCreateInternal(m, v6532, int32(_a_F_PostgresMainLoopOnce_186), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v6538 = m.ExcPending
	if v6538 != 0 {
		goto L1
	} else {
		goto L1568
	}
L1567:
	;
	goto L1566
L1568:
	;
	v6540 = v6537
	v6541 = v6537
	goto L1560
L1569:
	;
	v6550 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v6560 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1572
L1570:
	;
	goto L1571
L1571:
	;
	v6564 = F_raw_parser(m, v6373, int32(0))
	mBase = m.M
	v6565 = m.ExcPending
	if v6565 != 0 {
		goto L1
	} else {
		goto L1573
	}
L1572:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1571
L1573:
	;
	v6567 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])))
	if v6567 == int32(1) {
		goto L1574
	} else {
		goto L1575
	}
L1574:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_172))
	mBase = m.M
	v6572 = m.ExcPending
	if v6572 != 0 {
		goto L1
	} else {
		goto L1577
	}
L1575:
	;
	goto L1576
L1576:
	;
	if v6564 != 0 {
		goto L1579
	} else {
		goto L1580
	}
L1577:
	;
	goto L1576
L1578:
	;
	if v6541 != 0 {
		goto L1604
	} else {
		goto L1605
	}
L1579:
	;
	v6573 = *(*int32)(unsafe.Add(mBase, uint32(v6564)+4))
	if int32(2) <= v6573 {
		goto L197
	} else {
		goto L1582
	}
L1580:
	;
	goto L1581
L1581:
	;
	v6630 = int32(0)
	v6633 = F_CreateCachedPlan(m, v6630, v6373, v6630)
	mBase = m.M
	v6634 = m.ExcPending
	if v6634 != 0 {
		goto L1
	} else {
		goto L1603
	}
L1582:
	;
	v6576 = *(*int32)(unsafe.Add(mBase, uint32(v6564)+12))
	v6577 = *(*int32)(unsafe.Add(mBase, uint32(v6576)))
	v6579 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v6580 = *(*int32)(unsafe.Add(mBase, uint32(v6579)+24))
	goto L1583
L1583:
	;
	v6587 = *(*int32)(unsafe.Add(mBase, uint32(v6577)+4))
	if (v6580-int32(7))&int32(-9) == int32(0) {
		goto L1584
	} else {
		goto L1585
	}
L1584:
	;
	if v6587 == int32(0) {
		goto L196
	} else {
		goto L1587
	}
L1585:
	;
	goto L1586
L1586:
	;
	v6598 = F_CreateCommandTag(m, v6587)
	mBase = m.M
	v6599 = m.ExcPending
	if v6599 != 0 {
		goto L1
	} else {
		goto L1590
	}
L1587:
	;
	v6590 = *(*int32)(unsafe.Add(mBase, uint32(v6587)))
	if v6590 != int32(225) {
		goto L196
	} else {
		goto L1588
	}
L1588:
	;
	v6593 = *(*int32)(unsafe.Add(mBase, uint32(v6587)+4))
	if (v6593-int32(2))&int32(-6) != 0 {
		goto L196
	} else {
		goto L1589
	}
L1589:
	;
	goto L1586
L1590:
	;
	v6600 = F_CreateCachedPlan(m, v6577, v6373, v6598)
	mBase = m.M
	v6601 = m.ExcPending
	if v6601 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1591:
	;
	v6604 = *(*int32)(unsafe.Add(mBase, uint32(v6577)+4))
	v6605 = *(*int32)(unsafe.Add(mBase, uint32(v6604)))
	switch v6605 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v6609 = int32(1)
		goto L1593
	default:
		goto L1594
	}
L1592:
	;
	if v6609 == int32(0) {
		goto L1595
	} else {
		goto L1596
	}
L1593:
	;
	goto L1592
L1594:
	;
	v6609 = int32(0)
	goto L1593
L1595:
	;
	v6616 = F_pg_analyze_and_rewrite_varparams(m, v6577, v6373, v40+int32(512), v40+int32(460))
	mBase = m.M
	v6617 = m.ExcPending
	if v6617 != 0 {
		goto L1
	} else {
		goto L1598
	}
L1596:
	;
	goto L1597
L1597:
	;
	v6618 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v6619 = m.ExcPending
	if v6619 != 0 {
		goto L1
	} else {
		goto L1599
	}
L1598:
	;
	v6635 = v6600
	v6638 = v6616
	goto L1578
L1599:
	;
	F_PushActiveSnapshot(m, v6618)
	mBase = m.M
	v6621 = m.ExcPending
	if v6621 != 0 {
		goto L1
	} else {
		goto L1600
	}
L1600:
	;
	v6626 = F_pg_analyze_and_rewrite_varparams(m, v6577, v6373, v40+int32(512), v40+int32(460))
	mBase = m.M
	v6627 = m.ExcPending
	if v6627 != 0 {
		goto L1
	} else {
		goto L1601
	}
L1601:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v6629 = m.ExcPending
	if v6629 != 0 {
		goto L1
	} else {
		goto L1602
	}
L1602:
	;
	v6635 = v6600
	v6638 = v6626
	goto L1578
L1603:
	;
	v6635 = v6633
	v6638 = v6630
	goto L1578
L1604:
	;
	v6639 = *(*int32)(unsafe.Add(mBase, uint32(v6635)+56))
	v6641 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6645 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+16))
	if v6645 != v6641 {
		goto L1608
	} else {
		goto L1609
	}
L1605:
	;
	goto L1606
L1606:
	;
	v6674 = *(*int32)(unsafe.Add(mBase, uint32(v40)+512))
	v6675 = *(*int32)(unsafe.Add(mBase, uint32(v40)+460))
	v6676 = int32(0)
	F_CompleteCachedPlan(m, v6635, v6638, v6541, v6674, v6675, v6676, v6676, int32(2048), int32(1))
	mBase = m.M
	v6681 = m.ExcPending
	if v6681 != 0 {
		goto L1
	} else {
		goto L1624
	}
L1607:
	;
	goto L1606
L1608:
	;
	if v6645 == int32(0) {
		goto L1611
	} else {
		goto L1612
	}
L1609:
	;
	goto L1610
L1610:
	;
	goto L1607
L1611:
	;
	if v6641 != 0 {
		goto L1618
	} else {
		goto L1619
	}
L1612:
	;
	v6649 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+28))
	v6650 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+24))
	if v6650 != 0 {
		goto L1614
	} else {
		goto L1615
	}
L1613:
	;
	if v6649 == int32(0) {
		goto L1611
	} else {
		goto L1617
	}
L1614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6650)+28)) = v6649
	goto L1613
L1615:
	;
	goto L1616
L1616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6645)+20)) = v6649
	goto L1613
L1617:
	;
	v6655 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6649)+24)) = v6655
	goto L1611
L1618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6639)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6639)+16)) = v6641
	v6662 = *(*int32)(unsafe.Add(mBase, uint32(v6641)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6639)+28)) = v6662
	if v6662 != 0 {
		goto L1621
	} else {
		goto L1622
	}
L1619:
	;
	goto L1620
L1620:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6639)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6639)+16)) = int32(0)
	goto L1610
L1621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6662)+24)) = v6639
	goto L1623
L1622:
	;
	goto L1623
L1623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6641)+20)) = v6639
	goto L1607
L1624:
	;
	v6683 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v6683 != 0 {
		goto L1625
	} else {
		goto L1626
	}
L1625:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6685 = m.ExcPending
	if v6685 != 0 {
		goto L1
	} else {
		goto L1628
	}
L1626:
	;
	goto L1627
L1627:
	;
	if v6520 != 0 {
		goto L1630
	} else {
		goto L1631
	}
L1628:
	;
	goto L1627
L1629:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6543
	F_CommandCounterIncrement(m)
	mBase = m.M
	v6696 = m.ExcPending
	if v6696 != 0 {
		goto L1
	} else {
		goto L1635
	}
L1630:
	;
	F_StorePreparedStatement(m, v6371, v6635, int32(0))
	mBase = m.M
	v6688 = m.ExcPending
	if v6688 != 0 {
		goto L1
	} else {
		goto L1633
	}
L1631:
	;
	goto L1632
L1632:
	;
	F_SaveCachedPlan(m, v6635)
	mBase = m.M
	v6690 = m.ExcPending
	if v6690 != 0 {
		goto L1
	} else {
		goto L1634
	}
L1633:
	;
	goto L1629
L1634:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v6635
	goto L1629
L1635:
	;
	v6698 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v6698 == int32(2) {
		goto L1636
	} else {
		goto L1637
	}
L1636:
	;
	F_pq_putemptymessage(m, int32(49))
	mBase = m.M
	v6703 = m.ExcPending
	if v6703 != 0 {
		goto L1
	} else {
		goto L1639
	}
L1637:
	;
	goto L1638
L1638:
	;
	v6707 = F_check_log_duration(m, v40+int32(480), int32(0))
	mBase = m.M
	v6708 = m.ExcPending
	if v6708 != 0 {
		goto L1
	} else {
		goto L1644
	}
L1639:
	;
	goto L1638
L1640:
	;
	if v6480 != 0 {
		goto L1656
	} else {
		goto L1657
	}
L1641:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6749 = m.ExcPending
	if v6749 != 0 {
		goto L1
	} else {
		goto L1654
	}
L1642:
	;
	v6728 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6729 = m.ExcPending
	if v6729 != 0 {
		goto L1
	} else {
		goto L1648
	}
L1643:
	;
	v6713 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6714 = m.ExcPending
	if v6714 != 0 {
		goto L1
	} else {
		goto L1645
	}
L1644:
	;
	switch v6707 - int32(1) {
	case 0:
		goto L1643
	case 1:
		goto L1642
	default:
		goto L1640
	}
L1645:
	;
	if v6713 == int32(0) {
		goto L1640
	} else {
		goto L1646
	}
L1646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(32))
	mBase = m.M
	v6724 = m.ExcPending
	if v6724 != 0 {
		goto L1
	} else {
		goto L1647
	}
L1647:
	;
	v6747 = int32(1707)
	goto L1641
L1648:
	;
	if v6728 == int32(0) {
		goto L1640
	} else {
		goto L1649
	}
L1649:
	;
	v6732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6371))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v6373
	if v6732 != 0 {
		goto L1650
	} else {
		goto L1651
	}
L1650:
	;
	v6735 = v6371
	goto L1652
L1651:
	;
	v6735 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1652
L1652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+52)) = v6735
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_187), v40+int32(48))
	mBase = m.M
	v6744 = m.ExcPending
	if v6744 != 0 {
		goto L1
	} else {
		goto L1653
	}
L1653:
	;
	v6747 = int32(1715)
	goto L1641
L1654:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v6747, int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v6753 = m.ExcPending
	if v6753 != 0 {
		goto L1
	} else {
		goto L1655
	}
L1655:
	;
	goto L1640
L1656:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_188))
	mBase = m.M
	v6757 = m.ExcPending
	if v6757 != 0 {
		goto L1
	} else {
		goto L1659
	}
L1657:
	;
	goto L1658
L1658:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L198
L1659:
	;
	goto L1658
L1660:
	;
	v6766 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v6766 < int32(0) {
		goto L1662
	} else {
		goto L1663
	}
L1661:
	;
	v6773 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	v6775 = v40 + int32(440)
	v6776 = F_pq_getmsgstring(m, v6775)
	mBase = m.M
	v6777 = m.ExcPending
	if v6777 != 0 {
		goto L1
	} else {
		goto L1665
	}
L1662:
	;
	v6770 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v6770
	goto L1664
L1663:
	;
	goto L1664
L1664:
	;
	goto L1661
L1665:
	;
	v6778 = F_pq_getmsgstring(m, v6775)
	mBase = m.M
	v6779 = m.ExcPending
	if v6779 != 0 {
		goto L1
	} else {
		goto L1666
	}
L1666:
	;
	v6782 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6783 = m.ExcPending
	if v6783 != 0 {
		goto L1
	} else {
		goto L1667
	}
L1667:
	;
	if v6782 != 0 {
		goto L1668
	} else {
		goto L1669
	}
L1668:
	;
	v6784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6776))))
	v6786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6778))))
	if v6786 != 0 {
		goto L1671
	} else {
		goto L1672
	}
L1669:
	;
	goto L1670
L1670:
	;
	v6803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6778))))
	if v6803 != 0 {
		goto L1680
	} else {
		goto L1681
	}
L1671:
	;
	v6787 = v6778
	goto L1673
L1672:
	;
	v6787 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1673
L1673:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+212)) = v6787
	if v6784 != 0 {
		goto L1674
	} else {
		goto L1675
	}
L1674:
	;
	v6790 = v6776
	goto L1676
L1675:
	;
	v6790 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1676
L1676:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+208)) = v6790
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_189), v40+int32(208))
	mBase = m.M
	v6796 = m.ExcPending
	if v6796 != 0 {
		goto L1
	} else {
		goto L1677
	}
L1677:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1761), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v6801 = m.ExcPending
	if v6801 != 0 {
		goto L1
	} else {
		goto L1678
	}
L1678:
	;
	goto L1670
L1679:
	;
	v6814 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v6814
	F_pgstat_report_activity(m, int32(3), v6814)
	mBase = m.M
	v6818 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+60))
	if v6818 == int32(0) {
		goto L1685
	} else {
		goto L1686
	}
L1680:
	;
	v6805 = F_FetchPreparedStatement(m, v6778, int32(1))
	mBase = m.M
	v6806 = m.ExcPending
	if v6806 != 0 {
		goto L1
	} else {
		goto L1683
	}
L1681:
	;
	goto L1682
L1682:
	;
	v6809 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117]))
	if v6809 == int32(0) {
		goto L195
	} else {
		goto L1684
	}
L1683:
	;
	v6807 = *(*int32)(unsafe.Add(mBase, uint32(v6805)+64))
	v6812 = v6807
	goto L1679
L1684:
	;
	v6812 = v6809
	goto L1679
L1685:
	;
	if v6773&int32(1) != 0 {
		goto L1699
	} else {
		goto L1700
	}
L1686:
	;
	v6821 = *(*int32)(unsafe.Add(mBase, uint32(v6818)+4))
	if v6821 <= int32(0) {
		goto L1685
	} else {
		goto L1687
	}
L1687:
	;
	v6824 = *(*int32)(unsafe.Add(mBase, uint32(v6818)+12))
	v6832 = int32(0)
	goto L1688
L1688:
	;
	v6866 = *(*int32)(unsafe.Add(mBase, uint32(v6824+v6832<<(uint(int32(2))%32))))
	v6867 = *(*int64)(unsafe.Add(mBase, uint32(v6866)+16))
	if v6867 == int64(0) {
		goto L1690
	} else {
		goto L1691
	}
L1689:
	;
	v6876 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v6876 == int32(0) {
		goto L1695
	} else {
		goto L1696
	}
L1690:
	;
	v6871 = v6832 + int32(1)
	if v6871 != v6821 {
		v6832 = v6871
		goto L1688
	} else {
		goto L1693
	}
L1691:
	;
	goto L1692
L1692:
	;
	goto L1689
L1693:
	;
	goto L1685
L1694:
	;
	goto L1685
L1695:
	;
	goto L1694
L1696:
	;
	v6880 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v6880&int32(1) == int32(0) {
		goto L1695
	} else {
		goto L1697
	}
L1697:
	;
	v6887 = *(*int64)(unsafe.Add(mBase, uint32(v6876)+392))
	if int32(1)&base.B2i32(v6887 != int64(0)) != 0 {
		goto L1695
	} else {
		goto L1698
	}
L1698:
	;
	v6891 = int32(_a_F_PostgresMainLoopOnce_176)
	v6893 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v6894 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v6893 + v6894
	v6897 = *(*int32)(unsafe.Add(mBase, uint32(v6876)))
	*(*int32)(unsafe.Add(mBase, uint32(v6876))) = v6897 + v6894
	*(*int64)(unsafe.Add(mBase, uint32(v6876)+392)) = v6867
	*(*int32)(unsafe.Add(mBase, uint32(v6876))) = v6897 + int32(2)
	v6908 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v6908 - v6894
	goto L1695
L1699:
	;
	v6951 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v6961 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1702
L1700:
	;
	goto L1701
L1701:
	;
	F_start_xact_command(m)
	mBase = m.M
	v6965 = m.ExcPending
	if v6965 != 0 {
		goto L1
	} else {
		goto L1703
	}
L1702:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1701
L1703:
	;
	v6968 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6968
	v6973 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v6974 = m.ExcPending
	if v6974 != 0 {
		goto L1
	} else {
		goto L1704
	}
L1704:
	;
	if int32(0) < v6973 {
		goto L1705
	} else {
		goto L1706
	}
L1705:
	;
	v6980 = F_palloc(m, v6973<<(uint(int32(1))%32))
	mBase = m.M
	v6981 = m.ExcPending
	if v6981 != 0 {
		goto L1
	} else {
		goto L1708
	}
L1706:
	;
	v7032 = v1
	goto L1707
L1707:
	;
	v7071 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7072 = m.ExcPending
	if v7072 != 0 {
		goto L1
	} else {
		goto L1713
	}
L1708:
	;
	v6988 = int32(0)
	goto L1709
L1709:
	;
	v7025 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7026 = m.ExcPending
	if v7026 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1710:
	;
	v7032 = v6980
	goto L1707
L1711:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v6980+v6988<<(uint(int32(1))%32)))) = uint16(v7025)
	v7029 = v6988 + int32(1)
	if v7029 != v6973 {
		v6988 = v7029
		goto L1709
	} else {
		goto L1712
	}
L1712:
	;
	goto L1710
L1713:
	;
	if base.B2i32(v7071 != v6973)&base.B2i32(int32(2) <= v6973) != 0 {
		goto L194
	} else {
		goto L1714
	}
L1714:
	;
	v7077 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+24))
	if v7071 != v7077 {
		goto L193
	} else {
		goto L1715
	}
L1715:
	;
	v7080 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v7081 = *(*int32)(unsafe.Add(mBase, uint32(v7080)+24))
	goto L1716
L1716:
	;
	if (v7081-int32(7))&int32(-9) == int32(0) {
		goto L1717
	} else {
		goto L1718
	}
L1717:
	;
	v7088 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+4))
	if v7088 == int32(0) {
		goto L192
	} else {
		goto L1720
	}
L1718:
	;
	goto L1719
L1719:
	;
	v7104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6776))))
	if v7104 == int32(0) {
		goto L1725
	} else {
		goto L1726
	}
L1720:
	;
	v7091 = *(*int32)(unsafe.Add(mBase, uint32(v7088)+4))
	if v7091 == int32(0) {
		goto L192
	} else {
		goto L1721
	}
L1721:
	;
	v7094 = *(*int32)(unsafe.Add(mBase, uint32(v7091)))
	if v7094 != int32(225) {
		goto L192
	} else {
		goto L1722
	}
L1722:
	;
	v7097 = *(*int32)(unsafe.Add(mBase, uint32(v7091)+4))
	if (v7097-int32(2))&int32(-6)|v7071 != 0 {
		goto L192
	} else {
		goto L1723
	}
L1723:
	;
	goto L1719
L1724:
	;
	v7116 = int32(_a_F_PostgresMainLoopOnce_135)
	v7117 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7119 = *(*int32)(unsafe.Add(mBase, uint32(v7115)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7119
	v7121 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+12))
	v7122 = F_pstrdup(m, v7121)
	mBase = m.M
	v7123 = m.ExcPending
	if v7123 != 0 {
		goto L1
	} else {
		goto L1730
	}
L1725:
	;
	v7107 = int32(1)
	v7109 = F_CreatePortal(m, v6776, v7107, v7107)
	mBase = m.M
	v7110 = m.ExcPending
	if v7110 != 0 {
		goto L1
	} else {
		goto L1728
	}
L1726:
	;
	goto L1727
L1727:
	;
	v7111 = int32(0)
	v7113 = F_CreatePortal(m, v6776, v7111, v7111)
	mBase = m.M
	v7114 = m.ExcPending
	if v7114 != 0 {
		goto L1
	} else {
		goto L1729
	}
L1728:
	;
	v7115 = v7109
	goto L1724
L1729:
	;
	v7115 = v7113
	goto L1724
L1730:
	;
	v7124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6778))))
	if v7124 != 0 {
		goto L1731
	} else {
		goto L1732
	}
L1731:
	;
	v7125 = F_pstrdup(m, v6778)
	mBase = m.M
	v7126 = m.ExcPending
	if v7126 != 0 {
		goto L1
	} else {
		goto L1734
	}
L1732:
	;
	v7127 = v1
	goto L1733
L1733:
	;
	if v7071 <= int32(0) {
		goto L1737
	} else {
		goto L1738
	}
L1734:
	;
	v7127 = v7125
	goto L1733
L1735:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7117
	v7443 = *(*int32)(unsafe.Add(mBase, uint32(v7115)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = v7432
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v7443
	v7446 = int32(_a_F_PostgresMainLoopOnce_191)
	v7447 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v40 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+516)) = int32(1146)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v7447
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v40 + int32(460)
	v7462 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7463 = m.ExcPending
	if v7463 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1736:
	;
	v7432 = v7394
	v7440 = int32(1)
	goto L1735
L1737:
	;
	v7130 = int32(0)
	v7131 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+4))
	if v7131 == v7130 {
		v7432 = v1
		v7440 = v7130
		goto L1735
	} else {
		goto L1740
	}
L1738:
	;
	goto L1739
L1739:
	;
	v7149 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7150 = m.ExcPending
	if v7150 != 0 {
		goto L1
	} else {
		goto L1747
	}
L1740:
	;
	v7137 = *(*int32)(unsafe.Add(mBase, uint32(v7131)+4))
	v7138 = *(*int32)(unsafe.Add(mBase, uint32(v7137)))
	switch v7138 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v7142 = int32(1)
		goto L1742
	default:
		goto L1743
	}
L1741:
	;
	if v7142 == int32(0) {
		v7432 = v1
		v7440 = int32(0)
		goto L1735
	} else {
		goto L1744
	}
L1742:
	;
	goto L1741
L1743:
	;
	v7142 = int32(0)
	goto L1742
L1744:
	;
	v7145 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7146 = m.ExcPending
	if v7146 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1745:
	;
	F_PushActiveSnapshot(m, v7145)
	mBase = m.M
	v7148 = m.ExcPending
	if v7148 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	v7394 = v1
	goto L1736
L1747:
	;
	F_PushActiveSnapshot(m, v7149)
	mBase = m.M
	v7152 = m.ExcPending
	if v7152 != 0 {
		goto L1
	} else {
		goto L1748
	}
L1748:
	;
	v7153 = *(*int32)(unsafe.Add(mBase, uint32(v7115)))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+464)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v7153
	v7157 = int32(_a_F_PostgresMainLoopOnce_191)
	v7158 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v40 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+516)) = int32(1145)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v7158
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v40 + int32(460)
	v7171 = F_makeParamList(m, v7071)
	mBase = m.M
	v7172 = m.ExcPending
	if v7172 != 0 {
		goto L1
	} else {
		goto L1749
	}
L1749:
	;
	v7175 = int32(0)
	v7184 = v7175
	v7202 = v1
	goto L1750
L1750:
	;
	v7216 = v7184 << (uint(int32(2)) % 32)
	v7217 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+20))
	v7219 = *(*int32)(unsafe.Add(mBase, uint32(v7216+v7217)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = v7184
	v7226 = F_pq_getmsgint(m, v40+int32(440), int32(4))
	mBase = m.M
	v7227 = m.ExcPending
	if v7227 != 0 {
		goto L1
	} else {
		goto L1753
	}
L1751:
	;
	v7353 = int32(_a_F_PostgresMainLoopOnce_191)
	v7355 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v7356 = *(*int32)(unsafe.Add(mBase, uint32(v7355)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v7356
	v7359 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	if v7359 == int32(0) {
		v7394 = v7171
		goto L1736
	} else {
		goto L1799
	}
L1752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+480)) = v7243
	if int32(2) <= v6973 {
		goto L1762
	} else {
		goto L1763
	}
L1753:
	;
	v7229 = base.B2i32(v7226 == int32(-1))
	if v7226 == int32(-1) {
		goto L1754
	} else {
		goto L1755
	}
L1754:
	;
	v7230 = int32(0)
	v7243 = v7230
	v7244 = v7230
	goto L1752
L1755:
	;
	goto L1756
L1756:
	;
	v7234 = F_pq_getmsgbytes(m, v40+int32(440), v7226)
	mBase = m.M
	v7235 = m.ExcPending
	if v7235 != 0 {
		goto L1
	} else {
		goto L1757
	}
L1757:
	;
	v7236 = v7234 + v7226
	v7237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7236))))
	v7238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7236))) = uint8(v7238)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+488)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+484)) = v7226
	v7243 = v7234
	v7244 = v7237
	goto L1752
L1758:
	;
	if v7229 == int32(0) {
		goto L1795
	} else {
		goto L1796
	}
L1759:
	;
	F_getTypeBinaryInputInfo(m, v7219, v40+int32(472), v40+int32(456))
	mBase = m.M
	v7318 = m.ExcPending
	if v7318 != 0 {
		goto L1
	} else {
		goto L1788
	}
L1760:
	;
	F_getTypeInputInfo(m, v7219, v40+int32(472), v40+int32(456))
	mBase = m.M
	v7260 = m.ExcPending
	if v7260 != 0 {
		goto L1
	} else {
		goto L1766
	}
L1761:
	;
	v7253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7252))))
	switch v7253 {
	case 0:
		goto L1760
	case 1:
		goto L1759
	default:
		goto L190
	}
L1762:
	;
	v7252 = v7032 + v7184<<(uint(int32(1))%32)
	goto L1761
L1763:
	;
	goto L1764
L1764:
	;
	if v6973 <= v7175 {
		goto L1760
	} else {
		goto L1765
	}
L1765:
	;
	v7252 = v7032
	goto L1761
L1766:
	;
	if v7226 == int32(-1) {
		goto L1767
	} else {
		goto L1768
	}
L1767:
	;
	v7265 = int32(0)
	goto L1769
L1768:
	;
	v7262 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	v7263 = F_pg_client_to_server(m, v7262, v7226)
	mBase = m.M
	v7264 = m.ExcPending
	if v7264 != 0 {
		goto L1
	} else {
		goto L1770
	}
L1769:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v7265
	v7267 = *(*int32)(unsafe.Add(mBase, uint32(v40)+472))
	v7268 = *(*int32)(unsafe.Add(mBase, uint32(v40)+456))
	v7270 = F_OidInputFunctionCall(m, v7267, v7265, v7268, int32(-1))
	mBase = m.M
	v7271 = m.ExcPending
	if v7271 != 0 {
		goto L1
	} else {
		goto L1771
	}
L1770:
	;
	v7265 = v7263
	goto L1769
L1771:
	;
	v7272 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v7272
	if v7265 == v7272 {
		v7335 = v7270
		v7336 = v7202
		goto L1758
	} else {
		goto L1772
	}
L1772:
	;
	v7277 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	if v7277 != 0 {
		goto L1773
	} else {
		goto L1774
	}
L1773:
	;
	v7278 = int32(_a_F_PostgresMainLoopOnce_135)
	v7279 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7282 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7282
	if v7202 == int32(0) {
		goto L1777
	} else {
		goto L1778
	}
L1774:
	;
	v7308 = v7202
	goto L1775
L1775:
	;
	v7309 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	if v7265 == v7309 {
		v7335 = v7270
		v7336 = v7308
		goto L1758
	} else {
		goto L1786
	}
L1776:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7216+v7291))) = v7300
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7279
	v7308 = v7291
	goto L1775
L1777:
	;
	v7286 = F_palloc0(m, v7071<<(uint(int32(2))%32))
	mBase = m.M
	v7287 = m.ExcPending
	if v7287 != 0 {
		goto L1
	} else {
		goto L1780
	}
L1778:
	;
	v7290 = v7277
	v7291 = v7202
	goto L1779
L1779:
	;
	if v7290 < int32(0) {
		goto L1781
	} else {
		goto L1782
	}
L1780:
	;
	v7289 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v7290 = v7289
	v7291 = v7286
	goto L1779
L1781:
	;
	v7294 = F_pstrdup(m, v7265)
	mBase = m.M
	v7295 = m.ExcPending
	if v7295 != 0 {
		goto L1
	} else {
		goto L1784
	}
L1782:
	;
	goto L1783
L1783:
	;
	v7298 = F_pnstrdup(m, v7265, v7290+int32(8))
	mBase = m.M
	v7299 = m.ExcPending
	if v7299 != 0 {
		goto L1
	} else {
		goto L1785
	}
L1784:
	;
	v7300 = v7294
	goto L1776
L1785:
	;
	v7300 = v7298
	goto L1776
L1786:
	;
	F_pfree(m, v7265)
	mBase = m.M
	v7312 = m.ExcPending
	if v7312 != 0 {
		goto L1
	} else {
		goto L1787
	}
L1787:
	;
	v7335 = v7270
	v7336 = v7308
	goto L1758
L1788:
	;
	v7319 = *(*int32)(unsafe.Add(mBase, uint32(v40)+472))
	if v7226 == int32(-1) {
		goto L1789
	} else {
		goto L1790
	}
L1789:
	;
	v7323 = int32(0)
	goto L1791
L1790:
	;
	v7323 = v40 + int32(480)
	goto L1791
L1791:
	;
	v7324 = *(*int32)(unsafe.Add(mBase, uint32(v40)+456))
	v7326 = F_OidReceiveFunctionCall(m, v7319, v7323, v7324, int32(-1))
	mBase = m.M
	v7327 = m.ExcPending
	if v7327 != 0 {
		goto L1
	} else {
		goto L1792
	}
L1792:
	;
	if v7226 == int32(-1) {
		v7335 = v7326
		v7336 = v7202
		goto L1758
	} else {
		goto L1793
	}
L1793:
	;
	v7328 = *(*int32)(unsafe.Add(mBase, uint32(v40)+492))
	v7329 = *(*int32)(unsafe.Add(mBase, uint32(v40)+484))
	if v7328 != v7329 {
		goto L191
	} else {
		goto L1794
	}
L1794:
	;
	v7335 = v7326
	v7336 = v7202
	goto L1758
L1795:
	;
	v7339 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	*(*uint8)(unsafe.Add(mBase, uint32(v7339+v7226))) = uint8(v7244)
	goto L1797
L1796:
	;
	goto L1797
L1797:
	;
	v7344 = v7171 + int32(32) + v7184*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v7344)+8)) = v7219
	v7346 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7344)+6)) = uint16(v7346)
	*(*uint8)(unsafe.Add(mBase, uint32(v7344)+4)) = uint8(v7229)
	*(*int32)(unsafe.Add(mBase, uint32(v7344))) = v7335
	v7351 = v7184 + v7346
	if v7351 != v7071 {
		v7184 = v7351
		v7202 = v7336
		goto L1750
	} else {
		goto L1798
	}
L1798:
	;
	goto L1751
L1799:
	;
	v7362 = F_BuildParamLogString(m, v7171, v7336, v7359)
	mBase = m.M
	v7363 = m.ExcPending
	if v7363 != 0 {
		goto L1
	} else {
		goto L1800
	}
L1800:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7171)+24)) = v7362
	v7394 = v7171
	goto L1736
L1801:
	;
	if int32(0) < v7462 {
		goto L1802
	} else {
		goto L1803
	}
L1802:
	;
	v7469 = F_palloc(m, v7462<<(uint(int32(1))%32))
	mBase = m.M
	v7470 = m.ExcPending
	if v7470 != 0 {
		goto L1
	} else {
		goto L1805
	}
L1803:
	;
	v7530 = int32(0)
	goto L1804
L1804:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v7560 = m.ExcPending
	if v7560 != 0 {
		goto L1
	} else {
		goto L1810
	}
L1805:
	;
	v7477 = int32(0)
	goto L1806
L1806:
	;
	v7514 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7515 = m.ExcPending
	if v7515 != 0 {
		goto L1
	} else {
		goto L1808
	}
L1807:
	;
	v7530 = v7469
	goto L1804
L1808:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7469+v7477<<(uint(int32(1))%32)))) = uint16(v7514)
	v7518 = v7477 + int32(1)
	if v7518 != v7462 {
		v7477 = v7518
		goto L1806
	} else {
		goto L1809
	}
L1809:
	;
	goto L1807
L1810:
	;
	v7561 = int32(0)
	v7563 = F_GetCachedPlan(m, v6812, v7432, v7561, v7561)
	mBase = m.M
	v7564 = m.ExcPending
	if v7564 != 0 {
		goto L1
	} else {
		goto L1811
	}
L1811:
	;
	v7565 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+16))
	v7566 = *(*int32)(unsafe.Add(mBase, uint32(v7563)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v7115)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+40)) = v7565
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+32)) = v7122
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+4)) = v7127
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+60)) = v7563
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+56)) = v7566
	*(*int32)(unsafe.Add(mBase, uint32(v7115)+36)) = v7565
	goto L1812
L1812:
	;
	v7577 = *(*int32)(unsafe.Add(mBase, uint32(v7115)+56))
	if v7577 == int32(0) {
		goto L1813
	} else {
		goto L1814
	}
L1813:
	;
	if v7440 != 0 {
		goto L1827
	} else {
		goto L1828
	}
L1814:
	;
	v7580 = *(*int32)(unsafe.Add(mBase, uint32(v7577)+4))
	if v7580 <= int32(0) {
		goto L1813
	} else {
		goto L1815
	}
L1815:
	;
	v7583 = *(*int32)(unsafe.Add(mBase, uint32(v7577)+12))
	v7591 = int32(0)
	goto L1816
L1816:
	;
	v7625 = *(*int32)(unsafe.Add(mBase, uint32(v7583+v7591<<(uint(int32(2))%32))))
	v7626 = *(*int64)(unsafe.Add(mBase, uint32(v7625)+16))
	if v7626 == int64(0) {
		goto L1818
	} else {
		goto L1819
	}
L1817:
	;
	v7635 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v7635 == int32(0) {
		goto L1823
	} else {
		goto L1824
	}
L1818:
	;
	v7630 = v7591 + int32(1)
	if v7630 != v7580 {
		v7591 = v7630
		goto L1816
	} else {
		goto L1821
	}
L1819:
	;
	goto L1820
L1820:
	;
	goto L1817
L1821:
	;
	goto L1813
L1822:
	;
	goto L1813
L1823:
	;
	goto L1822
L1824:
	;
	v7639 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v7639&int32(1) == int32(0) {
		goto L1823
	} else {
		goto L1825
	}
L1825:
	;
	v7646 = *(*int64)(unsafe.Add(mBase, uint32(v7635)+400))
	if int32(1)&base.B2i32(v7646 != int64(0)) != 0 {
		goto L1823
	} else {
		goto L1826
	}
L1826:
	;
	v7650 = int32(_a_F_PostgresMainLoopOnce_176)
	v7652 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v7653 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v7652 + v7653
	v7656 = *(*int32)(unsafe.Add(mBase, uint32(v7635)))
	*(*int32)(unsafe.Add(mBase, uint32(v7635))) = v7656 + v7653
	*(*int64)(unsafe.Add(mBase, uint32(v7635)+400)) = v7626
	*(*int32)(unsafe.Add(mBase, uint32(v7635))) = v7656 + int32(2)
	v7667 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v7667 - v7653
	goto L1823
L1827:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v7709 = m.ExcPending
	if v7709 != 0 {
		goto L1
	} else {
		goto L1830
	}
L1828:
	;
	goto L1829
L1829:
	;
	v7710 = int32(0)
	F_PortalStart(m, v7115, v7432, v7710, v7710)
	mBase = m.M
	v7713 = m.ExcPending
	if v7713 != 0 {
		goto L1
	} else {
		goto L1831
	}
L1830:
	;
	goto L1829
L1831:
	;
	F_PortalSetResultFormat(m, v7115, v7462, v7530)
	mBase = m.M
	v7715 = m.ExcPending
	if v7715 != 0 {
		goto L1
	} else {
		goto L1832
	}
L1832:
	;
	v7716 = int32(_a_F_PostgresMainLoopOnce_191)
	v7718 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v7719 = *(*int32)(unsafe.Add(mBase, uint32(v7718)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v7719
	v7722 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v7722 == int32(2) {
		goto L1833
	} else {
		goto L1834
	}
L1833:
	;
	F_pq_putemptymessage(m, int32(50))
	mBase = m.M
	v7727 = m.ExcPending
	if v7727 != 0 {
		goto L1
	} else {
		goto L1836
	}
L1834:
	;
	goto L1835
L1835:
	;
	v7731 = F_check_log_duration(m, v40+int32(480), int32(0))
	mBase = m.M
	v7732 = m.ExcPending
	if v7732 != 0 {
		goto L1
	} else {
		goto L1841
	}
L1836:
	;
	goto L1835
L1837:
	;
	if v6773&int32(1) != 0 {
		goto L1867
	} else {
		goto L1868
	}
L1838:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v7808, int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v7812 = m.ExcPending
	if v7812 != 0 {
		goto L1
	} else {
		goto L1866
	}
L1839:
	;
	v7754 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7755 = m.ExcPending
	if v7755 != 0 {
		goto L1
	} else {
		goto L1846
	}
L1840:
	;
	v7737 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7738 = m.ExcPending
	if v7738 != 0 {
		goto L1
	} else {
		goto L1842
	}
L1841:
	;
	switch v7731 - int32(1) {
	case 0:
		goto L1840
	case 1:
		goto L1839
	default:
		goto L1837
	}
L1842:
	;
	if v7737 == int32(0) {
		goto L1837
	} else {
		goto L1843
	}
L1843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(96))
	mBase = m.M
	v7748 = m.ExcPending
	if v7748 != 0 {
		goto L1
	} else {
		goto L1844
	}
L1844:
	;
	F_errhidestmt(m)
	mBase = m.M
	v7750 = m.ExcPending
	if v7750 != 0 {
		goto L1
	} else {
		goto L1845
	}
L1845:
	;
	v7808 = int32(2185)
	goto L1838
L1846:
	;
	if v7754 == int32(0) {
		goto L1837
	} else {
		goto L1847
	}
L1847:
	;
	v7758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6778))))
	v7759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6776))))
	v7760 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+144)) = v7760
	if v7759 != 0 {
		goto L1848
	} else {
		goto L1849
	}
L1848:
	;
	v7763 = v6776
	goto L1850
L1849:
	;
	v7763 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1850
L1850:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v7763
	if v7759 != 0 {
		goto L1851
	} else {
		goto L1852
	}
L1851:
	;
	v7767 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L1853
L1852:
	;
	v7767 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1853
L1853:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+136)) = v7767
	if v7758 != 0 {
		goto L1854
	} else {
		goto L1855
	}
L1854:
	;
	v7770 = v6778
	goto L1856
L1855:
	;
	v7770 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1856
L1856:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+132)) = v7770
	*(*int32)(unsafe.Add(mBase, uint32(v40)+128)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_193), v40+int32(128))
	mBase = m.M
	v7779 = m.ExcPending
	if v7779 != 0 {
		goto L1
	} else {
		goto L1857
	}
L1857:
	;
	F_errhidestmt(m)
	mBase = m.M
	v7781 = m.ExcPending
	if v7781 != 0 {
		goto L1
	} else {
		goto L1858
	}
L1858:
	;
	v7782 = int32(2196)
	if v7432 == int32(0) {
		v7808 = v7782
		goto L1838
	} else {
		goto L1859
	}
L1859:
	;
	v7785 = *(*int32)(unsafe.Add(mBase, uint32(v7432)+28))
	if v7785 <= int32(0) {
		v7808 = v7782
		goto L1838
	} else {
		goto L1860
	}
L1860:
	;
	v7789 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128]))
	if v7789 == int32(0) {
		v7808 = v7782
		goto L1838
	} else {
		goto L1861
	}
L1861:
	;
	v7793 = F_BuildParamLogString(m, v7432, int32(0), v7789)
	mBase = m.M
	v7794 = m.ExcPending
	if v7794 != 0 {
		goto L1
	} else {
		goto L1862
	}
L1862:
	;
	if v7793 == int32(0) {
		v7808 = v7782
		goto L1838
	} else {
		goto L1863
	}
L1863:
	;
	v7797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7793))))
	if v7797 == int32(0) {
		v7808 = v7782
		goto L1838
	} else {
		goto L1864
	}
L1864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+112)) = v7793
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(112))
	mBase = m.M
	v7805 = m.ExcPending
	if v7805 != 0 {
		goto L1
	} else {
		goto L1865
	}
L1865:
	;
	v7808 = v7782
	goto L1838
L1866:
	;
	goto L1837
L1867:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_195))
	mBase = m.M
	v7820 = m.ExcPending
	if v7820 != 0 {
		goto L1
	} else {
		goto L1870
	}
L1868:
	;
	goto L1869
L1869:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L198
L1870:
	;
	goto L1869
L1871:
	;
	v7829 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v7829 < int32(0) {
		goto L1873
	} else {
		goto L1874
	}
L1872:
	;
	v7836 = v40 + int32(440)
	v7837 = F_pq_getmsgstring(m, v7836)
	mBase = m.M
	v7838 = m.ExcPending
	if v7838 != 0 {
		goto L1
	} else {
		goto L1876
	}
L1873:
	;
	v7833 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v7833
	goto L1875
L1874:
	;
	goto L1875
L1875:
	;
	goto L1872
L1876:
	;
	v7840 = F_pq_getmsgint(m, v7836, int32(4))
	mBase = m.M
	v7841 = m.ExcPending
	if v7841 != 0 {
		goto L1
	} else {
		goto L1877
	}
L1877:
	;
	F_pq_getmsgend(m, v7836)
	mBase = m.M
	v7843 = m.ExcPending
	if v7843 != 0 {
		goto L1
	} else {
		goto L1878
	}
L1878:
	;
	v7845 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v7847 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	v7848 = F_GetPortalByName(m, v7837)
	mBase = m.M
	v7849 = m.ExcPending
	if v7849 != 0 {
		goto L1
	} else {
		goto L1879
	}
L1879:
	;
	if v7848 == int32(0) {
		goto L189
	} else {
		goto L1880
	}
L1880:
	;
	if v7845 == int32(2) {
		goto L1881
	} else {
		goto L1882
	}
L1881:
	;
	v7855 = int32(3)
	goto L1883
L1882:
	;
	v7855 = v7845
	goto L1883
L1883:
	;
	v7856 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+36))
	if v7856 == int32(0) {
		goto L1884
	} else {
		goto L1885
	}
L1884:
	;
	F_NullCommand(m, v7855)
	mBase = m.M
	v7860 = m.ExcPending
	if v7860 != 0 {
		goto L1
	} else {
		goto L1887
	}
L1885:
	;
	goto L1886
L1886:
	;
	v7861 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+56))
	if v7861 == int32(0) {
		v7880 = v1
		goto L1888
	} else {
		goto L1889
	}
L1887:
	;
	goto L198
L1888:
	;
	v7881 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+32))
	v7882 = F_pstrdup(m, v7881)
	mBase = m.M
	v7883 = m.ExcPending
	if v7883 != 0 {
		goto L1
	} else {
		goto L1895
	}
L1889:
	;
	v7864 = *(*int32)(unsafe.Add(mBase, uint32(v7861)+4))
	if v7864 != int32(1) {
		v7880 = v1
		goto L1888
	} else {
		goto L1890
	}
L1890:
	;
	v7867 = *(*int32)(unsafe.Add(mBase, uint32(v7861)+12))
	v7868 = *(*int32)(unsafe.Add(mBase, uint32(v7867)))
	v7869 = *(*int32)(unsafe.Add(mBase, uint32(v7868)+4))
	if v7869 == int32(6) {
		goto L1891
	} else {
		goto L1892
	}
L1891:
	;
	v7873 = *(*int32)(unsafe.Add(mBase, uint32(v7868)+88))
	v7874 = *(*int32)(unsafe.Add(mBase, uint32(v7873)))
	if v7874 == int32(225) {
		v7880 = int32(1)
		goto L1888
	} else {
		goto L1894
	}
L1892:
	;
	goto L1893
L1893:
	;
	v7880 = int32(0)
	goto L1888
L1894:
	;
	goto L1893
L1895:
	;
	v7884 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+4))
	if v7884 != 0 {
		goto L1896
	} else {
		goto L1897
	}
L1896:
	;
	v7885 = F_pstrdup(m, v7884)
	mBase = m.M
	v7886 = m.ExcPending
	if v7886 != 0 {
		goto L1
	} else {
		goto L1899
	}
L1897:
	;
	v7888 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1898
L1898:
	;
	v7889 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v7882
	F_pgstat_report_activity(m, int32(3), v7882)
	mBase = m.M
	v7894 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+56))
	if v7894 == int32(0) {
		goto L1900
	} else {
		goto L1901
	}
L1899:
	;
	v7888 = v7885
	goto L1898
L1900:
	;
	v8131 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+36))
	v8136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8131<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[124]))))
	*(*int32)(unsafe.Add(mBase, uint32(v40+int32(456)))) = v8136
	goto L1934
L1901:
	;
	v7897 = *(*int32)(unsafe.Add(mBase, uint32(v7894)+4))
	if v7897 <= int32(0) {
		goto L1900
	} else {
		goto L1902
	}
L1902:
	;
	v7900 = int32(0)
	if v7900 < v7897 {
		goto L1903
	} else {
		goto L1904
	}
L1903:
	;
	v7903 = v7897
	goto L1905
L1904:
	;
	v7903 = v7900
	goto L1905
L1905:
	;
	v7904 = *(*int32)(unsafe.Add(mBase, uint32(v7894)+12))
	v7912 = int32(0)
	goto L1907
L1906:
	;
	if v8003 <= int32(0) {
		goto L1900
	} else {
		goto L1922
	}
L1907:
	;
	v7946 = *(*int32)(unsafe.Add(mBase, uint32(v7904+v7912<<(uint(int32(2))%32))))
	v7947 = *(*int64)(unsafe.Add(mBase, uint32(v7946)+8))
	if v7947 == int64(0) {
		goto L1909
	} else {
		goto L1910
	}
L1908:
	;
	v7956 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v7956 == int32(0) {
		goto L1914
	} else {
		goto L1915
	}
L1909:
	;
	v7951 = v7912 + int32(1)
	if v7951 != v7897 {
		v7912 = v7951
		goto L1907
	} else {
		goto L1912
	}
L1910:
	;
	goto L1911
L1911:
	;
	goto L1908
L1912:
	;
	v8000 = v7903
	v8001 = v7894
	v8003 = v7897
	goto L1906
L1913:
	;
	v7992 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+56))
	if v7992 == int32(0) {
		goto L1900
	} else {
		goto L1918
	}
L1914:
	;
	goto L1913
L1915:
	;
	v7960 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v7960&int32(1) == int32(0) {
		goto L1914
	} else {
		goto L1916
	}
L1916:
	;
	v7967 = *(*int64)(unsafe.Add(mBase, uint32(v7956)+392))
	if int32(1)&base.B2i32(v7967 != int64(0)) != 0 {
		goto L1914
	} else {
		goto L1917
	}
L1917:
	;
	v7971 = int32(_a_F_PostgresMainLoopOnce_176)
	v7973 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v7974 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v7973 + v7974
	v7977 = *(*int32)(unsafe.Add(mBase, uint32(v7956)))
	*(*int32)(unsafe.Add(mBase, uint32(v7956))) = v7977 + v7974
	*(*int64)(unsafe.Add(mBase, uint32(v7956)+392)) = v7947
	*(*int32)(unsafe.Add(mBase, uint32(v7956))) = v7977 + int32(2)
	v7988 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v7988 - v7974
	goto L1914
L1918:
	;
	v7995 = *(*int32)(unsafe.Add(mBase, uint32(v7992)+4))
	v7996 = int32(0)
	if v7996 < v7995 {
		goto L1919
	} else {
		goto L1920
	}
L1919:
	;
	v7999 = v7995
	goto L1921
L1920:
	;
	v7999 = v7996
	goto L1921
L1921:
	;
	v8000 = v7999
	v8001 = v7992
	v8003 = v7995
	goto L1906
L1922:
	;
	v8006 = *(*int32)(unsafe.Add(mBase, uint32(v8001)+12))
	v8014 = int32(0)
	goto L1923
L1923:
	;
	v8048 = *(*int32)(unsafe.Add(mBase, uint32(v8006+v8014<<(uint(int32(2))%32))))
	v8049 = *(*int64)(unsafe.Add(mBase, uint32(v8048)+16))
	if v8049 == int64(0) {
		goto L1925
	} else {
		goto L1926
	}
L1924:
	;
	v8058 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v8058 == int32(0) {
		goto L1930
	} else {
		goto L1931
	}
L1925:
	;
	v8053 = v8014 + int32(1)
	if v8000 != v8053 {
		v8014 = v8053
		goto L1923
	} else {
		goto L1928
	}
L1926:
	;
	goto L1927
L1927:
	;
	goto L1924
L1928:
	;
	goto L1900
L1929:
	;
	goto L1900
L1930:
	;
	goto L1929
L1931:
	;
	v8062 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v8062&int32(1) == int32(0) {
		goto L1930
	} else {
		goto L1932
	}
L1932:
	;
	v8069 = *(*int64)(unsafe.Add(mBase, uint32(v8058)+400))
	if int32(1)&base.B2i32(v8069 != int64(0)) != 0 {
		goto L1930
	} else {
		goto L1933
	}
L1933:
	;
	v8073 = int32(_a_F_PostgresMainLoopOnce_176)
	v8075 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	v8076 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v8075 + v8076
	v8079 = *(*int32)(unsafe.Add(mBase, uint32(v8058)))
	*(*int32)(unsafe.Add(mBase, uint32(v8058))) = v8079 + v8076
	*(*int64)(unsafe.Add(mBase, uint32(v8058)+400)) = v8049
	*(*int32)(unsafe.Add(mBase, uint32(v8058))) = v8079 + int32(2)
	v8090 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])) = v8090 - v8076
	goto L1930
L1934:
	;
	if v7847&int32(1) != 0 {
		goto L1935
	} else {
		goto L1936
	}
L1935:
	;
	v8142 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	v8152 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1938
L1936:
	;
	goto L1937
L1937:
	;
	v8156 = F_CreateDestReceiver(m, v7855)
	mBase = m.M
	v8157 = m.ExcPending
	if v8157 != 0 {
		goto L1
	} else {
		goto L1939
	}
L1938:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1937
L1939:
	;
	if v7855 == int32(3) {
		goto L1940
	} else {
		goto L1941
	}
L1940:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8156)+20)) = v7848
	goto L1942
L1941:
	;
	goto L1942
L1942:
	;
	F_start_xact_command(m)
	mBase = m.M
	v8162 = m.ExcPending
	if v8162 != 0 {
		goto L1
	} else {
		goto L1943
	}
L1943:
	;
	v8163 = int32(0)
	v8164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7848)+116)))
	v8166 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119]))
	switch v8166 {
	case 0:
		v8328 = v8163
		goto L1944
	default:
		goto L1946
	case 3:
		goto L1945
	}
L1944:
	;
	v8360 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8361 = *(*int32)(unsafe.Add(mBase, uint32(v8360)+24))
	goto L1978
L1945:
	;
	v8264 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8265 = m.ExcPending
	if v8265 != 0 {
		goto L1
	} else {
		goto L1954
	}
L1946:
	;
	v8167 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+56))
	if v8167 == int32(0) {
		v8328 = v8163
		goto L1944
	} else {
		goto L1947
	}
L1947:
	;
	v8170 = *(*int32)(unsafe.Add(mBase, uint32(v8167)+4))
	if v8170 <= int32(0) {
		v8328 = v8163
		goto L1944
	} else {
		goto L1948
	}
L1948:
	;
	v8179 = v8163
	goto L1949
L1949:
	;
	v8210 = *(*int32)(unsafe.Add(mBase, uint32(v8167)+12))
	v8214 = *(*int32)(unsafe.Add(mBase, uint32(v8210+v8179<<(uint(int32(2))%32))))
	v8215 = F_GetCommandLogLevel(m, v8214)
	mBase = m.M
	v8216 = m.ExcPending
	if v8216 != 0 {
		goto L1
	} else {
		goto L1951
	}
L1950:
	;
	v8328 = int32(0)
	goto L1944
L1951:
	;
	v8218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119]))
	if base.Ui32(v8215) <= base.Ui32(v8218) {
		goto L1945
	} else {
		goto L1952
	}
L1952:
	;
	v8221 = v8179 + int32(1)
	v8222 = *(*int32)(unsafe.Add(mBase, uint32(v8167)+4))
	if v8221 < v8222 {
		v8179 = v8221
		goto L1949
	} else {
		goto L1953
	}
L1953:
	;
	goto L1950
L1954:
	;
	if v8264 == int32(0) {
		goto L1955
	} else {
		goto L1956
	}
L1955:
	;
	v8328 = int32(1)
	goto L1944
L1956:
	;
	goto L1957
L1957:
	;
	v8269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7837))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+336)) = v7882
	*(*int32)(unsafe.Add(mBase, uint32(v40)+324)) = v7888
	if v8269 != 0 {
		goto L1958
	} else {
		goto L1959
	}
L1958:
	;
	v8273 = v7837
	goto L1960
L1959:
	;
	v8273 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1960
L1960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+332)) = v8273
	if v8269 != 0 {
		goto L1961
	} else {
		goto L1962
	}
L1961:
	;
	v8277 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L1963
L1962:
	;
	v8277 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1963
L1963:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+328)) = v8277
	v8279 = int32(1)
	if v8164&v8279 != 0 {
		goto L1964
	} else {
		goto L1965
	}
L1964:
	;
	v8284 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L1966
L1965:
	;
	v8284 = int32(_a_F_PostgresMainLoopOnce_197)
	goto L1966
L1966:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+320)) = v8284
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_198), v40+int32(320))
	mBase = m.M
	v8290 = m.ExcPending
	if v8290 != 0 {
		goto L1
	} else {
		goto L1967
	}
L1967:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8292 = m.ExcPending
	if v8292 != 0 {
		goto L1
	} else {
		goto L1968
	}
L1968:
	;
	if v7889 == int32(0) {
		goto L1969
	} else {
		goto L1970
	}
L1969:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2346), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v8321 = m.ExcPending
	if v8321 != 0 {
		goto L1
	} else {
		goto L1977
	}
L1970:
	;
	v8295 = *(*int32)(unsafe.Add(mBase, uint32(v7889)+28))
	if v8295 <= int32(0) {
		goto L1969
	} else {
		goto L1971
	}
L1971:
	;
	v8299 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128]))
	if v8299 == int32(0) {
		goto L1969
	} else {
		goto L1972
	}
L1972:
	;
	v8303 = F_BuildParamLogString(m, v7889, int32(0), v8299)
	mBase = m.M
	v8304 = m.ExcPending
	if v8304 != 0 {
		goto L1
	} else {
		goto L1973
	}
L1973:
	;
	if v8303 == int32(0) {
		goto L1969
	} else {
		goto L1974
	}
L1974:
	;
	v8307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8303))))
	if v8307 == int32(0) {
		goto L1969
	} else {
		goto L1975
	}
L1975:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+304)) = v8303
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(304))
	mBase = m.M
	v8315 = m.ExcPending
	if v8315 != 0 {
		goto L1
	} else {
		goto L1976
	}
L1976:
	;
	goto L1969
L1977:
	;
	v8328 = v8279
	goto L1944
L1978:
	;
	if (v8361-int32(7))&int32(-9) == int32(0) {
		goto L1979
	} else {
		goto L1980
	}
L1979:
	;
	v8368 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+56))
	if v8368 == int32(0) {
		goto L188
	} else {
		goto L1982
	}
L1980:
	;
	goto L1981
L1981:
	;
	v8392 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v8392 != 0 {
		goto L1988
	} else {
		goto L1989
	}
L1982:
	;
	v8371 = *(*int32)(unsafe.Add(mBase, uint32(v8368)+4))
	if v8371 != int32(1) {
		goto L188
	} else {
		goto L1983
	}
L1983:
	;
	v8374 = *(*int32)(unsafe.Add(mBase, uint32(v8368)+12))
	v8375 = *(*int32)(unsafe.Add(mBase, uint32(v8374)))
	v8376 = *(*int32)(unsafe.Add(mBase, uint32(v8375)+4))
	if v8376 != int32(6) {
		goto L188
	} else {
		goto L1984
	}
L1984:
	;
	v8379 = *(*int32)(unsafe.Add(mBase, uint32(v8375)+88))
	if v8379 == int32(0) {
		goto L188
	} else {
		goto L1985
	}
L1985:
	;
	v8382 = *(*int32)(unsafe.Add(mBase, uint32(v8379)))
	if v8382 != int32(225) {
		goto L188
	} else {
		goto L1986
	}
L1986:
	;
	v8385 = *(*int32)(unsafe.Add(mBase, uint32(v8379)+4))
	if (v8385-int32(2))&int32(-6) != 0 {
		goto L188
	} else {
		goto L1987
	}
L1987:
	;
	goto L1981
L1988:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v8394 = m.ExcPending
	if v8394 != 0 {
		goto L1
	} else {
		goto L1991
	}
L1989:
	;
	goto L1990
L1990:
	;
	v8395 = *(*int32)(unsafe.Add(mBase, uint32(v7848)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+476)) = v7889
	*(*int32)(unsafe.Add(mBase, uint32(v40)+472)) = v8395
	v8398 = int32(_a_F_PostgresMainLoopOnce_191)
	v8399 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v40 + int32(460)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = int32(1146)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v8399
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v40 + int32(472)
	if v7840 <= int32(0) {
		goto L1992
	} else {
		goto L1993
	}
L1991:
	;
	goto L1990
L1992:
	;
	v8413 = int32(2147483647)
	goto L1994
L1993:
	;
	v8413 = v7840
	goto L1994
L1994:
	;
	v8417 = F_PortalRun(m, v7848, v8413, int32(1), v8156, v8156, v40+int32(512))
	mBase = m.M
	v8418 = m.ExcPending
	if v8418 != 0 {
		goto L1
	} else {
		goto L1995
	}
L1995:
	;
	v8419 = *(*int32)(unsafe.Add(mBase, uint32(v8156)+12))
	m.T0[v8419].(func(*base.Module, int32))(m, v8156)
	mBase = m.M
	v8421 = m.ExcPending
	if v8421 != 0 {
		goto L1
	} else {
		goto L1996
	}
L1996:
	;
	v8422 = int32(_a_F_PostgresMainLoopOnce_191)
	v8424 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v8425 = *(*int32)(unsafe.Add(mBase, uint32(v8424)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v8425
	if v8417 != 0 {
		goto L1998
	} else {
		goto L1999
	}
L1997:
	;
	v8490 = F_check_log_duration(m, v40+int32(480), v8328)
	mBase = m.M
	v8491 = m.ExcPending
	if v8491 != 0 {
		goto L1
	} else {
		goto L2027
	}
L1998:
	;
	if v7880 == int32(0) {
		goto L2003
	} else {
		goto L2004
	}
L1999:
	;
	goto L2000
L2000:
	;
	v8475 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v8475 == int32(2) {
		goto L2019
	} else {
		goto L2020
	}
L2001:
	;
	F_EndCommand(m, v40+int32(512), v7855)
	mBase = m.M
	v8473 = m.ExcPending
	if v8473 != 0 {
		goto L1
	} else {
		goto L2018
	}
L2002:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v8453 = m.ExcPending
	if v8453 != 0 {
		goto L1
	} else {
		goto L2014
	}
L2003:
	;
	v8430 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v8430&int32(4) == int32(0) {
		goto L2002
	} else {
		goto L2006
	}
L2004:
	;
	goto L2005
L2005:
	;
	v8438 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L2007
L2006:
	;
	goto L2005
L2007:
	;
	if v8438 != 0 {
		goto L2008
	} else {
		goto L2009
	}
L2008:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8441 = m.ExcPending
	if v8441 != 0 {
		goto L1
	} else {
		goto L2011
	}
L2009:
	;
	goto L2010
L2010:
	;
	v8442 = int32(0)
	v8444 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v8444 == v8442 {
		v8469 = v8442
		goto L2001
	} else {
		goto L2012
	}
L2011:
	;
	goto L2010
L2012:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8448 = m.ExcPending
	if v8448 != 0 {
		goto L1
	} else {
		goto L2013
	}
L2013:
	;
	v8450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v8450)
	v8469 = v8442
	goto L2001
L2014:
	;
	v8454 = int32(_a_F_PostgresMainLoopOnce_200)
	v8456 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v8456 | int32(8)
	v8463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L2015
L2015:
	;
	if v8463 == int32(0) {
		v8469 = v7889
		goto L2001
	} else {
		goto L2016
	}
L2016:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8468 = m.ExcPending
	if v8468 != 0 {
		goto L1
	} else {
		goto L2017
	}
L2017:
	;
	v8469 = v7889
	goto L2001
L2018:
	;
	v8487 = v8469
	goto L1997
L2019:
	;
	F_pq_putemptymessage(m, int32(115))
	mBase = m.M
	v8480 = m.ExcPending
	if v8480 != 0 {
		goto L1
	} else {
		goto L2022
	}
L2020:
	;
	goto L2021
L2021:
	;
	v8481 = int32(_a_F_PostgresMainLoopOnce_200)
	v8483 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v8483 | int32(8)
	v8487 = v7889
	goto L1997
L2022:
	;
	goto L2021
L2023:
	;
	if v7847&int32(1) != 0 {
		goto L2053
	} else {
		goto L2054
	}
L2024:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v8568, int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v8572 = m.ExcPending
	if v8572 != 0 {
		goto L1
	} else {
		goto L2052
	}
L2025:
	;
	v8513 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8514 = m.ExcPending
	if v8514 != 0 {
		goto L1
	} else {
		goto L2032
	}
L2026:
	;
	v8496 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8497 = m.ExcPending
	if v8497 != 0 {
		goto L1
	} else {
		goto L2028
	}
L2027:
	;
	switch v8490 - int32(1) {
	case 0:
		goto L2026
	case 1:
		goto L2025
	default:
		goto L2023
	}
L2028:
	;
	if v8496 == int32(0) {
		goto L2023
	} else {
		goto L2029
	}
L2029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+240)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(240))
	mBase = m.M
	v8507 = m.ExcPending
	if v8507 != 0 {
		goto L1
	} else {
		goto L2030
	}
L2030:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8509 = m.ExcPending
	if v8509 != 0 {
		goto L1
	} else {
		goto L2031
	}
L2031:
	;
	v8568 = int32(2457)
	goto L2024
L2032:
	;
	if v8513 == int32(0) {
		goto L2023
	} else {
		goto L2033
	}
L2033:
	;
	v8517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7837))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+292)) = v7882
	if v8517 != 0 {
		goto L2034
	} else {
		goto L2035
	}
L2034:
	;
	v8520 = v7837
	goto L2036
L2035:
	;
	v8520 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L2036
L2036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+288)) = v8520
	*(*int32)(unsafe.Add(mBase, uint32(v40)+280)) = v7888
	if v8517 != 0 {
		goto L2037
	} else {
		goto L2038
	}
L2037:
	;
	v8525 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L2039
L2038:
	;
	v8525 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L2039
L2039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+284)) = v8525
	if v8164&int32(1) != 0 {
		goto L2040
	} else {
		goto L2041
	}
L2040:
	;
	v8531 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L2042
L2041:
	;
	v8531 = int32(_a_F_PostgresMainLoopOnce_197)
	goto L2042
L2042:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+276)) = v8531
	*(*int32)(unsafe.Add(mBase, uint32(v40)+272)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_201), v40+int32(272))
	mBase = m.M
	v8540 = m.ExcPending
	if v8540 != 0 {
		goto L1
	} else {
		goto L2043
	}
L2043:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8542 = m.ExcPending
	if v8542 != 0 {
		goto L1
	} else {
		goto L2044
	}
L2044:
	;
	v8543 = int32(2471)
	if v8487 == int32(0) {
		v8568 = v8543
		goto L2024
	} else {
		goto L2045
	}
L2045:
	;
	v8546 = *(*int32)(unsafe.Add(mBase, uint32(v8487)+28))
	if v8546 <= int32(0) {
		v8568 = v8543
		goto L2024
	} else {
		goto L2046
	}
L2046:
	;
	v8550 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128]))
	if v8550 == int32(0) {
		v8568 = v8543
		goto L2024
	} else {
		goto L2047
	}
L2047:
	;
	v8554 = F_BuildParamLogString(m, v8487, int32(0), v8550)
	mBase = m.M
	v8555 = m.ExcPending
	if v8555 != 0 {
		goto L1
	} else {
		goto L2048
	}
L2048:
	;
	if v8554 == int32(0) {
		v8568 = v8543
		goto L2024
	} else {
		goto L2049
	}
L2049:
	;
	v8558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8554))))
	if v8558 == int32(0) {
		v8568 = v8543
		goto L2024
	} else {
		goto L2050
	}
L2050:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+256)) = v8554
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(256))
	mBase = m.M
	v8566 = m.ExcPending
	if v8566 != 0 {
		goto L1
	} else {
		goto L2051
	}
L2051:
	;
	v8568 = v8543
	goto L2024
L2052:
	;
	goto L2023
L2053:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_202))
	mBase = m.M
	v8579 = m.ExcPending
	if v8579 != 0 {
		goto L1
	} else {
		goto L2056
	}
L2054:
	;
	goto L2055
L2055:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L198
L2056:
	;
	goto L2055
L2057:
	;
	v8588 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v8588 < int32(0) {
		goto L2059
	} else {
		goto L2060
	}
L2058:
	;
	F_pgstat_report_activity(m, int32(5), int32(0))
	mBase = m.M
	F_start_xact_command(m)
	mBase = m.M
	v8598 = m.ExcPending
	if v8598 != 0 {
		goto L1
	} else {
		goto L2062
	}
L2059:
	;
	v8592 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v8592
	goto L2061
L2060:
	;
	goto L2061
L2061:
	;
	goto L2058
L2062:
	;
	v8601 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v8601
	v8604 = v40 + int32(440)
	v8605 = m.G0
	v8607 = v8605 - int32(1568)
	m.G0 = v8607
	v8610 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8611 = *(*int32)(unsafe.Add(mBase, uint32(v8610)+24))
	goto L2073
L2063:
	;
	v9454 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v9454 != 0 {
		goto L2246
	} else {
		goto L2247
	}
L2064:
	;
	v9412 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+240))
	v9413 = m.T0[v9412].(func(*base.Module, int32) int32)(m, v8607+int32(740))
	mBase = m.M
	v9414 = m.ExcPending
	if v9414 != 0 {
		goto L1
	} else {
		goto L2245
	}
L2065:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9357 = m.ExcPending
	if v9357 != 0 {
		goto L1
	} else {
		goto L2241
	}
L2066:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9337 = m.ExcPending
	if v9337 != 0 {
		goto L1
	} else {
		goto L2237
	}
L2067:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9319 = m.ExcPending
	if v9319 != 0 {
		goto L1
	} else {
		goto L2233
	}
L2068:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9300 = m.ExcPending
	if v9300 != 0 {
		goto L1
	} else {
		goto L2229
	}
L2069:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9280 = m.ExcPending
	if v9280 != 0 {
		goto L1
	} else {
		goto L2225
	}
L2070:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9261 = m.ExcPending
	if v9261 != 0 {
		goto L1
	} else {
		goto L2222
	}
L2071:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9241 = m.ExcPending
	if v9241 != 0 {
		goto L1
	} else {
		goto L2218
	}
L2072:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9225 = m.ExcPending
	if v9225 != 0 {
		goto L1
	} else {
		goto L2214
	}
L2073:
	;
	if base.B2i32((v8611-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L2074
	} else {
		goto L2075
	}
L2074:
	;
	v8620 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v8621 = m.ExcPending
	if v8621 != 0 {
		goto L1
	} else {
		goto L2077
	}
L2075:
	;
	goto L2076
L2076:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9209 = m.ExcPending
	if v9209 != 0 {
		goto L1
	} else {
		goto L2210
	}
L2077:
	;
	F_PushActiveSnapshot(m, v8620)
	mBase = m.M
	v8623 = m.ExcPending
	if v8623 != 0 {
		goto L1
	} else {
		goto L2078
	}
L2078:
	;
	v8625 = F_pq_getmsgint(m, v8604, int32(4))
	mBase = m.M
	v8626 = m.ExcPending
	if v8626 != 0 {
		goto L1
	} else {
		goto L2079
	}
L2079:
	;
	base.MemoryFill(m, v8607+int32(236), int32(0), int32(504))
	v8633 = F_SearchSysCache1(m, int32(47), v8625)
	mBase = m.M
	v8634 = m.ExcPending
	if v8634 != 0 {
		goto L1
	} else {
		goto L2080
	}
L2080:
	;
	if v8633 == int32(0) {
		goto L2072
	} else {
		goto L2081
	}
L2081:
	;
	v8637 = *(*int32)(unsafe.Add(mBase, uint32(v8633)+16))
	v8638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8637)+22)))
	v8639 = v8637 + v8638
	v8640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8639)+96)))
	if v8640 != int32(102) {
		goto L2071
	} else {
		goto L2082
	}
L2082:
	;
	v8643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8639)+100)))
	if v8643 == int32(1) {
		goto L2071
	} else {
		goto L2083
	}
L2083:
	;
	v8646 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8639)+104)))
	if int32(101) <= v8646 {
		goto L2070
	} else {
		goto L2084
	}
L2084:
	;
	v8649 = *(*int32)(unsafe.Add(mBase, uint32(v8639)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+268)) = v8649
	v8651 = *(*int32)(unsafe.Add(mBase, uint32(v8639)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+272)) = v8651
	v8654 = v8607 + int32(276)
	v8655 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8639)+104)))
	v8657 = v8655 << (uint(int32(2)) % 32)
	if v8657 != 0 {
		goto L2085
	} else {
		goto L2086
	}
L2085:
	;
	base.MemoryCopy(m, v8654, v8639+int32(136), v8657)
	goto L2087
L2086:
	;
	goto L2087
L2087:
	;
	v8662 = v8607 + int32(240)
	v8664 = v8607 + int32(676)
	v8666 = v8639 + int32(4)
	goto L2091
L2088:
	;
	F_ReleaseCatCache(m, v8633)
	mBase = m.M
	v8787 = m.ExcPending
	if v8787 != 0 {
		goto L1
	} else {
		goto L2119
	}
L2089:
	;
	v8783 = F_strlen(m, v8772)
	mBase = m.M
	goto L2088
L2091:
	;
	goto L2092
L2092:
	;
	v8673 = int32(63)
	if (v8664^v8666)&int32(3) != 0 {
		goto L2096
	} else {
		goto L2097
	}
L2093:
	;
	v8776 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8773))) = uint8(v8776)
	goto L2089
L2094:
	;
	v8757 = v8752
	v8758 = v8753
	v8759 = v8754
	goto L2115
L2095:
	;
	if v8747 == int32(0) {
		v8772 = v8745
		v8773 = v8746
		goto L2093
	} else {
		goto L2114
	}
L2096:
	;
	v8745 = v8666
	v8746 = v8664
	v8747 = v8673
	goto L2095
L2097:
	;
	goto L2098
L2098:
	;
	v8677 = int32(0)
	if base.B2i32(v8666&int32(3) == v8677)|int32(0) == v8677 {
		goto L2100
	} else {
		goto L2101
	}
L2099:
	;
	if v8713 == int32(0) {
		v8772 = v8710
		v8773 = v8711
		goto L2093
	} else {
		goto L2108
	}
L2100:
	;
	v8689 = v8666
	v8690 = v8664
	v8691 = v8673
	goto L2103
L2101:
	;
	goto L2102
L2102:
	;
	v8710 = v8666
	v8711 = v8664
	v8712 = v8673
	v8713 = int32(1)
	goto L2099
L2103:
	;
	v8693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8689))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8690))) = uint8(v8693)
	if v8693 == int32(0) {
		v8752 = v8689
		v8753 = v8690
		v8754 = v8691
		goto L2094
	} else {
		goto L2105
	}
L2104:
	;
	v8710 = v8704
	v8711 = v8698
	v8712 = v8700
	v8713 = v8702
	goto L2099
L2105:
	;
	v8697 = int32(1)
	v8698 = v8690 + v8697
	v8700 = v8691 - v8697
	v8701 = int32(0)
	v8702 = base.B2i32(v8700 != v8701)
	v8704 = v8689 + v8697
	if v8704&int32(3) == v8701 {
		v8710 = v8704
		v8711 = v8698
		v8712 = v8700
		v8713 = v8702
		goto L2099
	} else {
		goto L2106
	}
L2106:
	;
	if v8700 != 0 {
		v8689 = v8704
		v8690 = v8698
		v8691 = v8700
		goto L2103
	} else {
		goto L2107
	}
L2107:
	;
	goto L2104
L2108:
	;
	v8716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8710))))
	if base.B2i32(v8716 == int32(0))|base.B2i32(base.Ui32(v8712) < base.Ui32(int32(4))) != 0 {
		v8745 = v8710
		v8746 = v8711
		v8747 = v8712
		goto L2095
	} else {
		goto L2109
	}
L2109:
	;
	v8723 = v8710
	v8724 = v8711
	v8725 = v8712
	goto L2110
L2110:
	;
	v8728 = *(*int32)(unsafe.Add(mBase, uint32(v8723)))
	v8731 = int32(-2139062144)
	if (int32(16843008)-v8728|v8728)&v8731 != v8731 {
		v8752 = v8723
		v8753 = v8724
		v8754 = v8725
		goto L2094
	} else {
		goto L2112
	}
L2111:
	;
	v8745 = v8739
	v8746 = v8737
	v8747 = v8741
	goto L2095
L2112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8724))) = v8728
	v8736 = int32(4)
	v8737 = v8724 + v8736
	v8739 = v8723 + v8736
	v8741 = v8725 - v8736
	if base.Ui32(int32(3)) < base.Ui32(v8741) {
		v8723 = v8739
		v8724 = v8737
		v8725 = v8741
		goto L2110
	} else {
		goto L2113
	}
L2113:
	;
	goto L2111
L2114:
	;
	v8752 = v8745
	v8753 = v8746
	v8754 = v8747
	goto L2094
L2115:
	;
	v8761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8757))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8758))) = uint8(v8761)
	if v8761 == int32(0) {
		v8772 = v8757
		v8773 = v8758
		goto L2093
	} else {
		goto L2117
	}
L2116:
	;
	v8772 = v8768
	v8773 = v8766
	goto L2093
L2117:
	;
	v8765 = int32(1)
	v8766 = v8758 + v8765
	v8768 = v8757 + v8765
	v8770 = v8759 - v8765
	if v8770 != 0 {
		v8757 = v8768
		v8758 = v8766
		v8759 = v8770
		goto L2115
	} else {
		goto L2118
	}
L2118:
	;
	goto L2116
L2119:
	;
	F_fmgr_info(m, v8625, v8662)
	mBase = m.M
	v8789 = m.ExcPending
	if v8789 != 0 {
		goto L1
	} else {
		goto L2120
	}
L2120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+236)) = v8625
	v8792 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119]))
	if v8792 != int32(3) {
		goto L2121
	} else {
		goto L2122
	}
L2121:
	;
	v8814 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+268))
	v8816 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v8818 = F_object_aclcheck(m, int32(2615), v8814, v8816, int64(256))
	mBase = m.M
	v8819 = m.ExcPending
	if v8819 != 0 {
		goto L1
	} else {
		goto L2127
	}
L2122:
	;
	v8797 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8798 = m.ExcPending
	if v8798 != 0 {
		goto L1
	} else {
		goto L2123
	}
L2123:
	;
	if v8797 == int32(0) {
		goto L2121
	} else {
		goto L2124
	}
L2124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+180)) = v8625
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+176)) = v8664
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_203), v8607+int32(176))
	mBase = m.M
	v8807 = m.ExcPending
	if v8807 != 0 {
		goto L1
	} else {
		goto L2125
	}
L2125:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(234), int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v8812 = m.ExcPending
	if v8812 != 0 {
		goto L1
	} else {
		goto L2126
	}
L2126:
	;
	goto L2121
L2127:
	;
	if v8818 != 0 {
		goto L2128
	} else {
		goto L2129
	}
L2128:
	;
	v8821 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+268))
	v8822 = F_get_namespace_name(m, v8821)
	mBase = m.M
	v8823 = m.ExcPending
	if v8823 != 0 {
		goto L1
	} else {
		goto L2131
	}
L2129:
	;
	goto L2130
L2130:
	;
	v8827 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v8827 != 0 {
		goto L2133
	} else {
		goto L2134
	}
L2131:
	;
	F_aclcheck_error(m, v8818, int32(36), v8822)
	mBase = m.M
	v8825 = m.ExcPending
	if v8825 != 0 {
		goto L1
	} else {
		goto L2132
	}
L2132:
	;
	goto L2130
L2133:
	;
	v8828 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+268))
	v8830 = F_RunNamespaceSearchHook(m, v8828, int32(1))
	mBase = m.M
	v8831 = m.ExcPending
	if v8831 != 0 {
		goto L1
	} else {
		goto L2136
	}
L2134:
	;
	goto L2135
L2135:
	;
	v8834 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v8836 = F_object_aclcheck(m, int32(1255), v8625, v8834, int64(128))
	mBase = m.M
	v8837 = m.ExcPending
	if v8837 != 0 {
		goto L1
	} else {
		goto L2137
	}
L2136:
	;
	goto L2135
L2137:
	;
	if v8836 != 0 {
		goto L2138
	} else {
		goto L2139
	}
L2138:
	;
	v8839 = F_get_func_name(m, v8625)
	mBase = m.M
	v8840 = m.ExcPending
	if v8840 != 0 {
		goto L1
	} else {
		goto L2141
	}
L2139:
	;
	goto L2140
L2140:
	;
	v8845 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v8845 != 0 {
		goto L2143
	} else {
		goto L2144
	}
L2141:
	;
	F_aclcheck_error(m, v8836, int32(19), v8839)
	mBase = m.M
	v8842 = m.ExcPending
	if v8842 != 0 {
		goto L1
	} else {
		goto L2142
	}
L2142:
	;
	goto L2140
L2143:
	;
	F_RunFunctionExecuteHook(m, v8625)
	mBase = m.M
	v8847 = m.ExcPending
	if v8847 != 0 {
		goto L1
	} else {
		goto L2146
	}
L2144:
	;
	goto L2145
L2145:
	;
	v8848 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8607)+744)) = v8848
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+740)) = v8662
	*(*int64)(unsafe.Add(mBase, uint32(v8607)+749)) = v8848
	v8854 = F_pq_getmsgint(m, v8604, int32(2))
	mBase = m.M
	v8855 = m.ExcPending
	if v8855 != 0 {
		goto L1
	} else {
		goto L2147
	}
L2146:
	;
	goto L2145
L2147:
	;
	if int32(0) < v8854 {
		goto L2148
	} else {
		goto L2149
	}
L2148:
	;
	v8860 = F_palloc(m, v8854<<(uint(int32(1))%32))
	mBase = m.M
	v8861 = m.ExcPending
	if v8861 != 0 {
		goto L1
	} else {
		goto L2151
	}
L2149:
	;
	v8909 = int32(0)
	goto L2150
L2150:
	;
	v8947 = F_pq_getmsgint(m, v8604, int32(2))
	mBase = m.M
	v8948 = m.ExcPending
	if v8948 != 0 {
		goto L1
	} else {
		goto L2156
	}
L2151:
	;
	v8872 = v1
	goto L2152
L2152:
	;
	v8903 = F_pq_getmsgint(m, v8604, int32(2))
	mBase = m.M
	v8904 = m.ExcPending
	if v8904 != 0 {
		goto L1
	} else {
		goto L2154
	}
L2153:
	;
	v8909 = v8860
	goto L2150
L2154:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8860+v8872<<(uint(int32(1))%32)))) = uint16(v8903)
	v8907 = v8872 + int32(1)
	if v8907 != v8854 {
		v8872 = v8907
		goto L2152
	} else {
		goto L2155
	}
L2155:
	;
	goto L2153
L2156:
	;
	if int32(100) < v8947 {
		goto L2069
	} else {
		goto L2157
	}
L2157:
	;
	v8951 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8607)+248)))
	if v8947 != v8951 {
		goto L2069
	} else {
		goto L2158
	}
L2158:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8607)+758)) = uint16(v8947)
	if base.B2i32(v8947 != v8854)&base.B2i32(int32(2) <= v8854) != 0 {
		goto L2068
	} else {
		goto L2159
	}
L2159:
	;
	F_initStringInfo(m, v8607+int32(192))
	mBase = m.M
	v8961 = m.ExcPending
	if v8961 != 0 {
		goto L1
	} else {
		goto L2160
	}
L2160:
	;
	if int32(0) < v8947 {
		goto L2161
	} else {
		goto L2162
	}
L2161:
	;
	v8965 = v8607 + int32(760)
	v8979 = int32(0)
	goto L2164
L2162:
	;
	goto L2163
L2163:
	;
	v9144 = F_pq_getmsgint(m, v8604, int32(2))
	mBase = m.M
	v9145 = m.ExcPending
	if v9145 != 0 {
		goto L1
	} else {
		goto L2200
	}
L2164:
	;
	v9007 = v8979 << (uint(int32(3)) % 32)
	v9010 = v9007 + (v8607 + int32(740))
	v9012 = F_pq_getmsgint(m, v8604, int32(4))
	mBase = m.M
	v9013 = m.ExcPending
	if v9013 != 0 {
		goto L1
	} else {
		goto L2167
	}
L2165:
	;
	goto L2163
L2166:
	;
	if base.B2i32(v8854 < int32(2)) == int32(0) {
		goto L2179
	} else {
		goto L2180
	}
L2167:
	;
	if v9012 == int32(-1) {
		goto L2168
	} else {
		goto L2169
	}
L2168:
	;
	v9016 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9010)+24)) = uint8(v9016)
	goto L2166
L2169:
	;
	goto L2170
L2170:
	;
	v9018 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9010)+24)) = uint8(v9018)
	if v9012 < v9018 {
		goto L2067
	} else {
		goto L2171
	}
L2171:
	;
	v9023 = v8607 + int32(192)
	v9024 = *(*int32)(unsafe.Add(mBase, uint32(v9023)))
	v9025 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9024))) = uint8(v9025)
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+12)) = v9025
	*(*int32)(unsafe.Add(mBase, uint32(v9023)+4)) = v9025
	goto L2172
L2172:
	;
	v9031 = F_pq_getmsgbytes(m, v8604, v9012)
	mBase = m.M
	v9032 = m.ExcPending
	if v9032 != 0 {
		goto L1
	} else {
		goto L2173
	}
L2173:
	;
	F_appendBinaryStringInfo(m, v9023, v9031, v9012)
	mBase = m.M
	v9034 = m.ExcPending
	if v9034 != 0 {
		goto L1
	} else {
		goto L2174
	}
L2174:
	;
	goto L2166
L2175:
	;
	v9104 = v8979 + int32(1)
	if v9104 != v8947 {
		v8979 = v9104
		goto L2164
	} else {
		goto L2199
	}
L2176:
	;
	v9079 = *(*int32)(unsafe.Add(mBase, uint32(v8654+v8979<<(uint(int32(2))%32))))
	F_getTypeBinaryInputInfo(m, v9079, v8607+int32(1564), v8607+int32(1560))
	mBase = m.M
	v9085 = m.ExcPending
	if v9085 != 0 {
		goto L1
	} else {
		goto L2192
	}
L2177:
	;
	v9049 = *(*int32)(unsafe.Add(mBase, uint32(v8654+v8979<<(uint(int32(2))%32))))
	F_getTypeInputInfo(m, v9049, v8607+int32(1564), v8607+int32(1560))
	mBase = m.M
	v9055 = m.ExcPending
	if v9055 != 0 {
		goto L1
	} else {
		goto L2183
	}
L2178:
	;
	v9044 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9043))))
	switch v9044 {
	case 0:
		goto L2177
	case 1:
		goto L2176
	default:
		goto L2065
	}
L2179:
	;
	v9043 = v8909 + v8979<<(uint(int32(1))%32)
	goto L2178
L2180:
	;
	goto L2181
L2181:
	;
	if v8854 <= int32(0) {
		goto L2177
	} else {
		goto L2182
	}
L2182:
	;
	v9043 = v8909
	goto L2178
L2183:
	;
	if v9012 == int32(-1) {
		goto L2184
	} else {
		goto L2185
	}
L2184:
	;
	v9062 = int32(0)
	goto L2186
L2185:
	;
	v9059 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+192))
	v9060 = F_pg_client_to_server(m, v9059, v9012)
	mBase = m.M
	v9061 = m.ExcPending
	if v9061 != 0 {
		goto L1
	} else {
		goto L2187
	}
L2186:
	;
	v9064 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1564))
	v9065 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1560))
	v9067 = F_OidInputFunctionCall(m, v9064, v9062, v9065, int32(-1))
	mBase = m.M
	v9068 = m.ExcPending
	if v9068 != 0 {
		goto L1
	} else {
		goto L2188
	}
L2187:
	;
	v9062 = v9060
	goto L2186
L2188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8965+v9007))) = v9067
	if v9062 == int32(0) {
		goto L2175
	} else {
		goto L2189
	}
L2189:
	;
	v9072 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+192))
	if v9062 == v9072 {
		goto L2175
	} else {
		goto L2190
	}
L2190:
	;
	F_pfree(m, v9062)
	mBase = m.M
	v9075 = m.ExcPending
	if v9075 != 0 {
		goto L1
	} else {
		goto L2191
	}
L2191:
	;
	goto L2175
L2192:
	;
	v9087 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1564))
	v9092 = base.B2i32(v9012 == int32(-1))
	if v9012 == int32(-1) {
		goto L2193
	} else {
		goto L2194
	}
L2193:
	;
	v9093 = int32(0)
	goto L2195
L2194:
	;
	v9093 = v8607 + int32(192)
	goto L2195
L2195:
	;
	v9094 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1560))
	v9096 = F_OidReceiveFunctionCall(m, v9087, v9093, v9094, int32(-1))
	mBase = m.M
	v9097 = m.ExcPending
	if v9097 != 0 {
		goto L1
	} else {
		goto L2196
	}
L2196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8965+v9007))) = v9096
	if v9012 == int32(-1) {
		goto L2175
	} else {
		goto L2197
	}
L2197:
	;
	v9099 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+204))
	v9100 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+196))
	if v9099 != v9100 {
		goto L2066
	} else {
		goto L2198
	}
L2198:
	;
	goto L2175
L2199:
	;
	goto L2165
L2200:
	;
	F_pq_getmsgend(m, v8604)
	mBase = m.M
	v9147 = m.ExcPending
	if v9147 != 0 {
		goto L1
	} else {
		goto L2201
	}
L2201:
	;
	v9148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8607)+250)))
	if v9148 != int32(1) {
		goto L2064
	} else {
		goto L2202
	}
L2202:
	;
	v9151 = int32(0)
	if v8947 <= v9151 {
		goto L2064
	} else {
		goto L2203
	}
L2203:
	;
	v9165 = v9151
	goto L2204
L2204:
	;
	v9197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8607+int32(740)+v9165<<(uint(int32(3))%32))+24)))
	if v9197 == int32(0) {
		goto L2206
	} else {
		goto L2207
	}
L2205:
	;
	v9203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8607)+756)) = uint8(v9203)
	v9452 = int32(0)
	goto L2063
L2206:
	;
	v9201 = v9165 + int32(1)
	if base.I32_extend16_s(v8947) != v9201 {
		v9165 = v9201
		goto L2204
	} else {
		goto L2209
	}
L2207:
	;
	goto L2208
L2208:
	;
	goto L2205
L2209:
	;
	goto L2064
L2210:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v9212 = m.ExcPending
	if v9212 != 0 {
		goto L1
	} else {
		goto L2211
	}
L2211:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v9216 = m.ExcPending
	if v9216 != 0 {
		goto L1
	} else {
		goto L2212
	}
L2212:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(209), int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v9221 = m.ExcPending
	if v9221 != 0 {
		goto L1
	} else {
		goto L2213
	}
L2213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2214:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v9228 = m.ExcPending
	if v9228 != 0 {
		goto L1
	} else {
		goto L2215
	}
L2215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607))) = v8625
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_206), v8607)
	mBase = m.M
	v9232 = m.ExcPending
	if v9232 != 0 {
		goto L1
	} else {
		goto L2216
	}
L2216:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(141), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9237 = m.ExcPending
	if v9237 != 0 {
		goto L1
	} else {
		goto L2217
	}
L2217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2218:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9244 = m.ExcPending
	if v9244 != 0 {
		goto L1
	} else {
		goto L2219
	}
L2219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+16)) = v8639 + int32(4)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_208), v8607+int32(16))
	mBase = m.M
	v9252 = m.ExcPending
	if v9252 != 0 {
		goto L1
	} else {
		goto L2220
	}
L2220:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(149), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9257 = m.ExcPending
	if v9257 != 0 {
		goto L1
	} else {
		goto L2221
	}
L2221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+32)) = v8639 + int32(4)
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_209), v8607+int32(32))
	mBase = m.M
	v9271 = m.ExcPending
	if v9271 != 0 {
		goto L1
	} else {
		goto L2223
	}
L2223:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(154), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9276 = m.ExcPending
	if v9276 != 0 {
		goto L1
	} else {
		goto L2224
	}
L2224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2225:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9283 = m.ExcPending
	if v9283 != 0 {
		goto L1
	} else {
		goto L2226
	}
L2226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+48)) = v8947
	v9285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8607)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+52)) = v9285
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_210), v8607+int32(48))
	mBase = m.M
	v9291 = m.ExcPending
	if v9291 != 0 {
		goto L1
	} else {
		goto L2227
	}
L2227:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(353), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9296 = m.ExcPending
	if v9296 != 0 {
		goto L1
	} else {
		goto L2228
	}
L2228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2229:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9303 = m.ExcPending
	if v9303 != 0 {
		goto L1
	} else {
		goto L2230
	}
L2230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+164)) = v8947
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+160)) = v8854
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_212), v8607+int32(160))
	mBase = m.M
	v9310 = m.ExcPending
	if v9310 != 0 {
		goto L1
	} else {
		goto L2231
	}
L2231:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(361), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9315 = m.ExcPending
	if v9315 != 0 {
		goto L1
	} else {
		goto L2232
	}
L2232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2233:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9322 = m.ExcPending
	if v9322 != 0 {
		goto L1
	} else {
		goto L2234
	}
L2234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+144)) = v9012
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_213), v8607+int32(144))
	mBase = m.M
	v9328 = m.ExcPending
	if v9328 != 0 {
		goto L1
	} else {
		goto L2235
	}
L2235:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(385), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9333 = m.ExcPending
	if v9333 != 0 {
		goto L1
	} else {
		goto L2236
	}
L2236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2237:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v9340 = m.ExcPending
	if v9340 != 0 {
		goto L1
	} else {
		goto L2238
	}
L2238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+128)) = v8979 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_214), v8607+int32(128))
	mBase = m.M
	v9348 = m.ExcPending
	if v9348 != 0 {
		goto L1
	} else {
		goto L2239
	}
L2239:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(448), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9353 = m.ExcPending
	if v9353 != 0 {
		goto L1
	} else {
		goto L2240
	}
L2240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2241:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9360 = m.ExcPending
	if v9360 != 0 {
		goto L1
	} else {
		goto L2242
	}
L2242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+64)) = base.I32_extend16_s(v9044)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v8607-int32(-64))
	mBase = m.M
	v9367 = m.ExcPending
	if v9367 != 0 {
		goto L1
	} else {
		goto L2243
	}
L2243:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(453), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9372 = m.ExcPending
	if v9372 != 0 {
		goto L1
	} else {
		goto L2244
	}
L2244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2245:
	;
	v9452 = v9413
	goto L2063
L2246:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9456 = m.ExcPending
	if v9456 != 0 {
		goto L1
	} else {
		goto L2249
	}
L2247:
	;
	goto L2248
L2248:
	;
	v9457 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+272))
	v9458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8607)+756)))
	v9460 = v8607 + int32(192)
	F_pq_beginmessage(m, v9460, int32(86))
	mBase = m.M
	v9463 = m.ExcPending
	if v9463 != 0 {
		goto L1
	} else {
		goto L2250
	}
L2249:
	;
	goto L2248
L2250:
	;
	if v9458 == int32(1) {
		goto L2252
	} else {
		goto L2253
	}
L2251:
	;
	v9615 = v8607 + int32(192)
	F_pq_endmessage(m, v9615)
	mBase = m.M
	v9617 = m.ExcPending
	if v9617 != 0 {
		goto L1
	} else {
		goto L2284
	}
L2252:
	;
	F_enlargeStringInfo(m, v9460, int32(4))
	mBase = m.M
	v9468 = m.ExcPending
	if v9468 != 0 {
		goto L1
	} else {
		goto L2255
	}
L2253:
	;
	goto L2254
L2254:
	;
	switch v9144 & int32(_a_F_PostgresMainLoopOnce_216) {
	case 0:
		goto L2256
	case 1:
		goto L2258
	default:
		goto L2257
	}
L2255:
	;
	v9469 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+196))
	v9470 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v9469+v9470))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+196)) = v9469 + int32(4)
	goto L2251
L2256:
	;
	F_getTypeOutputInfo(m, v9457, v8607+int32(1564), v8607+int32(1560))
	mBase = m.M
	v9599 = m.ExcPending
	if v9599 != 0 {
		goto L1
	} else {
		goto L2280
	}
L2257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9578 = m.ExcPending
	if v9578 != 0 {
		goto L1
	} else {
		goto L2276
	}
L2258:
	;
	F_getTypeBinaryOutputInfo(m, v9457, v8607+int32(1564), v8607+int32(1560))
	mBase = m.M
	v9484 = m.ExcPending
	if v9484 != 0 {
		goto L1
	} else {
		goto L2259
	}
L2259:
	;
	v9485 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1564))
	v9486 = m.G0
	v9488 = v9486 + int32(-64)
	m.G0 = v9488
	v9491 = v9486 + int32(-56)
	v9493 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_fmgr_info_cxt_security(m, v9485, v9491, v9493, int32(0))
	mBase = m.M
	v9496 = m.ExcPending
	if v9496 != 0 {
		goto L1
	} else {
		goto L2261
	}
L2260:
	;
	v9538 = *(*int32)(unsafe.Add(mBase, uint32(v9520)))
	v9540 = v8607 + int32(192)
	F_enlargeStringInfo(m, v9540, int32(4))
	mBase = m.M
	v9543 = m.ExcPending
	if v9543 != 0 {
		goto L1
	} else {
		goto L2273
	}
L2261:
	;
	v9497 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9488)+40)) = v9497
	*(*int64)(unsafe.Add(mBase, uint32(v9488)+45)) = v9497
	v9501 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9488)+60)) = uint8(v9501)
	*(*int32)(unsafe.Add(mBase, uint32(v9488)+56)) = v9452
	*(*int32)(unsafe.Add(mBase, uint32(v9488)+36)) = v9491
	v9505 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v9488)+54)) = uint16(v9505)
	v9509 = *(*int32)(unsafe.Add(mBase, uint32(v9488)+8))
	v9510 = m.T0[v9509].(func(*base.Module, int32) int32)(m, v9486+int32(-28))
	mBase = m.M
	v9511 = m.ExcPending
	if v9511 != 0 {
		goto L1
	} else {
		goto L2262
	}
L2262:
	;
	v9512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9488)+52)))
	if v9512 != int32(1) {
		goto L2263
	} else {
		goto L2264
	}
L2263:
	;
	v9515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9510))))
	if v9515&int32(3) != 0 {
		goto L2266
	} else {
		goto L2267
	}
L2264:
	;
	goto L2265
L2265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9527 = m.ExcPending
	if v9527 != 0 {
		goto L1
	} else {
		goto L2270
	}
L2266:
	;
	v9518 = F_detoast_attr(m, v9510)
	mBase = m.M
	v9519 = m.ExcPending
	if v9519 != 0 {
		goto L1
	} else {
		goto L2269
	}
L2267:
	;
	v9520 = v9510
	goto L2268
L2268:
	;
	m.G0 = v9488 - int32(-64)
	goto L2260
L2269:
	;
	v9520 = v9518
	goto L2268
L2270:
	;
	v9528 = *(*int32)(unsafe.Add(mBase, uint32(v9488)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9488))) = v9528
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_217), v9488)
	mBase = m.M
	v9532 = m.ExcPending
	if v9532 != 0 {
		goto L1
	} else {
		goto L2271
	}
L2271:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_218), int32(1143), int32(_a_F_PostgresMainLoopOnce_219))
	mBase = m.M
	v9537 = m.ExcPending
	if v9537 != 0 {
		goto L1
	} else {
		goto L2272
	}
L2272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2273:
	;
	v9544 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+196))
	v9545 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+192))
	v9547 = int32(2)
	v9549 = int32(4)
	v9550 = int32(base.Ui32(v9538)>>(uint(v9547)%32)) - v9549
	v9551 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v9544+v9545))) = base.I32_rotr(v9550&v9551, int32(8)) | base.I32_rotr(v9550, int32(24))&v9551
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+196)) = v9544 + v9549
	v9566 = *(*int32)(unsafe.Add(mBase, uint32(v9520)))
	F_appendBinaryStringInfo(m, v9540, v9520+v9549, int32(base.Ui32(v9566)>>(uint(v9547)%32))-v9549)
	mBase = m.M
	v9572 = m.ExcPending
	if v9572 != 0 {
		goto L1
	} else {
		goto L2274
	}
L2274:
	;
	F_pfree(m, v9520)
	mBase = m.M
	v9574 = m.ExcPending
	if v9574 != 0 {
		goto L1
	} else {
		goto L2275
	}
L2275:
	;
	goto L2251
L2276:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9581 = m.ExcPending
	if v9581 != 0 {
		goto L1
	} else {
		goto L2277
	}
L2277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+112)) = base.I32_extend16_s(v9144)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v8607+int32(112))
	mBase = m.M
	v9588 = m.ExcPending
	if v9588 != 0 {
		goto L1
	} else {
		goto L2278
	}
L2278:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(106), int32(_a_F_PostgresMainLoopOnce_220))
	mBase = m.M
	v9593 = m.ExcPending
	if v9593 != 0 {
		goto L1
	} else {
		goto L2279
	}
L2279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2280:
	;
	v9602 = *(*int32)(unsafe.Add(mBase, uint32(v8607)+1564))
	v9603 = F_OidOutputFunctionCall(m, v9602, v9452)
	mBase = m.M
	v9604 = m.ExcPending
	if v9604 != 0 {
		goto L1
	} else {
		goto L2281
	}
L2281:
	;
	v9605 = F_strlen(m, v9603)
	mBase = m.M
	F_pq_sendcountedtext(m, v8607+int32(192), v9603, v9605)
	mBase = m.M
	v9607 = m.ExcPending
	if v9607 != 0 {
		goto L1
	} else {
		goto L2282
	}
L2282:
	;
	F_pfree(m, v9603)
	mBase = m.M
	v9609 = m.ExcPending
	if v9609 != 0 {
		goto L1
	} else {
		goto L2283
	}
L2283:
	;
	goto L2251
L2284:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v9619 = m.ExcPending
	if v9619 != 0 {
		goto L1
	} else {
		goto L2285
	}
L2285:
	;
	v9623 = F_check_log_duration(m, v9615, base.B2i32(v8792 == int32(3)))
	mBase = m.M
	v9624 = m.ExcPending
	if v9624 != 0 {
		goto L1
	} else {
		goto L2290
	}
L2286:
	;
	m.G0 = v8607 + int32(1568)
	v9669 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L2298
L2287:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), v9659, int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v9662 = m.ExcPending
	if v9662 != 0 {
		goto L1
	} else {
		goto L2297
	}
L2288:
	;
	v9644 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9645 = m.ExcPending
	if v9645 != 0 {
		goto L1
	} else {
		goto L2294
	}
L2289:
	;
	v9629 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9630 = m.ExcPending
	if v9630 != 0 {
		goto L1
	} else {
		goto L2291
	}
L2290:
	;
	switch v9623 - int32(1) {
	case 0:
		goto L2289
	case 1:
		goto L2288
	default:
		goto L2286
	}
L2291:
	;
	if v9629 == int32(0) {
		goto L2286
	} else {
		goto L2292
	}
L2292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+80)) = v8607 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v8607+int32(80))
	mBase = m.M
	v9640 = m.ExcPending
	if v9640 != 0 {
		goto L1
	} else {
		goto L2293
	}
L2293:
	;
	v9659 = int32(312)
	goto L2287
L2294:
	;
	if v9644 == int32(0) {
		goto L2286
	} else {
		goto L2295
	}
L2295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+104)) = v8625
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+100)) = v8664
	*(*int32)(unsafe.Add(mBase, uint32(v8607)+96)) = v8607 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_221), v8607+int32(96))
	mBase = m.M
	v9657 = m.ExcPending
	if v9657 != 0 {
		goto L1
	} else {
		goto L2296
	}
L2296:
	;
	v9659 = int32(317)
	goto L2287
L2297:
	;
	goto L2286
L2298:
	;
	if v9669 != 0 {
		goto L2299
	} else {
		goto L2300
	}
L2299:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9672 = m.ExcPending
	if v9672 != 0 {
		goto L1
	} else {
		goto L2302
	}
L2300:
	;
	goto L2301
L2301:
	;
	v9674 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v9674 != 0 {
		goto L2303
	} else {
		goto L2304
	}
L2302:
	;
	goto L2301
L2303:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v9676 = m.ExcPending
	if v9676 != 0 {
		goto L1
	} else {
		goto L2306
	}
L2304:
	;
	goto L2305
L2305:
	;
	v9681 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v9681)
	goto L198
L2306:
	;
	v9678 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v9678)
	goto L2305
L2307:
	;
	v9688 = v40 + int32(440)
	v9689 = F_pq_getmsgbyte(m, v9688)
	mBase = m.M
	v9690 = m.ExcPending
	if v9690 != 0 {
		goto L1
	} else {
		goto L2308
	}
L2308:
	;
	v9691 = F_pq_getmsgstring(m, v9688)
	mBase = m.M
	v9692 = m.ExcPending
	if v9692 != 0 {
		goto L1
	} else {
		goto L2309
	}
L2309:
	;
	F_pq_getmsgend(m, v9688)
	mBase = m.M
	v9694 = m.ExcPending
	if v9694 != 0 {
		goto L1
	} else {
		goto L2310
	}
L2310:
	;
	switch v9689 - int32(80) {
	case 0:
		goto L2312
	default:
		goto L186
	case 3:
		goto L2313
	}
L2311:
	;
	v9719 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9719 != int32(2) {
		goto L198
	} else {
		goto L2323
	}
L2312:
	;
	v9710 = F_GetPortalByName(m, v9691)
	mBase = m.M
	v9711 = m.ExcPending
	if v9711 != 0 {
		goto L1
	} else {
		goto L2320
	}
L2313:
	;
	v9697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9691))))
	if v9697 != 0 {
		goto L2314
	} else {
		goto L2315
	}
L2314:
	;
	F_DropPreparedStatement(m, v9691, int32(0))
	mBase = m.M
	v9700 = m.ExcPending
	if v9700 != 0 {
		goto L1
	} else {
		goto L2317
	}
L2315:
	;
	goto L2316
L2316:
	;
	v9702 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117]))
	if v9702 == int32(0) {
		goto L2311
	} else {
		goto L2318
	}
L2317:
	;
	goto L2311
L2318:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int32(0)
	F_DropCachedPlan(m, v9702)
	mBase = m.M
	v9709 = m.ExcPending
	if v9709 != 0 {
		goto L1
	} else {
		goto L2319
	}
L2319:
	;
	goto L2311
L2320:
	;
	if v9710 == int32(0) {
		goto L2311
	} else {
		goto L2321
	}
L2321:
	;
	F_PortalDrop(m, v9710, int32(0))
	mBase = m.M
	v9716 = m.ExcPending
	if v9716 != 0 {
		goto L1
	} else {
		goto L2322
	}
L2322:
	;
	goto L2311
L2323:
	;
	F_pq_putemptymessage(m, int32(51))
	mBase = m.M
	v9724 = m.ExcPending
	if v9724 != 0 {
		goto L1
	} else {
		goto L2324
	}
L2324:
	;
	goto L198
L2325:
	;
	v9730 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v9730 < int32(0) {
		goto L2327
	} else {
		goto L2328
	}
L2326:
	;
	v9737 = v40 + int32(440)
	v9738 = F_pq_getmsgbyte(m, v9737)
	mBase = m.M
	v9739 = m.ExcPending
	if v9739 != 0 {
		goto L1
	} else {
		goto L2330
	}
L2327:
	;
	v9734 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v9734
	goto L2329
L2328:
	;
	goto L2329
L2329:
	;
	goto L2326
L2330:
	;
	v9740 = F_pq_getmsgstring(m, v9737)
	mBase = m.M
	v9741 = m.ExcPending
	if v9741 != 0 {
		goto L1
	} else {
		goto L2331
	}
L2331:
	;
	F_pq_getmsgend(m, v9737)
	mBase = m.M
	v9743 = m.ExcPending
	if v9743 != 0 {
		goto L1
	} else {
		goto L2332
	}
L2332:
	;
	switch v9738 - int32(80) {
	case 0:
		goto L2334
	default:
		goto L2333
	case 3:
		goto L2335
	}
L2333:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9962 = m.ExcPending
	if v9962 != 0 {
		goto L1
	} else {
		goto L2377
	}
L2334:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9926 = m.ExcPending
	if v9926 != 0 {
		goto L1
	} else {
		goto L2362
	}
L2335:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9747 = m.ExcPending
	if v9747 != 0 {
		goto L1
	} else {
		goto L2336
	}
L2336:
	;
	v9750 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v9750
	v9752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9740))))
	if v9752 != 0 {
		goto L2338
	} else {
		goto L2339
	}
L2337:
	;
	v9763 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v9764 = *(*int32)(unsafe.Add(mBase, uint32(v9763)+24))
	goto L2343
L2338:
	;
	v9754 = F_FetchPreparedStatement(m, v9740, int32(1))
	mBase = m.M
	v9755 = m.ExcPending
	if v9755 != 0 {
		goto L1
	} else {
		goto L2341
	}
L2339:
	;
	goto L2340
L2340:
	;
	v9758 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117]))
	if v9758 == int32(0) {
		goto L185
	} else {
		goto L2342
	}
L2341:
	;
	v9756 = *(*int32)(unsafe.Add(mBase, uint32(v9754)+64))
	v9761 = v9756
	goto L2337
L2342:
	;
	v9761 = v9758
	goto L2337
L2343:
	;
	if (v9764-int32(7))&int32(-9) == int32(0) {
		goto L2344
	} else {
		goto L2345
	}
L2344:
	;
	v9771 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+52))
	if v9771 != 0 {
		goto L184
	} else {
		goto L2347
	}
L2345:
	;
	goto L2346
L2346:
	;
	v9773 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9773 != int32(2) {
		goto L198
	} else {
		goto L2348
	}
L2347:
	;
	goto L2346
L2348:
	;
	v9776 = int32(_a_F_PostgresMainLoopOnce_222)
	F_resetStringInfo(m, v9776)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132])) = int32(116)
	goto L2349
L2349:
	;
	v9780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9761)+24)))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_222), int32(2))
	mBase = m.M
	v9784 = m.ExcPending
	if v9784 != 0 {
		goto L1
	} else {
		goto L2350
	}
L2350:
	;
	v9786 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	v9787 = int32(_a_F_PostgresMainLoopOnce_223)
	v9788 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	v9790 = int32(8)
	v9794 = v9780<<(uint(v9790)%32) | int32(base.Ui32(v9780)>>(uint(v9790)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v9786+v9788))) = uint16(v9794)
	v9798 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134])) = v9798 + int32(2)
	v9802 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+24))
	if int32(0) < v9802 {
		goto L2351
	} else {
		goto L2352
	}
L2351:
	;
	v9812 = int32(0)
	goto L2354
L2352:
	;
	goto L2353
L2353:
	;
	F_pq_endmessage_reuse(m, int32(_a_F_PostgresMainLoopOnce_222))
	mBase = m.M
	v9914 = m.ExcPending
	if v9914 != 0 {
		goto L1
	} else {
		goto L2358
	}
L2354:
	;
	v9843 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+20))
	v9847 = *(*int32)(unsafe.Add(mBase, uint32(v9843+v9812<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_222), int32(4))
	mBase = m.M
	v9851 = m.ExcPending
	if v9851 != 0 {
		goto L1
	} else {
		goto L2356
	}
L2355:
	;
	goto L2353
L2356:
	;
	v9852 = int32(_a_F_PostgresMainLoopOnce_223)
	v9853 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	v9855 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	v9859 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v9853+v9855))) = base.I32_rotr(v9847, int32(24))&v9859 | base.I32_rotr(v9847&v9859, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134])) = v9853 + int32(4)
	v9872 = v9812 + int32(1)
	v9873 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+24))
	if v9872 < v9873 {
		v9812 = v9872
		goto L2354
	} else {
		goto L2357
	}
L2357:
	;
	goto L2355
L2358:
	;
	v9915 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+52))
	if v9915 == int32(0) {
		goto L203
	} else {
		goto L2359
	}
L2359:
	;
	v9918 = F_CachedPlanGetTargetList(m, v9761)
	mBase = m.M
	v9919 = m.ExcPending
	if v9919 != 0 {
		goto L1
	} else {
		goto L2360
	}
L2360:
	;
	v9921 = *(*int32)(unsafe.Add(mBase, uint32(v9761)+52))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_222), v9921, v9918, int32(0))
	mBase = m.M
	v9924 = m.ExcPending
	if v9924 != 0 {
		goto L1
	} else {
		goto L2361
	}
L2361:
	;
	goto L198
L2362:
	;
	v9929 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v9929
	v9931 = F_GetPortalByName(m, v9740)
	mBase = m.M
	v9932 = m.ExcPending
	if v9932 != 0 {
		goto L1
	} else {
		goto L2363
	}
L2363:
	;
	if v9931 == int32(0) {
		goto L183
	} else {
		goto L2364
	}
L2364:
	;
	v9936 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v9937 = *(*int32)(unsafe.Add(mBase, uint32(v9936)+24))
	goto L2365
L2365:
	;
	if (v9937-int32(7))&int32(-9) == int32(0) {
		goto L2366
	} else {
		goto L2367
	}
L2366:
	;
	v9944 = *(*int32)(unsafe.Add(mBase, uint32(v9931)+92))
	if v9944 != 0 {
		goto L182
	} else {
		goto L2369
	}
L2367:
	;
	goto L2368
L2368:
	;
	v9946 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9946 != int32(2) {
		goto L198
	} else {
		goto L2370
	}
L2369:
	;
	goto L2368
L2370:
	;
	v9949 = *(*int32)(unsafe.Add(mBase, uint32(v9931)+92))
	if v9949 != 0 {
		goto L2371
	} else {
		goto L2372
	}
L2371:
	;
	v9951 = F_FetchPortalTargetList(m, v9931)
	mBase = m.M
	v9952 = m.ExcPending
	if v9952 != 0 {
		goto L1
	} else {
		goto L2374
	}
L2372:
	;
	goto L2373
L2373:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v9958 = m.ExcPending
	if v9958 != 0 {
		goto L1
	} else {
		goto L2376
	}
L2374:
	;
	v9953 = *(*int32)(unsafe.Add(mBase, uint32(v9931)+96))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_222), v9949, v9951, v9953)
	mBase = m.M
	v9955 = m.ExcPending
	if v9955 != 0 {
		goto L1
	} else {
		goto L2375
	}
L2375:
	;
	goto L198
L2376:
	;
	goto L198
L2377:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9965 = m.ExcPending
	if v9965 != 0 {
		goto L1
	} else {
		goto L2378
	}
L2378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+368)) = v9738
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_224), v40+int32(368))
	mBase = m.M
	v9971 = m.ExcPending
	if v9971 != 0 {
		goto L1
	} else {
		goto L2379
	}
L2379:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_225), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v9976 = m.ExcPending
	if v9976 != 0 {
		goto L1
	} else {
		goto L2380
	}
L2380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2381:
	;
	v9982 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9982 != int32(2) {
		goto L198
	} else {
		goto L2382
	}
L2382:
	;
	v9986 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103]))
	v9987 = *(*int32)(unsafe.Add(mBase, uint32(v9986)+4))
	v9988 = m.T0[v9987].(func(*base.Module) int32)(m)
	mBase = m.M
	v9989 = m.ExcPending
	if v9989 != 0 {
		goto L1
	} else {
		goto L2383
	}
L2383:
	;
	goto L198
L2384:
	;
	v9996 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v9997 = *(*int32)(unsafe.Add(mBase, uint32(v9996)+24))
	if v9997 == int32(4) {
		goto L2386
	} else {
		goto L2387
	}
L2385:
	;
	v10005 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])))
	goto L2389
L2386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9996)+24)) = int32(1)
	goto L2388
L2387:
	;
	goto L2388
L2388:
	;
	goto L2385
L2389:
	;
	if v10005 != 0 {
		goto L2390
	} else {
		goto L2391
	}
L2390:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v10008 = m.ExcPending
	if v10008 != 0 {
		goto L1
	} else {
		goto L2393
	}
L2391:
	;
	goto L2392
L2392:
	;
	v10010 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v10010 != 0 {
		goto L2394
	} else {
		goto L2395
	}
L2393:
	;
	goto L2392
L2394:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v10012 = m.ExcPending
	if v10012 != 0 {
		goto L1
	} else {
		goto L2397
	}
L2395:
	;
	goto L2396
L2396:
	;
	v10017 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v10017)
	goto L198
L2397:
	;
	v10014 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])) = uint8(v10014)
	goto L2396
L2398:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11])) = int32(0)
	goto L2400
L2399:
	;
	goto L2400
L2400:
	;
	v10030 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[135]))
	if v10030 != 0 {
		goto L181
	} else {
		goto L2401
	}
L2401:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v10033 = m.ExcPending
	if v10033 != 0 {
		goto L1
	} else {
		goto L2402
	}
L2402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2403:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10040 = m.ExcPending
	if v10040 != 0 {
		goto L1
	} else {
		goto L2404
	}
L2404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v592
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_11), v40+int32(16))
	mBase = m.M
	v10046 = m.ExcPending
	if v10046 != 0 {
		goto L1
	} else {
		goto L2405
	}
L2405:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_226), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10051 = m.ExcPending
	if v10051 != 0 {
		goto L1
	} else {
		goto L2406
	}
L2406:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2407:
	;
	goto L198
L2408:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v10101 = m.ExcPending
	if v10101 != 0 {
		goto L1
	} else {
		goto L2409
	}
L2409:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_227), int32(0))
	mBase = m.M
	v10105 = m.ExcPending
	if v10105 != 0 {
		goto L1
	} else {
		goto L2410
	}
L2410:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1581), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v10110 = m.ExcPending
	if v10110 != 0 {
		goto L1
	} else {
		goto L2411
	}
L2411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2412:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10117 = m.ExcPending
	if v10117 != 0 {
		goto L1
	} else {
		goto L2413
	}
L2413:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10121 = m.ExcPending
	if v10121 != 0 {
		goto L1
	} else {
		goto L2414
	}
L2414:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10123 = m.ExcPending
	if v10123 != 0 {
		goto L1
	} else {
		goto L2415
	}
L2415:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1603), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v10128 = m.ExcPending
	if v10128 != 0 {
		goto L1
	} else {
		goto L2416
	}
L2416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2417:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v10135 = m.ExcPending
	if v10135 != 0 {
		goto L1
	} else {
		goto L2418
	}
L2418:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_228), int32(0))
	mBase = m.M
	v10139 = m.ExcPending
	if v10139 != 0 {
		goto L1
	} else {
		goto L2419
	}
L2419:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1778), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10144 = m.ExcPending
	if v10144 != 0 {
		goto L1
	} else {
		goto L2420
	}
L2420:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2421:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10151 = m.ExcPending
	if v10151 != 0 {
		goto L1
	} else {
		goto L2422
	}
L2422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+196)) = v7071
	*(*int32)(unsafe.Add(mBase, uint32(v40)+192)) = v6973
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_229), v40+int32(192))
	mBase = m.M
	v10158 = m.ExcPending
	if v10158 != 0 {
		goto L1
	} else {
		goto L2423
	}
L2423:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1831), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10163 = m.ExcPending
	if v10163 != 0 {
		goto L1
	} else {
		goto L2424
	}
L2424:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2425:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10170 = m.ExcPending
	if v10170 != 0 {
		goto L1
	} else {
		goto L2426
	}
L2426:
	;
	v10171 = *(*int32)(unsafe.Add(mBase, uint32(v6812)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v10171
	*(*int32)(unsafe.Add(mBase, uint32(v40)+180)) = v6778
	*(*int32)(unsafe.Add(mBase, uint32(v40)+176)) = v7071
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_230), v40+int32(176))
	mBase = m.M
	v10179 = m.ExcPending
	if v10179 != 0 {
		goto L1
	} else {
		goto L2427
	}
L2427:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1837), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10184 = m.ExcPending
	if v10184 != 0 {
		goto L1
	} else {
		goto L2428
	}
L2428:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2429:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10192 = m.ExcPending
	if v10192 != 0 {
		goto L1
	} else {
		goto L2430
	}
L2430:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10196 = m.ExcPending
	if v10196 != 0 {
		goto L1
	} else {
		goto L2431
	}
L2431:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10198 = m.ExcPending
	if v10198 != 0 {
		goto L1
	} else {
		goto L2432
	}
L2432:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1855), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10203 = m.ExcPending
	if v10203 != 0 {
		goto L1
	} else {
		goto L2433
	}
L2433:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2434:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v10210 = m.ExcPending
	if v10210 != 0 {
		goto L1
	} else {
		goto L2435
	}
L2435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+160)) = v7184 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), v40+int32(160))
	mBase = m.M
	v10218 = m.ExcPending
	if v10218 != 0 {
		goto L1
	} else {
		goto L2436
	}
L2436:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2051), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10223 = m.ExcPending
	if v10223 != 0 {
		goto L1
	} else {
		goto L2437
	}
L2437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2438:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v10230 = m.ExcPending
	if v10230 != 0 {
		goto L1
	} else {
		goto L2439
	}
L2439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = base.I32_extend16_s(v7253)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v40+int32(80))
	mBase = m.M
	v10237 = m.ExcPending
	if v10237 != 0 {
		goto L1
	} else {
		goto L2440
	}
L2440:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2058), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10242 = m.ExcPending
	if v10242 != 0 {
		goto L1
	} else {
		goto L2441
	}
L2441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2442:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v10249 = m.ExcPending
	if v10249 != 0 {
		goto L1
	} else {
		goto L2443
	}
L2443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+224)) = v7837
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_232), v40+int32(224))
	mBase = m.M
	v10255 = m.ExcPending
	if v10255 != 0 {
		goto L1
	} else {
		goto L2444
	}
L2444:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2244), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v10260 = m.ExcPending
	if v10260 != 0 {
		goto L1
	} else {
		goto L2445
	}
L2445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2446:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10268 = m.ExcPending
	if v10268 != 0 {
		goto L1
	} else {
		goto L2447
	}
L2447:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10272 = m.ExcPending
	if v10272 != 0 {
		goto L1
	} else {
		goto L2448
	}
L2448:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10274 = m.ExcPending
	if v10274 != 0 {
		goto L1
	} else {
		goto L2449
	}
L2449:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2360), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v10279 = m.ExcPending
	if v10279 != 0 {
		goto L1
	} else {
		goto L2450
	}
L2450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2451:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10286 = m.ExcPending
	if v10286 != 0 {
		goto L1
	} else {
		goto L2452
	}
L2452:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_233), int32(0))
	mBase = m.M
	v10290 = m.ExcPending
	if v10290 != 0 {
		goto L1
	} else {
		goto L2453
	}
L2453:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_234), int32(_a_F_PostgresMainLoopOnce_235))
	mBase = m.M
	v10295 = m.ExcPending
	if v10295 != 0 {
		goto L1
	} else {
		goto L2454
	}
L2454:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2455:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10302 = m.ExcPending
	if v10302 != 0 {
		goto L1
	} else {
		goto L2456
	}
L2456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+352)) = v9689
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_236), v40+int32(352))
	mBase = m.M
	v10308 = m.ExcPending
	if v10308 != 0 {
		goto L1
	} else {
		goto L2457
	}
L2457:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_237), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10313 = m.ExcPending
	if v10313 != 0 {
		goto L1
	} else {
		goto L2458
	}
L2458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2459:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v10320 = m.ExcPending
	if v10320 != 0 {
		goto L1
	} else {
		goto L2460
	}
L2460:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_228), int32(0))
	mBase = m.M
	v10324 = m.ExcPending
	if v10324 != 0 {
		goto L1
	} else {
		goto L2461
	}
L2461:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2776), int32(_a_F_PostgresMainLoopOnce_238))
	mBase = m.M
	v10329 = m.ExcPending
	if v10329 != 0 {
		goto L1
	} else {
		goto L2462
	}
L2462:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2463:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10336 = m.ExcPending
	if v10336 != 0 {
		goto L1
	} else {
		goto L2464
	}
L2464:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10340 = m.ExcPending
	if v10340 != 0 {
		goto L1
	} else {
		goto L2465
	}
L2465:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10342 = m.ExcPending
	if v10342 != 0 {
		goto L1
	} else {
		goto L2466
	}
L2466:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2797), int32(_a_F_PostgresMainLoopOnce_238))
	mBase = m.M
	v10347 = m.ExcPending
	if v10347 != 0 {
		goto L1
	} else {
		goto L2467
	}
L2467:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2468:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v10354 = m.ExcPending
	if v10354 != 0 {
		goto L1
	} else {
		goto L2469
	}
L2469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+384)) = v9740
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_232), v40+int32(384))
	mBase = m.M
	v10360 = m.ExcPending
	if v10360 != 0 {
		goto L1
	} else {
		goto L2470
	}
L2470:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2858), int32(_a_F_PostgresMainLoopOnce_239))
	mBase = m.M
	v10365 = m.ExcPending
	if v10365 != 0 {
		goto L1
	} else {
		goto L2471
	}
L2471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2472:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10372 = m.ExcPending
	if v10372 != 0 {
		goto L1
	} else {
		goto L2473
	}
L2473:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10376 = m.ExcPending
	if v10376 != 0 {
		goto L1
	} else {
		goto L2474
	}
L2474:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10378 = m.ExcPending
	if v10378 != 0 {
		goto L1
	} else {
		goto L2475
	}
L2475:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2874), int32(_a_F_PostgresMainLoopOnce_239))
	mBase = m.M
	v10383 = m.ExcPending
	if v10383 != 0 {
		goto L1
	} else {
		goto L2476
	}
L2476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2477:
	;
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2478:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10397 = m.ExcPending
	if v10397 != 0 {
		goto L1
	} else {
		goto L2479
	}
L2479:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_240), int32(0))
	mBase = m.M
	v10401 = m.ExcPending
	if v10401 != 0 {
		goto L1
	} else {
		goto L2480
	}
L2480:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_241), int32(_a_F_PostgresMainLoopOnce_235))
	mBase = m.M
	v10406 = m.ExcPending
	if v10406 != 0 {
		goto L1
	} else {
		goto L2481
	}
L2481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
