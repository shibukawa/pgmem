package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v620 int32
	_ = v620
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v714 int32
	_ = v714
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1042 int32
	_ = v1042
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
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
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
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
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1241 int32
	_ = v1241
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1371 int32
	_ = v1371
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1560 int32
	_ = v1560
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1600 int32
	_ = v1600
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1641 int32
	_ = v1641
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
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
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
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
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1763 int32
	_ = v1763
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1851 int32
	_ = v1851
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2056 int32
	_ = v2056
	var v2063 int32
	_ = v2063
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	v6 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(256)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v40 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = int32(_a_F_ExecFindPartition_0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+131)))
	if v65 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v54 = v40
	v56 = v40 + int32(4)
	v57 = v43
	goto L1
L3:
	;
	goto L4
L4:
	;
	v44 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v49 = v44 + int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v51 != 0 {
		v54 = v51
		v56 = v49
		v57 = v50
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v52 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v54 = v52
	v56 = v49
	v57 = v50
	goto L1
L9:
	;
	v69 = F_ExecPartitionCheck(m, l1, l3, l4, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v71 == int32(0) {
		v2107 = v6
		v2111 = v37
		v2125 = v56
		v2127 = v57
		v2129 = v59
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2125))) = v2127
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v2129
	m.G0 = v2111 + int32(256)
	return v2107
L14:
	;
	v74 = l0
	v75 = l1
	v76 = l2
	v77 = l3
	v78 = l4
	v82 = v71
	v83 = v37
	v88 = l3
	v95 = v6
	v97 = v56
	v98 = v39
	v99 = v57
	v101 = v59
	goto L15
L15:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[1]))
	if v109 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v2063 == int32(0) {
		v2107 = v2047
		v2111 = v83
		v2125 = v97
		v2127 = v99
		v2129 = v101
		goto L13
	} else {
		goto L362
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v88
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if v117 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v127)+4)))
	if int32(0) < v128 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v126 = v124
	goto L21
L23:
	;
	if v115 != 0 {
		v123 = v115
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v121 = v115
	goto L25
L25:
	;
	if v121 != 0 {
		v123 = v121
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v118 = F_ExecPrepareExprList(m, v117, v78)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v118
	v121 = v118
	goto L25
L28:
	;
	v126 = int32(0)
	goto L21
L29:
	;
	v1070 = v1042 << (uint(int32(2)) % 32)
	v1071 = v82 + v1070
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+24))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+v1042))))
	if v1075 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+20)) = v541
	v1042 = v543
	goto L29
L31:
	;
	v683 = F_RelationGetPartitionKey(m, v113)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L5
	} else {
		goto L113
	}
L32:
	;
	if int32(0) <= v620 {
		v1042 = v620
		goto L29
	} else {
		goto L112
	}
L33:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	if v541 != v607 {
		goto L30
	} else {
		goto L111
	}
L34:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v275)+32))
	v620 = v606
	goto L32
L35:
	;
	if int32(0) <= v543 {
		goto L33
	} else {
		goto L110
	}
L36:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if v470 < int32(16) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	v398 = v328
	goto L93
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L90
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L5
	} else {
		goto L87
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L84
	}
L41:
	;
	v137 = v127
	v138 = v126
	v139 = int32(0)
	goto L44
L42:
	;
	v242 = v127
	v243 = v126
	v247 = v128
	goto L43
L43:
	;
	if v243 != 0 {
		goto L39
	} else {
		goto L64
	}
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166+v139<<(uint(int32(1))%32)))))
	if v170 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v242 = v234
	v243 = v217
	v247 = v235
	goto L43
L46:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83-int32(-64)+v139))) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v83+int32(96)+v139<<(uint(int32(2))%32)))) = v218
	v233 = v139 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+4)))
	if v233 < v235 {
		v137 = v234
		v138 = v217
		v139 = v233
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
	if v171 < v170 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v138 == int32(0) {
		goto L40
	} else {
		goto L54
	}
L50:
	;
	F_slot_getsomeattrs_int(m, v88, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v176 = v170 - int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v176<<(uint(int32(2))%32))))
	v216 = v176 + v177
	v217 = v138
	v218 = v183
	goto L46
L53:
	;
	goto L52
L54:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v78)+152))
	if v187 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v190 = F_MakePerTupleExprContext(m, v78)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	v192 = v187
	goto L57
