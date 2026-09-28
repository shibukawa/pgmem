package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int64
	_ = v160
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
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
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int64
	_ = v272
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v826 int64
	_ = v826
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int64
	_ = v969
	var v970 int64
	_ = v970
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1111 int32
	_ = v1111
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1156 int32
	_ = v1156
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1310 int32
	_ = v1310
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1371 int32
	_ = v1371
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int64
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1455 int32
	_ = v1455
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1480 int32
	_ = v1480
	var v1493 int32
	_ = v1493
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1538 int32
	_ = v1538
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1563 int32
	_ = v1563
	var v1576 int32
	_ = v1576
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1617 int32
	_ = v1617
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1659 int64
	_ = v1659
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
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
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1775 int32
	_ = v1775
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1790 int32
	_ = v1790
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1863 int32
	_ = v1863
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1906 int32
	_ = v1906
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2295 int32
	_ = v2295
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2350 int32
	_ = v2350
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2414 int32
	_ = v2414
	var v2420 int32
	_ = v2420
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2482 int32
	_ = v2482
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	var v2503 int32
	_ = v2503
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2531 int32
	_ = v2531
	var v2537 int32
	_ = v2537
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2559 int32
	_ = v2559
	var v2563 int32
	_ = v2563
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2620 int32
	_ = v2620
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2639 int32
	_ = v2639
	var v2643 int32
	_ = v2643
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2703 int32
	_ = v2703
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2758 int64
	_ = v2758
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2789 int64
	_ = v2789
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
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
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2836 int32
	_ = v2836
	var v2844 int32
	_ = v2844
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v3000 int32
	_ = v3000
	var v3012 int32
	_ = v3012
	var v3028 int32
	_ = v3028
	var v3044 int32
	_ = v3044
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3077 int32
	_ = v3077
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3103 int32
	_ = v3103
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3181 int32
	_ = v3181
	var v3184 int32
	_ = v3184
	var v3201 int32
	_ = v3201
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3217 int32
	_ = v3217
	var v3220 int32
	_ = v3220
	var v3233 int32
	_ = v3233
	var v3249 int32
	_ = v3249
	var v3267 int32
	_ = v3267
	var v3284 int32
	_ = v3284
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3332 int32
	_ = v3332
	var v3351 int32
	_ = v3351
	var v3361 int32
	_ = v3361
	var v3382 int32
	_ = v3382
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3396 int32
	_ = v3396
	var v3411 int32
	_ = v3411
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3440 int32
	_ = v3440
	var v3452 int32
	_ = v3452
	var v3466 int32
	_ = v3466
	var v3478 int32
	_ = v3478
	var v3485 int32
	_ = v3485
	var v3492 int32
	_ = v3492
	var v3504 int32
	_ = v3504
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3541 int32
	_ = v3541
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3573 int32
	_ = v3573
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3585 int32
	_ = v3585
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3597 int32
	_ = v3597
	var v3601 int32
	_ = v3601
	var v3606 int32
	_ = v3606
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3618 int32
	_ = v3618
	var v3625 int32
	_ = v3625
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3635 int32
	_ = v3635
	var v3640 int32
	_ = v3640
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3655 int32
	_ = v3655
	var v3657 int32
	_ = v3657
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3671 int32
	_ = v3671
	var v3674 int32
	_ = v3674
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3686 int32
	_ = v3686
	var v3688 int32
	_ = v3688
	var v3689 int64
	_ = v3689
	var v3697 int32
	_ = v3697
	var v3701 int32
	_ = v3701
	var v3705 int32
	_ = v3705
	var v3711 int32
	_ = v3711
	var v3715 int32
	_ = v3715
	var v3716 int32
	_ = v3716
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3729 int32
	_ = v3729
	var v3732 int32
	_ = v3732
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3745 int32
	_ = v3745
	var v3751 int32
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3755 int32
	_ = v3755
	var v3765 int32
	_ = v3765
	var v3773 int32
	_ = v3773
	var v3774 int32
	_ = v3774
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3783 int32
	_ = v3783
	var v3791 int32
	_ = v3791
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3803 int32
	_ = v3803
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3810 int32
	_ = v3810
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3817 int32
	_ = v3817
	var v3818 int32
	_ = v3818
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3837 int32
	_ = v3837
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3853 int32
	_ = v3853
	var v3857 int32
	_ = v3857
	var v3871 int32
	_ = v3871
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3900 int32
	_ = v3900
	var v3905 int32
	_ = v3905
	var v3910 int32
	_ = v3910
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3916 int32
	_ = v3916
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3926 int32
	_ = v3926
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3935 int32
	_ = v3935
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3949 int32
	_ = v3949
	var v3951 int32
	_ = v3951
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3967 int32
	_ = v3967
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3974 int32
	_ = v3974
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3982 int32
	_ = v3982
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3991 int32
	_ = v3991
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3996 int32
	_ = v3996
	var v3997 int32
	_ = v3997
	var v3999 int32
	_ = v3999
	var v4001 int32
	_ = v4001
	var v4004 int32
	_ = v4004
	var v4005 int32
	_ = v4005
	var v4008 int32
	_ = v4008
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4018 int32
	_ = v4018
	var v4022 int32
	_ = v4022
	var v4024 int32
	_ = v4024
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4043 int32
	_ = v4043
	var v4063 int32
	_ = v4063
	var v4070 int32
	_ = v4070
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4082 int32
	_ = v4082
	var v4083 int32
	_ = v4083
	var v4084 int64
	_ = v4084
	var v4092 int32
	_ = v4092
	var v4097 int32
	_ = v4097
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4108 int32
	_ = v4108
	var v4110 int32
	_ = v4110
	var v4112 int32
	_ = v4112
	var v4115 int32
	_ = v4115
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4120 int32
	_ = v4120
	var v4128 int32
	_ = v4128
	var v4130 int32
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4138 int32
	_ = v4138
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4141 int32
	_ = v4141
	var v4147 int32
	_ = v4147
	var v4148 int32
	_ = v4148
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4155 int32
	_ = v4155
	var v4156 int64
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4163 int32
	_ = v4163
	var v4164 int32
	_ = v4164
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4177 int32
	_ = v4177
	var v4178 int32
	_ = v4178
	var v4183 int32
	_ = v4183
	var v4187 int32
	_ = v4187
	var v4192 int32
	_ = v4192
	var v4196 int32
	_ = v4196
	var v4199 int32
	_ = v4199
	var v4203 int32
	_ = v4203
	var v4208 int32
	_ = v4208
	var v4219 int32
	_ = v4219
	var v4224 int32
	_ = v4224
	var v4227 int32
	_ = v4227
	var v4230 int32
	_ = v4230
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4240 int64
	_ = v4240
	var v4241 int64
	_ = v4241
	var v4252 int32
	_ = v4252
	var v4256 int32
	_ = v4256
	var v4258 int32
	_ = v4258
	var v4259 int32
	_ = v4259
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4274 int32
	_ = v4274
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4289 int32
	_ = v4289
	var v4292 int32
	_ = v4292
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4300 int32
	_ = v4300
	var v4305 int32
	_ = v4305
	var v4306 int32
	_ = v4306
	var v4308 int32
	_ = v4308
	var v4312 int32
	_ = v4312
	var v4316 int32
	_ = v4316
	var v4318 int32
	_ = v4318
	var v4322 int32
	_ = v4322
	var v4325 int32
	_ = v4325
	var v4327 int32
	_ = v4327
	var v4328 int32
	_ = v4328
	var v4329 int32
	_ = v4329
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4340 int32
	_ = v4340
	var v4349 int32
	_ = v4349
	var v4351 int32
	_ = v4351
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4360 int32
	_ = v4360
	var v4367 int32
	_ = v4367
	var v4376 int32
	_ = v4376
	var v4378 int32
	_ = v4378
	var v4381 int32
	_ = v4381
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4394 int32
	_ = v4394
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4405 int32
	_ = v4405
	var v4406 int32
	_ = v4406
	var v4412 int32
	_ = v4412
	var v4418 int32
	_ = v4418
	var v4427 int32
	_ = v4427
	var v4428 int32
	_ = v4428
	var v4434 int32
	_ = v4434
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4441 int32
	_ = v4441
	var v4444 int32
	_ = v4444
	var v4449 int32
	_ = v4449
	var v4451 int32
	_ = v4451
	var v4452 int32
	_ = v4452
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4460 int32
	_ = v4460
	var v4468 int32
	_ = v4468
	var v4485 int32
	_ = v4485
	var v4488 int32
	_ = v4488
	var v4491 int32
	_ = v4491
	var v4495 int32
	_ = v4495
	var v4503 int32
	_ = v4503
	var v4521 int32
	_ = v4521
	var v4524 int32
	_ = v4524
	var v4526 int32
	_ = v4526
	var v4528 int32
	_ = v4528
	var v4531 int32
	_ = v4531
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4563 int32
	_ = v4563
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4573 int32
	_ = v4573
	var v4577 int32
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4583 int32
	_ = v4583
	var v4585 int32
	_ = v4585
	var v4592 int32
	_ = v4592
	var v4594 int32
	_ = v4594
	var v4600 int32
	_ = v4600
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4608 int32
	_ = v4608
	var v4611 int32
	_ = v4611
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4614 int32
	_ = v4614
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4619 int32
	_ = v4619
	var v4622 int32
	_ = v4622
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4628 int32
	_ = v4628
	var v4630 int32
	_ = v4630
	var v4632 int32
	_ = v4632
	var v4640 int32
	_ = v4640
	var v4642 int32
	_ = v4642
	var v4648 int32
	_ = v4648
	var v4653 int32
	_ = v4653
	var v4657 int32
	_ = v4657
	var v4660 int32
	_ = v4660
	var v4664 int32
	_ = v4664
	var v4669 int32
	_ = v4669
	var v4673 int32
	_ = v4673
	var v4676 int32
	_ = v4676
	var v4683 int32
	_ = v4683
	var v4688 int32
	_ = v4688
	var v4692 int32
	_ = v4692
	var v4695 int32
	_ = v4695
	var v4702 int32
	_ = v4702
	var v4707 int32
	_ = v4707
	var v4711 int32
	_ = v4711
	var v4714 int32
	_ = v4714
	var v4718 int32
	_ = v4718
	var v4724 int32
	_ = v4724
	var v4725 int32
	_ = v4725
	var v4730 int32
	_ = v4730
	var v4734 int32
	_ = v4734
	var v4737 int32
	_ = v4737
	var v4743 int32
	_ = v4743
	var v4748 int32
	_ = v4748
	var v4749 int32
	_ = v4749
	var v4757 int32
	_ = v4757
	var v4759 int32
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4768 int32
	_ = v4768
	var v4769 int32
	_ = v4769
	var v4772 int32
	_ = v4772
	var v4775 int32
	_ = v4775
	var v4776 int32
	_ = v4776
	var v4777 int32
	_ = v4777
	var v4778 int32
	_ = v4778
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4781 int32
	_ = v4781
	var v4783 int32
	_ = v4783
	var v4786 int32
	_ = v4786
	var v4787 int32
	_ = v4787
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4793 int32
	_ = v4793
	var v4799 int32
	_ = v4799
	var v4800 int32
	_ = v4800
	var v4801 int32
	_ = v4801
	var v4807 int32
	_ = v4807
	var v4808 int32
	_ = v4808
	var v4814 int32
	_ = v4814
	var v4817 int32
	_ = v4817
	var v4823 int32
	_ = v4823
	var v4828 int32
	_ = v4828
	var v4830 int32
	_ = v4830
	var v4832 int32
	_ = v4832
	var v4839 int32
	_ = v4839
	var v4843 int32
	_ = v4843
	var v4855 int32
	_ = v4855
	var v4856 int32
	_ = v4856
	var v4857 int32
	_ = v4857
	var v4859 int32
	_ = v4859
	var v4863 int32
	_ = v4863
	var v4864 int32
	_ = v4864
	var v4866 int32
	_ = v4866
	var v4867 int32
	_ = v4867
	var v4868 int32
	_ = v4868
	var v4870 int32
	_ = v4870
	var v4876 int32
	_ = v4876
	var v4877 int32
	_ = v4877
	var v4878 int32
	_ = v4878
	var v4879 int32
	_ = v4879
	var v4882 int32
	_ = v4882
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4891 int32
	_ = v4891
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4902 int32
	_ = v4902
	var v4903 int32
	_ = v4903
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4911 int32
	_ = v4911
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4918 int32
	_ = v4918
	var v4919 int32
	_ = v4919
	var v4920 int32
	_ = v4920
	var v4923 int32
	_ = v4923
	var v4924 int32
	_ = v4924
	var v4925 int32
	_ = v4925
	var v4927 int32
	_ = v4927
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4934 int32
	_ = v4934
	var v4936 int32
	_ = v4936
	var v4938 int32
	_ = v4938
	var v4939 int32
	_ = v4939
	var v4942 int32
	_ = v4942
	var v4949 int32
	_ = v4949
	var v4952 int32
	_ = v4952
	var v4956 int32
	_ = v4956
	var v4959 int32
	_ = v4959
	var v4968 int32
	_ = v4968
	var v4972 int32
	_ = v4972
	var v4974 int32
	_ = v4974
	var v4975 int32
	_ = v4975
	var v4979 int32
	_ = v4979
	var v4980 int32
	_ = v4980
	var v4982 int32
	_ = v4982
	var v4986 int32
	_ = v4986
	var v4988 int32
	_ = v4988
	var v4990 int32
	_ = v4990
	var v4993 int32
	_ = v4993
	var v4998 int32
	_ = v4998
	var v4999 int32
	_ = v4999
	var v5000 int32
	_ = v5000
	var v5002 int32
	_ = v5002
	var v5003 int32
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5007 int32
	_ = v5007
	var v5010 int32
	_ = v5010
	var v5015 int32
	_ = v5015
	var v5016 int32
	_ = v5016
	var v5017 int32
	_ = v5017
	var v5024 int32
	_ = v5024
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5029 int32
	_ = v5029
	var v5038 int32
	_ = v5038
	var v5046 int32
	_ = v5046
	var v5049 int32
	_ = v5049
	var v5051 int32
	_ = v5051
	var v5053 int32
	_ = v5053
	var v5054 int32
	_ = v5054
	var v5055 int32
	_ = v5055
	var v5057 int32
	_ = v5057
	var v5061 int32
	_ = v5061
	var v5065 int32
	_ = v5065
	var v5069 int32
	_ = v5069
	var v5075 int32
	_ = v5075
	var v5080 int32
	_ = v5080
	var v5082 int32
	_ = v5082
	var v5084 int32
	_ = v5084
	var v5086 int32
	_ = v5086
	var v5088 int32
	_ = v5088
	var v5090 int32
	_ = v5090
	var v5093 int64
	_ = v5093
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5099 int32
	_ = v5099
	var v5100 int32
	_ = v5100
	var v5101 int32
	_ = v5101
	var v5102 int32
	_ = v5102
	var v5104 int32
	_ = v5104
	var v5107 int32
	_ = v5107
	var v5110 int32
	_ = v5110
	var v5113 int32
	_ = v5113
	var v5114 int32
	_ = v5114
	var v5117 int32
	_ = v5117
	var v5118 int32
	_ = v5118
	var v5121 int32
	_ = v5121
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5132 int32
	_ = v5132
	var v5136 int32
	_ = v5136
	var v5139 int32
	_ = v5139
	var v5144 int32
	_ = v5144
	var v5151 int32
	_ = v5151
	var v5153 int32
	_ = v5153
	var v5155 int32
	_ = v5155
	var v5156 int32
	_ = v5156
	var v5158 int32
	_ = v5158
	var v5161 int32
	_ = v5161
	var v5167 int32
	_ = v5167
	var v5168 int32
	_ = v5168
	var v5170 int32
	_ = v5170
	var v5172 int32
	_ = v5172
	var v5176 int32
	_ = v5176
	var v5177 int32
	_ = v5177
	var v5178 int32
	_ = v5178
	var v5184 int32
	_ = v5184
	var v5188 int32
	_ = v5188
	var v5191 int32
	_ = v5191
	var v5214 int32
	_ = v5214
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5221 int32
	_ = v5221
	var v5224 int32
	_ = v5224
	var v5228 int32
	_ = v5228
	var v5230 int32
	_ = v5230
	var v5237 int32
	_ = v5237
	var v5258 int32
	_ = v5258
	var v5262 int32
	_ = v5262
	var v5263 int32
	_ = v5263
	var v5290 int32
	_ = v5290
	var v5291 int32
	_ = v5291
	var v5293 int32
	_ = v5293
	var v5305 int32
	_ = v5305
	var v5309 int32
	_ = v5309
	var v5314 int32
	_ = v5314
	var v5317 int32
	_ = v5317
	var v5327 int32
	_ = v5327
	var v5331 int32
	_ = v5331
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5339 int32
	_ = v5339
	var v5342 int64
	_ = v5342
	var v5343 int32
	_ = v5343
	var v5345 int32
	_ = v5345
	var v5346 int32
	_ = v5346
	var v5349 int64
	_ = v5349
	var v5350 int32
	_ = v5350
	var v5352 int32
	_ = v5352
	var v5353 int32
	_ = v5353
	var v5355 int32
	_ = v5355
	var v5356 int32
	_ = v5356
	var v5360 int32
	_ = v5360
	var v5361 int32
	_ = v5361
	var v5364 int32
	_ = v5364
	var v5366 int32
	_ = v5366
	var v5370 int64
	_ = v5370
	var v5371 int32
	_ = v5371
	var v5372 int32
	_ = v5372
	var v5373 int32
	_ = v5373
	var v5374 int32
	_ = v5374
	var v5375 int32
	_ = v5375
	var v5376 int32
	_ = v5376
	var v5381 int32
	_ = v5381
	var v5382 int32
	_ = v5382
	var v5385 int32
	_ = v5385
	var v5386 int32
	_ = v5386
	var v5387 int32
	_ = v5387
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5400 int32
	_ = v5400
	var v5405 int32
	_ = v5405
	var v5408 int32
	_ = v5408
	var v5409 int32
	_ = v5409
	var v5410 int32
	_ = v5410
	var v5411 int32
	_ = v5411
	var v5412 int32
	_ = v5412
	var v5415 int32
	_ = v5415
	var v5424 int32
	_ = v5424
	var v5426 int32
	_ = v5426
	var v5430 int32
	_ = v5430
	var v5435 int32
	_ = v5435
	var v5440 int64
	_ = v5440
	var v5441 int32
	_ = v5441
	var v5442 int32
	_ = v5442
	var v5444 int32
	_ = v5444
	var v5445 int32
	_ = v5445
	var v5447 int32
	_ = v5447
	var v5452 int64
	_ = v5452
	var v5453 int32
	_ = v5453
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5457 int32
	_ = v5457
	var v5459 int32
	_ = v5459
	var v5460 int32
	_ = v5460
	var v5462 int32
	_ = v5462
	var v5463 int32
	_ = v5463
	var v5468 int32
	_ = v5468
	var v5469 int32
	_ = v5469
	var v5479 int32
	_ = v5479
	var v5483 int32
	_ = v5483
	var v5486 int32
	_ = v5486
	var v5489 int32
	_ = v5489
	var v5490 int32
	_ = v5490
	var v5493 int32
	_ = v5493
	var v5494 int32
	_ = v5494
	var v5497 int32
	_ = v5497
	var v5504 int32
	_ = v5504
	var v5505 int32
	_ = v5505
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
	var v5516 int32
	_ = v5516
	var v5522 int32
	_ = v5522
	var v5528 int32
	_ = v5528
	var v5529 int32
	_ = v5529
	var v5530 int32
	_ = v5530
	var v5531 int32
	_ = v5531
	var v5537 int32
	_ = v5537
	var v5540 int32
	_ = v5540
	var v5543 int32
	_ = v5543
	var v5548 int32
	_ = v5548
	var v5550 int32
	_ = v5550
	var v5552 int32
	_ = v5552
	var v5554 int32
	_ = v5554
	var v5556 int32
	_ = v5556
	var v5583 int32
	_ = v5583
	var v5585 int32
	_ = v5585
	var v5587 int32
	_ = v5587
	var v5589 int32
	_ = v5589
	var v5591 int32
	_ = v5591
	var v5596 int32
	_ = v5596
	var v5597 int32
	_ = v5597
	var v5599 int32
	_ = v5599
	var v5600 int32
	_ = v5600
	var v5601 int32
	_ = v5601
	var v5602 int32
	_ = v5602
	var v5605 int32
	_ = v5605
	var v5609 int32
	_ = v5609
	var v5613 int32
	_ = v5613
	var v5614 int32
	_ = v5614
	var v5618 int32
	_ = v5618
	var v5620 int32
	_ = v5620
	var v5623 int32
	_ = v5623
	var v5627 int32
	_ = v5627
	var v5633 int32
	_ = v5633
	var v5635 int32
	_ = v5635
	var v5638 int32
	_ = v5638
	var v5641 int32
	_ = v5641
	var v5642 int32
	_ = v5642
	var v5645 int32
	_ = v5645
	var v5647 int32
	_ = v5647
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5658 int32
	_ = v5658
	var v5660 int32
	_ = v5660
	var v5663 int32
	_ = v5663
	var v5664 int32
	_ = v5664
	var v5671 int32
	_ = v5671
	var v5675 int32
	_ = v5675
	var v5679 int32
	_ = v5679
	var v5683 int32
	_ = v5683
	var v5685 int32
	_ = v5685
	var v5687 int64
	_ = v5687
	var v5695 int32
	_ = v5695
	var v5700 int32
	_ = v5700
	var v5705 int32
	_ = v5705
	var v5710 int32
	_ = v5710
	var v5712 int32
	_ = v5712
	var v5715 int32
	_ = v5715
	var v5723 int32
	_ = v5723
	var v5726 int32
	_ = v5726
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5734 int32
	_ = v5734
	var v5738 int32
	_ = v5738
	var v5740 int32
	_ = v5740
	var v5744 int32
	_ = v5744
	var v5773 int32
	_ = v5773
	var v5775 int32
	_ = v5775
	var v5801 int32
	_ = v5801
	var v5802 int32
	_ = v5802
	var v5804 int32
	_ = v5804
	var v5807 int32
	_ = v5807
	var v5811 int32
	_ = v5811
	var v5814 int32
	_ = v5814
	var v5815 int32
	_ = v5815
	var v5821 int32
	_ = v5821
	var v5843 int32
	_ = v5843
	var v5847 int32
	_ = v5847
	var v5848 int32
	_ = v5848
	var v5849 int32
	_ = v5849
	var v5850 int32
	_ = v5850
	var v5855 int32
	_ = v5855
	var v5856 int32
	_ = v5856
	var v5859 int32
	_ = v5859
	var v5865 int32
	_ = v5865
	var v5866 int32
	_ = v5866
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5875 int32
	_ = v5875
	var v5876 int32
	_ = v5876
	var v5878 int32
	_ = v5878
	var v5879 int32
	_ = v5879
	var v5881 int32
	_ = v5881
	var v5883 int32
	_ = v5883
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5914 int32
	_ = v5914
	var v5941 int32
	_ = v5941
	var v5943 int32
	_ = v5943
	var v5953 int32
	_ = v5953
	var v5957 int32
	_ = v5957
	var v5962 int32
	_ = v5962
	var v5975 int32
	_ = v5975
	var v5994 int32
	_ = v5994
	var v5997 int32
	_ = v5997
	var v6005 int32
	_ = v6005
	var v6009 int32
	_ = v6009
	var v6014 int32
	_ = v6014
	var v6017 int32
	_ = v6017
	var v6025 int32
	_ = v6025
	var v6028 int32
	_ = v6028
	var v6029 int32
	_ = v6029
	var v6034 int32
	_ = v6034
	var v6038 int32
	_ = v6038
	var v6040 int32
	_ = v6040
	var v6046 int32
	_ = v6046
	var v6051 int32
	_ = v6051
	var v6055 int32
	_ = v6055
	var v6058 int32
	_ = v6058
	var v6066 int32
	_ = v6066
	var v6069 int32
	_ = v6069
	var v6074 int32
	_ = v6074
	var v6075 int32
	_ = v6075
	var v6080 int32
	_ = v6080
	var v6084 int32
	_ = v6084
	var v6087 int32
	_ = v6087
	var v6095 int32
	_ = v6095
	var v6100 int32
	_ = v6100
	var v6104 int32
	_ = v6104
	var v6107 int32
	_ = v6107
	var v6115 int32
	_ = v6115
	var v6118 int32
	_ = v6118
	var v6119 int32
	_ = v6119
	var v6124 int32
	_ = v6124
	var v6128 int32
	_ = v6128
	var v6131 int32
	_ = v6131
	var v6139 int32
	_ = v6139
	var v6144 int32
	_ = v6144
	var v6148 int32
	_ = v6148
	var v6152 int32
	_ = v6152
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
	var v6162 int32
	_ = v6162
	var v6167 int32
	_ = v6167
	var v6171 int32
	_ = v6171
	var v6175 int32
	_ = v6175
	var v6180 int32
	_ = v6180
	var v6181 int32
	_ = v6181
	var v6185 int32
	_ = v6185
	var v6190 int32
	_ = v6190
	var v6195 int32
	_ = v6195
	var v6198 int32
	_ = v6198
	var v6204 int32
	_ = v6204
	var v6207 int32
	_ = v6207
	var v6208 int32
	_ = v6208
	var v6213 int32
	_ = v6213
	v7 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(528)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+444)) = v7
	v36 = F_errstart(m, int32(12), v7)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(_a_F_InitPostgres_0), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[1]))
	F_ProcArrayAdd(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(740), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	F_on_shmem_exit(m, int32(1228), int64(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_pgstat_beinit(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v107 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	F_pgstat_bestart_initial(m)
	mBase = m.M
	F_SharedInvalBackendInit(m, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_SharedInvalBackendInit(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L25
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[3]))
	F_ProcSignalInit(m, int32(_a_F_InitPostgres_2), v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_RegisterTimeout(m, int32(1), int32(1839))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_RegisterTimeout(m, int32(3), int32(1840))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_RegisterTimeout(m, int32(2), int32(1841))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_RegisterTimeout(m, int32(7), int32(1842))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_RegisterTimeout(m, int32(8), int32(1843))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_RegisterTimeout(m, int32(9), int32(1844))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_RegisterTimeout(m, int32(11), int32(1845))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_RegisterTimeout(m, int32(10), int32(1846))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L11
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[3]))
	F_ProcSignalInit(m, int32(_a_F_InitPostgres_2), v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L11
L27:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_InitializeProcessXLogLogicalInfo(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L35
	}
L30:
	;
	F_StartupXLOG(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4])) = int32(0)
	F_before_shmem_exit(m, int32(988), int64(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_before_shmem_exit(m, int32(1847), int64(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v130 = m.G0
	v132 = v130 - int32(48)
	m.G0 = v132
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5]))
	if v135 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v132)+8)) = int64(34359738372)
	v146 = F_hash_create(m, int32(_a_F_InitPostgres_3), int64(400), v132, int32(40))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[6])) = v146
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5]))
	v152 = F_MemoryContextAlloc(m, v150, int32(32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[7])) = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[8])) = v152
	v160 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[9])) = v160
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[10])) = v160
	v166 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[11])) = v166
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[12])) = v166
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[13])) = v166
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[14])) = v166
	m.G0 = v132 + int32(48)
	v181 = m.G0
	v183 = v181 - int32(16)
	m.G0 = v183
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v166
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v166
	v198 = v166
	goto L44
