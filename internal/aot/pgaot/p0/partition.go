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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
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
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
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
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v708 int32
	_ = v708
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v848 int32
	_ = v848
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
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
	var v1041 int32
	_ = v1041
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
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
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
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
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1241 int32
	_ = v1241
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1323 int32
	_ = v1323
	var v1332 int32
	_ = v1332
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1379 int32
	_ = v1379
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1426 int32
	_ = v1426
	var v1434 int32
	_ = v1434
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1471 int32
	_ = v1471
	var v1482 int32
	_ = v1482
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1567 int32
	_ = v1567
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1607 int32
	_ = v1607
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1648 int32
	_ = v1648
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1787 int32
	_ = v1787
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
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1873 int32
	_ = v1873
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2060 int32
	_ = v2060
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2103 int32
	_ = v2103
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
	v57 = int32(4554128)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+131)))
	if v64 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v54 = v40
	v55 = v40 + int32(4)
	v56 = v43
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
		v55 = v49
		v56 = v50
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
	v55 = v49
	v56 = v50
	goto L1
L9:
	;
	v68 = F_ExecPartitionCheck(m, l1, l3, l4, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v70 == int32(0) {
		v2103 = v6
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v58
	m.G0 = v37 + int32(256)
	return v2103
L14:
	;
	v83 = v70
	v88 = l3
	v95 = v6
	goto L15
L15:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v108 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v2060 == int32(0) {
		v2103 = v2043
		goto L13
	} else {
		goto L360
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v88
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v116 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v126)+4)))
	if int32(0) < v127 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v125 = v123
	goto L21
L23:
	;
	if v114 != 0 {
		v122 = v114
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v120 = v114
	goto L25
L25:
	;
	if v120 != 0 {
		v122 = v120
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v117 = F_ExecPrepareExprList(m, v116, l4)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v117
	v120 = v117
	goto L25
L28:
	;
	v125 = int32(0)
	goto L21
L29:
	;
	v1070 = v1041 << (uint(int32(2)) % 32)
	v1073 = v83 + v1070 + int32(24)
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1041))))
	if v1077 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+20)) = v539
	v1041 = v540
	goto L29
L31:
	;
	v681 = F_RelationGetPartitionKey(m, v112)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L113
	}
L32:
	;
	if int32(0) <= v617 {
		v1041 = v617
		goto L29
	} else {
		goto L112
	}
L33:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	if v539 != v605 {
		goto L30
	} else {
		goto L111
	}
L34:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v275)+32))
	v617 = v604
	goto L32
L35:
	;
	if int32(0) <= v540 {
		goto L33
	} else {
		goto L110
	}
L36:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if v468 < int32(16) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	v395 = v326
	goto L93
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L90
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L87
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L84
	}
L41:
	;
	v136 = v126
	v137 = int32(0)
	v140 = v125
	goto L44
L42:
	;
	v242 = v126
	v246 = v125
	v249 = v127
	goto L43
L43:
	;
	if v246 != 0 {
		goto L39
	} else {
		goto L64
	}
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165+v137<<(uint(int32(1))%32)))))
	if v169 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v242 = v234
	v246 = v219
	v249 = v235
	goto L43
L46:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v137))) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(96)+v137<<(uint(int32(2))%32)))) = v220
	v233 = v137 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+4)))
	if v233 < v235 {
		v136 = v234
		v137 = v233
		v140 = v219
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
	if v170 < v169 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v140 == int32(0) {
		goto L40
	} else {
		goto L54
	}
L50:
	;
	F_slot_getsomeattrs_int(m, v88, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v175 = v169 - int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v175<<(uint(int32(2))%32))))
	v217 = v175 + v176
	v219 = v140
	v220 = v182
	goto L46
L53:
	;
	goto L52
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v186 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v189 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	v191 = v186
	goto L57
L57:
	;
	v192 = int32(4554128)
	v193 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v195
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v185)+20))
	v200 = m.T0[v199].(func(*base.Module, int32, int32, int32) int32)(m, v185, v191, v37+int32(240))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v191 = v189
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v193
	v205 = v140 + int32(4)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if base.Ui32(v205) < base.Ui32(v208+v209<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v214 = v205
	goto L62
L61:
	;
	v214 = int32(0)
	goto L62
L62:
	;
	v217 = v37 + int32(240)
	v219 = v214
	v220 = v200
	goto L46
L63:
	;
	goto L45
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v271 == int32(0) {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
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
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v337 = F_compute_partition_hash_value(m, v249, v331, v332, v37+int32(96), v37-int32(-64))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L83
	}
L67:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+240)) = uint8(v326)
	if v326 < v249 {
		goto L37
	} else {
		goto L82
	}
