package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v101 int32
	_ = v101
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
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v387 int32
	_ = v387
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
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
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int64
	_ = v796
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int64
	_ = v807
	var v809 int64
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int64
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1105 int64
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1156 int32
	_ = v1156
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1208 int32
	_ = v1208
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int64
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1383 int32
	_ = v1383
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1485 int64
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1492 float64
	_ = v1492
	var v1493 float64
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1535 float64
	_ = v1535
	var v1536 float64
	_ = v1536
	var v1538 float64
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int64
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int64
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int64
	_ = v1558
	var v1560 int64
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1570 float64
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1580 int64
	_ = v1580
	var v1582 int64
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1589 float64
	_ = v1589
	var v1593 float64
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int64
	_ = v1741
	var v1742 int64
	_ = v1742
	var v1752 int64
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1851 int64
	_ = v1851
	var v1852 int64
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1879 int64
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1912 int32
	_ = v1912
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1951 int32
	_ = v1951
	var v1966 int32
	_ = v1966
	var v1967 int64
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2071 int64
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2079 int64
	_ = v2079
	var v2082 int64
	_ = v2082
	var v2088 int64
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2112 int32
	_ = v2112
	var v2117 int32
	_ = v2117
	var v2118 int64
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2131 int64
	_ = v2131
	var v2132 int64
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2156 int64
	_ = v2156
	var v2157 int64
	_ = v2157
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2175 int32
	_ = v2175
	var v2177 int64
	_ = v2177
	var v2178 int64
	_ = v2178
	var v2190 int32
	_ = v2190
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2220 int32
	_ = v2220
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2230 int64
	_ = v2230
	var v2231 int64
	_ = v2231
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int64
	_ = v2291
	var v2293 int64
	_ = v2293
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2308 int64
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2315 int64
	_ = v2315
	var v2322 int64
	_ = v2322
	var v2327 int64
	_ = v2327
	var v2330 int64
	_ = v2330
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
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
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2408 int32
	_ = v2408
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2424 int32
	_ = v2424
	var v2429 int32
	_ = v2429
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2449 int64
	_ = v2449
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2477 int32
	_ = v2477
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2513 int32
	_ = v2513
	var v2518 int32
	_ = v2518
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2552 int32
	_ = v2552
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2563 int32
	_ = v2563
	var v2571 int32
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2590 int32
	_ = v2590
	var v2595 int32
	_ = v2595
	v3 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[47]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[47]))
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
	v33 = m.G0
	v35 = v33 - int32(1040)
	m.G0 = v35
	v39 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errmsg_internal(m, int32(93859), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v55 = F_LWLockAcquire(m, v51+int32(4608), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_errfinish(m, int32(512119), int32(2124), int32(125871))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if int32(0) < v58 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v35 + int32(1040)
	v177 = m.G0
	v179 = v177 - int32(1152)
	m.G0 = v179
	v181 = F_GetRedoRecPtr(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L37
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	v65 = v3
	v66 = v58
	v69 = v3
	v70 = v62
	goto L15
L13:
	;
	goto L14
L14:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v152+int32(4608))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v82 = v70 + v65*int32(288)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v83 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v140+int32(4608))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L33
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(90359)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v82 + int32(24)
	v94 = F_pg_sprintf(m, v35+int32(16), int32(185927), v35)
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
	v132 = v69
	v133 = v70
	goto L19
L19:
	;
	v137 = v65 + int32(1)
	if v137 < v131 {
		v65 = v137
		v66 = v131
		v69 = v132
		v70 = v133
		goto L15
	} else {
		goto L32
	}
L20:
	;
	if l1&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v82)+104))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v82)+280))
	F_SaveSlotToPath(m, v82, v35+int32(16), int32(15))
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
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(1)
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_s_lock(m, v82, int32(512119), int32(2162), int32(125871))
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
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
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
	v126 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	v131 = v126
	v132 = base.B2i32(v119 != v118) | v69
	v133 = v128
	goto L19
L32:
	;
	goto L16
L33:
	;
	if v132&int32(1) == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
	v183 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v186 = F_AllocateDir(m, int32(125992))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v189 = F_ReadDir(m, v186, int32(125992))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v189 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui64(v181) < base.Ui64(v183) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	F_FreeDir(m, v186)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L86
	}
L44:
	;
	v192 = v181
	goto L46
L45:
	;
	v192 = v183
	goto L46
L46:
	;
	v201 = v189
	goto L47
L47:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+19)))
	if v212 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L43
L49:
	;
	v338 = F_ReadDir(m, v186, int32(125992))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L84
	}
L50:
	;
	v225 = v201 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+84)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v179)+80)) = int32(125992)
	v235 = F_pg_snprintf(m, v179+int32(96), int32(1045), int32(185927), v179+int32(80))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L55
	}
L51:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+20)))
	if v215 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+20)))
	if v218 != int32(46) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+21)))
	if v221 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v242 = F_get_dirent_type(m, v179+int32(96), v201, int32(0), int32(14))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	F_errfinish(m, int32(519946), v332, int32(447946))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L83
	}
L57:
	;
	if v242&int32(-3) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v248 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+52)) = v179 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+48)) = v179 + int32(92)
	v270 = F_sscanf(m, v225, int32(248058), v179+int32(48))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	if v248 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+64)) = v179 + int32(96)
	F_errmsg_internal(m, int32(212717), v179-int32(-64))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v332 = int32(2009)
	goto L56
L64:
	;
	if v270 != int32(2) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v276 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v289 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v179)+88)))
	v290 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v179)+92)))
	if base.Ui64(v192-int64(1)) < base.Ui64(v289|v290<<(uint(int64(32))%64)) {
		goto L49
	} else {
		goto L71
	}
L68:
	;
	if v276 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+32)) = v179 + int32(96)
	F_errmsg(m, int32(742631), v179+int32(32))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v332 = int32(2025)
	goto L56
L71:
	;
	v297 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v297 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = v179 + int32(96)
	F_errmsg_internal(m, int32(187549), v179+int32(16))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v314 = F_unlink(m, v179+int32(96))
	mBase = m.M
	if int32(0) <= v314 {
		goto L49
	} else {
		goto L78
	}
L76:
	;
	F_errfinish(m, int32(519946), int32(2034), int32(447946))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v319 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v319 == int32(0) {
		goto L49
	} else {
		goto L80
	}
L80:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v179 + int32(96)
	F_errmsg(m, int32(310820), v179)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v332 = int32(2046)
	goto L56
L83:
	;
	goto L49
L84:
	;
	if v338 != 0 {
		v201 = v338
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L48
L86:
	;
	m.G0 = v179 + int32(1152)
	v362 = m.G0
	v364 = v362 - int32(1216)
	m.G0 = v364
	v366 = F_GetRedoRecPtr(m)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v368 = F_ReplicationSlotsComputeLogicalRestartLSN(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v371 = F_AllocateDir(m, int32(164463))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L93
	}
L89:
	;
	v698 = m.G0
	v700 = v698 - int32(112)
	m.G0 = v700
	*(*int32)(unsafe.Add(mBase, uint32(v700)+108)) = int32(307747550)
	v705 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v705 != 0 {
		goto L186
	} else {
		goto L187
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L178
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L174
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L171
	}
L93:
	;
	v374 = F_ReadDir(m, v371, int32(164463))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v374 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if base.Ui64(v366) < base.Ui64(v368) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	F_FreeDir(m, v371)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L169
	}
