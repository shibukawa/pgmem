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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
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
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v327 int32
	_ = v327
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v400 int32
	_ = v400
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v520 int64
	_ = v520
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int64
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1026 int32
	_ = v1026
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1174 int32
	_ = v1174
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1204 float64
	_ = v1204
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 float64
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1269 float64
	_ = v1269
	var v1273 float64
	_ = v1273
	var v1277 float64
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1285 float64
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1349 float64
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1376 int32
	_ = v1376
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 float64
	_ = v1466
	var v1467 int64
	_ = v1467
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1613 int64
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1686 int32
	_ = v1686
	var v1692 int32
	_ = v1692
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1759 float64
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1846 float64
	_ = v1846
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
	var v1873 int32
	_ = v1873
	var v1879 int32
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2121 int32
	_ = v2121
	v11 = int32(0)
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v33 = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v34) < base.Ui32(int32(_a_F_heapam_relation_copy_for_cluster_0)) {
		v43 = v33
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v47 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v38 == int32(99) {
		v43 = v33
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v41 = F_isTempToastNamespace(m, v38)
	mBase = m.M
	v43 = v41
	goto L2
L5:
	;
	return
L6:
	;
	v49 = F_palloc(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v53 = m.G0
	v55 = v53 - int32(112)
	m.G0 = v55
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v63 = F_AllocSetContextCreateInternal(m, v58, int32(_a_F_heapam_relation_copy_for_cluster_1), int32(0), int32(_a_F_heapam_relation_copy_for_cluster_2), int32(_a_F_heapam_relation_copy_for_cluster_3))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v65 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v63
	v70 = F_palloc0(m, int32(72))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = l0
	v77 = F_RelationGetNumberOfBlocksInFork(m, l1, v72)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+40)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v77
	v85 = F_smgr_bulk_start_rel(m, l1, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v55)+28)) = int64(103079215116)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+52)) = v90
	v95 = v55 + int32(12)
	v97 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_5), int32(128), v95, int32(1064))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+56)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = int32(20)
	v105 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_6), int32(128), v95, int32(1064))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+60)) = v105
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v66
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[1]))
	if v111 < int32(2) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v55 + int32(112)
	if l3 != 0 {
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)) = uint8(v184)
	goto L14
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+48))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+118)))
	if v116 != int32(112) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+56))
	goto L19
L18:
	;
	v141 = v55 + int32(60)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[2]))
	v147 = F_LWLockAcquire(m, v143+int32(512), int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L26
	}
L19:
	;
	if base.Ui32(v119) < base.Ui32(int32(_a_F_heapam_relation_copy_for_cluster_0)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)) = uint8(v122)
	goto L18
L21:
	;
	goto L22
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+180))
	if v125 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+119)))
	switch v129 - int32(109) {
	case 0, 5:
		goto L24
	default:
		goto L15
	}
L24:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+104)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)) = uint8(v132)
	if v132 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	if v141 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[3]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v151
	goto L29
L28:
	;
	goto L29
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[2]))
	F_LWLockRelease(m, v154+int32(512))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
	if v159 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(20)))) = uint8(v162)
	goto L14
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v159
	v165 = F_GetXLogInsertRecPtr(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v70)+48)) = v165
	*(*int64)(unsafe.Add(mBase, uint32(v55)+80)) = int64(4535485464580)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+104)) = v172
	v179 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_7), int32(128), v55-int32(-64), int32(1064))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+64)) = v179
	goto L14
L36:
	;
	v739 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L5
	} else {
		goto L100
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = int64(8589934593)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = int64(2)
	v520 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+56)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v520
	goto L83
L38:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v437 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L39:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	v194 = m.G0
	v196 = v194 - int32(16)
	m.G0 = v196
	v198 = int32(0)
	v200 = F_tuplesort_begin_common(m, v193, v198, v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
	v202 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v205
	v208 = F_palloc0(m, int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[6])))
	if v211 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l2)+192))
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = int32(1825)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+40)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v200)+60)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = int32(1826)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(1827)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = int32(1828)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = int32(1829)
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = int32(1830)
	v251 = F_BuildIndexInfo(m, l2)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L50
	}
