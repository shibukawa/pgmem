package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CheckPointGuts(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int64
	_ = v498
	var v499 int64
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int64
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int64
	_ = v791
	var v793 int64
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int64
	_ = v994
	var v995 int64
	_ = v995
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int64
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1088 int64
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1139 int32
	_ = v1139
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1222 int32
	_ = v1222
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1362 int32
	_ = v1362
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1464 int64
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1471 float64
	_ = v1471
	var v1472 float64
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1514 float64
	_ = v1514
	var v1515 float64
	_ = v1515
	var v1517 float64
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int64
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int64
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int64
	_ = v1537
	var v1539 int64
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1549 float64
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1559 int64
	_ = v1559
	var v1561 int64
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1568 float64
	_ = v1568
	var v1572 float64
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int64
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1728 int64
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
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
	var v1755 int32
	_ = v1755
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1823 int64
	_ = v1823
	var v1824 int64
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1851 int64
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1919 int32
	_ = v1919
	var v1936 int32
	_ = v1936
	var v1937 int64
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2041 int64
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2049 int64
	_ = v2049
	var v2052 int64
	_ = v2052
	var v2058 int64
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2088 int64
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2101 int64
	_ = v2101
	var v2102 int64
	_ = v2102
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2126 int64
	_ = v2126
	var v2127 int64
	_ = v2127
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2141 int32
	_ = v2141
	var v2147 int64
	_ = v2147
	var v2148 int64
	_ = v2148
	var v2160 int32
	_ = v2160
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2200 int64
	_ = v2200
	var v2201 int64
	_ = v2201
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int64
	_ = v2261
	var v2263 int64
	_ = v2263
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2278 int64
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2285 int64
	_ = v2285
	var v2292 int64
	_ = v2292
	var v2297 int64
	_ = v2297
	var v2300 int64
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2376 int32
	_ = v2376
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2392 int32
	_ = v2392
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2417 int64
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2463 int32
	_ = v2463
	var v2465 int32
	_ = v2465
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2481 int32
	_ = v2481
	var v2486 int32
	_ = v2486
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2520 int32
	_ = v2520
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2539 int32
	_ = v2539
	var v2544 int32
	_ = v2544
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2558 int32
	_ = v2558
	var v2563 int32
	_ = v2563
	v3 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v23 = F_LWLockAcquire(m, v19+int32(3200), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v26+int32(3200))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = m.G0
	v33 = v31 - int32(1040)
	m.G0 = v33
	v37 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_0), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v48 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v55 = F_LWLockAcquire(m, v51+int32(_a_F_CheckPointGuts_1), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_2), int32(2124), int32(_a_F_CheckPointGuts_3))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	if int32(0) < v58 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v33 + int32(1040)
	v175 = m.G0
	v177 = v175 - int32(1152)
	m.G0 = v177
	v179 = F_GetRedoRecPtr(m)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	v65 = v3
	v66 = v62
	v67 = v58
	v69 = v3
	goto L15
L13:
	;
	goto L14
L14:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v150+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v82 = v66 + v65*int32(288)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v83 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v138+int32(_a_F_CheckPointGuts_1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L33
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_CheckPointGuts_4)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v82 + int32(24)
	v94 = F_pg_sprintf(m, v33+int32(16), int32(_a_F_CheckPointGuts_5), v33)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v131 = v66
	v132 = v67
	v133 = v69
	goto L19
L19:
	;
	v135 = v65 + int32(1)
	if v135 < v132 {
		v65 = v135
		v66 = v131
		v67 = v132
		v69 = v133
		goto L15
	} else {
		goto L32
	}
L20:
	;
	if l1&v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v82)+104))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v82)+280))
	F_SaveSlotToPath(m, v82, v33+int32(16), int32(15))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	if v98 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v103 = base.AtomicRmwXchg32(m, v82, int32(0), int32(1))
	if v103 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_s_lock(m, v82, int32(_a_F_CheckPointGuts_2), int32(2162), int32(_a_F_CheckPointGuts_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v82)+112))
	if v109 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v115 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v82))), uint32(v115))
	goto L21
L29:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v82)+120))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v82)+264))
	if base.Ui64(v110) <= base.Ui64(v111) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v113 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)) = uint16(v113)
	goto L28
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[1]))
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[2]))
	v131 = v130
	v132 = v128
	v133 = base.B2i32(v118 != v119) | v69
	goto L19
L32:
	;
	goto L16
L33:
	;
	if v133&int32(1) == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L11
L36:
	;
	goto L11
L37:
	;
	v181 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v184 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v187 = F_ReadDir(m, v184, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v187 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui64(v179) < base.Ui64(v181) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	F_FreeDir(m, v184)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L44:
	;
	v190 = v179
	goto L46
L45:
	;
	v190 = v181
	goto L46
L46:
	;
	v197 = v187
	goto L47
L47:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+19)))
	if v210 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L43
L49:
	;
	v332 = F_ReadDir(m, v184, int32(_a_F_CheckPointGuts_6))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L84
	}
L50:
	;
	v223 = v197 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+84)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v177)+80)) = int32(_a_F_CheckPointGuts_6)
	v228 = v177 + int32(96)
	v233 = F_pg_snprintf(m, v228, int32(1045), int32(_a_F_CheckPointGuts_5), v177+int32(80))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L55
	}
L51:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)))
	if v213 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)))
	if v216 != int32(46) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+21)))
	if v219 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v238 = F_get_dirent_type(m, v228, v197, int32(0), int32(14))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_7), v325, int32(_a_F_CheckPointGuts_8))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L83
	}
L57:
	;
	if v238&int32(-3) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v244 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+52)) = v177 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+48)) = v177 + int32(92)
	v264 = F_sscanf(m, v223, int32(_a_F_CheckPointGuts_9), v177+int32(48))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	if v244 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+64)) = v228
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_10), v177-int32(-64))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v325 = int32(2009)
	goto L56
L64:
	;
	if v264 != int32(2) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v270 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v283 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v177)+88)))
	v284 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v177)+92)))
	if base.Ui64(v190-int64(1)) < base.Ui64(v283|v284<<(uint(int64(32))%64)) {
		goto L49
	} else {
		goto L71
	}
L68:
	;
	if v270 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+32)) = v177 + int32(96)
	F_errmsg(m, int32(_a_F_CheckPointGuts_11), v177+int32(32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v325 = int32(2025)
	goto L56
L71:
	;
	v291 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v291 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+16)) = v177 + int32(96)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_12), v177+int32(16))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v307 = v177 + int32(96)
	v308 = F_unlink(m, v307)
	mBase = m.M
	if int32(0) <= v308 {
		goto L49
	} else {
		goto L78
	}
