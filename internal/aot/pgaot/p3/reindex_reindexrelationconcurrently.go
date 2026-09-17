package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexRelationConcurrently(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v596 int32
	_ = v596
	var v622 int32
	_ = v622
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v660 int32
	_ = v660
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v846 int64
	_ = v846
	var v848 int64
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v913 int64
	_ = v913
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v927 int64
	_ = v927
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v941 int64
	_ = v941
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v955 int64
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
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
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1208 int32
	_ = v1208
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1295 int32
	_ = v1295
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int64
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int64
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1569 int64
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1574 int64
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1595 int32
	_ = v1595
	var v1602 int32
	_ = v1602
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1643 int32
	_ = v1643
	var v1649 int32
	_ = v1649
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int64
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1717 int32
	_ = v1717
	var v1723 int32
	_ = v1723
	var v1738 int32
	_ = v1738
	var v1743 int32
	_ = v1743
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1857 int32
	_ = v1857
	var v1863 int32
	_ = v1863
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1880 int32
	_ = v1880
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1987 int32
	_ = v1987
	var v1993 int32
	_ = v1993
	var v2001 int64
	_ = v2001
	var v2003 int64
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2043 int32
	_ = v2043
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2068 int64
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2082 int64
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2089 int32
	_ = v2089
	var v2096 int64
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2110 int64
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2259 int32
	_ = v2259
	var v2265 int32
	_ = v2265
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2329 int32
	_ = v2329
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2391 int32
	_ = v2391
	var v2397 int32
	_ = v2397
	var v2401 int64
	_ = v2401
	var v2405 int64
	_ = v2405
	var v2407 int64
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2447 int32
	_ = v2447
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2472 int64
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2479 int32
	_ = v2479
	var v2486 int64
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2493 int32
	_ = v2493
	var v2500 int64
	_ = v2500
	var v2503 int32
	_ = v2503
	var v2507 int32
	_ = v2507
	var v2514 int64
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2599 int32
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2608 int32
	_ = v2608
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2619 int32
	_ = v2619
	var v2627 int32
	_ = v2627
	var v2633 int32
	_ = v2633
	var v2639 int32
	_ = v2639
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2652 int32
	_ = v2652
	var v2656 int32
	_ = v2656
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2675 int32
	_ = v2675
	var v2681 int32
	_ = v2681
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2755 int32
	_ = v2755
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2796 int32
	_ = v2796
	var v2801 int32
	_ = v2801
	var v2804 int32
	_ = v2804
	var v2807 int32
	_ = v2807
	var v2809 int32
	_ = v2809
	var v2814 int32
	_ = v2814
	var v2818 int32
	_ = v2818
	var v2823 int32
	_ = v2823
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2829 int32
	_ = v2829
	var v2837 int32
	_ = v2837
	var v2843 int32
	_ = v2843
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2856 int32
	_ = v2856
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2966 int32
	_ = v2966
	var v2968 int32
	_ = v2968
	var v2973 int32
	_ = v2973
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2988 int32
	_ = v2988
	var v2996 int32
	_ = v2996
	var v3002 int32
	_ = v3002
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3021 int32
	_ = v3021
	var v3055 int32
	_ = v3055
	var v3059 int32
	_ = v3059
	var v3062 int32
	_ = v3062
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3109 int32
	_ = v3109
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3117 int32
	_ = v3117
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3158 int32
	_ = v3158
	var v3162 int32
	_ = v3162
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3222 int32
	_ = v3222
	var v3256 int32
	_ = v3256
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3278 int32
	_ = v3278
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
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
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3391 int32
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3407 int32
	_ = v3407
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3415 int32
	_ = v3415
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3440 int32
	_ = v3440
	var v3442 int32
	_ = v3442
	var v3446 int32
	_ = v3446
	var v3448 int32
	_ = v3448
	var v3450 int32
	_ = v3450
	var v3452 int32
	_ = v3452
	var v3463 int32
	_ = v3463
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3486 int32
	_ = v3486
	var v3493 int32
	_ = v3493
	var v3501 int32
	_ = v3501
	var v3506 int32
	_ = v3506
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3565 int32
	_ = v3565
	var v3597 int32
	_ = v3597
	var v3600 int32
	_ = v3600
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3619 int32
	_ = v3619
	var v3623 int32
	_ = v3623
	var v3659 int32
	_ = v3659
	var v3663 int32
	_ = v3663
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3682 int32
	_ = v3682
	var v3687 int32
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3692 int32
	_ = v3692
	var v3693 int32
	_ = v3693
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3752 int32
	_ = v3752
	var v3789 int64
	_ = v3789
	var v3796 int32
	_ = v3796
	var v3799 int32
	_ = v3799
	var v3804 int32
	_ = v3804
	var v3812 int32
	_ = v3812
	var v3815 int32
	_ = v3815
	var v3820 int32
	_ = v3820
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3845 int32
	_ = v3845
	var v3848 int32
	_ = v3848
	var v3851 int32
	_ = v3851
	var v3852 int32
	_ = v3852
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3870 int32
	_ = v3870
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3876 int32
	_ = v3876
	var v3880 int32
	_ = v3880
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3884 int64
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3894 int64
	_ = v3894
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3904 int32
	_ = v3904
	var v3907 int32
	_ = v3907
	var v3909 int32
	_ = v3909
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3919 int32
	_ = v3919
	var v3921 int32
	_ = v3921
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3931 int32
	_ = v3931
	var v3936 int32
	_ = v3936
	var v3945 int32
	_ = v3945
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3979 int32
	_ = v3979
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3985 int32
	_ = v3985
	var v3989 int32
	_ = v3989
	var v4027 int32
	_ = v4027
	var v4033 int32
	_ = v4033
	var v4036 int32
	_ = v4036
	var v4039 int32
	_ = v4039
	var v4042 int32
	_ = v4042
	var v4045 int32
	_ = v4045
	var v4048 int32
	_ = v4048
	var v4055 int32
	_ = v4055
	var v4059 int32
	_ = v4059
	var v4064 int32
	_ = v4064
	var v4068 int32
	_ = v4068
	var v4074 int32
	_ = v4074
	var v4079 int32
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4089 int32
	_ = v4089
	var v4094 int32
	_ = v4094
	var v4098 int32
	_ = v4098
	var v4104 int32
	_ = v4104
	var v4109 int32
	_ = v4109
	var v4113 int32
	_ = v4113
	var v4119 int32
	_ = v4119
	var v4124 int32
	_ = v4124
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4129 int32
	_ = v4129
	var v4133 int32
	_ = v4133
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4153 int32
	_ = v4153
	var v4177 int32
	_ = v4177
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4193 int32
	_ = v4193
	var v4197 int32
	_ = v4197
	var v4234 int32
	_ = v4234
	var v4237 int32
	_ = v4237
	var v4241 int32
	_ = v4241
	var v4246 int32
	_ = v4246
	var v4249 int32
	_ = v4249
	var v4251 int32
	_ = v4251
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4259 int32
	_ = v4259
	var v4269 int32
	_ = v4269
	var v4276 int32
	_ = v4276
	var v4315 int32
	_ = v4315
	var v4319 int32
	_ = v4319
	var v4324 int32
	_ = v4324
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4338 int32
	_ = v4338
	var v4343 int32
	_ = v4343
	v4 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(416)
	m.G0 = v38
	v41 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+232)) = v41
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+224)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[2]))
	v52 = F_AllocSetContextCreateInternal(m, v47, int32(_a_F_ReindexRelationConcurrently_0), v4, int32(1024), int32(_a_F_ReindexRelationConcurrently_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v56&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v59 = int32(_a_F_ReindexRelationConcurrently_2)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v63 = F_get_rel_name(m, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v78 = v4
	v79 = v4
	goto L5
L5:
	;
	v80 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v65 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v67 = F_get_namespace_name(m, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_getrusage(m, v38+int32(264))
	mBase = m.M
	F_gettimeofday(m, v38+int32(248))
	mBase = m.M
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v60
	v78 = v63
	v79 = v67
	goto L5
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L1
	} else {
		goto L727
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4315 = m.ExcPending
	if v4315 != 0 {
		goto L1
	} else {
		goto L724
	}
L12:
	;
	m.G0 = v38 + int32(416)
	return v4276
L13:
	;
	if v672 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L14:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+112))
	if v467 != 0 {
		goto L136
	} else {
		goto L137
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L132
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L127
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L123
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L119
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L115
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L111
	}
L21:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v263 = F_IndexGetRelation(m, l1, int32(base.Ui32(v258&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L78
	}
L22:
	;
	v86 = int32(_a_F_ReindexRelationConcurrently_2)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v91 = F_lappend_oid(m, int32(0), l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	switch v80&int32(255) - int32(105) {
	case 0:
		goto L21
	default:
		goto L15
	case 4, 9, 11:
		goto L22
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v87
	goto L25
L25:
	;
	if base.Ui32(l1) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v97&int32(4) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v107 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v101 = F_try_table_open(m, l1, int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v104 = F_table_open(m, l1, int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	if v101 != 0 {
		v106 = v101
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v4276 = v4
	goto L12
L33:
	;
	v106 = v104
	goto L27
L34:
	;
	v109 = int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+56))
	if base.Ui32(v110) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		v119 = v109
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v120 = F_RelationGetIndexList(m, v106)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	if v119 != 0 {
		goto L19
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+68))
	if v114 == int32(99) {
		v119 = v109
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v117 = F_isTempToastNamespace(m, v114)
	mBase = m.M
	v119 = v117
	goto L38
L41:
	;
	goto L36
L42:
	;
	if v120 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v442 = v4
	goto L14
L44:
	;
	goto L45
L45:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v124 <= int32(0) {
		v442 = v4
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v129 = int32(0)
	v139 = v4
	goto L47
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v129<<(uint(int32(2))%32))))
	v169 = F_index_open(m, v167, int32(4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v442 = v250
	goto L14
L49:
	;
	F_relation_close(m, v169, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L76
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)+192))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+18)))
	if v172 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v177 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+15)))
	if v206 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	if v177 == int32(0) {
		v250 = v139
		goto L49
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v184 = F_get_rel_namespace(m, v167)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v186 = F_get_namespace_name(m, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v188 = F_get_rel_name(m, v167)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+132)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v186
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_4), v38+int32(128))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(_a_F_ReindexRelationConcurrently_5), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3685), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v250 = v139
	goto L49