L57:
	;
	v193 = int32(_a_F_ExecFindPartition_0)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v196
	v199 = v83 + int32(240)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	v201 = m.T0[v200].(func(*base.Module, int32, int32, int32) int32)(m, v186, v192, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v192 = v190
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v194
	v206 = v138 + int32(4)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if base.Ui32(v206) < base.Ui32(v209+v210<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v215 = v206
	goto L62
L61:
	;
	v215 = int32(0)
	goto L62
L62:
	;
	v216 = v199
	v217 = v215
	v218 = v201
	goto L46
L63:
	;
	goto L45
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v271 == int32(0) {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	switch v276 - int32(104) {
	case 0:
		goto L66
	default:
		goto L38
	case 4:
		goto L68
	case 10:
		goto L67
	}
L66:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v339 = F_compute_partition_hash_value(m, v247, v333, v334, v83+int32(96), v83-int32(-64))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L83
	}
L67:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+240)) = uint8(v328)
	if v328 < v247 {
		goto L37
	} else {
		goto L82
	}
L68:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+64)))
	if v279 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v275)+28))
	if v282 == int32(-1) {
		goto L34
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v83)+96))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if int32(16) <= v286 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v620 = v282
	goto L32
L73:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v325+v294)))
	v620 = v327
	goto L32
L74:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v294 = v292 << (uint(int32(2)) % 32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294+v295)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v299 = F_FunctionCall2Coll(m, v289, v291, v298, v285)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	v305 = v285
	goto L76
L76:
	;
	v306 = int32(-1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v311 = F_partition_list_bsearch(m, v307, v308, v275, v305, v83+int32(240))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	if v299 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v83)+96))
	v305 = v303
	goto L76
L79:
	;
	if v311 < int32(0) {
		v541 = v311
		v543 = v306
		goto L35
	} else {
		goto L80
	}
L80:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+240)))
	if v315&int32(1) == int32(0) {
		v541 = v311
		v543 = v306
		goto L35
	} else {
		goto L81
	}
L81:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320+v311<<(uint(int32(2))%32))))
	v541 = v311
	v543 = v324
	goto L35
L82:
	;
	goto L36
L83:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v342 = int64(*(*int32)(unsafe.Add(mBase, uint32(v275)+20)))
	v343 = base.I64_rem_u_s(v339, v342)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v341+base.I32_wrap_i64(v343)<<(uint(int32(2))%32))))
	v620 = v348
	goto L32
L84:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1337), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1348), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+48)) = v379
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_4), v83+int32(48))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1571), int32(_a_F_ExecFindPartition_5))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83-int32(-64)+v398))))
	if v428 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v434 = int32(-1)
	v541 = v434
	v543 = v434
	goto L35
L95:
	;
	v432 = v398 + int32(1)
	if v247 != v432 {
		v398 = v432
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L36
L99:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v524 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v529 = F_partition_range_datum_bsearch(m, v522, v523, v275, v524, v83+int32(96), v83+int32(240))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L109
	}
L100:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v477 = v475 << (uint(int32(2)) % 32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477+v478)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v481+v477)))
	v486 = F_partition_rbound_datum_cmp(m, v473, v474, v480, v483, v83+int32(96), v247)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v486 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v490+v477)+4))
	v620 = v492
	goto L32
L103:
	;
	goto L104
L104:
	;
	if int32(0) <= v486 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v496 = v475 + int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v497 <= v496 {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v502 = v496 << (uint(int32(2)) % 32)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v502+v503)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v506+v502)))
	v511 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v512 = F_partition_rbound_datum_cmp(m, v499, v500, v505, v508, v83+int32(96), v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v512 <= int32(0) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v516+v502)))
	v620 = v518
	goto L32
L109:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531+v529<<(uint(int32(2))%32))+4))
	v541 = v529
	v543 = v535
	goto L35
L110:
	;
	goto L34
L111:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = v609 + int32(1)
	v1042 = v543
	goto L29
L112:
	;
	goto L31
L113:
	;
	v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v683)+4)))
	v686 = int32(0)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v113)+56))
	v690 = F_check_enable_rls(m, v687, v686, int32(1))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L5
	} else {
		goto L115
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L5
	} else {
		goto L164
	}