L42:
	;
	F_CacheRegisterRelcacheCallback(m, int32(1798))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L93
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L90
	}
L44:
	;
	v217 = v198 << (uint(int32(5)) % 32)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_InitPostgres[17])))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_InitPostgres[18])))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_InitPostgres[19])))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_InitPostgres[20])))
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5]))
	if v223 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	F_pg_qsort(m, int32(_a_F_InitPostgres_4), v453, int32(4), int32(1810))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L70
	}
L46:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v233 = F_AllocSetContextCreateInternal(m, v228, int32(_a_F_InitPostgres_5), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	v236 = v223
	goto L48
L48:
	;
	v237 = int32(_a_F_InitPostgres_8)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v236
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23]))
	if v242 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5])) = v233
	v236 = v233
	goto L48
L50:
	;
	v247 = F_palloc(m, int32(8))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v254 = v217 + int32(_a_F_InitPostgres_9)
	v258 = F_palloc_aligned(m, int32(328), int32(128), int32(4))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23])) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = int64(0)
	goto L52
L54:
	;
	v262 = F_palloc0(m, v221<<(uint(int32(3))%32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = v262
	v265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+96)) = uint8(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+92)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v258)+88)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v258)+84)) = int32(_a_F_InitPostgres_10)
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v198
	v272 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v258)+68)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = v265
	*(*int64)(unsafe.Add(mBase, uint32(v258)+76)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v258)+64)) = v220
	if v220 <= v265 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+100)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v258 + int32(100)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v198<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[24]))) = v258
	if v258 == int32(0) {
		goto L43
	} else {
		goto L68
	}
L57:
	;
	v283 = v220 & int32(3)
	v285 = v258 + int32(48)
	v286 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v220) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v298 = v286
	v305 = int32(0)
	goto L61
L59:
	;
	v354 = v286
	goto L60
L60:
	;
	v380 = v354
	v385 = int32(0)
	goto L65
L61:
	;
	v318 = v298 << (uint(int32(2)) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v254+v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v318))) = v321
	v323 = int32(4)
	v324 = v318 | v323
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v254+v324)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v324))) = v327
	v330 = v318 | int32(8)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v254+v330)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v330))) = v333
	v336 = v318 | int32(12)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v254+v336)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v336))) = v339
	v342 = v298 + v323
	v344 = v305 + v323
	if v344 != v220&int32(2147483644) {
		v298 = v342
		v305 = v344
		goto L61
	} else {
		goto L63
	}
L62:
	;
	if v283 == int32(0) {
		goto L56
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v354 = v342
	goto L60
L65:
	;
	v400 = v380 << (uint(int32(2)) % 32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v254+v400)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v400))) = v403
	v405 = int32(1)
	v408 = v385 + v405
	if v408 != v283 {
		v380 = v380 + v405
		v385 = v408
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L56
L67:
	;
	goto L66
L68:
	;
	v449 = int32(_a_F_InitPostgres_11)
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15]))
	v452 = int32(1)
	v453 = v451 + v452
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v453
	v455 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v451<<(uint(v455)%32))+uint32(_c_F_InitPostgres[25]))) = v218
	v460 = int32(_a_F_InitPostgres_12)
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	v463 = v461 << (uint(v455) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_c_F_InitPostgres[26]))) = v218
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v461 + v455
	*(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_c_F_InitPostgres[27]))) = v219
	v475 = v198 + v452
	if v475 != int32(85) {
		v198 = v475
		goto L44
	} else {
		goto L69
	}
L69:
	;
	goto L45
L70:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15]))
	if base.Ui32(int32(2)) <= base.Ui32(v485) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v495 = int32(0)
	v496 = int32(1)
	goto L74
L72:
	;
	v543 = v485
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v543
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	F_pg_qsort(m, int32(_a_F_InitPostgres_13), v564, int32(4), int32(1810))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v514 = int32(2)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v496<<(uint(v514)%32))+uint32(_c_F_InitPostgres[25])))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v495<<(uint(v514)%32))+uint32(_c_F_InitPostgres[25])))
	if v516 == v519 {
		v527 = v495
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v543 = v527 + int32(1)
	goto L73
L76:
	;
	v530 = v496 + int32(1)
	if v530 != v485 {
		v495 = v527
		v496 = v530
		goto L74
	} else {
		goto L79
	}
L77:
	;
	v522 = v495 + int32(1)
	if v522 == v496 {
		v527 = v496
		goto L76
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[25]))) = v516
	v527 = v522
	goto L76
L79:
	;
	goto L75
L80:
	;
	v569 = int32(_a_F_InitPostgres_12)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	if base.Ui32(int32(2)) <= base.Ui32(v571) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v581 = int32(0)
	v582 = int32(1)
	goto L84
L82:
	;
	v645 = v571
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v645
	m.G0 = v183 + int32(16)
	goto L42
L84:
	;
	v600 = int32(2)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v582<<(uint(v600)%32))+uint32(_c_F_InitPostgres[26])))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v581<<(uint(v600)%32))+uint32(_c_F_InitPostgres[26])))
	if v602 == v605 {
		v613 = v581
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v645 = v613 + int32(1)
	goto L83
L86:
	;
	v616 = v582 + int32(1)
	if v616 != v571 {
		v581 = v613
		v582 = v616
		goto L84
	} else {
		goto L89
	}
L87:
	;
	v608 = v581 + int32(1)
	if v608 == v582 {
		v613 = v582
		goto L86
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[26]))) = v602
	v613 = v608
	goto L86
L89:
	;
	goto L85
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v218
	F_errmsg_internal(m, int32(_a_F_InitPostgres_14), v183)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_15), int32(137), int32(_a_F_InitPostgres_16))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
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
	F_CacheRegisterSyscacheCallback(m, int32(47), int32(1799), int64(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_CacheRegisterSyscacheCallback(m, int32(82), int32(1799), int64(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_CacheRegisterSyscacheCallback(m, int32(38), int32(1800), int64(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_CacheRegisterSyscacheCallback(m, int32(40), int32(1800), int64(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(1800), int64(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_CacheRegisterSyscacheCallback(m, int32(32), int32(1800), int64(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_CacheRegisterSyscacheCallback(m, int32(30), int32(1800), int64(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_CacheRegisterSyscacheCallback(m, int32(9), int32(1801), int64(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_CacheRegisterSyscacheCallback(m, int32(11), int32(1801), int64(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_CacheRegisterSyscacheCallback(m, int32(21), int32(1801), int64(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v717 = m.G0
	v719 = v717 - int32(48)
	m.G0 = v719
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v728 = F_AllocSetContextCreateInternal(m, v723, int32(_a_F_InitPostgres_17), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[28])) = v728
	*(*int64)(unsafe.Add(mBase, uint32(v719)+8)) = int64(292057776192)
	v737 = F_hash_create(m, int32(_a_F_InitPostgres_18), int64(16), v719, int32(24))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[29])) = v737
	m.G0 = v719 + int32(48)
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v744 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_read_relmap_file(m, int32(_a_F_InitPostgres_19), int32(_a_F_InitPostgres_20), int32(0), int32(22))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v752 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v753 = int32(_a_F_InitPostgres_8)
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v757
	v760 = F_load_relcache_init_file(m, int32(1))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_before_shmem_exit(m, int32(1848), int64(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L123
	}
L113:
	;
	if v760 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_21), int32(1248), int32(1), int32(18), int32(_a_F_InitPostgres_22))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v754
	goto L112
L117:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_23), int32(2842), int32(1), int32(12), int32(_a_F_InitPostgres_24))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_25), int32(2843), int32(1), int32(7), int32(_a_F_InitPostgres_26))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_27), int32(4066), int32(1), int32(4), int32(_a_F_InitPostgres_28))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_29), int32(_a_F_InitPostgres_30), int32(1), int32(23), int32(_a_F_InitPostgres_31))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_32), int32(2173), int32(1), int32(3), int32(_a_F_InitPostgres_33))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L116
L123:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	if v814 == int32(3) {
		goto L135
	} else {
		goto L136
	}
L124:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6195 = m.ExcPending
	if v6195 != 0 {
		goto L1
	} else {
		goto L1422
	}
L125:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6171 = m.ExcPending
	if v6171 != 0 {
		goto L1
	} else {
		goto L1417
	}
L126:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6148 = m.ExcPending
	if v6148 != 0 {
		goto L1
	} else {
		goto L1412
	}
L127:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6128 = m.ExcPending
	if v6128 != 0 {
		goto L1
	} else {
		goto L1408
	}
L128:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6104 = m.ExcPending
	if v6104 != 0 {
		goto L1
	} else {
		goto L1403
	}
L129:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6084 = m.ExcPending
	if v6084 != 0 {
		goto L1
	} else {
		goto L1399
	}
L130:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6055 = m.ExcPending
	if v6055 != 0 {
		goto L1
	} else {
		goto L1394
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6038 = m.ExcPending
	if v6038 != 0 {
		goto L1
	} else {
		goto L1391
	}
L132:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v6017 = m.ExcPending
	if v6017 != 0 {
		goto L1
	} else {
		goto L1387
	}
L133:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v5994 = m.ExcPending
	if v5994 != 0 {
		goto L1
	} else {
		goto L1382
	}
L134:
	;
	m.G0 = v5975 + int32(528)
	return
L135:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	if v31 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L138:
	;
	v5975 = v28
	goto L134
L139:
	;
	v5801 = int32(0)
	v5802 = m.G0
	v5804 = v5802 - int32(32)
	m.G0 = v5804
	v5807 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[31])))
	if v5807 == v5801 {
		goto L1352
	} else {
		goto L1353
	}
L140:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v5773 = m.ExcPending
	if v5773 != 0 {
		goto L1
	} else {
		goto L1349
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32])) = v5038
	v5046 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v5046)+20)) = v5038
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v5049 = m.ExcPending
	if v5049 != 0 {
		goto L1
	} else {
		goto L1181
	}
L142:
	;
	F_LockSharedObject(m, int32(1262), v4749, int32(3))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L1094
	}
L143:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4734 = m.ExcPending
	if v4734 != 0 {
		goto L1
	} else {
		goto L1090
	}
L144:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4711 = m.ExcPending
	if v4711 != 0 {
		goto L1
	} else {
		goto L1085
	}
L145:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L1
	} else {
		goto L1081
	}
L146:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4673 = m.ExcPending
	if v4673 != 0 {
		goto L1
	} else {
		goto L1077
	}
L147:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4657 = m.ExcPending
	if v4657 != 0 {
		goto L1
	} else {
		goto L1073
	}
L148:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4640 = m.ExcPending
	if v4640 != 0 {
		goto L1
	} else {
		goto L1070
	}
L149:
	;
	v4381 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v4381 != 0 {
		goto L1013
	} else {
		goto L1014
	}
L150:
	;
	v4353 = F_superuser(m)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L1012
	}
L151:
	;
	v936 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[34])) = uint8(v936)
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[35]))
	if v941 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L152:
	;
	F_InitializeSessionUserId(m, l2, l3, int32(base.Ui32(l4&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L193
	}
L153:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L192
	}
L154:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[36]))
	if v822 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L159
	}
L156:
	;
	v826 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[37])) = v826
	goto L158
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[38])) = int32(1)
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	switch v834 - int32(4) {
	case 0, 3:
		goto L153
	default:
		goto L160
	}
L160:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v838 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v841 = int32(_a_F_InitPostgres_34)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v847 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[39])))
	if base.B2i32(v844 == int32(0))|base.B2i32(v844 != v847) != 0 {
		v865 = v844
		v866 = v847
		goto L166
	} else {
		goto L167
	}
L162:
	;
	goto L163
L163:
	;
	if v834 != int32(5) {
		goto L151
	} else {
		goto L190
	}
L164:
	;
	v882 = F_table_open(m, int32(1260), int32(1))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L178
	}
L165:
	;
	if v865-v866 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L166:
	;
	goto L165
L167:
	;
	v850 = l2
	v851 = v841
	goto L168
L168:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851)+1)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+1)))
	if v855 == int32(0) {
		v865 = v855
		v866 = v854
		goto L166
	} else {
		goto L170
	}
L169:
	;
	v865 = v855
	v866 = v854
	goto L166
L170:
	;
	v858 = int32(1)
	if v855 == v854 {
		v850 = v850 + v858
		v851 = v851 + v858
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v873 = int32(0)
	F_InitializeSessionUserId(m, l2, v873, v873)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v879 = int32(1)
	goto L164
L176:
	;
	v877 = F_superuser(m)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v879 = v877
	goto L164
L178:
	;
	v884 = int32(0)
	v886 = F_table_beginscan_catalog(m, v882, v884, v884)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v888 = F_heap_getnext(m, v886)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)+188))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	m.T0[v892].(func(*base.Module, int32))(m, v886)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_relation_close(m, v882, int32(1))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	if v888 != 0 {
		v4355 = l0
		v4356 = l1
		v4358 = v879
		v4359 = l4
		v4360 = l5
		v4367 = v28
		v4376 = v31
		v4378 = v7
		goto L149
	} else {
		goto L183
	}
L183:
	;
	v900 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	if v900 == int32(0) {
		v4355 = l0
		v4356 = l1
		v4358 = v879
		v4359 = l4
		v4360 = l5
		v4367 = v28
		v4376 = v31
		v4378 = v7
		goto L149
	} else {
		goto L185
	}
L185:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_35), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+432)) = l2
	F_errhint(m, int32(_a_F_InitPostgres_36), v28+int32(432))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(917), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v4355 = l0
	v4356 = l1
	v4358 = v879
	v4359 = l4
	v4360 = l5
	v4367 = v28
	v4376 = v31
	v4378 = v7
	goto L149
L190:
	;
	if l2|l3 != 0 {
		goto L152
	} else {
		goto L191
	}
L191:
	;
	goto L153
L192:
	;
	v4355 = l0
	v4356 = l1
	v4358 = int32(1)
	v4359 = l4
	v4360 = l5
	v4367 = v28
	v4376 = v31
	v4378 = v7
	goto L149
L193:
	;
	v4328 = l0
	v4329 = l1
	v4332 = l4
	v4333 = l5
	v4340 = v28
	v4349 = v31
	v4351 = v7
	goto L150
L194:
	;
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v951 = F_AllocSetContextCreateInternal(m, v946, int32(_a_F_InitPostgres_37), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v954 = F_load_hba(m)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L198
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[35])) = v951
	goto L196
L198:
	;
	if v954 == int32(0) {
		goto L148
	} else {
		goto L199
	}
L199:
	;
	v958 = F_load_ident(m)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v964 = m.G0
	v965 = int32(16)
	v966 = v964 - v965
	m.G0 = v966
	F_gettimeofday(m, v966)
	mBase = m.M
	v969 = *(*int64)(unsafe.Add(mBase, uint32(v966)))
	v970 = int64(*(*int32)(unsafe.Add(mBase, uint32(v966)+8)))
	m.G0 = v966 + v965
	goto L201
L201:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[40])) = v970 + v969*int64(1000000) - int64(946684800000000)
	v982 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	F_enable_timeout_after(m, int32(3), v982*int32(1000))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v987 = m.G0
	v989 = v987 - int32(2992)
	m.G0 = v989
	v991 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+444)) = v991
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+443)) = uint8(v991)
	v995 = m.G0
	v997 = v995 - int32(288)
	m.G0 = v997
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v1001 = F_get_role_oid(m, v999, int32(1))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[42]))
	if v1004 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+380)) = v2295
	m.G0 = v997 + int32(288)
	v2311 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v2311 != 0 {
		goto L520
	} else {
		goto L521
	}
L205:
	;
	v2277 = F_palloc0(m, int32(396))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L1
	} else {
		goto L519
	}
L206:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	if v1007 <= int32(0) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1011 = v939 + int32(144)
	v1034 = v7
	goto L208
L208:
	;
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011))))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+12))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1038+v1034<<(uint(int32(2))%32))))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+12))
	if v1043 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L209:
	;
	goto L205
L210:
	;
	v2248 = v1034 + int32(1)
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	if v2248 < v2249 {
		v1034 = v2248
		goto L208
	} else {
		goto L518
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+288)) = v1076
	*(*int32)(unsafe.Add(mBase, uint32(v939)+284)) = int32(-2)
	goto L210
L212:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+16))
	if v1891 == int32(0) {
		goto L210
	} else {
		goto L420
	}
L213:
	;
	if v1037 == int32(1) {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	if v1037 == int32(1) {
		goto L210
	} else {
		goto L217
	}
L216:
	;
	goto L210
L217:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+488)))
	if v1050 == int32(1) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+288))
	switch v1059 {
	case 0:
		goto L225
	case 1, 2:
		goto L224
	case 3:
		goto L212
	default:
		goto L210
	}
L219:
	;
	if base.Ui32(v1043-int32(3)) < base.Ui32(int32(2)) {
		goto L210
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	switch v1043 - int32(2) {
	case 0, 2:
		goto L210
	default:
		goto L218
	}
L222:
	;
	goto L218
L223:
	;
	v1799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1042)+24)))
	if v1037 != v1799 {
		goto L210
	} else {
		goto L409
	}
L224:
	;
	v1281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v997)+24)) = uint8(v1281)
	*(*int32)(unsafe.Add(mBase, uint32(v997)+20)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v997)+16)) = v1059
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44])) = v1281
	v1289 = v997 + int32(16)
	v1290 = m.G0
	v1292 = v1290 - int32(144)
	m.G0 = v1292
	v1296 = m.G0
	v1298 = v1296 - int32(272)
	m.G0 = v1298
	v1301 = v1298 + int32(8)
	goto L294
L225:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+292))
	if v1060 == int32(0) {
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v939)+284))
	if v1063 < int32(0) {
		goto L210
	} else {
		goto L227
	}
L227:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v939)+280))
	if v1066 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v939)+272))
	v1071 = v997 + int32(16)
	v1073 = int32(0)
	v1076 = F_pg_getnameinfo_all(m, v1011, v1069, v1071, int32(255), v1073, v1073, int32(8))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	v1081 = v1066
	goto L230
L230:
	;
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	if v1083 == int32(46) {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	if v1076 != 0 {
		goto L211
	} else {
		goto L232
	}
L232:
	;
	v1078 = F_pstrdup(m, v1071)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+280)) = v1078
	v1081 = v1078
	goto L230
L234:
	;
	v1086 = F_strlen(m, v1060)
	mBase = m.M
	v1087 = F_strlen(m, v1081)
	mBase = m.M
	if base.Ui32(v1087) < base.Ui32(v1086) {
		goto L210
	} else {
		goto L237
	}
L235:
	;
	v1093 = v1081
	goto L236
L236:
	;
	v1096 = v1060
	v1097 = v1093
	goto L239
L237:
	;
	v1093 = v1081 + (v1087 - v1086)
	goto L236
L238:
	;
	if v1134 != 0 {
		goto L210
	} else {
		goto L251
	}