L98:
	;
	v377 = v366
	goto L100
L99:
	;
	v377 = v368
	goto L100
L100:
	;
	v387 = v374
	goto L101
L101:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+19)))
	if v401 != int32(46) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L97
L103:
	;
	v618 = F_ReadDir(m, v371, int32(164463))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L167
	}
L104:
	;
	v414 = v387 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+132)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v364)+128)) = int32(164463)
	v424 = F_pg_snprintf(m, v364+int32(160), int32(1044), int32(185927), v364+int32(128))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+20)))
	if v404 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+20)))
	if v407 != int32(46) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+21)))
	if v410 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v430 = F_get_dirent_type(m, v364+int32(160), v387, int32(0), int32(14))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v430&int32(-3) != 0 {
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v434 = int32(694409)
	goto L114
L112:
	;
	if v471-v472 != 0 {
		goto L103
	} else {
		goto L126
	}
L114:
	;
	goto L115
L115:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v441 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v442 = v414
	v443 = v434
	v444 = int32(4)
	v445 = v441
	goto L120
L117:
	;
	v467 = v434
	v471 = int32(0)
	goto L118
L118:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	goto L112
L119:
	;
	v467 = v462
	v471 = v464
	goto L118
L120:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v445 != v447 {
		v462 = v443
		v464 = v445
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v462 = v456
	v464 = int32(0)
	goto L119
L122:
	;
	if v447 == int32(0) {
		v462 = v443
		v464 = v445
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v452 = v444 - int32(1)
	if v452 == int32(0) {
		v462 = v443
		v464 = v445
		goto L119
	} else {
		goto L124
	}
L124:
	;
	v455 = int32(1)
	v456 = v443 + v455
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	if v457 != 0 {
		v442 = v442 + v455
		v443 = v456
		v444 = v452
		v445 = v457
		goto L120
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364+int32(116)))) = v364 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v364+int32(112)))) = v364 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+108)) = v364 + int32(136)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+104)) = v364 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+100)) = v364 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+96)) = v364 + int32(156)
	v501 = F_sscanf(m, v414, int32(30856), v364+int32(96))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v501 != int32(6) {
		goto L92
	} else {
		goto L128
	}
L128:
	;
	if v368 != int64(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v557 = F_OpenTransientFile(m, v364+int32(160), int32(2))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L145
	}
L130:
	;
	v507 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v364)+136)))
	v508 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v364)+140)))
	if base.Ui64(v377-int64(1)) < base.Ui64(v507|v508<<(uint(int64(32))%64)) {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v515 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	if v515 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+64)) = v364 + int32(160)
	F_errmsg_internal(m, int32(744942), v364-int32(-64))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v532 = F_unlink(m, v364+int32(160))
	mBase = m.M
	if int32(0) <= v532 {
		goto L103
	} else {
		goto L140
	}
L138:
	;
	F_errfinish(m, int32(515656), int32(1210), int32(248852))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+48)) = v364 + int32(160)
	F_errmsg(m, int32(310820), v364+int32(48))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(515656), int32(1214), int32(248852))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
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
	if v557 < int32(0) {
		goto L91
	} else {
		goto L146
	}
L146:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = int32(167772194)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v567 != int32(1) {
		v581 = int32(0)
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = int32(0)
	v613 = F_CloseTransientFile(m, v557)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L165
	}
L148:
	;
	if v581 == int32(0) {
		goto L147
	} else {
		goto L155
	}
L149:
	;
	goto L148
L150:
	;
	goto L151
L151:
	;
	v572 = F_fsync(m, v557)
	mBase = m.M
	if v572 != int32(-1) {
		v581 = v572
		goto L149
	} else {
		goto L153
	}
L152:
	;
	v581 = int32(-1)
	goto L149
L153:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v576 == int32(27) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, _consts[45])))
	if v587 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v590 = F_errstart(m, v588, int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L160
	}
L157:
	;
	v588 = int32(21)
	goto L159
L158:
	;
	v588 = int32(23)
	goto L159
L159:
	;
	goto L156
L160:
	;
	if v590 == int32(0) {
		goto L147
	} else {
		goto L161
	}
L161:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+32)) = v364 + int32(160)
	F_errmsg(m, int32(311214), v364+int32(32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(515656), int32(1240), int32(248852))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L147
L165:
	;
	if v613 != 0 {
		goto L90
	} else {
		goto L166
	}
L166:
	;
	goto L103
L167:
	;
	if v618 != 0 {
		v387 = v618
		goto L101
	} else {
		goto L168
	}
L168:
	;
	goto L102
L169:
	;
	F_fsync_fname(m, int32(164463), int32(1))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	m.G0 = v364 + int32(1216)
	goto L89
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+80)) = v414
	F_errmsg_internal(m, int32(742315), v364+int32(80))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(515656), int32(1204), int32(248852))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v364 + int32(160)
	F_errmsg(m, int32(310266), v364)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(515656), int32(1229), int32(248852))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v364 + int32(160)
	F_errmsg(m, int32(311009), v364+int32(16))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(515656), int32(1246), int32(248852))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	v1006 = m.G0
	v1007 = int32(16)
	v1008 = v1006 - v1007
	m.G0 = v1008
	F___gettimeofday(m, v1008)
	mBase = m.M
	v1011 = *(*int64)(unsafe.Add(mBase, uint32(v1008)))
	v1012 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1008)+8)))
	m.G0 = v1008 + v1007
	goto L254
L183:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L250
	}
L184:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L246
	}
L185:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L242
	}
L186:
	;
	v707 = F_unlink(m, int32(245290))
	mBase = m.M
	if v707 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L188
L188:
	;
	m.G0 = v700 + int32(112)
	goto L182
L189:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v711 != int32(44) {
		goto L185
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v716 = F_OpenTransientFile(m, int32(245290), int32(193))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L193
	}
L192:
	;
	goto L191
L193:
	;
	if v716 < int32(0) {
		goto L184
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v725 = int32(4)
	v726 = F_write(m, v716, v700+int32(108), v725)
	mBase = m.M
	if v726 != v725 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v730 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v730 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	v758 = m.Env.Pgmem_crc32c(m, int32(-1), v700+int32(108), int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v700)+104)) = v758
	v761 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v765 = F_LWLockAcquire(m, v761+int32(5120), int32(1))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L205
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L200
L199:
	;
	goto L200
L200:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+64)) = int32(245290)
	F_errmsg(m, int32(309949), v700-int32(-64))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(516209), int32(639), int32(287663))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if int32(0) < v768 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v774 = v700 + int32(96)
	v777 = v758
	v779 = v768
	v781 = int32(0)
	v782 = v772
	goto L209
L207:
	;
	v865 = v758
	goto L208
L208:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v881+int32(5120))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L228
	}
L209:
	;
	v794 = v782 + v781*int32(56)
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v794))))
	if v795 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v865 = v856
	goto L208
L211:
	;
	v796 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v774))) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v700)+88)) = v796
	v801 = v794 + int32(40)
	v803 = F_LWLockAcquire(m, v801, int32(1))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	v856 = v777
	v857 = v779
	v858 = v782
	goto L213