L115:
	;
	if v690 == int32(2) {
		v976 = v686
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v697 = F_pg_class_aclcheck(m, v687, v695, int64(2))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v699 = int32(0)
	if base.B2i32(v697 == v699)|base.B2i32(v685 <= v699) == v699 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v714 = int32(0)
	goto L121
L119:
	;
	goto L120
L120:
	;
	v791 = v83 + int32(240)
	F_initStringInfo(m, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L127
	}
L121:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v683)+8))
	v745 = int32(*(*int16)(unsafe.Add(mBase, uint32(v741+v714<<(uint(int32(1))%32)))))
	if v745 == int32(0) {
		v976 = v686
		goto L114
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v751 = F_pg_attribute_aclcheck(m, v687, v745, v749, int64(2))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v751 != 0 {
		v976 = v686
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v754 = v714 + int32(1)
	if v754 != v685 {
		v714 = v754
		goto L121
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v797 = F_pg_get_partkeydef_worker(m, v687, int32(7), int32(1), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v797
	F_appendStringInfo(m, v791, int32(_a_F_ExecFindPartition_6), v83+int32(32))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v685 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_appendStringInfoChar(m, v83+int32(240), int32(41))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L163
	}
L131:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+64)))
	if v808 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v683)+32))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	F_getTypeOutputInfo(m, v812, v83+int32(236), v83+int32(235))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	v823 = int32(_a_F_ExecFindPartition_7)
	goto L134
L134:
	;
	v824 = F_strlen(m, v823)
	mBase = m.M
	if int32(65) <= v824 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v83)+236))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v83)+96))
	v821 = F_OidOutputFunctionCall(m, v819, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v823 = v821
	goto L134
L137:
	;
	v842 = int32(1)
	if v685 == v842 {
		goto L130
	} else {
		goto L145
	}
L138:
	;
	v828 = v83 + int32(240)
	v830 = F_pg_mbcliplen(m, v823, v824, int32(64))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_appendBinaryStringInfo(m, v83+int32(240), v823, v824)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L5
	} else {
		goto L144
	}
L141:
	;
	F_appendBinaryStringInfo(m, v828, v823, v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	F_appendStringInfoString(m, v828, int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	goto L137
L145:
	;
	v850 = v842
	goto L146
L146:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83-int32(-64)+v850))))
	if v883 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L130
L148:
	;
	v887 = v850 << (uint(int32(2)) % 32)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v683)+32))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v887+v888)))
	F_getTypeOutputInfo(m, v890, v83+int32(236), v83+int32(235))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	v905 = int32(_a_F_ExecFindPartition_7)
	goto L150
L150:
	;
	v907 = v83 + int32(240)
	F_appendStringInfoString(m, v907, int32(_a_F_ExecFindPartition_9))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v83)+236))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(96)+v887)))
	v902 = F_OidOutputFunctionCall(m, v897, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v905 = v902
	goto L150
L153:
	;
	v911 = F_strlen(m, v905)
	mBase = m.M
	if v911 <= int32(64) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v928 = v850 + int32(1)
	if v928 != v685 {
		v850 = v928
		goto L146
	} else {
		goto L162
	}
L155:
	;
	F_appendBinaryStringInfo(m, v907, v905, v911)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L5
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v917 = v83 + int32(240)
	v919 = F_pg_mbcliplen(m, v905, v911, int32(64))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L5
	} else {
		goto L159
	}
L158:
	;
	goto L154
L159:
	;
	F_appendBinaryStringInfo(m, v917, v905, v919)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	F_appendStringInfoString(m, v917, int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	goto L154
L162:
	;
	goto L147
L163:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v83)+240))
	v976 = v969
	goto L114
L164:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1011 + int32(4)
	F_errmsg(m, int32(_a_F_ExecFindPartition_10), v83+int32(16))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	if v976 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v976
	F_errdetail(m, int32(_a_F_ExecFindPartition_11), v83)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L5
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F_errtable(m, v113)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L5
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(335), int32(_a_F_ExecFindPartition_12))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2076)+32))
	if v2077 == v1042 {
		goto L352
	} else {
		goto L353
	}
L174:
	;
	if int32(0) <= v1072 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if int32(0) <= v1072 {
		goto L342
	} else {
		goto L343
	}
L177:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1080+v1072<<(uint(int32(2))%32))))
	v2047 = v1084
	v2050 = int32(0)
	v2056 = v88
	v2063 = v95
	goto L173
L178:
	;
	goto L179
