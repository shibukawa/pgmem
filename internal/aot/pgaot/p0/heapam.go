package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_relation_copy_for_cluster(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v409 int32
	_ = v409
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int64
	_ = v488
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v526 int64
	_ = v526
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int64
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int64
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1031 int32
	_ = v1031
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1209 float64
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1229 float64
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1276 float64
	_ = v1276
	var v1280 float64
	_ = v1280
	var v1284 float64
	_ = v1284
	var v1292 int32
	_ = v1292
	var v1293 float64
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1356 float64
	_ = v1356
	var v1361 int64
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 float64
	_ = v1514
	var v1518 int64
	_ = v1518
	var v1520 int64
	_ = v1520
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1666 int64
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1737 int32
	_ = v1737
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1809 float64
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1900 float64
	_ = v1900
	var v1905 int64
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1964 int32
	_ = v1964
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2108 int32
	_ = v2108
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	v11 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v35) < base.Ui32(int32(12000)) {
		v44 = v34
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v48 = F_palloc(m, v45<<(uint(int32(2))%32))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v39 == int32(99) {
		v44 = v34
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v42 = F_isTempToastNamespace(m, v39)
	mBase = m.M
	v44 = v42
	goto L2
L5:
	;
	return
L6:
	;
	v50 = F_palloc(m, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v54 = m.G0
	v56 = v54 - int32(112)
	m.G0 = v56
	v59 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v64 = F_AllocSetContextCreateInternal(m, v59, int32(357586), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v66 = int32(4536272)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
	v71 = F_palloc0(m, int32(72))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = l0
	v78 = F_RelationGetNumberOfBlocksInFork(m, l1, v73)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v71)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v71)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = v78
	v86 = F_smgr_bulk_start_rel(m, l1, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v56)+28)) = int64(103079215116)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+52)) = v91
	v98 = F_hash_create(m, int32(177259), int32(128), v56+int32(12), int32(1064))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+56)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v56)+32)) = int32(20)
	v108 = F_hash_create(m, int32(243614), int32(128), v56+int32(12), int32(1064))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+60)) = v108
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v67
	v114 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v114 < int32(2) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	m.G0 = v56 + int32(112)
	if l3 != 0 {
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v151 = v56 + int32(60)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v157 = F_LWLockAcquire(m, v153+int32(512), int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L26
	}
L16:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+104)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+20)) = uint8(v141)
	if v141 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L17:
	;
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+20)) = uint8(v139)
	goto L14
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+118)))
	if v119 != int32(112) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+56))
	goto L20
L20:
	;
	if base.Ui32(v122) < base.Ui32(int32(12000)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+20)) = uint8(v125)
	v149 = v71 + int32(20)
	goto L15
L22:
	;
	goto L23
L23:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+180))
	if v130 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+119)))
	switch v134 - int32(109) {
	case 0, 5:
		goto L16
	default:
		goto L17
	}
L25:
	;
	v149 = v71 + int32(20)
	goto L15
L26:
	;
	if v151 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v161
	goto L29
L28:
	;
	goto L29
L29:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v164+int32(512))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v56)+60))
	if v169 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v172)
	goto L14
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+32)) = v169
	v175 = F_GetXLogInsertRecPtr(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v71)+48)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v56)+80)) = int64(4535485464580)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+104)) = v182
	v189 = F_hash_create(m, int32(341728), int32(128), v56-int32(-64), int32(1064))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+64)) = v189
	goto L14
L36:
	;
	v746 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L100
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = int64(8589934593)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = int64(2)
	v526 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+56)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v526
	v533 = int32(0)
	v540 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v540 == v533 {
		goto L82
	} else {
		goto L83
	}
L38:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v447 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L39:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	v201 = m.G0
	v203 = v201 - int32(16)
	m.G0 = v203
	v205 = int32(0)
	v207 = F_tuplesort_begin_common(m, v200, v205, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L71
	}
L42:
	;
	v209 = int32(4536272)
	v210 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v212
	v215 = F_palloc0(m, int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
	if v218 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l2)+192))
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = int32(1841)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+40)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v207)+60)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = int32(1842)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+16)) = int32(1843)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = int32(1844)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = int32(1845)
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = int32(1846)
	v258 = F_BuildIndexInfo(m, l2)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L50
	}
