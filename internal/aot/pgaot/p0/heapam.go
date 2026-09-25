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
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
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
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v536 int64
	_ = v536
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int64
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v777 int32
	_ = v777
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int64
	_ = v835
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1244 float64
	_ = v1244
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 float64
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1309 float64
	_ = v1309
	var v1313 float64
	_ = v1313
	var v1317 float64
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1325 float64
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1389 float64
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
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
	var v1515 int64
	_ = v1515
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1665 int64
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1831 float64
	_ = v1831
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
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1918 float64
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1964 int32
	_ = v1964
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
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
	v763 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L5
	} else {
		goto L100
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = int64(8589934593)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = int64(2)
	v536 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+56)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v536
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
	v480 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v486 = m.T0[v485].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_heapam_relation_copy_for_cluster_12), v480, v480, v480, int32(449))
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
	v456 = int32(0)
	v458 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v459 = base.AtomicRmwOr32(m, v456, v458, v456)
	*(*int64)(unsafe.Add(mBase, uint32(v437+int32(8))+232)) = int64(1)
	v467 = base.AtomicRmwOr32(m, v456, v458, v456)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v468 + v449
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v474 - v449
	goto L73
L76:
	;
	v488 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v486)+36)))
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v491 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v746 = v432
	v751 = v486
	v756 = v11
	goto L36
L78:
	;
	goto L77
L79:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v495&int32(1) == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v500 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v503 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v502 + v503
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v506 + v503
	v510 = int32(0)
	v512 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v513 = base.AtomicRmwOr32(m, v510, v512, v510)
	*(*int64)(unsafe.Add(mBase, uint32(v491+int32(40))+232)) = v488
	v521 = base.AtomicRmwOr32(m, v510, v512, v510)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v522 + v503
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v528 - v503
	goto L78
L81:
	;
	v725 = int32(0)
	v728 = F_index_beginscan(m, l0, l2, int32(_a_F_heapam_relation_copy_for_cluster_12), v725, v725, v725)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L5
	} else {
		goto L98
	}
L82:
	;
	goto L81
L83:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v552 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v556&int32(1) == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v561 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v564 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v563 + v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v567 + v564
	v571 = int32(0)
	v574 = base.AtomicRmwOr32(m, v571, int32(_a_F_heapam_relation_copy_for_cluster_14), v571)
	goto L87
L86:
	;
	v701 = int32(0)
	v704 = base.AtomicRmwOr32(m, v701, int32(_a_F_heapam_relation_copy_for_cluster_14), v701)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v706 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v705 + v706
	v709 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v711 - v706
	goto L82
L87:
	;
	goto L89
L89:
	;
	goto L90
L90:
	;
	v666 = int32(0)
	v669 = int32(0)
	goto L95
L95:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8)+v669<<(uint(int32(2))%32))))
	v679 = int32(3)
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v26+int32(-32)+v669<<(uint(v679)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v552+int32(232)+v678<<(uint(v679)%32)))) = v685
	v687 = int32(1)
	v690 = v666 + v687
	if v690 != int32(2) {
		v666 = v690
		v669 = v669 + v687
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
	v730 = int32(0)
	F_index_rescan(m, v728, v730, v730, v730, v730)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v746 = v11
	v751 = v11
	v756 = v728
	goto L36
L100:
	;
	v777 = int32(-1)
	goto L104
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L5
	} else {
		goto L405
	}
L102:
	;
	if v763 != 0 {
		goto L317
	} else {
		goto L318
	}
L103:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1708)+188))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+12))
	m.T0[v1710].(func(*base.Module, int32))(m, v751)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L5
	} else {
		goto L316
	}
L104:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9]))
	if v791 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	F_index_endscan(m, v756)
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L5
	} else {
		goto L314
	}
