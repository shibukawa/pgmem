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
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v854 int64
	_ = v854
	var v856 int64
	_ = v856
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v925 int64
	_ = v925
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v939 int64
	_ = v939
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v953 int64
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v967 int64
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
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
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
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
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1224 int32
	_ = v1224
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1311 int32
	_ = v1311
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int64
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int64
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1585 int64
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1590 int64
	_ = v1590
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1611 int32
	_ = v1611
	var v1618 int32
	_ = v1618
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int64
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1733 int32
	_ = v1733
	var v1739 int32
	_ = v1739
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1887 int32
	_ = v1887
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2025 int32
	_ = v2025
	var v2033 int64
	_ = v2033
	var v2035 int64
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2079 int32
	_ = v2079
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2104 int64
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2118 int64
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2132 int64
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2146 int64
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2151 int32
	_ = v2151
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2271 int32
	_ = v2271
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2313 int32
	_ = v2313
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2377 int32
	_ = v2377
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2453 int32
	_ = v2453
	var v2457 int64
	_ = v2457
	var v2461 int64
	_ = v2461
	var v2463 int64
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2479 int32
	_ = v2479
	var v2483 int32
	_ = v2483
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2507 int32
	_ = v2507
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2532 int64
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2546 int64
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2560 int64
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2574 int64
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2668 int32
	_ = v2668
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2761 int32
	_ = v2761
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2807 int32
	_ = v2807
	var v2809 int32
	_ = v2809
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2828 int32
	_ = v2828
	var v2832 int32
	_ = v2832
	var v2835 int32
	_ = v2835
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2884 int32
	_ = v2884
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2894 int32
	_ = v2894
	var v2898 int32
	_ = v2898
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2931 int32
	_ = v2931
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3008 int32
	_ = v3008
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3061 int32
	_ = v3061
	var v3065 int32
	_ = v3065
	var v3070 int32
	_ = v3070
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3076 int32
	_ = v3076
	var v3080 int32
	_ = v3080
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3098 int32
	_ = v3098
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3117 int32
	_ = v3117
	var v3151 int32
	_ = v3151
	var v3155 int32
	_ = v3155
	var v3158 int32
	_ = v3158
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3205 int32
	_ = v3205
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3254 int32
	_ = v3254
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3318 int32
	_ = v3318
	var v3352 int32
	_ = v3352
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3374 int32
	_ = v3374
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3434 int32
	_ = v3434
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3456 int32
	_ = v3456
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3503 int32
	_ = v3503
	var v3507 int32
	_ = v3507
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3536 int32
	_ = v3536
	var v3538 int32
	_ = v3538
	var v3542 int32
	_ = v3542
	var v3544 int32
	_ = v3544
	var v3546 int32
	_ = v3546
	var v3548 int32
	_ = v3548
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3571 int32
	_ = v3571
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3582 int32
	_ = v3582
	var v3589 int32
	_ = v3589
	var v3597 int32
	_ = v3597
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3610 int32
	_ = v3610
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3651 int32
	_ = v3651
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
	var v3661 int32
	_ = v3661
	var v3693 int32
	_ = v3693
	var v3696 int32
	_ = v3696
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
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3711 int32
	_ = v3711
	var v3712 int32
	_ = v3712
	var v3715 int32
	_ = v3715
	var v3719 int32
	_ = v3719
	var v3755 int32
	_ = v3755
	var v3759 int32
	_ = v3759
	var v3761 int32
	_ = v3761
	var v3762 int32
	_ = v3762
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3774 int32
	_ = v3774
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3783 int32
	_ = v3783
	var v3785 int32
	_ = v3785
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3841 int32
	_ = v3841
	var v3843 int32
	_ = v3843
	var v3845 int32
	_ = v3845
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3885 int64
	_ = v3885
	var v3892 int32
	_ = v3892
	var v3895 int32
	_ = v3895
	var v3900 int32
	_ = v3900
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3916 int32
	_ = v3916
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3941 int32
	_ = v3941
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3948 int32
	_ = v3948
	var v3949 int32
	_ = v3949
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3952 int32
	_ = v3952
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3960 int32
	_ = v3960
	var v3962 int32
	_ = v3962
	var v3966 int32
	_ = v3966
	var v3968 int32
	_ = v3968
	var v3970 int32
	_ = v3970
	var v3972 int32
	_ = v3972
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3979 int32
	_ = v3979
	var v3980 int64
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3986 int32
	_ = v3986
	var v3987 int32
	_ = v3987
	var v3988 int32
	_ = v3988
	var v3989 int32
	_ = v3989
	var v3990 int64
	_ = v3990
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v4000 int32
	_ = v4000
	var v4003 int32
	_ = v4003
	var v4005 int32
	_ = v4005
	var v4009 int32
	_ = v4009
	var v4010 int32
	_ = v4010
	var v4015 int32
	_ = v4015
	var v4017 int32
	_ = v4017
	var v4020 int32
	_ = v4020
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4027 int32
	_ = v4027
	var v4032 int32
	_ = v4032
	var v4041 int32
	_ = v4041
	var v4063 int32
	_ = v4063
	var v4064 int32
	_ = v4064
	var v4065 int32
	_ = v4065
	var v4066 int32
	_ = v4066
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4081 int32
	_ = v4081
	var v4085 int32
	_ = v4085
	var v4123 int32
	_ = v4123
	var v4129 int32
	_ = v4129
	var v4132 int32
	_ = v4132
	var v4135 int32
	_ = v4135
	var v4138 int32
	_ = v4138
	var v4141 int32
	_ = v4141
	var v4144 int32
	_ = v4144
	var v4151 int32
	_ = v4151
	var v4155 int32
	_ = v4155
	var v4160 int32
	_ = v4160
	var v4164 int32
	_ = v4164
	var v4170 int32
	_ = v4170
	var v4175 int32
	_ = v4175
	var v4179 int32
	_ = v4179
	var v4185 int32
	_ = v4185
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4200 int32
	_ = v4200
	var v4205 int32
	_ = v4205
	var v4209 int32
	_ = v4209
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4229 int32
	_ = v4229
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4249 int32
	_ = v4249
	var v4273 int32
	_ = v4273
	var v4279 int32
	_ = v4279
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4289 int32
	_ = v4289
	var v4293 int32
	_ = v4293
	var v4330 int32
	_ = v4330
	var v4333 int32
	_ = v4333
	var v4337 int32
	_ = v4337
	var v4342 int32
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4351 int32
	_ = v4351
	var v4355 int32
	_ = v4355
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4373 int32
	_ = v4373
	var v4380 int32
	_ = v4380
	var v4419 int32
	_ = v4419
	var v4423 int32
	_ = v4423
	var v4428 int32
	_ = v4428
	var v4432 int32
	_ = v4432
	var v4435 int32
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4438 int32
	_ = v4438
	var v4442 int32
	_ = v4442
	var v4447 int32
	_ = v4447
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
	v4432 = m.ExcPending
	if v4432 != 0 {
		goto L1
	} else {
		goto L727
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4419 = m.ExcPending
	if v4419 != 0 {
		goto L1
	} else {
		goto L724
	}
L12:
	;
	m.G0 = v38 + int32(416)
	return v4380
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
	v4380 = v4
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
		v4380 = v4
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
	v4380 = v4
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
	v4380 = int32(0)
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
	v1611 = v702
	v1618 = v702
	goto L173
L173:
	;
	if v681 == int32(0) {
		v1733 = v1618
		v1739 = v4
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
	v1611 = v1541
	v1618 = v1555
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
	v854 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v747))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v854
	v856 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v747)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v856
	v860 = v38 + int32(224)
	v862 = v38 + int32(192)
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
	v825 = int32(0)
	v827 = int32(_a_F_ReindexRelationConcurrently_15)
	v828 = base.AtomicRmwOr32(m, v825, v827, v825)
	*(*int32)(unsafe.Add(mBase, uint32(v806)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v806)+224)) = v794
	base.MemoryFill(m, v806+int32(232), v825, int32(160))
	v839 = base.AtomicRmwOr32(m, v825, v827, v825)
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v840 + v818
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v846 - v818
	goto L189