L239:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096))))
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097))))
	if v1100 == v1101 {
		v1123 = v1100
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1134 = int32(0)
	goto L238
L241:
	;
	v1125 = int32(1)
	if v1123 != 0 {
		v1096 = v1096 + v1125
		v1097 = v1097 + v1125
		goto L239
	} else {
		goto L250
	}
L242:
	;
	if base.Ui32((v1100-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1111 = v1100 | int32(32)
	goto L245
L244:
	;
	v1111 = v1100
	goto L245
L245:
	;
	if base.Ui32((v1101-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1120 = v1101 | int32(32)
	goto L248
L247:
	;
	v1120 = v1101
	goto L248
L248:
	;
	if v1111 == v1120 {
		v1123 = v1111
		goto L241
	} else {
		goto L249
	}
L249:
	;
	v1134 = v1111 - v1120
	goto L238
L250:
	;
	goto L240
L251:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v939)+284))
	if v1135 == int32(1) {
		goto L212
	} else {
		goto L252
	}
L252:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v939)+280))
	v1139 = int32(0)
	v1143 = m.Env.Getaddrinfo(m, v1138, v1139, v1139, v997+int32(284))
	mBase = m.M
	if v1143 == v1139 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v997)+284))
	if v1146 != 0 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+288)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(v939)+284)) = int32(-2)
	goto L210
L256:
	;
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011))))
	v1156 = v1146
	goto L259
L257:
	;
	goto L258
L258:
	;
	v1265 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L286
	}
L259:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+20))
	v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175))))
	if v1176 != v1147 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+20))
	F_emscripten_builtin_free(m, v1235)
	mBase = m.M
	F_emscripten_builtin_free(m, v1146)
	mBase = m.M
	goto L285
L261:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+28))
	if v1234 != 0 {
		v1156 = v1234
		goto L259
	} else {
		goto L284
	}
L262:
	;
	switch v1147 - int32(2) {
	case 0:
		goto L264
	default:
		goto L261
	case 8:
		goto L265
	}
L263:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+20))
	F_emscripten_builtin_free(m, v1229)
	mBase = m.M
	F_emscripten_builtin_free(m, v1146)
	mBase = m.M
	goto L283
L264:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+4))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v939)+148))
	if v1226 != v1227 {
		goto L261
	} else {
		goto L282
	}
L265:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+8)))
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+152)))
	if v1178 != v1179 {
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+9)))
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+153)))
	if v1181 != v1182 {
		goto L261
	} else {
		goto L267
	}
L267:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+10)))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+154)))
	if v1184 != v1185 {
		goto L261
	} else {
		goto L268
	}
L268:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+11)))
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+155)))
	if v1187 != v1188 {
		goto L261
	} else {
		goto L269
	}
L269:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+12)))
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+156)))
	if v1190 != v1191 {
		goto L261
	} else {
		goto L270
	}
L270:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+13)))
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+157)))
	if v1193 != v1194 {
		goto L261
	} else {
		goto L271
	}
L271:
	;
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+14)))
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+158)))
	if v1196 != v1197 {
		goto L261
	} else {
		goto L272
	}
L272:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+15)))
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+159)))
	if v1199 != v1200 {
		goto L261
	} else {
		goto L273
	}
L273:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+16)))
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+160)))
	if v1202 != v1203 {
		goto L261
	} else {
		goto L274
	}
L274:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+17)))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+161)))
	if v1205 != v1206 {
		goto L261
	} else {
		goto L275
	}
L275:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+18)))
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+162)))
	if v1208 != v1209 {
		goto L261
	} else {
		goto L276
	}
L276:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+19)))
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+163)))
	if v1211 != v1212 {
		goto L261
	} else {
		goto L277
	}
L277:
	;
	v1214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+20)))
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+164)))
	if v1214 != v1215 {
		goto L261
	} else {
		goto L278
	}
L278:
	;
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+21)))
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+165)))
	if v1217 != v1218 {
		goto L261
	} else {
		goto L279
	}
L279:
	;
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+22)))
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+166)))
	if v1220 != v1221 {
		goto L261
	} else {
		goto L280
	}
L280:
	;
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+23)))
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939)+167)))
	if v1223 == v1224 {
		goto L263
	} else {
		goto L281
	}
L281:
	;
	goto L261
L282:
	;
	goto L263
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+284)) = int32(1)
	goto L212
L284:
	;
	goto L260
L285:
	;
	goto L258
L286:
	;
	if v1265 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v997))) = v1060
	F_errmsg_internal(m, int32(_a_F_InitPostgres_38), v997)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+284)) = int32(-1)
	goto L210
L290:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_39), int32(1154), int32(_a_F_InitPostgres_40))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1412 = int32(0)
	v1413 = F_socket(m, int32(16), int32(_a_F_InitPostgres_41), v1412)
	mBase = m.M
	if v1413 < v1412 {
		goto L304
	} else {
		goto L305
	}
L293:
	;
	goto L292
L294:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1301))) = uint8(v1281)
	v1310 = v1298 + int32(272)
	*(*uint8)(unsafe.Add(mBase, uint32(v1310-int32(1)))) = uint8(v1281)
	goto L295
L295:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1301)+2)) = uint8(v1281)
	*(*uint8)(unsafe.Add(mBase, uint32(v1301)+1)) = uint8(v1281)
	*(*uint8)(unsafe.Add(mBase, uint32(v1310-int32(3)))) = uint8(v1281)
	*(*uint8)(unsafe.Add(mBase, uint32(v1310-int32(2)))) = uint8(v1281)
	goto L296
L296:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1301)+3)) = uint8(v1281)
	*(*uint8)(unsafe.Add(mBase, uint32(v1310-int32(4)))) = uint8(v1281)
	goto L297
L297:
	;
	v1332 = int32(0)
	v1335 = (v1332 - v1301) & int32(3)
	v1336 = v1301 + v1335
	*(*int32)(unsafe.Add(mBase, uint32(v1336))) = v1332
	v1344 = (int32(264) - v1335) & int32(-4)
	v1345 = v1336 + v1344
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(4)))) = v1332
	if base.Ui32(v1344) < base.Ui32(int32(9)) {
		goto L293
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+8)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+4)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(8)))) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(12)))) = v1332
	if base.Ui32(v1344) < base.Ui32(int32(25)) {
		goto L293
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+24)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+20)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+16)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+12)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(16)))) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(20)))) = v1332
	v1371 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1345-v1371))) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1345-int32(28)))) = v1332
	v1380 = v1336&int32(4) | v1371
	v1381 = v1344 - v1380
	if base.Ui32(v1381) < base.Ui32(int32(32)) {
		goto L293
	} else {
		goto L300
	}
L300:
	;
	v1386 = base.I64_extend_i32_u(v1332) * int64(4294967297)
	v1389 = v1380 + v1336
	v1390 = v1381
	goto L301
L301:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1389)+24)) = v1386
	*(*int64)(unsafe.Add(mBase, uint32(v1389)+16)) = v1386
	*(*int64)(unsafe.Add(mBase, uint32(v1389)+8)) = v1386
	*(*int64)(unsafe.Add(mBase, uint32(v1389))) = v1386
	v1398 = int32(32)
	v1401 = v1390 - v1398
	if base.Ui32(int32(31)) < base.Ui32(v1401) {
		v1389 = v1389 + v1398
		v1390 = v1401
		goto L301
	} else {
		goto L303
	}
L302:
	;
	goto L293
L303:
	;
	goto L302
L304:
	;
	v1587 = int32(-1)
	goto L306
L305:
	;
	v1418 = int32(18)
	v1420 = v1298 + int32(8)
	v1421 = int32(0)
	v1423 = m.G0
	v1425 = v1423 + int32(-8192)
	m.G0 = v1425
	v1428 = int32(20)
	F___memset(m, v1425, v1421, v1428)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1423)+uint32(_c_F_InitPostgres[45]))) = uint8(v1421)
	*(*int32)(unsafe.Add(mBase, uint32(v1423)+uint32(_c_F_InitPostgres[46]))) = int32(1)
	v1433 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1423)+uint32(_c_F_InitPostgres[47]))) = uint16(v1433)
	*(*uint16)(unsafe.Add(mBase, uint32(v1423)+uint32(_c_F_InitPostgres[48]))) = uint16(v1418)
	*(*int32)(unsafe.Add(mBase, uint32(v1423)+uint32(_c_F_InitPostgres[49]))) = v1428
	v1439 = F_sendto(m, v1413, v1425, v1428)
	mBase = m.M
	if v1439 < v1421 {
		v1493 = v1439
		goto L308
	} else {
		goto L309
	}
L306:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+8))
	if v1587 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L307:
	;
	if v1493 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L308:
	;
	m.G0 = v1425 - int32(-8192)
	goto L307
L309:
	;
	v1442 = F_recvfrom(m, v1413, v1425)
	mBase = m.M
	if v1442 <= int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1493 = int32(-1)
	goto L308
L311:
	;
	v1447 = v1442
	goto L312
L312:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1447) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1493 = int32(0)
	goto L308
L314:
	;
	goto L313
L315:
	;
	v1455 = v1425
	goto L318
L316:
	;
	goto L317
L317:
	;
	v1480 = F_recvfrom(m, v1413, v1425)
	mBase = m.M
	if int32(0) < v1480 {
		v1447 = v1480
		goto L312
	} else {
		goto L323
	}
L318:
	;
	v1461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1455)+4)))
	switch v1461 - int32(2) {
	case 0:
		v1493 = int32(-1)
		goto L308
	case 1:
		goto L314
	default:
		goto L320
	}
L319:
	;
	goto L317
L320:
	;
	v1464 = F_netlink_msg_to_ifaddr(m, v1420, v1455)
	mBase = m.M
	if v1464 != 0 {
		v1493 = v1464
		goto L308
	} else {
		goto L321
	}
L321:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1455)))
	v1470 = v1455 + (v1465+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1447+v1425-v1470) {
		v1455 = v1470
		goto L318
	} else {
		goto L322
	}
L322:
	;
	goto L319
L323:
	;
	goto L310
L324:
	;
	v1503 = int32(22)
	v1504 = int32(0)
	v1506 = m.G0
	v1508 = v1506 + int32(-8192)
	m.G0 = v1508
	v1511 = int32(20)
	F___memset(m, v1508, v1504, v1511)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_InitPostgres[45]))) = uint8(v1504)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_InitPostgres[46]))) = int32(2)
	v1516 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_InitPostgres[47]))) = uint16(v1516)
	*(*uint16)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_InitPostgres[48]))) = uint16(v1503)
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+uint32(_c_F_InitPostgres[49]))) = v1511
	v1522 = F_sendto(m, v1413, v1508, v1511)
	mBase = m.M
	if v1522 < v1504 {
		v1576 = v1522
		goto L328
	} else {
		goto L329
	}
L325:
	;
	v1583 = v1493
	goto L326
L326:
	;
	v1584 = m.Wasi_snapshot_preview1.Fd_close(m, v1413)
	mBase = m.M
	v1587 = v1583
	goto L306
L327:
	;
	v1583 = v1576
	goto L326
L328:
	;
	m.G0 = v1508 - int32(-8192)
	goto L327
L329:
	;
	v1525 = F_recvfrom(m, v1413, v1508)
	mBase = m.M
	if v1525 <= int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1576 = int32(-1)
	goto L328
L331:
	;
	v1530 = v1525
	goto L332
L332:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1530) {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	v1576 = int32(0)
	goto L328
L334:
	;
	goto L333
L335:
	;
	v1538 = v1508
	goto L338
L336:
	;
	goto L337
L337:
	;
	v1563 = F_recvfrom(m, v1413, v1508)
	mBase = m.M
	if int32(0) < v1563 {
		v1530 = v1563
		goto L332
	} else {
		goto L343
	}
L338:
	;
	v1544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+4)))
	switch v1544 - int32(2) {
	case 0:
		v1576 = int32(-1)
		goto L328
	case 1:
		goto L334
	default:
		goto L340
	}
L339:
	;
	goto L337
L340:
	;
	v1547 = F_netlink_msg_to_ifaddr(m, v1420, v1538)
	mBase = m.M
	if v1547 != 0 {
		v1576 = v1547
		goto L328
	} else {
		goto L341
	}
L341:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1538)))
	v1553 = v1538 + (v1548+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1530+v1508-v1553) {
		v1538 = v1553
		goto L338
	} else {
		goto L342
	}
L342:
	;
	goto L339
L343:
	;
	goto L330
L344:
	;
	m.G0 = v1298 + int32(272)
	if v1587 < int32(0) {
		goto L355
	} else {
		goto L356
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1292+int32(12)))) = v1588
	goto L344
L346:
	;
	goto L347
L347:
	;
	if v1588 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L344
L349:
	;
	v1593 = v1588
	goto L352
L350:
	;
	goto L351
L351:
	;
	goto L348
L352:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1593)))
	F_emscripten_builtin_free(m, v1593)
	mBase = m.M
	if v1595 != 0 {
		v1593 = v1595
		goto L352
	} else {
		goto L354
	}
L353:
	;
	goto L351
L354:
	;
	goto L353
L355:
	;
	v1775 = int32(-1)
	goto L357
L356:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+12))
	if v1605 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	m.G0 = v1292 + int32(144)
	if v1775 < int32(0) {
		goto L401
	} else {
		goto L402
	}
L358:
	;
	v1607 = v1292 + int32(24)
	v1617 = v1605
	goto L361
L359:
	;
	v1741 = int32(0)
	goto L360
L360:
	;
	if v1741 != 0 {
		goto L395
	} else {
		goto L396
	}
L361:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+12))
	if v1633 != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+12))
	v1741 = v1714
	goto L360
L363:
	;
	v1634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1633))))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+16))
	if v1635 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L364:
	;
	goto L365
L365:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	if v1713 != 0 {
		v1617 = v1713
		goto L361
	} else {
		goto L393
	}
L366:
	;
	v1672 = m.G0
	v1674 = v1672 - int32(128)
	m.G0 = v1674
	v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289)+8)))
	if v1676 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L367:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+16)) = uint16(v1634)
	v1669 = v1292 + int32(16)
	goto L366
L368:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+16)) = int64(0)
	v1659 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1607)+8)) = v1659
	*(*int64)(unsafe.Add(mBase, uint32(v1607))) = v1659
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+40)) = int32(0)
	goto L367
L369:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+16)) = int64(-4294967296)
	goto L367
L370:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+4))
	if v1650 != 0 {
		v1669 = v1635
		goto L366
	} else {
		goto L379
	}
L371:
	;
	switch v1634 - int32(2) {
	case 0:
		goto L369
	default:
		v1669 = v1292 + int32(16)
		goto L366
	case 8:
		goto L368
	}
L372:
	;
	v1638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1635))))
	if v1638 != v1634 {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	switch v1634 - int32(2) {
	case 0:
		goto L370
	default:
		v1669 = v1635
		goto L366
	case 8:
		goto L374
	}
L374:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+8))
	if v1642 != 0 {
		v1669 = v1635
		goto L366
	} else {
		goto L375
	}
L375:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+12))
	if v1643 != 0 {
		v1669 = v1635
		goto L366
	} else {
		goto L376
	}
L376:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+16))
	if v1644 != 0 {
		v1669 = v1635
		goto L366
	} else {
		goto L377
	}
L377:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+20))
	if v1645 != 0 {
		v1669 = v1635
		goto L366
	} else {
		goto L378
	}
L378:
	;
	goto L368
L379:
	;
	goto L369
L380:
	;
	goto L365
L381:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1289)))
	if v1679 == int32(1) {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	goto L383
L383:
	;
	m.G0 = v1674 + int32(128)
	goto L380
L384:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1289)+8)) = uint8(v1703)
	goto L383
L385:
	;
	v1682 = int32(0)
	v1684 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1633))))
	v1685 = F_pg_sockaddr_cidr_mask(m, v1674, v1682, v1684)
	mBase = m.M
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+4))
	v1687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1686))))
	v1688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1633))))
	if v1687 != v1688 {
		v1703 = v1682
		goto L384
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+4))
	v1695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1694))))
	v1696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1633))))
	if v1695 != v1696 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	v1690 = F_pg_range_sockaddr(m, v1686, v1633, v1674)
	mBase = m.M
	if v1690 == int32(0) {
		v1703 = v1682
		goto L384
	} else {
		goto L389
	}
L389:
	;
	v1703 = int32(1)
	goto L384
L390:
	;
	v1703 = int32(0)
	goto L384
L391:
	;
	v1698 = F_pg_range_sockaddr(m, v1694, v1633, v1669)
	mBase = m.M
	if v1698 == int32(0) {
		goto L390
	} else {
		goto L392
	}
L392:
	;
	v1703 = int32(1)
	goto L384
L393:
	;
	goto L362
L394:
	;
	v1775 = int32(0)
	goto L357
L395:
	;
	v1743 = v1741
	goto L398
L396:
	;
	goto L397
L397:
	;
	goto L394
L398:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1743)))
	F_emscripten_builtin_free(m, v1743)
	mBase = m.M
	if v1745 != 0 {
		v1743 = v1745
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
L401:
	;
	v1783 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L1
	} else {
		goto L404
	}
L402:
	;
	goto L403
L403:
	;
	v1796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997)+24)))
	if v1796 == int32(0) {
		goto L210
	} else {
		goto L408
	}
L404:
	;
	if v1783 == int32(0) {
		goto L210
	} else {
		goto L405
	}
L405:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_42), int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_39), int32(1218), int32(_a_F_InitPostgres_43))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	goto L210
L408:
	;
	goto L212
L409:
	;
	v1807 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011))))
	switch v1807 - int32(2) {
	case 0:
		goto L413
	default:
		goto L411
	case 8:
		goto L412
	}
L410:
	;
	if v1863 == int32(0) {
		goto L210
	} else {
		goto L419
	}
L411:
	;
	v1863 = int32(0)
	goto L410
L412:
	;
	v1818 = v1042 + int32(164)
	v1820 = v1042 + int32(32)
	v1822 = v939 + int32(152)
	v1824 = int32(0)
	goto L414
L413:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1042+int32(156))+4))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1042+int32(24))+4))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+4))
	v1863 = base.B2i32(v1810&(v1811^v1812) == int32(0))
	goto L410
L414:
	;
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824+v1818))))
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824+v1820))))
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824+v1822))))
	if v1830&(v1832^v1834) != 0 {
		goto L411
	} else {
		goto L416
	}
L415:
	;
	v1863 = int32(1)
	goto L410
L416:
	;
	v1838 = v1824 | int32(1)
	v1840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820+v1838))))
	v1842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822+v1838))))
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818+v1838))))
	if (v1840^v1842)&v1845 != 0 {
		goto L411
	} else {
		goto L417
	}
L417:
	;
	v1848 = v1824 + int32(2)
	if v1848 != int32(16) {
		v1824 = v1848
		goto L414
	} else {
		goto L418
	}
L418:
	;
	goto L415
L419:
	;
	goto L212
L420:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+4))
	if v1894 <= int32(0) {
		goto L210
	} else {
		goto L421
	}
L421:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v939)+360))
	v1906 = int32(0)
	goto L422
L422:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+12))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1925+v1906<<(uint(int32(2))%32))))
	v1930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1929)+4)))
	v1932 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v1932 != int32(1) {
		goto L426
	} else {
		goto L427
	}
L423:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+20))
	v2215 = F_check_role_2(m, v2213, v1001, v2214)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L516
	}
L424:
	;
	goto L423
L425:
	;
	v2207 = v1906 + int32(1)
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+4))
	if v2207 < v2208 {
		v1906 = v2207
		goto L422
	} else {
		goto L515
	}
L426:
	;
	if v1930&int32(1) == int32(0) {
		goto L438
	} else {
		goto L439
	}
L427:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	if v1936&int32(1) != 0 {
		goto L426
	} else {
		goto L428
	}
L428:
	;
	if v1930&int32(1) != 0 {
		goto L425
	} else {
		goto L429
	}
L429:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	v1942 = int32(_a_F_InitPostgres_44)
	v1945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1941))))
	v1948 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[52])))
	if base.B2i32(v1945 == int32(0))|base.B2i32(v1945 != v1948) != 0 {
		v1966 = v1945
		v1967 = v1948
		goto L431
	} else {
		goto L432
	}
L430:
	;
	if v1966-v1967 != 0 {
		goto L425
	} else {
		goto L437
	}
L431:
	;
	goto L430
L432:
	;
	v1951 = v1941
	v1952 = v1942
	goto L433
L433:
	;
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1952)+1)))
	v1956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1951)+1)))
	if v1956 == int32(0) {
		v1966 = v1956
		v1967 = v1955
		goto L431
	} else {
		goto L435
	}
L434:
	;
	v1966 = v1956
	v1967 = v1955
	goto L431
L435:
	;
	v1959 = int32(1)
	if v1956 == v1955 {
		v1951 = v1951 + v1959
		v1952 = v1952 + v1959
		goto L433
	} else {
		goto L436
	}
L436:
	;
	goto L434
L437:
	;
	goto L424
L438:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	v1974 = int32(_a_F_InitPostgres_45)
	v1977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[53])))
	if base.B2i32(v1977 == int32(0))|base.B2i32(v1977 != v1980) != 0 {
		v1998 = v1977
		v1999 = v1980
		goto L442
	} else {
		goto L443
	}
L439:
	;
	goto L440
L440:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v1929)+8))
	if v2153 != 0 {
		goto L499
	} else {
		goto L500
	}
L441:
	;
	if v1998-v1999 == int32(0) {
		goto L424
	} else {
		goto L448
	}
L442:
	;
	goto L441
L443:
	;
	v1983 = v1973
	v1984 = v1974
	goto L444
L444:
	;
	v1987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1984)+1)))
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1983)+1)))
	if v1988 == int32(0) {
		v1998 = v1988
		v1999 = v1987
		goto L442
	} else {
		goto L446
	}
L445:
	;
	v1998 = v1988
	v1999 = v1987
	goto L442
L446:
	;
	v1991 = int32(1)
	if v1988 == v1987 {
		v1983 = v1983 + v1991
		v1984 = v1984 + v1991
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	v2003 = int32(_a_F_InitPostgres_46)
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[54])))
	if base.B2i32(v2006 == int32(0))|base.B2i32(v2006 != v2009) != 0 {
		v2027 = v2006
		v2028 = v2009
		goto L450
	} else {
		goto L451
	}
L449:
	;
	if v2027-v2028 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L450:
	;
	goto L449
L451:
	;
	v2012 = v1973
	v2013 = v2003
	goto L452
L452:
	;
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2013)+1)))
	v2017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2012)+1)))
	if v2017 == int32(0) {
		v2027 = v2017
		v2028 = v2016
		goto L450
	} else {
		goto L454
	}
L453:
	;
	v2027 = v2017
	v2028 = v2016
	goto L450
L454:
	;
	v2020 = int32(1)
	if v2017 == v2016 {
		v2012 = v2012 + v2020
		v2013 = v2013 + v2020
		goto L452
	} else {
		goto L455
	}
L455:
	;
	goto L453
L456:
	;
	v2034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1898))))
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897))))
	if base.B2i32(v2034 == int32(0))|base.B2i32(v2034 != v2037) != 0 {
		v2055 = v2034
		v2056 = v2037
		goto L460
	} else {
		goto L461
	}
L457:
	;
	goto L458
L458:
	;
	v2060 = int32(_a_F_InitPostgres_47)
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[55])))
	if base.B2i32(v2063 == int32(0))|base.B2i32(v2063 != v2066) != 0 {
		v2084 = v2063
		v2085 = v2066
		goto L469
	} else {
		goto L470
	}