L76:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_7), int32(2034), int32(_a_F_CheckPointGuts_8))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v313 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v313 == int32(0) {
		goto L49
	} else {
		goto L80
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v307
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v177)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v325 = int32(2046)
	goto L56
L83:
	;
	goto L49
L84:
	;
	if v332 != 0 {
		v197 = v332
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L48
L86:
	;
	m.G0 = v177 + int32(1152)
	v356 = m.G0
	v358 = v356 - int32(1216)
	m.G0 = v358
	v360 = F_GetRedoRecPtr(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v362 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v365 = F_AllocateDir(m, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L93
	}
L89:
	;
	v684 = m.G0
	v686 = v684 - int32(112)
	m.G0 = v686
	*(*int32)(unsafe.Add(mBase, uint32(v686)+108)) = int32(307747550)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	if v691 != 0 {
		goto L183
	} else {
		goto L184
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L175
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L171
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L168
	}
L93:
	;
	v368 = F_ReadDir(m, v365, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v368 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.Ui64(v360) < base.Ui64(v362) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	F_FreeDir(m, v365)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L166
	}
L98:
	;
	v371 = v360
	goto L100
L99:
	;
	v371 = v362
	goto L100
L100:
	;
	v384 = v368
	goto L101
L101:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+19)))
	if v395 != int32(46) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L97
L103:
	;
	v604 = F_ReadDir(m, v365, int32(_a_F_CheckPointGuts_14))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L164
	}
L104:
	;
	v408 = v384 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+132)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v358)+128)) = int32(_a_F_CheckPointGuts_14)
	v413 = v358 + int32(160)
	v418 = F_pg_snprintf(m, v413, int32(1044), int32(_a_F_CheckPointGuts_5), v358+int32(128))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+20)))
	if v398 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+20)))
	if v401 != int32(46) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+21)))
	if v404 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v422 = F_get_dirent_type(m, v413, v384, int32(0), int32(14))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v422&int32(-3) != 0 {
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v426 = int32(_a_F_CheckPointGuts_15)
	goto L114
L112:
	;
	if v464-v465 != 0 {
		goto L103
	} else {
		goto L125
	}
L114:
	;
	goto L115
L115:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if v433 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v434 = v408
	v435 = v426
	v436 = int32(4)
	v437 = v433
	goto L120
L117:
	;
	v460 = v426
	v464 = int32(0)
	goto L118
L118:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	goto L112
L119:
	;
	v460 = v455
	v464 = v457
	goto L118
L120:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if base.B2i32(v437 != v439)|base.B2i32(v439 == int32(0)) != 0 {
		v455 = v435
		v457 = v437
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v455 = v449
	v457 = int32(0)
	goto L119
L122:
	;
	v445 = v436 - int32(1)
	if v445 == int32(0) {
		v455 = v435
		v457 = v437
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v448 = int32(1)
	v449 = v435 + v448
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)))
	if v450 != 0 {
		v434 = v434 + v448
		v435 = v449
		v436 = v445
		v437 = v450
		goto L120
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358+int32(116)))) = v358 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v358+int32(112)))) = v358 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+108)) = v358 + int32(136)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+104)) = v358 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+100)) = v358 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+96)) = v358 + int32(156)
	v494 = F_sscanf(m, v408, int32(_a_F_CheckPointGuts_16), v358+int32(96))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v494 != int32(6) {
		goto L92
	} else {
		goto L127
	}
L127:
	;
	v498 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v358)+136)))
	v499 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v358)+140)))
	if base.Ui64(v498|v499<<(uint(int64(32))%64)) <= base.Ui64(v371-int64(1)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v506 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v542 = v358 + int32(160)
	v544 = F_OpenTransientFile(m, v542, int32(2))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L142
	}
L131:
	;
	if v506 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+64)) = v413
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_17), v358-int32(-64))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v520 = v358 + int32(160)
	v521 = F_unlink(m, v520)
	mBase = m.M
	if int32(0) <= v521 {
		goto L103
	} else {
		goto L137
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1210), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+48)) = v520
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v358+int32(48))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1214), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
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
	if v544 < int32(0) {
		goto L91
	} else {
		goto L143
	}
L143:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = int32(167772194)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v554 != int32(1) {
		v568 = int32(0)
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v595))) = int32(0)
	v598 = F_CloseTransientFile(m, v544)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L162
	}
L145:
	;
	if v568 == int32(0) {
		goto L144
	} else {
		goto L152
	}
L146:
	;
	goto L145
L147:
	;
	goto L148
L148:
	;
	v559 = F_fsync(m, v544)
	mBase = m.M
	if v559 != int32(-1) {
		v568 = v559
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v568 = int32(-1)
	goto L146
L150:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v563 == int32(27) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v574 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v577 = F_errstart(m, v575, int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L157
	}
L154:
	;
	v575 = int32(21)
	goto L156
L155:
	;
	v575 = int32(23)
	goto L156
L156:
	;
	goto L153
L157:
	;
	if v577 == int32(0) {
		goto L144
	} else {
		goto L158
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+32)) = v542
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v358+int32(32))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1240), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L144
L162:
	;
	if v598 != 0 {
		goto L90
	} else {
		goto L163
	}
L163:
	;
	goto L103
L164:
	;
	if v604 != 0 {
		v384 = v604
		goto L101
	} else {
		goto L165
	}
L165:
	;
	goto L102
L166:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_14), int32(1))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v358 + int32(1216)
	goto L89
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+80)) = v408
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_21), v358+int32(80))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1204), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v358 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_22), v358)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1229), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v358 + int32(160)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v358+int32(16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_18), int32(1246), int32(_a_F_CheckPointGuts_19))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
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
	v989 = m.G0
	v990 = int32(16)
	v991 = v989 - v990
	m.G0 = v991
	F_gettimeofday(m, v991)
	mBase = m.M
	v994 = *(*int64)(unsafe.Add(mBase, uint32(v991)))
	v995 = int64(*(*int32)(unsafe.Add(mBase, uint32(v991)+8)))
	m.G0 = v991 + v990
	goto L251
L180:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L247
	}
L181:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L243
	}
L182:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L239
	}
L183:
	;
	v693 = F_unlink(m, int32(_a_F_CheckPointGuts_24))
	mBase = m.M
	if v693 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L185
L185:
	;
	m.G0 = v686 + int32(112)
	goto L179