L45:
	;
	v216 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v216 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+8)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(v196)+4)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v221
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_8), v196)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_9), int32(275), int32(_a_F_heapam_relation_copy_for_cluster_10))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v251
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+12)))
	v255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+36)) = uint8(base.B2i32(v254 != v255))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v30
	v260 = F__bt_mkscankey(m, l2, v255)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+76))
	if v263 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v264 = F_CreateExecutorState(m)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v200)+40))
	v282 = F_palloc0(m, v279*int32(36))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L61
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v264
	v268 = F_MakeTupleTableSlot(m, v30, int32(_a_F_heapam_relation_copy_for_cluster_11))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+152))
	if v271 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v274 = v271
	goto L59
L58:
	;
	v272 = F_MakePerTupleExprContext(m, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+4)) = v268
	goto L54
L60:
	;
	v274 = v272
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+44)) = v282
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v200)+40))
	if v286 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_pfree(m, v260)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L70
	}
L63:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v260)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	v298 = int32(base.Ui32(v294)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+9)) = uint8(v298)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+10)) = uint16(v300)
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+20)) = uint8(v302)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v304&int32(16777216))>>(uint(int32(24))%32)), v282)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v200)+40))
	if v311 < int32(2) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v327 = int32(1)
	goto L66
L66:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v200)+44))
	v344 = v341 + v327*int32(36)
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v346
	v350 = v260 + int32(16) + v327*int32(48)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v357 = int32(base.Ui32(v353)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+9)) = uint8(v357)
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+4)))
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+20)) = uint8(v360)
	*(*uint16)(unsafe.Add(mBase, uint32(v344)+10)) = uint16(v359)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v363&int32(16777216))>>(uint(int32(24))%32)), v344)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L68
	}
L67:
	;
	goto L62
L68:
	;
	v371 = v327 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v200)+40))
	if v371 < v372 {
		v327 = v371
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v203
	m.G0 = v196 + int32(16)
	v432 = v200
	goto L38
L71:
	;
	v432 = int32(0)
	goto L38
L72:
	;
	v472 = int32(0)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+8))
	v478 = m.T0[v477].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_heapam_relation_copy_for_cluster_12), v472, v472, v472, int32(449))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v441&int32(1) == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v446 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v449 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v448 + v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v452 + v449
	*(*int64)(unsafe.Add(mBase, uint32(v437+int32(8))+232)) = int64(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v460 + v449
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v466 - v449
	goto L73
L76:
	;
	v480 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v478)+36)))
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v483 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v722 = v432
	v727 = v478
	v732 = v11
	goto L36
L78:
	;
	goto L77
L79:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v487&int32(1) == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v492 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v495 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v494 + v495
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v498 + v495
	*(*int64)(unsafe.Add(mBase, uint32(v483+int32(40))+232)) = v480
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v506 + v495
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v512 - v495
	goto L78
L81:
	;
	v701 = int32(0)
	v704 = F_index_beginscan(m, l0, l2, int32(_a_F_heapam_relation_copy_for_cluster_12), v701, v701, v701)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L98
	}
L82:
	;
	goto L81
L83:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v536 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v540&int32(1) == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v545 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v548 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v547 + v548
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v551 + v548
	goto L87
L86:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v682 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v681 + v682
	v685 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v687 - v682
	goto L82
L87:
	;
	goto L89
L89:
	;
	goto L90
L90:
	;
	v646 = int32(0)
	v649 = int32(0)
	goto L95