L192:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1045 = F_get_rel_name(m, v1044)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L209
	}
L193:
	;
	goto L192
L194:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v872 == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v876&int32(1) == int32(0) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v881 = int32(_a_F_ReindexRelationConcurrently_14)
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v884 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v883 + v884
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v887 + v884
	v891 = int32(0)
	v894 = base.AtomicRmwOr32(m, v891, int32(_a_F_ReindexRelationConcurrently_15), v891)
	goto L198
L197:
	;
	v1021 = int32(0)
	v1024 = base.AtomicRmwOr32(m, v1021, int32(_a_F_ReindexRelationConcurrently_15), v1021)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	v1026 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v1025 + v1026
	v1029 = int32(_a_F_ReindexRelationConcurrently_14)
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1031 - v1026
	goto L193
L198:
	;
	v900 = v872 + int32(232)
	goto L199
L199:
	;
	v906 = int32(0)
	v909 = int32(0)
	goto L202
L202:
	;
	v915 = int32(2)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v860+v909<<(uint(v915)%32))))
	v919 = int32(3)
	v925 = *(*int64)(unsafe.Add(mBase, uint32(v862+v909<<(uint(v919)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v900+v918<<(uint(v919)%32)))) = v925
	v928 = v909 | int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v860+v928<<(uint(v915)%32))))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v862+v928<<(uint(v919)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v900+v932<<(uint(v919)%32)))) = v939
	v942 = v909 | v915
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v860+v942<<(uint(v915)%32))))
	v953 = *(*int64)(unsafe.Add(mBase, uint32(v862+v942<<(uint(v919)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v900+v946<<(uint(v919)%32)))) = v953
	v956 = v909 | v919
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v860+v956<<(uint(v915)%32))))
	v967 = *(*int64)(unsafe.Add(mBase, uint32(v862+v956<<(uint(v919)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v900+v960<<(uint(v919)%32)))) = v967
	v969 = int32(4)
	v972 = v906 + v969
	if v972 != int32(4) {
		v906 = v972
		v909 = v909 + v969
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
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v750)+192))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+4))
	v1051 = F_get_rel_namespace(m, v1050)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1054 = F_ChooseRelationName(m, v1045, int32(0), int32(_a_F_ReindexRelationConcurrently_16), v1051, int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1056 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1065 = int32(0)
	v1068 = m.G0
	v1070 = v1068 - int32(48)
	m.G0 = v1070
	v1073 = F_index_open(m, v1064, int32(3))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L221
	}
L213:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v755)+48))
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+119)))
	if v1058 != int32(116) {
		v1063 = v1056
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v750)+48))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+92))
	v1063 = v1062
	goto L212
L216:
	;
	goto L215
L217:
	;
	v1525 = F_index_open(m, v1452, int32(4))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L292
	}
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L289
	}
L219:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L286
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L1
	} else {
		goto L283
	}
L221:
	;
	v1075 = F_BuildIndexInfo(m, v1073)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+92))
	if v1077 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1081 = F_SearchSysCache1(m, int32(34), v1064)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
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
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L279
	}
L226:
	;
	if v1081 == int32(0) {
		goto L220
	} else {
		goto L227
	}
L227:
	;
	v1087 = F_SysCacheGetAttrNotNull(m, int32(34), v1081, int32(18))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v1091 = F_SysCacheGetAttrNotNull(m, int32(34), v1081, int32(19))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v1094 = F_SearchSysCache1(m, int32(57), v1064)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if v1094 == int32(0) {
		goto L219
	} else {
		goto L231
	}
L231:
	;
	v1102 = F_SysCacheGetAttr(m, int32(57), v1094, int32(33), v1070+int32(47))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+76))
	if v1104 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1107 = F_SysCacheGetAttrNotNull(m, int32(34), v1081, int32(20))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	v1116 = v1065
	goto L235
L235:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+84))
	if v1117 != 0 {
		goto L240
	} else {
		goto L241
	}
L236:
	;
	v1109 = F_text_to_cstring(m, v1107)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1111 = F_stringToNode(m, v1109)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_pfree(m, v1109)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1116 = v1111
	goto L235
L240:
	;
	v1120 = F_SysCacheGetAttrNotNull(m, int32(34), v1081, int32(21))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L243
	}
L241:
	;
	v1131 = v1065
	goto L242
L242:
	;
	v1132 = int32(0)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+8))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+132))
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+116)))
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+117)))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+204))
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+28)))
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+124)))
	v1143 = F_makeIndexInfo(m, v1133, v1134, v1135, v1116, v1131, v1136, v1137, v1132, int32(1), v1141, v1142)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L248
	}
L243:
	;
	v1122 = F_text_to_cstring(m, v1120)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1124 = F_stringToNode(m, v1122)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1126 = F_make_ands_implicit(m, v1124)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_pfree(m, v1122)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1131 = v1126
	goto L242
L248:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	if int32(0) < v1145 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1148 = int32(12)
	v1165 = int32(0)
	v1167 = v1065
	goto L252
L250:
	;
	v1224 = v1065
	goto L251
L251:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	v1248 = F_palloc0(m, v1245<<(uint(int32(2))%32))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L256
	}
L252:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+52))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1188)))
	v1198 = F_lappend(m, v1167, v1188+v1189<<(uint(int32(4))%32)+v1165*int32(100)+int32(24))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L254
	}
L253:
	;
	v1224 = v1198
	goto L251
L254:
	;
	v1200 = int32(1)
	v1201 = v1165 << (uint(v1200) % 32)
	v1204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1075+v1148+v1201))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1143+v1148+v1201))) = uint16(v1204)
	v1207 = v1165 + v1200
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	if v1207 < v1208 {
		v1165 = v1207
		v1167 = v1198
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	if int32(0) < v1250 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1256 = v1132
	goto L260
L258:
	;
	v1311 = v1250
	goto L259
L259:
	;
	v1336 = F_palloc0(m, v1311<<(uint(int32(3))%32))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L264
	}
L260:
	;
	v1292 = v1256 + int32(1)
	v1294 = F_get_attoptions(m, v1064, base.I32_extend16_s(v1292))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L262
	}
L261:
	;
	v1311 = v1297
	goto L259
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1248+v1256<<(uint(int32(2))%32)))) = v1294
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	if v1292 < v1297 {
		v1256 = v1292
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1338 = int32(0)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	if v1338 < v1339 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1345 = v1338
	goto L268
L266:
	;
	goto L267
L267:
	;
	v1436 = int32(0)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+48))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+84))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+248))
	v1443 = int32(24)
	v1452 = F_index_create(m, v755, v1054, v1436, v1436, v1436, v1436, v1143, v1224, v1441, v1063, v1442, v1087+v1443, v1248, v1091+v1443, v1336, v1102, int32(12), v1436, int32(1), v1436, v1436)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L275
	}
L268:
	;
	v1379 = v1345 + int32(1)
	v1381 = F_SearchSysCache2(m, int32(7), v1064, base.I32_extend16_s(v1379))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L270
	}
L269:
	;
	goto L267
L270:
	;
	if v1381 == int32(0) {
		goto L218
	} else {
		goto L271
	}
L271:
	;
	v1389 = F_SysCacheGetAttr(m, int32(7), v1381, int32(21), v1070+int32(47))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_ReleaseCatCache(m, v1381)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1395 = v1336 + v1345<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1395))) = v1389
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070)+47)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1395)+4)) = uint8(v1397)
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	if v1379 < v1399 {
		v1345 = v1379
		goto L268
	} else {
		goto L274
	}
L274:
	;
	goto L269
L275:
	;
	F_relation_close(m, v1073, int32(0))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_ReleaseCatCache(m, v1081)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_ReleaseCatCache(m, v1094)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	m.G0 = v1070 + int32(48)
	goto L217
