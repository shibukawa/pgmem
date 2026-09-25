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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
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
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
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
	var v667 int32
	_ = v667
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int64
	_ = v918
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int64
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1110 int64
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1133 int64
	_ = v1133
	var v1134 int64
	_ = v1134
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int64
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1297 int64
	_ = v1297
	var v1300 int64
	_ = v1300
	var v1302 int64
	_ = v1302
	var v1303 int64
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1323 int64
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int64
	_ = v1333
	var v1334 int64
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int64
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int64
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int64
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int64
	_ = v1359
	var v1362 int64
	_ = v1362
	var v1368 int32
	_ = v1368
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int64
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1396 int64
	_ = v1396
	var v1399 int64
	_ = v1399
	var v1400 int64
	_ = v1400
	var v1403 int64
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1431 int64
	_ = v1431
	var v1434 int64
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1512 int32
	_ = v1512
	var v1524 int32
	_ = v1524
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1582 int64
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1640 int32
	_ = v1640
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1780 int32
	_ = v1780
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1842 int32
	_ = v1842
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
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2098 int32
	_ = v2098
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2122 int64
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2126 int64
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2144 int64
	_ = v2144
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2153 int64
	_ = v2153
	var v2156 int64
	_ = v2156
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2170 int64
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2175 int64
	_ = v2175
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int64
	_ = v2189
	var v2192 int64
	_ = v2192
	var v2198 int32
	_ = v2198
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int64
	_ = v2210
	var v2211 int64
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2215 int64
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2225 int64
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2228 int64
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int64
	_ = v2232
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2249 int32
	_ = v2249
	var v2251 int64
	_ = v2251
	var v2254 int64
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2271 int64
	_ = v2271
	var v2274 int64
	_ = v2274
	var v2275 int64
	_ = v2275
	var v2278 int64
	_ = v2278
	var v2284 int32
	_ = v2284
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2306 int32
	_ = v2306
	var v2311 int32
	_ = v2311
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2328 int64
	_ = v2328
	var v2329 int64
	_ = v2329
	var v2331 int64
	_ = v2331
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2360 int32
	_ = v2360
	var v2361 int64
	_ = v2361
	var v2362 int64
	_ = v2362
	var v2364 int64
	_ = v2364
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2397 int32
	_ = v2397
	var v2403 int32
	_ = v2403
	var v2408 int64
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2427 int32
	_ = v2427
	var v2432 int64
	_ = v2432
	var v2435 int64
	_ = v2435
	var v2441 int32
	_ = v2441
	var v2448 int32
	_ = v2448
	var v2455 int32
	_ = v2455
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2468 int64
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2476 int64
	_ = v2476
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int64
	_ = v2493
	var v2499 int32
	_ = v2499
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2569 int32
	_ = v2569
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2585 int32
	_ = v2585
	var v2590 int32
	_ = v2590
	var v2595 int64
	_ = v2595
	var v2603 int32
	_ = v2603
	var v2607 int32
	_ = v2607
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2617 int32
	_ = v2617
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2644 int32
	_ = v2644
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2679 int64
	_ = v2679
	var v2683 int64
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2695 int64
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2703 int32
	_ = v2703
	var v2706 int64
	_ = v2706
	var v2714 int32
	_ = v2714
	var v2723 int32
	_ = v2723
	var v2727 int32
	_ = v2727
	var v2731 int32
	_ = v2731
	var v2736 int32
	_ = v2736
	var v2737 int64
	_ = v2737
	var v2741 int64
	_ = v2741
	var v2744 int32
	_ = v2744
	var v2747 int64
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2755 int64
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2761 int64
	_ = v2761
	var v2769 int32
	_ = v2769
	var v2770 int64
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2774 int64
	_ = v2774
	var v2780 int64
	_ = v2780
	var v2796 int32
	_ = v2796
	var v2798 int64
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2811 int32
	_ = v2811
	var v2813 int64
	_ = v2813
	var v2818 int32
	_ = v2818
	var v2821 int64
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2827 int64
	_ = v2827
	var v2833 int32
	_ = v2833
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int64
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2851 int64
	_ = v2851
	var v2857 int32
	_ = v2857
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2870 int32
	_ = v2870
	var v2875 int32
	_ = v2875
	var v2879 int32
	_ = v2879
	var v2883 int32
	_ = v2883
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int64
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int64
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2914 int32
	_ = v2914
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2934 int32
	_ = v2934
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v2997 int64
	_ = v2997
	var v3005 int32
	_ = v3005
	var v3009 int32
	_ = v3009
	var v3013 int32
	_ = v3013
	var v3019 int32
	_ = v3019
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3040 int32
	_ = v3040
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3053 int32
	_ = v3053
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3088 int32
	_ = v3088
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3129 int32
	_ = v3129
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3157 int32
	_ = v3157
	var v3166 int32
	_ = v3166
	var v3194 int32
	_ = v3194
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3207 int32
	_ = v3207
	var v3212 int32
	_ = v3212
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3235 int32
	_ = v3235
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3246 int32
	_ = v3246
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3257 int32
	_ = v3257
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3275 int32
	_ = v3275
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3289 int32
	_ = v3289
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3307 int32
	_ = v3307
	var v3314 int32
	_ = v3314
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3329 int32
	_ = v3329
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3343 int32
	_ = v3343
	var v3348 int32
	_ = v3348
	var v3353 int32
	_ = v3353
	var v3357 int32
	_ = v3357
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3367 int32
	_ = v3367
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3401 int32
	_ = v3401
	var v3406 int32
	_ = v3406
	var v3415 int32
	_ = v3415
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3424 int32
	_ = v3424
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3435 int32
	_ = v3435
	var v3437 int32
	_ = v3437
	var v3440 int32
	_ = v3440
	var v3445 int32
	_ = v3445
	var v3450 int32
	_ = v3450
	var v3460 int32
	_ = v3460
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3491 int32
	_ = v3491
	var v3496 int32
	_ = v3496
	var v3500 int32
	_ = v3500
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3507 int32
	_ = v3507
	var v3511 int32
	_ = v3511
	var v3517 int32
	_ = v3517
	var v3520 int32
	_ = v3520
	var v3526 int32
	_ = v3526
	var v3530 int32
	_ = v3530
	var v3535 int32
	_ = v3535
	var v3539 int32
	_ = v3539
	var v3544 int32
	_ = v3544
	var v3551 int32
	_ = v3551
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3594 int64
	_ = v3594
	var v3596 int64
	_ = v3596
	var v3597 int64
	_ = v3597
	var v3604 int32
	_ = v3604
	var v3608 int32
	_ = v3608
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3618 int64
	_ = v3618
	var v3619 int64
	_ = v3619
	var v3628 int32
	_ = v3628
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3702 int32
	_ = v3702
	var v3704 int32
	_ = v3704
	var v3709 int32
	_ = v3709
	var v3711 int32
	_ = v3711
	var v3718 int32
	_ = v3718
	var v3720 int32
	_ = v3720
	var v3728 int32
	_ = v3728
	var v3733 int32
	_ = v3733
	var v3737 int32
	_ = v3737
	var v3739 int32
	_ = v3739
	var v3747 int32
	_ = v3747
	var v3752 int32
	_ = v3752
	var v3756 int32
	_ = v3756
	var v3758 int32
	_ = v3758
	var v3766 int32
	_ = v3766
	var v3771 int32
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3781 int32
	_ = v3781
	var v3786 int32
	_ = v3786
	var v3790 int32
	_ = v3790
	var v3793 int32
	_ = v3793
	var v3804 int32
	_ = v3804
	var v3809 int32
	_ = v3809
	var v3813 int32
	_ = v3813
	var v3816 int32
	_ = v3816
	var v3825 int32
	_ = v3825
	var v3830 int32
	_ = v3830
	var v3834 int32
	_ = v3834
	var v3837 int32
	_ = v3837
	var v3846 int32
	_ = v3846
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3861 int32
	_ = v3861
	var v3866 int32
	_ = v3866
	var v3870 int32
	_ = v3870
	var v3872 int32
	_ = v3872
	var v3880 int32
	_ = v3880
	var v3885 int32
	_ = v3885
	var v3889 int32
	_ = v3889
	var v3891 int32
	_ = v3891
	var v3900 int32
	_ = v3900
	var v3905 int32
	_ = v3905
	var v3909 int32
	_ = v3909
	var v3912 int32
	_ = v3912
	var v3918 int32
	_ = v3918
	var v3922 int32
	_ = v3922
	var v3927 int32
	_ = v3927
	var v3931 int32
	_ = v3931
	var v3934 int32
	_ = v3934
	var v3940 int32
	_ = v3940
	var v3944 int32
	_ = v3944
	var v3949 int32
	_ = v3949
	var v3987 int32
	_ = v3987
	var v3991 int32
	_ = v3991
	var v3995 int32
	_ = v3995
	var v4000 int32
	_ = v4000
	var v4002 int32
	_ = v4002
	var v4003 int32
	_ = v4003
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4010 int32
	_ = v4010
	var v4041 int32
	_ = v4041
	var v4044 int32
	_ = v4044
	var v4047 int32
	_ = v4047
	var v4050 int32
	_ = v4050
	var v4054 int32
	_ = v4054
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4061 int32
	_ = v4061
	var v4064 int32
	_ = v4064
	var v4065 int32
	_ = v4065
	var v4101 int32
	_ = v4101
	var v4103 int32
	_ = v4103
	var v4105 int32
	_ = v4105
	var v4106 int64
	_ = v4106
	var v4111 int32
	_ = v4111
	var v4112 int64
	_ = v4112
	var v4114 int32
	_ = v4114
	var v4115 int32
	_ = v4115
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4123 int64
	_ = v4123
	var v4125 int32
	_ = v4125
	var v4127 int32
	_ = v4127
	var v4130 int64
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4145 int32
	_ = v4145
	var v4150 int32
	_ = v4150
	var v4151 int32
	_ = v4151
	var v4155 int32
	_ = v4155
	var v4160 int32
	_ = v4160
	var v4163 int32
	_ = v4163
	var v4164 int32
	_ = v4164
	var v4168 int32
	_ = v4168
	var v4174 int32
	_ = v4174
	var v4176 int32
	_ = v4176
	var v4181 int32
	_ = v4181
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4196 int32
	_ = v4196
	var v4201 int32
	_ = v4201
	var v4211 int32
	_ = v4211
	var v4216 int32
	_ = v4216
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4228 int32
	_ = v4228
	var v4232 int32
	_ = v4232
	var v4234 int32
	_ = v4234
	var v4238 int32
	_ = v4238
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4278 int32
	_ = v4278
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4284 int64
	_ = v4284
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4290 int32
	_ = v4290
	var v4292 int64
	_ = v4292
	var v4295 int64
	_ = v4295
	var v4301 int32
	_ = v4301
	var v4306 int32
	_ = v4306
	var v4313 int32
	_ = v4313
	var v4318 int32
	_ = v4318
	var v4350 int32
	_ = v4350
	var v4352 int32
	_ = v4352
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4394 int32
	_ = v4394
	var v4401 int32
	_ = v4401
	var v4406 int32
	_ = v4406
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4419 int32
	_ = v4419
	var v4424 int32
	_ = v4424
	var v4428 int32
	_ = v4428
	var v4430 int32
	_ = v4430
	var v4437 int32
	_ = v4437
	var v4442 int32
	_ = v4442
	var v4446 int32
	_ = v4446
	var v4448 int32
	_ = v4448
	var v4458 int32
	_ = v4458
	var v4463 int32
	_ = v4463
	var v4467 int32
	_ = v4467
	var v4470 int32
	_ = v4470
	var v4474 int32
	_ = v4474
	var v4479 int32
	_ = v4479
	var v4483 int32
	_ = v4483
	var v4486 int32
	_ = v4486
	var v4493 int32
	_ = v4493
	var v4498 int32
	_ = v4498
	var v4502 int32
	_ = v4502
	var v4504 int32
	_ = v4504
	var v4511 int32
	_ = v4511
	var v4516 int32
	_ = v4516
	var v4518 int32
	_ = v4518
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4524 int64
	_ = v4524
	var v4526 int64
	_ = v4526
	var v4528 int64
	_ = v4528
	var v4530 int32
	_ = v4530
	var v4532 int32
	_ = v4532
	var v4534 int32
	_ = v4534
	var v4538 int32
	_ = v4538
	var v4539 int32
	_ = v4539
	var v4541 int32
	_ = v4541
	var v4542 int32
	_ = v4542
	var v4544 int32
	_ = v4544
	var v4545 int32
	_ = v4545
	var v4549 int32
	_ = v4549
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4585 int32
	_ = v4585
	var v4589 int32
	_ = v4589
	var v4591 int32
	_ = v4591
	var v4592 int64
	_ = v4592
	var v4600 int32
	_ = v4600
	var v4604 int32
	_ = v4604
	var v4608 int32
	_ = v4608
	var v4614 int32
	_ = v4614
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4626 int32
	_ = v4626
	var v4627 int32
	_ = v4627
	var v4628 int32
	_ = v4628
	var v4632 int32
	_ = v4632
	var v4635 int32
	_ = v4635
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4648 int32
	_ = v4648
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4658 int32
	_ = v4658
	var v4668 int32
	_ = v4668
	var v4674 int64
	_ = v4674
	var v4678 int32
	_ = v4678
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4684 int64
	_ = v4684
	var v4688 int32
	_ = v4688
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4728 int32
	_ = v4728
	var v4732 int32
	_ = v4732
	var v4734 int32
	_ = v4734
	var v4737 int32
	_ = v4737
	var v4739 int32
	_ = v4739
	var v4743 int32
	_ = v4743
	var v4745 int32
	_ = v4745
	var v4750 int32
	_ = v4750
	var v4751 int32
	_ = v4751
	var v4760 int32
	_ = v4760
	var v4764 int32
	_ = v4764
	var v4765 int32
	_ = v4765
	var v4769 int32
	_ = v4769
	var v4776 int32
	_ = v4776
	var v4780 int32
	_ = v4780
	var v4781 int32
	_ = v4781
	var v4785 int32
	_ = v4785
	var v4790 int32
	_ = v4790
	var v4792 int32
	_ = v4792
	var v4795 int32
	_ = v4795
	var v4799 int32
	_ = v4799
	var v4800 int32
	_ = v4800
	var v4801 int32
	_ = v4801
	var v4804 int64
	_ = v4804
	var v4805 int64
	_ = v4805
	var v4818 int32
	_ = v4818
	var v4860 int32
	_ = v4860
	var v4868 int32
	_ = v4868
	var v4872 int32
	_ = v4872
	var v4873 int32
	_ = v4873
	var v4878 int32
	_ = v4878
	var v4880 int32
	_ = v4880
	var v4884 int32
	_ = v4884
	var v4888 int32
	_ = v4888
	var v4892 int32
	_ = v4892
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4900 int32
	_ = v4900
	var v4901 int32
	_ = v4901
	var v4908 int32
	_ = v4908
	var v4913 int32
	_ = v4913
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4921 int32
	_ = v4921
	var v4926 int32
	_ = v4926
	var v4927 int32
	_ = v4927
	var v4931 int32
	_ = v4931
	var v4936 int32
	_ = v4936
	var v4941 int32
	_ = v4941
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4947 int32
	_ = v4947
	var v4950 int64
	_ = v4950
	var v4951 int64
	_ = v4951
	var v4961 int32
	_ = v4961
	var v5006 int32
	_ = v5006
	var v5014 int32
	_ = v5014
	var v5017 int32
	_ = v5017
	var v5018 int32
	_ = v5018
	var v5023 int32
	_ = v5023
	var v5025 int32
	_ = v5025
	var v5028 int32
	_ = v5028
	var v5032 int32
	_ = v5032
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5043 int32
	_ = v5043
	var v5044 int32
	_ = v5044
	var v5050 int32
	_ = v5050
	var v5055 int32
	_ = v5055
	var v5056 int32
	_ = v5056
	var v5096 int32
	_ = v5096
	var v5097 int32
	_ = v5097
	var v5104 int32
	_ = v5104
	var v5105 int32
	_ = v5105
	var v5110 int32
	_ = v5110
	var v5111 int32
	_ = v5111
	var v5120 int32
	_ = v5120
	var v5122 int32
	_ = v5122
	var v5124 int32
	_ = v5124
	var v5128 int32
	_ = v5128
	var v5136 int32
	_ = v5136
	var v5137 int32
	_ = v5137
	var v5147 int32
	_ = v5147
	var v5154 int32
	_ = v5154
	var v5158 int32
	_ = v5158
	var v5162 int32
	_ = v5162
	var v5167 int32
	_ = v5167
	var v5168 int32
	_ = v5168
	var v5178 int32
	_ = v5178
	var v5180 int32
	_ = v5180
	var v5184 int32
	_ = v5184
	var v5185 int32
	_ = v5185
	var v5192 int32
	_ = v5192
	var v5195 int32
	_ = v5195
	var v5196 int32
	_ = v5196
	var v5198 int32
	_ = v5198
	var v5199 int32
	_ = v5199
	var v5202 int32
	_ = v5202
	var v5203 int32
	_ = v5203
	var v5204 int32
	_ = v5204
	var v5208 int32
	_ = v5208
	var v5209 int32
	_ = v5209
	var v5212 int32
	_ = v5212
	var v5221 int32
	_ = v5221
	var v5224 int32
	_ = v5224
	var v5226 int32
	_ = v5226
	var v5233 int32
	_ = v5233
	var v5234 int32
	_ = v5234
	var v5239 int32
	_ = v5239
	var v5240 int32
	_ = v5240
	var v5249 int32
	_ = v5249
	var v5251 int32
	_ = v5251
	var v5253 int32
	_ = v5253
	var v5257 int32
	_ = v5257
	var v5265 int32
	_ = v5265
	var v5266 int32
	_ = v5266
	var v5271 int64
	_ = v5271
	var v5273 int64
	_ = v5273
	var v5279 int32
	_ = v5279
	var v5284 int32
	_ = v5284
	var v5292 int32
	_ = v5292
	var v5296 int32
	_ = v5296
	var v5297 int32
	_ = v5297
	var v5302 int64
	_ = v5302
	var v5304 int64
	_ = v5304
	var v5310 int32
	_ = v5310
	var v5316 int32
	_ = v5316
	var v5317 int32
	_ = v5317
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5331 int32
	_ = v5331
	var v5337 int32
	_ = v5337
	var v5338 int32
	_ = v5338
	var v5343 int32
	_ = v5343
	var v5344 int32
	_ = v5344
	var v5347 int32
	_ = v5347
	var v5354 int32
	_ = v5354
	var v5356 int32
	_ = v5356
	var v5358 int32
	_ = v5358
	var v5360 int32
	_ = v5360
	var v5370 int32
	_ = v5370
	var v5371 int32
	_ = v5371
	var v5380 int32
	_ = v5380
	var v5383 int32
	_ = v5383
	var v5393 int32
	_ = v5393
	var v5401 int32
	_ = v5401
	var v5405 int32
	_ = v5405
	var v5406 int32
	_ = v5406
	var v5415 int32
	_ = v5415
	var v5417 int32
	_ = v5417
	var v5418 int32
	_ = v5418
	var v5423 int32
	_ = v5423
	var v5424 int32
	_ = v5424
	var v5430 int32
	_ = v5430
	var v5435 int32
	_ = v5435
	var v5438 int32
	_ = v5438
	var v5439 int32
	_ = v5439
	var v5447 int32
	_ = v5447
	var v5452 int32
	_ = v5452
	var v5453 int32
	_ = v5453
	var v5454 int32
	_ = v5454
	var v5459 int32
	_ = v5459
	var v5460 int32
	_ = v5460
	var v5469 int32
	_ = v5469
	var v5475 int32
	_ = v5475
	var v5476 int32
	_ = v5476
	var v5478 int32
	_ = v5478
	var v5484 int32
	_ = v5484
	var v5489 int32
	_ = v5489
	var v5490 int64
	_ = v5490
	var v5492 int32
	_ = v5492
	var v5493 int32
	_ = v5493
	var v5498 int32
	_ = v5498
	var v5499 int32
	_ = v5499
	var v5511 int32
	_ = v5511
	var v5517 int32
	_ = v5517
	var v5522 int32
	_ = v5522
	var v5523 int32
	_ = v5523
	var v5524 int32
	_ = v5524
	var v5528 int32
	_ = v5528
	var v5530 int32
	_ = v5530
	var v5533 int32
	_ = v5533
	var v5534 int32
	_ = v5534
	var v5538 int64
	_ = v5538
	var v5540 int64
	_ = v5540
	var v5546 int32
	_ = v5546
	var v5548 int32
	_ = v5548
	var v5549 int32
	_ = v5549
	var v5550 int32
	_ = v5550
	var v5557 int32
	_ = v5557
	var v5572 int32
	_ = v5572
	var v5574 int32
	_ = v5574
	var v5582 int32
	_ = v5582
	var v5584 int32
	_ = v5584
	var v5586 int32
	_ = v5586
	var v5587 int32
	_ = v5587
	var v5589 int32
	_ = v5589
	var v5590 int32
	_ = v5590
	var v5592 int32
	_ = v5592
	var v5593 int32
	_ = v5593
	var v5594 int32
	_ = v5594
	var v5599 int32
	_ = v5599
	var v5600 int32
	_ = v5600
	var v5606 int32
	_ = v5606
	var v5608 int32
	_ = v5608
	var v5610 int32
	_ = v5610
	var v5613 int32
	_ = v5613
	var v5623 int32
	_ = v5623
	var v5629 int32
	_ = v5629
	var v5630 int32
	_ = v5630
	var v5631 int32
	_ = v5631
	var v5634 int32
	_ = v5634
	var v5635 int32
	_ = v5635
	var v5636 int32
	_ = v5636
	var v5640 int32
	_ = v5640
	var v5641 int32
	_ = v5641
	var v5645 int64
	_ = v5645
	var v5647 int64
	_ = v5647
	var v5653 int32
	_ = v5653
	var v5658 int32
	_ = v5658
	var v5660 int64
	_ = v5660
	var v5662 int64
	_ = v5662
	var v5668 int32
	_ = v5668
	var v5673 int32
	_ = v5673
	var v5674 int32
	_ = v5674
	var v5681 int32
	_ = v5681
	var v5682 int32
	_ = v5682
	var v5689 int32
	_ = v5689
	var v5694 int32
	_ = v5694
	var v5695 int32
	_ = v5695
	var v5705 int32
	_ = v5705
	var v5710 int32
	_ = v5710
	var v5713 int32
	_ = v5713
	var v5716 int32
	_ = v5716
	var v5717 int32
	_ = v5717
	var v5727 int32
	_ = v5727
	var v5732 int32
	_ = v5732
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5777 int32
	_ = v5777
	var v5782 int32
	_ = v5782
	var v5786 int32
	_ = v5786
	var v5787 int32
	_ = v5787
	var v5788 int32
	_ = v5788
	var v5791 int64
	_ = v5791
	var v5792 int64
	_ = v5792
	var v5802 int32
	_ = v5802
	var v5847 int32
	_ = v5847
	var v5855 int32
	_ = v5855
	var v5858 int32
	_ = v5858
	var v5859 int32
	_ = v5859
	var v5864 int32
	_ = v5864
	var v5866 int32
	_ = v5866
	var v5869 int32
	_ = v5869
	var v5873 int32
	_ = v5873
	var v5908 int32
	_ = v5908
	var v5909 int32
	_ = v5909
	var v5912 int32
	_ = v5912
	var v5913 int32
	_ = v5913
	var v5920 int32
	_ = v5920
	var v5925 int32
	_ = v5925
	var v5927 int32
	_ = v5927
	var v6000 int32
	_ = v6000
	var v6001 int32
	_ = v6001
	var v6004 int32
	_ = v6004
	var v6012 int32
	_ = v6012
	var v6017 int32
	_ = v6017
	var v6019 int32
	_ = v6019
	var v6026 int32
	_ = v6026
	var v6028 int32
	_ = v6028
	var v6030 int32
	_ = v6030
	var v6032 int32
	_ = v6032
	var v6036 int32
	_ = v6036
	var v6038 int32
	_ = v6038
	var v6040 int32
	_ = v6040
	var v6041 int32
	_ = v6041
	var v6044 int32
	_ = v6044
	var v6045 int32
	_ = v6045
	var v6049 int32
	_ = v6049
	var v6050 int32
	_ = v6050
	var v6051 int32
	_ = v6051
	var v6054 int32
	_ = v6054
	var v6055 int32
	_ = v6055
	var v6059 int32
	_ = v6059
	var v6060 int32
	_ = v6060
	var v6063 int32
	_ = v6063
	var v6067 int32
	_ = v6067
	var v6068 int64
	_ = v6068
	var v6070 int64
	_ = v6070
	var v6073 int32
	_ = v6073
	var v6076 int32
	_ = v6076
	var v6077 int32
	_ = v6077
	var v6079 int32
	_ = v6079
	var v6082 int32
	_ = v6082
	var v6083 int32
	_ = v6083
	var v6086 int32
	_ = v6086
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6122 int32
	_ = v6122
	var v6125 int32
	_ = v6125
	var v6128 int32
	_ = v6128
	var v6131 int32
	_ = v6131
	var v6138 int32
	_ = v6138
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6155 int32
	_ = v6155
	var v6159 int32
	_ = v6159
	var v6164 int32
	_ = v6164
	var v6168 int32
	_ = v6168
	var v6169 int32
	_ = v6169
	var v6205 int32
	_ = v6205
	var v6210 int32
	_ = v6210
	var v6214 int32
	_ = v6214
	var v6221 int32
	_ = v6221
	var v6222 int32
	_ = v6222
	var v6226 int32
	_ = v6226
	var v6231 int32
	_ = v6231
	var v6232 int32
	_ = v6232
	var v6234 int32
	_ = v6234
	var v6242 int32
	_ = v6242
	var v6244 int32
	_ = v6244
	var v6245 int32
	_ = v6245
	var v6253 int32
	_ = v6253
	var v6254 int32
	_ = v6254
	var v6258 int32
	_ = v6258
	var v6260 int32
	_ = v6260
	var v6262 int32
	_ = v6262
	var v6266 int32
	_ = v6266
	var v6267 int32
	_ = v6267
	var v6269 int32
	_ = v6269
	var v6272 int32
	_ = v6272
	var v6277 int64
	_ = v6277
	var v6280 int32
	_ = v6280
	var v6282 int32
	_ = v6282
	var v6287 int32
	_ = v6287
	var v6294 int32
	_ = v6294
	var v6295 int32
	_ = v6295
	var v6296 int32
	_ = v6296
	var v6298 int32
	_ = v6298
	var v6299 int32
	_ = v6299
	var v6303 int32
	_ = v6303
	var v6335 int32
	_ = v6335
	var v6341 int32
	_ = v6341
	var v6342 int32
	_ = v6342
	var v6346 int32
	_ = v6346
	var v6348 int32
	_ = v6348
	var v6352 int32
	_ = v6352
	var v6357 int32
	_ = v6357
	var v6389 int32
	_ = v6389
	var v6393 int32
	_ = v6393
	var v6398 int32
	_ = v6398
	var v6433 int32
	_ = v6433
	var v6435 int32
	_ = v6435
	var v6438 int32
	_ = v6438
	var v6441 int32
	_ = v6441
	var v6443 int32
	_ = v6443
	var v6450 int32
	_ = v6450
	var v6452 int64
	_ = v6452
	var v6454 int64
	_ = v6454
	var v6457 int32
	_ = v6457
	var v6462 int32
	_ = v6462
	var v6464 int32
	_ = v6464
	var v6465 int64
	_ = v6465
	var v6467 int64
	_ = v6467
	var v6470 int32
	_ = v6470
	var v6471 int64
	_ = v6471
	var v6472 int32
	_ = v6472
	var v6474 int32
	_ = v6474
	var v6475 int64
	_ = v6475
	var v6482 int32
	_ = v6482
	var v6491 int32
	_ = v6491
	var v6492 int32
	_ = v6492
	var v6493 int32
	_ = v6493
	var v6496 int64
	_ = v6496
	var v6497 int64
	_ = v6497
	var v6508 int32
	_ = v6508
	var v6513 int32
	_ = v6513
	var v6517 int32
	_ = v6517
	var v6524 int32
	_ = v6524
	var v6526 int32
	_ = v6526
	var v6528 int32
	_ = v6528
	var v6530 int32
	_ = v6530
	var v6532 int64
	_ = v6532
	var v6534 int64
	_ = v6534
	var v6537 int32
	_ = v6537
	var v6539 int32
	_ = v6539
	var v6541 int32
	_ = v6541
	var v6544 int32
	_ = v6544
	var v6545 int32
	_ = v6545
	var v6546 int32
	_ = v6546
	var v6549 int32
	_ = v6549
	var v6557 int32
	_ = v6557
	var v6559 int32
	_ = v6559
	var v6560 int64
	_ = v6560
	var v6563 int64
	_ = v6563
	var v6569 int32
	_ = v6569
	var v6574 int32
	_ = v6574
	var v6575 int32
	_ = v6575
	var v6579 int32
	_ = v6579
	var v6580 int32
	_ = v6580
	var v6581 int32
	_ = v6581
	var v6584 int32
	_ = v6584
	var v6585 int32
	_ = v6585
	var v6594 int32
	_ = v6594
	var v6599 int32
	_ = v6599
	var v6632 int32
	_ = v6632
	var v6635 int32
	_ = v6635
	var v6638 int32
	_ = v6638
	var v6642 int32
	_ = v6642
	var v6644 int32
	_ = v6644
	var v6647 int32
	_ = v6647
	var v6651 int32
	_ = v6651
	var v6654 int32
	_ = v6654
	var v6659 int32
	_ = v6659
	var v6660 int32
	_ = v6660
	var v6662 int32
	_ = v6662
	var v6663 int64
	_ = v6663
	var v6666 int64
	_ = v6666
	var v6672 int32
	_ = v6672
	var v6677 int32
	_ = v6677
	var v6680 int32
	_ = v6680
	var v6684 int32
	_ = v6684
	var v6685 int32
	_ = v6685
	var v6690 int32
	_ = v6690
	var v6720 int32
	_ = v6720
	var v6728 int32
	_ = v6728
	var v6730 int32
	_ = v6730
	var v6733 int32
	_ = v6733
	var v6734 int64
	_ = v6734
	var v6736 int64
	_ = v6736
	var v6742 int32
	_ = v6742
	var v6744 int32
	_ = v6744
	var v6759 int32
	_ = v6759
	var v6760 int32
	_ = v6760
	var v6764 int32
	_ = v6764
	var v6765 int64
	_ = v6765
	var v6766 int32
	_ = v6766
	var v6768 int32
	_ = v6768
	var v6770 int32
	_ = v6770
	var v6774 int64
	_ = v6774
	var v6780 int32
	_ = v6780
	var v6785 int32
	_ = v6785
	var v6788 int32
	_ = v6788
	var v6790 int32
	_ = v6790
	var v6791 int32
	_ = v6791
	var v6794 int32
	_ = v6794
	var v6796 int32
	_ = v6796
	var v6800 int32
	_ = v6800
	var v6802 int32
	_ = v6802
	var v6806 int32
	_ = v6806
	var v6813 int32
	_ = v6813
	var v6814 int32
	_ = v6814
	var v6818 int32
	_ = v6818
	var v6823 int32
	_ = v6823
	var v6825 int64
	_ = v6825
	var v6831 int32
	_ = v6831
	var v6842 int32
	_ = v6842
	var v6845 int64
	_ = v6845
	var v6847 int64
	_ = v6847
	var v6855 int32
	_ = v6855
	var v6865 int32
	_ = v6865
	var v6866 int32
	_ = v6866
	var v6870 int64
	_ = v6870
	var v6873 int64
	_ = v6873
	var v6879 int32
	_ = v6879
	var v6884 int32
	_ = v6884
	var v6886 int32
	_ = v6886
	var v6887 int32
	_ = v6887
	var v6890 int32
	_ = v6890
	var v6895 int32
	_ = v6895
	var v6897 int32
	_ = v6897
	var v6899 int32
	_ = v6899
	var v6900 int32
	_ = v6900
	var v6906 int64
	_ = v6906
	var v6911 int32
	_ = v6911
	var v6915 int32
	_ = v6915
	var v6917 int32
	_ = v6917
	var v6923 int32
	_ = v6923
	var v6926 int32
	_ = v6926
	var v6928 int32
	_ = v6928
	var v6934 int32
	_ = v6934
	var v6938 int32
	_ = v6938
	var v6940 int32
	_ = v6940
	var v6943 int32
	_ = v6943
	var v6947 int32
	_ = v6947
	var v6952 int32
	_ = v6952
	var v6953 int32
	_ = v6953
	var v6954 int32
	_ = v6954
	var v6957 int32
	_ = v6957
	var v6961 int32
	_ = v6961
	var v6966 int32
	_ = v6966
	var v6967 int32
	_ = v6967
	var v6968 int32
	_ = v6968
	var v6971 int32
	_ = v6971
	var v6975 int32
	_ = v6975
	var v6982 int32
	_ = v6982
	var v6985 int32
	_ = v6985
	var v6993 int32
	_ = v6993
	var v6994 int32
	_ = v6994
	var v6998 int32
	_ = v6998
	var v6999 int32
	_ = v6999
	var v7000 int32
	_ = v7000
	var v7005 int64
	_ = v7005
	var v7006 int64
	_ = v7006
	var v7014 int32
	_ = v7014
	var v7016 int32
	_ = v7016
	var v7017 int32
	_ = v7017
	var v7019 int32
	_ = v7019
	var v7020 int32
	_ = v7020
	var v7026 int64
	_ = v7026
	var v7031 int32
	_ = v7031
	var v7035 int32
	_ = v7035
	var v7037 int32
	_ = v7037
	var v7043 int32
	_ = v7043
	var v7046 int32
	_ = v7046
	var v7048 int32
	_ = v7048
	var v7054 int32
	_ = v7054
	var v7058 int32
	_ = v7058
	var v7060 int32
	_ = v7060
	var v7063 int32
	_ = v7063
	var v7067 int32
	_ = v7067
	var v7072 int32
	_ = v7072
	var v7073 int32
	_ = v7073
	var v7074 int32
	_ = v7074
	var v7077 int32
	_ = v7077
	var v7081 int32
	_ = v7081
	var v7088 int32
	_ = v7088
	var v7091 int32
	_ = v7091
	var v7099 int32
	_ = v7099
	var v7100 int32
	_ = v7100
	var v7104 int32
	_ = v7104
	var v7105 int32
	_ = v7105
	var v7106 int32
	_ = v7106
	var v7111 int64
	_ = v7111
	var v7112 int64
	_ = v7112
	var v7120 int32
	_ = v7120
	var v7121 int32
	_ = v7121
	var v7123 int32
	_ = v7123
	var v7124 int32
	_ = v7124
	var v7125 int32
	_ = v7125
	var v7127 int32
	_ = v7127
	var v7128 int32
	_ = v7128
	var v7131 int32
	_ = v7131
	var v7138 int32
	_ = v7138
	var v7140 int32
	_ = v7140
	var v7141 int32
	_ = v7141
	var v7142 int32
	_ = v7142
	var v7143 int32
	_ = v7143
	var v7153 int32
	_ = v7153
	var v7156 int32
	_ = v7156
	var v7166 int32
	_ = v7166
	var v7167 int64
	_ = v7167
	var v7171 int64
	_ = v7171
	var v7191 int32
	_ = v7191
	var v7195 int32
	_ = v7195
	var v7201 int32
	_ = v7201
	var v7207 int32
	_ = v7207
	var v7208 int32
	_ = v7208
	var v7209 int32
	_ = v7209
	var v7212 int32
	_ = v7212
	var v7214 int32
	_ = v7214
	var v7218 int32
	_ = v7218
	var v7219 int32
	_ = v7219
	var v7222 int32
	_ = v7222
	var v7228 int32
	_ = v7228
	var v7229 int64
	_ = v7229
	var v7233 int32
	_ = v7233
	var v7234 int32
	_ = v7234
	var v7235 int32
	_ = v7235
	var v7238 int64
	_ = v7238
	var v7239 int64
	_ = v7239
	var v7247 int64
	_ = v7247
	var v7251 int64
	_ = v7251
	var v7257 int64
	_ = v7257
	var v7266 int64
	_ = v7266
	var v7269 int32
	_ = v7269
	var v7273 int32
	_ = v7273
	var v7276 int32
	_ = v7276
	var v7281 int32
	_ = v7281
	var v7283 int32
	_ = v7283
	var v7284 int32
	_ = v7284
	var v7285 int32
	_ = v7285
	var v7321 int64
	_ = v7321
	var v7325 int32
	_ = v7325
	var v7326 int32
	_ = v7326
	var v7327 int32
	_ = v7327
	var v7330 int64
	_ = v7330
	var v7331 int64
	_ = v7331
	var v7339 int64
	_ = v7339
	var v7342 int64
	_ = v7342
	var v7348 int64
	_ = v7348
	var v7357 int64
	_ = v7357
	var v7360 int32
	_ = v7360
	var v7365 int32
	_ = v7365
	var v7366 int32
	_ = v7366
	var v7372 int32
	_ = v7372
	var v7377 int32
	_ = v7377
	var v7379 int32
	_ = v7379
	var v7384 int32
	_ = v7384
	var v7385 int32
	_ = v7385
	var v7387 int32
	_ = v7387
	var v7390 int32
	_ = v7390
	var v7395 int32
	_ = v7395
	var v7397 int32
	_ = v7397
	var v7398 int32
	_ = v7398
	var v7399 int32
	_ = v7399
	var v7437 int32
	_ = v7437
	var v7438 int32
	_ = v7438
	var v7443 int32
	_ = v7443
	var v7478 int32
	_ = v7478
	var v7479 int32
	_ = v7479
	var v7480 int32
	_ = v7480
	var v7486 int32
	_ = v7486
	var v7491 int32
	_ = v7491
	var v7493 int32
	_ = v7493
	var v7494 int32
	_ = v7494
	var v7495 int32
	_ = v7495
	var v7497 int32
	_ = v7497
	var v7501 int32
	_ = v7501
	var v7502 int32
	_ = v7502
	var v7503 int32
	_ = v7503
	var v7504 int32
	_ = v7504
	var v7506 int32
	_ = v7506
	var v7509 int64
	_ = v7509
	var v7511 int32
	_ = v7511
	var v7512 int32
	_ = v7512
	var v7518 int32
	_ = v7518
	var v7521 int32
	_ = v7521
	var v7524 int32
	_ = v7524
	var v7525 int32
	_ = v7525
	var v7528 int32
	_ = v7528
	var v7535 int32
	_ = v7535
	var v7536 int32
	_ = v7536
	var v7537 int32
	_ = v7537
	var v7539 int32
	_ = v7539
	var v7545 int32
	_ = v7545
	var v7551 int32
	_ = v7551
	var v7556 int64
	_ = v7556
	var v7559 int32
	_ = v7559
	var v7561 int32
	_ = v7561
	var v7564 int32
	_ = v7564
	var v7567 int32
	_ = v7567
	var v7570 int32
	_ = v7570
	var v7572 int32
	_ = v7572
	var v7579 int32
	_ = v7579
	var v7580 int64
	_ = v7580
	var v7582 int32
	_ = v7582
	var v7585 int32
	_ = v7585
	var v7589 int32
	_ = v7589
	var v7592 int32
	_ = v7592
	var v7596 int32
	_ = v7596
	var v7598 int32
	_ = v7598
	var v7599 int32
	_ = v7599
	var v7600 int32
	_ = v7600
	var v7602 int32
	_ = v7602
	var v7607 int32
	_ = v7607
	var v7608 int64
	_ = v7608
	var v7609 int64
	_ = v7609
	var v7611 int64
	_ = v7611
	var v7613 int64
	_ = v7613
	var v7620 int32
	_ = v7620
	var v7621 int32
	_ = v7621
	var v7622 int32
	_ = v7622
	var v7623 int32
	_ = v7623
	var v7627 int64
	_ = v7627
	var v7633 int32
	_ = v7633
	var v7638 int32
	_ = v7638
	var v7641 int64
	_ = v7641
	var v7643 int64
	_ = v7643
	var v7644 int32
	_ = v7644
	var v7645 int64
	_ = v7645
	var v7648 int32
	_ = v7648
	var v7649 int32
	_ = v7649
	var v7654 int32
	_ = v7654
	var v7659 int32
	_ = v7659
	var v7665 int64
	_ = v7665
	var v7668 int64
	_ = v7668
	var v7669 int64
	_ = v7669
	var v7672 int64
	_ = v7672
	var v7678 int32
	_ = v7678
	var v7683 int32
	_ = v7683
	var v7689 int32
	_ = v7689
	var v7691 int32
	_ = v7691
	var v7694 int32
	_ = v7694
	var v7698 int32
	_ = v7698
	var v7699 int32
	_ = v7699
	var v7701 int32
	_ = v7701
	var v7702 int32
	_ = v7702
	var v7707 int32
	_ = v7707
	var v7708 int32
	_ = v7708
	var v7710 int32
	_ = v7710
	var v7711 int32
	_ = v7711
	var v7713 int32
	_ = v7713
	var v7714 int32
	_ = v7714
	var v7715 int32
	_ = v7715
	var v7716 int32
	_ = v7716
	var v7721 int32
	_ = v7721
	var v7728 int32
	_ = v7728
	var v7758 int32
	_ = v7758
	var v7760 int32
	_ = v7760
	var v7762 int32
	_ = v7762
	var v7764 int32
	_ = v7764
	var v7765 int32
	_ = v7765
	var v7767 int32
	_ = v7767
	var v7768 int32
	_ = v7768
	var v7772 int32
	_ = v7772
	var v7773 int32
	_ = v7773
	var v7777 int32
	_ = v7777
	var v7778 int32
	_ = v7778
	var v7780 int64
	_ = v7780
	var v7782 int32
	_ = v7782
	var v7784 int32
	_ = v7784
	var v7792 int32
	_ = v7792
	var v7795 int32
	_ = v7795
	var v7799 int32
	_ = v7799
	var v7800 int64
	_ = v7800
	var v7802 int32
	_ = v7802
	var v7806 int32
	_ = v7806
	var v7807 int32
	_ = v7807
	var v7810 int32
	_ = v7810
	var v7811 int32
	_ = v7811
	var v7816 int32
	_ = v7816
	var v7818 int32
	_ = v7818
	var v7822 int32
	_ = v7822
	var v7828 int32
	_ = v7828
	var v7830 int32
	_ = v7830
	var v7836 int32
	_ = v7836
	var v7840 int32
	_ = v7840
	var v7841 int64
	_ = v7841
	var v7843 int32
	_ = v7843
	var v7844 int64
	_ = v7844
	var v7847 int64
	_ = v7847
	var v7851 int32
	_ = v7851
	var v7852 int32
	_ = v7852
	var v7853 int32
	_ = v7853
	var v7857 int32
	_ = v7857
	var v7858 int32
	_ = v7858
	var v7860 int32
	_ = v7860
	var v7862 int32
	_ = v7862
	var v7863 int32
	_ = v7863
	var v7865 int32
	_ = v7865
	var v7867 int32
	_ = v7867
	var v7869 int32
	_ = v7869
	var v7870 int32
	_ = v7870
	var v7878 int32
	_ = v7878
	var v7879 int32
	_ = v7879
	var v7880 int32
	_ = v7880
	var v7883 int32
	_ = v7883
	var v7884 int32
	_ = v7884
	var v7886 int32
	_ = v7886
	var v7887 int32
	_ = v7887
	var v7889 int32
	_ = v7889
	var v7891 int32
	_ = v7891
	var v7901 int32
	_ = v7901
	var v7902 int32
	_ = v7902
	var v7903 int32
	_ = v7903
	var v7906 int32
	_ = v7906
	var v7907 int32
	_ = v7907
	var v7908 int32
	_ = v7908
	var v7911 int32
	_ = v7911
	var v7912 int32
	_ = v7912
	var v7914 int32
	_ = v7914
	var v7919 int32
	_ = v7919
	var v7932 int32
	_ = v7932
	var v7935 int32
	_ = v7935
	var v7936 int32
	_ = v7936
	var v7937 int32
	_ = v7937
	var v7974 int32
	_ = v7974
	var v7977 int32
	_ = v7977
	var v7980 int32
	_ = v7980
	var v7982 int32
	_ = v7982
	var v7989 int32
	_ = v7989
	var v7991 int32
	_ = v7991
	var v7992 int64
	_ = v7992
	var v7994 int64
	_ = v7994
	var v7997 int32
	_ = v7997
	var v8001 int32
	_ = v8001
	var v8005 int32
	_ = v8005
	var v8010 int32
	_ = v8010
	var v8012 int32
	_ = v8012
	var v8014 int32
	_ = v8014
	var v8017 int32
	_ = v8017
	var v8019 int32
	_ = v8019
	var v8020 int64
	_ = v8020
	var v8022 int32
	_ = v8022
	var v8023 int32
	_ = v8023
	var v8025 int32
	_ = v8025
	var v8030 int32
	_ = v8030
	var v8034 int32
	_ = v8034
	var v8035 int32
	_ = v8035
	var v8036 int32
	_ = v8036
	var v8037 int32
	_ = v8037
	var v8039 int32
	_ = v8039
	var v8050 int32
	_ = v8050
	var v8052 int32
	_ = v8052
	var v8054 int32
	_ = v8054
	var v8057 int32
	_ = v8057
	var v8060 int32
	_ = v8060
	var v8063 int32
	_ = v8063
	var v8064 int32
	_ = v8064
	var v8067 int32
	_ = v8067
	var v8068 int32
	_ = v8068
	var v8071 int32
	_ = v8071
	var v8078 int32
	_ = v8078
	var v8079 int32
	_ = v8079
	var v8082 int32
	_ = v8082
	var v8091 int64
	_ = v8091
	var v8093 int32
	_ = v8093
	var v8100 int32
	_ = v8100
	var v8104 int32
	_ = v8104
	var v8116 int32
	_ = v8116
	var v8117 int32
	_ = v8117
	var v8118 int32
	_ = v8118
	var v8120 int32
	_ = v8120
	var v8124 int32
	_ = v8124
	var v8125 int32
	_ = v8125
	var v8127 int32
	_ = v8127
	var v8128 int32
	_ = v8128
	var v8129 int32
	_ = v8129
	var v8131 int32
	_ = v8131
	var v8137 int32
	_ = v8137
	var v8138 int32
	_ = v8138
	var v8139 int32
	_ = v8139
	var v8140 int32
	_ = v8140
	var v8143 int32
	_ = v8143
	var v8150 int32
	_ = v8150
	var v8151 int32
	_ = v8151
	var v8152 int32
	_ = v8152
	var v8155 int32
	_ = v8155
	var v8158 int32
	_ = v8158
	var v8163 int32
	_ = v8163
	var v8164 int32
	_ = v8164
	var v8166 int32
	_ = v8166
	var v8168 int32
	_ = v8168
	var v8172 int32
	_ = v8172
	var v8173 int32
	_ = v8173
	var v8174 int32
	_ = v8174
	var v8179 int32
	_ = v8179
	var v8180 int32
	_ = v8180
	var v8181 int32
	_ = v8181
	var v8184 int32
	_ = v8184
	var v8185 int32
	_ = v8185
	var v8186 int32
	_ = v8186
	var v8188 int32
	_ = v8188
	var v8192 int32
	_ = v8192
	var v8193 int32
	_ = v8193
	var v8195 int32
	_ = v8195
	var v8197 int32
	_ = v8197
	var v8199 int32
	_ = v8199
	var v8200 int32
	_ = v8200
	var v8203 int32
	_ = v8203
	var v8210 int32
	_ = v8210
	var v8215 int32
	_ = v8215
	var v8216 int32
	_ = v8216
	var v8220 int64
	_ = v8220
	var v8221 int32
	_ = v8221
	var v8222 int32
	_ = v8222
	var v8230 int32
	_ = v8230
	var v8235 int32
	_ = v8235
	var v8239 int32
	_ = v8239
	var v8244 int64
	_ = v8244
	var v8246 int64
	_ = v8246
	var v8249 int32
	_ = v8249
	var v8257 int32
	_ = v8257
	var v8264 int32
	_ = v8264
	var v8265 int32
	_ = v8265
	var v8269 int64
	_ = v8269
	var v8272 int64
	_ = v8272
	var v8278 int32
	_ = v8278
	var v8283 int32
	_ = v8283
	var v8288 int32
	_ = v8288
	var v8289 int32
	_ = v8289
	var v8290 int32
	_ = v8290
	var v8297 int32
	_ = v8297
	var v8300 int32
	_ = v8300
	var v8308 int32
	_ = v8308
	var v8309 int64
	_ = v8309
	var v8311 int32
	_ = v8311
	var v8314 int32
	_ = v8314
	var v8316 int32
	_ = v8316
	var v8323 int32
	_ = v8323
	var v8325 int32
	_ = v8325
	var v8327 int32
	_ = v8327
	var v8330 int32
	_ = v8330
	var v8332 int32
	_ = v8332
	var v8333 int64
	_ = v8333
	var v8335 int32
	_ = v8335
	var v8338 int32
	_ = v8338
	var v8339 int32
	_ = v8339
	var v8341 int32
	_ = v8341
	var v8342 int32
	_ = v8342
	var v8348 int64
	_ = v8348
	var v8353 int32
	_ = v8353
	var v8357 int32
	_ = v8357
	var v8359 int32
	_ = v8359
	var v8365 int32
	_ = v8365
	var v8368 int32
	_ = v8368
	var v8370 int32
	_ = v8370
	var v8376 int32
	_ = v8376
	var v8380 int32
	_ = v8380
	var v8382 int32
	_ = v8382
	var v8385 int32
	_ = v8385
	var v8389 int32
	_ = v8389
	var v8394 int32
	_ = v8394
	var v8395 int32
	_ = v8395
	var v8396 int32
	_ = v8396
	var v8399 int32
	_ = v8399
	var v8403 int32
	_ = v8403
	var v8408 int32
	_ = v8408
	var v8409 int32
	_ = v8409
	var v8410 int32
	_ = v8410
	var v8413 int32
	_ = v8413
	var v8417 int32
	_ = v8417
	var v8424 int32
	_ = v8424
	var v8427 int32
	_ = v8427
	var v8435 int32
	_ = v8435
	var v8436 int32
	_ = v8436
	var v8440 int32
	_ = v8440
	var v8441 int32
	_ = v8441
	var v8442 int32
	_ = v8442
	var v8447 int64
	_ = v8447
	var v8448 int64
	_ = v8448
	var v8456 int32
	_ = v8456
	var v8457 int32
	_ = v8457
	var v8458 int32
	_ = v8458
	var v8460 int32
	_ = v8460
	var v8461 int32
	_ = v8461
	var v8467 int64
	_ = v8467
	var v8472 int32
	_ = v8472
	var v8476 int32
	_ = v8476
	var v8478 int32
	_ = v8478
	var v8484 int32
	_ = v8484
	var v8487 int32
	_ = v8487
	var v8489 int32
	_ = v8489
	var v8495 int32
	_ = v8495
	var v8499 int32
	_ = v8499
	var v8501 int32
	_ = v8501
	var v8504 int32
	_ = v8504
	var v8508 int32
	_ = v8508
	var v8513 int32
	_ = v8513
	var v8514 int32
	_ = v8514
	var v8515 int32
	_ = v8515
	var v8518 int32
	_ = v8518
	var v8522 int32
	_ = v8522
	var v8529 int32
	_ = v8529
	var v8532 int32
	_ = v8532
	var v8540 int32
	_ = v8540
	var v8541 int32
	_ = v8541
	var v8545 int32
	_ = v8545
	var v8546 int32
	_ = v8546
	var v8547 int32
	_ = v8547
	var v8552 int64
	_ = v8552
	var v8553 int64
	_ = v8553
	var v8561 int32
	_ = v8561
	var v8562 int32
	_ = v8562
	var v8563 int32
	_ = v8563
	var v8565 int32
	_ = v8565
	var v8569 int32
	_ = v8569
	var v8575 int32
	_ = v8575
	var v8580 int32
	_ = v8580
	var v8588 int32
	_ = v8588
	var v8592 int32
	_ = v8592
	var v8593 int32
	_ = v8593
	var v8597 int32
	_ = v8597
	var v8599 int64
	_ = v8599
	var v8600 int32
	_ = v8600
	var v8601 int32
	_ = v8601
	var v8608 int32
	_ = v8608
	var v8613 int32
	_ = v8613
	var v8616 int32
	_ = v8616
	var v8617 int32
	_ = v8617
	var v8621 int32
	_ = v8621
	var v8623 int64
	_ = v8623
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8632 int32
	_ = v8632
	var v8637 int32
	_ = v8637
	var v8638 int32
	_ = v8638
	var v8645 int32
	_ = v8645
	var v8652 int32
	_ = v8652
	var v8653 int32
	_ = v8653
	var v8657 int32
	_ = v8657
	var v8662 int32
	_ = v8662
	var v8664 int32
	_ = v8664
	var v8667 int64
	_ = v8667
	var v8673 int32
	_ = v8673
	var v8681 int32
	_ = v8681
	var v8688 int32
	_ = v8688
	var v8693 int32
	_ = v8693
	var v8698 int32
	_ = v8698
	var v8705 int32
	_ = v8705
	var v8710 int32
	_ = v8710
	var v8714 int32
	_ = v8714
	var v8717 int64
	_ = v8717
	var v8720 int32
	_ = v8720
	var v8723 int64
	_ = v8723
	var v8729 int32
	_ = v8729
	var v8734 int32
	_ = v8734
	var v8738 int32
	_ = v8738
	var v8739 int64
	_ = v8739
	var v8741 int64
	_ = v8741
	var v8742 int64
	_ = v8742
	var v8746 int64
	_ = v8746
	var v8752 int32
	_ = v8752
	var v8757 int32
	_ = v8757
	var v8761 int32
	_ = v8761
	var v8764 int32
	_ = v8764
	var v8765 int32
	_ = v8765
	var v8771 int32
	_ = v8771
	var v8776 int32
	_ = v8776
	var v8780 int32
	_ = v8780
	var v8781 int32
	_ = v8781
	var v8783 int32
	_ = v8783
	var v8785 int64
	_ = v8785
	var v8787 int32
	_ = v8787
	var v8793 int32
	_ = v8793
	var v8798 int32
	_ = v8798
	var v8806 int32
	_ = v8806
	var v8808 int32
	_ = v8808
	var v8811 int32
	_ = v8811
	var v8812 int32
	_ = v8812
	var v8815 int32
	_ = v8815
	var v8816 int32
	_ = v8816
	var v8822 int32
	_ = v8822
	var v8827 int32
	_ = v8827
	var v8829 int64
	_ = v8829
	var v8839 int32
	_ = v8839
	var v8846 int32
	_ = v8846
	var v8847 int32
	_ = v8847
	var v8851 int32
	_ = v8851
	var v8853 int64
	_ = v8853
	var v8854 int32
	_ = v8854
	var v8855 int32
	_ = v8855
	var v8862 int32
	_ = v8862
	var v8867 int32
	_ = v8867
	var v8871 int32
	_ = v8871
	var v8873 int64
	_ = v8873
	var v8874 int32
	_ = v8874
	var v8875 int32
	_ = v8875
	var v8882 int32
	_ = v8882
	var v8887 int32
	_ = v8887
	var v8923 int32
	_ = v8923
	var v8926 int32
	_ = v8926
	var v8928 int32
	_ = v8928
	var v8931 int32
	_ = v8931
	var v8933 int32
	_ = v8933
	var v8936 int32
	_ = v8936
	var v8938 int32
	_ = v8938
	var v8945 int32
	_ = v8945
	var v8947 int32
	_ = v8947
	var v8948 int32
	_ = v8948
	var v8953 int32
	_ = v8953
	var v8958 int32
	_ = v8958
	var v8966 int32
	_ = v8966
	var v8997 int32
	_ = v8997
	var v9029 int32
	_ = v9029
	var v9032 int32
	_ = v9032
	var v9035 int32
	_ = v9035
	var v9039 int32
	_ = v9039
	var v9041 int32
	_ = v9041
	var v9044 int32
	_ = v9044
	var v9048 int32
	_ = v9048
	var v9051 int32
	_ = v9051
	var v9056 int32
	_ = v9056
	var v9057 int32
	_ = v9057
	var v9059 int32
	_ = v9059
	var v9060 int64
	_ = v9060
	var v9063 int32
	_ = v9063
	var v9064 int32
	_ = v9064
	var v9068 int64
	_ = v9068
	var v9074 int32
	_ = v9074
	var v9079 int32
	_ = v9079
	var v9082 int32
	_ = v9082
	var v9085 int32
	_ = v9085
	var v9087 int32
	_ = v9087
	var v9094 int32
	_ = v9094
	var v9096 int32
	_ = v9096
	var v9097 int64
	_ = v9097
	var v9098 int32
	_ = v9098
	var v9105 int32
	_ = v9105
	var v9106 int32
	_ = v9106
	var v9109 int32
	_ = v9109
	var v9110 int32
	_ = v9110
	var v9114 int32
	_ = v9114
	var v9119 int32
	_ = v9119
	var v9121 int32
	_ = v9121
	var v9130 int32
	_ = v9130
	var v9158 int32
	_ = v9158
	var v9164 int32
	_ = v9164
	var v9171 int32
	_ = v9171
	var v9175 int32
	_ = v9175
	var v9180 int32
	_ = v9180
	var v9184 int32
	_ = v9184
	var v9187 int32
	_ = v9187
	var v9191 int32
	_ = v9191
	var v9196 int32
	_ = v9196
	var v9232 int32
	_ = v9232
	var v9234 int32
	_ = v9234
	var v9237 int32
	_ = v9237
	var v9238 int32
	_ = v9238
	var v9240 int32
	_ = v9240
	var v9243 int32
	_ = v9243
	var v9246 int32
	_ = v9246
	var v9248 int32
	_ = v9248
	var v9255 int32
	_ = v9255
	var v9257 int32
	_ = v9257
	var v9258 int32
	_ = v9258
	var v9260 int32
	_ = v9260
	var v9263 int32
	_ = v9263
	var v9264 int32
	_ = v9264
	var v9270 int32
	_ = v9270
	var v9306 int32
	_ = v9306
	var v9310 int32
	_ = v9310
	var v9311 int32
	_ = v9311
	var v9317 int32
	_ = v9317
	var v9318 int32
	_ = v9318
	var v9323 int32
	_ = v9323
	var v9325 int32
	_ = v9325
	var v9329 int32
	_ = v9329
	var v9331 int32
	_ = v9331
	var v9334 int32
	_ = v9334
	var v9336 int32
	_ = v9336
	var v9343 int32
	_ = v9343
	var v9345 int32
	_ = v9345
	var v9346 int32
	_ = v9346
	var v9349 int32
	_ = v9349
	var v9354 int32
	_ = v9354
	var v9386 int32
	_ = v9386
	var v9391 int32
	_ = v9391
	var v9395 int32
	_ = v9395
	var v9399 int32
	_ = v9399
	var v9400 int32
	_ = v9400
	var v9402 int32
	_ = v9402
	var v9406 int32
	_ = v9406
	var v9413 int32
	_ = v9413
	var v9414 int32
	_ = v9414
	var v9415 int32
	_ = v9415
	var v9434 int64
	_ = v9434
	var v9443 int32
	_ = v9443
	var v9444 int32
	_ = v9444
	var v9447 int32
	_ = v9447
	var v9455 int32
	_ = v9455
	var v9456 int32
	_ = v9456
	var v9457 int32
	_ = v9457
	var v9460 int64
	_ = v9460
	var v9461 int64
	_ = v9461
	var v9470 int64
	_ = v9470
	var v9473 int32
	_ = v9473
	var v9478 int32
	_ = v9478
	var v9479 int32
	_ = v9479
	var v9483 int32
	_ = v9483
	var v9487 int32
	_ = v9487
	var v9489 int32
	_ = v9489
	var v9490 int32
	_ = v9490
	var v9491 int32
	_ = v9491
	var v9492 int64
	_ = v9492
	var v9494 int32
	_ = v9494
	var v9531 int32
	_ = v9531
	var v9535 int32
	_ = v9535
	var v9571 int32
	_ = v9571
	var v9574 int32
	_ = v9574
	var v9579 int32
	_ = v9579
	var v9580 int32
	_ = v9580
	var v9581 int32
	_ = v9581
	var v9583 int32
	_ = v9583
	var v9587 int32
	_ = v9587
	var v9588 int64
	_ = v9588
	var v9590 int32
	_ = v9590
	var v9592 int32
	_ = v9592
	var v9595 int32
	_ = v9595
	var v9596 int32
	_ = v9596
	var v9598 int32
	_ = v9598
	var v9599 int64
	_ = v9599
	var v9600 int32
	_ = v9600
	var v9603 int32
	_ = v9603
	var v9607 int32
	_ = v9607
	var v9610 int32
	_ = v9610
	var v9613 int32
	_ = v9613
	var v9619 int64
	_ = v9619
	var v9622 int32
	_ = v9622
	var v9623 int32
	_ = v9623
	var v9624 int32
	_ = v9624
	var v9626 int32
	_ = v9626
	var v9627 int32
	_ = v9627
	var v9632 int32
	_ = v9632
	var v9633 int64
	_ = v9633
	var v9637 int32
	_ = v9637
	var v9641 int32
	_ = v9641
	var v9646 int32
	_ = v9646
	var v9647 int32
	_ = v9647
	var v9653 int32
	_ = v9653
	var v9654 int32
	_ = v9654
	var v9656 int32
	_ = v9656
	var v9658 int64
	_ = v9658
	var v9659 int32
	_ = v9659
	var v9660 int32
	_ = v9660
	var v9664 int32
	_ = v9664
	var v9672 int32
	_ = v9672
	var v9673 int32
	_ = v9673
	var v9675 int64
	_ = v9675
	var v9680 int32
	_ = v9680
	var v9681 int32
	_ = v9681
	var v9684 int64
	_ = v9684
	var v9692 int32
	_ = v9692
	var v9693 int32
	_ = v9693
	var v9702 int32
	_ = v9702
	var v9703 int32
	_ = v9703
	var v9709 int32
	_ = v9709
	var v9710 int32
	_ = v9710
	var v9716 int32
	_ = v9716
	var v9717 int32
	_ = v9717
	var v9722 int32
	_ = v9722
	var v9723 int32
	_ = v9723
	var v9729 int64
	_ = v9729
	var v9732 int64
	_ = v9732
	var v9735 int32
	_ = v9735
	var v9738 int32
	_ = v9738
	var v9745 int32
	_ = v9745
	var v9748 int32
	_ = v9748
	var v9752 int32
	_ = v9752
	var v9754 int64
	_ = v9754
	var v9756 int64
	_ = v9756
	var v9760 int32
	_ = v9760
	var v9763 int32
	_ = v9763
	var v9766 int64
	_ = v9766
	var v9769 int32
	_ = v9769
	var v9775 int32
	_ = v9775
	var v9778 int32
	_ = v9778
	var v9782 int32
	_ = v9782
	var v9787 int32
	_ = v9787
	var v9790 int32
	_ = v9790
	var v9792 int32
	_ = v9792
	var v9794 int32
	_ = v9794
	var v9795 int32
	_ = v9795
	var v9797 int32
	_ = v9797
	var v9801 int32
	_ = v9801
	var v9802 int32
	_ = v9802
	var v9804 int32
	_ = v9804
	var v9805 int32
	_ = v9805
	var v9808 int32
	_ = v9808
	var v9812 int32
	_ = v9812
	var v9814 int32
	_ = v9814
	var v9817 int32
	_ = v9817
	var v9819 int32
	_ = v9819
	var v9820 int32
	_ = v9820
	var v9821 int32
	_ = v9821
	var v9823 int32
	_ = v9823
	var v9826 int32
	_ = v9826
	var v9827 int32
	_ = v9827
	var v9833 int32
	_ = v9833
	var v9838 int32
	_ = v9838
	var v9842 int32
	_ = v9842
	var v9846 int32
	_ = v9846
	var v9847 int64
	_ = v9847
	var v9848 int64
	_ = v9848
	var v9849 int64
	_ = v9849
	var v9854 int64
	_ = v9854
	var v9855 int64
	_ = v9855
	var v9858 int64
	_ = v9858
	var v9861 int32
	_ = v9861
	var v9866 int32
	_ = v9866
	var v9867 int32
	_ = v9867
	var v9869 int32
	_ = v9869
	var v9870 int32
	_ = v9870
	var v9873 int32
	_ = v9873
	var v9876 int32
	_ = v9876
	var v9881 int32
	_ = v9881
	var v9882 int32
	_ = v9882
	var v9883 int32
	_ = v9883
	var v9886 int32
	_ = v9886
	var v9887 int32
	_ = v9887
	var v9891 int32
	_ = v9891
	var v9898 int32
	_ = v9898
	var v9932 int32
	_ = v9932
	var v9943 int32
	_ = v9943
	var v9948 int32
	_ = v9948
	var v9951 int32
	_ = v9951
	var v9952 int32
	_ = v9952
	var v9957 int32
	_ = v9957
	var v9962 int32
	_ = v9962
	var v9972 int32
	_ = v9972
	var v9977 int32
	_ = v9977
	var v9979 int32
	_ = v9979
	var v9988 int32
	_ = v9988
	var v9993 int32
	_ = v9993
	var v9994 int32
	_ = v9994
	var v9998 int32
	_ = v9998
	var v10002 int32
	_ = v10002
	var v10004 int32
	_ = v10004
	var v10041 int32
	_ = v10041
	var v10046 int32
	_ = v10046
	var v10051 int32
	_ = v10051
	var v10055 int32
	_ = v10055
	var v10060 int32
	_ = v10060
	var v10066 int32
	_ = v10066
	var v10067 int32
	_ = v10067
	var v10069 int32
	_ = v10069
	var v10070 int32
	_ = v10070
	var v10074 int32
	_ = v10074
	var v10082 int32
	_ = v10082
	var v10087 int32
	_ = v10087
	var v10089 int32
	_ = v10089
	var v10092 int32
	_ = v10092
	var v10093 int32
	_ = v10093
	var v10094 int32
	_ = v10094
	var v10095 int32
	_ = v10095
	var v10102 int32
	_ = v10102
	var v10103 int32
	_ = v10103
	var v10107 int32
	_ = v10107
	var v10111 int32
	_ = v10111
	var v10116 int32
	_ = v10116
	var v10117 int32
	_ = v10117
	var v10118 int32
	_ = v10118
	var v10119 int32
	_ = v10119
	var v10157 int64
	_ = v10157
	var v10158 int64
	_ = v10158
	var v10159 int64
	_ = v10159
	var v10162 int64
	_ = v10162
	var v10165 int32
	_ = v10165
	var v10170 int32
	_ = v10170
	var v10171 int32
	_ = v10171
	var v10173 int32
	_ = v10173
	var v10174 int32
	_ = v10174
	var v10179 int32
	_ = v10179
	var v10180 int32
	_ = v10180
	var v10181 int32
	_ = v10181
	var v10186 int32
	_ = v10186
	var v10187 int32
	_ = v10187
	var v10189 int32
	_ = v10189
	var v10190 int32
	_ = v10190
	var v10191 int32
	_ = v10191
	var v10193 int32
	_ = v10193
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
	var v10207 int32
	_ = v10207
	var v10208 int32
	_ = v10208
	var v10212 int32
	_ = v10212
	var v10217 int32
	_ = v10217
	var v10222 int32
	_ = v10222
	var v10223 int32
	_ = v10223
	var v10229 int32
	_ = v10229
	var v10230 int32
	_ = v10230
	var v10238 int32
	_ = v10238
	var v10239 int32
	_ = v10239
	var v10244 int32
	_ = v10244
	var v10245 int32
	_ = v10245
	var v10249 int32
	_ = v10249
	var v10251 int32
	_ = v10251
	var v10252 int32
	_ = v10252
	var v10258 int32
	_ = v10258
	var v10260 int32
	_ = v10260
	var v10268 int32
	_ = v10268
	var v10300 int32
	_ = v10300
	var v10305 int32
	_ = v10305
	var v10309 int32
	_ = v10309
	var v10310 int32
	_ = v10310
	var v10312 int32
	_ = v10312
	var v10313 int32
	_ = v10313
	var v10314 int32
	_ = v10314
	var v10320 int32
	_ = v10320
	var v10324 int32
	_ = v10324
	var v10326 int32
	_ = v10326
	var v10331 int32
	_ = v10331
	var v10332 int32
	_ = v10332
	var v10337 int32
	_ = v10337
	var v10339 int32
	_ = v10339
	var v10345 int32
	_ = v10345
	var v10350 int32
	_ = v10350
	var v10351 int32
	_ = v10351
	var v10352 int32
	_ = v10352
	var v10354 int32
	_ = v10354
	var v10355 int32
	_ = v10355
	var v10358 int32
	_ = v10358
	var v10363 int32
	_ = v10363
	var v10365 int32
	_ = v10365
	var v10371 int32
	_ = v10371
	var v10376 int32
	_ = v10376
	var v10380 int32
	_ = v10380
	var v10382 int32
	_ = v10382
	var v10390 int32
	_ = v10390
	var v10395 int32
	_ = v10395
	var v10397 int32
	_ = v10397
	var v10404 int32
	_ = v10404
	var v10406 int32
	_ = v10406
	var v10414 int32
	_ = v10414
	var v10419 int32
	_ = v10419
	var v10423 int32
	_ = v10423
	var v10459 int64
	_ = v10459
	var v10462 int32
	_ = v10462
	var v10467 int32
	_ = v10467
	var v10468 int32
	_ = v10468
	var v10469 int32
	_ = v10469
	var v10474 int32
	_ = v10474
	var v10477 int32
	_ = v10477
	var v10479 int32
	_ = v10479
	var v10480 int32
	_ = v10480
	var v10481 int32
	_ = v10481
	var v10484 int32
	_ = v10484
	var v10489 int32
	_ = v10489
	var v10494 int32
	_ = v10494
	var v10498 int32
	_ = v10498
	var v10503 int32
	_ = v10503
	var v10509 int32
	_ = v10509
	var v10510 int32
	_ = v10510
	var v10512 int32
	_ = v10512
	var v10513 int32
	_ = v10513
	var v10517 int32
	_ = v10517
	var v10525 int32
	_ = v10525
	var v10530 int32
	_ = v10530
	var v10532 int32
	_ = v10532
	var v10535 int32
	_ = v10535
	var v10536 int32
	_ = v10536
	var v10539 int32
	_ = v10539
	var v10544 int32
	_ = v10544
	var v10545 int32
	_ = v10545
	var v10549 int32
	_ = v10549
	var v10550 int32
	_ = v10550
	var v10552 int32
	_ = v10552
	var v10557 int32
	_ = v10557
	var v10562 int32
	_ = v10562
	var v10563 int32
	_ = v10563
	var v10565 int32
	_ = v10565
	var v10570 int32
	_ = v10570
	var v10571 int32
	_ = v10571
	var v10573 int32
	_ = v10573
	var v10574 int32
	_ = v10574
	var v10577 int32
	_ = v10577
	var v10582 int32
	_ = v10582
	var v10584 int32
	_ = v10584
	var v10590 int32
	_ = v10590
	var v10595 int32
	_ = v10595
	var v10599 int32
	_ = v10599
	var v10601 int32
	_ = v10601
	var v10609 int32
	_ = v10609
	var v10614 int32
	_ = v10614
	var v10652 int32
	_ = v10652
	var v10654 int32
	_ = v10654
	var v10662 int32
	_ = v10662
	var v10667 int32
	_ = v10667
	var v10670 int32
	_ = v10670
	var v10671 int32
	_ = v10671
	var v10677 int32
	_ = v10677
	var v10682 int32
	_ = v10682
	var v10688 int32
	_ = v10688
	var v10718 int32
	_ = v10718
	var v10721 int32
	_ = v10721
	var v10723 int32
	_ = v10723
	var v10730 int32
	_ = v10730
	var v10732 int32
	_ = v10732
	var v10734 int32
	_ = v10734
	var v10736 int32
	_ = v10736
	var v10741 int64
	_ = v10741
	var v10742 int64
	_ = v10742
	var v10745 int32
	_ = v10745
	var v10747 int32
	_ = v10747
	var v10748 int64
	_ = v10748
	var v10749 int64
	_ = v10749
	var v10752 int64
	_ = v10752
	var v10753 int64
	_ = v10753
	var v10759 int32
	_ = v10759
	var v10761 int64
	_ = v10761
	var v10769 int32
	_ = v10769
	var v10782 int64
	_ = v10782
	var v10789 int32
	_ = v10789
	var v10791 int64
	_ = v10791
	var v10793 int64
	_ = v10793
	var v10796 int32
	_ = v10796
	var v10797 int64
	_ = v10797
	var v10803 int64
	_ = v10803
	var v10822 int64
	_ = v10822
	var v10830 int64
	_ = v10830
	var v10836 int32
	_ = v10836
	var v10839 int32
	_ = v10839
	var v10843 int64
	_ = v10843
	var v10844 int32
	_ = v10844
	var v10847 int32
	_ = v10847
	var v10848 int64
	_ = v10848
	var v10850 int32
	_ = v10850
	var v10851 int32
	_ = v10851
	var v10854 int32
	_ = v10854
	var v10858 int32
	_ = v10858
	var v10862 int64
	_ = v10862
	var v10863 int64
	_ = v10863
	var v10866 int64
	_ = v10866
	var v10867 int64
	_ = v10867
	var v10871 int32
	_ = v10871
	var v10872 int32
	_ = v10872
	var v10876 int64
	_ = v10876
	var v10883 int64
	_ = v10883
	var v10884 int32
	_ = v10884
	var v10885 int32
	_ = v10885
	var v10887 int64
	_ = v10887
	var v10889 int32
	_ = v10889
	var v10891 int64
	_ = v10891
	var v10893 int32
	_ = v10893
	var v10896 int32
	_ = v10896
	var v10900 int64
	_ = v10900
	var v10902 int32
	_ = v10902
	var v10914 int64
	_ = v10914
	var v10921 int32
	_ = v10921
	var v10922 int32
	_ = v10922
	var v10925 int32
	_ = v10925
	var v10926 int32
	_ = v10926
	var v10929 int32
	_ = v10929
	var v10931 int32
	_ = v10931
	var v10938 int32
	_ = v10938
	var v10940 int64
	_ = v10940
	var v10942 int32
	_ = v10942
	var v10946 int32
	_ = v10946
	var v10950 int32
	_ = v10950
	var v10951 int32
	_ = v10951
	var v10953 int32
	_ = v10953
	var v10954 int64
	_ = v10954
	var v10956 int64
	_ = v10956
	var v10993 int64
	_ = v10993
	var v11001 int64
	_ = v11001
	var v11041 int32
	_ = v11041
	var v11045 int32
	_ = v11045
	var v11047 int32
	_ = v11047
	var v11051 int32
	_ = v11051
	var v11053 int32
	_ = v11053
	var v11054 int32
	_ = v11054
	var v11056 int32
	_ = v11056
	var v11057 int64
	_ = v11057
	var v11061 int64
	_ = v11061
	var v11064 int32
	_ = v11064
	var v11065 int32
	_ = v11065
	var v11068 int32
	_ = v11068
	var v11070 int32
	_ = v11070
	var v11071 int32
	_ = v11071
	var v11072 int32
	_ = v11072
	var v11074 int32
	_ = v11074
	var v11077 int32
	_ = v11077
	var v11078 int32
	_ = v11078
	var v11080 int32
	_ = v11080
	var v11081 int32
	_ = v11081
	var v11082 int32
	_ = v11082
	var v11085 int32
	_ = v11085
	var v11087 int32
	_ = v11087
	var v11088 int32
	_ = v11088
	var v11089 int32
	_ = v11089
	var v11090 int32
	_ = v11090
	var v11091 int32
	_ = v11091
	var v11098 int32
	_ = v11098
	var v11101 int32
	_ = v11101
	var v11102 int32
	_ = v11102
	var v11105 int32
	_ = v11105
	var v11119 int32
	_ = v11119
	var v11121 int32
	_ = v11121
	var v11123 int32
	_ = v11123
	var v11130 int32
	_ = v11130
	var v11143 int32
	_ = v11143
	var v11144 int32
	_ = v11144
	var v11146 int32
	_ = v11146
	var v11155 int32
	_ = v11155
	var v11157 int32
	_ = v11157
	var v11161 int32
	_ = v11161
	var v11162 int32
	_ = v11162
	var v11164 int32
	_ = v11164
	var v11165 int32
	_ = v11165
	var v11166 int32
	_ = v11166
	var v11167 int32
	_ = v11167
	var v11168 int32
	_ = v11168
	var v11170 int32
	_ = v11170
	var v11174 int32
	_ = v11174
	var v11175 int32
	_ = v11175
	var v11176 int32
	_ = v11176
	var v11178 int32
	_ = v11178
	var v11179 int64
	_ = v11179
	var v11181 int64
	_ = v11181
	var v11183 int32
	_ = v11183
	var v11184 int32
	_ = v11184
	var v11186 int32
	_ = v11186
	var v11187 int32
	_ = v11187
	var v11190 int32
	_ = v11190
	var v11192 int32
	_ = v11192
	var v11193 int32
	_ = v11193
	var v11195 int32
	_ = v11195
	var v11199 int32
	_ = v11199
	var v11200 int32
	_ = v11200
	var v11203 int32
	_ = v11203
	var v11204 int32
	_ = v11204
	var v11205 int32
	_ = v11205
	var v11207 int32
	_ = v11207
	var v11208 int32
	_ = v11208
	var v11209 int32
	_ = v11209
	var v11212 int32
	_ = v11212
	var v11214 int32
	_ = v11214
	var v11215 int32
	_ = v11215
	var v11222 int32
	_ = v11222
	var v11231 int32
	_ = v11231
	var v11233 int32
	_ = v11233
	var v11235 int32
	_ = v11235
	var v11242 int32
	_ = v11242
	var v11248 int32
	_ = v11248
	var v11258 int32
	_ = v11258
	var v11259 int32
	_ = v11259
	var v11261 int32
	_ = v11261
	var v11264 int32
	_ = v11264
	var v11266 int32
	_ = v11266
	var v11268 int32
	_ = v11268
	var v11269 int64
	_ = v11269
	var v11271 int64
	_ = v11271
	var v11273 int32
	_ = v11273
	var v11275 int32
	_ = v11275
	var v11277 int32
	_ = v11277
	var v11278 int32
	_ = v11278
	var v11280 int32
	_ = v11280
	var v11281 int32
	_ = v11281
	var v11284 int32
	_ = v11284
	var v11286 int32
	_ = v11286
	var v11287 int32
	_ = v11287
	var v11290 int32
	_ = v11290
	var v11291 int32
	_ = v11291
	var v11292 int32
	_ = v11292
	var v11295 int32
	_ = v11295
	var v11300 int32
	_ = v11300
	var v11302 int32
	_ = v11302
	var v11303 int32
	_ = v11303
	var v11307 int32
	_ = v11307
	var v11308 int32
	_ = v11308
	var v11327 int32
	_ = v11327
	var v11333 int32
	_ = v11333
	var v11340 int32
	_ = v11340
	var v11341 int32
	_ = v11341
	var v11343 int32
	_ = v11343
	var v11346 int32
	_ = v11346
	var v11353 int32
	_ = v11353
	var v11357 int32
	_ = v11357
	var v11358 int32
	_ = v11358
	var v11360 int32
	_ = v11360
	var v11361 int32
	_ = v11361
	var v11364 int32
	_ = v11364
	var v11368 int32
	_ = v11368
	var v11371 int32
	_ = v11371
	var v11372 int32
	_ = v11372
	var v11373 int32
	_ = v11373
	var v11375 int32
	_ = v11375
	var v11378 int32
	_ = v11378
	var v11382 int32
	_ = v11382
	var v11383 int32
	_ = v11383
	var v11385 int32
	_ = v11385
	var v11386 int32
	_ = v11386
	var v11395 int32
	_ = v11395
	var v11401 int32
	_ = v11401
	var v11426 int32
	_ = v11426
	var v11427 int32
	_ = v11427
	var v11428 int64
	_ = v11428
	var v11429 int32
	_ = v11429
	var v11432 int32
	_ = v11432
	var v11433 int32
	_ = v11433
	var v11436 int32
	_ = v11436
	var v11437 int32
	_ = v11437
	var v11441 int32
	_ = v11441
	var v11446 int32
	_ = v11446
	var v11447 int32
	_ = v11447
	var v11448 int32
	_ = v11448
	var v11449 int32
	_ = v11449
	var v11450 int32
	_ = v11450
	var v11451 int32
	_ = v11451
	var v11452 int32
	_ = v11452
	var v11453 int32
	_ = v11453
	var v11455 int32
	_ = v11455
	var v11456 int64
	_ = v11456
	var v11457 int32
	_ = v11457
	var v11458 int32
	_ = v11458
	var v11460 int32
	_ = v11460
	var v11462 int32
	_ = v11462
	var v11466 int32
	_ = v11466
	var v11471 int32
	_ = v11471
	var v11474 int32
	_ = v11474
	var v11488 int32
	_ = v11488
	var v11498 int32
	_ = v11498
	var v11499 int32
	_ = v11499
	var v11500 int32
	_ = v11500
	var v11503 int32
	_ = v11503
	var v11504 int32
	_ = v11504
	var v11507 int32
	_ = v11507
	var v11512 int32
	_ = v11512
	var v11514 int32
	_ = v11514
	var v11519 int32
	_ = v11519
	var v11521 int32
	_ = v11521
	var v11523 int32
	_ = v11523
	var v11524 int32
	_ = v11524
	var v11525 int32
	_ = v11525
	var v11527 int32
	_ = v11527
	var v11532 int32
	_ = v11532
	var v11534 int32
	_ = v11534
	var v11538 int32
	_ = v11538
	var v11545 int32
	_ = v11545
	var v11573 int32
	_ = v11573
	var v11575 int32
	_ = v11575
	var v11578 int32
	_ = v11578
	var v11579 int32
	_ = v11579
	var v11580 int32
	_ = v11580
	var v11582 int32
	_ = v11582
	var v11583 int32
	_ = v11583
	var v11590 int32
	_ = v11590
	var v11593 int32
	_ = v11593
	var v11595 int32
	_ = v11595
	var v11597 int32
	_ = v11597
	var v11601 int32
	_ = v11601
	var v11602 int32
	_ = v11602
	var v11604 int32
	_ = v11604
	var v11608 int32
	_ = v11608
	var v11612 int32
	_ = v11612
	var v11617 int32
	_ = v11617
	var v11619 int32
	_ = v11619
	var v11623 int32
	_ = v11623
	var v11624 int32
	_ = v11624
	var v11660 int32
	_ = v11660
	var v11662 int32
	_ = v11662
	var v11663 int32
	_ = v11663
	var v11700 int32
	_ = v11700
	var v11704 int32
	_ = v11704
	var v11708 int32
	_ = v11708
	var v11710 int32
	_ = v11710
	var v11713 int32
	_ = v11713
	var v11718 int32
	_ = v11718
	var v11719 int32
	_ = v11719
	var v11720 int64
	_ = v11720
	var v11721 int32
	_ = v11721
	var v11722 int64
	_ = v11722
	var v11725 int32
	_ = v11725
	var v11726 int32
	_ = v11726
	var v11727 int32
	_ = v11727
	var v11729 int32
	_ = v11729
	var v11730 int32
	_ = v11730
	var v11735 int32
	_ = v11735
	var v11736 int64
	_ = v11736
	var v11741 int32
	_ = v11741
	var v11744 int32
	_ = v11744
	var v11749 int32
	_ = v11749
	var v11751 int32
	_ = v11751
	var v11753 int32
	_ = v11753
	var v11754 int32
	_ = v11754
	var v11756 int32
	_ = v11756
	var v11757 int32
	_ = v11757
	var v11759 int32
	_ = v11759
	var v11761 int32
	_ = v11761
	var v11763 int32
	_ = v11763
	var v11769 int32
	_ = v11769
	var v11770 int32
	_ = v11770
	var v11771 int32
	_ = v11771
	var v11775 int32
	_ = v11775
	var v11776 int32
	_ = v11776
	var v11777 int32
	_ = v11777
	var v11779 int32
	_ = v11779
	var v11783 int32
	_ = v11783
	var v11797 int32
	_ = v11797
	var v11802 int32
	_ = v11802
	var v11803 int32
	_ = v11803
	var v11805 int32
	_ = v11805
	var v11816 int32
	_ = v11816
	var v11823 int64
	_ = v11823
	var v11826 int32
	_ = v11826
	var v11831 int32
	_ = v11831
	var v11832 int64
	_ = v11832
	var v11833 int64
	_ = v11833
	var v11834 int32
	_ = v11834
	var v11838 int64
	_ = v11838
	var v11839 int64
	_ = v11839
	var v11841 int64
	_ = v11841
	var v11847 int64
	_ = v11847
	var v11848 int64
	_ = v11848
	var v11849 int64
	_ = v11849
	var v11859 int64
	_ = v11859
	var v11861 int64
	_ = v11861
	var v11865 int64
	_ = v11865
	var v11867 int32
	_ = v11867
	var v11869 int32
	_ = v11869
	var v11874 int32
	_ = v11874
	var v11879 int32
	_ = v11879
	var v11881 int32
	_ = v11881
	var v11883 int32
	_ = v11883
	var v11887 int32
	_ = v11887
	var v11892 int32
	_ = v11892
	var v11893 int32
	_ = v11893
	var v11896 int32
	_ = v11896
	var v11898 int32
	_ = v11898
	var v11902 int32
	_ = v11902
	var v11904 int32
	_ = v11904
	var v11905 int32
	_ = v11905
	var v11906 int32
	_ = v11906
	var v11908 int32
	_ = v11908
	var v11911 int32
	_ = v11911
	var v11913 int32
	_ = v11913
	var v11918 int32
	_ = v11918
	var v11919 int32
	_ = v11919
	var v11920 int32
	_ = v11920
	var v11923 int64
	_ = v11923
	var v11924 int64
	_ = v11924
	var v11938 int32
	_ = v11938
	var v11941 int64
	_ = v11941
	var v11942 int32
	_ = v11942
	var v11944 int64
	_ = v11944
	var v11947 int32
	_ = v11947
	var v11948 int32
	_ = v11948
	var v11950 int32
	_ = v11950
	var v11960 int32
	_ = v11960
	var v11963 int32
	_ = v11963
	var v11964 int32
	_ = v11964
	var v11968 int32
	_ = v11968
	var v11972 int32
	_ = v11972
	var v11977 int32
	_ = v11977
	var v11978 int32
	_ = v11978
	var v11982 int32
	_ = v11982
	var v11987 int32
	_ = v11987
	var v11988 int32
	_ = v11988
	var v11990 int32
	_ = v11990
	var v11997 int32
	_ = v11997
	var v11998 int32
	_ = v11998
	var v11999 int32
	_ = v11999
	var v12002 int64
	_ = v12002
	var v12003 int64
	_ = v12003
	var v12014 int32
	_ = v12014
	var v12017 int32
	_ = v12017
	var v12019 int32
	_ = v12019
	var v12020 int32
	_ = v12020
	var v12022 int32
	_ = v12022
	var v12025 int32
	_ = v12025
	var v12026 int32
	_ = v12026
	var v12027 int32
	_ = v12027
	var v12029 int32
	_ = v12029
	var v12034 int32
	_ = v12034
	var v12039 int32
	_ = v12039
	var v12042 int64
	_ = v12042
	var v12043 int32
	_ = v12043
	var v12045 int32
	_ = v12045
	var v12047 int32
	_ = v12047
	var v12051 int32
	_ = v12051
	var v12052 int32
	_ = v12052
	var v12054 int32
	_ = v12054
	var v12056 int32
	_ = v12056
	var v12059 int32
	_ = v12059
	var v12061 int32
	_ = v12061
	var v12063 int32
	_ = v12063
	var v12067 int32
	_ = v12067
	var v12068 int32
	_ = v12068
	var v12070 int32
	_ = v12070
	var v12076 int32
	_ = v12076
	var v12077 int32
	_ = v12077
	var v12081 int32
	_ = v12081
	var v12083 int32
	_ = v12083
	var v12084 int32
	_ = v12084
	var v12087 int32
	_ = v12087
	var v12088 int32
	_ = v12088
	var v12091 int32
	_ = v12091
	var v12092 int32
	_ = v12092
	var v12095 int32
	_ = v12095
	var v12096 int32
	_ = v12096
	var v12099 int32
	_ = v12099
	var v12100 int32
	_ = v12100
	var v12103 int32
	_ = v12103
	var v12104 int32
	_ = v12104
	var v12107 int32
	_ = v12107
	var v12108 int32
	_ = v12108
	var v12111 int32
	_ = v12111
	var v12112 int32
	_ = v12112
	var v12115 int32
	_ = v12115
	var v12122 int32
	_ = v12122
	var v12125 int32
	_ = v12125
	var v12128 int32
	_ = v12128
	var v12131 int32
	_ = v12131
	var v12134 int32
	_ = v12134
	var v12137 int32
	_ = v12137
	var v12140 int32
	_ = v12140
	var v12143 int32
	_ = v12143
	var v12148 int32
	_ = v12148
	var v12151 int64
	_ = v12151
	var v12152 int32
	_ = v12152
	var v12154 int32
	_ = v12154
	var v12156 int32
	_ = v12156
	var v12160 int32
	_ = v12160
	var v12161 int32
	_ = v12161
	var v12163 int32
	_ = v12163
	var v12165 int32
	_ = v12165
	var v12168 int32
	_ = v12168
	var v12171 int32
	_ = v12171
	var v12174 int32
	_ = v12174
	var v12177 int32
	_ = v12177
	var v12180 int32
	_ = v12180
	var v12183 int32
	_ = v12183
	var v12186 int32
	_ = v12186
	var v12189 int32
	_ = v12189
	var v12191 int32
	_ = v12191
	var v12193 int32
	_ = v12193
	var v12197 int32
	_ = v12197
	var v12200 int32
	_ = v12200
	var v12204 int32
	_ = v12204
	var v12207 int32
	_ = v12207
	var v12214 int32
	_ = v12214
	var v12216 int32
	_ = v12216
	var v12218 int32
	_ = v12218
	var v12226 int32
	_ = v12226
	var v12232 int64
	_ = v12232
	var v12233 int64
	_ = v12233
	var v12235 int64
	_ = v12235
	var v12236 int64
	_ = v12236
	var v12239 int64
	_ = v12239
	var v12242 int32
	_ = v12242
	var v12247 int32
	_ = v12247
	var v12248 int32
	_ = v12248
	var v12249 int32
	_ = v12249
	var v12251 int32
	_ = v12251
	var v12257 int32
	_ = v12257
	var v12262 int32
	_ = v12262
	var v12263 int32
	_ = v12263
	var v12264 int32
	_ = v12264
	var v12266 int32
	_ = v12266
	var v12269 int32
	_ = v12269
	var v12279 int32
	_ = v12279
	var v12280 int32
	_ = v12280
	var v12283 int32
	_ = v12283
	var v12291 int32
	_ = v12291
	var v12292 int32
	_ = v12292
	var v12295 int32
	_ = v12295
	var v12298 int32
	_ = v12298
	var v12303 int32
	_ = v12303
	var v12307 int32
	_ = v12307
	var v12311 int64
	_ = v12311
	var v12312 int64
	_ = v12312
	var v12313 int64
	_ = v12313
	var v12316 int64
	_ = v12316
	var v12319 int32
	_ = v12319
	var v12324 int32
	_ = v12324
	var v12325 int32
	_ = v12325
	var v12330 int32
	_ = v12330
	var v12335 int32
	_ = v12335
	var v12336 int32
	_ = v12336
	var v12339 int32
	_ = v12339
	var v12344 int32
	_ = v12344
	var v12345 int32
	_ = v12345
	var v12347 int32
	_ = v12347
	var v12349 int32
	_ = v12349
	var v12350 int32
	_ = v12350
	var v12352 int32
	_ = v12352
	var v12363 int32
	_ = v12363
	var v12367 int32
	_ = v12367
	var v12371 int32
	_ = v12371
	var v12372 int32
	_ = v12372
	var v12374 int32
	_ = v12374
	var v12375 int32
	_ = v12375
	var v12384 int32
	_ = v12384
	var v12390 int32
	_ = v12390
	var v12391 int32
	_ = v12391
	var v12393 int32
	_ = v12393
	var v12397 int32
	_ = v12397
	var v12399 int32
	_ = v12399
	var v12402 int32
	_ = v12402
	var v12406 int32
	_ = v12406
	var v12407 int32
	_ = v12407
	var v12409 int32
	_ = v12409
	var v12413 int32
	_ = v12413
	var v12416 int32
	_ = v12416
	var v12418 int32
	_ = v12418
	var v12425 int32
	_ = v12425
	var v12427 int32
	_ = v12427
	var v12430 int32
	_ = v12430
	var v12434 int32
	_ = v12434
	var v12436 int32
	_ = v12436
	var v12438 int32
	_ = v12438
	var v12440 int32
	_ = v12440
	var v12444 int32
	_ = v12444
	var v12446 int32
	_ = v12446
	var v12448 int32
	_ = v12448
	var v12449 int32
	_ = v12449
	var v12452 int32
	_ = v12452
	var v12455 int32
	_ = v12455
	var v12460 int32
	_ = v12460
	var v12463 int32
	_ = v12463
	var v12467 int32
	_ = v12467
	var v12472 int32
	_ = v12472
	var v12476 int32
	_ = v12476
	var v12479 int32
	_ = v12479
	var v12483 int32
	_ = v12483
	var v12488 int32
	_ = v12488
	var v12492 int32
	_ = v12492
	var v12494 int32
	_ = v12494
	var v12501 int32
	_ = v12501
	var v12506 int32
	_ = v12506
	var v12510 int32
	_ = v12510
	var v12512 int32
	_ = v12512
	var v12520 int32
	_ = v12520
	var v12525 int32
	_ = v12525
	var v12529 int32
	_ = v12529
	var v12537 int32
	_ = v12537
	var v12542 int32
	_ = v12542
	var v12546 int32
	_ = v12546
	var v12549 int32
	_ = v12549
	var v12553 int32
	_ = v12553
	var v12557 int32
	_ = v12557
	var v12562 int32
	_ = v12562
	var v12566 int32
	_ = v12566
	var v12568 int32
	_ = v12568
	var v12576 int32
	_ = v12576
	var v12581 int32
	_ = v12581
	var v12585 int32
	_ = v12585
	var v12587 int32
	_ = v12587
	var v12595 int32
	_ = v12595
	var v12600 int32
	_ = v12600
	var v12602 int32
	_ = v12602
	var v12610 int32
	_ = v12610
	var v12615 int32
	_ = v12615
	var v12616 int32
	_ = v12616
	var v12617 int32
	_ = v12617
	var v12619 int32
	_ = v12619
	var v12620 int32
	_ = v12620
	var v12623 int32
	_ = v12623
	var v12628 int32
	_ = v12628
	var v12630 int32
	_ = v12630
	var v12636 int32
	_ = v12636
	var v12641 int32
	_ = v12641
	var v12645 int32
	_ = v12645
	var v12647 int32
	_ = v12647
	var v12655 int32
	_ = v12655
	var v12660 int32
	_ = v12660
	var v12664 int32
	_ = v12664
	var v12666 int32
	_ = v12666
	var v12674 int32
	_ = v12674
	var v12679 int32
	_ = v12679
	var v12681 int32
	_ = v12681
	var v12683 int32
	_ = v12683
	var v12685 int32
	_ = v12685
	var v12687 int32
	_ = v12687
	var v12693 int32
	_ = v12693
	var v12695 int32
	_ = v12695
	var v12701 int32
	_ = v12701
	var v12706 int32
	_ = v12706
	var v12712 int32
	_ = v12712
	var v12716 int32
	_ = v12716
	var v12721 int32
	_ = v12721
	var v12725 int32
	_ = v12725
	var v12728 int64
	_ = v12728
	var v12734 int32
	_ = v12734
	var v12739 int32
	_ = v12739
	var v12743 int32
	_ = v12743
	var v12746 int64
	_ = v12746
	var v12752 int32
	_ = v12752
	var v12757 int32
	_ = v12757
	var v12761 int32
	_ = v12761
	var v12763 int64
	_ = v12763
	var v12766 int64
	_ = v12766
	var v12772 int32
	_ = v12772
	var v12777 int32
	_ = v12777
	var v12782 int32
	_ = v12782
	var v12786 int32
	_ = v12786
	var v12791 int32
	_ = v12791
	v1 = int32(0)
	v35 = m.G0
	v39 = (v35 - int32(_a_F_StartupXLOG_0)) & int32(-4096)
	m.G0 = v39
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[1])) = v43
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+32))
	if base.Ui64(int64(23)) < base.Ui64(v47&int64(8184)) {
		goto L18
	} else {
		goto L19
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12782 = m.ExcPending
	if v12782 != 0 {
		goto L32
	} else {
		goto L2674
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12761 = m.ExcPending
	if v12761 != 0 {
		goto L32
	} else {
		goto L2671
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12743 = m.ExcPending
	if v12743 != 0 {
		goto L32
	} else {
		goto L2668
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12725 = m.ExcPending
	if v12725 != 0 {
		goto L32
	} else {
		goto L2665
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12712 = m.ExcPending
	if v12712 != 0 {
		goto L32
	} else {
		goto L2662
	}
L6:
	;
	v12681 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v12683 = v39 + int32(_a_F_StartupXLOG_1)
	v12685 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	F_XLogFileName(m, v12683, v9823, v9849, v12685)
	mBase = m.M
	v12687 = m.ExcPending
	if v12687 != 0 {
		goto L32
	} else {
		goto L2657
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12664 = m.ExcPending
	if v12664 != 0 {
		goto L32
	} else {
		goto L2653
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12645 = m.ExcPending
	if v12645 != 0 {
		goto L32
	} else {
		goto L2649
	}
L9:
	;
	v12616 = int32(_a_F_StartupXLOG_2)
	v12617 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v12619 = v39 + int32(_a_F_StartupXLOG_3)
	v12620 = F_unlink(m, v12619)
	mBase = m.M
	if v12617 != 0 {
		goto L2642
	} else {
		goto L2643
	}
L10:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12602 = m.ExcPending
	if v12602 != 0 {
		goto L32
	} else {
		goto L2639
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12585 = m.ExcPending
	if v12585 != 0 {
		goto L32
	} else {
		goto L2635
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12566 = m.ExcPending
	if v12566 != 0 {
		goto L32
	} else {
		goto L2631
	}
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12546 = m.ExcPending
	if v12546 != 0 {
		goto L32
	} else {
		goto L2626
	}
L14:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12529 = m.ExcPending
	if v12529 != 0 {
		goto L32
	} else {
		goto L2623
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12510 = m.ExcPending
	if v12510 != 0 {
		goto L32
	} else {
		goto L2619
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12492 = m.ExcPending
	if v12492 != 0 {
		goto L32
	} else {
		goto L2615
	}
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12476 = m.ExcPending
	if v12476 != 0 {
		goto L32
	} else {
		goto L2611
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	switch v53 - int32(1) {
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
	v12460 = m.ExcPending
	if v12460 != 0 {
		goto L32
	} else {
		goto L2607
	}
L21:
	;
	v256 = v39 + int32(_a_F_StartupXLOG_4)
	v259 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_5), v256, int32(0))
	mBase = m.M
	goto L72
L22:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), v249, int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L32
	} else {
		goto L71
	}
L23:
	;
	v220 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L32
	} else {
		goto L65
	}
L24:
	;
	v186 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L32
	} else {
		goto L58
	}
L25:
	;
	v152 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L32
	} else {
		goto L51
	}
L26:
	;
	v122 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L32
	} else {
		goto L45
	}
L27:
	;
	v92 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L32
	} else {
		goto L39
	}
L28:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[5])))
	if v59 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v60 = int32(15)
	goto L31
L30:
	;
	v60 = int32(18)
	goto L31
L31:
	;
	v62 = F_errstart(m, v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return
L33:
	;
	if v62 == int32(0) {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v68
	v71 = F_palloc(m, int32(128))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v79 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v81 = F_pg_strftime(m, v71, int32(128), int32(_a_F_StartupXLOG_8), v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4000)) = v71
	F_errmsg(m, int32(_a_F_StartupXLOG_9), v39+int32(4000))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v249 = int32(_a_F_StartupXLOG_10)
	goto L22
L39:
	;
	if v92 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v98
	v101 = F_palloc(m, int32(128))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v109 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v111 = F_pg_strftime(m, v101, int32(128), int32(_a_F_StartupXLOG_8), v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4016)) = v101
	F_errmsg(m, int32(_a_F_StartupXLOG_11), v39+int32(4016))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	v249 = int32(_a_F_StartupXLOG_12)
	goto L22
L45:
	;
	if v122 == int32(0) {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v128
	v131 = F_palloc(m, int32(128))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v139 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L32
	} else {
		goto L48
	}
L48:
	;
	v141 = F_pg_strftime(m, v131, int32(128), int32(_a_F_StartupXLOG_8), v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4032)) = v131
	F_errmsg(m, int32(_a_F_StartupXLOG_13), v39+int32(4032))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v249 = int32(_a_F_StartupXLOG_14)
	goto L22
L51:
	;
	if v152 == int32(0) {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v158
	v161 = F_palloc(m, int32(128))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v169 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	v171 = F_pg_strftime(m, v161, int32(128), int32(_a_F_StartupXLOG_8), v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4048)) = v161
	F_errmsg(m, int32(_a_F_StartupXLOG_15), v39+int32(4048))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L32
	} else {
		goto L56
	}
L56:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_16), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L32
	} else {
		goto L57
	}
L57:
	;
	v249 = int32(_a_F_StartupXLOG_17)
	goto L22
L58:
	;
	if v186 == int32(0) {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v191)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v192
	v195 = F_palloc(m, int32(128))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L32
	} else {
		goto L60
	}
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v203 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v205 = F_pg_strftime(m, v195, int32(128), int32(_a_F_StartupXLOG_8), v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L32
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4064)) = v195
	F_errmsg(m, int32(_a_F_StartupXLOG_18), v39+int32(4064))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_19), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	v249 = int32(_a_F_StartupXLOG_20)
	goto L22
L65:
	;
	if v220 == int32(0) {
		goto L21
	} else {
		goto L66
	}
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v225)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v226
	v229 = F_palloc(m, int32(128))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L32
	} else {
		goto L67
	}
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[7]))
	v237 = F_pg_localtime(m, v39+int32(_a_F_StartupXLOG_1), v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	v239 = F_pg_strftime(m, v229, int32(128), int32(_a_F_StartupXLOG_8), v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4080)) = v229
	F_errmsg(m, int32(_a_F_StartupXLOG_21), v39+int32(4080))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v249 = int32(_a_F_StartupXLOG_22)
	goto L22
L71:
	;
	goto L21
L72:
	;
	if v259 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[8])))
	if v260&int32(_a_F_StartupXLOG_23) != int32(_a_F_StartupXLOG_0) {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	v266 = v39 + int32(_a_F_StartupXLOG_1)
	v270 = F_pg_snprintf(m, v266, int32(1024), int32(_a_F_StartupXLOG_24), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L32
	} else {
		goto L75
	}
L75:
	;
	v274 = F___fstatat(m, int32(-100), v266, v256, int32(0))
	mBase = m.M
	goto L77
L76:
	;
	v324 = v39 + int32(_a_F_StartupXLOG_1)
	v328 = F_pg_snprintf(m, v324, int32(1024), int32(_a_F_StartupXLOG_25), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L32
	} else {
		goto L94
	}
L77:
	;
	if v274 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[8])))
	if v277&int32(_a_F_StartupXLOG_23) == int32(_a_F_StartupXLOG_0) {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v301 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L32
	} else {
		goto L86
	}
L81:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L32
	} else {
		goto L82
	}
L82:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L32
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3936)) = v266
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v39+int32(3936))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L32
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_27), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L32
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	if v301 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3968)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_29), v39+int32(3968))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L32
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[9]))
	v320 = F_mkdir(m, v39+int32(_a_F_StartupXLOG_1), v319)
	mBase = m.M
	goto L92
L90:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_30), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L32
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v320 < int32(0) {
		goto L15
	} else {
		goto L93
	}
L93:
	;
	goto L76
L94:
	;
	v334 = F___fstatat(m, int32(-100), v324, v39+int32(_a_F_StartupXLOG_4), int32(0))
	mBase = m.M
	goto L96
L95:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[10]))
	if v382 != 0 {
		goto L112
	} else {
		goto L113
	}
L96:
	;
	if v334 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[8])))
	if v337&int32(_a_F_StartupXLOG_23) == int32(_a_F_StartupXLOG_0) {
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v359 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L32
	} else {
		goto L104
	}
L100:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L32
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3888)) = v324
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v39+int32(3888))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L32
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_31), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L32
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	if v359 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3920)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_29), v39+int32(3920))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L32
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[9]))
	v378 = F_mkdir(m, v39+int32(_a_F_StartupXLOG_1), v377)
	mBase = m.M
	goto L110
L108:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_32), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L32
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	if v378 < int32(0) {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	goto L95
L112:
	;
	F_RegisterTimeout(m, int32(12), int32(406))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L32
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	v391 = v389 - int32(3)
	if base.Ui32(v391) <= base.Ui32(int32(-3)) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v396 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L32
	} else {
		goto L119
	}
L117:
	;
	v667 = v388
	goto L118
L118:
	;
	v704 = m.G0
	v706 = v704 - int32(2176)
	m.G0 = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	if base.Ui32(v711) < base.Ui32(v710) {
		goto L184
	} else {
		goto L185
	}
L119:
	;
	if v396 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_33), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L32
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v408 = F_AllocateDir(m, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L32
	} else {
		goto L125
	}
L123:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3835), int32(_a_F_StartupXLOG_34))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L32
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v411 = F_ReadDir(m, v408, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L32
	} else {
		goto L126
	}
L126:
	;
	if v411 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v416 = v411
	goto L130
L128:
	;
	goto L129
L129:
	;
	F_FreeDir(m, v408)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L32
	} else {
		goto L154
	}
L130:
	;
	v448 = v416 + int32(19)
	v449 = int32(_a_F_StartupXLOG_35)
	goto L135
L131:
	;
	goto L129
L132:
	;
	v525 = F_ReadDir(m, v408, int32(_a_F_StartupXLOG_5))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L32
	} else {
		goto L152
	}
L133:
	;
	if v487-v488 != 0 {
		goto L132
	} else {
		goto L146
	}
L135:
	;
	goto L136
L136:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	if v456 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v457 = v448
	v458 = v449
	v459 = int32(9)
	v460 = v456
	goto L141
L138:
	;
	v483 = v449
	v487 = int32(0)
	goto L139
L139:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	goto L133
L140:
	;
	v483 = v478
	v487 = v480
	goto L139
L141:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if base.B2i32(v460 != v462)|base.B2i32(v462 == int32(0)) != 0 {
		v478 = v458
		v480 = v460
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v478 = v472
	v480 = int32(0)
	goto L140
L143:
	;
	v468 = v459 - int32(1)
	if v468 == int32(0) {
		v478 = v458
		v480 = v460
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v471 = int32(1)
	v472 = v458 + v471
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+1)))
	if v473 != 0 {
		v457 = v457 + v471
		v458 = v472
		v459 = v468
		v460 = v473
		goto L141
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3872)) = v448
	v498 = v39 + int32(_a_F_StartupXLOG_1)
	v503 = F_pg_snprintf(m, v498, int32(1024), int32(_a_F_StartupXLOG_36), v39+int32(3872))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L32
	} else {
		goto L147
	}
L147:
	;
	v505 = F_unlink(m, v498)
	mBase = m.M
	v508 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L32
	} else {
		goto L148
	}
L148:
	;
	if v508 == int32(0) {
		goto L132
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3856)) = v498
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_37), v39+int32(3856))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L32
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3847), int32(_a_F_StartupXLOG_34))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L32
	} else {
		goto L151
	}
L151:
	;
	goto L132
L152:
	;
	if v525 != 0 {
		v416 = v525
		goto L130
	} else {
		goto L153
	}
L153:
	;
	goto L131
L154:
	;
	v563 = m.G0
	v565 = v563 - int32(112)
	m.G0 = v565
	v568 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v568 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v576 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_5), v565+int32(16), int32(256))
	mBase = m.M
	goto L160
L156:
	;
	goto L157
L157:
	;
	m.G0 = v565 + int32(112)
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v667 = v663
	goto L118
L158:
	;
	F_walkdir(m, v646, int32(1093), int32(0), int32(15))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L32
	} else {
		goto L182
	}
L159:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1092), int32(1), int32(14))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L32
	} else {
		goto L180
	}
L160:
	;
	if v576 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v581 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L32
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v565)+20))
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L32
	} else {
		goto L173
	}
L164:
	;
	if v581 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L32
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L32
	} else {
		goto L171
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = int32(_a_F_StartupXLOG_5)
	F_errmsg(m, int32(_a_F_StartupXLOG_39), v565)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L32
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_40), int32(3635), int32(_a_F_StartupXLOG_41))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L32
	} else {
		goto L170
	}
L170:
	;
	goto L167
L171:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1092), int32(0), int32(14))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L32
	} else {
		goto L172
	}
L172:
	;
	goto L159
L173:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1092), int32(0), int32(14))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L32
	} else {
		goto L174
	}
L174:
	;
	if v603&int32(_a_F_StartupXLOG_23) != int32(_a_F_StartupXLOG_43) {
		goto L159
	} else {
		goto L175
	}
L175:
	;
	v616 = int32(_a_F_StartupXLOG_5)
	F_walkdir(m, v616, int32(1092), int32(0), int32(14))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L32
	} else {
		goto L176
	}
L176:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1092), int32(1), int32(14))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L32
	} else {
		goto L177
	}
L177:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L32
	} else {
		goto L178
	}
L178:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_42), int32(1093), int32(0), int32(15))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L32
	} else {
		goto L179
	}
L179:
	;
	v646 = v616
	goto L158
L180:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L32
	} else {
		goto L181
	}
L181:
	;
	v646 = int32(_a_F_StartupXLOG_42)
	goto L158
L182:
	;
	F_walkdir(m, int32(_a_F_StartupXLOG_38), int32(1093), int32(1), int32(15))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L32
	} else {
		goto L183
	}
L183:
	;
	goto L157
L184:
	;
	v713 = v710
	goto L186
L185:
	;
	v713 = v711
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12])) = v713
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[10]))
	if v716 != 0 {
		goto L196
	} else {
		goto L197
	}
L187:
	;
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2369 != int32(1) {
		goto L538
	} else {
		goto L539
	}
L188:
	;
	v2340 = int32(0)
	v2343 = v2316
	v2344 = v2317
	v2345 = v2318
	v2346 = v2319
	v2347 = v2320
	v2348 = v2321
	v2349 = v2322
	v2350 = v2323
	v2351 = v2324
	v2352 = v2325
	v2353 = v2326
	v2360 = v2327
	v2361 = v2328
	v2362 = v2329
	v2364 = v2331
	goto L187
L189:
	;
	v2291 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v2291 == int32(44) {
		v2316 = v1350
		v2317 = v1332
		v2318 = v1339
		v2319 = v1341
		v2320 = v1340
		v2321 = v1342
		v2322 = v1343
		v2323 = v1344
		v2324 = v1345
		v2325 = v1352
		v2326 = v1353
		v2327 = v1461
		v2328 = v1333
		v2329 = v1338
		v2331 = v1351
		goto L188
	} else {
		goto L533
	}
L190:
	;
	v2076 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_44), v706+int32(1152), int32(0))
	mBase = m.M
	goto L483
L191:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L32
	} else {
		goto L478
	}
L192:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v959 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L193:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v862 != 0 {
		goto L244
	} else {
		goto L245
	}
L194:
	;
	v837 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])) = uint8(v837)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])) = uint8(v837)
	v843 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v843 != 0 {
		goto L193
	} else {
		goto L238
	}
L195:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L32
	} else {
		goto L234
	}
L196:
	;
	v719 = v706 + int32(1152)
	v722 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_45), v719, int32(0))
	mBase = m.M
	goto L199
L197:
	;
	goto L198
L198:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v814 != int32(1) {
		goto L192
	} else {
		goto L233
	}
L199:
	;
	if v722 == int32(0) {
		goto L195
	} else {
		goto L200
	}
L200:
	;
	v726 = F_unlink(m, int32(_a_F_StartupXLOG_46))
	mBase = m.M
	v730 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_47), v719, int32(0))
	mBase = m.M
	goto L201
L201:
	;
	if v730 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v736 = F_BasicOpenFilePerm(m, int32(_a_F_StartupXLOG_47), int32(2), int32(384))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L32
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v766 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_48), v706+int32(1152), int32(0))
	mBase = m.M
	goto L216
L205:
	;
	if int32(0) <= v736 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v742 != int32(1) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	v759 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])) = uint8(v759)
	goto L194
L209:
	;
	v757 = F_close(m, v736)
	mBase = m.M
	goto L208
L210:
	;
	goto L209
L211:
	;
	goto L212
L212:
	;
	v747 = F_fsync(m, v736)
	mBase = m.M
	if v747 != int32(-1) {
		goto L210
	} else {
		goto L214
	}
L213:
	;
	goto L210
L214:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v751 == int32(27) {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	if v766 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v772 = F_BasicOpenFilePerm(m, int32(_a_F_StartupXLOG_48), int32(2), int32(384))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L32
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v799 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])) = uint8(v799)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])) = uint8(v799)
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])))
	if v805 != 0 {
		goto L194
	} else {
		goto L231
	}
L220:
	;
	if int32(0) <= v772 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v778 != int32(1) {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	goto L223
L223:
	;
	v795 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])) = uint8(v795)
	goto L219
L224:
	;
	v793 = F_close(m, v772)
	mBase = m.M
	goto L223
L225:
	;
	goto L224
L226:
	;
	goto L227
L227:
	;
	v783 = F_fsync(m, v772)
	mBase = m.M
	if v783 != int32(-1) {
		goto L225
	} else {
		goto L229
	}
L228:
	;
	goto L225
L229:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v787 == int32(27) {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])))
	if v807 == int32(0) {
		goto L192
	} else {
		goto L232
	}
L232:
	;
	v811 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])) = uint8(v811)
	goto L193
L233:
	;
	goto L193
L234:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L32
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+864)) = int32(_a_F_StartupXLOG_45)
	F_errmsg(m, int32(_a_F_StartupXLOG_49), v706+int32(864))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L32
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1060), int32(_a_F_StartupXLOG_51))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L32
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L32
	} else {
		goto L239
	}
L239:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L32
	} else {
		goto L240
	}
L240:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_52), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L32
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1124), int32(_a_F_StartupXLOG_51))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L32
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[18]))
	if v897 != 0 {
		goto L262
	} else {
		goto L263
	}
L244:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[19]))
	if v864 != 0 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[20]))
	if v889 == int32(0) {
		goto L191
	} else {
		goto L260
	}
L247:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	if v865 != 0 {
		goto L243
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[20]))
	if v867 != 0 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L249
L251:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	if v868 != 0 {
		goto L243
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v871 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L32
	} else {
		goto L255
	}
L254:
	;
	goto L253
L255:
	;
	if v871 == int32(0) {
		goto L243
	} else {
		goto L256
	}
L256:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_53), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L32
	} else {
		goto L257
	}
L257:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_54), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L32
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1142), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L32
	} else {
		goto L259
	}
L259:
	;
	goto L243
L260:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889))))
	if v892 == int32(0) {
		goto L191
	} else {
		goto L261
	}
L261:
	;
	goto L243
L262:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v906 == int32(2) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v899&int32(1) != 0 {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[18])) = int32(2)
	goto L262
L265:
	;
	v911 = int32(0)
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[23]))
	v916 = F_DirectFunctionCall3Coll(m, int32(411), v911, v913, v911, int32(-1))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L32
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v922 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[24]))
	switch v922 - int32(1) {
	case 0:
		goto L270
	case 1:
		goto L271
	default:
		goto L192
	}
L268:
	;
	v918 = *(*int64)(unsafe.Add(mBase, uint32(v916)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25])) = v918
	goto L267
L269:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12])) = v955
	goto L192
L270:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v952 = F_findNewestTimeLine(m, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L32
	} else {
		goto L279
	}
L271:
	;
	v925 = int32(1)
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[26]))
	if v927 == v925 {
		v955 = v925
		goto L269
	} else {
		goto L272
	}
L272:
	;
	v930 = F_existsTimeLineHistory(m, v927)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L32
	} else {
		goto L273
	}
L273:
	;
	if v930 != 0 {
		v955 = v927
		goto L269
	} else {
		goto L274
	}
L274:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L32
	} else {
		goto L275
	}
L275:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L32
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+848)) = v927
	F_errmsg(m, int32(_a_F_StartupXLOG_56), v706+int32(848))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L32
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1189), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L32
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v955 = v952
	goto L269
L280:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_OwnLatch(m, v963+int32(4))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L32
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v969 = F_palloc0(m, int32(12))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L32
	} else {
		goto L284
	}
L283:
	;
	goto L282
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+884)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+880)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+876)) = int32(412)
	v979 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v982 = F_XLogReaderAllocate(m, v979, v706+int32(876), v969)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L32
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28])) = v982
	if v982 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v667)))
	*(*int64)(unsafe.Add(mBase, uint32(v982)+16)) = v985
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v989 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v982)+116)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v982)+104)) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v982)+100)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v982)+112)) = v989
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v998 = m.G0
	v1000 = v998 - int32(48)
	m.G0 = v1000
	v1003 = F_palloc0(m, int32(136))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L32
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L32
	} else {
		goto L473
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1003))) = v997
	*(*int64)(unsafe.Add(mBase, uint32(v1000)+16)) = int64(171798691852)
	v1011 = F_hash_create(m, int32(_a_F_StartupXLOG_57), int32(1024), v1000, int32(40))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L32
	} else {
		goto L290
	}
L290:
	;
	v1014 = v1003 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+32)) = v1014
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+28)) = v1014
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+24)) = v1011
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1019)+56)) = int64(0)
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+128)) = v1025 - int32(1)
	m.G0 = v1000 + int32(48)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32])) = v1003
	v1036 = F_palloc(m, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L32
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33])) = v1036
	v1041 = F_palloc(m, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L32
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34])) = v1041
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = int64(0)
	v1048 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v1048
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v1048)
	v1055 = F_AllocateFile(m, int32(_a_F_StartupXLOG_59), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L32
	} else {
		goto L293
	}
L293:
	;
	if v1055 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v1060 == int32(44) {
		goto L190
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1082 = v706 + int32(1079)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+720)) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(v706)+716)) = v706 + int32(1088)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+712)) = v706 + int32(1084)
	v1091 = v706 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+708)) = v1091
	v1094 = v706 + int32(892)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+704)) = v1094
	v1099 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_61), v706+int32(704))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L32
	} else {
		goto L303
	}
L297:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L32
	} else {
		goto L298
	}
L298:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L32
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+832)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(832))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L32
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1258), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L32
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L32
	} else {
		goto L469
	}
L303:
	;
	if v1099 != int32(5) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1079)))
	if v1103 != int32(10) {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = v1107
	v1110 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+888)))
	v1111 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v1110 | v1111<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+688)) = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v706)+692)) = v1091
	*(*int32)(unsafe.Add(mBase, uint32(v706)+696)) = v1082
	v1122 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_64), v706+int32(688))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L32
	} else {
		goto L307
	}
L306:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L32
	} else {
		goto L465
	}
L307:
	;
	if v1122 != int32(3) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1079)))
	if v1126 != int32(10) {
		goto L306
	} else {
		goto L309
	}
L309:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v1130
	v1133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+888)))
	v1134 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = v1133 | v1134<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+672)) = v706 + int32(1056)
	v1145 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_65), v706+int32(672))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L32
	} else {
		goto L311
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+656)) = v706 + int32(1024)
	v1165 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_66), v706+int32(656))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L32
	} else {
		goto L314
	}
L311:
	;
	if v1145 != int32(1) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1149 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1064)))
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1056))
	if v1149|(v1150^int64(7234308641521824883)) != int64(0) {
		goto L310
	} else {
		goto L313
	}
L313:
	;
	v1157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v1157)
	goto L310
L314:
	;
	v1167 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1024))
	v1169 = v706 + int32(896)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+640)) = v1169
	v1174 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_67), v706+int32(640))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L32
	} else {
		goto L316
	}
L315:
	;
	v1198 = v706 + int32(1152)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+608)) = v1198
	v1203 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_68), v706+int32(608))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L32
	} else {
		goto L323
	}
L316:
	;
	if v1174 != int32(1) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1180 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L32
	} else {
		goto L318
	}
L318:
	;
	if v1180 == int32(0) {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+628)) = int32(_a_F_StartupXLOG_59)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+624)) = v1169
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_69), v706+int32(624))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L32
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1321), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L32
	} else {
		goto L321
	}
L321:
	;
	goto L315
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+576)) = v706 + int32(1080)
	v1232 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_70), v706+int32(576))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L32
	} else {
		goto L331
	}
L323:
	;
	if v1203 != int32(1) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1209 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L32
	} else {
		goto L325
	}
L325:
	;
	if v1209 == int32(0) {
		goto L322
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+596)) = int32(_a_F_StartupXLOG_59)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+592)) = v1198
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_71), v706+int32(592))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L32
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1326), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L32
	} else {
		goto L328
	}
L328:
	;
	goto L322
L329:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L32
	} else {
		goto L460
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+516)) = v706 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+512)) = v706 + int32(892)
	v1268 = F_fscanf(m, v1055, int32(_a_F_StartupXLOG_72), v706+int32(512))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L32
	} else {
		goto L338
	}
L331:
	;
	if v1232 != int32(1) {
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	if v1236 != v1237 {
		goto L329
	} else {
		goto L333
	}
L333:
	;
	v1241 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L32
	} else {
		goto L334
	}
L334:
	;
	if v1241 == int32(0) {
		goto L330
	} else {
		goto L335
	}
L335:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+528)) = v1245
	*(*int32)(unsafe.Add(mBase, uint32(v706)+532)) = int32(_a_F_StartupXLOG_59)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_73), v706+int32(528))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L32
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1343), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L32
	} else {
		goto L337
	}
L337:
	;
	goto L330
L338:
	;
	if v1268 <= int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1055)))
	goto L343
L340:
	;
	goto L341
L341:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L32
	} else {
		goto L455
	}
L342:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L32
	} else {
		goto L451
	}
L343:
	;
	if int32(base.Ui32(v1272)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1277 = F_FreeFile(m, v1055)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L32
	} else {
		goto L345
	}
L345:
	;
	if v1277 != 0 {
		goto L342
	} else {
		goto L346
	}
L346:
	;
	v1280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v1280)
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v1283 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1285 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v1285)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L32
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1291 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L32
	} else {
		goto L351
	}
L350:
	;
	goto L349
L351:
	;
	if v1291 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+480)) = v1294
	v1297 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+468)) = uint32(v1297)
	v1300 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+476)) = uint32(v1300)
	v1302 = int64(32)
	v1303 = int64(base.Ui64(v1297) >> (uint(v1302) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+464)) = uint32(v1303)
	v1306 = int64(base.Ui64(v1300) >> (uint(v1302) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+472)) = uint32(v1306)
	F_errmsg(m, int32(_a_F_StartupXLOG_74), v706+int32(464))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L32
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v1323 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v1325 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v1326 = F_ReadCheckpointRecord(m, v1321, v1323, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L32
	} else {
		goto L358
	}
L355:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(627), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L32
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	v1461 = base.B2i32(v1165 == int32(1)) & base.B2i32(v1167 == int64(34166655670121587))
	v1464 = F_AllocateFile(m, int32(_a_F_StartupXLOG_44), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L32
	} else {
		goto L380
	}
L358:
	;
	if v1326 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+96))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+64))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+8))
	v1333 = *(*int64)(unsafe.Add(mBase, uint32(v1331)))
	v1334 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+896)) = v1334
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+904)) = v1336
	v1338 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+24))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+32))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+36))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+40))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+44))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+48))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+52))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+56))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+1096)) = v1346
	v1348 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+1088)) = v1348
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326)+16)))
	v1351 = *(*int64)(unsafe.Add(mBase, uint32(v1331)+80))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+76))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+72))
	v1356 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L32
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L32
	} else {
		goto L376
	}
L362:
	;
	if v1356 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1359 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+452)) = uint32(v1359)
	v1362 = int64(base.Ui64(v1359) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+448)) = uint32(v1362)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_76), v706+int32(448))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L32
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1376 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v1376)
	v1379 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v1379) <= base.Ui64(v1333) {
		goto L357
	} else {
		goto L368
	}
L366:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(641), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L32
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	F_XLogPrefetcherBeginRead(m, v1382, v1333)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L32
	} else {
		goto L369
	}
L369:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v1389 = F_ReadRecord(m, v1386, int32(15), int32(0), v1332)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L32
	} else {
		goto L370
	}
L370:
	;
	if v1389 != 0 {
		goto L357
	} else {
		goto L371
	}
L371:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L32
	} else {
		goto L372
	}
L372:
	;
	v1396 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+92)) = uint32(v1396)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+84)) = uint32(v1333)
	v1399 = int64(32)
	v1400 = int64(base.Ui64(v1333) >> (uint(v1399) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+80)) = uint32(v1400)
	v1403 = int64(base.Ui64(v1396) >> (uint(v1399) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+88)) = uint32(v1403)
	F_errmsg(m, int32(_a_F_StartupXLOG_77), v706+int32(80))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L32
	} else {
		goto L373
	}
L373:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+64)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v706)+68)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v706)+72)) = v1411
	*(*int32)(unsafe.Add(mBase, uint32(v706)+76)) = v1411
	F_errhint(m, int32(_a_F_StartupXLOG_78), v706-int32(-64))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L32
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(661), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L32
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	v1431 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+52)) = uint32(v1431)
	v1434 = int64(base.Ui64(v1431) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+48)) = uint32(v1434)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v706+int32(48))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L32
	} else {
		goto L377
	}
L377:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+32)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v706)+36)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v706)+40)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v706)+44)) = v1442
	F_errhint(m, int32(_a_F_StartupXLOG_78), v706+int32(32))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L32
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(672), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L32
	} else {
		goto L379
	}
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L380:
	;
	if v1464 == int32(0) {
		goto L189
	} else {
		goto L381
	}
L381:
	;
	v1468 = F_do_getc(m, v1464)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L32
	} else {
		goto L383
	}
L382:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1464)))
	goto L426
L383:
	;
	if v1468 == int32(-1) {
		v1788 = v1
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1472 = int32(0)
	v1476 = v1
	v1479 = v1472
	v1480 = v1472
	v1481 = v1468
	goto L385
L385:
	;
	if v1480&int32(1) != 0 {
		goto L391
	} else {
		goto L392
	}
L386:
	;
	if base.B2i32(v1729 == int32(0))&(v1730^int32(1)) != 0 {
		v1788 = v1726
		goto L382
	} else {
		goto L420
	}
L387:
	;
	v1758 = F_do_getc(m, v1464)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L32
	} else {
		goto L418
	}
L388:
	;
	v1726 = v1691
	v1729 = v1694
	v1730 = int32(0)
	goto L387
L389:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L32
	} else {
		goto L414
	}
L390:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L32
	} else {
		goto L410
	}
L391:
	;
	v1640 = (v1480 ^ int32(1)) & base.B2i32(v1481 == int32(92))
	if v1640|base.B2i32(base.Ui32(int32(1022)) < base.Ui32(v1479)) != 0 {
		v1726 = v1476
		v1729 = v1479
		v1730 = v1640
		goto L387
	} else {
		goto L409
	}
L392:
	;
	switch v1481 - int32(10) {
	case 0, 3:
		goto L393
	default:
		goto L391
	}
L393:
	;
	if v1479 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1512 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v706+int32(1152)+v1479))) = uint8(v1512)
	v1524 = v1512
	goto L397
L395:
	;
	v1603 = v1476
	goto L396
L396:
	;
	v1691 = v1603
	v1694 = int32(0)
	goto L388
L397:
	;
	v1554 = v706 + int32(1152) + v1524
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1554))))
	v1556 = int32(32)
	if v1555|v1556 != v1556 {
		goto L399
	} else {
		goto L400
	}
L398:
	;
	if base.B2i32(v1524 <= int32(0))|base.B2i32(v1479-int32(1) <= v1524) != 0 {
		goto L390
	} else {
		goto L402
	}
L399:
	;
	v1524 = v1524 + int32(1)
	goto L397
L400:
	;
	goto L401
L401:
	;
	goto L398
L402:
	;
	v1568 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1554))) = uint8(v1568)
	v1571 = F_palloc0(m, int32(24))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L32
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v1582 = F_strtox_2(m, v706+int32(1152), v706+int32(1056), int32(10), int64(4294967295))
	mBase = m.M
	goto L404
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1571))) = base.I32_wrap_i64(v1582)
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1056))
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585))))
	if v1586 != 0 {
		goto L389
	} else {
		goto L405
	}
L405:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if base.B2i32(v1588 == int32(68))|base.B2i32(v1588 == int32(28)) != 0 {
		goto L389
	} else {
		goto L406
	}
L406:
	;
	v1596 = F_pstrdup(m, v1554+int32(1))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L32
	} else {
		goto L407
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1571)+4)) = v1596
	v1599 = F_lappend(m, v1476, v1571)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L32
	} else {
		goto L408
	}
L408:
	;
	v1603 = v1599
	goto L396
L409:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v706+int32(1152)+v1479))) = uint8(v1481)
	v1691 = v1476
	v1694 = v1479 + int32(1)
	goto L388
L410:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L32
	} else {
		goto L411
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+432)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(432))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L32
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1425), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L32
	} else {
		goto L413
	}
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L414:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L32
	} else {
		goto L415
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+416)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(416))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L32
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1434), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L32
	} else {
		goto L417
	}
L417:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L418:
	;
	if v1758 != int32(-1) {
		v1476 = v1726
		v1479 = v1729
		v1480 = v1730
		v1481 = v1758
		goto L385
	} else {
		goto L419
	}
L419:
	;
	goto L386
L420:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L32
	} else {
		goto L421
	}
L421:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L32
	} else {
		goto L422
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+400)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(400))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L32
	} else {
		goto L423
	}
L423:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1454), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L32
	} else {
		goto L424
	}
L424:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L425:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L32
	} else {
		goto L447
	}
L426:
	;
	if int32(base.Ui32(v1820)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1825 = F_FreeFile(m, v1464)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L32
	} else {
		goto L428
	}
L428:
	;
	if v1825 != 0 {
		goto L425
	} else {
		goto L429
	}
L429:
	;
	if v1788 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2340 = int32(1)
	v2343 = v1350
	v2344 = v1332
	v2345 = v1339
	v2346 = v1341
	v2347 = v1340
	v2348 = v1342
	v2349 = v1343
	v2350 = v1344
	v2351 = v1345
	v2352 = v1352
	v2353 = v1353
	v2360 = v1461
	v2361 = v1333
	v2362 = v1338
	v2364 = v1351
	goto L187
L431:
	;
	goto L432
L432:
	;
	v1830 = int32(1)
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	if v1831 <= int32(0) {
		v2340 = v1830
		v2343 = v1350
		v2344 = v1332
		v2345 = v1339
		v2346 = v1341
		v2347 = v1340
		v2348 = v1342
		v2349 = v1343
		v2350 = v1344
		v2351 = v1345
		v2352 = v1352
		v2353 = v1353
		v2360 = v1461
		v2361 = v1333
		v2362 = v1338
		v2364 = v1351
		goto L187
	} else {
		goto L433
	}
L433:
	;
	v1842 = int32(0)
	goto L434
L434:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+12))
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1869+v1842<<(uint(int32(2))%32))))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1873)))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+372)) = v1874
	*(*int32)(unsafe.Add(mBase, uint32(v706)+368)) = int32(_a_F_StartupXLOG_38)
	v1881 = F_psprintf(m, int32(_a_F_StartupXLOG_82), v706+int32(368))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L32
	} else {
		goto L437
	}
L435:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L32
	} else {
		goto L443
	}
L436:
	;
	goto L435
L437:
	;
	F_remove_tablespace_symlink(m, v1881)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L32
	} else {
		goto L438
	}
L438:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	v1886 = F_symlink(m, v1885, v1881)
	mBase = m.M
	if v1886 < int32(0) {
		goto L436
	} else {
		goto L439
	}
L439:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+4))
	F_pfree(m, v1889)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L32
	} else {
		goto L440
	}
L440:
	;
	F_pfree(m, v1873)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L32
	} else {
		goto L441
	}
L441:
	;
	v1895 = v1842 + int32(1)
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	if v1895 < v1896 {
		v1842 = v1895
		goto L434
	} else {
		goto L442
	}
L442:
	;
	v2340 = v1830
	v2343 = v1350
	v2344 = v1332
	v2345 = v1339
	v2346 = v1341
	v2347 = v1340
	v2348 = v1342
	v2349 = v1343
	v2350 = v1344
	v2351 = v1345
	v2352 = v1352
	v2353 = v1353
	v2360 = v1461
	v2361 = v1333
	v2362 = v1338
	v2364 = v1351
	goto L187
L443:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L32
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+352)) = v1881
	F_errmsg(m, int32(_a_F_StartupXLOG_83), v706+int32(352))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L32
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(698), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L32
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L32
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+384)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(384))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L32
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1460), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L32
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+496)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(496))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L32
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1356), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L32
	} else {
		goto L454
	}
L454:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L455:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L32
	} else {
		goto L456
	}
L456:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_84), int32(0))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L32
	} else {
		goto L457
	}
L457:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_85), int32(0))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L32
	} else {
		goto L458
	}
L458:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1350), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
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
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L32
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+560)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(560))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L32
	} else {
		goto L462
	}
L462:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+544)) = v1985
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+548)) = v1987
	F_errdetail(m, int32(_a_F_StartupXLOG_86), v706+int32(544))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L32
	} else {
		goto L463
	}
L463:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1339), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L32
	} else {
		goto L464
	}
L464:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L465:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L32
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+16)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(16))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L32
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1278), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
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
	F_errcode(m, int32(325))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L32
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706))) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L32
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1271), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
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
	F_errcode(m, int32(_a_F_StartupXLOG_87))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L32
	} else {
		goto L474
	}
L474:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_88), int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L32
	} else {
		goto L475
	}
L475:
	;
	F_errdetail(m, int32(_a_F_StartupXLOG_89), int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L32
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(572), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L32
	} else {
		goto L477
	}
L477:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L478:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L32
	} else {
		goto L479
	}
L479:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_90), int32(0))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L32
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1150), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L32
	} else {
		goto L481
	}
L481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L482:
	;
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2119 != int32(1) {
		goto L497
	} else {
		goto L498
	}
L483:
	;
	if v2076 != 0 {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2077 = int32(_a_F_StartupXLOG_91)
	v2078 = F_unlink(m, v2077)
	mBase = m.M
	v2082 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v2077, int32(14))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L32
	} else {
		goto L485
	}
L485:
	;
	v2086 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L32
	} else {
		goto L486
	}
L486:
	;
	if v2086 == int32(0) {
		goto L482
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+820)) = int32(_a_F_StartupXLOG_59)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+816)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_92), v706+int32(816))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L32
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+804)) = int32(_a_F_StartupXLOG_91)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+800)) = int32(_a_F_StartupXLOG_44)
	if v2082 != 0 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2105 = int32(_a_F_StartupXLOG_93)
	goto L491
L490:
	;
	v2105 = int32(_a_F_StartupXLOG_94)
	goto L491
L491:
	;
	F_errdetail(m, v2105, v706+int32(800))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L32
	} else {
		goto L492
	}
L492:
	;
	if v2082 != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v2113 = int32(739)
	goto L495
L494:
	;
	v2113 = int32(733)
	goto L495
L495:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2113, int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L32
	} else {
		goto L496
	}
L496:
	;
	goto L482
L497:
	;
	v2144 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	if v2144 == int64(0) {
		goto L506
	} else {
		goto L507
	}
L498:
	;
	v2122 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if v2122 != int64(0) {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v2133)
	v2136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2136 == int32(0) {
		goto L497
	} else {
		goto L504
	}
L500:
	;
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	if v2125 != 0 {
		goto L499
	} else {
		goto L501
	}
L501:
	;
	v2126 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	if v2126 != int64(0) {
		goto L499
	} else {
		goto L502
	}
L502:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	if v2129 != int32(1) {
		goto L497
	} else {
		goto L503
	}
L503:
	;
	goto L499
L504:
	;
	v2140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v2140)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L32
	} else {
		goto L505
	}
L505:
	;
	goto L497
L506:
	;
	v2170 = *(*int64)(unsafe.Add(mBase, uint32(v667)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = v2170
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v2173
	v2175 = *(*int64)(unsafe.Add(mBase, uint32(v667)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = v2173
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v2175
	v2181 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v2182 = F_ReadCheckpointRecord(m, v2181, v2170, v2173)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L32
	} else {
		goto L513
	}
L507:
	;
	v2149 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L32
	} else {
		goto L508
	}
L508:
	;
	if v2149 == int32(0) {
		goto L506
	} else {
		goto L509
	}
L509:
	;
	v2153 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+788)) = uint32(v2153)
	v2156 = int64(base.Ui64(v2153) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+784)) = uint32(v2156)
	F_errmsg(m, int32(_a_F_StartupXLOG_95), v706+int32(784))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L32
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(778), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L32
	} else {
		goto L511
	}
L511:
	;
	goto L506
L512:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L32
	} else {
		goto L530
	}
L513:
	;
	if v2182 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2186 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L32
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L32
	} else {
		goto L527
	}
L517:
	;
	if v2186 != 0 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2189 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+772)) = uint32(v2189)
	v2192 = int64(base.Ui64(v2189) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+768)) = uint32(v2192)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_76), v706+int32(768))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L32
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2206)+96))
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+64))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+8))
	v2210 = *(*int64)(unsafe.Add(mBase, uint32(v2208)))
	v2211 = *(*int64)(unsafe.Add(mBase, uint32(v2208)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+896)) = v2211
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+904)) = v2213
	v2215 = *(*int64)(unsafe.Add(mBase, uint32(v2208)+24))
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+32))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+36))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+40))
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+44))
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+48))
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+52))
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+56))
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+1096)) = v2223
	v2225 = *(*int64)(unsafe.Add(mBase, uint32(v2208)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+1088)) = v2225
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2182)+16)))
	v2228 = *(*int64)(unsafe.Add(mBase, uint32(v2208)+80))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+76))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+72))
	v2232 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v2232) <= base.Ui64(v2210) {
		v2316 = v2227
		v2317 = v2209
		v2318 = v2216
		v2319 = v2218
		v2320 = v2217
		v2321 = v2219
		v2322 = v2220
		v2323 = v2221
		v2324 = v2222
		v2325 = v2229
		v2326 = v2230
		v2327 = v1
		v2328 = v2210
		v2329 = v2215
		v2331 = v2228
		goto L188
	} else {
		goto L523
	}
L521:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(791), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L32
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	F_XLogPrefetcherBeginRead(m, v2235, v2210)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L32
	} else {
		goto L524
	}
L524:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v2242 = F_ReadRecord(m, v2239, int32(15), int32(0), v2209)
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L32
	} else {
		goto L525
	}
L525:
	;
	if v2242 == int32(0) {
		goto L512
	} else {
		goto L526
	}
L526:
	;
	v2316 = v2227
	v2317 = v2209
	v2318 = v2216
	v2319 = v2218
	v2320 = v2217
	v2321 = v2219
	v2322 = v2220
	v2323 = v2221
	v2324 = v2222
	v2325 = v2229
	v2326 = v2230
	v2327 = v1
	v2328 = v2210
	v2329 = v2215
	v2331 = v2228
	goto L188
L527:
	;
	v2251 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+740)) = uint32(v2251)
	v2254 = int64(base.Ui64(v2251) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+736)) = uint32(v2254)
	F_errmsg(m, int32(_a_F_StartupXLOG_96), v706+int32(736))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L32
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(803), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L32
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	v2271 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+764)) = uint32(v2271)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+756)) = uint32(v2210)
	v2274 = int64(32)
	v2275 = int64(base.Ui64(v2210) >> (uint(v2274) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+752)) = uint32(v2275)
	v2278 = int64(base.Ui64(v2271) >> (uint(v2274) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+760)) = uint32(v2278)
	F_errmsg(m, int32(_a_F_StartupXLOG_97), v706+int32(752))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L32
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(815), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
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
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L32
	} else {
		goto L534
	}
L534:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L32
	} else {
		goto L535
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+336)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(336))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L32
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1393), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L32
	} else {
		goto L537
	}
L537:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L538:
	;
	v2468 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v2470 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2471 = F_tliOfPointInHistory(m, v2468, v2470)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L32
	} else {
		goto L572
	}
L539:
	;
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2374 != 0 {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2460, int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L32
	} else {
		goto L567
	}
L541:
	;
	v2377 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L32
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v2387 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	v2390 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L32
	} else {
		goto L547
	}
L544:
	;
	if v2377 == int32(0) {
		goto L538
	} else {
		goto L545
	}
L545:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_98), int32(0))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L32
	} else {
		goto L546
	}
L546:
	;
	v2460 = int32(823)
	goto L540
L547:
	;
	switch v2387 - int32(1) {
	case 0:
		goto L553
	case 1:
		goto L552
	case 2:
		goto L551
	case 3:
		goto L550
	case 4:
		goto L549
	default:
		goto L548
	}
L548:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L565
	}
L549:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L563
	}
L550:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L561
	}
L551:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L559
	}
L552:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L556
	}
L553:
	;
	if v2390 == int32(0) {
		goto L538
	} else {
		goto L554
	}
L554:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+272)) = v2397
	F_errmsg(m, int32(_a_F_StartupXLOG_99), v706+int32(272))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L32
	} else {
		goto L555
	}
L555:
	;
	v2460 = int32(827)
	goto L540
L556:
	;
	v2408 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25]))
	v2409 = F_timestamptz_to_str(m, v2408)
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L32
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+288)) = v2409
	F_errmsg(m, int32(_a_F_StartupXLOG_100), v706+int32(288))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L32
	} else {
		goto L558
	}
L558:
	;
	v2460 = int32(831)
	goto L540
L559:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+304)) = v2421
	F_errmsg(m, int32(_a_F_StartupXLOG_101), v706+int32(304))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L32
	} else {
		goto L560
	}
L560:
	;
	v2460 = int32(835)
	goto L540
L561:
	;
	v2432 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+324)) = uint32(v2432)
	v2435 = int64(base.Ui64(v2432) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+320)) = uint32(v2435)
	F_errmsg(m, int32(_a_F_StartupXLOG_102), v706+int32(320))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L32
	} else {
		goto L562
	}
L562:
	;
	v2460 = int32(839)
	goto L540
L563:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_103), int32(0))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L32
	} else {
		goto L564
	}
L564:
	;
	v2460 = int32(842)
	goto L540
L565:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_104), int32(0))
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L32
	} else {
		goto L566
	}
L566:
	;
	v2460 = int32(845)
	goto L540
L567:
	;
	goto L538
L568:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+120))
	v2892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2890)+56)))
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+48))
	v2894 = *(*int64)(unsafe.Add(mBase, uint32(v2890)+40))
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+116))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+112))
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+96))
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+92))
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+88))
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+84))
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+80))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+76))
	v2903 = *(*int64)(unsafe.Add(mBase, uint32(v2890)+64))
	v2904 = int32(_a_F_StartupXLOG_105)
	v2905 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v2905))) = v2906
	*(*int64)(unsafe.Add(mBase, uint32(v2905)+8)) = v2903
	v2910 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	*(*int32)(unsafe.Add(mBase, uint32(v2910)+4)) = int32(0)
	F_MultiXactSetNextMXact(m, v2902, v2901)
	mBase = m.M
	v2914 = m.ExcPending
	if v2914 != 0 {
		goto L32
	} else {
		goto L681
	}
L569:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L32
	} else {
		goto L678
	}
L570:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L32
	} else {
		goto L675
	}
L571:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L32
	} else {
		goto L672
	}
L572:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if v2471 == v2474 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v2476 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if v2476 != int64(0) {
		goto L576
	} else {
		goto L577
	}
L574:
	;
	goto L575
L575:
	;
	v2796 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2798 = F_tliSwitchPoint(m, v2474, v2796, int32(0))
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L32
	} else {
		goto L664
	}
L576:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2483 = F_tliOfPointInHistory(m, v2476-int64(1), v2482)
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L32
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2489 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L32
	} else {
		goto L581
	}
L579:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	if v2483 != v2485 {
		goto L571
	} else {
		goto L580
	}
L580:
	;
	goto L578
L581:
	;
	if v2489 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+196)) = uint32(v2361)
	v2493 = int64(base.Ui64(v2361) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+192)) = uint32(v2493)
	if base.Ui32(v2343) < base.Ui32(int32(16)) {
		goto L585
	} else {
		goto L586
	}
L583:
	;
	goto L584
L584:
	;
	v2513 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L32
	} else {
		goto L590
	}
L585:
	;
	v2499 = int32(_a_F_StartupXLOG_106)
	goto L587
L586:
	;
	v2499 = int32(_a_F_StartupXLOG_107)
	goto L587
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+200)) = v2499
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_108), v706+int32(192))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L32
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(892), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L32
	} else {
		goto L589
	}
L589:
	;
	goto L584
L590:
	;
	if v2513 != 0 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+184)) = v2345
	*(*int64)(unsafe.Add(mBase, uint32(v706)+176)) = v2362
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_109), v706+int32(176))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L32
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v2529 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L32
	} else {
		goto L596
	}
L594:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(896), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L32
	} else {
		goto L595
	}
L595:
	;
	goto L593
L596:
	;
	if v2529 != 0 {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+164)) = v2346
	*(*int32)(unsafe.Add(mBase, uint32(v706)+160)) = v2347
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_110), v706+int32(160))
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L32
	} else {
		goto L600
	}
L598:
	;
	goto L599
L599:
	;
	v2545 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L32
	} else {
		goto L602
	}
L600:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(899), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L32
	} else {
		goto L601
	}
L601:
	;
	goto L599
L602:
	;
	if v2545 != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+148)) = v2349
	*(*int32)(unsafe.Add(mBase, uint32(v706)+144)) = v2348
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_111), v706+int32(144))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L32
	} else {
		goto L606
	}
L604:
	;
	goto L605
L605:
	;
	v2561 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L32
	} else {
		goto L608
	}
L606:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(902), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L32
	} else {
		goto L607
	}
L607:
	;
	goto L605
L608:
	;
	if v2561 != 0 {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+132)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v706)+128)) = v2350
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_112), v706+int32(128))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L32
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v2577 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L32
	} else {
		goto L614
	}
L612:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(905), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L32
	} else {
		goto L613
	}
L613:
	;
	goto L611
L614:
	;
	if v2577 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+116)) = v2352
	*(*int32)(unsafe.Add(mBase, uint32(v706)+112)) = v2353
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_113), v706+int32(112))
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L32
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	if base.Ui32(base.I32_wrap_i64(v2362)) <= base.Ui32(int32(2)) {
		goto L570
	} else {
		goto L620
	}
L618:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(909), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L32
	} else {
		goto L619
	}
L619:
	;
	goto L617
L620:
	;
	v2595 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v2595) < base.Ui64(v2361) {
		goto L569
	} else {
		goto L621
	}
L621:
	;
	if base.Ui64(v2361) < base.Ui64(v2595) {
		goto L628
	} else {
		goto L629
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49])) = v2771
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50])) = v2774
	v2780 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51])) = v2780
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v2780
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4095)))) = uint8(base.B2i32(base.Ui32(v2343) < base.Ui32(int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4093)))) = uint8(base.B2i32(v1055 != int32(0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4094)))) = uint8(v2340)
	m.G0 = v706 + int32(2176)
	goto L568
L623:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	v2770 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	v2771 = v2769
	v2774 = v2770
	goto L622
L624:
	;
	v2755 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v2755
	v2758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v2758)
	v2761 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2761
	if v2751&int32(1) != 0 {
		goto L623
	} else {
		goto L663
	}
L625:
	;
	if v2633&int32(1) != 0 {
		v2676 = int32(5)
		goto L638
	} else {
		goto L639
	}
L626:
	;
	v2628 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	v2630 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v2630 != int32(1) {
		v2751 = v2628
		goto L624
	} else {
		goto L637
	}
L627:
	;
	v2623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v2623)
	v2626 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	v2633 = v2626
	goto L625
L628:
	;
	if base.Ui32(int32(15)) < base.Ui32(v2343) {
		goto L627
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	if v2613 != int32(1) {
		goto L627
	} else {
		goto L635
	}
L631:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L32
	} else {
		goto L632
	}
L632:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_114), int32(0))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L32
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(928), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L32
	} else {
		goto L634
	}
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	v2617 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2617&int32(1) == int32(0) {
		goto L626
	} else {
		goto L636
	}
L636:
	;
	goto L627
L637:
	;
	v2633 = v2628
	goto L625
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+16)) = v2676
	v2679 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+48)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v667)+40)) = v2361
	*(*int64)(unsafe.Add(mBase, uint32(v667)+32)) = v2679
	v2683 = *(*int64)(unsafe.Add(mBase, uint32(v706)+896))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+52)) = v2683
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(v706)+904))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+60)) = v2685
	*(*int32)(unsafe.Add(mBase, uint32(v667)+96)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v667)+92)) = v2350
	*(*int32)(unsafe.Add(mBase, uint32(v667)+88)) = v2349
	*(*int32)(unsafe.Add(mBase, uint32(v667)+84)) = v2348
	*(*int32)(unsafe.Add(mBase, uint32(v667)+80)) = v2346
	*(*int32)(unsafe.Add(mBase, uint32(v667)+76)) = v2347
	*(*int32)(unsafe.Add(mBase, uint32(v667)+72)) = v2345
	*(*int64)(unsafe.Add(mBase, uint32(v667)+64)) = v2362
	v2695 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1088))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+100)) = v2695
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+108)) = v2697
	*(*int64)(unsafe.Add(mBase, uint32(v667)+120)) = v2364
	*(*int32)(unsafe.Add(mBase, uint32(v667)+116)) = v2352
	*(*int32)(unsafe.Add(mBase, uint32(v667)+112)) = v2353
	v2703 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	if v2703 != int32(1) {
		goto L651
	} else {
		goto L652
	}
L639:
	;
	v2639 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L32
	} else {
		goto L640
	}
L640:
	;
	if v2639 != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_115), int32(0))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L32
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	v2650 = int32(4)
	v2652 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	if base.Ui32(v2652) <= base.Ui32(v2653) {
		v2676 = v2650
		goto L638
	} else {
		goto L646
	}
L644:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(958), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L32
	} else {
		goto L645
	}
L645:
	;
	goto L643
L646:
	;
	v2657 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L32
	} else {
		goto L647
	}
L647:
	;
	if v2657 == int32(0) {
		v2676 = v2650
		goto L638
	} else {
		goto L648
	}
L648:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+96)) = v2661
	v2664 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+100)) = v2664
	F_errmsg(m, int32(_a_F_StartupXLOG_116), v706+int32(96))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L32
	} else {
		goto L649
	}
L649:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(964), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L32
	} else {
		goto L650
	}
L650:
	;
	v2676 = v2650
	goto L638
L651:
	;
	if v1055 == int32(0) {
		v2751 = v2703
		goto L624
	} else {
		goto L654
	}
L652:
	;
	v2706 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if base.Ui64(v2361) <= base.Ui64(v2706) {
		goto L651
	} else {
		goto L653
	}
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+144)) = v2344
	*(*int64)(unsafe.Add(mBase, uint32(v667)+136)) = v2361
	goto L651
L654:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v667)+152)) = v2361
	v2714 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])))
	*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)) = uint8(v2714)
	if v2360 == int32(0) {
		v2751 = v2703
		goto L624
	} else {
		goto L655
	}
L655:
	;
	switch v708 - int32(2) {
	case 0, 3:
		goto L656
	default:
		goto L657
	}
L656:
	;
	v2737 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+160)) = v2737
	v2741 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v2741
	v2744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v2744)
	v2747 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2747
	if v2703 != 0 {
		goto L623
	} else {
		goto L662
	}
L657:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		goto L32
	} else {
		goto L658
	}
L658:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_117), int32(0))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L32
	} else {
		goto L659
	}
L659:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_118), int32(0))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L32
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1006), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L32
	} else {
		goto L661
	}
L661:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L662:
	;
	v2771 = int32(0)
	v2774 = int64(0)
	goto L622
L663:
	;
	v2771 = int32(0)
	v2774 = int64(0)
	goto L622
L664:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L32
	} else {
		goto L665
	}
L665:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+256)) = v2805
	F_errmsg(m, int32(_a_F_StartupXLOG_119), v706+int32(256))
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L32
	} else {
		goto L666
	}
L666:
	;
	v2813 = int64(base.Ui64(v2798) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+240)) = uint32(v2813)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+244)) = uint32(v2798)
	if v1055 != 0 {
		goto L667
	} else {
		goto L668
	}
L667:
	;
	v2818 = int32(_a_F_StartupXLOG_59)
	goto L669
L668:
	;
	v2818 = int32(_a_F_StartupXLOG_120)
	goto L669
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+224)) = v2818
	v2821 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+232)) = uint32(v2821)
	v2824 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+236)) = v2824
	v2827 = int64(base.Ui64(v2821) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+228)) = uint32(v2827)
	F_errdetail(m, int32(_a_F_StartupXLOG_121), v706+int32(224))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L32
	} else {
		goto L670
	}
L670:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(873), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L32
	} else {
		goto L671
	}
L671:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L672:
	;
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	v2844 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	v2846 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+208)) = v2846
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+216)) = uint32(v2844)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+220)) = v2843
	v2851 = int64(base.Ui64(v2844) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+212)) = uint32(v2851)
	F_errmsg(m, int32(_a_F_StartupXLOG_122), v706+int32(208))
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L32
	} else {
		goto L673
	}
L673:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(887), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L32
	} else {
		goto L674
	}
L674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L675:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_123), int32(0))
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L32
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(912), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L32
	} else {
		goto L677
	}
L677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L678:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_124), int32(0))
	mBase = m.M
	v2883 = m.ExcPending
	if v2883 != 0 {
		goto L32
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(917), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		goto L32
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	F_AdvanceOldestClogXid(m, v2900)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L32
	} else {
		goto L682
	}
L682:
	;
	F_SetTransactionIdLimit(m, v2900, v2899)
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L32
	} else {
		goto L683
	}
L683:
	;
	F_SetMultiXactIdLimit(m, v2898, v2897, int32(1))
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		goto L32
	} else {
		goto L684
	}
L684:
	;
	F_SetCommitTsLimit(m, v2896, v2895)
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L32
	} else {
		goto L685
	}
L685:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v2925)+208)) = v2903
	v2927 = m.G0
	v2929 = v2927 - int32(1088)
	m.G0 = v2929
	*(*int32)(unsafe.Add(mBase, uint32(v2929)+16)) = int32(_a_F_StartupXLOG_125)
	v2934 = v2929 + int32(32)
	v2939 = F_pg_snprintf(m, v2934, int32(1050), int32(_a_F_StartupXLOG_126), v2929+int32(16))
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L32
	} else {
		goto L686
	}
L686:
	;
	F_unlink_initfile(m, v2934, int32(15))
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L32
	} else {
		goto L687
	}
L687:
	;
	F_RelationCacheInitFileRemoveInDir(m, int32(_a_F_StartupXLOG_127))
	mBase = m.M
	v2946 = m.ExcPending
	if v2946 != 0 {
		goto L32
	} else {
		goto L688
	}
L688:
	;
	v2948 = F_AllocateDir(m, int32(_a_F_StartupXLOG_38))
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L32
	} else {
		goto L689
	}
L689:
	;
	v2952 = F_ReadDirExtended(m, v2948, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L32
	} else {
		goto L690
	}
L690:
	;
	if v2952 != 0 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v2954 = v2952
	goto L694
L692:
	;
	goto L693
L693:
	;
	F_FreeDir(m, v2948)
	mBase = m.M
	v3129 = m.ExcPending
	if v3129 != 0 {
		goto L32
	} else {
		goto L722
	}
L694:
	;
	v2989 = v2954 + int32(19)
	v2990 = int32(_a_F_StartupXLOG_128)
	v2994 = m.G0
	v2996 = v2994 - int32(32)
	v2997 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2996)+24)) = v2997
	*(*int64)(unsafe.Add(mBase, uint32(v2996)+16)) = v2997
	*(*int64)(unsafe.Add(mBase, uint32(v2996)+8)) = v2997
	*(*int64)(unsafe.Add(mBase, uint32(v2996))) = v2997
	v3005 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[56])))
	if v3005 == int32(0) {
		goto L697
	} else {
		goto L698
	}
L695:
	;
	goto L693
L696:
	;
	v3074 = F_strlen(m, v2989)
	mBase = m.M
	if v3073 == v3074 {
		goto L715
	} else {
		goto L716
	}
L697:
	;
	v3073 = int32(0)
	goto L696
L698:
	;
	goto L699
L699:
	;
	v3009 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[57])))
	if v3009 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3013 = v2989
	goto L703
L701:
	;
	goto L702
L702:
	;
	v3023 = v2990
	v3024 = v3005
	goto L706
L703:
	;
	v3019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3013))))
	if v3019 == v3005 {
		v3013 = v3013 + int32(1)
		goto L703
	} else {
		goto L705
	}
L704:
	;
	v3073 = v3013 - v2989
	goto L696
L705:
	;
	goto L704
L706:
	;
	v3031 = v2996 + int32(base.Ui32(v3024)>>(uint(int32(3))%32))&int32(28)
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v3031)))
	v3033 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3031))) = v3032 | v3033<<(uint(v3024)%32)
	v3037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023)+1)))
	if v3037 != 0 {
		v3023 = v3023 + v3033
		v3024 = v3037
		goto L706
	} else {
		goto L708
	}
L707:
	;
	v3040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2989))))
	if v3040 == int32(0) {
		v3063 = v2989
		goto L709
	} else {
		goto L710
	}
L708:
	;
	goto L707
L709:
	;
	v3073 = v3063 - v2989
	goto L696
L710:
	;
	v3044 = v2989
	v3045 = v3040
	goto L711
L711:
	;
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(v2996+int32(base.Ui32(v3045)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3053)>>(uint(v3045)%32))&int32(1) == int32(0) {
		v3063 = v3044
		goto L709
	} else {
		goto L713
	}
L712:
	;
	v3063 = v3061
	goto L709
L713:
	;
	v3059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3044)+1)))
	v3061 = v3044 + int32(1)
	if v3059 != 0 {
		v3044 = v3061
		v3045 = v3059
		goto L711
	} else {
		goto L714
	}
L714:
	;
	goto L712
L715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2929)+8)) = int32(_a_F_StartupXLOG_129)
	*(*int32)(unsafe.Add(mBase, uint32(v2929)+4)) = v2989
	*(*int32)(unsafe.Add(mBase, uint32(v2929))) = int32(_a_F_StartupXLOG_38)
	v3082 = v2929 + int32(32)
	v3085 = F_pg_snprintf(m, v3082, int32(1050), int32(_a_F_StartupXLOG_130), v2929)
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L32
	} else {
		goto L718
	}
L716:
	;
	goto L717
L717:
	;
	v3092 = F_ReadDirExtended(m, v2948, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L32
	} else {
		goto L720
	}
L718:
	;
	F_RelationCacheInitFileRemoveInDir(m, v3082)
	mBase = m.M
	v3088 = m.ExcPending
	if v3088 != 0 {
		goto L32
	} else {
		goto L719
	}
L719:
	;
	goto L717
L720:
	;
	if v3092 != 0 {
		v2954 = v3092
		goto L694
	} else {
		goto L721
	}
L721:
	;
	goto L695
L722:
	;
	m.G0 = v2929 + int32(1088)
	v3133 = m.G0
	v3135 = v3133 - int32(3696)
	m.G0 = v3135
	v3139 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L32
	} else {
		goto L723
	}
L723:
	;
	if v3139 != 0 {
		goto L724
	} else {
		goto L725
	}
L724:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_131), int32(0))
	mBase = m.M
	v3144 = m.ExcPending
	if v3144 != 0 {
		goto L32
	} else {
		goto L727
	}
L725:
	;
	goto L726
L726:
	;
	v3151 = F_AllocateDir(m, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L32
	} else {
		goto L743
	}
L727:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2202), int32(_a_F_StartupXLOG_134))
	mBase = m.M
	v3149 = m.ExcPending
	if v3149 != 0 {
		goto L32
	} else {
		goto L728
	}
L728:
	;
	goto L726
L729:
	;
	v4002 = F_AllocateDir(m, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L32
	} else {
		goto L919
	}
L730:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L32
	} else {
		goto L915
	}
L731:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3931 = m.ExcPending
	if v3931 != 0 {
		goto L32
	} else {
		goto L910
	}
L732:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L32
	} else {
		goto L905
	}
L733:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3889 = m.ExcPending
	if v3889 != 0 {
		goto L32
	} else {
		goto L902
	}
L734:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L32
	} else {
		goto L898
	}
L735:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L32
	} else {
		goto L895
	}
L736:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3834 = m.ExcPending
	if v3834 != 0 {
		goto L32
	} else {
		goto L891
	}
L737:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3813 = m.ExcPending
	if v3813 != 0 {
		goto L32
	} else {
		goto L887
	}
L738:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3790 = m.ExcPending
	if v3790 != 0 {
		goto L32
	} else {
		goto L883
	}
L739:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
		goto L32
	} else {
		goto L880
	}
L740:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3756 = m.ExcPending
	if v3756 != 0 {
		goto L32
	} else {
		goto L876
	}
L741:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L32
	} else {
		goto L872
	}
L742:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3718 = m.ExcPending
	if v3718 != 0 {
		goto L32
	} else {
		goto L868
	}
L743:
	;
	v3154 = F_ReadDir(m, v3151, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L32
	} else {
		goto L744
	}
L744:
	;
	if v3154 != 0 {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	v3157 = v3135 + int32(3512)
	v3166 = v3154
	goto L748
L746:
	;
	goto L747
L747:
	;
	F_FreeDir(m, v3151)
	mBase = m.M
	v3702 = m.ExcPending
	if v3702 != 0 {
		goto L32
	} else {
		goto L862
	}
L748:
	;
	v3194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3166)+19)))
	if v3194 != int32(46) {
		goto L751
	} else {
		goto L752
	}
L749:
	;
	goto L747
L750:
	;
	v3665 = F_ReadDir(m, v3151, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3666 = m.ExcPending
	if v3666 != 0 {
		goto L32
	} else {
		goto L860
	}
L751:
	;
	v3207 = v3166 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+340)) = v3207
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+336)) = int32(_a_F_StartupXLOG_132)
	v3212 = v3135 + int32(352)
	v3217 = F_pg_snprintf(m, v3212, int32(1036), int32(_a_F_StartupXLOG_135), v3135+int32(336))
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L32
	} else {
		goto L756
	}
L752:
	;
	v3197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3166)+20)))
	if v3197 == int32(0) {
		goto L750
	} else {
		goto L753
	}
L753:
	;
	v3200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3166)+20)))
	if v3200 != int32(46) {
		goto L751
	} else {
		goto L754
	}
L754:
	;
	v3203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3166)+21)))
	if v3203 == int32(0) {
		goto L750
	} else {
		goto L755
	}
L755:
	;
	goto L751
L756:
	;
	v3221 = F_get_dirent_type(m, v3212, v3166, int32(0), int32(14))
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L32
	} else {
		goto L758
	}
L757:
	;
	v3223 = F_strlen(m, v3207)
	mBase = m.M
	v3225 = F_strlen(m, int32(_a_F_StartupXLOG_136))
	mBase = m.M
	if base.Ui32(v3225) <= base.Ui32(v3223) {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	switch v3221 {
	case 0, 3:
		goto L757
	default:
		goto L750
	}
L759:
	;
	v3228 = v3207 + (v3223 - v3225)
	v3229 = int32(_a_F_StartupXLOG_136)
	v3232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3228))))
	v3235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[58])))
	if base.B2i32(v3232 == int32(0))|base.B2i32(v3232 != v3235) != 0 {
		v3253 = v3232
		v3254 = v3235
		goto L763
	} else {
		goto L764
	}
L760:
	;
	v3257 = int32(1)
	goto L761
L761:
	;
	if v3257 == int32(0) {
		goto L769
	} else {
		goto L770
	}
L762:
	;
	v3257 = v3253 - v3254
	goto L761
L763:
	;
	goto L762
L764:
	;
	v3238 = v3228
	v3239 = v3229
	goto L765
L765:
	;
	v3242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3239)+1)))
	v3243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3238)+1)))
	if v3243 == int32(0) {
		v3253 = v3243
		v3254 = v3242
		goto L763
	} else {
		goto L767
	}
L766:
	;
	v3253 = v3243
	v3254 = v3242
	goto L763
L767:
	;
	v3246 = int32(1)
	if v3243 == v3242 {
		v3238 = v3238 + v3246
		v3239 = v3239 + v3246
		goto L765
	} else {
		goto L768
	}
L768:
	;
	goto L766
L769:
	;
	v3261 = v3135 + int32(352)
	v3262 = F_rmtree(m, v3261)
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L32
	} else {
		goto L772
	}
L770:
	;
	goto L771
L771:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+320)) = int32(_a_F_StartupXLOG_132)
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+324)) = v3207
	v3289 = v3135 + int32(2448)
	v3293 = F_pg_sprintf(m, v3289, int32(_a_F_StartupXLOG_135), v3135+int32(320))
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L32
	} else {
		goto L781
	}
L772:
	;
	if v3262 == int32(0) {
		goto L773
	} else {
		goto L774
	}
L773:
	;
	v3268 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3269 = m.ExcPending
	if v3269 != 0 {
		goto L32
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	F_fsync_fname(m, int32(_a_F_StartupXLOG_132), int32(1))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L32
	} else {
		goto L780
	}
L776:
	;
	if v3268 == int32(0) {
		goto L750
	} else {
		goto L777
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135))) = v3261
	F_errmsg(m, int32(_a_F_StartupXLOG_137), v3135)
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L32
	} else {
		goto L778
	}
L778:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2229), int32(_a_F_StartupXLOG_134))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L32
	} else {
		goto L779
	}
L779:
	;
	goto L750
L780:
	;
	goto L750
L781:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+304)) = v3289
	v3297 = v3135 + int32(1392)
	v3301 = F_pg_sprintf(m, v3297, int32(_a_F_StartupXLOG_138), v3135+int32(304))
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L32
	} else {
		goto L782
	}
L782:
	;
	v3303 = F_unlink(m, v3297)
	mBase = m.M
	if v3303 < int32(0) {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v3307 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3307 != int32(44) {
		goto L742
	} else {
		goto L786
	}
L784:
	;
	goto L785
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+272)) = v3135 + int32(2448)
	v3314 = v3135 + int32(1392)
	v3318 = F_pg_sprintf(m, v3314, int32(_a_F_StartupXLOG_139), v3135+int32(272))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L32
	} else {
		goto L787
	}
L786:
	;
	goto L785
L787:
	;
	v3322 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L32
	} else {
		goto L788
	}
L788:
	;
	if v3322 != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+256)) = v3314
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_140), v3135+int32(256))
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L32
	} else {
		goto L792
	}
L790:
	;
	goto L791
L791:
	;
	v3336 = v3135 + int32(1392)
	v3338 = F_OpenTransientFile(m, v3336, int32(2))
	mBase = m.M
	v3339 = m.ExcPending
	if v3339 != 0 {
		goto L32
	} else {
		goto L794
	}
L792:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2506), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L32
	} else {
		goto L793
	}
L793:
	;
	goto L791
L794:
	;
	if v3338 < int32(0) {
		goto L741
	} else {
		goto L795
	}
L795:
	;
	v3343 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3343))) = int32(167772207)
	v3348 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v3348 != int32(1) {
		v3362 = int32(0)
		goto L797
	} else {
		goto L798
	}
L796:
	;
	if v3362 != 0 {
		goto L740
	} else {
		goto L803
	}
L797:
	;
	goto L796
L798:
	;
	goto L799
L799:
	;
	v3353 = F_fsync(m, v3338)
	mBase = m.M
	if v3353 != int32(-1) {
		v3362 = v3353
		goto L797
	} else {
		goto L801
	}
L800:
	;
	v3362 = int32(-1)
	goto L797
L801:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3357 == int32(27) {
		goto L799
	} else {
		goto L802
	}
L802:
	;
	goto L800
L803:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3364))) = int32(0)
	v3367 = int32(_a_F_StartupXLOG_142)
	v3369 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	v3370 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v3369 + v3370
	F_fsync_fname(m, v3135+int32(2448), v3370)
	mBase = m.M
	v3377 = m.ExcPending
	if v3377 != 0 {
		goto L32
	} else {
		goto L804
	}
L804:
	;
	v3378 = int32(_a_F_StartupXLOG_142)
	v3380 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v3380 - int32(1)
	v3384 = int32(_a_F_StartupXLOG_143)
	v3385 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3385))) = int32(167772206)
	v3390 = int32(16)
	v3391 = F_read(m, v3338, v3135+int32(3496), v3390)
	mBase = m.M
	v3393 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3393))) = int32(0)
	if v3391 != v3390 {
		goto L805
	} else {
		goto L806
	}
L805:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3401 = m.ExcPending
	if v3401 != 0 {
		goto L32
	} else {
		goto L808
	}
L806:
	;
	goto L807
L807:
	;
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3496))
	if v3421 != int32(17112225) {
		goto L738
	} else {
		goto L813
	}
L808:
	;
	if v3391 < int32(0) {
		goto L739
	} else {
		goto L809
	}
L809:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3406 = m.ExcPending
	if v3406 != 0 {
		goto L32
	} else {
		goto L810
	}
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+232)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+228)) = v3391
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+224)) = v3336
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v3135+int32(224))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L32
	} else {
		goto L811
	}
L811:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2552), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L32
	} else {
		goto L812
	}
L812:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L813:
	;
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3504))
	if v3424 != int32(5) {
		goto L737
	} else {
		goto L814
	}
L814:
	;
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3508))
	if v3427 != int32(184) {
		goto L736
	} else {
		goto L815
	}
L815:
	;
	v3430 = int32(_a_F_StartupXLOG_143)
	v3431 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3431))) = int32(167772206)
	v3435 = F_read(m, v3338, v3157, int32(184))
	mBase = m.M
	v3437 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3437))) = int32(0)
	v3440 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3508))
	if v3440 != v3435 {
		goto L816
	} else {
		goto L817
	}
L816:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3445 = m.ExcPending
	if v3445 != 0 {
		goto L32
	} else {
		goto L819
	}
L817:
	;
	goto L818
L818:
	;
	v3466 = F_CloseTransientFile(m, v3338)
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L32
	} else {
		goto L824
	}
L819:
	;
	if v3435 < int32(0) {
		goto L735
	} else {
		goto L820
	}
L820:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3450 = m.ExcPending
	if v3450 != 0 {
		goto L32
	} else {
		goto L821
	}
L821:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+152)) = v3440
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+148)) = v3435
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+144)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v3135+int32(144))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L32
	} else {
		goto L822
	}
L822:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2592), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L32
	} else {
		goto L823
	}
L823:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L824:
	;
	if v3466 != 0 {
		goto L734
	} else {
		goto L825
	}
L825:
	;
	v3468 = int32(-1)
	v3470 = m.Env.Pgmem_crc32c(m, v3468, v3135+int32(3504), int32(192))
	mBase = m.M
	v3472 = v3470 ^ v3468
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3500))
	if v3472 != v3473 {
		goto L733
	} else {
		goto L826
	}
L826:
	;
	v3475 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3580))
	if v3475 != 0 {
		goto L827
	} else {
		goto L828
	}
L827:
	;
	v3477 = v3135 + int32(2448)
	v3478 = F_rmtree(m, v3477)
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L32
	} else {
		goto L831
	}
L828:
	;
	goto L829
L829:
	;
	v3502 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3576))
	if v3503 != 0 {
		goto L839
	} else {
		goto L840
	}
L830:
	;
	F_fsync_fname(m, int32(_a_F_StartupXLOG_132), int32(1))
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L32
	} else {
		goto L837
	}
L831:
	;
	if v3478 != 0 {
		goto L830
	} else {
		goto L832
	}
L832:
	;
	v3482 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L32
	} else {
		goto L833
	}
L833:
	;
	if v3482 == int32(0) {
		goto L830
	} else {
		goto L834
	}
L834:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+80)) = v3477
	F_errmsg(m, int32(_a_F_StartupXLOG_137), v3135+int32(80))
	mBase = m.M
	v3491 = m.ExcPending
	if v3491 != 0 {
		goto L32
	} else {
		goto L835
	}
L835:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2622), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L32
	} else {
		goto L836
	}
L836:
	;
	goto L830
L837:
	;
	goto L750
L838:
	;
	v3539 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if v3539 <= int32(0) {
		goto L730
	} else {
		goto L851
	}
L839:
	;
	if v3502 <= int32(1) {
		goto L732
	} else {
		goto L842
	}
L840:
	;
	goto L841
L841:
	;
	if v3502 <= int32(0) {
		goto L731
	} else {
		goto L850
	}
L842:
	;
	v3507 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v3507 != int32(1) {
		goto L838
	} else {
		goto L843
	}
L843:
	;
	v3511 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v3511&int32(1) != 0 {
		goto L838
	} else {
		goto L844
	}
L844:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L32
	} else {
		goto L845
	}
L845:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L32
	} else {
		goto L846
	}
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+64)) = v3157
	F_errmsg(m, int32(_a_F_StartupXLOG_145), v3135-int32(-64))
	mBase = m.M
	v3526 = m.ExcPending
	if v3526 != 0 {
		goto L32
	} else {
		goto L847
	}
L847:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_146), int32(0))
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L32
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2661), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		goto L32
	} else {
		goto L849
	}
L849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L850:
	;
	goto L838
L851:
	;
	v3544 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v3551 = int32(0)
	goto L852
L852:
	;
	v3581 = v3544 + v3551*int32(288)
	v3582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+4)))
	if v3582 != 0 {
		goto L854
	} else {
		goto L855
	}
L853:
	;
	base.MemoryCopy(m, v3581+int32(24), v3157, int32(184))
	v3590 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3584))
	*(*int32)(unsafe.Add(mBase, uint32(v3581)+16)) = v3590
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3588))
	*(*int32)(unsafe.Add(mBase, uint32(v3581)+20)) = v3592
	v3594 = *(*int64)(unsafe.Add(mBase, uint32(v3135)+3608))
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+264)) = v3594
	v3596 = *(*int64)(unsafe.Add(mBase, uint32(v3135)+3592))
	v3597 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+236)) = v3597
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+280)) = v3596
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+244)) = v3597
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+252)) = v3597
	v3604 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3581)+260)) = v3604
	*(*int32)(unsafe.Add(mBase, uint32(v3581)+8)) = v3604
	v3608 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3581)+4)) = uint8(v3608)
	v3613 = m.G0
	v3614 = int32(16)
	v3615 = v3613 - v3614
	m.G0 = v3615
	F_gettimeofday(m, v3615)
	mBase = m.M
	v3618 = *(*int64)(unsafe.Add(mBase, uint32(v3615)))
	v3619 = int64(*(*int32)(unsafe.Add(mBase, uint32(v3615)+8)))
	m.G0 = v3615 + v3614
	goto L858
L854:
	;
	v3584 = v3551 + int32(1)
	if v3539 != v3584 {
		v3551 = v3584
		goto L852
	} else {
		goto L857
	}
L855:
	;
	goto L856
L856:
	;
	goto L853
L857:
	;
	goto L730
L858:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v3581)+112))
	if v3628 != 0 {
		goto L750
	} else {
		goto L859
	}
L859:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3581)+272)) = v3619 + v3618*int64(1000000) - int64(946684800000000)
	goto L750
L860:
	;
	if v3665 != 0 {
		v3166 = v3665
		goto L748
	} else {
		goto L861
	}
L861:
	;
	goto L749
L862:
	;
	v3704 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v3704 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L32
	} else {
		goto L866
	}
L864:
	;
	goto L865
L865:
	;
	m.G0 = v3135 + int32(3696)
	goto L729
L866:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L32
	} else {
		goto L867
	}
L867:
	;
	goto L865
L868:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L32
	} else {
		goto L869
	}
L869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+288)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v3135+int32(288))
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L32
	} else {
		goto L870
	}
L870:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2502), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3733 = m.ExcPending
	if v3733 != 0 {
		goto L32
	} else {
		goto L871
	}
L871:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L872:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L32
	} else {
		goto L873
	}
L873:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+16)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v3135+int32(16))
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L32
	} else {
		goto L874
	}
L874:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2518), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L32
	} else {
		goto L875
	}
L875:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L876:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3758 = m.ExcPending
	if v3758 != 0 {
		goto L32
	} else {
		goto L877
	}
L877:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+240)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v3135+int32(240))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L32
	} else {
		goto L878
	}
L878:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2529), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3771 = m.ExcPending
	if v3771 != 0 {
		goto L32
	} else {
		goto L879
	}
L879:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L880:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+208)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v3135+int32(208))
	mBase = m.M
	v3781 = m.ExcPending
	if v3781 != 0 {
		goto L32
	} else {
		goto L881
	}
L881:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2546), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L32
	} else {
		goto L882
	}
L882:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L883:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3793 = m.ExcPending
	if v3793 != 0 {
		goto L32
	} else {
		goto L884
	}
L884:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+200)) = int32(17112225)
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+196)) = v3421
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+192)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_150), v3135+int32(192))
	mBase = m.M
	v3804 = m.ExcPending
	if v3804 != 0 {
		goto L32
	} else {
		goto L885
	}
L885:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2560), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3809 = m.ExcPending
	if v3809 != 0 {
		goto L32
	} else {
		goto L886
	}
L886:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L887:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L32
	} else {
		goto L888
	}
L888:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+180)) = v3424
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+176)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_151), v3135+int32(176))
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		goto L32
	} else {
		goto L889
	}
L889:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2567), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3830 = m.ExcPending
	if v3830 != 0 {
		goto L32
	} else {
		goto L890
	}
L890:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L891:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3837 = m.ExcPending
	if v3837 != 0 {
		goto L32
	} else {
		goto L892
	}
L892:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+164)) = v3427
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+160)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v3135+int32(160))
	mBase = m.M
	v3846 = m.ExcPending
	if v3846 != 0 {
		goto L32
	} else {
		goto L893
	}
L893:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2574), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L32
	} else {
		goto L894
	}
L894:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L895:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+128)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v3135+int32(128))
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L32
	} else {
		goto L896
	}
L896:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2587), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3866 = m.ExcPending
	if v3866 != 0 {
		goto L32
	} else {
		goto L897
	}
L897:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L898:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L32
	} else {
		goto L899
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+112)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v3135+int32(112))
	mBase = m.M
	v3880 = m.ExcPending
	if v3880 != 0 {
		goto L32
	} else {
		goto L900
	}
L900:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2598), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L32
	} else {
		goto L901
	}
L901:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L902:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+100)) = v3472
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+3500))
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+104)) = v3891
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+96)) = v3135 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_154), v3135+int32(96))
	mBase = m.M
	v3900 = m.ExcPending
	if v3900 != 0 {
		goto L32
	} else {
		goto L903
	}
L903:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2610), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
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
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L32
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+48)) = v3157
	F_errmsg(m, int32(_a_F_StartupXLOG_155), v3135+int32(48))
	mBase = m.M
	v3918 = m.ExcPending
	if v3918 != 0 {
		goto L32
	} else {
		goto L907
	}
L907:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_156), int32(0))
	mBase = m.M
	v3922 = m.ExcPending
	if v3922 != 0 {
		goto L32
	} else {
		goto L908
	}
L908:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2647), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3927 = m.ExcPending
	if v3927 != 0 {
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
	F_errcode(m, int32(325))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L32
	} else {
		goto L911
	}
L911:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+32)) = v3157
	F_errmsg(m, int32(_a_F_StartupXLOG_157), v3135+int32(32))
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		goto L32
	} else {
		goto L912
	}
L912:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_158), int32(0))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L32
	} else {
		goto L913
	}
L913:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2668), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3949 = m.ExcPending
	if v3949 != 0 {
		goto L32
	} else {
		goto L914
	}
L914:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L915:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_159), int32(0))
	mBase = m.M
	v3991 = m.ExcPending
	if v3991 != 0 {
		goto L32
	} else {
		goto L916
	}
L916:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_160), int32(0))
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L32
	} else {
		goto L917
	}
L917:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2716), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L32
	} else {
		goto L918
	}
L918:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L919:
	;
	v4005 = F_ReadDir(m, v4002, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L32
	} else {
		goto L920
	}
L920:
	;
	if v4005 != 0 {
		goto L921
	} else {
		goto L922
	}
L921:
	;
	v4010 = v4005
	goto L924
L922:
	;
	goto L923
L923:
	;
	F_FreeDir(m, v4002)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L32
	} else {
		goto L937
	}
L924:
	;
	v4041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4010)+19)))
	if v4041 != int32(46) {
		goto L927
	} else {
		goto L928
	}
L925:
	;
	goto L923
L926:
	;
	v4064 = F_ReadDir(m, v4002, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L32
	} else {
		goto L935
	}
L927:
	;
	v4054 = v4010 + int32(19)
	v4056 = F_ReplicationSlotValidateName(m, v4054, int32(13))
	mBase = m.M
	v4057 = m.ExcPending
	if v4057 != 0 {
		goto L32
	} else {
		goto L932
	}
L928:
	;
	v4044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4010)+20)))
	if v4044 == int32(0) {
		goto L926
	} else {
		goto L929
	}
L929:
	;
	v4047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4010)+20)))
	if v4047 != int32(46) {
		goto L927
	} else {
		goto L930
	}
L930:
	;
	v4050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4010)+21)))
	if v4050 == int32(0) {
		goto L926
	} else {
		goto L931
	}
L931:
	;
	goto L927
L932:
	;
	if v4056 == int32(0) {
		goto L926
	} else {
		goto L933
	}
L933:
	;
	F_ReorderBufferCleanupSerializedTXNs(m, v4054)
	mBase = m.M
	v4061 = m.ExcPending
	if v4061 != 0 {
		goto L32
	} else {
		goto L934
	}
L934:
	;
	goto L926
L935:
	;
	if v4064 != 0 {
		v4010 = v4064
		goto L924
	} else {
		goto L936
	}
L936:
	;
	goto L925
L937:
	;
	v4103 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v4105 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v4106 = *(*int64)(unsafe.Add(mBase, uint32(v4105)+8))
	v4111 = int32(48)
	v4112 = base.AtomicRmwXchg64(m, v4103, v4111, int64(base.Ui64(v4106)>>(uint(int64(15))%64))&int64(131071))
	v4114 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v4115 = *(*int32)(unsafe.Add(mBase, uint32(v4114)+4))
	v4117 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(v4114)))
	v4123 = base.AtomicRmwXchg64(m, v4117, v4111, base.I64_extend_i32_u(int32(base.Ui32(v4118)>>(uint(int32(11))%32))))
	v4125 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v4127 = base.I32_div_u_s(v4115, int32(1636))
	v4130 = base.AtomicRmwXchg64(m, v4125, v4111, base.I64_extend_i32_u(v4127))
	v4132 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4132)+200)))
	if v4133 == int32(1) {
		goto L938
	} else {
		goto L939
	}
L938:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L32
	} else {
		goto L941
	}
L939:
	;
	goto L940
L940:
	;
	v4138 = m.G0
	v4140 = v4138 - int32(176)
	m.G0 = v4140
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+172)) = int32(307747550)
	v4145 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	if v4145 == int32(0) {
		goto L950
	} else {
		goto L951
	}
L941:
	;
	goto L940
L942:
	;
	v4518 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v4520 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4521 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+16))
	if v4521 == int32(1) {
		goto L1020
	} else {
		goto L1021
	}
L943:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		goto L32
	} else {
		goto L1016
	}
L944:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L32
	} else {
		goto L1012
	}
L945:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L32
	} else {
		goto L1008
	}
L946:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4446 = m.ExcPending
	if v4446 != 0 {
		goto L32
	} else {
		goto L1004
	}
L947:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
		goto L32
	} else {
		goto L1000
	}
L948:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L32
	} else {
		goto L997
	}
L949:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L32
	} else {
		goto L994
	}
L950:
	;
	m.G0 = v4140 + int32(176)
	goto L942
L951:
	;
	v4150 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L32
	} else {
		goto L952
	}
L952:
	;
	if v4150 != 0 {
		goto L953
	} else {
		goto L954
	}
L953:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_161), int32(0))
	mBase = m.M
	v4155 = m.ExcPending
	if v4155 != 0 {
		goto L32
	} else {
		goto L956
	}
L954:
	;
	goto L955
L955:
	;
	v4163 = F_OpenTransientFile(m, int32(_a_F_StartupXLOG_162), int32(0))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L32
	} else {
		goto L958
	}
L956:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(745), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		goto L32
	} else {
		goto L957
	}
L957:
	;
	goto L955
L958:
	;
	if v4163 < int32(0) {
		goto L959
	} else {
		goto L960
	}
L959:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4168 == int32(44) {
		goto L950
	} else {
		goto L962
	}
L960:
	;
	goto L961
L961:
	;
	v4189 = int32(4)
	v4190 = F_read(m, v4163, v4140+int32(172), v4189)
	mBase = m.M
	if v4190 != v4189 {
		goto L967
	} else {
		goto L968
	}
L962:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		goto L32
	} else {
		goto L963
	}
L963:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4176 = m.ExcPending
	if v4176 != 0 {
		goto L32
	} else {
		goto L964
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140))) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v4140)
	mBase = m.M
	v4181 = m.ExcPending
	if v4181 != 0 {
		goto L32
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(759), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L32
	} else {
		goto L966
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L32
	} else {
		goto L970
	}
L968:
	;
	goto L969
L969:
	;
	v4221 = m.Env.Pgmem_crc32c(m, int32(-1), v4140+int32(172), int32(4))
	mBase = m.M
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+172))
	if v4222 != int32(307747550) {
		goto L948
	} else {
		goto L975
	}
L970:
	;
	if v4190 < int32(0) {
		goto L949
	} else {
		goto L971
	}
L971:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4201 = m.ExcPending
	if v4201 != 0 {
		goto L32
	} else {
		goto L972
	}
L972:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+136)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+132)) = v4190
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+128)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v4140+int32(128))
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L32
	} else {
		goto L973
	}
L973:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(774), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4216 = m.ExcPending
	if v4216 != 0 {
		goto L32
	} else {
		goto L974
	}
L974:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L975:
	;
	v4228 = F_read(m, v4163, v4140+int32(152), int32(16))
	mBase = m.M
	if v4228 != int32(4) {
		goto L976
	} else {
		goto L977
	}
L976:
	;
	v4232 = int32(0)
	v4234 = v4221
	v4238 = v4228
	goto L979
L977:
	;
	v4318 = v4221
	goto L978
L978:
	;
	v4350 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+152))
	v4352 = v4318 ^ int32(-1)
	if v4350 != v4352 {
		goto L944
	} else {
		goto L991
	}
L979:
	;
	if v4238 < int32(0) {
		goto L947
	} else {
		goto L981
	}
L980:
	;
	v4318 = v4273
	goto L978
L981:
	;
	if v4238 != int32(16) {
		goto L946
	} else {
		goto L982
	}
L982:
	;
	v4273 = m.Env.Pgmem_crc32c(m, v4234, v4140+int32(152), int32(16))
	mBase = m.M
	v4275 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	if v4232 == v4275 {
		goto L945
	} else {
		goto L983
	}
L983:
	;
	v4278 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v4281 = v4278 + v4232*int32(56)
	v4282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4140)+152)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4281))) = uint16(v4282)
	v4284 = *(*int64)(unsafe.Add(mBase, uint32(v4140)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v4281)+8)) = v4284
	v4288 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L32
	} else {
		goto L984
	}
L984:
	;
	if v4288 != 0 {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v4290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4140)+152)))
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+64)) = v4290
	v4292 = *(*int64)(unsafe.Add(mBase, uint32(v4140)+160))
	*(*uint32)(unsafe.Add(mBase, uint32(v4140)+72)) = uint32(v4292)
	v4295 = int64(base.Ui64(v4292) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4140)+68)) = uint32(v4295)
	F_errmsg(m, int32(_a_F_StartupXLOG_165), v4140-int32(-64))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		goto L32
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v4313 = F_read(m, v4163, v4140+int32(152), int32(16))
	mBase = m.M
	if v4313 != int32(4) {
		v4232 = v4232 + int32(1)
		v4234 = v4273
		v4238 = v4313
		goto L979
	} else {
		goto L990
	}
L988:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(831), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		goto L32
	} else {
		goto L989
	}
L989:
	;
	goto L987
L990:
	;
	goto L980
L991:
	;
	v4354 = F_CloseTransientFile(m, v4163)
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L32
	} else {
		goto L992
	}
L992:
	;
	if v4354 != 0 {
		goto L943
	} else {
		goto L993
	}
L993:
	;
	goto L950
L994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+112)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v4140+int32(112))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L32
	} else {
		goto L995
	}
L995:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(769), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4406 = m.ExcPending
	if v4406 != 0 {
		goto L32
	} else {
		goto L996
	}
L996:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+100)) = int32(307747550)
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+96)) = v4413
	F_errmsg(m, int32(_a_F_StartupXLOG_166), v4140+int32(96))
	mBase = m.M
	v4419 = m.ExcPending
	if v4419 != 0 {
		goto L32
	} else {
		goto L998
	}
L998:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(781), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4424 = m.ExcPending
	if v4424 != 0 {
		goto L32
	} else {
		goto L999
	}
L999:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1000:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4430 = m.ExcPending
	if v4430 != 0 {
		goto L32
	} else {
		goto L1001
	}
L1001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+48)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v4140+int32(48))
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		goto L32
	} else {
		goto L1002
	}
L1002:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(805), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v4448 = m.ExcPending
	if v4448 != 0 {
		goto L32
	} else {
		goto L1005
	}
L1005:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+88)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+84)) = v4238
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+80)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v4140+int32(80))
	mBase = m.M
	v4458 = m.ExcPending
	if v4458 != 0 {
		goto L32
	} else {
		goto L1006
	}
L1006:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(813), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4463 = m.ExcPending
	if v4463 != 0 {
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
	F_errcode(m, int32(_a_F_StartupXLOG_167))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L32
	} else {
		goto L1009
	}
L1009:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_168), int32(0))
	mBase = m.M
	v4474 = m.ExcPending
	if v4474 != 0 {
		goto L32
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(821), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4479 = m.ExcPending
	if v4479 != 0 {
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4486 = m.ExcPending
	if v4486 != 0 {
		goto L32
	} else {
		goto L1013
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+36)) = v4350
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+32)) = v4352
	F_errmsg(m, int32(_a_F_StartupXLOG_169), v4140+int32(32))
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L32
	} else {
		goto L1014
	}
L1014:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(840), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L32
	} else {
		goto L1017
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4140)+16)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v4140+int32(16))
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L32
	} else {
		goto L1018
	}
L1018:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(846), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4516 = m.ExcPending
	if v4516 != 0 {
		goto L32
	} else {
		goto L1019
	}
L1019:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1020:
	;
	v4524 = *(*int64)(unsafe.Add(mBase, uint32(v4520)+128))
	v4526 = v4524
	goto L1022
L1021:
	;
	v4526 = int64(1000)
	goto L1022
L1022:
	;
	v4528 = base.AtomicRmwXchg64(m, v4518, int32(240), v4526)
	v4530 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	F_restoreTimeLineHistoryFiles(m, v2893, v4530)
	mBase = m.M
	v4532 = m.ExcPending
	if v4532 != 0 {
		goto L32
	} else {
		goto L1023
	}
L1023:
	;
	v4534 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v4538 = F_LWLockAcquire(m, v4534+int32(2304), int32(0))
	mBase = m.M
	v4539 = m.ExcPending
	if v4539 != 0 {
		goto L32
	} else {
		goto L1024
	}
L1024:
	;
	v4541 = F_AllocateDir(m, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4542 = m.ExcPending
	if v4542 != 0 {
		goto L32
	} else {
		goto L1025
	}
L1025:
	;
	v4544 = F_ReadDir(m, v4541, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4545 = m.ExcPending
	if v4545 != 0 {
		goto L32
	} else {
		goto L1026
	}
L1026:
	;
	if v4544 != 0 {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v4549 = v4544
	goto L1030
L1028:
	;
	goto L1029
L1029:
	;
	v4728 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v4728+int32(2304))
	mBase = m.M
	v4732 = m.ExcPending
	if v4732 != 0 {
		goto L32
	} else {
		goto L1060
	}
L1030:
	;
	v4581 = v4549 + int32(19)
	v4582 = F_strlen(m, v4581)
	mBase = m.M
	if v4582 != int32(16) {
		goto L1032
	} else {
		goto L1033
	}
L1031:
	;
	goto L1029
L1032:
	;
	v4691 = F_ReadDir(m, v4541, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L32
	} else {
		goto L1058
	}
L1033:
	;
	v4585 = int32(_a_F_StartupXLOG_171)
	v4589 = m.G0
	v4591 = v4589 - int32(32)
	v4592 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4591)+24)) = v4592
	*(*int64)(unsafe.Add(mBase, uint32(v4591)+16)) = v4592
	*(*int64)(unsafe.Add(mBase, uint32(v4591)+8)) = v4592
	*(*int64)(unsafe.Add(mBase, uint32(v4591))) = v4592
	v4600 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[71])))
	if v4600 == int32(0) {
		goto L1035
	} else {
		goto L1036
	}
L1034:
	;
	if v4668 != int32(16) {
		goto L1032
	} else {
		goto L1053
	}
L1035:
	;
	v4668 = int32(0)
	goto L1034
L1036:
	;
	goto L1037
L1037:
	;
	v4604 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[72])))
	if v4604 == int32(0) {
		goto L1038
	} else {
		goto L1039
	}
L1038:
	;
	v4608 = v4581
	goto L1041
L1039:
	;
	goto L1040
L1040:
	;
	v4618 = v4585
	v4619 = v4600
	goto L1044
L1041:
	;
	v4614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4608))))
	if v4614 == v4600 {
		v4608 = v4608 + int32(1)
		goto L1041
	} else {
		goto L1043
	}
L1042:
	;
	v4668 = v4608 - v4581
	goto L1034
L1043:
	;
	goto L1042
L1044:
	;
	v4626 = v4591 + int32(base.Ui32(v4619)>>(uint(int32(3))%32))&int32(28)
	v4627 = *(*int32)(unsafe.Add(mBase, uint32(v4626)))
	v4628 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4626))) = v4627 | v4628<<(uint(v4619)%32)
	v4632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4618)+1)))
	if v4632 != 0 {
		v4618 = v4618 + v4628
		v4619 = v4632
		goto L1044
	} else {
		goto L1046
	}
L1045:
	;
	v4635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4581))))
	if v4635 == int32(0) {
		v4658 = v4581
		goto L1047
	} else {
		goto L1048
	}
L1046:
	;
	goto L1045
L1047:
	;
	v4668 = v4658 - v4581
	goto L1034
L1048:
	;
	v4639 = v4581
	v4640 = v4635
	goto L1049
L1049:
	;
	v4648 = *(*int32)(unsafe.Add(mBase, uint32(v4591+int32(base.Ui32(v4640)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v4648)>>(uint(v4640)%32))&int32(1) == int32(0) {
		v4658 = v4639
		goto L1047
	} else {
		goto L1051
	}
L1050:
	;
	v4658 = v4656
	goto L1047
L1051:
	;
	v4654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4639)+1)))
	v4656 = v4639 + int32(1)
	if v4654 != 0 {
		v4639 = v4656
		v4640 = v4654
		goto L1049
	} else {
		goto L1052
	}
L1052:
	;
	goto L1050
L1053:
	;
	v4674 = F_strtox_2(m, v4581, int32(0), int32(16), int64(-1))
	mBase = m.M
	goto L1054
L1054:
	;
	v4678 = int32(0)
	v4680 = F_ProcessTwoPhaseBuffer(m, base.I32_wrap_i64(v4674), int64(0), int32(1), v4678, v4678)
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L32
	} else {
		goto L1055
	}
L1055:
	;
	if v4680 == int32(0) {
		goto L1032
	} else {
		goto L1056
	}
L1056:
	;
	v4684 = int64(0)
	F_PrepareRedoAdd(m, v4680, v4684, v4684, int32(0))
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L32
	} else {
		goto L1057
	}
L1057:
	;
	goto L1032
L1058:
	;
	if v4691 != 0 {
		v4549 = v4691
		goto L1030
	} else {
		goto L1059
	}
L1059:
	;
	goto L1031
L1060:
	;
	F_FreeDir(m, v4541)
	mBase = m.M
	v4734 = m.ExcPending
	if v4734 != 0 {
		goto L32
	} else {
		goto L1061
	}
L1061:
	;
	if base.Ui32(v391) <= base.Ui32(int32(-3)) {
		goto L1063
	} else {
		goto L1064
	}
L1062:
	;
	v6000 = int32(1)
	v6001 = v2892 & v6000
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[73])) = uint8(v6001)
	v6004 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v6004)+200)) = v2894
	*(*int64)(unsafe.Add(mBase, uint32(v6004)+152)) = v2894
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[74])) = uint8(v6001)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[75])) = v2894
	v6012 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v6012 == v6000 {
		goto L1351
	} else {
		goto L1352
	}
L1063:
	;
	v4737 = m.G0
	v4739 = v4737 - int32(48)
	m.G0 = v4739
	v4743 = F_unlink(m, int32(_a_F_StartupXLOG_172))
	mBase = m.M
	if v4743 != 0 {
		goto L1068
	} else {
		goto L1069
	}
L1064:
	;
	goto L1065
L1065:
	;
	v4892 = m.G0
	v4894 = v4892 - int32(528)
	m.G0 = v4894
	v4897 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[76]))
	v4900 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4901 = m.ExcPending
	if v4901 != 0 {
		goto L32
	} else {
		goto L1101
	}
L1066:
	;
	v4799 = m.G0
	v4800 = int32(16)
	v4801 = v4799 - v4800
	m.G0 = v4801
	F_gettimeofday(m, v4801)
	mBase = m.M
	v4804 = *(*int64)(unsafe.Add(mBase, uint32(v4801)))
	v4805 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4801)+8)))
	m.G0 = v4801 + v4800
	goto L1086
L1067:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v4792, int32(_a_F_StartupXLOG_174))
	mBase = m.M
	v4795 = m.ExcPending
	if v4795 != 0 {
		goto L32
	} else {
		goto L1085
	}
L1068:
	;
	v4745 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4745 == int32(44) {
		goto L1071
	} else {
		goto L1072
	}
L1069:
	;
	goto L1070
L1070:
	;
	v4780 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4781 = m.ExcPending
	if v4781 != 0 {
		goto L32
	} else {
		goto L1081
	}
L1071:
	;
	v4750 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4751 = m.ExcPending
	if v4751 != 0 {
		goto L32
	} else {
		goto L1074
	}
L1072:
	;
	goto L1073
L1073:
	;
	v4764 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L32
	} else {
		goto L1077
	}
L1074:
	;
	if v4750 == int32(0) {
		goto L1066
	} else {
		goto L1075
	}
L1075:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4739)+16)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_175), v4739+int32(16))
	mBase = m.M
	v4760 = m.ExcPending
	if v4760 != 0 {
		goto L32
	} else {
		goto L1076
	}
L1076:
	;
	v4792 = int32(530)
	goto L1067
L1077:
	;
	if v4764 == int32(0) {
		goto L1066
	} else {
		goto L1078
	}
L1078:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4769 = m.ExcPending
	if v4769 != 0 {
		goto L32
	} else {
		goto L1079
	}
L1079:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4739)+32)) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_176), v4739+int32(32))
	mBase = m.M
	v4776 = m.ExcPending
	if v4776 != 0 {
		goto L32
	} else {
		goto L1080
	}
L1080:
	;
	v4792 = int32(535)
	goto L1067
L1081:
	;
	if v4780 == int32(0) {
		goto L1066
	} else {
		goto L1082
	}
L1082:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4785 = m.ExcPending
	if v4785 != 0 {
		goto L32
	} else {
		goto L1083
	}
L1083:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4739))) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_177), v4739)
	mBase = m.M
	v4790 = m.ExcPending
	if v4790 != 0 {
		goto L32
	} else {
		goto L1084
	}
L1084:
	;
	v4792 = int32(542)
	goto L1067
L1085:
	;
	goto L1066
L1086:
	;
	v4818 = int32(1)
	goto L1087
L1087:
	;
	if base.Ui32(v4818) <= base.Ui32(int32(12)) {
		goto L1091
	} else {
		goto L1092
	}
L1088:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v4888 = m.ExcPending
	if v4888 != 0 {
		goto L32
	} else {
		goto L1100
	}
L1089:
	;
	v4884 = v4818 + int32(1)
	if v4884 != int32(33) {
		v4818 = v4884
		goto L1087
	} else {
		goto L1099
	}
L1090:
	;
	v4873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4872))))
	if v4873&int32(1) == int32(0) {
		goto L1089
	} else {
		goto L1097
	}
L1091:
	;
	v4872 = v4818*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1090
L1092:
	;
	goto L1093
L1093:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4818-int32(24)) {
		goto L1089
	} else {
		goto L1094
	}
L1094:
	;
	v4860 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v4860 == int32(0) {
		goto L1089
	} else {
		goto L1095
	}
L1095:
	;
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v4860+v4818<<(uint(int32(2))%32)-int32(96))))
	if v4868 == int32(0) {
		goto L1089
	} else {
		goto L1096
	}
L1096:
	;
	v4872 = v4868
	goto L1090
L1097:
	;
	v4878 = *(*int32)(unsafe.Add(mBase, uint32(v4872)+60))
	m.T0[v4878].(func(*base.Module, int64))(m, v4805+v4804*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v4880 = m.ExcPending
	if v4880 != 0 {
		goto L32
	} else {
		goto L1098
	}
L1098:
	;
	goto L1089
L1099:
	;
	goto L1088
L1100:
	;
	m.G0 = v4739 + int32(48)
	goto L1062
L1101:
	;
	if v4900 != 0 {
		goto L1102
	} else {
		goto L1103
	}
L1102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+432)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_179), v4894+int32(432))
	mBase = m.M
	v4908 = m.ExcPending
	if v4908 != 0 {
		goto L32
	} else {
		goto L1105
	}
L1103:
	;
	goto L1104
L1104:
	;
	v4916 = F_AllocateFile(m, int32(_a_F_StartupXLOG_172), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v4917 = m.ExcPending
	if v4917 != 0 {
		goto L32
	} else {
		goto L1108
	}
L1105:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1765), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v4913 = m.ExcPending
	if v4913 != 0 {
		goto L32
	} else {
		goto L1106
	}
L1106:
	;
	goto L1104
L1107:
	;
	m.G0 = v4894 + int32(528)
	goto L1062
L1108:
	;
	if v4916 == int32(0) {
		goto L1109
	} else {
		goto L1110
	}
L1109:
	;
	v4921 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4921 == int32(44) {
		goto L1112
	} else {
		goto L1113
	}
L1110:
	;
	goto L1111
L1111:
	;
	v5037 = F_fread(m, v4894+int32(524), int32(1), int32(4), v4916)
	mBase = m.M
	v5038 = m.ExcPending
	if v5038 != 0 {
		goto L32
	} else {
		goto L1136
	}
L1112:
	;
	v4945 = m.G0
	v4946 = int32(16)
	v4947 = v4945 - v4946
	m.G0 = v4947
	F_gettimeofday(m, v4947)
	mBase = m.M
	v4950 = *(*int64)(unsafe.Add(mBase, uint32(v4947)))
	v4951 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4947)+8)))
	m.G0 = v4947 + v4946
	goto L1119
L1113:
	;
	v4926 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4927 = m.ExcPending
	if v4927 != 0 {
		goto L32
	} else {
		goto L1114
	}
L1114:
	;
	if v4926 == int32(0) {
		goto L1112
	} else {
		goto L1115
	}
L1115:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4931 = m.ExcPending
	if v4931 != 0 {
		goto L32
	} else {
		goto L1116
	}
L1116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894))) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_181), v4894)
	mBase = m.M
	v4936 = m.ExcPending
	if v4936 != 0 {
		goto L32
	} else {
		goto L1117
	}
L1117:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1782), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v4941 = m.ExcPending
	if v4941 != 0 {
		goto L32
	} else {
		goto L1118
	}
L1118:
	;
	goto L1112
L1119:
	;
	v4961 = int32(1)
	goto L1120
L1120:
	;
	if base.Ui32(v4961) <= base.Ui32(int32(12)) {
		goto L1124
	} else {
		goto L1125
	}
L1121:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5032 = m.ExcPending
	if v5032 != 0 {
		goto L32
	} else {
		goto L1133
	}
L1122:
	;
	v5028 = v4961 + int32(1)
	if v5028 != int32(33) {
		v4961 = v5028
		goto L1120
	} else {
		goto L1132
	}
L1123:
	;
	v5018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5017))))
	if v5018&int32(1) == int32(0) {
		goto L1122
	} else {
		goto L1130
	}
L1124:
	;
	v5017 = v4961*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1123
L1125:
	;
	goto L1126
L1126:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4961-int32(24)) {
		goto L1122
	} else {
		goto L1127
	}
L1127:
	;
	v5006 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5006 == int32(0) {
		goto L1122
	} else {
		goto L1128
	}
L1128:
	;
	v5014 = *(*int32)(unsafe.Add(mBase, uint32(v5006+v4961<<(uint(int32(2))%32)-int32(96))))
	if v5014 == int32(0) {
		goto L1122
	} else {
		goto L1129
	}
L1129:
	;
	v5017 = v5014
	goto L1123
L1130:
	;
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v5017)+60))
	m.T0[v5023].(func(*base.Module, int64))(m, v4951+v4950*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5025 = m.ExcPending
	if v5025 != 0 {
		goto L32
	} else {
		goto L1131
	}
L1131:
	;
	goto L1122
L1132:
	;
	goto L1121
L1133:
	;
	goto L1107
L1134:
	;
	v5908 = F_FreeFile(m, v4916)
	mBase = m.M
	v5909 = m.ExcPending
	if v5909 != 0 {
		goto L32
	} else {
		goto L1344
	}
L1135:
	;
	v5769 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5770 = m.ExcPending
	if v5770 != 0 {
		goto L32
	} else {
		goto L1323
	}
L1136:
	;
	if v5037 != int32(4) {
		goto L1137
	} else {
		goto L1138
	}
L1137:
	;
	v5043 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5044 = m.ExcPending
	if v5044 != 0 {
		goto L32
	} else {
		goto L1140
	}
L1138:
	;
	goto L1139
L1139:
	;
	v5056 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+524))
	if v5056 == int32(27638967) {
		goto L1144
	} else {
		goto L1145
	}
L1140:
	;
	if v5043 == int32(0) {
		goto L1135
	} else {
		goto L1141
	}
L1141:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_182), int32(0))
	mBase = m.M
	v5050 = m.ExcPending
	if v5050 != 0 {
		goto L32
	} else {
		goto L1142
	}
L1142:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1792), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L32
	} else {
		goto L1143
	}
L1143:
	;
	goto L1135
L1144:
	;
	goto L1152
L1145:
	;
	goto L1146
L1146:
	;
	v5716 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5717 = m.ExcPending
	if v5717 != 0 {
		goto L32
	} else {
		goto L1319
	}
L1147:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5710, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5713 = m.ExcPending
	if v5713 != 0 {
		goto L32
	} else {
		goto L1318
	}
L1148:
	;
	v5694 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5695 = m.ExcPending
	if v5695 != 0 {
		goto L32
	} else {
		goto L1315
	}
L1149:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5682, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5689 = m.ExcPending
	if v5689 != 0 {
		goto L32
	} else {
		goto L1314
	}
L1150:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5674, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5681 = m.ExcPending
	if v5681 != 0 {
		goto L32
	} else {
		goto L1313
	}
L1151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5658 = m.ExcPending
	if v5658 != 0 {
		goto L32
	} else {
		goto L1310
	}
L1152:
	;
	v5096 = F_do_getc(m, v4916)
	mBase = m.M
	v5097 = m.ExcPending
	if v5097 != 0 {
		goto L32
	} else {
		goto L1160
	}
L1153:
	;
	v5640 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5641 = m.ExcPending
	if v5641 != 0 {
		goto L32
	} else {
		goto L1307
	}
L1154:
	;
	v5517 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	v5522 = F_dshash_find_or_insert(m, v5517, v4894+int32(504), v4894+int32(523))
	mBase = m.M
	v5523 = m.ExcPending
	if v5523 != 0 {
		goto L32
	} else {
		goto L1278
	}
L1155:
	;
	v5454 = *(*int32)(unsafe.Add(mBase, uint32(v5453)+48))
	if v5454 == int32(0) {
		goto L1262
	} else {
		goto L1263
	}
L1156:
	;
	v5438 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5439 = m.ExcPending
	if v5439 != 0 {
		goto L32
	} else {
		goto L1258
	}
L1157:
	;
	v5417 = F_do_getc(m, v4916)
	mBase = m.M
	v5418 = m.ExcPending
	if v5418 != 0 {
		goto L32
	} else {
		goto L1252
	}
L1158:
	;
	v5224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	if v5224 != 0 {
		goto L1194
	} else {
		goto L1195
	}
L1159:
	;
	v5104 = F_fread(m, v4894+int32(436), int32(1), int32(4), v4916)
	mBase = m.M
	v5105 = m.ExcPending
	if v5105 != 0 {
		goto L32
	} else {
		goto L1161
	}
L1160:
	;
	switch v5096 - int32(69) {
	case 0:
		goto L1157
	case 1:
		goto L1159
	default:
		goto L1156
	case 9, 14:
		goto L1158
	}
L1161:
	;
	if v5104 != int32(4) {
		goto L1162
	} else {
		goto L1163
	}
L1162:
	;
	v5110 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5111 = m.ExcPending
	if v5111 != 0 {
		goto L32
	} else {
		goto L1165
	}
L1163:
	;
	goto L1164
L1164:
	;
	v5122 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+436))
	v5124 = v5122 - int32(24)
	v5128 = base.B2i32(base.Ui32(v5122-int32(1)) < base.Ui32(int32(12)))
	if v5128|base.B2i32(base.Ui32(v5124) < base.Ui32(int32(9))) == int32(0) {
		goto L1168
	} else {
		goto L1169
	}
L1165:
	;
	if v5110 == int32(0) {
		goto L1135
	} else {
		goto L1166
	}
L1166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+128)) = int32(70)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v4894+int32(128))
	mBase = m.M
	v5120 = m.ExcPending
	if v5120 != 0 {
		goto L32
	} else {
		goto L1167
	}
L1167:
	;
	v5710 = int32(1822)
	goto L1147
L1168:
	;
	v5136 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5137 = m.ExcPending
	if v5137 != 0 {
		goto L32
	} else {
		goto L1171
	}
L1169:
	;
	goto L1170
L1170:
	;
	if v5128 == int32(0) {
		goto L1176
	} else {
		goto L1177
	}
L1171:
	;
	if v5136 == int32(0) {
		goto L1135
	} else {
		goto L1172
	}
L1172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+116)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+112)) = v5122
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_184), v4894+int32(112))
	mBase = m.M
	v5147 = m.ExcPending
	if v5147 != 0 {
		goto L32
	} else {
		goto L1173
	}
L1173:
	;
	v5710 = int32(1829)
	goto L1147
L1174:
	;
	v5199 = *(*int32)(unsafe.Add(mBase, uint32(v5196)+16))
	v5202 = *(*int32)(unsafe.Add(mBase, uint32(v5196)+20))
	v5203 = F_fread(m, v5198+v5199, int32(1), v5202, v4916)
	mBase = m.M
	v5204 = m.ExcPending
	if v5204 != 0 {
		goto L32
	} else {
		goto L1189
	}
L1175:
	;
	v5195 = *(*int32)(unsafe.Add(mBase, uint32(v5158+(v4897+int32(_a_F_StartupXLOG_185)))))
	v5196 = v5162
	v5198 = v5195
	goto L1174
L1176:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5124) {
		goto L1180
	} else {
		goto L1181
	}
L1177:
	;
	goto L1178
L1178:
	;
	v5184 = v5122 * int32(72)
	v5185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5184)+uint32(_c_F_StartupXLOG[80]))))
	if v5185&int32(1) == int32(0) {
		goto L1148
	} else {
		goto L1188
	}
L1179:
	;
	v5180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5162))))
	if v5180&int32(1) != 0 {
		goto L1175
	} else {
		goto L1187
	}
L1180:
	;
	v5167 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5168 = m.ExcPending
	if v5168 != 0 {
		goto L32
	} else {
		goto L1184
	}
L1181:
	;
	v5154 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5154 == int32(0) {
		goto L1180
	} else {
		goto L1182
	}
L1182:
	;
	v5158 = v5122 << (uint(int32(2)) % 32)
	v5162 = *(*int32)(unsafe.Add(mBase, uint32(v5154+v5158-int32(96))))
	if v5162 != 0 {
		goto L1179
	} else {
		goto L1183
	}
L1183:
	;
	goto L1180
L1184:
	;
	if v5167 == int32(0) {
		goto L1135
	} else {
		goto L1185
	}
L1185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+100)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+96)) = v5122
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_186), v4894+int32(96))
	mBase = m.M
	v5178 = m.ExcPending
	if v5178 != 0 {
		goto L32
	} else {
		goto L1186
	}
L1186:
	;
	v5710 = int32(1837)
	goto L1147
L1187:
	;
	goto L1148
L1188:
	;
	v5192 = *(*int32)(unsafe.Add(mBase, uint32(v5184)+uint32(_c_F_StartupXLOG[81])))
	v5196 = v5184 + int32(_a_F_StartupXLOG_178)
	v5198 = v4897 + v5192
	goto L1174
L1189:
	;
	if v5203 == v5202 {
		goto L1152
	} else {
		goto L1190
	}
L1190:
	;
	v5208 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5209 = m.ExcPending
	if v5209 != 0 {
		goto L32
	} else {
		goto L1191
	}
L1191:
	;
	if v5208 == int32(0) {
		goto L1135
	} else {
		goto L1192
	}
L1192:
	;
	v5212 = *(*int32)(unsafe.Add(mBase, uint32(v5196)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+72)) = v5212
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+68)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+64)) = v5122
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_187), v4894-int32(-64))
	mBase = m.M
	v5221 = m.ExcPending
	if v5221 != 0 {
		goto L32
	} else {
		goto L1193
	}
L1193:
	;
	v5710 = int32(1863)
	goto L1147
L1194:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5226 = m.ExcPending
	if v5226 != 0 {
		goto L32
	} else {
		goto L1197
	}
L1195:
	;
	goto L1196
L1196:
	;
	if v5096 == int32(83) {
		goto L1198
	} else {
		goto L1199
	}
L1197:
	;
	goto L1196
L1198:
	;
	v5233 = F_fread(m, v4894+int32(504), int32(1), int32(16), v4916)
	mBase = m.M
	v5234 = m.ExcPending
	if v5234 != 0 {
		goto L32
	} else {
		goto L1201
	}
L1199:
	;
	goto L1200
L1200:
	;
	v5316 = F_fread(m, v4894+int32(500), int32(1), int32(4), v4916)
	mBase = m.M
	v5317 = m.ExcPending
	if v5317 != 0 {
		goto L32
	} else {
		goto L1222
	}
L1201:
	;
	if v5233 != int32(16) {
		goto L1202
	} else {
		goto L1203
	}
L1202:
	;
	v5239 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5240 = m.ExcPending
	if v5240 != 0 {
		goto L32
	} else {
		goto L1205
	}
L1203:
	;
	goto L1204
L1204:
	;
	v5251 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+504))
	v5253 = v5251 - int32(24)
	v5257 = base.B2i32(base.Ui32(v5251-int32(1)) < base.Ui32(int32(12)))
	if v5257|base.B2i32(base.Ui32(v5253) < base.Ui32(int32(9))) == int32(0) {
		goto L1208
	} else {
		goto L1209
	}
L1205:
	;
	if v5239 == int32(0) {
		goto L1135
	} else {
		goto L1206
	}
L1206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+304)) = int32(83)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_188), v4894+int32(304))
	mBase = m.M
	v5249 = m.ExcPending
	if v5249 != 0 {
		goto L32
	} else {
		goto L1207
	}
L1207:
	;
	v5682 = int32(1883)
	goto L1149
L1208:
	;
	v5265 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5266 = m.ExcPending
	if v5266 != 0 {
		goto L32
	} else {
		goto L1211
	}
L1209:
	;
	goto L1210
L1210:
	;
	if base.Ui32(v5251-int32(1)) < base.Ui32(int32(12)) {
		goto L1154
	} else {
		goto L1214
	}
L1211:
	;
	if v5265 == int32(0) {
		goto L1135
	} else {
		goto L1212
	}
L1212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+288)) = int32(83)
	v5271 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+272)) = v5271
	v5273 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+280)) = v5273
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_189), v4894+int32(272))
	mBase = m.M
	v5279 = m.ExcPending
	if v5279 != 0 {
		goto L32
	} else {
		goto L1213
	}
L1213:
	;
	v5682 = int32(1891)
	goto L1149
L1214:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5253) {
		goto L1215
	} else {
		goto L1216
	}
L1215:
	;
	v5296 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5297 = m.ExcPending
	if v5297 != 0 {
		goto L32
	} else {
		goto L1219
	}
L1216:
	;
	v5284 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5284 == int32(0) {
		goto L1215
	} else {
		goto L1217
	}
L1217:
	;
	v5292 = *(*int32)(unsafe.Add(mBase, uint32(v5284+v5251<<(uint(int32(2))%32)-int32(96))))
	if v5292 != 0 {
		goto L1154
	} else {
		goto L1218
	}
L1218:
	;
	goto L1215
L1219:
	;
	if v5296 == int32(0) {
		goto L1135
	} else {
		goto L1220
	}
L1220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+256)) = int32(83)
	v5302 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+240)) = v5302
	v5304 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+248)) = v5304
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_190), v4894+int32(240))
	mBase = m.M
	v5310 = m.ExcPending
	if v5310 != 0 {
		goto L32
	} else {
		goto L1221
	}
L1221:
	;
	v5682 = int32(1899)
	goto L1149
L1222:
	;
	if v5316 != int32(4) {
		goto L1223
	} else {
		goto L1224
	}
L1223:
	;
	v5322 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5323 = m.ExcPending
	if v5323 != 0 {
		goto L32
	} else {
		goto L1226
	}
L1224:
	;
	goto L1225
L1225:
	;
	v5337 = F_fread(m, v4894+int32(436), int32(1), int32(64), v4916)
	mBase = m.M
	v5338 = m.ExcPending
	if v5338 != 0 {
		goto L32
	} else {
		goto L1229
	}
L1226:
	;
	if v5322 == int32(0) {
		goto L1135
	} else {
		goto L1227
	}
L1227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+400)) = v5096
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v4894+int32(400))
	mBase = m.M
	v5331 = m.ExcPending
	if v5331 != 0 {
		goto L32
	} else {
		goto L1228
	}
L1228:
	;
	v5674 = int32(1912)
	goto L1150
L1229:
	;
	if v5337 != int32(64) {
		goto L1230
	} else {
		goto L1231
	}
L1230:
	;
	v5343 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5344 = m.ExcPending
	if v5344 != 0 {
		goto L32
	} else {
		goto L1233
	}
L1231:
	;
	goto L1232
L1232:
	;
	v5356 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+500))
	v5358 = v5356 - int32(24)
	v5360 = v5356 - int32(1)
	if base.B2i32(base.Ui32(v5360) < base.Ui32(int32(12)))|base.B2i32(base.Ui32(v5358) < base.Ui32(int32(9))) == int32(0) {
		goto L1236
	} else {
		goto L1237
	}
L1233:
	;
	if v5343 == int32(0) {
		goto L1135
	} else {
		goto L1234
	}
L1234:
	;
	v5347 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+384)) = v5347
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+388)) = v5096
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_191), v4894+int32(384))
	mBase = m.M
	v5354 = m.ExcPending
	if v5354 != 0 {
		goto L32
	} else {
		goto L1235
	}
L1235:
	;
	v5674 = int32(1918)
	goto L1150
L1236:
	;
	v5370 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5371 = m.ExcPending
	if v5371 != 0 {
		goto L32
	} else {
		goto L1239
	}
L1237:
	;
	goto L1238
L1238:
	;
	v5383 = base.B2i32(base.Ui32(int32(11)) < base.Ui32(v5360))
	if v5383 == int32(0) {
		goto L1242
	} else {
		goto L1243
	}
L1239:
	;
	if v5370 == int32(0) {
		goto L1135
	} else {
		goto L1240
	}
L1240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+372)) = v5096
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+368)) = v5356
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_184), v4894+int32(368))
	mBase = m.M
	v5380 = m.ExcPending
	if v5380 != 0 {
		goto L32
	} else {
		goto L1241
	}
L1241:
	;
	v5674 = int32(1924)
	goto L1150
L1242:
	;
	v5453 = v5356*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1155
L1243:
	;
	goto L1244
L1244:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5358) {
		goto L1245
	} else {
		goto L1246
	}
L1245:
	;
	v5405 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5406 = m.ExcPending
	if v5406 != 0 {
		goto L32
	} else {
		goto L1249
	}
L1246:
	;
	v5393 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5393 == int32(0) {
		goto L1245
	} else {
		goto L1247
	}
L1247:
	;
	v5401 = *(*int32)(unsafe.Add(mBase, uint32(v5393+v5356<<(uint(int32(2))%32)-int32(96))))
	if v5401 != 0 {
		v5453 = v5401
		goto L1155
	} else {
		goto L1248
	}
L1248:
	;
	goto L1245
L1249:
	;
	if v5405 == int32(0) {
		goto L1135
	} else {
		goto L1250
	}
L1250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+356)) = v5096
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+352)) = v5356
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_186), v4894+int32(352))
	mBase = m.M
	v5415 = m.ExcPending
	if v5415 != 0 {
		goto L32
	} else {
		goto L1251
	}
L1251:
	;
	v5674 = int32(1932)
	goto L1150
L1252:
	;
	if v5417 == int32(-1) {
		goto L1134
	} else {
		goto L1253
	}
L1253:
	;
	v5423 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5424 = m.ExcPending
	if v5424 != 0 {
		goto L32
	} else {
		goto L1254
	}
L1254:
	;
	if v5423 == int32(0) {
		goto L1135
	} else {
		goto L1255
	}
L1255:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_192), int32(0))
	mBase = m.M
	v5430 = m.ExcPending
	if v5430 != 0 {
		goto L32
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2010), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5435 = m.ExcPending
	if v5435 != 0 {
		goto L32
	} else {
		goto L1257
	}
L1257:
	;
	goto L1135
L1258:
	;
	if v5438 == int32(0) {
		goto L1135
	} else {
		goto L1259
	}
L1259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+48)) = v5096
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_193), v4894+int32(48))
	mBase = m.M
	v5447 = m.ExcPending
	if v5447 != 0 {
		goto L32
	} else {
		goto L1260
	}
L1260:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2017), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5452 = m.ExcPending
	if v5452 != 0 {
		goto L32
	} else {
		goto L1261
	}
L1261:
	;
	goto L1135
L1262:
	;
	v5459 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5460 = m.ExcPending
	if v5460 != 0 {
		goto L32
	} else {
		goto L1265
	}
L1263:
	;
	goto L1264
L1264:
	;
	v5475 = m.T0[v5454].(func(*base.Module, int32, int32) int32)(m, v4894+int32(436), v4894+int32(504))
	mBase = m.M
	v5476 = m.ExcPending
	if v5476 != 0 {
		goto L32
	} else {
		goto L1268
	}
L1265:
	;
	if v5459 == int32(0) {
		goto L1135
	} else {
		goto L1266
	}
L1266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+324)) = v5096
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+320)) = v5356
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_194), v4894+int32(320))
	mBase = m.M
	v5469 = m.ExcPending
	if v5469 != 0 {
		goto L32
	} else {
		goto L1267
	}
L1267:
	;
	v5674 = int32(1939)
	goto L1150
L1268:
	;
	if v5475 != 0 {
		goto L1154
	} else {
		goto L1269
	}
L1269:
	;
	if base.Ui32(int32(11)) < base.Ui32(v5360) {
		goto L1270
	} else {
		goto L1271
	}
L1270:
	;
	v5478 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	v5484 = *(*int32)(unsafe.Add(mBase, uint32(v5478+v5356<<(uint(int32(2))%32)-int32(96))))
	v5489 = v5484
	goto L1272
L1271:
	;
	v5489 = v5356*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1272
L1272:
	;
	v5490 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5489)+20)))
	v5492 = F___fseeko_unlocked(m, v4916, v5490, int32(1))
	mBase = m.M
	v5493 = m.ExcPending
	if v5493 != 0 {
		goto L32
	} else {
		goto L1273
	}
L1273:
	;
	if v5492 == int32(0) {
		goto L1152
	} else {
		goto L1274
	}
L1274:
	;
	v5498 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5499 = m.ExcPending
	if v5499 != 0 {
		goto L32
	} else {
		goto L1275
	}
L1275:
	;
	if v5498 == int32(0) {
		goto L1135
	} else {
		goto L1276
	}
L1276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+344)) = v5096
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+340)) = v5356
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+336)) = v4894 + int32(436)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_195), v4894+int32(336))
	mBase = m.M
	v5511 = m.ExcPending
	if v5511 != 0 {
		goto L32
	} else {
		goto L1277
	}
L1277:
	;
	v5674 = int32(1949)
	goto L1150
L1278:
	;
	v5524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4894)+523)))
	if v5524 == int32(1) {
		goto L1279
	} else {
		goto L1280
	}
L1279:
	;
	v5528 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	F_dshash_release_lock(m, v5528, v5522)
	mBase = m.M
	v5530 = m.ExcPending
	if v5530 != 0 {
		goto L32
	} else {
		goto L1282
	}
L1280:
	;
	goto L1281
L1281:
	;
	v5548 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+504))
	v5549 = int32(0)
	v5550 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5522)+20)) = v5550
	*(*uint8)(unsafe.Add(mBase, uint32(v5522)+16)) = uint8(v5549)
	*(*int32)(unsafe.Add(mBase, uint32(v5522)+24)) = v5549
	v5557 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	if base.Ui32(v5548-v5550) <= base.Ui32(int32(11)) {
		goto L1287
	} else {
		goto L1288
	}
L1282:
	;
	v5533 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5534 = m.ExcPending
	if v5534 != 0 {
		goto L32
	} else {
		goto L1283
	}
L1283:
	;
	if v5533 == int32(0) {
		goto L1135
	} else {
		goto L1284
	}
L1284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+160)) = v5096
	v5538 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+144)) = v5538
	v5540 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+152)) = v5540
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_196), v4894+int32(144))
	mBase = m.M
	v5546 = m.ExcPending
	if v5546 != 0 {
		goto L32
	} else {
		goto L1285
	}
L1285:
	;
	v5682 = int32(1972)
	goto L1149
L1286:
	;
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v5586)+4))
	v5589 = F_dsa_allocate_extended(m, v5557, v5587, int32(6))
	mBase = m.M
	v5590 = m.ExcPending
	if v5590 != 0 {
		goto L32
	} else {
		goto L1293
	}
L1287:
	;
	v5586 = v5548*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1286
L1288:
	;
	goto L1289
L1289:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5548-int32(24)) {
		v5584 = int32(0)
		goto L1290
	} else {
		goto L1291
	}
L1290:
	;
	v5586 = v5584
	goto L1286
L1291:
	;
	v5572 = int32(0)
	v5574 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5574 == v5572 {
		v5584 = v5572
		goto L1290
	} else {
		goto L1292
	}
L1292:
	;
	v5582 = *(*int32)(unsafe.Add(mBase, uint32(v5574+v5548<<(uint(int32(2))%32)-int32(96))))
	v5584 = v5582
	goto L1290
L1293:
	;
	if v5589 != 0 {
		goto L1294
	} else {
		goto L1295
	}
L1294:
	;
	v5592 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	v5593 = F_dsa_get_address(m, v5592, v5589)
	mBase = m.M
	v5594 = m.ExcPending
	if v5594 != 0 {
		goto L32
	} else {
		goto L1297
	}
L1295:
	;
	v5606 = v5549
	goto L1296
L1296:
	;
	v5608 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	F_dshash_release_lock(m, v5608, v5522)
	mBase = m.M
	v5610 = m.ExcPending
	if v5610 != 0 {
		goto L32
	} else {
		goto L1299
	}
L1297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5593))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v5522)+28)) = v5589
	v5599 = v5593 + int32(4)
	v5600 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v5599))) = uint16(v5600)
	*(*int32)(unsafe.Add(mBase, uint32(v5599)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v5599)+8)) = int64(-1)
	goto L1298
L1298:
	;
	v5606 = v5593
	goto L1296
L1299:
	;
	if v5606 == int32(0) {
		goto L1151
	} else {
		goto L1300
	}
L1300:
	;
	v5613 = *(*int32)(unsafe.Add(mBase, uint32(v4894)+504))
	if base.Ui32(v5613-int32(1)) <= base.Ui32(int32(11)) {
		goto L1302
	} else {
		goto L1303
	}
L1301:
	;
	v5631 = *(*int32)(unsafe.Add(mBase, uint32(v5630)+16))
	v5634 = *(*int32)(unsafe.Add(mBase, uint32(v5630)+20))
	v5635 = F_fread(m, v5606+v5631, int32(1), v5634, v4916)
	mBase = m.M
	v5636 = m.ExcPending
	if v5636 != 0 {
		goto L32
	} else {
		goto L1305
	}
L1302:
	;
	v5630 = v5613*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1301
L1303:
	;
	goto L1304
L1304:
	;
	v5623 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	v5629 = *(*int32)(unsafe.Add(mBase, uint32(v5623+v5613<<(uint(int32(2))%32)-int32(96))))
	v5630 = v5629
	goto L1301
L1305:
	;
	if v5635 == v5634 {
		goto L1152
	} else {
		goto L1306
	}
L1306:
	;
	goto L1153
L1307:
	;
	if v5640 == int32(0) {
		goto L1135
	} else {
		goto L1308
	}
L1308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+224)) = v5096
	v5645 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+208)) = v5645
	v5647 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+216)) = v5647
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_197), v4894+int32(208))
	mBase = m.M
	v5653 = m.ExcPending
	if v5653 != 0 {
		goto L32
	} else {
		goto L1309
	}
L1309:
	;
	v5682 = int32(1996)
	goto L1149
L1310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+192)) = v5096
	v5660 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+176)) = v5660
	v5662 = *(*int64)(unsafe.Add(mBase, uint32(v4894)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4894)+184)) = v5662
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_198), v4894+int32(176))
	mBase = m.M
	v5668 = m.ExcPending
	if v5668 != 0 {
		goto L32
	} else {
		goto L1311
	}
L1311:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1987), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5673 = m.ExcPending
	if v5673 != 0 {
		goto L32
	} else {
		goto L1312
	}
L1312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1313:
	;
	goto L1135
L1314:
	;
	goto L1135
L1315:
	;
	if v5694 == int32(0) {
		goto L1135
	} else {
		goto L1316
	}
L1316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+84)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+80)) = v5122
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_199), v4894+int32(80))
	mBase = m.M
	v5705 = m.ExcPending
	if v5705 != 0 {
		goto L32
	} else {
		goto L1317
	}
L1317:
	;
	v5710 = int32(1844)
	goto L1147
L1318:
	;
	goto L1135
L1319:
	;
	if v5716 == int32(0) {
		goto L1135
	} else {
		goto L1320
	}
L1320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+420)) = int32(27638967)
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+416)) = v5056
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_200), v4894+int32(416))
	mBase = m.M
	v5727 = m.ExcPending
	if v5727 != 0 {
		goto L32
	} else {
		goto L1321
	}
L1321:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1799), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5732 = m.ExcPending
	if v5732 != 0 {
		goto L32
	} else {
		goto L1322
	}
L1322:
	;
	goto L1135
L1323:
	;
	if v5769 != 0 {
		goto L1324
	} else {
		goto L1325
	}
L1324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+32)) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_201), v4894+int32(32))
	mBase = m.M
	v5777 = m.ExcPending
	if v5777 != 0 {
		goto L32
	} else {
		goto L1327
	}
L1325:
	;
	goto L1326
L1326:
	;
	v5786 = m.G0
	v5787 = int32(16)
	v5788 = v5786 - v5787
	m.G0 = v5788
	F_gettimeofday(m, v5788)
	mBase = m.M
	v5791 = *(*int64)(unsafe.Add(mBase, uint32(v5788)))
	v5792 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5788)+8)))
	m.G0 = v5788 + v5787
	goto L1329
L1327:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2032), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5782 = m.ExcPending
	if v5782 != 0 {
		goto L32
	} else {
		goto L1328
	}
L1328:
	;
	goto L1326
L1329:
	;
	v5802 = int32(1)
	goto L1330
L1330:
	;
	if base.Ui32(v5802) <= base.Ui32(int32(12)) {
		goto L1334
	} else {
		goto L1335
	}
L1331:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5873 = m.ExcPending
	if v5873 != 0 {
		goto L32
	} else {
		goto L1343
	}
L1332:
	;
	v5869 = v5802 + int32(1)
	if v5869 != int32(33) {
		v5802 = v5869
		goto L1330
	} else {
		goto L1342
	}
L1333:
	;
	v5859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5858))))
	if v5859&int32(1) == int32(0) {
		goto L1332
	} else {
		goto L1340
	}
L1334:
	;
	v5858 = v5802*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1333
L1335:
	;
	goto L1336
L1336:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5802-int32(24)) {
		goto L1332
	} else {
		goto L1337
	}
L1337:
	;
	v5847 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5847 == int32(0) {
		goto L1332
	} else {
		goto L1338
	}
L1338:
	;
	v5855 = *(*int32)(unsafe.Add(mBase, uint32(v5847+v5802<<(uint(int32(2))%32)-int32(96))))
	if v5855 == int32(0) {
		goto L1332
	} else {
		goto L1339
	}
L1339:
	;
	v5858 = v5855
	goto L1333
L1340:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v5858)+60))
	m.T0[v5864].(func(*base.Module, int64))(m, v5792+v5791*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5866 = m.ExcPending
	if v5866 != 0 {
		goto L32
	} else {
		goto L1341
	}
L1341:
	;
	goto L1332
L1342:
	;
	goto L1331
L1343:
	;
	goto L1134
L1344:
	;
	v5912 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v5913 = m.ExcPending
	if v5913 != 0 {
		goto L32
	} else {
		goto L1345
	}
L1345:
	;
	if v5912 != 0 {
		goto L1346
	} else {
		goto L1347
	}
L1346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4894)+16)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_202), v4894+int32(16))
	mBase = m.M
	v5920 = m.ExcPending
	if v5920 != 0 {
		goto L32
	} else {
		goto L1349
	}
L1347:
	;
	goto L1348
L1348:
	;
	v5927 = F_unlink(m, int32(_a_F_StartupXLOG_172))
	mBase = m.M
	goto L1107
L1349:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2025), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5925 = m.ExcPending
	if v5925 != 0 {
		goto L32
	} else {
		goto L1350
	}
L1350:
	;
	goto L1348
L1351:
	;
	v6017 = base.AtomicRmwXchg32(m, v6004, int32(440), int32(1))
	if v6017 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L1352:
	;
	goto L1353
L1353:
	;
	v9232 = m.G0
	v9234 = v9232 - int32(272)
	m.G0 = v9234
	v9237 = F_palloc(m, int32(72))
	mBase = m.M
	v9238 = m.ExcPending
	if v9238 != 0 {
		goto L32
	} else {
		goto L2015
	}
L1354:
	;
	v6019 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v6019+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_203), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v6026 = m.ExcPending
	if v6026 != 0 {
		goto L32
	} else {
		goto L1357
	}
L1355:
	;
	goto L1356
L1356:
	;
	v6028 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v6030 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	*(*int32)(unsafe.Add(mBase, uint32(v6028)+316)) = v6030
	v6032 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6028)+440)), uint32(v6032))
	v6036 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	v6038 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v6036, v6038)
	mBase = m.M
	v6040 = m.ExcPending
	if v6040 != 0 {
		goto L32
	} else {
		goto L1358
	}
L1357:
	;
	goto L1356
L1358:
	;
	v6041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4093)))
	if v6041 == int32(1) {
		goto L1359
	} else {
		goto L1360
	}
L1359:
	;
	v6044 = int32(_a_F_StartupXLOG_204)
	v6045 = F_unlink(m, v6044)
	mBase = m.M
	v6049 = F_durable_rename(m, int32(_a_F_StartupXLOG_59), v6044, int32(22))
	mBase = m.M
	v6050 = m.ExcPending
	if v6050 != 0 {
		goto L32
	} else {
		goto L1362
	}
L1360:
	;
	goto L1361
L1361:
	;
	v6051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4094)))
	if v6051 == int32(1) {
		goto L1363
	} else {
		goto L1364
	}
L1362:
	;
	goto L1361
L1363:
	;
	v6054 = int32(_a_F_StartupXLOG_91)
	v6055 = F_unlink(m, v6054)
	mBase = m.M
	v6059 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v6054, int32(22))
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
		goto L32
	} else {
		goto L1366
	}
L1364:
	;
	goto L1365
L1365:
	;
	v6063 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	if v6063 == int32(1) {
		goto L1367
	} else {
		goto L1368
	}
L1366:
	;
	goto L1365
L1367:
	;
	v6067 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v6068 = *(*int64)(unsafe.Add(mBase, uint32(v6067)+136))
	v6070 = v6068
	goto L1369
L1368:
	;
	v6070 = int64(0)
	goto L1369
L1369:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[83])) = v6070
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v6073 = m.ExcPending
	if v6073 != 0 {
		goto L32
	} else {
		goto L1370
	}
L1370:
	;
	F_ResetUnloggedRelations(m, int32(1))
	mBase = m.M
	v6076 = m.ExcPending
	if v6076 != 0 {
		goto L32
	} else {
		goto L1371
	}
L1371:
	;
	v6077 = m.G0
	v6079 = v6077 - int32(1072)
	m.G0 = v6079
	v6082 = F_AllocateDir(m, int32(_a_F_StartupXLOG_205))
	mBase = m.M
	v6083 = m.ExcPending
	if v6083 != 0 {
		goto L32
	} else {
		goto L1372
	}
L1372:
	;
	v6086 = F_ReadDirExtended(m, v6082, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6087 = m.ExcPending
	if v6087 != 0 {
		goto L32
	} else {
		goto L1373
	}
L1373:
	;
	if v6086 != 0 {
		goto L1374
	} else {
		goto L1375
	}
L1374:
	;
	v6088 = v6086
	goto L1377
L1375:
	;
	goto L1376
L1376:
	;
	F_FreeDir(m, v6082)
	mBase = m.M
	v6205 = m.ExcPending
	if v6205 != 0 {
		goto L32
	} else {
		goto L1394
	}
L1377:
	;
	v6122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6088)+19)))
	if v6122 != int32(46) {
		goto L1380
	} else {
		goto L1381
	}
L1378:
	;
	goto L1376
L1379:
	;
	v6168 = F_ReadDirExtended(m, v6082, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6169 = m.ExcPending
	if v6169 != 0 {
		goto L32
	} else {
		goto L1392
	}
L1380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6079)+16)) = v6088 + int32(19)
	v6138 = v6079 + int32(32)
	v6143 = F_pg_snprintf(m, v6138, int32(1037), int32(_a_F_StartupXLOG_206), v6079+int32(16))
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L32
	} else {
		goto L1385
	}
L1381:
	;
	v6125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6088)+20)))
	if v6125 == int32(0) {
		goto L1379
	} else {
		goto L1382
	}
L1382:
	;
	v6128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6088)+20)))
	if v6128 != int32(46) {
		goto L1380
	} else {
		goto L1383
	}
L1383:
	;
	v6131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6088)+21)))
	if v6131 == int32(0) {
		goto L1379
	} else {
		goto L1384
	}
L1384:
	;
	goto L1380
L1385:
	;
	v6145 = F_unlink(m, v6138)
	mBase = m.M
	if v6145 == int32(0) {
		goto L1379
	} else {
		goto L1386
	}
L1386:
	;
	v6150 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6151 = m.ExcPending
	if v6151 != 0 {
		goto L32
	} else {
		goto L1387
	}
L1387:
	;
	if v6150 == int32(0) {
		goto L1379
	} else {
		goto L1388
	}
L1388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v6155 = m.ExcPending
	if v6155 != 0 {
		goto L32
	} else {
		goto L1389
	}
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6079))) = v6138
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v6079)
	mBase = m.M
	v6159 = m.ExcPending
	if v6159 != 0 {
		goto L32
	} else {
		goto L1390
	}
L1390:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_207), int32(1609), int32(_a_F_StartupXLOG_208))
	mBase = m.M
	v6164 = m.ExcPending
	if v6164 != 0 {
		goto L32
	} else {
		goto L1391
	}
L1391:
	;
	goto L1379
L1392:
	;
	if v6168 != 0 {
		v6088 = v6168
		goto L1377
	} else {
		goto L1393
	}
L1393:
	;
	goto L1378
L1394:
	;
	m.G0 = v6079 + int32(1072)
	v6210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v6210 != int32(1) {
		goto L1395
	} else {
		goto L1396
	}
L1395:
	;
	v6433 = m.G0
	v6435 = v6433 - int32(848)
	m.G0 = v6435
	v6438 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6441 = base.AtomicRmwXchg32(m, v6438, int32(96), int32(1))
	if v6441 != 0 {
		goto L1426
	} else {
		goto L1427
	}
L1396:
	;
	v6214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v6214&int32(1) == int32(0) {
		goto L1395
	} else {
		goto L1397
	}
L1397:
	;
	v6221 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v6222 = m.ExcPending
	if v6222 != 0 {
		goto L32
	} else {
		goto L1398
	}
L1398:
	;
	if v6221 != 0 {
		goto L1399
	} else {
		goto L1400
	}
L1399:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_209), int32(0))
	mBase = m.M
	v6226 = m.ExcPending
	if v6226 != 0 {
		goto L32
	} else {
		goto L1402
	}
L1400:
	;
	goto L1401
L1401:
	;
	v6232 = m.G0
	v6234 = v6232 + int32(-64)
	m.G0 = v6234
	*(*int64)(unsafe.Add(mBase, uint32(v6234)+24)) = int64(68719476748)
	v6242 = v6232 + int32(-56)
	v6244 = F_hash_create(m, int32(_a_F_StartupXLOG_210), int32(64), v6242, int32(40))
	mBase = m.M
	v6245 = m.ExcPending
	if v6245 != 0 {
		goto L32
	} else {
		goto L1404
	}
L1402:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_211), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v6231 = m.ExcPending
	if v6231 != 0 {
		goto L32
	} else {
		goto L1403
	}
L1403:
	;
	goto L1401
L1404:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[84])) = v6244
	*(*int64)(unsafe.Add(mBase, uint32(v6234)+24)) = int64(34359738372)
	v6253 = F_hash_create(m, int32(_a_F_StartupXLOG_212), int32(64), v6242, int32(40))
	mBase = m.M
	v6254 = m.ExcPending
	if v6254 != 0 {
		goto L32
	} else {
		goto L1405
	}
L1405:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[85])) = v6253
	F_SharedInvalBackendInit(m, int32(1))
	mBase = m.M
	v6258 = m.ExcPending
	if v6258 != 0 {
		goto L32
	} else {
		goto L1406
	}
L1406:
	;
	v6260 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[86]))
	v6262 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v6260)+52)) = v6262
	*(*int32)(unsafe.Add(mBase, uint32(v6234)+56)) = v6262
	v6266 = int32(_a_F_StartupXLOG_213)
	v6267 = int32(1)
	v6269 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[88]))
	if base.Ui32(v6269) <= base.Ui32(v6267) {
		goto L1408
	} else {
		goto L1409
	}
L1407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6234)+60)) = v6272
	v6277 = *(*int64)(unsafe.Add(mBase, uint32(v6234)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v6234))) = v6277
	F_VirtualXactLockTableInsert(m, v6234)
	mBase = m.M
	v6280 = m.ExcPending
	if v6280 != 0 {
		goto L32
	} else {
		goto L1411
	}
L1408:
	;
	v6272 = v6267
	goto L1410
L1409:
	;
	v6272 = v6269
	goto L1410
L1410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[88])) = v6272 + int32(1)
	goto L1407
L1411:
	;
	v6282 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89])) = v6282
	m.G0 = v6234 - int32(-64)
	v6287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4095)))
	if v6287 == v6282 {
		goto L1412
	} else {
		goto L1413
	}
L1412:
	;
	v6294 = F_PrescanPreparedTransactions(m, v39+int32(_a_F_StartupXLOG_4), v39+int32(_a_F_StartupXLOG_3))
	mBase = m.M
	v6295 = m.ExcPending
	if v6295 != 0 {
		goto L32
	} else {
		goto L1415
	}
L1413:
	;
	v6296 = v2891
	goto L1414
L1414:
	;
	v6298 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v6299 = *(*int32)(unsafe.Add(mBase, uint32(v6298)+8))
	v6303 = v6299
	goto L1416
L1415:
	;
	v6296 = v6294
	goto L1414
L1416:
	;
	v6335 = v6303 - int32(1)
	if base.Ui32(v6335) < base.Ui32(int32(3)) {
		v6303 = v6335
		goto L1416
	} else {
		goto L1418
	}
L1417:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[90])) = v6335
	F_StartupSUBTRANS(m, v6296)
	mBase = m.M
	v6341 = m.ExcPending
	if v6341 != 0 {
		goto L32
	} else {
		goto L1419
	}
L1418:
	;
	goto L1417
L1419:
	;
	v6342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4095)))
	if v6342 != int32(1) {
		goto L1395
	} else {
		goto L1420
	}
L1420:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v6346 = m.ExcPending
	if v6346 != 0 {
		goto L32
	} else {
		goto L1421
	}
L1421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v6296
	v6348 = base.I32_wrap_i64(v2903)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v6348
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[93]))) = int64(8589934592)
	v6352 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[94])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v6352
	v6357 = v6348
	goto L1422
L1422:
	;
	v6389 = v6357 - int32(1)
	if base.Ui32(v6389) < base.Ui32(int32(3)) {
		v6357 = v6389
		goto L1422
	} else {
		goto L1424
	}
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[95]))) = v6389
	v6393 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[96])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[97]))) = v6393
	F_ProcArrayApplyRecoveryInfo(m, v39+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v6398 = m.ExcPending
	if v6398 != 0 {
		goto L32
	} else {
		goto L1425
	}
L1424:
	;
	goto L1423
L1425:
	;
	goto L1395
L1426:
	;
	v6443 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v6443+int32(96), int32(_a_F_StartupXLOG_50), int32(1682), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6450 = m.ExcPending
	if v6450 != 0 {
		goto L32
	} else {
		goto L1429
	}
L1427:
	;
	goto L1428
L1428:
	;
	v6452 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	v6454 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v6452) < base.Ui64(v6454) {
		goto L1431
	} else {
		goto L1432
	}
L1429:
	;
	goto L1428
L1430:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6470)+32)) = v6471
	v6474 = *(*int32)(unsafe.Add(mBase, uint32(v6472)))
	v6475 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6470)+64)) = v6475
	*(*int32)(unsafe.Add(mBase, uint32(v6470)+56)) = v6474
	*(*int64)(unsafe.Add(mBase, uint32(v6470)+48)) = v6471
	*(*int32)(unsafe.Add(mBase, uint32(v6470)+40)) = v6474
	*(*int64)(unsafe.Add(mBase, uint32(v6470)+72)) = v6475
	v6482 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6470)+80)) = v6482
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6470)+96)), uint32(v6482))
	v6491 = m.G0
	v6492 = int32(16)
	v6493 = v6491 - v6492
	m.G0 = v6493
	F_gettimeofday(m, v6493)
	mBase = m.M
	v6496 = *(*int64)(unsafe.Add(mBase, uint32(v6493)))
	v6497 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6493)+8)))
	m.G0 = v6493 + v6492
	goto L1434
L1431:
	;
	v6457 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v6457)+24)) = int64(0)
	v6470 = v6457
	v6471 = v6452
	v6472 = int32(_a_F_StartupXLOG_215)
	goto L1430
L1432:
	;
	goto L1433
L1433:
	;
	v6462 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6464 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6465 = *(*int64)(unsafe.Add(mBase, uint32(v6464)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v6462)+24)) = v6465
	v6467 = *(*int64)(unsafe.Add(mBase, uint32(v6464)+40))
	v6470 = v6462
	v6471 = v6467
	v6472 = int32(_a_F_StartupXLOG_216)
	goto L1430
L1434:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[98])) = v6497 + v6496*int64(1000000) - int64(946684800000000)
	v6508 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v6508 == int32(1) {
		goto L1435
	} else {
		goto L1436
	}
L1435:
	;
	v6513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v6513 == int32(1) {
		goto L1439
	} else {
		goto L1440
	}
L1436:
	;
	goto L1437
L1437:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L32
	} else {
		goto L1442
	}
L1438:
	;
	goto L1437
L1439:
	;
	v6517 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[99]))
	*(*int32)(unsafe.Add(mBase, uint32(v6517+int32(0)))) = int32(1)
	v6524 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	v6526 = F_pgmem_kill(m, v6524, int32(10))
	mBase = m.M
	goto L1441
L1440:
	;
	goto L1441
L1441:
	;
	goto L1438
L1442:
	;
	v6530 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v6532 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	v6534 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v6532) < base.Ui64(v6534) {
		goto L1452
	} else {
		goto L1453
	}
L1443:
	;
	goto L1353
L1444:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9184 = m.ExcPending
	if v9184 != 0 {
		goto L32
	} else {
		goto L2011
	}
L1445:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9171 = m.ExcPending
	if v9171 != 0 {
		goto L32
	} else {
		goto L2008
	}
L1446:
	;
	if v9130 != 0 {
		goto L2004
	} else {
		goto L2005
	}
L1447:
	;
	v8997 = int32(0)
	goto L1975
L1448:
	;
	v8923 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[101])))
	if v8923 == int32(0) {
		goto L1445
	} else {
		goto L1963
	}
L1449:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v7123
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = v8829
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = int64(0)
	v8839 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v8839)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v8839)
	v8846 = F_errstart(m, int32(15), v8839)
	mBase = m.M
	v8847 = m.ExcPending
	if v8847 != 0 {
		goto L32
	} else {
		goto L1951
	}
L1450:
	;
	v8815 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8816 = m.ExcPending
	if v8816 != 0 {
		goto L32
	} else {
		goto L1947
	}
L1451:
	;
	F_getrusage(m, v6435+int32(384))
	mBase = m.M
	F_gettimeofday(m, v6435+int32(368))
	mBase = m.M
	goto L1466
L1452:
	;
	v6537 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38]))
	F_XLogPrefetcherBeginRead(m, v6530, v6532)
	mBase = m.M
	v6539 = m.ExcPending
	if v6539 != 0 {
		goto L32
	} else {
		goto L1455
	}
L1453:
	;
	goto L1454
L1454:
	;
	v6575 = int32(0)
	v6579 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v6580 = F_ReadRecord(m, v6530, int32(15), v6575, v6579)
	mBase = m.M
	v6581 = m.ExcPending
	if v6581 != 0 {
		goto L32
	} else {
		goto L1464
	}
L1455:
	;
	v6541 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v6544 = F_ReadRecord(m, v6541, int32(23), int32(0), v6537)
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		goto L32
	} else {
		goto L1456
	}
L1456:
	;
	v6546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6544)+17)))
	if v6546 == int32(0) {
		goto L1457
	} else {
		goto L1458
	}
L1457:
	;
	v6549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6544)+16)))
	if v6549&int32(240) == int32(224) {
		v6584 = v6537
		v6585 = v6544
		goto L1451
	} else {
		goto L1460
	}
L1458:
	;
	goto L1459
L1459:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6557 = m.ExcPending
	if v6557 != 0 {
		goto L32
	} else {
		goto L1461
	}
L1460:
	;
	goto L1459
L1461:
	;
	v6559 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6560 = *(*int64)(unsafe.Add(mBase, uint32(v6559)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+356)) = uint32(v6560)
	v6563 = int64(base.Ui64(v6560) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+352)) = uint32(v6563)
	F_errmsg(m, int32(_a_F_StartupXLOG_217), v6435+int32(352))
	mBase = m.M
	v6569 = m.ExcPending
	if v6569 != 0 {
		goto L32
	} else {
		goto L1462
	}
L1462:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1737), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6574 = m.ExcPending
	if v6574 != 0 {
		goto L32
	} else {
		goto L1463
	}
L1463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1464:
	;
	if v6580 == int32(0) {
		goto L1450
	} else {
		goto L1465
	}
L1465:
	;
	v6584 = v6579
	v6585 = v6580
	goto L1451
L1466:
	;
	v6594 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v6594)
	v6599 = int32(0)
	goto L1467
L1467:
	;
	v6632 = v6599 << (uint(int32(5)) % 32)
	v6635 = *(*int32)(unsafe.Add(mBase, uint32(v6632)+uint32(_c_F_StartupXLOG[108])))
	if v6635 == int32(0) {
		goto L1469
	} else {
		goto L1470
	}
L1468:
	;
	v6659 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6660 = m.ExcPending
	if v6660 != 0 {
		goto L32
	} else {
		goto L1478
	}
L1469:
	;
	v6644 = *(*int32)(unsafe.Add(mBase, uint32(v6632)+uint32(_c_F_StartupXLOG[109])))
	if v6644 == int32(0) {
		goto L1473
	} else {
		goto L1474
	}
L1470:
	;
	v6638 = *(*int32)(unsafe.Add(mBase, uint32(v6632)+uint32(_c_F_StartupXLOG[110])))
	if v6638 == int32(0) {
		goto L1469
	} else {
		goto L1471
	}
L1471:
	;
	m.T0[v6638].(func(*base.Module))(m)
	mBase = m.M
	v6642 = m.ExcPending
	if v6642 != 0 {
		goto L32
	} else {
		goto L1472
	}
L1472:
	;
	goto L1469
L1473:
	;
	v6654 = v6599 + int32(2)
	if v6654 != int32(256) {
		v6599 = v6654
		goto L1467
	} else {
		goto L1477
	}
L1474:
	;
	v6647 = *(*int32)(unsafe.Add(mBase, uint32(v6632)+uint32(_c_F_StartupXLOG[111])))
	if v6647 == int32(0) {
		goto L1473
	} else {
		goto L1475
	}
L1475:
	;
	m.T0[v6647].(func(*base.Module))(m)
	mBase = m.M
	v6651 = m.ExcPending
	if v6651 != 0 {
		goto L32
	} else {
		goto L1476
	}
L1476:
	;
	goto L1473
L1477:
	;
	goto L1468
L1478:
	;
	if v6659 != 0 {
		goto L1479
	} else {
		goto L1480
	}
L1479:
	;
	v6662 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6663 = *(*int64)(unsafe.Add(mBase, uint32(v6662)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+340)) = uint32(v6663)
	v6666 = int64(base.Ui64(v6663) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+336)) = uint32(v6666)
	F_errmsg(m, int32(_a_F_StartupXLOG_218), v6435+int32(336))
	mBase = m.M
	v6672 = m.ExcPending
	if v6672 != 0 {
		goto L32
	} else {
		goto L1482
	}
L1480:
	;
	goto L1481
L1481:
	;
	v6680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v6680 == int32(0) {
		goto L1484
	} else {
		goto L1485
	}
L1482:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1760), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6677 = m.ExcPending
	if v6677 != 0 {
		goto L32
	} else {
		goto L1483
	}
L1483:
	;
	goto L1481
L1484:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v6684 = m.ExcPending
	if v6684 != 0 {
		goto L32
	} else {
		goto L1487
	}
L1485:
	;
	goto L1486
L1486:
	;
	v6685 = v6584
	v6690 = v6585
	goto L1488
L1487:
	;
	goto L1486
L1488:
	;
	v6720 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v6720 != 0 {
		goto L1490
	} else {
		goto L1491
	}
L1489:
	;
	v8966 = v8806
	goto L1447
L1490:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v6788 = m.ExcPending
	if v6788 != 0 {
		goto L32
	} else {
		goto L1501
	}
L1491:
	;
	v6728 = m.G0
	v6730 = v6728 - int32(16)
	m.G0 = v6730
	v6733 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[112]))
	if v6733 != 0 {
		goto L1493
	} else {
		goto L1494
	}
L1492:
	;
	if base.B2i32(v6733 != int32(0)) == int32(0) {
		goto L1490
	} else {
		goto L1496
	}
L1493:
	;
	v6734 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v6736 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[113]))
	F_TimestampDifference(m, v6736, v6734, v6730+int32(12), v6730+int32(8))
	mBase = m.M
	v6742 = *(*int32)(unsafe.Add(mBase, uint32(v6730)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6435+int32(560)))) = v6742
	v6744 = *(*int32)(unsafe.Add(mBase, uint32(v6730)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6435+int32(540)))) = v6744
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[112])) = int32(0)
	goto L1495
L1494:
	;
	goto L1495
L1495:
	;
	m.G0 = v6730 + int32(16)
	goto L1492
L1496:
	;
	v6759 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6760 = m.ExcPending
	if v6760 != 0 {
		goto L32
	} else {
		goto L1497
	}
L1497:
	;
	if v6759 == int32(0) {
		goto L1490
	} else {
		goto L1498
	}
L1498:
	;
	v6764 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6765 = *(*int64)(unsafe.Add(mBase, uint32(v6764)+32))
	v6766 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+320)) = v6766
	v6768 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+540))
	v6770 = base.I32_div_s(v6768, int32(_a_F_StartupXLOG_219))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+324)) = v6770
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+332)) = uint32(v6765)
	v6774 = int64(base.Ui64(v6765) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+328)) = uint32(v6774)
	F_errmsg(m, int32(_a_F_StartupXLOG_220), v6435+int32(320))
	mBase = m.M
	v6780 = m.ExcPending
	if v6780 != 0 {
		goto L32
	} else {
		goto L1499
	}
L1499:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1773), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6785 = m.ExcPending
	if v6785 != 0 {
		goto L32
	} else {
		goto L1500
	}
L1500:
	;
	goto L1490
L1501:
	;
	v6790 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6791 = *(*int32)(unsafe.Add(mBase, uint32(v6790)+80))
	if v6791 != 0 {
		goto L1502
	} else {
		goto L1503
	}
L1502:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v6794 = m.ExcPending
	if v6794 != 0 {
		goto L32
	} else {
		goto L1505
	}
L1503:
	;
	goto L1504
L1504:
	;
	v6796 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v6796 != int32(1) {
		goto L1506
	} else {
		goto L1507
	}
L1505:
	;
	goto L1504
L1506:
	;
	v7191 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114]))
	if v7191 <= int32(0) {
		goto L1591
	} else {
		goto L1592
	}
L1507:
	;
	v6800 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6802 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v6802 != int32(5) {
		goto L1508
	} else {
		goto L1509
	}
L1508:
	;
	if v6802 != int32(4) {
		goto L1517
	} else {
		goto L1518
	}
L1509:
	;
	v6806 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[101])))
	if v6806&int32(1) == int32(0) {
		goto L1508
	} else {
		goto L1510
	}
L1510:
	;
	v6813 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6814 = m.ExcPending
	if v6814 != 0 {
		goto L32
	} else {
		goto L1511
	}
L1511:
	;
	if v6813 != 0 {
		goto L1512
	} else {
		goto L1513
	}
L1512:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_221), int32(0))
	mBase = m.M
	v6818 = m.ExcPending
	if v6818 != 0 {
		goto L32
	} else {
		goto L1515
	}
L1513:
	;
	goto L1514
L1514:
	;
	v6825 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v6825
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = v6825
	v6831 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v6831
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v6831)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v6831)
	goto L1448
L1515:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2614), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v6823 = m.ExcPending
	if v6823 != 0 {
		goto L32
	} else {
		goto L1516
	}
L1516:
	;
	goto L1514
L1517:
	;
	v6886 = *(*int32)(unsafe.Add(mBase, uint32(v6800)+96))
	v6887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6886)+49)))
	if v6887 != int32(1) {
		goto L1506
	} else {
		goto L1525
	}
L1518:
	;
	v6842 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[115])))
	if v6842&int32(1) != 0 {
		goto L1517
	} else {
		goto L1519
	}
L1519:
	;
	v6845 = *(*int64)(unsafe.Add(mBase, uint32(v6800)+32))
	v6847 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	if base.Ui64(v6845) < base.Ui64(v6847) {
		goto L1517
	} else {
		goto L1520
	}
L1520:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v6845
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = int64(0)
	v6855 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v6855
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v6855)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v6855)
	v6865 = F_errstart(m, int32(15), v6855)
	mBase = m.M
	v6866 = m.ExcPending
	if v6866 != 0 {
		goto L32
	} else {
		goto L1521
	}
L1521:
	;
	if v6865 == int32(0) {
		goto L1448
	} else {
		goto L1522
	}
L1522:
	;
	v6870 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+308)) = uint32(v6870)
	v6873 = int64(base.Ui64(v6870) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+304)) = uint32(v6873)
	F_errmsg(m, int32(_a_F_StartupXLOG_223), v6435+int32(304))
	mBase = m.M
	v6879 = m.ExcPending
	if v6879 != 0 {
		goto L32
	} else {
		goto L1523
	}
L1523:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2636), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v6884 = m.ExcPending
	if v6884 != 0 {
		goto L32
	} else {
		goto L1524
	}
L1524:
	;
	goto L1448
L1525:
	;
	v6890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6886)+48)))
	switch int32(base.Ui32(v6890)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		goto L1531
	default:
		goto L1506
	case 2:
		goto L1529
	case 3:
		goto L1530
	case 4:
		goto L1528
	}
L1526:
	;
	v7125 = int32(0)
	v7127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[115])))
	v7128 = int32(1)
	v7131 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v7127&v7128|base.B2i32(v7131 != v7128) == v7125 {
		goto L1573
	} else {
		goto L1574
	}
L1527:
	;
	v7123 = v7121
	v7124 = int32(0)
	goto L1526
L1528:
	;
	v7017 = *(*int32)(unsafe.Add(mBase, uint32(v6886)+64))
	v7019 = v6435 + int32(560)
	v7020 = int32(0)
	base.MemoryFill(m, v7019, v7020, int32(264))
	v7026 = *(*int64)(unsafe.Add(mBase, uint32(v7017)))
	*(*int64)(unsafe.Add(mBase, uint32(v7019))) = v7026
	if v7020 <= base.I32_extend8_s(v6890) {
		goto L1555
	} else {
		goto L1556
	}
L1529:
	;
	v7016 = *(*int32)(unsafe.Add(mBase, uint32(v6886)+36))
	v7121 = v7016
	goto L1527
L1530:
	;
	v6897 = *(*int32)(unsafe.Add(mBase, uint32(v6886)+64))
	v6899 = v6435 + int32(560)
	v6900 = int32(0)
	base.MemoryFill(m, v6899, v6900, int32(288))
	v6906 = *(*int64)(unsafe.Add(mBase, uint32(v6897)))
	*(*int64)(unsafe.Add(mBase, uint32(v6899))) = v6906
	if v6900 <= base.I32_extend8_s(v6890) {
		goto L1533
	} else {
		goto L1534
	}
L1531:
	;
	v6895 = *(*int32)(unsafe.Add(mBase, uint32(v6886)+36))
	v7123 = v6895
	v7124 = int32(1)
	goto L1526
L1532:
	;
	v7014 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+612))
	v7123 = v7014
	v7124 = int32(1)
	goto L1526
L1533:
	;
	goto L1532
L1534:
	;
	v6911 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+8)) = v6911
	if v6911&int32(1) != 0 {
		goto L1535
	} else {
		goto L1536
	}
L1535:
	;
	v6915 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+12)) = v6915
	v6917 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+16)) = v6917
	v6923 = v6897 + int32(20)
	goto L1537
L1536:
	;
	v6923 = v6897 + int32(12)
	goto L1537
L1537:
	;
	if v6911&int32(2) != 0 {
		goto L1538
	} else {
		goto L1539
	}
L1538:
	;
	v6926 = *(*int32)(unsafe.Add(mBase, uint32(v6923)))
	v6928 = v6923 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+24)) = v6928
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+20)) = v6926
	v6934 = v6928 + v6926<<(uint(int32(2))%32)
	goto L1540
L1539:
	;
	v6934 = v6923
	goto L1540
L1540:
	;
	if v6911&int32(4) != 0 {
		goto L1541
	} else {
		goto L1542
	}
L1541:
	;
	v6938 = *(*int32)(unsafe.Add(mBase, uint32(v6934)))
	v6940 = v6934 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+32)) = v6940
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+28)) = v6938
	v6943 = *(*int32)(unsafe.Add(mBase, uint32(v6934)))
	v6947 = v6940 + v6943*int32(12)
	goto L1543
L1542:
	;
	v6947 = v6934
	goto L1543
L1543:
	;
	if v6911&int32(256) != 0 {
		goto L1544
	} else {
		goto L1545
	}
L1544:
	;
	v6952 = *(*int32)(unsafe.Add(mBase, uint32(v6947)))
	v6953 = int32(4)
	v6954 = v6947 + v6953
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+40)) = v6954
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+36)) = v6952
	v6957 = *(*int32)(unsafe.Add(mBase, uint32(v6947)))
	v6961 = v6954 + v6957<<(uint(v6953)%32)
	goto L1546
L1545:
	;
	v6961 = v6947
	goto L1546
L1546:
	;
	if v6911&int32(8) != 0 {
		goto L1547
	} else {
		goto L1548
	}
L1547:
	;
	v6966 = *(*int32)(unsafe.Add(mBase, uint32(v6961)))
	v6967 = int32(4)
	v6968 = v6961 + v6967
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+48)) = v6968
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+44)) = v6966
	v6971 = *(*int32)(unsafe.Add(mBase, uint32(v6961)))
	v6975 = v6968 + v6971<<(uint(v6967)%32)
	goto L1549
L1548:
	;
	v6975 = v6961
	goto L1549
L1549:
	;
	if v6911&int32(16) == int32(0) {
		v6999 = v6911
		v7000 = v6975
		goto L1550
	} else {
		goto L1551
	}
L1550:
	;
	if v6999&int32(32) == int32(0) {
		goto L1533
	} else {
		goto L1553
	}
L1551:
	;
	v6982 = *(*int32)(unsafe.Add(mBase, uint32(v6975)))
	*(*int32)(unsafe.Add(mBase, uint32(v6899)+52)) = v6982
	v6985 = v6975 + int32(4)
	if v6911&int32(128) == int32(0) {
		v6999 = v6911
		v7000 = v6985
		goto L1550
	} else {
		goto L1552
	}
L1552:
	;
	v6993 = F_strlcpy(m, v6435+int32(616), v6985, int32(200))
	mBase = m.M
	v6994 = F_strlen(m, v6985)
	mBase = m.M
	v6998 = *(*int32)(unsafe.Add(mBase, uint32(v6899)+8))
	v6999 = v6998
	v7000 = v6994 + v6985 + int32(1)
	goto L1550
L1553:
	;
	v7005 = *(*int64)(unsafe.Add(mBase, uint32(v7000)))
	v7006 = *(*int64)(unsafe.Add(mBase, uint32(v7000)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v6899)+280)) = v7006
	*(*int64)(unsafe.Add(mBase, uint32(v6899)+272)) = v7005
	goto L1533
L1554:
	;
	v7120 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+604))
	v7121 = v7120
	goto L1527
L1555:
	;
	goto L1554
L1556:
	;
	v7031 = *(*int32)(unsafe.Add(mBase, uint32(v7017)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+8)) = v7031
	if v7031&int32(1) != 0 {
		goto L1557
	} else {
		goto L1558
	}
L1557:
	;
	v7035 = *(*int32)(unsafe.Add(mBase, uint32(v7017)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+12)) = v7035
	v7037 = *(*int32)(unsafe.Add(mBase, uint32(v7017)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+16)) = v7037
	v7043 = v7017 + int32(20)
	goto L1559
L1558:
	;
	v7043 = v7017 + int32(12)
	goto L1559
L1559:
	;
	if v7031&int32(2) != 0 {
		goto L1560
	} else {
		goto L1561
	}
L1560:
	;
	v7046 = *(*int32)(unsafe.Add(mBase, uint32(v7043)))
	v7048 = v7043 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+24)) = v7048
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+20)) = v7046
	v7054 = v7048 + v7046<<(uint(int32(2))%32)
	goto L1562
L1561:
	;
	v7054 = v7043
	goto L1562
L1562:
	;
	if v7031&int32(4) != 0 {
		goto L1563
	} else {
		goto L1564
	}
L1563:
	;
	v7058 = *(*int32)(unsafe.Add(mBase, uint32(v7054)))
	v7060 = v7054 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+32)) = v7060
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+28)) = v7058
	v7063 = *(*int32)(unsafe.Add(mBase, uint32(v7054)))
	v7067 = v7060 + v7063*int32(12)
	goto L1565
L1564:
	;
	v7067 = v7054
	goto L1565
L1565:
	;
	if v7031&int32(256) != 0 {
		goto L1566
	} else {
		goto L1567
	}
L1566:
	;
	v7072 = *(*int32)(unsafe.Add(mBase, uint32(v7067)))
	v7073 = int32(4)
	v7074 = v7067 + v7073
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+40)) = v7074
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+36)) = v7072
	v7077 = *(*int32)(unsafe.Add(mBase, uint32(v7067)))
	v7081 = v7074 + v7077<<(uint(v7073)%32)
	goto L1568
L1567:
	;
	v7081 = v7067
	goto L1568
L1568:
	;
	if v7031&int32(16) == int32(0) {
		v7105 = v7031
		v7106 = v7081
		goto L1569
	} else {
		goto L1570
	}
L1569:
	;
	if v7105&int32(32) == int32(0) {
		goto L1555
	} else {
		goto L1572
	}
L1570:
	;
	v7088 = *(*int32)(unsafe.Add(mBase, uint32(v7081)))
	*(*int32)(unsafe.Add(mBase, uint32(v7019)+44)) = v7088
	v7091 = v7081 + int32(4)
	if v7031&int32(128) == int32(0) {
		v7105 = v7031
		v7106 = v7091
		goto L1569
	} else {
		goto L1571
	}
L1571:
	;
	v7099 = F_strlcpy(m, v6435+int32(608), v7091, int32(200))
	mBase = m.M
	v7100 = F_strlen(m, v7091)
	mBase = m.M
	v7104 = *(*int32)(unsafe.Add(mBase, uint32(v7019)+8))
	v7105 = v7104
	v7106 = v7100 + v7091 + int32(1)
	goto L1569
L1572:
	;
	v7111 = *(*int64)(unsafe.Add(mBase, uint32(v7106)))
	v7112 = *(*int64)(unsafe.Add(mBase, uint32(v7106)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7019)+256)) = v7112
	*(*int64)(unsafe.Add(mBase, uint32(v7019)+248)) = v7111
	goto L1555
L1573:
	;
	v7138 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v7140 = base.B2i32(v7123 == v7138)
	goto L1575
L1574:
	;
	v7140 = v7125
	goto L1575
L1575:
	;
	v7141 = *(*int32)(unsafe.Add(mBase, uint32(v6800)+96))
	v7142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7141)+48)))
	v7143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7141)+49)))
	if base.B2i32(v7143 == int32(0))&base.B2i32(v7142&int32(240) == int32(112)) != 0 {
		goto L1576
	} else {
		goto L1577
	}
L1576:
	;
	v7166 = *(*int32)(unsafe.Add(mBase, uint32(v7141)+64))
	v7167 = *(*int64)(unsafe.Add(mBase, uint32(v7166)))
	if v7131 == int32(2) {
		goto L1584
	} else {
		goto L1585
	}
L1577:
	;
	if v7143 != int32(1) {
		goto L1578
	} else {
		goto L1579
	}
L1578:
	;
	if v7140 == int32(0) {
		goto L1506
	} else {
		goto L1582
	}
L1579:
	;
	v7153 = int32(4)
	v7156 = int32(base.Ui32(v7142)>>(uint(v7153)%32)) & int32(7)
	if base.Ui32(v7153) < base.Ui32(v7156) {
		goto L1578
	} else {
		goto L1580
	}
L1580:
	;
	if v7156 != int32(1) {
		goto L1576
	} else {
		goto L1581
	}
L1581:
	;
	goto L1578
L1582:
	;
	v8829 = int64(0)
	goto L1449
L1583:
	;
	if v7171 <= v7167 {
		v8829 = v7167
		goto L1449
	} else {
		goto L1590
	}
L1584:
	;
	v7171 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25]))
	if v7127&int32(1) == int32(0) {
		goto L1583
	} else {
		goto L1587
	}
L1585:
	;
	goto L1586
L1586:
	;
	if v7140 == int32(0) {
		goto L1506
	} else {
		goto L1589
	}
L1587:
	;
	if v7171 < v7167 {
		v8829 = v7167
		goto L1449
	} else {
		goto L1588
	}
L1588:
	;
	goto L1506
L1589:
	;
	v8829 = v7167
	goto L1449
L1590:
	;
	goto L1506
L1591:
	;
	v7478 = int32(0)
	v7479 = int32(_a_F_StartupXLOG_224)
	v7480 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[116]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[116])) = v6435 + int32(540)
	v7486 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+548)) = v7486
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+544)) = int32(413)
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+540)) = v7480
	v7491 = *(*int32)(unsafe.Add(mBase, uint32(v6690)+4))
	F_AdvanceNextFullTransactionIdPastXid(m, v7491)
	mBase = m.M
	v7493 = m.ExcPending
	if v7493 != 0 {
		goto L32
	} else {
		goto L1633
	}
L1592:
	;
	v7195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[101])))
	if v7195&int32(1) == int32(0) {
		goto L1591
	} else {
		goto L1593
	}
L1593:
	;
	v7201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v7201&int32(1) == int32(0) {
		goto L1591
	} else {
		goto L1594
	}
L1594:
	;
	v7207 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7208 = *(*int32)(unsafe.Add(mBase, uint32(v7207)+96))
	v7209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7208)+49)))
	if v7209 != int32(1) {
		goto L1591
	} else {
		goto L1595
	}
L1595:
	;
	v7212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7208)+48)))
	v7214 = v7212 & int32(112)
	if v7214 != 0 {
		goto L1596
	} else {
		goto L1597
	}
L1596:
	;
	v7218 = base.B2i32(v7214 != int32(48))
	goto L1598
L1597:
	;
	v7218 = int32(0)
	goto L1598
L1598:
	;
	if v7218 != 0 {
		goto L1591
	} else {
		goto L1599
	}
L1599:
	;
	v7219 = int32(4)
	v7222 = int32(base.Ui32(v7212)>>(uint(v7219)%32)) & int32(7)
	if base.B2i32(base.Ui32(v7219) < base.Ui32(v7222))|base.B2i32(v7222 == int32(1)) != 0 {
		goto L1591
	} else {
		goto L1600
	}
L1600:
	;
	v7228 = *(*int32)(unsafe.Add(mBase, uint32(v7208)+64))
	v7229 = *(*int64)(unsafe.Add(mBase, uint32(v7228)))
	v7233 = m.G0
	v7234 = int32(16)
	v7235 = v7233 - v7234
	m.G0 = v7235
	F_gettimeofday(m, v7235)
	mBase = m.M
	v7238 = *(*int64)(unsafe.Add(mBase, uint32(v7235)))
	v7239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7235)+8)))
	m.G0 = v7235 + v7234
	v7247 = v7239 + v7238*int64(1000000) - int64(946684800000000)
	goto L1601
L1601:
	;
	v7251 = v7229 + base.I64_extend_i32_u(v7191)*int64(1000)
	if v7251 <= v7247 {
		v7269 = int32(0)
		goto L1603
	} else {
		goto L1604
	}
L1602:
	;
	if v7269 <= int32(0) {
		goto L1591
	} else {
		goto L1606
	}
L1603:
	;
	goto L1602
L1604:
	;
	v7257 = v7251 - v7247
	if base.B2i32(int64(0) < v7247)^base.B2i32(v7257 < v7251)|base.B2i32(int64(2147483646000) < v7257) != 0 {
		v7269 = int32(2147483647)
		goto L1603
	} else {
		goto L1605
	}
L1605:
	;
	v7266 = base.I64_div_s(v7257+int64(999), int64(1000))
	v7269 = base.I32_wrap_i64(v7266)
	goto L1603
L1606:
	;
	v7273 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7273+int32(4)))) = v7276
	v7281 = base.AtomicRmwOr32(m, v7276, int32(_a_F_StartupXLOG_225), v7276)
	goto L1607
L1607:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7283 = m.ExcPending
	if v7283 != 0 {
		goto L32
	} else {
		goto L1608
	}
L1608:
	;
	v7284 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7285 = m.ExcPending
	if v7285 != 0 {
		goto L32
	} else {
		goto L1610
	}
L1609:
	;
	v7437 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7438 = *(*int32)(unsafe.Add(mBase, uint32(v7437)+80))
	if v7438 == int32(0) {
		goto L1591
	} else {
		goto L1631
	}
L1610:
	;
	if v7284 != 0 {
		goto L1609
	} else {
		goto L1611
	}
L1611:
	;
	goto L1612
L1612:
	;
	v7321 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114])))
	v7325 = m.G0
	v7326 = int32(16)
	v7327 = v7325 - v7326
	m.G0 = v7327
	F_gettimeofday(m, v7327)
	mBase = m.M
	v7330 = *(*int64)(unsafe.Add(mBase, uint32(v7327)))
	v7331 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7327)+8)))
	m.G0 = v7327 + v7326
	v7339 = v7331 + v7330*int64(1000000) - int64(946684800000000)
	goto L1614
L1613:
	;
	goto L1609
L1614:
	;
	v7342 = v7321*int64(1000) + v7229
	if v7342 <= v7339 {
		v7360 = int32(0)
		goto L1616
	} else {
		goto L1617
	}
L1615:
	;
	if v7360 <= int32(0) {
		goto L1609
	} else {
		goto L1619
	}
L1616:
	;
	goto L1615
L1617:
	;
	v7348 = v7342 - v7339
	if base.B2i32(int64(0) < v7339)^base.B2i32(v7348 < v7342)|base.B2i32(int64(2147483646000) < v7348) != 0 {
		v7360 = int32(2147483647)
		goto L1616
	} else {
		goto L1618
	}
L1618:
	;
	v7357 = base.I64_div_s(v7348+int64(999), int64(1000))
	v7360 = base.I32_wrap_i64(v7357)
	goto L1616
L1619:
	;
	v7365 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v7366 = m.ExcPending
	if v7366 != 0 {
		goto L32
	} else {
		goto L1620
	}
L1620:
	;
	if v7365 != 0 {
		goto L1621
	} else {
		goto L1622
	}
L1621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+256)) = v7360
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_226), v6435+int32(256))
	mBase = m.M
	v7372 = m.ExcPending
	if v7372 != 0 {
		goto L32
	} else {
		goto L1624
	}
L1622:
	;
	goto L1623
L1623:
	;
	v7379 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7384 = F_WaitLatch(m, v7379+int32(4), int32(41), v7360, int32(150994947))
	mBase = m.M
	v7385 = m.ExcPending
	if v7385 != 0 {
		goto L32
	} else {
		goto L1626
	}
L1624:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(3078), int32(_a_F_StartupXLOG_227))
	mBase = m.M
	v7377 = m.ExcPending
	if v7377 != 0 {
		goto L32
	} else {
		goto L1625
	}
L1625:
	;
	goto L1623
L1626:
	;
	v7387 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7390 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7387+int32(4)))) = v7390
	v7395 = base.AtomicRmwOr32(m, v7390, int32(_a_F_StartupXLOG_225), v7390)
	goto L1627
L1627:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7397 = m.ExcPending
	if v7397 != 0 {
		goto L32
	} else {
		goto L1628
	}
L1628:
	;
	v7398 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7399 = m.ExcPending
	if v7399 != 0 {
		goto L32
	} else {
		goto L1629
	}
L1629:
	;
	if v7398 == int32(0) {
		goto L1612
	} else {
		goto L1630
	}
L1630:
	;
	goto L1613
L1631:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v7443 = m.ExcPending
	if v7443 != 0 {
		goto L32
	} else {
		goto L1632
	}
L1632:
	;
	goto L1591
L1633:
	;
	v7494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6690)+17)))
	if v7494 != 0 {
		v7561 = v6685
		v7564 = v7478
		goto L1641
	} else {
		goto L1642
	}
L1634:
	;
	v8806 = int32(0)
	v8808 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v8811 = F_ReadRecord(m, v8808, int32(15), v8806, v7561)
	mBase = m.M
	v8812 = m.ExcPending
	if v8812 != 0 {
		goto L32
	} else {
		goto L1945
	}
L1635:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8780 = m.ExcPending
	if v8780 != 0 {
		goto L32
	} else {
		goto L1942
	}
L1636:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8761 = m.ExcPending
	if v8761 != 0 {
		goto L32
	} else {
		goto L1938
	}
L1637:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8738 = m.ExcPending
	if v8738 != 0 {
		goto L32
	} else {
		goto L1935
	}
L1638:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8714 = m.ExcPending
	if v8714 != 0 {
		goto L32
	} else {
		goto L1932
	}
L1639:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8698 = m.ExcPending
	if v8698 != 0 {
		goto L32
	} else {
		goto L1929
	}
L1640:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8681 = m.ExcPending
	if v8681 != 0 {
		goto L32
	} else {
		goto L1926
	}
L1641:
	;
	v7567 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7570 = base.AtomicRmwXchg32(m, v7567, int32(96), int32(1))
	if v7570 != 0 {
		goto L1668
	} else {
		goto L1669
	}
L1642:
	;
	v7495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6690)+16)))
	v7497 = v7495 & int32(240)
	if v7497 != 0 {
		goto L1643
	} else {
		goto L1644
	}
L1643:
	;
	v7501 = base.B2i32(v7497 != int32(144))
	goto L1645
L1644:
	;
	v7501 = int32(0)
	goto L1645
L1645:
	;
	if v7501 != 0 {
		v7561 = v6685
		v7564 = v7478
		goto L1641
	} else {
		goto L1646
	}
L1646:
	;
	v7502 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7503 = *(*int32)(unsafe.Add(mBase, uint32(v7502)+64))
	v7504 = *(*int32)(unsafe.Add(mBase, uint32(v7503)+8))
	if v7504 == v6685 {
		v7561 = v6685
		v7564 = v7478
		goto L1641
	} else {
		goto L1647
	}
L1647:
	;
	v7506 = *(*int32)(unsafe.Add(mBase, uint32(v7503)+12))
	if v7506 != v6685 {
		goto L1640
	} else {
		goto L1648
	}
L1648:
	;
	if base.Ui32(v7504) < base.Ui32(v6685) {
		goto L1639
	} else {
		goto L1649
	}
L1649:
	;
	v7509 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	v7511 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v7512 = int32(0)
	if v7511 == v7512 {
		goto L1651
	} else {
		goto L1652
	}
L1650:
	;
	if v7551 == int32(0) {
		goto L1639
	} else {
		goto L1663
	}
L1651:
	;
	v7551 = int32(0)
	goto L1650
L1652:
	;
	goto L1653
L1653:
	;
	v7518 = *(*int32)(unsafe.Add(mBase, uint32(v7511)+4))
	if v7518 <= int32(0) {
		v7545 = v7512
		goto L1654
	} else {
		goto L1655
	}
L1654:
	;
	v7551 = v7545
	goto L1650
L1655:
	;
	v7521 = int32(0)
	if v7521 < v7518 {
		goto L1656
	} else {
		goto L1657
	}
L1656:
	;
	v7524 = v7518
	goto L1658
L1657:
	;
	v7524 = v7521
	goto L1658
L1658:
	;
	v7525 = *(*int32)(unsafe.Add(mBase, uint32(v7511)+12))
	v7528 = int32(0)
	goto L1659
L1659:
	;
	v7535 = *(*int32)(unsafe.Add(mBase, uint32(v7525+v7528<<(uint(int32(2))%32))))
	v7536 = *(*int32)(unsafe.Add(mBase, uint32(v7535)))
	v7537 = base.B2i32(v7536 == v7504)
	if v7536 == v7504 {
		v7545 = v7537
		goto L1654
	} else {
		goto L1661
	}
L1660:
	;
	v7545 = v7537
	goto L1654
L1661:
	;
	v7539 = v7528 + int32(1)
	if v7539 != v7524 {
		v7528 = v7539
		goto L1659
	} else {
		goto L1662
	}
L1662:
	;
	goto L1660
L1663:
	;
	v7556 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	if base.Ui64(v7509) < base.Ui64(v7556) {
		goto L1664
	} else {
		goto L1665
	}
L1664:
	;
	v7559 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	if base.Ui32(v7559) < base.Ui32(v7504) {
		goto L1638
	} else {
		goto L1667
	}
L1665:
	;
	goto L1666
L1666:
	;
	v7561 = v7504
	v7564 = int32(1)
	goto L1641
L1667:
	;
	goto L1666
L1668:
	;
	v7572 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v7572+int32(96), int32(_a_F_StartupXLOG_50), int32(1991), int32(_a_F_StartupXLOG_228))
	mBase = m.M
	v7579 = m.ExcPending
	if v7579 != 0 {
		goto L32
	} else {
		goto L1671
	}
L1669:
	;
	goto L1670
L1670:
	;
	v7580 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	v7582 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v7582)+56)) = v7561
	*(*int64)(unsafe.Add(mBase, uint32(v7582)+48)) = v7580
	v7585 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7582)+96)), uint32(v7585))
	v7589 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v7589 == v7585 {
		goto L1672
	} else {
		goto L1673
	}
L1671:
	;
	goto L1670
L1672:
	;
	v7598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6690)+17)))
	if v7598 != 0 {
		goto L1676
	} else {
		goto L1677
	}
L1673:
	;
	v7592 = *(*int32)(unsafe.Add(mBase, uint32(v6690)+4))
	if v7592 == int32(0) {
		goto L1672
	} else {
		goto L1674
	}
L1674:
	;
	F_RecordKnownAssignedTransactionIds(m, v7592)
	mBase = m.M
	v7596 = m.ExcPending
	if v7596 != 0 {
		goto L32
	} else {
		goto L1675
	}
L1675:
	;
	goto L1672
L1676:
	;
	v7689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6690)+17)))
	v7691 = v7689 << (uint(int32(5)) % 32)
	v7694 = *(*int32)(unsafe.Add(mBase, uint32(v7691)+uint32(_c_F_StartupXLOG[108])))
	if v7694 == int32(0) {
		goto L1702
	} else {
		goto L1703
	}
L1677:
	;
	v7599 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7599)+48)))
	v7602 = v7600 & int32(240)
	if v7602 != int32(80) {
		goto L1678
	} else {
		goto L1679
	}
L1678:
	;
	if v7602 != int32(208) {
		goto L1676
	} else {
		goto L1681
	}
L1679:
	;
	goto L1680
L1680:
	;
	v7641 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	v7643 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53]))
	v7644 = *(*int32)(unsafe.Add(mBase, uint32(v7599)+64))
	v7645 = *(*int64)(unsafe.Add(mBase, uint32(v7644)))
	v7648 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v7649 = m.ExcPending
	if v7649 != 0 {
		goto L32
	} else {
		goto L1690
	}
L1681:
	;
	v7607 = *(*int32)(unsafe.Add(mBase, uint32(v7599)+64))
	v7608 = *(*int64)(unsafe.Add(mBase, uint32(v7607)))
	v7609 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+64))
	if v7608 != v7609 {
		goto L1637
	} else {
		goto L1682
	}
L1682:
	;
	v7611 = *(*int64)(unsafe.Add(mBase, uint32(v7607)+8))
	v7613 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v7613
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51])) = v7613
	v7620 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7621 = m.ExcPending
	if v7621 != 0 {
		goto L32
	} else {
		goto L1683
	}
L1683:
	;
	if v7620 != 0 {
		goto L1684
	} else {
		goto L1685
	}
L1684:
	;
	v7622 = F_timestamptz_to_str(m, v7611)
	mBase = m.M
	v7623 = m.ExcPending
	if v7623 != 0 {
		goto L32
	} else {
		goto L1687
	}
L1685:
	;
	goto L1686
L1686:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7486)+64)) = int64(0)
	goto L1676
L1687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+168)) = v7622
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+164)) = uint32(v7608)
	v7627 = int64(base.Ui64(v7608) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+160)) = uint32(v7627)
	F_errmsg(m, int32(_a_F_StartupXLOG_229), v6435+int32(160))
	mBase = m.M
	v7633 = m.ExcPending
	if v7633 != 0 {
		goto L32
	} else {
		goto L1688
	}
L1688:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2117), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7638 = m.ExcPending
	if v7638 != 0 {
		goto L32
	} else {
		goto L1689
	}
L1689:
	;
	goto L1686
L1690:
	;
	if v7645 == v7643 {
		goto L1691
	} else {
		goto L1692
	}
L1691:
	;
	if v7648 != 0 {
		goto L1694
	} else {
		goto L1695
	}
L1692:
	;
	goto L1693
L1693:
	;
	if v7648 == int32(0) {
		goto L1676
	} else {
		goto L1699
	}
L1694:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_231), int32(0))
	mBase = m.M
	v7654 = m.ExcPending
	if v7654 != 0 {
		goto L32
	} else {
		goto L1697
	}
L1695:
	;
	goto L1696
L1696:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v7641
	goto L1676
L1697:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2138), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7659 = m.ExcPending
	if v7659 != 0 {
		goto L32
	} else {
		goto L1698
	}
L1698:
	;
	goto L1696
L1699:
	;
	v7665 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+204)) = uint32(v7665)
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+196)) = uint32(v7645)
	v7668 = int64(32)
	v7669 = int64(base.Ui64(v7645) >> (uint(v7668) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+192)) = uint32(v7669)
	v7672 = int64(base.Ui64(v7665) >> (uint(v7668) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+200)) = uint32(v7672)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_232), v6435+int32(192))
	mBase = m.M
	v7678 = m.ExcPending
	if v7678 != 0 {
		goto L32
	} else {
		goto L1700
	}
L1700:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2144), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v7683 = m.ExcPending
	if v7683 != 0 {
		goto L32
	} else {
		goto L1701
	}
L1701:
	;
	goto L1676
L1702:
	;
	F_RmgrNotFound(m, v7689)
	mBase = m.M
	v7698 = m.ExcPending
	if v7698 != 0 {
		goto L32
	} else {
		goto L1705
	}
L1703:
	;
	goto L1704
L1704:
	;
	v7699 = *(*int32)(unsafe.Add(mBase, uint32(v7691)+uint32(_c_F_StartupXLOG[117])))
	m.T0[v7699].(func(*base.Module, int32))(m, v7486)
	mBase = m.M
	v7701 = m.ExcPending
	if v7701 != 0 {
		goto L32
	} else {
		goto L1706
	}
L1705:
	;
	goto L1704
L1706:
	;
	v7702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6690)+16)))
	if v7702&int32(2) == int32(0) {
		goto L1707
	} else {
		goto L1708
	}
L1707:
	;
	v7974 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+540))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[116])) = v7974
	v7977 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7980 = base.AtomicRmwXchg32(m, v7977, int32(96), int32(1))
	if v7980 != 0 {
		goto L1769
	} else {
		goto L1770
	}
L1708:
	;
	v7707 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7707)+49)))
	v7710 = v7708 << (uint(int32(5)) % 32)
	v7711 = *(*int32)(unsafe.Add(mBase, uint32(v7710)+uint32(_c_F_StartupXLOG[108])))
	if v7711 != 0 {
		goto L1709
	} else {
		goto L1710
	}
L1709:
	;
	v7715 = v7707
	goto L1711
L1710:
	;
	F_RmgrNotFound(m, v7708)
	mBase = m.M
	v7713 = m.ExcPending
	if v7713 != 0 {
		goto L32
	} else {
		goto L1712
	}
L1711:
	;
	v7716 = *(*int32)(unsafe.Add(mBase, uint32(v7715)+72))
	if v7716 < int32(0) {
		goto L1707
	} else {
		goto L1713
	}
L1712:
	;
	v7714 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7715 = v7714
	goto L1711
L1713:
	;
	v7721 = *(*int32)(unsafe.Add(mBase, uint32(v7710)+uint32(_c_F_StartupXLOG[118])))
	v7728 = int32(0)
	goto L1714
L1714:
	;
	v7758 = v7728 & int32(255)
	v7760 = v6435 + int32(560)
	v7762 = v6435 + int32(556)
	v7764 = v6435 + int32(552)
	v7765 = int32(0)
	v7767 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7768 = *(*int32)(unsafe.Add(mBase, uint32(v7767)+72))
	if v7768 < v7758 {
		v7792 = v7765
		goto L1718
	} else {
		goto L1719
	}
L1715:
	;
	goto L1707
L1716:
	;
	v7935 = v7728 + int32(1)
	v7936 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7937 = *(*int32)(unsafe.Add(mBase, uint32(v7936)+72))
	if v7935 <= v7937 {
		v7728 = v7935
		goto L1714
	} else {
		goto L1768
	}
L1717:
	;
	if v7792 == int32(0) {
		goto L1716
	} else {
		goto L1731
	}
L1718:
	;
	goto L1717
L1719:
	;
	v7772 = v7767 + v7758*int32(52)
	v7773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7772)+76)))
	if v7773 != int32(1) {
		v7792 = v7765
		goto L1718
	} else {
		goto L1720
	}
L1720:
	;
	v7777 = v7772 + int32(76)
	if v7760 != 0 {
		goto L1721
	} else {
		goto L1722
	}
L1721:
	;
	v7778 = *(*int32)(unsafe.Add(mBase, uint32(v7777)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7760)+8)) = v7778
	v7780 = *(*int64)(unsafe.Add(mBase, uint32(v7777)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v7760))) = v7780
	goto L1723
L1722:
	;
	goto L1723
L1723:
	;
	if v7762 != 0 {
		goto L1724
	} else {
		goto L1725
	}
L1724:
	;
	v7782 = *(*int32)(unsafe.Add(mBase, uint32(v7777)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7762))) = v7782
	goto L1726
L1725:
	;
	goto L1726
L1726:
	;
	if v7764 != 0 {
		goto L1727
	} else {
		goto L1728
	}
L1727:
	;
	v7784 = *(*int32)(unsafe.Add(mBase, uint32(v7777)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7764))) = v7784
	goto L1729
L1728:
	;
	goto L1729
L1729:
	;
	v7792 = int32(1)
	goto L1718
L1731:
	;
	v7795 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+96))
	v7799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7795+v7728*int32(52))+106)))
	if v7799 != 0 {
		goto L1716
	} else {
		goto L1732
	}
L1732:
	;
	v7800 = *(*int64)(unsafe.Add(mBase, uint32(v6435)+560))
	*(*int64)(unsafe.Add(mBase, uint32(v6435)+144)) = v7800
	v7802 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+568))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+152)) = v7802
	v7806 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+556))
	v7807 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+552))
	v7810 = F_XLogReadBufferExtended(m, v6435+int32(144), v7806, v7807, int32(4), int32(0))
	mBase = m.M
	v7811 = m.ExcPending
	if v7811 != 0 {
		goto L32
	} else {
		goto L1733
	}
L1733:
	;
	if v7810 == int32(0) {
		goto L1716
	} else {
		goto L1734
	}
L1734:
	;
	F_LockBuffer(m, v7810, int32(2))
	mBase = m.M
	v7816 = m.ExcPending
	if v7816 != 0 {
		goto L32
	} else {
		goto L1735
	}
L1735:
	;
	v7818 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	if v7810 < int32(0) {
		goto L1737
	} else {
		goto L1738
	}
L1736:
	;
	base.MemoryCopy(m, v7818, v7836, int32(_a_F_StartupXLOG_58))
	F_UnlockReleaseBuffer(m, v7810)
	mBase = m.M
	v7840 = m.ExcPending
	if v7840 != 0 {
		goto L32
	} else {
		goto L1740
	}
L1737:
	;
	v7822 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[119]))
	v7828 = *(*int32)(unsafe.Add(mBase, uint32(v7822+(v7810^int32(-1))<<(uint(int32(2))%32))))
	v7836 = v7828
	goto L1736
L1738:
	;
	goto L1739
L1739:
	;
	v7830 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[120]))
	v7836 = v7830 + v7810<<(uint(int32(13))%32) + int32(-8192)
	goto L1736
L1740:
	;
	v7841 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	v7843 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7844 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7843))))
	v7847 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7843)+4)))
	if base.Ui64(v7841) < base.Ui64(v7844<<(uint(int64(32))%64)|v7847) {
		goto L1716
	} else {
		goto L1741
	}
L1741:
	;
	v7851 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7852 = F_RestoreBlockImage(m, v7486, v7758, v7851)
	mBase = m.M
	v7853 = m.ExcPending
	if v7853 != 0 {
		goto L32
	} else {
		goto L1742
	}
L1742:
	;
	if v7852 == int32(0) {
		goto L1636
	} else {
		goto L1743
	}
L1743:
	;
	if v7721 != 0 {
		goto L1744
	} else {
		goto L1745
	}
L1744:
	;
	v7857 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7858 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+552))
	m.T0[v7721].(func(*base.Module, int32, int32))(m, v7857, v7858)
	mBase = m.M
	v7860 = m.ExcPending
	if v7860 != 0 {
		goto L32
	} else {
		goto L1747
	}
L1745:
	;
	goto L1746
L1746:
	;
	v7867 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7869 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7870 = int32(_a_F_StartupXLOG_58)
	goto L1752
L1747:
	;
	v7862 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7863 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+552))
	m.T0[v7721].(func(*base.Module, int32, int32))(m, v7862, v7863)
	mBase = m.M
	v7865 = m.ExcPending
	if v7865 != 0 {
		goto L32
	} else {
		goto L1748
	}
L1748:
	;
	goto L1746
L1749:
	;
	if v7932 != 0 {
		goto L1635
	} else {
		goto L1767
	}
L1750:
	;
	v7932 = int32(0)
	goto L1749
L1751:
	;
	v7906 = v7901
	v7907 = v7902
	v7908 = v7903
	goto L1761
L1752:
	;
	if (v7867|v7869)&int32(3) != 0 {
		v7901 = v7867
		v7902 = v7869
		v7903 = v7870
		goto L1751
	} else {
		goto L1755
	}
L1754:
	;
	if v7891 == int32(0) {
		goto L1750
	} else {
		goto L1760
	}
L1755:
	;
	v7878 = v7867
	v7879 = v7869
	v7880 = v7870
	goto L1756
L1756:
	;
	v7883 = *(*int32)(unsafe.Add(mBase, uint32(v7878)))
	v7884 = *(*int32)(unsafe.Add(mBase, uint32(v7879)))
	if v7883 != v7884 {
		v7901 = v7878
		v7902 = v7879
		v7903 = v7880
		goto L1751
	} else {
		goto L1758
	}
L1757:
	;
	goto L1754
L1758:
	;
	v7886 = int32(4)
	v7887 = v7879 + v7886
	v7889 = v7878 + v7886
	v7891 = v7880 - v7886
	if base.Ui32(int32(3)) < base.Ui32(v7891) {
		v7878 = v7889
		v7879 = v7887
		v7880 = v7891
		goto L1756
	} else {
		goto L1759
	}
L1759:
	;
	goto L1757
L1760:
	;
	v7901 = v7889
	v7902 = v7887
	v7903 = v7891
	goto L1751
L1761:
	;
	v7911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7906))))
	v7912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7907))))
	if v7911 == v7912 {
		goto L1763
	} else {
		goto L1764
	}
L1762:
	;
	v7932 = v7911 - v7912
	goto L1749
L1763:
	;
	v7914 = int32(1)
	v7919 = v7908 - v7914
	if v7919 != 0 {
		v7906 = v7906 + v7914
		v7907 = v7907 + v7914
		v7908 = v7919
		goto L1761
	} else {
		goto L1766
	}
L1764:
	;
	goto L1765
L1765:
	;
	goto L1762
L1766:
	;
	goto L1750
L1767:
	;
	goto L1716
L1768:
	;
	goto L1715
L1769:
	;
	v7982 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v7982+int32(96), int32(_a_F_StartupXLOG_50), int32(2028), int32(_a_F_StartupXLOG_228))
	mBase = m.M
	v7989 = m.ExcPending
	if v7989 != 0 {
		goto L32
	} else {
		goto L1772
	}
L1770:
	;
	goto L1771
L1771:
	;
	v7991 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7992 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v7991)+24)) = v7992
	v7994 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7991)+40)) = v7561
	*(*int64)(unsafe.Add(mBase, uint32(v7991)+32)) = v7994
	v7997 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7991)+96)), uint32(v7997))
	v8001 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v8001 != int32(1) {
		goto L1773
	} else {
		goto L1774
	}
L1772:
	;
	goto L1771
L1773:
	;
	v8012 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[121])))
	if v8012 != 0 {
		goto L1777
	} else {
		goto L1778
	}
L1774:
	;
	v8005 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[122]))
	if v8005 <= int32(0) {
		goto L1773
	} else {
		goto L1775
	}
L1775:
	;
	F_WalSndWakeup(m, v7564, int32(1))
	mBase = m.M
	v8010 = m.ExcPending
	if v8010 != 0 {
		goto L32
	} else {
		goto L1776
	}
L1776:
	;
	goto L1773
L1777:
	;
	v8014 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[121])) = uint8(v8014)
	F_WalRcvForceReply(m)
	mBase = m.M
	v8017 = m.ExcPending
	if v8017 != 0 {
		goto L32
	} else {
		goto L1780
	}
L1778:
	;
	goto L1779
L1779:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v8019 = m.ExcPending
	if v8019 != 0 {
		goto L32
	} else {
		goto L1781
	}
L1780:
	;
	goto L1779
L1781:
	;
	if v7564 != 0 {
		goto L1782
	} else {
		goto L1783
	}
L1782:
	;
	v8020 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+40))
	F_RemoveNonParentXlogFiles(m, v8020, v7561)
	mBase = m.M
	v8022 = m.ExcPending
	if v8022 != 0 {
		goto L32
	} else {
		goto L1785
	}
L1783:
	;
	goto L1784
L1784:
	;
	v8030 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v8030 != int32(1) {
		goto L1634
	} else {
		goto L1787
	}
L1785:
	;
	v8023 = int32(_a_F_StartupXLOG_233)
	v8025 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31])) = v8025 + int32(1)
	goto L1786
L1786:
	;
	goto L1784
L1787:
	;
	v8034 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8035 = *(*int32)(unsafe.Add(mBase, uint32(v8034)+96))
	v8036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8035)+48)))
	v8037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8035)+49)))
	v8039 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v8037|base.B2i32(v8039 != int32(3))|base.B2i32(v8036&int32(240) != int32(112)) == int32(0) {
		goto L1788
	} else {
		goto L1789
	}
L1788:
	;
	v8050 = *(*int32)(unsafe.Add(mBase, uint32(v8035)+64))
	v8052 = v8050 + int32(8)
	v8054 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	v8057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8052))))
	v8060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8054))))
	if base.B2i32(v8057 == int32(0))|base.B2i32(v8057 != v8060) != 0 {
		v8078 = v8057
		v8079 = v8060
		goto L1792
	} else {
		goto L1793
	}
L1789:
	;
	goto L1790
L1790:
	;
	if v8039 != int32(4) {
		goto L1835
	} else {
		goto L1836
	}
L1791:
	;
	if v8078-v8079 != 0 {
		goto L1634
	} else {
		goto L1798
	}
L1792:
	;
	goto L1791
L1793:
	;
	v8063 = v8052
	v8064 = v8054
	goto L1794
L1794:
	;
	v8067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8064)+1)))
	v8068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8063)+1)))
	if v8068 == int32(0) {
		v8078 = v8068
		v8079 = v8067
		goto L1792
	} else {
		goto L1796
	}
L1795:
	;
	v8078 = v8068
	v8079 = v8067
	goto L1792
L1796:
	;
	v8071 = int32(1)
	if v8068 == v8067 {
		v8063 = v8063 + v8071
		v8064 = v8064 + v8071
		goto L1794
	} else {
		goto L1797
	}
L1797:
	;
	goto L1795
L1798:
	;
	v8082 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v8082)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = int32(0)
	v8091 = *(*int64)(unsafe.Add(mBase, uint32(v8050)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = v8091
	v8093 = int32(_a_F_StartupXLOG_234)
	goto L1802
L1799:
	;
	v8215 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8216 = m.ExcPending
	if v8216 != 0 {
		goto L32
	} else {
		goto L1830
	}
L1800:
	;
	v8210 = F_strlen(m, v8199)
	mBase = m.M
	goto L1799
L1802:
	;
	goto L1803
L1803:
	;
	v8100 = int32(63)
	if (v8093^v8052)&int32(3) != 0 {
		goto L1807
	} else {
		goto L1808
	}
L1804:
	;
	v8203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8200))) = uint8(v8203)
	goto L1800
L1805:
	;
	v8184 = v8179
	v8185 = v8180
	v8186 = v8181
	goto L1826
L1806:
	;
	if v8174 == int32(0) {
		v8199 = v8172
		v8200 = v8173
		goto L1804
	} else {
		goto L1825
	}
L1807:
	;
	v8172 = v8052
	v8173 = v8093
	v8174 = v8100
	goto L1806
L1808:
	;
	goto L1809
L1809:
	;
	v8104 = int32(0)
	if base.B2i32(v8052&int32(3) == v8104)|int32(0) == v8104 {
		goto L1811
	} else {
		goto L1812
	}
L1810:
	;
	if v8140 == int32(0) {
		v8199 = v8137
		v8200 = v8138
		goto L1804
	} else {
		goto L1819
	}
L1811:
	;
	v8116 = v8052
	v8117 = v8093
	v8118 = v8100
	goto L1814
L1812:
	;
	goto L1813
L1813:
	;
	v8137 = v8052
	v8138 = v8093
	v8139 = v8100
	v8140 = int32(1)
	goto L1810
L1814:
	;
	v8120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8117))) = uint8(v8120)
	if v8120 == int32(0) {
		v8179 = v8116
		v8180 = v8117
		v8181 = v8118
		goto L1805
	} else {
		goto L1816
	}
L1815:
	;
	v8137 = v8131
	v8138 = v8125
	v8139 = v8127
	v8140 = v8129
	goto L1810
L1816:
	;
	v8124 = int32(1)
	v8125 = v8117 + v8124
	v8127 = v8118 - v8124
	v8128 = int32(0)
	v8129 = base.B2i32(v8127 != v8128)
	v8131 = v8116 + v8124
	if v8131&int32(3) == v8128 {
		v8137 = v8131
		v8138 = v8125
		v8139 = v8127
		v8140 = v8129
		goto L1810
	} else {
		goto L1817
	}
L1817:
	;
	if v8127 != 0 {
		v8116 = v8131
		v8117 = v8125
		v8118 = v8127
		goto L1814
	} else {
		goto L1818
	}
L1818:
	;
	goto L1815
L1819:
	;
	v8143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8137))))
	if base.B2i32(v8143 == int32(0))|base.B2i32(base.Ui32(v8139) < base.Ui32(int32(4))) != 0 {
		v8172 = v8137
		v8173 = v8138
		v8174 = v8139
		goto L1806
	} else {
		goto L1820
	}
L1820:
	;
	v8150 = v8137
	v8151 = v8138
	v8152 = v8139
	goto L1821
L1821:
	;
	v8155 = *(*int32)(unsafe.Add(mBase, uint32(v8150)))
	v8158 = int32(-2139062144)
	if (int32(16843008)-v8155|v8155)&v8158 != v8158 {
		v8179 = v8150
		v8180 = v8151
		v8181 = v8152
		goto L1805
	} else {
		goto L1823
	}
L1822:
	;
	v8172 = v8166
	v8173 = v8164
	v8174 = v8168
	goto L1806
L1823:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8151))) = v8155
	v8163 = int32(4)
	v8164 = v8151 + v8163
	v8166 = v8150 + v8163
	v8168 = v8152 - v8163
	if base.Ui32(int32(3)) < base.Ui32(v8168) {
		v8150 = v8166
		v8151 = v8164
		v8152 = v8168
		goto L1821
	} else {
		goto L1824
	}
L1824:
	;
	goto L1822
L1825:
	;
	v8179 = v8172
	v8180 = v8173
	v8181 = v8174
	goto L1805
L1826:
	;
	v8188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8185))) = uint8(v8188)
	if v8188 == int32(0) {
		v8199 = v8184
		v8200 = v8185
		goto L1804
	} else {
		goto L1828
	}
L1827:
	;
	v8199 = v8195
	v8200 = v8193
	goto L1804
L1828:
	;
	v8192 = int32(1)
	v8193 = v8185 + v8192
	v8195 = v8184 + v8192
	v8197 = v8186 - v8192
	if v8197 != 0 {
		v8184 = v8195
		v8185 = v8193
		v8186 = v8197
		goto L1826
	} else {
		goto L1829
	}
L1829:
	;
	goto L1827
L1830:
	;
	if v8215 == int32(0) {
		goto L1448
	} else {
		goto L1831
	}
L1831:
	;
	v8220 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v8221 = F_timestamptz_to_str(m, v8220)
	mBase = m.M
	v8222 = m.ExcPending
	if v8222 != 0 {
		goto L32
	} else {
		goto L1832
	}
L1832:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+36)) = v8221
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+32)) = int32(_a_F_StartupXLOG_234)
	F_errmsg(m, int32(_a_F_StartupXLOG_235), v6435+int32(32))
	mBase = m.M
	v8230 = m.ExcPending
	if v8230 != 0 {
		goto L32
	} else {
		goto L1833
	}
L1833:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2787), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8235 = m.ExcPending
	if v8235 != 0 {
		goto L32
	} else {
		goto L1834
	}
L1834:
	;
	goto L1448
L1835:
	;
	if v8037 != int32(1) {
		goto L1634
	} else {
		goto L1843
	}
L1836:
	;
	v8239 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[115])))
	if v8239&int32(1) == int32(0) {
		goto L1835
	} else {
		goto L1837
	}
L1837:
	;
	v8244 = *(*int64)(unsafe.Add(mBase, uint32(v8034)+32))
	v8246 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	if base.Ui64(v8244) < base.Ui64(v8246) {
		goto L1835
	} else {
		goto L1838
	}
L1838:
	;
	v8249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v8249)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v8244
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = int64(0)
	v8257 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v8257
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v8257)
	v8264 = F_errstart(m, int32(15), v8257)
	mBase = m.M
	v8265 = m.ExcPending
	if v8265 != 0 {
		goto L32
	} else {
		goto L1839
	}
L1839:
	;
	if v8264 == int32(0) {
		goto L1448
	} else {
		goto L1840
	}
L1840:
	;
	v8269 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+84)) = uint32(v8269)
	v8272 = int64(base.Ui64(v8269) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+80)) = uint32(v8272)
	F_errmsg(m, int32(_a_F_StartupXLOG_237), v6435+int32(80))
	mBase = m.M
	v8278 = m.ExcPending
	if v8278 != 0 {
		goto L32
	} else {
		goto L1841
	}
L1841:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2804), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8283 = m.ExcPending
	if v8283 != 0 {
		goto L32
	} else {
		goto L1842
	}
L1842:
	;
	goto L1448
L1843:
	;
	v8288 = v8036 & int32(112)
	v8289 = int32(4)
	v8290 = int32(base.Ui32(v8288) >> (uint(v8289) % 32))
	if base.B2i32(base.Ui32(v8289) < base.Ui32(v8290))|base.B2i32(v8290 == int32(1)) != 0 {
		v8638 = v8039
		goto L1844
	} else {
		goto L1845
	}
L1844:
	;
	if v8638 != int32(5) {
		goto L1634
	} else {
		goto L1918
	}
L1845:
	;
	v8297 = int32(4)
	v8300 = int32(base.Ui32(v8036)>>(uint(v8297)%32)) & int32(7)
	if base.B2i32(base.Ui32(v8297) < base.Ui32(v8300))|base.B2i32(v8300 == int32(1)) == int32(0) {
		goto L1846
	} else {
		goto L1847
	}
L1846:
	;
	v8308 = *(*int32)(unsafe.Add(mBase, uint32(v8035)+64))
	v8309 = *(*int64)(unsafe.Add(mBase, uint32(v8308)))
	v8311 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8314 = base.AtomicRmwXchg32(m, v8311, int32(96), int32(1))
	if v8314 != 0 {
		goto L1849
	} else {
		goto L1850
	}
L1847:
	;
	v8332 = v8035
	v8333 = int64(0)
	goto L1848
L1848:
	;
	v8335 = v8288 - int32(48)
	if v8335 != 0 {
		goto L1856
	} else {
		goto L1857
	}
L1849:
	;
	v8316 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v8316+int32(96), int32(_a_F_StartupXLOG_50), int32(_a_F_StartupXLOG_238), int32(_a_F_StartupXLOG_239))
	mBase = m.M
	v8323 = m.ExcPending
	if v8323 != 0 {
		goto L32
	} else {
		goto L1852
	}
L1850:
	;
	goto L1851
L1851:
	;
	v8325 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v8325)+64)) = v8309
	v8327 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8325)+96)), uint32(v8327))
	v8330 = *(*int32)(unsafe.Add(mBase, uint32(v8034)+96))
	v8332 = v8330
	v8333 = v8309
	goto L1848
L1852:
	;
	goto L1851
L1853:
	;
	v8565 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v8565 != int32(1) {
		v8638 = v8565
		goto L1844
	} else {
		goto L1903
	}
L1854:
	;
	v8562 = *(*int32)(unsafe.Add(mBase, uint32(v8332)+36))
	v8563 = v8562
	goto L1853
L1855:
	;
	v8457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8332)+48)))
	v8458 = *(*int32)(unsafe.Add(mBase, uint32(v8332)+64))
	v8460 = v6435 + int32(560)
	v8461 = int32(0)
	base.MemoryFill(m, v8460, v8461, int32(264))
	v8467 = *(*int64)(unsafe.Add(mBase, uint32(v8458)))
	*(*int64)(unsafe.Add(mBase, uint32(v8460))) = v8467
	if v8461 <= base.I32_extend8_s(v8457) {
		goto L1885
	} else {
		goto L1886
	}
L1856:
	;
	if v8335 == int32(16) {
		goto L1859
	} else {
		goto L1860
	}
L1857:
	;
	goto L1858
L1858:
	;
	v8338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8332)+48)))
	v8339 = *(*int32)(unsafe.Add(mBase, uint32(v8332)+64))
	v8341 = v6435 + int32(560)
	v8342 = int32(0)
	base.MemoryFill(m, v8341, v8342, int32(288))
	v8348 = *(*int64)(unsafe.Add(mBase, uint32(v8339)))
	*(*int64)(unsafe.Add(mBase, uint32(v8341))) = v8348
	if v8342 <= base.I32_extend8_s(v8338) {
		goto L1863
	} else {
		goto L1864
	}
L1859:
	;
	goto L1855
L1860:
	;
	goto L1854
L1862:
	;
	v8456 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+612))
	v8563 = v8456
	goto L1853
L1863:
	;
	goto L1862
L1864:
	;
	v8353 = *(*int32)(unsafe.Add(mBase, uint32(v8339)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+8)) = v8353
	if v8353&int32(1) != 0 {
		goto L1865
	} else {
		goto L1866
	}
L1865:
	;
	v8357 = *(*int32)(unsafe.Add(mBase, uint32(v8339)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+12)) = v8357
	v8359 = *(*int32)(unsafe.Add(mBase, uint32(v8339)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+16)) = v8359
	v8365 = v8339 + int32(20)
	goto L1867
L1866:
	;
	v8365 = v8339 + int32(12)
	goto L1867
L1867:
	;
	if v8353&int32(2) != 0 {
		goto L1868
	} else {
		goto L1869
	}
L1868:
	;
	v8368 = *(*int32)(unsafe.Add(mBase, uint32(v8365)))
	v8370 = v8365 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+24)) = v8370
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+20)) = v8368
	v8376 = v8370 + v8368<<(uint(int32(2))%32)
	goto L1870
L1869:
	;
	v8376 = v8365
	goto L1870
L1870:
	;
	if v8353&int32(4) != 0 {
		goto L1871
	} else {
		goto L1872
	}
L1871:
	;
	v8380 = *(*int32)(unsafe.Add(mBase, uint32(v8376)))
	v8382 = v8376 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+32)) = v8382
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+28)) = v8380
	v8385 = *(*int32)(unsafe.Add(mBase, uint32(v8376)))
	v8389 = v8382 + v8385*int32(12)
	goto L1873
L1872:
	;
	v8389 = v8376
	goto L1873
L1873:
	;
	if v8353&int32(256) != 0 {
		goto L1874
	} else {
		goto L1875
	}
L1874:
	;
	v8394 = *(*int32)(unsafe.Add(mBase, uint32(v8389)))
	v8395 = int32(4)
	v8396 = v8389 + v8395
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+40)) = v8396
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+36)) = v8394
	v8399 = *(*int32)(unsafe.Add(mBase, uint32(v8389)))
	v8403 = v8396 + v8399<<(uint(v8395)%32)
	goto L1876
L1875:
	;
	v8403 = v8389
	goto L1876
L1876:
	;
	if v8353&int32(8) != 0 {
		goto L1877
	} else {
		goto L1878
	}
L1877:
	;
	v8408 = *(*int32)(unsafe.Add(mBase, uint32(v8403)))
	v8409 = int32(4)
	v8410 = v8403 + v8409
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+48)) = v8410
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+44)) = v8408
	v8413 = *(*int32)(unsafe.Add(mBase, uint32(v8403)))
	v8417 = v8410 + v8413<<(uint(v8409)%32)
	goto L1879
L1878:
	;
	v8417 = v8403
	goto L1879
L1879:
	;
	if v8353&int32(16) == int32(0) {
		v8441 = v8353
		v8442 = v8417
		goto L1880
	} else {
		goto L1881
	}
L1880:
	;
	if v8441&int32(32) == int32(0) {
		goto L1863
	} else {
		goto L1883
	}
L1881:
	;
	v8424 = *(*int32)(unsafe.Add(mBase, uint32(v8417)))
	*(*int32)(unsafe.Add(mBase, uint32(v8341)+52)) = v8424
	v8427 = v8417 + int32(4)
	if v8353&int32(128) == int32(0) {
		v8441 = v8353
		v8442 = v8427
		goto L1880
	} else {
		goto L1882
	}
L1882:
	;
	v8435 = F_strlcpy(m, v6435+int32(616), v8427, int32(200))
	mBase = m.M
	v8436 = F_strlen(m, v8427)
	mBase = m.M
	v8440 = *(*int32)(unsafe.Add(mBase, uint32(v8341)+8))
	v8441 = v8440
	v8442 = v8436 + v8427 + int32(1)
	goto L1880
L1883:
	;
	v8447 = *(*int64)(unsafe.Add(mBase, uint32(v8442)))
	v8448 = *(*int64)(unsafe.Add(mBase, uint32(v8442)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8341)+280)) = v8448
	*(*int64)(unsafe.Add(mBase, uint32(v8341)+272)) = v8447
	goto L1863
L1884:
	;
	v8561 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+604))
	v8563 = v8561
	goto L1853
L1885:
	;
	goto L1884
L1886:
	;
	v8472 = *(*int32)(unsafe.Add(mBase, uint32(v8458)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+8)) = v8472
	if v8472&int32(1) != 0 {
		goto L1887
	} else {
		goto L1888
	}
L1887:
	;
	v8476 = *(*int32)(unsafe.Add(mBase, uint32(v8458)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+12)) = v8476
	v8478 = *(*int32)(unsafe.Add(mBase, uint32(v8458)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+16)) = v8478
	v8484 = v8458 + int32(20)
	goto L1889
L1888:
	;
	v8484 = v8458 + int32(12)
	goto L1889
L1889:
	;
	if v8472&int32(2) != 0 {
		goto L1890
	} else {
		goto L1891
	}
L1890:
	;
	v8487 = *(*int32)(unsafe.Add(mBase, uint32(v8484)))
	v8489 = v8484 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+24)) = v8489
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+20)) = v8487
	v8495 = v8489 + v8487<<(uint(int32(2))%32)
	goto L1892
L1891:
	;
	v8495 = v8484
	goto L1892
L1892:
	;
	if v8472&int32(4) != 0 {
		goto L1893
	} else {
		goto L1894
	}
L1893:
	;
	v8499 = *(*int32)(unsafe.Add(mBase, uint32(v8495)))
	v8501 = v8495 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+32)) = v8501
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+28)) = v8499
	v8504 = *(*int32)(unsafe.Add(mBase, uint32(v8495)))
	v8508 = v8501 + v8504*int32(12)
	goto L1895
L1894:
	;
	v8508 = v8495
	goto L1895
L1895:
	;
	if v8472&int32(256) != 0 {
		goto L1896
	} else {
		goto L1897
	}
L1896:
	;
	v8513 = *(*int32)(unsafe.Add(mBase, uint32(v8508)))
	v8514 = int32(4)
	v8515 = v8508 + v8514
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+40)) = v8515
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+36)) = v8513
	v8518 = *(*int32)(unsafe.Add(mBase, uint32(v8508)))
	v8522 = v8515 + v8518<<(uint(v8514)%32)
	goto L1898
L1897:
	;
	v8522 = v8508
	goto L1898
L1898:
	;
	if v8472&int32(16) == int32(0) {
		v8546 = v8472
		v8547 = v8522
		goto L1899
	} else {
		goto L1900
	}
L1899:
	;
	if v8546&int32(32) == int32(0) {
		goto L1885
	} else {
		goto L1902
	}
L1900:
	;
	v8529 = *(*int32)(unsafe.Add(mBase, uint32(v8522)))
	*(*int32)(unsafe.Add(mBase, uint32(v8460)+44)) = v8529
	v8532 = v8522 + int32(4)
	if v8472&int32(128) == int32(0) {
		v8546 = v8472
		v8547 = v8532
		goto L1899
	} else {
		goto L1901
	}
L1901:
	;
	v8540 = F_strlcpy(m, v6435+int32(608), v8532, int32(200))
	mBase = m.M
	v8541 = F_strlen(m, v8532)
	mBase = m.M
	v8545 = *(*int32)(unsafe.Add(mBase, uint32(v8460)+8))
	v8546 = v8545
	v8547 = v8541 + v8532 + int32(1)
	goto L1899
L1902:
	;
	v8552 = *(*int64)(unsafe.Add(mBase, uint32(v8547)))
	v8553 = *(*int64)(unsafe.Add(mBase, uint32(v8547)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8460)+256)) = v8553
	*(*int64)(unsafe.Add(mBase, uint32(v8460)+248)) = v8552
	goto L1885
L1903:
	;
	v8569 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[115])))
	if v8569&int32(1) == int32(0) {
		v8638 = v8565
		goto L1844
	} else {
		goto L1904
	}
L1904:
	;
	v8575 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	if v8563 != v8575 {
		v8638 = v8565
		goto L1844
	} else {
		goto L1905
	}
L1905:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v8563
	v8580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v8580)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = v8333
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = int64(0)
	v8588 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v8588)
	switch v8290 {
	case 0, 3:
		goto L1907
	default:
		goto L1448
	case 2, 4:
		goto L1906
	}
L1906:
	;
	v8616 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8617 = m.ExcPending
	if v8617 != 0 {
		goto L32
	} else {
		goto L1913
	}
L1907:
	;
	v8592 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8593 = m.ExcPending
	if v8593 != 0 {
		goto L32
	} else {
		goto L1908
	}
L1908:
	;
	if v8592 == int32(0) {
		goto L1448
	} else {
		goto L1909
	}
L1909:
	;
	v8597 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	v8599 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v8600 = F_timestamptz_to_str(m, v8599)
	mBase = m.M
	v8601 = m.ExcPending
	if v8601 != 0 {
		goto L32
	} else {
		goto L1910
	}
L1910:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+52)) = v8600
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+48)) = v8597
	F_errmsg(m, int32(_a_F_StartupXLOG_240), v6435+int32(48))
	mBase = m.M
	v8608 = m.ExcPending
	if v8608 != 0 {
		goto L32
	} else {
		goto L1911
	}
L1911:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2872), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8613 = m.ExcPending
	if v8613 != 0 {
		goto L32
	} else {
		goto L1912
	}
L1912:
	;
	goto L1448
L1913:
	;
	if v8616 == int32(0) {
		goto L1448
	} else {
		goto L1914
	}
L1914:
	;
	v8621 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	v8623 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v8624 = F_timestamptz_to_str(m, v8623)
	mBase = m.M
	v8625 = m.ExcPending
	if v8625 != 0 {
		goto L32
	} else {
		goto L1915
	}
L1915:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+68)) = v8624
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+64)) = v8621
	F_errmsg(m, int32(_a_F_StartupXLOG_241), v6435-int32(-64))
	mBase = m.M
	v8632 = m.ExcPending
	if v8632 != 0 {
		goto L32
	} else {
		goto L1916
	}
L1916:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2880), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8637 = m.ExcPending
	if v8637 != 0 {
		goto L32
	} else {
		goto L1917
	}
L1917:
	;
	goto L1448
L1918:
	;
	v8645 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[101])))
	if v8645&int32(1) == int32(0) {
		goto L1634
	} else {
		goto L1919
	}
L1919:
	;
	v8652 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8653 = m.ExcPending
	if v8653 != 0 {
		goto L32
	} else {
		goto L1920
	}
L1920:
	;
	if v8652 != 0 {
		goto L1921
	} else {
		goto L1922
	}
L1921:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_221), int32(0))
	mBase = m.M
	v8657 = m.ExcPending
	if v8657 != 0 {
		goto L32
	} else {
		goto L1924
	}
L1922:
	;
	goto L1923
L1923:
	;
	v8664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v8664)
	v8667 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = v8667
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = v8667
	v8673 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v8673
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[106])) = uint8(v8673)
	goto L1448
L1924:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2890), int32(_a_F_StartupXLOG_236))
	mBase = m.M
	v8662 = m.ExcPending
	if v8662 != 0 {
		goto L32
	} else {
		goto L1925
	}
L1925:
	;
	goto L1923
L1926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+244)) = v6685
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+240)) = v7506
	F_errmsg(m, int32(_a_F_StartupXLOG_242), v6435+int32(240))
	mBase = m.M
	v8688 = m.ExcPending
	if v8688 != 0 {
		goto L32
	} else {
		goto L1927
	}
L1927:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2406), int32(_a_F_StartupXLOG_243))
	mBase = m.M
	v8693 = m.ExcPending
	if v8693 != 0 {
		goto L32
	} else {
		goto L1928
	}
L1928:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1929:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+212)) = v6685
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+208)) = v7504
	F_errmsg(m, int32(_a_F_StartupXLOG_244), v6435+int32(208))
	mBase = m.M
	v8705 = m.ExcPending
	if v8705 != 0 {
		goto L32
	} else {
		goto L1930
	}
L1930:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2415), int32(_a_F_StartupXLOG_243))
	mBase = m.M
	v8710 = m.ExcPending
	if v8710 != 0 {
		goto L32
	} else {
		goto L1931
	}
L1931:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1932:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+224)) = v7504
	v8717 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+232)) = uint32(v8717)
	v8720 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+236)) = v8720
	v8723 = int64(base.Ui64(v8717) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+228)) = uint32(v8723)
	F_errmsg(m, int32(_a_F_StartupXLOG_245), v6435+int32(224))
	mBase = m.M
	v8729 = m.ExcPending
	if v8729 != 0 {
		goto L32
	} else {
		goto L1933
	}
L1933:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2433), int32(_a_F_StartupXLOG_243))
	mBase = m.M
	v8734 = m.ExcPending
	if v8734 != 0 {
		goto L32
	} else {
		goto L1934
	}
L1934:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1935:
	;
	v8739 = *(*int64)(unsafe.Add(mBase, uint32(v7486)+64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+188)) = uint32(v8739)
	v8741 = int64(32)
	v8742 = int64(base.Ui64(v8739) >> (uint(v8741) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+184)) = uint32(v8742)
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+180)) = uint32(v7608)
	v8746 = int64(base.Ui64(v7608) >> (uint(v8741) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+176)) = uint32(v8746)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_246), v6435+int32(176))
	mBase = m.M
	v8752 = m.ExcPending
	if v8752 != 0 {
		goto L32
	} else {
		goto L1936
	}
L1936:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2108), int32(_a_F_StartupXLOG_230))
	mBase = m.M
	v8757 = m.ExcPending
	if v8757 != 0 {
		goto L32
	} else {
		goto L1937
	}
L1937:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1938:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v8764 = m.ExcPending
	if v8764 != 0 {
		goto L32
	} else {
		goto L1939
	}
L1939:
	;
	v8765 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+128)) = v8765
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_247), v6435+int32(128))
	mBase = m.M
	v8771 = m.ExcPending
	if v8771 != 0 {
		goto L32
	} else {
		goto L1940
	}
L1940:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2563), int32(_a_F_StartupXLOG_248))
	mBase = m.M
	v8776 = m.ExcPending
	if v8776 != 0 {
		goto L32
	} else {
		goto L1941
	}
L1941:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1942:
	;
	v8781 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+552))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+112)) = v8781
	v8783 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+96)) = v8783
	v8785 = *(*int64)(unsafe.Add(mBase, uint32(v6435)+564))
	*(*int64)(unsafe.Add(mBase, uint32(v6435)+100)) = v8785
	v8787 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+108)) = v8787
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_249), v6435+int32(96))
	mBase = m.M
	v8793 = m.ExcPending
	if v8793 != 0 {
		goto L32
	} else {
		goto L1943
	}
L1943:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2581), int32(_a_F_StartupXLOG_248))
	mBase = m.M
	v8798 = m.ExcPending
	if v8798 != 0 {
		goto L32
	} else {
		goto L1944
	}
L1944:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1945:
	;
	if v8811 != 0 {
		v6685 = v7561
		v6690 = v8811
		goto L1488
	} else {
		goto L1946
	}
L1946:
	;
	goto L1489
L1947:
	;
	if v8815 == int32(0) {
		v9130 = v6575
		goto L1446
	} else {
		goto L1948
	}
L1948:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_250), int32(0))
	mBase = m.M
	v8822 = m.ExcPending
	if v8822 != 0 {
		goto L32
	} else {
		goto L1949
	}
L1949:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1909), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v8827 = m.ExcPending
	if v8827 != 0 {
		goto L32
	} else {
		goto L1950
	}
L1950:
	;
	v9130 = v6575
	goto L1446
L1951:
	;
	if v7124 != 0 {
		goto L1952
	} else {
		goto L1953
	}
L1952:
	;
	if v8846 == int32(0) {
		goto L1448
	} else {
		goto L1955
	}
L1953:
	;
	goto L1954
L1954:
	;
	if v8846 == int32(0) {
		goto L1448
	} else {
		goto L1959
	}
L1955:
	;
	v8851 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	v8853 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v8854 = F_timestamptz_to_str(m, v8853)
	mBase = m.M
	v8855 = m.ExcPending
	if v8855 != 0 {
		goto L32
	} else {
		goto L1956
	}
L1956:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+276)) = v8854
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+272)) = v8851
	F_errmsg(m, int32(_a_F_StartupXLOG_251), v6435+int32(272))
	mBase = m.M
	v8862 = m.ExcPending
	if v8862 != 0 {
		goto L32
	} else {
		goto L1957
	}
L1957:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2727), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v8867 = m.ExcPending
	if v8867 != 0 {
		goto L32
	} else {
		goto L1958
	}
L1958:
	;
	goto L1448
L1959:
	;
	v8871 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	v8873 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v8874 = F_timestamptz_to_str(m, v8873)
	mBase = m.M
	v8875 = m.ExcPending
	if v8875 != 0 {
		goto L32
	} else {
		goto L1960
	}
L1960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+292)) = v8874
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+288)) = v8871
	F_errmsg(m, int32(_a_F_StartupXLOG_252), v6435+int32(288))
	mBase = m.M
	v8882 = m.ExcPending
	if v8882 != 0 {
		goto L32
	} else {
		goto L1961
	}
L1961:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2734), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v8887 = m.ExcPending
	if v8887 != 0 {
		goto L32
	} else {
		goto L1962
	}
L1962:
	;
	goto L1448
L1963:
	;
	v8926 = int32(1)
	v8928 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[18]))
	switch v8928 {
	case 0:
		goto L1964
	default:
		v8966 = v8926
		goto L1447
	case 2:
		goto L1965
	}
L1964:
	;
	v8933 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8936 = base.AtomicRmwXchg32(m, v8933, int32(96), int32(1))
	if v8936 != 0 {
		goto L1967
	} else {
		goto L1968
	}
L1965:
	;
	F_proc_exit(m, int32(3))
	mBase = m.M
	v8931 = m.ExcPending
	if v8931 != 0 {
		goto L32
	} else {
		goto L1966
	}
L1966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1967:
	;
	v8938 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v8938+int32(96), int32(_a_F_StartupXLOG_50), int32(3114), int32(_a_F_StartupXLOG_253))
	mBase = m.M
	v8945 = m.ExcPending
	if v8945 != 0 {
		goto L32
	} else {
		goto L1970
	}
L1968:
	;
	goto L1969
L1969:
	;
	v8947 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8948 = *(*int32)(unsafe.Add(mBase, uint32(v8947)+80))
	if v8948 == int32(0) {
		goto L1971
	} else {
		goto L1972
	}
L1970:
	;
	goto L1969
L1971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8947)+80)) = int32(1)
	goto L1973
L1972:
	;
	goto L1973
L1973:
	;
	v8953 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8947)+96)), uint32(v8953))
	F_recoveryPausesHere(m, int32(1))
	mBase = m.M
	v8958 = m.ExcPending
	if v8958 != 0 {
		goto L32
	} else {
		goto L1974
	}
L1974:
	;
	v8966 = v8926
	goto L1447
L1975:
	;
	v9029 = v8997 << (uint(int32(5)) % 32)
	v9032 = *(*int32)(unsafe.Add(mBase, uint32(v9029)+uint32(_c_F_StartupXLOG[108])))
	if v9032 == int32(0) {
		goto L1977
	} else {
		goto L1978
	}
L1976:
	;
	v9056 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9057 = m.ExcPending
	if v9057 != 0 {
		goto L32
	} else {
		goto L1986
	}
L1977:
	;
	v9041 = *(*int32)(unsafe.Add(mBase, uint32(v9029)+uint32(_c_F_StartupXLOG[109])))
	if v9041 == int32(0) {
		goto L1981
	} else {
		goto L1982
	}
L1978:
	;
	v9035 = *(*int32)(unsafe.Add(mBase, uint32(v9029)+uint32(_c_F_StartupXLOG[123])))
	if v9035 == int32(0) {
		goto L1977
	} else {
		goto L1979
	}
L1979:
	;
	m.T0[v9035].(func(*base.Module))(m)
	mBase = m.M
	v9039 = m.ExcPending
	if v9039 != 0 {
		goto L32
	} else {
		goto L1980
	}
L1980:
	;
	goto L1977
L1981:
	;
	v9051 = v8997 + int32(2)
	if v9051 != int32(256) {
		v8997 = v9051
		goto L1975
	} else {
		goto L1985
	}
L1982:
	;
	v9044 = *(*int32)(unsafe.Add(mBase, uint32(v9029)+uint32(_c_F_StartupXLOG[124])))
	if v9044 == int32(0) {
		goto L1981
	} else {
		goto L1983
	}
L1983:
	;
	m.T0[v9044].(func(*base.Module))(m)
	mBase = m.M
	v9048 = m.ExcPending
	if v9048 != 0 {
		goto L32
	} else {
		goto L1984
	}
L1984:
	;
	goto L1981
L1985:
	;
	goto L1976
L1986:
	;
	if v9056 != 0 {
		goto L1987
	} else {
		goto L1988
	}
L1987:
	;
	v9059 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9060 = *(*int64)(unsafe.Add(mBase, uint32(v9059)+32))
	v9063 = F_pg_rusage_show(m, v6435+int32(368))
	mBase = m.M
	v9064 = m.ExcPending
	if v9064 != 0 {
		goto L32
	} else {
		goto L1990
	}
L1988:
	;
	goto L1989
L1989:
	;
	v9082 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v9085 = base.AtomicRmwXchg32(m, v9082, int32(96), int32(1))
	if v9085 != 0 {
		goto L1993
	} else {
		goto L1994
	}
L1990:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+24)) = v9063
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+20)) = uint32(v9060)
	v9068 = int64(base.Ui64(v9060) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6435)+16)) = uint32(v9068)
	F_errmsg(m, int32(_a_F_StartupXLOG_254), v6435+int32(16))
	mBase = m.M
	v9074 = m.ExcPending
	if v9074 != 0 {
		goto L32
	} else {
		goto L1991
	}
L1991:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1896), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9079 = m.ExcPending
	if v9079 != 0 {
		goto L32
	} else {
		goto L1992
	}
L1992:
	;
	goto L1989
L1993:
	;
	v9087 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v9087+int32(96), int32(_a_F_StartupXLOG_50), int32(_a_F_StartupXLOG_255), int32(_a_F_StartupXLOG_256))
	mBase = m.M
	v9094 = m.ExcPending
	if v9094 != 0 {
		goto L32
	} else {
		goto L1996
	}
L1994:
	;
	goto L1995
L1995:
	;
	v9096 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v9097 = *(*int64)(unsafe.Add(mBase, uint32(v9096)+64))
	v9098 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9096)+96)), uint32(v9098))
	if v9097 == int64(0) {
		goto L1997
	} else {
		goto L1998
	}
L1996:
	;
	goto L1995
L1997:
	;
	v9121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[107])) = uint8(v9121)
	v9130 = v8966
	goto L1446
L1998:
	;
	v9105 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9106 = m.ExcPending
	if v9106 != 0 {
		goto L32
	} else {
		goto L1999
	}
L1999:
	;
	if v9105 == int32(0) {
		goto L1997
	} else {
		goto L2000
	}
L2000:
	;
	v9109 = F_timestamptz_to_str(m, v9097)
	mBase = m.M
	v9110 = m.ExcPending
	if v9110 != 0 {
		goto L32
	} else {
		goto L2001
	}
L2001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6435))) = v9109
	F_errmsg(m, int32(_a_F_StartupXLOG_257), v6435)
	mBase = m.M
	v9114 = m.ExcPending
	if v9114 != 0 {
		goto L32
	} else {
		goto L2002
	}
L2002:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1901), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9119 = m.ExcPending
	if v9119 != 0 {
		goto L32
	} else {
		goto L2003
	}
L2003:
	;
	goto L1997
L2004:
	;
	m.G0 = v6435 + int32(848)
	goto L1443
L2005:
	;
	v9158 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9158&int32(1) == int32(0) {
		goto L2004
	} else {
		goto L2006
	}
L2006:
	;
	v9164 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v9164 != 0 {
		goto L1444
	} else {
		goto L2007
	}
L2007:
	;
	goto L2004
L2008:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_258), int32(0))
	mBase = m.M
	v9175 = m.ExcPending
	if v9175 != 0 {
		goto L32
	} else {
		goto L2009
	}
L2009:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1862), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9180 = m.ExcPending
	if v9180 != 0 {
		goto L32
	} else {
		goto L2010
	}
L2010:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2011:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v9187 = m.ExcPending
	if v9187 != 0 {
		goto L32
	} else {
		goto L2012
	}
L2012:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_259), int32(0))
	mBase = m.M
	v9191 = m.ExcPending
	if v9191 != 0 {
		goto L32
	} else {
		goto L2013
	}
L2013:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1921), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9196 = m.ExcPending
	if v9196 != 0 {
		goto L32
	} else {
		goto L2014
	}
L2014:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2015:
	;
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v9240 = m.ExcPending
	if v9240 != 0 {
		goto L32
	} else {
		goto L2016
	}
L2016:
	;
	v9243 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	v9246 = base.AtomicRmwXchg32(m, v9243, int32(16), int32(1))
	if v9246 != 0 {
		goto L2017
	} else {
		goto L2018
	}
L2017:
	;
	v9248 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	F_s_lock(m, v9248+int32(16), int32(_a_F_StartupXLOG_260), int32(1589), int32(_a_F_StartupXLOG_261))
	mBase = m.M
	v9255 = m.ExcPending
	if v9255 != 0 {
		goto L32
	} else {
		goto L2020
	}
L2018:
	;
	goto L2019
L2019:
	;
	v9257 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	v9258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9257)+4)) = uint8(v9258)
	v9260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9257)+5)))
	if v9260 != v9258 {
		v9354 = v9257
		goto L2021
	} else {
		goto L2022
	}
L2020:
	;
	goto L2019
L2021:
	;
	v9386 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9354)+16)), uint32(v9386))
	v9391 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v9391 == int32(1) {
		goto L2039
	} else {
		goto L2040
	}
L2022:
	;
	v9263 = *(*int32)(unsafe.Add(mBase, uint32(v9257)))
	v9264 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9257)+16)), uint32(v9264))
	if v9263 != int32(-1) {
		goto L2023
	} else {
		goto L2024
	}
L2023:
	;
	v9270 = F_pgmem_kill(m, v9263, int32(10))
	mBase = m.M
	goto L2025
L2024:
	;
	goto L2025
L2025:
	;
	goto L2026
L2026:
	;
	v9306 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[126]))
	v9310 = F_WaitLatch(m, v9306, int32(41), int32(10), int32(83886092))
	mBase = m.M
	v9311 = m.ExcPending
	if v9311 != 0 {
		goto L32
	} else {
		goto L2029
	}
L2028:
	;
	v9331 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	v9334 = base.AtomicRmwXchg32(m, v9331, int32(16), int32(1))
	if v9334 != 0 {
		goto L2034
	} else {
		goto L2035
	}
L2029:
	;
	if v9310&int32(1) == int32(0) {
		goto L2028
	} else {
		goto L2030
	}
L2030:
	;
	v9317 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[126]))
	v9318 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9317))) = v9318
	v9323 = base.AtomicRmwOr32(m, v9318, int32(_a_F_StartupXLOG_225), v9318)
	goto L2031
L2031:
	;
	v9325 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	if v9325 == int32(0) {
		goto L2028
	} else {
		goto L2032
	}
L2032:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9329 = m.ExcPending
	if v9329 != 0 {
		goto L32
	} else {
		goto L2033
	}
L2033:
	;
	goto L2028
L2034:
	;
	v9336 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	F_s_lock(m, v9336+int32(16), int32(_a_F_StartupXLOG_260), int32(1631), int32(_a_F_StartupXLOG_261))
	mBase = m.M
	v9343 = m.ExcPending
	if v9343 != 0 {
		goto L32
	} else {
		goto L2037
	}
L2035:
	;
	goto L2036
L2036:
	;
	v9345 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	v9346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9345)+5)))
	if v9346 != int32(1) {
		v9354 = v9345
		goto L2021
	} else {
		goto L2038
	}
L2037:
	;
	goto L2036
L2038:
	;
	v9349 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9345)+16)), uint32(v9349))
	goto L2026
L2039:
	;
	v9395 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v9399 = F_LWLockAcquire(m, v9395+int32(_a_F_StartupXLOG_262), int32(1))
	mBase = m.M
	v9400 = m.ExcPending
	if v9400 != 0 {
		goto L32
	} else {
		goto L2042
	}
L2040:
	;
	goto L2041
L2041:
	;
	v9571 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v9571)
	v9574 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v9579 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v9579 != 0 {
		goto L2064
	} else {
		goto L2065
	}
L2042:
	;
	v9402 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v9402 {
		goto L2043
	} else {
		goto L2044
	}
L2043:
	;
	v9406 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9413 = v9386
	v9414 = v9402
	v9415 = v9406
	v9434 = int64(0)
	goto L2046
L2044:
	;
	goto L2045
L2045:
	;
	v9531 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v9531+int32(_a_F_StartupXLOG_262))
	mBase = m.M
	v9535 = m.ExcPending
	if v9535 != 0 {
		goto L32
	} else {
		goto L2063
	}
L2046:
	;
	v9443 = v9415 + v9413*int32(288)
	v9444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9443)+4)))
	if v9444 != int32(1) {
		v9490 = v9414
		v9491 = v9415
		v9492 = v9434
		goto L2048
	} else {
		goto L2049
	}
L2047:
	;
	goto L2045
L2048:
	;
	v9494 = v9413 + int32(1)
	if v9494 < v9490 {
		v9413 = v9494
		v9414 = v9490
		v9415 = v9491
		v9434 = v9492
		goto L2046
	} else {
		goto L2062
	}
L2049:
	;
	v9447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9443)+201)))
	if v9447 == int32(0) {
		v9490 = v9414
		v9491 = v9415
		v9492 = v9434
		goto L2048
	} else {
		goto L2050
	}
L2050:
	;
	if v9434 == int64(0) {
		goto L2051
	} else {
		goto L2052
	}
L2051:
	;
	v9455 = m.G0
	v9456 = int32(16)
	v9457 = v9455 - v9456
	m.G0 = v9457
	F_gettimeofday(m, v9457)
	mBase = m.M
	v9460 = *(*int64)(unsafe.Add(mBase, uint32(v9457)))
	v9461 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9457)+8)))
	m.G0 = v9457 + v9456
	goto L2054
L2052:
	;
	v9470 = v9434
	goto L2053
L2053:
	;
	v9473 = base.AtomicRmwXchg32(m, v9443, int32(0), int32(1))
	if v9473 != 0 {
		goto L2055
	} else {
		goto L2056
	}
L2054:
	;
	v9470 = v9461 + v9460*int64(1000000) - int64(946684800000000)
	goto L2053
L2055:
	;
	F_s_lock(m, v9443, int32(_a_F_StartupXLOG_263), int32(251), int32(_a_F_StartupXLOG_264))
	mBase = m.M
	v9478 = m.ExcPending
	if v9478 != 0 {
		goto L32
	} else {
		goto L2058
	}
L2056:
	;
	goto L2057
L2057:
	;
	v9479 = *(*int32)(unsafe.Add(mBase, uint32(v9443)+112))
	if v9479 == int32(0) {
		goto L2059
	} else {
		goto L2060
	}
L2058:
	;
	goto L2057
L2059:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9443)+272)) = v9470
	goto L2061
L2060:
	;
	goto L2061
L2061:
	;
	v9483 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9443))), uint32(v9483))
	v9487 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	v9489 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9490 = v9487
	v9491 = v9489
	v9492 = v9470
	goto L2048
L2062:
	;
	goto L2047
L2063:
	;
	goto L2041
L2064:
	;
	v9580 = v9574 + int32(40)
	goto L2066
L2065:
	;
	v9580 = int32(_a_F_StartupXLOG_216)
	goto L2066
L2066:
	;
	v9581 = *(*int32)(unsafe.Add(mBase, uint32(v9580)))
	v9583 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	if v9579 != 0 {
		goto L2067
	} else {
		goto L2068
	}
L2067:
	;
	v9587 = v9574 + int32(24)
	goto L2069
L2068:
	;
	v9587 = int32(_a_F_StartupXLOG_265)
	goto L2069
L2069:
	;
	v9588 = *(*int64)(unsafe.Add(mBase, uint32(v9587)))
	F_XLogPrefetcherBeginRead(m, v9583, v9588)
	mBase = m.M
	v9590 = m.ExcPending
	if v9590 != 0 {
		goto L32
	} else {
		goto L2070
	}
L2070:
	;
	v9592 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v9595 = F_ReadRecord(m, v9592, int32(23), int32(0), v9581)
	mBase = m.M
	v9596 = m.ExcPending
	if v9596 != 0 {
		goto L32
	} else {
		goto L2071
	}
L2071:
	;
	v9598 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9599 = *(*int64)(unsafe.Add(mBase, uint32(v9598)+40))
	v9600 = *(*int32)(unsafe.Add(mBase, uint32(v9598)+1184))
	*(*int32)(unsafe.Add(mBase, uint32(v9237)+24)) = v9600
	v9603 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9603 != int32(1) {
		goto L2072
	} else {
		goto L2073
	}
L2072:
	;
	v9619 = v9599 & int64(8191)
	if v9619 != int64(0) {
		goto L2075
	} else {
		goto L2076
	}
L2073:
	;
	v9607 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v9607)
	v9610 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	if v9610 < v9607 {
		goto L2072
	} else {
		goto L2074
	}
L2074:
	;
	v9613 = F_close(m, v9610)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127])) = int32(-1)
	goto L2072
L2075:
	;
	v9622 = base.I32_wrap_i64(v9619)
	v9623 = F_palloc(m, v9622)
	mBase = m.M
	v9624 = m.ExcPending
	if v9624 != 0 {
		goto L32
	} else {
		goto L2078
	}
L2076:
	;
	v9632 = int32(0)
	v9633 = v9599
	goto L2077
L2077:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9237)+40)) = v9632
	*(*int64)(unsafe.Add(mBase, uint32(v9237)+32)) = v9633
	v9637 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	switch v9637 - int32(1) {
	case 0:
		goto L2088
	case 1:
		goto L2087
	case 2:
		goto L2085
	case 3:
		goto L2086
	case 4:
		goto L2084
	default:
		goto L2083
	}
L2078:
	;
	if v9622 != 0 {
		goto L2079
	} else {
		goto L2080
	}
L2079:
	;
	v9626 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9627 = *(*int32)(unsafe.Add(mBase, uint32(v9626)+128))
	base.MemoryCopy(m, v9623, v9627, v9622)
	goto L2081
L2080:
	;
	goto L2081
L2081:
	;
	v9632 = v9623
	v9633 = v9599 & int64(-8192)
	goto L2077
L2082:
	;
	v9722 = F_pstrdup(m, v9234-int32(-64))
	mBase = m.M
	v9723 = m.ExcPending
	if v9723 != 0 {
		goto L32
	} else {
		goto L2105
	}
L2083:
	;
	v9716 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_266), int32(0))
	mBase = m.M
	v9717 = m.ExcPending
	if v9717 != 0 {
		goto L32
	} else {
		goto L2104
	}
L2084:
	;
	v9709 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_267), int32(0))
	mBase = m.M
	v9710 = m.ExcPending
	if v9710 != 0 {
		goto L32
	} else {
		goto L2103
	}
L2085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9234)+48)) = int32(_a_F_StartupXLOG_234)
	v9702 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_268), v9234+int32(48))
	mBase = m.M
	v9703 = m.ExcPending
	if v9703 != 0 {
		goto L32
	} else {
		goto L2102
	}
L2086:
	;
	v9675 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[104]))
	*(*uint32)(unsafe.Add(mBase, uint32(v9234)+40)) = uint32(v9675)
	v9680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])))
	if v9680 != 0 {
		goto L2098
	} else {
		goto L2099
	}
L2087:
	;
	v9656 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])))
	v9658 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[103]))
	v9659 = F_timestamptz_to_str(m, v9658)
	mBase = m.M
	v9660 = m.ExcPending
	if v9660 != 0 {
		goto L32
	} else {
		goto L2093
	}
L2088:
	;
	v9641 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	*(*int32)(unsafe.Add(mBase, uint32(v9234)+4)) = v9641
	v9646 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])))
	if v9646 != 0 {
		goto L2089
	} else {
		goto L2090
	}
L2089:
	;
	v9647 = int32(_a_F_StartupXLOG_269)
	goto L2091
L2090:
	;
	v9647 = int32(_a_F_StartupXLOG_270)
	goto L2091
L2091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9234))) = v9647
	v9653 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_271), v9234)
	mBase = m.M
	v9654 = m.ExcPending
	if v9654 != 0 {
		goto L32
	} else {
		goto L2092
	}
L2092:
	;
	goto L2082
L2093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9234)+20)) = v9659
	if v9656 != 0 {
		goto L2094
	} else {
		goto L2095
	}
L2094:
	;
	v9664 = int32(_a_F_StartupXLOG_269)
	goto L2096
L2095:
	;
	v9664 = int32(_a_F_StartupXLOG_270)
	goto L2096
L2096:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9234)+16)) = v9664
	v9672 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_272), v9234+int32(16))
	mBase = m.M
	v9673 = m.ExcPending
	if v9673 != 0 {
		goto L32
	} else {
		goto L2097
	}
L2097:
	;
	goto L2082
L2098:
	;
	v9681 = int32(_a_F_StartupXLOG_269)
	goto L2100
L2099:
	;
	v9681 = int32(_a_F_StartupXLOG_270)
	goto L2100
L2100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9234)+32)) = v9681
	v9684 = int64(base.Ui64(v9675) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v9234)+36)) = uint32(v9684)
	v9692 = F_pg_snprintf(m, v9234-int32(-64), int32(200), int32(_a_F_StartupXLOG_273), v9234+int32(32))
	mBase = m.M
	v9693 = m.ExcPending
	if v9693 != 0 {
		goto L32
	} else {
		goto L2101
	}
L2101:
	;
	goto L2082
L2102:
	;
	goto L2082
L2103:
	;
	goto L2082
L2104:
	;
	goto L2082
L2105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9237)+16)) = v9599
	*(*int32)(unsafe.Add(mBase, uint32(v9237)+8)) = v9581
	*(*int64)(unsafe.Add(mBase, uint32(v9237))) = v9588
	*(*int32)(unsafe.Add(mBase, uint32(v9237)+64)) = v9722
	v9729 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51]))
	*(*int64)(unsafe.Add(mBase, uint32(v9237)+48)) = v9729
	v9732 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52]))
	*(*int64)(unsafe.Add(mBase, uint32(v9237)+56)) = v9732
	v9735 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9237)+68)) = uint8(v9735)
	v9738 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9237)+69)) = uint8(v9738)
	m.G0 = v9234 + int32(272)
	v9745 = *(*int32)(unsafe.Add(mBase, uint32(v9237)+24))
	v9748 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v9748 == int32(1) {
		goto L2106
	} else {
		goto L2107
	}
L2106:
	;
	v9752 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v9754 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[83]))
	if base.Ui64(v9754) <= base.Ui64(v9599) {
		goto L2110
	} else {
		goto L2111
	}
L2107:
	;
	goto L2108
L2108:
	;
	v9792 = int32(0)
	v9794 = F_PrescanPreparedTransactions(m, v9792, v9792)
	mBase = m.M
	v9795 = m.ExcPending
	if v9795 != 0 {
		goto L32
	} else {
		goto L2125
	}
L2109:
	;
	F_ResetUnloggedRelations(m, int32(2))
	mBase = m.M
	v9790 = m.ExcPending
	if v9790 != 0 {
		goto L32
	} else {
		goto L2124
	}
L2110:
	;
	v9756 = *(*int64)(unsafe.Add(mBase, uint32(v9752)+152))
	if v9756 == int64(0) {
		goto L2109
	} else {
		goto L2113
	}
L2111:
	;
	goto L2112
L2112:
	;
	v9760 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9760 == int32(0) {
		goto L2114
	} else {
		goto L2115
	}
L2113:
	;
	goto L2112
L2114:
	;
	v9763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9752)+168)))
	if v9763 != int32(1) {
		goto L2109
	} else {
		goto L2117
	}
L2115:
	;
	goto L2116
L2116:
	;
	v9766 = *(*int64)(unsafe.Add(mBase, uint32(v9752)+152))
	if v9766 != int64(0) {
		goto L13
	} else {
		goto L2118
	}
L2117:
	;
	goto L2116
L2118:
	;
	v9769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9752)+168)))
	if v9769 == int32(1) {
		goto L13
	} else {
		goto L2119
	}
L2119:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9775 = m.ExcPending
	if v9775 != 0 {
		goto L32
	} else {
		goto L2120
	}
L2120:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v9778 = m.ExcPending
	if v9778 != 0 {
		goto L32
	} else {
		goto L2121
	}
L2121:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_274), int32(0))
	mBase = m.M
	v9782 = m.ExcPending
	if v9782 != 0 {
		goto L32
	} else {
		goto L2122
	}
L2122:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_275), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v9787 = m.ExcPending
	if v9787 != 0 {
		goto L32
	} else {
		goto L2123
	}
L2123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2124:
	;
	goto L2108
L2125:
	;
	v9797 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v9801 = F_LWLockAcquire(m, v9797+int32(1152), int32(0))
	mBase = m.M
	v9802 = m.ExcPending
	if v9802 != 0 {
		goto L32
	} else {
		goto L2126
	}
L2126:
	;
	v9804 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v9805 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9804)+320)) = uint8(v9805)
	v9808 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v9808+int32(1152))
	mBase = m.M
	v9812 = m.ExcPending
	if v9812 != 0 {
		goto L32
	} else {
		goto L2127
	}
L2127:
	;
	v9814 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9814 != int32(1) {
		goto L2129
	} else {
		goto L2130
	}
L2128:
	;
	v10718 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10721 = base.AtomicRmwXchg32(m, v10718, int32(440), int32(1))
	if v10721 != 0 {
		goto L2317
	} else {
		goto L2318
	}
L2129:
	;
	v9817 = *(*int32)(unsafe.Add(mBase, uint32(v9237)+8))
	v10688 = v9817
	goto L2128
L2130:
	;
	goto L2131
L2131:
	;
	v9819 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v9820 = F_findNewestTimeLine(m, v9819)
	mBase = m.M
	v9821 = m.ExcPending
	if v9821 != 0 {
		goto L32
	} else {
		goto L2132
	}
L2132:
	;
	v9823 = v9820 + int32(1)
	v9826 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9827 = m.ExcPending
	if v9827 != 0 {
		goto L32
	} else {
		goto L2133
	}
L2133:
	;
	if v9826 != 0 {
		goto L2134
	} else {
		goto L2135
	}
L2134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3840)) = v9823
	F_errmsg(m, int32(_a_F_StartupXLOG_276), v39+int32(3840))
	mBase = m.M
	v9833 = m.ExcPending
	if v9833 != 0 {
		goto L32
	} else {
		goto L2137
	}
L2135:
	;
	goto L2136
L2136:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v9842 = m.ExcPending
	if v9842 != 0 {
		goto L32
	} else {
		goto L2139
	}
L2137:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_277), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v9838 = m.ExcPending
	if v9838 != 0 {
		goto L32
	} else {
		goto L2138
	}
L2138:
	;
	goto L2136
L2139:
	;
	v9846 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v9847 = base.I64_extend_i32_s(v9846)
	v9848 = base.I64_div_u_s(v9599-int64(1), v9847)
	v9849 = base.I64_div_u_s(v9599, v9847)
	if v9848 == v9849 {
		goto L2141
	} else {
		goto L2142
	}
L2140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3680)) = v9823
	v10157 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v10158 = base.I64_div_u_s(int64(4294967296), v10157)
	v10159 = base.I64_div_u_s(v9849, v10158)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3684)) = uint32(v10159)
	v10162 = v9849 - v10158*v10159
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3688)) = uint32(v10162)
	v10165 = v39 + int32(_a_F_StartupXLOG_1)
	v10170 = F_pg_snprintf(m, v10165, int32(64), int32(_a_F_StartupXLOG_278), v39+int32(3680))
	mBase = m.M
	v10171 = m.ExcPending
	if v10171 != 0 {
		goto L32
	} else {
		goto L2203
	}
L2141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3808)) = v9745
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[128]))) = v9848
	v9854 = base.I64_div_u_s(int64(4294967296), v9847)
	v9855 = base.I64_div_u_s(v9848, v9854)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3812)) = uint32(v9855)
	v9858 = v9848 - v9854*v9855
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3816)) = uint32(v9858)
	v9861 = v39 + int32(_a_F_StartupXLOG_4)
	v9866 = F_pg_snprintf(m, v9861, int32(1024), int32(_a_F_StartupXLOG_279), v39+int32(3808))
	mBase = m.M
	v9867 = m.ExcPending
	if v9867 != 0 {
		goto L32
	} else {
		goto L2144
	}
L2142:
	;
	goto L2143
L2143:
	;
	v10117 = F_XLogFileInit(m, v9849, v9823)
	mBase = m.M
	v10118 = m.ExcPending
	if v10118 != 0 {
		goto L32
	} else {
		goto L2201
	}
L2144:
	;
	v9869 = F_OpenTransientFile(m, v9861, int32(0))
	mBase = m.M
	v9870 = m.ExcPending
	if v9870 != 0 {
		goto L32
	} else {
		goto L2145
	}
L2145:
	;
	if v9869 < int32(0) {
		goto L12
	} else {
		goto L2146
	}
L2146:
	;
	v9873 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3792)) = v9873
	v9876 = v39 + int32(_a_F_StartupXLOG_3)
	v9881 = F_pg_snprintf(m, v9876, int32(1024), int32(_a_F_StartupXLOG_280), v39+int32(3792))
	mBase = m.M
	v9882 = m.ExcPending
	if v9882 != 0 {
		goto L32
	} else {
		goto L2147
	}
L2147:
	;
	v9883 = F_unlink(m, v9876)
	mBase = m.M
	v9886 = F_OpenTransientFile(m, v9876, int32(194))
	mBase = m.M
	v9887 = m.ExcPending
	if v9887 != 0 {
		goto L32
	} else {
		goto L2148
	}
L2148:
	;
	if v9886 < int32(0) {
		goto L11
	} else {
		goto L2149
	}
L2149:
	;
	v9891 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if int32(0) < v9891 {
		goto L2150
	} else {
		goto L2151
	}
L2150:
	;
	v9898 = int32(0)
	goto L2153
L2151:
	;
	goto L2152
L2152:
	;
	v10041 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10041))) = int32(167772231)
	v10046 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v10046 != int32(1) {
		v10060 = int32(0)
		goto L2176
	} else {
		goto L2177
	}
L2153:
	;
	v9932 = base.I32_wrap_i64(v9599)&(v9846-int32(1)) - v9898
	if base.Ui32(v9932) <= base.Ui32(int32(_a_F_StartupXLOG_281)) {
		goto L2155
	} else {
		goto L2156
	}
L2154:
	;
	goto L2152
L2155:
	;
	base.MemoryFill(m, v39+int32(_a_F_StartupXLOG_1), int32(0), int32(_a_F_StartupXLOG_58))
	goto L2157
L2156:
	;
	goto L2157
L2157:
	;
	if int32(0) < v9932 {
		goto L2158
	} else {
		goto L2159
	}
L2158:
	;
	v9943 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9943))) = int32(167772230)
	v9948 = int32(_a_F_StartupXLOG_58)
	if base.Ui32(v9948) <= base.Ui32(v9932) {
		goto L2161
	} else {
		goto L2162
	}
L2159:
	;
	goto L2160
L2160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v9988 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9988))) = int32(167772232)
	v9993 = int32(_a_F_StartupXLOG_58)
	v9994 = F_write(m, v9886, v39+int32(_a_F_StartupXLOG_1), v9993)
	mBase = m.M
	if v9994 != v9993 {
		goto L9
	} else {
		goto L2172
	}
L2161:
	;
	v9951 = v9948
	goto L2163
L2162:
	;
	v9951 = v9932
	goto L2163
L2163:
	;
	v9952 = F_read(m, v9869, v39+int32(_a_F_StartupXLOG_1), v9951)
	mBase = m.M
	if v9952 != v9951 {
		goto L2164
	} else {
		goto L2165
	}
L2164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9957 = m.ExcPending
	if v9957 != 0 {
		goto L32
	} else {
		goto L2167
	}
L2165:
	;
	goto L2166
L2166:
	;
	v9979 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9979))) = int32(0)
	goto L2160
L2167:
	;
	if v9952 < int32(0) {
		goto L10
	} else {
		goto L2168
	}
L2168:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v9962 = m.ExcPending
	if v9962 != 0 {
		goto L32
	} else {
		goto L2169
	}
L2169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3784)) = v9951
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3780)) = v9952
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3776)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v39+int32(3776))
	mBase = m.M
	v9972 = m.ExcPending
	if v9972 != 0 {
		goto L32
	} else {
		goto L2170
	}
L2170:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3486), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v9977 = m.ExcPending
	if v9977 != 0 {
		goto L32
	} else {
		goto L2171
	}
L2171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2172:
	;
	v9998 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9998))) = int32(0)
	v10002 = v9898 - int32(-8192)
	v10004 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v10002 < v10004 {
		v9898 = v10002
		goto L2153
	} else {
		goto L2173
	}
L2173:
	;
	goto L2154
L2174:
	;
	v10089 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10089))) = int32(0)
	v10092 = F_CloseTransientFile(m, v9886)
	mBase = m.M
	v10093 = m.ExcPending
	if v10093 != 0 {
		goto L32
	} else {
		goto L2192
	}
L2175:
	;
	if v10060 == int32(0) {
		goto L2174
	} else {
		goto L2182
	}
L2176:
	;
	goto L2175
L2177:
	;
	goto L2178
L2178:
	;
	v10051 = F_fsync(m, v9886)
	mBase = m.M
	if v10051 != int32(-1) {
		v10060 = v10051
		goto L2176
	} else {
		goto L2180
	}
L2179:
	;
	v10060 = int32(-1)
	goto L2176
L2180:
	;
	v10055 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10055 == int32(27) {
		goto L2178
	} else {
		goto L2181
	}
L2181:
	;
	goto L2179
L2182:
	;
	v10066 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[129])))
	if v10066 != 0 {
		goto L2184
	} else {
		goto L2185
	}
L2183:
	;
	v10069 = F_errstart(m, v10067, int32(0))
	mBase = m.M
	v10070 = m.ExcPending
	if v10070 != 0 {
		goto L32
	} else {
		goto L2187
	}
L2184:
	;
	v10067 = int32(21)
	goto L2186
L2185:
	;
	v10067 = int32(23)
	goto L2186
L2186:
	;
	goto L2183
L2187:
	;
	if v10069 == int32(0) {
		goto L2174
	} else {
		goto L2188
	}
L2188:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10074 = m.ExcPending
	if v10074 != 0 {
		goto L32
	} else {
		goto L2189
	}
L2189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3728)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v39+int32(3728))
	mBase = m.M
	v10082 = m.ExcPending
	if v10082 != 0 {
		goto L32
	} else {
		goto L2190
	}
L2190:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3514), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v10087 = m.ExcPending
	if v10087 != 0 {
		goto L32
	} else {
		goto L2191
	}
L2191:
	;
	goto L2174
L2192:
	;
	if v10092 != 0 {
		goto L8
	} else {
		goto L2193
	}
L2193:
	;
	v10094 = F_CloseTransientFile(m, v9869)
	mBase = m.M
	v10095 = m.ExcPending
	if v10095 != 0 {
		goto L32
	} else {
		goto L2194
	}
L2194:
	;
	if v10094 != 0 {
		goto L7
	} else {
		goto L2195
	}
L2195:
	;
	v10102 = F_InstallXLogFileSegment(m, v39+int32(_a_F_StartupXLOG_283), v39+int32(_a_F_StartupXLOG_3), int32(0), int64(0), v9823)
	mBase = m.M
	v10103 = m.ExcPending
	if v10103 != 0 {
		goto L32
	} else {
		goto L2196
	}
L2196:
	;
	if v10102 != 0 {
		goto L2140
	} else {
		goto L2197
	}
L2197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10107 = m.ExcPending
	if v10107 != 0 {
		goto L32
	} else {
		goto L2198
	}
L2198:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_284), int32(0))
	mBase = m.M
	v10111 = m.ExcPending
	if v10111 != 0 {
		goto L32
	} else {
		goto L2199
	}
L2199:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3531), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v10116 = m.ExcPending
	if v10116 != 0 {
		goto L32
	} else {
		goto L2200
	}
L2200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2201:
	;
	v10119 = F_close(m, v10117)
	mBase = m.M
	if v10119 != 0 {
		goto L6
	} else {
		goto L2202
	}
L2202:
	;
	goto L2140
L2203:
	;
	F_XLogArchiveCleanup(m, v10165)
	mBase = m.M
	v10173 = m.ExcPending
	if v10173 != 0 {
		goto L32
	} else {
		goto L2204
	}
L2204:
	;
	v10174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9237)+68)))
	if v10174 == int32(1) {
		goto L2205
	} else {
		goto L2206
	}
L2205:
	;
	v10179 = F_durable_unlink(m, int32(_a_F_StartupXLOG_47), int32(22))
	mBase = m.M
	v10180 = m.ExcPending
	if v10180 != 0 {
		goto L32
	} else {
		goto L2208
	}
L2206:
	;
	goto L2207
L2207:
	;
	v10181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9237)+69)))
	if v10181 == int32(1) {
		goto L2209
	} else {
		goto L2210
	}
L2208:
	;
	goto L2207
L2209:
	;
	v10186 = F_durable_unlink(m, int32(_a_F_StartupXLOG_48), int32(22))
	mBase = m.M
	v10187 = m.ExcPending
	if v10187 != 0 {
		goto L32
	} else {
		goto L2212
	}
L2210:
	;
	goto L2211
L2211:
	;
	v10189 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v10190 = *(*int32)(unsafe.Add(mBase, uint32(v9237)+64))
	v10191 = m.G0
	v10193 = v10191 - int32(_a_F_StartupXLOG_285)
	m.G0 = v10193
	v10195 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+224)) = v10195
	v10198 = v10193 + int32(_a_F_StartupXLOG_286)
	v10203 = F_pg_snprintf(m, v10198, int32(1024), int32(_a_F_StartupXLOG_280), v10193+int32(224))
	mBase = m.M
	v10204 = m.ExcPending
	if v10204 != 0 {
		goto L32
	} else {
		goto L2213
	}
L2212:
	;
	goto L2211
L2213:
	;
	v10205 = F_unlink(m, v10198)
	mBase = m.M
	v10207 = F_OpenTransientFile(m, v10198, int32(194))
	mBase = m.M
	v10208 = m.ExcPending
	if v10208 != 0 {
		goto L32
	} else {
		goto L2220
	}
L2214:
	;
	v10670 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10671 = m.ExcPending
	if v10671 != 0 {
		goto L32
	} else {
		goto L2313
	}
L2215:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10652 = m.ExcPending
	if v10652 != 0 {
		goto L32
	} else {
		goto L2309
	}
L2216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+112)) = v10190
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+100)) = v10189
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+96)) = v10423
	*(*uint32)(unsafe.Add(mBase, uint32(v10193)+108)) = uint32(v9599)
	v10459 = int64(base.Ui64(v9599) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10193)+104)) = uint32(v10459)
	v10462 = v10193 + int32(240)
	v10467 = F_pg_snprintf(m, v10462, int32(_a_F_StartupXLOG_58), int32(_a_F_StartupXLOG_287), v10193+int32(96))
	mBase = m.M
	v10468 = m.ExcPending
	if v10468 != 0 {
		goto L32
	} else {
		goto L2266
	}
L2217:
	;
	v10397 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10397 == int32(44) {
		goto L2259
	} else {
		goto L2260
	}
L2218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10380 = m.ExcPending
	if v10380 != 0 {
		goto L32
	} else {
		goto L2255
	}
L2219:
	;
	v10351 = int32(_a_F_StartupXLOG_2)
	v10352 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10354 = v10193 + int32(_a_F_StartupXLOG_286)
	v10355 = F_unlink(m, v10354)
	mBase = m.M
	if v10352 != 0 {
		goto L2248
	} else {
		goto L2249
	}
L2220:
	;
	if int32(0) <= v10207 {
		goto L2221
	} else {
		goto L2222
	}
L2221:
	;
	v10212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v10212 == int32(1) {
		goto L2225
	} else {
		goto L2226
	}
L2222:
	;
	goto L2223
L2223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10337 = m.ExcPending
	if v10337 != 0 {
		goto L32
	} else {
		goto L2244
	}
L2224:
	;
	v10244 = F_OpenTransientFile(m, v10193+int32(_a_F_StartupXLOG_288), int32(0))
	mBase = m.M
	v10245 = m.ExcPending
	if v10245 != 0 {
		goto L32
	} else {
		goto L2231
	}
L2225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+192)) = v10189
	v10217 = v10193 + int32(_a_F_StartupXLOG_289)
	v10222 = F_pg_snprintf(m, v10217, int32(64), int32(_a_F_StartupXLOG_290), v10193+int32(192))
	mBase = m.M
	v10223 = m.ExcPending
	if v10223 != 0 {
		goto L32
	} else {
		goto L2228
	}
L2226:
	;
	goto L2227
L2227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+208)) = v10189
	v10238 = F_pg_snprintf(m, v10193+int32(_a_F_StartupXLOG_288), int32(1024), int32(_a_F_StartupXLOG_291), v10193+int32(208))
	mBase = m.M
	v10239 = m.ExcPending
	if v10239 != 0 {
		goto L32
	} else {
		goto L2230
	}
L2228:
	;
	v10229 = F_RestoreArchivedFile(m, v10193+int32(_a_F_StartupXLOG_288), v10217, int32(_a_F_StartupXLOG_292), int64(0), int32(0))
	mBase = m.M
	v10230 = m.ExcPending
	if v10230 != 0 {
		goto L32
	} else {
		goto L2229
	}
L2229:
	;
	goto L2224
L2230:
	;
	goto L2224
L2231:
	;
	if v10244 < int32(0) {
		goto L2217
	} else {
		goto L2232
	}
L2232:
	;
	v10249 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10249
	v10251 = int32(_a_F_StartupXLOG_143)
	v10252 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10252))) = int32(167772219)
	v10258 = F_read(m, v10244, v10193+int32(240), int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v10260 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10260))) = v10249
	if v10258 < v10249 {
		goto L2215
	} else {
		goto L2233
	}
L2233:
	;
	v10268 = v10258
	goto L2234
L2234:
	;
	v10300 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10300 != 0 {
		goto L2215
	} else {
		goto L2236
	}
L2235:
	;
	v10331 = F_CloseTransientFile(m, v10244)
	mBase = m.M
	v10332 = m.ExcPending
	if v10332 != 0 {
		goto L32
	} else {
		goto L2242
	}
L2236:
	;
	if v10268 != 0 {
		goto L2237
	} else {
		goto L2238
	}
L2237:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10305 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10305))) = int32(167772221)
	v10309 = v10193 + int32(240)
	v10310 = F_write(m, v10207, v10309, v10268)
	mBase = m.M
	if v10310 != v10268 {
		goto L2219
	} else {
		goto L2240
	}
L2238:
	;
	goto L2239
L2239:
	;
	goto L2235
L2240:
	;
	v10312 = int32(_a_F_StartupXLOG_143)
	v10313 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	v10314 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10313))) = v10314
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10314
	v10320 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10320))) = int32(167772219)
	v10324 = F_read(m, v10244, v10309, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v10326 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10326))) = v10314
	if v10314 <= v10324 {
		v10268 = v10324
		goto L2234
	} else {
		goto L2241
	}
L2241:
	;
	goto L2215
L2242:
	;
	if v10331 != 0 {
		goto L2218
	} else {
		goto L2243
	}
L2243:
	;
	v10423 = int32(_a_F_StartupXLOG_293)
	goto L2216
L2244:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10339 = m.ExcPending
	if v10339 != 0 {
		goto L32
	} else {
		goto L2245
	}
L2245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193))) = v10193 + int32(_a_F_StartupXLOG_286)
	F_errmsg(m, int32(_a_F_StartupXLOG_294), v10193)
	mBase = m.M
	v10345 = m.ExcPending
	if v10345 != 0 {
		goto L32
	} else {
		goto L2246
	}
L2246:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(329), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10350 = m.ExcPending
	if v10350 != 0 {
		goto L32
	} else {
		goto L2247
	}
L2247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2248:
	;
	v10358 = v10352
	goto L2250
L2249:
	;
	v10358 = int32(51)
	goto L2250
L2250:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10358
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10363 = m.ExcPending
	if v10363 != 0 {
		goto L32
	} else {
		goto L2251
	}
L2251:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10365 = m.ExcPending
	if v10365 != 0 {
		goto L32
	} else {
		goto L2252
	}
L2252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+176)) = v10354
	F_errmsg(m, int32(_a_F_StartupXLOG_297), v10193+int32(176))
	mBase = m.M
	v10371 = m.ExcPending
	if v10371 != 0 {
		goto L32
	} else {
		goto L2253
	}
L2253:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(384), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10376 = m.ExcPending
	if v10376 != 0 {
		goto L32
	} else {
		goto L2254
	}
L2254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2255:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10382 = m.ExcPending
	if v10382 != 0 {
		goto L32
	} else {
		goto L2256
	}
L2256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+160)) = v10193 + int32(_a_F_StartupXLOG_288)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v10193+int32(160))
	mBase = m.M
	v10390 = m.ExcPending
	if v10390 != 0 {
		goto L32
	} else {
		goto L2257
	}
L2257:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(392), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10395 = m.ExcPending
	if v10395 != 0 {
		goto L32
	} else {
		goto L2258
	}
L2258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2259:
	;
	v10423 = int32(_a_F_StartupXLOG_298)
	goto L2216
L2260:
	;
	goto L2261
L2261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10404 = m.ExcPending
	if v10404 != 0 {
		goto L32
	} else {
		goto L2262
	}
L2262:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10406 = m.ExcPending
	if v10406 != 0 {
		goto L32
	} else {
		goto L2263
	}
L2263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+128)) = v10193 + int32(_a_F_StartupXLOG_288)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v10193+int32(128))
	mBase = m.M
	v10414 = m.ExcPending
	if v10414 != 0 {
		goto L32
	} else {
		goto L2264
	}
L2264:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(348), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10419 = m.ExcPending
	if v10419 != 0 {
		goto L32
	} else {
		goto L2265
	}
L2265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2266:
	;
	v10469 = F_strlen(m, v10462)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10474 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10474))) = int32(167772221)
	v10477 = F_write(m, v10207, v10462, v10469)
	mBase = m.M
	if v10477 == v10469 {
		goto L2268
	} else {
		goto L2269
	}
L2267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10599 = m.ExcPending
	if v10599 != 0 {
		goto L32
	} else {
		goto L2305
	}
L2268:
	;
	v10479 = int32(_a_F_StartupXLOG_143)
	v10480 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	v10481 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10480))) = v10481
	v10484 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10484))) = int32(167772220)
	v10489 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v10489 != int32(1) {
		v10503 = v10481
		goto L2273
	} else {
		goto L2274
	}
L2269:
	;
	goto L2270
L2270:
	;
	v10570 = int32(_a_F_StartupXLOG_2)
	v10571 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10573 = v10193 + int32(_a_F_StartupXLOG_286)
	v10574 = F_unlink(m, v10573)
	mBase = m.M
	if v10571 != 0 {
		goto L2298
	} else {
		goto L2299
	}
L2271:
	;
	v10532 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10532))) = int32(0)
	v10535 = F_CloseTransientFile(m, v10207)
	mBase = m.M
	v10536 = m.ExcPending
	if v10536 != 0 {
		goto L32
	} else {
		goto L2289
	}
L2272:
	;
	if v10503 == int32(0) {
		goto L2271
	} else {
		goto L2279
	}
L2273:
	;
	goto L2272
L2274:
	;
	goto L2275
L2275:
	;
	v10494 = F_fsync(m, v10207)
	mBase = m.M
	if v10494 != int32(-1) {
		v10503 = v10494
		goto L2273
	} else {
		goto L2277
	}
L2276:
	;
	v10503 = int32(-1)
	goto L2273
L2277:
	;
	v10498 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10498 == int32(27) {
		goto L2275
	} else {
		goto L2278
	}
L2278:
	;
	goto L2276
L2279:
	;
	v10509 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[129])))
	if v10509 != 0 {
		goto L2281
	} else {
		goto L2282
	}
L2280:
	;
	v10512 = F_errstart(m, v10510, int32(0))
	mBase = m.M
	v10513 = m.ExcPending
	if v10513 != 0 {
		goto L32
	} else {
		goto L2284
	}
L2281:
	;
	v10510 = int32(21)
	goto L2283
L2282:
	;
	v10510 = int32(23)
	goto L2283
L2283:
	;
	goto L2280
L2284:
	;
	if v10512 == int32(0) {
		goto L2271
	} else {
		goto L2285
	}
L2285:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10517 = m.ExcPending
	if v10517 != 0 {
		goto L32
	} else {
		goto L2286
	}
L2286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+64)) = v10193 + int32(_a_F_StartupXLOG_286)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v10193-int32(-64))
	mBase = m.M
	v10525 = m.ExcPending
	if v10525 != 0 {
		goto L32
	} else {
		goto L2287
	}
L2287:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(432), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10530 = m.ExcPending
	if v10530 != 0 {
		goto L32
	} else {
		goto L2288
	}
L2288:
	;
	goto L2271
L2289:
	;
	if v10535 != 0 {
		goto L2267
	} else {
		goto L2290
	}
L2290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+32)) = v9823
	v10539 = v10193 + int32(_a_F_StartupXLOG_288)
	v10544 = F_pg_snprintf(m, v10539, int32(1024), int32(_a_F_StartupXLOG_291), v10193+int32(32))
	mBase = m.M
	v10545 = m.ExcPending
	if v10545 != 0 {
		goto L32
	} else {
		goto L2291
	}
L2291:
	;
	v10549 = F_durable_rename(m, v10193+int32(_a_F_StartupXLOG_286), v10539, int32(21))
	mBase = m.M
	v10550 = m.ExcPending
	if v10550 != 0 {
		goto L32
	} else {
		goto L2292
	}
L2292:
	;
	v10552 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[130]))
	if int32(0) < v10552 {
		goto L2293
	} else {
		goto L2294
	}
L2293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+16)) = v9823
	v10557 = v10193 + int32(_a_F_StartupXLOG_289)
	v10562 = F_pg_snprintf(m, v10557, int32(64), int32(_a_F_StartupXLOG_290), v10193+int32(16))
	mBase = m.M
	v10563 = m.ExcPending
	if v10563 != 0 {
		goto L32
	} else {
		goto L2296
	}
L2294:
	;
	goto L2295
L2295:
	;
	m.G0 = v10193 + int32(_a_F_StartupXLOG_285)
	goto L2214
L2296:
	;
	F_XLogArchiveNotify(m, v10557)
	mBase = m.M
	v10565 = m.ExcPending
	if v10565 != 0 {
		goto L32
	} else {
		goto L2297
	}
L2297:
	;
	goto L2295
L2298:
	;
	v10577 = v10571
	goto L2300
L2299:
	;
	v10577 = int32(51)
	goto L2300
L2300:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10577
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10582 = m.ExcPending
	if v10582 != 0 {
		goto L32
	} else {
		goto L2301
	}
L2301:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10584 = m.ExcPending
	if v10584 != 0 {
		goto L32
	} else {
		goto L2302
	}
L2302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+80)) = v10573
	F_errmsg(m, int32(_a_F_StartupXLOG_297), v10193+int32(80))
	mBase = m.M
	v10590 = m.ExcPending
	if v10590 != 0 {
		goto L32
	} else {
		goto L2303
	}
L2303:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(424), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10595 = m.ExcPending
	if v10595 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v10601 = m.ExcPending
	if v10601 != 0 {
		goto L32
	} else {
		goto L2306
	}
L2306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+48)) = v10193 + int32(_a_F_StartupXLOG_286)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v10193+int32(48))
	mBase = m.M
	v10609 = m.ExcPending
	if v10609 != 0 {
		goto L32
	} else {
		goto L2307
	}
L2307:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(438), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10614 = m.ExcPending
	if v10614 != 0 {
		goto L32
	} else {
		goto L2308
	}
L2308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2309:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10654 = m.ExcPending
	if v10654 != 0 {
		goto L32
	} else {
		goto L2310
	}
L2310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10193)+144)) = v10193 + int32(_a_F_StartupXLOG_288)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v10193+int32(144))
	mBase = m.M
	v10662 = m.ExcPending
	if v10662 != 0 {
		goto L32
	} else {
		goto L2311
	}
L2311:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_295), int32(362), int32(_a_F_StartupXLOG_296))
	mBase = m.M
	v10667 = m.ExcPending
	if v10667 != 0 {
		goto L32
	} else {
		goto L2312
	}
L2312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2313:
	;
	if v10670 == int32(0) {
		v10688 = v9823
		goto L2128
	} else {
		goto L2314
	}
L2314:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_299), int32(0))
	mBase = m.M
	v10677 = m.ExcPending
	if v10677 != 0 {
		goto L32
	} else {
		goto L2315
	}
L2315:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_300), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10682 = m.ExcPending
	if v10682 != 0 {
		goto L32
	} else {
		goto L2316
	}
L2316:
	;
	v10688 = v9823
	goto L2128
L2317:
	;
	v10723 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v10723+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_301), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10730 = m.ExcPending
	if v10730 != 0 {
		goto L32
	} else {
		goto L2320
	}
L2318:
	;
	goto L2319
L2319:
	;
	v10732 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v10732)+308)) = v10688
	v10734 = *(*int32)(unsafe.Add(mBase, uint32(v9237)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10732)+312)) = v10734
	v10736 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v10732)+440)), uint32(v10736))
	if v9732 == int64(0) {
		goto L2321
	} else {
		goto L2322
	}
L2320:
	;
	goto L2319
L2321:
	;
	v10741 = v9599
	goto L2323
L2322:
	;
	v10741 = v9732
	goto L2323
L2323:
	;
	v10742 = *(*int64)(unsafe.Add(mBase, uint32(v9237)))
	v10745 = base.I32_wrap_i64(v10742) & int32(_a_F_StartupXLOG_281)
	v10747 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v10748 = base.I64_extend_i32_s(v10747)
	v10749 = base.I64_div_u_s(v10742, v10748)
	v10752 = base.I64_extend_i32_s(v10747 - int32(1))
	v10753 = v10742 & v10752
	if v10753&int64(35184372080640) == int64(0) {
		goto L2325
	} else {
		goto L2326
	}
L2324:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10732)+16)) = v10791
	v10793 = base.I64_div_u_s(v10741, v10748)
	v10796 = base.I32_wrap_i64(v10741) & int32(_a_F_StartupXLOG_281)
	v10797 = v10741 & v10752
	if v10797&int64(35184372080640) == int64(0) {
		goto L2331
	} else {
		goto L2332
	}
L2325:
	;
	v10759 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[131]))
	v10761 = v10749 * base.I64_extend_i32_s(v10759)
	if v10745 == int32(0) {
		v10789 = v10759
		v10791 = v10761
		goto L2324
	} else {
		goto L2328
	}
L2326:
	;
	goto L2327
L2327:
	;
	v10769 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[131]))
	v10782 = v10749*base.I64_extend_i32_s(v10769) + (int64(base.Ui64(v10753)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v10745 == int32(0) {
		v10789 = v10769
		v10791 = v10782
		goto L2324
	} else {
		goto L2329
	}
L2328:
	;
	v10789 = v10759
	v10791 = v10761 + base.I64_extend_i32_u(v10745-int32(40))
	goto L2324
L2329:
	;
	v10789 = v10769
	v10791 = v10782 + base.I64_extend_i32_u(v10745-int32(24))
	goto L2324
L2330:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10732)+8)) = v10830
	if v10741&int64(8191) != int64(0) {
		goto L2336
	} else {
		goto L2337
	}
L2331:
	;
	v10803 = v10793 * base.I64_extend_i32_s(v10789)
	if v10796 == int32(0) {
		v10830 = v10803
		goto L2330
	} else {
		goto L2334
	}
L2332:
	;
	goto L2333
L2333:
	;
	v10822 = v10793*base.I64_extend_i32_s(v10789) + (int64(base.Ui64(v10797)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v10796 == int32(0) {
		v10830 = v10822
		goto L2330
	} else {
		goto L2335
	}
L2334:
	;
	v10830 = v10803 + base.I64_extend_i32_u(v10796-int32(40))
	goto L2330
L2335:
	;
	v10830 = v10822 + base.I64_extend_i32_u(v10796-int32(24))
	goto L2330
L2336:
	;
	v10836 = *(*int32)(unsafe.Add(mBase, uint32(v10732)+296))
	v10839 = *(*int32)(unsafe.Add(mBase, uint32(v10732)+304))
	v10843 = base.I64_rem_u_s(int64(base.Ui64(v10741)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v10839+int32(1)))
	v10844 = base.I32_wrap_i64(v10843)
	v10847 = v10836 + v10844<<(uint(int32(13))%32)
	v10848 = *(*int64)(unsafe.Add(mBase, uint32(v9237)+32))
	v10850 = base.I32_wrap_i64(v10741 - v10848)
	if v10850 != 0 {
		goto L2339
	} else {
		goto L2340
	}
L2337:
	;
	v10872 = v10732
	v10876 = v10741
	goto L2338
L2338:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10872)+288)) = v10876
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[132])) = v10741
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[133])) = v10741
	v10883 = base.AtomicRmwXchg64(m, v10872, int32(264), v10741)
	v10884 = int32(_a_F_StartupXLOG_302)
	v10885 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10887 = base.AtomicRmwXchg64(m, v10885, int32(272), v10741)
	v10889 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10891 = base.AtomicRmwXchg64(m, v10889, int32(280), v10741)
	v10893 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10893)+192)) = v10741
	*(*int64)(unsafe.Add(mBase, uint32(v10893)+184)) = v10741
	v10896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10893)+320)))
	if v10896 != int32(1) {
		goto L2345
	} else {
		goto L2346
	}
L2339:
	;
	v10851 = *(*int32)(unsafe.Add(mBase, uint32(v9237)+40))
	base.MemoryCopy(m, v10847, v10851, v10850)
	goto L2341
L2340:
	;
	goto L2341
L2341:
	;
	v10854 = int32(_a_F_StartupXLOG_58) - v10850
	if v10854 != 0 {
		goto L2342
	} else {
		goto L2343
	}
L2342:
	;
	base.MemoryFill(m, v10850+v10847, int32(0), v10854)
	goto L2344
L2343:
	;
	goto L2344
L2344:
	;
	v10858 = *(*int32)(unsafe.Add(mBase, uint32(v10732)+300))
	v10862 = *(*int64)(unsafe.Add(mBase, uint32(v9237)+32))
	v10863 = int64(-8192)
	v10866 = base.AtomicRmwXchg64(m, v10858+v10844<<(uint(int32(3))%32), int32(0), v10862-v10863)
	v10867 = *(*int64)(unsafe.Add(mBase, uint32(v9237)+32))
	v10871 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10872 = v10871
	v10876 = v10867 - v10863
	goto L2338
L2345:
	;
	v10938 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v10938)
	v10940 = F_time(m)
	mBase = m.M
	v10942 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10942)+256)) = v10741
	*(*int64)(unsafe.Add(mBase, uint32(v10942)+248)) = v10940
	v10946 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v10950 = F_LWLockAcquire(m, v10946+int32(512), v10938)
	mBase = m.M
	v10951 = m.ExcPending
	if v10951 != 0 {
		goto L32
	} else {
		goto L2353
	}
L2346:
	;
	v10900 = v10741 - int64(1)
	v10902 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if base.Ui64(v10900&base.I64_extend_i32_s(v10902-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v10902), float64(0.75))))) {
		goto L2345
	} else {
		goto L2347
	}
L2347:
	;
	v10914 = base.I64_div_u_s(v10900, base.I64_extend_i32_s(v10902))
	v10921 = F_XLogFileInitInternal(m, v10914+int64(1), v10688, v39+int32(_a_F_StartupXLOG_4), v39+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v10922 = m.ExcPending
	if v10922 != 0 {
		goto L32
	} else {
		goto L2348
	}
L2348:
	;
	if int32(0) <= v10921 {
		goto L2349
	} else {
		goto L2350
	}
L2349:
	;
	v10925 = F_close(m, v10921)
	mBase = m.M
	goto L2351
L2350:
	;
	goto L2351
L2351:
	;
	v10926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[96]))))
	if v10926 != int32(1) {
		goto L2345
	} else {
		goto L2352
	}
L2352:
	;
	v10929 = int32(_a_F_StartupXLOG_303)
	v10931 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[134]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[134])) = v10931 + int32(1)
	goto L2345
L2353:
	;
	v10953 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v10954 = *(*int64)(unsafe.Add(mBase, uint32(v10953)+8))
	v10956 = v10954 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10953)+48)) = v10956
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v10956)))|base.B2i32(base.Ui64(v10956) < base.Ui64(int64(3))) == int32(0) {
		goto L2354
	} else {
		goto L2355
	}
L2354:
	;
	v10993 = v10956
	goto L2357
L2355:
	;
	goto L2356
L2356:
	;
	v11041 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11041+int32(512))
	mBase = m.M
	v11045 = m.ExcPending
	if v11045 != 0 {
		goto L32
	} else {
		goto L2360
	}
L2357:
	;
	v11001 = v10993 - int64(1)
	if base.Ui32(base.I32_wrap_i64(v11001)) < base.Ui32(int32(3)) {
		v10993 = v11001
		goto L2357
	} else {
		goto L2359
	}
L2358:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10953)+48)) = v11001
	goto L2356
L2359:
	;
	goto L2358
L2360:
	;
	v11047 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v11047 == int32(0) {
		goto L2361
	} else {
		goto L2362
	}
L2361:
	;
	F_StartupSUBTRANS(m, v9794)
	mBase = m.M
	v11051 = m.ExcPending
	if v11051 != 0 {
		goto L32
	} else {
		goto L2364
	}
L2362:
	;
	goto L2363
L2363:
	;
	v11053 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11054 = *(*int32)(unsafe.Add(mBase, uint32(v11053)+28))
	v11056 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v11057 = *(*int64)(unsafe.Add(mBase, uint32(v11056)+8))
	v11061 = int64(base.Ui64(v11057)>>(uint(int64(15))%64)) & int64(131071)
	v11064 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[135])))
	v11065 = base.I32_rem_u_s(base.I32_wrap_i64(v11061), v11064)
	v11068 = v11054 + v11065<<(uint(int32(7))%32)
	v11070 = F_LWLockAcquire(m, v11068, int32(0))
	mBase = m.M
	v11071 = m.ExcPending
	if v11071 != 0 {
		goto L32
	} else {
		goto L2365
	}
L2364:
	;
	goto L2363
L2365:
	;
	v11072 = base.I32_wrap_i64(v11057)
	v11074 = v11072 & int32(_a_F_StartupXLOG_304)
	if v11074 != 0 {
		goto L2366
	} else {
		goto L2367
	}
L2366:
	;
	v11077 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_305), v11061, int32(0), v11072)
	mBase = m.M
	v11078 = m.ExcPending
	if v11078 != 0 {
		goto L32
	} else {
		goto L2369
	}
L2367:
	;
	goto L2368
L2368:
	;
	F_LWLockRelease(m, v11068)
	mBase = m.M
	v11155 = m.ExcPending
	if v11155 != 0 {
		goto L32
	} else {
		goto L2380
	}
L2369:
	;
	v11080 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11081 = *(*int32)(unsafe.Add(mBase, uint32(v11080)+4))
	v11082 = int32(2)
	v11085 = *(*int32)(unsafe.Add(mBase, uint32(v11081+v11077<<(uint(v11082)%32))))
	v11087 = int32(base.Ui32(v11074) >> (uint(v11082) % 32))
	v11088 = v11085 + v11087
	v11089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11088))))
	v11090 = int32(-1)
	v11091 = int32(1)
	v11098 = v11089 & (v11090<<(uint(v11072<<(uint(v11091)%32)&int32(6))%32) ^ v11090)
	*(*uint8)(unsafe.Add(mBase, uint32(v11088))) = uint8(v11098)
	v11101 = v11088 + v11091
	v11102 = int32(3)
	v11105 = v11087 ^ int32(_a_F_StartupXLOG_281)
	if v11101&v11102|v11105&v11102|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v11105)) == int32(0) {
		goto L2371
	} else {
		goto L2372
	}
L2370:
	;
	v11143 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11144 = *(*int32)(unsafe.Add(mBase, uint32(v11143)+12))
	v11146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11144+v11077))) = uint8(v11146)
	goto L2368
L2371:
	;
	if v11087 == int32(_a_F_StartupXLOG_281) {
		goto L2370
	} else {
		goto L2374
	}
L2372:
	;
	goto L2373
L2373:
	;
	if v11105 == int32(0) {
		goto L2370
	} else {
		goto L2379
	}
L2374:
	;
	v11119 = v11105 + v11085 + v11087 + int32(1)
	v11121 = v11088 + int32(5)
	if base.Ui32(v11121) < base.Ui32(v11119) {
		goto L2375
	} else {
		goto L2376
	}
L2375:
	;
	v11123 = v11119
	goto L2377
L2376:
	;
	v11123 = v11121
	goto L2377
L2377:
	;
	v11130 = (v11123-v11088-int32(2))&int32(-4) + int32(4)
	if v11130 == int32(0) {
		goto L2370
	} else {
		goto L2378
	}
L2378:
	;
	base.MemoryFill(m, v11101, int32(0), v11130)
	goto L2370
L2379:
	;
	base.MemoryFill(m, v11101, int32(0), v11105)
	goto L2370
L2380:
	;
	v11157 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11161 = F_LWLockAcquire(m, v11157+int32(1664), int32(1))
	mBase = m.M
	v11162 = m.ExcPending
	if v11162 != 0 {
		goto L32
	} else {
		goto L2381
	}
L2381:
	;
	v11164 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v11165 = *(*int32)(unsafe.Add(mBase, uint32(v11164)+16))
	v11166 = *(*int32)(unsafe.Add(mBase, uint32(v11164)+12))
	v11167 = *(*int32)(unsafe.Add(mBase, uint32(v11164)+4))
	v11168 = *(*int32)(unsafe.Add(mBase, uint32(v11164)))
	v11170 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11170+int32(1664))
	mBase = m.M
	v11174 = m.ExcPending
	if v11174 != 0 {
		goto L32
	} else {
		goto L2382
	}
L2382:
	;
	v11175 = int32(_a_F_StartupXLOG_306)
	v11176 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11178 = int32(base.Ui32(v11168) >> (uint(int32(11)) % 32))
	v11179 = base.I64_extend_i32_u(v11178)
	v11181 = base.AtomicRmwXchg64(m, v11176, int32(48), v11179)
	v11183 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11184 = *(*int32)(unsafe.Add(mBase, uint32(v11183)+28))
	v11186 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[136])))
	v11187 = base.I32_rem_u_s(v11178, v11186)
	v11190 = v11184 + v11187<<(uint(int32(7))%32)
	v11192 = F_LWLockAcquire(m, v11190, int32(0))
	mBase = m.M
	v11193 = m.ExcPending
	if v11193 != 0 {
		goto L32
	} else {
		goto L2383
	}
L2383:
	;
	v11195 = v11168 & int32(2047)
	if v11195 == int32(0) {
		goto L2385
	} else {
		goto L2386
	}
L2384:
	;
	v11207 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11208 = *(*int32)(unsafe.Add(mBase, uint32(v11207)+4))
	v11209 = int32(2)
	v11212 = *(*int32)(unsafe.Add(mBase, uint32(v11208+v11205<<(uint(v11209)%32))))
	v11214 = v11195 << (uint(v11209) % 32)
	v11215 = v11212 + v11214
	*(*int32)(unsafe.Add(mBase, uint32(v11215))) = v11167
	if base.Ui32(int32(2045)) < base.Ui32(v11195-int32(1)) {
		goto L2390
	} else {
		goto L2391
	}
L2385:
	;
	v11199 = F_SimpleLruZeroPage(m, int32(_a_F_StartupXLOG_306), v11179)
	mBase = m.M
	v11200 = m.ExcPending
	if v11200 != 0 {
		goto L32
	} else {
		goto L2388
	}
L2386:
	;
	goto L2387
L2387:
	;
	v11203 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_306), v11179, int32(1), v11168)
	mBase = m.M
	v11204 = m.ExcPending
	if v11204 != 0 {
		goto L32
	} else {
		goto L2389
	}
L2388:
	;
	v11205 = v11199
	goto L2384
L2389:
	;
	v11205 = v11203
	goto L2384
L2390:
	;
	v11258 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11259 = *(*int32)(unsafe.Add(mBase, uint32(v11258)+12))
	v11261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11259+v11205))) = uint8(v11261)
	F_LWLockRelease(m, v11190)
	mBase = m.M
	v11264 = m.ExcPending
	if v11264 != 0 {
		goto L32
	} else {
		goto L2400
	}
L2391:
	;
	v11222 = v11215 + int32(4)
	if v11222&int32(3)|base.B2i32(base.Ui32(v11195) < base.Ui32(int32(1791))) == int32(0) {
		goto L2392
	} else {
		goto L2393
	}
L2392:
	;
	v11231 = v11215 + int32(8)
	v11233 = v11212 - int32(-8192)
	if base.Ui32(v11233) < base.Ui32(v11231) {
		goto L2395
	} else {
		goto L2396
	}
L2393:
	;
	goto L2394
L2394:
	;
	v11248 = int32(_a_F_StartupXLOG_307) - v11214
	if v11248 == int32(0) {
		goto L2390
	} else {
		goto L2399
	}
L2395:
	;
	v11235 = v11231
	goto L2397
L2396:
	;
	v11235 = v11233
	goto L2397
L2397:
	;
	v11242 = (v11235-v11215-int32(5))&int32(-4) + int32(4)
	if v11242 == int32(0) {
		goto L2390
	} else {
		goto L2398
	}
L2398:
	;
	base.MemoryFill(m, v11222, int32(0), v11242)
	goto L2390
L2399:
	;
	base.MemoryFill(m, v11222, int32(0), v11248)
	goto L2390
L2400:
	;
	v11266 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11268 = base.I32_div_u_s(v11167, int32(1636))
	v11269 = base.I64_extend_i32_u(v11268)
	v11271 = base.AtomicRmwXchg64(m, v11266, int32(48), v11269)
	v11273 = int32(base.Ui32(v11167) >> (uint(int32(2)) % 32))
	v11275 = base.I32_rem_u_s(v11273, int32(409))
	if v11275 != 0 {
		goto L2401
	} else {
		goto L2402
	}
L2401:
	;
	v11277 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11278 = *(*int32)(unsafe.Add(mBase, uint32(v11277)+28))
	v11280 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[137])))
	v11281 = base.I32_rem_u_s(v11268, v11280)
	v11284 = v11278 + v11281<<(uint(int32(7))%32)
	v11286 = F_LWLockAcquire(m, v11284, int32(0))
	mBase = m.M
	v11287 = m.ExcPending
	if v11287 != 0 {
		goto L32
	} else {
		goto L2404
	}
L2402:
	;
	goto L2403
L2403:
	;
	v11353 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11357 = F_LWLockAcquire(m, v11353+int32(1664), int32(0))
	mBase = m.M
	v11358 = m.ExcPending
	if v11358 != 0 {
		goto L32
	} else {
		goto L2413
	}
L2404:
	;
	v11290 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_308), v11269, int32(1), v11167)
	mBase = m.M
	v11291 = m.ExcPending
	if v11291 != 0 {
		goto L32
	} else {
		goto L2405
	}
L2405:
	;
	v11292 = int32(2)
	v11295 = v11167 << (uint(v11292) % 32) & int32(12)
	v11300 = v11295 + v11275*int32(20) + int32(4)
	v11302 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11303 = *(*int32)(unsafe.Add(mBase, uint32(v11302)+4))
	v11307 = *(*int32)(unsafe.Add(mBase, uint32(v11303+v11290<<(uint(v11292)%32))))
	v11308 = v11300 + v11307
	if v11308&int32(3)|base.B2i32(base.Ui32(v11300) < base.Ui32(int32(_a_F_StartupXLOG_309))) == int32(0) {
		goto L2407
	} else {
		goto L2408
	}
L2406:
	;
	v11340 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11341 = *(*int32)(unsafe.Add(mBase, uint32(v11340)+12))
	v11343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11341+v11290))) = uint8(v11343)
	F_LWLockRelease(m, v11284)
	mBase = m.M
	v11346 = m.ExcPending
	if v11346 != 0 {
		goto L32
	} else {
		goto L2412
	}
L2407:
	;
	v11327 = (v11268*int32(_a_F_StartupXLOG_310)-v11295+v11273*int32(-20)+int32(_a_F_StartupXLOG_311))&int32(-4) + int32(4)
	if v11327 == int32(0) {
		goto L2406
	} else {
		goto L2410
	}
L2408:
	;
	goto L2409
L2409:
	;
	v11333 = int32(_a_F_StartupXLOG_58) - v11300
	if v11333 == int32(0) {
		goto L2406
	} else {
		goto L2411
	}
L2410:
	;
	base.MemoryFill(m, v11308, int32(0), v11327)
	goto L2406
L2411:
	;
	base.MemoryFill(m, v11308, int32(0), v11333)
	goto L2406
L2412:
	;
	goto L2403
L2413:
	;
	v11360 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v11361 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11360)+8)) = uint8(v11361)
	v11364 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11364+int32(1664))
	mBase = m.M
	v11368 = m.ExcPending
	if v11368 != 0 {
		goto L32
	} else {
		goto L2414
	}
L2414:
	;
	F_SetMultiXactIdLimit(m, v11166, v11165, int32(1))
	mBase = m.M
	v11371 = m.ExcPending
	if v11371 != 0 {
		goto L32
	} else {
		goto L2415
	}
L2415:
	;
	v11372 = int32(0)
	v11373 = m.G0
	v11375 = v11373 - int32(16)
	m.G0 = v11375
	v11378 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11382 = F_LWLockAcquire(m, v11378+int32(2304), v11372)
	mBase = m.M
	v11383 = m.ExcPending
	if v11383 != 0 {
		goto L32
	} else {
		goto L2416
	}
L2416:
	;
	v11385 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[138]))
	v11386 = *(*int32)(unsafe.Add(mBase, uint32(v11385)+4))
	if int32(0) < v11386 {
		goto L2417
	} else {
		goto L2418
	}
L2417:
	;
	v11395 = v11385
	v11401 = v11372
	goto L2420
L2418:
	;
	goto L2419
L2419:
	;
	v11700 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11700+int32(2304))
	mBase = m.M
	v11704 = m.ExcPending
	if v11704 != 0 {
		goto L32
	} else {
		goto L2461
	}
L2420:
	;
	v11426 = *(*int32)(unsafe.Add(mBase, uint32(v11395+v11401<<(uint(int32(2))%32))+8))
	v11427 = *(*int32)(unsafe.Add(mBase, uint32(v11426)+32))
	v11428 = *(*int64)(unsafe.Add(mBase, uint32(v11426)+16))
	v11429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11426)+45)))
	v11432 = F_ProcessTwoPhaseBuffer(m, v11427, v11428, v11429, int32(1), int32(0))
	mBase = m.M
	v11433 = m.ExcPending
	if v11433 != 0 {
		goto L32
	} else {
		goto L2422
	}
L2421:
	;
	goto L2419
L2422:
	;
	if v11432 != 0 {
		goto L2423
	} else {
		goto L2424
	}
L2423:
	;
	v11436 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11437 = m.ExcPending
	if v11437 != 0 {
		goto L32
	} else {
		goto L2426
	}
L2424:
	;
	goto L2425
L2425:
	;
	v11660 = v11401 + int32(1)
	v11662 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[138]))
	v11663 = *(*int32)(unsafe.Add(mBase, uint32(v11662)+4))
	if v11660 < v11663 {
		v11395 = v11662
		v11401 = v11660
		goto L2420
	} else {
		goto L2460
	}
L2426:
	;
	if v11436 != 0 {
		goto L2427
	} else {
		goto L2428
	}
L2427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11375))) = v11427
	F_errmsg(m, int32(_a_F_StartupXLOG_312), v11375)
	mBase = m.M
	v11441 = m.ExcPending
	if v11441 != 0 {
		goto L32
	} else {
		goto L2430
	}
L2428:
	;
	goto L2429
L2429:
	;
	v11447 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+48))
	v11448 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+44))
	v11449 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+40))
	v11450 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+36))
	v11451 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+32))
	v11452 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+28))
	v11453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11432)+54)))
	v11455 = v11432 + int32(72)
	v11456 = *(*int64)(unsafe.Add(mBase, uint32(v11432)+16))
	v11457 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+24))
	v11458 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+12))
	F_MarkAsPreparingGuts(m, v11426, v11427, v11455, v11456, v11457, v11458)
	mBase = m.M
	v11460 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11426)+46)) = uint8(v11460)
	v11462 = int32(7)
	v11466 = v11455 + (v11453+v11462)&int32(_a_F_StartupXLOG_313)
	v11471 = int32(-8)
	v11474 = int32(12)
	v11488 = int32(4)
	v11498 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[139]))
	v11499 = *(*int32)(unsafe.Add(mBase, uint32(v11498)))
	v11500 = *(*int32)(unsafe.Add(mBase, uint32(v11426)+4))
	v11503 = v11499 + v11500*int32(640)
	v11504 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+28))
	if int32(65) <= v11504 {
		goto L2434
	} else {
		goto L2435
	}
L2430:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_314), int32(2106), int32(_a_F_StartupXLOG_315))
	mBase = m.M
	v11446 = m.ExcPending
	if v11446 != 0 {
		goto L32
	} else {
		goto L2431
	}
L2431:
	;
	goto L2429
L2432:
	;
	v11525 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11426)+44)) = uint8(v11525)
	v11527 = *(*int32)(unsafe.Add(mBase, uint32(v11523)))
	F_ProcArrayAdd(m, v11527+v11524*int32(640))
	mBase = m.M
	v11532 = m.ExcPending
	if v11532 != 0 {
		goto L32
	} else {
		goto L2441
	}
L2433:
	;
	v11514 = v11512 << (uint(int32(2)) % 32)
	if v11514 != 0 {
		goto L2438
	} else {
		goto L2439
	}
L2434:
	;
	v11507 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11503)+277)) = uint8(v11507)
	v11512 = int32(64)
	goto L2433
L2435:
	;
	goto L2436
L2436:
	;
	if v11504 <= int32(0) {
		v11523 = v11498
		v11524 = v11500
		goto L2432
	} else {
		goto L2437
	}
L2437:
	;
	v11512 = v11504
	goto L2433
L2438:
	;
	base.MemoryCopy(m, v11503+int32(280), v11466, v11514)
	goto L2440
L2439:
	;
	goto L2440
L2440:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11503)+276)) = uint8(v11512)
	v11519 = *(*int32)(unsafe.Add(mBase, uint32(v11426)+4))
	v11521 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[139]))
	v11523 = v11521
	v11524 = v11519
	goto L2432
L2441:
	;
	v11534 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11534+int32(2304))
	mBase = m.M
	v11538 = m.ExcPending
	if v11538 != 0 {
		goto L32
	} else {
		goto L2442
	}
L2442:
	;
	v11545 = v11466 + (v11452<<(uint(int32(2))%32)+v11462)&v11471 + (v11451*v11474+v11462)&v11471 + (v11450*v11474+v11462)&v11471 + v11449<<(uint(v11488)%32) + v11448<<(uint(v11488)%32) + v11447<<(uint(v11488)%32)
	goto L2443
L2443:
	;
	v11573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11545)+4)))
	if v11573 != 0 {
		goto L2445
	} else {
		goto L2446
	}
L2444:
	;
	v11590 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if base.Ui32(int32(2)) <= base.Ui32(v11590) {
		goto L2452
	} else {
		goto L2453
	}
L2445:
	;
	v11575 = v11545 + int32(8)
	v11578 = *(*int32)(unsafe.Add(mBase, uint32(v11573<<(uint(int32(2))%32))+uint32(_c_F_StartupXLOG[140])))
	if v11578 != 0 {
		goto L2448
	} else {
		goto L2449
	}
L2446:
	;
	goto L2447
L2447:
	;
	goto L2444
L2448:
	;
	v11579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11545)+6)))
	v11580 = *(*int32)(unsafe.Add(mBase, uint32(v11545)))
	m.T0[v11578].(func(*base.Module, int32, int32, int32, int32))(m, v11427, v11579, v11575, v11580)
	mBase = m.M
	v11582 = m.ExcPending
	if v11582 != 0 {
		goto L32
	} else {
		goto L2451
	}
L2449:
	;
	goto L2450
L2450:
	;
	v11583 = *(*int32)(unsafe.Add(mBase, uint32(v11545)))
	v11545 = v11575 + (v11583+int32(7))&int32(-8)
	goto L2443
L2451:
	;
	goto L2450
L2452:
	;
	v11593 = *(*int32)(unsafe.Add(mBase, uint32(v11432)+28))
	F_StandbyReleaseLockTree(m, v11427, v11593, v11466)
	mBase = m.M
	v11595 = m.ExcPending
	if v11595 != 0 {
		goto L32
	} else {
		goto L2455
	}
L2453:
	;
	goto L2454
L2454:
	;
	v11597 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11601 = F_LWLockAcquire(m, v11597+int32(2304), int32(0))
	mBase = m.M
	v11602 = m.ExcPending
	if v11602 != 0 {
		goto L32
	} else {
		goto L2456
	}
L2455:
	;
	goto L2454
L2456:
	;
	v11604 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[141]))
	*(*int32)(unsafe.Add(mBase, uint32(v11604)+40)) = int32(-1)
	v11608 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11608+int32(2304))
	mBase = m.M
	v11612 = m.ExcPending
	if v11612 != 0 {
		goto L32
	} else {
		goto L2457
	}
L2457:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[141])) = int32(0)
	F_pfree(m, v11432)
	mBase = m.M
	v11617 = m.ExcPending
	if v11617 != 0 {
		goto L32
	} else {
		goto L2458
	}
L2458:
	;
	v11619 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11623 = F_LWLockAcquire(m, v11619+int32(2304), int32(0))
	mBase = m.M
	v11624 = m.ExcPending
	if v11624 != 0 {
		goto L32
	} else {
		goto L2459
	}
L2459:
	;
	goto L2425
L2460:
	;
	goto L2421
L2461:
	;
	m.G0 = v11375 + int32(16)
	v11708 = m.G0
	v11710 = v11708 - int32(1024)
	m.G0 = v11710
	v11713 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v11718 = *(*int32)(unsafe.Add(mBase, uint32(v11713)))
	v11719 = *(*int32)(unsafe.Add(mBase, uint32(v11718)+124))
	if v11719 != 0 {
		goto L2463
	} else {
		goto L2464
	}
L2462:
	;
	v11741 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127]))
	if int32(0) <= v11741 {
		goto L2466
	} else {
		goto L2467
	}
L2463:
	;
	v11720 = *(*int64)(unsafe.Add(mBase, uint32(v11719)+16))
	v11721 = *(*int32)(unsafe.Add(mBase, uint32(v11718)+120))
	v11722 = *(*int64)(unsafe.Add(mBase, uint32(v11721)+16))
	v11725 = base.I32_wrap_i64(v11720 - v11722)
	goto L2465
L2464:
	;
	v11725 = int32(0)
	goto L2465
L2465:
	;
	v11726 = *(*int32)(unsafe.Add(mBase, uint32(v11713)+112))
	v11727 = *(*int32)(unsafe.Add(mBase, uint32(v11726)+16))
	v11729 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[30]))
	v11730 = *(*int32)(unsafe.Add(mBase, uint32(v11726)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11729)+64)) = v11730
	*(*int32)(unsafe.Add(mBase, uint32(v11729)+56)) = v11725
	*(*int32)(unsafe.Add(mBase, uint32(v11729)+60)) = v11730 + v11727
	v11735 = *(*int32)(unsafe.Add(mBase, uint32(v11713)))
	v11736 = *(*int64)(unsafe.Add(mBase, uint32(v11735)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v11713)+16)) = v11736 - int64(-8192)
	goto L2462
L2466:
	;
	v11744 = F_close(m, v11741)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[127])) = int32(-1)
	goto L2468
L2467:
	;
	goto L2468
L2468:
	;
	v11749 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	F_XLogReaderFree(m, v11749)
	mBase = m.M
	v11751 = m.ExcPending
	if v11751 != 0 {
		goto L32
	} else {
		goto L2469
	}
L2469:
	;
	v11753 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v11754 = *(*int32)(unsafe.Add(mBase, uint32(v11753)+112))
	F_pfree(m, v11754)
	mBase = m.M
	v11756 = m.ExcPending
	if v11756 != 0 {
		goto L32
	} else {
		goto L2470
	}
L2470:
	;
	v11757 = *(*int32)(unsafe.Add(mBase, uint32(v11753)+24))
	F_hash_destroy(m, v11757)
	mBase = m.M
	v11759 = m.ExcPending
	if v11759 != 0 {
		goto L32
	} else {
		goto L2471
	}
L2471:
	;
	F_pfree(m, v11753)
	mBase = m.M
	v11761 = m.ExcPending
	if v11761 != 0 {
		goto L32
	} else {
		goto L2472
	}
L2472:
	;
	v11763 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11763 != int32(1) {
		goto L2473
	} else {
		goto L2474
	}
L2473:
	;
	m.G0 = v11710 + int32(1024)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[142])) = int32(1)
	if v9729 != int64(0) {
		goto L2479
	} else {
		goto L2480
	}
L2474:
	;
	v11769 = F_pg_snprintf(m, v11710, int32(1024), int32(_a_F_StartupXLOG_316), int32(0))
	mBase = m.M
	v11770 = m.ExcPending
	if v11770 != 0 {
		goto L32
	} else {
		goto L2475
	}
L2475:
	;
	v11771 = F_unlink(m, v11710)
	mBase = m.M
	v11775 = F_pg_snprintf(m, v11710, int32(1024), int32(_a_F_StartupXLOG_317), int32(0))
	mBase = m.M
	v11776 = m.ExcPending
	if v11776 != 0 {
		goto L32
	} else {
		goto L2476
	}
L2476:
	;
	v11777 = F_unlink(m, v11710)
	mBase = m.M
	v11779 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11779 != int32(1) {
		goto L2473
	} else {
		goto L2477
	}
L2477:
	;
	v11783 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v11783+int32(4))+12)) = int32(0)
	goto L2478
L2478:
	;
	goto L2473
L2479:
	;
	v11797 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[143])))
	if v11797 != int32(1) {
		goto L5
	} else {
		goto L2482
	}
L2480:
	;
	goto L2481
L2481:
	;
	v11960 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[73])))
	*(*uint8)(unsafe.Add(mBase, uint32(v10732)+160)) = uint8(v11960)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v11963 = m.ExcPending
	if v11963 != 0 {
		goto L32
	} else {
		goto L2512
	}
L2482:
	;
	v11802 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v11803 = *(*int32)(unsafe.Add(mBase, uint32(v11802)+316))
	v11805 = base.B2i32(v11803 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[143])) = uint8(v11805)
	if v11805 == int32(0) {
		goto L5
	} else {
		goto L2483
	}
L2483:
	;
	if v9732&int64(8191) != int64(0) {
		goto L4
	} else {
		goto L2484
	}
L2484:
	;
	v11816 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v9732&base.I64_extend_i32_s(v11816-int32(1)) == int64(0) {
		goto L2485
	} else {
		goto L2486
	}
L2485:
	;
	v11823 = int64(40)
	goto L2487
L2486:
	;
	v11823 = int64(24)
	goto L2487
L2487:
	;
	v11826 = base.AtomicRmwXchg32(m, v11802, int32(0), int32(1))
	if v11826 != 0 {
		goto L2488
	} else {
		goto L2489
	}
L2488:
	;
	F_s_lock(m, v11802, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_318), int32(_a_F_StartupXLOG_319))
	mBase = m.M
	v11831 = m.ExcPending
	if v11831 != 0 {
		goto L32
	} else {
		goto L2491
	}
L2489:
	;
	goto L2490
L2490:
	;
	v11832 = v11823 | v9732
	v11833 = *(*int64)(unsafe.Add(mBase, uint32(v11802)+8))
	v11834 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11802))), uint32(v11834))
	v11838 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[131])))
	v11839 = base.I64_div_u_s(v11833, v11838)
	v11841 = v11833 - v11839*v11838
	if base.Ui64(v11841) <= base.Ui64(int64(8151)) {
		goto L2493
	} else {
		goto L2494
	}
L2491:
	;
	goto L2490
L2492:
	;
	v11861 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v11865 = v11839*v11861 + v11859&int64(4294967295)
	if v11865 != v11832 {
		goto L3
	} else {
		goto L2496
	}
L2493:
	;
	v11859 = v11841 + int64(40)
	goto L2492
L2494:
	;
	goto L2495
L2495:
	;
	v11847 = v11841 - int64(8152)
	v11848 = int64(8168)
	v11849 = base.I64_div_u_s(v11847, v11848)
	v11859 = v11847 - v11849*v11848 + v11849<<(uint(int64(13))%64) + int64(8216)
	goto L2492
L2496:
	;
	v11867 = int32(_a_F_StartupXLOG_142)
	v11869 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v11869 + int32(1)
	v11874 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144]))
	if v11874 == int32(-1) {
		goto L2497
	} else {
		goto L2498
	}
L2497:
	;
	v11879 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[87]))
	v11881 = base.I32_rem_s(v11879, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144])) = v11881
	v11883 = v11881
	goto L2499
L2498:
	;
	v11883 = v11874
	goto L2499
L2499:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[145])) = v11883
	v11887 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[146]))
	v11892 = F_LWLockAcquire(m, v11887+v11883<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v11893 = m.ExcPending
	if v11893 != 0 {
		goto L32
	} else {
		goto L2500
	}
L2500:
	;
	if v11892 == int32(0) {
		goto L2501
	} else {
		goto L2502
	}
L2501:
	;
	v11896 = int32(_a_F_StartupXLOG_320)
	v11898 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144]))
	v11902 = base.I32_rem_s(v11898+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144])) = v11902
	goto L2503
L2502:
	;
	goto L2503
L2503:
	;
	v11904 = F_GetXLogBuffer(m, v9732, v10688)
	mBase = m.M
	v11905 = m.ExcPending
	if v11905 != 0 {
		goto L32
	} else {
		goto L2504
	}
L2504:
	;
	v11906 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11904)+2)))
	v11908 = v11906 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v11904)+2)) = uint16(v11908)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v11911 = m.ExcPending
	if v11911 != 0 {
		goto L32
	} else {
		goto L2505
	}
L2505:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v11913 = m.ExcPending
	if v11913 != 0 {
		goto L32
	} else {
		goto L2506
	}
L2506:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v9729
	v11918 = m.G0
	v11919 = int32(16)
	v11920 = v11918 - v11919
	m.G0 = v11920
	F_gettimeofday(m, v11920)
	mBase = m.M
	v11923 = *(*int64)(unsafe.Add(mBase, uint32(v11920)))
	v11924 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11920)+8)))
	m.G0 = v11920 + v11919
	goto L2507
L2507:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[147]))) = v11924 + v11923*int64(1000000) - int64(946684800000000)
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(16))
	mBase = m.M
	v11938 = m.ExcPending
	if v11938 != 0 {
		goto L32
	} else {
		goto L2508
	}
L2508:
	;
	v11941 = F_XLogInsert(m, int32(0), int32(208))
	mBase = m.M
	v11942 = m.ExcPending
	if v11942 != 0 {
		goto L32
	} else {
		goto L2509
	}
L2509:
	;
	v11944 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	if v11944 != v11832 {
		goto L2
	} else {
		goto L2510
	}
L2510:
	;
	F_XLogFlush(m, v11941)
	mBase = m.M
	v11947 = m.ExcPending
	if v11947 != 0 {
		goto L32
	} else {
		goto L2511
	}
L2511:
	;
	v11948 = int32(_a_F_StartupXLOG_142)
	v11950 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v11950 - int32(1)
	goto L2481
L2512:
	;
	v11964 = int32(0)
	if v6012 == v11964 {
		v12077 = v11964
		goto L2513
	} else {
		goto L2514
	}
L2513:
	;
	v12081 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	v12083 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12084 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+172))
	if v12081 != v12084 {
		goto L2534
	} else {
		goto L2535
	}
L2514:
	;
	v11968 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11968 != int32(1) {
		goto L2515
	} else {
		goto L2516
	}
L2515:
	;
	F_RequestCheckpoint(m, int32(38))
	mBase = m.M
	v12076 = m.ExcPending
	if v12076 != 0 {
		goto L32
	} else {
		goto L2532
	}
L2516:
	;
	v11972 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v11972&int32(1) == int32(0) {
		goto L2515
	} else {
		goto L2517
	}
L2517:
	;
	v11977 = F_PromoteIsTriggered(m)
	mBase = m.M
	v11978 = m.ExcPending
	if v11978 != 0 {
		goto L32
	} else {
		goto L2518
	}
L2518:
	;
	if v11977 == int32(0) {
		goto L2515
	} else {
		goto L2519
	}
L2519:
	;
	v11982 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[143])))
	if v11982 != int32(1) {
		goto L1
	} else {
		goto L2520
	}
L2520:
	;
	v11987 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v11988 = *(*int32)(unsafe.Add(mBase, uint32(v11987)+316))
	v11990 = base.B2i32(v11988 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[143])) = uint8(v11990)
	if v11990 == int32(0) {
		goto L1
	} else {
		goto L2521
	}
L2521:
	;
	v11997 = m.G0
	v11998 = int32(16)
	v11999 = v11997 - v11998
	m.G0 = v11999
	F_gettimeofday(m, v11999)
	mBase = m.M
	v12002 = *(*int64)(unsafe.Add(mBase, uint32(v11999)))
	v12003 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11999)+8)))
	m.G0 = v11999 + v11998
	goto L2522
L2522:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v12003 + v12002*int64(1000000) - int64(946684800000000)
	v12014 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v12014
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v12017 = m.ExcPending
	if v12017 != 0 {
		goto L32
	} else {
		goto L2523
	}
L2523:
	;
	v12019 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v12020 = *(*int32)(unsafe.Add(mBase, uint32(v12019)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[147]))) = v12020
	v12022 = *(*int32)(unsafe.Add(mBase, uint32(v12019)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v12022
	F_WALInsertLockRelease(m)
	mBase = m.M
	v12025 = m.ExcPending
	if v12025 != 0 {
		goto L32
	} else {
		goto L2524
	}
L2524:
	;
	v12026 = int32(1)
	v12027 = int32(_a_F_StartupXLOG_142)
	v12029 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v12029 + v12026
	F_XLogBeginInsert(m)
	mBase = m.M
	v12034 = m.ExcPending
	if v12034 != 0 {
		goto L32
	} else {
		goto L2525
	}
L2525:
	;
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(24))
	mBase = m.M
	v12039 = m.ExcPending
	if v12039 != 0 {
		goto L32
	} else {
		goto L2526
	}
L2526:
	;
	v12042 = F_XLogInsert(m, int32(0), int32(144))
	mBase = m.M
	v12043 = m.ExcPending
	if v12043 != 0 {
		goto L32
	} else {
		goto L2527
	}
L2527:
	;
	F_XLogFlush(m, v12042)
	mBase = m.M
	v12045 = m.ExcPending
	if v12045 != 0 {
		goto L32
	} else {
		goto L2528
	}
L2528:
	;
	v12047 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12051 = F_LWLockAcquire(m, v12047+int32(1152), int32(0))
	mBase = m.M
	v12052 = m.ExcPending
	if v12052 != 0 {
		goto L32
	} else {
		goto L2529
	}
L2529:
	;
	v12054 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v12054)+136)) = v12042
	v12056 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[147])))
	*(*int32)(unsafe.Add(mBase, uint32(v12054)+144)) = v12056
	v12059 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	F_update_controlfile(m, v12059, v12054)
	mBase = m.M
	v12061 = m.ExcPending
	if v12061 != 0 {
		goto L32
	} else {
		goto L2530
	}
L2530:
	;
	v12063 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12063+int32(1152))
	mBase = m.M
	v12067 = m.ExcPending
	if v12067 != 0 {
		goto L32
	} else {
		goto L2531
	}
L2531:
	;
	v12068 = int32(_a_F_StartupXLOG_142)
	v12070 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v12070 - int32(1)
	v12077 = v12026
	goto L2513
L2532:
	;
	v12077 = v11964
	goto L2513
L2533:
	;
	v12200 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v12200 != int32(1) {
		goto L2553
	} else {
		goto L2554
	}
L2534:
	;
	v12115 = int32(0)
	if base.B2i32(v12081 == v12084)&base.B2i32(v12081 <= v12115) == v12115 {
		goto L2543
	} else {
		goto L2544
	}
L2535:
	;
	v12087 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[149])))
	v12088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12083)+176)))
	if v12087 != v12088 {
		goto L2534
	} else {
		goto L2536
	}
L2536:
	;
	v12091 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	v12092 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+180))
	if v12091 != v12092 {
		goto L2534
	} else {
		goto L2537
	}
L2537:
	;
	v12095 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	v12096 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+184))
	if v12095 != v12096 {
		goto L2534
	} else {
		goto L2538
	}
L2538:
	;
	v12099 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[122]))
	v12100 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+188))
	if v12099 != v12100 {
		goto L2534
	} else {
		goto L2539
	}
L2539:
	;
	v12103 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	v12104 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+192))
	if v12103 != v12104 {
		goto L2534
	} else {
		goto L2540
	}
L2540:
	;
	v12107 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	v12108 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+196))
	if v12107 != v12108 {
		goto L2534
	} else {
		goto L2541
	}
L2541:
	;
	v12111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[154])))
	v12112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12083)+200)))
	if v12111 == v12112 {
		goto L2533
	} else {
		goto L2542
	}
L2542:
	;
	goto L2534
L2543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[155]))) = v12081
	v12122 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v12122
	v12125 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[93]))) = v12125
	v12128 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[122]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[147]))) = v12128
	v12131 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v12131
	v12134 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v12134
	v12137 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[149])))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[95]))) = uint8(v12137)
	v12140 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[154])))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[156]))) = uint8(v12140)
	F_XLogBeginInsert(m)
	mBase = m.M
	v12143 = m.ExcPending
	if v12143 != 0 {
		goto L32
	} else {
		goto L2546
	}
L2544:
	;
	goto L2545
L2545:
	;
	v12156 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12160 = F_LWLockAcquire(m, v12156+int32(1152), int32(0))
	mBase = m.M
	v12161 = m.ExcPending
	if v12161 != 0 {
		goto L32
	} else {
		goto L2550
	}
L2546:
	;
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(28))
	mBase = m.M
	v12148 = m.ExcPending
	if v12148 != 0 {
		goto L32
	} else {
		goto L2547
	}
L2547:
	;
	v12151 = F_XLogInsert(m, int32(0), int32(96))
	mBase = m.M
	v12152 = m.ExcPending
	if v12152 != 0 {
		goto L32
	} else {
		goto L2548
	}
L2548:
	;
	F_XLogFlush(m, v12151)
	mBase = m.M
	v12154 = m.ExcPending
	if v12154 != 0 {
		goto L32
	} else {
		goto L2549
	}
L2549:
	;
	goto L2545
L2550:
	;
	v12163 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12165 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+180)) = v12165
	v12168 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+184)) = v12168
	v12171 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[122]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+188)) = v12171
	v12174 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[152]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+192)) = v12174
	v12177 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[153]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+196)) = v12177
	v12180 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, uint32(v12163)+172)) = v12180
	v12183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[149])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12163)+176)) = uint8(v12183)
	v12186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[154])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12163)+200)) = uint8(v12186)
	v12189 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	F_update_controlfile(m, v12189, v12163)
	mBase = m.M
	v12191 = m.ExcPending
	if v12191 != 0 {
		goto L32
	} else {
		goto L2551
	}
L2551:
	;
	v12193 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12193+int32(1152))
	mBase = m.M
	v12197 = m.ExcPending
	if v12197 != 0 {
		goto L32
	} else {
		goto L2552
	}
L2552:
	;
	goto L2533
L2553:
	;
	v12363 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[154])))
	if v12363 == int32(0) {
		goto L2584
	} else {
		goto L2585
	}
L2554:
	;
	v12204 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[157]))
	if v12204 == int32(0) {
		goto L2555
	} else {
		goto L2556
	}
L2555:
	;
	F_RemoveNonParentXlogFiles(m, v10741, v10688)
	mBase = m.M
	v12216 = m.ExcPending
	if v12216 != 0 {
		goto L32
	} else {
		goto L2559
	}
L2556:
	;
	v12207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12204))))
	if v12207 == int32(0) {
		goto L2555
	} else {
		goto L2557
	}
L2557:
	;
	F_ExecuteRecoveryCommand(m, v12204, int32(_a_F_StartupXLOG_321), int32(1), int32(134217774))
	mBase = m.M
	v12214 = m.ExcPending
	if v12214 != 0 {
		goto L32
	} else {
		goto L2558
	}
L2558:
	;
	goto L2555
L2559:
	;
	v12218 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v10741&base.I64_extend_i32_s(v12218-int32(1)) == int64(0) {
		goto L2553
	} else {
		goto L2560
	}
L2560:
	;
	v12226 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[130]))
	if v12226 <= int32(0) {
		goto L2553
	} else {
		goto L2561
	}
L2561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3616)) = v9745
	v12232 = base.I64_extend_i32_s(v12218)
	v12233 = base.I64_div_u_s(v10741-int64(1), v12232)
	v12235 = base.I64_div_u_s(int64(4294967296), v12232)
	v12236 = base.I64_div_u_s(v12233, v12235)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3620)) = uint32(v12236)
	v12239 = v12233 - v12236*v12235
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3624)) = uint32(v12239)
	v12242 = v39 + int32(_a_F_StartupXLOG_3)
	v12247 = F_pg_snprintf(m, v12242, int32(64), int32(_a_F_StartupXLOG_278), v39+int32(3616))
	mBase = m.M
	v12248 = m.ExcPending
	if v12248 != 0 {
		goto L32
	} else {
		goto L2562
	}
L2562:
	;
	v12249 = m.G0
	v12251 = v12249 - int32(1168)
	m.G0 = v12251
	*(*int32)(unsafe.Add(mBase, uint32(v12251)+32)) = v12242
	*(*int32)(unsafe.Add(mBase, uint32(v12251)+36)) = int32(_a_F_StartupXLOG_322)
	v12257 = v12251 + int32(144)
	v12262 = F_pg_snprintf(m, v12257, int32(1024), int32(_a_F_StartupXLOG_323), v12251+int32(32))
	mBase = m.M
	v12263 = m.ExcPending
	if v12263 != 0 {
		goto L32
	} else {
		goto L2563
	}
L2563:
	;
	v12264 = int32(1)
	v12266 = v12251 + int32(48)
	v12269 = F___fstatat(m, int32(-100), v12257, v12266, int32(0))
	mBase = m.M
	goto L2565
L2564:
	;
	m.G0 = v12251 + int32(1168)
	if v12298 != 0 {
		goto L2553
	} else {
		goto L2572
	}
L2565:
	;
	if v12269 == int32(0) {
		v12298 = v12264
		goto L2564
	} else {
		goto L2566
	}
L2566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12251)+20)) = int32(_a_F_StartupXLOG_324)
	*(*int32)(unsafe.Add(mBase, uint32(v12251)+16)) = v12242
	v12279 = F_pg_snprintf(m, v12257, int32(1024), int32(_a_F_StartupXLOG_323), v12251+int32(16))
	mBase = m.M
	v12280 = m.ExcPending
	if v12280 != 0 {
		goto L32
	} else {
		goto L2567
	}
L2567:
	;
	v12283 = F___fstatat(m, int32(-100), v12257, v12266, int32(0))
	mBase = m.M
	goto L2568
L2568:
	;
	if v12283 == int32(0) {
		v12298 = v12264
		goto L2564
	} else {
		goto L2569
	}
L2569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12251)+4)) = int32(_a_F_StartupXLOG_322)
	*(*int32)(unsafe.Add(mBase, uint32(v12251))) = v12242
	v12291 = F_pg_snprintf(m, v12257, int32(1024), int32(_a_F_StartupXLOG_323), v12251)
	mBase = m.M
	v12292 = m.ExcPending
	if v12292 != 0 {
		goto L32
	} else {
		goto L2570
	}
L2570:
	;
	v12295 = F___fstatat(m, int32(-100), v12257, v12266, int32(0))
	mBase = m.M
	goto L2571
L2571:
	;
	v12298 = base.B2i32(v12295 == int32(0))
	goto L2564
L2572:
	;
	v12303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[158])))
	if v12303 == int32(1) {
		goto L2573
	} else {
		goto L2574
	}
L2573:
	;
	F_WaitForWalSummarization(m, v10741)
	mBase = m.M
	v12307 = m.ExcPending
	if v12307 != 0 {
		goto L32
	} else {
		goto L2576
	}
L2574:
	;
	goto L2575
L2575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3600)) = v9745
	v12311 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v12312 = base.I64_div_u_s(int64(4294967296), v12311)
	v12313 = base.I64_div_u_s(v12233, v12312)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3604)) = uint32(v12313)
	v12316 = v12233 - v12312*v12313
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3608)) = uint32(v12316)
	v12319 = v39 + int32(_a_F_StartupXLOG_1)
	v12324 = F_pg_snprintf(m, v12319, int32(1024), int32(_a_F_StartupXLOG_279), v39+int32(3600))
	mBase = m.M
	v12325 = m.ExcPending
	if v12325 != 0 {
		goto L32
	} else {
		goto L2577
	}
L2576:
	;
	goto L2575
L2577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3584)) = v39 + int32(_a_F_StartupXLOG_3)
	v12330 = v39 + int32(_a_F_StartupXLOG_283)
	v12335 = F_pg_snprintf(m, v12330, int32(64), int32(_a_F_StartupXLOG_325), v39+int32(3584))
	mBase = m.M
	v12336 = m.ExcPending
	if v12336 != 0 {
		goto L32
	} else {
		goto L2578
	}
L2578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3568)) = v12319
	v12339 = v39 + int32(_a_F_StartupXLOG_4)
	v12344 = F_pg_snprintf(m, v12339, int32(1024), int32(_a_F_StartupXLOG_325), v39+int32(3568))
	mBase = m.M
	v12345 = m.ExcPending
	if v12345 != 0 {
		goto L32
	} else {
		goto L2579
	}
L2579:
	;
	F_XLogArchiveCleanup(m, v12330)
	mBase = m.M
	v12347 = m.ExcPending
	if v12347 != 0 {
		goto L32
	} else {
		goto L2580
	}
L2580:
	;
	v12349 = F_durable_rename(m, v12319, v12339, int32(21))
	mBase = m.M
	v12350 = m.ExcPending
	if v12350 != 0 {
		goto L32
	} else {
		goto L2581
	}
L2581:
	;
	F_XLogArchiveNotify(m, v12330)
	mBase = m.M
	v12352 = m.ExcPending
	if v12352 != 0 {
		goto L32
	} else {
		goto L2582
	}
L2582:
	;
	goto L2553
L2583:
	;
	v12402 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12406 = F_LWLockAcquire(m, v12402+int32(1152), int32(0))
	mBase = m.M
	v12407 = m.ExcPending
	if v12407 != 0 {
		goto L32
	} else {
		goto L2591
	}
L2584:
	;
	v12367 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12371 = F_LWLockAcquire(m, v12367+int32(_a_F_StartupXLOG_326), int32(0))
	mBase = m.M
	v12372 = m.ExcPending
	if v12372 != 0 {
		goto L32
	} else {
		goto L2587
	}
L2585:
	;
	goto L2586
L2586:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v12399 = m.ExcPending
	if v12399 != 0 {
		goto L32
	} else {
		goto L2590
	}
L2587:
	;
	v12374 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[159]))
	v12375 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12374)+16)) = uint16(v12375)
	*(*int64)(unsafe.Add(mBase, uint32(v12374)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v12374))) = v12375
	*(*uint8)(unsafe.Add(mBase, uint32(v12374)+24)) = uint8(v12375)
	v12384 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	*(*int64)(unsafe.Add(mBase, uint32(v12384)+40)) = int64(0)
	v12390 = F_SlruScanDirectory(m, int32(_a_F_StartupXLOG_327), int32(290), v12375)
	mBase = m.M
	v12391 = m.ExcPending
	if v12391 != 0 {
		goto L32
	} else {
		goto L2588
	}
L2588:
	;
	v12393 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12393+int32(_a_F_StartupXLOG_326))
	mBase = m.M
	v12397 = m.ExcPending
	if v12397 != 0 {
		goto L32
	} else {
		goto L2589
	}
L2589:
	;
	goto L2583
L2590:
	;
	goto L2583
L2591:
	;
	v12409 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v12409)+16)) = int32(6)
	v12413 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v12416 = base.AtomicRmwXchg32(m, v12413, int32(440), int32(1))
	if v12416 != 0 {
		goto L2592
	} else {
		goto L2593
	}
L2592:
	;
	v12418 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v12418+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_328), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12425 = m.ExcPending
	if v12425 != 0 {
		goto L32
	} else {
		goto L2595
	}
L2593:
	;
	goto L2594
L2594:
	;
	v12427 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v12427)+316)) = int32(2)
	v12430 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12427)+440)), uint32(v12430))
	v12434 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	v12436 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v12434, v12436)
	mBase = m.M
	v12438 = m.ExcPending
	if v12438 != 0 {
		goto L32
	} else {
		goto L2596
	}
L2595:
	;
	goto L2594
L2596:
	;
	v12440 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12440+int32(1152))
	mBase = m.M
	v12444 = m.ExcPending
	if v12444 != 0 {
		goto L32
	} else {
		goto L2597
	}
L2597:
	;
	v12446 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v12446 != 0 {
		goto L2598
	} else {
		goto L2599
	}
L2598:
	;
	F_ShutdownRecoveryTransactionEnvironment(m)
	mBase = m.M
	v12448 = m.ExcPending
	if v12448 != 0 {
		goto L32
	} else {
		goto L2601
	}
L2599:
	;
	goto L2600
L2600:
	;
	v12449 = int32(1)
	F_WalSndWakeup(m, v12449, v12449)
	mBase = m.M
	v12452 = m.ExcPending
	if v12452 != 0 {
		goto L32
	} else {
		goto L2602
	}
L2601:
	;
	goto L2600
L2602:
	;
	if v12077 != 0 {
		goto L2603
	} else {
		goto L2604
	}
L2603:
	;
	F_RequestCheckpoint(m, int32(8))
	mBase = m.M
	v12455 = m.ExcPending
	if v12455 != 0 {
		goto L32
	} else {
		goto L2606
	}
L2604:
	;
	goto L2605
L2605:
	;
	m.G0 = v35
	return
L2606:
	;
	goto L2605
L2607:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12463 = m.ExcPending
	if v12463 != 0 {
		goto L32
	} else {
		goto L2608
	}
L2608:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_329), int32(0))
	mBase = m.M
	v12467 = m.ExcPending
	if v12467 != 0 {
		goto L32
	} else {
		goto L2609
	}
L2609:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_330), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12472 = m.ExcPending
	if v12472 != 0 {
		goto L32
	} else {
		goto L2610
	}
L2610:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2611:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12479 = m.ExcPending
	if v12479 != 0 {
		goto L32
	} else {
		goto L2612
	}
L2612:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_331), int32(0))
	mBase = m.M
	v12483 = m.ExcPending
	if v12483 != 0 {
		goto L32
	} else {
		goto L2613
	}
L2613:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_332), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12488 = m.ExcPending
	if v12488 != 0 {
		goto L32
	} else {
		goto L2614
	}
L2614:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2615:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12494 = m.ExcPending
	if v12494 != 0 {
		goto L32
	} else {
		goto L2616
	}
L2616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3984)) = int32(_a_F_StartupXLOG_5)
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v39+int32(3984))
	mBase = m.M
	v12501 = m.ExcPending
	if v12501 != 0 {
		goto L32
	} else {
		goto L2617
	}
L2617:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_333), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12506 = m.ExcPending
	if v12506 != 0 {
		goto L32
	} else {
		goto L2618
	}
L2618:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2619:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12512 = m.ExcPending
	if v12512 != 0 {
		goto L32
	} else {
		goto L2620
	}
L2620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3952)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_334), v39+int32(3952))
	mBase = m.M
	v12520 = m.ExcPending
	if v12520 != 0 {
		goto L32
	} else {
		goto L2621
	}
L2621:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_335), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12525 = m.ExcPending
	if v12525 != 0 {
		goto L32
	} else {
		goto L2622
	}
L2622:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3904)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_334), v39+int32(3904))
	mBase = m.M
	v12537 = m.ExcPending
	if v12537 != 0 {
		goto L32
	} else {
		goto L2624
	}
L2624:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_336), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12542 = m.ExcPending
	if v12542 != 0 {
		goto L32
	} else {
		goto L2625
	}
L2625:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2626:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v12549 = m.ExcPending
	if v12549 != 0 {
		goto L32
	} else {
		goto L2627
	}
L2627:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_337), int32(0))
	mBase = m.M
	v12553 = m.ExcPending
	if v12553 != 0 {
		goto L32
	} else {
		goto L2628
	}
L2628:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_338), int32(0))
	mBase = m.M
	v12557 = m.ExcPending
	if v12557 != 0 {
		goto L32
	} else {
		goto L2629
	}
L2629:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_339), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12562 = m.ExcPending
	if v12562 != 0 {
		goto L32
	} else {
		goto L2630
	}
L2630:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2631:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12568 = m.ExcPending
	if v12568 != 0 {
		goto L32
	} else {
		goto L2632
	}
L2632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3536)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v39+int32(3536))
	mBase = m.M
	v12576 = m.ExcPending
	if v12576 != 0 {
		goto L32
	} else {
		goto L2633
	}
L2633:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3435), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12581 = m.ExcPending
	if v12581 != 0 {
		goto L32
	} else {
		goto L2634
	}
L2634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2635:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12587 = m.ExcPending
	if v12587 != 0 {
		goto L32
	} else {
		goto L2636
	}
L2636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3552)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_294), v39+int32(3552))
	mBase = m.M
	v12595 = m.ExcPending
	if v12595 != 0 {
		goto L32
	} else {
		goto L2637
	}
L2637:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3449), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12600 = m.ExcPending
	if v12600 != 0 {
		goto L32
	} else {
		goto L2638
	}
L2638:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3760)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v39+int32(3760))
	mBase = m.M
	v12610 = m.ExcPending
	if v12610 != 0 {
		goto L32
	} else {
		goto L2640
	}
L2640:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3481), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12615 = m.ExcPending
	if v12615 != 0 {
		goto L32
	} else {
		goto L2641
	}
L2641:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2642:
	;
	v12623 = v12617
	goto L2644
L2643:
	;
	v12623 = int32(51)
	goto L2644
L2644:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v12623
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12628 = m.ExcPending
	if v12628 != 0 {
		goto L32
	} else {
		goto L2645
	}
L2645:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12630 = m.ExcPending
	if v12630 != 0 {
		goto L32
	} else {
		goto L2646
	}
L2646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3744)) = v12619
	F_errmsg(m, int32(_a_F_StartupXLOG_297), v39+int32(3744))
	mBase = m.M
	v12636 = m.ExcPending
	if v12636 != 0 {
		goto L32
	} else {
		goto L2647
	}
L2647:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3505), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12641 = m.ExcPending
	if v12641 != 0 {
		goto L32
	} else {
		goto L2648
	}
L2648:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2649:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12647 = m.ExcPending
	if v12647 != 0 {
		goto L32
	} else {
		goto L2650
	}
L2650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3712)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3712))
	mBase = m.M
	v12655 = m.ExcPending
	if v12655 != 0 {
		goto L32
	} else {
		goto L2651
	}
L2651:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3520), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12660 = m.ExcPending
	if v12660 != 0 {
		goto L32
	} else {
		goto L2652
	}
L2652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2653:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12666 = m.ExcPending
	if v12666 != 0 {
		goto L32
	} else {
		goto L2654
	}
L2654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3696)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3696))
	mBase = m.M
	v12674 = m.ExcPending
	if v12674 != 0 {
		goto L32
	} else {
		goto L2655
	}
L2655:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3525), int32(_a_F_StartupXLOG_282))
	mBase = m.M
	v12679 = m.ExcPending
	if v12679 != 0 {
		goto L32
	} else {
		goto L2656
	}
L2656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2657:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v12681
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12693 = m.ExcPending
	if v12693 != 0 {
		goto L32
	} else {
		goto L2658
	}
L2658:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12695 = m.ExcPending
	if v12695 != 0 {
		goto L32
	} else {
		goto L2659
	}
L2659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3824)) = v12683
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3824))
	mBase = m.M
	v12701 = m.ExcPending
	if v12701 != 0 {
		goto L32
	} else {
		goto L2660
	}
L2660:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_340), int32(_a_F_StartupXLOG_341))
	mBase = m.M
	v12706 = m.ExcPending
	if v12706 != 0 {
		goto L32
	} else {
		goto L2661
	}
L2661:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2662:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_342), int32(0))
	mBase = m.M
	v12716 = m.ExcPending
	if v12716 != 0 {
		goto L32
	} else {
		goto L2663
	}
L2663:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_343), int32(_a_F_StartupXLOG_344))
	mBase = m.M
	v12721 = m.ExcPending
	if v12721 != 0 {
		goto L32
	} else {
		goto L2664
	}
L2664:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2665:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3668)) = uint32(v9732)
	v12728 = int64(base.Ui64(v9732) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3664)) = uint32(v12728)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_345), v39+int32(3664))
	mBase = m.M
	v12734 = m.ExcPending
	if v12734 != 0 {
		goto L32
	} else {
		goto L2666
	}
L2666:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_346), int32(_a_F_StartupXLOG_344))
	mBase = m.M
	v12739 = m.ExcPending
	if v12739 != 0 {
		goto L32
	} else {
		goto L2667
	}
L2667:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2668:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3652)) = uint32(v11865)
	v12746 = int64(base.Ui64(v11865) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3648)) = uint32(v12746)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_347), v39+int32(3648))
	mBase = m.M
	v12752 = m.ExcPending
	if v12752 != 0 {
		goto L32
	} else {
		goto L2669
	}
L2669:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_348), int32(_a_F_StartupXLOG_344))
	mBase = m.M
	v12757 = m.ExcPending
	if v12757 != 0 {
		goto L32
	} else {
		goto L2670
	}
L2670:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2671:
	;
	v12763 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3636)) = uint32(v12763)
	v12766 = int64(base.Ui64(v12763) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3632)) = uint32(v12766)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_349), v39+int32(3632))
	mBase = m.M
	v12772 = m.ExcPending
	if v12772 != 0 {
		goto L32
	} else {
		goto L2672
	}
L2672:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_350), int32(_a_F_StartupXLOG_344))
	mBase = m.M
	v12777 = m.ExcPending
	if v12777 != 0 {
		goto L32
	} else {
		goto L2673
	}
L2673:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2674:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_351), int32(0))
	mBase = m.M
	v12786 = m.ExcPending
	if v12786 != 0 {
		goto L32
	} else {
		goto L2675
	}
L2675:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_352), int32(_a_F_StartupXLOG_353))
	mBase = m.M
	v12791 = m.ExcPending
	if v12791 != 0 {
		goto L32
	} else {
		goto L2676
	}
L2676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