L63:
	;
	v211 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v236 = int32(_a_F_ReindexRelationConcurrently_2)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v241 = F_palloc(m, int32(16))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	if v211 == int32(0) {
		v250 = v139
		goto L49
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v218 = F_get_rel_namespace(m, v167)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v220 = F_get_namespace_name(m, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v222 = F_get_rel_name(m, v167)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+116)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = v220
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_8), v38+int32(112))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3691), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v250 = v139
	goto L49
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v167
	v244 = F_lappend(m, v139, v241)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v237
	v250 = v244
	goto L49
L76:
	;
	v255 = v129 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v255 < v256 {
		v129 = v255
		v139 = v250
		goto L47
	} else {
		goto L77
	}
L77:
	;
	goto L48
L78:
	;
	if v263 == int32(0) {
		v4276 = v4
		goto L12
	} else {
		goto L79
	}
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(v263) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v269 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v269 != int32(99) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v275 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v273 = F_isTempToastNamespace(m, v269)
	mBase = m.M
	v275 = v273
	goto L86
L85:
	;
	v275 = int32(1)
	goto L86
L86:
	;
	goto L83
L87:
	;
	v276 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v280&int32(4) != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	if v276 == int32(0) {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v290 != 0 {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	v284 = F_try_table_open(m, v263, int32(4))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v287 = F_table_open(m, v263, int32(4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	if v284 != 0 {
		v289 = v284
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v4276 = v4
	goto L12
L98:
	;
	v289 = v287
	goto L92
L99:
	;
	v292 = int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+56))
	if base.Ui32(v293) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		v302 = v292
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	F_relation_close(m, v289, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L107
	}
L102:
	;
	if v302 != 0 {
		goto L16
	} else {
		goto L106
	}
L103:
	;
	goto L102
L104:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v289)+48))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+68))
	if v297 == int32(99) {
		v302 = v292
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v300 = F_isTempToastNamespace(m, v297)
	mBase = m.M
	v302 = v300
	goto L103
L106:
	;
	goto L101
L107:
	;
	v306 = int32(_a_F_ReindexRelationConcurrently_2)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+188)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+156)) = v263
	v315 = F_list_make1_impl(m, int32(472), v38+int32(156))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v318 = F_palloc(m, int32(16))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = l1
	v322 = F_lappend(m, int32(0), v318)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v307
	v672 = v322
	v681 = v315
	goto L13
L111:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_9), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3650), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+144)) = v349 + int32(4)
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_10), v38+int32(144))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3670), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_9), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3780), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_11), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3791), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v402 = F_get_rel_name(m, l1)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+160)) = v402
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_10), v38+int32(160))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3816), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_12), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3845), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
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
	v470 = F_table_open(m, v467, int32(4))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	v634 = v442
	v643 = v91
	goto L138
L138:
	;
	F_relation_close(m, v106, int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L166
	}
L139:
	;
	v472 = int32(_a_F_ReindexRelationConcurrently_2)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v476 = F_lappend_oid(m, v91, v467)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v473
	v480 = F_RelationGetIndexList(m, v470)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	F_relation_close(m, v470, int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L165
	}
L142:
	;
	if v480 == int32(0) {
		v596 = v442
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	if v484 <= int32(0) {
		v596 = v442
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v488 = int32(0)
	v498 = v442
	goto L145
L145:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v480)+12))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522+v488<<(uint(int32(2))%32))))
	v528 = F_index_open(m, v526, int32(4))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v596 = v576
	goto L141
L147:
	;
	F_relation_close(m, v528, int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L163
	}
L148:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v528)+192))
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+18)))
	if v531 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v532 = int32(_a_F_ReindexRelationConcurrently_2)
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v537 = F_palloc(m, int32(16))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v546 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v526
	v540 = F_lappend(m, v498, v537)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v533
	v576 = v540
	goto L147
L154:
	;
	if v546 == int32(0) {
		v576 = v498
		goto L147
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v553 = F_get_rel_namespace(m, v526)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v555 = F_get_namespace_name(m, v553)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v557 = F_get_rel_name(m, v526)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v555
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_4), v38+int32(96))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errhint(m, int32(_a_F_ReindexRelationConcurrently_5), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3738), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v576 = v498
	goto L147
L163:
	;
	v582 = v488 + int32(1)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	if v582 < v583 {
		v488 = v582
		v498 = v576
		goto L145
	} else {
		goto L164
	}
L164:
	;
	goto L146
L165:
	;
	v634 = v596
	v643 = v476
	goto L138
L166:
	;
	v672 = v634
	v681 = v643
	goto L13
L167:
	;
	v4276 = int32(0)
	goto L12
L168:
	;
	goto L169
L169:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v699 == int32(1664) {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v702 = int32(0)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v702 < v704 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v715 = int32(0)
	v717 = v702
	v724 = v702
	goto L174
L172:
	;
	v1595 = v702
	v1602 = v702
	goto L173
L173:
	;
	if v681 == int32(0) {
		v1717 = v1602
		v1723 = v4
		goto L309
	} else {
		goto L310
	}
L174:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743+v715<<(uint(int32(2))%32))))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v750 = F_index_open(m, v748, int32(4))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v1595 = v1525
	v1602 = v1539
	goto L173
L176:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v750)+192))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	v755 = F_table_open(m, v753, int32(4))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(184)))) = v762
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(180)))) = v765
	goto L178
L178:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v755)+48))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+80))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5])) = v769 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4])) = v768
	goto L179
L179:
	;
	v777 = int32(_a_F_ReindexRelationConcurrently_13)
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[6]))
	v781 = v779 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[6])) = v781
	goto L180
L180:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v785 = F_RelationGetIndexExpressions(m, v750)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	if v785 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v790 = int32(1)
	goto L185
L184:
	;
	v788 = F_RelationGetIndexPredicate(m, v750)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L186
	}
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v747)+12)) = uint8(base.B2i32(v790 == int32(0)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v755)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v747)+4)) = v794
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v750)+48))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v747)+8)) = v797
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v750)+48))
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799)+118)))
	if v800 == int32(116) {
		goto L11
	} else {
		goto L187
	}
L186:
	;
	v790 = v788
	goto L185
L187:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v806 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = int64(4)
	v846 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v747))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v846
	v848 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v747)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v848
	v852 = v38 + int32(224)
	v854 = v38 + int32(192)
	goto L194
L189:
	;
	goto L188
L190:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v810&int32(1) == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v815 = int32(_a_F_ReindexRelationConcurrently_14)
	v817 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v818 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v817 + v818
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v821 + v818
	*(*int32)(unsafe.Add(mBase, uint32(v806)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v806)+224)) = v794
	base.MemoryFill(m, v806+int32(232), int32(0), int32(160))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v832 + v818
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v838 - v818
	goto L189
L192:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1029 = F_get_rel_name(m, v1028)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L209
	}
L193:
	;
	goto L192
L194:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v864 == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v868&int32(1) == int32(0) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v873 = int32(_a_F_ReindexRelationConcurrently_14)
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v876 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v875 + v876
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	*(*int32)(unsafe.Add(mBase, uint32(v864))) = v879 + v876
	goto L198
