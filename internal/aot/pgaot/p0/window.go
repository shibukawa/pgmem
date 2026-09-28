package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecWindowAgg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int64
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int64
	_ = v611
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v615 int32
	_ = v615
	var v644 int32
	_ = v644
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int64
	_ = v679
	var v680 int64
	_ = v680
	var v683 int32
	_ = v683
	var v691 int64
	_ = v691
	var v693 int64
	_ = v693
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int64
	_ = v727
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int64
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v760 int32
	_ = v760
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int64
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v931 int32
	_ = v931
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v989 int32
	_ = v989
	var v992 int64
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int64
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int64
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int64
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1021 int64
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int64
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1044 int64
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int64
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1060 int64
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int64
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int64
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int64
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1120 int64
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1139 int32
	_ = v1139
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1197 int64
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int64
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int64
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int64
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int64
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1289 int64
	_ = v1289
	var v1290 int64
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1322 int32
	_ = v1322
	var v1327 int64
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int64
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1352 int32
	_ = v1352
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1372 int64
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int64
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int64
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1546 int32
	_ = v1546
	var v1547 int64
	_ = v1547
	var v1555 int32
	_ = v1555
	var v1557 int64
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int64
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1610 int64
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1618 int64
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int64
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int64
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1666 int64
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int64
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1686 int64
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1767 int64
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	var v1869 int32
	_ = v1869
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
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
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1914 int32
	_ = v1914
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1925 int64
	_ = v1925
	var v1926 int64
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1939 int64
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int64
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int64
	_ = v1988
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2011 int32
	_ = v2011
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2068 int32
	_ = v2068
	var v2075 int32
	_ = v2075
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2112 int32
	_ = v2112
	var v2113 int64
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2134 int64
	_ = v2134
	var v2135 int64
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2143 int64
	_ = v2143
	var v2144 int64
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2157 int64
	_ = v2157
	var v2158 int64
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2193 int64
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int64
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2201 int64
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2249 int64
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2262 int64
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2276 int32
	_ = v2276
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2308 int64
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2338 int32
	_ = v2338
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2373 int32
	_ = v2373
	var v2374 int64
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2397 int32
	_ = v2397
	var v2398 int64
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2420 int32
	_ = v2420
	var v2428 int32
	_ = v2428
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2447 int64
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2473 int32
	_ = v2473
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2497 int32
	_ = v2497
	var v2521 int32
	_ = v2521
	var v2525 int32
	_ = v2525
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2557 int32
	_ = v2557
	var v2583 int32
	_ = v2583
	var v2592 int32
	_ = v2592
	var v2612 int32
	_ = v2612
	var v2641 int32
	_ = v2641
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2652 int64
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2658 int32
	_ = v2658
	var v2661 float64
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	v25 = m.G0
	v27 = v25 - int32(1632)
	m.G0 = v27
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[0]))
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v35 == int32(0) {
		v2669 = int32(0)
		v2670 = v27
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v2670 + int32(1632)
	return v2669
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+384)))
	if v38 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = m.G0
	v43 = v41 - int32(16)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v46&int32(_a_F_ExecWindowAgg_0) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L10
L10:
	;
	v212 = l0
	v216 = v27
	v228 = v27 + int32(32)
	goto L50
L11:
	;
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L46
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L42
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L38
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L34
	}
L16:
	;
	if v46&int32(_a_F_ExecWindowAgg_1) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v51 = int32(_a_F_ExecWindowAgg_2)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int64)(m, v53, v45, v43+int32(15))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v52
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	if v64 == int32(1) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v69 = F_exprType(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_get_typlenbyval(m, v69, v43+int32(12), v43+int32(11))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+11)))
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+12)))
	v79 = F_datumCopy(m, v60, v77, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v79
	if v46&int32(12) == int32(0) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	if v60 < int64(0) {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+384)) = uint8(v135)
	m.G0 = v43 + int32(16)
	goto L11
L26:
	;
	v95 = int32(_a_F_ExecWindowAgg_2)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v99
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	v104 = m.T0[v103].(func(*base.Module, int32, int32, int32) int64)(m, v97, v45, v43+int32(15))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v96
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	if v108 == int32(1) {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v113 = F_exprType(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_get_typlenbyval(m, v113, v43+int32(12), v43+int32(11))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+11)))
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+12)))
	v123 = F_datumCopy(m, v104, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v123
	if v46&int32(12) == int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v104 < int64(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_3), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2237), int32(_a_F_ExecWindowAgg_5))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errcode(m, int32(50593922))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_6), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2251), int32(_a_F_ExecWindowAgg_5))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_7), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2264), int32(_a_F_ExecWindowAgg_5))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(50593922))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_8), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2278), int32(_a_F_ExecWindowAgg_5))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+386)))
	if v236 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v2669 = v2360
	v2670 = v216
	goto L6
L52:
	;
	F_spool_tuples(m, v212, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L57
	}
L53:
	;
	F_begin_partition(m, v212)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v242 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v212)+388)) = uint16(v242)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	v246 = v244 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v212)+176)) = v246
	v248 = v246
	goto L52
L56:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	v248 = v241
	goto L52
L57:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+385)))
	if v251 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v2665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+319)))
	if v2665 != 0 {
		goto L50
	} else {
		goto L452
	}
L59:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	if v2641 == int32(0) {
		v2669 = v2360
		v2670 = v216
		goto L6
	} else {
		goto L448
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+224)) = int32(2)
	goto L59
L61:
	;
	v2612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+224)) = v2612
	v2669 = v2612
	v2670 = v2592
	goto L6
L62:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	F_MemoryContextReset(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L68
	}
L63:
	;
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v212)+168))
	if v254 < v255 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_release_partition(m, v212)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+387)))
	if v259 != int32(1) {
		v2592 = v216
		goto L61
	} else {
		goto L66
	}
