package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_StartupXLOG(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v672 int32
	_ = v672
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
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
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v902 int64
	_ = v902
	var v905 int64
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int64
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1098 int64
	_ = v1098
	var v1099 int64
	_ = v1099
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1121 int64
	_ = v1121
	var v1122 int64
	_ = v1122
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int64
	_ = v1137
	var v1138 int64
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int64
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1285 int64
	_ = v1285
	var v1288 int64
	_ = v1288
	var v1290 int64
	_ = v1290
	var v1291 int64
	_ = v1291
	var v1294 int64
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1311 int64
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int64
	_ = v1321
	var v1322 int64
	_ = v1322
	var v1324 int64
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int64
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int64
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int64
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int64
	_ = v1346
	var v1349 int64
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1366 int64
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1383 int64
	_ = v1383
	var v1386 int64
	_ = v1386
	var v1387 int64
	_ = v1387
	var v1390 int64
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1418 int64
	_ = v1418
	var v1421 int64
	_ = v1421
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1500 int32
	_ = v1500
	var v1519 int32
	_ = v1519
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1572 int64
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1605 int32
	_ = v1605
	var v1632 int32
	_ = v1632
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1796 int32
	_ = v1796
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1845 int32
	_ = v1845
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1909 int32
	_ = v1909
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1927 int32
	_ = v1927
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2116 int32
	_ = v2116
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2140 int64
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2144 int64
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int64
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2171 int64
	_ = v2171
	var v2174 int64
	_ = v2174
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2188 int64
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2193 int64
	_ = v2193
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2207 int64
	_ = v2207
	var v2210 int64
	_ = v2210
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int64
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int64
	_ = v2231
	var v2233 int64
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2236 int64
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int64
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int64
	_ = v2247
	var v2249 int64
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2266 int32
	_ = v2266
	var v2268 int64
	_ = v2268
	var v2271 int64
	_ = v2271
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2288 int64
	_ = v2288
	var v2291 int64
	_ = v2291
	var v2292 int64
	_ = v2292
	var v2295 int64
	_ = v2295
	var v2301 int32
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2323 int32
	_ = v2323
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int64
	_ = v2344
	var v2347 int64
	_ = v2347
	var v2348 int64
	_ = v2348
	var v2349 int64
	_ = v2349
	var v2350 int64
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2379 int64
	_ = v2379
	var v2382 int64
	_ = v2382
	var v2383 int64
	_ = v2383
	var v2384 int64
	_ = v2384
	var v2386 int64
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2416 int32
	_ = v2416
	var v2422 int32
	_ = v2422
	var v2427 int64
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2435 int32
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2446 int32
	_ = v2446
	var v2451 int64
	_ = v2451
	var v2454 int64
	_ = v2454
	var v2460 int32
	_ = v2460
	var v2467 int32
	_ = v2467
	var v2474 int32
	_ = v2474
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2487 int64
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2495 int64
	_ = v2495
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2512 int64
	_ = v2512
	var v2518 int32
	_ = v2518
	var v2524 int32
	_ = v2524
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2556 int32
	_ = v2556
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2572 int32
	_ = v2572
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2588 int32
	_ = v2588
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2604 int32
	_ = v2604
	var v2609 int32
	_ = v2609
	var v2614 int64
	_ = v2614
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2642 int32
	_ = v2642
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2663 int32
	_ = v2663
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2689 int32
	_ = v2689
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2698 int64
	_ = v2698
	var v2702 int64
	_ = v2702
	var v2704 int64
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2721 int32
	_ = v2721
	var v2724 int64
	_ = v2724
	var v2732 int32
	_ = v2732
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2755 int64
	_ = v2755
	var v2759 int64
	_ = v2759
	var v2762 int32
	_ = v2762
	var v2765 int64
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2773 int64
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2779 int64
	_ = v2779
	var v2787 int32
	_ = v2787
	var v2788 int64
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2792 int64
	_ = v2792
	var v2798 int64
	_ = v2798
	var v2814 int32
	_ = v2814
	var v2816 int64
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2829 int32
	_ = v2829
	var v2831 int64
	_ = v2831
	var v2836 int32
	_ = v2836
	var v2839 int64
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2845 int64
	_ = v2845
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2856 int32
	_ = v2856
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int64
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2869 int64
	_ = v2869
	var v2875 int32
	_ = v2875
	var v2880 int32
	_ = v2880
	var v2884 int32
	_ = v2884
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2906 int32
	_ = v2906
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2912 int64
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int64
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int64
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2929 int32
	_ = v2929
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2949 int32
	_ = v2949
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3014 int64
	_ = v3014
	var v3022 int32
	_ = v3022
	var v3026 int32
	_ = v3026
	var v3030 int32
	_ = v3030
	var v3036 int32
	_ = v3036
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3054 int32
	_ = v3054
	var v3057 int32
	_ = v3057
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3070 int32
	_ = v3070
	var v3076 int32
	_ = v3076
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3099 int32
	_ = v3099
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3105 int32
	_ = v3105
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3148 int32
	_ = v3148
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3163 int32
	_ = v3163
	var v3168 int32
	_ = v3168
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3192 int32
	_ = v3192
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3221 int32
	_ = v3221
	var v3224 int32
	_ = v3224
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3267 int32
	_ = v3267
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3296 int32
	_ = v3296
	var v3301 int32
	_ = v3301
	var v3305 int32
	_ = v3305
	var v3310 int32
	_ = v3310
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3328 int32
	_ = v3328
	var v3335 int32
	_ = v3335
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3350 int32
	_ = v3350
	var v3355 int32
	_ = v3355
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3364 int32
	_ = v3364
	var v3369 int32
	_ = v3369
	var v3374 int32
	_ = v3374
	var v3378 int32
	_ = v3378
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3388 int32
	_ = v3388
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3414 int32
	_ = v3414
	var v3422 int32
	_ = v3422
	var v3427 int32
	_ = v3427
	var v3436 int32
	_ = v3436
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3445 int32
	_ = v3445
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3456 int32
	_ = v3456
	var v3458 int32
	_ = v3458
	var v3461 int32
	_ = v3461
	var v3466 int32
	_ = v3466
	var v3471 int32
	_ = v3471
	var v3481 int32
	_ = v3481
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3512 int32
	_ = v3512
	var v3517 int32
	_ = v3517
	var v3521 int32
	_ = v3521
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3528 int32
	_ = v3528
	var v3532 int32
	_ = v3532
	var v3538 int32
	_ = v3538
	var v3541 int32
	_ = v3541
	var v3547 int32
	_ = v3547
	var v3551 int32
	_ = v3551
	var v3556 int32
	_ = v3556
	var v3560 int32
	_ = v3560
	var v3565 int32
	_ = v3565
	var v3579 int32
	_ = v3579
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3607 int32
	_ = v3607
	var v3613 int32
	_ = v3613
	var v3615 int32
	_ = v3615
	var v3617 int64
	_ = v3617
	var v3619 int64
	_ = v3619
	var v3620 int64
	_ = v3620
	var v3631 int32
	_ = v3631
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3641 int64
	_ = v3641
	var v3642 int64
	_ = v3642
	var v3651 int32
	_ = v3651
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3729 int32
	_ = v3729
	var v3731 int32
	_ = v3731
	var v3733 int32
	_ = v3733
	var v3739 int32
	_ = v3739
	var v3741 int32
	_ = v3741
	var v3748 int32
	_ = v3748
	var v3750 int32
	_ = v3750
	var v3758 int32
	_ = v3758
	var v3763 int32
	_ = v3763
	var v3767 int32
	_ = v3767
	var v3769 int32
	_ = v3769
	var v3777 int32
	_ = v3777
	var v3782 int32
	_ = v3782
	var v3786 int32
	_ = v3786
	var v3788 int32
	_ = v3788
	var v3796 int32
	_ = v3796
	var v3801 int32
	_ = v3801
	var v3803 int32
	_ = v3803
	var v3811 int32
	_ = v3811
	var v3816 int32
	_ = v3816
	var v3820 int32
	_ = v3820
	var v3823 int32
	_ = v3823
	var v3834 int32
	_ = v3834
	var v3839 int32
	_ = v3839
	var v3843 int32
	_ = v3843
	var v3846 int32
	_ = v3846
	var v3855 int32
	_ = v3855
	var v3860 int32
	_ = v3860
	var v3864 int32
	_ = v3864
	var v3867 int32
	_ = v3867
	var v3876 int32
	_ = v3876
	var v3881 int32
	_ = v3881
	var v3883 int32
	_ = v3883
	var v3891 int32
	_ = v3891
	var v3896 int32
	_ = v3896
	var v3900 int32
	_ = v3900
	var v3902 int32
	_ = v3902
	var v3910 int32
	_ = v3910
	var v3915 int32
	_ = v3915
	var v3919 int32
	_ = v3919
	var v3921 int32
	_ = v3921
	var v3930 int32
	_ = v3930
	var v3935 int32
	_ = v3935
	var v3939 int32
	_ = v3939
	var v3942 int32
	_ = v3942
	var v3948 int32
	_ = v3948
	var v3952 int32
	_ = v3952
	var v3957 int32
	_ = v3957
	var v3961 int32
	_ = v3961
	var v3964 int32
	_ = v3964
	var v3970 int32
	_ = v3970
	var v3974 int32
	_ = v3974
	var v3979 int32
	_ = v3979
	var v4019 int32
	_ = v4019
	var v4023 int32
	_ = v4023
	var v4027 int32
	_ = v4027
	var v4032 int32
	_ = v4032
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4038 int32
	_ = v4038
	var v4041 int32
	_ = v4041
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4058 int32
	_ = v4058
	var v4086 int32
	_ = v4086
	var v4089 int32
	_ = v4089
	var v4092 int32
	_ = v4092
	var v4095 int32
	_ = v4095
	var v4099 int32
	_ = v4099
	var v4102 int32
	_ = v4102
	var v4103 int32
	_ = v4103
	var v4107 int32
	_ = v4107
	var v4110 int32
	_ = v4110
	var v4111 int32
	_ = v4111
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4153 int32
	_ = v4153
	var v4154 int64
	_ = v4154
	var v4159 int32
	_ = v4159
	var v4160 int64
	_ = v4160
	var v4162 int32
	_ = v4162
	var v4163 int64
	_ = v4163
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4171 int64
	_ = v4171
	var v4173 int32
	_ = v4173
	var v4175 int64
	_ = v4175
	var v4177 int64
	_ = v4177
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4187 int32
	_ = v4187
	var v4192 int32
	_ = v4192
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4202 int32
	_ = v4202
	var v4207 int32
	_ = v4207
	var v4210 int32
	_ = v4210
	var v4211 int32
	_ = v4211
	var v4215 int32
	_ = v4215
	var v4221 int32
	_ = v4221
	var v4223 int32
	_ = v4223
	var v4228 int32
	_ = v4228
	var v4233 int32
	_ = v4233
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4243 int32
	_ = v4243
	var v4248 int32
	_ = v4248
	var v4258 int32
	_ = v4258
	var v4263 int32
	_ = v4263
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4275 int32
	_ = v4275
	var v4279 int32
	_ = v4279
	var v4292 int32
	_ = v4292
	var v4293 int32
	_ = v4293
	var v4322 int32
	_ = v4322
	var v4324 int32
	_ = v4324
	var v4334 int32
	_ = v4334
	var v4339 int32
	_ = v4339
	var v4343 int32
	_ = v4343
	var v4345 int32
	_ = v4345
	var v4348 int32
	_ = v4348
	var v4351 int32
	_ = v4351
	var v4352 int32
	_ = v4352
	var v4354 int64
	_ = v4354
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4360 int32
	_ = v4360
	var v4362 int64
	_ = v4362
	var v4365 int64
	_ = v4365
	var v4371 int32
	_ = v4371
	var v4376 int32
	_ = v4376
	var v4383 int32
	_ = v4383
	var v4425 int32
	_ = v4425
	var v4427 int32
	_ = v4427
	var v4434 int32
	_ = v4434
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4444 int32
	_ = v4444
	var v4445 int32
	_ = v4445
	var v4486 int32
	_ = v4486
	var v4493 int32
	_ = v4493
	var v4498 int32
	_ = v4498
	var v4502 int32
	_ = v4502
	var v4505 int32
	_ = v4505
	var v4509 int32
	_ = v4509
	var v4514 int32
	_ = v4514
	var v4518 int32
	_ = v4518
	var v4521 int32
	_ = v4521
	var v4528 int32
	_ = v4528
	var v4533 int32
	_ = v4533
	var v4537 int32
	_ = v4537
	var v4539 int32
	_ = v4539
	var v4546 int32
	_ = v4546
	var v4551 int32
	_ = v4551
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4564 int32
	_ = v4564
	var v4569 int32
	_ = v4569
	var v4571 int32
	_ = v4571
	var v4573 int32
	_ = v4573
	var v4574 int32
	_ = v4574
	var v4577 int64
	_ = v4577
	var v4579 int64
	_ = v4579
	var v4581 int64
	_ = v4581
	var v4583 int32
	_ = v4583
	var v4585 int32
	_ = v4585
	var v4587 int32
	_ = v4587
	var v4591 int32
	_ = v4591
	var v4592 int32
	_ = v4592
	var v4594 int32
	_ = v4594
	var v4595 int32
	_ = v4595
	var v4597 int32
	_ = v4597
	var v4598 int32
	_ = v4598
	var v4607 int32
	_ = v4607
	var v4636 int32
	_ = v4636
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4644 int32
	_ = v4644
	var v4646 int32
	_ = v4646
	var v4647 int64
	_ = v4647
	var v4655 int32
	_ = v4655
	var v4659 int32
	_ = v4659
	var v4663 int32
	_ = v4663
	var v4669 int32
	_ = v4669
	var v4673 int32
	_ = v4673
	var v4674 int32
	_ = v4674
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4687 int32
	_ = v4687
	var v4690 int32
	_ = v4690
	var v4694 int32
	_ = v4694
	var v4695 int32
	_ = v4695
	var v4703 int32
	_ = v4703
	var v4709 int32
	_ = v4709
	var v4711 int32
	_ = v4711
	var v4713 int32
	_ = v4713
	var v4723 int32
	_ = v4723
	var v4729 int64
	_ = v4729
	var v4732 int32
	_ = v4732
	var v4734 int32
	_ = v4734
	var v4735 int32
	_ = v4735
	var v4738 int64
	_ = v4738
	var v4742 int32
	_ = v4742
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4785 int32
	_ = v4785
	var v4789 int32
	_ = v4789
	var v4791 int32
	_ = v4791
	var v4794 int32
	_ = v4794
	var v4796 int32
	_ = v4796
	var v4800 int32
	_ = v4800
	var v4802 int32
	_ = v4802
	var v4807 int32
	_ = v4807
	var v4808 int32
	_ = v4808
	var v4817 int32
	_ = v4817
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4826 int32
	_ = v4826
	var v4833 int32
	_ = v4833
	var v4837 int32
	_ = v4837
	var v4838 int32
	_ = v4838
	var v4842 int32
	_ = v4842
	var v4847 int32
	_ = v4847
	var v4849 int32
	_ = v4849
	var v4852 int32
	_ = v4852
	var v4862 int32
	_ = v4862
	var v4901 int32
	_ = v4901
	var v4909 int32
	_ = v4909
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4918 int32
	_ = v4918
	var v4921 int32
	_ = v4921
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4929 int32
	_ = v4929
	var v4932 int64
	_ = v4932
	var v4933 int64
	_ = v4933
	var v4951 int32
	_ = v4951
	var v4990 int32
	_ = v4990
	var v4998 int32
	_ = v4998
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5007 int32
	_ = v5007
	var v5009 int32
	_ = v5009
	var v5012 int32
	_ = v5012
	var v5016 int32
	_ = v5016
	var v5020 int32
	_ = v5020
	var v5022 int32
	_ = v5022
	var v5025 int32
	_ = v5025
	var v5028 int32
	_ = v5028
	var v5029 int32
	_ = v5029
	var v5036 int32
	_ = v5036
	var v5041 int32
	_ = v5041
	var v5044 int32
	_ = v5044
	var v5045 int32
	_ = v5045
	var v5049 int32
	_ = v5049
	var v5054 int32
	_ = v5054
	var v5055 int32
	_ = v5055
	var v5059 int32
	_ = v5059
	var v5064 int32
	_ = v5064
	var v5069 int32
	_ = v5069
	var v5073 int32
	_ = v5073
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5078 int64
	_ = v5078
	var v5079 int64
	_ = v5079
	var v5089 int32
	_ = v5089
	var v5136 int32
	_ = v5136
	var v5144 int32
	_ = v5144
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5154 int32
	_ = v5154
	var v5156 int32
	_ = v5156
	var v5160 int32
	_ = v5160
	var v5164 int32
	_ = v5164
	var v5170 int32
	_ = v5170
	var v5171 int32
	_ = v5171
	var v5176 int32
	_ = v5176
	var v5177 int32
	_ = v5177
	var v5183 int32
	_ = v5183
	var v5188 int32
	_ = v5188
	var v5189 int32
	_ = v5189
	var v5232 int32
	_ = v5232
	var v5233 int32
	_ = v5233
	var v5240 int32
	_ = v5240
	var v5241 int32
	_ = v5241
	var v5246 int32
	_ = v5246
	var v5247 int32
	_ = v5247
	var v5256 int32
	_ = v5256
	var v5258 int32
	_ = v5258
	var v5260 int32
	_ = v5260
	var v5264 int32
	_ = v5264
	var v5272 int32
	_ = v5272
	var v5273 int32
	_ = v5273
	var v5283 int32
	_ = v5283
	var v5290 int32
	_ = v5290
	var v5294 int32
	_ = v5294
	var v5298 int32
	_ = v5298
	var v5303 int32
	_ = v5303
	var v5304 int32
	_ = v5304
	var v5314 int32
	_ = v5314
	var v5316 int32
	_ = v5316
	var v5320 int32
	_ = v5320
	var v5321 int32
	_ = v5321
	var v5328 int32
	_ = v5328
	var v5331 int32
	_ = v5331
	var v5332 int32
	_ = v5332
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5338 int32
	_ = v5338
	var v5339 int32
	_ = v5339
	var v5340 int32
	_ = v5340
	var v5344 int32
	_ = v5344
	var v5345 int32
	_ = v5345
	var v5348 int32
	_ = v5348
	var v5357 int32
	_ = v5357
	var v5360 int32
	_ = v5360
	var v5362 int32
	_ = v5362
	var v5369 int32
	_ = v5369
	var v5370 int32
	_ = v5370
	var v5375 int32
	_ = v5375
	var v5376 int32
	_ = v5376
	var v5385 int32
	_ = v5385
	var v5387 int32
	_ = v5387
	var v5389 int32
	_ = v5389
	var v5391 int32
	_ = v5391
	var v5401 int32
	_ = v5401
	var v5402 int32
	_ = v5402
	var v5407 int64
	_ = v5407
	var v5409 int64
	_ = v5409
	var v5415 int32
	_ = v5415
	var v5426 int32
	_ = v5426
	var v5434 int32
	_ = v5434
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5445 int64
	_ = v5445
	var v5447 int64
	_ = v5447
	var v5453 int32
	_ = v5453
	var v5459 int32
	_ = v5459
	var v5460 int32
	_ = v5460
	var v5465 int32
	_ = v5465
	var v5466 int32
	_ = v5466
	var v5474 int32
	_ = v5474
	var v5480 int32
	_ = v5480
	var v5481 int32
	_ = v5481
	var v5486 int32
	_ = v5486
	var v5487 int32
	_ = v5487
	var v5490 int32
	_ = v5490
	var v5497 int32
	_ = v5497
	var v5499 int32
	_ = v5499
	var v5501 int32
	_ = v5501
	var v5503 int32
	_ = v5503
	var v5513 int32
	_ = v5513
	var v5514 int32
	_ = v5514
	var v5523 int32
	_ = v5523
	var v5526 int32
	_ = v5526
	var v5536 int32
	_ = v5536
	var v5544 int32
	_ = v5544
	var v5549 int32
	_ = v5549
	var v5550 int32
	_ = v5550
	var v5559 int32
	_ = v5559
	var v5562 int32
	_ = v5562
	var v5563 int32
	_ = v5563
	var v5568 int32
	_ = v5568
	var v5569 int32
	_ = v5569
	var v5575 int32
	_ = v5575
	var v5580 int32
	_ = v5580
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5592 int32
	_ = v5592
	var v5597 int32
	_ = v5597
	var v5599 int32
	_ = v5599
	var v5600 int32
	_ = v5600
	var v5605 int32
	_ = v5605
	var v5606 int32
	_ = v5606
	var v5615 int32
	_ = v5615
	var v5621 int32
	_ = v5621
	var v5622 int32
	_ = v5622
	var v5624 int32
	_ = v5624
	var v5630 int32
	_ = v5630
	var v5635 int32
	_ = v5635
	var v5636 int64
	_ = v5636
	var v5638 int32
	_ = v5638
	var v5639 int32
	_ = v5639
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5657 int32
	_ = v5657
	var v5664 int32
	_ = v5664
	var v5667 int32
	_ = v5667
	var v5672 int32
	_ = v5672
	var v5674 int32
	_ = v5674
	var v5679 int32
	_ = v5679
	var v5680 int32
	_ = v5680
	var v5681 int32
	_ = v5681
	var v5685 int32
	_ = v5685
	var v5687 int32
	_ = v5687
	var v5690 int32
	_ = v5690
	var v5691 int32
	_ = v5691
	var v5695 int64
	_ = v5695
	var v5697 int64
	_ = v5697
	var v5703 int32
	_ = v5703
	var v5706 int32
	_ = v5706
	var v5721 int32
	_ = v5721
	var v5723 int32
	_ = v5723
	var v5731 int32
	_ = v5731
	var v5733 int32
	_ = v5733
	var v5735 int32
	_ = v5735
	var v5738 int32
	_ = v5738
	var v5743 int32
	_ = v5743
	var v5744 int32
	_ = v5744
	var v5746 int32
	_ = v5746
	var v5747 int32
	_ = v5747
	var v5749 int32
	_ = v5749
	var v5750 int32
	_ = v5750
	var v5751 int32
	_ = v5751
	var v5755 int32
	_ = v5755
	var v5759 int32
	_ = v5759
	var v5767 int64
	_ = v5767
	var v5772 int32
	_ = v5772
	var v5773 int32
	_ = v5773
	var v5775 int32
	_ = v5775
	var v5777 int32
	_ = v5777
	var v5780 int32
	_ = v5780
	var v5790 int32
	_ = v5790
	var v5796 int32
	_ = v5796
	var v5797 int32
	_ = v5797
	var v5798 int32
	_ = v5798
	var v5801 int32
	_ = v5801
	var v5802 int32
	_ = v5802
	var v5803 int32
	_ = v5803
	var v5807 int32
	_ = v5807
	var v5808 int32
	_ = v5808
	var v5812 int64
	_ = v5812
	var v5814 int64
	_ = v5814
	var v5820 int32
	_ = v5820
	var v5822 int32
	_ = v5822
	var v5827 int32
	_ = v5827
	var v5828 int32
	_ = v5828
	var v5831 int32
	_ = v5831
	var v5832 int32
	_ = v5832
	var v5836 int64
	_ = v5836
	var v5838 int64
	_ = v5838
	var v5844 int32
	_ = v5844
	var v5849 int32
	_ = v5849
	var v5851 int64
	_ = v5851
	var v5853 int64
	_ = v5853
	var v5859 int32
	_ = v5859
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5874 int32
	_ = v5874
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5890 int32
	_ = v5890
	var v5895 int32
	_ = v5895
	var v5898 int32
	_ = v5898
	var v5901 int32
	_ = v5901
	var v5902 int32
	_ = v5902
	var v5912 int32
	_ = v5912
	var v5917 int32
	_ = v5917
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5964 int32
	_ = v5964
	var v5969 int32
	_ = v5969
	var v5973 int32
	_ = v5973
	var v5974 int32
	_ = v5974
	var v5975 int32
	_ = v5975
	var v5978 int64
	_ = v5978
	var v5979 int64
	_ = v5979
	var v5989 int32
	_ = v5989
	var v6036 int32
	_ = v6036
	var v6044 int32
	_ = v6044
	var v6048 int32
	_ = v6048
	var v6049 int32
	_ = v6049
	var v6054 int32
	_ = v6054
	var v6056 int32
	_ = v6056
	var v6060 int32
	_ = v6060
	var v6064 int32
	_ = v6064
	var v6102 int32
	_ = v6102
	var v6103 int32
	_ = v6103
	var v6104 int32
	_ = v6104
	var v6107 int32
	_ = v6107
	var v6108 int32
	_ = v6108
	var v6115 int32
	_ = v6115
	var v6120 int32
	_ = v6120
	var v6122 int32
	_ = v6122
	var v6136 int32
	_ = v6136
	var v6160 int32
	_ = v6160
	var v6207 int32
	_ = v6207
	var v6215 int32
	_ = v6215
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6224 int32
	_ = v6224
	var v6228 int32
	_ = v6228
	var v6271 int32
	_ = v6271
	var v6272 int32
	_ = v6272
	var v6275 int32
	_ = v6275
	var v6283 int32
	_ = v6283
	var v6288 int32
	_ = v6288
	var v6293 int32
	_ = v6293
	var v6295 int32
	_ = v6295
	var v6297 int32
	_ = v6297
	var v6299 int32
	_ = v6299
	var v6303 int32
	_ = v6303
	var v6305 int32
	_ = v6305
	var v6307 int32
	_ = v6307
	var v6308 int32
	_ = v6308
	var v6311 int32
	_ = v6311
	var v6312 int32
	_ = v6312
	var v6316 int32
	_ = v6316
	var v6317 int32
	_ = v6317
	var v6318 int32
	_ = v6318
	var v6321 int32
	_ = v6321
	var v6322 int32
	_ = v6322
	var v6326 int32
	_ = v6326
	var v6327 int32
	_ = v6327
	var v6330 int32
	_ = v6330
	var v6334 int32
	_ = v6334
	var v6335 int64
	_ = v6335
	var v6337 int64
	_ = v6337
	var v6340 int32
	_ = v6340
	var v6343 int32
	_ = v6343
	var v6344 int32
	_ = v6344
	var v6346 int32
	_ = v6346
	var v6349 int32
	_ = v6349
	var v6350 int32
	_ = v6350
	var v6353 int32
	_ = v6353
	var v6354 int32
	_ = v6354
	var v6355 int32
	_ = v6355
	var v6391 int32
	_ = v6391
	var v6394 int32
	_ = v6394
	var v6397 int32
	_ = v6397
	var v6400 int32
	_ = v6400
	var v6407 int32
	_ = v6407
	var v6412 int32
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6414 int32
	_ = v6414
	var v6419 int32
	_ = v6419
	var v6420 int32
	_ = v6420
	var v6424 int32
	_ = v6424
	var v6428 int32
	_ = v6428
	var v6433 int32
	_ = v6433
	var v6437 int32
	_ = v6437
	var v6438 int32
	_ = v6438
	var v6476 int32
	_ = v6476
	var v6481 int32
	_ = v6481
	var v6485 int32
	_ = v6485
	var v6492 int32
	_ = v6492
	var v6493 int32
	_ = v6493
	var v6497 int32
	_ = v6497
	var v6502 int32
	_ = v6502
	var v6503 int32
	_ = v6503
	var v6505 int32
	_ = v6505
	var v6513 int32
	_ = v6513
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6524 int32
	_ = v6524
	var v6525 int32
	_ = v6525
	var v6529 int32
	_ = v6529
	var v6531 int32
	_ = v6531
	var v6533 int32
	_ = v6533
	var v6537 int32
	_ = v6537
	var v6538 int32
	_ = v6538
	var v6540 int32
	_ = v6540
	var v6543 int32
	_ = v6543
	var v6548 int64
	_ = v6548
	var v6551 int32
	_ = v6551
	var v6553 int32
	_ = v6553
	var v6558 int32
	_ = v6558
	var v6565 int32
	_ = v6565
	var v6566 int32
	_ = v6566
	var v6567 int32
	_ = v6567
	var v6569 int32
	_ = v6569
	var v6570 int32
	_ = v6570
	var v6579 int32
	_ = v6579
	var v6608 int32
	_ = v6608
	var v6614 int32
	_ = v6614
	var v6616 int32
	_ = v6616
	var v6619 int32
	_ = v6619
	var v6624 int32
	_ = v6624
	var v6625 int32
	_ = v6625
	var v6626 int32
	_ = v6626
	var v6627 int32
	_ = v6627
	var v6630 int32
	_ = v6630
	var v6631 int32
	_ = v6631
	var v6634 int32
	_ = v6634
	var v6638 int32
	_ = v6638
	var v6640 int32
	_ = v6640
	var v6644 int32
	_ = v6644
	var v6654 int32
	_ = v6654
	var v6683 int32
	_ = v6683
	var v6687 int32
	_ = v6687
	var v6692 int32
	_ = v6692
	var v6729 int32
	_ = v6729
	var v6731 int32
	_ = v6731
	var v6734 int32
	_ = v6734
	var v6737 int32
	_ = v6737
	var v6742 int32
	_ = v6742
	var v6744 int64
	_ = v6744
	var v6746 int64
	_ = v6746
	var v6749 int32
	_ = v6749
	var v6754 int32
	_ = v6754
	var v6756 int32
	_ = v6756
	var v6757 int64
	_ = v6757
	var v6759 int64
	_ = v6759
	var v6762 int32
	_ = v6762
	var v6763 int64
	_ = v6763
	var v6764 int32
	_ = v6764
	var v6766 int32
	_ = v6766
	var v6767 int64
	_ = v6767
	var v6774 int32
	_ = v6774
	var v6783 int32
	_ = v6783
	var v6784 int32
	_ = v6784
	var v6785 int32
	_ = v6785
	var v6788 int64
	_ = v6788
	var v6789 int64
	_ = v6789
	var v6800 int32
	_ = v6800
	var v6805 int32
	_ = v6805
	var v6809 int32
	_ = v6809
	var v6816 int32
	_ = v6816
	var v6818 int32
	_ = v6818
	var v6820 int32
	_ = v6820
	var v6822 int32
	_ = v6822
	var v6824 int64
	_ = v6824
	var v6826 int64
	_ = v6826
	var v6829 int32
	_ = v6829
	var v6831 int32
	_ = v6831
	var v6833 int32
	_ = v6833
	var v6836 int32
	_ = v6836
	var v6837 int32
	_ = v6837
	var v6838 int32
	_ = v6838
	var v6841 int32
	_ = v6841
	var v6849 int32
	_ = v6849
	var v6851 int32
	_ = v6851
	var v6852 int64
	_ = v6852
	var v6855 int64
	_ = v6855
	var v6861 int32
	_ = v6861
	var v6866 int32
	_ = v6866
	var v6867 int32
	_ = v6867
	var v6871 int32
	_ = v6871
	var v6872 int32
	_ = v6872
	var v6873 int32
	_ = v6873
	var v6876 int32
	_ = v6876
	var v6877 int32
	_ = v6877
	var v6886 int32
	_ = v6886
	var v6892 int32
	_ = v6892
	var v6926 int32
	_ = v6926
	var v6929 int32
	_ = v6929
	var v6932 int32
	_ = v6932
	var v6936 int32
	_ = v6936
	var v6938 int32
	_ = v6938
	var v6941 int32
	_ = v6941
	var v6945 int32
	_ = v6945
	var v6948 int32
	_ = v6948
	var v6953 int32
	_ = v6953
	var v6954 int32
	_ = v6954
	var v6956 int32
	_ = v6956
	var v6957 int64
	_ = v6957
	var v6960 int64
	_ = v6960
	var v6966 int32
	_ = v6966
	var v6971 int32
	_ = v6971
	var v6974 int32
	_ = v6974
	var v6978 int32
	_ = v6978
	var v6979 int32
	_ = v6979
	var v6990 int32
	_ = v6990
	var v7016 int32
	_ = v7016
	var v7024 int32
	_ = v7024
	var v7026 int32
	_ = v7026
	var v7029 int32
	_ = v7029
	var v7030 int64
	_ = v7030
	var v7032 int64
	_ = v7032
	var v7038 int32
	_ = v7038
	var v7040 int32
	_ = v7040
	var v7055 int32
	_ = v7055
	var v7056 int32
	_ = v7056
	var v7060 int32
	_ = v7060
	var v7061 int64
	_ = v7061
	var v7062 int32
	_ = v7062
	var v7064 int32
	_ = v7064
	var v7066 int32
	_ = v7066
	var v7070 int64
	_ = v7070
	var v7076 int32
	_ = v7076
	var v7081 int32
	_ = v7081
	var v7084 int32
	_ = v7084
	var v7086 int32
	_ = v7086
	var v7087 int32
	_ = v7087
	var v7090 int32
	_ = v7090
	var v7092 int32
	_ = v7092
	var v7096 int32
	_ = v7096
	var v7098 int32
	_ = v7098
	var v7102 int32
	_ = v7102
	var v7109 int32
	_ = v7109
	var v7110 int32
	_ = v7110
	var v7114 int32
	_ = v7114
	var v7119 int32
	_ = v7119
	var v7121 int64
	_ = v7121
	var v7127 int32
	_ = v7127
	var v7138 int32
	_ = v7138
	var v7141 int64
	_ = v7141
	var v7143 int64
	_ = v7143
	var v7151 int32
	_ = v7151
	var v7161 int32
	_ = v7161
	var v7162 int32
	_ = v7162
	var v7166 int64
	_ = v7166
	var v7169 int64
	_ = v7169
	var v7175 int32
	_ = v7175
	var v7180 int32
	_ = v7180
	var v7182 int32
	_ = v7182
	var v7183 int32
	_ = v7183
	var v7186 int32
	_ = v7186
	var v7191 int32
	_ = v7191
	var v7193 int32
	_ = v7193
	var v7195 int32
	_ = v7195
	var v7196 int32
	_ = v7196
	var v7202 int64
	_ = v7202
	var v7207 int32
	_ = v7207
	var v7211 int32
	_ = v7211
	var v7213 int32
	_ = v7213
	var v7219 int32
	_ = v7219
	var v7222 int32
	_ = v7222
	var v7224 int32
	_ = v7224
	var v7230 int32
	_ = v7230
	var v7234 int32
	_ = v7234
	var v7236 int32
	_ = v7236
	var v7239 int32
	_ = v7239
	var v7243 int32
	_ = v7243
	var v7248 int32
	_ = v7248
	var v7249 int32
	_ = v7249
	var v7250 int32
	_ = v7250
	var v7253 int32
	_ = v7253
	var v7257 int32
	_ = v7257
	var v7262 int32
	_ = v7262
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7267 int32
	_ = v7267
	var v7271 int32
	_ = v7271
	var v7278 int32
	_ = v7278
	var v7281 int32
	_ = v7281
	var v7289 int32
	_ = v7289
	var v7290 int32
	_ = v7290
	var v7294 int32
	_ = v7294
	var v7295 int32
	_ = v7295
	var v7296 int32
	_ = v7296
	var v7301 int64
	_ = v7301
	var v7302 int64
	_ = v7302
	var v7310 int32
	_ = v7310
	var v7312 int32
	_ = v7312
	var v7313 int32
	_ = v7313
	var v7315 int32
	_ = v7315
	var v7316 int32
	_ = v7316
	var v7322 int64
	_ = v7322
	var v7327 int32
	_ = v7327
	var v7331 int32
	_ = v7331
	var v7333 int32
	_ = v7333
	var v7339 int32
	_ = v7339
	var v7342 int32
	_ = v7342
	var v7344 int32
	_ = v7344
	var v7350 int32
	_ = v7350
	var v7354 int32
	_ = v7354
	var v7356 int32
	_ = v7356
	var v7359 int32
	_ = v7359
	var v7363 int32
	_ = v7363
	var v7368 int32
	_ = v7368
	var v7369 int32
	_ = v7369
	var v7370 int32
	_ = v7370
	var v7373 int32
	_ = v7373
	var v7377 int32
	_ = v7377
	var v7384 int32
	_ = v7384
	var v7387 int32
	_ = v7387
	var v7395 int32
	_ = v7395
	var v7396 int32
	_ = v7396
	var v7400 int32
	_ = v7400
	var v7401 int32
	_ = v7401
	var v7402 int32
	_ = v7402
	var v7407 int64
	_ = v7407
	var v7408 int64
	_ = v7408
	var v7416 int32
	_ = v7416
	var v7417 int32
	_ = v7417
	var v7419 int32
	_ = v7419
	var v7420 int32
	_ = v7420
	var v7421 int32
	_ = v7421
	var v7423 int32
	_ = v7423
	var v7424 int32
	_ = v7424
	var v7427 int32
	_ = v7427
	var v7434 int32
	_ = v7434
	var v7436 int32
	_ = v7436
	var v7437 int32
	_ = v7437
	var v7438 int32
	_ = v7438
	var v7439 int32
	_ = v7439
	var v7449 int32
	_ = v7449
	var v7452 int32
	_ = v7452
	var v7462 int32
	_ = v7462
	var v7463 int64
	_ = v7463
	var v7467 int64
	_ = v7467
	var v7487 int32
	_ = v7487
	var v7491 int32
	_ = v7491
	var v7497 int32
	_ = v7497
	var v7503 int32
	_ = v7503
	var v7504 int32
	_ = v7504
	var v7505 int32
	_ = v7505
	var v7508 int32
	_ = v7508
	var v7510 int32
	_ = v7510
	var v7514 int32
	_ = v7514
	var v7515 int32
	_ = v7515
	var v7518 int32
	_ = v7518
	var v7524 int32
	_ = v7524
	var v7525 int64
	_ = v7525
	var v7529 int32
	_ = v7529
	var v7530 int32
	_ = v7530
	var v7531 int32
	_ = v7531
	var v7534 int64
	_ = v7534
	var v7535 int64
	_ = v7535
	var v7543 int64
	_ = v7543
	var v7547 int64
	_ = v7547
	var v7553 int64
	_ = v7553
	var v7562 int64
	_ = v7562
	var v7565 int32
	_ = v7565
	var v7569 int32
	_ = v7569
	var v7572 int32
	_ = v7572
	var v7577 int32
	_ = v7577
	var v7579 int32
	_ = v7579
	var v7580 int32
	_ = v7580
	var v7581 int32
	_ = v7581
	var v7619 int64
	_ = v7619
	var v7623 int32
	_ = v7623
	var v7624 int32
	_ = v7624
	var v7625 int32
	_ = v7625
	var v7628 int64
	_ = v7628
	var v7629 int64
	_ = v7629
	var v7637 int64
	_ = v7637
	var v7640 int64
	_ = v7640
	var v7646 int64
	_ = v7646
	var v7655 int64
	_ = v7655
	var v7658 int32
	_ = v7658
	var v7663 int32
	_ = v7663
	var v7664 int32
	_ = v7664
	var v7670 int32
	_ = v7670
	var v7675 int32
	_ = v7675
	var v7677 int32
	_ = v7677
	var v7682 int32
	_ = v7682
	var v7683 int32
	_ = v7683
	var v7685 int32
	_ = v7685
	var v7688 int32
	_ = v7688
	var v7693 int32
	_ = v7693
	var v7695 int32
	_ = v7695
	var v7696 int32
	_ = v7696
	var v7697 int32
	_ = v7697
	var v7737 int32
	_ = v7737
	var v7738 int32
	_ = v7738
	var v7743 int32
	_ = v7743
	var v7780 int32
	_ = v7780
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7788 int32
	_ = v7788
	var v7793 int32
	_ = v7793
	var v7795 int32
	_ = v7795
	var v7796 int32
	_ = v7796
	var v7797 int32
	_ = v7797
	var v7799 int32
	_ = v7799
	var v7803 int32
	_ = v7803
	var v7804 int32
	_ = v7804
	var v7805 int32
	_ = v7805
	var v7806 int32
	_ = v7806
	var v7808 int32
	_ = v7808
	var v7811 int64
	_ = v7811
	var v7813 int32
	_ = v7813
	var v7814 int32
	_ = v7814
	var v7820 int32
	_ = v7820
	var v7823 int32
	_ = v7823
	var v7826 int32
	_ = v7826
	var v7827 int32
	_ = v7827
	var v7830 int32
	_ = v7830
	var v7837 int32
	_ = v7837
	var v7838 int32
	_ = v7838
	var v7839 int32
	_ = v7839
	var v7841 int32
	_ = v7841
	var v7847 int32
	_ = v7847
	var v7853 int32
	_ = v7853
	var v7858 int64
	_ = v7858
	var v7861 int32
	_ = v7861
	var v7863 int32
	_ = v7863
	var v7866 int32
	_ = v7866
	var v7869 int32
	_ = v7869
	var v7872 int32
	_ = v7872
	var v7877 int32
	_ = v7877
	var v7878 int64
	_ = v7878
	var v7880 int32
	_ = v7880
	var v7883 int32
	_ = v7883
	var v7887 int32
	_ = v7887
	var v7890 int32
	_ = v7890
	var v7894 int32
	_ = v7894
	var v7896 int32
	_ = v7896
	var v7897 int32
	_ = v7897
	var v7898 int32
	_ = v7898
	var v7900 int32
	_ = v7900
	var v7905 int32
	_ = v7905
	var v7906 int64
	_ = v7906
	var v7907 int64
	_ = v7907
	var v7909 int64
	_ = v7909
	var v7911 int64
	_ = v7911
	var v7918 int32
	_ = v7918
	var v7919 int32
	_ = v7919
	var v7920 int32
	_ = v7920
	var v7921 int32
	_ = v7921
	var v7925 int64
	_ = v7925
	var v7931 int32
	_ = v7931
	var v7936 int32
	_ = v7936
	var v7939 int64
	_ = v7939
	var v7941 int64
	_ = v7941
	var v7942 int32
	_ = v7942
	var v7943 int64
	_ = v7943
	var v7946 int32
	_ = v7946
	var v7947 int32
	_ = v7947
	var v7952 int32
	_ = v7952
	var v7957 int32
	_ = v7957
	var v7963 int64
	_ = v7963
	var v7966 int64
	_ = v7966
	var v7967 int64
	_ = v7967
	var v7970 int64
	_ = v7970
	var v7976 int32
	_ = v7976
	var v7981 int32
	_ = v7981
	var v7987 int32
	_ = v7987
	var v7989 int32
	_ = v7989
	var v7992 int32
	_ = v7992
	var v7996 int32
	_ = v7996
	var v7997 int32
	_ = v7997
	var v7999 int32
	_ = v7999
	var v8000 int32
	_ = v8000
	var v8005 int32
	_ = v8005
	var v8006 int32
	_ = v8006
	var v8008 int32
	_ = v8008
	var v8009 int32
	_ = v8009
	var v8011 int32
	_ = v8011
	var v8012 int32
	_ = v8012
	var v8013 int32
	_ = v8013
	var v8014 int32
	_ = v8014
	var v8019 int32
	_ = v8019
	var v8032 int32
	_ = v8032
	var v8058 int32
	_ = v8058
	var v8060 int32
	_ = v8060
	var v8062 int32
	_ = v8062
	var v8064 int32
	_ = v8064
	var v8065 int32
	_ = v8065
	var v8067 int32
	_ = v8067
	var v8068 int32
	_ = v8068
	var v8072 int32
	_ = v8072
	var v8073 int32
	_ = v8073
	var v8077 int32
	_ = v8077
	var v8078 int32
	_ = v8078
	var v8080 int64
	_ = v8080
	var v8082 int32
	_ = v8082
	var v8084 int32
	_ = v8084
	var v8092 int32
	_ = v8092
	var v8095 int32
	_ = v8095
	var v8099 int32
	_ = v8099
	var v8100 int64
	_ = v8100
	var v8102 int32
	_ = v8102
	var v8106 int32
	_ = v8106
	var v8107 int32
	_ = v8107
	var v8110 int32
	_ = v8110
	var v8111 int32
	_ = v8111
	var v8116 int32
	_ = v8116
	var v8118 int32
	_ = v8118
	var v8122 int32
	_ = v8122
	var v8128 int32
	_ = v8128
	var v8130 int32
	_ = v8130
	var v8136 int32
	_ = v8136
	var v8140 int32
	_ = v8140
	var v8141 int64
	_ = v8141
	var v8143 int32
	_ = v8143
	var v8144 int64
	_ = v8144
	var v8149 int32
	_ = v8149
	var v8150 int32
	_ = v8150
	var v8151 int32
	_ = v8151
	var v8155 int32
	_ = v8155
	var v8156 int32
	_ = v8156
	var v8158 int32
	_ = v8158
	var v8160 int32
	_ = v8160
	var v8161 int32
	_ = v8161
	var v8163 int32
	_ = v8163
	var v8165 int32
	_ = v8165
	var v8167 int32
	_ = v8167
	var v8168 int32
	_ = v8168
	var v8176 int32
	_ = v8176
	var v8177 int32
	_ = v8177
	var v8178 int32
	_ = v8178
	var v8181 int32
	_ = v8181
	var v8182 int32
	_ = v8182
	var v8184 int32
	_ = v8184
	var v8185 int32
	_ = v8185
	var v8187 int32
	_ = v8187
	var v8189 int32
	_ = v8189
	var v8199 int32
	_ = v8199
	var v8200 int32
	_ = v8200
	var v8201 int32
	_ = v8201
	var v8204 int32
	_ = v8204
	var v8205 int32
	_ = v8205
	var v8206 int32
	_ = v8206
	var v8209 int32
	_ = v8209
	var v8210 int32
	_ = v8210
	var v8212 int32
	_ = v8212
	var v8217 int32
	_ = v8217
	var v8230 int32
	_ = v8230
	var v8233 int32
	_ = v8233
	var v8234 int32
	_ = v8234
	var v8235 int32
	_ = v8235
	var v8274 int32
	_ = v8274
	var v8277 int32
	_ = v8277
	var v8280 int32
	_ = v8280
	var v8285 int32
	_ = v8285
	var v8287 int32
	_ = v8287
	var v8288 int64
	_ = v8288
	var v8290 int64
	_ = v8290
	var v8293 int32
	_ = v8293
	var v8297 int32
	_ = v8297
	var v8301 int32
	_ = v8301
	var v8306 int32
	_ = v8306
	var v8308 int32
	_ = v8308
	var v8310 int32
	_ = v8310
	var v8313 int32
	_ = v8313
	var v8315 int32
	_ = v8315
	var v8316 int64
	_ = v8316
	var v8318 int32
	_ = v8318
	var v8319 int32
	_ = v8319
	var v8321 int32
	_ = v8321
	var v8327 int32
	_ = v8327
	var v8328 int64
	_ = v8328
	var v8330 int32
	_ = v8330
	var v8333 int32
	_ = v8333
	var v8334 int64
	_ = v8334
	var v8336 int32
	_ = v8336
	var v8339 int32
	_ = v8339
	var v8340 int64
	_ = v8340
	var v8342 int32
	_ = v8342
	var v8344 int32
	_ = v8344
	var v8348 int32
	_ = v8348
	var v8349 int32
	_ = v8349
	var v8350 int32
	_ = v8350
	var v8351 int32
	_ = v8351
	var v8353 int32
	_ = v8353
	var v8364 int32
	_ = v8364
	var v8366 int32
	_ = v8366
	var v8368 int32
	_ = v8368
	var v8371 int32
	_ = v8371
	var v8374 int32
	_ = v8374
	var v8377 int32
	_ = v8377
	var v8378 int32
	_ = v8378
	var v8381 int32
	_ = v8381
	var v8382 int32
	_ = v8382
	var v8385 int32
	_ = v8385
	var v8392 int32
	_ = v8392
	var v8393 int32
	_ = v8393
	var v8396 int32
	_ = v8396
	var v8405 int64
	_ = v8405
	var v8407 int32
	_ = v8407
	var v8414 int32
	_ = v8414
	var v8418 int32
	_ = v8418
	var v8430 int32
	_ = v8430
	var v8431 int32
	_ = v8431
	var v8432 int32
	_ = v8432
	var v8434 int32
	_ = v8434
	var v8438 int32
	_ = v8438
	var v8439 int32
	_ = v8439
	var v8441 int32
	_ = v8441
	var v8442 int32
	_ = v8442
	var v8443 int32
	_ = v8443
	var v8445 int32
	_ = v8445
	var v8451 int32
	_ = v8451
	var v8452 int32
	_ = v8452
	var v8453 int32
	_ = v8453
	var v8454 int32
	_ = v8454
	var v8457 int32
	_ = v8457
	var v8464 int32
	_ = v8464
	var v8465 int32
	_ = v8465
	var v8466 int32
	_ = v8466
	var v8469 int32
	_ = v8469
	var v8472 int32
	_ = v8472
	var v8477 int32
	_ = v8477
	var v8478 int32
	_ = v8478
	var v8480 int32
	_ = v8480
	var v8482 int32
	_ = v8482
	var v8486 int32
	_ = v8486
	var v8487 int32
	_ = v8487
	var v8488 int32
	_ = v8488
	var v8493 int32
	_ = v8493
	var v8494 int32
	_ = v8494
	var v8495 int32
	_ = v8495
	var v8498 int32
	_ = v8498
	var v8499 int32
	_ = v8499
	var v8500 int32
	_ = v8500
	var v8502 int32
	_ = v8502
	var v8506 int32
	_ = v8506
	var v8507 int32
	_ = v8507
	var v8509 int32
	_ = v8509
	var v8511 int32
	_ = v8511
	var v8513 int32
	_ = v8513
	var v8514 int32
	_ = v8514
	var v8517 int32
	_ = v8517
	var v8524 int32
	_ = v8524
	var v8529 int32
	_ = v8529
	var v8530 int32
	_ = v8530
	var v8534 int64
	_ = v8534
	var v8535 int32
	_ = v8535
	var v8536 int32
	_ = v8536
	var v8544 int32
	_ = v8544
	var v8549 int32
	_ = v8549
	var v8553 int32
	_ = v8553
	var v8558 int64
	_ = v8558
	var v8560 int64
	_ = v8560
	var v8563 int32
	_ = v8563
	var v8571 int32
	_ = v8571
	var v8578 int32
	_ = v8578
	var v8579 int32
	_ = v8579
	var v8583 int64
	_ = v8583
	var v8586 int64
	_ = v8586
	var v8592 int32
	_ = v8592
	var v8597 int32
	_ = v8597
	var v8602 int32
	_ = v8602
	var v8603 int32
	_ = v8603
	var v8604 int32
	_ = v8604
	var v8611 int32
	_ = v8611
	var v8614 int32
	_ = v8614
	var v8622 int32
	_ = v8622
	var v8623 int64
	_ = v8623
	var v8625 int32
	_ = v8625
	var v8628 int32
	_ = v8628
	var v8633 int32
	_ = v8633
	var v8635 int32
	_ = v8635
	var v8637 int32
	_ = v8637
	var v8640 int32
	_ = v8640
	var v8642 int32
	_ = v8642
	var v8643 int64
	_ = v8643
	var v8645 int32
	_ = v8645
	var v8648 int32
	_ = v8648
	var v8649 int32
	_ = v8649
	var v8651 int32
	_ = v8651
	var v8652 int32
	_ = v8652
	var v8658 int64
	_ = v8658
	var v8663 int32
	_ = v8663
	var v8667 int32
	_ = v8667
	var v8669 int32
	_ = v8669
	var v8675 int32
	_ = v8675
	var v8678 int32
	_ = v8678
	var v8680 int32
	_ = v8680
	var v8686 int32
	_ = v8686
	var v8690 int32
	_ = v8690
	var v8692 int32
	_ = v8692
	var v8695 int32
	_ = v8695
	var v8699 int32
	_ = v8699
	var v8704 int32
	_ = v8704
	var v8705 int32
	_ = v8705
	var v8706 int32
	_ = v8706
	var v8709 int32
	_ = v8709
	var v8713 int32
	_ = v8713
	var v8718 int32
	_ = v8718
	var v8719 int32
	_ = v8719
	var v8720 int32
	_ = v8720
	var v8723 int32
	_ = v8723
	var v8727 int32
	_ = v8727
	var v8734 int32
	_ = v8734
	var v8737 int32
	_ = v8737
	var v8745 int32
	_ = v8745
	var v8746 int32
	_ = v8746
	var v8750 int32
	_ = v8750
	var v8751 int32
	_ = v8751
	var v8752 int32
	_ = v8752
	var v8757 int64
	_ = v8757
	var v8758 int64
	_ = v8758
	var v8766 int32
	_ = v8766
	var v8767 int32
	_ = v8767
	var v8768 int32
	_ = v8768
	var v8770 int32
	_ = v8770
	var v8771 int32
	_ = v8771
	var v8777 int64
	_ = v8777
	var v8782 int32
	_ = v8782
	var v8786 int32
	_ = v8786
	var v8788 int32
	_ = v8788
	var v8794 int32
	_ = v8794
	var v8797 int32
	_ = v8797
	var v8799 int32
	_ = v8799
	var v8805 int32
	_ = v8805
	var v8809 int32
	_ = v8809
	var v8811 int32
	_ = v8811
	var v8814 int32
	_ = v8814
	var v8818 int32
	_ = v8818
	var v8823 int32
	_ = v8823
	var v8824 int32
	_ = v8824
	var v8825 int32
	_ = v8825
	var v8828 int32
	_ = v8828
	var v8832 int32
	_ = v8832
	var v8839 int32
	_ = v8839
	var v8842 int32
	_ = v8842
	var v8850 int32
	_ = v8850
	var v8851 int32
	_ = v8851
	var v8855 int32
	_ = v8855
	var v8856 int32
	_ = v8856
	var v8857 int32
	_ = v8857
	var v8862 int64
	_ = v8862
	var v8863 int64
	_ = v8863
	var v8871 int32
	_ = v8871
	var v8872 int32
	_ = v8872
	var v8873 int32
	_ = v8873
	var v8875 int32
	_ = v8875
	var v8879 int32
	_ = v8879
	var v8885 int32
	_ = v8885
	var v8890 int32
	_ = v8890
	var v8898 int32
	_ = v8898
	var v8902 int32
	_ = v8902
	var v8903 int32
	_ = v8903
	var v8907 int32
	_ = v8907
	var v8909 int64
	_ = v8909
	var v8910 int32
	_ = v8910
	var v8911 int32
	_ = v8911
	var v8918 int32
	_ = v8918
	var v8923 int32
	_ = v8923
	var v8926 int32
	_ = v8926
	var v8927 int32
	_ = v8927
	var v8931 int32
	_ = v8931
	var v8933 int64
	_ = v8933
	var v8934 int32
	_ = v8934
	var v8935 int32
	_ = v8935
	var v8942 int32
	_ = v8942
	var v8947 int32
	_ = v8947
	var v8949 int32
	_ = v8949
	var v8955 int32
	_ = v8955
	var v8962 int32
	_ = v8962
	var v8963 int32
	_ = v8963
	var v8967 int32
	_ = v8967
	var v8972 int32
	_ = v8972
	var v8974 int32
	_ = v8974
	var v8977 int64
	_ = v8977
	var v8983 int32
	_ = v8983
	var v8991 int32
	_ = v8991
	var v8998 int32
	_ = v8998
	var v9003 int32
	_ = v9003
	var v9008 int32
	_ = v9008
	var v9015 int32
	_ = v9015
	var v9020 int32
	_ = v9020
	var v9024 int32
	_ = v9024
	var v9027 int64
	_ = v9027
	var v9030 int32
	_ = v9030
	var v9033 int64
	_ = v9033
	var v9039 int32
	_ = v9039
	var v9044 int32
	_ = v9044
	var v9048 int32
	_ = v9048
	var v9049 int64
	_ = v9049
	var v9051 int64
	_ = v9051
	var v9052 int64
	_ = v9052
	var v9056 int64
	_ = v9056
	var v9062 int32
	_ = v9062
	var v9067 int32
	_ = v9067
	var v9071 int32
	_ = v9071
	var v9074 int32
	_ = v9074
	var v9075 int32
	_ = v9075
	var v9081 int32
	_ = v9081
	var v9086 int32
	_ = v9086
	var v9090 int32
	_ = v9090
	var v9091 int32
	_ = v9091
	var v9093 int32
	_ = v9093
	var v9095 int64
	_ = v9095
	var v9097 int32
	_ = v9097
	var v9103 int32
	_ = v9103
	var v9108 int32
	_ = v9108
	var v9117 int32
	_ = v9117
	var v9119 int32
	_ = v9119
	var v9122 int32
	_ = v9122
	var v9123 int32
	_ = v9123
	var v9126 int32
	_ = v9126
	var v9127 int32
	_ = v9127
	var v9133 int32
	_ = v9133
	var v9138 int32
	_ = v9138
	var v9140 int64
	_ = v9140
	var v9150 int32
	_ = v9150
	var v9157 int32
	_ = v9157
	var v9158 int32
	_ = v9158
	var v9162 int32
	_ = v9162
	var v9164 int64
	_ = v9164
	var v9165 int32
	_ = v9165
	var v9166 int32
	_ = v9166
	var v9173 int32
	_ = v9173
	var v9178 int32
	_ = v9178
	var v9182 int32
	_ = v9182
	var v9184 int64
	_ = v9184
	var v9185 int32
	_ = v9185
	var v9186 int32
	_ = v9186
	var v9193 int32
	_ = v9193
	var v9198 int32
	_ = v9198
	var v9236 int32
	_ = v9236
	var v9239 int32
	_ = v9239
	var v9241 int32
	_ = v9241
	var v9244 int32
	_ = v9244
	var v9246 int32
	_ = v9246
	var v9249 int32
	_ = v9249
	var v9254 int32
	_ = v9254
	var v9256 int32
	_ = v9256
	var v9257 int32
	_ = v9257
	var v9262 int32
	_ = v9262
	var v9267 int32
	_ = v9267
	var v9280 int32
	_ = v9280
	var v9313 int32
	_ = v9313
	var v9342 int32
	_ = v9342
	var v9345 int32
	_ = v9345
	var v9348 int32
	_ = v9348
	var v9352 int32
	_ = v9352
	var v9354 int32
	_ = v9354
	var v9357 int32
	_ = v9357
	var v9361 int32
	_ = v9361
	var v9364 int32
	_ = v9364
	var v9369 int32
	_ = v9369
	var v9370 int32
	_ = v9370
	var v9372 int32
	_ = v9372
	var v9373 int64
	_ = v9373
	var v9376 int32
	_ = v9376
	var v9377 int32
	_ = v9377
	var v9381 int64
	_ = v9381
	var v9387 int32
	_ = v9387
	var v9392 int32
	_ = v9392
	var v9395 int32
	_ = v9395
	var v9398 int32
	_ = v9398
	var v9403 int32
	_ = v9403
	var v9405 int32
	_ = v9405
	var v9406 int64
	_ = v9406
	var v9407 int32
	_ = v9407
	var v9414 int32
	_ = v9414
	var v9415 int32
	_ = v9415
	var v9418 int32
	_ = v9418
	var v9419 int32
	_ = v9419
	var v9423 int32
	_ = v9423
	var v9428 int32
	_ = v9428
	var v9430 int32
	_ = v9430
	var v9444 int32
	_ = v9444
	var v9469 int32
	_ = v9469
	var v9475 int32
	_ = v9475
	var v9482 int32
	_ = v9482
	var v9486 int32
	_ = v9486
	var v9491 int32
	_ = v9491
	var v9495 int32
	_ = v9495
	var v9498 int32
	_ = v9498
	var v9502 int32
	_ = v9502
	var v9507 int32
	_ = v9507
	var v9545 int32
	_ = v9545
	var v9547 int32
	_ = v9547
	var v9550 int32
	_ = v9550
	var v9551 int32
	_ = v9551
	var v9553 int32
	_ = v9553
	var v9556 int32
	_ = v9556
	var v9559 int32
	_ = v9559
	var v9564 int32
	_ = v9564
	var v9566 int32
	_ = v9566
	var v9567 int32
	_ = v9567
	var v9569 int32
	_ = v9569
	var v9572 int32
	_ = v9572
	var v9573 int32
	_ = v9573
	var v9580 int32
	_ = v9580
	var v9581 int32
	_ = v9581
	var v9619 int32
	_ = v9619
	var v9623 int32
	_ = v9623
	var v9624 int32
	_ = v9624
	var v9630 int32
	_ = v9630
	var v9631 int32
	_ = v9631
	var v9636 int32
	_ = v9636
	var v9638 int32
	_ = v9638
	var v9642 int32
	_ = v9642
	var v9644 int32
	_ = v9644
	var v9647 int32
	_ = v9647
	var v9652 int32
	_ = v9652
	var v9654 int32
	_ = v9654
	var v9655 int32
	_ = v9655
	var v9658 int32
	_ = v9658
	var v9664 int32
	_ = v9664
	var v9697 int32
	_ = v9697
	var v9702 int32
	_ = v9702
	var v9706 int32
	_ = v9706
	var v9710 int32
	_ = v9710
	var v9711 int32
	_ = v9711
	var v9713 int32
	_ = v9713
	var v9715 int32
	_ = v9715
	var v9720 int32
	_ = v9720
	var v9730 int32
	_ = v9730
	var v9733 int32
	_ = v9733
	var v9734 int32
	_ = v9734
	var v9735 int32
	_ = v9735
	var v9749 int64
	_ = v9749
	var v9759 int32
	_ = v9759
	var v9760 int32
	_ = v9760
	var v9763 int32
	_ = v9763
	var v9771 int32
	_ = v9771
	var v9772 int32
	_ = v9772
	var v9773 int32
	_ = v9773
	var v9776 int64
	_ = v9776
	var v9777 int64
	_ = v9777
	var v9786 int64
	_ = v9786
	var v9789 int32
	_ = v9789
	var v9792 int32
	_ = v9792
	var v9793 int32
	_ = v9793
	var v9797 int32
	_ = v9797
	var v9801 int32
	_ = v9801
	var v9803 int32
	_ = v9803
	var v9805 int32
	_ = v9805
	var v9806 int32
	_ = v9806
	var v9807 int32
	_ = v9807
	var v9808 int32
	_ = v9808
	var v9809 int64
	_ = v9809
	var v9811 int32
	_ = v9811
	var v9851 int32
	_ = v9851
	var v9855 int32
	_ = v9855
	var v9893 int32
	_ = v9893
	var v9896 int32
	_ = v9896
	var v9901 int32
	_ = v9901
	var v9902 int32
	_ = v9902
	var v9903 int32
	_ = v9903
	var v9905 int32
	_ = v9905
	var v9909 int32
	_ = v9909
	var v9910 int64
	_ = v9910
	var v9912 int32
	_ = v9912
	var v9914 int32
	_ = v9914
	var v9917 int32
	_ = v9917
	var v9918 int32
	_ = v9918
	var v9920 int32
	_ = v9920
	var v9921 int64
	_ = v9921
	var v9922 int32
	_ = v9922
	var v9925 int32
	_ = v9925
	var v9929 int32
	_ = v9929
	var v9932 int32
	_ = v9932
	var v9935 int32
	_ = v9935
	var v9941 int64
	_ = v9941
	var v9944 int32
	_ = v9944
	var v9945 int32
	_ = v9945
	var v9946 int32
	_ = v9946
	var v9948 int32
	_ = v9948
	var v9949 int32
	_ = v9949
	var v9954 int32
	_ = v9954
	var v9955 int64
	_ = v9955
	var v9959 int32
	_ = v9959
	var v9963 int32
	_ = v9963
	var v9968 int32
	_ = v9968
	var v9969 int32
	_ = v9969
	var v9975 int32
	_ = v9975
	var v9976 int32
	_ = v9976
	var v9978 int32
	_ = v9978
	var v9980 int64
	_ = v9980
	var v9981 int32
	_ = v9981
	var v9982 int32
	_ = v9982
	var v9986 int32
	_ = v9986
	var v9994 int32
	_ = v9994
	var v9995 int32
	_ = v9995
	var v9997 int64
	_ = v9997
	var v10002 int32
	_ = v10002
	var v10003 int32
	_ = v10003
	var v10006 int64
	_ = v10006
	var v10014 int32
	_ = v10014
	var v10015 int32
	_ = v10015
	var v10024 int32
	_ = v10024
	var v10025 int32
	_ = v10025
	var v10031 int32
	_ = v10031
	var v10032 int32
	_ = v10032
	var v10038 int32
	_ = v10038
	var v10039 int32
	_ = v10039
	var v10044 int32
	_ = v10044
	var v10045 int32
	_ = v10045
	var v10051 int64
	_ = v10051
	var v10054 int64
	_ = v10054
	var v10057 int32
	_ = v10057
	var v10060 int32
	_ = v10060
	var v10067 int32
	_ = v10067
	var v10070 int32
	_ = v10070
	var v10074 int32
	_ = v10074
	var v10076 int64
	_ = v10076
	var v10078 int64
	_ = v10078
	var v10082 int32
	_ = v10082
	var v10085 int32
	_ = v10085
	var v10088 int64
	_ = v10088
	var v10091 int32
	_ = v10091
	var v10097 int32
	_ = v10097
	var v10100 int32
	_ = v10100
	var v10104 int32
	_ = v10104
	var v10109 int32
	_ = v10109
	var v10112 int32
	_ = v10112
	var v10114 int32
	_ = v10114
	var v10116 int32
	_ = v10116
	var v10117 int32
	_ = v10117
	var v10119 int32
	_ = v10119
	var v10123 int32
	_ = v10123
	var v10124 int32
	_ = v10124
	var v10126 int32
	_ = v10126
	var v10127 int32
	_ = v10127
	var v10130 int32
	_ = v10130
	var v10134 int32
	_ = v10134
	var v10136 int32
	_ = v10136
	var v10139 int32
	_ = v10139
	var v10141 int32
	_ = v10141
	var v10142 int32
	_ = v10142
	var v10143 int32
	_ = v10143
	var v10145 int32
	_ = v10145
	var v10148 int32
	_ = v10148
	var v10149 int32
	_ = v10149
	var v10155 int32
	_ = v10155
	var v10160 int32
	_ = v10160
	var v10164 int32
	_ = v10164
	var v10168 int32
	_ = v10168
	var v10169 int64
	_ = v10169
	var v10170 int64
	_ = v10170
	var v10171 int64
	_ = v10171
	var v10176 int64
	_ = v10176
	var v10177 int64
	_ = v10177
	var v10180 int64
	_ = v10180
	var v10183 int32
	_ = v10183
	var v10188 int32
	_ = v10188
	var v10189 int32
	_ = v10189
	var v10191 int32
	_ = v10191
	var v10192 int32
	_ = v10192
	var v10195 int32
	_ = v10195
	var v10198 int32
	_ = v10198
	var v10203 int32
	_ = v10203
	var v10204 int32
	_ = v10204
	var v10205 int32
	_ = v10205
	var v10208 int32
	_ = v10208
	var v10209 int32
	_ = v10209
	var v10213 int32
	_ = v10213
	var v10220 int32
	_ = v10220
	var v10256 int32
	_ = v10256
	var v10267 int32
	_ = v10267
	var v10272 int32
	_ = v10272
	var v10275 int32
	_ = v10275
	var v10276 int32
	_ = v10276
	var v10281 int32
	_ = v10281
	var v10286 int32
	_ = v10286
	var v10296 int32
	_ = v10296
	var v10301 int32
	_ = v10301
	var v10303 int32
	_ = v10303
	var v10312 int32
	_ = v10312
	var v10317 int32
	_ = v10317
	var v10318 int32
	_ = v10318
	var v10322 int32
	_ = v10322
	var v10326 int32
	_ = v10326
	var v10328 int32
	_ = v10328
	var v10367 int32
	_ = v10367
	var v10372 int32
	_ = v10372
	var v10377 int32
	_ = v10377
	var v10381 int32
	_ = v10381
	var v10386 int32
	_ = v10386
	var v10392 int32
	_ = v10392
	var v10393 int32
	_ = v10393
	var v10395 int32
	_ = v10395
	var v10396 int32
	_ = v10396
	var v10400 int32
	_ = v10400
	var v10408 int32
	_ = v10408
	var v10413 int32
	_ = v10413
	var v10415 int32
	_ = v10415
	var v10418 int32
	_ = v10418
	var v10419 int32
	_ = v10419
	var v10420 int32
	_ = v10420
	var v10421 int32
	_ = v10421
	var v10428 int32
	_ = v10428
	var v10429 int32
	_ = v10429
	var v10433 int32
	_ = v10433
	var v10437 int32
	_ = v10437
	var v10442 int32
	_ = v10442
	var v10443 int32
	_ = v10443
	var v10444 int32
	_ = v10444
	var v10445 int32
	_ = v10445
	var v10485 int64
	_ = v10485
	var v10486 int64
	_ = v10486
	var v10487 int64
	_ = v10487
	var v10490 int64
	_ = v10490
	var v10493 int32
	_ = v10493
	var v10498 int32
	_ = v10498
	var v10499 int32
	_ = v10499
	var v10501 int32
	_ = v10501
	var v10502 int32
	_ = v10502
	var v10507 int32
	_ = v10507
	var v10508 int32
	_ = v10508
	var v10509 int32
	_ = v10509
	var v10514 int32
	_ = v10514
	var v10515 int32
	_ = v10515
	var v10517 int32
	_ = v10517
	var v10518 int32
	_ = v10518
	var v10519 int32
	_ = v10519
	var v10521 int32
	_ = v10521
	var v10523 int32
	_ = v10523
	var v10526 int32
	_ = v10526
	var v10531 int32
	_ = v10531
	var v10532 int32
	_ = v10532
	var v10533 int32
	_ = v10533
	var v10535 int32
	_ = v10535
	var v10536 int32
	_ = v10536
	var v10540 int32
	_ = v10540
	var v10545 int32
	_ = v10545
	var v10550 int32
	_ = v10550
	var v10551 int32
	_ = v10551
	var v10557 int32
	_ = v10557
	var v10558 int32
	_ = v10558
	var v10566 int32
	_ = v10566
	var v10567 int32
	_ = v10567
	var v10572 int32
	_ = v10572
	var v10573 int32
	_ = v10573
	var v10577 int32
	_ = v10577
	var v10579 int32
	_ = v10579
	var v10580 int32
	_ = v10580
	var v10586 int32
	_ = v10586
	var v10588 int32
	_ = v10588
	var v10601 int32
	_ = v10601
	var v10630 int32
	_ = v10630
	var v10635 int32
	_ = v10635
	var v10639 int32
	_ = v10639
	var v10640 int32
	_ = v10640
	var v10642 int32
	_ = v10642
	var v10643 int32
	_ = v10643
	var v10644 int32
	_ = v10644
	var v10650 int32
	_ = v10650
	var v10654 int32
	_ = v10654
	var v10656 int32
	_ = v10656
	var v10661 int32
	_ = v10661
	var v10662 int32
	_ = v10662
	var v10667 int32
	_ = v10667
	var v10669 int32
	_ = v10669
	var v10675 int32
	_ = v10675
	var v10680 int32
	_ = v10680
	var v10681 int32
	_ = v10681
	var v10682 int32
	_ = v10682
	var v10684 int32
	_ = v10684
	var v10685 int32
	_ = v10685
	var v10688 int32
	_ = v10688
	var v10693 int32
	_ = v10693
	var v10695 int32
	_ = v10695
	var v10701 int32
	_ = v10701
	var v10706 int32
	_ = v10706
	var v10710 int32
	_ = v10710
	var v10712 int32
	_ = v10712
	var v10720 int32
	_ = v10720
	var v10725 int32
	_ = v10725
	var v10727 int32
	_ = v10727
	var v10734 int32
	_ = v10734
	var v10736 int32
	_ = v10736
	var v10744 int32
	_ = v10744
	var v10749 int32
	_ = v10749
	var v10758 int32
	_ = v10758
	var v10791 int64
	_ = v10791
	var v10794 int32
	_ = v10794
	var v10799 int32
	_ = v10799
	var v10800 int32
	_ = v10800
	var v10801 int32
	_ = v10801
	var v10806 int32
	_ = v10806
	var v10809 int32
	_ = v10809
	var v10811 int32
	_ = v10811
	var v10812 int32
	_ = v10812
	var v10813 int32
	_ = v10813
	var v10816 int32
	_ = v10816
	var v10821 int32
	_ = v10821
	var v10826 int32
	_ = v10826
	var v10830 int32
	_ = v10830
	var v10835 int32
	_ = v10835
	var v10841 int32
	_ = v10841
	var v10842 int32
	_ = v10842
	var v10844 int32
	_ = v10844
	var v10845 int32
	_ = v10845
	var v10849 int32
	_ = v10849
	var v10857 int32
	_ = v10857
	var v10862 int32
	_ = v10862
	var v10864 int32
	_ = v10864
	var v10867 int32
	_ = v10867
	var v10868 int32
	_ = v10868
	var v10871 int32
	_ = v10871
	var v10876 int32
	_ = v10876
	var v10877 int32
	_ = v10877
	var v10881 int32
	_ = v10881
	var v10882 int32
	_ = v10882
	var v10884 int32
	_ = v10884
	var v10889 int32
	_ = v10889
	var v10894 int32
	_ = v10894
	var v10895 int32
	_ = v10895
	var v10897 int32
	_ = v10897
	var v10902 int32
	_ = v10902
	var v10903 int32
	_ = v10903
	var v10905 int32
	_ = v10905
	var v10906 int32
	_ = v10906
	var v10909 int32
	_ = v10909
	var v10914 int32
	_ = v10914
	var v10916 int32
	_ = v10916
	var v10922 int32
	_ = v10922
	var v10927 int32
	_ = v10927
	var v10931 int32
	_ = v10931
	var v10933 int32
	_ = v10933
	var v10941 int32
	_ = v10941
	var v10946 int32
	_ = v10946
	var v10986 int32
	_ = v10986
	var v10988 int32
	_ = v10988
	var v10996 int32
	_ = v10996
	var v11001 int32
	_ = v11001
	var v11004 int32
	_ = v11004
	var v11005 int32
	_ = v11005
	var v11011 int32
	_ = v11011
	var v11016 int32
	_ = v11016
	var v11028 int32
	_ = v11028
	var v11054 int32
	_ = v11054
	var v11057 int32
	_ = v11057
	var v11062 int32
	_ = v11062
	var v11064 int32
	_ = v11064
	var v11066 int32
	_ = v11066
	var v11068 int32
	_ = v11068
	var v11073 int64
	_ = v11073
	var v11074 int64
	_ = v11074
	var v11077 int32
	_ = v11077
	var v11079 int32
	_ = v11079
	var v11080 int64
	_ = v11080
	var v11081 int64
	_ = v11081
	var v11084 int64
	_ = v11084
	var v11085 int64
	_ = v11085
	var v11091 int32
	_ = v11091
	var v11093 int64
	_ = v11093
	var v11101 int32
	_ = v11101
	var v11114 int64
	_ = v11114
	var v11121 int32
	_ = v11121
	var v11123 int64
	_ = v11123
	var v11125 int64
	_ = v11125
	var v11128 int32
	_ = v11128
	var v11129 int64
	_ = v11129
	var v11135 int64
	_ = v11135
	var v11154 int64
	_ = v11154
	var v11162 int64
	_ = v11162
	var v11168 int32
	_ = v11168
	var v11171 int32
	_ = v11171
	var v11175 int64
	_ = v11175
	var v11176 int32
	_ = v11176
	var v11179 int32
	_ = v11179
	var v11180 int64
	_ = v11180
	var v11182 int32
	_ = v11182
	var v11183 int32
	_ = v11183
	var v11186 int32
	_ = v11186
	var v11190 int32
	_ = v11190
	var v11191 int32
	_ = v11191
	var v11192 int32
	_ = v11192
	var v11196 int64
	_ = v11196
	var v11197 int64
	_ = v11197
	var v11200 int64
	_ = v11200
	var v11202 int32
	_ = v11202
	var v11203 int64
	_ = v11203
	var v11206 int32
	_ = v11206
	var v11210 int64
	_ = v11210
	var v11217 int64
	_ = v11217
	var v11218 int32
	_ = v11218
	var v11219 int32
	_ = v11219
	var v11221 int64
	_ = v11221
	var v11223 int32
	_ = v11223
	var v11225 int64
	_ = v11225
	var v11227 int32
	_ = v11227
	var v11230 int32
	_ = v11230
	var v11234 int64
	_ = v11234
	var v11236 int32
	_ = v11236
	var v11248 int64
	_ = v11248
	var v11255 int32
	_ = v11255
	var v11256 int32
	_ = v11256
	var v11259 int32
	_ = v11259
	var v11260 int32
	_ = v11260
	var v11263 int32
	_ = v11263
	var v11265 int32
	_ = v11265
	var v11272 int32
	_ = v11272
	var v11274 int64
	_ = v11274
	var v11276 int32
	_ = v11276
	var v11280 int32
	_ = v11280
	var v11284 int32
	_ = v11284
	var v11285 int32
	_ = v11285
	var v11287 int32
	_ = v11287
	var v11288 int64
	_ = v11288
	var v11290 int64
	_ = v11290
	var v11328 int64
	_ = v11328
	var v11337 int64
	_ = v11337
	var v11379 int32
	_ = v11379
	var v11383 int32
	_ = v11383
	var v11385 int32
	_ = v11385
	var v11389 int32
	_ = v11389
	var v11390 int32
	_ = v11390
	var v11392 int32
	_ = v11392
	var v11395 int32
	_ = v11395
	var v11396 int64
	_ = v11396
	var v11397 int32
	_ = v11397
	var v11400 int32
	_ = v11400
	var v11401 int32
	_ = v11401
	var v11405 int64
	_ = v11405
	var v11408 int32
	_ = v11408
	var v11409 int32
	_ = v11409
	var v11412 int32
	_ = v11412
	var v11414 int32
	_ = v11414
	var v11415 int32
	_ = v11415
	var v11417 int32
	_ = v11417
	var v11422 int32
	_ = v11422
	var v11423 int32
	_ = v11423
	var v11425 int32
	_ = v11425
	var v11426 int32
	_ = v11426
	var v11427 int32
	_ = v11427
	var v11430 int32
	_ = v11430
	var v11432 int32
	_ = v11432
	var v11433 int32
	_ = v11433
	var v11434 int32
	_ = v11434
	var v11435 int32
	_ = v11435
	var v11436 int32
	_ = v11436
	var v11443 int32
	_ = v11443
	var v11446 int32
	_ = v11446
	var v11447 int32
	_ = v11447
	var v11450 int32
	_ = v11450
	var v11464 int32
	_ = v11464
	var v11466 int32
	_ = v11466
	var v11468 int32
	_ = v11468
	var v11476 int32
	_ = v11476
	var v11485 int32
	_ = v11485
	var v11486 int32
	_ = v11486
	var v11488 int32
	_ = v11488
	var v11497 int32
	_ = v11497
	var v11501 int32
	_ = v11501
	var v11503 int32
	_ = v11503
	var v11506 int32
	_ = v11506
	var v11510 int32
	_ = v11510
	var v11511 int32
	_ = v11511
	var v11513 int32
	_ = v11513
	var v11515 int32
	_ = v11515
	var v11516 int32
	_ = v11516
	var v11518 int32
	_ = v11518
	var v11519 int32
	_ = v11519
	var v11520 int64
	_ = v11520
	var v11524 int32
	_ = v11524
	var v11525 int32
	_ = v11525
	var v11526 int32
	_ = v11526
	var v11528 int32
	_ = v11528
	var v11529 int64
	_ = v11529
	var v11531 int64
	_ = v11531
	var v11532 int32
	_ = v11532
	var v11534 int32
	_ = v11534
	var v11535 int32
	_ = v11535
	var v11537 int32
	_ = v11537
	var v11538 int32
	_ = v11538
	var v11541 int32
	_ = v11541
	var v11543 int32
	_ = v11543
	var v11544 int32
	_ = v11544
	var v11546 int32
	_ = v11546
	var v11550 int32
	_ = v11550
	var v11554 int32
	_ = v11554
	var v11555 int32
	_ = v11555
	var v11560 int32
	_ = v11560
	var v11561 int32
	_ = v11561
	var v11562 int32
	_ = v11562
	var v11564 int32
	_ = v11564
	var v11565 int32
	_ = v11565
	var v11569 int32
	_ = v11569
	var v11571 int32
	_ = v11571
	var v11572 int32
	_ = v11572
	var v11579 int32
	_ = v11579
	var v11588 int32
	_ = v11588
	var v11590 int32
	_ = v11590
	var v11592 int32
	_ = v11592
	var v11604 int32
	_ = v11604
	var v11613 int32
	_ = v11613
	var v11614 int32
	_ = v11614
	var v11616 int32
	_ = v11616
	var v11619 int32
	_ = v11619
	var v11621 int32
	_ = v11621
	var v11623 int64
	_ = v11623
	var v11625 int64
	_ = v11625
	var v11627 int64
	_ = v11627
	var v11629 int64
	_ = v11629
	var v11635 int64
	_ = v11635
	var v11638 int32
	_ = v11638
	var v11639 int32
	_ = v11639
	var v11640 int64
	_ = v11640
	var v11644 int32
	_ = v11644
	var v11646 int32
	_ = v11646
	var v11647 int32
	_ = v11647
	var v11652 int32
	_ = v11652
	var v11653 int32
	_ = v11653
	var v11655 int32
	_ = v11655
	var v11658 int32
	_ = v11658
	var v11664 int32
	_ = v11664
	var v11666 int32
	_ = v11666
	var v11667 int32
	_ = v11667
	var v11671 int32
	_ = v11671
	var v11672 int32
	_ = v11672
	var v11696 int32
	_ = v11696
	var v11700 int32
	_ = v11700
	var v11701 int32
	_ = v11701
	var v11703 int32
	_ = v11703
	var v11706 int32
	_ = v11706
	var v11714 int32
	_ = v11714
	var v11718 int32
	_ = v11718
	var v11719 int32
	_ = v11719
	var v11721 int32
	_ = v11721
	var v11722 int32
	_ = v11722
	var v11725 int32
	_ = v11725
	var v11729 int32
	_ = v11729
	var v11731 int32
	_ = v11731
	var v11735 int32
	_ = v11735
	var v11736 int32
	_ = v11736
	var v11738 int32
	_ = v11738
	var v11741 int32
	_ = v11741
	var v11745 int32
	_ = v11745
	var v11746 int32
	_ = v11746
	var v11748 int32
	_ = v11748
	var v11749 int32
	_ = v11749
	var v11764 int32
	_ = v11764
	var v11776 int32
	_ = v11776
	var v11791 int32
	_ = v11791
	var v11792 int64
	_ = v11792
	var v11793 int64
	_ = v11793
	var v11794 int32
	_ = v11794
	var v11797 int32
	_ = v11797
	var v11798 int32
	_ = v11798
	var v11801 int32
	_ = v11801
	var v11802 int32
	_ = v11802
	var v11803 int64
	_ = v11803
	var v11806 int64
	_ = v11806
	var v11810 int32
	_ = v11810
	var v11815 int32
	_ = v11815
	var v11817 int32
	_ = v11817
	var v11818 int32
	_ = v11818
	var v11819 int32
	_ = v11819
	var v11820 int32
	_ = v11820
	var v11821 int32
	_ = v11821
	var v11822 int32
	_ = v11822
	var v11823 int32
	_ = v11823
	var v11824 int64
	_ = v11824
	var v11826 int32
	_ = v11826
	var v11827 int64
	_ = v11827
	var v11828 int32
	_ = v11828
	var v11829 int32
	_ = v11829
	var v11831 int32
	_ = v11831
	var v11833 int32
	_ = v11833
	var v11837 int32
	_ = v11837
	var v11842 int32
	_ = v11842
	var v11845 int32
	_ = v11845
	var v11859 int32
	_ = v11859
	var v11869 int32
	_ = v11869
	var v11870 int32
	_ = v11870
	var v11871 int32
	_ = v11871
	var v11874 int32
	_ = v11874
	var v11875 int32
	_ = v11875
	var v11878 int32
	_ = v11878
	var v11883 int32
	_ = v11883
	var v11885 int32
	_ = v11885
	var v11890 int32
	_ = v11890
	var v11892 int32
	_ = v11892
	var v11895 int32
	_ = v11895
	var v11896 int32
	_ = v11896
	var v11897 int32
	_ = v11897
	var v11899 int32
	_ = v11899
	var v11904 int32
	_ = v11904
	var v11906 int32
	_ = v11906
	var v11910 int32
	_ = v11910
	var v11923 int32
	_ = v11923
	var v11947 int32
	_ = v11947
	var v11949 int32
	_ = v11949
	var v11952 int32
	_ = v11952
	var v11953 int32
	_ = v11953
	var v11954 int32
	_ = v11954
	var v11956 int32
	_ = v11956
	var v11957 int32
	_ = v11957
	var v11964 int32
	_ = v11964
	var v11967 int32
	_ = v11967
	var v11968 int32
	_ = v11968
	var v11970 int32
	_ = v11970
	var v11972 int32
	_ = v11972
	var v11976 int32
	_ = v11976
	var v11977 int32
	_ = v11977
	var v11979 int32
	_ = v11979
	var v11983 int32
	_ = v11983
	var v11987 int32
	_ = v11987
	var v11992 int32
	_ = v11992
	var v11994 int32
	_ = v11994
	var v11998 int32
	_ = v11998
	var v11999 int32
	_ = v11999
	var v12037 int32
	_ = v12037
	var v12039 int32
	_ = v12039
	var v12040 int32
	_ = v12040
	var v12079 int32
	_ = v12079
	var v12083 int32
	_ = v12083
	var v12087 int32
	_ = v12087
	var v12089 int32
	_ = v12089
	var v12092 int32
	_ = v12092
	var v12097 int32
	_ = v12097
	var v12098 int32
	_ = v12098
	var v12099 int64
	_ = v12099
	var v12100 int32
	_ = v12100
	var v12101 int64
	_ = v12101
	var v12104 int32
	_ = v12104
	var v12105 int32
	_ = v12105
	var v12106 int32
	_ = v12106
	var v12108 int32
	_ = v12108
	var v12109 int32
	_ = v12109
	var v12114 int32
	_ = v12114
	var v12115 int64
	_ = v12115
	var v12120 int32
	_ = v12120
	var v12123 int32
	_ = v12123
	var v12128 int32
	_ = v12128
	var v12129 int32
	_ = v12129
	var v12131 int32
	_ = v12131
	var v12133 int32
	_ = v12133
	var v12135 int32
	_ = v12135
	var v12137 int32
	_ = v12137
	var v12138 int32
	_ = v12138
	var v12140 int32
	_ = v12140
	var v12141 int32
	_ = v12141
	var v12143 int32
	_ = v12143
	var v12145 int32
	_ = v12145
	var v12147 int32
	_ = v12147
	var v12153 int32
	_ = v12153
	var v12154 int32
	_ = v12154
	var v12155 int32
	_ = v12155
	var v12159 int32
	_ = v12159
	var v12160 int32
	_ = v12160
	var v12161 int32
	_ = v12161
	var v12163 int32
	_ = v12163
	var v12167 int32
	_ = v12167
	var v12181 int32
	_ = v12181
	var v12186 int32
	_ = v12186
	var v12187 int32
	_ = v12187
	var v12189 int32
	_ = v12189
	var v12200 int32
	_ = v12200
	var v12207 int64
	_ = v12207
	var v12210 int32
	_ = v12210
	var v12213 int32
	_ = v12213
	var v12214 int64
	_ = v12214
	var v12215 int64
	_ = v12215
	var v12216 int32
	_ = v12216
	var v12220 int64
	_ = v12220
	var v12221 int64
	_ = v12221
	var v12223 int64
	_ = v12223
	var v12229 int64
	_ = v12229
	var v12230 int64
	_ = v12230
	var v12231 int64
	_ = v12231
	var v12241 int64
	_ = v12241
	var v12243 int64
	_ = v12243
	var v12247 int64
	_ = v12247
	var v12249 int32
	_ = v12249
	var v12251 int32
	_ = v12251
	var v12256 int32
	_ = v12256
	var v12261 int32
	_ = v12261
	var v12263 int32
	_ = v12263
	var v12265 int32
	_ = v12265
	var v12269 int32
	_ = v12269
	var v12274 int32
	_ = v12274
	var v12275 int32
	_ = v12275
	var v12278 int32
	_ = v12278
	var v12280 int32
	_ = v12280
	var v12284 int32
	_ = v12284
	var v12286 int32
	_ = v12286
	var v12287 int32
	_ = v12287
	var v12288 int32
	_ = v12288
	var v12290 int32
	_ = v12290
	var v12293 int32
	_ = v12293
	var v12295 int32
	_ = v12295
	var v12300 int32
	_ = v12300
	var v12301 int32
	_ = v12301
	var v12302 int32
	_ = v12302
	var v12305 int64
	_ = v12305
	var v12306 int64
	_ = v12306
	var v12320 int32
	_ = v12320
	var v12323 int64
	_ = v12323
	var v12324 int32
	_ = v12324
	var v12326 int64
	_ = v12326
	var v12329 int32
	_ = v12329
	var v12330 int32
	_ = v12330
	var v12332 int32
	_ = v12332
	var v12342 int32
	_ = v12342
	var v12345 int32
	_ = v12345
	var v12346 int32
	_ = v12346
	var v12350 int32
	_ = v12350
	var v12354 int32
	_ = v12354
	var v12359 int32
	_ = v12359
	var v12360 int32
	_ = v12360
	var v12364 int32
	_ = v12364
	var v12369 int32
	_ = v12369
	var v12370 int32
	_ = v12370
	var v12372 int32
	_ = v12372
	var v12379 int32
	_ = v12379
	var v12380 int32
	_ = v12380
	var v12381 int32
	_ = v12381
	var v12384 int64
	_ = v12384
	var v12385 int64
	_ = v12385
	var v12396 int32
	_ = v12396
	var v12399 int32
	_ = v12399
	var v12401 int32
	_ = v12401
	var v12402 int32
	_ = v12402
	var v12404 int32
	_ = v12404
	var v12407 int32
	_ = v12407
	var v12408 int32
	_ = v12408
	var v12409 int32
	_ = v12409
	var v12411 int32
	_ = v12411
	var v12416 int32
	_ = v12416
	var v12421 int32
	_ = v12421
	var v12424 int64
	_ = v12424
	var v12425 int32
	_ = v12425
	var v12427 int32
	_ = v12427
	var v12429 int32
	_ = v12429
	var v12433 int32
	_ = v12433
	var v12434 int32
	_ = v12434
	var v12436 int32
	_ = v12436
	var v12438 int32
	_ = v12438
	var v12441 int32
	_ = v12441
	var v12443 int32
	_ = v12443
	var v12445 int32
	_ = v12445
	var v12449 int32
	_ = v12449
	var v12450 int32
	_ = v12450
	var v12452 int32
	_ = v12452
	var v12458 int32
	_ = v12458
	var v12459 int32
	_ = v12459
	var v12463 int32
	_ = v12463
	var v12465 int32
	_ = v12465
	var v12466 int32
	_ = v12466
	var v12469 int32
	_ = v12469
	var v12470 int32
	_ = v12470
	var v12473 int32
	_ = v12473
	var v12474 int32
	_ = v12474
	var v12477 int32
	_ = v12477
	var v12478 int32
	_ = v12478
	var v12481 int32
	_ = v12481
	var v12482 int32
	_ = v12482
	var v12485 int32
	_ = v12485
	var v12486 int32
	_ = v12486
	var v12489 int32
	_ = v12489
	var v12490 int32
	_ = v12490
	var v12493 int32
	_ = v12493
	var v12494 int32
	_ = v12494
	var v12497 int32
	_ = v12497
	var v12504 int32
	_ = v12504
	var v12507 int32
	_ = v12507
	var v12510 int32
	_ = v12510
	var v12513 int32
	_ = v12513
	var v12516 int32
	_ = v12516
	var v12519 int32
	_ = v12519
	var v12522 int32
	_ = v12522
	var v12525 int32
	_ = v12525
	var v12530 int32
	_ = v12530
	var v12533 int64
	_ = v12533
	var v12534 int32
	_ = v12534
	var v12536 int32
	_ = v12536
	var v12538 int32
	_ = v12538
	var v12542 int32
	_ = v12542
	var v12543 int32
	_ = v12543
	var v12545 int32
	_ = v12545
	var v12547 int32
	_ = v12547
	var v12550 int32
	_ = v12550
	var v12553 int32
	_ = v12553
	var v12556 int32
	_ = v12556
	var v12559 int32
	_ = v12559
	var v12562 int32
	_ = v12562
	var v12565 int32
	_ = v12565
	var v12568 int32
	_ = v12568
	var v12571 int32
	_ = v12571
	var v12573 int32
	_ = v12573
	var v12575 int32
	_ = v12575
	var v12579 int32
	_ = v12579
	var v12582 int32
	_ = v12582
	var v12586 int32
	_ = v12586
	var v12589 int32
	_ = v12589
	var v12596 int32
	_ = v12596
	var v12598 int32
	_ = v12598
	var v12600 int32
	_ = v12600
	var v12608 int32
	_ = v12608
	var v12614 int64
	_ = v12614
	var v12615 int64
	_ = v12615
	var v12617 int64
	_ = v12617
	var v12618 int64
	_ = v12618
	var v12621 int64
	_ = v12621
	var v12624 int32
	_ = v12624
	var v12629 int32
	_ = v12629
	var v12630 int32
	_ = v12630
	var v12631 int32
	_ = v12631
	var v12633 int32
	_ = v12633
	var v12639 int32
	_ = v12639
	var v12644 int32
	_ = v12644
	var v12645 int32
	_ = v12645
	var v12646 int32
	_ = v12646
	var v12648 int32
	_ = v12648
	var v12651 int32
	_ = v12651
	var v12661 int32
	_ = v12661
	var v12662 int32
	_ = v12662
	var v12665 int32
	_ = v12665
	var v12673 int32
	_ = v12673
	var v12674 int32
	_ = v12674
	var v12677 int32
	_ = v12677
	var v12680 int32
	_ = v12680
	var v12685 int32
	_ = v12685
	var v12689 int32
	_ = v12689
	var v12693 int64
	_ = v12693
	var v12694 int64
	_ = v12694
	var v12695 int64
	_ = v12695
	var v12698 int64
	_ = v12698
	var v12701 int32
	_ = v12701
	var v12706 int32
	_ = v12706
	var v12707 int32
	_ = v12707
	var v12712 int32
	_ = v12712
	var v12717 int32
	_ = v12717
	var v12718 int32
	_ = v12718
	var v12721 int32
	_ = v12721
	var v12726 int32
	_ = v12726
	var v12727 int32
	_ = v12727
	var v12729 int32
	_ = v12729
	var v12731 int32
	_ = v12731
	var v12732 int32
	_ = v12732
	var v12734 int32
	_ = v12734
	var v12745 int32
	_ = v12745
	var v12749 int32
	_ = v12749
	var v12753 int32
	_ = v12753
	var v12754 int32
	_ = v12754
	var v12755 int32
	_ = v12755
	var v12756 int32
	_ = v12756
	var v12757 int32
	_ = v12757
	var v12762 int32
	_ = v12762
	var v12768 int32
	_ = v12768
	var v12774 int32
	_ = v12774
	var v12775 int32
	_ = v12775
	var v12777 int32
	_ = v12777
	var v12781 int32
	_ = v12781
	var v12783 int32
	_ = v12783
	var v12785 int32
	_ = v12785
	var v12786 int32
	_ = v12786
	var v12788 int32
	_ = v12788
	var v12791 int32
	_ = v12791
	var v12795 int32
	_ = v12795
	var v12796 int32
	_ = v12796
	var v12798 int32
	_ = v12798
	var v12801 int32
	_ = v12801
	var v12802 int32
	_ = v12802
	var v12806 int32
	_ = v12806
	var v12808 int32
	_ = v12808
	var v12809 int32
	_ = v12809
	var v12813 int32
	_ = v12813
	var v12817 int32
	_ = v12817
	var v12818 int32
	_ = v12818
	var v12822 int32
	_ = v12822
	var v12827 int32
	_ = v12827
	var v12829 int32
	_ = v12829
	var v12830 int32
	_ = v12830
	var v12833 int32
	_ = v12833
	var v12837 int32
	_ = v12837
	var v12840 int32
	_ = v12840
	var v12845 int32
	_ = v12845
	var v12848 int64
	_ = v12848
	var v12849 int32
	_ = v12849
	var v12851 int32
	_ = v12851
	var v12853 int32
	_ = v12853
	var v12857 int32
	_ = v12857
	var v12859 int32
	_ = v12859
	var v12863 int64
	_ = v12863
	var v12864 int32
	_ = v12864
	var v12866 int32
	_ = v12866
	var v12870 int32
	_ = v12870
	var v12872 int32
	_ = v12872
	var v12873 int32
	_ = v12873
	var v12875 int32
	_ = v12875
	var v12877 int32
	_ = v12877
	var v12879 int32
	_ = v12879
	var v12883 int32
	_ = v12883
	var v12884 int32
	_ = v12884
	var v12886 int32
	_ = v12886
	var v12890 int32
	_ = v12890
	var v12893 int32
	_ = v12893
	var v12898 int32
	_ = v12898
	var v12900 int32
	_ = v12900
	var v12903 int32
	_ = v12903
	var v12907 int32
	_ = v12907
	var v12909 int32
	_ = v12909
	var v12911 int32
	_ = v12911
	var v12913 int32
	_ = v12913
	var v12917 int32
	_ = v12917
	var v12921 int32
	_ = v12921
	var v12922 int32
	_ = v12922
	var v12925 int32
	_ = v12925
	var v12935 int32
	_ = v12935
	var v12939 int32
	_ = v12939
	var v12943 int32
	_ = v12943
	var v12945 int32
	_ = v12945
	var v12947 int32
	_ = v12947
	var v12948 int32
	_ = v12948
	var v12951 int32
	_ = v12951
	var v12954 int32
	_ = v12954
	var v12959 int32
	_ = v12959
	var v12962 int32
	_ = v12962
	var v12966 int32
	_ = v12966
	var v12971 int32
	_ = v12971
	var v12975 int32
	_ = v12975
	var v12978 int32
	_ = v12978
	var v12982 int32
	_ = v12982
	var v12987 int32
	_ = v12987
	var v12991 int32
	_ = v12991
	var v12993 int32
	_ = v12993
	var v13000 int32
	_ = v13000
	var v13005 int32
	_ = v13005
	var v13009 int32
	_ = v13009
	var v13011 int32
	_ = v13011
	var v13019 int32
	_ = v13019
	var v13024 int32
	_ = v13024
	var v13028 int32
	_ = v13028
	var v13036 int32
	_ = v13036
	var v13041 int32
	_ = v13041
	var v13045 int32
	_ = v13045
	var v13048 int32
	_ = v13048
	var v13052 int32
	_ = v13052
	var v13056 int32
	_ = v13056
	var v13061 int32
	_ = v13061
	var v13065 int32
	_ = v13065
	var v13067 int32
	_ = v13067
	var v13075 int32
	_ = v13075
	var v13080 int32
	_ = v13080
	var v13084 int32
	_ = v13084
	var v13086 int32
	_ = v13086
	var v13094 int32
	_ = v13094
	var v13099 int32
	_ = v13099
	var v13101 int32
	_ = v13101
	var v13109 int32
	_ = v13109
	var v13114 int32
	_ = v13114
	var v13115 int32
	_ = v13115
	var v13116 int32
	_ = v13116
	var v13118 int32
	_ = v13118
	var v13119 int32
	_ = v13119
	var v13122 int32
	_ = v13122
	var v13127 int32
	_ = v13127
	var v13129 int32
	_ = v13129
	var v13135 int32
	_ = v13135
	var v13140 int32
	_ = v13140
	var v13144 int32
	_ = v13144
	var v13146 int32
	_ = v13146
	var v13154 int32
	_ = v13154
	var v13159 int32
	_ = v13159
	var v13163 int32
	_ = v13163
	var v13165 int32
	_ = v13165
	var v13173 int32
	_ = v13173
	var v13178 int32
	_ = v13178
	var v13180 int32
	_ = v13180
	var v13182 int32
	_ = v13182
	var v13184 int32
	_ = v13184
	var v13186 int32
	_ = v13186
	var v13192 int32
	_ = v13192
	var v13194 int32
	_ = v13194
	var v13200 int32
	_ = v13200
	var v13205 int32
	_ = v13205
	var v13211 int32
	_ = v13211
	var v13215 int32
	_ = v13215
	var v13220 int32
	_ = v13220
	var v13224 int32
	_ = v13224
	var v13227 int64
	_ = v13227
	var v13233 int32
	_ = v13233
	var v13238 int32
	_ = v13238
	var v13242 int32
	_ = v13242
	var v13245 int64
	_ = v13245
	var v13251 int32
	_ = v13251
	var v13256 int32
	_ = v13256
	var v13260 int32
	_ = v13260
	var v13262 int64
	_ = v13262
	var v13265 int64
	_ = v13265
	var v13271 int32
	_ = v13271
	var v13276 int32
	_ = v13276
	var v13281 int32
	_ = v13281
	var v13285 int32
	_ = v13285
	var v13290 int32
	_ = v13290
	v1 = int32(0)
	v37 = m.G0
	v41 = (v37 - int32(_a_F_StartupXLOG_0)) & int32(-4096)
	m.G0 = v41
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[1])) = v45
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
	if base.Ui64(int64(23)) < base.Ui64(v49&int64(8184)) {
		goto L18
	} else {
		goto L19
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13281 = m.ExcPending
	if v13281 != 0 {
		goto L32
	} else {
		goto L2765
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13260 = m.ExcPending
	if v13260 != 0 {
		goto L32
	} else {
		goto L2762
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13242 = m.ExcPending
	if v13242 != 0 {
		goto L32
	} else {
		goto L2759
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13224 = m.ExcPending
	if v13224 != 0 {
		goto L32
	} else {
		goto L2756
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13211 = m.ExcPending
	if v13211 != 0 {
		goto L32
	} else {
		goto L2753
	}
L6:
	;
	v13180 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v13182 = v41 + int32(_a_F_StartupXLOG_1)
	v13184 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	F_XLogFileName(m, v13182, v10145, v10171, v13184)
	mBase = m.M
	v13186 = m.ExcPending
	if v13186 != 0 {
		goto L32
	} else {
		goto L2748
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13163 = m.ExcPending
	if v13163 != 0 {
		goto L32
	} else {
		goto L2744
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13144 = m.ExcPending
	if v13144 != 0 {
		goto L32
	} else {
		goto L2740
	}
L9:
	;
	v13115 = int32(_a_F_StartupXLOG_2)
	v13116 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v13118 = v41 + int32(_a_F_StartupXLOG_3)
	v13119 = F_unlink(m, v13118)
	mBase = m.M
	if v13116 != 0 {
		goto L2733
	} else {
		goto L2734
	}
L10:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13101 = m.ExcPending
	if v13101 != 0 {
		goto L32
	} else {
		goto L2730
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13084 = m.ExcPending
	if v13084 != 0 {
		goto L32
	} else {
		goto L2726
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13065 = m.ExcPending
	if v13065 != 0 {
		goto L32
	} else {
		goto L2722
	}
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v13045 = m.ExcPending
	if v13045 != 0 {
		goto L32
	} else {
		goto L2717
	}
L14:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v13028 = m.ExcPending
	if v13028 != 0 {
		goto L32
	} else {
		goto L2714
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v13009 = m.ExcPending
	if v13009 != 0 {
		goto L32
	} else {
		goto L2710
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12991 = m.ExcPending
	if v12991 != 0 {
		goto L32
	} else {
		goto L2706
	}
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12975 = m.ExcPending
	if v12975 != 0 {
		goto L32
	} else {
		goto L2702
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	switch v55 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	case 5:
		goto L23
	default:
		goto L17
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12959 = m.ExcPending
	if v12959 != 0 {
		goto L32
	} else {
		goto L2698
	}
L21:
	;
	v252 = v41 + int32(_a_F_StartupXLOG_4)
	v255 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_5), v252, int32(0))
	mBase = m.M
	goto L66
L22:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), v245, int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L32
	} else {
		goto L65
	}
L23:
	;
	v217 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L32
	} else {
		goto L60
	}
L24:
	;
	v184 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L32
	} else {
		goto L54
	}
L25:
	;
	v151 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L32
	} else {
		goto L48
	}
L26:
	;
	v122 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L32
	} else {
		goto L43
	}
L27:
	;
	v93 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L32
	} else {
		goto L38
	}
L28:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[5])))
	if v61 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v62 = int32(15)
	goto L31
L30:
	;
	v62 = int32(18)
	goto L31
L31:
	;
	v64 = F_errstart(m, v62, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return
L33:
	;
	if v64 == int32(0) {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v70
	v73 = v41 + int32(3952)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v80 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v82 = F_pg_strftime(m, v73, int32(128), int32(_a_F_StartupXLOG_8), v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3856)) = v73
	F_errmsg(m, int32(_a_F_StartupXLOG_9), v41+int32(3856))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v245 = int32(_a_F_StartupXLOG_10)
	goto L22
L38:
	;
	if v93 == int32(0) {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v99
	v102 = v41 + int32(3952)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v109 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v111 = F_pg_strftime(m, v102, int32(128), int32(_a_F_StartupXLOG_8), v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3872)) = v102
	F_errmsg(m, int32(_a_F_StartupXLOG_11), v41+int32(3872))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v245 = int32(_a_F_StartupXLOG_12)
	goto L22
L43:
	;
	if v122 == int32(0) {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v128
	v131 = v41 + int32(3952)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v138 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v140 = F_pg_strftime(m, v131, int32(128), int32(_a_F_StartupXLOG_8), v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3888)) = v131
	F_errmsg(m, int32(_a_F_StartupXLOG_13), v41+int32(3888))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v245 = int32(_a_F_StartupXLOG_14)
	goto L22
L48:
	;
	if v151 == int32(0) {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v157
	v160 = v41 + int32(3952)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v167 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v169 = F_pg_strftime(m, v160, int32(128), int32(_a_F_StartupXLOG_8), v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3904)) = v160
	F_errmsg(m, int32(_a_F_StartupXLOG_15), v41+int32(3904))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L32
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_16), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v245 = int32(_a_F_StartupXLOG_17)
	goto L22
L54:
	;
	if v184 == int32(0) {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v189)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v190
	v193 = v41 + int32(3952)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v200 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L32
	} else {
		goto L56
	}
L56:
	;
	v202 = F_pg_strftime(m, v193, int32(128), int32(_a_F_StartupXLOG_8), v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L32
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3920)) = v193
	F_errmsg(m, int32(_a_F_StartupXLOG_18), v41+int32(3920))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L32
	} else {
		goto L58
	}
L58:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_19), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L32
	} else {
		goto L59
	}
L59:
	;
	v245 = int32(_a_F_StartupXLOG_20)
	goto L22
L60:
	;
	if v217 == int32(0) {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v222)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v223
	v226 = v41 + int32(3952)
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v233 = F_pg_localtime(m, v41+int32(_a_F_StartupXLOG_1), v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L32
	} else {
		goto L62
	}
L62:
	;
	v235 = F_pg_strftime(m, v226, int32(128), int32(_a_F_StartupXLOG_8), v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3936)) = v226
	F_errmsg(m, int32(_a_F_StartupXLOG_21), v41+int32(3936))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	v245 = int32(_a_F_StartupXLOG_22)
	goto L22
L65:
	;
	goto L21
L66:
	;
	if v255 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[8])))
	if v256&int32(_a_F_StartupXLOG_23) != int32(_a_F_StartupXLOG_0) {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	v262 = v41 + int32(_a_F_StartupXLOG_1)
	v266 = F_pg_snprintf(m, v262, int32(1024), int32(_a_F_StartupXLOG_24), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	v270 = F___fstatat(m, int32(-100), v262, v252, int32(0))
	mBase = m.M
	goto L71
L70:
	;
	v320 = v41 + int32(_a_F_StartupXLOG_1)
	v324 = F_pg_snprintf(m, v320, int32(1024), int32(_a_F_StartupXLOG_25), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L32
	} else {
		goto L88
	}
L71:
	;
	if v270 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[8])))
	if v273&int32(_a_F_StartupXLOG_23) == int32(_a_F_StartupXLOG_0) {
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v297 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L32
	} else {
		goto L80
	}
L75:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L32
	} else {
		goto L76
	}
L76:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L32
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3792)) = v262
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v41+int32(3792))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L32
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_27), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	if v297 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3824)) = v41 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_29), v41+int32(3824))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L32
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[9]))
	v316 = F_mkdir(m, v41+int32(_a_F_StartupXLOG_1), v315)
	mBase = m.M
	goto L86
L84:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_30), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L32
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	if v316 < int32(0) {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	goto L70
L88:
	;
	v330 = F___fstatat(m, int32(-100), v320, v41+int32(_a_F_StartupXLOG_4), int32(0))
	mBase = m.M
	goto L90
L89:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[10]))
	if v378 != 0 {
		goto L106
	} else {
		goto L107
	}
L90:
	;
	if v330 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[8])))
	if v333&int32(_a_F_StartupXLOG_23) == int32(_a_F_StartupXLOG_0) {
		goto L89
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v355 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L32
	} else {
		goto L98
	}
L94:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L32
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3744)) = v320
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v41+int32(3744))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L32
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_31), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	if v355 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3776)) = v41 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_29), v41+int32(3776))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L32
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[9]))
	v374 = F_mkdir(m, v41+int32(_a_F_StartupXLOG_1), v373)
	mBase = m.M
	goto L104
L102:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_32), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L32
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	if v374 < int32(0) {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	goto L89
L106:
	;
	F_RegisterTimeout(m, int32(12), int32(429))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L32
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+16))
	v387 = v385 - int32(3)
	if base.Ui32(v387) <= base.Ui32(int32(-3)) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v392 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L32
	} else {
		goto L113
	}
L111:
	;
	v672 = v384
	goto L112
L112:
	;
	v706 = m.G0
	v708 = v706 - int32(2176)
	m.G0 = v708
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v672)+16))
	v712 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])) = uint8(v712)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v672)+152))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v672)+48))
	if base.Ui32(v716) < base.Ui32(v715) {
		goto L178
	} else {
		goto L179
	}
L113:
	;
	if v392 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_33), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L32
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v404 = F_AllocateDir(m, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L32
	} else {
		goto L119
	}
L117:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3898), int32(_a_F_StartupXLOG_34))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L32
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v407 = F_ReadDir(m, v404, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L32
	} else {
		goto L120
	}
L120:
	;
	if v407 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v417 = v407
	goto L124
L122:
	;
	goto L123
L123:
	;
	F_FreeDir(m, v404)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L32
	} else {
		goto L148
	}
L124:
	;
	v446 = v417 + int32(19)
	v447 = int32(_a_F_StartupXLOG_35)
	goto L129
L125:
	;
	goto L123
L126:
	;
	v523 = F_ReadDir(m, v404, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L32
	} else {
		goto L146
	}
L127:
	;
	if v485-v486 != 0 {
		goto L126
	} else {
		goto L140
	}
L129:
	;
	goto L130
L130:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v454 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v455 = v446
	v456 = v447
	v457 = int32(9)
	v458 = v454
	goto L135
L132:
	;
	v481 = v447
	v485 = int32(0)
	goto L133
L133:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	goto L127
L134:
	;
	v481 = v476
	v485 = v478
	goto L133
L135:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if base.B2i32(v458 != v460)|base.B2i32(v460 == int32(0)) != 0 {
		v476 = v456
		v478 = v458
		goto L134
	} else {
		goto L137
	}
L136:
	;
	v476 = v470
	v478 = int32(0)
	goto L134
L137:
	;
	v466 = v457 - int32(1)
	if v466 == int32(0) {
		v476 = v456
		v478 = v458
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v469 = int32(1)
	v470 = v456 + v469
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+1)))
	if v471 != 0 {
		v455 = v455 + v469
		v456 = v470
		v457 = v466
		v458 = v471
		goto L135
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3728)) = v446
	v496 = v41 + int32(_a_F_StartupXLOG_1)
	v501 = F_pg_snprintf(m, v496, int32(1024), int32(_a_F_StartupXLOG_36), v41+int32(3728))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L32
	} else {
		goto L141
	}
L141:
	;
	v503 = F_unlink(m, v496)
	mBase = m.M
	v506 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L32
	} else {
		goto L142
	}
L142:
	;
	if v506 == int32(0) {
		goto L126
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3712)) = v496
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_37), v41+int32(3712))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L32
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3910), int32(_a_F_StartupXLOG_34))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L32
	} else {
		goto L145
	}
L145:
	;
	goto L126
L146:
	;
	if v523 != 0 {
		v417 = v523
		goto L124
	} else {
		goto L147
	}
L147:
	;
	goto L125
L148:
	;
	v563 = m.G0
	v565 = v563 - int32(112)
	m.G0 = v565
	v568 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v568 == int32(1) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v576 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_5), v565+int32(16), int32(256))
	mBase = m.M
	goto L154
L150:
	;
	goto L151
L151:
	;
	m.G0 = v565 + int32(112)
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v672 = v663
	goto L112
L152:
	;
	F_walkdir(m, v646, int32(1182), int32(0), int32(15))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L32
	} else {
		goto L176
	}
L153:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1181), int32(1), int32(14))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L32
	} else {
		goto L174
	}
L154:
	;
	if v576 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v581 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L32
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v565)+20))
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L32
	} else {
		goto L167
	}
L158:
	;
	if v581 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L32
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L32
	} else {
		goto L165
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = int32(_a_F_StartupXLOG_5)
	F_errmsg(m, int32(_a_F_StartupXLOG_39), v565)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L32
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_40), int32(3620), int32(_a_F_StartupXLOG_41))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L32
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1181), int32(0), int32(14))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L32
	} else {
		goto L166
	}
L166:
	;
	goto L153
L167:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1181), int32(0), int32(14))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L32
	} else {
		goto L168
	}
L168:
	;
	if v603&int32(_a_F_StartupXLOG_23) != int32(_a_F_StartupXLOG_43) {
		goto L153
	} else {
		goto L169
	}
L169:
	;
	v616 = int32(_a_F_StartupXLOG_5)
	F_walkdir(m, v616, int32(1181), int32(0), int32(14))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L32
	} else {
		goto L170
	}
L170:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1181), int32(1), int32(14))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L32
	} else {
		goto L171
	}
L171:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L32
	} else {
		goto L172
	}
L172:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1182), int32(0), int32(15))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L32
	} else {
		goto L173
	}
L173:
	;
	v646 = v616
	goto L152
L174:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L32
	} else {
		goto L175
	}
L175:
	;
	v646 = int32(_a_F_StartupXLOG_42)
	goto L152
L176:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1182), int32(1), int32(15))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L32
	} else {
		goto L177
	}
L177:
	;
	goto L151
L178:
	;
	v718 = v715
	goto L180
L179:
	;
	v718 = v716
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13])) = v718
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[10]))
	if v721 != 0 {
		goto L189
	} else {
		goto L190
	}
L181:
	;
	v2388 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2388 != int32(1) {
		goto L533
	} else {
		goto L534
	}
L182:
	;
	v2352 = v2330
	v2353 = v2331
	v2364 = v1
	v2366 = v2334
	v2368 = v2336
	v2369 = v2337
	v2371 = v2338
	v2372 = v2339
	v2373 = v2340
	v2374 = v2341
	v2375 = v2342
	v2376 = v2343
	v2379 = v2344
	v2382 = v2347
	v2383 = v2348
	v2384 = v2349
	v2386 = v2350
	goto L181
L183:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v2308 == int32(44) {
		v2330 = v1339
		v2331 = v1448
		v2334 = v1335
		v2336 = v1331
		v2337 = v1338
		v2338 = v1330
		v2339 = v1336
		v2340 = v1320
		v2341 = v1333
		v2342 = v1334
		v2343 = v1328
		v2344 = v1321
		v2347 = v1332
		v2348 = v1329
		v2349 = v1340
		v2350 = v1337
		goto L182
	} else {
		goto L528
	}
L184:
	;
	v2094 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_44), v708+int32(1152), int32(0))
	mBase = m.M
	goto L478
L185:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L32
	} else {
		goto L473
	}
L186:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L32
	} else {
		goto L469
	}
L187:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v947 == int32(1) {
		goto L271
	} else {
		goto L272
	}
L188:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v851 != 0 {
		goto L235
	} else {
		goto L236
	}
L189:
	;
	v724 = v708 + int32(1152)
	v727 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_45), v724, int32(0))
	mBase = m.M
	goto L192
L190:
	;
	goto L191
L191:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v846 != int32(1) {
		goto L187
	} else {
		goto L233
	}
L192:
	;
	if v727 == int32(0) {
		goto L185
	} else {
		goto L193
	}
L193:
	;
	v731 = F_unlink(m, int32(_a_F_StartupXLOG_46))
	mBase = m.M
	v735 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_47), v724, int32(0))
	mBase = m.M
	goto L194
L194:
	;
	if v735 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v741 = F_BasicOpenFilePerm(m, int32(_a_F_StartupXLOG_47), int32(2), int32(384))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L32
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v772 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_48), v708+int32(1152), int32(0))
	mBase = m.M
	goto L209
L198:
	;
	if int32(0) <= v741 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v747 != int32(1) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	goto L201
L201:
	;
	v764 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])) = uint8(v764)
	goto L197
L202:
	;
	v762 = F_close(m, v741)
	mBase = m.M
	goto L201
L203:
	;
	goto L202
L204:
	;
	goto L205
L205:
	;
	v752 = F_fsync(m, v741)
	mBase = m.M
	if v752 != int32(-1) {
		goto L203
	} else {
		goto L207
	}
L206:
	;
	goto L203
L207:
	;
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v756 == int32(27) {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	if v772 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v778 = F_BasicOpenFilePerm(m, int32(_a_F_StartupXLOG_48), int32(2), int32(384))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L32
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v805 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])) = uint8(v805)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])) = uint8(v805)
	v811 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])))
	if v811 == v805 {
		goto L224
	} else {
		goto L225
	}
L213:
	;
	if int32(0) <= v778 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v784 != int32(1) {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	v801 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])) = uint8(v801)
	goto L212
L217:
	;
	v799 = F_close(m, v778)
	mBase = m.M
	goto L216
L218:
	;
	goto L217
L219:
	;
	goto L220
L220:
	;
	v789 = F_fsync(m, v778)
	mBase = m.M
	if v789 != int32(-1) {
		goto L218
	} else {
		goto L222
	}
L221:
	;
	goto L218
L222:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v793 == int32(27) {
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])))
	if v815 == int32(0) {
		goto L187
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v822 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])) = uint8(v822)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])) = uint8(v822)
	v828 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[18])))
	if v828 != 0 {
		goto L188
	} else {
		goto L228
	}
L227:
	;
	v819 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])) = uint8(v819)
	goto L188
L228:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L32
	} else {
		goto L229
	}
L229:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L32
	} else {
		goto L230
	}
L230:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_49), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L32
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1073), int32(_a_F_StartupXLOG_51))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L32
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	goto L188
L234:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[19]))
	if v886 != 0 {
		goto L253
	} else {
		goto L254
	}
L235:
	;
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[20]))
	if v853 != 0 {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	goto L237
L237:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v878 == int32(0) {
		goto L186
	} else {
		goto L251
	}
L238:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	if v854 != 0 {
		goto L234
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v856 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v856 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	if v857 != 0 {
		goto L234
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v860 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L32
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	if v860 == int32(0) {
		goto L234
	} else {
		goto L247
	}
L247:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_52), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L32
	} else {
		goto L248
	}
L248:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_53), int32(0))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L32
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1091), int32(_a_F_StartupXLOG_54))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L32
	} else {
		goto L250
	}
L250:
	;
	goto L234
L251:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	if v881 == int32(0) {
		goto L186
	} else {
		goto L252
	}
L252:
	;
	goto L234
L253:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v895 == int32(2) {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[23])))
	if v888&int32(1) != 0 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[19])) = int32(2)
	goto L253
L256:
	;
	v902 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_StartupXLOG[24])))
	v905 = F_DirectFunctionCall3Coll(m, int32(439), int32(0), v902, int64(0), int64(-1))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L32
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[25]))
	switch v910 - int32(1) {
	case 0:
		goto L261
	case 1:
		goto L262
	default:
		goto L187
	}
L259:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[26])) = v905
	goto L258
L260:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13])) = v943
	goto L187
L261:
	;
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	v940 = F_findNewestTimeLine(m, v939)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L32
	} else {
		goto L270
	}
L262:
	;
	v913 = int32(1)
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	if v915 == v913 {
		v943 = v913
		goto L260
	} else {
		goto L263
	}
L263:
	;
	v918 = F_existsTimeLineHistory(m, v915)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L32
	} else {
		goto L264
	}
L264:
	;
	if v918 != 0 {
		v943 = v915
		goto L260
	} else {
		goto L265
	}
L265:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L32
	} else {
		goto L266
	}
L266:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L32
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+848)) = v915
	F_errmsg(m, int32(_a_F_StartupXLOG_55), v708+int32(848))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L32
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1138), int32(_a_F_StartupXLOG_54))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L32
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	v943 = v940
	goto L260
L271:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	F_OwnLatch(m, v951+int32(4))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L32
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v957 = F_palloc0(m, int32(12))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L32
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+884)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+880)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+876)) = int32(440)
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v970 = F_XLogReaderAllocate(m, v967, v708+int32(876), v957)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L32
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29])) = v970
	if v970 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v973 = *(*int64)(unsafe.Add(mBase, uint32(v672)))
	*(*int64)(unsafe.Add(mBase, uint32(v970)+16)) = v973
	v976 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[30]))
	v977 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v970)+116)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v970)+104)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v970)+100)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v970)+112)) = v977
	v985 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v986 = m.G0
	v988 = v986 - int32(48)
	m.G0 = v988
	v991 = F_palloc0(m, int32(136))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L32
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L32
	} else {
		goto L464
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v991))) = v985
	*(*int64)(unsafe.Add(mBase, uint32(v988)+8)) = int64(171798691852)
	v999 = F_hash_create(m, int32(_a_F_StartupXLOG_56), int64(1024), v988, int32(40))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L32
	} else {
		goto L281
	}
L281:
	;
	v1002 = v991 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v991)+32)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v991)+28)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v991)+24)) = v999
	v1007 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1007)+56)) = int64(0)
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v991)+128)) = v1013 - int32(1)
	m.G0 = v988 + int32(48)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33])) = v991
	v1024 = F_palloc(m, int32(_a_F_StartupXLOG_57))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L32
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34])) = v1024
	v1029 = F_palloc(m, int32(_a_F_StartupXLOG_57))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L32
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = v1029
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = int64(0)
	v1036 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = v1036
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = uint8(v1036)
	v1043 = F_AllocateFile(m, int32(_a_F_StartupXLOG_58), int32(_a_F_StartupXLOG_59))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L32
	} else {
		goto L284
	}
L284:
	;
	if v1043 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v1048 == int32(44) {
		goto L184
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v1070 = v708 + int32(1079)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+720)) = v1070
	*(*int32)(unsafe.Add(mBase, uint32(v708)+716)) = v708 + int32(1088)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+712)) = v708 + int32(1084)
	v1079 = v708 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+708)) = v1079
	v1082 = v708 + int32(892)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+704)) = v1082
	v1087 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_60), v708+int32(704))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L32
	} else {
		goto L294
	}
L288:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L32
	} else {
		goto L289
	}
L289:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L32
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+832)) = int32(_a_F_StartupXLOG_58)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v708+int32(832))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L32
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1207), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L32
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L32
	} else {
		goto L460
	}
L294:
	;
	if v1087 != int32(5) {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+1079)))
	if v1091 != int32(10) {
		goto L293
	} else {
		goto L296
	}
L296:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v1095
	v1098 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v708)+888)))
	v1099 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v708)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = v1098 | v1099<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+688)) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(v708)+692)) = v1079
	*(*int32)(unsafe.Add(mBase, uint32(v708)+696)) = v1070
	v1110 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_63), v708+int32(688))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L32
	} else {
		goto L298
	}
L297:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L32
	} else {
		goto L456
	}
L298:
	;
	if v1110 != int32(3) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+1079)))
	if v1114 != int32(10) {
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = v1118
	v1121 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v708)+888)))
	v1122 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v708)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v1121 | v1122<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+672)) = v708 + int32(1056)
	v1133 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_64), v708+int32(672))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L32
	} else {
		goto L302
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+656)) = v708 + int32(1024)
	v1153 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_65), v708+int32(656))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L32
	} else {
		goto L305
	}
L302:
	;
	if v1133 != int32(1) {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1137 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v708)+1064)))
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(v708)+1056))
	if v1137|(v1138^int64(7234308641521824883)) != int64(0) {
		goto L301
	} else {
		goto L304
	}
L304:
	;
	v1145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = uint8(v1145)
	goto L301
L305:
	;
	v1155 = *(*int64)(unsafe.Add(mBase, uint32(v708)+1024))
	v1157 = v708 + int32(896)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+640)) = v1157
	v1162 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_66), v708+int32(640))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L32
	} else {
		goto L307
	}
L306:
	;
	v1186 = v708 + int32(1152)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+608)) = v1186
	v1191 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_67), v708+int32(608))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L32
	} else {
		goto L314
	}
L307:
	;
	if v1162 != int32(1) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1168 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L32
	} else {
		goto L309
	}
L309:
	;
	if v1168 == int32(0) {
		goto L306
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+628)) = int32(_a_F_StartupXLOG_58)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+624)) = v1157
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_68), v708+int32(624))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L32
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1270), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L32
	} else {
		goto L312
	}
L312:
	;
	goto L306
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+576)) = v708 + int32(1080)
	v1220 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_69), v708+int32(576))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L32
	} else {
		goto L322
	}
L314:
	;
	if v1191 != int32(1) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1197 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L32
	} else {
		goto L316
	}
L316:
	;
	if v1197 == int32(0) {
		goto L313
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+596)) = int32(_a_F_StartupXLOG_58)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+592)) = v1186
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_70), v708+int32(592))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L32
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1275), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L32
	} else {
		goto L319
	}
L319:
	;
	goto L313
L320:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L32
	} else {
		goto L451
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+516)) = v708 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+512)) = v708 + int32(892)
	v1256 = F_fscanf(m, v1043, int32(_a_F_StartupXLOG_71), v708+int32(512))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L32
	} else {
		goto L329
	}
L322:
	;
	if v1220 != int32(1) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1084))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1080))
	if v1224 != v1225 {
		goto L320
	} else {
		goto L324
	}
L324:
	;
	v1229 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L32
	} else {
		goto L325
	}
L325:
	;
	if v1229 == int32(0) {
		goto L321
	} else {
		goto L326
	}
L326:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+528)) = v1233
	*(*int32)(unsafe.Add(mBase, uint32(v708)+532)) = int32(_a_F_StartupXLOG_58)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_72), v708+int32(528))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L32
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1292), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L32
	} else {
		goto L328
	}
L328:
	;
	goto L321
L329:
	;
	if v1256 <= int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1043)))
	goto L334
L331:
	;
	goto L332
L332:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L32
	} else {
		goto L446
	}
L333:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L32
	} else {
		goto L442
	}
L334:
	;
	if int32(base.Ui32(v1260)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1265 = F_FreeFile(m, v1043)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L32
	} else {
		goto L336
	}
L336:
	;
	if v1265 != 0 {
		goto L333
	} else {
		goto L337
	}
L337:
	;
	v1268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v1268)
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v1271 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v1273)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L32
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1279 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L32
	} else {
		goto L342
	}
L341:
	;
	goto L340
L342:
	;
	if v1279 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+480)) = v1282
	v1285 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[40]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+468)) = uint32(v1285)
	v1288 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+476)) = uint32(v1288)
	v1290 = int64(32)
	v1291 = int64(base.Ui64(v1285) >> (uint(v1290) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+464)) = uint32(v1291)
	v1294 = int64(base.Ui64(v1288) >> (uint(v1290) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+472)) = uint32(v1294)
	F_errmsg(m, int32(_a_F_StartupXLOG_73), v708+int32(464))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L32
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v1311 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v1313 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37]))
	v1314 = F_ReadCheckpointRecord(m, v1309, v1311, v1313)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L32
	} else {
		goto L349
	}
L346:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(574), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L32
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v1448 = base.B2i32(v1153 == int32(1)) & base.B2i32(v1155 == int64(34166655670121587))
	v1451 = F_AllocateFile(m, int32(_a_F_StartupXLOG_44), int32(_a_F_StartupXLOG_59))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L32
	} else {
		goto L371
	}
L349:
	;
	if v1314 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+96))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+64))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+8))
	v1321 = *(*int64)(unsafe.Add(mBase, uint32(v1319)))
	v1322 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v708)+896)) = v1322
	v1324 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v708)+904)) = v1324
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+912)) = v1326
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314)+16)))
	v1329 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+88))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+84))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+80))
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+72))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+68))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+64))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+60))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+56))
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+48))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+44))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+40))
	v1340 = *(*int64)(unsafe.Add(mBase, uint32(v1319)+32))
	v1343 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L32
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L32
	} else {
		goto L367
	}
L353:
	;
	if v1343 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1346 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+452)) = uint32(v1346)
	v1349 = int64(base.Ui64(v1346) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+448)) = uint32(v1349)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_75), v708+int32(448))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L32
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])) = uint8(v1363)
	v1366 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if base.Ui64(v1366) <= base.Ui64(v1321) {
		goto L348
	} else {
		goto L359
	}
L357:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(588), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L32
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	F_XLogPrefetcherBeginRead(m, v1369, v1321)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L32
	} else {
		goto L360
	}
L360:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v1376 = F_ReadRecord(m, v1373, int32(15), int32(0), v1320)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L32
	} else {
		goto L361
	}
L361:
	;
	if v1376 != 0 {
		goto L348
	} else {
		goto L362
	}
L362:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L32
	} else {
		goto L363
	}
L363:
	;
	v1383 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+92)) = uint32(v1383)
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+84)) = uint32(v1321)
	v1386 = int64(32)
	v1387 = int64(base.Ui64(v1321) >> (uint(v1386) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+80)) = uint32(v1387)
	v1390 = int64(base.Ui64(v1383) >> (uint(v1386) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+88)) = uint32(v1390)
	F_errmsg(m, int32(_a_F_StartupXLOG_76), v708+int32(80))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L32
	} else {
		goto L364
	}
L364:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+64)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v708)+68)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v708)+72)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v708)+76)) = v1398
	F_errhint(m, int32(_a_F_StartupXLOG_77), v708-int32(-64))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L32
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(608), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L32
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	v1418 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+52)) = uint32(v1418)
	v1421 = int64(base.Ui64(v1418) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+48)) = uint32(v1421)
	F_errmsg(m, int32(_a_F_StartupXLOG_78), v708+int32(48))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L32
	} else {
		goto L368
	}
L368:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+32)) = v1429
	*(*int32)(unsafe.Add(mBase, uint32(v708)+36)) = v1429
	*(*int32)(unsafe.Add(mBase, uint32(v708)+40)) = v1429
	*(*int32)(unsafe.Add(mBase, uint32(v708)+44)) = v1429
	F_errhint(m, int32(_a_F_StartupXLOG_77), v708+int32(32))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L32
	} else {
		goto L369
	}
L369:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(619), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L32
	} else {
		goto L370
	}
L370:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L371:
	;
	if v1451 == int32(0) {
		goto L183
	} else {
		goto L372
	}
L372:
	;
	v1455 = F_do_getc(m, v1451)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L32
	} else {
		goto L374
	}
L373:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1451)))
	goto L417
L374:
	;
	if v1455 == int32(-1) {
		v1796 = v1
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v1471 = int32(0)
	v1472 = v1455
	v1473 = v1
	v1474 = v1
	goto L376
L376:
	;
	if v1473&int32(1) != 0 {
		goto L382
	} else {
		goto L383
	}
L377:
	;
	if base.B2i32(v1729 == int32(0))&(v1731^int32(1)) != 0 {
		v1796 = v1732
		goto L373
	} else {
		goto L411
	}
L378:
	;
	v1754 = F_do_getc(m, v1451)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L32
	} else {
		goto L409
	}
L379:
	;
	v1729 = v1692
	v1731 = int32(0)
	v1732 = v1695
	goto L378
L380:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L32
	} else {
		goto L405
	}
L381:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L32
	} else {
		goto L401
	}
L382:
	;
	v1632 = (v1473 ^ int32(1)) & base.B2i32(v1472 == int32(92))
	if v1632|base.B2i32(base.Ui32(int32(1022)) < base.Ui32(v1471)) != 0 {
		v1729 = v1471
		v1731 = v1632
		v1732 = v1474
		goto L378
	} else {
		goto L400
	}
L383:
	;
	switch v1472 - int32(10) {
	case 0, 3:
		goto L384
	default:
		goto L382
	}
L384:
	;
	if v1471 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1500 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v708+int32(1152)+v1471))) = uint8(v1500)
	v1519 = v1500
	goto L388
L386:
	;
	v1605 = v1474
	goto L387
L387:
	;
	v1692 = int32(0)
	v1695 = v1605
	goto L379
L388:
	;
	v1544 = v708 + int32(1152) + v1519
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1544))))
	v1546 = int32(32)
	if v1545|v1546 != v1546 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	if base.B2i32(v1519 <= int32(0))|base.B2i32(v1471-int32(1) <= v1519) != 0 {
		goto L381
	} else {
		goto L393
	}
L390:
	;
	v1519 = v1519 + int32(1)
	goto L388
L391:
	;
	goto L392
L392:
	;
	goto L389
L393:
	;
	v1558 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1544))) = uint8(v1558)
	v1561 = F_palloc0(m, int32(24))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L32
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v1572 = F_strtox_2(m, v708+int32(1152), v708+int32(1088), int32(10), int64(4294967295))
	mBase = m.M
	goto L395
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1561))) = base.I32_wrap_i64(v1572)
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1088))
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1575))))
	if v1576 != 0 {
		goto L380
	} else {
		goto L396
	}
L396:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if base.B2i32(v1578 == int32(68))|base.B2i32(v1578 == int32(28)) != 0 {
		goto L380
	} else {
		goto L397
	}
L397:
	;
	v1586 = F_pstrdup(m, v1544+int32(1))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L32
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1561)+4)) = v1586
	v1589 = F_lappend(m, v1474, v1561)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L32
	} else {
		goto L399
	}
L399:
	;
	v1605 = v1589
	goto L387
L400:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v708+int32(1152)+v1471))) = uint8(v1472)
	v1692 = v1471 + int32(1)
	v1695 = v1474
	goto L379
L401:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L32
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+432)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708+int32(432))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L32
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1374), int32(_a_F_StartupXLOG_80))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L32
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L32
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+416)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708+int32(416))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L32
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1383), int32(_a_F_StartupXLOG_80))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L32
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	if v1754 != int32(-1) {
		v1471 = v1729
		v1472 = v1754
		v1473 = v1731
		v1474 = v1732
		goto L376
	} else {
		goto L410
	}
L410:
	;
	goto L377
L411:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L32
	} else {
		goto L412
	}
L412:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L32
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+400)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708+int32(400))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L32
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1403), int32(_a_F_StartupXLOG_80))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L32
	} else {
		goto L415
	}
L415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L416:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L32
	} else {
		goto L438
	}
L417:
	;
	if int32(base.Ui32(v1818)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L416
	} else {
		goto L418
	}
L418:
	;
	v1823 = F_FreeFile(m, v1451)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L32
	} else {
		goto L419
	}
L419:
	;
	if v1823 != 0 {
		goto L416
	} else {
		goto L420
	}
L420:
	;
	if v1796 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2352 = v1339
	v2353 = v1448
	v2364 = int32(1)
	v2366 = v1335
	v2368 = v1331
	v2369 = v1338
	v2371 = v1330
	v2372 = v1336
	v2373 = v1320
	v2374 = v1333
	v2375 = v1334
	v2376 = v1328
	v2379 = v1321
	v2382 = v1332
	v2383 = v1329
	v2384 = v1340
	v2386 = v1337
	goto L181
L422:
	;
	goto L423
L423:
	;
	v1828 = int32(1)
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+4))
	if v1829 <= int32(0) {
		v2352 = v1339
		v2353 = v1448
		v2364 = v1828
		v2366 = v1335
		v2368 = v1331
		v2369 = v1338
		v2371 = v1330
		v2372 = v1336
		v2373 = v1320
		v2374 = v1333
		v2375 = v1334
		v2376 = v1328
		v2379 = v1321
		v2382 = v1332
		v2383 = v1329
		v2384 = v1340
		v2386 = v1337
		goto L181
	} else {
		goto L424
	}
L424:
	;
	v1845 = int32(0)
	goto L425
L425:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+12))
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1869+v1845<<(uint(int32(2))%32))))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1873)))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+372)) = v1874
	*(*int32)(unsafe.Add(mBase, uint32(v708)+368)) = int32(_a_F_StartupXLOG_38)
	v1881 = F_psprintf(m, int32(_a_F_StartupXLOG_81), v708+int32(368))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L32
	} else {
		goto L428
	}
L426:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L32
	} else {
		goto L434
	}
L427:
	;
	goto L426
L428:
	;
	F_remove_tablespace_symlink(m, v1881)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L32
	} else {
		goto L429
	}
L429:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	v1886 = F_symlink(m, v1885, v1881)
	mBase = m.M
	if v1886 < int32(0) {
		goto L427
	} else {
		goto L430
	}
L430:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	F_pfree(m, v1889)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L32
	} else {
		goto L431
	}
L431:
	;
	F_pfree(m, v1873)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L32
	} else {
		goto L432
	}
L432:
	;
	v1895 = v1845 + int32(1)
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+4))
	if v1895 < v1896 {
		v1845 = v1895
		goto L425
	} else {
		goto L433
	}
L433:
	;
	v2352 = v1339
	v2353 = v1448
	v2364 = v1828
	v2366 = v1335
	v2368 = v1331
	v2369 = v1338
	v2371 = v1330
	v2372 = v1336
	v2373 = v1320
	v2374 = v1333
	v2375 = v1334
	v2376 = v1328
	v2379 = v1321
	v2382 = v1332
	v2383 = v1329
	v2384 = v1340
	v2386 = v1337
	goto L181
L434:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L32
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+352)) = v1881
	F_errmsg(m, int32(_a_F_StartupXLOG_82), v708+int32(352))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L32
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(645), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L32
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L32
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+384)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v708+int32(384))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L32
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1409), int32(_a_F_StartupXLOG_80))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L32
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L32
	} else {
		goto L443
	}
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+496)) = int32(_a_F_StartupXLOG_58)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v708+int32(496))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L32
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1305), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L32
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L32
	} else {
		goto L447
	}
L447:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_83), int32(0))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L32
	} else {
		goto L448
	}
L448:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_84), int32(0))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L32
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1299), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L32
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L32
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+560)) = int32(_a_F_StartupXLOG_58)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708+int32(560))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L32
	} else {
		goto L453
	}
L453:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+544)) = v1985
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v708)+1084))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+548)) = v1987
	v1992 = F_errdetail(m, int32(_a_F_StartupXLOG_85), v708+int32(544))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L32
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1288), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L32
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L32
	} else {
		goto L457
	}
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+16)) = int32(_a_F_StartupXLOG_58)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708+int32(16))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L32
	} else {
		goto L458
	}
L458:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1227), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L32
	} else {
		goto L459
	}
L459:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L460:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L32
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = int32(_a_F_StartupXLOG_58)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v708)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L32
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1220), int32(_a_F_StartupXLOG_62))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L32
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	F_errcode(m, int32(_a_F_StartupXLOG_86))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L32
	} else {
		goto L465
	}
L465:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_87), int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L32
	} else {
		goto L466
	}
L466:
	;
	v2048 = F_errdetail(m, int32(_a_F_StartupXLOG_88), int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L32
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(519), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L32
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L32
	} else {
		goto L470
	}
L470:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_89), int32(0))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L32
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1099), int32(_a_F_StartupXLOG_54))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L32
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L32
	} else {
		goto L474
	}
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+864)) = int32(_a_F_StartupXLOG_45)
	F_errmsg(m, int32(_a_F_StartupXLOG_90), v708+int32(864))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L32
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1007), int32(_a_F_StartupXLOG_51))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L32
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2137 != int32(1) {
		goto L492
	} else {
		goto L493
	}
L478:
	;
	if v2094 != 0 {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v2095 = int32(_a_F_StartupXLOG_91)
	v2096 = F_unlink(m, v2095)
	mBase = m.M
	v2100 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v2095, int32(14))
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L32
	} else {
		goto L480
	}
L480:
	;
	v2104 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L32
	} else {
		goto L481
	}
L481:
	;
	if v2104 == int32(0) {
		goto L477
	} else {
		goto L482
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+820)) = int32(_a_F_StartupXLOG_58)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+816)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_92), v708+int32(816))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L32
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+804)) = int32(_a_F_StartupXLOG_91)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+800)) = int32(_a_F_StartupXLOG_44)
	if v2100 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2123 = int32(_a_F_StartupXLOG_93)
	goto L486
L485:
	;
	v2123 = int32(_a_F_StartupXLOG_94)
	goto L486
L486:
	;
	v2126 = F_errdetail(m, v2123, v708+int32(800))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L32
	} else {
		goto L487
	}
L487:
	;
	if v2100 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2131 = int32(686)
	goto L490
L489:
	;
	v2131 = int32(680)
	goto L490
L490:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2131, int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L32
	} else {
		goto L491
	}
L491:
	;
	goto L477
L492:
	;
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v672)+160))
	if v2162 == int64(0) {
		goto L501
	} else {
		goto L502
	}
L493:
	;
	v2140 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	if v2140 != int64(0) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v2151)
	v2154 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v2154 == int32(0) {
		goto L492
	} else {
		goto L499
	}
L495:
	;
	v2143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+176)))
	if v2143 != 0 {
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v2144 = *(*int64)(unsafe.Add(mBase, uint32(v672)+168))
	if v2144 != int64(0) {
		goto L494
	} else {
		goto L497
	}
L497:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v672)+16))
	if v2147 != int32(1) {
		goto L492
	} else {
		goto L498
	}
L498:
	;
	goto L494
L499:
	;
	v2158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v2158)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L32
	} else {
		goto L500
	}
L500:
	;
	goto L492
L501:
	;
	v2188 = *(*int64)(unsafe.Add(mBase, uint32(v672)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v2188
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v672)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = v2191
	v2193 = *(*int64)(unsafe.Add(mBase, uint32(v672)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v2191
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = v2193
	v2199 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v2200 = F_ReadCheckpointRecord(m, v2199, v2188, v2191)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L32
	} else {
		goto L508
	}
L502:
	;
	v2167 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L32
	} else {
		goto L503
	}
L503:
	;
	if v2167 == int32(0) {
		goto L501
	} else {
		goto L504
	}
L504:
	;
	v2171 = *(*int64)(unsafe.Add(mBase, uint32(v672)+160))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+788)) = uint32(v2171)
	v2174 = int64(base.Ui64(v2171) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+784)) = uint32(v2174)
	F_errmsg(m, int32(_a_F_StartupXLOG_95), v708+int32(784))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L32
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(725), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L32
	} else {
		goto L506
	}
L506:
	;
	goto L501
L507:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L32
	} else {
		goto L525
	}
L508:
	;
	if v2200 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v2204 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L32
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L32
	} else {
		goto L522
	}
L512:
	;
	if v2204 != 0 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v2207 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+772)) = uint32(v2207)
	v2210 = int64(base.Ui64(v2207) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+768)) = uint32(v2210)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_75), v708+int32(768))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L32
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+96))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2225)+64))
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+8))
	v2228 = *(*int64)(unsafe.Add(mBase, uint32(v2226)))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+912)) = v2229
	v2231 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v708)+904)) = v2231
	v2233 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v708)+896)) = v2233
	v2235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2200)+16)))
	v2236 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+88))
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+84))
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+80))
	v2239 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+72))
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+68))
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+64))
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+60))
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+56))
	v2244 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+48))
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+44))
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+40))
	v2247 = *(*int64)(unsafe.Add(mBase, uint32(v2226)+32))
	v2249 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if base.Ui64(v2249) <= base.Ui64(v2228) {
		v2330 = v2246
		v2331 = v1
		v2334 = v2242
		v2336 = v2238
		v2337 = v2245
		v2338 = v2237
		v2339 = v2243
		v2340 = v2227
		v2341 = v2240
		v2342 = v2241
		v2343 = v2235
		v2344 = v2228
		v2347 = v2239
		v2348 = v2236
		v2349 = v2247
		v2350 = v2244
		goto L182
	} else {
		goto L518
	}
L516:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(738), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L32
	} else {
		goto L517
	}
L517:
	;
	goto L515
L518:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	F_XLogPrefetcherBeginRead(m, v2252, v2228)
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L32
	} else {
		goto L519
	}
L519:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v2259 = F_ReadRecord(m, v2256, int32(15), int32(0), v2227)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L32
	} else {
		goto L520
	}
L520:
	;
	if v2259 == int32(0) {
		goto L507
	} else {
		goto L521
	}
L521:
	;
	v2330 = v2246
	v2331 = v1
	v2334 = v2242
	v2336 = v2238
	v2337 = v2245
	v2338 = v2237
	v2339 = v2243
	v2340 = v2227
	v2341 = v2240
	v2342 = v2241
	v2343 = v2235
	v2344 = v2228
	v2347 = v2239
	v2348 = v2236
	v2349 = v2247
	v2350 = v2244
	goto L182
L522:
	;
	v2268 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+740)) = uint32(v2268)
	v2271 = int64(base.Ui64(v2268) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+736)) = uint32(v2271)
	F_errmsg(m, int32(_a_F_StartupXLOG_96), v708+int32(736))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L32
	} else {
		goto L523
	}
L523:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(750), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L32
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	v2288 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+764)) = uint32(v2288)
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+756)) = uint32(v2228)
	v2291 = int64(32)
	v2292 = int64(base.Ui64(v2228) >> (uint(v2291) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+752)) = uint32(v2292)
	v2295 = int64(base.Ui64(v2288) >> (uint(v2291) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+760)) = uint32(v2295)
	F_errmsg(m, int32(_a_F_StartupXLOG_76), v708+int32(752))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L32
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(762), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L32
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L32
	} else {
		goto L529
	}
L529:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L32
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+336)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v708+int32(336))
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L32
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1342), int32(_a_F_StartupXLOG_80))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L32
	} else {
		goto L532
	}
L532:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L533:
	;
	v2487 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v2489 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v2490 = F_tliOfPointInHistory(m, v2487, v2489)
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L32
	} else {
		goto L567
	}
L534:
	;
	v2393 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v2393 != 0 {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2479, int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L32
	} else {
		goto L562
	}
L536:
	;
	v2396 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L32
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	v2409 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L32
	} else {
		goto L542
	}
L539:
	;
	if v2396 == int32(0) {
		goto L533
	} else {
		goto L540
	}
L540:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_97), int32(0))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L32
	} else {
		goto L541
	}
L541:
	;
	v2479 = int32(770)
	goto L535
L542:
	;
	switch v2406 - int32(1) {
	case 0:
		goto L548
	case 1:
		goto L547
	case 2:
		goto L546
	case 3:
		goto L545
	case 4:
		goto L544
	default:
		goto L543
	}
L543:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L560
	}
L544:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L558
	}
L545:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L556
	}
L546:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L554
	}
L547:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L551
	}
L548:
	;
	if v2409 == int32(0) {
		goto L533
	} else {
		goto L549
	}
L549:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+272)) = v2416
	F_errmsg(m, int32(_a_F_StartupXLOG_98), v708+int32(272))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L32
	} else {
		goto L550
	}
L550:
	;
	v2479 = int32(774)
	goto L535
L551:
	;
	v2427 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[26]))
	v2428 = F_timestamptz_to_str(m, v2427)
	mBase = m.M
	v2429 = m.ExcPending
	if v2429 != 0 {
		goto L32
	} else {
		goto L552
	}
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+288)) = v2428
	F_errmsg(m, int32(_a_F_StartupXLOG_99), v708+int32(288))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L32
	} else {
		goto L553
	}
L553:
	;
	v2479 = int32(778)
	goto L535
L554:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+304)) = v2440
	F_errmsg(m, int32(_a_F_StartupXLOG_100), v708+int32(304))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L32
	} else {
		goto L555
	}
L555:
	;
	v2479 = int32(782)
	goto L535
L556:
	;
	v2451 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+324)) = uint32(v2451)
	v2454 = int64(base.Ui64(v2451) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+320)) = uint32(v2454)
	F_errmsg(m, int32(_a_F_StartupXLOG_101), v708+int32(320))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L32
	} else {
		goto L557
	}
L557:
	;
	v2479 = int32(786)
	goto L535
L558:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_102), int32(0))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L32
	} else {
		goto L559
	}
L559:
	;
	v2479 = int32(789)
	goto L535
L560:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_103), int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L32
	} else {
		goto L561
	}
L561:
	;
	v2479 = int32(792)
	goto L535
L562:
	;
	goto L533
L563:
	;
	v2908 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+128))
	v2910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2908)+56)))
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+48))
	v2912 = *(*int64)(unsafe.Add(mBase, uint32(v2908)+40))
	v2913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2908)+64)))
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+124))
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+120))
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+108))
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+104))
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+100))
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+96))
	v2920 = *(*int64)(unsafe.Add(mBase, uint32(v2908)+88))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+84))
	v2922 = *(*int64)(unsafe.Add(mBase, uint32(v2908)+72))
	v2923 = int32(_a_F_StartupXLOG_104)
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v2924))) = v2925
	*(*int64)(unsafe.Add(mBase, uint32(v2924)+8)) = v2922
	v2929 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v2929)+4)) = int32(0)
	F_MultiXactSetNextMXact(m, v2921, v2920)
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L32
	} else {
		goto L676
	}
L564:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L32
	} else {
		goto L673
	}
L565:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2884 = m.ExcPending
	if v2884 != 0 {
		goto L32
	} else {
		goto L670
	}
L566:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L32
	} else {
		goto L667
	}
L567:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37]))
	if v2490 == v2493 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v2495 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	if v2495 != int64(0) {
		goto L571
	} else {
		goto L572
	}
L569:
	;
	goto L570
L570:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v2816 = F_tliSwitchPoint(m, v2493, v2814, int32(0))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L32
	} else {
		goto L659
	}
L571:
	;
	v2501 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v2502 = F_tliOfPointInHistory(m, v2495-int64(1), v2501)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L32
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v2508 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L32
	} else {
		goto L576
	}
L574:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v672)+152))
	if v2502 != v2504 {
		goto L566
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	if v2508 != 0 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+196)) = uint32(v2379)
	v2512 = int64(base.Ui64(v2379) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+192)) = uint32(v2512)
	if base.Ui32(v2376) < base.Ui32(int32(16)) {
		goto L580
	} else {
		goto L581
	}
L578:
	;
	goto L579
L579:
	;
	v2532 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L32
	} else {
		goto L585
	}
L580:
	;
	v2518 = int32(_a_F_StartupXLOG_105)
	goto L582
L581:
	;
	v2518 = int32(_a_F_StartupXLOG_106)
	goto L582
L582:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+200)) = v2518
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_107), v708+int32(192))
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L32
	} else {
		goto L583
	}
L583:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(839), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L32
	} else {
		goto L584
	}
L584:
	;
	goto L579
L585:
	;
	if v2532 != 0 {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+184)) = v2352
	*(*int64)(unsafe.Add(mBase, uint32(v708)+176)) = v2384
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_108), v708+int32(176))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L32
	} else {
		goto L589
	}
L587:
	;
	goto L588
L588:
	;
	v2548 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L32
	} else {
		goto L591
	}
L589:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(843), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L32
	} else {
		goto L590
	}
L590:
	;
	goto L588
L591:
	;
	if v2548 != 0 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v708)+168)) = v2386
	*(*int32)(unsafe.Add(mBase, uint32(v708)+160)) = v2369
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_109), v708+int32(160))
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L32
	} else {
		goto L595
	}
L593:
	;
	goto L594
L594:
	;
	v2564 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L32
	} else {
		goto L597
	}
L595:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(846), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2561 = m.ExcPending
	if v2561 != 0 {
		goto L32
	} else {
		goto L596
	}
L596:
	;
	goto L594
L597:
	;
	if v2564 != 0 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+148)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v708)+144)) = v2372
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_110), v708+int32(144))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L32
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	v2580 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L32
	} else {
		goto L603
	}
L601:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(849), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L32
	} else {
		goto L602
	}
L602:
	;
	goto L600
L603:
	;
	if v2580 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+132)) = v2374
	*(*int32)(unsafe.Add(mBase, uint32(v708)+128)) = v2375
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_111), v708+int32(128))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L32
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	v2596 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L32
	} else {
		goto L609
	}
L607:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(852), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L32
	} else {
		goto L608
	}
L608:
	;
	goto L606
L609:
	;
	if v2596 != 0 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+116)) = v2371
	*(*int32)(unsafe.Add(mBase, uint32(v708)+112)) = v2368
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_112), v708+int32(112))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L32
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	if base.Ui32(base.I32_wrap_i64(v2384)) <= base.Ui32(int32(2)) {
		goto L565
	} else {
		goto L615
	}
L613:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(856), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L32
	} else {
		goto L614
	}
L614:
	;
	goto L612
L615:
	;
	v2614 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if base.Ui64(v2614) < base.Ui64(v2379) {
		goto L564
	} else {
		goto L616
	}
L616:
	;
	if base.Ui64(v2379) < base.Ui64(v2614) {
		goto L623
	} else {
		goto L624
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[50])) = v2789
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51])) = v2792
	v2798 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v2798
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v2798
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(4095)))) = uint8(base.B2i32(base.Ui32(v2376) < base.Ui32(int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(4093)))) = uint8(base.B2i32(v1043 != int32(0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(4094)))) = uint8(v2364)
	m.G0 = v708 + int32(2176)
	goto L563
L618:
	;
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v672)+152))
	v2788 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	v2789 = v2787
	v2792 = v2788
	goto L617
L619:
	;
	v2773 = *(*int64)(unsafe.Add(mBase, uint32(v672)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2773
	v2776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+176)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = uint8(v2776)
	v2779 = *(*int64)(unsafe.Add(mBase, uint32(v672)+168))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[55])) = v2779
	if v2769&int32(1) != 0 {
		goto L618
	} else {
		goto L658
	}
L620:
	;
	if v2652&int32(1) != 0 {
		v2695 = int32(5)
		goto L633
	} else {
		goto L634
	}
L621:
	;
	v2647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	v2649 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])))
	if v2649 != int32(1) {
		v2769 = v2647
		goto L619
	} else {
		goto L632
	}
L622:
	;
	v2642 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])) = uint8(v2642)
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	v2652 = v2645
	goto L620
L623:
	;
	if base.Ui32(int32(15)) < base.Ui32(v2376) {
		goto L622
	} else {
		goto L626
	}
L624:
	;
	goto L625
L625:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v672)+16))
	if v2632 != int32(1) {
		goto L622
	} else {
		goto L630
	}
L626:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L32
	} else {
		goto L627
	}
L627:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_113), int32(0))
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L32
	} else {
		goto L628
	}
L628:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(875), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L32
	} else {
		goto L629
	}
L629:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L630:
	;
	v2636 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2636&int32(1) == int32(0) {
		goto L621
	} else {
		goto L631
	}
L631:
	;
	goto L622
L632:
	;
	v2652 = v2647
	goto L620
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v672)+16)) = v2695
	v2698 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v672)+48)) = v2373
	*(*int64)(unsafe.Add(mBase, uint32(v672)+40)) = v2379
	*(*int64)(unsafe.Add(mBase, uint32(v672)+32)) = v2698
	v2702 = *(*int64)(unsafe.Add(mBase, uint32(v708)+896))
	*(*int64)(unsafe.Add(mBase, uint32(v672)+52)) = v2702
	v2704 = *(*int64)(unsafe.Add(mBase, uint32(v708)+904))
	*(*int64)(unsafe.Add(mBase, uint32(v672)+60)) = v2704
	v2706 = *(*int32)(unsafe.Add(mBase, uint32(v708)+912))
	*(*int32)(unsafe.Add(mBase, uint32(v672)+68)) = v2706
	*(*int64)(unsafe.Add(mBase, uint32(v672)+128)) = v2383
	*(*int32)(unsafe.Add(mBase, uint32(v672)+124)) = v2371
	*(*int32)(unsafe.Add(mBase, uint32(v672)+120)) = v2368
	*(*int64)(unsafe.Add(mBase, uint32(v672)+112)) = v2382
	*(*int32)(unsafe.Add(mBase, uint32(v672)+108)) = v2374
	*(*int32)(unsafe.Add(mBase, uint32(v672)+104)) = v2375
	*(*int32)(unsafe.Add(mBase, uint32(v672)+100)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v672)+96)) = v2372
	*(*int64)(unsafe.Add(mBase, uint32(v672)+88)) = v2386
	*(*int32)(unsafe.Add(mBase, uint32(v672)+84)) = v2369
	*(*int32)(unsafe.Add(mBase, uint32(v672)+80)) = v2352
	*(*int64)(unsafe.Add(mBase, uint32(v672)+72)) = v2384
	v2721 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v2721 != int32(1) {
		goto L646
	} else {
		goto L647
	}
L634:
	;
	v2658 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2659 = m.ExcPending
	if v2659 != 0 {
		goto L32
	} else {
		goto L635
	}
L635:
	;
	if v2658 != 0 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_114), int32(0))
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		goto L32
	} else {
		goto L639
	}
L637:
	;
	goto L638
L638:
	;
	v2669 = int32(4)
	v2671 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v672)+48))
	if base.Ui32(v2671) <= base.Ui32(v2672) {
		v2695 = v2669
		goto L633
	} else {
		goto L641
	}
L639:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(905), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2668 = m.ExcPending
	if v2668 != 0 {
		goto L32
	} else {
		goto L640
	}
L640:
	;
	goto L638
L641:
	;
	v2676 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L32
	} else {
		goto L642
	}
L642:
	;
	if v2676 == int32(0) {
		v2695 = v2669
		goto L633
	} else {
		goto L643
	}
L643:
	;
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v672)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+96)) = v2680
	v2683 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+100)) = v2683
	F_errmsg(m, int32(_a_F_StartupXLOG_115), v708+int32(96))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L32
	} else {
		goto L644
	}
L644:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(911), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L32
	} else {
		goto L645
	}
L645:
	;
	v2695 = v2669
	goto L633
L646:
	;
	if v1043 == int32(0) {
		v2769 = v2721
		goto L619
	} else {
		goto L649
	}
L647:
	;
	v2724 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	if base.Ui64(v2379) <= base.Ui64(v2724) {
		goto L646
	} else {
		goto L648
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v672)+152)) = v2373
	*(*int64)(unsafe.Add(mBase, uint32(v672)+144)) = v2379
	goto L646
L649:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v672)+160)) = v2379
	v2732 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[38])))
	*(*uint8)(unsafe.Add(mBase, uint32(v672)+176)) = uint8(v2732)
	if v2353 == int32(0) {
		v2769 = v2721
		goto L619
	} else {
		goto L650
	}
L650:
	;
	switch v710 - int32(2) {
	case 0, 3:
		goto L651
	default:
		goto L652
	}
L651:
	;
	v2755 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v672)+168)) = v2755
	v2759 = *(*int64)(unsafe.Add(mBase, uint32(v672)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2759
	v2762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+176)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = uint8(v2762)
	v2765 = *(*int64)(unsafe.Add(mBase, uint32(v672)+168))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[55])) = v2765
	if v2721 != 0 {
		goto L618
	} else {
		goto L657
	}
L652:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2741 = m.ExcPending
	if v2741 != 0 {
		goto L32
	} else {
		goto L653
	}
L653:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_116), int32(0))
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L32
	} else {
		goto L654
	}
L654:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_117), int32(0))
	mBase = m.M
	v2749 = m.ExcPending
	if v2749 != 0 {
		goto L32
	} else {
		goto L655
	}
L655:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(953), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L32
	} else {
		goto L656
	}
L656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L657:
	;
	v2789 = int32(0)
	v2792 = int64(0)
	goto L617
L658:
	;
	v2789 = int32(0)
	v2792 = int64(0)
	goto L617
L659:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L32
	} else {
		goto L660
	}
L660:
	;
	v2823 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+256)) = v2823
	F_errmsg(m, int32(_a_F_StartupXLOG_118), v708+int32(256))
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		goto L32
	} else {
		goto L661
	}
L661:
	;
	v2831 = int64(base.Ui64(v2816) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+240)) = uint32(v2831)
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+244)) = uint32(v2816)
	if v1043 != 0 {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v2836 = int32(_a_F_StartupXLOG_58)
	goto L664
L663:
	;
	v2836 = int32(_a_F_StartupXLOG_119)
	goto L664
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+224)) = v2836
	v2839 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+232)) = uint32(v2839)
	v2842 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+236)) = v2842
	v2845 = int64(base.Ui64(v2839) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+228)) = uint32(v2845)
	v2850 = F_errdetail(m, int32(_a_F_StartupXLOG_120), v708+int32(224))
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L32
	} else {
		goto L665
	}
L665:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(820), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L32
	} else {
		goto L666
	}
L666:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L667:
	;
	v2861 = *(*int32)(unsafe.Add(mBase, uint32(v672)+152))
	v2862 = *(*int64)(unsafe.Add(mBase, uint32(v672)+144))
	v2864 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+208)) = v2864
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+216)) = uint32(v2862)
	*(*int32)(unsafe.Add(mBase, uint32(v708)+220)) = v2861
	v2869 = int64(base.Ui64(v2862) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v708)+212)) = uint32(v2869)
	F_errmsg(m, int32(_a_F_StartupXLOG_121), v708+int32(208))
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L32
	} else {
		goto L668
	}
L668:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(834), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		goto L32
	} else {
		goto L669
	}
L669:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L670:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_122), int32(0))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		goto L32
	} else {
		goto L671
	}
L671:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(859), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L32
	} else {
		goto L672
	}
L672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L673:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_123), int32(0))
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L32
	} else {
		goto L674
	}
L674:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(864), int32(_a_F_StartupXLOG_74))
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L32
	} else {
		goto L675
	}
L675:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L676:
	;
	F_AdvanceOldestClogXid(m, v2919)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L32
	} else {
		goto L677
	}
L677:
	;
	F_SetTransactionIdLimit(m, v2919, v2918)
	mBase = m.M
	v2937 = m.ExcPending
	if v2937 != 0 {
		goto L32
	} else {
		goto L678
	}
L678:
	;
	F_SetMultiXactIdLimit(m, v2917, v2916)
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L32
	} else {
		goto L679
	}
L679:
	;
	F_SetCommitTsLimit(m, v2915, v2914)
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L32
	} else {
		goto L680
	}
L680:
	;
	v2942 = m.G0
	v2944 = v2942 - int32(1088)
	m.G0 = v2944
	*(*int32)(unsafe.Add(mBase, uint32(v2944)+16)) = int32(_a_F_StartupXLOG_124)
	v2949 = v2944 + int32(32)
	v2954 = F_pg_snprintf(m, v2949, int32(1050), int32(_a_F_StartupXLOG_125), v2944+int32(16))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L32
	} else {
		goto L681
	}
L681:
	;
	F_unlink_initfile(m, v2949, int32(15))
	mBase = m.M
	v2958 = m.ExcPending
	if v2958 != 0 {
		goto L32
	} else {
		goto L682
	}
L682:
	;
	F_RelationCacheInitFileRemoveInDir(m, int32(_a_F_StartupXLOG_126))
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L32
	} else {
		goto L683
	}
L683:
	;
	v2963 = F_AllocateDir(m, int32(_a_F_StartupXLOG_38))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L32
	} else {
		goto L684
	}
L684:
	;
	v2967 = F_ReadDirExtended(m, v2963, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L32
	} else {
		goto L685
	}
L685:
	;
	if v2967 != 0 {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v2969 = v2967
	goto L689
L687:
	;
	goto L688
L688:
	;
	F_FreeDir(m, v2963)
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L32
	} else {
		goto L717
	}
L689:
	;
	v3006 = v2969 + int32(19)
	v3007 = int32(_a_F_StartupXLOG_127)
	v3011 = m.G0
	v3013 = v3011 - int32(32)
	v3014 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3013)+24)) = v3014
	*(*int64)(unsafe.Add(mBase, uint32(v3013)+16)) = v3014
	*(*int64)(unsafe.Add(mBase, uint32(v3013)+8)) = v3014
	*(*int64)(unsafe.Add(mBase, uint32(v3013))) = v3014
	v3022 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[56])))
	if v3022 == int32(0) {
		goto L692
	} else {
		goto L693
	}
L690:
	;
	goto L688
L691:
	;
	v3091 = F_strlen(m, v3006)
	mBase = m.M
	if v3090 == v3091 {
		goto L710
	} else {
		goto L711
	}
L692:
	;
	v3090 = int32(0)
	goto L691
L693:
	;
	goto L694
L694:
	;
	v3026 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[57])))
	if v3026 == int32(0) {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v3030 = v3006
	goto L698
L696:
	;
	goto L697
L697:
	;
	v3040 = v3007
	v3041 = v3022
	goto L701
L698:
	;
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3030))))
	if v3036 == v3022 {
		v3030 = v3030 + int32(1)
		goto L698
	} else {
		goto L700
	}
L699:
	;
	v3090 = v3030 - v3006
	goto L691
L700:
	;
	goto L699
L701:
	;
	v3048 = v3013 + int32(base.Ui32(v3041)>>(uint(int32(3))%32))&int32(28)
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v3048)))
	v3050 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3048))) = v3049 | v3050<<(uint(v3041)%32)
	v3054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3040)+1)))
	if v3054 != 0 {
		v3040 = v3040 + v3050
		v3041 = v3054
		goto L701
	} else {
		goto L703
	}
L702:
	;
	v3057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3006))))
	if v3057 == int32(0) {
		v3080 = v3006
		goto L704
	} else {
		goto L705
	}
L703:
	;
	goto L702
L704:
	;
	v3090 = v3080 - v3006
	goto L691
L705:
	;
	v3061 = v3006
	v3062 = v3057
	goto L706
L706:
	;
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v3013+int32(base.Ui32(v3062)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3070)>>(uint(v3062)%32))&int32(1) == int32(0) {
		v3080 = v3061
		goto L704
	} else {
		goto L708
	}
L707:
	;
	v3080 = v3078
	goto L704
L708:
	;
	v3076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+1)))
	v3078 = v3061 + int32(1)
	if v3076 != 0 {
		v3061 = v3078
		v3062 = v3076
		goto L706
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2944)+8)) = int32(_a_F_StartupXLOG_128)
	*(*int32)(unsafe.Add(mBase, uint32(v2944)+4)) = v3006
	*(*int32)(unsafe.Add(mBase, uint32(v2944))) = int32(_a_F_StartupXLOG_38)
	v3099 = v2944 + int32(32)
	v3102 = F_pg_snprintf(m, v3099, int32(1050), int32(_a_F_StartupXLOG_129), v2944)
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L32
	} else {
		goto L713
	}
L711:
	;
	goto L712
L712:
	;
	v3109 = F_ReadDirExtended(m, v2963, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L32
	} else {
		goto L715
	}
L713:
	;
	F_RelationCacheInitFileRemoveInDir(m, v3099)
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L32
	} else {
		goto L714
	}
L714:
	;
	goto L712
L715:
	;
	if v3109 != 0 {
		v2969 = v3109
		goto L689
	} else {
		goto L716
	}
L716:
	;
	goto L690
L717:
	;
	m.G0 = v2944 + int32(1088)
	v3152 = m.G0
	v3154 = v3152 - int32(3696)
	m.G0 = v3154
	v3158 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L32
	} else {
		goto L718
	}
L718:
	;
	if v3158 != 0 {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_130), int32(0))
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L32
	} else {
		goto L722
	}
L720:
	;
	goto L721
L721:
	;
	v3170 = F_AllocateDir(m, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v3171 = m.ExcPending
	if v3171 != 0 {
		goto L32
	} else {
		goto L738
	}
L722:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2405), int32(_a_F_StartupXLOG_133))
	mBase = m.M
	v3168 = m.ExcPending
	if v3168 != 0 {
		goto L32
	} else {
		goto L723
	}
L723:
	;
	goto L721
L724:
	;
	v4034 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	if v4034 != 0 {
		goto L914
	} else {
		goto L915
	}
L725:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L32
	} else {
		goto L910
	}
L726:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L32
	} else {
		goto L905
	}
L727:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3939 = m.ExcPending
	if v3939 != 0 {
		goto L32
	} else {
		goto L900
	}
L728:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3919 = m.ExcPending
	if v3919 != 0 {
		goto L32
	} else {
		goto L897
	}
L729:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3900 = m.ExcPending
	if v3900 != 0 {
		goto L32
	} else {
		goto L893
	}
L730:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3883 = m.ExcPending
	if v3883 != 0 {
		goto L32
	} else {
		goto L890
	}
L731:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L32
	} else {
		goto L886
	}
L732:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L32
	} else {
		goto L882
	}
L733:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3820 = m.ExcPending
	if v3820 != 0 {
		goto L32
	} else {
		goto L878
	}
L734:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3803 = m.ExcPending
	if v3803 != 0 {
		goto L32
	} else {
		goto L875
	}
L735:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L32
	} else {
		goto L871
	}
L736:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3767 = m.ExcPending
	if v3767 != 0 {
		goto L32
	} else {
		goto L867
	}
L737:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L32
	} else {
		goto L863
	}
L738:
	;
	v3173 = F_ReadDir(m, v3170, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v3174 = m.ExcPending
	if v3174 != 0 {
		goto L32
	} else {
		goto L739
	}
L739:
	;
	if v3173 != 0 {
		goto L740
	} else {
		goto L741
	}
L740:
	;
	v3176 = v3154 + int32(3512)
	v3192 = v3173
	goto L743
L741:
	;
	goto L742
L742:
	;
	F_FreeDir(m, v3170)
	mBase = m.M
	v3729 = m.ExcPending
	if v3729 != 0 {
		goto L32
	} else {
		goto L857
	}
L743:
	;
	v3215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3192)+19)))
	if v3215 != int32(46) {
		goto L746
	} else {
		goto L747
	}
L744:
	;
	goto L742
L745:
	;
	v3690 = F_ReadDir(m, v3170, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v3691 = m.ExcPending
	if v3691 != 0 {
		goto L32
	} else {
		goto L855
	}
L746:
	;
	v3228 = v3192 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+340)) = v3228
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+336)) = int32(_a_F_StartupXLOG_131)
	v3233 = v3154 + int32(352)
	v3238 = F_pg_snprintf(m, v3233, int32(1036), int32(_a_F_StartupXLOG_134), v3154+int32(336))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L32
	} else {
		goto L751
	}
L747:
	;
	v3218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3192)+20)))
	if v3218 == int32(0) {
		goto L745
	} else {
		goto L748
	}
L748:
	;
	v3221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3192)+20)))
	if v3221 != int32(46) {
		goto L746
	} else {
		goto L749
	}
L749:
	;
	v3224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3192)+21)))
	if v3224 == int32(0) {
		goto L745
	} else {
		goto L750
	}
L750:
	;
	goto L746
L751:
	;
	v3242 = F_get_dirent_type(m, v3233, v3192, int32(0), int32(14))
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L32
	} else {
		goto L753
	}
L752:
	;
	v3244 = F_strlen(m, v3228)
	mBase = m.M
	v3246 = F_strlen(m, int32(_a_F_StartupXLOG_135))
	mBase = m.M
	if base.Ui32(v3246) <= base.Ui32(v3244) {
		goto L754
	} else {
		goto L755
	}
L753:
	;
	switch v3242 {
	case 0, 3:
		goto L752
	default:
		goto L745
	}
L754:
	;
	v3249 = v3228 + (v3244 - v3246)
	v3250 = int32(_a_F_StartupXLOG_135)
	v3253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3249))))
	v3256 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[59])))
	if base.B2i32(v3253 == int32(0))|base.B2i32(v3253 != v3256) != 0 {
		v3274 = v3253
		v3275 = v3256
		goto L758
	} else {
		goto L759
	}
L755:
	;
	v3278 = int32(1)
	goto L756
L756:
	;
	if v3278 == int32(0) {
		goto L764
	} else {
		goto L765
	}
L757:
	;
	v3278 = v3274 - v3275
	goto L756
L758:
	;
	goto L757
L759:
	;
	v3259 = v3249
	v3260 = v3250
	goto L760
L760:
	;
	v3263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3260)+1)))
	v3264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3259)+1)))
	if v3264 == int32(0) {
		v3274 = v3264
		v3275 = v3263
		goto L758
	} else {
		goto L762
	}
L761:
	;
	v3274 = v3264
	v3275 = v3263
	goto L758
L762:
	;
	v3267 = int32(1)
	if v3264 == v3263 {
		v3259 = v3259 + v3267
		v3260 = v3260 + v3267
		goto L760
	} else {
		goto L763
	}
L763:
	;
	goto L761
L764:
	;
	v3282 = v3154 + int32(352)
	v3283 = F_rmtree(m, v3282)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L32
	} else {
		goto L767
	}
L765:
	;
	goto L766
L766:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+320)) = int32(_a_F_StartupXLOG_131)
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+324)) = v3228
	v3310 = v3154 + int32(2448)
	v3314 = F_pg_sprintf(m, v3310, int32(_a_F_StartupXLOG_134), v3154+int32(320))
	mBase = m.M
	v3315 = m.ExcPending
	if v3315 != 0 {
		goto L32
	} else {
		goto L776
	}
L767:
	;
	if v3283 == int32(0) {
		goto L768
	} else {
		goto L769
	}
L768:
	;
	v3289 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L32
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	F_fsync_fname(m, int32(_a_F_StartupXLOG_131), int32(1))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L32
	} else {
		goto L775
	}
L771:
	;
	if v3289 == int32(0) {
		goto L745
	} else {
		goto L772
	}
L772:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154))) = v3282
	F_errmsg(m, int32(_a_F_StartupXLOG_136), v3154)
	mBase = m.M
	v3296 = m.ExcPending
	if v3296 != 0 {
		goto L32
	} else {
		goto L773
	}
L773:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2432), int32(_a_F_StartupXLOG_133))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L32
	} else {
		goto L774
	}
L774:
	;
	goto L745
L775:
	;
	goto L745
L776:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+304)) = v3310
	v3318 = v3154 + int32(1392)
	v3322 = F_pg_sprintf(m, v3318, int32(_a_F_StartupXLOG_137), v3154+int32(304))
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L32
	} else {
		goto L777
	}
L777:
	;
	v3324 = F_unlink(m, v3318)
	mBase = m.M
	if v3324 < int32(0) {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v3328 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3328 != int32(44) {
		goto L737
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+272)) = v3154 + int32(2448)
	v3335 = v3154 + int32(1392)
	v3339 = F_pg_sprintf(m, v3335, int32(_a_F_StartupXLOG_138), v3154+int32(272))
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L32
	} else {
		goto L782
	}
L781:
	;
	goto L780
L782:
	;
	v3343 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L32
	} else {
		goto L783
	}
L783:
	;
	if v3343 != 0 {
		goto L784
	} else {
		goto L785
	}
L784:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+256)) = v3335
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_139), v3154+int32(256))
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L32
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	v3357 = v3154 + int32(1392)
	v3359 = F_OpenTransientFile(m, v3357, int32(2))
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L32
	} else {
		goto L789
	}
L787:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2709), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3355 = m.ExcPending
	if v3355 != 0 {
		goto L32
	} else {
		goto L788
	}
L788:
	;
	goto L786
L789:
	;
	if v3359 < int32(0) {
		goto L736
	} else {
		goto L790
	}
L790:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3364))) = int32(167772209)
	v3369 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v3369 != int32(1) {
		v3383 = int32(0)
		goto L792
	} else {
		goto L793
	}
L791:
	;
	if v3383 != 0 {
		goto L735
	} else {
		goto L798
	}
L792:
	;
	goto L791
L793:
	;
	goto L794
L794:
	;
	v3374 = F_fsync(m, v3359)
	mBase = m.M
	if v3374 != int32(-1) {
		v3383 = v3374
		goto L792
	} else {
		goto L796
	}
L795:
	;
	v3383 = int32(-1)
	goto L792
L796:
	;
	v3378 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3378 == int32(27) {
		goto L794
	} else {
		goto L797
	}
L797:
	;
	goto L795
L798:
	;
	v3385 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3385))) = int32(0)
	v3388 = int32(_a_F_StartupXLOG_141)
	v3390 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	v3391 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v3390 + v3391
	F_fsync_fname(m, v3154+int32(2448), v3391)
	mBase = m.M
	v3398 = m.ExcPending
	if v3398 != 0 {
		goto L32
	} else {
		goto L799
	}
L799:
	;
	v3399 = int32(_a_F_StartupXLOG_141)
	v3401 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v3401 - int32(1)
	v3405 = int32(_a_F_StartupXLOG_142)
	v3406 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3406))) = int32(167772208)
	v3411 = int32(16)
	v3412 = F_read(m, v3359, v3154+int32(3496), v3411)
	mBase = m.M
	v3414 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3414))) = int32(0)
	if v3412 != v3411 {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L32
	} else {
		goto L803
	}
L801:
	;
	goto L802
L802:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3496))
	if v3442 != int32(17112225) {
		goto L733
	} else {
		goto L808
	}
L803:
	;
	if v3412 < int32(0) {
		goto L734
	} else {
		goto L804
	}
L804:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
		goto L32
	} else {
		goto L805
	}
L805:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+232)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+228)) = v3412
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+224)) = v3357
	F_errmsg(m, int32(_a_F_StartupXLOG_143), v3154+int32(224))
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		goto L32
	} else {
		goto L806
	}
L806:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2755), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L32
	} else {
		goto L807
	}
L807:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L808:
	;
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3504))
	if v3445 != int32(5) {
		goto L732
	} else {
		goto L809
	}
L809:
	;
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3508))
	if v3448 != int32(184) {
		goto L731
	} else {
		goto L810
	}
L810:
	;
	v3451 = int32(_a_F_StartupXLOG_142)
	v3452 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3452))) = int32(167772208)
	v3456 = F_read(m, v3359, v3176, int32(184))
	mBase = m.M
	v3458 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v3458))) = int32(0)
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3508))
	if v3461 != v3456 {
		goto L811
	} else {
		goto L812
	}
L811:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L32
	} else {
		goto L814
	}
L812:
	;
	goto L813
L813:
	;
	v3487 = F_CloseTransientFile(m, v3359)
	mBase = m.M
	v3488 = m.ExcPending
	if v3488 != 0 {
		goto L32
	} else {
		goto L819
	}
L814:
	;
	if v3456 < int32(0) {
		goto L730
	} else {
		goto L815
	}
L815:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L32
	} else {
		goto L816
	}
L816:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+152)) = v3461
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+148)) = v3456
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+144)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_143), v3154+int32(144))
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L32
	} else {
		goto L817
	}
L817:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2795), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L32
	} else {
		goto L818
	}
L818:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L819:
	;
	if v3487 != 0 {
		goto L729
	} else {
		goto L820
	}
L820:
	;
	v3489 = int32(-1)
	v3491 = m.Env.Pgmem_crc32c(m, v3489, v3154+int32(3504), int32(192))
	mBase = m.M
	v3493 = v3491 ^ v3489
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3500))
	if v3493 != v3494 {
		goto L728
	} else {
		goto L821
	}
L821:
	;
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3580))
	if v3496 != 0 {
		goto L822
	} else {
		goto L823
	}
L822:
	;
	v3498 = v3154 + int32(2448)
	v3499 = F_rmtree(m, v3498)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L32
	} else {
		goto L826
	}
L823:
	;
	goto L824
L824:
	;
	v3523 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3576))
	if v3524 != 0 {
		goto L834
	} else {
		goto L835
	}
L825:
	;
	F_fsync_fname(m, int32(_a_F_StartupXLOG_131), int32(1))
	mBase = m.M
	v3521 = m.ExcPending
	if v3521 != 0 {
		goto L32
	} else {
		goto L832
	}
L826:
	;
	if v3499 != 0 {
		goto L825
	} else {
		goto L827
	}
L827:
	;
	v3503 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L32
	} else {
		goto L828
	}
L828:
	;
	if v3503 == int32(0) {
		goto L825
	} else {
		goto L829
	}
L829:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+80)) = v3498
	F_errmsg(m, int32(_a_F_StartupXLOG_136), v3154+int32(80))
	mBase = m.M
	v3512 = m.ExcPending
	if v3512 != 0 {
		goto L32
	} else {
		goto L830
	}
L830:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2825), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L32
	} else {
		goto L831
	}
L831:
	;
	goto L825
L832:
	;
	goto L745
L833:
	;
	v3560 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if v3560 <= int32(0) {
		goto L725
	} else {
		goto L846
	}
L834:
	;
	if v3523 <= int32(0) {
		goto L727
	} else {
		goto L837
	}
L835:
	;
	goto L836
L836:
	;
	if v3523 <= int32(0) {
		goto L726
	} else {
		goto L845
	}
L837:
	;
	v3528 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v3528 != int32(1) {
		goto L833
	} else {
		goto L838
	}
L838:
	;
	v3532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[23])))
	if v3532&int32(1) != 0 {
		goto L833
	} else {
		goto L839
	}
L839:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3538 = m.ExcPending
	if v3538 != 0 {
		goto L32
	} else {
		goto L840
	}
L840:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L32
	} else {
		goto L841
	}
L841:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+64)) = v3176
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v3154-int32(-64))
	mBase = m.M
	v3547 = m.ExcPending
	if v3547 != 0 {
		goto L32
	} else {
		goto L842
	}
L842:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_145), int32(0))
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L32
	} else {
		goto L843
	}
L843:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2865), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L32
	} else {
		goto L844
	}
L844:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L845:
	;
	goto L833
L846:
	;
	v3565 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v3579 = int32(0)
	goto L847
L847:
	;
	v3604 = v3565 + v3579*int32(296)
	v3605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3604)+4)))
	if v3605 != 0 {
		goto L849
	} else {
		goto L850
	}
L848:
	;
	base.MemoryCopy(m, v3604+int32(24), v3176, int32(184))
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3584))
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+16)) = v3613
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3588))
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+20)) = v3615
	v3617 = *(*int64)(unsafe.Add(mBase, uint32(v3154)+3608))
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+264)) = v3617
	v3619 = *(*int64)(unsafe.Add(mBase, uint32(v3154)+3592))
	v3620 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+236)) = v3620
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+280)) = v3619
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+244)) = v3620
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+252)) = v3620
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+260)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+8)) = int32(-1)
	v3631 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3604)+4)) = uint8(v3631)
	v3636 = m.G0
	v3637 = int32(16)
	v3638 = v3636 - v3637
	m.G0 = v3638
	F_gettimeofday(m, v3638)
	mBase = m.M
	v3641 = *(*int64)(unsafe.Add(mBase, uint32(v3638)))
	v3642 = int64(*(*int32)(unsafe.Add(mBase, uint32(v3638)+8)))
	m.G0 = v3638 + v3637
	goto L853
L849:
	;
	v3607 = v3579 + int32(1)
	if v3560 != v3607 {
		v3579 = v3607
		goto L847
	} else {
		goto L852
	}
L850:
	;
	goto L851
L851:
	;
	goto L848
L852:
	;
	goto L725
L853:
	;
	v3651 = *(*int32)(unsafe.Add(mBase, uint32(v3604)+112))
	if v3651 != 0 {
		goto L745
	} else {
		goto L854
	}
L854:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+272)) = v3642 + v3641*int64(1000000) - int64(946684800000000)
	goto L745
L855:
	;
	if v3690 != 0 {
		v3192 = v3690
		goto L743
	} else {
		goto L856
	}
L856:
	;
	goto L744
L857:
	;
	v3731 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v3733 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v3731+v3733 {
		goto L858
	} else {
		goto L859
	}
L858:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L32
	} else {
		goto L861
	}
L859:
	;
	goto L860
L860:
	;
	m.G0 = v3154 + int32(3696)
	goto L724
L861:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v3741 = m.ExcPending
	if v3741 != 0 {
		goto L32
	} else {
		goto L862
	}
L862:
	;
	goto L860
L863:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L32
	} else {
		goto L864
	}
L864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+288)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_146), v3154+int32(288))
	mBase = m.M
	v3758 = m.ExcPending
	if v3758 != 0 {
		goto L32
	} else {
		goto L865
	}
L865:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2705), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3763 = m.ExcPending
	if v3763 != 0 {
		goto L32
	} else {
		goto L866
	}
L866:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L867:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3769 = m.ExcPending
	if v3769 != 0 {
		goto L32
	} else {
		goto L868
	}
L868:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+16)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v3154+int32(16))
	mBase = m.M
	v3777 = m.ExcPending
	if v3777 != 0 {
		goto L32
	} else {
		goto L869
	}
L869:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2721), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3782 = m.ExcPending
	if v3782 != 0 {
		goto L32
	} else {
		goto L870
	}
L870:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L871:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L32
	} else {
		goto L872
	}
L872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+240)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v3154+int32(240))
	mBase = m.M
	v3796 = m.ExcPending
	if v3796 != 0 {
		goto L32
	} else {
		goto L873
	}
L873:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2732), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3801 = m.ExcPending
	if v3801 != 0 {
		goto L32
	} else {
		goto L874
	}
L874:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L875:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+208)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v3154+int32(208))
	mBase = m.M
	v3811 = m.ExcPending
	if v3811 != 0 {
		goto L32
	} else {
		goto L876
	}
L876:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2749), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L32
	} else {
		goto L877
	}
L877:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L878:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3823 = m.ExcPending
	if v3823 != 0 {
		goto L32
	} else {
		goto L879
	}
L879:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+200)) = int32(17112225)
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+196)) = v3442
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+192)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v3154+int32(192))
	mBase = m.M
	v3834 = m.ExcPending
	if v3834 != 0 {
		goto L32
	} else {
		goto L880
	}
L880:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2763), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3839 = m.ExcPending
	if v3839 != 0 {
		goto L32
	} else {
		goto L881
	}
L881:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L882:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3846 = m.ExcPending
	if v3846 != 0 {
		goto L32
	} else {
		goto L883
	}
L883:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+180)) = v3445
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+176)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_150), v3154+int32(176))
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L32
	} else {
		goto L884
	}
L884:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2770), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3860 = m.ExcPending
	if v3860 != 0 {
		goto L32
	} else {
		goto L885
	}
L885:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L886:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L32
	} else {
		goto L887
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+164)) = v3448
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+160)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_151), v3154+int32(160))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L32
	} else {
		goto L888
	}
L888:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2777), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3881 = m.ExcPending
	if v3881 != 0 {
		goto L32
	} else {
		goto L889
	}
L889:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+128)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v3154+int32(128))
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L32
	} else {
		goto L891
	}
L891:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2790), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3896 = m.ExcPending
	if v3896 != 0 {
		goto L32
	} else {
		goto L892
	}
L892:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L893:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3902 = m.ExcPending
	if v3902 != 0 {
		goto L32
	} else {
		goto L894
	}
L894:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+112)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v3154+int32(112))
	mBase = m.M
	v3910 = m.ExcPending
	if v3910 != 0 {
		goto L32
	} else {
		goto L895
	}
L895:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2801), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3915 = m.ExcPending
	if v3915 != 0 {
		goto L32
	} else {
		goto L896
	}
L896:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L897:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+100)) = v3493
	v3921 = *(*int32)(unsafe.Add(mBase, uint32(v3154)+3500))
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+104)) = v3921
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+96)) = v3154 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v3154+int32(96))
	mBase = m.M
	v3930 = m.ExcPending
	if v3930 != 0 {
		goto L32
	} else {
		goto L898
	}
L898:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2813), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3935 = m.ExcPending
	if v3935 != 0 {
		goto L32
	} else {
		goto L899
	}
L899:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L900:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L32
	} else {
		goto L901
	}
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+48)) = v3176
	F_errmsg(m, int32(_a_F_StartupXLOG_154), v3154+int32(48))
	mBase = m.M
	v3948 = m.ExcPending
	if v3948 != 0 {
		goto L32
	} else {
		goto L902
	}
L902:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_155), int32(0))
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		goto L32
	} else {
		goto L903
	}
L903:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2850), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L32
	} else {
		goto L904
	}
L904:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L905:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L32
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3154)+32)) = v3176
	F_errmsg(m, int32(_a_F_StartupXLOG_156), v3154+int32(32))
	mBase = m.M
	v3970 = m.ExcPending
	if v3970 != 0 {
		goto L32
	} else {
		goto L907
	}
L907:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_155), int32(0))
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		goto L32
	} else {
		goto L908
	}
L908:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2872), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v3979 = m.ExcPending
	if v3979 != 0 {
		goto L32
	} else {
		goto L909
	}
L909:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L910:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_157), int32(0))
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L32
	} else {
		goto L911
	}
L911:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_158), int32(0))
	mBase = m.M
	v4027 = m.ExcPending
	if v4027 != 0 {
		goto L32
	} else {
		goto L912
	}
L912:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_132), int32(2926), int32(_a_F_StartupXLOG_140))
	mBase = m.M
	v4032 = m.ExcPending
	if v4032 != 0 {
		goto L32
	} else {
		goto L913
	}
L913:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L914:
	;
	v4035 = int32(_a_F_StartupXLOG_159)
	v4036 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v4038 = v2913 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4036))) = uint8(v4038)
	v4041 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	*(*uint8)(unsafe.Add(mBase, uint32(v4041)+1)) = uint8(v4038)
	goto L916
L915:
	;
	goto L916
L916:
	;
	v4045 = F_AllocateDir(m, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v4046 = m.ExcPending
	if v4046 != 0 {
		goto L32
	} else {
		goto L917
	}
L917:
	;
	v4048 = F_ReadDir(m, v4045, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v4049 = m.ExcPending
	if v4049 != 0 {
		goto L32
	} else {
		goto L918
	}
L918:
	;
	if v4048 != 0 {
		goto L919
	} else {
		goto L920
	}
L919:
	;
	v4058 = v4048
	goto L922
L920:
	;
	goto L921
L921:
	;
	F_FreeDir(m, v4045)
	mBase = m.M
	v4149 = m.ExcPending
	if v4149 != 0 {
		goto L32
	} else {
		goto L935
	}
L922:
	;
	v4086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4058)+19)))
	if v4086 != int32(46) {
		goto L925
	} else {
		goto L926
	}
L923:
	;
	goto L921
L924:
	;
	v4110 = F_ReadDir(m, v4045, int32(_a_F_StartupXLOG_131))
	mBase = m.M
	v4111 = m.ExcPending
	if v4111 != 0 {
		goto L32
	} else {
		goto L933
	}
L925:
	;
	v4099 = v4058 + int32(19)
	v4102 = F_ReplicationSlotValidateName(m, v4099, int32(1), int32(13))
	mBase = m.M
	v4103 = m.ExcPending
	if v4103 != 0 {
		goto L32
	} else {
		goto L930
	}
L926:
	;
	v4089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4058)+20)))
	if v4089 == int32(0) {
		goto L924
	} else {
		goto L927
	}
L927:
	;
	v4092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4058)+20)))
	if v4092 != int32(46) {
		goto L925
	} else {
		goto L928
	}
L928:
	;
	v4095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4058)+21)))
	if v4095 == int32(0) {
		goto L924
	} else {
		goto L929
	}
L929:
	;
	goto L925
L930:
	;
	if v4102 == int32(0) {
		goto L924
	} else {
		goto L931
	}
L931:
	;
	F_ReorderBufferCleanupSerializedTXNs(m, v4099)
	mBase = m.M
	v4107 = m.ExcPending
	if v4107 != 0 {
		goto L32
	} else {
		goto L932
	}
L932:
	;
	goto L924
L933:
	;
	if v4110 != 0 {
		v4058 = v4110
		goto L922
	} else {
		goto L934
	}
L934:
	;
	goto L923
L935:
	;
	v4151 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v4153 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	v4154 = *(*int64)(unsafe.Add(mBase, uint32(v4153)+8))
	v4159 = int32(48)
	v4160 = base.AtomicRmwXchg64(m, v4151, v4159, int64(base.Ui64(v4154)>>(uint(int64(15))%64))&int64(131071))
	v4162 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v4163 = *(*int64)(unsafe.Add(mBase, uint32(v4162)+8))
	v4165 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	v4166 = *(*int32)(unsafe.Add(mBase, uint32(v4162)))
	v4171 = base.AtomicRmwXchg64(m, v4165, v4159, base.I64_extend_i32_u(int32(base.Ui32(v4166)>>(uint(int32(10))%32))))
	v4173 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v4175 = base.I64_div_u_s(v4163, int64(1636))
	v4177 = base.AtomicRmwXchg64(m, v4173, v4159, v4175)
	v4179 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4179)+208)))
	if v4180 == int32(1) {
		goto L936
	} else {
		goto L937
	}
L936:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v4184 = m.ExcPending
	if v4184 != 0 {
		goto L32
	} else {
		goto L939
	}
L937:
	;
	goto L938
L938:
	;
	v4185 = m.G0
	v4187 = v4185 - int32(176)
	m.G0 = v4187
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+172)) = int32(307747550)
	v4192 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	if v4192 == int32(0) {
		goto L946
	} else {
		goto L947
	}
L939:
	;
	goto L938
L940:
	;
	v4571 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v4573 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4574 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+16))
	if v4574 == int32(1) {
		goto L1019
	} else {
		goto L1020
	}
L941:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4555 = m.ExcPending
	if v4555 != 0 {
		goto L32
	} else {
		goto L1016
	}
L942:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4537 = m.ExcPending
	if v4537 != 0 {
		goto L32
	} else {
		goto L1012
	}
L943:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4518 = m.ExcPending
	if v4518 != 0 {
		goto L32
	} else {
		goto L1008
	}
L944:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		goto L32
	} else {
		goto L1004
	}
L945:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4486 = m.ExcPending
	if v4486 != 0 {
		goto L32
	} else {
		goto L1001
	}
L946:
	;
	m.G0 = v4187 + int32(176)
	goto L940
L947:
	;
	v4197 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L32
	} else {
		goto L948
	}
L948:
	;
	if v4197 != 0 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_160), int32(0))
	mBase = m.M
	v4202 = m.ExcPending
	if v4202 != 0 {
		goto L32
	} else {
		goto L952
	}
L950:
	;
	goto L951
L951:
	;
	v4210 = F_OpenTransientFile(m, int32(_a_F_StartupXLOG_161), int32(0))
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L32
	} else {
		goto L954
	}
L952:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(763), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L32
	} else {
		goto L953
	}
L953:
	;
	goto L951
L954:
	;
	if v4210 < int32(0) {
		goto L955
	} else {
		goto L956
	}
L955:
	;
	v4215 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4215 == int32(44) {
		goto L946
	} else {
		goto L958
	}
L956:
	;
	goto L957
L957:
	;
	v4236 = int32(4)
	v4237 = F_read(m, v4210, v4187+int32(172), v4236)
	mBase = m.M
	if v4237 != v4236 {
		goto L963
	} else {
		goto L964
	}
L958:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4221 = m.ExcPending
	if v4221 != 0 {
		goto L32
	} else {
		goto L959
	}
L959:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L32
	} else {
		goto L960
	}
L960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187))) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v4187)
	mBase = m.M
	v4228 = m.ExcPending
	if v4228 != 0 {
		goto L32
	} else {
		goto L961
	}
L961:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(777), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4233 = m.ExcPending
	if v4233 != 0 {
		goto L32
	} else {
		goto L962
	}
L962:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L963:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4243 = m.ExcPending
	if v4243 != 0 {
		goto L32
	} else {
		goto L966
	}
L964:
	;
	goto L965
L965:
	;
	v4268 = m.Env.Pgmem_crc32c(m, int32(-1), v4187+int32(172), int32(4))
	mBase = m.M
	v4269 = *(*int32)(unsafe.Add(mBase, uint32(v4187)+172))
	if v4269 != int32(307747550) {
		goto L941
	} else {
		goto L971
	}
L966:
	;
	if v4237 < int32(0) {
		goto L945
	} else {
		goto L967
	}
L967:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L32
	} else {
		goto L968
	}
L968:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+136)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+132)) = v4237
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+128)) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_143), v4187+int32(128))
	mBase = m.M
	v4258 = m.ExcPending
	if v4258 != 0 {
		goto L32
	} else {
		goto L969
	}
L969:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(792), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4263 = m.ExcPending
	if v4263 != 0 {
		goto L32
	} else {
		goto L970
	}
L970:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L971:
	;
	v4275 = F_read(m, v4210, v4187+int32(152), int32(16))
	mBase = m.M
	if int32(0) <= v4275 {
		goto L973
	} else {
		goto L974
	}
L972:
	;
	v4440 = *(*int32)(unsafe.Add(mBase, uint32(v4187)+152))
	v4442 = v4293 ^ int32(-1)
	if v4440 != v4442 {
		goto L943
	} else {
		goto L998
	}
L973:
	;
	v4279 = int32(0)
	v4292 = v4275
	v4293 = v4268
	goto L976
L974:
	;
	goto L975
L975:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4425 = m.ExcPending
	if v4425 != 0 {
		goto L32
	} else {
		goto L994
	}
L976:
	;
	if v4292 != int32(16) {
		goto L978
	} else {
		goto L979
	}
L977:
	;
	goto L975
L978:
	;
	if v4292 == int32(4) {
		goto L972
	} else {
		goto L981
	}
L979:
	;
	goto L980
L980:
	;
	v4343 = m.Env.Pgmem_crc32c(m, v4293, v4187+int32(152), int32(16))
	mBase = m.M
	v4345 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	if v4279 == v4345 {
		goto L944
	} else {
		goto L986
	}
L981:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v4322 = m.ExcPending
	if v4322 != 0 {
		goto L32
	} else {
		goto L982
	}
L982:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L32
	} else {
		goto L983
	}
L983:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+40)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+36)) = v4292
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+32)) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_143), v4187+int32(32))
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		goto L32
	} else {
		goto L984
	}
L984:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(830), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L32
	} else {
		goto L985
	}
L985:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L986:
	;
	v4348 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[72]))
	v4351 = v4348 + v4279<<(uint(int32(6))%32)
	v4352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4187)+152)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4351))) = uint16(v4352)
	v4354 = *(*int64)(unsafe.Add(mBase, uint32(v4187)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v4351)+8)) = v4354
	v4358 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4359 = m.ExcPending
	if v4359 != 0 {
		goto L32
	} else {
		goto L987
	}
L987:
	;
	if v4358 != 0 {
		goto L988
	} else {
		goto L989
	}
L988:
	;
	v4360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4187)+152)))
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+80)) = v4360
	v4362 = *(*int64)(unsafe.Add(mBase, uint32(v4187)+160))
	*(*uint32)(unsafe.Add(mBase, uint32(v4187)+88)) = uint32(v4362)
	v4365 = int64(base.Ui64(v4362) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4187)+84)) = uint32(v4365)
	F_errmsg(m, int32(_a_F_StartupXLOG_164), v4187+int32(80))
	mBase = m.M
	v4371 = m.ExcPending
	if v4371 != 0 {
		goto L32
	} else {
		goto L991
	}
L989:
	;
	goto L990
L990:
	;
	v4383 = F_read(m, v4210, v4187+int32(152), int32(16))
	mBase = m.M
	if int32(0) <= v4383 {
		v4279 = v4279 + int32(1)
		v4292 = v4383
		v4293 = v4343
		goto L976
	} else {
		goto L993
	}
L991:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(848), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4376 = m.ExcPending
	if v4376 != 0 {
		goto L32
	} else {
		goto L992
	}
L992:
	;
	goto L990
L993:
	;
	goto L977
L994:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L32
	} else {
		goto L995
	}
L995:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+16)) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v4187+int32(16))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L32
	} else {
		goto L996
	}
L996:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(815), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4439 = m.ExcPending
	if v4439 != 0 {
		goto L32
	} else {
		goto L997
	}
L997:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L998:
	;
	v4444 = F_CloseTransientFile(m, v4210)
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L32
	} else {
		goto L999
	}
L999:
	;
	if v4444 != 0 {
		goto L942
	} else {
		goto L1000
	}
L1000:
	;
	goto L946
L1001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+112)) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v4187+int32(112))
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L32
	} else {
		goto L1002
	}
L1002:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(787), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L32
	} else {
		goto L1003
	}
L1003:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1004:
	;
	F_errcode(m, int32(_a_F_StartupXLOG_165))
	mBase = m.M
	v4505 = m.ExcPending
	if v4505 != 0 {
		goto L32
	} else {
		goto L1005
	}
L1005:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_166), int32(0))
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		goto L32
	} else {
		goto L1006
	}
L1006:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(838), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4514 = m.ExcPending
	if v4514 != 0 {
		goto L32
	} else {
		goto L1007
	}
L1007:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1008:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4521 = m.ExcPending
	if v4521 != 0 {
		goto L32
	} else {
		goto L1009
	}
L1009:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+68)) = v4440
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+64)) = v4442
	F_errmsg(m, int32(_a_F_StartupXLOG_167), v4187-int32(-64))
	mBase = m.M
	v4528 = m.ExcPending
	if v4528 != 0 {
		goto L32
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(857), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4533 = m.ExcPending
	if v4533 != 0 {
		goto L32
	} else {
		goto L1011
	}
L1011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1012:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4539 = m.ExcPending
	if v4539 != 0 {
		goto L32
	} else {
		goto L1013
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+48)) = int32(_a_F_StartupXLOG_161)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v4187+int32(48))
	mBase = m.M
	v4546 = m.ExcPending
	if v4546 != 0 {
		goto L32
	} else {
		goto L1014
	}
L1014:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(863), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4551 = m.ExcPending
	if v4551 != 0 {
		goto L32
	} else {
		goto L1015
	}
L1015:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1016:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+100)) = int32(307747550)
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v4187)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+96)) = v4558
	F_errmsg(m, int32(_a_F_StartupXLOG_168), v4187+int32(96))
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L32
	} else {
		goto L1017
	}
L1017:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_162), int32(799), int32(_a_F_StartupXLOG_163))
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L32
	} else {
		goto L1018
	}
L1018:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1019:
	;
	v4577 = *(*int64)(unsafe.Add(mBase, uint32(v4573)+136))
	v4579 = v4577
	goto L1021
L1020:
	;
	v4579 = int64(1000)
	goto L1021
L1021:
	;
	v4581 = base.AtomicRmwXchg64(m, v4571, int32(232), v4579)
	v4583 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	F_restoreTimeLineHistoryFiles(m, v2911, v4583)
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L32
	} else {
		goto L1022
	}
L1022:
	;
	v4587 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v4591 = F_LWLockAcquire(m, v4587+int32(2304), int32(0))
	mBase = m.M
	v4592 = m.ExcPending
	if v4592 != 0 {
		goto L32
	} else {
		goto L1023
	}
L1023:
	;
	v4594 = F_AllocateDir(m, int32(_a_F_StartupXLOG_169))
	mBase = m.M
	v4595 = m.ExcPending
	if v4595 != 0 {
		goto L32
	} else {
		goto L1024
	}
L1024:
	;
	v4597 = F_ReadDir(m, v4594, int32(_a_F_StartupXLOG_169))
	mBase = m.M
	v4598 = m.ExcPending
	if v4598 != 0 {
		goto L32
	} else {
		goto L1025
	}
L1025:
	;
	if v4597 != 0 {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v4607 = v4597
	goto L1029
L1027:
	;
	goto L1028
L1028:
	;
	v4785 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v4785+int32(2304))
	mBase = m.M
	v4789 = m.ExcPending
	if v4789 != 0 {
		goto L32
	} else {
		goto L1059
	}
L1029:
	;
	v4636 = v4607 + int32(19)
	v4637 = F_strlen(m, v4636)
	mBase = m.M
	if v4637 != int32(16) {
		goto L1031
	} else {
		goto L1032
	}
L1030:
	;
	goto L1028
L1031:
	;
	v4746 = F_ReadDir(m, v4594, int32(_a_F_StartupXLOG_169))
	mBase = m.M
	v4747 = m.ExcPending
	if v4747 != 0 {
		goto L32
	} else {
		goto L1057
	}
L1032:
	;
	v4640 = int32(_a_F_StartupXLOG_170)
	v4644 = m.G0
	v4646 = v4644 - int32(32)
	v4647 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4646)+24)) = v4647
	*(*int64)(unsafe.Add(mBase, uint32(v4646)+16)) = v4647
	*(*int64)(unsafe.Add(mBase, uint32(v4646)+8)) = v4647
	*(*int64)(unsafe.Add(mBase, uint32(v4646))) = v4647
	v4655 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[74])))
	if v4655 == int32(0) {
		goto L1034
	} else {
		goto L1035
	}
L1033:
	;
	if v4723 != int32(16) {
		goto L1031
	} else {
		goto L1052
	}
L1034:
	;
	v4723 = int32(0)
	goto L1033
L1035:
	;
	goto L1036
L1036:
	;
	v4659 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[75])))
	if v4659 == int32(0) {
		goto L1037
	} else {
		goto L1038
	}
L1037:
	;
	v4663 = v4636
	goto L1040
L1038:
	;
	goto L1039
L1039:
	;
	v4673 = v4640
	v4674 = v4655
	goto L1043
L1040:
	;
	v4669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4663))))
	if v4669 == v4655 {
		v4663 = v4663 + int32(1)
		goto L1040
	} else {
		goto L1042
	}
L1041:
	;
	v4723 = v4663 - v4636
	goto L1033
L1042:
	;
	goto L1041
L1043:
	;
	v4681 = v4646 + int32(base.Ui32(v4674)>>(uint(int32(3))%32))&int32(28)
	v4682 = *(*int32)(unsafe.Add(mBase, uint32(v4681)))
	v4683 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4681))) = v4682 | v4683<<(uint(v4674)%32)
	v4687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4673)+1)))
	if v4687 != 0 {
		v4673 = v4673 + v4683
		v4674 = v4687
		goto L1043
	} else {
		goto L1045
	}
L1044:
	;
	v4690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4636))))
	if v4690 == int32(0) {
		v4713 = v4636
		goto L1046
	} else {
		goto L1047
	}
L1045:
	;
	goto L1044
L1046:
	;
	v4723 = v4713 - v4636
	goto L1033
L1047:
	;
	v4694 = v4636
	v4695 = v4690
	goto L1048
L1048:
	;
	v4703 = *(*int32)(unsafe.Add(mBase, uint32(v4646+int32(base.Ui32(v4695)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v4703)>>(uint(v4695)%32))&int32(1) == int32(0) {
		v4713 = v4694
		goto L1046
	} else {
		goto L1050
	}
L1049:
	;
	v4713 = v4711
	goto L1046
L1050:
	;
	v4709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4694)+1)))
	v4711 = v4694 + int32(1)
	if v4709 != 0 {
		v4694 = v4711
		v4695 = v4709
		goto L1048
	} else {
		goto L1051
	}
L1051:
	;
	goto L1049
L1052:
	;
	v4729 = F_strtox_2(m, v4636, int32(0), int32(16), int64(-1))
	mBase = m.M
	goto L1053
L1053:
	;
	v4732 = int32(0)
	v4734 = F_ProcessTwoPhaseBuffer(m, v4729, int64(0), int32(1), v4732, v4732)
	mBase = m.M
	v4735 = m.ExcPending
	if v4735 != 0 {
		goto L32
	} else {
		goto L1054
	}
L1054:
	;
	if v4734 == int32(0) {
		goto L1031
	} else {
		goto L1055
	}
L1055:
	;
	v4738 = int64(0)
	F_PrepareRedoAdd(m, v4729, v4734, v4738, v4738, int32(0))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L32
	} else {
		goto L1056
	}
L1056:
	;
	goto L1031
L1057:
	;
	if v4746 != 0 {
		v4607 = v4746
		goto L1029
	} else {
		goto L1058
	}
L1058:
	;
	goto L1030
L1059:
	;
	F_FreeDir(m, v4594)
	mBase = m.M
	v4791 = m.ExcPending
	if v4791 != 0 {
		goto L32
	} else {
		goto L1060
	}
L1060:
	;
	if base.Ui32(v387) <= base.Ui32(int32(-3)) {
		goto L1062
	} else {
		goto L1063
	}
L1061:
	;
	v6271 = int32(1)
	v6272 = v2910 & v6271
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[76])) = uint8(v6272)
	v6275 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	*(*int64)(unsafe.Add(mBase, uint32(v6275)+200)) = v2912
	*(*int64)(unsafe.Add(mBase, uint32(v6275)+152)) = v2912
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[77])) = uint8(v6272)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[78])) = v2912
	v6283 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])))
	if v6283 == v6271 {
		goto L1389
	} else {
		goto L1390
	}
L1062:
	;
	v4794 = m.G0
	v4796 = v4794 - int32(48)
	m.G0 = v4796
	v4800 = F_unlink(m, int32(_a_F_StartupXLOG_171))
	mBase = m.M
	if v4800 != 0 {
		goto L1067
	} else {
		goto L1068
	}
L1063:
	;
	goto L1064
L1064:
	;
	v5020 = m.G0
	v5022 = v5020 - int32(560)
	m.G0 = v5022
	v5025 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	v5028 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v5029 = m.ExcPending
	if v5029 != 0 {
		goto L32
	} else {
		goto L1113
	}
L1065:
	;
	v4862 = int32(1)
	goto L1085
L1066:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), v4849, int32(_a_F_StartupXLOG_173))
	mBase = m.M
	v4852 = m.ExcPending
	if v4852 != 0 {
		goto L32
	} else {
		goto L1084
	}
L1067:
	;
	v4802 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4802 == int32(44) {
		goto L1070
	} else {
		goto L1071
	}
L1068:
	;
	goto L1069
L1069:
	;
	v4837 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4838 = m.ExcPending
	if v4838 != 0 {
		goto L32
	} else {
		goto L1080
	}
L1070:
	;
	v4807 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L32
	} else {
		goto L1073
	}
L1071:
	;
	goto L1072
L1072:
	;
	v4821 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4822 = m.ExcPending
	if v4822 != 0 {
		goto L32
	} else {
		goto L1076
	}
L1073:
	;
	if v4807 == int32(0) {
		goto L1065
	} else {
		goto L1074
	}
L1074:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4796)+16)) = int32(_a_F_StartupXLOG_171)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_174), v4796+int32(16))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L32
	} else {
		goto L1075
	}
L1075:
	;
	v4849 = int32(551)
	goto L1066
L1076:
	;
	if v4821 == int32(0) {
		goto L1065
	} else {
		goto L1077
	}
L1077:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4826 = m.ExcPending
	if v4826 != 0 {
		goto L32
	} else {
		goto L1078
	}
L1078:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4796)+32)) = int32(_a_F_StartupXLOG_171)
	F_errmsg(m, int32(_a_F_StartupXLOG_175), v4796+int32(32))
	mBase = m.M
	v4833 = m.ExcPending
	if v4833 != 0 {
		goto L32
	} else {
		goto L1079
	}
L1079:
	;
	v4849 = int32(556)
	goto L1066
L1080:
	;
	if v4837 == int32(0) {
		goto L1065
	} else {
		goto L1081
	}
L1081:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4842 = m.ExcPending
	if v4842 != 0 {
		goto L32
	} else {
		goto L1082
	}
L1082:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4796))) = int32(_a_F_StartupXLOG_171)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_176), v4796)
	mBase = m.M
	v4847 = m.ExcPending
	if v4847 != 0 {
		goto L32
	} else {
		goto L1083
	}
L1083:
	;
	v4849 = int32(563)
	goto L1066
L1084:
	;
	goto L1065
L1085:
	;
	if base.Ui32(v4862) <= base.Ui32(int32(13)) {
		goto L1089
	} else {
		goto L1090
	}
L1086:
	;
	v4927 = m.G0
	v4928 = int32(16)
	v4929 = v4927 - v4928
	m.G0 = v4929
	F_gettimeofday(m, v4929)
	mBase = m.M
	v4932 = *(*int64)(unsafe.Add(mBase, uint32(v4929)))
	v4933 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4929)+8)))
	m.G0 = v4929 + v4928
	goto L1098
L1087:
	;
	v4921 = v4862 + int32(1)
	if v4921 != int32(33) {
		v4862 = v4921
		goto L1085
	} else {
		goto L1097
	}
L1088:
	;
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v4912)+60))
	if v4913 == int32(0) {
		goto L1087
	} else {
		goto L1095
	}
L1089:
	;
	v4912 = v4862*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1088
L1090:
	;
	goto L1091
L1091:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4862-int32(24)) {
		goto L1087
	} else {
		goto L1092
	}
L1092:
	;
	v4901 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v4901 == int32(0) {
		goto L1087
	} else {
		goto L1093
	}
L1093:
	;
	v4909 = *(*int32)(unsafe.Add(mBase, uint32(v4901+v4862<<(uint(int32(2))%32)-int32(96))))
	if v4909 == int32(0) {
		goto L1087
	} else {
		goto L1094
	}
L1094:
	;
	v4912 = v4909
	goto L1088
L1095:
	;
	m.T0[v4913].(func(*base.Module, int32))(m, int32(2))
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L32
	} else {
		goto L1096
	}
L1096:
	;
	goto L1087
L1097:
	;
	goto L1086
L1098:
	;
	v4951 = int32(1)
	goto L1099
L1099:
	;
	if base.Ui32(v4951) <= base.Ui32(int32(13)) {
		goto L1103
	} else {
		goto L1104
	}
L1100:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5016 = m.ExcPending
	if v5016 != 0 {
		goto L32
	} else {
		goto L1112
	}
L1101:
	;
	v5012 = v4951 + int32(1)
	if v5012 != int32(33) {
		v4951 = v5012
		goto L1099
	} else {
		goto L1111
	}
L1102:
	;
	v5002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5001))))
	if v5002&int32(1) == int32(0) {
		goto L1101
	} else {
		goto L1109
	}
L1103:
	;
	v5001 = v4951*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1102
L1104:
	;
	goto L1105
L1105:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4951-int32(24)) {
		goto L1101
	} else {
		goto L1106
	}
L1106:
	;
	v4990 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v4990 == int32(0) {
		goto L1101
	} else {
		goto L1107
	}
L1107:
	;
	v4998 = *(*int32)(unsafe.Add(mBase, uint32(v4990+v4951<<(uint(int32(2))%32)-int32(96))))
	if v4998 == int32(0) {
		goto L1101
	} else {
		goto L1108
	}
L1108:
	;
	v5001 = v4998
	goto L1102
L1109:
	;
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v5001)+72))
	m.T0[v5007].(func(*base.Module, int64))(m, v4933+v4932*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L32
	} else {
		goto L1110
	}
L1110:
	;
	goto L1101
L1111:
	;
	goto L1100
L1112:
	;
	m.G0 = v4796 + int32(48)
	goto L1061
L1113:
	;
	if v5028 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+464)) = int32(_a_F_StartupXLOG_171)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_178), v5022+int32(464))
	mBase = m.M
	v5036 = m.ExcPending
	if v5036 != 0 {
		goto L32
	} else {
		goto L1117
	}
L1115:
	;
	goto L1116
L1116:
	;
	v5044 = F_AllocateFile(m, int32(_a_F_StartupXLOG_171), int32(_a_F_StartupXLOG_59))
	mBase = m.M
	v5045 = m.ExcPending
	if v5045 != 0 {
		goto L32
	} else {
		goto L1120
	}
L1117:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(1843), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5041 = m.ExcPending
	if v5041 != 0 {
		goto L32
	} else {
		goto L1118
	}
L1118:
	;
	goto L1116
L1119:
	;
	v6160 = int32(1)
	goto L1376
L1120:
	;
	if v5044 == int32(0) {
		goto L1121
	} else {
		goto L1122
	}
L1121:
	;
	v5049 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v5049 == int32(44) {
		goto L1124
	} else {
		goto L1125
	}
L1122:
	;
	goto L1123
L1123:
	;
	v5170 = F_fread(m, v5022+int32(556), int32(1), int32(4), v5044)
	mBase = m.M
	v5171 = m.ExcPending
	if v5171 != 0 {
		goto L32
	} else {
		goto L1148
	}
L1124:
	;
	v5073 = m.G0
	v5074 = int32(16)
	v5075 = v5073 - v5074
	m.G0 = v5075
	F_gettimeofday(m, v5075)
	mBase = m.M
	v5078 = *(*int64)(unsafe.Add(mBase, uint32(v5075)))
	v5079 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5075)+8)))
	m.G0 = v5075 + v5074
	goto L1131
L1125:
	;
	v5054 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L32
	} else {
		goto L1126
	}
L1126:
	;
	if v5054 == int32(0) {
		goto L1124
	} else {
		goto L1127
	}
L1127:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v5059 = m.ExcPending
	if v5059 != 0 {
		goto L32
	} else {
		goto L1128
	}
L1128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022))) = int32(_a_F_StartupXLOG_171)
	F_errmsg(m, int32(_a_F_StartupXLOG_180), v5022)
	mBase = m.M
	v5064 = m.ExcPending
	if v5064 != 0 {
		goto L32
	} else {
		goto L1129
	}
L1129:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(1860), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5069 = m.ExcPending
	if v5069 != 0 {
		goto L32
	} else {
		goto L1130
	}
L1130:
	;
	goto L1124
L1131:
	;
	v5089 = int32(1)
	goto L1132
L1132:
	;
	if base.Ui32(v5089) <= base.Ui32(int32(13)) {
		goto L1136
	} else {
		goto L1137
	}
L1133:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5164 = m.ExcPending
	if v5164 != 0 {
		goto L32
	} else {
		goto L1145
	}
L1134:
	;
	v5160 = v5089 + int32(1)
	if v5160 != int32(33) {
		v5089 = v5160
		goto L1132
	} else {
		goto L1144
	}
L1135:
	;
	v5149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5148))))
	if v5149&int32(1) == int32(0) {
		goto L1134
	} else {
		goto L1142
	}
L1136:
	;
	v5148 = v5089*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1135
L1137:
	;
	goto L1138
L1138:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5089-int32(24)) {
		goto L1134
	} else {
		goto L1139
	}
L1139:
	;
	v5136 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v5136 == int32(0) {
		goto L1134
	} else {
		goto L1140
	}
L1140:
	;
	v5144 = *(*int32)(unsafe.Add(mBase, uint32(v5136+v5089<<(uint(int32(2))%32)-int32(96))))
	if v5144 == int32(0) {
		goto L1134
	} else {
		goto L1141
	}
L1141:
	;
	v5148 = v5144
	goto L1135
L1142:
	;
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(v5148)+72))
	m.T0[v5154].(func(*base.Module, int64))(m, v5079+v5078*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5156 = m.ExcPending
	if v5156 != 0 {
		goto L32
	} else {
		goto L1143
	}
L1143:
	;
	goto L1134
L1144:
	;
	goto L1133
L1145:
	;
	v6136 = int32(2)
	goto L1119
L1146:
	;
	v6103 = F_FreeFile(m, v5044)
	mBase = m.M
	v6104 = m.ExcPending
	if v6104 != 0 {
		goto L32
	} else {
		goto L1369
	}
L1147:
	;
	v5956 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5957 = m.ExcPending
	if v5957 != 0 {
		goto L32
	} else {
		goto L1348
	}
L1148:
	;
	if v5170 != int32(4) {
		goto L1149
	} else {
		goto L1150
	}
L1149:
	;
	v5176 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5177 = m.ExcPending
	if v5177 != 0 {
		goto L32
	} else {
		goto L1152
	}
L1150:
	;
	goto L1151
L1151:
	;
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+556))
	if v5189 == int32(27638972) {
		goto L1156
	} else {
		goto L1157
	}
L1152:
	;
	if v5176 == int32(0) {
		goto L1147
	} else {
		goto L1153
	}
L1153:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_181), int32(0))
	mBase = m.M
	v5183 = m.ExcPending
	if v5183 != 0 {
		goto L32
	} else {
		goto L1154
	}
L1154:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(1871), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5188 = m.ExcPending
	if v5188 != 0 {
		goto L32
	} else {
		goto L1155
	}
L1155:
	;
	goto L1147
L1156:
	;
	goto L1163
L1157:
	;
	goto L1158
L1158:
	;
	v5901 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5902 = m.ExcPending
	if v5902 != 0 {
		goto L32
	} else {
		goto L1344
	}
L1159:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), v5895, int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5898 = m.ExcPending
	if v5898 != 0 {
		goto L32
	} else {
		goto L1343
	}
L1160:
	;
	v5879 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5880 = m.ExcPending
	if v5880 != 0 {
		goto L32
	} else {
		goto L1340
	}
L1161:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), v5865, int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5874 = m.ExcPending
	if v5874 != 0 {
		goto L32
	} else {
		goto L1339
	}
L1162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5849 = m.ExcPending
	if v5849 != 0 {
		goto L32
	} else {
		goto L1336
	}
L1163:
	;
	v5232 = F_do_getc(m, v5044)
	mBase = m.M
	v5233 = m.ExcPending
	if v5233 != 0 {
		goto L32
	} else {
		goto L1172
	}
L1164:
	;
	v5831 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5832 = m.ExcPending
	if v5832 != 0 {
		goto L32
	} else {
		goto L1333
	}
L1165:
	;
	v5674 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[81]))
	v5679 = F_dshash_find_or_insert_extended(m, v5674, v5022+int32(536), v5022+int32(555))
	mBase = m.M
	v5680 = m.ExcPending
	if v5680 != 0 {
		goto L32
	} else {
		goto L1293
	}
L1166:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), v5664, int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5667 = m.ExcPending
	if v5667 != 0 {
		goto L32
	} else {
		goto L1292
	}
L1167:
	;
	v5600 = *(*int32)(unsafe.Add(mBase, uint32(v5599)+48))
	if v5600 == int32(0) {
		goto L1276
	} else {
		goto L1277
	}
L1168:
	;
	v5583 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5584 = m.ExcPending
	if v5584 != 0 {
		goto L32
	} else {
		goto L1272
	}
L1169:
	;
	v5562 = F_do_getc(m, v5044)
	mBase = m.M
	v5563 = m.ExcPending
	if v5563 != 0 {
		goto L32
	} else {
		goto L1266
	}
L1170:
	;
	v5360 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	if v5360 != 0 {
		goto L1206
	} else {
		goto L1207
	}
L1171:
	;
	v5240 = F_fread(m, v5022+int32(468), int32(1), int32(4), v5044)
	mBase = m.M
	v5241 = m.ExcPending
	if v5241 != 0 {
		goto L32
	} else {
		goto L1173
	}
L1172:
	;
	switch v5232 - int32(69) {
	case 0:
		goto L1169
	case 1:
		goto L1171
	default:
		goto L1168
	case 9, 14:
		goto L1170
	}
L1173:
	;
	if v5240 != int32(4) {
		goto L1174
	} else {
		goto L1175
	}
L1174:
	;
	v5246 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5247 = m.ExcPending
	if v5247 != 0 {
		goto L32
	} else {
		goto L1177
	}
L1175:
	;
	goto L1176
L1176:
	;
	v5258 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+468))
	v5260 = v5258 - int32(24)
	v5264 = base.B2i32(base.Ui32(v5258-int32(1)) < base.Ui32(int32(13)))
	if v5264|base.B2i32(base.Ui32(v5260) < base.Ui32(int32(9))) == int32(0) {
		goto L1180
	} else {
		goto L1181
	}
L1177:
	;
	if v5246 == int32(0) {
		goto L1147
	} else {
		goto L1178
	}
L1178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+128)) = int32(70)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_182), v5022+int32(128))
	mBase = m.M
	v5256 = m.ExcPending
	if v5256 != 0 {
		goto L32
	} else {
		goto L1179
	}
L1179:
	;
	v5895 = int32(1901)
	goto L1159
L1180:
	;
	v5272 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5273 = m.ExcPending
	if v5273 != 0 {
		goto L32
	} else {
		goto L1183
	}
L1181:
	;
	goto L1182
L1182:
	;
	if v5264 == int32(0) {
		goto L1188
	} else {
		goto L1189
	}
L1183:
	;
	if v5272 == int32(0) {
		goto L1147
	} else {
		goto L1184
	}
L1184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+116)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+112)) = v5258
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v5022+int32(112))
	mBase = m.M
	v5283 = m.ExcPending
	if v5283 != 0 {
		goto L32
	} else {
		goto L1185
	}
L1185:
	;
	v5895 = int32(1908)
	goto L1159
L1186:
	;
	v5335 = *(*int32)(unsafe.Add(mBase, uint32(v5332)+16))
	v5338 = *(*int32)(unsafe.Add(mBase, uint32(v5332)+20))
	v5339 = F_fread(m, v5334+v5335, int32(1), v5338, v5044)
	mBase = m.M
	v5340 = m.ExcPending
	if v5340 != 0 {
		goto L32
	} else {
		goto L1201
	}
L1187:
	;
	v5331 = *(*int32)(unsafe.Add(mBase, uint32(v5294+(v5025+int32(_a_F_StartupXLOG_184)))))
	v5332 = v5298
	v5334 = v5331
	goto L1186
L1188:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5260) {
		goto L1192
	} else {
		goto L1193
	}
L1189:
	;
	goto L1190
L1190:
	;
	v5320 = v5258 * int32(84)
	v5321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5320)+uint32(_c_F_StartupXLOG[83]))))
	if v5321&int32(1) == int32(0) {
		goto L1160
	} else {
		goto L1200
	}
L1191:
	;
	v5316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5298))))
	if v5316&int32(1) != 0 {
		goto L1187
	} else {
		goto L1199
	}
L1192:
	;
	v5303 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5304 = m.ExcPending
	if v5304 != 0 {
		goto L32
	} else {
		goto L1196
	}
L1193:
	;
	v5290 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v5290 == int32(0) {
		goto L1192
	} else {
		goto L1194
	}
L1194:
	;
	v5294 = v5258 << (uint(int32(2)) % 32)
	v5298 = *(*int32)(unsafe.Add(mBase, uint32(v5290+v5294-int32(96))))
	if v5298 != 0 {
		goto L1191
	} else {
		goto L1195
	}
L1195:
	;
	goto L1192
L1196:
	;
	if v5303 == int32(0) {
		goto L1147
	} else {
		goto L1197
	}
L1197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+100)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+96)) = v5258
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_185), v5022+int32(96))
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L32
	} else {
		goto L1198
	}
L1198:
	;
	v5895 = int32(1916)
	goto L1159
L1199:
	;
	goto L1160
L1200:
	;
	v5328 = *(*int32)(unsafe.Add(mBase, uint32(v5320)+uint32(_c_F_StartupXLOG[84])))
	v5332 = v5320 + int32(_a_F_StartupXLOG_177)
	v5334 = v5025 + v5328
	goto L1186
L1201:
	;
	if v5339 == v5338 {
		goto L1163
	} else {
		goto L1202
	}
L1202:
	;
	v5344 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5345 = m.ExcPending
	if v5345 != 0 {
		goto L32
	} else {
		goto L1203
	}
L1203:
	;
	if v5344 == int32(0) {
		goto L1147
	} else {
		goto L1204
	}
L1204:
	;
	v5348 = *(*int32)(unsafe.Add(mBase, uint32(v5332)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+72)) = v5348
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+68)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+64)) = v5258
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_186), v5022-int32(-64))
	mBase = m.M
	v5357 = m.ExcPending
	if v5357 != 0 {
		goto L32
	} else {
		goto L1205
	}
L1205:
	;
	v5895 = int32(1942)
	goto L1159
L1206:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5362 = m.ExcPending
	if v5362 != 0 {
		goto L32
	} else {
		goto L1209
	}
L1207:
	;
	goto L1208
L1208:
	;
	if v5232 == int32(83) {
		goto L1210
	} else {
		goto L1211
	}
L1209:
	;
	goto L1208
L1210:
	;
	v5369 = F_fread(m, v5022+int32(536), int32(1), int32(16), v5044)
	mBase = m.M
	v5370 = m.ExcPending
	if v5370 != 0 {
		goto L32
	} else {
		goto L1213
	}
L1211:
	;
	goto L1212
L1212:
	;
	v5459 = F_fread(m, v5022+int32(532), int32(1), int32(4), v5044)
	mBase = m.M
	v5460 = m.ExcPending
	if v5460 != 0 {
		goto L32
	} else {
		goto L1236
	}
L1213:
	;
	if v5369 != int32(16) {
		goto L1214
	} else {
		goto L1215
	}
L1214:
	;
	v5375 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5376 = m.ExcPending
	if v5376 != 0 {
		goto L32
	} else {
		goto L1217
	}
L1215:
	;
	goto L1216
L1216:
	;
	v5387 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+536))
	v5389 = v5387 - int32(24)
	v5391 = v5387 - int32(1)
	if base.B2i32(base.Ui32(v5391) < base.Ui32(int32(13)))|base.B2i32(base.Ui32(v5389) < base.Ui32(int32(9))) == int32(0) {
		goto L1220
	} else {
		goto L1221
	}
L1217:
	;
	if v5375 == int32(0) {
		goto L1147
	} else {
		goto L1218
	}
L1218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+336)) = int32(83)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_187), v5022+int32(336))
	mBase = m.M
	v5385 = m.ExcPending
	if v5385 != 0 {
		goto L32
	} else {
		goto L1219
	}
L1219:
	;
	v5865 = int32(1963)
	goto L1161
L1220:
	;
	v5401 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5402 = m.ExcPending
	if v5402 != 0 {
		goto L32
	} else {
		goto L1223
	}
L1221:
	;
	goto L1222
L1222:
	;
	if base.Ui32(v5391) <= base.Ui32(int32(12)) {
		goto L1226
	} else {
		goto L1227
	}
L1223:
	;
	if v5401 == int32(0) {
		goto L1147
	} else {
		goto L1224
	}
L1224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+320)) = int32(83)
	v5407 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+304)) = v5407
	v5409 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+312)) = v5409
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_188), v5022+int32(304))
	mBase = m.M
	v5415 = m.ExcPending
	if v5415 != 0 {
		goto L32
	} else {
		goto L1225
	}
L1225:
	;
	v5865 = int32(1971)
	goto L1161
L1226:
	;
	v5672 = v5387*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1165
L1227:
	;
	goto L1228
L1228:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5389) {
		goto L1229
	} else {
		goto L1230
	}
L1229:
	;
	v5439 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5440 = m.ExcPending
	if v5440 != 0 {
		goto L32
	} else {
		goto L1233
	}
L1230:
	;
	v5426 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v5426 == int32(0) {
		goto L1229
	} else {
		goto L1231
	}
L1231:
	;
	v5434 = *(*int32)(unsafe.Add(mBase, uint32(v5426+v5387<<(uint(int32(2))%32)-int32(96))))
	if v5434 != 0 {
		v5672 = v5434
		goto L1165
	} else {
		goto L1232
	}
L1232:
	;
	goto L1229
L1233:
	;
	if v5439 == int32(0) {
		goto L1147
	} else {
		goto L1234
	}
L1234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+288)) = int32(83)
	v5445 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+272)) = v5445
	v5447 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+280)) = v5447
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_189), v5022+int32(272))
	mBase = m.M
	v5453 = m.ExcPending
	if v5453 != 0 {
		goto L32
	} else {
		goto L1235
	}
L1235:
	;
	v5865 = int32(1980)
	goto L1161
L1236:
	;
	if v5459 != int32(4) {
		goto L1237
	} else {
		goto L1238
	}
L1237:
	;
	v5465 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5466 = m.ExcPending
	if v5466 != 0 {
		goto L32
	} else {
		goto L1240
	}
L1238:
	;
	goto L1239
L1239:
	;
	v5480 = F_fread(m, v5022+int32(468), int32(1), int32(64), v5044)
	mBase = m.M
	v5481 = m.ExcPending
	if v5481 != 0 {
		goto L32
	} else {
		goto L1243
	}
L1240:
	;
	if v5465 == int32(0) {
		goto L1147
	} else {
		goto L1241
	}
L1241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+432)) = v5232
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_182), v5022+int32(432))
	mBase = m.M
	v5474 = m.ExcPending
	if v5474 != 0 {
		goto L32
	} else {
		goto L1242
	}
L1242:
	;
	v5664 = int32(1992)
	goto L1166
L1243:
	;
	if v5480 != int32(64) {
		goto L1244
	} else {
		goto L1245
	}
L1244:
	;
	v5486 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5487 = m.ExcPending
	if v5487 != 0 {
		goto L32
	} else {
		goto L1247
	}
L1245:
	;
	goto L1246
L1246:
	;
	v5499 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+532))
	v5501 = v5499 - int32(24)
	v5503 = v5499 - int32(1)
	if base.B2i32(base.Ui32(v5503) < base.Ui32(int32(13)))|base.B2i32(base.Ui32(v5501) < base.Ui32(int32(9))) == int32(0) {
		goto L1250
	} else {
		goto L1251
	}
L1247:
	;
	if v5486 == int32(0) {
		goto L1147
	} else {
		goto L1248
	}
L1248:
	;
	v5490 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+532))
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+416)) = v5490
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+420)) = v5232
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_190), v5022+int32(416))
	mBase = m.M
	v5497 = m.ExcPending
	if v5497 != 0 {
		goto L32
	} else {
		goto L1249
	}
L1249:
	;
	v5664 = int32(1998)
	goto L1166
L1250:
	;
	v5513 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5514 = m.ExcPending
	if v5514 != 0 {
		goto L32
	} else {
		goto L1253
	}
L1251:
	;
	goto L1252
L1252:
	;
	v5526 = base.B2i32(base.Ui32(int32(12)) < base.Ui32(v5503))
	if v5526 == int32(0) {
		goto L1256
	} else {
		goto L1257
	}
L1253:
	;
	if v5513 == int32(0) {
		goto L1147
	} else {
		goto L1254
	}
L1254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+404)) = v5232
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+400)) = v5499
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v5022+int32(400))
	mBase = m.M
	v5523 = m.ExcPending
	if v5523 != 0 {
		goto L32
	} else {
		goto L1255
	}
L1255:
	;
	v5664 = int32(2004)
	goto L1166
L1256:
	;
	v5599 = v5499*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1167
L1257:
	;
	goto L1258
L1258:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5501) {
		goto L1259
	} else {
		goto L1260
	}
L1259:
	;
	v5549 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5550 = m.ExcPending
	if v5550 != 0 {
		goto L32
	} else {
		goto L1263
	}
L1260:
	;
	v5536 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v5536 == int32(0) {
		goto L1259
	} else {
		goto L1261
	}
L1261:
	;
	v5544 = *(*int32)(unsafe.Add(mBase, uint32(v5536+v5499<<(uint(int32(2))%32)-int32(96))))
	if v5544 != 0 {
		v5599 = v5544
		goto L1167
	} else {
		goto L1262
	}
L1262:
	;
	goto L1259
L1263:
	;
	if v5549 == int32(0) {
		goto L1147
	} else {
		goto L1264
	}
L1264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+388)) = v5232
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+384)) = v5499
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_185), v5022+int32(384))
	mBase = m.M
	v5559 = m.ExcPending
	if v5559 != 0 {
		goto L32
	} else {
		goto L1265
	}
L1265:
	;
	v5664 = int32(2012)
	goto L1166
L1266:
	;
	if v5562 == int32(-1) {
		v6102 = int32(1)
		goto L1146
	} else {
		goto L1267
	}
L1267:
	;
	v5568 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5569 = m.ExcPending
	if v5569 != 0 {
		goto L32
	} else {
		goto L1268
	}
L1268:
	;
	if v5568 == int32(0) {
		goto L1147
	} else {
		goto L1269
	}
L1269:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_191), int32(0))
	mBase = m.M
	v5575 = m.ExcPending
	if v5575 != 0 {
		goto L32
	} else {
		goto L1270
	}
L1270:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(2102), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5580 = m.ExcPending
	if v5580 != 0 {
		goto L32
	} else {
		goto L1271
	}
L1271:
	;
	goto L1147
L1272:
	;
	if v5583 == int32(0) {
		goto L1147
	} else {
		goto L1273
	}
L1273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+48)) = v5232
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_192), v5022+int32(48))
	mBase = m.M
	v5592 = m.ExcPending
	if v5592 != 0 {
		goto L32
	} else {
		goto L1274
	}
L1274:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(2109), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5597 = m.ExcPending
	if v5597 != 0 {
		goto L32
	} else {
		goto L1275
	}
L1275:
	;
	goto L1147
L1276:
	;
	v5605 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5606 = m.ExcPending
	if v5606 != 0 {
		goto L32
	} else {
		goto L1279
	}
L1277:
	;
	goto L1278
L1278:
	;
	v5621 = m.T0[v5600].(func(*base.Module, int32, int32) int32)(m, v5022+int32(468), v5022+int32(536))
	mBase = m.M
	v5622 = m.ExcPending
	if v5622 != 0 {
		goto L32
	} else {
		goto L1282
	}
L1279:
	;
	if v5605 == int32(0) {
		goto L1147
	} else {
		goto L1280
	}
L1280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+356)) = v5232
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+352)) = v5499
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_193), v5022+int32(352))
	mBase = m.M
	v5615 = m.ExcPending
	if v5615 != 0 {
		goto L32
	} else {
		goto L1281
	}
L1281:
	;
	v5664 = int32(2019)
	goto L1166
L1282:
	;
	if v5621 != 0 {
		v5672 = v5599
		goto L1165
	} else {
		goto L1283
	}
L1283:
	;
	if base.Ui32(int32(12)) < base.Ui32(v5503) {
		goto L1284
	} else {
		goto L1285
	}
L1284:
	;
	v5624 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	v5630 = *(*int32)(unsafe.Add(mBase, uint32(v5624+v5499<<(uint(int32(2))%32)-int32(96))))
	v5635 = v5630
	goto L1286
L1285:
	;
	v5635 = v5499*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1286
L1286:
	;
	v5636 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5635)+20)))
	v5638 = F___fseeko_unlocked(m, v5044, v5636, int32(1))
	mBase = m.M
	v5639 = m.ExcPending
	if v5639 != 0 {
		goto L32
	} else {
		goto L1287
	}
L1287:
	;
	if v5638 == int32(0) {
		goto L1163
	} else {
		goto L1288
	}
L1288:
	;
	v5644 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5645 = m.ExcPending
	if v5645 != 0 {
		goto L32
	} else {
		goto L1289
	}
L1289:
	;
	if v5644 == int32(0) {
		goto L1147
	} else {
		goto L1290
	}
L1290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+376)) = v5232
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+372)) = v5499
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+368)) = v5022 + int32(468)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_194), v5022+int32(368))
	mBase = m.M
	v5657 = m.ExcPending
	if v5657 != 0 {
		goto L32
	} else {
		goto L1291
	}
L1291:
	;
	v5664 = int32(2029)
	goto L1166
L1292:
	;
	goto L1147
L1293:
	;
	v5681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5022)+555)))
	if v5681 == int32(1) {
		goto L1294
	} else {
		goto L1295
	}
L1294:
	;
	v5685 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[81]))
	F_dshash_release_lock(m, v5685, v5679)
	mBase = m.M
	v5687 = m.ExcPending
	if v5687 != 0 {
		goto L32
	} else {
		goto L1297
	}
L1295:
	;
	goto L1296
L1296:
	;
	v5706 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+536))
	if base.Ui32(v5706-int32(1)) <= base.Ui32(int32(12)) {
		goto L1302
	} else {
		goto L1303
	}
L1297:
	;
	v5690 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5691 = m.ExcPending
	if v5691 != 0 {
		goto L32
	} else {
		goto L1298
	}
L1298:
	;
	if v5690 == int32(0) {
		goto L1147
	} else {
		goto L1299
	}
L1299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+160)) = v5232
	v5695 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+144)) = v5695
	v5697 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+152)) = v5697
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_195), v5022+int32(144))
	mBase = m.M
	v5703 = m.ExcPending
	if v5703 != 0 {
		goto L32
	} else {
		goto L1300
	}
L1300:
	;
	v5865 = int32(2052)
	goto L1161
L1301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+20)) = int32(1)
	v5738 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5679)+16)) = uint8(v5738)
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+24)) = v5738
	v5743 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[85]))
	v5744 = *(*int32)(unsafe.Add(mBase, uint32(v5735)+4))
	v5746 = F_dsa_allocate_extended(m, v5743, v5744, int32(6))
	mBase = m.M
	v5747 = m.ExcPending
	if v5747 != 0 {
		goto L32
	} else {
		goto L1308
	}
L1302:
	;
	v5735 = v5706*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1301
L1303:
	;
	goto L1304
L1304:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5706-int32(24)) {
		v5733 = int32(0)
		goto L1305
	} else {
		goto L1306
	}
L1305:
	;
	v5735 = v5733
	goto L1301
L1306:
	;
	v5721 = int32(0)
	v5723 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v5723 == v5721 {
		v5733 = v5721
		goto L1305
	} else {
		goto L1307
	}
L1307:
	;
	v5731 = *(*int32)(unsafe.Add(mBase, uint32(v5723+v5706<<(uint(int32(2))%32)-int32(96))))
	v5733 = v5731
	goto L1305
L1308:
	;
	if v5746 != 0 {
		goto L1309
	} else {
		goto L1310
	}
L1309:
	;
	v5749 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[85]))
	v5750 = F_dsa_get_address(m, v5749, v5746)
	mBase = m.M
	v5751 = m.ExcPending
	if v5751 != 0 {
		goto L32
	} else {
		goto L1312
	}
L1310:
	;
	v5773 = int32(0)
	goto L1311
L1311:
	;
	v5775 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[81]))
	F_dshash_release_lock(m, v5775, v5679)
	mBase = m.M
	v5777 = m.ExcPending
	if v5777 != 0 {
		goto L32
	} else {
		goto L1317
	}
L1312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5750))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+28)) = v5746
	v5755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5735))))
	if v5755&int32(8) != 0 {
		goto L1313
	} else {
		goto L1314
	}
L1313:
	;
	v5759 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	v5767 = base.AtomicRmwAdd64(m, v5759+v5706<<(uint(int32(3))%32)+int32(16), int32(0), int64(1))
	goto L1315
L1314:
	;
	goto L1315
L1315:
	;
	F_LWLockInitialize(m, v5750+int32(4), int32(85))
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		goto L32
	} else {
		goto L1316
	}
L1316:
	;
	v5773 = v5750
	goto L1311
L1317:
	;
	if v5773 == int32(0) {
		goto L1162
	} else {
		goto L1318
	}
L1318:
	;
	v5780 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+536))
	if base.Ui32(v5780-int32(1)) <= base.Ui32(int32(12)) {
		goto L1320
	} else {
		goto L1321
	}
L1319:
	;
	v5798 = *(*int32)(unsafe.Add(mBase, uint32(v5797)+16))
	v5801 = *(*int32)(unsafe.Add(mBase, uint32(v5797)+20))
	v5802 = F_fread(m, v5773+v5798, int32(1), v5801, v5044)
	mBase = m.M
	v5803 = m.ExcPending
	if v5803 != 0 {
		goto L32
	} else {
		goto L1323
	}
L1320:
	;
	v5797 = v5780*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1319
L1321:
	;
	goto L1322
L1322:
	;
	v5790 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	v5796 = *(*int32)(unsafe.Add(mBase, uint32(v5790+v5780<<(uint(int32(2))%32)-int32(96))))
	v5797 = v5796
	goto L1319
L1323:
	;
	if v5802 != v5801 {
		goto L1324
	} else {
		goto L1325
	}
L1324:
	;
	v5807 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5808 = m.ExcPending
	if v5808 != 0 {
		goto L32
	} else {
		goto L1327
	}
L1325:
	;
	goto L1326
L1326:
	;
	v5822 = *(*int32)(unsafe.Add(mBase, uint32(v5672)+56))
	if v5822 == int32(0) {
		goto L1163
	} else {
		goto L1330
	}
L1327:
	;
	if v5807 == int32(0) {
		goto L1147
	} else {
		goto L1328
	}
L1328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+256)) = v5232
	v5812 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+240)) = v5812
	v5814 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+248)) = v5814
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_196), v5022+int32(240))
	mBase = m.M
	v5820 = m.ExcPending
	if v5820 != 0 {
		goto L32
	} else {
		goto L1329
	}
L1329:
	;
	v5865 = int32(2076)
	goto L1161
L1330:
	;
	v5827 = m.T0[v5822].(func(*base.Module, int32, int32, int32) int32)(m, v5022+int32(536), v5773, v5044)
	mBase = m.M
	v5828 = m.ExcPending
	if v5828 != 0 {
		goto L32
	} else {
		goto L1331
	}
L1331:
	;
	if v5827 != 0 {
		goto L1163
	} else {
		goto L1332
	}
L1332:
	;
	goto L1164
L1333:
	;
	if v5831 == int32(0) {
		goto L1147
	} else {
		goto L1334
	}
L1334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+224)) = v5232
	v5836 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+208)) = v5836
	v5838 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+216)) = v5838
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_197), v5022+int32(208))
	mBase = m.M
	v5844 = m.ExcPending
	if v5844 != 0 {
		goto L32
	} else {
		goto L1335
	}
L1335:
	;
	v5865 = int32(2087)
	goto L1161
L1336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+192)) = v5232
	v5851 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+176)) = v5851
	v5853 = *(*int64)(unsafe.Add(mBase, uint32(v5022)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+184)) = v5853
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_198), v5022+int32(176))
	mBase = m.M
	v5859 = m.ExcPending
	if v5859 != 0 {
		goto L32
	} else {
		goto L1337
	}
L1337:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(2067), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5864 = m.ExcPending
	if v5864 != 0 {
		goto L32
	} else {
		goto L1338
	}
L1338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1339:
	;
	goto L1147
L1340:
	;
	if v5879 == int32(0) {
		goto L1147
	} else {
		goto L1341
	}
L1341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+84)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+80)) = v5258
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_199), v5022+int32(80))
	mBase = m.M
	v5890 = m.ExcPending
	if v5890 != 0 {
		goto L32
	} else {
		goto L1342
	}
L1342:
	;
	v5895 = int32(1923)
	goto L1159
L1343:
	;
	goto L1147
L1344:
	;
	if v5901 == int32(0) {
		goto L1147
	} else {
		goto L1345
	}
L1345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+452)) = int32(27638972)
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+448)) = v5189
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_200), v5022+int32(448))
	mBase = m.M
	v5912 = m.ExcPending
	if v5912 != 0 {
		goto L32
	} else {
		goto L1346
	}
L1346:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(1878), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5917 = m.ExcPending
	if v5917 != 0 {
		goto L32
	} else {
		goto L1347
	}
L1347:
	;
	goto L1147
L1348:
	;
	if v5956 != 0 {
		goto L1349
	} else {
		goto L1350
	}
L1349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+32)) = int32(_a_F_StartupXLOG_171)
	F_errmsg(m, int32(_a_F_StartupXLOG_201), v5022+int32(32))
	mBase = m.M
	v5964 = m.ExcPending
	if v5964 != 0 {
		goto L32
	} else {
		goto L1352
	}
L1350:
	;
	goto L1351
L1351:
	;
	v5973 = m.G0
	v5974 = int32(16)
	v5975 = v5973 - v5974
	m.G0 = v5975
	F_gettimeofday(m, v5975)
	mBase = m.M
	v5978 = *(*int64)(unsafe.Add(mBase, uint32(v5975)))
	v5979 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5975)+8)))
	m.G0 = v5975 + v5974
	goto L1354
L1352:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(2135), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v5969 = m.ExcPending
	if v5969 != 0 {
		goto L32
	} else {
		goto L1353
	}
L1353:
	;
	goto L1351
L1354:
	;
	v5989 = int32(1)
	goto L1355
L1355:
	;
	if base.Ui32(v5989) <= base.Ui32(int32(13)) {
		goto L1359
	} else {
		goto L1360
	}
L1356:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v6064 = m.ExcPending
	if v6064 != 0 {
		goto L32
	} else {
		goto L1368
	}
L1357:
	;
	v6060 = v5989 + int32(1)
	if v6060 != int32(33) {
		v5989 = v6060
		goto L1355
	} else {
		goto L1367
	}
L1358:
	;
	v6049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6048))))
	if v6049&int32(1) == int32(0) {
		goto L1357
	} else {
		goto L1365
	}
L1359:
	;
	v6048 = v5989*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1358
L1360:
	;
	goto L1361
L1361:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5989-int32(24)) {
		goto L1357
	} else {
		goto L1362
	}
L1362:
	;
	v6036 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v6036 == int32(0) {
		goto L1357
	} else {
		goto L1363
	}
L1363:
	;
	v6044 = *(*int32)(unsafe.Add(mBase, uint32(v6036+v5989<<(uint(int32(2))%32)-int32(96))))
	if v6044 == int32(0) {
		goto L1357
	} else {
		goto L1364
	}
L1364:
	;
	v6048 = v6044
	goto L1358
L1365:
	;
	v6054 = *(*int32)(unsafe.Add(mBase, uint32(v6048)+72))
	m.T0[v6054].(func(*base.Module, int64))(m, v5979+v5978*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v6056 = m.ExcPending
	if v6056 != 0 {
		goto L32
	} else {
		goto L1366
	}
L1366:
	;
	goto L1357
L1367:
	;
	goto L1356
L1368:
	;
	v6102 = int32(2)
	goto L1146
L1369:
	;
	v6107 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6108 = m.ExcPending
	if v6108 != 0 {
		goto L32
	} else {
		goto L1370
	}
L1370:
	;
	if v6107 != 0 {
		goto L1371
	} else {
		goto L1372
	}
L1371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5022)+16)) = int32(_a_F_StartupXLOG_171)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_202), v5022+int32(16))
	mBase = m.M
	v6115 = m.ExcPending
	if v6115 != 0 {
		goto L32
	} else {
		goto L1374
	}
L1372:
	;
	goto L1373
L1373:
	;
	v6122 = F_unlink(m, int32(_a_F_StartupXLOG_171))
	mBase = m.M
	v6136 = v6102
	goto L1119
L1374:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_172), int32(2118), int32(_a_F_StartupXLOG_179))
	mBase = m.M
	v6120 = m.ExcPending
	if v6120 != 0 {
		goto L32
	} else {
		goto L1375
	}
L1375:
	;
	goto L1373
L1376:
	;
	if base.Ui32(v6160) <= base.Ui32(int32(13)) {
		goto L1380
	} else {
		goto L1381
	}
L1377:
	;
	m.G0 = v5022 + int32(560)
	goto L1061
L1378:
	;
	v6228 = v6160 + int32(1)
	if v6228 != int32(33) {
		v6160 = v6228
		goto L1376
	} else {
		goto L1388
	}
L1379:
	;
	v6220 = *(*int32)(unsafe.Add(mBase, uint32(v6219)+60))
	if v6220 == int32(0) {
		goto L1378
	} else {
		goto L1386
	}
L1380:
	;
	v6219 = v6160*int32(84) + int32(_a_F_StartupXLOG_177)
	goto L1379
L1381:
	;
	goto L1382
L1382:
	;
	if base.Ui32(int32(8)) < base.Ui32(v6160-int32(24)) {
		goto L1378
	} else {
		goto L1383
	}
L1383:
	;
	v6207 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[80]))
	if v6207 == int32(0) {
		goto L1378
	} else {
		goto L1384
	}
L1384:
	;
	v6215 = *(*int32)(unsafe.Add(mBase, uint32(v6207+v6160<<(uint(int32(2))%32)-int32(96))))
	if v6215 == int32(0) {
		goto L1378
	} else {
		goto L1385
	}
L1385:
	;
	v6219 = v6215
	goto L1379
L1386:
	;
	m.T0[v6220].(func(*base.Module, int32))(m, v6136)
	mBase = m.M
	v6224 = m.ExcPending
	if v6224 != 0 {
		goto L32
	} else {
		goto L1387
	}
L1387:
	;
	goto L1378
L1388:
	;
	goto L1377
L1389:
	;
	v6288 = base.AtomicRmwXchg32(m, v6275, int32(440), int32(1))
	if v6288 != 0 {
		goto L1392
	} else {
		goto L1393
	}
L1390:
	;
	goto L1391
L1391:
	;
	v9545 = m.G0
	v9547 = v9545 - int32(272)
	m.G0 = v9547
	v9550 = F_palloc(m, int32(72))
	mBase = m.M
	v9551 = m.ExcPending
	if v9551 != 0 {
		goto L32
	} else {
		goto L2060
	}
L1392:
	;
	F_s_lock(m, v6275+int32(440), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v6293 = m.ExcPending
	if v6293 != 0 {
		goto L32
	} else {
		goto L1395
	}
L1393:
	;
	goto L1394
L1394:
	;
	v6295 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v6297 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	*(*int32)(unsafe.Add(mBase, uint32(v6295)+308)) = v6297
	v6299 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6295)+440)), uint32(v6299))
	v6303 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v6305 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v6303, v6305)
	mBase = m.M
	v6307 = m.ExcPending
	if v6307 != 0 {
		goto L32
	} else {
		goto L1396
	}
L1395:
	;
	goto L1394
L1396:
	;
	v6308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4093)))
	if v6308 == int32(1) {
		goto L1397
	} else {
		goto L1398
	}
L1397:
	;
	v6311 = int32(_a_F_StartupXLOG_204)
	v6312 = F_unlink(m, v6311)
	mBase = m.M
	v6316 = F_durable_rename(m, int32(_a_F_StartupXLOG_58), v6311, int32(22))
	mBase = m.M
	v6317 = m.ExcPending
	if v6317 != 0 {
		goto L32
	} else {
		goto L1400
	}
L1398:
	;
	goto L1399
L1399:
	;
	v6318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4094)))
	if v6318 == int32(1) {
		goto L1401
	} else {
		goto L1402
	}
L1400:
	;
	goto L1399
L1401:
	;
	v6321 = int32(_a_F_StartupXLOG_91)
	v6322 = F_unlink(m, v6321)
	mBase = m.M
	v6326 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v6321, int32(22))
	mBase = m.M
	v6327 = m.ExcPending
	if v6327 != 0 {
		goto L32
	} else {
		goto L1404
	}
L1402:
	;
	goto L1403
L1403:
	;
	v6330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v6330 == int32(1) {
		goto L1405
	} else {
		goto L1406
	}
L1404:
	;
	goto L1403
L1405:
	;
	v6334 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v6335 = *(*int64)(unsafe.Add(mBase, uint32(v6334)+144))
	v6337 = v6335
	goto L1407
L1406:
	;
	v6337 = int64(0)
	goto L1407
L1407:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[86])) = v6337
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v6340 = m.ExcPending
	if v6340 != 0 {
		goto L32
	} else {
		goto L1408
	}
L1408:
	;
	F_ResetUnloggedRelations(m, int32(1))
	mBase = m.M
	v6343 = m.ExcPending
	if v6343 != 0 {
		goto L32
	} else {
		goto L1409
	}
L1409:
	;
	v6344 = m.G0
	v6346 = v6344 - int32(1072)
	m.G0 = v6346
	v6349 = F_AllocateDir(m, int32(_a_F_StartupXLOG_205))
	mBase = m.M
	v6350 = m.ExcPending
	if v6350 != 0 {
		goto L32
	} else {
		goto L1410
	}
L1410:
	;
	v6353 = F_ReadDirExtended(m, v6349, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6354 = m.ExcPending
	if v6354 != 0 {
		goto L32
	} else {
		goto L1411
	}
L1411:
	;
	if v6353 != 0 {
		goto L1412
	} else {
		goto L1413
	}
L1412:
	;
	v6355 = v6353
	goto L1415
L1413:
	;
	goto L1414
L1414:
	;
	F_FreeDir(m, v6349)
	mBase = m.M
	v6476 = m.ExcPending
	if v6476 != 0 {
		goto L32
	} else {
		goto L1432
	}
L1415:
	;
	v6391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6355)+19)))
	if v6391 != int32(46) {
		goto L1418
	} else {
		goto L1419
	}
L1416:
	;
	goto L1414
L1417:
	;
	v6437 = F_ReadDirExtended(m, v6349, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6438 = m.ExcPending
	if v6438 != 0 {
		goto L32
	} else {
		goto L1430
	}
L1418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6346)+16)) = v6355 + int32(19)
	v6407 = v6346 + int32(32)
	v6412 = F_pg_snprintf(m, v6407, int32(1037), int32(_a_F_StartupXLOG_206), v6346+int32(16))
	mBase = m.M
	v6413 = m.ExcPending
	if v6413 != 0 {
		goto L32
	} else {
		goto L1423
	}
L1419:
	;
	v6394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6355)+20)))
	if v6394 == int32(0) {
		goto L1417
	} else {
		goto L1420
	}
L1420:
	;
	v6397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6355)+20)))
	if v6397 != int32(46) {
		goto L1418
	} else {
		goto L1421
	}
L1421:
	;
	v6400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6355)+21)))
	if v6400 == int32(0) {
		goto L1417
	} else {
		goto L1422
	}
L1422:
	;
	goto L1418
L1423:
	;
	v6414 = F_unlink(m, v6407)
	mBase = m.M
	if v6414 == int32(0) {
		goto L1417
	} else {
		goto L1424
	}
L1424:
	;
	v6419 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6420 = m.ExcPending
	if v6420 != 0 {
		goto L32
	} else {
		goto L1425
	}
L1425:
	;
	if v6419 == int32(0) {
		goto L1417
	} else {
		goto L1426
	}
L1426:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v6424 = m.ExcPending
	if v6424 != 0 {
		goto L32
	} else {
		goto L1427
	}
L1427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6346))) = v6407
	F_errmsg(m, int32(_a_F_StartupXLOG_146), v6346)
	mBase = m.M
	v6428 = m.ExcPending
	if v6428 != 0 {
		goto L32
	} else {
		goto L1428
	}
L1428:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_207), int32(1646), int32(_a_F_StartupXLOG_208))
	mBase = m.M
	v6433 = m.ExcPending
	if v6433 != 0 {
		goto L32
	} else {
		goto L1429
	}
L1429:
	;
	goto L1417
L1430:
	;
	if v6437 != 0 {
		v6355 = v6437
		goto L1415
	} else {
		goto L1431
	}
L1431:
	;
	goto L1416
L1432:
	;
	m.G0 = v6346 + int32(1072)
	v6481 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v6481 != int32(1) {
		goto L1433
	} else {
		goto L1434
	}
L1433:
	;
	v6729 = m.G0
	v6731 = v6729 - int32(848)
	m.G0 = v6731
	v6734 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6737 = base.AtomicRmwXchg32(m, v6734, int32(96), int32(1))
	if v6737 != 0 {
		goto L1468
	} else {
		goto L1469
	}
L1434:
	;
	v6485 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[23])))
	if v6485&int32(1) == int32(0) {
		goto L1433
	} else {
		goto L1435
	}
L1435:
	;
	v6492 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v6493 = m.ExcPending
	if v6493 != 0 {
		goto L32
	} else {
		goto L1436
	}
L1436:
	;
	if v6492 != 0 {
		goto L1437
	} else {
		goto L1438
	}
L1437:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_209), int32(0))
	mBase = m.M
	v6497 = m.ExcPending
	if v6497 != 0 {
		goto L32
	} else {
		goto L1440
	}
L1438:
	;
	goto L1439
L1439:
	;
	v6503 = m.G0
	v6505 = v6503 + int32(-64)
	m.G0 = v6505
	*(*int64)(unsafe.Add(mBase, uint32(v6505)+16)) = int64(68719476748)
	v6513 = v6503 + int32(-56)
	v6515 = F_hash_create(m, int32(_a_F_StartupXLOG_210), int64(64), v6513, int32(40))
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L32
	} else {
		goto L1442
	}
L1440:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_211), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v6502 = m.ExcPending
	if v6502 != 0 {
		goto L32
	} else {
		goto L1441
	}
L1441:
	;
	goto L1439
L1442:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[87])) = v6515
	*(*int64)(unsafe.Add(mBase, uint32(v6505)+16)) = int64(34359738372)
	v6524 = F_hash_create(m, int32(_a_F_StartupXLOG_212), int64(64), v6513, int32(40))
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L32
	} else {
		goto L1443
	}
L1443:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[88])) = v6524
	F_SharedInvalBackendInit(m, int32(1))
	mBase = m.M
	v6529 = m.ExcPending
	if v6529 != 0 {
		goto L32
	} else {
		goto L1444
	}
L1444:
	;
	v6531 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	v6533 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v6531)+40)) = v6533
	*(*int32)(unsafe.Add(mBase, uint32(v6505)+56)) = v6533
	v6537 = int32(_a_F_StartupXLOG_213)
	v6538 = int32(1)
	v6540 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[91]))
	if base.Ui32(v6540) <= base.Ui32(v6538) {
		goto L1446
	} else {
		goto L1447
	}
L1445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6505)+60)) = v6543
	v6548 = *(*int64)(unsafe.Add(mBase, uint32(v6505)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v6505))) = v6548
	F_VirtualXactLockTableInsert(m, v6505)
	mBase = m.M
	v6551 = m.ExcPending
	if v6551 != 0 {
		goto L32
	} else {
		goto L1449
	}
L1446:
	;
	v6543 = v6538
	goto L1448
L1447:
	;
	v6543 = v6540
	goto L1448
L1448:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[91])) = v6543 + int32(1)
	goto L1445
L1449:
	;
	v6553 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[92])) = v6553
	m.G0 = v6505 - int32(-64)
	v6558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4095)))
	if v6558 == v6553 {
		goto L1450
	} else {
		goto L1451
	}
L1450:
	;
	v6565 = F_PrescanPreparedTransactions(m, v41+int32(_a_F_StartupXLOG_4), v41+int32(_a_F_StartupXLOG_3))
	mBase = m.M
	v6566 = m.ExcPending
	if v6566 != 0 {
		goto L32
	} else {
		goto L1453
	}
L1451:
	;
	v6567 = v2909
	goto L1452
L1452:
	;
	v6569 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	v6570 = *(*int32)(unsafe.Add(mBase, uint32(v6569)+8))
	v6579 = v6570
	goto L1454
L1453:
	;
	v6567 = v6565
	goto L1452
L1454:
	;
	v6608 = v6579 - int32(1)
	if base.Ui32(v6608) < base.Ui32(int32(3)) {
		v6579 = v6608
		goto L1454
	} else {
		goto L1456
	}
L1455:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[93])) = v6608
	F_StartupSUBTRANS(m, v6567)
	mBase = m.M
	v6614 = m.ExcPending
	if v6614 != 0 {
		goto L32
	} else {
		goto L1457
	}
L1456:
	;
	goto L1455
L1457:
	;
	v6616 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6619 = base.AtomicRmwXchg32(m, v6616, int32(96), int32(1))
	if v6619 != 0 {
		goto L1458
	} else {
		goto L1459
	}
L1458:
	;
	F_s_lock(m, v6616+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v6624 = m.ExcPending
	if v6624 != 0 {
		goto L32
	} else {
		goto L1461
	}
L1459:
	;
	goto L1460
L1460:
	;
	v6625 = int32(_a_F_StartupXLOG_214)
	v6626 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6626)+2)) = uint8(v6627)
	v6630 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6631 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6630)+96)), uint32(v6631))
	v6634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4095)))
	if v6634 != v6627 {
		goto L1433
	} else {
		goto L1462
	}
L1461:
	;
	goto L1460
L1462:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v6638 = m.ExcPending
	if v6638 != 0 {
		goto L32
	} else {
		goto L1463
	}
L1463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[94]))) = v6567
	v6640 = base.I32_wrap_i64(v2922)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[95]))) = v6640
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[96]))) = int64(8589934592)
	v6644 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[97])))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v6644
	v6654 = v6640
	goto L1464
L1464:
	;
	v6683 = v6654 - int32(1)
	if base.Ui32(v6683) < base.Ui32(int32(3)) {
		v6654 = v6683
		goto L1464
	} else {
		goto L1466
	}
L1465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[98]))) = v6683
	v6687 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[99])))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[100]))) = v6687
	F_ProcArrayApplyRecoveryInfo(m, v41+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v6692 = m.ExcPending
	if v6692 != 0 {
		goto L32
	} else {
		goto L1467
	}
L1466:
	;
	goto L1465
L1467:
	;
	goto L1433
L1468:
	;
	F_s_lock(m, v6734+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v6742 = m.ExcPending
	if v6742 != 0 {
		goto L32
	} else {
		goto L1471
	}
L1469:
	;
	goto L1470
L1470:
	;
	v6744 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[40]))
	v6746 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if base.Ui64(v6744) < base.Ui64(v6746) {
		goto L1473
	} else {
		goto L1474
	}
L1471:
	;
	goto L1470
L1472:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6762)+32)) = v6763
	v6766 = *(*int32)(unsafe.Add(mBase, uint32(v6764)))
	v6767 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6762)+64)) = v6767
	*(*int32)(unsafe.Add(mBase, uint32(v6762)+56)) = v6766
	*(*int64)(unsafe.Add(mBase, uint32(v6762)+48)) = v6763
	*(*int32)(unsafe.Add(mBase, uint32(v6762)+40)) = v6766
	*(*int64)(unsafe.Add(mBase, uint32(v6762)+72)) = v6767
	v6774 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6762)+80)) = v6774
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6762)+96)), uint32(v6774))
	v6783 = m.G0
	v6784 = int32(16)
	v6785 = v6783 - v6784
	m.G0 = v6785
	F_gettimeofday(m, v6785)
	mBase = m.M
	v6788 = *(*int64)(unsafe.Add(mBase, uint32(v6785)))
	v6789 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6785)+8)))
	m.G0 = v6785 + v6784
	goto L1476
L1473:
	;
	v6749 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int64)(unsafe.Add(mBase, uint32(v6749)+24)) = int64(0)
	v6762 = v6749
	v6763 = v6744
	v6764 = int32(_a_F_StartupXLOG_215)
	goto L1472
L1474:
	;
	goto L1475
L1475:
	;
	v6754 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6756 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v6757 = *(*int64)(unsafe.Add(mBase, uint32(v6756)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v6754)+24)) = v6757
	v6759 = *(*int64)(unsafe.Add(mBase, uint32(v6756)+40))
	v6762 = v6754
	v6763 = v6759
	v6764 = int32(_a_F_StartupXLOG_216)
	goto L1472
L1476:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v6789 + v6788*int64(1000000) - int64(946684800000000)
	v6800 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[18])))
	if v6800 == int32(1) {
		goto L1477
	} else {
		goto L1478
	}
L1477:
	;
	v6805 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[18])))
	if v6805 == int32(1) {
		goto L1481
	} else {
		goto L1482
	}
L1478:
	;
	goto L1479
L1479:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v6820 = m.ExcPending
	if v6820 != 0 {
		goto L32
	} else {
		goto L1484
	}
L1480:
	;
	goto L1479
L1481:
	;
	v6809 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	*(*int32)(unsafe.Add(mBase, uint32(v6809+int32(0)))) = int32(1)
	v6816 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v6818 = F_pgmem_kill(m, v6816, int32(10))
	mBase = m.M
	goto L1483
L1482:
	;
	goto L1483
L1483:
	;
	goto L1480
L1484:
	;
	v6822 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v6824 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[40]))
	v6826 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if base.Ui64(v6824) < base.Ui64(v6826) {
		goto L1494
	} else {
		goto L1495
	}
L1485:
	;
	goto L1391
L1486:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9495 = m.ExcPending
	if v9495 != 0 {
		goto L32
	} else {
		goto L2056
	}
L1487:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9482 = m.ExcPending
	if v9482 != 0 {
		goto L32
	} else {
		goto L2053
	}
L1488:
	;
	if v9444 != 0 {
		goto L2049
	} else {
		goto L2050
	}
L1489:
	;
	v9313 = int32(0)
	goto L2020
L1490:
	;
	v9236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v9236 == int32(0) {
		goto L1487
	} else {
		goto L2008
	}
L1491:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v7419
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = v9140
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = int64(0)
	v9150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v9150)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v9150)
	v9157 = F_errstart(m, int32(15), v9150)
	mBase = m.M
	v9158 = m.ExcPending
	if v9158 != 0 {
		goto L32
	} else {
		goto L1996
	}
L1492:
	;
	v9126 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9127 = m.ExcPending
	if v9127 != 0 {
		goto L32
	} else {
		goto L1992
	}
L1493:
	;
	F_getrusage(m, v6731+int32(384))
	mBase = m.M
	F_gettimeofday(m, v6731+int32(368))
	mBase = m.M
	goto L1508
L1494:
	;
	v6829 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	F_XLogPrefetcherBeginRead(m, v6822, v6824)
	mBase = m.M
	v6831 = m.ExcPending
	if v6831 != 0 {
		goto L32
	} else {
		goto L1497
	}
L1495:
	;
	goto L1496
L1496:
	;
	v6867 = int32(0)
	v6871 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[37]))
	v6872 = F_ReadRecord(m, v6822, int32(15), v6867, v6871)
	mBase = m.M
	v6873 = m.ExcPending
	if v6873 != 0 {
		goto L32
	} else {
		goto L1506
	}
L1497:
	;
	v6833 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v6836 = F_ReadRecord(m, v6833, int32(24), int32(0), v6829)
	mBase = m.M
	v6837 = m.ExcPending
	if v6837 != 0 {
		goto L32
	} else {
		goto L1498
	}
L1498:
	;
	v6838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6836)+17)))
	if v6838 == int32(0) {
		goto L1499
	} else {
		goto L1500
	}
L1499:
	;
	v6841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6836)+16)))
	if v6841&int32(240) == int32(224) {
		v6876 = v6829
		v6877 = v6836
		goto L1493
	} else {
		goto L1502
	}
L1500:
	;
	goto L1501
L1501:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6849 = m.ExcPending
	if v6849 != 0 {
		goto L32
	} else {
		goto L1503
	}
L1502:
	;
	goto L1501
L1503:
	;
	v6851 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v6852 = *(*int64)(unsafe.Add(mBase, uint32(v6851)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+356)) = uint32(v6852)
	v6855 = int64(base.Ui64(v6852) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+352)) = uint32(v6855)
	F_errmsg(m, int32(_a_F_StartupXLOG_217), v6731+int32(352))
	mBase = m.M
	v6861 = m.ExcPending
	if v6861 != 0 {
		goto L32
	} else {
		goto L1504
	}
L1504:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1687), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v6866 = m.ExcPending
	if v6866 != 0 {
		goto L32
	} else {
		goto L1505
	}
L1505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1506:
	;
	if v6872 == int32(0) {
		goto L1492
	} else {
		goto L1507
	}
L1507:
	;
	v6876 = v6871
	v6877 = v6872
	goto L1493
L1508:
	;
	v6886 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[109])) = uint8(v6886)
	v6892 = int32(0)
	goto L1509
L1509:
	;
	v6926 = v6892 << (uint(int32(5)) % 32)
	v6929 = *(*int32)(unsafe.Add(mBase, uint32(v6926)+uint32(_c_F_StartupXLOG[110])))
	if v6929 == int32(0) {
		goto L1511
	} else {
		goto L1512
	}
L1510:
	;
	v6953 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6954 = m.ExcPending
	if v6954 != 0 {
		goto L32
	} else {
		goto L1520
	}
L1511:
	;
	v6938 = *(*int32)(unsafe.Add(mBase, uint32(v6926)+uint32(_c_F_StartupXLOG[111])))
	if v6938 == int32(0) {
		goto L1515
	} else {
		goto L1516
	}
L1512:
	;
	v6932 = *(*int32)(unsafe.Add(mBase, uint32(v6926)+uint32(_c_F_StartupXLOG[112])))
	if v6932 == int32(0) {
		goto L1511
	} else {
		goto L1513
	}
L1513:
	;
	m.T0[v6932].(func(*base.Module))(m)
	mBase = m.M
	v6936 = m.ExcPending
	if v6936 != 0 {
		goto L32
	} else {
		goto L1514
	}
L1514:
	;
	goto L1511
L1515:
	;
	v6948 = v6892 + int32(2)
	if v6948 != int32(256) {
		v6892 = v6948
		goto L1509
	} else {
		goto L1519
	}
L1516:
	;
	v6941 = *(*int32)(unsafe.Add(mBase, uint32(v6926)+uint32(_c_F_StartupXLOG[113])))
	if v6941 == int32(0) {
		goto L1515
	} else {
		goto L1517
	}
L1517:
	;
	m.T0[v6941].(func(*base.Module))(m)
	mBase = m.M
	v6945 = m.ExcPending
	if v6945 != 0 {
		goto L32
	} else {
		goto L1518
	}
L1518:
	;
	goto L1515
L1519:
	;
	goto L1510
L1520:
	;
	if v6953 != 0 {
		goto L1521
	} else {
		goto L1522
	}
L1521:
	;
	v6956 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v6957 = *(*int64)(unsafe.Add(mBase, uint32(v6956)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+340)) = uint32(v6957)
	v6960 = int64(base.Ui64(v6957) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+336)) = uint32(v6960)
	F_errmsg(m, int32(_a_F_StartupXLOG_219), v6731+int32(336))
	mBase = m.M
	v6966 = m.ExcPending
	if v6966 != 0 {
		goto L32
	} else {
		goto L1524
	}
L1522:
	;
	goto L1523
L1523:
	;
	v6974 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v6974 == int32(0) {
		goto L1526
	} else {
		goto L1527
	}
L1524:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1710), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v6971 = m.ExcPending
	if v6971 != 0 {
		goto L32
	} else {
		goto L1525
	}
L1525:
	;
	goto L1523
L1526:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v6978 = m.ExcPending
	if v6978 != 0 {
		goto L32
	} else {
		goto L1529
	}
L1527:
	;
	goto L1528
L1528:
	;
	v6979 = v6876
	v6990 = v6877
	goto L1530
L1529:
	;
	goto L1528
L1530:
	;
	v7016 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v7016 != 0 {
		goto L1532
	} else {
		goto L1533
	}
L1531:
	;
	v9280 = v9117
	goto L1489
L1532:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7084 = m.ExcPending
	if v7084 != 0 {
		goto L32
	} else {
		goto L1543
	}
L1533:
	;
	v7024 = m.G0
	v7026 = v7024 - int32(16)
	m.G0 = v7026
	v7029 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114]))
	if v7029 != 0 {
		goto L1535
	} else {
		goto L1536
	}
L1534:
	;
	if base.B2i32(v7029 != int32(0)) == int32(0) {
		goto L1532
	} else {
		goto L1538
	}
L1535:
	;
	v7030 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v7032 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[115]))
	F_TimestampDifference(m, v7032, v7030, v7026+int32(12), v7026+int32(8))
	mBase = m.M
	v7038 = *(*int32)(unsafe.Add(mBase, uint32(v7026)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6731+int32(560)))) = v7038
	v7040 = *(*int32)(unsafe.Add(mBase, uint32(v7026)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6731+int32(540)))) = v7040
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114])) = int32(0)
	goto L1537
L1536:
	;
	goto L1537
L1537:
	;
	m.G0 = v7026 + int32(16)
	goto L1534
L1538:
	;
	v7055 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7056 = m.ExcPending
	if v7056 != 0 {
		goto L32
	} else {
		goto L1539
	}
L1539:
	;
	if v7055 == int32(0) {
		goto L1532
	} else {
		goto L1540
	}
L1540:
	;
	v7060 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v7061 = *(*int64)(unsafe.Add(mBase, uint32(v7060)+32))
	v7062 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+320)) = v7062
	v7064 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+540))
	v7066 = base.I32_div_s(v7064, int32(_a_F_StartupXLOG_220))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+324)) = v7066
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+332)) = uint32(v7061)
	v7070 = int64(base.Ui64(v7061) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+328)) = uint32(v7070)
	F_errmsg(m, int32(_a_F_StartupXLOG_221), v6731+int32(320))
	mBase = m.M
	v7076 = m.ExcPending
	if v7076 != 0 {
		goto L32
	} else {
		goto L1541
	}
L1541:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1723), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v7081 = m.ExcPending
	if v7081 != 0 {
		goto L32
	} else {
		goto L1542
	}
L1542:
	;
	goto L1532
L1543:
	;
	v7086 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7087 = *(*int32)(unsafe.Add(mBase, uint32(v7086)+80))
	if v7087 != 0 {
		goto L1544
	} else {
		goto L1545
	}
L1544:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v7090 = m.ExcPending
	if v7090 != 0 {
		goto L32
	} else {
		goto L1547
	}
L1545:
	;
	goto L1546
L1546:
	;
	v7092 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v7092 != int32(1) {
		goto L1548
	} else {
		goto L1549
	}
L1547:
	;
	goto L1546
L1548:
	;
	v7487 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[116]))
	if v7487 <= int32(0) {
		goto L1633
	} else {
		goto L1634
	}
L1549:
	;
	v7096 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v7098 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v7098 != int32(5) {
		goto L1550
	} else {
		goto L1551
	}
L1550:
	;
	if v7098 != int32(4) {
		goto L1559
	} else {
		goto L1560
	}
L1551:
	;
	v7102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v7102&int32(1) == int32(0) {
		goto L1550
	} else {
		goto L1552
	}
L1552:
	;
	v7109 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7110 = m.ExcPending
	if v7110 != 0 {
		goto L32
	} else {
		goto L1553
	}
L1553:
	;
	if v7109 != 0 {
		goto L1554
	} else {
		goto L1555
	}
L1554:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_222), int32(0))
	mBase = m.M
	v7114 = m.ExcPending
	if v7114 != 0 {
		goto L32
	} else {
		goto L1557
	}
L1555:
	;
	goto L1556
L1556:
	;
	v7121 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = v7121
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = v7121
	v7127 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v7127
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v7127)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v7127)
	goto L1490
L1557:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2578), int32(_a_F_StartupXLOG_223))
	mBase = m.M
	v7119 = m.ExcPending
	if v7119 != 0 {
		goto L32
	} else {
		goto L1558
	}
L1558:
	;
	goto L1556
L1559:
	;
	v7182 = *(*int32)(unsafe.Add(mBase, uint32(v7096)+96))
	v7183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7182)+49)))
	if v7183 != int32(1) {
		goto L1548
	} else {
		goto L1567
	}
L1560:
	;
	v7138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[117])))
	if v7138&int32(1) != 0 {
		goto L1559
	} else {
		goto L1561
	}
L1561:
	;
	v7141 = *(*int64)(unsafe.Add(mBase, uint32(v7096)+32))
	v7143 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	if base.Ui64(v7141) < base.Ui64(v7143) {
		goto L1559
	} else {
		goto L1562
	}
L1562:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = v7141
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = int64(0)
	v7151 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v7151
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v7151)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v7151)
	v7161 = F_errstart(m, int32(15), v7151)
	mBase = m.M
	v7162 = m.ExcPending
	if v7162 != 0 {
		goto L32
	} else {
		goto L1563
	}
L1563:
	;
	if v7161 == int32(0) {
		goto L1490
	} else {
		goto L1564
	}
L1564:
	;
	v7166 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+308)) = uint32(v7166)
	v7169 = int64(base.Ui64(v7166) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+304)) = uint32(v7169)
	F_errmsg(m, int32(_a_F_StartupXLOG_224), v6731+int32(304))
	mBase = m.M
	v7175 = m.ExcPending
	if v7175 != 0 {
		goto L32
	} else {
		goto L1565
	}
L1565:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2600), int32(_a_F_StartupXLOG_223))
	mBase = m.M
	v7180 = m.ExcPending
	if v7180 != 0 {
		goto L32
	} else {
		goto L1566
	}
L1566:
	;
	goto L1490
L1567:
	;
	v7186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7182)+48)))
	switch int32(base.Ui32(v7186)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		goto L1573
	default:
		goto L1548
	case 2:
		goto L1571
	case 3:
		goto L1572
	case 4:
		goto L1570
	}
L1568:
	;
	v7421 = int32(0)
	v7423 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[117])))
	v7424 = int32(1)
	v7427 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v7423&v7424|base.B2i32(v7427 != v7424) == v7421 {
		goto L1615
	} else {
		goto L1616
	}
L1569:
	;
	v7419 = v7417
	v7420 = int32(0)
	goto L1568
L1570:
	;
	v7313 = *(*int32)(unsafe.Add(mBase, uint32(v7182)+64))
	v7315 = v6731 + int32(560)
	v7316 = int32(0)
	base.MemoryFill(m, v7315, v7316, int32(264))
	v7322 = *(*int64)(unsafe.Add(mBase, uint32(v7313)))
	*(*int64)(unsafe.Add(mBase, uint32(v7315))) = v7322
	if v7316 <= base.I32_extend8_s(v7186) {
		goto L1597
	} else {
		goto L1598
	}
L1571:
	;
	v7312 = *(*int32)(unsafe.Add(mBase, uint32(v7182)+36))
	v7417 = v7312
	goto L1569
L1572:
	;
	v7193 = *(*int32)(unsafe.Add(mBase, uint32(v7182)+64))
	v7195 = v6731 + int32(560)
	v7196 = int32(0)
	base.MemoryFill(m, v7195, v7196, int32(288))
	v7202 = *(*int64)(unsafe.Add(mBase, uint32(v7193)))
	*(*int64)(unsafe.Add(mBase, uint32(v7195))) = v7202
	if v7196 <= base.I32_extend8_s(v7186) {
		goto L1575
	} else {
		goto L1576
	}
L1573:
	;
	v7191 = *(*int32)(unsafe.Add(mBase, uint32(v7182)+36))
	v7419 = v7191
	v7420 = int32(1)
	goto L1568
L1574:
	;
	v7310 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+612))
	v7419 = v7310
	v7420 = int32(1)
	goto L1568
L1575:
	;
	goto L1574
L1576:
	;
	v7207 = *(*int32)(unsafe.Add(mBase, uint32(v7193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+8)) = v7207
	if v7207&int32(1) != 0 {
		goto L1577
	} else {
		goto L1578
	}
L1577:
	;
	v7211 = *(*int32)(unsafe.Add(mBase, uint32(v7193)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+12)) = v7211
	v7213 = *(*int32)(unsafe.Add(mBase, uint32(v7193)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+16)) = v7213
	v7219 = v7193 + int32(20)
	goto L1579
L1578:
	;
	v7219 = v7193 + int32(12)
	goto L1579
L1579:
	;
	if v7207&int32(2) != 0 {
		goto L1580
	} else {
		goto L1581
	}
L1580:
	;
	v7222 = *(*int32)(unsafe.Add(mBase, uint32(v7219)))
	v7224 = v7219 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+24)) = v7224
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+20)) = v7222
	v7230 = v7224 + v7222<<(uint(int32(2))%32)
	goto L1582
L1581:
	;
	v7230 = v7219
	goto L1582
L1582:
	;
	if v7207&int32(4) != 0 {
		goto L1583
	} else {
		goto L1584
	}
L1583:
	;
	v7234 = *(*int32)(unsafe.Add(mBase, uint32(v7230)))
	v7236 = v7230 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+32)) = v7236
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+28)) = v7234
	v7239 = *(*int32)(unsafe.Add(mBase, uint32(v7230)))
	v7243 = v7236 + v7239*int32(12)
	goto L1585
L1584:
	;
	v7243 = v7230
	goto L1585
L1585:
	;
	if v7207&int32(256) != 0 {
		goto L1586
	} else {
		goto L1587
	}
L1586:
	;
	v7248 = *(*int32)(unsafe.Add(mBase, uint32(v7243)))
	v7249 = int32(4)
	v7250 = v7243 + v7249
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+40)) = v7250
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+36)) = v7248
	v7253 = *(*int32)(unsafe.Add(mBase, uint32(v7243)))
	v7257 = v7250 + v7253<<(uint(v7249)%32)
	goto L1588
L1587:
	;
	v7257 = v7243
	goto L1588
L1588:
	;
	if v7207&int32(8) != 0 {
		goto L1589
	} else {
		goto L1590
	}
L1589:
	;
	v7262 = *(*int32)(unsafe.Add(mBase, uint32(v7257)))
	v7263 = int32(4)
	v7264 = v7257 + v7263
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+48)) = v7264
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+44)) = v7262
	v7267 = *(*int32)(unsafe.Add(mBase, uint32(v7257)))
	v7271 = v7264 + v7267<<(uint(v7263)%32)
	goto L1591
L1590:
	;
	v7271 = v7257
	goto L1591
L1591:
	;
	if v7207&int32(16) == int32(0) {
		v7295 = v7207
		v7296 = v7271
		goto L1592
	} else {
		goto L1593
	}
L1592:
	;
	if v7295&int32(32) == int32(0) {
		goto L1575
	} else {
		goto L1595
	}
L1593:
	;
	v7278 = *(*int32)(unsafe.Add(mBase, uint32(v7271)))
	*(*int32)(unsafe.Add(mBase, uint32(v7195)+52)) = v7278
	v7281 = v7271 + int32(4)
	if v7207&int32(128) == int32(0) {
		v7295 = v7207
		v7296 = v7281
		goto L1592
	} else {
		goto L1594
	}
L1594:
	;
	v7289 = F_strlcpy(m, v6731+int32(616), v7281, int32(200))
	mBase = m.M
	v7290 = F_strlen(m, v7281)
	mBase = m.M
	v7294 = *(*int32)(unsafe.Add(mBase, uint32(v7195)+8))
	v7295 = v7294
	v7296 = v7290 + v7281 + int32(1)
	goto L1592
L1595:
	;
	v7301 = *(*int64)(unsafe.Add(mBase, uint32(v7296)))
	v7302 = *(*int64)(unsafe.Add(mBase, uint32(v7296)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7195)+280)) = v7302
	*(*int64)(unsafe.Add(mBase, uint32(v7195)+272)) = v7301
	goto L1575
L1596:
	;
	v7416 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+604))
	v7417 = v7416
	goto L1569
L1597:
	;
	goto L1596
L1598:
	;
	v7327 = *(*int32)(unsafe.Add(mBase, uint32(v7313)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+8)) = v7327
	if v7327&int32(1) != 0 {
		goto L1599
	} else {
		goto L1600
	}
L1599:
	;
	v7331 = *(*int32)(unsafe.Add(mBase, uint32(v7313)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+12)) = v7331
	v7333 = *(*int32)(unsafe.Add(mBase, uint32(v7313)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+16)) = v7333
	v7339 = v7313 + int32(20)
	goto L1601
L1600:
	;
	v7339 = v7313 + int32(12)
	goto L1601
L1601:
	;
	if v7327&int32(2) != 0 {
		goto L1602
	} else {
		goto L1603
	}
L1602:
	;
	v7342 = *(*int32)(unsafe.Add(mBase, uint32(v7339)))
	v7344 = v7339 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+24)) = v7344
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+20)) = v7342
	v7350 = v7344 + v7342<<(uint(int32(2))%32)
	goto L1604
L1603:
	;
	v7350 = v7339
	goto L1604
L1604:
	;
	if v7327&int32(4) != 0 {
		goto L1605
	} else {
		goto L1606
	}
L1605:
	;
	v7354 = *(*int32)(unsafe.Add(mBase, uint32(v7350)))
	v7356 = v7350 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+32)) = v7356
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+28)) = v7354
	v7359 = *(*int32)(unsafe.Add(mBase, uint32(v7350)))
	v7363 = v7356 + v7359*int32(12)
	goto L1607
L1606:
	;
	v7363 = v7350
	goto L1607
L1607:
	;
	if v7327&int32(256) != 0 {
		goto L1608
	} else {
		goto L1609
	}
L1608:
	;
	v7368 = *(*int32)(unsafe.Add(mBase, uint32(v7363)))
	v7369 = int32(4)
	v7370 = v7363 + v7369
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+40)) = v7370
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+36)) = v7368
	v7373 = *(*int32)(unsafe.Add(mBase, uint32(v7363)))
	v7377 = v7370 + v7373<<(uint(v7369)%32)
	goto L1610
L1609:
	;
	v7377 = v7363
	goto L1610
L1610:
	;
	if v7327&int32(16) == int32(0) {
		v7401 = v7327
		v7402 = v7377
		goto L1611
	} else {
		goto L1612
	}
L1611:
	;
	if v7401&int32(32) == int32(0) {
		goto L1597
	} else {
		goto L1614
	}
L1612:
	;
	v7384 = *(*int32)(unsafe.Add(mBase, uint32(v7377)))
	*(*int32)(unsafe.Add(mBase, uint32(v7315)+44)) = v7384
	v7387 = v7377 + int32(4)
	if v7327&int32(128) == int32(0) {
		v7401 = v7327
		v7402 = v7387
		goto L1611
	} else {
		goto L1613
	}
L1613:
	;
	v7395 = F_strlcpy(m, v6731+int32(608), v7387, int32(200))
	mBase = m.M
	v7396 = F_strlen(m, v7387)
	mBase = m.M
	v7400 = *(*int32)(unsafe.Add(mBase, uint32(v7315)+8))
	v7401 = v7400
	v7402 = v7396 + v7387 + int32(1)
	goto L1611
L1614:
	;
	v7407 = *(*int64)(unsafe.Add(mBase, uint32(v7402)))
	v7408 = *(*int64)(unsafe.Add(mBase, uint32(v7402)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7315)+256)) = v7408
	*(*int64)(unsafe.Add(mBase, uint32(v7315)+248)) = v7407
	goto L1597
L1615:
	;
	v7434 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	v7436 = base.B2i32(v7419 == v7434)
	goto L1617
L1616:
	;
	v7436 = v7421
	goto L1617
L1617:
	;
	v7437 = *(*int32)(unsafe.Add(mBase, uint32(v7096)+96))
	v7438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7437)+48)))
	v7439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7437)+49)))
	if base.B2i32(v7439 == int32(0))&base.B2i32(v7438&int32(240) == int32(112)) != 0 {
		goto L1618
	} else {
		goto L1619
	}
L1618:
	;
	v7462 = *(*int32)(unsafe.Add(mBase, uint32(v7437)+64))
	v7463 = *(*int64)(unsafe.Add(mBase, uint32(v7462)))
	if v7427 == int32(2) {
		goto L1626
	} else {
		goto L1627
	}
L1619:
	;
	if v7439 != int32(1) {
		goto L1620
	} else {
		goto L1621
	}
L1620:
	;
	if v7436 == int32(0) {
		goto L1548
	} else {
		goto L1624
	}
L1621:
	;
	v7449 = int32(4)
	v7452 = int32(base.Ui32(v7438)>>(uint(v7449)%32)) & int32(7)
	if base.Ui32(v7449) < base.Ui32(v7452) {
		goto L1620
	} else {
		goto L1622
	}
L1622:
	;
	if v7452 != int32(1) {
		goto L1618
	} else {
		goto L1623
	}
L1623:
	;
	goto L1620
L1624:
	;
	v9140 = int64(0)
	goto L1491
L1625:
	;
	if v7467 <= v7463 {
		v9140 = v7463
		goto L1491
	} else {
		goto L1632
	}
L1626:
	;
	v7467 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[26]))
	if v7423&int32(1) == int32(0) {
		goto L1625
	} else {
		goto L1629
	}
L1627:
	;
	goto L1628
L1628:
	;
	if v7436 == int32(0) {
		goto L1548
	} else {
		goto L1631
	}
L1629:
	;
	if v7467 < v7463 {
		v9140 = v7463
		goto L1491
	} else {
		goto L1630
	}
L1630:
	;
	goto L1548
L1631:
	;
	v9140 = v7463
	goto L1491
L1632:
	;
	goto L1548
L1633:
	;
	v7780 = int32(0)
	v7781 = int32(_a_F_StartupXLOG_225)
	v7782 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[118]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[118])) = v6731 + int32(540)
	v7788 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+548)) = v7788
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+544)) = int32(441)
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+540)) = v7782
	v7793 = *(*int32)(unsafe.Add(mBase, uint32(v6990)+4))
	F_AdvanceNextFullTransactionIdPastXid(m, v7793)
	mBase = m.M
	v7795 = m.ExcPending
	if v7795 != 0 {
		goto L32
	} else {
		goto L1675
	}
L1634:
	;
	v7491 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v7491&int32(1) == int32(0) {
		goto L1633
	} else {
		goto L1635
	}
L1635:
	;
	v7497 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v7497&int32(1) == int32(0) {
		goto L1633
	} else {
		goto L1636
	}
L1636:
	;
	v7503 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v7504 = *(*int32)(unsafe.Add(mBase, uint32(v7503)+96))
	v7505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7504)+49)))
	if v7505 != int32(1) {
		goto L1633
	} else {
		goto L1637
	}
L1637:
	;
	v7508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7504)+48)))
	v7510 = v7508 & int32(112)
	if v7510 != 0 {
		goto L1638
	} else {
		goto L1639
	}
L1638:
	;
	v7514 = base.B2i32(v7510 != int32(48))
	goto L1640
L1639:
	;
	v7514 = int32(0)
	goto L1640
L1640:
	;
	if v7514 != 0 {
		goto L1633
	} else {
		goto L1641
	}
L1641:
	;
	v7515 = int32(4)
	v7518 = int32(base.Ui32(v7508)>>(uint(v7515)%32)) & int32(7)
	if base.B2i32(base.Ui32(v7515) < base.Ui32(v7518))|base.B2i32(v7518 == int32(1)) != 0 {
		goto L1633
	} else {
		goto L1642
	}
L1642:
	;
	v7524 = *(*int32)(unsafe.Add(mBase, uint32(v7504)+64))
	v7525 = *(*int64)(unsafe.Add(mBase, uint32(v7524)))
	v7529 = m.G0
	v7530 = int32(16)
	v7531 = v7529 - v7530
	m.G0 = v7531
	F_gettimeofday(m, v7531)
	mBase = m.M
	v7534 = *(*int64)(unsafe.Add(mBase, uint32(v7531)))
	v7535 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7531)+8)))
	m.G0 = v7531 + v7530
	v7543 = v7535 + v7534*int64(1000000) - int64(946684800000000)
	goto L1643
L1643:
	;
	v7547 = v7525 + base.I64_extend_i32_u(v7487)*int64(1000)
	if v7547 <= v7543 {
		v7565 = int32(0)
		goto L1645
	} else {
		goto L1646
	}
L1644:
	;
	if v7565 <= int32(0) {
		goto L1633
	} else {
		goto L1648
	}
L1645:
	;
	goto L1644
L1646:
	;
	v7553 = v7547 - v7543
	if base.B2i32(int64(0) < v7543)^base.B2i32(v7553 < v7547)|base.B2i32(int64(2147483646000) < v7553) != 0 {
		v7565 = int32(2147483647)
		goto L1645
	} else {
		goto L1647
	}
L1647:
	;
	v7562 = base.I64_div_s(v7553+int64(999), int64(1000))
	v7565 = base.I32_wrap_i64(v7562)
	goto L1645
L1648:
	;
	v7569 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7572 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7569+int32(4)))) = v7572
	v7577 = base.AtomicRmwOr32(m, v7572, int32(_a_F_StartupXLOG_226), v7572)
	goto L1649
L1649:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7579 = m.ExcPending
	if v7579 != 0 {
		goto L32
	} else {
		goto L1650
	}
L1650:
	;
	v7580 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7581 = m.ExcPending
	if v7581 != 0 {
		goto L32
	} else {
		goto L1652
	}
L1651:
	;
	v7737 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7738 = *(*int32)(unsafe.Add(mBase, uint32(v7737)+80))
	if v7738 == int32(0) {
		goto L1633
	} else {
		goto L1673
	}
L1652:
	;
	if v7580 != 0 {
		goto L1651
	} else {
		goto L1653
	}
L1653:
	;
	goto L1654
L1654:
	;
	v7619 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[116])))
	v7623 = m.G0
	v7624 = int32(16)
	v7625 = v7623 - v7624
	m.G0 = v7625
	F_gettimeofday(m, v7625)
	mBase = m.M
	v7628 = *(*int64)(unsafe.Add(mBase, uint32(v7625)))
	v7629 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7625)+8)))
	m.G0 = v7625 + v7624
	v7637 = v7629 + v7628*int64(1000000) - int64(946684800000000)
	goto L1656
L1655:
	;
	goto L1651
L1656:
	;
	v7640 = v7619*int64(1000) + v7525
	if v7640 <= v7637 {
		v7658 = int32(0)
		goto L1658
	} else {
		goto L1659
	}
L1657:
	;
	if v7658 <= int32(0) {
		goto L1651
	} else {
		goto L1661
	}
L1658:
	;
	goto L1657
L1659:
	;
	v7646 = v7640 - v7637
	if base.B2i32(int64(0) < v7637)^base.B2i32(v7646 < v7640)|base.B2i32(int64(2147483646000) < v7646) != 0 {
		v7658 = int32(2147483647)
		goto L1658
	} else {
		goto L1660
	}
L1660:
	;
	v7655 = base.I64_div_s(v7646+int64(999), int64(1000))
	v7658 = base.I32_wrap_i64(v7655)
	goto L1658
L1661:
	;
	v7663 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v7664 = m.ExcPending
	if v7664 != 0 {
		goto L32
	} else {
		goto L1662
	}
L1662:
	;
	if v7663 != 0 {
		goto L1663
	} else {
		goto L1664
	}
L1663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+256)) = v7658
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_227), v6731+int32(256))
	mBase = m.M
	v7670 = m.ExcPending
	if v7670 != 0 {
		goto L32
	} else {
		goto L1666
	}
L1664:
	;
	goto L1665
L1665:
	;
	v7677 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7682 = F_WaitLatch(m, v7677+int32(4), int32(41), v7658, int32(150994948))
	mBase = m.M
	v7683 = m.ExcPending
	if v7683 != 0 {
		goto L32
	} else {
		goto L1668
	}
L1666:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(3042), int32(_a_F_StartupXLOG_228))
	mBase = m.M
	v7675 = m.ExcPending
	if v7675 != 0 {
		goto L32
	} else {
		goto L1667
	}
L1667:
	;
	goto L1665
L1668:
	;
	v7685 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7688 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7685+int32(4)))) = v7688
	v7693 = base.AtomicRmwOr32(m, v7688, int32(_a_F_StartupXLOG_226), v7688)
	goto L1669
L1669:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7695 = m.ExcPending
	if v7695 != 0 {
		goto L32
	} else {
		goto L1670
	}
L1670:
	;
	v7696 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7697 = m.ExcPending
	if v7697 != 0 {
		goto L32
	} else {
		goto L1671
	}
L1671:
	;
	if v7696 == int32(0) {
		goto L1654
	} else {
		goto L1672
	}
L1672:
	;
	goto L1655
L1673:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v7743 = m.ExcPending
	if v7743 != 0 {
		goto L32
	} else {
		goto L1674
	}
L1674:
	;
	goto L1633
L1675:
	;
	v7796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6990)+17)))
	if v7796 != 0 {
		v7863 = v6979
		v7866 = v7780
		goto L1683
	} else {
		goto L1684
	}
L1676:
	;
	v9117 = int32(0)
	v9119 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v9122 = F_ReadRecord(m, v9119, int32(15), v9117, v7863)
	mBase = m.M
	v9123 = m.ExcPending
	if v9123 != 0 {
		goto L32
	} else {
		goto L1990
	}
L1677:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9090 = m.ExcPending
	if v9090 != 0 {
		goto L32
	} else {
		goto L1987
	}
L1678:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9071 = m.ExcPending
	if v9071 != 0 {
		goto L32
	} else {
		goto L1983
	}
L1679:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9048 = m.ExcPending
	if v9048 != 0 {
		goto L32
	} else {
		goto L1980
	}
L1680:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v9024 = m.ExcPending
	if v9024 != 0 {
		goto L32
	} else {
		goto L1977
	}
L1681:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v9008 = m.ExcPending
	if v9008 != 0 {
		goto L32
	} else {
		goto L1974
	}
L1682:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v8991 = m.ExcPending
	if v8991 != 0 {
		goto L32
	} else {
		goto L1971
	}
L1683:
	;
	v7869 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7872 = base.AtomicRmwXchg32(m, v7869, int32(96), int32(1))
	if v7872 != 0 {
		goto L1710
	} else {
		goto L1711
	}
L1684:
	;
	v7797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6990)+16)))
	v7799 = v7797 & int32(240)
	if v7799 != 0 {
		goto L1685
	} else {
		goto L1686
	}
L1685:
	;
	v7803 = base.B2i32(v7799 != int32(144))
	goto L1687
L1686:
	;
	v7803 = int32(0)
	goto L1687
L1687:
	;
	if v7803 != 0 {
		v7863 = v6979
		v7866 = v7780
		goto L1683
	} else {
		goto L1688
	}
L1688:
	;
	v7804 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v7805 = *(*int32)(unsafe.Add(mBase, uint32(v7804)+64))
	v7806 = *(*int32)(unsafe.Add(mBase, uint32(v7805)+8))
	if v7806 == v6979 {
		v7863 = v6979
		v7866 = v7780
		goto L1683
	} else {
		goto L1689
	}
L1689:
	;
	v7808 = *(*int32)(unsafe.Add(mBase, uint32(v7805)+12))
	if v7808 != v6979 {
		goto L1682
	} else {
		goto L1690
	}
L1690:
	;
	if base.Ui32(v7806) < base.Ui32(v6979) {
		goto L1681
	} else {
		goto L1691
	}
L1691:
	;
	v7811 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	v7813 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v7814 = int32(0)
	if v7813 == v7814 {
		goto L1693
	} else {
		goto L1694
	}
L1692:
	;
	if v7853 == int32(0) {
		goto L1681
	} else {
		goto L1705
	}
L1693:
	;
	v7853 = int32(0)
	goto L1692
L1694:
	;
	goto L1695
L1695:
	;
	v7820 = *(*int32)(unsafe.Add(mBase, uint32(v7813)+4))
	if v7820 <= int32(0) {
		v7847 = v7814
		goto L1696
	} else {
		goto L1697
	}
L1696:
	;
	v7853 = v7847
	goto L1692
L1697:
	;
	v7823 = int32(0)
	if v7823 < v7820 {
		goto L1698
	} else {
		goto L1699
	}
L1698:
	;
	v7826 = v7820
	goto L1700
L1699:
	;
	v7826 = v7823
	goto L1700
L1700:
	;
	v7827 = *(*int32)(unsafe.Add(mBase, uint32(v7813)+12))
	v7830 = int32(0)
	goto L1701
L1701:
	;
	v7837 = *(*int32)(unsafe.Add(mBase, uint32(v7827+v7830<<(uint(int32(2))%32))))
	v7838 = *(*int32)(unsafe.Add(mBase, uint32(v7837)))
	v7839 = base.B2i32(v7838 == v7806)
	if v7838 == v7806 {
		v7847 = v7839
		goto L1696
	} else {
		goto L1703
	}
L1702:
	;
	v7847 = v7839
	goto L1696
L1703:
	;
	v7841 = v7830 + int32(1)
	if v7841 != v7826 {
		v7830 = v7841
		goto L1701
	} else {
		goto L1704
	}
L1704:
	;
	goto L1702
L1705:
	;
	v7858 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51]))
	if base.Ui64(v7811) < base.Ui64(v7858) {
		goto L1706
	} else {
		goto L1707
	}
L1706:
	;
	v7861 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	if base.Ui32(v7861) < base.Ui32(v7806) {
		goto L1680
	} else {
		goto L1709
	}
L1707:
	;
	goto L1708
L1708:
	;
	v7863 = v7806
	v7866 = int32(1)
	goto L1683
L1709:
	;
	goto L1708
L1710:
	;
	F_s_lock(m, v7869+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v7877 = m.ExcPending
	if v7877 != 0 {
		goto L32
	} else {
		goto L1713
	}
L1711:
	;
	goto L1712
L1712:
	;
	v7878 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	v7880 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v7880)+56)) = v7863
	*(*int64)(unsafe.Add(mBase, uint32(v7880)+48)) = v7878
	v7883 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7880)+96)), uint32(v7883))
	v7887 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[92]))
	if v7887 == v7883 {
		goto L1714
	} else {
		goto L1715
	}
L1713:
	;
	goto L1712
L1714:
	;
	v7896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6990)+17)))
	if v7896 != 0 {
		goto L1718
	} else {
		goto L1719
	}
L1715:
	;
	v7890 = *(*int32)(unsafe.Add(mBase, uint32(v6990)+4))
	if v7890 == int32(0) {
		goto L1714
	} else {
		goto L1716
	}
L1716:
	;
	F_RecordKnownAssignedTransactionIds(m, v7890)
	mBase = m.M
	v7894 = m.ExcPending
	if v7894 != 0 {
		goto L32
	} else {
		goto L1717
	}
L1717:
	;
	goto L1714
L1718:
	;
	v7987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6990)+17)))
	v7989 = v7987 << (uint(int32(5)) % 32)
	v7992 = *(*int32)(unsafe.Add(mBase, uint32(v7989)+uint32(_c_F_StartupXLOG[110])))
	if v7992 == int32(0) {
		goto L1744
	} else {
		goto L1745
	}
L1719:
	;
	v7897 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v7898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7897)+48)))
	v7900 = v7898 & int32(240)
	if v7900 != int32(80) {
		goto L1720
	} else {
		goto L1721
	}
L1720:
	;
	if v7900 != int32(208) {
		goto L1718
	} else {
		goto L1723
	}
L1721:
	;
	goto L1722
L1722:
	;
	v7939 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	v7941 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54]))
	v7942 = *(*int32)(unsafe.Add(mBase, uint32(v7897)+64))
	v7943 = *(*int64)(unsafe.Add(mBase, uint32(v7942)))
	v7946 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v7947 = m.ExcPending
	if v7947 != 0 {
		goto L32
	} else {
		goto L1732
	}
L1723:
	;
	v7905 = *(*int32)(unsafe.Add(mBase, uint32(v7897)+64))
	v7906 = *(*int64)(unsafe.Add(mBase, uint32(v7905)))
	v7907 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+64))
	if v7906 != v7907 {
		goto L1679
	} else {
		goto L1724
	}
L1724:
	;
	v7909 = *(*int64)(unsafe.Add(mBase, uint32(v7905)+8))
	v7911 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v7911
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v7911
	v7918 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7919 = m.ExcPending
	if v7919 != 0 {
		goto L32
	} else {
		goto L1725
	}
L1725:
	;
	if v7918 != 0 {
		goto L1726
	} else {
		goto L1727
	}
L1726:
	;
	v7920 = F_timestamptz_to_str(m, v7909)
	mBase = m.M
	v7921 = m.ExcPending
	if v7921 != 0 {
		goto L32
	} else {
		goto L1729
	}
L1727:
	;
	goto L1728
L1728:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7788)+64)) = int64(0)
	goto L1718
L1729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+168)) = v7920
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+164)) = uint32(v7906)
	v7925 = int64(base.Ui64(v7906) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+160)) = uint32(v7925)
	F_errmsg(m, int32(_a_F_StartupXLOG_229), v6731+int32(160))
	mBase = m.M
	v7931 = m.ExcPending
	if v7931 != 0 {
		goto L32
	} else {
		goto L1730
	}
L1730:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2081), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7936 = m.ExcPending
	if v7936 != 0 {
		goto L32
	} else {
		goto L1731
	}
L1731:
	;
	goto L1728
L1732:
	;
	if v7941 == v7943 {
		goto L1733
	} else {
		goto L1734
	}
L1733:
	;
	if v7946 != 0 {
		goto L1736
	} else {
		goto L1737
	}
L1734:
	;
	goto L1735
L1735:
	;
	if v7946 == int32(0) {
		goto L1718
	} else {
		goto L1741
	}
L1736:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_231), int32(0))
	mBase = m.M
	v7952 = m.ExcPending
	if v7952 != 0 {
		goto L32
	} else {
		goto L1739
	}
L1737:
	;
	goto L1738
L1738:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[55])) = v7939
	goto L1718
L1739:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2102), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7957 = m.ExcPending
	if v7957 != 0 {
		goto L32
	} else {
		goto L1740
	}
L1740:
	;
	goto L1738
L1741:
	;
	v7963 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+204)) = uint32(v7963)
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+196)) = uint32(v7943)
	v7966 = int64(32)
	v7967 = int64(base.Ui64(v7943) >> (uint(v7966) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+192)) = uint32(v7967)
	v7970 = int64(base.Ui64(v7963) >> (uint(v7966) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+200)) = uint32(v7970)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_232), v6731+int32(192))
	mBase = m.M
	v7976 = m.ExcPending
	if v7976 != 0 {
		goto L32
	} else {
		goto L1742
	}
L1742:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2108), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7981 = m.ExcPending
	if v7981 != 0 {
		goto L32
	} else {
		goto L1743
	}
L1743:
	;
	goto L1718
L1744:
	;
	F_RmgrNotFound(m, v7987)
	mBase = m.M
	v7996 = m.ExcPending
	if v7996 != 0 {
		goto L32
	} else {
		goto L1747
	}
L1745:
	;
	goto L1746
L1746:
	;
	v7997 = *(*int32)(unsafe.Add(mBase, uint32(v7989)+uint32(_c_F_StartupXLOG[119])))
	m.T0[v7997].(func(*base.Module, int32))(m, v7788)
	mBase = m.M
	v7999 = m.ExcPending
	if v7999 != 0 {
		goto L32
	} else {
		goto L1748
	}
L1747:
	;
	goto L1746
L1748:
	;
	v8000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6990)+16)))
	if v8000&int32(2) == int32(0) {
		goto L1749
	} else {
		goto L1750
	}
L1749:
	;
	v8274 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+540))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[118])) = v8274
	v8277 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8280 = base.AtomicRmwXchg32(m, v8277, int32(96), int32(1))
	if v8280 != 0 {
		goto L1811
	} else {
		goto L1812
	}
L1750:
	;
	v8005 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v8006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8005)+49)))
	v8008 = v8006 << (uint(int32(5)) % 32)
	v8009 = *(*int32)(unsafe.Add(mBase, uint32(v8008)+uint32(_c_F_StartupXLOG[110])))
	if v8009 != 0 {
		goto L1751
	} else {
		goto L1752
	}
L1751:
	;
	v8013 = v8005
	goto L1753
L1752:
	;
	F_RmgrNotFound(m, v8006)
	mBase = m.M
	v8011 = m.ExcPending
	if v8011 != 0 {
		goto L32
	} else {
		goto L1754
	}
L1753:
	;
	v8014 = *(*int32)(unsafe.Add(mBase, uint32(v8013)+72))
	if v8014 < int32(0) {
		goto L1749
	} else {
		goto L1755
	}
L1754:
	;
	v8012 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v8013 = v8012
	goto L1753
L1755:
	;
	v8019 = *(*int32)(unsafe.Add(mBase, uint32(v8008)+uint32(_c_F_StartupXLOG[120])))
	v8032 = int32(0)
	goto L1756
L1756:
	;
	v8058 = v8032 & int32(255)
	v8060 = v6731 + int32(560)
	v8062 = v6731 + int32(556)
	v8064 = v6731 + int32(552)
	v8065 = int32(0)
	v8067 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v8068 = *(*int32)(unsafe.Add(mBase, uint32(v8067)+72))
	if v8068 < v8058 {
		v8092 = v8065
		goto L1760
	} else {
		goto L1761
	}
L1757:
	;
	goto L1749
L1758:
	;
	v8233 = v8032 + int32(1)
	v8234 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v8235 = *(*int32)(unsafe.Add(mBase, uint32(v8234)+72))
	if v8233 <= v8235 {
		v8032 = v8233
		goto L1756
	} else {
		goto L1810
	}
L1759:
	;
	if v8092 == int32(0) {
		goto L1758
	} else {
		goto L1773
	}
L1760:
	;
	goto L1759
L1761:
	;
	v8072 = v8067 + v8058*int32(52)
	v8073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8072)+76)))
	if v8073 != int32(1) {
		v8092 = v8065
		goto L1760
	} else {
		goto L1762
	}
L1762:
	;
	v8077 = v8072 + int32(76)
	if v8060 != 0 {
		goto L1763
	} else {
		goto L1764
	}
L1763:
	;
	v8078 = *(*int32)(unsafe.Add(mBase, uint32(v8077)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+8)) = v8078
	v8080 = *(*int64)(unsafe.Add(mBase, uint32(v8077)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v8060))) = v8080
	goto L1765
L1764:
	;
	goto L1765
L1765:
	;
	if v8062 != 0 {
		goto L1766
	} else {
		goto L1767
	}
L1766:
	;
	v8082 = *(*int32)(unsafe.Add(mBase, uint32(v8077)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8062))) = v8082
	goto L1768
L1767:
	;
	goto L1768
L1768:
	;
	if v8064 != 0 {
		goto L1769
	} else {
		goto L1770
	}
L1769:
	;
	v8084 = *(*int32)(unsafe.Add(mBase, uint32(v8077)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8064))) = v8084
	goto L1771
L1770:
	;
	goto L1771
L1771:
	;
	v8092 = int32(1)
	goto L1760
L1773:
	;
	v8095 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+96))
	v8099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8095+v8032*int32(52))+106)))
	if v8099 != 0 {
		goto L1758
	} else {
		goto L1774
	}
L1774:
	;
	v8100 = *(*int64)(unsafe.Add(mBase, uint32(v6731)+560))
	*(*int64)(unsafe.Add(mBase, uint32(v6731)+144)) = v8100
	v8102 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+568))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+152)) = v8102
	v8106 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+556))
	v8107 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+552))
	v8110 = F_XLogReadBufferExtended(m, v6731+int32(144), v8106, v8107, int32(4), int32(0))
	mBase = m.M
	v8111 = m.ExcPending
	if v8111 != 0 {
		goto L32
	} else {
		goto L1775
	}
L1775:
	;
	if v8110 == int32(0) {
		goto L1758
	} else {
		goto L1776
	}
L1776:
	;
	F_LockBufferInternal(m, v8110, int32(3))
	mBase = m.M
	v8116 = m.ExcPending
	if v8116 != 0 {
		goto L32
	} else {
		goto L1777
	}
L1777:
	;
	v8118 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	if v8110 < int32(0) {
		goto L1779
	} else {
		goto L1780
	}
L1778:
	;
	base.MemoryCopy(m, v8118, v8136, int32(_a_F_StartupXLOG_57))
	F_UnlockReleaseBuffer(m, v8110)
	mBase = m.M
	v8140 = m.ExcPending
	if v8140 != 0 {
		goto L32
	} else {
		goto L1782
	}
L1779:
	;
	v8122 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[121]))
	v8128 = *(*int32)(unsafe.Add(mBase, uint32(v8122+(v8110^int32(-1))<<(uint(int32(2))%32))))
	v8136 = v8128
	goto L1778
L1780:
	;
	goto L1781
L1781:
	;
	v8130 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[122]))
	v8136 = v8130 + v8110<<(uint(int32(13))%32) + int32(-8192)
	goto L1778
L1782:
	;
	v8141 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	v8143 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v8144 = *(*int64)(unsafe.Add(mBase, uint32(v8143)))
	if base.Ui64(v8141) < base.Ui64(base.I64_rotl(v8144, int64(32))) {
		goto L1758
	} else {
		goto L1783
	}
L1783:
	;
	v8149 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v8150 = F_RestoreBlockImage(m, v7788, v8058, v8149)
	mBase = m.M
	v8151 = m.ExcPending
	if v8151 != 0 {
		goto L32
	} else {
		goto L1784
	}
L1784:
	;
	if v8150 == int32(0) {
		goto L1678
	} else {
		goto L1785
	}
L1785:
	;
	if v8019 != 0 {
		goto L1786
	} else {
		goto L1787
	}
L1786:
	;
	v8155 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v8156 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+552))
	m.T0[v8019].(func(*base.Module, int32, int32))(m, v8155, v8156)
	mBase = m.M
	v8158 = m.ExcPending
	if v8158 != 0 {
		goto L32
	} else {
		goto L1789
	}
L1787:
	;
	goto L1788
L1788:
	;
	v8165 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v8167 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v8168 = int32(_a_F_StartupXLOG_57)
	goto L1794
L1789:
	;
	v8160 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v8161 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+552))
	m.T0[v8019].(func(*base.Module, int32, int32))(m, v8160, v8161)
	mBase = m.M
	v8163 = m.ExcPending
	if v8163 != 0 {
		goto L32
	} else {
		goto L1790
	}
L1790:
	;
	goto L1788
L1791:
	;
	if v8230 != 0 {
		goto L1677
	} else {
		goto L1809
	}
L1792:
	;
	v8230 = int32(0)
	goto L1791
L1793:
	;
	v8204 = v8199
	v8205 = v8200
	v8206 = v8201
	goto L1803
L1794:
	;
	if (v8165|v8167)&int32(3) != 0 {
		v8199 = v8165
		v8200 = v8167
		v8201 = v8168
		goto L1793
	} else {
		goto L1797
	}
L1796:
	;
	if v8189 == int32(0) {
		goto L1792
	} else {
		goto L1802
	}
L1797:
	;
	v8176 = v8165
	v8177 = v8167
	v8178 = v8168
	goto L1798
L1798:
	;
	v8181 = *(*int32)(unsafe.Add(mBase, uint32(v8176)))
	v8182 = *(*int32)(unsafe.Add(mBase, uint32(v8177)))
	if v8181 != v8182 {
		v8199 = v8176
		v8200 = v8177
		v8201 = v8178
		goto L1793
	} else {
		goto L1800
	}
L1799:
	;
	goto L1796
L1800:
	;
	v8184 = int32(4)
	v8185 = v8177 + v8184
	v8187 = v8176 + v8184
	v8189 = v8178 - v8184
	if base.Ui32(int32(3)) < base.Ui32(v8189) {
		v8176 = v8187
		v8177 = v8185
		v8178 = v8189
		goto L1798
	} else {
		goto L1801
	}
L1801:
	;
	goto L1799
L1802:
	;
	v8199 = v8187
	v8200 = v8185
	v8201 = v8189
	goto L1793
L1803:
	;
	v8209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8204))))
	v8210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8205))))
	if v8209 == v8210 {
		goto L1805
	} else {
		goto L1806
	}
L1804:
	;
	v8230 = v8209 - v8210
	goto L1791
L1805:
	;
	v8212 = int32(1)
	v8217 = v8206 - v8212
	if v8217 != 0 {
		v8204 = v8204 + v8212
		v8205 = v8205 + v8212
		v8206 = v8217
		goto L1803
	} else {
		goto L1808
	}
L1806:
	;
	goto L1807
L1807:
	;
	goto L1804
L1808:
	;
	goto L1792
L1809:
	;
	goto L1758
L1810:
	;
	goto L1757
L1811:
	;
	F_s_lock(m, v8277+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v8285 = m.ExcPending
	if v8285 != 0 {
		goto L32
	} else {
		goto L1814
	}
L1812:
	;
	goto L1813
L1813:
	;
	v8287 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8288 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v8287)+24)) = v8288
	v8290 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8287)+40)) = v7863
	*(*int64)(unsafe.Add(mBase, uint32(v8287)+32)) = v8290
	v8293 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8287)+96)), uint32(v8293))
	v8297 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[23])))
	if v8297 != int32(1) {
		goto L1815
	} else {
		goto L1816
	}
L1814:
	;
	goto L1813
L1815:
	;
	v8308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[123])))
	if v8308 != 0 {
		goto L1819
	} else {
		goto L1820
	}
L1816:
	;
	v8301 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	if v8301 <= int32(0) {
		goto L1815
	} else {
		goto L1817
	}
L1817:
	;
	F_WalSndWakeup(m, v7866, int32(1))
	mBase = m.M
	v8306 = m.ExcPending
	if v8306 != 0 {
		goto L32
	} else {
		goto L1818
	}
L1818:
	;
	goto L1815
L1819:
	;
	v8310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[123])) = uint8(v8310)
	F_WalRcvRequestApplyReply(m)
	mBase = m.M
	v8313 = m.ExcPending
	if v8313 != 0 {
		goto L32
	} else {
		goto L1822
	}
L1820:
	;
	goto L1821
L1821:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v8315 = m.ExcPending
	if v8315 != 0 {
		goto L32
	} else {
		goto L1823
	}
L1822:
	;
	goto L1821
L1823:
	;
	if v7866 != 0 {
		goto L1824
	} else {
		goto L1825
	}
L1824:
	;
	v8316 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+40))
	F_RemoveNonParentXlogFiles(m, v8316, v7863)
	mBase = m.M
	v8318 = m.ExcPending
	if v8318 != 0 {
		goto L32
	} else {
		goto L1827
	}
L1825:
	;
	goto L1826
L1826:
	;
	v8327 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8328 = *(*int64)(unsafe.Add(mBase, uint32(v8327)+32))
	F_WaitLSNWakeup(m, int32(0), v8328)
	mBase = m.M
	v8330 = m.ExcPending
	if v8330 != 0 {
		goto L32
	} else {
		goto L1829
	}
L1827:
	;
	v8319 = int32(_a_F_StartupXLOG_233)
	v8321 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32])) = v8321 + int32(1)
	goto L1828
L1828:
	;
	goto L1826
L1829:
	;
	v8333 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8334 = *(*int64)(unsafe.Add(mBase, uint32(v8333)+32))
	F_WaitLSNWakeup(m, int32(1), v8334)
	mBase = m.M
	v8336 = m.ExcPending
	if v8336 != 0 {
		goto L32
	} else {
		goto L1830
	}
L1830:
	;
	v8339 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8340 = *(*int64)(unsafe.Add(mBase, uint32(v8339)+32))
	F_WaitLSNWakeup(m, int32(2), v8340)
	mBase = m.M
	v8342 = m.ExcPending
	if v8342 != 0 {
		goto L32
	} else {
		goto L1831
	}
L1831:
	;
	v8344 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v8344 != int32(1) {
		goto L1676
	} else {
		goto L1832
	}
L1832:
	;
	v8348 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v8349 = *(*int32)(unsafe.Add(mBase, uint32(v8348)+96))
	v8350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8349)+48)))
	v8351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8349)+49)))
	v8353 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v8351|base.B2i32(v8353 != int32(3))|base.B2i32(v8350&int32(240) != int32(112)) == int32(0) {
		goto L1833
	} else {
		goto L1834
	}
L1833:
	;
	v8364 = *(*int32)(unsafe.Add(mBase, uint32(v8349)+64))
	v8366 = v8364 + int32(8)
	v8368 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	v8371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8366))))
	v8374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8368))))
	if base.B2i32(v8371 == int32(0))|base.B2i32(v8371 != v8374) != 0 {
		v8392 = v8371
		v8393 = v8374
		goto L1837
	} else {
		goto L1838
	}
L1834:
	;
	goto L1835
L1835:
	;
	if v8353 != int32(4) {
		goto L1880
	} else {
		goto L1881
	}
L1836:
	;
	if v8392-v8393 != 0 {
		goto L1676
	} else {
		goto L1843
	}
L1837:
	;
	goto L1836
L1838:
	;
	v8377 = v8366
	v8378 = v8368
	goto L1839
L1839:
	;
	v8381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8378)+1)))
	v8382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+1)))
	if v8382 == int32(0) {
		v8392 = v8382
		v8393 = v8381
		goto L1837
	} else {
		goto L1841
	}
L1840:
	;
	v8392 = v8382
	v8393 = v8381
	goto L1837
L1841:
	;
	v8385 = int32(1)
	if v8382 == v8381 {
		v8377 = v8377 + v8385
		v8378 = v8378 + v8385
		goto L1839
	} else {
		goto L1842
	}
L1842:
	;
	goto L1840
L1843:
	;
	v8396 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v8396)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = int32(0)
	v8405 = *(*int64)(unsafe.Add(mBase, uint32(v8364)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = v8405
	v8407 = int32(_a_F_StartupXLOG_234)
	goto L1847
L1844:
	;
	v8529 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8530 = m.ExcPending
	if v8530 != 0 {
		goto L32
	} else {
		goto L1875
	}
L1845:
	;
	v8524 = F_strlen(m, v8513)
	mBase = m.M
	goto L1844
L1847:
	;
	goto L1848
L1848:
	;
	v8414 = int32(63)
	if (v8407^v8366)&int32(3) != 0 {
		goto L1852
	} else {
		goto L1853
	}
L1849:
	;
	v8517 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8514))) = uint8(v8517)
	goto L1845
L1850:
	;
	v8498 = v8493
	v8499 = v8494
	v8500 = v8495
	goto L1871
L1851:
	;
	if v8488 == int32(0) {
		v8513 = v8486
		v8514 = v8487
		goto L1849
	} else {
		goto L1870
	}
L1852:
	;
	v8486 = v8366
	v8487 = v8407
	v8488 = v8414
	goto L1851
L1853:
	;
	goto L1854
L1854:
	;
	v8418 = int32(0)
	if base.B2i32(v8366&int32(3) == v8418)|int32(0) == v8418 {
		goto L1856
	} else {
		goto L1857
	}
L1855:
	;
	if v8454 == int32(0) {
		v8513 = v8451
		v8514 = v8452
		goto L1849
	} else {
		goto L1864
	}
L1856:
	;
	v8430 = v8366
	v8431 = v8407
	v8432 = v8414
	goto L1859
L1857:
	;
	goto L1858
L1858:
	;
	v8451 = v8366
	v8452 = v8407
	v8453 = v8414
	v8454 = int32(1)
	goto L1855
L1859:
	;
	v8434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8430))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8431))) = uint8(v8434)
	if v8434 == int32(0) {
		v8493 = v8430
		v8494 = v8431
		v8495 = v8432
		goto L1850
	} else {
		goto L1861
	}
L1860:
	;
	v8451 = v8445
	v8452 = v8439
	v8453 = v8441
	v8454 = v8443
	goto L1855
L1861:
	;
	v8438 = int32(1)
	v8439 = v8431 + v8438
	v8441 = v8432 - v8438
	v8442 = int32(0)
	v8443 = base.B2i32(v8441 != v8442)
	v8445 = v8430 + v8438
	if v8445&int32(3) == v8442 {
		v8451 = v8445
		v8452 = v8439
		v8453 = v8441
		v8454 = v8443
		goto L1855
	} else {
		goto L1862
	}
L1862:
	;
	if v8441 != 0 {
		v8430 = v8445
		v8431 = v8439
		v8432 = v8441
		goto L1859
	} else {
		goto L1863
	}
L1863:
	;
	goto L1860
L1864:
	;
	v8457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8451))))
	if base.B2i32(v8457 == int32(0))|base.B2i32(base.Ui32(v8453) < base.Ui32(int32(4))) != 0 {
		v8486 = v8451
		v8487 = v8452
		v8488 = v8453
		goto L1851
	} else {
		goto L1865
	}
L1865:
	;
	v8464 = v8451
	v8465 = v8452
	v8466 = v8453
	goto L1866
L1866:
	;
	v8469 = *(*int32)(unsafe.Add(mBase, uint32(v8464)))
	v8472 = int32(-2139062144)
	if (int32(16843008)-v8469|v8469)&v8472 != v8472 {
		v8493 = v8464
		v8494 = v8465
		v8495 = v8466
		goto L1850
	} else {
		goto L1868
	}
L1867:
	;
	v8486 = v8480
	v8487 = v8478
	v8488 = v8482
	goto L1851
L1868:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8465))) = v8469
	v8477 = int32(4)
	v8478 = v8465 + v8477
	v8480 = v8464 + v8477
	v8482 = v8466 - v8477
	if base.Ui32(int32(3)) < base.Ui32(v8482) {
		v8464 = v8480
		v8465 = v8478
		v8466 = v8482
		goto L1866
	} else {
		goto L1869
	}
L1869:
	;
	goto L1867
L1870:
	;
	v8493 = v8486
	v8494 = v8487
	v8495 = v8488
	goto L1850
L1871:
	;
	v8502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8498))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8499))) = uint8(v8502)
	if v8502 == int32(0) {
		v8513 = v8498
		v8514 = v8499
		goto L1849
	} else {
		goto L1873
	}
L1872:
	;
	v8513 = v8509
	v8514 = v8507
	goto L1849
L1873:
	;
	v8506 = int32(1)
	v8507 = v8499 + v8506
	v8509 = v8498 + v8506
	v8511 = v8500 - v8506
	if v8511 != 0 {
		v8498 = v8509
		v8499 = v8507
		v8500 = v8511
		goto L1871
	} else {
		goto L1874
	}
L1874:
	;
	goto L1872
L1875:
	;
	if v8529 == int32(0) {
		goto L1490
	} else {
		goto L1876
	}
L1876:
	;
	v8534 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v8535 = F_timestamptz_to_str(m, v8534)
	mBase = m.M
	v8536 = m.ExcPending
	if v8536 != 0 {
		goto L32
	} else {
		goto L1877
	}
L1877:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+36)) = v8535
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+32)) = int32(_a_F_StartupXLOG_234)
	F_errmsg(m, int32(_a_F_StartupXLOG_235), v6731+int32(32))
	mBase = m.M
	v8544 = m.ExcPending
	if v8544 != 0 {
		goto L32
	} else {
		goto L1878
	}
L1878:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2751), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8549 = m.ExcPending
	if v8549 != 0 {
		goto L32
	} else {
		goto L1879
	}
L1879:
	;
	goto L1490
L1880:
	;
	if v8351 != int32(1) {
		goto L1676
	} else {
		goto L1888
	}
L1881:
	;
	v8553 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[117])))
	if v8553&int32(1) == int32(0) {
		goto L1880
	} else {
		goto L1882
	}
L1882:
	;
	v8558 = *(*int64)(unsafe.Add(mBase, uint32(v8348)+32))
	v8560 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	if base.Ui64(v8558) < base.Ui64(v8560) {
		goto L1880
	} else {
		goto L1883
	}
L1883:
	;
	v8563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v8563)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = v8558
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = int64(0)
	v8571 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v8571
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v8571)
	v8578 = F_errstart(m, int32(15), v8571)
	mBase = m.M
	v8579 = m.ExcPending
	if v8579 != 0 {
		goto L32
	} else {
		goto L1884
	}
L1884:
	;
	if v8578 == int32(0) {
		goto L1490
	} else {
		goto L1885
	}
L1885:
	;
	v8583 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+84)) = uint32(v8583)
	v8586 = int64(base.Ui64(v8583) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+80)) = uint32(v8586)
	F_errmsg(m, int32(_a_F_StartupXLOG_237), v6731+int32(80))
	mBase = m.M
	v8592 = m.ExcPending
	if v8592 != 0 {
		goto L32
	} else {
		goto L1886
	}
L1886:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2768), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8597 = m.ExcPending
	if v8597 != 0 {
		goto L32
	} else {
		goto L1887
	}
L1887:
	;
	goto L1490
L1888:
	;
	v8602 = v8350 & int32(112)
	v8603 = int32(4)
	v8604 = int32(base.Ui32(v8602) >> (uint(v8603) % 32))
	if base.B2i32(base.Ui32(v8603) < base.Ui32(v8604))|base.B2i32(v8604 == int32(1)) != 0 {
		v8949 = v8353
		goto L1889
	} else {
		goto L1890
	}
L1889:
	;
	if v8949 != int32(5) {
		goto L1676
	} else {
		goto L1963
	}
L1890:
	;
	v8611 = int32(4)
	v8614 = int32(base.Ui32(v8350)>>(uint(v8611)%32)) & int32(7)
	if base.B2i32(base.Ui32(v8611) < base.Ui32(v8614))|base.B2i32(v8614 == int32(1)) == int32(0) {
		goto L1891
	} else {
		goto L1892
	}
L1891:
	;
	v8622 = *(*int32)(unsafe.Add(mBase, uint32(v8349)+64))
	v8623 = *(*int64)(unsafe.Add(mBase, uint32(v8622)))
	v8625 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8628 = base.AtomicRmwXchg32(m, v8625, int32(96), int32(1))
	if v8628 != 0 {
		goto L1894
	} else {
		goto L1895
	}
L1892:
	;
	v8642 = v8349
	v8643 = int64(0)
	goto L1893
L1893:
	;
	v8645 = v8602 - int32(48)
	if v8645 != 0 {
		goto L1901
	} else {
		goto L1902
	}
L1894:
	;
	F_s_lock(m, v8625+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v8633 = m.ExcPending
	if v8633 != 0 {
		goto L32
	} else {
		goto L1897
	}
L1895:
	;
	goto L1896
L1896:
	;
	v8635 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int64)(unsafe.Add(mBase, uint32(v8635)+64)) = v8623
	v8637 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8635)+96)), uint32(v8637))
	v8640 = *(*int32)(unsafe.Add(mBase, uint32(v8348)+96))
	v8642 = v8640
	v8643 = v8623
	goto L1893
L1897:
	;
	goto L1896
L1898:
	;
	v8875 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v8875 != int32(1) {
		v8949 = v8875
		goto L1889
	} else {
		goto L1948
	}
L1899:
	;
	v8872 = *(*int32)(unsafe.Add(mBase, uint32(v8642)+36))
	v8873 = v8872
	goto L1898
L1900:
	;
	v8767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8642)+48)))
	v8768 = *(*int32)(unsafe.Add(mBase, uint32(v8642)+64))
	v8770 = v6731 + int32(560)
	v8771 = int32(0)
	base.MemoryFill(m, v8770, v8771, int32(264))
	v8777 = *(*int64)(unsafe.Add(mBase, uint32(v8768)))
	*(*int64)(unsafe.Add(mBase, uint32(v8770))) = v8777
	if v8771 <= base.I32_extend8_s(v8767) {
		goto L1930
	} else {
		goto L1931
	}
L1901:
	;
	if v8645 == int32(16) {
		goto L1904
	} else {
		goto L1905
	}
L1902:
	;
	goto L1903
L1903:
	;
	v8648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8642)+48)))
	v8649 = *(*int32)(unsafe.Add(mBase, uint32(v8642)+64))
	v8651 = v6731 + int32(560)
	v8652 = int32(0)
	base.MemoryFill(m, v8651, v8652, int32(288))
	v8658 = *(*int64)(unsafe.Add(mBase, uint32(v8649)))
	*(*int64)(unsafe.Add(mBase, uint32(v8651))) = v8658
	if v8652 <= base.I32_extend8_s(v8648) {
		goto L1908
	} else {
		goto L1909
	}
L1904:
	;
	goto L1900
L1905:
	;
	goto L1899
L1907:
	;
	v8766 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+612))
	v8873 = v8766
	goto L1898
L1908:
	;
	goto L1907
L1909:
	;
	v8663 = *(*int32)(unsafe.Add(mBase, uint32(v8649)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+8)) = v8663
	if v8663&int32(1) != 0 {
		goto L1910
	} else {
		goto L1911
	}
L1910:
	;
	v8667 = *(*int32)(unsafe.Add(mBase, uint32(v8649)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+12)) = v8667
	v8669 = *(*int32)(unsafe.Add(mBase, uint32(v8649)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+16)) = v8669
	v8675 = v8649 + int32(20)
	goto L1912
L1911:
	;
	v8675 = v8649 + int32(12)
	goto L1912
L1912:
	;
	if v8663&int32(2) != 0 {
		goto L1913
	} else {
		goto L1914
	}
L1913:
	;
	v8678 = *(*int32)(unsafe.Add(mBase, uint32(v8675)))
	v8680 = v8675 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+24)) = v8680
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+20)) = v8678
	v8686 = v8680 + v8678<<(uint(int32(2))%32)
	goto L1915
L1914:
	;
	v8686 = v8675
	goto L1915
L1915:
	;
	if v8663&int32(4) != 0 {
		goto L1916
	} else {
		goto L1917
	}
L1916:
	;
	v8690 = *(*int32)(unsafe.Add(mBase, uint32(v8686)))
	v8692 = v8686 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+32)) = v8692
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+28)) = v8690
	v8695 = *(*int32)(unsafe.Add(mBase, uint32(v8686)))
	v8699 = v8692 + v8695*int32(12)
	goto L1918
L1917:
	;
	v8699 = v8686
	goto L1918
L1918:
	;
	if v8663&int32(256) != 0 {
		goto L1919
	} else {
		goto L1920
	}
L1919:
	;
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v8699)))
	v8705 = int32(4)
	v8706 = v8699 + v8705
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+40)) = v8706
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+36)) = v8704
	v8709 = *(*int32)(unsafe.Add(mBase, uint32(v8699)))
	v8713 = v8706 + v8709<<(uint(v8705)%32)
	goto L1921
L1920:
	;
	v8713 = v8699
	goto L1921
L1921:
	;
	if v8663&int32(8) != 0 {
		goto L1922
	} else {
		goto L1923
	}
L1922:
	;
	v8718 = *(*int32)(unsafe.Add(mBase, uint32(v8713)))
	v8719 = int32(4)
	v8720 = v8713 + v8719
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+48)) = v8720
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+44)) = v8718
	v8723 = *(*int32)(unsafe.Add(mBase, uint32(v8713)))
	v8727 = v8720 + v8723<<(uint(v8719)%32)
	goto L1924
L1923:
	;
	v8727 = v8713
	goto L1924
L1924:
	;
	if v8663&int32(16) == int32(0) {
		v8751 = v8663
		v8752 = v8727
		goto L1925
	} else {
		goto L1926
	}
L1925:
	;
	if v8751&int32(32) == int32(0) {
		goto L1908
	} else {
		goto L1928
	}
L1926:
	;
	v8734 = *(*int32)(unsafe.Add(mBase, uint32(v8727)))
	*(*int32)(unsafe.Add(mBase, uint32(v8651)+52)) = v8734
	v8737 = v8727 + int32(4)
	if v8663&int32(128) == int32(0) {
		v8751 = v8663
		v8752 = v8737
		goto L1925
	} else {
		goto L1927
	}
L1927:
	;
	v8745 = F_strlcpy(m, v6731+int32(616), v8737, int32(200))
	mBase = m.M
	v8746 = F_strlen(m, v8737)
	mBase = m.M
	v8750 = *(*int32)(unsafe.Add(mBase, uint32(v8651)+8))
	v8751 = v8750
	v8752 = v8746 + v8737 + int32(1)
	goto L1925
L1928:
	;
	v8757 = *(*int64)(unsafe.Add(mBase, uint32(v8752)))
	v8758 = *(*int64)(unsafe.Add(mBase, uint32(v8752)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8651)+280)) = v8758
	*(*int64)(unsafe.Add(mBase, uint32(v8651)+272)) = v8757
	goto L1908
L1929:
	;
	v8871 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+604))
	v8873 = v8871
	goto L1898
L1930:
	;
	goto L1929
L1931:
	;
	v8782 = *(*int32)(unsafe.Add(mBase, uint32(v8768)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+8)) = v8782
	if v8782&int32(1) != 0 {
		goto L1932
	} else {
		goto L1933
	}
L1932:
	;
	v8786 = *(*int32)(unsafe.Add(mBase, uint32(v8768)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+12)) = v8786
	v8788 = *(*int32)(unsafe.Add(mBase, uint32(v8768)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+16)) = v8788
	v8794 = v8768 + int32(20)
	goto L1934
L1933:
	;
	v8794 = v8768 + int32(12)
	goto L1934
L1934:
	;
	if v8782&int32(2) != 0 {
		goto L1935
	} else {
		goto L1936
	}
L1935:
	;
	v8797 = *(*int32)(unsafe.Add(mBase, uint32(v8794)))
	v8799 = v8794 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+24)) = v8799
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+20)) = v8797
	v8805 = v8799 + v8797<<(uint(int32(2))%32)
	goto L1937
L1936:
	;
	v8805 = v8794
	goto L1937
L1937:
	;
	if v8782&int32(4) != 0 {
		goto L1938
	} else {
		goto L1939
	}
L1938:
	;
	v8809 = *(*int32)(unsafe.Add(mBase, uint32(v8805)))
	v8811 = v8805 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+32)) = v8811
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+28)) = v8809
	v8814 = *(*int32)(unsafe.Add(mBase, uint32(v8805)))
	v8818 = v8811 + v8814*int32(12)
	goto L1940
L1939:
	;
	v8818 = v8805
	goto L1940
L1940:
	;
	if v8782&int32(256) != 0 {
		goto L1941
	} else {
		goto L1942
	}
L1941:
	;
	v8823 = *(*int32)(unsafe.Add(mBase, uint32(v8818)))
	v8824 = int32(4)
	v8825 = v8818 + v8824
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+40)) = v8825
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+36)) = v8823
	v8828 = *(*int32)(unsafe.Add(mBase, uint32(v8818)))
	v8832 = v8825 + v8828<<(uint(v8824)%32)
	goto L1943
L1942:
	;
	v8832 = v8818
	goto L1943
L1943:
	;
	if v8782&int32(16) == int32(0) {
		v8856 = v8782
		v8857 = v8832
		goto L1944
	} else {
		goto L1945
	}
L1944:
	;
	if v8856&int32(32) == int32(0) {
		goto L1930
	} else {
		goto L1947
	}
L1945:
	;
	v8839 = *(*int32)(unsafe.Add(mBase, uint32(v8832)))
	*(*int32)(unsafe.Add(mBase, uint32(v8770)+44)) = v8839
	v8842 = v8832 + int32(4)
	if v8782&int32(128) == int32(0) {
		v8856 = v8782
		v8857 = v8842
		goto L1944
	} else {
		goto L1946
	}
L1946:
	;
	v8850 = F_strlcpy(m, v6731+int32(608), v8842, int32(200))
	mBase = m.M
	v8851 = F_strlen(m, v8842)
	mBase = m.M
	v8855 = *(*int32)(unsafe.Add(mBase, uint32(v8770)+8))
	v8856 = v8855
	v8857 = v8851 + v8842 + int32(1)
	goto L1944
L1947:
	;
	v8862 = *(*int64)(unsafe.Add(mBase, uint32(v8857)))
	v8863 = *(*int64)(unsafe.Add(mBase, uint32(v8857)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8770)+256)) = v8863
	*(*int64)(unsafe.Add(mBase, uint32(v8770)+248)) = v8862
	goto L1930
L1948:
	;
	v8879 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[117])))
	if v8879&int32(1) == int32(0) {
		v8949 = v8875
		goto L1889
	} else {
		goto L1949
	}
L1949:
	;
	v8885 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	if v8873 != v8885 {
		v8949 = v8875
		goto L1889
	} else {
		goto L1950
	}
L1950:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v8873
	v8890 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v8890)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = v8643
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = int64(0)
	v8898 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v8898)
	switch v8604 {
	case 0, 3:
		goto L1952
	default:
		goto L1490
	case 2, 4:
		goto L1951
	}
L1951:
	;
	v8926 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8927 = m.ExcPending
	if v8927 != 0 {
		goto L32
	} else {
		goto L1958
	}
L1952:
	;
	v8902 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8903 = m.ExcPending
	if v8903 != 0 {
		goto L32
	} else {
		goto L1953
	}
L1953:
	;
	if v8902 == int32(0) {
		goto L1490
	} else {
		goto L1954
	}
L1954:
	;
	v8907 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	v8909 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v8910 = F_timestamptz_to_str(m, v8909)
	mBase = m.M
	v8911 = m.ExcPending
	if v8911 != 0 {
		goto L32
	} else {
		goto L1955
	}
L1955:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+52)) = v8910
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+48)) = v8907
	F_errmsg(m, int32(_a_F_StartupXLOG_238), v6731+int32(48))
	mBase = m.M
	v8918 = m.ExcPending
	if v8918 != 0 {
		goto L32
	} else {
		goto L1956
	}
L1956:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2836), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8923 = m.ExcPending
	if v8923 != 0 {
		goto L32
	} else {
		goto L1957
	}
L1957:
	;
	goto L1490
L1958:
	;
	if v8926 == int32(0) {
		goto L1490
	} else {
		goto L1959
	}
L1959:
	;
	v8931 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	v8933 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v8934 = F_timestamptz_to_str(m, v8933)
	mBase = m.M
	v8935 = m.ExcPending
	if v8935 != 0 {
		goto L32
	} else {
		goto L1960
	}
L1960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+68)) = v8934
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+64)) = v8931
	F_errmsg(m, int32(_a_F_StartupXLOG_239), v6731-int32(-64))
	mBase = m.M
	v8942 = m.ExcPending
	if v8942 != 0 {
		goto L32
	} else {
		goto L1961
	}
L1961:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2844), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8947 = m.ExcPending
	if v8947 != 0 {
		goto L32
	} else {
		goto L1962
	}
L1962:
	;
	goto L1490
L1963:
	;
	v8955 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v8955&int32(1) == int32(0) {
		goto L1676
	} else {
		goto L1964
	}
L1964:
	;
	v8962 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8963 = m.ExcPending
	if v8963 != 0 {
		goto L32
	} else {
		goto L1965
	}
L1965:
	;
	if v8962 != 0 {
		goto L1966
	} else {
		goto L1967
	}
L1966:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_222), int32(0))
	mBase = m.M
	v8967 = m.ExcPending
	if v8967 != 0 {
		goto L32
	} else {
		goto L1969
	}
L1967:
	;
	goto L1968
L1968:
	;
	v8974 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v8974)
	v8977 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = v8977
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = v8977
	v8983 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v8983
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[108])) = uint8(v8983)
	goto L1490
L1969:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2854), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8972 = m.ExcPending
	if v8972 != 0 {
		goto L32
	} else {
		goto L1970
	}
L1970:
	;
	goto L1968
L1971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+244)) = v6979
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+240)) = v7808
	F_errmsg(m, int32(_a_F_StartupXLOG_240), v6731+int32(240))
	mBase = m.M
	v8998 = m.ExcPending
	if v8998 != 0 {
		goto L32
	} else {
		goto L1972
	}
L1972:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2370), int32(_a_F_StartupXLOG_241))
	mBase = m.M
	v9003 = m.ExcPending
	if v9003 != 0 {
		goto L32
	} else {
		goto L1973
	}
L1973:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1974:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+212)) = v6979
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+208)) = v7806
	F_errmsg(m, int32(_a_F_StartupXLOG_242), v6731+int32(208))
	mBase = m.M
	v9015 = m.ExcPending
	if v9015 != 0 {
		goto L32
	} else {
		goto L1975
	}
L1975:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2379), int32(_a_F_StartupXLOG_241))
	mBase = m.M
	v9020 = m.ExcPending
	if v9020 != 0 {
		goto L32
	} else {
		goto L1976
	}
L1976:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1977:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+224)) = v7806
	v9027 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+232)) = uint32(v9027)
	v9030 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+236)) = v9030
	v9033 = int64(base.Ui64(v9027) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+228)) = uint32(v9033)
	F_errmsg(m, int32(_a_F_StartupXLOG_243), v6731+int32(224))
	mBase = m.M
	v9039 = m.ExcPending
	if v9039 != 0 {
		goto L32
	} else {
		goto L1978
	}
L1978:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2397), int32(_a_F_StartupXLOG_241))
	mBase = m.M
	v9044 = m.ExcPending
	if v9044 != 0 {
		goto L32
	} else {
		goto L1979
	}
L1979:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1980:
	;
	v9049 = *(*int64)(unsafe.Add(mBase, uint32(v7788)+64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+188)) = uint32(v9049)
	v9051 = int64(32)
	v9052 = int64(base.Ui64(v9049) >> (uint(v9051) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+184)) = uint32(v9052)
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+180)) = uint32(v7906)
	v9056 = int64(base.Ui64(v7906) >> (uint(v9051) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+176)) = uint32(v9056)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_244), v6731+int32(176))
	mBase = m.M
	v9062 = m.ExcPending
	if v9062 != 0 {
		goto L32
	} else {
		goto L1981
	}
L1981:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2072), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v9067 = m.ExcPending
	if v9067 != 0 {
		goto L32
	} else {
		goto L1982
	}
L1982:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1983:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v9074 = m.ExcPending
	if v9074 != 0 {
		goto L32
	} else {
		goto L1984
	}
L1984:
	;
	v9075 = *(*int32)(unsafe.Add(mBase, uint32(v7788)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+128)) = v9075
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_245), v6731+int32(128))
	mBase = m.M
	v9081 = m.ExcPending
	if v9081 != 0 {
		goto L32
	} else {
		goto L1985
	}
L1985:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2527), int32(_a_F_StartupXLOG_246))
	mBase = m.M
	v9086 = m.ExcPending
	if v9086 != 0 {
		goto L32
	} else {
		goto L1986
	}
L1986:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1987:
	;
	v9091 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+552))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+112)) = v9091
	v9093 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+96)) = v9093
	v9095 = *(*int64)(unsafe.Add(mBase, uint32(v6731)+564))
	*(*int64)(unsafe.Add(mBase, uint32(v6731)+100)) = v9095
	v9097 = *(*int32)(unsafe.Add(mBase, uint32(v6731)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+108)) = v9097
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_247), v6731+int32(96))
	mBase = m.M
	v9103 = m.ExcPending
	if v9103 != 0 {
		goto L32
	} else {
		goto L1988
	}
L1988:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2545), int32(_a_F_StartupXLOG_246))
	mBase = m.M
	v9108 = m.ExcPending
	if v9108 != 0 {
		goto L32
	} else {
		goto L1989
	}
L1989:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1990:
	;
	if v9122 != 0 {
		v6979 = v7863
		v6990 = v9122
		goto L1530
	} else {
		goto L1991
	}
L1991:
	;
	goto L1531
L1992:
	;
	if v9126 == int32(0) {
		v9444 = v6867
		goto L1488
	} else {
		goto L1993
	}
L1993:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_248), int32(0))
	mBase = m.M
	v9133 = m.ExcPending
	if v9133 != 0 {
		goto L32
	} else {
		goto L1994
	}
L1994:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1873), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v9138 = m.ExcPending
	if v9138 != 0 {
		goto L32
	} else {
		goto L1995
	}
L1995:
	;
	v9444 = v6867
	goto L1488
L1996:
	;
	if v7420 != 0 {
		goto L1997
	} else {
		goto L1998
	}
L1997:
	;
	if v9157 == int32(0) {
		goto L1490
	} else {
		goto L2000
	}
L1998:
	;
	goto L1999
L1999:
	;
	if v9157 == int32(0) {
		goto L1490
	} else {
		goto L2004
	}
L2000:
	;
	v9162 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	v9164 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v9165 = F_timestamptz_to_str(m, v9164)
	mBase = m.M
	v9166 = m.ExcPending
	if v9166 != 0 {
		goto L32
	} else {
		goto L2001
	}
L2001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+276)) = v9165
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+272)) = v9162
	F_errmsg(m, int32(_a_F_StartupXLOG_249), v6731+int32(272))
	mBase = m.M
	v9173 = m.ExcPending
	if v9173 != 0 {
		goto L32
	} else {
		goto L2002
	}
L2002:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2691), int32(_a_F_StartupXLOG_223))
	mBase = m.M
	v9178 = m.ExcPending
	if v9178 != 0 {
		goto L32
	} else {
		goto L2003
	}
L2003:
	;
	goto L1490
L2004:
	;
	v9182 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	v9184 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v9185 = F_timestamptz_to_str(m, v9184)
	mBase = m.M
	v9186 = m.ExcPending
	if v9186 != 0 {
		goto L32
	} else {
		goto L2005
	}
L2005:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+292)) = v9185
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+288)) = v9182
	F_errmsg(m, int32(_a_F_StartupXLOG_250), v6731+int32(288))
	mBase = m.M
	v9193 = m.ExcPending
	if v9193 != 0 {
		goto L32
	} else {
		goto L2006
	}
L2006:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2698), int32(_a_F_StartupXLOG_223))
	mBase = m.M
	v9198 = m.ExcPending
	if v9198 != 0 {
		goto L32
	} else {
		goto L2007
	}
L2007:
	;
	goto L1490
L2008:
	;
	v9239 = int32(1)
	v9241 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[19]))
	switch v9241 {
	case 0:
		goto L2009
	default:
		v9280 = v9239
		goto L1489
	case 2:
		goto L2010
	}
L2009:
	;
	v9246 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9249 = base.AtomicRmwXchg32(m, v9246, int32(96), int32(1))
	if v9249 != 0 {
		goto L2012
	} else {
		goto L2013
	}
L2010:
	;
	F_proc_exit(m, int32(3))
	mBase = m.M
	v9244 = m.ExcPending
	if v9244 != 0 {
		goto L32
	} else {
		goto L2011
	}
L2011:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2012:
	;
	F_s_lock(m, v9246+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v9254 = m.ExcPending
	if v9254 != 0 {
		goto L32
	} else {
		goto L2015
	}
L2013:
	;
	goto L2014
L2014:
	;
	v9256 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9257 = *(*int32)(unsafe.Add(mBase, uint32(v9256)+80))
	if v9257 == int32(0) {
		goto L2016
	} else {
		goto L2017
	}
L2015:
	;
	goto L2014
L2016:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9256)+80)) = int32(1)
	goto L2018
L2017:
	;
	goto L2018
L2018:
	;
	v9262 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9256)+96)), uint32(v9262))
	F_recoveryPausesHere(m, int32(1))
	mBase = m.M
	v9267 = m.ExcPending
	if v9267 != 0 {
		goto L32
	} else {
		goto L2019
	}
L2019:
	;
	v9280 = v9239
	goto L1489
L2020:
	;
	v9342 = v9313 << (uint(int32(5)) % 32)
	v9345 = *(*int32)(unsafe.Add(mBase, uint32(v9342)+uint32(_c_F_StartupXLOG[110])))
	if v9345 == int32(0) {
		goto L2022
	} else {
		goto L2023
	}
L2021:
	;
	v9369 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9370 = m.ExcPending
	if v9370 != 0 {
		goto L32
	} else {
		goto L2031
	}
L2022:
	;
	v9354 = *(*int32)(unsafe.Add(mBase, uint32(v9342)+uint32(_c_F_StartupXLOG[111])))
	if v9354 == int32(0) {
		goto L2026
	} else {
		goto L2027
	}
L2023:
	;
	v9348 = *(*int32)(unsafe.Add(mBase, uint32(v9342)+uint32(_c_F_StartupXLOG[125])))
	if v9348 == int32(0) {
		goto L2022
	} else {
		goto L2024
	}
L2024:
	;
	m.T0[v9348].(func(*base.Module))(m)
	mBase = m.M
	v9352 = m.ExcPending
	if v9352 != 0 {
		goto L32
	} else {
		goto L2025
	}
L2025:
	;
	goto L2022
L2026:
	;
	v9364 = v9313 + int32(2)
	if v9364 != int32(256) {
		v9313 = v9364
		goto L2020
	} else {
		goto L2030
	}
L2027:
	;
	v9357 = *(*int32)(unsafe.Add(mBase, uint32(v9342)+uint32(_c_F_StartupXLOG[126])))
	if v9357 == int32(0) {
		goto L2026
	} else {
		goto L2028
	}
L2028:
	;
	m.T0[v9357].(func(*base.Module))(m)
	mBase = m.M
	v9361 = m.ExcPending
	if v9361 != 0 {
		goto L32
	} else {
		goto L2029
	}
L2029:
	;
	goto L2026
L2030:
	;
	goto L2021
L2031:
	;
	if v9369 != 0 {
		goto L2032
	} else {
		goto L2033
	}
L2032:
	;
	v9372 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v9373 = *(*int64)(unsafe.Add(mBase, uint32(v9372)+32))
	v9376 = F_pg_rusage_show(m, v6731+int32(368))
	mBase = m.M
	v9377 = m.ExcPending
	if v9377 != 0 {
		goto L32
	} else {
		goto L2035
	}
L2033:
	;
	goto L2034
L2034:
	;
	v9395 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9398 = base.AtomicRmwXchg32(m, v9395, int32(96), int32(1))
	if v9398 != 0 {
		goto L2038
	} else {
		goto L2039
	}
L2035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731)+24)) = v9376
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+20)) = uint32(v9373)
	v9381 = int64(base.Ui64(v9373) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6731)+16)) = uint32(v9381)
	F_errmsg(m, int32(_a_F_StartupXLOG_251), v6731+int32(16))
	mBase = m.M
	v9387 = m.ExcPending
	if v9387 != 0 {
		goto L32
	} else {
		goto L2036
	}
L2036:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1860), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v9392 = m.ExcPending
	if v9392 != 0 {
		goto L32
	} else {
		goto L2037
	}
L2037:
	;
	goto L2034
L2038:
	;
	F_s_lock(m, v9395+int32(96), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v9403 = m.ExcPending
	if v9403 != 0 {
		goto L32
	} else {
		goto L2041
	}
L2039:
	;
	goto L2040
L2040:
	;
	v9405 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9406 = *(*int64)(unsafe.Add(mBase, uint32(v9405)+64))
	v9407 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9405)+96)), uint32(v9407))
	if v9406 == int64(0) {
		goto L2042
	} else {
		goto L2043
	}
L2041:
	;
	goto L2040
L2042:
	;
	v9430 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[109])) = uint8(v9430)
	v9444 = v9280
	goto L1488
L2043:
	;
	v9414 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9415 = m.ExcPending
	if v9415 != 0 {
		goto L32
	} else {
		goto L2044
	}
L2044:
	;
	if v9414 == int32(0) {
		goto L2042
	} else {
		goto L2045
	}
L2045:
	;
	v9418 = F_timestamptz_to_str(m, v9406)
	mBase = m.M
	v9419 = m.ExcPending
	if v9419 != 0 {
		goto L32
	} else {
		goto L2046
	}
L2046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6731))) = v9418
	F_errmsg(m, int32(_a_F_StartupXLOG_252), v6731)
	mBase = m.M
	v9423 = m.ExcPending
	if v9423 != 0 {
		goto L32
	} else {
		goto L2047
	}
L2047:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1865), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v9428 = m.ExcPending
	if v9428 != 0 {
		goto L32
	} else {
		goto L2048
	}
L2048:
	;
	goto L2042
L2049:
	;
	m.G0 = v6731 + int32(848)
	goto L1485
L2050:
	;
	v9469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v9469&int32(1) == int32(0) {
		goto L2049
	} else {
		goto L2051
	}
L2051:
	;
	v9475 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	if v9475 != 0 {
		goto L1486
	} else {
		goto L2052
	}
L2052:
	;
	goto L2049
L2053:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_253), int32(0))
	mBase = m.M
	v9486 = m.ExcPending
	if v9486 != 0 {
		goto L32
	} else {
		goto L2054
	}
L2054:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1825), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v9491 = m.ExcPending
	if v9491 != 0 {
		goto L32
	} else {
		goto L2055
	}
L2055:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2056:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v9498 = m.ExcPending
	if v9498 != 0 {
		goto L32
	} else {
		goto L2057
	}
L2057:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_254), int32(0))
	mBase = m.M
	v9502 = m.ExcPending
	if v9502 != 0 {
		goto L32
	} else {
		goto L2058
	}
L2058:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1885), int32(_a_F_StartupXLOG_218))
	mBase = m.M
	v9507 = m.ExcPending
	if v9507 != 0 {
		goto L32
	} else {
		goto L2059
	}
L2059:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2060:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v9553 = m.ExcPending
	if v9553 != 0 {
		goto L32
	} else {
		goto L2061
	}
L2061:
	;
	v9556 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	v9559 = base.AtomicRmwXchg32(m, v9556, int32(16), int32(1))
	if v9559 != 0 {
		goto L2062
	} else {
		goto L2063
	}
L2062:
	;
	F_s_lock(m, v9556+int32(16), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v9564 = m.ExcPending
	if v9564 != 0 {
		goto L32
	} else {
		goto L2065
	}
L2063:
	;
	goto L2064
L2064:
	;
	v9566 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	v9567 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9566)+4)) = uint8(v9567)
	v9569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9566)+5)))
	if v9569 != v9567 {
		v9664 = v9566
		goto L2066
	} else {
		goto L2067
	}
L2065:
	;
	goto L2064
L2066:
	;
	v9697 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9664)+16)), uint32(v9697))
	v9702 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v9702 == int32(1) {
		goto L2085
	} else {
		goto L2086
	}
L2067:
	;
	v9572 = *(*int32)(unsafe.Add(mBase, uint32(v9566)))
	v9573 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9566)+16)), uint32(v9573))
	if v9572 != int32(-1) {
		goto L2068
	} else {
		goto L2069
	}
L2068:
	;
	v9580 = F_SendProcSignal(m, v9572, int32(7), int32(-1))
	mBase = m.M
	v9581 = m.ExcPending
	if v9581 != 0 {
		goto L32
	} else {
		goto L2071
	}
L2069:
	;
	goto L2070
L2070:
	;
	goto L2072
L2071:
	;
	goto L2070
L2072:
	;
	v9619 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[128]))
	v9623 = F_WaitLatch(m, v9619, int32(41), int32(10), int32(83886092))
	mBase = m.M
	v9624 = m.ExcPending
	if v9624 != 0 {
		goto L32
	} else {
		goto L2075
	}
L2074:
	;
	v9644 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	v9647 = base.AtomicRmwXchg32(m, v9644, int32(16), int32(1))
	if v9647 != 0 {
		goto L2080
	} else {
		goto L2081
	}
L2075:
	;
	if v9623&int32(1) == int32(0) {
		goto L2074
	} else {
		goto L2076
	}
L2076:
	;
	v9630 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[128]))
	v9631 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9630))) = v9631
	v9636 = base.AtomicRmwOr32(m, v9631, int32(_a_F_StartupXLOG_226), v9631)
	goto L2077
L2077:
	;
	v9638 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	if v9638 == int32(0) {
		goto L2074
	} else {
		goto L2078
	}
L2078:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9642 = m.ExcPending
	if v9642 != 0 {
		goto L32
	} else {
		goto L2079
	}
L2079:
	;
	goto L2074
L2080:
	;
	F_s_lock(m, v9644+int32(16), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v9652 = m.ExcPending
	if v9652 != 0 {
		goto L32
	} else {
		goto L2083
	}
L2081:
	;
	goto L2082
L2082:
	;
	v9654 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	v9655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9654)+5)))
	if v9655 != int32(1) {
		v9664 = v9654
		goto L2066
	} else {
		goto L2084
	}
L2083:
	;
	goto L2082
L2084:
	;
	v9658 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9654)+16)), uint32(v9658))
	goto L2072
L2085:
	;
	v9706 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v9710 = F_LWLockAcquire(m, v9706+int32(_a_F_StartupXLOG_255), int32(1))
	mBase = m.M
	v9711 = m.ExcPending
	if v9711 != 0 {
		goto L32
	} else {
		goto L2088
	}
L2086:
	;
	goto L2087
L2087:
	;
	v9893 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v9893)
	v9896 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9901 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])))
	if v9901 != 0 {
		goto L2110
	} else {
		goto L2111
	}
L2088:
	;
	v9713 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v9715 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v9713+v9715 {
		goto L2089
	} else {
		goto L2090
	}
L2089:
	;
	v9720 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9730 = v9715
	v9733 = v9713
	v9734 = v9697
	v9735 = v9720
	v9749 = int64(0)
	goto L2092
L2090:
	;
	goto L2091
L2091:
	;
	v9851 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v9851+int32(_a_F_StartupXLOG_255))
	mBase = m.M
	v9855 = m.ExcPending
	if v9855 != 0 {
		goto L32
	} else {
		goto L2109
	}
L2092:
	;
	v9759 = v9735 + v9734*int32(296)
	v9760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9759)+4)))
	if v9760 != int32(1) {
		v9806 = v9730
		v9807 = v9733
		v9808 = v9735
		v9809 = v9749
		goto L2094
	} else {
		goto L2095
	}
L2093:
	;
	goto L2091
L2094:
	;
	v9811 = v9734 + int32(1)
	if v9811 < v9806+v9807 {
		v9730 = v9806
		v9733 = v9807
		v9734 = v9811
		v9735 = v9808
		v9749 = v9809
		goto L2092
	} else {
		goto L2108
	}
L2095:
	;
	v9763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9759)+201)))
	if v9763 != int32(1) {
		v9806 = v9730
		v9807 = v9733
		v9808 = v9735
		v9809 = v9749
		goto L2094
	} else {
		goto L2096
	}
L2096:
	;
	if v9749 == int64(0) {
		goto L2097
	} else {
		goto L2098
	}
L2097:
	;
	v9771 = m.G0
	v9772 = int32(16)
	v9773 = v9771 - v9772
	m.G0 = v9773
	F_gettimeofday(m, v9773)
	mBase = m.M
	v9776 = *(*int64)(unsafe.Add(mBase, uint32(v9773)))
	v9777 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9773)+8)))
	m.G0 = v9773 + v9772
	goto L2100
L2098:
	;
	v9786 = v9749
	goto L2099
L2099:
	;
	v9789 = base.AtomicRmwXchg32(m, v9759, int32(0), int32(1))
	if v9789 != 0 {
		goto L2101
	} else {
		goto L2102
	}
L2100:
	;
	v9786 = v9777 + v9776*int64(1000000) - int64(946684800000000)
	goto L2099
L2101:
	;
	F_s_lock(m, v9759, int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v9792 = m.ExcPending
	if v9792 != 0 {
		goto L32
	} else {
		goto L2104
	}
L2102:
	;
	goto L2103
L2103:
	;
	v9793 = *(*int32)(unsafe.Add(mBase, uint32(v9759)+112))
	if v9793 == int32(0) {
		goto L2105
	} else {
		goto L2106
	}
L2104:
	;
	goto L2103
L2105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9759)+272)) = v9786
	goto L2107
L2106:
	;
	goto L2107
L2107:
	;
	v9797 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9759))), uint32(v9797))
	v9801 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v9803 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	v9805 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9806 = v9803
	v9807 = v9801
	v9808 = v9805
	v9809 = v9786
	goto L2094
L2108:
	;
	goto L2093
L2109:
	;
	goto L2087
L2110:
	;
	v9902 = v9896 + int32(40)
	goto L2112
L2111:
	;
	v9902 = int32(_a_F_StartupXLOG_216)
	goto L2112
L2112:
	;
	v9903 = *(*int32)(unsafe.Add(mBase, uint32(v9902)))
	v9905 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	if v9901 != 0 {
		goto L2113
	} else {
		goto L2114
	}
L2113:
	;
	v9909 = v9896 + int32(24)
	goto L2115
L2114:
	;
	v9909 = int32(_a_F_StartupXLOG_256)
	goto L2115
L2115:
	;
	v9910 = *(*int64)(unsafe.Add(mBase, uint32(v9909)))
	F_XLogPrefetcherBeginRead(m, v9905, v9910)
	mBase = m.M
	v9912 = m.ExcPending
	if v9912 != 0 {
		goto L32
	} else {
		goto L2116
	}
L2116:
	;
	v9914 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v9917 = F_ReadRecord(m, v9914, int32(24), int32(0), v9903)
	mBase = m.M
	v9918 = m.ExcPending
	if v9918 != 0 {
		goto L32
	} else {
		goto L2117
	}
L2117:
	;
	v9920 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v9921 = *(*int64)(unsafe.Add(mBase, uint32(v9920)+40))
	v9922 = *(*int32)(unsafe.Add(mBase, uint32(v9920)+1184))
	*(*int32)(unsafe.Add(mBase, uint32(v9550)+24)) = v9922
	v9925 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v9925 != int32(1) {
		goto L2118
	} else {
		goto L2119
	}
L2118:
	;
	v9941 = v9921 & int64(8191)
	if v9941 != int64(0) {
		goto L2121
	} else {
		goto L2122
	}
L2119:
	;
	v9929 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v9929)
	v9932 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129]))
	if v9932 < v9929 {
		goto L2118
	} else {
		goto L2120
	}
L2120:
	;
	v9935 = F_close(m, v9932)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129])) = int32(-1)
	goto L2118
L2121:
	;
	v9944 = base.I32_wrap_i64(v9941)
	v9945 = F_palloc(m, v9944)
	mBase = m.M
	v9946 = m.ExcPending
	if v9946 != 0 {
		goto L32
	} else {
		goto L2124
	}
L2122:
	;
	v9954 = int32(0)
	v9955 = v9921
	goto L2123
L2123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9550)+40)) = v9954
	*(*int64)(unsafe.Add(mBase, uint32(v9550)+32)) = v9955
	v9959 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[22]))
	switch v9959 - int32(1) {
	case 0:
		goto L2134
	case 1:
		goto L2133
	case 2:
		goto L2131
	case 3:
		goto L2132
	case 4:
		goto L2130
	default:
		goto L2129
	}
L2124:
	;
	if v9944 != 0 {
		goto L2125
	} else {
		goto L2126
	}
L2125:
	;
	v9948 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v9949 = *(*int32)(unsafe.Add(mBase, uint32(v9948)+128))
	base.MemoryCopy(m, v9945, v9949, v9944)
	goto L2127
L2126:
	;
	goto L2127
L2127:
	;
	v9954 = v9945
	v9955 = v9921 & int64(-8192)
	goto L2123
L2128:
	;
	v10044 = F_pstrdup(m, v9547-int32(-64))
	mBase = m.M
	v10045 = m.ExcPending
	if v10045 != 0 {
		goto L32
	} else {
		goto L2151
	}
L2129:
	;
	v10038 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_257), int32(0))
	mBase = m.M
	v10039 = m.ExcPending
	if v10039 != 0 {
		goto L32
	} else {
		goto L2150
	}
L2130:
	;
	v10031 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_258), int32(0))
	mBase = m.M
	v10032 = m.ExcPending
	if v10032 != 0 {
		goto L32
	} else {
		goto L2149
	}
L2131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+48)) = int32(_a_F_StartupXLOG_234)
	v10024 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_259), v9547+int32(48))
	mBase = m.M
	v10025 = m.ExcPending
	if v10025 != 0 {
		goto L32
	} else {
		goto L2148
	}
L2132:
	;
	v9997 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[106]))
	*(*uint32)(unsafe.Add(mBase, uint32(v9547)+40)) = uint32(v9997)
	v10002 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])))
	if v10002 != 0 {
		goto L2144
	} else {
		goto L2145
	}
L2133:
	;
	v9978 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])))
	v9980 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[105]))
	v9981 = F_timestamptz_to_str(m, v9980)
	mBase = m.M
	v9982 = m.ExcPending
	if v9982 != 0 {
		goto L32
	} else {
		goto L2139
	}
L2134:
	;
	v9963 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+4)) = v9963
	v9968 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])))
	if v9968 != 0 {
		goto L2135
	} else {
		goto L2136
	}
L2135:
	;
	v9969 = int32(_a_F_StartupXLOG_260)
	goto L2137
L2136:
	;
	v9969 = int32(_a_F_StartupXLOG_261)
	goto L2137
L2137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9547))) = v9969
	v9975 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_262), v9547)
	mBase = m.M
	v9976 = m.ExcPending
	if v9976 != 0 {
		goto L32
	} else {
		goto L2138
	}
L2138:
	;
	goto L2128
L2139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+20)) = v9981
	if v9978 != 0 {
		goto L2140
	} else {
		goto L2141
	}
L2140:
	;
	v9986 = int32(_a_F_StartupXLOG_260)
	goto L2142
L2141:
	;
	v9986 = int32(_a_F_StartupXLOG_261)
	goto L2142
L2142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+16)) = v9986
	v9994 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_263), v9547+int32(16))
	mBase = m.M
	v9995 = m.ExcPending
	if v9995 != 0 {
		goto L32
	} else {
		goto L2143
	}
L2143:
	;
	goto L2128
L2144:
	;
	v10003 = int32(_a_F_StartupXLOG_260)
	goto L2146
L2145:
	;
	v10003 = int32(_a_F_StartupXLOG_261)
	goto L2146
L2146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9547)+32)) = v10003
	v10006 = int64(base.Ui64(v9997) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v9547)+36)) = uint32(v10006)
	v10014 = F_pg_snprintf(m, v9547-int32(-64), int32(200), int32(_a_F_StartupXLOG_264), v9547+int32(32))
	mBase = m.M
	v10015 = m.ExcPending
	if v10015 != 0 {
		goto L32
	} else {
		goto L2147
	}
L2147:
	;
	goto L2128
L2148:
	;
	goto L2128
L2149:
	;
	goto L2128
L2150:
	;
	goto L2128
L2151:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9550)+16)) = v9921
	*(*int32)(unsafe.Add(mBase, uint32(v9550)+8)) = v9903
	*(*int64)(unsafe.Add(mBase, uint32(v9550))) = v9910
	*(*int32)(unsafe.Add(mBase, uint32(v9550)+64)) = v10044
	v10051 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52]))
	*(*int64)(unsafe.Add(mBase, uint32(v9550)+48)) = v10051
	v10054 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53]))
	*(*int64)(unsafe.Add(mBase, uint32(v9550)+56)) = v10054
	v10057 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9550)+68)) = uint8(v10057)
	v10060 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9550)+69)) = uint8(v10060)
	m.G0 = v9547 + int32(272)
	v10067 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+24))
	v10070 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])))
	if v10070 == int32(1) {
		goto L2152
	} else {
		goto L2153
	}
L2152:
	;
	v10074 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v10076 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[86]))
	if base.Ui64(v10076) <= base.Ui64(v9921) {
		goto L2156
	} else {
		goto L2157
	}
L2153:
	;
	goto L2154
L2154:
	;
	v10114 = int32(0)
	v10116 = F_PrescanPreparedTransactions(m, v10114, v10114)
	mBase = m.M
	v10117 = m.ExcPending
	if v10117 != 0 {
		goto L32
	} else {
		goto L2171
	}
L2155:
	;
	F_ResetUnloggedRelations(m, int32(2))
	mBase = m.M
	v10112 = m.ExcPending
	if v10112 != 0 {
		goto L32
	} else {
		goto L2170
	}
L2156:
	;
	v10078 = *(*int64)(unsafe.Add(mBase, uint32(v10074)+160))
	if v10078 == int64(0) {
		goto L2155
	} else {
		goto L2159
	}
L2157:
	;
	goto L2158
L2158:
	;
	v10082 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v10082 == int32(0) {
		goto L2160
	} else {
		goto L2161
	}
L2159:
	;
	goto L2158
L2160:
	;
	v10085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10074)+176)))
	if v10085 != int32(1) {
		goto L2155
	} else {
		goto L2163
	}
L2161:
	;
	goto L2162
L2162:
	;
	v10088 = *(*int64)(unsafe.Add(mBase, uint32(v10074)+160))
	if v10088 != int64(0) {
		goto L13
	} else {
		goto L2164
	}
L2163:
	;
	goto L2162
L2164:
	;
	v10091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10074)+176)))
	if v10091 == int32(1) {
		goto L13
	} else {
		goto L2165
	}
L2165:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v10097 = m.ExcPending
	if v10097 != 0 {
		goto L32
	} else {
		goto L2166
	}
L2166:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v10100 = m.ExcPending
	if v10100 != 0 {
		goto L32
	} else {
		goto L2167
	}
L2167:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_265), int32(0))
	mBase = m.M
	v10104 = m.ExcPending
	if v10104 != 0 {
		goto L32
	} else {
		goto L2168
	}
L2168:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_266), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10109 = m.ExcPending
	if v10109 != 0 {
		goto L32
	} else {
		goto L2169
	}
L2169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2170:
	;
	goto L2154
L2171:
	;
	v10119 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v10123 = F_LWLockAcquire(m, v10119+int32(1152), int32(0))
	mBase = m.M
	v10124 = m.ExcPending
	if v10124 != 0 {
		goto L32
	} else {
		goto L2172
	}
L2172:
	;
	v10126 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v10127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10126)+312)) = uint8(v10127)
	v10130 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v10130+int32(1152))
	mBase = m.M
	v10134 = m.ExcPending
	if v10134 != 0 {
		goto L32
	} else {
		goto L2173
	}
L2173:
	;
	v10136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v10136 != int32(1) {
		goto L2175
	} else {
		goto L2176
	}
L2174:
	;
	v11054 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v11057 = base.AtomicRmwXchg32(m, v11054, int32(440), int32(1))
	if v11057 != 0 {
		goto L2363
	} else {
		goto L2364
	}
L2175:
	;
	v10139 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+8))
	v11028 = v10139
	goto L2174
L2176:
	;
	goto L2177
L2177:
	;
	v10141 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	v10142 = F_findNewestTimeLine(m, v10141)
	mBase = m.M
	v10143 = m.ExcPending
	if v10143 != 0 {
		goto L32
	} else {
		goto L2178
	}
L2178:
	;
	v10145 = v10142 + int32(1)
	v10148 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10149 = m.ExcPending
	if v10149 != 0 {
		goto L32
	} else {
		goto L2179
	}
L2179:
	;
	if v10148 != 0 {
		goto L2180
	} else {
		goto L2181
	}
L2180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3696)) = v10145
	F_errmsg(m, int32(_a_F_StartupXLOG_267), v41+int32(3696))
	mBase = m.M
	v10155 = m.ExcPending
	if v10155 != 0 {
		goto L32
	} else {
		goto L2183
	}
L2181:
	;
	goto L2182
L2182:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v10164 = m.ExcPending
	if v10164 != 0 {
		goto L32
	} else {
		goto L2185
	}
L2183:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_268), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10160 = m.ExcPending
	if v10160 != 0 {
		goto L32
	} else {
		goto L2184
	}
L2184:
	;
	goto L2182
L2185:
	;
	v10168 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v10169 = base.I64_extend_i32_s(v10168)
	v10170 = base.I64_div_u_s(v9921-int64(1), v10169)
	v10171 = base.I64_div_u_s(v9921, v10169)
	if v10170 == v10171 {
		goto L2187
	} else {
		goto L2188
	}
L2186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3536)) = v10145
	v10485 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v10486 = base.I64_div_u_s(int64(4294967296), v10485)
	v10487 = base.I64_div_u_s(v10171, v10486)
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3540)) = uint32(v10487)
	v10490 = v10171 - v10487*v10486
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3544)) = uint32(v10490)
	v10493 = v41 + int32(_a_F_StartupXLOG_1)
	v10498 = F_pg_snprintf(m, v10493, int32(64), int32(_a_F_StartupXLOG_269), v41+int32(3536))
	mBase = m.M
	v10499 = m.ExcPending
	if v10499 != 0 {
		goto L32
	} else {
		goto L2249
	}
L2187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3664)) = v10067
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[130]))) = v10170
	v10176 = base.I64_div_u_s(int64(4294967296), v10169)
	v10177 = base.I64_div_u_s(v10170, v10176)
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3668)) = uint32(v10177)
	v10180 = v10170 - v10177*v10176
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3672)) = uint32(v10180)
	v10183 = v41 + int32(_a_F_StartupXLOG_4)
	v10188 = F_pg_snprintf(m, v10183, int32(1024), int32(_a_F_StartupXLOG_270), v41+int32(3664))
	mBase = m.M
	v10189 = m.ExcPending
	if v10189 != 0 {
		goto L32
	} else {
		goto L2190
	}
L2188:
	;
	goto L2189
L2189:
	;
	v10443 = F_XLogFileInit(m, v10171, v10145)
	mBase = m.M
	v10444 = m.ExcPending
	if v10444 != 0 {
		goto L32
	} else {
		goto L2247
	}
L2190:
	;
	v10191 = F_OpenTransientFile(m, v10183, int32(0))
	mBase = m.M
	v10192 = m.ExcPending
	if v10192 != 0 {
		goto L32
	} else {
		goto L2191
	}
L2191:
	;
	if v10191 < int32(0) {
		goto L12
	} else {
		goto L2192
	}
L2192:
	;
	v10195 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3648)) = v10195
	v10198 = v41 + int32(_a_F_StartupXLOG_3)
	v10203 = F_pg_snprintf(m, v10198, int32(1024), int32(_a_F_StartupXLOG_271), v41+int32(3648))
	mBase = m.M
	v10204 = m.ExcPending
	if v10204 != 0 {
		goto L32
	} else {
		goto L2193
	}
L2193:
	;
	v10205 = F_unlink(m, v10198)
	mBase = m.M
	v10208 = F_OpenTransientFile(m, v10198, int32(194))
	mBase = m.M
	v10209 = m.ExcPending
	if v10209 != 0 {
		goto L32
	} else {
		goto L2194
	}
L2194:
	;
	if v10208 < int32(0) {
		goto L11
	} else {
		goto L2195
	}
L2195:
	;
	v10213 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if int32(0) < v10213 {
		goto L2196
	} else {
		goto L2197
	}
L2196:
	;
	v10220 = int32(0)
	goto L2199
L2197:
	;
	goto L2198
L2198:
	;
	v10367 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10367))) = int32(167772233)
	v10372 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v10372 != int32(1) {
		v10386 = int32(0)
		goto L2222
	} else {
		goto L2223
	}
L2199:
	;
	v10256 = base.I32_wrap_i64(v9921)&(v10168-int32(1)) - v10220
	if base.Ui32(v10256) <= base.Ui32(int32(_a_F_StartupXLOG_272)) {
		goto L2201
	} else {
		goto L2202
	}
L2200:
	;
	goto L2198
L2201:
	;
	base.MemoryFill(m, v41+int32(_a_F_StartupXLOG_1), int32(0), int32(_a_F_StartupXLOG_57))
	goto L2203
L2202:
	;
	goto L2203
L2203:
	;
	if int32(0) < v10256 {
		goto L2204
	} else {
		goto L2205
	}
L2204:
	;
	v10267 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10267))) = int32(167772232)
	v10272 = int32(_a_F_StartupXLOG_57)
	if base.Ui32(v10272) <= base.Ui32(v10256) {
		goto L2207
	} else {
		goto L2208
	}
L2205:
	;
	goto L2206
L2206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10312 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10312))) = int32(167772234)
	v10317 = int32(_a_F_StartupXLOG_57)
	v10318 = F_write(m, v10208, v41+int32(_a_F_StartupXLOG_1), v10317)
	mBase = m.M
	if v10318 != v10317 {
		goto L9
	} else {
		goto L2218
	}
L2207:
	;
	v10275 = v10272
	goto L2209
L2208:
	;
	v10275 = v10256
	goto L2209
L2209:
	;
	v10276 = F_read(m, v10191, v41+int32(_a_F_StartupXLOG_1), v10275)
	mBase = m.M
	if v10276 != v10275 {
		goto L2210
	} else {
		goto L2211
	}
L2210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10281 = m.ExcPending
	if v10281 != 0 {
		goto L32
	} else {
		goto L2213
	}
L2211:
	;
	goto L2212
L2212:
	;
	v10303 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10303))) = int32(0)
	goto L2206
L2213:
	;
	if v10276 < int32(0) {
		goto L10
	} else {
		goto L2214
	}
L2214:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v10286 = m.ExcPending
	if v10286 != 0 {
		goto L32
	} else {
		goto L2215
	}
L2215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3640)) = v10275
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3636)) = v10276
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3632)) = v41 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_143), v41+int32(3632))
	mBase = m.M
	v10296 = m.ExcPending
	if v10296 != 0 {
		goto L32
	} else {
		goto L2216
	}
L2216:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3549), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v10301 = m.ExcPending
	if v10301 != 0 {
		goto L32
	} else {
		goto L2217
	}
L2217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2218:
	;
	v10322 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10322))) = int32(0)
	v10326 = v10220 - int32(-8192)
	v10328 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v10326 < v10328 {
		v10220 = v10326
		goto L2199
	} else {
		goto L2219
	}
L2219:
	;
	goto L2200
L2220:
	;
	v10415 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10415))) = int32(0)
	v10418 = F_CloseTransientFile(m, v10208)
	mBase = m.M
	v10419 = m.ExcPending
	if v10419 != 0 {
		goto L32
	} else {
		goto L2238
	}
L2221:
	;
	if v10386 == int32(0) {
		goto L2220
	} else {
		goto L2228
	}
L2222:
	;
	goto L2221
L2223:
	;
	goto L2224
L2224:
	;
	v10377 = F_fsync(m, v10208)
	mBase = m.M
	if v10377 != int32(-1) {
		v10386 = v10377
		goto L2222
	} else {
		goto L2226
	}
L2225:
	;
	v10386 = int32(-1)
	goto L2222
L2226:
	;
	v10381 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10381 == int32(27) {
		goto L2224
	} else {
		goto L2227
	}
L2227:
	;
	goto L2225
L2228:
	;
	v10392 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[131])))
	if v10392 != 0 {
		goto L2230
	} else {
		goto L2231
	}
L2229:
	;
	v10395 = F_errstart(m, v10393, int32(0))
	mBase = m.M
	v10396 = m.ExcPending
	if v10396 != 0 {
		goto L32
	} else {
		goto L2233
	}
L2230:
	;
	v10393 = int32(21)
	goto L2232
L2231:
	;
	v10393 = int32(24)
	goto L2232
L2232:
	;
	goto L2229
L2233:
	;
	if v10395 == int32(0) {
		goto L2220
	} else {
		goto L2234
	}
L2234:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10400 = m.ExcPending
	if v10400 != 0 {
		goto L32
	} else {
		goto L2235
	}
L2235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3584)) = v41 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v41+int32(3584))
	mBase = m.M
	v10408 = m.ExcPending
	if v10408 != 0 {
		goto L32
	} else {
		goto L2236
	}
L2236:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3577), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v10413 = m.ExcPending
	if v10413 != 0 {
		goto L32
	} else {
		goto L2237
	}
L2237:
	;
	goto L2220
L2238:
	;
	if v10418 != 0 {
		goto L8
	} else {
		goto L2239
	}
L2239:
	;
	v10420 = F_CloseTransientFile(m, v10191)
	mBase = m.M
	v10421 = m.ExcPending
	if v10421 != 0 {
		goto L32
	} else {
		goto L2240
	}
L2240:
	;
	if v10420 != 0 {
		goto L7
	} else {
		goto L2241
	}
L2241:
	;
	v10428 = F_InstallXLogFileSegment(m, v41+int32(_a_F_StartupXLOG_274), v41+int32(_a_F_StartupXLOG_3), int32(0), int64(0), v10145)
	mBase = m.M
	v10429 = m.ExcPending
	if v10429 != 0 {
		goto L32
	} else {
		goto L2242
	}
L2242:
	;
	if v10428 != 0 {
		goto L2186
	} else {
		goto L2243
	}
L2243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10433 = m.ExcPending
	if v10433 != 0 {
		goto L32
	} else {
		goto L2244
	}
L2244:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_275), int32(0))
	mBase = m.M
	v10437 = m.ExcPending
	if v10437 != 0 {
		goto L32
	} else {
		goto L2245
	}
L2245:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3594), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v10442 = m.ExcPending
	if v10442 != 0 {
		goto L32
	} else {
		goto L2246
	}
L2246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2247:
	;
	v10445 = F_close(m, v10443)
	mBase = m.M
	if v10445 != 0 {
		goto L6
	} else {
		goto L2248
	}
L2248:
	;
	goto L2186
L2249:
	;
	F_XLogArchiveCleanup(m, v10493)
	mBase = m.M
	v10501 = m.ExcPending
	if v10501 != 0 {
		goto L32
	} else {
		goto L2250
	}
L2250:
	;
	v10502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9550)+68)))
	if v10502 == int32(1) {
		goto L2251
	} else {
		goto L2252
	}
L2251:
	;
	v10507 = F_durable_unlink(m, int32(_a_F_StartupXLOG_47), int32(22))
	mBase = m.M
	v10508 = m.ExcPending
	if v10508 != 0 {
		goto L32
	} else {
		goto L2254
	}
L2252:
	;
	goto L2253
L2253:
	;
	v10509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9550)+69)))
	if v10509 == int32(1) {
		goto L2255
	} else {
		goto L2256
	}
L2254:
	;
	goto L2253
L2255:
	;
	v10514 = F_durable_unlink(m, int32(_a_F_StartupXLOG_48), int32(22))
	mBase = m.M
	v10515 = m.ExcPending
	if v10515 != 0 {
		goto L32
	} else {
		goto L2258
	}
L2256:
	;
	goto L2257
L2257:
	;
	v10517 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[13]))
	v10518 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+64))
	v10519 = m.G0
	v10521 = v10519 - int32(_a_F_StartupXLOG_276)
	m.G0 = v10521
	v10523 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+224)) = v10523
	v10526 = v10521 + int32(_a_F_StartupXLOG_277)
	v10531 = F_pg_snprintf(m, v10526, int32(1024), int32(_a_F_StartupXLOG_271), v10521+int32(224))
	mBase = m.M
	v10532 = m.ExcPending
	if v10532 != 0 {
		goto L32
	} else {
		goto L2259
	}
L2258:
	;
	goto L2257
L2259:
	;
	v10533 = F_unlink(m, v10526)
	mBase = m.M
	v10535 = F_OpenTransientFile(m, v10526, int32(194))
	mBase = m.M
	v10536 = m.ExcPending
	if v10536 != 0 {
		goto L32
	} else {
		goto L2266
	}
L2260:
	;
	v11004 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11005 = m.ExcPending
	if v11005 != 0 {
		goto L32
	} else {
		goto L2359
	}
L2261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10986 = m.ExcPending
	if v10986 != 0 {
		goto L32
	} else {
		goto L2355
	}
L2262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+112)) = v10518
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+100)) = v10517
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+96)) = v10758
	*(*uint32)(unsafe.Add(mBase, uint32(v10521)+108)) = uint32(v9921)
	v10791 = int64(base.Ui64(v9921) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10521)+104)) = uint32(v10791)
	v10794 = v10521 + int32(240)
	v10799 = F_pg_snprintf(m, v10794, int32(_a_F_StartupXLOG_57), int32(_a_F_StartupXLOG_278), v10521+int32(96))
	mBase = m.M
	v10800 = m.ExcPending
	if v10800 != 0 {
		goto L32
	} else {
		goto L2312
	}
L2263:
	;
	v10727 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10727 == int32(44) {
		goto L2305
	} else {
		goto L2306
	}
L2264:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10710 = m.ExcPending
	if v10710 != 0 {
		goto L32
	} else {
		goto L2301
	}
L2265:
	;
	v10681 = int32(_a_F_StartupXLOG_2)
	v10682 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10684 = v10521 + int32(_a_F_StartupXLOG_277)
	v10685 = F_unlink(m, v10684)
	mBase = m.M
	if v10682 != 0 {
		goto L2294
	} else {
		goto L2295
	}
L2266:
	;
	if int32(0) <= v10535 {
		goto L2267
	} else {
		goto L2268
	}
L2267:
	;
	v10540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v10540 == int32(1) {
		goto L2271
	} else {
		goto L2272
	}
L2268:
	;
	goto L2269
L2269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10667 = m.ExcPending
	if v10667 != 0 {
		goto L32
	} else {
		goto L2290
	}
L2270:
	;
	v10572 = F_OpenTransientFile(m, v10521+int32(_a_F_StartupXLOG_279), int32(0))
	mBase = m.M
	v10573 = m.ExcPending
	if v10573 != 0 {
		goto L32
	} else {
		goto L2277
	}
L2271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+192)) = v10517
	v10545 = v10521 + int32(_a_F_StartupXLOG_280)
	v10550 = F_pg_snprintf(m, v10545, int32(64), int32(_a_F_StartupXLOG_281), v10521+int32(192))
	mBase = m.M
	v10551 = m.ExcPending
	if v10551 != 0 {
		goto L32
	} else {
		goto L2274
	}
L2272:
	;
	goto L2273
L2273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+208)) = v10517
	v10566 = F_pg_snprintf(m, v10521+int32(_a_F_StartupXLOG_279), int32(1024), int32(_a_F_StartupXLOG_282), v10521+int32(208))
	mBase = m.M
	v10567 = m.ExcPending
	if v10567 != 0 {
		goto L32
	} else {
		goto L2276
	}
L2274:
	;
	v10557 = F_RestoreArchivedFile(m, v10521+int32(_a_F_StartupXLOG_279), v10545, int32(_a_F_StartupXLOG_283), int64(0), int32(0))
	mBase = m.M
	v10558 = m.ExcPending
	if v10558 != 0 {
		goto L32
	} else {
		goto L2275
	}
L2275:
	;
	goto L2270
L2276:
	;
	goto L2270
L2277:
	;
	if v10572 < int32(0) {
		goto L2263
	} else {
		goto L2278
	}
L2278:
	;
	v10577 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10577
	v10579 = int32(_a_F_StartupXLOG_142)
	v10580 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10580))) = int32(167772221)
	v10586 = F_read(m, v10572, v10521+int32(240), int32(_a_F_StartupXLOG_57))
	mBase = m.M
	v10588 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10588))) = v10577
	if v10586 < v10577 {
		goto L2261
	} else {
		goto L2279
	}
L2279:
	;
	v10601 = v10586
	goto L2280
L2280:
	;
	v10630 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10630 != 0 {
		goto L2261
	} else {
		goto L2282
	}
L2281:
	;
	v10661 = F_CloseTransientFile(m, v10572)
	mBase = m.M
	v10662 = m.ExcPending
	if v10662 != 0 {
		goto L32
	} else {
		goto L2288
	}
L2282:
	;
	if v10601 != 0 {
		goto L2283
	} else {
		goto L2284
	}
L2283:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10635 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10635))) = int32(167772223)
	v10639 = v10521 + int32(240)
	v10640 = F_write(m, v10535, v10639, v10601)
	mBase = m.M
	if v10640 != v10601 {
		goto L2265
	} else {
		goto L2286
	}
L2284:
	;
	goto L2285
L2285:
	;
	goto L2281
L2286:
	;
	v10642 = int32(_a_F_StartupXLOG_142)
	v10643 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	v10644 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10643))) = v10644
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10644
	v10650 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10650))) = int32(167772221)
	v10654 = F_read(m, v10572, v10639, int32(_a_F_StartupXLOG_57))
	mBase = m.M
	v10656 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10656))) = v10644
	if v10644 <= v10654 {
		v10601 = v10654
		goto L2280
	} else {
		goto L2287
	}
L2287:
	;
	goto L2261
L2288:
	;
	if v10661 != 0 {
		goto L2264
	} else {
		goto L2289
	}
L2289:
	;
	v10758 = int32(_a_F_StartupXLOG_284)
	goto L2262
L2290:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10669 = m.ExcPending
	if v10669 != 0 {
		goto L32
	} else {
		goto L2291
	}
L2291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521))) = v10521 + int32(_a_F_StartupXLOG_277)
	F_errmsg(m, int32(_a_F_StartupXLOG_285), v10521)
	mBase = m.M
	v10675 = m.ExcPending
	if v10675 != 0 {
		goto L32
	} else {
		goto L2292
	}
L2292:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(330), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10680 = m.ExcPending
	if v10680 != 0 {
		goto L32
	} else {
		goto L2293
	}
L2293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2294:
	;
	v10688 = v10682
	goto L2296
L2295:
	;
	v10688 = int32(51)
	goto L2296
L2296:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10688
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10693 = m.ExcPending
	if v10693 != 0 {
		goto L32
	} else {
		goto L2297
	}
L2297:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10695 = m.ExcPending
	if v10695 != 0 {
		goto L32
	} else {
		goto L2298
	}
L2298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+176)) = v10684
	F_errmsg(m, int32(_a_F_StartupXLOG_288), v10521+int32(176))
	mBase = m.M
	v10701 = m.ExcPending
	if v10701 != 0 {
		goto L32
	} else {
		goto L2299
	}
L2299:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(385), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10706 = m.ExcPending
	if v10706 != 0 {
		goto L32
	} else {
		goto L2300
	}
L2300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2301:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10712 = m.ExcPending
	if v10712 != 0 {
		goto L32
	} else {
		goto L2302
	}
L2302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+160)) = v10521 + int32(_a_F_StartupXLOG_279)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v10521+int32(160))
	mBase = m.M
	v10720 = m.ExcPending
	if v10720 != 0 {
		goto L32
	} else {
		goto L2303
	}
L2303:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(393), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10725 = m.ExcPending
	if v10725 != 0 {
		goto L32
	} else {
		goto L2304
	}
L2304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2305:
	;
	v10758 = int32(_a_F_StartupXLOG_289)
	goto L2262
L2306:
	;
	goto L2307
L2307:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10734 = m.ExcPending
	if v10734 != 0 {
		goto L32
	} else {
		goto L2308
	}
L2308:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10736 = m.ExcPending
	if v10736 != 0 {
		goto L32
	} else {
		goto L2309
	}
L2309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+128)) = v10521 + int32(_a_F_StartupXLOG_279)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v10521+int32(128))
	mBase = m.M
	v10744 = m.ExcPending
	if v10744 != 0 {
		goto L32
	} else {
		goto L2310
	}
L2310:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(349), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10749 = m.ExcPending
	if v10749 != 0 {
		goto L32
	} else {
		goto L2311
	}
L2311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2312:
	;
	v10801 = F_strlen(m, v10794)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10806 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10806))) = int32(167772223)
	v10809 = F_write(m, v10535, v10794, v10801)
	mBase = m.M
	if v10809 == v10801 {
		goto L2314
	} else {
		goto L2315
	}
L2313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10931 = m.ExcPending
	if v10931 != 0 {
		goto L32
	} else {
		goto L2351
	}
L2314:
	;
	v10811 = int32(_a_F_StartupXLOG_142)
	v10812 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	v10813 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10812))) = v10813
	v10816 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10816))) = int32(167772222)
	v10821 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[12])))
	if v10821 != int32(1) {
		v10835 = v10813
		goto L2319
	} else {
		goto L2320
	}
L2315:
	;
	goto L2316
L2316:
	;
	v10902 = int32(_a_F_StartupXLOG_2)
	v10903 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10905 = v10521 + int32(_a_F_StartupXLOG_277)
	v10906 = F_unlink(m, v10905)
	mBase = m.M
	if v10903 != 0 {
		goto L2344
	} else {
		goto L2345
	}
L2317:
	;
	v10864 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, uint32(v10864))) = int32(0)
	v10867 = F_CloseTransientFile(m, v10535)
	mBase = m.M
	v10868 = m.ExcPending
	if v10868 != 0 {
		goto L32
	} else {
		goto L2335
	}
L2318:
	;
	if v10835 == int32(0) {
		goto L2317
	} else {
		goto L2325
	}
L2319:
	;
	goto L2318
L2320:
	;
	goto L2321
L2321:
	;
	v10826 = F_fsync(m, v10535)
	mBase = m.M
	if v10826 != int32(-1) {
		v10835 = v10826
		goto L2319
	} else {
		goto L2323
	}
L2322:
	;
	v10835 = int32(-1)
	goto L2319
L2323:
	;
	v10830 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10830 == int32(27) {
		goto L2321
	} else {
		goto L2324
	}
L2324:
	;
	goto L2322
L2325:
	;
	v10841 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[131])))
	if v10841 != 0 {
		goto L2327
	} else {
		goto L2328
	}
L2326:
	;
	v10844 = F_errstart(m, v10842, int32(0))
	mBase = m.M
	v10845 = m.ExcPending
	if v10845 != 0 {
		goto L32
	} else {
		goto L2330
	}
L2327:
	;
	v10842 = int32(21)
	goto L2329
L2328:
	;
	v10842 = int32(24)
	goto L2329
L2329:
	;
	goto L2326
L2330:
	;
	if v10844 == int32(0) {
		goto L2317
	} else {
		goto L2331
	}
L2331:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10849 = m.ExcPending
	if v10849 != 0 {
		goto L32
	} else {
		goto L2332
	}
L2332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+64)) = v10521 + int32(_a_F_StartupXLOG_277)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v10521-int32(-64))
	mBase = m.M
	v10857 = m.ExcPending
	if v10857 != 0 {
		goto L32
	} else {
		goto L2333
	}
L2333:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(433), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10862 = m.ExcPending
	if v10862 != 0 {
		goto L32
	} else {
		goto L2334
	}
L2334:
	;
	goto L2317
L2335:
	;
	if v10867 != 0 {
		goto L2313
	} else {
		goto L2336
	}
L2336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+32)) = v10145
	v10871 = v10521 + int32(_a_F_StartupXLOG_279)
	v10876 = F_pg_snprintf(m, v10871, int32(1024), int32(_a_F_StartupXLOG_282), v10521+int32(32))
	mBase = m.M
	v10877 = m.ExcPending
	if v10877 != 0 {
		goto L32
	} else {
		goto L2337
	}
L2337:
	;
	v10881 = F_durable_rename(m, v10521+int32(_a_F_StartupXLOG_277), v10871, int32(21))
	mBase = m.M
	v10882 = m.ExcPending
	if v10882 != 0 {
		goto L32
	} else {
		goto L2338
	}
L2338:
	;
	v10884 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[132]))
	if int32(0) < v10884 {
		goto L2339
	} else {
		goto L2340
	}
L2339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+16)) = v10145
	v10889 = v10521 + int32(_a_F_StartupXLOG_280)
	v10894 = F_pg_snprintf(m, v10889, int32(64), int32(_a_F_StartupXLOG_281), v10521+int32(16))
	mBase = m.M
	v10895 = m.ExcPending
	if v10895 != 0 {
		goto L32
	} else {
		goto L2342
	}
L2340:
	;
	goto L2341
L2341:
	;
	m.G0 = v10521 + int32(_a_F_StartupXLOG_276)
	goto L2260
L2342:
	;
	F_XLogArchiveNotify(m, v10889)
	mBase = m.M
	v10897 = m.ExcPending
	if v10897 != 0 {
		goto L32
	} else {
		goto L2343
	}
L2343:
	;
	goto L2341
L2344:
	;
	v10909 = v10903
	goto L2346
L2345:
	;
	v10909 = int32(51)
	goto L2346
L2346:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10909
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10914 = m.ExcPending
	if v10914 != 0 {
		goto L32
	} else {
		goto L2347
	}
L2347:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10916 = m.ExcPending
	if v10916 != 0 {
		goto L32
	} else {
		goto L2348
	}
L2348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+80)) = v10905
	F_errmsg(m, int32(_a_F_StartupXLOG_288), v10521+int32(80))
	mBase = m.M
	v10922 = m.ExcPending
	if v10922 != 0 {
		goto L32
	} else {
		goto L2349
	}
L2349:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(425), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10927 = m.ExcPending
	if v10927 != 0 {
		goto L32
	} else {
		goto L2350
	}
L2350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2351:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10933 = m.ExcPending
	if v10933 != 0 {
		goto L32
	} else {
		goto L2352
	}
L2352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+48)) = v10521 + int32(_a_F_StartupXLOG_277)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v10521+int32(48))
	mBase = m.M
	v10941 = m.ExcPending
	if v10941 != 0 {
		goto L32
	} else {
		goto L2353
	}
L2353:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(439), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v10946 = m.ExcPending
	if v10946 != 0 {
		goto L32
	} else {
		goto L2354
	}
L2354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2355:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10988 = m.ExcPending
	if v10988 != 0 {
		goto L32
	} else {
		goto L2356
	}
L2356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10521)+144)) = v10521 + int32(_a_F_StartupXLOG_279)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v10521+int32(144))
	mBase = m.M
	v10996 = m.ExcPending
	if v10996 != 0 {
		goto L32
	} else {
		goto L2357
	}
L2357:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_286), int32(363), int32(_a_F_StartupXLOG_287))
	mBase = m.M
	v11001 = m.ExcPending
	if v11001 != 0 {
		goto L32
	} else {
		goto L2358
	}
L2358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2359:
	;
	if v11004 == int32(0) {
		v11028 = v10145
		goto L2174
	} else {
		goto L2360
	}
L2360:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_290), int32(0))
	mBase = m.M
	v11011 = m.ExcPending
	if v11011 != 0 {
		goto L32
	} else {
		goto L2361
	}
L2361:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_291), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v11016 = m.ExcPending
	if v11016 != 0 {
		goto L32
	} else {
		goto L2362
	}
L2362:
	;
	v11028 = v10145
	goto L2174
L2363:
	;
	F_s_lock(m, v11054+int32(440), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v11062 = m.ExcPending
	if v11062 != 0 {
		goto L32
	} else {
		goto L2366
	}
L2364:
	;
	goto L2365
L2365:
	;
	v11064 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	*(*int32)(unsafe.Add(mBase, uint32(v11064)+300)) = v11028
	v11066 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11064)+304)) = v11066
	v11068 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11064)+440)), uint32(v11068))
	if v10054 == int64(0) {
		goto L2367
	} else {
		goto L2368
	}
L2366:
	;
	goto L2365
L2367:
	;
	v11073 = v9921
	goto L2369
L2368:
	;
	v11073 = v10054
	goto L2369
L2369:
	;
	v11074 = *(*int64)(unsafe.Add(mBase, uint32(v9550)))
	v11077 = base.I32_wrap_i64(v11074) & int32(_a_F_StartupXLOG_272)
	v11079 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v11080 = base.I64_extend_i32_s(v11079)
	v11081 = base.I64_div_u_s(v11074, v11080)
	v11084 = base.I64_extend_i32_s(v11079 - int32(1))
	v11085 = v11074 & v11084
	if v11085&int64(35184372080640) == int64(0) {
		goto L2371
	} else {
		goto L2372
	}
L2370:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11064)+16)) = v11123
	v11125 = base.I64_div_u_s(v11073, v11080)
	v11128 = base.I32_wrap_i64(v11073) & int32(_a_F_StartupXLOG_272)
	v11129 = v11073 & v11084
	if v11129&int64(35184372080640) == int64(0) {
		goto L2377
	} else {
		goto L2378
	}
L2371:
	;
	v11091 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[133]))
	v11093 = v11081 * base.I64_extend_i32_s(v11091)
	if v11077 == int32(0) {
		v11121 = v11091
		v11123 = v11093
		goto L2370
	} else {
		goto L2374
	}
L2372:
	;
	goto L2373
L2373:
	;
	v11101 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[133]))
	v11114 = v11081*base.I64_extend_i32_s(v11101) + (int64(base.Ui64(v11085)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v11077 == int32(0) {
		v11121 = v11101
		v11123 = v11114
		goto L2370
	} else {
		goto L2375
	}
L2374:
	;
	v11121 = v11091
	v11123 = v11093 + base.I64_extend_i32_u(v11077-int32(40))
	goto L2370
L2375:
	;
	v11121 = v11101
	v11123 = v11114 + base.I64_extend_i32_u(v11077-int32(24))
	goto L2370
L2376:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11064)+8)) = v11162
	if v11073&int64(8191) != int64(0) {
		goto L2382
	} else {
		goto L2383
	}
L2377:
	;
	v11135 = v11125 * base.I64_extend_i32_s(v11121)
	if v11128 == int32(0) {
		v11162 = v11135
		goto L2376
	} else {
		goto L2380
	}
L2378:
	;
	goto L2379
L2379:
	;
	v11154 = v11125*base.I64_extend_i32_s(v11121) + (int64(base.Ui64(v11129)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v11128 == int32(0) {
		v11162 = v11154
		goto L2376
	} else {
		goto L2381
	}
L2380:
	;
	v11162 = v11135 + base.I64_extend_i32_u(v11128-int32(40))
	goto L2376
L2381:
	;
	v11162 = v11154 + base.I64_extend_i32_u(v11128-int32(24))
	goto L2376
L2382:
	;
	v11168 = *(*int32)(unsafe.Add(mBase, uint32(v11064)+288))
	v11171 = *(*int32)(unsafe.Add(mBase, uint32(v11064)+296))
	v11175 = base.I64_rem_u_s(int64(base.Ui64(v11073)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v11171+int32(1)))
	v11176 = base.I32_wrap_i64(v11175)
	v11179 = v11168 + v11176<<(uint(int32(13))%32)
	v11180 = *(*int64)(unsafe.Add(mBase, uint32(v9550)+32))
	v11182 = base.I32_wrap_i64(v11073 - v11180)
	if v11182 != 0 {
		goto L2385
	} else {
		goto L2386
	}
L2383:
	;
	v11206 = v11064
	v11210 = v11073
	goto L2384
L2384:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11206)+280)) = v11210
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[134])) = v11073
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[135])) = v11073
	v11217 = base.AtomicRmwXchg64(m, v11206, int32(256), v11073)
	v11218 = int32(_a_F_StartupXLOG_292)
	v11219 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v11221 = base.AtomicRmwXchg64(m, v11219, int32(264), v11073)
	v11223 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v11225 = base.AtomicRmwXchg64(m, v11223, int32(272), v11073)
	v11227 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	*(*int64)(unsafe.Add(mBase, uint32(v11227)+192)) = v11073
	*(*int64)(unsafe.Add(mBase, uint32(v11227)+184)) = v11073
	v11230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11227)+312)))
	if v11230 != int32(1) {
		goto L2391
	} else {
		goto L2392
	}
L2385:
	;
	v11183 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+40))
	base.MemoryCopy(m, v11179, v11183, v11182)
	goto L2387
L2386:
	;
	goto L2387
L2387:
	;
	v11186 = int32(_a_F_StartupXLOG_57) - v11182
	if v11186 != 0 {
		goto L2388
	} else {
		goto L2389
	}
L2388:
	;
	base.MemoryFill(m, v11182+v11179, int32(0), v11186)
	goto L2390
L2389:
	;
	goto L2390
L2390:
	;
	v11190 = int32(_a_F_StartupXLOG_292)
	v11191 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v11192 = *(*int32)(unsafe.Add(mBase, uint32(v11191)+292))
	v11196 = *(*int64)(unsafe.Add(mBase, uint32(v9550)+32))
	v11197 = int64(-8192)
	v11200 = base.AtomicRmwXchg64(m, v11192+v11176<<(uint(int32(3))%32), int32(0), v11196-v11197)
	v11202 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v11203 = *(*int64)(unsafe.Add(mBase, uint32(v9550)+32))
	v11206 = v11202
	v11210 = v11203 - v11197
	goto L2384
L2391:
	;
	v11272 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[43])) = uint8(v11272)
	v11274 = F_time(m)
	mBase = m.M
	v11276 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	*(*int64)(unsafe.Add(mBase, uint32(v11276)+248)) = v11073
	*(*int64)(unsafe.Add(mBase, uint32(v11276)+240)) = v11274
	v11280 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11284 = F_LWLockAcquire(m, v11280+int32(512), v11272)
	mBase = m.M
	v11285 = m.ExcPending
	if v11285 != 0 {
		goto L32
	} else {
		goto L2399
	}
L2392:
	;
	v11234 = v11073 - int64(1)
	v11236 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if base.Ui64(v11234&base.I64_extend_i32_s(v11236-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v11236), float64(0.75))))) {
		goto L2391
	} else {
		goto L2393
	}
L2393:
	;
	v11248 = base.I64_div_u_s(v11234, base.I64_extend_i32_s(v11236))
	v11255 = F_XLogFileInitInternal(m, v11248+int64(1), v11028, v41+int32(_a_F_StartupXLOG_4), v41+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v11256 = m.ExcPending
	if v11256 != 0 {
		goto L32
	} else {
		goto L2394
	}
L2394:
	;
	if int32(0) <= v11255 {
		goto L2395
	} else {
		goto L2396
	}
L2395:
	;
	v11259 = F_close(m, v11255)
	mBase = m.M
	goto L2397
L2396:
	;
	goto L2397
L2397:
	;
	v11260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[99]))))
	if v11260 != int32(1) {
		goto L2391
	} else {
		goto L2398
	}
L2398:
	;
	v11263 = int32(_a_F_StartupXLOG_293)
	v11265 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[136]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[136])) = v11265 + int32(1)
	goto L2391
L2399:
	;
	v11287 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	v11288 = *(*int64)(unsafe.Add(mBase, uint32(v11287)+8))
	v11290 = v11288 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v11287)+48)) = v11290
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v11290)))|base.B2i32(base.Ui64(v11290) < base.Ui64(int64(3))) == int32(0) {
		goto L2400
	} else {
		goto L2401
	}
L2400:
	;
	v11328 = v11290
	goto L2403
L2401:
	;
	goto L2402
L2402:
	;
	v11379 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v11379+int32(512))
	mBase = m.M
	v11383 = m.ExcPending
	if v11383 != 0 {
		goto L32
	} else {
		goto L2406
	}
L2403:
	;
	v11337 = v11328 - int64(1)
	if base.Ui32(base.I32_wrap_i64(v11337)) < base.Ui32(int32(3)) {
		v11328 = v11337
		goto L2403
	} else {
		goto L2405
	}
L2404:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11287)+48)) = v11337
	goto L2402
L2405:
	;
	goto L2404
L2406:
	;
	v11385 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[92]))
	if v11385 == int32(0) {
		goto L2407
	} else {
		goto L2408
	}
L2407:
	;
	F_StartupSUBTRANS(m, v10116)
	mBase = m.M
	v11389 = m.ExcPending
	if v11389 != 0 {
		goto L32
	} else {
		goto L2410
	}
L2408:
	;
	goto L2409
L2409:
	;
	v11390 = m.G0
	v11392 = v11390 - int32(16)
	m.G0 = v11392
	v11395 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	v11396 = *(*int64)(unsafe.Add(mBase, uint32(v11395)+8))
	v11397 = base.I32_wrap_i64(v11396)
	*(*int32)(unsafe.Add(mBase, uint32(v11392)+12)) = v11397
	v11400 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11401 = *(*int32)(unsafe.Add(mBase, uint32(v11400)+28))
	v11405 = int64(base.Ui64(v11396)>>(uint(int64(15))%64)) & int64(131071)
	v11408 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[137])))
	v11409 = base.I32_rem_u_s(base.I32_wrap_i64(v11405), v11408)
	v11412 = v11401 + v11409<<(uint(int32(7))%32)
	v11414 = F_LWLockAcquire(m, v11412, int32(0))
	mBase = m.M
	v11415 = m.ExcPending
	if v11415 != 0 {
		goto L32
	} else {
		goto L2411
	}
L2410:
	;
	goto L2409
L2411:
	;
	v11417 = v11397 & int32(_a_F_StartupXLOG_294)
	if v11417 != 0 {
		goto L2412
	} else {
		goto L2413
	}
L2412:
	;
	v11422 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_295), v11405, int32(0), v11392+int32(12))
	mBase = m.M
	v11423 = m.ExcPending
	if v11423 != 0 {
		goto L32
	} else {
		goto L2415
	}
L2413:
	;
	goto L2414
L2414:
	;
	F_LWLockRelease(m, v11412)
	mBase = m.M
	v11497 = m.ExcPending
	if v11497 != 0 {
		goto L32
	} else {
		goto L2425
	}
L2415:
	;
	v11425 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11426 = *(*int32)(unsafe.Add(mBase, uint32(v11425)+4))
	v11427 = int32(2)
	v11430 = *(*int32)(unsafe.Add(mBase, uint32(v11426+v11422<<(uint(v11427)%32))))
	v11432 = int32(base.Ui32(v11417) >> (uint(v11427) % 32))
	v11433 = v11430 + v11432
	v11434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11433))))
	v11435 = int32(-1)
	v11436 = int32(1)
	v11443 = v11434 & (v11435<<(uint(v11397<<(uint(v11436)%32)&int32(6))%32) ^ v11435)
	*(*uint8)(unsafe.Add(mBase, uint32(v11433))) = uint8(v11443)
	v11446 = v11433 + v11436
	v11447 = int32(3)
	v11450 = v11432 ^ int32(_a_F_StartupXLOG_272)
	if v11446&v11447|v11450&v11447|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v11450)) == int32(0) {
		goto L2417
	} else {
		goto L2418
	}
L2416:
	;
	v11485 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11486 = *(*int32)(unsafe.Add(mBase, uint32(v11485)+12))
	v11488 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11486+v11422))) = uint8(v11488)
	goto L2414
L2417:
	;
	if v11432 == int32(_a_F_StartupXLOG_272) {
		goto L2416
	} else {
		goto L2420
	}
L2418:
	;
	v11476 = v11450
	goto L2419
L2419:
	;
	if v11476 == int32(0) {
		goto L2416
	} else {
		goto L2424
	}
L2420:
	;
	v11464 = v11450 + v11430 + v11432 + int32(1)
	v11466 = v11433 + int32(5)
	if base.Ui32(v11466) < base.Ui32(v11464) {
		goto L2421
	} else {
		goto L2422
	}
L2421:
	;
	v11468 = v11464
	goto L2423
L2422:
	;
	v11468 = v11466
	goto L2423
L2423:
	;
	v11476 = (v11468-v11433-int32(2))&int32(-4) + int32(4)
	goto L2419
L2424:
	;
	base.MemoryFill(m, v11446, int32(0), v11476)
	goto L2416
L2425:
	;
	m.G0 = v11392 + int32(16)
	v11501 = m.G0
	v11503 = v11501 - int32(32)
	m.G0 = v11503
	v11506 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11510 = F_LWLockAcquire(m, v11506+int32(1664), int32(1))
	mBase = m.M
	v11511 = m.ExcPending
	if v11511 != 0 {
		goto L32
	} else {
		goto L2426
	}
L2426:
	;
	v11513 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11515 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11516 = *(*int32)(unsafe.Add(mBase, uint32(v11515)))
	*(*int32)(unsafe.Add(mBase, uint32(v11503)+28)) = v11516
	v11518 = *(*int32)(unsafe.Add(mBase, uint32(v11515)+24))
	v11519 = *(*int32)(unsafe.Add(mBase, uint32(v11515)+20))
	v11520 = *(*int64)(unsafe.Add(mBase, uint32(v11515)+8))
	F_LWLockRelease(m, v11513+int32(1664))
	mBase = m.M
	v11524 = m.ExcPending
	if v11524 != 0 {
		goto L32
	} else {
		goto L2427
	}
L2427:
	;
	v11525 = int32(_a_F_StartupXLOG_296)
	v11526 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	v11528 = int32(base.Ui32(v11516) >> (uint(int32(10)) % 32))
	v11529 = base.I64_extend_i32_u(v11528)
	v11531 = base.AtomicRmwXchg64(m, v11526, int32(48), v11529)
	v11532 = *(*int32)(unsafe.Add(mBase, uint32(v11503)+28))
	v11534 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	v11535 = *(*int32)(unsafe.Add(mBase, uint32(v11534)+28))
	v11537 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[138])))
	v11538 = base.I32_rem_u_s(v11528, v11537)
	v11541 = v11535 + v11538<<(uint(int32(7))%32)
	v11543 = F_LWLockAcquire(m, v11541, int32(0))
	mBase = m.M
	v11544 = m.ExcPending
	if v11544 != 0 {
		goto L32
	} else {
		goto L2428
	}
L2428:
	;
	v11546 = v11532 & int32(1023)
	if v11532 != int32(1) {
		goto L2430
	} else {
		goto L2431
	}
L2429:
	;
	v11564 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	v11565 = *(*int32)(unsafe.Add(mBase, uint32(v11564)+4))
	v11569 = *(*int32)(unsafe.Add(mBase, uint32(v11565+v11562<<(uint(int32(2))%32))))
	v11571 = v11546 << (uint(int32(3)) % 32)
	v11572 = v11569 + v11571
	*(*int64)(unsafe.Add(mBase, uint32(v11572))) = v11520
	if base.Ui32(int32(1021)) < base.Ui32(v11546-int32(1)) {
		goto L2438
	} else {
		goto L2439
	}
L2430:
	;
	v11550 = v11546
	goto L2432
L2431:
	;
	v11550 = int32(0)
	goto L2432
L2432:
	;
	if v11550 == int32(0) {
		goto L2433
	} else {
		goto L2434
	}
L2433:
	;
	v11554 = F_SimpleLruZeroPage(m, int32(_a_F_StartupXLOG_297), v11529)
	mBase = m.M
	v11555 = m.ExcPending
	if v11555 != 0 {
		goto L32
	} else {
		goto L2436
	}
L2434:
	;
	goto L2435
L2435:
	;
	v11560 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_297), v11529, int32(1), v11503+int32(28))
	mBase = m.M
	v11561 = m.ExcPending
	if v11561 != 0 {
		goto L32
	} else {
		goto L2437
	}
L2436:
	;
	v11562 = v11554
	goto L2429
L2437:
	;
	v11562 = v11560
	goto L2429
L2438:
	;
	v11613 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	v11614 = *(*int32)(unsafe.Add(mBase, uint32(v11613)+12))
	v11616 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11614+v11562))) = uint8(v11616)
	F_LWLockRelease(m, v11541)
	mBase = m.M
	v11619 = m.ExcPending
	if v11619 != 0 {
		goto L32
	} else {
		goto L2448
	}
L2439:
	;
	v11579 = v11572 + int32(8)
	if v11579&int32(3)|base.B2i32(base.Ui32(v11546) < base.Ui32(int32(895))) == int32(0) {
		goto L2441
	} else {
		goto L2442
	}
L2440:
	;
	if v11604 == int32(0) {
		goto L2438
	} else {
		goto L2447
	}
L2441:
	;
	v11588 = v11572 + int32(12)
	v11590 = v11569 - int32(-8192)
	if base.Ui32(v11590) < base.Ui32(v11588) {
		goto L2444
	} else {
		goto L2445
	}
L2442:
	;
	goto L2443
L2443:
	;
	v11604 = int32(_a_F_StartupXLOG_298) - v11571
	goto L2440
L2444:
	;
	v11592 = v11588
	goto L2446
L2445:
	;
	v11592 = v11590
	goto L2446
L2446:
	;
	v11604 = (v11592-v11572-int32(9))&int32(-4) + int32(4)
	goto L2440
L2447:
	;
	base.MemoryFill(m, v11579, int32(0), v11604)
	goto L2438
L2448:
	;
	v11621 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v11623 = base.I64_div_u_s(v11520, int64(1636))
	v11625 = base.AtomicRmwXchg64(m, v11621, int32(48), v11623)
	v11627 = int64(base.Ui64(v11520) >> (uint(int64(2)) % 64))
	v11629 = base.I64_rem_u_s(v11627, int64(409))
	if v11629 != int64(0) {
		goto L2449
	} else {
		goto L2450
	}
L2449:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11503)+8)) = int64(0)
	v11635 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[139])))
	*(*int64)(unsafe.Add(mBase, uint32(v11503)+16)) = v11520
	v11638 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v11639 = *(*int32)(unsafe.Add(mBase, uint32(v11638)+28))
	v11640 = base.I64_rem_u_s(v11623, v11635)
	v11644 = v11639 + base.I32_wrap_i64(v11640)<<(uint(int32(7))%32)
	v11646 = F_LWLockAcquire(m, v11644, int32(0))
	mBase = m.M
	v11647 = m.ExcPending
	if v11647 != 0 {
		goto L32
	} else {
		goto L2452
	}
L2450:
	;
	goto L2451
L2451:
	;
	v11714 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11718 = F_LWLockAcquire(m, v11714+int32(1664), int32(0))
	mBase = m.M
	v11719 = m.ExcPending
	if v11719 != 0 {
		goto L32
	} else {
		goto L2462
	}
L2452:
	;
	v11652 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_299), v11623, int32(1), v11503+int32(8))
	mBase = m.M
	v11653 = m.ExcPending
	if v11653 != 0 {
		goto L32
	} else {
		goto L2453
	}
L2453:
	;
	v11655 = int32(2)
	v11658 = base.I32_wrap_i64(v11520) << (uint(v11655) % 32) & int32(12)
	v11664 = v11658 + base.I32_wrap_i64(v11629)*int32(20) + int32(4)
	v11666 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v11667 = *(*int32)(unsafe.Add(mBase, uint32(v11666)+4))
	v11671 = *(*int32)(unsafe.Add(mBase, uint32(v11667+v11652<<(uint(v11655)%32))))
	v11672 = v11664 + v11671
	if v11672&int32(3)|base.B2i32(base.Ui32(v11664) < base.Ui32(int32(_a_F_StartupXLOG_300))) == int32(0) {
		goto L2455
	} else {
		goto L2456
	}
L2454:
	;
	if v11696 != 0 {
		goto L2458
	} else {
		goto L2459
	}
L2455:
	;
	v11696 = (base.I32_wrap_i64(v11623)*int32(_a_F_StartupXLOG_301)-v11658+base.I32_wrap_i64(v11627)*int32(-20)+int32(_a_F_StartupXLOG_302))&int32(-4) + int32(4)
	goto L2454
L2456:
	;
	goto L2457
L2457:
	;
	v11696 = int32(_a_F_StartupXLOG_57) - v11664
	goto L2454
L2458:
	;
	base.MemoryFill(m, v11672, int32(0), v11696)
	goto L2460
L2459:
	;
	goto L2460
L2460:
	;
	v11700 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v11701 = *(*int32)(unsafe.Add(mBase, uint32(v11700)+12))
	v11703 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11701+v11652))) = uint8(v11703)
	F_LWLockRelease(m, v11644)
	mBase = m.M
	v11706 = m.ExcPending
	if v11706 != 0 {
		goto L32
	} else {
		goto L2461
	}
L2461:
	;
	goto L2451
L2462:
	;
	v11721 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11722 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11721)+16)) = uint8(v11722)
	v11725 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v11725+int32(1664))
	mBase = m.M
	v11729 = m.ExcPending
	if v11729 != 0 {
		goto L32
	} else {
		goto L2463
	}
L2463:
	;
	F_SetMultiXactIdLimit(m, v11519, v11518)
	mBase = m.M
	v11731 = m.ExcPending
	if v11731 != 0 {
		goto L32
	} else {
		goto L2464
	}
L2464:
	;
	m.G0 = v11503 + int32(32)
	v11735 = int32(0)
	v11736 = m.G0
	v11738 = v11736 - int32(16)
	m.G0 = v11738
	v11741 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11745 = F_LWLockAcquire(m, v11741+int32(2304), v11735)
	mBase = m.M
	v11746 = m.ExcPending
	if v11746 != 0 {
		goto L32
	} else {
		goto L2465
	}
L2465:
	;
	v11748 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[140]))
	v11749 = *(*int32)(unsafe.Add(mBase, uint32(v11748)+4))
	if int32(0) < v11749 {
		goto L2466
	} else {
		goto L2467
	}
L2466:
	;
	v11764 = v11748
	v11776 = v11735
	goto L2469
L2467:
	;
	goto L2468
L2468:
	;
	v12079 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12079+int32(2304))
	mBase = m.M
	v12083 = m.ExcPending
	if v12083 != 0 {
		goto L32
	} else {
		goto L2510
	}
L2469:
	;
	v11791 = *(*int32)(unsafe.Add(mBase, uint32(v11764+v11776<<(uint(int32(2))%32))+8))
	v11792 = *(*int64)(unsafe.Add(mBase, uint32(v11791)+32))
	v11793 = *(*int64)(unsafe.Add(mBase, uint32(v11791)+16))
	v11794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11791)+49)))
	v11797 = F_ProcessTwoPhaseBuffer(m, v11792, v11793, v11794, int32(1), int32(0))
	mBase = m.M
	v11798 = m.ExcPending
	if v11798 != 0 {
		goto L32
	} else {
		goto L2471
	}
L2470:
	;
	goto L2468
L2471:
	;
	if v11797 != 0 {
		goto L2472
	} else {
		goto L2473
	}
L2472:
	;
	v11801 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11802 = m.ExcPending
	if v11802 != 0 {
		goto L32
	} else {
		goto L2475
	}
L2473:
	;
	goto L2474
L2474:
	;
	v12037 = v11776 + int32(1)
	v12039 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[140]))
	v12040 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+4))
	if v12037 < v12040 {
		v11764 = v12039
		v11776 = v12037
		goto L2469
	} else {
		goto L2509
	}
L2475:
	;
	if v11801 != 0 {
		goto L2476
	} else {
		goto L2477
	}
L2476:
	;
	v11803 = *(*int64)(unsafe.Add(mBase, uint32(v11791)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v11738))) = uint32(v11803)
	v11806 = int64(base.Ui64(v11803) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11738)+4)) = uint32(v11806)
	F_errmsg(m, int32(_a_F_StartupXLOG_303), v11738)
	mBase = m.M
	v11810 = m.ExcPending
	if v11810 != 0 {
		goto L32
	} else {
		goto L2479
	}
L2477:
	;
	goto L2478
L2478:
	;
	v11817 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+48))
	v11818 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+44))
	v11819 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+40))
	v11820 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+36))
	v11821 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+32))
	v11822 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+28))
	v11823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11797)+54)))
	v11824 = *(*int64)(unsafe.Add(mBase, uint32(v11791)+32))
	v11826 = v11797 + int32(72)
	v11827 = *(*int64)(unsafe.Add(mBase, uint32(v11797)+16))
	v11828 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+24))
	v11829 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+12))
	F_MarkAsPreparingGuts(m, v11791, v11824, v11826, v11827, v11828, v11829)
	mBase = m.M
	v11831 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11791)+50)) = uint8(v11831)
	v11833 = int32(7)
	v11837 = v11826 + (v11823+v11833)&int32(_a_F_StartupXLOG_304)
	v11842 = int32(-8)
	v11845 = int32(12)
	v11859 = int32(4)
	v11869 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[141]))
	v11870 = *(*int32)(unsafe.Add(mBase, uint32(v11869)))
	v11871 = *(*int32)(unsafe.Add(mBase, uint32(v11791)+4))
	v11874 = v11870 + v11871*int32(768)
	v11875 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+28))
	if int32(65) <= v11875 {
		goto L2483
	} else {
		goto L2484
	}
L2479:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_305), int32(2122), int32(_a_F_StartupXLOG_306))
	mBase = m.M
	v11815 = m.ExcPending
	if v11815 != 0 {
		goto L32
	} else {
		goto L2480
	}
L2480:
	;
	goto L2478
L2481:
	;
	v11897 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11791)+48)) = uint8(v11897)
	v11899 = *(*int32)(unsafe.Add(mBase, uint32(v11895)))
	F_ProcArrayAdd(m, v11899+v11896*int32(768))
	mBase = m.M
	v11904 = m.ExcPending
	if v11904 != 0 {
		goto L32
	} else {
		goto L2490
	}
L2482:
	;
	v11885 = v11883 << (uint(int32(2)) % 32)
	if v11885 != 0 {
		goto L2487
	} else {
		goto L2488
	}
L2483:
	;
	v11878 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11874)+57)) = uint8(v11878)
	v11883 = int32(64)
	goto L2482
L2484:
	;
	goto L2485
L2485:
	;
	if v11875 <= int32(0) {
		v11895 = v11869
		v11896 = v11871
		goto L2481
	} else {
		goto L2486
	}
L2486:
	;
	v11883 = v11875
	goto L2482
L2487:
	;
	base.MemoryCopy(m, v11874+int32(60), v11837, v11885)
	goto L2489
L2488:
	;
	goto L2489
L2489:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11874)+56)) = uint8(v11883)
	v11890 = *(*int32)(unsafe.Add(mBase, uint32(v11791)+4))
	v11892 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[141]))
	v11895 = v11892
	v11896 = v11890
	goto L2481
L2490:
	;
	v11906 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v11906+int32(2304))
	mBase = m.M
	v11910 = m.ExcPending
	if v11910 != 0 {
		goto L32
	} else {
		goto L2491
	}
L2491:
	;
	v11923 = v11837 + (v11822<<(uint(int32(2))%32)+v11833)&v11842 + (v11821*v11845+v11833)&v11842 + (v11820*v11845+v11833)&v11842 + v11819<<(uint(v11859)%32) + v11818<<(uint(v11859)%32) + v11817<<(uint(v11859)%32)
	goto L2492
L2492:
	;
	v11947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11923)+4)))
	if v11947 != 0 {
		goto L2494
	} else {
		goto L2495
	}
L2493:
	;
	v11964 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[92]))
	if base.Ui32(int32(2)) <= base.Ui32(v11964) {
		goto L2501
	} else {
		goto L2502
	}
L2494:
	;
	v11949 = v11923 + int32(8)
	v11952 = *(*int32)(unsafe.Add(mBase, uint32(v11947<<(uint(int32(2))%32))+uint32(_c_F_StartupXLOG[142])))
	if v11952 != 0 {
		goto L2497
	} else {
		goto L2498
	}
L2495:
	;
	goto L2496
L2496:
	;
	goto L2493
L2497:
	;
	v11953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11923)+6)))
	v11954 = *(*int32)(unsafe.Add(mBase, uint32(v11923)))
	m.T0[v11952].(func(*base.Module, int64, int32, int32, int32))(m, v11792, v11953, v11949, v11954)
	mBase = m.M
	v11956 = m.ExcPending
	if v11956 != 0 {
		goto L32
	} else {
		goto L2500
	}
L2498:
	;
	goto L2499
L2499:
	;
	v11957 = *(*int32)(unsafe.Add(mBase, uint32(v11923)))
	v11923 = v11949 + (v11957+int32(7))&int32(-8)
	goto L2492
L2500:
	;
	goto L2499
L2501:
	;
	v11967 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+8))
	v11968 = *(*int32)(unsafe.Add(mBase, uint32(v11797)+28))
	F_StandbyReleaseLockTree(m, v11967, v11968, v11837)
	mBase = m.M
	v11970 = m.ExcPending
	if v11970 != 0 {
		goto L32
	} else {
		goto L2504
	}
L2502:
	;
	goto L2503
L2503:
	;
	v11972 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11976 = F_LWLockAcquire(m, v11972+int32(2304), int32(0))
	mBase = m.M
	v11977 = m.ExcPending
	if v11977 != 0 {
		goto L32
	} else {
		goto L2505
	}
L2504:
	;
	goto L2503
L2505:
	;
	v11979 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[143]))
	*(*int32)(unsafe.Add(mBase, uint32(v11979)+44)) = int32(-1)
	v11983 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v11983+int32(2304))
	mBase = m.M
	v11987 = m.ExcPending
	if v11987 != 0 {
		goto L32
	} else {
		goto L2506
	}
L2506:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[143])) = int32(0)
	F_pfree(m, v11797)
	mBase = m.M
	v11992 = m.ExcPending
	if v11992 != 0 {
		goto L32
	} else {
		goto L2507
	}
L2507:
	;
	v11994 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v11998 = F_LWLockAcquire(m, v11994+int32(2304), int32(0))
	mBase = m.M
	v11999 = m.ExcPending
	if v11999 != 0 {
		goto L32
	} else {
		goto L2508
	}
L2508:
	;
	goto L2474
L2509:
	;
	goto L2470
L2510:
	;
	m.G0 = v11738 + int32(16)
	v12087 = m.G0
	v12089 = v12087 - int32(1024)
	m.G0 = v12089
	v12092 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v12097 = *(*int32)(unsafe.Add(mBase, uint32(v12092)))
	v12098 = *(*int32)(unsafe.Add(mBase, uint32(v12097)+124))
	if v12098 != 0 {
		goto L2512
	} else {
		goto L2513
	}
L2511:
	;
	v12120 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129]))
	if int32(0) <= v12120 {
		goto L2515
	} else {
		goto L2516
	}
L2512:
	;
	v12099 = *(*int64)(unsafe.Add(mBase, uint32(v12098)+16))
	v12100 = *(*int32)(unsafe.Add(mBase, uint32(v12097)+120))
	v12101 = *(*int64)(unsafe.Add(mBase, uint32(v12100)+16))
	v12104 = base.I32_wrap_i64(v12099 - v12101)
	goto L2514
L2513:
	;
	v12104 = int32(0)
	goto L2514
L2514:
	;
	v12105 = *(*int32)(unsafe.Add(mBase, uint32(v12092)+112))
	v12106 = *(*int32)(unsafe.Add(mBase, uint32(v12105)+16))
	v12108 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	v12109 = *(*int32)(unsafe.Add(mBase, uint32(v12105)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12108)+64)) = v12109
	*(*int32)(unsafe.Add(mBase, uint32(v12108)+56)) = v12104
	*(*int32)(unsafe.Add(mBase, uint32(v12108)+60)) = v12109 + v12106
	v12114 = *(*int32)(unsafe.Add(mBase, uint32(v12092)))
	v12115 = *(*int64)(unsafe.Add(mBase, uint32(v12114)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v12092)+16)) = v12115 - int64(-8192)
	goto L2511
L2515:
	;
	v12123 = F_close(m, v12120)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129])) = int32(-1)
	goto L2517
L2516:
	;
	goto L2517
L2517:
	;
	v12128 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v12129 = *(*int32)(unsafe.Add(mBase, uint32(v12128)+24))
	F_pfree(m, v12129)
	mBase = m.M
	v12131 = m.ExcPending
	if v12131 != 0 {
		goto L32
	} else {
		goto L2518
	}
L2518:
	;
	v12133 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	F_XLogReaderFree(m, v12133)
	mBase = m.M
	v12135 = m.ExcPending
	if v12135 != 0 {
		goto L32
	} else {
		goto L2519
	}
L2519:
	;
	v12137 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v12138 = *(*int32)(unsafe.Add(mBase, uint32(v12137)+112))
	F_pfree(m, v12138)
	mBase = m.M
	v12140 = m.ExcPending
	if v12140 != 0 {
		goto L32
	} else {
		goto L2520
	}
L2520:
	;
	v12141 = *(*int32)(unsafe.Add(mBase, uint32(v12137)+24))
	F_hash_destroy(m, v12141)
	mBase = m.M
	v12143 = m.ExcPending
	if v12143 != 0 {
		goto L32
	} else {
		goto L2521
	}
L2521:
	;
	F_pfree(m, v12137)
	mBase = m.M
	v12145 = m.ExcPending
	if v12145 != 0 {
		goto L32
	} else {
		goto L2522
	}
L2522:
	;
	v12147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v12147 != int32(1) {
		goto L2523
	} else {
		goto L2524
	}
L2523:
	;
	m.G0 = v12089 + int32(1024)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144])) = int32(1)
	if v10051 != int64(0) {
		goto L2529
	} else {
		goto L2530
	}
L2524:
	;
	v12153 = F_pg_snprintf(m, v12089, int32(1024), int32(_a_F_StartupXLOG_307), int32(0))
	mBase = m.M
	v12154 = m.ExcPending
	if v12154 != 0 {
		goto L32
	} else {
		goto L2525
	}
L2525:
	;
	v12155 = F_unlink(m, v12089)
	mBase = m.M
	v12159 = F_pg_snprintf(m, v12089, int32(1024), int32(_a_F_StartupXLOG_308), int32(0))
	mBase = m.M
	v12160 = m.ExcPending
	if v12160 != 0 {
		goto L32
	} else {
		goto L2526
	}
L2526:
	;
	v12161 = F_unlink(m, v12089)
	mBase = m.M
	v12163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v12163 != int32(1) {
		goto L2523
	} else {
		goto L2527
	}
L2527:
	;
	v12167 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v12167+int32(4))+12)) = int32(0)
	goto L2528
L2528:
	;
	goto L2523
L2529:
	;
	v12181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[145])))
	if v12181 != int32(1) {
		goto L5
	} else {
		goto L2532
	}
L2530:
	;
	goto L2531
L2531:
	;
	v12342 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[76])))
	*(*uint8)(unsafe.Add(mBase, uint32(v11064)+160)) = uint8(v12342)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v12345 = m.ExcPending
	if v12345 != 0 {
		goto L32
	} else {
		goto L2562
	}
L2532:
	;
	v12186 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v12187 = *(*int32)(unsafe.Add(mBase, uint32(v12186)+308))
	v12189 = base.B2i32(v12187 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[145])) = uint8(v12189)
	if v12189 == int32(0) {
		goto L5
	} else {
		goto L2533
	}
L2533:
	;
	if v10054&int64(8191) != int64(0) {
		goto L4
	} else {
		goto L2534
	}
L2534:
	;
	v12200 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v10054&base.I64_extend_i32_s(v12200-int32(1)) == int64(0) {
		goto L2535
	} else {
		goto L2536
	}
L2535:
	;
	v12207 = int64(40)
	goto L2537
L2536:
	;
	v12207 = int64(24)
	goto L2537
L2537:
	;
	v12210 = base.AtomicRmwXchg32(m, v12186, int32(0), int32(1))
	if v12210 != 0 {
		goto L2538
	} else {
		goto L2539
	}
L2538:
	;
	F_s_lock(m, v12186, int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v12213 = m.ExcPending
	if v12213 != 0 {
		goto L32
	} else {
		goto L2541
	}
L2539:
	;
	goto L2540
L2540:
	;
	v12214 = v12207 | v10054
	v12215 = *(*int64)(unsafe.Add(mBase, uint32(v12186)+8))
	v12216 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12186))), uint32(v12216))
	v12220 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[133])))
	v12221 = base.I64_div_u_s(v12215, v12220)
	v12223 = v12215 - v12221*v12220
	if base.Ui64(v12223) <= base.Ui64(int64(8151)) {
		goto L2543
	} else {
		goto L2544
	}
L2541:
	;
	goto L2540
L2542:
	;
	v12243 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v12247 = v12221*v12243 + v12241&int64(4294967295)
	if v12247 != v12214 {
		goto L3
	} else {
		goto L2546
	}
L2543:
	;
	v12241 = v12223 + int64(40)
	goto L2542
L2544:
	;
	goto L2545
L2545:
	;
	v12229 = v12223 - int64(8152)
	v12230 = int64(8168)
	v12231 = base.I64_div_u_s(v12229, v12230)
	v12241 = v12229 - v12231*v12230 + v12231<<(uint(int64(13))%64) + int64(8216)
	goto L2542
L2546:
	;
	v12249 = int32(_a_F_StartupXLOG_141)
	v12251 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v12251 + int32(1)
	v12256 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[146]))
	if v12256 == int32(-1) {
		goto L2547
	} else {
		goto L2548
	}
L2547:
	;
	v12261 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[90]))
	v12263 = base.I32_rem_s(v12261, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[146])) = v12263
	v12265 = v12263
	goto L2549
L2548:
	;
	v12265 = v12256
	goto L2549
L2549:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[147])) = v12265
	v12269 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	v12274 = F_LWLockAcquire(m, v12269+v12265<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v12275 = m.ExcPending
	if v12275 != 0 {
		goto L32
	} else {
		goto L2550
	}
L2550:
	;
	if v12274 == int32(0) {
		goto L2551
	} else {
		goto L2552
	}
L2551:
	;
	v12278 = int32(_a_F_StartupXLOG_309)
	v12280 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[146]))
	v12284 = base.I32_rem_s(v12280+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[146])) = v12284
	goto L2553
L2552:
	;
	goto L2553
L2553:
	;
	v12286 = F_GetXLogBuffer(m, v10054, v11028)
	mBase = m.M
	v12287 = m.ExcPending
	if v12287 != 0 {
		goto L32
	} else {
		goto L2554
	}
L2554:
	;
	v12288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12286)+2)))
	v12290 = v12288 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v12286)+2)) = uint16(v12290)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v12293 = m.ExcPending
	if v12293 != 0 {
		goto L32
	} else {
		goto L2555
	}
L2555:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v12295 = m.ExcPending
	if v12295 != 0 {
		goto L32
	} else {
		goto L2556
	}
L2556:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v10051
	v12300 = m.G0
	v12301 = int32(16)
	v12302 = v12300 - v12301
	m.G0 = v12302
	F_gettimeofday(m, v12302)
	mBase = m.M
	v12305 = *(*int64)(unsafe.Add(mBase, uint32(v12302)))
	v12306 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12302)+8)))
	m.G0 = v12302 + v12301
	goto L2557
L2557:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[149]))) = v12306 + v12305*int64(1000000) - int64(946684800000000)
	F_XLogRegisterData(m, v41+int32(_a_F_StartupXLOG_1), int32(16))
	mBase = m.M
	v12320 = m.ExcPending
	if v12320 != 0 {
		goto L32
	} else {
		goto L2558
	}
L2558:
	;
	v12323 = F_XLogInsert(m, int32(0), int32(208))
	mBase = m.M
	v12324 = m.ExcPending
	if v12324 != 0 {
		goto L32
	} else {
		goto L2559
	}
L2559:
	;
	v12326 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	if v12326 != v12214 {
		goto L2
	} else {
		goto L2560
	}
L2560:
	;
	F_XLogFlush(m, v12323)
	mBase = m.M
	v12329 = m.ExcPending
	if v12329 != 0 {
		goto L32
	} else {
		goto L2561
	}
L2561:
	;
	v12330 = int32(_a_F_StartupXLOG_141)
	v12332 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v12332 - int32(1)
	goto L2531
L2562:
	;
	v12346 = int32(0)
	if v6283 == v12346 {
		v12459 = v12346
		goto L2563
	} else {
		goto L2564
	}
L2563:
	;
	v12463 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	v12465 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12466 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+180))
	if v12463 != v12466 {
		goto L2584
	} else {
		goto L2585
	}
L2564:
	;
	v12350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v12350 != int32(1) {
		goto L2565
	} else {
		goto L2566
	}
L2565:
	;
	F_RequestCheckpoint(m, int32(38))
	mBase = m.M
	v12458 = m.ExcPending
	if v12458 != 0 {
		goto L32
	} else {
		goto L2582
	}
L2566:
	;
	v12354 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[18])))
	if v12354&int32(1) == int32(0) {
		goto L2565
	} else {
		goto L2567
	}
L2567:
	;
	v12359 = F_PromoteIsTriggered(m)
	mBase = m.M
	v12360 = m.ExcPending
	if v12360 != 0 {
		goto L32
	} else {
		goto L2568
	}
L2568:
	;
	if v12359 == int32(0) {
		goto L2565
	} else {
		goto L2569
	}
L2569:
	;
	v12364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[145])))
	if v12364 != int32(1) {
		goto L1
	} else {
		goto L2570
	}
L2570:
	;
	v12369 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v12370 = *(*int32)(unsafe.Add(mBase, uint32(v12369)+308))
	v12372 = base.B2i32(v12370 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[145])) = uint8(v12372)
	if v12372 == int32(0) {
		goto L1
	} else {
		goto L2571
	}
L2571:
	;
	v12379 = m.G0
	v12380 = int32(16)
	v12381 = v12379 - v12380
	m.G0 = v12381
	F_gettimeofday(m, v12381)
	mBase = m.M
	v12384 = *(*int64)(unsafe.Add(mBase, uint32(v12381)))
	v12385 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12381)+8)))
	m.G0 = v12381 + v12380
	goto L2572
L2572:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v12385 + v12384*int64(1000000) - int64(946684800000000)
	v12396 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[94]))) = v12396
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v12399 = m.ExcPending
	if v12399 != 0 {
		goto L32
	} else {
		goto L2573
	}
L2573:
	;
	v12401 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v12402 = *(*int32)(unsafe.Add(mBase, uint32(v12401)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[149]))) = v12402
	v12404 = *(*int32)(unsafe.Add(mBase, uint32(v12401)+304))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[95]))) = v12404
	F_WALInsertLockRelease(m)
	mBase = m.M
	v12407 = m.ExcPending
	if v12407 != 0 {
		goto L32
	} else {
		goto L2574
	}
L2574:
	;
	v12408 = int32(1)
	v12409 = int32(_a_F_StartupXLOG_141)
	v12411 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v12411 + v12408
	F_XLogBeginInsert(m)
	mBase = m.M
	v12416 = m.ExcPending
	if v12416 != 0 {
		goto L32
	} else {
		goto L2575
	}
L2575:
	;
	F_XLogRegisterData(m, v41+int32(_a_F_StartupXLOG_1), int32(24))
	mBase = m.M
	v12421 = m.ExcPending
	if v12421 != 0 {
		goto L32
	} else {
		goto L2576
	}
L2576:
	;
	v12424 = F_XLogInsert(m, int32(0), int32(144))
	mBase = m.M
	v12425 = m.ExcPending
	if v12425 != 0 {
		goto L32
	} else {
		goto L2577
	}
L2577:
	;
	F_XLogFlush(m, v12424)
	mBase = m.M
	v12427 = m.ExcPending
	if v12427 != 0 {
		goto L32
	} else {
		goto L2578
	}
L2578:
	;
	v12429 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v12433 = F_LWLockAcquire(m, v12429+int32(1152), int32(0))
	mBase = m.M
	v12434 = m.ExcPending
	if v12434 != 0 {
		goto L32
	} else {
		goto L2579
	}
L2579:
	;
	v12436 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v12436)+144)) = v12424
	v12438 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[149])))
	*(*int32)(unsafe.Add(mBase, uint32(v12436)+152)) = v12438
	v12441 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	F_update_controlfile(m, v12441, v12436)
	mBase = m.M
	v12443 = m.ExcPending
	if v12443 != 0 {
		goto L32
	} else {
		goto L2580
	}
L2580:
	;
	v12445 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12445+int32(1152))
	mBase = m.M
	v12449 = m.ExcPending
	if v12449 != 0 {
		goto L32
	} else {
		goto L2581
	}
L2581:
	;
	v12450 = int32(_a_F_StartupXLOG_141)
	v12452 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61])) = v12452 - int32(1)
	v12459 = v12408
	goto L2563
L2582:
	;
	v12459 = v12346
	goto L2563
L2583:
	;
	v12582 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v12582 != int32(1) {
		goto L2603
	} else {
		goto L2604
	}
L2584:
	;
	v12497 = int32(0)
	if base.B2i32(v12466 == v12463)&base.B2i32(v12463 <= v12497) == v12497 {
		goto L2593
	} else {
		goto L2594
	}
L2585:
	;
	v12469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[151])))
	v12470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12465)+184)))
	if v12469 != v12470 {
		goto L2584
	} else {
		goto L2586
	}
L2586:
	;
	v12473 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	v12474 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+188))
	if v12473 != v12474 {
		goto L2584
	} else {
		goto L2587
	}
L2587:
	;
	v12477 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	v12478 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+192))
	if v12477 != v12478 {
		goto L2584
	} else {
		goto L2588
	}
L2588:
	;
	v12481 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	v12482 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+196))
	if v12481 != v12482 {
		goto L2584
	} else {
		goto L2589
	}
L2589:
	;
	v12485 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[154]))
	v12486 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+200))
	if v12485 != v12486 {
		goto L2584
	} else {
		goto L2590
	}
L2590:
	;
	v12489 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[155]))
	v12490 = *(*int32)(unsafe.Add(mBase, uint32(v12465)+204))
	if v12489 != v12490 {
		goto L2584
	} else {
		goto L2591
	}
L2591:
	;
	v12493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[156])))
	v12494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12465)+208)))
	if v12493 == v12494 {
		goto L2583
	} else {
		goto L2592
	}
L2592:
	;
	goto L2584
L2593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[157]))) = v12463
	v12504 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[6]))) = v12504
	v12507 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[96]))) = v12507
	v12510 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[149]))) = v12510
	v12513 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[154]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[95]))) = v12513
	v12516 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[155]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[94]))) = v12516
	v12519 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[151])))
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[98]))) = uint8(v12519)
	v12522 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[156])))
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_StartupXLOG[158]))) = uint8(v12522)
	F_XLogBeginInsert(m)
	mBase = m.M
	v12525 = m.ExcPending
	if v12525 != 0 {
		goto L32
	} else {
		goto L2596
	}
L2594:
	;
	goto L2595
L2595:
	;
	v12538 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v12542 = F_LWLockAcquire(m, v12538+int32(1152), int32(0))
	mBase = m.M
	v12543 = m.ExcPending
	if v12543 != 0 {
		goto L32
	} else {
		goto L2600
	}
L2596:
	;
	F_XLogRegisterData(m, v41+int32(_a_F_StartupXLOG_1), int32(28))
	mBase = m.M
	v12530 = m.ExcPending
	if v12530 != 0 {
		goto L32
	} else {
		goto L2597
	}
L2597:
	;
	v12533 = F_XLogInsert(m, int32(0), int32(96))
	mBase = m.M
	v12534 = m.ExcPending
	if v12534 != 0 {
		goto L32
	} else {
		goto L2598
	}
L2598:
	;
	F_XLogFlush(m, v12533)
	mBase = m.M
	v12536 = m.ExcPending
	if v12536 != 0 {
		goto L32
	} else {
		goto L2599
	}
L2599:
	;
	goto L2595
L2600:
	;
	v12545 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12547 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+188)) = v12547
	v12550 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+192)) = v12550
	v12553 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+196)) = v12553
	v12556 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[154]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+200)) = v12556
	v12559 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[155]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+204)) = v12559
	v12562 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	*(*int32)(unsafe.Add(mBase, uint32(v12545)+180)) = v12562
	v12565 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[151])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12545)+184)) = uint8(v12565)
	v12568 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[156])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12545)+208)) = uint8(v12568)
	v12571 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	F_update_controlfile(m, v12571, v12545)
	mBase = m.M
	v12573 = m.ExcPending
	if v12573 != 0 {
		goto L32
	} else {
		goto L2601
	}
L2601:
	;
	v12575 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12575+int32(1152))
	mBase = m.M
	v12579 = m.ExcPending
	if v12579 != 0 {
		goto L32
	} else {
		goto L2602
	}
L2602:
	;
	goto L2583
L2603:
	;
	v12745 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[156])))
	if v12745 == int32(0) {
		goto L2634
	} else {
		goto L2635
	}
L2604:
	;
	v12586 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[159]))
	if v12586 == int32(0) {
		goto L2605
	} else {
		goto L2606
	}
L2605:
	;
	F_RemoveNonParentXlogFiles(m, v11073, v11028)
	mBase = m.M
	v12598 = m.ExcPending
	if v12598 != 0 {
		goto L32
	} else {
		goto L2609
	}
L2606:
	;
	v12589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12586))))
	if v12589 == int32(0) {
		goto L2605
	} else {
		goto L2607
	}
L2607:
	;
	F_ExecuteRecoveryCommand(m, v12586, int32(_a_F_StartupXLOG_310), int32(1), int32(134217774))
	mBase = m.M
	v12596 = m.ExcPending
	if v12596 != 0 {
		goto L32
	} else {
		goto L2608
	}
L2608:
	;
	goto L2605
L2609:
	;
	v12600 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v11073&base.I64_extend_i32_s(v12600-int32(1)) == int64(0) {
		goto L2603
	} else {
		goto L2610
	}
L2610:
	;
	v12608 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[132]))
	if v12608 <= int32(0) {
		goto L2603
	} else {
		goto L2611
	}
L2611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3472)) = v10067
	v12614 = base.I64_extend_i32_s(v12600)
	v12615 = base.I64_div_u_s(v11073-int64(1), v12614)
	v12617 = base.I64_div_u_s(int64(4294967296), v12614)
	v12618 = base.I64_div_u_s(v12615, v12617)
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3476)) = uint32(v12618)
	v12621 = v12615 - v12618*v12617
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3480)) = uint32(v12621)
	v12624 = v41 + int32(_a_F_StartupXLOG_3)
	v12629 = F_pg_snprintf(m, v12624, int32(64), int32(_a_F_StartupXLOG_269), v41+int32(3472))
	mBase = m.M
	v12630 = m.ExcPending
	if v12630 != 0 {
		goto L32
	} else {
		goto L2612
	}
L2612:
	;
	v12631 = m.G0
	v12633 = v12631 - int32(1168)
	m.G0 = v12633
	*(*int32)(unsafe.Add(mBase, uint32(v12633)+32)) = v12624
	*(*int32)(unsafe.Add(mBase, uint32(v12633)+36)) = int32(_a_F_StartupXLOG_311)
	v12639 = v12633 + int32(144)
	v12644 = F_pg_snprintf(m, v12639, int32(1024), int32(_a_F_StartupXLOG_312), v12633+int32(32))
	mBase = m.M
	v12645 = m.ExcPending
	if v12645 != 0 {
		goto L32
	} else {
		goto L2613
	}
L2613:
	;
	v12646 = int32(1)
	v12648 = v12633 + int32(48)
	v12651 = F___fstatat(m, int32(-100), v12639, v12648, int32(0))
	mBase = m.M
	goto L2615
L2614:
	;
	m.G0 = v12633 + int32(1168)
	if v12680 != 0 {
		goto L2603
	} else {
		goto L2622
	}
L2615:
	;
	if v12651 == int32(0) {
		v12680 = v12646
		goto L2614
	} else {
		goto L2616
	}
L2616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12633)+20)) = int32(_a_F_StartupXLOG_313)
	*(*int32)(unsafe.Add(mBase, uint32(v12633)+16)) = v12624
	v12661 = F_pg_snprintf(m, v12639, int32(1024), int32(_a_F_StartupXLOG_312), v12633+int32(16))
	mBase = m.M
	v12662 = m.ExcPending
	if v12662 != 0 {
		goto L32
	} else {
		goto L2617
	}
L2617:
	;
	v12665 = F___fstatat(m, int32(-100), v12639, v12648, int32(0))
	mBase = m.M
	goto L2618
L2618:
	;
	if v12665 == int32(0) {
		v12680 = v12646
		goto L2614
	} else {
		goto L2619
	}
L2619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12633)+4)) = int32(_a_F_StartupXLOG_311)
	*(*int32)(unsafe.Add(mBase, uint32(v12633))) = v12624
	v12673 = F_pg_snprintf(m, v12639, int32(1024), int32(_a_F_StartupXLOG_312), v12633)
	mBase = m.M
	v12674 = m.ExcPending
	if v12674 != 0 {
		goto L32
	} else {
		goto L2620
	}
L2620:
	;
	v12677 = F___fstatat(m, int32(-100), v12639, v12648, int32(0))
	mBase = m.M
	goto L2621
L2621:
	;
	v12680 = base.B2i32(v12677 == int32(0))
	goto L2614
L2622:
	;
	v12685 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[160])))
	if v12685 == int32(1) {
		goto L2623
	} else {
		goto L2624
	}
L2623:
	;
	F_WaitForWalSummarization(m, v11073)
	mBase = m.M
	v12689 = m.ExcPending
	if v12689 != 0 {
		goto L32
	} else {
		goto L2626
	}
L2624:
	;
	goto L2625
L2625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3456)) = v10067
	v12693 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v12694 = base.I64_div_u_s(int64(4294967296), v12693)
	v12695 = base.I64_div_u_s(v12615, v12694)
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3460)) = uint32(v12695)
	v12698 = v12615 - v12695*v12694
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3464)) = uint32(v12698)
	v12701 = v41 + int32(_a_F_StartupXLOG_1)
	v12706 = F_pg_snprintf(m, v12701, int32(1024), int32(_a_F_StartupXLOG_270), v41+int32(3456))
	mBase = m.M
	v12707 = m.ExcPending
	if v12707 != 0 {
		goto L32
	} else {
		goto L2627
	}
L2626:
	;
	goto L2625
L2627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3440)) = v41 + int32(_a_F_StartupXLOG_3)
	v12712 = v41 + int32(_a_F_StartupXLOG_274)
	v12717 = F_pg_snprintf(m, v12712, int32(64), int32(_a_F_StartupXLOG_314), v41+int32(3440))
	mBase = m.M
	v12718 = m.ExcPending
	if v12718 != 0 {
		goto L32
	} else {
		goto L2628
	}
L2628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3424)) = v12701
	v12721 = v41 + int32(_a_F_StartupXLOG_4)
	v12726 = F_pg_snprintf(m, v12721, int32(1024), int32(_a_F_StartupXLOG_314), v41+int32(3424))
	mBase = m.M
	v12727 = m.ExcPending
	if v12727 != 0 {
		goto L32
	} else {
		goto L2629
	}
L2629:
	;
	F_XLogArchiveCleanup(m, v12712)
	mBase = m.M
	v12729 = m.ExcPending
	if v12729 != 0 {
		goto L32
	} else {
		goto L2630
	}
L2630:
	;
	v12731 = F_durable_rename(m, v12701, v12721, int32(21))
	mBase = m.M
	v12732 = m.ExcPending
	if v12732 != 0 {
		goto L32
	} else {
		goto L2631
	}
L2631:
	;
	F_XLogArchiveNotify(m, v12712)
	mBase = m.M
	v12734 = m.ExcPending
	if v12734 != 0 {
		goto L32
	} else {
		goto L2632
	}
L2632:
	;
	goto L2603
L2633:
	;
	v12785 = int32(0)
	v12786 = m.G0
	v12788 = v12786 - int32(16)
	m.G0 = v12788
	v12791 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v12795 = F_LWLockAcquire(m, v12791+int32(_a_F_StartupXLOG_315), v12785)
	mBase = m.M
	v12796 = m.ExcPending
	if v12796 != 0 {
		goto L32
	} else {
		goto L2641
	}
L2634:
	;
	v12749 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v12753 = F_LWLockAcquire(m, v12749+int32(_a_F_StartupXLOG_316), int32(0))
	mBase = m.M
	v12754 = m.ExcPending
	if v12754 != 0 {
		goto L32
	} else {
		goto L2637
	}
L2635:
	;
	goto L2636
L2636:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v12783 = m.ExcPending
	if v12783 != 0 {
		goto L32
	} else {
		goto L2640
	}
L2637:
	;
	v12755 = int32(_a_F_StartupXLOG_317)
	v12756 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[161]))
	v12757 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12756))) = v12757
	*(*uint8)(unsafe.Add(mBase, uint32(v12756)+24)) = uint8(v12757)
	v12762 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[161]))
	*(*uint16)(unsafe.Add(mBase, uint32(v12762)+16)) = uint16(v12757)
	*(*int64)(unsafe.Add(mBase, uint32(v12762)+8)) = int64(-9223372036854775807 - 1)
	v12768 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	*(*int64)(unsafe.Add(mBase, uint32(v12768)+40)) = int64(0)
	v12774 = F_SlruScanDirectory(m, int32(_a_F_StartupXLOG_318), int32(302), v12757)
	mBase = m.M
	v12775 = m.ExcPending
	if v12775 != 0 {
		goto L32
	} else {
		goto L2638
	}
L2638:
	;
	v12777 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12777+int32(_a_F_StartupXLOG_316))
	mBase = m.M
	v12781 = m.ExcPending
	if v12781 != 0 {
		goto L32
	} else {
		goto L2639
	}
L2639:
	;
	goto L2633
L2640:
	;
	goto L2633
L2641:
	;
	v12798 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[58]))
	if v12798 != int32(2) {
		goto L2643
	} else {
		goto L2644
	}
L2642:
	;
	v12808 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v12809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12808)+1)))
	if v12806 != v12809 {
		goto L2649
	} else {
		goto L2650
	}
L2643:
	;
	v12801 = F_CheckLogicalSlotExists(m)
	mBase = m.M
	v12802 = m.ExcPending
	if v12802 != 0 {
		goto L32
	} else {
		goto L2646
	}
L2644:
	;
	goto L2645
L2645:
	;
	v12806 = int32(1)
	goto L2642
L2646:
	;
	if v12801 == int32(0) {
		v12806 = v12785
		goto L2642
	} else {
		goto L2647
	}
L2647:
	;
	goto L2645
L2648:
	;
	v12859 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[18])))
	if v12859 == int32(1) {
		goto L2664
	} else {
		goto L2665
	}
L2649:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12808))) = uint8(v12806)
	v12813 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	*(*uint8)(unsafe.Add(mBase, uint32(v12813)+1)) = uint8(v12806)
	v12817 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v12818 = m.ExcPending
	if v12818 != 0 {
		goto L32
	} else {
		goto L2652
	}
L2650:
	;
	goto L2651
L2651:
	;
	v12853 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12853+int32(_a_F_StartupXLOG_315))
	mBase = m.M
	v12857 = m.ExcPending
	if v12857 != 0 {
		goto L32
	} else {
		goto L2663
	}
L2652:
	;
	if v12817 != 0 {
		goto L2653
	} else {
		goto L2654
	}
L2653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12788))) = v12806
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_319), v12788)
	mBase = m.M
	v12822 = m.ExcPending
	if v12822 != 0 {
		goto L32
	} else {
		goto L2656
	}
L2654:
	;
	goto L2655
L2655:
	;
	v12829 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v12830 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12829)+2)) = uint8(v12830)
	v12833 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12833+int32(_a_F_StartupXLOG_315))
	mBase = m.M
	v12837 = m.ExcPending
	if v12837 != 0 {
		goto L32
	} else {
		goto L2658
	}
L2656:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_320), int32(659), int32(_a_F_StartupXLOG_321))
	mBase = m.M
	v12827 = m.ExcPending
	if v12827 != 0 {
		goto L32
	} else {
		goto L2657
	}
L2657:
	;
	goto L2655
L2658:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12788)+15)) = uint8(v12806)
	F_XLogBeginInsert(m)
	mBase = m.M
	v12840 = m.ExcPending
	if v12840 != 0 {
		goto L32
	} else {
		goto L2659
	}
L2659:
	;
	F_XLogRegisterData(m, v12788+int32(15), int32(1))
	mBase = m.M
	v12845 = m.ExcPending
	if v12845 != 0 {
		goto L32
	} else {
		goto L2660
	}
L2660:
	;
	v12848 = F_XLogInsert(m, int32(0), int32(240))
	mBase = m.M
	v12849 = m.ExcPending
	if v12849 != 0 {
		goto L32
	} else {
		goto L2661
	}
L2661:
	;
	F_XLogFlush(m, v12848)
	mBase = m.M
	v12851 = m.ExcPending
	if v12851 != 0 {
		goto L32
	} else {
		goto L2662
	}
L2662:
	;
	goto L2648
L2663:
	;
	goto L2648
L2664:
	;
	v12863 = F_EmitProcSignalBarrier(m, int32(1))
	mBase = m.M
	v12864 = m.ExcPending
	if v12864 != 0 {
		goto L32
	} else {
		goto L2667
	}
L2665:
	;
	goto L2666
L2666:
	;
	m.G0 = v12788 + int32(16)
	v12870 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+40))
	if v12870 != 0 {
		goto L2669
	} else {
		goto L2670
	}
L2667:
	;
	F_WaitForProcSignalBarrier(m, v12863)
	mBase = m.M
	v12866 = m.ExcPending
	if v12866 != 0 {
		goto L32
	} else {
		goto L2668
	}
L2668:
	;
	goto L2666
L2669:
	;
	F_pfree(m, v12870)
	mBase = m.M
	v12872 = m.ExcPending
	if v12872 != 0 {
		goto L32
	} else {
		goto L2672
	}
L2670:
	;
	goto L2671
L2671:
	;
	v12873 = *(*int32)(unsafe.Add(mBase, uint32(v9550)+64))
	F_pfree(m, v12873)
	mBase = m.M
	v12875 = m.ExcPending
	if v12875 != 0 {
		goto L32
	} else {
		goto L2673
	}
L2672:
	;
	goto L2671
L2673:
	;
	F_pfree(m, v9550)
	mBase = m.M
	v12877 = m.ExcPending
	if v12877 != 0 {
		goto L32
	} else {
		goto L2674
	}
L2674:
	;
	v12879 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	v12883 = F_LWLockAcquire(m, v12879+int32(1152), int32(0))
	mBase = m.M
	v12884 = m.ExcPending
	if v12884 != 0 {
		goto L32
	} else {
		goto L2675
	}
L2675:
	;
	v12886 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v12886)+16)) = int32(6)
	v12890 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	v12893 = base.AtomicRmwXchg32(m, v12890, int32(440), int32(1))
	if v12893 != 0 {
		goto L2676
	} else {
		goto L2677
	}
L2676:
	;
	F_s_lock(m, v12890+int32(440), int32(_a_F_StartupXLOG_203))
	mBase = m.M
	v12898 = m.ExcPending
	if v12898 != 0 {
		goto L32
	} else {
		goto L2679
	}
L2677:
	;
	goto L2678
L2678:
	;
	v12900 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[71]))
	*(*int32)(unsafe.Add(mBase, uint32(v12900)+308)) = int32(2)
	v12903 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12900)+440)), uint32(v12903))
	v12907 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v12909 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v12907, v12909)
	mBase = m.M
	v12911 = m.ExcPending
	if v12911 != 0 {
		goto L32
	} else {
		goto L2680
	}
L2679:
	;
	goto L2678
L2680:
	;
	v12913 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[73]))
	F_LWLockRelease(m, v12913+int32(1152))
	mBase = m.M
	v12917 = m.ExcPending
	if v12917 != 0 {
		goto L32
	} else {
		goto L2681
	}
L2681:
	;
	v12921 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[141]))
	v12922 = *(*int32)(unsafe.Add(mBase, uint32(v12921)+68))
	if v12922 != int32(-1) {
		goto L2683
	} else {
		goto L2684
	}
L2682:
	;
	F_WaitLSNWakeup(m, int32(0), int64(0))
	mBase = m.M
	v12935 = m.ExcPending
	if v12935 != 0 {
		goto L32
	} else {
		goto L2686
	}
L2683:
	;
	v12925 = *(*int32)(unsafe.Add(mBase, uint32(v12921)))
	F_SetLatch(m, v12925+v12922*int32(768)+int32(316))
	mBase = m.M
	goto L2685
L2684:
	;
	goto L2685
L2685:
	;
	goto L2682
L2686:
	;
	F_WaitLSNWakeup(m, int32(1), int64(0))
	mBase = m.M
	v12939 = m.ExcPending
	if v12939 != 0 {
		goto L32
	} else {
		goto L2687
	}
L2687:
	;
	F_WaitLSNWakeup(m, int32(2), int64(0))
	mBase = m.M
	v12943 = m.ExcPending
	if v12943 != 0 {
		goto L32
	} else {
		goto L2688
	}
L2688:
	;
	v12945 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[92]))
	if v12945 != 0 {
		goto L2689
	} else {
		goto L2690
	}
L2689:
	;
	F_ShutdownRecoveryTransactionEnvironment(m)
	mBase = m.M
	v12947 = m.ExcPending
	if v12947 != 0 {
		goto L32
	} else {
		goto L2692
	}
L2690:
	;
	goto L2691
L2691:
	;
	v12948 = int32(1)
	F_WalSndWakeup(m, v12948, v12948)
	mBase = m.M
	v12951 = m.ExcPending
	if v12951 != 0 {
		goto L32
	} else {
		goto L2693
	}
L2692:
	;
	goto L2691
L2693:
	;
	if v12459 != 0 {
		goto L2694
	} else {
		goto L2695
	}
L2694:
	;
	F_RequestCheckpoint(m, int32(8))
	mBase = m.M
	v12954 = m.ExcPending
	if v12954 != 0 {
		goto L32
	} else {
		goto L2697
	}
L2695:
	;
	goto L2696
L2696:
	;
	m.G0 = v37
	return
L2697:
	;
	goto L2696
L2698:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12962 = m.ExcPending
	if v12962 != 0 {
		goto L32
	} else {
		goto L2699
	}
L2699:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_322), int32(0))
	mBase = m.M
	v12966 = m.ExcPending
	if v12966 != 0 {
		goto L32
	} else {
		goto L2700
	}
L2700:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_323), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12971 = m.ExcPending
	if v12971 != 0 {
		goto L32
	} else {
		goto L2701
	}
L2701:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2702:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12978 = m.ExcPending
	if v12978 != 0 {
		goto L32
	} else {
		goto L2703
	}
L2703:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_324), int32(0))
	mBase = m.M
	v12982 = m.ExcPending
	if v12982 != 0 {
		goto L32
	} else {
		goto L2704
	}
L2704:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_325), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12987 = m.ExcPending
	if v12987 != 0 {
		goto L32
	} else {
		goto L2705
	}
L2705:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2706:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12993 = m.ExcPending
	if v12993 != 0 {
		goto L32
	} else {
		goto L2707
	}
L2707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3840)) = int32(_a_F_StartupXLOG_5)
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v41+int32(3840))
	mBase = m.M
	v13000 = m.ExcPending
	if v13000 != 0 {
		goto L32
	} else {
		goto L2708
	}
L2708:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_326), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v13005 = m.ExcPending
	if v13005 != 0 {
		goto L32
	} else {
		goto L2709
	}
L2709:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2710:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13011 = m.ExcPending
	if v13011 != 0 {
		goto L32
	} else {
		goto L2711
	}
L2711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3808)) = v41 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_327), v41+int32(3808))
	mBase = m.M
	v13019 = m.ExcPending
	if v13019 != 0 {
		goto L32
	} else {
		goto L2712
	}
L2712:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_328), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v13024 = m.ExcPending
	if v13024 != 0 {
		goto L32
	} else {
		goto L2713
	}
L2713:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2714:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3760)) = v41 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_327), v41+int32(3760))
	mBase = m.M
	v13036 = m.ExcPending
	if v13036 != 0 {
		goto L32
	} else {
		goto L2715
	}
L2715:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_329), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v13041 = m.ExcPending
	if v13041 != 0 {
		goto L32
	} else {
		goto L2716
	}
L2716:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2717:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v13048 = m.ExcPending
	if v13048 != 0 {
		goto L32
	} else {
		goto L2718
	}
L2718:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_330), int32(0))
	mBase = m.M
	v13052 = m.ExcPending
	if v13052 != 0 {
		goto L32
	} else {
		goto L2719
	}
L2719:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_331), int32(0))
	mBase = m.M
	v13056 = m.ExcPending
	if v13056 != 0 {
		goto L32
	} else {
		goto L2720
	}
L2720:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_332), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v13061 = m.ExcPending
	if v13061 != 0 {
		goto L32
	} else {
		goto L2721
	}
L2721:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2722:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13067 = m.ExcPending
	if v13067 != 0 {
		goto L32
	} else {
		goto L2723
	}
L2723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3392)) = v41 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v41+int32(3392))
	mBase = m.M
	v13075 = m.ExcPending
	if v13075 != 0 {
		goto L32
	} else {
		goto L2724
	}
L2724:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3498), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13080 = m.ExcPending
	if v13080 != 0 {
		goto L32
	} else {
		goto L2725
	}
L2725:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2726:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13086 = m.ExcPending
	if v13086 != 0 {
		goto L32
	} else {
		goto L2727
	}
L2727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3408)) = v41 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_285), v41+int32(3408))
	mBase = m.M
	v13094 = m.ExcPending
	if v13094 != 0 {
		goto L32
	} else {
		goto L2728
	}
L2728:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3512), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13099 = m.ExcPending
	if v13099 != 0 {
		goto L32
	} else {
		goto L2729
	}
L2729:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2730:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3616)) = v41 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_61), v41+int32(3616))
	mBase = m.M
	v13109 = m.ExcPending
	if v13109 != 0 {
		goto L32
	} else {
		goto L2731
	}
L2731:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3544), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13114 = m.ExcPending
	if v13114 != 0 {
		goto L32
	} else {
		goto L2732
	}
L2732:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2733:
	;
	v13122 = v13116
	goto L2735
L2734:
	;
	v13122 = int32(51)
	goto L2735
L2735:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v13122
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13127 = m.ExcPending
	if v13127 != 0 {
		goto L32
	} else {
		goto L2736
	}
L2736:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13129 = m.ExcPending
	if v13129 != 0 {
		goto L32
	} else {
		goto L2737
	}
L2737:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3600)) = v13118
	F_errmsg(m, int32(_a_F_StartupXLOG_288), v41+int32(3600))
	mBase = m.M
	v13135 = m.ExcPending
	if v13135 != 0 {
		goto L32
	} else {
		goto L2738
	}
L2738:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3568), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13140 = m.ExcPending
	if v13140 != 0 {
		goto L32
	} else {
		goto L2739
	}
L2739:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2740:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13146 = m.ExcPending
	if v13146 != 0 {
		goto L32
	} else {
		goto L2741
	}
L2741:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3568)) = v41 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v41+int32(3568))
	mBase = m.M
	v13154 = m.ExcPending
	if v13154 != 0 {
		goto L32
	} else {
		goto L2742
	}
L2742:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3583), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13159 = m.ExcPending
	if v13159 != 0 {
		goto L32
	} else {
		goto L2743
	}
L2743:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2744:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13165 = m.ExcPending
	if v13165 != 0 {
		goto L32
	} else {
		goto L2745
	}
L2745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3552)) = v41 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v41+int32(3552))
	mBase = m.M
	v13173 = m.ExcPending
	if v13173 != 0 {
		goto L32
	} else {
		goto L2746
	}
L2746:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3588), int32(_a_F_StartupXLOG_273))
	mBase = m.M
	v13178 = m.ExcPending
	if v13178 != 0 {
		goto L32
	} else {
		goto L2747
	}
L2747:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2748:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v13180
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13192 = m.ExcPending
	if v13192 != 0 {
		goto L32
	} else {
		goto L2749
	}
L2749:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v13194 = m.ExcPending
	if v13194 != 0 {
		goto L32
	} else {
		goto L2750
	}
L2750:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+3680)) = v13182
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v41+int32(3680))
	mBase = m.M
	v13200 = m.ExcPending
	if v13200 != 0 {
		goto L32
	} else {
		goto L2751
	}
L2751:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_333), int32(_a_F_StartupXLOG_334))
	mBase = m.M
	v13205 = m.ExcPending
	if v13205 != 0 {
		goto L32
	} else {
		goto L2752
	}
L2752:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2753:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_335), int32(0))
	mBase = m.M
	v13215 = m.ExcPending
	if v13215 != 0 {
		goto L32
	} else {
		goto L2754
	}
L2754:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_336), int32(_a_F_StartupXLOG_337))
	mBase = m.M
	v13220 = m.ExcPending
	if v13220 != 0 {
		goto L32
	} else {
		goto L2755
	}
L2755:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2756:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3524)) = uint32(v10054)
	v13227 = int64(base.Ui64(v10054) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3520)) = uint32(v13227)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_338), v41+int32(3520))
	mBase = m.M
	v13233 = m.ExcPending
	if v13233 != 0 {
		goto L32
	} else {
		goto L2757
	}
L2757:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_339), int32(_a_F_StartupXLOG_337))
	mBase = m.M
	v13238 = m.ExcPending
	if v13238 != 0 {
		goto L32
	} else {
		goto L2758
	}
L2758:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2759:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3508)) = uint32(v12247)
	v13245 = int64(base.Ui64(v12247) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3504)) = uint32(v13245)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_340), v41+int32(3504))
	mBase = m.M
	v13251 = m.ExcPending
	if v13251 != 0 {
		goto L32
	} else {
		goto L2760
	}
L2760:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_341), int32(_a_F_StartupXLOG_337))
	mBase = m.M
	v13256 = m.ExcPending
	if v13256 != 0 {
		goto L32
	} else {
		goto L2761
	}
L2761:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2762:
	;
	v13262 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3492)) = uint32(v13262)
	v13265 = int64(base.Ui64(v13262) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v41)+3488)) = uint32(v13265)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_342), v41+int32(3488))
	mBase = m.M
	v13271 = m.ExcPending
	if v13271 != 0 {
		goto L32
	} else {
		goto L2763
	}
L2763:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_343), int32(_a_F_StartupXLOG_337))
	mBase = m.M
	v13276 = m.ExcPending
	if v13276 != 0 {
		goto L32
	} else {
		goto L2764
	}
L2764:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2765:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_344), int32(0))
	mBase = m.M
	v13285 = m.ExcPending
	if v13285 != 0 {
		goto L32
	} else {
		goto L2766
	}
L2766:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_345), int32(_a_F_StartupXLOG_346))
	mBase = m.M
	v13290 = m.ExcPending
	if v13290 != 0 {
		goto L32
	} else {
		goto L2767
	}
L2767:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