L197:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	v1010 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v864))) = v1009 + v1010
	v1013 = int32(_a_F_ReindexRelationConcurrently_14)
	v1015 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1015 - v1010
	goto L193
L198:
	;
	v888 = v864 + int32(232)
	goto L199
L199:
	;
	v894 = int32(0)
	v897 = int32(0)
	goto L202
L202:
	;
	v903 = int32(2)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v852+v897<<(uint(v903)%32))))
	v907 = int32(3)
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v854+v897<<(uint(v907)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v888+v906<<(uint(v907)%32)))) = v913
	v916 = v897 | int32(1)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v852+v916<<(uint(v903)%32))))
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v854+v916<<(uint(v907)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v888+v920<<(uint(v907)%32)))) = v927
	v930 = v897 | v903
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v852+v930<<(uint(v903)%32))))
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v854+v930<<(uint(v907)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v888+v934<<(uint(v907)%32)))) = v941
	v944 = v897 | v907
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v852+v944<<(uint(v903)%32))))
	v955 = *(*int64)(unsafe.Add(mBase, uint32(v854+v944<<(uint(v907)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v888+v948<<(uint(v907)%32)))) = v955
	v957 = int32(4)
	v960 = v894 + v957
	if v960 != int32(4) {
		v894 = v960
		v897 = v897 + v957
		goto L202
	} else {
		goto L204
	}
L203:
	;
	goto L197
L204:
	;
	goto L203
L209:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v750)+192))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+4))
	v1035 = F_get_rel_namespace(m, v1034)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1038 = F_ChooseRelationName(m, v1029, int32(0), int32(_a_F_ReindexRelationConcurrently_15), v1035, int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1040 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1049 = int32(0)
	v1052 = m.G0
	v1054 = v1052 - int32(48)
	m.G0 = v1054
	v1057 = F_index_open(m, v1048, int32(3))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L221
	}
L213:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v755)+48))
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041)+119)))
	if v1042 != int32(116) {
		v1047 = v1040
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v750)+48))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+92))
	v1047 = v1046
	goto L212
L216:
	;
	goto L215
L217:
	;
	v1509 = F_index_open(m, v1436, int32(4))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L292
	}
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L1
	} else {
		goto L289
	}
L219:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L286
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L283
	}
L221:
	;
	v1059 = F_BuildIndexInfo(m, v1057)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+92))
	if v1061 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1065 = F_SearchSysCache1(m, int32(34), v1048)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L279
	}
L226:
	;
	if v1065 == int32(0) {
		goto L220
	} else {
		goto L227
	}
L227:
	;
	v1071 = F_SysCacheGetAttrNotNull(m, int32(34), v1065, int32(18))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v1075 = F_SysCacheGetAttrNotNull(m, int32(34), v1065, int32(19))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v1078 = F_SearchSysCache1(m, int32(57), v1048)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if v1078 == int32(0) {
		goto L219
	} else {
		goto L231
	}
L231:
	;
	v1086 = F_SysCacheGetAttr(m, int32(57), v1078, int32(33), v1054+int32(47))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+76))
	if v1088 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1091 = F_SysCacheGetAttrNotNull(m, int32(34), v1065, int32(20))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	v1100 = v1049
	goto L235
L235:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+84))
	if v1101 != 0 {
		goto L240
	} else {
		goto L241
	}
L236:
	;
	v1093 = F_text_to_cstring(m, v1091)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1095 = F_stringToNode(m, v1093)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_pfree(m, v1093)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1100 = v1095
	goto L235
L240:
	;
	v1104 = F_SysCacheGetAttrNotNull(m, int32(34), v1065, int32(21))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L243
	}
L241:
	;
	v1115 = v1049
	goto L242
L242:
	;
	v1116 = int32(0)
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+8))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+132))
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059)+116)))
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059)+117)))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+204))
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+28)))
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059)+124)))
	v1127 = F_makeIndexInfo(m, v1117, v1118, v1119, v1100, v1115, v1120, v1121, v1116, int32(1), v1125, v1126)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L248
	}
L243:
	;
	v1106 = F_text_to_cstring(m, v1104)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1108 = F_stringToNode(m, v1106)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1110 = F_make_ands_implicit(m, v1108)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_pfree(m, v1106)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1115 = v1110
	goto L242
L248:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	if int32(0) < v1129 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1132 = int32(12)
	v1149 = int32(0)
	v1151 = v1049
	goto L252
L250:
	;
	v1208 = v1049
	goto L251
L251:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	v1232 = F_palloc0(m, v1229<<(uint(int32(2))%32))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L256
	}
L252:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+52))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	v1182 = F_lappend(m, v1151, v1172+v1173<<(uint(int32(4))%32)+v1149*int32(100)+int32(24))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L254
	}
L253:
	;
	v1208 = v1182
	goto L251
L254:
	;
	v1184 = int32(1)
	v1185 = v1149 << (uint(v1184) % 32)
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1059+v1132+v1185))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1127+v1132+v1185))) = uint16(v1188)
	v1191 = v1149 + v1184
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	if v1191 < v1192 {
		v1149 = v1191
		v1151 = v1182
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if int32(0) < v1234 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1240 = v1116
	goto L260
L258:
	;
	v1295 = v1234
	goto L259
L259:
	;
	v1320 = F_palloc0(m, v1295<<(uint(int32(3))%32))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L1
	} else {
		goto L264
	}
L260:
	;
	v1276 = v1240 + int32(1)
	v1278 = F_get_attoptions(m, v1048, base.I32_extend16_s(v1276))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L1
	} else {
		goto L262
	}
L261:
	;
	v1295 = v1281
	goto L259
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1232+v1240<<(uint(int32(2))%32)))) = v1278
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1276 < v1281 {
		v1240 = v1276
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1322 = int32(0)
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1322 < v1323 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1329 = v1322
	goto L268
L266:
	;
	goto L267
L267:
	;
	v1420 = int32(0)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+48))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+84))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+248))
	v1427 = int32(24)
	v1436 = F_index_create(m, v755, v1038, v1420, v1420, v1420, v1420, v1127, v1208, v1425, v1047, v1426, v1071+v1427, v1232, v1075+v1427, v1320, v1086, int32(12), v1420, int32(1), v1420, v1420)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L275
	}
L268:
	;
	v1363 = v1329 + int32(1)
	v1365 = F_SearchSysCache2(m, int32(7), v1048, base.I32_extend16_s(v1363))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L270
	}
L269:
	;
	goto L267
L270:
	;
	if v1365 == int32(0) {
		goto L218
	} else {
		goto L271
	}
L271:
	;
	v1373 = F_SysCacheGetAttr(m, int32(7), v1365, int32(21), v1054+int32(47))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_ReleaseCatCache(m, v1365)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1379 = v1320 + v1329<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1379))) = v1373
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054)+47)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1379)+4)) = uint8(v1381)
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1363 < v1383 {
		v1329 = v1363
		goto L268
	} else {
		goto L274
	}
L274:
	;
	goto L269
L275:
	;
	F_relation_close(m, v1057, int32(0))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_ReleaseCatCache(m, v1065)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_ReleaseCatCache(m, v1078)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	m.G0 = v1054 + int32(48)
	goto L217
L279:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_16), int32(0))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1333), int32(_a_F_ReindexRelationConcurrently_18))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054))) = v1048
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_19), v1054)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1338), int32(_a_F_ReindexRelationConcurrently_18))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+16)) = v1048
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_20), v1054+int32(16))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1350), int32(_a_F_ReindexRelationConcurrently_18))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+36)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+32)) = v1363
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_21), v1054+int32(32))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1431), int32(_a_F_ReindexRelationConcurrently_18))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	v1511 = int32(_a_F_ReindexRelationConcurrently_2)
	v1512 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v1516 = F_palloc(m, int32(16))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1516))) = v1436
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1516)+12)) = uint8(v1519)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+4)) = v1521
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+8)) = v1523
	v1525 = F_lappend(m, v717, v1516)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1528 = F_palloc(m, int32(8))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1530 = *(*int64)(unsafe.Add(mBase, uint32(v750)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1528))) = v1530
	v1532 = F_lappend(m, v724, v1528)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1535 = F_palloc(m, int32(8))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1537 = *(*int64)(unsafe.Add(mBase, uint32(v1509)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1535))) = v1537
	v1539 = F_lappend(m, v1532, v1535)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1512
	F_relation_close(m, v750, int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_relation_close(m, v1509, int32(0))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	F_AtEOXact_GUC(m, int32(0), v781)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v38)+184))
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5])) = v1553
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4])) = v1552
	goto L302