L68:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+64)))
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
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if int32(16) <= v286 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v617 = v282
	goto L32
L73:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323+v294)))
	v617 = v325
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
	v311 = F_partition_list_bsearch(m, v307, v308, v275, v305, v37+int32(240))
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
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v305 = v303
	goto L76
L79:
	;
	if v311 < int32(0) {
		v539 = v311
		v540 = v306
		goto L35
	} else {
		goto L80
	}
L80:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+240)))
	if v315 != int32(1) {
		v539 = v311
		v540 = v306
		goto L35
	} else {
		goto L81
	}
L81:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+v311<<(uint(int32(2))%32))))
	v539 = v311
	v540 = v322
	goto L35
L82:
	;
	goto L36
L83:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v340 = int64(*(*int32)(unsafe.Add(mBase, uint32(v275)+20)))
	v341 = base.I64_rem_u_s(v337, v340)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v339+base.I32_wrap_i64(v341)<<(uint(int32(2))%32))))
	v617 = v346
	goto L32
L84:
	;
	F_errmsg_internal(m, int32(153103), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(519337), int32(1337), int32(299548))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
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
	F_errmsg_internal(m, int32(153103), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(519337), int32(1348), int32(299548))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
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
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v377
	F_errmsg_internal(m, int32(502140), v37+int32(48))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(519337), int32(1571), int32(401267))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
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
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v395))))
	if v426 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v432 = int32(-1)
	v539 = v432
	v540 = v432
	goto L35
L95:
	;
	v430 = v395 + int32(1)
	if v249 != v430 {
		v395 = v430
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
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v527 = F_partition_range_datum_bsearch(m, v520, v521, v275, v522, v37+int32(96), v37+int32(240))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L5
	} else {
		goto L109
	}
L100:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v475 = v473 << (uint(int32(2)) % 32)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475+v476)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v479+v475)))
	v484 = F_partition_rbound_datum_cmp(m, v471, v472, v478, v481, v37+int32(96), v249)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v484 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488+v475)+4))
	v617 = v490
	goto L32
L103:
	;
	goto L104
L104:
	;
	if int32(0) <= v484 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v494 = v473 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v495 <= v494 {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v500 = v494 << (uint(int32(2)) % 32)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v500+v501)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504+v500)))
	v509 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v510 = F_partition_rbound_datum_cmp(m, v497, v498, v503, v506, v37+int32(96), v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v510 <= int32(0) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514+v500)))
	v617 = v516
	goto L32
L109:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529+v527<<(uint(int32(2))%32))+4))
	v539 = v527
	v540 = v533
	goto L35
L110:
	;
	goto L34
L111:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = v607 + int32(1)
	v1041 = v540
	goto L29
L112:
	;
	goto L31
L113:
	;
	v683 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681)+4)))
	v684 = int32(0)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	v688 = F_check_enable_rls(m, v685, v684, int32(1))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
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
	if v688 == int32(2) {
		v979 = v684
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v695 = F_pg_class_aclcheck(m, v685, v693, int64(2))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L5
	} else {
		goto L118
	}
L117:
	;
	F_initStringInfo(m, v37+int32(240))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L127
	}
L118:
	;
	if v695 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	if v683 <= int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v708 = int32(0)
	goto L121
L121:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v681)+8))
	v740 = int32(*(*int16)(unsafe.Add(mBase, uint32(v736+v708<<(uint(int32(1))%32)))))
	if v740 == int32(0) {
		v979 = v684
		goto L114
	} else {
		goto L123
	}
L122:
	;
	goto L117
L123:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v746 = F_pg_attribute_aclcheck(m, v685, v740, v744, int64(2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v746 != 0 {
		v979 = v684
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v749 = v708 + int32(1)
	if v749 != v683 {
		v708 = v749
		goto L121
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v792 = F_pg_get_partkeydef_worker(m, v685, int32(7), int32(1), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v792
	F_appendStringInfo(m, v37+int32(240), int32(716843), v37+int32(32))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v683 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_appendStringInfoChar(m, v37+int32(240), int32(41))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L163
	}
L131:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+64)))
	if v805 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v681)+32))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	F_getTypeOutputInfo(m, v809, v37+int32(236), v37+int32(235))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	v820 = int32(317344)
	goto L134
