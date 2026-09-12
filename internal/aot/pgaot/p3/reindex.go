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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
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
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v598 int32
	_ = v598
	var v626 int32
	_ = v626
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v872 int64
	_ = v872
	var v874 int64
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
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
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v981 int64
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
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
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
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
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
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
	var v1139 int32
	_ = v1139
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
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1179 int32
	_ = v1179
	var v1189 int32
	_ = v1189
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1246 int32
	_ = v1246
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1275 int32
	_ = v1275
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1325 int32
	_ = v1325
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1364 int32
	_ = v1364
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int64
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int64
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1598 int64
	_ = v1598
	var v1600 int64
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	var v1668 int32
	_ = v1668
	var v1675 int32
	_ = v1675
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int64
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1774 int32
	_ = v1774
	var v1781 int32
	_ = v1781
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1938 int32
	_ = v1938
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2037 int32
	_ = v2037
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2079 int64
	_ = v2079
	var v2081 int64
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2095 int32
	_ = v2095
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2121 int32
	_ = v2121
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2146 int64
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2153 int32
	_ = v2153
	var v2160 int64
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2174 int64
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2188 int64
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2335 int32
	_ = v2335
	var v2341 int32
	_ = v2341
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2358 int32
	_ = v2358
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2454 int32
	_ = v2454
	var v2461 int32
	_ = v2461
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2499 int64
	_ = v2499
	var v2503 int64
	_ = v2503
	var v2505 int64
	_ = v2505
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2519 int32
	_ = v2519
	var v2525 int32
	_ = v2525
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2536 int32
	_ = v2536
	var v2545 int32
	_ = v2545
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2570 int64
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2584 int64
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2598 int64
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2605 int32
	_ = v2605
	var v2612 int64
	_ = v2612
	var v2614 int32
	_ = v2614
	var v2617 int32
	_ = v2617
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2702 int32
	_ = v2702
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2723 int32
	_ = v2723
	var v2729 int32
	_ = v2729
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2769 int32
	_ = v2769
	var v2775 int32
	_ = v2775
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2846 int32
	_ = v2846
	var v2849 int32
	_ = v2849
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2890 int32
	_ = v2890
	var v2893 int32
	_ = v2893
	var v2897 int32
	_ = v2897
	var v2900 int32
	_ = v2900
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2910 int32
	_ = v2910
	var v2914 int32
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2931 int32
	_ = v2931
	var v2937 int32
	_ = v2937
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2984 int32
	_ = v2984
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3006 int32
	_ = v3006
	var v3009 int32
	_ = v3009
	var v3011 int32
	_ = v3011
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3060 int32
	_ = v3060
	var v3062 int32
	_ = v3062
	var v3067 int32
	_ = v3067
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3088 int32
	_ = v3088
	var v3094 int32
	_ = v3094
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3113 int32
	_ = v3113
	var v3147 int32
	_ = v3147
	var v3151 int32
	_ = v3151
	var v3154 int32
	_ = v3154
	var v3161 int32
	_ = v3161
	var v3163 int32
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3201 int32
	_ = v3201
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3216 int32
	_ = v3216
	var v3250 int32
	_ = v3250
	var v3254 int32
	_ = v3254
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3314 int32
	_ = v3314
	var v3348 int32
	_ = v3348
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3370 int32
	_ = v3370
	var v3375 int32
	_ = v3375
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3427 int32
	_ = v3427
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3480 int32
	_ = v3480
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3496 int32
	_ = v3496
	var v3500 int32
	_ = v3500
	var v3502 int32
	_ = v3502
	var v3504 int32
	_ = v3504
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3529 int32
	_ = v3529
	var v3531 int32
	_ = v3531
	var v3535 int32
	_ = v3535
	var v3537 int32
	_ = v3537
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3552 int32
	_ = v3552
	var v3556 int32
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3564 int32
	_ = v3564
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3575 int32
	_ = v3575
	var v3582 int32
	_ = v3582
	var v3590 int32
	_ = v3590
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3605 int32
	_ = v3605
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3656 int32
	_ = v3656
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3693 int32
	_ = v3693
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3748 int32
	_ = v3748
	var v3752 int32
	_ = v3752
	var v3754 int32
	_ = v3754
	var v3755 int32
	_ = v3755
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3767 int32
	_ = v3767
	var v3769 int32
	_ = v3769
	var v3776 int32
	_ = v3776
	var v3778 int32
	_ = v3778
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3825 int32
	_ = v3825
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3836 int32
	_ = v3836
	var v3838 int32
	_ = v3838
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3880 int64
	_ = v3880
	var v3887 int32
	_ = v3887
	var v3895 int32
	_ = v3895
	var v3903 int32
	_ = v3903
	var v3906 int32
	_ = v3906
	var v3911 int32
	_ = v3911
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3938 int32
	_ = v3938
	var v3941 int32
	_ = v3941
	var v3944 int32
	_ = v3944
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
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
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3963 int32
	_ = v3963
	var v3965 int32
	_ = v3965
	var v3967 int32
	_ = v3967
	var v3969 int32
	_ = v3969
	var v3973 int32
	_ = v3973
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3977 int64
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3979 int32
	_ = v3979
	var v3983 int32
	_ = v3983
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3987 int64
	_ = v3987
	var v3989 int32
	_ = v3989
	var v3990 int32
	_ = v3990
	var v3991 int32
	_ = v3991
	var v3995 int32
	_ = v3995
	var v3998 int32
	_ = v3998
	var v4001 int32
	_ = v4001
	var v4003 int32
	_ = v4003
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4013 int32
	_ = v4013
	var v4015 int32
	_ = v4015
	var v4018 int32
	_ = v4018
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4021 int32
	_ = v4021
	var v4025 int32
	_ = v4025
	var v4032 int32
	_ = v4032
	var v4036 int32
	_ = v4036
	var v4061 int32
	_ = v4061
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4064 int32
	_ = v4064
	var v4069 int32
	_ = v4069
	var v4070 int32
	_ = v4070
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4076 int32
	_ = v4076
	var v4077 int32
	_ = v4077
	var v4079 int32
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4121 int32
	_ = v4121
	var v4127 int32
	_ = v4127
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4136 int32
	_ = v4136
	var v4139 int32
	_ = v4139
	var v4142 int32
	_ = v4142
	var v4149 int32
	_ = v4149
	var v4153 int32
	_ = v4153
	var v4158 int32
	_ = v4158
	var v4162 int32
	_ = v4162
	var v4168 int32
	_ = v4168
	var v4173 int32
	_ = v4173
	var v4177 int32
	_ = v4177
	var v4183 int32
	_ = v4183
	var v4188 int32
	_ = v4188
	var v4192 int32
	_ = v4192
	var v4198 int32
	_ = v4198
	var v4203 int32
	_ = v4203
	var v4207 int32
	_ = v4207
	var v4213 int32
	_ = v4213
	var v4218 int32
	_ = v4218
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4223 int32
	_ = v4223
	var v4227 int32
	_ = v4227
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4245 int32
	_ = v4245
	var v4271 int32
	_ = v4271
	var v4277 int32
	_ = v4277
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4287 int32
	_ = v4287
	var v4291 int32
	_ = v4291
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4335 int32
	_ = v4335
	var v4338 int32
	_ = v4338
	var v4341 int32
	_ = v4341
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4347 int32
	_ = v4347
	var v4351 int32
	_ = v4351
	var v4361 int32
	_ = v4361
	var v4370 int32
	_ = v4370
	v4 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(416)
	m.G0 = v38
	v41 = *(*int64)(unsafe.Add(mBase, _consts[472]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+232)) = v41
	v44 = *(*int64)(unsafe.Add(mBase, _consts[473]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+224)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v52 = F_AllocSetContextCreateInternal(m, v47, int32(91972), v4, int32(1024), int32(8192))
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
	v59 = int32(4470560)
	v60 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
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
		goto L21
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
	F___gettimeofday(m, v38+int32(248))
	mBase = m.M
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v60
	v78 = v63
	v79 = v67
	goto L5
L10:
	;
	m.G0 = v38 + int32(416)
	return v4370
L11:
	;
	if v674 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L12:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+112))
	if v469 != 0 {
		goto L135
	} else {
		goto L136
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L131
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L126
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L122
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L118
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L114
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L110
	}
L19:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v264 = F_IndexGetRelation(m, l1, int32(base.Ui32(v259&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L77
	}
L20:
	;
	v86 = int32(4470560)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v91 = F_lappend_oid(m, int32(0), l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	switch v80&int32(255) - int32(105) {
	case 0:
		goto L19
	default:
		goto L13
	case 4, 9, 11:
		goto L20
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v87
	goto L23
L23:
	;
	if base.Ui32(l1) < base.Ui32(int32(12000)) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v97&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v107 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v101 = F_try_table_open(m, l1, int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v104 = F_table_open(m, l1, int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	if v101 != 0 {
		v106 = v101
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v4370 = v4
	goto L10
L31:
	;
	v106 = v104
	goto L25
L32:
	;
	v109 = int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+56))
	if base.Ui32(v110) < base.Ui32(int32(12000)) {
		v119 = v109
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v120 = F_RelationGetIndexList(m, v106)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	if v119 != 0 {
		goto L17
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+68))
	if v114 == int32(99) {
		v119 = v109
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v117 = F_isTempToastNamespace(m, v114)
	mBase = m.M
	v119 = v117
	goto L36
L39:
	;
	goto L34
L40:
	;
	v130 = v122
	v138 = int32(0)
	goto L46
L41:
	;
	if v120 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v122 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v122 < v123 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v442 = int32(0)
	goto L12
L45:
	;
	goto L44
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164+v130<<(uint(int32(2))%32))))
	v170 = F_index_open(m, v168, int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	v442 = v251
	goto L12
L48:
	;
	F_relation_close(m, v170, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L75
	}
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+192))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+18)))
	if v173 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v178 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+15)))
	if v207 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	if v178 == int32(0) {
		v251 = v138
		goto L48
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v185 = F_get_rel_namespace(m, v168)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v187 = F_get_namespace_name(m, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v189 = F_get_rel_name(m, v168)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+132)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v187
	F_errmsg(m, int32(665404), v38+int32(128))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(630921), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(485797), int32(3685), int32(18922))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v251 = v138
	goto L48