L279:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_17), int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1333), int32(_a_F_ReindexRelationConcurrently_19))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1070))) = v1064
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_20), v1070)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1338), int32(_a_F_ReindexRelationConcurrently_19))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1070)+16)) = v1064
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_21), v1070+int32(16))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1350), int32(_a_F_ReindexRelationConcurrently_19))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1070)+36)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(v1070)+32)) = v1379
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_22), v1070+int32(32))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1431), int32(_a_F_ReindexRelationConcurrently_19))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
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
	v1527 = int32(_a_F_ReindexRelationConcurrently_2)
	v1528 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v1532 = F_palloc(m, int32(16))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1532))) = v1452
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1532)+12)) = uint8(v1535)
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+4)) = v1537
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1532)+8)) = v1539
	v1541 = F_lappend(m, v717, v1532)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1544 = F_palloc(m, int32(8))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1546 = *(*int64)(unsafe.Add(mBase, uint32(v750)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1544))) = v1546
	v1548 = F_lappend(m, v724, v1544)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1551 = F_palloc(m, int32(8))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1553 = *(*int64)(unsafe.Add(mBase, uint32(v1525)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1551))) = v1553
	v1555 = F_lappend(m, v1548, v1551)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1528
	F_relation_close(m, v750, int32(0))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_relation_close(m, v1525, int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	F_AtEOXact_GUC(m, int32(0), v781)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v38)+184))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5])) = v1569
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4])) = v1568
	goto L302
L302:
	;
	F_relation_close(m, v755, int32(0))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v1580 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = v1580
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v1580
	v1585 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v1585
	v1588 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v1588
	v1590 = *(*int64)(unsafe.Add(mBase, uint32(v38)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+80)) = v1590
	F_EventTriggerCollectSimpleCommand(m, v38+int32(80), v38-int32(-64), l0)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1599 = v715 + int32(1)
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v1599 < v1600 {
		v715 = v1599
		v717 = v1541
		v724 = v1555
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
	if v1733 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L310:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v1639 <= int32(0) {
		v1733 = v1618
		v1739 = v4
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1643 = int32(0)
	v1659 = v1618
	v1665 = v4
	goto L312
L312:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1678+v1643<<(uint(int32(2))%32))))
	v1684 = F_table_open(m, v1682, int32(4))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L1
	} else {
		goto L314
	}
L313:
	;
	v1733 = v1695
	v1739 = v1706
	goto L309
L314:
	;
	v1686 = int32(_a_F_ReindexRelationConcurrently_2)
	v1687 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v52
	v1691 = F_palloc(m, int32(8))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1693 = *(*int64)(unsafe.Add(mBase, uint32(v1684)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1691))) = v1693
	v1695 = F_lappend(m, v1659, v1691)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1698 = F_palloc(m, int32(16))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1698))) = v1700
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1691)))
	*(*int64)(unsafe.Add(mBase, uint32(v1698)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v1698)+4)) = v1702
	v1706 = F_lappend(m, v1665, v1698)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1687
	F_relation_close(m, v1684, int32(0))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1714 = v1643 + int32(1)
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	if v1714 < v1715 {
		v1643 = v1714
		v1659 = v1695
		v1665 = v1706
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
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L328
	}
L322:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+4))
	if v1754 <= int32(0) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1759 = int32(0)
	goto L324
L324:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+12))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1793+v1759<<(uint(int32(2))%32))))
	F_LockRelationIdForSession(m, v1797, int32(4))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L1
	} else {
		goto L326
	}
L325:
	;
	goto L321
L326:
	;
	v1802 = v1759 + int32(1)
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+4))
	if v1802 < v1803 {
		v1759 = v1802
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
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1850 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	F_WaitForLockersMultiple(m, v1739, int32(5), int32(1))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L335
	}
L332:
	;
	goto L331
L333:
	;
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1854&int32(1) == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1859 = int32(_a_F_ReindexRelationConcurrently_14)
	v1861 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1862 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1861 + v1862
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1850)))
	*(*int32)(unsafe.Add(mBase, uint32(v1850))) = v1865 + v1862
	v1869 = int32(0)
	v1871 = int32(_a_F_ReindexRelationConcurrently_15)
	v1872 = base.AtomicRmwOr32(m, v1869, v1871, v1869)
	*(*int64)(unsafe.Add(mBase, uint32(v1850+int32(72))+232)) = int64(1)
	v1880 = base.AtomicRmwOr32(m, v1869, v1871, v1869)
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1850)))
	*(*int32)(unsafe.Add(mBase, uint32(v1850))) = v1881 + v1862
	v1887 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1887 - v1862
	goto L332
L335:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	if v1611 != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L1
	} else {
		goto L446
	}
L338:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if int32(0) < v1897 {
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
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L1
	} else {
		goto L439
	}
L341:
	;
	v1904 = int32(0)
	goto L344
L342:
	;
	goto L343
L343:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L1
	} else {
		goto L383
	}
L344:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1936+v1904<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L1
	} else {
		goto L346
	}
L345:
	;
	goto L343
L346:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v1944 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1940)+12)))
	if v1947 == int32(1) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L349
L351:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v1955 = F_LWLockAcquire(m, v1951+int32(512), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	v1977 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L356
	}
L354:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1958)+124)))
	v1961 = v1959 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1958)+124)) = uint8(v1961)
	v1964 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1964)+12))
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1958)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v1965+v1966))) = uint8(v1961)
	v1970 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v1970+int32(512))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	F_PushActiveSnapshot(m, v1977)
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+4))
	v1985 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1985 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = int64(4)
	v2033 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1940))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2033
	v2035 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1940)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2035
	v2039 = v38 + int32(224)
	v2041 = v38 + int32(192)
	goto L364
L359:
	;
	goto L358
L360:
	;
	v1989 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1989&int32(1) == int32(0) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1994 = int32(_a_F_ReindexRelationConcurrently_14)
	v1996 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1997 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1996 + v1997
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1985)))
	*(*int32)(unsafe.Add(mBase, uint32(v1985))) = v2000 + v1997
	v2004 = int32(0)
	v2006 = int32(_a_F_ReindexRelationConcurrently_15)
	v2007 = base.AtomicRmwOr32(m, v2004, v2006, v2004)
	*(*int32)(unsafe.Add(mBase, uint32(v1985)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1985)+224)) = v1982
	base.MemoryFill(m, v1985+int32(232), v2004, int32(160))
	v2018 = base.AtomicRmwOr32(m, v2004, v2006, v2004)
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v1985)))
	*(*int32)(unsafe.Add(mBase, uint32(v1985))) = v2019 + v1997
	v2025 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2025 - v1997
	goto L359
L362:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+4))
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v1940)))
	F_index_concurrently_build(m, v2223, v2224)
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L379
	}
L363:
	;
	goto L362
L364:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2051 == int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2055&int32(1) == int32(0) {
		goto L363
	} else {
		goto L366
	}
L366:
	;
	v2060 = int32(_a_F_ReindexRelationConcurrently_14)
	v2062 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2063 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2062 + v2063
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2051)))
	*(*int32)(unsafe.Add(mBase, uint32(v2051))) = v2066 + v2063
	v2070 = int32(0)
	v2073 = base.AtomicRmwOr32(m, v2070, int32(_a_F_ReindexRelationConcurrently_15), v2070)
	goto L368
L367:
	;
	v2200 = int32(0)
	v2203 = base.AtomicRmwOr32(m, v2200, int32(_a_F_ReindexRelationConcurrently_15), v2200)
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2051)))
	v2205 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2051))) = v2204 + v2205
	v2208 = int32(_a_F_ReindexRelationConcurrently_14)
	v2210 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2210 - v2205
	goto L363
L368:
	;
	v2079 = v2051 + int32(232)
	goto L369
L369:
	;
	v2085 = int32(0)
	v2088 = int32(0)
	goto L372