L134:
	;
	v821 = F_strlen(m, v820)
	mBase = m.M
	if int32(65) <= v821 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v818 = F_OidOutputFunctionCall(m, v816, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v820 = v818
	goto L134
L137:
	;
	v840 = int32(1)
	if v683 == v840 {
		goto L130
	} else {
		goto L145
	}
L138:
	;
	v827 = F_pg_mbcliplen(m, v820, v821, int32(64))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v820, v821)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L144
	}
L141:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v820, v827)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(689203))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
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
	v848 = v840
	goto L146
L146:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v848))))
	if v881 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L130
L148:
	;
	v885 = v848 << (uint(int32(2)) % 32)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v681)+32))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v885+v886)))
	F_getTypeOutputInfo(m, v888, v37+int32(236), v37+int32(235))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	v902 = int32(317344)
	goto L150
L150:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(778892))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(96)+v885)))
	v900 = F_OidOutputFunctionCall(m, v895, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v902 = v900
	goto L150
L153:
	;
	v908 = F_strlen(m, v902)
	mBase = m.M
	if v908 <= int32(64) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v928 = v848 + int32(1)
	if v928 != v683 {
		v848 = v928
		goto L146
	} else {
		goto L162
	}
L155:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v902, v908)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L5
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v918 = F_pg_mbcliplen(m, v902, v908, int32(64))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L5
	} else {
		goto L159
	}
L158:
	;
	goto L154
L159:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v902, v918)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(689203))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
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
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v37)+240))
	v979 = v969
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
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v112)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v1011 + int32(4)
	F_errmsg(m, int32(31555), v37+int32(16))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	if v979 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v979
	F_errdetail(m, int32(630805), v37)
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
	F_errtable(m, v112)
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
	F_errfinish(m, int32(519337), int32(335), int32(261471))
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
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+32))
	if v2073 == v1041 {
		goto L350
	} else {
		goto L351
	}
L174:
	;
	if int32(0) <= v1074 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if int32(0) <= v1074 {
		goto L340
	} else {
		goto L341
	}
L177:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1082+v1074<<(uint(int32(2))%32))))
	v2043 = v1086
	v2048 = int32(0)
	v2053 = v88
	v2060 = v95
	goto L173
L178:
	;
	goto L179
L179:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1088+v1070)))
	v1093 = F_ExecLookupResultRelByOid(m, l0, v1090, int32(1), int32(0))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	if v1093 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1096 != 0 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	v1107 = int32(0)
	v1109 = m.G0
	v1111 = v1109 - int32(16)
	m.G0 = v1111
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+8))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1114+v1041<<(uint(int32(2))%32))))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+8))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+4))
	v1123 = int32(4554128)
	v1124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1126
	v1129 = F_table_open(m, v1118, int32(3))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L5
	} else {
		goto L189
	}
L184:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+132))
	v1099 = v1097
	goto L186
L185:
	;
	v1099 = int32(0)
	goto L186
L186:
	;
	F_CheckValidResultRel(m, v1093, int32(3), v1099, int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1093, v1041, int32(1))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v2043 = v1093
	v2048 = int32(0)
	v2053 = v88
	v2060 = v95
	goto L173
L189:
	;
	v1132 = F_palloc0(m, int32(216))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132))) = int32(388)
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+132))
	F_InitResultRelInfo(m, v1132, v1129, int32(0), l1, v1137)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	if v1119 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+132))
	v1143 = v1141
	goto L194
L193:
	;
	v1143 = int32(0)
	goto L194
L194:
	;
	F_CheckValidResultRel(m, v1132, int32(3), v1143, int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147)+116)))
	if v1148 != int32(1) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	if v1119 != 0 {
		goto L204
	} else {
		goto L205
	}
L197:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+16))
	if v1151 != 0 {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	if v1119 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+132))
	v1156 = base.B2i32(v1152 != int32(0))
	goto L201
L200:
	;
	v1156 = int32(0)
	goto L201
L201:
	;
	F_ExecOpenIndices(m, v1132, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L5
	} else {
		goto L202
	}
L202:
	;
	goto L196