L186:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v697 != int32(44) {
		goto L182
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v702 = F_OpenTransientFile(m, int32(_a_F_CheckPointGuts_24), int32(193))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	if v702 < int32(0) {
		goto L181
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v711 = int32(4)
	v712 = F_write(m, v702, v686+int32(108), v711)
	mBase = m.M
	if v712 != v711 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v716 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v744 = m.Env.Pgmem_crc32c(m, int32(-1), v686+int32(108), int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v686)+104)) = v744
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v751 = F_LWLockAcquire(m, v747+int32(_a_F_CheckPointGuts_25), int32(1))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L202
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L197
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+64)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v686-int32(-64))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(639), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
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
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	if int32(0) < v754 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[8]))
	v762 = int32(0)
	v763 = v744
	v764 = v758
	v765 = v754
	goto L206
L204:
	;
	v850 = v744
	goto L205
L205:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v864+int32(_a_F_CheckPointGuts_25))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L225
	}
L206:
	;
	v778 = v764 + v762*int32(56)
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778))))
	if v779 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v850 = v840
	goto L205
L208:
	;
	v780 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v686)+96)) = v780
	*(*int64)(unsafe.Add(mBase, uint32(v686)+88)) = v780
	v785 = v778 + int32(40)
	v787 = F_LWLockAcquire(m, v785, int32(1))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v840 = v763
	v841 = v764
	v842 = v765
	goto L210
L210:
	;
	v844 = v762 + int32(1)
	if v844 < v842 {
		v762 = v844
		v763 = v840
		v764 = v841
		v765 = v842
		goto L206
	} else {
		goto L224
	}
L211:
	;
	v789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778))))
	*(*uint16)(unsafe.Add(mBase, uint32(v686)+88)) = uint16(v789)
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v778)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v686)+96)) = v791
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v778)+16))
	F_LWLockRelease(m, v785)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_XLogFlush(m, v793)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v803 = int32(16)
	v804 = F_write(m, v702, v686+int32(88), v803)
	mBase = m.M
	if v804 != v803 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v808 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v835 = m.Env.Pgmem_crc32c(m, v763, v686+int32(88), int32(16))
	mBase = m.M
	v837 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[8]))
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[3]))
	v840 = v835
	v841 = v837
	v842 = v839
	goto L210
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L219
L218:
	;
	goto L219
L219:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+48)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v686+int32(48))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(681), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
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
	goto L207
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+104)) = v850 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v877 = int32(4)
	v878 = F_write(m, v702, v686+int32(104), v877)
	mBase = m.M
	if v878 != v877 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v882 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v906 = F_CloseTransientFile(m, v702)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L236
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L231
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+32)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_26), v686+int32(32))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(700), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	if v906 != 0 {
		goto L180
	} else {
		goto L237
	}
L237:
	;
	v911 = F_durable_rename(m, int32(_a_F_CheckPointGuts_24), int32(_a_F_CheckPointGuts_29), int32(23))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	goto L185
L239:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+80)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_13), v686+int32(80))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(615), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686))) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_30), v686)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(627), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v686)+16)) = int32(_a_F_CheckPointGuts_24)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v686+int32(16))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_27), int32(707), int32(_a_F_CheckPointGuts_28))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[9])) = v995 + v994*int64(1000000) - int64(946684800000000)
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_31))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_32))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_33))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_34))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_35))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v1025 = F_LWLockAcquire(m, v1021+int32(_a_F_CheckPointGuts_36), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[10]))
	v1029 = *(*int64)(unsafe.Add(mBase, uint32(v1028)))
	if v1029 < int64(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1105 = int32(0)
	v1108 = m.G0
	v1110 = v1108 - int32(_a_F_CheckPointGuts_37)
	m.G0 = v1110
	v1113 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[11]))
	if v1113 <= v1105 {
		goto L281
	} else {
		goto L282
	}
L259:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1033+int32(_a_F_CheckPointGuts_36))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+12))
	if v1038 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L258
L263:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v1090+int32(_a_F_CheckPointGuts_36))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L278
	}
L264:
	;
	v1041 = int32(4)
	v1042 = v1038&int32(-1024) | v1041
	v1045 = base.I32_wrap_i64(v1029) << (uint(int32(10)) % 32)
	v1047 = v1045 | v1041
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1047))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1042)) == int32(0) {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	goto L266
L266:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1028))) = int64(-1)
	v1088 = v1029
	goto L263
L267:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[10]))
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v1082)))
	v1088 = v1083
	goto L263
L268:
	;
	if v1059 == int32(0) {
		goto L267
	} else {
		goto L272
	}
L269:
	;
	v1059 = base.B2i32(base.Ui32(v1042) < base.Ui32(v1047))
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1059 = int32(base.Ui32(v1042-v1047) >> (uint(int32(31)) % 32))
	goto L268
L272:
	;
	v1063 = v1045 + int32(1027)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1063))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1042)) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	if v1075 == int32(0) {
		goto L267
	} else {
		goto L277
	}
L274:
	;
	v1075 = base.B2i32(base.Ui32(v1042) < base.Ui32(v1063))
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1075 = int32(base.Ui32(v1042-v1063) >> (uint(int32(31)) % 32))
	goto L273
L277:
	;
	v1088 = base.I64_extend_i32_u(int32(base.Ui32(v1038) >> (uint(int32(10)) % 32)))
	goto L263
L278:
	;
	F_SimpleLruTruncate(m, int32(_a_F_CheckPointGuts_38), v1088)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_SimpleLruWriteAll(m, int32(_a_F_CheckPointGuts_38))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L258
L281:
	;
	m.G0 = v1110 + int32(_a_F_CheckPointGuts_37)
	v1712 = m.G0
	v1713 = int32(16)
	v1714 = v1712 - v1713
	m.G0 = v1714
	F_gettimeofday(m, v1714)
	mBase = m.M
	v1717 = *(*int64)(unsafe.Add(mBase, uint32(v1714)))
	v1718 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1714)+8)))
	m.G0 = v1714 + v1713
	goto L413
L282:
	;
	if l1&int32(19) != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1120 = int32(-8388609)
	goto L285
L284:
	;
	v1120 = int32(2139095039)
	goto L285
L285:
	;
	v1129 = v1105
	v1130 = v1105
	goto L286
L286:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+28)) = int32(_a_F_CheckPointGuts_39)
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+24)) = int32(_a_F_CheckPointGuts_40)
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+20)) = int32(_a_F_CheckPointGuts_41)
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1110)+8)) = int64(0)
	v1152 = v1139 + v1129<<(uint(int32(6))%32)
	v1153 = int32(_a_F_CheckPointGuts_42)
	v1155 = base.AtomicRmwOr32(m, v1152, int32(24), v1153)
	if v1155&v1153 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if v1246 == int32(0) {
		goto L281
	} else {
		goto L314
	}
L288:
	;
	goto L291
L289:
	;
	v1189 = v1155
	goto L290