L372:
	;
	v2094 = int32(2)
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2039+v2088<<(uint(v2094)%32))))
	v2098 = int32(3)
	v2104 = *(*int64)(unsafe.Add(mBase, uint32(v2041+v2088<<(uint(v2098)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2079+v2097<<(uint(v2098)%32)))) = v2104
	v2107 = v2088 | int32(1)
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2039+v2107<<(uint(v2094)%32))))
	v2118 = *(*int64)(unsafe.Add(mBase, uint32(v2041+v2107<<(uint(v2098)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2079+v2111<<(uint(v2098)%32)))) = v2118
	v2121 = v2088 | v2094
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2039+v2121<<(uint(v2094)%32))))
	v2132 = *(*int64)(unsafe.Add(mBase, uint32(v2041+v2121<<(uint(v2098)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2079+v2125<<(uint(v2098)%32)))) = v2132
	v2135 = v2088 | v2098
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2039+v2135<<(uint(v2094)%32))))
	v2146 = *(*int64)(unsafe.Add(mBase, uint32(v2041+v2135<<(uint(v2098)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2079+v2139<<(uint(v2098)%32)))) = v2146
	v2148 = int32(4)
	v2151 = v2085 + v2148
	if v2151 != int32(4) {
		v2085 = v2151
		v2088 = v2088 + v2148
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
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	v2232 = v1904 + int32(1)
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v2232 < v2233 {
		v1904 = v2232
		goto L344
	} else {
		goto L382
	}
L382:
	;
	goto L345
L383:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2276 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	F_WaitForLockersMultiple(m, v1739, int32(5), int32(1))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L1
	} else {
		goto L388
	}
L385:
	;
	goto L384
L386:
	;
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2280&int32(1) == int32(0) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v2285 = int32(_a_F_ReindexRelationConcurrently_14)
	v2287 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2288 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2287 + v2288
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2276)))
	*(*int32)(unsafe.Add(mBase, uint32(v2276))) = v2291 + v2288
	v2295 = int32(0)
	v2297 = int32(_a_F_ReindexRelationConcurrently_15)
	v2298 = base.AtomicRmwOr32(m, v2295, v2297, v2295)
	*(*int64)(unsafe.Add(mBase, uint32(v2276+int32(72))+232)) = int64(3)
	v2306 = base.AtomicRmwOr32(m, v2295, v2297, v2295)
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v2276)))
	*(*int32)(unsafe.Add(mBase, uint32(v2276))) = v2307 + v2288
	v2313 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2313 - v2288
	goto L385
L388:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v2323 = int32(0)
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v2324 <= v2323 {
		goto L337
	} else {
		goto L390
	}
L390:
	;
	v2327 = v2323
	goto L391
L391:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v2362+v2327<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L393
	}
L392:
	;
	goto L337
L393:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2370 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L1
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v2373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2366)+12)))
	if v2373 == int32(1) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	goto L396
L398:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v2381 = F_LWLockAcquire(m, v2377+int32(512), int32(0))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L1
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v2403 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L1
	} else {
		goto L403
	}
L401:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2384)+124)))
	v2387 = v2385 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2384)+124)) = uint8(v2387)
	v2390 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+12))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2384)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2391+v2392))) = uint8(v2387)
	v2396 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v2396+int32(512))
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	v2405 = F_RegisterSnapshot(m, v2403)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	F_PushActiveSnapshot(m, v2405)
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v2366)+4))
	v2413 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2413 == int32(0) {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	v2457 = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = v2457
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = v2457
	v2461 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2366))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2461
	v2463 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2366)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2463
	v2467 = v38 + int32(224)
	v2469 = v38 + int32(192)
	goto L412
L407:
	;
	goto L406
L408:
	;
	v2417 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2417&int32(1) == int32(0) {
		goto L407
	} else {
		goto L409
	}
L409:
	;
	v2422 = int32(_a_F_ReindexRelationConcurrently_14)
	v2424 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2425 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2424 + v2425
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2413)))
	*(*int32)(unsafe.Add(mBase, uint32(v2413))) = v2428 + v2425
	v2432 = int32(0)
	v2434 = int32(_a_F_ReindexRelationConcurrently_15)
	v2435 = base.AtomicRmwOr32(m, v2432, v2434, v2432)
	*(*int32)(unsafe.Add(mBase, uint32(v2413)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2413)+224)) = v2410
	base.MemoryFill(m, v2413+int32(232), v2432, int32(160))
	v2446 = base.AtomicRmwOr32(m, v2432, v2434, v2432)
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2413)))
	*(*int32)(unsafe.Add(mBase, uint32(v2413))) = v2447 + v2425
	v2453 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2453 - v2425
	goto L407
L410:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2366)+4))
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v2366)))
	F_validate_index(m, v2651, v2652, v2405)
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L1
	} else {
		goto L427
	}
L411:
	;
	goto L410
L412:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2479 == int32(0) {
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v2483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2483&int32(1) == int32(0) {
		goto L411
	} else {
		goto L414
	}
L414:
	;
	v2488 = int32(_a_F_ReindexRelationConcurrently_14)
	v2490 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2491 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2490 + v2491
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v2479)))
	*(*int32)(unsafe.Add(mBase, uint32(v2479))) = v2494 + v2491
	v2498 = int32(0)
	v2501 = base.AtomicRmwOr32(m, v2498, int32(_a_F_ReindexRelationConcurrently_15), v2498)
	goto L416
L415:
	;
	v2628 = int32(0)
	v2631 = base.AtomicRmwOr32(m, v2628, int32(_a_F_ReindexRelationConcurrently_15), v2628)
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2479)))
	v2633 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2479))) = v2632 + v2633
	v2636 = int32(_a_F_ReindexRelationConcurrently_14)
	v2638 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2638 - v2633
	goto L411
L416:
	;
	v2507 = v2479 + int32(232)
	goto L417
L417:
	;
	v2513 = int32(0)
	v2516 = int32(0)
	goto L420
L420:
	;
	v2522 = int32(2)
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2516<<(uint(v2522)%32))))
	v2526 = int32(3)
	v2532 = *(*int64)(unsafe.Add(mBase, uint32(v2469+v2516<<(uint(v2526)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2507+v2525<<(uint(v2526)%32)))) = v2532
	v2535 = v2516 | int32(1)
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2535<<(uint(v2522)%32))))
	v2546 = *(*int64)(unsafe.Add(mBase, uint32(v2469+v2535<<(uint(v2526)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2507+v2539<<(uint(v2526)%32)))) = v2546
	v2549 = v2516 | v2522
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2549<<(uint(v2522)%32))))
	v2560 = *(*int64)(unsafe.Add(mBase, uint32(v2469+v2549<<(uint(v2526)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2507+v2553<<(uint(v2526)%32)))) = v2560
	v2563 = v2516 | v2526
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2563<<(uint(v2522)%32))))
	v2574 = *(*int64)(unsafe.Add(mBase, uint32(v2469+v2563<<(uint(v2526)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2507+v2567<<(uint(v2526)%32)))) = v2574
	v2576 = int32(4)
	v2579 = v2513 + v2576
	if v2579 != int32(4) {
		v2513 = v2579
		v2516 = v2516 + v2576
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
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v2405)+4))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	F_UnregisterSnapshot(m, v2405)
	mBase = m.M
	v2659 = m.ExcPending
	if v2659 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2661 = m.ExcPending
	if v2661 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2668 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	F_WaitForOlderSnapshots(m, v2655, int32(1))
	mBase = m.M
	v2711 = m.ExcPending
	if v2711 != 0 {
		goto L1
	} else {
		goto L436
	}
L433:
	;
	goto L432