L66:
	;
	F_begin_partition(m, v212)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+224)) = int32(1)
	goto L62
L68:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v212)+144))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v212)+148))
	F_tuplestore_select_read_pointer(m, v270, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v212)+228))
	if v274&int32(_a_F_ExecWindowAgg_9) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v212)+152))
	if int32(0) <= v2338 {
		goto L415
	} else {
		goto L416
	}
L71:
	;
	v2208 = int32(0)
	if v670 != int32(1) {
		goto L408
	} else {
		goto L409
	}
L72:
	;
	v1869 = int32(0)
	goto L354
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L4
	} else {
		goto L351
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L4
	} else {
		goto L348
	}
L75:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v212)+224))
	if v355 != int32(1) {
		goto L70
	} else {
		goto L94
	}
L76:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v212)+144))
	v345 = int32(1)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v212)+112))
	v348 = F_tuplestore_gettupleslot(m, v344, v345, v345, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L92
	}
L77:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	if v279 <= int64(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v212)+412))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v212)+112))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+32))
	m.T0[v285].(func(*base.Module, int32, int32))(m, v282, v283)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v212)+144))
	v289 = int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v212)+112))
	v292 = F_tuplestore_gettupleslot(m, v288, v289, v289, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	if v292 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L81:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+96))
	if v297 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v212)+412))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	m.T0[v341].(func(*base.Module, int32))(m, v339)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L91
	}
L83:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v212)+412))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v212)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v301)+8)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v301)+12)) = v300
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v212)+140))
	if v305 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	F_MemoryContextReset(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v311 = int32(_a_F_ExecWindowAgg_2)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v314
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v305)+24))
	v319 = m.T0[v318].(func(*base.Module, int32, int32, int32) int64)(m, v305, v301, v216+int32(8))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L82
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v312
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	F_MemoryContextReset(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v319 != int64(0) {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+390)) = uint8(v328)
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+352)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v212)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+328)) = v332 + int64(1)
	goto L82
L91:
	;
	goto L75
L92:
	;
	if v348 == int32(0) {
		goto L74
	} else {
		goto L93
	}
L93:
	;
	goto L75
L94:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v212)+120))
	if int32(0) < v358 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L4
	} else {
		goto L344
	}
L96:
	;
	v370 = int32(0)
	goto L99
L97:
	;
	goto L98
L98:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v212)+124))
	if v670 <= int32(0) {
		goto L70
	} else {
		goto L125
	}
L99:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v212)+128))
	v389 = v386 + v370*int32(56)
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+47)))
	if v390 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	v397 = int32(_a_F_ExecWindowAgg_2)
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	if int32(101) <= v403 {
		goto L95
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v644 = v370 + int32(1)
	if v644 != v358 {
		v370 = v644
		goto L99
	} else {
		goto L124
	}
L104:
	;
	v407 = v389 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v389)+52))
	v410 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v409
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v389)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)) = uint8(v410)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v413
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+26)) = uint16(v403)
	if v410 < v403 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v421 = v403 & int32(7)
	v422 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v403) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v569 = v407
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = int32(0)
	v595 = v393 + v396<<(uint(int32(3))%32)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v599 = m.T0[v598].(func(*base.Module, int32) int64)(m, v216+int32(8))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L119
	}
L108:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v569 = v565
	goto L107
L109:
	;
	v432 = v422
	v438 = int32(0)
	goto L112
L110:
	;
	v486 = v422
	goto L111
L111:
	;
	v510 = v486
	v512 = v422
	goto L116
L112:
	;
	v453 = int32(8)
	v457 = v216 + v453 + v432<<(uint(int32(4))%32)
	v458 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+144)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+128)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+112)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+96)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+80)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457-int32(-64)))) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+48)) = uint8(v458)
	*(*uint8)(unsafe.Add(mBase, uint32(v457)+32)) = uint8(v458)
	v477 = v432 + v453
	v479 = v438 + v453
	if v479 != v403&int32(2147483640) {
		v432 = v477
		v438 = v479
		goto L112
	} else {
		goto L114
	}
L113:
	;
	if v421 == int32(0) {
		goto L108
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v486 = v477
	goto L111
L116:
	;
	v534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v510<<(uint(int32(4))%32)+v216)+40)) = uint8(v534)
	v539 = v512 + v534
	if v539 != v421 {
		v510 = v510 + v534
		v512 = v539
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L108
L118:
	;
	goto L117
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v595))) = v599
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v396+v394))) = uint8(v602)
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+46)))
	if (v602|v604)&int32(1) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v398
	goto L103
L121:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v212)+120))
	if v608 < int32(2) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v595)))
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v389)+44)))
	v614 = F_datumCopy(m, v611, int32(0), v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v595))) = v614
	goto L120
L124:
	;
	goto L100
L125:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v212)+408))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v212)+404))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v212)+200))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	F_update_frameheadpos(m, v212)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v212)+184))
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v212)+208))
	if v680 <= v679 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v679 != v680 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L4
	} else {
		goto L341
	}
L130:
	;
	v697 = int32(0)
	v700 = v697
	v705 = v697
	goto L135
L131:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v212)+228))
	if base.B2i32(v683&int32(1280) == int32(0))|v683&int32(_a_F_ExecWindowAgg_10) != 0 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	if v691 < v679 {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v693 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	if v691 < v693 {
		goto L71
	} else {
		goto L134
	}
L134:
	;
	goto L130
L135:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v726 = v723 + v700*int32(184)
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v212)+176))
	if v727 == int64(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	if v670 <= v748 {
		v1179 = v748
		goto L150
	} else {
		goto L151
	}
L137:
	;
	v751 = v700 + int32(1)
	if v751 != v670 {
		v700 = v751
		v705 = v748
		goto L135
	} else {
		goto L147
	}
L138:
	;
	v746 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+176)) = uint8(v746)
	v748 = v705
	goto L137