L213:
	;
	v861 = v781 + int32(1)
	if v861 < v857 {
		v777 = v856
		v779 = v857
		v781 = v861
		v782 = v858
		goto L209
	} else {
		goto L227
	}
L214:
	;
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v794))))
	*(*uint16)(unsafe.Add(mBase, uint32(v700)+88)) = uint16(v805)
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v794)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v774))) = v807
	v809 = *(*int64)(unsafe.Add(mBase, uint32(v794)+16))
	F_LWLockRelease(m, v801)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_XLogFlush(m, v809)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v819 = int32(16)
	v820 = F_write(m, v716, v700+int32(88), v819)
	mBase = m.M
	if v820 != v819 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v824 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	v851 = m.Env.Pgmem_crc32c(m, v777, v700+int32(88), int32(16))
	mBase = m.M
	v853 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v855 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v856 = v851
	v857 = v853
	v858 = v855
	goto L213
L220:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L222
L221:
	;
	goto L222
L222:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+48)) = int32(245290)
	F_errmsg(m, int32(309949), v700+int32(48))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(516209), int32(681), int32(287663))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	goto L210
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+104)) = v865 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v894 = int32(4)
	v895 = F_write(m, v716, v700+int32(104), v894)
	mBase = m.M
	if v895 != v894 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v899 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	v923 = F_CloseTransientFile(m, v716)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L239
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L234
L233:
	;
	goto L234
L234:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+32)) = int32(245290)
	F_errmsg(m, int32(309949), v700+int32(32))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(516209), int32(700), int32(287663))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
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
	if v923 != 0 {
		goto L183
	} else {
		goto L240
	}
L240:
	;
	v928 = F_durable_rename(m, int32(245290), int32(93790), int32(23))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	goto L188
L242:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+80)) = int32(245290)
	F_errmsg(m, int32(310820), v700+int32(80))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(516209), int32(615), int32(287663))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = int32(245290)
	F_errmsg(m, int32(310945), v700)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(516209), int32(627), int32(287663))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+16)) = int32(245290)
	F_errmsg(m, int32(311009), v700+int32(16))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(516209), int32(707), int32(287663))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	*(*int64)(unsafe.Add(mBase, _consts[333])) = v1012 + v1011*int64(1000000) - int64(946684800000000)
	F_SimpleLruWriteAll(m, int32(4443296))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_SimpleLruWriteAll(m, int32(4443380))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_SimpleLruWriteAll(m, int32(4443672))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_SimpleLruWriteAll(m, int32(4443476))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_SimpleLruWriteAll(m, int32(4443556))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v1042 = F_LWLockAcquire(m, v1038+int32(6656), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(v1045)))
	if v1046 < int64(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1122 = int32(0)
	v1125 = m.G0
	v1127 = v1125 - int32(5136)
	m.G0 = v1127
	v1130 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	if v1130 <= v1122 {
		goto L284
	} else {
		goto L285
	}
L262:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v1050+int32(6656))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+12))
	if v1055 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L261
L266:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v1107+int32(6656))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L281
	}
L267:
	;
	v1058 = int32(4)
	v1059 = v1055&int32(-1024) | v1058
	v1062 = base.I32_wrap_i64(v1046) << (uint(int32(10)) % 32)
	v1064 = v1062 | v1058
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1064))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1059)) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L268:
	;
	goto L269
L269:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1045))) = int64(-1)
	v1105 = v1046
	goto L266
L270:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1099)))
	v1105 = v1100
	goto L266
L271:
	;
	if v1076 == int32(0) {
		goto L270
	} else {
		goto L275
	}
L272:
	;
	v1076 = base.B2i32(base.Ui32(v1059) < base.Ui32(v1064))
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1076 = int32(base.Ui32(v1059-v1064) >> (uint(int32(31)) % 32))
	goto L271
L275:
	;
	v1080 = v1062 + int32(1027)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1080))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1059)) == int32(0) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if v1092 == int32(0) {
		goto L270
	} else {
		goto L280
	}
L277:
	;
	v1092 = base.B2i32(base.Ui32(v1059) < base.Ui32(v1080))
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1092 = int32(base.Ui32(v1059-v1080) >> (uint(int32(31)) % 32))
	goto L276
L280:
	;
	v1105 = base.I64_extend_i32_u(int32(base.Ui32(v1055) >> (uint(int32(10)) % 32)))
	goto L266
L281:
	;
	F_SimpleLruTruncate(m, int32(4472036), v1105)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_SimpleLruWriteAll(m, int32(4472036))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L261
L284:
	;
	m.G0 = v1127 + int32(5136)
	v1736 = m.G0
	v1737 = int32(16)
	v1738 = v1736 - v1737
	m.G0 = v1738
	F___gettimeofday(m, v1738)
	mBase = m.M
	v1741 = *(*int64)(unsafe.Add(mBase, uint32(v1738)))
	v1742 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1738)+8)))
	m.G0 = v1738 + v1737
	goto L416
L285:
	;
	if l1&int32(19) != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1137 = int32(-8388609)
	goto L288
L287:
	;
	v1137 = int32(2139095039)
	goto L288
L288:
	;
	v1142 = v1122
	v1146 = v1122
	goto L289
L289:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+28)) = int32(239002)
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+20)) = int32(514763)
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1127)+8)) = int64(0)
	v1169 = v1156 + v1142<<(uint(int32(6))%32)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+24))
	v1171 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+24)) = v1170 | v1171
	if v1170&v1171 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	if v1265 == int32(0) {
		goto L284
	} else {
		goto L317
	}
L291:
	;
	goto L294
L292:
	;
	v1208 = v1170
	goto L293
L293:
	;
	v1223 = int32(4155052)
	v1224 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1127+int32(8))+8))
	if v1226 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L294:
	;
	F_perform_spin_delay(m, v1127+int32(8))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L296
	}
L295:
	;
	v1208 = v1197
	goto L293
L296:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+24))
	v1198 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+24)) = v1197 | v1198
	if v1197&v1198 != 0 {
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	if v1137|v1208 == int32(-1) {
		goto L309
	} else {
		goto L310
	}
L299:
	;
	goto L298
L300:
	;
	*(*int32)(unsafe.Add(mBase, _consts[335])) = v1241
	goto L299
L301:
	;
	if int32(999) < v1224 {
		goto L299
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	if v1224 < int32(11) {
		goto L299
	} else {
		goto L308
	}
L304:
	;
	v1231 = int32(900)
	if v1231 <= v1224 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1234 = v1231
	goto L307
L306:
	;
	v1234 = v1224
	goto L307
L307:
	;
	v1241 = v1234 + int32(100)
	goto L300
L308:
	;
	v1241 = v1224 - int32(1)
	goto L300
L309:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	v1250 = v1247 + v1146*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+16)) = v1142
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1252
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+4)) = v1254
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+8)) = v1256
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+12)) = v1258
	v1265 = v1146 + int32(1)
	v1266 = v1208 | int32(1077936128)
	goto L311