L106:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v756 != 0 {
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
	v936 = int32(0)
	v938 = F_ExecFetchSlotHeapTuple(m, v763, v936, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L138
	}
L112:
	;
	v795 = F_index_getnext_slot(m, v756, int32(1), v763)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v763)+36)) = v816
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	if v819 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	if v795 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756)+72)))
	if v799 != int32(1) {
		v935 = v777
		goto L111
	} else {
		goto L117
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_16), int32(798), int32(_a_F_heapam_relation_copy_for_cluster_17))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
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
	v821 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[11])))
	if v821&int32(1) == int32(0) {
		goto L101
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+188))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+20))
	v830 = m.T0[v829].(func(*base.Module, int32, int32, int32) int32)(m, v751, int32(1), v763)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	if v830 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v835 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v751)+36)))
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v838 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v751)+52))
	if v777 == v879 {
		v935 = v777
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
	v842 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v842&int32(1) == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v847 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v850 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v849 + v850
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = v853 + v850
	v857 = int32(0)
	v859 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v860 = base.AtomicRmwOr32(m, v857, v859, v857)
	*(*int64)(unsafe.Add(mBase, uint32(v838+int32(48))+232)) = v835
	v868 = base.AtomicRmwOr32(m, v857, v859, v857)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = v869 + v850
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v875 - v850
	goto L130
L133:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v751)+36))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v751)+40))
	v886 = base.I32_rem_u_s(v879+v882-v884, v882)
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v892 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v751)+52))
	v935 = v933
	goto L111
L135:
	;
	goto L134
L136:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v896&int32(1) == int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v901 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v904 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v903 + v904
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	*(*int32)(unsafe.Add(mBase, uint32(v892))) = v907 + v904
	v911 = int32(0)
	v913 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v914 = base.AtomicRmwOr32(m, v911, v913, v911)
	*(*int64)(unsafe.Add(mBase, uint32(v892+int32(48))+232)) = base.I64_extend_i32_u(v886 + int32(1))
	v922 = base.AtomicRmwOr32(m, v911, v913, v911)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	*(*int32)(unsafe.Add(mBase, uint32(v892))) = v923 + v904
	v929 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v929 - v904
	goto L135
L138:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v763)+68))
	F_LockBuffer(m, v940, int32(1))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	v944 = F_HeapTupleSatisfiesVacuum(m, v938, l4, v940)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L5
	} else {
		goto L146
	}
L140:
	;
	F_LockBuffer(m, v940, int32(0))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L5
	} else {
		goto L262
	}
L141:
	;
	v1317 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1317, float64(1))
	goto L140
L142:
	;
	F_LockBuffer(m, v940, int32(0))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L5
	} else {
		goto L251
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
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
	switch v944 {
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
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+20)))
	v948 = int32(768)
	if v947&v948 != v948 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v954 = v952
	goto L150
L149:
	;
	v954 = int32(2)
	goto L150
L150:
	;
	if base.Ui32(v954) < base.Ui32(int32(3)) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v1074 != 0 {
		goto L140
	} else {
		goto L191
	}
L152:
	;
	v1074 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v965 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v965 == v954 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1074 = int32(1)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v969 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v969 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1074 = v1066
	goto L151
L159:
	;
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v973 == int32(0) {
		v1066 = int32(0)
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	v1037 = int32(0)
	v1039 = v969 - int32(1)
	goto L181
L162:
	;
	v978 = v973
	goto L163
L163:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v978)+20))
	if v983 == int32(4) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1066 = int32(0)
	goto L158
L165:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v978)+80))
	if v1030 != 0 {
		v978 = v1030
		goto L163
	} else {
		goto L180
	}
L166:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	if v986 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v989 = int32(1)
	if v954 == v986 {
		v1066 = v989
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v978)+52))
	v993 = v991 - int32(1)
	if v993 < int32(0) {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v998 = int32(0)
	v1000 = v993
	goto L170
L170:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v978)+48))
	v1006 = int32(2)
	v1007 = base.I32_div_s(v1000-v998, v1006)
	v1008 = v1007 + v998
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1004+v1008<<(uint(v1006)%32))))
	if v1012 == v954 {
		v1066 = v989
		goto L158
	} else {
		goto L172
	}