L45:
	;
	v223 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v223 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v227)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v228
	F_errmsg_internal(m, int32(514630), v203)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(504643), int32(275), int32(219576))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+4)) = v258
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+12)))
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+36)) = uint8(base.B2i32(v261 != v262))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v31
	v267 = F__bt_mkscankey(m, l2, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+76))
	if v270 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v271 = F_CreateExecutorState(m)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v207)+40))
	v289 = F_palloc0(m, v286*int32(36))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L61
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v271
	v275 = F_MakeSingleTupleTableSlot(m, v31, int32(1632200))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+152))
	if v278 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v281 = v278
	goto L59
L58:
	;
	v279 = F_MakePerTupleExprContext(m, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v275
	goto L54
L60:
	;
	v281 = v279
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+44)) = v289
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v207)+40))
	if v293 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_pfree(m, v267)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L70
	}
L63:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v267)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	v305 = int32(base.Ui32(v301)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+9)) = uint8(v305)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+10)) = uint16(v307)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+20)) = uint8(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v311&int32(16777216))>>(uint(int32(24))%32)), v289)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v207)+40))
	if v318 < int32(2) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v335 = int32(1)
	goto L66
L66:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v207)+44))
	v352 = v349 + v335*int32(36)
	v354 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v354
	v358 = v267 + int32(16) + v335*int32(48)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v365 = int32(base.Ui32(v361)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+9)) = uint8(v365)
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358)+4)))
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+20)) = uint8(v368)
	*(*uint16)(unsafe.Add(mBase, uint32(v352)+10)) = uint16(v367)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v371&int32(16777216))>>(uint(int32(24))%32)), v352)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L68
	}
L67:
	;
	goto L62
L68:
	;
	v379 = v335 + int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v207)+40))
	if v379 < v380 {
		v335 = v379
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v210
	m.G0 = v203 + int32(16)
	v442 = v207
	goto L38
L71:
	;
	v442 = int32(0)
	goto L38
L72:
	;
	v480 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v486 = m.T0[v485].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(4194416), v480, v480, v480, int32(449))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v451 != int32(1) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v454 = int32(4530932)
	v456 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v457 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v456 + v457
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v460 + v457
	*(*int64)(unsafe.Add(mBase, uint32(v447+int32(8))+232)) = int64(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v447))) = v468 + v457
	v474 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v474 - v457
	goto L73
L76:
	;
	v488 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v486)+36)))
	v491 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v491 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v731 = v442
	v732 = v486
	v737 = v11
	goto L36
L78:
	;
	goto L77
L79:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v495 != int32(1) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v498 = int32(4530932)
	v500 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v501 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v500 + v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v504 + v501
	*(*int64)(unsafe.Add(mBase, uint32(v491+int32(40))+232)) = v488
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v512 + v501
	v518 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v518 - v501
	goto L78
L81:
	;
	v707 = int32(0)
	v710 = F_index_beginscan(m, l0, l2, int32(4194416), v707, v707, v707)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L98
	}
L82:
	;
	goto L81
L83:
	;
	goto L84
L84:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v546&int32(1) == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v551 = int32(4530932)
	v553 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v554 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v553 + v554
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v557 + v554
	goto L87
L86:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v688 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v687 + v688
	v691 = int32(4530932)
	v693 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v693 - v688
	goto L82
L87:
	;
	goto L89
L89:
	;
	goto L90
L90:
	;
	goto L94
L94:
	;
	v652 = int32(0)
	v655 = v533
	goto L95
L95:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-8)+v655<<(uint(int32(2))%32))))
	v665 = int32(3)
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v27+int32(-32)+v655<<(uint(v665)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v540+int32(232)+v664<<(uint(v665)%32)))) = v671
	v673 = int32(1)
	v676 = v652 + v673
	if v676 != int32(2) {
		v652 = v676
		v655 = v655 + v673
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L86
L97:
	;
	goto L96
L98:
	;
	v712 = int32(0)
	F_index_rescan(m, v710, v712, v712, v712, v712)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v731 = v11
	v732 = v11
	v737 = v710
	goto L36
L100:
	;
	v764 = int32(-1)
	goto L104
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L5
	} else {
		goto L423
	}
L102:
	;
	if v746 != 0 {
		goto L328
	} else {
		goto L329
	}