L62:
	;
	v212 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v237 = int32(4470560)
	v238 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v242 = F_palloc(m, int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L73
	}
L65:
	;
	if v212 == int32(0) {
		v251 = v138
		goto L48
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v219 = F_get_rel_namespace(m, v168)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v221 = F_get_namespace_name(m, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v223 = F_get_rel_name(m, v168)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+116)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = v221
	F_errmsg(m, int32(325696), v38+int32(112))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(485797), int32(3691), int32(18922))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v251 = v138
	goto L48
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v168
	v245 = F_lappend(m, v138, v242)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v238
	v251 = v245
	goto L48
L75:
	;
	v256 = v130 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v256 < v257 {
		v130 = v256
		v138 = v251
		goto L46
	} else {
		goto L76
	}
L76:
	;
	goto L47
L77:
	;
	if v264 == int32(0) {
		v4370 = v4
		goto L10
	} else {
		goto L78
	}
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(v264) < base.Ui32(int32(12000)) {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	v270 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v270 != int32(99) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v277 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v276 = F_isTempToastNamespace(m, v270)
	mBase = m.M
	v277 = v276
	goto L85
L84:
	;
	v277 = int32(1)
	goto L85
L85:
	;
	goto L82
L86:
	;
	v278 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v282&int32(4) != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v278 == int32(0) {
		goto L15
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v292 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	v286 = F_try_table_open(m, v264, int32(4))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v289 = F_table_open(m, v264, int32(4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	if v286 != 0 {
		v291 = v286
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v4370 = v4
	goto L10
L97:
	;
	v291 = v289
	goto L91
L98:
	;
	v294 = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)+56))
	if base.Ui32(v295) < base.Ui32(int32(12000)) {
		v304 = v294
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	F_sequence_close(m, v291, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L106
	}
L101:
	;
	if v304 != 0 {
		goto L14
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v291)+48))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+68))
	if v299 == int32(99) {
		v304 = v294
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v302 = F_isTempToastNamespace(m, v299)
	mBase = m.M
	v304 = v302
	goto L102
L105:
	;
	goto L100
L106:
	;
	v308 = int32(4470560)
	v309 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+188)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+156)) = v264
	v317 = F_list_make1_impl(m, int32(472), v38+int32(156))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v320 = F_palloc(m, int32(16))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = l1
	v324 = F_lappend(m, int32(0), v320)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v309
	v674 = v324
	v682 = v317
	goto L11
L110:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(18349), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(485797), int32(3650), int32(18922))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+144)) = v351 + int32(4)
	F_errmsg(m, int32(681384), v38+int32(144))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(485797), int32(3670), int32(18922))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(18349), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(485797), int32(3780), int32(18922))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(388654), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(485797), int32(3791), int32(18922))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v404 = F_get_rel_name(m, l1)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+160)) = v404
	F_errmsg(m, int32(681384), v38+int32(160))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(485797), int32(3816), int32(18922))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(18393), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(485797), int32(3845), int32(18922))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v472 = F_table_open(m, v469, int32(4))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	v636 = v442
	v644 = v91
	goto L137
L137:
	;
	F_sequence_close(m, v106, int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L165
	}
L138:
	;
	v474 = int32(4470560)
	v475 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v478 = F_lappend_oid(m, v91, v469)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v475
	v482 = F_RelationGetIndexList(m, v472)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	F_sequence_close(m, v472, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L164
	}
L141:
	;
	if v482 == int32(0) {
		v598 = v442
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	if v486 <= int32(0) {
		v598 = v442
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v490 = int32(0)
	v498 = v442
	goto L144
L144:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524+v490<<(uint(int32(2))%32))))
	v530 = F_index_open(m, v528, int32(4))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	v598 = v581
	goto L140
L146:
	;
	F_relation_close(m, v530, int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L162
	}
L147:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v530)+192))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+18)))
	if v533 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v538 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v567 = int32(4470560)
	v568 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v572 = F_palloc(m, int32(16))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L160
	}
L151:
	;
	if v538 == int32(0) {
		v581 = v498
		goto L146
	} else {
		goto L152
	}
L152:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v545 = F_get_rel_namespace(m, v528)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v547 = F_get_namespace_name(m, v545)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v549 = F_get_rel_name(m, v528)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v547
	F_errmsg(m, int32(665404), v38+int32(96))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errhint(m, int32(630921), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(485797), int32(3738), int32(18922))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v581 = v498
	goto L146
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v528
	v575 = F_lappend(m, v498, v572)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v568
	v581 = v575
	goto L146
L162:
	;
	v586 = v490 + int32(1)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	if v586 < v587 {
		v490 = v586
		v498 = v581
		goto L144
	} else {
		goto L163
	}
L163:
	;
	goto L145
L164:
	;
	v636 = v598
	v644 = v478
	goto L137
L165:
	;
	v674 = v636
	v682 = v644
	goto L11
L166:
	;
	v4370 = int32(0)
	goto L10
L167:
	;
	goto L168
L168:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v703 != int32(1664) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v1774 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L170:
	;
	v706 = int32(0)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v706 < v708 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L1
	} else {
		goto L334
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L1
	} else {
		goto L331
	}
L174:
	;
	v720 = int32(0)
	v723 = v706
	v727 = v706
	goto L177
L175:
	;
	v1623 = v706
	v1627 = v706
	goto L176
L176:
	;
	if v682 == int32(0) {
		v1774 = v1627
		v1781 = v4
		goto L169
	} else {
		goto L320
	}
L177:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v747+v720<<(uint(int32(2))%32))))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v754 = F_index_open(m, v752, int32(4))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	v1623 = v1551
	v1627 = v1565
	goto L176
L179:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v754)+192))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	v759 = F_table_open(m, v757, int32(4))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(184)))) = v766
	v769 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(180)))) = v769
	goto L181
L181:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v759)+48))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+80))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v773 | int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v772
	goto L182
L182:
	;
	v781 = int32(4468600)
	v783 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v785 = v783 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v785
	goto L183
L183:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v789 = F_RelationGetIndexExpressions(m, v754)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	if v789 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v794 = int32(1)
	goto L188
L187:
	;
	v792 = F_RelationGetIndexPredicate(m, v754)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v751)+12)) = uint8(base.B2i32(v794 == int32(0)))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v759)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v751)+4)) = v798
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v754)+48))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v800)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v751)+8)) = v801
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v754)+48))
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+118)))
	if v804 == int32(116) {
		goto L173
	} else {
		goto L190
	}
L189:
	;
	v794 = v792
	goto L188
L190:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v810 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = int64(4)
	v872 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v751))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v872
	v874 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v751)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v874
	v878 = v38 + int32(224)
	v880 = v38 + int32(192)
	v881 = int32(0)
	v888 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v888 == v881 {
		goto L204
	} else {
		goto L205
	}
L192:
	;
	goto L191
L193:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v814 != int32(1) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v817 = int32(4465220)
	v819 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v820 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v819 + v820
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v823 + v820
	*(*int32)(unsafe.Add(mBase, uint32(v810)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v810)+224)) = v798
	v830 = v810 + int32(232)
	if v830&int32(3) == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	v857 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v856 + v857
	v860 = int32(4465220)
	v862 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v862 - v857
	goto L192
L196:
	;
	v836 = v810 + int32(392)
	if base.Ui32(v836) <= base.Ui32(v830) {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v853 = F___memset(m, v830, int32(0), int32(160))
	mBase = m.M
	goto L195
L199:
	;
	v840 = v810 + int32(236)
	if base.Ui32(v840) < base.Ui32(v836) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v842 = v836
	goto L202
L201:
	;
	v842 = v840
	goto L202
L202:
	;
	v850 = F___memset(m, v830, int32(0), (v842-v810-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L195
L203:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v1055 = F_get_rel_name(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L220
	}
L204:
	;
	goto L203
L205:
	;
	goto L206
L206:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v894&int32(1) == int32(0) {
		goto L204
	} else {
		goto L207
	}
L207:
	;
	v899 = int32(4465220)
	v901 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v902 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v901 + v902
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	*(*int32)(unsafe.Add(mBase, uint32(v888))) = v905 + v902
	goto L209
L208:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	v1036 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v888))) = v1035 + v1036
	v1039 = int32(4465220)
	v1041 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1041 - v1036
	goto L204
L209:
	;
	v914 = v888 + int32(232)
	goto L210