L459:
	;
	if v2055-v2056 == int32(0) {
		goto L424
	} else {
		goto L466
	}
L460:
	;
	goto L459
L461:
	;
	v2040 = v1898
	v2041 = v1897
	goto L462
L462:
	;
	v2044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2041)+1)))
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2040)+1)))
	if v2045 == int32(0) {
		v2055 = v2045
		v2056 = v2044
		goto L460
	} else {
		goto L464
	}
L463:
	;
	v2055 = v2045
	v2056 = v2044
	goto L460
L464:
	;
	v2048 = int32(1)
	if v2045 == v2044 {
		v2040 = v2040 + v2048
		v2041 = v2041 + v2048
		goto L462
	} else {
		goto L465
	}
L465:
	;
	goto L463
L466:
	;
	goto L425
L467:
	;
	v2123 = int32(_a_F_InitPostgres_44)
	v2126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[52])))
	if base.B2i32(v2126 == int32(0))|base.B2i32(v2126 != v2129) != 0 {
		v2147 = v2126
		v2148 = v2129
		goto L492
	} else {
		goto L493
	}
L468:
	;
	if v2084-v2085 != 0 {
		goto L475
	} else {
		goto L476
	}
L469:
	;
	goto L468
L470:
	;
	v2069 = v1973
	v2070 = v2060
	goto L471
L471:
	;
	v2073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2070)+1)))
	v2074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2069)+1)))
	if v2074 == int32(0) {
		v2084 = v2074
		v2085 = v2073
		goto L469
	} else {
		goto L473
	}
L472:
	;
	v2084 = v2074
	v2085 = v2073
	goto L469
L473:
	;
	v2077 = int32(1)
	if v2074 == v2073 {
		v2069 = v2069 + v2077
		v2070 = v2070 + v2077
		goto L471
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	v2087 = int32(_a_F_InitPostgres_48)
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v2093 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[56])))
	if base.B2i32(v2090 == int32(0))|base.B2i32(v2090 != v2093) != 0 {
		v2111 = v2090
		v2112 = v2093
		goto L479
	} else {
		goto L480
	}
L476:
	;
	goto L477
L477:
	;
	if v1001 == int32(0) {
		goto L425
	} else {
		goto L486
	}
L478:
	;
	if v2111-v2112 != 0 {
		goto L467
	} else {
		goto L485
	}
L479:
	;
	goto L478
L480:
	;
	v2096 = v1973
	v2097 = v2087
	goto L481
L481:
	;
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2097)+1)))
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2096)+1)))
	if v2101 == int32(0) {
		v2111 = v2101
		v2112 = v2100
		goto L479
	} else {
		goto L483
	}
L482:
	;
	v2111 = v2101
	v2112 = v2100
	goto L479
L483:
	;
	v2104 = int32(1)
	if v2101 == v2100 {
		v2096 = v2096 + v2104
		v2097 = v2097 + v2104
		goto L481
	} else {
		goto L484
	}
L484:
	;
	goto L482
L485:
	;
	goto L477
L486:
	;
	v2117 = F_get_role_oid(m, v1898, int32(1))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	if v2117 == int32(0) {
		goto L425
	} else {
		goto L488
	}
L488:
	;
	v2121 = F_is_member_of_role_nosuper(m, v1001, v2117)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	if v2121 != 0 {
		goto L424
	} else {
		goto L490
	}
L490:
	;
	goto L425
L491:
	;
	if v2147-v2148 == int32(0) {
		goto L425
	} else {
		goto L498
	}
L492:
	;
	goto L491
L493:
	;
	v2132 = v1973
	v2133 = v2123
	goto L494
L494:
	;
	v2136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+1)))
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2132)+1)))
	if v2137 == int32(0) {
		v2147 = v2137
		v2148 = v2136
		goto L492
	} else {
		goto L496
	}
L495:
	;
	v2147 = v2137
	v2148 = v2136
	goto L492
L496:
	;
	v2140 = int32(1)
	if v2137 == v2136 {
		v2132 = v2132 + v2140
		v2133 = v2133 + v2140
		goto L494
	} else {
		goto L497
	}
L497:
	;
	goto L495
L498:
	;
	goto L440
L499:
	;
	v2154 = F_strlen(m, v1898)
	mBase = m.M
	v2159 = F_palloc(m, v2154<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L1
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	v2177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2174))))
	v2180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1898))))
	if base.B2i32(v2177 == int32(0))|base.B2i32(v2177 != v2180) != 0 {
		v2198 = v2177
		v2199 = v2180
		goto L508
	} else {
		goto L509
	}
L502:
	;
	v2161 = F_strlen(m, v1898)
	mBase = m.M
	v2162 = F_pg_mb2wchar_with_len(m, v1898, v2159, v2161)
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1929)+8))
	v2165 = int32(0)
	v2168 = F_pg_regexec(m, v2164, v2159, v2162, v2165, v2165, v2165)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	F_pfree(m, v2159)
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	if v2168 == int32(0) {
		goto L424
	} else {
		goto L506
	}
L506:
	;
	goto L425
L507:
	;
	if v2198-v2199 == int32(0) {
		goto L424
	} else {
		goto L514
	}
L508:
	;
	goto L507
L509:
	;
	v2183 = v2174
	v2184 = v1898
	goto L510
L510:
	;
	v2187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184)+1)))
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2183)+1)))
	if v2188 == int32(0) {
		v2198 = v2188
		v2199 = v2187
		goto L508
	} else {
		goto L512
	}
L511:
	;
	v2198 = v2188
	v2199 = v2187
	goto L508
L512:
	;
	v2191 = int32(1)
	if v2188 == v2187 {
		v2183 = v2183 + v2191
		v2184 = v2184 + v2191
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
L514:
	;
	goto L425
L515:
	;
	goto L210
L516:
	;
	if v2215 == int32(0) {
		goto L210
	} else {
		goto L517
	}
L517:
	;
	v2295 = v1042
	goto L204
L518:
	;
	goto L209
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2277)+296)) = int32(1)
	v2295 = v2277
	goto L204
L520:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+356))
	if v2315 == int32(0) {
		goto L527
	} else {
		goto L528
	}
L523:
	;
	goto L522
L524:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v4230 = m.ExcPending
	if v4230 != 0 {
		goto L1
	} else {
		goto L983
	}
L525:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v4227 = m.ExcPending
	if v4227 != 0 {
		goto L1
	} else {
		goto L982
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+104)) = int32(_a_F_InitPostgres_49)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+100)) = v2342
	*(*int32)(unsafe.Add(mBase, uint32(v989)+96)) = v989 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_50), v989+int32(96))
	mBase = m.M
	v4219 = m.ExcPending
	if v4219 != 0 {
		goto L1
	} else {
		goto L980
	}
L527:
	;
	v2318 = int32(-1)
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+296))
	switch v2319 {
	case 0:
		goto L538
	case 1:
		goto L537
	case 2, 12:
		goto L532
	case 3:
		goto L535
	case 4:
		goto L533
	case 5, 6:
		goto L534
	default:
		v4043 = v2318
		goto L530
	case 13:
		goto L536
	case 14:
		goto L531
	}
L528:
	;
	goto L529
L529:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L1
	} else {
		goto L976
	}
L530:
	;
	v4063 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[57])))
	if base.B2i32(v4063&int32(2) == int32(0))|v4043 != 0 {
		goto L929
	} else {
		goto L930
	}
L531:
	;
	v4035 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_51), v939, int32(0), v989+int32(444), v989+int32(443))
	mBase = m.M
	v4036 = m.ExcPending
	if v4036 != 0 {
		goto L1
	} else {
		goto L928
	}
L532:
	;
	v4043 = int32(0)
	goto L530
L533:
	;
	v3947 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v3947 != 0 {
		goto L895
	} else {
		goto L896
	}
L534:
	;
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v3552 = F_get_role_password(m, v3549, v989+int32(444))
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L1
	} else {
		goto L792
	}
L535:
	;
	v2710 = v989 + int32(2860)
	v2713 = int32(132)
	base.MemoryCopy(m, v2710, v939+int32(144), v2713)
	v2716 = v989 + int32(2728)
	base.MemoryCopy(m, v2716, v939+int32(12), v2713)
	v2721 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1516)) = v2721
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1512)) = v2721
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(v989)+2988))
	v2728 = v989 + int32(1952)
	v2734 = F_pg_getnameinfo_all(m, v2710, v2726, v2728, int32(255), v989+int32(1920), int32(32), int32(3))
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L645
	}
L536:
	;
	v2627 = m.G0
	v2629 = v2627 - int32(16)
	m.G0 = v2629
	*(*int32)(unsafe.Add(mBase, uint32(v2629))) = int32(12)
	goto L623
L537:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v939)+272))
	v2370 = v989 + int32(448)
	v2372 = int32(0)
	v2375 = F_pg_getnameinfo_all(m, v939+int32(144), v2368, v2370, int32(255), v2372, v2372, int32(1))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L545
	}
L538:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v939)+272))
	v2324 = v989 + int32(448)
	v2326 = int32(0)
	v2329 = F_pg_getnameinfo_all(m, v939+int32(144), v2322, v2324, int32(255), v2326, v2326, int32(1))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	v2332 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v2343 = int32(1)
	if base.B2i32(v2332&v2343 == int32(0))&base.B2i32(v2334 == v2343) != 0 {
		goto L526
	} else {
		goto L542
	}
L542:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v939)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+92)) = int32(_a_F_InitPostgres_49)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+88)) = v2350
	*(*int32)(unsafe.Add(mBase, uint32(v989)+84)) = v2342
	*(*int32)(unsafe.Add(mBase, uint32(v989)+80)) = v2324
	F_errmsg(m, int32(_a_F_InitPostgres_52), v989+int32(80))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(474), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L545:
	;
	v2378 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v2389 = int32(1)
	if v2378&v2389|base.B2i32(v2380 != v2389) == int32(0) {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+296)) = int32(_a_F_InitPostgres_49)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+292)) = v2388
	*(*int32)(unsafe.Add(mBase, uint32(v989)+288)) = v2370
	F_errmsg(m, int32(_a_F_InitPostgres_55), v989+int32(288))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L1
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(v939)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+204)) = int32(_a_F_InitPostgres_49)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+200)) = v2509
	*(*int32)(unsafe.Add(mBase, uint32(v989)+196)) = v2388
	*(*int32)(unsafe.Add(mBase, uint32(v989)+192)) = v989 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_56), v989+int32(192))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L587
	}
L551:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v939)+284))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v939)+280))
	if v2406 != 0 {
		goto L553
	} else {
		goto L554
	}
L552:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(534), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L1
	} else {
		goto L586
	}
L553:
	;
	switch v2405 + int32(2) {
	case 0:
		goto L556
	case 1:
		goto L557
	case 2:
		goto L558
	case 3:
		goto L559
	default:
		goto L552
	}
L554:
	;
	goto L555
L555:
	;
	if v2405 != int32(-2) {
		goto L552
	} else {
		goto L574
	}
L556:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v939)+288))
	v2430 = int32(_a_F_InitPostgres_57)
	v2432 = v2427 + int32(1)
	if v2432 == int32(0) {
		v2452 = v2430
		goto L564
	} else {
		goto L565
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+256)) = v2406
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v989+int32(256))
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L562
	}
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+240)) = v2406
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v989+int32(240))
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L1
	} else {
		goto L561
	}
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+224)) = v2406
	F_errdetail_log(m, int32(_a_F_InitPostgres_60), v989+int32(224))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L1
	} else {
		goto L560
	}
L560:
	;
	goto L552
L561:
	;
	goto L552
L562:
	;
	goto L552
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+276)) = v2452 + base.B2i32(v2454 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+272)) = v2406
	F_errdetail_log(m, int32(_a_F_InitPostgres_61), v989+int32(272))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L573
	}
L564:
	;
	v2454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2452))))
	goto L563
L565:
	;
	v2436 = v2430
	v2437 = v2432
	goto L566
L566:
	;
	v2438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2436))))
	if v2438 == int32(0) {
		v2452 = v2436
		goto L564
	} else {
		goto L568
	}
L567:
	;
	v2452 = v2448
	goto L564
L568:
	;
	v2442 = v2436
	goto L569
L569:
	;
	v2446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2442)+1)))
	if v2446 != 0 {
		v2442 = v2442 + int32(1)
		goto L569
	} else {
		goto L571
	}
L570:
	;
	v2448 = v2442 + int32(2)
	v2450 = v2437 + int32(1)
	if v2450 != 0 {
		v2436 = v2448
		v2437 = v2450
		goto L566
	} else {
		goto L572
	}
L571:
	;
	goto L570
L572:
	;
	goto L567
L573:
	;
	goto L552
L574:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v939)+288))
	v2470 = int32(_a_F_InitPostgres_57)
	v2472 = v2467 + int32(1)
	if v2472 == int32(0) {
		v2492 = v2470
		goto L576
	} else {
		goto L577
	}
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+208)) = v2492 + base.B2i32(v2494 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_62), v989+int32(208))
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L1
	} else {
		goto L585
	}
L576:
	;
	v2494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492))))
	goto L575
L577:
	;
	v2476 = v2470
	v2477 = v2472
	goto L578
L578:
	;
	v2478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476))))
	if v2478 == int32(0) {
		v2492 = v2476
		goto L576
	} else {
		goto L580
	}
L579:
	;
	v2492 = v2488
	goto L576
L580:
	;
	v2482 = v2476
	goto L581
L581:
	;
	v2486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2482)+1)))
	if v2486 != 0 {
		v2482 = v2482 + int32(1)
		goto L581
	} else {
		goto L583
	}
L582:
	;
	v2488 = v2482 + int32(2)
	v2490 = v2477 + int32(1)
	if v2490 != 0 {
		v2476 = v2488
		v2477 = v2490
		goto L578
	} else {
		goto L584
	}
L583:
	;
	goto L582
L584:
	;
	goto L579
L585:
	;
	goto L552
L586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L587:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v939)+284))
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v939)+280))
	if v2523 != 0 {
		goto L589
	} else {
		goto L590
	}
L588:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(543), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L1
	} else {
		goto L622
	}
L589:
	;
	switch v2522 + int32(2) {
	case 0:
		goto L592
	case 1:
		goto L593
	case 2:
		goto L594
	case 3:
		goto L595
	default:
		goto L588
	}
L590:
	;
	goto L591
L591:
	;
	if v2522 != int32(-2) {
		goto L588
	} else {
		goto L610
	}
L592:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v939)+288))
	v2547 = int32(_a_F_InitPostgres_57)
	v2549 = v2544 + int32(1)
	if v2549 == int32(0) {
		v2569 = v2547
		goto L600
	} else {
		goto L601
	}
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+160)) = v2523
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v989+int32(160))
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L1
	} else {
		goto L598
	}
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+144)) = v2523
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v989+int32(144))
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L1
	} else {
		goto L597
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+128)) = v2523
	F_errdetail_log(m, int32(_a_F_InitPostgres_60), v989+int32(128))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	goto L588
L597:
	;
	goto L588
L598:
	;
	goto L588
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+180)) = v2569 + base.B2i32(v2571 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+176)) = v2523
	F_errdetail_log(m, int32(_a_F_InitPostgres_61), v989+int32(176))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L1
	} else {
		goto L609
	}
L600:
	;
	v2571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2569))))
	goto L599
L601:
	;
	v2553 = v2547
	v2554 = v2549
	goto L602
L602:
	;
	v2555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2553))))
	if v2555 == int32(0) {
		v2569 = v2553
		goto L600
	} else {
		goto L604
	}
L603:
	;
	v2569 = v2565
	goto L600
L604:
	;
	v2559 = v2553
	goto L605
L605:
	;
	v2563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2559)+1)))
	if v2563 != 0 {
		v2559 = v2559 + int32(1)
		goto L605
	} else {
		goto L607
	}
L606:
	;
	v2565 = v2559 + int32(2)
	v2567 = v2554 + int32(1)
	if v2567 != 0 {
		v2553 = v2565
		v2554 = v2567
		goto L602
	} else {
		goto L608
	}
L607:
	;
	goto L606
L608:
	;
	goto L603
L609:
	;
	goto L588
L610:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v939)+288))
	v2587 = int32(_a_F_InitPostgres_57)
	v2589 = v2584 + int32(1)
	if v2589 == int32(0) {
		v2609 = v2587
		goto L612
	} else {
		goto L613
	}
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+112)) = v2609 + base.B2i32(v2611 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_62), v989+int32(112))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L621
	}
L612:
	;
	v2611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2609))))
	goto L611
L613:
	;
	v2593 = v2587
	v2594 = v2589
	goto L614
L614:
	;
	v2595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2593))))
	if v2595 == int32(0) {
		v2609 = v2593
		goto L612
	} else {
		goto L616
	}
L615:
	;
	v2609 = v2605
	goto L612
L616:
	;
	v2599 = v2593
	goto L617
L617:
	;
	v2603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2599)+1)))
	if v2603 != 0 {
		v2599 = v2599 + int32(1)
		goto L617
	} else {
		goto L619
	}
L618:
	;
	v2605 = v2599 + int32(2)
	v2607 = v2594 + int32(1)
	if v2607 != 0 {
		v2593 = v2605
		v2594 = v2607
		goto L614
	} else {
		goto L620
	}
L619:
	;
	goto L618
L620:
	;
	goto L615
L621:
	;
	goto L588
L622:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L623:
	;
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v989+int32(1952)))) = v2639
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v989+int32(1664)))) = v2643
	goto L625
L625:
	;
	m.G0 = v2629 + int32(16)
	goto L627
L627:
	;
	goto L628
L628:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44])) = int32(44)
	v2693 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	if v2693 == int32(0) {
		v4043 = v2318
		goto L530
	} else {
		goto L642
	}
L642:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1952))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+320)) = v2697
	F_errmsg(m, int32(_a_F_InitPostgres_63), v989+int32(320))
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(1935), int32(_a_F_InitPostgres_64))
	mBase = m.M
	v2708 = m.ExcPending
	if v2708 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	v4043 = v2318
	goto L530
L645:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v989)+2856))
	v2738 = v989 + int32(1664)
	v2744 = F_pg_getnameinfo_all(m, v2716, v2736, v2738, int32(255), v989+int32(1632), int32(32), int32(3))
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+432)) = int32(113)
	v2749 = v989 + int32(1600)
	v2754 = F_pg_snprintf(m, v2749, int32(32), int32(_a_F_InitPostgres_65), v989+int32(432))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1480)) = int32(4)
	v2758 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v989)+1492)) = v2758
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1488)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v989)+1500)) = v2758
	v2764 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1508)) = v2764
	v2766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v989)+2860)))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1484)) = v2766
	v2769 = v989 + int32(1480)
	v2772 = F_pg_getaddrinfo_all(m, v2728, v2749, v2769, v989+int32(1516))
	mBase = m.M
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1516))
	if v2772|base.B2i32(v2773 == v2764) == v2764 {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1480)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1488)) = int32(1)
	v2783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v989)+2728)))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+1484)) = v2783
	v2786 = v989 + int32(1492)
	v2787 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2786)+16)) = v2787
	v2789 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2786)+8)) = v2789
	*(*int64)(unsafe.Add(mBase, uint32(v2786))) = v2789
	v2796 = F_pg_getaddrinfo_all(m, v2738, v2787, v2769, v989+int32(1512))
	mBase = m.M
	if v2796 != 0 {
		v3466 = v2721
		goto L651
	} else {
		goto L652
	}
L649:
	;
	v3485 = v2773
	v3492 = v2721
	goto L650
L650:
	;
	if v3485 != 0 {
		goto L762
	} else {
		goto L763
	}
L651:
	;
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1516))
	v3485 = v3478
	v3492 = v3466
	goto L650
L652:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1512))
	if v2797 == int32(0) {
		v3466 = v2721
		goto L651
	} else {
		goto L653
	}
L653:
	;
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1516))
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v2800)+4))
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v2800)+8))
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2800)+12))
	v2804 = F_socket(m, v2801, v2802, v2803)
	mBase = m.M
	if v2804 == int32(-1) {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	v2809 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L1
	} else {
		goto L657
	}
L655:
	;
	goto L656
L656:
	;
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1512))
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v2825)+20))
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2825)+16))
	v2828 = F_bind(m, v2804, v2826, v2827)
	mBase = m.M
	if v2828 != 0 {
		goto L664
	} else {
		goto L665
	}
L657:
	;
	if v2809 == int32(0) {
		v3466 = v2721
		goto L651
	} else {
		goto L658
	}
L658:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_66), int32(0))
	mBase = m.M
	v2818 = m.ExcPending
	if v2818 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(1789), int32(_a_F_InitPostgres_67))
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	v3466 = v2721
	goto L651
L662:
	;
	v3452 = F_close(m, v2804)
	mBase = m.M
	v3466 = v3440
	goto L651
L663:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), v3423, int32(_a_F_InitPostgres_67))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L761
	}
L664:
	;
	v2831 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+388)) = v989 + int32(1632)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+384)) = v989 + int32(1920)
	v2861 = F_pg_snprintf(m, v989+int32(1520), int32(80), int32(_a_F_InitPostgres_68), v989+int32(384))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L1
	} else {
		goto L671
	}
L667:
	;
	if v2831 == int32(0) {
		v3440 = v2721
		goto L662
	} else {
		goto L668
	}
L668:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v2836 = m.ExcPending
	if v2836 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+416)) = v989 + int32(1664)
	F_errmsg(m, int32(_a_F_InitPostgres_69), v989+int32(416))
	mBase = m.M
	v2844 = m.ExcPending
	if v2844 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	v3411 = v2721
	v3423 = int32(1805)
	goto L663
L671:
	;
	goto L673
L672:
	;
	v2982 = v989 + int32(448)
	v2984 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2982+v2930))) = uint8(v2984)
	v2988 = v989 + int32(2208)
	v2990 = m.G0
	v2992 = v2990 - int32(80)
	m.G0 = v2992
	v2994 = F_strlen(m, v2982)
	mBase = m.M
	if base.Ui32(v2994) < base.Ui32(int32(2)) {
		v3361 = v2984
		goto L699
	} else {
		goto L700
	}
L673:
	;
	v2889 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v2889 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	v2963 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L1
	} else {
		goto L695
	}
L675:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L1
	} else {
		goto L678
	}
L676:
	;
	goto L677
L677:
	;
	v2893 = v989 + int32(1520)
	v2894 = F_strlen(m, v2893)
	mBase = m.M
	v2895 = F_pgmem_send(m, v2804, v2893, v2894)
	mBase = m.M
	if int32(0) <= v2895 {
		goto L679
	} else {
		goto L680
	}
L678:
	;
	goto L677
L679:
	;
	goto L682
L680:
	;
	goto L681
L681:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v2958 == int32(27) {
		goto L673
	} else {
		goto L694
	}
L682:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v2924 != 0 {
		goto L684
	} else {
		goto L685
	}
L683:
	;
	v2939 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L1
	} else {
		goto L690
	}
L684:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2926 = m.ExcPending
	if v2926 != 0 {
		goto L1
	} else {
		goto L687
	}
L685:
	;
	goto L686
L686:
	;
	v2930 = F_pgmem_recv(m, v2804, v989+int32(448), int32(591))
	mBase = m.M
	if int32(0) <= v2930 {
		goto L672
	} else {
		goto L688
	}