L302:
	;
	F_relation_close(m, v755, int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	if l0 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v1436
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v1564 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = v1564
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v1564
	v1569 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v1569
	v1572 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v1572
	v1574 = *(*int64)(unsafe.Add(mBase, uint32(v38)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+80)) = v1574
	F_EventTriggerCollectSimpleCommand(m, v38+int32(80), v38-int32(-64), l0)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1583 = v715 + int32(1)
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v1583 < v1584 {
		v715 = v1583
		v717 = v1525
		v724 = v1539
		goto L174
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	goto L175
L309:
	;
	if v1717 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L310:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v1623 <= int32(0) {
		v1717 = v1602
		v1723 = v4
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1627 = int32(0)
	v1643 = v1602
	v1649 = v4
	goto L312
L312:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1662+v1627<<(uint(int32(2))%32))))
	v1668 = F_table_open(m, v1666, int32(4))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L1
	} else {
		goto L314
	}
L313:
	;
	v1717 = v1679
	v1723 = v1690
	goto L309
L314:
	;
	v1670 = int32(_a_F_ReindexRelationConcurrently_2)
	v1671 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v1675 = F_palloc(m, int32(8))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1677 = *(*int64)(unsafe.Add(mBase, uint32(v1668)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1675))) = v1677
	v1679 = F_lappend(m, v1643, v1675)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1682 = F_palloc(m, int32(16))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1682))) = v1684
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1675)))
	*(*int64)(unsafe.Add(mBase, uint32(v1682)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v1682)+4)) = v1686
	v1690 = F_lappend(m, v1649, v1682)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1671
	F_relation_close(m, v1668, int32(0))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1698 = v1627 + int32(1)
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v1698 < v1699 {
		v1627 = v1698
		v1643 = v1679
		v1649 = v1690
		goto L312
	} else {
		goto L320
	}
L320:
	;
	goto L313
L321:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L1
	} else {
		goto L328
	}
L322:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+4))
	if v1738 <= int32(0) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1743 = int32(0)
	goto L324
L324:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+12))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1777+v1743<<(uint(int32(2))%32))))
	F_LockRelationIdForSession(m, v1781, int32(4))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L1
	} else {
		goto L326
	}
L325:
	;
	goto L321
L326:
	;
	v1786 = v1743 + int32(1)
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+4))
	if v1786 < v1787 {
		v1743 = v1786
		goto L324
	} else {
		goto L327
	}
L327:
	;
	goto L325
L328:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1834 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	F_WaitForLockersMultiple(m, v1723, int32(5), int32(1))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L1
	} else {
		goto L335
	}
L332:
	;
	goto L331
L333:
	;
	v1838 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1838&int32(1) == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1843 = int32(_a_F_ReindexRelationConcurrently_14)
	v1845 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1846 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1845 + v1846
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	*(*int32)(unsafe.Add(mBase, uint32(v1834))) = v1849 + v1846
	*(*int64)(unsafe.Add(mBase, uint32(v1834+int32(72))+232)) = int64(1)
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	*(*int32)(unsafe.Add(mBase, uint32(v1834))) = v1857 + v1846
	v1863 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1863 - v1846
	goto L332
L335:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	if v1595 != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L1
	} else {
		goto L446
	}
L338:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if int32(0) < v1873 {
		goto L341
	} else {
		goto L342
	}
L339:
	;
	goto L340
L340:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L1
	} else {
		goto L439
	}
L341:
	;
	v1880 = int32(0)
	goto L344
L342:
	;
	goto L343
L343:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L1
	} else {
		goto L383
	}
L344:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+12))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1912+v1880<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L346
	}
L345:
	;
	goto L343
L346:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v1920 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1916)+12)))
	if v1923 == int32(1) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L349
L351:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v1931 = F_LWLockAcquire(m, v1927+int32(512), int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L1
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	v1953 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L356
	}
L354:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1934)+124)))
	v1937 = v1935 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1934)+124)) = uint8(v1937)
	v1940 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+12))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v1941+v1942))) = uint8(v1937)
	v1946 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v1946+int32(512))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	F_PushActiveSnapshot(m, v1953)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1916)+4))
	v1961 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1961 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = int64(4)
	v2001 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1916))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2001
	v2003 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1916)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2003
	v2007 = v38 + int32(224)
	v2009 = v38 + int32(192)
	goto L364
L359:
	;
	goto L358
L360:
	;
	v1965 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1965&int32(1) == int32(0) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1970 = int32(_a_F_ReindexRelationConcurrently_14)
	v1972 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1973 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1972 + v1973
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1961)))
	*(*int32)(unsafe.Add(mBase, uint32(v1961))) = v1976 + v1973
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1961)+224)) = v1958
	base.MemoryFill(m, v1961+int32(232), int32(0), int32(160))
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1961)))
	*(*int32)(unsafe.Add(mBase, uint32(v1961))) = v1987 + v1973
	v1993 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1993 - v1973
	goto L359
L362:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v1916)+4))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v1916)))
	F_index_concurrently_build(m, v2183, v2184)
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L1
	} else {
		goto L379
	}
L363:
	;
	goto L362
L364:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2019 == int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v2023 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2023&int32(1) == int32(0) {
		goto L363
	} else {
		goto L366
	}
L366:
	;
	v2028 = int32(_a_F_ReindexRelationConcurrently_14)
	v2030 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2031 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2030 + v2031
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2019)))
	*(*int32)(unsafe.Add(mBase, uint32(v2019))) = v2034 + v2031
	goto L368
L367:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v2019)))
	v2165 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2019))) = v2164 + v2165
	v2168 = int32(_a_F_ReindexRelationConcurrently_14)
	v2170 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2170 - v2165
	goto L363
L368:
	;
	v2043 = v2019 + int32(232)
	goto L369
L369:
	;
	v2049 = int32(0)
	v2052 = int32(0)
	goto L372
L372:
	;
	v2058 = int32(2)
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2007+v2052<<(uint(v2058)%32))))
	v2062 = int32(3)
	v2068 = *(*int64)(unsafe.Add(mBase, uint32(v2009+v2052<<(uint(v2062)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2043+v2061<<(uint(v2062)%32)))) = v2068
	v2071 = v2052 | int32(1)
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2007+v2071<<(uint(v2058)%32))))
	v2082 = *(*int64)(unsafe.Add(mBase, uint32(v2009+v2071<<(uint(v2062)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2043+v2075<<(uint(v2062)%32)))) = v2082
	v2085 = v2052 | v2058
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2007+v2085<<(uint(v2058)%32))))
	v2096 = *(*int64)(unsafe.Add(mBase, uint32(v2009+v2085<<(uint(v2062)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2043+v2089<<(uint(v2062)%32)))) = v2096
	v2099 = v2052 | v2062
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2007+v2099<<(uint(v2058)%32))))
	v2110 = *(*int64)(unsafe.Add(mBase, uint32(v2009+v2099<<(uint(v2062)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2043+v2103<<(uint(v2062)%32)))) = v2110
	v2112 = int32(4)
	v2115 = v2049 + v2112
	if v2115 != int32(4) {
		v2049 = v2115
		v2052 = v2052 + v2112
		goto L372
	} else {
		goto L374
	}
L373:
	;
	goto L367
L374:
	;
	goto L373
L379:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	v2192 = v1880 + int32(1)
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if v2192 < v2193 {
		v1880 = v2192
		goto L344
	} else {
		goto L382
	}
L382:
	;
	goto L345
L383:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2236 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	F_WaitForLockersMultiple(m, v1723, int32(5), int32(1))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L1
	} else {
		goto L388
	}
L385:
	;
	goto L384
L386:
	;
	v2240 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2240&int32(1) == int32(0) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v2245 = int32(_a_F_ReindexRelationConcurrently_14)
	v2247 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2248 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2247 + v2248
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2236)))
	*(*int32)(unsafe.Add(mBase, uint32(v2236))) = v2251 + v2248
	*(*int64)(unsafe.Add(mBase, uint32(v2236+int32(72))+232)) = int64(3)
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2236)))
	*(*int32)(unsafe.Add(mBase, uint32(v2236))) = v2259 + v2248
	v2265 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2265 - v2248
	goto L385
L388:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v2275 = int32(0)
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if v2276 <= v2275 {
		goto L337
	} else {
		goto L390
	}
L390:
	;
	v2279 = v2275
	goto L391
L391:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+12))
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v2314+v2279<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L1
	} else {
		goto L393
	}
L392:
	;
	goto L337
L393:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2322 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L1
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v2325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2318)+12)))
	if v2325 == int32(1) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	goto L396
L398:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v2333 = F_LWLockAcquire(m, v2329+int32(512), int32(0))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L1
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v2355 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L403
	}