L139:
	;
	v742 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+176)) = uint8(v742)
	v748 = v705 + v742
	goto L137
L140:
	;
	v730 = *(*int64)(unsafe.Add(mBase, uint32(v212)+184))
	v731 = *(*int64)(unsafe.Add(mBase, uint32(v212)+208))
	if v730 != v731 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	if v733 == int32(0) {
		goto L139
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+229)))
	if v736&int32(896) != 0 {
		goto L139
	} else {
		goto L145
	}
L144:
	;
	goto L143
L145:
	;
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	if v730 < v739 {
		goto L138
	} else {
		goto L146
	}
L146:
	;
	goto L139
L147:
	;
	goto L136
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L4
	} else {
		goto L338
	}
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L4
	} else {
		goto L335
	}
L150:
	;
	v1197 = *(*int64)(unsafe.Add(mBase, uint32(v212)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+208)) = v1197
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v675)+16))
	if int32(0) <= v1199 {
		goto L219
	} else {
		goto L220
	}
L151:
	;
	v760 = v748
	goto L152
L152:
	;
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v212)+208))
	v779 = *(*int64)(unsafe.Add(mBase, uint32(v212)+184))
	if v779 <= v778 {
		v1179 = v760
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v1179 = v1139
	goto L150
L154:
	;
	v781 = F_window_gettupleslot(m, v675, v778, v673)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	if v781 == int32(0) {
		goto L148
	} else {
		goto L156
	}
L156:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	*(*int32)(unsafe.Add(mBase, uint32(v785)+12)) = v673
	v794 = v760
	v799 = int32(0)
	goto L157
L157:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v815 = v812 + v799*int32(184)
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+176)))
	if v816 != 0 {
		v1139 = v794
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+20))
	F_MemoryContextReset(m, v1161)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L4
	} else {
		goto L216
	}
L159:
	;
	v1158 = v799 + int32(1)
	if v1158 != v670 {
		v794 = v1139
		v799 = v1158
		goto L157
	} else {
		goto L215
	}
L160:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v212)+128))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v815)+140))
	v821 = v817 + v818*int32(56)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+8))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+12))
	v825 = int32(_a_F_ExecWindowAgg_2)
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v829
	if v824 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v823)+8))
	if v847 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v824)+24))
	v836 = m.T0[v835].(func(*base.Module, int32, int32, int32) int64)(m, v824, v828, v216+int32(7))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+7)))
	if base.B2i32(v838 == int32(0))&base.B2i32(v836 != int64(0)) != 0 {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v826
	v1139 = v794
	goto L159
L165:
	;
	v923 = int32(1)
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+50)))
	if base.B2i32(v924 != v923)|base.B2i32(v822 <= int32(0)) != 0 {
		goto L172
	} else {
		goto L173
	}
L166:
	;
	v851 = int32(0)
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v852 <= v851 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v856 = int32(1)
	v858 = v851
	goto L168
L168:
	;
	v881 = v228 + v856<<(uint(int32(4))%32)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v847)+12))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v882+v858<<(uint(int32(2))%32))))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v886)+24))
	v890 = m.T0[v889].(func(*base.Module, int32, int32, int32) int64)(m, v886, v828, v881+int32(8))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L4
	} else {
		goto L170
	}
L169:
	;
	goto L165
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v881))) = v890
	v893 = int32(1)
	v896 = v858 + v893
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v896 < v897 {
		v856 = v856 + v893
		v858 = v896
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+160)))
	if v989 == int32(1) {
		goto L149
	} else {
		goto L180
	}
L173:
	;
	v931 = v923
	goto L174
L174:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931<<(uint(int32(4))%32)+v216)+40)))
	if v957 != int32(1) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v826
	v1139 = v794
	goto L159
L176:
	;
	v961 = v931 + int32(1)
	if v961 <= v822 {
		v931 = v961
		goto L174
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	goto L172
L180:
	;
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v815)+168))
	if v992 == int64(1) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v826
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v815)+144))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v212)+372))
	if v997 != v998 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	v1031 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v815 + int32(40)
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v821)+40))
	v1039 = v822 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+26)) = uint16(v1039)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)) = uint8(v1031)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v1037
	v1044 = *(*int64)(unsafe.Add(mBase, uint32(v815)+152))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+40)) = uint8(v1031)
	*(*int64)(unsafe.Add(mBase, uint32(v216)+32)) = v1044
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v815)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = v1048
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v815)+40))
	v1053 = m.T0[v1052].(func(*base.Module, int32) int64)(m, v216+int32(8))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L4
	} else {
		goto L193
	}
L184:
	;
	F_MemoryContextReset(m, v997)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+112)))
	if v1002 == int32(1) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L186
L188:
	;
	v1023 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v815)+168)) = v1023
	*(*uint8)(unsafe.Add(mBase, uint32(v815)+160)) = uint8(v1022)
	*(*int64)(unsafe.Add(mBase, uint32(v815)+152)) = v1021
	v1027 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v815)+128)) = uint8(v1027)
	*(*int64)(unsafe.Add(mBase, uint32(v815)+120)) = v1023
	v1139 = v794
	goto L159
L189:
	;
	v1005 = *(*int64)(unsafe.Add(mBase, uint32(v815)+104))
	v1021 = v1005
	v1022 = int32(1)
	goto L188
L190:
	;
	goto L191
L191:
	;
	v1007 = int32(_a_F_ExecWindowAgg_2)
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v815)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1010
	v1012 = *(*int64)(unsafe.Add(mBase, uint32(v815)+104))
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+138)))
	v1014 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815)+134)))
	v1015 = F_datumCopy(m, v1012, v1013, v1014)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1008
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+112)))
	v1021 = v1015
	v1022 = v1019
	goto L188