L434:
	;
	v2672 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2672&int32(1) == int32(0) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2677 = int32(_a_F_ReindexRelationConcurrently_14)
	v2679 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2680 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2679 + v2680
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2668)))
	*(*int32)(unsafe.Add(mBase, uint32(v2668))) = v2683 + v2680
	v2687 = int32(0)
	v2689 = int32(_a_F_ReindexRelationConcurrently_15)
	v2690 = base.AtomicRmwOr32(m, v2687, v2689, v2687)
	*(*int64)(unsafe.Add(mBase, uint32(v2668+int32(72))+232)) = int64(7)
	v2698 = base.AtomicRmwOr32(m, v2687, v2689, v2687)
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v2668)))
	*(*int32)(unsafe.Add(mBase, uint32(v2668))) = v2699 + v2680
	v2705 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2705 - v2680
	goto L433
L436:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v2715 = v2327 + int32(1)
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v2715 < v2716 {
		v2327 = v2715
		goto L391
	} else {
		goto L438
	}
L438:
	;
	goto L392
L439:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2724 == int32(0) {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	F_WaitForLockersMultiple(m, v1739, int32(5), int32(1))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L1
	} else {
		goto L444
	}
L441:
	;
	goto L440
L442:
	;
	v2728 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2728&int32(1) == int32(0) {
		goto L441
	} else {
		goto L443
	}
L443:
	;
	v2733 = int32(_a_F_ReindexRelationConcurrently_14)
	v2735 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2736 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2735 + v2736
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2724)))
	*(*int32)(unsafe.Add(mBase, uint32(v2724))) = v2739 + v2736
	v2743 = int32(0)
	v2745 = int32(_a_F_ReindexRelationConcurrently_15)
	v2746 = base.AtomicRmwOr32(m, v2743, v2745, v2743)
	*(*int64)(unsafe.Add(mBase, uint32(v2724+int32(72))+232)) = int64(3)
	v2754 = base.AtomicRmwOr32(m, v2743, v2745, v2743)
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v2724)))
	*(*int32)(unsafe.Add(mBase, uint32(v2724))) = v2755 + v2736
	v2761 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2761 - v2736
	goto L441
L444:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	goto L337
L446:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v2813 = F_LWLockAcquire(m, v2809+int32(512), int32(0))
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v2817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2816)+124)))
	v2819 = v2817 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2816)+124)) = uint8(v2819)
	v2822 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v2822)+12))
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2816)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2823+v2824))) = uint8(v2819)
	v2828 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v2828+int32(512))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v2835 = int32(0)
	goto L451
L449:
	;
	F_MemoryContextDelete(m, v52)
	mBase = m.M
	v4330 = m.ExcPending
	if v4330 != 0 {
		goto L1
	} else {
		goto L718
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v79
	F_errmsg(m, v4249, v38+int32(32))
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L1
	} else {
		goto L714
	}
L451:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v2835 < v2870 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v4232 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v4233 = m.ExcPending
	if v4233 != 0 {
		goto L1
	} else {
		goto L712
	}
L453:
	;
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v2876 = v2872 + v2835<<(uint(int32(2))%32)
	goto L455
L454:
	;
	v2876 = int32(0)
	goto L455
L455:
	;
	if v1611 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	goto L452
L457:
	;
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v2884+v2835<<(uint(int32(2))%32))))
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v2876)))
	v3434 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v3434 != 0 {
		goto L536
	} else {
		goto L537
	}
L458:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L1
	} else {
		goto L462
	}
L459:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if base.B2i32(v2876 == int32(0))|base.B2i32(v2881 <= v2835) != 0 {
		goto L458
	} else {
		goto L460
	}
L460:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	if v2884 != 0 {
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
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v2894 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2894 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	F_WaitForLockersMultiple(m, v1739, int32(8), int32(1))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L1
	} else {
		goto L468
	}
L465:
	;
	goto L464
L466:
	;
	v2898 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2898&int32(1) == int32(0) {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v2903 = int32(_a_F_ReindexRelationConcurrently_14)
	v2905 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2906 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2905 + v2906
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2894)))
	*(*int32)(unsafe.Add(mBase, uint32(v2894))) = v2909 + v2906
	v2913 = int32(0)
	v2915 = int32(_a_F_ReindexRelationConcurrently_15)
	v2916 = base.AtomicRmwOr32(m, v2913, v2915, v2913)
	*(*int64)(unsafe.Add(mBase, uint32(v2894+int32(72))+232)) = int64(8)
	v2924 = base.AtomicRmwOr32(m, v2913, v2915, v2913)
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2894)))
	*(*int32)(unsafe.Add(mBase, uint32(v2894))) = v2925 + v2906
	v2931 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2931 - v2906
	goto L465
L468:
	;
	v2939 = int32(0)
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v2939 < v2940 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2944 = v2939
	goto L472
L470:
	;
	goto L471
L471:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3054 = m.ExcPending
	if v3054 != 0 {
		goto L1
	} else {
		goto L489
	}
L472:
	;
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2978+v2944<<(uint(int32(2))%32))))
	v2984 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2984 != 0 {
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
	v2986 = m.ExcPending
	if v2986 != 0 {
		goto L1
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	v2987 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L1
	} else {
		goto L478
	}
L477:
	;
	goto L476
L478:
	;
	F_PushActiveSnapshot(m, v2987)
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	v2991 = *(*int32)(unsafe.Add(mBase, uint32(v2982)))
	v2992 = *(*int32)(unsafe.Add(mBase, uint32(v2982)+4))
	v2994 = F_table_open(m, v2992, int32(4))
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v2997 = F_index_open(m, v2991, int32(4))
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	F_TransferPredicateLocksToHeapRelation(m, v2997)
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_index_set_state_flags(m, v2991, int32(3))
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	F_CacheInvalidateRelcache(m, v2994)
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	F_relation_close(m, v2994, int32(0))
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	F_relation_close(m, v2997, int32(0))
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	v3015 = v2944 + int32(1)
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v3015 < v3016 {
		v2944 = v3015
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
	v3056 = m.ExcPending
	if v3056 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v3061 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v3061 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	F_WaitForLockersMultiple(m, v1739, int32(8), int32(1))
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L1
	} else {
		goto L495
	}
L492:
	;
	goto L491
L493:
	;
	v3065 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v3065&int32(1) == int32(0) {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v3070 = int32(_a_F_ReindexRelationConcurrently_14)
	v3072 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v3073 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v3072 + v3073
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v3061)))
	*(*int32)(unsafe.Add(mBase, uint32(v3061))) = v3076 + v3073
	v3080 = int32(0)
	v3082 = int32(_a_F_ReindexRelationConcurrently_15)
	v3083 = base.AtomicRmwOr32(m, v3080, v3082, v3080)
	*(*int64)(unsafe.Add(mBase, uint32(v3061+int32(72))+232)) = int64(9)
	v3091 = base.AtomicRmwOr32(m, v3080, v3082, v3080)
	v3092 = *(*int32)(unsafe.Add(mBase, uint32(v3061)))
	*(*int32)(unsafe.Add(mBase, uint32(v3061))) = v3092 + v3073
	v3098 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v3098 - v3073
	goto L492
L495:
	;
	v3106 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	F_PushActiveSnapshot(m, v3106)
	mBase = m.M
	v3109 = m.ExcPending
	if v3109 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v3110 = F_new_object_addresses(m)
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if int32(0) < v3112 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v3117 = int32(0)
	goto L502
L500:
	;
	goto L501
L501:
	;
	v3205 = int32(0)
	F_performMultipleDeletions(m, v3110, v3205, int32(33))
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L1
	} else {
		goto L506
	}
L502:
	;
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v3151+v3117<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v3155)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v3158
	F_add_exact_object_address(m, v38+int32(168), v3110)
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L1
	} else {
		goto L504
	}
L503:
	;
	goto L501
L504:
	;
	v3167 = v3117 + int32(1)
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v3167 < v3168 {
		v3117 = v3167
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
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	if v1733 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L1
	} else {
		goto L516
	}
L510:
	;
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+4))
	if v3216 <= int32(0) {
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v3220 = v3205
	goto L512
L512:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+12))
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v3254+v3220<<(uint(int32(2))%32))))
	F_UnlockRelationIdForSession(m, v3258, int32(4))
	mBase = m.M
	v3261 = m.ExcPending
	if v3261 != 0 {
		goto L1
	} else {
		goto L514
	}