L290:
	;
	v1204 = int32(_a_F_CheckPointGuts_43)
	v1205 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13]))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1110+int32(8))+8))
	if v1207 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	F_perform_spin_delay(m, v1110+int32(8))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	v1189 = v1181
	goto L290
L293:
	;
	v1179 = int32(_a_F_CheckPointGuts_42)
	v1181 = base.AtomicRmwOr32(m, v1152, int32(24), v1179)
	if v1181&v1179 != 0 {
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	if v1189|v1120 == int32(-1) {
		goto L306
	} else {
		goto L307
	}
L296:
	;
	goto L295
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[13])) = v1222
	goto L296
L298:
	;
	if int32(999) < v1205 {
		goto L296
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	if v1205 < int32(11) {
		goto L296
	} else {
		goto L305
	}
L301:
	;
	v1212 = int32(900)
	if v1212 <= v1205 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1215 = v1212
	goto L304
L303:
	;
	v1215 = v1205
	goto L304
L304:
	;
	v1222 = v1215 + int32(100)
	goto L297
L305:
	;
	v1222 = v1205 - int32(1)
	goto L297
L306:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1231 = v1228 + v1130*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+16)) = v1129
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v1233
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+4)) = v1235
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+8)) = v1237
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+12)) = v1239
	v1246 = v1130 + int32(1)
	v1247 = v1189 | int32(1077936128)
	goto L308
L307:
	;
	v1246 = v1130
	v1247 = v1189
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1152)+24)) = v1247 & int32(-4194305)
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1252 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1256 = v1129 + int32(1)
	v1258 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[11]))
	if v1256 < v1258 {
		v1129 = v1256
		v1130 = v1246
		goto L286
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	goto L287
L314:
	;
	v1262 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+12)) = v1262
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+8)) = int32(_a_F_CheckPointGuts_44)
	v1268 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	F_sort_checkpoint_bufferids(m, v1268, v1246)
	mBase = m.M
	if v1262 < v1246 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	F_binaryheap_build(m, v1399)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L346
	}
L316:
	;
	v1272 = int32(0)
	v1276 = v1262
	v1277 = v1105
	v1281 = v1272
	v1282 = v1272
	goto L319
L317:
	;
	goto L318
L318:
	;
	v1386 = int32(0)
	v1390 = F_binaryheap_allocate(m, v1386, int32(1077), v1386)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L345
	}
L319:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1293+v1281*int32(20))))
	if v1277 == v1297 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1349 = int32(0)
	v1352 = F_binaryheap_allocate(m, v1333, int32(1077), v1349)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L339
	}
L321:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+24)) = v1338 + int32(1)
	v1343 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1343 != 0 {
		goto L334
	} else {
		goto L335
	}
L322:
	;
	v1299 = v1277
	goto L324
L323:
	;
	v1299 = int32(0)
	goto L324
L324:
	;
	if v1299 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1303 = v1276 + int32(1)
	v1305 = v1303 * int32(40)
	if v1282 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L326:
	;
	goto L327
L327:
	;
	v1328 = int32(40)
	v1333 = v1276
	v1334 = v1277
	v1335 = v1282 + v1276*v1328 - v1328
	v1337 = v1282
	goto L321
L328:
	;
	v1315 = v1312 + v1276*int32(40)
	v1316 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1315)+32)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v1315))) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v1315)+24)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v1315)+16)) = v1316
	*(*int64)(unsafe.Add(mBase, uint32(v1315)+8)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+32)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v1315))) = v1297
	v1333 = v1303
	v1334 = v1297
	v1335 = v1315
	v1337 = v1312
	goto L321
L329:
	;
	v1308 = F_palloc(m, v1305)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1310 = F_repalloc(m, v1282, v1305)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L333
	}
L332:
	;
	v1312 = v1308
	goto L328
L333:
	;
	v1312 = v1310
	goto L328
L334:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1347 = v1281 + int32(1)
	if v1347 != v1246 {
		v1276 = v1333
		v1277 = v1334
		v1281 = v1347
		v1282 = v1337
		goto L319
	} else {
		goto L338
	}
L337:
	;
	goto L336
L338:
	;
	goto L320
L339:
	;
	if v1333 <= int32(0) {
		v1399 = v1352
		v1400 = v1337
		goto L315
	} else {
		goto L340
	}
L340:
	;
	v1362 = v1349
	goto L341
L341:
	;
	v1376 = v1337 + v1362*int32(40)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v1376)+16)) = base.F64_div(base.F64_convert_i32_u(v1246), base.F64_convert_i32_s(v1377))
	F_binaryheap_add_unordered(m, v1352, v1376)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L343
	}
L342:
	;
	v1399 = v1352
	v1400 = v1337
	goto L315
L343:
	;
	v1384 = v1362 + int32(1)
	if v1384 != v1333 {
		v1362 = v1384
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1399 = v1390
	v1400 = v1386
	goto L315
L346:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	if v1411 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	F_IssuePendingWritebacks(m, v1110+int32(8), int32(3))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L1
	} else {
		goto L410
	}
L348:
	;
	v1659 = int32(0)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1416 = int32(0)
	v1420 = v1416
	v1422 = v1416
	goto L351
L351:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[12]))
	v1438 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[14]))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1399)+20))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+32))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1438+v1440*int32(20))+16))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1436+v1444<<(uint(int32(6))%32))+24))
	if v1448&int32(1073741824) == int32(0) {
		v1470 = v1420
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1659 = v1470
	goto L347
L353:
	;
	v1471 = *(*float64)(unsafe.Add(mBase, uint32(v1439)+16))
	v1472 = *(*float64)(unsafe.Add(mBase, uint32(v1439)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1439)+8)) = base.F64_add(v1471, v1472)
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+28))
	v1476 = int32(1)
	v1477 = v1475 + v1476
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+28)) = v1477
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1439)+32)) = v1479 + v1476
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+24))
	if v1483 == v1477 {
		goto L358
	} else {
		goto L359
	}
L354:
	;
	v1456 = F_SyncOneBuffer(m, v1444, int32(0), v1110+int32(8))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	if v1456&int32(1) == int32(0) {
		v1470 = v1420
		goto L353
	} else {
		goto L356
	}
L356:
	;
	v1462 = int32(_a_F_CheckPointGuts_45)
	v1464 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[16])) = v1464 + int64(1)
	v1470 = v1420 + int32(1)
	goto L353
L357:
	;
	v1490 = v1422 + int32(1)
	v1494 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[17]))
	if v1494 != int32(11) {
		goto L363
	} else {
		goto L364
	}
L358:
	;
	v1485 = F_binaryheap_remove_first(m, v1399)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_binaryheap_replace_first(m, v1399, v1439)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L1
	} else {
		goto L362
	}