L179:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1086+v1070)))
	v1091 = F_ExecLookupResultRelByOid(m, v74, v1088, int32(1), int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	if v1091 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v1094 != 0 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	v1105 = int32(0)
	v1107 = m.G0
	v1109 = v1107 - int32(16)
	m.G0 = v1109
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+8))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1112+v1042<<(uint(int32(2))%32))))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v74)+116))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+8))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+4))
	v1121 = int32(_a_F_ExecFindPartition_0)
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v76)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1124
	v1127 = F_table_open(m, v1116, int32(3))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L5
	} else {
		goto L190
	}
L184:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+132))
	v1097 = v1095
	goto L186
L185:
	;
	v1097 = int32(0)
	goto L186
L186:
	;
	F_CheckValidResultRel(m, v1091, int32(3), v1097, int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	F_ExecInitRoutingInfo(m, v74, v78, v76, v82, v1091, v1042, int32(1))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v2047 = v1091
	v2050 = int32(0)
	v2056 = v88
	v2063 = v95
	goto L173
L189:
	;
	v2047 = v1130
	v2050 = int32(0)
	v2056 = v88
	v2063 = v95
	goto L173
L190:
	;
	v1130 = F_palloc0(m, int32(216))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130))) = int32(388)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v78)+132))
	F_InitResultRelInfo(m, v1130, v1127, int32(0), v75, v1135)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v1117 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+132))
	v1141 = v1139
	goto L195
L194:
	;
	v1141 = int32(0)
	goto L195
L195:
	;
	F_CheckValidResultRel(m, v1130, int32(3), v1141, int32(0))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+116)))
	if v1146 != int32(1) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if v1117 != 0 {
		goto L206
	} else {
		goto L207
	}
L198:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+16))
	if v1149 != 0 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	if v1117 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+132))
	v1154 = base.B2i32(v1150 != int32(0))
	goto L202
L201:
	;
	v1154 = int32(0)
	goto L202
L202:
	;
	F_ExecOpenIndices(m, v1130, v1154)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	goto L197
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L5
	} else {
		goto L338
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1122
	m.G0 = v1109 + int32(16)
	goto L189
L206:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+100))
	if v1157 != 0 {
		goto L212
	} else {
		goto L213
	}
L207:
	;
	goto L208
L208:
	;
	F_ExecInitRoutingInfo(m, v74, v78, v76, v82, v1130, v1042, int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L5
	} else {
		goto L336
	}
L209:
	;
	v1395 = int32(0)
	F_ExecInitRoutingInfo(m, v74, v78, v76, v82, v1130, v1042, v1395)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L5
	} else {
		goto L231
	}
L210:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+72))
	v1352 = F_map_variable_attnos(m, v1326, v1120, v1324, v1349, v1109+int32(15))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L229
	}
L211:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+52))
	v1312 = F_build_attrmap_by_name(m, v1309, v1310, int32(0))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L5
	} else {
		goto L228
	}
L212:
	;
	v1158 = int32(0)
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+12))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+52))
	v1164 = F_build_attrmap_by_name(m, v1161, v1162, v1158)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L5
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+112))
	if v1269 == int32(0) {
		v1371 = v1105
		goto L209
	} else {
		goto L227
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+120)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+116)) = v1170
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+112))
	if v1262 == int32(0) {
		v1371 = v1164
		goto L209
	} else {
		goto L225
	}
L216:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+72))
	v1170 = F_map_variable_attnos(m, v1160, v1120, v1164, v1167, v1109+int32(15))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	if v1170 == int32(0) {
		v1241 = v1158
		goto L215
	} else {
		goto L218
	}
L218:
	;
	v1174 = int32(0)
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+4))
	if v1175 <= v1174 {
		v1241 = v1158
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v1190 = v1174
	v1193 = v1158
	goto L220
L220:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+12))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1212+v1190<<(uint(int32(2))%32))))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+16))
	v1218 = F_ExecInitQual(m, v1217, v74)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L5
	} else {
		goto L222
	}
L221:
	;
	v1241 = v1220
	goto L215
L222:
	;
	v1220 = F_lappend(m, v1193, v1218)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	v1223 = v1190 + int32(1)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+4))
	if v1223 < v1224 {
		v1190 = v1223
		v1193 = v1220
		goto L220
	} else {
		goto L224
	}
L224:
	;
	goto L221
L225:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+12))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1265)))
	if v1164 == int32(0) {
		v1308 = v1266
		goto L211
	} else {
		goto L226
	}
L226:
	;
	v1324 = v1164
	v1326 = v1266
	goto L210
