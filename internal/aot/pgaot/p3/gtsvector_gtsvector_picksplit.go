package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1060 int32
	_ = v1060
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1193 int32
	_ = v1193
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1239 int64
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1271 int64
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int64
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int64
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int64
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int64
	_ = v1287
	var v1291 int64
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1326 int64
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1337 int32
	_ = v1337
	var v1359 int64
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1363 int64
	_ = v1363
	var v1364 int64
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1374 int64
	_ = v1374
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int64
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1417 int64
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1431 int64
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int64
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int64
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int64
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int64
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1479 int64
	_ = v1479
	var v1480 int64
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1492 int64
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1501 int64
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1505 int64
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int64
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int64
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int64
	_ = v1517
	var v1521 int64
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1532 int64
	_ = v1532
	var v1562 int64
	_ = v1562
	var v1572 int32
	_ = v1572
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1678 int32
	_ = v1678
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1725 int64
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1757 int64
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1761 int64
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1765 int64
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1769 int64
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1773 int64
	_ = v1773
	var v1777 int64
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1812 int64
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1822 int32
	_ = v1822
	var v1845 int64
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int64
	_ = v1849
	var v1850 int64
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1860 int64
	_ = v1860
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1887 int64
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1903 int64
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1917 int64
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1927 int64
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int64
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1945 int64
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int64
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1965 int64
	_ = v1965
	var v1966 int64
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1978 int64
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1987 int64
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int64
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1995 int64
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1999 int64
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2003 int64
	_ = v2003
	var v2007 int64
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2018 int64
	_ = v2018
	var v2048 int64
	_ = v2048
	var v2057 int32
	_ = v2057
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2120 int32
	_ = v2120
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2207 int32
	_ = v2207
	var v2213 int32
	_ = v2213
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2279 int32
	_ = v2279
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2312 int32
	_ = v2312
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2399 int32
	_ = v2399
	var v2405 int32
	_ = v2405
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2471 int32
	_ = v2471
	var v2522 int32
	_ = v2522
	var v2527 int32
	_ = v2527
	var v2538 int32
	_ = v2538
	var v2555 int32
	_ = v2555
	var v2560 int32
	_ = v2560
	var v2570 int32
	_ = v2570
	v2 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 == v2 {
		v51 = v2
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
		if v38 == int32(0) {
			v51 = v2
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
			if v41 != int32(7) {
				v51 = v2
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				if v44 != int32(17) {
					v51 = v2
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
					v51 = v47 ^ int32(1)
				}
			}
		}
	}
	if v51&int32(1) != 0 {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v55 = F_get_fn_opclass_options(m, v54)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
			v60 = v59
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			v63 = v61 + int32(65534)
			v65 = v63 & int32(65535)
			v67 = v65 + int32(2)
			v69 = v67 << (uint(int32(1)) % 32)
			v70 = F_palloc(m, v69)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v70
				v73 = F_palloc(m, v69)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v73
					v79 = F_palloc(m, v67<<(uint(int32(3))%32))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v82 = F_palloc(m, v60*v67)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v84 = int32(0)
							v85 = v2
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v79+v84<<(uint(int32(3))%32))+4)) = v82 + v84*v60
								v121 = v85 + int32(1)
								v123 = v121 & int32(65535)
								if base.Ui32(v123) < base.Ui32(v67) {
									v84 = v123
									v85 = v121
									continue
								} else {
									break
								}
								break
							}
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
							v126 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)) = uint8(v126)
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
							if v128&int32(1) != 0 {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
								v135 = int32(base.Ui32(v131)>>(uint(int32(2))%32)) - int32(8)
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								if v136&int32(3) != 0 {
									v158 = v60
									v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
									mBase = m.M
								} else {
									if base.Ui32(int32(1024)) < base.Ui32(v60) {
										v158 = v60
										v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
										mBase = m.M
									} else {
										if v60&int32(3) != 0 {
											v158 = v60
											v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
											mBase = m.M
										} else {
											v143 = v136 + v60
											if base.Ui32(v143) <= base.Ui32(v136) {
											} else {
												v148 = v136 + int32(4)
												if base.Ui32(v148) < base.Ui32(v143) {
													v150 = v143
												} else {
													v150 = v148
												}
												v158 = (v136^int32(-1)+v150)&int32(-4) + int32(4)
												v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
												mBase = m.M
											}
										}
									}
								}
								if base.Ui32(v135) < base.Ui32(int32(4)) {
								} else {
									v167 = v125 + int32(8)
									v169 = v60 << (uint(int32(3)) % 32)
									v170 = int32(1)
									v172 = int32(base.Ui32(v135) >> (uint(int32(2)) % 32))
									if base.Ui32(v172) <= base.Ui32(v170) {
										v175 = v170
									} else {
										v175 = v172
									}
									v178 = int32(0)
									if base.Ui32(int32(8)) <= base.Ui32(v135) {
										v184 = v178
										v188 = int32(0)
										for {
											v214 = int32(2)
											v216 = v167 + v184<<(uint(v214)%32)
											v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
											v218 = base.I32_rem_u_s(v217, v169)
											v219 = int32(3)
											v221 = v136 + int32(base.Ui32(v218)>>(uint(v219)%32))
											v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
											v223 = int32(1)
											v224 = int32(7)
											v227 = v222 | v223<<(uint(v218&v224)%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v227)
											v229 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
											v230 = base.I32_rem_u_s(v229, v169)
											v233 = v136 + int32(base.Ui32(v230)>>(uint(v219)%32))
											v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
											v239 = v234 | v223<<(uint(v230&v224)%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v239)
											v242 = v184 + v214
											v244 = v188 + v214
											if v244 != v175&int32(1073741822) {
												v184 = v242
												v188 = v244
												continue
											} else {
												break
											}
											break
										}
										v246 = v242
									} else {
										v246 = v178
									}
									if v175&int32(1) == int32(0) {
									} else {
										v281 = *(*int32)(unsafe.Add(mBase, uint32(v167+v246<<(uint(int32(2))%32))))
										v282 = base.I32_rem_u_s(v281, v169)
										v285 = v136 + int32(base.Ui32(v282)>>(uint(int32(3))%32))
										v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
										v291 = v286 | int32(1)<<(uint(v282&int32(7))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v291)
									}
								}
							} else {
								if v128&int32(4) != 0 {
									v295 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)) = uint8(v295)
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									if v60 != 0 {
										v300 = F__emscripten_memcpy_bulkmem(m, v297, v125+int32(8), v60)
										mBase = m.M
									} else {
									}
								}
							}
							v333 = v32 + int32(4)
							if base.Ui32(int32(2)) <= base.Ui32(v65) {
								v336 = int32(3)
								v344 = v60 << (uint(v336) % 32)
								v362 = int32(-1)
								v363 = int32(1)
								v364 = v2
								v366 = v2
								for {
									v381 = v363 + int32(1)
									v382 = v381
									v388 = v381
									v397 = v362
									v399 = v364
									v401 = v366
									for {
										if v363 != int32(1) {
										} else {
											v417 = *(*int32)(unsafe.Add(mBase, uint32(v333+v388<<(uint(int32(4))%32))))
											v420 = v79 + v388<<(uint(int32(3))%32)
											v421 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v421)
											v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
											if v423&int32(1) != 0 {
												v426 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
												v430 = int32(base.Ui32(v426)>>(uint(int32(2))%32)) - int32(8)
												v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
												v434 = int32(0)
												if (base.B2i32(v431&int32(3) != v434)|(base.B2i32(v60&v336 != int32(0))|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v60))))&int32(1) == v434 {
													v441 = v431 + v60
													if base.Ui32(v441) <= base.Ui32(v431) {
													} else {
														v446 = v431 + int32(4)
														if base.Ui32(v446) < base.Ui32(v441) {
															v448 = v441
														} else {
															v448 = v446
														}
														v454 = (v431^int32(-1)+v448)&int32(-4) + int32(4)
														v458 = F__emscripten_memset_bulkmem(m, v431, base.I32_extend8_s(int32(0)), v454)
														mBase = m.M
													}
												} else {
													v454 = v60
													v458 = F__emscripten_memset_bulkmem(m, v431, base.I32_extend8_s(int32(0)), v454)
													mBase = m.M
												}
												if base.Ui32(v430) < base.Ui32(int32(4)) {
												} else {
													v464 = v417 + int32(8)
													v465 = int32(1)
													v467 = int32(base.Ui32(v430) >> (uint(int32(2)) % 32))
													if base.Ui32(v467) <= base.Ui32(v465) {
														v470 = v465
													} else {
														v470 = v467
													}
													v473 = int32(0)
													if base.Ui32(int32(8)) <= base.Ui32(v430) {
														v479 = v473
														v483 = int32(0)
														for {
															v509 = int32(2)
															v511 = v464 + v479<<(uint(v509)%32)
															v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
															v513 = base.I32_rem_u_s(v512, v344)
															v514 = int32(3)
															v516 = v431 + int32(base.Ui32(v513)>>(uint(v514)%32))
															v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
															v518 = int32(1)
															v519 = int32(7)
															v522 = v517 | v518<<(uint(v513&v519)%32)
															*(*uint8)(unsafe.Add(mBase, uint32(v516))) = uint8(v522)
															v524 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
															v525 = base.I32_rem_u_s(v524, v344)
															v528 = v431 + int32(base.Ui32(v525)>>(uint(v514)%32))
															v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
															v534 = v529 | v518<<(uint(v525&v519)%32)
															*(*uint8)(unsafe.Add(mBase, uint32(v528))) = uint8(v534)
															v537 = v479 + v509
															v539 = v483 + v509
															if v539 != v470&int32(1073741822) {
																v479 = v537
																v483 = v539
																continue
															} else {
																break
															}
															break
														}
														v541 = v537
													} else {
														v541 = v473
													}
													if v470&int32(1) == int32(0) {
													} else {
														v576 = *(*int32)(unsafe.Add(mBase, uint32(v464+v541<<(uint(int32(2))%32))))
														v577 = base.I32_rem_u_s(v576, v344)
														v580 = v431 + int32(base.Ui32(v577)>>(uint(int32(3))%32))
														v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
														v586 = v581 | int32(1)<<(uint(v577&int32(7))%32)
														*(*uint8)(unsafe.Add(mBase, uint32(v580))) = uint8(v586)
													}
												}
											} else {
												if v423&int32(4) != 0 {
													v590 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v590)
												} else {
													v592 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
													if v60 != 0 {
														v595 = F__emscripten_memcpy_bulkmem(m, v592, v417+int32(8), v60)
														mBase = m.M
													} else {
													}
												}
											}
										}
										v630 = F_hemdistcache_1(m, v79+v388<<(uint(int32(3))%32), v79+v363<<(uint(int32(3))%32), v60)
										mBase = m.M
										v631 = base.B2i32(v397 < v630)
										if v397 < v630 {
											v632 = v630
										} else {
											v632 = v397
										}
										if v397 < v630 {
											v633 = v382
										} else {
											v633 = v401
										}
										if v397 < v630 {
											v634 = v363
										} else {
											v634 = v399
										}
										v636 = v382 + int32(1)
										v637 = int32(65535)
										v638 = v636 & v637
										if base.Ui32(v638) <= base.Ui32(v63&v637) {
											v382 = v636
											v388 = v638
											v397 = v632
											v399 = v634
											v401 = v633
											continue
										} else {
											break
										}
										break
									}
									if v381 != v65 {
										v362 = v632
										v363 = v381
										v364 = v634
										v366 = v633
										continue
									} else {
										break
									}
									break
								}
								v660 = v634
								v662 = v633
							} else {
								v660 = v2
								v662 = v2
							}
							v673 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v673
							*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v673
							v678 = int32(65535)
							v686 = base.B2i32(v660&v678 == v673) | base.B2i32(v662&v678 == v673)
							if v686 != 0 {
								v687 = int32(1)
							} else {
								v687 = v660
							}
							v692 = v79 + v687&int32(65535)<<(uint(int32(3))%32)
							v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
							v694 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
							v695 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
							v696 = int32(8)
							v698 = v60 + v696
							v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
							if v699 != 0 {
								v700 = v696
							} else {
								v700 = v698
							}
							v701 = F_palloc(m, v700)
							mBase = m.M
							v702 = m.ExcPending
							if v702 != 0 {
								return int32(0)
							} else {
								if v699 != 0 {
									v705 = int32(6)
								} else {
									v705 = int32(2)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v705
								v707 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v701))) = v700 << (uint(v707) % 32)
								if v686 != 0 {
									v711 = v707
								} else {
									v711 = v662
								}
								if v693 == int32(0) {
								} else {
									if v699 != 0 {
									} else {
										if v60 != 0 {
											v716 = F__emscripten_memcpy_bulkmem(m, v701+int32(8), v693, v60)
											mBase = m.M
										} else {
										}
									}
								}
								v722 = v79 + v711&int32(65535)<<(uint(int32(3))%32)
								v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
								v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722))))
								if v725 != 0 {
									v726 = int32(8)
								} else {
									v726 = v698
								}
								v727 = F_palloc(m, v726)
								mBase = m.M
								v728 = m.ExcPending
								if v728 != 0 {
									return int32(0)
								} else {
									if v725 != 0 {
										v731 = int32(6)
									} else {
										v731 = int32(2)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v731
									*(*int32)(unsafe.Add(mBase, uint32(v727))) = v726 << (uint(int32(2)) % 32)
									if v723 == int32(0) {
									} else {
										if v725 != 0 {
										} else {
											if v60 != 0 {
												v740 = F__emscripten_memcpy_bulkmem(m, v727+int32(8), v723, v60)
												mBase = m.M
											} else {
											}
										}
									}
									v742 = int32(65535)
									v743 = v61 + v742
									v745 = v743 & v742
									v749 = *(*int32)(unsafe.Add(mBase, uint32(v333+v745<<(uint(int32(4))%32))))
									v752 = v79 + v745<<(uint(int32(3))%32)
									v753 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v752))) = uint8(v753)
									v755 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
									if v755&int32(1) != 0 {
										v758 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
										v762 = int32(base.Ui32(v758)>>(uint(int32(2))%32)) - int32(8)
										v763 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
										if v763&int32(3) != 0 {
											v785 = v60
											v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
											mBase = m.M
										} else {
											if base.Ui32(int32(1024)) < base.Ui32(v60) {
												v785 = v60
												v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
												mBase = m.M
											} else {
												if v60&int32(3) != 0 {
													v785 = v60
													v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
													mBase = m.M
												} else {
													v770 = v763 + v60
													if base.Ui32(v770) <= base.Ui32(v763) {
													} else {
														v775 = v763 + int32(4)
														if base.Ui32(v775) < base.Ui32(v770) {
															v777 = v770
														} else {
															v777 = v775
														}
														v785 = (v763^int32(-1)+v777)&int32(-4) + int32(4)
														v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
														mBase = m.M
													}
												}
											}
										}
										if base.Ui32(v762) < base.Ui32(int32(4)) {
										} else {
											v794 = v749 + int32(8)
											v796 = v60 << (uint(int32(3)) % 32)
											v797 = int32(1)
											v799 = int32(base.Ui32(v762) >> (uint(int32(2)) % 32))
											if base.Ui32(v799) <= base.Ui32(v797) {
												v802 = v797
											} else {
												v802 = v799
											}
											v805 = int32(0)
											if base.Ui32(int32(8)) <= base.Ui32(v762) {
												v811 = v805
												v815 = int32(0)
												for {
													v841 = int32(2)
													v843 = v794 + v811<<(uint(v841)%32)
													v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
													v845 = base.I32_rem_u_s(v844, v796)
													v846 = int32(3)
													v848 = v763 + int32(base.Ui32(v845)>>(uint(v846)%32))
													v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
													v850 = int32(1)
													v851 = int32(7)
													v854 = v849 | v850<<(uint(v845&v851)%32)
													*(*uint8)(unsafe.Add(mBase, uint32(v848))) = uint8(v854)
													v856 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
													v857 = base.I32_rem_u_s(v856, v796)
													v860 = v763 + int32(base.Ui32(v857)>>(uint(v846)%32))
													v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860))))
													v866 = v861 | v850<<(uint(v857&v851)%32)
													*(*uint8)(unsafe.Add(mBase, uint32(v860))) = uint8(v866)
													v869 = v811 + v841
													v871 = v815 + v841
													if v871 != v802&int32(1073741822) {
														v811 = v869
														v815 = v871
														continue
													} else {
														break
													}
													break
												}
												v873 = v869
											} else {
												v873 = v805
											}
											if v802&int32(1) == int32(0) {
											} else {
												v908 = *(*int32)(unsafe.Add(mBase, uint32(v794+v873<<(uint(int32(2))%32))))
												v909 = base.I32_rem_u_s(v908, v796)
												v912 = v763 + int32(base.Ui32(v909)>>(uint(int32(3))%32))
												v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
												v918 = v913 | int32(1)<<(uint(v909&int32(7))%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v918)
											}
										}
									} else {
										if v755&int32(4) != 0 {
											v922 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v752))) = uint8(v922)
										} else {
											v924 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
											if v60 != 0 {
												v927 = F__emscripten_memcpy_bulkmem(m, v924, v749+int32(8), v60)
												mBase = m.M
											} else {
											}
										}
									}
									v961 = F_palloc(m, v745<<(uint(int32(3))%32))
									mBase = m.M
									v962 = m.ExcPending
									if v962 != 0 {
										return int32(0)
									} else {
										if v61&int32(65535) == int32(1) {
											F_pg_qsort(m, v961, v745, int32(8), int32(1522))
											mBase = m.M
											v970 = m.ExcPending
											if v970 != 0 {
												return int32(0)
											} else {
												v2555 = v694
												v2560 = v695
												v2570 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v2555))) = uint16(v2570)
												*(*uint16)(unsafe.Add(mBase, uint32(v2560))) = uint16(v2570)
												*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v727
												*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v701
												return v31
											}
										} else {
											v971 = int32(8)
											v972 = v727 + v971
											v974 = v701 + v971
											v975 = int32(1)
											v977 = v975
											v981 = v975
											for {
												v1008 = v981 << (uint(int32(3)) % 32)
												v1009 = v961 + v1008
												*(*uint16)(unsafe.Add(mBase, uint32(v1009-int32(8)))) = uint16(v977)
												v1015 = v1008 + v79
												v1016 = F_hemdistcache_1(m, v692, v1015, v60)
												mBase = m.M
												v1017 = F_hemdistcache_1(m, v722, v1015, v60)
												mBase = m.M
												v1018 = v1016 - v1017
												v1020 = v1018 >> (uint(int32(31)) % 32)
												*(*int32)(unsafe.Add(mBase, uint32(v1009-int32(4)))) = v1018 ^ v1020 - v1020
												v1025 = v977 + int32(1)
												v1026 = int32(65535)
												v1027 = v1025 & v1026
												if base.Ui32(v1027) <= base.Ui32(v743&v1026) {
													v977 = v1025
													v981 = v1027
													continue
												} else {
													break
												}
												break
											}
											F_pg_qsort(m, v961, v745, int32(8), int32(1522))
											mBase = m.M
											v1034 = m.ExcPending
											if v1034 != 0 {
												return int32(0)
											} else {
												v1035 = int32(1)
												if base.Ui32(v745) <= base.Ui32(v1035) {
													v1038 = v1035
												} else {
													v1038 = v745
												}
												v1040 = v60 & int32(2147483644)
												v1041 = int32(3)
												v1042 = v60 & v1041
												v1044 = v60 & int32(-4)
												v1046 = v60 & int32(2147483646)
												v1047 = int32(1)
												v1048 = v60 & v1047
												v1050 = v60 - v1047
												v1052 = v60 << (uint(v1041) % 32)
												v1060 = int32(0)
												v1071 = v694
												v1076 = v695
												for {
													v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961+v1060<<(uint(int32(3))%32)))))
													if v687&int32(65535) == v1089 {
														*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v687)
														v1092 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1092 + int32(1)
														v2522 = v1071 + int32(2)
														v2527 = v1076
													} else {
														if v711&int32(65535) == v1089 {
															*(*uint16)(unsafe.Add(mBase, uint32(v1076))) = uint16(v711)
															v1102 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1102 + int32(1)
															v2522 = v1071
															v2527 = v1076 + int32(2)
														} else {
															v1108 = v79 + v1089<<(uint(int32(3))%32)
															v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
															v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+4)))
															if v1110&int32(4) == int32(0) {
																if v1109&int32(1) == int32(0) {
																	if v60 <= int32(0) {
																		v1572 = int32(0)
																	} else {
																		v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1128 = int32(0)
																		if v1050 != 0 {
																			v1131 = v1128
																			v1138 = v1128
																			v1141 = v1128
																			for {
																				v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131+v974))))
																				v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131+v1127))))
																				v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162^v1164)+uint32(_consts[1053]))))
																				v1171 = v1131 | int32(1)
																				v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974+v1171))))
																				v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127+v1171))))
																				v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173^v1175)+uint32(_consts[1053]))))
																				v1180 = v1138 + v1168 + v1179
																				v1181 = int32(2)
																				v1182 = v1131 + v1181
																				v1184 = v1141 + v1181
																				if v1184 != v1046 {
																					v1131 = v1182
																					v1138 = v1180
																					v1141 = v1184
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1186 = v1182
																			v1193 = v1180
																		} else {
																			v1186 = v1128
																			v1193 = v1128
																		}
																		if v1048 == int32(0) {
																			v1572 = v1193
																		} else {
																			v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v974))))
																			v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v1127))))
																			v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219^v1221)+uint32(_consts[1053]))))
																			v1572 = v1193 + v1225
																		}
																	}
																} else {
																	if v1109&int32(1) == int32(0) {
																		v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1232 = v1231
																	} else {
																		v1232 = v974
																	}
																	if v60 <= int32(3) {
																		if v60 == int32(0) {
																			v1562 = int64(0)
																		} else {
																			v1239 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																				v1242 = v1232
																				v1249 = int32(0)
																				v1271 = v1239
																				for {
																					v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+3)))
																					v1275 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+uint32(_consts[1053]))))
																					v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+2)))
																					v1279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+uint32(_consts[1053]))))
																					v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
																					v1283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+uint32(_consts[1053]))))
																					v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
																					v1287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1284)+uint32(_consts[1053]))))
																					v1291 = v1275 + (v1279 + (v1283 + (v1271 + v1287)))
																					v1292 = int32(4)
																					v1293 = v1242 + v1292
																					v1295 = v1249 + v1292
																					if v1295 != v1044 {
																						v1242 = v1293
																						v1249 = v1295
																						v1271 = v1291
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1297 = v1293
																				v1326 = v1291
																			} else {
																				v1297 = v1232
																				v1326 = v1239
																			}
																			v1327 = int32(0)
																			if v1042 == v1327 {
																				v1562 = v1326
																			} else {
																				v1330 = v1297
																				v1337 = v1327
																				v1359 = v1326
																				for {
																					v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330))))
																					v1363 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+uint32(_consts[1053]))))
																					v1364 = v1359 + v1363
																					v1365 = int32(1)
																					v1368 = v1337 + v1365
																					if v1368 != v1042 {
																						v1330 = v1330 + v1365
																						v1337 = v1368
																						v1359 = v1364
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1562 = v1364
																			}
																		}
																	} else {
																		v1374 = int64(0)
																		if v60 < int32(4) {
																			v1453 = v1232
																			v1454 = v60
																			v1459 = v1374
																		} else {
																			if v1232 != (v1232+int32(3))&int32(-4) {
																				v1453 = v1232
																				v1454 = v60
																				v1459 = v1374
																			} else {
																				v1383 = v60 - int32(4)
																				v1387 = int32(base.Ui32(v1383)>>(uint(int32(2))%32)) + int32(1)
																				v1389 = v1387 & int32(3)
																				if base.Ui32(v1383) < base.Ui32(int32(12)) {
																					v1425 = v1232
																					v1426 = v60
																					v1431 = v1374
																				} else {
																					v1395 = v1232
																					v1396 = v60
																					v1397 = int32(0)
																					v1401 = v1374
																					for {
																						v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
																						v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+8))
																						v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
																						v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1395)))
																						v1417 = base.I64_extend_i32_u(base.I32_popcnt(v1402)) + (base.I64_extend_i32_u(base.I32_popcnt(v1405)) + (base.I64_extend_i32_u(base.I32_popcnt(v1408)) + (v1401 + base.I64_extend_i32_u(base.I32_popcnt(v1411)))))
																						v1418 = int32(16)
																						v1419 = v1396 - v1418
																						v1421 = v1395 + v1418
																						v1423 = v1397 + int32(4)
																						if v1423 != v1387&int32(2147483644) {
																							v1395 = v1421
																							v1396 = v1419
																							v1397 = v1423
																							v1401 = v1417
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1425 = v1421
																					v1426 = v1419
																					v1431 = v1417
																				}
																				if v1389 == int32(0) {
																					v1453 = v1425
																					v1454 = v1426
																					v1459 = v1431
																				} else {
																					v1436 = v1426
																					v1437 = v1425
																					v1438 = int32(0)
																					v1441 = v1431
																					for {
																						v1442 = int32(4)
																						v1443 = v1436 - v1442
																						v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
																						v1447 = v1441 + base.I64_extend_i32_u(base.I32_popcnt(v1444))
																						v1449 = v1437 + v1442
																						v1451 = v1438 + int32(1)
																						if v1451 != v1389 {
																							v1436 = v1443
																							v1437 = v1449
																							v1438 = v1451
																							v1441 = v1447
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1453 = v1449
																					v1454 = v1443
																					v1459 = v1447
																				}
																			}
																		}
																		if v1454 == int32(0) {
																			v1532 = v1459
																		} else {
																			v1463 = v1454 & int32(3)
																			if v1463 == int32(0) {
																				v1486 = v1453
																				v1488 = v1454
																				v1492 = v1459
																			} else {
																				v1469 = v1454
																				v1470 = v1453
																				v1471 = int32(0)
																				v1473 = v1459
																				for {
																					v1474 = int32(1)
																					v1475 = v1469 - v1474
																					v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
																					v1479 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+uint32(_consts[1053]))))
																					v1480 = v1473 + v1479
																					v1482 = v1470 + v1474
																					v1484 = v1471 + v1474
																					if v1484 != v1463 {
																						v1469 = v1475
																						v1470 = v1482
																						v1471 = v1484
																						v1473 = v1480
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1486 = v1482
																				v1488 = v1475
																				v1492 = v1480
																			}
																			if base.Ui32(v1454) < base.Ui32(int32(4)) {
																				v1532 = v1492
																			} else {
																				v1495 = v1486
																				v1497 = v1488
																				v1501 = v1492
																				for {
																					v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+3)))
																					v1505 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1502)+uint32(_consts[1053]))))
																					v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+2)))
																					v1509 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_consts[1053]))))
																					v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+1)))
																					v1513 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+uint32(_consts[1053]))))
																					v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495))))
																					v1517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1514)+uint32(_consts[1053]))))
																					v1521 = v1505 + (v1509 + (v1513 + (v1501 + v1517)))
																					v1522 = int32(4)
																					v1525 = v1497 - v1522
																					if v1525 != 0 {
																						v1495 = v1495 + v1522
																						v1497 = v1525
																						v1501 = v1521
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1532 = v1521
																			}
																		}
																		v1562 = v1532
																	}
																	v1572 = v1052 - base.I32_wrap_i64(v1562)
																}
															} else {
																if v1109&int32(1) == int32(0) {
																	if v1109&int32(1) == int32(0) {
																		v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1232 = v1231
																	} else {
																		v1232 = v974
																	}
																	if v60 <= int32(3) {
																		if v60 == int32(0) {
																			v1562 = int64(0)
																		} else {
																			v1239 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																				v1242 = v1232
																				v1249 = int32(0)
																				v1271 = v1239
																				for {
																					v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+3)))
																					v1275 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+uint32(_consts[1053]))))
																					v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+2)))
																					v1279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+uint32(_consts[1053]))))
																					v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
																					v1283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+uint32(_consts[1053]))))
																					v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
																					v1287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1284)+uint32(_consts[1053]))))
																					v1291 = v1275 + (v1279 + (v1283 + (v1271 + v1287)))
																					v1292 = int32(4)
																					v1293 = v1242 + v1292
																					v1295 = v1249 + v1292
																					if v1295 != v1044 {
																						v1242 = v1293
																						v1249 = v1295
																						v1271 = v1291
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1297 = v1293
																				v1326 = v1291
																			} else {
																				v1297 = v1232
																				v1326 = v1239
																			}
																			v1327 = int32(0)
																			if v1042 == v1327 {
																				v1562 = v1326
																			} else {
																				v1330 = v1297
																				v1337 = v1327
																				v1359 = v1326
																				for {
																					v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330))))
																					v1363 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+uint32(_consts[1053]))))
																					v1364 = v1359 + v1363
																					v1365 = int32(1)
																					v1368 = v1337 + v1365
																					if v1368 != v1042 {
																						v1330 = v1330 + v1365
																						v1337 = v1368
																						v1359 = v1364
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1562 = v1364
																			}
																		}
																	} else {
																		v1374 = int64(0)
																		if v60 < int32(4) {
																			v1453 = v1232
																			v1454 = v60
																			v1459 = v1374
																		} else {
																			if v1232 != (v1232+int32(3))&int32(-4) {
																				v1453 = v1232
																				v1454 = v60
																				v1459 = v1374
																			} else {
																				v1383 = v60 - int32(4)
																				v1387 = int32(base.Ui32(v1383)>>(uint(int32(2))%32)) + int32(1)
																				v1389 = v1387 & int32(3)
																				if base.Ui32(v1383) < base.Ui32(int32(12)) {
																					v1425 = v1232
																					v1426 = v60
																					v1431 = v1374
																				} else {
																					v1395 = v1232
																					v1396 = v60
																					v1397 = int32(0)
																					v1401 = v1374
																					for {
																						v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
																						v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+8))
																						v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
																						v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1395)))
																						v1417 = base.I64_extend_i32_u(base.I32_popcnt(v1402)) + (base.I64_extend_i32_u(base.I32_popcnt(v1405)) + (base.I64_extend_i32_u(base.I32_popcnt(v1408)) + (v1401 + base.I64_extend_i32_u(base.I32_popcnt(v1411)))))
																						v1418 = int32(16)
																						v1419 = v1396 - v1418
																						v1421 = v1395 + v1418
																						v1423 = v1397 + int32(4)
																						if v1423 != v1387&int32(2147483644) {
																							v1395 = v1421
																							v1396 = v1419
																							v1397 = v1423
																							v1401 = v1417
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1425 = v1421
																					v1426 = v1419
																					v1431 = v1417
																				}
																				if v1389 == int32(0) {
																					v1453 = v1425
																					v1454 = v1426
																					v1459 = v1431
																				} else {
																					v1436 = v1426
																					v1437 = v1425
																					v1438 = int32(0)
																					v1441 = v1431
																					for {
																						v1442 = int32(4)
																						v1443 = v1436 - v1442
																						v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
																						v1447 = v1441 + base.I64_extend_i32_u(base.I32_popcnt(v1444))
																						v1449 = v1437 + v1442
																						v1451 = v1438 + int32(1)
																						if v1451 != v1389 {
																							v1436 = v1443
																							v1437 = v1449
																							v1438 = v1451
																							v1441 = v1447
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1453 = v1449
																					v1454 = v1443
																					v1459 = v1447
																				}
																			}
																		}
																		if v1454 == int32(0) {
																			v1532 = v1459
																		} else {
																			v1463 = v1454 & int32(3)
																			if v1463 == int32(0) {
																				v1486 = v1453
																				v1488 = v1454
																				v1492 = v1459
																			} else {
																				v1469 = v1454
																				v1470 = v1453
																				v1471 = int32(0)
																				v1473 = v1459
																				for {
																					v1474 = int32(1)
																					v1475 = v1469 - v1474
																					v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
																					v1479 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+uint32(_consts[1053]))))
																					v1480 = v1473 + v1479
																					v1482 = v1470 + v1474
																					v1484 = v1471 + v1474
																					if v1484 != v1463 {
																						v1469 = v1475
																						v1470 = v1482
																						v1471 = v1484
																						v1473 = v1480
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1486 = v1482
																				v1488 = v1475
																				v1492 = v1480
																			}
																			if base.Ui32(v1454) < base.Ui32(int32(4)) {
																				v1532 = v1492
																			} else {
																				v1495 = v1486
																				v1497 = v1488
																				v1501 = v1492
																				for {
																					v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+3)))
																					v1505 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1502)+uint32(_consts[1053]))))
																					v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+2)))
																					v1509 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_consts[1053]))))
																					v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+1)))
																					v1513 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+uint32(_consts[1053]))))
																					v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495))))
																					v1517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1514)+uint32(_consts[1053]))))
																					v1521 = v1505 + (v1509 + (v1513 + (v1501 + v1517)))
																					v1522 = int32(4)
																					v1525 = v1497 - v1522
																					if v1525 != 0 {
																						v1495 = v1495 + v1522
																						v1497 = v1525
																						v1501 = v1521
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1532 = v1521
																			}
																		}
																		v1562 = v1532
																	}
																	v1572 = v1052 - base.I32_wrap_i64(v1562)
																} else {
																	v1572 = int32(0)
																}
															}
															v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
															v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+4)))
															if v1596&int32(4) == int32(0) {
																if v1595&int32(1) == int32(0) {
																	if v60 <= int32(0) {
																		v2057 = int32(0)
																	} else {
																		v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1614 = int32(0)
																		if v1050 != 0 {
																			v1617 = v1614
																			v1623 = v1614
																			v1625 = v1614
																			for {
																				v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617+v972))))
																				v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617+v1613))))
																				v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648^v1650)+uint32(_consts[1053]))))
																				v1657 = v1617 | int32(1)
																				v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972+v1657))))
																				v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657+v1613))))
																				v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1659^v1661)+uint32(_consts[1053]))))
																				v1666 = v1623 + v1654 + v1665
																				v1667 = int32(2)
																				v1668 = v1617 + v1667
																				v1670 = v1625 + v1667
																				if v1670 != v1046 {
																					v1617 = v1668
																					v1623 = v1666
																					v1625 = v1670
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1672 = v1668
																			v1678 = v1666
																		} else {
																			v1672 = v1614
																			v1678 = v1614
																		}
																		if v1048 == int32(0) {
																			v2057 = v1678
																		} else {
																			v1705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672+v972))))
																			v1707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672+v1613))))
																			v1711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1705^v1707)+uint32(_consts[1053]))))
																			v2057 = v1678 + v1711
																		}
																	}
																} else {
																	if v1595&int32(1) == int32(0) {
																		v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1718 = v1717
																	} else {
																		v1718 = v972
																	}
																	if v60 <= int32(3) {
																		if v60 == int32(0) {
																			v2048 = int64(0)
																		} else {
																			v1725 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																				v1728 = v1718
																				v1734 = int32(0)
																				v1757 = v1725
																				for {
																					v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+3)))
																					v1761 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+uint32(_consts[1053]))))
																					v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+2)))
																					v1765 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1762)+uint32(_consts[1053]))))
																					v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+1)))
																					v1769 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1766)+uint32(_consts[1053]))))
																					v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
																					v1773 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+uint32(_consts[1053]))))
																					v1777 = v1761 + (v1765 + (v1769 + (v1757 + v1773)))
																					v1778 = int32(4)
																					v1779 = v1728 + v1778
																					v1781 = v1734 + v1778
																					if v1781 != v1044 {
																						v1728 = v1779
																						v1734 = v1781
																						v1757 = v1777
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1783 = v1779
																				v1812 = v1777
																			} else {
																				v1783 = v1718
																				v1812 = v1725
																			}
																			v1813 = int32(0)
																			if v1042 == v1813 {
																				v2048 = v1812
																			} else {
																				v1816 = v1783
																				v1822 = v1813
																				v1845 = v1812
																				for {
																					v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816))))
																					v1849 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+uint32(_consts[1053]))))
																					v1850 = v1845 + v1849
																					v1851 = int32(1)
																					v1854 = v1822 + v1851
																					if v1854 != v1042 {
																						v1816 = v1816 + v1851
																						v1822 = v1854
																						v1845 = v1850
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2048 = v1850
																			}
																		}
																	} else {
																		v1860 = int64(0)
																		if v60 < int32(4) {
																			v1939 = v1718
																			v1940 = v60
																			v1945 = v1860
																		} else {
																			if v1718 != (v1718+int32(3))&int32(-4) {
																				v1939 = v1718
																				v1940 = v60
																				v1945 = v1860
																			} else {
																				v1869 = v60 - int32(4)
																				v1873 = int32(base.Ui32(v1869)>>(uint(int32(2))%32)) + int32(1)
																				v1875 = v1873 & int32(3)
																				if base.Ui32(v1869) < base.Ui32(int32(12)) {
																					v1911 = v1718
																					v1912 = v60
																					v1917 = v1860
																				} else {
																					v1881 = v1718
																					v1882 = v60
																					v1883 = int32(0)
																					v1887 = v1860
																					for {
																						v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+12))
																						v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+8))
																						v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+4))
																						v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1881)))
																						v1903 = base.I64_extend_i32_u(base.I32_popcnt(v1888)) + (base.I64_extend_i32_u(base.I32_popcnt(v1891)) + (base.I64_extend_i32_u(base.I32_popcnt(v1894)) + (v1887 + base.I64_extend_i32_u(base.I32_popcnt(v1897)))))
																						v1904 = int32(16)
																						v1905 = v1882 - v1904
																						v1907 = v1881 + v1904
																						v1909 = v1883 + int32(4)
																						if v1909 != v1873&int32(2147483644) {
																							v1881 = v1907
																							v1882 = v1905
																							v1883 = v1909
																							v1887 = v1903
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1911 = v1907
																					v1912 = v1905
																					v1917 = v1903
																				}
																				if v1875 == int32(0) {
																					v1939 = v1911
																					v1940 = v1912
																					v1945 = v1917
																				} else {
																					v1922 = v1912
																					v1923 = v1911
																					v1924 = int32(0)
																					v1927 = v1917
																					for {
																						v1928 = int32(4)
																						v1929 = v1922 - v1928
																						v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
																						v1933 = v1927 + base.I64_extend_i32_u(base.I32_popcnt(v1930))
																						v1935 = v1923 + v1928
																						v1937 = v1924 + int32(1)
																						if v1937 != v1875 {
																							v1922 = v1929
																							v1923 = v1935
																							v1924 = v1937
																							v1927 = v1933
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1939 = v1935
																					v1940 = v1929
																					v1945 = v1933
																				}
																			}
																		}
																		if v1940 == int32(0) {
																			v2018 = v1945
																		} else {
																			v1949 = v1940 & int32(3)
																			if v1949 == int32(0) {
																				v1972 = v1939
																				v1974 = v1940
																				v1978 = v1945
																			} else {
																				v1955 = v1940
																				v1956 = v1939
																				v1957 = int32(0)
																				v1959 = v1945
																				for {
																					v1960 = int32(1)
																					v1961 = v1955 - v1960
																					v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956))))
																					v1965 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+uint32(_consts[1053]))))
																					v1966 = v1959 + v1965
																					v1968 = v1956 + v1960
																					v1970 = v1957 + v1960
																					if v1970 != v1949 {
																						v1955 = v1961
																						v1956 = v1968
																						v1957 = v1970
																						v1959 = v1966
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1972 = v1968
																				v1974 = v1961
																				v1978 = v1966
																			}
																			if base.Ui32(v1940) < base.Ui32(int32(4)) {
																				v2018 = v1978
																			} else {
																				v1981 = v1972
																				v1983 = v1974
																				v1987 = v1978
																				for {
																					v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+3)))
																					v1991 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1988)+uint32(_consts[1053]))))
																					v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+2)))
																					v1995 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1992)+uint32(_consts[1053]))))
																					v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+1)))
																					v1999 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1996)+uint32(_consts[1053]))))
																					v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981))))
																					v2003 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2000)+uint32(_consts[1053]))))
																					v2007 = v1991 + (v1995 + (v1999 + (v1987 + v2003)))
																					v2008 = int32(4)
																					v2011 = v1983 - v2008
																					if v2011 != 0 {
																						v1981 = v1981 + v2008
																						v1983 = v2011
																						v1987 = v2007
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2018 = v2007
																			}
																		}
																		v2048 = v2018
																	}
																	v2057 = v1052 - base.I32_wrap_i64(v2048)
																}
															} else {
																if v1595&int32(1) == int32(0) {
																	if v1595&int32(1) == int32(0) {
																		v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v1718 = v1717
																	} else {
																		v1718 = v972
																	}
																	if v60 <= int32(3) {
																		if v60 == int32(0) {
																			v2048 = int64(0)
																		} else {
																			v1725 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																				v1728 = v1718
																				v1734 = int32(0)
																				v1757 = v1725
																				for {
																					v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+3)))
																					v1761 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+uint32(_consts[1053]))))
																					v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+2)))
																					v1765 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1762)+uint32(_consts[1053]))))
																					v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+1)))
																					v1769 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1766)+uint32(_consts[1053]))))
																					v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
																					v1773 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+uint32(_consts[1053]))))
																					v1777 = v1761 + (v1765 + (v1769 + (v1757 + v1773)))
																					v1778 = int32(4)
																					v1779 = v1728 + v1778
																					v1781 = v1734 + v1778
																					if v1781 != v1044 {
																						v1728 = v1779
																						v1734 = v1781
																						v1757 = v1777
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1783 = v1779
																				v1812 = v1777
																			} else {
																				v1783 = v1718
																				v1812 = v1725
																			}
																			v1813 = int32(0)
																			if v1042 == v1813 {
																				v2048 = v1812
																			} else {
																				v1816 = v1783
																				v1822 = v1813
																				v1845 = v1812
																				for {
																					v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816))))
																					v1849 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+uint32(_consts[1053]))))
																					v1850 = v1845 + v1849
																					v1851 = int32(1)
																					v1854 = v1822 + v1851
																					if v1854 != v1042 {
																						v1816 = v1816 + v1851
																						v1822 = v1854
																						v1845 = v1850
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2048 = v1850
																			}
																		}
																	} else {
																		v1860 = int64(0)
																		if v60 < int32(4) {
																			v1939 = v1718
																			v1940 = v60
																			v1945 = v1860
																		} else {
																			if v1718 != (v1718+int32(3))&int32(-4) {
																				v1939 = v1718
																				v1940 = v60
																				v1945 = v1860
																			} else {
																				v1869 = v60 - int32(4)
																				v1873 = int32(base.Ui32(v1869)>>(uint(int32(2))%32)) + int32(1)
																				v1875 = v1873 & int32(3)
																				if base.Ui32(v1869) < base.Ui32(int32(12)) {
																					v1911 = v1718
																					v1912 = v60
																					v1917 = v1860
																				} else {
																					v1881 = v1718
																					v1882 = v60
																					v1883 = int32(0)
																					v1887 = v1860
																					for {
																						v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+12))
																						v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+8))
																						v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+4))
																						v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1881)))
																						v1903 = base.I64_extend_i32_u(base.I32_popcnt(v1888)) + (base.I64_extend_i32_u(base.I32_popcnt(v1891)) + (base.I64_extend_i32_u(base.I32_popcnt(v1894)) + (v1887 + base.I64_extend_i32_u(base.I32_popcnt(v1897)))))
																						v1904 = int32(16)
																						v1905 = v1882 - v1904
																						v1907 = v1881 + v1904
																						v1909 = v1883 + int32(4)
																						if v1909 != v1873&int32(2147483644) {
																							v1881 = v1907
																							v1882 = v1905
																							v1883 = v1909
																							v1887 = v1903
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1911 = v1907
																					v1912 = v1905
																					v1917 = v1903
																				}
																				if v1875 == int32(0) {
																					v1939 = v1911
																					v1940 = v1912
																					v1945 = v1917
																				} else {
																					v1922 = v1912
																					v1923 = v1911
																					v1924 = int32(0)
																					v1927 = v1917
																					for {
																						v1928 = int32(4)
																						v1929 = v1922 - v1928
																						v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
																						v1933 = v1927 + base.I64_extend_i32_u(base.I32_popcnt(v1930))
																						v1935 = v1923 + v1928
																						v1937 = v1924 + int32(1)
																						if v1937 != v1875 {
																							v1922 = v1929
																							v1923 = v1935
																							v1924 = v1937
																							v1927 = v1933
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1939 = v1935
																					v1940 = v1929
																					v1945 = v1933
																				}
																			}
																		}
																		if v1940 == int32(0) {
																			v2018 = v1945
																		} else {
																			v1949 = v1940 & int32(3)
																			if v1949 == int32(0) {
																				v1972 = v1939
																				v1974 = v1940
																				v1978 = v1945
																			} else {
																				v1955 = v1940
																				v1956 = v1939
																				v1957 = int32(0)
																				v1959 = v1945
																				for {
																					v1960 = int32(1)
																					v1961 = v1955 - v1960
																					v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956))))
																					v1965 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+uint32(_consts[1053]))))
																					v1966 = v1959 + v1965
																					v1968 = v1956 + v1960
																					v1970 = v1957 + v1960
																					if v1970 != v1949 {
																						v1955 = v1961
																						v1956 = v1968
																						v1957 = v1970
																						v1959 = v1966
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1972 = v1968
																				v1974 = v1961
																				v1978 = v1966
																			}
																			if base.Ui32(v1940) < base.Ui32(int32(4)) {
																				v2018 = v1978
																			} else {
																				v1981 = v1972
																				v1983 = v1974
																				v1987 = v1978
																				for {
																					v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+3)))
																					v1991 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1988)+uint32(_consts[1053]))))
																					v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+2)))
																					v1995 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1992)+uint32(_consts[1053]))))
																					v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+1)))
																					v1999 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1996)+uint32(_consts[1053]))))
																					v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981))))
																					v2003 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2000)+uint32(_consts[1053]))))
																					v2007 = v1991 + (v1995 + (v1999 + (v1987 + v2003)))
																					v2008 = int32(4)
																					v2011 = v1983 - v2008
																					if v2011 != 0 {
																						v1981 = v1981 + v2008
																						v1983 = v2011
																						v1987 = v2007
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2018 = v2007
																			}
																		}
																		v2048 = v2018
																	}
																	v2057 = v1052 - base.I32_wrap_i64(v2048)
																} else {
																	v2057 = int32(0)
																}
															}
															v2083 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
															v2084 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
															v2085 = v2083 - v2084
															if base.F64_lt(base.F64_convert_i32_s(v1572), base.F64_add(base.F64_convert_i32_s(v2057), base.F64_mul(base.F64_convert_i32_s(v2085*v2085*v2085), float64(-0.1)))) != 0 {
																v2093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+4)))
																if v2093&int32(4) != 0 {
																} else {
																	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																	if v2096 == int32(1) {
																		v2101 = F__emscripten_memset_bulkmem(m, v974, base.I32_extend8_s(int32(255)), v60)
																		mBase = m.M
																	} else {
																		if v60 <= int32(0) {
																		} else {
																			v2104 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																			v2105 = int32(0)
																			if base.Ui32(int32(4)) <= base.Ui32(v60) {
																				v2110 = v2105
																				v2120 = v2105
																				for {
																					v2140 = v2110 + v974
																					v2141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140))))
																					v2143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2110+v2104))))
																					v2144 = v2141 | v2143
																					*(*uint8)(unsafe.Add(mBase, uint32(v2140))) = uint8(v2144)
																					v2147 = v2110 | int32(1)
																					v2148 = v974 + v2147
																					v2149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2148))))
																					v2151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2147))))
																					v2152 = v2149 | v2151
																					*(*uint8)(unsafe.Add(mBase, uint32(v2148))) = uint8(v2152)
																					v2155 = v2110 | int32(2)
																					v2156 = v974 + v2155
																					v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156))))
																					v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2155))))
																					v2160 = v2157 | v2159
																					*(*uint8)(unsafe.Add(mBase, uint32(v2156))) = uint8(v2160)
																					v2163 = v2110 | int32(3)
																					v2164 = v974 + v2163
																					v2165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2164))))
																					v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2163))))
																					v2168 = v2165 | v2167
																					*(*uint8)(unsafe.Add(mBase, uint32(v2164))) = uint8(v2168)
																					v2170 = int32(4)
																					v2171 = v2110 + v2170
																					v2173 = v2120 + v2170
																					if v2173 != v1040 {
																						v2110 = v2171
																						v2120 = v2173
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2175 = v2171
																			} else {
																				v2175 = v2105
																			}
																			if v1042 == int32(0) {
																			} else {
																				v2207 = v2175
																				v2213 = v2105
																				for {
																					v2237 = v2207 + v974
																					v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2237))))
																					v2240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2207+v2104))))
																					v2241 = v2238 | v2240
																					*(*uint8)(unsafe.Add(mBase, uint32(v2237))) = uint8(v2241)
																					v2243 = int32(1)
																					v2246 = v2213 + v2243
																					if v2246 != v1042 {
																						v2207 = v2207 + v2243
																						v2213 = v2246
																						continue
																					} else {
																						break
																					}
																					break
																				}
																			}
																		}
																	}
																}
																*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v1089)
																v2279 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v2279 + int32(1)
																v2522 = v1071 + int32(2)
																v2527 = v1076
															} else {
																v2285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+4)))
																if v2285&int32(4) != 0 {
																} else {
																	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																	if v2288 == int32(1) {
																		v2293 = F__emscripten_memset_bulkmem(m, v972, base.I32_extend8_s(int32(255)), v60)
																		mBase = m.M
																	} else {
																		if v60 <= int32(0) {
																		} else {
																			v2296 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																			v2297 = int32(0)
																			if base.Ui32(int32(4)) <= base.Ui32(v60) {
																				v2302 = v2297
																				v2312 = v2297
																				for {
																					v2332 = v2302 + v972
																					v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2332))))
																					v2335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302+v2296))))
																					v2336 = v2333 | v2335
																					*(*uint8)(unsafe.Add(mBase, uint32(v2332))) = uint8(v2336)
																					v2339 = v2302 | int32(1)
																					v2340 = v972 + v2339
																					v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2340))))
																					v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2339))))
																					v2344 = v2341 | v2343
																					*(*uint8)(unsafe.Add(mBase, uint32(v2340))) = uint8(v2344)
																					v2347 = v2302 | int32(2)
																					v2348 = v972 + v2347
																					v2349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2348))))
																					v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2347))))
																					v2352 = v2349 | v2351
																					*(*uint8)(unsafe.Add(mBase, uint32(v2348))) = uint8(v2352)
																					v2355 = v2302 | int32(3)
																					v2356 = v972 + v2355
																					v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2356))))
																					v2359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2355))))
																					v2360 = v2357 | v2359
																					*(*uint8)(unsafe.Add(mBase, uint32(v2356))) = uint8(v2360)
																					v2362 = int32(4)
																					v2363 = v2302 + v2362
																					v2365 = v2312 + v2362
																					if v2365 != v1040 {
																						v2302 = v2363
																						v2312 = v2365
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v2367 = v2363
																			} else {
																				v2367 = v2297
																			}
																			if v1042 == int32(0) {
																			} else {
																				v2399 = v2367
																				v2405 = v2297
																				for {
																					v2429 = v2399 + v972
																					v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2429))))
																					v2432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2399+v2296))))
																					v2433 = v2430 | v2432
																					*(*uint8)(unsafe.Add(mBase, uint32(v2429))) = uint8(v2433)
																					v2435 = int32(1)
																					v2438 = v2405 + v2435
																					if v2438 != v1042 {
																						v2399 = v2399 + v2435
																						v2405 = v2438
																						continue
																					} else {
																						break
																					}
																					break
																				}
																			}
																		}
																	}
																}
																*(*uint16)(unsafe.Add(mBase, uint32(v1076))) = uint16(v1089)
																v2471 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
																*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v2471 + int32(1)
																v2522 = v1071
																v2527 = v1076 + int32(2)
															}
														}
													}
													v2538 = v1060 + int32(1)
													if v2538 != v1038 {
														v1060 = v2538
														v1071 = v2522
														v1076 = v2527
														continue
													} else {
														break
													}
													break
												}
												v2555 = v2522
												v2560 = v2527
												v2570 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v2555))) = uint16(v2570)
												*(*uint16)(unsafe.Add(mBase, uint32(v2560))) = uint16(v2570)
												*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v727
												*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v701
												return v31
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
	} else {
		v60 = int32(124)
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v63 = v61 + int32(65534)
		v65 = v63 & int32(65535)
		v67 = v65 + int32(2)
		v69 = v67 << (uint(int32(1)) % 32)
		v70 = F_palloc(m, v69)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v70
			v73 = F_palloc(m, v69)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v73
				v79 = F_palloc(m, v67<<(uint(int32(3))%32))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v82 = F_palloc(m, v60*v67)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = int32(0)
						v85 = v2
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v79+v84<<(uint(int32(3))%32))+4)) = v82 + v84*v60
							v121 = v85 + int32(1)
							v123 = v121 & int32(65535)
							if base.Ui32(v123) < base.Ui32(v67) {
								v84 = v123
								v85 = v121
								continue
							} else {
								break
							}
							break
						}
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
						v126 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)) = uint8(v126)
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
						if v128&int32(1) != 0 {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
							v135 = int32(base.Ui32(v131)>>(uint(int32(2))%32)) - int32(8)
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
							if v136&int32(3) != 0 {
								v158 = v60
								v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
								mBase = m.M
							} else {
								if base.Ui32(int32(1024)) < base.Ui32(v60) {
									v158 = v60
									v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
									mBase = m.M
								} else {
									if v60&int32(3) != 0 {
										v158 = v60
										v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
										mBase = m.M
									} else {
										v143 = v136 + v60
										if base.Ui32(v143) <= base.Ui32(v136) {
										} else {
											v148 = v136 + int32(4)
											if base.Ui32(v148) < base.Ui32(v143) {
												v150 = v143
											} else {
												v150 = v148
											}
											v158 = (v136^int32(-1)+v150)&int32(-4) + int32(4)
											v161 = F__emscripten_memset_bulkmem(m, v136, base.I32_extend8_s(int32(0)), v158)
											mBase = m.M
										}
									}
								}
							}
							if base.Ui32(v135) < base.Ui32(int32(4)) {
							} else {
								v167 = v125 + int32(8)
								v169 = v60 << (uint(int32(3)) % 32)
								v170 = int32(1)
								v172 = int32(base.Ui32(v135) >> (uint(int32(2)) % 32))
								if base.Ui32(v172) <= base.Ui32(v170) {
									v175 = v170
								} else {
									v175 = v172
								}
								v178 = int32(0)
								if base.Ui32(int32(8)) <= base.Ui32(v135) {
									v184 = v178
									v188 = int32(0)
									for {
										v214 = int32(2)
										v216 = v167 + v184<<(uint(v214)%32)
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
										v218 = base.I32_rem_u_s(v217, v169)
										v219 = int32(3)
										v221 = v136 + int32(base.Ui32(v218)>>(uint(v219)%32))
										v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
										v223 = int32(1)
										v224 = int32(7)
										v227 = v222 | v223<<(uint(v218&v224)%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v227)
										v229 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
										v230 = base.I32_rem_u_s(v229, v169)
										v233 = v136 + int32(base.Ui32(v230)>>(uint(v219)%32))
										v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
										v239 = v234 | v223<<(uint(v230&v224)%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v239)
										v242 = v184 + v214
										v244 = v188 + v214
										if v244 != v175&int32(1073741822) {
											v184 = v242
											v188 = v244
											continue
										} else {
											break
										}
										break
									}
									v246 = v242
								} else {
									v246 = v178
								}
								if v175&int32(1) == int32(0) {
								} else {
									v281 = *(*int32)(unsafe.Add(mBase, uint32(v167+v246<<(uint(int32(2))%32))))
									v282 = base.I32_rem_u_s(v281, v169)
									v285 = v136 + int32(base.Ui32(v282)>>(uint(int32(3))%32))
									v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
									v291 = v286 | int32(1)<<(uint(v282&int32(7))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v291)
								}
							}
						} else {
							if v128&int32(4) != 0 {
								v295 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)) = uint8(v295)
							} else {
								v297 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								if v60 != 0 {
									v300 = F__emscripten_memcpy_bulkmem(m, v297, v125+int32(8), v60)
									mBase = m.M
								} else {
								}
							}
						}
						v333 = v32 + int32(4)
						if base.Ui32(int32(2)) <= base.Ui32(v65) {
							v336 = int32(3)
							v344 = v60 << (uint(v336) % 32)
							v362 = int32(-1)
							v363 = int32(1)
							v364 = v2
							v366 = v2
							for {
								v381 = v363 + int32(1)
								v382 = v381
								v388 = v381
								v397 = v362
								v399 = v364
								v401 = v366
								for {
									if v363 != int32(1) {
									} else {
										v417 = *(*int32)(unsafe.Add(mBase, uint32(v333+v388<<(uint(int32(4))%32))))
										v420 = v79 + v388<<(uint(int32(3))%32)
										v421 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v421)
										v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
										if v423&int32(1) != 0 {
											v426 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
											v430 = int32(base.Ui32(v426)>>(uint(int32(2))%32)) - int32(8)
											v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
											v434 = int32(0)
											if (base.B2i32(v431&int32(3) != v434)|(base.B2i32(v60&v336 != int32(0))|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v60))))&int32(1) == v434 {
												v441 = v431 + v60
												if base.Ui32(v441) <= base.Ui32(v431) {
												} else {
													v446 = v431 + int32(4)
													if base.Ui32(v446) < base.Ui32(v441) {
														v448 = v441
													} else {
														v448 = v446
													}
													v454 = (v431^int32(-1)+v448)&int32(-4) + int32(4)
													v458 = F__emscripten_memset_bulkmem(m, v431, base.I32_extend8_s(int32(0)), v454)
													mBase = m.M
												}
											} else {
												v454 = v60
												v458 = F__emscripten_memset_bulkmem(m, v431, base.I32_extend8_s(int32(0)), v454)
												mBase = m.M
											}
											if base.Ui32(v430) < base.Ui32(int32(4)) {
											} else {
												v464 = v417 + int32(8)
												v465 = int32(1)
												v467 = int32(base.Ui32(v430) >> (uint(int32(2)) % 32))
												if base.Ui32(v467) <= base.Ui32(v465) {
													v470 = v465
												} else {
													v470 = v467
												}
												v473 = int32(0)
												if base.Ui32(int32(8)) <= base.Ui32(v430) {
													v479 = v473
													v483 = int32(0)
													for {
														v509 = int32(2)
														v511 = v464 + v479<<(uint(v509)%32)
														v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
														v513 = base.I32_rem_u_s(v512, v344)
														v514 = int32(3)
														v516 = v431 + int32(base.Ui32(v513)>>(uint(v514)%32))
														v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
														v518 = int32(1)
														v519 = int32(7)
														v522 = v517 | v518<<(uint(v513&v519)%32)
														*(*uint8)(unsafe.Add(mBase, uint32(v516))) = uint8(v522)
														v524 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
														v525 = base.I32_rem_u_s(v524, v344)
														v528 = v431 + int32(base.Ui32(v525)>>(uint(v514)%32))
														v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
														v534 = v529 | v518<<(uint(v525&v519)%32)
														*(*uint8)(unsafe.Add(mBase, uint32(v528))) = uint8(v534)
														v537 = v479 + v509
														v539 = v483 + v509
														if v539 != v470&int32(1073741822) {
															v479 = v537
															v483 = v539
															continue
														} else {
															break
														}
														break
													}
													v541 = v537
												} else {
													v541 = v473
												}
												if v470&int32(1) == int32(0) {
												} else {
													v576 = *(*int32)(unsafe.Add(mBase, uint32(v464+v541<<(uint(int32(2))%32))))
													v577 = base.I32_rem_u_s(v576, v344)
													v580 = v431 + int32(base.Ui32(v577)>>(uint(int32(3))%32))
													v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
													v586 = v581 | int32(1)<<(uint(v577&int32(7))%32)
													*(*uint8)(unsafe.Add(mBase, uint32(v580))) = uint8(v586)
												}
											}
										} else {
											if v423&int32(4) != 0 {
												v590 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v420))) = uint8(v590)
											} else {
												v592 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
												if v60 != 0 {
													v595 = F__emscripten_memcpy_bulkmem(m, v592, v417+int32(8), v60)
													mBase = m.M
												} else {
												}
											}
										}
									}
									v630 = F_hemdistcache_1(m, v79+v388<<(uint(int32(3))%32), v79+v363<<(uint(int32(3))%32), v60)
									mBase = m.M
									v631 = base.B2i32(v397 < v630)
									if v397 < v630 {
										v632 = v630
									} else {
										v632 = v397
									}
									if v397 < v630 {
										v633 = v382
									} else {
										v633 = v401
									}
									if v397 < v630 {
										v634 = v363
									} else {
										v634 = v399
									}
									v636 = v382 + int32(1)
									v637 = int32(65535)
									v638 = v636 & v637
									if base.Ui32(v638) <= base.Ui32(v63&v637) {
										v382 = v636
										v388 = v638
										v397 = v632
										v399 = v634
										v401 = v633
										continue
									} else {
										break
									}
									break
								}
								if v381 != v65 {
									v362 = v632
									v363 = v381
									v364 = v634
									v366 = v633
									continue
								} else {
									break
								}
								break
							}
							v660 = v634
							v662 = v633
						} else {
							v660 = v2
							v662 = v2
						}
						v673 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v673
						*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v673
						v678 = int32(65535)
						v686 = base.B2i32(v660&v678 == v673) | base.B2i32(v662&v678 == v673)
						if v686 != 0 {
							v687 = int32(1)
						} else {
							v687 = v660
						}
						v692 = v79 + v687&int32(65535)<<(uint(int32(3))%32)
						v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
						v694 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						v695 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
						v696 = int32(8)
						v698 = v60 + v696
						v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
						if v699 != 0 {
							v700 = v696
						} else {
							v700 = v698
						}
						v701 = F_palloc(m, v700)
						mBase = m.M
						v702 = m.ExcPending
						if v702 != 0 {
							return int32(0)
						} else {
							if v699 != 0 {
								v705 = int32(6)
							} else {
								v705 = int32(2)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v705
							v707 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v701))) = v700 << (uint(v707) % 32)
							if v686 != 0 {
								v711 = v707
							} else {
								v711 = v662
							}
							if v693 == int32(0) {
							} else {
								if v699 != 0 {
								} else {
									if v60 != 0 {
										v716 = F__emscripten_memcpy_bulkmem(m, v701+int32(8), v693, v60)
										mBase = m.M
									} else {
									}
								}
							}
							v722 = v79 + v711&int32(65535)<<(uint(int32(3))%32)
							v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
							v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722))))
							if v725 != 0 {
								v726 = int32(8)
							} else {
								v726 = v698
							}
							v727 = F_palloc(m, v726)
							mBase = m.M
							v728 = m.ExcPending
							if v728 != 0 {
								return int32(0)
							} else {
								if v725 != 0 {
									v731 = int32(6)
								} else {
									v731 = int32(2)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v731
								*(*int32)(unsafe.Add(mBase, uint32(v727))) = v726 << (uint(int32(2)) % 32)
								if v723 == int32(0) {
								} else {
									if v725 != 0 {
									} else {
										if v60 != 0 {
											v740 = F__emscripten_memcpy_bulkmem(m, v727+int32(8), v723, v60)
											mBase = m.M
										} else {
										}
									}
								}
								v742 = int32(65535)
								v743 = v61 + v742
								v745 = v743 & v742
								v749 = *(*int32)(unsafe.Add(mBase, uint32(v333+v745<<(uint(int32(4))%32))))
								v752 = v79 + v745<<(uint(int32(3))%32)
								v753 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v752))) = uint8(v753)
								v755 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
								if v755&int32(1) != 0 {
									v758 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
									v762 = int32(base.Ui32(v758)>>(uint(int32(2))%32)) - int32(8)
									v763 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
									if v763&int32(3) != 0 {
										v785 = v60
										v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
										mBase = m.M
									} else {
										if base.Ui32(int32(1024)) < base.Ui32(v60) {
											v785 = v60
											v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
											mBase = m.M
										} else {
											if v60&int32(3) != 0 {
												v785 = v60
												v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
												mBase = m.M
											} else {
												v770 = v763 + v60
												if base.Ui32(v770) <= base.Ui32(v763) {
												} else {
													v775 = v763 + int32(4)
													if base.Ui32(v775) < base.Ui32(v770) {
														v777 = v770
													} else {
														v777 = v775
													}
													v785 = (v763^int32(-1)+v777)&int32(-4) + int32(4)
													v788 = F__emscripten_memset_bulkmem(m, v763, base.I32_extend8_s(int32(0)), v785)
													mBase = m.M
												}
											}
										}
									}
									if base.Ui32(v762) < base.Ui32(int32(4)) {
									} else {
										v794 = v749 + int32(8)
										v796 = v60 << (uint(int32(3)) % 32)
										v797 = int32(1)
										v799 = int32(base.Ui32(v762) >> (uint(int32(2)) % 32))
										if base.Ui32(v799) <= base.Ui32(v797) {
											v802 = v797
										} else {
											v802 = v799
										}
										v805 = int32(0)
										if base.Ui32(int32(8)) <= base.Ui32(v762) {
											v811 = v805
											v815 = int32(0)
											for {
												v841 = int32(2)
												v843 = v794 + v811<<(uint(v841)%32)
												v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
												v845 = base.I32_rem_u_s(v844, v796)
												v846 = int32(3)
												v848 = v763 + int32(base.Ui32(v845)>>(uint(v846)%32))
												v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
												v850 = int32(1)
												v851 = int32(7)
												v854 = v849 | v850<<(uint(v845&v851)%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v848))) = uint8(v854)
												v856 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
												v857 = base.I32_rem_u_s(v856, v796)
												v860 = v763 + int32(base.Ui32(v857)>>(uint(v846)%32))
												v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860))))
												v866 = v861 | v850<<(uint(v857&v851)%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v860))) = uint8(v866)
												v869 = v811 + v841
												v871 = v815 + v841
												if v871 != v802&int32(1073741822) {
													v811 = v869
													v815 = v871
													continue
												} else {
													break
												}
												break
											}
											v873 = v869
										} else {
											v873 = v805
										}
										if v802&int32(1) == int32(0) {
										} else {
											v908 = *(*int32)(unsafe.Add(mBase, uint32(v794+v873<<(uint(int32(2))%32))))
											v909 = base.I32_rem_u_s(v908, v796)
											v912 = v763 + int32(base.Ui32(v909)>>(uint(int32(3))%32))
											v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
											v918 = v913 | int32(1)<<(uint(v909&int32(7))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v918)
										}
									}
								} else {
									if v755&int32(4) != 0 {
										v922 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v752))) = uint8(v922)
									} else {
										v924 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
										if v60 != 0 {
											v927 = F__emscripten_memcpy_bulkmem(m, v924, v749+int32(8), v60)
											mBase = m.M
										} else {
										}
									}
								}
								v961 = F_palloc(m, v745<<(uint(int32(3))%32))
								mBase = m.M
								v962 = m.ExcPending
								if v962 != 0 {
									return int32(0)
								} else {
									if v61&int32(65535) == int32(1) {
										F_pg_qsort(m, v961, v745, int32(8), int32(1522))
										mBase = m.M
										v970 = m.ExcPending
										if v970 != 0 {
											return int32(0)
										} else {
											v2555 = v694
											v2560 = v695
											v2570 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v2555))) = uint16(v2570)
											*(*uint16)(unsafe.Add(mBase, uint32(v2560))) = uint16(v2570)
											*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v727
											*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v701
											return v31
										}
									} else {
										v971 = int32(8)
										v972 = v727 + v971
										v974 = v701 + v971
										v975 = int32(1)
										v977 = v975
										v981 = v975
										for {
											v1008 = v981 << (uint(int32(3)) % 32)
											v1009 = v961 + v1008
											*(*uint16)(unsafe.Add(mBase, uint32(v1009-int32(8)))) = uint16(v977)
											v1015 = v1008 + v79
											v1016 = F_hemdistcache_1(m, v692, v1015, v60)
											mBase = m.M
											v1017 = F_hemdistcache_1(m, v722, v1015, v60)
											mBase = m.M
											v1018 = v1016 - v1017
											v1020 = v1018 >> (uint(int32(31)) % 32)
											*(*int32)(unsafe.Add(mBase, uint32(v1009-int32(4)))) = v1018 ^ v1020 - v1020
											v1025 = v977 + int32(1)
											v1026 = int32(65535)
											v1027 = v1025 & v1026
											if base.Ui32(v1027) <= base.Ui32(v743&v1026) {
												v977 = v1025
												v981 = v1027
												continue
											} else {
												break
											}
											break
										}
										F_pg_qsort(m, v961, v745, int32(8), int32(1522))
										mBase = m.M
										v1034 = m.ExcPending
										if v1034 != 0 {
											return int32(0)
										} else {
											v1035 = int32(1)
											if base.Ui32(v745) <= base.Ui32(v1035) {
												v1038 = v1035
											} else {
												v1038 = v745
											}
											v1040 = v60 & int32(2147483644)
											v1041 = int32(3)
											v1042 = v60 & v1041
											v1044 = v60 & int32(-4)
											v1046 = v60 & int32(2147483646)
											v1047 = int32(1)
											v1048 = v60 & v1047
											v1050 = v60 - v1047
											v1052 = v60 << (uint(v1041) % 32)
											v1060 = int32(0)
											v1071 = v694
											v1076 = v695
											for {
												v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961+v1060<<(uint(int32(3))%32)))))
												if v687&int32(65535) == v1089 {
													*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v687)
													v1092 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1092 + int32(1)
													v2522 = v1071 + int32(2)
													v2527 = v1076
												} else {
													if v711&int32(65535) == v1089 {
														*(*uint16)(unsafe.Add(mBase, uint32(v1076))) = uint16(v711)
														v1102 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1102 + int32(1)
														v2522 = v1071
														v2527 = v1076 + int32(2)
													} else {
														v1108 = v79 + v1089<<(uint(int32(3))%32)
														v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
														v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+4)))
														if v1110&int32(4) == int32(0) {
															if v1109&int32(1) == int32(0) {
																if v60 <= int32(0) {
																	v1572 = int32(0)
																} else {
																	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1128 = int32(0)
																	if v1050 != 0 {
																		v1131 = v1128
																		v1138 = v1128
																		v1141 = v1128
																		for {
																			v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131+v974))))
																			v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131+v1127))))
																			v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162^v1164)+uint32(_consts[1053]))))
																			v1171 = v1131 | int32(1)
																			v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974+v1171))))
																			v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127+v1171))))
																			v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173^v1175)+uint32(_consts[1053]))))
																			v1180 = v1138 + v1168 + v1179
																			v1181 = int32(2)
																			v1182 = v1131 + v1181
																			v1184 = v1141 + v1181
																			if v1184 != v1046 {
																				v1131 = v1182
																				v1138 = v1180
																				v1141 = v1184
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v1186 = v1182
																		v1193 = v1180
																	} else {
																		v1186 = v1128
																		v1193 = v1128
																	}
																	if v1048 == int32(0) {
																		v1572 = v1193
																	} else {
																		v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v974))))
																		v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v1127))))
																		v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219^v1221)+uint32(_consts[1053]))))
																		v1572 = v1193 + v1225
																	}
																}
															} else {
																if v1109&int32(1) == int32(0) {
																	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1232 = v1231
																} else {
																	v1232 = v974
																}
																if v60 <= int32(3) {
																	if v60 == int32(0) {
																		v1562 = int64(0)
																	} else {
																		v1239 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																			v1242 = v1232
																			v1249 = int32(0)
																			v1271 = v1239
																			for {
																				v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+3)))
																				v1275 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+uint32(_consts[1053]))))
																				v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+2)))
																				v1279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+uint32(_consts[1053]))))
																				v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
																				v1283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+uint32(_consts[1053]))))
																				v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
																				v1287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1284)+uint32(_consts[1053]))))
																				v1291 = v1275 + (v1279 + (v1283 + (v1271 + v1287)))
																				v1292 = int32(4)
																				v1293 = v1242 + v1292
																				v1295 = v1249 + v1292
																				if v1295 != v1044 {
																					v1242 = v1293
																					v1249 = v1295
																					v1271 = v1291
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1297 = v1293
																			v1326 = v1291
																		} else {
																			v1297 = v1232
																			v1326 = v1239
																		}
																		v1327 = int32(0)
																		if v1042 == v1327 {
																			v1562 = v1326
																		} else {
																			v1330 = v1297
																			v1337 = v1327
																			v1359 = v1326
																			for {
																				v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330))))
																				v1363 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+uint32(_consts[1053]))))
																				v1364 = v1359 + v1363
																				v1365 = int32(1)
																				v1368 = v1337 + v1365
																				if v1368 != v1042 {
																					v1330 = v1330 + v1365
																					v1337 = v1368
																					v1359 = v1364
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1562 = v1364
																		}
																	}
																} else {
																	v1374 = int64(0)
																	if v60 < int32(4) {
																		v1453 = v1232
																		v1454 = v60
																		v1459 = v1374
																	} else {
																		if v1232 != (v1232+int32(3))&int32(-4) {
																			v1453 = v1232
																			v1454 = v60
																			v1459 = v1374
																		} else {
																			v1383 = v60 - int32(4)
																			v1387 = int32(base.Ui32(v1383)>>(uint(int32(2))%32)) + int32(1)
																			v1389 = v1387 & int32(3)
																			if base.Ui32(v1383) < base.Ui32(int32(12)) {
																				v1425 = v1232
																				v1426 = v60
																				v1431 = v1374
																			} else {
																				v1395 = v1232
																				v1396 = v60
																				v1397 = int32(0)
																				v1401 = v1374
																				for {
																					v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
																					v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+8))
																					v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
																					v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1395)))
																					v1417 = base.I64_extend_i32_u(base.I32_popcnt(v1402)) + (base.I64_extend_i32_u(base.I32_popcnt(v1405)) + (base.I64_extend_i32_u(base.I32_popcnt(v1408)) + (v1401 + base.I64_extend_i32_u(base.I32_popcnt(v1411)))))
																					v1418 = int32(16)
																					v1419 = v1396 - v1418
																					v1421 = v1395 + v1418
																					v1423 = v1397 + int32(4)
																					if v1423 != v1387&int32(2147483644) {
																						v1395 = v1421
																						v1396 = v1419
																						v1397 = v1423
																						v1401 = v1417
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1425 = v1421
																				v1426 = v1419
																				v1431 = v1417
																			}
																			if v1389 == int32(0) {
																				v1453 = v1425
																				v1454 = v1426
																				v1459 = v1431
																			} else {
																				v1436 = v1426
																				v1437 = v1425
																				v1438 = int32(0)
																				v1441 = v1431
																				for {
																					v1442 = int32(4)
																					v1443 = v1436 - v1442
																					v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
																					v1447 = v1441 + base.I64_extend_i32_u(base.I32_popcnt(v1444))
																					v1449 = v1437 + v1442
																					v1451 = v1438 + int32(1)
																					if v1451 != v1389 {
																						v1436 = v1443
																						v1437 = v1449
																						v1438 = v1451
																						v1441 = v1447
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1453 = v1449
																				v1454 = v1443
																				v1459 = v1447
																			}
																		}
																	}
																	if v1454 == int32(0) {
																		v1532 = v1459
																	} else {
																		v1463 = v1454 & int32(3)
																		if v1463 == int32(0) {
																			v1486 = v1453
																			v1488 = v1454
																			v1492 = v1459
																		} else {
																			v1469 = v1454
																			v1470 = v1453
																			v1471 = int32(0)
																			v1473 = v1459
																			for {
																				v1474 = int32(1)
																				v1475 = v1469 - v1474
																				v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
																				v1479 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+uint32(_consts[1053]))))
																				v1480 = v1473 + v1479
																				v1482 = v1470 + v1474
																				v1484 = v1471 + v1474
																				if v1484 != v1463 {
																					v1469 = v1475
																					v1470 = v1482
																					v1471 = v1484
																					v1473 = v1480
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1486 = v1482
																			v1488 = v1475
																			v1492 = v1480
																		}
																		if base.Ui32(v1454) < base.Ui32(int32(4)) {
																			v1532 = v1492
																		} else {
																			v1495 = v1486
																			v1497 = v1488
																			v1501 = v1492
																			for {
																				v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+3)))
																				v1505 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1502)+uint32(_consts[1053]))))
																				v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+2)))
																				v1509 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_consts[1053]))))
																				v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+1)))
																				v1513 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+uint32(_consts[1053]))))
																				v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495))))
																				v1517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1514)+uint32(_consts[1053]))))
																				v1521 = v1505 + (v1509 + (v1513 + (v1501 + v1517)))
																				v1522 = int32(4)
																				v1525 = v1497 - v1522
																				if v1525 != 0 {
																					v1495 = v1495 + v1522
																					v1497 = v1525
																					v1501 = v1521
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1532 = v1521
																		}
																	}
																	v1562 = v1532
																}
																v1572 = v1052 - base.I32_wrap_i64(v1562)
															}
														} else {
															if v1109&int32(1) == int32(0) {
																if v1109&int32(1) == int32(0) {
																	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1232 = v1231
																} else {
																	v1232 = v974
																}
																if v60 <= int32(3) {
																	if v60 == int32(0) {
																		v1562 = int64(0)
																	} else {
																		v1239 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																			v1242 = v1232
																			v1249 = int32(0)
																			v1271 = v1239
																			for {
																				v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+3)))
																				v1275 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+uint32(_consts[1053]))))
																				v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+2)))
																				v1279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+uint32(_consts[1053]))))
																				v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
																				v1283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+uint32(_consts[1053]))))
																				v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
																				v1287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1284)+uint32(_consts[1053]))))
																				v1291 = v1275 + (v1279 + (v1283 + (v1271 + v1287)))
																				v1292 = int32(4)
																				v1293 = v1242 + v1292
																				v1295 = v1249 + v1292
																				if v1295 != v1044 {
																					v1242 = v1293
																					v1249 = v1295
																					v1271 = v1291
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1297 = v1293
																			v1326 = v1291
																		} else {
																			v1297 = v1232
																			v1326 = v1239
																		}
																		v1327 = int32(0)
																		if v1042 == v1327 {
																			v1562 = v1326
																		} else {
																			v1330 = v1297
																			v1337 = v1327
																			v1359 = v1326
																			for {
																				v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330))))
																				v1363 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+uint32(_consts[1053]))))
																				v1364 = v1359 + v1363
																				v1365 = int32(1)
																				v1368 = v1337 + v1365
																				if v1368 != v1042 {
																					v1330 = v1330 + v1365
																					v1337 = v1368
																					v1359 = v1364
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1562 = v1364
																		}
																	}
																} else {
																	v1374 = int64(0)
																	if v60 < int32(4) {
																		v1453 = v1232
																		v1454 = v60
																		v1459 = v1374
																	} else {
																		if v1232 != (v1232+int32(3))&int32(-4) {
																			v1453 = v1232
																			v1454 = v60
																			v1459 = v1374
																		} else {
																			v1383 = v60 - int32(4)
																			v1387 = int32(base.Ui32(v1383)>>(uint(int32(2))%32)) + int32(1)
																			v1389 = v1387 & int32(3)
																			if base.Ui32(v1383) < base.Ui32(int32(12)) {
																				v1425 = v1232
																				v1426 = v60
																				v1431 = v1374
																			} else {
																				v1395 = v1232
																				v1396 = v60
																				v1397 = int32(0)
																				v1401 = v1374
																				for {
																					v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
																					v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+8))
																					v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
																					v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1395)))
																					v1417 = base.I64_extend_i32_u(base.I32_popcnt(v1402)) + (base.I64_extend_i32_u(base.I32_popcnt(v1405)) + (base.I64_extend_i32_u(base.I32_popcnt(v1408)) + (v1401 + base.I64_extend_i32_u(base.I32_popcnt(v1411)))))
																					v1418 = int32(16)
																					v1419 = v1396 - v1418
																					v1421 = v1395 + v1418
																					v1423 = v1397 + int32(4)
																					if v1423 != v1387&int32(2147483644) {
																						v1395 = v1421
																						v1396 = v1419
																						v1397 = v1423
																						v1401 = v1417
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1425 = v1421
																				v1426 = v1419
																				v1431 = v1417
																			}
																			if v1389 == int32(0) {
																				v1453 = v1425
																				v1454 = v1426
																				v1459 = v1431
																			} else {
																				v1436 = v1426
																				v1437 = v1425
																				v1438 = int32(0)
																				v1441 = v1431
																				for {
																					v1442 = int32(4)
																					v1443 = v1436 - v1442
																					v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1437)))
																					v1447 = v1441 + base.I64_extend_i32_u(base.I32_popcnt(v1444))
																					v1449 = v1437 + v1442
																					v1451 = v1438 + int32(1)
																					if v1451 != v1389 {
																						v1436 = v1443
																						v1437 = v1449
																						v1438 = v1451
																						v1441 = v1447
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1453 = v1449
																				v1454 = v1443
																				v1459 = v1447
																			}
																		}
																	}
																	if v1454 == int32(0) {
																		v1532 = v1459
																	} else {
																		v1463 = v1454 & int32(3)
																		if v1463 == int32(0) {
																			v1486 = v1453
																			v1488 = v1454
																			v1492 = v1459
																		} else {
																			v1469 = v1454
																			v1470 = v1453
																			v1471 = int32(0)
																			v1473 = v1459
																			for {
																				v1474 = int32(1)
																				v1475 = v1469 - v1474
																				v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
																				v1479 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+uint32(_consts[1053]))))
																				v1480 = v1473 + v1479
																				v1482 = v1470 + v1474
																				v1484 = v1471 + v1474
																				if v1484 != v1463 {
																					v1469 = v1475
																					v1470 = v1482
																					v1471 = v1484
																					v1473 = v1480
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1486 = v1482
																			v1488 = v1475
																			v1492 = v1480
																		}
																		if base.Ui32(v1454) < base.Ui32(int32(4)) {
																			v1532 = v1492
																		} else {
																			v1495 = v1486
																			v1497 = v1488
																			v1501 = v1492
																			for {
																				v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+3)))
																				v1505 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1502)+uint32(_consts[1053]))))
																				v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+2)))
																				v1509 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_consts[1053]))))
																				v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495)+1)))
																				v1513 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+uint32(_consts[1053]))))
																				v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1495))))
																				v1517 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1514)+uint32(_consts[1053]))))
																				v1521 = v1505 + (v1509 + (v1513 + (v1501 + v1517)))
																				v1522 = int32(4)
																				v1525 = v1497 - v1522
																				if v1525 != 0 {
																					v1495 = v1495 + v1522
																					v1497 = v1525
																					v1501 = v1521
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1532 = v1521
																		}
																	}
																	v1562 = v1532
																}
																v1572 = v1052 - base.I32_wrap_i64(v1562)
															} else {
																v1572 = int32(0)
															}
														}
														v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
														v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+4)))
														if v1596&int32(4) == int32(0) {
															if v1595&int32(1) == int32(0) {
																if v60 <= int32(0) {
																	v2057 = int32(0)
																} else {
																	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1614 = int32(0)
																	if v1050 != 0 {
																		v1617 = v1614
																		v1623 = v1614
																		v1625 = v1614
																		for {
																			v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617+v972))))
																			v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617+v1613))))
																			v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648^v1650)+uint32(_consts[1053]))))
																			v1657 = v1617 | int32(1)
																			v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972+v1657))))
																			v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657+v1613))))
																			v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1659^v1661)+uint32(_consts[1053]))))
																			v1666 = v1623 + v1654 + v1665
																			v1667 = int32(2)
																			v1668 = v1617 + v1667
																			v1670 = v1625 + v1667
																			if v1670 != v1046 {
																				v1617 = v1668
																				v1623 = v1666
																				v1625 = v1670
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v1672 = v1668
																		v1678 = v1666
																	} else {
																		v1672 = v1614
																		v1678 = v1614
																	}
																	if v1048 == int32(0) {
																		v2057 = v1678
																	} else {
																		v1705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672+v972))))
																		v1707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672+v1613))))
																		v1711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1705^v1707)+uint32(_consts[1053]))))
																		v2057 = v1678 + v1711
																	}
																}
															} else {
																if v1595&int32(1) == int32(0) {
																	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1718 = v1717
																} else {
																	v1718 = v972
																}
																if v60 <= int32(3) {
																	if v60 == int32(0) {
																		v2048 = int64(0)
																	} else {
																		v1725 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																			v1728 = v1718
																			v1734 = int32(0)
																			v1757 = v1725
																			for {
																				v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+3)))
																				v1761 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+uint32(_consts[1053]))))
																				v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+2)))
																				v1765 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1762)+uint32(_consts[1053]))))
																				v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+1)))
																				v1769 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1766)+uint32(_consts[1053]))))
																				v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
																				v1773 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+uint32(_consts[1053]))))
																				v1777 = v1761 + (v1765 + (v1769 + (v1757 + v1773)))
																				v1778 = int32(4)
																				v1779 = v1728 + v1778
																				v1781 = v1734 + v1778
																				if v1781 != v1044 {
																					v1728 = v1779
																					v1734 = v1781
																					v1757 = v1777
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1783 = v1779
																			v1812 = v1777
																		} else {
																			v1783 = v1718
																			v1812 = v1725
																		}
																		v1813 = int32(0)
																		if v1042 == v1813 {
																			v2048 = v1812
																		} else {
																			v1816 = v1783
																			v1822 = v1813
																			v1845 = v1812
																			for {
																				v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816))))
																				v1849 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+uint32(_consts[1053]))))
																				v1850 = v1845 + v1849
																				v1851 = int32(1)
																				v1854 = v1822 + v1851
																				if v1854 != v1042 {
																					v1816 = v1816 + v1851
																					v1822 = v1854
																					v1845 = v1850
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2048 = v1850
																		}
																	}
																} else {
																	v1860 = int64(0)
																	if v60 < int32(4) {
																		v1939 = v1718
																		v1940 = v60
																		v1945 = v1860
																	} else {
																		if v1718 != (v1718+int32(3))&int32(-4) {
																			v1939 = v1718
																			v1940 = v60
																			v1945 = v1860
																		} else {
																			v1869 = v60 - int32(4)
																			v1873 = int32(base.Ui32(v1869)>>(uint(int32(2))%32)) + int32(1)
																			v1875 = v1873 & int32(3)
																			if base.Ui32(v1869) < base.Ui32(int32(12)) {
																				v1911 = v1718
																				v1912 = v60
																				v1917 = v1860
																			} else {
																				v1881 = v1718
																				v1882 = v60
																				v1883 = int32(0)
																				v1887 = v1860
																				for {
																					v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+12))
																					v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+8))
																					v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+4))
																					v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1881)))
																					v1903 = base.I64_extend_i32_u(base.I32_popcnt(v1888)) + (base.I64_extend_i32_u(base.I32_popcnt(v1891)) + (base.I64_extend_i32_u(base.I32_popcnt(v1894)) + (v1887 + base.I64_extend_i32_u(base.I32_popcnt(v1897)))))
																					v1904 = int32(16)
																					v1905 = v1882 - v1904
																					v1907 = v1881 + v1904
																					v1909 = v1883 + int32(4)
																					if v1909 != v1873&int32(2147483644) {
																						v1881 = v1907
																						v1882 = v1905
																						v1883 = v1909
																						v1887 = v1903
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1911 = v1907
																				v1912 = v1905
																				v1917 = v1903
																			}
																			if v1875 == int32(0) {
																				v1939 = v1911
																				v1940 = v1912
																				v1945 = v1917
																			} else {
																				v1922 = v1912
																				v1923 = v1911
																				v1924 = int32(0)
																				v1927 = v1917
																				for {
																					v1928 = int32(4)
																					v1929 = v1922 - v1928
																					v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
																					v1933 = v1927 + base.I64_extend_i32_u(base.I32_popcnt(v1930))
																					v1935 = v1923 + v1928
																					v1937 = v1924 + int32(1)
																					if v1937 != v1875 {
																						v1922 = v1929
																						v1923 = v1935
																						v1924 = v1937
																						v1927 = v1933
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1939 = v1935
																				v1940 = v1929
																				v1945 = v1933
																			}
																		}
																	}
																	if v1940 == int32(0) {
																		v2018 = v1945
																	} else {
																		v1949 = v1940 & int32(3)
																		if v1949 == int32(0) {
																			v1972 = v1939
																			v1974 = v1940
																			v1978 = v1945
																		} else {
																			v1955 = v1940
																			v1956 = v1939
																			v1957 = int32(0)
																			v1959 = v1945
																			for {
																				v1960 = int32(1)
																				v1961 = v1955 - v1960
																				v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956))))
																				v1965 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+uint32(_consts[1053]))))
																				v1966 = v1959 + v1965
																				v1968 = v1956 + v1960
																				v1970 = v1957 + v1960
																				if v1970 != v1949 {
																					v1955 = v1961
																					v1956 = v1968
																					v1957 = v1970
																					v1959 = v1966
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1972 = v1968
																			v1974 = v1961
																			v1978 = v1966
																		}
																		if base.Ui32(v1940) < base.Ui32(int32(4)) {
																			v2018 = v1978
																		} else {
																			v1981 = v1972
																			v1983 = v1974
																			v1987 = v1978
																			for {
																				v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+3)))
																				v1991 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1988)+uint32(_consts[1053]))))
																				v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+2)))
																				v1995 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1992)+uint32(_consts[1053]))))
																				v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+1)))
																				v1999 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1996)+uint32(_consts[1053]))))
																				v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981))))
																				v2003 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2000)+uint32(_consts[1053]))))
																				v2007 = v1991 + (v1995 + (v1999 + (v1987 + v2003)))
																				v2008 = int32(4)
																				v2011 = v1983 - v2008
																				if v2011 != 0 {
																					v1981 = v1981 + v2008
																					v1983 = v2011
																					v1987 = v2007
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2018 = v2007
																		}
																	}
																	v2048 = v2018
																}
																v2057 = v1052 - base.I32_wrap_i64(v2048)
															}
														} else {
															if v1595&int32(1) == int32(0) {
																if v1595&int32(1) == int32(0) {
																	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																	v1718 = v1717
																} else {
																	v1718 = v972
																}
																if v60 <= int32(3) {
																	if v60 == int32(0) {
																		v2048 = int64(0)
																	} else {
																		v1725 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v1050) {
																			v1728 = v1718
																			v1734 = int32(0)
																			v1757 = v1725
																			for {
																				v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+3)))
																				v1761 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1758)+uint32(_consts[1053]))))
																				v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+2)))
																				v1765 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1762)+uint32(_consts[1053]))))
																				v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+1)))
																				v1769 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1766)+uint32(_consts[1053]))))
																				v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
																				v1773 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+uint32(_consts[1053]))))
																				v1777 = v1761 + (v1765 + (v1769 + (v1757 + v1773)))
																				v1778 = int32(4)
																				v1779 = v1728 + v1778
																				v1781 = v1734 + v1778
																				if v1781 != v1044 {
																					v1728 = v1779
																					v1734 = v1781
																					v1757 = v1777
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1783 = v1779
																			v1812 = v1777
																		} else {
																			v1783 = v1718
																			v1812 = v1725
																		}
																		v1813 = int32(0)
																		if v1042 == v1813 {
																			v2048 = v1812
																		} else {
																			v1816 = v1783
																			v1822 = v1813
																			v1845 = v1812
																			for {
																				v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816))))
																				v1849 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+uint32(_consts[1053]))))
																				v1850 = v1845 + v1849
																				v1851 = int32(1)
																				v1854 = v1822 + v1851
																				if v1854 != v1042 {
																					v1816 = v1816 + v1851
																					v1822 = v1854
																					v1845 = v1850
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2048 = v1850
																		}
																	}
																} else {
																	v1860 = int64(0)
																	if v60 < int32(4) {
																		v1939 = v1718
																		v1940 = v60
																		v1945 = v1860
																	} else {
																		if v1718 != (v1718+int32(3))&int32(-4) {
																			v1939 = v1718
																			v1940 = v60
																			v1945 = v1860
																		} else {
																			v1869 = v60 - int32(4)
																			v1873 = int32(base.Ui32(v1869)>>(uint(int32(2))%32)) + int32(1)
																			v1875 = v1873 & int32(3)
																			if base.Ui32(v1869) < base.Ui32(int32(12)) {
																				v1911 = v1718
																				v1912 = v60
																				v1917 = v1860
																			} else {
																				v1881 = v1718
																				v1882 = v60
																				v1883 = int32(0)
																				v1887 = v1860
																				for {
																					v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+12))
																					v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+8))
																					v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+4))
																					v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1881)))
																					v1903 = base.I64_extend_i32_u(base.I32_popcnt(v1888)) + (base.I64_extend_i32_u(base.I32_popcnt(v1891)) + (base.I64_extend_i32_u(base.I32_popcnt(v1894)) + (v1887 + base.I64_extend_i32_u(base.I32_popcnt(v1897)))))
																					v1904 = int32(16)
																					v1905 = v1882 - v1904
																					v1907 = v1881 + v1904
																					v1909 = v1883 + int32(4)
																					if v1909 != v1873&int32(2147483644) {
																						v1881 = v1907
																						v1882 = v1905
																						v1883 = v1909
																						v1887 = v1903
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1911 = v1907
																				v1912 = v1905
																				v1917 = v1903
																			}
																			if v1875 == int32(0) {
																				v1939 = v1911
																				v1940 = v1912
																				v1945 = v1917
																			} else {
																				v1922 = v1912
																				v1923 = v1911
																				v1924 = int32(0)
																				v1927 = v1917
																				for {
																					v1928 = int32(4)
																					v1929 = v1922 - v1928
																					v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
																					v1933 = v1927 + base.I64_extend_i32_u(base.I32_popcnt(v1930))
																					v1935 = v1923 + v1928
																					v1937 = v1924 + int32(1)
																					if v1937 != v1875 {
																						v1922 = v1929
																						v1923 = v1935
																						v1924 = v1937
																						v1927 = v1933
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1939 = v1935
																				v1940 = v1929
																				v1945 = v1933
																			}
																		}
																	}
																	if v1940 == int32(0) {
																		v2018 = v1945
																	} else {
																		v1949 = v1940 & int32(3)
																		if v1949 == int32(0) {
																			v1972 = v1939
																			v1974 = v1940
																			v1978 = v1945
																		} else {
																			v1955 = v1940
																			v1956 = v1939
																			v1957 = int32(0)
																			v1959 = v1945
																			for {
																				v1960 = int32(1)
																				v1961 = v1955 - v1960
																				v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1956))))
																				v1965 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+uint32(_consts[1053]))))
																				v1966 = v1959 + v1965
																				v1968 = v1956 + v1960
																				v1970 = v1957 + v1960
																				if v1970 != v1949 {
																					v1955 = v1961
																					v1956 = v1968
																					v1957 = v1970
																					v1959 = v1966
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1972 = v1968
																			v1974 = v1961
																			v1978 = v1966
																		}
																		if base.Ui32(v1940) < base.Ui32(int32(4)) {
																			v2018 = v1978
																		} else {
																			v1981 = v1972
																			v1983 = v1974
																			v1987 = v1978
																			for {
																				v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+3)))
																				v1991 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1988)+uint32(_consts[1053]))))
																				v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+2)))
																				v1995 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1992)+uint32(_consts[1053]))))
																				v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981)+1)))
																				v1999 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1996)+uint32(_consts[1053]))))
																				v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981))))
																				v2003 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2000)+uint32(_consts[1053]))))
																				v2007 = v1991 + (v1995 + (v1999 + (v1987 + v2003)))
																				v2008 = int32(4)
																				v2011 = v1983 - v2008
																				if v2011 != 0 {
																					v1981 = v1981 + v2008
																					v1983 = v2011
																					v1987 = v2007
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2018 = v2007
																		}
																	}
																	v2048 = v2018
																}
																v2057 = v1052 - base.I32_wrap_i64(v2048)
															} else {
																v2057 = int32(0)
															}
														}
														v2083 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
														v2084 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
														v2085 = v2083 - v2084
														if base.F64_lt(base.F64_convert_i32_s(v1572), base.F64_add(base.F64_convert_i32_s(v2057), base.F64_mul(base.F64_convert_i32_s(v2085*v2085*v2085), float64(-0.1)))) != 0 {
															v2093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+4)))
															if v2093&int32(4) != 0 {
															} else {
																v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																if v2096 == int32(1) {
																	v2101 = F__emscripten_memset_bulkmem(m, v974, base.I32_extend8_s(int32(255)), v60)
																	mBase = m.M
																} else {
																	if v60 <= int32(0) {
																	} else {
																		v2104 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v2105 = int32(0)
																		if base.Ui32(int32(4)) <= base.Ui32(v60) {
																			v2110 = v2105
																			v2120 = v2105
																			for {
																				v2140 = v2110 + v974
																				v2141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140))))
																				v2143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2110+v2104))))
																				v2144 = v2141 | v2143
																				*(*uint8)(unsafe.Add(mBase, uint32(v2140))) = uint8(v2144)
																				v2147 = v2110 | int32(1)
																				v2148 = v974 + v2147
																				v2149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2148))))
																				v2151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2147))))
																				v2152 = v2149 | v2151
																				*(*uint8)(unsafe.Add(mBase, uint32(v2148))) = uint8(v2152)
																				v2155 = v2110 | int32(2)
																				v2156 = v974 + v2155
																				v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156))))
																				v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2155))))
																				v2160 = v2157 | v2159
																				*(*uint8)(unsafe.Add(mBase, uint32(v2156))) = uint8(v2160)
																				v2163 = v2110 | int32(3)
																				v2164 = v974 + v2163
																				v2165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2164))))
																				v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2163))))
																				v2168 = v2165 | v2167
																				*(*uint8)(unsafe.Add(mBase, uint32(v2164))) = uint8(v2168)
																				v2170 = int32(4)
																				v2171 = v2110 + v2170
																				v2173 = v2120 + v2170
																				if v2173 != v1040 {
																					v2110 = v2171
																					v2120 = v2173
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2175 = v2171
																		} else {
																			v2175 = v2105
																		}
																		if v1042 == int32(0) {
																		} else {
																			v2207 = v2175
																			v2213 = v2105
																			for {
																				v2237 = v2207 + v974
																				v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2237))))
																				v2240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2207+v2104))))
																				v2241 = v2238 | v2240
																				*(*uint8)(unsafe.Add(mBase, uint32(v2237))) = uint8(v2241)
																				v2243 = int32(1)
																				v2246 = v2213 + v2243
																				if v2246 != v1042 {
																					v2207 = v2207 + v2243
																					v2213 = v2246
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	}
																}
															}
															*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v1089)
															v2279 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v2279 + int32(1)
															v2522 = v1071 + int32(2)
															v2527 = v1076
														} else {
															v2285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+4)))
															if v2285&int32(4) != 0 {
															} else {
																v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																if v2288 == int32(1) {
																	v2293 = F__emscripten_memset_bulkmem(m, v972, base.I32_extend8_s(int32(255)), v60)
																	mBase = m.M
																} else {
																	if v60 <= int32(0) {
																	} else {
																		v2296 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
																		v2297 = int32(0)
																		if base.Ui32(int32(4)) <= base.Ui32(v60) {
																			v2302 = v2297
																			v2312 = v2297
																			for {
																				v2332 = v2302 + v972
																				v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2332))))
																				v2335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302+v2296))))
																				v2336 = v2333 | v2335
																				*(*uint8)(unsafe.Add(mBase, uint32(v2332))) = uint8(v2336)
																				v2339 = v2302 | int32(1)
																				v2340 = v972 + v2339
																				v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2340))))
																				v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2339))))
																				v2344 = v2341 | v2343
																				*(*uint8)(unsafe.Add(mBase, uint32(v2340))) = uint8(v2344)
																				v2347 = v2302 | int32(2)
																				v2348 = v972 + v2347
																				v2349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2348))))
																				v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2347))))
																				v2352 = v2349 | v2351
																				*(*uint8)(unsafe.Add(mBase, uint32(v2348))) = uint8(v2352)
																				v2355 = v2302 | int32(3)
																				v2356 = v972 + v2355
																				v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2356))))
																				v2359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2296+v2355))))
																				v2360 = v2357 | v2359
																				*(*uint8)(unsafe.Add(mBase, uint32(v2356))) = uint8(v2360)
																				v2362 = int32(4)
																				v2363 = v2302 + v2362
																				v2365 = v2312 + v2362
																				if v2365 != v1040 {
																					v2302 = v2363
																					v2312 = v2365
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v2367 = v2363
																		} else {
																			v2367 = v2297
																		}
																		if v1042 == int32(0) {
																		} else {
																			v2399 = v2367
																			v2405 = v2297
																			for {
																				v2429 = v2399 + v972
																				v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2429))))
																				v2432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2399+v2296))))
																				v2433 = v2430 | v2432
																				*(*uint8)(unsafe.Add(mBase, uint32(v2429))) = uint8(v2433)
																				v2435 = int32(1)
																				v2438 = v2405 + v2435
																				if v2438 != v1042 {
																					v2399 = v2399 + v2435
																					v2405 = v2438
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	}
																}
															}
															*(*uint16)(unsafe.Add(mBase, uint32(v1076))) = uint16(v1089)
															v2471 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
															*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v2471 + int32(1)
															v2522 = v1071
															v2527 = v1076 + int32(2)
														}
													}
												}
												v2538 = v1060 + int32(1)
												if v2538 != v1038 {
													v1060 = v2538
													v1071 = v2522
													v1076 = v2527
													continue
												} else {
													break
												}
												break
											}
											v2555 = v2522
											v2560 = v2527
											v2570 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v2555))) = uint16(v2570)
											*(*uint16)(unsafe.Add(mBase, uint32(v2560))) = uint16(v2570)
											*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v727
											*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v701
											return v31
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