L171:
	;
	goto L165
L172:
	;
	v1016 = F_TransactionIdPrecedes(m, v1012, v954)
	mBase = m.M
	if v1016 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1017 = v1008 + int32(1)
	goto L175
L174:
	;
	v1017 = v998
	goto L175
L175:
	;
	if v1016 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1020 = v1000
	goto L178
L177:
	;
	v1020 = v1008 - int32(1)
	goto L178
L178:
	;
	if v1017 <= v1020 {
		v998 = v1017
		v1000 = v1020
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
	v1044 = int32(2)
	v1045 = base.I32_div_s(v1039-v1037, v1044)
	v1046 = v1045 + v1037
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v1046<<(uint(v1044)%32))))
	v1051 = base.B2i32(v1050 == v954)
	if v1050 == v954 {
		v1066 = v1051
		goto L158
	} else {
		goto L183
	}
L182:
	;
	v1066 = v1051
	goto L158
L183:
	;
	v1054 = base.B2i32(base.Ui32(v1050) < base.Ui32(v954))
	if base.Ui32(v1050) < base.Ui32(v954) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1055 = v1046 + int32(1)
	goto L186
L185:
	;
	v1055 = v1037
	goto L186
L186:
	;
	if base.Ui32(v1050) < base.Ui32(v954) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1058 = v1039
	goto L189
L188:
	;
	v1058 = v1046 - int32(1)
	goto L189
L189:
	;
	if v1055 <= v1058 {
		v1037 = v1055
		v1039 = v1058
		goto L181
	} else {
		goto L190
	}
L190:
	;
	goto L182
L191:
	;
	v1077 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v1077 == int32(0) {
		goto L140
	} else {
		goto L193
	}
L193:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v1081 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_18), v28)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_16), int32(868), int32(_a_F_heapam_relation_copy_for_cluster_17))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	goto L140
L196:
	;
	v1244 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1244, float64(1))
	goto L140
L197:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v1094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093)+20)))
	if v1094&int32(_a_F_heapam_relation_copy_for_cluster_19) == int32(_a_F_heapam_relation_copy_for_cluster_20) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if base.Ui32(v1102) < base.Ui32(int32(3)) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v1099 = F_HeapTupleGetUpdateXid(m, v1093)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L5
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1102 = v1101
	goto L198
L202:
	;
	v1102 = v1099
	goto L198
L203:
	;
	if v1222 != 0 {
		goto L196
	} else {
		goto L243
	}
L204:
	;
	v1222 = int32(0)
	goto L203
L205:
	;
	goto L206
L206:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v1113 == v1102 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1222 = int32(1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v1117 <= int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v1222 = v1214
	goto L203
L211:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v1121 == int32(0) {
		v1214 = int32(0)
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	v1185 = int32(0)
	v1187 = v1117 - int32(1)
	goto L233
L214:
	;
	v1126 = v1121
	goto L215
L215:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+20))
	if v1131 == int32(4) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1214 = int32(0)
	goto L210
L217:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+80))
	if v1178 != 0 {
		v1126 = v1178
		goto L215
	} else {
		goto L232
	}
L218:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	if v1134 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1137 = int32(1)
	if v1102 == v1134 {
		v1214 = v1137
		goto L210
	} else {
		goto L220
	}
L220:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+52))
	v1141 = v1139 - int32(1)
	if v1141 < int32(0) {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	v1146 = int32(0)
	v1148 = v1141
	goto L222
L222:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+48))
	v1154 = int32(2)
	v1155 = base.I32_div_s(v1148-v1146, v1154)
	v1156 = v1155 + v1146
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1152+v1156<<(uint(v1154)%32))))
	if v1160 == v1102 {
		v1214 = v1137
		goto L210
	} else {
		goto L224
	}