L687:
	;
	goto L686
L688:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v2934 == int32(27) {
		goto L682
	} else {
		goto L689
	}
L689:
	;
	goto L683
L690:
	;
	if v2939 == int32(0) {
		v3440 = v2721
		goto L662
	} else {
		goto L691
	}
L691:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+356)) = v989 + int32(1600)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+352)) = v989 + int32(1952)
	F_errmsg(m, int32(_a_F_InitPostgres_70), v989+int32(352))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L1
	} else {
		goto L693
	}
L693:
	;
	v3411 = v2721
	v3423 = int32(1856)
	goto L663
L694:
	;
	goto L674
L695:
	;
	if v2963 == int32(0) {
		v3440 = v2721
		goto L662
	} else {
		goto L696
	}
L696:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+340)) = v989 + int32(1600)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+336)) = v989 + int32(1952)
	F_errmsg(m, int32(_a_F_InitPostgres_71), v989+int32(336))
	mBase = m.M
	v2979 = m.ExcPending
	if v2979 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	v3411 = v2721
	v3423 = int32(1839)
	goto L663
L699:
	;
	m.G0 = v2992 + int32(80)
	if v3361 != 0 {
		v3440 = int32(1)
		goto L662
	} else {
		goto L757
	}
L700:
	;
	v3000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2994+v2982-int32(2)))))
	if v3000 != int32(13) {
		v3361 = v2984
		goto L699
	} else {
		goto L701
	}
L701:
	;
	v3012 = v2982
	goto L702
L702:
	;
	v3028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3012))))
	if v3028 == int32(13) {
		v3361 = v2984
		goto L699
	} else {
		goto L704
	}
L703:
	;
	v3044 = v3012
	goto L708
L704:
	;
	if v3028 != int32(58) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v3012 = v3012 + int32(1)
	goto L702
L706:
	;
	goto L707
L707:
	;
	goto L703
L708:
	;
	v3061 = v3044 + int32(1)
	v3062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3044)+1)))
	if base.B2i32(v3062 == int32(9))|base.B2i32(v3062 == int32(32)) != 0 {
		v3044 = v3061
		goto L708
	} else {
		goto L710
	}
L709:
	;
	v3074 = v3062
	v3075 = v2984
	v3077 = v3061
	goto L718
L710:
	;
	goto L709
L711:
	;
	v3173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3171+v2992))) = uint8(v3173)
	v3181 = v3169
	v3184 = v3170
	goto L737
L712:
	;
	v3169 = v3114
	v3170 = v3077 + int32(1)
	v3171 = v3113
	goto L711
L713:
	;
	v3169 = v3114
	v3170 = v3077 + int32(1)
	v3171 = v3113
	goto L711
L714:
	;
	v3169 = v3123
	v3170 = v3077 + int32(2)
	v3171 = v3122
	goto L711
L715:
	;
	v3169 = v3123
	v3170 = v3077 + int32(2)
	v3171 = v3122
	goto L711
L716:
	;
	v3169 = v3132
	v3170 = v3077 + int32(3)
	v3171 = v3131
	goto L711
L717:
	;
	v3169 = v3132
	v3170 = v3077 + int32(3)
	v3171 = int32(79)
	goto L711
L718:
	;
	v3094 = v3074 & int32(255)
	v3096 = v3094 - int32(9)
	v3103 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v3096))|base.B2i32(int32(1)<<(uint(v3096)%32)&int32(_a_F_InitPostgres_72) == v3103) == v3103 {
		goto L720
	} else {
		goto L721
	}
L719:
	;
	v3169 = v3132
	v3170 = v3077 + int32(3)
	v3171 = v3131
	goto L711
L720:
	;
	v3169 = v3074
	v3170 = v3077
	v3171 = v3075
	goto L711
L721:
	;
	goto L722
L722:
	;
	if v3094 == int32(58) {
		goto L723
	} else {
		goto L724
	}
L723:
	;
	v3169 = v3074
	v3170 = v3077
	v3171 = v3075
	goto L711
L724:
	;
	goto L725
L725:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3075+v2992))) = uint8(v3074)
	v3113 = v3075 | int32(1)
	v3114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3077)+1)))
	switch v3114 - int32(9) {
	case 0, 23:
		goto L713
	case 1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22:
		goto L726
	case 4:
		goto L712
	default:
		goto L727
	}
L726:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3113+v2992))) = uint8(v3114)
	v3122 = v3075 | int32(2)
	v3123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3077)+2)))
	switch v3123 - int32(9) {
	case 0, 23:
		goto L715
	case 1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22:
		goto L729
	case 4:
		goto L714
	default:
		goto L730
	}
L727:
	;
	if v3114 == int32(58) {
		goto L712
	} else {
		goto L728
	}
L728:
	;
	goto L726
L729:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3122+v2992))) = uint8(v3123)
	v3131 = v3075 | int32(3)
	v3132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3077)+3)))
	if base.B2i32(v3132 == int32(13))|base.B2i32(v3132 == int32(58)) != 0 {
		goto L716
	} else {
		goto L732
	}
L730:
	;
	if v3123 == int32(58) {
		goto L714
	} else {
		goto L731
	}
L731:
	;
	goto L729
L732:
	;
	if v3075 == int32(76) {
		goto L717
	} else {
		goto L733
	}
L733:
	;
	if base.B2i32(v3132 == int32(9))|base.B2i32(v3132 == int32(32)) == int32(0) {
		goto L734
	} else {
		goto L735
	}
L734:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3131+v2992))) = uint8(v3132)
	v3149 = int32(4)
	v3151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3077)+4)))
	v3074 = v3151
	v3075 = v3075 + v3149
	v3077 = v3077 + v3149
	goto L718
L735:
	;
	goto L736
L736:
	;
	goto L719
L737:
	;
	v3201 = v3181 & int32(255)
	if base.B2i32(v3201 != int32(32))&base.B2i32(v3201 != int32(9)) == int32(0) {
		goto L739
	} else {
		goto L740
	}
L738:
	;
	v3212 = int32(0)
	if v3181&int32(255) != int32(58) {
		v3361 = v3212
		goto L699
	} else {
		goto L742
	}
L739:
	;
	v3209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3184)+1)))
	v3181 = v3209
	v3184 = v3184 + int32(1)
	goto L737
L740:
	;
	goto L741
L741:
	;
	goto L738
L742:
	;
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v2992)))
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v2992)+3))
	if v3217^int32(1380275029)|(v3220^int32(_a_F_InitPostgres_73)) != 0 {
		v3361 = v3212
		goto L699
	} else {
		goto L743
	}
L743:
	;
	v3233 = v3184
	goto L744
L744:
	;
	v3249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3233)+1)))
	if v3249 == int32(13) {
		v3361 = v3212
		goto L699
	} else {
		goto L746
	}
L745:
	;
	v3267 = v3233 + int32(2)
	goto L749
L746:
	;
	if v3249 != int32(58) {
		v3233 = v3233 + int32(1)
		goto L744
	} else {
		goto L747
	}
L747:
	;
	goto L745
L748:
	;
	v3351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3332+v2988))) = uint8(v3351)
	v3361 = int32(1)
	goto L699
L749:
	;
	v3284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3267))))
	switch v3284 - int32(9) {
	case 0, 23:
		goto L752
	default:
		goto L751
	case 4:
		v3332 = v3212
		goto L748
	}
L750:
	;
	v3295 = int32(0)
	v3298 = v3267
	v3300 = v3284
	goto L753
L751:
	;
	goto L750
L752:
	;
	v3267 = v3267 + int32(1)
	goto L749
L753:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3295+v2988))) = uint8(v3300)
	v3317 = v3295 + int32(1)
	v3318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3298)+1)))
	if v3318 == int32(13) {
		v3332 = v3317
		goto L748
	} else {
		goto L755
	}
L754:
	;
	v3332 = v3317
	goto L748
L755:
	;
	if base.Ui32(v3295) < base.Ui32(int32(511)) {
		v3295 = v3317
		v3298 = v3298 + int32(1)
		v3300 = v3318
		goto L753
	} else {
		goto L756
	}
L756:
	;
	goto L754
L757:
	;
	v3382 = int32(0)
	v3385 = F_errstart(m, int32(15), v3382)
	mBase = m.M
	v3386 = m.ExcPending
	if v3386 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	if v3385 == int32(0) {
		v3440 = v3382
		goto L662
	} else {
		goto L759
	}
L759:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+368)) = v989 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_74), v989+int32(368))
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L1
	} else {
		goto L760
	}
L760:
	;
	v3411 = v3382
	v3423 = int32(1866)
	goto L663
L761:
	;
	v3440 = v3411
	goto L662
L762:
	;
	v3504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v989)+2860)))
	if v3504 == int32(1) {
		goto L767
	} else {
		goto L768
	}
L763:
	;
	goto L764
L764:
	;
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(v989)+1512))
	if v3520 != 0 {
		goto L775
	} else {
		goto L776
	}
L765:
	;
	goto L764
L766:
	;
	goto L765
L767:
	;
	if v3485 == int32(0) {
		goto L766
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	if v3485 == int32(0) {
		goto L766
	} else {
		goto L774
	}
L770:
	;
	v3510 = v3485
	goto L771
L771:
	;
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3510)+28))
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v3510)+20))
	F_emscripten_builtin_free(m, v3512)
	mBase = m.M
	F_emscripten_builtin_free(m, v3510)
	mBase = m.M
	if v3511 != 0 {
		v3510 = v3511
		goto L771
	} else {
		goto L773
	}
L772:
	;
	goto L766
L773:
	;
	goto L772
L774:
	;
	F_freeaddrinfo(m, v3485)
	mBase = m.M
	goto L766
L775:
	;
	v3521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v989)+2728)))
	if v3521 == int32(1) {
		goto L780
	} else {
		goto L781
	}
L776:
	;
	goto L777
L777:
	;
	if v3492 == int32(0) {
		v4043 = int32(-1)
		goto L530
	} else {
		goto L788
	}
L778:
	;
	goto L777
L779:
	;
	goto L778
L780:
	;
	if v3520 == int32(0) {
		goto L779
	} else {
		goto L783
	}
L781:
	;
	goto L782
L782:
	;
	if v3520 == int32(0) {
		goto L779
	} else {
		goto L787
	}
L783:
	;
	v3527 = v3520
	goto L784
L784:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3527)+28))
	v3529 = *(*int32)(unsafe.Add(mBase, uint32(v3527)+20))
	F_emscripten_builtin_free(m, v3529)
	mBase = m.M
	F_emscripten_builtin_free(m, v3527)
	mBase = m.M
	if v3528 != 0 {
		v3527 = v3528
		goto L784
	} else {
		goto L786
	}
L785:
	;
	goto L779
L786:
	;
	goto L785
L787:
	;
	F_freeaddrinfo(m, v3520)
	mBase = m.M
	goto L779
L788:
	;
	v3541 = v989 + int32(2208)
	F_set_authn_id(m, v939, v3541)
	mBase = m.M
	v3543 = m.ExcPending
	if v3543 != 0 {
		goto L1
	} else {
		goto L789
	}
L789:
	;
	v3544 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v3544)+300))
	v3546 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v3547 = F_check_usermap(m, v3545, v3546, v3541)
	mBase = m.M
	v3548 = m.ExcPending
	if v3548 != 0 {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	v4043 = v3547
	goto L530
L791:
	;
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(v3561)+296))
	if base.B2i32(v3562 != int32(5))|base.B2i32(v3560 != int32(1)) == int32(0) {
		goto L798
	} else {
		goto L799
	}
L792:
	;
	if v3552 == int32(0) {
		goto L793
	} else {
		goto L794
	}
L793:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[58]))
	v3560 = v3557
	goto L791
L794:
	;
	goto L795
L795:
	;
	v3558 = F_get_password_type(m, v3552)
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	v3560 = v3558
	goto L791
L797:
	;
	if v3552 != 0 {
		goto L889
	} else {
		goto L890
	}
L798:
	;
	v3573 = int32(0)
	v3577 = m.G0
	v3579 = v3577 - int32(16)
	m.G0 = v3579
	*(*int32)(unsafe.Add(mBase, uint32(v3579))) = v3573
	v3585 = F_open(m, int32(_a_F_InitPostgres_75), v3573, v3579)
	mBase = m.M
	if v3585 != int32(-1) {
		goto L802
	} else {
		goto L803
	}
L799:
	;
	goto L800
L800:
	;
	v3933 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_76), v939, v3552, v989+int32(444), int32(0))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L1
	} else {
		goto L888
	}
L801:
	;
	if v3618 == int32(0) {
		goto L814
	} else {
		goto L815
	}
L802:
	;
	goto L806
L803:
	;
	v3618 = v3573
	goto L804
L804:
	;
	m.G0 = v3579 + int32(16)
	goto L801
L805:
	;
	v3613 = F_close(m, v3585)
	mBase = m.M
	v3618 = v3611
	goto L804
L806:
	;
	v3591 = v989 + int32(448)
	v3592 = int32(4)
	goto L807
L807:
	;
	v3597 = F_read(m, v3585, v3591, v3592)
	mBase = m.M
	if v3597 <= int32(0) {
		goto L809
	} else {
		goto L810
	}
L808:
	;
	v3611 = int32(1)
	goto L805
L809:
	;
	v3601 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v3601 == int32(27) {
		goto L807
	} else {
		goto L812
	}
L810:
	;
	goto L811
L811:
	;
	v3606 = v3592 - v3597
	if v3606 != 0 {
		v3591 = v3591 + v3597
		v3592 = v3606
		goto L807
	} else {
		goto L813
	}
L812:
	;
	v3611 = int32(0)
	goto L805
L813:
	;
	goto L808
L814:
	;
	v3625 = int32(-1)
	v3628 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L817
	}
L815:
	;
	goto L816
L816:
	;
	F_sendAuthRequest(m, int32(5), v989+int32(448), int32(4))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L821
	}
L817:
	;
	if v3628 == int32(0) {
		v3935 = v3625
		goto L797
	} else {
		goto L818
	}
L818:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_77), int32(0))
	mBase = m.M
	v3635 = m.ExcPending
	if v3635 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(906), int32(_a_F_InitPostgres_78))
	mBase = m.M
	v3640 = m.ExcPending
	if v3640 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	v3935 = v3625
	goto L797
L821:
	;
	v3647 = F_recv_password_packet(m)
	mBase = m.M
	v3648 = m.ExcPending
	if v3648 != 0 {
		goto L1
	} else {
		goto L822
	}
L822:
	;
	if v3647 == int32(0) {
		goto L823
	} else {
		goto L824
	}
L823:
	;
	v3935 = int32(-2)
	goto L797
L824:
	;
	goto L825
L825:
	;
	if v3552 == int32(0) {
		goto L826
	} else {
		goto L827
	}
L826:
	;
	F_pfree(m, v3647)
	mBase = m.M
	v3655 = m.ExcPending
	if v3655 != 0 {
		goto L1
	} else {
		goto L829
	}
L827:
	;
	goto L828
L828:
	;
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v3660 = m.G0
	v3662 = v3660 - int32(128)
	m.G0 = v3662
	v3664 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3662)+28)) = v3664
	*(*int32)(unsafe.Add(mBase, uint32(v3662)+116)) = v3664
	v3668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552))))
	if v3668 != int32(109) {
		goto L832
	} else {
		goto L833
	}
L829:
	;
	v3935 = int32(-1)
	goto L797
L830:
	;
	m.G0 = v3662 + int32(128)
	F_pfree(m, v3647)
	mBase = m.M
	v3910 = m.ExcPending
	if v3910 != 0 {
		goto L1
	} else {
		goto L883
	}
L831:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+444)) = v3900
	v3905 = int32(-1)
	goto L830
L832:
	;
	v3890 = F_parse_scram_secret(m, v3552, v3662+int32(120), v3662+int32(112), v3662+int32(116), v3662+int32(124), v3662+int32(32), v3662+int32(80))
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L1
	} else {
		goto L881
	}
L833:
	;
	v3671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552)+1)))
	if v3671 != int32(100) {
		goto L832
	} else {
		goto L834
	}
L834:
	;
	v3674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552)+2)))
	if v3674 != int32(53) {
		goto L832
	} else {
		goto L835
	}
L835:
	;
	v3677 = F_strlen(m, v3552)
	mBase = m.M
	if v3677 != int32(35) {
		goto L832
	} else {
		goto L836
	}
L836:
	;
	v3681 = v3552 + int32(3)
	v3682 = int32(_a_F_InitPostgres_79)
	v3686 = m.G0
	v3688 = v3686 - int32(32)
	v3689 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3688)+24)) = v3689
	*(*int64)(unsafe.Add(mBase, uint32(v3688)+16)) = v3689
	*(*int64)(unsafe.Add(mBase, uint32(v3688)+8)) = v3689
	*(*int64)(unsafe.Add(mBase, uint32(v3688))) = v3689
	v3697 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[59])))
	if v3697 == int32(0) {
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if v3765 != int32(32) {
		goto L832
	} else {
		goto L856
	}
L838:
	;
	v3765 = int32(0)
	goto L837
L839:
	;
	goto L840
L840:
	;
	v3701 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[60])))
	if v3701 == int32(0) {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v3705 = v3681
	goto L844
L842:
	;
	goto L843
L843:
	;
	v3715 = v3682
	v3716 = v3697
	goto L847
L844:
	;
	v3711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3705))))
	if v3711 == v3697 {
		v3705 = v3705 + int32(1)
		goto L844
	} else {
		goto L846
	}
L845:
	;
	v3765 = v3705 - v3681
	goto L837
L846:
	;
	goto L845
L847:
	;
	v3723 = v3688 + int32(base.Ui32(v3716)>>(uint(int32(3))%32))&int32(28)
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v3723)))
	v3725 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3723))) = v3724 | v3725<<(uint(v3716)%32)
	v3729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3715)+1)))
	if v3729 != 0 {
		v3715 = v3715 + v3725
		v3716 = v3729
		goto L847
	} else {
		goto L849
	}
L848:
	;
	v3732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3681))))
	if v3732 == int32(0) {
		v3755 = v3681
		goto L850
	} else {
		goto L851
	}
L849:
	;
	goto L848
L850:
	;
	v3765 = v3755 - v3681
	goto L837
L851:
	;
	v3736 = v3681
	v3737 = v3732
	goto L852
L852:
	;
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3688+int32(base.Ui32(v3737)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3745)>>(uint(v3737)%32))&int32(1) == int32(0) {
		v3755 = v3736
		goto L850
	} else {
		goto L854
	}
L853:
	;
	v3755 = v3753
	goto L850
L854:
	;
	v3751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3736)+1)))
	v3753 = v3736 + int32(1)
	if v3751 != 0 {
		v3736 = v3753
		v3737 = v3751
		goto L852
	} else {
		goto L855
	}
L855:
	;
	goto L853
L856:
	;
	v3773 = F_pg_md5_encrypt(m, v3681, v989+int32(448), int32(4), v3662+int32(32), v3662+int32(28))
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L857
	}
L857:
	;
	if v3773 == int32(0) {
		goto L858
	} else {
		goto L859
	}
L858:
	;
	v3777 = *(*int32)(unsafe.Add(mBase, uint32(v3662)+28))
	v3900 = v3777
	goto L831
L859:
	;
	goto L860
L860:
	;
	v3778 = F_strlen(m, v3647)
	mBase = m.M
	v3780 = v3662 + int32(32)
	v3781 = F_strlen(m, v3780)
	mBase = m.M
	if v3778 != v3781 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3662))) = v3657
	v3875 = F_psprintf(m, int32(_a_F_InitPostgres_80), v3662)
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L1
	} else {
		goto L880
	}
L862:
	;
	v3783 = int32(0)
	if v3778 == v3783 {
		goto L864
	} else {
		goto L865
	}
L863:
	;
	if v3871 != 0 {
		goto L861
	} else {
		goto L879
	}
L864:
	;
	v3871 = int32(0)
	goto L863
L865:
	;
	goto L866
L866:
	;
	v3791 = v3778 & int32(3)
	if base.Ui32(v3778) < base.Ui32(int32(4)) {
		goto L869
	} else {
		goto L870
	}
L867:
	;
	v3871 = base.B2i32(v3857 != int32(0))
	goto L863
L868:
	;
	v3837 = v3830
	v3838 = v3831
	v3839 = v3832
	v3843 = v3783
	goto L876
L869:
	;
	v3830 = v3647
	v3831 = v3780
	v3832 = int32(0)
	goto L868
L870:
	;
	goto L871
L871:
	;
	v3798 = v3647
	v3799 = v3780
	v3800 = int32(0)
	v3803 = v3783
	goto L872
L872:
	;
	v3805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3799))))
	v3806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3798))))
	v3809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3799)+1)))
	v3810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3798)+1)))
	v3813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3799)+2)))
	v3814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3798)+2)))
	v3817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3799)+3)))
	v3818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3798)+3)))
	v3820 = v3800 | (v3805 ^ v3806) | (v3809 ^ v3810) | (v3813 ^ v3814) | (v3817 ^ v3818)
	v3821 = int32(4)
	v3822 = v3799 + v3821
	v3824 = v3798 + v3821
	v3826 = v3803 + v3821
	if v3826 != v3778&int32(-4) {
		v3798 = v3824
		v3799 = v3822
		v3800 = v3820
		v3803 = v3826
		goto L872
	} else {
		goto L874
	}
L873:
	;
	if v3791 == int32(0) {
		v3857 = v3820
		goto L867
	} else {
		goto L875
	}
L874:
	;
	goto L873
L875:
	;
	v3830 = v3824
	v3831 = v3822
	v3832 = v3820
	goto L868
L876:
	;
	v3844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3838))))
	v3845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3837))))
	v3847 = v3839 | (v3844 ^ v3845)
	v3848 = int32(1)
	v3853 = v3843 + v3848
	if v3853 != v3791 {
		v3837 = v3837 + v3848
		v3838 = v3838 + v3848
		v3839 = v3847
		v3843 = v3853
		goto L876
	} else {
		goto L878
	}
L877:
	;
	v3857 = v3847
	goto L867
L878:
	;
	goto L877
L879:
	;
	v3905 = int32(0)
	goto L830
L880:
	;
	v3900 = v3875
	goto L831
L881:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3662)+16)) = v3657
	v3896 = F_psprintf(m, int32(_a_F_InitPostgres_81), v3662+int32(16))
	mBase = m.M
	v3897 = m.ExcPending
	if v3897 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	v3900 = v3896
	goto L831
L883:
	;
	if v3905 != 0 {
		v3935 = v3905
		goto L797
	} else {
		goto L884
	}
L884:
	;
	v3912 = int32(_a_F_InitPostgres_8)
	v3913 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v3916 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v3916
	v3919 = F_pstrdup(m, int32(_a_F_InitPostgres_82))
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	v3922 = F_pstrdup(m, int32(_a_F_InitPostgres_83))
	mBase = m.M
	v3923 = m.ExcPending
	if v3923 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	F_StoreConnectionWarning(m, v3919, v3922, int32(845))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v3913
	v3935 = int32(0)
	goto L797
L888:
	;
	v3935 = v3933
	goto L797