L95:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8)+v649<<(uint(int32(2))%32))))
	v659 = int32(3)
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v26+int32(-32)+v649<<(uint(v659)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v536+int32(232)+v658<<(uint(v659)%32)))) = v665
	v667 = int32(1)
	v670 = v646 + v667
	if v670 != int32(2) {
		v646 = v670
		v649 = v649 + v667
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
	v706 = int32(0)
	F_index_rescan(m, v704, v706, v706, v706, v706)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v722 = v11
	v727 = v11
	v732 = v704
	goto L36
L100:
	;
	v753 = int32(-1)
	goto L104
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L5
	} else {
		goto L405
	}
L102:
	;
	if v739 != 0 {
		goto L317
	} else {
		goto L318
	}
L103:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+188))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1653)+12))
	m.T0[v1654].(func(*base.Module, int32))(m, v727)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L5
	} else {
		goto L316
	}
L104:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9]))
	if v767 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	F_index_endscan(m, v732)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L5
	} else {
		goto L314
	}
L106:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L5
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v732 != 0 {
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
	v896 = int32(0)
	v898 = F_ExecFetchSlotHeapTuple(m, v739, v896, v896)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L138
	}
L112:
	;
	v771 = F_index_getnext_slot(m, v732, int32(1), v739)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v739)+36)) = v792
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	if v795 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	if v771 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732)+72)))
	if v775 != int32(1) {
		v895 = v753
		goto L111
	} else {
		goto L117
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_14), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(798), int32(_a_F_heapam_relation_copy_for_cluster_16))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
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
	v797 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[11])))
	if v797&int32(1) == int32(0) {
		goto L101
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+188))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v804)+20))
	v806 = m.T0[v805].(func(*base.Module, int32, int32, int32) int32)(m, v727, int32(1), v739)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L5
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	if v806 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v811 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v727)+36)))
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v814 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v727)+52))
	if v753 == v847 {
		v895 = v753
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
	v818 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v818&int32(1) == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v823 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v826 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v825 + v826
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	*(*int32)(unsafe.Add(mBase, uint32(v814))) = v829 + v826
	*(*int64)(unsafe.Add(mBase, uint32(v814+int32(48))+232)) = v811
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	*(*int32)(unsafe.Add(mBase, uint32(v814))) = v837 + v826
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v843 - v826
	goto L130
L133:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v727)+36))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v727)+40))
	v854 = base.I32_rem_u_s(v847+v850-v852, v850)
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v860 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v727)+52))
	v895 = v893
	goto L111
L135:
	;
	goto L134
L136:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v864&int32(1) == int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v869 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v871 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v872 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v871 + v872
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v860)))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = v875 + v872
	*(*int64)(unsafe.Add(mBase, uint32(v860+int32(48))+232)) = base.I64_extend_i32_u(v854 + int32(1))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v860)))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = v883 + v872
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v889 - v872
	goto L135
L138:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v739)+68))
	F_LockBuffer(m, v900, int32(1))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	v904 = F_HeapTupleSatisfiesVacuum(m, v898, l4, v900)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L5
	} else {
		goto L146
	}
L140:
	;
	F_LockBuffer(m, v900, int32(0))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L5
	} else {
		goto L262
	}
L141:
	;
	v1277 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1277, float64(1))
	goto L140
L142:
	;
	F_LockBuffer(m, v900, int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L5
	} else {
		goto L251
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L5
	} else {
		goto L248
	}
L144:
	;
	if v43 != 0 {
		goto L196
	} else {
		goto L197
	}
L145:
	;
	if v43 != 0 {
		goto L140
	} else {
		goto L147
	}
L146:
	;
	switch v904 {
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
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v898)+16))
	v907 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v906)+20)))
	v908 = int32(768)
	if v907&v908 != v908 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v914 = v912
	goto L150
L149:
	;
	v914 = int32(2)
	goto L150
L150:
	;
	if base.Ui32(v914) < base.Ui32(int32(3)) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v1034 != 0 {
		goto L140
	} else {
		goto L191
	}