L103:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1705)+188))
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+12))
	m.T0[v1707].(func(*base.Module, int32))(m, v732)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L5
	} else {
		goto L327
	}
L104:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v775 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	F_index_endscan(m, v737)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L5
	} else {
		goto L325
	}
L106:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v737 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L108
L110:
	;
	goto L105
L111:
	;
	v901 = int32(0)
	v903 = F_ExecFetchSlotHeapTuple(m, v746, v901, v901)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L138
	}
L112:
	;
	v779 = F_index_getnext_slot(m, v737, int32(1), v746)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v746)+36)) = v800
	v803 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v803 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	if v779 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+72)))
	if v783 != int32(1) {
		v900 = v764
		goto L111
	} else {
		goto L117
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg_internal(m, int32(142169), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(506907), int32(798), int32(219527))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _consts[26])))
	if v805&int32(1) == int32(0) {
		goto L101
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)+188))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+20))
	v814 = m.T0[v813].(func(*base.Module, int32, int32, int32) int32)(m, v732, int32(1), v746)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	if v814 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v819 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v732)+36)))
	v822 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v822 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v732)+52))
	if v764 == v853 {
		v900 = v764
		goto L111
	} else {
		goto L133
	}
L129:
	;
	goto L103
L130:
	;
	goto L129
L131:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v826 != int32(1) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v829 = int32(4530932)
	v831 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v832 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v831 + v832
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	*(*int32)(unsafe.Add(mBase, uint32(v822))) = v835 + v832
	*(*int64)(unsafe.Add(mBase, uint32(v822+int32(48))+232)) = v819
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	*(*int32)(unsafe.Add(mBase, uint32(v822))) = v843 + v832
	v849 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v849 - v832
	goto L130
L133:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v732)+36))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v732)+40))
	v860 = base.I32_rem_u_s(v856+v853-v858, v856)
	v866 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v866 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v732)+52))
	v900 = v897
	goto L111
L135:
	;
	goto L134
L136:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v870 != int32(1) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v873 = int32(4530932)
	v875 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v876 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v875 + v876
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = v879 + v876
	*(*int64)(unsafe.Add(mBase, uint32(v866+int32(48))+232)) = base.I64_extend_i32_u(v860 + int32(1))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = v887 + v876
	v893 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v893 - v876
	goto L135
L138:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v746)+68))
	F_LockBuffer(m, v905, int32(1))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	v909 = F_HeapTupleSatisfiesVacuum(m, v903, l4, v905)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L146
	}
L140:
	;
	F_LockBuffer(m, v905, int32(0))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L5
	} else {
		goto L262
	}
L141:
	;
	v1284 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1284, float64(1))
	goto L140
L142:
	;
	F_LockBuffer(m, v905, int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L251
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L5
	} else {
		goto L248
	}
L144:
	;
	if v44 != 0 {
		goto L196
	} else {
		goto L197
	}
L145:
	;
	if v44 != 0 {
		goto L140
	} else {
		goto L147
	}
L146:
	;
	switch v909 {
	case 0:
		goto L142
	case 1:
		goto L140
	case 2:
		goto L141
	case 3:
		goto L145
	case 4:
		goto L144
	default:
		goto L143
	}
L147:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v903)+16))
	v912 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v911)+20)))
	v913 = int32(768)
	if v912&v913 != v913 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v911)))
	v919 = v917
	goto L150
L149:
	;
	v919 = int32(2)
	goto L150
L150:
	;
	if base.Ui32(v919) < base.Ui32(int32(3)) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v1039 != 0 {
		goto L140
	} else {
		goto L191
	}
L152:
	;
	v1039 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _consts[35]))
	if v930 == v919 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1039 = int32(1)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	if v934 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1039 = v1031
	goto L151
L159:
	;
	v938 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v938 == int32(0) {
		v1031 = int32(0)
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	v1002 = int32(0)
	v1004 = v934 - int32(1)
	goto L181
L162:
	;
	v943 = v938
	goto L163
L163:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v943)+20))
	if v948 == int32(4) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1031 = int32(0)
	goto L158
L165:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v943)+80))
	if v995 != 0 {
		v943 = v995
		goto L163
	} else {
		goto L180
	}