L361:
	;
	goto L357
L362:
	;
	goto L357
L363:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	if v1656 != 0 {
		v1420 = v1470
		v1422 = v1490
		goto L351
	} else {
		goto L409
	}
L364:
	;
	if l1&int32(4) != 0 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[15]))
	if v1648 == int32(0) {
		goto L363
	} else {
		goto L407
	}
L366:
	;
	v1631 = int32(_a_F_CheckPointGuts_46)
	v1633 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18]))
	v1635 = v1633 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = v1635
	if int32(0) < v1635 {
		goto L365
	} else {
		goto L405
	}
L367:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[19]))
	if v1500 != 0 {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[20]))
	if v1502 != 0 {
		goto L366
	} else {
		goto L369
	}
L369:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[21]))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+20))
	if v1505&int32(4) != 0 {
		goto L366
	} else {
		goto L370
	}
L370:
	;
	v1508 = m.G0
	v1510 = v1508 - int32(16)
	m.G0 = v1510
	v1514 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[22]))
	v1515 = base.F64_mul(base.F64_div(base.F64_convert_i32_s(v1490), base.F64_convert_i32_s(v1246)), v1514)
	v1517 = *(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[23]))
	if base.F64_lt(v1515, v1517) != 0 {
		v1577 = int32(0)
		goto L371
	} else {
		goto L372
	}
L371:
	;
	m.G0 = v1510 + int32(16)
	if v1577 == int32(0) {
		goto L366
	} else {
		goto L387
	}
L372:
	;
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[24])))
	if v1521 == int32(1) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	v1539 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[25]))
	v1543 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[26]))
	v1547 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[27]))
	v1549 = base.F64_div(base.F64_div(base.F64_convert_i64_u(v1537-v1539), base.F64_convert_i32_s(v1543)), base.F64_convert_i32_s(v1547))
	if base.F64_lt(v1515, v1549) == int32(0) {
		goto L383
	} else {
		goto L384
	}
L374:
	;
	if v1531 != 0 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[28]))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1526)+316))
	v1529 = base.B2i32(v1527 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[24])) = uint8(v1529)
	v1531 = v1529
	goto L377
L376:
	;
	v1531 = int32(0)
	goto L377
L377:
	;
	goto L374
L378:
	;
	v1533 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1535 = F_GetInsertRecPtr(m)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L382
	}
L381:
	;
	v1537 = v1533
	goto L373
L382:
	;
	v1537 = v1535
	goto L373
L383:
	;
	F_gettimeofday(m, v1510)
	mBase = m.M
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+8))
	v1559 = *(*int64)(unsafe.Add(mBase, uint32(v1510)))
	v1561 = *(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[29]))
	v1566 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[30]))
	v1568 = base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v1555), float64(1e+06)), base.F64_convert_i64_s(v1559-v1561)), base.F64_convert_i32_s(v1566))
	if base.F64_lt(v1515, v1568) == int32(0) {
		v1577 = int32(1)
		goto L371
	} else {
		goto L386
	}
L384:
	;
	v1572 = v1549
	goto L385
L385:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CheckPointGuts[23])) = v1572
	v1577 = int32(0)
	goto L371
L386:
	;
	v1572 = v1568
	goto L385
L387:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[31]))
	if v1584 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[31])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L1
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L1
	} else {
		goto L400
	}
L391:
	;
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1597 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	if v1597 != 0 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_47), int32(0))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	goto L390
L398:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_48), int32(1402), int32(_a_F_CheckPointGuts_49))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = int32(1000)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32]))
	v1622 = F_WaitLatch(m, v1618, int32(41), int32(100), int32(150994945))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1625))) = int32(0)
	goto L404
L404:
	;
	goto L365
L405:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[18])) = int32(1000)
	goto L365
L407:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	goto L363
L409:
	;
	goto L352
L410:
	;
	F_pfree(m, v1400)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_pfree(m, v1399)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v1683 = int32(_a_F_CheckPointGuts_50)
	v1685 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33]))
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[33])) = v1685 + v1659
	goto L281
L413:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[34])) = v1718 + v1717*int64(1000000) - int64(946684800000000)
	v1728 = int64(0)
	v1729 = int32(0)
	v1731 = m.G0
	v1733 = v1731 - int32(1152)
	m.G0 = v1733
	v1736 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	if v1736 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	v2195 = m.G0
	v2196 = int32(16)
	v2197 = v2195 - v2196
	m.G0 = v2197
	F_gettimeofday(m, v2197)
	mBase = m.M
	v2200 = *(*int64)(unsafe.Add(mBase, uint32(v2197)))
	v2201 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2197)+8)))
	m.G0 = v2197 + v2196
	goto L516
L415:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L513
	}
L418:
	;
	v1740 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])))
	if v1740 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1795 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])) = uint8(v1795)
	v1797 = int32(_a_F_CheckPointGuts_51)
	v1799 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	v1801 = v1799 + v1795
	*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])) = uint16(v1801)
	v1804 = v1733 + int32(1116)
	v1806 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	F_hash_seq_init(m, v1804, v1806)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L1
	} else {
		goto L428
	}
L420:
	;
	v1744 = v1733 + int32(1116)
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	F_hash_seq_init(m, v1744, v1746)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v1749 = F_hash_seq_search(m, v1744)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	if v1749 == int32(0) {
		goto L419
	} else {
		goto L423
	}
L423:
	;
	v1755 = v1749
	goto L424
L424:
	;
	v1771 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1755)+24)) = uint16(v1771)
	v1775 = F_hash_seq_search(m, v1733+int32(1116))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L1
	} else {
		goto L426
	}
L425:
	;
	goto L419
L426:
	;
	if v1775 != 0 {
		v1755 = v1775
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	v1809 = F_hash_seq_search(m, v1804)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L430
	}
L429:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L510
	}
L430:
	;
	if v1809 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1814 = v1809
	v1816 = int32(10)
	v1817 = v1729
	v1823 = v1728
	v1824 = v1728
	goto L434
L432:
	;
	v2141 = v1729
	v2147 = v1728
	v2148 = v1728
	goto L433
L433:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[38])) = v2147
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[39])) = v2148
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[40])) = v2141
	v2160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[36])) = uint8(v2160)
	m.G0 = v1733 + int32(1152)
	goto L414
L434:
	;
	v1829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1814)+24)))
	v1831 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CheckPointGuts[37])))
	if v1829 != v1831 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v2141 = v2120
	v2147 = v2126
	v2148 = v2127
	goto L433