L401:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v2337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2336)+124)))
	v2339 = v2337 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2336)+124)) = uint8(v2339)
	v2342 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2342)+12))
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2343+v2344))) = uint8(v2339)
	v2348 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v2348+int32(512))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	v2357 = F_RegisterSnapshot(m, v2355)
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	F_PushActiveSnapshot(m, v2357)
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+4))
	v2365 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2365 == int32(0) {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	v2401 = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = v2401
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = v2401
	v2405 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2318))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2405
	v2407 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2318)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2407
	v2411 = v38 + int32(224)
	v2413 = v38 + int32(192)
	goto L412
L407:
	;
	goto L406
L408:
	;
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2369&int32(1) == int32(0) {
		goto L407
	} else {
		goto L409
	}
L409:
	;
	v2374 = int32(_a_F_ReindexRelationConcurrently_14)
	v2376 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2377 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2376 + v2377
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2365)))
	*(*int32)(unsafe.Add(mBase, uint32(v2365))) = v2380 + v2377
	*(*int32)(unsafe.Add(mBase, uint32(v2365)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2365)+224)) = v2362
	base.MemoryFill(m, v2365+int32(232), int32(0), int32(160))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2365)))
	*(*int32)(unsafe.Add(mBase, uint32(v2365))) = v2391 + v2377
	v2397 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2397 - v2377
	goto L407
L410:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+4))
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2318)))
	F_validate_index(m, v2587, v2588, v2357)
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L1
	} else {
		goto L427
	}
L411:
	;
	goto L410
L412:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2423 == int32(0) {
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v2427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2427&int32(1) == int32(0) {
		goto L411
	} else {
		goto L414
	}
L414:
	;
	v2432 = int32(_a_F_ReindexRelationConcurrently_14)
	v2434 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2435 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2434 + v2435
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2423)))
	*(*int32)(unsafe.Add(mBase, uint32(v2423))) = v2438 + v2435
	goto L416
L415:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2423)))
	v2569 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2423))) = v2568 + v2569
	v2572 = int32(_a_F_ReindexRelationConcurrently_14)
	v2574 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2574 - v2569
	goto L411
L416:
	;
	v2447 = v2423 + int32(232)
	goto L417
L417:
	;
	v2453 = int32(0)
	v2456 = int32(0)
	goto L420
L420:
	;
	v2462 = int32(2)
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2411+v2456<<(uint(v2462)%32))))
	v2466 = int32(3)
	v2472 = *(*int64)(unsafe.Add(mBase, uint32(v2413+v2456<<(uint(v2466)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2447+v2465<<(uint(v2466)%32)))) = v2472
	v2475 = v2456 | int32(1)
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2411+v2475<<(uint(v2462)%32))))
	v2486 = *(*int64)(unsafe.Add(mBase, uint32(v2413+v2475<<(uint(v2466)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2447+v2479<<(uint(v2466)%32)))) = v2486
	v2489 = v2456 | v2462
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2411+v2489<<(uint(v2462)%32))))
	v2500 = *(*int64)(unsafe.Add(mBase, uint32(v2413+v2489<<(uint(v2466)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2447+v2493<<(uint(v2466)%32)))) = v2500
	v2503 = v2456 | v2466
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2411+v2503<<(uint(v2462)%32))))
	v2514 = *(*int64)(unsafe.Add(mBase, uint32(v2413+v2503<<(uint(v2466)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2447+v2507<<(uint(v2466)%32)))) = v2514
	v2516 = int32(4)
	v2519 = v2453 + v2516
	if v2519 != int32(4) {
		v2453 = v2519
		v2456 = v2456 + v2516
		goto L420
	} else {
		goto L422
	}
L421:
	;
	goto L415
L422:
	;
	goto L421
L427:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+4))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	F_UnregisterSnapshot(m, v2357)
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2604 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	F_WaitForOlderSnapshots(m, v2591, int32(1))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L1
	} else {
		goto L436
	}
L433:
	;
	goto L432
L434:
	;
	v2608 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2608&int32(1) == int32(0) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2613 = int32(_a_F_ReindexRelationConcurrently_14)
	v2615 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2616 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2615 + v2616
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v2604)))
	*(*int32)(unsafe.Add(mBase, uint32(v2604))) = v2619 + v2616
	*(*int64)(unsafe.Add(mBase, uint32(v2604+int32(72))+232)) = int64(7)
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2604)))
	*(*int32)(unsafe.Add(mBase, uint32(v2604))) = v2627 + v2616
	v2633 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2633 - v2616
	goto L433
L436:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v2643 = v2279 + int32(1)
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if v2643 < v2644 {
		v2279 = v2643
		goto L391
	} else {
		goto L438
	}
L438:
	;
	goto L392
L439:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2652 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	F_WaitForLockersMultiple(m, v1723, int32(5), int32(1))
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L1
	} else {
		goto L444
	}
L441:
	;
	goto L440
L442:
	;
	v2656 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2656&int32(1) == int32(0) {
		goto L441
	} else {
		goto L443
	}
L443:
	;
	v2661 = int32(_a_F_ReindexRelationConcurrently_14)
	v2663 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2664 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2663 + v2664
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2652)))
	*(*int32)(unsafe.Add(mBase, uint32(v2652))) = v2667 + v2664
	*(*int64)(unsafe.Add(mBase, uint32(v2652+int32(72))+232)) = int64(3)
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2652)))
	*(*int32)(unsafe.Add(mBase, uint32(v2652))) = v2675 + v2664
	v2681 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2681 - v2664
	goto L441
L444:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	goto L337
L446:
	;
	v2729 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v2733 = F_LWLockAcquire(m, v2729+int32(512), int32(0))
	mBase = m.M
	v2734 = m.ExcPending
	if v2734 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v2737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2736)+124)))
	v2739 = v2737 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2736)+124)) = uint8(v2739)
	v2742 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2742)+12))
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v2736)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2743+v2744))) = uint8(v2739)
	v2748 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v2748+int32(512))
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v2755 = int32(0)
	goto L451
L449:
	;
	F_MemoryContextDelete(m, v52)
	mBase = m.M
	v4234 = m.ExcPending
	if v4234 != 0 {
		goto L1
	} else {
		goto L718
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v79
	F_errmsg(m, v4153, v38+int32(32))
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L1
	} else {
		goto L714
	}
L451:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v2755 < v2790 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v4136 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L1
	} else {
		goto L712
	}
L453:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v2796 = v2792 + v2755<<(uint(int32(2))%32)
	goto L455
L454:
	;
	v2796 = int32(0)
	goto L455
L455:
	;
	if v1595 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	goto L452
L457:
	;
	v3335 = *(*int32)(unsafe.Add(mBase, uint32(v2804+v2755<<(uint(int32(2))%32))))
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(v2796)))
	v3338 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v3338 != 0 {
		goto L536
	} else {
		goto L537
	}
L458:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L1
	} else {
		goto L462
	}
L459:
	;
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if base.B2i32(v2796 == int32(0))|base.B2i32(v2801 <= v2755) != 0 {
		goto L458
	} else {
		goto L460
	}
L460:
	;
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+12))
	if v2804 != 0 {
		goto L457
	} else {
		goto L461
	}
L461:
	;
	goto L458
L462:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2814 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	F_WaitForLockersMultiple(m, v1723, int32(8), int32(1))
	mBase = m.M
	v2850 = m.ExcPending
	if v2850 != 0 {
		goto L1
	} else {
		goto L468
	}
L465:
	;
	goto L464
L466:
	;
	v2818 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2818&int32(1) == int32(0) {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v2823 = int32(_a_F_ReindexRelationConcurrently_14)
	v2825 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2826 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2825 + v2826
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2814)))
	*(*int32)(unsafe.Add(mBase, uint32(v2814))) = v2829 + v2826
	*(*int64)(unsafe.Add(mBase, uint32(v2814+int32(72))+232)) = int64(8)
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2814)))
	*(*int32)(unsafe.Add(mBase, uint32(v2814))) = v2837 + v2826
	v2843 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2843 - v2826
	goto L465
L468:
	;
	v2851 = int32(0)
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v2851 < v2852 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2856 = v2851
	goto L472
L470:
	;
	goto L471
L471:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L1
	} else {
		goto L489
	}