L210:
	;
	v920 = int32(0)
	v923 = v881
	goto L213
L212:
	;
	goto L208
L213:
	;
	v929 = int32(2)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v878+v923<<(uint(v929)%32))))
	v933 = int32(3)
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v880+v923<<(uint(v933)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v914+v932<<(uint(v933)%32)))) = v939
	v942 = v923 | int32(1)
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v878+v942<<(uint(v929)%32))))
	v953 = *(*int64)(unsafe.Add(mBase, uint32(v880+v942<<(uint(v933)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v914+v946<<(uint(v933)%32)))) = v953
	v956 = v923 | v929
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v878+v956<<(uint(v929)%32))))
	v967 = *(*int64)(unsafe.Add(mBase, uint32(v880+v956<<(uint(v933)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v914+v960<<(uint(v933)%32)))) = v967
	v970 = v923 | v933
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v878+v970<<(uint(v929)%32))))
	v981 = *(*int64)(unsafe.Add(mBase, uint32(v880+v970<<(uint(v933)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v914+v974<<(uint(v933)%32)))) = v981
	v983 = int32(4)
	v986 = v920 + v983
	if v986 != int32(4) {
		v920 = v986
		v923 = v923 + v983
		goto L213
	} else {
		goto L215
	}
L214:
	;
	goto L212
L215:
	;
	goto L214
L220:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v754)+192))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	v1061 = F_get_rel_namespace(m, v1060)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1064 = F_ChooseRelationName(m, v1055, int32(0), int32(31484), v1061, int32(0))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1066 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v1075 = int32(0)
	v1078 = m.G0
	v1080 = v1078 - int32(48)
	m.G0 = v1080
	v1083 = F_index_open(m, v1074, int32(3))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L232
	}
L224:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v759)+48))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+119)))
	if v1068 != int32(116) {
		v1073 = v1066
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v754)+48))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+92))
	v1073 = v1072
	goto L223
L227:
	;
	goto L226
L228:
	;
	v1535 = F_index_open(m, v1462, int32(4))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L303
	}
L229:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L1
	} else {
		goto L300
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L297
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L1
	} else {
		goto L294
	}
L232:
	;
	v1085 = F_BuildIndexInfo(m, v1083)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+92))
	if v1087 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1091 = F_SearchSysCache1(m, int32(34), v1074)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L290
	}
L237:
	;
	if v1091 == int32(0) {
		goto L231
	} else {
		goto L238
	}
L238:
	;
	v1097 = F_SysCacheGetAttrNotNull(m, int32(34), v1091, int32(18))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1101 = F_SysCacheGetAttrNotNull(m, int32(34), v1091, int32(19))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v1104 = F_SearchSysCache1(m, int32(57), v1074)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	if v1104 == int32(0) {
		goto L230
	} else {
		goto L242
	}
L242:
	;
	v1112 = F_SysCacheGetAttr(m, int32(57), v1104, int32(33), v1080+int32(47))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+76))
	if v1114 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1117 = F_SysCacheGetAttrNotNull(m, int32(34), v1091, int32(20))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	v1126 = v1075
	goto L246
L246:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+84))
	if v1127 != 0 {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v1119 = F_text_to_cstring(m, v1117)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v1121 = F_stringToNode(m, v1119)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_pfree(m, v1119)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v1126 = v1121
	goto L246
L251:
	;
	v1130 = F_SysCacheGetAttrNotNull(m, int32(34), v1091, int32(21))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	v1141 = v1075
	goto L253
L253:
	;
	v1142 = int32(0)
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+8))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+132))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085)+116)))
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085)+117)))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+204))
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+28)))
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085)+124)))
	v1153 = F_makeIndexInfo(m, v1143, v1144, v1145, v1126, v1141, v1146, v1147, v1142, int32(1), v1151, v1152)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L259
	}
L254:
	;
	v1132 = F_text_to_cstring(m, v1130)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1134 = F_stringToNode(m, v1132)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1136 = F_make_ands_implicit(m, v1134)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_pfree(m, v1132)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1141 = v1136
	goto L253
L259:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	if int32(0) < v1155 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1158 = int32(12)
	v1179 = int32(0)
	v1189 = v1075
	goto L263
L261:
	;
	v1246 = v1075
	goto L262
L262:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	v1258 = F_palloc0(m, v1255<<(uint(int32(2))%32))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L267
	}
L263:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+52))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)))
	v1208 = F_lappend(m, v1189, v1198+v1199<<(uint(int32(4))%32)+v1179*int32(100)+int32(24))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L1
	} else {
		goto L265
	}
L264:
	;
	v1246 = v1208
	goto L262
L265:
	;
	v1210 = int32(1)
	v1211 = v1179 << (uint(v1210) % 32)
	v1214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211+(v1085+v1158)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1153+v1158+v1211))) = uint16(v1214)
	v1217 = v1179 + v1210
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	if v1217 < v1218 {
		v1179 = v1217
		v1189 = v1208
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	if int32(0) < v1260 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1275 = v1142
	goto L271
L269:
	;
	v1325 = v1260
	goto L270
L270:
	;
	v1346 = F_palloc0(m, v1325<<(uint(int32(3))%32))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L275
	}
L271:
	;
	v1302 = v1275 + int32(1)
	v1304 = F_get_attoptions(m, v1074, base.I32_extend16_s(v1302))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L273
	}
L272:
	;
	v1325 = v1307
	goto L270
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258+v1275<<(uint(int32(2))%32)))) = v1304
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	if v1302 < v1307 {
		v1275 = v1302
		goto L271
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v1348 = int32(0)
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	if v1348 < v1349 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1364 = v1348
	goto L279
L277:
	;
	goto L278
L278:
	;
	v1446 = int32(0)
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+48))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+84))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+248))
	v1453 = int32(24)
	v1462 = F_index_create(m, v759, v1064, v1446, v1446, v1446, v1446, v1153, v1246, v1451, v1073, v1452, v1097+v1453, v1258, v1101+v1453, v1346, v1112, int32(12), v1446, int32(1), v1446, v1446)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L286
	}
L279:
	;
	v1389 = v1364 + int32(1)
	v1391 = F_SearchSysCache2(m, int32(7), v1074, base.I32_extend16_s(v1389))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L281
	}
L280:
	;
	goto L278
L281:
	;
	if v1391 == int32(0) {
		goto L229
	} else {
		goto L282
	}
L282:
	;
	v1399 = F_SysCacheGetAttr(m, int32(7), v1391, int32(21), v1080+int32(47))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_ReleaseCatCache(m, v1391)
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v1405 = v1346 + v1364<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1405))) = v1399
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080)+47)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1405)+4)) = uint8(v1407)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+4))
	if v1389 < v1409 {
		v1364 = v1389
		goto L279
	} else {
		goto L285
	}
L285:
	;
	goto L280
L286:
	;
	F_relation_close(m, v1083, int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_ReleaseCatCache(m, v1091)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_ReleaseCatCache(m, v1104)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	m.G0 = v1080 + int32(48)
	goto L228
L290:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	F_errmsg(m, int32(434823), int32(0))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(484069), int32(1333), int32(18084))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1074
	F_errmsg_internal(m, int32(39514), v1080)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(484069), int32(1338), int32(18084))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080)+16)) = v1074
	F_errmsg_internal(m, int32(45662), v1080+int32(16))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(484069), int32(1350), int32(18084))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080)+36)) = v1074
	*(*int32)(unsafe.Add(mBase, uint32(v1080)+32)) = v1389
	F_errmsg_internal(m, int32(45850), v1080+int32(32))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(484069), int32(1431), int32(18084))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	v1537 = int32(4470560)
	v1538 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v1542 = F_palloc(m, int32(16))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1542))) = v1462
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1542)+12)) = uint8(v1545)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v751)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1542)+4)) = v1547
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v751)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1542)+8)) = v1549
	v1551 = F_lappend(m, v723, v1542)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1554 = F_palloc(m, int32(8))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v1556 = *(*int64)(unsafe.Add(mBase, uint32(v754)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1554))) = v1556
	v1558 = F_lappend(m, v727, v1554)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1561 = F_palloc(m, int32(8))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v1563 = *(*int64)(unsafe.Add(mBase, uint32(v1535)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1561))) = v1563
	v1565 = F_lappend(m, v1558, v1561)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1538
	F_relation_close(m, v754, int32(0))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_relation_close(m, v1535, int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_AtEOXact_GUC(m, int32(0), v785)
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v38)+184))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v1579
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1578
	goto L313
L313:
	;
	F_sequence_close(m, v759, int32(0))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	if l0 != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1587 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = v1587
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v1587
	v1592 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v1592
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v1462
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v1598 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v1598
	v1600 = *(*int64)(unsafe.Add(mBase, uint32(v38)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+80)) = v1600
	F_EventTriggerCollectSimpleCommand(m, v38+int32(80), v38-int32(-64), l0)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1609 = v720 + int32(1)
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v1609 < v1610 {
		v720 = v1609
		v723 = v1551
		v727 = v1565
		goto L177
	} else {
		goto L319
	}
L318:
	;
	goto L317
L319:
	;
	goto L178
L320:
	;
	v1649 = int32(0)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	if v1650 <= v1649 {
		v1774 = v1627
		v1781 = v4
		goto L169
	} else {
		goto L321
	}