L310:
	;
	v1265 = v1146
	v1266 = v1208
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+24)) = v1266 & int32(-4194305)
	v1271 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	if v1271 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1275 = v1142 + int32(1)
	v1277 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	if v1275 < v1277 {
		v1142 = v1275
		v1146 = v1265
		goto L289
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	goto L290
L317:
	;
	v1281 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+12)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+8)) = int32(4464552)
	v1287 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	F_sort_checkpoint_bufferids(m, v1287, v1265)
	mBase = m.M
	if v1281 < v1265 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	F_binaryheap_build(m, v1423)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L349
	}
L319:
	;
	v1291 = int32(0)
	v1295 = v1281
	v1297 = v1291
	v1300 = v1122
	v1303 = v1291
	goto L322
L320:
	;
	goto L321
L321:
	;
	v1407 = int32(0)
	v1411 = F_binaryheap_allocate(m, v1407, int32(1077), v1407)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L348
	}
L322:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1312+v1303*int32(20))))
	if v1300 == v1316 {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	v1370 = int32(0)
	v1373 = F_binaryheap_allocate(m, v1354, int32(1077), v1370)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L342
	}
L324:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1356)+24)) = v1359 + int32(1)
	v1364 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	if v1364 != 0 {
		goto L337
	} else {
		goto L338
	}
L325:
	;
	v1318 = v1300
	goto L327
L326:
	;
	v1318 = int32(0)
	goto L327
L327:
	;
	if v1318 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1322 = v1295 + int32(1)
	v1324 = v1322 * int32(40)
	if v1297 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L329:
	;
	goto L330
L330:
	;
	v1349 = int32(40)
	v1354 = v1295
	v1355 = v1297
	v1356 = v1297 + v1295*v1349 - v1349
	v1357 = v1300
	goto L324
L331:
	;
	v1334 = v1331 + v1295*int32(40)
	v1335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1334))) = v1335
	v1338 = v1334 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1338))) = v1335
	*(*int64)(unsafe.Add(mBase, uint32(v1334)+24)) = v1335
	*(*int64)(unsafe.Add(mBase, uint32(v1334)+16)) = v1335
	*(*int64)(unsafe.Add(mBase, uint32(v1334)+8)) = v1335
	*(*int32)(unsafe.Add(mBase, uint32(v1334))) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v1338))) = v1303
	v1354 = v1322
	v1355 = v1331
	v1356 = v1334
	v1357 = v1316
	goto L324
L332:
	;
	v1327 = F_palloc(m, v1324)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v1329 = F_repalloc(m, v1297, v1324)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L1
	} else {
		goto L336
	}
L335:
	;
	v1331 = v1327
	goto L331
L336:
	;
	v1331 = v1329
	goto L331
L337:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1368 = v1303 + int32(1)
	if v1368 != v1265 {
		v1295 = v1354
		v1297 = v1355
		v1300 = v1357
		v1303 = v1368
		goto L322
	} else {
		goto L341
	}
L340:
	;
	goto L339
L341:
	;
	goto L323
L342:
	;
	if v1354 <= int32(0) {
		v1417 = v1355
		v1423 = v1373
		goto L318
	} else {
		goto L343
	}
L343:
	;
	v1383 = v1370
	goto L344
L344:
	;
	v1397 = v1355 + v1383*int32(40)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v1397)+16)) = base.F64_div(base.F64_convert_i32_s(v1265), base.F64_convert_i32_s(v1398))
	F_binaryheap_add_unordered(m, v1373, v1397)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L346
	}
L345:
	;
	v1417 = v1355
	v1423 = v1373
	goto L318
L346:
	;
	v1405 = v1383 + int32(1)
	if v1405 != v1354 {
		v1383 = v1405
		goto L344
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v1417 = v1407
	v1423 = v1411
	goto L318
L349:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	if v1432 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	F_IssuePendingWritebacks(m, v1127+int32(8), int32(3))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L413
	}
L351:
	;
	v1683 = int32(0)
	goto L350
L352:
	;
	goto L353
L353:
	;
	v1437 = int32(0)
	v1441 = v1437
	v1445 = v1437
	goto L354
L354:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1459 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+20))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+32))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1459+v1461*int32(20))+16))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1457+v1465<<(uint(int32(6))%32))+24))
	if v1469&int32(1073741824) == int32(0) {
		v1491 = v1441
		goto L356
	} else {
		goto L357
	}
L355:
	;
	v1683 = v1491
	goto L350
L356:
	;
	v1492 = *(*float64)(unsafe.Add(mBase, uint32(v1460)+16))
	v1493 = *(*float64)(unsafe.Add(mBase, uint32(v1460)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v1460)+8)) = base.F64_add(v1492, v1493)
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+28))
	v1497 = int32(1)
	v1498 = v1496 + v1497
	*(*int32)(unsafe.Add(mBase, uint32(v1460)+28)) = v1498
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1460)+32)) = v1500 + v1497
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+24))
	if v1504 == v1498 {
		goto L361
	} else {
		goto L362
	}
L357:
	;
	v1477 = F_SyncOneBuffer(m, v1465, int32(0), v1127+int32(8))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	if v1477&int32(1) == int32(0) {
		v1491 = v1441
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v1483 = int32(4528464)
	v1485 = *(*int64)(unsafe.Add(mBase, _consts[338]))
	*(*int64)(unsafe.Add(mBase, _consts[338])) = v1485 + int64(1)
	v1491 = v1441 + int32(1)
	goto L356
L360:
	;
	v1511 = v1445 + int32(1)
	v1515 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if v1515 != int32(11) {
		goto L366
	} else {
		goto L367
	}
L361:
	;
	v1506 = F_binaryheap_remove_first(m, v1423)
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	F_binaryheap_replace_first(m, v1423, v1460)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	goto L360
L365:
	;
	goto L360
L366:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	if v1680 != 0 {
		v1441 = v1491
		v1445 = v1511
		goto L354
	} else {
		goto L412
	}
L367:
	;
	if l1&int32(4) != 0 {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	if v1671 == int32(0) {
		goto L366
	} else {
		goto L410
	}
L369:
	;
	v1653 = int32(4154544)
	v1655 = *(*int32)(unsafe.Add(mBase, _consts[339]))
	v1657 = v1655 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[339])) = v1657
	if int32(0) < v1657 {
		goto L368
	} else {
		goto L408
	}
L370:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, _consts[340]))
	if v1521 != 0 {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, _consts[341]))
	if v1523 != 0 {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _consts[342]))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+20))
	if v1526&int32(4) != 0 {
		goto L369
	} else {
		goto L373
	}
L373:
	;
	v1529 = m.G0
	v1531 = v1529 - int32(16)
	m.G0 = v1531
	v1535 = *(*float64)(unsafe.Add(mBase, _consts[343]))
	v1536 = base.F64_mul(base.F64_div(base.F64_convert_i32_s(v1511), base.F64_convert_i32_s(v1265)), v1535)
	v1538 = *(*float64)(unsafe.Add(mBase, _consts[344]))
	if base.F64_lt(v1536, v1538) != 0 {
		v1598 = int32(0)
		goto L374
	} else {
		goto L375
	}
L374:
	;
	m.G0 = v1531 + int32(16)
	if v1598 == int32(0) {
		goto L369
	} else {
		goto L390
	}
L375:
	;
	v1542 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v1542 == int32(1) {
		goto L378
	} else {
		goto L379
	}