L472:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2890+v2856<<(uint(int32(2))%32))))
	v2896 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2896 != 0 {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	goto L471
L474:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2898 = m.ExcPending
	if v2898 != 0 {
		goto L1
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	v2899 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L478
	}
L477:
	;
	goto L476
L478:
	;
	F_PushActiveSnapshot(m, v2899)
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v2894)))
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v2894)+4))
	v2906 = F_table_open(m, v2904, int32(4))
	mBase = m.M
	v2907 = m.ExcPending
	if v2907 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v2909 = F_index_open(m, v2903, int32(4))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	F_TransferPredicateLocksToHeapRelation(m, v2909)
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_index_set_state_flags(m, v2903, int32(3))
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	F_CacheInvalidateRelcache(m, v2906)
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	F_relation_close(m, v2906, int32(0))
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	F_relation_close(m, v2909, int32(0))
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	v2927 = v2856 + int32(1)
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v2927 < v2928 {
		v2856 = v2927
		goto L472
	} else {
		goto L488
	}
L488:
	;
	goto L473
L489:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v2973 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2973 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	F_WaitForLockersMultiple(m, v1723, int32(8), int32(1))
	mBase = m.M
	v3009 = m.ExcPending
	if v3009 != 0 {
		goto L1
	} else {
		goto L495
	}
L492:
	;
	goto L491
L493:
	;
	v2977 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2977&int32(1) == int32(0) {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v2982 = int32(_a_F_ReindexRelationConcurrently_14)
	v2984 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2985 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2984 + v2985
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2973)))
	*(*int32)(unsafe.Add(mBase, uint32(v2973))) = v2988 + v2985
	*(*int64)(unsafe.Add(mBase, uint32(v2973+int32(72))+232)) = int64(9)
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v2973)))
	*(*int32)(unsafe.Add(mBase, uint32(v2973))) = v2996 + v2985
	v3002 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v3002 - v2985
	goto L492
L495:
	;
	v3010 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	F_PushActiveSnapshot(m, v3010)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v3014 = F_new_object_addresses(m)
	mBase = m.M
	v3015 = m.ExcPending
	if v3015 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if int32(0) < v3016 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v3021 = int32(0)
	goto L502
L500:
	;
	goto L501
L501:
	;
	v3109 = int32(0)
	F_performMultipleDeletions(m, v3014, v3109, int32(33))
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L1
	} else {
		goto L506
	}
L502:
	;
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v3055+v3021<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3059)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v3062
	F_add_exact_object_address(m, v38+int32(168), v3014)
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L1
	} else {
		goto L504
	}
L503:
	;
	goto L501
L504:
	;
	v3071 = v3021 + int32(1)
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v3071 < v3072 {
		v3021 = v3071
		goto L502
	} else {
		goto L505
	}
L505:
	;
	goto L503
L506:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3117 = m.ExcPending
	if v3117 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	if v1717 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L1
	} else {
		goto L516
	}
L510:
	;
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+4))
	if v3120 <= int32(0) {
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v3124 = v3109
	goto L512
L512:
	;
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+12))
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3158+v3124<<(uint(int32(2))%32))))
	F_UnlockRelationIdForSession(m, v3162, int32(4))
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L1
	} else {
		goto L514
	}
L513:
	;
	goto L509
L514:
	;
	v3167 = v3124 + int32(1)
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+4))
	if v3167 < v3168 {
		v3124 = v3167
		goto L512
	} else {
		goto L515
	}
L515:
	;
	goto L513
L516:
	;
	v3207 = int32(1)
	v3208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v3208&v3207 == int32(0) {
		goto L449
	} else {
		goto L517
	}
L517:
	;
	if v80 == int32(105) {
		goto L456
	} else {
		goto L518
	}
L518:
	;
	if v1595 == int32(0) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v3326 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3327 = m.ExcPending
	if v3327 != 0 {
		goto L1
	} else {
		goto L534
	}
L520:
	;
	v3217 = int32(0)
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if v3218 <= v3217 {
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v3222 = v3217
	goto L522
L522:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+12))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3256+v3222<<(uint(int32(2))%32))))
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3260)))
	v3264 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3265 = m.ExcPending
	if v3265 != 0 {
		goto L1
	} else {
		goto L524
	}
L523:
	;
	goto L519
L524:
	;
	if v3264 != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v3266 = F_get_rel_namespace(m, v3261)
	mBase = m.M
	v3267 = m.ExcPending
	if v3267 != 0 {
		goto L1
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v3286 = v3222 + int32(1)
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	if v3286 < v3287 {
		v3222 = v3286
		goto L522
	} else {
		goto L533
	}
L528:
	;
	v3268 = F_get_namespace_name(m, v3266)
	mBase = m.M
	v3269 = m.ExcPending
	if v3269 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v3270 = F_get_rel_name(m, v3261)
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v3270
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v3268
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_22), v38+int32(48))
	mBase = m.M
	v3278 = m.ExcPending
	if v3278 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(_a_F_ReindexRelationConcurrently_23), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	goto L527
L533:
	;
	goto L523
L534:
	;
	if v3326 == int32(0) {
		goto L449
	} else {
		goto L535
	}
L535:
	;
	v4153 = int32(_a_F_ReindexRelationConcurrently_24)
	v4177 = int32(_a_F_ReindexRelationConcurrently_25)
	goto L450
L536:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v3336)))
	v3342 = F_get_rel_name(m, v3341)
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L1
	} else {
		goto L540
	}
L539:
	;
	goto L538
L540:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(v3336)+4))
	v3347 = F_get_rel_namespace(m, v3346)
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v3350 = F_ChooseRelationName(m, v3342, int32(0), int32(_a_F_ReindexRelationConcurrently_26), v3347, int32(0))
	mBase = m.M
	v3351 = m.ExcPending
	if v3351 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v3352 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3353 = m.ExcPending
	if v3353 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	F_PushActiveSnapshot(m, v3352)
	mBase = m.M
	v3355 = m.ExcPending
	if v3355 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v3335)))
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v3336)))
	v3358 = m.G0
	v3360 = v3358 - int32(240)
	m.G0 = v3360
	v3363 = F_relation_open(m, v3357, int32(4))
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	v3366 = F_relation_open(m, v3356, int32(4))
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	v3370 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v3371 = m.ExcPending
	if v3371 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	v3374 = F_SearchSysCacheCopy(m, int32(57), v3357, int32(0))
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L553
	}
L548:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4126 = m.ExcPending
	if v4126 != 0 {
		goto L1
	} else {
		goto L709
	}
L549:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4113 = m.ExcPending
	if v4113 != 0 {
		goto L1
	} else {
		goto L706
	}
L550:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4098 = m.ExcPending
	if v4098 != 0 {
		goto L1
	} else {
		goto L703
	}
L551:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4083 = m.ExcPending
	if v4083 != 0 {
		goto L1
	} else {
		goto L700
	}
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4068 = m.ExcPending
	if v4068 != 0 {
		goto L1
	} else {
		goto L697
	}
L553:
	;
	if v3374 != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3378 = F_SearchSysCacheCopy(m, int32(57), v3356, int32(0))
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
		goto L1
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L1
	} else {
		goto L694
	}
L557:
	;
	if v3378 == int32(0) {
		goto L552
	} else {
		goto L558
	}
L558:
	;
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3378)+16))
	v3383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3382)+22)))
	v3384 = v3382 + v3383
	v3385 = int32(4)
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+16))
	v3388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3387)+22)))
	v3389 = v3387 + v3388
	v3391 = v3389 + v3385
	v3393 = F_strncpy(m, v3384+v3385, v3391, int32(64))
	mBase = m.M
	v3394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3393)+63)) = uint8(v3394)
	goto L559
L559:
	;
	v3397 = F_strncpy(m, v3391, v3350, int32(64))
	mBase = m.M
	v3398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3397)+63)) = uint8(v3398)
	goto L560
L560:
	;
	v3400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3384)+131)))
	v3401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3389)+131)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3384)+131)) = uint8(v3401)
	*(*uint8)(unsafe.Add(mBase, uint32(v3389)+131)) = uint8(v3400)
	F_CatalogTupleUpdate(m, v3370, v3374+int32(4), v3374)
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	F_CatalogTupleUpdate(m, v3370, v3378+int32(4), v3378)
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	F_pfree(m, v3374)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_pfree(m, v3378)
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	v3418 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v3419 = m.ExcPending
	if v3419 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	v3422 = F_SearchSysCacheCopy(m, int32(34), v3357, int32(0))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	if v3422 == int32(0) {
		goto L551
	} else {
		goto L567
	}
L567:
	;
	v3428 = F_SearchSysCacheCopy(m, int32(34), v3356, int32(0))
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	if v3428 == int32(0) {
		goto L550
	} else {
		goto L569
	}