L152:
	;
	v1034 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v925 == v914 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1034 = int32(1)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v929 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1034 = v1026
	goto L151
L159:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v933 == int32(0) {
		v1026 = int32(0)
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	v997 = int32(0)
	v999 = v929 - int32(1)
	goto L181
L162:
	;
	v938 = v933
	goto L163
L163:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v938)+20))
	if v943 == int32(4) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1026 = int32(0)
	goto L158
L165:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v938)+80))
	if v990 != 0 {
		v938 = v990
		goto L163
	} else {
		goto L180
	}
L166:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v938)))
	if v946 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v949 = int32(1)
	if v914 == v946 {
		v1026 = v949
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v938)+52))
	v953 = v951 - int32(1)
	if v953 < int32(0) {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v958 = int32(0)
	v960 = v953
	goto L170
L170:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v938)+48))
	v966 = int32(2)
	v967 = base.I32_div_s(v960-v958, v966)
	v968 = v967 + v958
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v964+v968<<(uint(v966)%32))))
	if v972 == v914 {
		v1026 = v949
		goto L158
	} else {
		goto L172
	}
L171:
	;
	goto L165
L172:
	;
	v976 = F_TransactionIdPrecedes(m, v972, v914)
	mBase = m.M
	if v976 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v977 = v968 + int32(1)
	goto L175
L174:
	;
	v977 = v958
	goto L175
L175:
	;
	if v976 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v980 = v960
	goto L178
L177:
	;
	v980 = v968 - int32(1)
	goto L178
L178:
	;
	if v977 <= v980 {
		v958 = v977
		v960 = v980
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
	v1004 = int32(2)
	v1005 = base.I32_div_s(v999-v997, v1004)
	v1006 = v1005 + v997
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v995+v1006<<(uint(v1004)%32))))
	v1011 = base.B2i32(v1010 == v914)
	if v1010 == v914 {
		v1026 = v1011
		goto L158
	} else {
		goto L183
	}
L182:
	;
	v1026 = v1011
	goto L158
L183:
	;
	v1014 = base.B2i32(base.Ui32(v1010) < base.Ui32(v914))
	if base.Ui32(v1010) < base.Ui32(v914) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1015 = v1006 + int32(1)
	goto L186
L185:
	;
	v1015 = v997
	goto L186
L186:
	;
	if base.Ui32(v1010) < base.Ui32(v914) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1018 = v999
	goto L189
L188:
	;
	v1018 = v1006 - int32(1)
	goto L189
L189:
	;
	if v1015 <= v1018 {
		v997 = v1015
		v999 = v1018
		goto L181
	} else {
		goto L190
	}
L190:
	;
	goto L182
L191:
	;
	v1037 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v1037 == int32(0) {
		goto L140
	} else {
		goto L193
	}
L193:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v1041 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_17), v28)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(868), int32(_a_F_heapam_relation_copy_for_cluster_16))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	goto L140
L196:
	;
	v1204 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1204, float64(1))
	goto L140
L197:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v898)+16))
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053)+20)))
	if v1054&int32(_a_F_heapam_relation_copy_for_cluster_18) == int32(_a_F_heapam_relation_copy_for_cluster_19) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if base.Ui32(v1062) < base.Ui32(int32(3)) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v1059 = F_HeapTupleGetUpdateXid(m, v1053)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L5
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	v1062 = v1061
	goto L198
L202:
	;
	v1062 = v1059
	goto L198
L203:
	;
	if v1182 != 0 {
		goto L196
	} else {
		goto L243
	}
L204:
	;
	v1182 = int32(0)
	goto L203
L205:
	;
	goto L206
L206:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v1073 == v1062 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1182 = int32(1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v1077 <= int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v1182 = v1174
	goto L203
L211:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v1081 == int32(0) {
		v1174 = int32(0)
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	v1145 = int32(0)
	v1147 = v1077 - int32(1)
	goto L233
L214:
	;
	v1086 = v1081
	goto L215
L215:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+20))
	if v1091 == int32(4) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1174 = int32(0)
	goto L210