L203:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1806
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v1809 = F_lappend(m, v1808, v1132)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L5
	} else {
		goto L301
	}
L204:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+100))
	if v1159 != 0 {
		goto L210
	} else {
		goto L211
	}
L205:
	;
	goto L206
L206:
	;
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1132, v1041, int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L5
	} else {
		goto L300
	}
L207:
	;
	v1397 = int32(0)
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1132, v1041, v1397)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L5
	} else {
		goto L229
	}
L208:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+72))
	v1354 = F_map_variable_attnos(m, v1323, v1122, v1332, v1351, v1111+int32(15))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L5
	} else {
		goto L227
	}
L209:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+52))
	v1314 = F_build_attrmap_by_name(m, v1311, v1312, int32(0))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L5
	} else {
		goto L226
	}
L210:
	;
	v1160 = int32(0)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+12))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1161)))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+52))
	v1166 = F_build_attrmap_by_name(m, v1163, v1164, v1160)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L5
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+112))
	if v1271 == int32(0) {
		v1379 = v1107
		goto L207
	} else {
		goto L225
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+120)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+116)) = v1172
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+112))
	if v1264 == int32(0) {
		v1379 = v1166
		goto L207
	} else {
		goto L223
	}
L214:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+72))
	v1172 = F_map_variable_attnos(m, v1162, v1122, v1166, v1169, v1111+int32(15))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	if v1172 == int32(0) {
		v1241 = v1160
		goto L213
	} else {
		goto L216
	}
L216:
	;
	v1176 = int32(0)
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1177 <= v1176 {
		v1241 = v1160
		goto L213
	} else {
		goto L217
	}
L217:
	;
	v1187 = v1176
	v1193 = v1160
	goto L218
L218:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+12))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1214+v1187<<(uint(int32(2))%32))))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+16))
	v1220 = F_ExecInitQual(m, v1219, l0)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L5
	} else {
		goto L220
	}
L219:
	;
	v1241 = v1222
	goto L213
L220:
	;
	v1222 = F_lappend(m, v1193, v1220)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	v1225 = v1187 + int32(1)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	if v1225 < v1226 {
		v1187 = v1225
		v1193 = v1222
		goto L218
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1267)))
	if v1166 == int32(0) {
		v1310 = v1268
		goto L209
	} else {
		goto L224
	}
L224:
	;
	v1323 = v1268
	v1332 = v1166
	goto L208
L225:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+12))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1274)))
	v1310 = v1275
	goto L209
L226:
	;
	v1323 = v1310
	v1332 = v1314
	goto L208
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+148)) = v1354
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1360 = F_ExecBuildProjectionInfo(m, v1354, v1357, v1358, l0, v1359)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L5
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+152)) = v1360
	v1379 = v1332
	goto L207
L229:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+132))
	if v1401 == int32(0) {
		v1787 = v1379
		goto L203
	} else {
		goto L230
	}
L230:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1406 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L5
	} else {
		goto L297
	}
L232:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+8))
	v1408 = F_RelationGetIndexList(m, v1407)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L5
	} else {
		goto L236
	}
L233:
	;
	v1648 = v1397
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+156)) = v1648
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+132))
	if v1665 != int32(2) {
		v1787 = v1379
		goto L203
	} else {
		goto L274
	}
L235:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1623 != 0 {
		goto L267
	} else {
		goto L268
	}
L236:
	;
	if v1408 == int32(0) {
		v1607 = v1397
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1412 = int32(0)
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+4))
	if v1413 <= v1412 {
		v1607 = v1397
		goto L235
	} else {
		goto L238
	}
L238:
	;
	v1426 = v1412
	v1434 = v1397
	goto L239
L239:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+12))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1450+v1426<<(uint(int32(2))%32))))
	v1455 = F_get_partition_ancestors(m, v1454)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L5
	} else {
		goto L241
	}
L240:
	;
	v1607 = v1567
	goto L235
L241:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1457 == int32(0) {
		v1567 = v1434
		goto L242
	} else {
		goto L243
	}
L242:
	;
	F_list_free(m, v1455)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L5
	} else {
		goto L265
	}
L243:
	;
	v1460 = int32(0)
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+4))
	if v1461 <= v1460 {
		v1567 = v1434
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v1471 = v1460
	v1482 = v1434
	goto L245
L245:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+12))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1498+v1471<<(uint(int32(2))%32))))
	v1503 = int32(0)
	if v1455 == v1503 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	v1567 = v1544
	goto L242