L321:
	;
	v1656 = v1649
	v1668 = v1627
	v1675 = v4
	goto L322
L322:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v682)+12))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1688+v1656<<(uint(int32(2))%32))))
	v1694 = F_table_open(m, v1692, int32(4))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L324
	}
L323:
	;
	v1774 = v1705
	v1781 = v1716
	goto L169
L324:
	;
	v1696 = int32(4470560)
	v1697 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v1701 = F_palloc(m, int32(8))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1703 = *(*int64)(unsafe.Add(mBase, uint32(v1694)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1701))) = v1703
	v1705 = F_lappend(m, v1668, v1701)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	v1708 = F_palloc(m, int32(16))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1708))) = v1710
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1701)))
	*(*int64)(unsafe.Add(mBase, uint32(v1708)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v1708)+4)) = v1712
	v1716 = F_lappend(m, v1675, v1708)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1697
	F_sequence_close(m, v1694, int32(0))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	v1724 = v1656 + int32(1)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	if v1724 < v1725 {
		v1656 = v1724
		v1668 = v1705
		v1675 = v1716
		goto L322
	} else {
		goto L330
	}
L330:
	;
	goto L323
L331:
	;
	F_errmsg_internal(m, int32(18443), int32(0))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(485797), int32(3939), int32(18922))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1748 = F_get_tablespace_name(m, v1747)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1748
	F_errmsg(m, int32(696418), v38)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(485797), int32(3864), int32(18922))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L346
	}
L340:
	;
	v1796 = int32(0)
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+4))
	if v1797 <= v1796 {
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v1801 = v1796
	goto L342
L342:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+12))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1835+v1801<<(uint(int32(2))%32))))
	F_LockRelationIdForSession(m, v1839, int32(4))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L1
	} else {
		goto L344
	}
L343:
	;
	goto L339
L344:
	;
	v1844 = v1801 + int32(1)
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+4))
	if v1844 < v1845 {
		v1801 = v1844
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v1892 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	F_WaitForLockersMultiple(m, v1781, int32(5), int32(1))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L353
	}
L350:
	;
	goto L349
L351:
	;
	v1896 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v1896 != int32(1) {
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v1899 = int32(4465220)
	v1901 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v1902 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1901 + v1902
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1892)))
	*(*int32)(unsafe.Add(mBase, uint32(v1892))) = v1905 + v1902
	*(*int64)(unsafe.Add(mBase, uint32(v1892+int32(72))+232)) = int64(1)
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1892)))
	*(*int32)(unsafe.Add(mBase, uint32(v1892))) = v1913 + v1902
	v1919 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1919 - v1902
	goto L350
L353:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	if v1623 != 0 {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L1
	} else {
		goto L480
	}
L356:
	;
	v1929 = int32(0)
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v1929 < v1930 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	goto L358
L358:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L1
	} else {
		goto L473
	}
L359:
	;
	v1938 = v1929
	goto L362
L360:
	;
	goto L361
L361:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L409
	}
L362:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+12))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1968+v1938<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L364
	}
L363:
	;
	goto L361
L364:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v1976 != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972)+12)))
	if v1979 == int32(1) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	goto L367
L369:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v1987 = F_LWLockAcquire(m, v1983+int32(512), int32(0))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	v2009 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1990)+124)))
	v1993 = v1991 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1990)+124)) = uint8(v1993)
	v1996 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+12))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v1997+v1998))) = uint8(v1993)
	v2002 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v2002+int32(512))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	goto L371
L374:
	;
	F_PushActiveSnapshot(m, v2009)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+4))
	v2017 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2017 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = int64(4)
	v2079 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1972))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2079
	v2081 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1972)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2081
	v2085 = v38 + int32(224)
	v2087 = v38 + int32(192)
	v2088 = int32(0)
	v2095 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2095 == v2088 {
		goto L389
	} else {
		goto L390
	}
L377:
	;
	goto L376
L378:
	;
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2021 != int32(1) {
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v2024 = int32(4465220)
	v2026 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2027 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2026 + v2027
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v2017)))
	*(*int32)(unsafe.Add(mBase, uint32(v2017))) = v2030 + v2027
	*(*int32)(unsafe.Add(mBase, uint32(v2017)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2017)+224)) = v2014
	v2037 = v2017 + int32(232)
	if v2037&int32(3) == int32(0) {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2017)))
	v2064 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2017))) = v2063 + v2064
	v2067 = int32(4465220)
	v2069 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2069 - v2064
	goto L377
L381:
	;
	v2043 = v2017 + int32(392)
	if base.Ui32(v2043) <= base.Ui32(v2037) {
		goto L380
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	v2060 = F___memset(m, v2037, int32(0), int32(160))
	mBase = m.M
	goto L380
L384:
	;
	v2047 = v2017 + int32(236)
	if base.Ui32(v2047) < base.Ui32(v2043) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2049 = v2043
	goto L387
L386:
	;
	v2049 = v2047
	goto L387
L387:
	;
	v2057 = F___memset(m, v2037, int32(0), (v2049-v2017-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L380
L388:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v1972)+4))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v1972)))
	F_index_concurrently_build(m, v2261, v2262)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L405
	}
L389:
	;
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2101&int32(1) == int32(0) {
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v2106 = int32(4465220)
	v2108 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2109 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2108 + v2109
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	*(*int32)(unsafe.Add(mBase, uint32(v2095))) = v2112 + v2109
	goto L394
L393:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	v2243 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2095))) = v2242 + v2243
	v2246 = int32(4465220)
	v2248 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2248 - v2243
	goto L389
L394:
	;
	v2121 = v2095 + int32(232)
	goto L395
L395:
	;
	v2127 = int32(0)
	v2130 = v2088
	goto L398
L397:
	;
	goto L393
L398:
	;
	v2136 = int32(2)
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2085+v2130<<(uint(v2136)%32))))
	v2140 = int32(3)
	v2146 = *(*int64)(unsafe.Add(mBase, uint32(v2087+v2130<<(uint(v2140)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2121+v2139<<(uint(v2140)%32)))) = v2146
	v2149 = v2130 | int32(1)
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2085+v2149<<(uint(v2136)%32))))
	v2160 = *(*int64)(unsafe.Add(mBase, uint32(v2087+v2149<<(uint(v2140)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2121+v2153<<(uint(v2140)%32)))) = v2160
	v2163 = v2130 | v2136
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2085+v2163<<(uint(v2136)%32))))
	v2174 = *(*int64)(unsafe.Add(mBase, uint32(v2087+v2163<<(uint(v2140)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2121+v2167<<(uint(v2140)%32)))) = v2174
	v2177 = v2130 | v2140
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2085+v2177<<(uint(v2136)%32))))
	v2188 = *(*int64)(unsafe.Add(mBase, uint32(v2087+v2177<<(uint(v2140)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2121+v2181<<(uint(v2140)%32)))) = v2188
	v2190 = int32(4)
	v2193 = v2127 + v2190
	if v2193 != int32(4) {
		v2127 = v2193
		v2130 = v2130 + v2190
		goto L398
	} else {
		goto L400
	}
L399:
	;
	goto L397
L400:
	;
	goto L399
L405:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v2270 = v1938 + int32(1)
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v2270 < v2271 {
		v1938 = v2270
		goto L362
	} else {
		goto L408
	}
L408:
	;
	goto L363
L409:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2314 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	F_WaitForLockersMultiple(m, v1781, int32(5), int32(1))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L1
	} else {
		goto L414
	}
L411:
	;
	goto L410
L412:
	;
	v2318 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2318 != int32(1) {
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v2321 = int32(4465220)
	v2323 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2324 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2323 + v2324
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	*(*int32)(unsafe.Add(mBase, uint32(v2314))) = v2327 + v2324
	*(*int64)(unsafe.Add(mBase, uint32(v2314+int32(72))+232)) = int64(3)
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	*(*int32)(unsafe.Add(mBase, uint32(v2314))) = v2335 + v2324
	v2341 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2341 - v2324
	goto L411
L414:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v2351 = int32(0)
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v2352 <= v2351 {
		goto L355
	} else {
		goto L416
	}
L416:
	;
	v2358 = v2351
	goto L417
L417:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+12))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2390+v2358<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L1
	} else {
		goto L419
	}
L418:
	;
	goto L355
L419:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v2398 != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v2401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2394)+12)))
	if v2401 == int32(1) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	goto L422
L424:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v2409 = F_LWLockAcquire(m, v2405+int32(512), int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L1
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	v2431 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L429
	}
L427:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v2413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2412)+124)))
	v2415 = v2413 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2412)+124)) = uint8(v2415)
	v2418 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v2418)+12))
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2412)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2419+v2420))) = uint8(v2415)
	v2424 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v2424+int32(512))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	goto L426
L429:
	;
	v2433 = F_RegisterSnapshot(m, v2431)
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_PushActiveSnapshot(m, v2433)
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2394)+4))
	v2441 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2441 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	v2499 = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+200)) = v2499
	*(*int64)(unsafe.Add(mBase, uint32(v38)+192)) = v2499
	v2503 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2394))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = v2503
	v2505 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2394)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+216)) = v2505
	v2509 = v38 + int32(224)
	v2511 = v38 + int32(192)
	v2512 = int32(0)
	v2519 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2519 == v2512 {
		goto L445
	} else {
		goto L446
	}