L513:
	;
	goto L509
L514:
	;
	v3263 = v3220 + int32(1)
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+4))
	if v3263 < v3264 {
		v3220 = v3263
		goto L512
	} else {
		goto L515
	}
L515:
	;
	goto L513
L516:
	;
	v3303 = int32(1)
	v3304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v3304&v3303 == int32(0) {
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
	if v1611 == int32(0) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v3422 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L1
	} else {
		goto L534
	}
L520:
	;
	v3313 = int32(0)
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v3314 <= v3313 {
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v3318 = v3313
	goto L522
L522:
	;
	v3352 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v3352+v3318<<(uint(int32(2))%32))))
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v3356)))
	v3360 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3361 = m.ExcPending
	if v3361 != 0 {
		goto L1
	} else {
		goto L524
	}
L523:
	;
	goto L519
L524:
	;
	if v3360 != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v3362 = F_get_rel_namespace(m, v3357)
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L1
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v3382 = v3318 + int32(1)
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v3382 < v3383 {
		v3318 = v3382
		goto L522
	} else {
		goto L533
	}
L528:
	;
	v3364 = F_get_namespace_name(m, v3362)
	mBase = m.M
	v3365 = m.ExcPending
	if v3365 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v3366 = F_get_rel_name(m, v3357)
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v3366
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v3364
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_23), v38+int32(48))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(_a_F_ReindexRelationConcurrently_24), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
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
	if v3422 == int32(0) {
		goto L449
	} else {
		goto L535
	}
L535:
	;
	v4249 = int32(_a_F_ReindexRelationConcurrently_25)
	v4273 = int32(_a_F_ReindexRelationConcurrently_26)
	goto L450
L536:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		goto L1
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v3432)))
	v3438 = F_get_rel_name(m, v3437)
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L540
	}
L539:
	;
	goto L538
L540:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v3432)+4))
	v3443 = F_get_rel_namespace(m, v3442)
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v3446 = F_ChooseRelationName(m, v3438, int32(0), int32(_a_F_ReindexRelationConcurrently_27), v3443, int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v3448 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3449 = m.ExcPending
	if v3449 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	F_PushActiveSnapshot(m, v3448)
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3431)))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3432)))
	v3454 = m.G0
	v3456 = v3454 - int32(240)
	m.G0 = v3456
	v3459 = F_relation_open(m, v3453, int32(4))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L1
	} else {
		goto L545
	}
L545:
	;
	v3462 = F_relation_open(m, v3452, int32(4))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	v3466 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	v3470 = F_SearchSysCacheCopy(m, int32(57), v3453, int32(0))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L1
	} else {
		goto L553
	}
L548:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4222 = m.ExcPending
	if v4222 != 0 {
		goto L1
	} else {
		goto L709
	}
L549:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4209 = m.ExcPending
	if v4209 != 0 {
		goto L1
	} else {
		goto L706
	}
L550:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L1
	} else {
		goto L703
	}
L551:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4179 = m.ExcPending
	if v4179 != 0 {
		goto L1
	} else {
		goto L700
	}
L552:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L697
	}
L553:
	;
	if v3470 != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3474 = F_SearchSysCacheCopy(m, int32(57), v3452, int32(0))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
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
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L1
	} else {
		goto L694
	}
L557:
	;
	if v3474 == int32(0) {
		goto L552
	} else {
		goto L558
	}
L558:
	;
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v3474)+16))
	v3479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3478)+22)))
	v3480 = v3478 + v3479
	v3481 = int32(4)
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v3470)+16))
	v3484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483)+22)))
	v3485 = v3483 + v3484
	v3487 = v3485 + v3481
	v3489 = F_strncpy(m, v3480+v3481, v3487, int32(64))
	mBase = m.M
	v3490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3489)+63)) = uint8(v3490)
	goto L559
L559:
	;
	v3493 = F_strncpy(m, v3487, v3446, int32(64))
	mBase = m.M
	v3494 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3493)+63)) = uint8(v3494)
	goto L560
L560:
	;
	v3496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3480)+131)))
	v3497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3485)+131)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3480)+131)) = uint8(v3497)
	*(*uint8)(unsafe.Add(mBase, uint32(v3485)+131)) = uint8(v3496)
	F_CatalogTupleUpdate(m, v3466, v3470+int32(4), v3470)
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	F_CatalogTupleUpdate(m, v3466, v3474+int32(4), v3474)
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	F_pfree(m, v3470)
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	F_pfree(m, v3474)
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	v3514 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	v3518 = F_SearchSysCacheCopy(m, int32(34), v3453, int32(0))
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	if v3518 == int32(0) {
		goto L551
	} else {
		goto L567
	}
L567:
	;
	v3524 = F_SearchSysCacheCopy(m, int32(34), v3452, int32(0))
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	if v3524 == int32(0) {
		goto L550
	} else {
		goto L569
	}
L569:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3524)+16))
	v3529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3528)+22)))
	v3530 = v3528 + v3529
	v3531 = *(*int32)(unsafe.Add(mBase, uint32(v3518)+16))
	v3532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3531)+22)))
	v3533 = v3531 + v3532
	v3534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3533)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+14)) = uint8(v3534)
	v3536 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3533)+14)) = uint8(v3536)
	v3538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3533)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+15)) = uint8(v3538)
	*(*uint8)(unsafe.Add(mBase, uint32(v3533)+15)) = uint8(v3536)
	v3542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3533)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+16)) = uint8(v3542)
	v3544 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3533)+16)) = uint8(v3544)
	v3546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3533)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+22)) = uint8(v3546)
	v3548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3533)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+18)) = uint8(v3544)
	*(*uint8)(unsafe.Add(mBase, uint32(v3530)+17)) = uint8(v3548)
	*(*uint8)(unsafe.Add(mBase, uint32(v3533)+22)) = uint8(v3536)
	*(*uint16)(unsafe.Add(mBase, uint32(v3533)+17)) = uint16(v3536)
	F_CatalogTupleUpdate(m, v3514, v3518+int32(4), v3518)
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	F_CatalogTupleUpdate(m, v3514, v3524+int32(4), v3524)
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_pfree(m, v3518)
	mBase = m.M
	v3565 = m.ExcPending
	if v3565 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	F_pfree(m, v3524)
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	v3568 = int32(0)
	v3569 = m.G0
	v3571 = v3569 - int32(144)
	m.G0 = v3571
	v3575 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	F_ScanKeyInit(m, v3571, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_ScanKeyInit(m, v3571+int32(48), int32(5), int32(3), int32(184), v3453)
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	F_ScanKeyInit(m, v3571+int32(96), int32(6), int32(3), int32(65), int32(0))
	mBase = m.M
	v3597 = m.ExcPending
	if v3597 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v3602 = F_systable_beginscan(m, v3575, int32(2674), int32(1), int32(0), int32(3), v3571)
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v3604 = F_systable_getnext(m, v3602)
	mBase = m.M
	v3605 = m.ExcPending
	if v3605 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	if v3604 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v3606 = v3604
	v3610 = v3568
	goto L583
L581:
	;
	v3661 = v3568
	goto L582
L582:
	;
	F_systable_endscan(m, v3602)
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L592
	}
L583:
	;
	v3641 = *(*int32)(unsafe.Add(mBase, uint32(v3606)+16))
	v3642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3641)+22)))
	v3643 = v3641 + v3642
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v3643)))
	if v3644 != int32(2606) {
		v3654 = v3610
		goto L585
	} else {
		goto L586
	}
L584:
	;
	v3661 = v3654
	goto L582
L585:
	;
	v3655 = F_systable_getnext(m, v3602)
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L1
	} else {
		goto L590
	}
L586:
	;
	v3647 = *(*int32)(unsafe.Add(mBase, uint32(v3643)+8))
	if v3647 != 0 {
		v3654 = v3610
		goto L585
	} else {
		goto L587
	}