L217:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+80))
	if v1138 != 0 {
		v1086 = v1138
		goto L215
	} else {
		goto L232
	}
L218:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	if v1094 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1097 = int32(1)
	if v1062 == v1094 {
		v1174 = v1097
		goto L210
	} else {
		goto L220
	}
L220:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+52))
	v1101 = v1099 - int32(1)
	if v1101 < int32(0) {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	v1106 = int32(0)
	v1108 = v1101
	goto L222
L222:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+48))
	v1114 = int32(2)
	v1115 = base.I32_div_s(v1108-v1106, v1114)
	v1116 = v1115 + v1106
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1112+v1116<<(uint(v1114)%32))))
	if v1120 == v1062 {
		v1174 = v1097
		goto L210
	} else {
		goto L224
	}
L223:
	;
	goto L217
L224:
	;
	v1124 = F_TransactionIdPrecedes(m, v1120, v1062)
	mBase = m.M
	if v1124 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1125 = v1116 + int32(1)
	goto L227
L226:
	;
	v1125 = v1106
	goto L227
L227:
	;
	if v1124 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1128 = v1108
	goto L230
L229:
	;
	v1128 = v1116 - int32(1)
	goto L230
L230:
	;
	if v1125 <= v1128 {
		v1106 = v1125
		v1108 = v1128
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
	v1152 = int32(2)
	v1153 = base.I32_div_s(v1147-v1145, v1152)
	v1154 = v1153 + v1145
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1143+v1154<<(uint(v1152)%32))))
	v1159 = base.B2i32(v1158 == v1062)
	if v1158 == v1062 {
		v1174 = v1159
		goto L210
	} else {
		goto L235
	}
L234:
	;
	v1174 = v1159
	goto L210
L235:
	;
	v1162 = base.B2i32(base.Ui32(v1158) < base.Ui32(v1062))
	if base.Ui32(v1158) < base.Ui32(v1062) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1163 = v1154 + int32(1)
	goto L238
L237:
	;
	v1163 = v1145
	goto L238
L238:
	;
	if base.Ui32(v1158) < base.Ui32(v1062) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1166 = v1147
	goto L241
L240:
	;
	v1166 = v1154 - int32(1)
	goto L241
L241:
	;
	if v1163 <= v1166 {
		v1145 = v1163
		v1147 = v1166
		goto L233
	} else {
		goto L242
	}
L242:
	;
	goto L234
L243:
	;
	v1185 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L5
	} else {
		goto L244
	}
L244:
	;
	if v1185 == int32(0) {
		goto L196
	} else {
		goto L245
	}
L245:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v1189 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_20), v26+int32(-48))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(880), int32(_a_F_heapam_relation_copy_for_cluster_16))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L196
L248:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_21), int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(886), int32(_a_F_heapam_relation_copy_for_cluster_16))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
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
	v1224 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1224, float64(1))
	v1228 = m.G0
	v1230 = v1228 - int32(16)
	m.G0 = v1230
	*(*int32)(unsafe.Add(mBase, uint32(v1230)+12)) = int32(0)
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v898)+16))
	v1235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1234)+20)))
	v1236 = int32(768)
	if v1235&v1236 != v1236 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1234)))
	v1242 = v1240
	goto L254
L253:
	;
	v1242 = int32(2)
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1230)+4)) = v1242
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v898)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1230)+8)) = v1244
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v898)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1230)+12)) = uint16(v1246)
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v1250 = v1230 + int32(4)
	v1251 = int32(0)
	v1253 = F_hash_search(m, v1248, v1250, v1251, v1251)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	if v1253 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+20))
	F_pfree(m, v1255)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	m.G0 = v1230 + int32(16)
	if v1253 == int32(0) {
		v753 = v895
		goto L104
	} else {
		goto L261
	}