L166:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	if v951 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v954 = int32(1)
	if v919 == v951 {
		v1031 = v954
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v943)+52))
	v958 = v956 - int32(1)
	if v958 < int32(0) {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v963 = int32(0)
	v965 = v958
	goto L170
L170:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v943)+48))
	v971 = int32(2)
	v972 = base.I32_div_s(v965-v963, v971)
	v973 = v972 + v963
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v969+v973<<(uint(v971)%32))))
	if v977 == v919 {
		v1031 = v954
		goto L158
	} else {
		goto L172
	}
L171:
	;
	goto L165
L172:
	;
	v981 = F_TransactionIdPrecedes(m, v977, v919)
	mBase = m.M
	if v981 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v982 = v973 + int32(1)
	goto L175
L174:
	;
	v982 = v963
	goto L175
L175:
	;
	if v981 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v985 = v965
	goto L178
L177:
	;
	v985 = v973 - int32(1)
	goto L178
L178:
	;
	if v982 <= v985 {
		v963 = v982
		v965 = v985
		goto L170
	} else {
		goto L179
	}
L179:
	;
	goto L171
L180:
	;
	goto L164
L181:
	;
	v1009 = int32(2)
	v1010 = base.I32_div_s(v1004-v1002, v1009)
	v1011 = v1010 + v1002
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1000+v1011<<(uint(v1009)%32))))
	v1016 = base.B2i32(v1015 == v919)
	if v1015 == v919 {
		v1031 = v1016
		goto L158
	} else {
		goto L183
	}
L182:
	;
	v1031 = v1016
	goto L158
L183:
	;
	v1019 = base.B2i32(base.Ui32(v1015) < base.Ui32(v919))
	if base.Ui32(v1015) < base.Ui32(v919) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1020 = v1011 + int32(1)
	goto L186
L185:
	;
	v1020 = v1002
	goto L186
L186:
	;
	if base.Ui32(v1015) < base.Ui32(v919) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1023 = v1004
	goto L189
L188:
	;
	v1023 = v1011 - int32(1)
	goto L189
L189:
	;
	if v1020 <= v1023 {
		v1002 = v1020
		v1004 = v1023
		goto L181
	} else {
		goto L190
	}
L190:
	;
	goto L182
L191:
	;
	v1042 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v1042 == int32(0) {
		goto L140
	} else {
		goto L193
	}
L193:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v1046 + int32(4)
	F_errmsg_internal(m, int32(733114), v29)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(506907), int32(868), int32(219527))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	goto L140
L196:
	;
	v1209 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1209, float64(1))
	goto L140
L197:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v903)+16))
	v1059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+20)))
	if v1059&int32(6272) == int32(4096) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if base.Ui32(v1067) < base.Ui32(int32(3)) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v1064 = F_HeapTupleGetUpdateXid(m, v1058)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	v1067 = v1066
	goto L198
L202:
	;
	v1067 = v1064
	goto L198
L203:
	;
	if v1187 != 0 {
		goto L196
	} else {
		goto L243
	}
L204:
	;
	v1187 = int32(0)
	goto L203
L205:
	;
	goto L206
L206:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, _consts[35]))
	if v1078 == v1067 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1187 = int32(1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	if v1082 <= int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v1187 = v1179
	goto L203
L211:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v1086 == int32(0) {
		v1179 = int32(0)
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	v1150 = int32(0)
	v1152 = v1082 - int32(1)
	goto L233
L214:
	;
	v1091 = v1086
	goto L215
L215:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+20))
	if v1096 == int32(4) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1179 = int32(0)
	goto L210
L217:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+80))
	if v1143 != 0 {
		v1091 = v1143
		goto L215
	} else {
		goto L232
	}
L218:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1091)))
	if v1099 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1102 = int32(1)
	if v1067 == v1099 {
		v1179 = v1102
		goto L210
	} else {
		goto L220
	}
L220:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+52))
	v1106 = v1104 - int32(1)
	if v1106 < int32(0) {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	v1111 = int32(0)
	v1113 = v1106
	goto L222
L222:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+48))
	v1119 = int32(2)
	v1120 = base.I32_div_s(v1113-v1111, v1119)
	v1121 = v1120 + v1111
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1117+v1121<<(uint(v1119)%32))))
	if v1125 == v1067 {
		v1179 = v1102
		goto L210
	} else {
		goto L224
	}
L223:
	;
	goto L217
L224:
	;
	v1129 = F_TransactionIdPrecedes(m, v1125, v1067)
	mBase = m.M
	if v1129 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1130 = v1121 + int32(1)
	goto L227