L433:
	;
	goto L432
L434:
	;
	v2445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2445 != int32(1) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2448 = int32(4465220)
	v2450 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2451 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2450 + v2451
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2441)))
	*(*int32)(unsafe.Add(mBase, uint32(v2441))) = v2454 + v2451
	*(*int32)(unsafe.Add(mBase, uint32(v2441)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2441)+224)) = v2438
	v2461 = v2441 + int32(232)
	if v2461&int32(3) == int32(0) {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2441)))
	v2488 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2441))) = v2487 + v2488
	v2491 = int32(4465220)
	v2493 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2493 - v2488
	goto L433
L437:
	;
	v2467 = v2441 + int32(392)
	if base.Ui32(v2467) <= base.Ui32(v2461) {
		goto L436
	} else {
		goto L440
	}
L438:
	;
	goto L439
L439:
	;
	v2484 = F___memset(m, v2461, int32(0), int32(160))
	mBase = m.M
	goto L436
L440:
	;
	v2471 = v2441 + int32(236)
	if base.Ui32(v2471) < base.Ui32(v2467) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v2473 = v2467
	goto L443
L442:
	;
	v2473 = v2471
	goto L443
L443:
	;
	v2481 = F___memset(m, v2461, int32(0), (v2473-v2441-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L436
L444:
	;
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(v2394)+4))
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2394)))
	F_validate_index(m, v2685, v2686, v2433)
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L1
	} else {
		goto L461
	}
L445:
	;
	goto L444
L446:
	;
	goto L447
L447:
	;
	v2525 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2525&int32(1) == int32(0) {
		goto L445
	} else {
		goto L448
	}
L448:
	;
	v2530 = int32(4465220)
	v2532 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2533 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2532 + v2533
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2519)))
	*(*int32)(unsafe.Add(mBase, uint32(v2519))) = v2536 + v2533
	goto L450
L449:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2519)))
	v2667 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2519))) = v2666 + v2667
	v2670 = int32(4465220)
	v2672 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2672 - v2667
	goto L445
L450:
	;
	v2545 = v2519 + int32(232)
	goto L451
L451:
	;
	v2551 = int32(0)
	v2554 = v2512
	goto L454
L453:
	;
	goto L449
L454:
	;
	v2560 = int32(2)
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2509+v2554<<(uint(v2560)%32))))
	v2564 = int32(3)
	v2570 = *(*int64)(unsafe.Add(mBase, uint32(v2511+v2554<<(uint(v2564)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2545+v2563<<(uint(v2564)%32)))) = v2570
	v2573 = v2554 | int32(1)
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2509+v2573<<(uint(v2560)%32))))
	v2584 = *(*int64)(unsafe.Add(mBase, uint32(v2511+v2573<<(uint(v2564)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2545+v2577<<(uint(v2564)%32)))) = v2584
	v2587 = v2554 | v2560
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2509+v2587<<(uint(v2560)%32))))
	v2598 = *(*int64)(unsafe.Add(mBase, uint32(v2511+v2587<<(uint(v2564)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2545+v2591<<(uint(v2564)%32)))) = v2598
	v2601 = v2554 | v2564
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(v2509+v2601<<(uint(v2560)%32))))
	v2612 = *(*int64)(unsafe.Add(mBase, uint32(v2511+v2601<<(uint(v2564)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2545+v2605<<(uint(v2564)%32)))) = v2612
	v2614 = int32(4)
	v2617 = v2551 + v2614
	if v2617 != int32(4) {
		v2551 = v2617
		v2554 = v2554 + v2614
		goto L454
	} else {
		goto L456
	}
L455:
	;
	goto L453
L456:
	;
	goto L455
L461:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+4))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2691 = m.ExcPending
	if v2691 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	F_UnregisterSnapshot(m, v2433)
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2702 == int32(0) {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	F_WaitForOlderSnapshots(m, v2689, int32(1))
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L470
	}
L467:
	;
	goto L466
L468:
	;
	v2706 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2706 != int32(1) {
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v2709 = int32(4465220)
	v2711 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2712 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2711 + v2712
	v2715 = *(*int32)(unsafe.Add(mBase, uint32(v2702)))
	*(*int32)(unsafe.Add(mBase, uint32(v2702))) = v2715 + v2712
	*(*int64)(unsafe.Add(mBase, uint32(v2702+int32(72))+232)) = int64(7)
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(v2702)))
	*(*int32)(unsafe.Add(mBase, uint32(v2702))) = v2723 + v2712
	v2729 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2729 - v2712
	goto L467
L470:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v2739 = v2358 + int32(1)
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v2739 < v2740 {
		v2358 = v2739
		goto L417
	} else {
		goto L472
	}
L472:
	;
	goto L418
L473:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2748 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	F_WaitForLockersMultiple(m, v1781, int32(5), int32(1))
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L1
	} else {
		goto L478
	}
L475:
	;
	goto L474
L476:
	;
	v2752 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2752 != int32(1) {
		goto L475
	} else {
		goto L477
	}
L477:
	;
	v2755 = int32(4465220)
	v2757 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2758 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2757 + v2758
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2748)))
	*(*int32)(unsafe.Add(mBase, uint32(v2748))) = v2761 + v2758
	*(*int64)(unsafe.Add(mBase, uint32(v2748+int32(72))+232)) = int64(3)
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2748)))
	*(*int32)(unsafe.Add(mBase, uint32(v2748))) = v2769 + v2758
	v2775 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2775 - v2758
	goto L475
L478:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	goto L355
L480:
	;
	v2823 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v2827 = F_LWLockAcquire(m, v2823+int32(512), int32(0))
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	v2830 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v2831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2830)+124)))
	v2833 = v2831 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2830)+124)) = uint8(v2833)
	v2836 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2836)+12))
	v2838 = *(*int32)(unsafe.Add(mBase, uint32(v2830)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v2837+v2838))) = uint8(v2833)
	v2842 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v2842+int32(512))
	mBase = m.M
	v2846 = m.ExcPending
	if v2846 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	v2849 = int32(0)
	goto L485
L483:
	;
	F_MemoryContextDelete(m, v52)
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L1
	} else {
		goto L757
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v79
	F_errmsg(m, v4245, v38+int32(32))
	mBase = m.M
	v4277 = m.ExcPending
	if v4277 != 0 {
		goto L1
	} else {
		goto L753
	}
L485:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v2849 < v2884 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	v4230 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		goto L1
	} else {
		goto L751
	}
L487:
	;
	v2886 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	v2890 = v2886 + v2849<<(uint(int32(2))%32)
	goto L489
L488:
	;
	v2890 = int32(0)
	goto L489
L489:
	;
	if v1623 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L490:
	;
	goto L486
L491:
	;
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v2900)))
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v2890)))
	v3427 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v3427 != 0 {
		goto L571
	} else {
		goto L572
	}
L492:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L1
	} else {
		goto L497
	}
L493:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v2893 <= v2849 {
		goto L492
	} else {
		goto L494
	}
L494:
	;
	if v2890 == int32(0) {
		goto L492
	} else {
		goto L495
	}
L495:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+12))
	v2900 = v2897 + v2849<<(uint(int32(2))%32)
	if v2900 != 0 {
		goto L491
	} else {
		goto L496
	}
L496:
	;
	goto L492
L497:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v2910 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v2910 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L499:
	;
	F_WaitForLockersMultiple(m, v1781, int32(8), int32(1))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L503
	}
L500:
	;
	goto L499
L501:
	;
	v2914 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v2914 != int32(1) {
		goto L500
	} else {
		goto L502
	}
L502:
	;
	v2917 = int32(4465220)
	v2919 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v2920 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2919 + v2920
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v2910)))
	*(*int32)(unsafe.Add(mBase, uint32(v2910))) = v2923 + v2920
	*(*int64)(unsafe.Add(mBase, uint32(v2910+int32(72))+232)) = int64(8)
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2910)))
	*(*int32)(unsafe.Add(mBase, uint32(v2910))) = v2931 + v2920
	v2937 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v2937 - v2920
	goto L500
L503:
	;
	v2945 = int32(0)
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v2945 < v2946 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v2950 = v2945
	goto L507
L505:
	;
	goto L506
L506:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3060 = m.ExcPending
	if v3060 != 0 {
		goto L1
	} else {
		goto L524
	}
L507:
	;
	v2984 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2984+v2950<<(uint(int32(2))%32))))
	v2990 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v2990 != 0 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	goto L506
L509:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L1
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v2993 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2994 = m.ExcPending
	if v2994 != 0 {
		goto L1
	} else {
		goto L513
	}
L512:
	;
	goto L511
L513:
	;
	F_PushActiveSnapshot(m, v2993)
	mBase = m.M
	v2996 = m.ExcPending
	if v2996 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, uint32(v2988)))
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v2988)+4))
	v3000 = F_table_open(m, v2998, int32(4))
	mBase = m.M
	v3001 = m.ExcPending
	if v3001 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v3003 = F_index_open(m, v2997, int32(4))
	mBase = m.M
	v3004 = m.ExcPending
	if v3004 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_TransferPredicateLocksToHeapRelation(m, v3003)
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_index_set_state_flags(m, v2997, int32(3))
	mBase = m.M
	v3009 = m.ExcPending
	if v3009 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	F_CacheInvalidateRelcache(m, v3000)
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	F_sequence_close(m, v3000, int32(0))
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	F_relation_close(m, v3003, int32(0))
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3019 = m.ExcPending
	if v3019 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	v3021 = v2950 + int32(1)
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v3021 < v3022 {
		v2950 = v3021
		goto L507
	} else {
		goto L523
	}