L247:
	;
	if v1541 != 0 {
		goto L260
	} else {
		goto L261
	}
L248:
	;
	v1541 = int32(0)
	goto L247
L249:
	;
	goto L250
L250:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+4))
	if v1509 <= int32(0) {
		v1534 = v1503
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1541 = v1534
	goto L247
L252:
	;
	v1512 = int32(0)
	if v1512 < v1509 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1515 = v1509
	goto L255
L254:
	;
	v1515 = v1512
	goto L255
L255:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+12))
	v1518 = int32(0)
	goto L256
L256:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1516+v1518<<(uint(int32(2))%32))))
	v1527 = base.B2i32(v1526 == v1502)
	if v1526 == v1502 {
		v1534 = v1527
		goto L251
	} else {
		goto L258
	}
L257:
	;
	v1534 = v1527
	goto L251
L258:
	;
	v1529 = v1518 + int32(1)
	if v1529 != v1515 {
		v1518 = v1529
		goto L256
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	v1542 = F_lappend_oid(m, v1482, v1454)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L5
	} else {
		goto L263
	}
L261:
	;
	v1544 = v1482
	goto L262
L262:
	;
	v1546 = v1471 + int32(1)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+4))
	if v1546 < v1547 {
		v1471 = v1546
		v1482 = v1544
		goto L245
	} else {
		goto L264
	}
L263:
	;
	v1544 = v1542
	goto L262
L264:
	;
	goto L246
L265:
	;
	v1586 = v1426 + int32(1)
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+4))
	if v1586 < v1587 {
		v1426 = v1586
		v1434 = v1567
		goto L239
	} else {
		goto L266
	}
L266:
	;
	goto L240
L267:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+4))
	v1625 = v1624
	goto L269
L268:
	;
	v1625 = v1107
	goto L269
L269:
	;
	if v1607 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1607)+4))
	v1628 = v1626
	goto L272
L271:
	;
	v1628 = int32(0)
	goto L272
L272:
	;
	if v1628 != v1625 {
		goto L231
	} else {
		goto L273
	}
L273:
	;
	v1648 = v1607
	goto L234
L274:
	;
	v1669 = F_palloc0(m, int32(20))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669))) = int32(386)
	v1673 = F_ExecGetRootToChildMap(m, v1132, l4)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+160)) = v1669
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+8))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1680 = F_table_slot_create(m, v1676, v1677+int32(104))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+4)) = v1680
	if v1673 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+8)) = v1686
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+12)) = v1689
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+16)) = v1692
	v1787 = v1379
	goto L203
L279:
	;
	goto L280
L280:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+140))
	v1695 = F_copyObjectImpl(m, v1694)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	if v1379 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+52))
	v1703 = F_build_attrmap_by_name(m, v1700, v1701, int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L5
	} else {
		goto L285
	}
L283:
	;
	v1705 = v1379
	goto L284
L284:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+72))
	v1710 = F_map_variable_attnos(m, v1695, int32(-1), v1705, v1707, v1111+int32(15))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L5
	} else {
		goto L286
	}
L285:
	;
	v1705 = v1703
	goto L284
L286:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1712)+72))
	v1716 = F_map_variable_attnos(m, v1710, v1122, v1705, v1713, v1111+int32(15))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+144))
	v1719 = F_ExecGetChildToRootMap(m, v1132)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+8))
	v1722 = F_adjust_partition_colnos_using_map(m, v1718, v1721)
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1727 = F_table_slot_create(m, v1129, v1724+int32(104))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+8)) = v1727
	v1731 = F_ExecBuildUpdateProjection(m, v1716, int32(1), v1722, v1405, v1404, v1727, l0)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+12)) = v1731
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+148))
	if v1734 == int32(0) {
		v1787 = v1705
		goto L203
	} else {
		goto L292
	}
L292:
	;
	v1737 = F_copyObjectImpl(m, v1734)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1740)+72))
	v1744 = F_map_variable_attnos(m, v1737, int32(-1), v1705, v1741, v1111+int32(15))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+72))
	v1750 = F_map_variable_attnos(m, v1744, v1122, v1705, v1747, v1111+int32(15))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	v1752 = F_ExecInitQual(m, v1750, l0)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+16)) = v1752
	v1787 = v1705
	goto L203