L569:
	;
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v3428)+16))
	v3433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3432)+22)))
	v3434 = v3432 + v3433
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(v3422)+16))
	v3436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3435)+22)))
	v3437 = v3435 + v3436
	v3438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3437)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+14)) = uint8(v3438)
	v3440 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3437)+14)) = uint8(v3440)
	v3442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3437)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+15)) = uint8(v3442)
	*(*uint8)(unsafe.Add(mBase, uint32(v3437)+15)) = uint8(v3440)
	v3446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3437)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+16)) = uint8(v3446)
	v3448 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3437)+16)) = uint8(v3448)
	v3450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3437)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+22)) = uint8(v3450)
	v3452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3437)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+18)) = uint8(v3448)
	*(*uint8)(unsafe.Add(mBase, uint32(v3434)+17)) = uint8(v3452)
	*(*uint8)(unsafe.Add(mBase, uint32(v3437)+22)) = uint8(v3440)
	*(*uint16)(unsafe.Add(mBase, uint32(v3437)+17)) = uint16(v3440)
	F_CatalogTupleUpdate(m, v3418, v3422+int32(4), v3422)
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	F_CatalogTupleUpdate(m, v3418, v3428+int32(4), v3428)
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_pfree(m, v3422)
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	F_pfree(m, v3428)
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	v3472 = int32(0)
	v3473 = m.G0
	v3475 = v3473 - int32(144)
	m.G0 = v3475
	v3479 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	F_ScanKeyInit(m, v3475, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_ScanKeyInit(m, v3475+int32(48), int32(5), int32(3), int32(184), v3357)
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	F_ScanKeyInit(m, v3475+int32(96), int32(6), int32(3), int32(65), int32(0))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v3506 = F_systable_beginscan(m, v3479, int32(2674), int32(1), int32(0), int32(3), v3475)
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v3508 = F_systable_getnext(m, v3506)
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	if v3508 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v3510 = v3508
	v3514 = v3472
	goto L583
L581:
	;
	v3565 = v3472
	goto L582
L582:
	;
	F_systable_endscan(m, v3506)
	mBase = m.M
	v3597 = m.ExcPending
	if v3597 != 0 {
		goto L1
	} else {
		goto L592
	}
L583:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v3510)+16))
	v3546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3545)+22)))
	v3547 = v3545 + v3546
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3547)))
	if v3548 != int32(2606) {
		v3558 = v3514
		goto L585
	} else {
		goto L586
	}
L584:
	;
	v3565 = v3558
	goto L582
L585:
	;
	v3559 = F_systable_getnext(m, v3506)
	mBase = m.M
	v3560 = m.ExcPending
	if v3560 != 0 {
		goto L1
	} else {
		goto L590
	}
L586:
	;
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3547)+8))
	if v3551 != 0 {
		v3558 = v3514
		goto L585
	} else {
		goto L587
	}
L587:
	;
	v3552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3547)+24)))
	if v3552 != int32(110) {
		v3558 = v3514
		goto L585
	} else {
		goto L588
	}
L588:
	;
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(v3547)+4))
	v3556 = F_lappend_oid(m, v3514, v3555)
	mBase = m.M
	v3557 = m.ExcPending
	if v3557 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v3558 = v3556
	goto L585
L590:
	;
	if v3559 != 0 {
		v3510 = v3559
		v3514 = v3558
		goto L583
	} else {
		goto L591
	}
L591:
	;
	goto L584
L592:
	;
	F_relation_close(m, v3479, int32(1))
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	m.G0 = v3475 + int32(144)
	v3604 = F_get_index_constraint(m, v3357)
	mBase = m.M
	v3605 = m.ExcPending
	if v3605 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	if v3604 != 0 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v3606 = F_lappend_oid(m, v3565, v3604)
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L1
	} else {
		goto L598
	}
L596:
	;
	v3608 = v3565
	goto L597
L597:
	;
	v3611 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L1
	} else {
		goto L599
	}
L598:
	;
	v3608 = v3606
	goto L597
L599:
	;
	v3615 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v3616 = m.ExcPending
	if v3616 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	if v3608 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v3789 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3360)+88)) = v3789
	*(*int64)(unsafe.Add(mBase, uint32(v3360)+80)) = v3789
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+80)) = v3356
	v3796 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+72)) = v3796
	v3799 = v3360 + int32(96)
	F_ScanKeyInit(m, v3799, v3796, int32(3), int32(184), v3357)
	mBase = m.M
	v3804 = m.ExcPending
	if v3804 != 0 {
		goto L1
	} else {
		goto L627
	}
L602:
	;
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(v3608)+4))
	if v3619 <= int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3623 = int32(0)
	goto L604
L604:
	;
	v3659 = *(*int32)(unsafe.Add(mBase, uint32(v3608)+12))
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3659+v3623<<(uint(int32(2))%32))))
	v3665 = F_SearchSysCacheCopy(m, int32(19), v3663, int32(0))
	mBase = m.M
	v3666 = m.ExcPending
	if v3666 != 0 {
		goto L1
	} else {
		goto L606
	}
L605:
	;
	goto L601
L606:
	;
	if v3665 == int32(0) {
		goto L549
	} else {
		goto L607
	}
L607:
	;
	v3669 = *(*int32)(unsafe.Add(mBase, uint32(v3665)+16))
	v3670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3669)+22)))
	v3671 = v3669 + v3670
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3671)+88))
	if v3357 == v3672 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3671)+88)) = v3356
	F_CatalogTupleUpdate(m, v3611, v3665+int32(4), v3665)
	mBase = m.M
	v3678 = m.ExcPending
	if v3678 != 0 {
		goto L1
	} else {
		goto L611
	}
L609:
	;
	goto L610
L610:
	;
	F_pfree(m, v3665)
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L1
	} else {
		goto L612
	}
L611:
	;
	goto L610
L612:
	;
	v3682 = v3360 + int32(96)
	F_ScanKeyInit(m, v3682, int32(11), int32(3), int32(184), v3663)
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v3689 = int32(1)
	v3692 = F_systable_beginscan(m, v3615, int32(2699), v3689, int32(0), v3689, v3682)
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	goto L615
L615:
	;
	v3729 = F_systable_getnext(m, v3692)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L617
	}
L616:
	;
	F_systable_endscan(m, v3692)
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L1
	} else {
		goto L625
	}
L617:
	;
	if v3729 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+16))
	v3732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3731)+22)))
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3731+v3732)+88))
	if v3734 != v3357 {
		goto L615
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	goto L616
L621:
	;
	v3736 = F_heap_copytuple(m, v3729)
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	v3738 = *(*int32)(unsafe.Add(mBase, uint32(v3736)+16))
	v3739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3738)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3738+v3739)+88)) = v3356
	F_CatalogTupleUpdate(m, v3615, v3736+int32(4), v3736)
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	F_pfree(m, v3736)
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	goto L615
L625:
	;
	v3751 = v3623 + int32(1)
	v3752 = *(*int32)(unsafe.Add(mBase, uint32(v3608)+4))
	if v3751 < v3752 {
		v3623 = v3751
		goto L604
	} else {
		goto L626
	}
L626:
	;
	goto L605
L627:
	;
	F_ScanKeyInit(m, v3360+int32(144), int32(2), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3812 = m.ExcPending
	if v3812 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v3815 = int32(3)
	F_ScanKeyInit(m, v3360+int32(192), v3815, v3815, int32(65), int32(0))
	mBase = m.M
	v3820 = m.ExcPending
	if v3820 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	v3823 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v3824 = m.ExcPending
	if v3824 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	v3829 = F_systable_beginscan(m, v3823, int32(2675), int32(1), int32(0), int32(3), v3799)
	mBase = m.M
	v3830 = m.ExcPending
	if v3830 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	v3831 = F_systable_getnext(m, v3829)
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	if v3831 != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3833 = *(*int32)(unsafe.Add(mBase, uint32(v3823)+52))
	v3840 = F_heap_modify_tuple(m, v3831, v3833, v3360+int32(80), v3360+int32(76), v3360+int32(72))
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		goto L1
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	F_systable_endscan(m, v3829)
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	F_CatalogTupleUpdate(m, v3823, v3840+int32(4), v3840)
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	goto L635
L638:
	;
	F_relation_close(m, v3823, int32(0))
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	v3852 = F_get_rel_relispartition(m, v3357)
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	if v3852 != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v3854 = F_get_partition_ancestors(m, v3357)
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L1
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	F_changeDependenciesOf(m, v3356, v3357)
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L648
	}
L644:
	;
	v3856 = *(*int32)(unsafe.Add(mBase, uint32(v3854)+12))
	v3857 = *(*int32)(unsafe.Add(mBase, uint32(v3856)))
	v3858 = int32(0)
	v3860 = F_DeleteInheritsTuple(m, v3357, v3857, v3858, v3858)
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	F_StoreSingleInheritance(m, v3356, v3857, int32(1))
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	F_list_free(m, v3854)
	mBase = m.M
	v3866 = m.ExcPending
	if v3866 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	goto L643