L193:
	;
	v1055 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = v1055
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	if v1057 == v1055 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v815)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v815)+168)) = v1060 - int64(1)
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+138)))
	if v1064 != 0 {
		v1120 = v1053
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v826
	v1129 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v815)+176)) = uint8(v1129)
	v1139 = v794 + v1129
	goto L159
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v826
	*(*int64)(unsafe.Add(mBase, uint32(v815)+152)) = v1120
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v815)+160)) = uint8(v1125)
	v1139 = v794
	goto L159
L198:
	;
	v1065 = base.I32_wrap_i64(v1053)
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v815)+152))
	if v1065 == v1066 {
		v1120 = v1053
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v1068 = int32(0)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v815)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1070
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+134)))
	if v1072 != int32(_a_F_ExecWindowAgg_11) {
		v1091 = v1072
		v1092 = v1068
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+160)))
	if v1101 != 0 {
		v1120 = v1100
		goto L197
	} else {
		goto L208
	}
L201:
	;
	v1096 = F_datumCopy(m, v1053, v1092&int32(1), base.I32_extend16_s(v1091))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L4
	} else {
		goto L207
	}
L202:
	;
	v1075 = int32(_a_F_ExecWindowAgg_11)
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	if v1076 != int32(1) {
		v1091 = v1075
		v1092 = v1068
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065)+1)))
	if v1079 != int32(3) {
		v1091 = v1075
		v1092 = v1068
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1053))+2))
	goto L205
L205:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+8))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+16))
	v1087 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	if v1085 == v1087 {
		v1100 = v1053
		goto L200
	} else {
		goto L206
	}
L206:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+138)))
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+134)))
	v1091 = v1090
	v1092 = v1089
	goto L201
L207:
	;
	v1100 = v1096
	goto L200
L208:
	;
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v815)+152))
	v1103 = base.I32_wrap_i64(v1102)
	v1104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+134)))
	if v1104 != int32(_a_F_ExecWindowAgg_11) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_pfree(m, v1103)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L4
	} else {
		goto L214
	}
L210:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103))))
	if v1107 != int32(1) {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103)+1)))
	if v1110 != int32(3) {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	F_DeleteExpandedObject(m, v1102)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v1120 = v1100
	goto L197
L214:
	;
	v1120 = v1100
	goto L197
L215:
	;
	goto L158
L216:
	;
	v1164 = *(*int64)(unsafe.Add(mBase, uint32(v212)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+208)) = v1164 + int64(1)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v673)+8))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	m.T0[v1169].(func(*base.Module, int32))(m, v673)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	if v1139 < v670 {
		v760 = v1139
		goto L152
	} else {
		goto L218
	}
L218:
	;
	goto L153
L219:
	;
	F_WinSetMarkPosition(m, v675, v1197)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v1204 = int32(0)
	v1205 = base.B2i32(v1179 <= v1204)
	if v1205 == v1204 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L221
L223:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v212)+372))
	F_MemoryContextReset(m, v1208)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1215 = int32(0)
	goto L227
L226:
	;
	goto L225
L227:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v1239 = v1236 + v1215*int32(184)
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+176)))
	if v1240 == int32(1) {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v1289 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	if v1179 <= v1204 {
		goto L247
	} else {
		goto L248
	}
L229:
	;
	v1287 = v1215 + int32(1)
	if v1287 != v670 {
		v1215 = v1287
		goto L227
	} else {
		goto L246
	}
L230:
	;
	v1280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1239)+128)) = uint8(v1280)
	*(*int64)(unsafe.Add(mBase, uint32(v1239)+120)) = int64(0)
	goto L229
L231:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+144))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v212)+372))
	if v1243 != v1244 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+128)))
	if v1273 != 0 {
		goto L229
	} else {
		goto L243
	}
L234:
	;
	F_MemoryContextReset(m, v1243)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L4
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+112)))
	if v1248 == int32(1) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L236
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1239)+168)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1239)+160)) = uint8(v1268)
	*(*int64)(unsafe.Add(mBase, uint32(v1239)+152)) = v1267
	goto L230
L239:
	;
	v1251 = *(*int64)(unsafe.Add(mBase, uint32(v1239)+104))
	v1267 = v1251
	v1268 = int32(1)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1253 = int32(_a_F_ExecWindowAgg_2)
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1256
	v1258 = *(*int64)(unsafe.Add(mBase, uint32(v1239)+104))
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+138)))
	v1260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1239)+134)))
	v1261 = F_datumCopy(m, v1258, v1259, v1260)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1254
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+112)))
	v1267 = v1261
	v1268 = v1265
	goto L238
L243:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+137)))
	if v1274 != 0 {
		goto L230
	} else {
		goto L244
	}
L244:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+120))
	F_pfree(m, v1275)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	goto L230
L246:
	;
	goto L228
L247:
	;
	goto L251
L248:
	;
	v1290 = *(*int64)(unsafe.Add(mBase, uint32(v212)+184))
	if v1289 == v1290 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v212)+216)) = v1290
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v674)+8))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+12))
	m.T0[v1294].(func(*base.Module, int32))(m, v674)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	if v674 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	v1334 = F_row_is_in_frame(m, v675, v1332, v674, int32(0))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L4
	} else {
		goto L260
	}
L254:
	;
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+4)))
	if v1322&int32(2) == int32(0) {
		goto L253
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1327 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	v1328 = F_window_gettupleslot(m, v675, v1327, v674)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L4
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	if v1328 == int32(0) {
		goto L72
	} else {
		goto L259
	}
L259:
	;
	goto L253
L260:
	;
	if v1334 < int32(0) {
		goto L72
	} else {
		goto L261
	}
L261:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	if v1334 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+12)) = v674
	v1352 = int32(0)
	goto L265
L263:
	;
	v1763 = v1338
	goto L264