L436:
	;
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v1834 != int32(1) {
		v2094 = v1816
		v2095 = v1817
		v2101 = v1823
		v2102 = v1824
		goto L439
	} else {
		goto L440
	}
L437:
	;
	v2119 = v1816
	v2120 = v1817
	v2126 = v1823
	v2127 = v1824
	goto L438
L438:
	;
	v2134 = F_hash_seq_search(m, v1733+int32(1116))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L508
	}
L439:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[35]))
	v2111 = F_hash_search(m, v2108, v1814, int32(2), int32(0))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L1
	} else {
		goto L506
	}
L440:
	;
	v1838 = v1816 - int32(1)
	if v1838 <= int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L1
	} else {
		goto L444
	}
L442:
	;
	v1844 = v1838
	goto L443
L443:
	;
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+26)))
	if v1845 != 0 {
		v2094 = v1844
		v2095 = v1817
		v2101 = v1823
		v2102 = v1824
		goto L439
	} else {
		goto L445
	}
L444:
	;
	v1844 = int32(10)
	goto L443
L445:
	;
	F___clock_gettime(m, int32(1), v1733+int32(1136))
	mBase = m.M
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+1144))
	v1851 = *(*int64)(unsafe.Add(mBase, uint32(v1733)+1136))
	v1853 = v1733 + int32(80)
	v1854 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1814))))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1854*int32(12))+uint32(_c_F_CheckPointGuts[41])))
	v1860 = m.T0[v1859].(func(*base.Module, int32, int32) int32)(m, v1814, v1853)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L447
	}
L446:
	;
	v2045 = int32(1)
	F___clock_gettime(m, v2045, v1733+int32(1136))
	mBase = m.M
	v2049 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1733)+1144)))
	v2052 = *(*int64)(unsafe.Add(mBase, uint32(v1733)+1136))
	v2058 = base.I64_div_s(v2049-base.I64_extend_i32_s(v2031)+(v2052-v2041)*int64(1000000000), int64(1000))
	v2061 = v1817 + v2045
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[42])))
	if v2063 != v2045 {
		goto L497
	} else {
		goto L498
	}
L447:
	;
	if v1860 == int32(0) {
		v2031 = v1850
		v2032 = v1844
		v2041 = v1851
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v1866 == int32(44) {
		goto L451
	} else {
		goto L452
	}
L449:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L1
	} else {
		goto L467
	}
L450:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), v1905, int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L466
	}
L451:
	;
	v1871 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v1887 != 0 {
		goto L459
	} else {
		goto L460
	}
L454:
	;
	if v1871 == int32(0) {
		goto L449
	} else {
		goto L455
	}
L455:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+48)) = v1853
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_54), v1733+int32(48))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1905 = int32(452)
	goto L450
L458:
	;
	v1890 = F_errstart(m, v1888, int32(0))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L1
	} else {
		goto L462
	}
L459:
	;
	v1888 = int32(21)
	goto L461
L460:
	;
	v1888 = int32(23)
	goto L461
L461:
	;
	goto L458
L462:
	;
	if v1890 == int32(0) {
		goto L449
	} else {
		goto L463
	}
L463:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+64)) = v1733 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v1733-int32(-64))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v1905 = int32(447)
	goto L450
L466:
	;
	goto L449
L467:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+26)))
	if v1912 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1919 = int32(1)
	goto L471
L469:
	;
	goto L470
L470:
	;
	v2094 = int32(10)
	v2095 = v1817
	v2101 = v1823
	v2102 = v1824
	goto L439
L471:
	;
	F___clock_gettime(m, int32(1), v1733+int32(1136))
	mBase = m.M
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+1144))
	v1937 = *(*int64)(unsafe.Add(mBase, uint32(v1733)+1136))
	v1940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1814))))
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1940*int32(12))+uint32(_c_F_CheckPointGuts[41])))
	v1946 = m.T0[v1945].(func(*base.Module, int32, int32) int32)(m, v1814, v1733+int32(80))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L1
	} else {
		goto L473
	}
L472:
	;
	goto L470
L473:
	;
	if v1946 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2031 = v1936
	v2032 = int32(10)
	v2041 = v1937
	goto L446
L475:
	;
	goto L476
L476:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	v1956 = int32(0)
	if base.B2i32(v1953 == int32(44))&base.B2i32(v1919 <= v1956) == v1956 {
		goto L479
	} else {
		goto L480
	}
L477:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L495
	}
L478:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), v1999, int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L494
	}
L479:
	;
	v1964 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[7])))
	if v1964 != 0 {
		goto L483
	} else {
		goto L484
	}
L480:
	;
	goto L481
L481:
	;
	v1984 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L490
	}
L482:
	;
	v1967 = F_errstart(m, v1965, int32(0))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L486
	}
L483:
	;
	v1965 = int32(21)
	goto L485
L484:
	;
	v1965 = int32(23)
	goto L485
L485:
	;
	goto L482
L486:
	;
	if v1967 == int32(0) {
		goto L477
	} else {
		goto L487
	}
L487:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+16)) = v1733 + int32(80)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v1733+int32(16))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	v1999 = int32(447)
	goto L478
L490:
	;
	if v1984 == int32(0) {
		goto L477
	} else {
		goto L491
	}
L491:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+32)) = v1733 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_54), v1733+int32(32))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v1999 = int32(452)
	goto L478
L494:
	;
	goto L477
L495:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+26)))
	if v2007 != int32(1) {
		v1919 = v1919 + int32(1)
		goto L471
	} else {
		goto L496
	}
L496:
	;
	goto L472
L497:
	;
	if base.Ui64(v1824) < base.Ui64(v2058) {
		goto L503
	} else {
		goto L504
	}
L498:
	;
	v2068 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	if v2068 == int32(0) {
		goto L497
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1733))) = v2061
	*(*float64)(unsafe.Add(mBase, uint32(v1733)+8)) = base.F64_div(base.F64_convert_i64_u(v2058), float64(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v1733)+4)) = v1733 + int32(80)
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_55), v1733)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(432), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	goto L497
L503:
	;
	v2088 = v2058
	goto L505
L504:
	;
	v2088 = v1824
	goto L505
L505:
	;
	v2094 = v2032
	v2095 = v2061
	v2101 = v1823 + v2058
	v2102 = v2088
	goto L439
L506:
	;
	if v2111 == int32(0) {
		goto L429
	} else {
		goto L507
	}
L507:
	;
	v2119 = v2094
	v2120 = v2095
	v2126 = v2101
	v2127 = v2102
	goto L438
L508:
	;
	if v2134 != 0 {
		v1814 = v2134
		v1816 = v2119
		v1817 = v2120
		v1823 = v2126
		v1824 = v2127
		goto L434
	} else {
		goto L509
	}