L226:
	;
	v1130 = v1111
	goto L227
L227:
	;
	if v1129 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1133 = v1113
	goto L230
L229:
	;
	v1133 = v1121 - int32(1)
	goto L230
L230:
	;
	if v1130 <= v1133 {
		v1111 = v1130
		v1113 = v1133
		goto L222
	} else {
		goto L231
	}
L231:
	;
	goto L223
L232:
	;
	goto L216
L233:
	;
	v1157 = int32(2)
	v1158 = base.I32_div_s(v1152-v1150, v1157)
	v1159 = v1158 + v1150
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1148+v1159<<(uint(v1157)%32))))
	v1164 = base.B2i32(v1163 == v1067)
	if v1163 == v1067 {
		v1179 = v1164
		goto L210
	} else {
		goto L235
	}
L234:
	;
	v1179 = v1164
	goto L210
L235:
	;
	v1167 = base.B2i32(base.Ui32(v1163) < base.Ui32(v1067))
	if base.Ui32(v1163) < base.Ui32(v1067) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1168 = v1159 + int32(1)
	goto L238
L237:
	;
	v1168 = v1150
	goto L238
L238:
	;
	if base.Ui32(v1163) < base.Ui32(v1067) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1171 = v1152
	goto L241
L240:
	;
	v1171 = v1159 - int32(1)
	goto L241
L241:
	;
	if v1168 <= v1171 {
		v1150 = v1168
		v1152 = v1171
		goto L233
	} else {
		goto L242
	}
L242:
	;
	goto L234
L243:
	;
	v1190 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L5
	} else {
		goto L244
	}
L244:
	;
	if v1190 == int32(0) {
		goto L196
	} else {
		goto L245
	}
L245:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1194 + int32(4)
	F_errmsg_internal(m, int32(733162), v27+int32(-48))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(506907), int32(880), int32(219527))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L196
L248:
	;
	F_errmsg_internal(m, int32(99829), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(506907), int32(886), int32(219527))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L5
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
	v1229 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1229, float64(1))
	v1233 = m.G0
	v1235 = v1233 - int32(16)
	m.G0 = v1235
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+12)) = int32(0)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v903)+16))
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1239)+20)))
	v1241 = int32(768)
	if v1240&v1241 != v1241 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	v1247 = v1245
	goto L254
L253:
	;
	v1247 = int32(2)
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+4)) = v1247
	v1249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1235)+12)) = uint16(v1249)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v903)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1235)+8)) = v1251
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v1256 = int32(0)
	v1258 = F_hash_search(m, v1253, v1235+int32(4), v1256, v1256)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	if v1258 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+20))
	F_pfree(m, v1260)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	m.G0 = v1235 + int32(16)
	if v1258 == int32(0) {
		v764 = v900
		goto L104
	} else {
		goto L261
	}
L259:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v1269 = F_hash_search(m, v1263, v1235+int32(4), int32(2), v1235+int32(3))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1276 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1276, float64(1))
	v1280 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1280, float64(-1))
	v764 = v900
	goto L104
L262:
	;
	v1293 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_add(v1293, float64(1))
	if v731 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1297 = m.G0
	v1299 = v1297 - int32(16)
	m.G0 = v1299
	v1301 = int32(4536272)
	v1302 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v731)+32))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1304
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v731)+60))
	v1307 = F_heap_copytuple(m, v903)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L5
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = int64(17179869187)
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v903, v1433, v48, v50)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L5
	} else {
		goto L291
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1299))) = v1307
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+36)))
	if v1310 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+4))
	v1314 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1313)+12)))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	v1318 = F_heap_getattr_1(m, v1307, v1314, v1315, v1299+int32(8))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L5
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+52)))
	if v1321&int32(2) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1299)+4)) = v1318
	goto L269
L271:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+36)))
	if v1335 != int32(1) {
		v1346 = int32(0)
		goto L276
	} else {
		goto L277
	}
L272:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	v1333 = (v1326 + int32(31)) & int32(-8)
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1331 = F_GetMemoryChunkSpace(m, v1307)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	v1333 = v1331
	goto L271
L276:
	;
	F_tuplesort_puttuple_common(m, v731, v1299, v1346&int32(1), v1333)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L5
	} else {
		goto L279
	}