L264:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1763)+20))
	F_MemoryContextReset(m, v1764)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L4
	} else {
		goto L333
	}
L265:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v1368 = v1365 + v1352*int32(184)
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+176)))
	if v1369 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	v1763 = v1738
	goto L264
L267:
	;
	v1736 = v1352 + int32(1)
	if v1736 != v670 {
		v1352 = v1736
		goto L265
	} else {
		goto L332
	}
L268:
	;
	v1372 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	if v1372 < v1289 {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v212)+128))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+140))
	v1378 = v1374 + v1375*int32(56)
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+8))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1378)))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+12))
	v1382 = int32(_a_F_ExecWindowAgg_2)
	v1383 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v212)+380))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1386
	if v1381 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L270
L272:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+8))
	if v1404 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+24))
	v1393 = m.T0[v1392].(func(*base.Module, int32, int32, int32) int64)(m, v1381, v1385, v216+int32(7))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+7)))
	if base.B2i32(v1395 == int32(0))&base.B2i32(v1393 != int64(0)) != 0 {
		goto L272
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1383
	goto L267
L276:
	;
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+22)))
	if v1480 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L277:
	;
	v1408 = int32(0)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	if v1409 <= v1408 {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1413 = int32(1)
	v1415 = v1408
	goto L279
L279:
	;
	v1438 = v228 + v1413<<(uint(int32(4))%32)
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+12))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1439+v1415<<(uint(int32(2))%32))))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+24))
	v1447 = m.T0[v1446].(func(*base.Module, int32, int32, int32) int64)(m, v1443, v1385, v1438+int32(8))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L4
	} else {
		goto L281
	}
L280:
	;
	goto L276
L281:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1438))) = v1447
	v1450 = int32(1)
	v1453 = v1415 + v1450
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	if v1453 < v1454 {
		v1413 = v1413 + v1450
		v1415 = v1453
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L4
	} else {
		goto L328
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1383
	goto L267
L285:
	;
	v1597 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v1597
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v1368 + int32(12)
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+40))
	v1605 = v1379 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+26)) = uint16(v1605)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)) = uint8(v1597)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v1603
	v1610 = *(*int64)(unsafe.Add(mBase, uint32(v1368)+152))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+40)) = uint8(v1596)
	*(*int64)(unsafe.Add(mBase, uint32(v216)+32)) = v1610
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = v1613
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+12))
	v1618 = m.T0[v1617].(func(*base.Module, int32) int64)(m, v216+int32(8))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L4
	} else {
		goto L304
	}
L286:
	;
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+160)))
	v1596 = v1483
	goto L285
L287:
	;
	goto L288
L288:
	;
	if v1379 <= int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+160)))
	v1547 = *(*int64)(unsafe.Add(mBase, uint32(v1368)+168))
	if v1547 == int64(0) {
		goto L298
	} else {
		goto L299
	}
L290:
	;
	v1488 = int32(1)
	goto L291
L291:
	;
	v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1488<<(uint(int32(4))%32)+v216)+40)))
	if v1514 != int32(1) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1383
	goto L267
L293:
	;
	v1518 = v1488 + int32(1)
	if v1518 <= v1379 {
		v1488 = v1518
		goto L291
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	goto L292
L296:
	;
	goto L289
L297:
	;
	v1596 = int32(0)
	goto L285
L298:
	;
	if v1546&int32(1) == int32(0) {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	if v1546&int32(1) != 0 {
		goto L284
	} else {
		goto L303
	}
L301:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1555
	v1557 = *(*int64)(unsafe.Add(mBase, uint32(v216)+48))
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+138)))
	v1559 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1368)+134)))
	v1560 = F_datumCopy(m, v1557, v1558, v1559)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1368)+168)) = int64(1)
	v1564 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+160)) = uint8(v1564)
	*(*int64)(unsafe.Add(mBase, uint32(v1368)+152)) = v1560
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1383
	goto L267
L303:
	;
	goto L297
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = int32(0)
	v1622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	if v1622 == int32(1) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+4))
	if v1625 != 0 {
		goto L283
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v1626 = *(*int64)(unsafe.Add(mBase, uint32(v1368)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v1368)+168)) = v1626 + int64(1)
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+138)))
	if v1630 != 0 {
		v1686 = v1618
		goto L309
	} else {
		goto L310
	}
L308:
	;
	goto L307
L309:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1383
	*(*int64)(unsafe.Add(mBase, uint32(v1368)+152)) = v1686
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+160)) = uint8(v1691)
	goto L267
L310:
	;
	v1631 = base.I32_wrap_i64(v1618)
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+152))
	if v1631 == v1632 {
		v1686 = v1618
		goto L309
	} else {
		goto L311
	}
L311:
	;
	if v1622 != 0 {
		v1666 = v1618
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+160)))
	if v1667 != 0 {
		v1686 = v1666
		goto L309
	} else {
		goto L321
	}
L313:
	;
	v1634 = int32(0)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1636
	v1638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1368)+134)))
	if v1638 != int32(_a_F_ExecWindowAgg_11) {
		v1657 = v1638
		v1658 = v1634
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1662 = F_datumCopy(m, v1618, v1658&int32(1), base.I32_extend16_s(v1657))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L4
	} else {
		goto L320
	}
L315:
	;
	v1641 = int32(_a_F_ExecWindowAgg_11)
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1631))))
	if v1642 != int32(1) {
		v1657 = v1641
		v1658 = v1634
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1631)+1)))
	if v1645 != int32(3) {
		v1657 = v1641
		v1658 = v1634
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1618))+2))
	goto L318
L318:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+8))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+16))
	v1653 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	if v1651 == v1653 {
		v1666 = v1618
		goto L312
	} else {
		goto L319
	}
L319:
	;
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+138)))
	v1656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1368)+134)))
	v1657 = v1656
	v1658 = v1655
	goto L314