L223:
	;
	goto L217
L224:
	;
	v1164 = F_TransactionIdPrecedes(m, v1160, v1102)
	mBase = m.M
	if v1164 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1165 = v1156 + int32(1)
	goto L227
L226:
	;
	v1165 = v1146
	goto L227
L227:
	;
	if v1164 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1168 = v1148
	goto L230
L229:
	;
	v1168 = v1156 - int32(1)
	goto L230
L230:
	;
	if v1165 <= v1168 {
		v1146 = v1165
		v1148 = v1168
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
	v1192 = int32(2)
	v1193 = base.I32_div_s(v1187-v1185, v1192)
	v1194 = v1193 + v1185
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1183+v1194<<(uint(v1192)%32))))
	v1199 = base.B2i32(v1198 == v1102)
	if v1198 == v1102 {
		v1214 = v1199
		goto L210
	} else {
		goto L235
	}
L234:
	;
	v1214 = v1199
	goto L210
L235:
	;
	v1202 = base.B2i32(base.Ui32(v1198) < base.Ui32(v1102))
	if base.Ui32(v1198) < base.Ui32(v1102) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1203 = v1194 + int32(1)
	goto L238
L237:
	;
	v1203 = v1185
	goto L238
L238:
	;
	if base.Ui32(v1198) < base.Ui32(v1102) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1206 = v1187
	goto L241
L240:
	;
	v1206 = v1194 - int32(1)
	goto L241
L241:
	;
	if v1203 <= v1206 {
		v1185 = v1203
		v1187 = v1206
		goto L233
	} else {
		goto L242
	}
L242:
	;
	goto L234
L243:
	;
	v1225 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L5
	} else {
		goto L244
	}
L244:
	;
	if v1225 == int32(0) {
		goto L196
	} else {
		goto L245
	}
L245:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v1229 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_21), v26+int32(-48))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_16), int32(880), int32(_a_F_heapam_relation_copy_for_cluster_17))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L196
L248:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_22), int32(0))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_16), int32(886), int32(_a_F_heapam_relation_copy_for_cluster_17))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
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
	v1264 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1264, float64(1))
	v1268 = m.G0
	v1270 = v1268 - int32(16)
	m.G0 = v1270
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+12)) = int32(0)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v1275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1274)+20)))
	v1276 = int32(768)
	if v1275&v1276 != v1276 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1274)))
	v1282 = v1280
	goto L254
L253:
	;
	v1282 = int32(2)
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+4)) = v1282
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v938)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+8)) = v1284
	v1286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v938)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1270)+12)) = uint16(v1286)
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v1290 = v1270 + int32(4)
	v1291 = int32(0)
	v1293 = F_hash_search(m, v1288, v1290, v1291, v1291)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	if v1293 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+20))
	F_pfree(m, v1295)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	m.G0 = v1270 + int32(16)
	if v1293 == int32(0) {
		v777 = v935
		goto L104
	} else {
		goto L261
	}
L259:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v1302 = F_hash_search(m, v1298, v1290, int32(2), v1270+int32(3))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1309 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1309, float64(1))
	v1313 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1313, float64(-1))
	v777 = v935
	goto L104
L262:
	;
	v1325 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_add(v1325, float64(1))
	if v746 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1329 = m.G0
	v1331 = v1329 - int32(16)
	m.G0 = v1331
	v1333 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v1334 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v746)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1336
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v746)+60))
	v1339 = F_heap_copytuple(m, v938)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
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
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v938, v1437, v47, v49)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L5
	} else {
		goto L284
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1331))) = v1339
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+36)))
	if v1342 == int32(1) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	v1346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1345)+12)))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1350 = F_heap_getattr_1(m, v1339, v1346, v1347, v1331+int32(8))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L5
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+52)))
	if v1353&int32(2) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+4)) = v1350
	goto L269