L376:
	;
	v1560 = *(*int64)(unsafe.Add(mBase, _consts[345]))
	v1564 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	v1568 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v1570 = base.F64_div(base.F64_div(base.F64_convert_i64_u(v1558-v1560), base.F64_convert_i32_s(v1564)), base.F64_convert_i32_s(v1568))
	if base.F64_lt(v1536, v1570) == int32(0) {
		goto L386
	} else {
		goto L387
	}
L377:
	;
	if v1552 != 0 {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+316))
	v1550 = base.B2i32(v1548 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v1550)
	v1552 = v1550
	goto L380
L379:
	;
	v1552 = int32(0)
	goto L380
L380:
	;
	goto L377
L381:
	;
	v1554 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L1
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	v1556 = F_GetInsertRecPtr(m)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L385
	}
L384:
	;
	v1558 = v1554
	goto L376
L385:
	;
	v1558 = v1556
	goto L376
L386:
	;
	F___gettimeofday(m, v1531)
	mBase = m.M
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+8))
	v1580 = *(*int64)(unsafe.Add(mBase, uint32(v1531)))
	v1582 = *(*int64)(unsafe.Add(mBase, _consts[346]))
	v1587 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v1589 = base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v1576), float64(1e+06)), base.F64_convert_i64_s(v1580-v1582)), base.F64_convert_i32_s(v1587))
	if base.F64_lt(v1536, v1589) == int32(0) {
		v1598 = int32(1)
		goto L374
	} else {
		goto L389
	}
L387:
	;
	v1593 = v1570
	goto L388
L388:
	;
	*(*float64)(unsafe.Add(mBase, _consts[344])) = v1593
	v1598 = int32(0)
	goto L374
L389:
	;
	v1593 = v1589
	goto L388
L390:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, _consts[348]))
	if v1605 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, _consts[348])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L1
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L403
	}
L394:
	;
	F_SyncRepUpdateSyncStandbysDefined(m)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1618 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	if v1618 != 0 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	F_errmsg_internal(m, int32(166299), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	goto L393
L401:
	;
	F_errfinish(m, int32(514853), int32(1390), int32(351275))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	*(*int32)(unsafe.Add(mBase, _consts[339])) = int32(1000)
	F_CheckArchiveTimeout(m)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	F_pgstat_report_checkpointer(m)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	v1643 = F_WaitLatch(m, v1639, int32(41), int32(100), int32(150994945))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	*(*int32)(unsafe.Add(mBase, uint32(v1646))) = int32(0)
	goto L407
L407:
	;
	goto L368
L408:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, _consts[339])) = int32(1000)
	goto L368
L410:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	goto L366
L412:
	;
	goto L355
L413:
	;
	F_pfree(m, v1417)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	F_pfree(m, v1423)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v1707 = int32(4444480)
	v1709 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v1709 + v1683
	goto L284
L416:
	;
	*(*int64)(unsafe.Add(mBase, _consts[350])) = v1742 + v1741*int64(1000000) - int64(946684800000000)
	v1752 = int64(0)
	v1753 = int32(0)
	v1755 = m.G0
	v1757 = v1755 - int32(1152)
	m.G0 = v1757
	v1760 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	if v1760 != 0 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v2225 = m.G0
	v2226 = int32(16)
	v2227 = v2225 - v2226
	m.G0 = v2227
	F___gettimeofday(m, v2227)
	mBase = m.M
	v2230 = *(*int64)(unsafe.Add(mBase, uint32(v2227)))
	v2231 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2227)+8)))
	m.G0 = v2227 + v2226
	goto L519
L418:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L1
	} else {
		goto L516
	}
L421:
	;
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, _consts[352])))
	if v1764 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1821 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[352])) = uint8(v1821)
	v1823 = int32(4472272)
	v1825 = int32(*(*uint16)(unsafe.Add(mBase, _consts[353])))
	v1827 = v1825 + v1821
	*(*uint16)(unsafe.Add(mBase, _consts[353])) = uint16(v1827)
	v1832 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	F_hash_seq_init(m, v1757+int32(1116), v1832)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L431
	}
L423:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	F_hash_seq_init(m, v1757+int32(1116), v1770)
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v1775 = F_hash_seq_search(m, v1757+int32(1116))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	if v1775 == int32(0) {
		goto L422
	} else {
		goto L426
	}
L426:
	;
	v1781 = v1775
	goto L427
L427:
	;
	v1797 = int32(*(*uint16)(unsafe.Add(mBase, _consts[353])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1781)+24)) = uint16(v1797)
	v1801 = F_hash_seq_search(m, v1757+int32(1116))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L429
	}
L428:
	;
	goto L422
L429:
	;
	if v1801 != 0 {
		v1781 = v1801
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	v1837 = F_hash_seq_search(m, v1757+int32(1116))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L1
	} else {
		goto L433
	}
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L1
	} else {
		goto L513
	}
L433:
	;
	if v1837 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1842 = v1837
	v1846 = int32(10)
	v1849 = v1753
	v1851 = v1752
	v1852 = v1752
	goto L437
L435:
	;
	v2175 = v1753
	v2177 = v1752
	v2178 = v1752
	goto L436
L436:
	;
	*(*int64)(unsafe.Add(mBase, _consts[354])) = v2177
	*(*int64)(unsafe.Add(mBase, _consts[355])) = v2178
	*(*int32)(unsafe.Add(mBase, _consts[356])) = v2175
	v2190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[352])) = uint8(v2190)
	m.G0 = v1757 + int32(1152)
	goto L417
L437:
	;
	v1857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1842)+24)))
	v1859 = int32(*(*uint16)(unsafe.Add(mBase, _consts[353])))
	if v1857 != v1859 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	v2175 = v2154
	v2177 = v2156
	v2178 = v2157
	goto L436
L439:
	;
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v1862 != int32(1) {
		v2126 = v1846
		v2129 = v1849
		v2131 = v1851
		v2132 = v1852
		goto L442
	} else {
		goto L443
	}
L440:
	;
	v2151 = v1846
	v2154 = v1849
	v2156 = v1851
	v2157 = v1852
	goto L441
L441:
	;
	v2164 = F_hash_seq_search(m, v1757+int32(1116))
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L1
	} else {
		goto L511
	}
L442:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	v2141 = F_hash_search(m, v2138, v1842, int32(2), int32(0))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L509
	}
L443:
	;
	v1866 = v1846 - int32(1)
	if v1866 <= int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L1
	} else {
		goto L447
	}
L445:
	;
	v1872 = v1866
	goto L446
L446:
	;
	v1873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842)+26)))
	if v1873 != 0 {
		v2126 = v1872
		v2129 = v1849
		v2131 = v1851
		v2132 = v1852
		goto L442
	} else {
		goto L448
	}
L447:
	;
	v1872 = int32(10)
	goto L446
L448:
	;
	F___clock_gettime(m, int32(1), v1757+int32(1136))
	mBase = m.M
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+1144))
	v1879 = *(*int64)(unsafe.Add(mBase, uint32(v1757)+1136))
	v1882 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1842))))
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1882*int32(12))+uint32(_consts[357])))
	v1888 = m.T0[v1887].(func(*base.Module, int32, int32) int32)(m, v1842, v1757+int32(80))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L450
	}