L277:
	;
	v1338 = int32(0)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v731)+44))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+24))
	if v1340 == v1338 {
		v1346 = v1338
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+8)))
	v1346 = v1343 ^ int32(1)
	goto L276
L279:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1302
	m.G0 = v1299 + int32(16)
	v1356 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	if base.F64_lt(base.F64_abs(v1356), float64(9.223372036854776e+18)) != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1361 = base.I64_trunc_f64_s(v1356)
	v1364 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1364 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	goto L282
L282:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1399 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L283:
	;
	v764 = v900
	goto L104
L284:
	;
	goto L283
L285:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1368 != int32(1) {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1371 = int32(4530932)
	v1373 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1374 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1373 + v1374
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	*(*int32)(unsafe.Add(mBase, uint32(v1364))) = v1377 + v1374
	*(*int64)(unsafe.Add(mBase, uint32(v1364+int32(24))+232)) = v1361
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	*(*int32)(unsafe.Add(mBase, uint32(v1364))) = v1385 + v1374
	v1391 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1391 - v1374
	goto L284
L287:
	;
	v764 = v900
	goto L104
L288:
	;
	goto L287
L289:
	;
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1403 != int32(1) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1406 = int32(4530932)
	v1408 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1409 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1408 + v1409
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = v1412 + v1409
	*(*int64)(unsafe.Add(mBase, uint32(v1399+int32(24))+232)) = int64(-9223372036854775807 - 1)
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = v1420 + v1409
	v1426 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1426 - v1409
	goto L288
L291:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1432)))
	if int32(0) < v1436 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1444 = int32(0)
	v1448 = v1436
	goto L295
L293:
	;
	goto L294
L294:
	;
	v1508 = F_heap_form_tuple(m, v1432, v48, v50)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L5
	} else {
		goto L301
	}
L295:
	;
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1432+int32(29)+v1444<<(uint(int32(4))%32)))))
	if v1471 == int32(1) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L294
L297:
	;
	v1475 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1444+v50))) = uint8(v1475)
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1432)))
	v1478 = v1477
	goto L299
L298:
	;
	v1478 = v1448
	goto L299
L299:
	;
	v1480 = v1444 + int32(1)
	if v1480 < v1478 {
		v1444 = v1480
		v1448 = v1478
		goto L295
	} else {
		goto L300
	}
L300:
	;
	goto L296
L301:
	;
	F_rewrite_heap_tuple(m, v71, v903, v1508)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	F_pfree(m, v1508)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L5
	} else {
		goto L303
	}
L303:
	;
	v1514 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	if base.F64_lt(base.F64_abs(v1514), float64(9.223372036854776e+18)) != 0 {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v1520
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v1520
	v1528 = int32(0)
	v1535 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1535 == v1528 {
		goto L309
	} else {
		goto L310
	}
L305:
	;
	v1518 = base.I64_trunc_f64_s(v1514)
	v1520 = v1518
	goto L304
L306:
	;
	goto L307
L307:
	;
	v1520 = int64(-9223372036854775807 - 1)
	goto L304
L308:
	;
	v764 = v900
	goto L104
L309:
	;
	goto L308
L310:
	;
	goto L311
L311:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1541&int32(1) == int32(0) {
		goto L309
	} else {
		goto L312
	}
L312:
	;
	v1546 = int32(4530932)
	v1548 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1549 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1548 + v1549
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1535)))
	*(*int32)(unsafe.Add(mBase, uint32(v1535))) = v1552 + v1549
	goto L314
L313:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1535)))
	v1683 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1535))) = v1682 + v1683
	v1686 = int32(4530932)
	v1688 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1688 - v1683
	goto L309
L314:
	;
	goto L316
L316:
	;
	goto L317
L317:
	;
	goto L321
L321:
	;
	v1647 = int32(0)
	v1650 = v1528
	goto L322