L271:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+36)))
	if v1367 != int32(1) {
		v1378 = int32(0)
		goto L276
	} else {
		goto L277
	}
L272:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1339)))
	v1365 = (v1358 + int32(31)) & int32(-8)
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1363 = F_GetMemoryChunkSpace(m, v1339)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	v1365 = v1363
	goto L271
L276:
	;
	F_tuplesort_puttuple_common(m, v746, v1331, v1378&int32(1), v1365)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L5
	} else {
		goto L279
	}
L277:
	;
	v1370 = int32(0)
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v746)+44))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+24))
	if v1372 == v1370 {
		v1378 = v1370
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331)+8)))
	v1378 = v1375 ^ int32(1)
	goto L276
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1334
	m.G0 = v1331 + int32(16)
	v1389 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	v1393 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1393 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v777 = v935
	goto L104
L281:
	;
	goto L280
L282:
	;
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1397&int32(1) == int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1402 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1404 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1405 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1404 + v1405
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1408 + v1405
	v1412 = int32(0)
	v1414 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v1415 = base.AtomicRmwOr32(m, v1412, v1414, v1412)
	*(*int64)(unsafe.Add(mBase, uint32(v1393+int32(24))+232)) = base.I64_trunc_sat_f64_s(v1389)
	v1423 = base.AtomicRmwOr32(m, v1412, v1414, v1412)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1424 + v1405
	v1430 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1430 - v1405
	goto L281
L284:
	;
	v1440 = int32(0)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1436)))
	if v1440 < v1441 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1446 = v1440
	v1449 = v1441
	goto L288
L286:
	;
	goto L287
L287:
	;
	v1508 = F_heap_form_tuple(m, v1436, v47, v49)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L5
	} else {
		goto L294
	}
L288:
	;
	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1446<<(uint(int32(4))%32))+29)))
	if v1472 == int32(1) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	goto L287
L290:
	;
	v1476 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1446+v49))) = uint8(v1476)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1436)))
	v1479 = v1478
	goto L292
L291:
	;
	v1479 = v1449
	goto L292
L292:
	;
	v1481 = v1446 + int32(1)
	if v1481 < v1479 {
		v1446 = v1481
		v1449 = v1479
		goto L288
	} else {
		goto L293
	}
L293:
	;
	goto L289
L294:
	;
	F_rewrite_heap_tuple(m, v70, v938, v1508)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	F_pfree(m, v1508)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v1514 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
	v1515 = base.I64_trunc_sat_f64_s(v1514)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v1515
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v1515
	goto L299
L297:
	;
	v777 = v935
	goto L104
L298:
	;
	goto L297
L299:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1532 == int32(0) {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1536&int32(1) == int32(0) {
		goto L298
	} else {
		goto L301
	}
L301:
	;
	v1541 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1543 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1544 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1543 + v1544
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1532)))
	*(*int32)(unsafe.Add(mBase, uint32(v1532))) = v1547 + v1544
	v1551 = int32(0)
	v1554 = base.AtomicRmwOr32(m, v1551, int32(_a_F_heapam_relation_copy_for_cluster_14), v1551)
	goto L303
L302:
	;
	v1681 = int32(0)
	v1684 = base.AtomicRmwOr32(m, v1681, int32(_a_F_heapam_relation_copy_for_cluster_14), v1681)
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1532)))
	v1686 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1532))) = v1685 + v1686
	v1689 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1691 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1691 - v1686
	goto L298
L303:
	;
	goto L305
L305:
	;
	goto L306
L306:
	;
	v1646 = int32(0)
	v1649 = int32(0)
	goto L311