L297:
	;
	F_errmsg_internal(m, int32(79967), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(519337), int32(730), int32(254237))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L5
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
	v1787 = v1107
	goto L203
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v1809
	if v1119 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1124
	m.G0 = v1111 + int32(16)
	v2043 = v1132
	v2048 = int32(0)
	v2053 = v88
	v2060 = v95
	goto L173
L303:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+72))
	if v1814 != int32(5) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+160))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+12))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1819)))
	if v1787 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+52))
	v1826 = F_build_attrmap_by_name(m, v1823, v1824, int32(0))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L5
	} else {
		goto L308
	}
L306:
	;
	v1828 = v1787
	goto L307
L307:
	;
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+48)))
	if v1829 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v1828 = v1826
	goto L307
L309:
	;
	F_ExecInitMergeTupleSlots(m, l0, v1132)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L5
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+164))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+12))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1836)))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+72))
	v1842 = F_map_variable_attnos(m, v1837, v1122, v1828, v1839, v1111+int32(15))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L5
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v1844 = F_ExecInitQual(m, v1842, l0)
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L5
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+176)) = v1844
	if v1820 == int32(0) {
		goto L302
	} else {
		goto L315
	}
L315:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+4))
	if v1849 <= int32(0) {
		goto L302
	} else {
		goto L316
	}
L316:
	;
	v1853 = v1132 + int32(164)
	v1873 = int32(0)
	goto L317
L317:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+12))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1888+v1873<<(uint(int32(2))%32))))
	v1893 = F_copyObjectImpl(m, v1892)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L5
	} else {
		goto L319
	}
L318:
	;
	goto L302
L319:
	;
	v1896 = F_palloc0(m, int32(16))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+4)) = v1893
	*(*int32)(unsafe.Add(mBase, uint32(v1896))) = int32(387)
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+4))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1853+v1901<<(uint(int32(2))%32))))
	v1906 = F_lappend(m, v1905, v1896)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+4))
	v1909 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1853+v1908<<(uint(v1909)%32)))) = v1906
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+8))
	switch v1913 - v1909 {
	case 0:
		goto L326
	case 1:
		goto L324
	case 2, 5:
		goto L322
	default:
		goto L325
	}
L322:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+16))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+48))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1952)+72))
	v1956 = F_map_variable_attnos(m, v1951, v1122, v1828, v1953, v1111+int32(15))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L5
	} else {
		goto L336
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+8)) = v1948
	goto L322
L324:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+20))
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+40))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+52))
	v1945 = F_ExecBuildProjectionInfo(m, v1942, v1817, v1943, l0, v1944)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L5
	} else {
		goto L335
	}
L325:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L5
	} else {
		goto L332
	}
L326:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+24))
	if v1828 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1917 = F_adjust_partition_colnos_using_map(m, v1916, v1828)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L5
	} else {
		goto L330
	}
L328:
	;
	v1920 = v1916
	goto L329
L329:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1893)+20))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+8))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+52))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+40))
	v1927 = F_ExecBuildUpdateProjection(m, v1921, int32(1), v1920, v1924, v1817, v1925, int32(0))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L5
	} else {
		goto L331
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1893)+24)) = v1917
	v1920 = v1917
	goto L329
L331:
	;
	v1948 = v1927
	goto L323
L332:
	;
	F_errmsg_internal(m, int32(376128), int32(0))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(519337), int32(968), int32(254237))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	v1948 = v1945
	goto L323
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1893)+16)) = v1956
	v1959 = F_ExecInitQual(m, v1956, l0)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+12)) = v1959
	v1963 = v1873 + int32(1)
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+4))
	if v1963 < v1964 {
		v1873 = v1963
		goto L317
	} else {
		goto L338
	}
L338:
	;
	goto L318
L339:
	;
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2025)))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+16))
	if v2028 == int32(0) {
		v2043 = v2027
		v2048 = v2026
		v2053 = v88
		v2060 = v95
		goto L173
	} else {
		goto L344
	}
L340:
	;
	v2009 = v1074 << (uint(int32(2)) % 32)
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2009+v39)))
	v2025 = v2009 + v2010
	v2026 = v2013
	goto L339
L341:
	;
	goto L342
L342:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2014+v1070)))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2018 = F_ExecInitPartitionDispatchInfo(m, l4, l2, v2016, v83, v1041, v2017)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L5
	} else {
		goto L343
	}