L259:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v1262 = F_hash_search(m, v1258, v1250, int32(2), v1230+int32(3))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1269 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1269, float64(1))
	v1273 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1273, float64(-1))
	v753 = v895
	goto L104
L262:
	;
	v1285 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_add(v1285, float64(1))
	if v722 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1289 = m.G0
	v1291 = v1289 - int32(16)
	m.G0 = v1291
	v1293 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v722)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1296
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v722)+60))
	v1299 = F_heap_copytuple(m, v898)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L5
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = int64(17179869187)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v898, v1389, v47, v49)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L5
	} else {
		goto L284
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1291))) = v1299
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+36)))
	if v1302 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+4))
	v1306 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1305)+12)))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1298)))
	v1310 = F_heap_getattr_1(m, v1299, v1306, v1307, v1291+int32(8))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L5
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+52)))
	if v1313&int32(2) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+4)) = v1310
	goto L269
L271:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+36)))
	if v1327 != int32(1) {
		v1338 = int32(0)
		goto L276
	} else {
		goto L277
	}
L272:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1299)))
	v1325 = (v1318 + int32(31)) & int32(-8)
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1323 = F_GetMemoryChunkSpace(m, v1299)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	v1325 = v1323
	goto L271
L276:
	;
	F_tuplesort_puttuple_common(m, v722, v1291, v1338&int32(1), v1325)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L5
	} else {
		goto L279
	}
L277:
	;
	v1330 = int32(0)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v722)+44))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+24))
	if v1332 == v1330 {
		v1338 = v1330
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+8)))
	v1338 = v1335 ^ int32(1)
	goto L276
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1294
	m.G0 = v1291 + int32(16)
	v1349 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1353 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v753 = v895
	goto L104
L281:
	;
	goto L280
L282:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1357&int32(1) == int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1362 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1364 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1365 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1364 + v1365
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1353)))
	*(*int32)(unsafe.Add(mBase, uint32(v1353))) = v1368 + v1365
	*(*int64)(unsafe.Add(mBase, uint32(v1353+int32(24))+232)) = base.I64_trunc_sat_f64_s(v1349)
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1353)))
	*(*int32)(unsafe.Add(mBase, uint32(v1353))) = v1376 + v1365
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1382 - v1365
	goto L281
L284:
	;
	v1392 = int32(0)
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1388)))
	if v1392 < v1393 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1398 = v1392
	v1401 = v1393
	goto L288
L286:
	;
	goto L287
L287:
	;
	v1460 = F_heap_form_tuple(m, v1388, v47, v49)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L5
	} else {
		goto L294
	}
L288:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1388+v1398<<(uint(int32(4))%32))+29)))
	if v1424 == int32(1) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	goto L287
L290:
	;
	v1428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1398+v49))) = uint8(v1428)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1388)))
	v1431 = v1430
	goto L292
L291:
	;
	v1431 = v1401
	goto L292
L292:
	;
	v1433 = v1398 + int32(1)
	if v1433 < v1431 {
		v1398 = v1433
		v1401 = v1431
		goto L288
	} else {
		goto L293
	}
L293:
	;
	goto L289
L294:
	;
	F_rewrite_heap_tuple(m, v70, v898, v1460)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	F_pfree(m, v1460)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v1466 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	v1467 = base.I64_trunc_sat_f64_s(v1466)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v1467
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v1467
	goto L299
L297:
	;
	v753 = v895
	goto L104
L298:
	;
	goto L297
L299:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1484 == int32(0) {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1488&int32(1) == int32(0) {
		goto L298
	} else {
		goto L301
	}
L301:
	;
	v1493 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1496 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1495 + v1496
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1484)))
	*(*int32)(unsafe.Add(mBase, uint32(v1484))) = v1499 + v1496
	goto L303
L302:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1484)))
	v1630 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1484))) = v1629 + v1630
	v1633 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1635 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1635 - v1630
	goto L298
L303:
	;
	goto L305