L320:
	;
	v1666 = v1662
	goto L312
L321:
	;
	v1668 = *(*int64)(unsafe.Add(mBase, uint32(v1368)+152))
	v1669 = base.I32_wrap_i64(v1668)
	v1670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1368)+134)))
	if v1670 != int32(_a_F_ExecWindowAgg_11) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	F_pfree(m, v1669)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L4
	} else {
		goto L327
	}
L323:
	;
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669))))
	if v1673 != int32(1) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+1)))
	if v1676 != int32(3) {
		goto L322
	} else {
		goto L325
	}
L325:
	;
	F_DeleteExpandedObject(m, v1668)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L4
	} else {
		goto L326
	}
L326:
	;
	v1686 = v1666
	goto L309
L327:
	;
	v1686 = v1666
	goto L309
L328:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L329
	}
L329:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_12), int32(0))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L4
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(402), int32(_a_F_ExecWindowAgg_13))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L4
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
	goto L266
L333:
	;
	v1767 = *(*int64)(unsafe.Add(mBase, uint32(v212)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+216)) = v1767 + int64(1)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v674)+8))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+12))
	m.T0[v1772].(func(*base.Module, int32))(m, v674)
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L4
	} else {
		goto L334
	}
L334:
	;
	goto L251
L335:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_14), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(534), int32(_a_F_ExecWindowAgg_15))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L4
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_16), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L4
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(862), int32(_a_F_ExecWindowAgg_17))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_18), int32(0))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(784), int32(_a_F_ExecWindowAgg_17))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	v1821 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v1821
	F_errmsg_plural(m, int32(_a_F_ExecWindowAgg_19), int32(_a_F_ExecWindowAgg_20), v1821, v216)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(1100), int32(_a_F_ExecWindowAgg_21))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_22), int32(0))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2406), int32(_a_F_ExecWindowAgg_23))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L351:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_22), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L4
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2392), int32(_a_F_ExecWindowAgg_23))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L4
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v1888 = v1885 + v1869*int32(184)
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+140))
	v1890 = int32(_a_F_ExecWindowAgg_2)
	v1891 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v212)+128))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v676)+32))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v676)+36))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1896)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1897
	v1899 = v1889 + v1894
	v1902 = v1893 + v1889<<(uint(int32(3))%32)
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+8))
	if v1903 != 0 {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	goto L70
L356:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1891
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+137)))
	if v2188 != 0 {
		goto L403
	} else {
		goto L404
	}
L357:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+96))
	v1905 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v1905
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v1888 + int32(68)
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1892+v1889*int32(56))+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+26)) = uint16(v1904)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)) = uint8(v1905)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v1914
	v1919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+160)))
	if v1919 == v1905 {
		goto L362
	} else {
		goto L363
	}
L358:
	;
	goto L359
L359:
	;
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+160)))
	if v2137 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L360:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+40)) = uint8(v1941)
	*(*int64)(unsafe.Add(mBase, uint32(v216)+32)) = v1942
	if v1904 < int32(2) {
		v2075 = v1941
		goto L370
	} else {
		goto L371
	}
L361:
	;
	v1926 = *(*int64)(unsafe.Add(mBase, uint32(v1888)+152))
	v1928 = base.I32_wrap_i64(v1926)
	v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1928))))
	if v1929 != int32(1) {
		v1939 = v1926
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1888)+134)))
	if v1922 == int32(_a_F_ExecWindowAgg_11) {
		goto L361
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1925 = *(*int64)(unsafe.Add(mBase, uint32(v1888)+152))
	v1941 = v1919
	v1942 = v1925
	goto L360
L365:
	;
	goto L364
L366:
	;
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+160)))
	v1941 = v1940
	v1942 = v1939
	goto L360
L367:
	;
	goto L366
L368:
	;
	v1932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1928)+1)))
	if v1932 != int32(3) {
		v1939 = v1926
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+2))
	v1939 = base.I64_extend_i32_u(v1935 + int32(18))
	goto L367
L370:
	;
	v2094 = int32(1)
	v2096 = int32(0)
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+78)))
	if base.B2i32(v2075&v2094 == v2096)|base.B2i32(v2098 != v2094) == v2096 {
		goto L381
	} else {
		goto L382
	}
L371:
	;
	v1947 = int32(1)
	v1948 = v1904 - v1947
	v1949 = int32(3)
	v1950 = v1948 & v1949
	if base.Ui32(v1904-int32(2)) < base.Ui32(v1949) {
		v2011 = v1947
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v2034 = int32(0)
	v2036 = v2011
	goto L378
L373:
	;
	v1962 = v1947
	v1964 = int32(0)
	goto L374
L374:
	;
	v1983 = int32(4)
	v1985 = v228 + v1962<<(uint(v1983)%32)
	v1986 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1985)+56)) = uint8(v1986)
	v1988 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1985)+48)) = v1988
	*(*uint8)(unsafe.Add(mBase, uint32(v1985)+40)) = uint8(v1986)
	*(*int64)(unsafe.Add(mBase, uint32(v1985)+32)) = v1988
	*(*uint8)(unsafe.Add(mBase, uint32(v1985)+24)) = uint8(v1986)
	*(*int64)(unsafe.Add(mBase, uint32(v1985)+16)) = v1988
	*(*uint8)(unsafe.Add(mBase, uint32(v1985)+8)) = uint8(v1986)
	*(*int64)(unsafe.Add(mBase, uint32(v1985))) = v1988
	v2003 = v1962 + v1983
	v2005 = v1964 + v1983
	if v2005 != v1948&int32(-4) {
		v1962 = v2003
		v1964 = v2005
		goto L374
	} else {
		goto L376
	}
L375:
	;
	if v1950 != 0 {
		v2011 = v2003
		goto L372
	} else {
		goto L377
	}