L889:
	;
	F_pfree(m, v3552)
	mBase = m.M
	v3941 = m.ExcPending
	if v3941 != 0 {
		goto L1
	} else {
		goto L892
	}
L890:
	;
	goto L891
L891:
	;
	if v3935 != 0 {
		v4043 = v3935
		goto L530
	} else {
		goto L893
	}
L892:
	;
	goto L891
L893:
	;
	v3942 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	F_set_authn_id(m, v939, v3942)
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L1
	} else {
		goto L894
	}
L894:
	;
	v4043 = int32(0)
	goto L530
L895:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3949 = m.ExcPending
	if v3949 != 0 {
		goto L1
	} else {
		goto L898
	}
L896:
	;
	goto L897
L897:
	;
	v3951 = v989 + int32(448)
	F_pq_beginmessage(m, v3951, int32(82))
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L899
	}
L898:
	;
	goto L897
L899:
	;
	F_enlargeStringInfo(m, v3951, int32(4))
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v3958 = *(*int32)(unsafe.Add(mBase, uint32(v989)+452))
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v989)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v3958+v3959))) = int32(50331648)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+452)) = v3958 + int32(4)
	F_pq_endmessage(m, v3951)
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	v3969 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[61]))
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3969)+4))
	v3971 = m.T0[v3970].(func(*base.Module) int32)(m)
	mBase = m.M
	v3972 = m.ExcPending
	if v3972 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	v3974 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v3974 != 0 {
		goto L903
	} else {
		goto L904
	}
L903:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L1
	} else {
		goto L906
	}
L904:
	;
	goto L905
L905:
	;
	v3977 = F_recv_password_packet(m)
	mBase = m.M
	v3978 = m.ExcPending
	if v3978 != 0 {
		goto L1
	} else {
		goto L907
	}
L906:
	;
	goto L905
L907:
	;
	if v3977 == int32(0) {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v4043 = int32(-2)
	goto L530
L909:
	;
	goto L910
L910:
	;
	v3982 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v3985 = F_get_role_password(m, v3982, v989+int32(444))
	mBase = m.M
	v3986 = m.ExcPending
	if v3986 != 0 {
		goto L1
	} else {
		goto L911
	}
L911:
	;
	if v3985 == int32(0) {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	F_pfree(m, v3977)
	mBase = m.M
	v3990 = m.ExcPending
	if v3990 != 0 {
		goto L1
	} else {
		goto L915
	}
L913:
	;
	goto L914
L914:
	;
	v3991 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v3994 = F_plain_crypt_verify(m, v3991, v3985, v3977, v989+int32(444))
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L1
	} else {
		goto L916
	}
L915:
	;
	v4043 = v2318
	goto L530
L916:
	;
	v3996 = F_get_password_type(m, v3985)
	mBase = m.M
	v3997 = m.ExcPending
	if v3997 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	F_pfree(m, v3985)
	mBase = m.M
	v3999 = m.ExcPending
	if v3999 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	F_pfree(m, v3977)
	mBase = m.M
	v4001 = m.ExcPending
	if v4001 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	if v3994 != 0 {
		v4043 = v3994
		goto L530
	} else {
		goto L920
	}
L920:
	;
	if v3996 == int32(1) {
		goto L921
	} else {
		goto L922
	}
L921:
	;
	v4004 = int32(_a_F_InitPostgres_8)
	v4005 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v4008 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v4008
	v4011 = F_pstrdup(m, int32(_a_F_InitPostgres_82))
	mBase = m.M
	v4012 = m.ExcPending
	if v4012 != 0 {
		goto L1
	} else {
		goto L924
	}
L922:
	;
	goto L923
L923:
	;
	v4022 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	F_set_authn_id(m, v939, v4022)
	mBase = m.M
	v4024 = m.ExcPending
	if v4024 != 0 {
		goto L1
	} else {
		goto L927
	}
L924:
	;
	v4014 = F_pstrdup(m, int32(_a_F_InitPostgres_83))
	mBase = m.M
	v4015 = m.ExcPending
	if v4015 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	F_StoreConnectionWarning(m, v4011, v4014, int32(845))
	mBase = m.M
	v4018 = m.ExcPending
	if v4018 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v4005
	goto L923
L927:
	;
	goto L532
L928:
	;
	v4043 = v4035
	goto L530
L929:
	;
	v4102 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[62]))
	if v4102 != 0 {
		goto L937
	} else {
		goto L938
	}
L930:
	;
	v4070 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[63]))
	if v4070 != 0 {
		goto L929
	} else {
		goto L931
	}
L931:
	;
	v4073 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	if v4073 == int32(0) {
		goto L929
	} else {
		goto L933
	}
L933:
	;
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	v4078 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v4078)+296))
	v4082 = *(*int32)(unsafe.Add(mBase, uint32(v4079<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[64])))
	goto L934
L934:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v4084 = *(*int64)(unsafe.Add(mBase, uint32(v4083)))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+68)) = v4082
	*(*int64)(unsafe.Add(mBase, uint32(v989)+72)) = v4084
	*(*int32)(unsafe.Add(mBase, uint32(v989)+64)) = v4077
	F_errmsg(m, int32(_a_F_InitPostgres_84), v989-int32(-64))
	mBase = m.M
	v4092 = m.ExcPending
	if v4092 != 0 {
		goto L1
	} else {
		goto L935
	}
L935:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(664), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v4097 = m.ExcPending
	if v4097 != 0 {
		goto L1
	} else {
		goto L936
	}
L936:
	;
	goto L929
L937:
	;
	m.T0[v4102].(func(*base.Module, int32, int32))(m, v939, v4043)
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L1
	} else {
		goto L940
	}
L938:
	;
	goto L939
L939:
	;
	if v4043 == int32(0) {
		goto L941
	} else {
		goto L942
	}
L940:
	;
	goto L939
L941:
	;
	v4108 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4108 != 0 {
		goto L944
	} else {
		goto L945
	}
L942:
	;
	goto L943
L943:
	;
	if v4043 == int32(-2) {
		goto L525
	} else {
		goto L955
	}
L944:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4110 = m.ExcPending
	if v4110 != 0 {
		goto L1
	} else {
		goto L947
	}
L945:
	;
	goto L946
L946:
	;
	v4112 = v989 + int32(448)
	F_pq_beginmessage(m, v4112, int32(82))
	mBase = m.M
	v4115 = m.ExcPending
	if v4115 != 0 {
		goto L1
	} else {
		goto L948
	}
L947:
	;
	goto L946
L948:
	;
	F_enlargeStringInfo(m, v4112, int32(4))
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v989)+452))
	v4120 = *(*int32)(unsafe.Add(mBase, uint32(v989)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v4119+v4120))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+452)) = v4119 + int32(4)
	F_pq_endmessage(m, v4112)
	mBase = m.M
	v4128 = m.ExcPending
	if v4128 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	v4130 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4130 != 0 {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4132 = m.ExcPending
	if v4132 != 0 {
		goto L1
	} else {
		goto L954
	}
L952:
	;
	goto L953
L953:
	;
	m.G0 = v989 + int32(2992)
	goto L524
L954:
	;
	goto L953
L955:
	;
	v4138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+443)))
	v4139 = *(*int32)(unsafe.Add(mBase, uint32(v989)+444))
	v4140 = *(*int32)(unsafe.Add(mBase, uint32(v939)+380))
	v4141 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+296))
	if base.Ui32(int32(14)) < base.Ui32(v4141) {
		goto L957
	} else {
		goto L958
	}
L956:
	;
	if v4138 != 0 {
		goto L960
	} else {
		goto L961
	}
L957:
	;
	v4151 = int32(514)
	v4152 = int32(_a_F_InitPostgres_85)
	goto L956
L958:
	;
	goto L959
L959:
	;
	v4147 = v4141 << (uint(int32(2)) % 32)
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v4147)+uint32(_c_F_InitPostgres[65])))
	v4149 = *(*int32)(unsafe.Add(mBase, uint32(v4147)+uint32(_c_F_InitPostgres[66])))
	v4151 = v4148
	v4152 = v4149
	goto L956
L960:
	;
	v4155 = int32(23)
	goto L962
L961:
	;
	v4155 = int32(22)
	goto L962
L962:
	;
	v4156 = *(*int64)(unsafe.Add(mBase, uint32(v4140)))
	v4157 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+56)) = v4157
	*(*int64)(unsafe.Add(mBase, uint32(v989)+48)) = v4156
	v4163 = F_psprintf(m, int32(_a_F_InitPostgres_86), v989+int32(48))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L963
	}
L963:
	;
	if v4139 != 0 {
		goto L964
	} else {
		goto L965
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+36)) = v4163
	*(*int32)(unsafe.Add(mBase, uint32(v989)+32)) = v4139
	v4170 = F_psprintf(m, int32(_a_F_InitPostgres_87), v989+int32(32))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L1
	} else {
		goto L967
	}
L965:
	;
	v4172 = v4163
	goto L966
L966:
	;
	v4174 = F_errstart(m, v4155, int32(0))
	mBase = m.M
	v4175 = m.ExcPending
	if v4175 != 0 {
		goto L1
	} else {
		goto L968
	}
L967:
	;
	v4172 = v4170
	goto L966
L968:
	;
	F_errcode(m, v4151)
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	v4178 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+16)) = v4178
	F_errmsg(m, v4152, v989+int32(16))
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	if v4172 != 0 {
		goto L971
	} else {
		goto L972
	}
L971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989))) = v4172
	F_errdetail_log(m, int32(_a_F_InitPostgres_88), v989)
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L974
	}
L972:
	;
	goto L973
L973:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(316), int32(_a_F_InitPostgres_89))
	mBase = m.M
	v4192 = m.ExcPending
	if v4192 != 0 {
		goto L1
	} else {
		goto L975
	}
L974:
	;
	goto L973
L975:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L976:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v4199 = m.ExcPending
	if v4199 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_90), int32(0))
	mBase = m.M
	v4203 = m.ExcPending
	if v4203 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(411), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v4208 = m.ExcPending
	if v4208 != 0 {
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
	F_errfinish(m, int32(_a_F_InitPostgres_53), int32(466), int32(_a_F_InitPostgres_54))
	mBase = m.M
	v4224 = m.ExcPending
	if v4224 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L982:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L983:
	;
	v4235 = m.G0
	v4236 = int32(16)
	v4237 = v4235 - v4236
	m.G0 = v4237
	F_gettimeofday(m, v4237)
	mBase = m.M
	v4240 = *(*int64)(unsafe.Add(mBase, uint32(v4237)))
	v4241 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4237)+8)))
	m.G0 = v4237 + v4236
	goto L984
L984:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[67])) = v4241 + v4240*int64(1000000) - int64(946684800000000)
	v4252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[57])))
	if v4252&int32(4) != 0 {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v4256 = v28 + int32(448)
	F_initStringInfo(m, v4256)
	mBase = m.M
	v4258 = m.ExcPending
	if v4258 != 0 {
		goto L1
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v4312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[34])) = uint8(v4312)
	F_InitializeSessionUserId(m, l2, l3, v4312)
	mBase = m.M
	v4316 = m.ExcPending
	if v4316 != 0 {
		goto L1
	} else {
		goto L1008
	}
L988:
	;
	v4259 = *(*int32)(unsafe.Add(mBase, uint32(v939)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+400)) = v4259
	v4264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v4264 != 0 {
		goto L989
	} else {
		goto L990
	}
L989:
	;
	v4265 = int32(_a_F_InitPostgres_91)
	goto L991
L990:
	;
	v4265 = int32(_a_F_InitPostgres_92)
	goto L991
L991:
	;
	F_appendStringInfo(m, v4256, v4265, v28+int32(400))
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	v4271 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v4271 == int32(0) {
		goto L993
	} else {
		goto L994
	}
L993:
	;
	v4274 = *(*int32)(unsafe.Add(mBase, uint32(v939)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+384)) = v4274
	F_appendStringInfo(m, v4256, int32(_a_F_InitPostgres_93), v28+int32(384))
	mBase = m.M
	v4280 = m.ExcPending
	if v4280 != 0 {
		goto L1
	} else {
		goto L996
	}
L994:
	;
	goto L995
L995:
	;
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(v939)+376))
	if v4281 != 0 {
		goto L997
	} else {
		goto L998
	}
L996:
	;
	goto L995
L997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+368)) = v4281
	F_appendStringInfo(m, v28+int32(448), int32(_a_F_InitPostgres_94), v28+int32(368))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L1000
	}
L998:
	;
	goto L999
L999:
	;
	v4292 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4293 = m.ExcPending
	if v4293 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1000:
	;
	goto L999
L1001:
	;
	if v4292 != 0 {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	v4294 = *(*int32)(unsafe.Add(mBase, uint32(v28)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+352)) = v4294
	F_errmsg_internal(m, int32(_a_F_InitPostgres_88), v28+int32(352))
	mBase = m.M
	v4300 = m.ExcPending
	if v4300 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1003:
	;
	goto L1004
L1004:
	;
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(v28)+448))
	F_pfree(m, v4306)
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1005:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(324), int32(_a_F_InitPostgres_95))
	mBase = m.M
	v4305 = m.ExcPending
	if v4305 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1006:
	;
	goto L1004
L1007:
	;
	goto L987
L1008:
	;
	v4318 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[63]))
	if v4318 == int32(0) {
		v4328 = l0
		v4329 = l1
		v4332 = l4
		v4333 = l5
		v4340 = v28
		v4349 = v31
		v4351 = v7
		goto L150
	} else {
		goto L1009
	}
L1009:
	;
	v4322 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[68]))
	v4325 = *(*int32)(unsafe.Add(mBase, uint32(v4322<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[64])))
	goto L1010
L1010:
	;
	F_InitializeSystemUser(m, v4318, v4325)
	mBase = m.M
	v4327 = m.ExcPending
	if v4327 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	v4328 = l0
	v4329 = l1
	v4332 = l4
	v4333 = l5
	v4340 = v28
	v4349 = v31
	v4351 = v7
	goto L150
L1012:
	;
	v4355 = v4328
	v4356 = v4329
	v4358 = v4353
	v4359 = v4332
	v4360 = v4333
	v4367 = v4340
	v4376 = v4349
	v4378 = v4351
	goto L149
L1013:
	;
	v4382 = int32(_a_F_InitPostgres_96)
	v4384 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[69]))
	v4385 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[69])) = v4384 + v4385
	v4389 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[70]))
	v4390 = *(*int32)(unsafe.Add(mBase, uint32(v4389)))
	*(*int32)(unsafe.Add(mBase, uint32(v4389))) = v4390 + v4385
	v4394 = int32(0)
	v4396 = int32(_a_F_InitPostgres_97)
	v4397 = base.AtomicRmwOr32(m, v4394, v4396, v4394)
	*(*uint8)(unsafe.Add(mBase, uint32(v4389)+192)) = uint8(v4394)
	*(*uint8)(unsafe.Add(mBase, uint32(v4389)+200)) = uint8(v4394)
	v4405 = base.AtomicRmwOr32(m, v4394, v4396, v4394)
	v4406 = *(*int32)(unsafe.Add(mBase, uint32(v4389)))
	*(*int32)(unsafe.Add(mBase, uint32(v4389))) = v4406 + v4385
	v4412 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[69])) = v4412 - v4385
	goto L1015
L1014:
	;
	goto L1015
L1015:
	;
	v4418 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[71])))
	if (v4418^int32(-1)|v4358)&int32(1) == int32(0) {
		goto L147
	} else {
		goto L1016
	}
L1016:
	;
	v4427 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v4428 = int32(1)
	if (base.B2i32(v4427 != v4428)|v4358)&v4428 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	v4563 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v4563 != int32(1) {
		goto L1036
	} else {
		goto L1037
	}
L1018:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[72]))
	v4436 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[73]))
	v4437 = v4434 + v4436
	if v4437 <= int32(0) {
		goto L1017
	} else {
		goto L1019
	}
L1019:
	;
	v4441 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	v4444 = base.AtomicRmwXchg32(m, v4441, int32(20), int32(1))
	if v4444 != 0 {
		goto L1020
	} else {
		goto L1021
	}
L1020:
	;
	F_s_lock(m, v4441+int32(20), int32(_a_F_InitPostgres_98))
	mBase = m.M
	v4449 = m.ExcPending
	if v4449 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1021:
	;
	goto L1022
L1022:
	;
	v4451 = v4367 + int32(444)
	v4452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4451))) = v4452
	v4455 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	v4456 = *(*int32)(unsafe.Add(mBase, uint32(v4455)+28))
	if v4456 == v4452 {
		v4503 = v4455
		goto L1024
	} else {
		goto L1025
	}
L1023:
	;
	goto L1022
L1024:
	;
	v4521 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4503)+20)), uint32(v4521))
	v4524 = *(*int32)(unsafe.Add(mBase, uint32(v4451)))
	if v4524 == v4437 {
		goto L1017
	} else {
		goto L1032
	}
L1025:
	;
	v4460 = v4455 + int32(24)
	if v4456 == v4460 {
		v4503 = v4455
		goto L1024
	} else {
		goto L1026
	}
L1026:
	;
	v4468 = v4456
	v4485 = v4378
	goto L1027
L1027:
	;
	v4488 = v4485 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4451))) = v4488
	if v4437 == v4488 {
		goto L1029
	} else {
		goto L1030
	}
L1028:
	;
	v4495 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	v4503 = v4495
	goto L1024
L1029:
	;
	goto L1028
L1030:
	;
	v4491 = *(*int32)(unsafe.Add(mBase, uint32(v4468)+4))
	if v4491 != v4460 {
		v4468 = v4491
		v4485 = v4488
		goto L1027
	} else {
		goto L1031
	}
L1031:
	;
	goto L1029
L1032:
	;
	v4526 = *(*int32)(unsafe.Add(mBase, uint32(v4367)+444))
	v4528 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[73]))
	if v4526 < v4528 {
		goto L146
	} else {
		goto L1033
	}
L1033:
	;
	v4531 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	v4533 = F_has_privs_of_role(m, v4531, int32(_a_F_InitPostgres_99))
	mBase = m.M
	v4534 = m.ExcPending
	if v4534 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	if v4533 == int32(0) {
		goto L145
	} else {
		goto L1035
	}
L1035:
	;
	goto L1017
L1036:
	;
	if v4376 != 0 {
		goto L1050
	} else {
		goto L1051
	}
L1037:
	;
	v4567 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	v4568 = F_has_rolreplication(m, v4567)
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	if v4568 == int32(0) {
		goto L144
	} else {
		goto L1039
	}
L1039:
	;
	v4573 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v4573 != int32(1) {
		goto L1036
	} else {
		goto L1040
	}
L1040:
	;
	v4577 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	if v4577&int32(1) != 0 {
		goto L1036
	} else {
		goto L1041
	}
L1041:
	;
	v4581 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v4581 != 0 {
		goto L1042
	} else {
		goto L1043
	}
L1042:
	;
	F_process_startup_options(m, v4581, v4358)
	mBase = m.M
	v4583 = m.ExcPending
	if v4583 != 0 {
		goto L1
	} else {
		goto L1045
	}
L1043:
	;
	goto L1044
L1044:
	;
	v4585 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	if int32(0) < v4585 {
		goto L1046
	} else {
		goto L1047
	}
L1045:
	;
	goto L1044
L1046:
	;
	F_pg_usleep(m, v4585*int32(_a_F_InitPostgres_100))
	mBase = m.M
	goto L1048
L1047:
	;
	goto L1048
L1048:
	;
	F_InitializeClientEncoding(m)
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1049:
	;
	goto L140
L1050:
	;
	if v4355 != 0 {
		goto L1053
	} else {
		goto L1054
	}
L1051:
	;
	goto L1052
L1052:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77])) = int32(1663)
	v5038 = int32(1)
	goto L141
L1053:
	;
	v4594 = v4367 + int32(448)
	F_ScanKeyInit(m, v4594, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v4355))
	mBase = m.M
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1054:
	;
	goto L1055
L1055:
	;
	if v4356 != 0 {
		v4749 = v4356
		goto L142
	} else {
		goto L1067
	}
L1056:
	;
	v4604 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v4605 = m.ExcPending
	if v4605 != 0 {
		goto L1
	} else {
		goto L1057
	}
L1057:
	;
	v4608 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[78])))
	v4611 = F_systable_beginscan(m, v4604, int32(2671), v4608, int32(0), int32(1), v4594)
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1058:
	;
	v4613 = F_systable_getnext(m, v4611)
	mBase = m.M
	v4614 = m.ExcPending
	if v4614 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1059:
	;
	if v4613 != 0 {
		goto L1060
	} else {
		goto L1061
	}
L1060:
	;
	v4615 = F_heap_copytuple(m, v4613)
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1061:
	;
	v4617 = int32(0)
	goto L1062
L1062:
	;
	F_systable_endscan(m, v4611)
	mBase = m.M
	v4619 = m.ExcPending
	if v4619 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1063:
	;
	v4617 = v4615
	goto L1062
L1064:
	;
	F_relation_close(m, v4604, int32(1))
	mBase = m.M
	v4622 = m.ExcPending
	if v4622 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	if v4617 == int32(0) {
		goto L143
	} else {
		goto L1066
	}
L1066:
	;
	v4625 = *(*int32)(unsafe.Add(mBase, uint32(v4617)+16))
	v4626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4625)+22)))
	v4628 = *(*int32)(unsafe.Add(mBase, uint32(v4625+v4626)))
	v4749 = v4628
	goto L142
L1067:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v4630 = m.ExcPending
	if v4630 != 0 {
		goto L1
	} else {
		goto L1068
	}
L1068:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v4632 = m.ExcPending
	if v4632 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1069:
	;
	v5975 = v4367
	goto L134
L1070:
	;
	v4642 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[79]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+416)) = v4642
	F_errmsg(m, int32(_a_F_InitPostgres_101), v28+int32(416))
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1071:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(240), int32(_a_F_InitPostgres_95))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1073:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1074:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_102), int32(0))
	mBase = m.M
	v4664 = m.ExcPending
	if v4664 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1075:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(961), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1076:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1077:
	;
	F_errcode(m, int32(_a_F_InitPostgres_103))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+320)) = int32(_a_F_InitPostgres_104)
	F_errmsg(m, int32(_a_F_InitPostgres_105), v4367+int32(320))
	mBase = m.M
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(982), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L1
	} else {
		goto L1080
	}
L1080:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1081:
	;
	F_errcode(m, int32(_a_F_InitPostgres_103))
	mBase = m.M
	v4695 = m.ExcPending
	if v4695 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+336)) = int32(_a_F_InitPostgres_106)
	F_errmsg(m, int32(_a_F_InitPostgres_107), v4367+int32(336))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(988), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4707 = m.ExcPending
	if v4707 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1084:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1085:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v4714 = m.ExcPending
	if v4714 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_108), int32(0))
	mBase = m.M
	v4718 = m.ExcPending
	if v4718 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+304)) = int32(_a_F_InitPostgres_109)
	v4724 = F_errdetail(m, int32(_a_F_InitPostgres_110), v4367+int32(304))
	mBase = m.M
	v4725 = m.ExcPending
	if v4725 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1001), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4730 = m.ExcPending
	if v4730 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1090:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v4737 = m.ExcPending
	if v4737 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+288)) = v4355
	F_errmsg(m, int32(_a_F_InitPostgres_111), v4367+int32(288))
	mBase = m.M
	v4743 = m.ExcPending
	if v4743 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1056), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4748 = m.ExcPending
	if v4748 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1094:
	;
	v4759 = v4367 + int32(448)
	F_ScanKeyInit(m, v4759, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v4749))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	v4768 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v4769 = m.ExcPending
	if v4769 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	v4772 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[78])))
	v4775 = F_systable_beginscan(m, v4768, int32(2672), v4772, int32(0), int32(1), v4759)
	mBase = m.M
	v4776 = m.ExcPending
	if v4776 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	v4777 = F_systable_getnext(m, v4775)
	mBase = m.M
	v4778 = m.ExcPending
	if v4778 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	if v4777 != 0 {
		goto L1099
	} else {
		goto L1100
	}
