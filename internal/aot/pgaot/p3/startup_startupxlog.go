package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int64
	_ = v919
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int64
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int64
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1135 int64
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int64
	_ = v1151
	var v1152 int64
	_ = v1152
	var v1159 int32
	_ = v1159
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int64
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1299 int64
	_ = v1299
	var v1302 int64
	_ = v1302
	var v1304 int64
	_ = v1304
	var v1305 int64
	_ = v1305
	var v1308 int64
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1325 int64
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int64
	_ = v1335
	var v1336 int64
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1340 int64
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
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int64
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int64
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int64
	_ = v1361
	var v1364 int64
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1381 int64
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int64
	_ = v1398
	var v1401 int64
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1405 int64
	_ = v1405
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1433 int64
	_ = v1433
	var v1436 int64
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1514 int32
	_ = v1514
	var v1526 int32
	_ = v1526
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1584 int64
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1642 int32
	_ = v1642
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1844 int32
	_ = v1844
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1929 int32
	_ = v1929
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1995 int32
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2100 int32
	_ = v2100
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2124 int64
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2128 int64
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int64
	_ = v2146
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int64
	_ = v2155
	var v2158 int64
	_ = v2158
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2172 int64
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2177 int64
	_ = v2177
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int64
	_ = v2191
	var v2194 int64
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int64
	_ = v2212
	var v2213 int64
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int64
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
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int64
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2230 int64
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int64
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2251 int32
	_ = v2251
	var v2253 int64
	_ = v2253
	var v2256 int64
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2273 int64
	_ = v2273
	var v2276 int64
	_ = v2276
	var v2277 int64
	_ = v2277
	var v2280 int64
	_ = v2280
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
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
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int64
	_ = v2330
	var v2331 int64
	_ = v2331
	var v2333 int64
	_ = v2333
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
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
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2363 int64
	_ = v2363
	var v2364 int64
	_ = v2364
	var v2366 int64
	_ = v2366
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2399 int32
	_ = v2399
	var v2405 int32
	_ = v2405
	var v2410 int64
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2429 int32
	_ = v2429
	var v2434 int64
	_ = v2434
	var v2437 int64
	_ = v2437
	var v2443 int32
	_ = v2443
	var v2450 int32
	_ = v2450
	var v2457 int32
	_ = v2457
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2470 int64
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2478 int64
	_ = v2478
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2495 int64
	_ = v2495
	var v2501 int32
	_ = v2501
	var v2507 int32
	_ = v2507
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2555 int32
	_ = v2555
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2571 int32
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2587 int32
	_ = v2587
	var v2592 int32
	_ = v2592
	var v2597 int64
	_ = v2597
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2619 int32
	_ = v2619
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2646 int32
	_ = v2646
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2663 int32
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2681 int64
	_ = v2681
	var v2685 int64
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2697 int64
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2705 int32
	_ = v2705
	var v2708 int64
	_ = v2708
	var v2716 int32
	_ = v2716
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2738 int32
	_ = v2738
	var v2739 int64
	_ = v2739
	var v2743 int64
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2749 int64
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2757 int64
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2763 int64
	_ = v2763
	var v2771 int32
	_ = v2771
	var v2772 int64
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2776 int64
	_ = v2776
	var v2782 int64
	_ = v2782
	var v2798 int32
	_ = v2798
	var v2800 int64
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2813 int32
	_ = v2813
	var v2815 int64
	_ = v2815
	var v2820 int32
	_ = v2820
	var v2823 int64
	_ = v2823
	var v2826 int32
	_ = v2826
	var v2829 int64
	_ = v2829
	var v2835 int32
	_ = v2835
	var v2840 int32
	_ = v2840
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int64
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2853 int64
	_ = v2853
	var v2859 int32
	_ = v2859
	var v2864 int32
	_ = v2864
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int64
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
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2905 int64
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2912 int32
	_ = v2912
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2931 int32
	_ = v2931
	var v2936 int32
	_ = v2936
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2945 int32
	_ = v2945
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v2998 int32
	_ = v2998
	var v2999 int64
	_ = v2999
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3021 int32
	_ = v3021
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
	var v3042 int32
	_ = v3042
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3055 int32
	_ = v3055
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3065 int32
	_ = v3065
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3084 int32
	_ = v3084
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3090 int32
	_ = v3090
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3131 int32
	_ = v3131
	var v3135 int32
	_ = v3135
	var v3137 int32
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3146 int32
	_ = v3146
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3168 int32
	_ = v3168
	var v3196 int32
	_ = v3196
	var v3199 int32
	_ = v3199
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3209 int32
	_ = v3209
	var v3214 int32
	_ = v3214
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3234 int32
	_ = v3234
	var v3237 int32
	_ = v3237
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3248 int32
	_ = v3248
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3259 int32
	_ = v3259
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3277 int32
	_ = v3277
	var v3282 int32
	_ = v3282
	var v3286 int32
	_ = v3286
	var v3291 int32
	_ = v3291
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3309 int32
	_ = v3309
	var v3316 int32
	_ = v3316
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3331 int32
	_ = v3331
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3345 int32
	_ = v3345
	var v3350 int32
	_ = v3350
	var v3355 int32
	_ = v3355
	var v3359 int32
	_ = v3359
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3369 int32
	_ = v3369
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3382 int32
	_ = v3382
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3395 int32
	_ = v3395
	var v3403 int32
	_ = v3403
	var v3408 int32
	_ = v3408
	var v3417 int32
	_ = v3417
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3447 int32
	_ = v3447
	var v3452 int32
	_ = v3452
	var v3462 int32
	_ = v3462
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3477 int32
	_ = v3477
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3493 int32
	_ = v3493
	var v3498 int32
	_ = v3498
	var v3502 int32
	_ = v3502
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3509 int32
	_ = v3509
	var v3513 int32
	_ = v3513
	var v3519 int32
	_ = v3519
	var v3522 int32
	_ = v3522
	var v3528 int32
	_ = v3528
	var v3532 int32
	_ = v3532
	var v3537 int32
	_ = v3537
	var v3541 int32
	_ = v3541
	var v3546 int32
	_ = v3546
	var v3553 int32
	_ = v3553
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3586 int32
	_ = v3586
	var v3592 int32
	_ = v3592
	var v3594 int32
	_ = v3594
	var v3596 int64
	_ = v3596
	var v3598 int64
	_ = v3598
	var v3599 int64
	_ = v3599
	var v3606 int32
	_ = v3606
	var v3610 int32
	_ = v3610
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3620 int64
	_ = v3620
	var v3621 int64
	_ = v3621
	var v3630 int32
	_ = v3630
	var v3667 int32
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3704 int32
	_ = v3704
	var v3706 int32
	_ = v3706
	var v3711 int32
	_ = v3711
	var v3713 int32
	_ = v3713
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3730 int32
	_ = v3730
	var v3735 int32
	_ = v3735
	var v3739 int32
	_ = v3739
	var v3741 int32
	_ = v3741
	var v3749 int32
	_ = v3749
	var v3754 int32
	_ = v3754
	var v3758 int32
	_ = v3758
	var v3760 int32
	_ = v3760
	var v3768 int32
	_ = v3768
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3783 int32
	_ = v3783
	var v3788 int32
	_ = v3788
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3806 int32
	_ = v3806
	var v3811 int32
	_ = v3811
	var v3815 int32
	_ = v3815
	var v3818 int32
	_ = v3818
	var v3827 int32
	_ = v3827
	var v3832 int32
	_ = v3832
	var v3836 int32
	_ = v3836
	var v3839 int32
	_ = v3839
	var v3848 int32
	_ = v3848
	var v3853 int32
	_ = v3853
	var v3855 int32
	_ = v3855
	var v3863 int32
	_ = v3863
	var v3868 int32
	_ = v3868
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3882 int32
	_ = v3882
	var v3887 int32
	_ = v3887
	var v3891 int32
	_ = v3891
	var v3893 int32
	_ = v3893
	var v3902 int32
	_ = v3902
	var v3907 int32
	_ = v3907
	var v3911 int32
	_ = v3911
	var v3914 int32
	_ = v3914
	var v3920 int32
	_ = v3920
	var v3924 int32
	_ = v3924
	var v3929 int32
	_ = v3929
	var v3933 int32
	_ = v3933
	var v3936 int32
	_ = v3936
	var v3942 int32
	_ = v3942
	var v3946 int32
	_ = v3946
	var v3951 int32
	_ = v3951
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3997 int32
	_ = v3997
	var v4002 int32
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4005 int32
	_ = v4005
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4012 int32
	_ = v4012
	var v4043 int32
	_ = v4043
	var v4046 int32
	_ = v4046
	var v4049 int32
	_ = v4049
	var v4052 int32
	_ = v4052
	var v4056 int32
	_ = v4056
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4063 int32
	_ = v4063
	var v4066 int32
	_ = v4066
	var v4067 int32
	_ = v4067
	var v4103 int32
	_ = v4103
	var v4105 int32
	_ = v4105
	var v4107 int32
	_ = v4107
	var v4108 int64
	_ = v4108
	var v4115 int32
	_ = v4115
	var v4116 int32
	_ = v4116
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4125 int32
	_ = v4125
	var v4127 int32
	_ = v4127
	var v4131 int32
	_ = v4131
	var v4132 int32
	_ = v4132
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4144 int32
	_ = v4144
	var v4149 int32
	_ = v4149
	var v4150 int32
	_ = v4150
	var v4154 int32
	_ = v4154
	var v4159 int32
	_ = v4159
	var v4162 int32
	_ = v4162
	var v4163 int32
	_ = v4163
	var v4167 int32
	_ = v4167
	var v4173 int32
	_ = v4173
	var v4175 int32
	_ = v4175
	var v4180 int32
	_ = v4180
	var v4185 int32
	_ = v4185
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4195 int32
	_ = v4195
	var v4200 int32
	_ = v4200
	var v4210 int32
	_ = v4210
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4227 int32
	_ = v4227
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4237 int32
	_ = v4237
	var v4272 int32
	_ = v4272
	var v4274 int32
	_ = v4274
	var v4277 int32
	_ = v4277
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4283 int64
	_ = v4283
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4291 int64
	_ = v4291
	var v4294 int64
	_ = v4294
	var v4300 int32
	_ = v4300
	var v4305 int32
	_ = v4305
	var v4312 int32
	_ = v4312
	var v4316 int32
	_ = v4316
	var v4349 int32
	_ = v4349
	var v4351 int32
	_ = v4351
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4393 int32
	_ = v4393
	var v4400 int32
	_ = v4400
	var v4405 int32
	_ = v4405
	var v4409 int32
	_ = v4409
	var v4412 int32
	_ = v4412
	var v4418 int32
	_ = v4418
	var v4423 int32
	_ = v4423
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4436 int32
	_ = v4436
	var v4441 int32
	_ = v4441
	var v4445 int32
	_ = v4445
	var v4447 int32
	_ = v4447
	var v4457 int32
	_ = v4457
	var v4462 int32
	_ = v4462
	var v4466 int32
	_ = v4466
	var v4469 int32
	_ = v4469
	var v4473 int32
	_ = v4473
	var v4478 int32
	_ = v4478
	var v4482 int32
	_ = v4482
	var v4485 int32
	_ = v4485
	var v4492 int32
	_ = v4492
	var v4497 int32
	_ = v4497
	var v4501 int32
	_ = v4501
	var v4503 int32
	_ = v4503
	var v4510 int32
	_ = v4510
	var v4515 int32
	_ = v4515
	var v4517 int32
	_ = v4517
	var v4519 int32
	_ = v4519
	var v4520 int32
	_ = v4520
	var v4523 int64
	_ = v4523
	var v4525 int64
	_ = v4525
	var v4528 int32
	_ = v4528
	var v4530 int32
	_ = v4530
	var v4532 int32
	_ = v4532
	var v4536 int32
	_ = v4536
	var v4537 int32
	_ = v4537
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4542 int32
	_ = v4542
	var v4543 int32
	_ = v4543
	var v4547 int32
	_ = v4547
	var v4579 int32
	_ = v4579
	var v4580 int32
	_ = v4580
	var v4583 int32
	_ = v4583
	var v4587 int32
	_ = v4587
	var v4589 int32
	_ = v4589
	var v4590 int64
	_ = v4590
	var v4598 int32
	_ = v4598
	var v4602 int32
	_ = v4602
	var v4606 int32
	_ = v4606
	var v4612 int32
	_ = v4612
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4624 int32
	_ = v4624
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4630 int32
	_ = v4630
	var v4633 int32
	_ = v4633
	var v4637 int32
	_ = v4637
	var v4638 int32
	_ = v4638
	var v4646 int32
	_ = v4646
	var v4652 int32
	_ = v4652
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4666 int32
	_ = v4666
	var v4672 int64
	_ = v4672
	var v4676 int32
	_ = v4676
	var v4678 int32
	_ = v4678
	var v4679 int32
	_ = v4679
	var v4682 int64
	_ = v4682
	var v4686 int32
	_ = v4686
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4726 int32
	_ = v4726
	var v4730 int32
	_ = v4730
	var v4732 int32
	_ = v4732
	var v4735 int32
	_ = v4735
	var v4737 int32
	_ = v4737
	var v4741 int32
	_ = v4741
	var v4743 int32
	_ = v4743
	var v4748 int32
	_ = v4748
	var v4749 int32
	_ = v4749
	var v4758 int32
	_ = v4758
	var v4762 int32
	_ = v4762
	var v4763 int32
	_ = v4763
	var v4767 int32
	_ = v4767
	var v4774 int32
	_ = v4774
	var v4778 int32
	_ = v4778
	var v4779 int32
	_ = v4779
	var v4783 int32
	_ = v4783
	var v4788 int32
	_ = v4788
	var v4790 int32
	_ = v4790
	var v4793 int32
	_ = v4793
	var v4797 int32
	_ = v4797
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4802 int64
	_ = v4802
	var v4803 int64
	_ = v4803
	var v4816 int32
	_ = v4816
	var v4858 int32
	_ = v4858
	var v4866 int32
	_ = v4866
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4876 int32
	_ = v4876
	var v4878 int32
	_ = v4878
	var v4882 int32
	_ = v4882
	var v4886 int32
	_ = v4886
	var v4890 int32
	_ = v4890
	var v4892 int32
	_ = v4892
	var v4895 int32
	_ = v4895
	var v4898 int32
	_ = v4898
	var v4899 int32
	_ = v4899
	var v4906 int32
	_ = v4906
	var v4911 int32
	_ = v4911
	var v4914 int32
	_ = v4914
	var v4915 int32
	_ = v4915
	var v4919 int32
	_ = v4919
	var v4924 int32
	_ = v4924
	var v4925 int32
	_ = v4925
	var v4929 int32
	_ = v4929
	var v4934 int32
	_ = v4934
	var v4939 int32
	_ = v4939
	var v4943 int32
	_ = v4943
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4948 int64
	_ = v4948
	var v4949 int64
	_ = v4949
	var v4959 int32
	_ = v4959
	var v5004 int32
	_ = v5004
	var v5012 int32
	_ = v5012
	var v5015 int32
	_ = v5015
	var v5016 int32
	_ = v5016
	var v5021 int32
	_ = v5021
	var v5023 int32
	_ = v5023
	var v5026 int32
	_ = v5026
	var v5030 int32
	_ = v5030
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5041 int32
	_ = v5041
	var v5042 int32
	_ = v5042
	var v5048 int32
	_ = v5048
	var v5053 int32
	_ = v5053
	var v5054 int32
	_ = v5054
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5102 int32
	_ = v5102
	var v5103 int32
	_ = v5103
	var v5108 int32
	_ = v5108
	var v5109 int32
	_ = v5109
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5122 int32
	_ = v5122
	var v5126 int32
	_ = v5126
	var v5134 int32
	_ = v5134
	var v5135 int32
	_ = v5135
	var v5145 int32
	_ = v5145
	var v5152 int32
	_ = v5152
	var v5156 int32
	_ = v5156
	var v5160 int32
	_ = v5160
	var v5165 int32
	_ = v5165
	var v5166 int32
	_ = v5166
	var v5176 int32
	_ = v5176
	var v5178 int32
	_ = v5178
	var v5182 int32
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5190 int32
	_ = v5190
	var v5193 int32
	_ = v5193
	var v5194 int32
	_ = v5194
	var v5196 int32
	_ = v5196
	var v5197 int32
	_ = v5197
	var v5200 int32
	_ = v5200
	var v5201 int32
	_ = v5201
	var v5202 int32
	_ = v5202
	var v5206 int32
	_ = v5206
	var v5207 int32
	_ = v5207
	var v5210 int32
	_ = v5210
	var v5219 int32
	_ = v5219
	var v5222 int32
	_ = v5222
	var v5224 int32
	_ = v5224
	var v5231 int32
	_ = v5231
	var v5232 int32
	_ = v5232
	var v5237 int32
	_ = v5237
	var v5238 int32
	_ = v5238
	var v5247 int32
	_ = v5247
	var v5249 int32
	_ = v5249
	var v5251 int32
	_ = v5251
	var v5255 int32
	_ = v5255
	var v5263 int32
	_ = v5263
	var v5264 int32
	_ = v5264
	var v5269 int64
	_ = v5269
	var v5271 int64
	_ = v5271
	var v5277 int32
	_ = v5277
	var v5282 int32
	_ = v5282
	var v5290 int32
	_ = v5290
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5300 int64
	_ = v5300
	var v5302 int64
	_ = v5302
	var v5308 int32
	_ = v5308
	var v5314 int32
	_ = v5314
	var v5315 int32
	_ = v5315
	var v5320 int32
	_ = v5320
	var v5321 int32
	_ = v5321
	var v5329 int32
	_ = v5329
	var v5335 int32
	_ = v5335
	var v5336 int32
	_ = v5336
	var v5341 int32
	_ = v5341
	var v5342 int32
	_ = v5342
	var v5345 int32
	_ = v5345
	var v5352 int32
	_ = v5352
	var v5354 int32
	_ = v5354
	var v5356 int32
	_ = v5356
	var v5358 int32
	_ = v5358
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5378 int32
	_ = v5378
	var v5381 int32
	_ = v5381
	var v5391 int32
	_ = v5391
	var v5399 int32
	_ = v5399
	var v5403 int32
	_ = v5403
	var v5404 int32
	_ = v5404
	var v5413 int32
	_ = v5413
	var v5415 int32
	_ = v5415
	var v5416 int32
	_ = v5416
	var v5421 int32
	_ = v5421
	var v5422 int32
	_ = v5422
	var v5428 int32
	_ = v5428
	var v5433 int32
	_ = v5433
	var v5436 int32
	_ = v5436
	var v5437 int32
	_ = v5437
	var v5445 int32
	_ = v5445
	var v5450 int32
	_ = v5450
	var v5451 int32
	_ = v5451
	var v5452 int32
	_ = v5452
	var v5457 int32
	_ = v5457
	var v5458 int32
	_ = v5458
	var v5467 int32
	_ = v5467
	var v5473 int32
	_ = v5473
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5482 int32
	_ = v5482
	var v5487 int32
	_ = v5487
	var v5488 int64
	_ = v5488
	var v5490 int32
	_ = v5490
	var v5491 int32
	_ = v5491
	var v5496 int32
	_ = v5496
	var v5497 int32
	_ = v5497
	var v5509 int32
	_ = v5509
	var v5515 int32
	_ = v5515
	var v5520 int32
	_ = v5520
	var v5521 int32
	_ = v5521
	var v5522 int32
	_ = v5522
	var v5526 int32
	_ = v5526
	var v5528 int32
	_ = v5528
	var v5531 int32
	_ = v5531
	var v5532 int32
	_ = v5532
	var v5536 int64
	_ = v5536
	var v5538 int64
	_ = v5538
	var v5544 int32
	_ = v5544
	var v5546 int32
	_ = v5546
	var v5547 int32
	_ = v5547
	var v5548 int32
	_ = v5548
	var v5555 int32
	_ = v5555
	var v5570 int32
	_ = v5570
	var v5572 int32
	_ = v5572
	var v5580 int32
	_ = v5580
	var v5582 int32
	_ = v5582
	var v5584 int32
	_ = v5584
	var v5585 int32
	_ = v5585
	var v5587 int32
	_ = v5587
	var v5588 int32
	_ = v5588
	var v5590 int32
	_ = v5590
	var v5591 int32
	_ = v5591
	var v5592 int32
	_ = v5592
	var v5597 int32
	_ = v5597
	var v5598 int32
	_ = v5598
	var v5604 int32
	_ = v5604
	var v5606 int32
	_ = v5606
	var v5608 int32
	_ = v5608
	var v5611 int32
	_ = v5611
	var v5621 int32
	_ = v5621
	var v5627 int32
	_ = v5627
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5632 int32
	_ = v5632
	var v5633 int32
	_ = v5633
	var v5634 int32
	_ = v5634
	var v5638 int32
	_ = v5638
	var v5639 int32
	_ = v5639
	var v5643 int64
	_ = v5643
	var v5645 int64
	_ = v5645
	var v5651 int32
	_ = v5651
	var v5656 int32
	_ = v5656
	var v5658 int64
	_ = v5658
	var v5660 int64
	_ = v5660
	var v5666 int32
	_ = v5666
	var v5671 int32
	_ = v5671
	var v5672 int32
	_ = v5672
	var v5679 int32
	_ = v5679
	var v5680 int32
	_ = v5680
	var v5687 int32
	_ = v5687
	var v5692 int32
	_ = v5692
	var v5693 int32
	_ = v5693
	var v5703 int32
	_ = v5703
	var v5708 int32
	_ = v5708
	var v5711 int32
	_ = v5711
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5725 int32
	_ = v5725
	var v5730 int32
	_ = v5730
	var v5767 int32
	_ = v5767
	var v5768 int32
	_ = v5768
	var v5775 int32
	_ = v5775
	var v5780 int32
	_ = v5780
	var v5784 int32
	_ = v5784
	var v5785 int32
	_ = v5785
	var v5786 int32
	_ = v5786
	var v5789 int64
	_ = v5789
	var v5790 int64
	_ = v5790
	var v5800 int32
	_ = v5800
	var v5845 int32
	_ = v5845
	var v5853 int32
	_ = v5853
	var v5856 int32
	_ = v5856
	var v5857 int32
	_ = v5857
	var v5862 int32
	_ = v5862
	var v5864 int32
	_ = v5864
	var v5867 int32
	_ = v5867
	var v5871 int32
	_ = v5871
	var v5906 int32
	_ = v5906
	var v5907 int32
	_ = v5907
	var v5910 int32
	_ = v5910
	var v5911 int32
	_ = v5911
	var v5918 int32
	_ = v5918
	var v5923 int32
	_ = v5923
	var v5925 int32
	_ = v5925
	var v5998 int32
	_ = v5998
	var v5999 int32
	_ = v5999
	var v6002 int32
	_ = v6002
	var v6010 int32
	_ = v6010
	var v6013 int32
	_ = v6013
	var v6017 int32
	_ = v6017
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6028 int32
	_ = v6028
	var v6033 int32
	_ = v6033
	var v6035 int32
	_ = v6035
	var v6037 int32
	_ = v6037
	var v6038 int32
	_ = v6038
	var v6041 int32
	_ = v6041
	var v6042 int32
	_ = v6042
	var v6046 int32
	_ = v6046
	var v6047 int32
	_ = v6047
	var v6048 int32
	_ = v6048
	var v6051 int32
	_ = v6051
	var v6052 int32
	_ = v6052
	var v6056 int32
	_ = v6056
	var v6057 int32
	_ = v6057
	var v6060 int32
	_ = v6060
	var v6064 int32
	_ = v6064
	var v6065 int64
	_ = v6065
	var v6067 int64
	_ = v6067
	var v6070 int32
	_ = v6070
	var v6073 int32
	_ = v6073
	var v6074 int32
	_ = v6074
	var v6076 int32
	_ = v6076
	var v6079 int32
	_ = v6079
	var v6080 int32
	_ = v6080
	var v6083 int32
	_ = v6083
	var v6084 int32
	_ = v6084
	var v6085 int32
	_ = v6085
	var v6119 int32
	_ = v6119
	var v6122 int32
	_ = v6122
	var v6125 int32
	_ = v6125
	var v6128 int32
	_ = v6128
	var v6135 int32
	_ = v6135
	var v6140 int32
	_ = v6140
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6152 int32
	_ = v6152
	var v6156 int32
	_ = v6156
	var v6161 int32
	_ = v6161
	var v6165 int32
	_ = v6165
	var v6166 int32
	_ = v6166
	var v6202 int32
	_ = v6202
	var v6207 int32
	_ = v6207
	var v6211 int32
	_ = v6211
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6223 int32
	_ = v6223
	var v6228 int32
	_ = v6228
	var v6229 int32
	_ = v6229
	var v6231 int32
	_ = v6231
	var v6239 int32
	_ = v6239
	var v6241 int32
	_ = v6241
	var v6242 int32
	_ = v6242
	var v6250 int32
	_ = v6250
	var v6251 int32
	_ = v6251
	var v6255 int32
	_ = v6255
	var v6257 int32
	_ = v6257
	var v6259 int32
	_ = v6259
	var v6263 int32
	_ = v6263
	var v6264 int32
	_ = v6264
	var v6266 int32
	_ = v6266
	var v6269 int32
	_ = v6269
	var v6274 int64
	_ = v6274
	var v6277 int32
	_ = v6277
	var v6279 int32
	_ = v6279
	var v6284 int32
	_ = v6284
	var v6291 int32
	_ = v6291
	var v6292 int32
	_ = v6292
	var v6293 int32
	_ = v6293
	var v6295 int32
	_ = v6295
	var v6296 int32
	_ = v6296
	var v6300 int32
	_ = v6300
	var v6332 int32
	_ = v6332
	var v6338 int32
	_ = v6338
	var v6339 int32
	_ = v6339
	var v6343 int32
	_ = v6343
	var v6345 int32
	_ = v6345
	var v6349 int32
	_ = v6349
	var v6354 int32
	_ = v6354
	var v6386 int32
	_ = v6386
	var v6390 int32
	_ = v6390
	var v6395 int32
	_ = v6395
	var v6430 int32
	_ = v6430
	var v6432 int32
	_ = v6432
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6440 int32
	_ = v6440
	var v6447 int32
	_ = v6447
	var v6449 int64
	_ = v6449
	var v6451 int64
	_ = v6451
	var v6454 int32
	_ = v6454
	var v6459 int32
	_ = v6459
	var v6461 int32
	_ = v6461
	var v6462 int64
	_ = v6462
	var v6464 int64
	_ = v6464
	var v6467 int32
	_ = v6467
	var v6468 int64
	_ = v6468
	var v6469 int32
	_ = v6469
	var v6471 int32
	_ = v6471
	var v6472 int64
	_ = v6472
	var v6479 int32
	_ = v6479
	var v6487 int32
	_ = v6487
	var v6488 int32
	_ = v6488
	var v6489 int32
	_ = v6489
	var v6492 int64
	_ = v6492
	var v6493 int64
	_ = v6493
	var v6504 int32
	_ = v6504
	var v6509 int32
	_ = v6509
	var v6511 int32
	_ = v6511
	var v6513 int32
	_ = v6513
	var v6515 int64
	_ = v6515
	var v6517 int64
	_ = v6517
	var v6520 int32
	_ = v6520
	var v6522 int32
	_ = v6522
	var v6524 int32
	_ = v6524
	var v6527 int32
	_ = v6527
	var v6528 int32
	_ = v6528
	var v6529 int32
	_ = v6529
	var v6532 int32
	_ = v6532
	var v6540 int32
	_ = v6540
	var v6542 int32
	_ = v6542
	var v6543 int64
	_ = v6543
	var v6546 int64
	_ = v6546
	var v6552 int32
	_ = v6552
	var v6557 int32
	_ = v6557
	var v6558 int32
	_ = v6558
	var v6562 int32
	_ = v6562
	var v6563 int32
	_ = v6563
	var v6564 int32
	_ = v6564
	var v6567 int32
	_ = v6567
	var v6568 int32
	_ = v6568
	var v6577 int32
	_ = v6577
	var v6581 int32
	_ = v6581
	var v6615 int32
	_ = v6615
	var v6618 int32
	_ = v6618
	var v6621 int32
	_ = v6621
	var v6625 int32
	_ = v6625
	var v6627 int32
	_ = v6627
	var v6630 int32
	_ = v6630
	var v6634 int32
	_ = v6634
	var v6637 int32
	_ = v6637
	var v6642 int32
	_ = v6642
	var v6643 int32
	_ = v6643
	var v6645 int32
	_ = v6645
	var v6646 int64
	_ = v6646
	var v6649 int64
	_ = v6649
	var v6655 int32
	_ = v6655
	var v6660 int32
	_ = v6660
	var v6663 int32
	_ = v6663
	var v6667 int32
	_ = v6667
	var v6668 int32
	_ = v6668
	var v6673 int32
	_ = v6673
	var v6703 int32
	_ = v6703
	var v6711 int32
	_ = v6711
	var v6713 int32
	_ = v6713
	var v6716 int32
	_ = v6716
	var v6717 int64
	_ = v6717
	var v6719 int64
	_ = v6719
	var v6725 int32
	_ = v6725
	var v6727 int32
	_ = v6727
	var v6742 int32
	_ = v6742
	var v6743 int32
	_ = v6743
	var v6747 int32
	_ = v6747
	var v6748 int64
	_ = v6748
	var v6749 int32
	_ = v6749
	var v6751 int32
	_ = v6751
	var v6753 int32
	_ = v6753
	var v6757 int64
	_ = v6757
	var v6763 int32
	_ = v6763
	var v6768 int32
	_ = v6768
	var v6771 int32
	_ = v6771
	var v6773 int32
	_ = v6773
	var v6774 int32
	_ = v6774
	var v6777 int32
	_ = v6777
	var v6779 int32
	_ = v6779
	var v6783 int32
	_ = v6783
	var v6785 int32
	_ = v6785
	var v6789 int32
	_ = v6789
	var v6796 int32
	_ = v6796
	var v6797 int32
	_ = v6797
	var v6801 int32
	_ = v6801
	var v6806 int32
	_ = v6806
	var v6808 int64
	_ = v6808
	var v6814 int32
	_ = v6814
	var v6825 int32
	_ = v6825
	var v6828 int64
	_ = v6828
	var v6830 int64
	_ = v6830
	var v6838 int32
	_ = v6838
	var v6848 int32
	_ = v6848
	var v6849 int32
	_ = v6849
	var v6853 int64
	_ = v6853
	var v6856 int64
	_ = v6856
	var v6862 int32
	_ = v6862
	var v6867 int32
	_ = v6867
	var v6869 int32
	_ = v6869
	var v6870 int32
	_ = v6870
	var v6873 int32
	_ = v6873
	var v6878 int32
	_ = v6878
	var v6880 int32
	_ = v6880
	var v6882 int32
	_ = v6882
	var v6883 int32
	_ = v6883
	var v6889 int64
	_ = v6889
	var v6894 int32
	_ = v6894
	var v6898 int32
	_ = v6898
	var v6900 int32
	_ = v6900
	var v6906 int32
	_ = v6906
	var v6909 int32
	_ = v6909
	var v6911 int32
	_ = v6911
	var v6917 int32
	_ = v6917
	var v6921 int32
	_ = v6921
	var v6923 int32
	_ = v6923
	var v6926 int32
	_ = v6926
	var v6930 int32
	_ = v6930
	var v6935 int32
	_ = v6935
	var v6936 int32
	_ = v6936
	var v6937 int32
	_ = v6937
	var v6940 int32
	_ = v6940
	var v6944 int32
	_ = v6944
	var v6949 int32
	_ = v6949
	var v6950 int32
	_ = v6950
	var v6951 int32
	_ = v6951
	var v6954 int32
	_ = v6954
	var v6958 int32
	_ = v6958
	var v6965 int32
	_ = v6965
	var v6968 int32
	_ = v6968
	var v6976 int32
	_ = v6976
	var v6977 int32
	_ = v6977
	var v6981 int32
	_ = v6981
	var v6982 int32
	_ = v6982
	var v6983 int32
	_ = v6983
	var v6988 int64
	_ = v6988
	var v6989 int64
	_ = v6989
	var v6997 int32
	_ = v6997
	var v6999 int32
	_ = v6999
	var v7000 int32
	_ = v7000
	var v7002 int32
	_ = v7002
	var v7003 int32
	_ = v7003
	var v7009 int64
	_ = v7009
	var v7014 int32
	_ = v7014
	var v7018 int32
	_ = v7018
	var v7020 int32
	_ = v7020
	var v7026 int32
	_ = v7026
	var v7029 int32
	_ = v7029
	var v7031 int32
	_ = v7031
	var v7037 int32
	_ = v7037
	var v7041 int32
	_ = v7041
	var v7043 int32
	_ = v7043
	var v7046 int32
	_ = v7046
	var v7050 int32
	_ = v7050
	var v7055 int32
	_ = v7055
	var v7056 int32
	_ = v7056
	var v7057 int32
	_ = v7057
	var v7060 int32
	_ = v7060
	var v7064 int32
	_ = v7064
	var v7071 int32
	_ = v7071
	var v7074 int32
	_ = v7074
	var v7082 int32
	_ = v7082
	var v7083 int32
	_ = v7083
	var v7087 int32
	_ = v7087
	var v7088 int32
	_ = v7088
	var v7089 int32
	_ = v7089
	var v7094 int64
	_ = v7094
	var v7095 int64
	_ = v7095
	var v7103 int32
	_ = v7103
	var v7104 int32
	_ = v7104
	var v7106 int32
	_ = v7106
	var v7107 int32
	_ = v7107
	var v7108 int32
	_ = v7108
	var v7110 int32
	_ = v7110
	var v7111 int32
	_ = v7111
	var v7114 int32
	_ = v7114
	var v7121 int32
	_ = v7121
	var v7123 int32
	_ = v7123
	var v7124 int32
	_ = v7124
	var v7125 int32
	_ = v7125
	var v7126 int32
	_ = v7126
	var v7136 int32
	_ = v7136
	var v7139 int32
	_ = v7139
	var v7149 int32
	_ = v7149
	var v7150 int64
	_ = v7150
	var v7154 int64
	_ = v7154
	var v7174 int32
	_ = v7174
	var v7178 int32
	_ = v7178
	var v7184 int32
	_ = v7184
	var v7190 int32
	_ = v7190
	var v7191 int32
	_ = v7191
	var v7192 int32
	_ = v7192
	var v7195 int32
	_ = v7195
	var v7197 int32
	_ = v7197
	var v7201 int32
	_ = v7201
	var v7202 int32
	_ = v7202
	var v7205 int32
	_ = v7205
	var v7211 int32
	_ = v7211
	var v7212 int64
	_ = v7212
	var v7216 int32
	_ = v7216
	var v7217 int32
	_ = v7217
	var v7218 int32
	_ = v7218
	var v7221 int64
	_ = v7221
	var v7222 int64
	_ = v7222
	var v7230 int64
	_ = v7230
	var v7234 int64
	_ = v7234
	var v7240 int64
	_ = v7240
	var v7249 int64
	_ = v7249
	var v7252 int32
	_ = v7252
	var v7256 int32
	_ = v7256
	var v7262 int32
	_ = v7262
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7300 int64
	_ = v7300
	var v7304 int32
	_ = v7304
	var v7305 int32
	_ = v7305
	var v7306 int32
	_ = v7306
	var v7309 int64
	_ = v7309
	var v7310 int64
	_ = v7310
	var v7318 int64
	_ = v7318
	var v7321 int64
	_ = v7321
	var v7327 int64
	_ = v7327
	var v7336 int64
	_ = v7336
	var v7339 int32
	_ = v7339
	var v7344 int32
	_ = v7344
	var v7345 int32
	_ = v7345
	var v7351 int32
	_ = v7351
	var v7356 int32
	_ = v7356
	var v7358 int32
	_ = v7358
	var v7363 int32
	_ = v7363
	var v7364 int32
	_ = v7364
	var v7366 int32
	_ = v7366
	var v7372 int32
	_ = v7372
	var v7373 int32
	_ = v7373
	var v7374 int32
	_ = v7374
	var v7412 int32
	_ = v7412
	var v7413 int32
	_ = v7413
	var v7418 int32
	_ = v7418
	var v7453 int32
	_ = v7453
	var v7454 int32
	_ = v7454
	var v7455 int32
	_ = v7455
	var v7461 int32
	_ = v7461
	var v7466 int32
	_ = v7466
	var v7468 int32
	_ = v7468
	var v7469 int32
	_ = v7469
	var v7470 int32
	_ = v7470
	var v7472 int32
	_ = v7472
	var v7476 int32
	_ = v7476
	var v7477 int32
	_ = v7477
	var v7478 int32
	_ = v7478
	var v7479 int32
	_ = v7479
	var v7481 int32
	_ = v7481
	var v7484 int64
	_ = v7484
	var v7486 int32
	_ = v7486
	var v7487 int32
	_ = v7487
	var v7493 int32
	_ = v7493
	var v7496 int32
	_ = v7496
	var v7499 int32
	_ = v7499
	var v7500 int32
	_ = v7500
	var v7503 int32
	_ = v7503
	var v7510 int32
	_ = v7510
	var v7511 int32
	_ = v7511
	var v7512 int32
	_ = v7512
	var v7514 int32
	_ = v7514
	var v7520 int32
	_ = v7520
	var v7526 int32
	_ = v7526
	var v7531 int64
	_ = v7531
	var v7534 int32
	_ = v7534
	var v7536 int32
	_ = v7536
	var v7539 int32
	_ = v7539
	var v7542 int32
	_ = v7542
	var v7543 int32
	_ = v7543
	var v7547 int32
	_ = v7547
	var v7554 int32
	_ = v7554
	var v7555 int64
	_ = v7555
	var v7557 int32
	_ = v7557
	var v7558 int32
	_ = v7558
	var v7563 int32
	_ = v7563
	var v7566 int32
	_ = v7566
	var v7570 int32
	_ = v7570
	var v7572 int32
	_ = v7572
	var v7573 int32
	_ = v7573
	var v7574 int32
	_ = v7574
	var v7576 int32
	_ = v7576
	var v7581 int32
	_ = v7581
	var v7582 int64
	_ = v7582
	var v7583 int64
	_ = v7583
	var v7585 int64
	_ = v7585
	var v7587 int64
	_ = v7587
	var v7594 int32
	_ = v7594
	var v7595 int32
	_ = v7595
	var v7596 int32
	_ = v7596
	var v7597 int32
	_ = v7597
	var v7601 int64
	_ = v7601
	var v7607 int32
	_ = v7607
	var v7612 int32
	_ = v7612
	var v7615 int64
	_ = v7615
	var v7617 int64
	_ = v7617
	var v7618 int32
	_ = v7618
	var v7619 int64
	_ = v7619
	var v7622 int32
	_ = v7622
	var v7623 int32
	_ = v7623
	var v7628 int32
	_ = v7628
	var v7633 int32
	_ = v7633
	var v7639 int64
	_ = v7639
	var v7642 int64
	_ = v7642
	var v7643 int64
	_ = v7643
	var v7646 int64
	_ = v7646
	var v7652 int32
	_ = v7652
	var v7657 int32
	_ = v7657
	var v7663 int32
	_ = v7663
	var v7665 int32
	_ = v7665
	var v7668 int32
	_ = v7668
	var v7672 int32
	_ = v7672
	var v7673 int32
	_ = v7673
	var v7675 int32
	_ = v7675
	var v7676 int32
	_ = v7676
	var v7681 int32
	_ = v7681
	var v7682 int32
	_ = v7682
	var v7684 int32
	_ = v7684
	var v7685 int32
	_ = v7685
	var v7687 int32
	_ = v7687
	var v7688 int32
	_ = v7688
	var v7689 int32
	_ = v7689
	var v7690 int32
	_ = v7690
	var v7695 int32
	_ = v7695
	var v7702 int32
	_ = v7702
	var v7732 int32
	_ = v7732
	var v7734 int32
	_ = v7734
	var v7736 int32
	_ = v7736
	var v7738 int32
	_ = v7738
	var v7739 int32
	_ = v7739
	var v7741 int32
	_ = v7741
	var v7742 int32
	_ = v7742
	var v7746 int32
	_ = v7746
	var v7747 int32
	_ = v7747
	var v7751 int32
	_ = v7751
	var v7752 int32
	_ = v7752
	var v7754 int64
	_ = v7754
	var v7756 int32
	_ = v7756
	var v7758 int32
	_ = v7758
	var v7766 int32
	_ = v7766
	var v7769 int32
	_ = v7769
	var v7773 int32
	_ = v7773
	var v7774 int64
	_ = v7774
	var v7776 int32
	_ = v7776
	var v7780 int32
	_ = v7780
	var v7781 int32
	_ = v7781
	var v7784 int32
	_ = v7784
	var v7785 int32
	_ = v7785
	var v7790 int32
	_ = v7790
	var v7792 int32
	_ = v7792
	var v7796 int32
	_ = v7796
	var v7802 int32
	_ = v7802
	var v7804 int32
	_ = v7804
	var v7810 int32
	_ = v7810
	var v7814 int32
	_ = v7814
	var v7815 int64
	_ = v7815
	var v7817 int32
	_ = v7817
	var v7818 int64
	_ = v7818
	var v7821 int64
	_ = v7821
	var v7825 int32
	_ = v7825
	var v7826 int32
	_ = v7826
	var v7827 int32
	_ = v7827
	var v7831 int32
	_ = v7831
	var v7832 int32
	_ = v7832
	var v7834 int32
	_ = v7834
	var v7836 int32
	_ = v7836
	var v7837 int32
	_ = v7837
	var v7839 int32
	_ = v7839
	var v7841 int32
	_ = v7841
	var v7843 int32
	_ = v7843
	var v7844 int32
	_ = v7844
	var v7852 int32
	_ = v7852
	var v7853 int32
	_ = v7853
	var v7854 int32
	_ = v7854
	var v7857 int32
	_ = v7857
	var v7858 int32
	_ = v7858
	var v7860 int32
	_ = v7860
	var v7861 int32
	_ = v7861
	var v7863 int32
	_ = v7863
	var v7865 int32
	_ = v7865
	var v7875 int32
	_ = v7875
	var v7876 int32
	_ = v7876
	var v7877 int32
	_ = v7877
	var v7880 int32
	_ = v7880
	var v7881 int32
	_ = v7881
	var v7882 int32
	_ = v7882
	var v7885 int32
	_ = v7885
	var v7886 int32
	_ = v7886
	var v7888 int32
	_ = v7888
	var v7893 int32
	_ = v7893
	var v7906 int32
	_ = v7906
	var v7909 int32
	_ = v7909
	var v7910 int32
	_ = v7910
	var v7911 int32
	_ = v7911
	var v7948 int32
	_ = v7948
	var v7951 int32
	_ = v7951
	var v7952 int32
	_ = v7952
	var v7956 int32
	_ = v7956
	var v7963 int32
	_ = v7963
	var v7965 int32
	_ = v7965
	var v7966 int64
	_ = v7966
	var v7968 int64
	_ = v7968
	var v7974 int32
	_ = v7974
	var v7978 int32
	_ = v7978
	var v7983 int32
	_ = v7983
	var v7985 int32
	_ = v7985
	var v7987 int32
	_ = v7987
	var v7990 int32
	_ = v7990
	var v7992 int32
	_ = v7992
	var v7993 int64
	_ = v7993
	var v7995 int32
	_ = v7995
	var v7996 int32
	_ = v7996
	var v7998 int32
	_ = v7998
	var v8003 int32
	_ = v8003
	var v8007 int32
	_ = v8007
	var v8008 int32
	_ = v8008
	var v8009 int32
	_ = v8009
	var v8010 int32
	_ = v8010
	var v8012 int32
	_ = v8012
	var v8023 int32
	_ = v8023
	var v8025 int32
	_ = v8025
	var v8027 int32
	_ = v8027
	var v8030 int32
	_ = v8030
	var v8033 int32
	_ = v8033
	var v8036 int32
	_ = v8036
	var v8037 int32
	_ = v8037
	var v8040 int32
	_ = v8040
	var v8041 int32
	_ = v8041
	var v8044 int32
	_ = v8044
	var v8051 int32
	_ = v8051
	var v8052 int32
	_ = v8052
	var v8055 int32
	_ = v8055
	var v8064 int64
	_ = v8064
	var v8066 int32
	_ = v8066
	var v8073 int32
	_ = v8073
	var v8077 int32
	_ = v8077
	var v8089 int32
	_ = v8089
	var v8090 int32
	_ = v8090
	var v8091 int32
	_ = v8091
	var v8093 int32
	_ = v8093
	var v8097 int32
	_ = v8097
	var v8098 int32
	_ = v8098
	var v8100 int32
	_ = v8100
	var v8101 int32
	_ = v8101
	var v8102 int32
	_ = v8102
	var v8104 int32
	_ = v8104
	var v8110 int32
	_ = v8110
	var v8111 int32
	_ = v8111
	var v8112 int32
	_ = v8112
	var v8113 int32
	_ = v8113
	var v8116 int32
	_ = v8116
	var v8123 int32
	_ = v8123
	var v8124 int32
	_ = v8124
	var v8125 int32
	_ = v8125
	var v8128 int32
	_ = v8128
	var v8131 int32
	_ = v8131
	var v8136 int32
	_ = v8136
	var v8137 int32
	_ = v8137
	var v8139 int32
	_ = v8139
	var v8141 int32
	_ = v8141
	var v8145 int32
	_ = v8145
	var v8146 int32
	_ = v8146
	var v8147 int32
	_ = v8147
	var v8152 int32
	_ = v8152
	var v8153 int32
	_ = v8153
	var v8154 int32
	_ = v8154
	var v8157 int32
	_ = v8157
	var v8158 int32
	_ = v8158
	var v8159 int32
	_ = v8159
	var v8161 int32
	_ = v8161
	var v8165 int32
	_ = v8165
	var v8166 int32
	_ = v8166
	var v8168 int32
	_ = v8168
	var v8170 int32
	_ = v8170
	var v8172 int32
	_ = v8172
	var v8173 int32
	_ = v8173
	var v8176 int32
	_ = v8176
	var v8183 int32
	_ = v8183
	var v8188 int32
	_ = v8188
	var v8189 int32
	_ = v8189
	var v8193 int64
	_ = v8193
	var v8194 int32
	_ = v8194
	var v8195 int32
	_ = v8195
	var v8203 int32
	_ = v8203
	var v8208 int32
	_ = v8208
	var v8212 int32
	_ = v8212
	var v8217 int64
	_ = v8217
	var v8219 int64
	_ = v8219
	var v8222 int32
	_ = v8222
	var v8230 int32
	_ = v8230
	var v8237 int32
	_ = v8237
	var v8238 int32
	_ = v8238
	var v8242 int64
	_ = v8242
	var v8245 int64
	_ = v8245
	var v8251 int32
	_ = v8251
	var v8256 int32
	_ = v8256
	var v8261 int32
	_ = v8261
	var v8262 int32
	_ = v8262
	var v8263 int32
	_ = v8263
	var v8270 int32
	_ = v8270
	var v8273 int32
	_ = v8273
	var v8282 int32
	_ = v8282
	var v8283 int32
	_ = v8283
	var v8284 int32
	_ = v8284
	var v8285 int64
	_ = v8285
	var v8289 int32
	_ = v8289
	var v8296 int32
	_ = v8296
	var v8298 int32
	_ = v8298
	var v8302 int32
	_ = v8302
	var v8304 int32
	_ = v8304
	var v8305 int64
	_ = v8305
	var v8307 int32
	_ = v8307
	var v8310 int32
	_ = v8310
	var v8311 int32
	_ = v8311
	var v8313 int32
	_ = v8313
	var v8314 int32
	_ = v8314
	var v8320 int64
	_ = v8320
	var v8325 int32
	_ = v8325
	var v8329 int32
	_ = v8329
	var v8331 int32
	_ = v8331
	var v8337 int32
	_ = v8337
	var v8340 int32
	_ = v8340
	var v8342 int32
	_ = v8342
	var v8348 int32
	_ = v8348
	var v8352 int32
	_ = v8352
	var v8354 int32
	_ = v8354
	var v8357 int32
	_ = v8357
	var v8361 int32
	_ = v8361
	var v8366 int32
	_ = v8366
	var v8367 int32
	_ = v8367
	var v8368 int32
	_ = v8368
	var v8371 int32
	_ = v8371
	var v8375 int32
	_ = v8375
	var v8380 int32
	_ = v8380
	var v8381 int32
	_ = v8381
	var v8382 int32
	_ = v8382
	var v8385 int32
	_ = v8385
	var v8389 int32
	_ = v8389
	var v8396 int32
	_ = v8396
	var v8399 int32
	_ = v8399
	var v8407 int32
	_ = v8407
	var v8408 int32
	_ = v8408
	var v8412 int32
	_ = v8412
	var v8413 int32
	_ = v8413
	var v8414 int32
	_ = v8414
	var v8419 int64
	_ = v8419
	var v8420 int64
	_ = v8420
	var v8428 int32
	_ = v8428
	var v8429 int32
	_ = v8429
	var v8430 int32
	_ = v8430
	var v8432 int32
	_ = v8432
	var v8433 int32
	_ = v8433
	var v8439 int64
	_ = v8439
	var v8444 int32
	_ = v8444
	var v8448 int32
	_ = v8448
	var v8450 int32
	_ = v8450
	var v8456 int32
	_ = v8456
	var v8459 int32
	_ = v8459
	var v8461 int32
	_ = v8461
	var v8467 int32
	_ = v8467
	var v8471 int32
	_ = v8471
	var v8473 int32
	_ = v8473
	var v8476 int32
	_ = v8476
	var v8480 int32
	_ = v8480
	var v8485 int32
	_ = v8485
	var v8486 int32
	_ = v8486
	var v8487 int32
	_ = v8487
	var v8490 int32
	_ = v8490
	var v8494 int32
	_ = v8494
	var v8501 int32
	_ = v8501
	var v8504 int32
	_ = v8504
	var v8512 int32
	_ = v8512
	var v8513 int32
	_ = v8513
	var v8517 int32
	_ = v8517
	var v8518 int32
	_ = v8518
	var v8519 int32
	_ = v8519
	var v8524 int64
	_ = v8524
	var v8525 int64
	_ = v8525
	var v8533 int32
	_ = v8533
	var v8534 int32
	_ = v8534
	var v8535 int32
	_ = v8535
	var v8537 int32
	_ = v8537
	var v8541 int32
	_ = v8541
	var v8547 int32
	_ = v8547
	var v8552 int32
	_ = v8552
	var v8560 int32
	_ = v8560
	var v8564 int32
	_ = v8564
	var v8565 int32
	_ = v8565
	var v8569 int32
	_ = v8569
	var v8571 int64
	_ = v8571
	var v8572 int32
	_ = v8572
	var v8573 int32
	_ = v8573
	var v8580 int32
	_ = v8580
	var v8585 int32
	_ = v8585
	var v8588 int32
	_ = v8588
	var v8589 int32
	_ = v8589
	var v8593 int32
	_ = v8593
	var v8595 int64
	_ = v8595
	var v8596 int32
	_ = v8596
	var v8597 int32
	_ = v8597
	var v8604 int32
	_ = v8604
	var v8609 int32
	_ = v8609
	var v8610 int32
	_ = v8610
	var v8617 int32
	_ = v8617
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8629 int32
	_ = v8629
	var v8634 int32
	_ = v8634
	var v8636 int32
	_ = v8636
	var v8639 int64
	_ = v8639
	var v8645 int32
	_ = v8645
	var v8653 int32
	_ = v8653
	var v8660 int32
	_ = v8660
	var v8665 int32
	_ = v8665
	var v8670 int32
	_ = v8670
	var v8677 int32
	_ = v8677
	var v8682 int32
	_ = v8682
	var v8686 int32
	_ = v8686
	var v8689 int64
	_ = v8689
	var v8692 int32
	_ = v8692
	var v8695 int64
	_ = v8695
	var v8701 int32
	_ = v8701
	var v8706 int32
	_ = v8706
	var v8710 int32
	_ = v8710
	var v8711 int64
	_ = v8711
	var v8713 int64
	_ = v8713
	var v8714 int64
	_ = v8714
	var v8718 int64
	_ = v8718
	var v8724 int32
	_ = v8724
	var v8729 int32
	_ = v8729
	var v8733 int32
	_ = v8733
	var v8736 int32
	_ = v8736
	var v8737 int32
	_ = v8737
	var v8743 int32
	_ = v8743
	var v8748 int32
	_ = v8748
	var v8752 int32
	_ = v8752
	var v8753 int32
	_ = v8753
	var v8755 int32
	_ = v8755
	var v8757 int64
	_ = v8757
	var v8759 int32
	_ = v8759
	var v8765 int32
	_ = v8765
	var v8770 int32
	_ = v8770
	var v8778 int32
	_ = v8778
	var v8780 int32
	_ = v8780
	var v8783 int32
	_ = v8783
	var v8784 int32
	_ = v8784
	var v8787 int32
	_ = v8787
	var v8788 int32
	_ = v8788
	var v8794 int32
	_ = v8794
	var v8799 int32
	_ = v8799
	var v8801 int64
	_ = v8801
	var v8811 int32
	_ = v8811
	var v8818 int32
	_ = v8818
	var v8819 int32
	_ = v8819
	var v8823 int32
	_ = v8823
	var v8825 int64
	_ = v8825
	var v8826 int32
	_ = v8826
	var v8827 int32
	_ = v8827
	var v8834 int32
	_ = v8834
	var v8839 int32
	_ = v8839
	var v8843 int32
	_ = v8843
	var v8845 int64
	_ = v8845
	var v8846 int32
	_ = v8846
	var v8847 int32
	_ = v8847
	var v8854 int32
	_ = v8854
	var v8859 int32
	_ = v8859
	var v8895 int32
	_ = v8895
	var v8898 int32
	_ = v8898
	var v8900 int32
	_ = v8900
	var v8903 int32
	_ = v8903
	var v8905 int32
	_ = v8905
	var v8906 int32
	_ = v8906
	var v8910 int32
	_ = v8910
	var v8917 int32
	_ = v8917
	var v8919 int32
	_ = v8919
	var v8920 int32
	_ = v8920
	var v8929 int32
	_ = v8929
	var v8937 int32
	_ = v8937
	var v8968 int32
	_ = v8968
	var v9000 int32
	_ = v9000
	var v9003 int32
	_ = v9003
	var v9006 int32
	_ = v9006
	var v9010 int32
	_ = v9010
	var v9012 int32
	_ = v9012
	var v9015 int32
	_ = v9015
	var v9019 int32
	_ = v9019
	var v9022 int32
	_ = v9022
	var v9027 int32
	_ = v9027
	var v9028 int32
	_ = v9028
	var v9030 int32
	_ = v9030
	var v9031 int64
	_ = v9031
	var v9034 int32
	_ = v9034
	var v9035 int32
	_ = v9035
	var v9039 int64
	_ = v9039
	var v9045 int32
	_ = v9045
	var v9050 int32
	_ = v9050
	var v9053 int32
	_ = v9053
	var v9054 int32
	_ = v9054
	var v9058 int32
	_ = v9058
	var v9065 int32
	_ = v9065
	var v9067 int32
	_ = v9067
	var v9070 int64
	_ = v9070
	var v9075 int32
	_ = v9075
	var v9076 int32
	_ = v9076
	var v9079 int32
	_ = v9079
	var v9080 int32
	_ = v9080
	var v9084 int32
	_ = v9084
	var v9089 int32
	_ = v9089
	var v9091 int32
	_ = v9091
	var v9100 int32
	_ = v9100
	var v9128 int32
	_ = v9128
	var v9134 int32
	_ = v9134
	var v9141 int32
	_ = v9141
	var v9145 int32
	_ = v9145
	var v9150 int32
	_ = v9150
	var v9154 int32
	_ = v9154
	var v9157 int32
	_ = v9157
	var v9161 int32
	_ = v9161
	var v9166 int32
	_ = v9166
	var v9202 int32
	_ = v9202
	var v9204 int32
	_ = v9204
	var v9207 int32
	_ = v9207
	var v9208 int32
	_ = v9208
	var v9210 int32
	_ = v9210
	var v9214 int32
	_ = v9214
	var v9215 int32
	_ = v9215
	var v9219 int32
	_ = v9219
	var v9226 int32
	_ = v9226
	var v9228 int32
	_ = v9228
	var v9229 int32
	_ = v9229
	var v9231 int32
	_ = v9231
	var v9236 int32
	_ = v9236
	var v9240 int32
	_ = v9240
	var v9241 int32
	_ = v9241
	var v9277 int32
	_ = v9277
	var v9281 int32
	_ = v9281
	var v9282 int32
	_ = v9282
	var v9288 int32
	_ = v9288
	var v9292 int32
	_ = v9292
	var v9296 int32
	_ = v9296
	var v9298 int32
	_ = v9298
	var v9299 int32
	_ = v9299
	var v9303 int32
	_ = v9303
	var v9310 int32
	_ = v9310
	var v9312 int32
	_ = v9312
	var v9313 int32
	_ = v9313
	var v9319 int32
	_ = v9319
	var v9355 int32
	_ = v9355
	var v9359 int32
	_ = v9359
	var v9363 int32
	_ = v9363
	var v9364 int32
	_ = v9364
	var v9366 int32
	_ = v9366
	var v9370 int32
	_ = v9370
	var v9377 int32
	_ = v9377
	var v9378 int32
	_ = v9378
	var v9380 int32
	_ = v9380
	var v9398 int64
	_ = v9398
	var v9407 int32
	_ = v9407
	var v9408 int32
	_ = v9408
	var v9411 int32
	_ = v9411
	var v9419 int32
	_ = v9419
	var v9420 int32
	_ = v9420
	var v9421 int32
	_ = v9421
	var v9424 int64
	_ = v9424
	var v9425 int64
	_ = v9425
	var v9434 int64
	_ = v9434
	var v9435 int32
	_ = v9435
	var v9442 int32
	_ = v9442
	var v9443 int32
	_ = v9443
	var v9450 int32
	_ = v9450
	var v9452 int32
	_ = v9452
	var v9453 int32
	_ = v9453
	var v9454 int32
	_ = v9454
	var v9455 int64
	_ = v9455
	var v9457 int32
	_ = v9457
	var v9494 int32
	_ = v9494
	var v9498 int32
	_ = v9498
	var v9534 int32
	_ = v9534
	var v9537 int32
	_ = v9537
	var v9542 int32
	_ = v9542
	var v9543 int32
	_ = v9543
	var v9544 int32
	_ = v9544
	var v9546 int32
	_ = v9546
	var v9550 int32
	_ = v9550
	var v9551 int64
	_ = v9551
	var v9553 int32
	_ = v9553
	var v9555 int32
	_ = v9555
	var v9558 int32
	_ = v9558
	var v9559 int32
	_ = v9559
	var v9561 int32
	_ = v9561
	var v9562 int64
	_ = v9562
	var v9563 int32
	_ = v9563
	var v9566 int32
	_ = v9566
	var v9570 int32
	_ = v9570
	var v9573 int32
	_ = v9573
	var v9576 int32
	_ = v9576
	var v9582 int64
	_ = v9582
	var v9585 int32
	_ = v9585
	var v9586 int32
	_ = v9586
	var v9587 int32
	_ = v9587
	var v9589 int32
	_ = v9589
	var v9590 int32
	_ = v9590
	var v9595 int32
	_ = v9595
	var v9596 int64
	_ = v9596
	var v9600 int32
	_ = v9600
	var v9604 int32
	_ = v9604
	var v9609 int32
	_ = v9609
	var v9610 int32
	_ = v9610
	var v9616 int32
	_ = v9616
	var v9617 int32
	_ = v9617
	var v9619 int32
	_ = v9619
	var v9621 int64
	_ = v9621
	var v9622 int32
	_ = v9622
	var v9623 int32
	_ = v9623
	var v9627 int32
	_ = v9627
	var v9635 int32
	_ = v9635
	var v9636 int32
	_ = v9636
	var v9638 int64
	_ = v9638
	var v9643 int32
	_ = v9643
	var v9644 int32
	_ = v9644
	var v9647 int64
	_ = v9647
	var v9655 int32
	_ = v9655
	var v9656 int32
	_ = v9656
	var v9665 int32
	_ = v9665
	var v9666 int32
	_ = v9666
	var v9672 int32
	_ = v9672
	var v9673 int32
	_ = v9673
	var v9679 int32
	_ = v9679
	var v9680 int32
	_ = v9680
	var v9685 int32
	_ = v9685
	var v9686 int32
	_ = v9686
	var v9692 int64
	_ = v9692
	var v9695 int64
	_ = v9695
	var v9698 int32
	_ = v9698
	var v9701 int32
	_ = v9701
	var v9708 int32
	_ = v9708
	var v9711 int32
	_ = v9711
	var v9715 int32
	_ = v9715
	var v9717 int64
	_ = v9717
	var v9719 int64
	_ = v9719
	var v9723 int32
	_ = v9723
	var v9726 int32
	_ = v9726
	var v9729 int64
	_ = v9729
	var v9732 int32
	_ = v9732
	var v9738 int32
	_ = v9738
	var v9741 int32
	_ = v9741
	var v9745 int32
	_ = v9745
	var v9750 int32
	_ = v9750
	var v9753 int32
	_ = v9753
	var v9755 int32
	_ = v9755
	var v9757 int32
	_ = v9757
	var v9758 int32
	_ = v9758
	var v9760 int32
	_ = v9760
	var v9764 int32
	_ = v9764
	var v9765 int32
	_ = v9765
	var v9767 int32
	_ = v9767
	var v9768 int32
	_ = v9768
	var v9771 int32
	_ = v9771
	var v9775 int32
	_ = v9775
	var v9777 int32
	_ = v9777
	var v9780 int32
	_ = v9780
	var v9782 int32
	_ = v9782
	var v9783 int32
	_ = v9783
	var v9784 int32
	_ = v9784
	var v9786 int32
	_ = v9786
	var v9789 int32
	_ = v9789
	var v9790 int32
	_ = v9790
	var v9796 int32
	_ = v9796
	var v9801 int32
	_ = v9801
	var v9805 int32
	_ = v9805
	var v9809 int32
	_ = v9809
	var v9810 int64
	_ = v9810
	var v9811 int64
	_ = v9811
	var v9812 int64
	_ = v9812
	var v9817 int64
	_ = v9817
	var v9818 int64
	_ = v9818
	var v9821 int64
	_ = v9821
	var v9824 int32
	_ = v9824
	var v9829 int32
	_ = v9829
	var v9830 int32
	_ = v9830
	var v9832 int32
	_ = v9832
	var v9833 int32
	_ = v9833
	var v9839 int32
	_ = v9839
	var v9844 int32
	_ = v9844
	var v9845 int32
	_ = v9845
	var v9846 int32
	_ = v9846
	var v9849 int32
	_ = v9849
	var v9850 int32
	_ = v9850
	var v9854 int32
	_ = v9854
	var v9861 int32
	_ = v9861
	var v9895 int32
	_ = v9895
	var v9906 int32
	_ = v9906
	var v9911 int32
	_ = v9911
	var v9914 int32
	_ = v9914
	var v9915 int32
	_ = v9915
	var v9920 int32
	_ = v9920
	var v9925 int32
	_ = v9925
	var v9935 int32
	_ = v9935
	var v9940 int32
	_ = v9940
	var v9942 int32
	_ = v9942
	var v9951 int32
	_ = v9951
	var v9956 int32
	_ = v9956
	var v9957 int32
	_ = v9957
	var v9961 int32
	_ = v9961
	var v9965 int32
	_ = v9965
	var v9967 int32
	_ = v9967
	var v10004 int32
	_ = v10004
	var v10009 int32
	_ = v10009
	var v10014 int32
	_ = v10014
	var v10018 int32
	_ = v10018
	var v10023 int32
	_ = v10023
	var v10029 int32
	_ = v10029
	var v10030 int32
	_ = v10030
	var v10032 int32
	_ = v10032
	var v10033 int32
	_ = v10033
	var v10037 int32
	_ = v10037
	var v10045 int32
	_ = v10045
	var v10050 int32
	_ = v10050
	var v10052 int32
	_ = v10052
	var v10055 int32
	_ = v10055
	var v10056 int32
	_ = v10056
	var v10057 int32
	_ = v10057
	var v10058 int32
	_ = v10058
	var v10065 int32
	_ = v10065
	var v10066 int32
	_ = v10066
	var v10070 int32
	_ = v10070
	var v10074 int32
	_ = v10074
	var v10079 int32
	_ = v10079
	var v10080 int32
	_ = v10080
	var v10081 int32
	_ = v10081
	var v10082 int32
	_ = v10082
	var v10120 int64
	_ = v10120
	var v10121 int64
	_ = v10121
	var v10122 int64
	_ = v10122
	var v10125 int64
	_ = v10125
	var v10128 int32
	_ = v10128
	var v10133 int32
	_ = v10133
	var v10134 int32
	_ = v10134
	var v10136 int32
	_ = v10136
	var v10137 int32
	_ = v10137
	var v10142 int32
	_ = v10142
	var v10143 int32
	_ = v10143
	var v10144 int32
	_ = v10144
	var v10149 int32
	_ = v10149
	var v10150 int32
	_ = v10150
	var v10152 int32
	_ = v10152
	var v10153 int32
	_ = v10153
	var v10154 int32
	_ = v10154
	var v10156 int32
	_ = v10156
	var v10161 int32
	_ = v10161
	var v10166 int32
	_ = v10166
	var v10167 int32
	_ = v10167
	var v10168 int32
	_ = v10168
	var v10170 int32
	_ = v10170
	var v10171 int32
	_ = v10171
	var v10175 int32
	_ = v10175
	var v10180 int32
	_ = v10180
	var v10185 int32
	_ = v10185
	var v10186 int32
	_ = v10186
	var v10192 int32
	_ = v10192
	var v10193 int32
	_ = v10193
	var v10201 int32
	_ = v10201
	var v10202 int32
	_ = v10202
	var v10207 int32
	_ = v10207
	var v10208 int32
	_ = v10208
	var v10212 int32
	_ = v10212
	var v10214 int32
	_ = v10214
	var v10215 int32
	_ = v10215
	var v10221 int32
	_ = v10221
	var v10223 int32
	_ = v10223
	var v10231 int32
	_ = v10231
	var v10263 int32
	_ = v10263
	var v10268 int32
	_ = v10268
	var v10272 int32
	_ = v10272
	var v10273 int32
	_ = v10273
	var v10275 int32
	_ = v10275
	var v10276 int32
	_ = v10276
	var v10277 int32
	_ = v10277
	var v10283 int32
	_ = v10283
	var v10287 int32
	_ = v10287
	var v10289 int32
	_ = v10289
	var v10294 int32
	_ = v10294
	var v10295 int32
	_ = v10295
	var v10300 int32
	_ = v10300
	var v10302 int32
	_ = v10302
	var v10308 int32
	_ = v10308
	var v10313 int32
	_ = v10313
	var v10314 int32
	_ = v10314
	var v10315 int32
	_ = v10315
	var v10317 int32
	_ = v10317
	var v10318 int32
	_ = v10318
	var v10321 int32
	_ = v10321
	var v10326 int32
	_ = v10326
	var v10328 int32
	_ = v10328
	var v10334 int32
	_ = v10334
	var v10339 int32
	_ = v10339
	var v10343 int32
	_ = v10343
	var v10345 int32
	_ = v10345
	var v10353 int32
	_ = v10353
	var v10358 int32
	_ = v10358
	var v10360 int32
	_ = v10360
	var v10367 int32
	_ = v10367
	var v10369 int32
	_ = v10369
	var v10377 int32
	_ = v10377
	var v10382 int32
	_ = v10382
	var v10386 int32
	_ = v10386
	var v10422 int64
	_ = v10422
	var v10425 int32
	_ = v10425
	var v10430 int32
	_ = v10430
	var v10431 int32
	_ = v10431
	var v10432 int32
	_ = v10432
	var v10437 int32
	_ = v10437
	var v10440 int32
	_ = v10440
	var v10442 int32
	_ = v10442
	var v10443 int32
	_ = v10443
	var v10444 int32
	_ = v10444
	var v10447 int32
	_ = v10447
	var v10452 int32
	_ = v10452
	var v10457 int32
	_ = v10457
	var v10461 int32
	_ = v10461
	var v10466 int32
	_ = v10466
	var v10472 int32
	_ = v10472
	var v10473 int32
	_ = v10473
	var v10475 int32
	_ = v10475
	var v10476 int32
	_ = v10476
	var v10480 int32
	_ = v10480
	var v10488 int32
	_ = v10488
	var v10493 int32
	_ = v10493
	var v10495 int32
	_ = v10495
	var v10498 int32
	_ = v10498
	var v10499 int32
	_ = v10499
	var v10502 int32
	_ = v10502
	var v10507 int32
	_ = v10507
	var v10508 int32
	_ = v10508
	var v10512 int32
	_ = v10512
	var v10513 int32
	_ = v10513
	var v10515 int32
	_ = v10515
	var v10520 int32
	_ = v10520
	var v10525 int32
	_ = v10525
	var v10526 int32
	_ = v10526
	var v10528 int32
	_ = v10528
	var v10533 int32
	_ = v10533
	var v10534 int32
	_ = v10534
	var v10536 int32
	_ = v10536
	var v10537 int32
	_ = v10537
	var v10540 int32
	_ = v10540
	var v10545 int32
	_ = v10545
	var v10547 int32
	_ = v10547
	var v10553 int32
	_ = v10553
	var v10558 int32
	_ = v10558
	var v10562 int32
	_ = v10562
	var v10564 int32
	_ = v10564
	var v10572 int32
	_ = v10572
	var v10577 int32
	_ = v10577
	var v10615 int32
	_ = v10615
	var v10617 int32
	_ = v10617
	var v10625 int32
	_ = v10625
	var v10630 int32
	_ = v10630
	var v10633 int32
	_ = v10633
	var v10634 int32
	_ = v10634
	var v10640 int32
	_ = v10640
	var v10645 int32
	_ = v10645
	var v10651 int32
	_ = v10651
	var v10681 int32
	_ = v10681
	var v10682 int32
	_ = v10682
	var v10686 int32
	_ = v10686
	var v10693 int32
	_ = v10693
	var v10695 int32
	_ = v10695
	var v10697 int32
	_ = v10697
	var v10703 int64
	_ = v10703
	var v10704 int64
	_ = v10704
	var v10707 int32
	_ = v10707
	var v10709 int32
	_ = v10709
	var v10710 int64
	_ = v10710
	var v10711 int64
	_ = v10711
	var v10714 int64
	_ = v10714
	var v10715 int64
	_ = v10715
	var v10721 int32
	_ = v10721
	var v10723 int64
	_ = v10723
	var v10731 int32
	_ = v10731
	var v10744 int64
	_ = v10744
	var v10751 int32
	_ = v10751
	var v10753 int64
	_ = v10753
	var v10755 int64
	_ = v10755
	var v10758 int32
	_ = v10758
	var v10759 int64
	_ = v10759
	var v10765 int64
	_ = v10765
	var v10784 int64
	_ = v10784
	var v10792 int64
	_ = v10792
	var v10798 int32
	_ = v10798
	var v10801 int32
	_ = v10801
	var v10805 int64
	_ = v10805
	var v10806 int32
	_ = v10806
	var v10809 int32
	_ = v10809
	var v10810 int64
	_ = v10810
	var v10812 int32
	_ = v10812
	var v10813 int32
	_ = v10813
	var v10816 int32
	_ = v10816
	var v10820 int32
	_ = v10820
	var v10824 int64
	_ = v10824
	var v10825 int64
	_ = v10825
	var v10828 int64
	_ = v10828
	var v10832 int32
	_ = v10832
	var v10833 int32
	_ = v10833
	var v10837 int64
	_ = v10837
	var v10844 int32
	_ = v10844
	var v10845 int32
	_ = v10845
	var v10848 int32
	_ = v10848
	var v10851 int32
	_ = v10851
	var v10854 int32
	_ = v10854
	var v10858 int64
	_ = v10858
	var v10860 int32
	_ = v10860
	var v10872 int64
	_ = v10872
	var v10879 int32
	_ = v10879
	var v10880 int32
	_ = v10880
	var v10883 int32
	_ = v10883
	var v10884 int32
	_ = v10884
	var v10887 int32
	_ = v10887
	var v10889 int32
	_ = v10889
	var v10896 int32
	_ = v10896
	var v10898 int64
	_ = v10898
	var v10900 int32
	_ = v10900
	var v10904 int32
	_ = v10904
	var v10908 int32
	_ = v10908
	var v10909 int32
	_ = v10909
	var v10911 int32
	_ = v10911
	var v10912 int64
	_ = v10912
	var v10914 int64
	_ = v10914
	var v10951 int64
	_ = v10951
	var v10959 int64
	_ = v10959
	var v10999 int32
	_ = v10999
	var v11003 int32
	_ = v11003
	var v11005 int32
	_ = v11005
	var v11009 int32
	_ = v11009
	var v11011 int32
	_ = v11011
	var v11012 int32
	_ = v11012
	var v11014 int32
	_ = v11014
	var v11015 int64
	_ = v11015
	var v11019 int64
	_ = v11019
	var v11022 int32
	_ = v11022
	var v11023 int32
	_ = v11023
	var v11026 int32
	_ = v11026
	var v11028 int32
	_ = v11028
	var v11029 int32
	_ = v11029
	var v11030 int32
	_ = v11030
	var v11032 int32
	_ = v11032
	var v11035 int32
	_ = v11035
	var v11036 int32
	_ = v11036
	var v11038 int32
	_ = v11038
	var v11039 int32
	_ = v11039
	var v11040 int32
	_ = v11040
	var v11043 int32
	_ = v11043
	var v11045 int32
	_ = v11045
	var v11046 int32
	_ = v11046
	var v11047 int32
	_ = v11047
	var v11048 int32
	_ = v11048
	var v11049 int32
	_ = v11049
	var v11056 int32
	_ = v11056
	var v11059 int32
	_ = v11059
	var v11060 int32
	_ = v11060
	var v11063 int32
	_ = v11063
	var v11077 int32
	_ = v11077
	var v11079 int32
	_ = v11079
	var v11081 int32
	_ = v11081
	var v11088 int32
	_ = v11088
	var v11101 int32
	_ = v11101
	var v11102 int32
	_ = v11102
	var v11104 int32
	_ = v11104
	var v11113 int32
	_ = v11113
	var v11115 int32
	_ = v11115
	var v11119 int32
	_ = v11119
	var v11120 int32
	_ = v11120
	var v11122 int32
	_ = v11122
	var v11123 int32
	_ = v11123
	var v11124 int32
	_ = v11124
	var v11125 int32
	_ = v11125
	var v11126 int32
	_ = v11126
	var v11128 int32
	_ = v11128
	var v11132 int32
	_ = v11132
	var v11133 int32
	_ = v11133
	var v11134 int32
	_ = v11134
	var v11136 int32
	_ = v11136
	var v11137 int64
	_ = v11137
	var v11140 int32
	_ = v11140
	var v11141 int32
	_ = v11141
	var v11143 int32
	_ = v11143
	var v11144 int32
	_ = v11144
	var v11147 int32
	_ = v11147
	var v11149 int32
	_ = v11149
	var v11150 int32
	_ = v11150
	var v11152 int32
	_ = v11152
	var v11156 int32
	_ = v11156
	var v11157 int32
	_ = v11157
	var v11160 int32
	_ = v11160
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
	var v11169 int32
	_ = v11169
	var v11171 int32
	_ = v11171
	var v11172 int32
	_ = v11172
	var v11179 int32
	_ = v11179
	var v11188 int32
	_ = v11188
	var v11190 int32
	_ = v11190
	var v11192 int32
	_ = v11192
	var v11199 int32
	_ = v11199
	var v11205 int32
	_ = v11205
	var v11215 int32
	_ = v11215
	var v11216 int32
	_ = v11216
	var v11218 int32
	_ = v11218
	var v11221 int32
	_ = v11221
	var v11223 int32
	_ = v11223
	var v11225 int32
	_ = v11225
	var v11226 int64
	_ = v11226
	var v11229 int32
	_ = v11229
	var v11231 int32
	_ = v11231
	var v11233 int32
	_ = v11233
	var v11234 int32
	_ = v11234
	var v11236 int32
	_ = v11236
	var v11237 int32
	_ = v11237
	var v11240 int32
	_ = v11240
	var v11242 int32
	_ = v11242
	var v11243 int32
	_ = v11243
	var v11246 int32
	_ = v11246
	var v11247 int32
	_ = v11247
	var v11248 int32
	_ = v11248
	var v11251 int32
	_ = v11251
	var v11256 int32
	_ = v11256
	var v11258 int32
	_ = v11258
	var v11259 int32
	_ = v11259
	var v11263 int32
	_ = v11263
	var v11264 int32
	_ = v11264
	var v11283 int32
	_ = v11283
	var v11289 int32
	_ = v11289
	var v11296 int32
	_ = v11296
	var v11297 int32
	_ = v11297
	var v11299 int32
	_ = v11299
	var v11302 int32
	_ = v11302
	var v11309 int32
	_ = v11309
	var v11313 int32
	_ = v11313
	var v11314 int32
	_ = v11314
	var v11316 int32
	_ = v11316
	var v11317 int32
	_ = v11317
	var v11320 int32
	_ = v11320
	var v11324 int32
	_ = v11324
	var v11327 int32
	_ = v11327
	var v11328 int32
	_ = v11328
	var v11329 int32
	_ = v11329
	var v11331 int32
	_ = v11331
	var v11334 int32
	_ = v11334
	var v11338 int32
	_ = v11338
	var v11339 int32
	_ = v11339
	var v11341 int32
	_ = v11341
	var v11342 int32
	_ = v11342
	var v11351 int32
	_ = v11351
	var v11357 int32
	_ = v11357
	var v11382 int32
	_ = v11382
	var v11383 int32
	_ = v11383
	var v11384 int64
	_ = v11384
	var v11385 int32
	_ = v11385
	var v11388 int32
	_ = v11388
	var v11389 int32
	_ = v11389
	var v11392 int32
	_ = v11392
	var v11393 int32
	_ = v11393
	var v11397 int32
	_ = v11397
	var v11402 int32
	_ = v11402
	var v11403 int32
	_ = v11403
	var v11404 int32
	_ = v11404
	var v11405 int32
	_ = v11405
	var v11406 int32
	_ = v11406
	var v11407 int32
	_ = v11407
	var v11408 int32
	_ = v11408
	var v11409 int32
	_ = v11409
	var v11411 int32
	_ = v11411
	var v11412 int64
	_ = v11412
	var v11413 int32
	_ = v11413
	var v11414 int32
	_ = v11414
	var v11416 int32
	_ = v11416
	var v11418 int32
	_ = v11418
	var v11422 int32
	_ = v11422
	var v11427 int32
	_ = v11427
	var v11430 int32
	_ = v11430
	var v11444 int32
	_ = v11444
	var v11454 int32
	_ = v11454
	var v11455 int32
	_ = v11455
	var v11456 int32
	_ = v11456
	var v11459 int32
	_ = v11459
	var v11460 int32
	_ = v11460
	var v11463 int32
	_ = v11463
	var v11468 int32
	_ = v11468
	var v11470 int32
	_ = v11470
	var v11475 int32
	_ = v11475
	var v11477 int32
	_ = v11477
	var v11479 int32
	_ = v11479
	var v11480 int32
	_ = v11480
	var v11481 int32
	_ = v11481
	var v11483 int32
	_ = v11483
	var v11488 int32
	_ = v11488
	var v11490 int32
	_ = v11490
	var v11494 int32
	_ = v11494
	var v11501 int32
	_ = v11501
	var v11529 int32
	_ = v11529
	var v11531 int32
	_ = v11531
	var v11534 int32
	_ = v11534
	var v11535 int32
	_ = v11535
	var v11536 int32
	_ = v11536
	var v11538 int32
	_ = v11538
	var v11539 int32
	_ = v11539
	var v11546 int32
	_ = v11546
	var v11549 int32
	_ = v11549
	var v11551 int32
	_ = v11551
	var v11553 int32
	_ = v11553
	var v11557 int32
	_ = v11557
	var v11558 int32
	_ = v11558
	var v11560 int32
	_ = v11560
	var v11564 int32
	_ = v11564
	var v11568 int32
	_ = v11568
	var v11573 int32
	_ = v11573
	var v11575 int32
	_ = v11575
	var v11579 int32
	_ = v11579
	var v11580 int32
	_ = v11580
	var v11616 int32
	_ = v11616
	var v11618 int32
	_ = v11618
	var v11619 int32
	_ = v11619
	var v11656 int32
	_ = v11656
	var v11660 int32
	_ = v11660
	var v11664 int32
	_ = v11664
	var v11666 int32
	_ = v11666
	var v11669 int32
	_ = v11669
	var v11674 int32
	_ = v11674
	var v11675 int32
	_ = v11675
	var v11676 int64
	_ = v11676
	var v11677 int32
	_ = v11677
	var v11678 int64
	_ = v11678
	var v11681 int32
	_ = v11681
	var v11682 int32
	_ = v11682
	var v11683 int32
	_ = v11683
	var v11685 int32
	_ = v11685
	var v11686 int32
	_ = v11686
	var v11691 int32
	_ = v11691
	var v11692 int64
	_ = v11692
	var v11697 int32
	_ = v11697
	var v11700 int32
	_ = v11700
	var v11705 int32
	_ = v11705
	var v11707 int32
	_ = v11707
	var v11709 int32
	_ = v11709
	var v11710 int32
	_ = v11710
	var v11712 int32
	_ = v11712
	var v11713 int32
	_ = v11713
	var v11715 int32
	_ = v11715
	var v11717 int32
	_ = v11717
	var v11719 int32
	_ = v11719
	var v11725 int32
	_ = v11725
	var v11726 int32
	_ = v11726
	var v11727 int32
	_ = v11727
	var v11731 int32
	_ = v11731
	var v11732 int32
	_ = v11732
	var v11733 int32
	_ = v11733
	var v11735 int32
	_ = v11735
	var v11739 int32
	_ = v11739
	var v11753 int32
	_ = v11753
	var v11758 int32
	_ = v11758
	var v11759 int32
	_ = v11759
	var v11761 int32
	_ = v11761
	var v11769 int32
	_ = v11769
	var v11771 int32
	_ = v11771
	var v11778 int32
	_ = v11778
	var v11787 int64
	_ = v11787
	var v11788 int64
	_ = v11788
	var v11791 int64
	_ = v11791
	var v11793 int64
	_ = v11793
	var v11794 int64
	_ = v11794
	var v11796 int64
	_ = v11796
	var v11802 int64
	_ = v11802
	var v11803 int64
	_ = v11803
	var v11804 int64
	_ = v11804
	var v11814 int64
	_ = v11814
	var v11816 int64
	_ = v11816
	var v11820 int64
	_ = v11820
	var v11822 int32
	_ = v11822
	var v11824 int32
	_ = v11824
	var v11829 int32
	_ = v11829
	var v11834 int32
	_ = v11834
	var v11836 int32
	_ = v11836
	var v11838 int32
	_ = v11838
	var v11842 int32
	_ = v11842
	var v11847 int32
	_ = v11847
	var v11848 int32
	_ = v11848
	var v11851 int32
	_ = v11851
	var v11853 int32
	_ = v11853
	var v11857 int32
	_ = v11857
	var v11859 int32
	_ = v11859
	var v11860 int32
	_ = v11860
	var v11861 int32
	_ = v11861
	var v11863 int32
	_ = v11863
	var v11866 int32
	_ = v11866
	var v11868 int32
	_ = v11868
	var v11873 int32
	_ = v11873
	var v11874 int32
	_ = v11874
	var v11875 int32
	_ = v11875
	var v11878 int64
	_ = v11878
	var v11879 int64
	_ = v11879
	var v11893 int32
	_ = v11893
	var v11896 int64
	_ = v11896
	var v11897 int32
	_ = v11897
	var v11899 int64
	_ = v11899
	var v11902 int32
	_ = v11902
	var v11903 int32
	_ = v11903
	var v11905 int32
	_ = v11905
	var v11916 int32
	_ = v11916
	var v11919 int32
	_ = v11919
	var v11920 int32
	_ = v11920
	var v11924 int32
	_ = v11924
	var v11928 int32
	_ = v11928
	var v11933 int32
	_ = v11933
	var v11934 int32
	_ = v11934
	var v11938 int32
	_ = v11938
	var v11943 int32
	_ = v11943
	var v11944 int32
	_ = v11944
	var v11946 int32
	_ = v11946
	var v11953 int32
	_ = v11953
	var v11954 int32
	_ = v11954
	var v11955 int32
	_ = v11955
	var v11958 int64
	_ = v11958
	var v11959 int64
	_ = v11959
	var v11970 int32
	_ = v11970
	var v11973 int32
	_ = v11973
	var v11975 int32
	_ = v11975
	var v11976 int32
	_ = v11976
	var v11978 int32
	_ = v11978
	var v11981 int32
	_ = v11981
	var v11982 int32
	_ = v11982
	var v11983 int32
	_ = v11983
	var v11985 int32
	_ = v11985
	var v11990 int32
	_ = v11990
	var v11995 int32
	_ = v11995
	var v11998 int64
	_ = v11998
	var v11999 int32
	_ = v11999
	var v12001 int32
	_ = v12001
	var v12003 int32
	_ = v12003
	var v12007 int32
	_ = v12007
	var v12008 int32
	_ = v12008
	var v12010 int32
	_ = v12010
	var v12012 int32
	_ = v12012
	var v12015 int32
	_ = v12015
	var v12017 int32
	_ = v12017
	var v12019 int32
	_ = v12019
	var v12023 int32
	_ = v12023
	var v12024 int32
	_ = v12024
	var v12026 int32
	_ = v12026
	var v12032 int32
	_ = v12032
	var v12033 int32
	_ = v12033
	var v12037 int32
	_ = v12037
	var v12039 int32
	_ = v12039
	var v12040 int32
	_ = v12040
	var v12043 int32
	_ = v12043
	var v12044 int32
	_ = v12044
	var v12047 int32
	_ = v12047
	var v12048 int32
	_ = v12048
	var v12051 int32
	_ = v12051
	var v12052 int32
	_ = v12052
	var v12055 int32
	_ = v12055
	var v12056 int32
	_ = v12056
	var v12059 int32
	_ = v12059
	var v12060 int32
	_ = v12060
	var v12063 int32
	_ = v12063
	var v12064 int32
	_ = v12064
	var v12067 int32
	_ = v12067
	var v12068 int32
	_ = v12068
	var v12071 int32
	_ = v12071
	var v12078 int32
	_ = v12078
	var v12081 int32
	_ = v12081
	var v12084 int32
	_ = v12084
	var v12087 int32
	_ = v12087
	var v12090 int32
	_ = v12090
	var v12093 int32
	_ = v12093
	var v12096 int32
	_ = v12096
	var v12099 int32
	_ = v12099
	var v12104 int32
	_ = v12104
	var v12107 int64
	_ = v12107
	var v12108 int32
	_ = v12108
	var v12110 int32
	_ = v12110
	var v12112 int32
	_ = v12112
	var v12116 int32
	_ = v12116
	var v12117 int32
	_ = v12117
	var v12119 int32
	_ = v12119
	var v12121 int32
	_ = v12121
	var v12124 int32
	_ = v12124
	var v12127 int32
	_ = v12127
	var v12130 int32
	_ = v12130
	var v12133 int32
	_ = v12133
	var v12136 int32
	_ = v12136
	var v12139 int32
	_ = v12139
	var v12142 int32
	_ = v12142
	var v12145 int32
	_ = v12145
	var v12147 int32
	_ = v12147
	var v12149 int32
	_ = v12149
	var v12153 int32
	_ = v12153
	var v12156 int32
	_ = v12156
	var v12160 int32
	_ = v12160
	var v12163 int32
	_ = v12163
	var v12170 int32
	_ = v12170
	var v12172 int32
	_ = v12172
	var v12174 int32
	_ = v12174
	var v12182 int32
	_ = v12182
	var v12188 int64
	_ = v12188
	var v12189 int64
	_ = v12189
	var v12191 int64
	_ = v12191
	var v12192 int64
	_ = v12192
	var v12195 int64
	_ = v12195
	var v12198 int32
	_ = v12198
	var v12203 int32
	_ = v12203
	var v12204 int32
	_ = v12204
	var v12205 int32
	_ = v12205
	var v12207 int32
	_ = v12207
	var v12213 int32
	_ = v12213
	var v12218 int32
	_ = v12218
	var v12219 int32
	_ = v12219
	var v12220 int32
	_ = v12220
	var v12222 int32
	_ = v12222
	var v12225 int32
	_ = v12225
	var v12235 int32
	_ = v12235
	var v12236 int32
	_ = v12236
	var v12239 int32
	_ = v12239
	var v12247 int32
	_ = v12247
	var v12248 int32
	_ = v12248
	var v12251 int32
	_ = v12251
	var v12254 int32
	_ = v12254
	var v12259 int32
	_ = v12259
	var v12263 int32
	_ = v12263
	var v12267 int64
	_ = v12267
	var v12268 int64
	_ = v12268
	var v12269 int64
	_ = v12269
	var v12272 int64
	_ = v12272
	var v12275 int32
	_ = v12275
	var v12280 int32
	_ = v12280
	var v12281 int32
	_ = v12281
	var v12286 int32
	_ = v12286
	var v12291 int32
	_ = v12291
	var v12292 int32
	_ = v12292
	var v12295 int32
	_ = v12295
	var v12300 int32
	_ = v12300
	var v12301 int32
	_ = v12301
	var v12303 int32
	_ = v12303
	var v12305 int32
	_ = v12305
	var v12306 int32
	_ = v12306
	var v12308 int32
	_ = v12308
	var v12319 int32
	_ = v12319
	var v12323 int32
	_ = v12323
	var v12327 int32
	_ = v12327
	var v12328 int32
	_ = v12328
	var v12330 int32
	_ = v12330
	var v12331 int32
	_ = v12331
	var v12340 int32
	_ = v12340
	var v12346 int32
	_ = v12346
	var v12347 int32
	_ = v12347
	var v12349 int32
	_ = v12349
	var v12353 int32
	_ = v12353
	var v12355 int32
	_ = v12355
	var v12358 int32
	_ = v12358
	var v12362 int32
	_ = v12362
	var v12363 int32
	_ = v12363
	var v12365 int32
	_ = v12365
	var v12369 int32
	_ = v12369
	var v12370 int32
	_ = v12370
	var v12374 int32
	_ = v12374
	var v12381 int32
	_ = v12381
	var v12383 int32
	_ = v12383
	var v12389 int32
	_ = v12389
	var v12391 int32
	_ = v12391
	var v12393 int32
	_ = v12393
	var v12395 int32
	_ = v12395
	var v12399 int32
	_ = v12399
	var v12401 int32
	_ = v12401
	var v12403 int32
	_ = v12403
	var v12404 int32
	_ = v12404
	var v12407 int32
	_ = v12407
	var v12410 int32
	_ = v12410
	var v12415 int32
	_ = v12415
	var v12418 int32
	_ = v12418
	var v12422 int32
	_ = v12422
	var v12427 int32
	_ = v12427
	var v12431 int32
	_ = v12431
	var v12434 int32
	_ = v12434
	var v12438 int32
	_ = v12438
	var v12443 int32
	_ = v12443
	var v12447 int32
	_ = v12447
	var v12449 int32
	_ = v12449
	var v12456 int32
	_ = v12456
	var v12461 int32
	_ = v12461
	var v12465 int32
	_ = v12465
	var v12467 int32
	_ = v12467
	var v12475 int32
	_ = v12475
	var v12480 int32
	_ = v12480
	var v12484 int32
	_ = v12484
	var v12492 int32
	_ = v12492
	var v12497 int32
	_ = v12497
	var v12501 int32
	_ = v12501
	var v12504 int32
	_ = v12504
	var v12508 int32
	_ = v12508
	var v12512 int32
	_ = v12512
	var v12517 int32
	_ = v12517
	var v12521 int32
	_ = v12521
	var v12523 int32
	_ = v12523
	var v12531 int32
	_ = v12531
	var v12536 int32
	_ = v12536
	var v12540 int32
	_ = v12540
	var v12542 int32
	_ = v12542
	var v12550 int32
	_ = v12550
	var v12555 int32
	_ = v12555
	var v12557 int32
	_ = v12557
	var v12565 int32
	_ = v12565
	var v12570 int32
	_ = v12570
	var v12571 int32
	_ = v12571
	var v12572 int32
	_ = v12572
	var v12574 int32
	_ = v12574
	var v12575 int32
	_ = v12575
	var v12578 int32
	_ = v12578
	var v12583 int32
	_ = v12583
	var v12585 int32
	_ = v12585
	var v12591 int32
	_ = v12591
	var v12596 int32
	_ = v12596
	var v12600 int32
	_ = v12600
	var v12602 int32
	_ = v12602
	var v12610 int32
	_ = v12610
	var v12615 int32
	_ = v12615
	var v12619 int32
	_ = v12619
	var v12621 int32
	_ = v12621
	var v12629 int32
	_ = v12629
	var v12634 int32
	_ = v12634
	var v12636 int32
	_ = v12636
	var v12638 int32
	_ = v12638
	var v12640 int32
	_ = v12640
	var v12642 int32
	_ = v12642
	var v12648 int32
	_ = v12648
	var v12650 int32
	_ = v12650
	var v12656 int32
	_ = v12656
	var v12661 int32
	_ = v12661
	var v12667 int32
	_ = v12667
	var v12671 int32
	_ = v12671
	var v12676 int32
	_ = v12676
	var v12680 int32
	_ = v12680
	var v12683 int64
	_ = v12683
	var v12689 int32
	_ = v12689
	var v12694 int32
	_ = v12694
	var v12698 int32
	_ = v12698
	var v12701 int64
	_ = v12701
	var v12707 int32
	_ = v12707
	var v12712 int32
	_ = v12712
	var v12716 int32
	_ = v12716
	var v12718 int64
	_ = v12718
	var v12721 int64
	_ = v12721
	var v12727 int32
	_ = v12727
	var v12732 int32
	_ = v12732
	var v12737 int32
	_ = v12737
	var v12741 int32
	_ = v12741
	var v12746 int32
	_ = v12746
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
	v12737 = m.ExcPending
	if v12737 != 0 {
		goto L32
	} else {
		goto L2672
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12716 = m.ExcPending
	if v12716 != 0 {
		goto L32
	} else {
		goto L2669
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12698 = m.ExcPending
	if v12698 != 0 {
		goto L32
	} else {
		goto L2666
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12680 = m.ExcPending
	if v12680 != 0 {
		goto L32
	} else {
		goto L2663
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12667 = m.ExcPending
	if v12667 != 0 {
		goto L32
	} else {
		goto L2660
	}
L6:
	;
	v12636 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v12638 = v39 + int32(_a_F_StartupXLOG_1)
	v12640 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	F_XLogFileName(m, v12638, v9786, v9812, v12640)
	mBase = m.M
	v12642 = m.ExcPending
	if v12642 != 0 {
		goto L32
	} else {
		goto L2655
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12619 = m.ExcPending
	if v12619 != 0 {
		goto L32
	} else {
		goto L2651
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12600 = m.ExcPending
	if v12600 != 0 {
		goto L32
	} else {
		goto L2647
	}
L9:
	;
	v12571 = int32(_a_F_StartupXLOG_2)
	v12572 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v12574 = v39 + int32(_a_F_StartupXLOG_3)
	v12575 = F_unlink(m, v12574)
	mBase = m.M
	if v12572 != 0 {
		goto L2640
	} else {
		goto L2641
	}
L10:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12557 = m.ExcPending
	if v12557 != 0 {
		goto L32
	} else {
		goto L2637
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12540 = m.ExcPending
	if v12540 != 0 {
		goto L32
	} else {
		goto L2633
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12521 = m.ExcPending
	if v12521 != 0 {
		goto L32
	} else {
		goto L2629
	}
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12501 = m.ExcPending
	if v12501 != 0 {
		goto L32
	} else {
		goto L2624
	}
L14:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12484 = m.ExcPending
	if v12484 != 0 {
		goto L32
	} else {
		goto L2621
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12465 = m.ExcPending
	if v12465 != 0 {
		goto L32
	} else {
		goto L2617
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12447 = m.ExcPending
	if v12447 != 0 {
		goto L32
	} else {
		goto L2613
	}
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v12431 = m.ExcPending
	if v12431 != 0 {
		goto L32
	} else {
		goto L2609
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
	v12415 = m.ExcPending
	if v12415 != 0 {
		goto L32
	} else {
		goto L2605
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
	v2371 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2371 != int32(1) {
		goto L538
	} else {
		goto L539
	}
L188:
	;
	v2342 = int32(0)
	v2344 = v2318
	v2346 = v2319
	v2347 = v2320
	v2348 = v2321
	v2349 = v2322
	v2350 = v2323
	v2351 = v2324
	v2352 = v2325
	v2353 = v2326
	v2354 = v2327
	v2355 = v2328
	v2356 = v2329
	v2363 = v2330
	v2364 = v2331
	v2366 = v2333
	goto L187
L189:
	;
	v2293 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v2293 == int32(44) {
		v2318 = v1352
		v2319 = v1334
		v2320 = v1341
		v2321 = v1343
		v2322 = v1342
		v2323 = v1344
		v2324 = v1345
		v2325 = v1346
		v2326 = v1347
		v2327 = v1463
		v2328 = v1354
		v2329 = v1355
		v2330 = v1335
		v2331 = v1340
		v2333 = v1353
		goto L188
	} else {
		goto L533
	}
L190:
	;
	v2078 = F___fstatat(m, int32(-100), int32(_a_F_StartupXLOG_44), v706+int32(1152), int32(0))
	mBase = m.M
	goto L483
L191:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L32
	} else {
		goto L478
	}
L192:
	;
	v961 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v961 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L193:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v863 != 0 {
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
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[18]))
	if v898 != 0 {
		goto L262
	} else {
		goto L263
	}
L244:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[19]))
	if v865 != 0 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v890 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[20]))
	if v890 == int32(0) {
		goto L191
	} else {
		goto L260
	}
L247:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865))))
	if v866 != 0 {
		goto L243
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[20]))
	if v868 != 0 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L249
L251:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868))))
	if v869 != 0 {
		goto L243
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v872 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L32
	} else {
		goto L255
	}
L254:
	;
	goto L253
L255:
	;
	if v872 == int32(0) {
		goto L243
	} else {
		goto L256
	}
L256:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_53), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L32
	} else {
		goto L257
	}
L257:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_54), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L32
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1142), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L32
	} else {
		goto L259
	}
L259:
	;
	goto L243
L260:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890))))
	if v893 == int32(0) {
		goto L191
	} else {
		goto L261
	}
L261:
	;
	goto L243
L262:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v907 == int32(2) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v900&int32(1) != 0 {
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
	v912 = int32(0)
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[23]))
	v917 = F_DirectFunctionCall3Coll(m, int32(411), v912, v914, v912, int32(-1))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L32
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[24]))
	switch v923 - int32(1) {
	case 0:
		goto L270
	case 1:
		goto L271
	default:
		goto L192
	}
L268:
	;
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v917)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25])) = v919
	goto L267
L269:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12])) = v956
	goto L192
L270:
	;
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v953 = F_findNewestTimeLine(m, v952)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L32
	} else {
		goto L279
	}
L271:
	;
	v926 = int32(1)
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[26]))
	if v928 == v926 {
		v956 = v926
		goto L269
	} else {
		goto L272
	}
L272:
	;
	v931 = F_existsTimeLineHistory(m, v928)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L32
	} else {
		goto L273
	}
L273:
	;
	if v931 != 0 {
		v956 = v928
		goto L269
	} else {
		goto L274
	}
L274:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L32
	} else {
		goto L275
	}
L275:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L32
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+848)) = v928
	F_errmsg(m, int32(_a_F_StartupXLOG_56), v706+int32(848))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L32
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1189), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
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
	v956 = v953
	goto L269
L280:
	;
	v965 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_OwnLatch(m, v965+int32(4))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L32
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v971 = F_palloc0(m, int32(12))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
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
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v984 = F_XLogReaderAllocate(m, v981, v706+int32(876), v971)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L32
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28])) = v984
	if v984 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v667)))
	*(*int64)(unsafe.Add(mBase, uint32(v984)+16)) = v987
	v990 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[29]))
	v991 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v984)+116)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v984)+104)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v984)+100)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v984)+112)) = v991
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v1000 = m.G0
	v1002 = v1000 - int32(48)
	m.G0 = v1002
	v1005 = F_palloc0(m, int32(136))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
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
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L32
	} else {
		goto L473
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1005))) = v999
	*(*int64)(unsafe.Add(mBase, uint32(v1002)+16)) = int64(171798691852)
	v1013 = F_hash_create(m, int32(_a_F_StartupXLOG_57), int32(1024), v1002, int32(40))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L32
	} else {
		goto L290
	}
L290:
	;
	v1016 = v1005 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v1005)+32)) = v1016
	*(*int32)(unsafe.Add(mBase, uint32(v1005)+28)) = v1016
	*(*int32)(unsafe.Add(mBase, uint32(v1005)+24)) = v1013
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1021)+56)) = int64(0)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1005)+128)) = v1027 - int32(1)
	m.G0 = v1002 + int32(48)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32])) = v1005
	v1038 = F_palloc(m, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L32
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33])) = v1038
	v1043 = F_palloc(m, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L32
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34])) = v1043
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = int64(0)
	v1050 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v1050
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v1050)
	v1057 = F_AllocateFile(m, int32(_a_F_StartupXLOG_59), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L32
	} else {
		goto L293
	}
L293:
	;
	if v1057 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v1062 == int32(44) {
		goto L190
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1084 = v706 + int32(1079)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+720)) = v1084
	*(*int32)(unsafe.Add(mBase, uint32(v706)+716)) = v706 + int32(1088)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+712)) = v706 + int32(1084)
	v1093 = v706 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+708)) = v1093
	v1096 = v706 + int32(892)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+704)) = v1096
	v1101 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_61), v706+int32(704))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L32
	} else {
		goto L303
	}
L297:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L32
	} else {
		goto L298
	}
L298:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L32
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+832)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(832))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L32
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1258), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
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
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L32
	} else {
		goto L469
	}
L303:
	;
	if v1101 != int32(5) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1079)))
	if v1105 != int32(10) {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = v1109
	v1112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+888)))
	v1113 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v1112 | v1113<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+688)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(v706)+692)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(v706)+696)) = v1084
	v1124 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_64), v706+int32(688))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L32
	} else {
		goto L307
	}
L306:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L32
	} else {
		goto L465
	}
L307:
	;
	if v1124 != int32(3) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1079)))
	if v1128 != int32(10) {
		goto L306
	} else {
		goto L309
	}
L309:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v1132
	v1135 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+888)))
	v1136 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v706)+892)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = v1135 | v1136<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+672)) = v706 + int32(1056)
	v1147 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_65), v706+int32(672))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L32
	} else {
		goto L311
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+656)) = v706 + int32(1024)
	v1167 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_66), v706+int32(656))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L32
	} else {
		goto L314
	}
L311:
	;
	if v1147 != int32(1) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1064)))
	v1152 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1056))
	if v1151|(v1152^int64(7234308641521824883)) != int64(0) {
		goto L310
	} else {
		goto L313
	}
L313:
	;
	v1159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v1159)
	goto L310
L314:
	;
	v1169 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1024))
	v1171 = v706 + int32(896)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+640)) = v1171
	v1176 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_67), v706+int32(640))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L32
	} else {
		goto L316
	}
L315:
	;
	v1200 = v706 + int32(1152)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+608)) = v1200
	v1205 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_68), v706+int32(608))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L32
	} else {
		goto L323
	}
L316:
	;
	if v1176 != int32(1) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1182 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L32
	} else {
		goto L318
	}
L318:
	;
	if v1182 == int32(0) {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+628)) = int32(_a_F_StartupXLOG_59)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+624)) = v1171
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_69), v706+int32(624))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L32
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1321), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
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
	v1234 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_70), v706+int32(576))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L32
	} else {
		goto L331
	}
L323:
	;
	if v1205 != int32(1) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1211 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L32
	} else {
		goto L325
	}
L325:
	;
	if v1211 == int32(0) {
		goto L322
	} else {
		goto L326
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+596)) = int32(_a_F_StartupXLOG_59)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+592)) = v1200
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_71), v706+int32(592))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L32
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1326), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
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
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L32
	} else {
		goto L460
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+516)) = v706 + int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+512)) = v706 + int32(892)
	v1270 = F_fscanf(m, v1057, int32(_a_F_StartupXLOG_72), v706+int32(512))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L32
	} else {
		goto L338
	}
L331:
	;
	if v1234 != int32(1) {
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	if v1238 != v1239 {
		goto L329
	} else {
		goto L333
	}
L333:
	;
	v1243 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L32
	} else {
		goto L334
	}
L334:
	;
	if v1243 == int32(0) {
		goto L330
	} else {
		goto L335
	}
L335:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+528)) = v1247
	*(*int32)(unsafe.Add(mBase, uint32(v706)+532)) = int32(_a_F_StartupXLOG_59)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_73), v706+int32(528))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L32
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1343), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L32
	} else {
		goto L337
	}
L337:
	;
	goto L330
L338:
	;
	if v1270 <= int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1057)))
	goto L343
L340:
	;
	goto L341
L341:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L32
	} else {
		goto L455
	}
L342:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L32
	} else {
		goto L451
	}
L343:
	;
	if int32(base.Ui32(v1274)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1279 = F_FreeFile(m, v1057)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L32
	} else {
		goto L345
	}
L345:
	;
	if v1279 != 0 {
		goto L342
	} else {
		goto L346
	}
L346:
	;
	v1282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v1282)
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v1285 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1287 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v1287)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L32
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1293 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L32
	} else {
		goto L351
	}
L350:
	;
	goto L349
L351:
	;
	if v1293 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+480)) = v1296
	v1299 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+468)) = uint32(v1299)
	v1302 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+476)) = uint32(v1302)
	v1304 = int64(32)
	v1305 = int64(base.Ui64(v1299) >> (uint(v1304) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+464)) = uint32(v1305)
	v1308 = int64(base.Ui64(v1302) >> (uint(v1304) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+472)) = uint32(v1308)
	F_errmsg(m, int32(_a_F_StartupXLOG_74), v706+int32(464))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L32
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v1325 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v1327 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v1328 = F_ReadCheckpointRecord(m, v1323, v1325, v1327)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L32
	} else {
		goto L358
	}
L355:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(627), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L32
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	v1463 = base.B2i32(v1167 == int32(1)) & base.B2i32(v1169 == int64(34166655670121587))
	v1466 = F_AllocateFile(m, int32(_a_F_StartupXLOG_44), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L32
	} else {
		goto L380
	}
L358:
	;
	if v1328 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+96))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+64))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+8))
	v1335 = *(*int64)(unsafe.Add(mBase, uint32(v1333)))
	v1336 = *(*int64)(unsafe.Add(mBase, uint32(v1333)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+896)) = v1336
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+904)) = v1338
	v1340 = *(*int64)(unsafe.Add(mBase, uint32(v1333)+24))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+32))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+36))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+40))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+44))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+48))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+52))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+56))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+1096)) = v1348
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v1333)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+1088)) = v1350
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328)+16)))
	v1353 = *(*int64)(unsafe.Add(mBase, uint32(v1333)+80))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+76))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+72))
	v1358 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
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
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L32
	} else {
		goto L376
	}
L362:
	;
	if v1358 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1361 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+452)) = uint32(v1361)
	v1364 = int64(base.Ui64(v1361) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+448)) = uint32(v1364)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_76), v706+int32(448))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L32
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v1378)
	v1381 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v1381) <= base.Ui64(v1335) {
		goto L357
	} else {
		goto L368
	}
L366:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(641), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L32
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	F_XLogPrefetcherBeginRead(m, v1384, v1335)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L32
	} else {
		goto L369
	}
L369:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v1391 = F_ReadRecord(m, v1388, int32(15), int32(0), v1334)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L32
	} else {
		goto L370
	}
L370:
	;
	if v1391 != 0 {
		goto L357
	} else {
		goto L371
	}
L371:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L32
	} else {
		goto L372
	}
L372:
	;
	v1398 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+92)) = uint32(v1398)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+84)) = uint32(v1335)
	v1401 = int64(32)
	v1402 = int64(base.Ui64(v1335) >> (uint(v1401) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+80)) = uint32(v1402)
	v1405 = int64(base.Ui64(v1398) >> (uint(v1401) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+88)) = uint32(v1405)
	F_errmsg(m, int32(_a_F_StartupXLOG_77), v706+int32(80))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L32
	} else {
		goto L373
	}
L373:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+64)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v706)+68)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v706)+72)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v706)+76)) = v1413
	F_errhint(m, int32(_a_F_StartupXLOG_78), v706-int32(-64))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L32
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(661), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
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
	v1433 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+52)) = uint32(v1433)
	v1436 = int64(base.Ui64(v1433) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+48)) = uint32(v1436)
	F_errmsg(m, int32(_a_F_StartupXLOG_79), v706+int32(48))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L32
	} else {
		goto L377
	}
L377:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+32)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v706)+36)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v706)+40)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v706)+44)) = v1444
	F_errhint(m, int32(_a_F_StartupXLOG_78), v706+int32(32))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L32
	} else {
		goto L378
	}
L378:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(672), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
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
	if v1466 == int32(0) {
		goto L189
	} else {
		goto L381
	}
L381:
	;
	v1470 = F_do_getc(m, v1466)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L32
	} else {
		goto L383
	}
L382:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1466)))
	goto L426
L383:
	;
	if v1470 == int32(-1) {
		v1789 = v1
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1474 = int32(0)
	v1477 = v1
	v1481 = v1474
	v1482 = v1474
	v1483 = v1470
	goto L385
L385:
	;
	if v1482&int32(1) != 0 {
		goto L391
	} else {
		goto L392
	}
L386:
	;
	if base.B2i32(v1731 == int32(0))&(v1732^int32(1)) != 0 {
		v1789 = v1727
		goto L382
	} else {
		goto L420
	}
L387:
	;
	v1760 = F_do_getc(m, v1466)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L32
	} else {
		goto L418
	}
L388:
	;
	v1727 = v1692
	v1731 = v1696
	v1732 = int32(0)
	goto L387
L389:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L32
	} else {
		goto L414
	}
L390:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L32
	} else {
		goto L410
	}
L391:
	;
	v1642 = (v1482 ^ int32(1)) & base.B2i32(v1483 == int32(92))
	if v1642|base.B2i32(base.Ui32(int32(1022)) < base.Ui32(v1481)) != 0 {
		v1727 = v1477
		v1731 = v1481
		v1732 = v1642
		goto L387
	} else {
		goto L409
	}
L392:
	;
	switch v1483 - int32(10) {
	case 0, 3:
		goto L393
	default:
		goto L391
	}
L393:
	;
	if v1481 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1514 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v706+int32(1152)+v1481))) = uint8(v1514)
	v1526 = v1514
	goto L397
L395:
	;
	v1604 = v1477
	goto L396
L396:
	;
	v1692 = v1604
	v1696 = int32(0)
	goto L388
L397:
	;
	v1556 = v706 + int32(1152) + v1526
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556))))
	v1558 = int32(32)
	if v1557|v1558 != v1558 {
		goto L399
	} else {
		goto L400
	}
L398:
	;
	if base.B2i32(v1526 <= int32(0))|base.B2i32(v1481-int32(1) <= v1526) != 0 {
		goto L390
	} else {
		goto L402
	}
L399:
	;
	v1526 = v1526 + int32(1)
	goto L397
L400:
	;
	goto L401
L401:
	;
	goto L398
L402:
	;
	v1570 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1556))) = uint8(v1570)
	v1573 = F_palloc0(m, int32(24))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L32
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v1584 = F_strtox_2(m, v706+int32(1152), v706+int32(1056), int32(10), int64(4294967295))
	mBase = m.M
	goto L404
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1573))) = base.I32_wrap_i64(v1584)
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1056))
	v1588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587))))
	if v1588 != 0 {
		goto L389
	} else {
		goto L405
	}
L405:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if base.B2i32(v1590 == int32(68))|base.B2i32(v1590 == int32(28)) != 0 {
		goto L389
	} else {
		goto L406
	}
L406:
	;
	v1598 = F_pstrdup(m, v1556+int32(1))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L32
	} else {
		goto L407
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1573)+4)) = v1598
	v1601 = F_lappend(m, v1477, v1573)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L32
	} else {
		goto L408
	}
L408:
	;
	v1604 = v1601
	goto L396
L409:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v706+int32(1152)+v1481))) = uint8(v1483)
	v1692 = v1477
	v1696 = v1481 + int32(1)
	goto L388
L410:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L32
	} else {
		goto L411
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+432)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(432))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L32
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1425), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
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
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L32
	} else {
		goto L415
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+416)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(416))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L32
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1434), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
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
	if v1760 != int32(-1) {
		v1477 = v1727
		v1481 = v1731
		v1482 = v1732
		v1483 = v1760
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
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L32
	} else {
		goto L421
	}
L421:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L32
	} else {
		goto L422
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+400)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(400))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L32
	} else {
		goto L423
	}
L423:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1454), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
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
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L32
	} else {
		goto L447
	}
L426:
	;
	if int32(base.Ui32(v1822)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1827 = F_FreeFile(m, v1466)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L32
	} else {
		goto L428
	}
L428:
	;
	if v1827 != 0 {
		goto L425
	} else {
		goto L429
	}
L429:
	;
	if v1789 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2342 = int32(1)
	v2344 = v1352
	v2346 = v1334
	v2347 = v1341
	v2348 = v1343
	v2349 = v1342
	v2350 = v1344
	v2351 = v1345
	v2352 = v1346
	v2353 = v1347
	v2354 = v1463
	v2355 = v1354
	v2356 = v1355
	v2363 = v1335
	v2364 = v1340
	v2366 = v1353
	goto L187
L431:
	;
	goto L432
L432:
	;
	v1832 = int32(1)
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+4))
	if v1833 <= int32(0) {
		v2342 = v1832
		v2344 = v1352
		v2346 = v1334
		v2347 = v1341
		v2348 = v1343
		v2349 = v1342
		v2350 = v1344
		v2351 = v1345
		v2352 = v1346
		v2353 = v1347
		v2354 = v1463
		v2355 = v1354
		v2356 = v1355
		v2363 = v1335
		v2364 = v1340
		v2366 = v1353
		goto L187
	} else {
		goto L433
	}
L433:
	;
	v1844 = int32(0)
	goto L434
L434:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+12))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1871+v1844<<(uint(int32(2))%32))))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1875)))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+372)) = v1876
	*(*int32)(unsafe.Add(mBase, uint32(v706)+368)) = int32(_a_F_StartupXLOG_38)
	v1883 = F_psprintf(m, int32(_a_F_StartupXLOG_82), v706+int32(368))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L32
	} else {
		goto L437
	}
L435:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L32
	} else {
		goto L443
	}
L436:
	;
	goto L435
L437:
	;
	F_remove_tablespace_symlink(m, v1883)
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L32
	} else {
		goto L438
	}
L438:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	v1888 = F_symlink(m, v1887, v1883)
	mBase = m.M
	if v1888 < int32(0) {
		goto L436
	} else {
		goto L439
	}
L439:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	F_pfree(m, v1891)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L32
	} else {
		goto L440
	}
L440:
	;
	F_pfree(m, v1875)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L32
	} else {
		goto L441
	}
L441:
	;
	v1897 = v1844 + int32(1)
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+4))
	if v1897 < v1898 {
		v1844 = v1897
		goto L434
	} else {
		goto L442
	}
L442:
	;
	v2342 = v1832
	v2344 = v1352
	v2346 = v1334
	v2347 = v1341
	v2348 = v1343
	v2349 = v1342
	v2350 = v1344
	v2351 = v1345
	v2352 = v1346
	v2353 = v1347
	v2354 = v1463
	v2355 = v1354
	v2356 = v1355
	v2363 = v1335
	v2364 = v1340
	v2366 = v1353
	goto L187
L443:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L32
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+352)) = v1883
	F_errmsg(m, int32(_a_F_StartupXLOG_83), v706+int32(352))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L32
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(698), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
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
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L32
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+384)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(384))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L32
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1460), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
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
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L32
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+496)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(496))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L32
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1356), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
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
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L32
	} else {
		goto L456
	}
L456:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_84), int32(0))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L32
	} else {
		goto L457
	}
L457:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_85), int32(0))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L32
	} else {
		goto L458
	}
L458:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1350), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
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
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L32
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+560)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(560))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L32
	} else {
		goto L462
	}
L462:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1080))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+544)) = v1987
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1084))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+548)) = v1989
	F_errdetail(m, int32(_a_F_StartupXLOG_86), v706+int32(544))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L32
	} else {
		goto L463
	}
L463:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1339), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
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
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L32
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+16)) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706+int32(16))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L32
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1278), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
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
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L32
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706))) = int32(_a_F_StartupXLOG_59)
	F_errmsg(m, int32(_a_F_StartupXLOG_80), v706)
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L32
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1271), int32(_a_F_StartupXLOG_63))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
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
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L32
	} else {
		goto L474
	}
L474:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_88), int32(0))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L32
	} else {
		goto L475
	}
L475:
	;
	F_errdetail(m, int32(_a_F_StartupXLOG_89), int32(0))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L32
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(572), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
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
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L32
	} else {
		goto L479
	}
L479:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_90), int32(0))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L32
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1150), int32(_a_F_StartupXLOG_55))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
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
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2121 != int32(1) {
		goto L497
	} else {
		goto L498
	}
L483:
	;
	if v2078 != 0 {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2079 = int32(_a_F_StartupXLOG_91)
	v2080 = F_unlink(m, v2079)
	mBase = m.M
	v2084 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v2079, int32(14))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L32
	} else {
		goto L485
	}
L485:
	;
	v2088 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L32
	} else {
		goto L486
	}
L486:
	;
	if v2088 == int32(0) {
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
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L32
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+804)) = int32(_a_F_StartupXLOG_91)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+800)) = int32(_a_F_StartupXLOG_44)
	if v2084 != 0 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2107 = int32(_a_F_StartupXLOG_93)
	goto L491
L490:
	;
	v2107 = int32(_a_F_StartupXLOG_94)
	goto L491
L491:
	;
	F_errdetail(m, v2107, v706+int32(800))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L32
	} else {
		goto L492
	}
L492:
	;
	if v2084 != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v2115 = int32(739)
	goto L495
L494:
	;
	v2115 = int32(733)
	goto L495
L495:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2115, int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L32
	} else {
		goto L496
	}
L496:
	;
	goto L482
L497:
	;
	v2146 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	if v2146 == int64(0) {
		goto L506
	} else {
		goto L507
	}
L498:
	;
	v2124 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if v2124 != int64(0) {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v2135)
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2138 == int32(0) {
		goto L497
	} else {
		goto L504
	}
L500:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	if v2127 != 0 {
		goto L499
	} else {
		goto L501
	}
L501:
	;
	v2128 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	if v2128 != int64(0) {
		goto L499
	} else {
		goto L502
	}
L502:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	if v2131 != int32(1) {
		goto L497
	} else {
		goto L503
	}
L503:
	;
	goto L499
L504:
	;
	v2142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v2142)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L32
	} else {
		goto L505
	}
L505:
	;
	goto L497
L506:
	;
	v2172 = *(*int64)(unsafe.Add(mBase, uint32(v667)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35])) = v2172
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36])) = v2175
	v2177 = *(*int64)(unsafe.Add(mBase, uint32(v667)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38])) = v2175
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39])) = v2177
	v2183 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v2184 = F_ReadCheckpointRecord(m, v2183, v2172, v2175)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L32
	} else {
		goto L513
	}
L507:
	;
	v2151 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L32
	} else {
		goto L508
	}
L508:
	;
	if v2151 == int32(0) {
		goto L506
	} else {
		goto L509
	}
L509:
	;
	v2155 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+788)) = uint32(v2155)
	v2158 = int64(base.Ui64(v2155) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+784)) = uint32(v2158)
	F_errmsg(m, int32(_a_F_StartupXLOG_95), v706+int32(784))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L32
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(778), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
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
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L32
	} else {
		goto L530
	}
L513:
	;
	if v2184 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2188 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
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
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L32
	} else {
		goto L527
	}
L517:
	;
	if v2188 != 0 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2191 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+772)) = uint32(v2191)
	v2194 = int64(base.Ui64(v2191) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+768)) = uint32(v2194)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_76), v706+int32(768))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L32
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+96))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2209)+64))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+8))
	v2212 = *(*int64)(unsafe.Add(mBase, uint32(v2210)))
	v2213 = *(*int64)(unsafe.Add(mBase, uint32(v2210)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+896)) = v2213
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+904)) = v2215
	v2217 = *(*int64)(unsafe.Add(mBase, uint32(v2210)+24))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+32))
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+36))
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+40))
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+44))
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+48))
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+52))
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+56))
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+1096)) = v2225
	v2227 = *(*int64)(unsafe.Add(mBase, uint32(v2210)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v706)+1088)) = v2227
	v2229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184)+16)))
	v2230 = *(*int64)(unsafe.Add(mBase, uint32(v2210)+80))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+76))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+72))
	v2234 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v2234) <= base.Ui64(v2212) {
		v2318 = v2229
		v2319 = v2211
		v2320 = v2218
		v2321 = v2220
		v2322 = v2219
		v2323 = v2221
		v2324 = v2222
		v2325 = v2223
		v2326 = v2224
		v2327 = v1
		v2328 = v2231
		v2329 = v2232
		v2330 = v2212
		v2331 = v2217
		v2333 = v2230
		goto L188
	} else {
		goto L523
	}
L521:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(791), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L32
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	F_XLogPrefetcherBeginRead(m, v2237, v2212)
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L32
	} else {
		goto L524
	}
L524:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v2244 = F_ReadRecord(m, v2241, int32(15), int32(0), v2211)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L32
	} else {
		goto L525
	}
L525:
	;
	if v2244 == int32(0) {
		goto L512
	} else {
		goto L526
	}
L526:
	;
	v2318 = v2229
	v2319 = v2211
	v2320 = v2218
	v2321 = v2220
	v2322 = v2219
	v2323 = v2221
	v2324 = v2222
	v2325 = v2223
	v2326 = v2224
	v2327 = v1
	v2328 = v2231
	v2329 = v2232
	v2330 = v2212
	v2331 = v2217
	v2333 = v2230
	goto L188
L527:
	;
	v2253 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+740)) = uint32(v2253)
	v2256 = int64(base.Ui64(v2253) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+736)) = uint32(v2256)
	F_errmsg(m, int32(_a_F_StartupXLOG_96), v706+int32(736))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L32
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(803), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
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
	v2273 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+764)) = uint32(v2273)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+756)) = uint32(v2212)
	v2276 = int64(32)
	v2277 = int64(base.Ui64(v2212) >> (uint(v2276) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+752)) = uint32(v2277)
	v2280 = int64(base.Ui64(v2273) >> (uint(v2276) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+760)) = uint32(v2280)
	F_errmsg(m, int32(_a_F_StartupXLOG_97), v706+int32(752))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L32
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(815), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
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
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L32
	} else {
		goto L534
	}
L534:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L32
	} else {
		goto L535
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+336)) = int32(_a_F_StartupXLOG_44)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v706+int32(336))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L32
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1393), int32(_a_F_StartupXLOG_81))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
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
	v2470 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	v2472 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2473 = F_tliOfPointInHistory(m, v2470, v2472)
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L32
	} else {
		goto L572
	}
L539:
	;
	v2376 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[14])))
	if v2376 != 0 {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), v2462, int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L32
	} else {
		goto L567
	}
L541:
	;
	v2379 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L32
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v2389 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	v2392 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2393 = m.ExcPending
	if v2393 != 0 {
		goto L32
	} else {
		goto L547
	}
L544:
	;
	if v2379 == int32(0) {
		goto L538
	} else {
		goto L545
	}
L545:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_98), int32(0))
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L32
	} else {
		goto L546
	}
L546:
	;
	v2462 = int32(823)
	goto L540
L547:
	;
	switch v2389 - int32(1) {
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
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L565
	}
L549:
	;
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L563
	}
L550:
	;
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L561
	}
L551:
	;
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L559
	}
L552:
	;
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L556
	}
L553:
	;
	if v2392 == int32(0) {
		goto L538
	} else {
		goto L554
	}
L554:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+272)) = v2399
	F_errmsg(m, int32(_a_F_StartupXLOG_99), v706+int32(272))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L32
	} else {
		goto L555
	}
L555:
	;
	v2462 = int32(827)
	goto L540
L556:
	;
	v2410 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25]))
	v2411 = F_timestamptz_to_str(m, v2410)
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L32
	} else {
		goto L557
	}
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+288)) = v2411
	F_errmsg(m, int32(_a_F_StartupXLOG_100), v706+int32(288))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L32
	} else {
		goto L558
	}
L558:
	;
	v2462 = int32(831)
	goto L540
L559:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+304)) = v2423
	F_errmsg(m, int32(_a_F_StartupXLOG_101), v706+int32(304))
	mBase = m.M
	v2429 = m.ExcPending
	if v2429 != 0 {
		goto L32
	} else {
		goto L560
	}
L560:
	;
	v2462 = int32(835)
	goto L540
L561:
	;
	v2434 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+324)) = uint32(v2434)
	v2437 = int64(base.Ui64(v2434) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+320)) = uint32(v2437)
	F_errmsg(m, int32(_a_F_StartupXLOG_102), v706+int32(320))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L32
	} else {
		goto L562
	}
L562:
	;
	v2462 = int32(839)
	goto L540
L563:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_103), int32(0))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L32
	} else {
		goto L564
	}
L564:
	;
	v2462 = int32(842)
	goto L540
L565:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_104), int32(0))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L32
	} else {
		goto L566
	}
L566:
	;
	v2462 = int32(845)
	goto L540
L567:
	;
	goto L538
L568:
	;
	v2892 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+120))
	v2894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2892)+56)))
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+48))
	v2896 = *(*int64)(unsafe.Add(mBase, uint32(v2892)+40))
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+116))
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+112))
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+96))
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+92))
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+88))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+84))
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+80))
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+76))
	v2905 = *(*int64)(unsafe.Add(mBase, uint32(v2892)+64))
	v2906 = int32(_a_F_StartupXLOG_105)
	v2907 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v2892)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v2907))) = v2908
	*(*int64)(unsafe.Add(mBase, uint32(v2907)+8)) = v2905
	v2912 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	*(*int32)(unsafe.Add(mBase, uint32(v2912)+4)) = int32(0)
	F_MultiXactSetNextMXact(m, v2904, v2903)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L32
	} else {
		goto L681
	}
L569:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		goto L32
	} else {
		goto L678
	}
L570:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L32
	} else {
		goto L675
	}
L571:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2844 = m.ExcPending
	if v2844 != 0 {
		goto L32
	} else {
		goto L672
	}
L572:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	if v2473 == v2476 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v2478 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if v2478 != int64(0) {
		goto L576
	} else {
		goto L577
	}
L574:
	;
	goto L575
L575:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2800 = F_tliSwitchPoint(m, v2476, v2798, int32(0))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L32
	} else {
		goto L664
	}
L576:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v2485 = F_tliOfPointInHistory(m, v2478-int64(1), v2484)
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L32
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2491 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L32
	} else {
		goto L581
	}
L579:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	if v2485 != v2487 {
		goto L571
	} else {
		goto L580
	}
L580:
	;
	goto L578
L581:
	;
	if v2491 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+196)) = uint32(v2363)
	v2495 = int64(base.Ui64(v2363) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+192)) = uint32(v2495)
	if base.Ui32(v2344) < base.Ui32(int32(16)) {
		goto L585
	} else {
		goto L586
	}
L583:
	;
	goto L584
L584:
	;
	v2515 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L32
	} else {
		goto L590
	}
L585:
	;
	v2501 = int32(_a_F_StartupXLOG_106)
	goto L587
L586:
	;
	v2501 = int32(_a_F_StartupXLOG_107)
	goto L587
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+200)) = v2501
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_108), v706+int32(192))
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L32
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(892), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L32
	} else {
		goto L589
	}
L589:
	;
	goto L584
L590:
	;
	if v2515 != 0 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+184)) = v2347
	*(*int64)(unsafe.Add(mBase, uint32(v706)+176)) = v2364
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_109), v706+int32(176))
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L32
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v2531 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L32
	} else {
		goto L596
	}
L594:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(896), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L32
	} else {
		goto L595
	}
L595:
	;
	goto L593
L596:
	;
	if v2531 != 0 {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+164)) = v2348
	*(*int32)(unsafe.Add(mBase, uint32(v706)+160)) = v2349
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_110), v706+int32(160))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L32
	} else {
		goto L600
	}
L598:
	;
	goto L599
L599:
	;
	v2547 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L32
	} else {
		goto L602
	}
L600:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(899), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L32
	} else {
		goto L601
	}
L601:
	;
	goto L599
L602:
	;
	if v2547 != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+148)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v706)+144)) = v2350
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_111), v706+int32(144))
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L32
	} else {
		goto L606
	}
L604:
	;
	goto L605
L605:
	;
	v2563 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		goto L32
	} else {
		goto L608
	}
L606:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(902), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2560 = m.ExcPending
	if v2560 != 0 {
		goto L32
	} else {
		goto L607
	}
L607:
	;
	goto L605
L608:
	;
	if v2563 != 0 {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+132)) = v2353
	*(*int32)(unsafe.Add(mBase, uint32(v706)+128)) = v2352
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_112), v706+int32(128))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		goto L32
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v2579 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L32
	} else {
		goto L614
	}
L612:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(905), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L32
	} else {
		goto L613
	}
L613:
	;
	goto L611
L614:
	;
	if v2579 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+116)) = v2355
	*(*int32)(unsafe.Add(mBase, uint32(v706)+112)) = v2356
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_113), v706+int32(112))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L32
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	if base.Ui32(base.I32_wrap_i64(v2364)) <= base.Ui32(int32(2)) {
		goto L570
	} else {
		goto L620
	}
L618:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(909), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L32
	} else {
		goto L619
	}
L619:
	;
	goto L617
L620:
	;
	v2597 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v2597) < base.Ui64(v2363) {
		goto L569
	} else {
		goto L621
	}
L621:
	;
	if base.Ui64(v2363) < base.Ui64(v2597) {
		goto L628
	} else {
		goto L629
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49])) = v2773
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50])) = v2776
	v2782 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51])) = v2782
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v2782
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4095)))) = uint8(base.B2i32(base.Ui32(v2344) < base.Ui32(int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4093)))) = uint8(base.B2i32(v1057 != int32(0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(4094)))) = uint8(v2342)
	m.G0 = v706 + int32(2176)
	goto L568
L623:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	v2772 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	v2773 = v2771
	v2776 = v2772
	goto L622
L624:
	;
	v2757 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v2757
	v2760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v2760)
	v2763 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2763
	if v2753&int32(1) != 0 {
		goto L623
	} else {
		goto L663
	}
L625:
	;
	if v2635&int32(1) != 0 {
		v2678 = int32(5)
		goto L638
	} else {
		goto L639
	}
L626:
	;
	v2630 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v2632 != int32(1) {
		v2753 = v2630
		goto L624
	} else {
		goto L637
	}
L627:
	;
	v2625 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v2625)
	v2628 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	v2635 = v2628
	goto L625
L628:
	;
	if base.Ui32(int32(15)) < base.Ui32(v2344) {
		goto L627
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	if v2615 != int32(1) {
		goto L627
	} else {
		goto L635
	}
L631:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L32
	} else {
		goto L632
	}
L632:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_114), int32(0))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L32
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(928), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
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
	v2619 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v2619&int32(1) == int32(0) {
		goto L626
	} else {
		goto L636
	}
L636:
	;
	goto L627
L637:
	;
	v2635 = v2630
	goto L625
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+16)) = v2678
	v2681 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+48)) = v2346
	*(*int64)(unsafe.Add(mBase, uint32(v667)+40)) = v2363
	*(*int64)(unsafe.Add(mBase, uint32(v667)+32)) = v2681
	v2685 = *(*int64)(unsafe.Add(mBase, uint32(v706)+896))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+52)) = v2685
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v706)+904))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+60)) = v2687
	*(*int32)(unsafe.Add(mBase, uint32(v667)+96)) = v2353
	*(*int32)(unsafe.Add(mBase, uint32(v667)+92)) = v2352
	*(*int32)(unsafe.Add(mBase, uint32(v667)+88)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v667)+84)) = v2350
	*(*int32)(unsafe.Add(mBase, uint32(v667)+80)) = v2348
	*(*int32)(unsafe.Add(mBase, uint32(v667)+76)) = v2349
	*(*int32)(unsafe.Add(mBase, uint32(v667)+72)) = v2347
	*(*int64)(unsafe.Add(mBase, uint32(v667)+64)) = v2364
	v2697 = *(*int64)(unsafe.Add(mBase, uint32(v706)+1088))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+100)) = v2697
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v706)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+108)) = v2699
	*(*int64)(unsafe.Add(mBase, uint32(v667)+120)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v667)+116)) = v2355
	*(*int32)(unsafe.Add(mBase, uint32(v667)+112)) = v2356
	v2705 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	if v2705 != int32(1) {
		goto L651
	} else {
		goto L652
	}
L639:
	;
	v2641 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L32
	} else {
		goto L640
	}
L640:
	;
	if v2641 != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_115), int32(0))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L32
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	v2652 = int32(4)
	v2654 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	if base.Ui32(v2654) <= base.Ui32(v2655) {
		v2678 = v2652
		goto L638
	} else {
		goto L646
	}
L644:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(958), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L32
	} else {
		goto L645
	}
L645:
	;
	goto L643
L646:
	;
	v2659 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L32
	} else {
		goto L647
	}
L647:
	;
	if v2659 == int32(0) {
		v2678 = v2652
		goto L638
	} else {
		goto L648
	}
L648:
	;
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+96)) = v2663
	v2666 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+100)) = v2666
	F_errmsg(m, int32(_a_F_StartupXLOG_116), v706+int32(96))
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L32
	} else {
		goto L649
	}
L649:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(964), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L32
	} else {
		goto L650
	}
L650:
	;
	v2678 = v2652
	goto L638
L651:
	;
	if v1057 == int32(0) {
		v2753 = v2705
		goto L624
	} else {
		goto L654
	}
L652:
	;
	v2708 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	if base.Ui64(v2363) <= base.Ui64(v2708) {
		goto L651
	} else {
		goto L653
	}
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+144)) = v2346
	*(*int64)(unsafe.Add(mBase, uint32(v667)+136)) = v2363
	goto L651
L654:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v667)+152)) = v2363
	v2716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])))
	*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)) = uint8(v2716)
	if v2354 == int32(0) {
		v2753 = v2705
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
	v2739 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+160)) = v2739
	v2743 = *(*int64)(unsafe.Add(mBase, uint32(v667)+152))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53])) = v2743
	v2746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+168)))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[37])) = uint8(v2746)
	v2749 = *(*int64)(unsafe.Add(mBase, uint32(v667)+160))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v2749
	if v2705 != 0 {
		goto L623
	} else {
		goto L662
	}
L657:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L32
	} else {
		goto L658
	}
L658:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_117), int32(0))
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L32
	} else {
		goto L659
	}
L659:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_118), int32(0))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L32
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1006), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2738 = m.ExcPending
	if v2738 != 0 {
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
	v2773 = int32(0)
	v2776 = int64(0)
	goto L622
L663:
	;
	v2773 = int32(0)
	v2776 = int64(0)
	goto L622
L664:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L32
	} else {
		goto L665
	}
L665:
	;
	v2807 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+256)) = v2807
	F_errmsg(m, int32(_a_F_StartupXLOG_119), v706+int32(256))
	mBase = m.M
	v2813 = m.ExcPending
	if v2813 != 0 {
		goto L32
	} else {
		goto L666
	}
L666:
	;
	v2815 = int64(base.Ui64(v2800) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+240)) = uint32(v2815)
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+244)) = uint32(v2800)
	if v1057 != 0 {
		goto L667
	} else {
		goto L668
	}
L667:
	;
	v2820 = int32(_a_F_StartupXLOG_59)
	goto L669
L668:
	;
	v2820 = int32(_a_F_StartupXLOG_120)
	goto L669
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+224)) = v2820
	v2823 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+232)) = uint32(v2823)
	v2826 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+236)) = v2826
	v2829 = int64(base.Ui64(v2823) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+228)) = uint32(v2829)
	F_errdetail(m, int32(_a_F_StartupXLOG_121), v706+int32(224))
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L32
	} else {
		goto L670
	}
L670:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(873), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
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
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v667)+144))
	v2846 = *(*int64)(unsafe.Add(mBase, uint32(v667)+136))
	v2848 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+208)) = v2848
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+216)) = uint32(v2846)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+220)) = v2845
	v2853 = int64(base.Ui64(v2846) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v706)+212)) = uint32(v2853)
	F_errmsg(m, int32(_a_F_StartupXLOG_122), v706+int32(208))
	mBase = m.M
	v2859 = m.ExcPending
	if v2859 != 0 {
		goto L32
	} else {
		goto L673
	}
L673:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(887), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
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
	v2872 = m.ExcPending
	if v2872 != 0 {
		goto L32
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(912), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
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
	v2885 = m.ExcPending
	if v2885 != 0 {
		goto L32
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(917), int32(_a_F_StartupXLOG_75))
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
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
	F_AdvanceOldestClogXid(m, v2902)
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L32
	} else {
		goto L682
	}
L682:
	;
	F_SetTransactionIdLimit(m, v2902, v2901)
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L32
	} else {
		goto L683
	}
L683:
	;
	F_SetMultiXactIdLimit(m, v2900, v2899, int32(1))
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L32
	} else {
		goto L684
	}
L684:
	;
	F_SetCommitTsLimit(m, v2898, v2897)
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L32
	} else {
		goto L685
	}
L685:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v2927)+208)) = v2905
	v2929 = m.G0
	v2931 = v2929 - int32(1088)
	m.G0 = v2931
	*(*int32)(unsafe.Add(mBase, uint32(v2931)+16)) = int32(_a_F_StartupXLOG_125)
	v2936 = v2931 + int32(32)
	v2941 = F_pg_snprintf(m, v2936, int32(1050), int32(_a_F_StartupXLOG_126), v2931+int32(16))
	mBase = m.M
	v2942 = m.ExcPending
	if v2942 != 0 {
		goto L32
	} else {
		goto L686
	}
L686:
	;
	F_unlink_initfile(m, v2936, int32(15))
	mBase = m.M
	v2945 = m.ExcPending
	if v2945 != 0 {
		goto L32
	} else {
		goto L687
	}
L687:
	;
	F_RelationCacheInitFileRemoveInDir(m, int32(_a_F_StartupXLOG_127))
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L32
	} else {
		goto L688
	}
L688:
	;
	v2950 = F_AllocateDir(m, int32(_a_F_StartupXLOG_38))
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		goto L32
	} else {
		goto L689
	}
L689:
	;
	v2954 = F_ReadDirExtended(m, v2950, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L32
	} else {
		goto L690
	}
L690:
	;
	if v2954 != 0 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v2956 = v2954
	goto L694
L692:
	;
	goto L693
L693:
	;
	F_FreeDir(m, v2950)
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L32
	} else {
		goto L722
	}
L694:
	;
	v2991 = v2956 + int32(19)
	v2992 = int32(_a_F_StartupXLOG_128)
	v2996 = m.G0
	v2998 = v2996 - int32(32)
	v2999 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2998)+24)) = v2999
	*(*int64)(unsafe.Add(mBase, uint32(v2998)+16)) = v2999
	*(*int64)(unsafe.Add(mBase, uint32(v2998)+8)) = v2999
	*(*int64)(unsafe.Add(mBase, uint32(v2998))) = v2999
	v3007 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[56])))
	if v3007 == int32(0) {
		goto L697
	} else {
		goto L698
	}
L695:
	;
	goto L693
L696:
	;
	v3076 = F_strlen(m, v2991)
	mBase = m.M
	if v3075 == v3076 {
		goto L715
	} else {
		goto L716
	}
L697:
	;
	v3075 = int32(0)
	goto L696
L698:
	;
	goto L699
L699:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[57])))
	if v3011 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3015 = v2991
	goto L703
L701:
	;
	goto L702
L702:
	;
	v3025 = v2992
	v3026 = v3007
	goto L706
L703:
	;
	v3021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3015))))
	if v3021 == v3007 {
		v3015 = v3015 + int32(1)
		goto L703
	} else {
		goto L705
	}
L704:
	;
	v3075 = v3015 - v2991
	goto L696
L705:
	;
	goto L704
L706:
	;
	v3033 = v2998 + int32(base.Ui32(v3026)>>(uint(int32(3))%32))&int32(28)
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3033)))
	v3035 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3033))) = v3034 | v3035<<(uint(v3026)%32)
	v3039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3025)+1)))
	if v3039 != 0 {
		v3025 = v3025 + v3035
		v3026 = v3039
		goto L706
	} else {
		goto L708
	}
L707:
	;
	v3042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2991))))
	if v3042 == int32(0) {
		v3065 = v2991
		goto L709
	} else {
		goto L710
	}
L708:
	;
	goto L707
L709:
	;
	v3075 = v3065 - v2991
	goto L696
L710:
	;
	v3046 = v2991
	v3047 = v3042
	goto L711
L711:
	;
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v2998+int32(base.Ui32(v3047)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3055)>>(uint(v3047)%32))&int32(1) == int32(0) {
		v3065 = v3046
		goto L709
	} else {
		goto L713
	}
L712:
	;
	v3065 = v3063
	goto L709
L713:
	;
	v3061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3046)+1)))
	v3063 = v3046 + int32(1)
	if v3061 != 0 {
		v3046 = v3063
		v3047 = v3061
		goto L711
	} else {
		goto L714
	}
L714:
	;
	goto L712
L715:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2931)+8)) = int32(_a_F_StartupXLOG_129)
	*(*int32)(unsafe.Add(mBase, uint32(v2931)+4)) = v2991
	*(*int32)(unsafe.Add(mBase, uint32(v2931))) = int32(_a_F_StartupXLOG_38)
	v3084 = v2931 + int32(32)
	v3087 = F_pg_snprintf(m, v3084, int32(1050), int32(_a_F_StartupXLOG_130), v2931)
	mBase = m.M
	v3088 = m.ExcPending
	if v3088 != 0 {
		goto L32
	} else {
		goto L718
	}
L716:
	;
	goto L717
L717:
	;
	v3094 = F_ReadDirExtended(m, v2950, int32(_a_F_StartupXLOG_38), int32(15))
	mBase = m.M
	v3095 = m.ExcPending
	if v3095 != 0 {
		goto L32
	} else {
		goto L720
	}
L718:
	;
	F_RelationCacheInitFileRemoveInDir(m, v3084)
	mBase = m.M
	v3090 = m.ExcPending
	if v3090 != 0 {
		goto L32
	} else {
		goto L719
	}
L719:
	;
	goto L717
L720:
	;
	if v3094 != 0 {
		v2956 = v3094
		goto L694
	} else {
		goto L721
	}
L721:
	;
	goto L695
L722:
	;
	m.G0 = v2931 + int32(1088)
	v3135 = m.G0
	v3137 = v3135 - int32(3696)
	m.G0 = v3137
	v3141 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L32
	} else {
		goto L723
	}
L723:
	;
	if v3141 != 0 {
		goto L724
	} else {
		goto L725
	}
L724:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_131), int32(0))
	mBase = m.M
	v3146 = m.ExcPending
	if v3146 != 0 {
		goto L32
	} else {
		goto L727
	}
L725:
	;
	goto L726
L726:
	;
	v3153 = F_AllocateDir(m, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L32
	} else {
		goto L743
	}
L727:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2202), int32(_a_F_StartupXLOG_134))
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L32
	} else {
		goto L728
	}
L728:
	;
	goto L726
L729:
	;
	v4004 = F_AllocateDir(m, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4005 = m.ExcPending
	if v4005 != 0 {
		goto L32
	} else {
		goto L919
	}
L730:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L32
	} else {
		goto L915
	}
L731:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		goto L32
	} else {
		goto L910
	}
L732:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L32
	} else {
		goto L905
	}
L733:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L32
	} else {
		goto L902
	}
L734:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L32
	} else {
		goto L898
	}
L735:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L32
	} else {
		goto L895
	}
L736:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3836 = m.ExcPending
	if v3836 != 0 {
		goto L32
	} else {
		goto L891
	}
L737:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3815 = m.ExcPending
	if v3815 != 0 {
		goto L32
	} else {
		goto L887
	}
L738:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L32
	} else {
		goto L883
	}
L739:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v3775 = m.ExcPending
	if v3775 != 0 {
		goto L32
	} else {
		goto L880
	}
L740:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3758 = m.ExcPending
	if v3758 != 0 {
		goto L32
	} else {
		goto L876
	}
L741:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L32
	} else {
		goto L872
	}
L742:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L32
	} else {
		goto L868
	}
L743:
	;
	v3156 = F_ReadDir(m, v3153, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3157 = m.ExcPending
	if v3157 != 0 {
		goto L32
	} else {
		goto L744
	}
L744:
	;
	if v3156 != 0 {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	v3159 = v3137 + int32(3512)
	v3168 = v3156
	goto L748
L746:
	;
	goto L747
L747:
	;
	F_FreeDir(m, v3153)
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L32
	} else {
		goto L862
	}
L748:
	;
	v3196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3168)+19)))
	if v3196 != int32(46) {
		goto L751
	} else {
		goto L752
	}
L749:
	;
	goto L747
L750:
	;
	v3667 = F_ReadDir(m, v3153, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L32
	} else {
		goto L860
	}
L751:
	;
	v3209 = v3168 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+340)) = v3209
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+336)) = int32(_a_F_StartupXLOG_132)
	v3214 = v3137 + int32(352)
	v3219 = F_pg_snprintf(m, v3214, int32(1036), int32(_a_F_StartupXLOG_135), v3137+int32(336))
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L32
	} else {
		goto L756
	}
L752:
	;
	v3199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3168)+20)))
	if v3199 == int32(0) {
		goto L750
	} else {
		goto L753
	}
L753:
	;
	v3202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3168)+20)))
	if v3202 != int32(46) {
		goto L751
	} else {
		goto L754
	}
L754:
	;
	v3205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3168)+21)))
	if v3205 == int32(0) {
		goto L750
	} else {
		goto L755
	}
L755:
	;
	goto L751
L756:
	;
	v3223 = F_get_dirent_type(m, v3214, v3168, int32(0), int32(14))
	mBase = m.M
	v3224 = m.ExcPending
	if v3224 != 0 {
		goto L32
	} else {
		goto L758
	}
L757:
	;
	v3225 = F_strlen(m, v3209)
	mBase = m.M
	v3227 = F_strlen(m, int32(_a_F_StartupXLOG_136))
	mBase = m.M
	if base.Ui32(v3227) <= base.Ui32(v3225) {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	switch v3223 {
	case 0, 3:
		goto L757
	default:
		goto L750
	}
L759:
	;
	v3230 = v3209 + (v3225 - v3227)
	v3231 = int32(_a_F_StartupXLOG_136)
	v3234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3230))))
	v3237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[58])))
	if base.B2i32(v3234 == int32(0))|base.B2i32(v3234 != v3237) != 0 {
		v3255 = v3234
		v3256 = v3237
		goto L763
	} else {
		goto L764
	}
L760:
	;
	v3259 = int32(1)
	goto L761
L761:
	;
	if v3259 == int32(0) {
		goto L769
	} else {
		goto L770
	}
L762:
	;
	v3259 = v3255 - v3256
	goto L761
L763:
	;
	goto L762
L764:
	;
	v3240 = v3230
	v3241 = v3231
	goto L765
L765:
	;
	v3244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3241)+1)))
	v3245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3240)+1)))
	if v3245 == int32(0) {
		v3255 = v3245
		v3256 = v3244
		goto L763
	} else {
		goto L767
	}
L766:
	;
	v3255 = v3245
	v3256 = v3244
	goto L763
L767:
	;
	v3248 = int32(1)
	if v3245 == v3244 {
		v3240 = v3240 + v3248
		v3241 = v3241 + v3248
		goto L765
	} else {
		goto L768
	}
L768:
	;
	goto L766
L769:
	;
	v3263 = v3137 + int32(352)
	v3264 = F_rmtree(m, v3263)
	mBase = m.M
	v3265 = m.ExcPending
	if v3265 != 0 {
		goto L32
	} else {
		goto L772
	}
L770:
	;
	goto L771
L771:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+320)) = int32(_a_F_StartupXLOG_132)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+324)) = v3209
	v3291 = v3137 + int32(2448)
	v3295 = F_pg_sprintf(m, v3291, int32(_a_F_StartupXLOG_135), v3137+int32(320))
	mBase = m.M
	v3296 = m.ExcPending
	if v3296 != 0 {
		goto L32
	} else {
		goto L781
	}
L772:
	;
	if v3264 == int32(0) {
		goto L773
	} else {
		goto L774
	}
L773:
	;
	v3270 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
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
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L32
	} else {
		goto L780
	}
L776:
	;
	if v3270 == int32(0) {
		goto L750
	} else {
		goto L777
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137))) = v3263
	F_errmsg(m, int32(_a_F_StartupXLOG_137), v3137)
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L32
	} else {
		goto L778
	}
L778:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2229), int32(_a_F_StartupXLOG_134))
	mBase = m.M
	v3282 = m.ExcPending
	if v3282 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+304)) = v3291
	v3299 = v3137 + int32(1392)
	v3303 = F_pg_sprintf(m, v3299, int32(_a_F_StartupXLOG_138), v3137+int32(304))
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L32
	} else {
		goto L782
	}
L782:
	;
	v3305 = F_unlink(m, v3299)
	mBase = m.M
	if v3305 < int32(0) {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v3309 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3309 != int32(44) {
		goto L742
	} else {
		goto L786
	}
L784:
	;
	goto L785
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+272)) = v3137 + int32(2448)
	v3316 = v3137 + int32(1392)
	v3320 = F_pg_sprintf(m, v3316, int32(_a_F_StartupXLOG_139), v3137+int32(272))
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L32
	} else {
		goto L787
	}
L786:
	;
	goto L785
L787:
	;
	v3324 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3325 = m.ExcPending
	if v3325 != 0 {
		goto L32
	} else {
		goto L788
	}
L788:
	;
	if v3324 != 0 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+256)) = v3316
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_140), v3137+int32(256))
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L32
	} else {
		goto L792
	}
L790:
	;
	goto L791
L791:
	;
	v3338 = v3137 + int32(1392)
	v3340 = F_OpenTransientFile(m, v3338, int32(2))
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L32
	} else {
		goto L794
	}
L792:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2506), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3336 = m.ExcPending
	if v3336 != 0 {
		goto L32
	} else {
		goto L793
	}
L793:
	;
	goto L791
L794:
	;
	if v3340 < int32(0) {
		goto L741
	} else {
		goto L795
	}
L795:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3345))) = int32(167772207)
	v3350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v3350 != int32(1) {
		v3364 = int32(0)
		goto L797
	} else {
		goto L798
	}
L796:
	;
	if v3364 != 0 {
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
	v3355 = F_fsync(m, v3340)
	mBase = m.M
	if v3355 != int32(-1) {
		v3364 = v3355
		goto L797
	} else {
		goto L801
	}
L800:
	;
	v3364 = int32(-1)
	goto L797
L801:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v3359 == int32(27) {
		goto L799
	} else {
		goto L802
	}
L802:
	;
	goto L800
L803:
	;
	v3366 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3366))) = int32(0)
	v3369 = int32(_a_F_StartupXLOG_142)
	v3371 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	v3372 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v3371 + v3372
	F_fsync_fname(m, v3137+int32(2448), v3372)
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
		goto L32
	} else {
		goto L804
	}
L804:
	;
	v3380 = int32(_a_F_StartupXLOG_142)
	v3382 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v3382 - int32(1)
	v3386 = int32(_a_F_StartupXLOG_143)
	v3387 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3387))) = int32(167772206)
	v3392 = int32(16)
	v3393 = F_read(m, v3340, v3137+int32(3496), v3392)
	mBase = m.M
	v3395 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3395))) = int32(0)
	if v3393 != v3392 {
		goto L805
	} else {
		goto L806
	}
L805:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L32
	} else {
		goto L808
	}
L806:
	;
	goto L807
L807:
	;
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3496))
	if v3423 != int32(17112225) {
		goto L738
	} else {
		goto L813
	}
L808:
	;
	if v3393 < int32(0) {
		goto L739
	} else {
		goto L809
	}
L809:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L32
	} else {
		goto L810
	}
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+232)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+228)) = v3393
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+224)) = v3338
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v3137+int32(224))
	mBase = m.M
	v3417 = m.ExcPending
	if v3417 != 0 {
		goto L32
	} else {
		goto L811
	}
L811:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2552), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
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
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3504))
	if v3426 != int32(5) {
		goto L737
	} else {
		goto L814
	}
L814:
	;
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3508))
	if v3429 != int32(184) {
		goto L736
	} else {
		goto L815
	}
L815:
	;
	v3432 = int32(_a_F_StartupXLOG_143)
	v3433 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3433))) = int32(167772206)
	v3437 = F_read(m, v3340, v3159, int32(184))
	mBase = m.M
	v3439 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v3439))) = int32(0)
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3508))
	if v3442 != v3437 {
		goto L816
	} else {
		goto L817
	}
L816:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L32
	} else {
		goto L819
	}
L817:
	;
	goto L818
L818:
	;
	v3468 = F_CloseTransientFile(m, v3340)
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L32
	} else {
		goto L824
	}
L819:
	;
	if v3437 < int32(0) {
		goto L735
	} else {
		goto L820
	}
L820:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v3452 = m.ExcPending
	if v3452 != 0 {
		goto L32
	} else {
		goto L821
	}
L821:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+152)) = v3442
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+148)) = v3437
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+144)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v3137+int32(144))
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L32
	} else {
		goto L822
	}
L822:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2592), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
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
	if v3468 != 0 {
		goto L734
	} else {
		goto L825
	}
L825:
	;
	v3470 = int32(-1)
	v3472 = m.Env.Pgmem_crc32c(m, v3470, v3137+int32(3504), int32(192))
	mBase = m.M
	v3474 = v3472 ^ v3470
	v3475 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3500))
	if v3474 != v3475 {
		goto L733
	} else {
		goto L826
	}
L826:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3580))
	if v3477 != 0 {
		goto L827
	} else {
		goto L828
	}
L827:
	;
	v3479 = v3137 + int32(2448)
	v3480 = F_rmtree(m, v3479)
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L32
	} else {
		goto L831
	}
L828:
	;
	goto L829
L829:
	;
	v3504 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3576))
	if v3505 != 0 {
		goto L839
	} else {
		goto L840
	}
L830:
	;
	F_fsync_fname(m, int32(_a_F_StartupXLOG_132), int32(1))
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L32
	} else {
		goto L837
	}
L831:
	;
	if v3480 != 0 {
		goto L830
	} else {
		goto L832
	}
L832:
	;
	v3484 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		goto L32
	} else {
		goto L833
	}
L833:
	;
	if v3484 == int32(0) {
		goto L830
	} else {
		goto L834
	}
L834:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+80)) = v3479
	F_errmsg(m, int32(_a_F_StartupXLOG_137), v3137+int32(80))
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L32
	} else {
		goto L835
	}
L835:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2622), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
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
	v3541 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if v3541 <= int32(0) {
		goto L730
	} else {
		goto L851
	}
L839:
	;
	if v3504 <= int32(1) {
		goto L732
	} else {
		goto L842
	}
L840:
	;
	goto L841
L841:
	;
	if v3504 <= int32(0) {
		goto L731
	} else {
		goto L850
	}
L842:
	;
	v3509 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v3509 != int32(1) {
		goto L838
	} else {
		goto L843
	}
L843:
	;
	v3513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v3513&int32(1) != 0 {
		goto L838
	} else {
		goto L844
	}
L844:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L32
	} else {
		goto L845
	}
L845:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L32
	} else {
		goto L846
	}
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+64)) = v3159
	F_errmsg(m, int32(_a_F_StartupXLOG_145), v3137-int32(-64))
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L32
	} else {
		goto L847
	}
L847:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_146), int32(0))
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L32
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2661), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
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
	v3546 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v3553 = int32(0)
	goto L852
L852:
	;
	v3583 = v3546 + v3553*int32(288)
	v3584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3583)+4)))
	if v3584 != 0 {
		goto L854
	} else {
		goto L855
	}
L853:
	;
	base.MemoryCopy(m, v3583+int32(24), v3159, int32(184))
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3584))
	*(*int32)(unsafe.Add(mBase, uint32(v3583)+16)) = v3592
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3588))
	*(*int32)(unsafe.Add(mBase, uint32(v3583)+20)) = v3594
	v3596 = *(*int64)(unsafe.Add(mBase, uint32(v3137)+3608))
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+264)) = v3596
	v3598 = *(*int64)(unsafe.Add(mBase, uint32(v3137)+3592))
	v3599 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+236)) = v3599
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+280)) = v3598
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+244)) = v3599
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+252)) = v3599
	v3606 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3583)+260)) = v3606
	*(*int32)(unsafe.Add(mBase, uint32(v3583)+8)) = v3606
	v3610 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3583)+4)) = uint8(v3610)
	v3615 = m.G0
	v3616 = int32(16)
	v3617 = v3615 - v3616
	m.G0 = v3617
	F_gettimeofday(m, v3617)
	mBase = m.M
	v3620 = *(*int64)(unsafe.Add(mBase, uint32(v3617)))
	v3621 = int64(*(*int32)(unsafe.Add(mBase, uint32(v3617)+8)))
	m.G0 = v3617 + v3616
	goto L858
L854:
	;
	v3586 = v3553 + int32(1)
	if v3541 != v3586 {
		v3553 = v3586
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
	v3630 = *(*int32)(unsafe.Add(mBase, uint32(v3583)+112))
	if v3630 != 0 {
		goto L750
	} else {
		goto L859
	}
L859:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+272)) = v3621 + v3620*int64(1000000) - int64(946684800000000)
	goto L750
L860:
	;
	if v3667 != 0 {
		v3168 = v3667
		goto L748
	} else {
		goto L861
	}
L861:
	;
	goto L749
L862:
	;
	v3706 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v3706 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L32
	} else {
		goto L866
	}
L864:
	;
	goto L865
L865:
	;
	m.G0 = v3137 + int32(3696)
	goto L729
L866:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
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
	v3722 = m.ExcPending
	if v3722 != 0 {
		goto L32
	} else {
		goto L869
	}
L869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+288)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v3137+int32(288))
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L32
	} else {
		goto L870
	}
L870:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2502), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
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
	v3741 = m.ExcPending
	if v3741 != 0 {
		goto L32
	} else {
		goto L873
	}
L873:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+16)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v3137+int32(16))
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L32
	} else {
		goto L874
	}
L874:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2518), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3754 = m.ExcPending
	if v3754 != 0 {
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
	v3760 = m.ExcPending
	if v3760 != 0 {
		goto L32
	} else {
		goto L877
	}
L877:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+240)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v3137+int32(240))
	mBase = m.M
	v3768 = m.ExcPending
	if v3768 != 0 {
		goto L32
	} else {
		goto L878
	}
L878:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2529), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+208)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v3137+int32(208))
	mBase = m.M
	v3783 = m.ExcPending
	if v3783 != 0 {
		goto L32
	} else {
		goto L881
	}
L881:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2546), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
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
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L32
	} else {
		goto L884
	}
L884:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+200)) = int32(17112225)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+196)) = v3423
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+192)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_150), v3137+int32(192))
	mBase = m.M
	v3806 = m.ExcPending
	if v3806 != 0 {
		goto L32
	} else {
		goto L885
	}
L885:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2560), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3811 = m.ExcPending
	if v3811 != 0 {
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
	v3818 = m.ExcPending
	if v3818 != 0 {
		goto L32
	} else {
		goto L888
	}
L888:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+180)) = v3426
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+176)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_151), v3137+int32(176))
	mBase = m.M
	v3827 = m.ExcPending
	if v3827 != 0 {
		goto L32
	} else {
		goto L889
	}
L889:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2567), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
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
	v3839 = m.ExcPending
	if v3839 != 0 {
		goto L32
	} else {
		goto L892
	}
L892:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+164)) = v3429
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+160)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_152), v3137+int32(160))
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L32
	} else {
		goto L893
	}
L893:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2574), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+128)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v3137+int32(128))
	mBase = m.M
	v3863 = m.ExcPending
	if v3863 != 0 {
		goto L32
	} else {
		goto L896
	}
L896:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2587), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3868 = m.ExcPending
	if v3868 != 0 {
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
	v3874 = m.ExcPending
	if v3874 != 0 {
		goto L32
	} else {
		goto L899
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+112)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v3137+int32(112))
	mBase = m.M
	v3882 = m.ExcPending
	if v3882 != 0 {
		goto L32
	} else {
		goto L900
	}
L900:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2598), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3887 = m.ExcPending
	if v3887 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+100)) = v3474
	v3893 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+3500))
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+104)) = v3893
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+96)) = v3137 + int32(1392)
	F_errmsg(m, int32(_a_F_StartupXLOG_154), v3137+int32(96))
	mBase = m.M
	v3902 = m.ExcPending
	if v3902 != 0 {
		goto L32
	} else {
		goto L903
	}
L903:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2610), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3907 = m.ExcPending
	if v3907 != 0 {
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
	v3914 = m.ExcPending
	if v3914 != 0 {
		goto L32
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+48)) = v3159
	F_errmsg(m, int32(_a_F_StartupXLOG_155), v3137+int32(48))
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L32
	} else {
		goto L907
	}
L907:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_156), int32(0))
	mBase = m.M
	v3924 = m.ExcPending
	if v3924 != 0 {
		goto L32
	} else {
		goto L908
	}
L908:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2647), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3929 = m.ExcPending
	if v3929 != 0 {
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
	v3936 = m.ExcPending
	if v3936 != 0 {
		goto L32
	} else {
		goto L911
	}
L911:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+32)) = v3159
	F_errmsg(m, int32(_a_F_StartupXLOG_157), v3137+int32(32))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L32
	} else {
		goto L912
	}
L912:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_158), int32(0))
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L32
	} else {
		goto L913
	}
L913:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2668), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v3951 = m.ExcPending
	if v3951 != 0 {
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
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L32
	} else {
		goto L916
	}
L916:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_160), int32(0))
	mBase = m.M
	v3997 = m.ExcPending
	if v3997 != 0 {
		goto L32
	} else {
		goto L917
	}
L917:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_133), int32(2716), int32(_a_F_StartupXLOG_141))
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
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
	v4007 = F_ReadDir(m, v4004, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4008 = m.ExcPending
	if v4008 != 0 {
		goto L32
	} else {
		goto L920
	}
L920:
	;
	if v4007 != 0 {
		goto L921
	} else {
		goto L922
	}
L921:
	;
	v4012 = v4007
	goto L924
L922:
	;
	goto L923
L923:
	;
	F_FreeDir(m, v4004)
	mBase = m.M
	v4103 = m.ExcPending
	if v4103 != 0 {
		goto L32
	} else {
		goto L937
	}
L924:
	;
	v4043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4012)+19)))
	if v4043 != int32(46) {
		goto L927
	} else {
		goto L928
	}
L925:
	;
	goto L923
L926:
	;
	v4066 = F_ReadDir(m, v4004, int32(_a_F_StartupXLOG_132))
	mBase = m.M
	v4067 = m.ExcPending
	if v4067 != 0 {
		goto L32
	} else {
		goto L935
	}
L927:
	;
	v4056 = v4012 + int32(19)
	v4058 = F_ReplicationSlotValidateName(m, v4056, int32(13))
	mBase = m.M
	v4059 = m.ExcPending
	if v4059 != 0 {
		goto L32
	} else {
		goto L932
	}
L928:
	;
	v4046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4012)+20)))
	if v4046 == int32(0) {
		goto L926
	} else {
		goto L929
	}
L929:
	;
	v4049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4012)+20)))
	if v4049 != int32(46) {
		goto L927
	} else {
		goto L930
	}
L930:
	;
	v4052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4012)+21)))
	if v4052 == int32(0) {
		goto L926
	} else {
		goto L931
	}
L931:
	;
	goto L927
L932:
	;
	if v4058 == int32(0) {
		goto L926
	} else {
		goto L933
	}
L933:
	;
	F_ReorderBufferCleanupSerializedTXNs(m, v4056)
	mBase = m.M
	v4063 = m.ExcPending
	if v4063 != 0 {
		goto L32
	} else {
		goto L934
	}
L934:
	;
	goto L926
L935:
	;
	if v4066 != 0 {
		v4012 = v4066
		goto L924
	} else {
		goto L936
	}
L936:
	;
	goto L925
L937:
	;
	v4105 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v4107 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v4108 = *(*int64)(unsafe.Add(mBase, uint32(v4107)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4105)+48)) = int64(base.Ui64(v4108)>>(uint(int64(15))%64)) & int64(131071)
	v4115 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v4116 = *(*int32)(unsafe.Add(mBase, uint32(v4115)+4))
	v4118 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v4115)))
	*(*int64)(unsafe.Add(mBase, uint32(v4118)+48)) = base.I64_extend_i32_u(int32(base.Ui32(v4119) >> (uint(int32(11)) % 32)))
	v4125 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v4127 = base.I32_div_u_s(v4116, int32(1636))
	*(*int64)(unsafe.Add(mBase, uint32(v4125)+48)) = base.I64_extend_i32_u(v4127)
	v4131 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4131)+200)))
	if v4132 == int32(1) {
		goto L938
	} else {
		goto L939
	}
L938:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L32
	} else {
		goto L941
	}
L939:
	;
	goto L940
L940:
	;
	v4137 = m.G0
	v4139 = v4137 - int32(176)
	m.G0 = v4139
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+172)) = int32(307747550)
	v4144 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	if v4144 == int32(0) {
		goto L950
	} else {
		goto L951
	}
L941:
	;
	goto L940
L942:
	;
	v4517 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v4519 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v4520 = *(*int32)(unsafe.Add(mBase, uint32(v4519)+16))
	if v4520 == int32(1) {
		goto L1020
	} else {
		goto L1021
	}
L943:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4501 = m.ExcPending
	if v4501 != 0 {
		goto L32
	} else {
		goto L1016
	}
L944:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4482 = m.ExcPending
	if v4482 != 0 {
		goto L32
	} else {
		goto L1012
	}
L945:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4466 = m.ExcPending
	if v4466 != 0 {
		goto L32
	} else {
		goto L1008
	}
L946:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L32
	} else {
		goto L1004
	}
L947:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L32
	} else {
		goto L1000
	}
L948:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4409 = m.ExcPending
	if v4409 != 0 {
		goto L32
	} else {
		goto L997
	}
L949:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4393 = m.ExcPending
	if v4393 != 0 {
		goto L32
	} else {
		goto L994
	}
L950:
	;
	m.G0 = v4139 + int32(176)
	goto L942
L951:
	;
	v4149 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L32
	} else {
		goto L952
	}
L952:
	;
	if v4149 != 0 {
		goto L953
	} else {
		goto L954
	}
L953:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_161), int32(0))
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L32
	} else {
		goto L956
	}
L954:
	;
	goto L955
L955:
	;
	v4162 = F_OpenTransientFile(m, int32(_a_F_StartupXLOG_162), int32(0))
	mBase = m.M
	v4163 = m.ExcPending
	if v4163 != 0 {
		goto L32
	} else {
		goto L958
	}
L956:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(745), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4159 = m.ExcPending
	if v4159 != 0 {
		goto L32
	} else {
		goto L957
	}
L957:
	;
	goto L955
L958:
	;
	if v4162 < int32(0) {
		goto L959
	} else {
		goto L960
	}
L959:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4167 == int32(44) {
		goto L950
	} else {
		goto L962
	}
L960:
	;
	goto L961
L961:
	;
	v4188 = int32(4)
	v4189 = F_read(m, v4162, v4139+int32(172), v4188)
	mBase = m.M
	if v4189 != v4188 {
		goto L967
	} else {
		goto L968
	}
L962:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L32
	} else {
		goto L963
	}
L963:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4175 = m.ExcPending
	if v4175 != 0 {
		goto L32
	} else {
		goto L964
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139))) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v4139)
	mBase = m.M
	v4180 = m.ExcPending
	if v4180 != 0 {
		goto L32
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(759), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
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
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L32
	} else {
		goto L970
	}
L968:
	;
	goto L969
L969:
	;
	v4220 = m.Env.Pgmem_crc32c(m, int32(-1), v4139+int32(172), int32(4))
	mBase = m.M
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v4139)+172))
	if v4221 != int32(307747550) {
		goto L948
	} else {
		goto L975
	}
L970:
	;
	if v4189 < int32(0) {
		goto L949
	} else {
		goto L971
	}
L971:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		goto L32
	} else {
		goto L972
	}
L972:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+136)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+132)) = v4189
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+128)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v4139+int32(128))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L32
	} else {
		goto L973
	}
L973:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(774), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
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
	v4227 = F_read(m, v4162, v4139+int32(152), int32(16))
	mBase = m.M
	if v4227 != int32(4) {
		goto L976
	} else {
		goto L977
	}
L976:
	;
	v4231 = int32(0)
	v4232 = v4220
	v4237 = v4227
	goto L979
L977:
	;
	v4316 = v4220
	goto L978
L978:
	;
	v4349 = *(*int32)(unsafe.Add(mBase, uint32(v4139)+152))
	v4351 = v4316 ^ int32(-1)
	if v4349 != v4351 {
		goto L944
	} else {
		goto L991
	}
L979:
	;
	if v4237 < int32(0) {
		goto L947
	} else {
		goto L981
	}
L980:
	;
	v4316 = v4272
	goto L978
L981:
	;
	if v4237 != int32(16) {
		goto L946
	} else {
		goto L982
	}
L982:
	;
	v4272 = m.Env.Pgmem_crc32c(m, v4232, v4139+int32(152), int32(16))
	mBase = m.M
	v4274 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[68]))
	if v4231 == v4274 {
		goto L945
	} else {
		goto L983
	}
L983:
	;
	v4277 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[69]))
	v4280 = v4277 + v4231*int32(56)
	v4281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4139)+152)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4280))) = uint16(v4281)
	v4283 = *(*int64)(unsafe.Add(mBase, uint32(v4139)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v4280)+8)) = v4283
	v4287 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4288 = m.ExcPending
	if v4288 != 0 {
		goto L32
	} else {
		goto L984
	}
L984:
	;
	if v4287 != 0 {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v4289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4139)+152)))
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+64)) = v4289
	v4291 = *(*int64)(unsafe.Add(mBase, uint32(v4139)+160))
	*(*uint32)(unsafe.Add(mBase, uint32(v4139)+72)) = uint32(v4291)
	v4294 = int64(base.Ui64(v4291) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4139)+68)) = uint32(v4294)
	F_errmsg(m, int32(_a_F_StartupXLOG_165), v4139-int32(-64))
	mBase = m.M
	v4300 = m.ExcPending
	if v4300 != 0 {
		goto L32
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v4312 = F_read(m, v4162, v4139+int32(152), int32(16))
	mBase = m.M
	if v4312 != int32(4) {
		v4231 = v4231 + int32(1)
		v4232 = v4272
		v4237 = v4312
		goto L979
	} else {
		goto L990
	}
L988:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(831), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4305 = m.ExcPending
	if v4305 != 0 {
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
	v4353 = F_CloseTransientFile(m, v4162)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L32
	} else {
		goto L992
	}
L992:
	;
	if v4353 != 0 {
		goto L943
	} else {
		goto L993
	}
L993:
	;
	goto L950
L994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+112)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v4139+int32(112))
	mBase = m.M
	v4400 = m.ExcPending
	if v4400 != 0 {
		goto L32
	} else {
		goto L995
	}
L995:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(769), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+100)) = int32(307747550)
	v4412 = *(*int32)(unsafe.Add(mBase, uint32(v4139)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+96)) = v4412
	F_errmsg(m, int32(_a_F_StartupXLOG_166), v4139+int32(96))
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		goto L32
	} else {
		goto L998
	}
L998:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(781), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
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
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L32
	} else {
		goto L1001
	}
L1001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+48)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v4139+int32(48))
	mBase = m.M
	v4436 = m.ExcPending
	if v4436 != 0 {
		goto L32
	} else {
		goto L1002
	}
L1002:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(805), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4441 = m.ExcPending
	if v4441 != 0 {
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
	v4447 = m.ExcPending
	if v4447 != 0 {
		goto L32
	} else {
		goto L1005
	}
L1005:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+88)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+84)) = v4237
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+80)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v4139+int32(80))
	mBase = m.M
	v4457 = m.ExcPending
	if v4457 != 0 {
		goto L32
	} else {
		goto L1006
	}
L1006:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(813), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4462 = m.ExcPending
	if v4462 != 0 {
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
	v4469 = m.ExcPending
	if v4469 != 0 {
		goto L32
	} else {
		goto L1009
	}
L1009:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_168), int32(0))
	mBase = m.M
	v4473 = m.ExcPending
	if v4473 != 0 {
		goto L32
	} else {
		goto L1010
	}
L1010:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(821), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4478 = m.ExcPending
	if v4478 != 0 {
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
	v4485 = m.ExcPending
	if v4485 != 0 {
		goto L32
	} else {
		goto L1013
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+36)) = v4349
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+32)) = v4351
	F_errmsg(m, int32(_a_F_StartupXLOG_169), v4139+int32(32))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L32
	} else {
		goto L1014
	}
L1014:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(840), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4497 = m.ExcPending
	if v4497 != 0 {
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
	v4503 = m.ExcPending
	if v4503 != 0 {
		goto L32
	} else {
		goto L1017
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4139)+16)) = int32(_a_F_StartupXLOG_162)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v4139+int32(16))
	mBase = m.M
	v4510 = m.ExcPending
	if v4510 != 0 {
		goto L32
	} else {
		goto L1018
	}
L1018:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_163), int32(846), int32(_a_F_StartupXLOG_164))
	mBase = m.M
	v4515 = m.ExcPending
	if v4515 != 0 {
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
	v4523 = *(*int64)(unsafe.Add(mBase, uint32(v4519)+128))
	v4525 = v4523
	goto L1022
L1021:
	;
	v4525 = int64(1000)
	goto L1022
L1022:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4517)+240)) = v4525
	v4528 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	F_restoreTimeLineHistoryFiles(m, v2895, v4528)
	mBase = m.M
	v4530 = m.ExcPending
	if v4530 != 0 {
		goto L32
	} else {
		goto L1023
	}
L1023:
	;
	v4532 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v4536 = F_LWLockAcquire(m, v4532+int32(2304), int32(0))
	mBase = m.M
	v4537 = m.ExcPending
	if v4537 != 0 {
		goto L32
	} else {
		goto L1024
	}
L1024:
	;
	v4539 = F_AllocateDir(m, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		goto L32
	} else {
		goto L1025
	}
L1025:
	;
	v4542 = F_ReadDir(m, v4539, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L32
	} else {
		goto L1026
	}
L1026:
	;
	if v4542 != 0 {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v4547 = v4542
	goto L1030
L1028:
	;
	goto L1029
L1029:
	;
	v4726 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v4726+int32(2304))
	mBase = m.M
	v4730 = m.ExcPending
	if v4730 != 0 {
		goto L32
	} else {
		goto L1060
	}
L1030:
	;
	v4579 = v4547 + int32(19)
	v4580 = F_strlen(m, v4579)
	mBase = m.M
	if v4580 != int32(16) {
		goto L1032
	} else {
		goto L1033
	}
L1031:
	;
	goto L1029
L1032:
	;
	v4689 = F_ReadDir(m, v4539, int32(_a_F_StartupXLOG_170))
	mBase = m.M
	v4690 = m.ExcPending
	if v4690 != 0 {
		goto L32
	} else {
		goto L1058
	}
L1033:
	;
	v4583 = int32(_a_F_StartupXLOG_171)
	v4587 = m.G0
	v4589 = v4587 - int32(32)
	v4590 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4589)+24)) = v4590
	*(*int64)(unsafe.Add(mBase, uint32(v4589)+16)) = v4590
	*(*int64)(unsafe.Add(mBase, uint32(v4589)+8)) = v4590
	*(*int64)(unsafe.Add(mBase, uint32(v4589))) = v4590
	v4598 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[71])))
	if v4598 == int32(0) {
		goto L1035
	} else {
		goto L1036
	}
L1034:
	;
	if v4666 != int32(16) {
		goto L1032
	} else {
		goto L1053
	}
L1035:
	;
	v4666 = int32(0)
	goto L1034
L1036:
	;
	goto L1037
L1037:
	;
	v4602 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[72])))
	if v4602 == int32(0) {
		goto L1038
	} else {
		goto L1039
	}
L1038:
	;
	v4606 = v4579
	goto L1041
L1039:
	;
	goto L1040
L1040:
	;
	v4616 = v4583
	v4617 = v4598
	goto L1044
L1041:
	;
	v4612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4606))))
	if v4612 == v4598 {
		v4606 = v4606 + int32(1)
		goto L1041
	} else {
		goto L1043
	}
L1042:
	;
	v4666 = v4606 - v4579
	goto L1034
L1043:
	;
	goto L1042
L1044:
	;
	v4624 = v4589 + int32(base.Ui32(v4617)>>(uint(int32(3))%32))&int32(28)
	v4625 = *(*int32)(unsafe.Add(mBase, uint32(v4624)))
	v4626 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4624))) = v4625 | v4626<<(uint(v4617)%32)
	v4630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4616)+1)))
	if v4630 != 0 {
		v4616 = v4616 + v4626
		v4617 = v4630
		goto L1044
	} else {
		goto L1046
	}
L1045:
	;
	v4633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4579))))
	if v4633 == int32(0) {
		v4656 = v4579
		goto L1047
	} else {
		goto L1048
	}
L1046:
	;
	goto L1045
L1047:
	;
	v4666 = v4656 - v4579
	goto L1034
L1048:
	;
	v4637 = v4579
	v4638 = v4633
	goto L1049
L1049:
	;
	v4646 = *(*int32)(unsafe.Add(mBase, uint32(v4589+int32(base.Ui32(v4638)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v4646)>>(uint(v4638)%32))&int32(1) == int32(0) {
		v4656 = v4637
		goto L1047
	} else {
		goto L1051
	}
L1050:
	;
	v4656 = v4654
	goto L1047
L1051:
	;
	v4652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4637)+1)))
	v4654 = v4637 + int32(1)
	if v4652 != 0 {
		v4637 = v4654
		v4638 = v4652
		goto L1049
	} else {
		goto L1052
	}
L1052:
	;
	goto L1050
L1053:
	;
	v4672 = F_strtox_2(m, v4579, int32(0), int32(16), int64(-1))
	mBase = m.M
	goto L1054
L1054:
	;
	v4676 = int32(0)
	v4678 = F_ProcessTwoPhaseBuffer(m, base.I32_wrap_i64(v4672), int64(0), int32(1), v4676, v4676)
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L32
	} else {
		goto L1055
	}
L1055:
	;
	if v4678 == int32(0) {
		goto L1032
	} else {
		goto L1056
	}
L1056:
	;
	v4682 = int64(0)
	F_PrepareRedoAdd(m, v4678, v4682, v4682, int32(0))
	mBase = m.M
	v4686 = m.ExcPending
	if v4686 != 0 {
		goto L32
	} else {
		goto L1057
	}
L1057:
	;
	goto L1032
L1058:
	;
	if v4689 != 0 {
		v4547 = v4689
		goto L1030
	} else {
		goto L1059
	}
L1059:
	;
	goto L1031
L1060:
	;
	F_FreeDir(m, v4539)
	mBase = m.M
	v4732 = m.ExcPending
	if v4732 != 0 {
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
	v5998 = int32(1)
	v5999 = v2894 & v5998
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[73])) = uint8(v5999)
	v6002 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v6002)+200)) = v2896
	*(*int64)(unsafe.Add(mBase, uint32(v6002)+152)) = v2896
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[74])) = uint8(v5999)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[75])) = v2896
	v6010 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v6010 == v5998 {
		goto L1351
	} else {
		goto L1352
	}
L1063:
	;
	v4735 = m.G0
	v4737 = v4735 - int32(48)
	m.G0 = v4737
	v4741 = F_unlink(m, int32(_a_F_StartupXLOG_172))
	mBase = m.M
	if v4741 != 0 {
		goto L1068
	} else {
		goto L1069
	}
L1064:
	;
	goto L1065
L1065:
	;
	v4890 = m.G0
	v4892 = v4890 - int32(528)
	m.G0 = v4892
	v4895 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[76]))
	v4898 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4899 = m.ExcPending
	if v4899 != 0 {
		goto L32
	} else {
		goto L1101
	}
L1066:
	;
	v4797 = m.G0
	v4798 = int32(16)
	v4799 = v4797 - v4798
	m.G0 = v4799
	F_gettimeofday(m, v4799)
	mBase = m.M
	v4802 = *(*int64)(unsafe.Add(mBase, uint32(v4799)))
	v4803 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4799)+8)))
	m.G0 = v4799 + v4798
	goto L1086
L1067:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v4790, int32(_a_F_StartupXLOG_174))
	mBase = m.M
	v4793 = m.ExcPending
	if v4793 != 0 {
		goto L32
	} else {
		goto L1085
	}
L1068:
	;
	v4743 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4743 == int32(44) {
		goto L1071
	} else {
		goto L1072
	}
L1069:
	;
	goto L1070
L1070:
	;
	v4778 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4779 = m.ExcPending
	if v4779 != 0 {
		goto L32
	} else {
		goto L1081
	}
L1071:
	;
	v4748 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		goto L32
	} else {
		goto L1074
	}
L1072:
	;
	goto L1073
L1073:
	;
	v4762 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4763 = m.ExcPending
	if v4763 != 0 {
		goto L32
	} else {
		goto L1077
	}
L1074:
	;
	if v4748 == int32(0) {
		goto L1066
	} else {
		goto L1075
	}
L1075:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4737)+16)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_175), v4737+int32(16))
	mBase = m.M
	v4758 = m.ExcPending
	if v4758 != 0 {
		goto L32
	} else {
		goto L1076
	}
L1076:
	;
	v4790 = int32(530)
	goto L1067
L1077:
	;
	if v4762 == int32(0) {
		goto L1066
	} else {
		goto L1078
	}
L1078:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		goto L32
	} else {
		goto L1079
	}
L1079:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4737)+32)) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_176), v4737+int32(32))
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L32
	} else {
		goto L1080
	}
L1080:
	;
	v4790 = int32(535)
	goto L1067
L1081:
	;
	if v4778 == int32(0) {
		goto L1066
	} else {
		goto L1082
	}
L1082:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4783 = m.ExcPending
	if v4783 != 0 {
		goto L32
	} else {
		goto L1083
	}
L1083:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4737))) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_177), v4737)
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L32
	} else {
		goto L1084
	}
L1084:
	;
	v4790 = int32(542)
	goto L1067
L1085:
	;
	goto L1066
L1086:
	;
	v4816 = int32(1)
	goto L1087
L1087:
	;
	if base.Ui32(v4816) <= base.Ui32(int32(12)) {
		goto L1091
	} else {
		goto L1092
	}
L1088:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v4886 = m.ExcPending
	if v4886 != 0 {
		goto L32
	} else {
		goto L1100
	}
L1089:
	;
	v4882 = v4816 + int32(1)
	if v4882 != int32(33) {
		v4816 = v4882
		goto L1087
	} else {
		goto L1099
	}
L1090:
	;
	v4871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4870))))
	if v4871&int32(1) == int32(0) {
		goto L1089
	} else {
		goto L1097
	}
L1091:
	;
	v4870 = v4816*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1090
L1092:
	;
	goto L1093
L1093:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4816-int32(24)) {
		goto L1089
	} else {
		goto L1094
	}
L1094:
	;
	v4858 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v4858 == int32(0) {
		goto L1089
	} else {
		goto L1095
	}
L1095:
	;
	v4866 = *(*int32)(unsafe.Add(mBase, uint32(v4858+v4816<<(uint(int32(2))%32)-int32(96))))
	if v4866 == int32(0) {
		goto L1089
	} else {
		goto L1096
	}
L1096:
	;
	v4870 = v4866
	goto L1090
L1097:
	;
	v4876 = *(*int32)(unsafe.Add(mBase, uint32(v4870)+60))
	m.T0[v4876].(func(*base.Module, int64))(m, v4803+v4802*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v4878 = m.ExcPending
	if v4878 != 0 {
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
	m.G0 = v4737 + int32(48)
	goto L1062
L1101:
	;
	if v4898 != 0 {
		goto L1102
	} else {
		goto L1103
	}
L1102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+432)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_179), v4892+int32(432))
	mBase = m.M
	v4906 = m.ExcPending
	if v4906 != 0 {
		goto L32
	} else {
		goto L1105
	}
L1103:
	;
	goto L1104
L1104:
	;
	v4914 = F_AllocateFile(m, int32(_a_F_StartupXLOG_172), int32(_a_F_StartupXLOG_60))
	mBase = m.M
	v4915 = m.ExcPending
	if v4915 != 0 {
		goto L32
	} else {
		goto L1108
	}
L1105:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1765), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v4911 = m.ExcPending
	if v4911 != 0 {
		goto L32
	} else {
		goto L1106
	}
L1106:
	;
	goto L1104
L1107:
	;
	m.G0 = v4892 + int32(528)
	goto L1062
L1108:
	;
	if v4914 == int32(0) {
		goto L1109
	} else {
		goto L1110
	}
L1109:
	;
	v4919 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v4919 == int32(44) {
		goto L1112
	} else {
		goto L1113
	}
L1110:
	;
	goto L1111
L1111:
	;
	v5035 = F_fread(m, v4892+int32(524), int32(1), int32(4), v4914)
	mBase = m.M
	v5036 = m.ExcPending
	if v5036 != 0 {
		goto L32
	} else {
		goto L1136
	}
L1112:
	;
	v4943 = m.G0
	v4944 = int32(16)
	v4945 = v4943 - v4944
	m.G0 = v4945
	F_gettimeofday(m, v4945)
	mBase = m.M
	v4948 = *(*int64)(unsafe.Add(mBase, uint32(v4945)))
	v4949 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4945)+8)))
	m.G0 = v4945 + v4944
	goto L1119
L1113:
	;
	v4924 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4925 = m.ExcPending
	if v4925 != 0 {
		goto L32
	} else {
		goto L1114
	}
L1114:
	;
	if v4924 == int32(0) {
		goto L1112
	} else {
		goto L1115
	}
L1115:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v4929 = m.ExcPending
	if v4929 != 0 {
		goto L32
	} else {
		goto L1116
	}
L1116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892))) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_181), v4892)
	mBase = m.M
	v4934 = m.ExcPending
	if v4934 != 0 {
		goto L32
	} else {
		goto L1117
	}
L1117:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1782), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v4939 = m.ExcPending
	if v4939 != 0 {
		goto L32
	} else {
		goto L1118
	}
L1118:
	;
	goto L1112
L1119:
	;
	v4959 = int32(1)
	goto L1120
L1120:
	;
	if base.Ui32(v4959) <= base.Ui32(int32(12)) {
		goto L1124
	} else {
		goto L1125
	}
L1121:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5030 = m.ExcPending
	if v5030 != 0 {
		goto L32
	} else {
		goto L1133
	}
L1122:
	;
	v5026 = v4959 + int32(1)
	if v5026 != int32(33) {
		v4959 = v5026
		goto L1120
	} else {
		goto L1132
	}
L1123:
	;
	v5016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5015))))
	if v5016&int32(1) == int32(0) {
		goto L1122
	} else {
		goto L1130
	}
L1124:
	;
	v5015 = v4959*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1123
L1125:
	;
	goto L1126
L1126:
	;
	if base.Ui32(int32(8)) < base.Ui32(v4959-int32(24)) {
		goto L1122
	} else {
		goto L1127
	}
L1127:
	;
	v5004 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5004 == int32(0) {
		goto L1122
	} else {
		goto L1128
	}
L1128:
	;
	v5012 = *(*int32)(unsafe.Add(mBase, uint32(v5004+v4959<<(uint(int32(2))%32)-int32(96))))
	if v5012 == int32(0) {
		goto L1122
	} else {
		goto L1129
	}
L1129:
	;
	v5015 = v5012
	goto L1123
L1130:
	;
	v5021 = *(*int32)(unsafe.Add(mBase, uint32(v5015)+60))
	m.T0[v5021].(func(*base.Module, int64))(m, v4949+v4948*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5023 = m.ExcPending
	if v5023 != 0 {
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
	v5906 = F_FreeFile(m, v4914)
	mBase = m.M
	v5907 = m.ExcPending
	if v5907 != 0 {
		goto L32
	} else {
		goto L1344
	}
L1135:
	;
	v5767 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5768 = m.ExcPending
	if v5768 != 0 {
		goto L32
	} else {
		goto L1323
	}
L1136:
	;
	if v5035 != int32(4) {
		goto L1137
	} else {
		goto L1138
	}
L1137:
	;
	v5041 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5042 = m.ExcPending
	if v5042 != 0 {
		goto L32
	} else {
		goto L1140
	}
L1138:
	;
	goto L1139
L1139:
	;
	v5054 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+524))
	if v5054 == int32(27638967) {
		goto L1144
	} else {
		goto L1145
	}
L1140:
	;
	if v5041 == int32(0) {
		goto L1135
	} else {
		goto L1141
	}
L1141:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_182), int32(0))
	mBase = m.M
	v5048 = m.ExcPending
	if v5048 != 0 {
		goto L32
	} else {
		goto L1142
	}
L1142:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1792), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5053 = m.ExcPending
	if v5053 != 0 {
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
	v5714 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5715 = m.ExcPending
	if v5715 != 0 {
		goto L32
	} else {
		goto L1319
	}
L1147:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5708, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5711 = m.ExcPending
	if v5711 != 0 {
		goto L32
	} else {
		goto L1318
	}
L1148:
	;
	v5692 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5693 = m.ExcPending
	if v5693 != 0 {
		goto L32
	} else {
		goto L1315
	}
L1149:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5680, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5687 = m.ExcPending
	if v5687 != 0 {
		goto L32
	} else {
		goto L1314
	}
L1150:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), v5672, int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5679 = m.ExcPending
	if v5679 != 0 {
		goto L32
	} else {
		goto L1313
	}
L1151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5656 = m.ExcPending
	if v5656 != 0 {
		goto L32
	} else {
		goto L1310
	}
L1152:
	;
	v5094 = F_do_getc(m, v4914)
	mBase = m.M
	v5095 = m.ExcPending
	if v5095 != 0 {
		goto L32
	} else {
		goto L1160
	}
L1153:
	;
	v5638 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5639 = m.ExcPending
	if v5639 != 0 {
		goto L32
	} else {
		goto L1307
	}
L1154:
	;
	v5515 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	v5520 = F_dshash_find_or_insert(m, v5515, v4892+int32(504), v4892+int32(523))
	mBase = m.M
	v5521 = m.ExcPending
	if v5521 != 0 {
		goto L32
	} else {
		goto L1278
	}
L1155:
	;
	v5452 = *(*int32)(unsafe.Add(mBase, uint32(v5451)+48))
	if v5452 == int32(0) {
		goto L1262
	} else {
		goto L1263
	}
L1156:
	;
	v5436 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5437 = m.ExcPending
	if v5437 != 0 {
		goto L32
	} else {
		goto L1258
	}
L1157:
	;
	v5415 = F_do_getc(m, v4914)
	mBase = m.M
	v5416 = m.ExcPending
	if v5416 != 0 {
		goto L32
	} else {
		goto L1252
	}
L1158:
	;
	v5222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	if v5222 != 0 {
		goto L1194
	} else {
		goto L1195
	}
L1159:
	;
	v5102 = F_fread(m, v4892+int32(436), int32(1), int32(4), v4914)
	mBase = m.M
	v5103 = m.ExcPending
	if v5103 != 0 {
		goto L32
	} else {
		goto L1161
	}
L1160:
	;
	switch v5094 - int32(69) {
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
	if v5102 != int32(4) {
		goto L1162
	} else {
		goto L1163
	}
L1162:
	;
	v5108 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5109 = m.ExcPending
	if v5109 != 0 {
		goto L32
	} else {
		goto L1165
	}
L1163:
	;
	goto L1164
L1164:
	;
	v5120 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+436))
	v5122 = v5120 - int32(24)
	v5126 = base.B2i32(base.Ui32(v5120-int32(1)) < base.Ui32(int32(12)))
	if v5126|base.B2i32(base.Ui32(v5122) < base.Ui32(int32(9))) == int32(0) {
		goto L1168
	} else {
		goto L1169
	}
L1165:
	;
	if v5108 == int32(0) {
		goto L1135
	} else {
		goto L1166
	}
L1166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+128)) = int32(70)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v4892+int32(128))
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L32
	} else {
		goto L1167
	}
L1167:
	;
	v5708 = int32(1822)
	goto L1147
L1168:
	;
	v5134 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5135 = m.ExcPending
	if v5135 != 0 {
		goto L32
	} else {
		goto L1171
	}
L1169:
	;
	goto L1170
L1170:
	;
	if v5126 == int32(0) {
		goto L1176
	} else {
		goto L1177
	}
L1171:
	;
	if v5134 == int32(0) {
		goto L1135
	} else {
		goto L1172
	}
L1172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+116)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+112)) = v5120
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_184), v4892+int32(112))
	mBase = m.M
	v5145 = m.ExcPending
	if v5145 != 0 {
		goto L32
	} else {
		goto L1173
	}
L1173:
	;
	v5708 = int32(1829)
	goto L1147
L1174:
	;
	v5197 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+16))
	v5200 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+20))
	v5201 = F_fread(m, v5196+v5197, int32(1), v5200, v4914)
	mBase = m.M
	v5202 = m.ExcPending
	if v5202 != 0 {
		goto L32
	} else {
		goto L1189
	}
L1175:
	;
	v5193 = *(*int32)(unsafe.Add(mBase, uint32(v5156+(v4895+int32(_a_F_StartupXLOG_185)))))
	v5194 = v5160
	v5196 = v5193
	goto L1174
L1176:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5122) {
		goto L1180
	} else {
		goto L1181
	}
L1177:
	;
	goto L1178
L1178:
	;
	v5182 = v5120 * int32(72)
	v5183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5182)+uint32(_c_F_StartupXLOG[80]))))
	if v5183&int32(1) == int32(0) {
		goto L1148
	} else {
		goto L1188
	}
L1179:
	;
	v5178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5160))))
	if v5178&int32(1) != 0 {
		goto L1175
	} else {
		goto L1187
	}
L1180:
	;
	v5165 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5166 = m.ExcPending
	if v5166 != 0 {
		goto L32
	} else {
		goto L1184
	}
L1181:
	;
	v5152 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5152 == int32(0) {
		goto L1180
	} else {
		goto L1182
	}
L1182:
	;
	v5156 = v5120 << (uint(int32(2)) % 32)
	v5160 = *(*int32)(unsafe.Add(mBase, uint32(v5152+v5156-int32(96))))
	if v5160 != 0 {
		goto L1179
	} else {
		goto L1183
	}
L1183:
	;
	goto L1180
L1184:
	;
	if v5165 == int32(0) {
		goto L1135
	} else {
		goto L1185
	}
L1185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+100)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+96)) = v5120
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_186), v4892+int32(96))
	mBase = m.M
	v5176 = m.ExcPending
	if v5176 != 0 {
		goto L32
	} else {
		goto L1186
	}
L1186:
	;
	v5708 = int32(1837)
	goto L1147
L1187:
	;
	goto L1148
L1188:
	;
	v5190 = *(*int32)(unsafe.Add(mBase, uint32(v5182)+uint32(_c_F_StartupXLOG[81])))
	v5194 = v5182 + int32(_a_F_StartupXLOG_178)
	v5196 = v4895 + v5190
	goto L1174
L1189:
	;
	if v5201 == v5200 {
		goto L1152
	} else {
		goto L1190
	}
L1190:
	;
	v5206 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5207 = m.ExcPending
	if v5207 != 0 {
		goto L32
	} else {
		goto L1191
	}
L1191:
	;
	if v5206 == int32(0) {
		goto L1135
	} else {
		goto L1192
	}
L1192:
	;
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+72)) = v5210
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+68)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+64)) = v5120
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_187), v4892-int32(-64))
	mBase = m.M
	v5219 = m.ExcPending
	if v5219 != 0 {
		goto L32
	} else {
		goto L1193
	}
L1193:
	;
	v5708 = int32(1863)
	goto L1147
L1194:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5224 = m.ExcPending
	if v5224 != 0 {
		goto L32
	} else {
		goto L1197
	}
L1195:
	;
	goto L1196
L1196:
	;
	if v5094 == int32(83) {
		goto L1198
	} else {
		goto L1199
	}
L1197:
	;
	goto L1196
L1198:
	;
	v5231 = F_fread(m, v4892+int32(504), int32(1), int32(16), v4914)
	mBase = m.M
	v5232 = m.ExcPending
	if v5232 != 0 {
		goto L32
	} else {
		goto L1201
	}
L1199:
	;
	goto L1200
L1200:
	;
	v5314 = F_fread(m, v4892+int32(500), int32(1), int32(4), v4914)
	mBase = m.M
	v5315 = m.ExcPending
	if v5315 != 0 {
		goto L32
	} else {
		goto L1222
	}
L1201:
	;
	if v5231 != int32(16) {
		goto L1202
	} else {
		goto L1203
	}
L1202:
	;
	v5237 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5238 = m.ExcPending
	if v5238 != 0 {
		goto L32
	} else {
		goto L1205
	}
L1203:
	;
	goto L1204
L1204:
	;
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+504))
	v5251 = v5249 - int32(24)
	v5255 = base.B2i32(base.Ui32(v5249-int32(1)) < base.Ui32(int32(12)))
	if v5255|base.B2i32(base.Ui32(v5251) < base.Ui32(int32(9))) == int32(0) {
		goto L1208
	} else {
		goto L1209
	}
L1205:
	;
	if v5237 == int32(0) {
		goto L1135
	} else {
		goto L1206
	}
L1206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+304)) = int32(83)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_188), v4892+int32(304))
	mBase = m.M
	v5247 = m.ExcPending
	if v5247 != 0 {
		goto L32
	} else {
		goto L1207
	}
L1207:
	;
	v5680 = int32(1883)
	goto L1149
L1208:
	;
	v5263 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5264 = m.ExcPending
	if v5264 != 0 {
		goto L32
	} else {
		goto L1211
	}
L1209:
	;
	goto L1210
L1210:
	;
	if base.Ui32(v5249-int32(1)) < base.Ui32(int32(12)) {
		goto L1154
	} else {
		goto L1214
	}
L1211:
	;
	if v5263 == int32(0) {
		goto L1135
	} else {
		goto L1212
	}
L1212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+288)) = int32(83)
	v5269 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+272)) = v5269
	v5271 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+280)) = v5271
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_189), v4892+int32(272))
	mBase = m.M
	v5277 = m.ExcPending
	if v5277 != 0 {
		goto L32
	} else {
		goto L1213
	}
L1213:
	;
	v5680 = int32(1891)
	goto L1149
L1214:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5251) {
		goto L1215
	} else {
		goto L1216
	}
L1215:
	;
	v5294 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5295 = m.ExcPending
	if v5295 != 0 {
		goto L32
	} else {
		goto L1219
	}
L1216:
	;
	v5282 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5282 == int32(0) {
		goto L1215
	} else {
		goto L1217
	}
L1217:
	;
	v5290 = *(*int32)(unsafe.Add(mBase, uint32(v5282+v5249<<(uint(int32(2))%32)-int32(96))))
	if v5290 != 0 {
		goto L1154
	} else {
		goto L1218
	}
L1218:
	;
	goto L1215
L1219:
	;
	if v5294 == int32(0) {
		goto L1135
	} else {
		goto L1220
	}
L1220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+256)) = int32(83)
	v5300 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+240)) = v5300
	v5302 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+248)) = v5302
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_190), v4892+int32(240))
	mBase = m.M
	v5308 = m.ExcPending
	if v5308 != 0 {
		goto L32
	} else {
		goto L1221
	}
L1221:
	;
	v5680 = int32(1899)
	goto L1149
L1222:
	;
	if v5314 != int32(4) {
		goto L1223
	} else {
		goto L1224
	}
L1223:
	;
	v5320 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5321 = m.ExcPending
	if v5321 != 0 {
		goto L32
	} else {
		goto L1226
	}
L1224:
	;
	goto L1225
L1225:
	;
	v5335 = F_fread(m, v4892+int32(436), int32(1), int32(64), v4914)
	mBase = m.M
	v5336 = m.ExcPending
	if v5336 != 0 {
		goto L32
	} else {
		goto L1229
	}
L1226:
	;
	if v5320 == int32(0) {
		goto L1135
	} else {
		goto L1227
	}
L1227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+400)) = v5094
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_183), v4892+int32(400))
	mBase = m.M
	v5329 = m.ExcPending
	if v5329 != 0 {
		goto L32
	} else {
		goto L1228
	}
L1228:
	;
	v5672 = int32(1912)
	goto L1150
L1229:
	;
	if v5335 != int32(64) {
		goto L1230
	} else {
		goto L1231
	}
L1230:
	;
	v5341 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5342 = m.ExcPending
	if v5342 != 0 {
		goto L32
	} else {
		goto L1233
	}
L1231:
	;
	goto L1232
L1232:
	;
	v5354 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+500))
	v5356 = v5354 - int32(24)
	v5358 = v5354 - int32(1)
	if base.B2i32(base.Ui32(v5358) < base.Ui32(int32(12)))|base.B2i32(base.Ui32(v5356) < base.Ui32(int32(9))) == int32(0) {
		goto L1236
	} else {
		goto L1237
	}
L1233:
	;
	if v5341 == int32(0) {
		goto L1135
	} else {
		goto L1234
	}
L1234:
	;
	v5345 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+384)) = v5345
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+388)) = v5094
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_191), v4892+int32(384))
	mBase = m.M
	v5352 = m.ExcPending
	if v5352 != 0 {
		goto L32
	} else {
		goto L1235
	}
L1235:
	;
	v5672 = int32(1918)
	goto L1150
L1236:
	;
	v5368 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5369 = m.ExcPending
	if v5369 != 0 {
		goto L32
	} else {
		goto L1239
	}
L1237:
	;
	goto L1238
L1238:
	;
	v5381 = base.B2i32(base.Ui32(int32(11)) < base.Ui32(v5358))
	if v5381 == int32(0) {
		goto L1242
	} else {
		goto L1243
	}
L1239:
	;
	if v5368 == int32(0) {
		goto L1135
	} else {
		goto L1240
	}
L1240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+372)) = v5094
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+368)) = v5354
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_184), v4892+int32(368))
	mBase = m.M
	v5378 = m.ExcPending
	if v5378 != 0 {
		goto L32
	} else {
		goto L1241
	}
L1241:
	;
	v5672 = int32(1924)
	goto L1150
L1242:
	;
	v5451 = v5354*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1155
L1243:
	;
	goto L1244
L1244:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5356) {
		goto L1245
	} else {
		goto L1246
	}
L1245:
	;
	v5403 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5404 = m.ExcPending
	if v5404 != 0 {
		goto L32
	} else {
		goto L1249
	}
L1246:
	;
	v5391 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5391 == int32(0) {
		goto L1245
	} else {
		goto L1247
	}
L1247:
	;
	v5399 = *(*int32)(unsafe.Add(mBase, uint32(v5391+v5354<<(uint(int32(2))%32)-int32(96))))
	if v5399 != 0 {
		v5451 = v5399
		goto L1155
	} else {
		goto L1248
	}
L1248:
	;
	goto L1245
L1249:
	;
	if v5403 == int32(0) {
		goto L1135
	} else {
		goto L1250
	}
L1250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+356)) = v5094
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+352)) = v5354
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_186), v4892+int32(352))
	mBase = m.M
	v5413 = m.ExcPending
	if v5413 != 0 {
		goto L32
	} else {
		goto L1251
	}
L1251:
	;
	v5672 = int32(1932)
	goto L1150
L1252:
	;
	if v5415 == int32(-1) {
		goto L1134
	} else {
		goto L1253
	}
L1253:
	;
	v5421 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5422 = m.ExcPending
	if v5422 != 0 {
		goto L32
	} else {
		goto L1254
	}
L1254:
	;
	if v5421 == int32(0) {
		goto L1135
	} else {
		goto L1255
	}
L1255:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_192), int32(0))
	mBase = m.M
	v5428 = m.ExcPending
	if v5428 != 0 {
		goto L32
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2010), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5433 = m.ExcPending
	if v5433 != 0 {
		goto L32
	} else {
		goto L1257
	}
L1257:
	;
	goto L1135
L1258:
	;
	if v5436 == int32(0) {
		goto L1135
	} else {
		goto L1259
	}
L1259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+48)) = v5094
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_193), v4892+int32(48))
	mBase = m.M
	v5445 = m.ExcPending
	if v5445 != 0 {
		goto L32
	} else {
		goto L1260
	}
L1260:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2017), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5450 = m.ExcPending
	if v5450 != 0 {
		goto L32
	} else {
		goto L1261
	}
L1261:
	;
	goto L1135
L1262:
	;
	v5457 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5458 = m.ExcPending
	if v5458 != 0 {
		goto L32
	} else {
		goto L1265
	}
L1263:
	;
	goto L1264
L1264:
	;
	v5473 = m.T0[v5452].(func(*base.Module, int32, int32) int32)(m, v4892+int32(436), v4892+int32(504))
	mBase = m.M
	v5474 = m.ExcPending
	if v5474 != 0 {
		goto L32
	} else {
		goto L1268
	}
L1265:
	;
	if v5457 == int32(0) {
		goto L1135
	} else {
		goto L1266
	}
L1266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+324)) = v5094
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+320)) = v5354
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_194), v4892+int32(320))
	mBase = m.M
	v5467 = m.ExcPending
	if v5467 != 0 {
		goto L32
	} else {
		goto L1267
	}
L1267:
	;
	v5672 = int32(1939)
	goto L1150
L1268:
	;
	if v5473 != 0 {
		goto L1154
	} else {
		goto L1269
	}
L1269:
	;
	if base.Ui32(int32(11)) < base.Ui32(v5358) {
		goto L1270
	} else {
		goto L1271
	}
L1270:
	;
	v5476 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	v5482 = *(*int32)(unsafe.Add(mBase, uint32(v5476+v5354<<(uint(int32(2))%32)-int32(96))))
	v5487 = v5482
	goto L1272
L1271:
	;
	v5487 = v5354*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1272
L1272:
	;
	v5488 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5487)+20)))
	v5490 = F___fseeko_unlocked(m, v4914, v5488, int32(1))
	mBase = m.M
	v5491 = m.ExcPending
	if v5491 != 0 {
		goto L32
	} else {
		goto L1273
	}
L1273:
	;
	if v5490 == int32(0) {
		goto L1152
	} else {
		goto L1274
	}
L1274:
	;
	v5496 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5497 = m.ExcPending
	if v5497 != 0 {
		goto L32
	} else {
		goto L1275
	}
L1275:
	;
	if v5496 == int32(0) {
		goto L1135
	} else {
		goto L1276
	}
L1276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+344)) = v5094
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+340)) = v5354
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+336)) = v4892 + int32(436)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_195), v4892+int32(336))
	mBase = m.M
	v5509 = m.ExcPending
	if v5509 != 0 {
		goto L32
	} else {
		goto L1277
	}
L1277:
	;
	v5672 = int32(1949)
	goto L1150
L1278:
	;
	v5522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4892)+523)))
	if v5522 == int32(1) {
		goto L1279
	} else {
		goto L1280
	}
L1279:
	;
	v5526 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	F_dshash_release_lock(m, v5526, v5520)
	mBase = m.M
	v5528 = m.ExcPending
	if v5528 != 0 {
		goto L32
	} else {
		goto L1282
	}
L1280:
	;
	goto L1281
L1281:
	;
	v5546 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+504))
	v5547 = int32(0)
	v5548 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+20)) = v5548
	*(*uint8)(unsafe.Add(mBase, uint32(v5520)+16)) = uint8(v5547)
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+24)) = v5547
	v5555 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	if base.Ui32(v5546-v5548) <= base.Ui32(int32(11)) {
		goto L1287
	} else {
		goto L1288
	}
L1282:
	;
	v5531 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5532 = m.ExcPending
	if v5532 != 0 {
		goto L32
	} else {
		goto L1283
	}
L1283:
	;
	if v5531 == int32(0) {
		goto L1135
	} else {
		goto L1284
	}
L1284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+160)) = v5094
	v5536 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+144)) = v5536
	v5538 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+152)) = v5538
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_196), v4892+int32(144))
	mBase = m.M
	v5544 = m.ExcPending
	if v5544 != 0 {
		goto L32
	} else {
		goto L1285
	}
L1285:
	;
	v5680 = int32(1972)
	goto L1149
L1286:
	;
	v5585 = *(*int32)(unsafe.Add(mBase, uint32(v5584)+4))
	v5587 = F_dsa_allocate_extended(m, v5555, v5585, int32(6))
	mBase = m.M
	v5588 = m.ExcPending
	if v5588 != 0 {
		goto L32
	} else {
		goto L1293
	}
L1287:
	;
	v5584 = v5546*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1286
L1288:
	;
	goto L1289
L1289:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5546-int32(24)) {
		v5582 = int32(0)
		goto L1290
	} else {
		goto L1291
	}
L1290:
	;
	v5584 = v5582
	goto L1286
L1291:
	;
	v5570 = int32(0)
	v5572 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5572 == v5570 {
		v5582 = v5570
		goto L1290
	} else {
		goto L1292
	}
L1292:
	;
	v5580 = *(*int32)(unsafe.Add(mBase, uint32(v5572+v5546<<(uint(int32(2))%32)-int32(96))))
	v5582 = v5580
	goto L1290
L1293:
	;
	if v5587 != 0 {
		goto L1294
	} else {
		goto L1295
	}
L1294:
	;
	v5590 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[82]))
	v5591 = F_dsa_get_address(m, v5590, v5587)
	mBase = m.M
	v5592 = m.ExcPending
	if v5592 != 0 {
		goto L32
	} else {
		goto L1297
	}
L1295:
	;
	v5604 = v5547
	goto L1296
L1296:
	;
	v5606 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[78]))
	F_dshash_release_lock(m, v5606, v5520)
	mBase = m.M
	v5608 = m.ExcPending
	if v5608 != 0 {
		goto L32
	} else {
		goto L1299
	}
L1297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5591))) = int32(-559038737)
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+28)) = v5587
	v5597 = v5591 + int32(4)
	v5598 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v5597))) = uint16(v5598)
	*(*int32)(unsafe.Add(mBase, uint32(v5597)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v5597)+8)) = int64(-1)
	goto L1298
L1298:
	;
	v5604 = v5591
	goto L1296
L1299:
	;
	if v5604 == int32(0) {
		goto L1151
	} else {
		goto L1300
	}
L1300:
	;
	v5611 = *(*int32)(unsafe.Add(mBase, uint32(v4892)+504))
	if base.Ui32(v5611-int32(1)) <= base.Ui32(int32(11)) {
		goto L1302
	} else {
		goto L1303
	}
L1301:
	;
	v5629 = *(*int32)(unsafe.Add(mBase, uint32(v5628)+16))
	v5632 = *(*int32)(unsafe.Add(mBase, uint32(v5628)+20))
	v5633 = F_fread(m, v5604+v5629, int32(1), v5632, v4914)
	mBase = m.M
	v5634 = m.ExcPending
	if v5634 != 0 {
		goto L32
	} else {
		goto L1305
	}
L1302:
	;
	v5628 = v5611*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1301
L1303:
	;
	goto L1304
L1304:
	;
	v5621 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	v5627 = *(*int32)(unsafe.Add(mBase, uint32(v5621+v5611<<(uint(int32(2))%32)-int32(96))))
	v5628 = v5627
	goto L1301
L1305:
	;
	if v5633 == v5632 {
		goto L1152
	} else {
		goto L1306
	}
L1306:
	;
	goto L1153
L1307:
	;
	if v5638 == int32(0) {
		goto L1135
	} else {
		goto L1308
	}
L1308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+224)) = v5094
	v5643 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+208)) = v5643
	v5645 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+216)) = v5645
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_197), v4892+int32(208))
	mBase = m.M
	v5651 = m.ExcPending
	if v5651 != 0 {
		goto L32
	} else {
		goto L1309
	}
L1309:
	;
	v5680 = int32(1996)
	goto L1149
L1310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+192)) = v5094
	v5658 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+176)) = v5658
	v5660 = *(*int64)(unsafe.Add(mBase, uint32(v4892)+512))
	*(*int64)(unsafe.Add(mBase, uint32(v4892)+184)) = v5660
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_198), v4892+int32(176))
	mBase = m.M
	v5666 = m.ExcPending
	if v5666 != 0 {
		goto L32
	} else {
		goto L1311
	}
L1311:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1987), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5671 = m.ExcPending
	if v5671 != 0 {
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
	if v5692 == int32(0) {
		goto L1135
	} else {
		goto L1316
	}
L1316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+84)) = int32(70)
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+80)) = v5120
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_199), v4892+int32(80))
	mBase = m.M
	v5703 = m.ExcPending
	if v5703 != 0 {
		goto L32
	} else {
		goto L1317
	}
L1317:
	;
	v5708 = int32(1844)
	goto L1147
L1318:
	;
	goto L1135
L1319:
	;
	if v5714 == int32(0) {
		goto L1135
	} else {
		goto L1320
	}
L1320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+420)) = int32(27638967)
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+416)) = v5054
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_200), v4892+int32(416))
	mBase = m.M
	v5725 = m.ExcPending
	if v5725 != 0 {
		goto L32
	} else {
		goto L1321
	}
L1321:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(1799), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5730 = m.ExcPending
	if v5730 != 0 {
		goto L32
	} else {
		goto L1322
	}
L1322:
	;
	goto L1135
L1323:
	;
	if v5767 != 0 {
		goto L1324
	} else {
		goto L1325
	}
L1324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+32)) = int32(_a_F_StartupXLOG_172)
	F_errmsg(m, int32(_a_F_StartupXLOG_201), v4892+int32(32))
	mBase = m.M
	v5775 = m.ExcPending
	if v5775 != 0 {
		goto L32
	} else {
		goto L1327
	}
L1325:
	;
	goto L1326
L1326:
	;
	v5784 = m.G0
	v5785 = int32(16)
	v5786 = v5784 - v5785
	m.G0 = v5786
	F_gettimeofday(m, v5786)
	mBase = m.M
	v5789 = *(*int64)(unsafe.Add(mBase, uint32(v5786)))
	v5790 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5786)+8)))
	m.G0 = v5786 + v5785
	goto L1329
L1327:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2032), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5780 = m.ExcPending
	if v5780 != 0 {
		goto L32
	} else {
		goto L1328
	}
L1328:
	;
	goto L1326
L1329:
	;
	v5800 = int32(1)
	goto L1330
L1330:
	;
	if base.Ui32(v5800) <= base.Ui32(int32(12)) {
		goto L1334
	} else {
		goto L1335
	}
L1331:
	;
	F_pgstat_drop_all_entries(m)
	mBase = m.M
	v5871 = m.ExcPending
	if v5871 != 0 {
		goto L32
	} else {
		goto L1343
	}
L1332:
	;
	v5867 = v5800 + int32(1)
	if v5867 != int32(33) {
		v5800 = v5867
		goto L1330
	} else {
		goto L1342
	}
L1333:
	;
	v5857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5856))))
	if v5857&int32(1) == int32(0) {
		goto L1332
	} else {
		goto L1340
	}
L1334:
	;
	v5856 = v5800*int32(72) + int32(_a_F_StartupXLOG_178)
	goto L1333
L1335:
	;
	goto L1336
L1336:
	;
	if base.Ui32(int32(8)) < base.Ui32(v5800-int32(24)) {
		goto L1332
	} else {
		goto L1337
	}
L1337:
	;
	v5845 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[77]))
	if v5845 == int32(0) {
		goto L1332
	} else {
		goto L1338
	}
L1338:
	;
	v5853 = *(*int32)(unsafe.Add(mBase, uint32(v5845+v5800<<(uint(int32(2))%32)-int32(96))))
	if v5853 == int32(0) {
		goto L1332
	} else {
		goto L1339
	}
L1339:
	;
	v5856 = v5853
	goto L1333
L1340:
	;
	v5862 = *(*int32)(unsafe.Add(mBase, uint32(v5856)+60))
	m.T0[v5862].(func(*base.Module, int64))(m, v5790+v5789*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v5864 = m.ExcPending
	if v5864 != 0 {
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
	v5910 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v5911 = m.ExcPending
	if v5911 != 0 {
		goto L32
	} else {
		goto L1345
	}
L1345:
	;
	if v5910 != 0 {
		goto L1346
	} else {
		goto L1347
	}
L1346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4892)+16)) = int32(_a_F_StartupXLOG_172)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_202), v4892+int32(16))
	mBase = m.M
	v5918 = m.ExcPending
	if v5918 != 0 {
		goto L32
	} else {
		goto L1349
	}
L1347:
	;
	goto L1348
L1348:
	;
	v5925 = F_unlink(m, int32(_a_F_StartupXLOG_172))
	mBase = m.M
	goto L1107
L1349:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_173), int32(2025), int32(_a_F_StartupXLOG_180))
	mBase = m.M
	v5923 = m.ExcPending
	if v5923 != 0 {
		goto L32
	} else {
		goto L1350
	}
L1350:
	;
	goto L1348
L1351:
	;
	v6013 = *(*int32)(unsafe.Add(mBase, uint32(v6002)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v6002)+440)) = int32(1)
	if v6013 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L1352:
	;
	goto L1353
L1353:
	;
	v9202 = m.G0
	v9204 = v9202 - int32(272)
	m.G0 = v9204
	v9207 = F_palloc(m, int32(72))
	mBase = m.M
	v9208 = m.ExcPending
	if v9208 != 0 {
		goto L32
	} else {
		goto L2012
	}
L1354:
	;
	v6017 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v6017+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_203), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L32
	} else {
		goto L1357
	}
L1355:
	;
	goto L1356
L1356:
	;
	v6026 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	v6028 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v6028)+440)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6028)+316)) = v6026
	v6033 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	v6035 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v6033, v6035)
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L32
	} else {
		goto L1358
	}
L1357:
	;
	goto L1356
L1358:
	;
	v6038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4093)))
	if v6038 == int32(1) {
		goto L1359
	} else {
		goto L1360
	}
L1359:
	;
	v6041 = int32(_a_F_StartupXLOG_204)
	v6042 = F_unlink(m, v6041)
	mBase = m.M
	v6046 = F_durable_rename(m, int32(_a_F_StartupXLOG_59), v6041, int32(22))
	mBase = m.M
	v6047 = m.ExcPending
	if v6047 != 0 {
		goto L32
	} else {
		goto L1362
	}
L1360:
	;
	goto L1361
L1361:
	;
	v6048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4094)))
	if v6048 == int32(1) {
		goto L1363
	} else {
		goto L1364
	}
L1362:
	;
	goto L1361
L1363:
	;
	v6051 = int32(_a_F_StartupXLOG_91)
	v6052 = F_unlink(m, v6051)
	mBase = m.M
	v6056 = F_durable_rename(m, int32(_a_F_StartupXLOG_44), v6051, int32(22))
	mBase = m.M
	v6057 = m.ExcPending
	if v6057 != 0 {
		goto L32
	} else {
		goto L1366
	}
L1364:
	;
	goto L1365
L1365:
	;
	v6060 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])))
	if v6060 == int32(1) {
		goto L1367
	} else {
		goto L1368
	}
L1366:
	;
	goto L1365
L1367:
	;
	v6064 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v6065 = *(*int64)(unsafe.Add(mBase, uint32(v6064)+136))
	v6067 = v6065
	goto L1369
L1368:
	;
	v6067 = int64(0)
	goto L1369
L1369:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[83])) = v6067
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v6070 = m.ExcPending
	if v6070 != 0 {
		goto L32
	} else {
		goto L1370
	}
L1370:
	;
	F_ResetUnloggedRelations(m, int32(1))
	mBase = m.M
	v6073 = m.ExcPending
	if v6073 != 0 {
		goto L32
	} else {
		goto L1371
	}
L1371:
	;
	v6074 = m.G0
	v6076 = v6074 - int32(1072)
	m.G0 = v6076
	v6079 = F_AllocateDir(m, int32(_a_F_StartupXLOG_205))
	mBase = m.M
	v6080 = m.ExcPending
	if v6080 != 0 {
		goto L32
	} else {
		goto L1372
	}
L1372:
	;
	v6083 = F_ReadDirExtended(m, v6079, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6084 = m.ExcPending
	if v6084 != 0 {
		goto L32
	} else {
		goto L1373
	}
L1373:
	;
	if v6083 != 0 {
		goto L1374
	} else {
		goto L1375
	}
L1374:
	;
	v6085 = v6083
	goto L1377
L1375:
	;
	goto L1376
L1376:
	;
	F_FreeDir(m, v6079)
	mBase = m.M
	v6202 = m.ExcPending
	if v6202 != 0 {
		goto L32
	} else {
		goto L1394
	}
L1377:
	;
	v6119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6085)+19)))
	if v6119 != int32(46) {
		goto L1380
	} else {
		goto L1381
	}
L1378:
	;
	goto L1376
L1379:
	;
	v6165 = F_ReadDirExtended(m, v6079, int32(_a_F_StartupXLOG_205), int32(15))
	mBase = m.M
	v6166 = m.ExcPending
	if v6166 != 0 {
		goto L32
	} else {
		goto L1392
	}
L1380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6076)+16)) = v6085 + int32(19)
	v6135 = v6076 + int32(32)
	v6140 = F_pg_snprintf(m, v6135, int32(1037), int32(_a_F_StartupXLOG_206), v6076+int32(16))
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L32
	} else {
		goto L1385
	}
L1381:
	;
	v6122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6085)+20)))
	if v6122 == int32(0) {
		goto L1379
	} else {
		goto L1382
	}
L1382:
	;
	v6125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6085)+20)))
	if v6125 != int32(46) {
		goto L1380
	} else {
		goto L1383
	}
L1383:
	;
	v6128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6085)+21)))
	if v6128 == int32(0) {
		goto L1379
	} else {
		goto L1384
	}
L1384:
	;
	goto L1380
L1385:
	;
	v6142 = F_unlink(m, v6135)
	mBase = m.M
	if v6142 == int32(0) {
		goto L1379
	} else {
		goto L1386
	}
L1386:
	;
	v6147 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6148 = m.ExcPending
	if v6148 != 0 {
		goto L32
	} else {
		goto L1387
	}
L1387:
	;
	if v6147 == int32(0) {
		goto L1379
	} else {
		goto L1388
	}
L1388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v6152 = m.ExcPending
	if v6152 != 0 {
		goto L32
	} else {
		goto L1389
	}
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6076))) = v6135
	F_errmsg(m, int32(_a_F_StartupXLOG_147), v6076)
	mBase = m.M
	v6156 = m.ExcPending
	if v6156 != 0 {
		goto L32
	} else {
		goto L1390
	}
L1390:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_207), int32(1609), int32(_a_F_StartupXLOG_208))
	mBase = m.M
	v6161 = m.ExcPending
	if v6161 != 0 {
		goto L32
	} else {
		goto L1391
	}
L1391:
	;
	goto L1379
L1392:
	;
	if v6165 != 0 {
		v6085 = v6165
		goto L1377
	} else {
		goto L1393
	}
L1393:
	;
	goto L1378
L1394:
	;
	m.G0 = v6076 + int32(1072)
	v6207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v6207 != int32(1) {
		goto L1395
	} else {
		goto L1396
	}
L1395:
	;
	v6430 = m.G0
	v6432 = v6430 - int32(848)
	m.G0 = v6432
	v6435 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6436 = *(*int32)(unsafe.Add(mBase, uint32(v6435)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v6435)+96)) = int32(1)
	if v6436 != 0 {
		goto L1426
	} else {
		goto L1427
	}
L1396:
	;
	v6211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v6211&int32(1) == int32(0) {
		goto L1395
	} else {
		goto L1397
	}
L1397:
	;
	v6218 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v6219 = m.ExcPending
	if v6219 != 0 {
		goto L32
	} else {
		goto L1398
	}
L1398:
	;
	if v6218 != 0 {
		goto L1399
	} else {
		goto L1400
	}
L1399:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_209), int32(0))
	mBase = m.M
	v6223 = m.ExcPending
	if v6223 != 0 {
		goto L32
	} else {
		goto L1402
	}
L1400:
	;
	goto L1401
L1401:
	;
	v6229 = m.G0
	v6231 = v6229 + int32(-64)
	m.G0 = v6231
	*(*int64)(unsafe.Add(mBase, uint32(v6231)+24)) = int64(68719476748)
	v6239 = v6229 + int32(-56)
	v6241 = F_hash_create(m, int32(_a_F_StartupXLOG_210), int32(64), v6239, int32(40))
	mBase = m.M
	v6242 = m.ExcPending
	if v6242 != 0 {
		goto L32
	} else {
		goto L1404
	}
L1402:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_211), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v6228 = m.ExcPending
	if v6228 != 0 {
		goto L32
	} else {
		goto L1403
	}
L1403:
	;
	goto L1401
L1404:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[84])) = v6241
	*(*int64)(unsafe.Add(mBase, uint32(v6231)+24)) = int64(34359738372)
	v6250 = F_hash_create(m, int32(_a_F_StartupXLOG_212), int32(64), v6239, int32(40))
	mBase = m.M
	v6251 = m.ExcPending
	if v6251 != 0 {
		goto L32
	} else {
		goto L1405
	}
L1405:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[85])) = v6250
	F_SharedInvalBackendInit(m, int32(1))
	mBase = m.M
	v6255 = m.ExcPending
	if v6255 != 0 {
		goto L32
	} else {
		goto L1406
	}
L1406:
	;
	v6257 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[86]))
	v6259 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[87]))
	*(*int32)(unsafe.Add(mBase, uint32(v6257)+52)) = v6259
	*(*int32)(unsafe.Add(mBase, uint32(v6231)+56)) = v6259
	v6263 = int32(_a_F_StartupXLOG_213)
	v6264 = int32(1)
	v6266 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[88]))
	if base.Ui32(v6266) <= base.Ui32(v6264) {
		goto L1408
	} else {
		goto L1409
	}
L1407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6231)+60)) = v6269
	v6274 = *(*int64)(unsafe.Add(mBase, uint32(v6231)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v6231))) = v6274
	F_VirtualXactLockTableInsert(m, v6231)
	mBase = m.M
	v6277 = m.ExcPending
	if v6277 != 0 {
		goto L32
	} else {
		goto L1411
	}
L1408:
	;
	v6269 = v6264
	goto L1410
L1409:
	;
	v6269 = v6266
	goto L1410
L1410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[88])) = v6269 + int32(1)
	goto L1407
L1411:
	;
	v6279 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89])) = v6279
	m.G0 = v6231 - int32(-64)
	v6284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4095)))
	if v6284 == v6279 {
		goto L1412
	} else {
		goto L1413
	}
L1412:
	;
	v6291 = F_PrescanPreparedTransactions(m, v39+int32(_a_F_StartupXLOG_4), v39+int32(_a_F_StartupXLOG_3))
	mBase = m.M
	v6292 = m.ExcPending
	if v6292 != 0 {
		goto L32
	} else {
		goto L1415
	}
L1413:
	;
	v6293 = v2893
	goto L1414
L1414:
	;
	v6295 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v6296 = *(*int32)(unsafe.Add(mBase, uint32(v6295)+8))
	v6300 = v6296
	goto L1416
L1415:
	;
	v6293 = v6291
	goto L1414
L1416:
	;
	v6332 = v6300 - int32(1)
	if base.Ui32(v6332) < base.Ui32(int32(3)) {
		v6300 = v6332
		goto L1416
	} else {
		goto L1418
	}
L1417:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[90])) = v6332
	F_StartupSUBTRANS(m, v6293)
	mBase = m.M
	v6338 = m.ExcPending
	if v6338 != 0 {
		goto L32
	} else {
		goto L1419
	}
L1418:
	;
	goto L1417
L1419:
	;
	v6339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4095)))
	if v6339 != int32(1) {
		goto L1395
	} else {
		goto L1420
	}
L1420:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v6343 = m.ExcPending
	if v6343 != 0 {
		goto L32
	} else {
		goto L1421
	}
L1421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v6293
	v6345 = base.I32_wrap_i64(v2905)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v6345
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[93]))) = int64(8589934592)
	v6349 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[94])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v6349
	v6354 = v6345
	goto L1422
L1422:
	;
	v6386 = v6354 - int32(1)
	if base.Ui32(v6386) < base.Ui32(int32(3)) {
		v6354 = v6386
		goto L1422
	} else {
		goto L1424
	}
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[95]))) = v6386
	v6390 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[96])))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[97]))) = v6390
	F_ProcArrayApplyRecoveryInfo(m, v39+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v6395 = m.ExcPending
	if v6395 != 0 {
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
	v6440 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v6440+int32(96), int32(_a_F_StartupXLOG_50), int32(1682), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6447 = m.ExcPending
	if v6447 != 0 {
		goto L32
	} else {
		goto L1429
	}
L1427:
	;
	goto L1428
L1428:
	;
	v6449 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	v6451 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v6449) < base.Ui64(v6451) {
		goto L1431
	} else {
		goto L1432
	}
L1429:
	;
	goto L1428
L1430:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6467)+32)) = v6468
	v6471 = *(*int32)(unsafe.Add(mBase, uint32(v6469)))
	v6472 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6467)+64)) = v6472
	*(*int32)(unsafe.Add(mBase, uint32(v6467)+56)) = v6471
	*(*int64)(unsafe.Add(mBase, uint32(v6467)+48)) = v6468
	*(*int32)(unsafe.Add(mBase, uint32(v6467)+40)) = v6471
	*(*int64)(unsafe.Add(mBase, uint32(v6467)+72)) = v6472
	v6479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6467)+80)) = v6479
	*(*int32)(unsafe.Add(mBase, uint32(v6467)+96)) = v6479
	v6487 = m.G0
	v6488 = int32(16)
	v6489 = v6487 - v6488
	m.G0 = v6489
	F_gettimeofday(m, v6489)
	mBase = m.M
	v6492 = *(*int64)(unsafe.Add(mBase, uint32(v6489)))
	v6493 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6489)+8)))
	m.G0 = v6489 + v6488
	goto L1434
L1431:
	;
	v6454 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v6454)+24)) = int64(0)
	v6467 = v6454
	v6468 = v6449
	v6469 = int32(_a_F_StartupXLOG_215)
	goto L1430
L1432:
	;
	goto L1433
L1433:
	;
	v6459 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6461 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6462 = *(*int64)(unsafe.Add(mBase, uint32(v6461)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v6459)+24)) = v6462
	v6464 = *(*int64)(unsafe.Add(mBase, uint32(v6461)+40))
	v6467 = v6459
	v6468 = v6464
	v6469 = int32(_a_F_StartupXLOG_216)
	goto L1430
L1434:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[98])) = v6493 + v6492*int64(1000000) - int64(946684800000000)
	v6504 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v6504 == int32(1) {
		goto L1435
	} else {
		goto L1436
	}
L1435:
	;
	F_SendPostmasterSignal(m, int32(0))
	mBase = m.M
	v6509 = m.ExcPending
	if v6509 != 0 {
		goto L32
	} else {
		goto L1438
	}
L1436:
	;
	goto L1437
L1437:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v6511 = m.ExcPending
	if v6511 != 0 {
		goto L32
	} else {
		goto L1439
	}
L1438:
	;
	goto L1437
L1439:
	;
	v6513 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v6515 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[39]))
	v6517 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[35]))
	if base.Ui64(v6515) < base.Ui64(v6517) {
		goto L1449
	} else {
		goto L1450
	}
L1440:
	;
	goto L1353
L1441:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9154 = m.ExcPending
	if v9154 != 0 {
		goto L32
	} else {
		goto L2008
	}
L1442:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9141 = m.ExcPending
	if v9141 != 0 {
		goto L32
	} else {
		goto L2005
	}
L1443:
	;
	if v9100 != 0 {
		goto L2001
	} else {
		goto L2002
	}
L1444:
	;
	v8968 = int32(0)
	goto L1972
L1445:
	;
	v8895 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[99])))
	if v8895 == int32(0) {
		goto L1442
	} else {
		goto L1960
	}
L1446:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v7106
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v8801
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = int64(0)
	v8811 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v8811)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v8811)
	v8818 = F_errstart(m, int32(15), v8811)
	mBase = m.M
	v8819 = m.ExcPending
	if v8819 != 0 {
		goto L32
	} else {
		goto L1948
	}
L1447:
	;
	v8787 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8788 = m.ExcPending
	if v8788 != 0 {
		goto L32
	} else {
		goto L1944
	}
L1448:
	;
	F_getrusage(m, v6432+int32(384))
	mBase = m.M
	F_gettimeofday(m, v6432+int32(368))
	mBase = m.M
	goto L1463
L1449:
	;
	v6520 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[38]))
	F_XLogPrefetcherBeginRead(m, v6513, v6515)
	mBase = m.M
	v6522 = m.ExcPending
	if v6522 != 0 {
		goto L32
	} else {
		goto L1452
	}
L1450:
	;
	goto L1451
L1451:
	;
	v6558 = int32(0)
	v6562 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[36]))
	v6563 = F_ReadRecord(m, v6513, int32(15), v6558, v6562)
	mBase = m.M
	v6564 = m.ExcPending
	if v6564 != 0 {
		goto L32
	} else {
		goto L1461
	}
L1452:
	;
	v6524 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v6527 = F_ReadRecord(m, v6524, int32(23), int32(0), v6520)
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L32
	} else {
		goto L1453
	}
L1453:
	;
	v6529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6527)+17)))
	if v6529 == int32(0) {
		goto L1454
	} else {
		goto L1455
	}
L1454:
	;
	v6532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6527)+16)))
	if v6532&int32(240) == int32(224) {
		v6567 = v6520
		v6568 = v6527
		goto L1448
	} else {
		goto L1457
	}
L1455:
	;
	goto L1456
L1456:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6540 = m.ExcPending
	if v6540 != 0 {
		goto L32
	} else {
		goto L1458
	}
L1457:
	;
	goto L1456
L1458:
	;
	v6542 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6543 = *(*int64)(unsafe.Add(mBase, uint32(v6542)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+356)) = uint32(v6543)
	v6546 = int64(base.Ui64(v6543) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+352)) = uint32(v6546)
	F_errmsg(m, int32(_a_F_StartupXLOG_217), v6432+int32(352))
	mBase = m.M
	v6552 = m.ExcPending
	if v6552 != 0 {
		goto L32
	} else {
		goto L1459
	}
L1459:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1737), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6557 = m.ExcPending
	if v6557 != 0 {
		goto L32
	} else {
		goto L1460
	}
L1460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1461:
	;
	if v6563 == int32(0) {
		goto L1447
	} else {
		goto L1462
	}
L1462:
	;
	v6567 = v6562
	v6568 = v6563
	goto L1448
L1463:
	;
	v6577 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v6577)
	v6581 = int32(0)
	goto L1464
L1464:
	;
	v6615 = v6581 << (uint(int32(5)) % 32)
	v6618 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+uint32(_c_F_StartupXLOG[106])))
	if v6618 == int32(0) {
		goto L1466
	} else {
		goto L1467
	}
L1465:
	;
	v6642 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6643 = m.ExcPending
	if v6643 != 0 {
		goto L32
	} else {
		goto L1475
	}
L1466:
	;
	v6627 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+uint32(_c_F_StartupXLOG[107])))
	if v6627 == int32(0) {
		goto L1470
	} else {
		goto L1471
	}
L1467:
	;
	v6621 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+uint32(_c_F_StartupXLOG[108])))
	if v6621 == int32(0) {
		goto L1466
	} else {
		goto L1468
	}
L1468:
	;
	m.T0[v6621].(func(*base.Module))(m)
	mBase = m.M
	v6625 = m.ExcPending
	if v6625 != 0 {
		goto L32
	} else {
		goto L1469
	}
L1469:
	;
	goto L1466
L1470:
	;
	v6637 = v6581 + int32(2)
	if v6637 != int32(256) {
		v6581 = v6637
		goto L1464
	} else {
		goto L1474
	}
L1471:
	;
	v6630 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+uint32(_c_F_StartupXLOG[109])))
	if v6630 == int32(0) {
		goto L1470
	} else {
		goto L1472
	}
L1472:
	;
	m.T0[v6630].(func(*base.Module))(m)
	mBase = m.M
	v6634 = m.ExcPending
	if v6634 != 0 {
		goto L32
	} else {
		goto L1473
	}
L1473:
	;
	goto L1470
L1474:
	;
	goto L1465
L1475:
	;
	if v6642 != 0 {
		goto L1476
	} else {
		goto L1477
	}
L1476:
	;
	v6645 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6646 = *(*int64)(unsafe.Add(mBase, uint32(v6645)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+340)) = uint32(v6646)
	v6649 = int64(base.Ui64(v6646) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+336)) = uint32(v6649)
	F_errmsg(m, int32(_a_F_StartupXLOG_218), v6432+int32(336))
	mBase = m.M
	v6655 = m.ExcPending
	if v6655 != 0 {
		goto L32
	} else {
		goto L1479
	}
L1477:
	;
	goto L1478
L1478:
	;
	v6663 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v6663 == int32(0) {
		goto L1481
	} else {
		goto L1482
	}
L1479:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1760), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6660 = m.ExcPending
	if v6660 != 0 {
		goto L32
	} else {
		goto L1480
	}
L1480:
	;
	goto L1478
L1481:
	;
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v6667 = m.ExcPending
	if v6667 != 0 {
		goto L32
	} else {
		goto L1484
	}
L1482:
	;
	goto L1483
L1483:
	;
	v6668 = v6567
	v6673 = v6568
	goto L1485
L1484:
	;
	goto L1483
L1485:
	;
	v6703 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v6703 != 0 {
		goto L1487
	} else {
		goto L1488
	}
L1486:
	;
	v8937 = v8778
	goto L1444
L1487:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v6771 = m.ExcPending
	if v6771 != 0 {
		goto L32
	} else {
		goto L1498
	}
L1488:
	;
	v6711 = m.G0
	v6713 = v6711 - int32(16)
	m.G0 = v6713
	v6716 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[110]))
	if v6716 != 0 {
		goto L1490
	} else {
		goto L1491
	}
L1489:
	;
	if base.B2i32(v6716 != int32(0)) == int32(0) {
		goto L1487
	} else {
		goto L1493
	}
L1490:
	;
	v6717 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v6719 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[111]))
	F_TimestampDifference(m, v6719, v6717, v6713+int32(12), v6713+int32(8))
	mBase = m.M
	v6725 = *(*int32)(unsafe.Add(mBase, uint32(v6713)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6432+int32(560)))) = v6725
	v6727 = *(*int32)(unsafe.Add(mBase, uint32(v6713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6432+int32(540)))) = v6727
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[110])) = int32(0)
	goto L1492
L1491:
	;
	goto L1492
L1492:
	;
	m.G0 = v6713 + int32(16)
	goto L1489
L1493:
	;
	v6742 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6743 = m.ExcPending
	if v6743 != 0 {
		goto L32
	} else {
		goto L1494
	}
L1494:
	;
	if v6742 == int32(0) {
		goto L1487
	} else {
		goto L1495
	}
L1495:
	;
	v6747 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6748 = *(*int64)(unsafe.Add(mBase, uint32(v6747)+32))
	v6749 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+320)) = v6749
	v6751 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+540))
	v6753 = base.I32_div_s(v6751, int32(_a_F_StartupXLOG_219))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+324)) = v6753
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+332)) = uint32(v6748)
	v6757 = int64(base.Ui64(v6748) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+328)) = uint32(v6757)
	F_errmsg(m, int32(_a_F_StartupXLOG_220), v6432+int32(320))
	mBase = m.M
	v6763 = m.ExcPending
	if v6763 != 0 {
		goto L32
	} else {
		goto L1496
	}
L1496:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1773), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v6768 = m.ExcPending
	if v6768 != 0 {
		goto L32
	} else {
		goto L1497
	}
L1497:
	;
	goto L1487
L1498:
	;
	v6773 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v6774 = *(*int32)(unsafe.Add(mBase, uint32(v6773)+80))
	if v6774 != 0 {
		goto L1499
	} else {
		goto L1500
	}
L1499:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v6777 = m.ExcPending
	if v6777 != 0 {
		goto L32
	} else {
		goto L1502
	}
L1500:
	;
	goto L1501
L1501:
	;
	v6779 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v6779 != int32(1) {
		goto L1503
	} else {
		goto L1504
	}
L1502:
	;
	goto L1501
L1503:
	;
	v7174 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[112]))
	if v7174 <= int32(0) {
		goto L1588
	} else {
		goto L1589
	}
L1504:
	;
	v6783 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v6785 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v6785 != int32(5) {
		goto L1505
	} else {
		goto L1506
	}
L1505:
	;
	if v6785 != int32(4) {
		goto L1514
	} else {
		goto L1515
	}
L1506:
	;
	v6789 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[99])))
	if v6789&int32(1) == int32(0) {
		goto L1505
	} else {
		goto L1507
	}
L1507:
	;
	v6796 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6797 = m.ExcPending
	if v6797 != 0 {
		goto L32
	} else {
		goto L1508
	}
L1508:
	;
	if v6796 != 0 {
		goto L1509
	} else {
		goto L1510
	}
L1509:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_221), int32(0))
	mBase = m.M
	v6801 = m.ExcPending
	if v6801 != 0 {
		goto L32
	} else {
		goto L1512
	}
L1510:
	;
	goto L1511
L1511:
	;
	v6808 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v6808
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v6808
	v6814 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v6814
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v6814)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v6814)
	goto L1445
L1512:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2614), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v6806 = m.ExcPending
	if v6806 != 0 {
		goto L32
	} else {
		goto L1513
	}
L1513:
	;
	goto L1511
L1514:
	;
	v6869 = *(*int32)(unsafe.Add(mBase, uint32(v6783)+96))
	v6870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6869)+49)))
	if v6870 != int32(1) {
		goto L1503
	} else {
		goto L1522
	}
L1515:
	;
	v6825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[113])))
	if v6825&int32(1) != 0 {
		goto L1514
	} else {
		goto L1516
	}
L1516:
	;
	v6828 = *(*int64)(unsafe.Add(mBase, uint32(v6783)+32))
	v6830 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	if base.Ui64(v6828) < base.Ui64(v6830) {
		goto L1514
	} else {
		goto L1517
	}
L1517:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v6828
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = int64(0)
	v6838 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v6838
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v6838)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v6838)
	v6848 = F_errstart(m, int32(15), v6838)
	mBase = m.M
	v6849 = m.ExcPending
	if v6849 != 0 {
		goto L32
	} else {
		goto L1518
	}
L1518:
	;
	if v6848 == int32(0) {
		goto L1445
	} else {
		goto L1519
	}
L1519:
	;
	v6853 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+308)) = uint32(v6853)
	v6856 = int64(base.Ui64(v6853) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+304)) = uint32(v6856)
	F_errmsg(m, int32(_a_F_StartupXLOG_223), v6432+int32(304))
	mBase = m.M
	v6862 = m.ExcPending
	if v6862 != 0 {
		goto L32
	} else {
		goto L1520
	}
L1520:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2636), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v6867 = m.ExcPending
	if v6867 != 0 {
		goto L32
	} else {
		goto L1521
	}
L1521:
	;
	goto L1445
L1522:
	;
	v6873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6869)+48)))
	switch int32(base.Ui32(v6873)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		goto L1528
	default:
		goto L1503
	case 2:
		goto L1526
	case 3:
		goto L1527
	case 4:
		goto L1525
	}
L1523:
	;
	v7108 = int32(0)
	v7110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[113])))
	v7111 = int32(1)
	v7114 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v7110&v7111|base.B2i32(v7114 != v7111) == v7108 {
		goto L1570
	} else {
		goto L1571
	}
L1524:
	;
	v7106 = v7104
	v7107 = int32(0)
	goto L1523
L1525:
	;
	v7000 = *(*int32)(unsafe.Add(mBase, uint32(v6869)+64))
	v7002 = v6432 + int32(560)
	v7003 = int32(0)
	base.MemoryFill(m, v7002, v7003, int32(264))
	v7009 = *(*int64)(unsafe.Add(mBase, uint32(v7000)))
	*(*int64)(unsafe.Add(mBase, uint32(v7002))) = v7009
	if v7003 <= base.I32_extend8_s(v6873) {
		goto L1552
	} else {
		goto L1553
	}
L1526:
	;
	v6999 = *(*int32)(unsafe.Add(mBase, uint32(v6869)+36))
	v7104 = v6999
	goto L1524
L1527:
	;
	v6880 = *(*int32)(unsafe.Add(mBase, uint32(v6869)+64))
	v6882 = v6432 + int32(560)
	v6883 = int32(0)
	base.MemoryFill(m, v6882, v6883, int32(288))
	v6889 = *(*int64)(unsafe.Add(mBase, uint32(v6880)))
	*(*int64)(unsafe.Add(mBase, uint32(v6882))) = v6889
	if v6883 <= base.I32_extend8_s(v6873) {
		goto L1530
	} else {
		goto L1531
	}
L1528:
	;
	v6878 = *(*int32)(unsafe.Add(mBase, uint32(v6869)+36))
	v7106 = v6878
	v7107 = int32(1)
	goto L1523
L1529:
	;
	v6997 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+612))
	v7106 = v6997
	v7107 = int32(1)
	goto L1523
L1530:
	;
	goto L1529
L1531:
	;
	v6894 = *(*int32)(unsafe.Add(mBase, uint32(v6880)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+8)) = v6894
	if v6894&int32(1) != 0 {
		goto L1532
	} else {
		goto L1533
	}
L1532:
	;
	v6898 = *(*int32)(unsafe.Add(mBase, uint32(v6880)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+12)) = v6898
	v6900 = *(*int32)(unsafe.Add(mBase, uint32(v6880)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+16)) = v6900
	v6906 = v6880 + int32(20)
	goto L1534
L1533:
	;
	v6906 = v6880 + int32(12)
	goto L1534
L1534:
	;
	if v6894&int32(2) != 0 {
		goto L1535
	} else {
		goto L1536
	}
L1535:
	;
	v6909 = *(*int32)(unsafe.Add(mBase, uint32(v6906)))
	v6911 = v6906 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+24)) = v6911
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+20)) = v6909
	v6917 = v6911 + v6909<<(uint(int32(2))%32)
	goto L1537
L1536:
	;
	v6917 = v6906
	goto L1537
L1537:
	;
	if v6894&int32(4) != 0 {
		goto L1538
	} else {
		goto L1539
	}
L1538:
	;
	v6921 = *(*int32)(unsafe.Add(mBase, uint32(v6917)))
	v6923 = v6917 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+32)) = v6923
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+28)) = v6921
	v6926 = *(*int32)(unsafe.Add(mBase, uint32(v6917)))
	v6930 = v6923 + v6926*int32(12)
	goto L1540
L1539:
	;
	v6930 = v6917
	goto L1540
L1540:
	;
	if v6894&int32(256) != 0 {
		goto L1541
	} else {
		goto L1542
	}
L1541:
	;
	v6935 = *(*int32)(unsafe.Add(mBase, uint32(v6930)))
	v6936 = int32(4)
	v6937 = v6930 + v6936
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+40)) = v6937
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+36)) = v6935
	v6940 = *(*int32)(unsafe.Add(mBase, uint32(v6930)))
	v6944 = v6937 + v6940<<(uint(v6936)%32)
	goto L1543
L1542:
	;
	v6944 = v6930
	goto L1543
L1543:
	;
	if v6894&int32(8) != 0 {
		goto L1544
	} else {
		goto L1545
	}
L1544:
	;
	v6949 = *(*int32)(unsafe.Add(mBase, uint32(v6944)))
	v6950 = int32(4)
	v6951 = v6944 + v6950
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+48)) = v6951
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+44)) = v6949
	v6954 = *(*int32)(unsafe.Add(mBase, uint32(v6944)))
	v6958 = v6951 + v6954<<(uint(v6950)%32)
	goto L1546
L1545:
	;
	v6958 = v6944
	goto L1546
L1546:
	;
	if v6894&int32(16) == int32(0) {
		v6982 = v6894
		v6983 = v6958
		goto L1547
	} else {
		goto L1548
	}
L1547:
	;
	if v6982&int32(32) == int32(0) {
		goto L1530
	} else {
		goto L1550
	}
L1548:
	;
	v6965 = *(*int32)(unsafe.Add(mBase, uint32(v6958)))
	*(*int32)(unsafe.Add(mBase, uint32(v6882)+52)) = v6965
	v6968 = v6958 + int32(4)
	if v6894&int32(128) == int32(0) {
		v6982 = v6894
		v6983 = v6968
		goto L1547
	} else {
		goto L1549
	}
L1549:
	;
	v6976 = F_strlcpy(m, v6432+int32(616), v6968, int32(200))
	mBase = m.M
	v6977 = F_strlen(m, v6968)
	mBase = m.M
	v6981 = *(*int32)(unsafe.Add(mBase, uint32(v6882)+8))
	v6982 = v6981
	v6983 = v6977 + v6968 + int32(1)
	goto L1547
L1550:
	;
	v6988 = *(*int64)(unsafe.Add(mBase, uint32(v6983)))
	v6989 = *(*int64)(unsafe.Add(mBase, uint32(v6983)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v6882)+280)) = v6989
	*(*int64)(unsafe.Add(mBase, uint32(v6882)+272)) = v6988
	goto L1530
L1551:
	;
	v7103 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+604))
	v7104 = v7103
	goto L1524
L1552:
	;
	goto L1551
L1553:
	;
	v7014 = *(*int32)(unsafe.Add(mBase, uint32(v7000)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+8)) = v7014
	if v7014&int32(1) != 0 {
		goto L1554
	} else {
		goto L1555
	}
L1554:
	;
	v7018 = *(*int32)(unsafe.Add(mBase, uint32(v7000)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+12)) = v7018
	v7020 = *(*int32)(unsafe.Add(mBase, uint32(v7000)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+16)) = v7020
	v7026 = v7000 + int32(20)
	goto L1556
L1555:
	;
	v7026 = v7000 + int32(12)
	goto L1556
L1556:
	;
	if v7014&int32(2) != 0 {
		goto L1557
	} else {
		goto L1558
	}
L1557:
	;
	v7029 = *(*int32)(unsafe.Add(mBase, uint32(v7026)))
	v7031 = v7026 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+24)) = v7031
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+20)) = v7029
	v7037 = v7031 + v7029<<(uint(int32(2))%32)
	goto L1559
L1558:
	;
	v7037 = v7026
	goto L1559
L1559:
	;
	if v7014&int32(4) != 0 {
		goto L1560
	} else {
		goto L1561
	}
L1560:
	;
	v7041 = *(*int32)(unsafe.Add(mBase, uint32(v7037)))
	v7043 = v7037 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+32)) = v7043
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+28)) = v7041
	v7046 = *(*int32)(unsafe.Add(mBase, uint32(v7037)))
	v7050 = v7043 + v7046*int32(12)
	goto L1562
L1561:
	;
	v7050 = v7037
	goto L1562
L1562:
	;
	if v7014&int32(256) != 0 {
		goto L1563
	} else {
		goto L1564
	}
L1563:
	;
	v7055 = *(*int32)(unsafe.Add(mBase, uint32(v7050)))
	v7056 = int32(4)
	v7057 = v7050 + v7056
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+40)) = v7057
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+36)) = v7055
	v7060 = *(*int32)(unsafe.Add(mBase, uint32(v7050)))
	v7064 = v7057 + v7060<<(uint(v7056)%32)
	goto L1565
L1564:
	;
	v7064 = v7050
	goto L1565
L1565:
	;
	if v7014&int32(16) == int32(0) {
		v7088 = v7014
		v7089 = v7064
		goto L1566
	} else {
		goto L1567
	}
L1566:
	;
	if v7088&int32(32) == int32(0) {
		goto L1552
	} else {
		goto L1569
	}
L1567:
	;
	v7071 = *(*int32)(unsafe.Add(mBase, uint32(v7064)))
	*(*int32)(unsafe.Add(mBase, uint32(v7002)+44)) = v7071
	v7074 = v7064 + int32(4)
	if v7014&int32(128) == int32(0) {
		v7088 = v7014
		v7089 = v7074
		goto L1566
	} else {
		goto L1568
	}
L1568:
	;
	v7082 = F_strlcpy(m, v6432+int32(608), v7074, int32(200))
	mBase = m.M
	v7083 = F_strlen(m, v7074)
	mBase = m.M
	v7087 = *(*int32)(unsafe.Add(mBase, uint32(v7002)+8))
	v7088 = v7087
	v7089 = v7083 + v7074 + int32(1)
	goto L1566
L1569:
	;
	v7094 = *(*int64)(unsafe.Add(mBase, uint32(v7089)))
	v7095 = *(*int64)(unsafe.Add(mBase, uint32(v7089)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7002)+256)) = v7095
	*(*int64)(unsafe.Add(mBase, uint32(v7002)+248)) = v7094
	goto L1552
L1570:
	;
	v7121 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	v7123 = base.B2i32(v7106 == v7121)
	goto L1572
L1571:
	;
	v7123 = v7108
	goto L1572
L1572:
	;
	v7124 = *(*int32)(unsafe.Add(mBase, uint32(v6783)+96))
	v7125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7124)+48)))
	v7126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7124)+49)))
	if base.B2i32(v7126 == int32(0))&base.B2i32(v7125&int32(240) == int32(112)) != 0 {
		goto L1573
	} else {
		goto L1574
	}
L1573:
	;
	v7149 = *(*int32)(unsafe.Add(mBase, uint32(v7124)+64))
	v7150 = *(*int64)(unsafe.Add(mBase, uint32(v7149)))
	if v7114 == int32(2) {
		goto L1581
	} else {
		goto L1582
	}
L1574:
	;
	if v7126 != int32(1) {
		goto L1575
	} else {
		goto L1576
	}
L1575:
	;
	if v7123 == int32(0) {
		goto L1503
	} else {
		goto L1579
	}
L1576:
	;
	v7136 = int32(4)
	v7139 = int32(base.Ui32(v7125)>>(uint(v7136)%32)) & int32(7)
	if base.Ui32(v7136) < base.Ui32(v7139) {
		goto L1575
	} else {
		goto L1577
	}
L1577:
	;
	if v7139 != int32(1) {
		goto L1573
	} else {
		goto L1578
	}
L1578:
	;
	goto L1575
L1579:
	;
	v8801 = int64(0)
	goto L1446
L1580:
	;
	if v7154 <= v7150 {
		v8801 = v7150
		goto L1446
	} else {
		goto L1587
	}
L1581:
	;
	v7154 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[25]))
	if v7110&int32(1) == int32(0) {
		goto L1580
	} else {
		goto L1584
	}
L1582:
	;
	goto L1583
L1583:
	;
	if v7123 == int32(0) {
		goto L1503
	} else {
		goto L1586
	}
L1584:
	;
	if v7154 < v7150 {
		v8801 = v7150
		goto L1446
	} else {
		goto L1585
	}
L1585:
	;
	goto L1503
L1586:
	;
	v8801 = v7150
	goto L1446
L1587:
	;
	goto L1503
L1588:
	;
	v7453 = int32(0)
	v7454 = int32(_a_F_StartupXLOG_224)
	v7455 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114])) = v6432 + int32(540)
	v7461 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+548)) = v7461
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+544)) = int32(413)
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+540)) = v7455
	v7466 = *(*int32)(unsafe.Add(mBase, uint32(v6673)+4))
	F_AdvanceNextFullTransactionIdPastXid(m, v7466)
	mBase = m.M
	v7468 = m.ExcPending
	if v7468 != 0 {
		goto L32
	} else {
		goto L1630
	}
L1589:
	;
	v7178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[99])))
	if v7178&int32(1) == int32(0) {
		goto L1588
	} else {
		goto L1590
	}
L1590:
	;
	v7184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v7184&int32(1) == int32(0) {
		goto L1588
	} else {
		goto L1591
	}
L1591:
	;
	v7190 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v7191 = *(*int32)(unsafe.Add(mBase, uint32(v7190)+96))
	v7192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7191)+49)))
	if v7192 != int32(1) {
		goto L1588
	} else {
		goto L1592
	}
L1592:
	;
	v7195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7191)+48)))
	v7197 = v7195 & int32(112)
	if v7197 != 0 {
		goto L1593
	} else {
		goto L1594
	}
L1593:
	;
	v7201 = base.B2i32(v7197 != int32(48))
	goto L1595
L1594:
	;
	v7201 = int32(0)
	goto L1595
L1595:
	;
	if v7201 != 0 {
		goto L1588
	} else {
		goto L1596
	}
L1596:
	;
	v7202 = int32(4)
	v7205 = int32(base.Ui32(v7195)>>(uint(v7202)%32)) & int32(7)
	if base.B2i32(base.Ui32(v7202) < base.Ui32(v7205))|base.B2i32(v7205 == int32(1)) != 0 {
		goto L1588
	} else {
		goto L1597
	}
L1597:
	;
	v7211 = *(*int32)(unsafe.Add(mBase, uint32(v7191)+64))
	v7212 = *(*int64)(unsafe.Add(mBase, uint32(v7211)))
	v7216 = m.G0
	v7217 = int32(16)
	v7218 = v7216 - v7217
	m.G0 = v7218
	F_gettimeofday(m, v7218)
	mBase = m.M
	v7221 = *(*int64)(unsafe.Add(mBase, uint32(v7218)))
	v7222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7218)+8)))
	m.G0 = v7218 + v7217
	v7230 = v7222 + v7221*int64(1000000) - int64(946684800000000)
	goto L1598
L1598:
	;
	v7234 = v7212 + base.I64_extend_i32_u(v7174)*int64(1000)
	if v7234 <= v7230 {
		v7252 = int32(0)
		goto L1600
	} else {
		goto L1601
	}
L1599:
	;
	if v7252 <= int32(0) {
		goto L1588
	} else {
		goto L1603
	}
L1600:
	;
	goto L1599
L1601:
	;
	v7240 = v7234 - v7230
	if base.B2i32(int64(0) < v7230)^base.B2i32(v7240 < v7234)|base.B2i32(int64(2147483646000) < v7240) != 0 {
		v7252 = int32(2147483647)
		goto L1600
	} else {
		goto L1602
	}
L1602:
	;
	v7249 = base.I64_div_s(v7240+int64(999), int64(1000))
	v7252 = base.I32_wrap_i64(v7249)
	goto L1600
L1603:
	;
	v7256 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v7256+int32(4)))) = int32(0)
	goto L1604
L1604:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7262 = m.ExcPending
	if v7262 != 0 {
		goto L32
	} else {
		goto L1605
	}
L1605:
	;
	v7263 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7264 = m.ExcPending
	if v7264 != 0 {
		goto L32
	} else {
		goto L1607
	}
L1606:
	;
	v7412 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7413 = *(*int32)(unsafe.Add(mBase, uint32(v7412)+80))
	if v7413 == int32(0) {
		goto L1588
	} else {
		goto L1628
	}
L1607:
	;
	if v7263 != 0 {
		goto L1606
	} else {
		goto L1608
	}
L1608:
	;
	goto L1609
L1609:
	;
	v7300 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[112])))
	v7304 = m.G0
	v7305 = int32(16)
	v7306 = v7304 - v7305
	m.G0 = v7306
	F_gettimeofday(m, v7306)
	mBase = m.M
	v7309 = *(*int64)(unsafe.Add(mBase, uint32(v7306)))
	v7310 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7306)+8)))
	m.G0 = v7306 + v7305
	v7318 = v7310 + v7309*int64(1000000) - int64(946684800000000)
	goto L1611
L1610:
	;
	goto L1606
L1611:
	;
	v7321 = v7300*int64(1000) + v7212
	if v7321 <= v7318 {
		v7339 = int32(0)
		goto L1613
	} else {
		goto L1614
	}
L1612:
	;
	if v7339 <= int32(0) {
		goto L1606
	} else {
		goto L1616
	}
L1613:
	;
	goto L1612
L1614:
	;
	v7327 = v7321 - v7318
	if base.B2i32(int64(0) < v7318)^base.B2i32(v7327 < v7321)|base.B2i32(int64(2147483646000) < v7327) != 0 {
		v7339 = int32(2147483647)
		goto L1613
	} else {
		goto L1615
	}
L1615:
	;
	v7336 = base.I64_div_s(v7327+int64(999), int64(1000))
	v7339 = base.I32_wrap_i64(v7336)
	goto L1613
L1616:
	;
	v7344 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v7345 = m.ExcPending
	if v7345 != 0 {
		goto L32
	} else {
		goto L1617
	}
L1617:
	;
	if v7344 != 0 {
		goto L1618
	} else {
		goto L1619
	}
L1618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+256)) = v7339
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_225), v6432+int32(256))
	mBase = m.M
	v7351 = m.ExcPending
	if v7351 != 0 {
		goto L32
	} else {
		goto L1621
	}
L1619:
	;
	goto L1620
L1620:
	;
	v7358 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7363 = F_WaitLatch(m, v7358+int32(4), int32(41), v7339, int32(150994947))
	mBase = m.M
	v7364 = m.ExcPending
	if v7364 != 0 {
		goto L32
	} else {
		goto L1623
	}
L1621:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(3078), int32(_a_F_StartupXLOG_226))
	mBase = m.M
	v7356 = m.ExcPending
	if v7356 != 0 {
		goto L32
	} else {
		goto L1622
	}
L1622:
	;
	goto L1620
L1623:
	;
	v7366 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v7366+int32(4)))) = int32(0)
	goto L1624
L1624:
	;
	F_ProcessStartupProcInterrupts(m)
	mBase = m.M
	v7372 = m.ExcPending
	if v7372 != 0 {
		goto L32
	} else {
		goto L1625
	}
L1625:
	;
	v7373 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v7374 = m.ExcPending
	if v7374 != 0 {
		goto L32
	} else {
		goto L1626
	}
L1626:
	;
	if v7373 == int32(0) {
		goto L1609
	} else {
		goto L1627
	}
L1627:
	;
	goto L1610
L1628:
	;
	F_recoveryPausesHere(m, int32(0))
	mBase = m.M
	v7418 = m.ExcPending
	if v7418 != 0 {
		goto L32
	} else {
		goto L1629
	}
L1629:
	;
	goto L1588
L1630:
	;
	v7469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6673)+17)))
	if v7469 != 0 {
		v7536 = v6668
		v7539 = v7453
		goto L1638
	} else {
		goto L1639
	}
L1631:
	;
	v8778 = int32(0)
	v8780 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v8783 = F_ReadRecord(m, v8780, int32(15), v8778, v7536)
	mBase = m.M
	v8784 = m.ExcPending
	if v8784 != 0 {
		goto L32
	} else {
		goto L1942
	}
L1632:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8752 = m.ExcPending
	if v8752 != 0 {
		goto L32
	} else {
		goto L1939
	}
L1633:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8733 = m.ExcPending
	if v8733 != 0 {
		goto L32
	} else {
		goto L1935
	}
L1634:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8710 = m.ExcPending
	if v8710 != 0 {
		goto L32
	} else {
		goto L1932
	}
L1635:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8686 = m.ExcPending
	if v8686 != 0 {
		goto L32
	} else {
		goto L1929
	}
L1636:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8670 = m.ExcPending
	if v8670 != 0 {
		goto L32
	} else {
		goto L1926
	}
L1637:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v8653 = m.ExcPending
	if v8653 != 0 {
		goto L32
	} else {
		goto L1923
	}
L1638:
	;
	v7542 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7543 = *(*int32)(unsafe.Add(mBase, uint32(v7542)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v7542)+96)) = int32(1)
	if v7543 != 0 {
		goto L1665
	} else {
		goto L1666
	}
L1639:
	;
	v7470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6673)+16)))
	v7472 = v7470 & int32(240)
	if v7472 != 0 {
		goto L1640
	} else {
		goto L1641
	}
L1640:
	;
	v7476 = base.B2i32(v7472 != int32(144))
	goto L1642
L1641:
	;
	v7476 = int32(0)
	goto L1642
L1642:
	;
	if v7476 != 0 {
		v7536 = v6668
		v7539 = v7453
		goto L1638
	} else {
		goto L1643
	}
L1643:
	;
	v7477 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7478 = *(*int32)(unsafe.Add(mBase, uint32(v7477)+64))
	v7479 = *(*int32)(unsafe.Add(mBase, uint32(v7478)+8))
	if v7479 == v6668 {
		v7536 = v6668
		v7539 = v7453
		goto L1638
	} else {
		goto L1644
	}
L1644:
	;
	v7481 = *(*int32)(unsafe.Add(mBase, uint32(v7478)+12))
	if v7481 != v6668 {
		goto L1637
	} else {
		goto L1645
	}
L1645:
	;
	if base.Ui32(v7479) < base.Ui32(v6668) {
		goto L1636
	} else {
		goto L1646
	}
L1646:
	;
	v7484 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	v7486 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[44]))
	v7487 = int32(0)
	if v7486 == v7487 {
		goto L1648
	} else {
		goto L1649
	}
L1647:
	;
	if v7526 == int32(0) {
		goto L1636
	} else {
		goto L1660
	}
L1648:
	;
	v7526 = int32(0)
	goto L1647
L1649:
	;
	goto L1650
L1650:
	;
	v7493 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+4))
	if v7493 <= int32(0) {
		v7520 = v7487
		goto L1651
	} else {
		goto L1652
	}
L1651:
	;
	v7526 = v7520
	goto L1647
L1652:
	;
	v7496 = int32(0)
	if v7496 < v7493 {
		goto L1653
	} else {
		goto L1654
	}
L1653:
	;
	v7499 = v7493
	goto L1655
L1654:
	;
	v7499 = v7496
	goto L1655
L1655:
	;
	v7500 = *(*int32)(unsafe.Add(mBase, uint32(v7486)+12))
	v7503 = int32(0)
	goto L1656
L1656:
	;
	v7510 = *(*int32)(unsafe.Add(mBase, uint32(v7500+v7503<<(uint(int32(2))%32))))
	v7511 = *(*int32)(unsafe.Add(mBase, uint32(v7510)))
	v7512 = base.B2i32(v7511 == v7479)
	if v7511 == v7479 {
		v7520 = v7512
		goto L1651
	} else {
		goto L1658
	}
L1657:
	;
	v7520 = v7512
	goto L1651
L1658:
	;
	v7514 = v7503 + int32(1)
	if v7514 != v7499 {
		v7503 = v7514
		goto L1656
	} else {
		goto L1659
	}
L1659:
	;
	goto L1657
L1660:
	;
	v7531 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	if base.Ui64(v7484) < base.Ui64(v7531) {
		goto L1661
	} else {
		goto L1662
	}
L1661:
	;
	v7534 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	if base.Ui32(v7534) < base.Ui32(v7479) {
		goto L1635
	} else {
		goto L1664
	}
L1662:
	;
	goto L1663
L1663:
	;
	v7536 = v7479
	v7539 = int32(1)
	goto L1638
L1664:
	;
	goto L1663
L1665:
	;
	v7547 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v7547+int32(96), int32(_a_F_StartupXLOG_50), int32(1991), int32(_a_F_StartupXLOG_227))
	mBase = m.M
	v7554 = m.ExcPending
	if v7554 != 0 {
		goto L32
	} else {
		goto L1668
	}
L1666:
	;
	goto L1667
L1667:
	;
	v7555 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	v7557 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7558 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7557)+96)) = v7558
	*(*int32)(unsafe.Add(mBase, uint32(v7557)+56)) = v7536
	*(*int64)(unsafe.Add(mBase, uint32(v7557)+48)) = v7555
	v7563 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v7563 == v7558 {
		goto L1669
	} else {
		goto L1670
	}
L1668:
	;
	goto L1667
L1669:
	;
	v7572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6673)+17)))
	if v7572 != 0 {
		goto L1673
	} else {
		goto L1674
	}
L1670:
	;
	v7566 = *(*int32)(unsafe.Add(mBase, uint32(v6673)+4))
	if v7566 == int32(0) {
		goto L1669
	} else {
		goto L1671
	}
L1671:
	;
	F_RecordKnownAssignedTransactionIds(m, v7566)
	mBase = m.M
	v7570 = m.ExcPending
	if v7570 != 0 {
		goto L32
	} else {
		goto L1672
	}
L1672:
	;
	goto L1669
L1673:
	;
	v7663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6673)+17)))
	v7665 = v7663 << (uint(int32(5)) % 32)
	v7668 = *(*int32)(unsafe.Add(mBase, uint32(v7665)+uint32(_c_F_StartupXLOG[106])))
	if v7668 == int32(0) {
		goto L1699
	} else {
		goto L1700
	}
L1674:
	;
	v7573 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7573)+48)))
	v7576 = v7574 & int32(240)
	if v7576 != int32(80) {
		goto L1675
	} else {
		goto L1676
	}
L1675:
	;
	if v7576 != int32(208) {
		goto L1673
	} else {
		goto L1678
	}
L1676:
	;
	goto L1677
L1677:
	;
	v7615 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	v7617 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53]))
	v7618 = *(*int32)(unsafe.Add(mBase, uint32(v7573)+64))
	v7619 = *(*int64)(unsafe.Add(mBase, uint32(v7618)))
	v7622 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v7623 = m.ExcPending
	if v7623 != 0 {
		goto L32
	} else {
		goto L1687
	}
L1678:
	;
	v7581 = *(*int32)(unsafe.Add(mBase, uint32(v7573)+64))
	v7582 = *(*int64)(unsafe.Add(mBase, uint32(v7581)))
	v7583 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+64))
	if v7582 != v7583 {
		goto L1634
	} else {
		goto L1679
	}
L1679:
	;
	v7585 = *(*int64)(unsafe.Add(mBase, uint32(v7581)+8))
	v7587 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52])) = v7587
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51])) = v7587
	v7594 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v7595 = m.ExcPending
	if v7595 != 0 {
		goto L32
	} else {
		goto L1680
	}
L1680:
	;
	if v7594 != 0 {
		goto L1681
	} else {
		goto L1682
	}
L1681:
	;
	v7596 = F_timestamptz_to_str(m, v7585)
	mBase = m.M
	v7597 = m.ExcPending
	if v7597 != 0 {
		goto L32
	} else {
		goto L1684
	}
L1682:
	;
	goto L1683
L1683:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7461)+64)) = int64(0)
	goto L1673
L1684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+168)) = v7596
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+164)) = uint32(v7582)
	v7601 = int64(base.Ui64(v7582) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+160)) = uint32(v7601)
	F_errmsg(m, int32(_a_F_StartupXLOG_228), v6432+int32(160))
	mBase = m.M
	v7607 = m.ExcPending
	if v7607 != 0 {
		goto L32
	} else {
		goto L1685
	}
L1685:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2117), int32(_a_F_StartupXLOG_229))
	mBase = m.M
	v7612 = m.ExcPending
	if v7612 != 0 {
		goto L32
	} else {
		goto L1686
	}
L1686:
	;
	goto L1683
L1687:
	;
	if v7619 == v7617 {
		goto L1688
	} else {
		goto L1689
	}
L1688:
	;
	if v7622 != 0 {
		goto L1691
	} else {
		goto L1692
	}
L1689:
	;
	goto L1690
L1690:
	;
	if v7622 == int32(0) {
		goto L1673
	} else {
		goto L1696
	}
L1691:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_230), int32(0))
	mBase = m.M
	v7628 = m.ExcPending
	if v7628 != 0 {
		goto L32
	} else {
		goto L1694
	}
L1692:
	;
	goto L1693
L1693:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[54])) = v7615
	goto L1673
L1694:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2138), int32(_a_F_StartupXLOG_229))
	mBase = m.M
	v7633 = m.ExcPending
	if v7633 != 0 {
		goto L32
	} else {
		goto L1695
	}
L1695:
	;
	goto L1693
L1696:
	;
	v7639 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[53]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+204)) = uint32(v7639)
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+196)) = uint32(v7619)
	v7642 = int64(32)
	v7643 = int64(base.Ui64(v7619) >> (uint(v7642) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+192)) = uint32(v7643)
	v7646 = int64(base.Ui64(v7639) >> (uint(v7642) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+200)) = uint32(v7646)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_231), v6432+int32(192))
	mBase = m.M
	v7652 = m.ExcPending
	if v7652 != 0 {
		goto L32
	} else {
		goto L1697
	}
L1697:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2144), int32(_a_F_StartupXLOG_229))
	mBase = m.M
	v7657 = m.ExcPending
	if v7657 != 0 {
		goto L32
	} else {
		goto L1698
	}
L1698:
	;
	goto L1673
L1699:
	;
	F_RmgrNotFound(m, v7663)
	mBase = m.M
	v7672 = m.ExcPending
	if v7672 != 0 {
		goto L32
	} else {
		goto L1702
	}
L1700:
	;
	goto L1701
L1701:
	;
	v7673 = *(*int32)(unsafe.Add(mBase, uint32(v7665)+uint32(_c_F_StartupXLOG[115])))
	m.T0[v7673].(func(*base.Module, int32))(m, v7461)
	mBase = m.M
	v7675 = m.ExcPending
	if v7675 != 0 {
		goto L32
	} else {
		goto L1703
	}
L1702:
	;
	goto L1701
L1703:
	;
	v7676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6673)+16)))
	if v7676&int32(2) == int32(0) {
		goto L1704
	} else {
		goto L1705
	}
L1704:
	;
	v7948 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+540))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[114])) = v7948
	v7951 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7952 = *(*int32)(unsafe.Add(mBase, uint32(v7951)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v7951)+96)) = int32(1)
	if v7952 != 0 {
		goto L1766
	} else {
		goto L1767
	}
L1705:
	;
	v7681 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7681)+49)))
	v7684 = v7682 << (uint(int32(5)) % 32)
	v7685 = *(*int32)(unsafe.Add(mBase, uint32(v7684)+uint32(_c_F_StartupXLOG[106])))
	if v7685 != 0 {
		goto L1706
	} else {
		goto L1707
	}
L1706:
	;
	v7689 = v7681
	goto L1708
L1707:
	;
	F_RmgrNotFound(m, v7682)
	mBase = m.M
	v7687 = m.ExcPending
	if v7687 != 0 {
		goto L32
	} else {
		goto L1709
	}
L1708:
	;
	v7690 = *(*int32)(unsafe.Add(mBase, uint32(v7689)+72))
	if v7690 < int32(0) {
		goto L1704
	} else {
		goto L1710
	}
L1709:
	;
	v7688 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7689 = v7688
	goto L1708
L1710:
	;
	v7695 = *(*int32)(unsafe.Add(mBase, uint32(v7684)+uint32(_c_F_StartupXLOG[116])))
	v7702 = int32(0)
	goto L1711
L1711:
	;
	v7732 = v7702 & int32(255)
	v7734 = v6432 + int32(560)
	v7736 = v6432 + int32(556)
	v7738 = v6432 + int32(552)
	v7739 = int32(0)
	v7741 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7742 = *(*int32)(unsafe.Add(mBase, uint32(v7741)+72))
	if v7742 < v7732 {
		v7766 = v7739
		goto L1715
	} else {
		goto L1716
	}
L1712:
	;
	goto L1704
L1713:
	;
	v7909 = v7702 + int32(1)
	v7910 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7911 = *(*int32)(unsafe.Add(mBase, uint32(v7910)+72))
	if v7909 <= v7911 {
		v7702 = v7909
		goto L1711
	} else {
		goto L1765
	}
L1714:
	;
	if v7766 == int32(0) {
		goto L1713
	} else {
		goto L1728
	}
L1715:
	;
	goto L1714
L1716:
	;
	v7746 = v7741 + v7732*int32(52)
	v7747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7746)+76)))
	if v7747 != int32(1) {
		v7766 = v7739
		goto L1715
	} else {
		goto L1717
	}
L1717:
	;
	v7751 = v7746 + int32(76)
	if v7734 != 0 {
		goto L1718
	} else {
		goto L1719
	}
L1718:
	;
	v7752 = *(*int32)(unsafe.Add(mBase, uint32(v7751)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7734)+8)) = v7752
	v7754 = *(*int64)(unsafe.Add(mBase, uint32(v7751)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v7734))) = v7754
	goto L1720
L1719:
	;
	goto L1720
L1720:
	;
	if v7736 != 0 {
		goto L1721
	} else {
		goto L1722
	}
L1721:
	;
	v7756 = *(*int32)(unsafe.Add(mBase, uint32(v7751)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7736))) = v7756
	goto L1723
L1722:
	;
	goto L1723
L1723:
	;
	if v7738 != 0 {
		goto L1724
	} else {
		goto L1725
	}
L1724:
	;
	v7758 = *(*int32)(unsafe.Add(mBase, uint32(v7751)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7738))) = v7758
	goto L1726
L1725:
	;
	goto L1726
L1726:
	;
	v7766 = int32(1)
	goto L1715
L1728:
	;
	v7769 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+96))
	v7773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7769+v7702*int32(52))+106)))
	if v7773 != 0 {
		goto L1713
	} else {
		goto L1729
	}
L1729:
	;
	v7774 = *(*int64)(unsafe.Add(mBase, uint32(v6432)+560))
	*(*int64)(unsafe.Add(mBase, uint32(v6432)+144)) = v7774
	v7776 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+568))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+152)) = v7776
	v7780 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+556))
	v7781 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+552))
	v7784 = F_XLogReadBufferExtended(m, v6432+int32(144), v7780, v7781, int32(4), int32(0))
	mBase = m.M
	v7785 = m.ExcPending
	if v7785 != 0 {
		goto L32
	} else {
		goto L1730
	}
L1730:
	;
	if v7784 == int32(0) {
		goto L1713
	} else {
		goto L1731
	}
L1731:
	;
	F_LockBuffer(m, v7784, int32(2))
	mBase = m.M
	v7790 = m.ExcPending
	if v7790 != 0 {
		goto L32
	} else {
		goto L1732
	}
L1732:
	;
	v7792 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	if v7784 < int32(0) {
		goto L1734
	} else {
		goto L1735
	}
L1733:
	;
	base.MemoryCopy(m, v7792, v7810, int32(_a_F_StartupXLOG_58))
	F_UnlockReleaseBuffer(m, v7784)
	mBase = m.M
	v7814 = m.ExcPending
	if v7814 != 0 {
		goto L32
	} else {
		goto L1737
	}
L1734:
	;
	v7796 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[117]))
	v7802 = *(*int32)(unsafe.Add(mBase, uint32(v7796+(v7784^int32(-1))<<(uint(int32(2))%32))))
	v7810 = v7802
	goto L1733
L1735:
	;
	goto L1736
L1736:
	;
	v7804 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[118]))
	v7810 = v7804 + v7784<<(uint(int32(13))%32) + int32(-8192)
	goto L1733
L1737:
	;
	v7815 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	v7817 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7818 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7817))))
	v7821 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7817)+4)))
	if base.Ui64(v7815) < base.Ui64(v7818<<(uint(int64(32))%64)|v7821) {
		goto L1713
	} else {
		goto L1738
	}
L1738:
	;
	v7825 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7826 = F_RestoreBlockImage(m, v7461, v7732, v7825)
	mBase = m.M
	v7827 = m.ExcPending
	if v7827 != 0 {
		goto L32
	} else {
		goto L1739
	}
L1739:
	;
	if v7826 == int32(0) {
		goto L1633
	} else {
		goto L1740
	}
L1740:
	;
	if v7695 != 0 {
		goto L1741
	} else {
		goto L1742
	}
L1741:
	;
	v7831 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7832 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+552))
	m.T0[v7695].(func(*base.Module, int32, int32))(m, v7831, v7832)
	mBase = m.M
	v7834 = m.ExcPending
	if v7834 != 0 {
		goto L32
	} else {
		goto L1744
	}
L1742:
	;
	goto L1743
L1743:
	;
	v7841 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[33]))
	v7843 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7844 = int32(_a_F_StartupXLOG_58)
	goto L1749
L1744:
	;
	v7836 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[34]))
	v7837 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+552))
	m.T0[v7695].(func(*base.Module, int32, int32))(m, v7836, v7837)
	mBase = m.M
	v7839 = m.ExcPending
	if v7839 != 0 {
		goto L32
	} else {
		goto L1745
	}
L1745:
	;
	goto L1743
L1746:
	;
	if v7906 != 0 {
		goto L1632
	} else {
		goto L1764
	}
L1747:
	;
	v7906 = int32(0)
	goto L1746
L1748:
	;
	v7880 = v7875
	v7881 = v7876
	v7882 = v7877
	goto L1758
L1749:
	;
	if (v7841|v7843)&int32(3) != 0 {
		v7875 = v7841
		v7876 = v7843
		v7877 = v7844
		goto L1748
	} else {
		goto L1752
	}
L1751:
	;
	if v7865 == int32(0) {
		goto L1747
	} else {
		goto L1757
	}
L1752:
	;
	v7852 = v7841
	v7853 = v7843
	v7854 = v7844
	goto L1753
L1753:
	;
	v7857 = *(*int32)(unsafe.Add(mBase, uint32(v7852)))
	v7858 = *(*int32)(unsafe.Add(mBase, uint32(v7853)))
	if v7857 != v7858 {
		v7875 = v7852
		v7876 = v7853
		v7877 = v7854
		goto L1748
	} else {
		goto L1755
	}
L1754:
	;
	goto L1751
L1755:
	;
	v7860 = int32(4)
	v7861 = v7853 + v7860
	v7863 = v7852 + v7860
	v7865 = v7854 - v7860
	if base.Ui32(int32(3)) < base.Ui32(v7865) {
		v7852 = v7863
		v7853 = v7861
		v7854 = v7865
		goto L1753
	} else {
		goto L1756
	}
L1756:
	;
	goto L1754
L1757:
	;
	v7875 = v7863
	v7876 = v7861
	v7877 = v7865
	goto L1748
L1758:
	;
	v7885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7880))))
	v7886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7881))))
	if v7885 == v7886 {
		goto L1760
	} else {
		goto L1761
	}
L1759:
	;
	v7906 = v7885 - v7886
	goto L1746
L1760:
	;
	v7888 = int32(1)
	v7893 = v7882 - v7888
	if v7893 != 0 {
		v7880 = v7880 + v7888
		v7881 = v7881 + v7888
		v7882 = v7893
		goto L1758
	} else {
		goto L1763
	}
L1761:
	;
	goto L1762
L1762:
	;
	goto L1759
L1763:
	;
	goto L1747
L1764:
	;
	goto L1713
L1765:
	;
	goto L1712
L1766:
	;
	v7956 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v7956+int32(96), int32(_a_F_StartupXLOG_50), int32(2028), int32(_a_F_StartupXLOG_227))
	mBase = m.M
	v7963 = m.ExcPending
	if v7963 != 0 {
		goto L32
	} else {
		goto L1769
	}
L1767:
	;
	goto L1768
L1768:
	;
	v7965 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v7966 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v7965)+24)) = v7966
	v7968 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7965)+96)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7965)+40)) = v7536
	*(*int64)(unsafe.Add(mBase, uint32(v7965)+32)) = v7968
	v7974 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[22])))
	if v7974 != int32(1) {
		goto L1770
	} else {
		goto L1771
	}
L1769:
	;
	goto L1768
L1770:
	;
	v7985 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[119])))
	if v7985 != 0 {
		goto L1774
	} else {
		goto L1775
	}
L1771:
	;
	v7978 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[120]))
	if v7978 <= int32(0) {
		goto L1770
	} else {
		goto L1772
	}
L1772:
	;
	F_WalSndWakeup(m, v7539, int32(1))
	mBase = m.M
	v7983 = m.ExcPending
	if v7983 != 0 {
		goto L32
	} else {
		goto L1773
	}
L1773:
	;
	goto L1770
L1774:
	;
	v7987 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[119])) = uint8(v7987)
	F_WalRcvForceReply(m)
	mBase = m.M
	v7990 = m.ExcPending
	if v7990 != 0 {
		goto L32
	} else {
		goto L1777
	}
L1775:
	;
	goto L1776
L1776:
	;
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v7992 = m.ExcPending
	if v7992 != 0 {
		goto L32
	} else {
		goto L1778
	}
L1777:
	;
	goto L1776
L1778:
	;
	if v7539 != 0 {
		goto L1779
	} else {
		goto L1780
	}
L1779:
	;
	v7993 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+40))
	F_RemoveNonParentXlogFiles(m, v7993, v7536)
	mBase = m.M
	v7995 = m.ExcPending
	if v7995 != 0 {
		goto L32
	} else {
		goto L1782
	}
L1780:
	;
	goto L1781
L1781:
	;
	v8003 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v8003 != int32(1) {
		goto L1631
	} else {
		goto L1784
	}
L1782:
	;
	v7996 = int32(_a_F_StartupXLOG_232)
	v7998 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[31])) = v7998 + int32(1)
	goto L1783
L1783:
	;
	goto L1781
L1784:
	;
	v8007 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v8008 = *(*int32)(unsafe.Add(mBase, uint32(v8007)+96))
	v8009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8008)+48)))
	v8010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8008)+49)))
	v8012 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v8010|base.B2i32(v8012 != int32(3))|base.B2i32(v8009&int32(240) != int32(112)) == int32(0) {
		goto L1785
	} else {
		goto L1786
	}
L1785:
	;
	v8023 = *(*int32)(unsafe.Add(mBase, uint32(v8008)+64))
	v8025 = v8023 + int32(8)
	v8027 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[46]))
	v8030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8025))))
	v8033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8027))))
	if base.B2i32(v8030 == int32(0))|base.B2i32(v8030 != v8033) != 0 {
		v8051 = v8030
		v8052 = v8033
		goto L1789
	} else {
		goto L1790
	}
L1786:
	;
	goto L1787
L1787:
	;
	if v8012 != int32(4) {
		goto L1832
	} else {
		goto L1833
	}
L1788:
	;
	if v8051-v8052 != 0 {
		goto L1631
	} else {
		goto L1795
	}
L1789:
	;
	goto L1788
L1790:
	;
	v8036 = v8025
	v8037 = v8027
	goto L1791
L1791:
	;
	v8040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8037)+1)))
	v8041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8036)+1)))
	if v8041 == int32(0) {
		v8051 = v8041
		v8052 = v8040
		goto L1789
	} else {
		goto L1793
	}
L1792:
	;
	v8051 = v8041
	v8052 = v8040
	goto L1789
L1793:
	;
	v8044 = int32(1)
	if v8041 == v8040 {
		v8036 = v8036 + v8044
		v8037 = v8037 + v8044
		goto L1791
	} else {
		goto L1794
	}
L1794:
	;
	goto L1792
L1795:
	;
	v8055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v8055)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = int32(0)
	v8064 = *(*int64)(unsafe.Add(mBase, uint32(v8023)))
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v8064
	v8066 = int32(_a_F_StartupXLOG_233)
	goto L1799
L1796:
	;
	v8188 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8189 = m.ExcPending
	if v8189 != 0 {
		goto L32
	} else {
		goto L1827
	}
L1797:
	;
	v8183 = F_strlen(m, v8172)
	mBase = m.M
	goto L1796
L1799:
	;
	goto L1800
L1800:
	;
	v8073 = int32(63)
	if (v8066^v8025)&int32(3) != 0 {
		goto L1804
	} else {
		goto L1805
	}
L1801:
	;
	v8176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8173))) = uint8(v8176)
	goto L1797
L1802:
	;
	v8157 = v8152
	v8158 = v8153
	v8159 = v8154
	goto L1823
L1803:
	;
	if v8147 == int32(0) {
		v8172 = v8145
		v8173 = v8146
		goto L1801
	} else {
		goto L1822
	}
L1804:
	;
	v8145 = v8025
	v8146 = v8066
	v8147 = v8073
	goto L1803
L1805:
	;
	goto L1806
L1806:
	;
	v8077 = int32(0)
	if base.B2i32(v8025&int32(3) == v8077)|int32(0) == v8077 {
		goto L1808
	} else {
		goto L1809
	}
L1807:
	;
	if v8113 == int32(0) {
		v8172 = v8110
		v8173 = v8111
		goto L1801
	} else {
		goto L1816
	}
L1808:
	;
	v8089 = v8025
	v8090 = v8066
	v8091 = v8073
	goto L1811
L1809:
	;
	goto L1810
L1810:
	;
	v8110 = v8025
	v8111 = v8066
	v8112 = v8073
	v8113 = int32(1)
	goto L1807
L1811:
	;
	v8093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8089))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8090))) = uint8(v8093)
	if v8093 == int32(0) {
		v8152 = v8089
		v8153 = v8090
		v8154 = v8091
		goto L1802
	} else {
		goto L1813
	}
L1812:
	;
	v8110 = v8104
	v8111 = v8098
	v8112 = v8100
	v8113 = v8102
	goto L1807
L1813:
	;
	v8097 = int32(1)
	v8098 = v8090 + v8097
	v8100 = v8091 - v8097
	v8101 = int32(0)
	v8102 = base.B2i32(v8100 != v8101)
	v8104 = v8089 + v8097
	if v8104&int32(3) == v8101 {
		v8110 = v8104
		v8111 = v8098
		v8112 = v8100
		v8113 = v8102
		goto L1807
	} else {
		goto L1814
	}
L1814:
	;
	if v8100 != 0 {
		v8089 = v8104
		v8090 = v8098
		v8091 = v8100
		goto L1811
	} else {
		goto L1815
	}
L1815:
	;
	goto L1812
L1816:
	;
	v8116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8110))))
	if base.B2i32(v8116 == int32(0))|base.B2i32(base.Ui32(v8112) < base.Ui32(int32(4))) != 0 {
		v8145 = v8110
		v8146 = v8111
		v8147 = v8112
		goto L1803
	} else {
		goto L1817
	}
L1817:
	;
	v8123 = v8110
	v8124 = v8111
	v8125 = v8112
	goto L1818
L1818:
	;
	v8128 = *(*int32)(unsafe.Add(mBase, uint32(v8123)))
	v8131 = int32(-2139062144)
	if (int32(16843008)-v8128|v8128)&v8131 != v8131 {
		v8152 = v8123
		v8153 = v8124
		v8154 = v8125
		goto L1802
	} else {
		goto L1820
	}
L1819:
	;
	v8145 = v8139
	v8146 = v8137
	v8147 = v8141
	goto L1803
L1820:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8124))) = v8128
	v8136 = int32(4)
	v8137 = v8124 + v8136
	v8139 = v8123 + v8136
	v8141 = v8125 - v8136
	if base.Ui32(int32(3)) < base.Ui32(v8141) {
		v8123 = v8139
		v8124 = v8137
		v8125 = v8141
		goto L1818
	} else {
		goto L1821
	}
L1821:
	;
	goto L1819
L1822:
	;
	v8152 = v8145
	v8153 = v8146
	v8154 = v8147
	goto L1802
L1823:
	;
	v8161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8157))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8158))) = uint8(v8161)
	if v8161 == int32(0) {
		v8172 = v8157
		v8173 = v8158
		goto L1801
	} else {
		goto L1825
	}
L1824:
	;
	v8172 = v8168
	v8173 = v8166
	goto L1801
L1825:
	;
	v8165 = int32(1)
	v8166 = v8158 + v8165
	v8168 = v8157 + v8165
	v8170 = v8159 - v8165
	if v8170 != 0 {
		v8157 = v8168
		v8158 = v8166
		v8159 = v8170
		goto L1823
	} else {
		goto L1826
	}
L1826:
	;
	goto L1824
L1827:
	;
	if v8188 == int32(0) {
		goto L1445
	} else {
		goto L1828
	}
L1828:
	;
	v8193 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v8194 = F_timestamptz_to_str(m, v8193)
	mBase = m.M
	v8195 = m.ExcPending
	if v8195 != 0 {
		goto L32
	} else {
		goto L1829
	}
L1829:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+36)) = v8194
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+32)) = int32(_a_F_StartupXLOG_233)
	F_errmsg(m, int32(_a_F_StartupXLOG_234), v6432+int32(32))
	mBase = m.M
	v8203 = m.ExcPending
	if v8203 != 0 {
		goto L32
	} else {
		goto L1830
	}
L1830:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2787), int32(_a_F_StartupXLOG_235))
	mBase = m.M
	v8208 = m.ExcPending
	if v8208 != 0 {
		goto L32
	} else {
		goto L1831
	}
L1831:
	;
	goto L1445
L1832:
	;
	if v8010 != int32(1) {
		goto L1631
	} else {
		goto L1840
	}
L1833:
	;
	v8212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[113])))
	if v8212&int32(1) == int32(0) {
		goto L1832
	} else {
		goto L1834
	}
L1834:
	;
	v8217 = *(*int64)(unsafe.Add(mBase, uint32(v8007)+32))
	v8219 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[47]))
	if base.Ui64(v8217) < base.Ui64(v8219) {
		goto L1832
	} else {
		goto L1835
	}
L1835:
	;
	v8222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v8222)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v8217
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = int64(0)
	v8230 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v8230
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v8230)
	v8237 = F_errstart(m, int32(15), v8230)
	mBase = m.M
	v8238 = m.ExcPending
	if v8238 != 0 {
		goto L32
	} else {
		goto L1836
	}
L1836:
	;
	if v8237 == int32(0) {
		goto L1445
	} else {
		goto L1837
	}
L1837:
	;
	v8242 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+84)) = uint32(v8242)
	v8245 = int64(base.Ui64(v8242) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+80)) = uint32(v8245)
	F_errmsg(m, int32(_a_F_StartupXLOG_236), v6432+int32(80))
	mBase = m.M
	v8251 = m.ExcPending
	if v8251 != 0 {
		goto L32
	} else {
		goto L1838
	}
L1838:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2804), int32(_a_F_StartupXLOG_235))
	mBase = m.M
	v8256 = m.ExcPending
	if v8256 != 0 {
		goto L32
	} else {
		goto L1839
	}
L1839:
	;
	goto L1445
L1840:
	;
	v8261 = v8009 & int32(112)
	v8262 = int32(4)
	v8263 = int32(base.Ui32(v8261) >> (uint(v8262) % 32))
	if base.B2i32(base.Ui32(v8262) < base.Ui32(v8263))|base.B2i32(v8263 == int32(1)) != 0 {
		v8610 = v8012
		goto L1841
	} else {
		goto L1842
	}
L1841:
	;
	if v8610 != int32(5) {
		goto L1631
	} else {
		goto L1915
	}
L1842:
	;
	v8270 = int32(4)
	v8273 = int32(base.Ui32(v8009)>>(uint(v8270)%32)) & int32(7)
	if base.B2i32(base.Ui32(v8270) < base.Ui32(v8273))|base.B2i32(v8273 == int32(1)) == int32(0) {
		goto L1843
	} else {
		goto L1844
	}
L1843:
	;
	v8282 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8283 = *(*int32)(unsafe.Add(mBase, uint32(v8282)+96))
	v8284 = *(*int32)(unsafe.Add(mBase, uint32(v8008)+64))
	v8285 = *(*int64)(unsafe.Add(mBase, uint32(v8284)))
	*(*int32)(unsafe.Add(mBase, uint32(v8282)+96)) = int32(1)
	if v8283 != 0 {
		goto L1846
	} else {
		goto L1847
	}
L1844:
	;
	v8304 = v8008
	v8305 = int64(0)
	goto L1845
L1845:
	;
	v8307 = v8261 - int32(48)
	if v8307 != 0 {
		goto L1853
	} else {
		goto L1854
	}
L1846:
	;
	v8289 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v8289+int32(96), int32(_a_F_StartupXLOG_50), int32(_a_F_StartupXLOG_237), int32(_a_F_StartupXLOG_238))
	mBase = m.M
	v8296 = m.ExcPending
	if v8296 != 0 {
		goto L32
	} else {
		goto L1849
	}
L1847:
	;
	goto L1848
L1848:
	;
	v8298 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v8298)+96)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8298)+64)) = v8285
	v8302 = *(*int32)(unsafe.Add(mBase, uint32(v8007)+96))
	v8304 = v8302
	v8305 = v8285
	goto L1845
L1849:
	;
	goto L1848
L1850:
	;
	v8537 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v8537 != int32(1) {
		v8610 = v8537
		goto L1841
	} else {
		goto L1900
	}
L1851:
	;
	v8534 = *(*int32)(unsafe.Add(mBase, uint32(v8304)+36))
	v8535 = v8534
	goto L1850
L1852:
	;
	v8429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8304)+48)))
	v8430 = *(*int32)(unsafe.Add(mBase, uint32(v8304)+64))
	v8432 = v6432 + int32(560)
	v8433 = int32(0)
	base.MemoryFill(m, v8432, v8433, int32(264))
	v8439 = *(*int64)(unsafe.Add(mBase, uint32(v8430)))
	*(*int64)(unsafe.Add(mBase, uint32(v8432))) = v8439
	if v8433 <= base.I32_extend8_s(v8429) {
		goto L1882
	} else {
		goto L1883
	}
L1853:
	;
	if v8307 == int32(16) {
		goto L1856
	} else {
		goto L1857
	}
L1854:
	;
	goto L1855
L1855:
	;
	v8310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8304)+48)))
	v8311 = *(*int32)(unsafe.Add(mBase, uint32(v8304)+64))
	v8313 = v6432 + int32(560)
	v8314 = int32(0)
	base.MemoryFill(m, v8313, v8314, int32(288))
	v8320 = *(*int64)(unsafe.Add(mBase, uint32(v8311)))
	*(*int64)(unsafe.Add(mBase, uint32(v8313))) = v8320
	if v8314 <= base.I32_extend8_s(v8310) {
		goto L1860
	} else {
		goto L1861
	}
L1856:
	;
	goto L1852
L1857:
	;
	goto L1851
L1859:
	;
	v8428 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+612))
	v8535 = v8428
	goto L1850
L1860:
	;
	goto L1859
L1861:
	;
	v8325 = *(*int32)(unsafe.Add(mBase, uint32(v8311)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+8)) = v8325
	if v8325&int32(1) != 0 {
		goto L1862
	} else {
		goto L1863
	}
L1862:
	;
	v8329 = *(*int32)(unsafe.Add(mBase, uint32(v8311)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+12)) = v8329
	v8331 = *(*int32)(unsafe.Add(mBase, uint32(v8311)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+16)) = v8331
	v8337 = v8311 + int32(20)
	goto L1864
L1863:
	;
	v8337 = v8311 + int32(12)
	goto L1864
L1864:
	;
	if v8325&int32(2) != 0 {
		goto L1865
	} else {
		goto L1866
	}
L1865:
	;
	v8340 = *(*int32)(unsafe.Add(mBase, uint32(v8337)))
	v8342 = v8337 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+24)) = v8342
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+20)) = v8340
	v8348 = v8342 + v8340<<(uint(int32(2))%32)
	goto L1867
L1866:
	;
	v8348 = v8337
	goto L1867
L1867:
	;
	if v8325&int32(4) != 0 {
		goto L1868
	} else {
		goto L1869
	}
L1868:
	;
	v8352 = *(*int32)(unsafe.Add(mBase, uint32(v8348)))
	v8354 = v8348 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+32)) = v8354
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+28)) = v8352
	v8357 = *(*int32)(unsafe.Add(mBase, uint32(v8348)))
	v8361 = v8354 + v8357*int32(12)
	goto L1870
L1869:
	;
	v8361 = v8348
	goto L1870
L1870:
	;
	if v8325&int32(256) != 0 {
		goto L1871
	} else {
		goto L1872
	}
L1871:
	;
	v8366 = *(*int32)(unsafe.Add(mBase, uint32(v8361)))
	v8367 = int32(4)
	v8368 = v8361 + v8367
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+40)) = v8368
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+36)) = v8366
	v8371 = *(*int32)(unsafe.Add(mBase, uint32(v8361)))
	v8375 = v8368 + v8371<<(uint(v8367)%32)
	goto L1873
L1872:
	;
	v8375 = v8361
	goto L1873
L1873:
	;
	if v8325&int32(8) != 0 {
		goto L1874
	} else {
		goto L1875
	}
L1874:
	;
	v8380 = *(*int32)(unsafe.Add(mBase, uint32(v8375)))
	v8381 = int32(4)
	v8382 = v8375 + v8381
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+48)) = v8382
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+44)) = v8380
	v8385 = *(*int32)(unsafe.Add(mBase, uint32(v8375)))
	v8389 = v8382 + v8385<<(uint(v8381)%32)
	goto L1876
L1875:
	;
	v8389 = v8375
	goto L1876
L1876:
	;
	if v8325&int32(16) == int32(0) {
		v8413 = v8325
		v8414 = v8389
		goto L1877
	} else {
		goto L1878
	}
L1877:
	;
	if v8413&int32(32) == int32(0) {
		goto L1860
	} else {
		goto L1880
	}
L1878:
	;
	v8396 = *(*int32)(unsafe.Add(mBase, uint32(v8389)))
	*(*int32)(unsafe.Add(mBase, uint32(v8313)+52)) = v8396
	v8399 = v8389 + int32(4)
	if v8325&int32(128) == int32(0) {
		v8413 = v8325
		v8414 = v8399
		goto L1877
	} else {
		goto L1879
	}
L1879:
	;
	v8407 = F_strlcpy(m, v6432+int32(616), v8399, int32(200))
	mBase = m.M
	v8408 = F_strlen(m, v8399)
	mBase = m.M
	v8412 = *(*int32)(unsafe.Add(mBase, uint32(v8313)+8))
	v8413 = v8412
	v8414 = v8408 + v8399 + int32(1)
	goto L1877
L1880:
	;
	v8419 = *(*int64)(unsafe.Add(mBase, uint32(v8414)))
	v8420 = *(*int64)(unsafe.Add(mBase, uint32(v8414)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8313)+280)) = v8420
	*(*int64)(unsafe.Add(mBase, uint32(v8313)+272)) = v8419
	goto L1860
L1881:
	;
	v8533 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+604))
	v8535 = v8533
	goto L1850
L1882:
	;
	goto L1881
L1883:
	;
	v8444 = *(*int32)(unsafe.Add(mBase, uint32(v8430)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+8)) = v8444
	if v8444&int32(1) != 0 {
		goto L1884
	} else {
		goto L1885
	}
L1884:
	;
	v8448 = *(*int32)(unsafe.Add(mBase, uint32(v8430)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+12)) = v8448
	v8450 = *(*int32)(unsafe.Add(mBase, uint32(v8430)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+16)) = v8450
	v8456 = v8430 + int32(20)
	goto L1886
L1885:
	;
	v8456 = v8430 + int32(12)
	goto L1886
L1886:
	;
	if v8444&int32(2) != 0 {
		goto L1887
	} else {
		goto L1888
	}
L1887:
	;
	v8459 = *(*int32)(unsafe.Add(mBase, uint32(v8456)))
	v8461 = v8456 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+24)) = v8461
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+20)) = v8459
	v8467 = v8461 + v8459<<(uint(int32(2))%32)
	goto L1889
L1888:
	;
	v8467 = v8456
	goto L1889
L1889:
	;
	if v8444&int32(4) != 0 {
		goto L1890
	} else {
		goto L1891
	}
L1890:
	;
	v8471 = *(*int32)(unsafe.Add(mBase, uint32(v8467)))
	v8473 = v8467 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+32)) = v8473
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+28)) = v8471
	v8476 = *(*int32)(unsafe.Add(mBase, uint32(v8467)))
	v8480 = v8473 + v8476*int32(12)
	goto L1892
L1891:
	;
	v8480 = v8467
	goto L1892
L1892:
	;
	if v8444&int32(256) != 0 {
		goto L1893
	} else {
		goto L1894
	}
L1893:
	;
	v8485 = *(*int32)(unsafe.Add(mBase, uint32(v8480)))
	v8486 = int32(4)
	v8487 = v8480 + v8486
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+40)) = v8487
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+36)) = v8485
	v8490 = *(*int32)(unsafe.Add(mBase, uint32(v8480)))
	v8494 = v8487 + v8490<<(uint(v8486)%32)
	goto L1895
L1894:
	;
	v8494 = v8480
	goto L1895
L1895:
	;
	if v8444&int32(16) == int32(0) {
		v8518 = v8444
		v8519 = v8494
		goto L1896
	} else {
		goto L1897
	}
L1896:
	;
	if v8518&int32(32) == int32(0) {
		goto L1882
	} else {
		goto L1899
	}
L1897:
	;
	v8501 = *(*int32)(unsafe.Add(mBase, uint32(v8494)))
	*(*int32)(unsafe.Add(mBase, uint32(v8432)+44)) = v8501
	v8504 = v8494 + int32(4)
	if v8444&int32(128) == int32(0) {
		v8518 = v8444
		v8519 = v8504
		goto L1896
	} else {
		goto L1898
	}
L1898:
	;
	v8512 = F_strlcpy(m, v6432+int32(608), v8504, int32(200))
	mBase = m.M
	v8513 = F_strlen(m, v8504)
	mBase = m.M
	v8517 = *(*int32)(unsafe.Add(mBase, uint32(v8432)+8))
	v8518 = v8517
	v8519 = v8513 + v8504 + int32(1)
	goto L1896
L1899:
	;
	v8524 = *(*int64)(unsafe.Add(mBase, uint32(v8519)))
	v8525 = *(*int64)(unsafe.Add(mBase, uint32(v8519)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8432)+256)) = v8525
	*(*int64)(unsafe.Add(mBase, uint32(v8432)+248)) = v8524
	goto L1882
L1900:
	;
	v8541 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[113])))
	if v8541&int32(1) == int32(0) {
		v8610 = v8537
		goto L1841
	} else {
		goto L1901
	}
L1901:
	;
	v8547 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[45]))
	if v8535 != v8547 {
		v8610 = v8537
		goto L1841
	} else {
		goto L1902
	}
L1902:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v8535
	v8552 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v8552)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v8305
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = int64(0)
	v8560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v8560)
	switch v8263 {
	case 0, 3:
		goto L1904
	default:
		goto L1445
	case 2, 4:
		goto L1903
	}
L1903:
	;
	v8588 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8589 = m.ExcPending
	if v8589 != 0 {
		goto L32
	} else {
		goto L1910
	}
L1904:
	;
	v8564 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8565 = m.ExcPending
	if v8565 != 0 {
		goto L32
	} else {
		goto L1905
	}
L1905:
	;
	if v8564 == int32(0) {
		goto L1445
	} else {
		goto L1906
	}
L1906:
	;
	v8569 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	v8571 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v8572 = F_timestamptz_to_str(m, v8571)
	mBase = m.M
	v8573 = m.ExcPending
	if v8573 != 0 {
		goto L32
	} else {
		goto L1907
	}
L1907:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+52)) = v8572
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+48)) = v8569
	F_errmsg(m, int32(_a_F_StartupXLOG_239), v6432+int32(48))
	mBase = m.M
	v8580 = m.ExcPending
	if v8580 != 0 {
		goto L32
	} else {
		goto L1908
	}
L1908:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2872), int32(_a_F_StartupXLOG_235))
	mBase = m.M
	v8585 = m.ExcPending
	if v8585 != 0 {
		goto L32
	} else {
		goto L1909
	}
L1909:
	;
	goto L1445
L1910:
	;
	if v8588 == int32(0) {
		goto L1445
	} else {
		goto L1911
	}
L1911:
	;
	v8593 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	v8595 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v8596 = F_timestamptz_to_str(m, v8595)
	mBase = m.M
	v8597 = m.ExcPending
	if v8597 != 0 {
		goto L32
	} else {
		goto L1912
	}
L1912:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+68)) = v8596
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+64)) = v8593
	F_errmsg(m, int32(_a_F_StartupXLOG_240), v6432-int32(-64))
	mBase = m.M
	v8604 = m.ExcPending
	if v8604 != 0 {
		goto L32
	} else {
		goto L1913
	}
L1913:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2880), int32(_a_F_StartupXLOG_235))
	mBase = m.M
	v8609 = m.ExcPending
	if v8609 != 0 {
		goto L32
	} else {
		goto L1914
	}
L1914:
	;
	goto L1445
L1915:
	;
	v8617 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[99])))
	if v8617&int32(1) == int32(0) {
		goto L1631
	} else {
		goto L1916
	}
L1916:
	;
	v8624 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v8625 = m.ExcPending
	if v8625 != 0 {
		goto L32
	} else {
		goto L1917
	}
L1917:
	;
	if v8624 != 0 {
		goto L1918
	} else {
		goto L1919
	}
L1918:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_221), int32(0))
	mBase = m.M
	v8629 = m.ExcPending
	if v8629 != 0 {
		goto L32
	} else {
		goto L1921
	}
L1919:
	;
	goto L1920
L1920:
	;
	v8636 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])) = uint8(v8636)
	v8639 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101])) = v8639
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102])) = v8639
	v8645 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100])) = v8645
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[104])) = uint8(v8645)
	goto L1445
L1921:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2890), int32(_a_F_StartupXLOG_235))
	mBase = m.M
	v8634 = m.ExcPending
	if v8634 != 0 {
		goto L32
	} else {
		goto L1922
	}
L1922:
	;
	goto L1920
L1923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+244)) = v6668
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+240)) = v7481
	F_errmsg(m, int32(_a_F_StartupXLOG_241), v6432+int32(240))
	mBase = m.M
	v8660 = m.ExcPending
	if v8660 != 0 {
		goto L32
	} else {
		goto L1924
	}
L1924:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2406), int32(_a_F_StartupXLOG_242))
	mBase = m.M
	v8665 = m.ExcPending
	if v8665 != 0 {
		goto L32
	} else {
		goto L1925
	}
L1925:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+212)) = v6668
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+208)) = v7479
	F_errmsg(m, int32(_a_F_StartupXLOG_243), v6432+int32(208))
	mBase = m.M
	v8677 = m.ExcPending
	if v8677 != 0 {
		goto L32
	} else {
		goto L1927
	}
L1927:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2415), int32(_a_F_StartupXLOG_242))
	mBase = m.M
	v8682 = m.ExcPending
	if v8682 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+224)) = v7479
	v8689 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[50]))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+232)) = uint32(v8689)
	v8692 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+236)) = v8692
	v8695 = int64(base.Ui64(v8689) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+228)) = uint32(v8695)
	F_errmsg(m, int32(_a_F_StartupXLOG_244), v6432+int32(224))
	mBase = m.M
	v8701 = m.ExcPending
	if v8701 != 0 {
		goto L32
	} else {
		goto L1930
	}
L1930:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2433), int32(_a_F_StartupXLOG_242))
	mBase = m.M
	v8706 = m.ExcPending
	if v8706 != 0 {
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
	v8711 = *(*int64)(unsafe.Add(mBase, uint32(v7461)+64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+188)) = uint32(v8711)
	v8713 = int64(32)
	v8714 = int64(base.Ui64(v8711) >> (uint(v8713) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+184)) = uint32(v8714)
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+180)) = uint32(v7582)
	v8718 = int64(base.Ui64(v7582) >> (uint(v8713) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+176)) = uint32(v8718)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_245), v6432+int32(176))
	mBase = m.M
	v8724 = m.ExcPending
	if v8724 != 0 {
		goto L32
	} else {
		goto L1933
	}
L1933:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2108), int32(_a_F_StartupXLOG_229))
	mBase = m.M
	v8729 = m.ExcPending
	if v8729 != 0 {
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
	F_errcode(m, int32(2600))
	mBase = m.M
	v8736 = m.ExcPending
	if v8736 != 0 {
		goto L32
	} else {
		goto L1936
	}
L1936:
	;
	v8737 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+1252))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+128)) = v8737
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_246), v6432+int32(128))
	mBase = m.M
	v8743 = m.ExcPending
	if v8743 != 0 {
		goto L32
	} else {
		goto L1937
	}
L1937:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2563), int32(_a_F_StartupXLOG_247))
	mBase = m.M
	v8748 = m.ExcPending
	if v8748 != 0 {
		goto L32
	} else {
		goto L1938
	}
L1938:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1939:
	;
	v8753 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+552))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+112)) = v8753
	v8755 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+560))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+96)) = v8755
	v8757 = *(*int64)(unsafe.Add(mBase, uint32(v6432)+564))
	*(*int64)(unsafe.Add(mBase, uint32(v6432)+100)) = v8757
	v8759 = *(*int32)(unsafe.Add(mBase, uint32(v6432)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+108)) = v8759
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_248), v6432+int32(96))
	mBase = m.M
	v8765 = m.ExcPending
	if v8765 != 0 {
		goto L32
	} else {
		goto L1940
	}
L1940:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2581), int32(_a_F_StartupXLOG_247))
	mBase = m.M
	v8770 = m.ExcPending
	if v8770 != 0 {
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
	if v8783 != 0 {
		v6668 = v7536
		v6673 = v8783
		goto L1485
	} else {
		goto L1943
	}
L1943:
	;
	goto L1486
L1944:
	;
	if v8787 == int32(0) {
		v9100 = v6558
		goto L1443
	} else {
		goto L1945
	}
L1945:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_249), int32(0))
	mBase = m.M
	v8794 = m.ExcPending
	if v8794 != 0 {
		goto L32
	} else {
		goto L1946
	}
L1946:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1909), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v8799 = m.ExcPending
	if v8799 != 0 {
		goto L32
	} else {
		goto L1947
	}
L1947:
	;
	v9100 = v6558
	goto L1443
L1948:
	;
	if v7107 != 0 {
		goto L1949
	} else {
		goto L1950
	}
L1949:
	;
	if v8818 == int32(0) {
		goto L1445
	} else {
		goto L1952
	}
L1950:
	;
	goto L1951
L1951:
	;
	if v8818 == int32(0) {
		goto L1445
	} else {
		goto L1956
	}
L1952:
	;
	v8823 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	v8825 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v8826 = F_timestamptz_to_str(m, v8825)
	mBase = m.M
	v8827 = m.ExcPending
	if v8827 != 0 {
		goto L32
	} else {
		goto L1953
	}
L1953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+276)) = v8826
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+272)) = v8823
	F_errmsg(m, int32(_a_F_StartupXLOG_250), v6432+int32(272))
	mBase = m.M
	v8834 = m.ExcPending
	if v8834 != 0 {
		goto L32
	} else {
		goto L1954
	}
L1954:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2727), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v8839 = m.ExcPending
	if v8839 != 0 {
		goto L32
	} else {
		goto L1955
	}
L1955:
	;
	goto L1445
L1956:
	;
	v8843 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	v8845 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v8846 = F_timestamptz_to_str(m, v8845)
	mBase = m.M
	v8847 = m.ExcPending
	if v8847 != 0 {
		goto L32
	} else {
		goto L1957
	}
L1957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+292)) = v8846
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+288)) = v8843
	F_errmsg(m, int32(_a_F_StartupXLOG_251), v6432+int32(288))
	mBase = m.M
	v8854 = m.ExcPending
	if v8854 != 0 {
		goto L32
	} else {
		goto L1958
	}
L1958:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(2734), int32(_a_F_StartupXLOG_222))
	mBase = m.M
	v8859 = m.ExcPending
	if v8859 != 0 {
		goto L32
	} else {
		goto L1959
	}
L1959:
	;
	goto L1445
L1960:
	;
	v8898 = int32(1)
	v8900 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[18]))
	switch v8900 {
	case 0:
		goto L1961
	default:
		v8937 = v8898
		goto L1444
	case 2:
		goto L1962
	}
L1961:
	;
	v8905 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8906 = *(*int32)(unsafe.Add(mBase, uint32(v8905)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v8905)+96)) = int32(1)
	if v8906 != 0 {
		goto L1964
	} else {
		goto L1965
	}
L1962:
	;
	F_proc_exit(m, int32(3))
	mBase = m.M
	v8903 = m.ExcPending
	if v8903 != 0 {
		goto L32
	} else {
		goto L1963
	}
L1963:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1964:
	;
	v8910 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v8910+int32(96), int32(_a_F_StartupXLOG_50), int32(3114), int32(_a_F_StartupXLOG_252))
	mBase = m.M
	v8917 = m.ExcPending
	if v8917 != 0 {
		goto L32
	} else {
		goto L1967
	}
L1965:
	;
	goto L1966
L1966:
	;
	v8919 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v8920 = *(*int32)(unsafe.Add(mBase, uint32(v8919)+80))
	if v8920 == int32(0) {
		goto L1968
	} else {
		goto L1969
	}
L1967:
	;
	goto L1966
L1968:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8919)+80)) = int32(1)
	goto L1970
L1969:
	;
	goto L1970
L1970:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8919)+96)) = int32(0)
	F_recoveryPausesHere(m, int32(1))
	mBase = m.M
	v8929 = m.ExcPending
	if v8929 != 0 {
		goto L32
	} else {
		goto L1971
	}
L1971:
	;
	v8937 = v8898
	goto L1444
L1972:
	;
	v9000 = v8968 << (uint(int32(5)) % 32)
	v9003 = *(*int32)(unsafe.Add(mBase, uint32(v9000)+uint32(_c_F_StartupXLOG[106])))
	if v9003 == int32(0) {
		goto L1974
	} else {
		goto L1975
	}
L1973:
	;
	v9027 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9028 = m.ExcPending
	if v9028 != 0 {
		goto L32
	} else {
		goto L1983
	}
L1974:
	;
	v9012 = *(*int32)(unsafe.Add(mBase, uint32(v9000)+uint32(_c_F_StartupXLOG[107])))
	if v9012 == int32(0) {
		goto L1978
	} else {
		goto L1979
	}
L1975:
	;
	v9006 = *(*int32)(unsafe.Add(mBase, uint32(v9000)+uint32(_c_F_StartupXLOG[121])))
	if v9006 == int32(0) {
		goto L1974
	} else {
		goto L1976
	}
L1976:
	;
	m.T0[v9006].(func(*base.Module))(m)
	mBase = m.M
	v9010 = m.ExcPending
	if v9010 != 0 {
		goto L32
	} else {
		goto L1977
	}
L1977:
	;
	goto L1974
L1978:
	;
	v9022 = v8968 + int32(2)
	if v9022 != int32(256) {
		v8968 = v9022
		goto L1972
	} else {
		goto L1982
	}
L1979:
	;
	v9015 = *(*int32)(unsafe.Add(mBase, uint32(v9000)+uint32(_c_F_StartupXLOG[122])))
	if v9015 == int32(0) {
		goto L1978
	} else {
		goto L1980
	}
L1980:
	;
	m.T0[v9015].(func(*base.Module))(m)
	mBase = m.M
	v9019 = m.ExcPending
	if v9019 != 0 {
		goto L32
	} else {
		goto L1981
	}
L1981:
	;
	goto L1978
L1982:
	;
	goto L1973
L1983:
	;
	if v9027 != 0 {
		goto L1984
	} else {
		goto L1985
	}
L1984:
	;
	v9030 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9031 = *(*int64)(unsafe.Add(mBase, uint32(v9030)+32))
	v9034 = F_pg_rusage_show(m, v6432+int32(368))
	mBase = m.M
	v9035 = m.ExcPending
	if v9035 != 0 {
		goto L32
	} else {
		goto L1987
	}
L1985:
	;
	goto L1986
L1986:
	;
	v9053 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v9054 = *(*int32)(unsafe.Add(mBase, uint32(v9053)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v9053)+96)) = int32(1)
	if v9054 != 0 {
		goto L1990
	} else {
		goto L1991
	}
L1987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432)+24)) = v9034
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+20)) = uint32(v9031)
	v9039 = int64(base.Ui64(v9031) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6432)+16)) = uint32(v9039)
	F_errmsg(m, int32(_a_F_StartupXLOG_253), v6432+int32(16))
	mBase = m.M
	v9045 = m.ExcPending
	if v9045 != 0 {
		goto L32
	} else {
		goto L1988
	}
L1988:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1896), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9050 = m.ExcPending
	if v9050 != 0 {
		goto L32
	} else {
		goto L1989
	}
L1989:
	;
	goto L1986
L1990:
	;
	v9058 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	F_s_lock(m, v9058+int32(96), int32(_a_F_StartupXLOG_50), int32(_a_F_StartupXLOG_254), int32(_a_F_StartupXLOG_255))
	mBase = m.M
	v9065 = m.ExcPending
	if v9065 != 0 {
		goto L32
	} else {
		goto L1993
	}
L1991:
	;
	goto L1992
L1992:
	;
	v9067 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v9067)+96)) = int32(0)
	v9070 = *(*int64)(unsafe.Add(mBase, uint32(v9067)+64))
	if v9070 == int64(0) {
		goto L1994
	} else {
		goto L1995
	}
L1993:
	;
	goto L1992
L1994:
	;
	v9091 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[105])) = uint8(v9091)
	v9100 = v8937
	goto L1443
L1995:
	;
	v9075 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9076 = m.ExcPending
	if v9076 != 0 {
		goto L32
	} else {
		goto L1996
	}
L1996:
	;
	if v9075 == int32(0) {
		goto L1994
	} else {
		goto L1997
	}
L1997:
	;
	v9079 = F_timestamptz_to_str(m, v9070)
	mBase = m.M
	v9080 = m.ExcPending
	if v9080 != 0 {
		goto L32
	} else {
		goto L1998
	}
L1998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6432))) = v9079
	F_errmsg(m, int32(_a_F_StartupXLOG_256), v6432)
	mBase = m.M
	v9084 = m.ExcPending
	if v9084 != 0 {
		goto L32
	} else {
		goto L1999
	}
L1999:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1901), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9089 = m.ExcPending
	if v9089 != 0 {
		goto L32
	} else {
		goto L2000
	}
L2000:
	;
	goto L1994
L2001:
	;
	m.G0 = v6432 + int32(848)
	goto L1440
L2002:
	;
	v9128 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9128&int32(1) == int32(0) {
		goto L2001
	} else {
		goto L2003
	}
L2003:
	;
	v9134 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	if v9134 != 0 {
		goto L1441
	} else {
		goto L2004
	}
L2004:
	;
	goto L2001
L2005:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_257), int32(0))
	mBase = m.M
	v9145 = m.ExcPending
	if v9145 != 0 {
		goto L32
	} else {
		goto L2006
	}
L2006:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1862), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9150 = m.ExcPending
	if v9150 != 0 {
		goto L32
	} else {
		goto L2007
	}
L2007:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2008:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v9157 = m.ExcPending
	if v9157 != 0 {
		goto L32
	} else {
		goto L2009
	}
L2009:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_258), int32(0))
	mBase = m.M
	v9161 = m.ExcPending
	if v9161 != 0 {
		goto L32
	} else {
		goto L2010
	}
L2010:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_50), int32(1921), int32(_a_F_StartupXLOG_214))
	mBase = m.M
	v9166 = m.ExcPending
	if v9166 != 0 {
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
	F_XLogShutdownWalRcv(m)
	mBase = m.M
	v9210 = m.ExcPending
	if v9210 != 0 {
		goto L32
	} else {
		goto L2013
	}
L2013:
	;
	v9214 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	v9215 = *(*int32)(unsafe.Add(mBase, uint32(v9214)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9214)+16)) = int32(1)
	if v9215 != 0 {
		goto L2014
	} else {
		goto L2015
	}
L2014:
	;
	v9219 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	F_s_lock(m, v9219+int32(16), int32(_a_F_StartupXLOG_259), int32(1589), int32(_a_F_StartupXLOG_260))
	mBase = m.M
	v9226 = m.ExcPending
	if v9226 != 0 {
		goto L32
	} else {
		goto L2017
	}
L2015:
	;
	goto L2016
L2016:
	;
	v9228 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	v9229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9228)+4)) = uint8(v9229)
	v9231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9228)+5)))
	if v9231 != v9229 {
		v9319 = v9228
		goto L2018
	} else {
		goto L2019
	}
L2017:
	;
	goto L2016
L2018:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9319)+16)) = int32(0)
	v9355 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])))
	if v9355 == int32(1) {
		goto L2037
	} else {
		goto L2038
	}
L2019:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9228)+16)) = int32(0)
	v9236 = *(*int32)(unsafe.Add(mBase, uint32(v9228)))
	if v9236 != int32(-1) {
		goto L2020
	} else {
		goto L2021
	}
L2020:
	;
	v9240 = F_kill(m, v9236, int32(10))
	mBase = m.M
	v9241 = m.ExcPending
	if v9241 != 0 {
		goto L32
	} else {
		goto L2023
	}
L2021:
	;
	goto L2022
L2022:
	;
	goto L2024
L2023:
	;
	goto L2022
L2024:
	;
	v9277 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	v9281 = F_WaitLatch(m, v9277, int32(41), int32(10), int32(83886092))
	mBase = m.M
	v9282 = m.ExcPending
	if v9282 != 0 {
		goto L32
	} else {
		goto L2027
	}
L2026:
	;
	v9298 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	v9299 = *(*int32)(unsafe.Add(mBase, uint32(v9298)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9298)+16)) = int32(1)
	if v9299 != 0 {
		goto L2032
	} else {
		goto L2033
	}
L2027:
	;
	if v9281&int32(1) == int32(0) {
		goto L2026
	} else {
		goto L2028
	}
L2028:
	;
	v9288 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[124]))
	*(*int32)(unsafe.Add(mBase, uint32(v9288))) = int32(0)
	goto L2029
L2029:
	;
	v9292 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[79]))
	if v9292 == int32(0) {
		goto L2026
	} else {
		goto L2030
	}
L2030:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v9296 = m.ExcPending
	if v9296 != 0 {
		goto L32
	} else {
		goto L2031
	}
L2031:
	;
	goto L2026
L2032:
	;
	v9303 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	F_s_lock(m, v9303+int32(16), int32(_a_F_StartupXLOG_259), int32(1631), int32(_a_F_StartupXLOG_260))
	mBase = m.M
	v9310 = m.ExcPending
	if v9310 != 0 {
		goto L32
	} else {
		goto L2035
	}
L2033:
	;
	goto L2034
L2034:
	;
	v9312 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[123]))
	v9313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9312)+5)))
	if v9313 != int32(1) {
		v9319 = v9312
		goto L2018
	} else {
		goto L2036
	}
L2035:
	;
	goto L2034
L2036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9312)+16)) = int32(0)
	goto L2024
L2037:
	;
	v9359 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v9363 = F_LWLockAcquire(m, v9359+int32(_a_F_StartupXLOG_261), int32(1))
	mBase = m.M
	v9364 = m.ExcPending
	if v9364 != 0 {
		goto L32
	} else {
		goto L2040
	}
L2038:
	;
	goto L2039
L2039:
	;
	v9534 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[41])) = uint8(v9534)
	v9537 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	v9542 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v9542 != 0 {
		goto L2062
	} else {
		goto L2063
	}
L2040:
	;
	v9366 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	if int32(0) < v9366 {
		goto L2041
	} else {
		goto L2042
	}
L2041:
	;
	v9370 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9377 = v9366
	v9378 = int32(0)
	v9380 = v9370
	v9398 = int64(0)
	goto L2044
L2042:
	;
	goto L2043
L2043:
	;
	v9494 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v9494+int32(_a_F_StartupXLOG_261))
	mBase = m.M
	v9498 = m.ExcPending
	if v9498 != 0 {
		goto L32
	} else {
		goto L2061
	}
L2044:
	;
	v9407 = v9380 + v9378*int32(288)
	v9408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9407)+4)))
	if v9408 != int32(1) {
		v9453 = v9377
		v9454 = v9380
		v9455 = v9398
		goto L2046
	} else {
		goto L2047
	}
L2045:
	;
	goto L2043
L2046:
	;
	v9457 = v9378 + int32(1)
	if v9457 < v9453 {
		v9377 = v9453
		v9378 = v9457
		v9380 = v9454
		v9398 = v9455
		goto L2044
	} else {
		goto L2060
	}
L2047:
	;
	v9411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9407)+201)))
	if v9411 == int32(0) {
		v9453 = v9377
		v9454 = v9380
		v9455 = v9398
		goto L2046
	} else {
		goto L2048
	}
L2048:
	;
	if v9398 == int64(0) {
		goto L2049
	} else {
		goto L2050
	}
L2049:
	;
	v9419 = m.G0
	v9420 = int32(16)
	v9421 = v9419 - v9420
	m.G0 = v9421
	F_gettimeofday(m, v9421)
	mBase = m.M
	v9424 = *(*int64)(unsafe.Add(mBase, uint32(v9421)))
	v9425 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9421)+8)))
	m.G0 = v9421 + v9420
	goto L2052
L2050:
	;
	v9434 = v9398
	goto L2051
L2051:
	;
	v9435 = *(*int32)(unsafe.Add(mBase, uint32(v9407)))
	*(*int32)(unsafe.Add(mBase, uint32(v9407))) = int32(1)
	if v9435 != 0 {
		goto L2053
	} else {
		goto L2054
	}
L2052:
	;
	v9434 = v9425 + v9424*int64(1000000) - int64(946684800000000)
	goto L2051
L2053:
	;
	F_s_lock(m, v9407, int32(_a_F_StartupXLOG_262), int32(251), int32(_a_F_StartupXLOG_263))
	mBase = m.M
	v9442 = m.ExcPending
	if v9442 != 0 {
		goto L32
	} else {
		goto L2056
	}
L2054:
	;
	goto L2055
L2055:
	;
	v9443 = *(*int32)(unsafe.Add(mBase, uint32(v9407)+112))
	if v9443 == int32(0) {
		goto L2057
	} else {
		goto L2058
	}
L2056:
	;
	goto L2055
L2057:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9407)+272)) = v9434
	goto L2059
L2058:
	;
	goto L2059
L2059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9407))) = int32(0)
	v9450 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[62]))
	v9452 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[63]))
	v9453 = v9450
	v9454 = v9452
	v9455 = v9434
	goto L2046
L2060:
	;
	goto L2045
L2061:
	;
	goto L2039
L2062:
	;
	v9543 = v9537 + int32(40)
	goto L2064
L2063:
	;
	v9543 = int32(_a_F_StartupXLOG_216)
	goto L2064
L2064:
	;
	v9544 = *(*int32)(unsafe.Add(mBase, uint32(v9543)))
	v9546 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	if v9542 != 0 {
		goto L2065
	} else {
		goto L2066
	}
L2065:
	;
	v9550 = v9537 + int32(24)
	goto L2067
L2066:
	;
	v9550 = int32(_a_F_StartupXLOG_264)
	goto L2067
L2067:
	;
	v9551 = *(*int64)(unsafe.Add(mBase, uint32(v9550)))
	F_XLogPrefetcherBeginRead(m, v9546, v9551)
	mBase = m.M
	v9553 = m.ExcPending
	if v9553 != 0 {
		goto L32
	} else {
		goto L2068
	}
L2068:
	;
	v9555 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v9558 = F_ReadRecord(m, v9555, int32(23), int32(0), v9544)
	mBase = m.M
	v9559 = m.ExcPending
	if v9559 != 0 {
		goto L32
	} else {
		goto L2069
	}
L2069:
	;
	v9561 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9562 = *(*int64)(unsafe.Add(mBase, uint32(v9561)+40))
	v9563 = *(*int32)(unsafe.Add(mBase, uint32(v9561)+1184))
	*(*int32)(unsafe.Add(mBase, uint32(v9207)+24)) = v9563
	v9566 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9566 != int32(1) {
		goto L2070
	} else {
		goto L2071
	}
L2070:
	;
	v9582 = v9562 & int64(8191)
	if v9582 != int64(0) {
		goto L2073
	} else {
		goto L2074
	}
L2071:
	;
	v9570 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[40])) = uint8(v9570)
	v9573 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	if v9573 < v9570 {
		goto L2070
	} else {
		goto L2072
	}
L2072:
	;
	v9576 = F_close(m, v9573)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125])) = int32(-1)
	goto L2070
L2073:
	;
	v9585 = base.I32_wrap_i64(v9582)
	v9586 = F_palloc(m, v9585)
	mBase = m.M
	v9587 = m.ExcPending
	if v9587 != 0 {
		goto L32
	} else {
		goto L2076
	}
L2074:
	;
	v9595 = int32(0)
	v9596 = v9562
	goto L2075
L2075:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9207)+40)) = v9595
	*(*int64)(unsafe.Add(mBase, uint32(v9207)+32)) = v9596
	v9600 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[21]))
	switch v9600 - int32(1) {
	case 0:
		goto L2086
	case 1:
		goto L2085
	case 2:
		goto L2083
	case 3:
		goto L2084
	case 4:
		goto L2082
	default:
		goto L2081
	}
L2076:
	;
	if v9585 != 0 {
		goto L2077
	} else {
		goto L2078
	}
L2077:
	;
	v9589 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	v9590 = *(*int32)(unsafe.Add(mBase, uint32(v9589)+128))
	base.MemoryCopy(m, v9586, v9590, v9585)
	goto L2079
L2078:
	;
	goto L2079
L2079:
	;
	v9595 = v9586
	v9596 = v9562 & int64(-8192)
	goto L2075
L2080:
	;
	v9685 = F_pstrdup(m, v9204-int32(-64))
	mBase = m.M
	v9686 = m.ExcPending
	if v9686 != 0 {
		goto L32
	} else {
		goto L2103
	}
L2081:
	;
	v9679 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_265), int32(0))
	mBase = m.M
	v9680 = m.ExcPending
	if v9680 != 0 {
		goto L32
	} else {
		goto L2102
	}
L2082:
	;
	v9672 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_266), int32(0))
	mBase = m.M
	v9673 = m.ExcPending
	if v9673 != 0 {
		goto L32
	} else {
		goto L2101
	}
L2083:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9204)+48)) = int32(_a_F_StartupXLOG_233)
	v9665 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_267), v9204+int32(48))
	mBase = m.M
	v9666 = m.ExcPending
	if v9666 != 0 {
		goto L32
	} else {
		goto L2100
	}
L2084:
	;
	v9638 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[102]))
	*(*uint32)(unsafe.Add(mBase, uint32(v9204)+40)) = uint32(v9638)
	v9643 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])))
	if v9643 != 0 {
		goto L2096
	} else {
		goto L2097
	}
L2085:
	;
	v9619 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])))
	v9621 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[101]))
	v9622 = F_timestamptz_to_str(m, v9621)
	mBase = m.M
	v9623 = m.ExcPending
	if v9623 != 0 {
		goto L32
	} else {
		goto L2091
	}
L2086:
	;
	v9604 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v9204)+4)) = v9604
	v9609 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[103])))
	if v9609 != 0 {
		goto L2087
	} else {
		goto L2088
	}
L2087:
	;
	v9610 = int32(_a_F_StartupXLOG_268)
	goto L2089
L2088:
	;
	v9610 = int32(_a_F_StartupXLOG_269)
	goto L2089
L2089:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9204))) = v9610
	v9616 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_270), v9204)
	mBase = m.M
	v9617 = m.ExcPending
	if v9617 != 0 {
		goto L32
	} else {
		goto L2090
	}
L2090:
	;
	goto L2080
L2091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9204)+20)) = v9622
	if v9619 != 0 {
		goto L2092
	} else {
		goto L2093
	}
L2092:
	;
	v9627 = int32(_a_F_StartupXLOG_268)
	goto L2094
L2093:
	;
	v9627 = int32(_a_F_StartupXLOG_269)
	goto L2094
L2094:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9204)+16)) = v9627
	v9635 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_271), v9204+int32(16))
	mBase = m.M
	v9636 = m.ExcPending
	if v9636 != 0 {
		goto L32
	} else {
		goto L2095
	}
L2095:
	;
	goto L2080
L2096:
	;
	v9644 = int32(_a_F_StartupXLOG_268)
	goto L2098
L2097:
	;
	v9644 = int32(_a_F_StartupXLOG_269)
	goto L2098
L2098:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9204)+32)) = v9644
	v9647 = int64(base.Ui64(v9638) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v9204)+36)) = uint32(v9647)
	v9655 = F_pg_snprintf(m, v9204-int32(-64), int32(200), int32(_a_F_StartupXLOG_272), v9204+int32(32))
	mBase = m.M
	v9656 = m.ExcPending
	if v9656 != 0 {
		goto L32
	} else {
		goto L2099
	}
L2099:
	;
	goto L2080
L2100:
	;
	goto L2080
L2101:
	;
	goto L2080
L2102:
	;
	goto L2080
L2103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9207)+16)) = v9562
	*(*int32)(unsafe.Add(mBase, uint32(v9207)+8)) = v9544
	*(*int64)(unsafe.Add(mBase, uint32(v9207))) = v9551
	*(*int32)(unsafe.Add(mBase, uint32(v9207)+64)) = v9685
	v9692 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[51]))
	*(*int64)(unsafe.Add(mBase, uint32(v9207)+48)) = v9692
	v9695 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[52]))
	*(*int64)(unsafe.Add(mBase, uint32(v9207)+56)) = v9695
	v9698 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[16])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9207)+68)) = uint8(v9698)
	v9701 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[17])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9207)+69)) = uint8(v9701)
	m.G0 = v9204 + int32(272)
	v9708 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+24))
	v9711 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])))
	if v9711 == int32(1) {
		goto L2104
	} else {
		goto L2105
	}
L2104:
	;
	v9715 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v9717 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[83]))
	if base.Ui64(v9717) <= base.Ui64(v9562) {
		goto L2108
	} else {
		goto L2109
	}
L2105:
	;
	goto L2106
L2106:
	;
	v9755 = int32(0)
	v9757 = F_PrescanPreparedTransactions(m, v9755, v9755)
	mBase = m.M
	v9758 = m.ExcPending
	if v9758 != 0 {
		goto L32
	} else {
		goto L2123
	}
L2107:
	;
	F_ResetUnloggedRelations(m, int32(2))
	mBase = m.M
	v9753 = m.ExcPending
	if v9753 != 0 {
		goto L32
	} else {
		goto L2122
	}
L2108:
	;
	v9719 = *(*int64)(unsafe.Add(mBase, uint32(v9715)+152))
	if v9719 == int64(0) {
		goto L2107
	} else {
		goto L2111
	}
L2109:
	;
	goto L2110
L2110:
	;
	v9723 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9723 == int32(0) {
		goto L2112
	} else {
		goto L2113
	}
L2111:
	;
	goto L2110
L2112:
	;
	v9726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9715)+168)))
	if v9726 != int32(1) {
		goto L2107
	} else {
		goto L2115
	}
L2113:
	;
	goto L2114
L2114:
	;
	v9729 = *(*int64)(unsafe.Add(mBase, uint32(v9715)+152))
	if v9729 != int64(0) {
		goto L13
	} else {
		goto L2116
	}
L2115:
	;
	goto L2114
L2116:
	;
	v9732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9715)+168)))
	if v9732 == int32(1) {
		goto L13
	} else {
		goto L2117
	}
L2117:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v9738 = m.ExcPending
	if v9738 != 0 {
		goto L32
	} else {
		goto L2118
	}
L2118:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v9741 = m.ExcPending
	if v9741 != 0 {
		goto L32
	} else {
		goto L2119
	}
L2119:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_273), int32(0))
	mBase = m.M
	v9745 = m.ExcPending
	if v9745 != 0 {
		goto L32
	} else {
		goto L2120
	}
L2120:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_274), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v9750 = m.ExcPending
	if v9750 != 0 {
		goto L32
	} else {
		goto L2121
	}
L2121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2122:
	;
	goto L2106
L2123:
	;
	v9760 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v9764 = F_LWLockAcquire(m, v9760+int32(1152), int32(0))
	mBase = m.M
	v9765 = m.ExcPending
	if v9765 != 0 {
		goto L32
	} else {
		goto L2124
	}
L2124:
	;
	v9767 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v9768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9767)+320)) = uint8(v9768)
	v9771 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v9771+int32(1152))
	mBase = m.M
	v9775 = m.ExcPending
	if v9775 != 0 {
		goto L32
	} else {
		goto L2125
	}
L2125:
	;
	v9777 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v9777 != int32(1) {
		goto L2127
	} else {
		goto L2128
	}
L2126:
	;
	v10681 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10682 = *(*int32)(unsafe.Add(mBase, uint32(v10681)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v10681)+440)) = int32(1)
	if v10682 != 0 {
		goto L2315
	} else {
		goto L2316
	}
L2127:
	;
	v9780 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+8))
	v10651 = v9780
	goto L2126
L2128:
	;
	goto L2129
L2129:
	;
	v9782 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v9783 = F_findNewestTimeLine(m, v9782)
	mBase = m.M
	v9784 = m.ExcPending
	if v9784 != 0 {
		goto L32
	} else {
		goto L2130
	}
L2130:
	;
	v9786 = v9783 + int32(1)
	v9789 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v9790 = m.ExcPending
	if v9790 != 0 {
		goto L32
	} else {
		goto L2131
	}
L2131:
	;
	if v9789 != 0 {
		goto L2132
	} else {
		goto L2133
	}
L2132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3840)) = v9786
	F_errmsg(m, int32(_a_F_StartupXLOG_275), v39+int32(3840))
	mBase = m.M
	v9796 = m.ExcPending
	if v9796 != 0 {
		goto L32
	} else {
		goto L2135
	}
L2133:
	;
	goto L2134
L2134:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v9805 = m.ExcPending
	if v9805 != 0 {
		goto L32
	} else {
		goto L2137
	}
L2135:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_276), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v9801 = m.ExcPending
	if v9801 != 0 {
		goto L32
	} else {
		goto L2136
	}
L2136:
	;
	goto L2134
L2137:
	;
	v9809 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v9810 = base.I64_extend_i32_s(v9809)
	v9811 = base.I64_div_u_s(v9562-int64(1), v9810)
	v9812 = base.I64_div_u_s(v9562, v9810)
	if v9811 == v9812 {
		goto L2139
	} else {
		goto L2140
	}
L2138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3680)) = v9786
	v10120 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v10121 = base.I64_div_u_s(int64(4294967296), v10120)
	v10122 = base.I64_div_u_s(v9812, v10121)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3684)) = uint32(v10122)
	v10125 = v9812 - v10121*v10122
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3688)) = uint32(v10125)
	v10128 = v39 + int32(_a_F_StartupXLOG_1)
	v10133 = F_pg_snprintf(m, v10128, int32(64), int32(_a_F_StartupXLOG_277), v39+int32(3680))
	mBase = m.M
	v10134 = m.ExcPending
	if v10134 != 0 {
		goto L32
	} else {
		goto L2201
	}
L2139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3808)) = v9708
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[126]))) = v9811
	v9817 = base.I64_div_u_s(int64(4294967296), v9810)
	v9818 = base.I64_div_u_s(v9811, v9817)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3812)) = uint32(v9818)
	v9821 = v9811 - v9817*v9818
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3816)) = uint32(v9821)
	v9824 = v39 + int32(_a_F_StartupXLOG_4)
	v9829 = F_pg_snprintf(m, v9824, int32(1024), int32(_a_F_StartupXLOG_278), v39+int32(3808))
	mBase = m.M
	v9830 = m.ExcPending
	if v9830 != 0 {
		goto L32
	} else {
		goto L2142
	}
L2140:
	;
	goto L2141
L2141:
	;
	v10080 = F_XLogFileInit(m, v9812, v9786)
	mBase = m.M
	v10081 = m.ExcPending
	if v10081 != 0 {
		goto L32
	} else {
		goto L2199
	}
L2142:
	;
	v9832 = F_OpenTransientFile(m, v9824, int32(0))
	mBase = m.M
	v9833 = m.ExcPending
	if v9833 != 0 {
		goto L32
	} else {
		goto L2143
	}
L2143:
	;
	if v9832 < int32(0) {
		goto L12
	} else {
		goto L2144
	}
L2144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3792)) = int32(42)
	v9839 = v39 + int32(_a_F_StartupXLOG_3)
	v9844 = F_pg_snprintf(m, v9839, int32(1024), int32(_a_F_StartupXLOG_279), v39+int32(3792))
	mBase = m.M
	v9845 = m.ExcPending
	if v9845 != 0 {
		goto L32
	} else {
		goto L2145
	}
L2145:
	;
	v9846 = F_unlink(m, v9839)
	mBase = m.M
	v9849 = F_OpenTransientFile(m, v9839, int32(194))
	mBase = m.M
	v9850 = m.ExcPending
	if v9850 != 0 {
		goto L32
	} else {
		goto L2146
	}
L2146:
	;
	if v9849 < int32(0) {
		goto L11
	} else {
		goto L2147
	}
L2147:
	;
	v9854 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if int32(0) < v9854 {
		goto L2148
	} else {
		goto L2149
	}
L2148:
	;
	v9861 = int32(0)
	goto L2151
L2149:
	;
	goto L2150
L2150:
	;
	v10004 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10004))) = int32(167772231)
	v10009 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v10009 != int32(1) {
		v10023 = int32(0)
		goto L2174
	} else {
		goto L2175
	}
L2151:
	;
	v9895 = base.I32_wrap_i64(v9562)&(v9809-int32(1)) - v9861
	if base.Ui32(v9895) <= base.Ui32(int32(_a_F_StartupXLOG_280)) {
		goto L2153
	} else {
		goto L2154
	}
L2152:
	;
	goto L2150
L2153:
	;
	base.MemoryFill(m, v39+int32(_a_F_StartupXLOG_1), int32(0), int32(_a_F_StartupXLOG_58))
	goto L2155
L2154:
	;
	goto L2155
L2155:
	;
	if int32(0) < v9895 {
		goto L2156
	} else {
		goto L2157
	}
L2156:
	;
	v9906 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9906))) = int32(167772230)
	v9911 = int32(_a_F_StartupXLOG_58)
	if base.Ui32(v9911) <= base.Ui32(v9895) {
		goto L2159
	} else {
		goto L2160
	}
L2157:
	;
	goto L2158
L2158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v9951 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9951))) = int32(167772232)
	v9956 = int32(_a_F_StartupXLOG_58)
	v9957 = F_write(m, v9849, v39+int32(_a_F_StartupXLOG_1), v9956)
	mBase = m.M
	if v9957 != v9956 {
		goto L9
	} else {
		goto L2170
	}
L2159:
	;
	v9914 = v9911
	goto L2161
L2160:
	;
	v9914 = v9895
	goto L2161
L2161:
	;
	v9915 = F_read(m, v9832, v39+int32(_a_F_StartupXLOG_1), v9914)
	mBase = m.M
	if v9915 != v9914 {
		goto L2162
	} else {
		goto L2163
	}
L2162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v9920 = m.ExcPending
	if v9920 != 0 {
		goto L32
	} else {
		goto L2165
	}
L2163:
	;
	goto L2164
L2164:
	;
	v9942 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9942))) = int32(0)
	goto L2158
L2165:
	;
	if v9915 < int32(0) {
		goto L10
	} else {
		goto L2166
	}
L2166:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v9925 = m.ExcPending
	if v9925 != 0 {
		goto L32
	} else {
		goto L2167
	}
L2167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3784)) = v9914
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3780)) = v9915
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3776)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_144), v39+int32(3776))
	mBase = m.M
	v9935 = m.ExcPending
	if v9935 != 0 {
		goto L32
	} else {
		goto L2168
	}
L2168:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3486), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v9940 = m.ExcPending
	if v9940 != 0 {
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
	v9961 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v9961))) = int32(0)
	v9965 = v9861 - int32(-8192)
	v9967 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v9965 < v9967 {
		v9861 = v9965
		goto L2151
	} else {
		goto L2171
	}
L2171:
	;
	goto L2152
L2172:
	;
	v10052 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10052))) = int32(0)
	v10055 = F_CloseTransientFile(m, v9849)
	mBase = m.M
	v10056 = m.ExcPending
	if v10056 != 0 {
		goto L32
	} else {
		goto L2190
	}
L2173:
	;
	if v10023 == int32(0) {
		goto L2172
	} else {
		goto L2180
	}
L2174:
	;
	goto L2173
L2175:
	;
	goto L2176
L2176:
	;
	v10014 = F_fsync(m, v9849)
	mBase = m.M
	if v10014 != int32(-1) {
		v10023 = v10014
		goto L2174
	} else {
		goto L2178
	}
L2177:
	;
	v10023 = int32(-1)
	goto L2174
L2178:
	;
	v10018 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10018 == int32(27) {
		goto L2176
	} else {
		goto L2179
	}
L2179:
	;
	goto L2177
L2180:
	;
	v10029 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[127])))
	if v10029 != 0 {
		goto L2182
	} else {
		goto L2183
	}
L2181:
	;
	v10032 = F_errstart(m, v10030, int32(0))
	mBase = m.M
	v10033 = m.ExcPending
	if v10033 != 0 {
		goto L32
	} else {
		goto L2185
	}
L2182:
	;
	v10030 = int32(21)
	goto L2184
L2183:
	;
	v10030 = int32(23)
	goto L2184
L2184:
	;
	goto L2181
L2185:
	;
	if v10032 == int32(0) {
		goto L2172
	} else {
		goto L2186
	}
L2186:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10037 = m.ExcPending
	if v10037 != 0 {
		goto L32
	} else {
		goto L2187
	}
L2187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3728)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v39+int32(3728))
	mBase = m.M
	v10045 = m.ExcPending
	if v10045 != 0 {
		goto L32
	} else {
		goto L2188
	}
L2188:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3514), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v10050 = m.ExcPending
	if v10050 != 0 {
		goto L32
	} else {
		goto L2189
	}
L2189:
	;
	goto L2172
L2190:
	;
	if v10055 != 0 {
		goto L8
	} else {
		goto L2191
	}
L2191:
	;
	v10057 = F_CloseTransientFile(m, v9832)
	mBase = m.M
	v10058 = m.ExcPending
	if v10058 != 0 {
		goto L32
	} else {
		goto L2192
	}
L2192:
	;
	if v10057 != 0 {
		goto L7
	} else {
		goto L2193
	}
L2193:
	;
	v10065 = F_InstallXLogFileSegment(m, v39+int32(_a_F_StartupXLOG_282), v39+int32(_a_F_StartupXLOG_3), int32(0), int64(0), v9786)
	mBase = m.M
	v10066 = m.ExcPending
	if v10066 != 0 {
		goto L32
	} else {
		goto L2194
	}
L2194:
	;
	if v10065 != 0 {
		goto L2138
	} else {
		goto L2195
	}
L2195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10070 = m.ExcPending
	if v10070 != 0 {
		goto L32
	} else {
		goto L2196
	}
L2196:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_283), int32(0))
	mBase = m.M
	v10074 = m.ExcPending
	if v10074 != 0 {
		goto L32
	} else {
		goto L2197
	}
L2197:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3531), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v10079 = m.ExcPending
	if v10079 != 0 {
		goto L32
	} else {
		goto L2198
	}
L2198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2199:
	;
	v10082 = F_close(m, v10080)
	mBase = m.M
	if v10082 != 0 {
		goto L6
	} else {
		goto L2200
	}
L2200:
	;
	goto L2138
L2201:
	;
	F_XLogArchiveCleanup(m, v10128)
	mBase = m.M
	v10136 = m.ExcPending
	if v10136 != 0 {
		goto L32
	} else {
		goto L2202
	}
L2202:
	;
	v10137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9207)+68)))
	if v10137 == int32(1) {
		goto L2203
	} else {
		goto L2204
	}
L2203:
	;
	v10142 = F_durable_unlink(m, int32(_a_F_StartupXLOG_47), int32(22))
	mBase = m.M
	v10143 = m.ExcPending
	if v10143 != 0 {
		goto L32
	} else {
		goto L2206
	}
L2204:
	;
	goto L2205
L2205:
	;
	v10144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9207)+69)))
	if v10144 == int32(1) {
		goto L2207
	} else {
		goto L2208
	}
L2206:
	;
	goto L2205
L2207:
	;
	v10149 = F_durable_unlink(m, int32(_a_F_StartupXLOG_48), int32(22))
	mBase = m.M
	v10150 = m.ExcPending
	if v10150 != 0 {
		goto L32
	} else {
		goto L2210
	}
L2208:
	;
	goto L2209
L2209:
	;
	v10152 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[12]))
	v10153 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+64))
	v10154 = m.G0
	v10156 = v10154 - int32(_a_F_StartupXLOG_284)
	m.G0 = v10156
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+224)) = int32(42)
	v10161 = v10156 + int32(_a_F_StartupXLOG_285)
	v10166 = F_pg_snprintf(m, v10161, int32(1024), int32(_a_F_StartupXLOG_279), v10156+int32(224))
	mBase = m.M
	v10167 = m.ExcPending
	if v10167 != 0 {
		goto L32
	} else {
		goto L2211
	}
L2210:
	;
	goto L2209
L2211:
	;
	v10168 = F_unlink(m, v10161)
	mBase = m.M
	v10170 = F_OpenTransientFile(m, v10161, int32(194))
	mBase = m.M
	v10171 = m.ExcPending
	if v10171 != 0 {
		goto L32
	} else {
		goto L2218
	}
L2212:
	;
	v10633 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v10634 = m.ExcPending
	if v10634 != 0 {
		goto L32
	} else {
		goto L2311
	}
L2213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10615 = m.ExcPending
	if v10615 != 0 {
		goto L32
	} else {
		goto L2307
	}
L2214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+112)) = v10153
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+100)) = v10152
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+96)) = v10386
	*(*uint32)(unsafe.Add(mBase, uint32(v10156)+108)) = uint32(v9562)
	v10422 = int64(base.Ui64(v9562) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v10156)+104)) = uint32(v10422)
	v10425 = v10156 + int32(240)
	v10430 = F_pg_snprintf(m, v10425, int32(_a_F_StartupXLOG_58), int32(_a_F_StartupXLOG_286), v10156+int32(96))
	mBase = m.M
	v10431 = m.ExcPending
	if v10431 != 0 {
		goto L32
	} else {
		goto L2264
	}
L2215:
	;
	v10360 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10360 == int32(44) {
		goto L2257
	} else {
		goto L2258
	}
L2216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10343 = m.ExcPending
	if v10343 != 0 {
		goto L32
	} else {
		goto L2253
	}
L2217:
	;
	v10314 = int32(_a_F_StartupXLOG_2)
	v10315 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10317 = v10156 + int32(_a_F_StartupXLOG_285)
	v10318 = F_unlink(m, v10317)
	mBase = m.M
	if v10315 != 0 {
		goto L2246
	} else {
		goto L2247
	}
L2218:
	;
	if int32(0) <= v10170 {
		goto L2219
	} else {
		goto L2220
	}
L2219:
	;
	v10175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v10175 == int32(1) {
		goto L2223
	} else {
		goto L2224
	}
L2220:
	;
	goto L2221
L2221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10300 = m.ExcPending
	if v10300 != 0 {
		goto L32
	} else {
		goto L2242
	}
L2222:
	;
	v10207 = F_OpenTransientFile(m, v10156+int32(_a_F_StartupXLOG_287), int32(0))
	mBase = m.M
	v10208 = m.ExcPending
	if v10208 != 0 {
		goto L32
	} else {
		goto L2229
	}
L2223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+192)) = v10152
	v10180 = v10156 + int32(_a_F_StartupXLOG_288)
	v10185 = F_pg_snprintf(m, v10180, int32(64), int32(_a_F_StartupXLOG_289), v10156+int32(192))
	mBase = m.M
	v10186 = m.ExcPending
	if v10186 != 0 {
		goto L32
	} else {
		goto L2226
	}
L2224:
	;
	goto L2225
L2225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+208)) = v10152
	v10201 = F_pg_snprintf(m, v10156+int32(_a_F_StartupXLOG_287), int32(1024), int32(_a_F_StartupXLOG_290), v10156+int32(208))
	mBase = m.M
	v10202 = m.ExcPending
	if v10202 != 0 {
		goto L32
	} else {
		goto L2228
	}
L2226:
	;
	v10192 = F_RestoreArchivedFile(m, v10156+int32(_a_F_StartupXLOG_287), v10180, int32(_a_F_StartupXLOG_291), int64(0), int32(0))
	mBase = m.M
	v10193 = m.ExcPending
	if v10193 != 0 {
		goto L32
	} else {
		goto L2227
	}
L2227:
	;
	goto L2222
L2228:
	;
	goto L2222
L2229:
	;
	if v10207 < int32(0) {
		goto L2215
	} else {
		goto L2230
	}
L2230:
	;
	v10212 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10212
	v10214 = int32(_a_F_StartupXLOG_143)
	v10215 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10215))) = int32(167772219)
	v10221 = F_read(m, v10207, v10156+int32(240), int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v10223 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10223))) = v10212
	if v10221 < v10212 {
		goto L2213
	} else {
		goto L2231
	}
L2231:
	;
	v10231 = v10221
	goto L2232
L2232:
	;
	v10263 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10263 != 0 {
		goto L2213
	} else {
		goto L2234
	}
L2233:
	;
	v10294 = F_CloseTransientFile(m, v10207)
	mBase = m.M
	v10295 = m.ExcPending
	if v10295 != 0 {
		goto L32
	} else {
		goto L2240
	}
L2234:
	;
	if v10231 != 0 {
		goto L2235
	} else {
		goto L2236
	}
L2235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10268 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10268))) = int32(167772221)
	v10272 = v10156 + int32(240)
	v10273 = F_write(m, v10170, v10272, v10231)
	mBase = m.M
	if v10273 != v10231 {
		goto L2217
	} else {
		goto L2238
	}
L2236:
	;
	goto L2237
L2237:
	;
	goto L2233
L2238:
	;
	v10275 = int32(_a_F_StartupXLOG_143)
	v10276 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	v10277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10276))) = v10277
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10277
	v10283 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10283))) = int32(167772219)
	v10287 = F_read(m, v10207, v10272, int32(_a_F_StartupXLOG_58))
	mBase = m.M
	v10289 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10289))) = v10277
	if v10277 <= v10287 {
		v10231 = v10287
		goto L2232
	} else {
		goto L2239
	}
L2239:
	;
	goto L2213
L2240:
	;
	if v10294 != 0 {
		goto L2216
	} else {
		goto L2241
	}
L2241:
	;
	v10386 = int32(_a_F_StartupXLOG_292)
	goto L2214
L2242:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10302 = m.ExcPending
	if v10302 != 0 {
		goto L32
	} else {
		goto L2243
	}
L2243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156))) = v10156 + int32(_a_F_StartupXLOG_285)
	F_errmsg(m, int32(_a_F_StartupXLOG_293), v10156)
	mBase = m.M
	v10308 = m.ExcPending
	if v10308 != 0 {
		goto L32
	} else {
		goto L2244
	}
L2244:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(329), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10313 = m.ExcPending
	if v10313 != 0 {
		goto L32
	} else {
		goto L2245
	}
L2245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2246:
	;
	v10321 = v10315
	goto L2248
L2247:
	;
	v10321 = int32(51)
	goto L2248
L2248:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10321
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10326 = m.ExcPending
	if v10326 != 0 {
		goto L32
	} else {
		goto L2249
	}
L2249:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10328 = m.ExcPending
	if v10328 != 0 {
		goto L32
	} else {
		goto L2250
	}
L2250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+176)) = v10317
	F_errmsg(m, int32(_a_F_StartupXLOG_296), v10156+int32(176))
	mBase = m.M
	v10334 = m.ExcPending
	if v10334 != 0 {
		goto L32
	} else {
		goto L2251
	}
L2251:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(384), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10339 = m.ExcPending
	if v10339 != 0 {
		goto L32
	} else {
		goto L2252
	}
L2252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2253:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10345 = m.ExcPending
	if v10345 != 0 {
		goto L32
	} else {
		goto L2254
	}
L2254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+160)) = v10156 + int32(_a_F_StartupXLOG_287)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v10156+int32(160))
	mBase = m.M
	v10353 = m.ExcPending
	if v10353 != 0 {
		goto L32
	} else {
		goto L2255
	}
L2255:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(392), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10358 = m.ExcPending
	if v10358 != 0 {
		goto L32
	} else {
		goto L2256
	}
L2256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2257:
	;
	v10386 = int32(_a_F_StartupXLOG_297)
	goto L2214
L2258:
	;
	goto L2259
L2259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10367 = m.ExcPending
	if v10367 != 0 {
		goto L32
	} else {
		goto L2260
	}
L2260:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10369 = m.ExcPending
	if v10369 != 0 {
		goto L32
	} else {
		goto L2261
	}
L2261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+128)) = v10156 + int32(_a_F_StartupXLOG_287)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v10156+int32(128))
	mBase = m.M
	v10377 = m.ExcPending
	if v10377 != 0 {
		goto L32
	} else {
		goto L2262
	}
L2262:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(348), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10382 = m.ExcPending
	if v10382 != 0 {
		goto L32
	} else {
		goto L2263
	}
L2263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2264:
	;
	v10432 = F_strlen(m, v10425)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = int32(0)
	v10437 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10437))) = int32(167772221)
	v10440 = F_write(m, v10170, v10425, v10432)
	mBase = m.M
	if v10440 == v10432 {
		goto L2266
	} else {
		goto L2267
	}
L2265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10562 = m.ExcPending
	if v10562 != 0 {
		goto L32
	} else {
		goto L2303
	}
L2266:
	;
	v10442 = int32(_a_F_StartupXLOG_143)
	v10443 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	v10444 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10443))) = v10444
	v10447 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10447))) = int32(167772220)
	v10452 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[11])))
	if v10452 != int32(1) {
		v10466 = v10444
		goto L2271
	} else {
		goto L2272
	}
L2267:
	;
	goto L2268
L2268:
	;
	v10533 = int32(_a_F_StartupXLOG_2)
	v10534 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	v10536 = v10156 + int32(_a_F_StartupXLOG_285)
	v10537 = F_unlink(m, v10536)
	mBase = m.M
	if v10534 != 0 {
		goto L2296
	} else {
		goto L2297
	}
L2269:
	;
	v10495 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v10495))) = int32(0)
	v10498 = F_CloseTransientFile(m, v10170)
	mBase = m.M
	v10499 = m.ExcPending
	if v10499 != 0 {
		goto L32
	} else {
		goto L2287
	}
L2270:
	;
	if v10466 == int32(0) {
		goto L2269
	} else {
		goto L2277
	}
L2271:
	;
	goto L2270
L2272:
	;
	goto L2273
L2273:
	;
	v10457 = F_fsync(m, v10170)
	mBase = m.M
	if v10457 != int32(-1) {
		v10466 = v10457
		goto L2271
	} else {
		goto L2275
	}
L2274:
	;
	v10466 = int32(-1)
	goto L2271
L2275:
	;
	v10461 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3]))
	if v10461 == int32(27) {
		goto L2273
	} else {
		goto L2276
	}
L2276:
	;
	goto L2274
L2277:
	;
	v10472 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[127])))
	if v10472 != 0 {
		goto L2279
	} else {
		goto L2280
	}
L2278:
	;
	v10475 = F_errstart(m, v10473, int32(0))
	mBase = m.M
	v10476 = m.ExcPending
	if v10476 != 0 {
		goto L32
	} else {
		goto L2282
	}
L2279:
	;
	v10473 = int32(21)
	goto L2281
L2280:
	;
	v10473 = int32(23)
	goto L2281
L2281:
	;
	goto L2278
L2282:
	;
	if v10475 == int32(0) {
		goto L2269
	} else {
		goto L2283
	}
L2283:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10480 = m.ExcPending
	if v10480 != 0 {
		goto L32
	} else {
		goto L2284
	}
L2284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+64)) = v10156 + int32(_a_F_StartupXLOG_285)
	F_errmsg(m, int32(_a_F_StartupXLOG_149), v10156-int32(-64))
	mBase = m.M
	v10488 = m.ExcPending
	if v10488 != 0 {
		goto L32
	} else {
		goto L2285
	}
L2285:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(432), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10493 = m.ExcPending
	if v10493 != 0 {
		goto L32
	} else {
		goto L2286
	}
L2286:
	;
	goto L2269
L2287:
	;
	if v10498 != 0 {
		goto L2265
	} else {
		goto L2288
	}
L2288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+32)) = v9786
	v10502 = v10156 + int32(_a_F_StartupXLOG_287)
	v10507 = F_pg_snprintf(m, v10502, int32(1024), int32(_a_F_StartupXLOG_290), v10156+int32(32))
	mBase = m.M
	v10508 = m.ExcPending
	if v10508 != 0 {
		goto L32
	} else {
		goto L2289
	}
L2289:
	;
	v10512 = F_durable_rename(m, v10156+int32(_a_F_StartupXLOG_285), v10502, int32(21))
	mBase = m.M
	v10513 = m.ExcPending
	if v10513 != 0 {
		goto L32
	} else {
		goto L2290
	}
L2290:
	;
	v10515 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[128]))
	if int32(0) < v10515 {
		goto L2291
	} else {
		goto L2292
	}
L2291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+16)) = v9786
	v10520 = v10156 + int32(_a_F_StartupXLOG_288)
	v10525 = F_pg_snprintf(m, v10520, int32(64), int32(_a_F_StartupXLOG_289), v10156+int32(16))
	mBase = m.M
	v10526 = m.ExcPending
	if v10526 != 0 {
		goto L32
	} else {
		goto L2294
	}
L2292:
	;
	goto L2293
L2293:
	;
	m.G0 = v10156 + int32(_a_F_StartupXLOG_284)
	goto L2212
L2294:
	;
	F_XLogArchiveNotify(m, v10520)
	mBase = m.M
	v10528 = m.ExcPending
	if v10528 != 0 {
		goto L32
	} else {
		goto L2295
	}
L2295:
	;
	goto L2293
L2296:
	;
	v10540 = v10534
	goto L2298
L2297:
	;
	v10540 = int32(51)
	goto L2298
L2298:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v10540
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10545 = m.ExcPending
	if v10545 != 0 {
		goto L32
	} else {
		goto L2299
	}
L2299:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10547 = m.ExcPending
	if v10547 != 0 {
		goto L32
	} else {
		goto L2300
	}
L2300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+80)) = v10536
	F_errmsg(m, int32(_a_F_StartupXLOG_296), v10156+int32(80))
	mBase = m.M
	v10553 = m.ExcPending
	if v10553 != 0 {
		goto L32
	} else {
		goto L2301
	}
L2301:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(424), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10558 = m.ExcPending
	if v10558 != 0 {
		goto L32
	} else {
		goto L2302
	}
L2302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2303:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10564 = m.ExcPending
	if v10564 != 0 {
		goto L32
	} else {
		goto L2304
	}
L2304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+48)) = v10156 + int32(_a_F_StartupXLOG_285)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v10156+int32(48))
	mBase = m.M
	v10572 = m.ExcPending
	if v10572 != 0 {
		goto L32
	} else {
		goto L2305
	}
L2305:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(438), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10577 = m.ExcPending
	if v10577 != 0 {
		goto L32
	} else {
		goto L2306
	}
L2306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2307:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v10617 = m.ExcPending
	if v10617 != 0 {
		goto L32
	} else {
		goto L2308
	}
L2308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+144)) = v10156 + int32(_a_F_StartupXLOG_287)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v10156+int32(144))
	mBase = m.M
	v10625 = m.ExcPending
	if v10625 != 0 {
		goto L32
	} else {
		goto L2309
	}
L2309:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_294), int32(362), int32(_a_F_StartupXLOG_295))
	mBase = m.M
	v10630 = m.ExcPending
	if v10630 != 0 {
		goto L32
	} else {
		goto L2310
	}
L2310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2311:
	;
	if v10633 == int32(0) {
		v10651 = v9786
		goto L2126
	} else {
		goto L2312
	}
L2312:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_298), int32(0))
	mBase = m.M
	v10640 = m.ExcPending
	if v10640 != 0 {
		goto L32
	} else {
		goto L2313
	}
L2313:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_299), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10645 = m.ExcPending
	if v10645 != 0 {
		goto L32
	} else {
		goto L2314
	}
L2314:
	;
	v10651 = v9786
	goto L2126
L2315:
	;
	v10686 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v10686+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_300), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v10693 = m.ExcPending
	if v10693 != 0 {
		goto L32
	} else {
		goto L2318
	}
L2316:
	;
	goto L2317
L2317:
	;
	v10695 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v10695)+308)) = v10651
	v10697 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10695)+440)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10695)+312)) = v10697
	if v9695 == int64(0) {
		goto L2319
	} else {
		goto L2320
	}
L2318:
	;
	goto L2317
L2319:
	;
	v10703 = v9562
	goto L2321
L2320:
	;
	v10703 = v9695
	goto L2321
L2321:
	;
	v10704 = *(*int64)(unsafe.Add(mBase, uint32(v9207)))
	v10707 = base.I32_wrap_i64(v10704) & int32(_a_F_StartupXLOG_280)
	v10709 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	v10710 = base.I64_extend_i32_s(v10709)
	v10711 = base.I64_div_u_s(v10704, v10710)
	v10714 = base.I64_extend_i32_s(v10709 - int32(1))
	v10715 = v10704 & v10714
	if v10715&int64(35184372080640) == int64(0) {
		goto L2323
	} else {
		goto L2324
	}
L2322:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10695)+16)) = v10753
	v10755 = base.I64_div_u_s(v10703, v10710)
	v10758 = base.I32_wrap_i64(v10703) & int32(_a_F_StartupXLOG_280)
	v10759 = v10703 & v10714
	if v10759&int64(35184372080640) == int64(0) {
		goto L2329
	} else {
		goto L2330
	}
L2323:
	;
	v10721 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129]))
	v10723 = v10711 * base.I64_extend_i32_s(v10721)
	if v10707 == int32(0) {
		v10751 = v10721
		v10753 = v10723
		goto L2322
	} else {
		goto L2326
	}
L2324:
	;
	goto L2325
L2325:
	;
	v10731 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129]))
	v10744 = v10711*base.I64_extend_i32_s(v10731) + (int64(base.Ui64(v10715)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v10707 == int32(0) {
		v10751 = v10731
		v10753 = v10744
		goto L2322
	} else {
		goto L2327
	}
L2326:
	;
	v10751 = v10721
	v10753 = v10723 + base.I64_extend_i32_u(v10707-int32(40))
	goto L2322
L2327:
	;
	v10751 = v10731
	v10753 = v10744 + base.I64_extend_i32_u(v10707-int32(24))
	goto L2322
L2328:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10695)+8)) = v10792
	if v10703&int64(8191) != int64(0) {
		goto L2334
	} else {
		goto L2335
	}
L2329:
	;
	v10765 = v10755 * base.I64_extend_i32_s(v10751)
	if v10758 == int32(0) {
		v10792 = v10765
		goto L2328
	} else {
		goto L2332
	}
L2330:
	;
	goto L2331
L2331:
	;
	v10784 = v10755*base.I64_extend_i32_s(v10751) + (int64(base.Ui64(v10759)>>(uint(int64(13))%64))*int64(8168)+int64(4294959128))&int64(4294967288) + int64(8152)
	if v10758 == int32(0) {
		v10792 = v10784
		goto L2328
	} else {
		goto L2333
	}
L2332:
	;
	v10792 = v10765 + base.I64_extend_i32_u(v10758-int32(40))
	goto L2328
L2333:
	;
	v10792 = v10784 + base.I64_extend_i32_u(v10758-int32(24))
	goto L2328
L2334:
	;
	v10798 = *(*int32)(unsafe.Add(mBase, uint32(v10695)+296))
	v10801 = *(*int32)(unsafe.Add(mBase, uint32(v10695)+304))
	v10805 = base.I64_rem_u_s(int64(base.Ui64(v10703)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v10801+int32(1)))
	v10806 = base.I32_wrap_i64(v10805)
	v10809 = v10798 + v10806<<(uint(int32(13))%32)
	v10810 = *(*int64)(unsafe.Add(mBase, uint32(v9207)+32))
	v10812 = base.I32_wrap_i64(v10703 - v10810)
	if v10812 != 0 {
		goto L2337
	} else {
		goto L2338
	}
L2335:
	;
	v10833 = v10695
	v10837 = v10703
	goto L2336
L2336:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10833)+288)) = v10837
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[130])) = v10703
	*(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[131])) = v10703
	*(*int64)(unsafe.Add(mBase, uint32(v10833)+264)) = v10703
	v10844 = int32(_a_F_StartupXLOG_301)
	v10845 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10845)+272)) = v10703
	v10848 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10848)+280)) = v10703
	v10851 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10851)+192)) = v10703
	*(*int64)(unsafe.Add(mBase, uint32(v10851)+184)) = v10703
	v10854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10851)+320)))
	if v10854 != int32(1) {
		goto L2343
	} else {
		goto L2344
	}
L2337:
	;
	v10813 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+40))
	base.MemoryCopy(m, v10809, v10813, v10812)
	goto L2339
L2338:
	;
	goto L2339
L2339:
	;
	v10816 = int32(_a_F_StartupXLOG_58) - v10812
	if v10816 != 0 {
		goto L2340
	} else {
		goto L2341
	}
L2340:
	;
	base.MemoryFill(m, v10812+v10809, int32(0), v10816)
	goto L2342
L2341:
	;
	goto L2342
L2342:
	;
	v10820 = *(*int32)(unsafe.Add(mBase, uint32(v10695)+300))
	v10824 = *(*int64)(unsafe.Add(mBase, uint32(v9207)+32))
	v10825 = int64(-8192)
	*(*int64)(unsafe.Add(mBase, uint32(v10820+v10806<<(uint(int32(3))%32)))) = v10824 - v10825
	v10828 = *(*int64)(unsafe.Add(mBase, uint32(v9207)+32))
	v10832 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v10833 = v10832
	v10837 = v10828 - v10825
	goto L2336
L2343:
	;
	v10896 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[42])) = uint8(v10896)
	v10898 = F_time(m)
	mBase = m.M
	v10900 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int64)(unsafe.Add(mBase, uint32(v10900)+256)) = v10703
	*(*int64)(unsafe.Add(mBase, uint32(v10900)+248)) = v10898
	v10904 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v10908 = F_LWLockAcquire(m, v10904+int32(512), v10896)
	mBase = m.M
	v10909 = m.ExcPending
	if v10909 != 0 {
		goto L32
	} else {
		goto L2351
	}
L2344:
	;
	v10858 = v10703 - int64(1)
	v10860 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if base.Ui64(v10858&base.I64_extend_i32_s(v10860-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v10860), float64(0.75))))) {
		goto L2343
	} else {
		goto L2345
	}
L2345:
	;
	v10872 = base.I64_div_u_s(v10858, base.I64_extend_i32_s(v10860))
	v10879 = F_XLogFileInitInternal(m, v10872+int64(1), v10651, v39+int32(_a_F_StartupXLOG_4), v39+int32(_a_F_StartupXLOG_1))
	mBase = m.M
	v10880 = m.ExcPending
	if v10880 != 0 {
		goto L32
	} else {
		goto L2346
	}
L2346:
	;
	if int32(0) <= v10879 {
		goto L2347
	} else {
		goto L2348
	}
L2347:
	;
	v10883 = F_close(m, v10879)
	mBase = m.M
	goto L2349
L2348:
	;
	goto L2349
L2349:
	;
	v10884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[96]))))
	if v10884 != int32(1) {
		goto L2343
	} else {
		goto L2350
	}
L2350:
	;
	v10887 = int32(_a_F_StartupXLOG_302)
	v10889 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[132]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[132])) = v10889 + int32(1)
	goto L2343
L2351:
	;
	v10911 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v10912 = *(*int64)(unsafe.Add(mBase, uint32(v10911)+8))
	v10914 = v10912 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10911)+48)) = v10914
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(base.I32_wrap_i64(v10914)))|base.B2i32(base.Ui64(v10914) < base.Ui64(int64(3))) == int32(0) {
		goto L2352
	} else {
		goto L2353
	}
L2352:
	;
	v10951 = v10914
	goto L2355
L2353:
	;
	goto L2354
L2354:
	;
	v10999 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v10999+int32(512))
	mBase = m.M
	v11003 = m.ExcPending
	if v11003 != 0 {
		goto L32
	} else {
		goto L2358
	}
L2355:
	;
	v10959 = v10951 - int64(1)
	if base.Ui32(base.I32_wrap_i64(v10959)) < base.Ui32(int32(3)) {
		v10951 = v10959
		goto L2355
	} else {
		goto L2357
	}
L2356:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10911)+48)) = v10959
	goto L2354
L2357:
	;
	goto L2356
L2358:
	;
	v11005 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v11005 == int32(0) {
		goto L2359
	} else {
		goto L2360
	}
L2359:
	;
	F_StartupSUBTRANS(m, v9757)
	mBase = m.M
	v11009 = m.ExcPending
	if v11009 != 0 {
		goto L32
	} else {
		goto L2362
	}
L2360:
	;
	goto L2361
L2361:
	;
	v11011 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11012 = *(*int32)(unsafe.Add(mBase, uint32(v11011)+28))
	v11014 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	v11015 = *(*int64)(unsafe.Add(mBase, uint32(v11014)+8))
	v11019 = int64(base.Ui64(v11015)>>(uint(int64(15))%64)) & int64(131071)
	v11022 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[133])))
	v11023 = base.I32_rem_u_s(base.I32_wrap_i64(v11019), v11022)
	v11026 = v11012 + v11023<<(uint(int32(7))%32)
	v11028 = F_LWLockAcquire(m, v11026, int32(0))
	mBase = m.M
	v11029 = m.ExcPending
	if v11029 != 0 {
		goto L32
	} else {
		goto L2363
	}
L2362:
	;
	goto L2361
L2363:
	;
	v11030 = base.I32_wrap_i64(v11015)
	v11032 = v11030 & int32(_a_F_StartupXLOG_303)
	if v11032 != 0 {
		goto L2364
	} else {
		goto L2365
	}
L2364:
	;
	v11035 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_304), v11019, int32(0), v11030)
	mBase = m.M
	v11036 = m.ExcPending
	if v11036 != 0 {
		goto L32
	} else {
		goto L2367
	}
L2365:
	;
	goto L2366
L2366:
	;
	F_LWLockRelease(m, v11026)
	mBase = m.M
	v11113 = m.ExcPending
	if v11113 != 0 {
		goto L32
	} else {
		goto L2378
	}
L2367:
	;
	v11038 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11039 = *(*int32)(unsafe.Add(mBase, uint32(v11038)+4))
	v11040 = int32(2)
	v11043 = *(*int32)(unsafe.Add(mBase, uint32(v11039+v11035<<(uint(v11040)%32))))
	v11045 = int32(base.Ui32(v11032) >> (uint(v11040) % 32))
	v11046 = v11043 + v11045
	v11047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11046))))
	v11048 = int32(-1)
	v11049 = int32(1)
	v11056 = v11047 & (v11048<<(uint(v11030<<(uint(v11049)%32)&int32(6))%32) ^ v11048)
	*(*uint8)(unsafe.Add(mBase, uint32(v11046))) = uint8(v11056)
	v11059 = v11046 + v11049
	v11060 = int32(3)
	v11063 = v11045 ^ int32(_a_F_StartupXLOG_280)
	if v11059&v11060|v11063&v11060|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v11063)) == int32(0) {
		goto L2369
	} else {
		goto L2370
	}
L2368:
	;
	v11101 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[64]))
	v11102 = *(*int32)(unsafe.Add(mBase, uint32(v11101)+12))
	v11104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11102+v11035))) = uint8(v11104)
	goto L2366
L2369:
	;
	if v11045 == int32(_a_F_StartupXLOG_280) {
		goto L2368
	} else {
		goto L2372
	}
L2370:
	;
	goto L2371
L2371:
	;
	if v11063 == int32(0) {
		goto L2368
	} else {
		goto L2377
	}
L2372:
	;
	v11077 = v11063 + v11043 + v11045 + int32(1)
	v11079 = v11046 + int32(5)
	if base.Ui32(v11079) < base.Ui32(v11077) {
		goto L2373
	} else {
		goto L2374
	}
L2373:
	;
	v11081 = v11077
	goto L2375
L2374:
	;
	v11081 = v11079
	goto L2375
L2375:
	;
	v11088 = (v11081-v11046-int32(2))&int32(-4) + int32(4)
	if v11088 == int32(0) {
		goto L2368
	} else {
		goto L2376
	}
L2376:
	;
	base.MemoryFill(m, v11059, int32(0), v11088)
	goto L2368
L2377:
	;
	base.MemoryFill(m, v11059, int32(0), v11063)
	goto L2368
L2378:
	;
	v11115 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11119 = F_LWLockAcquire(m, v11115+int32(1664), int32(1))
	mBase = m.M
	v11120 = m.ExcPending
	if v11120 != 0 {
		goto L32
	} else {
		goto L2379
	}
L2379:
	;
	v11122 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v11123 = *(*int32)(unsafe.Add(mBase, uint32(v11122)+16))
	v11124 = *(*int32)(unsafe.Add(mBase, uint32(v11122)+12))
	v11125 = *(*int32)(unsafe.Add(mBase, uint32(v11122)+4))
	v11126 = *(*int32)(unsafe.Add(mBase, uint32(v11122)))
	v11128 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11128+int32(1664))
	mBase = m.M
	v11132 = m.ExcPending
	if v11132 != 0 {
		goto L32
	} else {
		goto L2380
	}
L2380:
	;
	v11133 = int32(_a_F_StartupXLOG_305)
	v11134 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11136 = int32(base.Ui32(v11126) >> (uint(int32(11)) % 32))
	v11137 = base.I64_extend_i32_u(v11136)
	*(*int64)(unsafe.Add(mBase, uint32(v11134)+48)) = v11137
	v11140 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11141 = *(*int32)(unsafe.Add(mBase, uint32(v11140)+28))
	v11143 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[134])))
	v11144 = base.I32_rem_u_s(v11136, v11143)
	v11147 = v11141 + v11144<<(uint(int32(7))%32)
	v11149 = F_LWLockAcquire(m, v11147, int32(0))
	mBase = m.M
	v11150 = m.ExcPending
	if v11150 != 0 {
		goto L32
	} else {
		goto L2381
	}
L2381:
	;
	v11152 = v11126 & int32(2047)
	if v11152 == int32(0) {
		goto L2383
	} else {
		goto L2384
	}
L2382:
	;
	v11164 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11165 = *(*int32)(unsafe.Add(mBase, uint32(v11164)+4))
	v11166 = int32(2)
	v11169 = *(*int32)(unsafe.Add(mBase, uint32(v11165+v11162<<(uint(v11166)%32))))
	v11171 = v11152 << (uint(v11166) % 32)
	v11172 = v11169 + v11171
	*(*int32)(unsafe.Add(mBase, uint32(v11172))) = v11125
	if base.Ui32(int32(2045)) < base.Ui32(v11152-int32(1)) {
		goto L2388
	} else {
		goto L2389
	}
L2383:
	;
	v11156 = F_SimpleLruZeroPage(m, int32(_a_F_StartupXLOG_305), v11137)
	mBase = m.M
	v11157 = m.ExcPending
	if v11157 != 0 {
		goto L32
	} else {
		goto L2386
	}
L2384:
	;
	goto L2385
L2385:
	;
	v11160 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_305), v11137, int32(1), v11126)
	mBase = m.M
	v11161 = m.ExcPending
	if v11161 != 0 {
		goto L32
	} else {
		goto L2387
	}
L2386:
	;
	v11162 = v11156
	goto L2382
L2387:
	;
	v11162 = v11160
	goto L2382
L2388:
	;
	v11215 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[66]))
	v11216 = *(*int32)(unsafe.Add(mBase, uint32(v11215)+12))
	v11218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11216+v11162))) = uint8(v11218)
	F_LWLockRelease(m, v11147)
	mBase = m.M
	v11221 = m.ExcPending
	if v11221 != 0 {
		goto L32
	} else {
		goto L2398
	}
L2389:
	;
	v11179 = v11172 + int32(4)
	if v11179&int32(3)|base.B2i32(base.Ui32(v11152) < base.Ui32(int32(1791))) == int32(0) {
		goto L2390
	} else {
		goto L2391
	}
L2390:
	;
	v11188 = v11172 + int32(8)
	v11190 = v11169 - int32(-8192)
	if base.Ui32(v11190) < base.Ui32(v11188) {
		goto L2393
	} else {
		goto L2394
	}
L2391:
	;
	goto L2392
L2392:
	;
	v11205 = int32(_a_F_StartupXLOG_306) - v11171
	if v11205 == int32(0) {
		goto L2388
	} else {
		goto L2397
	}
L2393:
	;
	v11192 = v11188
	goto L2395
L2394:
	;
	v11192 = v11190
	goto L2395
L2395:
	;
	v11199 = (v11192-v11172-int32(5))&int32(-4) + int32(4)
	if v11199 == int32(0) {
		goto L2388
	} else {
		goto L2396
	}
L2396:
	;
	base.MemoryFill(m, v11179, int32(0), v11199)
	goto L2388
L2397:
	;
	base.MemoryFill(m, v11179, int32(0), v11205)
	goto L2388
L2398:
	;
	v11223 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11225 = base.I32_div_u_s(v11125, int32(1636))
	v11226 = base.I64_extend_i32_u(v11225)
	*(*int64)(unsafe.Add(mBase, uint32(v11223)+48)) = v11226
	v11229 = int32(base.Ui32(v11125) >> (uint(int32(2)) % 32))
	v11231 = base.I32_rem_u_s(v11229, int32(409))
	if v11231 != 0 {
		goto L2399
	} else {
		goto L2400
	}
L2399:
	;
	v11233 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11234 = *(*int32)(unsafe.Add(mBase, uint32(v11233)+28))
	v11236 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_StartupXLOG[135])))
	v11237 = base.I32_rem_u_s(v11225, v11236)
	v11240 = v11234 + v11237<<(uint(int32(7))%32)
	v11242 = F_LWLockAcquire(m, v11240, int32(0))
	mBase = m.M
	v11243 = m.ExcPending
	if v11243 != 0 {
		goto L32
	} else {
		goto L2402
	}
L2400:
	;
	goto L2401
L2401:
	;
	v11309 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11313 = F_LWLockAcquire(m, v11309+int32(1664), int32(0))
	mBase = m.M
	v11314 = m.ExcPending
	if v11314 != 0 {
		goto L32
	} else {
		goto L2411
	}
L2402:
	;
	v11246 = F_SimpleLruReadPage(m, int32(_a_F_StartupXLOG_307), v11226, int32(1), v11125)
	mBase = m.M
	v11247 = m.ExcPending
	if v11247 != 0 {
		goto L32
	} else {
		goto L2403
	}
L2403:
	;
	v11248 = int32(2)
	v11251 = v11125 << (uint(v11248) % 32) & int32(12)
	v11256 = v11251 + v11231*int32(20) + int32(4)
	v11258 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11259 = *(*int32)(unsafe.Add(mBase, uint32(v11258)+4))
	v11263 = *(*int32)(unsafe.Add(mBase, uint32(v11259+v11246<<(uint(v11248)%32))))
	v11264 = v11256 + v11263
	if v11264&int32(3)|base.B2i32(base.Ui32(v11256) < base.Ui32(int32(_a_F_StartupXLOG_308))) == int32(0) {
		goto L2405
	} else {
		goto L2406
	}
L2404:
	;
	v11296 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[67]))
	v11297 = *(*int32)(unsafe.Add(mBase, uint32(v11296)+12))
	v11299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11297+v11246))) = uint8(v11299)
	F_LWLockRelease(m, v11240)
	mBase = m.M
	v11302 = m.ExcPending
	if v11302 != 0 {
		goto L32
	} else {
		goto L2410
	}
L2405:
	;
	v11283 = (v11225*int32(_a_F_StartupXLOG_309)-v11251+v11229*int32(-20)+int32(_a_F_StartupXLOG_310))&int32(-4) + int32(4)
	if v11283 == int32(0) {
		goto L2404
	} else {
		goto L2408
	}
L2406:
	;
	goto L2407
L2407:
	;
	v11289 = int32(_a_F_StartupXLOG_58) - v11256
	if v11289 == int32(0) {
		goto L2404
	} else {
		goto L2409
	}
L2408:
	;
	base.MemoryFill(m, v11264, int32(0), v11283)
	goto L2404
L2409:
	;
	base.MemoryFill(m, v11264, int32(0), v11289)
	goto L2404
L2410:
	;
	goto L2401
L2411:
	;
	v11316 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[65]))
	v11317 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11316)+8)) = uint8(v11317)
	v11320 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11320+int32(1664))
	mBase = m.M
	v11324 = m.ExcPending
	if v11324 != 0 {
		goto L32
	} else {
		goto L2412
	}
L2412:
	;
	F_SetMultiXactIdLimit(m, v11124, v11123, int32(1))
	mBase = m.M
	v11327 = m.ExcPending
	if v11327 != 0 {
		goto L32
	} else {
		goto L2413
	}
L2413:
	;
	v11328 = int32(0)
	v11329 = m.G0
	v11331 = v11329 - int32(16)
	m.G0 = v11331
	v11334 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11338 = F_LWLockAcquire(m, v11334+int32(2304), v11328)
	mBase = m.M
	v11339 = m.ExcPending
	if v11339 != 0 {
		goto L32
	} else {
		goto L2414
	}
L2414:
	;
	v11341 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[136]))
	v11342 = *(*int32)(unsafe.Add(mBase, uint32(v11341)+4))
	if int32(0) < v11342 {
		goto L2415
	} else {
		goto L2416
	}
L2415:
	;
	v11351 = v11341
	v11357 = v11328
	goto L2418
L2416:
	;
	goto L2417
L2417:
	;
	v11656 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11656+int32(2304))
	mBase = m.M
	v11660 = m.ExcPending
	if v11660 != 0 {
		goto L32
	} else {
		goto L2459
	}
L2418:
	;
	v11382 = *(*int32)(unsafe.Add(mBase, uint32(v11351+v11357<<(uint(int32(2))%32))+8))
	v11383 = *(*int32)(unsafe.Add(mBase, uint32(v11382)+32))
	v11384 = *(*int64)(unsafe.Add(mBase, uint32(v11382)+16))
	v11385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11382)+45)))
	v11388 = F_ProcessTwoPhaseBuffer(m, v11383, v11384, v11385, int32(1), int32(0))
	mBase = m.M
	v11389 = m.ExcPending
	if v11389 != 0 {
		goto L32
	} else {
		goto L2420
	}
L2419:
	;
	goto L2417
L2420:
	;
	if v11388 != 0 {
		goto L2421
	} else {
		goto L2422
	}
L2421:
	;
	v11392 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11393 = m.ExcPending
	if v11393 != 0 {
		goto L32
	} else {
		goto L2424
	}
L2422:
	;
	goto L2423
L2423:
	;
	v11616 = v11357 + int32(1)
	v11618 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[136]))
	v11619 = *(*int32)(unsafe.Add(mBase, uint32(v11618)+4))
	if v11616 < v11619 {
		v11351 = v11618
		v11357 = v11616
		goto L2418
	} else {
		goto L2458
	}
L2424:
	;
	if v11392 != 0 {
		goto L2425
	} else {
		goto L2426
	}
L2425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11331))) = v11383
	F_errmsg(m, int32(_a_F_StartupXLOG_311), v11331)
	mBase = m.M
	v11397 = m.ExcPending
	if v11397 != 0 {
		goto L32
	} else {
		goto L2428
	}
L2426:
	;
	goto L2427
L2427:
	;
	v11403 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+48))
	v11404 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+44))
	v11405 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+40))
	v11406 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+36))
	v11407 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+32))
	v11408 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+28))
	v11409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11388)+54)))
	v11411 = v11388 + int32(72)
	v11412 = *(*int64)(unsafe.Add(mBase, uint32(v11388)+16))
	v11413 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+24))
	v11414 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+12))
	F_MarkAsPreparingGuts(m, v11382, v11383, v11411, v11412, v11413, v11414)
	mBase = m.M
	v11416 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11382)+46)) = uint8(v11416)
	v11418 = int32(7)
	v11422 = v11411 + (v11409+v11418)&int32(_a_F_StartupXLOG_312)
	v11427 = int32(-8)
	v11430 = int32(12)
	v11444 = int32(4)
	v11454 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[137]))
	v11455 = *(*int32)(unsafe.Add(mBase, uint32(v11454)))
	v11456 = *(*int32)(unsafe.Add(mBase, uint32(v11382)+4))
	v11459 = v11455 + v11456*int32(640)
	v11460 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+28))
	if int32(65) <= v11460 {
		goto L2432
	} else {
		goto L2433
	}
L2428:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_313), int32(2106), int32(_a_F_StartupXLOG_314))
	mBase = m.M
	v11402 = m.ExcPending
	if v11402 != 0 {
		goto L32
	} else {
		goto L2429
	}
L2429:
	;
	goto L2427
L2430:
	;
	v11481 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11382)+44)) = uint8(v11481)
	v11483 = *(*int32)(unsafe.Add(mBase, uint32(v11479)))
	F_ProcArrayAdd(m, v11483+v11480*int32(640))
	mBase = m.M
	v11488 = m.ExcPending
	if v11488 != 0 {
		goto L32
	} else {
		goto L2439
	}
L2431:
	;
	v11470 = v11468 << (uint(int32(2)) % 32)
	if v11470 != 0 {
		goto L2436
	} else {
		goto L2437
	}
L2432:
	;
	v11463 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11459)+277)) = uint8(v11463)
	v11468 = int32(64)
	goto L2431
L2433:
	;
	goto L2434
L2434:
	;
	if v11460 <= int32(0) {
		v11479 = v11454
		v11480 = v11456
		goto L2430
	} else {
		goto L2435
	}
L2435:
	;
	v11468 = v11460
	goto L2431
L2436:
	;
	base.MemoryCopy(m, v11459+int32(280), v11422, v11470)
	goto L2438
L2437:
	;
	goto L2438
L2438:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11459)+276)) = uint8(v11468)
	v11475 = *(*int32)(unsafe.Add(mBase, uint32(v11382)+4))
	v11477 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[137]))
	v11479 = v11477
	v11480 = v11475
	goto L2430
L2439:
	;
	v11490 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11490+int32(2304))
	mBase = m.M
	v11494 = m.ExcPending
	if v11494 != 0 {
		goto L32
	} else {
		goto L2440
	}
L2440:
	;
	v11501 = v11422 + (v11408<<(uint(int32(2))%32)+v11418)&v11427 + (v11407*v11430+v11418)&v11427 + (v11406*v11430+v11418)&v11427 + v11405<<(uint(v11444)%32) + v11404<<(uint(v11444)%32) + v11403<<(uint(v11444)%32)
	goto L2441
L2441:
	;
	v11529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11501)+4)))
	if v11529 != 0 {
		goto L2443
	} else {
		goto L2444
	}
L2442:
	;
	v11546 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if base.Ui32(int32(2)) <= base.Ui32(v11546) {
		goto L2450
	} else {
		goto L2451
	}
L2443:
	;
	v11531 = v11501 + int32(8)
	v11534 = *(*int32)(unsafe.Add(mBase, uint32(v11529<<(uint(int32(2))%32))+uint32(_c_F_StartupXLOG[138])))
	if v11534 != 0 {
		goto L2446
	} else {
		goto L2447
	}
L2444:
	;
	goto L2445
L2445:
	;
	goto L2442
L2446:
	;
	v11535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11501)+6)))
	v11536 = *(*int32)(unsafe.Add(mBase, uint32(v11501)))
	m.T0[v11534].(func(*base.Module, int32, int32, int32, int32))(m, v11383, v11535, v11531, v11536)
	mBase = m.M
	v11538 = m.ExcPending
	if v11538 != 0 {
		goto L32
	} else {
		goto L2449
	}
L2447:
	;
	goto L2448
L2448:
	;
	v11539 = *(*int32)(unsafe.Add(mBase, uint32(v11501)))
	v11501 = v11531 + (v11539+int32(7))&int32(-8)
	goto L2441
L2449:
	;
	goto L2448
L2450:
	;
	v11549 = *(*int32)(unsafe.Add(mBase, uint32(v11388)+28))
	F_StandbyReleaseLockTree(m, v11383, v11549, v11422)
	mBase = m.M
	v11551 = m.ExcPending
	if v11551 != 0 {
		goto L32
	} else {
		goto L2453
	}
L2451:
	;
	goto L2452
L2452:
	;
	v11553 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11557 = F_LWLockAcquire(m, v11553+int32(2304), int32(0))
	mBase = m.M
	v11558 = m.ExcPending
	if v11558 != 0 {
		goto L32
	} else {
		goto L2454
	}
L2453:
	;
	goto L2452
L2454:
	;
	v11560 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[139]))
	*(*int32)(unsafe.Add(mBase, uint32(v11560)+40)) = int32(-1)
	v11564 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v11564+int32(2304))
	mBase = m.M
	v11568 = m.ExcPending
	if v11568 != 0 {
		goto L32
	} else {
		goto L2455
	}
L2455:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[139])) = int32(0)
	F_pfree(m, v11388)
	mBase = m.M
	v11573 = m.ExcPending
	if v11573 != 0 {
		goto L32
	} else {
		goto L2456
	}
L2456:
	;
	v11575 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v11579 = F_LWLockAcquire(m, v11575+int32(2304), int32(0))
	mBase = m.M
	v11580 = m.ExcPending
	if v11580 != 0 {
		goto L32
	} else {
		goto L2457
	}
L2457:
	;
	goto L2423
L2458:
	;
	goto L2419
L2459:
	;
	m.G0 = v11331 + int32(16)
	v11664 = m.G0
	v11666 = v11664 - int32(1024)
	m.G0 = v11666
	v11669 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v11674 = *(*int32)(unsafe.Add(mBase, uint32(v11669)))
	v11675 = *(*int32)(unsafe.Add(mBase, uint32(v11674)+124))
	if v11675 != 0 {
		goto L2461
	} else {
		goto L2462
	}
L2460:
	;
	v11697 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125]))
	if int32(0) <= v11697 {
		goto L2464
	} else {
		goto L2465
	}
L2461:
	;
	v11676 = *(*int64)(unsafe.Add(mBase, uint32(v11675)+16))
	v11677 = *(*int32)(unsafe.Add(mBase, uint32(v11674)+120))
	v11678 = *(*int64)(unsafe.Add(mBase, uint32(v11677)+16))
	v11681 = base.I32_wrap_i64(v11676 - v11678)
	goto L2463
L2462:
	;
	v11681 = int32(0)
	goto L2463
L2463:
	;
	v11682 = *(*int32)(unsafe.Add(mBase, uint32(v11669)+112))
	v11683 = *(*int32)(unsafe.Add(mBase, uint32(v11682)+16))
	v11685 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[30]))
	v11686 = *(*int32)(unsafe.Add(mBase, uint32(v11682)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11685)+64)) = v11686
	*(*int32)(unsafe.Add(mBase, uint32(v11685)+56)) = v11681
	*(*int32)(unsafe.Add(mBase, uint32(v11685)+60)) = v11686 + v11683
	v11691 = *(*int32)(unsafe.Add(mBase, uint32(v11669)))
	v11692 = *(*int64)(unsafe.Add(mBase, uint32(v11691)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v11669)+16)) = v11692 - int64(-8192)
	goto L2460
L2464:
	;
	v11700 = F_close(m, v11697)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[125])) = int32(-1)
	goto L2466
L2465:
	;
	goto L2466
L2466:
	;
	v11705 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[28]))
	F_XLogReaderFree(m, v11705)
	mBase = m.M
	v11707 = m.ExcPending
	if v11707 != 0 {
		goto L32
	} else {
		goto L2467
	}
L2467:
	;
	v11709 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[32]))
	v11710 = *(*int32)(unsafe.Add(mBase, uint32(v11709)+112))
	F_pfree(m, v11710)
	mBase = m.M
	v11712 = m.ExcPending
	if v11712 != 0 {
		goto L32
	} else {
		goto L2468
	}
L2468:
	;
	v11713 = *(*int32)(unsafe.Add(mBase, uint32(v11709)+24))
	F_hash_destroy(m, v11713)
	mBase = m.M
	v11715 = m.ExcPending
	if v11715 != 0 {
		goto L32
	} else {
		goto L2469
	}
L2469:
	;
	F_pfree(m, v11709)
	mBase = m.M
	v11717 = m.ExcPending
	if v11717 != 0 {
		goto L32
	} else {
		goto L2470
	}
L2470:
	;
	v11719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11719 != int32(1) {
		goto L2471
	} else {
		goto L2472
	}
L2471:
	;
	m.G0 = v11666 + int32(1024)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[140])) = int32(1)
	if v9692 != int64(0) {
		goto L2477
	} else {
		goto L2478
	}
L2472:
	;
	v11725 = F_pg_snprintf(m, v11666, int32(1024), int32(_a_F_StartupXLOG_315), int32(0))
	mBase = m.M
	v11726 = m.ExcPending
	if v11726 != 0 {
		goto L32
	} else {
		goto L2473
	}
L2473:
	;
	v11727 = F_unlink(m, v11666)
	mBase = m.M
	v11731 = F_pg_snprintf(m, v11666, int32(1024), int32(_a_F_StartupXLOG_316), int32(0))
	mBase = m.M
	v11732 = m.ExcPending
	if v11732 != 0 {
		goto L32
	} else {
		goto L2474
	}
L2474:
	;
	v11733 = F_unlink(m, v11666)
	mBase = m.M
	v11735 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11735 != int32(1) {
		goto L2471
	} else {
		goto L2475
	}
L2475:
	;
	v11739 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v11739+int32(4))+12)) = int32(0)
	goto L2476
L2476:
	;
	goto L2471
L2477:
	;
	v11753 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[141])))
	if v11753 != int32(1) {
		goto L5
	} else {
		goto L2480
	}
L2478:
	;
	goto L2479
L2479:
	;
	v11916 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[73])))
	*(*uint8)(unsafe.Add(mBase, uint32(v10695)+160)) = uint8(v11916)
	F_UpdateFullPageWrites(m)
	mBase = m.M
	v11919 = m.ExcPending
	if v11919 != 0 {
		goto L32
	} else {
		goto L2510
	}
L2480:
	;
	v11758 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v11759 = *(*int32)(unsafe.Add(mBase, uint32(v11758)+316))
	v11761 = base.B2i32(v11759 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[141])) = uint8(v11761)
	if v11761 == int32(0) {
		goto L5
	} else {
		goto L2481
	}
L2481:
	;
	if v9695&int64(8191) != int64(0) {
		goto L4
	} else {
		goto L2482
	}
L2482:
	;
	v11769 = *(*int32)(unsafe.Add(mBase, uint32(v11758)))
	v11771 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v11758))) = int32(1)
	if v11769 != 0 {
		goto L2483
	} else {
		goto L2484
	}
L2483:
	;
	F_s_lock(m, v11758, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_317), int32(_a_F_StartupXLOG_318))
	mBase = m.M
	v11778 = m.ExcPending
	if v11778 != 0 {
		goto L32
	} else {
		goto L2486
	}
L2484:
	;
	goto L2485
L2485:
	;
	if v9695&base.I64_extend_i32_s(v11771-int32(1)) == int64(0) {
		goto L2487
	} else {
		goto L2488
	}
L2486:
	;
	goto L2485
L2487:
	;
	v11787 = int64(40)
	goto L2489
L2488:
	;
	v11787 = int64(24)
	goto L2489
L2489:
	;
	v11788 = v11787 | v9695
	*(*int32)(unsafe.Add(mBase, uint32(v11758))) = int32(0)
	v11791 = *(*int64)(unsafe.Add(mBase, uint32(v11758)+8))
	v11793 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[129])))
	v11794 = base.I64_div_u_s(v11791, v11793)
	v11796 = v11791 - v11794*v11793
	if base.Ui64(v11796) <= base.Ui64(int64(8151)) {
		goto L2491
	} else {
		goto L2492
	}
L2490:
	;
	v11816 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v11820 = v11794*v11816 + v11814&int64(4294967295)
	if v11820 != v11788 {
		goto L3
	} else {
		goto L2494
	}
L2491:
	;
	v11814 = v11796 + int64(40)
	goto L2490
L2492:
	;
	goto L2493
L2493:
	;
	v11802 = v11796 - int64(8152)
	v11803 = int64(8168)
	v11804 = base.I64_div_u_s(v11802, v11803)
	v11814 = v11802 - v11804*v11803 + v11804<<(uint(int64(13))%64) + int64(8216)
	goto L2490
L2494:
	;
	v11822 = int32(_a_F_StartupXLOG_142)
	v11824 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v11824 + int32(1)
	v11829 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[142]))
	if v11829 == int32(-1) {
		goto L2495
	} else {
		goto L2496
	}
L2495:
	;
	v11834 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[87]))
	v11836 = base.I32_rem_s(v11834, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[142])) = v11836
	v11838 = v11836
	goto L2497
L2496:
	;
	v11838 = v11829
	goto L2497
L2497:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[143])) = v11838
	v11842 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[144]))
	v11847 = F_LWLockAcquire(m, v11842+v11838<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v11848 = m.ExcPending
	if v11848 != 0 {
		goto L32
	} else {
		goto L2498
	}
L2498:
	;
	if v11847 == int32(0) {
		goto L2499
	} else {
		goto L2500
	}
L2499:
	;
	v11851 = int32(_a_F_StartupXLOG_319)
	v11853 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[142]))
	v11857 = base.I32_rem_s(v11853+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[142])) = v11857
	goto L2501
L2500:
	;
	goto L2501
L2501:
	;
	v11859 = F_GetXLogBuffer(m, v9695, v10651)
	mBase = m.M
	v11860 = m.ExcPending
	if v11860 != 0 {
		goto L32
	} else {
		goto L2502
	}
L2502:
	;
	v11861 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11859)+2)))
	v11863 = v11861 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v11859)+2)) = uint16(v11863)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v11866 = m.ExcPending
	if v11866 != 0 {
		goto L32
	} else {
		goto L2503
	}
L2503:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v11868 = m.ExcPending
	if v11868 != 0 {
		goto L32
	} else {
		goto L2504
	}
L2504:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v9692
	v11873 = m.G0
	v11874 = int32(16)
	v11875 = v11873 - v11874
	m.G0 = v11875
	F_gettimeofday(m, v11875)
	mBase = m.M
	v11878 = *(*int64)(unsafe.Add(mBase, uint32(v11875)))
	v11879 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11875)+8)))
	m.G0 = v11875 + v11874
	goto L2505
L2505:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[145]))) = v11879 + v11878*int64(1000000) - int64(946684800000000)
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(16))
	mBase = m.M
	v11893 = m.ExcPending
	if v11893 != 0 {
		goto L32
	} else {
		goto L2506
	}
L2506:
	;
	v11896 = F_XLogInsert(m, int32(0), int32(208))
	mBase = m.M
	v11897 = m.ExcPending
	if v11897 != 0 {
		goto L32
	} else {
		goto L2507
	}
L2507:
	;
	v11899 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[146]))
	if v11899 != v11788 {
		goto L2
	} else {
		goto L2508
	}
L2508:
	;
	F_XLogFlush(m, v11896)
	mBase = m.M
	v11902 = m.ExcPending
	if v11902 != 0 {
		goto L32
	} else {
		goto L2509
	}
L2509:
	;
	v11903 = int32(_a_F_StartupXLOG_142)
	v11905 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v11905 - int32(1)
	goto L2479
L2510:
	;
	v11920 = int32(0)
	if v6010 == v11920 {
		v12033 = v11920
		goto L2511
	} else {
		goto L2512
	}
L2511:
	;
	v12037 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	v12039 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12040 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+172))
	if v12037 != v12040 {
		goto L2532
	} else {
		goto L2533
	}
L2512:
	;
	v11924 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v11924 != int32(1) {
		goto L2513
	} else {
		goto L2514
	}
L2513:
	;
	F_RequestCheckpoint(m, int32(38))
	mBase = m.M
	v12032 = m.ExcPending
	if v12032 != 0 {
		goto L32
	} else {
		goto L2530
	}
L2514:
	;
	v11928 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[15])))
	if v11928&int32(1) == int32(0) {
		goto L2513
	} else {
		goto L2515
	}
L2515:
	;
	v11933 = F_PromoteIsTriggered(m)
	mBase = m.M
	v11934 = m.ExcPending
	if v11934 != 0 {
		goto L32
	} else {
		goto L2516
	}
L2516:
	;
	if v11933 == int32(0) {
		goto L2513
	} else {
		goto L2517
	}
L2517:
	;
	v11938 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[141])))
	if v11938 != int32(1) {
		goto L1
	} else {
		goto L2518
	}
L2518:
	;
	v11943 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v11944 = *(*int32)(unsafe.Add(mBase, uint32(v11943)+316))
	v11946 = base.B2i32(v11944 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[141])) = uint8(v11946)
	if v11946 == int32(0) {
		goto L1
	} else {
		goto L2519
	}
L2519:
	;
	v11953 = m.G0
	v11954 = int32(16)
	v11955 = v11953 - v11954
	m.G0 = v11955
	F_gettimeofday(m, v11955)
	mBase = m.M
	v11958 = *(*int64)(unsafe.Add(mBase, uint32(v11955)))
	v11959 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11955)+8)))
	m.G0 = v11955 + v11954
	goto L2520
L2520:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v11959 + v11958*int64(1000000) - int64(946684800000000)
	v11970 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v11970
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v11973 = m.ExcPending
	if v11973 != 0 {
		goto L32
	} else {
		goto L2521
	}
L2521:
	;
	v11975 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v11976 = *(*int32)(unsafe.Add(mBase, uint32(v11975)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[145]))) = v11976
	v11978 = *(*int32)(unsafe.Add(mBase, uint32(v11975)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v11978
	F_WALInsertLockRelease(m)
	mBase = m.M
	v11981 = m.ExcPending
	if v11981 != 0 {
		goto L32
	} else {
		goto L2522
	}
L2522:
	;
	v11982 = int32(1)
	v11983 = int32(_a_F_StartupXLOG_142)
	v11985 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v11985 + v11982
	F_XLogBeginInsert(m)
	mBase = m.M
	v11990 = m.ExcPending
	if v11990 != 0 {
		goto L32
	} else {
		goto L2523
	}
L2523:
	;
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(24))
	mBase = m.M
	v11995 = m.ExcPending
	if v11995 != 0 {
		goto L32
	} else {
		goto L2524
	}
L2524:
	;
	v11998 = F_XLogInsert(m, int32(0), int32(144))
	mBase = m.M
	v11999 = m.ExcPending
	if v11999 != 0 {
		goto L32
	} else {
		goto L2525
	}
L2525:
	;
	F_XLogFlush(m, v11998)
	mBase = m.M
	v12001 = m.ExcPending
	if v12001 != 0 {
		goto L32
	} else {
		goto L2526
	}
L2526:
	;
	v12003 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12007 = F_LWLockAcquire(m, v12003+int32(1152), int32(0))
	mBase = m.M
	v12008 = m.ExcPending
	if v12008 != 0 {
		goto L32
	} else {
		goto L2527
	}
L2527:
	;
	v12010 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v12010)+136)) = v11998
	v12012 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[145])))
	*(*int32)(unsafe.Add(mBase, uint32(v12010)+144)) = v12012
	v12015 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	F_update_controlfile(m, v12015, v12010)
	mBase = m.M
	v12017 = m.ExcPending
	if v12017 != 0 {
		goto L32
	} else {
		goto L2528
	}
L2528:
	;
	v12019 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12019+int32(1152))
	mBase = m.M
	v12023 = m.ExcPending
	if v12023 != 0 {
		goto L32
	} else {
		goto L2529
	}
L2529:
	;
	v12024 = int32(_a_F_StartupXLOG_142)
	v12026 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[60])) = v12026 - int32(1)
	v12033 = v11982
	goto L2511
L2530:
	;
	v12033 = v11920
	goto L2511
L2531:
	;
	v12156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[13])))
	if v12156 != int32(1) {
		goto L2551
	} else {
		goto L2552
	}
L2532:
	;
	v12071 = int32(0)
	if base.B2i32(v12037 == v12040)&base.B2i32(v12037 <= v12071) == v12071 {
		goto L2541
	} else {
		goto L2542
	}
L2533:
	;
	v12043 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[147])))
	v12044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12039)+176)))
	if v12043 != v12044 {
		goto L2532
	} else {
		goto L2534
	}
L2534:
	;
	v12047 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	v12048 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+180))
	if v12047 != v12048 {
		goto L2532
	} else {
		goto L2535
	}
L2535:
	;
	v12051 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[149]))
	v12052 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+184))
	if v12051 != v12052 {
		goto L2532
	} else {
		goto L2536
	}
L2536:
	;
	v12055 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[120]))
	v12056 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+188))
	if v12055 != v12056 {
		goto L2532
	} else {
		goto L2537
	}
L2537:
	;
	v12059 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	v12060 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+192))
	if v12059 != v12060 {
		goto L2532
	} else {
		goto L2538
	}
L2538:
	;
	v12063 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	v12064 = *(*int32)(unsafe.Add(mBase, uint32(v12039)+196))
	if v12063 != v12064 {
		goto L2532
	} else {
		goto L2539
	}
L2539:
	;
	v12067 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[152])))
	v12068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12039)+200)))
	if v12067 == v12068 {
		goto L2531
	} else {
		goto L2540
	}
L2540:
	;
	goto L2532
L2541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[153]))) = v12037
	v12078 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[6]))) = v12078
	v12081 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[149]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[93]))) = v12081
	v12084 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[145]))) = v12084
	v12087 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[92]))) = v12087
	v12090 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[91]))) = v12090
	v12093 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[147])))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[95]))) = uint8(v12093)
	v12096 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[152])))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_StartupXLOG[154]))) = uint8(v12096)
	F_XLogBeginInsert(m)
	mBase = m.M
	v12099 = m.ExcPending
	if v12099 != 0 {
		goto L32
	} else {
		goto L2544
	}
L2542:
	;
	goto L2543
L2543:
	;
	v12112 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12116 = F_LWLockAcquire(m, v12112+int32(1152), int32(0))
	mBase = m.M
	v12117 = m.ExcPending
	if v12117 != 0 {
		goto L32
	} else {
		goto L2548
	}
L2544:
	;
	F_XLogRegisterData(m, v39+int32(_a_F_StartupXLOG_1), int32(28))
	mBase = m.M
	v12104 = m.ExcPending
	if v12104 != 0 {
		goto L32
	} else {
		goto L2545
	}
L2545:
	;
	v12107 = F_XLogInsert(m, int32(0), int32(96))
	mBase = m.M
	v12108 = m.ExcPending
	if v12108 != 0 {
		goto L32
	} else {
		goto L2546
	}
L2546:
	;
	F_XLogFlush(m, v12107)
	mBase = m.M
	v12110 = m.ExcPending
	if v12110 != 0 {
		goto L32
	} else {
		goto L2547
	}
L2547:
	;
	goto L2543
L2548:
	;
	v12119 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	v12121 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+180)) = v12121
	v12124 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[149]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+184)) = v12124
	v12127 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+188)) = v12127
	v12130 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[150]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+192)) = v12130
	v12133 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+196)) = v12133
	v12136 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[61]))
	*(*int32)(unsafe.Add(mBase, uint32(v12119)+172)) = v12136
	v12139 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[147])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12119)+176)) = uint8(v12139)
	v12142 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[152])))
	*(*uint8)(unsafe.Add(mBase, uint32(v12119)+200)) = uint8(v12142)
	v12145 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	F_update_controlfile(m, v12145, v12119)
	mBase = m.M
	v12147 = m.ExcPending
	if v12147 != 0 {
		goto L32
	} else {
		goto L2549
	}
L2549:
	;
	v12149 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12149+int32(1152))
	mBase = m.M
	v12153 = m.ExcPending
	if v12153 != 0 {
		goto L32
	} else {
		goto L2550
	}
L2550:
	;
	goto L2531
L2551:
	;
	v12319 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[152])))
	if v12319 == int32(0) {
		goto L2582
	} else {
		goto L2583
	}
L2552:
	;
	v12160 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[155]))
	if v12160 == int32(0) {
		goto L2553
	} else {
		goto L2554
	}
L2553:
	;
	F_RemoveNonParentXlogFiles(m, v10703, v10651)
	mBase = m.M
	v12172 = m.ExcPending
	if v12172 != 0 {
		goto L32
	} else {
		goto L2557
	}
L2554:
	;
	v12163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12160))))
	if v12163 == int32(0) {
		goto L2553
	} else {
		goto L2555
	}
L2555:
	;
	F_ExecuteRecoveryCommand(m, v12160, int32(_a_F_StartupXLOG_320), int32(1), int32(134217774))
	mBase = m.M
	v12170 = m.ExcPending
	if v12170 != 0 {
		goto L32
	} else {
		goto L2556
	}
L2556:
	;
	goto L2553
L2557:
	;
	v12174 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4]))
	if v10703&base.I64_extend_i32_s(v12174-int32(1)) == int64(0) {
		goto L2551
	} else {
		goto L2558
	}
L2558:
	;
	v12182 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[128]))
	if v12182 <= int32(0) {
		goto L2551
	} else {
		goto L2559
	}
L2559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3616)) = v9708
	v12188 = base.I64_extend_i32_s(v12174)
	v12189 = base.I64_div_u_s(v10703-int64(1), v12188)
	v12191 = base.I64_div_u_s(int64(4294967296), v12188)
	v12192 = base.I64_div_u_s(v12189, v12191)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3620)) = uint32(v12192)
	v12195 = v12189 - v12192*v12191
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3624)) = uint32(v12195)
	v12198 = v39 + int32(_a_F_StartupXLOG_3)
	v12203 = F_pg_snprintf(m, v12198, int32(64), int32(_a_F_StartupXLOG_277), v39+int32(3616))
	mBase = m.M
	v12204 = m.ExcPending
	if v12204 != 0 {
		goto L32
	} else {
		goto L2560
	}
L2560:
	;
	v12205 = m.G0
	v12207 = v12205 - int32(1168)
	m.G0 = v12207
	*(*int32)(unsafe.Add(mBase, uint32(v12207)+32)) = v12198
	*(*int32)(unsafe.Add(mBase, uint32(v12207)+36)) = int32(_a_F_StartupXLOG_321)
	v12213 = v12207 + int32(144)
	v12218 = F_pg_snprintf(m, v12213, int32(1024), int32(_a_F_StartupXLOG_322), v12207+int32(32))
	mBase = m.M
	v12219 = m.ExcPending
	if v12219 != 0 {
		goto L32
	} else {
		goto L2561
	}
L2561:
	;
	v12220 = int32(1)
	v12222 = v12207 + int32(48)
	v12225 = F___fstatat(m, int32(-100), v12213, v12222, int32(0))
	mBase = m.M
	goto L2563
L2562:
	;
	m.G0 = v12207 + int32(1168)
	if v12254 != 0 {
		goto L2551
	} else {
		goto L2570
	}
L2563:
	;
	if v12225 == int32(0) {
		v12254 = v12220
		goto L2562
	} else {
		goto L2564
	}
L2564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12207)+20)) = int32(_a_F_StartupXLOG_323)
	*(*int32)(unsafe.Add(mBase, uint32(v12207)+16)) = v12198
	v12235 = F_pg_snprintf(m, v12213, int32(1024), int32(_a_F_StartupXLOG_322), v12207+int32(16))
	mBase = m.M
	v12236 = m.ExcPending
	if v12236 != 0 {
		goto L32
	} else {
		goto L2565
	}
L2565:
	;
	v12239 = F___fstatat(m, int32(-100), v12213, v12222, int32(0))
	mBase = m.M
	goto L2566
L2566:
	;
	if v12239 == int32(0) {
		v12254 = v12220
		goto L2562
	} else {
		goto L2567
	}
L2567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12207)+4)) = int32(_a_F_StartupXLOG_321)
	*(*int32)(unsafe.Add(mBase, uint32(v12207))) = v12198
	v12247 = F_pg_snprintf(m, v12213, int32(1024), int32(_a_F_StartupXLOG_322), v12207)
	mBase = m.M
	v12248 = m.ExcPending
	if v12248 != 0 {
		goto L32
	} else {
		goto L2568
	}
L2568:
	;
	v12251 = F___fstatat(m, int32(-100), v12213, v12222, int32(0))
	mBase = m.M
	goto L2569
L2569:
	;
	v12254 = base.B2i32(v12251 == int32(0))
	goto L2562
L2570:
	;
	v12259 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupXLOG[156])))
	if v12259 == int32(1) {
		goto L2571
	} else {
		goto L2572
	}
L2571:
	;
	F_WaitForWalSummarization(m, v10703)
	mBase = m.M
	v12263 = m.ExcPending
	if v12263 != 0 {
		goto L32
	} else {
		goto L2574
	}
L2572:
	;
	goto L2573
L2573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3600)) = v9708
	v12267 = int64(*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[4])))
	v12268 = base.I64_div_u_s(int64(4294967296), v12267)
	v12269 = base.I64_div_u_s(v12189, v12268)
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3604)) = uint32(v12269)
	v12272 = v12189 - v12268*v12269
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3608)) = uint32(v12272)
	v12275 = v39 + int32(_a_F_StartupXLOG_1)
	v12280 = F_pg_snprintf(m, v12275, int32(1024), int32(_a_F_StartupXLOG_278), v39+int32(3600))
	mBase = m.M
	v12281 = m.ExcPending
	if v12281 != 0 {
		goto L32
	} else {
		goto L2575
	}
L2574:
	;
	goto L2573
L2575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3584)) = v39 + int32(_a_F_StartupXLOG_3)
	v12286 = v39 + int32(_a_F_StartupXLOG_282)
	v12291 = F_pg_snprintf(m, v12286, int32(64), int32(_a_F_StartupXLOG_324), v39+int32(3584))
	mBase = m.M
	v12292 = m.ExcPending
	if v12292 != 0 {
		goto L32
	} else {
		goto L2576
	}
L2576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3568)) = v12275
	v12295 = v39 + int32(_a_F_StartupXLOG_4)
	v12300 = F_pg_snprintf(m, v12295, int32(1024), int32(_a_F_StartupXLOG_324), v39+int32(3568))
	mBase = m.M
	v12301 = m.ExcPending
	if v12301 != 0 {
		goto L32
	} else {
		goto L2577
	}
L2577:
	;
	F_XLogArchiveCleanup(m, v12286)
	mBase = m.M
	v12303 = m.ExcPending
	if v12303 != 0 {
		goto L32
	} else {
		goto L2578
	}
L2578:
	;
	v12305 = F_durable_rename(m, v12275, v12295, int32(21))
	mBase = m.M
	v12306 = m.ExcPending
	if v12306 != 0 {
		goto L32
	} else {
		goto L2579
	}
L2579:
	;
	F_XLogArchiveNotify(m, v12286)
	mBase = m.M
	v12308 = m.ExcPending
	if v12308 != 0 {
		goto L32
	} else {
		goto L2580
	}
L2580:
	;
	goto L2551
L2581:
	;
	v12358 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12362 = F_LWLockAcquire(m, v12358+int32(1152), int32(0))
	mBase = m.M
	v12363 = m.ExcPending
	if v12363 != 0 {
		goto L32
	} else {
		goto L2589
	}
L2582:
	;
	v12323 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	v12327 = F_LWLockAcquire(m, v12323+int32(_a_F_StartupXLOG_325), int32(0))
	mBase = m.M
	v12328 = m.ExcPending
	if v12328 != 0 {
		goto L32
	} else {
		goto L2585
	}
L2583:
	;
	goto L2584
L2584:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v12355 = m.ExcPending
	if v12355 != 0 {
		goto L32
	} else {
		goto L2588
	}
L2585:
	;
	v12330 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[157]))
	v12331 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12330)+16)) = uint16(v12331)
	*(*int64)(unsafe.Add(mBase, uint32(v12330)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v12330))) = v12331
	*(*uint8)(unsafe.Add(mBase, uint32(v12330)+24)) = uint8(v12331)
	v12340 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[48]))
	*(*int64)(unsafe.Add(mBase, uint32(v12340)+40)) = int64(0)
	v12346 = F_SlruScanDirectory(m, int32(_a_F_StartupXLOG_326), int32(290), v12331)
	mBase = m.M
	v12347 = m.ExcPending
	if v12347 != 0 {
		goto L32
	} else {
		goto L2586
	}
L2586:
	;
	v12349 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12349+int32(_a_F_StartupXLOG_325))
	mBase = m.M
	v12353 = m.ExcPending
	if v12353 != 0 {
		goto L32
	} else {
		goto L2587
	}
L2587:
	;
	goto L2581
L2588:
	;
	goto L2581
L2589:
	;
	v12365 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v12365)+16)) = int32(6)
	v12369 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	v12370 = *(*int32)(unsafe.Add(mBase, uint32(v12369)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v12369)+440)) = int32(1)
	if v12370 != 0 {
		goto L2590
	} else {
		goto L2591
	}
L2590:
	;
	v12374 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	F_s_lock(m, v12374+int32(440), int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_327), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12381 = m.ExcPending
	if v12381 != 0 {
		goto L32
	} else {
		goto L2593
	}
L2591:
	;
	goto L2592
L2592:
	;
	v12383 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v12383)+440)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12383)+316)) = int32(2)
	v12389 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[43]))
	v12391 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[2]))
	F_update_controlfile(m, v12389, v12391)
	mBase = m.M
	v12393 = m.ExcPending
	if v12393 != 0 {
		goto L32
	} else {
		goto L2594
	}
L2593:
	;
	goto L2592
L2594:
	;
	v12395 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[70]))
	F_LWLockRelease(m, v12395+int32(1152))
	mBase = m.M
	v12399 = m.ExcPending
	if v12399 != 0 {
		goto L32
	} else {
		goto L2595
	}
L2595:
	;
	v12401 = *(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[89]))
	if v12401 != 0 {
		goto L2596
	} else {
		goto L2597
	}
L2596:
	;
	F_ShutdownRecoveryTransactionEnvironment(m)
	mBase = m.M
	v12403 = m.ExcPending
	if v12403 != 0 {
		goto L32
	} else {
		goto L2599
	}
L2597:
	;
	goto L2598
L2598:
	;
	v12404 = int32(1)
	F_WalSndWakeup(m, v12404, v12404)
	mBase = m.M
	v12407 = m.ExcPending
	if v12407 != 0 {
		goto L32
	} else {
		goto L2600
	}
L2599:
	;
	goto L2598
L2600:
	;
	if v12033 != 0 {
		goto L2601
	} else {
		goto L2602
	}
L2601:
	;
	F_RequestCheckpoint(m, int32(8))
	mBase = m.M
	v12410 = m.ExcPending
	if v12410 != 0 {
		goto L32
	} else {
		goto L2604
	}
L2602:
	;
	goto L2603
L2603:
	;
	m.G0 = v35
	return
L2604:
	;
	goto L2603
L2605:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12418 = m.ExcPending
	if v12418 != 0 {
		goto L32
	} else {
		goto L2606
	}
L2606:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_328), int32(0))
	mBase = m.M
	v12422 = m.ExcPending
	if v12422 != 0 {
		goto L32
	} else {
		goto L2607
	}
L2607:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_329), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12427 = m.ExcPending
	if v12427 != 0 {
		goto L32
	} else {
		goto L2608
	}
L2608:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2609:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v12434 = m.ExcPending
	if v12434 != 0 {
		goto L32
	} else {
		goto L2610
	}
L2610:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_330), int32(0))
	mBase = m.M
	v12438 = m.ExcPending
	if v12438 != 0 {
		goto L32
	} else {
		goto L2611
	}
L2611:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_331), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12443 = m.ExcPending
	if v12443 != 0 {
		goto L32
	} else {
		goto L2612
	}
L2612:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2613:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12449 = m.ExcPending
	if v12449 != 0 {
		goto L32
	} else {
		goto L2614
	}
L2614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3984)) = int32(_a_F_StartupXLOG_5)
	F_errmsg(m, int32(_a_F_StartupXLOG_26), v39+int32(3984))
	mBase = m.M
	v12456 = m.ExcPending
	if v12456 != 0 {
		goto L32
	} else {
		goto L2615
	}
L2615:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_332), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12461 = m.ExcPending
	if v12461 != 0 {
		goto L32
	} else {
		goto L2616
	}
L2616:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2617:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12467 = m.ExcPending
	if v12467 != 0 {
		goto L32
	} else {
		goto L2618
	}
L2618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3952)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_333), v39+int32(3952))
	mBase = m.M
	v12475 = m.ExcPending
	if v12475 != 0 {
		goto L32
	} else {
		goto L2619
	}
L2619:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_334), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12480 = m.ExcPending
	if v12480 != 0 {
		goto L32
	} else {
		goto L2620
	}
L2620:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3904)) = v39 + int32(_a_F_StartupXLOG_1)
	F_errmsg(m, int32(_a_F_StartupXLOG_333), v39+int32(3904))
	mBase = m.M
	v12492 = m.ExcPending
	if v12492 != 0 {
		goto L32
	} else {
		goto L2622
	}
L2622:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_335), int32(_a_F_StartupXLOG_28))
	mBase = m.M
	v12497 = m.ExcPending
	if v12497 != 0 {
		goto L32
	} else {
		goto L2623
	}
L2623:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2624:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v12504 = m.ExcPending
	if v12504 != 0 {
		goto L32
	} else {
		goto L2625
	}
L2625:
	;
	F_errmsg(m, int32(_a_F_StartupXLOG_336), int32(0))
	mBase = m.M
	v12508 = m.ExcPending
	if v12508 != 0 {
		goto L32
	} else {
		goto L2626
	}
L2626:
	;
	F_errhint(m, int32(_a_F_StartupXLOG_337), int32(0))
	mBase = m.M
	v12512 = m.ExcPending
	if v12512 != 0 {
		goto L32
	} else {
		goto L2627
	}
L2627:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_338), int32(_a_F_StartupXLOG_7))
	mBase = m.M
	v12517 = m.ExcPending
	if v12517 != 0 {
		goto L32
	} else {
		goto L2628
	}
L2628:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2629:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12523 = m.ExcPending
	if v12523 != 0 {
		goto L32
	} else {
		goto L2630
	}
L2630:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3536)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_148), v39+int32(3536))
	mBase = m.M
	v12531 = m.ExcPending
	if v12531 != 0 {
		goto L32
	} else {
		goto L2631
	}
L2631:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3435), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12536 = m.ExcPending
	if v12536 != 0 {
		goto L32
	} else {
		goto L2632
	}
L2632:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2633:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12542 = m.ExcPending
	if v12542 != 0 {
		goto L32
	} else {
		goto L2634
	}
L2634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3552)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_293), v39+int32(3552))
	mBase = m.M
	v12550 = m.ExcPending
	if v12550 != 0 {
		goto L32
	} else {
		goto L2635
	}
L2635:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3449), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12555 = m.ExcPending
	if v12555 != 0 {
		goto L32
	} else {
		goto L2636
	}
L2636:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3760)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_62), v39+int32(3760))
	mBase = m.M
	v12565 = m.ExcPending
	if v12565 != 0 {
		goto L32
	} else {
		goto L2638
	}
L2638:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3481), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12570 = m.ExcPending
	if v12570 != 0 {
		goto L32
	} else {
		goto L2639
	}
L2639:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2640:
	;
	v12578 = v12572
	goto L2642
L2641:
	;
	v12578 = int32(51)
	goto L2642
L2642:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v12578
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12583 = m.ExcPending
	if v12583 != 0 {
		goto L32
	} else {
		goto L2643
	}
L2643:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12585 = m.ExcPending
	if v12585 != 0 {
		goto L32
	} else {
		goto L2644
	}
L2644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3744)) = v12574
	F_errmsg(m, int32(_a_F_StartupXLOG_296), v39+int32(3744))
	mBase = m.M
	v12591 = m.ExcPending
	if v12591 != 0 {
		goto L32
	} else {
		goto L2645
	}
L2645:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3505), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12596 = m.ExcPending
	if v12596 != 0 {
		goto L32
	} else {
		goto L2646
	}
L2646:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2647:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12602 = m.ExcPending
	if v12602 != 0 {
		goto L32
	} else {
		goto L2648
	}
L2648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3712)) = v39 + int32(_a_F_StartupXLOG_3)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3712))
	mBase = m.M
	v12610 = m.ExcPending
	if v12610 != 0 {
		goto L32
	} else {
		goto L2649
	}
L2649:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3520), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12615 = m.ExcPending
	if v12615 != 0 {
		goto L32
	} else {
		goto L2650
	}
L2650:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2651:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12621 = m.ExcPending
	if v12621 != 0 {
		goto L32
	} else {
		goto L2652
	}
L2652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3696)) = v39 + int32(_a_F_StartupXLOG_4)
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3696))
	mBase = m.M
	v12629 = m.ExcPending
	if v12629 != 0 {
		goto L32
	} else {
		goto L2653
	}
L2653:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(3525), int32(_a_F_StartupXLOG_281))
	mBase = m.M
	v12634 = m.ExcPending
	if v12634 != 0 {
		goto L32
	} else {
		goto L2654
	}
L2654:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2655:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartupXLOG[3])) = v12636
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12648 = m.ExcPending
	if v12648 != 0 {
		goto L32
	} else {
		goto L2656
	}
L2656:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v12650 = m.ExcPending
	if v12650 != 0 {
		goto L32
	} else {
		goto L2657
	}
L2657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+3824)) = v12638
	F_errmsg(m, int32(_a_F_StartupXLOG_153), v39+int32(3824))
	mBase = m.M
	v12656 = m.ExcPending
	if v12656 != 0 {
		goto L32
	} else {
		goto L2658
	}
L2658:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_339), int32(_a_F_StartupXLOG_340))
	mBase = m.M
	v12661 = m.ExcPending
	if v12661 != 0 {
		goto L32
	} else {
		goto L2659
	}
L2659:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2660:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_341), int32(0))
	mBase = m.M
	v12671 = m.ExcPending
	if v12671 != 0 {
		goto L32
	} else {
		goto L2661
	}
L2661:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_342), int32(_a_F_StartupXLOG_343))
	mBase = m.M
	v12676 = m.ExcPending
	if v12676 != 0 {
		goto L32
	} else {
		goto L2662
	}
L2662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2663:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3668)) = uint32(v9695)
	v12683 = int64(base.Ui64(v9695) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3664)) = uint32(v12683)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_344), v39+int32(3664))
	mBase = m.M
	v12689 = m.ExcPending
	if v12689 != 0 {
		goto L32
	} else {
		goto L2664
	}
L2664:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_345), int32(_a_F_StartupXLOG_343))
	mBase = m.M
	v12694 = m.ExcPending
	if v12694 != 0 {
		goto L32
	} else {
		goto L2665
	}
L2665:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2666:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3652)) = uint32(v11820)
	v12701 = int64(base.Ui64(v11820) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3648)) = uint32(v12701)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_346), v39+int32(3648))
	mBase = m.M
	v12707 = m.ExcPending
	if v12707 != 0 {
		goto L32
	} else {
		goto L2667
	}
L2667:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_347), int32(_a_F_StartupXLOG_343))
	mBase = m.M
	v12712 = m.ExcPending
	if v12712 != 0 {
		goto L32
	} else {
		goto L2668
	}
L2668:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2669:
	;
	v12718 = *(*int64)(unsafe.Add(mBase, _c_F_StartupXLOG[146]))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3636)) = uint32(v12718)
	v12721 = int64(base.Ui64(v12718) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v39)+3632)) = uint32(v12721)
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_348), v39+int32(3632))
	mBase = m.M
	v12727 = m.ExcPending
	if v12727 != 0 {
		goto L32
	} else {
		goto L2670
	}
L2670:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_349), int32(_a_F_StartupXLOG_343))
	mBase = m.M
	v12732 = m.ExcPending
	if v12732 != 0 {
		goto L32
	} else {
		goto L2671
	}
L2671:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2672:
	;
	F_errmsg_internal(m, int32(_a_F_StartupXLOG_350), int32(0))
	mBase = m.M
	v12741 = m.ExcPending
	if v12741 != 0 {
		goto L32
	} else {
		goto L2673
	}
L2673:
	;
	F_errfinish(m, int32(_a_F_StartupXLOG_6), int32(_a_F_StartupXLOG_351), int32(_a_F_StartupXLOG_352))
	mBase = m.M
	v12746 = m.ExcPending
	if v12746 != 0 {
		goto L32
	} else {
		goto L2674
	}
L2674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