L449:
	;
	v2075 = int32(1)
	F___clock_gettime(m, v2075, v1757+int32(1136))
	mBase = m.M
	v2079 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1757)+1144)))
	v2082 = *(*int64)(unsafe.Add(mBase, uint32(v1757)+1136))
	v2088 = base.I64_div_s(v2079-base.I64_extend_i32_s(v2065)+(v2082-v2071)*int64(1000000000), int64(1000))
	v2091 = v1849 + v2075
	v2093 = int32(*(*uint8)(unsafe.Add(mBase, _consts[358])))
	if v2093 != v2075 {
		goto L500
	} else {
		goto L501
	}
L450:
	;
	if v1888 == int32(0) {
		v2064 = v1872
		v2065 = v1878
		v2071 = v1879
		goto L449
	} else {
		goto L451
	}
L451:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v1894 == int32(44) {
		goto L454
	} else {
		goto L455
	}
L452:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L1
	} else {
		goto L470
	}
L453:
	;
	F_errfinish(m, int32(520367), v1935, int32(124781))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L469
	}
L454:
	;
	v1899 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L1
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	v1917 = int32(*(*uint8)(unsafe.Add(mBase, _consts[45])))
	if v1917 != 0 {
		goto L462
	} else {
		goto L463
	}
L457:
	;
	if v1899 == int32(0) {
		goto L452
	} else {
		goto L458
	}
L458:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+48)) = v1757 + int32(80)
	F_errmsg_internal(m, int32(305784), v1757+int32(48))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	v1935 = int32(452)
	goto L453
L461:
	;
	v1920 = F_errstart(m, v1918, int32(0))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L465
	}
L462:
	;
	v1918 = int32(21)
	goto L464
L463:
	;
	v1918 = int32(23)
	goto L464
L464:
	;
	goto L461
L465:
	;
	if v1920 == int32(0) {
		goto L452
	} else {
		goto L466
	}
L466:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+64)) = v1757 + int32(80)
	F_errmsg(m, int32(311214), v1757-int32(-64))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v1935 = int32(447)
	goto L453
L469:
	;
	goto L452
L470:
	;
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842)+26)))
	if v1942 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v1951 = int32(1)
	goto L474
L472:
	;
	goto L473
L473:
	;
	v2126 = int32(10)
	v2129 = v1849
	v2131 = v1851
	v2132 = v1852
	goto L442
L474:
	;
	F___clock_gettime(m, int32(1), v1757+int32(1136))
	mBase = m.M
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+1144))
	v1967 = *(*int64)(unsafe.Add(mBase, uint32(v1757)+1136))
	v1970 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1842))))
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1970*int32(12))+uint32(_consts[357])))
	v1976 = m.T0[v1975].(func(*base.Module, int32, int32) int32)(m, v1842, v1757+int32(80))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L1
	} else {
		goto L476
	}
L475:
	;
	goto L473
L476:
	;
	if v1976 == int32(0) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v2064 = int32(10)
	v2065 = v1966
	v2071 = v1967
	goto L449
L478:
	;
	goto L479
L479:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v1986 = int32(0)
	if base.B2i32(v1983 == int32(44))&base.B2i32(v1951 <= v1986) == v1986 {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L498
	}
L481:
	;
	F_errfinish(m, int32(520367), v2029, int32(124781))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L497
	}
L482:
	;
	v1994 = int32(*(*uint8)(unsafe.Add(mBase, _consts[45])))
	if v1994 != 0 {
		goto L486
	} else {
		goto L487
	}
L483:
	;
	goto L484
L484:
	;
	v2014 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L493
	}
L485:
	;
	v1997 = F_errstart(m, v1995, int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L489
	}
L486:
	;
	v1995 = int32(21)
	goto L488
L487:
	;
	v1995 = int32(23)
	goto L488
L488:
	;
	goto L485
L489:
	;
	if v1997 == int32(0) {
		goto L480
	} else {
		goto L490
	}
L490:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+16)) = v1757 + int32(80)
	F_errmsg(m, int32(311214), v1757+int32(16))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v2029 = int32(447)
	goto L481
L493:
	;
	if v2014 == int32(0) {
		goto L480
	} else {
		goto L494
	}
L494:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+32)) = v1757 + int32(80)
	F_errmsg_internal(m, int32(305784), v1757+int32(32))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	v2029 = int32(452)
	goto L481
L497:
	;
	goto L480
L498:
	;
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842)+26)))
	if v2037 != int32(1) {
		v1951 = v1951 + int32(1)
		goto L474
	} else {
		goto L499
	}
L499:
	;
	goto L475
L500:
	;
	if base.Ui64(v1852) < base.Ui64(v2088) {
		goto L506
	} else {
		goto L507
	}
L501:
	;
	v2098 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	if v2098 == int32(0) {
		goto L500
	} else {
		goto L503
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1757))) = v2091
	*(*float64)(unsafe.Add(mBase, uint32(v1757)+8)) = base.F64_div(base.F64_convert_i64_u(v2088), float64(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+4)) = v1757 + int32(80)
	F_errmsg_internal(m, int32(159335), v1757)
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(520367), int32(432), int32(124781))
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	goto L500
L506:
	;
	v2118 = v2088
	goto L508
L507:
	;
	v2118 = v1852
	goto L508
L508:
	;
	v2126 = v2064
	v2129 = v2091
	v2131 = v1851 + v2088
	v2132 = v2118
	goto L442
L509:
	;
	if v2141 == int32(0) {
		goto L432
	} else {
		goto L510
	}
L510:
	;
	v2151 = v2126
	v2154 = v2129
	v2156 = v2131
	v2157 = v2132
	goto L441
L511:
	;
	if v2164 != 0 {
		v1842 = v2164
		v1846 = v2151
		v1849 = v2154
		v1851 = v2156
		v1852 = v2157
		goto L437
	} else {
		goto L512
	}
L512:
	;
	goto L438
L513:
	;
	F_errmsg_internal(m, int32(461972), int32(0))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(520367), int32(465), int32(124781))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
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
	F_errmsg_internal(m, int32(408301), int32(0))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(520367), int32(308), int32(124781))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
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
	*(*int64)(unsafe.Add(mBase, _consts[359])) = v2231 + v2230*int64(1000000) - int64(946684800000000)
	v2241 = int32(0)
	v2243 = m.G0
	v2245 = v2243 - int32(1152)
	m.G0 = v2245
	v2248 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v2248 <= v2241 {
		goto L524
	} else {
		goto L525
	}
L520:
	;
	return
L521:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L1
	} else {
		goto L598
	}
L522:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2561 = m.ExcPending
	if v2561 != 0 {
		goto L1
	} else {
		goto L594
	}
L523:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L1
	} else {
		goto L590
	}
L524:
	;
	m.G0 = v2245 + int32(1152)
	goto L520
L525:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v2256 = F_LWLockAcquire(m, v2252+int32(2304), int32(1))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v2259)+4))
	if int32(0) < v2260 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v2265 = v2259
	v2270 = v2241
	v2272 = v2241
	goto L530
L528:
	;
	v2477 = v2241
	goto L529
L529:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v2488+int32(2304))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L582
	}
L530:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2265+v2272<<(uint(int32(2))%32))+8))
	v2284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283)+44)))
	if v2284 == int32(0) {
		goto L533
	} else {
		goto L534
	}