L1099:
	;
	v4779 = F_heap_copytuple(m, v4777)
	mBase = m.M
	v4780 = m.ExcPending
	if v4780 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1100:
	;
	v4781 = int32(0)
	goto L1101
L1101:
	;
	F_systable_endscan(m, v4775)
	mBase = m.M
	v4783 = m.ExcPending
	if v4783 != 0 {
		goto L1
	} else {
		goto L1103
	}
L1102:
	;
	v4781 = v4779
	goto L1101
L1103:
	;
	F_relation_close(m, v4768, int32(1))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1104:
	;
	if v4781 != 0 {
		goto L1106
	} else {
		goto L1107
	}
L1105:
	;
	v4830 = v4367 + int32(448)
	v4832 = v4789 + int32(4)
	goto L1129
L1106:
	;
	v4787 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+16))
	v4788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4787)+22)))
	v4789 = v4787 + v4788
	if v4355 == int32(0) {
		goto L1105
	} else {
		goto L1109
	}
L1107:
	;
	goto L1108
L1108:
	;
	if v4355 != 0 {
		goto L124
	} else {
		goto L1121
	}
L1109:
	;
	v4793 = v4789 + int32(4)
	if v4793|v4355 != 0 {
		goto L1111
	} else {
		goto L1112
	}
L1110:
	;
	if v4808 == int32(0) {
		goto L1105
	} else {
		goto L1120
	}
L1111:
	;
	v4799 = int32(-1)
	goto L1113
L1112:
	;
	v4799 = int32(0)
	goto L1113
L1113:
	;
	if v4793 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	v4800 = int32(1)
	goto L1116
L1115:
	;
	v4800 = v4799
	goto L1116
L1116:
	;
	v4801 = int32(0)
	if base.B2i32(v4793 == v4801)|base.B2i32(v4355 == v4801) != 0 {
		goto L1117
	} else {
		goto L1118
	}
L1117:
	;
	v4808 = v4800
	goto L1119
L1118:
	;
	v4807 = F_strncmp(m, v4793, v4355, int32(64))
	mBase = m.M
	v4808 = v4807
	goto L1119
L1119:
	;
	goto L1110
L1120:
	;
	goto L124
L1121:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4814 = m.ExcPending
	if v4814 != 0 {
		goto L1
	} else {
		goto L1122
	}
L1122:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+240)) = v4749
	F_errmsg(m, int32(_a_F_InitPostgres_112), v4367+int32(240))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1124:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1125), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		goto L1
	} else {
		goto L1125
	}
L1125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1126:
	;
	v4952 = *(*int32)(unsafe.Add(mBase, uint32(v4789)+80))
	goto L1157
L1127:
	;
	v4949 = F_strlen(m, v4938)
	mBase = m.M
	goto L1126
L1129:
	;
	goto L1130
L1130:
	;
	v4839 = int32(63)
	if (v4830^v4832)&int32(3) != 0 {
		goto L1134
	} else {
		goto L1135
	}
L1131:
	;
	v4942 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4939))) = uint8(v4942)
	goto L1127
L1132:
	;
	v4923 = v4918
	v4924 = v4919
	v4925 = v4920
	goto L1153
L1133:
	;
	if v4913 == int32(0) {
		v4938 = v4911
		v4939 = v4912
		goto L1131
	} else {
		goto L1152
	}
L1134:
	;
	v4911 = v4832
	v4912 = v4830
	v4913 = v4839
	goto L1133
L1135:
	;
	goto L1136
L1136:
	;
	v4843 = int32(0)
	if base.B2i32(v4832&int32(3) == v4843)|int32(0) == v4843 {
		goto L1138
	} else {
		goto L1139
	}
L1137:
	;
	if v4879 == int32(0) {
		v4938 = v4876
		v4939 = v4877
		goto L1131
	} else {
		goto L1146
	}
L1138:
	;
	v4855 = v4832
	v4856 = v4830
	v4857 = v4839
	goto L1141
L1139:
	;
	goto L1140
L1140:
	;
	v4876 = v4832
	v4877 = v4830
	v4878 = v4839
	v4879 = int32(1)
	goto L1137
L1141:
	;
	v4859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4855))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4856))) = uint8(v4859)
	if v4859 == int32(0) {
		v4918 = v4855
		v4919 = v4856
		v4920 = v4857
		goto L1132
	} else {
		goto L1143
	}
L1142:
	;
	v4876 = v4870
	v4877 = v4864
	v4878 = v4866
	v4879 = v4868
	goto L1137
L1143:
	;
	v4863 = int32(1)
	v4864 = v4856 + v4863
	v4866 = v4857 - v4863
	v4867 = int32(0)
	v4868 = base.B2i32(v4866 != v4867)
	v4870 = v4855 + v4863
	if v4870&int32(3) == v4867 {
		v4876 = v4870
		v4877 = v4864
		v4878 = v4866
		v4879 = v4868
		goto L1137
	} else {
		goto L1144
	}
L1144:
	;
	if v4866 != 0 {
		v4855 = v4870
		v4856 = v4864
		v4857 = v4866
		goto L1141
	} else {
		goto L1145
	}
L1145:
	;
	goto L1142
L1146:
	;
	v4882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4876))))
	if base.B2i32(v4882 == int32(0))|base.B2i32(base.Ui32(v4878) < base.Ui32(int32(4))) != 0 {
		v4911 = v4876
		v4912 = v4877
		v4913 = v4878
		goto L1133
	} else {
		goto L1147
	}
L1147:
	;
	v4889 = v4876
	v4890 = v4877
	v4891 = v4878
	goto L1148
L1148:
	;
	v4894 = *(*int32)(unsafe.Add(mBase, uint32(v4889)))
	v4897 = int32(-2139062144)
	if (int32(16843008)-v4894|v4894)&v4897 != v4897 {
		v4918 = v4889
		v4919 = v4890
		v4920 = v4891
		goto L1132
	} else {
		goto L1150
	}
L1149:
	;
	v4911 = v4905
	v4912 = v4903
	v4913 = v4907
	goto L1133
L1150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4890))) = v4894
	v4902 = int32(4)
	v4903 = v4890 + v4902
	v4905 = v4889 + v4902
	v4907 = v4891 - v4902
	if base.Ui32(int32(3)) < base.Ui32(v4907) {
		v4889 = v4905
		v4890 = v4903
		v4891 = v4907
		goto L1148
	} else {
		goto L1151
	}
L1151:
	;
	goto L1149
L1152:
	;
	v4918 = v4911
	v4919 = v4912
	v4920 = v4913
	goto L1132
L1153:
	;
	v4927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4923))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4924))) = uint8(v4927)
	if v4927 == int32(0) {
		v4938 = v4923
		v4939 = v4924
		goto L1131
	} else {
		goto L1155
	}
L1154:
	;
	v4938 = v4934
	v4939 = v4932
	goto L1131
L1155:
	;
	v4931 = int32(1)
	v4932 = v4924 + v4931
	v4934 = v4923 + v4931
	v4936 = v4925 - v4931
	if v4936 != 0 {
		v4923 = v4934
		v4924 = v4932
		v4925 = v4936
		goto L1153
	} else {
		goto L1156
	}
L1156:
	;
	goto L1154
L1157:
	;
	if v4952 == int32(-2) {
		goto L133
	} else {
		goto L1158
	}
L1158:
	;
	v4956 = *(*int32)(unsafe.Add(mBase, uint32(v4789)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77])) = v4956
	v4959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4789)+79)))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[80])) = uint8(v4959)
	if v4360 == int32(0) {
		v5038 = v4749
		goto L141
	} else {
		goto L1159
	}
L1159:
	;
	if (v4830^v4360)&int32(3) != 0 {
		goto L1163
	} else {
		goto L1164
	}
L1160:
	;
	v5038 = v4749
	goto L141
L1161:
	;
	goto L1160
L1162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v5017))) = uint8(v5016)
	if v5016&int32(255) == int32(0) {
		goto L1161
	} else {
		goto L1177
	}
L1163:
	;
	v4968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4830))))
	v5015 = v4830
	v5016 = v4968
	v5017 = v4360
	goto L1162
L1164:
	;
	goto L1165
L1165:
	;
	if v4830&int32(3) != 0 {
		goto L1166
	} else {
		goto L1167
	}
L1166:
	;
	v4972 = v4830
	v4974 = v4360
	goto L1169
L1167:
	;
	v4986 = v4830
	v4988 = v4360
	goto L1168
L1168:
	;
	v4990 = *(*int32)(unsafe.Add(mBase, uint32(v4986)))
	v4993 = int32(-2139062144)
	if (int32(16843008)-v4990|v4990)&v4993 != v4993 {
		v5015 = v4986
		v5016 = v4990
		v5017 = v4988
		goto L1162
	} else {
		goto L1173
	}
L1169:
	;
	v4975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4972))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4974))) = uint8(v4975)
	if v4975 == int32(0) {
		goto L1161
	} else {
		goto L1171
	}
L1170:
	;
	v4986 = v4982
	v4988 = v4980
	goto L1168
L1171:
	;
	v4979 = int32(1)
	v4980 = v4974 + v4979
	v4982 = v4972 + v4979
	if v4982&int32(3) != 0 {
		v4972 = v4982
		v4974 = v4980
		goto L1169
	} else {
		goto L1172
	}
L1172:
	;
	goto L1170
L1173:
	;
	v4998 = v4986
	v4999 = v4990
	v5000 = v4988
	goto L1174
L1174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5000))) = v4999
	v5002 = int32(4)
	v5003 = v5000 + v5002
	v5005 = v4998 + v5002
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v4998)+4))
	v5010 = int32(-2139062144)
	if (int32(16843008)-v5007|v5007)&v5010 == v5010 {
		v4998 = v5005
		v4999 = v5007
		v5000 = v5003
		goto L1174
	} else {
		goto L1176
	}
L1175:
	;
	v5015 = v5005
	v5016 = v5007
	v5017 = v5003
	goto L1162
L1176:
	;
	goto L1175
L1177:
	;
	v5024 = v5015
	v5026 = v5017
	goto L1178
L1178:
	;
	v5027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5024)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5026)+1)) = uint8(v5027)
	v5029 = int32(1)
	if v5027 != 0 {
		v5024 = v5024 + v5029
		v5026 = v5026 + v5029
		goto L1178
	} else {
		goto L1180
	}
L1179:
	;
	goto L1161
L1180:
	;
	goto L1179
L1181:
	;
	v5051 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	v5053 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	v5054 = F_GetDatabasePath(m, v5051, v5053)
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L1
	} else {
		goto L1182
	}
L1182:
	;
	if v4376 != 0 {
		goto L1184
	} else {
		goto L1185
	}
L1183:
	;
	v5583 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v5583 != 0 {
		goto L1313
	} else {
		goto L1314
	}
L1184:
	;
	v5057 = F_access(m, v5054, int32(0))
	mBase = m.M
	if v5057 == int32(-1) {
		goto L1187
	} else {
		goto L1188
	}
L1185:
	;
	goto L1186
L1186:
	;
	F_SetDatabasePath(m, v5054)
	mBase = m.M
	v5550 = m.ExcPending
	if v5550 != 0 {
		goto L1
	} else {
		goto L1309
	}
L1187:
	;
	v5061 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v5065 = m.ExcPending
	if v5065 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1188:
	;
	goto L1189
L1189:
	;
	F_ValidatePgVersion(m, v5054)
	mBase = m.M
	v5082 = m.ExcPending
	if v5082 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1190:
	;
	if v5061 == int32(44) {
		goto L132
	} else {
		goto L1191
	}
L1191:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v5069 = m.ExcPending
	if v5069 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+32)) = v5054
	F_errmsg(m, int32(_a_F_InitPostgres_113), v4367+int32(32))
	mBase = m.M
	v5075 = m.ExcPending
	if v5075 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1201), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v5080 = m.ExcPending
	if v5080 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1195:
	;
	F_SetDatabasePath(m, v5054)
	mBase = m.M
	v5084 = m.ExcPending
	if v5084 != 0 {
		goto L1
	} else {
		goto L1196
	}
L1196:
	;
	F_pfree(m, v5054)
	mBase = m.M
	v5086 = m.ExcPending
	if v5086 != 0 {
		goto L1
	} else {
		goto L1197
	}
L1197:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v5088 = m.ExcPending
	if v5088 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	F_initialize_acl(m)
	mBase = m.M
	v5090 = m.ExcPending
	if v5090 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	v5093 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_InitPostgres[32])))
	v5094 = F_SearchSysCache1(m, int32(21), v5093)
	mBase = m.M
	v5095 = m.ExcPending
	if v5095 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1200:
	;
	if v5094 == int32(0) {
		goto L131
	} else {
		goto L1201
	}
L1201:
	;
	v5099 = v4367 + int32(448)
	v5100 = *(*int32)(unsafe.Add(mBase, uint32(v5094)+16))
	v5101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5100)+22)))
	v5102 = v5100 + v5101
	v5104 = v5102 + int32(4)
	v5107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5099))))
	v5110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5104))))
	if base.B2i32(v5107 == int32(0))|base.B2i32(v5107 != v5110) != 0 {
		v5128 = v5107
		v5129 = v5110
		goto L1203
	} else {
		goto L1204
	}
L1202:
	;
	if v5128-v5129 != 0 {
		goto L130
	} else {
		goto L1209
	}
L1203:
	;
	goto L1202
L1204:
	;
	v5113 = v5099
	v5114 = v5104
	goto L1205
L1205:
	;
	v5117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5114)+1)))
	v5118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5113)+1)))
	if v5118 == int32(0) {
		v5128 = v5118
		v5129 = v5117
		goto L1203
	} else {
		goto L1207
	}
L1206:
	;
	v5128 = v5118
	v5129 = v5117
	goto L1203
L1207:
	;
	v5121 = int32(1)
	if v5118 == v5117 {
		v5113 = v5113 + v5121
		v5114 = v5114 + v5121
		goto L1205
	} else {
		goto L1208
	}
L1208:
	;
	goto L1206
L1209:
	;
	v5132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v5132 != int32(1) {
		goto L1210
	} else {
		goto L1211
	}
L1210:
	;
	v5290 = *(*int32)(unsafe.Add(mBase, uint32(v5102)+72))
	v5291 = m.G0
	v5293 = v5291 - int32(16)
	m.G0 = v5293
	if base.B2i32(v5290 != int32(7))&base.B2i32(base.Ui32(v5290) <= base.Ui32(int32(34))) == int32(0) {
		goto L1238
	} else {
		goto L1239
	}
L1211:
	;
	v5136 = v4359 & int32(2)
	if v5136 == int32(0) {
		goto L1212
	} else {
		goto L1213
	}
L1212:
	;
	v5139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5102)+78)))
	if v5139&int32(1) == int32(0) {
		goto L129
	} else {
		goto L1215
	}
L1213:
	;
	goto L1214
L1214:
	;
	v5144 = int32(0)
	if base.B2i32(v5136 != v5144)|v4358 == v5144 {
		goto L1216
	} else {
		goto L1217
	}
L1215:
	;
	goto L1214
L1216:
	;
	v5151 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	v5153 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	v5155 = F_object_aclcheck(m, int32(1262), v5151, v5153, int64(2048))
	mBase = m.M
	v5156 = m.ExcPending
	if v5156 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1217:
	;
	goto L1218
L1218:
	;
	v5158 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v5161 = *(*int32)(unsafe.Add(mBase, uint32(v5102)+80))
	if v4358|(base.B2i32(v5158 != int32(1))|base.B2i32(v5161 < int32(0))) != 0 {
		goto L1210
	} else {
		goto L1221
	}
L1219:
	;
	if v5155 != 0 {
		goto L128
	} else {
		goto L1220
	}
L1220:
	;
	goto L1218
L1221:
	;
	v5167 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	v5168 = int32(0)
	v5170 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[81]))
	v5172 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[82]))
	v5176 = F_LWLockAcquire(m, v5172+int32(512), int32(1))
	mBase = m.M
	v5177 = m.ExcPending
	if v5177 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	v5178 = *(*int32)(unsafe.Add(mBase, uint32(v5170)))
	if int32(0) < v5178 {
		goto L1223
	} else {
		goto L1224
	}
L1223:
	;
	v5184 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[83]))
	v5188 = int32(0)
	v5191 = v5168
	goto L1226
L1224:
	;
	v5237 = v5168
	goto L1225
L1225:
	;
	v5258 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[82]))
	F_LWLockRelease(m, v5258+int32(512))
	mBase = m.M
	v5262 = m.ExcPending
	if v5262 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1226:
	;
	v5214 = *(*int32)(unsafe.Add(mBase, uint32(v5170+int32(36)+v5188<<(uint(int32(2))%32))))
	v5217 = v5184 + v5214*int32(768)
	v5218 = *(*int32)(unsafe.Add(mBase, uint32(v5217)+12))
	if v5218 == int32(0) {
		v5228 = v5191
		goto L1228
	} else {
		goto L1229
	}
L1227:
	;
	v5237 = v5228
	goto L1225
L1228:
	;
	v5230 = v5188 + int32(1)
	if v5230 != v5178 {
		v5188 = v5230
		v5191 = v5228
		goto L1226
	} else {
		goto L1235
	}
L1229:
	;
	v5221 = *(*int32)(unsafe.Add(mBase, uint32(v5217)+16))
	if v5221 != int32(1) {
		v5228 = v5191
		goto L1228
	} else {
		goto L1230
	}
L1230:
	;
	if v5167 != 0 {
		goto L1231
	} else {
		goto L1232
	}
L1231:
	;
	v5224 = *(*int32)(unsafe.Add(mBase, uint32(v5217)+20))
	if v5224 != v5167 {
		v5228 = v5191
		goto L1228
	} else {
		goto L1234
	}
L1232:
	;
	goto L1233
L1233:
	;
	v5228 = v5191 + int32(1)
	goto L1228
L1234:
	;
	goto L1233
L1235:
	;
	goto L1227
L1236:
	;
	v5263 = *(*int32)(unsafe.Add(mBase, uint32(v5102)+80))
	if v5263 < v5237 {
		goto L127
	} else {
		goto L1237
	}
L1237:
	;
	goto L1210
L1238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5305 = m.ExcPending
	if v5305 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1239:
	;
	goto L1240
L1240:
	;
	v5317 = v5290 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84])) = v5317 + int32(_a_F_InitPostgres_114)
	m.G0 = v5293 + int32(16)
	v5327 = *(*int32)(unsafe.Add(mBase, uint32(v5317)+uint32(_c_F_InitPostgres[85])))
	goto L1244
L1241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5293))) = v5290
	F_errmsg_internal(m, int32(_a_F_InitPostgres_115), v5293)
	mBase = m.M
	v5309 = m.ExcPending
	if v5309 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1242:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_116), int32(1293), int32(_a_F_InitPostgres_117))
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1244:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_118), v5327, int32(0), int32(1))
	mBase = m.M
	v5331 = m.ExcPending
	if v5331 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	v5334 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84]))
	v5335 = *(*int32)(unsafe.Add(mBase, uint32(v5334)))
	goto L1246
L1246:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_119), v5335, int32(4), int32(1))
	mBase = m.M
	v5339 = m.ExcPending
	if v5339 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	v5342 = F_SysCacheGetAttrNotNull(m, int32(21), v5094, int32(13))
	mBase = m.M
	v5343 = m.ExcPending
	if v5343 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	v5345 = F_text_to_cstring(m, base.I32_wrap_i64(v5342))
	mBase = m.M
	v5346 = m.ExcPending
	if v5346 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	v5349 = F_SysCacheGetAttrNotNull(m, int32(21), v5094, int32(14))
	mBase = m.M
	v5350 = m.ExcPending
	if v5350 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	v5352 = F_text_to_cstring(m, base.I32_wrap_i64(v5349))
	mBase = m.M
	v5353 = m.ExcPending
	if v5353 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	v5355 = F_pg_perm_setlocale(m, int32(3), v5345)
	mBase = m.M
	v5356 = m.ExcPending
	if v5356 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	if v5355 == int32(0) {
		goto L126
	} else {
		goto L1253
	}
L1253:
	;
	v5360 = F_pg_perm_setlocale(m, int32(0), v5352)
	mBase = m.M
	v5361 = m.ExcPending
	if v5361 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	if v5360 == int32(0) {
		goto L125
	} else {
		goto L1255
	}
L1255:
	;
	v5364 = m.G0
	v5366 = v5364 - int32(32)
	m.G0 = v5366
	v5370 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_InitPostgres[32])))
	v5371 = F_SearchSysCache1(m, int32(21), v5370)
	mBase = m.M
	v5372 = m.ExcPending
	if v5372 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1256:
	;
	v5440 = F_SysCacheGetAttr(m, int32(21), v5094, int32(17), v4367+int32(527))
	mBase = m.M
	v5441 = m.ExcPending
	if v5441 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1257:
	;
	if v5371 != 0 {
		goto L1258
	} else {
		goto L1259
	}
L1258:
	;
	v5373 = *(*int32)(unsafe.Add(mBase, uint32(v5371)+16))
	v5374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5373)+22)))
	v5375 = v5373 + v5374
	v5376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5375)+76)))
	switch v5376 - int32(98) {
	case 0:
		goto L1262
	case 1:
		goto L1264
	default:
		goto L1263
	case 7:
		goto L1265
	}
L1259:
	;
	goto L1260
L1260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5424 = m.ExcPending
	if v5424 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1261:
	;
	v5412 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5411)+3)) = uint8(v5412)
	F_ReleaseCatCache(m, v5371)
	mBase = m.M
	v5415 = m.ExcPending
	if v5415 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1262:
	;
	v5408 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v5409 = F_create_pg_locale_builtin(m, int32(100), v5408)
	mBase = m.M
	v5410 = m.ExcPending
	if v5410 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5391 = m.ExcPending
	if v5391 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1264:
	;
	v5385 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v5386 = F_create_pg_locale_libc(m, int32(100), v5385)
	mBase = m.M
	v5387 = m.ExcPending
	if v5387 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1265:
	;
	v5381 = F_create_pg_locale_icu(m)
	mBase = m.M
	v5382 = m.ExcPending
	if v5382 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1266:
	;
	v5411 = v5381
	goto L1261
L1267:
	;
	v5411 = v5386
	goto L1261
L1268:
	;
	v5392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5375)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v5366)+20)) = v5392
	*(*int32)(unsafe.Add(mBase, uint32(v5366)+16)) = int32(_a_F_InitPostgres_120)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_121), v5366+int32(16))
	mBase = m.M
	v5400 = m.ExcPending
	if v5400 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1269:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_122), int32(1156), int32(_a_F_InitPostgres_120))
	mBase = m.M
	v5405 = m.ExcPending
	if v5405 != 0 {
		goto L1
	} else {
		goto L1270
	}