L227:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+12))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	v1308 = v1273
	goto L211
L228:
	;
	v1324 = v1312
	v1326 = v1308
	goto L210
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+148)) = v1352
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v74)+60))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1358 = F_ExecBuildProjectionInfo(m, v1352, v1355, v1356, v74, v1357)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L5
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+152)) = v1358
	v1371 = v1324
	goto L209
L231:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+132))
	if v1399 == int32(0) {
		v1763 = v1371
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v78)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1788
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v78)+80))
	v1791 = F_lappend(m, v1790, v1130)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L5
	} else {
		goto L300
	}
L233:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v75)+156))
	if v1404 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+8))
	v1406 = F_RelationGetIndexList(m, v1405)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L5
	} else {
		goto L238
	}
L235:
	;
	v1641 = v1395
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+156)) = v1641
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+132))
	if v1663 != int32(2) {
		v1763 = v1371
		goto L232
	} else {
		goto L276
	}
L237:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v75)+156))
	if v1621 != 0 {
		goto L269
	} else {
		goto L270
	}
L238:
	;
	if v1406 == int32(0) {
		v1600 = v1395
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	if v1410 <= int32(0) {
		v1600 = v1395
		goto L237
	} else {
		goto L240
	}
L240:
	;
	v1422 = int32(0)
	v1427 = v1395
	goto L241
L241:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+12))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v1422<<(uint(int32(2))%32))))
	v1453 = F_get_partition_ancestors(m, v1452)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L5
	} else {
		goto L243
	}
L242:
	;
	v1600 = v1560
	goto L237
L243:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v75)+156))
	if v1455 == int32(0) {
		v1560 = v1427
		goto L244
	} else {
		goto L245
	}
L244:
	;
	F_list_free(m, v1453)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L5
	} else {
		goto L267
	}
L245:
	;
	v1458 = int32(0)
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+4))
	if v1459 <= v1458 {
		v1560 = v1427
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1474 = v1458
	v1475 = v1427
	goto L247
L247:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+12))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1496+v1474<<(uint(int32(2))%32))))
	v1501 = int32(0)
	if v1453 == v1501 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	v1560 = v1542
	goto L244
L249:
	;
	if v1539 != 0 {
		goto L262
	} else {
		goto L263
	}
L250:
	;
	v1539 = int32(0)
	goto L249
L251:
	;
	goto L252
L252:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+4))
	if v1507 <= int32(0) {
		v1533 = v1501
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1539 = v1533
	goto L249
L254:
	;
	v1510 = int32(0)
	if v1510 < v1507 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1513 = v1507
	goto L257
L256:
	;
	v1513 = v1510
	goto L257
L257:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+12))
	v1516 = int32(0)
	goto L258
L258:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1514+v1516<<(uint(int32(2))%32))))
	v1525 = base.B2i32(v1524 == v1500)
	if v1524 == v1500 {
		v1533 = v1525
		goto L253
	} else {
		goto L260
	}
L259:
	;
	v1533 = v1525
	goto L253
L260:
	;
	v1527 = v1516 + int32(1)
	if v1527 != v1513 {
		v1516 = v1527
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v1540 = F_lappend_oid(m, v1475, v1452)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L5
	} else {
		goto L265
	}
L263:
	;
	v1542 = v1475
	goto L264
L264:
	;
	v1544 = v1474 + int32(1)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+4))
	if v1544 < v1545 {
		v1474 = v1544
		v1475 = v1542
		goto L247
	} else {
		goto L266
	}
L265:
	;
	v1542 = v1540
	goto L264
L266:
	;
	goto L248
L267:
	;
	v1584 = v1422 + int32(1)
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	if v1584 < v1585 {
		v1422 = v1584
		v1427 = v1560
		goto L241
	} else {
		goto L268
	}
L268:
	;
	goto L242
L269:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+4))
	v1623 = v1622
	goto L271
L270:
	;
	v1623 = v1105
	goto L271
L271:
	;
	if v1600 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+4))
	v1626 = v1624
	goto L274
L273:
	;
	v1626 = int32(0)
	goto L274
L274:
	;
	if v1626 != v1623 {
		goto L204
	} else {
		goto L275
	}
L275:
	;
	v1641 = v1600
	goto L236