L531:
	;
	v2477 = v2463
	goto L529
L532:
	;
	v2467 = v2272 + int32(1)
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+4))
	if v2467 < v2468 {
		v2265 = v2459
		v2270 = v2463
		v2272 = v2467
		goto L530
	} else {
		goto L581
	}
L533:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283)+46)))
	if v2287 != int32(1) {
		v2459 = v2265
		v2463 = v2270
		goto L532
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	v2290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283)+45)))
	if v2290 != 0 {
		v2459 = v2265
		v2463 = v2270
		goto L532
	} else {
		goto L537
	}
L536:
	;
	goto L535
L537:
	;
	v2291 = *(*int64)(unsafe.Add(mBase, uint32(v2283)+24))
	if base.Ui64(l0) < base.Ui64(v2291) {
		v2459 = v2265
		v2463 = v2270
		goto L532
	} else {
		goto L538
	}
L538:
	;
	v2293 = *(*int64)(unsafe.Add(mBase, uint32(v2283)+16))
	F_XlogReadTwoPhaseData(m, v2293, v2245+int32(120), v2245+int32(116))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2283)+32))
	v2301 = int32(-1)
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+120))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+116))
	v2304 = m.Env.Pgmem_crc32c(m, v2301, v2302, v2303)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+124)) = v2304 ^ v2301
	v2308 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	if base.Ui32(v2300) <= base.Ui32(int32(2)) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2327 = base.I64_extend_i32_u(v2300)
	goto L543
L542:
	;
	v2315 = int64(base.Ui64(v2308) >> (uint(int64(32)) % 64))
	if base.Ui32(base.I32_wrap_i64(v2308)) < base.Ui32(v2300) {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v2245)+100)) = uint32(v2327)
	v2330 = int64(base.Ui64(v2327) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2245)+96)) = uint32(v2330)
	v2338 = F_pg_snprintf(m, v2245+int32(128), int32(1024), int32(531012), v2245+int32(96))
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L547
	}
L544:
	;
	v2322 = (v2315 - int64(1)) & int64(4294967295)
	goto L546
L545:
	;
	v2322 = v2315
	goto L546
L546:
	;
	v2327 = base.I64_extend_i32_u(v2300) | v2322<<(uint(int64(32))%64)
	goto L543
L547:
	;
	v2343 = F_OpenTransientFile(m, v2245+int32(128), int32(577))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	if v2343 < int32(0) {
		goto L523
	} else {
		goto L549
	}
L549:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v2351 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v2351))) = int32(167772224)
	v2354 = F_write(m, v2343, v2302, v2303)
	mBase = m.M
	if v2354 != v2303 {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2357 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L551:
	;
	goto L552
L552:
	;
	v2384 = int32(4)
	v2385 = F_write(m, v2343, v2245+int32(124), v2384)
	mBase = m.M
	if v2385 != v2384 {
		goto L560
	} else {
		goto L561
	}
L553:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L555
L554:
	;
	goto L555
L555:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+80)) = v2245 + int32(128)
	F_errmsg(m, int32(310851), v2245+int32(80))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(518801), int32(1756), int32(405635))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L1
	} else {
		goto L559
	}
L559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L560:
	;
	v2389 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2389 == int32(0) {
		goto L563
	} else {
		goto L564
	}
L561:
	;
	goto L562
L562:
	;
	v2414 = int32(4155132)
	v2415 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	v2416 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2415))) = v2416
	v2419 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v2419))) = int32(167772223)
	v2424 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v2424 != int32(1) {
		v2438 = v2416
		goto L571
	} else {
		goto L572
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L565
L564:
	;
	goto L565
L565:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+64)) = v2245 + int32(128)
	F_errmsg(m, int32(310851), v2245-int32(-64))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	F_errfinish(m, int32(518801), int32(1765), int32(405635))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L570:
	;
	if v2438 != 0 {
		goto L522
	} else {
		goto L577
	}
L571:
	;
	goto L570
L572:
	;
	goto L573
L573:
	;
	v2429 = F_fsync(m, v2343)
	mBase = m.M
	if v2429 != int32(-1) {
		v2438 = v2429
		goto L571
	} else {
		goto L575
	}
L574:
	;
	v2438 = int32(-1)
	goto L571
L575:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2433 == int32(27) {
		goto L573
	} else {
		goto L576
	}
L576:
	;
	goto L574
L577:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v2440))) = int32(0)
	v2443 = F_CloseTransientFile(m, v2343)
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	if v2443 != 0 {
		goto L521
	} else {
		goto L579
	}
L579:
	;
	v2445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2283)+45)) = uint8(v2445)
	v2448 = v2283 + int32(16)
	v2449 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2448)+8)) = v2449
	*(*int64)(unsafe.Add(mBase, uint32(v2448))) = v2449
	F_pfree(m, v2302)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v2459 = v2458
	v2463 = v2270 + int32(1)
	goto L532
L581:
	;
	goto L531
L582:
	;
	F_fsync_fname(m, int32(376250), int32(1))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	v2498 = int32(*(*uint8)(unsafe.Add(mBase, _consts[358])))
	if v2498 != int32(1) {
		goto L524
	} else {
		goto L584
	}
L584:
	;
	if v2477 <= int32(0) {
		goto L524
	} else {
		goto L585
	}
L585:
	;
	v2505 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	if v2505 == int32(0) {
		goto L524
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245))) = v2477
	F_errmsg_plural(m, int32(267941), int32(150086), v2477, v2245)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(518801), int32(1876), int32(376272))
	mBase = m.M
	v2518 = m.ExcPending
	if v2518 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	goto L524