L311:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8)+v1649<<(uint(int32(2))%32))))
	v1659 = int32(3)
	v1665 = *(*int64)(unsafe.Add(mBase, uint32(v26+int32(-32)+v1649<<(uint(v1659)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1532+int32(232)+v1658<<(uint(v1659)%32)))) = v1665
	v1667 = int32(1)
	v1670 = v1646 + v1667
	if v1670 != int32(2) {
		v1646 = v1670
		v1649 = v1649 + v1667
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
	if v751 == int32(0) {
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
	F_ExecDropSingleTupleTableSlot(m, v763)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L5
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	if v746 != 0 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	goto L319
L321:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1719 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L322:
	;
	goto L323
L323:
	;
	v1990 = m.G0
	v1992 = v1990 - int32(48)
	m.G0 = v1992
	v1995 = v1992 + int32(8)
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	F_hash_seq_init(m, v1995, v1996)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L5
	} else {
		goto L361
	}
L324:
	;
	F_tuplesort_performsort(m, v746)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L5
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1723&int32(1) == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1728 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1730 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1731 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1730 + v1731
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1719)))
	*(*int32)(unsafe.Add(mBase, uint32(v1719))) = v1734 + v1731
	v1738 = int32(0)
	v1740 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v1741 = base.AtomicRmwOr32(m, v1738, v1740, v1738)
	*(*int64)(unsafe.Add(mBase, uint32(v1719+int32(8))+232)) = int64(3)
	v1749 = base.AtomicRmwOr32(m, v1738, v1740, v1738)
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1719)))
	*(*int32)(unsafe.Add(mBase, uint32(v1719))) = v1750 + v1731
	v1756 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1756 - v1731
	goto L325
L328:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1766 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1831 = float64(0)
	goto L333
L330:
	;
	goto L329
L331:
	;
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1770&int32(1) == int32(0) {
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1775 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1777 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1778 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1777 + v1778
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1766)))
	*(*int32)(unsafe.Add(mBase, uint32(v1766))) = v1781 + v1778
	v1785 = int32(0)
	v1787 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v1788 = base.AtomicRmwOr32(m, v1785, v1787, v1785)
	*(*int64)(unsafe.Add(mBase, uint32(v1766+int32(8))+232)) = int64(4)
	v1796 = base.AtomicRmwOr32(m, v1785, v1787, v1785)
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1766)))
	*(*int32)(unsafe.Add(mBase, uint32(v1766))) = v1797 + v1778
	v1803 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1803 - v1778
	goto L330
L333:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9]))
	if v1833 != 0 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	F_tuplesort_end(m, v746)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L5
	} else {
		goto L360
	}
L335:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L5
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1836 = F_tuplesort_getheaptuple(m, v746)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L5
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	if v1836 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, v1836, v1839, v47, v49)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
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
	v1842 = int32(0)
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1838)))
	if v1842 < v1843 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1848 = v1842
	v1851 = v1843
	goto L347
L345:
	;
	goto L346
L346:
	;
	v1910 = F_heap_form_tuple(m, v1838, v47, v49)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L5
	} else {
		goto L353
	}
L347:
	;
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1838+v1848<<(uint(int32(4))%32))+29)))
	if v1874 == int32(1) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L346
L349:
	;
	v1878 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1848+v49))) = uint8(v1878)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1838)))
	v1881 = v1880
	goto L351
L350:
	;
	v1881 = v1851
	goto L351
L351:
	;
	v1883 = v1848 + int32(1)
	if v1883 < v1881 {
		v1848 = v1883
		v1851 = v1881
		goto L347
	} else {
		goto L352
	}
L352:
	;
	goto L348
L353:
	;
	F_rewrite_heap_tuple(m, v70, v1836, v1910)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	F_pfree(m, v1910)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	v1918 = base.F64_add(v1831, float64(1))
	v1922 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	if v1922 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1831 = v1918
	goto L333
L357:
	;
	goto L356