L276:
	;
	v1667 = F_palloc0(m, int32(20))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667))) = int32(386)
	v1671 = F_ExecGetRootToChildMap(m, v1130, v78)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+160)) = v1667
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+8))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v1678 = F_table_slot_create(m, v1674, v1675+int32(104))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+4)) = v1678
	if v1671 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+16)) = v1751
	v1763 = v1746
	goto L232
L281:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v75)+160))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1683)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+8)) = v1684
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v75)+160))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1686)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+12)) = v1687
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v75)+160))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+16))
	v1746 = v1371
	v1751 = v1690
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+140))
	v1692 = F_copyObjectImpl(m, v1691)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L5
	} else {
		goto L284
	}
L284:
	;
	if v1371 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+52))
	v1700 = F_build_attrmap_by_name(m, v1697, v1698, int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L5
	} else {
		goto L288
	}
L286:
	;
	v1702 = v1371
	goto L287
L287:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1703)+72))
	v1706 = v1109 + int32(15)
	v1707 = F_map_variable_attnos(m, v1692, int32(-1), v1702, v1704, v1706)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L5
	} else {
		goto L289
	}
L288:
	;
	v1702 = v1700
	goto L287
L289:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+72))
	v1711 = F_map_variable_attnos(m, v1707, v1120, v1702, v1710, v1706)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+144))
	v1714 = F_ExecGetChildToRootMap(m, v1130)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1714)+8))
	v1717 = F_adjust_partition_colnos_using_map(m, v1713, v1716)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L5
	} else {
		goto L292
	}
L292:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v1722 = F_table_slot_create(m, v1127, v1719+int32(104))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+8)) = v1722
	v1726 = F_ExecBuildUpdateProjection(m, v1711, int32(1), v1717, v1403, v1402, v1722, v74)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+12)) = v1726
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+148))
	if v1729 == int32(0) {
		v1763 = v1702
		goto L232
	} else {
		goto L295
	}
L295:
	;
	v1732 = F_copyObjectImpl(m, v1729)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+72))
	v1737 = F_map_variable_attnos(m, v1732, int32(-1), v1702, v1736, v1706)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+72))
	v1741 = F_map_variable_attnos(m, v1737, v1120, v1702, v1740, v1706)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v1743 = F_ExecInitQual(m, v1741, v74)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1746 = v1702
	v1751 = v1743
	goto L280
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+80)) = v1791
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+72))
	if v1794 != int32(5) {
		goto L205
	} else {
		goto L301
	}
L301:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v74)+64))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+160))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+12))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1799)))
	if v1763 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+52))
	v1806 = F_build_attrmap_by_name(m, v1803, v1804, int32(0))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L5
	} else {
		goto L305
	}
L303:
	;
	v1808 = v1763
	goto L304
L304:
	;
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130)+48)))
	if v1809 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v1808 = v1806
	goto L304
L306:
	;
	F_ExecInitMergeTupleSlots(m, v74, v1130)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L5
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+164))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+12))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1816)))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+72))
	v1822 = F_map_variable_attnos(m, v1817, v1120, v1808, v1819, v1109+int32(15))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L5
	} else {
		goto L310
	}
L309:
	;
	goto L308
L310:
	;
	v1824 = F_ExecInitQual(m, v1822, v74)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+176)) = v1824
	if v1800 == int32(0) {
		goto L205
	} else {
		goto L312
	}
L312:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+4))
	if v1829 <= int32(0) {
		goto L205
	} else {
		goto L313
	}
L313:
	;
	v1833 = v1130 + int32(164)
	v1851 = int32(0)
	goto L314
L314:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+12))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1868+v1851<<(uint(int32(2))%32))))
	v1873 = F_copyObjectImpl(m, v1872)
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L5
	} else {
		goto L316
	}
L315:
	;
	goto L205
L316:
	;
	v1876 = F_palloc0(m, int32(16))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1876)+4)) = v1873
	*(*int32)(unsafe.Add(mBase, uint32(v1876))) = int32(387)
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1833+v1881<<(uint(int32(2))%32))))
	v1886 = F_lappend(m, v1885, v1876)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	v1889 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1833+v1888<<(uint(v1889)%32)))) = v1886
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+8))
	switch v1893 - v1889 {
	case 0:
		goto L323
	case 1:
		goto L321
	case 2, 5:
		goto L319
	default:
		goto L322
	}
L319:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+16))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+48))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+72))
	v1936 = F_map_variable_attnos(m, v1931, v1120, v1808, v1933, v1109+int32(15))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L5
	} else {
		goto L333
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1876)+8)) = v1928
	goto L319