L587:
	;
	v3648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3643)+24)))
	if v3648 != int32(110) {
		v3654 = v3610
		goto L585
	} else {
		goto L588
	}
L588:
	;
	v3651 = *(*int32)(unsafe.Add(mBase, uint32(v3643)+4))
	v3652 = F_lappend_oid(m, v3610, v3651)
	mBase = m.M
	v3653 = m.ExcPending
	if v3653 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v3654 = v3652
	goto L585
L590:
	;
	if v3655 != 0 {
		v3606 = v3655
		v3610 = v3654
		goto L583
	} else {
		goto L591
	}
L591:
	;
	goto L584
L592:
	;
	F_relation_close(m, v3575, int32(1))
	mBase = m.M
	v3696 = m.ExcPending
	if v3696 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	m.G0 = v3571 + int32(144)
	v3700 = F_get_index_constraint(m, v3453)
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	if v3700 != 0 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v3702 = F_lappend_oid(m, v3661, v3700)
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L1
	} else {
		goto L598
	}
L596:
	;
	v3704 = v3661
	goto L597
L597:
	;
	v3707 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L1
	} else {
		goto L599
	}
L598:
	;
	v3704 = v3702
	goto L597
L599:
	;
	v3711 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	if v3704 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v3885 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3456)+88)) = v3885
	*(*int64)(unsafe.Add(mBase, uint32(v3456)+80)) = v3885
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+80)) = v3452
	v3892 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+72)) = v3892
	v3895 = v3456 + int32(96)
	F_ScanKeyInit(m, v3895, v3892, int32(3), int32(184), v3453)
	mBase = m.M
	v3900 = m.ExcPending
	if v3900 != 0 {
		goto L1
	} else {
		goto L627
	}
L602:
	;
	v3715 = *(*int32)(unsafe.Add(mBase, uint32(v3704)+4))
	if v3715 <= int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3719 = int32(0)
	goto L604
L604:
	;
	v3755 = *(*int32)(unsafe.Add(mBase, uint32(v3704)+12))
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3755+v3719<<(uint(int32(2))%32))))
	v3761 = F_SearchSysCacheCopy(m, int32(19), v3759, int32(0))
	mBase = m.M
	v3762 = m.ExcPending
	if v3762 != 0 {
		goto L1
	} else {
		goto L606
	}
L605:
	;
	goto L601
L606:
	;
	if v3761 == int32(0) {
		goto L549
	} else {
		goto L607
	}
L607:
	;
	v3765 = *(*int32)(unsafe.Add(mBase, uint32(v3761)+16))
	v3766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3765)+22)))
	v3767 = v3765 + v3766
	v3768 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+88))
	if v3453 == v3768 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3767)+88)) = v3452
	F_CatalogTupleUpdate(m, v3707, v3761+int32(4), v3761)
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L611
	}
L609:
	;
	goto L610
L610:
	;
	F_pfree(m, v3761)
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L1
	} else {
		goto L612
	}
L611:
	;
	goto L610
L612:
	;
	v3778 = v3456 + int32(96)
	F_ScanKeyInit(m, v3778, int32(11), int32(3), int32(184), v3759)
	mBase = m.M
	v3783 = m.ExcPending
	if v3783 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v3785 = int32(1)
	v3788 = F_systable_beginscan(m, v3711, int32(2699), v3785, int32(0), v3785, v3778)
	mBase = m.M
	v3789 = m.ExcPending
	if v3789 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	goto L615
L615:
	;
	v3825 = F_systable_getnext(m, v3788)
	mBase = m.M
	v3826 = m.ExcPending
	if v3826 != 0 {
		goto L1
	} else {
		goto L617
	}
L616:
	;
	F_systable_endscan(m, v3788)
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L1
	} else {
		goto L625
	}
L617:
	;
	if v3825 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v3825)+16))
	v3828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3827)+22)))
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(v3827+v3828)+88))
	if v3830 != v3453 {
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
	v3832 = F_heap_copytuple(m, v3825)
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(v3832)+16))
	v3835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3834)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3834+v3835)+88)) = v3452
	F_CatalogTupleUpdate(m, v3711, v3832+int32(4), v3832)
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	F_pfree(m, v3832)
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	goto L615
L625:
	;
	v3847 = v3719 + int32(1)
	v3848 = *(*int32)(unsafe.Add(mBase, uint32(v3704)+4))
	if v3847 < v3848 {
		v3719 = v3847
		goto L604
	} else {
		goto L626
	}
L626:
	;
	goto L605
L627:
	;
	F_ScanKeyInit(m, v3456+int32(144), int32(2), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3908 = m.ExcPending
	if v3908 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v3911 = int32(3)
	F_ScanKeyInit(m, v3456+int32(192), v3911, v3911, int32(65), int32(0))
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	v3919 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	v3925 = F_systable_beginscan(m, v3919, int32(2675), int32(1), int32(0), int32(3), v3895)
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	v3927 = F_systable_getnext(m, v3925)
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	if v3927 != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3929 = *(*int32)(unsafe.Add(mBase, uint32(v3919)+52))
	v3936 = F_heap_modify_tuple(m, v3927, v3929, v3456+int32(80), v3456+int32(76), v3456+int32(72))
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L1
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	F_systable_endscan(m, v3925)
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	F_CatalogTupleUpdate(m, v3919, v3936+int32(4), v3936)
	mBase = m.M
	v3941 = m.ExcPending
	if v3941 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	goto L635
L638:
	;
	F_relation_close(m, v3919, int32(0))
	mBase = m.M
	v3947 = m.ExcPending
	if v3947 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	v3948 = F_get_rel_relispartition(m, v3453)
	mBase = m.M
	v3949 = m.ExcPending
	if v3949 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	if v3948 != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v3950 = F_get_partition_ancestors(m, v3453)
	mBase = m.M
	v3951 = m.ExcPending
	if v3951 != 0 {
		goto L1
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	F_changeDependenciesOf(m, v3452, v3453)
	mBase = m.M
	v3966 = m.ExcPending
	if v3966 != 0 {
		goto L1
	} else {
		goto L648
	}
L644:
	;
	v3952 = *(*int32)(unsafe.Add(mBase, uint32(v3950)+12))
	v3953 = *(*int32)(unsafe.Add(mBase, uint32(v3952)))
	v3954 = int32(0)
	v3956 = F_DeleteInheritsTuple(m, v3453, v3953, v3954, v3954)
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	F_StoreSingleInheritance(m, v3452, v3953, int32(1))
	mBase = m.M
	v3960 = m.ExcPending
	if v3960 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	F_list_free(m, v3950)
	mBase = m.M
	v3962 = m.ExcPending
	if v3962 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	goto L643
L648:
	;
	F_changeDependenciesOn(m, v3452, v3453)
	mBase = m.M
	v3968 = m.ExcPending
	if v3968 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	F_changeDependenciesOf(m, v3453, v3452)
	mBase = m.M
	v3970 = m.ExcPending
	if v3970 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	F_changeDependenciesOn(m, v3453, v3452)
	mBase = m.M
	v3972 = m.ExcPending
	if v3972 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v3976 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3977 = *(*int32)(unsafe.Add(mBase, uint32(v3459)+48))
	v3978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3977)+117)))
	if v3978 != 0 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3979 = int32(0)
	goto L654
L653:
	;
	v3979 = v3976
	goto L654
L654:
	;
	v3980 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3459)+56)))
	v3981 = F_pgstat_fetch_entry(m, int32(2), v3979, v3980)
	mBase = m.M
	v3982 = m.ExcPending
	if v3982 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	if v3981 != 0 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3986 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3987 = *(*int32)(unsafe.Add(mBase, uint32(v3462)+48))
	v3988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3987)+117)))
	if v3988 != 0 {
		goto L659
	} else {
		goto L660
	}
L657:
	;
	goto L658