L523:
	;
	goto L508
L524:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v3067 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v3067 == int32(0) {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	F_WaitForLockersMultiple(m, v1781, int32(8), int32(1))
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L1
	} else {
		goto L530
	}
L527:
	;
	goto L526
L528:
	;
	v3071 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v3071 != int32(1) {
		goto L527
	} else {
		goto L529
	}
L529:
	;
	v3074 = int32(4465220)
	v3076 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v3077 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v3076 + v3077
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(v3067)))
	*(*int32)(unsafe.Add(mBase, uint32(v3067))) = v3080 + v3077
	*(*int64)(unsafe.Add(mBase, uint32(v3067+int32(72))+232)) = int64(9)
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v3067)))
	*(*int32)(unsafe.Add(mBase, uint32(v3067))) = v3088 + v3077
	v3094 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v3094 - v3077
	goto L527
L530:
	;
	v3102 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F_PushActiveSnapshot(m, v3102)
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v3106 = F_new_object_addresses(m)
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if int32(0) < v3108 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v3113 = int32(0)
	goto L537
L535:
	;
	goto L536
L536:
	;
	v3201 = int32(0)
	F_performMultipleDeletions(m, v3106, v3201, int32(33))
	mBase = m.M
	v3205 = m.ExcPending
	if v3205 != 0 {
		goto L1
	} else {
		goto L541
	}
L537:
	;
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v3147+v3113<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+168)) = int32(1259)
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(v3151)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+172)) = v3154
	F_add_exact_object_address(m, v38+int32(168), v3106)
	mBase = m.M
	v3161 = m.ExcPending
	if v3161 != 0 {
		goto L1
	} else {
		goto L539
	}
L538:
	;
	goto L536
L539:
	;
	v3163 = v3113 + int32(1)
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if v3163 < v3164 {
		v3113 = v3163
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3207 = m.ExcPending
	if v3207 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	if v1774 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v3298 = m.ExcPending
	if v3298 != 0 {
		goto L1
	} else {
		goto L551
	}
L545:
	;
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+4))
	if v3212 <= int32(0) {
		goto L544
	} else {
		goto L546
	}
L546:
	;
	v3216 = v3201
	goto L547
L547:
	;
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+12))
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v3250+v3216<<(uint(int32(2))%32))))
	F_UnlockRelationIdForSession(m, v3254, int32(4))
	mBase = m.M
	v3257 = m.ExcPending
	if v3257 != 0 {
		goto L1
	} else {
		goto L549
	}
L548:
	;
	goto L544
L549:
	;
	v3259 = v3216 + int32(1)
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+4))
	if v3259 < v3260 {
		v3216 = v3259
		goto L547
	} else {
		goto L550
	}
L550:
	;
	goto L548
L551:
	;
	v3299 = int32(1)
	v3300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v3300&v3299 == int32(0) {
		goto L483
	} else {
		goto L552
	}
L552:
	;
	if v80 == int32(105) {
		goto L490
	} else {
		goto L553
	}
L553:
	;
	if v1623 == int32(0) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3418 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3419 = m.ExcPending
	if v3419 != 0 {
		goto L1
	} else {
		goto L569
	}
L555:
	;
	v3309 = int32(0)
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v3310 <= v3309 {
		goto L554
	} else {
		goto L556
	}
L556:
	;
	v3314 = v3309
	goto L557
L557:
	;
	v3348 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+12))
	v3352 = *(*int32)(unsafe.Add(mBase, uint32(v3348+v3314<<(uint(int32(2))%32))))
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(v3352)))
	v3356 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3357 = m.ExcPending
	if v3357 != 0 {
		goto L1
	} else {
		goto L559
	}
L558:
	;
	goto L554
L559:
	;
	if v3356 != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3358 = F_get_rel_namespace(m, v3353)
	mBase = m.M
	v3359 = m.ExcPending
	if v3359 != 0 {
		goto L1
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	v3378 = v3314 + int32(1)
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	if v3378 < v3379 {
		v3314 = v3378
		goto L557
	} else {
		goto L568
	}
L563:
	;
	v3360 = F_get_namespace_name(m, v3358)
	mBase = m.M
	v3361 = m.ExcPending
	if v3361 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	v3362 = F_get_rel_name(m, v3353)
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v3362
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v3360
	F_errmsg(m, int32(431366), v38+int32(48))
	mBase = m.M
	v3370 = m.ExcPending
	if v3370 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	F_errfinish(m, int32(485797), int32(4417), int32(18922))
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	goto L562
L568:
	;
	goto L558
L569:
	;
	if v3418 == int32(0) {
		goto L483
	} else {
		goto L570
	}
L570:
	;
	v4245 = int32(431394)
	v4271 = int32(4425)
	goto L484
L571:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L1
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v3425)))
	v3431 = F_get_rel_name(m, v3430)
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L1
	} else {
		goto L575
	}
L574:
	;
	goto L573
L575:
	;
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(v3425)+4))
	v3436 = F_get_rel_namespace(m, v3435)
	mBase = m.M
	v3437 = m.ExcPending
	if v3437 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	v3439 = F_ChooseRelationName(m, v3431, int32(0), int32(422667), v3436, int32(0))
	mBase = m.M
	v3440 = m.ExcPending
	if v3440 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v3441 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v3442 = m.ExcPending
	if v3442 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	F_PushActiveSnapshot(m, v3441)
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3424)))
	v3446 = *(*int32)(unsafe.Add(mBase, uint32(v3425)))
	v3447 = m.G0
	v3449 = v3447 - int32(240)
	m.G0 = v3449
	v3452 = F_relation_open(m, v3446, int32(4))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v3455 = F_relation_open(m, v3445, int32(4))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	v3459 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	v3463 = F_SearchSysCacheCopy(m, int32(57), v3446, int32(0))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L1
	} else {
		goto L588
	}
L583:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v4220 = m.ExcPending
	if v4220 != 0 {
		goto L1
	} else {
		goto L748
	}
L584:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L1
	} else {
		goto L745
	}
L585:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4192 = m.ExcPending
	if v4192 != 0 {
		goto L1
	} else {
		goto L742
	}
L586:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L739
	}
L587:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		goto L1
	} else {
		goto L736
	}
L588:
	;
	if v3463 != 0 {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v3467 = F_SearchSysCacheCopy(m, int32(57), v3445, int32(0))
	mBase = m.M
	v3468 = m.ExcPending
	if v3468 != 0 {
		goto L1
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4149 = m.ExcPending
	if v4149 != 0 {
		goto L1
	} else {
		goto L733
	}
L592:
	;
	if v3467 == int32(0) {
		goto L587
	} else {
		goto L593
	}
L593:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3467)+16))
	v3472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3471)+22)))
	v3473 = v3471 + v3472
	v3474 = int32(4)
	v3476 = *(*int32)(unsafe.Add(mBase, uint32(v3463)+16))
	v3477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3476)+22)))
	v3478 = v3476 + v3477
	v3480 = v3478 + v3474
	v3482 = F_strncpy(m, v3473+v3474, v3480, int32(64))
	mBase = m.M
	v3483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3482)+63)) = uint8(v3483)
	goto L594
L594:
	;
	v3486 = F_strncpy(m, v3480, v3439, int32(64))
	mBase = m.M
	v3487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3486)+63)) = uint8(v3487)
	goto L595
L595:
	;
	v3489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3473)+131)))
	v3490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3478)+131)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+131)) = uint8(v3490)
	*(*uint8)(unsafe.Add(mBase, uint32(v3478)+131)) = uint8(v3489)
	F_CatalogTupleUpdate(m, v3459, v3463+int32(4), v3463)
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_CatalogTupleUpdate(m, v3459, v3467+int32(4), v3467)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	F_pfree(m, v3463)
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	F_pfree(m, v3467)
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	v3507 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v3508 = m.ExcPending
	if v3508 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v3511 = F_SearchSysCacheCopy(m, int32(34), v3446, int32(0))
	mBase = m.M
	v3512 = m.ExcPending
	if v3512 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	if v3511 == int32(0) {
		goto L586
	} else {
		goto L602
	}
L602:
	;
	v3517 = F_SearchSysCacheCopy(m, int32(34), v3445, int32(0))
	mBase = m.M
	v3518 = m.ExcPending
	if v3518 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	if v3517 == int32(0) {
		goto L585
	} else {
		goto L604
	}