L321:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+20))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+40))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+52))
	v1925 = F_ExecBuildProjectionInfo(m, v1922, v1797, v1923, v74, v1924)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L5
	} else {
		goto L332
	}
L322:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L5
	} else {
		goto L329
	}
L323:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+24))
	if v1808 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1897 = F_adjust_partition_colnos_using_map(m, v1896, v1808)
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L5
	} else {
		goto L327
	}
L325:
	;
	v1900 = v1896
	goto L326
L326:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+20))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+8))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+52))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+40))
	v1907 = F_ExecBuildUpdateProjection(m, v1901, int32(1), v1900, v1904, v1797, v1905, int32(0))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L5
	} else {
		goto L328
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1873)+24)) = v1897
	v1900 = v1897
	goto L326
L328:
	;
	v1928 = v1907
	goto L320
L329:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_13), int32(0))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(968), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	v1928 = v1925
	goto L320
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1873)+16)) = v1936
	v1939 = F_ExecInitQual(m, v1936, v74)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1876)+12)) = v1939
	v1943 = v1851 + int32(1)
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+4))
	if v1943 < v1944 {
		v1851 = v1943
		goto L314
	} else {
		goto L335
	}
L335:
	;
	goto L315
L336:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v78)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1950
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v78)+80))
	v1953 = F_lappend(m, v1952, v1130)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+80)) = v1953
	goto L205
L338:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_15), int32(0))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(730), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L341:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+16))
	if v2032 == int32(0) {
		v2047 = v2031
		v2050 = v2030
		v2056 = v88
		v2063 = v95
		goto L173
	} else {
		goto L346
	}
L342:
	;
	v2012 = v1072 << (uint(int32(2)) % 32)
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2012+v98)))
	v2028 = v2012 + v2013
	v2030 = v2016
	goto L341
L343:
	;
	goto L344
L344:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2017+v1070)))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v74)+120))
	v2021 = F_ExecInitPartitionDispatchInfo(m, v78, v76, v2019, v82, v1042, v2020)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+24))
	v2028 = v2023 + v2024<<(uint(int32(2))%32)
	v2030 = v2021
	goto L341
L346:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+20))
	v2036 = F_execute_attr_map_slot(m, v2035, v88, v2032)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L5
	} else {
		goto L347
	}
L347:
	;
	if v95 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+12))
	m.T0[v2039].(func(*base.Module, int32))(m, v95)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L5
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	v2047 = v2031
	v2050 = v2030
	v2056 = v2036
	v2063 = v2032
	goto L173
L351:
	;
	goto L350
L352:
	;
	if v1075 == int32(0) {
		v2090 = v2056
		goto L355
	} else {
		goto L356
	}
L353:
	;
	v2095 = v2056
	goto L354
L354:
	;
	if v2050 != 0 {
		v82 = v2050
		v88 = v2095
		v95 = v2063
		goto L15
	} else {
		goto L361
	}
L355:
	;
	v2092 = F_ExecPartitionCheck(m, v2047, v2090, v78, int32(1))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L5
	} else {
		goto L360
	}
L356:
	;
	v2081 = F_ExecGetRootToChildMap(m, v2047, v78)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L5
	} else {
		goto L357
	}
L357:
	;
	if v2081 == int32(0) {
		v2090 = v77
		goto L355
	} else {
		goto L358
	}
L358:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+8))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+204))
	v2087 = F_execute_attr_map_slot(m, v2085, v77, v2086)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L5
	} else {
		goto L359
	}
L359:
	;
	v2090 = v2087
	goto L355
L360:
	;
	v2095 = v2090
	goto L354
L361:
	;
	goto L16