L590:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+16)) = v2245 + int32(128)
	F_errmsg(m, int32(310912), v2245+int32(16))
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(518801), int32(1744), int32(405635))
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
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
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+48)) = v2245 + int32(128)
	F_errmsg(m, int32(311214), v2245+int32(48))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_errfinish(m, int32(518801), int32(1777), int32(405635))
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L598:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+32)) = v2245 + int32(128)
	F_errmsg(m, int32(311009), v2245+int32(32))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	F_errfinish(m, int32(518801), int32(1783), int32(405635))
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
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
	var v45 float64
	_ = v45
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v69 float64
	_ = v69
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_abs(base.F64_sub(v8, v9))
	if base.F64_eq(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v15 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v8), v15)&base.F64_ne(base.F64_abs(v9), v15) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
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
				if base.F64_eq(base.F64_abs(v36), v38) != 0 {
					v69 = v38
					return v69
				} else {
					if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v69 = v38
						return v69
					} else {
						F_float_overflow_error(m)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v69 = v21
				return v69
			}
		}
	} else {
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		v31 = base.F64_abs(base.F64_sub(v28, v29))
		if base.F64_ne(v31, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v45 = math.Float64frombits(uint64(0x7ff8000000000000))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)) {
				v69 = v45
				return v69
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)) {
					v69 = v45
					return v69
				} else {
					v52 = base.F64_lt(v11, v31)
					if v52 != 0 {
						v53 = v31
					} else {
						v53 = v11
					}
					if v52 != 0 {
						v54 = v11
					} else {
						v54 = v31
					}
					if base.F64_eq(v54, float64(0)) != 0 {
						v69 = v53
						return v69
					} else {
						v57 = base.F64_div(v54, v53)
						v62 = base.F64_mul(v53, base.F64_sqrt(base.F64_add(base.F64_mul(v57, v57), float64(1))))
						if base.F64_eq(base.F64_abs(v62), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if base.F64_eq(v62, float64(0)) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v69 = v62
								return v69
							}
						}
					}
				}
			}
		} else {
			v36 = v28
			v37 = v29
			v38 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v36), v38) != 0 {
				v69 = v38
				return v69
			} else {
				if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v69 = v38
					return v69
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
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
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v40 int32
	_ = v40
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v150 int32
	_ = v150
	v12 = int32(0)
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v18 = base.F64_sub(v16, v17)
	if base.F64_ne(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L31
	}
L2:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = base.F64_sub(v28, v29)
	if base.F64_ne(base.F64_abs(v30), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if base.F64_eq(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if base.F64_ne(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
L6:
	;
	v40 = int32(2)
	if l1 < v40 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	if base.F64_eq(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	return v132
L11:
	;
	v111 = F_lseg_crossing(m, v18, v30, v100, v99)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L25
	} else {
		goto L29
	}
L12:
	;
	v99 = v30
	v100 = v18
	v108 = v12
	goto L11
L13:
	;
	goto L14
L14:
	;
	v52 = v18
	v53 = v30
	v55 = int32(1)
	v56 = v12
	goto L15
L15:
	;
	v61 = l2 + v55<<(uint(int32(4))%32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v64 = base.F64_sub(v62, v63)
	if base.F64_ne(base.F64_abs(v64), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v99 = v76
	v100 = v64
	v108 = v92
	goto L11
L17:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v61)+8))
	v75 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = base.F64_sub(v74, v75)
	if base.F64_ne(base.F64_abs(v76), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if base.F64_eq(base.F64_abs(v62), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.F64_ne(base.F64_abs(v63), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v86 = F_lseg_crossing(m, v64, v76, v52, v53)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if base.F64_eq(base.F64_abs(v74), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.F64_ne(base.F64_abs(v75), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	return int32(0)
L26:
	;
	if v86 == int32(2147483647) {
		v132 = v40
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v92 = v56 + v86
	v94 = v55 + int32(1)
	if v94 != l1 {
		v52 = v64
		v53 = v76
		v55 = v94
		v56 = v92
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	if v111 == int32(2147483647) {
		v132 = v40
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v132 = base.B2i32(v108 != int32(0)-v111)
	goto L10
L31:
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v60 float64
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v3 = float64(0)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v10, v11) != 0 {
		v60 = v3
		return v60
	} else {
		v13 = base.F64_sub(v10, v11)
		v14 = base.F64_abs(v13)
		if base.F64_le(v14, float64(1e-06)) != 0 {
			v60 = v3
			return v60
		} else {
			v17 = math.Float64frombits(uint64(0x7ff0000000000000))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v18, v19) != 0 {
				v60 = v17
				return v60
			} else {
				if base.F64_le(base.F64_abs(base.F64_sub(v18, v19)), float64(1e-06)) != 0 {
					v60 = v17
					return v60
				} else {
					if base.F64_ne(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v33 = base.F64_sub(v19, v18)
						v34 = base.F64_abs(v33)
						if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
								F_float_zero_divide_error(m)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v49 = base.F64_div(v13, v33)
								v51 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_ne(v49, float64(0)) != 0 {
										v60 = v49
										return v60
									} else {
										if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v60 = v49
											return v60
										}
									}
								}
							}
						} else {
							if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
									F_float_zero_divide_error(m)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v49 = base.F64_div(v13, v33)
									v51 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v49, float64(0)) != 0 {
											v60 = v49
											return v60
										} else {
											if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v60 = v49
												return v60
											}
										}
									}
								}
							} else {
								if base.F64_ne(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
										F_float_zero_divide_error(m)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v49 = base.F64_div(v13, v33)
										v51 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v49, float64(0)) != 0 {
												v60 = v49
												return v60
											} else {
												if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v60 = v49
													return v60
												}
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v33 = base.F64_sub(v19, v18)
							v34 = base.F64_abs(v33)
							if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
									F_float_zero_divide_error(m)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v49 = base.F64_div(v13, v33)
									v51 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v49, float64(0)) != 0 {
											v60 = v49
											return v60
										} else {
											if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v60 = v49
												return v60
											}
										}
									}
								}
							} else {
								if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
										F_float_zero_divide_error(m)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v49 = base.F64_div(v13, v33)
										v51 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v49, float64(0)) != 0 {
												v60 = v49
												return v60
											} else {
												if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v60 = v49
													return v60
												}
											}
										}
									}
								} else {
									if base.F64_ne(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
											F_float_zero_divide_error(m)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v49 = base.F64_div(v13, v33)
											v51 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_ne(v49, float64(0)) != 0 {
													v60 = v49
													return v60
												} else {
													if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return float64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v60 = v49
														return v60
													}
												}
											}
										}
									}
								}
							}
						} else {
							if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v33 = base.F64_sub(v19, v18)
								v34 = base.F64_abs(v33)
								if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
										F_float_zero_divide_error(m)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v49 = base.F64_div(v13, v33)
										v51 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v49, float64(0)) != 0 {
												v60 = v49
												return v60
											} else {
												if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v60 = v49
													return v60
												}
											}
										}
									}
								} else {
									if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
											F_float_zero_divide_error(m)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v49 = base.F64_div(v13, v33)
											v51 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_ne(v49, float64(0)) != 0 {
													v60 = v49
													return v60
												} else {
													if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return float64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v60 = v49
														return v60
													}
												}
											}
										}
									} else {
										if base.F64_ne(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v33, float64(0)) != 0 {
												F_float_zero_divide_error(m)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v49 = base.F64_div(v13, v33)
												v51 = math.Float64frombits(uint64(0x7ff0000000000000))
												if base.F64_eq(base.F64_abs(v49), v51)&base.F64_ne(v14, v51) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.F64_ne(v49, float64(0)) != 0 {
														v60 = v49
														return v60
													} else {
														if base.F64_ne(v34, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
															F_float_underflow_error(m)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return float64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v60 = v49
															return v60
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
				}
			}
		}
	}
}
func F_point_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initStringInfo(m, v7+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_appendStringInfoChar(m, v7+int32(16), int32(40))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
			v23 = F_float8out_internal(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_float8out_internal(m, v21)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v23
					F_appendStringInfo(m, v7+int32(16), int32(186144), v7)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v23)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v25)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoChar(m, v7+int32(16), int32(41))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									m.G0 = v7 + int32(32)
									return v43
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
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v46 int32
	_ = v46
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v17 = base.F64_sub(v15, v16)
		if base.F64_ne(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v29 = base.F64_sub(v27, v28)
			if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				return v11
			} else {
				if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					return v11
				} else {
					if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						return v11
					}
				}
			}
		} else {
			if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v29 = base.F64_sub(v27, v28)
				if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					return v11
				} else {
					if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						return v11
					} else {
						if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							return v11
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
					v29 = base.F64_sub(v27, v28)
					if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						return v11
					} else {
						if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							return v11
						} else {
							if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
								return v11
							}
						}
					}
				}
			}
		}
	}
}