L658:
	;
	v4003 = m.G0
	v4005 = v4003 - int32(48)
	m.G0 = v4005
	v4009 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v4010 = m.ExcPending
	if v4010 != 0 {
		goto L1
	} else {
		goto L664
	}
L659:
	;
	v3989 = int32(0)
	goto L661
L660:
	;
	v3989 = v3986
	goto L661
L661:
	;
	v3990 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3462)+56)))
	v3992 = F_pgstat_get_entry_ref_locked(m, int32(2), v3989, v3990, int32(0))
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(v3992)+4))
	base.MemoryCopy(m, v3994+int32(24), v3981, int32(216))
	F_pgstat_unlock_entry(m, v3992)
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	goto L658
L664:
	;
	F_ScanKeyInit(m, v4005, int32(1), int32(3), int32(184), v3453)
	mBase = m.M
	v4015 = m.ExcPending
	if v4015 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v4017 = int32(1)
	v4020 = F_systable_beginscan(m, v4009, int32(2696), v4017, int32(0), v4017, v4005)
	mBase = m.M
	v4021 = m.ExcPending
	if v4021 != 0 {
		goto L1
	} else {
		goto L667
	}
L666:
	;
	F_relation_close(m, v4009, int32(3))
	mBase = m.M
	v4123 = m.ExcPending
	if v4123 != 0 {
		goto L1
	} else {
		goto L687
	}
L667:
	;
	v4022 = F_systable_getnext(m, v4020)
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	if v4022 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L669:
	;
	F_systable_endscan(m, v4020)
	mBase = m.M
	v4027 = m.ExcPending
	if v4027 != 0 {
		goto L1
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	v4032 = int32(0)
	v4041 = v4022
	goto L673
L672:
	;
	goto L666
L673:
	;
	v4063 = F_heap_copytuple(m, v4041)
	mBase = m.M
	v4064 = m.ExcPending
	if v4064 != 0 {
		goto L1
	} else {
		goto L675
	}
L674:
	;
	F_systable_endscan(m, v4020)
	mBase = m.M
	v4081 = m.ExcPending
	if v4081 != 0 {
		goto L1
	} else {
		goto L684
	}
L675:
	;
	v4065 = *(*int32)(unsafe.Add(mBase, uint32(v4063)+16))
	v4066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4065)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v4065+v4066))) = v3452
	if v4032 == int32(0) {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v4071 = F_CatalogOpenIndexes(m, v4009)
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		goto L1
	} else {
		goto L679
	}
L677:
	;
	v4073 = v4032
	goto L678
L678:
	;
	F_CatalogTupleInsertWithInfo(m, v4009, v4063, v4073)
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L1
	} else {
		goto L680
	}
L679:
	;
	v4073 = v4071
	goto L678
L680:
	;
	F_pfree(m, v4063)
	mBase = m.M
	v4077 = m.ExcPending
	if v4077 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	v4078 = F_systable_getnext(m, v4020)
	mBase = m.M
	v4079 = m.ExcPending
	if v4079 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	if v4078 != 0 {
		v4032 = v4073
		v4041 = v4078
		goto L673
	} else {
		goto L683
	}
L683:
	;
	goto L674
L684:
	;
	if v4073 == int32(0) {
		goto L666
	} else {
		goto L685
	}
L685:
	;
	F_CatalogCloseIndexes(m, v4073)
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	goto L666
L687:
	;
	m.G0 = v4005 + int32(48)
	F_relation_close(m, v3466, int32(3))
	mBase = m.M
	v4129 = m.ExcPending
	if v4129 != 0 {
		goto L1
	} else {
		goto L688
	}
L688:
	;
	F_relation_close(m, v3514, int32(3))
	mBase = m.M
	v4132 = m.ExcPending
	if v4132 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	F_relation_close(m, v3707, int32(3))
	mBase = m.M
	v4135 = m.ExcPending
	if v4135 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	F_relation_close(m, v3711, int32(3))
	mBase = m.M
	v4138 = m.ExcPending
	if v4138 != 0 {
		goto L1
	} else {
		goto L691
	}
L691:
	;
	F_relation_close(m, v3459, int32(0))
	mBase = m.M
	v4141 = m.ExcPending
	if v4141 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	F_relation_close(m, v3462, int32(0))
	mBase = m.M
	v4144 = m.ExcPending
	if v4144 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	m.G0 = v3456 + int32(240)
	goto L548
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3456))) = v3453
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_28), v3456)
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		goto L1
	} else {
		goto L695
	}
L695:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1585), int32(_a_F_ReindexRelationConcurrently_29))
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+16)) = v3452
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_28), v3456+int32(16))
	mBase = m.M
	v4170 = m.ExcPending
	if v4170 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1589), int32(_a_F_ReindexRelationConcurrently_29))
	mBase = m.M
	v4175 = m.ExcPending
	if v4175 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+32)) = v3453
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_28), v3456+int32(32))
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1615), int32(_a_F_ReindexRelationConcurrently_29))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+48)) = v3452
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_28), v3456+int32(48))
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1619), int32(_a_F_ReindexRelationConcurrently_29))
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3456)+64)) = v3759
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_30), v3456-int32(-64))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_18), int32(1683), int32(_a_F_ReindexRelationConcurrently_29))
	mBase = m.M
	v4220 = m.ExcPending
	if v4220 != 0 {
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
	v4223 = *(*int32)(unsafe.Add(mBase, uint32(v3432)+4))
	F_CacheInvalidateRelcacheByRelid(m, v4223)
	mBase = m.M
	v4225 = m.ExcPending
	if v4225 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4229 = m.ExcPending
	if v4229 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	v2835 = v2835 + int32(1)
	goto L451
L712:
	;
	if v4232 == int32(0) {
		goto L449
	} else {
		goto L713
	}
L713:
	;
	v4249 = int32(_a_F_ReindexRelationConcurrently_23)
	v4273 = int32(_a_F_ReindexRelationConcurrently_31)
	goto L450
L714:
	;
	v4282 = F_pg_rusage_show(m, v38+int32(248))
	mBase = m.M
	v4283 = m.ExcPending
	if v4283 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v4282
	F_errdetail(m, int32(_a_F_ReindexRelationConcurrently_32), v38+int32(16))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), v4273, int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4293 = m.ExcPending
	if v4293 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	goto L449
L718:
	;
	v4333 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v4333 == int32(0) {
		goto L720
	} else {
		goto L721
	}
L719:
	;
	v4380 = v3303
	goto L12
L720:
	;
	goto L719
L721:
	;
	v4337 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v4337&int32(1) == int32(0) {
		goto L720
	} else {
		goto L722
	}
L722:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, uint32(v4333)+220))
	if v4342 == int32(0) {
		goto L720
	} else {
		goto L723
	}
L723:
	;
	v4345 = int32(_a_F_ReindexRelationConcurrently_14)
	v4347 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v4348 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v4347 + v4348
	v4351 = *(*int32)(unsafe.Add(mBase, uint32(v4333)))
	*(*int32)(unsafe.Add(mBase, uint32(v4333))) = v4351 + v4348
	v4355 = int32(0)
	v4357 = int32(_a_F_ReindexRelationConcurrently_15)
	v4358 = base.AtomicRmwOr32(m, v4355, v4357, v4355)
	*(*int32)(unsafe.Add(mBase, uint32(v4333)+220)) = v4355
	*(*int32)(unsafe.Add(mBase, uint32(v4333)+224)) = v4355
	v4366 = base.AtomicRmwOr32(m, v4355, v4357, v4355)
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v4333)))
	*(*int32)(unsafe.Add(mBase, uint32(v4333))) = v4367 + v4348
	v4373 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v4373 - v4348
	goto L720
L724:
	;
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_33), int32(0))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3939), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
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
	v4435 = m.ExcPending
	if v4435 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v4437 = F_get_tablespace_name(m, v4436)
	mBase = m.M
	v4438 = m.ExcPending
	if v4438 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v4437
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_34), v38)
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3864), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v4447 = m.ExcPending
	if v4447 != 0 {
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