L509:
	;
	goto L435
L510:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_56), int32(0))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(465), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L513:
	;
	F_errmsg_internal(m, int32(_a_F_CheckPointGuts_57), int32(0))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_52), int32(308), int32(_a_F_CheckPointGuts_53))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
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
	*(*int64)(unsafe.Add(mBase, _c_F_CheckPointGuts[43])) = v2201 + v2200*int64(1000000) - int64(946684800000000)
	v2211 = int32(0)
	v2213 = m.G0
	v2215 = v2213 - int32(1152)
	m.G0 = v2215
	v2218 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[44]))
	if v2218 <= v2211 {
		goto L521
	} else {
		goto L522
	}
L517:
	;
	return
L518:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L1
	} else {
		goto L594
	}
L519:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L590
	}
L520:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L1
	} else {
		goto L586
	}
L521:
	;
	m.G0 = v2215 + int32(1152)
	goto L517
L522:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	v2226 = F_LWLockAcquire(m, v2222+int32(2304), int32(1))
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[45]))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2229)+4))
	if int32(0) < v2230 {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2235 = v2229
	v2236 = v2211
	v2238 = v2211
	goto L527
L525:
	;
	v2440 = v2211
	goto L526
L526:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[0]))
	F_LWLockRelease(m, v2455+int32(2304))
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L1
	} else {
		goto L579
	}
L527:
	;
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v2235+v2238<<(uint(int32(2))%32))+8))
	v2254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2253)+44)))
	if v2254 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L528:
	;
	v2440 = v2428
	goto L526
L529:
	;
	v2434 = v2238 + int32(1)
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2427)+4))
	if v2434 < v2435 {
		v2235 = v2427
		v2236 = v2428
		v2238 = v2434
		goto L527
	} else {
		goto L578
	}
L530:
	;
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2253)+46)))
	if v2257 != int32(1) {
		v2427 = v2235
		v2428 = v2236
		goto L529
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2253)+45)))
	if v2260 != 0 {
		v2427 = v2235
		v2428 = v2236
		goto L529
	} else {
		goto L534
	}
L533:
	;
	goto L532
L534:
	;
	v2261 = *(*int64)(unsafe.Add(mBase, uint32(v2253)+24))
	if base.Ui64(l0) < base.Ui64(v2261) {
		v2427 = v2235
		v2428 = v2236
		goto L529
	} else {
		goto L535
	}
L535:
	;
	v2263 = *(*int64)(unsafe.Add(mBase, uint32(v2253)+16))
	F_XlogReadTwoPhaseData(m, v2263, v2215+int32(120), v2215+int32(116))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+32))
	v2271 = int32(-1)
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+120))
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+116))
	v2274 = m.Env.Pgmem_crc32c(m, v2271, v2272, v2273)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+124)) = v2274 ^ v2271
	v2278 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	if base.Ui32(v2270) <= base.Ui32(int32(2)) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2297 = base.I64_extend_i32_u(v2270)
	goto L540
L539:
	;
	v2285 = int64(base.Ui64(v2278) >> (uint(int64(32)) % 64))
	if base.Ui32(base.I32_wrap_i64(v2278)) < base.Ui32(v2270) {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v2215)+100)) = uint32(v2297)
	v2300 = int64(base.Ui64(v2297) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2215)+96)) = uint32(v2300)
	v2303 = v2215 + int32(128)
	v2308 = F_pg_snprintf(m, v2303, int32(1024), int32(_a_F_CheckPointGuts_58), v2215+int32(96))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L544
	}
L541:
	;
	v2292 = (v2285 - int64(1)) & int64(4294967295)
	goto L543
L542:
	;
	v2292 = v2285
	goto L543
L543:
	;
	v2297 = base.I64_extend_i32_u(v2270) | v2292<<(uint(int64(32))%64)
	goto L540
L544:
	;
	v2311 = F_OpenTransientFile(m, v2303, int32(577))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	if v2311 < int32(0) {
		goto L520
	} else {
		goto L546
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(0)
	v2319 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2319))) = int32(167772224)
	v2322 = F_write(m, v2311, v2272, v2273)
	mBase = m.M
	if v2322 != v2273 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2325 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L548:
	;
	goto L549
L549:
	;
	v2352 = int32(4)
	v2353 = F_write(m, v2311, v2215+int32(124), v2352)
	mBase = m.M
	if v2353 != v2352 {
		goto L557
	} else {
		goto L558
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L552
L551:
	;
	goto L552
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+80)) = v2215 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_59), v2215+int32(80))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1756), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L557:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2357 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L558:
	;
	goto L559
L559:
	;
	v2382 = int32(_a_F_CheckPointGuts_62)
	v2383 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	v2384 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2383))) = v2384
	v2387 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2387))) = int32(167772223)
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[5])))
	if v2392 != int32(1) {
		v2406 = v2384
		goto L568
	} else {
		goto L569
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6])) = int32(51)
	goto L562
L561:
	;
	goto L562
L562:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+64)) = v2215 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_59), v2215-int32(-64))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1765), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L567:
	;
	if v2406 != 0 {
		goto L519
	} else {
		goto L574
	}
L568:
	;
	goto L567
L569:
	;
	goto L570
L570:
	;
	v2397 = F_fsync(m, v2311)
	mBase = m.M
	if v2397 != int32(-1) {
		v2406 = v2397
		goto L568
	} else {
		goto L572
	}
L571:
	;
	v2406 = int32(-1)
	goto L568
L572:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[6]))
	if v2401 == int32(27) {
		goto L570
	} else {
		goto L573
	}
L573:
	;
	goto L571
L574:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v2408))) = int32(0)
	v2411 = F_CloseTransientFile(m, v2311)
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	if v2411 != 0 {
		goto L518
	} else {
		goto L576
	}
L576:
	;
	v2413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2253)+45)) = uint8(v2413)
	v2416 = v2253 + int32(16)
	v2417 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2416)+8)) = v2417
	*(*int64)(unsafe.Add(mBase, uint32(v2416))) = v2417
	F_pfree(m, v2272)
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPointGuts[45]))
	v2427 = v2426
	v2428 = v2236 + int32(1)
	goto L529
L578:
	;
	goto L528
L579:
	;
	F_fsync_fname(m, int32(_a_F_CheckPointGuts_63), int32(1))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v2465 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckPointGuts[42])))
	if base.B2i32(v2465 != int32(1))|base.B2i32(v2440 <= int32(0)) != 0 {
		goto L521
	} else {
		goto L581
	}