L343:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	v2025 = v2020 + v2021<<(uint(int32(2))%32)
	v2026 = v2018
	goto L339
L344:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+20))
	v2032 = F_execute_attr_map_slot(m, v2031, v88, v2028)
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	if v95 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+12))
	m.T0[v2035].(func(*base.Module, int32))(m, v95)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L5
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v2043 = v2027
	v2048 = v2026
	v2053 = v2032
	v2060 = v2028
	goto L173
L349:
	;
	goto L348
L350:
	;
	if v1077 == int32(0) {
		v2086 = v2053
		goto L353
	} else {
		goto L354
	}
L351:
	;
	v2091 = v2053
	goto L352
L352:
	;
	if v2048 != 0 {
		v83 = v2048
		v88 = v2091
		v95 = v2060
		goto L15
	} else {
		goto L359
	}
L353:
	;
	v2088 = F_ExecPartitionCheck(m, v2043, v2086, l4, int32(1))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L5
	} else {
		goto L358
	}
L354:
	;
	v2077 = F_ExecGetRootToChildMap(m, v2043, l4)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	if v2077 == int32(0) {
		v2086 = l3
		goto L353
	} else {
		goto L356
	}
L356:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+8))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2043)+204))
	v2083 = F_execute_attr_map_slot(m, v2081, l3, v2082)
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L5
	} else {
		goto L357
	}
L357:
	;
	v2086 = v2083
	goto L353
L358:
	;
	v2091 = v2086
	goto L352
L359:
	;
	goto L16
L360:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2060)+8))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+12))
	m.T0[v2095].(func(*base.Module, int32))(m, v2060)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L5
	} else {
		goto L361
	}
L361:
	;
	v2103 = v2043
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
										F_errmsg(m, int32(96051), v11+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											if v43 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
												F_errdetail(m, int32(630851), v11)
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
														F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
													F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
				v25 = F_MakeTupleTableSlot(m, v18, int32(1650644))
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
												F_errmsg(m, int32(96051), v11+int32(16))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													if v43 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
														F_errdetail(m, int32(630851), v11)
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
																F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
															F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
								F_errmsg(m, int32(96051), v11+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v43 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
										F_errdetail(m, int32(630851), v11)
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
												F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
											F_errfinish(m, int32(519741), int32(1969), int32(223230))
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
							F_errmsg_internal(m, int32(479754), v8+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519325), int32(69), int32(100144))
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
						F_sequence_close(m, v12, int32(1))
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
					F_sequence_close(m, v12, int32(1))
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
					F_errmsg_internal(m, int32(50041), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519325), int32(65), int32(100144))
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v4 {
		v134 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v134
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v134 = v4
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
		v134 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = v4
	v35 = v28
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v34<<(uint(int32(1))%32)))))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v134 = int32(0)
	goto L1
L12:
	;
	v128 = v34 + int32(1)
	if v128 != v25 {
		v34 = v128
		v35 = v125
		goto L10
	} else {
		goto L43
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
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
		v125 = v35
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(1)
	if l2 == int32(0) {
		v134 = v52
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	v134 = v52
	goto L1
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v68 = int32(0)
	if l1 == v68 {
		v109 = v68
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v109 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	goto L20
L22:
	;
	if v67 == int32(0) {
		v109 = v68
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v77 < v78 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v77
	goto L26
L25:
	;
	v80 = v78
	goto L26
L26:
	;
	if v80 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = int32(1)
	goto L29
L28:
	;
	v83 = v80
	goto L29
L29:
	;
	v84 = int32(8)
	v89 = int32(0)
	goto L30
L30:
	;
	v96 = v89 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v67+v84+v96)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+(l1+v84))))
	v101 = v98 & v100
	v103 = base.B2i32(v101 != int32(0))
	if v101 != 0 {
		v109 = v103
		goto L21
	} else {
		goto L32
	}
L31:
	;
	v109 = v103
	goto L21
L32:
	;
	v105 = v89 + int32(1)
	if v105 != v83 {
		v89 = v105
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v117 = v35 + int32(4)
	if base.Ui32(v117) < base.Ui32(v66+v65<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v113)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v134 = int32(1)
	goto L1
L40:
	;
	v123 = v117
	goto L42
L41:
	;
	v123 = int32(0)
	goto L42
L42:
	;
	v125 = v123
	goto L12
L43:
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
