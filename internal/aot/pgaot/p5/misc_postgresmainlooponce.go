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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v669 int64
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
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
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int64
	_ = v808
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v921 int32
	_ = v921
	var v938 int32
	_ = v938
	var v952 int32
	_ = v952
	var v961 int32
	_ = v961
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1230 int32
	_ = v1230
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int64
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1352 int32
	_ = v1352
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int64
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1611 int64
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1933 int32
	_ = v1933
	var v1934 int64
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1962 int64
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1966 int64
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1971 int64
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int64
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int64
	_ = v1981
	var v1986 int32
	_ = v1986
	var v1987 int64
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int64
	_ = v1996
	var v2000 int64
	_ = v2000
	var v2003 int64
	_ = v2003
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2064 int64
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2151 int64
	_ = v2151
	var v2153 int64
	_ = v2153
	var v2160 int32
	_ = v2160
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2180 int64
	_ = v2180
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2197 int64
	_ = v2197
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2215 int64
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int64
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2386 int32
	_ = v2386
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2444 int32
	_ = v2444
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2480 int32
	_ = v2480
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2644 int32
	_ = v2644
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2739 int32
	_ = v2739
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2782 int32
	_ = v2782
	var v2788 int32
	_ = v2788
	var v2791 int32
	_ = v2791
	var v2795 int32
	_ = v2795
	var v2800 int32
	_ = v2800
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2822 int32
	_ = v2822
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2836 int32
	_ = v2836
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2845 int32
	_ = v2845
	var v2849 int32
	_ = v2849
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int32
	_ = v2857
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2886 int32
	_ = v2886
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2903 int32
	_ = v2903
	var v2908 int32
	_ = v2908
	var v2915 int32
	_ = v2915
	var v2918 int32
	_ = v2918
	var v2923 int32
	_ = v2923
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2941 int32
	_ = v2941
	var v2945 int32
	_ = v2945
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v2957 int32
	_ = v2957
	var v2961 int32
	_ = v2961
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2983 int32
	_ = v2983
	var v2987 int32
	_ = v2987
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3005 int32
	_ = v3005
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3015 int32
	_ = v3015
	var v3020 int64
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3024 int64
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3029 int64
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3032 int64
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int64
	_ = v3039
	var v3044 int32
	_ = v3044
	var v3045 int64
	_ = v3045
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3054 int64
	_ = v3054
	var v3058 int64
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3062 int32
	_ = v3062
	var v3065 int32
	_ = v3065
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3074 int64
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3080 int64
	_ = v3080
	var v3085 int32
	_ = v3085
	var v3086 int64
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3091 int64
	_ = v3091
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3102 int64
	_ = v3102
	var v3108 int32
	_ = v3108
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3121 int32
	_ = v3121
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3139 int32
	_ = v3139
	var v3143 int32
	_ = v3143
	var v3148 int32
	_ = v3148
	var v3152 int32
	_ = v3152
	var v3155 int32
	_ = v3155
	var v3159 int32
	_ = v3159
	var v3164 int32
	_ = v3164
	var v3169 int64
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3179 int32
	_ = v3179
	var v3184 int64
	_ = v3184
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3194 int32
	_ = v3194
	var v3203 int32
	_ = v3203
	var v3213 int32
	_ = v3213
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int64
	_ = v3249
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3266 int32
	_ = v3266
	var v3268 int64
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3275 int32
	_ = v3275
	var v3281 int32
	_ = v3281
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3292 int32
	_ = v3292
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3309 int32
	_ = v3309
	var v3311 int32
	_ = v3311
	var v3313 int32
	_ = v3313
	var v3317 int64
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3322 int64
	_ = v3322
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3342 int32
	_ = v3342
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3352 int64
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3366 int32
	_ = v3366
	var v3368 int32
	_ = v3368
	var v3374 int32
	_ = v3374
	var v3381 int32
	_ = v3381
	var v3382 int64
	_ = v3382
	var v3384 int64
	_ = v3384
	var v3385 int64
	_ = v3385
	var v3389 int64
	_ = v3389
	var v3393 int32
	_ = v3393
	var v3398 int32
	_ = v3398
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3408 int32
	_ = v3408
	var v3410 int32
	_ = v3410
	var v3415 int32
	_ = v3415
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3438 int32
	_ = v3438
	var v3442 int32
	_ = v3442
	var v3443 int64
	_ = v3443
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3461 int32
	_ = v3461
	var v3464 int32
	_ = v3464
	var v3465 int32
	_ = v3465
	var v3468 int32
	_ = v3468
	var v3477 int32
	_ = v3477
	var v3487 int32
	_ = v3487
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3507 int32
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3517 int32
	_ = v3517
	var v3519 int32
	_ = v3519
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3527 int32
	_ = v3527
	var v3528 int64
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3534 int64
	_ = v3534
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3542 int32
	_ = v3542
	var v3549 int32
	_ = v3549
	var v3551 int32
	_ = v3551
	var v3552 int64
	_ = v3552
	var v3554 int32
	_ = v3554
	var v3562 int32
	_ = v3562
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3569 int32
	_ = v3569
	var v3571 int32
	_ = v3571
	var v3576 int32
	_ = v3576
	var v3578 int32
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3603 int32
	_ = v3603
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3619 int32
	_ = v3619
	var v3624 int32
	_ = v3624
	var v3625 int32
	_ = v3625
	var v3633 int32
	_ = v3633
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3646 int32
	_ = v3646
	var v3648 int32
	_ = v3648
	var v3653 int32
	_ = v3653
	var v3658 int32
	_ = v3658
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3669 int32
	_ = v3669
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3682 int32
	_ = v3682
	var v3702 int32
	_ = v3702
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3710 int64
	_ = v3710
	var v3712 int64
	_ = v3712
	var v3715 int64
	_ = v3715
	var v3717 int64
	_ = v3717
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3731 int32
	_ = v3731
	var v3733 int32
	_ = v3733
	var v3780 int64
	_ = v3780
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3793 int32
	_ = v3793
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3807 int32
	_ = v3807
	var v3809 int64
	_ = v3809
	var v3848 int32
	_ = v3848
	var v3849 int32
	_ = v3849
	var v3853 int32
	_ = v3853
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3869 int32
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3900 int32
	_ = v3900
	var v3901 int32
	_ = v3901
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3931 int32
	_ = v3931
	var v3938 int32
	_ = v3938
	var v3939 int32
	_ = v3939
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3952 int32
	_ = v3952
	var v3953 int32
	_ = v3953
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3964 int32
	_ = v3964
	var v3968 int32
	_ = v3968
	var v3970 int32
	_ = v3970
	var v3972 int32
	_ = v3972
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3978 int32
	_ = v3978
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3985 int32
	_ = v3985
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4018 int32
	_ = v4018
	var v4020 int32
	_ = v4020
	var v4021 int32
	_ = v4021
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4037 int32
	_ = v4037
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4046 int32
	_ = v4046
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4059 int32
	_ = v4059
	var v4067 int32
	_ = v4067
	var v4069 int32
	_ = v4069
	var v4070 int32
	_ = v4070
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4109 int32
	_ = v4109
	var v4111 int32
	_ = v4111
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4133 int32
	_ = v4133
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4144 int32
	_ = v4144
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4149 int32
	_ = v4149
	var v4153 int32
	_ = v4153
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4175 int32
	_ = v4175
	var v4182 int32
	_ = v4182
	var v4185 int32
	_ = v4185
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4194 int32
	_ = v4194
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4201 int32
	_ = v4201
	var v4203 int32
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4212 int32
	_ = v4212
	var v4218 int32
	_ = v4218
	var v4219 int32
	_ = v4219
	var v4221 int32
	_ = v4221
	var v4227 int32
	_ = v4227
	var v4234 int32
	_ = v4234
	var v4238 int32
	_ = v4238
	var v4241 int32
	_ = v4241
	var v4247 int32
	_ = v4247
	var v4254 int32
	_ = v4254
	var v4258 int32
	_ = v4258
	var v4261 int32
	_ = v4261
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4280 int32
	_ = v4280
	var v4282 int32
	_ = v4282
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4295 int32
	_ = v4295
	var v4297 int32
	_ = v4297
	var v4299 int32
	_ = v4299
	var v4304 int32
	_ = v4304
	var v4320 int32
	_ = v4320
	var v4326 int32
	_ = v4326
	var v4329 int32
	_ = v4329
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4339 int32
	_ = v4339
	var v4344 int32
	_ = v4344
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4351 int32
	_ = v4351
	var v4352 int32
	_ = v4352
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4358 int32
	_ = v4358
	var v4361 int32
	_ = v4361
	var v4362 int32
	_ = v4362
	var v4364 int32
	_ = v4364
	var v4366 int32
	_ = v4366
	var v4368 int32
	_ = v4368
	var v4372 int32
	_ = v4372
	var v4374 int32
	_ = v4374
	var v4376 int32
	_ = v4376
	var v4380 int32
	_ = v4380
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4390 int32
	_ = v4390
	var v4397 int32
	_ = v4397
	var v4416 int32
	_ = v4416
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4425 int32
	_ = v4425
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4444 int32
	_ = v4444
	var v4448 int32
	_ = v4448
	var v4451 int32
	_ = v4451
	var v4455 int32
	_ = v4455
	var v4460 int32
	_ = v4460
	var v4464 int32
	_ = v4464
	var v4467 int32
	_ = v4467
	var v4473 int32
	_ = v4473
	var v4478 int32
	_ = v4478
	var v4482 int32
	_ = v4482
	var v4485 int32
	_ = v4485
	var v4489 int32
	_ = v4489
	var v4494 int32
	_ = v4494
	var v4498 int32
	_ = v4498
	var v4501 int32
	_ = v4501
	var v4508 int32
	_ = v4508
	var v4513 int32
	_ = v4513
	var v4517 int32
	_ = v4517
	var v4520 int32
	_ = v4520
	var v4524 int32
	_ = v4524
	var v4529 int32
	_ = v4529
	var v4533 int32
	_ = v4533
	var v4536 int32
	_ = v4536
	var v4540 int32
	_ = v4540
	var v4545 int32
	_ = v4545
	var v4549 int32
	_ = v4549
	var v4552 int32
	_ = v4552
	var v4556 int32
	_ = v4556
	var v4561 int32
	_ = v4561
	var v4565 int32
	_ = v4565
	var v4568 int32
	_ = v4568
	var v4572 int32
	_ = v4572
	var v4577 int32
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4588 int32
	_ = v4588
	var v4593 int32
	_ = v4593
	var v4597 int32
	_ = v4597
	var v4600 int32
	_ = v4600
	var v4604 int32
	_ = v4604
	var v4609 int32
	_ = v4609
	var v4613 int32
	_ = v4613
	var v4616 int32
	_ = v4616
	var v4620 int32
	_ = v4620
	var v4625 int32
	_ = v4625
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4636 int32
	_ = v4636
	var v4641 int32
	_ = v4641
	var v4644 int32
	_ = v4644
	var v4648 int32
	_ = v4648
	var v4650 int32
	_ = v4650
	var v4658 int32
	_ = v4658
	var v4663 int32
	_ = v4663
	var v4667 int32
	_ = v4667
	var v4669 int32
	_ = v4669
	var v4677 int32
	_ = v4677
	var v4682 int32
	_ = v4682
	var v4686 int32
	_ = v4686
	var v4688 int32
	_ = v4688
	var v4696 int32
	_ = v4696
	var v4701 int32
	_ = v4701
	var v4705 int32
	_ = v4705
	var v4707 int32
	_ = v4707
	var v4715 int32
	_ = v4715
	var v4720 int32
	_ = v4720
	var v4724 int32
	_ = v4724
	var v4727 int32
	_ = v4727
	var v4738 int32
	_ = v4738
	var v4743 int32
	_ = v4743
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4757 int32
	_ = v4757
	var v4762 int32
	_ = v4762
	var v4766 int32
	_ = v4766
	var v4769 int32
	_ = v4769
	var v4773 int32
	_ = v4773
	var v4778 int32
	_ = v4778
	var v4785 int32
	_ = v4785
	var v4788 int32
	_ = v4788
	var v4792 int32
	_ = v4792
	var v4797 int32
	_ = v4797
	var v4801 int32
	_ = v4801
	var v4804 int32
	_ = v4804
	var v4810 int32
	_ = v4810
	var v4815 int32
	_ = v4815
	var v4817 int32
	_ = v4817
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4825 int32
	_ = v4825
	var v4827 int32
	_ = v4827
	var v4843 int32
	_ = v4843
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4856 int32
	_ = v4856
	var v4864 int32
	_ = v4864
	var v4867 int32
	_ = v4867
	var v4868 int32
	_ = v4868
	var v4872 int32
	_ = v4872
	var v4873 int32
	_ = v4873
	var v4874 int32
	_ = v4874
	var v4878 int32
	_ = v4878
	var v4884 int32
	_ = v4884
	var v4886 int32
	_ = v4886
	var v4887 int32
	_ = v4887
	var v4889 int32
	_ = v4889
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4893 int32
	_ = v4893
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4898 int32
	_ = v4898
	var v4899 int32
	_ = v4899
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4909 int32
	_ = v4909
	var v4910 int32
	_ = v4910
	var v4918 int32
	_ = v4918
	var v4925 int32
	_ = v4925
	var v4930 int32
	_ = v4930
	var v4932 int32
	_ = v4932
	var v4933 int32
	_ = v4933
	var v4939 int32
	_ = v4939
	var v4943 int32
	_ = v4943
	var v4946 int32
	_ = v4946
	var v4948 int32
	_ = v4948
	var v4952 int32
	_ = v4952
	var v4953 int32
	_ = v4953
	var v4956 int32
	_ = v4956
	var v4957 int32
	_ = v4957
	var v4970 int32
	_ = v4970
	var v4971 int32
	_ = v4971
	var v4976 int32
	_ = v4976
	var v4977 int32
	_ = v4977
	var v4978 int32
	_ = v4978
	var v4980 int32
	_ = v4980
	var v4983 int32
	_ = v4983
	var v4984 int32
	_ = v4984
	var v4990 int32
	_ = v4990
	var v4992 int32
	_ = v4992
	var v4996 int32
	_ = v4996
	var v4999 int32
	_ = v4999
	var v5001 int32
	_ = v5001
	var v5006 int32
	_ = v5006
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
	var v5014 int32
	_ = v5014
	var v5020 int32
	_ = v5020
	var v5025 int32
	_ = v5025
	var v5033 int32
	_ = v5033
	var v5037 int32
	_ = v5037
	var v5042 int32
	_ = v5042
	var v5046 int32
	_ = v5046
	var v5050 int32
	_ = v5050
	var v5055 int32
	_ = v5055
	var v5056 int32
	_ = v5056
	var v5057 int32
	_ = v5057
	var v5058 int32
	_ = v5058
	var v5060 int32
	_ = v5060
	var v5062 int32
	_ = v5062
	var v5065 int32
	_ = v5065
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5070 int32
	_ = v5070
	var v5074 int32
	_ = v5074
	var v5076 int32
	_ = v5076
	var v5077 int64
	_ = v5077
	var v5080 int64
	_ = v5080
	var v5088 int32
	_ = v5088
	var v5089 int32
	_ = v5089
	var v5091 int32
	_ = v5091
	var v5092 int32
	_ = v5092
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5100 int32
	_ = v5100
	var v5105 int32
	_ = v5105
	var v5110 int32
	_ = v5110
	var v5115 int32
	_ = v5115
	var v5117 int32
	_ = v5117
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5123 int32
	_ = v5123
	var v5124 int32
	_ = v5124
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5131 int32
	_ = v5131
	var v5132 int32
	_ = v5132
	var v5134 int32
	_ = v5134
	var v5136 int32
	_ = v5136
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5140 int32
	_ = v5140
	var v5147 int32
	_ = v5147
	var v5149 int32
	_ = v5149
	var v5151 int32
	_ = v5151
	var v5162 int32
	_ = v5162
	var v5190 int32
	_ = v5190
	var v5194 int32
	_ = v5194
	var v5196 int32
	_ = v5196
	var v5242 int32
	_ = v5242
	var v5249 int32
	_ = v5249
	var v5254 int32
	_ = v5254
	var v5258 int32
	_ = v5258
	var v5265 int32
	_ = v5265
	var v5270 int32
	_ = v5270
	var v5274 int32
	_ = v5274
	var v5281 int32
	_ = v5281
	var v5286 int32
	_ = v5286
	var v5290 int32
	_ = v5290
	var v5297 int32
	_ = v5297
	var v5302 int32
	_ = v5302
	var v5306 int32
	_ = v5306
	var v5313 int32
	_ = v5313
	var v5318 int32
	_ = v5318
	var v5355 int32
	_ = v5355
	var v5357 int32
	_ = v5357
	var v5362 int32
	_ = v5362
	var v5364 int32
	_ = v5364
	var v5369 int32
	_ = v5369
	var v5370 int32
	_ = v5370
	var v5374 int32
	_ = v5374
	var v5376 int32
	_ = v5376
	var v5379 int32
	_ = v5379
	var v5384 int32
	_ = v5384
	var v5397 int32
	_ = v5397
	var v5398 int64
	_ = v5398
	var v5399 int32
	_ = v5399
	var v5400 int32
	_ = v5400
	var v5404 int64
	_ = v5404
	var v5405 int32
	_ = v5405
	var v5412 int32
	_ = v5412
	var v5419 int32
	_ = v5419
	var v5421 int32
	_ = v5421
	var v5426 int32
	_ = v5426
	var v5427 int32
	_ = v5427
	var v5428 int32
	_ = v5428
	var v5431 int32
	_ = v5431
	var v5434 int32
	_ = v5434
	var v5437 int32
	_ = v5437
	var v5438 int32
	_ = v5438
	var v5442 int32
	_ = v5442
	var v5444 int32
	_ = v5444
	var v5447 int32
	_ = v5447
	var v5452 int32
	_ = v5452
	var v5465 int32
	_ = v5465
	var v5466 int64
	_ = v5466
	var v5467 int32
	_ = v5467
	var v5468 int32
	_ = v5468
	var v5472 int64
	_ = v5472
	var v5473 int32
	_ = v5473
	var v5480 int32
	_ = v5480
	var v5487 int32
	_ = v5487
	var v5488 int32
	_ = v5488
	var v5490 int32
	_ = v5490
	var v5495 int32
	_ = v5495
	var v5497 int32
	_ = v5497
	var v5500 int32
	_ = v5500
	var v5501 int32
	_ = v5501
	var v5504 int32
	_ = v5504
	var v5540 int32
	_ = v5540
	var v5544 int32
	_ = v5544
	var v5545 int32
	_ = v5545
	var v5546 int32
	_ = v5546
	var v5548 int32
	_ = v5548
	var v5551 int32
	_ = v5551
	var v5552 int32
	_ = v5552
	var v5570 int32
	_ = v5570
	var v5628 int32
	_ = v5628
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5640 int32
	_ = v5640
	var v5642 int32
	_ = v5642
	var v5645 int32
	_ = v5645
	var v5649 int32
	_ = v5649
	var v5656 int32
	_ = v5656
	var v5685 int32
	_ = v5685
	var v5689 int32
	_ = v5689
	var v5690 int32
	_ = v5690
	var v5691 int32
	_ = v5691
	var v5694 int32
	_ = v5694
	var v5696 int32
	_ = v5696
	var v5697 int32
	_ = v5697
	var v5698 int32
	_ = v5698
	var v5699 int32
	_ = v5699
	var v5701 int32
	_ = v5701
	var v5705 int32
	_ = v5705
	var v5706 int32
	_ = v5706
	var v5712 int32
	_ = v5712
	var v5753 int32
	_ = v5753
	var v5777 int32
	_ = v5777
	var v5794 int32
	_ = v5794
	var v5811 int32
	_ = v5811
	var v5818 int32
	_ = v5818
	var v5831 int32
	_ = v5831
	var v5832 int32
	_ = v5832
	var v5849 int32
	_ = v5849
	var v5871 int32
	_ = v5871
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5880 int32
	_ = v5880
	var v5884 int32
	_ = v5884
	var v5889 int64
	_ = v5889
	var v5893 int32
	_ = v5893
	var v5895 int32
	_ = v5895
	var v5896 int32
	_ = v5896
	var v5899 int32
	_ = v5899
	var v5910 int32
	_ = v5910
	var v5918 int32
	_ = v5918
	var v5922 int32
	_ = v5922
	var v5927 int64
	_ = v5927
	var v5931 int32
	_ = v5931
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5937 int32
	_ = v5937
	var v5948 int32
	_ = v5948
	var v5952 int32
	_ = v5952
	var v5953 int32
	_ = v5953
	var v5954 int32
	_ = v5954
	var v5961 int32
	_ = v5961
	var v5968 int32
	_ = v5968
	var v5969 int32
	_ = v5969
	var v5978 int32
	_ = v5978
	var v5981 int32
	_ = v5981
	var v5984 int32
	_ = v5984
	var v5994 int32
	_ = v5994
	var v5997 int32
	_ = v5997
	var v6001 int32
	_ = v6001
	var v6003 int32
	_ = v6003
	var v6008 int32
	_ = v6008
	var v6011 int32
	_ = v6011
	var v6013 int32
	_ = v6013
	var v6018 int32
	_ = v6018
	var v6019 int32
	_ = v6019
	var v6025 int32
	_ = v6025
	var v6027 int32
	_ = v6027
	var v6030 int32
	_ = v6030
	var v6031 int32
	_ = v6031
	var v6035 int32
	_ = v6035
	var v6036 int32
	_ = v6036
	var v6037 int32
	_ = v6037
	var v6039 int32
	_ = v6039
	var v6040 int32
	_ = v6040
	var v6042 int32
	_ = v6042
	var v6044 int32
	_ = v6044
	var v6047 int32
	_ = v6047
	var v6048 int32
	_ = v6048
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
	var v6066 int32
	_ = v6066
	var v6069 int32
	_ = v6069
	var v6070 int32
	_ = v6070
	var v6074 int32
	_ = v6074
	var v6076 int32
	_ = v6076
	var v6079 int32
	_ = v6079
	var v6084 int32
	_ = v6084
	var v6097 int32
	_ = v6097
	var v6098 int64
	_ = v6098
	var v6099 int32
	_ = v6099
	var v6100 int32
	_ = v6100
	var v6104 int64
	_ = v6104
	var v6105 int32
	_ = v6105
	var v6112 int32
	_ = v6112
	var v6118 int32
	_ = v6118
	var v6121 int32
	_ = v6121
	var v6122 int32
	_ = v6122
	var v6124 int32
	_ = v6124
	var v6129 int32
	_ = v6129
	var v6130 int32
	_ = v6130
	var v6131 int32
	_ = v6131
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6137 int32
	_ = v6137
	var v6139 int32
	_ = v6139
	var v6141 int32
	_ = v6141
	var v6143 int32
	_ = v6143
	var v6145 int32
	_ = v6145
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6161 int32
	_ = v6161
	var v6165 int32
	_ = v6165
	var v6166 int32
	_ = v6166
	var v6168 int32
	_ = v6168
	var v6169 int32
	_ = v6169
	var v6172 int32
	_ = v6172
	var v6173 int32
	_ = v6173
	var v6174 int32
	_ = v6174
	var v6175 int32
	_ = v6175
	var v6178 int32
	_ = v6178
	var v6183 int32
	_ = v6183
	var v6190 int32
	_ = v6190
	var v6191 int32
	_ = v6191
	var v6192 int32
	_ = v6192
	var v6202 int32
	_ = v6202
	var v6203 int32
	_ = v6203
	var v6204 int32
	_ = v6204
	var v6206 int32
	_ = v6206
	var v6209 int32
	_ = v6209
	var v6210 int32
	_ = v6210
	var v6211 int32
	_ = v6211
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6231 int32
	_ = v6231
	var v6234 int32
	_ = v6234
	var v6236 int32
	_ = v6236
	var v6239 int32
	_ = v6239
	var v6240 int32
	_ = v6240
	var v6248 int32
	_ = v6248
	var v6251 int32
	_ = v6251
	var v6253 int32
	_ = v6253
	var v6255 int32
	_ = v6255
	var v6261 int32
	_ = v6261
	var v6266 int32
	_ = v6266
	var v6268 int32
	_ = v6268
	var v6270 int32
	_ = v6270
	var v6275 int32
	_ = v6275
	var v6277 int32
	_ = v6277
	var v6279 int32
	_ = v6279
	var v6280 int32
	_ = v6280
	var v6306 int32
	_ = v6306
	var v6343 int32
	_ = v6343
	var v6356 int32
	_ = v6356
	var v6362 int32
	_ = v6362
	var v6365 int32
	_ = v6365
	var v6367 int32
	_ = v6367
	var v6369 int32
	_ = v6369
	var v6371 int32
	_ = v6371
	var v6374 int32
	_ = v6374
	var v6377 int32
	_ = v6377
	var v6378 int32
	_ = v6378
	var v6383 int32
	_ = v6383
	var v6384 int32
	_ = v6384
	var v6392 int32
	_ = v6392
	var v6394 int32
	_ = v6394
	var v6398 int32
	_ = v6398
	var v6399 int32
	_ = v6399
	var v6410 int32
	_ = v6410
	var v6412 int32
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6414 int32
	_ = v6414
	var v6418 int32
	_ = v6418
	var v6425 int32
	_ = v6425
	var v6454 int32
	_ = v6454
	var v6458 int32
	_ = v6458
	var v6459 int32
	_ = v6459
	var v6460 int32
	_ = v6460
	var v6463 int32
	_ = v6463
	var v6465 int32
	_ = v6465
	var v6466 int32
	_ = v6466
	var v6467 int32
	_ = v6467
	var v6468 int32
	_ = v6468
	var v6470 int32
	_ = v6470
	var v6472 int32
	_ = v6472
	var v6473 int32
	_ = v6473
	var v6479 int32
	_ = v6479
	var v6491 int32
	_ = v6491
	var v6519 int32
	_ = v6519
	var v6558 int32
	_ = v6558
	var v6602 int32
	_ = v6602
	var v6606 int32
	_ = v6606
	var v6610 int32
	_ = v6610
	var v6614 int64
	_ = v6614
	var v6618 int32
	_ = v6618
	var v6619 int32
	_ = v6619
	var v6622 int32
	_ = v6622
	var v6623 int32
	_ = v6623
	var v6627 int32
	_ = v6627
	var v6628 int32
	_ = v6628
	var v6634 int32
	_ = v6634
	var v6635 int32
	_ = v6635
	var v6638 int32
	_ = v6638
	var v6678 int32
	_ = v6678
	var v6679 int32
	_ = v6679
	var v6682 int32
	_ = v6682
	var v6687 int32
	_ = v6687
	var v6723 int32
	_ = v6723
	var v6729 int32
	_ = v6729
	var v6734 int32
	_ = v6734
	var v6735 int32
	_ = v6735
	var v6739 int32
	_ = v6739
	var v6741 int32
	_ = v6741
	var v6744 int32
	_ = v6744
	var v6749 int32
	_ = v6749
	var v6762 int32
	_ = v6762
	var v6763 int64
	_ = v6763
	var v6764 int32
	_ = v6764
	var v6765 int32
	_ = v6765
	var v6769 int64
	_ = v6769
	var v6770 int32
	_ = v6770
	var v6777 int32
	_ = v6777
	var v6785 int32
	_ = v6785
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
	var v6804 int32
	_ = v6804
	var v6805 int32
	_ = v6805
	var v6807 int32
	_ = v6807
	var v6810 int32
	_ = v6810
	var v6815 int32
	_ = v6815
	var v6817 int32
	_ = v6817
	var v6822 int32
	_ = v6822
	var v6823 int32
	_ = v6823
	var v6824 int32
	_ = v6824
	var v6825 int32
	_ = v6825
	var v6826 int32
	_ = v6826
	var v6827 int32
	_ = v6827
	var v6831 int32
	_ = v6831
	var v6834 int32
	_ = v6834
	var v6835 int32
	_ = v6835
	var v6839 int32
	_ = v6839
	var v6841 int32
	_ = v6841
	var v6844 int32
	_ = v6844
	var v6849 int32
	_ = v6849
	var v6862 int32
	_ = v6862
	var v6863 int64
	_ = v6863
	var v6864 int32
	_ = v6864
	var v6865 int32
	_ = v6865
	var v6869 int64
	_ = v6869
	var v6870 int32
	_ = v6870
	var v6877 int32
	_ = v6877
	var v6884 int32
	_ = v6884
	var v6885 int32
	_ = v6885
	var v6887 int32
	_ = v6887
	var v6892 int32
	_ = v6892
	var v6893 int32
	_ = v6893
	var v6896 int32
	_ = v6896
	var v6897 int32
	_ = v6897
	var v6899 int32
	_ = v6899
	var v6900 int32
	_ = v6900
	var v6907 int32
	_ = v6907
	var v6910 int32
	_ = v6910
	var v6913 int32
	_ = v6913
	var v6918 int32
	_ = v6918
	var v6919 int32
	_ = v6919
	var v6920 int32
	_ = v6920
	var v6921 int32
	_ = v6921
	var v6924 int32
	_ = v6924
	var v6925 int32
	_ = v6925
	var v6929 int32
	_ = v6929
	var v6936 int32
	_ = v6936
	var v6937 int32
	_ = v6937
	var v6938 int32
	_ = v6938
	var v6939 int32
	_ = v6939
	var v6941 int32
	_ = v6941
	var v6946 int32
	_ = v6946
	var v6947 int32
	_ = v6947
	var v6949 int32
	_ = v6949
	var v6950 int32
	_ = v6950
	var v6953 int32
	_ = v6953
	var v6954 int32
	_ = v6954
	var v6955 int32
	_ = v6955
	var v6956 int32
	_ = v6956
	var v6958 int32
	_ = v6958
	var v6960 int32
	_ = v6960
	var v6964 int32
	_ = v6964
	var v6968 int32
	_ = v6968
	var v6969 int32
	_ = v6969
	var v6974 int32
	_ = v6974
	var v6981 int32
	_ = v6981
	var v6994 int32
	_ = v6994
	var v6995 int32
	_ = v6995
	var v6996 int32
	_ = v6996
	var v7001 int32
	_ = v7001
	var v7003 int32
	_ = v7003
	var v7005 int32
	_ = v7005
	var v7008 int32
	_ = v7008
	var v7010 int32
	_ = v7010
	var v7016 int32
	_ = v7016
	var v7018 int32
	_ = v7018
	var v7023 int32
	_ = v7023
	var v7027 int32
	_ = v7027
	var v7028 int32
	_ = v7028
	var v7033 int32
	_ = v7033
	var v7034 int32
	_ = v7034
	var v7044 int32
	_ = v7044
	var v7048 int32
	_ = v7048
	var v7049 int32
	_ = v7049
	var v7052 int32
	_ = v7052
	var v7055 int32
	_ = v7055
	var v7064 int32
	_ = v7064
	var v7067 int32
	_ = v7067
	var v7069 int32
	_ = v7069
	var v7073 int32
	_ = v7073
	var v7077 int32
	_ = v7077
	var v7082 int32
	_ = v7082
	var v7086 int32
	_ = v7086
	var v7090 int64
	_ = v7090
	var v7093 int32
	_ = v7093
	var v7096 int32
	_ = v7096
	var v7097 int32
	_ = v7097
	var v7100 int32
	_ = v7100
	var v7101 int32
	_ = v7101
	var v7104 int32
	_ = v7104
	var v7105 int32
	_ = v7105
	var v7106 int32
	_ = v7106
	var v7108 int32
	_ = v7108
	var v7109 int32
	_ = v7109
	var v7112 int32
	_ = v7112
	var v7118 int32
	_ = v7118
	var v7123 int32
	_ = v7123
	var v7125 int32
	_ = v7125
	var v7127 int32
	_ = v7127
	var v7128 int32
	_ = v7128
	var v7129 int32
	_ = v7129
	var v7131 int32
	_ = v7131
	var v7134 int32
	_ = v7134
	var v7136 int32
	_ = v7136
	var v7140 int32
	_ = v7140
	var v7143 int32
	_ = v7143
	var v7146 int32
	_ = v7146
	var v7150 int32
	_ = v7150
	var v7187 int32
	_ = v7187
	var v7188 int64
	_ = v7188
	var v7192 int32
	_ = v7192
	var v7197 int32
	_ = v7197
	var v7201 int32
	_ = v7201
	var v7206 int64
	_ = v7206
	var v7210 int32
	_ = v7210
	var v7212 int32
	_ = v7212
	var v7213 int32
	_ = v7213
	var v7216 int32
	_ = v7216
	var v7227 int32
	_ = v7227
	var v7269 int32
	_ = v7269
	var v7270 int32
	_ = v7270
	var v7274 int32
	_ = v7274
	var v7276 int32
	_ = v7276
	var v7279 int32
	_ = v7279
	var v7284 int32
	_ = v7284
	var v7297 int32
	_ = v7297
	var v7298 int64
	_ = v7298
	var v7299 int32
	_ = v7299
	var v7300 int32
	_ = v7300
	var v7304 int64
	_ = v7304
	var v7305 int32
	_ = v7305
	var v7312 int32
	_ = v7312
	var v7319 int32
	_ = v7319
	var v7322 int32
	_ = v7322
	var v7327 int32
	_ = v7327
	var v7328 int32
	_ = v7328
	var v7334 int32
	_ = v7334
	var v7335 int32
	_ = v7335
	var v7338 int32
	_ = v7338
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7382 int32
	_ = v7382
	var v7390 int32
	_ = v7390
	var v7423 int32
	_ = v7423
	var v7424 int32
	_ = v7424
	var v7429 int32
	_ = v7429
	var v7432 int32
	_ = v7432
	var v7433 int32
	_ = v7433
	var v7440 int32
	_ = v7440
	var v7443 int32
	_ = v7443
	var v7446 int32
	_ = v7446
	var v7449 int32
	_ = v7449
	var v7456 int32
	_ = v7456
	var v7459 int32
	_ = v7459
	var v7461 int32
	_ = v7461
	var v7462 int32
	_ = v7462
	var v7463 int32
	_ = v7463
	var v7465 int32
	_ = v7465
	var v7466 int32
	_ = v7466
	var v7467 int32
	_ = v7467
	var v7468 int32
	_ = v7468
	var v7469 int32
	_ = v7469
	var v7471 int32
	_ = v7471
	var v7473 int32
	_ = v7473
	var v7474 int32
	_ = v7474
	var v7475 int32
	_ = v7475
	var v7476 int32
	_ = v7476
	var v7477 int32
	_ = v7477
	var v7478 int32
	_ = v7478
	var v7479 int32
	_ = v7479
	var v7482 int32
	_ = v7482
	var v7483 int32
	_ = v7483
	var v7489 int32
	_ = v7489
	var v7490 int32
	_ = v7490
	var v7494 int32
	_ = v7494
	var v7497 int32
	_ = v7497
	var v7498 int32
	_ = v7498
	var v7500 int32
	_ = v7500
	var v7501 int32
	_ = v7501
	var v7502 int32
	_ = v7502
	var v7504 int32
	_ = v7504
	var v7505 int32
	_ = v7505
	var v7509 int32
	_ = v7509
	var v7510 int32
	_ = v7510
	var v7523 int32
	_ = v7523
	var v7524 int32
	_ = v7524
	var v7527 int32
	_ = v7527
	var v7532 int32
	_ = v7532
	var v7547 int32
	_ = v7547
	var v7567 int32
	_ = v7567
	var v7568 int32
	_ = v7568
	var v7570 int32
	_ = v7570
	var v7577 int32
	_ = v7577
	var v7578 int32
	_ = v7578
	var v7580 int32
	_ = v7580
	var v7581 int32
	_ = v7581
	var v7585 int32
	_ = v7585
	var v7586 int32
	_ = v7586
	var v7587 int32
	_ = v7587
	var v7588 int32
	_ = v7588
	var v7589 int32
	_ = v7589
	var v7594 int32
	_ = v7594
	var v7595 int32
	_ = v7595
	var v7603 int32
	_ = v7603
	var v7604 int32
	_ = v7604
	var v7611 int32
	_ = v7611
	var v7613 int32
	_ = v7613
	var v7614 int32
	_ = v7614
	var v7615 int32
	_ = v7615
	var v7616 int32
	_ = v7616
	var v7618 int32
	_ = v7618
	var v7619 int32
	_ = v7619
	var v7621 int32
	_ = v7621
	var v7622 int32
	_ = v7622
	var v7623 int32
	_ = v7623
	var v7628 int32
	_ = v7628
	var v7629 int32
	_ = v7629
	var v7630 int32
	_ = v7630
	var v7633 int32
	_ = v7633
	var v7637 int32
	_ = v7637
	var v7638 int32
	_ = v7638
	var v7640 int32
	_ = v7640
	var v7641 int32
	_ = v7641
	var v7642 int32
	_ = v7642
	var v7645 int32
	_ = v7645
	var v7646 int32
	_ = v7646
	var v7649 int32
	_ = v7649
	var v7650 int32
	_ = v7650
	var v7651 int32
	_ = v7651
	var v7657 int32
	_ = v7657
	var v7659 int32
	_ = v7659
	var v7662 int32
	_ = v7662
	var v7668 int32
	_ = v7668
	var v7669 int32
	_ = v7669
	var v7673 int32
	_ = v7673
	var v7674 int32
	_ = v7674
	var v7676 int32
	_ = v7676
	var v7677 int32
	_ = v7677
	var v7678 int32
	_ = v7678
	var v7679 int32
	_ = v7679
	var v7683 int32
	_ = v7683
	var v7684 int32
	_ = v7684
	var v7688 int32
	_ = v7688
	var v7693 int32
	_ = v7693
	var v7695 int32
	_ = v7695
	var v7700 int32
	_ = v7700
	var v7702 int32
	_ = v7702
	var v7704 int32
	_ = v7704
	var v7705 int32
	_ = v7705
	var v7708 int32
	_ = v7708
	var v7711 int32
	_ = v7711
	var v7712 int32
	_ = v7712
	var v7737 int32
	_ = v7737
	var v7774 int32
	_ = v7774
	var v7787 int32
	_ = v7787
	var v7790 int32
	_ = v7790
	var v7793 int32
	_ = v7793
	var v7794 int32
	_ = v7794
	var v7809 int32
	_ = v7809
	var v7810 int32
	_ = v7810
	var v7816 int32
	_ = v7816
	var v7817 int32
	_ = v7817
	var v7820 int32
	_ = v7820
	var v7860 int32
	_ = v7860
	var v7861 int32
	_ = v7861
	var v7864 int32
	_ = v7864
	var v7869 int32
	_ = v7869
	var v7905 int32
	_ = v7905
	var v7906 int32
	_ = v7906
	var v7908 int32
	_ = v7908
	var v7909 int32
	_ = v7909
	var v7910 int32
	_ = v7910
	var v7911 int32
	_ = v7911
	var v7922 int32
	_ = v7922
	var v7925 int32
	_ = v7925
	var v7928 int32
	_ = v7928
	var v7932 int32
	_ = v7932
	var v7969 int32
	_ = v7969
	var v7970 int64
	_ = v7970
	var v7974 int32
	_ = v7974
	var v7979 int32
	_ = v7979
	var v7983 int32
	_ = v7983
	var v7988 int64
	_ = v7988
	var v7992 int32
	_ = v7992
	var v7994 int32
	_ = v7994
	var v7995 int32
	_ = v7995
	var v7998 int32
	_ = v7998
	var v8009 int32
	_ = v8009
	var v8050 int32
	_ = v8050
	var v8051 int32
	_ = v8051
	var v8054 int32
	_ = v8054
	var v8056 int32
	_ = v8056
	var v8057 int32
	_ = v8057
	var v8059 int32
	_ = v8059
	var v8060 int32
	_ = v8060
	var v8063 int32
	_ = v8063
	var v8068 int32
	_ = v8068
	var v8072 int32
	_ = v8072
	var v8073 int32
	_ = v8073
	var v8078 int32
	_ = v8078
	var v8079 int32
	_ = v8079
	var v8089 int32
	_ = v8089
	var v8091 int32
	_ = v8091
	var v8095 int32
	_ = v8095
	var v8096 int32
	_ = v8096
	var v8099 int32
	_ = v8099
	var v8100 int32
	_ = v8100
	var v8101 int32
	_ = v8101
	var v8104 int32
	_ = v8104
	var v8108 int32
	_ = v8108
	var v8111 int32
	_ = v8111
	var v8120 int32
	_ = v8120
	var v8122 int32
	_ = v8122
	var v8123 int32
	_ = v8123
	var v8126 int32
	_ = v8126
	var v8130 int32
	_ = v8130
	var v8134 int32
	_ = v8134
	var v8135 int32
	_ = v8135
	var v8138 int32
	_ = v8138
	var v8146 int32
	_ = v8146
	var v8147 int32
	_ = v8147
	var v8152 int32
	_ = v8152
	var v8159 int32
	_ = v8159
	var v8164 int32
	_ = v8164
	var v8168 int32
	_ = v8168
	var v8172 int64
	_ = v8172
	var v8176 int32
	_ = v8176
	var v8177 int32
	_ = v8177
	var v8181 int32
	_ = v8181
	var v8182 int32
	_ = v8182
	var v8186 int32
	_ = v8186
	var v8188 int32
	_ = v8188
	var v8190 int32
	_ = v8190
	var v8191 int32
	_ = v8191
	var v8192 int32
	_ = v8192
	var v8198 int32
	_ = v8198
	var v8199 int32
	_ = v8199
	var v8203 int32
	_ = v8203
	var v8204 int32
	_ = v8204
	var v8207 int32
	_ = v8207
	var v8210 int32
	_ = v8210
	var v8211 int32
	_ = v8211
	var v8212 int32
	_ = v8212
	var v8216 int32
	_ = v8216
	var v8217 int32
	_ = v8217
	var v8223 int32
	_ = v8223
	var v8224 int32
	_ = v8224
	var v8225 int32
	_ = v8225
	var v8226 int32
	_ = v8226
	var v8227 int32
	_ = v8227
	var v8228 int32
	_ = v8228
	var v8229 int32
	_ = v8229
	var v8231 int32
	_ = v8231
	var v8232 int32
	_ = v8232
	var v8237 int32
	_ = v8237
	var v8240 int32
	_ = v8240
	var v8243 int32
	_ = v8243
	var v8246 int32
	_ = v8246
	var v8247 int32
	_ = v8247
	var v8251 int32
	_ = v8251
	var v8288 int32
	_ = v8288
	var v8289 int64
	_ = v8289
	var v8293 int32
	_ = v8293
	var v8298 int32
	_ = v8298
	var v8302 int32
	_ = v8302
	var v8307 int64
	_ = v8307
	var v8311 int32
	_ = v8311
	var v8313 int32
	_ = v8313
	var v8314 int32
	_ = v8314
	var v8317 int32
	_ = v8317
	var v8328 int32
	_ = v8328
	var v8332 int32
	_ = v8332
	var v8335 int32
	_ = v8335
	var v8336 int32
	_ = v8336
	var v8339 int32
	_ = v8339
	var v8340 int32
	_ = v8340
	var v8342 int32
	_ = v8342
	var v8343 int32
	_ = v8343
	var v8346 int32
	_ = v8346
	var v8350 int32
	_ = v8350
	var v8387 int32
	_ = v8387
	var v8388 int64
	_ = v8388
	var v8392 int32
	_ = v8392
	var v8397 int32
	_ = v8397
	var v8401 int32
	_ = v8401
	var v8406 int64
	_ = v8406
	var v8410 int32
	_ = v8410
	var v8412 int32
	_ = v8412
	var v8413 int32
	_ = v8413
	var v8416 int32
	_ = v8416
	var v8427 int32
	_ = v8427
	var v8467 int32
	_ = v8467
	var v8474 int32
	_ = v8474
	var v8482 int32
	_ = v8482
	var v8483 int32
	_ = v8483
	var v8487 int32
	_ = v8487
	var v8489 int32
	_ = v8489
	var v8492 int32
	_ = v8492
	var v8497 int32
	_ = v8497
	var v8510 int32
	_ = v8510
	var v8511 int64
	_ = v8511
	var v8512 int32
	_ = v8512
	var v8513 int32
	_ = v8513
	var v8517 int64
	_ = v8517
	var v8518 int32
	_ = v8518
	var v8525 int32
	_ = v8525
	var v8532 int32
	_ = v8532
	var v8533 int32
	_ = v8533
	var v8538 int32
	_ = v8538
	var v8539 int32
	_ = v8539
	var v8540 int32
	_ = v8540
	var v8542 int32
	_ = v8542
	var v8543 int32
	_ = v8543
	var v8546 int32
	_ = v8546
	var v8551 int32
	_ = v8551
	var v8585 int32
	_ = v8585
	var v8589 int32
	_ = v8589
	var v8590 int32
	_ = v8590
	var v8591 int32
	_ = v8591
	var v8593 int32
	_ = v8593
	var v8596 int32
	_ = v8596
	var v8597 int32
	_ = v8597
	var v8638 int32
	_ = v8638
	var v8639 int32
	_ = v8639
	var v8643 int32
	_ = v8643
	var v8647 int32
	_ = v8647
	var v8651 int32
	_ = v8651
	var v8653 int32
	_ = v8653
	var v8658 int32
	_ = v8658
	var v8664 int32
	_ = v8664
	var v8666 int32
	_ = v8666
	var v8669 int32
	_ = v8669
	var v8673 int32
	_ = v8673
	var v8677 int32
	_ = v8677
	var v8678 int32
	_ = v8678
	var v8681 int32
	_ = v8681
	var v8689 int32
	_ = v8689
	var v8695 int32
	_ = v8695
	var v8698 int32
	_ = v8698
	var v8733 int32
	_ = v8733
	var v8734 int32
	_ = v8734
	var v8741 int32
	_ = v8741
	var v8744 int32
	_ = v8744
	var v8747 int32
	_ = v8747
	var v8748 int32
	_ = v8748
	var v8749 int32
	_ = v8749
	var v8752 int32
	_ = v8752
	var v8755 int32
	_ = v8755
	var v8758 int32
	_ = v8758
	var v8765 int32
	_ = v8765
	var v8767 int32
	_ = v8767
	var v8768 int32
	_ = v8768
	var v8771 int32
	_ = v8771
	var v8772 int32
	_ = v8772
	var v8786 int32
	_ = v8786
	var v8790 int32
	_ = v8790
	var v8791 int32
	_ = v8791
	var v8792 int32
	_ = v8792
	var v8794 int32
	_ = v8794
	var v8795 int32
	_ = v8795
	var v8797 int32
	_ = v8797
	var v8798 int32
	_ = v8798
	var v8803 int32
	_ = v8803
	var v8813 int32
	_ = v8813
	var v8816 int32
	_ = v8816
	var v8817 int32
	_ = v8817
	var v8819 int32
	_ = v8819
	var v8823 int32
	_ = v8823
	var v8825 int32
	_ = v8825
	var v8828 int32
	_ = v8828
	var v8829 int32
	_ = v8829
	var v8831 int32
	_ = v8831
	var v8840 int32
	_ = v8840
	var v8845 int32
	_ = v8845
	var v8846 int32
	_ = v8846
	var v8850 int32
	_ = v8850
	var v8852 int32
	_ = v8852
	var v8857 int32
	_ = v8857
	var v8858 int32
	_ = v8858
	var v8860 int32
	_ = v8860
	var v8864 int32
	_ = v8864
	var v8867 int32
	_ = v8867
	var v8868 int32
	_ = v8868
	var v8873 int32
	_ = v8873
	var v8874 int32
	_ = v8874
	var v8884 int32
	_ = v8884
	var v8886 int32
	_ = v8886
	var v8890 int32
	_ = v8890
	var v8891 int32
	_ = v8891
	var v8894 int32
	_ = v8894
	var v8897 int32
	_ = v8897
	var v8902 int32
	_ = v8902
	var v8908 int32
	_ = v8908
	var v8917 int32
	_ = v8917
	var v8919 int32
	_ = v8919
	var v8920 int32
	_ = v8920
	var v8923 int32
	_ = v8923
	var v8927 int32
	_ = v8927
	var v8931 int32
	_ = v8931
	var v8932 int32
	_ = v8932
	var v8935 int32
	_ = v8935
	var v8943 int32
	_ = v8943
	var v8944 int32
	_ = v8944
	var v8949 int32
	_ = v8949
	var v8956 int32
	_ = v8956
	var v8961 int32
	_ = v8961
	var v8965 int32
	_ = v8965
	var v8969 int64
	_ = v8969
	var v8975 int32
	_ = v8975
	var v8978 int32
	_ = v8978
	var v8981 int32
	_ = v8981
	var v8982 int32
	_ = v8982
	var v8984 int32
	_ = v8984
	var v8987 int32
	_ = v8987
	var v8988 int32
	_ = v8988
	var v8997 int32
	_ = v8997
	var v8998 int32
	_ = v8998
	var v9000 int32
	_ = v9000
	var v9002 int32
	_ = v9002
	var v9003 int32
	_ = v9003
	var v9005 int32
	_ = v9005
	var v9012 int32
	_ = v9012
	var v9014 int32
	_ = v9014
	var v9016 int32
	_ = v9016
	var v9023 int32
	_ = v9023
	var v9027 int32
	_ = v9027
	var v9028 int32
	_ = v9028
	var v9031 int32
	_ = v9031
	var v9032 int32
	_ = v9032
	var v9033 int32
	_ = v9033
	var v9034 int32
	_ = v9034
	var v9037 int32
	_ = v9037
	var v9040 int32
	_ = v9040
	var v9043 int32
	_ = v9043
	var v9045 int32
	_ = v9045
	var v9048 int32
	_ = v9048
	var v9051 int32
	_ = v9051
	var v9053 int32
	_ = v9053
	var v9054 int32
	_ = v9054
	var v9055 int32
	_ = v9055
	var v9057 int32
	_ = v9057
	var v9059 int32
	_ = v9059
	var v9066 int32
	_ = v9066
	var v9079 int32
	_ = v9079
	var v9080 int32
	_ = v9080
	var v9081 int32
	_ = v9081
	var v9083 int32
	_ = v9083
	var v9087 int32
	_ = v9087
	var v9088 int32
	_ = v9088
	var v9090 int32
	_ = v9090
	var v9091 int32
	_ = v9091
	var v9092 int32
	_ = v9092
	var v9094 int32
	_ = v9094
	var v9100 int32
	_ = v9100
	var v9101 int32
	_ = v9101
	var v9102 int32
	_ = v9102
	var v9103 int32
	_ = v9103
	var v9106 int32
	_ = v9106
	var v9112 int32
	_ = v9112
	var v9113 int32
	_ = v9113
	var v9114 int32
	_ = v9114
	var v9117 int32
	_ = v9117
	var v9120 int32
	_ = v9120
	var v9125 int32
	_ = v9125
	var v9126 int32
	_ = v9126
	var v9128 int32
	_ = v9128
	var v9130 int32
	_ = v9130
	var v9134 int32
	_ = v9134
	var v9135 int32
	_ = v9135
	var v9136 int32
	_ = v9136
	var v9141 int32
	_ = v9141
	var v9142 int32
	_ = v9142
	var v9143 int32
	_ = v9143
	var v9146 int32
	_ = v9146
	var v9147 int32
	_ = v9147
	var v9148 int32
	_ = v9148
	var v9150 int32
	_ = v9150
	var v9154 int32
	_ = v9154
	var v9155 int32
	_ = v9155
	var v9157 int32
	_ = v9157
	var v9159 int32
	_ = v9159
	var v9161 int32
	_ = v9161
	var v9162 int32
	_ = v9162
	var v9165 int32
	_ = v9165
	var v9172 int32
	_ = v9172
	var v9176 int32
	_ = v9176
	var v9178 int32
	_ = v9178
	var v9180 int32
	_ = v9180
	var v9183 int32
	_ = v9183
	var v9188 int32
	_ = v9188
	var v9189 int32
	_ = v9189
	var v9198 int32
	_ = v9198
	var v9203 int32
	_ = v9203
	var v9205 int32
	_ = v9205
	var v9207 int32
	_ = v9207
	var v9209 int32
	_ = v9209
	var v9210 int32
	_ = v9210
	var v9212 int32
	_ = v9212
	var v9213 int32
	_ = v9213
	var v9214 int32
	_ = v9214
	var v9216 int32
	_ = v9216
	var v9218 int32
	_ = v9218
	var v9219 int32
	_ = v9219
	var v9221 int32
	_ = v9221
	var v9222 int32
	_ = v9222
	var v9225 int32
	_ = v9225
	var v9227 int32
	_ = v9227
	var v9228 int32
	_ = v9228
	var v9230 int32
	_ = v9230
	var v9231 int32
	_ = v9231
	var v9233 int32
	_ = v9233
	var v9235 int32
	_ = v9235
	var v9237 int32
	_ = v9237
	var v9238 int64
	_ = v9238
	var v9244 int32
	_ = v9244
	var v9245 int32
	_ = v9245
	var v9251 int32
	_ = v9251
	var v9252 int32
	_ = v9252
	var v9256 int32
	_ = v9256
	var v9293 int32
	_ = v9293
	var v9294 int32
	_ = v9294
	var v9297 int32
	_ = v9297
	var v9305 int32
	_ = v9305
	var v9336 int32
	_ = v9336
	var v9337 int32
	_ = v9337
	var v9340 int32
	_ = v9340
	var v9350 int32
	_ = v9350
	var v9354 int32
	_ = v9354
	var v9363 int32
	_ = v9363
	var v9397 int32
	_ = v9397
	var v9398 int32
	_ = v9398
	var v9400 int32
	_ = v9400
	var v9401 int32
	_ = v9401
	var v9404 int32
	_ = v9404
	var v9406 int32
	_ = v9406
	var v9411 int32
	_ = v9411
	var v9412 int32
	_ = v9412
	var v9413 int32
	_ = v9413
	var v9421 int32
	_ = v9421
	var v9422 int32
	_ = v9422
	var v9424 int32
	_ = v9424
	var v9432 int32
	_ = v9432
	var v9433 int32
	_ = v9433
	var v9438 int32
	_ = v9438
	var v9444 int32
	_ = v9444
	var v9448 int32
	_ = v9448
	var v9449 int32
	_ = v9449
	var v9450 int32
	_ = v9450
	var v9451 int32
	_ = v9451
	var v9453 int32
	_ = v9453
	var v9454 int32
	_ = v9454
	var v9456 int32
	_ = v9456
	var v9457 int32
	_ = v9457
	var v9461 int32
	_ = v9461
	var v9464 int32
	_ = v9464
	var v9468 int32
	_ = v9468
	var v9474 int32
	_ = v9474
	var v9476 int32
	_ = v9476
	var v9481 int32
	_ = v9481
	var v9482 int32
	_ = v9482
	var v9483 int32
	_ = v9483
	var v9485 int32
	_ = v9485
	var v9486 int32
	_ = v9486
	var v9488 int32
	_ = v9488
	var v9489 int32
	_ = v9489
	var v9494 int32
	_ = v9494
	var v9533 int32
	_ = v9533
	var v9534 int32
	_ = v9534
	var v9536 int32
	_ = v9536
	var v9537 int32
	_ = v9537
	var v9549 int32
	_ = v9549
	var v9585 int32
	_ = v9585
	var v9589 int32
	_ = v9589
	var v9591 int32
	_ = v9591
	var v9597 int32
	_ = v9597
	var v9600 int32
	_ = v9600
	var v9604 int32
	_ = v9604
	var v9609 int32
	_ = v9609
	var v9613 int32
	_ = v9613
	var v9616 int32
	_ = v9616
	var v9620 int32
	_ = v9620
	var v9625 int32
	_ = v9625
	var v9629 int32
	_ = v9629
	var v9632 int32
	_ = v9632
	var v9640 int32
	_ = v9640
	var v9645 int32
	_ = v9645
	var v9649 int32
	_ = v9649
	var v9659 int32
	_ = v9659
	var v9664 int32
	_ = v9664
	var v9668 int32
	_ = v9668
	var v9671 int32
	_ = v9671
	var v9673 int32
	_ = v9673
	var v9679 int32
	_ = v9679
	var v9684 int32
	_ = v9684
	var v9688 int32
	_ = v9688
	var v9691 int32
	_ = v9691
	var v9698 int32
	_ = v9698
	var v9703 int32
	_ = v9703
	var v9707 int32
	_ = v9707
	var v9710 int32
	_ = v9710
	var v9716 int32
	_ = v9716
	var v9721 int32
	_ = v9721
	var v9725 int32
	_ = v9725
	var v9728 int32
	_ = v9728
	var v9736 int32
	_ = v9736
	var v9741 int32
	_ = v9741
	var v9745 int32
	_ = v9745
	var v9748 int32
	_ = v9748
	var v9755 int32
	_ = v9755
	var v9760 int32
	_ = v9760
	var v9799 int32
	_ = v9799
	var v9800 int32
	_ = v9800
	var v9801 int32
	_ = v9801
	var v9802 int32
	_ = v9802
	var v9839 int32
	_ = v9839
	var v9841 int32
	_ = v9841
	var v9843 int32
	_ = v9843
	var v9844 int32
	_ = v9844
	var v9845 int32
	_ = v9845
	var v9850 int32
	_ = v9850
	var v9857 int32
	_ = v9857
	var v9858 int32
	_ = v9858
	var v9859 int32
	_ = v9859
	var v9873 int32
	_ = v9873
	var v9874 int32
	_ = v9874
	var v9875 int32
	_ = v9875
	var v9877 int32
	_ = v9877
	var v9882 int32
	_ = v9882
	var v9885 int32
	_ = v9885
	var v9886 int64
	_ = v9886
	var v9890 int32
	_ = v9890
	var v9896 int32
	_ = v9896
	var v9900 int32
	_ = v9900
	var v9901 int32
	_ = v9901
	var v9902 int32
	_ = v9902
	var v9903 int32
	_ = v9903
	var v9906 int32
	_ = v9906
	var v9909 int32
	_ = v9909
	var v9910 int32
	_ = v9910
	var v9911 int32
	_ = v9911
	var v9918 int32
	_ = v9918
	var v9919 int32
	_ = v9919
	var v9923 int32
	_ = v9923
	var v9928 int32
	_ = v9928
	var v9929 int32
	_ = v9929
	var v9934 int32
	_ = v9934
	var v9935 int32
	_ = v9935
	var v9936 int32
	_ = v9936
	var v9938 int32
	_ = v9938
	var v9940 int32
	_ = v9940
	var v9941 int32
	_ = v9941
	var v9942 int32
	_ = v9942
	var v9944 int32
	_ = v9944
	var v9946 int32
	_ = v9946
	var v9965 int32
	_ = v9965
	var v9971 int32
	_ = v9971
	var v9973 int32
	_ = v9973
	var v9977 int32
	_ = v9977
	var v9980 int32
	_ = v9980
	var v9987 int32
	_ = v9987
	var v9992 int32
	_ = v9992
	var v9998 int32
	_ = v9998
	var v10001 int32
	_ = v10001
	var v10002 int32
	_ = v10002
	var v10003 int32
	_ = v10003
	var v10004 int32
	_ = v10004
	var v10006 int32
	_ = v10006
	var v10008 int32
	_ = v10008
	var v10016 int32
	_ = v10016
	var v10018 int32
	_ = v10018
	var v10024 int32
	_ = v10024
	var v10025 int32
	_ = v10025
	var v10030 int32
	_ = v10030
	var v10031 int32
	_ = v10031
	var v10041 int32
	_ = v10041
	var v10045 int32
	_ = v10045
	var v10046 int32
	_ = v10046
	var v10058 int32
	_ = v10058
	var v10060 int32
	_ = v10060
	var v10063 int32
	_ = v10063
	var v10072 int32
	_ = v10072
	var v10075 int32
	_ = v10075
	var v10077 int32
	_ = v10077
	var v10079 int32
	_ = v10079
	var v10081 int32
	_ = v10081
	var v10084 int32
	_ = v10084
	var v10087 int32
	_ = v10087
	var v10092 int32
	_ = v10092
	var v10093 int32
	_ = v10093
	var v10096 int32
	_ = v10096
	var v10097 int32
	_ = v10097
	var v10101 int32
	_ = v10101
	var v10104 int32
	_ = v10104
	var v10107 int32
	_ = v10107
	var v10109 int32
	_ = v10109
	var v10116 int32
	_ = v10116
	var v10117 int32
	_ = v10117
	var v10118 int32
	_ = v10118
	var v10123 int32
	_ = v10123
	var v10126 int32
	_ = v10126
	var v10131 int32
	_ = v10131
	var v10133 int32
	_ = v10133
	var v10137 int32
	_ = v10137
	var v10141 int64
	_ = v10141
	var v10145 int32
	_ = v10145
	var v10146 int32
	_ = v10146
	var v10149 int32
	_ = v10149
	var v10150 int32
	_ = v10150
	var v10154 int32
	_ = v10154
	var v10158 int32
	_ = v10158
	var v10161 int32
	_ = v10161
	var v10163 int32
	_ = v10163
	var v10165 int32
	_ = v10165
	var v10166 int32
	_ = v10166
	var v10167 int32
	_ = v10167
	var v10169 int32
	_ = v10169
	var v10172 int32
	_ = v10172
	var v10174 int32
	_ = v10174
	var v10175 int32
	_ = v10175
	var v10182 int32
	_ = v10182
	var v10184 int32
	_ = v10184
	var v10187 int32
	_ = v10187
	var v10191 int32
	_ = v10191
	var v10195 int32
	_ = v10195
	var v10197 int32
	_ = v10197
	var v10198 int32
	_ = v10198
	var v10199 int32
	_ = v10199
	var v10201 int32
	_ = v10201
	var v10205 int32
	_ = v10205
	var v10209 int32
	_ = v10209
	var v10213 int32
	_ = v10213
	var v10220 int32
	_ = v10220
	var v10253 int32
	_ = v10253
	var v10257 int32
	_ = v10257
	var v10261 int32
	_ = v10261
	var v10262 int32
	_ = v10262
	var v10263 int32
	_ = v10263
	var v10265 int32
	_ = v10265
	var v10267 int32
	_ = v10267
	var v10269 int32
	_ = v10269
	var v10271 int32
	_ = v10271
	var v10288 int32
	_ = v10288
	var v10289 int32
	_ = v10289
	var v10329 int32
	_ = v10329
	var v10330 int32
	_ = v10330
	var v10333 int32
	_ = v10333
	var v10334 int32
	_ = v10334
	var v10336 int32
	_ = v10336
	var v10339 int32
	_ = v10339
	var v10341 int32
	_ = v10341
	var v10344 int32
	_ = v10344
	var v10346 int32
	_ = v10346
	var v10347 int32
	_ = v10347
	var v10351 int32
	_ = v10351
	var v10352 int32
	_ = v10352
	var v10359 int32
	_ = v10359
	var v10361 int32
	_ = v10361
	var v10364 int32
	_ = v10364
	var v10366 int32
	_ = v10366
	var v10367 int32
	_ = v10367
	var v10368 int32
	_ = v10368
	var v10370 int32
	_ = v10370
	var v10373 int32
	_ = v10373
	var v10377 int32
	_ = v10377
	var v10380 int32
	_ = v10380
	var v10386 int32
	_ = v10386
	var v10391 int32
	_ = v10391
	var v10395 int32
	_ = v10395
	var v10397 int32
	_ = v10397
	var v10401 int32
	_ = v10401
	var v10402 int32
	_ = v10402
	var v10403 int32
	_ = v10403
	var v10404 int32
	_ = v10404
	var v10408 int32
	_ = v10408
	var v10411 int32
	_ = v10411
	var v10412 int32
	_ = v10412
	var v10422 int32
	_ = v10422
	var v10425 int32
	_ = v10425
	var v10427 int32
	_ = v10427
	var v10429 int32
	_ = v10429
	var v10431 int32
	_ = v10431
	var v10434 int32
	_ = v10434
	var v10440 int32
	_ = v10440
	var v10447 int32
	_ = v10447
	var v10450 int32
	_ = v10450
	var v10454 int32
	_ = v10454
	var v10457 int32
	_ = v10457
	var v10463 int32
	_ = v10463
	var v10468 int32
	_ = v10468
	var v10471 int32
	_ = v10471
	var v10514 int32
	_ = v10514
	var v10517 int32
	_ = v10517
	var v10521 int32
	_ = v10521
	var v10526 int32
	_ = v10526
	var v10530 int32
	_ = v10530
	var v10533 int32
	_ = v10533
	var v10537 int32
	_ = v10537
	var v10542 int32
	_ = v10542
	var v10546 int32
	_ = v10546
	var v10549 int32
	_ = v10549
	var v10553 int32
	_ = v10553
	var v10555 int32
	_ = v10555
	var v10560 int32
	_ = v10560
	var v10564 int32
	_ = v10564
	var v10567 int32
	_ = v10567
	var v10571 int32
	_ = v10571
	var v10576 int32
	_ = v10576
	var v10580 int32
	_ = v10580
	var v10583 int32
	_ = v10583
	var v10587 int32
	_ = v10587
	var v10592 int32
	_ = v10592
	var v10596 int32
	_ = v10596
	var v10599 int32
	_ = v10599
	var v10606 int32
	_ = v10606
	var v10611 int32
	_ = v10611
	var v10615 int32
	_ = v10615
	var v10618 int32
	_ = v10618
	var v10619 int32
	_ = v10619
	var v10627 int32
	_ = v10627
	var v10632 int32
	_ = v10632
	var v10637 int32
	_ = v10637
	var v10640 int32
	_ = v10640
	var v10644 int32
	_ = v10644
	var v10646 int32
	_ = v10646
	var v10651 int32
	_ = v10651
	var v10655 int32
	_ = v10655
	var v10658 int32
	_ = v10658
	var v10666 int32
	_ = v10666
	var v10671 int32
	_ = v10671
	var v10675 int32
	_ = v10675
	var v10678 int32
	_ = v10678
	var v10685 int32
	_ = v10685
	var v10690 int32
	_ = v10690
	var v10694 int32
	_ = v10694
	var v10697 int32
	_ = v10697
	var v10701 int32
	_ = v10701
	var v10706 int32
	_ = v10706
	var v10710 int32
	_ = v10710
	var v10713 int32
	_ = v10713
	var v10719 int32
	_ = v10719
	var v10724 int32
	_ = v10724
	var v10729 int32
	_ = v10729
	var v10732 int32
	_ = v10732
	var v10736 int32
	_ = v10736
	var v10738 int32
	_ = v10738
	var v10743 int32
	_ = v10743
	var v10747 int32
	_ = v10747
	var v10750 int32
	_ = v10750
	var v10754 int32
	_ = v10754
	var v10759 int32
	_ = v10759
	var v10763 int32
	_ = v10763
	var v10766 int32
	_ = v10766
	var v10770 int32
	_ = v10770
	var v10775 int32
	_ = v10775
	var v10779 int32
	_ = v10779
	var v10782 int32
	_ = v10782
	var v10788 int32
	_ = v10788
	var v10793 int32
	_ = v10793
	var v10797 int32
	_ = v10797
	var v10800 int32
	_ = v10800
	var v10804 int32
	_ = v10804
	var v10809 int32
	_ = v10809
	var v10813 int32
	_ = v10813
	var v10816 int32
	_ = v10816
	var v10820 int32
	_ = v10820
	var v10825 int32
	_ = v10825
	var v10829 int32
	_ = v10829
	var v10832 int32
	_ = v10832
	var v10836 int32
	_ = v10836
	var v10838 int32
	_ = v10838
	var v10843 int32
	_ = v10843
	var v10847 int32
	_ = v10847
	var v10850 int32
	_ = v10850
	var v10856 int32
	_ = v10856
	var v10861 int32
	_ = v10861
	var v10865 int32
	_ = v10865
	var v10868 int32
	_ = v10868
	var v10872 int32
	_ = v10872
	var v10874 int32
	_ = v10874
	var v10879 int32
	_ = v10879
	v1 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(528)
	m.G0 = v39
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v43
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[2])) = uint8(v1)
	F_MemoryContextReset(m, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_initStringInfo(m, v39+int32(440))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[3]))
	if v55 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])))
	if v109 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[5]))
	if v59 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[6]))
	if v61 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v64 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_pairingheap_remove(m, int32(_a_F_PostgresMainLoopOnce_0), v55+int32(52))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[3])) = v70
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
	v101 = v70
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[8])) = v101
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+40)) = v101
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
	v101 = v100
	goto L13
L19:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
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
	if (v114-int32(7))&int32(-9) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v122 = int32(0)
	F_pgstat_report_activity(m, int32(6), v122)
	mBase = m.M
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[12]))
	if v125 <= v122 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	goto L33
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[13]))
	if v129 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v132 = base.B2i32(v129 <= v125)
	goto L30
L29:
	;
	v132 = int32(0)
	goto L30
L30:
	;
	if v132 != 0 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v134)
	F_enable_timeout_after(m, int32(7), v125)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	if v141 != int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v145 = int32(0)
	F_pgstat_report_activity(m, int32(4), v145)
	mBase = m.M
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[12]))
	if v148 <= v145 {
		goto L22
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[15]))
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[13]))
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v155 = base.B2i32(v152 <= v148)
	goto L40
L39:
	;
	v155 = int32(0)
	goto L40
L40:
	;
	if v155 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v157)
	F_enable_timeout_after(m, int32(7), v148)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v168 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	if int32(0) < v168 {
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
	F_enable_timeout_after(m, int32(10), v168)
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
	F___gettimeofday(m, v223)
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
	*(*float64)(unsafe.Add(mBase, uint32(v39)+432)) = base.F64_div(base.F64_convert_i64_u(v256), float64(1000))
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
	*(*float64)(unsafe.Add(mBase, uint32(v39)+424)) = base.F64_div(base.F64_convert_i64_u(v264), float64(1000))
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
	*(*float64)(unsafe.Add(mBase, uint32(v39)+416)) = base.F64_div(base.F64_convert_i64_u(v272), float64(1000))
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_1), v39+int32(416))
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
	v622 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])))
	if v622 == int32(1) {
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
	v417 = F_pq_getmessage(m, v39+int32(440), v416)
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
		v587 = v334
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
	v587 = v334
	goto L79
L104:
	;
	if v359 == int32(0) {
		v587 = v356
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
	v587 = v356
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
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v325
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_11), v39)
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
	v587 = int32(-1)
	goto L79
L115:
	;
	goto L116
L116:
	;
	v420 = int32(_a_F_PostgresMainLoopOnce_5)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v422 - int32(1)
	v587 = v325
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
	v435 = v39 + int32(440)
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
	F_appendStringInfoChar(m, v39+int32(440), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L158
	}
L121:
	;
	F_appendStringInfoChar(m, v39+int32(440), int32(10))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L157
	}
L122:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v482 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v39)+444))
	if v558 != 0 {
		goto L120
	} else {
		goto L156
	}
L124:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v485 = F_do_getc(m, v444)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32]))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])))
	if v490 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32])) = v488
	switch v485 + int32(1) {
	case 0:
		goto L145
	default:
		goto L147
	case 11:
		goto L148
	}
L130:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v492 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[33]))
	if v507 == int32(0) {
		goto L129
	} else {
		goto L143
	}
L133:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[34]))
	if v496 != 0 {
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
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[15]))
	if v500 == int32(0) {
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
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L129
L143:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[35]))
	F_SetLatch(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
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
	if v518 <= int32(0) {
		goto L121
	} else {
		goto L154
	}
L147:
	;
	F_appendStringInfoChar(m, v39+int32(440), base.I32_extend8_s(v485))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L153
	}
L148:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v39)+444))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[36])))
	if v520 == int32(0) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	if v518 < int32(2) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v39)+440))
	v526 = v525 + v518
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526-int32(1)))))
	if v529 != int32(10) {
		goto L147
	} else {
		goto L151
	}
L151:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526-int32(2)))))
	if v534 == int32(59) {
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
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v39)+440))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v518-int32(1)))))
	if v549 != int32(92) {
		goto L121
	} else {
		goto L155
	}
L155:
	;
	v553 = v518 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+444)) = v553
	v556 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v545+v553))) = uint8(v556)
	goto L122
L156:
	;
	v587 = int32(-1)
	goto L79
L157:
	;
	goto L120
L158:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[37])))
	if v574 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v39)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+400)) = v575
	F_pg_printf(m, int32(_a_F_PostgresMainLoopOnce_12), v39+int32(400))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v582 = F_fflush(m, v431)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v587 = int32(81)
	goto L79
L164:
	;
	F_disable_timeout(m, int32(7))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])))
	if v632 == int32(1) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v629 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[14])) = uint8(v629)
	goto L166
L168:
	;
	F_disable_timeout(m, int32(9))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v642 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v639 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[18])) = uint8(v639)
	goto L170
L172:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v646 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[10])) = uint8(v646)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[38]))
	if v649 != 0 {
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
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v587 != int32(-1) {
		goto L203
	} else {
		goto L204
	}
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[39])) = int32(99)
	goto L2600
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10865 = m.ExcPending
	if v10865 != 0 {
		goto L1
	} else {
		goto L2595
	}
L182:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10847 = m.ExcPending
	if v10847 != 0 {
		goto L1
	} else {
		goto L2591
	}
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10829 = m.ExcPending
	if v10829 != 0 {
		goto L1
	} else {
		goto L2586
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10813 = m.ExcPending
	if v10813 != 0 {
		goto L1
	} else {
		goto L2582
	}
L185:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10797 = m.ExcPending
	if v10797 != 0 {
		goto L1
	} else {
		goto L2578
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10779 = m.ExcPending
	if v10779 != 0 {
		goto L1
	} else {
		goto L2574
	}
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10763 = m.ExcPending
	if v10763 != 0 {
		goto L1
	} else {
		goto L2570
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10747 = m.ExcPending
	if v10747 != 0 {
		goto L1
	} else {
		goto L2566
	}
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10729 = m.ExcPending
	if v10729 != 0 {
		goto L1
	} else {
		goto L2561
	}
L190:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10710 = m.ExcPending
	if v10710 != 0 {
		goto L1
	} else {
		goto L2557
	}
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10694 = m.ExcPending
	if v10694 != 0 {
		goto L1
	} else {
		goto L2553
	}
L192:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10675 = m.ExcPending
	if v10675 != 0 {
		goto L1
	} else {
		goto L2549
	}
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10655 = m.ExcPending
	if v10655 != 0 {
		goto L1
	} else {
		goto L2545
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10637 = m.ExcPending
	if v10637 != 0 {
		goto L1
	} else {
		goto L2540
	}
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10615 = m.ExcPending
	if v10615 != 0 {
		goto L1
	} else {
		goto L2536
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10596 = m.ExcPending
	if v10596 != 0 {
		goto L1
	} else {
		goto L2532
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10580 = m.ExcPending
	if v10580 != 0 {
		goto L1
	} else {
		goto L2528
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10564 = m.ExcPending
	if v10564 != 0 {
		goto L1
	} else {
		goto L2524
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10546 = m.ExcPending
	if v10546 != 0 {
		goto L1
	} else {
		goto L2519
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10530 = m.ExcPending
	if v10530 != 0 {
		goto L1
	} else {
		goto L2515
	}
L201:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10514 = m.ExcPending
	if v10514 != 0 {
		goto L1
	} else {
		goto L2511
	}
L202:
	;
	m.G0 = v39 + int32(528)
	return
L203:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[28])))
	if v659&int32(1) != 0 {
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	switch v587 + int32(1) {
	case 0:
		goto L210
	default:
		goto L208
	case 67:
		goto L217
	case 68:
		goto L214
	case 69:
		goto L213
	case 70:
		goto L216
	case 71:
		goto L215
	case 73:
		goto L212
	case 81:
		goto L218
	case 82:
		goto L219
	case 84:
		goto L211
	case 89:
		goto L209
	case 100, 101, 103:
		goto L202
	}
L206:
	;
	goto L205
L207:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v10471 = m.ExcPending
	if v10471 != 0 {
		goto L1
	} else {
		goto L2510
	}
L208:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v10454 = m.ExcPending
	if v10454 != 0 {
		goto L1
	} else {
		goto L2506
	}
L209:
	;
	v10440 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10440 == int32(2) {
		goto L2501
	} else {
		goto L2502
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[40])) = int32(2)
	goto L209
L211:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v10408 = m.ExcPending
	if v10408 != 0 {
		goto L1
	} else {
		goto L2487
	}
L212:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v10395 = m.ExcPending
	if v10395 != 0 {
		goto L1
	} else {
		goto L2484
	}
L213:
	;
	v10133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v10133 == int32(1) {
		goto L185
	} else {
		goto L2428
	}
L214:
	;
	v10087 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v10087 == int32(1) {
		goto L187
	} else {
		goto L2410
	}
L215:
	;
	v8961 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v8961 == int32(1) {
		goto L188
	} else {
		goto L2154
	}
L216:
	;
	v8164 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v8164 == int32(1) {
		goto L191
	} else {
		goto L1964
	}
L217:
	;
	v7082 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v7082 == int32(1) {
		goto L198
	} else {
		goto L1750
	}
L218:
	;
	v6606 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v6606 == int32(1) {
		goto L201
	} else {
		goto L1613
	}
L219:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v665 < int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v673 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L224
	}
L221:
	;
	v669 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v669
	goto L223
L222:
	;
	goto L223
L223:
	;
	goto L220
L224:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[41])))
	if v680 == int32(1) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v6602 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v6602)
	goto L202
L227:
	;
	v683 = m.G0
	v685 = v683 - int32(_a_F_PostgresMainLoopOnce_13)
	m.G0 = v685
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v692 == int32(0) {
		v714 = v688
		goto L231
	} else {
		goto L232
	}
L228:
	;
	goto L229
L229:
	;
	v5355 = m.G0
	v5357 = v5355 - int32(112)
	m.G0 = v5357
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v673
	v5362 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v5364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	F_pgstat_report_activity(m, int32(3), v673)
	mBase = m.M
	if v5364 == int32(1) {
		goto L1380
	} else {
		goto L1381
	}
L230:
	;
	if v1147 != 0 {
		goto L226
	} else {
		goto L1379
	}
L231:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v716 != int32(4) {
		goto L267
	} else {
		goto L268
	}
L232:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	if v695 == int32(4) {
		v714 = v688
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v688)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v688)+76)) = int32(1)
	if v698 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_s_lock(m, v688+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v688)+4)) = int32(4)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v714 = v713
	goto L231
L237:
	;
	goto L236
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5306 = m.ExcPending
	if v5306 != 0 {
		goto L1
	} else {
		goto L1376
	}
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5290 = m.ExcPending
	if v5290 != 0 {
		goto L1
	} else {
		goto L1373
	}
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5274 = m.ExcPending
	if v5274 != 0 {
		goto L1
	} else {
		goto L1370
	}
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5258 = m.ExcPending
	if v5258 != 0 {
		goto L1
	} else {
		goto L1367
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5242 = m.ExcPending
	if v5242 != 0 {
		goto L1
	} else {
		goto L1364
	}
L243:
	;
	m.G0 = v685 + int32(_a_F_PostgresMainLoopOnce_13)
	goto L230
L244:
	;
	F_EndReplicationCommand(m, v5162)
	mBase = m.M
	v5190 = m.ExcPending
	if v5190 != 0 {
		goto L1
	} else {
		goto L1362
	}
L245:
	;
	v4864 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+8))
	if v4864 == int32(0) {
		goto L1275
	} else {
		goto L1276
	}
L246:
	;
	v4817 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L1
	} else {
		goto L1270
	}
L247:
	;
	if v4117 == int32(-1) {
		goto L1259
	} else {
		goto L1260
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4766 = m.ExcPending
	if v4766 != 0 {
		goto L1
	} else {
		goto L1255
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4747 = m.ExcPending
	if v4747 != 0 {
		goto L1
	} else {
		goto L1251
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L1
	} else {
		goto L1247
	}
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4705 = m.ExcPending
	if v4705 != 0 {
		goto L1
	} else {
		goto L1243
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4686 = m.ExcPending
	if v4686 != 0 {
		goto L1
	} else {
		goto L1239
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4667 = m.ExcPending
	if v4667 != 0 {
		goto L1
	} else {
		goto L1235
	}
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L1
	} else {
		goto L1231
	}
L255:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L1
	} else {
		goto L1230
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4629 = m.ExcPending
	if v4629 != 0 {
		goto L1
	} else {
		goto L1227
	}
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4613 = m.ExcPending
	if v4613 != 0 {
		goto L1
	} else {
		goto L1223
	}
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4597 = m.ExcPending
	if v4597 != 0 {
		goto L1
	} else {
		goto L1219
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4581 = m.ExcPending
	if v4581 != 0 {
		goto L1
	} else {
		goto L1216
	}
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4565 = m.ExcPending
	if v4565 != 0 {
		goto L1
	} else {
		goto L1212
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		goto L1
	} else {
		goto L1208
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4533 = m.ExcPending
	if v4533 != 0 {
		goto L1
	} else {
		goto L1204
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4517 = m.ExcPending
	if v4517 != 0 {
		goto L1
	} else {
		goto L1200
	}
L264:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L1
	} else {
		goto L1196
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4482 = m.ExcPending
	if v4482 != 0 {
		goto L1
	} else {
		goto L1192
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L1
	} else {
		goto L1188
	}
L267:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48])))
	if v720 != 0 {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	goto L269
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4448 = m.ExcPending
	if v4448 != 0 {
		goto L1
	} else {
		goto L1184
	}
L270:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v749 != 0 {
		goto L281
	} else {
		goto L282
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L278
	}
L272:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+20))
	goto L275
L273:
	;
	goto L274
L274:
	;
	goto L270
L275:
	;
	if base.B2i32(v723 == int32(2)) == int32(0) {
		goto L271
	} else {
		goto L276
	}
L276:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49]))
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])) = v729
	goto L274
L278:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_16), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(609), int32(_a_F_PostgresMainLoopOnce_18))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	if v753 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L283
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v770
	v774 = v685 + int32(428)
	v776 = F_palloc0(m, int32(20))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L291
	}
L286:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[52]))
	v763 = F_AllocSetContextCreateInternal(m, v758, int32(_a_F_PostgresMainLoopOnce_19), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	F_MemoryContextReset(m, v753)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L290
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51])) = v763
	v770 = v763
	goto L285
L290:
	;
	v769 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	v770 = v769
	goto L285
L291:
	;
	if v774 != 0 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v804 = F__emscripten_memset_bulkmem(m, v780, base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L301
L293:
	;
	v780 = F_palloc(m, int32(96))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	v786 = int32(28)
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[32])) = v786
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L298
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = v780
	if v780 != 0 {
		goto L292
	} else {
		goto L297
	}
L297:
	;
	v786 = int32(48)
	goto L295
L298:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_22), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_23), int32(275), int32(_a_F_PostgresMainLoopOnce_24))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v806 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v805)+60)) = v806
	v808 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v805)+52)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v805)+44)) = v806
	*(*int64)(unsafe.Add(mBase, uint32(v805)+36)) = v808
	*(*int64)(unsafe.Add(mBase, uint32(v805)+4)) = v808
	*(*int64)(unsafe.Add(mBase, uint32(v805)+12)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v805)+20)) = v806
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	*(*int32)(unsafe.Add(mBase, uint32(v820))) = v776
	v822 = F_strlen(m, v673)
	mBase = m.M
	if base.Ui32(v822) < base.Ui32(int32(-2)) {
		goto L305
	} else {
		goto L306
	}
L302:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	v1129 = m.G0
	v1131 = v1129 - int32(16)
	m.G0 = v1131
	v1135 = F_replication_yylex(m, v1131+int32(8), v1128)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L341
	}
L303:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_25))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L339
	}
L304:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_26))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L338
	}
L305:
	;
	v827 = v822 + int32(2)
	v828 = F_palloc(m, v827)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_27))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L337
	}
L308:
	;
	if v828 == int32(0) {
		goto L304
	} else {
		goto L309
	}
L309:
	;
	if v822 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1023 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v822+v828))) = uint16(v1023)
	if base.Ui32(v827) < base.Ui32(int32(2)) {
		v1110 = v1023
		goto L324
	} else {
		goto L325
	}
L311:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v822) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v857 = v1
	v858 = v1
	goto L315
L313:
	;
	v921 = v1
	goto L314
L314:
	;
	v938 = v822 & int32(3)
	if v938 == int32(0) {
		goto L310
	} else {
		goto L318
	}
L315:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v858))))
	*(*uint8)(unsafe.Add(mBase, uint32(v828+v858))) = uint8(v876)
	v879 = v858 | int32(1)
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v879))))
	*(*uint8)(unsafe.Add(mBase, uint32(v828+v879))) = uint8(v882)
	v885 = v858 | int32(2)
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v885))))
	*(*uint8)(unsafe.Add(mBase, uint32(v828+v885))) = uint8(v888)
	v891 = v858 | int32(3)
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v891))))
	*(*uint8)(unsafe.Add(mBase, uint32(v828+v891))) = uint8(v894)
	v896 = int32(4)
	v897 = v858 + v896
	v899 = v857 + v896
	if v899 != v822&int32(-4) {
		v857 = v899
		v858 = v897
		goto L315
	} else {
		goto L317
	}
L316:
	;
	v921 = v897
	goto L314
L317:
	;
	goto L316
L318:
	;
	v952 = v806
	v961 = v921
	goto L319
L319:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+v961))))
	*(*uint8)(unsafe.Add(mBase, uint32(v828+v961))) = uint8(v979)
	v981 = int32(1)
	v984 = v952 + v981
	if v984 != v938 {
		v952 = v984
		v961 = v961 + v981
		goto L319
	} else {
		goto L321
	}
L320:
	;
	goto L310
L321:
	;
	goto L320
L322:
	;
	if v1110 == int32(0) {
		goto L303
	} else {
		goto L336
	}
L323:
	;
	F_yy_fatal_error_3(m, int32(_a_F_PostgresMainLoopOnce_28))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L335
	}
L324:
	;
	goto L322
L325:
	;
	v1029 = v827 - int32(2)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828+v1029))))
	if v1031 != 0 {
		v1110 = v1023
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827+v828-int32(1)))))
	if v1035 != 0 {
		v1110 = v1023
		goto L324
	} else {
		goto L327
	}
L327:
	;
	v1037 = F_palloc(m, int32(48))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	if v1037 == int32(0) {
		goto L323
	} else {
		goto L329
	}
L329:
	;
	v1041 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+20)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+8)) = v828
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+4)) = v828
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+12)) = v1029
	*(*int64)(unsafe.Add(mBase, uint32(v1037)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1037)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v1037)+16)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v1037))) = v1041
	F_replication_yyensure_buffer_stack(m, v820)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v820)+20))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1055+v1056<<(uint(int32(2))%32))))
	if v1060 == v1037 {
		v1110 = v1037
		goto L324
	} else {
		goto L331
	}
L331:
	;
	if v1060 != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v820)+36))
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1062))) = uint8(v1063)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v820)+20))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v1067 = int32(2)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1065+v1066<<(uint(v1067)%32))))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v820)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1070)+8)) = v1071
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v820)+20))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1074<<(uint(v1067)%32))))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v820)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1078)+16)) = v1079
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v820)+20))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v1083 = v1081
	v1084 = v1082
	goto L334
L333:
	;
	v1083 = v1055
	v1084 = v1056
	goto L334
L334:
	;
	v1085 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1084<<(uint(v1085)%32)+v1083))) = v1037
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v820)+20))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v1093 = v1089 + v1090<<(uint(v1085)%32)
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+28)) = v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+36)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v820)+80)) = v1098
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+4)) = v1102
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1098))))
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+24)) = uint8(v1104)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+48)) = int32(1)
	v1110 = v1037
	goto L324
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+20)) = int32(1)
	goto L302
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	m.G0 = v1131 + int32(16)
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	if v1147 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L341:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v1135-int32(262)) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	if v1135 != int32(282) {
		v1147 = int32(0)
		goto L340
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1128)))
	*(*int32)(unsafe.Add(mBase, uint32(v1144))) = v1135
	v1147 = int32(1)
	goto L340
L345:
	;
	goto L344
L346:
	;
	F_replication_scanner_finish(m, v1151)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v1182 = m.G0
	v1184 = v1182 - int32(2080)
	m.G0 = v1184
	v1188 = v1184 - int32(-64)
	v1190 = v1184 + int32(1664)
	v1192 = v1190
	v1202 = v1188
	v1203 = v1188
	v1204 = v1190
	v1205 = int32(-2)
	v1209 = int32(200)
	v1216 = v1
	goto L360
L349:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v690
	v1159 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	F_MemoryContextReset(m, v1159)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	if v1163 != 0 {
		goto L243
	} else {
		goto L351
	}
L351:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_29), int32(0))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2064), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
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
	if v1342 != 0 {
		goto L266
	} else {
		goto L530
	}
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L526
	}
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L1
	} else {
		goto L522
	}
L359:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_31))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L521
	}
L360:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1204))) = uint16(v1216)
	v1230 = v1209 << (uint(int32(1)) % 32)
	if base.Ui32(v1192+v1230-int32(2)) <= base.Ui32(v1204) {
		goto L366
	} else {
		goto L367
	}
L361:
	;
	F_replication_yyerror(m, int32(_a_F_PostgresMainLoopOnce_32))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L520
	}
L362:
	;
	goto L361
L363:
	;
	v1192 = v1287
	v1202 = v1824
	v1203 = v1292
	v1204 = v1825 + int32(2)
	v1205 = v1826
	v1209 = v1294
	v1216 = v1829
	goto L360
L364:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216)+uint32(_c_F_PostgresMainLoopOnce[54]))))
	if v1361 == int32(0) {
		goto L362
	} else {
		goto L412
	}
L365:
	;
	if v1184+int32(1664) != v1339 {
		goto L408
	} else {
		goto L409
	}
L366:
	;
	if base.Ui32(int32(_a_F_PostgresMainLoopOnce_33)) < base.Ui32(v1209) {
		goto L359
	} else {
		goto L369
	}
L367:
	;
	v1287 = v1192
	v1291 = v1202
	v1292 = v1203
	v1293 = v1204
	v1294 = v1209
	goto L368
L368:
	;
	v1297 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1216)+uint32(_c_F_PostgresMainLoopOnce[55]))))
	if v1297 == int32(-36) {
		v1358 = v1205
		goto L364
	} else {
		goto L390
	}
L369:
	;
	v1237 = int32(_a_F_PostgresMainLoopOnce_7)
	if base.Ui32(v1237) <= base.Ui32(v1230) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1240 = v1237
	goto L372
L371:
	;
	v1240 = v1230
	goto L372
L372:
	;
	v1245 = F_palloc(m, v1240*int32(10)+int32(7))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	if v1245 == int32(0) {
		goto L359
	} else {
		goto L374
	}
L374:
	;
	v1250 = int32(1)
	v1253 = (v1204-v1192)>>(uint(v1250)%32) + v1250
	v1255 = v1253 << (uint(v1250) % 32)
	if v1255 != 0 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1264 = v1257 + (v1240<<(uint(int32(1))%32)+int32(7))&int32(2147483640)
	v1266 = v1253 << (uint(int32(3)) % 32)
	if v1266 != 0 {
		goto L380
	} else {
		goto L381
	}
L376:
	;
	v1256 = F__emscripten_memcpy_bulkmem(m, v1245, v1192, v1255)
	mBase = m.M
	v1257 = v1256
	goto L378
L377:
	;
	v1257 = v1245
	goto L378
L378:
	;
	goto L375
L379:
	;
	if v1184+int32(1664) != v1192 {
		goto L383
	} else {
		goto L384
	}
L380:
	;
	v1267 = F__emscripten_memcpy_bulkmem(m, v1264, v1203, v1266)
	mBase = m.M
	v1268 = v1267
	goto L382
L381:
	;
	v1268 = v1264
	goto L382
L382:
	;
	goto L379
L383:
	;
	F_pfree(m, v1192)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1274 = int32(1)
	v1277 = v1257 + v1253<<(uint(v1274)%32)
	if base.Ui32(v1257+v1240<<(uint(v1274)%32)) <= base.Ui32(v1277) {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	goto L385
L387:
	;
	v1339 = v1257
	v1342 = v1274
	goto L365
L388:
	;
	goto L389
L389:
	;
	v1287 = v1257
	v1291 = v1268 + v1266 - int32(8)
	v1292 = v1268
	v1293 = v1277 - int32(2)
	v1294 = v1240
	goto L368
L390:
	;
	if v1205 == int32(-2) {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v1319 = v1318 + v1297
	if base.Ui32(int32(80)) < base.Ui32(v1319) {
		v1358 = v1317
		goto L364
	} else {
		goto L400
	}
L392:
	;
	v1304 = F_replication_yylex(m, v1184+int32(2072), v1151)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L395
	}
L393:
	;
	v1306 = v1205
	goto L394
L394:
	;
	if v1306 <= int32(0) {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	v1306 = v1304
	goto L394
L396:
	;
	v1309 = int32(0)
	v1317 = v1309
	v1318 = v1309
	goto L391
L397:
	;
	goto L398
L398:
	;
	if base.Ui32(int32(282)) < base.Ui32(v1306) {
		v1317 = v1306
		v1318 = int32(2)
		goto L391
	} else {
		goto L399
	}
L399:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306)+uint32(_c_F_PostgresMainLoopOnce[56]))))
	v1317 = v1306
	v1318 = v1316
	goto L391
L400:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319)+uint32(_c_F_PostgresMainLoopOnce[57]))))
	if v1318 != v1324 {
		v1358 = v1317
		goto L364
	} else {
		goto L401
	}
L401:
	;
	if v1319 != int32(57) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1329 = v1291 + int32(8)
	v1330 = *(*int64)(unsafe.Add(mBase, uint32(v1184)+2072))
	*(*int64)(unsafe.Add(mBase, uint32(v1329))) = v1330
	if v1317 != 0 {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	goto L404
L404:
	;
	v1339 = v1287
	v1342 = int32(0)
	goto L365
L405:
	;
	v1334 = int32(-2)
	goto L407
L406:
	;
	v1334 = int32(0)
	goto L407
L407:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319)+uint32(_c_F_PostgresMainLoopOnce[58]))))
	v1824 = v1329
	v1825 = v1293
	v1826 = v1334
	v1829 = v1337
	goto L363
L408:
	;
	F_pfree(m, v1339)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	m.G0 = v1184 + int32(2080)
	goto L356
L411:
	;
	goto L410
L412:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1361)+uint32(_c_F_PostgresMainLoopOnce[59]))))
	v1371 = v1291 + (int32(1)-v1367)<<(uint(int32(3))%32)
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1371)))
	v1374 = int32(base.Ui32(v1372) >> (uint(int32(8)) % 32))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+4))
	switch v1361 - int32(2) {
	case 0:
		goto L475
	default:
		v1784 = v1372
		v1785 = v1374
		goto L413
	case 14:
		goto L474
	case 15:
		goto L473
	case 16:
		goto L472
	case 17:
		goto L471
	case 18:
		goto L470
	case 19:
		goto L469
	case 20:
		goto L468
	case 21:
		goto L467
	case 22:
		goto L466
	case 23:
		goto L465
	case 24:
		goto L464
	case 25:
		goto L463
	case 26, 44, 46, 48, 53:
		goto L462
	case 27:
		goto L461
	case 28:
		goto L460
	case 29:
		goto L459
	case 30:
		goto L458
	case 31:
		goto L457
	case 32:
		goto L456
	case 33:
		goto L455
	case 34:
		goto L454
	case 35:
		goto L453
	case 36:
		goto L452
	case 37:
		goto L451
	case 38:
		goto L450
	case 41:
		goto L449
	case 42:
		goto L448
	case 43:
		goto L447
	case 45:
		goto L446
	case 47:
		goto L445
	case 49:
		goto L444
	case 50:
		goto L443
	case 51:
		goto L442
	case 52:
		goto L441
	case 54:
		goto L440
	case 55:
		goto L439
	case 56:
		goto L438
	case 57:
		goto L437
	case 58:
		goto L436
	case 59:
		goto L435
	case 60:
		goto L434
	case 61:
		goto L433
	case 62:
		goto L432
	case 63:
		goto L431
	case 64:
		goto L430
	case 65:
		goto L429
	case 66:
		goto L428
	case 67:
		goto L427
	case 68:
		goto L426
	case 69:
		goto L425
	case 70:
		goto L424
	case 71:
		goto L423
	case 72:
		goto L422
	case 73:
		goto L421
	case 74:
		goto L420
	case 75:
		goto L419
	case 76:
		goto L418
	case 77:
		goto L417
	case 78:
		goto L416
	case 79:
		goto L415
	case 80:
		goto L414
	}
L413:
	;
	v1788 = v1291 - v1367<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1788)+12)) = v1375
	v1790 = int32(8)
	v1791 = v1788 + v1790
	*(*int32)(unsafe.Add(mBase, uint32(v1791))) = v1784&int32(255) | v1785<<(uint(v1790)%32)
	v1800 = v1293 - v1367<<(uint(int32(1))%32)
	v1801 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1800))))
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1361)+uint32(_c_F_PostgresMainLoopOnce[60]))))
	v1807 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1804)+uint32(_c_F_PostgresMainLoopOnce[61]))))
	v1808 = v1801 + v1807
	if base.Ui32(int32(80)) < base.Ui32(v1808) {
		goto L517
	} else {
		goto L518
	}
L414:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_34)
	v1785 = int32(302)
	goto L413
L415:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_35)
	v1785 = int32(337)
	goto L413
L416:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_36)
	v1785 = int32(337)
	goto L413
L417:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_37)
	v1785 = int32(337)
	goto L413
L418:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_38)
	v1785 = int32(1408)
	goto L413
L419:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_39)
	v1785 = int32(68)
	goto L413
L420:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_40)
	v1785 = int32(1201)
	goto L413
L421:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_41)
	v1785 = int32(336)
	goto L413
L422:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_42)
	v1785 = int32(1224)
	goto L413
L423:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_43)
	v1785 = int32(1224)
	goto L413
L424:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_44)
	v1785 = int32(1455)
	goto L413
L425:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_45)
	v1785 = int32(404)
	goto L413
L426:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_46)
	v1785 = int32(50)
	goto L413
L427:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_47)
	v1785 = int32(331)
	goto L413
L428:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_48)
	v1785 = int32(331)
	goto L413
L429:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_49)
	v1785 = int32(331)
	goto L413
L430:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_50)
	v1785 = int32(1037)
	goto L413
L431:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_51)
	v1785 = int32(124)
	goto L413
L432:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_52)
	v1785 = int32(1127)
	goto L413
L433:
	;
	v1784 = int32(_a_F_PostgresMainLoopOnce_53)
	v1785 = int32(909)
	goto L413
L434:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1784 = v1741
	v1785 = int32(base.Ui32(v1741) >> (uint(int32(8)) % 32))
	goto L413
L435:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1734 = F_makeInteger(m, v1733)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L515
	}
L436:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1723 = F_makeString(m, v1722)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L1
	} else {
		goto L513
	}
L437:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1712 = F_makeString(m, v1711)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L511
	}
L438:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1704 = F_makeDefElem(m, v1701, int32(0), int32(-1))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L1
	} else {
		goto L510
	}
L439:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+52)) = v1691
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+56)) = v1691
	v1697 = F_list_make1_impl(m, int32(1), v1184+int32(52))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L1
	} else {
		goto L509
	}
L440:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(16))))
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1687 = F_lappend(m, v1685, v1686)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L508
	}
L441:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1679 = F_makeString(m, v1678)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L507
	}
L442:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1674 = F_makeDefElem(m, v1671, v1672, int32(-1))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L506
	}
L443:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(16))))
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1665 = F_lappend(m, v1663, v1664)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L1
	} else {
		goto L505
	}
L444:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+48)) = v1651
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+60)) = v1651
	v1657 = F_list_make1_impl(m, int32(1), v1184+int32(48))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L504
	}
L445:
	;
	v1646 = int32(8)
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1646)))
	v1784 = v1648
	v1785 = int32(base.Ui32(v1648) >> (uint(v1646) % 32))
	goto L413
L446:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	if v1641 == int32(0) {
		goto L357
	} else {
		goto L503
	}
L447:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1784 = v1638
	v1785 = int32(base.Ui32(v1638) >> (uint(int32(8)) % 32))
	goto L413
L448:
	;
	v1784 = int32(0)
	v1785 = v1374
	goto L413
L449:
	;
	v1784 = int32(1)
	v1785 = v1374
	goto L413
L450:
	;
	v1630 = F_palloc0(m, int32(4))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L1
	} else {
		goto L502
	}
L451:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	if v1617 == int32(0) {
		goto L358
	} else {
		goto L500
	}
L452:
	;
	v1601 = F_palloc0(m, int32(32))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L499
	}
L453:
	;
	v1584 = F_palloc0(m, int32(32))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L1
	} else {
		goto L498
	}
L454:
	;
	v1569 = F_palloc0(m, int32(12))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L497
	}
L455:
	;
	v1556 = F_palloc0(m, int32(12))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L496
	}
L456:
	;
	v1545 = F_palloc0(m, int32(12))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L495
	}
L457:
	;
	v1537 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L493
	}
L458:
	;
	v1528 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L491
	}
L459:
	;
	v1519 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_54))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L489
	}
L460:
	;
	v1510 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_55))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L487
	}
L461:
	;
	v1501 = F_makeString(m, int32(_a_F_PostgresMainLoopOnce_56))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L485
	}
L462:
	;
	v1497 = int32(0)
	v1784 = v1497
	v1785 = v1497
	goto L413
L463:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1493 = F_lappend(m, v1491, v1492)
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L484
	}
L464:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1784 = v1486
	v1785 = int32(base.Ui32(v1486) >> (uint(int32(8)) % 32))
	goto L413
L465:
	;
	v1481 = int32(8)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1481)))
	v1784 = v1483
	v1785 = int32(base.Ui32(v1483) >> (uint(v1481) % 32))
	goto L413
L466:
	;
	v1459 = F_palloc0(m, int32(24))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L483
	}
L467:
	;
	v1440 = F_palloc0(m, int32(24))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L482
	}
L468:
	;
	v1433 = F_palloc0(m, int32(8))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L481
	}
L469:
	;
	v1422 = F_palloc0(m, int32(8))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L480
	}
L470:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(16))))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+4)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v1184))) = v1412
	v1417 = F_psprintf(m, int32(_a_F_PostgresMainLoopOnce_57), v1184)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L479
	}
L471:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1784 = v1407
	v1785 = int32(base.Ui32(v1407) >> (uint(int32(8)) % 32))
	goto L413
L472:
	;
	v1399 = F_palloc0(m, int32(8))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L478
	}
L473:
	;
	v1390 = F_palloc0(m, int32(8))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L477
	}
L474:
	;
	v1383 = F_palloc0(m, int32(4))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L476
	}
L475:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v685+int32(424)))) = v1380
	v1784 = v1372
	v1785 = v1374
	goto L413
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1383))) = int32(448)
	v1784 = v1383
	v1785 = int32(base.Ui32(v1383) >> (uint(int32(8)) % 32))
	goto L413
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1390))) = int32(454)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1390)+4)) = v1394
	v1784 = v1390
	v1785 = int32(base.Ui32(v1390) >> (uint(int32(8)) % 32))
	goto L413
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = int32(159)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+4)) = v1403
	v1784 = v1399
	v1785 = int32(base.Ui32(v1399) >> (uint(int32(8)) % 32))
	goto L413
L479:
	;
	v1784 = v1417
	v1785 = int32(base.Ui32(v1417) >> (uint(int32(8)) % 32))
	goto L413
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1422))) = int32(449)
	v1426 = int32(8)
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1426)))
	*(*int32)(unsafe.Add(mBase, uint32(v1422)+4)) = v1428
	v1784 = v1422
	v1785 = int32(base.Ui32(v1422) >> (uint(v1426) % 32))
	goto L413
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1433))) = int32(449)
	v1784 = v1433
	v1785 = int32(base.Ui32(v1433) >> (uint(int32(8)) % 32))
	goto L413
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = int32(450)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+4)) = v1448
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1440)+16)) = uint8(v1452)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+20)) = v1454
	v1784 = v1440
	v1785 = int32(base.Ui32(v1440) >> (uint(int32(8)) % 32))
	goto L413
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1459))) = int32(450)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+4)) = v1467
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291-int32(24)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1459)+16)) = uint8(v1471)
	v1473 = int32(8)
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1473)))
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+12)) = v1475
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1459)+20)) = v1477
	v1784 = v1459
	v1785 = int32(base.Ui32(v1459) >> (uint(v1473) % 32))
	goto L413
L484:
	;
	v1784 = v1493
	v1785 = int32(base.Ui32(v1493) >> (uint(int32(8)) % 32))
	goto L413
L485:
	;
	v1504 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_58), v1501, int32(-1))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	v1784 = v1504
	v1785 = int32(base.Ui32(v1504) >> (uint(int32(8)) % 32))
	goto L413
L487:
	;
	v1513 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_58), v1510, int32(-1))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	v1784 = v1513
	v1785 = int32(base.Ui32(v1513) >> (uint(int32(8)) % 32))
	goto L413
L489:
	;
	v1522 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_58), v1519, int32(-1))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v1784 = v1522
	v1785 = int32(base.Ui32(v1522) >> (uint(int32(8)) % 32))
	goto L413
L491:
	;
	v1531 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_40), v1528, int32(-1))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v1784 = v1531
	v1785 = int32(base.Ui32(v1531) >> (uint(int32(8)) % 32))
	goto L413
L493:
	;
	v1540 = F_makeDefElem(m, int32(_a_F_PostgresMainLoopOnce_38), v1537, int32(-1))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	v1784 = v1540
	v1785 = int32(base.Ui32(v1540) >> (uint(int32(8)) % 32))
	goto L413
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1545))) = int32(451)
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1550 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1545)+8)) = uint8(v1550)
	*(*int32)(unsafe.Add(mBase, uint32(v1545)+4)) = v1549
	v1784 = v1545
	v1785 = int32(base.Ui32(v1545) >> (uint(int32(8)) % 32))
	goto L413
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1556))) = int32(451)
	v1560 = int32(8)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1560)))
	v1563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1556)+8)) = uint8(v1563)
	*(*int32)(unsafe.Add(mBase, uint32(v1556)+4)) = v1562
	v1784 = v1556
	v1785 = int32(base.Ui32(v1556) >> (uint(v1560) % 32))
	goto L413
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1569))) = int32(452)
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+4)) = v1575
	v1577 = int32(8)
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1291-v1577)))
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+8)) = v1579
	v1784 = v1569
	v1785 = int32(base.Ui32(v1569) >> (uint(v1577) % 32))
	goto L413
L498:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1584))) = int64(453)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1584)+8)) = v1590
	v1592 = int32(8)
	v1594 = *(*int64)(unsafe.Add(mBase, uint32(v1291-v1592)))
	*(*int64)(unsafe.Add(mBase, uint32(v1584)+16)) = v1594
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1584)+12)) = v1596
	v1784 = v1584
	v1785 = int32(base.Ui32(v1584) >> (uint(v1592) % 32))
	goto L413
L499:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1601))) = int64(4294967749)
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1291-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+8)) = v1607
	v1609 = int32(8)
	v1611 = *(*int64)(unsafe.Add(mBase, uint32(v1291-v1609)))
	*(*int64)(unsafe.Add(mBase, uint32(v1601)+16)) = v1611
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+24)) = v1613
	v1784 = v1601
	v1785 = int32(base.Ui32(v1601) >> (uint(v1609) % 32))
	goto L413
L500:
	;
	v1621 = F_palloc0(m, int32(8))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1621))) = int32(455)
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1621)+4)) = v1625
	v1784 = v1621
	v1785 = int32(base.Ui32(v1621) >> (uint(int32(8)) % 32))
	goto L413
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1630))) = int32(456)
	v1784 = v1630
	v1785 = int32(base.Ui32(v1630) >> (uint(int32(8)) % 32))
	goto L413
L503:
	;
	v1784 = v1641
	v1785 = int32(base.Ui32(v1641) >> (uint(int32(8)) % 32))
	goto L413
L504:
	;
	v1784 = v1657
	v1785 = int32(base.Ui32(v1657) >> (uint(int32(8)) % 32))
	goto L413
L505:
	;
	v1784 = v1665
	v1785 = int32(base.Ui32(v1665) >> (uint(int32(8)) % 32))
	goto L413
L506:
	;
	v1784 = v1674
	v1785 = int32(base.Ui32(v1674) >> (uint(int32(8)) % 32))
	goto L413
L507:
	;
	v1784 = v1679
	v1785 = int32(base.Ui32(v1679) >> (uint(int32(8)) % 32))
	goto L413
L508:
	;
	v1784 = v1687
	v1785 = int32(base.Ui32(v1687) >> (uint(int32(8)) % 32))
	goto L413
L509:
	;
	v1784 = v1697
	v1785 = int32(base.Ui32(v1697) >> (uint(int32(8)) % 32))
	goto L413
L510:
	;
	v1784 = v1704
	v1785 = int32(base.Ui32(v1704) >> (uint(int32(8)) % 32))
	goto L413
L511:
	;
	v1715 = F_makeDefElem(m, v1710, v1712, int32(-1))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1784 = v1715
	v1785 = int32(base.Ui32(v1715) >> (uint(int32(8)) % 32))
	goto L413
L513:
	;
	v1726 = F_makeDefElem(m, v1721, v1723, int32(-1))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	v1784 = v1726
	v1785 = int32(base.Ui32(v1726) >> (uint(int32(8)) % 32))
	goto L413
L515:
	;
	v1737 = F_makeDefElem(m, v1732, v1734, int32(-1))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	v1784 = v1737
	v1785 = int32(base.Ui32(v1737) >> (uint(int32(8)) % 32))
	goto L413
L517:
	;
	v1820 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1804)+uint32(_c_F_PostgresMainLoopOnce[62]))))
	v1824 = v1791
	v1825 = v1800
	v1826 = v1358
	v1829 = v1820
	goto L363
L518:
	;
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1808)+uint32(_c_F_PostgresMainLoopOnce[57]))))
	if v1813 != v1801 {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v1817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1808)+uint32(_c_F_PostgresMainLoopOnce[58]))))
	v1824 = v1791
	v1825 = v1800
	v1826 = v1358
	v1829 = v1817
	goto L363
L520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+16)) = v1847
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_59), v1184+int32(16))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_60), int32(322), int32(_a_F_PostgresMainLoopOnce_61))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L526:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+32)) = v1866
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_59), v1184+int32(32))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_60), int32(363), int32(_a_F_PostgresMainLoopOnce_61))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v685)+428))
	F_replication_scanner_finish(m, v1878)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v673
	F_pgstat_report_activity(m, int32(3), v673)
	mBase = m.M
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[63])))
	if v1888 != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1889 = int32(15)
	goto L534
L533:
	;
	v1889 = int32(14)
	goto L534
L534:
	;
	v1891 = F_errstart(m, v1889, int32(0))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	if v1891 != 0 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+400)) = v673
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_62), v685+int32(400))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L1
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1905)+24))
	goto L541
L539:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2095), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	if (v1906-int32(7))&int32(-9) == int32(0) {
		goto L265
	} else {
		goto L542
	}
L542:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v1914 != 0 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L1
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_63))
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L1
	} else {
		goto L547
	}
L546:
	;
	goto L545
L547:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_64))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	F_initStringInfo(m, int32(_a_F_PostgresMainLoopOnce_65))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)))
	switch v1927 - int32(448) {
	case 0:
		goto L559
	case 1:
		goto L550
	case 2:
		goto L557
	case 3:
		goto L556
	case 4:
		goto L555
	case 5:
		goto L554
	case 6:
		goto L558
	case 7:
		goto L553
	case 8:
		goto L552
	default:
		goto L551
	}
L550:
	;
	v4435 = int32(_a_F_PostgresMainLoopOnce_66)
	F_PreventInTransactionBlock(m, int32(1), v4435)
	mBase = m.M
	v4439 = m.ExcPending
	if v4439 != 0 {
		goto L1
	} else {
		goto L1182
	}
L551:
	;
	if v1927 == int32(159) {
		goto L246
	} else {
		goto L1178
	}
L552:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_67))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L1042
	}
L553:
	;
	F_PreventInTransactionBlock(m, int32(1), int32(_a_F_PostgresMainLoopOnce_68))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L1
	} else {
		goto L1013
	}
L554:
	;
	v2967 = int32(_a_F_PostgresMainLoopOnce_69)
	F_PreventInTransactionBlock(m, int32(1), v2967)
	mBase = m.M
	v2971 = m.ExcPending
	if v2971 != 0 {
		goto L1
	} else {
		goto L854
	}
L555:
	;
	v2576 = int32(0)
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+8))
	if v2578 == v2576 {
		v2733 = v1
		v2734 = v1
		v2739 = v1
		v2748 = v2576
		goto L754
	} else {
		goto L755
	}
L556:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+4))
	v2568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2566)+8)))
	F_ReplicationSlotDrop(m, v2567, (v2568^int32(-1))&int32(1))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L1
	} else {
		goto L753
	}
L557:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v2253 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = v2253
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+20))
	if v2255 == v2253 {
		v4843 = v1
		v4853 = v1
		v4854 = v1
		v4856 = v1
		goto L245
	} else {
		goto L656
	}
L558:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[65]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = int64(0)
	v2095 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L607
	}
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = int32(0)
	v1933 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[66]))
	v1934 = *(*int64)(unsafe.Add(mBase, uint32(v1933)))
	goto L560
L560:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v685)+32)) = v1934
	v1938 = int32(32)
	v1942 = F_pg_snprintf(m, v685+int32(_a_F_PostgresMainLoopOnce_70), v1938, int32(_a_F_PostgresMainLoopOnce_71), v685+v1938)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v1947 == int32(1) {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])) = uint8(v1957)
	if v1957 != 0 {
		goto L567
	} else {
		goto L568
	}
L563:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1952)+316))
	v1955 = base.B2i32(v1953 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v1955)
	v1957 = v1955
	goto L565
L564:
	;
	v1957 = int32(0)
	goto L565
L565:
	;
	goto L562
L566:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+20)) = uint32(v2000)
	v2003 = int64(base.Ui64(v2000) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+16)) = uint32(v2003)
	v2011 = F_pg_snprintf(m, v685+int32(464), int32(64), int32(_a_F_PostgresMainLoopOnce_72), v685+int32(16))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L582
	}
L567:
	;
	v1962 = F_GetWalRcvFlushRecPtr(m, int32(0), v685+int32(_a_F_PostgresMainLoopOnce_73))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	v1976 = v685 + int32(440)
	v1979 = int32(_a_F_PostgresMainLoopOnce_74)
	v1980 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1981 = *(*int64)(unsafe.Add(mBase, uint32(v1980)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v1980)+280)) = v1981
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = v1981
	v1986 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1987 = *(*int64)(unsafe.Add(mBase, uint32(v1986)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v1986)+272)) = v1987
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = v1987
	if v1976 != 0 {
		goto L579
	} else {
		goto L580
	}
L570:
	;
	v1966 = F_GetXLogReplayRecPtr(m, v685+int32(464))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+440)) = v1968
	if base.Ui64(v1966) < base.Ui64(v1962) {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v1971 = v1962
	goto L574
L573:
	;
	v1971 = v1966
	goto L574
L574:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72])))
	if v1968 == v1972 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v1974 = v1971
	goto L577
L576:
	;
	v1974 = v1966
	goto L577
L577:
	;
	v2000 = v1974
	goto L566
L578:
	;
	v2000 = v1996
	goto L566
L579:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v1976))) = v1993
	goto L581
L580:
	;
	goto L581
L581:
	;
	v1996 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70]))
	goto L578
L582:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	if v2015 != 0 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L586
	}
L584:
	;
	v2029 = int32(0)
	goto L585
L585:
	;
	v2031 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L590
	}
L586:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[53]))
	v2022 = F_get_database_name(m, v2021)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	v2024 = F_MemoryContextStrdup(m, v2017, v2022)
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v2029 = v2024
	goto L585
L590:
	;
	v2034 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	F_TupleDescInitBuiltinEntry(m, v2034, int32(1), int32(_a_F_PostgresMainLoopOnce_75), int32(25))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	F_TupleDescInitBuiltinEntry(m, v2034, int32(2), int32(_a_F_PostgresMainLoopOnce_44), int32(20))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	F_TupleDescInitBuiltinEntry(m, v2034, int32(3), int32(_a_F_PostgresMainLoopOnce_76), int32(25))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	F_TupleDescInitBuiltinEntry(m, v2034, int32(4), int32(_a_F_PostgresMainLoopOnce_77), int32(25))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v2057 = F_begin_tup_output_tupdesc(m, v2031, v2034, int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	v2061 = F_cstring_to_text(m, v685+int32(_a_F_PostgresMainLoopOnce_70))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = v2061
	v2064 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v685)+440)))
	v2065 = F_Int64GetDatum(m, v2064)
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = v2065
	v2070 = F_cstring_to_text(m, v685+int32(464))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = v2070
	if v2029 != 0 {
		goto L601
	} else {
		goto L602
	}
L600:
	;
	F_do_tup_output(m, v2057, v685+int32(_a_F_PostgresMainLoopOnce_73), v685+int32(_a_F_PostgresMainLoopOnce_79))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L1
	} else {
		goto L605
	}
L601:
	;
	v2073 = F_cstring_to_text(m, v2029)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v2076 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[75]))) = uint8(v2076)
	goto L600
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[76]))) = v2073
	goto L600
L605:
	;
	F_end_tup_output(m, v2057)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_80)
	goto L244
L607:
	;
	F_TupleDescInitBuiltinEntry(m, v2095, int32(1), int32(_a_F_PostgresMainLoopOnce_81), int32(25))
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	F_TupleDescInitBuiltinEntry(m, v2095, int32(2), int32(_a_F_PostgresMainLoopOnce_82), int32(25))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	F_TupleDescInitBuiltinEntry(m, v2095, int32(3), int32(_a_F_PostgresMainLoopOnce_83), int32(20))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v2112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+462)) = uint8(v2112)
	v2114 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+460)) = uint16(v2114)
	v2117 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	v2121 = F_LWLockAcquire(m, v2117+int32(_a_F_PostgresMainLoopOnce_84), v2112)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2087)+4))
	v2125 = F_SearchNamedReplicationSlot(m, v2123, int32(0))
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L1
	} else {
		goto L614
	}
L612:
	;
	v2238 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L1
	} else {
		goto L652
	}
L613:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2125)))
	*(*int32)(unsafe.Add(mBase, uint32(v2125))) = int32(1)
	if v2134 != 0 {
		goto L620
	} else {
		goto L621
	}
L614:
	;
	if v2125 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125)+4)))
	if v2127 != 0 {
		goto L613
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	F_LWLockRelease(m, v2129+int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L1
	} else {
		goto L619
	}
L618:
	;
	goto L617
L619:
	;
	goto L612
L620:
	;
	F_s_lock(m, v2125, int32(_a_F_PostgresMainLoopOnce_14), int32(511), int32(_a_F_PostgresMainLoopOnce_85))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L1
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	goto L625
L623:
	;
	goto L622
L624:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2125)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+456)) = v2147
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2125)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+448)) = v2149
	v2151 = *(*int64)(unsafe.Add(mBase, uint32(v2125)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v685)+440)) = v2151
	v2153 = *(*int64)(unsafe.Add(mBase, uint32(v2125)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v685)+432)) = v2153
	goto L629
L625:
	;
	v2145 = F__emscripten_memcpy_bulkmem(m, v685+int32(_a_F_PostgresMainLoopOnce_70), v2125, int32(88))
	mBase = m.M
	goto L627
L627:
	;
	goto L624
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2125))) = int32(0)
	v2165 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[77]))
	F_LWLockRelease(m, v2165+int32(_a_F_PostgresMainLoopOnce_84))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L1
	} else {
		goto L632
	}
L629:
	;
	v2160 = F__emscripten_memcpy_bulkmem(m, v685+int32(464), v2125+int32(112), int32(176))
	mBase = m.M
	goto L631
L631:
	;
	goto L628
L632:
	;
	if v2147 != 0 {
		goto L264
	} else {
		goto L633
	}
L633:
	;
	v2171 = F_cstring_to_text(m, int32(_a_F_PostgresMainLoopOnce_43))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	v2173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+460)) = uint8(v2173)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64]))) = v2171
	if v2153 == int64(0) {
		goto L612
	} else {
		goto L635
	}
L635:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+52)) = uint32(v2153)
	v2180 = int64(base.Ui64(v2153) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+48)) = uint32(v2180)
	v2188 = F_pg_snprintf(m, v685+int32(_a_F_PostgresMainLoopOnce_73), int32(64), int32(_a_F_PostgresMainLoopOnce_72), v685+int32(48))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	v2192 = F_cstring_to_text(m, v685+int32(_a_F_PostgresMainLoopOnce_73))
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	v2194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+461)) = uint8(v2194)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v2192
	v2197 = *(*int64)(unsafe.Add(mBase, uint32(v685)+432))
	if v2197 == int64(0) {
		goto L612
	} else {
		goto L638
	}
L638:
	;
	v2202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v2202 == int32(1) {
		goto L641
	} else {
		goto L642
	}
L639:
	;
	v2223 = F_readTimeLineHistory(m, v2222)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L1
	} else {
		goto L649
	}
L640:
	;
	if v2212 != 0 {
		goto L644
	} else {
		goto L645
	}
L641:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+316))
	v2210 = base.B2i32(v2208 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v2210)
	v2212 = v2210
	goto L643
L642:
	;
	v2212 = int32(0)
	goto L643
L643:
	;
	goto L640
L644:
	;
	v2215 = F_GetXLogReplayRecPtr(m, v685+int32(_a_F_PostgresMainLoopOnce_73))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2219)+308))
	goto L648
L647:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72])))
	v2222 = v2217
	goto L639
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = v2220
	v2222 = v2220
	goto L639
L649:
	;
	v2225 = *(*int64)(unsafe.Add(mBase, uint32(v685)+432))
	v2226 = F_tliOfPointInHistory(m, v2225, v2223)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	v2229 = F_Int64GetDatum(m, base.I64_extend_i32_u(v2226))
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v2231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+462)) = uint8(v2231)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[65]))) = v2229
	goto L612
L652:
	;
	v2241 = F_begin_tup_output_tupdesc(m, v2238, v2095, int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	F_do_tup_output(m, v2241, v685+int32(_a_F_PostgresMainLoopOnce_79), v685+int32(460))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	F_end_tup_output(m, v2241)
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_86)
	goto L244
L656:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2255)+4))
	if v2258 <= int32(0) {
		v4843 = v1
		v4853 = v1
		v4854 = v1
		v4856 = v1
		goto L245
	} else {
		goto L657
	}
L657:
	;
	v2262 = int32(0)
	v2277 = v1
	v2278 = v1
	v2283 = v1
	v2284 = v1
	v2287 = v1
	v2288 = v1
	v2289 = v1
	v2290 = v1
	goto L658
L658:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v2255)+12))
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2298+v2262<<(uint(int32(2))%32))))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+8))
	v2304 = int32(_a_F_PostgresMainLoopOnce_58)
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[79])))
	v2308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303))))
	if v2308 == int32(0) {
		v2327 = v2307
		v2328 = v2308
		goto L662
	} else {
		goto L663
	}
L659:
	;
	v4843 = v2553
	v4853 = v2557
	v4854 = v2558
	v4856 = v2560
	goto L245
L660:
	;
	v2562 = v2262 + int32(1)
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2255)+4))
	if v2562 < v2563 {
		v2262 = v2562
		v2277 = v2553
		v2278 = v2554
		v2283 = v2555
		v2284 = v2556
		v2287 = v2557
		v2288 = v2558
		v2289 = v2559
		v2290 = v2560
		goto L658
	} else {
		goto L752
	}
L661:
	;
	if v2328-v2327 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L662:
	;
	goto L661
L663:
	;
	if v2307 != v2308 {
		v2327 = v2307
		v2328 = v2308
		goto L662
	} else {
		goto L664
	}
L664:
	;
	v2312 = v2303
	v2313 = v2304
	goto L665
L665:
	;
	v2316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2313)+1)))
	v2317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2312)+1)))
	if v2317 == int32(0) {
		v2327 = v2316
		v2328 = v2317
		goto L662
	} else {
		goto L667
	}
L666:
	;
	v2327 = v2316
	v2328 = v2317
	goto L662
L667:
	;
	v2320 = int32(1)
	if v2316 == v2317 {
		v2312 = v2312 + v2320
		v2313 = v2313 + v2320
		goto L665
	} else {
		goto L668
	}
L668:
	;
	goto L666
L669:
	;
	if v2283&int32(1) != 0 {
		goto L263
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	v2450 = int32(_a_F_PostgresMainLoopOnce_40)
	v2453 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[80])))
	v2454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303))))
	if v2454 == int32(0) {
		v2473 = v2453
		v2474 = v2454
		goto L713
	} else {
		goto L714
	}
L672:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+8))
	if v2334 != int32(1) {
		goto L263
	} else {
		goto L673
	}
L673:
	;
	v2337 = F_defGetString(m, v2302)
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	v2339 = int32(_a_F_PostgresMainLoopOnce_56)
	v2342 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[81])))
	v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337))))
	if v2343 == int32(0) {
		v2362 = v2342
		v2363 = v2343
		goto L676
	} else {
		goto L677
	}
L675:
	;
	if v2363-v2362 == int32(0) {
		goto L683
	} else {
		goto L684
	}
L676:
	;
	goto L675
L677:
	;
	if v2342 != v2343 {
		v2362 = v2342
		v2363 = v2343
		goto L676
	} else {
		goto L678
	}
L678:
	;
	v2347 = v2337
	v2348 = v2339
	goto L679
L679:
	;
	v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2348)+1)))
	v2352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2347)+1)))
	if v2352 == int32(0) {
		v2362 = v2351
		v2363 = v2352
		goto L676
	} else {
		goto L681
	}
L680:
	;
	v2362 = v2351
	v2363 = v2352
	goto L676
L681:
	;
	v2355 = int32(1)
	if v2351 == v2352 {
		v2347 = v2347 + v2355
		v2348 = v2348 + v2355
		goto L679
	} else {
		goto L682
	}
L682:
	;
	goto L680
L683:
	;
	v2553 = v2277
	v2554 = v2278
	v2555 = int32(1)
	v2556 = v2284
	v2557 = int32(0)
	v2558 = v2288
	v2559 = v2289
	v2560 = v2290
	goto L660
L684:
	;
	goto L685
L685:
	;
	v2369 = int32(1)
	v2370 = int32(_a_F_PostgresMainLoopOnce_55)
	v2373 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[82])))
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337))))
	if v2374 == int32(0) {
		v2393 = v2373
		v2394 = v2374
		goto L687
	} else {
		goto L688
	}
L686:
	;
	if v2394-v2393 == int32(0) {
		goto L694
	} else {
		goto L695
	}
L687:
	;
	goto L686
L688:
	;
	if v2373 != v2374 {
		v2393 = v2373
		v2394 = v2374
		goto L687
	} else {
		goto L689
	}
L689:
	;
	v2378 = v2337
	v2379 = v2370
	goto L690
L690:
	;
	v2382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2379)+1)))
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2378)+1)))
	if v2383 == int32(0) {
		v2393 = v2382
		v2394 = v2383
		goto L687
	} else {
		goto L692
	}
L691:
	;
	v2393 = v2382
	v2394 = v2383
	goto L687
L692:
	;
	v2386 = int32(1)
	if v2382 == v2383 {
		v2378 = v2378 + v2386
		v2379 = v2379 + v2386
		goto L690
	} else {
		goto L693
	}
L693:
	;
	goto L691
L694:
	;
	v2553 = v2277
	v2554 = v2278
	v2555 = v2369
	v2556 = v2284
	v2557 = int32(1)
	v2558 = v2288
	v2559 = v2289
	v2560 = v2290
	goto L660
L695:
	;
	goto L696
L696:
	;
	v2399 = int32(_a_F_PostgresMainLoopOnce_54)
	v2402 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[83])))
	v2403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337))))
	if v2403 == int32(0) {
		v2422 = v2402
		v2423 = v2403
		goto L698
	} else {
		goto L699
	}
L697:
	;
	if v2423-v2422 == int32(0) {
		goto L705
	} else {
		goto L706
	}
L698:
	;
	goto L697
L699:
	;
	if v2402 != v2403 {
		v2422 = v2402
		v2423 = v2403
		goto L698
	} else {
		goto L700
	}
L700:
	;
	v2407 = v2337
	v2408 = v2399
	goto L701
L701:
	;
	v2411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2408)+1)))
	v2412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2407)+1)))
	if v2412 == int32(0) {
		v2422 = v2411
		v2423 = v2412
		goto L698
	} else {
		goto L703
	}
L702:
	;
	v2422 = v2411
	v2423 = v2412
	goto L698
L703:
	;
	v2415 = int32(1)
	if v2411 == v2412 {
		v2407 = v2407 + v2415
		v2408 = v2408 + v2415
		goto L701
	} else {
		goto L704
	}
L704:
	;
	goto L702
L705:
	;
	v2553 = v2277
	v2554 = v2278
	v2555 = v2369
	v2556 = v2284
	v2557 = int32(2)
	v2558 = v2288
	v2559 = v2289
	v2560 = v2290
	goto L660
L706:
	;
	goto L707
L707:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+200)) = v2337
	*(*int32)(unsafe.Add(mBase, uint32(v685)+196)) = v2435
	*(*int32)(unsafe.Add(mBase, uint32(v685)+192)) = int32(_a_F_PostgresMainLoopOnce_87)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_88), v685+int32(192))
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1152), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	if v2474-v2473 == int32(0) {
		goto L720
	} else {
		goto L721
	}
L713:
	;
	goto L712
L714:
	;
	if v2453 != v2454 {
		v2473 = v2453
		v2474 = v2454
		goto L713
	} else {
		goto L715
	}
L715:
	;
	v2458 = v2303
	v2459 = v2450
	goto L716
L716:
	;
	v2462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2459)+1)))
	v2463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2458)+1)))
	if v2463 == int32(0) {
		v2473 = v2462
		v2474 = v2463
		goto L713
	} else {
		goto L718
	}
L717:
	;
	v2473 = v2462
	v2474 = v2463
	goto L713
L718:
	;
	v2466 = int32(1)
	if v2462 == v2463 {
		v2458 = v2458 + v2466
		v2459 = v2459 + v2466
		goto L716
	} else {
		goto L719
	}
L719:
	;
	goto L717
L720:
	;
	if v2289&int32(1) != 0 {
		goto L262
	} else {
		goto L723
	}
L721:
	;
	goto L722
L722:
	;
	v2484 = int32(_a_F_PostgresMainLoopOnce_38)
	v2487 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[84])))
	v2488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303))))
	if v2488 == int32(0) {
		v2507 = v2487
		v2508 = v2488
		goto L727
	} else {
		goto L728
	}
L723:
	;
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+8))
	if v2480 != 0 {
		goto L262
	} else {
		goto L724
	}
L724:
	;
	v2482 = F_defGetBoolean(m, v2302)
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	v2553 = v2482
	v2554 = v2278
	v2555 = v2283
	v2556 = v2284
	v2557 = v2287
	v2558 = v2288
	v2559 = int32(1)
	v2560 = v2290
	goto L660
L726:
	;
	if v2508-v2507 == int32(0) {
		goto L734
	} else {
		goto L735
	}
L727:
	;
	goto L726
L728:
	;
	if v2487 != v2488 {
		v2507 = v2487
		v2508 = v2488
		goto L727
	} else {
		goto L729
	}
L729:
	;
	v2492 = v2303
	v2493 = v2484
	goto L730
L730:
	;
	v2496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2493)+1)))
	v2497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492)+1)))
	if v2497 == int32(0) {
		v2507 = v2496
		v2508 = v2497
		goto L727
	} else {
		goto L732
	}
L731:
	;
	v2507 = v2496
	v2508 = v2497
	goto L727
L732:
	;
	v2500 = int32(1)
	if v2496 == v2497 {
		v2492 = v2492 + v2500
		v2493 = v2493 + v2500
		goto L730
	} else {
		goto L733
	}
L733:
	;
	goto L731
L734:
	;
	if v2278 != 0 {
		goto L261
	} else {
		goto L737
	}
L735:
	;
	goto L736
L736:
	;
	v2518 = int32(_a_F_PostgresMainLoopOnce_90)
	v2521 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[85])))
	v2522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303))))
	if v2522 == int32(0) {
		v2541 = v2521
		v2542 = v2522
		goto L741
	} else {
		goto L742
	}
L737:
	;
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+8))
	if v2512 != int32(1) {
		goto L261
	} else {
		goto L738
	}
L738:
	;
	v2516 = F_defGetBoolean(m, v2302)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	v2553 = v2277
	v2554 = int32(1)
	v2555 = v2283
	v2556 = v2284
	v2557 = v2287
	v2558 = v2516
	v2559 = v2289
	v2560 = v2290
	goto L660
L740:
	;
	if v2542-v2541 != 0 {
		goto L259
	} else {
		goto L748
	}
L741:
	;
	goto L740
L742:
	;
	if v2521 != v2522 {
		v2541 = v2521
		v2542 = v2522
		goto L741
	} else {
		goto L743
	}
L743:
	;
	v2526 = v2303
	v2527 = v2518
	goto L744
L744:
	;
	v2530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2527)+1)))
	v2531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2526)+1)))
	if v2531 == int32(0) {
		v2541 = v2530
		v2542 = v2531
		goto L741
	} else {
		goto L746
	}
L745:
	;
	v2541 = v2530
	v2542 = v2531
	goto L741
L746:
	;
	v2534 = int32(1)
	if v2530 == v2531 {
		v2526 = v2526 + v2534
		v2527 = v2527 + v2534
		goto L744
	} else {
		goto L747
	}
L747:
	;
	goto L745
L748:
	;
	if v2284&int32(1) != 0 {
		goto L260
	} else {
		goto L749
	}
L749:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+8))
	if v2546 != int32(1) {
		goto L260
	} else {
		goto L750
	}
L750:
	;
	v2550 = F_defGetBoolean(m, v2302)
	mBase = m.M
	v2551 = m.ExcPending
	if v2551 != 0 {
		goto L1
	} else {
		goto L751
	}
L751:
	;
	v2553 = v2277
	v2554 = v2278
	v2555 = v2283
	v2556 = int32(1)
	v2557 = v2287
	v2558 = v2288
	v2559 = v2289
	v2560 = v2550
	goto L660
L752:
	;
	goto L659
L753:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_91)
	goto L244
L754:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+464)) = uint8(v2739)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[86]))) = uint8(v2734)
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+4))
	v2752 = int32(0)
	v2753 = m.G0
	v2755 = v2753 - int32(1072)
	m.G0 = v2755
	F_ReplicationSlotAcquire(m, v2751, v2752, int32(1))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L1
	} else {
		goto L791
	}
L755:
	;
	v2581 = int32(0)
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v2578)+4))
	if v2582 <= v2581 {
		v2733 = v1
		v2734 = v1
		v2739 = v1
		v2748 = v2581
		goto L754
	} else {
		goto L756
	}
L756:
	;
	v2586 = int32(0)
	v2607 = v1
	v2608 = v1
	v2611 = v1
	v2613 = v1
	goto L757
L757:
	;
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2578)+12))
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v2622+v2586<<(uint(int32(2))%32))))
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2626)+8))
	v2628 = int32(_a_F_PostgresMainLoopOnce_90)
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[85])))
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2627))))
	if v2632 == int32(0) {
		v2651 = v2631
		v2652 = v2632
		goto L761
	} else {
		goto L762
	}
L758:
	;
	if v2692&int32(1) != 0 {
		goto L785
	} else {
		goto L786
	}
L759:
	;
	v2697 = v2586 + int32(1)
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2578)+4))
	if v2697 < v2698 {
		v2586 = v2697
		v2607 = v2692
		v2608 = v2693
		v2611 = v2694
		v2613 = v2695
		goto L757
	} else {
		goto L784
	}
L760:
	;
	if v2652-v2651 == int32(0) {
		goto L768
	} else {
		goto L769
	}
L761:
	;
	goto L760
L762:
	;
	if v2631 != v2632 {
		v2651 = v2631
		v2652 = v2632
		goto L761
	} else {
		goto L763
	}
L763:
	;
	v2636 = v2627
	v2637 = v2628
	goto L764
L764:
	;
	v2640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2637)+1)))
	v2641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2636)+1)))
	if v2641 == int32(0) {
		v2651 = v2640
		v2652 = v2641
		goto L761
	} else {
		goto L766
	}
L765:
	;
	v2651 = v2640
	v2652 = v2641
	goto L761
L766:
	;
	v2644 = int32(1)
	if v2640 == v2641 {
		v2636 = v2636 + v2644
		v2637 = v2637 + v2644
		goto L764
	} else {
		goto L767
	}
L767:
	;
	goto L765
L768:
	;
	if v2611&int32(1) != 0 {
		goto L258
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	v2661 = int32(_a_F_PostgresMainLoopOnce_38)
	v2664 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[84])))
	v2665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2627))))
	if v2665 == int32(0) {
		v2684 = v2664
		v2685 = v2665
		goto L774
	} else {
		goto L775
	}
L771:
	;
	v2659 = F_defGetBoolean(m, v2626)
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	v2692 = v2607
	v2693 = v2608
	v2694 = int32(1)
	v2695 = v2659
	goto L759
L773:
	;
	if v2685-v2684 != 0 {
		goto L256
	} else {
		goto L781
	}
L774:
	;
	goto L773
L775:
	;
	if v2664 != v2665 {
		v2684 = v2664
		v2685 = v2665
		goto L774
	} else {
		goto L776
	}
L776:
	;
	v2669 = v2627
	v2670 = v2661
	goto L777
L777:
	;
	v2673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2670)+1)))
	v2674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2669)+1)))
	if v2674 == int32(0) {
		v2684 = v2673
		v2685 = v2674
		goto L774
	} else {
		goto L779
	}
L778:
	;
	v2684 = v2673
	v2685 = v2674
	goto L774
L779:
	;
	v2677 = int32(1)
	if v2673 == v2674 {
		v2669 = v2669 + v2677
		v2670 = v2670 + v2677
		goto L777
	} else {
		goto L780
	}
L780:
	;
	goto L778
L781:
	;
	if v2607&int32(1) != 0 {
		goto L257
	} else {
		goto L782
	}
L782:
	;
	v2690 = F_defGetBoolean(m, v2626)
	mBase = m.M
	v2691 = m.ExcPending
	if v2691 != 0 {
		goto L1
	} else {
		goto L783
	}
L783:
	;
	v2692 = int32(1)
	v2693 = v2690
	v2694 = v2611
	v2695 = v2613
	goto L759
L784:
	;
	goto L758
L785:
	;
	v2705 = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	goto L787
L786:
	;
	v2705 = int32(0)
	goto L787
L787:
	;
	if v2694&int32(1) != 0 {
		goto L788
	} else {
		goto L789
	}
L788:
	;
	v2711 = v685 + int32(464)
	goto L790
L789:
	;
	v2711 = int32(0)
	goto L790
L790:
	;
	v2733 = v2705
	v2734 = v2693
	v2739 = v2695
	v2748 = v2711
	goto L754
L791:
	;
	v2762 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+88))
	if v2763 != 0 {
		goto L795
	} else {
		goto L796
	}
L792:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_92)
	goto L244
L793:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		goto L1
	} else {
		goto L850
	}
L794:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L1
	} else {
		goto L845
	}
L795:
	;
	v2766 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v2766 == int32(1) {
		goto L801
	} else {
		goto L802
	}
L796:
	;
	goto L797
L797:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L1
	} else {
		goto L841
	}
L798:
	;
	if v2733 == int32(0) {
		goto L826
	} else {
		goto L827
	}
L799:
	;
	if v2806&int32(1) != 0 {
		goto L815
	} else {
		goto L816
	}
L800:
	;
	if v2776 != 0 {
		goto L804
	} else {
		goto L805
	}
L801:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2771)+316))
	v2774 = base.B2i32(v2772 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v2774)
	v2776 = v2774
	goto L803
L802:
	;
	v2776 = int32(0)
	goto L803
L803:
	;
	goto L800
L804:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2778)+201)))
	if v2779 != 0 {
		goto L794
	} else {
		goto L807
	}
L805:
	;
	goto L806
L806:
	;
	if v2748 == int32(0) {
		v2836 = v2752
		goto L798
	} else {
		goto L814
	}
L807:
	;
	if v2748 == int32(0) {
		v2836 = v2752
		goto L798
	} else {
		goto L808
	}
L808:
	;
	v2782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2748))))
	if v2782 != int32(1) {
		v2806 = v2782
		v2807 = v2778
		goto L799
	} else {
		goto L809
	}
L809:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L1
	} else {
		goto L810
	}
L810:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_93), int32(0))
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_94), int32(913), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L814:
	;
	v2804 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2748))))
	v2806 = v2805
	v2807 = v2804
	goto L799
L815:
	;
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+92))
	if v2810 == int32(2) {
		goto L793
	} else {
		goto L818
	}
L816:
	;
	goto L817
L817:
	;
	v2813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2807)+202)))
	if v2813 == v2806&int32(255) {
		v2836 = v2752
		goto L798
	} else {
		goto L819
	}
L818:
	;
	goto L817
L819:
	;
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(v2807)))
	v2818 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2807))) = v2818
	if v2817 != 0 {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	v2822 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	F_s_lock(m, v2822, int32(_a_F_PostgresMainLoopOnce_94), int32(929), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L1
	} else {
		goto L823
	}
L821:
	;
	goto L822
L822:
	;
	v2828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2748))))
	v2830 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v2830))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2830)+202)) = uint8(v2828)
	v2836 = v2818
	goto L798
L823:
	;
	goto L822
L824:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L840
	}
L825:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)))
	*(*int32)(unsafe.Add(mBase, uint32(v2868))) = int32(1)
	if v2869 != 0 {
		goto L834
	} else {
		goto L835
	}
L826:
	;
	if v2836 == int32(0) {
		goto L824
	} else {
		goto L833
	}
L827:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2841)+136)))
	v2843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2733))))
	if v2842 == v2843 {
		goto L826
	} else {
		goto L828
	}
L828:
	;
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2841)))
	*(*int32)(unsafe.Add(mBase, uint32(v2841))) = int32(1)
	if v2845 != 0 {
		goto L829
	} else {
		goto L830
	}
L829:
	;
	v2849 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	F_s_lock(m, v2849, int32(_a_F_PostgresMainLoopOnce_94), int32(939), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L1
	} else {
		goto L832
	}
L830:
	;
	goto L831
L831:
	;
	v2855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2733))))
	v2857 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v2857))) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2857)+136)) = uint8(v2855)
	goto L825
L832:
	;
	goto L831
L833:
	;
	goto L825
L834:
	;
	F_s_lock(m, v2868, int32(_a_F_PostgresMainLoopOnce_94), int32(1107), int32(_a_F_PostgresMainLoopOnce_96))
	mBase = m.M
	v2876 = m.ExcPending
	if v2876 != 0 {
		goto L1
	} else {
		goto L837
	}
L835:
	;
	goto L836
L836:
	;
	v2877 = int32(_a_F_PostgresMainLoopOnce_97)
	v2878 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v2879 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v2878)+12)) = uint16(v2879)
	*(*int32)(unsafe.Add(mBase, uint32(v2868))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2755)+16)) = int32(_a_F_PostgresMainLoopOnce_98)
	v2886 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v2755)+20)) = v2886 + int32(24)
	v2895 = F_pg_sprintf(m, v2755+int32(48), int32(_a_F_PostgresMainLoopOnce_99), v2755+int32(16))
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L1
	} else {
		goto L838
	}
L837:
	;
	goto L836
L838:
	;
	v2898 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	F_SaveSlotToPath(m, v2898, v2755+int32(48), int32(21))
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L1
	} else {
		goto L839
	}
L839:
	;
	goto L824
L840:
	;
	m.G0 = v2755 + int32(1072)
	goto L792
L841:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L1
	} else {
		goto L842
	}
L842:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2755))) = int32(_a_F_PostgresMainLoopOnce_92)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_100), v2755)
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L1
	} else {
		goto L843
	}
L843:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_94), int32(891), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2928 = m.ExcPending
	if v2928 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L845:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L846
	}
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2755)+32)) = v2751
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_101), v2755+int32(32))
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_102), int32(0))
	mBase = m.M
	v2945 = m.ExcPending
	if v2945 != 0 {
		goto L1
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_94), int32(903), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2950 = m.ExcPending
	if v2950 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L1
	} else {
		goto L851
	}
L851:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_103), int32(0))
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L1
	} else {
		goto L852
	}
L852:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_94), int32(925), int32(_a_F_PostgresMainLoopOnce_95))
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L1
	} else {
		goto L853
	}
L853:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L854:
	;
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+4))
	if v2972 == int32(0) {
		goto L855
	} else {
		goto L856
	}
L855:
	;
	v2975 = m.G0
	v2977 = v2975 - int32(144)
	m.G0 = v2977
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+116)) = int32(1031)
	v2983 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+112)) = v2983
	v2987 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[88]))
	v2991 = F_XLogReaderAllocate(m, v2987, v2977+int32(112), v2983)
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L1
	} else {
		goto L858
	}
L856:
	;
	goto L857
L857:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L1
	} else {
		goto L967
	}
L858:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89])) = v2991
	if v2991 != 0 {
		goto L866
	} else {
		goto L867
	}
L859:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L1
	} else {
		goto L966
	}
L860:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L1
	} else {
		goto L963
	}
L861:
	;
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+8))
	if v3309 != 0 {
		goto L945
	} else {
		goto L946
	}
L862:
	;
	v3190 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v3190)+4))
	if v3191 != int32(2) {
		goto L920
	} else {
		goto L921
	}
L863:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90])) = v3169
	v3173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[91])) = uint8(v3173)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[92])) = uint8(v3173)
	v3179 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])))
	if v3179&int32(1) == v3173 {
		goto L862
	} else {
		goto L918
	}
L864:
	;
	v3169 = int64(0)
	goto L863
L865:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L1
	} else {
		goto L914
	}
L866:
	;
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+8))
	if v2994 != 0 {
		goto L869
	} else {
		goto L870
	}
L867:
	;
	goto L868
L868:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L909
	}
L869:
	;
	v2995 = int32(1)
	F_ReplicationSlotAcquire(m, v2994, v2995, v2995)
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L1
	} else {
		goto L872
	}
L870:
	;
	goto L871
L871:
	;
	v3005 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v3005 == int32(1) {
		goto L875
	} else {
		goto L876
	}
L872:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v3001 = *(*int32)(unsafe.Add(mBase, uint32(v3000)+88))
	if v3001 != 0 {
		goto L865
	} else {
		goto L873
	}
L873:
	;
	goto L871
L874:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])) = uint8(v3015)
	if v3015 != 0 {
		goto L879
	} else {
		goto L880
	}
L875:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3010)+316))
	v3013 = base.B2i32(v3011 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v3013)
	v3015 = v3013
	goto L877
L876:
	;
	v3015 = int32(0)
	goto L877
L877:
	;
	goto L874
L878:
	;
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+12))
	if v3059 != 0 {
		goto L894
	} else {
		goto L895
	}
L879:
	;
	v3020 = F_GetWalRcvFlushRecPtr(m, int32(0), v2977+int32(128))
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		goto L1
	} else {
		goto L882
	}
L880:
	;
	goto L881
L881:
	;
	v3034 = v2977 + int32(124)
	v3037 = int32(_a_F_PostgresMainLoopOnce_74)
	v3038 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3039 = *(*int64)(unsafe.Add(mBase, uint32(v3038)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v3038)+280)) = v3039
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70])) = v3039
	v3044 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3045 = *(*int64)(unsafe.Add(mBase, uint32(v3044)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v3044)+272)) = v3045
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[71])) = v3045
	if v3034 != 0 {
		goto L891
	} else {
		goto L892
	}
L882:
	;
	v3024 = F_GetXLogReplayRecPtr(m, v2977+int32(80))
	mBase = m.M
	v3025 = m.ExcPending
	if v3025 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+124)) = v3026
	if base.Ui64(v3024) < base.Ui64(v3020) {
		goto L884
	} else {
		goto L885
	}
L884:
	;
	v3029 = v3020
	goto L886
L885:
	;
	v3029 = v3024
	goto L886
L886:
	;
	v3030 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+128))
	if v3026 == v3030 {
		goto L887
	} else {
		goto L888
	}
L887:
	;
	v3032 = v3029
	goto L889
L888:
	;
	v3032 = v3024
	goto L889
L889:
	;
	v3058 = v3032
	goto L878
L890:
	;
	v3058 = v3054
	goto L878
L891:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(v3050)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v3034))) = v3051
	goto L893
L892:
	;
	goto L893
L893:
	;
	v3054 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[70]))
	goto L890
L894:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[94])) = v3059
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+124))
	if v3062 == v3059 {
		goto L897
	} else {
		goto L898
	}
L895:
	;
	goto L896
L896:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[94])) = v3115
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90])) = int64(0)
	v3121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])) = uint8(v3121)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[91])) = uint8(v3121)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[92])) = uint8(v3121)
	goto L862
L897:
	;
	v3065 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])) = uint8(v3065)
	goto L864
L898:
	;
	goto L899
L899:
	;
	v3068 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])) = uint8(v3068)
	v3070 = F_readTimeLineHistory(m, v3062)
	mBase = m.M
	v3071 = m.ExcPending
	if v3071 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+12))
	v3074 = F_tliSwitchPoint(m, v3072, v3070, int32(_a_F_PostgresMainLoopOnce_104))
	mBase = m.M
	v3075 = m.ExcPending
	if v3075 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	F_list_free_deep(m, v3070)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	if v3074 == int64(0) {
		goto L864
	} else {
		goto L903
	}
L903:
	;
	v3080 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	if base.Ui64(v3080) <= base.Ui64(v3074) {
		v3169 = v3074
		goto L863
	} else {
		goto L904
	}
L904:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	v3086 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+56)) = v3087
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+52)) = uint32(v3086)
	v3091 = int64(base.Ui64(v3086) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+48)) = uint32(v3091)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_105), v2977+int32(48))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	v3098 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+32)) = v3098
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+40)) = uint32(v3074)
	v3102 = int64(base.Ui64(v3074) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+36)) = uint32(v3102)
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_106), v2977+int32(32))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(914), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L1
	} else {
		goto L908
	}
L908:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L909:
	;
	F_errcode(m, int32(_a_F_PostgresMainLoopOnce_108))
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_109), int32(0))
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L1
	} else {
		goto L911
	}
L911:
	;
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_110), int32(0))
	mBase = m.M
	v3143 = m.ExcPending
	if v3143 != 0 {
		goto L1
	} else {
		goto L912
	}
L912:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(826), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L1
	} else {
		goto L913
	}
L913:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L914:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_111), int32(0))
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(843), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3164 = m.ExcPending
	if v3164 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L918:
	;
	v3184 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	if base.Ui64(v3169) <= base.Ui64(v3184) {
		goto L861
	} else {
		goto L919
	}
L919:
	;
	goto L862
L920:
	;
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v3190)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3190)+76)) = int32(1)
	if v3194 != 0 {
		goto L923
	} else {
		goto L924
	}
L921:
	;
	goto L922
L922:
	;
	F_pq_beginmessage(m, v2977+int32(128), int32(87))
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L1
	} else {
		goto L927
	}
L923:
	;
	F_s_lock(m, v3190+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3203 = m.ExcPending
	if v3203 != 0 {
		goto L1
	} else {
		goto L926
	}
L924:
	;
	goto L925
L925:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3190)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3190)+4)) = int32(2)
	goto L922
L926:
	;
	goto L925
L927:
	;
	F_enlargeStringInfo(m, v2977+int32(128), int32(1))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L928
	}
L928:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+132))
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+128))
	v3222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3219+v3220))) = uint8(v3222)
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+132)) = v3219 + int32(1)
	F_enlargeStringInfo(m, v2977+int32(128), int32(2))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L1
	} else {
		goto L929
	}
L929:
	;
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+132))
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v2977)+128))
	v3235 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3232+v3233))) = uint16(v3235)
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+132)) = v3232 + int32(2)
	F_pq_endmessage(m, v2977+int32(128))
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L1
	} else {
		goto L930
	}
L930:
	;
	v3245 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[95]))
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3245)+4))
	v3247 = m.T0[v3246].(func(*base.Module) int32)(m)
	mBase = m.M
	v3248 = m.ExcPending
	if v3248 != 0 {
		goto L1
	} else {
		goto L931
	}
L931:
	;
	v3249 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	if base.Ui64(v3058) < base.Ui64(v3249) {
		goto L860
	} else {
		goto L932
	}
L932:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[96])) = v3249
	v3254 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v3254)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3254)+76)) = int32(1)
	if v3255 != 0 {
		goto L933
	} else {
		goto L934
	}
L933:
	;
	v3259 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	F_s_lock(m, v3259+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(965), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L1
	} else {
		goto L936
	}
L934:
	;
	goto L935
L935:
	;
	v3268 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[96]))
	v3270 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v3270)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3270)+8)) = v3268
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L1
	} else {
		goto L937
	}
L936:
	;
	goto L935
L937:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = int32(1)
	F_WalSndLoop(m, int32(1037))
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L1
	} else {
		goto L938
	}
L938:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = int32(0)
	v3286 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v3286 != 0 {
		goto L859
	} else {
		goto L939
	}
L939:
	;
	v3288 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+4))
	if v3289 == int32(0) {
		goto L861
	} else {
		goto L940
	}
L940:
	;
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3288)+76)) = int32(1)
	if v3292 != 0 {
		goto L941
	} else {
		goto L942
	}
L941:
	;
	F_s_lock(m, v3288+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L1
	} else {
		goto L944
	}
L942:
	;
	goto L943
L943:
	;
	v3302 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3288)+76)) = v3302
	*(*int32)(unsafe.Add(mBase, uint32(v3288)+4)) = v3302
	goto L861
L944:
	;
	goto L943
L945:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3311 = m.ExcPending
	if v3311 != 0 {
		goto L1
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	v3313 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[93])))
	if v3313 == int32(1) {
		goto L949
	} else {
		goto L950
	}
L948:
	;
	goto L947
L949:
	;
	v3317 = *(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[90]))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+20)) = uint32(v3317)
	v3319 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2977)+70)) = uint16(v3319)
	v3322 = int64(base.Ui64(v3317) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+16)) = uint32(v3322)
	v3330 = F_pg_snprintf(m, v2977+int32(80), int32(18), int32(_a_F_PostgresMainLoopOnce_72), v2977+int32(16))
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L1
	} else {
		goto L952
	}
L950:
	;
	goto L951
L951:
	;
	F_EndReplicationCommand(m, int32(_a_F_PostgresMainLoopOnce_112))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L962
	}
L952:
	;
	v3333 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	v3336 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3337 = m.ExcPending
	if v3337 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	F_TupleDescInitBuiltinEntry(m, v3336, int32(1), int32(_a_F_PostgresMainLoopOnce_113), int32(20))
	mBase = m.M
	v3342 = m.ExcPending
	if v3342 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	F_TupleDescInitBuiltinEntry(m, v3336, int32(2), int32(_a_F_PostgresMainLoopOnce_114), int32(25))
	mBase = m.M
	v3347 = m.ExcPending
	if v3347 != 0 {
		goto L1
	} else {
		goto L956
	}
L956:
	;
	v3349 = F_begin_tup_output_tupdesc(m, v3333, v3336, int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L1
	} else {
		goto L957
	}
L957:
	;
	v3352 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[98])))
	v3353 = F_Int64GetDatum(m, v3352)
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L1
	} else {
		goto L958
	}
L958:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+72)) = v3353
	v3358 = F_cstring_to_text(m, v2977+int32(80))
	mBase = m.M
	v3359 = m.ExcPending
	if v3359 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2977)+76)) = v3358
	F_do_tup_output(m, v3349, v2977+int32(72), v2977+int32(70))
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	F_end_tup_output(m, v3349)
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	goto L951
L962:
	;
	m.G0 = v2977 + int32(144)
	v5162 = v2967
	goto L244
L963:
	;
	v3382 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+4)) = uint32(v3382)
	v3384 = int64(32)
	v3385 = int64(base.Ui64(v3382) >> (uint(v3384) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977))) = uint32(v3385)
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+12)) = uint32(v3058)
	v3389 = int64(base.Ui64(v3058) >> (uint(v3384) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2977)+8)) = uint32(v3389)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_115), v2977)
	mBase = m.M
	v3393 = m.ExcPending
	if v3393 != 0 {
		goto L1
	} else {
		goto L964
	}
L964:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(958), int32(_a_F_PostgresMainLoopOnce_107))
	mBase = m.M
	v3398 = m.ExcPending
	if v3398 != 0 {
		goto L1
	} else {
		goto L965
	}
L965:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+8))
	v3405 = int32(1)
	F_ReplicationSlotAcquire(m, v3404, v3405, v3405)
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	v3410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[68])))
	if v3410 != int32(1) {
		goto L969
	} else {
		goto L970
	}
L969:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+24))
	v3443 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = int32(1032)
	v3457 = F_CreateDecodingContext(m, v3443, v3442, int32(0), v685+int32(_a_F_PostgresMainLoopOnce_73), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v3458 = m.ExcPending
	if v3458 != 0 {
		goto L1
	} else {
		goto L982
	}
L970:
	;
	v3415 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])))
	if v3415 == int32(1) {
		goto L972
	} else {
		goto L973
	}
L971:
	;
	if v3425 != 0 {
		goto L969
	} else {
		goto L975
	}
L972:
	;
	v3420 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[69]))
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(v3420)+316))
	v3423 = base.B2i32(v3421 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[67])) = uint8(v3423)
	v3425 = v3423
	goto L974
L973:
	;
	v3425 = int32(0)
	goto L974
L974:
	;
	goto L971
L975:
	;
	v3428 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L1
	} else {
		goto L976
	}
L976:
	;
	if v3428 != 0 {
		goto L977
	} else {
		goto L978
	}
L977:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_116), int32(0))
	mBase = m.M
	v3433 = m.ExcPending
	if v3433 != 0 {
		goto L1
	} else {
		goto L980
	}
L978:
	;
	goto L979
L979:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45])) = int32(1)
	goto L969
L980:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1467), int32(_a_F_PostgresMainLoopOnce_117))
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	goto L979
L982:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99])) = v3457
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(v3457)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[89])) = v3461
	v3464 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3465 = *(*int32)(unsafe.Add(mBase, uint32(v3464)+4))
	if v3465 != int32(2) {
		goto L983
	} else {
		goto L984
	}
L983:
	;
	v3468 = *(*int32)(unsafe.Add(mBase, uint32(v3464)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3464)+76)) = int32(1)
	if v3468 != 0 {
		goto L986
	} else {
		goto L987
	}
L984:
	;
	goto L985
L985:
	;
	F_pq_beginmessage(m, v685+int32(464), int32(87))
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L990
	}
L986:
	;
	F_s_lock(m, v3464+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3477 = m.ExcPending
	if v3477 != 0 {
		goto L1
	} else {
		goto L989
	}
L987:
	;
	goto L988
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3464)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3464)+4)) = int32(2)
	goto L985
L989:
	;
	goto L988
L990:
	;
	F_enlargeStringInfo(m, v685+int32(464), int32(1))
	mBase = m.M
	v3492 = m.ExcPending
	if v3492 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	v3493 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v3496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3493+v3494))) = uint8(v3496)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v3493 + int32(1)
	F_enlargeStringInfo(m, v685+int32(464), int32(2))
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v3507 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v3509 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3506+v3507))) = uint16(v3509)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v3506 + int32(2)
	F_pq_endmessage(m, v685+int32(464))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	v3519 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[95]))
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(v3519)+4))
	v3521 = m.T0[v3520].(func(*base.Module) int32)(m)
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99]))
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v3524)+8))
	v3527 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v3528 = *(*int64)(unsafe.Add(mBase, uint32(v3527)+104))
	F_XLogBeginRead(m, v3525, v3528)
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L1
	} else {
		goto L995
	}
L995:
	;
	v3533 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v3534 = *(*int64)(unsafe.Add(mBase, uint32(v3533)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[96])) = v3534
	v3537 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3538 = *(*int32)(unsafe.Add(mBase, uint32(v3537)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3537)+76)) = int32(1)
	if v3538 != 0 {
		goto L996
	} else {
		goto L997
	}
L996:
	;
	v3542 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	F_s_lock(m, v3542+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(1507), int32(_a_F_PostgresMainLoopOnce_117))
	mBase = m.M
	v3549 = m.ExcPending
	if v3549 != 0 {
		goto L1
	} else {
		goto L999
	}
L997:
	;
	goto L998
L998:
	;
	v3551 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v3552 = *(*int64)(unsafe.Add(mBase, uint32(v3551)+104))
	v3554 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v3554)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3554)+8)) = v3552
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = int32(1)
	F_SyncRepInitConfig(m)
	mBase = m.M
	v3562 = m.ExcPending
	if v3562 != 0 {
		goto L1
	} else {
		goto L1000
	}
L999:
	;
	goto L998
L1000:
	;
	F_WalSndLoop(m, int32(1036))
	mBase = m.M
	v3565 = m.ExcPending
	if v3565 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	v3567 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[99]))
	F_FreeDecodingContext(m, v3567)
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[97])) = int32(0)
	v3576 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[45]))
	if v3576 != 0 {
		goto L255
	} else {
		goto L1004
	}
L1004:
	;
	v3578 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[44]))
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v3578)+4))
	if v3579 != 0 {
		goto L1005
	} else {
		goto L1006
	}
L1005:
	;
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v3578)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3578)+76)) = int32(1)
	if v3580 != 0 {
		goto L1008
	} else {
		goto L1009
	}
L1006:
	;
	goto L1007
L1007:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[100]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[86]))) = int32(56)
	F_EndCommand(m, v685+int32(_a_F_PostgresMainLoopOnce_70), int32(2))
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1008:
	;
	F_s_lock(m, v3578+int32(76), int32(_a_F_PostgresMainLoopOnce_14), int32(3869), int32(_a_F_PostgresMainLoopOnce_15))
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1009:
	;
	goto L1010
L1010:
	;
	v3590 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3578)+76)) = v3590
	*(*int32)(unsafe.Add(mBase, uint32(v3578)+4)) = v3590
	goto L1007
L1011:
	;
	goto L1010
L1012:
	;
	v5162 = v2967
	goto L244
L1013:
	;
	v3608 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v3610 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L1
	} else {
		goto L1014
	}
L1014:
	;
	v3613 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v3614 = m.ExcPending
	if v3614 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	F_TupleDescInitBuiltinEntry(m, v3613, int32(1), int32(_a_F_PostgresMainLoopOnce_118), int32(25))
	mBase = m.M
	v3619 = m.ExcPending
	if v3619 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	F_TupleDescInitBuiltinEntry(m, v3613, int32(2), int32(_a_F_PostgresMainLoopOnce_119), int32(25))
	mBase = m.M
	v3624 = m.ExcPending
	if v3624 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	v3625 = *(*int32)(unsafe.Add(mBase, uint32(v3608)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+352)) = v3625
	v3633 = F_pg_snprintf(m, v685+int32(_a_F_PostgresMainLoopOnce_73), int32(64), int32(_a_F_PostgresMainLoopOnce_120), v685+int32(352))
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L1
	} else {
		goto L1018
	}
L1018:
	;
	v3635 = *(*int32)(unsafe.Add(mBase, uint32(v3608)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+336)) = v3635
	v3643 = F_pg_snprintf(m, v685+int32(_a_F_PostgresMainLoopOnce_70), int32(1024), int32(_a_F_PostgresMainLoopOnce_121), v685+int32(336))
	mBase = m.M
	v3644 = m.ExcPending
	if v3644 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1019:
	;
	v3646 = *(*int32)(unsafe.Add(mBase, uint32(v3610)+4))
	m.T0[v3646].(func(*base.Module, int32, int32, int32))(m, v3610, int32(1), v3613)
	mBase = m.M
	v3648 = m.ExcPending
	if v3648 != 0 {
		goto L1
	} else {
		goto L1020
	}
L1020:
	;
	F_pq_beginmessage(m, v685+int32(_a_F_PostgresMainLoopOnce_79), int32(68))
	mBase = m.M
	v3653 = m.ExcPending
	if v3653 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	F_enlargeStringInfo(m, v685+int32(_a_F_PostgresMainLoopOnce_79), int32(2))
	mBase = m.M
	v3658 = m.ExcPending
	if v3658 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1022:
	;
	v3659 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v3660 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v3662 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3659+v3660))) = uint16(v3662)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v3659 + int32(2)
	v3669 = F_strlen(m, v685+int32(_a_F_PostgresMainLoopOnce_73))
	mBase = m.M
	F_enlargeStringInfo(m, v685+int32(_a_F_PostgresMainLoopOnce_79), int32(4))
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v3676 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v3678 = int32(24)
	v3680 = int32(_a_F_PostgresMainLoopOnce_122)
	v3682 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v3675+v3676))) = v3669<<(uint(v3678)%32) | v3669&v3680<<(uint(v3682)%32) | (int32(base.Ui32(v3669)>>(uint(v3682)%32))&v3680 | int32(base.Ui32(v3669)>>(uint(v3678)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v3675 + int32(4)
	F_pq_sendbytes(m, v685+int32(_a_F_PostgresMainLoopOnce_79), v685+int32(_a_F_PostgresMainLoopOnce_73), v3669)
	mBase = m.M
	v3702 = m.ExcPending
	if v3702 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	v3706 = F_OpenTransientFile(m, v685+int32(_a_F_PostgresMainLoopOnce_70), int32(0))
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	if v3706 < int32(0) {
		goto L254
	} else {
		goto L1026
	}
L1026:
	;
	v3710 = int64(0)
	v3712 = F___lseek(m, v3706, v3710, int32(2))
	mBase = m.M
	if v3712 < v3710 {
		goto L253
	} else {
		goto L1027
	}
L1027:
	;
	v3715 = int64(0)
	v3717 = F___lseek(m, v3706, v3715, int32(0))
	mBase = m.M
	if v3717 != v3715 {
		goto L252
	} else {
		goto L1028
	}
L1028:
	;
	F_enlargeStringInfo(m, v685+int32(_a_F_PostgresMainLoopOnce_79), int32(4))
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1029:
	;
	v3725 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78])))
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[64])))
	v3728 = base.I32_wrap_i64(v3712)
	v3729 = int32(24)
	v3731 = int32(_a_F_PostgresMainLoopOnce_122)
	v3733 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v3725+v3726))) = v3728<<(uint(v3729)%32) | v3728&v3731<<(uint(v3733)%32) | (int32(base.Ui32(v3728)>>(uint(v3733)%32))&v3731 | int32(base.Ui32(v3728)>>(uint(v3729)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[78]))) = v3725 + int32(4)
	if v3712 != int64(0) {
		goto L1030
	} else {
		goto L1031
	}
L1030:
	;
	v3780 = v3712
	goto L1033
L1031:
	;
	goto L1032
L1032:
	;
	v3848 = F_CloseTransientFile(m, v3706)
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1033:
	;
	v3786 = int32(_a_F_PostgresMainLoopOnce_123)
	v3787 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101]))
	*(*int32)(unsafe.Add(mBase, uint32(v3787))) = int32(167772227)
	v3793 = F_read(m, v3706, v685+int32(464), int32(_a_F_PostgresMainLoopOnce_20))
	mBase = m.M
	v3795 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[101]))
	v3796 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3795))) = v3796
	if v3793 < v3796 {
		goto L251
	} else {
		goto L1035
	}
L1034:
	;
	goto L1032
L1035:
	;
	if v3793 == int32(0) {
		goto L250
	} else {
		goto L1036
	}
L1036:
	;
	F_pq_sendbytes(m, v685+int32(_a_F_PostgresMainLoopOnce_79), v685+int32(464), v3793)
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	v3809 = v3780 - base.I64_extend_i32_u(v3793)
	if int64(0) < v3809 {
		v3780 = v3809
		goto L1033
	} else {
		goto L1038
	}
L1038:
	;
	goto L1034
L1039:
	;
	if v3848 != 0 {
		goto L249
	} else {
		goto L1040
	}
L1040:
	;
	F_pq_endmessage(m, v685+int32(_a_F_PostgresMainLoopOnce_79))
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L1
	} else {
		goto L1041
	}
L1041:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_68)
	goto L244
L1042:
	;
	v3861 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[102]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50])) = v3861
	v3864 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v3869 = F_AllocSetContextCreateInternal(m, v3864, int32(_a_F_PostgresMainLoopOnce_124), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L1043
	}
L1043:
	;
	v3871 = int32(_a_F_PostgresMainLoopOnce_125)
	v3872 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v3869
	v3876 = F_palloc0(m, int32(36))
	mBase = m.M
	v3877 = m.ExcPending
	if v3877 != 0 {
		goto L1
	} else {
		goto L1044
	}
L1044:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3876))) = v3869
	F_initStringInfo(m, v3876+int32(4))
	mBase = m.M
	v3882 = m.ExcPending
	if v3882 != 0 {
		goto L1
	} else {
		goto L1045
	}
L1045:
	;
	v3884 = F_MemoryContextAllocZero(m, v3869, int32(32))
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L1
	} else {
		goto L1046
	}
L1046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3884)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3884)+24)) = v3869
	v3891 = F_MemoryContextAllocExtended(m, v3869, int32(_a_F_PostgresMainLoopOnce_126), int32(5))
	mBase = m.M
	v3892 = m.ExcPending
	if v3892 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1047:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3884)+12)) = int64(63329292795903)
	*(*int64)(unsafe.Add(mBase, uint32(v3884))) = int64(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v3884)+20)) = v3891
	*(*int32)(unsafe.Add(mBase, uint32(v3876)+24)) = v3884
	v3900 = F_palloc0(m, int32(24))
	mBase = m.M
	v3901 = m.ExcPending
	if v3901 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3900)+20)) = int32(427)
	*(*int32)(unsafe.Add(mBase, uint32(v3900)+16)) = int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v3900)+12)) = int32(429)
	*(*int32)(unsafe.Add(mBase, uint32(v3900)+8)) = int32(430)
	*(*int32)(unsafe.Add(mBase, uint32(v3900)+4)) = int32(431)
	*(*int32)(unsafe.Add(mBase, uint32(v3900))) = v3876
	v3914 = F_palloc(m, int32(112))
	mBase = m.M
	v3915 = m.ExcPending
	if v3915 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1049:
	;
	v3917 = F_palloc(m, int32(68))
	mBase = m.M
	v3918 = m.ExcPending
	if v3918 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	v3919 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3917)+52)) = uint8(v3919)
	*(*int32)(unsafe.Add(mBase, uint32(v3917)+4)) = v3919
	*(*int32)(unsafe.Add(mBase, uint32(v3917))) = v3900
	if v3914 == v3919 {
		goto L1053
	} else {
		goto L1054
	}
L1051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+104)) = int32(_a_F_PostgresMainLoopOnce_127)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+100)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3914)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+88)) = int32(_a_F_PostgresMainLoopOnce_128)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+84)) = int32(_a_F_PostgresMainLoopOnce_129)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+80)) = int32(_a_F_PostgresMainLoopOnce_130)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+76)) = int32(_a_F_PostgresMainLoopOnce_131)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+72)) = int32(_a_F_PostgresMainLoopOnce_132)
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+68)) = v3917
	v4012 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v4013 = m.ExcPending
	if v4013 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+40)) = int32(1)
	v3945 = F_palloc0(m, int32(20))
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1053:
	;
	v3927 = F_palloc0(m, int32(68))
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1054:
	;
	goto L1055
L1055:
	;
	v3938 = F__emscripten_memset_bulkmem(m, v3914, base.I32_extend8_s(int32(0)), int32(68))
	mBase = m.M
	goto L1058
L1056:
	;
	if v3927 == int32(0) {
		goto L1051
	} else {
		goto L1057
	}
L1057:
	;
	v3931 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v3927)+36)) = v3931 | int32(1)
	v3939 = v3927
	goto L1052
L1058:
	;
	v3939 = v3914
	goto L1052
L1059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+52)) = v3945
	v3949 = F_palloc0(m, int32(28))
	mBase = m.M
	v3950 = m.ExcPending
	if v3950 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	v3952 = F_palloc(m, int32(640))
	mBase = m.M
	v3953 = m.ExcPending
	if v3953 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1061:
	;
	v3955 = F_palloc(m, int32(256))
	mBase = m.M
	v3956 = m.ExcPending
	if v3956 != 0 {
		goto L1
	} else {
		goto L1062
	}
L1062:
	;
	v3958 = F_palloc(m, int32(64))
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1063:
	;
	v3960 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+52))
	F_initStringInfo(m, v3960+int32(4))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+48)) = v3949
	*(*int32)(unsafe.Add(mBase, uint32(v3949))) = int32(64)
	v3968 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3968)+4)) = v3952
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3970)+12)) = v3955
	v3972 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3972)+16)) = v3958
	v3974 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+48))
	v3975 = *(*int32)(unsafe.Add(mBase, uint32(v3974)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3975))) = int32(0)
	v3978 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3939)+56)) = uint8(v3978)
	*(*uint8)(unsafe.Add(mBase, uint32(v3939)+24)) = uint8(v3978)
	v3982 = F_makeStringInfo(m)
	mBase = m.M
	v3983 = m.ExcPending
	if v3983 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+60)) = v3982
	v3985 = *(*int32)(unsafe.Add(mBase, uint32(v3939)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+36)) = v3985 | int32(2)
	goto L1051
L1066:
	;
	if v4012 == int32(0) {
		goto L1067
	} else {
		goto L1068
	}
L1067:
	;
	v4018 = *(*int32)(unsafe.Add(mBase, uint32(v3900)+20))
	m.T0[v4018].(func(*base.Module, int32, int32, int32))(m, v3900, int32(_a_F_PostgresMainLoopOnce_109), int32(0))
	mBase = m.M
	v4020 = m.ExcPending
	if v4020 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1068:
	;
	goto L1069
L1069:
	;
	v4021 = F_pg_cryptohash_init(m, v4012)
	mBase = m.M
	if v4021 < int32(0) {
		goto L1071
	} else {
		goto L1072
	}
L1070:
	;
	goto L1069
L1071:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v3900)+20))
	m.T0[v4026].(func(*base.Module, int32, int32, int32))(m, v3900, int32(_a_F_PostgresMainLoopOnce_133), int32(0))
	mBase = m.M
	v4028 = m.ExcPending
	if v4028 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1072:
	;
	goto L1073
L1073:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3914)+108)) = v4012
	*(*int32)(unsafe.Add(mBase, uint32(v3876)+32)) = v3914
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v3872
	F_pq_beginmessage(m, v685+int32(464), int32(71))
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1074:
	;
	goto L1073
L1075:
	;
	F_enlargeStringInfo(m, v685+int32(464), int32(1))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1076:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4046 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4043+v4044))) = uint8(v4046)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v4043 + int32(1)
	F_enlargeStringInfo(m, v685+int32(464), int32(2))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4057 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4059 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4056+v4057))) = uint16(v4059)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+468)) = v4056 + int32(2)
	F_pq_endmessage_reuse(m, v685+int32(464))
	mBase = m.M
	v4067 = m.ExcPending
	if v4067 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	v4069 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[95]))
	v4070 = *(*int32)(unsafe.Add(mBase, uint32(v4069)+4))
	v4071 = m.T0[v4070].(func(*base.Module) int32)(m)
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	goto L1081
L1080:
	;
	v4345 = int32(_a_F_PostgresMainLoopOnce_125)
	v4346 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4348 = *(*int32)(unsafe.Add(mBase, uint32(v3876)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4348
	v4350 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+32))
	v4351 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+4))
	v4352 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+8))
	F_json_parse_manifest_incremental_chunk(m, v4350, v4351, v4352, int32(1))
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1081:
	;
	v4109 = int32(_a_F_PostgresMainLoopOnce_5)
	v4111 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4111 + int32(1)
	F_pq_startmsgread(m)
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1082:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4326 = m.ExcPending
	if v4326 != 0 {
		goto L1
	} else {
		goto L1146
	}
L1083:
	;
	v4117 = F_pq_getbyte(m)
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1084:
	;
	v4120 = v4117 - int32(72)
	if base.Ui32(int32(30)) < base.Ui32(v4120) {
		goto L247
	} else {
		goto L1085
	}
L1085:
	;
	if int32(1)<<(uint(v4120)%32)&int32(1207961601) == int32(0) {
		goto L1087
	} else {
		goto L1088
	}
L1086:
	;
	v4136 = F_pq_getmessage(m, v685+int32(464), v4133)
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1087:
	;
	if v4120 != int32(28) {
		goto L247
	} else {
		goto L1090
	}
L1088:
	;
	goto L1089
L1089:
	;
	v4133 = int32(_a_F_PostgresMainLoopOnce_7)
	goto L1086
L1090:
	;
	v4133 = int32(1073741822)
	goto L1086
L1091:
	;
	if v4136 != 0 {
		goto L248
	} else {
		goto L1092
	}
L1092:
	;
	v4138 = int32(_a_F_PostgresMainLoopOnce_5)
	v4140 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[27])) = v4140 - int32(1)
	switch v4120 {
	case 0, 11:
		goto L1081
	default:
		goto L1080
	case 28:
		goto L1094
	case 30:
		goto L1093
	}
L1093:
	;
	goto L1082
L1094:
	;
	v4144 = *(*int32)(unsafe.Add(mBase, uint32(v685)+464))
	v4145 = *(*int32)(unsafe.Add(mBase, uint32(v685)+468))
	v4146 = int32(_a_F_PostgresMainLoopOnce_125)
	v4147 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v4149 = *(*int32)(unsafe.Add(mBase, uint32(v3876)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4149
	v4153 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+8))
	if v4153 < int32(1025) {
		goto L1095
	} else {
		goto L1096
	}
L1095:
	;
	F_appendBinaryStringInfo(m, v3876+int32(4), v4144, v4145)
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1096:
	;
	if v4145+v4153 < int32(_a_F_PostgresMainLoopOnce_134) {
		goto L1095
	} else {
		goto L1097
	}
L1097:
	;
	v4159 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+32))
	v4160 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+4))
	F_json_parse_manifest_incremental_chunk(m, v4159, v4160, v4153-int32(1024), int32(0))
	mBase = m.M
	v4165 = m.ExcPending
	if v4165 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	v4166 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+4))
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+8))
	v4170 = v4166 + v4167 - int32(1024)
	v4171 = int32(1025)
	if v4166 == v4170 {
		goto L1100
	} else {
		goto L1101
	}
L1099:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3876)+8)) = int32(1024)
	goto L1095
L1100:
	;
	goto L1099
L1101:
	;
	v4175 = v4166 + v4171
	if base.Ui32(v4170-v4175) <= base.Ui32(int32(-2050)) {
		goto L1102
	} else {
		goto L1103
	}
L1102:
	;
	v4182 = F___memcpy(m, v4166, v4170, v4171)
	mBase = m.M
	goto L1099
L1103:
	;
	goto L1104
L1104:
	;
	v4185 = (v4166 ^ v4170) & int32(3)
	if base.Ui32(v4166) < base.Ui32(v4170) {
		goto L1107
	} else {
		goto L1108
	}
L1105:
	;
	if v4287 == int32(0) {
		goto L1100
	} else {
		goto L1141
	}
L1106:
	;
	if base.Ui32(v4265) <= base.Ui32(int32(3)) {
		v4286 = v4264
		v4287 = v4265
		v4288 = v4266
		goto L1105
	} else {
		goto L1137
	}
L1107:
	;
	if v4185 != 0 {
		goto L1110
	} else {
		goto L1111
	}
L1108:
	;
	goto L1109
L1109:
	;
	if v4185 != 0 {
		v4247 = v4171
		goto L1120
	} else {
		goto L1121
	}
L1110:
	;
	v4286 = v4170
	v4287 = v4171
	v4288 = v4166
	goto L1105
L1111:
	;
	goto L1112
L1112:
	;
	if v4166&int32(3) == int32(0) {
		goto L1113
	} else {
		goto L1114
	}
L1113:
	;
	v4264 = v4170
	v4265 = v4171
	v4266 = v4166
	goto L1106
L1114:
	;
	goto L1115
L1115:
	;
	v4192 = v4170
	v4193 = v4171
	v4194 = v4166
	goto L1116
L1116:
	;
	if v4193 == int32(0) {
		goto L1100
	} else {
		goto L1118
	}
L1117:
	;
	v4264 = v4201
	v4265 = v4203
	v4266 = v4205
	goto L1106
L1118:
	;
	v4198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4194))) = uint8(v4198)
	v4200 = int32(1)
	v4201 = v4192 + v4200
	v4203 = v4193 - v4200
	v4205 = v4194 + v4200
	if v4205&int32(3) != 0 {
		v4192 = v4201
		v4193 = v4203
		v4194 = v4205
		goto L1116
	} else {
		goto L1119
	}
L1119:
	;
	goto L1117
L1120:
	;
	if v4247 == int32(0) {
		goto L1100
	} else {
		goto L1133
	}
L1121:
	;
	if v4175&int32(3) != 0 {
		goto L1122
	} else {
		goto L1123
	}
L1122:
	;
	v4212 = v4171
	goto L1125
L1123:
	;
	v4227 = v4171
	goto L1124
L1124:
	;
	if base.Ui32(v4227) <= base.Ui32(int32(3)) {
		v4247 = v4227
		goto L1120
	} else {
		goto L1129
	}
L1125:
	;
	if v4212 == int32(0) {
		goto L1100
	} else {
		goto L1127
	}
L1126:
	;
	v4227 = v4218
	goto L1124
L1127:
	;
	v4218 = v4212 - int32(1)
	v4219 = v4166 + v4218
	v4221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4170+v4218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4219))) = uint8(v4221)
	if v4219&int32(3) != 0 {
		v4212 = v4218
		goto L1125
	} else {
		goto L1128
	}
L1128:
	;
	goto L1126
L1129:
	;
	v4234 = v4227
	goto L1130
L1130:
	;
	v4238 = v4234 - int32(4)
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4170+v4238)))
	*(*int32)(unsafe.Add(mBase, uint32(v4166+v4238))) = v4241
	if base.Ui32(int32(3)) < base.Ui32(v4238) {
		v4234 = v4238
		goto L1130
	} else {
		goto L1132
	}
L1131:
	;
	v4247 = v4238
	goto L1120
L1132:
	;
	goto L1131
L1133:
	;
	v4254 = v4247
	goto L1134
L1134:
	;
	v4258 = v4254 - int32(1)
	v4261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4170+v4258))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4166+v4258))) = uint8(v4261)
	if v4258 != 0 {
		v4254 = v4258
		goto L1134
	} else {
		goto L1136
	}
L1135:
	;
	goto L1100
L1136:
	;
	goto L1135
L1137:
	;
	v4271 = v4264
	v4272 = v4265
	v4273 = v4266
	goto L1138
L1138:
	;
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(v4271)))
	*(*int32)(unsafe.Add(mBase, uint32(v4273))) = v4275
	v4277 = int32(4)
	v4278 = v4271 + v4277
	v4280 = v4273 + v4277
	v4282 = v4272 - v4277
	if base.Ui32(int32(3)) < base.Ui32(v4282) {
		v4271 = v4278
		v4272 = v4282
		v4273 = v4280
		goto L1138
	} else {
		goto L1140
	}
L1139:
	;
	v4286 = v4278
	v4287 = v4282
	v4288 = v4280
	goto L1105
L1140:
	;
	goto L1139
L1141:
	;
	v4293 = v4286
	v4294 = v4287
	v4295 = v4288
	goto L1142
L1142:
	;
	v4297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4293))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4295))) = uint8(v4297)
	v4299 = int32(1)
	v4304 = v4294 - v4299
	if v4304 != 0 {
		v4293 = v4293 + v4299
		v4294 = v4304
		v4295 = v4295 + v4299
		goto L1142
	} else {
		goto L1144
	}
L1143:
	;
	goto L1100
L1144:
	;
	goto L1143
L1145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4147
	goto L1081
L1146:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	v4332 = F_pq_getmsgstring(m, v685+int32(464))
	mBase = m.M
	v4333 = m.ExcPending
	if v4333 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+384)) = v4332
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_135), v685+int32(384))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(794), int32(_a_F_PostgresMainLoopOnce_136))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1151:
	;
	v4356 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+4))
	F_pfree(m, v4356)
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L1152
	}
L1152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3876)+4)) = int32(0)
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v3876)+32))
	v4362 = *(*int32)(unsafe.Add(mBase, uint32(v4361)+68))
	F_pfree(m, v4362)
	mBase = m.M
	v4364 = m.ExcPending
	if v4364 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1153:
	;
	F_freeJsonLexContext(m, v4361)
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1154:
	;
	F_pfree(m, v4361)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v4346
	v4372 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103]))
	if v4372 != 0 {
		goto L1156
	} else {
		goto L1157
	}
L1156:
	;
	F_MemoryContextDelete(m, v4372)
	mBase = m.M
	v4374 = m.ExcPending
	if v4374 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1157:
	;
	goto L1158
L1158:
	;
	v4376 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[104]))
	v4380 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+16))
	if v4380 != v4376 {
		goto L1161
	} else {
		goto L1162
	}
L1159:
	;
	goto L1158
L1160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[103])) = v3869
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105])) = v3876
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1161:
	;
	if v4380 == int32(0) {
		goto L1164
	} else {
		goto L1165
	}
L1162:
	;
	goto L1163
L1163:
	;
	goto L1160
L1164:
	;
	if v4376 != 0 {
		goto L1171
	} else {
		goto L1172
	}
L1165:
	;
	v4384 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+28))
	v4385 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+24))
	if v4385 != 0 {
		goto L1167
	} else {
		goto L1168
	}
L1166:
	;
	if v4384 == int32(0) {
		goto L1164
	} else {
		goto L1170
	}
L1167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4385)+28)) = v4384
	goto L1166
L1168:
	;
	goto L1169
L1169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4380)+20)) = v4384
	goto L1166
L1170:
	;
	v4390 = *(*int32)(unsafe.Add(mBase, uint32(v3869)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4384)+24)) = v4390
	goto L1164
L1171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3869)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3869)+16)) = v4376
	v4397 = *(*int32)(unsafe.Add(mBase, uint32(v4376)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3869)+28)) = v4397
	if v4397 != 0 {
		goto L1174
	} else {
		goto L1175
	}
L1172:
	;
	goto L1173
L1173:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3869)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3869)+16)) = int32(0)
	goto L1163
L1174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4397)+24)) = v3869
	goto L1176
L1175:
	;
	goto L1176
L1176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4376)+20)) = v3869
	goto L1160
L1177:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_67)
	goto L244
L1178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	v4424 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(v4424)))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v4425
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_137), v685)
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2215), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1182:
	;
	v4440 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	v4442 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[105]))
	F_SendBaseBackup(m, v4440, v4442)
	mBase = m.M
	v4444 = m.ExcPending
	if v4444 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1183:
	;
	v5162 = v4435
	goto L244
L1184:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		goto L1
	} else {
		goto L1185
	}
L1185:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_138), int32(0))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2010), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v4460 = m.ExcPending
	if v4460 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+416)) = v1342
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_139), v685+int32(416))
	mBase = m.M
	v4473 = m.ExcPending
	if v4473 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1190:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2078), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v4478 = m.ExcPending
	if v4478 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1192:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v4485 = m.ExcPending
	if v4485 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v4489 = m.ExcPending
	if v4489 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(2104), int32(_a_F_PostgresMainLoopOnce_30))
	mBase = m.M
	v4494 = m.ExcPending
	if v4494 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1196:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4501 = m.ExcPending
	if v4501 != 0 {
		goto L1
	} else {
		goto L1197
	}
L1197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+64)) = int32(_a_F_PostgresMainLoopOnce_86)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_141), v685-int32(-64))
	mBase = m.M
	v4508 = m.ExcPending
	if v4508 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(520), int32(_a_F_PostgresMainLoopOnce_85))
	mBase = m.M
	v4513 = m.ExcPending
	if v4513 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1200:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4520 = m.ExcPending
	if v4520 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1137), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v4529 = m.ExcPending
	if v4529 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1204:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4536 = m.ExcPending
	if v4536 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		goto L1
	} else {
		goto L1206
	}
L1206:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1159), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v4545 = m.ExcPending
	if v4545 != 0 {
		goto L1
	} else {
		goto L1207
	}
L1207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1208:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4552 = m.ExcPending
	if v4552 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1209:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4556 = m.ExcPending
	if v4556 != 0 {
		goto L1
	} else {
		goto L1210
	}
L1210:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1169), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v4561 = m.ExcPending
	if v4561 != 0 {
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
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4572 = m.ExcPending
	if v4572 != 0 {
		goto L1
	} else {
		goto L1214
	}
L1214:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1178), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
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
	v4582 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+208)) = v4582
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_143), v685+int32(208))
	mBase = m.M
	v4588 = m.ExcPending
	if v4588 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1217:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1183), int32(_a_F_PostgresMainLoopOnce_89))
	mBase = m.M
	v4593 = m.ExcPending
	if v4593 != 0 {
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
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1220:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4604 = m.ExcPending
	if v4604 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1420), int32(_a_F_PostgresMainLoopOnce_144))
	mBase = m.M
	v4609 = m.ExcPending
	if v4609 != 0 {
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
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_142), int32(0))
	mBase = m.M
	v4620 = m.ExcPending
	if v4620 != 0 {
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1429), int32(_a_F_PostgresMainLoopOnce_144))
	mBase = m.M
	v4625 = m.ExcPending
	if v4625 != 0 {
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
	v4630 = *(*int32)(unsafe.Add(mBase, uint32(v2626)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+224)) = v4630
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_143), v685+int32(224))
	mBase = m.M
	v4636 = m.ExcPending
	if v4636 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1228:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1434), int32(_a_F_PostgresMainLoopOnce_144))
	mBase = m.M
	v4641 = m.ExcPending
	if v4641 != 0 {
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
	base.Wasm_trap_unreachable()
	for {
	}
L1231:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4650 = m.ExcPending
	if v4650 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+240)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_145), v685+int32(240))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1233:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(616), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4663 = m.ExcPending
	if v4663 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1235:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+256)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_147), v685+int32(256))
	mBase = m.M
	v4677 = m.ExcPending
	if v4677 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(623), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4682 = m.ExcPending
	if v4682 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+320)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_148), v685+int32(320))
	mBase = m.M
	v4696 = m.ExcPending
	if v4696 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(627), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4701 = m.ExcPending
	if v4701 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v4707 = m.ExcPending
	if v4707 != 0 {
		goto L1
	} else {
		goto L1244
	}
L1244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+288)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_149), v685+int32(288))
	mBase = m.M
	v4715 = m.ExcPending
	if v4715 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(644), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4720 = m.ExcPending
	if v4720 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1247:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4727 = m.ExcPending
	if v4727 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+312)) = uint32(v3780)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+308)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+304)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_150), v685+int32(304))
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(649), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4743 = m.ExcPending
	if v4743 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1251:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+272)) = v685 + int32(_a_F_PostgresMainLoopOnce_70)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_151), v685+int32(272))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1253:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(658), int32(_a_F_PostgresMainLoopOnce_146))
	mBase = m.M
	v4762 = m.ExcPending
	if v4762 != 0 {
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
	F_errcode(m, int32(100663808))
	mBase = m.M
	v4769 = m.ExcPending
	if v4769 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1256:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v4773 = m.ExcPending
	if v4773 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1257:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(772), int32(_a_F_PostgresMainLoopOnce_136))
	mBase = m.M
	v4778 = m.ExcPending
	if v4778 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4785 = m.ExcPending
	if v4785 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1260:
	;
	goto L1261
L1261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4801 = m.ExcPending
	if v4801 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1262:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1263:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_8), int32(0))
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(746), int32(_a_F_PostgresMainLoopOnce_136))
	mBase = m.M
	v4797 = m.ExcPending
	if v4797 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v4804 = m.ExcPending
	if v4804 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+368)) = v4117
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_152), v685+int32(368))
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(763), int32(_a_F_PostgresMainLoopOnce_136))
	mBase = m.M
	v4815 = m.ExcPending
	if v4815 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1270:
	;
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(v685)+424))
	F_StartTransactionCommand(m)
	mBase = m.M
	v4822 = m.ExcPending
	if v4822 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	v4823 = *(*int32)(unsafe.Add(mBase, uint32(v4819)+4))
	F_GetPGVariable(m, v4823, v4817)
	mBase = m.M
	v4825 = m.ExcPending
	if v4825 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4827 = m.ExcPending
	if v4827 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1273:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_153)
	goto L244
L1274:
	;
	v5076 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v5077 = *(*int64)(unsafe.Add(mBase, uint32(v5076)+120))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+84)) = uint32(v5077)
	v5080 = int64(base.Ui64(v5077) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v685)+80)) = uint32(v5080)
	v5088 = F_pg_snprintf(m, v685+int32(464), int32(64), int32(_a_F_PostgresMainLoopOnce_72), v685+int32(80))
	mBase = m.M
	v5089 = m.ExcPending
	if v5089 != 0 {
		goto L1
	} else {
		goto L1339
	}
L1275:
	;
	v4867 = int32(0)
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+4))
	v4872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2252)+16)))
	if v4872 != 0 {
		goto L1278
	} else {
		goto L1279
	}
L1276:
	;
	goto L1277
L1277:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v4891 = m.ExcPending
	if v4891 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1278:
	;
	v4873 = int32(2)
	goto L1280
L1279:
	;
	v4873 = v4867
	goto L1280
L1280:
	;
	v4874 = int32(0)
	F_ReplicationSlotCreate(m, v4868, v4867, v4873, v4874, v4874, v4874)
	mBase = m.M
	v4878 = m.ExcPending
	if v4878 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1281:
	;
	if v4843&int32(1) == int32(0) {
		v5074 = v4867
		goto L1274
	} else {
		goto L1282
	}
L1282:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v4884 = m.ExcPending
	if v4884 != 0 {
		goto L1
	} else {
		goto L1283
	}
L1283:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v4886 = m.ExcPending
	if v4886 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1284:
	;
	v4887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2252)+16)))
	if v4887 != 0 {
		v5074 = v4867
		goto L1274
	} else {
		goto L1285
	}
L1285:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v4889 = m.ExcPending
	if v4889 != 0 {
		goto L1
	} else {
		goto L1286
	}
L1286:
	;
	v5074 = v4867
	goto L1274
L1287:
	;
	v4892 = int32(0)
	v4893 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+4))
	v4894 = int32(1)
	v4897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2252)+16)))
	if v4897 != 0 {
		goto L1288
	} else {
		goto L1289
	}
L1288:
	;
	v4898 = int32(2)
	goto L1290
L1289:
	;
	v4898 = v4894
	goto L1290
L1290:
	;
	v4899 = int32(1)
	F_ReplicationSlotCreate(m, v4893, v4894, v4898, v4854&v4899, v4856&v4899, int32(0))
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	switch v4853 {
	case 0:
		goto L1294
	default:
		v4956 = int32(0)
		goto L1292
	case 2:
		goto L1293
	}
L1292:
	;
	v4957 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[74]))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[73]))) = int32(1031)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[72]))) = int32(1032)
	v4970 = F_CreateInitDecodingContext(m, v4957, v4956, int64(0), v685+int32(_a_F_PostgresMainLoopOnce_73), int32(1033), int32(1034), int32(1035))
	mBase = m.M
	v4971 = m.ExcPending
	if v4971 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1293:
	;
	v4932 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v4933 = *(*int32)(unsafe.Add(mBase, uint32(v4932)+24))
	goto L1300
L1294:
	;
	v4907 = int32(1)
	v4909 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v4910 = *(*int32)(unsafe.Add(mBase, uint32(v4909)+24))
	goto L1295
L1295:
	;
	if base.B2i32(base.Ui32(v4907) < base.Ui32(v4910)) == int32(0) {
		v4956 = v4907
		goto L1292
	} else {
		goto L1296
	}
L1296:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+96)) = int32(_a_F_PostgresMainLoopOnce_154)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_155), v685+int32(96))
	mBase = m.M
	v4925 = m.ExcPending
	if v4925 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1298:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1258), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v4930 = m.ExcPending
	if v4930 != 0 {
		goto L1
	} else {
		goto L1299
	}
L1299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1300:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v4933)) == int32(0) {
		goto L242
	} else {
		goto L1301
	}
L1301:
	;
	v4939 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106]))
	if v4939 != int32(2) {
		goto L241
	} else {
		goto L1302
	}
L1302:
	;
	v4943 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107])))
	if v4943 == int32(0) {
		goto L240
	} else {
		goto L1303
	}
L1303:
	;
	v4946 = int32(1)
	v4948 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[108])))
	if v4948 == v4946 {
		goto L239
	} else {
		goto L1304
	}
L1304:
	;
	v4952 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v4953 = *(*int32)(unsafe.Add(mBase, uint32(v4952)+28))
	goto L1305
L1305:
	;
	if int32(1) < v4953 {
		goto L238
	} else {
		goto L1306
	}
L1306:
	;
	v4956 = v4946
	goto L1292
L1307:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[109])) = int64(0)
	F_DecodingContextFindStartpoint(m, v4970)
	mBase = m.M
	v4976 = m.ExcPending
	if v4976 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1308:
	;
	switch v4853 {
	case 0:
		goto L1311
	default:
		v5065 = v4892
		goto L1309
	case 2:
		goto L1310
	}
L1309:
	;
	F_FreeDecodingContext(m, v4970)
	mBase = m.M
	v5067 = m.ExcPending
	if v5067 != 0 {
		goto L1
	} else {
		goto L1336
	}
L1310:
	;
	v5056 = *(*int32)(unsafe.Add(mBase, uint32(v4970)+16))
	v5057 = F_SnapBuildInitialSnapshot(m, v5056)
	mBase = m.M
	v5058 = m.ExcPending
	if v5058 != 0 {
		goto L1
	} else {
		goto L1334
	}
L1311:
	;
	v4977 = *(*int32)(unsafe.Add(mBase, uint32(v4970)+16))
	v4978 = m.G0
	v4980 = v4978 - int32(16)
	m.G0 = v4980
	v4983 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v4984 = *(*int32)(unsafe.Add(mBase, uint32(v4983)+24))
	goto L1314
L1312:
	;
	v5065 = v5008
	goto L1309
L1313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5046 = m.ExcPending
	if v5046 != 0 {
		goto L1
	} else {
		goto L1331
	}
L1314:
	;
	if base.B2i32(v4984 != int32(0)) == int32(0) {
		goto L1315
	} else {
		goto L1316
	}
L1315:
	;
	v4990 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49]))
	if v4990 != 0 {
		goto L1313
	} else {
		goto L1318
	}
L1316:
	;
	goto L1317
L1317:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5033 = m.ExcPending
	if v5033 != 0 {
		goto L1
	} else {
		goto L1328
	}
L1318:
	;
	v4992 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[48])) = uint8(v4992)
	v4996 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[50]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[49])) = v4996
	F_StartTransactionCommand(m)
	mBase = m.M
	v4999 = m.ExcPending
	if v4999 != 0 {
		goto L1
	} else {
		goto L1319
	}
L1319:
	;
	v5001 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[107])) = uint8(v5001)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[106])) = int32(2)
	v5006 = F_SnapBuildInitialSnapshot(m, v4977)
	mBase = m.M
	v5007 = m.ExcPending
	if v5007 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1320:
	;
	v5008 = F_ExportSnapshot(m, v5006)
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	v5012 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5013 = m.ExcPending
	if v5013 != 0 {
		goto L1
	} else {
		goto L1322
	}
L1322:
	;
	if v5012 != 0 {
		goto L1323
	} else {
		goto L1324
	}
L1323:
	;
	v5014 = *(*int32)(unsafe.Add(mBase, uint32(v5006)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4980)+4)) = v5014
	*(*int32)(unsafe.Add(mBase, uint32(v4980))) = v5008
	F_errmsg_plural(m, int32(_a_F_PostgresMainLoopOnce_157), int32(_a_F_PostgresMainLoopOnce_158), v5014, v4980)
	mBase = m.M
	v5020 = m.ExcPending
	if v5020 != 0 {
		goto L1
	} else {
		goto L1326
	}
L1324:
	;
	goto L1325
L1325:
	;
	m.G0 = v4980 + int32(16)
	goto L1312
L1326:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(571), int32(_a_F_PostgresMainLoopOnce_159))
	mBase = m.M
	v5025 = m.ExcPending
	if v5025 != 0 {
		goto L1
	} else {
		goto L1327
	}
L1327:
	;
	goto L1325
L1328:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_160), int32(0))
	mBase = m.M
	v5037 = m.ExcPending
	if v5037 != 0 {
		goto L1
	} else {
		goto L1329
	}
L1329:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(545), int32(_a_F_PostgresMainLoopOnce_159))
	mBase = m.M
	v5042 = m.ExcPending
	if v5042 != 0 {
		goto L1
	} else {
		goto L1330
	}
L1330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1331:
	;
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_161), int32(0))
	mBase = m.M
	v5050 = m.ExcPending
	if v5050 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1332:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_17), int32(548), int32(_a_F_PostgresMainLoopOnce_159))
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L1
	} else {
		goto L1333
	}
L1333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1334:
	;
	v5060 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[7]))
	F_RestoreTransactionSnapshot(m, v5057, v5060)
	mBase = m.M
	v5062 = m.ExcPending
	if v5062 != 0 {
		goto L1
	} else {
		goto L1335
	}
L1335:
	;
	v5065 = v4892
	goto L1309
L1336:
	;
	v5068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2252)+16)))
	if v5068 != 0 {
		v5074 = v5065
		goto L1274
	} else {
		goto L1337
	}
L1337:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v5070 = m.ExcPending
	if v5070 != 0 {
		goto L1
	} else {
		goto L1338
	}
L1338:
	;
	v5074 = v5065
	goto L1274
L1339:
	;
	v5091 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v5092 = m.ExcPending
	if v5092 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1340:
	;
	v5094 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v5095 = m.ExcPending
	if v5095 != 0 {
		goto L1
	} else {
		goto L1341
	}
L1341:
	;
	F_TupleDescInitBuiltinEntry(m, v5094, int32(1), int32(_a_F_PostgresMainLoopOnce_162), int32(25))
	mBase = m.M
	v5100 = m.ExcPending
	if v5100 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1342:
	;
	F_TupleDescInitBuiltinEntry(m, v5094, int32(2), int32(_a_F_PostgresMainLoopOnce_163), int32(25))
	mBase = m.M
	v5105 = m.ExcPending
	if v5105 != 0 {
		goto L1
	} else {
		goto L1343
	}
L1343:
	;
	F_TupleDescInitBuiltinEntry(m, v5094, int32(3), int32(_a_F_PostgresMainLoopOnce_164), int32(25))
	mBase = m.M
	v5110 = m.ExcPending
	if v5110 != 0 {
		goto L1
	} else {
		goto L1344
	}
L1344:
	;
	F_TupleDescInitBuiltinEntry(m, v5094, int32(4), int32(_a_F_PostgresMainLoopOnce_165), int32(25))
	mBase = m.M
	v5115 = m.ExcPending
	if v5115 != 0 {
		goto L1
	} else {
		goto L1345
	}
L1345:
	;
	v5117 = F_begin_tup_output_tupdesc(m, v5091, v5094, int32(_a_F_PostgresMainLoopOnce_78))
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L1
	} else {
		goto L1346
	}
L1346:
	;
	v5120 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[87]))
	v5123 = F_cstring_to_text(m, v5120+int32(24))
	mBase = m.M
	v5124 = m.ExcPending
	if v5124 != 0 {
		goto L1
	} else {
		goto L1347
	}
L1347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[86]))) = v5123
	v5128 = F_cstring_to_text(m, v685+int32(464))
	mBase = m.M
	v5129 = m.ExcPending
	if v5129 != 0 {
		goto L1
	} else {
		goto L1348
	}
L1348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[110]))) = v5128
	if v5074 != 0 {
		goto L1350
	} else {
		goto L1351
	}
L1349:
	;
	v5136 = *(*int32)(unsafe.Add(mBase, uint32(v2252)+12))
	if v5136 != 0 {
		goto L1355
	} else {
		goto L1356
	}
L1350:
	;
	v5131 = F_cstring_to_text(m, v5074)
	mBase = m.M
	v5132 = m.ExcPending
	if v5132 != 0 {
		goto L1
	} else {
		goto L1353
	}
L1351:
	;
	goto L1352
L1352:
	;
	v5134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[111]))) = uint8(v5134)
	goto L1349
L1353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[100]))) = v5131
	goto L1349
L1354:
	;
	F_do_tup_output(m, v5117, v685+int32(_a_F_PostgresMainLoopOnce_70), v685+int32(_a_F_PostgresMainLoopOnce_79))
	mBase = m.M
	v5147 = m.ExcPending
	if v5147 != 0 {
		goto L1
	} else {
		goto L1359
	}
L1355:
	;
	v5137 = F_cstring_to_text(m, v5136)
	mBase = m.M
	v5138 = m.ExcPending
	if v5138 != 0 {
		goto L1
	} else {
		goto L1358
	}
L1356:
	;
	goto L1357
L1357:
	;
	v5140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[75]))) = uint8(v5140)
	goto L1354
L1358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+uint32(_c_F_PostgresMainLoopOnce[112]))) = v5137
	goto L1354
L1359:
	;
	F_end_tup_output(m, v5117)
	mBase = m.M
	v5149 = m.ExcPending
	if v5149 != 0 {
		goto L1
	} else {
		goto L1360
	}
L1360:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v5151 = m.ExcPending
	if v5151 != 0 {
		goto L1
	} else {
		goto L1361
	}
L1361:
	;
	v5162 = int32(_a_F_PostgresMainLoopOnce_87)
	goto L244
L1362:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v690
	v5194 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[51]))
	F_MemoryContextReset(m, v5194)
	mBase = m.M
	v5196 = m.ExcPending
	if v5196 != 0 {
		goto L1
	} else {
		goto L1363
	}
L1363:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L243
L1364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+176)) = int32(_a_F_PostgresMainLoopOnce_166)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_167), v685+int32(176))
	mBase = m.M
	v5249 = m.ExcPending
	if v5249 != 0 {
		goto L1
	} else {
		goto L1365
	}
L1365:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1268), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v5254 = m.ExcPending
	if v5254 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v685)+160)) = int32(_a_F_PostgresMainLoopOnce_166)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_168), v685+int32(160))
	mBase = m.M
	v5265 = m.ExcPending
	if v5265 != 0 {
		goto L1
	} else {
		goto L1368
	}
L1368:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1274), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v5270 = m.ExcPending
	if v5270 != 0 {
		goto L1
	} else {
		goto L1369
	}
L1369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+144)) = int32(_a_F_PostgresMainLoopOnce_166)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_169), v685+int32(144))
	mBase = m.M
	v5281 = m.ExcPending
	if v5281 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1371:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1279), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v5286 = m.ExcPending
	if v5286 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+112)) = int32(_a_F_PostgresMainLoopOnce_166)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_170), v685+int32(112))
	mBase = m.M
	v5297 = m.ExcPending
	if v5297 != 0 {
		goto L1
	} else {
		goto L1374
	}
L1374:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1285), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v5302 = m.ExcPending
	if v5302 != 0 {
		goto L1
	} else {
		goto L1375
	}
L1375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+128)) = int32(_a_F_PostgresMainLoopOnce_166)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_171), v685+int32(128))
	mBase = m.M
	v5313 = m.ExcPending
	if v5313 != 0 {
		goto L1
	} else {
		goto L1377
	}
L1377:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_14), int32(1291), int32(_a_F_PostgresMainLoopOnce_156))
	mBase = m.M
	v5318 = m.ExcPending
	if v5318 != 0 {
		goto L1
	} else {
		goto L1378
	}
L1378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1379:
	;
	goto L229
L1380:
	;
	v5369 = int32(_a_F_PostgresMainLoopOnce_172)
	v5370 = int32(0)
	v5374 = m.G0
	v5376 = v5374 - int32(16)
	m.G0 = v5376
	v5379 = int32(_a_F_PostgresMainLoopOnce_173)
	v5384 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v5370, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1384
L1381:
	;
	goto L1382
L1382:
	;
	F_start_xact_command(m)
	mBase = m.M
	v5419 = m.ExcPending
	if v5419 != 0 {
		goto L1
	} else {
		goto L1387
	}
L1383:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1382
L1384:
	;
	v5397 = F___memcpy(m, v5376, v5379, int32(16))
	mBase = m.M
	v5398 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5376))))
	v5399 = *(*int32)(unsafe.Add(mBase, uint32(v5376)+4))
	v5400 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v5400
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v5399
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v5398
	v5404 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5376)+8)))
	v5405 = *(*int32)(unsafe.Add(mBase, uint32(v5376)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v5400
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v5405
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v5404
	goto L1386
L1386:
	;
	v5412 = F___syscall_ret(m, v5370)
	mBase = m.M
	m.G0 = v5376 + int32(16)
	goto L1383
L1387:
	;
	v5421 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v5421 != 0 {
		goto L1388
	} else {
		goto L1389
	}
L1388:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])) = int32(0)
	F_DropCachedPlan(m, v5421)
	mBase = m.M
	v5426 = m.ExcPending
	if v5426 != 0 {
		goto L1
	} else {
		goto L1391
	}
L1389:
	;
	goto L1390
L1390:
	;
	v5427 = int32(_a_F_PostgresMainLoopOnce_125)
	v5428 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v5431 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5431
	v5434 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v5434 == int32(1) {
		goto L1392
	} else {
		goto L1393
	}
L1391:
	;
	goto L1390
L1392:
	;
	v5437 = int32(_a_F_PostgresMainLoopOnce_172)
	v5438 = int32(0)
	v5442 = m.G0
	v5444 = v5442 - int32(16)
	m.G0 = v5444
	v5447 = int32(_a_F_PostgresMainLoopOnce_173)
	v5452 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v5438, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1396
L1393:
	;
	goto L1394
L1394:
	;
	v5487 = F_raw_parser(m, v673, int32(0))
	mBase = m.M
	v5488 = m.ExcPending
	if v5488 != 0 {
		goto L1
	} else {
		goto L1399
	}
L1395:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1394
L1396:
	;
	v5465 = F___memcpy(m, v5444, v5447, int32(16))
	mBase = m.M
	v5466 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5444))))
	v5467 = *(*int32)(unsafe.Add(mBase, uint32(v5444)+4))
	v5468 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v5468
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v5467
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v5466
	v5472 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5444)+8)))
	v5473 = *(*int32)(unsafe.Add(mBase, uint32(v5444)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v5468
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v5473
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v5472
	goto L1398
L1398:
	;
	v5480 = F___syscall_ret(m, v5438)
	mBase = m.M
	m.G0 = v5444 + int32(16)
	goto L1395
L1399:
	;
	v5490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v5490 == int32(1) {
		goto L1400
	} else {
		goto L1401
	}
L1400:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_176))
	mBase = m.M
	v5495 = m.ExcPending
	if v5495 != 0 {
		goto L1
	} else {
		goto L1403
	}
L1401:
	;
	goto L1402
L1402:
	;
	v5497 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	switch v5497 {
	case 0:
		v5777 = v1
		goto L1407
	default:
		goto L1412
	case 3:
		goto L1411
	}
L1403:
	;
	goto L1402
L1404:
	;
	v6362 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L1570
L1405:
	;
	v6343 = v6306
	v6356 = int32(1)
	goto L1404
L1406:
	;
	v5831 = int32(0)
	v5832 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v5832 <= v5831 {
		v6343 = v5818
		v6356 = v5831
		goto L1404
	} else {
		goto L1439
	}
L1407:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5428
	if v5487 == int32(0) {
		v6306 = v5777
		goto L1405
	} else {
		goto L1438
	}
L1408:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1179), int32(_a_F_PostgresMainLoopOnce_177))
	mBase = m.M
	v5753 = m.ExcPending
	if v5753 != 0 {
		goto L1
	} else {
		goto L1437
	}
L1409:
	;
	v5705 = *(*int32)(unsafe.Add(mBase, uint32(v5696)+64))
	v5706 = *(*int32)(unsafe.Add(mBase, uint32(v5705)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5357)+48)) = v5706
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_178), v5357+int32(48))
	mBase = m.M
	v5712 = m.ExcPending
	if v5712 != 0 {
		goto L1
	} else {
		goto L1436
	}
L1410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5428
	v6306 = v1
	goto L1405
L1411:
	;
	v5628 = int32(1)
	v5631 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5632 = m.ExcPending
	if v5632 != 0 {
		goto L1
	} else {
		goto L1422
	}
L1412:
	;
	if v5487 == int32(0) {
		goto L1410
	} else {
		goto L1413
	}
L1413:
	;
	v5500 = int32(0)
	v5501 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v5500 < v5501 {
		goto L1414
	} else {
		goto L1415
	}
L1414:
	;
	v5504 = v5500
	goto L1417
L1415:
	;
	v5570 = v5501
	goto L1416
L1416:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v5428
	v5811 = v5570
	v5818 = v1
	goto L1406
L1417:
	;
	v5540 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v5544 = *(*int32)(unsafe.Add(mBase, uint32(v5540+v5504<<(uint(int32(2))%32))))
	v5545 = F_GetCommandLogLevel(m, v5544)
	mBase = m.M
	v5546 = m.ExcPending
	if v5546 != 0 {
		goto L1
	} else {
		goto L1419
	}
L1418:
	;
	v5570 = v5552
	goto L1416
L1419:
	;
	v5548 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	if base.Ui32(v5545) <= base.Ui32(v5548) {
		goto L1411
	} else {
		goto L1420
	}
L1420:
	;
	v5551 = v5504 + int32(1)
	v5552 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v5551 < v5552 {
		v5504 = v5551
		goto L1417
	} else {
		goto L1421
	}
L1421:
	;
	goto L1418
L1422:
	;
	if v5631 == int32(0) {
		v5777 = v5628
		goto L1407
	} else {
		goto L1423
	}
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5357)+64)) = v673
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_179), v5357-int32(-64))
	mBase = m.M
	v5640 = m.ExcPending
	if v5640 != 0 {
		goto L1
	} else {
		goto L1424
	}
L1424:
	;
	F_errhidestmt(m)
	mBase = m.M
	v5642 = m.ExcPending
	if v5642 != 0 {
		goto L1
	} else {
		goto L1425
	}
L1425:
	;
	if v5487 == int32(0) {
		goto L1408
	} else {
		goto L1426
	}
L1426:
	;
	v5645 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v5645 <= int32(0) {
		goto L1408
	} else {
		goto L1427
	}
L1427:
	;
	v5649 = int32(0)
	v5656 = v5645
	goto L1428
L1428:
	;
	v5685 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v5689 = *(*int32)(unsafe.Add(mBase, uint32(v5685+v5649<<(uint(int32(2))%32))))
	v5690 = *(*int32)(unsafe.Add(mBase, uint32(v5689)+4))
	v5691 = *(*int32)(unsafe.Add(mBase, uint32(v5690)))
	if v5691 == int32(253) {
		goto L1430
	} else {
		goto L1431
	}
L1429:
	;
	goto L1408
L1430:
	;
	v5694 = *(*int32)(unsafe.Add(mBase, uint32(v5690)+4))
	v5696 = F_FetchPreparedStatement(m, v5694, int32(0))
	mBase = m.M
	v5697 = m.ExcPending
	if v5697 != 0 {
		goto L1
	} else {
		goto L1433
	}
L1431:
	;
	v5699 = v5656
	goto L1432
L1432:
	;
	v5701 = v5649 + int32(1)
	if v5701 < v5699 {
		v5649 = v5701
		v5656 = v5699
		goto L1428
	} else {
		goto L1435
	}
L1433:
	;
	if v5696 != 0 {
		goto L1409
	} else {
		goto L1434
	}
L1434:
	;
	v5698 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	v5699 = v5698
	goto L1432
L1435:
	;
	goto L1429
L1436:
	;
	goto L1408
L1437:
	;
	v5777 = v5628
	goto L1407
L1438:
	;
	v5794 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	v5811 = v5794
	v5818 = v5777
	goto L1406
L1439:
	;
	v5849 = v1
	goto L1440
L1440:
	;
	v5871 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v5874 = v5871 + v5849<<(uint(int32(2))%32)
	v5875 = *(*int32)(unsafe.Add(mBase, uint32(v5874)))
	v5880 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v5880 == int32(0) {
		goto L1443
	} else {
		goto L1444
	}
L1441:
	;
	v6343 = v5818
	v6356 = int32(0)
	goto L1404
L1442:
	;
	v5918 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v5918 == int32(0) {
		goto L1448
	} else {
		goto L1449
	}
L1443:
	;
	goto L1442
L1444:
	;
	v5884 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v5884 != int32(1) {
		goto L1443
	} else {
		goto L1445
	}
L1445:
	;
	v5889 = *(*int64)(unsafe.Add(mBase, uint32(v5880)+392))
	if int32(0)&base.B2i32(v5889 != int64(0)) != 0 {
		goto L1443
	} else {
		goto L1446
	}
L1446:
	;
	v5893 = int32(_a_F_PostgresMainLoopOnce_180)
	v5895 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v5896 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v5895 + v5896
	v5899 = *(*int32)(unsafe.Add(mBase, uint32(v5880)))
	*(*int32)(unsafe.Add(mBase, uint32(v5880))) = v5899 + v5896
	*(*int64)(unsafe.Add(mBase, uint32(v5880)+392)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5880))) = v5899 + int32(2)
	v5910 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v5910 - v5896
	goto L1443
L1447:
	;
	v5952 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+4))
	v5953 = F_CreateCommandTag(m, v5952)
	mBase = m.M
	v5954 = m.ExcPending
	if v5954 != 0 {
		goto L1
	} else {
		goto L1452
	}
L1448:
	;
	goto L1447
L1449:
	;
	v5922 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v5922 != int32(1) {
		goto L1448
	} else {
		goto L1450
	}
L1450:
	;
	v5927 = *(*int64)(unsafe.Add(mBase, uint32(v5918)+400))
	if int32(0)&base.B2i32(v5927 != int64(0)) != 0 {
		goto L1448
	} else {
		goto L1451
	}
L1451:
	;
	v5931 = int32(_a_F_PostgresMainLoopOnce_180)
	v5933 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v5934 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v5933 + v5934
	v5937 = *(*int32)(unsafe.Add(mBase, uint32(v5918)))
	*(*int32)(unsafe.Add(mBase, uint32(v5918))) = v5937 + v5934
	*(*int64)(unsafe.Add(mBase, uint32(v5918)+400)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5918))) = v5937 + int32(2)
	v5948 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v5948 - v5934
	goto L1448
L1452:
	;
	v5961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5953<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[128]))))
	*(*int32)(unsafe.Add(mBase, uint32(v5357+int32(72)))) = v5961
	goto L1453
L1453:
	;
	v5968 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v5969 = *(*int32)(unsafe.Add(mBase, uint32(v5968)+24))
	goto L1455
L1454:
	;
	F_start_xact_command(m)
	mBase = m.M
	v6011 = m.ExcPending
	if v6011 != 0 {
		goto L1
	} else {
		goto L1466
	}
L1455:
	;
	if base.B2i32((v5969-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L1454
	} else {
		goto L1456
	}
L1456:
	;
	v5978 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+4))
	if v5978 == int32(0) {
		goto L1457
	} else {
		goto L1458
	}
L1457:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5994 = m.ExcPending
	if v5994 != 0 {
		goto L1
	} else {
		goto L1461
	}
L1458:
	;
	v5981 = *(*int32)(unsafe.Add(mBase, uint32(v5978)))
	if v5981 != int32(225) {
		goto L1457
	} else {
		goto L1459
	}
L1459:
	;
	v5984 = *(*int32)(unsafe.Add(mBase, uint32(v5978)+4))
	if (v5984-int32(2))&int32(-6) == int32(0) {
		goto L1454
	} else {
		goto L1460
	}
L1460:
	;
	goto L1457
L1461:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v5997 = m.ExcPending
	if v5997 != 0 {
		goto L1
	} else {
		goto L1462
	}
L1462:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L1
	} else {
		goto L1463
	}
L1463:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v6003 = m.ExcPending
	if v6003 != 0 {
		goto L1
	} else {
		goto L1464
	}
L1464:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1246), int32(_a_F_PostgresMainLoopOnce_177))
	mBase = m.M
	v6008 = m.ExcPending
	if v6008 != 0 {
		goto L1
	} else {
		goto L1465
	}
L1465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1466:
	;
	v6013 = base.B2i32(v5811 < int32(2))
	if v6013 == int32(0) {
		goto L1467
	} else {
		goto L1468
	}
L1467:
	;
	v6018 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v6019 = *(*int32)(unsafe.Add(mBase, uint32(v6018)+24))
	if v6019 == int32(1) {
		goto L1471
	} else {
		goto L1472
	}
L1468:
	;
	goto L1469
L1469:
	;
	v6025 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v6025 != 0 {
		goto L1474
	} else {
		goto L1475
	}
L1470:
	;
	goto L1469
L1471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6018)+24)) = int32(4)
	goto L1473
L1472:
	;
	goto L1473
L1473:
	;
	goto L1470
L1474:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6027 = m.ExcPending
	if v6027 != 0 {
		goto L1
	} else {
		goto L1477
	}
L1475:
	;
	goto L1476
L1476:
	;
	v6030 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+4))
	v6031 = *(*int32)(unsafe.Add(mBase, uint32(v6030)))
	switch v6031 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v6035 = int32(1)
		goto L1479
	default:
		goto L1480
	}
L1477:
	;
	goto L1476
L1478:
	;
	if v6035 != 0 {
		goto L1481
	} else {
		goto L1482
	}
L1479:
	;
	goto L1478
L1480:
	;
	v6035 = int32(0)
	goto L1479
L1481:
	;
	v6036 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L1
	} else {
		goto L1484
	}
L1482:
	;
	goto L1483
L1483:
	;
	v6040 = int32(0)
	v6042 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6044 = v5874 + int32(4)
	if v6044 == v6040 {
		v6059 = v6042
		v6060 = v6040
		goto L1486
	} else {
		goto L1487
	}
L1484:
	;
	F_PushActiveSnapshot(m, v6036)
	mBase = m.M
	v6039 = m.ExcPending
	if v6039 != 0 {
		goto L1
	} else {
		goto L1485
	}
L1485:
	;
	goto L1483
L1486:
	;
	v6061 = int32(_a_F_PostgresMainLoopOnce_125)
	v6062 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6059
	v6066 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v6066 == int32(1) {
		goto L1490
	} else {
		goto L1491
	}
L1487:
	;
	v6047 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v6048 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if base.Ui32(v6047+v6048<<(uint(int32(2))%32)) <= base.Ui32(v6044) {
		v6059 = v6042
		v6060 = v6040
		goto L1486
	} else {
		goto L1488
	}
L1488:
	;
	v6057 = F_AllocSetContextCreateInternal(m, v6042, int32(_a_F_PostgresMainLoopOnce_181), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v6058 = m.ExcPending
	if v6058 != 0 {
		goto L1
	} else {
		goto L1489
	}
L1489:
	;
	v6059 = v6057
	v6060 = v6057
	goto L1486
L1490:
	;
	v6069 = int32(_a_F_PostgresMainLoopOnce_172)
	v6070 = int32(0)
	v6074 = m.G0
	v6076 = v6074 - int32(16)
	m.G0 = v6076
	v6079 = int32(_a_F_PostgresMainLoopOnce_173)
	v6084 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v6070, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1494
L1491:
	;
	goto L1492
L1492:
	;
	v6118 = int32(0)
	v6121 = F_parse_analyze_fixedparams(m, v5875, v673, v6118, v6118, v6118)
	mBase = m.M
	v6122 = m.ExcPending
	if v6122 != 0 {
		goto L1
	} else {
		goto L1497
	}
L1493:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1492
L1494:
	;
	v6097 = F___memcpy(m, v6076, v6079, int32(16))
	mBase = m.M
	v6098 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6076))))
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(v6076)+4))
	v6100 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v6100
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v6099
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v6098
	v6104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6076)+8)))
	v6105 = *(*int32)(unsafe.Add(mBase, uint32(v6076)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v6100
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v6105
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v6104
	goto L1496
L1496:
	;
	v6112 = F___syscall_ret(m, v6070)
	mBase = m.M
	m.G0 = v6076 + int32(16)
	goto L1493
L1497:
	;
	v6124 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v6124 == int32(1) {
		goto L1498
	} else {
		goto L1499
	}
L1498:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_182))
	mBase = m.M
	v6129 = m.ExcPending
	if v6129 != 0 {
		goto L1
	} else {
		goto L1501
	}
L1499:
	;
	goto L1500
L1500:
	;
	v6130 = F_pg_rewrite_query(m, v6121)
	mBase = m.M
	v6131 = m.ExcPending
	if v6131 != 0 {
		goto L1
	} else {
		goto L1502
	}
L1501:
	;
	goto L1500
L1502:
	;
	v6134 = F_pg_plan_queries(m, v6130, v673, int32(2048), int32(0))
	mBase = m.M
	v6135 = m.ExcPending
	if v6135 != 0 {
		goto L1
	} else {
		goto L1503
	}
L1503:
	;
	if v6035 != 0 {
		goto L1504
	} else {
		goto L1505
	}
L1504:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v6137 = m.ExcPending
	if v6137 != 0 {
		goto L1
	} else {
		goto L1507
	}
L1505:
	;
	goto L1506
L1506:
	;
	v6139 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v6139 != 0 {
		goto L1508
	} else {
		goto L1509
	}
L1507:
	;
	goto L1506
L1508:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L1
	} else {
		goto L1511
	}
L1509:
	;
	goto L1510
L1510:
	;
	v6143 = int32(1)
	v6145 = F_CreatePortal(m, int32(_a_F_PostgresMainLoopOnce_183), v6143, v6143)
	mBase = m.M
	v6146 = m.ExcPending
	if v6146 != 0 {
		goto L1
	} else {
		goto L1512
	}
L1511:
	;
	goto L1510
L1512:
	;
	v6147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6145)+136)) = uint8(v6147)
	*(*int64)(unsafe.Add(mBase, uint32(v6145)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+40)) = v5953
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+32)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+4)) = v6147
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+60)) = v6147
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+56)) = v6134
	*(*int32)(unsafe.Add(mBase, uint32(v6145)+36)) = v5953
	goto L1513
L1513:
	;
	v6161 = int32(0)
	F_PortalStart(m, v6145, v6161, v6161, v6161)
	mBase = m.M
	v6165 = m.ExcPending
	if v6165 != 0 {
		goto L1
	} else {
		goto L1514
	}
L1514:
	;
	v6166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5357)+78)) = uint16(v6166)
	v6168 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+4))
	v6169 = *(*int32)(unsafe.Add(mBase, uint32(v6168)))
	if v6169 != int32(203) {
		goto L1515
	} else {
		goto L1516
	}
L1515:
	;
	F_PortalSetResultFormat(m, v6145, int32(1), v5357+int32(78))
	mBase = m.M
	v6190 = m.ExcPending
	if v6190 != 0 {
		goto L1
	} else {
		goto L1521
	}
L1516:
	;
	v6172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6168)+16)))
	if v6172 != 0 {
		goto L1515
	} else {
		goto L1517
	}
L1517:
	;
	v6173 = *(*int32)(unsafe.Add(mBase, uint32(v6168)+12))
	v6174 = F_GetPortalByName(m, v6173)
	mBase = m.M
	v6175 = m.ExcPending
	if v6175 != 0 {
		goto L1
	} else {
		goto L1518
	}
L1518:
	;
	if v6174 == int32(0) {
		goto L1515
	} else {
		goto L1519
	}
L1519:
	;
	v6178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6174)+76)))
	if v6178&int32(1) == int32(0) {
		goto L1515
	} else {
		goto L1520
	}
L1520:
	;
	v6183 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5357)+78)) = uint16(v6183)
	goto L1515
L1521:
	;
	v6191 = F_CreateDestReceiver(m, v5362)
	mBase = m.M
	v6192 = m.ExcPending
	if v6192 != 0 {
		goto L1
	} else {
		goto L1522
	}
L1522:
	;
	if v5362 == int32(2) {
		goto L1523
	} else {
		goto L1524
	}
L1523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6191)+20)) = v6145
	goto L1526
L1524:
	;
	goto L1525
L1525:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6062
	v6202 = F_PortalRun(m, v6145, int32(2147483647), int32(1), v6191, v6191, v5357+int32(80))
	mBase = m.M
	v6203 = m.ExcPending
	if v6203 != 0 {
		goto L1
	} else {
		goto L1527
	}
L1526:
	;
	goto L1525
L1527:
	;
	v6204 = *(*int32)(unsafe.Add(mBase, uint32(v6191)+12))
	m.T0[v6204].(func(*base.Module, int32))(m, v6191)
	mBase = m.M
	v6206 = m.ExcPending
	if v6206 != 0 {
		goto L1
	} else {
		goto L1528
	}
L1528:
	;
	F_PortalDrop(m, v6145, int32(0))
	mBase = m.M
	v6209 = m.ExcPending
	if v6209 != 0 {
		goto L1
	} else {
		goto L1529
	}
L1529:
	;
	if v6044 != 0 {
		goto L1533
	} else {
		goto L1534
	}
L1530:
	;
	F_EndCommand(m, v5357+int32(80), v5362)
	mBase = m.M
	v6275 = m.ExcPending
	if v6275 != 0 {
		goto L1
	} else {
		goto L1564
	}
L1531:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6268 = m.ExcPending
	if v6268 != 0 {
		goto L1
	} else {
		goto L1563
	}
L1532:
	;
	v6239 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+4))
	v6240 = *(*int32)(unsafe.Add(mBase, uint32(v6239)))
	if v6240 == int32(225) {
		goto L1550
	} else {
		goto L1551
	}
L1533:
	;
	v6210 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v6211 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if base.Ui32(v6044) < base.Ui32(v6210+v6211<<(uint(int32(2))%32)) {
		goto L1532
	} else {
		goto L1536
	}
L1534:
	;
	goto L1535
L1535:
	;
	if v6013 == int32(0) {
		goto L1537
	} else {
		goto L1538
	}
L1536:
	;
	goto L1535
L1537:
	;
	v6220 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v6221 = *(*int32)(unsafe.Add(mBase, uint32(v6220)+24))
	if v6221 == int32(4) {
		goto L1541
	} else {
		goto L1542
	}
L1538:
	;
	goto L1539
L1539:
	;
	v6231 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L1544
L1540:
	;
	goto L1539
L1541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6220)+24)) = int32(1)
	goto L1543
L1542:
	;
	goto L1543
L1543:
	;
	goto L1540
L1544:
	;
	if v6231 != 0 {
		goto L1545
	} else {
		goto L1546
	}
L1545:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6234 = m.ExcPending
	if v6234 != 0 {
		goto L1
	} else {
		goto L1548
	}
L1546:
	;
	goto L1547
L1547:
	;
	v6236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v6236 == int32(0) {
		goto L1530
	} else {
		goto L1549
	}
L1548:
	;
	goto L1547
L1549:
	;
	goto L1531
L1550:
	;
	v6248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L1553
L1551:
	;
	goto L1552
L1552:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v6255 = m.ExcPending
	if v6255 != 0 {
		goto L1
	} else {
		goto L1559
	}
L1553:
	;
	if v6248 != 0 {
		goto L1554
	} else {
		goto L1555
	}
L1554:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6251 = m.ExcPending
	if v6251 != 0 {
		goto L1
	} else {
		goto L1557
	}
L1555:
	;
	goto L1556
L1556:
	;
	v6253 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v6253 != 0 {
		goto L1531
	} else {
		goto L1558
	}
L1557:
	;
	goto L1556
L1558:
	;
	goto L1530
L1559:
	;
	v6261 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L1560
L1560:
	;
	if v6261 == int32(0) {
		goto L1530
	} else {
		goto L1561
	}
L1561:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6266 = m.ExcPending
	if v6266 != 0 {
		goto L1
	} else {
		goto L1562
	}
L1562:
	;
	goto L1530
L1563:
	;
	v6270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = uint8(v6270)
	goto L1530
L1564:
	;
	if v6060 != 0 {
		goto L1565
	} else {
		goto L1566
	}
L1565:
	;
	F_MemoryContextDelete(m, v6060)
	mBase = m.M
	v6277 = m.ExcPending
	if v6277 != 0 {
		goto L1
	} else {
		goto L1568
	}
L1566:
	;
	goto L1567
L1567:
	;
	v6279 = v5849 + int32(1)
	v6280 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v6279 < v6280 {
		v5849 = v6279
		goto L1440
	} else {
		goto L1569
	}
L1568:
	;
	goto L1567
L1569:
	;
	goto L1441
L1570:
	;
	if v6362 != 0 {
		goto L1571
	} else {
		goto L1572
	}
L1571:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6365 = m.ExcPending
	if v6365 != 0 {
		goto L1
	} else {
		goto L1574
	}
L1572:
	;
	goto L1573
L1573:
	;
	v6367 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v6367 != 0 {
		goto L1575
	} else {
		goto L1576
	}
L1574:
	;
	goto L1573
L1575:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6369 = m.ExcPending
	if v6369 != 0 {
		goto L1
	} else {
		goto L1578
	}
L1576:
	;
	goto L1577
L1577:
	;
	if v6356 != 0 {
		goto L1579
	} else {
		goto L1580
	}
L1578:
	;
	v6371 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = uint8(v6371)
	goto L1577
L1579:
	;
	F_NullCommand(m, v5362)
	mBase = m.M
	v6374 = m.ExcPending
	if v6374 != 0 {
		goto L1
	} else {
		goto L1582
	}
L1580:
	;
	goto L1581
L1581:
	;
	v6377 = F_check_log_duration(m, v5357+int32(80), v6343)
	mBase = m.M
	v6378 = m.ExcPending
	if v6378 != 0 {
		goto L1
	} else {
		goto L1587
	}
L1582:
	;
	goto L1581
L1583:
	;
	if v5364 != 0 {
		goto L1609
	} else {
		goto L1610
	}
L1584:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v6491, int32(_a_F_PostgresMainLoopOnce_177))
	mBase = m.M
	v6519 = m.ExcPending
	if v6519 != 0 {
		goto L1
	} else {
		goto L1608
	}
L1585:
	;
	v6398 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6399 = m.ExcPending
	if v6399 != 0 {
		goto L1
	} else {
		goto L1592
	}
L1586:
	;
	v6383 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6384 = m.ExcPending
	if v6384 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1587:
	;
	switch v6377 - int32(1) {
	case 0:
		goto L1586
	case 1:
		goto L1585
	default:
		goto L1583
	}
L1588:
	;
	if v6383 == int32(0) {
		goto L1583
	} else {
		goto L1589
	}
L1589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5357))) = v5357 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_184), v5357)
	mBase = m.M
	v6392 = m.ExcPending
	if v6392 != 0 {
		goto L1
	} else {
		goto L1590
	}
L1590:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6394 = m.ExcPending
	if v6394 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1591:
	;
	v6491 = int32(1471)
	goto L1584
L1592:
	;
	if v6398 == int32(0) {
		goto L1583
	} else {
		goto L1593
	}
L1593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5357)+36)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v5357)+32)) = v5357 + int32(80)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_185), v5357+int32(32))
	mBase = m.M
	v6410 = m.ExcPending
	if v6410 != 0 {
		goto L1
	} else {
		goto L1594
	}
L1594:
	;
	F_errhidestmt(m)
	mBase = m.M
	v6412 = m.ExcPending
	if v6412 != 0 {
		goto L1
	} else {
		goto L1595
	}
L1595:
	;
	v6413 = int32(1478)
	if v6356 != 0 {
		v6491 = v6413
		goto L1584
	} else {
		goto L1596
	}
L1596:
	;
	v6414 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	if v6414 <= int32(0) {
		v6491 = v6413
		goto L1584
	} else {
		goto L1597
	}
L1597:
	;
	v6418 = int32(0)
	v6425 = v6414
	goto L1598
L1598:
	;
	v6454 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+12))
	v6458 = *(*int32)(unsafe.Add(mBase, uint32(v6454+v6418<<(uint(int32(2))%32))))
	v6459 = *(*int32)(unsafe.Add(mBase, uint32(v6458)+4))
	v6460 = *(*int32)(unsafe.Add(mBase, uint32(v6459)))
	if v6460 == int32(253) {
		goto L1601
	} else {
		goto L1602
	}
L1599:
	;
	v6472 = *(*int32)(unsafe.Add(mBase, uint32(v6465)+64))
	v6473 = *(*int32)(unsafe.Add(mBase, uint32(v6472)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5357)+16)) = v6473
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_178), v5357+int32(16))
	mBase = m.M
	v6479 = m.ExcPending
	if v6479 != 0 {
		goto L1
	} else {
		goto L1607
	}
L1600:
	;
	goto L1599
L1601:
	;
	v6463 = *(*int32)(unsafe.Add(mBase, uint32(v6459)+4))
	v6465 = F_FetchPreparedStatement(m, v6463, int32(0))
	mBase = m.M
	v6466 = m.ExcPending
	if v6466 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1602:
	;
	v6468 = v6425
	goto L1603
L1603:
	;
	v6470 = v6418 + int32(1)
	if v6470 < v6468 {
		v6418 = v6470
		v6425 = v6468
		goto L1598
	} else {
		goto L1606
	}
L1604:
	;
	if v6465 != 0 {
		goto L1600
	} else {
		goto L1605
	}
L1605:
	;
	v6467 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+4))
	v6468 = v6467
	goto L1603
L1606:
	;
	v6491 = v6413
	goto L1584
L1607:
	;
	v6491 = v6413
	goto L1584
L1608:
	;
	goto L1583
L1609:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_186))
	mBase = m.M
	v6558 = m.ExcPending
	if v6558 != 0 {
		goto L1
	} else {
		goto L1612
	}
L1610:
	;
	goto L1611
L1611:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	m.G0 = v5357 + int32(112)
	goto L226
L1612:
	;
	goto L1611
L1613:
	;
	v6610 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v6610 < int32(0) {
		goto L1615
	} else {
		goto L1616
	}
L1614:
	;
	v6618 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v6619 = m.ExcPending
	if v6619 != 0 {
		goto L1
	} else {
		goto L1618
	}
L1615:
	;
	v6614 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v6614
	goto L1617
L1616:
	;
	goto L1617
L1617:
	;
	goto L1614
L1618:
	;
	v6622 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v6623 = m.ExcPending
	if v6623 != 0 {
		goto L1
	} else {
		goto L1619
	}
L1619:
	;
	v6627 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v6628 = m.ExcPending
	if v6628 != 0 {
		goto L1
	} else {
		goto L1620
	}
L1620:
	;
	if int32(0) < v6627 {
		goto L1621
	} else {
		goto L1622
	}
L1621:
	;
	v6634 = F_palloc(m, v6627<<(uint(int32(2))%32))
	mBase = m.M
	v6635 = m.ExcPending
	if v6635 != 0 {
		goto L1
	} else {
		goto L1624
	}
L1622:
	;
	v6687 = int32(0)
	goto L1623
L1623:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v6723 = m.ExcPending
	if v6723 != 0 {
		goto L1
	} else {
		goto L1629
	}
L1624:
	;
	v6638 = int32(0)
	goto L1625
L1625:
	;
	v6678 = F_pq_getmsgint(m, v39+int32(440), int32(4))
	mBase = m.M
	v6679 = m.ExcPending
	if v6679 != 0 {
		goto L1
	} else {
		goto L1627
	}
L1626:
	;
	v6687 = v6634
	goto L1623
L1627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6634+v6638<<(uint(int32(2))%32)))) = v6678
	v6682 = v6638 + int32(1)
	if v6682 != v6627 {
		v6638 = v6682
		goto L1625
	} else {
		goto L1628
	}
L1628:
	;
	goto L1626
L1629:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v6622
	*(*int32)(unsafe.Add(mBase, uint32(v39)+512)) = v6687
	*(*int32)(unsafe.Add(mBase, uint32(v39)+460)) = v6627
	v6729 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	F_pgstat_report_activity(m, int32(3), v6622)
	mBase = m.M
	if v6729 == int32(1) {
		goto L1630
	} else {
		goto L1631
	}
L1630:
	;
	v6734 = int32(_a_F_PostgresMainLoopOnce_172)
	v6735 = int32(0)
	v6739 = m.G0
	v6741 = v6739 - int32(16)
	m.G0 = v6741
	v6744 = int32(_a_F_PostgresMainLoopOnce_173)
	v6749 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v6735, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1634
L1631:
	;
	goto L1632
L1632:
	;
	v6785 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6786 = m.ExcPending
	if v6786 != 0 {
		goto L1
	} else {
		goto L1637
	}
L1633:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1632
L1634:
	;
	v6762 = F___memcpy(m, v6741, v6744, int32(16))
	mBase = m.M
	v6763 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6741))))
	v6764 = *(*int32)(unsafe.Add(mBase, uint32(v6741)+4))
	v6765 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v6765
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v6764
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v6763
	v6769 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6741)+8)))
	v6770 = *(*int32)(unsafe.Add(mBase, uint32(v6741)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v6765
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v6770
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v6769
	goto L1636
L1636:
	;
	v6777 = F___syscall_ret(m, v6735)
	mBase = m.M
	m.G0 = v6741 + int32(16)
	goto L1633
L1637:
	;
	if v6785 != 0 {
		goto L1638
	} else {
		goto L1639
	}
L1638:
	;
	v6787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6618))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v6622
	if v6787 != 0 {
		goto L1641
	} else {
		goto L1642
	}
L1639:
	;
	goto L1640
L1640:
	;
	F_start_xact_command(m)
	mBase = m.M
	v6804 = m.ExcPending
	if v6804 != 0 {
		goto L1
	} else {
		goto L1646
	}
L1641:
	;
	v6790 = v6618
	goto L1643
L1642:
	;
	v6790 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1643
L1643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v6790
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_188), v39-int32(-64))
	mBase = m.M
	v6796 = m.ExcPending
	if v6796 != 0 {
		goto L1
	} else {
		goto L1644
	}
L1644:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1526), int32(_a_F_PostgresMainLoopOnce_189))
	mBase = m.M
	v6801 = m.ExcPending
	if v6801 != 0 {
		goto L1
	} else {
		goto L1645
	}
L1645:
	;
	goto L1640
L1646:
	;
	v6805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6618))))
	if v6805 != 0 {
		goto L1648
	} else {
		goto L1649
	}
L1647:
	;
	v6826 = int32(_a_F_PostgresMainLoopOnce_125)
	v6827 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6824
	v6831 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v6831 == int32(1) {
		goto L1656
	} else {
		goto L1657
	}
L1648:
	;
	v6807 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6824 = v6807
	v6825 = int32(0)
	goto L1647
L1649:
	;
	goto L1650
L1650:
	;
	v6810 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v6810 != 0 {
		goto L1651
	} else {
		goto L1652
	}
L1651:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])) = int32(0)
	F_DropCachedPlan(m, v6810)
	mBase = m.M
	v6815 = m.ExcPending
	if v6815 != 0 {
		goto L1
	} else {
		goto L1654
	}
L1652:
	;
	goto L1653
L1653:
	;
	v6817 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6822 = F_AllocSetContextCreateInternal(m, v6817, int32(_a_F_PostgresMainLoopOnce_190), int32(0), int32(_a_F_PostgresMainLoopOnce_20), int32(_a_F_PostgresMainLoopOnce_21))
	mBase = m.M
	v6823 = m.ExcPending
	if v6823 != 0 {
		goto L1
	} else {
		goto L1655
	}
L1654:
	;
	goto L1653
L1655:
	;
	v6824 = v6822
	v6825 = v6822
	goto L1647
L1656:
	;
	v6834 = int32(_a_F_PostgresMainLoopOnce_172)
	v6835 = int32(0)
	v6839 = m.G0
	v6841 = v6839 - int32(16)
	m.G0 = v6841
	v6844 = int32(_a_F_PostgresMainLoopOnce_173)
	v6849 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v6835, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1660
L1657:
	;
	goto L1658
L1658:
	;
	v6884 = F_raw_parser(m, v6622, int32(0))
	mBase = m.M
	v6885 = m.ExcPending
	if v6885 != 0 {
		goto L1
	} else {
		goto L1663
	}
L1659:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1658
L1660:
	;
	v6862 = F___memcpy(m, v6841, v6844, int32(16))
	mBase = m.M
	v6863 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6841))))
	v6864 = *(*int32)(unsafe.Add(mBase, uint32(v6841)+4))
	v6865 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v6865
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v6864
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v6863
	v6869 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6841)+8)))
	v6870 = *(*int32)(unsafe.Add(mBase, uint32(v6841)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v6865
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v6870
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v6869
	goto L1662
L1662:
	;
	v6877 = F___syscall_ret(m, v6835)
	mBase = m.M
	m.G0 = v6841 + int32(16)
	goto L1659
L1663:
	;
	v6887 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[122])))
	if v6887 == int32(1) {
		goto L1664
	} else {
		goto L1665
	}
L1664:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_176))
	mBase = m.M
	v6892 = m.ExcPending
	if v6892 != 0 {
		goto L1
	} else {
		goto L1667
	}
L1665:
	;
	goto L1666
L1666:
	;
	if v6884 != 0 {
		goto L1669
	} else {
		goto L1670
	}
L1667:
	;
	goto L1666
L1668:
	;
	if v6825 != 0 {
		goto L1694
	} else {
		goto L1695
	}
L1669:
	;
	v6893 = *(*int32)(unsafe.Add(mBase, uint32(v6884)+4))
	if int32(2) <= v6893 {
		goto L200
	} else {
		goto L1672
	}
L1670:
	;
	goto L1671
L1671:
	;
	v6950 = int32(0)
	v6953 = F_CreateCachedPlan(m, v6950, v6622, v6950)
	mBase = m.M
	v6954 = m.ExcPending
	if v6954 != 0 {
		goto L1
	} else {
		goto L1693
	}
L1672:
	;
	v6896 = *(*int32)(unsafe.Add(mBase, uint32(v6884)+12))
	v6897 = *(*int32)(unsafe.Add(mBase, uint32(v6896)))
	v6899 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v6900 = *(*int32)(unsafe.Add(mBase, uint32(v6899)+24))
	goto L1673
L1673:
	;
	v6907 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+4))
	if (v6900-int32(7))&int32(-9) == int32(0) {
		goto L1674
	} else {
		goto L1675
	}
L1674:
	;
	if v6907 == int32(0) {
		goto L199
	} else {
		goto L1677
	}
L1675:
	;
	goto L1676
L1676:
	;
	v6918 = F_CreateCommandTag(m, v6907)
	mBase = m.M
	v6919 = m.ExcPending
	if v6919 != 0 {
		goto L1
	} else {
		goto L1680
	}
L1677:
	;
	v6910 = *(*int32)(unsafe.Add(mBase, uint32(v6907)))
	if v6910 != int32(225) {
		goto L199
	} else {
		goto L1678
	}
L1678:
	;
	v6913 = *(*int32)(unsafe.Add(mBase, uint32(v6907)+4))
	if (v6913-int32(2))&int32(-6) != 0 {
		goto L199
	} else {
		goto L1679
	}
L1679:
	;
	goto L1676
L1680:
	;
	v6920 = F_CreateCachedPlan(m, v6897, v6622, v6918)
	mBase = m.M
	v6921 = m.ExcPending
	if v6921 != 0 {
		goto L1
	} else {
		goto L1681
	}
L1681:
	;
	v6924 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+4))
	v6925 = *(*int32)(unsafe.Add(mBase, uint32(v6924)))
	switch v6925 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v6929 = int32(1)
		goto L1683
	default:
		goto L1684
	}
L1682:
	;
	if v6929 == int32(0) {
		goto L1685
	} else {
		goto L1686
	}
L1683:
	;
	goto L1682
L1684:
	;
	v6929 = int32(0)
	goto L1683
L1685:
	;
	v6936 = F_pg_analyze_and_rewrite_varparams(m, v6897, v6622, v39+int32(512), v39+int32(460))
	mBase = m.M
	v6937 = m.ExcPending
	if v6937 != 0 {
		goto L1
	} else {
		goto L1688
	}
L1686:
	;
	goto L1687
L1687:
	;
	v6938 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v6939 = m.ExcPending
	if v6939 != 0 {
		goto L1
	} else {
		goto L1689
	}
L1688:
	;
	v6955 = v6920
	v6956 = v6936
	goto L1668
L1689:
	;
	F_PushActiveSnapshot(m, v6938)
	mBase = m.M
	v6941 = m.ExcPending
	if v6941 != 0 {
		goto L1
	} else {
		goto L1690
	}
L1690:
	;
	v6946 = F_pg_analyze_and_rewrite_varparams(m, v6897, v6622, v39+int32(512), v39+int32(460))
	mBase = m.M
	v6947 = m.ExcPending
	if v6947 != 0 {
		goto L1
	} else {
		goto L1691
	}
L1691:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v6949 = m.ExcPending
	if v6949 != 0 {
		goto L1
	} else {
		goto L1692
	}
L1692:
	;
	v6955 = v6920
	v6956 = v6946
	goto L1668
L1693:
	;
	v6955 = v6953
	v6956 = v6950
	goto L1668
L1694:
	;
	v6958 = *(*int32)(unsafe.Add(mBase, uint32(v6955)+56))
	v6960 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	v6964 = *(*int32)(unsafe.Add(mBase, uint32(v6958)+16))
	if v6964 != v6960 {
		goto L1698
	} else {
		goto L1699
	}
L1695:
	;
	goto L1696
L1696:
	;
	v6994 = *(*int32)(unsafe.Add(mBase, uint32(v39)+512))
	v6995 = *(*int32)(unsafe.Add(mBase, uint32(v39)+460))
	v6996 = int32(0)
	F_CompleteCachedPlan(m, v6955, v6956, v6825, v6994, v6995, v6996, v6996, int32(2048), int32(1))
	mBase = m.M
	v7001 = m.ExcPending
	if v7001 != 0 {
		goto L1
	} else {
		goto L1714
	}
L1697:
	;
	goto L1696
L1698:
	;
	if v6964 == int32(0) {
		goto L1701
	} else {
		goto L1702
	}
L1699:
	;
	goto L1700
L1700:
	;
	goto L1697
L1701:
	;
	if v6960 != 0 {
		goto L1708
	} else {
		goto L1709
	}
L1702:
	;
	v6968 = *(*int32)(unsafe.Add(mBase, uint32(v6958)+28))
	v6969 = *(*int32)(unsafe.Add(mBase, uint32(v6958)+24))
	if v6969 != 0 {
		goto L1704
	} else {
		goto L1705
	}
L1703:
	;
	if v6968 == int32(0) {
		goto L1701
	} else {
		goto L1707
	}
L1704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6969)+28)) = v6968
	goto L1703
L1705:
	;
	goto L1706
L1706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6964)+20)) = v6968
	goto L1703
L1707:
	;
	v6974 = *(*int32)(unsafe.Add(mBase, uint32(v6958)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6968)+24)) = v6974
	goto L1701
L1708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6958)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6958)+16)) = v6960
	v6981 = *(*int32)(unsafe.Add(mBase, uint32(v6960)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6958)+28)) = v6981
	if v6981 != 0 {
		goto L1711
	} else {
		goto L1712
	}
L1709:
	;
	goto L1710
L1710:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6958)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6958)+16)) = int32(0)
	goto L1700
L1711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6981)+24)) = v6958
	goto L1713
L1712:
	;
	goto L1713
L1713:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6960)+20)) = v6958
	goto L1697
L1714:
	;
	v7003 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v7003 != 0 {
		goto L1715
	} else {
		goto L1716
	}
L1715:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v7005 = m.ExcPending
	if v7005 != 0 {
		goto L1
	} else {
		goto L1718
	}
L1716:
	;
	goto L1717
L1717:
	;
	if v6805 != 0 {
		goto L1720
	} else {
		goto L1721
	}
L1718:
	;
	goto L1717
L1719:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v6827
	F_CommandCounterIncrement(m)
	mBase = m.M
	v7016 = m.ExcPending
	if v7016 != 0 {
		goto L1
	} else {
		goto L1725
	}
L1720:
	;
	F_StorePreparedStatement(m, v6618, v6955, int32(0))
	mBase = m.M
	v7008 = m.ExcPending
	if v7008 != 0 {
		goto L1
	} else {
		goto L1723
	}
L1721:
	;
	goto L1722
L1722:
	;
	F_SaveCachedPlan(m, v6955)
	mBase = m.M
	v7010 = m.ExcPending
	if v7010 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1723:
	;
	goto L1719
L1724:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])) = v6955
	goto L1719
L1725:
	;
	v7018 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v7018 == int32(2) {
		goto L1726
	} else {
		goto L1727
	}
L1726:
	;
	F_pq_putemptymessage(m, int32(49))
	mBase = m.M
	v7023 = m.ExcPending
	if v7023 != 0 {
		goto L1
	} else {
		goto L1729
	}
L1727:
	;
	goto L1728
L1728:
	;
	v7027 = F_check_log_duration(m, v39+int32(480), int32(0))
	mBase = m.M
	v7028 = m.ExcPending
	if v7028 != 0 {
		goto L1
	} else {
		goto L1734
	}
L1729:
	;
	goto L1728
L1730:
	;
	if v6729 != 0 {
		goto L1746
	} else {
		goto L1747
	}
L1731:
	;
	F_errhidestmt(m)
	mBase = m.M
	v7069 = m.ExcPending
	if v7069 != 0 {
		goto L1
	} else {
		goto L1744
	}
L1732:
	;
	v7048 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7049 = m.ExcPending
	if v7049 != 0 {
		goto L1
	} else {
		goto L1738
	}
L1733:
	;
	v7033 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7034 = m.ExcPending
	if v7034 != 0 {
		goto L1
	} else {
		goto L1735
	}
L1734:
	;
	switch v7027 - int32(1) {
	case 0:
		goto L1733
	case 1:
		goto L1732
	default:
		goto L1730
	}
L1735:
	;
	if v7033 == int32(0) {
		goto L1730
	} else {
		goto L1736
	}
L1736:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_184), v39+int32(32))
	mBase = m.M
	v7044 = m.ExcPending
	if v7044 != 0 {
		goto L1
	} else {
		goto L1737
	}
L1737:
	;
	v7067 = int32(1707)
	goto L1731
L1738:
	;
	if v7048 == int32(0) {
		goto L1730
	} else {
		goto L1739
	}
L1739:
	;
	v7052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6618))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v6622
	if v7052 != 0 {
		goto L1740
	} else {
		goto L1741
	}
L1740:
	;
	v7055 = v6618
	goto L1742
L1741:
	;
	v7055 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1742
L1742:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v7055
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_191), v39+int32(48))
	mBase = m.M
	v7064 = m.ExcPending
	if v7064 != 0 {
		goto L1
	} else {
		goto L1743
	}
L1743:
	;
	v7067 = int32(1715)
	goto L1731
L1744:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v7067, int32(_a_F_PostgresMainLoopOnce_189))
	mBase = m.M
	v7073 = m.ExcPending
	if v7073 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1745:
	;
	goto L1730
L1746:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_192))
	mBase = m.M
	v7077 = m.ExcPending
	if v7077 != 0 {
		goto L1
	} else {
		goto L1749
	}
L1747:
	;
	goto L1748
L1748:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L202
L1749:
	;
	goto L1748
L1750:
	;
	v7086 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v7086 < int32(0) {
		goto L1752
	} else {
		goto L1753
	}
L1751:
	;
	v7093 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	v7096 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v7097 = m.ExcPending
	if v7097 != 0 {
		goto L1
	} else {
		goto L1755
	}
L1752:
	;
	v7090 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v7090
	goto L1754
L1753:
	;
	goto L1754
L1754:
	;
	goto L1751
L1755:
	;
	v7100 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v7101 = m.ExcPending
	if v7101 != 0 {
		goto L1
	} else {
		goto L1756
	}
L1756:
	;
	v7104 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v7105 = m.ExcPending
	if v7105 != 0 {
		goto L1
	} else {
		goto L1757
	}
L1757:
	;
	if v7104 != 0 {
		goto L1758
	} else {
		goto L1759
	}
L1758:
	;
	v7106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7096))))
	v7108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100))))
	if v7108 != 0 {
		goto L1761
	} else {
		goto L1762
	}
L1759:
	;
	goto L1760
L1760:
	;
	v7125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100))))
	if v7125 != 0 {
		goto L1770
	} else {
		goto L1771
	}
L1761:
	;
	v7109 = v7100
	goto L1763
L1762:
	;
	v7109 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1763
L1763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+212)) = v7109
	if v7106 != 0 {
		goto L1764
	} else {
		goto L1765
	}
L1764:
	;
	v7112 = v7096
	goto L1766
L1765:
	;
	v7112 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1766
L1766:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+208)) = v7112
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_193), v39+int32(208))
	mBase = m.M
	v7118 = m.ExcPending
	if v7118 != 0 {
		goto L1
	} else {
		goto L1767
	}
L1767:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1761), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v7123 = m.ExcPending
	if v7123 != 0 {
		goto L1
	} else {
		goto L1768
	}
L1768:
	;
	goto L1760
L1769:
	;
	v7136 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v7136
	F_pgstat_report_activity(m, int32(3), v7136)
	mBase = m.M
	v7140 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+60))
	if v7140 == int32(0) {
		goto L1775
	} else {
		goto L1776
	}
L1770:
	;
	v7127 = F_FetchPreparedStatement(m, v7100, int32(1))
	mBase = m.M
	v7128 = m.ExcPending
	if v7128 != 0 {
		goto L1
	} else {
		goto L1773
	}
L1771:
	;
	goto L1772
L1772:
	;
	v7131 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v7131 == int32(0) {
		goto L197
	} else {
		goto L1774
	}
L1773:
	;
	v7129 = *(*int32)(unsafe.Add(mBase, uint32(v7127)+64))
	v7134 = v7129
	goto L1769
L1774:
	;
	v7134 = v7131
	goto L1769
L1775:
	;
	if v7093&int32(1) != 0 {
		goto L1789
	} else {
		goto L1790
	}
L1776:
	;
	v7143 = *(*int32)(unsafe.Add(mBase, uint32(v7140)+4))
	if v7143 <= int32(0) {
		goto L1775
	} else {
		goto L1777
	}
L1777:
	;
	v7146 = *(*int32)(unsafe.Add(mBase, uint32(v7140)+12))
	v7150 = int32(0)
	goto L1778
L1778:
	;
	v7187 = *(*int32)(unsafe.Add(mBase, uint32(v7146+v7150<<(uint(int32(2))%32))))
	v7188 = *(*int64)(unsafe.Add(mBase, uint32(v7187)+16))
	if v7188 == int64(0) {
		goto L1780
	} else {
		goto L1781
	}
L1779:
	;
	v7197 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v7197 == int32(0) {
		goto L1785
	} else {
		goto L1786
	}
L1780:
	;
	v7192 = v7150 + int32(1)
	if v7192 != v7143 {
		v7150 = v7192
		goto L1778
	} else {
		goto L1783
	}
L1781:
	;
	goto L1782
L1782:
	;
	goto L1779
L1783:
	;
	goto L1775
L1784:
	;
	goto L1775
L1785:
	;
	goto L1784
L1786:
	;
	v7201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v7201 != int32(1) {
		goto L1785
	} else {
		goto L1787
	}
L1787:
	;
	v7206 = *(*int64)(unsafe.Add(mBase, uint32(v7197)+392))
	if int32(1)&base.B2i32(v7206 != int64(0)) != 0 {
		goto L1785
	} else {
		goto L1788
	}
L1788:
	;
	v7210 = int32(_a_F_PostgresMainLoopOnce_180)
	v7212 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v7213 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v7212 + v7213
	v7216 = *(*int32)(unsafe.Add(mBase, uint32(v7197)))
	*(*int32)(unsafe.Add(mBase, uint32(v7197))) = v7216 + v7213
	*(*int64)(unsafe.Add(mBase, uint32(v7197)+392)) = v7188
	*(*int32)(unsafe.Add(mBase, uint32(v7197))) = v7216 + int32(2)
	v7227 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v7227 - v7213
	goto L1785
L1789:
	;
	v7269 = int32(_a_F_PostgresMainLoopOnce_172)
	v7270 = int32(0)
	v7274 = m.G0
	v7276 = v7274 - int32(16)
	m.G0 = v7276
	v7279 = int32(_a_F_PostgresMainLoopOnce_173)
	v7284 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v7270, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L1793
L1790:
	;
	goto L1791
L1791:
	;
	F_start_xact_command(m)
	mBase = m.M
	v7319 = m.ExcPending
	if v7319 != 0 {
		goto L1
	} else {
		goto L1796
	}
L1792:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L1791
L1793:
	;
	v7297 = F___memcpy(m, v7276, v7279, int32(16))
	mBase = m.M
	v7298 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7276))))
	v7299 = *(*int32)(unsafe.Add(mBase, uint32(v7276)+4))
	v7300 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v7300
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v7299
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v7298
	v7304 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7276)+8)))
	v7305 = *(*int32)(unsafe.Add(mBase, uint32(v7276)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v7300
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v7305
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v7304
	goto L1795
L1795:
	;
	v7312 = F___syscall_ret(m, v7270)
	mBase = m.M
	m.G0 = v7276 + int32(16)
	goto L1792
L1796:
	;
	v7322 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7322
	v7327 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v7328 = m.ExcPending
	if v7328 != 0 {
		goto L1
	} else {
		goto L1797
	}
L1797:
	;
	if int32(0) < v7327 {
		goto L1798
	} else {
		goto L1799
	}
L1798:
	;
	v7334 = F_palloc(m, v7327<<(uint(int32(1))%32))
	mBase = m.M
	v7335 = m.ExcPending
	if v7335 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1799:
	;
	v7390 = v1
	goto L1800
L1800:
	;
	v7423 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v7424 = m.ExcPending
	if v7424 != 0 {
		goto L1
	} else {
		goto L1806
	}
L1801:
	;
	v7338 = int32(0)
	goto L1802
L1802:
	;
	v7378 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v7379 = m.ExcPending
	if v7379 != 0 {
		goto L1
	} else {
		goto L1804
	}
L1803:
	;
	v7390 = v7334
	goto L1800
L1804:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7334+v7338<<(uint(int32(1))%32)))) = uint16(v7378)
	v7382 = v7338 + int32(1)
	if v7382 != v7327 {
		v7338 = v7382
		goto L1802
	} else {
		goto L1805
	}
L1805:
	;
	goto L1803
L1806:
	;
	if base.B2i32(v7423 != v7327)&base.B2i32(int32(2) <= v7327) != 0 {
		goto L196
	} else {
		goto L1807
	}
L1807:
	;
	v7429 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+24))
	if v7423 != v7429 {
		goto L195
	} else {
		goto L1808
	}
L1808:
	;
	v7432 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v7433 = *(*int32)(unsafe.Add(mBase, uint32(v7432)+24))
	goto L1809
L1809:
	;
	if (v7433-int32(7))&int32(-9) == int32(0) {
		goto L1810
	} else {
		goto L1811
	}
L1810:
	;
	v7440 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+4))
	if v7440 == int32(0) {
		goto L194
	} else {
		goto L1813
	}
L1811:
	;
	goto L1812
L1812:
	;
	v7456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7096))))
	if v7456 == int32(0) {
		goto L1818
	} else {
		goto L1819
	}
L1813:
	;
	v7443 = *(*int32)(unsafe.Add(mBase, uint32(v7440)+4))
	if v7443 == int32(0) {
		goto L194
	} else {
		goto L1814
	}
L1814:
	;
	v7446 = *(*int32)(unsafe.Add(mBase, uint32(v7443)))
	if v7446 != int32(225) {
		goto L194
	} else {
		goto L1815
	}
L1815:
	;
	v7449 = *(*int32)(unsafe.Add(mBase, uint32(v7443)+4))
	if (v7449-int32(2))&int32(-6)|v7423 != 0 {
		goto L194
	} else {
		goto L1816
	}
L1816:
	;
	goto L1812
L1817:
	;
	v7468 = int32(_a_F_PostgresMainLoopOnce_125)
	v7469 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7471 = *(*int32)(unsafe.Add(mBase, uint32(v7467)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7471
	v7473 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+12))
	v7474 = F_pstrdup(m, v7473)
	mBase = m.M
	v7475 = m.ExcPending
	if v7475 != 0 {
		goto L1
	} else {
		goto L1823
	}
L1818:
	;
	v7459 = int32(1)
	v7461 = F_CreatePortal(m, v7096, v7459, v7459)
	mBase = m.M
	v7462 = m.ExcPending
	if v7462 != 0 {
		goto L1
	} else {
		goto L1821
	}
L1819:
	;
	goto L1820
L1820:
	;
	v7463 = int32(0)
	v7465 = F_CreatePortal(m, v7096, v7463, v7463)
	mBase = m.M
	v7466 = m.ExcPending
	if v7466 != 0 {
		goto L1
	} else {
		goto L1822
	}
L1821:
	;
	v7467 = v7461
	goto L1817
L1822:
	;
	v7467 = v7465
	goto L1817
L1823:
	;
	v7476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100))))
	if v7476 != 0 {
		goto L1824
	} else {
		goto L1825
	}
L1824:
	;
	v7477 = F_pstrdup(m, v7100)
	mBase = m.M
	v7478 = m.ExcPending
	if v7478 != 0 {
		goto L1
	} else {
		goto L1827
	}
L1825:
	;
	v7479 = v1
	goto L1826
L1826:
	;
	if v7423 <= int32(0) {
		goto L1830
	} else {
		goto L1831
	}
L1827:
	;
	v7479 = v7477
	goto L1826
L1828:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7469
	v7790 = *(*int32)(unsafe.Add(mBase, uint32(v7467)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+464)) = v7774
	*(*int32)(unsafe.Add(mBase, uint32(v39)+460)) = v7790
	v7793 = int32(_a_F_PostgresMainLoopOnce_195)
	v7794 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v39 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+516)) = int32(1162)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+512)) = v7794
	*(*int32)(unsafe.Add(mBase, uint32(v39)+520)) = v39 + int32(460)
	v7809 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v7810 = m.ExcPending
	if v7810 != 0 {
		goto L1
	} else {
		goto L1894
	}
L1829:
	;
	v7774 = v7737
	v7787 = int32(1)
	goto L1828
L1830:
	;
	v7482 = int32(0)
	v7483 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+4))
	if v7483 == v7482 {
		v7774 = v1
		v7787 = v7482
		goto L1828
	} else {
		goto L1833
	}
L1831:
	;
	goto L1832
L1832:
	;
	v7501 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7502 = m.ExcPending
	if v7502 != 0 {
		goto L1
	} else {
		goto L1840
	}
L1833:
	;
	v7489 = *(*int32)(unsafe.Add(mBase, uint32(v7483)+4))
	v7490 = *(*int32)(unsafe.Add(mBase, uint32(v7489)))
	switch v7490 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v7494 = int32(1)
		goto L1835
	default:
		goto L1836
	}
L1834:
	;
	if v7494 == int32(0) {
		v7774 = v1
		v7787 = int32(0)
		goto L1828
	} else {
		goto L1837
	}
L1835:
	;
	goto L1834
L1836:
	;
	v7494 = int32(0)
	goto L1835
L1837:
	;
	v7497 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v7498 = m.ExcPending
	if v7498 != 0 {
		goto L1
	} else {
		goto L1838
	}
L1838:
	;
	F_PushActiveSnapshot(m, v7497)
	mBase = m.M
	v7500 = m.ExcPending
	if v7500 != 0 {
		goto L1
	} else {
		goto L1839
	}
L1839:
	;
	v7737 = v1
	goto L1829
L1840:
	;
	F_PushActiveSnapshot(m, v7501)
	mBase = m.M
	v7504 = m.ExcPending
	if v7504 != 0 {
		goto L1
	} else {
		goto L1841
	}
L1841:
	;
	v7505 = *(*int32)(unsafe.Add(mBase, uint32(v7467)))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+464)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+460)) = v7505
	v7509 = int32(_a_F_PostgresMainLoopOnce_195)
	v7510 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v39 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+516)) = int32(1161)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+512)) = v7510
	*(*int32)(unsafe.Add(mBase, uint32(v39)+520)) = v39 + int32(460)
	v7523 = F_makeParamList(m, v7423)
	mBase = m.M
	v7524 = m.ExcPending
	if v7524 != 0 {
		goto L1
	} else {
		goto L1842
	}
L1842:
	;
	v7527 = int32(0)
	v7532 = v7527
	v7547 = v1
	goto L1843
L1843:
	;
	v7567 = v7532 << (uint(int32(2)) % 32)
	v7568 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+20))
	v7570 = *(*int32)(unsafe.Add(mBase, uint32(v7567+v7568)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+468)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+464)) = v7532
	v7577 = F_pq_getmsgint(m, v39+int32(440), int32(4))
	mBase = m.M
	v7578 = m.ExcPending
	if v7578 != 0 {
		goto L1
	} else {
		goto L1846
	}
L1844:
	;
	v7702 = int32(_a_F_PostgresMainLoopOnce_195)
	v7704 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v7705 = *(*int32)(unsafe.Add(mBase, uint32(v7704)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v7705
	v7708 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v7708 == int32(0) {
		v7737 = v7523
		goto L1829
	} else {
		goto L1892
	}
L1845:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+480)) = v7594
	if int32(2) <= v7327 {
		goto L1855
	} else {
		goto L1856
	}
L1846:
	;
	v7580 = base.B2i32(v7577 == int32(-1))
	if v7577 == int32(-1) {
		goto L1847
	} else {
		goto L1848
	}
L1847:
	;
	v7581 = int32(0)
	v7594 = v7581
	v7595 = v7581
	goto L1845
L1848:
	;
	goto L1849
L1849:
	;
	v7585 = F_pq_getmsgbytes(m, v39+int32(440), v7577)
	mBase = m.M
	v7586 = m.ExcPending
	if v7586 != 0 {
		goto L1
	} else {
		goto L1850
	}
L1850:
	;
	v7587 = v7585 + v7577
	v7588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7587))))
	v7589 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7587))) = uint8(v7589)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+488)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+484)) = v7577
	v7594 = v7585
	v7595 = v7588
	goto L1845
L1851:
	;
	if v7580 == int32(0) {
		goto L1888
	} else {
		goto L1889
	}
L1852:
	;
	F_getTypeBinaryInputInfo(m, v7570, v39+int32(472), v39+int32(456))
	mBase = m.M
	v7668 = m.ExcPending
	if v7668 != 0 {
		goto L1
	} else {
		goto L1881
	}
L1853:
	;
	F_getTypeInputInfo(m, v7570, v39+int32(472), v39+int32(456))
	mBase = m.M
	v7611 = m.ExcPending
	if v7611 != 0 {
		goto L1
	} else {
		goto L1859
	}
L1854:
	;
	v7604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7603))))
	switch v7604 {
	case 0:
		goto L1853
	case 1:
		goto L1852
	default:
		goto L192
	}
L1855:
	;
	v7603 = v7390 + v7532<<(uint(int32(1))%32)
	goto L1854
L1856:
	;
	goto L1857
L1857:
	;
	if v7327 <= v7527 {
		goto L1853
	} else {
		goto L1858
	}
L1858:
	;
	v7603 = v7390
	goto L1854
L1859:
	;
	if v7577 == int32(-1) {
		goto L1860
	} else {
		goto L1861
	}
L1860:
	;
	v7616 = int32(0)
	goto L1862
L1861:
	;
	v7613 = *(*int32)(unsafe.Add(mBase, uint32(v39)+480))
	v7614 = F_pg_client_to_server(m, v7613, v7577)
	mBase = m.M
	v7615 = m.ExcPending
	if v7615 != 0 {
		goto L1
	} else {
		goto L1863
	}
L1862:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+468)) = v7616
	v7618 = *(*int32)(unsafe.Add(mBase, uint32(v39)+472))
	v7619 = *(*int32)(unsafe.Add(mBase, uint32(v39)+456))
	v7621 = F_OidInputFunctionCall(m, v7618, v7616, v7619, int32(-1))
	mBase = m.M
	v7622 = m.ExcPending
	if v7622 != 0 {
		goto L1
	} else {
		goto L1864
	}
L1863:
	;
	v7616 = v7614
	goto L1862
L1864:
	;
	v7623 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+468)) = v7623
	if v7616 == v7623 {
		v7683 = v7621
		v7684 = v7547
		goto L1851
	} else {
		goto L1865
	}
L1865:
	;
	v7628 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	if v7628 != 0 {
		goto L1866
	} else {
		goto L1867
	}
L1866:
	;
	v7629 = int32(_a_F_PostgresMainLoopOnce_125)
	v7630 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	v7633 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7633
	if v7547 == int32(0) {
		goto L1870
	} else {
		goto L1871
	}
L1867:
	;
	v7657 = v7547
	goto L1868
L1868:
	;
	v7659 = *(*int32)(unsafe.Add(mBase, uint32(v39)+480))
	if v7616 == v7659 {
		v7683 = v7621
		v7684 = v7657
		goto L1851
	} else {
		goto L1879
	}
L1869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7567+v7642))) = v7651
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v7630
	v7657 = v7642
	goto L1868
L1870:
	;
	v7637 = F_palloc0(m, v7423<<(uint(int32(2))%32))
	mBase = m.M
	v7638 = m.ExcPending
	if v7638 != 0 {
		goto L1
	} else {
		goto L1873
	}
L1871:
	;
	v7641 = v7628
	v7642 = v7547
	goto L1872
L1872:
	;
	if v7641 < int32(0) {
		goto L1874
	} else {
		goto L1875
	}
L1873:
	;
	v7640 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[131]))
	v7641 = v7640
	v7642 = v7637
	goto L1872
L1874:
	;
	v7645 = F_pstrdup(m, v7616)
	mBase = m.M
	v7646 = m.ExcPending
	if v7646 != 0 {
		goto L1
	} else {
		goto L1877
	}
L1875:
	;
	goto L1876
L1876:
	;
	v7649 = F_pnstrdup(m, v7616, v7641+int32(8))
	mBase = m.M
	v7650 = m.ExcPending
	if v7650 != 0 {
		goto L1
	} else {
		goto L1878
	}
L1877:
	;
	v7651 = v7645
	goto L1869
L1878:
	;
	v7651 = v7649
	goto L1869
L1879:
	;
	F_pfree(m, v7616)
	mBase = m.M
	v7662 = m.ExcPending
	if v7662 != 0 {
		goto L1
	} else {
		goto L1880
	}
L1880:
	;
	v7683 = v7621
	v7684 = v7657
	goto L1851
L1881:
	;
	v7669 = *(*int32)(unsafe.Add(mBase, uint32(v39)+472))
	if v7577 == int32(-1) {
		goto L1882
	} else {
		goto L1883
	}
L1882:
	;
	v7673 = int32(0)
	goto L1884
L1883:
	;
	v7673 = v39 + int32(480)
	goto L1884
L1884:
	;
	v7674 = *(*int32)(unsafe.Add(mBase, uint32(v39)+456))
	v7676 = F_OidReceiveFunctionCall(m, v7669, v7673, v7674, int32(-1))
	mBase = m.M
	v7677 = m.ExcPending
	if v7677 != 0 {
		goto L1
	} else {
		goto L1885
	}
L1885:
	;
	if v7577 == int32(-1) {
		v7683 = v7676
		v7684 = v7547
		goto L1851
	} else {
		goto L1886
	}
L1886:
	;
	v7678 = *(*int32)(unsafe.Add(mBase, uint32(v39)+492))
	v7679 = *(*int32)(unsafe.Add(mBase, uint32(v39)+484))
	if v7678 != v7679 {
		goto L193
	} else {
		goto L1887
	}
L1887:
	;
	v7683 = v7676
	v7684 = v7547
	goto L1851
L1888:
	;
	v7688 = *(*int32)(unsafe.Add(mBase, uint32(v39)+480))
	*(*uint8)(unsafe.Add(mBase, uint32(v7688+v7577))) = uint8(v7595)
	goto L1890
L1889:
	;
	goto L1890
L1890:
	;
	v7693 = v7523 + int32(32) + v7532*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v7693)+8)) = v7570
	v7695 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7693)+6)) = uint16(v7695)
	*(*uint8)(unsafe.Add(mBase, uint32(v7693)+4)) = uint8(v7580)
	*(*int32)(unsafe.Add(mBase, uint32(v7693))) = v7683
	v7700 = v7532 + v7695
	if v7700 != v7423 {
		v7532 = v7700
		v7547 = v7684
		goto L1843
	} else {
		goto L1891
	}
L1891:
	;
	goto L1844
L1892:
	;
	v7711 = F_BuildParamLogString(m, v7523, v7684, v7708)
	mBase = m.M
	v7712 = m.ExcPending
	if v7712 != 0 {
		goto L1
	} else {
		goto L1893
	}
L1893:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7523)+24)) = v7711
	v7737 = v7523
	goto L1829
L1894:
	;
	if int32(0) < v7809 {
		goto L1895
	} else {
		goto L1896
	}
L1895:
	;
	v7816 = F_palloc(m, v7809<<(uint(int32(1))%32))
	mBase = m.M
	v7817 = m.ExcPending
	if v7817 != 0 {
		goto L1
	} else {
		goto L1898
	}
L1896:
	;
	v7869 = int32(0)
	goto L1897
L1897:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v7905 = m.ExcPending
	if v7905 != 0 {
		goto L1
	} else {
		goto L1903
	}
L1898:
	;
	v7820 = int32(0)
	goto L1899
L1899:
	;
	v7860 = F_pq_getmsgint(m, v39+int32(440), int32(2))
	mBase = m.M
	v7861 = m.ExcPending
	if v7861 != 0 {
		goto L1
	} else {
		goto L1901
	}
L1900:
	;
	v7869 = v7816
	goto L1897
L1901:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7816+v7820<<(uint(int32(1))%32)))) = uint16(v7860)
	v7864 = v7820 + int32(1)
	if v7864 != v7809 {
		v7820 = v7864
		goto L1899
	} else {
		goto L1902
	}
L1902:
	;
	goto L1900
L1903:
	;
	v7906 = int32(0)
	v7908 = F_GetCachedPlan(m, v7134, v7774, v7906, v7906)
	mBase = m.M
	v7909 = m.ExcPending
	if v7909 != 0 {
		goto L1
	} else {
		goto L1904
	}
L1904:
	;
	v7910 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+16))
	v7911 = *(*int32)(unsafe.Add(mBase, uint32(v7908)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v7467)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+40)) = v7910
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+32)) = v7474
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+4)) = v7479
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+60)) = v7908
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+56)) = v7911
	*(*int32)(unsafe.Add(mBase, uint32(v7467)+36)) = v7910
	goto L1905
L1905:
	;
	v7922 = *(*int32)(unsafe.Add(mBase, uint32(v7467)+56))
	if v7922 == int32(0) {
		goto L1906
	} else {
		goto L1907
	}
L1906:
	;
	if v7787 != 0 {
		goto L1920
	} else {
		goto L1921
	}
L1907:
	;
	v7925 = *(*int32)(unsafe.Add(mBase, uint32(v7922)+4))
	if v7925 <= int32(0) {
		goto L1906
	} else {
		goto L1908
	}
L1908:
	;
	v7928 = *(*int32)(unsafe.Add(mBase, uint32(v7922)+12))
	v7932 = int32(0)
	goto L1909
L1909:
	;
	v7969 = *(*int32)(unsafe.Add(mBase, uint32(v7928+v7932<<(uint(int32(2))%32))))
	v7970 = *(*int64)(unsafe.Add(mBase, uint32(v7969)+16))
	if v7970 == int64(0) {
		goto L1911
	} else {
		goto L1912
	}
L1910:
	;
	v7979 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v7979 == int32(0) {
		goto L1916
	} else {
		goto L1917
	}
L1911:
	;
	v7974 = v7932 + int32(1)
	if v7974 != v7925 {
		v7932 = v7974
		goto L1909
	} else {
		goto L1914
	}
L1912:
	;
	goto L1913
L1913:
	;
	goto L1910
L1914:
	;
	goto L1906
L1915:
	;
	goto L1906
L1916:
	;
	goto L1915
L1917:
	;
	v7983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v7983 != int32(1) {
		goto L1916
	} else {
		goto L1918
	}
L1918:
	;
	v7988 = *(*int64)(unsafe.Add(mBase, uint32(v7979)+400))
	if int32(1)&base.B2i32(v7988 != int64(0)) != 0 {
		goto L1916
	} else {
		goto L1919
	}
L1919:
	;
	v7992 = int32(_a_F_PostgresMainLoopOnce_180)
	v7994 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v7995 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v7994 + v7995
	v7998 = *(*int32)(unsafe.Add(mBase, uint32(v7979)))
	*(*int32)(unsafe.Add(mBase, uint32(v7979))) = v7998 + v7995
	*(*int64)(unsafe.Add(mBase, uint32(v7979)+400)) = v7970
	*(*int32)(unsafe.Add(mBase, uint32(v7979))) = v7998 + int32(2)
	v8009 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v8009 - v7995
	goto L1916
L1920:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v8050 = m.ExcPending
	if v8050 != 0 {
		goto L1
	} else {
		goto L1923
	}
L1921:
	;
	goto L1922
L1922:
	;
	v8051 = int32(0)
	F_PortalStart(m, v7467, v7774, v8051, v8051)
	mBase = m.M
	v8054 = m.ExcPending
	if v8054 != 0 {
		goto L1
	} else {
		goto L1924
	}
L1923:
	;
	goto L1922
L1924:
	;
	F_PortalSetResultFormat(m, v7467, v7809, v7869)
	mBase = m.M
	v8056 = m.ExcPending
	if v8056 != 0 {
		goto L1
	} else {
		goto L1925
	}
L1925:
	;
	v8057 = int32(_a_F_PostgresMainLoopOnce_195)
	v8059 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v8060 = *(*int32)(unsafe.Add(mBase, uint32(v8059)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v8060
	v8063 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v8063 == int32(2) {
		goto L1926
	} else {
		goto L1927
	}
L1926:
	;
	F_pq_putemptymessage(m, int32(50))
	mBase = m.M
	v8068 = m.ExcPending
	if v8068 != 0 {
		goto L1
	} else {
		goto L1929
	}
L1927:
	;
	goto L1928
L1928:
	;
	v8072 = F_check_log_duration(m, v39+int32(480), int32(0))
	mBase = m.M
	v8073 = m.ExcPending
	if v8073 != 0 {
		goto L1
	} else {
		goto L1934
	}
L1929:
	;
	goto L1928
L1930:
	;
	if v7093&int32(1) != 0 {
		goto L1960
	} else {
		goto L1961
	}
L1931:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v8147, int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v8152 = m.ExcPending
	if v8152 != 0 {
		goto L1
	} else {
		goto L1959
	}
L1932:
	;
	v8095 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8096 = m.ExcPending
	if v8096 != 0 {
		goto L1
	} else {
		goto L1939
	}
L1933:
	;
	v8078 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8079 = m.ExcPending
	if v8079 != 0 {
		goto L1
	} else {
		goto L1935
	}
L1934:
	;
	switch v8072 - int32(1) {
	case 0:
		goto L1933
	case 1:
		goto L1932
	default:
		goto L1930
	}
L1935:
	;
	if v8078 == int32(0) {
		goto L1930
	} else {
		goto L1936
	}
L1936:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_184), v39+int32(96))
	mBase = m.M
	v8089 = m.ExcPending
	if v8089 != 0 {
		goto L1
	} else {
		goto L1937
	}
L1937:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8091 = m.ExcPending
	if v8091 != 0 {
		goto L1
	} else {
		goto L1938
	}
L1938:
	;
	v8147 = int32(2185)
	goto L1931
L1939:
	;
	if v8095 == int32(0) {
		goto L1930
	} else {
		goto L1940
	}
L1940:
	;
	v8099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100))))
	v8100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7096))))
	v8101 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v8101
	if v8100 != 0 {
		goto L1941
	} else {
		goto L1942
	}
L1941:
	;
	v8104 = v7096
	goto L1943
L1942:
	;
	v8104 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1943
L1943:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+140)) = v8104
	if v8100 != 0 {
		goto L1944
	} else {
		goto L1945
	}
L1944:
	;
	v8108 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L1946
L1945:
	;
	v8108 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L1946
L1946:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = v8108
	if v8099 != 0 {
		goto L1947
	} else {
		goto L1948
	}
L1947:
	;
	v8111 = v7100
	goto L1949
L1948:
	;
	v8111 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1949
L1949:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v8111
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_197), v39+int32(128))
	mBase = m.M
	v8120 = m.ExcPending
	if v8120 != 0 {
		goto L1
	} else {
		goto L1950
	}
L1950:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8122 = m.ExcPending
	if v8122 != 0 {
		goto L1
	} else {
		goto L1951
	}
L1951:
	;
	v8123 = int32(2196)
	if v7774 == int32(0) {
		v8147 = v8123
		goto L1931
	} else {
		goto L1952
	}
L1952:
	;
	v8126 = *(*int32)(unsafe.Add(mBase, uint32(v7774)+28))
	if v8126 <= int32(0) {
		v8147 = v8123
		goto L1931
	} else {
		goto L1953
	}
L1953:
	;
	v8130 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132]))
	if v8130 == int32(0) {
		v8147 = v8123
		goto L1931
	} else {
		goto L1954
	}
L1954:
	;
	v8134 = F_BuildParamLogString(m, v7774, int32(0), v8130)
	mBase = m.M
	v8135 = m.ExcPending
	if v8135 != 0 {
		goto L1
	} else {
		goto L1955
	}
L1955:
	;
	if v8134 == int32(0) {
		v8147 = v8123
		goto L1931
	} else {
		goto L1956
	}
L1956:
	;
	v8138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8134))))
	if v8138 == int32(0) {
		v8147 = v8123
		goto L1931
	} else {
		goto L1957
	}
L1957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v8134
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_198), v39+int32(112))
	mBase = m.M
	v8146 = m.ExcPending
	if v8146 != 0 {
		goto L1
	} else {
		goto L1958
	}
L1958:
	;
	v8147 = v8123
	goto L1931
L1959:
	;
	goto L1930
L1960:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_199))
	mBase = m.M
	v8159 = m.ExcPending
	if v8159 != 0 {
		goto L1
	} else {
		goto L1963
	}
L1961:
	;
	goto L1962
L1962:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L202
L1963:
	;
	goto L1962
L1964:
	;
	v8168 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v8168 < int32(0) {
		goto L1966
	} else {
		goto L1967
	}
L1965:
	;
	v8176 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v8177 = m.ExcPending
	if v8177 != 0 {
		goto L1
	} else {
		goto L1969
	}
L1966:
	;
	v8172 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v8172
	goto L1968
L1967:
	;
	goto L1968
L1968:
	;
	goto L1965
L1969:
	;
	v8181 = F_pq_getmsgint(m, v39+int32(440), int32(4))
	mBase = m.M
	v8182 = m.ExcPending
	if v8182 != 0 {
		goto L1
	} else {
		goto L1970
	}
L1970:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v8186 = m.ExcPending
	if v8186 != 0 {
		goto L1
	} else {
		goto L1971
	}
L1971:
	;
	v8188 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	v8190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[47])))
	v8191 = F_GetPortalByName(m, v8176)
	mBase = m.M
	v8192 = m.ExcPending
	if v8192 != 0 {
		goto L1
	} else {
		goto L1972
	}
L1972:
	;
	if v8191 == int32(0) {
		goto L190
	} else {
		goto L1973
	}
L1973:
	;
	if v8188 == int32(2) {
		goto L1974
	} else {
		goto L1975
	}
L1974:
	;
	v8198 = int32(3)
	goto L1976
L1975:
	;
	v8198 = v8188
	goto L1976
L1976:
	;
	v8199 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+36))
	if v8199 == int32(0) {
		goto L1977
	} else {
		goto L1978
	}
L1977:
	;
	F_NullCommand(m, v8198)
	mBase = m.M
	v8203 = m.ExcPending
	if v8203 != 0 {
		goto L1
	} else {
		goto L1980
	}
L1978:
	;
	goto L1979
L1979:
	;
	v8204 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+56))
	if v8204 == int32(0) {
		v8223 = v1
		goto L1981
	} else {
		goto L1982
	}
L1980:
	;
	goto L202
L1981:
	;
	v8224 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+32))
	v8225 = F_pstrdup(m, v8224)
	mBase = m.M
	v8226 = m.ExcPending
	if v8226 != 0 {
		goto L1
	} else {
		goto L1988
	}
L1982:
	;
	v8207 = *(*int32)(unsafe.Add(mBase, uint32(v8204)+4))
	if v8207 != int32(1) {
		v8223 = v1
		goto L1981
	} else {
		goto L1983
	}
L1983:
	;
	v8210 = *(*int32)(unsafe.Add(mBase, uint32(v8204)+12))
	v8211 = *(*int32)(unsafe.Add(mBase, uint32(v8210)))
	v8212 = *(*int32)(unsafe.Add(mBase, uint32(v8211)+4))
	if v8212 == int32(6) {
		goto L1984
	} else {
		goto L1985
	}
L1984:
	;
	v8216 = *(*int32)(unsafe.Add(mBase, uint32(v8211)+88))
	v8217 = *(*int32)(unsafe.Add(mBase, uint32(v8216)))
	if v8217 == int32(225) {
		v8223 = int32(1)
		goto L1981
	} else {
		goto L1987
	}
L1985:
	;
	goto L1986
L1986:
	;
	v8223 = int32(0)
	goto L1981
L1987:
	;
	goto L1986
L1988:
	;
	v8227 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+4))
	if v8227 != 0 {
		goto L1989
	} else {
		goto L1990
	}
L1989:
	;
	v8228 = F_pstrdup(m, v8227)
	mBase = m.M
	v8229 = m.ExcPending
	if v8229 != 0 {
		goto L1
	} else {
		goto L1992
	}
L1990:
	;
	v8231 = int32(_a_F_PostgresMainLoopOnce_187)
	goto L1991
L1991:
	;
	v8232 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = v8225
	F_pgstat_report_activity(m, int32(3), v8225)
	mBase = m.M
	v8237 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+56))
	if v8237 == int32(0) {
		goto L1993
	} else {
		goto L1994
	}
L1992:
	;
	v8231 = v8228
	goto L1991
L1993:
	;
	v8467 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+36))
	v8474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8467<<(uint(int32(3))%32))+uint32(_c_F_PostgresMainLoopOnce[128]))))
	*(*int32)(unsafe.Add(mBase, uint32(v39+int32(456)))) = v8474
	goto L2027
L1994:
	;
	v8240 = *(*int32)(unsafe.Add(mBase, uint32(v8237)+4))
	if v8240 <= int32(0) {
		goto L1993
	} else {
		goto L1995
	}
L1995:
	;
	v8243 = int32(0)
	if v8243 < v8240 {
		goto L1996
	} else {
		goto L1997
	}
L1996:
	;
	v8246 = v8240
	goto L1998
L1997:
	;
	v8246 = v8243
	goto L1998
L1998:
	;
	v8247 = *(*int32)(unsafe.Add(mBase, uint32(v8237)+12))
	v8251 = int32(0)
	goto L2000
L1999:
	;
	if v8340 <= int32(0) {
		goto L1993
	} else {
		goto L2015
	}
L2000:
	;
	v8288 = *(*int32)(unsafe.Add(mBase, uint32(v8247+v8251<<(uint(int32(2))%32))))
	v8289 = *(*int64)(unsafe.Add(mBase, uint32(v8288)+8))
	if v8289 == int64(0) {
		goto L2002
	} else {
		goto L2003
	}
L2001:
	;
	v8298 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v8298 == int32(0) {
		goto L2007
	} else {
		goto L2008
	}
L2002:
	;
	v8293 = v8251 + int32(1)
	if v8293 != v8240 {
		v8251 = v8293
		goto L2000
	} else {
		goto L2005
	}
L2003:
	;
	goto L2004
L2004:
	;
	goto L2001
L2005:
	;
	v8340 = v8240
	v8342 = v8237
	v8343 = v8246
	goto L1999
L2006:
	;
	v8332 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+56))
	if v8332 == int32(0) {
		goto L1993
	} else {
		goto L2011
	}
L2007:
	;
	goto L2006
L2008:
	;
	v8302 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v8302 != int32(1) {
		goto L2007
	} else {
		goto L2009
	}
L2009:
	;
	v8307 = *(*int64)(unsafe.Add(mBase, uint32(v8298)+392))
	if int32(1)&base.B2i32(v8307 != int64(0)) != 0 {
		goto L2007
	} else {
		goto L2010
	}
L2010:
	;
	v8311 = int32(_a_F_PostgresMainLoopOnce_180)
	v8313 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v8314 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v8313 + v8314
	v8317 = *(*int32)(unsafe.Add(mBase, uint32(v8298)))
	*(*int32)(unsafe.Add(mBase, uint32(v8298))) = v8317 + v8314
	*(*int64)(unsafe.Add(mBase, uint32(v8298)+392)) = v8289
	*(*int32)(unsafe.Add(mBase, uint32(v8298))) = v8317 + int32(2)
	v8328 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v8328 - v8314
	goto L2007
L2011:
	;
	v8335 = *(*int32)(unsafe.Add(mBase, uint32(v8332)+4))
	v8336 = int32(0)
	if v8336 < v8335 {
		goto L2012
	} else {
		goto L2013
	}
L2012:
	;
	v8339 = v8335
	goto L2014
L2013:
	;
	v8339 = v8336
	goto L2014
L2014:
	;
	v8340 = v8335
	v8342 = v8332
	v8343 = v8339
	goto L1999
L2015:
	;
	v8346 = *(*int32)(unsafe.Add(mBase, uint32(v8342)+12))
	v8350 = int32(0)
	goto L2016
L2016:
	;
	v8387 = *(*int32)(unsafe.Add(mBase, uint32(v8346+v8350<<(uint(int32(2))%32))))
	v8388 = *(*int64)(unsafe.Add(mBase, uint32(v8387)+16))
	if v8388 == int64(0) {
		goto L2018
	} else {
		goto L2019
	}
L2017:
	;
	v8397 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[125]))
	if v8397 == int32(0) {
		goto L2023
	} else {
		goto L2024
	}
L2018:
	;
	v8392 = v8350 + int32(1)
	if v8343 != v8392 {
		v8350 = v8392
		goto L2016
	} else {
		goto L2021
	}
L2019:
	;
	goto L2020
L2020:
	;
	goto L2017
L2021:
	;
	goto L1993
L2022:
	;
	goto L1993
L2023:
	;
	goto L2022
L2024:
	;
	v8401 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[126])))
	if v8401 != int32(1) {
		goto L2023
	} else {
		goto L2025
	}
L2025:
	;
	v8406 = *(*int64)(unsafe.Add(mBase, uint32(v8397)+400))
	if int32(1)&base.B2i32(v8406 != int64(0)) != 0 {
		goto L2023
	} else {
		goto L2026
	}
L2026:
	;
	v8410 = int32(_a_F_PostgresMainLoopOnce_180)
	v8412 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	v8413 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v8412 + v8413
	v8416 = *(*int32)(unsafe.Add(mBase, uint32(v8397)))
	*(*int32)(unsafe.Add(mBase, uint32(v8397))) = v8416 + v8413
	*(*int64)(unsafe.Add(mBase, uint32(v8397)+400)) = v8388
	*(*int32)(unsafe.Add(mBase, uint32(v8397))) = v8416 + int32(2)
	v8427 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[127])) = v8427 - v8413
	goto L2023
L2027:
	;
	if v8190&int32(1) != 0 {
		goto L2028
	} else {
		goto L2029
	}
L2028:
	;
	v8482 = int32(_a_F_PostgresMainLoopOnce_172)
	v8483 = int32(0)
	v8487 = m.G0
	v8489 = v8487 - int32(16)
	m.G0 = v8489
	v8492 = int32(_a_F_PostgresMainLoopOnce_173)
	v8497 = F___memset(m, int32(_a_F_PostgresMainLoopOnce_174), v8483, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[113])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[114])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = int64(1)
	goto L2032
L2029:
	;
	goto L2030
L2030:
	;
	v8532 = F_CreateDestReceiver(m, v8198)
	mBase = m.M
	v8533 = m.ExcPending
	if v8533 != 0 {
		goto L1
	} else {
		goto L2035
	}
L2031:
	;
	F___gettimeofday(m, int32(_a_F_PostgresMainLoopOnce_175))
	mBase = m.M
	goto L2030
L2032:
	;
	v8510 = F___memcpy(m, v8489, v8492, int32(16))
	mBase = m.M
	v8511 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8489))))
	v8512 = *(*int32)(unsafe.Add(mBase, uint32(v8489)+4))
	v8513 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[117])) = v8513
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[118])) = v8512
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[119])) = v8511
	v8517 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8489)+8)))
	v8518 = *(*int32)(unsafe.Add(mBase, uint32(v8489)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[120])) = v8513
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[115])) = v8518
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[116])) = v8517
	goto L2034
L2034:
	;
	v8525 = F___syscall_ret(m, v8483)
	mBase = m.M
	m.G0 = v8489 + int32(16)
	goto L2031
L2035:
	;
	if v8198 == int32(3) {
		goto L2036
	} else {
		goto L2037
	}
L2036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8532)+20)) = v8191
	goto L2039
L2037:
	;
	goto L2038
L2038:
	;
	F_start_xact_command(m)
	mBase = m.M
	v8538 = m.ExcPending
	if v8538 != 0 {
		goto L1
	} else {
		goto L2040
	}
L2039:
	;
	goto L2038
L2040:
	;
	v8539 = int32(0)
	v8540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8191)+116)))
	v8542 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	switch v8542 {
	case 0:
		v8698 = v8539
		goto L2041
	default:
		goto L2043
	case 3:
		goto L2042
	}
L2041:
	;
	v8733 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8734 = *(*int32)(unsafe.Add(mBase, uint32(v8733)+24))
	goto L2075
L2042:
	;
	v8638 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8639 = m.ExcPending
	if v8639 != 0 {
		goto L1
	} else {
		goto L2051
	}
L2043:
	;
	v8543 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+56))
	if v8543 == int32(0) {
		v8698 = v8539
		goto L2041
	} else {
		goto L2044
	}
L2044:
	;
	v8546 = *(*int32)(unsafe.Add(mBase, uint32(v8543)+4))
	if v8546 <= int32(0) {
		v8698 = v8539
		goto L2041
	} else {
		goto L2045
	}
L2045:
	;
	v8551 = v8539
	goto L2046
L2046:
	;
	v8585 = *(*int32)(unsafe.Add(mBase, uint32(v8543)+12))
	v8589 = *(*int32)(unsafe.Add(mBase, uint32(v8585+v8551<<(uint(int32(2))%32))))
	v8590 = F_GetCommandLogLevel(m, v8589)
	mBase = m.M
	v8591 = m.ExcPending
	if v8591 != 0 {
		goto L1
	} else {
		goto L2048
	}
L2047:
	;
	v8698 = int32(0)
	goto L2041
L2048:
	;
	v8593 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	if base.Ui32(v8590) <= base.Ui32(v8593) {
		goto L2042
	} else {
		goto L2049
	}
L2049:
	;
	v8596 = v8551 + int32(1)
	v8597 = *(*int32)(unsafe.Add(mBase, uint32(v8543)+4))
	if v8596 < v8597 {
		v8551 = v8596
		goto L2046
	} else {
		goto L2050
	}
L2050:
	;
	goto L2047
L2051:
	;
	if v8638 == int32(0) {
		goto L2052
	} else {
		goto L2053
	}
L2052:
	;
	v8698 = int32(1)
	goto L2041
L2053:
	;
	goto L2054
L2054:
	;
	v8643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8176))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+336)) = v8225
	*(*int32)(unsafe.Add(mBase, uint32(v39)+324)) = v8231
	if v8643 != 0 {
		goto L2055
	} else {
		goto L2056
	}
L2055:
	;
	v8647 = v8176
	goto L2057
L2056:
	;
	v8647 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L2057
L2057:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+332)) = v8647
	if v8643 != 0 {
		goto L2058
	} else {
		goto L2059
	}
L2058:
	;
	v8651 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L2060
L2059:
	;
	v8651 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L2060
L2060:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+328)) = v8651
	v8653 = int32(1)
	if v8540&v8653 != 0 {
		goto L2061
	} else {
		goto L2062
	}
L2061:
	;
	v8658 = int32(_a_F_PostgresMainLoopOnce_200)
	goto L2063
L2062:
	;
	v8658 = int32(_a_F_PostgresMainLoopOnce_201)
	goto L2063
L2063:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+320)) = v8658
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_202), v39+int32(320))
	mBase = m.M
	v8664 = m.ExcPending
	if v8664 != 0 {
		goto L1
	} else {
		goto L2064
	}
L2064:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8666 = m.ExcPending
	if v8666 != 0 {
		goto L1
	} else {
		goto L2065
	}
L2065:
	;
	if v8232 == int32(0) {
		goto L2066
	} else {
		goto L2067
	}
L2066:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2346), int32(_a_F_PostgresMainLoopOnce_203))
	mBase = m.M
	v8695 = m.ExcPending
	if v8695 != 0 {
		goto L1
	} else {
		goto L2074
	}
L2067:
	;
	v8669 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+28))
	if v8669 <= int32(0) {
		goto L2066
	} else {
		goto L2068
	}
L2068:
	;
	v8673 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132]))
	if v8673 == int32(0) {
		goto L2066
	} else {
		goto L2069
	}
L2069:
	;
	v8677 = F_BuildParamLogString(m, v8232, int32(0), v8673)
	mBase = m.M
	v8678 = m.ExcPending
	if v8678 != 0 {
		goto L1
	} else {
		goto L2070
	}
L2070:
	;
	if v8677 == int32(0) {
		goto L2066
	} else {
		goto L2071
	}
L2071:
	;
	v8681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8677))))
	if v8681 == int32(0) {
		goto L2066
	} else {
		goto L2072
	}
L2072:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+304)) = v8677
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_198), v39+int32(304))
	mBase = m.M
	v8689 = m.ExcPending
	if v8689 != 0 {
		goto L1
	} else {
		goto L2073
	}
L2073:
	;
	goto L2066
L2074:
	;
	v8698 = v8653
	goto L2041
L2075:
	;
	if (v8734-int32(7))&int32(-9) == int32(0) {
		goto L2076
	} else {
		goto L2077
	}
L2076:
	;
	v8741 = *(*int32)(unsafe.Add(mBase, uint32(v8191)+56))
	if v8741 == int32(0) {
		goto L189
	} else {
		goto L2079
	}
L2077:
	;
	goto L2078
L2078:
	;
	v8765 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v8765 != 0 {
		goto L2085
	} else {
		goto L2086
	}
L2079:
	;
	v8744 = *(*int32)(unsafe.Add(mBase, uint32(v8741)+4))
	if v8744 != int32(1) {
		goto L189
	} else {
		goto L2080
	}
L2080:
	;
	v8747 = *(*int32)(unsafe.Add(mBase, uint32(v8741)+12))
	v8748 = *(*int32)(unsafe.Add(mBase, uint32(v8747)))
	v8749 = *(*int32)(unsafe.Add(mBase, uint32(v8748)+4))
	if v8749 != int32(6) {
		goto L189
	} else {
		goto L2081
	}
L2081:
	;
	v8752 = *(*int32)(unsafe.Add(mBase, uint32(v8748)+88))
	if v8752 == int32(0) {
		goto L189
	} else {
		goto L2082
	}
L2082:
	;
	v8755 = *(*int32)(unsafe.Add(mBase, uint32(v8752)))
	if v8755 != int32(225) {
		goto L189
	} else {
		goto L2083
	}
L2083:
	;
	v8758 = *(*int32)(unsafe.Add(mBase, uint32(v8752)+4))
	if (v8758-int32(2))&int32(-6) != 0 {
		goto L189
	} else {
		goto L2084
	}
L2084:
	;
	goto L2078
L2085:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v8767 = m.ExcPending
	if v8767 != 0 {
		goto L1
	} else {
		goto L2088
	}
L2086:
	;
	goto L2087
L2087:
	;
	v8768 = *(*int32)(unsafe.Add(mBase, uint32(v8191)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+476)) = v8232
	*(*int32)(unsafe.Add(mBase, uint32(v39)+472)) = v8768
	v8771 = int32(_a_F_PostgresMainLoopOnce_195)
	v8772 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v39 + int32(460)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+464)) = int32(1162)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+460)) = v8772
	*(*int32)(unsafe.Add(mBase, uint32(v39)+468)) = v39 + int32(472)
	if v8181 <= int32(0) {
		goto L2089
	} else {
		goto L2090
	}
L2088:
	;
	goto L2087
L2089:
	;
	v8786 = int32(2147483647)
	goto L2091
L2090:
	;
	v8786 = v8181
	goto L2091
L2091:
	;
	v8790 = F_PortalRun(m, v8191, v8786, int32(1), v8532, v8532, v39+int32(512))
	mBase = m.M
	v8791 = m.ExcPending
	if v8791 != 0 {
		goto L1
	} else {
		goto L2092
	}
L2092:
	;
	v8792 = *(*int32)(unsafe.Add(mBase, uint32(v8532)+12))
	m.T0[v8792].(func(*base.Module, int32))(m, v8532)
	mBase = m.M
	v8794 = m.ExcPending
	if v8794 != 0 {
		goto L1
	} else {
		goto L2093
	}
L2093:
	;
	v8795 = int32(_a_F_PostgresMainLoopOnce_195)
	v8797 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130]))
	v8798 = *(*int32)(unsafe.Add(mBase, uint32(v8797)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[130])) = v8798
	if v8790 != 0 {
		goto L2095
	} else {
		goto L2096
	}
L2094:
	;
	v8867 = F_check_log_duration(m, v39+int32(480), v8698)
	mBase = m.M
	v8868 = m.ExcPending
	if v8868 != 0 {
		goto L1
	} else {
		goto L2124
	}
L2095:
	;
	if v8223 == int32(0) {
		goto L2100
	} else {
		goto L2101
	}
L2096:
	;
	goto L2097
L2097:
	;
	v8852 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v8852 == int32(2) {
		goto L2116
	} else {
		goto L2117
	}
L2098:
	;
	F_EndCommand(m, v39+int32(512), v8198)
	mBase = m.M
	v8850 = m.ExcPending
	if v8850 != 0 {
		goto L1
	} else {
		goto L2115
	}
L2099:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v8828 = m.ExcPending
	if v8828 != 0 {
		goto L1
	} else {
		goto L2111
	}
L2100:
	;
	v8803 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133])))
	if v8803&int32(4) == int32(0) {
		goto L2099
	} else {
		goto L2103
	}
L2101:
	;
	goto L2102
L2102:
	;
	v8813 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L2104
L2103:
	;
	goto L2102
L2104:
	;
	if v8813 != 0 {
		goto L2105
	} else {
		goto L2106
	}
L2105:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8816 = m.ExcPending
	if v8816 != 0 {
		goto L1
	} else {
		goto L2108
	}
L2106:
	;
	goto L2107
L2107:
	;
	v8817 = int32(0)
	v8819 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v8819 == v8817 {
		v8846 = v8817
		goto L2098
	} else {
		goto L2109
	}
L2108:
	;
	goto L2107
L2109:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8823 = m.ExcPending
	if v8823 != 0 {
		goto L1
	} else {
		goto L2110
	}
L2110:
	;
	v8825 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = uint8(v8825)
	v8846 = v8817
	goto L2098
L2111:
	;
	v8829 = int32(_a_F_PostgresMainLoopOnce_204)
	v8831 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133])) = v8831 | int32(8)
	v8840 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L2112
L2112:
	;
	if v8840 == int32(0) {
		v8846 = v8232
		goto L2098
	} else {
		goto L2113
	}
L2113:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v8845 = m.ExcPending
	if v8845 != 0 {
		goto L1
	} else {
		goto L2114
	}
L2114:
	;
	v8846 = v8232
	goto L2098
L2115:
	;
	v8864 = v8846
	goto L2094
L2116:
	;
	F_pq_putemptymessage(m, int32(115))
	mBase = m.M
	v8857 = m.ExcPending
	if v8857 != 0 {
		goto L1
	} else {
		goto L2119
	}
L2117:
	;
	goto L2118
L2118:
	;
	v8858 = int32(_a_F_PostgresMainLoopOnce_204)
	v8860 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[133])) = v8860 | int32(8)
	v8864 = v8232
	goto L2094
L2119:
	;
	goto L2118
L2120:
	;
	if v8190&int32(1) != 0 {
		goto L2150
	} else {
		goto L2151
	}
L2121:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), v8944, int32(_a_F_PostgresMainLoopOnce_203))
	mBase = m.M
	v8949 = m.ExcPending
	if v8949 != 0 {
		goto L1
	} else {
		goto L2149
	}
L2122:
	;
	v8890 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8891 = m.ExcPending
	if v8891 != 0 {
		goto L1
	} else {
		goto L2129
	}
L2123:
	;
	v8873 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8874 = m.ExcPending
	if v8874 != 0 {
		goto L1
	} else {
		goto L2125
	}
L2124:
	;
	switch v8867 - int32(1) {
	case 0:
		goto L2123
	case 1:
		goto L2122
	default:
		goto L2120
	}
L2125:
	;
	if v8873 == int32(0) {
		goto L2120
	} else {
		goto L2126
	}
L2126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+240)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_184), v39+int32(240))
	mBase = m.M
	v8884 = m.ExcPending
	if v8884 != 0 {
		goto L1
	} else {
		goto L2127
	}
L2127:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8886 = m.ExcPending
	if v8886 != 0 {
		goto L1
	} else {
		goto L2128
	}
L2128:
	;
	v8944 = int32(2457)
	goto L2121
L2129:
	;
	if v8890 == int32(0) {
		goto L2120
	} else {
		goto L2130
	}
L2130:
	;
	v8894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8176))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+292)) = v8225
	if v8894 != 0 {
		goto L2131
	} else {
		goto L2132
	}
L2131:
	;
	v8897 = v8176
	goto L2133
L2132:
	;
	v8897 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L2133
L2133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+288)) = v8897
	*(*int32)(unsafe.Add(mBase, uint32(v39)+280)) = v8231
	if v8894 != 0 {
		goto L2134
	} else {
		goto L2135
	}
L2134:
	;
	v8902 = int32(_a_F_PostgresMainLoopOnce_196)
	goto L2136
L2135:
	;
	v8902 = int32(_a_F_PostgresMainLoopOnce_183)
	goto L2136
L2136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+284)) = v8902
	if v8540&int32(1) != 0 {
		goto L2137
	} else {
		goto L2138
	}
L2137:
	;
	v8908 = int32(_a_F_PostgresMainLoopOnce_200)
	goto L2139
L2138:
	;
	v8908 = int32(_a_F_PostgresMainLoopOnce_201)
	goto L2139
L2139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+276)) = v8908
	*(*int32)(unsafe.Add(mBase, uint32(v39)+272)) = v39 + int32(480)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_205), v39+int32(272))
	mBase = m.M
	v8917 = m.ExcPending
	if v8917 != 0 {
		goto L1
	} else {
		goto L2140
	}
L2140:
	;
	F_errhidestmt(m)
	mBase = m.M
	v8919 = m.ExcPending
	if v8919 != 0 {
		goto L1
	} else {
		goto L2141
	}
L2141:
	;
	v8920 = int32(2471)
	if v8864 == int32(0) {
		v8944 = v8920
		goto L2121
	} else {
		goto L2142
	}
L2142:
	;
	v8923 = *(*int32)(unsafe.Add(mBase, uint32(v8864)+28))
	if v8923 <= int32(0) {
		v8944 = v8920
		goto L2121
	} else {
		goto L2143
	}
L2143:
	;
	v8927 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[132]))
	if v8927 == int32(0) {
		v8944 = v8920
		goto L2121
	} else {
		goto L2144
	}
L2144:
	;
	v8931 = F_BuildParamLogString(m, v8864, int32(0), v8927)
	mBase = m.M
	v8932 = m.ExcPending
	if v8932 != 0 {
		goto L1
	} else {
		goto L2145
	}
L2145:
	;
	if v8931 == int32(0) {
		v8944 = v8920
		goto L2121
	} else {
		goto L2146
	}
L2146:
	;
	v8935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8931))))
	if v8935 == int32(0) {
		v8944 = v8920
		goto L2121
	} else {
		goto L2147
	}
L2147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+256)) = v8931
	F_errdetail(m, int32(_a_F_PostgresMainLoopOnce_198), v39+int32(256))
	mBase = m.M
	v8943 = m.ExcPending
	if v8943 != 0 {
		goto L1
	} else {
		goto L2148
	}
L2148:
	;
	v8944 = v8920
	goto L2121
L2149:
	;
	goto L2120
L2150:
	;
	F_ShowUsage(m, int32(_a_F_PostgresMainLoopOnce_206))
	mBase = m.M
	v8956 = m.ExcPending
	if v8956 != 0 {
		goto L1
	} else {
		goto L2153
	}
L2151:
	;
	goto L2152
L2152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[46])) = int32(0)
	goto L202
L2153:
	;
	goto L2152
L2154:
	;
	v8965 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v8965 < int32(0) {
		goto L2156
	} else {
		goto L2157
	}
L2155:
	;
	F_pgstat_report_activity(m, int32(5), int32(0))
	mBase = m.M
	F_start_xact_command(m)
	mBase = m.M
	v8975 = m.ExcPending
	if v8975 != 0 {
		goto L1
	} else {
		goto L2159
	}
L2156:
	;
	v8969 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v8969
	goto L2158
L2157:
	;
	goto L2158
L2158:
	;
	goto L2155
L2159:
	;
	v8978 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v8978
	v8981 = v39 + int32(440)
	v8982 = m.G0
	v8984 = v8982 - int32(1568)
	m.G0 = v8984
	v8987 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v8988 = *(*int32)(unsafe.Add(mBase, uint32(v8987)+24))
	goto L2170
L2160:
	;
	v9841 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[31]))
	if v9841 != 0 {
		goto L2349
	} else {
		goto L2350
	}
L2161:
	;
	v9799 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+740))
	v9800 = *(*int32)(unsafe.Add(mBase, uint32(v9799)))
	v9801 = m.T0[v9800].(func(*base.Module, int32) int32)(m, v8984+int32(740))
	mBase = m.M
	v9802 = m.ExcPending
	if v9802 != 0 {
		goto L1
	} else {
		goto L2348
	}
L2162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9745 = m.ExcPending
	if v9745 != 0 {
		goto L1
	} else {
		goto L2344
	}
L2163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9725 = m.ExcPending
	if v9725 != 0 {
		goto L1
	} else {
		goto L2340
	}
L2164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9707 = m.ExcPending
	if v9707 != 0 {
		goto L1
	} else {
		goto L2336
	}
L2165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9688 = m.ExcPending
	if v9688 != 0 {
		goto L1
	} else {
		goto L2332
	}
L2166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9668 = m.ExcPending
	if v9668 != 0 {
		goto L1
	} else {
		goto L2328
	}
L2167:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9649 = m.ExcPending
	if v9649 != 0 {
		goto L1
	} else {
		goto L2325
	}
L2168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9629 = m.ExcPending
	if v9629 != 0 {
		goto L1
	} else {
		goto L2321
	}
L2169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9613 = m.ExcPending
	if v9613 != 0 {
		goto L1
	} else {
		goto L2317
	}
L2170:
	;
	if base.B2i32((v8988-int32(7))&int32(-9) == int32(0)) == int32(0) {
		goto L2171
	} else {
		goto L2172
	}
L2171:
	;
	v8997 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v8998 = m.ExcPending
	if v8998 != 0 {
		goto L1
	} else {
		goto L2174
	}
L2172:
	;
	goto L2173
L2173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9597 = m.ExcPending
	if v9597 != 0 {
		goto L1
	} else {
		goto L2313
	}
L2174:
	;
	F_PushActiveSnapshot(m, v8997)
	mBase = m.M
	v9000 = m.ExcPending
	if v9000 != 0 {
		goto L1
	} else {
		goto L2175
	}
L2175:
	;
	v9002 = F_pq_getmsgint(m, v8981, int32(4))
	mBase = m.M
	v9003 = m.ExcPending
	if v9003 != 0 {
		goto L1
	} else {
		goto L2176
	}
L2176:
	;
	v9005 = v8984 + int32(236)
	v9012 = v8984 + int32(740)
	v9014 = v8984 + int32(240)
	if base.Ui32(v9014) < base.Ui32(v9012) {
		goto L2177
	} else {
		goto L2178
	}
L2177:
	;
	v9016 = v9012
	goto L2179
L2178:
	;
	v9016 = v9014
	goto L2179
L2179:
	;
	v9023 = F__emscripten_memset_bulkmem(m, v9005, base.I32_extend8_s(int32(0)), (v9005^int32(-1)+v9016)&int32(-4)+int32(4))
	mBase = m.M
	goto L2180
L2180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+236)) = int32(0)
	v9027 = F_SearchSysCache1(m, int32(47), v9002)
	mBase = m.M
	v9028 = m.ExcPending
	if v9028 != 0 {
		goto L1
	} else {
		goto L2181
	}
L2181:
	;
	if v9027 == int32(0) {
		goto L2169
	} else {
		goto L2182
	}
L2182:
	;
	v9031 = *(*int32)(unsafe.Add(mBase, uint32(v9027)+16))
	v9032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9031)+22)))
	v9033 = v9031 + v9032
	v9034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9033)+96)))
	if v9034 != int32(102) {
		goto L2168
	} else {
		goto L2183
	}
L2183:
	;
	v9037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9033)+100)))
	if v9037 == int32(1) {
		goto L2168
	} else {
		goto L2184
	}
L2184:
	;
	v9040 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9033)+104)))
	if int32(101) <= v9040 {
		goto L2167
	} else {
		goto L2185
	}
L2185:
	;
	v9043 = *(*int32)(unsafe.Add(mBase, uint32(v9033)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+268)) = v9043
	v9045 = *(*int32)(unsafe.Add(mBase, uint32(v9033)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+272)) = v9045
	v9048 = v8984 + int32(276)
	v9051 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9033)+104)))
	v9053 = v9051 << (uint(int32(2)) % 32)
	if v9053 != 0 {
		goto L2187
	} else {
		goto L2188
	}
L2186:
	;
	v9057 = v8984 + int32(676)
	v9059 = v9033 + int32(4)
	goto L2193
L2187:
	;
	v9054 = F__emscripten_memcpy_bulkmem(m, v9048, v9033+int32(136), v9053)
	mBase = m.M
	v9055 = v9054
	goto L2189
L2188:
	;
	v9055 = v9048
	goto L2189
L2189:
	;
	goto L2186
L2190:
	;
	F_ReleaseCatCache(m, v9027)
	mBase = m.M
	v9176 = m.ExcPending
	if v9176 != 0 {
		goto L1
	} else {
		goto L2222
	}
L2191:
	;
	v9172 = F_strlen(m, v9161)
	mBase = m.M
	goto L2190
L2193:
	;
	goto L2194
L2194:
	;
	v9066 = int32(63)
	if (v9057^v9059)&int32(3) != 0 {
		goto L2198
	} else {
		goto L2199
	}
L2195:
	;
	v9165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9162))) = uint8(v9165)
	goto L2191
L2196:
	;
	v9146 = v9141
	v9147 = v9142
	v9148 = v9143
	goto L2218
L2197:
	;
	if v9136 == int32(0) {
		v9161 = v9134
		v9162 = v9135
		goto L2195
	} else {
		goto L2217
	}
L2198:
	;
	v9134 = v9059
	v9135 = v9057
	v9136 = v9066
	goto L2197
L2199:
	;
	goto L2200
L2200:
	;
	if v9059&int32(3) == int32(0) {
		goto L2202
	} else {
		goto L2203
	}
L2201:
	;
	if v9103 == int32(0) {
		v9161 = v9100
		v9162 = v9101
		goto L2195
	} else {
		goto L2210
	}
L2202:
	;
	v9100 = v9059
	v9101 = v9057
	v9102 = v9066
	v9103 = int32(1)
	goto L2201
L2203:
	;
	goto L2204
L2204:
	;
	v9079 = v9059
	v9080 = v9057
	v9081 = v9066
	goto L2205
L2205:
	;
	v9083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9079))))
	*(*uint8)(unsafe.Add(mBase, uint32(v9080))) = uint8(v9083)
	if v9083 == int32(0) {
		v9141 = v9079
		v9142 = v9080
		v9143 = v9081
		goto L2196
	} else {
		goto L2207
	}
L2206:
	;
	v9100 = v9094
	v9101 = v9088
	v9102 = v9090
	v9103 = v9092
	goto L2201
L2207:
	;
	v9087 = int32(1)
	v9088 = v9080 + v9087
	v9090 = v9081 - v9087
	v9091 = int32(0)
	v9092 = base.B2i32(v9090 != v9091)
	v9094 = v9079 + v9087
	if v9094&int32(3) == v9091 {
		v9100 = v9094
		v9101 = v9088
		v9102 = v9090
		v9103 = v9092
		goto L2201
	} else {
		goto L2208
	}
L2208:
	;
	if v9090 != 0 {
		v9079 = v9094
		v9080 = v9088
		v9081 = v9090
		goto L2205
	} else {
		goto L2209
	}
L2209:
	;
	goto L2206
L2210:
	;
	v9106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9100))))
	if v9106 == int32(0) {
		v9134 = v9100
		v9135 = v9101
		v9136 = v9102
		goto L2197
	} else {
		goto L2211
	}
L2211:
	;
	if base.Ui32(v9102) < base.Ui32(int32(4)) {
		v9134 = v9100
		v9135 = v9101
		v9136 = v9102
		goto L2197
	} else {
		goto L2212
	}
L2212:
	;
	v9112 = v9100
	v9113 = v9101
	v9114 = v9102
	goto L2213
L2213:
	;
	v9117 = *(*int32)(unsafe.Add(mBase, uint32(v9112)))
	v9120 = int32(-2139062144)
	if (int32(16843008)-v9117|v9117)&v9120 != v9120 {
		v9141 = v9112
		v9142 = v9113
		v9143 = v9114
		goto L2196
	} else {
		goto L2215
	}
L2214:
	;
	v9134 = v9128
	v9135 = v9126
	v9136 = v9130
	goto L2197
L2215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9113))) = v9117
	v9125 = int32(4)
	v9126 = v9113 + v9125
	v9128 = v9112 + v9125
	v9130 = v9114 - v9125
	if base.Ui32(int32(3)) < base.Ui32(v9130) {
		v9112 = v9128
		v9113 = v9126
		v9114 = v9130
		goto L2213
	} else {
		goto L2216
	}
L2216:
	;
	goto L2214
L2217:
	;
	v9141 = v9134
	v9142 = v9135
	v9143 = v9136
	goto L2196
L2218:
	;
	v9150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v9147))) = uint8(v9150)
	if v9150 == int32(0) {
		v9161 = v9146
		v9162 = v9147
		goto L2195
	} else {
		goto L2220
	}
L2219:
	;
	v9161 = v9157
	v9162 = v9155
	goto L2195
L2220:
	;
	v9154 = int32(1)
	v9155 = v9147 + v9154
	v9157 = v9146 + v9154
	v9159 = v9148 - v9154
	if v9159 != 0 {
		v9146 = v9157
		v9147 = v9155
		v9148 = v9159
		goto L2218
	} else {
		goto L2221
	}
L2221:
	;
	goto L2219
L2222:
	;
	v9178 = v8984 + int32(240)
	F_fmgr_info(m, v9002, v9178)
	mBase = m.M
	v9180 = m.ExcPending
	if v9180 != 0 {
		goto L1
	} else {
		goto L2223
	}
L2223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+236)) = v9002
	v9183 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[123]))
	if v9183 != int32(3) {
		goto L2224
	} else {
		goto L2225
	}
L2224:
	;
	v9205 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+268))
	v9207 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	v9209 = F_object_aclcheck(m, int32(2615), v9205, v9207, int64(256))
	mBase = m.M
	v9210 = m.ExcPending
	if v9210 != 0 {
		goto L1
	} else {
		goto L2230
	}
L2225:
	;
	v9188 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9189 = m.ExcPending
	if v9189 != 0 {
		goto L1
	} else {
		goto L2226
	}
L2226:
	;
	if v9188 == int32(0) {
		goto L2224
	} else {
		goto L2227
	}
L2227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+180)) = v9002
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+176)) = v9057
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_207), v8984+int32(176))
	mBase = m.M
	v9198 = m.ExcPending
	if v9198 != 0 {
		goto L1
	} else {
		goto L2228
	}
L2228:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(234), int32(_a_F_PostgresMainLoopOnce_209))
	mBase = m.M
	v9203 = m.ExcPending
	if v9203 != 0 {
		goto L1
	} else {
		goto L2229
	}
L2229:
	;
	goto L2224
L2230:
	;
	if v9209 != 0 {
		goto L2231
	} else {
		goto L2232
	}
L2231:
	;
	v9212 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+268))
	v9213 = F_get_namespace_name(m, v9212)
	mBase = m.M
	v9214 = m.ExcPending
	if v9214 != 0 {
		goto L1
	} else {
		goto L2234
	}
L2232:
	;
	goto L2233
L2233:
	;
	v9218 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[135]))
	if v9218 != 0 {
		goto L2236
	} else {
		goto L2237
	}
L2234:
	;
	F_aclcheck_error(m, v9209, int32(36), v9213)
	mBase = m.M
	v9216 = m.ExcPending
	if v9216 != 0 {
		goto L1
	} else {
		goto L2235
	}
L2235:
	;
	goto L2233
L2236:
	;
	v9219 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+268))
	v9221 = F_RunNamespaceSearchHook(m, v9219, int32(1))
	mBase = m.M
	v9222 = m.ExcPending
	if v9222 != 0 {
		goto L1
	} else {
		goto L2239
	}
L2237:
	;
	goto L2238
L2238:
	;
	v9225 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[134]))
	v9227 = F_object_aclcheck(m, int32(1255), v9002, v9225, int64(128))
	mBase = m.M
	v9228 = m.ExcPending
	if v9228 != 0 {
		goto L1
	} else {
		goto L2240
	}
L2239:
	;
	goto L2238
L2240:
	;
	if v9227 != 0 {
		goto L2241
	} else {
		goto L2242
	}
L2241:
	;
	v9230 = F_get_func_name(m, v9002)
	mBase = m.M
	v9231 = m.ExcPending
	if v9231 != 0 {
		goto L1
	} else {
		goto L2244
	}
L2242:
	;
	goto L2243
L2243:
	;
	v9235 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[135]))
	if v9235 != 0 {
		goto L2246
	} else {
		goto L2247
	}
L2244:
	;
	F_aclcheck_error(m, v9227, int32(19), v9230)
	mBase = m.M
	v9233 = m.ExcPending
	if v9233 != 0 {
		goto L1
	} else {
		goto L2245
	}
L2245:
	;
	goto L2243
L2246:
	;
	F_RunFunctionExecuteHook(m, v9002)
	mBase = m.M
	v9237 = m.ExcPending
	if v9237 != 0 {
		goto L1
	} else {
		goto L2249
	}
L2247:
	;
	goto L2248
L2248:
	;
	v9238 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8984)+749)) = v9238
	*(*int64)(unsafe.Add(mBase, uint32(v8984)+744)) = v9238
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+740)) = v9178
	v9244 = F_pq_getmsgint(m, v8981, int32(2))
	mBase = m.M
	v9245 = m.ExcPending
	if v9245 != 0 {
		goto L1
	} else {
		goto L2250
	}
L2249:
	;
	goto L2248
L2250:
	;
	if int32(0) < v9244 {
		goto L2251
	} else {
		goto L2252
	}
L2251:
	;
	v9251 = F_palloc(m, v9244<<(uint(int32(1))%32))
	mBase = m.M
	v9252 = m.ExcPending
	if v9252 != 0 {
		goto L1
	} else {
		goto L2254
	}
L2252:
	;
	v9305 = v1
	goto L2253
L2253:
	;
	v9336 = F_pq_getmsgint(m, v8981, int32(2))
	mBase = m.M
	v9337 = m.ExcPending
	if v9337 != 0 {
		goto L1
	} else {
		goto L2259
	}
L2254:
	;
	v9256 = int32(0)
	goto L2255
L2255:
	;
	v9293 = F_pq_getmsgint(m, v8981, int32(2))
	mBase = m.M
	v9294 = m.ExcPending
	if v9294 != 0 {
		goto L1
	} else {
		goto L2257
	}
L2256:
	;
	v9305 = v9251
	goto L2253
L2257:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v9251+v9256<<(uint(int32(1))%32)))) = uint16(v9293)
	v9297 = v9256 + int32(1)
	if v9297 != v9244 {
		v9256 = v9297
		goto L2255
	} else {
		goto L2258
	}
L2258:
	;
	goto L2256
L2259:
	;
	if int32(100) < v9336 {
		goto L2166
	} else {
		goto L2260
	}
L2260:
	;
	v9340 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8984)+248)))
	if v9336 != v9340 {
		goto L2166
	} else {
		goto L2261
	}
L2261:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8984)+758)) = uint16(v9336)
	if base.B2i32(v9336 != v9244)&base.B2i32(int32(2) <= v9244) != 0 {
		goto L2165
	} else {
		goto L2262
	}
L2262:
	;
	F_initStringInfo(m, v8984+int32(192))
	mBase = m.M
	v9350 = m.ExcPending
	if v9350 != 0 {
		goto L1
	} else {
		goto L2263
	}
L2263:
	;
	if int32(0) < v9336 {
		goto L2264
	} else {
		goto L2265
	}
L2264:
	;
	v9354 = v8984 + int32(760)
	v9363 = int32(0)
	goto L2267
L2265:
	;
	goto L2266
L2266:
	;
	v9533 = F_pq_getmsgint(m, v8981, int32(2))
	mBase = m.M
	v9534 = m.ExcPending
	if v9534 != 0 {
		goto L1
	} else {
		goto L2303
	}
L2267:
	;
	v9397 = v9363 << (uint(int32(3)) % 32)
	v9398 = v8984 + int32(764) + v9397
	v9400 = F_pq_getmsgint(m, v8981, int32(4))
	mBase = m.M
	v9401 = m.ExcPending
	if v9401 != 0 {
		goto L1
	} else {
		goto L2270
	}
L2268:
	;
	goto L2266
L2269:
	;
	if base.B2i32(v9244 < int32(2)) == int32(0) {
		goto L2282
	} else {
		goto L2283
	}
L2270:
	;
	if v9400 == int32(-1) {
		goto L2271
	} else {
		goto L2272
	}
L2271:
	;
	v9404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9398))) = uint8(v9404)
	goto L2269
L2272:
	;
	goto L2273
L2273:
	;
	v9406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9398))) = uint8(v9406)
	if v9400 < v9406 {
		goto L2164
	} else {
		goto L2274
	}
L2274:
	;
	v9411 = v8984 + int32(192)
	v9412 = *(*int32)(unsafe.Add(mBase, uint32(v9411)))
	v9413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9412))) = uint8(v9413)
	*(*int32)(unsafe.Add(mBase, uint32(v9411)+12)) = v9413
	*(*int32)(unsafe.Add(mBase, uint32(v9411)+4)) = v9413
	goto L2275
L2275:
	;
	v9421 = F_pq_getmsgbytes(m, v8981, v9400)
	mBase = m.M
	v9422 = m.ExcPending
	if v9422 != 0 {
		goto L1
	} else {
		goto L2276
	}
L2276:
	;
	F_appendBinaryStringInfo(m, v8984+int32(192), v9421, v9400)
	mBase = m.M
	v9424 = m.ExcPending
	if v9424 != 0 {
		goto L1
	} else {
		goto L2277
	}
L2277:
	;
	goto L2269
L2278:
	;
	v9494 = v9363 + int32(1)
	if v9494 != v9336 {
		v9363 = v9494
		goto L2267
	} else {
		goto L2302
	}
L2279:
	;
	v9468 = *(*int32)(unsafe.Add(mBase, uint32(v9055+v9363<<(uint(int32(2))%32))))
	F_getTypeBinaryInputInfo(m, v9468, v8984+int32(1564), v8984+int32(1560))
	mBase = m.M
	v9474 = m.ExcPending
	if v9474 != 0 {
		goto L1
	} else {
		goto L2295
	}
L2280:
	;
	v9438 = *(*int32)(unsafe.Add(mBase, uint32(v9055+v9363<<(uint(int32(2))%32))))
	F_getTypeInputInfo(m, v9438, v8984+int32(1564), v8984+int32(1560))
	mBase = m.M
	v9444 = m.ExcPending
	if v9444 != 0 {
		goto L1
	} else {
		goto L2286
	}
L2281:
	;
	v9433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9432))))
	switch v9433 {
	case 0:
		goto L2280
	case 1:
		goto L2279
	default:
		goto L2162
	}
L2282:
	;
	v9432 = v9305 + v9363<<(uint(int32(1))%32)
	goto L2281
L2283:
	;
	goto L2284
L2284:
	;
	if v9244 <= int32(0) {
		goto L2280
	} else {
		goto L2285
	}
L2285:
	;
	v9432 = v9305
	goto L2281
L2286:
	;
	if v9400 == int32(-1) {
		goto L2287
	} else {
		goto L2288
	}
L2287:
	;
	v9451 = int32(0)
	goto L2289
L2288:
	;
	v9448 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+192))
	v9449 = F_pg_client_to_server(m, v9448, v9400)
	mBase = m.M
	v9450 = m.ExcPending
	if v9450 != 0 {
		goto L1
	} else {
		goto L2290
	}
L2289:
	;
	v9453 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1564))
	v9454 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1560))
	v9456 = F_OidInputFunctionCall(m, v9453, v9451, v9454, int32(-1))
	mBase = m.M
	v9457 = m.ExcPending
	if v9457 != 0 {
		goto L1
	} else {
		goto L2291
	}
L2290:
	;
	v9451 = v9449
	goto L2289
L2291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9354+v9397))) = v9456
	if v9451 == int32(0) {
		goto L2278
	} else {
		goto L2292
	}
L2292:
	;
	v9461 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+192))
	if v9451 == v9461 {
		goto L2278
	} else {
		goto L2293
	}
L2293:
	;
	F_pfree(m, v9451)
	mBase = m.M
	v9464 = m.ExcPending
	if v9464 != 0 {
		goto L1
	} else {
		goto L2294
	}
L2294:
	;
	goto L2278
L2295:
	;
	v9476 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1564))
	v9481 = base.B2i32(v9400 == int32(-1))
	if v9400 == int32(-1) {
		goto L2296
	} else {
		goto L2297
	}
L2296:
	;
	v9482 = int32(0)
	goto L2298
L2297:
	;
	v9482 = v8984 + int32(192)
	goto L2298
L2298:
	;
	v9483 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1560))
	v9485 = F_OidReceiveFunctionCall(m, v9476, v9482, v9483, int32(-1))
	mBase = m.M
	v9486 = m.ExcPending
	if v9486 != 0 {
		goto L1
	} else {
		goto L2299
	}
L2299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9354+v9397))) = v9485
	if v9400 == int32(-1) {
		goto L2278
	} else {
		goto L2300
	}
L2300:
	;
	v9488 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+204))
	v9489 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+196))
	if v9488 != v9489 {
		goto L2163
	} else {
		goto L2301
	}
L2301:
	;
	goto L2278
L2302:
	;
	goto L2268
L2303:
	;
	F_pq_getmsgend(m, v8981)
	mBase = m.M
	v9536 = m.ExcPending
	if v9536 != 0 {
		goto L1
	} else {
		goto L2304
	}
L2304:
	;
	v9537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8984)+250)))
	if v9537 != int32(1) {
		goto L2161
	} else {
		goto L2305
	}
L2305:
	;
	if v9336 <= int32(0) {
		goto L2161
	} else {
		goto L2306
	}
L2306:
	;
	v9549 = int32(0)
	goto L2307
L2307:
	;
	v9585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8984+int32(764)+v9549<<(uint(int32(3))%32)))))
	if v9585 == int32(0) {
		goto L2309
	} else {
		goto L2310
	}
L2308:
	;
	v9591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8984)+756)) = uint8(v9591)
	v9839 = int32(0)
	goto L2160
L2309:
	;
	v9589 = v9549 + int32(1)
	if base.I32_extend16_s(v9336) != v9589 {
		v9549 = v9589
		goto L2307
	} else {
		goto L2312
	}
L2310:
	;
	goto L2311
L2311:
	;
	goto L2308
L2312:
	;
	goto L2161
L2313:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v9600 = m.ExcPending
	if v9600 != 0 {
		goto L1
	} else {
		goto L2314
	}
L2314:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v9604 = m.ExcPending
	if v9604 != 0 {
		goto L1
	} else {
		goto L2315
	}
L2315:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(209), int32(_a_F_PostgresMainLoopOnce_209))
	mBase = m.M
	v9609 = m.ExcPending
	if v9609 != 0 {
		goto L1
	} else {
		goto L2316
	}
L2316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2317:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v9616 = m.ExcPending
	if v9616 != 0 {
		goto L1
	} else {
		goto L2318
	}
L2318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984))) = v9002
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_210), v8984)
	mBase = m.M
	v9620 = m.ExcPending
	if v9620 != 0 {
		goto L1
	} else {
		goto L2319
	}
L2319:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(141), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9625 = m.ExcPending
	if v9625 != 0 {
		goto L1
	} else {
		goto L2320
	}
L2320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2321:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9632 = m.ExcPending
	if v9632 != 0 {
		goto L1
	} else {
		goto L2322
	}
L2322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+16)) = v9033 + int32(4)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_212), v8984+int32(16))
	mBase = m.M
	v9640 = m.ExcPending
	if v9640 != 0 {
		goto L1
	} else {
		goto L2323
	}
L2323:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(149), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9645 = m.ExcPending
	if v9645 != 0 {
		goto L1
	} else {
		goto L2324
	}
L2324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+32)) = v9033 + int32(4)
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_213), v8984+int32(32))
	mBase = m.M
	v9659 = m.ExcPending
	if v9659 != 0 {
		goto L1
	} else {
		goto L2326
	}
L2326:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(154), int32(_a_F_PostgresMainLoopOnce_211))
	mBase = m.M
	v9664 = m.ExcPending
	if v9664 != 0 {
		goto L1
	} else {
		goto L2327
	}
L2327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2328:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9671 = m.ExcPending
	if v9671 != 0 {
		goto L1
	} else {
		goto L2329
	}
L2329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+48)) = v9336
	v9673 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8984)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+52)) = v9673
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_214), v8984+int32(48))
	mBase = m.M
	v9679 = m.ExcPending
	if v9679 != 0 {
		goto L1
	} else {
		goto L2330
	}
L2330:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(353), int32(_a_F_PostgresMainLoopOnce_215))
	mBase = m.M
	v9684 = m.ExcPending
	if v9684 != 0 {
		goto L1
	} else {
		goto L2331
	}
L2331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2332:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9691 = m.ExcPending
	if v9691 != 0 {
		goto L1
	} else {
		goto L2333
	}
L2333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+164)) = v9336
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+160)) = v9244
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_216), v8984+int32(160))
	mBase = m.M
	v9698 = m.ExcPending
	if v9698 != 0 {
		goto L1
	} else {
		goto L2334
	}
L2334:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(361), int32(_a_F_PostgresMainLoopOnce_215))
	mBase = m.M
	v9703 = m.ExcPending
	if v9703 != 0 {
		goto L1
	} else {
		goto L2335
	}
L2335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2336:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v9710 = m.ExcPending
	if v9710 != 0 {
		goto L1
	} else {
		goto L2337
	}
L2337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+144)) = v9400
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_217), v8984+int32(144))
	mBase = m.M
	v9716 = m.ExcPending
	if v9716 != 0 {
		goto L1
	} else {
		goto L2338
	}
L2338:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(385), int32(_a_F_PostgresMainLoopOnce_215))
	mBase = m.M
	v9721 = m.ExcPending
	if v9721 != 0 {
		goto L1
	} else {
		goto L2339
	}
L2339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2340:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v9728 = m.ExcPending
	if v9728 != 0 {
		goto L1
	} else {
		goto L2341
	}
L2341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+128)) = v9363 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_218), v8984+int32(128))
	mBase = m.M
	v9736 = m.ExcPending
	if v9736 != 0 {
		goto L1
	} else {
		goto L2342
	}
L2342:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(448), int32(_a_F_PostgresMainLoopOnce_215))
	mBase = m.M
	v9741 = m.ExcPending
	if v9741 != 0 {
		goto L1
	} else {
		goto L2343
	}
L2343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2344:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9748 = m.ExcPending
	if v9748 != 0 {
		goto L1
	} else {
		goto L2345
	}
L2345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+64)) = base.I32_extend16_s(v9433)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_219), v8984-int32(-64))
	mBase = m.M
	v9755 = m.ExcPending
	if v9755 != 0 {
		goto L1
	} else {
		goto L2346
	}
L2346:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(453), int32(_a_F_PostgresMainLoopOnce_215))
	mBase = m.M
	v9760 = m.ExcPending
	if v9760 != 0 {
		goto L1
	} else {
		goto L2347
	}
L2347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2348:
	;
	v9839 = v9801
	goto L2160
L2349:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9843 = m.ExcPending
	if v9843 != 0 {
		goto L1
	} else {
		goto L2352
	}
L2350:
	;
	goto L2351
L2351:
	;
	v9844 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+272))
	v9845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8984)+756)))
	F_pq_beginmessage(m, v8984+int32(192), int32(86))
	mBase = m.M
	v9850 = m.ExcPending
	if v9850 != 0 {
		goto L1
	} else {
		goto L2353
	}
L2352:
	;
	goto L2351
L2353:
	;
	if v9845 == int32(1) {
		goto L2355
	} else {
		goto L2356
	}
L2354:
	;
	F_pq_endmessage(m, v8984+int32(192))
	mBase = m.M
	v10016 = m.ExcPending
	if v10016 != 0 {
		goto L1
	} else {
		goto L2387
	}
L2355:
	;
	F_enlargeStringInfo(m, v8984+int32(192), int32(4))
	mBase = m.M
	v9857 = m.ExcPending
	if v9857 != 0 {
		goto L1
	} else {
		goto L2358
	}
L2356:
	;
	goto L2357
L2357:
	;
	switch v9533 & int32(_a_F_PostgresMainLoopOnce_220) {
	case 0:
		goto L2359
	case 1:
		goto L2361
	default:
		goto L2360
	}
L2358:
	;
	v9858 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+196))
	v9859 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v9858+v9859))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+196)) = v9858 + int32(4)
	goto L2354
L2359:
	;
	F_getTypeOutputInfo(m, v9844, v8984+int32(1564), v8984+int32(1560))
	mBase = m.M
	v9998 = m.ExcPending
	if v9998 != 0 {
		goto L1
	} else {
		goto L2383
	}
L2360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9977 = m.ExcPending
	if v9977 != 0 {
		goto L1
	} else {
		goto L2379
	}
L2361:
	;
	F_getTypeBinaryOutputInfo(m, v9844, v8984+int32(1564), v8984+int32(1560))
	mBase = m.M
	v9873 = m.ExcPending
	if v9873 != 0 {
		goto L1
	} else {
		goto L2362
	}
L2362:
	;
	v9874 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1564))
	v9875 = m.G0
	v9877 = v9875 + int32(-64)
	m.G0 = v9877
	v9882 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1]))
	F_fmgr_info_cxt_security(m, v9874, v9875+int32(-56), v9882, int32(0))
	mBase = m.M
	v9885 = m.ExcPending
	if v9885 != 0 {
		goto L1
	} else {
		goto L2364
	}
L2363:
	;
	v9929 = *(*int32)(unsafe.Add(mBase, uint32(v9911)))
	F_enlargeStringInfo(m, v8984+int32(192), int32(4))
	mBase = m.M
	v9934 = m.ExcPending
	if v9934 != 0 {
		goto L1
	} else {
		goto L2376
	}
L2364:
	;
	v9886 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9877)+45)) = v9886
	*(*int64)(unsafe.Add(mBase, uint32(v9877)+40)) = v9886
	v9890 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9877)+60)) = uint8(v9890)
	*(*int32)(unsafe.Add(mBase, uint32(v9877)+56)) = v9839
	*(*int32)(unsafe.Add(mBase, uint32(v9877)+36)) = v9875 + int32(-56)
	v9896 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v9877)+54)) = uint16(v9896)
	v9900 = *(*int32)(unsafe.Add(mBase, uint32(v9877)+8))
	v9901 = m.T0[v9900].(func(*base.Module, int32) int32)(m, v9875+int32(-28))
	mBase = m.M
	v9902 = m.ExcPending
	if v9902 != 0 {
		goto L1
	} else {
		goto L2365
	}
L2365:
	;
	v9903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9877)+52)))
	if v9903 != int32(1) {
		goto L2366
	} else {
		goto L2367
	}
L2366:
	;
	v9906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9901))))
	if v9906&int32(3) != 0 {
		goto L2369
	} else {
		goto L2370
	}
L2367:
	;
	goto L2368
L2368:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9918 = m.ExcPending
	if v9918 != 0 {
		goto L1
	} else {
		goto L2373
	}
L2369:
	;
	v9909 = F_detoast_attr(m, v9901)
	mBase = m.M
	v9910 = m.ExcPending
	if v9910 != 0 {
		goto L1
	} else {
		goto L2372
	}
L2370:
	;
	v9911 = v9901
	goto L2371
L2371:
	;
	m.G0 = v9877 - int32(-64)
	goto L2363
L2372:
	;
	v9911 = v9909
	goto L2371
L2373:
	;
	v9919 = *(*int32)(unsafe.Add(mBase, uint32(v9877)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9877))) = v9919
	F_errmsg_internal(m, int32(_a_F_PostgresMainLoopOnce_221), v9877)
	mBase = m.M
	v9923 = m.ExcPending
	if v9923 != 0 {
		goto L1
	} else {
		goto L2374
	}
L2374:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_222), int32(1143), int32(_a_F_PostgresMainLoopOnce_223))
	mBase = m.M
	v9928 = m.ExcPending
	if v9928 != 0 {
		goto L1
	} else {
		goto L2375
	}
L2375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2376:
	;
	v9935 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+196))
	v9936 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+192))
	v9938 = int32(2)
	v9940 = int32(4)
	v9941 = int32(base.Ui32(v9929)>>(uint(v9938)%32)) - v9940
	v9942 = int32(24)
	v9944 = int32(_a_F_PostgresMainLoopOnce_122)
	v9946 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9935+v9936))) = v9941<<(uint(v9942)%32) | v9941&v9944<<(uint(v9946)%32) | (int32(base.Ui32(v9941)>>(uint(v9946)%32))&v9944 | int32(base.Ui32(v9941)>>(uint(v9942)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+196)) = v9935 + v9940
	v9965 = *(*int32)(unsafe.Add(mBase, uint32(v9911)))
	F_pq_sendbytes(m, v8984+int32(192), v9911+v9940, int32(base.Ui32(v9965)>>(uint(v9938)%32))-v9940)
	mBase = m.M
	v9971 = m.ExcPending
	if v9971 != 0 {
		goto L1
	} else {
		goto L2377
	}
L2377:
	;
	F_pfree(m, v9911)
	mBase = m.M
	v9973 = m.ExcPending
	if v9973 != 0 {
		goto L1
	} else {
		goto L2378
	}
L2378:
	;
	goto L2354
L2379:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v9980 = m.ExcPending
	if v9980 != 0 {
		goto L1
	} else {
		goto L2380
	}
L2380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+112)) = base.I32_extend16_s(v9533)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_219), v8984+int32(112))
	mBase = m.M
	v9987 = m.ExcPending
	if v9987 != 0 {
		goto L1
	} else {
		goto L2381
	}
L2381:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), int32(106), int32(_a_F_PostgresMainLoopOnce_224))
	mBase = m.M
	v9992 = m.ExcPending
	if v9992 != 0 {
		goto L1
	} else {
		goto L2382
	}
L2382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2383:
	;
	v10001 = *(*int32)(unsafe.Add(mBase, uint32(v8984)+1564))
	v10002 = F_OidOutputFunctionCall(m, v10001, v9839)
	mBase = m.M
	v10003 = m.ExcPending
	if v10003 != 0 {
		goto L1
	} else {
		goto L2384
	}
L2384:
	;
	v10004 = F_strlen(m, v10002)
	mBase = m.M
	F_pq_sendcountedtext(m, v8984+int32(192), v10002, v10004)
	mBase = m.M
	v10006 = m.ExcPending
	if v10006 != 0 {
		goto L1
	} else {
		goto L2385
	}
L2385:
	;
	F_pfree(m, v10002)
	mBase = m.M
	v10008 = m.ExcPending
	if v10008 != 0 {
		goto L1
	} else {
		goto L2386
	}
L2386:
	;
	goto L2354
L2387:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v10018 = m.ExcPending
	if v10018 != 0 {
		goto L1
	} else {
		goto L2388
	}
L2388:
	;
	v10024 = F_check_log_duration(m, v8984+int32(192), base.B2i32(v9183 == int32(3)))
	mBase = m.M
	v10025 = m.ExcPending
	if v10025 != 0 {
		goto L1
	} else {
		goto L2393
	}
L2389:
	;
	m.G0 = v8984 + int32(1568)
	v10072 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L2401
L2390:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_208), v10060, int32(_a_F_PostgresMainLoopOnce_209))
	mBase = m.M
	v10063 = m.ExcPending
	if v10063 != 0 {
		goto L1
	} else {
		goto L2400
	}
L2391:
	;
	v10045 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10046 = m.ExcPending
	if v10046 != 0 {
		goto L1
	} else {
		goto L2397
	}
L2392:
	;
	v10030 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10031 = m.ExcPending
	if v10031 != 0 {
		goto L1
	} else {
		goto L2394
	}
L2393:
	;
	switch v10024 - int32(1) {
	case 0:
		goto L2392
	case 1:
		goto L2391
	default:
		goto L2389
	}
L2394:
	;
	if v10030 == int32(0) {
		goto L2389
	} else {
		goto L2395
	}
L2395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+80)) = v8984 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_184), v8984+int32(80))
	mBase = m.M
	v10041 = m.ExcPending
	if v10041 != 0 {
		goto L1
	} else {
		goto L2396
	}
L2396:
	;
	v10060 = int32(312)
	goto L2390
L2397:
	;
	if v10045 == int32(0) {
		goto L2389
	} else {
		goto L2398
	}
L2398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+104)) = v9002
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+100)) = v9057
	*(*int32)(unsafe.Add(mBase, uint32(v8984)+96)) = v8984 + int32(192)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_225), v8984+int32(96))
	mBase = m.M
	v10058 = m.ExcPending
	if v10058 != 0 {
		goto L1
	} else {
		goto L2399
	}
L2399:
	;
	v10060 = int32(317)
	goto L2390
L2400:
	;
	goto L2389
L2401:
	;
	if v10072 != 0 {
		goto L2402
	} else {
		goto L2403
	}
L2402:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v10075 = m.ExcPending
	if v10075 != 0 {
		goto L1
	} else {
		goto L2405
	}
L2403:
	;
	goto L2404
L2404:
	;
	v10077 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v10077 != 0 {
		goto L2406
	} else {
		goto L2407
	}
L2405:
	;
	goto L2404
L2406:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v10079 = m.ExcPending
	if v10079 != 0 {
		goto L1
	} else {
		goto L2409
	}
L2407:
	;
	goto L2408
L2408:
	;
	v10084 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v10084)
	goto L202
L2409:
	;
	v10081 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = uint8(v10081)
	goto L2408
L2410:
	;
	v10092 = F_pq_getmsgbyte(m, v39+int32(440))
	mBase = m.M
	v10093 = m.ExcPending
	if v10093 != 0 {
		goto L1
	} else {
		goto L2411
	}
L2411:
	;
	v10096 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v10097 = m.ExcPending
	if v10097 != 0 {
		goto L1
	} else {
		goto L2412
	}
L2412:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v10101 = m.ExcPending
	if v10101 != 0 {
		goto L1
	} else {
		goto L2413
	}
L2413:
	;
	switch v10092 - int32(80) {
	case 0:
		goto L2415
	default:
		goto L186
	case 3:
		goto L2416
	}
L2414:
	;
	v10126 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10126 != int32(2) {
		goto L202
	} else {
		goto L2426
	}
L2415:
	;
	v10117 = F_GetPortalByName(m, v10096)
	mBase = m.M
	v10118 = m.ExcPending
	if v10118 != 0 {
		goto L1
	} else {
		goto L2423
	}
L2416:
	;
	v10104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10096))))
	if v10104 != 0 {
		goto L2417
	} else {
		goto L2418
	}
L2417:
	;
	F_DropPreparedStatement(m, v10096, int32(0))
	mBase = m.M
	v10107 = m.ExcPending
	if v10107 != 0 {
		goto L1
	} else {
		goto L2420
	}
L2418:
	;
	goto L2419
L2419:
	;
	v10109 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v10109 == int32(0) {
		goto L2414
	} else {
		goto L2421
	}
L2420:
	;
	goto L2414
L2421:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121])) = int32(0)
	F_DropCachedPlan(m, v10109)
	mBase = m.M
	v10116 = m.ExcPending
	if v10116 != 0 {
		goto L1
	} else {
		goto L2422
	}
L2422:
	;
	goto L2414
L2423:
	;
	if v10117 == int32(0) {
		goto L2414
	} else {
		goto L2424
	}
L2424:
	;
	F_PortalDrop(m, v10117, int32(0))
	mBase = m.M
	v10123 = m.ExcPending
	if v10123 != 0 {
		goto L1
	} else {
		goto L2425
	}
L2425:
	;
	goto L2414
L2426:
	;
	F_pq_putemptymessage(m, int32(51))
	mBase = m.M
	v10131 = m.ExcPending
	if v10131 != 0 {
		goto L1
	} else {
		goto L2427
	}
L2427:
	;
	goto L202
L2428:
	;
	v10137 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[42]))
	if v10137 < int32(0) {
		goto L2430
	} else {
		goto L2431
	}
L2429:
	;
	v10145 = F_pq_getmsgbyte(m, v39+int32(440))
	mBase = m.M
	v10146 = m.ExcPending
	if v10146 != 0 {
		goto L1
	} else {
		goto L2433
	}
L2430:
	;
	v10141 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[43])) = v10141
	goto L2432
L2431:
	;
	goto L2432
L2432:
	;
	goto L2429
L2433:
	;
	v10149 = F_pq_getmsgstring(m, v39+int32(440))
	mBase = m.M
	v10150 = m.ExcPending
	if v10150 != 0 {
		goto L1
	} else {
		goto L2434
	}
L2434:
	;
	F_pq_getmsgend(m, v39+int32(440))
	mBase = m.M
	v10154 = m.ExcPending
	if v10154 != 0 {
		goto L1
	} else {
		goto L2435
	}
L2435:
	;
	switch v10145 - int32(80) {
	case 0:
		goto L2437
	default:
		goto L2436
	case 3:
		goto L2438
	}
L2436:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10377 = m.ExcPending
	if v10377 != 0 {
		goto L1
	} else {
		goto L2480
	}
L2437:
	;
	F_start_xact_command(m)
	mBase = m.M
	v10341 = m.ExcPending
	if v10341 != 0 {
		goto L1
	} else {
		goto L2465
	}
L2438:
	;
	F_start_xact_command(m)
	mBase = m.M
	v10158 = m.ExcPending
	if v10158 != 0 {
		goto L1
	} else {
		goto L2439
	}
L2439:
	;
	v10161 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v10161
	v10163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10149))))
	if v10163 != 0 {
		goto L2441
	} else {
		goto L2442
	}
L2440:
	;
	v10174 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v10175 = *(*int32)(unsafe.Add(mBase, uint32(v10174)+24))
	goto L2446
L2441:
	;
	v10165 = F_FetchPreparedStatement(m, v10149, int32(1))
	mBase = m.M
	v10166 = m.ExcPending
	if v10166 != 0 {
		goto L1
	} else {
		goto L2444
	}
L2442:
	;
	goto L2443
L2443:
	;
	v10169 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[121]))
	if v10169 == int32(0) {
		goto L184
	} else {
		goto L2445
	}
L2444:
	;
	v10167 = *(*int32)(unsafe.Add(mBase, uint32(v10165)+64))
	v10172 = v10167
	goto L2440
L2445:
	;
	v10172 = v10169
	goto L2440
L2446:
	;
	if (v10175-int32(7))&int32(-9) == int32(0) {
		goto L2447
	} else {
		goto L2448
	}
L2447:
	;
	v10182 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+52))
	if v10182 != 0 {
		goto L183
	} else {
		goto L2450
	}
L2448:
	;
	goto L2449
L2449:
	;
	v10184 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10184 != int32(2) {
		goto L202
	} else {
		goto L2451
	}
L2450:
	;
	goto L2449
L2451:
	;
	v10187 = int32(_a_F_PostgresMainLoopOnce_226)
	F_resetStringInfo(m, v10187)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[136])) = int32(116)
	goto L2452
L2452:
	;
	v10191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10172)+24)))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_226), int32(2))
	mBase = m.M
	v10195 = m.ExcPending
	if v10195 != 0 {
		goto L1
	} else {
		goto L2453
	}
L2453:
	;
	v10197 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137]))
	v10198 = int32(_a_F_PostgresMainLoopOnce_227)
	v10199 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138]))
	v10201 = int32(8)
	v10205 = v10191<<(uint(v10201)%32) | int32(base.Ui32(v10191)>>(uint(v10201)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v10197+v10199))) = uint16(v10205)
	v10209 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138])) = v10209 + int32(2)
	v10213 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+24))
	if int32(0) < v10213 {
		goto L2454
	} else {
		goto L2455
	}
L2454:
	;
	v10220 = int32(0)
	goto L2457
L2455:
	;
	goto L2456
L2456:
	;
	F_pq_endmessage_reuse(m, int32(_a_F_PostgresMainLoopOnce_226))
	mBase = m.M
	v10329 = m.ExcPending
	if v10329 != 0 {
		goto L1
	} else {
		goto L2461
	}
L2457:
	;
	v10253 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+20))
	v10257 = *(*int32)(unsafe.Add(mBase, uint32(v10253+v10220<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, int32(_a_F_PostgresMainLoopOnce_226), int32(4))
	mBase = m.M
	v10261 = m.ExcPending
	if v10261 != 0 {
		goto L1
	} else {
		goto L2459
	}
L2458:
	;
	goto L2456
L2459:
	;
	v10262 = int32(_a_F_PostgresMainLoopOnce_227)
	v10263 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138]))
	v10265 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[137]))
	v10267 = int32(24)
	v10269 = int32(_a_F_PostgresMainLoopOnce_122)
	v10271 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10263+v10265))) = v10257<<(uint(v10267)%32) | v10257&v10269<<(uint(v10271)%32) | (int32(base.Ui32(v10257)>>(uint(v10271)%32))&v10269 | int32(base.Ui32(v10257)>>(uint(v10267)%32)))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[138])) = v10263 + int32(4)
	v10288 = v10220 + int32(1)
	v10289 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+24))
	if v10288 < v10289 {
		v10220 = v10288
		goto L2457
	} else {
		goto L2460
	}
L2460:
	;
	goto L2458
L2461:
	;
	v10330 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+52))
	if v10330 == int32(0) {
		goto L207
	} else {
		goto L2462
	}
L2462:
	;
	v10333 = F_CachedPlanGetTargetList(m, v10172)
	mBase = m.M
	v10334 = m.ExcPending
	if v10334 != 0 {
		goto L1
	} else {
		goto L2463
	}
L2463:
	;
	v10336 = *(*int32)(unsafe.Add(mBase, uint32(v10172)+52))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_226), v10336, v10333, int32(0))
	mBase = m.M
	v10339 = m.ExcPending
	if v10339 != 0 {
		goto L1
	} else {
		goto L2464
	}
L2464:
	;
	goto L202
L2465:
	;
	v10344 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[1])) = v10344
	v10346 = F_GetPortalByName(m, v10149)
	mBase = m.M
	v10347 = m.ExcPending
	if v10347 != 0 {
		goto L1
	} else {
		goto L2466
	}
L2466:
	;
	if v10346 == int32(0) {
		goto L182
	} else {
		goto L2467
	}
L2467:
	;
	v10351 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v10352 = *(*int32)(unsafe.Add(mBase, uint32(v10351)+24))
	goto L2468
L2468:
	;
	if (v10352-int32(7))&int32(-9) == int32(0) {
		goto L2469
	} else {
		goto L2470
	}
L2469:
	;
	v10359 = *(*int32)(unsafe.Add(mBase, uint32(v10346)+92))
	if v10359 != 0 {
		goto L181
	} else {
		goto L2472
	}
L2470:
	;
	goto L2471
L2471:
	;
	v10361 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10361 != int32(2) {
		goto L202
	} else {
		goto L2473
	}
L2472:
	;
	goto L2471
L2473:
	;
	v10364 = *(*int32)(unsafe.Add(mBase, uint32(v10346)+92))
	if v10364 != 0 {
		goto L2474
	} else {
		goto L2475
	}
L2474:
	;
	v10366 = F_FetchPortalTargetList(m, v10346)
	mBase = m.M
	v10367 = m.ExcPending
	if v10367 != 0 {
		goto L1
	} else {
		goto L2477
	}
L2475:
	;
	goto L2476
L2476:
	;
	F_pq_putemptymessage(m, int32(110))
	mBase = m.M
	v10373 = m.ExcPending
	if v10373 != 0 {
		goto L1
	} else {
		goto L2479
	}
L2477:
	;
	v10368 = *(*int32)(unsafe.Add(mBase, uint32(v10346)+96))
	F_SendRowDescriptionMessage(m, int32(_a_F_PostgresMainLoopOnce_226), v10364, v10366, v10368)
	mBase = m.M
	v10370 = m.ExcPending
	if v10370 != 0 {
		goto L1
	} else {
		goto L2478
	}
L2478:
	;
	goto L202
L2479:
	;
	goto L202
L2480:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10380 = m.ExcPending
	if v10380 != 0 {
		goto L1
	} else {
		goto L2481
	}
L2481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+368)) = v10145
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_228), v39+int32(368))
	mBase = m.M
	v10386 = m.ExcPending
	if v10386 != 0 {
		goto L1
	} else {
		goto L2482
	}
L2482:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_229), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10391 = m.ExcPending
	if v10391 != 0 {
		goto L1
	} else {
		goto L2483
	}
L2483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2484:
	;
	v10397 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11]))
	if v10397 != int32(2) {
		goto L202
	} else {
		goto L2485
	}
L2485:
	;
	v10401 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[95]))
	v10402 = *(*int32)(unsafe.Add(mBase, uint32(v10401)+4))
	v10403 = m.T0[v10402].(func(*base.Module) int32)(m)
	mBase = m.M
	v10404 = m.ExcPending
	if v10404 != 0 {
		goto L1
	} else {
		goto L2486
	}
L2486:
	;
	goto L202
L2487:
	;
	v10411 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[9]))
	v10412 = *(*int32)(unsafe.Add(mBase, uint32(v10411)+24))
	if v10412 == int32(4) {
		goto L2489
	} else {
		goto L2490
	}
L2488:
	;
	v10422 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[124])))
	goto L2492
L2489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10411)+24)) = int32(1)
	goto L2491
L2490:
	;
	goto L2491
L2491:
	;
	goto L2488
L2492:
	;
	if v10422 != 0 {
		goto L2493
	} else {
		goto L2494
	}
L2493:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v10425 = m.ExcPending
	if v10425 != 0 {
		goto L1
	} else {
		goto L2496
	}
L2494:
	;
	goto L2495
L2495:
	;
	v10427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])))
	if v10427 != 0 {
		goto L2497
	} else {
		goto L2498
	}
L2496:
	;
	goto L2495
L2497:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v10429 = m.ExcPending
	if v10429 != 0 {
		goto L1
	} else {
		goto L2500
	}
L2498:
	;
	goto L2499
L2499:
	;
	v10434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[4])) = uint8(v10434)
	goto L202
L2500:
	;
	v10431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[129])) = uint8(v10431)
	goto L2499
L2501:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[11])) = int32(0)
	goto L2503
L2502:
	;
	goto L2503
L2503:
	;
	v10447 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresMainLoopOnce[139]))
	if v10447 != 0 {
		goto L180
	} else {
		goto L2504
	}
L2504:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v10450 = m.ExcPending
	if v10450 != 0 {
		goto L1
	} else {
		goto L2505
	}
L2505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2506:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10457 = m.ExcPending
	if v10457 != 0 {
		goto L1
	} else {
		goto L2507
	}
L2507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v587
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_11), v39+int32(16))
	mBase = m.M
	v10463 = m.ExcPending
	if v10463 != 0 {
		goto L1
	} else {
		goto L2508
	}
L2508:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_230), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10468 = m.ExcPending
	if v10468 != 0 {
		goto L1
	} else {
		goto L2509
	}
L2509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2510:
	;
	goto L202
L2511:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10517 = m.ExcPending
	if v10517 != 0 {
		goto L1
	} else {
		goto L2512
	}
L2512:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), int32(0))
	mBase = m.M
	v10521 = m.ExcPending
	if v10521 != 0 {
		goto L1
	} else {
		goto L2513
	}
L2513:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_232), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10526 = m.ExcPending
	if v10526 != 0 {
		goto L1
	} else {
		goto L2514
	}
L2514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2515:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v10533 = m.ExcPending
	if v10533 != 0 {
		goto L1
	} else {
		goto L2516
	}
L2516:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_234), int32(0))
	mBase = m.M
	v10537 = m.ExcPending
	if v10537 != 0 {
		goto L1
	} else {
		goto L2517
	}
L2517:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1581), int32(_a_F_PostgresMainLoopOnce_189))
	mBase = m.M
	v10542 = m.ExcPending
	if v10542 != 0 {
		goto L1
	} else {
		goto L2518
	}
L2518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2519:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10549 = m.ExcPending
	if v10549 != 0 {
		goto L1
	} else {
		goto L2520
	}
L2520:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v10553 = m.ExcPending
	if v10553 != 0 {
		goto L1
	} else {
		goto L2521
	}
L2521:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10555 = m.ExcPending
	if v10555 != 0 {
		goto L1
	} else {
		goto L2522
	}
L2522:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1603), int32(_a_F_PostgresMainLoopOnce_189))
	mBase = m.M
	v10560 = m.ExcPending
	if v10560 != 0 {
		goto L1
	} else {
		goto L2523
	}
L2523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2524:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10567 = m.ExcPending
	if v10567 != 0 {
		goto L1
	} else {
		goto L2525
	}
L2525:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), int32(0))
	mBase = m.M
	v10571 = m.ExcPending
	if v10571 != 0 {
		goto L1
	} else {
		goto L2526
	}
L2526:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_232), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10576 = m.ExcPending
	if v10576 != 0 {
		goto L1
	} else {
		goto L2527
	}
L2527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2528:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v10583 = m.ExcPending
	if v10583 != 0 {
		goto L1
	} else {
		goto L2529
	}
L2529:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_235), int32(0))
	mBase = m.M
	v10587 = m.ExcPending
	if v10587 != 0 {
		goto L1
	} else {
		goto L2530
	}
L2530:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1778), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10592 = m.ExcPending
	if v10592 != 0 {
		goto L1
	} else {
		goto L2531
	}
L2531:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2532:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10599 = m.ExcPending
	if v10599 != 0 {
		goto L1
	} else {
		goto L2533
	}
L2533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+196)) = v7423
	*(*int32)(unsafe.Add(mBase, uint32(v39)+192)) = v7327
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_236), v39+int32(192))
	mBase = m.M
	v10606 = m.ExcPending
	if v10606 != 0 {
		goto L1
	} else {
		goto L2534
	}
L2534:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1831), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10611 = m.ExcPending
	if v10611 != 0 {
		goto L1
	} else {
		goto L2535
	}
L2535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2536:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10618 = m.ExcPending
	if v10618 != 0 {
		goto L1
	} else {
		goto L2537
	}
L2537:
	;
	v10619 = *(*int32)(unsafe.Add(mBase, uint32(v7134)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v10619
	*(*int32)(unsafe.Add(mBase, uint32(v39)+180)) = v7100
	*(*int32)(unsafe.Add(mBase, uint32(v39)+176)) = v7423
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_237), v39+int32(176))
	mBase = m.M
	v10627 = m.ExcPending
	if v10627 != 0 {
		goto L1
	} else {
		goto L2538
	}
L2538:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1837), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10632 = m.ExcPending
	if v10632 != 0 {
		goto L1
	} else {
		goto L2539
	}
L2539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2540:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10640 = m.ExcPending
	if v10640 != 0 {
		goto L1
	} else {
		goto L2541
	}
L2541:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v10644 = m.ExcPending
	if v10644 != 0 {
		goto L1
	} else {
		goto L2542
	}
L2542:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10646 = m.ExcPending
	if v10646 != 0 {
		goto L1
	} else {
		goto L2543
	}
L2543:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(1855), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10651 = m.ExcPending
	if v10651 != 0 {
		goto L1
	} else {
		goto L2544
	}
L2544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2545:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v10658 = m.ExcPending
	if v10658 != 0 {
		goto L1
	} else {
		goto L2546
	}
L2546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+160)) = v7532 + int32(1)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_238), v39+int32(160))
	mBase = m.M
	v10666 = m.ExcPending
	if v10666 != 0 {
		goto L1
	} else {
		goto L2547
	}
L2547:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2051), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10671 = m.ExcPending
	if v10671 != 0 {
		goto L1
	} else {
		goto L2548
	}
L2548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2549:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v10678 = m.ExcPending
	if v10678 != 0 {
		goto L1
	} else {
		goto L2550
	}
L2550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = base.I32_extend16_s(v7604)
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_219), v39+int32(80))
	mBase = m.M
	v10685 = m.ExcPending
	if v10685 != 0 {
		goto L1
	} else {
		goto L2551
	}
L2551:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2058), int32(_a_F_PostgresMainLoopOnce_194))
	mBase = m.M
	v10690 = m.ExcPending
	if v10690 != 0 {
		goto L1
	} else {
		goto L2552
	}
L2552:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2553:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10697 = m.ExcPending
	if v10697 != 0 {
		goto L1
	} else {
		goto L2554
	}
L2554:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), int32(0))
	mBase = m.M
	v10701 = m.ExcPending
	if v10701 != 0 {
		goto L1
	} else {
		goto L2555
	}
L2555:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_232), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10706 = m.ExcPending
	if v10706 != 0 {
		goto L1
	} else {
		goto L2556
	}
L2556:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2557:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v10713 = m.ExcPending
	if v10713 != 0 {
		goto L1
	} else {
		goto L2558
	}
L2558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v8176
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_239), v39+int32(224))
	mBase = m.M
	v10719 = m.ExcPending
	if v10719 != 0 {
		goto L1
	} else {
		goto L2559
	}
L2559:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2244), int32(_a_F_PostgresMainLoopOnce_203))
	mBase = m.M
	v10724 = m.ExcPending
	if v10724 != 0 {
		goto L1
	} else {
		goto L2560
	}
L2560:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2561:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10732 = m.ExcPending
	if v10732 != 0 {
		goto L1
	} else {
		goto L2562
	}
L2562:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v10736 = m.ExcPending
	if v10736 != 0 {
		goto L1
	} else {
		goto L2563
	}
L2563:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10738 = m.ExcPending
	if v10738 != 0 {
		goto L1
	} else {
		goto L2564
	}
L2564:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2360), int32(_a_F_PostgresMainLoopOnce_203))
	mBase = m.M
	v10743 = m.ExcPending
	if v10743 != 0 {
		goto L1
	} else {
		goto L2565
	}
L2565:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2566:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10750 = m.ExcPending
	if v10750 != 0 {
		goto L1
	} else {
		goto L2567
	}
L2567:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_240), int32(0))
	mBase = m.M
	v10754 = m.ExcPending
	if v10754 != 0 {
		goto L1
	} else {
		goto L2568
	}
L2568:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_241), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10759 = m.ExcPending
	if v10759 != 0 {
		goto L1
	} else {
		goto L2569
	}
L2569:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2570:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10766 = m.ExcPending
	if v10766 != 0 {
		goto L1
	} else {
		goto L2571
	}
L2571:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), int32(0))
	mBase = m.M
	v10770 = m.ExcPending
	if v10770 != 0 {
		goto L1
	} else {
		goto L2572
	}
L2572:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_232), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10775 = m.ExcPending
	if v10775 != 0 {
		goto L1
	} else {
		goto L2573
	}
L2573:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2574:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10782 = m.ExcPending
	if v10782 != 0 {
		goto L1
	} else {
		goto L2575
	}
L2575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+352)) = v10092
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_242), v39+int32(352))
	mBase = m.M
	v10788 = m.ExcPending
	if v10788 != 0 {
		goto L1
	} else {
		goto L2576
	}
L2576:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_243), int32(_a_F_PostgresMainLoopOnce_4))
	mBase = m.M
	v10793 = m.ExcPending
	if v10793 != 0 {
		goto L1
	} else {
		goto L2577
	}
L2577:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2578:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v10800 = m.ExcPending
	if v10800 != 0 {
		goto L1
	} else {
		goto L2579
	}
L2579:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_231), int32(0))
	mBase = m.M
	v10804 = m.ExcPending
	if v10804 != 0 {
		goto L1
	} else {
		goto L2580
	}
L2580:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(_a_F_PostgresMainLoopOnce_232), int32(_a_F_PostgresMainLoopOnce_233))
	mBase = m.M
	v10809 = m.ExcPending
	if v10809 != 0 {
		goto L1
	} else {
		goto L2581
	}
L2581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2582:
	;
	F_errcode(m, int32(386))
	mBase = m.M
	v10816 = m.ExcPending
	if v10816 != 0 {
		goto L1
	} else {
		goto L2583
	}
L2583:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_235), int32(0))
	mBase = m.M
	v10820 = m.ExcPending
	if v10820 != 0 {
		goto L1
	} else {
		goto L2584
	}
L2584:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2776), int32(_a_F_PostgresMainLoopOnce_244))
	mBase = m.M
	v10825 = m.ExcPending
	if v10825 != 0 {
		goto L1
	} else {
		goto L2585
	}
L2585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2586:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10832 = m.ExcPending
	if v10832 != 0 {
		goto L1
	} else {
		goto L2587
	}
L2587:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v10836 = m.ExcPending
	if v10836 != 0 {
		goto L1
	} else {
		goto L2588
	}
L2588:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10838 = m.ExcPending
	if v10838 != 0 {
		goto L1
	} else {
		goto L2589
	}
L2589:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2797), int32(_a_F_PostgresMainLoopOnce_244))
	mBase = m.M
	v10843 = m.ExcPending
	if v10843 != 0 {
		goto L1
	} else {
		goto L2590
	}
L2590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2591:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v10850 = m.ExcPending
	if v10850 != 0 {
		goto L1
	} else {
		goto L2592
	}
L2592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+384)) = v10149
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_239), v39+int32(384))
	mBase = m.M
	v10856 = m.ExcPending
	if v10856 != 0 {
		goto L1
	} else {
		goto L2593
	}
L2593:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2858), int32(_a_F_PostgresMainLoopOnce_245))
	mBase = m.M
	v10861 = m.ExcPending
	if v10861 != 0 {
		goto L1
	} else {
		goto L2594
	}
L2594:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2595:
	;
	F_errcode(m, int32(33685826))
	mBase = m.M
	v10868 = m.ExcPending
	if v10868 != 0 {
		goto L1
	} else {
		goto L2596
	}
L2596:
	;
	F_errmsg(m, int32(_a_F_PostgresMainLoopOnce_140), int32(0))
	mBase = m.M
	v10872 = m.ExcPending
	if v10872 != 0 {
		goto L1
	} else {
		goto L2597
	}
L2597:
	;
	F_errdetail_abort(m)
	mBase = m.M
	v10874 = m.ExcPending
	if v10874 != 0 {
		goto L1
	} else {
		goto L2598
	}
L2598:
	;
	F_errfinish(m, int32(_a_F_PostgresMainLoopOnce_2), int32(2874), int32(_a_F_PostgresMainLoopOnce_245))
	mBase = m.M
	v10879 = m.ExcPending
	if v10879 != 0 {
		goto L1
	} else {
		goto L2599
	}
L2599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2600:
	;
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