L322:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-8)+v1650<<(uint(int32(2))%32))))
	v1660 = int32(3)
	v1666 = *(*int64)(unsafe.Add(mBase, uint32(v27+int32(-32)+v1650<<(uint(v1660)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1535+int32(232)+v1659<<(uint(v1660)%32)))) = v1666
	v1668 = int32(1)
	v1671 = v1647 + v1668
	if v1671 != int32(2) {
		v1647 = v1671
		v1650 = v1650 + v1668
		goto L322
	} else {
		goto L324
	}
L323:
	;
	goto L313
L324:
	;
	goto L323
L325:
	;
	if v732 == int32(0) {
		goto L102
	} else {
		goto L326
	}
L326:
	;
	goto L103
L327:
	;
	goto L102
L328:
	;
	F_ExecDropSingleTupleTableSlot(m, v746)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L5
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	if v731 != 0 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	goto L330
L332:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1716 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L334
L334:
	;
	v2002 = m.G0
	v2004 = v2002 - int32(48)
	m.G0 = v2004
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	F_hash_seq_init(m, v2004+int32(8), v2008)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L5
	} else {
		goto L379
	}
L335:
	;
	F_tuplesort_performsort(m, v731)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L5
	} else {
		goto L339
	}
L336:
	;
	goto L335
L337:
	;
	v1720 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1720 != int32(1) {
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1723 = int32(4530932)
	v1725 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1726 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1725 + v1726
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1716)))
	*(*int32)(unsafe.Add(mBase, uint32(v1716))) = v1729 + v1726
	*(*int64)(unsafe.Add(mBase, uint32(v1716+int32(8))+232)) = int64(3)
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1716)))
	*(*int32)(unsafe.Add(mBase, uint32(v1716))) = v1737 + v1726
	v1743 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1743 - v1726
	goto L336
L339:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1753 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1809 = float64(0)
	goto L344
L341:
	;
	goto L340
L342:
	;
	v1757 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1757 != int32(1) {
		goto L341
	} else {
		goto L343
	}
L343:
	;
	v1760 = int32(4530932)
	v1762 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1763 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1762 + v1763
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1753)))
	*(*int32)(unsafe.Add(mBase, uint32(v1753))) = v1766 + v1763
	*(*int64)(unsafe.Add(mBase, uint32(v1753+int32(8))+232)) = int64(4)
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1753)))
	*(*int32)(unsafe.Add(mBase, uint32(v1753))) = v1774 + v1763
	v1780 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1780 - v1763
	goto L341
L344:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v1812 != 0 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	F_tuplesort_end(m, v731)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L5
	} else {
		goto L378
	}
L346:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L5
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v1815 = F_tuplesort_getheaptuple(m, v731)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L5
	} else {
		goto L350
	}
L349:
	;
	goto L348
L350:
	;
	if v1815 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v1815, v1818, v48, v50)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L5
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	goto L345
L354:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1817)))
	if int32(0) < v1821 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1829 = int32(0)
	v1833 = v1821
	goto L358
L356:
	;
	goto L357
L357:
	;
	v1893 = F_heap_form_tuple(m, v1817, v48, v50)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L5
	} else {
		goto L364
	}
L358:
	;
	v1856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1817+int32(29)+v1829<<(uint(int32(4))%32)))))
	if v1856 == int32(1) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	goto L357
L360:
	;
	v1860 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1829+v50))) = uint8(v1860)
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1817)))
	v1863 = v1862
	goto L362
L361:
	;
	v1863 = v1833
	goto L362
L362:
	;
	v1865 = v1829 + int32(1)
	if v1865 < v1863 {
		v1829 = v1865
		v1833 = v1863
		goto L358
	} else {
		goto L363
	}
L363:
	;
	goto L359
L364:
	;
	F_rewrite_heap_tuple(m, v71, v1815, v1893)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	F_pfree(m, v1893)
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L5
	} else {
		goto L366
	}
L366:
	;
	v1900 = base.F64_add(v1809, float64(1))
	if base.F64_lt(base.F64_abs(v1900), float64(9.223372036854776e+18)) != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1905 = base.I64_trunc_f64_s(v1900)
	v1908 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1908 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L368:
	;
	goto L369
L369:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v1943 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L370:
	;
	v1809 = v1900
	goto L344
L371:
	;
	goto L370
L372:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1912 != int32(1) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1915 = int32(4530932)
	v1917 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1918 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1917 + v1918
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1908)))
	*(*int32)(unsafe.Add(mBase, uint32(v1908))) = v1921 + v1918
	*(*int64)(unsafe.Add(mBase, uint32(v1908+int32(32))+232)) = v1905
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1908)))
	*(*int32)(unsafe.Add(mBase, uint32(v1908))) = v1929 + v1918
	v1935 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1935 - v1918
	goto L371