L305:
	;
	goto L306
L306:
	;
	v1594 = int32(0)
	v1597 = int32(0)
	goto L311
L311:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8)+v1597<<(uint(int32(2))%32))))
	v1607 = int32(3)
	v1613 = *(*int64)(unsafe.Add(mBase, uint32(v26+int32(-32)+v1597<<(uint(v1607)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1484+int32(232)+v1606<<(uint(v1607)%32)))) = v1613
	v1615 = int32(1)
	v1618 = v1594 + v1615
	if v1618 != int32(2) {
		v1594 = v1618
		v1597 = v1597 + v1615
		goto L311
	} else {
		goto L313
	}
L312:
	;
	goto L302
L313:
	;
	goto L312
L314:
	;
	if v727 == int32(0) {
		goto L102
	} else {
		goto L315
	}
L315:
	;
	goto L103
L316:
	;
	goto L102
L317:
	;
	F_ExecDropSingleTupleTableSlot(m, v739)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L5
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	if v722 != 0 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	goto L319
L321:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1663 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L322:
	;
	goto L323
L323:
	;
	v1910 = m.G0
	v1912 = v1910 - int32(48)
	m.G0 = v1912
	v1915 = v1912 + int32(8)
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	F_hash_seq_init(m, v1915, v1916)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L5
	} else {
		goto L361
	}
L324:
	;
	F_tuplesort_performsort(m, v722)
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L5
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1667&int32(1) == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1672 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1674 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1675 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1674 + v1675
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	*(*int32)(unsafe.Add(mBase, uint32(v1663))) = v1678 + v1675
	*(*int64)(unsafe.Add(mBase, uint32(v1663+int32(8))+232)) = int64(3)
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	*(*int32)(unsafe.Add(mBase, uint32(v1663))) = v1686 + v1675
	v1692 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1692 - v1675
	goto L325
L328:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1702 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1759 = float64(0)
	goto L333
L330:
	;
	goto L329
L331:
	;
	v1706 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1706&int32(1) == int32(0) {
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1711 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1713 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1714 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1713 + v1714
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1702)))
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = v1717 + v1714
	*(*int64)(unsafe.Add(mBase, uint32(v1702+int32(8))+232)) = int64(4)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1702)))
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = v1725 + v1714
	v1731 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1731 - v1714
	goto L330
L333:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9]))
	if v1761 != 0 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	F_tuplesort_end(m, v722)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L5
	} else {
		goto L360
	}
L335:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L5
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1764 = F_tuplesort_getheaptuple(m, v722)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L5
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	if v1764 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v1764, v1767, v47, v49)
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L5
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	goto L334
L343:
	;
	v1770 = int32(0)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1766)))
	if v1770 < v1771 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1776 = v1770
	v1779 = v1771
	goto L347
L345:
	;
	goto L346
L346:
	;
	v1838 = F_heap_form_tuple(m, v1766, v47, v49)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L5
	} else {
		goto L353
	}
L347:
	;
	v1802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766+v1776<<(uint(int32(4))%32))+29)))
	if v1802 == int32(1) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L346
L349:
	;
	v1806 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1776+v49))) = uint8(v1806)
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1766)))
	v1809 = v1808
	goto L351
L350:
	;
	v1809 = v1779
	goto L351
L351:
	;
	v1811 = v1776 + int32(1)
	if v1811 < v1809 {
		v1776 = v1811
		v1779 = v1809
		goto L347
	} else {
		goto L352
	}
L352:
	;
	goto L348
L353:
	;
	F_rewrite_heap_tuple(m, v70, v1764, v1838)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	F_pfree(m, v1838)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	v1846 = base.F64_add(v1759, float64(1))
	v1850 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1850 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1759 = v1846
	goto L333
L357:
	;
	goto L356