L376:
	;
	goto L375
L377:
	;
	v2075 = int32(1)
	goto L370
L378:
	;
	v2057 = int32(1)
	v2060 = v228 + v2036<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2060)+8)) = uint8(v2057)
	*(*int64)(unsafe.Add(mBase, uint32(v2060))) = int64(0)
	v2068 = v2034 + v2057
	if v2068 != v1950 {
		v2034 = v2068
		v2036 = v2036 + v2057
		goto L378
	} else {
		goto L380
	}
L379:
	;
	v2075 = v2057
	goto L370
L380:
	;
	goto L379
L381:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1902))) = int64(0)
	v2106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1899))) = uint8(v2106)
	goto L356
L382:
	;
	goto L383
L383:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = v2108
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+68))
	v2113 = m.T0[v2112].(func(*base.Module, int32) int64)(m, v216+int32(8))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+376)) = int32(0)
	v2117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1899))) = uint8(v2117)
	if v2117 != 0 {
		v2135 = v2113
		goto L385
	} else {
		goto L386
	}
L385:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1902))) = v2135
	goto L356
L386:
	;
	v2119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1888)+132)))
	if v2119 != int32(_a_F_ExecWindowAgg_11) {
		v2135 = v2113
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v2123 = base.I32_wrap_i64(v2113)
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2123))))
	if v2124 != int32(1) {
		v2134 = v2113
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v2135 = v2134
	goto L385
L389:
	;
	goto L388
L390:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2123)+1)))
	if v2127 != int32(3) {
		v2134 = v2113
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2123)+2))
	v2134 = base.I64_extend_i32_u(v2130 + int32(18))
	goto L389
L392:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1902))) = v2158
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+160)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1899))) = uint8(v2160)
	goto L356
L393:
	;
	v2144 = *(*int64)(unsafe.Add(mBase, uint32(v1888)+152))
	v2146 = base.I32_wrap_i64(v2144)
	v2147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2146))))
	if v2147 != int32(1) {
		v2157 = v2144
		goto L399
	} else {
		goto L400
	}
L394:
	;
	v2140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1888)+134)))
	if v2140 == int32(_a_F_ExecWindowAgg_11) {
		goto L393
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v2143 = *(*int64)(unsafe.Add(mBase, uint32(v1888)+152))
	v2158 = v2143
	goto L392
L397:
	;
	goto L396
L398:
	;
	v2158 = v2157
	goto L392
L399:
	;
	goto L398
L400:
	;
	v2150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2146)+1)))
	if v2150 != int32(3) {
		v2157 = v2144
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+2))
	v2157 = base.I64_extend_i32_u(v2153 + int32(18))
	goto L399
L402:
	;
	v2203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1899))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1888)+128)) = uint8(v2203)
	v2206 = v1869 + int32(1)
	if v2206 != v670 {
		v1869 = v2206
		goto L354
	} else {
		goto L407
	}
L403:
	;
	v2201 = *(*int64)(unsafe.Add(mBase, uint32(v1902)))
	*(*int64)(unsafe.Add(mBase, uint32(v1888)+120)) = v2201
	goto L402
L404:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1899))))
	if v2189 != 0 {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v1888)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2191
	v2193 = *(*int64)(unsafe.Add(mBase, uint32(v1902)))
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+137)))
	v2195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1888)+132)))
	v2196 = F_datumCopy(m, v2193, v2194, v2195)
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L4
	} else {
		goto L406
	}
L406:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1888)+120)) = v2196
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1891
	goto L402
L407:
	;
	goto L355
L408:
	;
	v2217 = v2208
	v2221 = int32(0)
	goto L411
L409:
	;
	v2276 = v2208
	goto L410
L410:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v676)+32))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v2303 = v2300 + v2276*int32(184)
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2303)+140))
	v2308 = *(*int64)(unsafe.Add(mBase, uint32(v2303)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v2299+v2304<<(uint(int32(3))%32)))) = v2308
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v676)+36))
	v2312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2304+v2310))) = uint8(v2312)
	goto L70
L411:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v676)+32))
	v2242 = v2217 * int32(184)
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v2244 = v2242 + v2243
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+140))
	v2246 = int32(3)
	v2249 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v2240+v2245<<(uint(v2246)%32)))) = v2249
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v676)+36))
	v2253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2244)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2245+v2251))) = uint8(v2253)
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v676)+32))
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v212)+132))
	v2257 = v2256 + v2242
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2257)+324))
	v2262 = *(*int64)(unsafe.Add(mBase, uint32(v2257)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v2255+v2258<<(uint(v2246)%32)))) = v2262
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v676)+36))
	v2266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2257)+312)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2258+v2264))) = uint8(v2266)
	v2268 = int32(2)
	v2269 = v2217 + v2268
	v2271 = v2221 + v2268
	if v2271 != v670&int32(2147483646) {
		v2217 = v2269
		v2221 = v2271
		goto L411
	} else {
		goto L413
	}
L412:
	;
	if v670&int32(1) == int32(0) {
		goto L70
	} else {
		goto L414
	}
L413:
	;
	goto L412
L414:
	;
	v2276 = v2269
	goto L410
L415:
	;
	F_update_frameheadpos(m, v212)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L4
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v212)+156))
	if int32(0) <= v2343 {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	goto L417
L419:
	;
	F_update_frametailpos(m, v212)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L4
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v212)+160))
	if int32(0) <= v2348 {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	goto L421
L423:
	;
	F_update_grouptailpos(m, v212)
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L4
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v212)+144))
	F_tuplestore_trim(m, v2353)
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L4
	} else {
		goto L427
	}
L426:
	;
	goto L425