L581:
	;
	v2473 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	if v2473 == int32(0) {
		goto L521
	} else {
		goto L583
	}
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215))) = v2440
	F_errmsg_plural(m, int32(_a_F_CheckPointGuts_64), int32(_a_F_CheckPointGuts_65), v2440, v2215)
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1876), int32(_a_F_CheckPointGuts_66))
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	goto L521
L586:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+16)) = v2215 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_67), v2215+int32(16))
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1744), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L590:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+48)) = v2215 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_20), v2215+int32(48))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1777), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2550 = m.ExcPending
	if v2550 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2215)+32)) = v2215 + int32(128)
	F_errmsg(m, int32(_a_F_CheckPointGuts_23), v2215+int32(32))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_errfinish(m, int32(_a_F_CheckPointGuts_60), int32(1783), int32(_a_F_CheckPointGuts_61))
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_dt(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v64 float64
	_ = v64
	var v71 float64
	_ = v71
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_abs(base.F64_sub(v8, v9))
	if base.F64_eq(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v15 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v8), v15)&base.F64_ne(base.F64_abs(v9), v15) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return float64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = math.Float64frombits(uint64(0x7ff0000000000000))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(base.F64_abs(base.F64_sub(v22, v23)), v21) != 0 {
				v36 = v22
				v37 = v23
				v38 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v36), v38)|base.F64_eq(base.F64_abs(v37), v38) != 0 {
					v71 = v38
					return v71
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v71 = v21
				return v71
			}
		}
	} else {
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		v31 = base.F64_abs(base.F64_sub(v28, v29))
		if base.F64_ne(v31, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v48 = int64(9218868437227405312)
			if base.B2i32(base.Ui64(v48) < base.Ui64(base.I64_reinterpret_f64(v11)))|base.B2i32(base.Ui64(v48) < base.Ui64(base.I64_reinterpret_f64(v31))) != 0 {
				v71 = math.Float64frombits(uint64(0x7ff8000000000000))
				return v71
			} else {
				v54 = base.F64_gt(v31, v11)
				if v54 != 0 {
					v55 = v31
				} else {
					v55 = v11
				}
				if v54 != 0 {
					v56 = v11
				} else {
					v56 = v31
				}
				if base.F64_eq(v56, float64(0)) != 0 {
					v71 = v55
					return v71
				} else {
					v59 = base.F64_div(v56, v55)
					v64 = base.F64_mul(v55, base.F64_sqrt(base.F64_add(base.F64_mul(v59, v59), float64(1))))
					if base.F64_eq(base.F64_abs(v64), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_eq(v64, float64(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v71 = v64
							return v71
						}
					}
				}
			}
		} else {
			v36 = v28
			v37 = v29
			v38 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v36), v38)|base.F64_eq(base.F64_abs(v37), v38) != 0 {
				v71 = v38
				return v71
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_point_inside(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v20 float64
	_ = v20
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v48 int32
	_ = v48
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v147 int32
	_ = v147
	var v166 int32
	_ = v166
	v12 = int32(0)
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v18 = base.F64_sub(v16, v17)
	v20 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v18), v20)|base.F64_eq(base.F64_abs(v16), v20) == v12)&base.F64_ne(base.F64_abs(v17), v20) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L19
	}
L2:
	;
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = base.F64_sub(v32, v33)
	v36 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36) == int32(0))&base.F64_ne(base.F64_abs(v33), v36) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = int32(2)
	if l1 < v48 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v147
L5:
	;
	v127 = F_lseg_crossing(m, v18, v34, v118, v117)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L17
	}
L6:
	;
	v117 = v34
	v118 = v18
	v126 = v12
	goto L5
L7:
	;
	goto L8
L8:
	;
	v55 = v18
	v56 = v34
	v63 = int32(1)
	v66 = v12
	goto L9
L9:
	;
	v69 = l2 + v63<<(uint(int32(4))%32)
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v72 = base.F64_sub(v70, v71)
	v74 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v72), v74)|base.F64_eq(base.F64_abs(v70), v74) == int32(0))&base.F64_ne(base.F64_abs(v71), v74) != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v117 = v88
	v118 = v72
	v126 = v108
	goto L5
L11:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
	v87 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = base.F64_sub(v86, v87)
	v90 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v88), v90)|base.F64_eq(base.F64_abs(v86), v90) == int32(0))&base.F64_ne(base.F64_abs(v87), v90) != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v102 = F_lseg_crossing(m, v72, v88, v55, v56)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v102 == int32(2147483647) {
		v147 = v48
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v108 = v102 + v66
	v110 = v63 + int32(1)
	if v110 != l1 {
		v55 = v72
		v56 = v88
		v63 = v110
		v66 = v108
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v127 == int32(2147483647) {
		v147 = v48
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v147 = base.B2i32(v126 != int32(0)-v127)
	goto L4
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_point_invsl(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v58 float64
	_ = v58
	var v68 float64
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v3 = float64(0)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v10, v11) != 0 {
		v68 = v3
		return v68
	} else {
		v13 = base.F64_sub(v10, v11)
		v14 = base.F64_abs(v13)
		if base.F64_le(v14, float64(1e-06)) != 0 {
			v68 = v3
			return v68
		} else {
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v18, v19)|base.F64_le(base.F64_abs(base.F64_sub(v18, v19)), float64(1e-06)) != 0 {
				v68 = math.Float64frombits(uint64(0x7ff0000000000000))
				return v68
			} else {
				v26 = math.Float64frombits(uint64(0x7ff0000000000000))
				v27 = base.F64_ne(v14, v26)
				if base.B2i32(v27|base.F64_eq(base.F64_abs(v10), v26) == int32(0))&base.F64_ne(base.F64_abs(v11), v26) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v39 = math.Float64frombits(uint64(0x7ff0000000000000))
					v41 = base.F64_sub(v19, v18)
					v42 = base.F64_abs(v41)
					if base.B2i32(base.F64_eq(base.F64_abs(v18), v39)|base.F64_ne(v42, v39) == int32(0))&base.F64_ne(base.F64_abs(v19), v39) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v41, float64(0)) != 0 {
							F_float_zero_divide_error(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v58 = base.F64_div(v13, v41)
							if v27&base.F64_eq(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_ne(v58, float64(0)) != 0 {
									v68 = v58
									return v68
								} else {
									if base.F64_ne(v42, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v68 = v58
										return v68
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = v8 + int32(16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_appendStringInfoChar(m, v12, int32(40))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
			v22 = F_float8out_internal(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_float8out_internal(m, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
					F_appendStringInfo(m, v12, int32(_a_F_point_out_0), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v22)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v24)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoChar(m, v12, int32(41))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									m.G0 = v8 + int32(32)
									return v38
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_point_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v54 int32
	_ = v54
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				return v11
			}
		}
	}
}