L362:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+8))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+12))
	m.T0[v2099].(func(*base.Module, int32))(m, v2063)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	v2107 = v2047
	v2111 = v83
	v2125 = v97
	v2127 = v99
	v2129 = v101
	goto L13
}
func F_ExecPartitionCheckEmitError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
		v20 = F_build_attrmap_by_name_if_req(m, v17, v18, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 == int32(0) {
				v32 = l1
				v33 = v13
				v34 = v18
				v35 = v15
				v37 = F_ExecGetInsertedCols(m, v33, l2)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = F_ExecGetUpdatedCols(m, v33, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = F_bms_union(m, v37, v39)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errcode(m, int32(67391682))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
										F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											if v43 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
												F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													F_errtable(m, v66)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_errtable(m, v66)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
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
							}
						}
					}
				}
			} else {
				v25 = F_MakeTupleTableSlot(m, v18, int32(_a_F_ExecPartitionCheckEmitError_4))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = F_execute_attr_map_slot(m, v20, l1, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v32 = v27
						v33 = v13
						v34 = v18
						v35 = v15
						v37 = F_ExecGetInsertedCols(m, v33, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = F_ExecGetUpdatedCols(m, v33, l2)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = F_bms_union(m, v37, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_errcode(m, int32(67391682))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
												F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													if v43 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
														F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return
														} else {
															v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															F_errtable(m, v66)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														F_errtable(m, v66)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return
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
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
		v32 = l1
		v33 = l0
		v34 = v30
		v35 = v31
		v37 = F_ExecGetInsertedCols(m, v33, l2)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = F_ExecGetUpdatedCols(m, v33, l2)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = F_bms_union(m, v37, v39)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errcode(m, int32(67391682))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
								F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v43 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
										F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_errtable(m, v66)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_errtable(m, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
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
					}
				}
			}
		}
	}
}
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F_get_partition_parent_worker(m, v12, l0, v8+int32(31))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				if l1 == int32(0) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v22&int32(1) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							F_errmsg_internal(m, int32(_a_F_get_partition_parent_0), v8+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(69), int32(_a_F_get_partition_parent_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_relation_close(m, v12, int32(1))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return v18
						}
					}
				} else {
					F_relation_close(m, v12, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return v18
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_get_partition_parent_3), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(65), int32(_a_F_get_partition_parent_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
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
}
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v4 {
		v136 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v136
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v136 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = v27
	goto L8
L7:
	;
	v28 = v4
	goto L8
L8:
	;
	if v25 <= int32(0) {
		v136 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = v28
	v35 = v4
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v35<<(uint(int32(1))%32)))))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v136 = int32(0)
	goto L1
L12:
	;
	v129 = v35 + int32(1)
	if v129 != v25 {
		v34 = v125
		v35 = v129
		goto L10
	} else {
		goto L42
	}
L13:
	;
	v48 = F_bms_is_member(m, v45+int32(7), l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	F_pull_varattnos(m, v57, int32(1), v13+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L19
	}
L16:
	;
	if v48 == int32(0) {
		v125 = v34
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(1)
	if l2 == int32(0) {
		v136 = v52
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	v136 = v52
	goto L1
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v68 = int32(0)
	if base.B2i32(l1 == v68)|base.B2i32(v67 == v68) != 0 {
		v113 = v68
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v113 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	goto L20
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v78 < v79 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = v78
	goto L25
L24:
	;
	v81 = v79
	goto L25
L25:
	;
	if v81 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = int32(1)
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v85 = int32(8)
	v90 = int32(0)
	goto L29
L29:
	;
	v97 = v90 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67+v85+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85+v97)))
	v102 = v99 & v101
	v104 = base.B2i32(v102 != int32(0))
	if v102 != 0 {
		v113 = v104
		goto L21
	} else {
		goto L31
	}
L30:
	;
	v113 = v104
	goto L21
L31:
	;
	v106 = v90 + int32(1)
	if v106 != v84 {
		v90 = v106
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v118 = v34 + int32(4)
	if base.Ui32(v118) < base.Ui32(v66+v65<<(uint(int32(2))%32)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v114)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v136 = int32(1)
	goto L1
L39:
	;
	v124 = v118
	goto L41
L40:
	;
	v124 = int32(0)
	goto L41
L41:
	;
	v125 = v124
	goto L12
L42:
	;
	goto L11
}
func F_partition_list_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v13 = v11 - int32(1)
	if v13 < int32(0) {
		v56 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v56
L2:
	;
	v21 = v10
	v22 = v13
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v30 = int32(2)
	v31 = base.I32_div_s(v21+v22+int32(1), v30)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(v30)%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = F_FunctionCall2Coll(m, l0, v25, v36, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v56 = v48
	goto L1
L5:
	;
	if v48 < v49 {
		v21 = v48
		v22 = v49
		goto L3
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v37 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(base.B2i32(v37 == int32(0)))
	if v37 != 0 {
		v48 = v31
		v49 = v22
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v48 = v21
	v49 = v31 - int32(1)
	goto L5
L11:
	;
	v56 = v31
	goto L1
L12:
	;
	goto L4
}