L604:
	;
	v3521 = *(*int32)(unsafe.Add(mBase, uint32(v3517)+16))
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3521)+22)))
	v3523 = v3521 + v3522
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(v3511)+16))
	v3525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3524)+22)))
	v3526 = v3524 + v3525
	v3527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+14)) = uint8(v3527)
	v3529 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3526)+14)) = uint8(v3529)
	v3531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+15)) = uint8(v3531)
	*(*uint8)(unsafe.Add(mBase, uint32(v3526)+15)) = uint8(v3529)
	v3535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+16)) = uint8(v3535)
	v3537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3526)+16)) = uint8(v3537)
	v3539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+22)) = uint8(v3539)
	v3541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3526)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+18)) = uint8(v3537)
	*(*uint8)(unsafe.Add(mBase, uint32(v3523)+17)) = uint8(v3541)
	*(*uint8)(unsafe.Add(mBase, uint32(v3526)+22)) = uint8(v3529)
	*(*uint16)(unsafe.Add(mBase, uint32(v3526)+17)) = uint16(v3529)
	F_CatalogTupleUpdate(m, v3507, v3511+int32(4), v3511)
	mBase = m.M
	v3552 = m.ExcPending
	if v3552 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	F_CatalogTupleUpdate(m, v3507, v3517+int32(4), v3517)
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	F_pfree(m, v3511)
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	F_pfree(m, v3517)
	mBase = m.M
	v3560 = m.ExcPending
	if v3560 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v3561 = int32(0)
	v3562 = m.G0
	v3564 = v3562 - int32(144)
	m.G0 = v3564
	v3568 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v3569 = m.ExcPending
	if v3569 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	F_ScanKeyInit(m, v3564, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3575 = m.ExcPending
	if v3575 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	F_ScanKeyInit(m, v3564+int32(48), int32(5), int32(3), int32(184), v3446)
	mBase = m.M
	v3582 = m.ExcPending
	if v3582 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	F_ScanKeyInit(m, v3564+int32(96), int32(6), int32(3), int32(65), int32(0))
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	v3595 = F_systable_beginscan(m, v3568, int32(2674), int32(1), int32(0), int32(3), v3564)
	mBase = m.M
	v3596 = m.ExcPending
	if v3596 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v3597 = F_systable_getnext(m, v3595)
	mBase = m.M
	v3598 = m.ExcPending
	if v3598 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	if v3597 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3599 = v3597
	v3605 = v3561
	goto L618
L616:
	;
	v3656 = v3561
	goto L617
L617:
	;
	F_systable_endscan(m, v3595)
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L627
	}
L618:
	;
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3599)+16))
	v3635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3634)+22)))
	v3636 = v3634 + v3635
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(v3636)))
	if v3637 != int32(2606) {
		v3647 = v3605
		goto L620
	} else {
		goto L621
	}
L619:
	;
	v3656 = v3647
	goto L617
L620:
	;
	v3648 = F_systable_getnext(m, v3595)
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L1
	} else {
		goto L625
	}
L621:
	;
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(v3636)+8))
	if v3640 != 0 {
		v3647 = v3605
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v3641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3636)+24)))
	if v3641 != int32(110) {
		v3647 = v3605
		goto L620
	} else {
		goto L623
	}
L623:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v3636)+4))
	v3645 = F_lappend_oid(m, v3605, v3644)
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	v3647 = v3645
	goto L620
L625:
	;
	if v3648 != 0 {
		v3599 = v3648
		v3605 = v3647
		goto L618
	} else {
		goto L626
	}
L626:
	;
	goto L619
L627:
	;
	F_sequence_close(m, v3568, int32(1))
	mBase = m.M
	v3689 = m.ExcPending
	if v3689 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	m.G0 = v3564 + int32(144)
	v3693 = F_get_index_constraint(m, v3446)
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	if v3693 != 0 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v3695 = F_lappend_oid(m, v3656, v3693)
	mBase = m.M
	v3696 = m.ExcPending
	if v3696 != 0 {
		goto L1
	} else {
		goto L633
	}
L631:
	;
	v3697 = v3656
	goto L632
L632:
	;
	v3700 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L1
	} else {
		goto L634
	}
L633:
	;
	v3697 = v3695
	goto L632
L634:
	;
	v3704 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	if v3697 == int32(0) {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v3880 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3449)+88)) = v3880
	*(*int64)(unsafe.Add(mBase, uint32(v3449)+80)) = v3880
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+80)) = v3445
	v3887 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+72)) = v3887
	F_ScanKeyInit(m, v3449+int32(96), v3887, int32(3), int32(184), v3446)
	mBase = m.M
	v3895 = m.ExcPending
	if v3895 != 0 {
		goto L1
	} else {
		goto L662
	}
L637:
	;
	v3708 = int32(0)
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(v3697)+4))
	if v3709 <= v3708 {
		goto L636
	} else {
		goto L638
	}
L638:
	;
	v3712 = v3708
	goto L639
L639:
	;
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3697)+12))
	v3752 = *(*int32)(unsafe.Add(mBase, uint32(v3748+v3712<<(uint(int32(2))%32))))
	v3754 = F_SearchSysCacheCopy(m, int32(19), v3752, int32(0))
	mBase = m.M
	v3755 = m.ExcPending
	if v3755 != 0 {
		goto L1
	} else {
		goto L641
	}
L640:
	;
	goto L636
L641:
	;
	if v3754 == int32(0) {
		goto L584
	} else {
		goto L642
	}
L642:
	;
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v3754)+16))
	v3759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3758)+22)))
	v3760 = v3758 + v3759
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(v3760)+88))
	if v3446 == v3761 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3760)+88)) = v3445
	F_CatalogTupleUpdate(m, v3700, v3754+int32(4), v3754)
	mBase = m.M
	v3767 = m.ExcPending
	if v3767 != 0 {
		goto L1
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	F_pfree(m, v3754)
	mBase = m.M
	v3769 = m.ExcPending
	if v3769 != 0 {
		goto L1
	} else {
		goto L647
	}
L646:
	;
	goto L645
L647:
	;
	F_ScanKeyInit(m, v3449+int32(96), int32(11), int32(3), int32(184), v3752)
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	v3778 = int32(1)
	v3783 = F_systable_beginscan(m, v3704, int32(2699), v3778, int32(0), v3778, v3449+int32(96))
	mBase = m.M
	v3784 = m.ExcPending
	if v3784 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	goto L650
L650:
	;
	v3820 = F_systable_getnext(m, v3783)
	mBase = m.M
	v3821 = m.ExcPending
	if v3821 != 0 {
		goto L1
	} else {
		goto L652
	}
L651:
	;
	F_systable_endscan(m, v3783)
	mBase = m.M
	v3840 = m.ExcPending
	if v3840 != 0 {
		goto L1
	} else {
		goto L660
	}
L652:
	;
	if v3820 != 0 {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v3820)+16))
	v3823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3822)+22)))
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v3822+v3823)+88))
	if v3825 != v3446 {
		goto L650
	} else {
		goto L656
	}
L654:
	;
	goto L655
L655:
	;
	goto L651
L656:
	;
	v3827 = F_heap_copytuple(m, v3820)
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	v3829 = *(*int32)(unsafe.Add(mBase, uint32(v3827)+16))
	v3830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3829)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3829+v3830)+88)) = v3445
	F_CatalogTupleUpdate(m, v3704, v3827+int32(4), v3827)
	mBase = m.M
	v3836 = m.ExcPending
	if v3836 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	F_pfree(m, v3827)
	mBase = m.M
	v3838 = m.ExcPending
	if v3838 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	goto L650
L660:
	;
	v3842 = v3712 + int32(1)
	v3843 = *(*int32)(unsafe.Add(mBase, uint32(v3697)+4))
	if v3842 < v3843 {
		v3712 = v3842
		goto L639
	} else {
		goto L661
	}
L661:
	;
	goto L640
L662:
	;
	F_ScanKeyInit(m, v3449+int32(144), int32(2), int32(3), int32(184), int32(1259))
	mBase = m.M
	v3903 = m.ExcPending
	if v3903 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v3906 = int32(3)
	F_ScanKeyInit(m, v3449+int32(192), v3906, v3906, int32(65), int32(0))
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	v3914 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v3915 = m.ExcPending
	if v3915 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v3922 = F_systable_beginscan(m, v3914, int32(2675), int32(1), int32(0), int32(3), v3449+int32(96))
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	v3924 = F_systable_getnext(m, v3922)
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L1
	} else {
		goto L667
	}
L667:
	;
	if v3924 != 0 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v3926 = *(*int32)(unsafe.Add(mBase, uint32(v3914)+52))
	v3933 = F_heap_modify_tuple(m, v3924, v3926, v3449+int32(80), v3449+int32(76), v3449+int32(72))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L1
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	F_systable_endscan(m, v3922)
	mBase = m.M
	v3941 = m.ExcPending
	if v3941 != 0 {
		goto L1
	} else {
		goto L673
	}
L671:
	;
	F_CatalogTupleUpdate(m, v3914, v3933+int32(4), v3933)
	mBase = m.M
	v3938 = m.ExcPending
	if v3938 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	goto L670