L427:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v212)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v266)+12)) = v2356
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v212)+68))
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2358)+80))
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2358)+24))
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2360)+8))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2361)+12))
	m.T0[v2362].(func(*base.Module, int32))(m, v2360)
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L4
	} else {
		goto L428
	}
L428:
	;
	v2365 = int32(_a_F_ExecWindowAgg_2)
	v2366 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2368
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2358)+32))
	v2374 = m.T0[v2373].(func(*base.Module, int32, int32, int32) int64)(m, v2358+int32(8), v2359, int32(0))
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L4
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2366
	v2378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2360)+4)))
	v2380 = v2378 & int32(_a_F_ExecWindowAgg_24)
	*(*uint16)(unsafe.Add(mBase, uint32(v2360)+4)) = uint16(v2380)
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2360)+12))
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2382)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2360)+6)) = uint16(v2383)
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v212)+224))
	if v2385 != int32(1) {
		goto L58
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v2360
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v212)+320))
	if v2389 == int32(0) {
		goto L59
	} else {
		goto L431
	}
L431:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2393
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2389)+24))
	v2398 = m.T0[v2397].(func(*base.Module, int32, int32, int32) int64)(m, v2389, v266, v216+int32(8))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L4
	} else {
		goto L432
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2366
	if v2398 != int64(0) {
		goto L59
	} else {
		goto L433
	}
L433:
	;
	v2404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+318)))
	if v2404 != int32(1) {
		v2592 = v216
		goto L61
	} else {
		goto L434
	}
L434:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v212)+120))
	if v2407 <= int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+319)))
	if v2583 != int32(1) {
		goto L60
	} else {
		goto L447
	}
L436:
	;
	v2411 = v2407 & int32(3)
	v2412 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2407) {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2420 = v2412
	v2428 = int32(0)
	goto L440
L438:
	;
	v2497 = v2412
	goto L439
L439:
	;
	v2521 = v2497
	v2525 = v2412
	goto L444
L440:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	v2444 = int32(3)
	v2447 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2443+v2420<<(uint(v2444)%32)))) = v2447
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v2451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2449+v2420))) = uint8(v2451)
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	v2455 = v2420 | v2451
	*(*int64)(unsafe.Add(mBase, uint32(v2453+v2455<<(uint(v2444)%32)))) = v2447
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2461+v2455))) = uint8(v2451)
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	v2467 = v2420 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v2465+v2467<<(uint(v2444)%32)))) = v2447
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2473+v2467))) = uint8(v2451)
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	v2479 = v2420 | v2444
	*(*int64)(unsafe.Add(mBase, uint32(v2477+v2479<<(uint(v2444)%32)))) = v2447
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2485+v2479))) = uint8(v2451)
	v2489 = int32(4)
	v2490 = v2420 + v2489
	v2492 = v2428 + v2489
	if v2492 != v2407&int32(2147483644) {
		v2420 = v2490
		v2428 = v2492
		goto L440
	} else {
		goto L442
	}
L441:
	;
	if v2411 == int32(0) {
		goto L435
	} else {
		goto L443
	}
L442:
	;
	goto L441
L443:
	;
	v2497 = v2490
	goto L439
L444:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2544+v2521<<(uint(int32(3))%32)))) = int64(0)
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v2552 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2550+v2521))) = uint8(v2552)
	v2557 = v2525 + v2552
	if v2557 != v2411 {
		v2521 = v2521 + v2552
		v2525 = v2557
		goto L444
	} else {
		goto L446
	}
L445:
	;
	goto L435
L446:
	;
	goto L445
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+224)) = int32(3)
	goto L50
L448:
	;
	v2644 = int32(_a_F_ExecWindowAgg_2)
	v2645 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2647
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2641)+24))
	v2652 = m.T0[v2651].(func(*base.Module, int32, int32, int32) int64)(m, v2641, v266, v216+int32(8))
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2645
	if v2652 != int64(0) {
		v2669 = v2360
		v2670 = v216
		goto L6
	} else {
		goto L450
	}
L450:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	if v2658 == int32(0) {
		goto L50
	} else {
		goto L451
	}
L451:
	;
	v2661 = *(*float64)(unsafe.Add(mBase, uint32(v2658)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v2658)+424)) = base.F64_add(v2661, float64(1))
	goto L50
L452:
	;
	goto L51
}
func F_window_lag_with_offset_and_default(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_leadlag_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_window_lead_with_offset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v10, int32(1), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v18 = v8 + int32(15)
		v19 = F_WinGetFuncArgCurrent(m, v10, int32(1), v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if v21 == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = int32(0)
				if v25 == v27 {
					v73 = v27
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
					if v31 == int32(0) {
						v73 = v27
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						v36 = v34 - int32(11)
						if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v36)%32))&int32(1) == int32(0))|int32(0) != 0 {
							v73 = v27
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v36<<(uint(int32(2))%32))+uint32(_c_F_window_lead_with_offset[0])))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v31+v51)))
							if v53 == int32(0) {
								v73 = v27
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
								if v56 <= int32(1) {
									v73 = v27
								} else {
									v58 = int32(1)
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(4))))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
									switch v64 - int32(7) {
									case 0:
										v73 = v58
									case 1:
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
										if v67 == int32(0) {
											v73 = v58
										} else {
											v73 = int32(0)
										}
									default:
										v73 = int32(0)
									}
								}
							}
						}
					}
				}
				v76 = F_WinGetFuncArgInPartition(m, v10, base.I32_wrap_i64(v19), v73, v18, v8+int32(14))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int64(0)
				} else {
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v78 != int32(1) {
						v85 = v76
					} else {
						v82 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
						v85 = int64(0)
					}
					m.G0 = v8 + int32(16)
					return v85
				}
			} else {
				v82 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
				v85 = int64(0)
				m.G0 = v8 + int32(16)
				return v85
			}
		}
	}
}
func F_window_lead_with_offset_and_default(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_leadlag_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
