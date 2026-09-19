package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v712 int64
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v885 int32
	_ = v885
	var v895 int32
	_ = v895
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v951 int32
	_ = v951
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1356 int64
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1602 int64
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1619 int64
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
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
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1902 int32
	_ = v1902
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1963 int32
	_ = v1963
	var v1964 int64
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1992 int64
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int64
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int64
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2004 int64
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int64
	_ = v2010
	var v2013 int64
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2021 int64
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2028 int64
	_ = v2028
	var v2032 int64
	_ = v2032
	var v2035 int64
	_ = v2035
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int64
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
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
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2179 int64
	_ = v2179
	var v2181 int64
	_ = v2181
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2208 int64
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2223 int64
	_ = v2223
	var v2228 int32
	_ = v2228
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int64
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int64
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2323 int32
	_ = v2323
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2346 int32
	_ = v2346
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2473 int32
	_ = v2473
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2543 int32
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2622 int32
	_ = v2622
	var v2627 int32
	_ = v2627
	var v2634 int32
	_ = v2634
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2649 int32
	_ = v2649
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2689 int32
	_ = v2689
	var v2696 int32
	_ = v2696
	var v2701 int32
	_ = v2701
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2710 int32
	_ = v2710
	var v2714 int32
	_ = v2714
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2791 int32
	_ = v2791
	var v2796 int32
	_ = v2796
	var v2804 int32
	_ = v2804
	var v2808 int32
	_ = v2808
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2821 int32
	_ = v2821
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2847 int32
	_ = v2847
	var v2849 int32
	_ = v2849
	var v2850 int64
	_ = v2850
	var v2853 int64
	_ = v2853
	var v2856 int32
	_ = v2856
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2873 int32
	_ = v2873
	var v2878 int32
	_ = v2878
	var v2883 int32
	_ = v2883
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2945 int32
	_ = v2945
	var v2966 int32
	_ = v2966
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2982 int32
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3005 int32
	_ = v3005
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3028 int32
	_ = v3028
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3039 int32
	_ = v3039
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3067 int32
	_ = v3067
	var v3073 int32
	_ = v3073
	var v3095 int32
	_ = v3095
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3111 int32
	_ = v3111
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3117 int32
	_ = v3117
	var v3122 int32
	_ = v3122
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3128 int32
	_ = v3128
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3136 int32
	_ = v3136
	var v3138 int32
	_ = v3138
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3151 int32
	_ = v3151
	var v3154 int32
	_ = v3154
	var v3158 int32
	_ = v3158
	var v3163 int32
	_ = v3163
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3180 int32
	_ = v3180
	var v3183 int32
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3190 int32
	_ = v3190
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3195 int32
	_ = v3195
	var v3201 int32
	_ = v3201
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3231 int32
	_ = v3231
	var v3234 int32
	_ = v3234
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3244 int32
	_ = v3244
	var v3250 int32
	_ = v3250
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3268 int32
	_ = v3268
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3283 int32
	_ = v3283
	var v3288 int32
	_ = v3288
	var v3292 int32
	_ = v3292
	var v3295 int32
	_ = v3295
	var v3301 int32
	_ = v3301
	var v3305 int32
	_ = v3305
	var v3310 int32
	_ = v3310
	var v3314 int32
	_ = v3314
	var v3317 int32
	_ = v3317
	var v3321 int32
	_ = v3321
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3343 int32
	_ = v3343
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3365 int32
	_ = v3365
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3380 int64
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3384 int64
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3389 int64
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int64
	_ = v3392
	var v3394 int32
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3398 int64
	_ = v3398
	var v3401 int64
	_ = v3401
	var v3405 int32
	_ = v3405
	var v3409 int64
	_ = v3409
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3416 int64
	_ = v3416
	var v3420 int64
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3424 int32
	_ = v3424
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3436 int64
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3442 int64
	_ = v3442
	var v3447 int32
	_ = v3447
	var v3448 int64
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3453 int64
	_ = v3453
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3464 int64
	_ = v3464
	var v3470 int32
	_ = v3470
	var v3475 int32
	_ = v3475
	var v3477 int32
	_ = v3477
	var v3483 int32
	_ = v3483
	var v3494 int32
	_ = v3494
	var v3497 int32
	_ = v3497
	var v3501 int32
	_ = v3501
	var v3505 int32
	_ = v3505
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3517 int32
	_ = v3517
	var v3521 int32
	_ = v3521
	var v3526 int32
	_ = v3526
	var v3531 int64
	_ = v3531
	var v3535 int32
	_ = v3535
	var v3541 int32
	_ = v3541
	var v3544 int64
	_ = v3544
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3555 int32
	_ = v3555
	var v3562 int32
	_ = v3562
	var v3565 int32
	_ = v3565
	var v3569 int32
	_ = v3569
	var v3572 int32
	_ = v3572
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3590 int32
	_ = v3590
	var v3596 int32
	_ = v3596
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3602 int64
	_ = v3602
	var v3607 int32
	_ = v3607
	var v3610 int32
	_ = v3610
	var v3612 int32
	_ = v3612
	var v3619 int32
	_ = v3619
	var v3621 int32
	_ = v3621
	var v3623 int64
	_ = v3623
	var v3625 int32
	_ = v3625
	var v3629 int32
	_ = v3629
	var v3635 int32
	_ = v3635
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3648 int32
	_ = v3648
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3665 int32
	_ = v3665
	var v3667 int32
	_ = v3667
	var v3669 int32
	_ = v3669
	var v3673 int64
	_ = v3673
	var v3675 int32
	_ = v3675
	var v3678 int64
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3693 int32
	_ = v3693
	var v3698 int32
	_ = v3698
	var v3703 int32
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3708 int64
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3729 int32
	_ = v3729
	var v3736 int32
	_ = v3736
	var v3737 int64
	_ = v3737
	var v3739 int64
	_ = v3739
	var v3740 int64
	_ = v3740
	var v3744 int64
	_ = v3744
	var v3748 int32
	_ = v3748
	var v3753 int32
	_ = v3753
	var v3756 int32
	_ = v3756
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3763 int32
	_ = v3763
	var v3765 int32
	_ = v3765
	var v3770 int32
	_ = v3770
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3788 int32
	_ = v3788
	var v3793 int32
	_ = v3793
	var v3797 int32
	_ = v3797
	var v3798 int64
	_ = v3798
	var v3812 int32
	_ = v3812
	var v3813 int32
	_ = v3813
	var v3816 int32
	_ = v3816
	var v3819 int32
	_ = v3819
	var v3820 int32
	_ = v3820
	var v3825 int32
	_ = v3825
	var v3832 int32
	_ = v3832
	var v3835 int32
	_ = v3835
	var v3839 int32
	_ = v3839
	var v3842 int32
	_ = v3842
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3847 int32
	_ = v3847
	var v3849 int32
	_ = v3849
	var v3856 int32
	_ = v3856
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3860 int32
	_ = v3860
	var v3866 int32
	_ = v3866
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3873 int32
	_ = v3873
	var v3874 int32
	_ = v3874
	var v3876 int32
	_ = v3876
	var v3877 int64
	_ = v3877
	var v3879 int32
	_ = v3879
	var v3882 int32
	_ = v3882
	var v3883 int64
	_ = v3883
	var v3886 int32
	_ = v3886
	var v3889 int32
	_ = v3889
	var v3891 int32
	_ = v3891
	var v3898 int32
	_ = v3898
	var v3900 int32
	_ = v3900
	var v3902 int32
	_ = v3902
	var v3903 int64
	_ = v3903
	var v3905 int32
	_ = v3905
	var v3912 int32
	_ = v3912
	var v3915 int32
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3919 int32
	_ = v3919
	var v3921 int32
	_ = v3921
	var v3926 int32
	_ = v3926
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3932 int32
	_ = v3932
	var v3939 int32
	_ = v3939
	var v3940 int32
	_ = v3940
	var v3953 int32
	_ = v3953
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3969 int32
	_ = v3969
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3978 int32
	_ = v3978
	var v3983 int32
	_ = v3983
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3988 int32
	_ = v3988
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v3996 int32
	_ = v3996
	var v3998 int32
	_ = v3998
	var v4000 int32
	_ = v4000
	var v4003 int32
	_ = v4003
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4010 int32
	_ = v4010
	var v4015 int32
	_ = v4015
	var v4018 int32
	_ = v4018
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4024 int32
	_ = v4024
	var v4036 int32
	_ = v4036
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4042 int64
	_ = v4042
	var v4044 int64
	_ = v4044
	var v4047 int64
	_ = v4047
	var v4049 int64
	_ = v4049
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4106 int64
	_ = v4106
	var v4111 int32
	_ = v4111
	var v4112 int32
	_ = v4112
	var v4116 int32
	_ = v4116
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4130 int32
	_ = v4130
	var v4132 int64
	_ = v4132
	var v4172 int32
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4177 int32
	_ = v4177
	var v4182 int32
	_ = v4182
	var v4185 int32
	_ = v4185
	var v4188 int32
	_ = v4188
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4200 int32
	_ = v4200
	var v4201 int32
	_ = v4201
	var v4206 int32
	_ = v4206
	var v4208 int32
	_ = v4208
	var v4209 int32
	_ = v4209
	var v4215 int32
	_ = v4215
	var v4216 int32
	_ = v4216
	var v4224 int32
	_ = v4224
	var v4225 int32
	_ = v4225
	var v4238 int32
	_ = v4238
	var v4239 int32
	_ = v4239
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4251 int32
	_ = v4251
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4262 int32
	_ = v4262
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4287 int32
	_ = v4287
	var v4291 int32
	_ = v4291
	var v4293 int32
	_ = v4293
	var v4295 int32
	_ = v4295
	var v4297 int32
	_ = v4297
	var v4298 int32
	_ = v4298
	var v4301 int32
	_ = v4301
	var v4305 int32
	_ = v4305
	var v4306 int32
	_ = v4306
	var v4308 int32
	_ = v4308
	var v4335 int32
	_ = v4335
	var v4336 int32
	_ = v4336
	var v4341 int32
	_ = v4341
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4349 int32
	_ = v4349
	var v4351 int32
	_ = v4351
	var v4357 int32
	_ = v4357
	var v4360 int32
	_ = v4360
	var v4363 int32
	_ = v4363
	var v4364 int32
	_ = v4364
	var v4365 int32
	_ = v4365
	var v4367 int32
	_ = v4367
	var v4374 int32
	_ = v4374
	var v4375 int32
	_ = v4375
	var v4376 int32
	_ = v4376
	var v4378 int32
	_ = v4378
	var v4384 int32
	_ = v4384
	var v4386 int32
	_ = v4386
	var v4387 int32
	_ = v4387
	var v4388 int32
	_ = v4388
	var v4389 int32
	_ = v4389
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4438 int32
	_ = v4438
	var v4451 int32
	_ = v4451
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4458 int32
	_ = v4458
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4467 int32
	_ = v4467
	var v4469 int32
	_ = v4469
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4484 int32
	_ = v4484
	var v4485 int32
	_ = v4485
	var v4486 int32
	_ = v4486
	var v4488 int32
	_ = v4488
	var v4498 int32
	_ = v4498
	var v4504 int32
	_ = v4504
	var v4507 int32
	_ = v4507
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4517 int32
	_ = v4517
	var v4522 int32
	_ = v4522
	var v4523 int32
	_ = v4523
	var v4524 int32
	_ = v4524
	var v4526 int32
	_ = v4526
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4536 int32
	_ = v4536
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4542 int32
	_ = v4542
	var v4544 int32
	_ = v4544
	var v4546 int32
	_ = v4546
	var v4550 int32
	_ = v4550
	var v4552 int32
	_ = v4552
	var v4554 int32
	_ = v4554
	var v4558 int32
	_ = v4558
	var v4562 int32
	_ = v4562
	var v4563 int32
	_ = v4563
	var v4568 int32
	_ = v4568
	var v4575 int32
	_ = v4575
	var v4593 int32
	_ = v4593
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4602 int32
	_ = v4602
	var v4606 int32
	_ = v4606
	var v4611 int32
	_ = v4611
	var v4612 int32
	_ = v4612
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4619 int32
	_ = v4619
	var v4621 int32
	_ = v4621
	var v4625 int32
	_ = v4625
	var v4628 int32
	_ = v4628
	var v4632 int32
	_ = v4632
	var v4637 int32
	_ = v4637
	var v4641 int32
	_ = v4641
	var v4644 int32
	_ = v4644
	var v4650 int32
	_ = v4650
	var v4655 int32
	_ = v4655
	var v4659 int32
	_ = v4659
	var v4662 int32
	_ = v4662
	var v4666 int32
	_ = v4666
	var v4671 int32
	_ = v4671
	var v4675 int32
	_ = v4675
	var v4678 int32
	_ = v4678
	var v4685 int32
	_ = v4685
	var v4690 int32
	_ = v4690
	var v4694 int32
	_ = v4694
	var v4697 int32
	_ = v4697
	var v4701 int32
	_ = v4701
	var v4706 int32
	_ = v4706
	var v4710 int32
	_ = v4710
	var v4713 int32
	_ = v4713
	var v4717 int32
	_ = v4717
	var v4722 int32
	_ = v4722
	var v4726 int32
	_ = v4726
	var v4729 int32
	_ = v4729
	var v4733 int32
	_ = v4733
	var v4738 int32
	_ = v4738
	var v4742 int32
	_ = v4742
	var v4745 int32
	_ = v4745
	var v4749 int32
	_ = v4749
	var v4754 int32
	_ = v4754
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4770 int32
	_ = v4770
	var v4774 int32
	_ = v4774
	var v4781 int32
	_ = v4781
	var v4786 int32
	_ = v4786
	var v4790 int32
	_ = v4790
	var v4797 int32
	_ = v4797
	var v4802 int32
	_ = v4802
	var v4806 int32
	_ = v4806
	var v4813 int32
	_ = v4813
	var v4818 int32
	_ = v4818
	var v4822 int32
	_ = v4822
	var v4829 int32
	_ = v4829
	var v4834 int32
	_ = v4834
	var v4838 int32
	_ = v4838
	var v4845 int32
	_ = v4845
	var v4850 int32
	_ = v4850
	var v4854 int32
	_ = v4854
	var v4857 int32
	_ = v4857
	var v4861 int32
	_ = v4861
	var v4866 int32
	_ = v4866
	var v4870 int32
	_ = v4870
	var v4873 int32
	_ = v4873
	var v4877 int32
	_ = v4877
	var v4882 int32
	_ = v4882
	var v4886 int32
	_ = v4886
	var v4887 int32
	_ = v4887
	var v4893 int32
	_ = v4893
	var v4898 int32
	_ = v4898
	var v4901 int32
	_ = v4901
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4915 int32
	_ = v4915
	var v4920 int32
	_ = v4920
	var v4924 int32
	_ = v4924
	var v4926 int32
	_ = v4926
	var v4934 int32
	_ = v4934
	var v4939 int32
	_ = v4939
	var v4943 int32
	_ = v4943
	var v4945 int32
	_ = v4945
	var v4953 int32
	_ = v4953
	var v4958 int32
	_ = v4958
	var v4962 int32
	_ = v4962
	var v4964 int32
	_ = v4964
	var v4972 int32
	_ = v4972
	var v4977 int32
	_ = v4977
	var v4981 int32
	_ = v4981
	var v4984 int32
	_ = v4984
	var v4995 int32
	_ = v4995
	var v5000 int32
	_ = v5000
	var v5004 int32
	_ = v5004
	var v5006 int32
	_ = v5006
	var v5014 int32
	_ = v5014
	var v5019 int32
	_ = v5019
	var v5023 int32
	_ = v5023
	var v5026 int32
	_ = v5026
	var v5030 int32
	_ = v5030
	var v5035 int32
	_ = v5035
	var v5042 int32
	_ = v5042
	var v5045 int32
	_ = v5045
	var v5049 int32
	_ = v5049
	var v5054 int32
	_ = v5054
	var v5058 int32
	_ = v5058
	var v5061 int32
	_ = v5061
	var v5067 int32
	_ = v5067
	var v5072 int32
	_ = v5072
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5076 int32
	_ = v5076
	var v5079 int32
	_ = v5079
	var v5080 int32
	_ = v5080
	var v5082 int32
	_ = v5082
	var v5084 int32
	_ = v5084
	var v5096 int32
	_ = v5096
	var v5123 int32
	_ = v5123
	var v5127 int32
	_ = v5127
	var v5129 int32
	_ = v5129
	var v5211 int32
	_ = v5211
	var v5213 int32
	_ = v5213
	var v5218 int32
	_ = v5218
	var v5220 int32
	_ = v5220
	var v5225 int32
	_ = v5225
	var v5235 int32
	_ = v5235
	var v5239 int32
	_ = v5239
	var v5241 int32
	_ = v5241
	var v5246 int32
	_ = v5246
	var v5247 int32
	_ = v5247
	var v5248 int32
	_ = v5248
	var v5251 int32
	_ = v5251
	var v5254 int32
	_ = v5254
	var v5257 int32
	_ = v5257
	var v5267 int32
	_ = v5267
	var v5271 int32
	_ = v5271
	var v5272 int32
	_ = v5272
	var v5274 int32
	_ = v5274
	var v5279 int32
	_ = v5279
	var v5281 int32
	_ = v5281
	var v5284 int32
	_ = v5284
	var v5287 int32
	_ = v5287
	var v5324 int32
	_ = v5324
	var v5328 int32
	_ = v5328
	var v5329 int32
	_ = v5329
	var v5330 int32
	_ = v5330
	var v5332 int32
	_ = v5332
	var v5335 int32
	_ = v5335
	var v5336 int32
	_ = v5336
	var v5354 int32
	_ = v5354
	var v5414 int32
	_ = v5414
	var v5417 int32
	_ = v5417
	var v5418 int32
	_ = v5418
	var v5426 int32
	_ = v5426
	var v5428 int32
	_ = v5428
	var v5431 int32
	_ = v5431
	var v5435 int32
	_ = v5435
	var v5450 int32
	_ = v5450
	var v5472 int32
	_ = v5472
	var v5476 int32
	_ = v5476
	var v5477 int32
	_ = v5477
	var v5478 int32
	_ = v5478
	var v5481 int32
	_ = v5481
	var v5483 int32
	_ = v5483
	var v5484 int32
	_ = v5484
	var v5485 int32
	_ = v5485
	var v5487 int32
	_ = v5487
	var v5489 int32
	_ = v5489
	var v5493 int32
	_ = v5493
	var v5494 int32
	_ = v5494
	var v5500 int32
	_ = v5500
	var v5542 int32
	_ = v5542
	var v5572 int32
	_ = v5572
	var v5584 int32
	_ = v5584
	var v5601 int32
	_ = v5601
	var v5614 int32
	_ = v5614
	var v5622 int32
	_ = v5622
	var v5647 int32
	_ = v5647
	var v5662 int32
	_ = v5662
	var v5665 int32
	_ = v5665
	var v5666 int32
	_ = v5666
	var v5671 int32
	_ = v5671
	var v5675 int32
	_ = v5675
	var v5682 int64
	_ = v5682
	var v5686 int32
	_ = v5686
	var v5688 int32
	_ = v5688
	var v5689 int32
	_ = v5689
	var v5692 int32
	_ = v5692
	var v5703 int32
	_ = v5703
	var v5711 int32
	_ = v5711
	var v5715 int32
	_ = v5715
	var v5722 int64
	_ = v5722
	var v5726 int32
	_ = v5726
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5732 int32
	_ = v5732
	var v5743 int32
	_ = v5743
	var v5747 int32
	_ = v5747
	var v5748 int32
	_ = v5748
	var v5749 int32
	_ = v5749
	var v5754 int32
	_ = v5754
	var v5759 int32
	_ = v5759
	var v5760 int32
	_ = v5760
	var v5769 int32
	_ = v5769
	var v5772 int32
	_ = v5772
	var v5775 int32
	_ = v5775
	var v5785 int32
	_ = v5785
	var v5788 int32
	_ = v5788
	var v5792 int32
	_ = v5792
	var v5794 int32
	_ = v5794
	var v5799 int32
	_ = v5799
	var v5802 int32
	_ = v5802
	var v5804 int32
	_ = v5804
	var v5809 int32
	_ = v5809
	var v5810 int32
	_ = v5810
	var v5816 int32
	_ = v5816
	var v5818 int32
	_ = v5818
	var v5821 int32
	_ = v5821
	var v5822 int32
	_ = v5822
	var v5826 int32
	_ = v5826
	var v5827 int32
	_ = v5827
	var v5828 int32
	_ = v5828
	var v5830 int32
	_ = v5830
	var v5833 int32
	_ = v5833
	var v5835 int32
	_ = v5835
	var v5836 int32
	_ = v5836
	var v5837 int32
	_ = v5837
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
	var v5855 int32
	_ = v5855
	var v5858 int32
	_ = v5858
	var v5868 int32
	_ = v5868
	var v5871 int32
	_ = v5871
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5877 int32
	_ = v5877
	var v5882 int32
	_ = v5882
	var v5883 int32
	_ = v5883
	var v5884 int32
	_ = v5884
	var v5887 int32
	_ = v5887
	var v5888 int32
	_ = v5888
	var v5890 int32
	_ = v5890
	var v5892 int32
	_ = v5892
	var v5894 int32
	_ = v5894
	var v5896 int32
	_ = v5896
	var v5898 int32
	_ = v5898
	var v5899 int32
	_ = v5899
	var v5900 int32
	_ = v5900
	var v5914 int32
	_ = v5914
	var v5918 int32
	_ = v5918
	var v5919 int32
	_ = v5919
	var v5921 int32
	_ = v5921
	var v5922 int32
	_ = v5922
	var v5925 int32
	_ = v5925
	var v5926 int32
	_ = v5926
	var v5927 int32
	_ = v5927
	var v5928 int32
	_ = v5928
	var v5931 int32
	_ = v5931
	var v5936 int32
	_ = v5936
	var v5943 int32
	_ = v5943
	var v5944 int32
	_ = v5944
	var v5945 int32
	_ = v5945
	var v5955 int32
	_ = v5955
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5959 int32
	_ = v5959
	var v5962 int32
	_ = v5962
	var v5963 int32
	_ = v5963
	var v5964 int32
	_ = v5964
	var v5973 int32
	_ = v5973
	var v5974 int32
	_ = v5974
	var v5982 int32
	_ = v5982
	var v5985 int32
	_ = v5985
	var v5987 int32
	_ = v5987
	var v5990 int32
	_ = v5990
	var v5991 int32
	_ = v5991
	var v5997 int32
	_ = v5997
	var v6000 int32
	_ = v6000
	var v6002 int32
	_ = v6002
	var v6004 int32
	_ = v6004
	var v6008 int32
	_ = v6008
	var v6013 int32
	_ = v6013
	var v6015 int32
	_ = v6015
	var v6017 int32
	_ = v6017
	var v6022 int32
	_ = v6022
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6027 int32
	_ = v6027
	var v6058 int32
	_ = v6058
	var v6069 int32
	_ = v6069
	var v6072 int32
	_ = v6072
	var v6074 int32
	_ = v6074
	var v6076 int32
	_ = v6076
	var v6078 int32
	_ = v6078
	var v6081 int32
	_ = v6081
	var v6123 int32
	_ = v6123
	var v6126 int32
	_ = v6126
	var v6127 int32
	_ = v6127
	var v6129 int32
	_ = v6129
	var v6133 int32
	_ = v6133
	var v6135 int32
	_ = v6135
	var v6167 int32
	_ = v6167
	var v6175 int32
	_ = v6175
	var v6178 int32
	_ = v6178
	var v6179 int32
	_ = v6179
	var v6184 int32
	_ = v6184
	var v6185 int32
	_ = v6185
	var v6193 int32
	_ = v6193
	var v6195 int32
	_ = v6195
	var v6199 int32
	_ = v6199
	var v6200 int32
	_ = v6200
	var v6211 int32
	_ = v6211
	var v6213 int32
	_ = v6213
	var v6214 int32
	_ = v6214
	var v6215 int32
	_ = v6215
	var v6219 int32
	_ = v6219
	var v6234 int32
	_ = v6234
	var v6256 int32
	_ = v6256
	var v6260 int32
	_ = v6260
	var v6261 int32
	_ = v6261
	var v6262 int32
	_ = v6262
	var v6265 int32
	_ = v6265
	var v6267 int32
	_ = v6267
	var v6268 int32
	_ = v6268
	var v6269 int32
	_ = v6269
	var v6271 int32
	_ = v6271
	var v6273 int32
	_ = v6273
	var v6275 int32
	_ = v6275
	var v6276 int32
	_ = v6276
	var v6282 int32
	_ = v6282
	var v6303 int32
	_ = v6303
	var v6323 int32
	_ = v6323
	var v6363 int32
	_ = v6363
	var v6408 int32
	_ = v6408
	var v6411 int32
	_ = v6411
	var v6415 int32
	_ = v6415
	var v6419 int64
	_ = v6419
	var v6422 int32
	_ = v6422
	var v6423 int32
	_ = v6423
	var v6424 int32
	_ = v6424
	var v6425 int32
	_ = v6425
	var v6426 int32
	_ = v6426
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6445 int32
	_ = v6445
	var v6480 int32
	_ = v6480
	var v6481 int32
	_ = v6481
	var v6484 int32
	_ = v6484
	var v6495 int32
	_ = v6495
	var v6526 int32
	_ = v6526
	var v6532 int32
	_ = v6532
	var v6537 int32
	_ = v6537
	var v6547 int32
	_ = v6547
	var v6552 int32
	_ = v6552
	var v6553 int32
	_ = v6553
	var v6554 int32
	_ = v6554
	var v6557 int32
	_ = v6557
	var v6563 int32
	_ = v6563
	var v6568 int32
	_ = v6568
	var v6571 int32
	_ = v6571
	var v6572 int32
	_ = v6572
	var v6574 int32
	_ = v6574
	var v6577 int32
	_ = v6577
	var v6582 int32
	_ = v6582
	var v6584 int32
	_ = v6584
	var v6589 int32
	_ = v6589
	var v6590 int32
	_ = v6590
	var v6592 int32
	_ = v6592
	var v6593 int32
	_ = v6593
	var v6594 int32
	_ = v6594
	var v6595 int32
	_ = v6595
	var v6599 int32
	_ = v6599
	var v6602 int32
	_ = v6602
	var v6612 int32
	_ = v6612
	var v6616 int32
	_ = v6616
	var v6617 int32
	_ = v6617
	var v6619 int32
	_ = v6619
	var v6624 int32
	_ = v6624
	var v6625 int32
	_ = v6625
	var v6628 int32
	_ = v6628
	var v6629 int32
	_ = v6629
	var v6631 int32
	_ = v6631
	var v6632 int32
	_ = v6632
	var v6639 int32
	_ = v6639
	var v6642 int32
	_ = v6642
	var v6645 int32
	_ = v6645
	var v6650 int32
	_ = v6650
	var v6651 int32
	_ = v6651
	var v6652 int32
	_ = v6652
	var v6653 int32
	_ = v6653
	var v6656 int32
	_ = v6656
	var v6657 int32
	_ = v6657
	var v6661 int32
	_ = v6661
	var v6668 int32
	_ = v6668
	var v6669 int32
	_ = v6669
	var v6670 int32
	_ = v6670
	var v6671 int32
	_ = v6671
	var v6673 int32
	_ = v6673
	var v6678 int32
	_ = v6678
	var v6679 int32
	_ = v6679
	var v6681 int32
	_ = v6681
	var v6682 int32
	_ = v6682
	var v6685 int32
	_ = v6685
	var v6686 int32
	_ = v6686
	var v6688 int32
	_ = v6688
	var v6689 int32
	_ = v6689
	var v6691 int32
	_ = v6691
	var v6693 int32
	_ = v6693
	var v6697 int32
	_ = v6697
	var v6701 int32
	_ = v6701
	var v6702 int32
	_ = v6702
	var v6707 int32
	_ = v6707
	var v6714 int32
	_ = v6714
	var v6726 int32
	_ = v6726
	var v6727 int32
	_ = v6727
	var v6728 int32
	_ = v6728
	var v6733 int32
	_ = v6733
	var v6735 int32
	_ = v6735
	var v6737 int32
	_ = v6737
	var v6740 int32
	_ = v6740
	var v6742 int32
	_ = v6742
	var v6748 int32
	_ = v6748
	var v6750 int32
	_ = v6750
	var v6755 int32
	_ = v6755
	var v6759 int32
	_ = v6759
	var v6760 int32
	_ = v6760
	var v6765 int32
	_ = v6765
	var v6766 int32
	_ = v6766
	var v6776 int32
	_ = v6776
	var v6780 int32
	_ = v6780
	var v6781 int32
	_ = v6781
	var v6784 int32
	_ = v6784
	var v6787 int32
	_ = v6787
	var v6796 int32
	_ = v6796
	var v6799 int32
	_ = v6799
	var v6801 int32
	_ = v6801
	var v6805 int32
	_ = v6805
	var v6809 int32
	_ = v6809
	var v6814 int32
	_ = v6814
	var v6818 int32
	_ = v6818
	var v6822 int64
	_ = v6822
	var v6825 int32
	_ = v6825
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6829 int32
	_ = v6829
	var v6830 int32
	_ = v6830
	var v6831 int32
	_ = v6831
	var v6834 int32
	_ = v6834
	var v6835 int32
	_ = v6835
	var v6836 int32
	_ = v6836
	var v6838 int32
	_ = v6838
	var v6839 int32
	_ = v6839
	var v6842 int32
	_ = v6842
	var v6848 int32
	_ = v6848
	var v6853 int32
	_ = v6853
	var v6855 int32
	_ = v6855
	var v6857 int32
	_ = v6857
	var v6858 int32
	_ = v6858
	var v6859 int32
	_ = v6859
	var v6861 int32
	_ = v6861
	var v6864 int32
	_ = v6864
	var v6866 int32
	_ = v6866
	var v6870 int32
	_ = v6870
	var v6873 int32
	_ = v6873
	var v6876 int32
	_ = v6876
	var v6886 int32
	_ = v6886
	var v6918 int32
	_ = v6918
	var v6919 int64
	_ = v6919
	var v6923 int32
	_ = v6923
	var v6928 int32
	_ = v6928
	var v6932 int32
	_ = v6932
	var v6939 int64
	_ = v6939
	var v6943 int32
	_ = v6943
	var v6945 int32
	_ = v6945
	var v6946 int32
	_ = v6946
	var v6949 int32
	_ = v6949
	var v6960 int32
	_ = v6960
	var v7003 int32
	_ = v7003
	var v7013 int32
	_ = v7013
	var v7017 int32
	_ = v7017
	var v7020 int32
	_ = v7020
	var v7025 int32
	_ = v7025
	var v7026 int32
	_ = v7026
	var v7032 int32
	_ = v7032
	var v7033 int32
	_ = v7033
	var v7042 int32
	_ = v7042
	var v7077 int32
	_ = v7077
	var v7078 int32
	_ = v7078
	var v7081 int32
	_ = v7081
	var v7084 int32
	_ = v7084
	var v7123 int32
	_ = v7123
	var v7124 int32
	_ = v7124
	var v7129 int32
	_ = v7129
	var v7132 int32
	_ = v7132
	var v7133 int32
	_ = v7133
	var v7140 int32
	_ = v7140
	var v7143 int32
	_ = v7143
	var v7146 int32
	_ = v7146
	var v7149 int32
	_ = v7149
	var v7156 int32
	_ = v7156
	var v7159 int32
	_ = v7159
	var v7161 int32
	_ = v7161
	var v7162 int32
	_ = v7162
	var v7163 int32
	_ = v7163
	var v7165 int32
	_ = v7165
	var v7166 int32
	_ = v7166
	var v7167 int32
	_ = v7167
	var v7168 int32
	_ = v7168
	var v7169 int32
	_ = v7169
	var v7171 int32
	_ = v7171
	var v7173 int32
	_ = v7173
	var v7174 int32
	_ = v7174
	var v7175 int32
	_ = v7175
	var v7176 int32
	_ = v7176
	var v7177 int32
	_ = v7177
	var v7178 int32
	_ = v7178
	var v7179 int32
	_ = v7179
	var v7182 int32
	_ = v7182
	var v7183 int32
	_ = v7183
	var v7189 int32
	_ = v7189
	var v7190 int32
	_ = v7190
	var v7194 int32
	_ = v7194
	var v7197 int32
	_ = v7197
	var v7198 int32
	_ = v7198
	var v7200 int32
	_ = v7200
	var v7201 int32
	_ = v7201
	var v7202 int32
	_ = v7202
	var v7204 int32
	_ = v7204
	var v7205 int32
	_ = v7205
	var v7209 int32
	_ = v7209
	var v7210 int32
	_ = v7210
	var v7223 int32
	_ = v7223
	var v7224 int32
	_ = v7224
	var v7227 int32
	_ = v7227
	var v7238 int32
	_ = v7238
	var v7255 int32
	_ = v7255
	var v7268 int32
	_ = v7268
	var v7269 int32
	_ = v7269
	var v7271 int32
	_ = v7271
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7286 int32
	_ = v7286
	var v7287 int32
	_ = v7287
	var v7288 int32
	_ = v7288
	var v7289 int32
	_ = v7289
	var v7290 int32
	_ = v7290
	var v7295 int32
	_ = v7295
	var v7297 int32
	_ = v7297
	var v7304 int32
	_ = v7304
	var v7305 int32
	_ = v7305
	var v7312 int32
	_ = v7312
	var v7314 int32
	_ = v7314
	var v7315 int32
	_ = v7315
	var v7316 int32
	_ = v7316
	var v7317 int32
	_ = v7317
	var v7319 int32
	_ = v7319
	var v7320 int32
	_ = v7320
	var v7322 int32
	_ = v7322
	var v7323 int32
	_ = v7323
	var v7324 int32
	_ = v7324
	var v7329 int32
	_ = v7329
	var v7330 int32
	_ = v7330
	var v7331 int32
	_ = v7331
	var v7334 int32
	_ = v7334
	var v7338 int32
	_ = v7338
	var v7339 int32
	_ = v7339
	var v7341 int32
	_ = v7341
	var v7342 int32
	_ = v7342
	var v7343 int32
	_ = v7343
	var v7346 int32
	_ = v7346
	var v7347 int32
	_ = v7347
	var v7350 int32
	_ = v7350
	var v7351 int32
	_ = v7351
	var v7352 int32
	_ = v7352
	var v7360 int32
	_ = v7360
	var v7361 int32
	_ = v7361
	var v7364 int32
	_ = v7364
	var v7370 int32
	_ = v7370
	var v7371 int32
	_ = v7371
	var v7375 int32
	_ = v7375
	var v7376 int32
	_ = v7376
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7380 int32
	_ = v7380
	var v7381 int32
	_ = v7381
	var v7385 int32
	_ = v7385
	var v7388 int32
	_ = v7388
	var v7391 int32
	_ = v7391
	var v7396 int32
	_ = v7396
	var v7398 int32
	_ = v7398
	var v7403 int32
	_ = v7403
	var v7405 int32
	_ = v7405
	var v7407 int32
	_ = v7407
	var v7408 int32
	_ = v7408
	var v7411 int32
	_ = v7411
	var v7414 int32
	_ = v7414
	var v7415 int32
	_ = v7415
	var v7446 int32
	_ = v7446
	var v7484 int32
	_ = v7484
	var v7492 int32
	_ = v7492
	var v7495 int32
	_ = v7495
	var v7498 int32
	_ = v7498
	var v7499 int32
	_ = v7499
	var v7514 int32
	_ = v7514
	var v7515 int32
	_ = v7515
	var v7521 int32
	_ = v7521
	var v7522 int32
	_ = v7522
	var v7531 int32
	_ = v7531
	var v7566 int32
	_ = v7566
	var v7567 int32
	_ = v7567
	var v7570 int32
	_ = v7570
	var v7581 int32
	_ = v7581
	var v7612 int32
	_ = v7612
	var v7613 int32
	_ = v7613
	var v7615 int32
	_ = v7615
	var v7616 int32
	_ = v7616
	var v7617 int32
	_ = v7617
	var v7618 int32
	_ = v7618
	var v7629 int32
	_ = v7629
	var v7632 int32
	_ = v7632
	var v7635 int32
	_ = v7635
	var v7645 int32
	_ = v7645
	var v7677 int32
	_ = v7677
	var v7678 int64
	_ = v7678
	var v7682 int32
	_ = v7682
	var v7687 int32
	_ = v7687
	var v7691 int32
	_ = v7691
	var v7698 int64
	_ = v7698
	var v7702 int32
	_ = v7702
	var v7704 int32
	_ = v7704
	var v7705 int32
	_ = v7705
	var v7708 int32
	_ = v7708
	var v7719 int32
	_ = v7719
	var v7761 int32
	_ = v7761
	var v7762 int32
	_ = v7762
	var v7765 int32
	_ = v7765
	var v7767 int32
	_ = v7767
	var v7768 int32
	_ = v7768
	var v7770 int32
	_ = v7770
	var v7771 int32
	_ = v7771
	var v7774 int32
	_ = v7774
	var v7779 int32
	_ = v7779
	var v7783 int32
	_ = v7783
	var v7784 int32
	_ = v7784
	var v7789 int32
	_ = v7789
	var v7790 int32
	_ = v7790
	var v7800 int32
	_ = v7800
	var v7802 int32
	_ = v7802
	var v7806 int32
	_ = v7806
	var v7807 int32
	_ = v7807
	var v7810 int32
	_ = v7810
	var v7811 int32
	_ = v7811
	var v7812 int32
	_ = v7812
	var v7815 int32
	_ = v7815
	var v7819 int32
	_ = v7819
	var v7822 int32
	_ = v7822
	var v7831 int32
	_ = v7831
	var v7833 int32
	_ = v7833
	var v7834 int32
	_ = v7834
	var v7837 int32
	_ = v7837
	var v7841 int32
	_ = v7841
	var v7845 int32
	_ = v7845
	var v7846 int32
	_ = v7846
	var v7849 int32
	_ = v7849
	var v7857 int32
	_ = v7857
	var v7860 int32
	_ = v7860
	var v7864 int32
	_ = v7864
	var v7872 int32
	_ = v7872
	var v7877 int32
	_ = v7877
	var v7881 int32
	_ = v7881
	var v7885 int64
	_ = v7885
	var v7888 int32
	_ = v7888
	var v7889 int32
	_ = v7889
	var v7890 int32
	_ = v7890
	var v7892 int32
	_ = v7892
	var v7893 int32
	_ = v7893
	var v7895 int32
	_ = v7895
	var v7897 int32
	_ = v7897
	var v7899 int32
	_ = v7899
	var v7900 int32
	_ = v7900
	var v7901 int32
	_ = v7901
	var v7907 int32
	_ = v7907
	var v7908 int32
	_ = v7908
	var v7912 int32
	_ = v7912
	var v7913 int32
	_ = v7913
	var v7916 int32
	_ = v7916
	var v7919 int32
	_ = v7919
	var v7920 int32
	_ = v7920
	var v7921 int32
	_ = v7921
	var v7925 int32
	_ = v7925
	var v7926 int32
	_ = v7926
	var v7932 int32
	_ = v7932
	var v7933 int32
	_ = v7933
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
	var v7940 int32
	_ = v7940
	var v7941 int32
	_ = v7941
	var v7946 int32
	_ = v7946
	var v7949 int32
	_ = v7949
	var v7952 int32
	_ = v7952
	var v7955 int32
	_ = v7955
	var v7956 int32
	_ = v7956
	var v7966 int32
	_ = v7966
	var v7998 int32
	_ = v7998
	var v7999 int64
	_ = v7999
	var v8003 int32
	_ = v8003
	var v8008 int32
	_ = v8008
	var v8012 int32
	_ = v8012
	var v8019 int64
	_ = v8019
	var v8023 int32
	_ = v8023
	var v8025 int32
	_ = v8025
	var v8026 int32
	_ = v8026
	var v8029 int32
	_ = v8029
	var v8040 int32
	_ = v8040
	var v8044 int32
	_ = v8044
	var v8047 int32
	_ = v8047
	var v8048 int32
	_ = v8048
	var v8051 int32
	_ = v8051
	var v8052 int32
	_ = v8052
	var v8054 int32
	_ = v8054
	var v8055 int32
	_ = v8055
	var v8058 int32
	_ = v8058
	var v8068 int32
	_ = v8068
	var v8100 int32
	_ = v8100
	var v8101 int64
	_ = v8101
	var v8105 int32
	_ = v8105
	var v8110 int32
	_ = v8110
	var v8114 int32
	_ = v8114
	var v8121 int64
	_ = v8121
	var v8125 int32
	_ = v8125
	var v8127 int32
	_ = v8127
	var v8128 int32
	_ = v8128
	var v8131 int32
	_ = v8131
	var v8142 int32
	_ = v8142
	var v8183 int32
	_ = v8183
	var v8188 int32
	_ = v8188
	var v8194 int32
	_ = v8194
	var v8204 int32
	_ = v8204
	var v8208 int32
	_ = v8208
	var v8209 int32
	_ = v8209
	var v8214 int32
	_ = v8214
	var v8215 int32
	_ = v8215
	var v8216 int32
	_ = v8216
	var v8218 int32
	_ = v8218
	var v8219 int32
	_ = v8219
	var v8222 int32
	_ = v8222
	var v8233 int32
	_ = v8233
	var v8262 int32
	_ = v8262
	var v8266 int32
	_ = v8266
	var v8267 int32
	_ = v8267
	var v8268 int32
	_ = v8268
	var v8270 int32
	_ = v8270
	var v8273 int32
	_ = v8273
	var v8274 int32
	_ = v8274
	var v8316 int32
	_ = v8316
	var v8317 int32
	_ = v8317
	var v8321 int32
	_ = v8321
	var v8325 int32
	_ = v8325
	var v8329 int32
	_ = v8329
	var v8331 int32
	_ = v8331
	var v8336 int32
	_ = v8336
	var v8342 int32
	_ = v8342
	var v8344 int32
	_ = v8344
	var v8347 int32
	_ = v8347
	var v8351 int32
	_ = v8351
	var v8355 int32
	_ = v8355
	var v8356 int32
	_ = v8356
	var v8359 int32
	_ = v8359
	var v8367 int32
	_ = v8367
	var v8373 int32
	_ = v8373
	var v8382 int32
	_ = v8382
	var v8412 int32
	_ = v8412
	var v8413 int32
	_ = v8413
	var v8420 int32
	_ = v8420
	var v8423 int32
	_ = v8423
	var v8426 int32
	_ = v8426
	var v8427 int32
	_ = v8427
	var v8428 int32
	_ = v8428
	var v8431 int32
	_ = v8431
	var v8434 int32
	_ = v8434
	var v8437 int32
	_ = v8437
	var v8444 int32
	_ = v8444
	var v8446 int32
	_ = v8446
	var v8447 int32
	_ = v8447
	var v8450 int32
	_ = v8450
	var v8451 int32
	_ = v8451
	var v8465 int32
	_ = v8465
	var v8469 int32
	_ = v8469
	var v8470 int32
	_ = v8470
	var v8471 int32
	_ = v8471
	var v8473 int32
	_ = v8473
	var v8474 int32
	_ = v8474
	var v8476 int32
	_ = v8476
	var v8477 int32
	_ = v8477
	var v8482 int32
	_ = v8482
	var v8490 int32
	_ = v8490
	var v8493 int32
	_ = v8493
	var v8494 int32
	_ = v8494
	var v8496 int32
	_ = v8496
	var v8500 int32
	_ = v8500
	var v8502 int32
	_ = v8502
	var v8505 int32
	_ = v8505
	var v8506 int32
	_ = v8506
	var v8508 int32
	_ = v8508
	var v8515 int32
	_ = v8515
	var v8520 int32
	_ = v8520
	var v8521 int32
	_ = v8521
	var v8525 int32
	_ = v8525
	var v8527 int32
	_ = v8527
	var v8532 int32
	_ = v8532
	var v8533 int32
	_ = v8533
	var v8535 int32
	_ = v8535
	var v8539 int32
	_ = v8539
	var v8542 int32
	_ = v8542
	var v8543 int32
	_ = v8543
	var v8548 int32
	_ = v8548
	var v8549 int32
	_ = v8549
	var v8559 int32
	_ = v8559
	var v8561 int32
	_ = v8561
	var v8565 int32
	_ = v8565
	var v8566 int32
	_ = v8566
	var v8569 int32
	_ = v8569
	var v8572 int32
	_ = v8572
	var v8577 int32
	_ = v8577
	var v8583 int32
	_ = v8583
	var v8592 int32
	_ = v8592
	var v8594 int32
	_ = v8594
	var v8595 int32
	_ = v8595
	var v8598 int32
	_ = v8598
	var v8602 int32
	_ = v8602
	var v8606 int32
	_ = v8606
	var v8607 int32
	_ = v8607
	var v8610 int32
	_ = v8610
	var v8618 int32
	_ = v8618
	var v8620 int32
	_ = v8620
	var v8624 int32
	_ = v8624
	var v8631 int32
	_ = v8631
	var v8636 int32
	_ = v8636
	var v8640 int32
	_ = v8640
	var v8644 int64
	_ = v8644
	var v8650 int32
	_ = v8650
	var v8653 int32
	_ = v8653
	var v8656 int32
	_ = v8656
	var v8657 int32
	_ = v8657
	var v8659 int32
	_ = v8659
	var v8662 int32
	_ = v8662
	var v8663 int32
	_ = v8663
	var v8672 int32
	_ = v8672
	var v8673 int32
	_ = v8673
	var v8675 int32
	_ = v8675
	var v8677 int32
	_ = v8677
	var v8678 int32
	_ = v8678
	var v8685 int32
	_ = v8685
	var v8686 int32
	_ = v8686
	var v8689 int32
	_ = v8689
	var v8690 int32
	_ = v8690
	var v8691 int32
	_ = v8691
	var v8692 int32
	_ = v8692
	var v8695 int32
	_ = v8695
	var v8698 int32
	_ = v8698
	var v8701 int32
	_ = v8701
	var v8703 int32
	_ = v8703
	var v8706 int32
	_ = v8706
	var v8707 int32
	_ = v8707
	var v8709 int32
	_ = v8709
	var v8714 int32
	_ = v8714
	var v8716 int32
	_ = v8716
	var v8718 int32
	_ = v8718
	var v8725 int32
	_ = v8725
	var v8729 int32
	_ = v8729
	var v8741 int32
	_ = v8741
	var v8742 int32
	_ = v8742
	var v8743 int32
	_ = v8743
	var v8745 int32
	_ = v8745
	var v8749 int32
	_ = v8749
	var v8750 int32
	_ = v8750
	var v8752 int32
	_ = v8752
	var v8753 int32
	_ = v8753
	var v8754 int32
	_ = v8754
	var v8756 int32
	_ = v8756
	var v8762 int32
	_ = v8762
	var v8763 int32
	_ = v8763
	var v8764 int32
	_ = v8764
	var v8765 int32
	_ = v8765
	var v8768 int32
	_ = v8768
	var v8775 int32
	_ = v8775
	var v8776 int32
	_ = v8776
	var v8777 int32
	_ = v8777
	var v8780 int32
	_ = v8780
	var v8783 int32
	_ = v8783
	var v8788 int32
	_ = v8788
	var v8789 int32
	_ = v8789
	var v8791 int32
	_ = v8791
	var v8793 int32
	_ = v8793
	var v8797 int32
	_ = v8797
	var v8798 int32
	_ = v8798
	var v8799 int32
	_ = v8799
	var v8804 int32
	_ = v8804
	var v8805 int32
	_ = v8805
	var v8806 int32
	_ = v8806
	var v8809 int32
	_ = v8809
	var v8810 int32
	_ = v8810
	var v8811 int32
	_ = v8811
	var v8813 int32
	_ = v8813
	var v8817 int32
	_ = v8817
	var v8818 int32
	_ = v8818
	var v8820 int32
	_ = v8820
	var v8822 int32
	_ = v8822
	var v8824 int32
	_ = v8824
	var v8825 int32
	_ = v8825
	var v8828 int32
	_ = v8828
	var v8835 int32
	_ = v8835
	var v8839 int32
	_ = v8839
	var v8841 int32
	_ = v8841
	var v8844 int32
	_ = v8844
	var v8849 int32
	_ = v8849
	var v8850 int32
	_ = v8850
	var v8859 int32
	_ = v8859
	var v8864 int32
	_ = v8864
	var v8866 int32
	_ = v8866
	var v8868 int32
	_ = v8868
	var v8870 int32
	_ = v8870
	var v8871 int32
	_ = v8871
	var v8873 int32
	_ = v8873
	var v8874 int32
	_ = v8874
	var v8875 int32
	_ = v8875
	var v8877 int32
	_ = v8877
	var v8879 int32
	_ = v8879
	var v8880 int32
	_ = v8880
	var v8882 int32
	_ = v8882
	var v8883 int32
	_ = v8883
	var v8886 int32
	_ = v8886
	var v8888 int32
	_ = v8888
	var v8889 int32
	_ = v8889
	var v8891 int32
	_ = v8891
	var v8892 int32
	_ = v8892
	var v8894 int32
	_ = v8894
	var v8897 int32
	_ = v8897
	var v8899 int32
	_ = v8899
	var v8900 int64
	_ = v8900
	var v8906 int32
	_ = v8906
	var v8907 int32
	_ = v8907
	var v8912 int32
	_ = v8912
	var v8913 int32
	_ = v8913
	var v8923 int32
	_ = v8923
	var v8955 int32
	_ = v8955
	var v8956 int32
	_ = v8956
	var v8959 int32
	_ = v8959
	var v8961 int32
	_ = v8961
	var v8999 int32
	_ = v8999
	var v9000 int32
	_ = v9000
	var v9003 int32
	_ = v9003
	var v9013 int32
	_ = v9013
	var v9017 int32
	_ = v9017
	var v9030 int32
	_ = v9030
	var v9059 int32
	_ = v9059
	var v9062 int32
	_ = v9062
	var v9064 int32
	_ = v9064
	var v9065 int32
	_ = v9065
	var v9068 int32
	_ = v9068
	var v9070 int32
	_ = v9070
	var v9075 int32
	_ = v9075
	var v9076 int32
	_ = v9076
	var v9077 int32
	_ = v9077
	var v9083 int32
	_ = v9083
	var v9084 int32
	_ = v9084
	var v9086 int32
	_ = v9086
	var v9095 int32
	_ = v9095
	var v9096 int32
	_ = v9096
	var v9101 int32
	_ = v9101
	var v9107 int32
	_ = v9107
	var v9111 int32
	_ = v9111
	var v9112 int32
	_ = v9112
	var v9113 int32
	_ = v9113
	var v9114 int32
	_ = v9114
	var v9116 int32
	_ = v9116
	var v9117 int32
	_ = v9117
	var v9119 int32
	_ = v9119
	var v9120 int32
	_ = v9120
	var v9124 int32
	_ = v9124
	var v9127 int32
	_ = v9127
	var v9131 int32
	_ = v9131
	var v9137 int32
	_ = v9137
	var v9139 int32
	_ = v9139
	var v9144 int32
	_ = v9144
	var v9145 int32
	_ = v9145
	var v9146 int32
	_ = v9146
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
	var v9196 int32
	_ = v9196
	var v9197 int32
	_ = v9197
	var v9199 int32
	_ = v9199
	var v9200 int32
	_ = v9200
	var v9203 int32
	_ = v9203
	var v9216 int32
	_ = v9216
	var v9249 int32
	_ = v9249
	var v9253 int32
	_ = v9253
	var v9255 int32
	_ = v9255
	var v9261 int32
	_ = v9261
	var v9264 int32
	_ = v9264
	var v9268 int32
	_ = v9268
	var v9273 int32
	_ = v9273
	var v9277 int32
	_ = v9277
	var v9280 int32
	_ = v9280
	var v9284 int32
	_ = v9284
	var v9289 int32
	_ = v9289
	var v9293 int32
	_ = v9293
	var v9296 int32
	_ = v9296
	var v9304 int32
	_ = v9304
	var v9309 int32
	_ = v9309
	var v9313 int32
	_ = v9313
	var v9323 int32
	_ = v9323
	var v9328 int32
	_ = v9328
	var v9332 int32
	_ = v9332
	var v9335 int32
	_ = v9335
	var v9337 int32
	_ = v9337
	var v9343 int32
	_ = v9343
	var v9348 int32
	_ = v9348
	var v9352 int32
	_ = v9352
	var v9355 int32
	_ = v9355
	var v9362 int32
	_ = v9362
	var v9367 int32
	_ = v9367
	var v9371 int32
	_ = v9371
	var v9374 int32
	_ = v9374
	var v9380 int32
	_ = v9380
	var v9385 int32
	_ = v9385
	var v9389 int32
	_ = v9389
	var v9392 int32
	_ = v9392
	var v9400 int32
	_ = v9400
	var v9405 int32
	_ = v9405
	var v9409 int32
	_ = v9409
	var v9412 int32
	_ = v9412
	var v9419 int32
	_ = v9419
	var v9424 int32
	_ = v9424
	var v9464 int32
	_ = v9464
	var v9465 int32
	_ = v9465
	var v9466 int32
	_ = v9466
	var v9504 int32
	_ = v9504
	var v9506 int32
	_ = v9506
	var v9508 int32
	_ = v9508
	var v9509 int32
	_ = v9509
	var v9510 int32
	_ = v9510
	var v9512 int32
	_ = v9512
	var v9515 int32
	_ = v9515
	var v9520 int32
	_ = v9520
	var v9521 int32
	_ = v9521
	var v9522 int32
	_ = v9522
	var v9536 int32
	_ = v9536
	var v9537 int32
	_ = v9537
	var v9538 int32
	_ = v9538
	var v9540 int32
	_ = v9540
	var v9543 int32
	_ = v9543
	var v9545 int32
	_ = v9545
	var v9548 int32
	_ = v9548
	var v9549 int64
	_ = v9549
	var v9553 int32
	_ = v9553
	var v9557 int32
	_ = v9557
	var v9561 int32
	_ = v9561
	var v9562 int32
	_ = v9562
	var v9563 int32
	_ = v9563
	var v9564 int32
	_ = v9564
	var v9567 int32
	_ = v9567
	var v9570 int32
	_ = v9570
	var v9571 int32
	_ = v9571
	var v9572 int32
	_ = v9572
	var v9579 int32
	_ = v9579
	var v9580 int32
	_ = v9580
	var v9584 int32
	_ = v9584
	var v9589 int32
	_ = v9589
	var v9590 int32
	_ = v9590
	var v9592 int32
	_ = v9592
	var v9595 int32
	_ = v9595
	var v9596 int32
	_ = v9596
	var v9597 int32
	_ = v9597
	var v9599 int32
	_ = v9599
	var v9601 int32
	_ = v9601
	var v9602 int32
	_ = v9602
	var v9603 int32
	_ = v9603
	var v9618 int32
	_ = v9618
	var v9624 int32
	_ = v9624
	var v9626 int32
	_ = v9626
	var v9630 int32
	_ = v9630
	var v9633 int32
	_ = v9633
	var v9640 int32
	_ = v9640
	var v9645 int32
	_ = v9645
	var v9651 int32
	_ = v9651
	var v9654 int32
	_ = v9654
	var v9655 int32
	_ = v9655
	var v9656 int32
	_ = v9656
	var v9657 int32
	_ = v9657
	var v9659 int32
	_ = v9659
	var v9661 int32
	_ = v9661
	var v9667 int32
	_ = v9667
	var v9669 int32
	_ = v9669
	var v9671 int32
	_ = v9671
	var v9675 int32
	_ = v9675
	var v9676 int32
	_ = v9676
	var v9681 int32
	_ = v9681
	var v9682 int32
	_ = v9682
	var v9692 int32
	_ = v9692
	var v9696 int32
	_ = v9696
	var v9697 int32
	_ = v9697
	var v9709 int32
	_ = v9709
	var v9711 int32
	_ = v9711
	var v9714 int32
	_ = v9714
	var v9721 int32
	_ = v9721
	var v9724 int32
	_ = v9724
	var v9726 int32
	_ = v9726
	var v9728 int32
	_ = v9728
	var v9730 int32
	_ = v9730
	var v9733 int32
	_ = v9733
	var v9736 int32
	_ = v9736
	var v9740 int32
	_ = v9740
	var v9741 int32
	_ = v9741
	var v9742 int32
	_ = v9742
	var v9743 int32
	_ = v9743
	var v9744 int32
	_ = v9744
	var v9746 int32
	_ = v9746
	var v9749 int32
	_ = v9749
	var v9752 int32
	_ = v9752
	var v9754 int32
	_ = v9754
	var v9761 int32
	_ = v9761
	var v9762 int32
	_ = v9762
	var v9763 int32
	_ = v9763
	var v9768 int32
	_ = v9768
	var v9771 int32
	_ = v9771
	var v9776 int32
	_ = v9776
	var v9778 int32
	_ = v9778
	var v9782 int32
	_ = v9782
	var v9786 int64
	_ = v9786
	var v9789 int32
	_ = v9789
	var v9790 int32
	_ = v9790
	var v9791 int32
	_ = v9791
	var v9792 int32
	_ = v9792
	var v9793 int32
	_ = v9793
	var v9795 int32
	_ = v9795
	var v9799 int32
	_ = v9799
	var v9802 int32
	_ = v9802
	var v9804 int32
	_ = v9804
	var v9806 int32
	_ = v9806
	var v9807 int32
	_ = v9807
	var v9808 int32
	_ = v9808
	var v9810 int32
	_ = v9810
	var v9813 int32
	_ = v9813
	var v9815 int32
	_ = v9815
	var v9816 int32
	_ = v9816
	var v9823 int32
	_ = v9823
	var v9825 int32
	_ = v9825
	var v9828 int32
	_ = v9828
	var v9832 int32
	_ = v9832
	var v9836 int32
	_ = v9836
	var v9838 int32
	_ = v9838
	var v9839 int32
	_ = v9839
	var v9840 int32
	_ = v9840
	var v9842 int32
	_ = v9842
	var v9846 int32
	_ = v9846
	var v9850 int32
	_ = v9850
	var v9854 int32
	_ = v9854
	var v9866 int32
	_ = v9866
	var v9895 int32
	_ = v9895
	var v9899 int32
	_ = v9899
	var v9903 int32
	_ = v9903
	var v9904 int32
	_ = v9904
	var v9905 int32
	_ = v9905
	var v9907 int32
	_ = v9907
	var v9911 int32
	_ = v9911
	var v9924 int32
	_ = v9924
	var v9925 int32
	_ = v9925
	var v9966 int32
	_ = v9966
	var v9967 int32
	_ = v9967
	var v9970 int32
	_ = v9970
	var v9971 int32
	_ = v9971
	var v9973 int32
	_ = v9973
	var v9976 int32
	_ = v9976
	var v9978 int32
	_ = v9978
	var v9981 int32
	_ = v9981
	var v9983 int32
	_ = v9983
	var v9984 int32
	_ = v9984
	var v9988 int32
	_ = v9988
	var v9989 int32
	_ = v9989
	var v9996 int32
	_ = v9996
	var v9998 int32
	_ = v9998
	var v10001 int32
	_ = v10001
	var v10003 int32
	_ = v10003
	var v10004 int32
	_ = v10004
	var v10005 int32
	_ = v10005
	var v10007 int32
	_ = v10007
	var v10010 int32
	_ = v10010
	var v10014 int32
	_ = v10014
	var v10017 int32
	_ = v10017
	var v10023 int32
	_ = v10023
	var v10028 int32
	_ = v10028
	var v10032 int32
	_ = v10032
	var v10034 int32
	_ = v10034
	var v10038 int32
	_ = v10038
	var v10039 int32
	_ = v10039
	var v10040 int32
	_ = v10040
	var v10041 int32
	_ = v10041
	var v10045 int32
	_ = v10045
	var v10048 int32
	_ = v10048
	var v10049 int32
	_ = v10049
	var v10057 int32
	_ = v10057
	var v10060 int32
	_ = v10060
	var v10062 int32
	_ = v10062
	var v10064 int32
	_ = v10064
	var v10066 int32
	_ = v10066
	var v10069 int32
	_ = v10069
	var v10075 int32
	_ = v10075
	var v10082 int32
	_ = v10082
	var v10085 int32
	_ = v10085
	var v10089 int32
	_ = v10089
	var v10092 int32
	_ = v10092
	var v10098 int32
	_ = v10098
	var v10103 int32
	_ = v10103
	var v10106 int32
	_ = v10106
	var v10150 int32
	_ = v10150
	var v10153 int32
	_ = v10153
	var v10157 int32
	_ = v10157
	var v10162 int32
	_ = v10162
	var v10166 int32
	_ = v10166
	var v10169 int32
	_ = v10169
	var v10173 int32
	_ = v10173
	var v10175 int32
	_ = v10175
	var v10180 int32
	_ = v10180
	var v10184 int32
	_ = v10184
	var v10187 int32
	_ = v10187
	var v10191 int32
	_ = v10191
	var v10196 int32
	_ = v10196
	var v10200 int32
	_ = v10200
	var v10203 int32
	_ = v10203
	var v10210 int32
	_ = v10210
	var v10215 int32
	_ = v10215
	var v10219 int32
	_ = v10219
	var v10222 int32
	_ = v10222
	var v10223 int32
	_ = v10223
	var v10231 int32
	_ = v10231
	var v10236 int32
	_ = v10236
	var v10241 int32
	_ = v10241
	var v10244 int32
	_ = v10244
	var v10248 int32
	_ = v10248
	var v10250 int32
	_ = v10250
	var v10255 int32
	_ = v10255
	var v10259 int32
	_ = v10259
	var v10262 int32
	_ = v10262
	var v10270 int32
	_ = v10270
	var v10275 int32
	_ = v10275
	var v10279 int32
	_ = v10279
	var v10282 int32
	_ = v10282
	var v10289 int32
	_ = v10289
	var v10294 int32
	_ = v10294
	var v10298 int32
	_ = v10298
	var v10301 int32
	_ = v10301
	var v10307 int32
	_ = v10307
	var v10312 int32
	_ = v10312
	var v10317 int32
	_ = v10317
	var v10320 int32
	_ = v10320
	var v10324 int32
	_ = v10324
	var v10326 int32
	_ = v10326
	var v10331 int32
	_ = v10331
	var v10335 int32
	_ = v10335
	var v10338 int32
	_ = v10338
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
	var v10381 int32
	_ = v10381
	var v10385 int32
	_ = v10385
	var v10388 int32
	_ = v10388
	var v10392 int32
	_ = v10392
	var v10394 int32
	_ = v10394
	var v10399 int32
	_ = v10399
	var v10403 int32
	_ = v10403
	var v10406 int32
	_ = v10406
	var v10412 int32
	_ = v10412
	var v10417 int32
	_ = v10417
	var v10421 int32
	_ = v10421
	var v10424 int32
	_ = v10424
	var v10428 int32
	_ = v10428
	var v10430 int32
	_ = v10430
	var v10435 int32
	_ = v10435
	var v10446 int32
	_ = v10446
	var v10449 int32
	_ = v10449
	var v10453 int32
	_ = v10453
	var v10458 int32
	_ = v10458
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
	v665 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])))
	if v665 == int32(1) {
		goto L177
	} else {
		goto L178
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
		v635 = v334
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
	v635 = v334
	goto L79
L104:
	;
	if v359 == int32(0) {
		v635 = v356
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
	v635 = v356
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
	v635 = int32(-1)
	goto L79
L115:
	;
	goto L116
L116:
	;
	v420 = int32(_a_F_PostgresMainLoopOnce_5)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v422 - int32(1)
	v635 = v325
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
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L171
	}
L121:
	;
	F_appendStringInfoChar(m, v40+int32(440), int32(10))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L170
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
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v40)+444))
	if v600 != 0 {
		goto L120
	} else {
		goto L169
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
		goto L158
	default:
		goto L160
	case 11:
		goto L161
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
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	if v513 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L129
L145:
	;
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = int32(1)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v516 == int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+12))
	if v519 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[36]))
	if v523 == v519 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v525 = m.G0
	v527 = v525 - int32(16)
	m.G0 = v527
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[37]))
	if v530 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	v553 = F_pgmem_kill(m, v519, int32(23))
	mBase = m.M
	goto L145
L152:
	;
	m.G0 = v527 + int32(16)
	goto L144
L153:
	;
	v533 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v527)+15)) = uint8(v533)
	goto L154
L154:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[38]))
	v541 = F_write(m, v537, v527+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v541 {
		goto L152
	} else {
		goto L156
	}
L155:
	;
	goto L152
L156:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32]))
	if v545 == int32(27) {
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	goto L123
L159:
	;
	if v560 <= int32(0) {
		goto L121
	} else {
		goto L167
	}
L160:
	;
	F_appendStringInfoChar(m, v40+int32(440), base.I32_extend8_s(v486))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L166
	}
L161:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v40)+444))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[39])))
	if v562 == int32(0) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	if v560 < int32(2) {
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	v568 = v567 + v560
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568-int32(1)))))
	if v571 != int32(10) {
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568-int32(2)))))
	if v576 == int32(59) {
		goto L120
	} else {
		goto L165
	}
L165:
	;
	goto L160
L166:
	;
	goto L122
L167:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587+v560-int32(1)))))
	if v591 != int32(92) {
		goto L121
	} else {
		goto L168
	}
L168:
	;
	v595 = v560 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+444)) = v595
	v598 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v595+v587))) = uint8(v598)
	goto L122
L169:
	;
	v635 = int32(-1)
	goto L79
L170:
	;
	goto L120
L171:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[40])))
	if v616 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v40)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+400)) = v617
	F_pg_printf(m, int32(_a_F_PostgresMainLoopOnce_12), v40+int32(400))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v624 = F_fflush(m, v431)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v635 = int32(81)
	goto L79
L177:
	;
	F_disable_timeout(m, int32(7))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])))
	if v675 == int32(1) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v672 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v672)
	goto L179
L181:
	;
	F_disable_timeout(m, int32(9))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v685 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v682 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])) = uint8(v682)
	goto L183
L185:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v689 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])) = uint8(v689)
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41]))
	if v692 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L187
L189:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	if v635 != int32(-1) {
		goto L212
	} else {
		goto L213
	}
L192:
	;
	goto L191
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10446 = m.ExcPending
	if v10446 != 0 {
		goto L1
	} else {
		goto L2491
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42])) = int32(99)
	goto L2490
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10421 = m.ExcPending
	if v10421 != 0 {
		goto L1
	} else {
		goto L2485
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10403 = m.ExcPending
	if v10403 != 0 {
		goto L1
	} else {
		goto L2481
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10385 = m.ExcPending
	if v10385 != 0 {
		goto L1
	} else {
		goto L2476
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10369 = m.ExcPending
	if v10369 != 0 {
		goto L1
	} else {
		goto L2472
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10351 = m.ExcPending
	if v10351 != 0 {
		goto L1
	} else {
		goto L2468
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10335 = m.ExcPending
	if v10335 != 0 {
		goto L1
	} else {
		goto L2464
	}
L201:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10317 = m.ExcPending
	if v10317 != 0 {
		goto L1
	} else {
		goto L2459
	}
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10298 = m.ExcPending
	if v10298 != 0 {
		goto L1
	} else {
		goto L2455
	}
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10279 = m.ExcPending
	if v10279 != 0 {
		goto L1
	} else {
		goto L2451
	}
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10259 = m.ExcPending
	if v10259 != 0 {
		goto L1
	} else {
		goto L2447
	}
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10241 = m.ExcPending
	if v10241 != 0 {
		goto L1
	} else {
		goto L2442
	}
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10219 = m.ExcPending
	if v10219 != 0 {
		goto L1
	} else {
		goto L2438
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10200 = m.ExcPending
	if v10200 != 0 {
		goto L1
	} else {
		goto L2434
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10184 = m.ExcPending
	if v10184 != 0 {
		goto L1
	} else {
		goto L2430
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10166 = m.ExcPending
	if v10166 != 0 {
		goto L1
	} else {
		goto L2425
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10150 = m.ExcPending
	if v10150 != 0 {
		goto L1
	} else {
		goto L2421
	}
L211:
	;
	m.G0 = v40 + int32(528)
	return
L212:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[28])))
	if v702&int32(1) != 0 {
		goto L211
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	switch v635 + int32(1) {
	case 0:
		goto L219
	default:
		goto L217
	case 67:
		goto L226
	case 68:
		goto L223
	case 69:
		goto L222
	case 70:
		goto L225
	case 71:
		goto L224
	case 73:
		goto L221
	case 81:
		goto L227
	case 82:
		goto L228
	case 84:
		goto L220
	case 89:
		goto L218
	case 100, 101, 103:
		goto L211
	}
L215:
	;
	goto L214
L216:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v10106 = m.ExcPending
	if v10106 != 0 {
		goto L1
	} else {
		goto L2420
	}
L217:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v10089 = m.ExcPending
	if v10089 != 0 {
		goto L1
	} else {
		goto L2416
	}
L218:
	;
	v10075 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10075 == int32(2) {
		goto L2411
	} else {
		goto L2412
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = int32(2)
	goto L218
L220:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v10045 = m.ExcPending
	if v10045 != 0 {
		goto L1
	} else {
		goto L2397
	}
L221:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v10032 = m.ExcPending
	if v10032 != 0 {
		goto L1
	} else {
		goto L2394
	}
L222:
	;
	v9778 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v9778 == int32(1) {
		goto L193
	} else {
		goto L2338
	}
L223:
	;
	v9736 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v9736 == int32(1) {
		goto L193
	} else {
		goto L2320
	}
L224:
	;
	v8636 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v8636 == int32(1) {
		goto L200
	} else {
		goto L2070
	}
L225:
	;
	v7877 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v7877 == int32(1) {
		goto L193
	} else {
		goto L1884
	}
L226:
	;
	v6814 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v6814 == int32(1) {
		goto L193
	} else {
		goto L1673
	}
L227:
	;
	v6411 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v6411 == int32(1) {
		goto L193
	} else {
		goto L1542
	}
L228:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v708 < int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v715 = v40 + int32(440)
	v716 = F_pq_getmsgstring(m, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L233
	}
L230:
	;
	v712 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v712
	goto L232
L231:
	;
	goto L232
L232:
	;
	goto L229
L233:
	;
	F_pq_getmsgend(m, v715)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44])))
	if v721 == int32(1) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v6408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v6408)
	goto L211
L236:
	;
	v724 = m.G0
	v726 = v724 - int32(_a_F_PostgresMainLoopOnce_13)
	m.G0 = v726
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48]))
	if v733 == int32(0) {
		v756 = v729
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L238
L238:
	;
	v5211 = m.G0
	v5213 = v5211 - int32(112)
	m.G0 = v5213
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v716
	v5218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v5220 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])))
	F_pgstat_report_activity(m, int32(3), v716)
	mBase = m.M
	if v5220 == int32(1) {
		goto L1316
	} else {
		goto L1317
	}
L239:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v757 != int32(4) {
		goto L274
	} else {
		goto L275
	}
L240:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	if v736 == int32(4) {
		v756 = v729
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v741 = base.AtomicRmwXchg32(m, v729, int32(76), int32(1))
	if v741 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	F_s_lock(m, v729+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729)+4)) = int32(4)
	v751 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v729)+76)), uint32(v751))
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v756 = v755
	goto L239
L245:
	;
	goto L244
L246:
	;
	m.G0 = v726 + int32(_a_F_PostgresMainLoopOnce_13)
	if v1183 != 0 {
		goto L235
	} else {
		goto L1315
	}
L247:
	;
	F_EndReplicationCommand(m, v5096)
	mBase = m.M
	v5123 = m.ExcPending
	if v5123 != 0 {
		goto L1
	} else {
		goto L1313
	}
L248:
	;
	v5074 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v5075 = m.ExcPending
	if v5075 != 0 {
		goto L1
	} else {
		goto L1309
	}
L249:
	;
	if v4435 == int32(-1) {
		goto L1298
	} else {
		goto L1299
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5023 = m.ExcPending
	if v5023 != 0 {
		goto L1
	} else {
		goto L1294
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5004 = m.ExcPending
	if v5004 != 0 {
		goto L1
	} else {
		goto L1290
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4981 = m.ExcPending
	if v4981 != 0 {
		goto L1
	} else {
		goto L1286
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4962 = m.ExcPending
	if v4962 != 0 {
		goto L1
	} else {
		goto L1282
	}
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4943 = m.ExcPending
	if v4943 != 0 {
		goto L1
	} else {
		goto L1278
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4924 = m.ExcPending
	if v4924 != 0 {
		goto L1
	} else {
		goto L1274
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		goto L1
	} else {
		goto L1270
	}
L257:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v4901 = m.ExcPending
	if v4901 != 0 {
		goto L1
	} else {
		goto L1269
	}
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4886 = m.ExcPending
	if v4886 != 0 {
		goto L1
	} else {
		goto L1266
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4870 = m.ExcPending
	if v4870 != 0 {
		goto L1
	} else {
		goto L1262
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4854 = m.ExcPending
	if v4854 != 0 {
		goto L1
	} else {
		goto L1258
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4838 = m.ExcPending
	if v4838 != 0 {
		goto L1
	} else {
		goto L1255
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4822 = m.ExcPending
	if v4822 != 0 {
		goto L1
	} else {
		goto L1252
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4806 = m.ExcPending
	if v4806 != 0 {
		goto L1
	} else {
		goto L1249
	}
L264:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4790 = m.ExcPending
	if v4790 != 0 {
		goto L1
	} else {
		goto L1246
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L1
	} else {
		goto L1243
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4758 = m.ExcPending
	if v4758 != 0 {
		goto L1
	} else {
		goto L1240
	}
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L1
	} else {
		goto L1236
	}
L268:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4726 = m.ExcPending
	if v4726 != 0 {
		goto L1
	} else {
		goto L1232
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4710 = m.ExcPending
	if v4710 != 0 {
		goto L1
	} else {
		goto L1228
	}
L270:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4694 = m.ExcPending
	if v4694 != 0 {
		goto L1
	} else {
		goto L1224
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4675 = m.ExcPending
	if v4675 != 0 {
		goto L1
	} else {
		goto L1220
	}
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4659 = m.ExcPending
	if v4659 != 0 {
		goto L1
	} else {
		goto L1216
	}
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4641 = m.ExcPending
	if v4641 != 0 {
		goto L1
	} else {
		goto L1212
	}
L274:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51])))
	if v761 != 0 {
		goto L279
	} else {
		goto L280
	}
L275:
	;
	goto L276
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4625 = m.ExcPending
	if v4625 != 0 {
		goto L1
	} else {
		goto L1208
	}
L277:
	;
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v790 != 0 {
		goto L288
	} else {
		goto L289
	}
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L285
	}
L279:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+20))
	goto L282
L280:
	;
	goto L281
L281:
	;
	goto L277
L282:
	;
	if base.B2i32(v764 == int32(2)) == int32(0) {
		goto L278
	} else {
		goto L283
	}
L283:
	;
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[52]))
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53])) = v770
	goto L281
L285:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_16), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(609), int32(_a_F_PostgresMainLoopOnce_18))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[54]))
	if v794 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	goto L290
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v811
	v815 = v726 + int32(428)
	v817 = F_palloc0(m, int32(20))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L298
	}
L293:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[55]))
	v804 = F_AllocSetContextCreateInternal(m, v799, int32(_a_F_PostgresMainLoopOnce_19), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	F_MemoryContextReset(m, v794)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[54])) = v804
	v811 = v804
	goto L292
L297:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[54]))
	v811 = v810
	goto L292
L298:
	;
	if v815 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v842 = int32(0)
	base.MemoryFill(m, v821, v842, int32(96))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	*(*int32)(unsafe.Add(mBase, uint32(v845)+60)) = v842
	v848 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v845)+52)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v845)+44)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v845)+36)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v845)+4)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v845)+12)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v845)+20)) = v842
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = v817
	v862 = F_strlen(m, v716)
	mBase = m.M
	v864 = v862 + int32(2)
	v865 = F_palloc(m, v864)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L310
	}
L300:
	;
	v821 = F_palloc(m, int32(96))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	v827 = int32(28)
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32])) = v827
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L305
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v815))) = v821
	if v821 != 0 {
		goto L299
	} else {
		goto L304
	}
L304:
	;
	v827 = int32(48)
	goto L302
L305:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_22), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_23), int32(275), int32(_a_F_PostgresMainLoopOnce_24))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v726)+428))
	v1165 = m.G0
	v1167 = v1165 - int32(16)
	m.G0 = v1167
	v1171 = F_replication_yylex(m, v1167+int32(8), v1164)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L344
	}
L309:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_25))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L342
	}
L310:
	;
	if v865 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	if v862 <= int32(0) {
		goto L314
	} else {
		goto L315
	}
L312:
	;
	goto L313
L313:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_26))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L341
	}
L314:
	;
	v1062 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v862+v865))) = uint16(v1062)
	if base.Ui32(v864) < base.Ui32(int32(2)) {
		v1150 = v1062
		goto L328
	} else {
		goto L329
	}
L315:
	;
	v870 = v862 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v862) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v885 = v1
	v895 = v1
	goto L319
L317:
	;
	v951 = v1
	goto L318
L318:
	;
	v988 = v951
	v990 = v1
	goto L323
L319:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885+v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v885+v865))) = uint8(v914)
	v917 = v885 | int32(1)
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917+v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v865+v917))) = uint8(v920)
	v923 = v885 | int32(2)
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923+v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v865+v923))) = uint8(v926)
	v929 = v885 | int32(3)
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929+v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v865+v929))) = uint8(v932)
	v934 = int32(4)
	v935 = v885 + v934
	v937 = v895 + v934
	if v937 != v862&int32(2147483644) {
		v885 = v935
		v895 = v937
		goto L319
	} else {
		goto L321
	}
L320:
	;
	if v870 == int32(0) {
		goto L314
	} else {
		goto L322
	}
L321:
	;
	goto L320
L322:
	;
	v951 = v935
	goto L318
L323:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988+v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v988+v865))) = uint8(v1017)
	v1019 = int32(1)
	v1022 = v990 + v1019
	if v1022 != v870 {
		v988 = v988 + v1019
		v990 = v1022
		goto L323
	} else {
		goto L325
	}
L324:
	;
	goto L314
L325:
	;
	goto L324
L326:
	;
	if v1150 == int32(0) {
		goto L309
	} else {
		goto L340
	}
L327:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_27))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L339
	}
L328:
	;
	goto L326
L329:
	;
	v1068 = v864 - int32(2)
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865+v1068))))
	if v1070 != 0 {
		v1150 = v1062
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+v865-int32(1)))))
	if v1074 != 0 {
		v1150 = v1062
		goto L328
	} else {
		goto L331
	}
L331:
	;
	v1076 = F_palloc(m, int32(48))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	if v1076 == int32(0) {
		goto L327
	} else {
		goto L333
	}
L333:
	;
	v1080 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+20)) = v1080
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+8)) = v865
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+4)) = v865
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+12)) = v1068
	*(*int64)(unsafe.Add(mBase, uint32(v1076)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1076)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+16)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1080
	F_replication_yyensure_buffer_stack(m, v860)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1095<<(uint(int32(2))%32))))
	if v1099 == v1076 {
		v1150 = v1076
		goto L328
	} else {
		goto L335
	}
L335:
	;
	if v1099 != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v860)+36))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1101))) = uint8(v1102)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1106 = int32(2)
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1104+v1105<<(uint(v1106)%32))))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v860)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+8)) = v1110
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1112+v1113<<(uint(v1106)%32))))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v860)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+16)) = v1118
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1122 = v1120
	v1123 = v1121
	goto L338
L337:
	;
	v1122 = v1094
	v1123 = v1095
	goto L338
L338:
	;
	v1124 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1123<<(uint(v1124)%32)+v1122))) = v1076
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v860)+12))
	v1132 = v1128 + v1129<<(uint(v1124)%32)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v860)+28)) = v1134
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v860)+36)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v860)+80)) = v1137
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)))
	*(*int32)(unsafe.Add(mBase, uint32(v860)+4)) = v1141
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137))))
	*(*uint8)(unsafe.Add(mBase, uint32(v860)+24)) = uint8(v1143)
	*(*int32)(unsafe.Add(mBase, uint32(v860)+48)) = int32(1)
	v1150 = v1076
	goto L328
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+20)) = int32(1)
	goto L308
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	m.G0 = v1167 + int32(16)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v726)+428))
	if v1183 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L344:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v1171-int32(262)) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if v1171 != int32(282) {
		v1183 = int32(0)
		goto L343
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	*(*int32)(unsafe.Add(mBase, uint32(v1180))) = v1171
	v1183 = int32(1)
	goto L343
L348:
	;
	goto L347
L349:
	;
	F_replication_scanner_finish(m, v1187)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1218 = m.G0
	v1220 = v1218 - int32(1872)
	m.G0 = v1220
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+1864)) = int64(0)
	v1226 = v1220 - int32(-64)
	v1228 = v1220 + int32(1664)
	v1231 = v1228
	v1236 = v1
	v1240 = v1226
	v1243 = v1226
	v1250 = v1228
	v1253 = int32(200)
	v1254 = int32(-2)
	goto L365
L352:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v731
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[54]))
	F_MemoryContextReset(m, v1195)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[56]))
	if v1199 != 0 {
		goto L246
	} else {
		goto L354
	}
L354:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_28), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2064), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	if v1861 != 0 {
		goto L273
	} else {
		goto L532
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L528
	}
L361:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L524
	}
L362:
	;
	if v1220+int32(1664) != v1851 {
		goto L520
	} else {
		goto L521
	}
L363:
	;
	v1851 = v1284
	v1861 = int32(1)
	goto L362
L364:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L1
	} else {
		goto L519
	}
L365:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1250))) = uint8(v1236)
	if base.Ui32(v1231+v1253-int32(1)) <= base.Ui32(v1250) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_31))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L1
	} else {
		goto L518
	}
L367:
	;
	if int32(_a_F_PostgresMainLoopOnce_32) < v1253 {
		goto L364
	} else {
		goto L370
	}
L368:
	;
	v1315 = v1231
	v1318 = v1240
	v1320 = v1243
	v1321 = v1250
	v1322 = v1253
	goto L369
L369:
	;
	if v1236 == int32(34) {
		goto L387
	} else {
		goto L388
	}
L370:
	;
	v1274 = int32(_a_F_PostgresMainLoopOnce_7)
	v1276 = v1253 << (uint(int32(1)) % 32)
	if v1274 <= v1276 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1279 = v1274
	goto L373
L372:
	;
	v1279 = v1276
	goto L373
L373:
	;
	v1284 = F_palloc(m, v1279*int32(9)+int32(7))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	if v1284 == int32(0) {
		goto L364
	} else {
		goto L375
	}
L375:
	;
	v1288 = v1250 - v1231
	v1290 = v1288 + int32(1)
	if v1290 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	base.MemoryCopy(m, v1284, v1231, v1290)
	goto L378
L377:
	;
	goto L378
L378:
	;
	v1295 = base.I32_div_s(v1279+int32(7), int32(8))
	v1296 = int32(3)
	v1298 = v1284 + v1295<<(uint(v1296)%32)
	v1300 = v1290 << (uint(v1296) % 32)
	if v1300 != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	base.MemoryCopy(m, v1298, v1240, v1300)
	goto L381
L380:
	;
	goto L381
L381:
	;
	if v1220+int32(1664) != v1231 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	F_pfree(m, v1231)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	if v1279-int32(1) <= v1288 {
		goto L363
	} else {
		goto L386
	}
L385:
	;
	goto L384
L386:
	;
	v1315 = v1284
	v1318 = v1298
	v1320 = v1298 + v1300 - int32(8)
	v1321 = v1284 + v1288
	v1322 = v1279
	goto L369
L387:
	;
	v1851 = v1315
	v1861 = int32(0)
	goto L362
L388:
	;
	goto L389
L389:
	;
	v1328 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1236)+uint32(_c_F_PostgresMainLoopOnce[57]))))
	if v1328 == int32(-36) {
		v1364 = v1254
		goto L392
	} else {
		goto L393
	}
L390:
	;
	goto L366
L391:
	;
	v1231 = v1315
	v1236 = base.I32_extend8_s(v1838)
	v1240 = v1318
	v1243 = v1835
	v1250 = v1836 + int32(1)
	v1253 = v1322
	v1254 = v1837
	goto L365
L392:
	;
	v1367 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1236)+uint32(_c_F_PostgresMainLoopOnce[58]))))
	v1369 = v1367 & int32(255)
	if v1369 == int32(0) {
		goto L390
	} else {
		goto L408
	}
L393:
	;
	if v1254 == int32(-2) {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	v1351 = v1328 + v1350
	if base.Ui32(int32(80)) < base.Ui32(v1351) {
		v1364 = v1349
		goto L392
	} else {
		goto L406
	}
L395:
	;
	v1335 = F_replication_yylex(m, v1220+int32(1864), v1187)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L398
	}
L396:
	;
	v1337 = v1254
	goto L397
L397:
	;
	if v1337 <= int32(0) {
		goto L399
	} else {
		goto L400
	}
L398:
	;
	v1337 = v1335
	goto L397
L399:
	;
	v1340 = int32(0)
	v1349 = v1340
	v1350 = v1340
	goto L394
L400:
	;
	goto L401
L401:
	;
	if v1337 == int32(256) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1851 = v1315
	v1861 = int32(1)
	goto L362
L403:
	;
	goto L404
L404:
	;
	if base.Ui32(int32(282)) < base.Ui32(v1337) {
		v1349 = v1337
		v1350 = int32(2)
		goto L394
	} else {
		goto L405
	}
L405:
	;
	v1348 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1337)+uint32(_c_F_PostgresMainLoopOnce[59]))))
	v1349 = v1337
	v1350 = v1348
	goto L394
L406:
	;
	v1354 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1351)+uint32(_c_F_PostgresMainLoopOnce[60]))))
	if v1350 != v1354 {
		v1364 = v1349
		goto L392
	} else {
		goto L407
	}
L407:
	;
	v1356 = *(*int64)(unsafe.Add(mBase, uint32(v1220)+1864))
	*(*int64)(unsafe.Add(mBase, uint32(v1320)+8)) = v1356
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351)+uint32(_c_F_PostgresMainLoopOnce[61]))))
	v1835 = v1320 + int32(8)
	v1836 = v1321
	v1837 = int32(-2)
	v1838 = v1361
	goto L391
L408:
	;
	v1375 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_PostgresMainLoopOnce[62]))))
	v1379 = v1320 + (int32(1)-v1375)<<(uint(int32(3))%32)
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1379)))
	v1382 = int32(base.Ui32(v1380) >> (uint(int32(8)) % 32))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1379)+4))
	switch v1369 - int32(2) {
	case 0:
		goto L471
	default:
		v1793 = v1380
		v1794 = v1382
		goto L409
	case 14:
		goto L470
	case 15:
		goto L469
	case 16:
		goto L468
	case 17:
		goto L467
	case 18:
		goto L466
	case 19:
		goto L465
	case 20:
		goto L464
	case 21:
		goto L463
	case 22:
		goto L462
	case 23:
		goto L461
	case 24:
		goto L460
	case 25:
		goto L459
	case 26, 44, 46, 48, 53:
		goto L458
	case 27:
		goto L457
	case 28:
		goto L456
	case 29:
		goto L455
	case 30:
		goto L454
	case 31:
		goto L453
	case 32:
		goto L452
	case 33:
		goto L451
	case 34:
		goto L450
	case 35:
		goto L449
	case 36:
		goto L448
	case 37:
		goto L447
	case 38:
		goto L446
	case 41:
		goto L445
	case 42:
		goto L444
	case 43:
		goto L443
	case 45:
		goto L442
	case 47:
		goto L441
	case 49:
		goto L440
	case 50:
		goto L439
	case 51:
		goto L438
	case 52:
		goto L437
	case 54:
		goto L436
	case 55:
		goto L435
	case 56:
		goto L434
	case 57:
		goto L433
	case 58:
		goto L432
	case 59:
		goto L431
	case 60:
		goto L430
	case 61:
		goto L429
	case 62:
		goto L428
	case 63:
		goto L427
	case 64:
		goto L426
	case 65:
		goto L425
	case 66:
		goto L424
	case 67:
		goto L423
	case 68:
		goto L422
	case 69:
		goto L421
	case 70:
		goto L420
	case 71:
		goto L419
	case 72:
		goto L418
	case 73:
		goto L417
	case 74:
		goto L416
	case 75:
		goto L415
	case 76:
		goto L414
	case 77:
		goto L413
	case 78:
		goto L412
	case 79:
		goto L411
	case 80:
		goto L410
	}
L409:
	;
	v1797 = v1320 - v1375<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1797)+12)) = v1383
	v1801 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1797)+8)) = v1793&int32(255) | v1794<<(uint(v1801)%32)
	v1806 = v1797 + v1801
	v1807 = v1321 - v1375
	v1808 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1807))))
	v1811 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_PostgresMainLoopOnce[63]))))
	v1814 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1811)+uint32(_c_F_PostgresMainLoopOnce[64]))))
	v1815 = v1808 + v1814
	if base.Ui32(v1815) <= base.Ui32(int32(80)) {
		goto L514
	} else {
		goto L515
	}
L410:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_33)
	v1794 = int32(327)
	goto L409
L411:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_34)
	v1794 = int32(363)
	goto L409
L412:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_35)
	v1794 = int32(362)
	goto L409
L413:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_36)
	v1794 = int32(362)
	goto L409
L414:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_37)
	v1794 = int32(1485)
	goto L409
L415:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_38)
	v1794 = int32(69)
	goto L409
L416:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_39)
	v1794 = int32(1266)
	goto L409
L417:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_40)
	v1794 = int32(361)
	goto L409
L418:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_41)
	v1794 = int32(1290)
	goto L409
L419:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_42)
	v1794 = int32(1290)
	goto L409
L420:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_43)
	v1794 = int32(1533)
	goto L409
L421:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_44)
	v1794 = int32(432)
	goto L409
L422:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_45)
	v1794 = int32(50)
	goto L409
L423:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_46)
	v1794 = int32(356)
	goto L409
L424:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_47)
	v1794 = int32(357)
	goto L409
L425:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_48)
	v1794 = int32(357)
	goto L409
L426:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_49)
	v1794 = int32(1094)
	goto L409
L427:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_50)
	v1794 = int32(130)
	goto L409
L428:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_51)
	v1794 = int32(1188)
	goto L409
L429:
	;
	v1793 = int32(_a_F_PostgresMainLoopOnce_52)
	v1794 = int32(961)
	goto L409
L430:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1793 = v1749
	v1794 = int32(base.Ui32(v1749) >> (uint(int32(8)) % 32))
	goto L409
L431:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1742 = F_makeInteger(m, v1741)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L1
	} else {
		goto L511
	}
L432:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1731 = F_makeString(m, v1730)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L1
	} else {
		goto L509
	}
L433:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1720 = F_makeString(m, v1719)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L1
	} else {
		goto L507
	}
L434:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1712 = F_makeDefElem(m, v1709, int32(0), int32(-1))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L506
	}
L435:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+52)) = v1699
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+56)) = v1699
	v1705 = F_list_make1_impl(m, int32(1), v1220+int32(52))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L505
	}
L436:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(16))))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1695 = F_lappend(m, v1693, v1694)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L504
	}
L437:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1687 = F_makeString(m, v1686)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L503
	}
L438:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1682 = F_makeDefElem(m, v1679, v1680, int32(-1))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L502
	}
L439:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(16))))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1673 = F_lappend(m, v1671, v1672)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L1
	} else {
		goto L501
	}
L440:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+48)) = v1659
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+60)) = v1659
	v1665 = F_list_make1_impl(m, int32(1), v1220+int32(48))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L1
	} else {
		goto L500
	}
L441:
	;
	v1654 = int32(8)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1654)))
	v1793 = v1656
	v1794 = int32(base.Ui32(v1656) >> (uint(v1654) % 32))
	goto L409
L442:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	if v1649 == int32(0) {
		goto L360
	} else {
		goto L499
	}
L443:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1793 = v1646
	v1794 = int32(base.Ui32(v1646) >> (uint(int32(8)) % 32))
	goto L409
L444:
	;
	v1793 = int32(0)
	v1794 = v1382
	goto L409
L445:
	;
	v1793 = int32(1)
	v1794 = v1382
	goto L409
L446:
	;
	v1638 = F_palloc0(m, int32(4))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L1
	} else {
		goto L498
	}
L447:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	if v1625 == int32(0) {
		goto L361
	} else {
		goto L496
	}
L448:
	;
	v1609 = F_palloc0(m, int32(32))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L495
	}
L449:
	;
	v1592 = F_palloc0(m, int32(32))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L494
	}
L450:
	;
	v1577 = F_palloc0(m, int32(12))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L493
	}
L451:
	;
	v1564 = F_palloc0(m, int32(12))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L492
	}
L452:
	;
	v1553 = F_palloc0(m, int32(12))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L491
	}
L453:
	;
	v1545 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L489
	}
L454:
	;
	v1536 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L1
	} else {
		goto L487
	}
L455:
	;
	v1527 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_53))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L485
	}
L456:
	;
	v1518 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_54))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L1
	} else {
		goto L483
	}
L457:
	;
	v1509 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_55))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L481
	}
L458:
	;
	v1505 = int32(0)
	v1793 = v1505
	v1794 = v1505
	goto L409
L459:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1501 = F_lappend(m, v1499, v1500)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L480
	}
L460:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1793 = v1494
	v1794 = int32(base.Ui32(v1494) >> (uint(int32(8)) % 32))
	goto L409
L461:
	;
	v1489 = int32(8)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1489)))
	v1793 = v1491
	v1794 = int32(base.Ui32(v1491) >> (uint(v1489) % 32))
	goto L409
L462:
	;
	v1467 = F_palloc0(m, int32(24))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L479
	}
L463:
	;
	v1448 = F_palloc0(m, int32(24))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L478
	}
L464:
	;
	v1441 = F_palloc0(m, int32(8))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L477
	}
L465:
	;
	v1430 = F_palloc0(m, int32(8))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L476
	}
L466:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(16))))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+4)) = v1421
	*(*int32)(unsafe.Add(mBase, uint32(v1220))) = v1420
	v1425 = F_psprintf(m, int32(_a_F_PostgresMainLoopOnce_56), v1220)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L1
	} else {
		goto L475
	}
L467:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1793 = v1415
	v1794 = int32(base.Ui32(v1415) >> (uint(int32(8)) % 32))
	goto L409
L468:
	;
	v1407 = F_palloc0(m, int32(8))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L1
	} else {
		goto L474
	}
L469:
	;
	v1398 = F_palloc0(m, int32(8))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L473
	}
L470:
	;
	v1391 = F_palloc0(m, int32(4))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L472
	}
L471:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v726+int32(424)))) = v1388
	v1793 = v1380
	v1794 = v1382
	goto L409
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1391))) = int32(448)
	v1793 = v1391
	v1794 = int32(base.Ui32(v1391) >> (uint(int32(8)) % 32))
	goto L409
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1398))) = int32(454)
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+4)) = v1402
	v1793 = v1398
	v1794 = int32(base.Ui32(v1398) >> (uint(int32(8)) % 32))
	goto L409
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1407))) = int32(159)
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1407)+4)) = v1411
	v1793 = v1407
	v1794 = int32(base.Ui32(v1407) >> (uint(int32(8)) % 32))
	goto L409
L475:
	;
	v1793 = v1425
	v1794 = int32(base.Ui32(v1425) >> (uint(int32(8)) % 32))
	goto L409
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1430))) = int32(449)
	v1434 = int32(8)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1434)))
	*(*int32)(unsafe.Add(mBase, uint32(v1430)+4)) = v1436
	v1793 = v1430
	v1794 = int32(base.Ui32(v1430) >> (uint(v1434) % 32))
	goto L409
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1441))) = int32(449)
	v1793 = v1441
	v1794 = int32(base.Ui32(v1441) >> (uint(int32(8)) % 32))
	goto L409
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1448))) = int32(450)
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+4)) = v1456
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1320-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1448)+16)) = uint8(v1460)
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1448)+20)) = v1462
	v1793 = v1448
	v1794 = int32(base.Ui32(v1448) >> (uint(int32(8)) % 32))
	goto L409
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1467))) = int32(450)
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+4)) = v1475
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1320-int32(24)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1467)+16)) = uint8(v1479)
	v1481 = int32(8)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1481)))
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+12)) = v1483
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+20)) = v1485
	v1793 = v1467
	v1794 = int32(base.Ui32(v1467) >> (uint(v1481) % 32))
	goto L409
L480:
	;
	v1793 = v1501
	v1794 = int32(base.Ui32(v1501) >> (uint(int32(8)) % 32))
	goto L409
L481:
	;
	v1512 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1509, int32(-1))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v1793 = v1512
	v1794 = int32(base.Ui32(v1512) >> (uint(int32(8)) % 32))
	goto L409
L483:
	;
	v1521 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1518, int32(-1))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	v1793 = v1521
	v1794 = int32(base.Ui32(v1521) >> (uint(int32(8)) % 32))
	goto L409
L485:
	;
	v1530 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_57), v1527, int32(-1))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	v1793 = v1530
	v1794 = int32(base.Ui32(v1530) >> (uint(int32(8)) % 32))
	goto L409
L487:
	;
	v1539 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_39), v1536, int32(-1))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	v1793 = v1539
	v1794 = int32(base.Ui32(v1539) >> (uint(int32(8)) % 32))
	goto L409
L489:
	;
	v1548 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_37), v1545, int32(-1))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v1793 = v1548
	v1794 = int32(base.Ui32(v1548) >> (uint(int32(8)) % 32))
	goto L409
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1553))) = int32(451)
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1558 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1553)+8)) = uint8(v1558)
	*(*int32)(unsafe.Add(mBase, uint32(v1553)+4)) = v1557
	v1793 = v1553
	v1794 = int32(base.Ui32(v1553) >> (uint(int32(8)) % 32))
	goto L409
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1564))) = int32(451)
	v1568 = int32(8)
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1568)))
	v1571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1564)+8)) = uint8(v1571)
	*(*int32)(unsafe.Add(mBase, uint32(v1564)+4)) = v1570
	v1793 = v1564
	v1794 = int32(base.Ui32(v1564) >> (uint(v1568) % 32))
	goto L409
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1577))) = int32(452)
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1577)+4)) = v1583
	v1585 = int32(8)
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1320-v1585)))
	*(*int32)(unsafe.Add(mBase, uint32(v1577)+8)) = v1587
	v1793 = v1577
	v1794 = int32(base.Ui32(v1577) >> (uint(v1585) % 32))
	goto L409
L494:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1592))) = int64(453)
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+8)) = v1598
	v1600 = int32(8)
	v1602 = *(*int64)(unsafe.Add(mBase, uint32(v1320-v1600)))
	*(*int64)(unsafe.Add(mBase, uint32(v1592)+16)) = v1602
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+12)) = v1604
	v1793 = v1592
	v1794 = int32(base.Ui32(v1592) >> (uint(v1600) % 32))
	goto L409
L495:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1609))) = int64(4294967749)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1320-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1609)+8)) = v1615
	v1617 = int32(8)
	v1619 = *(*int64)(unsafe.Add(mBase, uint32(v1320-v1617)))
	*(*int64)(unsafe.Add(mBase, uint32(v1609)+16)) = v1619
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609)+24)) = v1621
	v1793 = v1609
	v1794 = int32(base.Ui32(v1609) >> (uint(v1617) % 32))
	goto L409
L496:
	;
	v1629 = F_palloc0(m, int32(8))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1629))) = int32(455)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+4)) = v1633
	v1793 = v1629
	v1794 = int32(base.Ui32(v1629) >> (uint(int32(8)) % 32))
	goto L409
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1638))) = int32(456)
	v1793 = v1638
	v1794 = int32(base.Ui32(v1638) >> (uint(int32(8)) % 32))
	goto L409
L499:
	;
	v1793 = v1649
	v1794 = int32(base.Ui32(v1649) >> (uint(int32(8)) % 32))
	goto L409
L500:
	;
	v1793 = v1665
	v1794 = int32(base.Ui32(v1665) >> (uint(int32(8)) % 32))
	goto L409
L501:
	;
	v1793 = v1673
	v1794 = int32(base.Ui32(v1673) >> (uint(int32(8)) % 32))
	goto L409
L502:
	;
	v1793 = v1682
	v1794 = int32(base.Ui32(v1682) >> (uint(int32(8)) % 32))
	goto L409
L503:
	;
	v1793 = v1687
	v1794 = int32(base.Ui32(v1687) >> (uint(int32(8)) % 32))
	goto L409
L504:
	;
	v1793 = v1695
	v1794 = int32(base.Ui32(v1695) >> (uint(int32(8)) % 32))
	goto L409
L505:
	;
	v1793 = v1705
	v1794 = int32(base.Ui32(v1705) >> (uint(int32(8)) % 32))
	goto L409
L506:
	;
	v1793 = v1712
	v1794 = int32(base.Ui32(v1712) >> (uint(int32(8)) % 32))
	goto L409
L507:
	;
	v1723 = F_makeDefElem(m, v1718, v1720, int32(-1))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	v1793 = v1723
	v1794 = int32(base.Ui32(v1723) >> (uint(int32(8)) % 32))
	goto L409
L509:
	;
	v1734 = F_makeDefElem(m, v1729, v1731, int32(-1))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	v1793 = v1734
	v1794 = int32(base.Ui32(v1734) >> (uint(int32(8)) % 32))
	goto L409
L511:
	;
	v1745 = F_makeDefElem(m, v1740, v1742, int32(-1))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1793 = v1745
	v1794 = int32(base.Ui32(v1745) >> (uint(int32(8)) % 32))
	goto L409
L513:
	;
	v1827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_PostgresMainLoopOnce[61]))))
	v1835 = v1806
	v1836 = v1807
	v1837 = v1364
	v1838 = v1827
	goto L391
L514:
	;
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+uint32(_c_F_PostgresMainLoopOnce[60]))))
	if v1818 == v1808&int32(255) {
		goto L513
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v1824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1811)+uint32(_c_F_PostgresMainLoopOnce[65]))))
	v1835 = v1806
	v1836 = v1807
	v1837 = v1364
	v1838 = v1824
	goto L391
L517:
	;
	goto L516
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	F_pfree(m, v1851)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L1
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	m.G0 = v1220 + int32(1872)
	goto L359
L523:
	;
	goto L522
L524:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+16)) = v1877
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_58), v1220+int32(16))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_59), int32(322), int32(_a_F_PostgresMainLoopOnce_60))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	*(*int32)(unsafe.Add(mBase, uint32(v1220)+32)) = v1896
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_58), v1220+int32(32))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_59), int32(363), int32(_a_F_PostgresMainLoopOnce_60))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L532:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v726)+428))
	F_replication_scanner_finish(m, v1908)
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v716
	F_pgstat_report_activity(m, int32(3), v716)
	mBase = m.M
	v1918 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[66])))
	if v1918 != 0 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v1919 = int32(15)
	goto L536
L535:
	;
	v1919 = int32(14)
	goto L536
L536:
	;
	v1921 = F_errstart(m, v1919, int32(0))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	if v1921 != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+400)) = v716
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_61), v726+int32(400))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1935)+24))
	goto L543
L541:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2095), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L540
L543:
	;
	if (v1936-int32(7))&int32(-9) == int32(0) {
		goto L272
	} else {
		goto L544
	}
L544:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v1944 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L1
	} else {
		goto L548
	}
L546:
	;
	goto L547
L547:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_62))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L1
	} else {
		goto L549
	}
L548:
	;
	goto L547
L549:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_63))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_64))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	switch v1957 - int32(448) {
	case 0:
		goto L561
	case 1:
		goto L552
	case 2:
		goto L559
	case 3:
		goto L558
	case 4:
		goto L557
	case 5:
		goto L556
	case 6:
		goto L560
	case 7:
		goto L555
	case 8:
		goto L554
	default:
		goto L553
	}
L552:
	;
	v4612 = int32(_a_F_PostgresMainLoopOnce_65)
	F_PreventInTransactionBlock(m, int32(1), v4612)
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L1
	} else {
		goto L1206
	}
L553:
	;
	if v1957 == int32(159) {
		goto L248
	} else {
		goto L1202
	}
L554:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_66))
	mBase = m.M
	v4182 = m.ExcPending
	if v4182 != 0 {
		goto L1
	} else {
		goto L1113
	}
L555:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_67))
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L1
	} else {
		goto L1084
	}
L556:
	;
	v3327 = int32(_a_F_PostgresMainLoopOnce_68)
	F_PreventInTransactionBlock(m, int32(1), v3327)
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L1
	} else {
		goto L925
	}
L557:
	;
	v2935 = int32(0)
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+8))
	if v2937 == v2935 {
		v3095 = v1
		v3100 = v1
		v3101 = v1
		v3111 = v2935
		goto L827
	} else {
		goto L828
	}
L558:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+4))
	v2927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2925)+8)))
	F_ReplicationSlotDrop(m, v2926, (v2927^int32(-1))&int32(1))
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L1
	} else {
		goto L826
	}
L559:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v2277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67]))) = v2277
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+20))
	if v2279 == v2277 {
		v2612 = v1
		v2615 = v1
		v2622 = v1
		v2627 = v1
		goto L650
	} else {
		goto L651
	}
L560:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[68]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67]))) = int64(0)
	v2124 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L1
	} else {
		goto L609
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67]))) = int32(0)
	v1963 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1964 = *(*int64)(unsafe.Add(mBase, uint32(v1963)))
	goto L562
L562:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v726)+32)) = v1964
	v1968 = int32(32)
	v1972 = F_pg_snprintf(m, v726+int32(_a_F_PostgresMainLoopOnce_69), v1968, int32(_a_F_PostgresMainLoopOnce_70), v726+v1968)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])))
	if v1977 == int32(1) {
		goto L565
	} else {
		goto L566
	}
L564:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = uint8(v1987)
	if v1987 != 0 {
		goto L569
	} else {
		goto L570
	}
L565:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1982)+316))
	v1985 = base.B2i32(v1983 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = uint8(v1985)
	v1987 = v1985
	goto L567
L566:
	;
	v1987 = int32(0)
	goto L567
L567:
	;
	goto L564
L568:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+20)) = uint32(v2032)
	v2035 = int64(base.Ui64(v2032) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+16)) = uint32(v2035)
	v2043 = F_pg_snprintf(m, v726+int32(464), int32(64), int32(_a_F_PostgresMainLoopOnce_71), v726+int32(16))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L584
	}
L569:
	;
	v1992 = F_GetWalRcvFlushRecPtr(m, int32(0), v726+int32(_a_F_PostgresMainLoopOnce_72))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L1
	} else {
		goto L572
	}
L570:
	;
	goto L571
L571:
	;
	v2006 = v726 + int32(440)
	v2008 = int32(_a_F_PostgresMainLoopOnce_73)
	v2009 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v2010 = int64(0)
	v2013 = base.AtomicRmwCmpxchg64(m, v2009, int32(280), v2010, v2010)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[73])) = v2013
	v2017 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v2021 = base.AtomicRmwCmpxchg64(m, v2017, int32(272), v2010, v2010)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[74])) = v2021
	if v2006 != 0 {
		goto L581
	} else {
		goto L582
	}
L572:
	;
	v1996 = F_GetXLogReplayRecPtr(m, v726+int32(464))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+440)) = v1998
	if base.Ui64(v1996) < base.Ui64(v1992) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	v2001 = v1992
	goto L576
L575:
	;
	v2001 = v1996
	goto L576
L576:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75])))
	if v1998 == v2002 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2004 = v2001
	goto L579
L578:
	;
	v2004 = v1996
	goto L579
L579:
	;
	v2032 = v2004
	goto L568
L580:
	;
	v2032 = v2028
	goto L568
L581:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v2006))) = v2025
	goto L583
L582:
	;
	goto L583
L583:
	;
	v2028 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[73]))
	goto L580
L584:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[56]))
	if v2047 != 0 {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L588
	}
L586:
	;
	v2060 = int32(0)
	goto L587
L587:
	;
	v2062 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L1
	} else {
		goto L592
	}
L588:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[56]))
	v2054 = F_get_database_name(m, v2053)
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v2056 = F_MemoryContextStrdup(m, v2049, v2054)
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v2060 = v2056
	goto L587
L592:
	;
	v2065 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	F_TupleDescInitBuiltinEntry(m, v2065, int32(1), int32(_a_F_PostgresMainLoopOnce_74), int32(25))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	F_TupleDescInitBuiltinEntry(m, v2065, int32(2), int32(_a_F_PostgresMainLoopOnce_43), int32(20))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	F_TupleDescInitBuiltinEntry(m, v2065, int32(3), int32(_a_F_PostgresMainLoopOnce_75), int32(25))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_TupleDescInitBuiltinEntry(m, v2065, int32(4), int32(_a_F_PostgresMainLoopOnce_76), int32(25))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v2088 = F_begin_tup_output_tupdesc(m, v2062, v2065, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v2092 = F_cstring_to_text(m, v726+int32(_a_F_PostgresMainLoopOnce_69))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75]))) = v2092
	v2095 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v726)+440)))
	v2096 = F_Int64GetDatum(m, v2095)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[76]))) = v2096
	v2101 = F_cstring_to_text(m, v726+int32(464))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[77]))) = v2101
	if v2060 != 0 {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	F_do_tup_output(m, v2088, v726+int32(_a_F_PostgresMainLoopOnce_72), v726+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L607
	}
L603:
	;
	v2104 = F_cstring_to_text(m, v2060)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L606
	}
L604:
	;
	goto L605
L605:
	;
	v2107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[78]))) = uint8(v2107)
	goto L602
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[79]))) = v2104
	goto L602
L607:
	;
	F_end_tup_output(m, v2088)
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_79)
	goto L247
L609:
	;
	F_TupleDescInitBuiltinEntry(m, v2124, int32(1), int32(_a_F_PostgresMainLoopOnce_80), int32(25))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	F_TupleDescInitBuiltinEntry(m, v2124, int32(2), int32(_a_F_PostgresMainLoopOnce_81), int32(25))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	F_TupleDescInitBuiltinEntry(m, v2124, int32(3), int32(_a_F_PostgresMainLoopOnce_82), int32(20))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	v2141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+462)) = uint8(v2141)
	v2143 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+460)) = uint16(v2143)
	v2146 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[80]))
	v2150 = F_LWLockAcquire(m, v2146+int32(_a_F_PostgresMainLoopOnce_83), v2141)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2118)+4))
	v2154 = F_SearchNamedReplicationSlot(m, v2152, int32(0))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L616
	}
L614:
	;
	v2262 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L1
	} else {
		goto L646
	}
L615:
	;
	v2165 = base.AtomicRmwXchg32(m, v2154, int32(0), int32(1))
	if v2165 != 0 {
		goto L622
	} else {
		goto L623
	}
L616:
	;
	if v2154 != 0 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2154)+4)))
	if v2156 != 0 {
		goto L615
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[80]))
	F_LWLockRelease(m, v2158+int32(_a_F_PostgresMainLoopOnce_83))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L621
	}
L620:
	;
	goto L619
L621:
	;
	goto L614
L622:
	;
	F_s_lock(m, v2154, int32(_a_F_PostgresMainLoopOnce_14), int32(511), int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L1
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	base.MemoryCopy(m, v726+int32(_a_F_PostgresMainLoopOnce_69), v2154, int32(88))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2154)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+456)) = v2175
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2154)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+448)) = v2177
	v2179 = *(*int64)(unsafe.Add(mBase, uint32(v2154)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v726)+440)) = v2179
	v2181 = *(*int64)(unsafe.Add(mBase, uint32(v2154)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v726)+432)) = v2181
	base.MemoryCopy(m, v726+int32(464), v2154+int32(112), int32(176))
	v2189 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2154))), uint32(v2189))
	v2193 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[80]))
	F_LWLockRelease(m, v2193+int32(_a_F_PostgresMainLoopOnce_83))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L1
	} else {
		goto L626
	}
L625:
	;
	goto L624
L626:
	;
	if v2175 != 0 {
		goto L271
	} else {
		goto L627
	}
L627:
	;
	v2199 = F_cstring_to_text(m, int32(_a_F_PostgresMainLoopOnce_42))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v2201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+460)) = uint8(v2201)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67]))) = v2199
	if v2181 == int64(0) {
		goto L614
	} else {
		goto L629
	}
L629:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+52)) = uint32(v2181)
	v2208 = int64(base.Ui64(v2181) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+48)) = uint32(v2208)
	v2211 = v726 + int32(_a_F_PostgresMainLoopOnce_72)
	v2216 = F_pg_snprintf(m, v2211, int32(64), int32(_a_F_PostgresMainLoopOnce_71), v726+int32(48))
	mBase = m.M
	v2217 = m.ExcPending
	if v2217 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	v2218 = F_cstring_to_text(m, v2211)
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	v2220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+461)) = uint8(v2220)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81]))) = v2218
	v2223 = *(*int64)(unsafe.Add(mBase, uint32(v726)+432))
	if v2223 == int64(0) {
		goto L614
	} else {
		goto L632
	}
L632:
	;
	v2228 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])))
	if v2228 == int32(1) {
		goto L635
	} else {
		goto L636
	}
L633:
	;
	v2247 = F_readTimeLineHistory(m, v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L1
	} else {
		goto L643
	}
L634:
	;
	if v2238 != 0 {
		goto L638
	} else {
		goto L639
	}
L635:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v2233)+316))
	v2236 = base.B2i32(v2234 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = uint8(v2236)
	v2238 = v2236
	goto L637
L636:
	;
	v2238 = int32(0)
	goto L637
L637:
	;
	goto L634
L638:
	;
	v2239 = F_GetXLogReplayRecPtr(m, v2211)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L1
	} else {
		goto L641
	}
L639:
	;
	goto L640
L640:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2243)+308))
	goto L642
L641:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75])))
	v2246 = v2241
	goto L633
L642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75]))) = v2244
	v2246 = v2244
	goto L633
L643:
	;
	v2249 = *(*int64)(unsafe.Add(mBase, uint32(v726)+432))
	v2250 = F_tliOfPointInHistory(m, v2249, v2247)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	v2253 = F_Int64GetDatum(m, base.I64_extend_i32_u(v2250))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	v2255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+462)) = uint8(v2255)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[68]))) = v2253
	goto L614
L646:
	;
	v2265 = F_begin_tup_output_tupdesc(m, v2262, v2124, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	F_do_tup_output(m, v2265, v726+int32(_a_F_PostgresMainLoopOnce_78), v726+int32(460))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	F_end_tup_output(m, v2265)
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_85)
	goto L247
L650:
	;
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	if v2634 == int32(0) {
		goto L742
	} else {
		goto L743
	}
L651:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2279)+4))
	if v2282 <= int32(0) {
		v2612 = v1
		v2615 = v1
		v2622 = v1
		v2627 = v1
		goto L650
	} else {
		goto L652
	}
L652:
	;
	v2286 = int32(0)
	v2301 = v1
	v2303 = v1
	v2304 = v1
	v2307 = v1
	v2311 = v1
	v2312 = v1
	v2313 = v1
	v2316 = v1
	goto L653
L653:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2279)+12))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2323+v2286<<(uint(int32(2))%32))))
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(v2327)+8))
	v2329 = int32(_a_F_PostgresMainLoopOnce_57)
	v2332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328))))
	v2335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[82])))
	if base.B2i32(v2332 == int32(0))|base.B2i32(v2332 != v2335) != 0 {
		v2353 = v2332
		v2354 = v2335
		goto L657
	} else {
		goto L658
	}
L654:
	;
	v2612 = v2585
	v2615 = v2587
	v2622 = v2589
	v2627 = v2592
	goto L650
L655:
	;
	v2594 = v2286 + int32(1)
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v2279)+4))
	if v2594 < v2595 {
		v2286 = v2594
		v2301 = v2585
		v2303 = v2586
		v2304 = v2587
		v2307 = v2588
		v2311 = v2589
		v2312 = v2590
		v2313 = v2591
		v2316 = v2592
		goto L653
	} else {
		goto L740
	}
L656:
	;
	if v2353-v2354 == int32(0) {
		goto L663
	} else {
		goto L664
	}
L657:
	;
	goto L656
L658:
	;
	v2338 = v2328
	v2339 = v2329
	goto L659
L659:
	;
	v2342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2339)+1)))
	v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2338)+1)))
	if v2343 == int32(0) {
		v2353 = v2343
		v2354 = v2342
		goto L657
	} else {
		goto L661
	}
L660:
	;
	v2353 = v2343
	v2354 = v2342
	goto L657
L661:
	;
	v2346 = int32(1)
	if v2343 == v2342 {
		v2338 = v2338 + v2346
		v2339 = v2339 + v2346
		goto L659
	} else {
		goto L662
	}
L662:
	;
	goto L660
L663:
	;
	if v2307&int32(1) != 0 {
		goto L270
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	v2479 = int32(_a_F_PostgresMainLoopOnce_39)
	v2482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328))))
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[83])))
	if base.B2i32(v2482 == int32(0))|base.B2i32(v2482 != v2485) != 0 {
		v2503 = v2482
		v2504 = v2485
		goto L704
	} else {
		goto L705
	}
L666:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	if v2360 != int32(1) {
		goto L270
	} else {
		goto L667
	}
L667:
	;
	v2363 = F_defGetString(m, v2327)
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	v2365 = int32(_a_F_PostgresMainLoopOnce_55)
	v2368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2363))))
	v2371 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[84])))
	if base.B2i32(v2368 == int32(0))|base.B2i32(v2368 != v2371) != 0 {
		v2389 = v2368
		v2390 = v2371
		goto L670
	} else {
		goto L671
	}
L669:
	;
	if v2389-v2390 == int32(0) {
		goto L676
	} else {
		goto L677
	}
L670:
	;
	goto L669
L671:
	;
	v2374 = v2363
	v2375 = v2365
	goto L672
L672:
	;
	v2378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2375)+1)))
	v2379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2374)+1)))
	if v2379 == int32(0) {
		v2389 = v2379
		v2390 = v2378
		goto L670
	} else {
		goto L674
	}
L673:
	;
	v2389 = v2379
	v2390 = v2378
	goto L670
L674:
	;
	v2382 = int32(1)
	if v2379 == v2378 {
		v2374 = v2374 + v2382
		v2375 = v2375 + v2382
		goto L672
	} else {
		goto L675
	}
L675:
	;
	goto L673
L676:
	;
	v2585 = v2301
	v2586 = v2303
	v2587 = v2304
	v2588 = int32(1)
	v2589 = int32(0)
	v2590 = v2312
	v2591 = v2313
	v2592 = v2316
	goto L655
L677:
	;
	goto L678
L678:
	;
	v2396 = int32(1)
	v2397 = int32(_a_F_PostgresMainLoopOnce_54)
	v2400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2363))))
	v2403 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[85])))
	if base.B2i32(v2400 == int32(0))|base.B2i32(v2400 != v2403) != 0 {
		v2421 = v2400
		v2422 = v2403
		goto L680
	} else {
		goto L681
	}
L679:
	;
	if v2421-v2422 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L680:
	;
	goto L679
L681:
	;
	v2406 = v2363
	v2407 = v2397
	goto L682
L682:
	;
	v2410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2407)+1)))
	v2411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2406)+1)))
	if v2411 == int32(0) {
		v2421 = v2411
		v2422 = v2410
		goto L680
	} else {
		goto L684
	}
L683:
	;
	v2421 = v2411
	v2422 = v2410
	goto L680
L684:
	;
	v2414 = int32(1)
	if v2411 == v2410 {
		v2406 = v2406 + v2414
		v2407 = v2407 + v2414
		goto L682
	} else {
		goto L685
	}
L685:
	;
	goto L683
L686:
	;
	v2585 = v2301
	v2586 = v2303
	v2587 = v2304
	v2588 = v2396
	v2589 = int32(1)
	v2590 = v2312
	v2591 = v2313
	v2592 = v2316
	goto L655
L687:
	;
	goto L688
L688:
	;
	v2427 = int32(_a_F_PostgresMainLoopOnce_53)
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2363))))
	v2433 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[86])))
	if base.B2i32(v2430 == int32(0))|base.B2i32(v2430 != v2433) != 0 {
		v2451 = v2430
		v2452 = v2433
		goto L690
	} else {
		goto L691
	}
L689:
	;
	if v2451-v2452 == int32(0) {
		goto L696
	} else {
		goto L697
	}
L690:
	;
	goto L689
L691:
	;
	v2436 = v2363
	v2437 = v2427
	goto L692
L692:
	;
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2437)+1)))
	v2441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2436)+1)))
	if v2441 == int32(0) {
		v2451 = v2441
		v2452 = v2440
		goto L690
	} else {
		goto L694
	}
L693:
	;
	v2451 = v2441
	v2452 = v2440
	goto L690
L694:
	;
	v2444 = int32(1)
	if v2441 == v2440 {
		v2436 = v2436 + v2444
		v2437 = v2437 + v2444
		goto L692
	} else {
		goto L695
	}
L695:
	;
	goto L693
L696:
	;
	v2585 = v2301
	v2586 = v2303
	v2587 = v2304
	v2588 = v2396
	v2589 = int32(2)
	v2590 = v2312
	v2591 = v2313
	v2592 = v2316
	goto L655
L697:
	;
	goto L698
L698:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L699
	}
L699:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2327)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+200)) = v2363
	*(*int32)(unsafe.Add(mBase, uint32(v726)+196)) = v2464
	*(*int32)(unsafe.Add(mBase, uint32(v726)+192)) = int32(_a_F_PostgresMainLoopOnce_86)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_87), v726+int32(192))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1152), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	if v2503-v2504 == int32(0) {
		goto L710
	} else {
		goto L711
	}
L704:
	;
	goto L703
L705:
	;
	v2488 = v2328
	v2489 = v2479
	goto L706
L706:
	;
	v2492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2489)+1)))
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488)+1)))
	if v2493 == int32(0) {
		v2503 = v2493
		v2504 = v2492
		goto L704
	} else {
		goto L708
	}
L707:
	;
	v2503 = v2493
	v2504 = v2492
	goto L704
L708:
	;
	v2496 = int32(1)
	if v2493 == v2492 {
		v2488 = v2488 + v2496
		v2489 = v2489 + v2496
		goto L706
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	if v2312&int32(1) != 0 {
		goto L269
	} else {
		goto L713
	}
L711:
	;
	goto L712
L712:
	;
	v2514 = int32(_a_F_PostgresMainLoopOnce_37)
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328))))
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87])))
	if base.B2i32(v2517 == int32(0))|base.B2i32(v2517 != v2520) != 0 {
		v2538 = v2517
		v2539 = v2520
		goto L717
	} else {
		goto L718
	}
L713:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	if v2510 != 0 {
		goto L269
	} else {
		goto L714
	}
L714:
	;
	v2512 = F_defGetBoolean(m, v2327)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	v2585 = v2301
	v2586 = v2303
	v2587 = v2304
	v2588 = v2307
	v2589 = v2311
	v2590 = int32(1)
	v2591 = v2313
	v2592 = v2512
	goto L655
L716:
	;
	if v2538-v2539 == int32(0) {
		goto L723
	} else {
		goto L724
	}
L717:
	;
	goto L716
L718:
	;
	v2523 = v2328
	v2524 = v2514
	goto L719
L719:
	;
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2524)+1)))
	v2528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2523)+1)))
	if v2528 == int32(0) {
		v2538 = v2528
		v2539 = v2527
		goto L717
	} else {
		goto L721
	}
L720:
	;
	v2538 = v2528
	v2539 = v2527
	goto L717
L721:
	;
	v2531 = int32(1)
	if v2528 == v2527 {
		v2523 = v2523 + v2531
		v2524 = v2524 + v2531
		goto L719
	} else {
		goto L722
	}
L722:
	;
	goto L720
L723:
	;
	if v2303 != 0 {
		goto L268
	} else {
		goto L726
	}
L724:
	;
	goto L725
L725:
	;
	v2549 = int32(_a_F_PostgresMainLoopOnce_89)
	v2552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328))))
	v2555 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[88])))
	if base.B2i32(v2552 == int32(0))|base.B2i32(v2552 != v2555) != 0 {
		v2573 = v2552
		v2574 = v2555
		goto L730
	} else {
		goto L731
	}
L726:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	if v2543 != int32(1) {
		goto L268
	} else {
		goto L727
	}
L727:
	;
	v2547 = F_defGetBoolean(m, v2327)
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v2585 = v2301
	v2586 = int32(1)
	v2587 = v2547
	v2588 = v2307
	v2589 = v2311
	v2590 = v2312
	v2591 = v2313
	v2592 = v2316
	goto L655
L729:
	;
	if v2573-v2574 != 0 {
		goto L266
	} else {
		goto L736
	}
L730:
	;
	goto L729
L731:
	;
	v2558 = v2328
	v2559 = v2549
	goto L732
L732:
	;
	v2562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2559)+1)))
	v2563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2558)+1)))
	if v2563 == int32(0) {
		v2573 = v2563
		v2574 = v2562
		goto L730
	} else {
		goto L734
	}
L733:
	;
	v2573 = v2563
	v2574 = v2562
	goto L730
L734:
	;
	v2566 = int32(1)
	if v2563 == v2562 {
		v2558 = v2558 + v2566
		v2559 = v2559 + v2566
		goto L732
	} else {
		goto L735
	}
L735:
	;
	goto L733
L736:
	;
	if v2313&int32(1) != 0 {
		goto L267
	} else {
		goto L737
	}
L737:
	;
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	if v2578 != int32(1) {
		goto L267
	} else {
		goto L738
	}
L738:
	;
	v2582 = F_defGetBoolean(m, v2327)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	v2585 = v2582
	v2586 = v2303
	v2587 = v2304
	v2588 = v2307
	v2589 = v2311
	v2590 = v2312
	v2591 = int32(1)
	v2592 = v2316
	goto L655
L740:
	;
	goto L654
L741:
	;
	v2849 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v2850 = *(*int64)(unsafe.Add(mBase, uint32(v2849)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+84)) = uint32(v2850)
	v2853 = int64(base.Ui64(v2850) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+80)) = uint32(v2853)
	v2856 = v726 + int32(464)
	v2861 = F_pg_snprintf(m, v2856, int32(64), int32(_a_F_PostgresMainLoopOnce_71), v726+int32(80))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L1
	} else {
		goto L803
	}
L742:
	;
	v2637 = int32(0)
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+4))
	v2640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2276)+16)))
	F_ReplicationSlotCreate(m, v2638, v2637, v2640<<(uint(int32(1))%32)&int32(2), v2637, v2637, v2637)
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L1
	} else {
		goto L745
	}
L743:
	;
	goto L744
L744:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L1
	} else {
		goto L751
	}
L745:
	;
	if v2627&int32(1) == int32(0) {
		v2847 = v2637
		goto L741
	} else {
		goto L746
	}
L746:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v2655 = m.ExcPending
	if v2655 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	v2658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2276)+16)))
	if v2658 != 0 {
		v2847 = v2637
		goto L741
	} else {
		goto L749
	}
L749:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L1
	} else {
		goto L750
	}
L750:
	;
	v2847 = v2637
	goto L741
L751:
	;
	v2663 = int32(0)
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+4))
	v2665 = int32(1)
	v2668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2276)+16)))
	if v2668 != 0 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v2669 = int32(2)
	goto L754
L753:
	;
	v2669 = v2665
	goto L754
L754:
	;
	v2670 = int32(1)
	F_ReplicationSlotCreate(m, v2664, v2665, v2669, v2615&v2670, v2612&v2670, int32(0))
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	switch v2622 {
	case 0:
		goto L758
	default:
		v2727 = int32(0)
		goto L756
	case 2:
		goto L757
	}
L756:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[77]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[76]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75]))) = int32(1032)
	v2741 = F_CreateInitDecodingContext(m, v2728, v2727, int64(0), v726+int32(_a_F_PostgresMainLoopOnce_72), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L771
	}
L757:
	;
	v2703 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2703)+24))
	goto L764
L758:
	;
	v2678 = int32(1)
	v2680 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v2680)+24))
	goto L759
L759:
	;
	if base.B2i32(base.Ui32(v2678) < base.Ui32(v2681)) == int32(0) {
		v2727 = v2678
		goto L756
	} else {
		goto L760
	}
L760:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L1
	} else {
		goto L761
	}
L761:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+96)) = int32(_a_F_PostgresMainLoopOnce_90)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_91), v726+int32(96))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		goto L1
	} else {
		goto L762
	}
L762:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1258), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L1
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
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2704)) == int32(0) {
		goto L265
	} else {
		goto L765
	}
L765:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90]))
	if v2710 != int32(2) {
		goto L264
	} else {
		goto L766
	}
L766:
	;
	v2714 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[91])))
	if v2714 == int32(0) {
		goto L263
	} else {
		goto L767
	}
L767:
	;
	v2717 = int32(1)
	v2719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[92])))
	if v2719 == v2717 {
		goto L262
	} else {
		goto L768
	}
L768:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2723)+28))
	goto L769
L769:
	;
	if int32(1) < v2724 {
		goto L261
	} else {
		goto L770
	}
L770:
	;
	v2727 = v2717
	goto L756
L771:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])) = int64(0)
	F_DecodingContextFindStartpoint(m, v2741)
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	switch v2622 {
	case 0:
		goto L775
	default:
		v2837 = v2663
		goto L773
	case 2:
		goto L774
	}
L773:
	;
	F_FreeDecodingContext(m, v2741)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L1
	} else {
		goto L800
	}
L774:
	;
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2741)+16))
	v2828 = F_SnapBuildInitialSnapshot(m, v2827)
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		goto L1
	} else {
		goto L798
	}
L775:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2741)+16))
	v2749 = m.G0
	v2751 = v2749 - int32(16)
	m.G0 = v2751
	v2754 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v2754)+24))
	goto L778
L776:
	;
	v2837 = v2779
	goto L773
L777:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L1
	} else {
		goto L795
	}
L778:
	;
	if base.B2i32(v2755 != int32(0)) == int32(0) {
		goto L779
	} else {
		goto L780
	}
L779:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[52]))
	if v2761 != 0 {
		goto L777
	} else {
		goto L782
	}
L780:
	;
	goto L781
L781:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L1
	} else {
		goto L792
	}
L782:
	;
	v2763 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51])) = uint8(v2763)
	v2767 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[52])) = v2767
	F_StartTransactionCommand(m)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L1
	} else {
		goto L783
	}
L783:
	;
	v2772 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[91])) = uint8(v2772)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90])) = int32(2)
	v2777 = F_SnapBuildInitialSnapshot(m, v2748)
	mBase = m.M
	v2778 = m.ExcPending
	if v2778 != 0 {
		goto L1
	} else {
		goto L784
	}
L784:
	;
	v2779 = F_ExportSnapshot(m, v2777)
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L1
	} else {
		goto L785
	}
L785:
	;
	v2783 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L786
	}
L786:
	;
	if v2783 != 0 {
		goto L787
	} else {
		goto L788
	}
L787:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v2777)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2751)+4)) = v2785
	*(*int32)(unsafe.Add(mBase, uint32(v2751))) = v2779
	F_errmsg_plural(m, int32(_a_F_PostgresMainLoopOnce_93), int32(_a_F_PostgresMainLoopOnce_94), v2785, v2751)
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L1
	} else {
		goto L790
	}
L788:
	;
	goto L789
L789:
	;
	m.G0 = v2751 + int32(16)
	goto L776
L790:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(571), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	goto L789
L792:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_96), int32(0))
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(545), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2813 = m.ExcPending
	if v2813 != 0 {
		goto L1
	} else {
		goto L794
	}
L794:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L795:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_97), int32(0))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(548), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2826 = m.ExcPending
	if v2826 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L798:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	F_RestoreTransactionSnapshot(m, v2828, v2831)
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L1
	} else {
		goto L799
	}
L799:
	;
	v2837 = v2663
	goto L773
L800:
	;
	v2840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2276)+16)))
	if v2840 != 0 {
		v2847 = v2837
		goto L741
	} else {
		goto L801
	}
L801:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L1
	} else {
		goto L802
	}
L802:
	;
	v2847 = v2837
	goto L741
L803:
	;
	v2864 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		goto L1
	} else {
		goto L804
	}
L804:
	;
	v2867 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L1
	} else {
		goto L805
	}
L805:
	;
	F_TupleDescInitBuiltinEntry(m, v2867, int32(1), int32(_a_F_PostgresMainLoopOnce_98), int32(25))
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	F_TupleDescInitBuiltinEntry(m, v2867, int32(2), int32(_a_F_PostgresMainLoopOnce_99), int32(25))
	mBase = m.M
	v2878 = m.ExcPending
	if v2878 != 0 {
		goto L1
	} else {
		goto L807
	}
L807:
	;
	F_TupleDescInitBuiltinEntry(m, v2867, int32(3), int32(_a_F_PostgresMainLoopOnce_100), int32(25))
	mBase = m.M
	v2883 = m.ExcPending
	if v2883 != 0 {
		goto L1
	} else {
		goto L808
	}
L808:
	;
	F_TupleDescInitBuiltinEntry(m, v2867, int32(4), int32(_a_F_PostgresMainLoopOnce_101), int32(25))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		goto L1
	} else {
		goto L809
	}
L809:
	;
	v2890 = F_begin_tup_output_tupdesc(m, v2864, v2867, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v2896 = F_cstring_to_text(m, v2893+int32(24))
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[94]))) = v2896
	v2899 = F_cstring_to_text(m, v2856)
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[95]))) = v2899
	if v2847 != 0 {
		goto L814
	} else {
		goto L815
	}
L813:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	if v2907 != 0 {
		goto L819
	} else {
		goto L820
	}
L814:
	;
	v2902 = F_cstring_to_text(m, v2847)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L1
	} else {
		goto L817
	}
L815:
	;
	goto L816
L816:
	;
	v2905 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[96]))) = uint8(v2905)
	goto L813
L817:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[97]))) = v2902
	goto L813
L818:
	;
	F_do_tup_output(m, v2890, v726+int32(_a_F_PostgresMainLoopOnce_69), v726+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L1
	} else {
		goto L823
	}
L819:
	;
	v2908 = F_cstring_to_text(m, v2907)
	mBase = m.M
	v2909 = m.ExcPending
	if v2909 != 0 {
		goto L1
	} else {
		goto L822
	}
L820:
	;
	goto L821
L821:
	;
	v2911 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[78]))) = uint8(v2911)
	goto L818
L822:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[98]))) = v2908
	goto L818
L823:
	;
	F_end_tup_output(m, v2890)
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L1
	} else {
		goto L825
	}
L825:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_86)
	goto L247
L826:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_102)
	goto L247
L827:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+464)) = uint8(v3100)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[94]))) = uint8(v3101)
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+4))
	v3115 = m.G0
	v3117 = v3115 - int32(1072)
	m.G0 = v3117
	F_ReplicationSlotAcquire(m, v3114, int32(0), int32(1))
	mBase = m.M
	v3122 = m.ExcPending
	if v3122 != 0 {
		goto L1
	} else {
		goto L862
	}
L828:
	;
	v2940 = int32(0)
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v2937)+4))
	if v2941 <= v2940 {
		v3095 = v1
		v3100 = v1
		v3101 = v1
		v3111 = v2940
		goto L827
	} else {
		goto L829
	}
L829:
	;
	v2945 = int32(0)
	v2966 = v1
	v2970 = v1
	v2971 = v1
	v2972 = v1
	goto L830
L830:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2937)+12))
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v2982+v2945<<(uint(int32(2))%32))))
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2986)+8))
	v2988 = int32(_a_F_PostgresMainLoopOnce_89)
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987))))
	v2994 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[88])))
	if base.B2i32(v2991 == int32(0))|base.B2i32(v2991 != v2994) != 0 {
		v3012 = v2991
		v3013 = v2994
		goto L834
	} else {
		goto L835
	}
L831:
	;
	if v3054&int32(1) != 0 {
		goto L856
	} else {
		goto L857
	}
L832:
	;
	v3059 = v2945 + int32(1)
	v3060 = *(*int32)(unsafe.Add(mBase, uint32(v2937)+4))
	if v3059 < v3060 {
		v2945 = v3059
		v2966 = v3054
		v2970 = v3055
		v2971 = v3056
		v2972 = v3057
		goto L830
	} else {
		goto L855
	}
L833:
	;
	if v3012-v3013 == int32(0) {
		goto L840
	} else {
		goto L841
	}
L834:
	;
	goto L833
L835:
	;
	v2997 = v2987
	v2998 = v2988
	goto L836
L836:
	;
	v3001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2998)+1)))
	v3002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2997)+1)))
	if v3002 == int32(0) {
		v3012 = v3002
		v3013 = v3001
		goto L834
	} else {
		goto L838
	}
L837:
	;
	v3012 = v3002
	v3013 = v3001
	goto L834
L838:
	;
	v3005 = int32(1)
	if v3002 == v3001 {
		v2997 = v2997 + v3005
		v2998 = v2998 + v3005
		goto L836
	} else {
		goto L839
	}
L839:
	;
	goto L837
L840:
	;
	if v2970&int32(1) != 0 {
		goto L260
	} else {
		goto L843
	}
L841:
	;
	goto L842
L842:
	;
	v3022 = int32(_a_F_PostgresMainLoopOnce_37)
	v3025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987))))
	v3028 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87])))
	if base.B2i32(v3025 == int32(0))|base.B2i32(v3025 != v3028) != 0 {
		v3046 = v3025
		v3047 = v3028
		goto L846
	} else {
		goto L847
	}
L843:
	;
	v3020 = F_defGetBoolean(m, v2986)
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	v3054 = v2966
	v3055 = int32(1)
	v3056 = v3020
	v3057 = v2972
	goto L832
L845:
	;
	if v3046-v3047 != 0 {
		goto L258
	} else {
		goto L852
	}
L846:
	;
	goto L845
L847:
	;
	v3031 = v2987
	v3032 = v3022
	goto L848
L848:
	;
	v3035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3032)+1)))
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3031)+1)))
	if v3036 == int32(0) {
		v3046 = v3036
		v3047 = v3035
		goto L846
	} else {
		goto L850
	}
L849:
	;
	v3046 = v3036
	v3047 = v3035
	goto L846
L850:
	;
	v3039 = int32(1)
	if v3036 == v3035 {
		v3031 = v3031 + v3039
		v3032 = v3032 + v3039
		goto L848
	} else {
		goto L851
	}
L851:
	;
	goto L849
L852:
	;
	if v2966&int32(1) != 0 {
		goto L259
	} else {
		goto L853
	}
L853:
	;
	v3052 = F_defGetBoolean(m, v2986)
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L1
	} else {
		goto L854
	}
L854:
	;
	v3054 = int32(1)
	v3055 = v2970
	v3056 = v2971
	v3057 = v3052
	goto L832
L855:
	;
	goto L831
L856:
	;
	v3067 = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	goto L858
L857:
	;
	v3067 = int32(0)
	goto L858
L858:
	;
	if v3055&int32(1) != 0 {
		goto L859
	} else {
		goto L860
	}
L859:
	;
	v3073 = v726 + int32(464)
	goto L861
L860:
	;
	v3073 = int32(0)
	goto L861
L861:
	;
	v3095 = v3067
	v3100 = v3056
	v3101 = v3057
	v3111 = v3073
	goto L827
L862:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(v3124)+88))
	if v3125 != 0 {
		goto L866
	} else {
		goto L867
	}
L863:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_103)
	goto L247
L864:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L1
	} else {
		goto L921
	}
L865:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L1
	} else {
		goto L916
	}
L866:
	;
	v3128 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])))
	if v3128 == int32(1) {
		goto L872
	} else {
		goto L873
	}
L867:
	;
	goto L868
L868:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L1
	} else {
		goto L912
	}
L869:
	;
	if v3095 == int32(0) {
		goto L897
	} else {
		goto L898
	}
L870:
	;
	v3178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3177)+202)))
	if v3176 == v3178 {
		v3201 = v1
		goto L869
	} else {
		goto L890
	}
L871:
	;
	if v3138 != 0 {
		goto L875
	} else {
		goto L876
	}
L872:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v3133)+316))
	v3136 = base.B2i32(v3134 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = uint8(v3136)
	v3138 = v3136
	goto L874
L873:
	;
	v3138 = int32(0)
	goto L874
L874:
	;
	goto L871
L875:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3140)+201)))
	if v3141 != 0 {
		goto L865
	} else {
		goto L878
	}
L876:
	;
	goto L877
L877:
	;
	if v3111 == int32(0) {
		v3201 = v1
		goto L869
	} else {
		goto L885
	}
L878:
	;
	if v3111 == int32(0) {
		v3201 = v1
		goto L869
	} else {
		goto L879
	}
L879:
	;
	v3145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3111))))
	if v3145 != int32(1) {
		v3176 = int32(0)
		v3177 = v3140
		goto L870
	} else {
		goto L880
	}
L880:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L1
	} else {
		goto L881
	}
L881:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_104), int32(0))
	mBase = m.M
	v3158 = m.ExcPending
	if v3158 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(913), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L885:
	;
	v3166 = int32(1)
	v3168 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3111))))
	if v3169 != v3166 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v3176 = int32(0)
	v3177 = v3168
	goto L870
L887:
	;
	goto L888
L888:
	;
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(v3168)+92))
	if v3173 == int32(2) {
		goto L864
	} else {
		goto L889
	}
L889:
	;
	v3176 = v3166
	v3177 = v3168
	goto L870
L890:
	;
	v3180 = int32(1)
	v3183 = base.AtomicRmwXchg32(m, v3177, int32(0), v3180)
	if v3183 != 0 {
		goto L891
	} else {
		goto L892
	}
L891:
	;
	v3185 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	F_s_lock(m, v3185, int32(_a_F_PostgresMainLoopOnce_105), int32(929), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L1
	} else {
		goto L894
	}
L892:
	;
	goto L893
L893:
	;
	v3192 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3192)+202)) = uint8(v3193)
	v3195 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3192))), uint32(v3195))
	v3201 = v3180
	goto L869
L894:
	;
	goto L893
L895:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		goto L1
	} else {
		goto L911
	}
L896:
	;
	v3231 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3234 = base.AtomicRmwXchg32(m, v3231, int32(0), int32(1))
	if v3234 != 0 {
		goto L905
	} else {
		goto L906
	}
L897:
	;
	if v3201 == int32(0) {
		goto L895
	} else {
		goto L904
	}
L898:
	;
	v3205 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3205)+136)))
	v3207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3095))))
	if v3206 == v3207 {
		goto L897
	} else {
		goto L899
	}
L899:
	;
	v3211 = base.AtomicRmwXchg32(m, v3205, int32(0), int32(1))
	if v3211 != 0 {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	F_s_lock(m, v3213, int32(_a_F_PostgresMainLoopOnce_105), int32(939), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L903
	}
L901:
	;
	goto L902
L902:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3095))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3220)+136)) = uint8(v3221)
	v3223 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3220))), uint32(v3223))
	goto L896
L903:
	;
	goto L902
L904:
	;
	goto L896
L905:
	;
	F_s_lock(m, v3231, int32(_a_F_PostgresMainLoopOnce_105), int32(1107), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L1
	} else {
		goto L908
	}
L906:
	;
	goto L907
L907:
	;
	v3240 = int32(_a_F_PostgresMainLoopOnce_108)
	v3241 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3242 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3241)+12)) = uint16(v3242)
	v3244 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3231))), uint32(v3244))
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+16)) = int32(_a_F_PostgresMainLoopOnce_109)
	v3250 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+20)) = v3250 + int32(24)
	v3255 = v3117 + int32(48)
	v3259 = F_pg_sprintf(m, v3255, int32(_a_F_PostgresMainLoopOnce_110), v3117+int32(16))
	mBase = m.M
	v3260 = m.ExcPending
	if v3260 != 0 {
		goto L1
	} else {
		goto L909
	}
L908:
	;
	goto L907
L909:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	F_SaveSlotToPath(m, v3262, v3255, int32(21))
	mBase = m.M
	v3265 = m.ExcPending
	if v3265 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	goto L895
L911:
	;
	m.G0 = v3117 + int32(1072)
	goto L863
L912:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3278 = m.ExcPending
	if v3278 != 0 {
		goto L1
	} else {
		goto L913
	}
L913:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3117))) = int32(_a_F_PostgresMainLoopOnce_103)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_111), v3117)
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L1
	} else {
		goto L914
	}
L914:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(891), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3288 = m.ExcPending
	if v3288 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L916:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+32)) = v3114
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_112), v3117+int32(32))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_113), int32(0))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(903), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3310 = m.ExcPending
	if v3310 != 0 {
		goto L1
	} else {
		goto L920
	}
L920:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L921:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3317 = m.ExcPending
	if v3317 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_114), int32(0))
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_105), int32(925), int32(_a_F_PostgresMainLoopOnce_106))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L925:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+4))
	if v3332 == int32(0) {
		goto L926
	} else {
		goto L927
	}
L926:
	;
	v3335 = m.G0
	v3337 = v3335 - int32(144)
	m.G0 = v3337
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+116)) = int32(1031)
	v3343 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+112)) = v3343
	v3347 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99]))
	v3351 = F_XLogReaderAllocate(m, v3347, v3337+int32(112), v3343)
	mBase = m.M
	v3352 = m.ExcPending
	if v3352 != 0 {
		goto L1
	} else {
		goto L929
	}
L927:
	;
	goto L928
L928:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v3758 = m.ExcPending
	if v3758 != 0 {
		goto L1
	} else {
		goto L1038
	}
L929:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[100])) = v3351
	if v3351 != 0 {
		goto L937
	} else {
		goto L938
	}
L930:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v3756 = m.ExcPending
	if v3756 != 0 {
		goto L1
	} else {
		goto L1037
	}
L931:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3736 = m.ExcPending
	if v3736 != 0 {
		goto L1
	} else {
		goto L1034
	}
L932:
	;
	v3665 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+8))
	if v3665 != 0 {
		goto L1016
	} else {
		goto L1017
	}
L933:
	;
	v3549 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3550 = *(*int32)(unsafe.Add(mBase, uint32(v3549)+4))
	if v3550 != int32(2) {
		goto L991
	} else {
		goto L992
	}
L934:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])) = v3531
	v3535 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[102])) = uint8(v3535)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103])) = uint8(v3535)
	v3541 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])))
	if v3541 != int32(1) {
		goto L933
	} else {
		goto L989
	}
L935:
	;
	v3531 = int64(0)
	goto L934
L936:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3514 = m.ExcPending
	if v3514 != 0 {
		goto L1
	} else {
		goto L985
	}
L937:
	;
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+8))
	if v3354 != 0 {
		goto L940
	} else {
		goto L941
	}
L938:
	;
	goto L939
L939:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L1
	} else {
		goto L980
	}
L940:
	;
	v3355 = int32(1)
	F_ReplicationSlotAcquire(m, v3354, v3355, v3355)
	mBase = m.M
	v3358 = m.ExcPending
	if v3358 != 0 {
		goto L1
	} else {
		goto L943
	}
L941:
	;
	goto L942
L942:
	;
	v3365 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])))
	if v3365 == int32(1) {
		goto L946
	} else {
		goto L947
	}
L943:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3361 = *(*int32)(unsafe.Add(mBase, uint32(v3360)+88))
	if v3361 != 0 {
		goto L936
	} else {
		goto L944
	}
L944:
	;
	goto L942
L945:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = uint8(v3375)
	if v3375 != 0 {
		goto L950
	} else {
		goto L951
	}
L946:
	;
	v3370 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3371 = *(*int32)(unsafe.Add(mBase, uint32(v3370)+316))
	v3373 = base.B2i32(v3371 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = uint8(v3373)
	v3375 = v3373
	goto L948
L947:
	;
	v3375 = int32(0)
	goto L948
L948:
	;
	goto L945
L949:
	;
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	if v3421 != 0 {
		goto L965
	} else {
		goto L966
	}
L950:
	;
	v3380 = F_GetWalRcvFlushRecPtr(m, int32(0), v3337+int32(128))
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L1
	} else {
		goto L953
	}
L951:
	;
	goto L952
L952:
	;
	v3394 = v3337 + int32(124)
	v3396 = int32(_a_F_PostgresMainLoopOnce_73)
	v3397 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3398 = int64(0)
	v3401 = base.AtomicRmwCmpxchg64(m, v3397, int32(280), v3398, v3398)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[73])) = v3401
	v3405 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3409 = base.AtomicRmwCmpxchg64(m, v3405, int32(272), v3398, v3398)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[74])) = v3409
	if v3394 != 0 {
		goto L962
	} else {
		goto L963
	}
L953:
	;
	v3384 = F_GetXLogReplayRecPtr(m, v3337+int32(80))
	mBase = m.M
	v3385 = m.ExcPending
	if v3385 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	v3386 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+124)) = v3386
	if base.Ui64(v3384) < base.Ui64(v3380) {
		goto L955
	} else {
		goto L956
	}
L955:
	;
	v3389 = v3380
	goto L957
L956:
	;
	v3389 = v3384
	goto L957
L957:
	;
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+128))
	if v3386 == v3390 {
		goto L958
	} else {
		goto L959
	}
L958:
	;
	v3392 = v3389
	goto L960
L959:
	;
	v3392 = v3384
	goto L960
L960:
	;
	v3420 = v3392
	goto L949
L961:
	;
	v3420 = v3416
	goto L949
L962:
	;
	v3412 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(v3412)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v3394))) = v3413
	goto L964
L963:
	;
	goto L964
L964:
	;
	v3416 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[73]))
	goto L961
L965:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = v3421
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+124))
	if v3421 == v3424 {
		goto L968
	} else {
		goto L969
	}
L966:
	;
	goto L967
L967:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = v3477
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101])) = int64(0)
	v3483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])) = uint8(v3483)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[102])) = uint8(v3483)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103])) = uint8(v3483)
	goto L933
L968:
	;
	v3427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])) = uint8(v3427)
	goto L935
L969:
	;
	goto L970
L970:
	;
	v3430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])) = uint8(v3430)
	v3432 = F_readTimeLineHistory(m, v3424)
	mBase = m.M
	v3433 = m.ExcPending
	if v3433 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	v3436 = F_tliSwitchPoint(m, v3434, v3432, int32(_a_F_PostgresMainLoopOnce_115))
	mBase = m.M
	v3437 = m.ExcPending
	if v3437 != 0 {
		goto L1
	} else {
		goto L972
	}
L972:
	;
	F_list_free_deep(m, v3432)
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	if v3436 == int64(0) {
		goto L935
	} else {
		goto L974
	}
L974:
	;
	v3442 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	if base.Ui64(v3442) <= base.Ui64(v3436) {
		v3531 = v3436
		goto L934
	} else {
		goto L975
	}
L975:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L976
	}
L976:
	;
	v3448 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+56)) = v3449
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+52)) = uint32(v3448)
	v3453 = int64(base.Ui64(v3448) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+48)) = uint32(v3453)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_116), v3337+int32(48))
	mBase = m.M
	v3459 = m.ExcPending
	if v3459 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+32)) = v3460
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+40)) = uint32(v3436)
	v3464 = int64(base.Ui64(v3436) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+36)) = uint32(v3464)
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_117), v3337+int32(32))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(914), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L980:
	;
	F_errcode(m, int32(_a_F_PostgresMainLoopOnce_119))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_120), int32(0))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		goto L1
	} else {
		goto L982
	}
L982:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_121), int32(0))
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(826), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L1
	} else {
		goto L984
	}
L984:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L985:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L1
	} else {
		goto L986
	}
L986:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_122), int32(0))
	mBase = m.M
	v3521 = m.ExcPending
	if v3521 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(843), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3526 = m.ExcPending
	if v3526 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L989:
	;
	v3544 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	if base.Ui64(v3531) <= base.Ui64(v3544) {
		goto L932
	} else {
		goto L990
	}
L990:
	;
	goto L933
L991:
	;
	v3555 = base.AtomicRmwXchg32(m, v3549, int32(76), int32(1))
	if v3555 != 0 {
		goto L994
	} else {
		goto L995
	}
L992:
	;
	goto L993
L993:
	;
	v3569 = v3337 + int32(128)
	F_pq_beginmessage(m, v3569, int32(87))
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L1
	} else {
		goto L998
	}
L994:
	;
	F_s_lock(m, v3549+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3562 = m.ExcPending
	if v3562 != 0 {
		goto L1
	} else {
		goto L997
	}
L995:
	;
	goto L996
L996:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3549)+4)) = int32(2)
	v3565 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3549)+76)), uint32(v3565))
	goto L993
L997:
	;
	goto L996
L998:
	;
	F_enlargeStringInfo(m, v3569, int32(1))
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		goto L1
	} else {
		goto L999
	}
L999:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+132))
	v3577 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+128))
	v3579 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3576+v3577))) = uint8(v3579)
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+132)) = v3576 + int32(1)
	F_enlargeStringInfo(m, v3569, int32(2))
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L1
	} else {
		goto L1000
	}
L1000:
	;
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+132))
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(v3337)+128))
	v3590 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3587+v3588))) = uint16(v3590)
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+132)) = v3587 + int32(2)
	F_pq_endmessage(m, v3569)
	mBase = m.M
	v3596 = m.ExcPending
	if v3596 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	v3598 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106]))
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(v3598)+4))
	v3600 = m.T0[v3599].(func(*base.Module) int32)(m)
	mBase = m.M
	v3601 = m.ExcPending
	if v3601 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	v3602 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	if base.Ui64(v3420) < base.Ui64(v3602) {
		goto L931
	} else {
		goto L1003
	}
L1003:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107])) = v3602
	v3607 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3610 = base.AtomicRmwXchg32(m, v3607, int32(76), int32(1))
	if v3610 != 0 {
		goto L1004
	} else {
		goto L1005
	}
L1004:
	;
	v3612 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	F_s_lock(m, v3612+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(965), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3619 = m.ExcPending
	if v3619 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1005:
	;
	goto L1006
L1006:
	;
	v3621 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3623 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107]))
	*(*int64)(unsafe.Add(mBase, uint32(v3621)+8)) = v3623
	v3625 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3621)+76)), uint32(v3625))
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1007:
	;
	goto L1006
L1008:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108])) = int32(1)
	F_WalSndLoop(m, int32(1037))
	mBase = m.M
	v3635 = m.ExcPending
	if v3635 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108])) = int32(0)
	v3640 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48]))
	if v3640 != 0 {
		goto L930
	} else {
		goto L1010
	}
L1010:
	;
	v3642 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3643 = *(*int32)(unsafe.Add(mBase, uint32(v3642)+4))
	if v3643 == int32(0) {
		goto L932
	} else {
		goto L1011
	}
L1011:
	;
	v3648 = base.AtomicRmwXchg32(m, v3642, int32(76), int32(1))
	if v3648 != 0 {
		goto L1012
	} else {
		goto L1013
	}
L1012:
	;
	F_s_lock(m, v3642+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3655 = m.ExcPending
	if v3655 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1013:
	;
	goto L1014
L1014:
	;
	v3656 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+4)) = v3656
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3642)+76)), uint32(v3656))
	goto L932
L1015:
	;
	goto L1014
L1016:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3667 = m.ExcPending
	if v3667 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1017:
	;
	goto L1018
L1018:
	;
	v3669 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104])))
	if v3669 == int32(1) {
		goto L1020
	} else {
		goto L1021
	}
L1019:
	;
	goto L1018
L1020:
	;
	v3673 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101]))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+20)) = uint32(v3673)
	v3675 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3337)+70)) = uint16(v3675)
	v3678 = int64(base.Ui64(v3673) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+16)) = uint32(v3678)
	v3681 = v3337 + int32(80)
	v3686 = F_pg_snprintf(m, v3681, int32(18), int32(_a_F_PostgresMainLoopOnce_71), v3337+int32(16))
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1021:
	;
	goto L1022
L1022:
	;
	F_EndReplicationCommand(m, int32(_a_F_PostgresMainLoopOnce_123))
	mBase = m.M
	v3729 = m.ExcPending
	if v3729 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1023:
	;
	v3689 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3690 = m.ExcPending
	if v3690 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	v3692 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	F_TupleDescInitBuiltinEntry(m, v3692, int32(1), int32(_a_F_PostgresMainLoopOnce_124), int32(20))
	mBase = m.M
	v3698 = m.ExcPending
	if v3698 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	F_TupleDescInitBuiltinEntry(m, v3692, int32(2), int32(_a_F_PostgresMainLoopOnce_125), int32(25))
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L1
	} else {
		goto L1027
	}
L1027:
	;
	v3705 = F_begin_tup_output_tupdesc(m, v3689, v3692, int32(_a_F_PostgresMainLoopOnce_77))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	v3708 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[109])))
	v3709 = F_Int64GetDatum(m, v3708)
	mBase = m.M
	v3710 = m.ExcPending
	if v3710 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+72)) = v3709
	v3712 = F_cstring_to_text(m, v3681)
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3337)+76)) = v3712
	F_do_tup_output(m, v3705, v3337+int32(72), v3337+int32(70))
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	F_end_tup_output(m, v3705)
	mBase = m.M
	v3722 = m.ExcPending
	if v3722 != 0 {
		goto L1
	} else {
		goto L1032
	}
L1032:
	;
	goto L1022
L1033:
	;
	m.G0 = v3337 + int32(144)
	v5096 = v3327
	goto L247
L1034:
	;
	v3737 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+4)) = uint32(v3737)
	v3739 = int64(32)
	v3740 = int64(base.Ui64(v3737) >> (uint(v3739) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337))) = uint32(v3740)
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+12)) = uint32(v3420)
	v3744 = int64(base.Ui64(v3420) >> (uint(v3739) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3337)+8)) = uint32(v3744)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_126), v3337)
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(958), int32(_a_F_PostgresMainLoopOnce_118))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1036:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1037:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1038:
	;
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+8))
	v3760 = int32(1)
	F_ReplicationSlotAcquire(m, v3759, v3760, v3760)
	mBase = m.M
	v3763 = m.ExcPending
	if v3763 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	v3765 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])))
	if v3765 != int32(1) {
		goto L1040
	} else {
		goto L1041
	}
L1040:
	;
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+24))
	v3798 = *(*int64)(unsafe.Add(mBase, uint32(v1956)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[77]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[76]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[75]))) = int32(1032)
	v3812 = F_CreateDecodingContext(m, v3798, v3797, int32(0), v726+int32(_a_F_PostgresMainLoopOnce_72), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v3813 = m.ExcPending
	if v3813 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1041:
	;
	v3770 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])))
	if v3770 == int32(1) {
		goto L1043
	} else {
		goto L1044
	}
L1042:
	;
	if v3780 != 0 {
		goto L1040
	} else {
		goto L1046
	}
L1043:
	;
	v3775 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[72]))
	v3776 = *(*int32)(unsafe.Add(mBase, uint32(v3775)+316))
	v3778 = base.B2i32(v3776 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = uint8(v3778)
	v3780 = v3778
	goto L1045
L1044:
	;
	v3780 = int32(0)
	goto L1045
L1045:
	;
	goto L1042
L1046:
	;
	v3783 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3784 = m.ExcPending
	if v3784 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1047:
	;
	if v3783 != 0 {
		goto L1048
	} else {
		goto L1049
	}
L1048:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_127), int32(0))
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1049:
	;
	goto L1050
L1050:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48])) = int32(1)
	goto L1040
L1051:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1467), int32(_a_F_PostgresMainLoopOnce_128))
	mBase = m.M
	v3793 = m.ExcPending
	if v3793 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	goto L1050
L1053:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[110])) = v3812
	v3816 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[100])) = v3816
	v3819 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3820 = *(*int32)(unsafe.Add(mBase, uint32(v3819)+4))
	if v3820 != int32(2) {
		goto L1054
	} else {
		goto L1055
	}
L1054:
	;
	v3825 = base.AtomicRmwXchg32(m, v3819, int32(76), int32(1))
	if v3825 != 0 {
		goto L1057
	} else {
		goto L1058
	}
L1055:
	;
	goto L1056
L1056:
	;
	v3839 = v726 + int32(464)
	F_pq_beginmessage(m, v3839, int32(87))
	mBase = m.M
	v3842 = m.ExcPending
	if v3842 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1057:
	;
	F_s_lock(m, v3819+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1058:
	;
	goto L1059
L1059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3819)+4)) = int32(2)
	v3835 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3819)+76)), uint32(v3835))
	goto L1056
L1060:
	;
	goto L1059
L1061:
	;
	F_enlargeStringInfo(m, v3839, int32(1))
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L1
	} else {
		goto L1062
	}
L1062:
	;
	v3846 = *(*int32)(unsafe.Add(mBase, uint32(v726)+468))
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	v3849 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3846+v3847))) = uint8(v3849)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+468)) = v3846 + int32(1)
	F_enlargeStringInfo(m, v3839, int32(2))
	mBase = m.M
	v3856 = m.ExcPending
	if v3856 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1063:
	;
	v3857 = *(*int32)(unsafe.Add(mBase, uint32(v726)+468))
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	v3860 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3857+v3858))) = uint16(v3860)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+468)) = v3857 + int32(2)
	F_pq_endmessage(m, v3839)
	mBase = m.M
	v3866 = m.ExcPending
	if v3866 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1064:
	;
	v3868 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106]))
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(v3868)+4))
	v3870 = m.T0[v3869].(func(*base.Module) int32)(m)
	mBase = m.M
	v3871 = m.ExcPending
	if v3871 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	v3873 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[110]))
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3873)+8))
	v3876 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3877 = *(*int64)(unsafe.Add(mBase, uint32(v3876)+104))
	F_XLogBeginRead(m, v3874, v3877)
	mBase = m.M
	v3879 = m.ExcPending
	if v3879 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1066:
	;
	v3882 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3883 = *(*int64)(unsafe.Add(mBase, uint32(v3882)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107])) = v3883
	v3886 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3889 = base.AtomicRmwXchg32(m, v3886, int32(76), int32(1))
	if v3889 != 0 {
		goto L1067
	} else {
		goto L1068
	}
L1067:
	;
	v3891 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	F_s_lock(m, v3891+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(1507), int32(_a_F_PostgresMainLoopOnce_128))
	mBase = m.M
	v3898 = m.ExcPending
	if v3898 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1068:
	;
	goto L1069
L1069:
	;
	v3900 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3902 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89]))
	v3903 = *(*int64)(unsafe.Add(mBase, uint32(v3902)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v3900)+8)) = v3903
	v3905 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3900)+76)), uint32(v3905))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108])) = int32(1)
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1070:
	;
	goto L1069
L1071:
	;
	F_WalSndLoop(m, int32(1036))
	mBase = m.M
	v3915 = m.ExcPending
	if v3915 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	v3917 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[110]))
	F_FreeDecodingContext(m, v3917)
	mBase = m.M
	v3919 = m.ExcPending
	if v3919 != 0 {
		goto L1
	} else {
		goto L1073
	}
L1073:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108])) = int32(0)
	v3926 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48]))
	if v3926 != 0 {
		goto L257
	} else {
		goto L1075
	}
L1075:
	;
	v3928 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47]))
	v3929 = *(*int32)(unsafe.Add(mBase, uint32(v3928)+4))
	if v3929 != 0 {
		goto L1076
	} else {
		goto L1077
	}
L1076:
	;
	v3932 = base.AtomicRmwXchg32(m, v3928, int32(76), int32(1))
	if v3932 != 0 {
		goto L1079
	} else {
		goto L1080
	}
L1077:
	;
	goto L1078
L1078:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[97]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[94]))) = int32(56)
	F_EndCommand(m, v726+int32(_a_F_PostgresMainLoopOnce_69), int32(2))
	mBase = m.M
	v3953 = m.ExcPending
	if v3953 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1079:
	;
	F_s_lock(m, v3928+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3939 = m.ExcPending
	if v3939 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1080:
	;
	goto L1081
L1081:
	;
	v3940 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3928)+4)) = v3940
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3928)+76)), uint32(v3940))
	goto L1078
L1082:
	;
	goto L1081
L1083:
	;
	v5096 = v3327
	goto L247
L1084:
	;
	v3958 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v3960 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1085:
	;
	v3963 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	F_TupleDescInitBuiltinEntry(m, v3963, int32(1), int32(_a_F_PostgresMainLoopOnce_129), int32(25))
	mBase = m.M
	v3969 = m.ExcPending
	if v3969 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	F_TupleDescInitBuiltinEntry(m, v3963, int32(2), int32(_a_F_PostgresMainLoopOnce_130), int32(25))
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	v3975 = *(*int32)(unsafe.Add(mBase, uint32(v3958)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+352)) = v3975
	v3978 = v726 + int32(_a_F_PostgresMainLoopOnce_72)
	v3983 = F_pg_snprintf(m, v3978, int32(64), int32(_a_F_PostgresMainLoopOnce_131), v726+int32(352))
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	v3985 = *(*int32)(unsafe.Add(mBase, uint32(v3958)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+336)) = v3985
	v3988 = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	v3993 = F_pg_snprintf(m, v3988, int32(1024), int32(_a_F_PostgresMainLoopOnce_132), v726+int32(336))
	mBase = m.M
	v3994 = m.ExcPending
	if v3994 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	v3996 = *(*int32)(unsafe.Add(mBase, uint32(v3960)+4))
	m.T0[v3996].(func(*base.Module, int32, int32, int32))(m, v3960, int32(1), v3963)
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	v4000 = v726 + int32(_a_F_PostgresMainLoopOnce_78)
	F_pq_beginmessage(m, v4000, int32(68))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	F_enlargeStringInfo(m, v4000, int32(2))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	v4007 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81])))
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67])))
	v4010 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v4007+v4008))) = uint16(v4010)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81]))) = v4007 + int32(2)
	v4015 = F_strlen(m, v3978)
	mBase = m.M
	F_enlargeStringInfo(m, v4000, int32(4))
	mBase = m.M
	v4018 = m.ExcPending
	if v4018 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	v4019 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81])))
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67])))
	v4024 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v4019+v4020))) = base.I32_rotr(v4015, int32(24))&v4024 | base.I32_rotr(v4015&v4024, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81]))) = v4019 + int32(4)
	F_appendBinaryStringInfo(m, v4000, v3978, v4015)
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	v4038 = F_OpenTransientFile(m, v3988, int32(0))
	mBase = m.M
	v4039 = m.ExcPending
	if v4039 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	if v4038 < int32(0) {
		goto L256
	} else {
		goto L1097
	}
L1097:
	;
	v4042 = int64(0)
	v4044 = F___lseek(m, v4038, v4042, int32(2))
	mBase = m.M
	if v4044 < v4042 {
		goto L255
	} else {
		goto L1098
	}
L1098:
	;
	v4047 = int64(0)
	v4049 = F___lseek(m, v4038, v4047, int32(0))
	mBase = m.M
	if v4049 != v4047 {
		goto L254
	} else {
		goto L1099
	}
L1099:
	;
	F_enlargeStringInfo(m, v4000, int32(4))
	mBase = m.M
	v4054 = m.ExcPending
	if v4054 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1100:
	;
	v4055 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81])))
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[67])))
	v4058 = base.I32_wrap_i64(v4044)
	v4059 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v4055+v4056))) = base.I32_rotr(v4058&v4059, int32(8)) | base.I32_rotr(v4058, int32(24))&v4059
	*(*int32)(unsafe.Add(mBase, uint32(v726)+uint32(_c_F_PostgresMainLoopOnce[81]))) = v4055 + int32(4)
	if v4044 != int64(0) {
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	v4106 = v4044
	goto L1104
L1102:
	;
	goto L1103
L1103:
	;
	v4172 = F_CloseTransientFile(m, v4038)
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1104:
	;
	v4111 = int32(_a_F_PostgresMainLoopOnce_133)
	v4112 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v4112))) = int32(167772227)
	v4116 = v726 + int32(464)
	v4118 = F_read(m, v4038, v4116, int32(_a_F_PostgresMainLoopOnce_20))
	mBase = m.M
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[111]))
	v4121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4120))) = v4121
	if v4118 < v4121 {
		goto L253
	} else {
		goto L1106
	}
L1105:
	;
	goto L1103
L1106:
	;
	if v4118 == int32(0) {
		goto L252
	} else {
		goto L1107
	}
L1107:
	;
	F_appendBinaryStringInfo(m, v726+int32(_a_F_PostgresMainLoopOnce_78), v4116, v4118)
	mBase = m.M
	v4130 = m.ExcPending
	if v4130 != 0 {
		goto L1
	} else {
		goto L1108
	}
L1108:
	;
	v4132 = v4106 - base.I64_extend_i32_u(v4118)
	if int64(0) < v4132 {
		v4106 = v4132
		goto L1104
	} else {
		goto L1109
	}
L1109:
	;
	goto L1105
L1110:
	;
	if v4172 != 0 {
		goto L251
	} else {
		goto L1111
	}
L1111:
	;
	F_pq_endmessage(m, v726+int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1112:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_67)
	goto L247
L1113:
	;
	v4185 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[112]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53])) = v4185
	v4188 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4193 = F_AllocSetContextCreateInternal(m, v4188, int32(_a_F_PostgresMainLoopOnce_134), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1114:
	;
	v4195 = int32(_a_F_PostgresMainLoopOnce_135)
	v4196 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4193
	v4200 = F_palloc0(m, int32(36))
	mBase = m.M
	v4201 = m.ExcPending
	if v4201 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4200))) = v4193
	F_initStringInfo(m, v4200+int32(4))
	mBase = m.M
	v4206 = m.ExcPending
	if v4206 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1116:
	;
	v4208 = F_MemoryContextAllocZero(m, v4193, int32(32))
	mBase = m.M
	v4209 = m.ExcPending
	if v4209 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4208)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4208)+24)) = v4193
	v4215 = F_MemoryContextAllocExtended(m, v4193, int32(_a_F_PostgresMainLoopOnce_136), int32(5))
	mBase = m.M
	v4216 = m.ExcPending
	if v4216 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4208)+12)) = int64(63329292795903)
	*(*int64)(unsafe.Add(mBase, uint32(v4208))) = int64(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v4208)+20)) = v4215
	*(*int32)(unsafe.Add(mBase, uint32(v4200)+24)) = v4208
	v4224 = F_palloc0(m, int32(24))
	mBase = m.M
	v4225 = m.ExcPending
	if v4225 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4224)+20)) = int32(427)
	*(*int32)(unsafe.Add(mBase, uint32(v4224)+16)) = int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v4224)+12)) = int32(429)
	*(*int32)(unsafe.Add(mBase, uint32(v4224)+8)) = int32(430)
	*(*int32)(unsafe.Add(mBase, uint32(v4224)+4)) = int32(431)
	*(*int32)(unsafe.Add(mBase, uint32(v4224))) = v4200
	v4238 = F_palloc(m, int32(112))
	mBase = m.M
	v4239 = m.ExcPending
	if v4239 != 0 {
		goto L1
	} else {
		goto L1120
	}
L1120:
	;
	v4241 = F_palloc(m, int32(68))
	mBase = m.M
	v4242 = m.ExcPending
	if v4242 != 0 {
		goto L1
	} else {
		goto L1121
	}
L1121:
	;
	v4243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4241)+52)) = uint8(v4243)
	*(*int32)(unsafe.Add(mBase, uint32(v4241)+4)) = v4243
	*(*int32)(unsafe.Add(mBase, uint32(v4241))) = v4224
	if v4238 == v4243 {
		goto L1124
	} else {
		goto L1125
	}
L1122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+104)) = int32(_a_F_PostgresMainLoopOnce_137)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+100)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4238)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+88)) = int32(_a_F_PostgresMainLoopOnce_138)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+84)) = int32(_a_F_PostgresMainLoopOnce_139)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+80)) = int32(_a_F_PostgresMainLoopOnce_140)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+76)) = int32(_a_F_PostgresMainLoopOnce_141)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+72)) = int32(_a_F_PostgresMainLoopOnce_142)
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+68)) = v4241
	v4335 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v4336 = m.ExcPending
	if v4336 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+40)) = int32(1)
	v4268 = F_palloc0(m, int32(20))
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1124:
	;
	v4251 = F_palloc0(m, int32(68))
	mBase = m.M
	v4252 = m.ExcPending
	if v4252 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1125:
	;
	goto L1126
L1126:
	;
	base.MemoryFill(m, v4238, int32(0), int32(68))
	v4262 = v4238
	goto L1123
L1127:
	;
	if v4251 == int32(0) {
		goto L1122
	} else {
		goto L1128
	}
L1128:
	;
	v4255 = *(*int32)(unsafe.Add(mBase, uint32(v4251)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4251)+36)) = v4255 | int32(1)
	v4262 = v4251
	goto L1123
L1129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+52)) = v4268
	v4272 = F_palloc0(m, int32(28))
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1130:
	;
	v4275 = F_palloc(m, int32(640))
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1131:
	;
	v4278 = F_palloc(m, int32(256))
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1132:
	;
	v4281 = F_palloc(m, int32(64))
	mBase = m.M
	v4282 = m.ExcPending
	if v4282 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+52))
	F_initStringInfo(m, v4283+int32(4))
	mBase = m.M
	v4287 = m.ExcPending
	if v4287 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+48)) = v4272
	*(*int32)(unsafe.Add(mBase, uint32(v4272))) = int32(64)
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4291)+4)) = v4275
	v4293 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4293)+12)) = v4278
	v4295 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4295)+16)) = v4281
	v4297 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+48))
	v4298 = *(*int32)(unsafe.Add(mBase, uint32(v4297)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v4298))) = int32(0)
	v4301 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4262)+56)) = uint8(v4301)
	*(*uint8)(unsafe.Add(mBase, uint32(v4262)+24)) = uint8(v4301)
	v4305 = F_makeStringInfo(m)
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+60)) = v4305
	v4308 = *(*int32)(unsafe.Add(mBase, uint32(v4262)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4262)+36)) = v4308 | int32(2)
	goto L1122
L1136:
	;
	if v4335 == int32(0) {
		goto L1137
	} else {
		goto L1138
	}
L1137:
	;
	v4341 = *(*int32)(unsafe.Add(mBase, uint32(v4224)+20))
	m.T0[v4341].(func(*base.Module, int32, int32, int32))(m, v4224, int32(_a_F_PostgresMainLoopOnce_120), int32(0))
	mBase = m.M
	v4343 = m.ExcPending
	if v4343 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1138:
	;
	goto L1139
L1139:
	;
	v4344 = F_pg_cryptohash_init(m, v4335)
	mBase = m.M
	if v4344 < int32(0) {
		goto L1141
	} else {
		goto L1142
	}
L1140:
	;
	goto L1139
L1141:
	;
	v4349 = *(*int32)(unsafe.Add(mBase, uint32(v4224)+20))
	m.T0[v4349].(func(*base.Module, int32, int32, int32))(m, v4224, int32(_a_F_PostgresMainLoopOnce_143), int32(0))
	mBase = m.M
	v4351 = m.ExcPending
	if v4351 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1142:
	;
	goto L1143
L1143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4238)+108)) = v4335
	*(*int32)(unsafe.Add(mBase, uint32(v4200)+32)) = v4238
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4196
	v4357 = v726 + int32(464)
	F_pq_beginmessage(m, v4357, int32(71))
	mBase = m.M
	v4360 = m.ExcPending
	if v4360 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1144:
	;
	goto L1143
L1145:
	;
	F_enlargeStringInfo(m, v4357, int32(1))
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L1
	} else {
		goto L1146
	}
L1146:
	;
	v4364 = *(*int32)(unsafe.Add(mBase, uint32(v726)+468))
	v4365 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	v4367 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4364+v4365))) = uint8(v4367)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+468)) = v4364 + int32(1)
	F_enlargeStringInfo(m, v4357, int32(2))
	mBase = m.M
	v4374 = m.ExcPending
	if v4374 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	v4375 = *(*int32)(unsafe.Add(mBase, uint32(v726)+468))
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	v4378 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4375+v4376))) = uint16(v4378)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+468)) = v4375 + int32(2)
	F_pq_endmessage_reuse(m, v4357)
	mBase = m.M
	v4384 = m.ExcPending
	if v4384 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	v4386 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106]))
	v4387 = *(*int32)(unsafe.Add(mBase, uint32(v4386)+4))
	v4388 = m.T0[v4387].(func(*base.Module) int32)(m)
	mBase = m.M
	v4389 = m.ExcPending
	if v4389 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	goto L1151
L1150:
	;
	v4523 = int32(_a_F_PostgresMainLoopOnce_135)
	v4524 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4526 = *(*int32)(unsafe.Add(mBase, uint32(v4200)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4526
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+32))
	v4529 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+4))
	v4530 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+8))
	F_json_parse_manifest_incremental_chunk(m, v4528, v4529, v4530, int32(1))
	mBase = m.M
	v4533 = m.ExcPending
	if v4533 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1151:
	;
	v4427 = int32(_a_F_PostgresMainLoopOnce_5)
	v4429 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4429 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1153:
	;
	v4435 = F_pq_getbyte(m)
	mBase = m.M
	v4436 = m.ExcPending
	if v4436 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1154:
	;
	v4438 = v4435 - int32(72)
	if base.Ui32(int32(30)) < base.Ui32(v4438) {
		goto L249
	} else {
		goto L1155
	}
L1155:
	;
	if int32(1)<<(uint(v4438)%32)&int32(1207961601) == int32(0) {
		goto L1157
	} else {
		goto L1158
	}
L1156:
	;
	v4454 = F_pq_getmessage(m, v726+int32(464), v4451)
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1157:
	;
	if v4438 != int32(28) {
		goto L249
	} else {
		goto L1160
	}
L1158:
	;
	goto L1159
L1159:
	;
	v4451 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L1156
L1160:
	;
	v4451 = int32(1073741822)
	goto L1156
L1161:
	;
	if v4454 != 0 {
		goto L250
	} else {
		goto L1162
	}
L1162:
	;
	v4456 = int32(_a_F_PostgresMainLoopOnce_5)
	v4458 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4458 - int32(1)
	switch v4438 {
	case 0, 11:
		goto L1151
	default:
		goto L1150
	case 28:
		goto L1164
	case 30:
		goto L1163
	}
L1163:
	;
	goto L1152
L1164:
	;
	v4462 = *(*int32)(unsafe.Add(mBase, uint32(v726)+464))
	v4463 = *(*int32)(unsafe.Add(mBase, uint32(v726)+468))
	v4464 = int32(_a_F_PostgresMainLoopOnce_135)
	v4465 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4467 = *(*int32)(unsafe.Add(mBase, uint32(v4200)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4467
	v4469 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+8))
	if base.B2i32(v4469 < int32(1025))|base.B2i32(v4469+v4463 < int32(_a_F_PostgresMainLoopOnce_144)) == int32(0) {
		goto L1165
	} else {
		goto L1166
	}
L1165:
	;
	v4478 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+32))
	v4479 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+4))
	F_json_parse_manifest_incremental_chunk(m, v4478, v4479, v4469-int32(1024), int32(0))
	mBase = m.M
	v4484 = m.ExcPending
	if v4484 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1166:
	;
	goto L1167
L1167:
	;
	F_appendBinaryStringInfo(m, v4200+int32(4), v4462, v4463)
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1168:
	;
	v4485 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+4))
	v4486 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+8))
	v4488 = int32(1024)
	base.MemoryCopy(m, v4485, v4485+v4486-v4488, int32(1025))
	*(*int32)(unsafe.Add(mBase, uint32(v4200)+8)) = v4488
	goto L1167
L1169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4465
	goto L1151
L1170:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v4507 = m.ExcPending
	if v4507 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	v4510 = F_pq_getmsgstring(m, v726+int32(464))
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L1
	} else {
		goto L1172
	}
L1172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+384)) = v4510
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_145), v726+int32(384))
	mBase = m.M
	v4517 = m.ExcPending
	if v4517 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1173:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(794), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4522 = m.ExcPending
	if v4522 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1175:
	;
	v4534 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+4))
	F_pfree(m, v4534)
	mBase = m.M
	v4536 = m.ExcPending
	if v4536 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4200)+4)) = int32(0)
	v4539 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+32))
	v4540 = *(*int32)(unsafe.Add(mBase, uint32(v4539)+68))
	F_pfree(m, v4540)
	mBase = m.M
	v4542 = m.ExcPending
	if v4542 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	F_freeJsonLexContext(m, v4539)
	mBase = m.M
	v4544 = m.ExcPending
	if v4544 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1178:
	;
	F_pfree(m, v4539)
	mBase = m.M
	v4546 = m.ExcPending
	if v4546 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4524
	v4550 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113]))
	if v4550 != 0 {
		goto L1180
	} else {
		goto L1181
	}
L1180:
	;
	F_MemoryContextDelete(m, v4550)
	mBase = m.M
	v4552 = m.ExcPending
	if v4552 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1181:
	;
	goto L1182
L1182:
	;
	v4554 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114]))
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v4193)+16))
	if v4558 != v4554 {
		goto L1185
	} else {
		goto L1186
	}
L1183:
	;
	goto L1182
L1184:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = v4193
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v4200
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v4593 = m.ExcPending
	if v4593 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1185:
	;
	if v4558 == int32(0) {
		goto L1188
	} else {
		goto L1189
	}
L1186:
	;
	goto L1187
L1187:
	;
	goto L1184
L1188:
	;
	if v4554 != 0 {
		goto L1195
	} else {
		goto L1196
	}
L1189:
	;
	v4562 = *(*int32)(unsafe.Add(mBase, uint32(v4193)+28))
	v4563 = *(*int32)(unsafe.Add(mBase, uint32(v4193)+24))
	if v4563 != 0 {
		goto L1191
	} else {
		goto L1192
	}
L1190:
	;
	if v4562 == int32(0) {
		goto L1188
	} else {
		goto L1194
	}
L1191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4563)+28)) = v4562
	goto L1190
L1192:
	;
	goto L1193
L1193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4558)+20)) = v4562
	goto L1190
L1194:
	;
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(v4193)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4562)+24)) = v4568
	goto L1188
L1195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4193)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4193)+16)) = v4554
	v4575 = *(*int32)(unsafe.Add(mBase, uint32(v4554)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4193)+28)) = v4575
	if v4575 != 0 {
		goto L1198
	} else {
		goto L1199
	}
L1196:
	;
	goto L1197
L1197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4193)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4193)+16)) = int32(0)
	goto L1187
L1198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4575)+24)) = v4193
	goto L1200
L1199:
	;
	goto L1200
L1200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4554)+20)) = v4193
	goto L1184
L1201:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_66)
	goto L247
L1202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	v4601 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v4602 = *(*int32)(unsafe.Add(mBase, uint32(v4601)))
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = v4602
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_147), v726)
	mBase = m.M
	v4606 = m.ExcPending
	if v4606 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2215), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4611 = m.ExcPending
	if v4611 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1206:
	;
	v4617 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	v4619 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115]))
	F_SendBaseBackup(m, v4617, v4619)
	mBase = m.M
	v4621 = m.ExcPending
	if v4621 != 0 {
		goto L1
	} else {
		goto L1207
	}
L1207:
	;
	v5096 = v4612
	goto L247
L1208:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1209:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_148), int32(0))
	mBase = m.M
	v4632 = m.ExcPending
	if v4632 != 0 {
		goto L1
	} else {
		goto L1210
	}
L1210:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2010), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L1
	} else {
		goto L1211
	}
L1211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+416)) = v1861
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_149), v726+int32(416))
	mBase = m.M
	v4650 = m.ExcPending
	if v4650 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1214:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2078), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4655 = m.ExcPending
	if v4655 != 0 {
		goto L1
	} else {
		goto L1215
	}
L1215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1216:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v4662 = m.ExcPending
	if v4662 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1217:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v4666 = m.ExcPending
	if v4666 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2104), int32(_a_F_PostgresMainLoopOnce_29))
	mBase = m.M
	v4671 = m.ExcPending
	if v4671 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1220:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4678 = m.ExcPending
	if v4678 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+64)) = int32(_a_F_PostgresMainLoopOnce_85)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_151), v726-int32(-64))
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(520), int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v4690 = m.ExcPending
	if v4690 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1224:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4701 = m.ExcPending
	if v4701 != 0 {
		goto L1
	} else {
		goto L1226
	}
L1226:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1137), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4706 = m.ExcPending
	if v4706 != 0 {
		goto L1
	} else {
		goto L1227
	}
L1227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1228:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L1
	} else {
		goto L1229
	}
L1229:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		goto L1
	} else {
		goto L1230
	}
L1230:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1159), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		goto L1
	} else {
		goto L1231
	}
L1231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1232:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1233:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4733 = m.ExcPending
	if v4733 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1234:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1169), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		goto L1
	} else {
		goto L1238
	}
L1238:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1178), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1240:
	;
	v4759 = *(*int32)(unsafe.Add(mBase, uint32(v2327)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+208)) = v4759
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_153), v726+int32(208))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1183), int32(_a_F_PostgresMainLoopOnce_88))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+176)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_155), v726+int32(176))
	mBase = m.M
	v4781 = m.ExcPending
	if v4781 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1268), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+160)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_156), v726+int32(160))
	mBase = m.M
	v4797 = m.ExcPending
	if v4797 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1274), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4802 = m.ExcPending
	if v4802 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v726)+144)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_157), v726+int32(144))
	mBase = m.M
	v4813 = m.ExcPending
	if v4813 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1279), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+112)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_158), v726+int32(112))
	mBase = m.M
	v4829 = m.ExcPending
	if v4829 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1253:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1285), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+128)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_159), v726+int32(128))
	mBase = m.M
	v4845 = m.ExcPending
	if v4845 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1291), int32(_a_F_PostgresMainLoopOnce_92))
	mBase = m.M
	v4850 = m.ExcPending
	if v4850 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1258:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4857 = m.ExcPending
	if v4857 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1259:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4861 = m.ExcPending
	if v4861 != 0 {
		goto L1
	} else {
		goto L1260
	}
L1260:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1420), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4866 = m.ExcPending
	if v4866 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1262:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4873 = m.ExcPending
	if v4873 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), int32(0))
	mBase = m.M
	v4877 = m.ExcPending
	if v4877 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1429), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4882 = m.ExcPending
	if v4882 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1266:
	;
	v4887 = *(*int32)(unsafe.Add(mBase, uint32(v2986)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+224)) = v4887
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_153), v726+int32(224))
	mBase = m.M
	v4893 = m.ExcPending
	if v4893 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1434), int32(_a_F_PostgresMainLoopOnce_160))
	mBase = m.M
	v4898 = m.ExcPending
	if v4898 != 0 {
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
	base.Wasm_trap_unreachable()
	for {
	}
L1270:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4907 = m.ExcPending
	if v4907 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+240)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_161), v726+int32(240))
	mBase = m.M
	v4915 = m.ExcPending
	if v4915 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(616), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4920 = m.ExcPending
	if v4920 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1274:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4926 = m.ExcPending
	if v4926 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+256)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_163), v726+int32(256))
	mBase = m.M
	v4934 = m.ExcPending
	if v4934 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1276:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(623), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4939 = m.ExcPending
	if v4939 != 0 {
		goto L1
	} else {
		goto L1277
	}
L1277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1278:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4945 = m.ExcPending
	if v4945 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+320)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_164), v726+int32(320))
	mBase = m.M
	v4953 = m.ExcPending
	if v4953 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(627), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4958 = m.ExcPending
	if v4958 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1282:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4964 = m.ExcPending
	if v4964 != 0 {
		goto L1
	} else {
		goto L1283
	}
L1283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+288)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_165), v726+int32(288))
	mBase = m.M
	v4972 = m.ExcPending
	if v4972 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1284:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(644), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v4977 = m.ExcPending
	if v4977 != 0 {
		goto L1
	} else {
		goto L1285
	}
L1285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1286:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4984 = m.ExcPending
	if v4984 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1287:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v726)+312)) = uint32(v4106)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+308)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+304)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_166), v726+int32(304))
	mBase = m.M
	v4995 = m.ExcPending
	if v4995 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1288:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(649), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v5000 = m.ExcPending
	if v5000 != 0 {
		goto L1
	} else {
		goto L1289
	}
L1289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1290:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v5006 = m.ExcPending
	if v5006 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+272)) = v726 + int32(_a_F_PostgresMainLoopOnce_69)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_167), v726+int32(272))
	mBase = m.M
	v5014 = m.ExcPending
	if v5014 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1292:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(658), int32(_a_F_PostgresMainLoopOnce_162))
	mBase = m.M
	v5019 = m.ExcPending
	if v5019 != 0 {
		goto L1
	} else {
		goto L1293
	}
L1293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1294:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v5026 = m.ExcPending
	if v5026 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1295:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v5030 = m.ExcPending
	if v5030 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(772), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v5035 = m.ExcPending
	if v5035 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1298:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5042 = m.ExcPending
	if v5042 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1299:
	;
	goto L1300
L1300:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5058 = m.ExcPending
	if v5058 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1301:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v5045 = m.ExcPending
	if v5045 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1302:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v5049 = m.ExcPending
	if v5049 != 0 {
		goto L1
	} else {
		goto L1303
	}
L1303:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(746), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v5054 = m.ExcPending
	if v5054 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1305:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v5061 = m.ExcPending
	if v5061 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+368)) = v4435
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_168), v726+int32(368))
	mBase = m.M
	v5067 = m.ExcPending
	if v5067 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1307:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(763), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v5072 = m.ExcPending
	if v5072 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1309:
	;
	v5076 = *(*int32)(unsafe.Add(mBase, uint32(v726)+424))
	F_StartTransactionCommand(m)
	mBase = m.M
	v5079 = m.ExcPending
	if v5079 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1310:
	;
	v5080 = *(*int32)(unsafe.Add(mBase, uint32(v5076)+4))
	F_GetPGVariable(m, v5080, v5074)
	mBase = m.M
	v5082 = m.ExcPending
	if v5082 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1311:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5084 = m.ExcPending
	if v5084 != 0 {
		goto L1
	} else {
		goto L1312
	}
L1312:
	;
	v5096 = int32(_a_F_PostgresMainLoopOnce_169)
	goto L247
L1313:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v731
	v5127 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[54]))
	F_MemoryContextReset(m, v5127)
	mBase = m.M
	v5129 = m.ExcPending
	if v5129 != 0 {
		goto L1
	} else {
		goto L1314
	}
L1314:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = int32(0)
	goto L246
L1315:
	;
	goto L238
L1316:
	;
	v5225 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v5235 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1319
L1317:
	;
	goto L1318
L1318:
	;
	F_start_xact_command(m)
	mBase = m.M
	v5239 = m.ExcPending
	if v5239 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1319:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1318
L1320:
	;
	v5241 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120]))
	if v5241 != 0 {
		goto L1321
	} else {
		goto L1322
	}
L1321:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = int32(0)
	F_DropCachedPlan(m, v5241)
	mBase = m.M
	v5246 = m.ExcPending
	if v5246 != 0 {
		goto L1
	} else {
		goto L1324
	}
L1322:
	;
	goto L1323
L1323:
	;
	v5247 = int32(_a_F_PostgresMainLoopOnce_135)
	v5248 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v5251 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5251
	v5254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v5254 == int32(1) {
		goto L1325
	} else {
		goto L1326
	}
L1324:
	;
	goto L1323
L1325:
	;
	v5257 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v5267 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1328
L1326:
	;
	goto L1327
L1327:
	;
	v5271 = F_raw_parser(m, v716, int32(0))
	mBase = m.M
	v5272 = m.ExcPending
	if v5272 != 0 {
		goto L1
	} else {
		goto L1329
	}
L1328:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1327
L1329:
	;
	v5274 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v5274 == int32(1) {
		goto L1330
	} else {
		goto L1331
	}
L1330:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_172))
	mBase = m.M
	v5279 = m.ExcPending
	if v5279 != 0 {
		goto L1
	} else {
		goto L1333
	}
L1331:
	;
	goto L1332
L1332:
	;
	v5281 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122]))
	switch v5281 {
	case 0:
		v5572 = v1
		goto L1338
	default:
		goto L1343
	case 3:
		goto L1342
	}
L1333:
	;
	goto L1332
L1334:
	;
	v6178 = F_check_log_duration(m, v5213+int32(80), v6167)
	mBase = m.M
	v6179 = m.ExcPending
	if v6179 != 0 {
		goto L1
	} else {
		goto L1516
	}
L1335:
	;
	v6123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L1505
L1336:
	;
	v6069 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L1495
L1337:
	;
	v5622 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if v5622 <= int32(0) {
		goto L1335
	} else {
		goto L1370
	}
L1338:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5248
	if v5271 == int32(0) {
		v6058 = v5572
		goto L1336
	} else {
		goto L1369
	}
L1339:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1179), int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v5542 = m.ExcPending
	if v5542 != 0 {
		goto L1
	} else {
		goto L1368
	}
L1340:
	;
	v5493 = *(*int32)(unsafe.Add(mBase, uint32(v5483)+64))
	v5494 = *(*int32)(unsafe.Add(mBase, uint32(v5493)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5213)+48)) = v5494
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_174), v5213+int32(48))
	mBase = m.M
	v5500 = m.ExcPending
	if v5500 != 0 {
		goto L1
	} else {
		goto L1367
	}
L1341:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5248
	v6058 = v1
	goto L1336
L1342:
	;
	v5414 = int32(1)
	v5417 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5418 = m.ExcPending
	if v5418 != 0 {
		goto L1
	} else {
		goto L1353
	}
L1343:
	;
	if v5271 == int32(0) {
		goto L1341
	} else {
		goto L1344
	}
L1344:
	;
	v5284 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if int32(0) < v5284 {
		goto L1345
	} else {
		goto L1346
	}
L1345:
	;
	v5287 = int32(0)
	goto L1348
L1346:
	;
	v5354 = v5284
	goto L1347
L1347:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5248
	v5601 = v5354
	v5614 = v1
	goto L1337
L1348:
	;
	v5324 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v5328 = *(*int32)(unsafe.Add(mBase, uint32(v5324+v5287<<(uint(int32(2))%32))))
	v5329 = F_GetCommandLogLevel(m, v5328)
	mBase = m.M
	v5330 = m.ExcPending
	if v5330 != 0 {
		goto L1
	} else {
		goto L1350
	}
L1349:
	;
	v5354 = v5336
	goto L1347
L1350:
	;
	v5332 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122]))
	if base.Ui32(v5329) <= base.Ui32(v5332) {
		goto L1342
	} else {
		goto L1351
	}
L1351:
	;
	v5335 = v5287 + int32(1)
	v5336 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if v5335 < v5336 {
		v5287 = v5335
		goto L1348
	} else {
		goto L1352
	}
L1352:
	;
	goto L1349
L1353:
	;
	if v5417 == int32(0) {
		v5572 = v5414
		goto L1338
	} else {
		goto L1354
	}
L1354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5213)+64)) = v716
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_175), v5213-int32(-64))
	mBase = m.M
	v5426 = m.ExcPending
	if v5426 != 0 {
		goto L1
	} else {
		goto L1355
	}
L1355:
	;
	F_errhidestmt(m)
	mBase = m.M
	v5428 = m.ExcPending
	if v5428 != 0 {
		goto L1
	} else {
		goto L1356
	}
L1356:
	;
	if v5271 == int32(0) {
		goto L1339
	} else {
		goto L1357
	}
L1357:
	;
	v5431 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if v5431 <= int32(0) {
		goto L1339
	} else {
		goto L1358
	}
L1358:
	;
	v5435 = int32(0)
	v5450 = v5431
	goto L1359
L1359:
	;
	v5472 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v5476 = *(*int32)(unsafe.Add(mBase, uint32(v5472+v5435<<(uint(int32(2))%32))))
	v5477 = *(*int32)(unsafe.Add(mBase, uint32(v5476)+4))
	v5478 = *(*int32)(unsafe.Add(mBase, uint32(v5477)))
	if v5478 == int32(253) {
		goto L1361
	} else {
		goto L1362
	}
L1360:
	;
	goto L1339
L1361:
	;
	v5481 = *(*int32)(unsafe.Add(mBase, uint32(v5477)+4))
	v5483 = F_FetchPreparedStatement(m, v5481, int32(0))
	mBase = m.M
	v5484 = m.ExcPending
	if v5484 != 0 {
		goto L1
	} else {
		goto L1364
	}
L1362:
	;
	v5487 = v5450
	goto L1363
L1363:
	;
	v5489 = v5435 + int32(1)
	if v5489 < v5487 {
		v5435 = v5489
		v5450 = v5487
		goto L1359
	} else {
		goto L1366
	}
L1364:
	;
	if v5483 != 0 {
		goto L1340
	} else {
		goto L1365
	}
L1365:
	;
	v5485 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	v5487 = v5485
	goto L1363
L1366:
	;
	goto L1360
L1367:
	;
	goto L1339
L1368:
	;
	v5572 = v5414
	goto L1338
L1369:
	;
	v5584 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	v5601 = v5584
	v5614 = v5572
	goto L1337
L1370:
	;
	v5647 = v1
	goto L1371
L1371:
	;
	v5662 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v5665 = v5662 + v5647<<(uint(int32(2))%32)
	v5666 = *(*int32)(unsafe.Add(mBase, uint32(v5665)))
	v5671 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v5671 == int32(0) {
		goto L1374
	} else {
		goto L1375
	}
L1372:
	;
	goto L1335
L1373:
	;
	v5711 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v5711 == int32(0) {
		goto L1379
	} else {
		goto L1380
	}
L1374:
	;
	goto L1373
L1375:
	;
	v5675 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v5675&int32(1) == int32(0) {
		goto L1374
	} else {
		goto L1376
	}
L1376:
	;
	v5682 = *(*int64)(unsafe.Add(mBase, uint32(v5671)+392))
	if int32(0)&base.B2i32(v5682 != int64(0)) != 0 {
		goto L1374
	} else {
		goto L1377
	}
L1377:
	;
	v5686 = int32(_a_F_PostgresMainLoopOnce_176)
	v5688 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v5689 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v5688 + v5689
	v5692 = *(*int32)(unsafe.Add(mBase, uint32(v5671)))
	*(*int32)(unsafe.Add(mBase, uint32(v5671))) = v5692 + v5689
	*(*int64)(unsafe.Add(mBase, uint32(v5671)+392)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5671))) = v5692 + int32(2)
	v5703 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v5703 - v5689
	goto L1374
L1378:
	;
	v5747 = *(*int32)(unsafe.Add(mBase, uint32(v5666)+4))
	v5748 = F_CreateCommandTag(m, v5747)
	mBase = m.M
	v5749 = m.ExcPending
	if v5749 != 0 {
		goto L1
	} else {
		goto L1383
	}
L1379:
	;
	goto L1378
L1380:
	;
	v5715 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v5715&int32(1) == int32(0) {
		goto L1379
	} else {
		goto L1381
	}
L1381:
	;
	v5722 = *(*int64)(unsafe.Add(mBase, uint32(v5711)+400))
	if int32(0)&base.B2i32(v5722 != int64(0)) != 0 {
		goto L1379
	} else {
		goto L1382
	}
L1382:
	;
	v5726 = int32(_a_F_PostgresMainLoopOnce_176)
	v5728 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v5729 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v5728 + v5729
	v5732 = *(*int32)(unsafe.Add(mBase, uint32(v5711)))
	*(*int32)(unsafe.Add(mBase, uint32(v5711))) = v5732 + v5729
	*(*int64)(unsafe.Add(mBase, uint32(v5711)+400)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5711))) = v5732 + int32(2)
	v5743 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v5743 - v5729
	goto L1379
L1383:
	;
	v5754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5748<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[127]))))
	*(*int32)(unsafe.Add(mBase, uint32(v5213+int32(72)))) = v5754
	goto L1384
L1384:
	;
	v5759 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5760 = *(*int32)(unsafe.Add(mBase, uint32(v5759)+24))
	goto L1386
L1385:
	;
	F_start_xact_command(m)
	mBase = m.M
	v5802 = m.ExcPending
	if v5802 != 0 {
		goto L1
	} else {
		goto L1397
	}
L1386:
	;
	if base.B2i32((v5760-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L1385
	} else {
		goto L1387
	}
L1387:
	;
	v5769 = *(*int32)(unsafe.Add(mBase, uint32(v5666)+4))
	if v5769 == int32(0) {
		goto L1388
	} else {
		goto L1389
	}
L1388:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5785 = m.ExcPending
	if v5785 != 0 {
		goto L1
	} else {
		goto L1392
	}
L1389:
	;
	v5772 = *(*int32)(unsafe.Add(mBase, uint32(v5769)))
	if v5772 != int32(225) {
		goto L1388
	} else {
		goto L1390
	}
L1390:
	;
	v5775 = *(*int32)(unsafe.Add(mBase, uint32(v5769)+4))
	if (v5775-int32(2))&int32(-6) == int32(0) {
		goto L1385
	} else {
		goto L1391
	}
L1391:
	;
	goto L1388
L1392:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v5788 = m.ExcPending
	if v5788 != 0 {
		goto L1
	} else {
		goto L1393
	}
L1393:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v5792 = m.ExcPending
	if v5792 != 0 {
		goto L1
	} else {
		goto L1394
	}
L1394:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v5794 = m.ExcPending
	if v5794 != 0 {
		goto L1
	} else {
		goto L1395
	}
L1395:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1246), int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v5799 = m.ExcPending
	if v5799 != 0 {
		goto L1
	} else {
		goto L1396
	}
L1396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1397:
	;
	v5804 = base.B2i32(v5601 < int32(2))
	if v5804 == int32(0) {
		goto L1398
	} else {
		goto L1399
	}
L1398:
	;
	v5809 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5810 = *(*int32)(unsafe.Add(mBase, uint32(v5809)+24))
	if v5810 == int32(1) {
		goto L1402
	} else {
		goto L1403
	}
L1399:
	;
	goto L1400
L1400:
	;
	v5816 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v5816 != 0 {
		goto L1405
	} else {
		goto L1406
	}
L1401:
	;
	goto L1400
L1402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5809)+24)) = int32(4)
	goto L1404
L1403:
	;
	goto L1404
L1404:
	;
	goto L1401
L1405:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5818 = m.ExcPending
	if v5818 != 0 {
		goto L1
	} else {
		goto L1408
	}
L1406:
	;
	goto L1407
L1407:
	;
	v5821 = *(*int32)(unsafe.Add(mBase, uint32(v5666)+4))
	v5822 = *(*int32)(unsafe.Add(mBase, uint32(v5821)))
	switch v5822 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v5826 = int32(1)
		goto L1410
	default:
		goto L1411
	}
L1408:
	;
	goto L1407
L1409:
	;
	if v5826 != 0 {
		goto L1412
	} else {
		goto L1413
	}
L1410:
	;
	goto L1409
L1411:
	;
	v5826 = int32(0)
	goto L1410
L1412:
	;
	v5827 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v5828 = m.ExcPending
	if v5828 != 0 {
		goto L1
	} else {
		goto L1415
	}
L1413:
	;
	goto L1414
L1414:
	;
	v5833 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v5835 = v5665 + int32(4)
	v5836 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v5837 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if base.Ui32(v5835) < base.Ui32(v5836+v5837<<(uint(int32(2))%32)) {
		goto L1417
	} else {
		goto L1418
	}
L1415:
	;
	F_PushActiveSnapshot(m, v5827)
	mBase = m.M
	v5830 = m.ExcPending
	if v5830 != 0 {
		goto L1
	} else {
		goto L1416
	}
L1416:
	;
	goto L1414
L1417:
	;
	v5846 = F_AllocSetContextCreateInternal(m, v5833, int32(_a_F_PostgresMainLoopOnce_177), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v5847 = m.ExcPending
	if v5847 != 0 {
		goto L1
	} else {
		goto L1420
	}
L1418:
	;
	v5848 = v5833
	v5849 = int32(0)
	goto L1419
L1419:
	;
	v5850 = int32(_a_F_PostgresMainLoopOnce_135)
	v5851 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5848
	v5855 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v5855 == int32(1) {
		goto L1421
	} else {
		goto L1422
	}
L1420:
	;
	v5848 = v5846
	v5849 = v5846
	goto L1419
L1421:
	;
	v5858 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v5868 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1424
L1422:
	;
	goto L1423
L1423:
	;
	v5871 = int32(0)
	v5874 = F_parse_analyze_fixedparams(m, v5666, v716, v5871, v5871, v5871)
	mBase = m.M
	v5875 = m.ExcPending
	if v5875 != 0 {
		goto L1
	} else {
		goto L1425
	}
L1424:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1423
L1425:
	;
	v5877 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v5877 == int32(1) {
		goto L1426
	} else {
		goto L1427
	}
L1426:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_178))
	mBase = m.M
	v5882 = m.ExcPending
	if v5882 != 0 {
		goto L1
	} else {
		goto L1429
	}
L1427:
	;
	goto L1428
L1428:
	;
	v5883 = F_pg_rewrite_query(m, v5874)
	mBase = m.M
	v5884 = m.ExcPending
	if v5884 != 0 {
		goto L1
	} else {
		goto L1430
	}
L1429:
	;
	goto L1428
L1430:
	;
	v5887 = F_pg_plan_queries(m, v5883, v716, int32(2048), int32(0))
	mBase = m.M
	v5888 = m.ExcPending
	if v5888 != 0 {
		goto L1
	} else {
		goto L1431
	}
L1431:
	;
	if v5826 != 0 {
		goto L1432
	} else {
		goto L1433
	}
L1432:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v5890 = m.ExcPending
	if v5890 != 0 {
		goto L1
	} else {
		goto L1435
	}
L1433:
	;
	goto L1434
L1434:
	;
	v5892 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v5892 != 0 {
		goto L1436
	} else {
		goto L1437
	}
L1435:
	;
	goto L1434
L1436:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5894 = m.ExcPending
	if v5894 != 0 {
		goto L1
	} else {
		goto L1439
	}
L1437:
	;
	goto L1438
L1438:
	;
	v5896 = int32(1)
	v5898 = F_CreatePortal(m, int32(_a_F_PostgresMainLoopOnce_179), v5896, v5896)
	mBase = m.M
	v5899 = m.ExcPending
	if v5899 != 0 {
		goto L1
	} else {
		goto L1440
	}
L1439:
	;
	goto L1438
L1440:
	;
	v5900 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5898)+136)) = uint8(v5900)
	*(*int64)(unsafe.Add(mBase, uint32(v5898)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+40)) = v5748
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+32)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+4)) = v5900
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+60)) = v5900
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+56)) = v5887
	*(*int32)(unsafe.Add(mBase, uint32(v5898)+36)) = v5748
	goto L1441
L1441:
	;
	v5914 = int32(0)
	F_PortalStart(m, v5898, v5914, v5914, v5914)
	mBase = m.M
	v5918 = m.ExcPending
	if v5918 != 0 {
		goto L1
	} else {
		goto L1442
	}
L1442:
	;
	v5919 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5213)+78)) = uint16(v5919)
	v5921 = *(*int32)(unsafe.Add(mBase, uint32(v5666)+4))
	v5922 = *(*int32)(unsafe.Add(mBase, uint32(v5921)))
	if v5922 != int32(203) {
		goto L1443
	} else {
		goto L1444
	}
L1443:
	;
	F_PortalSetResultFormat(m, v5898, int32(1), v5213+int32(78))
	mBase = m.M
	v5943 = m.ExcPending
	if v5943 != 0 {
		goto L1
	} else {
		goto L1449
	}
L1444:
	;
	v5925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5921)+16)))
	if v5925 != 0 {
		goto L1443
	} else {
		goto L1445
	}
L1445:
	;
	v5926 = *(*int32)(unsafe.Add(mBase, uint32(v5921)+12))
	v5927 = F_GetPortalByName(m, v5926)
	mBase = m.M
	v5928 = m.ExcPending
	if v5928 != 0 {
		goto L1
	} else {
		goto L1446
	}
L1446:
	;
	if v5927 == int32(0) {
		goto L1443
	} else {
		goto L1447
	}
L1447:
	;
	v5931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5927)+76)))
	if v5931&int32(1) == int32(0) {
		goto L1443
	} else {
		goto L1448
	}
L1448:
	;
	v5936 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5213)+78)) = uint16(v5936)
	goto L1443
L1449:
	;
	v5944 = F_CreateDestReceiver(m, v5218)
	mBase = m.M
	v5945 = m.ExcPending
	if v5945 != 0 {
		goto L1
	} else {
		goto L1450
	}
L1450:
	;
	if v5218 == int32(2) {
		goto L1451
	} else {
		goto L1452
	}
L1451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5944)+20)) = v5898
	goto L1453
L1452:
	;
	goto L1453
L1453:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5851
	v5955 = F_PortalRun(m, v5898, int32(2147483647), int32(1), v5944, v5944, v5213+int32(80))
	mBase = m.M
	v5956 = m.ExcPending
	if v5956 != 0 {
		goto L1
	} else {
		goto L1454
	}
L1454:
	;
	v5957 = *(*int32)(unsafe.Add(mBase, uint32(v5944)+12))
	m.T0[v5957].(func(*base.Module, int32))(m, v5944)
	mBase = m.M
	v5959 = m.ExcPending
	if v5959 != 0 {
		goto L1
	} else {
		goto L1455
	}
L1455:
	;
	F_PortalDrop(m, v5898, int32(0))
	mBase = m.M
	v5962 = m.ExcPending
	if v5962 != 0 {
		goto L1
	} else {
		goto L1456
	}
L1456:
	;
	v5963 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v5964 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if base.Ui32(v5963+v5964<<(uint(int32(2))%32)) <= base.Ui32(v5835) {
		goto L1459
	} else {
		goto L1460
	}
L1457:
	;
	F_EndCommand(m, v5213+int32(80), v5218)
	mBase = m.M
	v6022 = m.ExcPending
	if v6022 != 0 {
		goto L1
	} else {
		goto L1489
	}
L1458:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6015 = m.ExcPending
	if v6015 != 0 {
		goto L1
	} else {
		goto L1488
	}
L1459:
	;
	if v5804 == int32(0) {
		goto L1462
	} else {
		goto L1463
	}
L1460:
	;
	goto L1461
L1461:
	;
	v5990 = *(*int32)(unsafe.Add(mBase, uint32(v5666)+4))
	v5991 = *(*int32)(unsafe.Add(mBase, uint32(v5990)))
	if v5991 == int32(225) {
		goto L1475
	} else {
		goto L1476
	}
L1462:
	;
	v5973 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5974 = *(*int32)(unsafe.Add(mBase, uint32(v5973)+24))
	if v5974 == int32(4) {
		goto L1466
	} else {
		goto L1467
	}
L1463:
	;
	goto L1464
L1464:
	;
	v5982 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L1469
L1465:
	;
	goto L1464
L1466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5973)+24)) = int32(1)
	goto L1468
L1467:
	;
	goto L1468
L1468:
	;
	goto L1465
L1469:
	;
	if v5982 != 0 {
		goto L1470
	} else {
		goto L1471
	}
L1470:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v5985 = m.ExcPending
	if v5985 != 0 {
		goto L1
	} else {
		goto L1473
	}
L1471:
	;
	goto L1472
L1472:
	;
	v5987 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v5987 == int32(0) {
		goto L1457
	} else {
		goto L1474
	}
L1473:
	;
	goto L1472
L1474:
	;
	goto L1458
L1475:
	;
	v5997 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L1478
L1476:
	;
	goto L1477
L1477:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v6004 = m.ExcPending
	if v6004 != 0 {
		goto L1
	} else {
		goto L1484
	}
L1478:
	;
	if v5997 != 0 {
		goto L1479
	} else {
		goto L1480
	}
L1479:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6000 = m.ExcPending
	if v6000 != 0 {
		goto L1
	} else {
		goto L1482
	}
L1480:
	;
	goto L1481
L1481:
	;
	v6002 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v6002 != 0 {
		goto L1458
	} else {
		goto L1483
	}
L1482:
	;
	goto L1481
L1483:
	;
	goto L1457
L1484:
	;
	v6008 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L1485
L1485:
	;
	if v6008 == int32(0) {
		goto L1457
	} else {
		goto L1486
	}
L1486:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6013 = m.ExcPending
	if v6013 != 0 {
		goto L1
	} else {
		goto L1487
	}
L1487:
	;
	goto L1457
L1488:
	;
	v6017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v6017)
	goto L1457
L1489:
	;
	if v5849 != 0 {
		goto L1490
	} else {
		goto L1491
	}
L1490:
	;
	F_MemoryContextDelete(m, v5849)
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L1
	} else {
		goto L1493
	}
L1491:
	;
	goto L1492
L1492:
	;
	v6026 = v5647 + int32(1)
	v6027 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if v6026 < v6027 {
		v5647 = v6026
		goto L1371
	} else {
		goto L1494
	}
L1493:
	;
	goto L1492
L1494:
	;
	goto L1372
L1495:
	;
	if v6069 != 0 {
		goto L1496
	} else {
		goto L1497
	}
L1496:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6072 = m.ExcPending
	if v6072 != 0 {
		goto L1
	} else {
		goto L1499
	}
L1497:
	;
	goto L1498
L1498:
	;
	v6074 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v6074 != 0 {
		goto L1500
	} else {
		goto L1501
	}
L1499:
	;
	goto L1498
L1500:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6076 = m.ExcPending
	if v6076 != 0 {
		goto L1
	} else {
		goto L1503
	}
L1501:
	;
	goto L1502
L1502:
	;
	F_NullCommand(m, v5218)
	mBase = m.M
	v6081 = m.ExcPending
	if v6081 != 0 {
		goto L1
	} else {
		goto L1504
	}
L1503:
	;
	v6078 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v6078)
	goto L1502
L1504:
	;
	v6167 = v6058
	v6175 = int32(1)
	goto L1334
L1505:
	;
	if v6123 != 0 {
		goto L1506
	} else {
		goto L1507
	}
L1506:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6126 = m.ExcPending
	if v6126 != 0 {
		goto L1
	} else {
		goto L1509
	}
L1507:
	;
	goto L1508
L1508:
	;
	v6127 = int32(0)
	v6129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v6129 == v6127 {
		v6167 = v5614
		v6175 = v6127
		goto L1334
	} else {
		goto L1510
	}
L1509:
	;
	goto L1508
L1510:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6133 = m.ExcPending
	if v6133 != 0 {
		goto L1
	} else {
		goto L1511
	}
L1511:
	;
	v6135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v6135)
	v6167 = v5614
	v6175 = v6135
	goto L1334
L1512:
	;
	if v5220 != 0 {
		goto L1538
	} else {
		goto L1539
	}
L1513:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v6303, int32(_a_F_PostgresMainLoopOnce_173))
	mBase = m.M
	v6323 = m.ExcPending
	if v6323 != 0 {
		goto L1
	} else {
		goto L1537
	}
L1514:
	;
	v6199 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6200 = m.ExcPending
	if v6200 != 0 {
		goto L1
	} else {
		goto L1521
	}
L1515:
	;
	v6184 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6185 = m.ExcPending
	if v6185 != 0 {
		goto L1
	} else {
		goto L1517
	}
L1516:
	;
	switch v6178 - int32(1) {
	case 0:
		goto L1515
	case 1:
		goto L1514
	default:
		goto L1512
	}
L1517:
	;
	if v6184 == int32(0) {
		goto L1512
	} else {
		goto L1518
	}
L1518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5213))) = v5213 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v5213)
	mBase = m.M
	v6193 = m.ExcPending
	if v6193 != 0 {
		goto L1
	} else {
		goto L1519
	}
L1519:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6195 = m.ExcPending
	if v6195 != 0 {
		goto L1
	} else {
		goto L1520
	}
L1520:
	;
	v6303 = int32(1471)
	goto L1513
L1521:
	;
	if v6199 == int32(0) {
		goto L1512
	} else {
		goto L1522
	}
L1522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5213)+36)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v5213)+32)) = v5213 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_181), v5213+int32(32))
	mBase = m.M
	v6211 = m.ExcPending
	if v6211 != 0 {
		goto L1
	} else {
		goto L1523
	}
L1523:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6213 = m.ExcPending
	if v6213 != 0 {
		goto L1
	} else {
		goto L1524
	}
L1524:
	;
	v6214 = int32(1478)
	if v6175 != 0 {
		v6303 = v6214
		goto L1513
	} else {
		goto L1525
	}
L1525:
	;
	v6215 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	if v6215 <= int32(0) {
		v6303 = v6214
		goto L1513
	} else {
		goto L1526
	}
L1526:
	;
	v6219 = int32(0)
	v6234 = v6215
	goto L1527
L1527:
	;
	v6256 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+12))
	v6260 = *(*int32)(unsafe.Add(mBase, uint32(v6256+v6219<<(uint(int32(2))%32))))
	v6261 = *(*int32)(unsafe.Add(mBase, uint32(v6260)+4))
	v6262 = *(*int32)(unsafe.Add(mBase, uint32(v6261)))
	if v6262 == int32(253) {
		goto L1530
	} else {
		goto L1531
	}
L1528:
	;
	v6275 = *(*int32)(unsafe.Add(mBase, uint32(v6267)+64))
	v6276 = *(*int32)(unsafe.Add(mBase, uint32(v6275)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5213)+16)) = v6276
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_174), v5213+int32(16))
	mBase = m.M
	v6282 = m.ExcPending
	if v6282 != 0 {
		goto L1
	} else {
		goto L1536
	}
L1529:
	;
	goto L1528
L1530:
	;
	v6265 = *(*int32)(unsafe.Add(mBase, uint32(v6261)+4))
	v6267 = F_FetchPreparedStatement(m, v6265, int32(0))
	mBase = m.M
	v6268 = m.ExcPending
	if v6268 != 0 {
		goto L1
	} else {
		goto L1533
	}
L1531:
	;
	v6271 = v6234
	goto L1532
L1532:
	;
	v6273 = v6219 + int32(1)
	if v6273 < v6271 {
		v6219 = v6273
		v6234 = v6271
		goto L1527
	} else {
		goto L1535
	}
L1533:
	;
	if v6267 != 0 {
		goto L1529
	} else {
		goto L1534
	}
L1534:
	;
	v6269 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+4))
	v6271 = v6269
	goto L1532
L1535:
	;
	v6303 = v6214
	goto L1513
L1536:
	;
	v6303 = v6214
	goto L1513
L1537:
	;
	goto L1512
L1538:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_182))
	mBase = m.M
	v6363 = m.ExcPending
	if v6363 != 0 {
		goto L1
	} else {
		goto L1541
	}
L1539:
	;
	goto L1540
L1540:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = int32(0)
	m.G0 = v5213 + int32(112)
	goto L235
L1541:
	;
	goto L1540
L1542:
	;
	v6415 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v6415 < int32(0) {
		goto L1544
	} else {
		goto L1545
	}
L1543:
	;
	v6422 = v40 + int32(440)
	v6423 = F_pq_getmsgstring(m, v6422)
	mBase = m.M
	v6424 = m.ExcPending
	if v6424 != 0 {
		goto L1
	} else {
		goto L1547
	}
L1544:
	;
	v6419 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v6419
	goto L1546
L1545:
	;
	goto L1546
L1546:
	;
	goto L1543
L1547:
	;
	v6425 = F_pq_getmsgstring(m, v6422)
	mBase = m.M
	v6426 = m.ExcPending
	if v6426 != 0 {
		goto L1
	} else {
		goto L1548
	}
L1548:
	;
	v6428 = F_pq_getmsgint(m, v6422, int32(2))
	mBase = m.M
	v6429 = m.ExcPending
	if v6429 != 0 {
		goto L1
	} else {
		goto L1549
	}
L1549:
	;
	if int32(0) < v6428 {
		goto L1550
	} else {
		goto L1551
	}
L1550:
	;
	v6435 = F_palloc(m, v6428<<(uint(int32(2))%32))
	mBase = m.M
	v6436 = m.ExcPending
	if v6436 != 0 {
		goto L1
	} else {
		goto L1553
	}
L1551:
	;
	v6495 = v1
	goto L1552
L1552:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v6526 = m.ExcPending
	if v6526 != 0 {
		goto L1
	} else {
		goto L1558
	}
L1553:
	;
	v6445 = int32(0)
	goto L1554
L1554:
	;
	v6480 = F_pq_getmsgint(m, v40+int32(440), int32(4))
	mBase = m.M
	v6481 = m.ExcPending
	if v6481 != 0 {
		goto L1
	} else {
		goto L1556
	}
L1555:
	;
	v6495 = v6435
	goto L1552
L1556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435+v6445<<(uint(int32(2))%32)))) = v6480
	v6484 = v6445 + int32(1)
	if v6484 != v6428 {
		v6445 = v6484
		goto L1554
	} else {
		goto L1557
	}
L1557:
	;
	goto L1555
L1558:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v6425
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v6495
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v6428
	v6532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])))
	F_pgstat_report_activity(m, int32(3), v6425)
	mBase = m.M
	if v6532 == int32(1) {
		goto L1559
	} else {
		goto L1560
	}
L1559:
	;
	v6537 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v6547 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1562
L1560:
	;
	goto L1561
L1561:
	;
	v6552 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6553 = m.ExcPending
	if v6553 != 0 {
		goto L1
	} else {
		goto L1563
	}
L1562:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1561
L1563:
	;
	if v6552 != 0 {
		goto L1564
	} else {
		goto L1565
	}
L1564:
	;
	v6554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6423))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v6425
	if v6554 != 0 {
		goto L1567
	} else {
		goto L1568
	}
L1565:
	;
	goto L1566
L1566:
	;
	F_start_xact_command(m)
	mBase = m.M
	v6571 = m.ExcPending
	if v6571 != 0 {
		goto L1
	} else {
		goto L1572
	}
L1567:
	;
	v6557 = v6423
	goto L1569
L1568:
	;
	v6557 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1569
L1569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v6557
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_184), v40-int32(-64))
	mBase = m.M
	v6563 = m.ExcPending
	if v6563 != 0 {
		goto L1
	} else {
		goto L1570
	}
L1570:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1526), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v6568 = m.ExcPending
	if v6568 != 0 {
		goto L1
	} else {
		goto L1571
	}
L1571:
	;
	goto L1566
L1572:
	;
	v6572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6423))))
	if v6572 != 0 {
		goto L1574
	} else {
		goto L1575
	}
L1573:
	;
	v6594 = int32(_a_F_PostgresMainLoopOnce_135)
	v6595 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6592
	v6599 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v6599 == int32(1) {
		goto L1582
	} else {
		goto L1583
	}
L1574:
	;
	v6574 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6592 = v6574
	v6593 = int32(0)
	goto L1573
L1575:
	;
	goto L1576
L1576:
	;
	v6577 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120]))
	if v6577 != 0 {
		goto L1577
	} else {
		goto L1578
	}
L1577:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = int32(0)
	F_DropCachedPlan(m, v6577)
	mBase = m.M
	v6582 = m.ExcPending
	if v6582 != 0 {
		goto L1
	} else {
		goto L1580
	}
L1578:
	;
	goto L1579
L1579:
	;
	v6584 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6589 = F_AllocSetContextCreateInternal(m, v6584, int32(_a_F_PostgresMainLoopOnce_186), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v6590 = m.ExcPending
	if v6590 != 0 {
		goto L1
	} else {
		goto L1581
	}
L1580:
	;
	goto L1579
L1581:
	;
	v6592 = v6589
	v6593 = v6589
	goto L1573
L1582:
	;
	v6602 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v6612 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1585
L1583:
	;
	goto L1584
L1584:
	;
	v6616 = F_raw_parser(m, v6425, int32(0))
	mBase = m.M
	v6617 = m.ExcPending
	if v6617 != 0 {
		goto L1
	} else {
		goto L1586
	}
L1585:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1584
L1586:
	;
	v6619 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])))
	if v6619 == int32(1) {
		goto L1587
	} else {
		goto L1588
	}
L1587:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_172))
	mBase = m.M
	v6624 = m.ExcPending
	if v6624 != 0 {
		goto L1
	} else {
		goto L1590
	}
L1588:
	;
	goto L1589
L1589:
	;
	if v6616 != 0 {
		goto L1592
	} else {
		goto L1593
	}
L1590:
	;
	goto L1589
L1591:
	;
	if v6593 != 0 {
		goto L1617
	} else {
		goto L1618
	}
L1592:
	;
	v6625 = *(*int32)(unsafe.Add(mBase, uint32(v6616)+4))
	if int32(2) <= v6625 {
		goto L210
	} else {
		goto L1595
	}
L1593:
	;
	goto L1594
L1594:
	;
	v6682 = int32(0)
	v6685 = F_CreateCachedPlan(m, v6682, v6425, v6682)
	mBase = m.M
	v6686 = m.ExcPending
	if v6686 != 0 {
		goto L1
	} else {
		goto L1616
	}
L1595:
	;
	v6628 = *(*int32)(unsafe.Add(mBase, uint32(v6616)+12))
	v6629 = *(*int32)(unsafe.Add(mBase, uint32(v6628)))
	v6631 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v6632 = *(*int32)(unsafe.Add(mBase, uint32(v6631)+24))
	goto L1596
L1596:
	;
	v6639 = *(*int32)(unsafe.Add(mBase, uint32(v6629)+4))
	if (v6632-int32(7))&int32(-9) == int32(0) {
		goto L1597
	} else {
		goto L1598
	}
L1597:
	;
	if v6639 == int32(0) {
		goto L209
	} else {
		goto L1600
	}
L1598:
	;
	goto L1599
L1599:
	;
	v6650 = F_CreateCommandTag(m, v6639)
	mBase = m.M
	v6651 = m.ExcPending
	if v6651 != 0 {
		goto L1
	} else {
		goto L1603
	}
L1600:
	;
	v6642 = *(*int32)(unsafe.Add(mBase, uint32(v6639)))
	if v6642 != int32(225) {
		goto L209
	} else {
		goto L1601
	}
L1601:
	;
	v6645 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+4))
	if (v6645-int32(2))&int32(-6) != 0 {
		goto L209
	} else {
		goto L1602
	}
L1602:
	;
	goto L1599
L1603:
	;
	v6652 = F_CreateCachedPlan(m, v6629, v6425, v6650)
	mBase = m.M
	v6653 = m.ExcPending
	if v6653 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1604:
	;
	v6656 = *(*int32)(unsafe.Add(mBase, uint32(v6629)+4))
	v6657 = *(*int32)(unsafe.Add(mBase, uint32(v6656)))
	switch v6657 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v6661 = int32(1)
		goto L1606
	default:
		goto L1607
	}
L1605:
	;
	if v6661 == int32(0) {
		goto L1608
	} else {
		goto L1609
	}
L1606:
	;
	goto L1605
L1607:
	;
	v6661 = int32(0)
	goto L1606
L1608:
	;
	v6668 = F_pg_analyze_and_rewrite_varparams(m, v6629, v6425, v40+int32(512), v40+int32(460))
	mBase = m.M
	v6669 = m.ExcPending
	if v6669 != 0 {
		goto L1
	} else {
		goto L1611
	}
L1609:
	;
	goto L1610
L1610:
	;
	v6670 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v6671 = m.ExcPending
	if v6671 != 0 {
		goto L1
	} else {
		goto L1612
	}
L1611:
	;
	v6688 = v6652
	v6689 = v6668
	goto L1591
L1612:
	;
	F_PushActiveSnapshot(m, v6670)
	mBase = m.M
	v6673 = m.ExcPending
	if v6673 != 0 {
		goto L1
	} else {
		goto L1613
	}
L1613:
	;
	v6678 = F_pg_analyze_and_rewrite_varparams(m, v6629, v6425, v40+int32(512), v40+int32(460))
	mBase = m.M
	v6679 = m.ExcPending
	if v6679 != 0 {
		goto L1
	} else {
		goto L1614
	}
L1614:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v6681 = m.ExcPending
	if v6681 != 0 {
		goto L1
	} else {
		goto L1615
	}
L1615:
	;
	v6688 = v6652
	v6689 = v6678
	goto L1591
L1616:
	;
	v6688 = v6685
	v6689 = v6682
	goto L1591
L1617:
	;
	v6691 = *(*int32)(unsafe.Add(mBase, uint32(v6688)+56))
	v6693 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6697 = *(*int32)(unsafe.Add(mBase, uint32(v6691)+16))
	if v6697 != v6693 {
		goto L1621
	} else {
		goto L1622
	}
L1618:
	;
	goto L1619
L1619:
	;
	v6726 = *(*int32)(unsafe.Add(mBase, uint32(v40)+512))
	v6727 = *(*int32)(unsafe.Add(mBase, uint32(v40)+460))
	v6728 = int32(0)
	F_CompleteCachedPlan(m, v6688, v6689, v6593, v6726, v6727, v6728, v6728, int32(2048), int32(1))
	mBase = m.M
	v6733 = m.ExcPending
	if v6733 != 0 {
		goto L1
	} else {
		goto L1637
	}
L1620:
	;
	goto L1619
L1621:
	;
	if v6697 == int32(0) {
		goto L1624
	} else {
		goto L1625
	}
L1622:
	;
	goto L1623
L1623:
	;
	goto L1620
L1624:
	;
	if v6693 != 0 {
		goto L1631
	} else {
		goto L1632
	}
L1625:
	;
	v6701 = *(*int32)(unsafe.Add(mBase, uint32(v6691)+28))
	v6702 = *(*int32)(unsafe.Add(mBase, uint32(v6691)+24))
	if v6702 != 0 {
		goto L1627
	} else {
		goto L1628
	}
L1626:
	;
	if v6701 == int32(0) {
		goto L1624
	} else {
		goto L1630
	}
L1627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6702)+28)) = v6701
	goto L1626
L1628:
	;
	goto L1629
L1629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6697)+20)) = v6701
	goto L1626
L1630:
	;
	v6707 = *(*int32)(unsafe.Add(mBase, uint32(v6691)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6701)+24)) = v6707
	goto L1624
L1631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6691)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6691)+16)) = v6693
	v6714 = *(*int32)(unsafe.Add(mBase, uint32(v6693)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6691)+28)) = v6714
	if v6714 != 0 {
		goto L1634
	} else {
		goto L1635
	}
L1632:
	;
	goto L1633
L1633:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6691)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6691)+16)) = int32(0)
	goto L1623
L1634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6714)+24)) = v6691
	goto L1636
L1635:
	;
	goto L1636
L1636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6693)+20)) = v6691
	goto L1620
L1637:
	;
	v6735 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v6735 != 0 {
		goto L1638
	} else {
		goto L1639
	}
L1638:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6737 = m.ExcPending
	if v6737 != 0 {
		goto L1
	} else {
		goto L1641
	}
L1639:
	;
	goto L1640
L1640:
	;
	if v6572 != 0 {
		goto L1643
	} else {
		goto L1644
	}
L1641:
	;
	goto L1640
L1642:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6595
	F_CommandCounterIncrement(m)
	mBase = m.M
	v6748 = m.ExcPending
	if v6748 != 0 {
		goto L1
	} else {
		goto L1648
	}
L1643:
	;
	F_StorePreparedStatement(m, v6423, v6688, int32(0))
	mBase = m.M
	v6740 = m.ExcPending
	if v6740 != 0 {
		goto L1
	} else {
		goto L1646
	}
L1644:
	;
	goto L1645
L1645:
	;
	F_SaveCachedPlan(m, v6688)
	mBase = m.M
	v6742 = m.ExcPending
	if v6742 != 0 {
		goto L1
	} else {
		goto L1647
	}
L1646:
	;
	goto L1642
L1647:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v6688
	goto L1642
L1648:
	;
	v6750 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v6750 == int32(2) {
		goto L1649
	} else {
		goto L1650
	}
L1649:
	;
	F_pq_putemptymessage(m, int32(49))
	mBase = m.M
	v6755 = m.ExcPending
	if v6755 != 0 {
		goto L1
	} else {
		goto L1652
	}
L1650:
	;
	goto L1651
L1651:
	;
	v6759 = F_check_log_duration(m, v40+int32(480), int32(0))
	mBase = m.M
	v6760 = m.ExcPending
	if v6760 != 0 {
		goto L1
	} else {
		goto L1657
	}
L1652:
	;
	goto L1651
L1653:
	;
	if v6532 != 0 {
		goto L1669
	} else {
		goto L1670
	}
L1654:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6801 = m.ExcPending
	if v6801 != 0 {
		goto L1
	} else {
		goto L1667
	}
L1655:
	;
	v6780 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6781 = m.ExcPending
	if v6781 != 0 {
		goto L1
	} else {
		goto L1661
	}
L1656:
	;
	v6765 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6766 = m.ExcPending
	if v6766 != 0 {
		goto L1
	} else {
		goto L1658
	}
L1657:
	;
	switch v6759 - int32(1) {
	case 0:
		goto L1656
	case 1:
		goto L1655
	default:
		goto L1653
	}
L1658:
	;
	if v6765 == int32(0) {
		goto L1653
	} else {
		goto L1659
	}
L1659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(32))
	mBase = m.M
	v6776 = m.ExcPending
	if v6776 != 0 {
		goto L1
	} else {
		goto L1660
	}
L1660:
	;
	v6799 = int32(1707)
	goto L1654
L1661:
	;
	if v6780 == int32(0) {
		goto L1653
	} else {
		goto L1662
	}
L1662:
	;
	v6784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6423))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v6425
	if v6784 != 0 {
		goto L1663
	} else {
		goto L1664
	}
L1663:
	;
	v6787 = v6423
	goto L1665
L1664:
	;
	v6787 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1665
L1665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+52)) = v6787
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_187), v40+int32(48))
	mBase = m.M
	v6796 = m.ExcPending
	if v6796 != 0 {
		goto L1
	} else {
		goto L1666
	}
L1666:
	;
	v6799 = int32(1715)
	goto L1654
L1667:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v6799, int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v6805 = m.ExcPending
	if v6805 != 0 {
		goto L1
	} else {
		goto L1668
	}
L1668:
	;
	goto L1653
L1669:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_188))
	mBase = m.M
	v6809 = m.ExcPending
	if v6809 != 0 {
		goto L1
	} else {
		goto L1672
	}
L1670:
	;
	goto L1671
L1671:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = int32(0)
	goto L211
L1672:
	;
	goto L1671
L1673:
	;
	v6818 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v6818 < int32(0) {
		goto L1675
	} else {
		goto L1676
	}
L1674:
	;
	v6825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])))
	v6827 = v40 + int32(440)
	v6828 = F_pq_getmsgstring(m, v6827)
	mBase = m.M
	v6829 = m.ExcPending
	if v6829 != 0 {
		goto L1
	} else {
		goto L1678
	}
L1675:
	;
	v6822 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v6822
	goto L1677
L1676:
	;
	goto L1677
L1677:
	;
	goto L1674
L1678:
	;
	v6830 = F_pq_getmsgstring(m, v6827)
	mBase = m.M
	v6831 = m.ExcPending
	if v6831 != 0 {
		goto L1
	} else {
		goto L1679
	}
L1679:
	;
	v6834 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6835 = m.ExcPending
	if v6835 != 0 {
		goto L1
	} else {
		goto L1680
	}
L1680:
	;
	if v6834 != 0 {
		goto L1681
	} else {
		goto L1682
	}
L1681:
	;
	v6836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6828))))
	v6838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6830))))
	if v6838 != 0 {
		goto L1684
	} else {
		goto L1685
	}
L1682:
	;
	goto L1683
L1683:
	;
	v6855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6830))))
	if v6855 != 0 {
		goto L1693
	} else {
		goto L1694
	}
L1684:
	;
	v6839 = v6830
	goto L1686
L1685:
	;
	v6839 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1686
L1686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+212)) = v6839
	if v6836 != 0 {
		goto L1687
	} else {
		goto L1688
	}
L1687:
	;
	v6842 = v6828
	goto L1689
L1688:
	;
	v6842 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1689
L1689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+208)) = v6842
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_189), v40+int32(208))
	mBase = m.M
	v6848 = m.ExcPending
	if v6848 != 0 {
		goto L1
	} else {
		goto L1690
	}
L1690:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1761), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v6853 = m.ExcPending
	if v6853 != 0 {
		goto L1
	} else {
		goto L1691
	}
L1691:
	;
	goto L1683
L1692:
	;
	v6866 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v6866
	F_pgstat_report_activity(m, int32(3), v6866)
	mBase = m.M
	v6870 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+60))
	if v6870 == int32(0) {
		goto L1698
	} else {
		goto L1699
	}
L1693:
	;
	v6857 = F_FetchPreparedStatement(m, v6830, int32(1))
	mBase = m.M
	v6858 = m.ExcPending
	if v6858 != 0 {
		goto L1
	} else {
		goto L1696
	}
L1694:
	;
	goto L1695
L1695:
	;
	v6861 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120]))
	if v6861 == int32(0) {
		goto L208
	} else {
		goto L1697
	}
L1696:
	;
	v6859 = *(*int32)(unsafe.Add(mBase, uint32(v6857)+64))
	v6864 = v6859
	goto L1692
L1697:
	;
	v6864 = v6861
	goto L1692
L1698:
	;
	if v6825&int32(1) != 0 {
		goto L1712
	} else {
		goto L1713
	}
L1699:
	;
	v6873 = *(*int32)(unsafe.Add(mBase, uint32(v6870)+4))
	if v6873 <= int32(0) {
		goto L1698
	} else {
		goto L1700
	}
L1700:
	;
	v6876 = *(*int32)(unsafe.Add(mBase, uint32(v6870)+12))
	v6886 = int32(0)
	goto L1701
L1701:
	;
	v6918 = *(*int32)(unsafe.Add(mBase, uint32(v6876+v6886<<(uint(int32(2))%32))))
	v6919 = *(*int64)(unsafe.Add(mBase, uint32(v6918)+16))
	if v6919 == int64(0) {
		goto L1703
	} else {
		goto L1704
	}
L1702:
	;
	v6928 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v6928 == int32(0) {
		goto L1708
	} else {
		goto L1709
	}
L1703:
	;
	v6923 = v6886 + int32(1)
	if v6923 != v6873 {
		v6886 = v6923
		goto L1701
	} else {
		goto L1706
	}
L1704:
	;
	goto L1705
L1705:
	;
	goto L1702
L1706:
	;
	goto L1698
L1707:
	;
	goto L1698
L1708:
	;
	goto L1707
L1709:
	;
	v6932 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v6932&int32(1) == int32(0) {
		goto L1708
	} else {
		goto L1710
	}
L1710:
	;
	v6939 = *(*int64)(unsafe.Add(mBase, uint32(v6928)+392))
	if int32(1)&base.B2i32(v6939 != int64(0)) != 0 {
		goto L1708
	} else {
		goto L1711
	}
L1711:
	;
	v6943 = int32(_a_F_PostgresMainLoopOnce_176)
	v6945 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v6946 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v6945 + v6946
	v6949 = *(*int32)(unsafe.Add(mBase, uint32(v6928)))
	*(*int32)(unsafe.Add(mBase, uint32(v6928))) = v6949 + v6946
	*(*int64)(unsafe.Add(mBase, uint32(v6928)+392)) = v6919
	*(*int32)(unsafe.Add(mBase, uint32(v6928))) = v6949 + int32(2)
	v6960 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v6960 - v6946
	goto L1708
L1712:
	;
	v7003 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v7013 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1715
L1713:
	;
	goto L1714
L1714:
	;
	F_start_xact_command(m)
	mBase = m.M
	v7017 = m.ExcPending
	if v7017 != 0 {
		goto L1
	} else {
		goto L1716
	}
L1715:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1714
L1716:
	;
	v7020 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7020
	v7025 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7026 = m.ExcPending
	if v7026 != 0 {
		goto L1
	} else {
		goto L1717
	}
L1717:
	;
	if int32(0) < v7025 {
		goto L1718
	} else {
		goto L1719
	}
L1718:
	;
	v7032 = F_palloc(m, v7025<<(uint(int32(1))%32))
	mBase = m.M
	v7033 = m.ExcPending
	if v7033 != 0 {
		goto L1
	} else {
		goto L1721
	}
L1719:
	;
	v7084 = v1
	goto L1720
L1720:
	;
	v7123 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7124 = m.ExcPending
	if v7124 != 0 {
		goto L1
	} else {
		goto L1726
	}
L1721:
	;
	v7042 = int32(0)
	goto L1722
L1722:
	;
	v7077 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7078 = m.ExcPending
	if v7078 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1723:
	;
	v7084 = v7032
	goto L1720
L1724:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7032+v7042<<(uint(int32(1))%32)))) = uint16(v7077)
	v7081 = v7042 + int32(1)
	if v7081 != v7025 {
		v7042 = v7081
		goto L1722
	} else {
		goto L1725
	}
L1725:
	;
	goto L1723
L1726:
	;
	if base.B2i32(v7123 != v7025)&base.B2i32(int32(2) <= v7025) != 0 {
		goto L207
	} else {
		goto L1727
	}
L1727:
	;
	v7129 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+24))
	if v7123 != v7129 {
		goto L206
	} else {
		goto L1728
	}
L1728:
	;
	v7132 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v7133 = *(*int32)(unsafe.Add(mBase, uint32(v7132)+24))
	goto L1729
L1729:
	;
	if (v7133-int32(7))&int32(-9) == int32(0) {
		goto L1730
	} else {
		goto L1731
	}
L1730:
	;
	v7140 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+4))
	if v7140 == int32(0) {
		goto L205
	} else {
		goto L1733
	}
L1731:
	;
	goto L1732
L1732:
	;
	v7156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6828))))
	if v7156 == int32(0) {
		goto L1738
	} else {
		goto L1739
	}
L1733:
	;
	v7143 = *(*int32)(unsafe.Add(mBase, uint32(v7140)+4))
	if v7143 == int32(0) {
		goto L205
	} else {
		goto L1734
	}
L1734:
	;
	v7146 = *(*int32)(unsafe.Add(mBase, uint32(v7143)))
	if v7146 != int32(225) {
		goto L205
	} else {
		goto L1735
	}
L1735:
	;
	v7149 = *(*int32)(unsafe.Add(mBase, uint32(v7143)+4))
	if (v7149-int32(2))&int32(-6)|v7123 != 0 {
		goto L205
	} else {
		goto L1736
	}
L1736:
	;
	goto L1732
L1737:
	;
	v7168 = int32(_a_F_PostgresMainLoopOnce_135)
	v7169 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7171 = *(*int32)(unsafe.Add(mBase, uint32(v7167)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7171
	v7173 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+12))
	v7174 = F_pstrdup(m, v7173)
	mBase = m.M
	v7175 = m.ExcPending
	if v7175 != 0 {
		goto L1
	} else {
		goto L1743
	}
L1738:
	;
	v7159 = int32(1)
	v7161 = F_CreatePortal(m, v6828, v7159, v7159)
	mBase = m.M
	v7162 = m.ExcPending
	if v7162 != 0 {
		goto L1
	} else {
		goto L1741
	}
L1739:
	;
	goto L1740
L1740:
	;
	v7163 = int32(0)
	v7165 = F_CreatePortal(m, v6828, v7163, v7163)
	mBase = m.M
	v7166 = m.ExcPending
	if v7166 != 0 {
		goto L1
	} else {
		goto L1742
	}
L1741:
	;
	v7167 = v7161
	goto L1737
L1742:
	;
	v7167 = v7165
	goto L1737
L1743:
	;
	v7176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6830))))
	if v7176 != 0 {
		goto L1744
	} else {
		goto L1745
	}
L1744:
	;
	v7177 = F_pstrdup(m, v6830)
	mBase = m.M
	v7178 = m.ExcPending
	if v7178 != 0 {
		goto L1
	} else {
		goto L1747
	}
L1745:
	;
	v7179 = v1
	goto L1746
L1746:
	;
	if v7123 <= int32(0) {
		goto L1750
	} else {
		goto L1751
	}
L1747:
	;
	v7179 = v7177
	goto L1746
L1748:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7169
	v7495 = *(*int32)(unsafe.Add(mBase, uint32(v7167)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = v7484
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v7495
	v7498 = int32(_a_F_PostgresMainLoopOnce_191)
	v7499 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v40 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+516)) = int32(1146)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v7499
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v40 + int32(460)
	v7514 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7515 = m.ExcPending
	if v7515 != 0 {
		goto L1
	} else {
		goto L1814
	}
L1749:
	;
	v7484 = v7446
	v7492 = int32(1)
	goto L1748
L1750:
	;
	v7182 = int32(0)
	v7183 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+4))
	if v7183 == v7182 {
		v7484 = v1
		v7492 = v7182
		goto L1748
	} else {
		goto L1753
	}
L1751:
	;
	goto L1752
L1752:
	;
	v7201 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7202 = m.ExcPending
	if v7202 != 0 {
		goto L1
	} else {
		goto L1760
	}
L1753:
	;
	v7189 = *(*int32)(unsafe.Add(mBase, uint32(v7183)+4))
	v7190 = *(*int32)(unsafe.Add(mBase, uint32(v7189)))
	switch v7190 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v7194 = int32(1)
		goto L1755
	default:
		goto L1756
	}
L1754:
	;
	if v7194 == int32(0) {
		v7484 = v1
		v7492 = int32(0)
		goto L1748
	} else {
		goto L1757
	}
L1755:
	;
	goto L1754
L1756:
	;
	v7194 = int32(0)
	goto L1755
L1757:
	;
	v7197 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7198 = m.ExcPending
	if v7198 != 0 {
		goto L1
	} else {
		goto L1758
	}
L1758:
	;
	F_PushActiveSnapshot(m, v7197)
	mBase = m.M
	v7200 = m.ExcPending
	if v7200 != 0 {
		goto L1
	} else {
		goto L1759
	}
L1759:
	;
	v7446 = v1
	goto L1749
L1760:
	;
	F_PushActiveSnapshot(m, v7201)
	mBase = m.M
	v7204 = m.ExcPending
	if v7204 != 0 {
		goto L1
	} else {
		goto L1761
	}
L1761:
	;
	v7205 = *(*int32)(unsafe.Add(mBase, uint32(v7167)))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+464)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v7205
	v7209 = int32(_a_F_PostgresMainLoopOnce_191)
	v7210 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v40 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+516)) = int32(1145)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+512)) = v7210
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v40 + int32(460)
	v7223 = F_makeParamList(m, v7123)
	mBase = m.M
	v7224 = m.ExcPending
	if v7224 != 0 {
		goto L1
	} else {
		goto L1762
	}
L1762:
	;
	v7227 = int32(0)
	v7238 = v7227
	v7255 = v1
	goto L1763
L1763:
	;
	v7268 = v7238 << (uint(int32(2)) % 32)
	v7269 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+20))
	v7271 = *(*int32)(unsafe.Add(mBase, uint32(v7268+v7269)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = v7238
	v7278 = F_pq_getmsgint(m, v40+int32(440), int32(4))
	mBase = m.M
	v7279 = m.ExcPending
	if v7279 != 0 {
		goto L1
	} else {
		goto L1766
	}
L1764:
	;
	v7405 = int32(_a_F_PostgresMainLoopOnce_191)
	v7407 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	v7408 = *(*int32)(unsafe.Add(mBase, uint32(v7407)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v7408
	v7411 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	if v7411 == int32(0) {
		v7446 = v7223
		goto L1749
	} else {
		goto L1812
	}
L1765:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+480)) = v7295
	if int32(2) <= v7025 {
		goto L1775
	} else {
		goto L1776
	}
L1766:
	;
	v7281 = base.B2i32(v7278 == int32(-1))
	if v7278 == int32(-1) {
		goto L1767
	} else {
		goto L1768
	}
L1767:
	;
	v7282 = int32(0)
	v7295 = v7282
	v7297 = v7282
	goto L1765
L1768:
	;
	goto L1769
L1769:
	;
	v7286 = F_pq_getmsgbytes(m, v40+int32(440), v7278)
	mBase = m.M
	v7287 = m.ExcPending
	if v7287 != 0 {
		goto L1
	} else {
		goto L1770
	}
L1770:
	;
	v7288 = v7286 + v7278
	v7289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7288))))
	v7290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7288))) = uint8(v7290)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+488)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+484)) = v7278
	v7295 = v7286
	v7297 = v7289
	goto L1765
L1771:
	;
	if v7281 == int32(0) {
		goto L1808
	} else {
		goto L1809
	}
L1772:
	;
	F_getTypeBinaryInputInfo(m, v7271, v40+int32(472), v40+int32(456))
	mBase = m.M
	v7370 = m.ExcPending
	if v7370 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1773:
	;
	F_getTypeInputInfo(m, v7271, v40+int32(472), v40+int32(456))
	mBase = m.M
	v7312 = m.ExcPending
	if v7312 != 0 {
		goto L1
	} else {
		goto L1779
	}
L1774:
	;
	v7305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7304))))
	switch v7305 {
	case 0:
		goto L1773
	case 1:
		goto L1772
	default:
		goto L203
	}
L1775:
	;
	v7304 = v7084 + v7238<<(uint(int32(1))%32)
	goto L1774
L1776:
	;
	goto L1777
L1777:
	;
	if v7025 <= v7227 {
		goto L1773
	} else {
		goto L1778
	}
L1778:
	;
	v7304 = v7084
	goto L1774
L1779:
	;
	if v7278 == int32(-1) {
		goto L1780
	} else {
		goto L1781
	}
L1780:
	;
	v7317 = int32(0)
	goto L1782
L1781:
	;
	v7314 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	v7315 = F_pg_client_to_server(m, v7314, v7278)
	mBase = m.M
	v7316 = m.ExcPending
	if v7316 != 0 {
		goto L1
	} else {
		goto L1783
	}
L1782:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v7317
	v7319 = *(*int32)(unsafe.Add(mBase, uint32(v40)+472))
	v7320 = *(*int32)(unsafe.Add(mBase, uint32(v40)+456))
	v7322 = F_OidInputFunctionCall(m, v7319, v7317, v7320, int32(-1))
	mBase = m.M
	v7323 = m.ExcPending
	if v7323 != 0 {
		goto L1
	} else {
		goto L1784
	}
L1783:
	;
	v7317 = v7315
	goto L1782
L1784:
	;
	v7324 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v7324
	if v7317 == v7324 {
		v7385 = v7322
		v7388 = v7255
		goto L1771
	} else {
		goto L1785
	}
L1785:
	;
	v7329 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	if v7329 != 0 {
		goto L1786
	} else {
		goto L1787
	}
L1786:
	;
	v7330 = int32(_a_F_PostgresMainLoopOnce_135)
	v7331 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7334 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7334
	if v7255 == int32(0) {
		goto L1790
	} else {
		goto L1791
	}
L1787:
	;
	v7360 = v7255
	goto L1788
L1788:
	;
	v7361 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	if v7317 == v7361 {
		v7385 = v7322
		v7388 = v7360
		goto L1771
	} else {
		goto L1799
	}
L1789:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7268+v7343))) = v7352
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7331
	v7360 = v7343
	goto L1788
L1790:
	;
	v7338 = F_palloc0(m, v7123<<(uint(int32(2))%32))
	mBase = m.M
	v7339 = m.ExcPending
	if v7339 != 0 {
		goto L1
	} else {
		goto L1793
	}
L1791:
	;
	v7342 = v7329
	v7343 = v7255
	goto L1792
L1792:
	;
	if v7342 < int32(0) {
		goto L1794
	} else {
		goto L1795
	}
L1793:
	;
	v7341 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v7342 = v7341
	v7343 = v7338
	goto L1792
L1794:
	;
	v7346 = F_pstrdup(m, v7317)
	mBase = m.M
	v7347 = m.ExcPending
	if v7347 != 0 {
		goto L1
	} else {
		goto L1797
	}
L1795:
	;
	goto L1796
L1796:
	;
	v7350 = F_pnstrdup(m, v7317, v7342+int32(8))
	mBase = m.M
	v7351 = m.ExcPending
	if v7351 != 0 {
		goto L1
	} else {
		goto L1798
	}
L1797:
	;
	v7352 = v7346
	goto L1789
L1798:
	;
	v7352 = v7350
	goto L1789
L1799:
	;
	F_pfree(m, v7317)
	mBase = m.M
	v7364 = m.ExcPending
	if v7364 != 0 {
		goto L1
	} else {
		goto L1800
	}
L1800:
	;
	v7385 = v7322
	v7388 = v7360
	goto L1771
L1801:
	;
	v7371 = *(*int32)(unsafe.Add(mBase, uint32(v40)+472))
	if v7278 == int32(-1) {
		goto L1802
	} else {
		goto L1803
	}
L1802:
	;
	v7375 = int32(0)
	goto L1804
L1803:
	;
	v7375 = v40 + int32(480)
	goto L1804
L1804:
	;
	v7376 = *(*int32)(unsafe.Add(mBase, uint32(v40)+456))
	v7378 = F_OidReceiveFunctionCall(m, v7371, v7375, v7376, int32(-1))
	mBase = m.M
	v7379 = m.ExcPending
	if v7379 != 0 {
		goto L1
	} else {
		goto L1805
	}
L1805:
	;
	if v7278 == int32(-1) {
		v7385 = v7378
		v7388 = v7255
		goto L1771
	} else {
		goto L1806
	}
L1806:
	;
	v7380 = *(*int32)(unsafe.Add(mBase, uint32(v40)+492))
	v7381 = *(*int32)(unsafe.Add(mBase, uint32(v40)+484))
	if v7380 != v7381 {
		goto L204
	} else {
		goto L1807
	}
L1807:
	;
	v7385 = v7378
	v7388 = v7255
	goto L1771
L1808:
	;
	v7391 = *(*int32)(unsafe.Add(mBase, uint32(v40)+480))
	*(*uint8)(unsafe.Add(mBase, uint32(v7391+v7278))) = uint8(v7297)
	goto L1810
L1809:
	;
	goto L1810
L1810:
	;
	v7396 = v7223 + int32(32) + v7238*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v7396)+8)) = v7271
	v7398 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7396)+6)) = uint16(v7398)
	*(*uint8)(unsafe.Add(mBase, uint32(v7396)+4)) = uint8(v7281)
	*(*int32)(unsafe.Add(mBase, uint32(v7396))) = v7385
	v7403 = v7238 + v7398
	if v7403 != v7123 {
		v7238 = v7403
		v7255 = v7388
		goto L1763
	} else {
		goto L1811
	}
L1811:
	;
	goto L1764
L1812:
	;
	v7414 = F_BuildParamLogString(m, v7223, v7388, v7411)
	mBase = m.M
	v7415 = m.ExcPending
	if v7415 != 0 {
		goto L1
	} else {
		goto L1813
	}
L1813:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7223)+24)) = v7414
	v7446 = v7223
	goto L1749
L1814:
	;
	if int32(0) < v7514 {
		goto L1815
	} else {
		goto L1816
	}
L1815:
	;
	v7521 = F_palloc(m, v7514<<(uint(int32(1))%32))
	mBase = m.M
	v7522 = m.ExcPending
	if v7522 != 0 {
		goto L1
	} else {
		goto L1818
	}
L1816:
	;
	v7581 = int32(0)
	goto L1817
L1817:
	;
	F_pq_getmsgend(m, v40+int32(440))
	mBase = m.M
	v7612 = m.ExcPending
	if v7612 != 0 {
		goto L1
	} else {
		goto L1823
	}
L1818:
	;
	v7531 = int32(0)
	goto L1819
L1819:
	;
	v7566 = F_pq_getmsgint(m, v40+int32(440), int32(2))
	mBase = m.M
	v7567 = m.ExcPending
	if v7567 != 0 {
		goto L1
	} else {
		goto L1821
	}
L1820:
	;
	v7581 = v7521
	goto L1817
L1821:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7521+v7531<<(uint(int32(1))%32)))) = uint16(v7566)
	v7570 = v7531 + int32(1)
	if v7570 != v7514 {
		v7531 = v7570
		goto L1819
	} else {
		goto L1822
	}
L1822:
	;
	goto L1820
L1823:
	;
	v7613 = int32(0)
	v7615 = F_GetCachedPlan(m, v6864, v7484, v7613, v7613)
	mBase = m.M
	v7616 = m.ExcPending
	if v7616 != 0 {
		goto L1
	} else {
		goto L1824
	}
L1824:
	;
	v7617 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+16))
	v7618 = *(*int32)(unsafe.Add(mBase, uint32(v7615)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v7167)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+40)) = v7617
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+32)) = v7174
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+4)) = v7179
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+60)) = v7615
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+56)) = v7618
	*(*int32)(unsafe.Add(mBase, uint32(v7167)+36)) = v7617
	goto L1825
L1825:
	;
	v7629 = *(*int32)(unsafe.Add(mBase, uint32(v7167)+56))
	if v7629 == int32(0) {
		goto L1826
	} else {
		goto L1827
	}
L1826:
	;
	if v7492 != 0 {
		goto L1840
	} else {
		goto L1841
	}
L1827:
	;
	v7632 = *(*int32)(unsafe.Add(mBase, uint32(v7629)+4))
	if v7632 <= int32(0) {
		goto L1826
	} else {
		goto L1828
	}
L1828:
	;
	v7635 = *(*int32)(unsafe.Add(mBase, uint32(v7629)+12))
	v7645 = int32(0)
	goto L1829
L1829:
	;
	v7677 = *(*int32)(unsafe.Add(mBase, uint32(v7635+v7645<<(uint(int32(2))%32))))
	v7678 = *(*int64)(unsafe.Add(mBase, uint32(v7677)+16))
	if v7678 == int64(0) {
		goto L1831
	} else {
		goto L1832
	}
L1830:
	;
	v7687 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v7687 == int32(0) {
		goto L1836
	} else {
		goto L1837
	}
L1831:
	;
	v7682 = v7645 + int32(1)
	if v7682 != v7632 {
		v7645 = v7682
		goto L1829
	} else {
		goto L1834
	}
L1832:
	;
	goto L1833
L1833:
	;
	goto L1830
L1834:
	;
	goto L1826
L1835:
	;
	goto L1826
L1836:
	;
	goto L1835
L1837:
	;
	v7691 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v7691&int32(1) == int32(0) {
		goto L1836
	} else {
		goto L1838
	}
L1838:
	;
	v7698 = *(*int64)(unsafe.Add(mBase, uint32(v7687)+400))
	if int32(1)&base.B2i32(v7698 != int64(0)) != 0 {
		goto L1836
	} else {
		goto L1839
	}
L1839:
	;
	v7702 = int32(_a_F_PostgresMainLoopOnce_176)
	v7704 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v7705 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v7704 + v7705
	v7708 = *(*int32)(unsafe.Add(mBase, uint32(v7687)))
	*(*int32)(unsafe.Add(mBase, uint32(v7687))) = v7708 + v7705
	*(*int64)(unsafe.Add(mBase, uint32(v7687)+400)) = v7678
	*(*int32)(unsafe.Add(mBase, uint32(v7687))) = v7708 + int32(2)
	v7719 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v7719 - v7705
	goto L1836
L1840:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v7761 = m.ExcPending
	if v7761 != 0 {
		goto L1
	} else {
		goto L1843
	}
L1841:
	;
	goto L1842
L1842:
	;
	v7762 = int32(0)
	F_PortalStart(m, v7167, v7484, v7762, v7762)
	mBase = m.M
	v7765 = m.ExcPending
	if v7765 != 0 {
		goto L1
	} else {
		goto L1844
	}
L1843:
	;
	goto L1842
L1844:
	;
	F_PortalSetResultFormat(m, v7167, v7514, v7581)
	mBase = m.M
	v7767 = m.ExcPending
	if v7767 != 0 {
		goto L1
	} else {
		goto L1845
	}
L1845:
	;
	v7768 = int32(_a_F_PostgresMainLoopOnce_191)
	v7770 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	v7771 = *(*int32)(unsafe.Add(mBase, uint32(v7770)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v7771
	v7774 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v7774 == int32(2) {
		goto L1846
	} else {
		goto L1847
	}
L1846:
	;
	F_pq_putemptymessage(m, int32(50))
	mBase = m.M
	v7779 = m.ExcPending
	if v7779 != 0 {
		goto L1
	} else {
		goto L1849
	}
L1847:
	;
	goto L1848
L1848:
	;
	v7783 = F_check_log_duration(m, v40+int32(480), int32(0))
	mBase = m.M
	v7784 = m.ExcPending
	if v7784 != 0 {
		goto L1
	} else {
		goto L1854
	}
L1849:
	;
	goto L1848
L1850:
	;
	if v6825&int32(1) != 0 {
		goto L1880
	} else {
		goto L1881
	}
L1851:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v7860, int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v7864 = m.ExcPending
	if v7864 != 0 {
		goto L1
	} else {
		goto L1879
	}
L1852:
	;
	v7806 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7807 = m.ExcPending
	if v7807 != 0 {
		goto L1
	} else {
		goto L1859
	}
L1853:
	;
	v7789 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7790 = m.ExcPending
	if v7790 != 0 {
		goto L1
	} else {
		goto L1855
	}
L1854:
	;
	switch v7783 - int32(1) {
	case 0:
		goto L1853
	case 1:
		goto L1852
	default:
		goto L1850
	}
L1855:
	;
	if v7789 == int32(0) {
		goto L1850
	} else {
		goto L1856
	}
L1856:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(96))
	mBase = m.M
	v7800 = m.ExcPending
	if v7800 != 0 {
		goto L1
	} else {
		goto L1857
	}
L1857:
	;
	F_errhidestmt(m)
	mBase = m.M
	v7802 = m.ExcPending
	if v7802 != 0 {
		goto L1
	} else {
		goto L1858
	}
L1858:
	;
	v7860 = int32(2185)
	goto L1851
L1859:
	;
	if v7806 == int32(0) {
		goto L1850
	} else {
		goto L1860
	}
L1860:
	;
	v7810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6830))))
	v7811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6828))))
	v7812 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+144)) = v7812
	if v7811 != 0 {
		goto L1861
	} else {
		goto L1862
	}
L1861:
	;
	v7815 = v6828
	goto L1863
L1862:
	;
	v7815 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1863
L1863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v7815
	if v7811 != 0 {
		goto L1864
	} else {
		goto L1865
	}
L1864:
	;
	v7819 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L1866
L1865:
	;
	v7819 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1866
L1866:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+136)) = v7819
	if v7810 != 0 {
		goto L1867
	} else {
		goto L1868
	}
L1867:
	;
	v7822 = v6830
	goto L1869
L1868:
	;
	v7822 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1869
L1869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+132)) = v7822
	*(*int32)(unsafe.Add(mBase, uint32(v40)+128)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_193), v40+int32(128))
	mBase = m.M
	v7831 = m.ExcPending
	if v7831 != 0 {
		goto L1
	} else {
		goto L1870
	}
L1870:
	;
	F_errhidestmt(m)
	mBase = m.M
	v7833 = m.ExcPending
	if v7833 != 0 {
		goto L1
	} else {
		goto L1871
	}
L1871:
	;
	v7834 = int32(2196)
	if v7484 == int32(0) {
		v7860 = v7834
		goto L1851
	} else {
		goto L1872
	}
L1872:
	;
	v7837 = *(*int32)(unsafe.Add(mBase, uint32(v7484)+28))
	if v7837 <= int32(0) {
		v7860 = v7834
		goto L1851
	} else {
		goto L1873
	}
L1873:
	;
	v7841 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v7841 == int32(0) {
		v7860 = v7834
		goto L1851
	} else {
		goto L1874
	}
L1874:
	;
	v7845 = F_BuildParamLogString(m, v7484, int32(0), v7841)
	mBase = m.M
	v7846 = m.ExcPending
	if v7846 != 0 {
		goto L1
	} else {
		goto L1875
	}
L1875:
	;
	if v7845 == int32(0) {
		v7860 = v7834
		goto L1851
	} else {
		goto L1876
	}
L1876:
	;
	v7849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7845))))
	if v7849 == int32(0) {
		v7860 = v7834
		goto L1851
	} else {
		goto L1877
	}
L1877:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+112)) = v7845
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(112))
	mBase = m.M
	v7857 = m.ExcPending
	if v7857 != 0 {
		goto L1
	} else {
		goto L1878
	}
L1878:
	;
	v7860 = v7834
	goto L1851
L1879:
	;
	goto L1850
L1880:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_195))
	mBase = m.M
	v7872 = m.ExcPending
	if v7872 != 0 {
		goto L1
	} else {
		goto L1883
	}
L1881:
	;
	goto L1882
L1882:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = int32(0)
	goto L211
L1883:
	;
	goto L1882
L1884:
	;
	v7881 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v7881 < int32(0) {
		goto L1886
	} else {
		goto L1887
	}
L1885:
	;
	v7888 = v40 + int32(440)
	v7889 = F_pq_getmsgstring(m, v7888)
	mBase = m.M
	v7890 = m.ExcPending
	if v7890 != 0 {
		goto L1
	} else {
		goto L1889
	}
L1886:
	;
	v7885 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v7885
	goto L1888
L1887:
	;
	goto L1888
L1888:
	;
	goto L1885
L1889:
	;
	v7892 = F_pq_getmsgint(m, v7888, int32(4))
	mBase = m.M
	v7893 = m.ExcPending
	if v7893 != 0 {
		goto L1
	} else {
		goto L1890
	}
L1890:
	;
	F_pq_getmsgend(m, v7888)
	mBase = m.M
	v7895 = m.ExcPending
	if v7895 != 0 {
		goto L1
	} else {
		goto L1891
	}
L1891:
	;
	v7897 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v7899 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])))
	v7900 = F_GetPortalByName(m, v7889)
	mBase = m.M
	v7901 = m.ExcPending
	if v7901 != 0 {
		goto L1
	} else {
		goto L1892
	}
L1892:
	;
	if v7900 == int32(0) {
		goto L202
	} else {
		goto L1893
	}
L1893:
	;
	if v7897 == int32(2) {
		goto L1894
	} else {
		goto L1895
	}
L1894:
	;
	v7907 = int32(3)
	goto L1896
L1895:
	;
	v7907 = v7897
	goto L1896
L1896:
	;
	v7908 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+36))
	if v7908 == int32(0) {
		goto L1897
	} else {
		goto L1898
	}
L1897:
	;
	F_NullCommand(m, v7907)
	mBase = m.M
	v7912 = m.ExcPending
	if v7912 != 0 {
		goto L1
	} else {
		goto L1900
	}
L1898:
	;
	goto L1899
L1899:
	;
	v7913 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+56))
	if v7913 == int32(0) {
		v7932 = v1
		goto L1901
	} else {
		goto L1902
	}
L1900:
	;
	goto L211
L1901:
	;
	v7933 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+32))
	v7934 = F_pstrdup(m, v7933)
	mBase = m.M
	v7935 = m.ExcPending
	if v7935 != 0 {
		goto L1
	} else {
		goto L1908
	}
L1902:
	;
	v7916 = *(*int32)(unsafe.Add(mBase, uint32(v7913)+4))
	if v7916 != int32(1) {
		v7932 = v1
		goto L1901
	} else {
		goto L1903
	}
L1903:
	;
	v7919 = *(*int32)(unsafe.Add(mBase, uint32(v7913)+12))
	v7920 = *(*int32)(unsafe.Add(mBase, uint32(v7919)))
	v7921 = *(*int32)(unsafe.Add(mBase, uint32(v7920)+4))
	if v7921 == int32(6) {
		goto L1904
	} else {
		goto L1905
	}
L1904:
	;
	v7925 = *(*int32)(unsafe.Add(mBase, uint32(v7920)+88))
	v7926 = *(*int32)(unsafe.Add(mBase, uint32(v7925)))
	if v7926 == int32(225) {
		v7932 = int32(1)
		goto L1901
	} else {
		goto L1907
	}
L1905:
	;
	goto L1906
L1906:
	;
	v7932 = int32(0)
	goto L1901
L1907:
	;
	goto L1906
L1908:
	;
	v7936 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+4))
	if v7936 != 0 {
		goto L1909
	} else {
		goto L1910
	}
L1909:
	;
	v7937 = F_pstrdup(m, v7936)
	mBase = m.M
	v7938 = m.ExcPending
	if v7938 != 0 {
		goto L1
	} else {
		goto L1912
	}
L1910:
	;
	v7940 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1911
L1911:
	;
	v7941 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v7934
	F_pgstat_report_activity(m, int32(3), v7934)
	mBase = m.M
	v7946 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+56))
	if v7946 == int32(0) {
		goto L1913
	} else {
		goto L1914
	}
L1912:
	;
	v7940 = v7937
	goto L1911
L1913:
	;
	v8183 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+36))
	v8188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8183<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[127]))))
	*(*int32)(unsafe.Add(mBase, uint32(v40+int32(456)))) = v8188
	goto L1947
L1914:
	;
	v7949 = *(*int32)(unsafe.Add(mBase, uint32(v7946)+4))
	if v7949 <= int32(0) {
		goto L1913
	} else {
		goto L1915
	}
L1915:
	;
	v7952 = int32(0)
	if v7952 < v7949 {
		goto L1916
	} else {
		goto L1917
	}
L1916:
	;
	v7955 = v7949
	goto L1918
L1917:
	;
	v7955 = v7952
	goto L1918
L1918:
	;
	v7956 = *(*int32)(unsafe.Add(mBase, uint32(v7946)+12))
	v7966 = int32(0)
	goto L1920
L1919:
	;
	if v8054 <= int32(0) {
		goto L1913
	} else {
		goto L1935
	}
L1920:
	;
	v7998 = *(*int32)(unsafe.Add(mBase, uint32(v7956+v7966<<(uint(int32(2))%32))))
	v7999 = *(*int64)(unsafe.Add(mBase, uint32(v7998)+8))
	if v7999 == int64(0) {
		goto L1922
	} else {
		goto L1923
	}
L1921:
	;
	v8008 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v8008 == int32(0) {
		goto L1927
	} else {
		goto L1928
	}
L1922:
	;
	v8003 = v7966 + int32(1)
	if v8003 != v7949 {
		v7966 = v8003
		goto L1920
	} else {
		goto L1925
	}
L1923:
	;
	goto L1924
L1924:
	;
	goto L1921
L1925:
	;
	v8052 = v7955
	v8054 = v7949
	v8055 = v7946
	goto L1919
L1926:
	;
	v8044 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+56))
	if v8044 == int32(0) {
		goto L1913
	} else {
		goto L1931
	}
L1927:
	;
	goto L1926
L1928:
	;
	v8012 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v8012&int32(1) == int32(0) {
		goto L1927
	} else {
		goto L1929
	}
L1929:
	;
	v8019 = *(*int64)(unsafe.Add(mBase, uint32(v8008)+392))
	if int32(1)&base.B2i32(v8019 != int64(0)) != 0 {
		goto L1927
	} else {
		goto L1930
	}
L1930:
	;
	v8023 = int32(_a_F_PostgresMainLoopOnce_176)
	v8025 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v8026 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v8025 + v8026
	v8029 = *(*int32)(unsafe.Add(mBase, uint32(v8008)))
	*(*int32)(unsafe.Add(mBase, uint32(v8008))) = v8029 + v8026
	*(*int64)(unsafe.Add(mBase, uint32(v8008)+392)) = v7999
	*(*int32)(unsafe.Add(mBase, uint32(v8008))) = v8029 + int32(2)
	v8040 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v8040 - v8026
	goto L1927
L1931:
	;
	v8047 = *(*int32)(unsafe.Add(mBase, uint32(v8044)+4))
	v8048 = int32(0)
	if v8048 < v8047 {
		goto L1932
	} else {
		goto L1933
	}
L1932:
	;
	v8051 = v8047
	goto L1934
L1933:
	;
	v8051 = v8048
	goto L1934
L1934:
	;
	v8052 = v8051
	v8054 = v8047
	v8055 = v8044
	goto L1919
L1935:
	;
	v8058 = *(*int32)(unsafe.Add(mBase, uint32(v8055)+12))
	v8068 = int32(0)
	goto L1936
L1936:
	;
	v8100 = *(*int32)(unsafe.Add(mBase, uint32(v8058+v8068<<(uint(int32(2))%32))))
	v8101 = *(*int64)(unsafe.Add(mBase, uint32(v8100)+16))
	if v8101 == int64(0) {
		goto L1938
	} else {
		goto L1939
	}
L1937:
	;
	v8110 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124]))
	if v8110 == int32(0) {
		goto L1943
	} else {
		goto L1944
	}
L1938:
	;
	v8105 = v8068 + int32(1)
	if v8052 != v8105 {
		v8068 = v8105
		goto L1936
	} else {
		goto L1941
	}
L1939:
	;
	goto L1940
L1940:
	;
	goto L1937
L1941:
	;
	goto L1913
L1942:
	;
	goto L1913
L1943:
	;
	goto L1942
L1944:
	;
	v8114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125])))
	if v8114&int32(1) == int32(0) {
		goto L1943
	} else {
		goto L1945
	}
L1945:
	;
	v8121 = *(*int64)(unsafe.Add(mBase, uint32(v8110)+400))
	if int32(1)&base.B2i32(v8121 != int64(0)) != 0 {
		goto L1943
	} else {
		goto L1946
	}
L1946:
	;
	v8125 = int32(_a_F_PostgresMainLoopOnce_176)
	v8127 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	v8128 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v8127 + v8128
	v8131 = *(*int32)(unsafe.Add(mBase, uint32(v8110)))
	*(*int32)(unsafe.Add(mBase, uint32(v8110))) = v8131 + v8128
	*(*int64)(unsafe.Add(mBase, uint32(v8110)+400)) = v8101
	*(*int32)(unsafe.Add(mBase, uint32(v8110))) = v8131 + int32(2)
	v8142 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])) = v8142 - v8128
	goto L1943
L1947:
	;
	if v7899&int32(1) != 0 {
		goto L1948
	} else {
		goto L1949
	}
L1948:
	;
	v8194 = int32(_a_F_PostgresMainLoopOnce_170)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = int64(1)
	v8204 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L1951
L1949:
	;
	goto L1950
L1950:
	;
	v8208 = F_CreateDestReceiver(m, v7907)
	mBase = m.M
	v8209 = m.ExcPending
	if v8209 != 0 {
		goto L1
	} else {
		goto L1952
	}
L1951:
	;
	F_gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_171))
	mBase = m.M
	goto L1950
L1952:
	;
	if v7907 == int32(3) {
		goto L1953
	} else {
		goto L1954
	}
L1953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8208)+20)) = v7900
	goto L1955
L1954:
	;
	goto L1955
L1955:
	;
	F_start_xact_command(m)
	mBase = m.M
	v8214 = m.ExcPending
	if v8214 != 0 {
		goto L1
	} else {
		goto L1956
	}
L1956:
	;
	v8215 = int32(0)
	v8216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7900)+116)))
	v8218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122]))
	switch v8218 {
	case 0:
		v8382 = v8215
		goto L1957
	default:
		goto L1959
	case 3:
		goto L1958
	}
L1957:
	;
	v8412 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8413 = *(*int32)(unsafe.Add(mBase, uint32(v8412)+24))
	goto L1991
L1958:
	;
	v8316 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8317 = m.ExcPending
	if v8317 != 0 {
		goto L1
	} else {
		goto L1967
	}
L1959:
	;
	v8219 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+56))
	if v8219 == int32(0) {
		v8382 = v8215
		goto L1957
	} else {
		goto L1960
	}
L1960:
	;
	v8222 = *(*int32)(unsafe.Add(mBase, uint32(v8219)+4))
	if v8222 <= int32(0) {
		v8382 = v8215
		goto L1957
	} else {
		goto L1961
	}
L1961:
	;
	v8233 = v8215
	goto L1962
L1962:
	;
	v8262 = *(*int32)(unsafe.Add(mBase, uint32(v8219)+12))
	v8266 = *(*int32)(unsafe.Add(mBase, uint32(v8262+v8233<<(uint(int32(2))%32))))
	v8267 = F_GetCommandLogLevel(m, v8266)
	mBase = m.M
	v8268 = m.ExcPending
	if v8268 != 0 {
		goto L1
	} else {
		goto L1964
	}
L1963:
	;
	v8382 = int32(0)
	goto L1957
L1964:
	;
	v8270 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122]))
	if base.Ui32(v8267) <= base.Ui32(v8270) {
		goto L1958
	} else {
		goto L1965
	}
L1965:
	;
	v8273 = v8233 + int32(1)
	v8274 = *(*int32)(unsafe.Add(mBase, uint32(v8219)+4))
	if v8273 < v8274 {
		v8233 = v8273
		goto L1962
	} else {
		goto L1966
	}
L1966:
	;
	goto L1963
L1967:
	;
	if v8316 == int32(0) {
		goto L1968
	} else {
		goto L1969
	}
L1968:
	;
	v8382 = int32(1)
	goto L1957
L1969:
	;
	goto L1970
L1970:
	;
	v8321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7889))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+336)) = v7934
	*(*int32)(unsafe.Add(mBase, uint32(v40)+324)) = v7940
	if v8321 != 0 {
		goto L1971
	} else {
		goto L1972
	}
L1971:
	;
	v8325 = v7889
	goto L1973
L1972:
	;
	v8325 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1973
L1973:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+332)) = v8325
	if v8321 != 0 {
		goto L1974
	} else {
		goto L1975
	}
L1974:
	;
	v8329 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L1976
L1975:
	;
	v8329 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L1976
L1976:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+328)) = v8329
	v8331 = int32(1)
	if v8216&v8331 != 0 {
		goto L1977
	} else {
		goto L1978
	}
L1977:
	;
	v8336 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L1979
L1978:
	;
	v8336 = int32(_a_F_PostgresMainLoopOnce_197)
	goto L1979
L1979:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+320)) = v8336
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_198), v40+int32(320))
	mBase = m.M
	v8342 = m.ExcPending
	if v8342 != 0 {
		goto L1
	} else {
		goto L1980
	}
L1980:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8344 = m.ExcPending
	if v8344 != 0 {
		goto L1
	} else {
		goto L1981
	}
L1981:
	;
	if v7941 == int32(0) {
		goto L1982
	} else {
		goto L1983
	}
L1982:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2346), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v8373 = m.ExcPending
	if v8373 != 0 {
		goto L1
	} else {
		goto L1990
	}
L1983:
	;
	v8347 = *(*int32)(unsafe.Add(mBase, uint32(v7941)+28))
	if v8347 <= int32(0) {
		goto L1982
	} else {
		goto L1984
	}
L1984:
	;
	v8351 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v8351 == int32(0) {
		goto L1982
	} else {
		goto L1985
	}
L1985:
	;
	v8355 = F_BuildParamLogString(m, v7941, int32(0), v8351)
	mBase = m.M
	v8356 = m.ExcPending
	if v8356 != 0 {
		goto L1
	} else {
		goto L1986
	}
L1986:
	;
	if v8355 == int32(0) {
		goto L1982
	} else {
		goto L1987
	}
L1987:
	;
	v8359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8355))))
	if v8359 == int32(0) {
		goto L1982
	} else {
		goto L1988
	}
L1988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+304)) = v8355
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(304))
	mBase = m.M
	v8367 = m.ExcPending
	if v8367 != 0 {
		goto L1
	} else {
		goto L1989
	}
L1989:
	;
	goto L1982
L1990:
	;
	v8382 = v8331
	goto L1957
L1991:
	;
	if (v8413-int32(7))&int32(-9) == int32(0) {
		goto L1992
	} else {
		goto L1993
	}
L1992:
	;
	v8420 = *(*int32)(unsafe.Add(mBase, uint32(v7900)+56))
	if v8420 == int32(0) {
		goto L201
	} else {
		goto L1995
	}
L1993:
	;
	goto L1994
L1994:
	;
	v8444 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v8444 != 0 {
		goto L2001
	} else {
		goto L2002
	}
L1995:
	;
	v8423 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+4))
	if v8423 != int32(1) {
		goto L201
	} else {
		goto L1996
	}
L1996:
	;
	v8426 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+12))
	v8427 = *(*int32)(unsafe.Add(mBase, uint32(v8426)))
	v8428 = *(*int32)(unsafe.Add(mBase, uint32(v8427)+4))
	if v8428 != int32(6) {
		goto L201
	} else {
		goto L1997
	}
L1997:
	;
	v8431 = *(*int32)(unsafe.Add(mBase, uint32(v8427)+88))
	if v8431 == int32(0) {
		goto L201
	} else {
		goto L1998
	}
L1998:
	;
	v8434 = *(*int32)(unsafe.Add(mBase, uint32(v8431)))
	if v8434 != int32(225) {
		goto L201
	} else {
		goto L1999
	}
L1999:
	;
	v8437 = *(*int32)(unsafe.Add(mBase, uint32(v8431)+4))
	if (v8437-int32(2))&int32(-6) != 0 {
		goto L201
	} else {
		goto L2000
	}
L2000:
	;
	goto L1994
L2001:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v8446 = m.ExcPending
	if v8446 != 0 {
		goto L1
	} else {
		goto L2004
	}
L2002:
	;
	goto L2003
L2003:
	;
	v8447 = *(*int32)(unsafe.Add(mBase, uint32(v7900)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+476)) = v7941
	*(*int32)(unsafe.Add(mBase, uint32(v40)+472)) = v8447
	v8450 = int32(_a_F_PostgresMainLoopOnce_191)
	v8451 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v40 + int32(460)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+464)) = int32(1146)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+460)) = v8451
	*(*int32)(unsafe.Add(mBase, uint32(v40)+468)) = v40 + int32(472)
	if v7892 <= int32(0) {
		goto L2005
	} else {
		goto L2006
	}
L2004:
	;
	goto L2003
L2005:
	;
	v8465 = int32(2147483647)
	goto L2007
L2006:
	;
	v8465 = v7892
	goto L2007
L2007:
	;
	v8469 = F_PortalRun(m, v7900, v8465, int32(1), v8208, v8208, v40+int32(512))
	mBase = m.M
	v8470 = m.ExcPending
	if v8470 != 0 {
		goto L1
	} else {
		goto L2008
	}
L2008:
	;
	v8471 = *(*int32)(unsafe.Add(mBase, uint32(v8208)+12))
	m.T0[v8471].(func(*base.Module, int32))(m, v8208)
	mBase = m.M
	v8473 = m.ExcPending
	if v8473 != 0 {
		goto L1
	} else {
		goto L2009
	}
L2009:
	;
	v8474 = int32(_a_F_PostgresMainLoopOnce_191)
	v8476 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129]))
	v8477 = *(*int32)(unsafe.Add(mBase, uint32(v8476)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = v8477
	if v8469 != 0 {
		goto L2011
	} else {
		goto L2012
	}
L2010:
	;
	v8542 = F_check_log_duration(m, v40+int32(480), v8382)
	mBase = m.M
	v8543 = m.ExcPending
	if v8543 != 0 {
		goto L1
	} else {
		goto L2040
	}
L2011:
	;
	if v7932 == int32(0) {
		goto L2016
	} else {
		goto L2017
	}
L2012:
	;
	goto L2013
L2013:
	;
	v8527 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v8527 == int32(2) {
		goto L2032
	} else {
		goto L2033
	}
L2014:
	;
	F_EndCommand(m, v40+int32(512), v7907)
	mBase = m.M
	v8525 = m.ExcPending
	if v8525 != 0 {
		goto L1
	} else {
		goto L2031
	}
L2015:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v8505 = m.ExcPending
	if v8505 != 0 {
		goto L1
	} else {
		goto L2027
	}
L2016:
	;
	v8482 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132])))
	if v8482&int32(4) == int32(0) {
		goto L2015
	} else {
		goto L2019
	}
L2017:
	;
	goto L2018
L2018:
	;
	v8490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L2020
L2019:
	;
	goto L2018
L2020:
	;
	if v8490 != 0 {
		goto L2021
	} else {
		goto L2022
	}
L2021:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8493 = m.ExcPending
	if v8493 != 0 {
		goto L1
	} else {
		goto L2024
	}
L2022:
	;
	goto L2023
L2023:
	;
	v8494 = int32(0)
	v8496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v8496 == v8494 {
		v8521 = v8494
		goto L2014
	} else {
		goto L2025
	}
L2024:
	;
	goto L2023
L2025:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8500 = m.ExcPending
	if v8500 != 0 {
		goto L1
	} else {
		goto L2026
	}
L2026:
	;
	v8502 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v8502)
	v8521 = v8494
	goto L2014
L2027:
	;
	v8506 = int32(_a_F_PostgresMainLoopOnce_200)
	v8508 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132])) = v8508 | int32(8)
	v8515 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L2028
L2028:
	;
	if v8515 == int32(0) {
		v8521 = v7941
		goto L2014
	} else {
		goto L2029
	}
L2029:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8520 = m.ExcPending
	if v8520 != 0 {
		goto L1
	} else {
		goto L2030
	}
L2030:
	;
	v8521 = v7941
	goto L2014
L2031:
	;
	v8539 = v8521
	goto L2010
L2032:
	;
	F_pq_putemptymessage(m, int32(115))
	mBase = m.M
	v8532 = m.ExcPending
	if v8532 != 0 {
		goto L1
	} else {
		goto L2035
	}
L2033:
	;
	goto L2034
L2034:
	;
	v8533 = int32(_a_F_PostgresMainLoopOnce_200)
	v8535 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132])) = v8535 | int32(8)
	v8539 = v7941
	goto L2010
L2035:
	;
	goto L2034
L2036:
	;
	if v7899&int32(1) != 0 {
		goto L2066
	} else {
		goto L2067
	}
L2037:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v8620, int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v8624 = m.ExcPending
	if v8624 != 0 {
		goto L1
	} else {
		goto L2065
	}
L2038:
	;
	v8565 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8566 = m.ExcPending
	if v8566 != 0 {
		goto L1
	} else {
		goto L2045
	}
L2039:
	;
	v8548 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8549 = m.ExcPending
	if v8549 != 0 {
		goto L1
	} else {
		goto L2041
	}
L2040:
	;
	switch v8542 - int32(1) {
	case 0:
		goto L2039
	case 1:
		goto L2038
	default:
		goto L2036
	}
L2041:
	;
	if v8548 == int32(0) {
		goto L2036
	} else {
		goto L2042
	}
L2042:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+240)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v40+int32(240))
	mBase = m.M
	v8559 = m.ExcPending
	if v8559 != 0 {
		goto L1
	} else {
		goto L2043
	}
L2043:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8561 = m.ExcPending
	if v8561 != 0 {
		goto L1
	} else {
		goto L2044
	}
L2044:
	;
	v8620 = int32(2457)
	goto L2037
L2045:
	;
	if v8565 == int32(0) {
		goto L2036
	} else {
		goto L2046
	}
L2046:
	;
	v8569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7889))))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+292)) = v7934
	if v8569 != 0 {
		goto L2047
	} else {
		goto L2048
	}
L2047:
	;
	v8572 = v7889
	goto L2049
L2048:
	;
	v8572 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L2049
L2049:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+288)) = v8572
	*(*int32)(unsafe.Add(mBase, uint32(v40)+280)) = v7940
	if v8569 != 0 {
		goto L2050
	} else {
		goto L2051
	}
L2050:
	;
	v8577 = int32(_a_F_PostgresMainLoopOnce_192)
	goto L2052
L2051:
	;
	v8577 = int32(_a_F_PostgresMainLoopOnce_179)
	goto L2052
L2052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+284)) = v8577
	if v8216&int32(1) != 0 {
		goto L2053
	} else {
		goto L2054
	}
L2053:
	;
	v8583 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L2055
L2054:
	;
	v8583 = int32(_a_F_PostgresMainLoopOnce_197)
	goto L2055
L2055:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+276)) = v8583
	*(*int32)(unsafe.Add(mBase, uint32(v40)+272)) = v40 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_201), v40+int32(272))
	mBase = m.M
	v8592 = m.ExcPending
	if v8592 != 0 {
		goto L1
	} else {
		goto L2056
	}
L2056:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8594 = m.ExcPending
	if v8594 != 0 {
		goto L1
	} else {
		goto L2057
	}
L2057:
	;
	v8595 = int32(2471)
	if v8539 == int32(0) {
		v8620 = v8595
		goto L2037
	} else {
		goto L2058
	}
L2058:
	;
	v8598 = *(*int32)(unsafe.Add(mBase, uint32(v8539)+28))
	if v8598 <= int32(0) {
		v8620 = v8595
		goto L2037
	} else {
		goto L2059
	}
L2059:
	;
	v8602 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v8602 == int32(0) {
		v8620 = v8595
		goto L2037
	} else {
		goto L2060
	}
L2060:
	;
	v8606 = F_BuildParamLogString(m, v8539, int32(0), v8602)
	mBase = m.M
	v8607 = m.ExcPending
	if v8607 != 0 {
		goto L1
	} else {
		goto L2061
	}
L2061:
	;
	if v8606 == int32(0) {
		v8620 = v8595
		goto L2037
	} else {
		goto L2062
	}
L2062:
	;
	v8610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8606))))
	if v8610 == int32(0) {
		v8620 = v8595
		goto L2037
	} else {
		goto L2063
	}
L2063:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+256)) = v8606
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_194), v40+int32(256))
	mBase = m.M
	v8618 = m.ExcPending
	if v8618 != 0 {
		goto L1
	} else {
		goto L2064
	}
L2064:
	;
	v8620 = v8595
	goto L2037
L2065:
	;
	goto L2036
L2066:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_202))
	mBase = m.M
	v8631 = m.ExcPending
	if v8631 != 0 {
		goto L1
	} else {
		goto L2069
	}
L2067:
	;
	goto L2068
L2068:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = int32(0)
	goto L211
L2069:
	;
	goto L2068
L2070:
	;
	v8640 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v8640 < int32(0) {
		goto L2072
	} else {
		goto L2073
	}
L2071:
	;
	F_pgstat_report_activity(m, int32(5), int32(0))
	mBase = m.M
	F_start_xact_command(m)
	mBase = m.M
	v8650 = m.ExcPending
	if v8650 != 0 {
		goto L1
	} else {
		goto L2075
	}
L2072:
	;
	v8644 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v8644
	goto L2074
L2073:
	;
	goto L2074
L2074:
	;
	goto L2071
L2075:
	;
	v8653 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v8653
	v8656 = v40 + int32(440)
	v8657 = m.G0
	v8659 = v8657 - int32(1568)
	m.G0 = v8659
	v8662 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8663 = *(*int32)(unsafe.Add(mBase, uint32(v8662)+24))
	goto L2086
L2076:
	;
	v9506 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v9506 != 0 {
		goto L2259
	} else {
		goto L2260
	}
L2077:
	;
	v9464 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+240))
	v9465 = m.T0[v9464].(func(*base.Module, int32) int32)(m, v8659+int32(740))
	mBase = m.M
	v9466 = m.ExcPending
	if v9466 != 0 {
		goto L1
	} else {
		goto L2258
	}
L2078:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9409 = m.ExcPending
	if v9409 != 0 {
		goto L1
	} else {
		goto L2254
	}
L2079:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9389 = m.ExcPending
	if v9389 != 0 {
		goto L1
	} else {
		goto L2250
	}
L2080:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9371 = m.ExcPending
	if v9371 != 0 {
		goto L1
	} else {
		goto L2246
	}
L2081:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9352 = m.ExcPending
	if v9352 != 0 {
		goto L1
	} else {
		goto L2242
	}
L2082:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9332 = m.ExcPending
	if v9332 != 0 {
		goto L1
	} else {
		goto L2238
	}
L2083:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9313 = m.ExcPending
	if v9313 != 0 {
		goto L1
	} else {
		goto L2235
	}
L2084:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9293 = m.ExcPending
	if v9293 != 0 {
		goto L1
	} else {
		goto L2231
	}
L2085:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9277 = m.ExcPending
	if v9277 != 0 {
		goto L1
	} else {
		goto L2227
	}
L2086:
	;
	if base.B2i32((v8663-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L2087
	} else {
		goto L2088
	}
L2087:
	;
	v8672 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v8673 = m.ExcPending
	if v8673 != 0 {
		goto L1
	} else {
		goto L2090
	}
L2088:
	;
	goto L2089
L2089:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9261 = m.ExcPending
	if v9261 != 0 {
		goto L1
	} else {
		goto L2223
	}
L2090:
	;
	F_PushActiveSnapshot(m, v8672)
	mBase = m.M
	v8675 = m.ExcPending
	if v8675 != 0 {
		goto L1
	} else {
		goto L2091
	}
L2091:
	;
	v8677 = F_pq_getmsgint(m, v8656, int32(4))
	mBase = m.M
	v8678 = m.ExcPending
	if v8678 != 0 {
		goto L1
	} else {
		goto L2092
	}
L2092:
	;
	base.MemoryFill(m, v8659+int32(236), int32(0), int32(504))
	v8685 = F_SearchSysCache1(m, int32(47), v8677)
	mBase = m.M
	v8686 = m.ExcPending
	if v8686 != 0 {
		goto L1
	} else {
		goto L2093
	}
L2093:
	;
	if v8685 == int32(0) {
		goto L2085
	} else {
		goto L2094
	}
L2094:
	;
	v8689 = *(*int32)(unsafe.Add(mBase, uint32(v8685)+16))
	v8690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8689)+22)))
	v8691 = v8689 + v8690
	v8692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8691)+96)))
	if v8692 != int32(102) {
		goto L2084
	} else {
		goto L2095
	}
L2095:
	;
	v8695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8691)+100)))
	if v8695 == int32(1) {
		goto L2084
	} else {
		goto L2096
	}
L2096:
	;
	v8698 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8691)+104)))
	if int32(101) <= v8698 {
		goto L2083
	} else {
		goto L2097
	}
L2097:
	;
	v8701 = *(*int32)(unsafe.Add(mBase, uint32(v8691)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+268)) = v8701
	v8703 = *(*int32)(unsafe.Add(mBase, uint32(v8691)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+272)) = v8703
	v8706 = v8659 + int32(276)
	v8707 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8691)+104)))
	v8709 = v8707 << (uint(int32(2)) % 32)
	if v8709 != 0 {
		goto L2098
	} else {
		goto L2099
	}
L2098:
	;
	base.MemoryCopy(m, v8706, v8691+int32(136), v8709)
	goto L2100
L2099:
	;
	goto L2100
L2100:
	;
	v8714 = v8659 + int32(240)
	v8716 = v8659 + int32(676)
	v8718 = v8691 + int32(4)
	goto L2104
L2101:
	;
	F_ReleaseCatCache(m, v8685)
	mBase = m.M
	v8839 = m.ExcPending
	if v8839 != 0 {
		goto L1
	} else {
		goto L2132
	}
L2102:
	;
	v8835 = F_strlen(m, v8824)
	mBase = m.M
	goto L2101
L2104:
	;
	goto L2105
L2105:
	;
	v8725 = int32(63)
	if (v8716^v8718)&int32(3) != 0 {
		goto L2109
	} else {
		goto L2110
	}
L2106:
	;
	v8828 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8825))) = uint8(v8828)
	goto L2102
L2107:
	;
	v8809 = v8804
	v8810 = v8805
	v8811 = v8806
	goto L2128
L2108:
	;
	if v8799 == int32(0) {
		v8824 = v8797
		v8825 = v8798
		goto L2106
	} else {
		goto L2127
	}
L2109:
	;
	v8797 = v8718
	v8798 = v8716
	v8799 = v8725
	goto L2108
L2110:
	;
	goto L2111
L2111:
	;
	v8729 = int32(0)
	if base.B2i32(v8718&int32(3) == v8729)|int32(0) == v8729 {
		goto L2113
	} else {
		goto L2114
	}
L2112:
	;
	if v8765 == int32(0) {
		v8824 = v8762
		v8825 = v8763
		goto L2106
	} else {
		goto L2121
	}
L2113:
	;
	v8741 = v8718
	v8742 = v8716
	v8743 = v8725
	goto L2116
L2114:
	;
	goto L2115
L2115:
	;
	v8762 = v8718
	v8763 = v8716
	v8764 = v8725
	v8765 = int32(1)
	goto L2112
L2116:
	;
	v8745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8741))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8742))) = uint8(v8745)
	if v8745 == int32(0) {
		v8804 = v8741
		v8805 = v8742
		v8806 = v8743
		goto L2107
	} else {
		goto L2118
	}
L2117:
	;
	v8762 = v8756
	v8763 = v8750
	v8764 = v8752
	v8765 = v8754
	goto L2112
L2118:
	;
	v8749 = int32(1)
	v8750 = v8742 + v8749
	v8752 = v8743 - v8749
	v8753 = int32(0)
	v8754 = base.B2i32(v8752 != v8753)
	v8756 = v8741 + v8749
	if v8756&int32(3) == v8753 {
		v8762 = v8756
		v8763 = v8750
		v8764 = v8752
		v8765 = v8754
		goto L2112
	} else {
		goto L2119
	}
L2119:
	;
	if v8752 != 0 {
		v8741 = v8756
		v8742 = v8750
		v8743 = v8752
		goto L2116
	} else {
		goto L2120
	}
L2120:
	;
	goto L2117
L2121:
	;
	v8768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8762))))
	if base.B2i32(v8768 == int32(0))|base.B2i32(base.Ui32(v8764) < base.Ui32(int32(4))) != 0 {
		v8797 = v8762
		v8798 = v8763
		v8799 = v8764
		goto L2108
	} else {
		goto L2122
	}
L2122:
	;
	v8775 = v8762
	v8776 = v8763
	v8777 = v8764
	goto L2123
L2123:
	;
	v8780 = *(*int32)(unsafe.Add(mBase, uint32(v8775)))
	v8783 = int32(-2139062144)
	if (int32(16843008)-v8780|v8780)&v8783 != v8783 {
		v8804 = v8775
		v8805 = v8776
		v8806 = v8777
		goto L2107
	} else {
		goto L2125
	}
L2124:
	;
	v8797 = v8791
	v8798 = v8789
	v8799 = v8793
	goto L2108
L2125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8776))) = v8780
	v8788 = int32(4)
	v8789 = v8776 + v8788
	v8791 = v8775 + v8788
	v8793 = v8777 - v8788
	if base.Ui32(int32(3)) < base.Ui32(v8793) {
		v8775 = v8791
		v8776 = v8789
		v8777 = v8793
		goto L2123
	} else {
		goto L2126
	}
L2126:
	;
	goto L2124
L2127:
	;
	v8804 = v8797
	v8805 = v8798
	v8806 = v8799
	goto L2107
L2128:
	;
	v8813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8809))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8810))) = uint8(v8813)
	if v8813 == int32(0) {
		v8824 = v8809
		v8825 = v8810
		goto L2106
	} else {
		goto L2130
	}
L2129:
	;
	v8824 = v8820
	v8825 = v8818
	goto L2106
L2130:
	;
	v8817 = int32(1)
	v8818 = v8810 + v8817
	v8820 = v8809 + v8817
	v8822 = v8811 - v8817
	if v8822 != 0 {
		v8809 = v8820
		v8810 = v8818
		v8811 = v8822
		goto L2128
	} else {
		goto L2131
	}
L2131:
	;
	goto L2129
L2132:
	;
	F_fmgr_info(m, v8677, v8714)
	mBase = m.M
	v8841 = m.ExcPending
	if v8841 != 0 {
		goto L1
	} else {
		goto L2133
	}
L2133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+236)) = v8677
	v8844 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122]))
	if v8844 != int32(3) {
		goto L2134
	} else {
		goto L2135
	}
L2134:
	;
	v8866 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+268))
	v8868 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	v8870 = F_object_aclcheck(m, int32(2615), v8866, v8868, int64(256))
	mBase = m.M
	v8871 = m.ExcPending
	if v8871 != 0 {
		goto L1
	} else {
		goto L2140
	}
L2135:
	;
	v8849 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8850 = m.ExcPending
	if v8850 != 0 {
		goto L1
	} else {
		goto L2136
	}
L2136:
	;
	if v8849 == int32(0) {
		goto L2134
	} else {
		goto L2137
	}
L2137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+180)) = v8677
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+176)) = v8716
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_203), v8659+int32(176))
	mBase = m.M
	v8859 = m.ExcPending
	if v8859 != 0 {
		goto L1
	} else {
		goto L2138
	}
L2138:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(234), int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v8864 = m.ExcPending
	if v8864 != 0 {
		goto L1
	} else {
		goto L2139
	}
L2139:
	;
	goto L2134
L2140:
	;
	if v8870 != 0 {
		goto L2141
	} else {
		goto L2142
	}
L2141:
	;
	v8873 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+268))
	v8874 = F_get_namespace_name(m, v8873)
	mBase = m.M
	v8875 = m.ExcPending
	if v8875 != 0 {
		goto L1
	} else {
		goto L2144
	}
L2142:
	;
	goto L2143
L2143:
	;
	v8879 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	if v8879 != 0 {
		goto L2146
	} else {
		goto L2147
	}
L2144:
	;
	F_aclcheck_error(m, v8870, int32(36), v8874)
	mBase = m.M
	v8877 = m.ExcPending
	if v8877 != 0 {
		goto L1
	} else {
		goto L2145
	}
L2145:
	;
	goto L2143
L2146:
	;
	v8880 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+268))
	v8882 = F_RunNamespaceSearchHook(m, v8880, int32(1))
	mBase = m.M
	v8883 = m.ExcPending
	if v8883 != 0 {
		goto L1
	} else {
		goto L2149
	}
L2147:
	;
	goto L2148
L2148:
	;
	v8886 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	v8888 = F_object_aclcheck(m, int32(1255), v8677, v8886, int64(128))
	mBase = m.M
	v8889 = m.ExcPending
	if v8889 != 0 {
		goto L1
	} else {
		goto L2150
	}
L2149:
	;
	goto L2148
L2150:
	;
	if v8888 != 0 {
		goto L2151
	} else {
		goto L2152
	}
L2151:
	;
	v8891 = F_get_func_name(m, v8677)
	mBase = m.M
	v8892 = m.ExcPending
	if v8892 != 0 {
		goto L1
	} else {
		goto L2154
	}
L2152:
	;
	goto L2153
L2153:
	;
	v8897 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	if v8897 != 0 {
		goto L2156
	} else {
		goto L2157
	}
L2154:
	;
	F_aclcheck_error(m, v8888, int32(19), v8891)
	mBase = m.M
	v8894 = m.ExcPending
	if v8894 != 0 {
		goto L1
	} else {
		goto L2155
	}
L2155:
	;
	goto L2153
L2156:
	;
	F_RunFunctionExecuteHook(m, v8677)
	mBase = m.M
	v8899 = m.ExcPending
	if v8899 != 0 {
		goto L1
	} else {
		goto L2159
	}
L2157:
	;
	goto L2158
L2158:
	;
	v8900 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8659)+744)) = v8900
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+740)) = v8714
	*(*int64)(unsafe.Add(mBase, uint32(v8659)+749)) = v8900
	v8906 = F_pq_getmsgint(m, v8656, int32(2))
	mBase = m.M
	v8907 = m.ExcPending
	if v8907 != 0 {
		goto L1
	} else {
		goto L2160
	}
L2159:
	;
	goto L2158
L2160:
	;
	if int32(0) < v8906 {
		goto L2161
	} else {
		goto L2162
	}
L2161:
	;
	v8912 = F_palloc(m, v8906<<(uint(int32(1))%32))
	mBase = m.M
	v8913 = m.ExcPending
	if v8913 != 0 {
		goto L1
	} else {
		goto L2164
	}
L2162:
	;
	v8961 = int32(0)
	goto L2163
L2163:
	;
	v8999 = F_pq_getmsgint(m, v8656, int32(2))
	mBase = m.M
	v9000 = m.ExcPending
	if v9000 != 0 {
		goto L1
	} else {
		goto L2169
	}
L2164:
	;
	v8923 = v1
	goto L2165
L2165:
	;
	v8955 = F_pq_getmsgint(m, v8656, int32(2))
	mBase = m.M
	v8956 = m.ExcPending
	if v8956 != 0 {
		goto L1
	} else {
		goto L2167
	}
L2166:
	;
	v8961 = v8912
	goto L2163
L2167:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8912+v8923<<(uint(int32(1))%32)))) = uint16(v8955)
	v8959 = v8923 + int32(1)
	if v8959 != v8906 {
		v8923 = v8959
		goto L2165
	} else {
		goto L2168
	}
L2168:
	;
	goto L2166
L2169:
	;
	if int32(100) < v8999 {
		goto L2082
	} else {
		goto L2170
	}
L2170:
	;
	v9003 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8659)+248)))
	if v8999 != v9003 {
		goto L2082
	} else {
		goto L2171
	}
L2171:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8659)+758)) = uint16(v8999)
	if base.B2i32(v8999 != v8906)&base.B2i32(int32(2) <= v8906) != 0 {
		goto L2081
	} else {
		goto L2172
	}
L2172:
	;
	F_initStringInfo(m, v8659+int32(192))
	mBase = m.M
	v9013 = m.ExcPending
	if v9013 != 0 {
		goto L1
	} else {
		goto L2173
	}
L2173:
	;
	if int32(0) < v8999 {
		goto L2174
	} else {
		goto L2175
	}
L2174:
	;
	v9017 = v8659 + int32(760)
	v9030 = int32(0)
	goto L2177
L2175:
	;
	goto L2176
L2176:
	;
	v9196 = F_pq_getmsgint(m, v8656, int32(2))
	mBase = m.M
	v9197 = m.ExcPending
	if v9197 != 0 {
		goto L1
	} else {
		goto L2213
	}
L2177:
	;
	v9059 = v9030 << (uint(int32(3)) % 32)
	v9062 = v9059 + (v8659 + int32(740))
	v9064 = F_pq_getmsgint(m, v8656, int32(4))
	mBase = m.M
	v9065 = m.ExcPending
	if v9065 != 0 {
		goto L1
	} else {
		goto L2180
	}
L2178:
	;
	goto L2176
L2179:
	;
	if base.B2i32(v8906 < int32(2)) == int32(0) {
		goto L2192
	} else {
		goto L2193
	}
L2180:
	;
	if v9064 == int32(-1) {
		goto L2181
	} else {
		goto L2182
	}
L2181:
	;
	v9068 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9062)+24)) = uint8(v9068)
	goto L2179
L2182:
	;
	goto L2183
L2183:
	;
	v9070 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9062)+24)) = uint8(v9070)
	if v9064 < v9070 {
		goto L2080
	} else {
		goto L2184
	}
L2184:
	;
	v9075 = v8659 + int32(192)
	v9076 = *(*int32)(unsafe.Add(mBase, uint32(v9075)))
	v9077 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9076))) = uint8(v9077)
	*(*int32)(unsafe.Add(mBase, uint32(v9075)+12)) = v9077
	*(*int32)(unsafe.Add(mBase, uint32(v9075)+4)) = v9077
	goto L2185
L2185:
	;
	v9083 = F_pq_getmsgbytes(m, v8656, v9064)
	mBase = m.M
	v9084 = m.ExcPending
	if v9084 != 0 {
		goto L1
	} else {
		goto L2186
	}
L2186:
	;
	F_appendBinaryStringInfo(m, v9075, v9083, v9064)
	mBase = m.M
	v9086 = m.ExcPending
	if v9086 != 0 {
		goto L1
	} else {
		goto L2187
	}
L2187:
	;
	goto L2179
L2188:
	;
	v9156 = v9030 + int32(1)
	if v9156 != v8999 {
		v9030 = v9156
		goto L2177
	} else {
		goto L2212
	}
L2189:
	;
	v9131 = *(*int32)(unsafe.Add(mBase, uint32(v8706+v9030<<(uint(int32(2))%32))))
	F_getTypeBinaryInputInfo(m, v9131, v8659+int32(1564), v8659+int32(1560))
	mBase = m.M
	v9137 = m.ExcPending
	if v9137 != 0 {
		goto L1
	} else {
		goto L2205
	}
L2190:
	;
	v9101 = *(*int32)(unsafe.Add(mBase, uint32(v8706+v9030<<(uint(int32(2))%32))))
	F_getTypeInputInfo(m, v9101, v8659+int32(1564), v8659+int32(1560))
	mBase = m.M
	v9107 = m.ExcPending
	if v9107 != 0 {
		goto L1
	} else {
		goto L2196
	}
L2191:
	;
	v9096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9095))))
	switch v9096 {
	case 0:
		goto L2190
	case 1:
		goto L2189
	default:
		goto L2078
	}
L2192:
	;
	v9095 = v8961 + v9030<<(uint(int32(1))%32)
	goto L2191
L2193:
	;
	goto L2194
L2194:
	;
	if v8906 <= int32(0) {
		goto L2190
	} else {
		goto L2195
	}
L2195:
	;
	v9095 = v8961
	goto L2191
L2196:
	;
	if v9064 == int32(-1) {
		goto L2197
	} else {
		goto L2198
	}
L2197:
	;
	v9114 = int32(0)
	goto L2199
L2198:
	;
	v9111 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+192))
	v9112 = F_pg_client_to_server(m, v9111, v9064)
	mBase = m.M
	v9113 = m.ExcPending
	if v9113 != 0 {
		goto L1
	} else {
		goto L2200
	}
L2199:
	;
	v9116 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1564))
	v9117 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1560))
	v9119 = F_OidInputFunctionCall(m, v9116, v9114, v9117, int32(-1))
	mBase = m.M
	v9120 = m.ExcPending
	if v9120 != 0 {
		goto L1
	} else {
		goto L2201
	}
L2200:
	;
	v9114 = v9112
	goto L2199
L2201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9017+v9059))) = v9119
	if v9114 == int32(0) {
		goto L2188
	} else {
		goto L2202
	}
L2202:
	;
	v9124 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+192))
	if v9114 == v9124 {
		goto L2188
	} else {
		goto L2203
	}
L2203:
	;
	F_pfree(m, v9114)
	mBase = m.M
	v9127 = m.ExcPending
	if v9127 != 0 {
		goto L1
	} else {
		goto L2204
	}
L2204:
	;
	goto L2188
L2205:
	;
	v9139 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1564))
	v9144 = base.B2i32(v9064 == int32(-1))
	if v9064 == int32(-1) {
		goto L2206
	} else {
		goto L2207
	}
L2206:
	;
	v9145 = int32(0)
	goto L2208
L2207:
	;
	v9145 = v8659 + int32(192)
	goto L2208
L2208:
	;
	v9146 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1560))
	v9148 = F_OidReceiveFunctionCall(m, v9139, v9145, v9146, int32(-1))
	mBase = m.M
	v9149 = m.ExcPending
	if v9149 != 0 {
		goto L1
	} else {
		goto L2209
	}
L2209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9017+v9059))) = v9148
	if v9064 == int32(-1) {
		goto L2188
	} else {
		goto L2210
	}
L2210:
	;
	v9151 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+204))
	v9152 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+196))
	if v9151 != v9152 {
		goto L2079
	} else {
		goto L2211
	}
L2211:
	;
	goto L2188
L2212:
	;
	goto L2178
L2213:
	;
	F_pq_getmsgend(m, v8656)
	mBase = m.M
	v9199 = m.ExcPending
	if v9199 != 0 {
		goto L1
	} else {
		goto L2214
	}
L2214:
	;
	v9200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8659)+250)))
	if v9200 != int32(1) {
		goto L2077
	} else {
		goto L2215
	}
L2215:
	;
	v9203 = int32(0)
	if v8999 <= v9203 {
		goto L2077
	} else {
		goto L2216
	}
L2216:
	;
	v9216 = v9203
	goto L2217
L2217:
	;
	v9249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8659+int32(740)+v9216<<(uint(int32(3))%32))+24)))
	if v9249 == int32(0) {
		goto L2219
	} else {
		goto L2220
	}
L2218:
	;
	v9255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8659)+756)) = uint8(v9255)
	v9504 = int32(0)
	goto L2076
L2219:
	;
	v9253 = v9216 + int32(1)
	if base.I32_extend16_s(v8999) != v9253 {
		v9216 = v9253
		goto L2217
	} else {
		goto L2222
	}
L2220:
	;
	goto L2221
L2221:
	;
	goto L2218
L2222:
	;
	goto L2077
L2223:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v9264 = m.ExcPending
	if v9264 != 0 {
		goto L1
	} else {
		goto L2224
	}
L2224:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v9268 = m.ExcPending
	if v9268 != 0 {
		goto L1
	} else {
		goto L2225
	}
L2225:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(209), int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v9273 = m.ExcPending
	if v9273 != 0 {
		goto L1
	} else {
		goto L2226
	}
L2226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2227:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v9280 = m.ExcPending
	if v9280 != 0 {
		goto L1
	} else {
		goto L2228
	}
L2228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659))) = v8677
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_206), v8659)
	mBase = m.M
	v9284 = m.ExcPending
	if v9284 != 0 {
		goto L1
	} else {
		goto L2229
	}
L2229:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(141), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9289 = m.ExcPending
	if v9289 != 0 {
		goto L1
	} else {
		goto L2230
	}
L2230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2231:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9296 = m.ExcPending
	if v9296 != 0 {
		goto L1
	} else {
		goto L2232
	}
L2232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+16)) = v8691 + int32(4)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_208), v8659+int32(16))
	mBase = m.M
	v9304 = m.ExcPending
	if v9304 != 0 {
		goto L1
	} else {
		goto L2233
	}
L2233:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(149), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9309 = m.ExcPending
	if v9309 != 0 {
		goto L1
	} else {
		goto L2234
	}
L2234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+32)) = v8691 + int32(4)
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_209), v8659+int32(32))
	mBase = m.M
	v9323 = m.ExcPending
	if v9323 != 0 {
		goto L1
	} else {
		goto L2236
	}
L2236:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(154), int32(_a_F_PostgresMainLoopOnce_207))
	mBase = m.M
	v9328 = m.ExcPending
	if v9328 != 0 {
		goto L1
	} else {
		goto L2237
	}
L2237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2238:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9335 = m.ExcPending
	if v9335 != 0 {
		goto L1
	} else {
		goto L2239
	}
L2239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+48)) = v8999
	v9337 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8659)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+52)) = v9337
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_210), v8659+int32(48))
	mBase = m.M
	v9343 = m.ExcPending
	if v9343 != 0 {
		goto L1
	} else {
		goto L2240
	}
L2240:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(353), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9348 = m.ExcPending
	if v9348 != 0 {
		goto L1
	} else {
		goto L2241
	}
L2241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2242:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9355 = m.ExcPending
	if v9355 != 0 {
		goto L1
	} else {
		goto L2243
	}
L2243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+164)) = v8999
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+160)) = v8906
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_212), v8659+int32(160))
	mBase = m.M
	v9362 = m.ExcPending
	if v9362 != 0 {
		goto L1
	} else {
		goto L2244
	}
L2244:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(361), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9367 = m.ExcPending
	if v9367 != 0 {
		goto L1
	} else {
		goto L2245
	}
L2245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2246:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9374 = m.ExcPending
	if v9374 != 0 {
		goto L1
	} else {
		goto L2247
	}
L2247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+144)) = v9064
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_213), v8659+int32(144))
	mBase = m.M
	v9380 = m.ExcPending
	if v9380 != 0 {
		goto L1
	} else {
		goto L2248
	}
L2248:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(385), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9385 = m.ExcPending
	if v9385 != 0 {
		goto L1
	} else {
		goto L2249
	}
L2249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2250:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v9392 = m.ExcPending
	if v9392 != 0 {
		goto L1
	} else {
		goto L2251
	}
L2251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+128)) = v9030 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_214), v8659+int32(128))
	mBase = m.M
	v9400 = m.ExcPending
	if v9400 != 0 {
		goto L1
	} else {
		goto L2252
	}
L2252:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(448), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9405 = m.ExcPending
	if v9405 != 0 {
		goto L1
	} else {
		goto L2253
	}
L2253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2254:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9412 = m.ExcPending
	if v9412 != 0 {
		goto L1
	} else {
		goto L2255
	}
L2255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+64)) = base.I32_extend16_s(v9096)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v8659-int32(-64))
	mBase = m.M
	v9419 = m.ExcPending
	if v9419 != 0 {
		goto L1
	} else {
		goto L2256
	}
L2256:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(453), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9424 = m.ExcPending
	if v9424 != 0 {
		goto L1
	} else {
		goto L2257
	}
L2257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2258:
	;
	v9504 = v9465
	goto L2076
L2259:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9508 = m.ExcPending
	if v9508 != 0 {
		goto L1
	} else {
		goto L2262
	}
L2260:
	;
	goto L2261
L2261:
	;
	v9509 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+272))
	v9510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8659)+756)))
	v9512 = v8659 + int32(192)
	F_pq_beginmessage(m, v9512, int32(86))
	mBase = m.M
	v9515 = m.ExcPending
	if v9515 != 0 {
		goto L1
	} else {
		goto L2263
	}
L2262:
	;
	goto L2261
L2263:
	;
	if v9510 == int32(1) {
		goto L2265
	} else {
		goto L2266
	}
L2264:
	;
	v9667 = v8659 + int32(192)
	F_pq_endmessage(m, v9667)
	mBase = m.M
	v9669 = m.ExcPending
	if v9669 != 0 {
		goto L1
	} else {
		goto L2297
	}
L2265:
	;
	F_enlargeStringInfo(m, v9512, int32(4))
	mBase = m.M
	v9520 = m.ExcPending
	if v9520 != 0 {
		goto L1
	} else {
		goto L2268
	}
L2266:
	;
	goto L2267
L2267:
	;
	switch v9196 & int32(_a_F_PostgresMainLoopOnce_216) {
	case 0:
		goto L2269
	case 1:
		goto L2271
	default:
		goto L2270
	}
L2268:
	;
	v9521 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+196))
	v9522 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v9521+v9522))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+196)) = v9521 + int32(4)
	goto L2264
L2269:
	;
	F_getTypeOutputInfo(m, v9509, v8659+int32(1564), v8659+int32(1560))
	mBase = m.M
	v9651 = m.ExcPending
	if v9651 != 0 {
		goto L1
	} else {
		goto L2293
	}
L2270:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9630 = m.ExcPending
	if v9630 != 0 {
		goto L1
	} else {
		goto L2289
	}
L2271:
	;
	F_getTypeBinaryOutputInfo(m, v9509, v8659+int32(1564), v8659+int32(1560))
	mBase = m.M
	v9536 = m.ExcPending
	if v9536 != 0 {
		goto L1
	} else {
		goto L2272
	}
L2272:
	;
	v9537 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1564))
	v9538 = m.G0
	v9540 = v9538 + int32(-64)
	m.G0 = v9540
	v9543 = v9538 + int32(-56)
	v9545 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_fmgr_info_cxt_security(m, v9537, v9543, v9545, int32(0))
	mBase = m.M
	v9548 = m.ExcPending
	if v9548 != 0 {
		goto L1
	} else {
		goto L2274
	}
L2273:
	;
	v9590 = *(*int32)(unsafe.Add(mBase, uint32(v9572)))
	v9592 = v8659 + int32(192)
	F_enlargeStringInfo(m, v9592, int32(4))
	mBase = m.M
	v9595 = m.ExcPending
	if v9595 != 0 {
		goto L1
	} else {
		goto L2286
	}
L2274:
	;
	v9549 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9540)+40)) = v9549
	*(*int64)(unsafe.Add(mBase, uint32(v9540)+45)) = v9549
	v9553 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9540)+60)) = uint8(v9553)
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+56)) = v9504
	*(*int32)(unsafe.Add(mBase, uint32(v9540)+36)) = v9543
	v9557 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v9540)+54)) = uint16(v9557)
	v9561 = *(*int32)(unsafe.Add(mBase, uint32(v9540)+8))
	v9562 = m.T0[v9561].(func(*base.Module, int32) int32)(m, v9538+int32(-28))
	mBase = m.M
	v9563 = m.ExcPending
	if v9563 != 0 {
		goto L1
	} else {
		goto L2275
	}
L2275:
	;
	v9564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9540)+52)))
	if v9564 != int32(1) {
		goto L2276
	} else {
		goto L2277
	}
L2276:
	;
	v9567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9562))))
	if v9567&int32(3) != 0 {
		goto L2279
	} else {
		goto L2280
	}
L2277:
	;
	goto L2278
L2278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9579 = m.ExcPending
	if v9579 != 0 {
		goto L1
	} else {
		goto L2283
	}
L2279:
	;
	v9570 = F_detoast_attr(m, v9562)
	mBase = m.M
	v9571 = m.ExcPending
	if v9571 != 0 {
		goto L1
	} else {
		goto L2282
	}
L2280:
	;
	v9572 = v9562
	goto L2281
L2281:
	;
	m.G0 = v9540 - int32(-64)
	goto L2273
L2282:
	;
	v9572 = v9570
	goto L2281
L2283:
	;
	v9580 = *(*int32)(unsafe.Add(mBase, uint32(v9540)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9540))) = v9580
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_217), v9540)
	mBase = m.M
	v9584 = m.ExcPending
	if v9584 != 0 {
		goto L1
	} else {
		goto L2284
	}
L2284:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_218), int32(1143), int32(_a_F_PostgresMainLoopOnce_219))
	mBase = m.M
	v9589 = m.ExcPending
	if v9589 != 0 {
		goto L1
	} else {
		goto L2285
	}
L2285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2286:
	;
	v9596 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+196))
	v9597 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+192))
	v9599 = int32(2)
	v9601 = int32(4)
	v9602 = int32(base.Ui32(v9590)>>(uint(v9599)%32)) - v9601
	v9603 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v9596+v9597))) = base.I32_rotr(v9602&v9603, int32(8)) | base.I32_rotr(v9602, int32(24))&v9603
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+196)) = v9596 + v9601
	v9618 = *(*int32)(unsafe.Add(mBase, uint32(v9572)))
	F_appendBinaryStringInfo(m, v9592, v9572+v9601, int32(base.Ui32(v9618)>>(uint(v9599)%32))-v9601)
	mBase = m.M
	v9624 = m.ExcPending
	if v9624 != 0 {
		goto L1
	} else {
		goto L2287
	}
L2287:
	;
	F_pfree(m, v9572)
	mBase = m.M
	v9626 = m.ExcPending
	if v9626 != 0 {
		goto L1
	} else {
		goto L2288
	}
L2288:
	;
	goto L2264
L2289:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9633 = m.ExcPending
	if v9633 != 0 {
		goto L1
	} else {
		goto L2290
	}
L2290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+112)) = base.I32_extend16_s(v9196)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v8659+int32(112))
	mBase = m.M
	v9640 = m.ExcPending
	if v9640 != 0 {
		goto L1
	} else {
		goto L2291
	}
L2291:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), int32(106), int32(_a_F_PostgresMainLoopOnce_220))
	mBase = m.M
	v9645 = m.ExcPending
	if v9645 != 0 {
		goto L1
	} else {
		goto L2292
	}
L2292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2293:
	;
	v9654 = *(*int32)(unsafe.Add(mBase, uint32(v8659)+1564))
	v9655 = F_OidOutputFunctionCall(m, v9654, v9504)
	mBase = m.M
	v9656 = m.ExcPending
	if v9656 != 0 {
		goto L1
	} else {
		goto L2294
	}
L2294:
	;
	v9657 = F_strlen(m, v9655)
	mBase = m.M
	F_pq_sendcountedtext(m, v8659+int32(192), v9655, v9657)
	mBase = m.M
	v9659 = m.ExcPending
	if v9659 != 0 {
		goto L1
	} else {
		goto L2295
	}
L2295:
	;
	F_pfree(m, v9655)
	mBase = m.M
	v9661 = m.ExcPending
	if v9661 != 0 {
		goto L1
	} else {
		goto L2296
	}
L2296:
	;
	goto L2264
L2297:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v9671 = m.ExcPending
	if v9671 != 0 {
		goto L1
	} else {
		goto L2298
	}
L2298:
	;
	v9675 = F_check_log_duration(m, v9667, base.B2i32(v8844 == int32(3)))
	mBase = m.M
	v9676 = m.ExcPending
	if v9676 != 0 {
		goto L1
	} else {
		goto L2303
	}
L2299:
	;
	m.G0 = v8659 + int32(1568)
	v9721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L2311
L2300:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_204), v9711, int32(_a_F_PostgresMainLoopOnce_205))
	mBase = m.M
	v9714 = m.ExcPending
	if v9714 != 0 {
		goto L1
	} else {
		goto L2310
	}
L2301:
	;
	v9696 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9697 = m.ExcPending
	if v9697 != 0 {
		goto L1
	} else {
		goto L2307
	}
L2302:
	;
	v9681 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9682 = m.ExcPending
	if v9682 != 0 {
		goto L1
	} else {
		goto L2304
	}
L2303:
	;
	switch v9675 - int32(1) {
	case 0:
		goto L2302
	case 1:
		goto L2301
	default:
		goto L2299
	}
L2304:
	;
	if v9681 == int32(0) {
		goto L2299
	} else {
		goto L2305
	}
L2305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+80)) = v8659 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_180), v8659+int32(80))
	mBase = m.M
	v9692 = m.ExcPending
	if v9692 != 0 {
		goto L1
	} else {
		goto L2306
	}
L2306:
	;
	v9711 = int32(312)
	goto L2300
L2307:
	;
	if v9696 == int32(0) {
		goto L2299
	} else {
		goto L2308
	}
L2308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+104)) = v8677
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+100)) = v8716
	*(*int32)(unsafe.Add(mBase, uint32(v8659)+96)) = v8659 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_221), v8659+int32(96))
	mBase = m.M
	v9709 = m.ExcPending
	if v9709 != 0 {
		goto L1
	} else {
		goto L2309
	}
L2309:
	;
	v9711 = int32(317)
	goto L2300
L2310:
	;
	goto L2299
L2311:
	;
	if v9721 != 0 {
		goto L2312
	} else {
		goto L2313
	}
L2312:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v9724 = m.ExcPending
	if v9724 != 0 {
		goto L1
	} else {
		goto L2315
	}
L2313:
	;
	goto L2314
L2314:
	;
	v9726 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v9726 != 0 {
		goto L2316
	} else {
		goto L2317
	}
L2315:
	;
	goto L2314
L2316:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v9728 = m.ExcPending
	if v9728 != 0 {
		goto L1
	} else {
		goto L2319
	}
L2317:
	;
	goto L2318
L2318:
	;
	v9733 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v9733)
	goto L211
L2319:
	;
	v9730 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v9730)
	goto L2318
L2320:
	;
	v9740 = v40 + int32(440)
	v9741 = F_pq_getmsgbyte(m, v9740)
	mBase = m.M
	v9742 = m.ExcPending
	if v9742 != 0 {
		goto L1
	} else {
		goto L2321
	}
L2321:
	;
	v9743 = F_pq_getmsgstring(m, v9740)
	mBase = m.M
	v9744 = m.ExcPending
	if v9744 != 0 {
		goto L1
	} else {
		goto L2322
	}
L2322:
	;
	F_pq_getmsgend(m, v9740)
	mBase = m.M
	v9746 = m.ExcPending
	if v9746 != 0 {
		goto L1
	} else {
		goto L2323
	}
L2323:
	;
	switch v9741 - int32(80) {
	case 0:
		goto L2325
	default:
		goto L199
	case 3:
		goto L2326
	}
L2324:
	;
	v9771 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9771 != int32(2) {
		goto L211
	} else {
		goto L2336
	}
L2325:
	;
	v9762 = F_GetPortalByName(m, v9743)
	mBase = m.M
	v9763 = m.ExcPending
	if v9763 != 0 {
		goto L1
	} else {
		goto L2333
	}
L2326:
	;
	v9749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9743))))
	if v9749 != 0 {
		goto L2327
	} else {
		goto L2328
	}
L2327:
	;
	F_DropPreparedStatement(m, v9743, int32(0))
	mBase = m.M
	v9752 = m.ExcPending
	if v9752 != 0 {
		goto L1
	} else {
		goto L2330
	}
L2328:
	;
	goto L2329
L2329:
	;
	v9754 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120]))
	if v9754 == int32(0) {
		goto L2324
	} else {
		goto L2331
	}
L2330:
	;
	goto L2324
L2331:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = int32(0)
	F_DropCachedPlan(m, v9754)
	mBase = m.M
	v9761 = m.ExcPending
	if v9761 != 0 {
		goto L1
	} else {
		goto L2332
	}
L2332:
	;
	goto L2324
L2333:
	;
	if v9762 == int32(0) {
		goto L2324
	} else {
		goto L2334
	}
L2334:
	;
	F_PortalDrop(m, v9762, int32(0))
	mBase = m.M
	v9768 = m.ExcPending
	if v9768 != 0 {
		goto L1
	} else {
		goto L2335
	}
L2335:
	;
	goto L2324
L2336:
	;
	F_pq_putemptymessage(m, int32(51))
	mBase = m.M
	v9776 = m.ExcPending
	if v9776 != 0 {
		goto L1
	} else {
		goto L2337
	}
L2337:
	;
	goto L211
L2338:
	;
	v9782 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v9782 < int32(0) {
		goto L2340
	} else {
		goto L2341
	}
L2339:
	;
	v9789 = v40 + int32(440)
	v9790 = F_pq_getmsgbyte(m, v9789)
	mBase = m.M
	v9791 = m.ExcPending
	if v9791 != 0 {
		goto L1
	} else {
		goto L2343
	}
L2340:
	;
	v9786 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v9786
	goto L2342
L2341:
	;
	goto L2342
L2342:
	;
	goto L2339
L2343:
	;
	v9792 = F_pq_getmsgstring(m, v9789)
	mBase = m.M
	v9793 = m.ExcPending
	if v9793 != 0 {
		goto L1
	} else {
		goto L2344
	}
L2344:
	;
	F_pq_getmsgend(m, v9789)
	mBase = m.M
	v9795 = m.ExcPending
	if v9795 != 0 {
		goto L1
	} else {
		goto L2345
	}
L2345:
	;
	switch v9790 - int32(80) {
	case 0:
		goto L2347
	default:
		goto L2346
	case 3:
		goto L2348
	}
L2346:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10014 = m.ExcPending
	if v10014 != 0 {
		goto L1
	} else {
		goto L2390
	}
L2347:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9978 = m.ExcPending
	if v9978 != 0 {
		goto L1
	} else {
		goto L2375
	}
L2348:
	;
	F_start_xact_command(m)
	mBase = m.M
	v9799 = m.ExcPending
	if v9799 != 0 {
		goto L1
	} else {
		goto L2349
	}
L2349:
	;
	v9802 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v9802
	v9804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9792))))
	if v9804 != 0 {
		goto L2351
	} else {
		goto L2352
	}
L2350:
	;
	v9815 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v9816 = *(*int32)(unsafe.Add(mBase, uint32(v9815)+24))
	goto L2356
L2351:
	;
	v9806 = F_FetchPreparedStatement(m, v9792, int32(1))
	mBase = m.M
	v9807 = m.ExcPending
	if v9807 != 0 {
		goto L1
	} else {
		goto L2354
	}
L2352:
	;
	goto L2353
L2353:
	;
	v9810 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120]))
	if v9810 == int32(0) {
		goto L198
	} else {
		goto L2355
	}
L2354:
	;
	v9808 = *(*int32)(unsafe.Add(mBase, uint32(v9806)+64))
	v9813 = v9808
	goto L2350
L2355:
	;
	v9813 = v9810
	goto L2350
L2356:
	;
	if (v9816-int32(7))&int32(-9) == int32(0) {
		goto L2357
	} else {
		goto L2358
	}
L2357:
	;
	v9823 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+52))
	if v9823 != 0 {
		goto L197
	} else {
		goto L2360
	}
L2358:
	;
	goto L2359
L2359:
	;
	v9825 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9825 != int32(2) {
		goto L211
	} else {
		goto L2361
	}
L2360:
	;
	goto L2359
L2361:
	;
	v9828 = int32(_a_F_PostgresMainLoopOnce_222)
	F_resetStringInfo(m, v9828)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[135])) = int32(116)
	goto L2362
L2362:
	;
	v9832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9813)+24)))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_222), int32(2))
	mBase = m.M
	v9836 = m.ExcPending
	if v9836 != 0 {
		goto L1
	} else {
		goto L2363
	}
L2363:
	;
	v9838 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[136]))
	v9839 = int32(_a_F_PostgresMainLoopOnce_223)
	v9840 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137]))
	v9842 = int32(8)
	v9846 = v9832<<(uint(v9842)%32) | int32(base.Ui32(v9832)>>(uint(v9842)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v9838+v9840))) = uint16(v9846)
	v9850 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137])) = v9850 + int32(2)
	v9854 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+24))
	if int32(0) < v9854 {
		goto L2364
	} else {
		goto L2365
	}
L2364:
	;
	v9866 = int32(0)
	goto L2367
L2365:
	;
	goto L2366
L2366:
	;
	F_pq_endmessage_reuse(m, int32(_a_F_PostgresMainLoopOnce_222))
	mBase = m.M
	v9966 = m.ExcPending
	if v9966 != 0 {
		goto L1
	} else {
		goto L2371
	}
L2367:
	;
	v9895 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+20))
	v9899 = *(*int32)(unsafe.Add(mBase, uint32(v9895+v9866<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_222), int32(4))
	mBase = m.M
	v9903 = m.ExcPending
	if v9903 != 0 {
		goto L1
	} else {
		goto L2369
	}
L2368:
	;
	goto L2366
L2369:
	;
	v9904 = int32(_a_F_PostgresMainLoopOnce_223)
	v9905 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137]))
	v9907 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[136]))
	v9911 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v9905+v9907))) = base.I32_rotr(v9899, int32(24))&v9911 | base.I32_rotr(v9899&v9911, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137])) = v9905 + int32(4)
	v9924 = v9866 + int32(1)
	v9925 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+24))
	if v9924 < v9925 {
		v9866 = v9924
		goto L2367
	} else {
		goto L2370
	}
L2370:
	;
	goto L2368
L2371:
	;
	v9967 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+52))
	if v9967 == int32(0) {
		goto L216
	} else {
		goto L2372
	}
L2372:
	;
	v9970 = F_CachedPlanGetTargetList(m, v9813)
	mBase = m.M
	v9971 = m.ExcPending
	if v9971 != 0 {
		goto L1
	} else {
		goto L2373
	}
L2373:
	;
	v9973 = *(*int32)(unsafe.Add(mBase, uint32(v9813)+52))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_222), v9973, v9970, int32(0))
	mBase = m.M
	v9976 = m.ExcPending
	if v9976 != 0 {
		goto L1
	} else {
		goto L2374
	}
L2374:
	;
	goto L211
L2375:
	;
	v9981 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v9981
	v9983 = F_GetPortalByName(m, v9792)
	mBase = m.M
	v9984 = m.ExcPending
	if v9984 != 0 {
		goto L1
	} else {
		goto L2376
	}
L2376:
	;
	if v9983 == int32(0) {
		goto L196
	} else {
		goto L2377
	}
L2377:
	;
	v9988 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v9989 = *(*int32)(unsafe.Add(mBase, uint32(v9988)+24))
	goto L2378
L2378:
	;
	if (v9989-int32(7))&int32(-9) == int32(0) {
		goto L2379
	} else {
		goto L2380
	}
L2379:
	;
	v9996 = *(*int32)(unsafe.Add(mBase, uint32(v9983)+92))
	if v9996 != 0 {
		goto L195
	} else {
		goto L2382
	}
L2380:
	;
	goto L2381
L2381:
	;
	v9998 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v9998 != int32(2) {
		goto L211
	} else {
		goto L2383
	}
L2382:
	;
	goto L2381
L2383:
	;
	v10001 = *(*int32)(unsafe.Add(mBase, uint32(v9983)+92))
	if v10001 != 0 {
		goto L2384
	} else {
		goto L2385
	}
L2384:
	;
	v10003 = F_FetchPortalTargetList(m, v9983)
	mBase = m.M
	v10004 = m.ExcPending
	if v10004 != 0 {
		goto L1
	} else {
		goto L2387
	}
L2385:
	;
	goto L2386
L2386:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v10010 = m.ExcPending
	if v10010 != 0 {
		goto L1
	} else {
		goto L2389
	}
L2387:
	;
	v10005 = *(*int32)(unsafe.Add(mBase, uint32(v9983)+96))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_222), v10001, v10003, v10005)
	mBase = m.M
	v10007 = m.ExcPending
	if v10007 != 0 {
		goto L1
	} else {
		goto L2388
	}
L2388:
	;
	goto L211
L2389:
	;
	goto L211
L2390:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10017 = m.ExcPending
	if v10017 != 0 {
		goto L1
	} else {
		goto L2391
	}
L2391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+368)) = v9790
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_224), v40+int32(368))
	mBase = m.M
	v10023 = m.ExcPending
	if v10023 != 0 {
		goto L1
	} else {
		goto L2392
	}
L2392:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_225), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10028 = m.ExcPending
	if v10028 != 0 {
		goto L1
	} else {
		goto L2393
	}
L2393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2394:
	;
	v10034 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10034 != int32(2) {
		goto L211
	} else {
		goto L2395
	}
L2395:
	;
	v10038 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106]))
	v10039 = *(*int32)(unsafe.Add(mBase, uint32(v10038)+4))
	v10040 = m.T0[v10039].(func(*base.Module) int32)(m)
	mBase = m.M
	v10041 = m.ExcPending
	if v10041 != 0 {
		goto L1
	} else {
		goto L2396
	}
L2396:
	;
	goto L211
L2397:
	;
	v10048 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v10049 = *(*int32)(unsafe.Add(mBase, uint32(v10048)+24))
	if v10049 == int32(4) {
		goto L2399
	} else {
		goto L2400
	}
L2398:
	;
	v10057 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123])))
	goto L2402
L2399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10048)+24)) = int32(1)
	goto L2401
L2400:
	;
	goto L2401
L2401:
	;
	goto L2398
L2402:
	;
	if v10057 != 0 {
		goto L2403
	} else {
		goto L2404
	}
L2403:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v10060 = m.ExcPending
	if v10060 != 0 {
		goto L1
	} else {
		goto L2406
	}
L2404:
	;
	goto L2405
L2405:
	;
	v10062 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])))
	if v10062 != 0 {
		goto L2407
	} else {
		goto L2408
	}
L2406:
	;
	goto L2405
L2407:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v10064 = m.ExcPending
	if v10064 != 0 {
		goto L1
	} else {
		goto L2410
	}
L2408:
	;
	goto L2409
L2409:
	;
	v10069 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v10069)
	goto L211
L2410:
	;
	v10066 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[128])) = uint8(v10066)
	goto L2409
L2411:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11])) = int32(0)
	goto L2413
L2412:
	;
	goto L2413
L2413:
	;
	v10082 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138]))
	if v10082 != 0 {
		goto L194
	} else {
		goto L2414
	}
L2414:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v10085 = m.ExcPending
	if v10085 != 0 {
		goto L1
	} else {
		goto L2415
	}
L2415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2416:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10092 = m.ExcPending
	if v10092 != 0 {
		goto L1
	} else {
		goto L2417
	}
L2417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v635
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_11), v40+int32(16))
	mBase = m.M
	v10098 = m.ExcPending
	if v10098 != 0 {
		goto L1
	} else {
		goto L2418
	}
L2418:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_226), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10103 = m.ExcPending
	if v10103 != 0 {
		goto L1
	} else {
		goto L2419
	}
L2419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2420:
	;
	goto L211
L2421:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v10153 = m.ExcPending
	if v10153 != 0 {
		goto L1
	} else {
		goto L2422
	}
L2422:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_227), int32(0))
	mBase = m.M
	v10157 = m.ExcPending
	if v10157 != 0 {
		goto L1
	} else {
		goto L2423
	}
L2423:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1581), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v10162 = m.ExcPending
	if v10162 != 0 {
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
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10169 = m.ExcPending
	if v10169 != 0 {
		goto L1
	} else {
		goto L2426
	}
L2426:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10173 = m.ExcPending
	if v10173 != 0 {
		goto L1
	} else {
		goto L2427
	}
L2427:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10175 = m.ExcPending
	if v10175 != 0 {
		goto L1
	} else {
		goto L2428
	}
L2428:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1603), int32(_a_F_PostgresMainLoopOnce_185))
	mBase = m.M
	v10180 = m.ExcPending
	if v10180 != 0 {
		goto L1
	} else {
		goto L2429
	}
L2429:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2430:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v10187 = m.ExcPending
	if v10187 != 0 {
		goto L1
	} else {
		goto L2431
	}
L2431:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_228), int32(0))
	mBase = m.M
	v10191 = m.ExcPending
	if v10191 != 0 {
		goto L1
	} else {
		goto L2432
	}
L2432:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1778), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10196 = m.ExcPending
	if v10196 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10203 = m.ExcPending
	if v10203 != 0 {
		goto L1
	} else {
		goto L2435
	}
L2435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+196)) = v7123
	*(*int32)(unsafe.Add(mBase, uint32(v40)+192)) = v7025
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_229), v40+int32(192))
	mBase = m.M
	v10210 = m.ExcPending
	if v10210 != 0 {
		goto L1
	} else {
		goto L2436
	}
L2436:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1831), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10215 = m.ExcPending
	if v10215 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10222 = m.ExcPending
	if v10222 != 0 {
		goto L1
	} else {
		goto L2439
	}
L2439:
	;
	v10223 = *(*int32)(unsafe.Add(mBase, uint32(v6864)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v10223
	*(*int32)(unsafe.Add(mBase, uint32(v40)+180)) = v6830
	*(*int32)(unsafe.Add(mBase, uint32(v40)+176)) = v7123
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_230), v40+int32(176))
	mBase = m.M
	v10231 = m.ExcPending
	if v10231 != 0 {
		goto L1
	} else {
		goto L2440
	}
L2440:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1837), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10236 = m.ExcPending
	if v10236 != 0 {
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
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10244 = m.ExcPending
	if v10244 != 0 {
		goto L1
	} else {
		goto L2443
	}
L2443:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10248 = m.ExcPending
	if v10248 != 0 {
		goto L1
	} else {
		goto L2444
	}
L2444:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10250 = m.ExcPending
	if v10250 != 0 {
		goto L1
	} else {
		goto L2445
	}
L2445:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1855), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10255 = m.ExcPending
	if v10255 != 0 {
		goto L1
	} else {
		goto L2446
	}
L2446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2447:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v10262 = m.ExcPending
	if v10262 != 0 {
		goto L1
	} else {
		goto L2448
	}
L2448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+160)) = v7238 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), v40+int32(160))
	mBase = m.M
	v10270 = m.ExcPending
	if v10270 != 0 {
		goto L1
	} else {
		goto L2449
	}
L2449:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2051), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10275 = m.ExcPending
	if v10275 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v10282 = m.ExcPending
	if v10282 != 0 {
		goto L1
	} else {
		goto L2452
	}
L2452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = base.I32_extend16_s(v7305)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_215), v40+int32(80))
	mBase = m.M
	v10289 = m.ExcPending
	if v10289 != 0 {
		goto L1
	} else {
		goto L2453
	}
L2453:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2058), int32(_a_F_PostgresMainLoopOnce_190))
	mBase = m.M
	v10294 = m.ExcPending
	if v10294 != 0 {
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
	F_errcode(m, int32(259))
	mBase = m.M
	v10301 = m.ExcPending
	if v10301 != 0 {
		goto L1
	} else {
		goto L2456
	}
L2456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+224)) = v7889
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_232), v40+int32(224))
	mBase = m.M
	v10307 = m.ExcPending
	if v10307 != 0 {
		goto L1
	} else {
		goto L2457
	}
L2457:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2244), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v10312 = m.ExcPending
	if v10312 != 0 {
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
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10320 = m.ExcPending
	if v10320 != 0 {
		goto L1
	} else {
		goto L2460
	}
L2460:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10324 = m.ExcPending
	if v10324 != 0 {
		goto L1
	} else {
		goto L2461
	}
L2461:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10326 = m.ExcPending
	if v10326 != 0 {
		goto L1
	} else {
		goto L2462
	}
L2462:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2360), int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v10331 = m.ExcPending
	if v10331 != 0 {
		goto L1
	} else {
		goto L2463
	}
L2463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2464:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10338 = m.ExcPending
	if v10338 != 0 {
		goto L1
	} else {
		goto L2465
	}
L2465:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_233), int32(0))
	mBase = m.M
	v10342 = m.ExcPending
	if v10342 != 0 {
		goto L1
	} else {
		goto L2466
	}
L2466:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_234), int32(_a_F_PostgresMainLoopOnce_235))
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10354 = m.ExcPending
	if v10354 != 0 {
		goto L1
	} else {
		goto L2469
	}
L2469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+352)) = v9741
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_236), v40+int32(352))
	mBase = m.M
	v10360 = m.ExcPending
	if v10360 != 0 {
		goto L1
	} else {
		goto L2470
	}
L2470:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_237), int32(_a_F_PostgresMainLoopOnce_4))
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
	F_errcode(m, int32(386))
	mBase = m.M
	v10372 = m.ExcPending
	if v10372 != 0 {
		goto L1
	} else {
		goto L2473
	}
L2473:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_228), int32(0))
	mBase = m.M
	v10376 = m.ExcPending
	if v10376 != 0 {
		goto L1
	} else {
		goto L2474
	}
L2474:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2776), int32(_a_F_PostgresMainLoopOnce_238))
	mBase = m.M
	v10381 = m.ExcPending
	if v10381 != 0 {
		goto L1
	} else {
		goto L2475
	}
L2475:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2476:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10388 = m.ExcPending
	if v10388 != 0 {
		goto L1
	} else {
		goto L2477
	}
L2477:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10392 = m.ExcPending
	if v10392 != 0 {
		goto L1
	} else {
		goto L2478
	}
L2478:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10394 = m.ExcPending
	if v10394 != 0 {
		goto L1
	} else {
		goto L2479
	}
L2479:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2797), int32(_a_F_PostgresMainLoopOnce_238))
	mBase = m.M
	v10399 = m.ExcPending
	if v10399 != 0 {
		goto L1
	} else {
		goto L2480
	}
L2480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2481:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v10406 = m.ExcPending
	if v10406 != 0 {
		goto L1
	} else {
		goto L2482
	}
L2482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+384)) = v9792
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_232), v40+int32(384))
	mBase = m.M
	v10412 = m.ExcPending
	if v10412 != 0 {
		goto L1
	} else {
		goto L2483
	}
L2483:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2858), int32(_a_F_PostgresMainLoopOnce_239))
	mBase = m.M
	v10417 = m.ExcPending
	if v10417 != 0 {
		goto L1
	} else {
		goto L2484
	}
L2484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2485:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10424 = m.ExcPending
	if v10424 != 0 {
		goto L1
	} else {
		goto L2486
	}
L2486:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), int32(0))
	mBase = m.M
	v10428 = m.ExcPending
	if v10428 != 0 {
		goto L1
	} else {
		goto L2487
	}
L2487:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10430 = m.ExcPending
	if v10430 != 0 {
		goto L1
	} else {
		goto L2488
	}
L2488:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2874), int32(_a_F_PostgresMainLoopOnce_239))
	mBase = m.M
	v10435 = m.ExcPending
	if v10435 != 0 {
		goto L1
	} else {
		goto L2489
	}
L2489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2490:
	;
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2491:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10449 = m.ExcPending
	if v10449 != 0 {
		goto L1
	} else {
		goto L2492
	}
L2492:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_240), int32(0))
	mBase = m.M
	v10453 = m.ExcPending
	if v10453 != 0 {
		goto L1
	} else {
		goto L2493
	}
L2493:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_241), int32(_a_F_PostgresMainLoopOnce_235))
	mBase = m.M
	v10458 = m.ExcPending
	if v10458 != 0 {
		goto L1
	} else {
		goto L2494
	}
L2494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