L358:
	;
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1854&int32(1) == int32(0) {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1859 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1861 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1862 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1861 + v1862
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1850)))
	*(*int32)(unsafe.Add(mBase, uint32(v1850))) = v1865 + v1862
	*(*int64)(unsafe.Add(mBase, uint32(v1850+int32(32))+232)) = base.I64_trunc_sat_f64_s(v1846)
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1850)))
	*(*int32)(unsafe.Add(mBase, uint32(v1850))) = v1873 + v1862
	v1879 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1879 - v1862
	goto L357
L360:
	;
	goto L323
L361:
	;
	v1919 = F_hash_seq_search(m, v1915)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	if v1919 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1921 = v1919
	goto L366
L364:
	;
	goto L365
L365:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v1984 != 0 {
		goto L371
	} else {
		goto L372
	}
L366:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+20))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+16))
	v1948 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1947)+16)) = uint16(v1948)
	*(*int32)(unsafe.Add(mBase, uint32(v1947)+12)) = int32(-1)
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+20))
	F_raw_heap_insert(m, v70, v1952)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L5
	} else {
		goto L368
	}
L367:
	;
	goto L365
L368:
	;
	v1957 = F_hash_seq_search(m, v1912+int32(8))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L5
	} else {
		goto L369
	}
L369:
	;
	if v1957 != 0 {
		v1921 = v1957
		goto L366
	} else {
		goto L370
	}
L370:
	;
	goto L367
L371:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	F_smgr_bulk_write(m, v1985, v1986, v1984, int32(1))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L5
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	F_smgr_bulk_finish(m, v1992)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L5
	} else {
		goto L375
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(0)
	goto L373
L375:
	;
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)))
	if v1995 != int32(1) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	F_MemoryContextDelete(m, v2096)
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L5
	} else {
		goto L402
	}
L377:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if v1998 != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	F_logical_heap_rewrite_flush_mappings(m, v70)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L5
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v2002 = v1912 + int32(28)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v70)+64))
	F_hash_seq_init(m, v2002, v2003)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L5
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	v2006 = F_hash_seq_search(m, v2002)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	if v2006 == int32(0) {
		goto L376
	} else {
		goto L384
	}
L384:
	;
	v2010 = v2006
	goto L385
L385:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+4))
	v2037 = F_FileSync(m, v2035, int32(167772197))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L5
	} else {
		goto L388
	}
L386:
	;
	goto L376
L387:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+4))
	F_FileClose(m, v2064)
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L5
	} else {
		goto L399
	}
L388:
	;
	if v2037 == int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v2044 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[16])))
	if v2044 != 0 {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v2047 = F_errstart(m, v2045, int32(0))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L5
	} else {
		goto L394
	}
L391:
	;
	v2045 = int32(21)
	goto L393
L392:
	;
	v2045 = int32(23)
	goto L393
L393:
	;
	goto L390
L394:
	;
	if v2047 == int32(0) {
		goto L387
	} else {
		goto L395
	}
L395:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L5
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1912))) = v2010 + int32(28)
	F_errmsg(m, int32(_a_F_heapam_relation_copy_for_cluster_22), v1912)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L5
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_23), int32(925), int32(_a_F_heapam_relation_copy_for_cluster_24))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L5
	} else {
		goto L398
	}
L398:
	;
	goto L387
L399:
	;
	v2069 = F_hash_seq_search(m, v1912+int32(28))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L5
	} else {
		goto L400
	}
L400:
	;
	if v2069 != 0 {
		v2010 = v2069
		goto L385
	} else {
		goto L401
	}
L401:
	;
	goto L386
L402:
	;
	m.G0 = v1912 + int32(48)
	F_pfree(m, v47)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	F_pfree(m, v49)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L5
	} else {
		goto L404
	}
L404:
	;
	m.G0 = v28 - int32(-64)
	return
L405:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_25), int32(0))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_26), int32(1034), int32(_a_F_heapam_relation_copy_for_cluster_27))
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L5
	} else {
		goto L407
	}
L407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