L673:
	;
	F_sequence_close(m, v3914, int32(0))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	v3945 = F_get_rel_relispartition(m, v3446)
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	if v3945 != 0 {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v3947 = F_get_partition_ancestors(m, v3446)
	mBase = m.M
	v3948 = m.ExcPending
	if v3948 != 0 {
		goto L1
	} else {
		goto L679
	}
L677:
	;
	goto L678
L678:
	;
	F_changeDependenciesOf(m, v3445, v3446)
	mBase = m.M
	v3963 = m.ExcPending
	if v3963 != 0 {
		goto L1
	} else {
		goto L683
	}
L679:
	;
	v3949 = *(*int32)(unsafe.Add(mBase, uint32(v3947)+12))
	v3950 = *(*int32)(unsafe.Add(mBase, uint32(v3949)))
	v3951 = int32(0)
	v3953 = F_DeleteInheritsTuple(m, v3446, v3950, v3951, v3951)
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	F_StoreSingleInheritance(m, v3445, v3950, int32(1))
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	F_list_free(m, v3947)
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	goto L678
L683:
	;
	F_changeDependenciesOn(m, v3445, v3446)
	mBase = m.M
	v3965 = m.ExcPending
	if v3965 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	F_changeDependenciesOf(m, v3446, v3445)
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	F_changeDependenciesOn(m, v3446, v3445)
	mBase = m.M
	v3969 = m.ExcPending
	if v3969 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	v3973 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v3974 = *(*int32)(unsafe.Add(mBase, uint32(v3452)+48))
	v3975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3974)+117)))
	if v3975 != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v3976 = int32(0)
	goto L689
L688:
	;
	v3976 = v3973
	goto L689
L689:
	;
	v3977 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3452)+56)))
	v3978 = F_pgstat_fetch_entry(m, int32(2), v3976, v3977)
	mBase = m.M
	v3979 = m.ExcPending
	if v3979 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	if v3978 != 0 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v3983 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3455)+48))
	v3985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3984)+117)))
	if v3985 != 0 {
		goto L694
	} else {
		goto L695
	}
L692:
	;
	goto L693
L693:
	;
	v4001 = m.G0
	v4003 = v4001 - int32(48)
	m.G0 = v4003
	v4007 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v4008 = m.ExcPending
	if v4008 != 0 {
		goto L1
	} else {
		goto L703
	}
L694:
	;
	v3986 = int32(0)
	goto L696
L695:
	;
	v3986 = v3983
	goto L696
L696:
	;
	v3987 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3455)+56)))
	v3989 = F_pgstat_get_entry_ref_locked(m, int32(2), v3986, v3987, int32(0))
	mBase = m.M
	v3990 = m.ExcPending
	if v3990 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	v3991 = *(*int32)(unsafe.Add(mBase, uint32(v3989)+4))
	goto L699
L698:
	;
	F_pgstat_unlock_entry(m, v3989)
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L1
	} else {
		goto L702
	}
L699:
	;
	v3995 = F__emscripten_memcpy_bulkmem(m, v3991+int32(24), v3978, int32(216))
	mBase = m.M
	goto L701
L701:
	;
	goto L698
L702:
	;
	goto L693
L703:
	;
	F_ScanKeyInit(m, v4003, int32(1), int32(3), int32(184), v3446)
	mBase = m.M
	v4013 = m.ExcPending
	if v4013 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	v4015 = int32(1)
	v4018 = F_systable_beginscan(m, v4007, int32(2696), v4015, int32(0), v4015, v4003)
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L1
	} else {
		goto L706
	}
L705:
	;
	F_sequence_close(m, v4007, int32(3))
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L1
	} else {
		goto L726
	}
L706:
	;
	v4020 = F_systable_getnext(m, v4018)
	mBase = m.M
	v4021 = m.ExcPending
	if v4021 != 0 {
		goto L1
	} else {
		goto L707
	}
L707:
	;
	if v4020 == int32(0) {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	F_systable_endscan(m, v4018)
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L1
	} else {
		goto L711
	}
L709:
	;
	goto L710
L710:
	;
	v4032 = int32(0)
	v4036 = v4020
	goto L712
L711:
	;
	goto L705
L712:
	;
	v4061 = F_heap_copytuple(m, v4036)
	mBase = m.M
	v4062 = m.ExcPending
	if v4062 != 0 {
		goto L1
	} else {
		goto L714
	}
L713:
	;
	F_systable_endscan(m, v4018)
	mBase = m.M
	v4079 = m.ExcPending
	if v4079 != 0 {
		goto L1
	} else {
		goto L723
	}
L714:
	;
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+16))
	v4064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4063)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v4063+v4064))) = v3445
	if v4032 == int32(0) {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v4069 = F_CatalogOpenIndexes(m, v4007)
	mBase = m.M
	v4070 = m.ExcPending
	if v4070 != 0 {
		goto L1
	} else {
		goto L718
	}
L716:
	;
	v4071 = v4032
	goto L717
L717:
	;
	F_CatalogTupleInsertWithInfo(m, v4007, v4061, v4071)
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L1
	} else {
		goto L719
	}
L718:
	;
	v4071 = v4069
	goto L717
L719:
	;
	F_pfree(m, v4061)
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	v4076 = F_systable_getnext(m, v4018)
	mBase = m.M
	v4077 = m.ExcPending
	if v4077 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	if v4076 != 0 {
		v4032 = v4071
		v4036 = v4076
		goto L712
	} else {
		goto L722
	}
L722:
	;
	goto L713
L723:
	;
	if v4071 == int32(0) {
		goto L705
	} else {
		goto L724
	}
L724:
	;
	F_CatalogCloseIndexes(m, v4071)
	mBase = m.M
	v4083 = m.ExcPending
	if v4083 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	goto L705
L726:
	;
	m.G0 = v4003 + int32(48)
	F_sequence_close(m, v3459, int32(3))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	F_sequence_close(m, v3507, int32(3))
	mBase = m.M
	v4130 = m.ExcPending
	if v4130 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	F_sequence_close(m, v3700, int32(3))
	mBase = m.M
	v4133 = m.ExcPending
	if v4133 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	F_sequence_close(m, v3704, int32(3))
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L1
	} else {
		goto L730
	}
L730:
	;
	F_relation_close(m, v3452, int32(0))
	mBase = m.M
	v4139 = m.ExcPending
	if v4139 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	F_relation_close(m, v3455, int32(0))
	mBase = m.M
	v4142 = m.ExcPending
	if v4142 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	m.G0 = v3449 + int32(240)
	goto L583
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3449))) = v3446
	F_errmsg_internal(m, int32(45551), v3449)
	mBase = m.M
	v4153 = m.ExcPending
	if v4153 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	F_errfinish(m, int32(484069), int32(1585), int32(234258))
	mBase = m.M
	v4158 = m.ExcPending
	if v4158 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L736:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+16)) = v3445
	F_errmsg_internal(m, int32(45551), v3449+int32(16))
	mBase = m.M
	v4168 = m.ExcPending
	if v4168 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	F_errfinish(m, int32(484069), int32(1589), int32(234258))
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L738
	}
L738:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+32)) = v3446
	F_errmsg_internal(m, int32(45551), v3449+int32(32))
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	F_errfinish(m, int32(484069), int32(1615), int32(234258))
	mBase = m.M
	v4188 = m.ExcPending
	if v4188 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L742:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+48)) = v3445
	F_errmsg_internal(m, int32(45551), v3449+int32(48))
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	F_errfinish(m, int32(484069), int32(1619), int32(234258))
	mBase = m.M
	v4203 = m.ExcPending
	if v4203 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+64)) = v3752
	F_errmsg_internal(m, int32(40387), v3449-int32(-64))
	mBase = m.M
	v4213 = m.ExcPending
	if v4213 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	F_errfinish(m, int32(484069), int32(1683), int32(234258))
	mBase = m.M
	v4218 = m.ExcPending
	if v4218 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L748:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v3425)+4))
	F_CacheInvalidateRelcacheByRelid(m, v4221)
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4227 = m.ExcPending
	if v4227 != 0 {
		goto L1
	} else {
		goto L750
	}
L750:
	;
	v2849 = v2849 + int32(1)
	goto L485
L751:
	;
	if v4230 == int32(0) {
		goto L483
	} else {
		goto L752
	}
L752:
	;
	v4245 = int32(431366)
	v4271 = int32(4406)
	goto L484
L753:
	;
	v4280 = F_pg_rusage_show(m, v38+int32(248))
	mBase = m.M
	v4281 = m.ExcPending
	if v4281 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v4280
	F_errdetail(m, int32(581597), v38+int32(16))
	mBase = m.M
	v4287 = m.ExcPending
	if v4287 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	F_errfinish(m, int32(485797), v4271, int32(18922))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	goto L483
L757:
	;
	v4331 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v4331 == int32(0) {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	v4370 = v3299
	goto L10
L759:
	;
	goto L758
L760:
	;
	v4335 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v4335 != int32(1) {
		goto L759
	} else {
		goto L761
	}
L761:
	;
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v4331)+220))
	if v4338 == int32(0) {
		goto L759
	} else {
		goto L762
	}
L762:
	;
	v4341 = int32(4465220)
	v4343 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v4344 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v4343 + v4344
	v4347 = *(*int32)(unsafe.Add(mBase, uint32(v4331)))
	*(*int32)(unsafe.Add(mBase, uint32(v4331))) = v4347 + v4344
	v4351 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4331)+220)) = v4351
	*(*int32)(unsafe.Add(mBase, uint32(v4331)+224)) = v4351
	*(*int32)(unsafe.Add(mBase, uint32(v4331))) = v4347 + int32(2)
	v4361 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v4361 - v4344
	goto L759
}
func F_ResetReindexState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[420]))
	if l0 <= v3 {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[46])) = v6
		*(*int32)(unsafe.Add(mBase, _consts[139])) = v6
		*(*int32)(unsafe.Add(mBase, _consts[47])) = v6
		*(*int32)(unsafe.Add(mBase, _consts[420])) = v6
	} else {
	}
	return
}