L374:
	;
	v1809 = v1900
	goto L344
L375:
	;
	goto L374
L376:
	;
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v1947 != int32(1) {
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1950 = int32(4530932)
	v1952 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1953 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1952 + v1953
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1943)))
	*(*int32)(unsafe.Add(mBase, uint32(v1943))) = v1956 + v1953
	*(*int64)(unsafe.Add(mBase, uint32(v1943+int32(32))+232)) = int64(-9223372036854775807 - 1)
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1943)))
	*(*int32)(unsafe.Add(mBase, uint32(v1943))) = v1964 + v1953
	v1970 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1970 - v1953
	goto L375
L378:
	;
	goto L334
L379:
	;
	v2013 = F_hash_seq_search(m, v2004+int32(8))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L5
	} else {
		goto L380
	}
L380:
	;
	if v2013 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2015 = v2013
	goto L384
L382:
	;
	goto L383
L383:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v2080 != 0 {
		goto L389
	} else {
		goto L390
	}
L384:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+20))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+16))
	v2043 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2042)+16)) = uint16(v2043)
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+12)) = int32(-1)
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+20))
	F_raw_heap_insert(m, v71, v2047)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L5
	} else {
		goto L386
	}
L385:
	;
	goto L383
L386:
	;
	v2052 = F_hash_seq_search(m, v2004+int32(8))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L5
	} else {
		goto L387
	}
L387:
	;
	if v2052 != 0 {
		v2015 = v2052
		goto L384
	} else {
		goto L388
	}
L388:
	;
	goto L385
L389:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	F_smgr_bulk_write(m, v2081, v2082, v2080, int32(1))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L5
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	F_smgr_bulk_finish(m, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L5
	} else {
		goto L393
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(0)
	goto L391
L393:
	;
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+20)))
	if v2091 != int32(1) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	F_MemoryContextDelete(m, v2196)
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L5
	} else {
		goto L420
	}
L395:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v71)+68))
	if v2094 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	F_logical_heap_rewrite_flush_mappings(m, v71)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L5
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v71)+64))
	F_hash_seq_init(m, v2004+int32(28), v2099)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L5
	} else {
		goto L400
	}
L399:
	;
	goto L398
L400:
	;
	v2104 = F_hash_seq_search(m, v2004+int32(28))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L5
	} else {
		goto L401
	}
L401:
	;
	if v2104 == int32(0) {
		goto L394
	} else {
		goto L402
	}
L402:
	;
	v2108 = v2104
	goto L403
L403:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2108)+4))
	v2136 = F_FileSync(m, v2134, int32(167772197))
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L5
	} else {
		goto L406
	}
L404:
	;
	goto L394
L405:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2108)+4))
	F_FileClose(m, v2163)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L5
	} else {
		goto L417
	}
L406:
	;
	if v2136 == int32(0) {
		goto L405
	} else {
		goto L407
	}
L407:
	;
	v2143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
	if v2143 != 0 {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	v2146 = F_errstart(m, v2144, int32(0))
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L5
	} else {
		goto L412
	}
L409:
	;
	v2144 = int32(21)
	goto L411
L410:
	;
	v2144 = int32(23)
	goto L411
L411:
	;
	goto L408
L412:
	;
	if v2146 == int32(0) {
		goto L405
	} else {
		goto L413
	}
L413:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L5
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2004))) = v2108 + int32(28)
	F_errmsg(m, int32(305761), v2004)
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(507593), int32(925), int32(357323))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L5
	} else {
		goto L416
	}
L416:
	;
	goto L405
L417:
	;
	v2168 = F_hash_seq_search(m, v2004+int32(28))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L5
	} else {
		goto L418
	}
L418:
	;
	if v2168 != 0 {
		v2108 = v2168
		goto L403
	} else {
		goto L419
	}
L419:
	;
	goto L404
L420:
	;
	m.G0 = v2004 + int32(48)
	F_pfree(m, v48)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L5
	} else {
		goto L421
	}
L421:
	;
	F_pfree(m, v50)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L5
	} else {
		goto L422
	}
L422:
	;
	m.G0 = v29 - int32(-64)
	return
L423:
	;
	F_errmsg_internal(m, int32(343590), int32(0))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L5
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(333904), int32(1034), int32(86386))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L5
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