L648:
	;
	F_changeDependenciesOn(m, v3356, v3357)
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	F_changeDependenciesOf(m, v3357, v3356)
	mBase = m.M
	v3874 = m.ExcPending
	if v3874 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	F_changeDependenciesOn(m, v3357, v3356)
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v3880 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3881 = *(*int32)(unsafe.Add(mBase, uint32(v3363)+48))
	v3882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3881)+117)))
	if v3882 != 0 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3883 = int32(0)
	goto L654
L653:
	;
	v3883 = v3880
	goto L654
L654:
	;
	v3884 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3363)+56)))
	v3885 = F_pgstat_fetch_entry(m, int32(2), v3883, v3884)
	mBase = m.M
	v3886 = m.ExcPending
	if v3886 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	if v3885 != 0 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3890 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(v3366)+48))
	v3892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3891)+117)))
	if v3892 != 0 {
		goto L659
	} else {
		goto L660
	}
L657:
	;
	goto L658
L658:
	;
	v3907 = m.G0
	v3909 = v3907 - int32(48)
	m.G0 = v3909
	v3913 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v3914 = m.ExcPending
	if v3914 != 0 {
		goto L1
	} else {
		goto L664
	}
L659:
	;
	v3893 = int32(0)
	goto L661
L660:
	;
	v3893 = v3890
	goto L661
L661:
	;
	v3894 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3366)+56)))
	v3896 = F_pgstat_get_entry_ref_locked(m, int32(2), v3893, v3894, int32(0))
	mBase = m.M
	v3897 = m.ExcPending
	if v3897 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3896)+4))
	base.MemoryCopy(m, v3898+int32(24), v3885, int32(216))
	F_pgstat_unlock_entry(m, v3896)
	mBase = m.M
	v3904 = m.ExcPending
	if v3904 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	goto L658
L664:
	;
	F_ScanKeyInit(m, v3909, int32(1), int32(3), int32(184), v3357)
	mBase = m.M
	v3919 = m.ExcPending
	if v3919 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v3921 = int32(1)
	v3924 = F_systable_beginscan(m, v3913, int32(2696), v3921, int32(0), v3921, v3909)
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L1
	} else {
		goto L667
	}
L666:
	;
	F_relation_close(m, v3913, int32(3))
	mBase = m.M
	v4027 = m.ExcPending
	if v4027 != 0 {
		goto L1
	} else {
		goto L687
	}
L667:
	;
	v3926 = F_systable_getnext(m, v3924)
	mBase = m.M
	v3927 = m.ExcPending
	if v3927 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	if v3926 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L669:
	;
	F_systable_endscan(m, v3924)
	mBase = m.M
	v3931 = m.ExcPending
	if v3931 != 0 {
		goto L1
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	v3936 = int32(0)
	v3945 = v3926
	goto L673
L672:
	;
	goto L666
L673:
	;
	v3967 = F_heap_copytuple(m, v3945)
	mBase = m.M
	v3968 = m.ExcPending
	if v3968 != 0 {
		goto L1
	} else {
		goto L675
	}
L674:
	;
	F_systable_endscan(m, v3924)
	mBase = m.M
	v3985 = m.ExcPending
	if v3985 != 0 {
		goto L1
	} else {
		goto L684
	}
L675:
	;
	v3969 = *(*int32)(unsafe.Add(mBase, uint32(v3967)+16))
	v3970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3969)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3969+v3970))) = v3356
	if v3936 == int32(0) {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v3975 = F_CatalogOpenIndexes(m, v3913)
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L1
	} else {
		goto L679
	}
L677:
	;
	v3977 = v3936
	goto L678
L678:
	;
	F_CatalogTupleInsertWithInfo(m, v3913, v3967, v3977)
	mBase = m.M
	v3979 = m.ExcPending
	if v3979 != 0 {
		goto L1
	} else {
		goto L680
	}
L679:
	;
	v3977 = v3975
	goto L678
L680:
	;
	F_pfree(m, v3967)
	mBase = m.M
	v3981 = m.ExcPending
	if v3981 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	v3982 = F_systable_getnext(m, v3924)
	mBase = m.M
	v3983 = m.ExcPending
	if v3983 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	if v3982 != 0 {
		v3936 = v3977
		v3945 = v3982
		goto L673
	} else {
		goto L683
	}
L683:
	;
	goto L674
L684:
	;
	if v3977 == int32(0) {
		goto L666
	} else {
		goto L685
	}
L685:
	;
	F_CatalogCloseIndexes(m, v3977)
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	goto L666
L687:
	;
	m.G0 = v3909 + int32(48)
	F_relation_close(m, v3370, int32(3))
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		goto L1
	} else {
		goto L688
	}
L688:
	;
	F_relation_close(m, v3418, int32(3))
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	F_relation_close(m, v3611, int32(3))
	mBase = m.M
	v4039 = m.ExcPending
	if v4039 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	F_relation_close(m, v3615, int32(3))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L1
	} else {
		goto L691
	}
L691:
	;
	F_relation_close(m, v3363, int32(0))
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	F_relation_close(m, v3366, int32(0))
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	m.G0 = v3360 + int32(240)
	goto L548
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360))) = v3357
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_27), v3360)
	mBase = m.M
	v4059 = m.ExcPending
	if v4059 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1585), int32(_a_F_ReindexRelationConcurrently_28))
	mBase = m.M
	v4064 = m.ExcPending
	if v4064 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+16)) = v3356
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_27), v3360+int32(16))
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1589), int32(_a_F_ReindexRelationConcurrently_28))
	mBase = m.M
	v4079 = m.ExcPending
	if v4079 != 0 {
		goto L1
	} else {
		goto L699
	}
L699:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L700:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+32)) = v3357
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_27), v3360+int32(32))
	mBase = m.M
	v4089 = m.ExcPending
	if v4089 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1615), int32(_a_F_ReindexRelationConcurrently_28))
	mBase = m.M
	v4094 = m.ExcPending
	if v4094 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+48)) = v3356
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_27), v3360+int32(48))
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1619), int32(_a_F_ReindexRelationConcurrently_28))
	mBase = m.M
	v4109 = m.ExcPending
	if v4109 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+64)) = v3663
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_29), v3360-int32(-64))
	mBase = m.M
	v4119 = m.ExcPending
	if v4119 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_17), int32(1683), int32(_a_F_ReindexRelationConcurrently_28))
	mBase = m.M
	v4124 = m.ExcPending
	if v4124 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L709:
	;
	v4127 = *(*int32)(unsafe.Add(mBase, uint32(v3336)+4))
	F_CacheInvalidateRelcacheByRelid(m, v4127)
	mBase = m.M
	v4129 = m.ExcPending
	if v4129 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4133 = m.ExcPending
	if v4133 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	v2755 = v2755 + int32(1)
	goto L451
L712:
	;
	if v4136 == int32(0) {
		goto L449
	} else {
		goto L713
	}
L713:
	;
	v4153 = int32(_a_F_ReindexRelationConcurrently_22)
	v4177 = int32(_a_F_ReindexRelationConcurrently_30)
	goto L450
L714:
	;
	v4186 = F_pg_rusage_show(m, v38+int32(248))
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v4186
	F_errdetail(m, int32(_a_F_ReindexRelationConcurrently_31), v38+int32(16))
	mBase = m.M
	v4193 = m.ExcPending
	if v4193 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), v4177, int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4197 = m.ExcPending
	if v4197 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	goto L449
L718:
	;
	v4237 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v4237 == int32(0) {
		goto L720
	} else {
		goto L721
	}
L719:
	;
	v4276 = v3207
	goto L12
L720:
	;
	goto L719
L721:
	;
	v4241 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v4241&int32(1) == int32(0) {
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v4246 = *(*int32)(unsafe.Add(mBase, uint32(v4237)+220))
	if v4246 == int32(0) {
		goto L720
	} else {
		goto L723
	}
L723:
	;
	v4249 = int32(_a_F_ReindexRelationConcurrently_14)
	v4251 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v4252 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v4251 + v4252
	v4255 = *(*int32)(unsafe.Add(mBase, uint32(v4237)))
	*(*int32)(unsafe.Add(mBase, uint32(v4237))) = v4255 + v4252
	v4259 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4237)+220)) = v4259
	*(*int32)(unsafe.Add(mBase, uint32(v4237)+224)) = v4259
	*(*int32)(unsafe.Add(mBase, uint32(v4237))) = v4255 + int32(2)
	v4269 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v4269 - v4252
	goto L720
L724:
	;
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_32), int32(0))
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3939), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L727:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v4332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v4333 = F_get_tablespace_name(m, v4332)
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v4333
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_33), v38)
	mBase = m.M
	v4338 = m.ExcPending
	if v4338 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3864), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4343 = m.ExcPending
	if v4343 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