L358:
	;
	v1926 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v1926&int32(1) == int32(0) {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1931 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1933 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	v1934 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1933 + v1934
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1922)))
	*(*int32)(unsafe.Add(mBase, uint32(v1922))) = v1937 + v1934
	v1941 = int32(0)
	v1943 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	v1944 = base.AtomicRmwOr32(m, v1941, v1943, v1941)
	*(*int64)(unsafe.Add(mBase, uint32(v1922+int32(32))+232)) = base.I64_trunc_sat_f64_s(v1918)
	v1952 = base.AtomicRmwOr32(m, v1941, v1943, v1941)
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1922)))
	*(*int32)(unsafe.Add(mBase, uint32(v1922))) = v1953 + v1934
	v1959 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8])) = v1959 - v1934
	goto L357
L360:
	;
	goto L323
L361:
	;
	v1999 = F_hash_seq_search(m, v1995)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	if v1999 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v2001 = v1999
	goto L366
L364:
	;
	goto L365
L365:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v2064 != 0 {
		goto L371
	} else {
		goto L372
	}
L366:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v2001)+20))
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+16))
	v2028 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2027)+16)) = uint16(v2028)
	*(*int32)(unsafe.Add(mBase, uint32(v2027)+12)) = int32(-1)
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2001)+20))
	F_raw_heap_insert(m, v70, v2032)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L5
	} else {
		goto L368
	}
L367:
	;
	goto L365
L368:
	;
	v2037 = F_hash_seq_search(m, v1992+int32(8))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L5
	} else {
		goto L369
	}
L369:
	;
	if v2037 != 0 {
		v2001 = v2037
		goto L366
	} else {
		goto L370
	}
L370:
	;
	goto L367
L371:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	F_smgr_bulk_write(m, v2065, v2066, v2064, int32(1))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L5
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	F_smgr_bulk_finish(m, v2072)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
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
	v2075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)))
	if v2075 != int32(1) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	F_MemoryContextDelete(m, v2176)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L5
	} else {
		goto L402
	}
L377:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if v2078 != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	F_logical_heap_rewrite_flush_mappings(m, v70)
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L5
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v2082 = v1992 + int32(28)
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v70)+64))
	F_hash_seq_init(m, v2082, v2083)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L5
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	v2086 = F_hash_seq_search(m, v2082)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	if v2086 == int32(0) {
		goto L376
	} else {
		goto L384
	}
L384:
	;
	v2090 = v2086
	goto L385
L385:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2090)+4))
	v2117 = F_FileSync(m, v2115, int32(167772197))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L5
	} else {
		goto L388
	}
L386:
	;
	goto L376
L387:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v2090)+4))
	F_FileClose(m, v2144)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L5
	} else {
		goto L399
	}
L388:
	;
	if v2117 == int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[16])))
	if v2124 != 0 {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v2127 = F_errstart(m, v2125, int32(0))
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L5
	} else {
		goto L394
	}
L391:
	;
	v2125 = int32(21)
	goto L393
L392:
	;
	v2125 = int32(23)
	goto L393
L393:
	;
	goto L390
L394:
	;
	if v2127 == int32(0) {
		goto L387
	} else {
		goto L395
	}
L395:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L5
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1992))) = v2090 + int32(28)
	F_errmsg(m, int32(_a_F_heapam_relation_copy_for_cluster_23), v1992)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L5
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_24), int32(925), int32(_a_F_heapam_relation_copy_for_cluster_25))
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L5
	} else {
		goto L398
	}
L398:
	;
	goto L387
L399:
	;
	v2149 = F_hash_seq_search(m, v1992+int32(28))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L5
	} else {
		goto L400
	}
L400:
	;
	if v2149 != 0 {
		v2090 = v2149
		goto L385
	} else {
		goto L401
	}
L401:
	;
	goto L386
L402:
	;
	m.G0 = v1992 + int32(48)
	F_pfree(m, v47)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	F_pfree(m, v49)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_26), int32(0))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_27), int32(1034), int32(_a_F_heapam_relation_copy_for_cluster_28))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
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