L1270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1271:
	;
	v5411 = v5409
	goto L1261
L1272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[86])) = v5411
	m.G0 = v5366 + int32(32)
	goto L1256
L1273:
	;
	v5426 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v5366))) = v5426
	F_errmsg_internal(m, int32(_a_F_InitPostgres_123), v5366)
	mBase = m.M
	v5430 = m.ExcPending
	if v5430 != 0 {
		goto L1
	} else {
		goto L1274
	}
L1274:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_122), int32(1142), int32(_a_F_InitPostgres_120))
	mBase = m.M
	v5435 = m.ExcPending
	if v5435 != 0 {
		goto L1
	} else {
		goto L1275
	}
L1275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1276:
	;
	v5442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4367)+527)))
	if v5442 != 0 {
		goto L1277
	} else {
		goto L1278
	}
L1277:
	;
	F_ReleaseCatCache(m, v5094)
	mBase = m.M
	v5548 = m.ExcPending
	if v5548 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1278:
	;
	v5444 = F_text_to_cstring(m, base.I32_wrap_i64(v5440))
	mBase = m.M
	v5445 = m.ExcPending
	if v5445 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1279:
	;
	v5447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5102)+76)))
	if v5447 != int32(99) {
		goto L1281
	} else {
		goto L1282
	}
L1280:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), v5540, int32(_a_F_InitPostgres_124))
	mBase = m.M
	v5543 = m.ExcPending
	if v5543 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1281:
	;
	v5452 = F_SysCacheGetAttrNotNull(m, int32(21), v5094, int32(15))
	mBase = m.M
	v5453 = m.ExcPending
	if v5453 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1282:
	;
	v5459 = v5345
	v5460 = int32(99)
	goto L1283
L1283:
	;
	v5462 = F_get_collation_actual_version(m, base.I32_extend8_s(v5460), v5459)
	mBase = m.M
	v5463 = m.ExcPending
	if v5463 != 0 {
		goto L1
	} else {
		goto L1286
	}
L1284:
	;
	v5455 = F_text_to_cstring(m, base.I32_wrap_i64(v5452))
	mBase = m.M
	v5456 = m.ExcPending
	if v5456 != 0 {
		goto L1
	} else {
		goto L1285
	}
L1285:
	;
	v5457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5102)+76)))
	v5459 = v5455
	v5460 = v5457
	goto L1283
L1286:
	;
	if v5462 == int32(0) {
		goto L1287
	} else {
		goto L1288
	}
L1287:
	;
	v5468 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5469 = m.ExcPending
	if v5469 != 0 {
		goto L1
	} else {
		goto L1290
	}
L1288:
	;
	goto L1289
L1289:
	;
	v5483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5462))))
	v5486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5444))))
	if base.B2i32(v5483 == int32(0))|base.B2i32(v5483 != v5486) != 0 {
		v5504 = v5483
		v5505 = v5486
		goto L1294
	} else {
		goto L1295
	}
L1290:
	;
	if v5468 == int32(0) {
		goto L1277
	} else {
		goto L1291
	}
L1291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+96)) = v4367 + int32(448)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_125), v4367+int32(96))
	mBase = m.M
	v5479 = m.ExcPending
	if v5479 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1292:
	;
	v5540 = int32(479)
	goto L1280
L1293:
	;
	if v5504-v5505 == int32(0) {
		goto L1277
	} else {
		goto L1300
	}
L1294:
	;
	goto L1293
L1295:
	;
	v5489 = v5462
	v5490 = v5444
	goto L1296
L1296:
	;
	v5493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5490)+1)))
	v5494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5489)+1)))
	if v5494 == int32(0) {
		v5504 = v5494
		v5505 = v5493
		goto L1294
	} else {
		goto L1298
	}
L1297:
	;
	v5504 = v5494
	v5505 = v5493
	goto L1294
L1298:
	;
	v5497 = int32(1)
	if v5494 == v5493 {
		v5489 = v5489 + v5497
		v5490 = v5490 + v5497
		goto L1296
	} else {
		goto L1299
	}
L1299:
	;
	goto L1297
L1300:
	;
	v5511 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5512 = m.ExcPending
	if v5512 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1301:
	;
	if v5511 == int32(0) {
		goto L1277
	} else {
		goto L1302
	}
L1302:
	;
	v5516 = v4367 + int32(448)
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+144)) = v5516
	F_errmsg(m, int32(_a_F_InitPostgres_126), v4367+int32(144))
	mBase = m.M
	v5522 = m.ExcPending
	if v5522 != 0 {
		goto L1
	} else {
		goto L1303
	}
L1303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+132)) = v5462
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+128)) = v5444
	v5528 = F_errdetail(m, int32(_a_F_InitPostgres_127), v4367+int32(128))
	mBase = m.M
	v5529 = m.ExcPending
	if v5529 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	v5530 = F_quote_identifier(m, v5516)
	mBase = m.M
	v5531 = m.ExcPending
	if v5531 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+112)) = v5530
	F_errhint(m, int32(_a_F_InitPostgres_128), v4367+int32(112))
	mBase = m.M
	v5537 = m.ExcPending
	if v5537 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1306:
	;
	v5540 = int32(490)
	goto L1280
L1307:
	;
	goto L1277
L1308:
	;
	goto L1183
L1309:
	;
	F_pfree(m, v5054)
	mBase = m.M
	v5552 = m.ExcPending
	if v5552 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1310:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v5554 = m.ExcPending
	if v5554 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1311:
	;
	F_initialize_acl(m)
	mBase = m.M
	v5556 = m.ExcPending
	if v5556 != 0 {
		goto L1
	} else {
		goto L1312
	}
L1312:
	;
	goto L1183
L1313:
	;
	F_process_startup_options(m, v5583, v4358)
	mBase = m.M
	v5585 = m.ExcPending
	if v5585 != 0 {
		goto L1
	} else {
		goto L1316
	}
L1314:
	;
	goto L1315
L1315:
	;
	v5587 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	v5589 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[87]))
	v5591 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v5591 == int32(1) {
		goto L1317
	} else {
		goto L1318
	}
L1316:
	;
	goto L1315
L1317:
	;
	v5596 = F_table_open(m, int32(2964), int32(1))
	mBase = m.M
	v5597 = m.ExcPending
	if v5597 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1318:
	;
	goto L1319
L1319:
	;
	v5627 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	if int32(0) < v5627 {
		goto L1329
	} else {
		goto L1330
	}
L1320:
	;
	v5599 = F_GetCatalogSnapshot(m, int32(2964))
	mBase = m.M
	v5600 = m.ExcPending
	if v5600 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	v5601 = F_RegisterSnapshot(m, v5599)
	mBase = m.M
	v5602 = m.ExcPending
	if v5602 != 0 {
		goto L1
	} else {
		goto L1322
	}
L1322:
	;
	F_ApplySetting(m, v5601, v5587, v5589, v5596, int32(8))
	mBase = m.M
	v5605 = m.ExcPending
	if v5605 != 0 {
		goto L1
	} else {
		goto L1323
	}
L1323:
	;
	F_ApplySetting(m, v5601, int32(0), v5589, v5596, int32(7))
	mBase = m.M
	v5609 = m.ExcPending
	if v5609 != 0 {
		goto L1
	} else {
		goto L1324
	}
L1324:
	;
	F_ApplySetting(m, v5601, v5587, int32(0), v5596, int32(6))
	mBase = m.M
	v5613 = m.ExcPending
	if v5613 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	v5614 = int32(0)
	F_ApplySetting(m, v5601, v5614, v5614, v5596, int32(5))
	mBase = m.M
	v5618 = m.ExcPending
	if v5618 != 0 {
		goto L1
	} else {
		goto L1326
	}
L1326:
	;
	F_UnregisterSnapshot(m, v5601)
	mBase = m.M
	v5620 = m.ExcPending
	if v5620 != 0 {
		goto L1
	} else {
		goto L1327
	}
L1327:
	;
	F_relation_close(m, v5596, int32(1))
	mBase = m.M
	v5623 = m.ExcPending
	if v5623 != 0 {
		goto L1
	} else {
		goto L1328
	}
L1328:
	;
	goto L1319
L1329:
	;
	F_pg_usleep(m, v5627*int32(_a_F_InitPostgres_100))
	mBase = m.M
	goto L1331
L1330:
	;
	goto L1331
L1331:
	;
	v5633 = m.G0
	v5635 = v5633 - int32(16)
	m.G0 = v5635
	v5638 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v5638 == int32(0) {
		goto L1333
	} else {
		goto L1334
	}
L1332:
	;
	m.G0 = v5635 + int32(16)
	F_InitializeClientEncoding(m)
	mBase = m.M
	v5723 = m.ExcPending
	if v5723 != 0 {
		goto L1
	} else {
		goto L1341
	}
L1333:
	;
	v5641 = int32(_a_F_InitPostgres_8)
	v5642 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v5645 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v5645
	v5647 = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v5635)+12)) = v5647
	*(*int32)(unsafe.Add(mBase, uint32(v5635)+8)) = v5647
	v5654 = F_list_make1_impl(m, int32(480), v5635+int32(8))
	mBase = m.M
	v5655 = m.ExcPending
	if v5655 != 0 {
		goto L1
	} else {
		goto L1336
	}
L1334:
	;
	goto L1335
L1335:
	;
	F_CacheRegisterSyscacheCallback(m, int32(38), int32(505), int64(0))
	mBase = m.M
	v5695 = m.ExcPending
	if v5695 != 0 {
		goto L1
	} else {
		goto L1337
	}
L1336:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v5642
	v5658 = int32(_a_F_InitPostgres_129)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[88])) = v5654
	v5660 = int32(_a_F_InitPostgres_130)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[89])) = int32(11)
	v5663 = int32(_a_F_InitPostgres_131)
	v5664 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[90])) = uint8(v5664)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[91])) = uint8(v5664)
	v5671 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[92])) = v5671
	v5675 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[88]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[93])) = v5675
	v5679 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[89]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[94])) = v5679
	v5683 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[90])))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[95])) = uint8(v5683)
	v5685 = int32(_a_F_InitPostgres_132)
	v5687 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[96]))
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[96])) = v5687 + int64(1)
	goto L1332
L1337:
	;
	F_CacheRegisterSyscacheCallback(m, int32(11), int32(505), int64(0))
	mBase = m.M
	v5700 = m.ExcPending
	if v5700 != 0 {
		goto L1
	} else {
		goto L1338
	}
L1338:
	;
	F_CacheRegisterSyscacheCallback(m, int32(9), int32(505), int64(0))
	mBase = m.M
	v5705 = m.ExcPending
	if v5705 != 0 {
		goto L1
	} else {
		goto L1339
	}
L1339:
	;
	F_CacheRegisterSyscacheCallback(m, int32(21), int32(505), int64(0))
	mBase = m.M
	v5710 = m.ExcPending
	if v5710 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1340:
	;
	v5712 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[91])) = uint8(v5712)
	v5715 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[97])) = uint8(v5715)
	goto L1332
L1341:
	;
	v5726 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v5728 = F_MemoryContextAllocZero(m, v5726, int32(20))
	mBase = m.M
	v5729 = m.ExcPending
	if v5729 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1342:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[98])) = v5728
	if v4359&int32(1) != 0 {
		goto L1343
	} else {
		goto L1344
	}
L1343:
	;
	v5734 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[99]))
	F_load_libraries(m, v5734, int32(_a_F_InitPostgres_133), int32(0))
	mBase = m.M
	v5738 = m.ExcPending
	if v5738 != 0 {
		goto L1
	} else {
		goto L1346
	}
L1344:
	;
	goto L1345
L1345:
	;
	if v4376 == int32(0) {
		goto L139
	} else {
		goto L1348
	}
L1346:
	;
	v5740 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[100]))
	F_load_libraries(m, v5740, int32(_a_F_InitPostgres_134), int32(1))
	mBase = m.M
	v5744 = m.ExcPending
	if v5744 != 0 {
		goto L1
	} else {
		goto L1347
	}
L1347:
	;
	goto L1345
L1348:
	;
	goto L140
L1349:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v5775 = m.ExcPending
	if v5775 != 0 {
		goto L1
	} else {
		goto L1350
	}
L1350:
	;
	goto L139
L1351:
	;
	v5975 = v4367
	goto L134
L1352:
	;
	v5811 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[31])) = uint8(v5811)
	v5814 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[101]))
	if v5814 != 0 {
		goto L1355
	} else {
		goto L1356
	}
L1353:
	;
	goto L1354
L1354:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5953 = m.ExcPending
	if v5953 != 0 {
		goto L1
	} else {
		goto L1379
	}
L1355:
	;
	v5815 = *(*int32)(unsafe.Add(mBase, uint32(v5814)+4))
	if int32(0) < v5815 {
		goto L1358
	} else {
		goto L1359
	}
L1356:
	;
	v5941 = int32(0)
	goto L1357
L1357:
	;
	F_list_free(m, v5941)
	mBase = m.M
	v5943 = m.ExcPending
	if v5943 != 0 {
		goto L1
	} else {
		goto L1378
	}
L1358:
	;
	v5821 = v5801
	goto L1361
L1359:
	;
	goto L1360
L1360:
	;
	v5914 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[101]))
	v5941 = v5914
	goto L1357
L1361:
	;
	v5843 = *(*int32)(unsafe.Add(mBase, uint32(v5814)+12))
	v5847 = *(*int32)(unsafe.Add(mBase, uint32(v5843+v5821<<(uint(int32(2))%32))))
	v5848 = *(*int32)(unsafe.Add(mBase, uint32(v5847)+8))
	if v5848 != 0 {
		goto L1364
	} else {
		goto L1365
	}
L1362:
	;
	goto L1360
L1363:
	;
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5847)))
	F_pfree(m, v5876)
	mBase = m.M
	v5878 = m.ExcPending
	if v5878 != 0 {
		goto L1
	} else {
		goto L1374
	}
L1364:
	;
	v5849 = m.T0[v5848].(func(*base.Module) int32)(m)
	mBase = m.M
	v5850 = m.ExcPending
	if v5850 != 0 {
		goto L1
	} else {
		goto L1367
	}
L1365:
	;
	goto L1366
L1366:
	;
	v5855 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5856 = m.ExcPending
	if v5856 != 0 {
		goto L1
	} else {
		goto L1369
	}
L1367:
	;
	if v5849 == int32(0) {
		goto L1363
	} else {
		goto L1368
	}
L1368:
	;
	goto L1366
L1369:
	;
	if v5855 == int32(0) {
		goto L1363
	} else {
		goto L1370
	}
L1370:
	;
	v5859 = *(*int32)(unsafe.Add(mBase, uint32(v5847)))
	*(*int32)(unsafe.Add(mBase, uint32(v5804)+16)) = v5859
	F_errmsg(m, int32(_a_F_InitPostgres_88), v5804+int32(16))
	mBase = m.M
	v5865 = m.ExcPending
	if v5865 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1371:
	;
	v5866 = *(*int32)(unsafe.Add(mBase, uint32(v5847)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5804))) = v5866
	v5869 = F_errdetail(m, int32(_a_F_InitPostgres_88), v5804)
	mBase = m.M
	v5870 = m.ExcPending
	if v5870 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1372:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1547), int32(_a_F_InitPostgres_135))
	mBase = m.M
	v5875 = m.ExcPending
	if v5875 != 0 {
		goto L1
	} else {
		goto L1373
	}
L1373:
	;
	goto L1363
L1374:
	;
	v5879 = *(*int32)(unsafe.Add(mBase, uint32(v5847)+4))
	F_pfree(m, v5879)
	mBase = m.M
	v5881 = m.ExcPending
	if v5881 != 0 {
		goto L1
	} else {
		goto L1375
	}
L1375:
	;
	F_pfree(m, v5847)
	mBase = m.M
	v5883 = m.ExcPending
	if v5883 != 0 {
		goto L1
	} else {
		goto L1376
	}
L1376:
	;
	v5885 = v5821 + int32(1)
	v5886 = *(*int32)(unsafe.Add(mBase, uint32(v5814)+4))
	if v5885 < v5886 {
		v5821 = v5885
		goto L1361
	} else {
		goto L1377
	}
L1377:
	;
	goto L1362
L1378:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[101])) = int32(0)
	m.G0 = v5804 + int32(32)
	goto L1351
L1379:
	;
	F_errmsg_internal(m, int32(_a_F_InitPostgres_136), int32(0))
	mBase = m.M
	v5957 = m.ExcPending
	if v5957 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1380:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1538), int32(_a_F_InitPostgres_135))
	mBase = m.M
	v5962 = m.ExcPending
	if v5962 != 0 {
		goto L1
	} else {
		goto L1381
	}
L1381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1382:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v5997 = m.ExcPending
	if v5997 != 0 {
		goto L1
	} else {
		goto L1383
	}
L1383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+272)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_137), v4367+int32(272))
	mBase = m.M
	v6005 = m.ExcPending
	if v6005 != 0 {
		goto L1
	} else {
		goto L1384
	}
L1384:
	;
	F_errhint(m, int32(_a_F_InitPostgres_138), int32(0))
	mBase = m.M
	v6009 = m.ExcPending
	if v6009 != 0 {
		goto L1
	} else {
		goto L1385
	}
L1385:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1135), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v6014 = m.ExcPending
	if v6014 != 0 {
		goto L1
	} else {
		goto L1386
	}
L1386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+16)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_111), v4367+int32(16))
	mBase = m.M
	v6025 = m.ExcPending
	if v6025 != 0 {
		goto L1
	} else {
		goto L1388
	}
L1388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367))) = v5054
	v6028 = F_errdetail(m, int32(_a_F_InitPostgres_139), v4367)
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		goto L1
	} else {
		goto L1389
	}
L1389:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1196), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v6034 = m.ExcPending
	if v6034 != 0 {
		goto L1
	} else {
		goto L1390
	}
L1390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1391:
	;
	v6040 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+48)) = v6040
	F_errmsg_internal(m, int32(_a_F_InitPostgres_123), v4367+int32(48))
	mBase = m.M
	v6046 = m.ExcPending
	if v6046 != 0 {
		goto L1
	} else {
		goto L1392
	}
L1392:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(350), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6051 = m.ExcPending
	if v6051 != 0 {
		goto L1
	} else {
		goto L1393
	}
L1393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1394:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v6058 = m.ExcPending
	if v6058 != 0 {
		goto L1
	} else {
		goto L1395
	}
L1395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+224)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_140), v4367+int32(224))
	mBase = m.M
	v6066 = m.ExcPending
	if v6066 != 0 {
		goto L1
	} else {
		goto L1396
	}
L1396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+212)) = v5104
	v6069 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+208)) = v6069
	v6074 = F_errdetail(m, int32(_a_F_InitPostgres_141), v4367+int32(208))
	mBase = m.M
	v6075 = m.ExcPending
	if v6075 != 0 {
		goto L1
	} else {
		goto L1397
	}
L1397:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(360), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6080 = m.ExcPending
	if v6080 != 0 {
		goto L1
	} else {
		goto L1398
	}
L1398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1399:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v6087 = m.ExcPending
	if v6087 != 0 {
		goto L1
	} else {
		goto L1400
	}
L1400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+192)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_142), v4367+int32(192))
	mBase = m.M
	v6095 = m.ExcPending
	if v6095 != 0 {
		goto L1
	} else {
		goto L1401
	}
L1401:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(380), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6100 = m.ExcPending
	if v6100 != 0 {
		goto L1
	} else {
		goto L1402
	}
L1402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1403:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v6107 = m.ExcPending
	if v6107 != 0 {
		goto L1
	} else {
		goto L1404
	}
L1404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+176)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_143), v4367+int32(176))
	mBase = m.M
	v6115 = m.ExcPending
	if v6115 != 0 {
		goto L1
	} else {
		goto L1405
	}
L1405:
	;
	v6118 = F_errdetail(m, int32(_a_F_InitPostgres_144), int32(0))
	mBase = m.M
	v6119 = m.ExcPending
	if v6119 != 0 {
		goto L1
	} else {
		goto L1406
	}
L1406:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(393), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6124 = m.ExcPending
	if v6124 != 0 {
		goto L1
	} else {
		goto L1407
	}
L1407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1408:
	;
	F_errcode(m, int32(_a_F_InitPostgres_103))
	mBase = m.M
	v6131 = m.ExcPending
	if v6131 != 0 {
		goto L1
	} else {
		goto L1409
	}
L1409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+160)) = v4367 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_145), v4367+int32(160))
	mBase = m.M
	v6139 = m.ExcPending
	if v6139 != 0 {
		goto L1
	} else {
		goto L1410
	}
L1410:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(414), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L1
	} else {
		goto L1411
	}
L1411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1412:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_146), int32(0))
	mBase = m.M
	v6152 = m.ExcPending
	if v6152 != 0 {
		goto L1
	} else {
		goto L1413
	}
L1413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+64)) = v5345
	v6157 = F_errdetail(m, int32(_a_F_InitPostgres_147), v4367-int32(-64))
	mBase = m.M
	v6158 = m.ExcPending
	if v6158 != 0 {
		goto L1
	} else {
		goto L1414
	}
L1414:
	;
	F_errhint(m, int32(_a_F_InitPostgres_148), int32(0))
	mBase = m.M
	v6162 = m.ExcPending
	if v6162 != 0 {
		goto L1
	} else {
		goto L1415
	}
L1415:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(440), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6167 = m.ExcPending
	if v6167 != 0 {
		goto L1
	} else {
		goto L1416
	}
L1416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1417:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_146), int32(0))
	mBase = m.M
	v6175 = m.ExcPending
	if v6175 != 0 {
		goto L1
	} else {
		goto L1418
	}
L1418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+80)) = v5352
	v6180 = F_errdetail(m, int32(_a_F_InitPostgres_149), v4367+int32(80))
	mBase = m.M
	v6181 = m.ExcPending
	if v6181 != 0 {
		goto L1
	} else {
		goto L1419
	}
L1419:
	;
	F_errhint(m, int32(_a_F_InitPostgres_148), int32(0))
	mBase = m.M
	v6185 = m.ExcPending
	if v6185 != 0 {
		goto L1
	} else {
		goto L1420
	}
L1420:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(447), int32(_a_F_InitPostgres_124))
	mBase = m.M
	v6190 = m.ExcPending
	if v6190 != 0 {
		goto L1
	} else {
		goto L1421
	}
L1421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1422:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v6198 = m.ExcPending
	if v6198 != 0 {
		goto L1
	} else {
		goto L1423
	}
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4367)+256)) = v4355
	F_errmsg(m, int32(_a_F_InitPostgres_111), v4367+int32(256))
	mBase = m.M
	v6204 = m.ExcPending
	if v6204 != 0 {
		goto L1
	} else {
		goto L1424
	}
L1424:
	;
	v6207 = F_errdetail(m, int32(_a_F_InitPostgres_150), int32(0))
	mBase = m.M
	v6208 = m.ExcPending
	if v6208 != 0 {
		goto L1
	} else {
		goto L1425
	}
L1425:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1121), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v6213 = m.ExcPending
	if v6213 != 0 {
		goto L1
	} else {
		goto L1426
	}
L1426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
