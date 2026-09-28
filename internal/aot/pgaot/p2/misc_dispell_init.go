package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dispell_init(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v160 int32
	_ = v160
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v273 int32
	_ = v273
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v782 int32
	_ = v782
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v857 int32
	_ = v857
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1013 int32
	_ = v1013
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1065 int32
	_ = v1065
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1331 int32
	_ = v1331
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1399 int32
	_ = v1399
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1486 int32
	_ = v1486
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1522 int32
	_ = v1522
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1553 int32
	_ = v1553
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
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
	var v1724 int32
	_ = v1724
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2196 int32
	_ = v2196
	var v2219 int32
	_ = v2219
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2345 int32
	_ = v2345
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2357 int32
	_ = v2357
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2421 int32
	_ = v2421
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2510 int32
	_ = v2510
	var v2525 int32
	_ = v2525
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2568 int32
	_ = v2568
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2578 int32
	_ = v2578
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2595 int32
	_ = v2595
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2620 int32
	_ = v2620
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2638 int32
	_ = v2638
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2651 int32
	_ = v2651
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2668 int32
	_ = v2668
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2726 int32
	_ = v2726
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2743 int32
	_ = v2743
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2770 int32
	_ = v2770
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2790 int32
	_ = v2790
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2801 int32
	_ = v2801
	var v2808 int32
	_ = v2808
	var v2811 int32
	_ = v2811
	var v2818 int32
	_ = v2818
	var v2823 int32
	_ = v2823
	var v2826 int32
	_ = v2826
	var v2829 int32
	_ = v2829
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2845 int32
	_ = v2845
	var v2849 int32
	_ = v2849
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2891 int32
	_ = v2891
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2906 int32
	_ = v2906
	var v2913 int32
	_ = v2913
	var v2916 int32
	_ = v2916
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2969 int32
	_ = v2969
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3007 int32
	_ = v3007
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3028 int32
	_ = v3028
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3042 int32
	_ = v3042
	var v3045 int32
	_ = v3045
	var v3051 int32
	_ = v3051
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3096 int32
	_ = v3096
	var v3100 int32
	_ = v3100
	var v3106 int32
	_ = v3106
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3148 int32
	_ = v3148
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3175 int32
	_ = v3175
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3196 int32
	_ = v3196
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3213 int32
	_ = v3213
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3228 int32
	_ = v3228
	var v3240 int32
	_ = v3240
	var v3245 int32
	_ = v3245
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3251 int32
	_ = v3251
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3271 int64
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3276 int32
	_ = v3276
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3286 int32
	_ = v3286
	var v3290 int32
	_ = v3290
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3303 int32
	_ = v3303
	var v3308 int32
	_ = v3308
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3357 int32
	_ = v3357
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3367 int32
	_ = v3367
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3375 int32
	_ = v3375
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3389 int32
	_ = v3389
	var v3398 int32
	_ = v3398
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3418 int32
	_ = v3418
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3424 int32
	_ = v3424
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3471 int32
	_ = v3471
	var v3476 int32
	_ = v3476
	var v3480 int32
	_ = v3480
	var v3483 int32
	_ = v3483
	var v3489 int32
	_ = v3489
	var v3494 int32
	_ = v3494
	var v3498 int32
	_ = v3498
	var v3501 int32
	_ = v3501
	var v3507 int32
	_ = v3507
	var v3512 int32
	_ = v3512
	var v3516 int32
	_ = v3516
	var v3519 int32
	_ = v3519
	var v3525 int32
	_ = v3525
	var v3530 int32
	_ = v3530
	var v3534 int32
	_ = v3534
	var v3540 int32
	_ = v3540
	var v3545 int32
	_ = v3545
	var v3549 int32
	_ = v3549
	var v3552 int32
	_ = v3552
	var v3558 int32
	_ = v3558
	var v3563 int32
	_ = v3563
	var v3567 int32
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3574 int32
	_ = v3574
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3583 int32
	_ = v3583
	var v3586 int32
	_ = v3586
	var v3589 int32
	_ = v3589
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3600 int32
	_ = v3600
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3630 int32
	_ = v3630
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3649 int32
	_ = v3649
	var v3652 int32
	_ = v3652
	var v3656 int32
	_ = v3656
	var v3661 int32
	_ = v3661
	var v3665 int32
	_ = v3665
	var v3668 int32
	_ = v3668
	var v3672 int32
	_ = v3672
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3684 int32
	_ = v3684
	var v3688 int32
	_ = v3688
	var v3693 int32
	_ = v3693
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3705 int32
	_ = v3705
	var v3710 int32
	_ = v3710
	var v3713 int32
	_ = v3713
	var v3726 int32
	_ = v3726
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3754 int32
	_ = v3754
	var v3758 int32
	_ = v3758
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3789 int32
	_ = v3789
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3797 int32
	_ = v3797
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3805 int64
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3808 int32
	_ = v3808
	var v3810 int32
	_ = v3810
	var v3811 int32
	_ = v3811
	var v3814 int32
	_ = v3814
	var v3819 int32
	_ = v3819
	var v3821 int32
	_ = v3821
	var v3838 int32
	_ = v3838
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3844 int32
	_ = v3844
	var v3847 int32
	_ = v3847
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3858 int32
	_ = v3858
	var v3859 int32
	_ = v3859
	var v3862 int32
	_ = v3862
	var v3865 int32
	_ = v3865
	var v3867 int32
	_ = v3867
	var v3871 int32
	_ = v3871
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3905 int32
	_ = v3905
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3912 int32
	_ = v3912
	var v3915 int32
	_ = v3915
	var v3916 int32
	_ = v3916
	var v3919 int32
	_ = v3919
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3931 int32
	_ = v3931
	var v3933 int32
	_ = v3933
	var v3939 int32
	_ = v3939
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3966 int32
	_ = v3966
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
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3983 int32
	_ = v3983
	var v3984 int32
	_ = v3984
	var v3986 int32
	_ = v3986
	var v3987 int32
	_ = v3987
	var v3988 int32
	_ = v3988
	var v3989 int32
	_ = v3989
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3996 int32
	_ = v3996
	var v4004 int32
	_ = v4004
	var v4008 int32
	_ = v4008
	var v4010 int32
	_ = v4010
	var v4011 int32
	_ = v4011
	var v4015 int32
	_ = v4015
	var v4016 int32
	_ = v4016
	var v4018 int32
	_ = v4018
	var v4022 int32
	_ = v4022
	var v4024 int32
	_ = v4024
	var v4026 int32
	_ = v4026
	var v4029 int32
	_ = v4029
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4046 int32
	_ = v4046
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4060 int32
	_ = v4060
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4065 int32
	_ = v4065
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4076 int32
	_ = v4076
	var v4077 int32
	_ = v4077
	var v4080 int32
	_ = v4080
	var v4081 int32
	_ = v4081
	var v4084 int32
	_ = v4084
	var v4086 int32
	_ = v4086
	var v4091 int32
	_ = v4091
	var v4093 int32
	_ = v4093
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4121 int32
	_ = v4121
	var v4122 int32
	_ = v4122
	var v4123 int32
	_ = v4123
	var v4127 int32
	_ = v4127
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4140 int32
	_ = v4140
	var v4141 int32
	_ = v4141
	var v4144 int32
	_ = v4144
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4164 int32
	_ = v4164
	var v4165 int32
	_ = v4165
	var v4167 int32
	_ = v4167
	var v4168 int32
	_ = v4168
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4172 int32
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4179 int32
	_ = v4179
	var v4186 int32
	_ = v4186
	var v4190 int32
	_ = v4190
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4204 int32
	_ = v4204
	var v4206 int32
	_ = v4206
	var v4208 int32
	_ = v4208
	var v4211 int32
	_ = v4211
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4218 int32
	_ = v4218
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4223 int32
	_ = v4223
	var v4225 int32
	_ = v4225
	var v4228 int32
	_ = v4228
	var v4233 int32
	_ = v4233
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4242 int32
	_ = v4242
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4255 int32
	_ = v4255
	var v4257 int32
	_ = v4257
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4266 int32
	_ = v4266
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4273 int32
	_ = v4273
	var v4276 int32
	_ = v4276
	var v4279 int32
	_ = v4279
	var v4280 int32
	_ = v4280
	var v4282 int32
	_ = v4282
	var v4311 int32
	_ = v4311
	var v4338 int32
	_ = v4338
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4355 int32
	_ = v4355
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4363 int32
	_ = v4363
	var v4364 int32
	_ = v4364
	var v4368 int32
	_ = v4368
	var v4373 int32
	_ = v4373
	var v4377 int32
	_ = v4377
	var v4380 int32
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4386 int32
	_ = v4386
	var v4392 int32
	_ = v4392
	var v4397 int32
	_ = v4397
	var v4401 int32
	_ = v4401
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4416 int32
	_ = v4416
	var v4421 int32
	_ = v4421
	var v4423 int32
	_ = v4423
	var v4425 int32
	_ = v4425
	var v4427 int32
	_ = v4427
	var v4431 int32
	_ = v4431
	var v4435 int32
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4440 int32
	_ = v4440
	var v4441 int32
	_ = v4441
	var v4445 int32
	_ = v4445
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4461 int32
	_ = v4461
	var v4474 int32
	_ = v4474
	var v4475 int32
	_ = v4475
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4484 int32
	_ = v4484
	var v4491 int32
	_ = v4491
	var v4494 int32
	_ = v4494
	var v4501 int32
	_ = v4501
	var v4523 int32
	_ = v4523
	var v4524 int32
	_ = v4524
	var v4528 int32
	_ = v4528
	var v4557 int32
	_ = v4557
	var v4558 int32
	_ = v4558
	var v4564 int32
	_ = v4564
	var v4567 int32
	_ = v4567
	var v4570 int32
	_ = v4570
	var v4573 int32
	_ = v4573
	var v4574 int32
	_ = v4574
	var v4577 int32
	_ = v4577
	var v4578 int32
	_ = v4578
	var v4581 int32
	_ = v4581
	var v4588 int32
	_ = v4588
	var v4589 int32
	_ = v4589
	var v4618 int32
	_ = v4618
	var v4620 int32
	_ = v4620
	var v4621 int32
	_ = v4621
	var v4625 int32
	_ = v4625
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4631 int32
	_ = v4631
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4634 int32
	_ = v4634
	var v4636 int32
	_ = v4636
	var v4639 int32
	_ = v4639
	var v4643 int32
	_ = v4643
	var v4644 int32
	_ = v4644
	var v4650 int32
	_ = v4650
	var v4655 int32
	_ = v4655
	var v4656 int32
	_ = v4656
	var v4678 int32
	_ = v4678
	var v4680 int32
	_ = v4680
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4685 int32
	_ = v4685
	var v4687 int32
	_ = v4687
	var v4694 int32
	_ = v4694
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4751 int32
	_ = v4751
	var v4753 int32
	_ = v4753
	var v4766 int32
	_ = v4766
	var v4789 int32
	_ = v4789
	var v4791 int32
	_ = v4791
	var v4792 int32
	_ = v4792
	var v4797 int32
	_ = v4797
	var v4798 int32
	_ = v4798
	var v4823 int32
	_ = v4823
	var v4827 int32
	_ = v4827
	var v4828 int32
	_ = v4828
	var v4830 int32
	_ = v4830
	var v4833 int32
	_ = v4833
	var v4834 int32
	_ = v4834
	var v4836 int32
	_ = v4836
	var v4839 int32
	_ = v4839
	var v4840 int32
	_ = v4840
	var v4844 int32
	_ = v4844
	var v4847 int32
	_ = v4847
	var v4878 int32
	_ = v4878
	var v4880 int32
	_ = v4880
	var v4881 int32
	_ = v4881
	var v4897 int32
	_ = v4897
	var v4900 int32
	_ = v4900
	var v4904 int32
	_ = v4904
	var v4909 int32
	_ = v4909
	var v4940 int32
	_ = v4940
	var v4943 int32
	_ = v4943
	var v4947 int32
	_ = v4947
	var v4952 int32
	_ = v4952
	v2 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = F_palloc0(m, int32(96))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v39 = v34 + int32(8)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[0]))
	v46 = F_AllocSetContextCreateInternal(m, v41, int32(_a_F_dispell_init_0), int32(0), int32(_a_F_dispell_init_1), int32(_a_F_dispell_init_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v46
	if v32 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4940 = m.ExcPending
	if v4940 != 0 {
		goto L1
	} else {
		goto L1157
	}
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v51 <= int32(0) {
		v3713 = v39
		v3726 = v2
		v3734 = v2
		v3735 = v30
		v3737 = v34
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v3738 = int32(0)
	if base.B2i32(v3726 == v3738)|base.B2i32(v3734 == v3738) == v3738 {
		goto L957
	} else {
		goto L958
	}
L7:
	;
	v55 = int32(0)
	v57 = v39
	v70 = v2
	v73 = v32
	v78 = v2
	v79 = v30
	v80 = v2
	v81 = v34
	goto L11
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L953
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3681 = m.ExcPending
	if v3681 != 0 {
		goto L1
	} else {
		goto L949
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3665 = m.ExcPending
	if v3665 != 0 {
		goto L1
	} else {
		goto L945
	}
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v55<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v88 = int32(_a_F_dispell_init_3)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dispell_init[1])))
	if base.B2i32(v91 == int32(0))|base.B2i32(v91 != v94) != 0 {
		v112 = v91
		v113 = v94
		goto L16
	} else {
		goto L17
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L1
	} else {
		goto L941
	}
L13:
	;
	goto L12
L14:
	;
	v3643 = v55 + int32(1)
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v3643 < v3644 {
		v55 = v3643
		v70 = v3630
		v78 = v3638
		v80 = v3640
		goto L11
	} else {
		goto L940
	}
L15:
	;
	if v112-v113 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v97 = v87
	v98 = v88
	goto L18
L18:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v102
		v113 = v101
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v112 = v102
	v113 = v101
	goto L16
L20:
	;
	v105 = int32(1)
	if v102 == v101 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v78 != 0 {
		goto L13
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v608 = int32(_a_F_dispell_init_4)
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v614 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dispell_init[2])))
	if base.B2i32(v611 == int32(0))|base.B2i32(v611 != v614) != 0 {
		v632 = v611
		v633 = v614
		goto L140
	} else {
		goto L141
	}
L25:
	;
	v117 = F_defGetString(m, v86)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v120 = F_get_tsearch_config_filename(m, v117, int32(_a_F_dispell_init_5))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v122 = m.G0
	v124 = v122 - int32(48)
	m.G0 = v124
	v127 = v124 + int32(4)
	v128 = F_tsearch_readline_begin(m, v127, v120)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	F_pfree(m, v120)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L138
	}
L29:
	;
	if v128 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v130 = F_tsearch_readline(m, v127)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L134
	}
L33:
	;
	if v130 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v138 = v130
	goto L37
L35:
	;
	goto L36
L36:
	;
	F_tsearch_readline_end(m, v124+int32(4))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L133
	}
L37:
	;
	v160 = v138
	goto L40
L38:
	;
	goto L36
L39:
	;
	v273 = v138
	goto L58
L40:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v186 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v195)
	v198 = v160 + int32(1)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v199 == v195 {
		v250 = v198
		goto L39
	} else {
		goto L49
	}
L42:
	;
	v250 = int32(_a_F_dispell_init_6)
	goto L39
L43:
	;
	goto L44
L44:
	;
	if v186 != int32(47) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v192 = F_pg_mblen_cstr(m, v160)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L41
L48:
	;
	v160 = v192 + v160
	goto L40
L49:
	;
	v205 = v198
	goto L50
L50:
	;
	v229 = F_pg_mblen_cstr(m, v205)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v243)
	v250 = v198
	goto L39
L52:
	;
	goto L51
L53:
	;
	if v229 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if base.Ui32((v233-int32(127))&int32(255)) < base.Ui32(int32(162)) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v240 != 0 {
		v205 = v205 + int32(1)
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v250 = v198
	goto L39
L57:
	;
	v305 = int32(_a_F_dispell_init_7)
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v308
	v310 = F_strlen(m, v138)
	mBase = m.M
	v312 = F_str_tolower(m, v138, v310, int32(100))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L63
	}
L58:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	switch v299 {
	case 0:
		goto L57
	default:
		goto L61
	case 9, 10, 11, 12, 13, 32:
		goto L60
	}
L59:
	;
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v303)
	goto L57
L60:
	;
	goto L59
L61:
	;
	v300 = F_pg_mblen_cstr(m, v273)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v273 = v300 + v273
	goto L58
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v306
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v57)+76))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
	if v316 <= v317 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v316 != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	v336 = F_strlen(m, v312)
	mBase = m.M
	v339 = F_MemoryContextAlloc(m, v335, v336+int32(9))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L73
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+68)) = v333
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+76)) = v316 << (uint(int32(1)) % 32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v325 = F_repalloc(m, v322, v316<<(uint(int32(3))%32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+76)) = int32(_a_F_dispell_init_8)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	v331 = F_MemoryContextAlloc(m, v329, int32(_a_F_dispell_init_9))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v333 = v325
	goto L67
L72:
	;
	v333 = v331
	goto L67
L73:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
	v343 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(v343)%32)))) = v339
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347+v348<<(uint(v343)%32))))
	v354 = v352 + int32(8)
	if (v312^v354)&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v429 != 0 {
		goto L95
	} else {
		goto L96
	}
L75:
	;
	goto L74
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v408)
	if v408&int32(255) == int32(0) {
		goto L75
	} else {
		goto L91
	}
L77:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	v407 = v312
	v408 = v360
	v409 = v354
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v312&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v364 = v312
	v366 = v354
	goto L83
L81:
	;
	v378 = v312
	v380 = v354
	goto L82
L82:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v385 = int32(-2139062144)
	if (int32(16843008)-v382|v382)&v385 != v385 {
		v407 = v378
		v408 = v382
		v409 = v380
		goto L76
	} else {
		goto L87
	}
L83:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v367)
	if v367 == int32(0) {
		goto L75
	} else {
		goto L85
	}
L84:
	;
	v378 = v374
	v380 = v372
	goto L82
L85:
	;
	v371 = int32(1)
	v372 = v366 + v371
	v374 = v364 + v371
	if v374&int32(3) != 0 {
		v364 = v374
		v366 = v372
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v390 = v378
	v391 = v382
	v392 = v380
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392))) = v391
	v394 = int32(4)
	v395 = v392 + v394
	v397 = v390 + v394
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v402 = int32(-2139062144)
	if (int32(16843008)-v399|v399)&v402 == v402 {
		v390 = v397
		v391 = v399
		v392 = v395
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v407 = v397
	v408 = v399
	v409 = v395
	goto L76
L90:
	;
	goto L89
L91:
	;
	v416 = v407
	v418 = v409
	goto L92
L92:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)) = uint8(v419)
	v421 = int32(1)
	if v419 != 0 {
		v416 = v416 + v421
		v418 = v418 + v421
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L75
L94:
	;
	goto L93
L95:
	;
	v430 = F_strlen(m, v250)
	mBase = m.M
	v432 = v430 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v432) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v535 = int32(_a_F_dispell_init_6)
	goto L97
L97:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v536+v537<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v535
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+72)) = v543 + int32(1)
	F_pfree(m, v312)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L129
	}
L98:
	;
	if (v250^v455)&int32(3) != 0 {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	v435 = F_palloc0(m, v432)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v440 = (v430 + int32(8)) & int32(4088)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	if base.Ui32(v440) <= base.Ui32(v441) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v455 = v435
	goto L98
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+84)) = v448 - v440
	*(*int32)(unsafe.Add(mBase, uint32(v57)+80)) = v440 + v449
	v455 = v449
	goto L98
L104:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	v448 = v441
	v449 = v443
	goto L103
L105:
	;
	goto L106
L106:
	;
	v444 = int32(_a_F_dispell_init_1)
	v446 = F_palloc0(m, v444)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v448 = v444
	v449 = v446
	goto L103
L108:
	;
	v535 = v455
	goto L97
L109:
	;
	goto L108
L110:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v511))) = uint8(v510)
	if v510&int32(255) == int32(0) {
		goto L109
	} else {
		goto L125
	}
L111:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v509 = v250
	v510 = v462
	v511 = v455
	goto L110
L112:
	;
	goto L113
L113:
	;
	if v250&int32(3) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v466 = v250
	v468 = v455
	goto L117
L115:
	;
	v480 = v250
	v482 = v455
	goto L116
L116:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v487 = int32(-2139062144)
	if (int32(16843008)-v484|v484)&v487 != v487 {
		v509 = v480
		v510 = v484
		v511 = v482
		goto L110
	} else {
		goto L121
	}
L117:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v469)
	if v469 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L118:
	;
	v480 = v476
	v482 = v474
	goto L116
L119:
	;
	v473 = int32(1)
	v474 = v468 + v473
	v476 = v466 + v473
	if v476&int32(3) != 0 {
		v466 = v476
		v468 = v474
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v492 = v480
	v493 = v484
	v494 = v482
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v493
	v496 = int32(4)
	v497 = v494 + v496
	v499 = v492 + v496
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v504 = int32(-2139062144)
	if (int32(16843008)-v501|v501)&v504 == v504 {
		v492 = v499
		v493 = v501
		v494 = v497
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v509 = v499
	v510 = v501
	v511 = v497
	goto L110
L124:
	;
	goto L123
L125:
	;
	v518 = v509
	v520 = v511
	goto L126
L126:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+1)) = uint8(v521)
	v523 = int32(1)
	if v521 != 0 {
		v518 = v518 + v523
		v520 = v520 + v523
		goto L126
	} else {
		goto L128
	}
L127:
	;
	goto L109
L128:
	;
	goto L127
L129:
	;
	F_pfree(m, v138)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v553 = F_tsearch_readline(m, v124+int32(4))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v553 != 0 {
		v138 = v553
		goto L37
	} else {
		goto L132
	}
L132:
	;
	goto L38
L133:
	;
	m.G0 = v124 + int32(48)
	goto L28
L134:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v120
	F_errmsg(m, int32(_a_F_dispell_init_10), v124)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(529), int32(_a_F_dispell_init_12))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v3630 = v70
	v3638 = int32(1)
	v3640 = v80
	goto L14
L139:
	;
	if v632-v633 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	goto L139
L141:
	;
	v617 = v87
	v618 = v608
	goto L142
L142:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+1)))
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	if v622 == int32(0) {
		v632 = v622
		v633 = v621
		goto L140
	} else {
		goto L144
	}
L143:
	;
	v632 = v622
	v633 = v621
	goto L140
L144:
	;
	v625 = int32(1)
	if v622 == v621 {
		v617 = v617 + v625
		v618 = v618 + v625
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	if v70 != 0 {
		goto L10
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v3583 = int32(_a_F_dispell_init_13)
	v3586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v3589 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dispell_init[4])))
	if base.B2i32(v3586 == int32(0))|base.B2i32(v3586 != v3589) != 0 {
		v3607 = v3586
		v3608 = v3589
		goto L930
	} else {
		goto L931
	}
L149:
	;
	v637 = F_defGetString(m, v86)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v640 = F_get_tsearch_config_filename(m, v637, int32(_a_F_dispell_init_14))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v642 = int32(0)
	v647 = m.G0
	v649 = v647 - int32(_a_F_dispell_init_15)
	m.G0 = v649
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+3248)) = uint8(v642)
	v654 = v649 + int32(132)
	v655 = F_tsearch_readline_begin(m, v654, v640)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L162
	}
L152:
	;
	F_pfree(m, v640)
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L1
	} else {
		goto L928
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L1
	} else {
		goto L924
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3549 = m.ExcPending
	if v3549 != 0 {
		goto L1
	} else {
		goto L920
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L1
	} else {
		goto L917
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L1
	} else {
		goto L913
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L1
	} else {
		goto L909
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L1
	} else {
		goto L905
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L1
	} else {
		goto L901
	}
L160:
	;
	m.G0 = v649 + int32(_a_F_dispell_init_15)
	goto L152
L161:
	;
	if v674&int32(1) != 0 {
		goto L153
	} else {
		goto L409
	}
L162:
	;
	if v655 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v657 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = v657
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)) = uint8(v657)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+44)) = uint8(v657)
	v663 = F_tsearch_readline(m, v654)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L1
	} else {
		goto L405
	}
L166:
	;
	if v663 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v666 = v649 + int32(1200)
	v670 = v649 + int32(3248)
	v674 = v642
	v676 = v663
	v690 = v642
	v692 = v642
	v693 = v642
	goto L170
L168:
	;
	goto L169
L169:
	;
	F_tsearch_readline_end(m, v649+int32(132))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L403
	}
L170:
	;
	v698 = F_strlen(m, v676)
	mBase = m.M
	v700 = F_str_tolower(m, v676, v698, int32(100))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L174
	}
L171:
	;
	goto L169
L172:
	;
	F_pfree(m, v676)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L399
	}
L173:
	;
	v1553 = v674
	v1569 = v1542
	v1571 = v692
	v1572 = v1545
	goto L172
L174:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if base.B2i32(v702 == int32(10))|base.B2i32(v702 == int32(35)) != 0 {
		v1542 = v690
		v1545 = v693
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v708 = int32(_a_F_dispell_init_16)
	goto L179
L176:
	;
	v909 = int32(_a_F_dispell_init_17)
	goto L215
L177:
	;
	if v746-v747 != 0 {
		goto L176
	} else {
		goto L190
	}
L179:
	;
	goto L180
L180:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v715 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v716 = v700
	v717 = v708
	v718 = int32(13)
	v719 = v715
	goto L185
L182:
	;
	v742 = v708
	v746 = int32(0)
	goto L183
L183:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	goto L177
L184:
	;
	v742 = v737
	v746 = v739
	goto L183
L185:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if base.B2i32(v719 != v721)|base.B2i32(v721 == int32(0)) != 0 {
		v737 = v717
		v739 = v719
		goto L184
	} else {
		goto L187
	}
L186:
	;
	v737 = v731
	v739 = int32(0)
	goto L184
L187:
	;
	v727 = v718 - int32(1)
	if v727 == int32(0) {
		v737 = v717
		v739 = v719
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v730 = int32(1)
	v731 = v717 + v730
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716)+1)))
	if v732 != 0 {
		v716 = v716 + v730
		v717 = v731
		v718 = v727
		v719 = v732
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	v756 = v676
	goto L191
L191:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	if v782 == int32(0) {
		goto L176
	} else {
		goto L193
	}
L192:
	;
	v796 = v756
	v799 = v782
	goto L198
L193:
	;
	if base.B2i32(v782 == int32(108))|base.B2i32(v782 == int32(76)) == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v792 = F_pg_mblen_cstr(m, v756)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	goto L192
L197:
	;
	v756 = v792 + v756
	goto L191
L198:
	;
	switch v799 & int32(255) {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L200
	default:
		goto L201
	}
L199:
	;
	v828 = int32(1)
	v830 = v796
	v833 = v799
	goto L203
L200:
	;
	goto L199
L201:
	;
	v824 = F_pg_mblen_cstr(m, v796)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v826 = v824 + v796
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	v796 = v826
	v799 = v827
	goto L198
L203:
	;
	v857 = v833 & int32(255)
	if base.B2i32(base.Ui32(v857-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v857 == int32(32)) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	if v857 == int32(0) {
		v1553 = v828
		v1569 = v690
		v1571 = v692
		v1572 = v693
		goto L172
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v878 = F_pg_mblen_cstr(m, v830)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L212
	}
L208:
	;
	v869 = F_pg_mblen_cstr(m, v830)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	if v869 != int32(1) {
		v1553 = v828
		v1569 = v690
		v1571 = v692
		v1572 = v693
		goto L172
	} else {
		goto L210
	}
L210:
	;
	F_addCompoundAffixFlagValue(m, v57, v830, int32(14))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v876 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+44)) = uint8(v876)
	v1553 = v828
	v1569 = v690
	v1571 = v692
	v1572 = v693
	goto L172
L212:
	;
	v880 = v878 + v830
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880))))
	v830 = v880
	v833 = v881
	goto L203
L213:
	;
	if v947-v948 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L215:
	;
	goto L216
L216:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v916 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v917 = v700
	v918 = v909
	v919 = int32(8)
	v920 = v916
	goto L221
L218:
	;
	v943 = v909
	v947 = int32(0)
	goto L219
L219:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
	goto L213
L220:
	;
	v943 = v938
	v947 = v940
	goto L219
L221:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918))))
	if base.B2i32(v920 != v922)|base.B2i32(v922 == int32(0)) != 0 {
		v938 = v918
		v940 = v920
		goto L220
	} else {
		goto L223
	}
L222:
	;
	v938 = v932
	v940 = int32(0)
	goto L220
L223:
	;
	v928 = v919 - int32(1)
	if v928 == int32(0) {
		v938 = v918
		v940 = v920
		goto L220
	} else {
		goto L224
	}
L224:
	;
	v931 = int32(1)
	v932 = v918 + v931
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917)+1)))
	if v933 != 0 {
		v917 = v917 + v931
		v918 = v932
		v919 = v928
		v920 = v933
		goto L221
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	v959 = int32(1)
	v1553 = v959
	v1569 = v959
	v1571 = v692
	v1572 = int32(0)
	goto L172
L227:
	;
	goto L228
L228:
	;
	v961 = int32(1)
	v962 = int32(_a_F_dispell_init_18)
	goto L231
L229:
	;
	if v1000-v1001 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L231:
	;
	goto L232
L232:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v969 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v970 = v700
	v971 = v962
	v972 = int32(8)
	v973 = v969
	goto L237
L234:
	;
	v996 = v962
	v1000 = int32(0)
	goto L235
L235:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	goto L229
L236:
	;
	v996 = v991
	v1000 = v993
	goto L235
L237:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	if base.B2i32(v973 != v975)|base.B2i32(v975 == int32(0)) != 0 {
		v991 = v971
		v993 = v973
		goto L236
	} else {
		goto L239
	}
L238:
	;
	v991 = v985
	v993 = int32(0)
	goto L236
L239:
	;
	v981 = v972 - int32(1)
	if v981 == int32(0) {
		v991 = v971
		v993 = v973
		goto L236
	} else {
		goto L240
	}
L240:
	;
	v984 = int32(1)
	v985 = v971 + v984
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+1)))
	if v986 != 0 {
		v970 = v970 + v984
		v971 = v985
		v972 = v981
		v973 = v986
		goto L237
	} else {
		goto L241
	}
L241:
	;
	goto L238
L242:
	;
	v1553 = v961
	v1569 = int32(0)
	v1571 = v692
	v1572 = int32(1)
	goto L172
L243:
	;
	goto L244
L244:
	;
	v1013 = int32(_a_F_dispell_init_19)
	goto L247
L245:
	;
	if v1051-v1052 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L247:
	;
	goto L248
L248:
	;
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v1020 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1021 = v700
	v1022 = v1013
	v1023 = int32(4)
	v1024 = v1020
	goto L253
L250:
	;
	v1047 = v1013
	v1051 = int32(0)
	goto L251
L251:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	goto L245
L252:
	;
	v1047 = v1042
	v1051 = v1044
	goto L251
L253:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022))))
	if base.B2i32(v1024 != v1026)|base.B2i32(v1026 == int32(0)) != 0 {
		v1042 = v1022
		v1044 = v1024
		goto L252
	} else {
		goto L255
	}
L254:
	;
	v1042 = v1036
	v1044 = int32(0)
	goto L252
L255:
	;
	v1032 = v1023 - int32(1)
	if v1032 == int32(0) {
		v1042 = v1022
		v1044 = v1024
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v1035 = int32(1)
	v1036 = v1022 + v1035
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)))
	if v1037 != 0 {
		v1021 = v1021 + v1035
		v1022 = v1036
		v1023 = v1032
		v1024 = v1037
		goto L253
	} else {
		goto L257
	}
L257:
	;
	goto L254
L258:
	;
	v1065 = v676 + int32(4)
	goto L262
L259:
	;
	goto L260
L260:
	;
	v1129 = int32(_a_F_dispell_init_20)
	goto L275
L261:
	;
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	v1116 = v1111 + base.B2i32(v1113 == int32(92))
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	if v1117 == int32(0) {
		goto L161
	} else {
		goto L270
	}
L262:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	if base.Ui32(v1091-int32(9)) < base.Ui32(int32(5)) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v1111 = v1065 + int32(1)
	v1112 = int32(64)
	goto L261
L264:
	;
	goto L263
L265:
	;
	v1105 = F_pg_mblen_cstr(m, v1065)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L269
	}
L266:
	;
	v1096 = int32(0)
	switch v1091 - int32(32) {
	case 0:
		goto L265
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v1111 = v1065
		v1112 = v1096
		goto L261
	case 10:
		goto L264
	default:
		goto L267
	}
L267:
	;
	if v1091 != int32(126) {
		v1111 = v1065
		v1112 = v1096
		goto L261
	} else {
		goto L268
	}
L268:
	;
	v1101 = int32(1)
	v1111 = v1065 + v1101
	v1112 = v1101
	goto L261
L269:
	;
	v1065 = v1105 + v1065
	goto L262
L270:
	;
	v1120 = F_pg_mblen_cstr(m, v1116)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	if v1120 != int32(1) {
		goto L161
	} else {
		goto L272
	}
L272:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	v1125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+3249)) = uint8(v1125)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+3248)) = uint8(v1124)
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116)+1)))
	switch v1128 {
	case 0, 9, 10, 11, 12, 13, 32, 35, 58:
		v1553 = v961
		v1569 = v690
		v1571 = v1112
		v1572 = v693
		goto L172
	default:
		goto L161
	}
L273:
	;
	if v1167-v1168 == int32(0) {
		goto L161
	} else {
		goto L286
	}
L275:
	;
	goto L276
L276:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v1136 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1137 = v676
	v1138 = v1129
	v1139 = int32(12)
	v1140 = v1136
	goto L281
L278:
	;
	v1163 = v1129
	v1167 = int32(0)
	goto L279
L279:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	goto L273
L280:
	;
	v1163 = v1158
	v1167 = v1160
	goto L279
L281:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138))))
	if base.B2i32(v1140 != v1142)|base.B2i32(v1142 == int32(0)) != 0 {
		v1158 = v1138
		v1160 = v1140
		goto L280
	} else {
		goto L283
	}
L282:
	;
	v1158 = v1152
	v1160 = int32(0)
	goto L280
L283:
	;
	v1148 = v1139 - int32(1)
	if v1148 == int32(0) {
		v1158 = v1138
		v1160 = v1140
		goto L280
	} else {
		goto L284
	}
L284:
	;
	v1151 = int32(1)
	v1152 = v1138 + v1151
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+1)))
	if v1153 != 0 {
		v1137 = v1137 + v1151
		v1138 = v1152
		v1139 = v1148
		v1140 = v1153
		goto L281
	} else {
		goto L285
	}
L285:
	;
	goto L282
L286:
	;
	v1178 = int32(_a_F_dispell_init_21)
	goto L289
L287:
	;
	if v1216-v1217 == int32(0) {
		goto L161
	} else {
		goto L300
	}
L289:
	;
	goto L290
L290:
	;
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v1185 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1186 = v676
	v1187 = v1178
	v1188 = int32(11)
	v1189 = v1185
	goto L295
L292:
	;
	v1212 = v1178
	v1216 = int32(0)
	goto L293
L293:
	;
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212))))
	goto L287
L294:
	;
	v1212 = v1207
	v1216 = v1209
	goto L293
L295:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187))))
	if base.B2i32(v1189 != v1191)|base.B2i32(v1191 == int32(0)) != 0 {
		v1207 = v1187
		v1209 = v1189
		goto L294
	} else {
		goto L297
	}
L296:
	;
	v1207 = v1201
	v1209 = int32(0)
	goto L294
L297:
	;
	v1197 = v1188 - int32(1)
	if v1197 == int32(0) {
		v1207 = v1187
		v1209 = v1189
		goto L294
	} else {
		goto L298
	}
L298:
	;
	v1200 = int32(1)
	v1201 = v1187 + v1200
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186)+1)))
	if v1202 != 0 {
		v1186 = v1186 + v1200
		v1187 = v1201
		v1188 = v1197
		v1189 = v1202
		goto L295
	} else {
		goto L299
	}
L299:
	;
	goto L296
L300:
	;
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	switch v1227 - int32(80) {
	case 0:
		goto L303
	default:
		goto L301
	case 3:
		goto L302
	}
L301:
	;
	if (v690|v693)&int32(1) == int32(0) {
		goto L308
	} else {
		goto L309
	}
L302:
	;
	v1236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v1236 != int32(70) {
		goto L301
	} else {
		goto L306
	}
L303:
	;
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v1230 != int32(70) {
		goto L301
	} else {
		goto L304
	}
L304:
	;
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+2)))
	if v1233 != int32(88) {
		goto L301
	} else {
		goto L305
	}
L305:
	;
	goto L161
L306:
	;
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+2)))
	if v1239 == int32(88) {
		goto L161
	} else {
		goto L307
	}
L307:
	;
	goto L301
L308:
	;
	v1247 = int32(0)
	v1542 = v1247
	v1545 = v1247
	goto L173
L309:
	;
	goto L310
L310:
	;
	v1249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+1200)) = uint8(v1249)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+176)) = uint8(v1249)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+2224)) = uint8(v1249)
	v1257 = v649 + int32(1200)
	v1259 = v649 + int32(176)
	v1261 = v649 + int32(2224)
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v1262 == v1249 {
		v1480 = v1259
		v1481 = v1261
		v1486 = v1257
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1480))) = uint8(v1496)
	*(*uint8)(unsafe.Add(mBase, uint32(v1486))) = uint8(v1496)
	*(*uint8)(unsafe.Add(mBase, uint32(v1481))) = uint8(v1496)
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+2224)))
	if v1502 == v1496 {
		v1542 = v690
		v1545 = v693
		goto L173
	} else {
		goto L396
	}
L312:
	;
	v1266 = v700
	v1275 = v1249
	v1276 = v1259
	v1277 = v1261
	v1282 = v1257
	goto L313
L313:
	;
	v1292 = F_pg_mblen_cstr(m, v1266)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L315
	}
L314:
	;
	v1480 = v1464
	v1481 = v1465
	v1486 = v1466
	goto L311
L315:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	switch v1275 - int32(1) {
	case 0:
		goto L322
	case 1:
		goto L321
	case 2:
		goto L320
	case 3:
		goto L319
	case 4:
		goto L318
	default:
		goto L323
	}
L316:
	;
	v1467 = v1266 + v1292
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1467))))
	if v1468 != 0 {
		v1266 = v1467
		v1275 = v1463
		v1276 = v1464
		v1277 = v1465
		v1282 = v1466
		goto L313
	} else {
		goto L395
	}
L317:
	;
	v1459 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1277))) = uint8(v1459)
	v1463 = int32(2)
	v1464 = v1276
	v1465 = v1277
	v1466 = v1282
	goto L316
L318:
	;
	if v1294 == int32(35) {
		goto L379
	} else {
		goto L380
	}
L319:
	;
	if v1294 == int32(45) {
		v1480 = v1276
		v1481 = v1277
		v1486 = v1282
		goto L311
	} else {
		goto L365
	}
L320:
	;
	if v1294 == int32(44) {
		goto L349
	} else {
		goto L350
	}
L321:
	;
	if v1294 == int32(45) {
		v1463 = int32(3)
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L335
	}
L322:
	;
	v1309 = int32(1)
	switch v1294 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v1463 = v1309
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	default:
		goto L330
	case 53:
		goto L317
	}
L323:
	;
	v1297 = int32(0)
	if base.Ui32(v1294-int32(9)) < base.Ui32(int32(5)) {
		v1463 = v1297
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L324
	}
L324:
	;
	switch v1294 - int32(32) {
	case 0:
		v1463 = v1297
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	default:
		goto L325
	case 3:
		v1553 = v674
		v1569 = v690
		v1571 = v692
		v1572 = v693
		goto L172
	}
L325:
	;
	v1304 = int32(1)
	if base.Ui32(v670-v1292) <= base.Ui32(v1277) {
		v1463 = v1304
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L326
	}
L326:
	;
	if v1292 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	base.MemoryCopy(m, v1277, v1266, v1292)
	goto L329
L328:
	;
	goto L329
L329:
	;
	v1463 = v1304
	v1464 = v1276
	v1465 = v1292 + v1277
	v1466 = v1282
	goto L316
L330:
	;
	if base.Ui32(v670-v1292) <= base.Ui32(v1277) {
		v1463 = v1309
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L331
	}
L331:
	;
	if v1292 != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	base.MemoryCopy(m, v1277, v1266, v1292)
	goto L334
L333:
	;
	goto L334
L334:
	;
	v1463 = v1309
	v1464 = v1276
	v1465 = v1292 + v1277
	v1466 = v1282
	goto L316
L335:
	;
	v1319 = F_t_isalpha_cstr(m, v1266)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L338
	}
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L345
	}
L337:
	;
	v1331 = int32(5)
	if base.Ui32(v666-v1292) <= base.Ui32(v1276) {
		v1463 = v1331
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L341
	}
L338:
	;
	if v1319 != 0 {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1321 = int32(2)
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	if base.Ui32(v1322-int32(9)) < base.Ui32(int32(5)) {
		v1463 = v1321
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L340
	}
L340:
	;
	switch v1322 - int32(32) {
	case 0:
		v1463 = v1321
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	default:
		goto L336
	case 7:
		goto L337
	}
L341:
	;
	if v1292 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	base.MemoryCopy(m, v1276, v1266, v1292)
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1463 = v1331
	v1464 = v1292 + v1276
	v1465 = v1277
	v1466 = v1282
	goto L316
L345:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	F_errmsg(m, int32(_a_F_dispell_init_22), int32(0))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(972), int32(_a_F_dispell_init_23))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	v1354 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1282))) = uint8(v1354)
	v1463 = int32(4)
	v1464 = v1276
	v1465 = v1277
	v1466 = v1282
	goto L316
L350:
	;
	goto L351
L351:
	;
	v1357 = F_t_isalpha_cstr(m, v1266)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	if v1357 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1359 = int32(3)
	if base.Ui32(v649+int32(2224)-v1292) <= base.Ui32(v1282) {
		v1463 = v1359
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	if base.B2i32(base.Ui32(v1365-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v1365 == int32(32)) != 0 {
		v1463 = int32(3)
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L360
	}
L356:
	;
	if v1292 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	base.MemoryCopy(m, v1282, v1266, v1292)
	goto L359
L358:
	;
	goto L359
L359:
	;
	v1463 = v1359
	v1464 = v1276
	v1465 = v1277
	v1466 = v1292 + v1282
	goto L316
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errmsg(m, int32(_a_F_dispell_init_22), int32(0))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(989), int32(_a_F_dispell_init_23))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	v1391 = F_t_isalpha_cstr(m, v1266)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	if v1391 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1393 = int32(5)
	if base.Ui32(v666-v1292) <= base.Ui32(v1276) {
		v1463 = v1393
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	if base.B2i32(base.Ui32(v1399-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v1399 == int32(32)) != 0 {
		v1463 = int32(4)
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L374
	}
L370:
	;
	if v1292 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	base.MemoryCopy(m, v1276, v1266, v1292)
	goto L373
L372:
	;
	goto L373
L373:
	;
	v1463 = v1393
	v1464 = v1292 + v1276
	v1465 = v1277
	v1466 = v1282
	goto L316
L374:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	F_errmsg(m, int32(_a_F_dispell_init_22), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1006), int32(_a_F_dispell_init_23))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	v1425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1276))) = uint8(v1425)
	v1480 = v1276
	v1481 = v1277
	v1486 = v1282
	goto L311
L380:
	;
	goto L381
L381:
	;
	v1427 = F_t_isalpha_cstr(m, v1266)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	if v1427 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1429 = int32(5)
	if base.Ui32(v666-v1292) <= base.Ui32(v1276) {
		v1463 = v1429
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1434 = int32(5)
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	if base.B2i32(base.Ui32(v1435-int32(9)) < base.Ui32(v1434))|base.B2i32(v1435 == int32(32)) != 0 {
		v1463 = v1434
		v1464 = v1276
		v1465 = v1277
		v1466 = v1282
		goto L316
	} else {
		goto L390
	}
L386:
	;
	if v1292 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	base.MemoryCopy(m, v1276, v1266, v1292)
	goto L389
L388:
	;
	goto L389
L389:
	;
	v1463 = v1429
	v1464 = v1292 + v1276
	v1465 = v1277
	v1466 = v1282
	goto L316
L390:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errmsg(m, int32(_a_F_dispell_init_22), int32(0))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1023), int32(_a_F_dispell_init_23))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	goto L314
L396:
	;
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+1200)))
	v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+176)))
	if v1505|v1506 == int32(0) {
		v1542 = v690
		v1545 = v693
		goto L173
	} else {
		goto L397
	}
L397:
	;
	F_NIAddAffix(m, v57, v649+int32(3248), base.I32_extend8_s(v692), v649+int32(2224), v649+int32(1200), v649+int32(176), v690&int32(1))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	v1542 = v690
	v1545 = v693
	goto L173
L399:
	;
	F_pfree(m, v700)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1583 = F_tsearch_readline(m, v649+int32(132))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	if v1583 != 0 {
		v674 = v1553
		v676 = v1583
		v690 = v1569
		v692 = v1571
		v693 = v1572
		goto L170
	} else {
		goto L402
	}
L402:
	;
	goto L171
L403:
	;
	F_finalizeCompoundAffixFlags(m, v57)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	goto L160
L405:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+128)) = v640
	F_errmsg(m, int32(_a_F_dispell_init_24), v649+int32(128))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1526), int32(_a_F_dispell_init_25))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
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
	F_tsearch_readline_end(m, v649+int32(132))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v1669 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = v1669
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)) = uint8(v1669)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+44)) = uint8(v1669)
	v1676 = v649 + int32(_a_F_dispell_init_26)
	v1677 = F_tsearch_readline_begin(m, v1676, v640)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	if v1677 == int32(0) {
		goto L154
	} else {
		goto L412
	}
L412:
	;
	v1681 = F_tsearch_readline(m, v1676)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	if v1681 != 0 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1684 = v1681
	goto L417
L415:
	;
	goto L416
L416:
	;
	F_tsearch_readline_end(m, v649+int32(_a_F_dispell_init_26))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L1
	} else {
		goto L620
	}
L417:
	;
	v1710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	switch v1710 {
	case 0, 9, 10, 11, 12, 13, 32, 35:
		goto L419
	default:
		goto L420
	}
L418:
	;
	goto L416
L419:
	;
	F_pfree(m, v1684)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L1
	} else {
		goto L617
	}
L420:
	;
	v1711 = int32(_a_F_dispell_init_20)
	goto L423
L421:
	;
	if v1749-v1750 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L423:
	;
	goto L424
L424:
	;
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1718 != 0 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1719 = v1684
	v1720 = v1711
	v1721 = int32(12)
	v1722 = v1718
	goto L429
L426:
	;
	v1745 = v1711
	v1749 = int32(0)
	goto L427
L427:
	;
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1745))))
	goto L421
L428:
	;
	v1745 = v1740
	v1749 = v1742
	goto L427
L429:
	;
	v1724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1720))))
	if base.B2i32(v1722 != v1724)|base.B2i32(v1724 == int32(0)) != 0 {
		v1740 = v1720
		v1742 = v1722
		goto L428
	} else {
		goto L431
	}
L430:
	;
	v1740 = v1734
	v1742 = int32(0)
	goto L428
L431:
	;
	v1730 = v1721 - int32(1)
	if v1730 == int32(0) {
		v1740 = v1720
		v1742 = v1722
		goto L428
	} else {
		goto L432
	}
L432:
	;
	v1733 = int32(1)
	v1734 = v1720 + v1733
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1719)+1)))
	if v1735 != 0 {
		v1719 = v1719 + v1733
		v1720 = v1734
		v1721 = v1730
		v1722 = v1735
		goto L429
	} else {
		goto L433
	}
L433:
	;
	goto L430
L434:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(12), int32(14))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L1
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	v1765 = int32(_a_F_dispell_init_27)
	goto L440
L437:
	;
	goto L419
L438:
	;
	if v1803-v1804 == int32(0) {
		goto L451
	} else {
		goto L452
	}
L440:
	;
	goto L441
L441:
	;
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1772 != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1773 = v1684
	v1774 = v1765
	v1775 = int32(13)
	v1776 = v1772
	goto L446
L443:
	;
	v1799 = v1765
	v1803 = int32(0)
	goto L444
L444:
	;
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1799))))
	goto L438
L445:
	;
	v1799 = v1794
	v1803 = v1796
	goto L444
L446:
	;
	v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774))))
	if base.B2i32(v1776 != v1778)|base.B2i32(v1778 == int32(0)) != 0 {
		v1794 = v1774
		v1796 = v1776
		goto L445
	} else {
		goto L448
	}
L447:
	;
	v1794 = v1788
	v1796 = int32(0)
	goto L445
L448:
	;
	v1784 = v1775 - int32(1)
	if v1784 == int32(0) {
		v1794 = v1774
		v1796 = v1776
		goto L445
	} else {
		goto L449
	}
L449:
	;
	v1787 = int32(1)
	v1788 = v1774 + v1787
	v1789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773)+1)))
	if v1789 != 0 {
		v1773 = v1773 + v1787
		v1774 = v1788
		v1775 = v1784
		v1776 = v1789
		goto L446
	} else {
		goto L450
	}
L450:
	;
	goto L447
L451:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(13), int32(2))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1819 = int32(_a_F_dispell_init_28)
	goto L457
L454:
	;
	goto L419
L455:
	;
	if v1857-v1858 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L457:
	;
	goto L458
L458:
	;
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1826 != 0 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1827 = v1684
	v1828 = v1819
	v1829 = int32(12)
	v1830 = v1826
	goto L463
L460:
	;
	v1853 = v1819
	v1857 = int32(0)
	goto L461
L461:
	;
	v1858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1853))))
	goto L455
L462:
	;
	v1853 = v1848
	v1857 = v1850
	goto L461
L463:
	;
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1828))))
	if base.B2i32(v1830 != v1832)|base.B2i32(v1832 == int32(0)) != 0 {
		v1848 = v1828
		v1850 = v1830
		goto L462
	} else {
		goto L465
	}
L464:
	;
	v1848 = v1842
	v1850 = int32(0)
	goto L462
L465:
	;
	v1838 = v1829 - int32(1)
	if v1838 == int32(0) {
		v1848 = v1828
		v1850 = v1830
		goto L462
	} else {
		goto L466
	}
L466:
	;
	v1841 = int32(1)
	v1842 = v1828 + v1841
	v1843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1827)+1)))
	if v1843 != 0 {
		v1827 = v1827 + v1841
		v1828 = v1842
		v1829 = v1838
		v1830 = v1843
		goto L463
	} else {
		goto L467
	}
L467:
	;
	goto L464
L468:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(12), int32(8))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	v1873 = int32(_a_F_dispell_init_29)
	goto L474
L471:
	;
	goto L419
L472:
	;
	if v1911-v1912 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L474:
	;
	goto L475
L475:
	;
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1880 != 0 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1881 = v1684
	v1882 = v1873
	v1883 = int32(11)
	v1884 = v1880
	goto L480
L477:
	;
	v1907 = v1873
	v1911 = int32(0)
	goto L478
L478:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1907))))
	goto L472
L479:
	;
	v1907 = v1902
	v1911 = v1904
	goto L478
L480:
	;
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1882))))
	if base.B2i32(v1884 != v1886)|base.B2i32(v1886 == int32(0)) != 0 {
		v1902 = v1882
		v1904 = v1884
		goto L479
	} else {
		goto L482
	}
L481:
	;
	v1902 = v1896
	v1904 = int32(0)
	goto L479
L482:
	;
	v1892 = v1883 - int32(1)
	if v1892 == int32(0) {
		v1902 = v1882
		v1904 = v1884
		goto L479
	} else {
		goto L483
	}
L483:
	;
	v1895 = int32(1)
	v1896 = v1882 + v1895
	v1897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1881)+1)))
	if v1897 != 0 {
		v1881 = v1881 + v1895
		v1882 = v1896
		v1883 = v1892
		v1884 = v1897
		goto L480
	} else {
		goto L484
	}
L484:
	;
	goto L481
L485:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(11), int32(8))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L1
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	v1927 = int32(_a_F_dispell_init_30)
	goto L491
L488:
	;
	goto L419
L489:
	;
	if v1965-v1966 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L491:
	;
	goto L492
L492:
	;
	v1934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1934 != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1935 = v1684
	v1936 = v1927
	v1937 = int32(14)
	v1938 = v1934
	goto L497
L494:
	;
	v1961 = v1927
	v1965 = int32(0)
	goto L495
L495:
	;
	v1966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	goto L489
L496:
	;
	v1961 = v1956
	v1965 = v1958
	goto L495
L497:
	;
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1936))))
	if base.B2i32(v1938 != v1940)|base.B2i32(v1940 == int32(0)) != 0 {
		v1956 = v1936
		v1958 = v1938
		goto L496
	} else {
		goto L499
	}
L498:
	;
	v1956 = v1950
	v1958 = int32(0)
	goto L496
L499:
	;
	v1946 = v1937 - int32(1)
	if v1946 == int32(0) {
		v1956 = v1936
		v1958 = v1938
		goto L496
	} else {
		goto L500
	}
L500:
	;
	v1949 = int32(1)
	v1950 = v1936 + v1949
	v1951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935)+1)))
	if v1951 != 0 {
		v1935 = v1935 + v1949
		v1936 = v1950
		v1937 = v1946
		v1938 = v1951
		goto L497
	} else {
		goto L501
	}
L501:
	;
	goto L498
L502:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(14), int32(4))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v1981 = int32(_a_F_dispell_init_31)
	goto L508
L505:
	;
	goto L419
L506:
	;
	if v2019-v2020 == int32(0) {
		goto L519
	} else {
		goto L520
	}
L508:
	;
	goto L509
L509:
	;
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v1988 != 0 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v1989 = v1684
	v1990 = v1981
	v1991 = int32(14)
	v1992 = v1988
	goto L514
L511:
	;
	v2015 = v1981
	v2019 = int32(0)
	goto L512
L512:
	;
	v2020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2015))))
	goto L506
L513:
	;
	v2015 = v2010
	v2019 = v2012
	goto L512
L514:
	;
	v1994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1990))))
	if base.B2i32(v1992 != v1994)|base.B2i32(v1994 == int32(0)) != 0 {
		v2010 = v1990
		v2012 = v1992
		goto L513
	} else {
		goto L516
	}
L515:
	;
	v2010 = v2004
	v2012 = int32(0)
	goto L513
L516:
	;
	v2000 = v1991 - int32(1)
	if v2000 == int32(0) {
		v2010 = v1990
		v2012 = v1992
		goto L513
	} else {
		goto L517
	}
L517:
	;
	v2003 = int32(1)
	v2004 = v1990 + v2003
	v2005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1989)+1)))
	if v2005 != 0 {
		v1989 = v1989 + v2003
		v1990 = v2004
		v1991 = v2000
		v1992 = v2005
		goto L514
	} else {
		goto L518
	}
L518:
	;
	goto L515
L519:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(14), int32(1))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L1
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	v2035 = int32(_a_F_dispell_init_32)
	goto L525
L522:
	;
	goto L419
L523:
	;
	if v2073-v2074 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L525:
	;
	goto L526
L526:
	;
	v2042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v2042 != 0 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v2043 = v1684
	v2044 = v2035
	v2045 = int32(18)
	v2046 = v2042
	goto L531
L528:
	;
	v2069 = v2035
	v2073 = int32(0)
	goto L529
L529:
	;
	v2074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2069))))
	goto L523
L530:
	;
	v2069 = v2064
	v2073 = v2066
	goto L529
L531:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2044))))
	if base.B2i32(v2046 != v2048)|base.B2i32(v2048 == int32(0)) != 0 {
		v2064 = v2044
		v2066 = v2046
		goto L530
	} else {
		goto L533
	}
L532:
	;
	v2064 = v2058
	v2066 = int32(0)
	goto L530
L533:
	;
	v2054 = v2045 - int32(1)
	if v2054 == int32(0) {
		v2064 = v2044
		v2066 = v2046
		goto L530
	} else {
		goto L534
	}
L534:
	;
	v2057 = int32(1)
	v2058 = v2044 + v2057
	v2059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2043)+1)))
	if v2059 != 0 {
		v2043 = v2043 + v2057
		v2044 = v2058
		v2045 = v2054
		v2046 = v2059
		goto L531
	} else {
		goto L535
	}
L535:
	;
	goto L532
L536:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(18), int32(16))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	v2089 = int32(_a_F_dispell_init_33)
	goto L542
L539:
	;
	goto L419
L540:
	;
	if v2127-v2128 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L542:
	;
	goto L543
L543:
	;
	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v2096 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2097 = v1684
	v2098 = v2089
	v2099 = int32(18)
	v2100 = v2096
	goto L548
L545:
	;
	v2123 = v2089
	v2127 = int32(0)
	goto L546
L546:
	;
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2123))))
	goto L540
L547:
	;
	v2123 = v2118
	v2127 = v2120
	goto L546
L548:
	;
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2098))))
	if base.B2i32(v2100 != v2102)|base.B2i32(v2102 == int32(0)) != 0 {
		v2118 = v2098
		v2120 = v2100
		goto L547
	} else {
		goto L550
	}
L549:
	;
	v2118 = v2112
	v2120 = int32(0)
	goto L547
L550:
	;
	v2108 = v2099 - int32(1)
	if v2108 == int32(0) {
		v2118 = v2098
		v2120 = v2100
		goto L547
	} else {
		goto L551
	}
L551:
	;
	v2111 = int32(1)
	v2112 = v2098 + v2111
	v2113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2097)+1)))
	if v2113 != 0 {
		v2097 = v2097 + v2111
		v2098 = v2112
		v2099 = v2108
		v2100 = v2113
		goto L548
	} else {
		goto L552
	}
L552:
	;
	goto L549
L553:
	;
	F_addCompoundAffixFlagValue(m, v57, v1684+int32(18), int32(32))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	v2143 = int32(_a_F_dispell_init_34)
	goto L559
L556:
	;
	goto L419
L557:
	;
	if v2181-v2182 != 0 {
		goto L419
	} else {
		goto L570
	}
L559:
	;
	goto L560
L560:
	;
	v2150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684))))
	if v2150 != 0 {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v2151 = v1684
	v2152 = v2143
	v2153 = int32(4)
	v2154 = v2150
	goto L565
L562:
	;
	v2177 = v2143
	v2181 = int32(0)
	goto L563
L563:
	;
	v2182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177))))
	goto L557
L564:
	;
	v2177 = v2172
	v2181 = v2174
	goto L563
L565:
	;
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2152))))
	if base.B2i32(v2154 != v2156)|base.B2i32(v2156 == int32(0)) != 0 {
		v2172 = v2152
		v2174 = v2154
		goto L564
	} else {
		goto L567
	}
L566:
	;
	v2172 = v2166
	v2174 = int32(0)
	goto L564
L567:
	;
	v2162 = v2153 - int32(1)
	if v2162 == int32(0) {
		v2172 = v2152
		v2174 = v2154
		goto L564
	} else {
		goto L568
	}
L568:
	;
	v2165 = int32(1)
	v2166 = v2152 + v2165
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2151)+1)))
	if v2167 != 0 {
		v2151 = v2151 + v2165
		v2152 = v2166
		v2153 = v2162
		v2154 = v2167
		goto L565
	} else {
		goto L569
	}
L569:
	;
	goto L566
L570:
	;
	v2196 = v1684 + int32(4)
	goto L571
L571:
	;
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	if base.B2i32(base.Ui32(v2219-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v2219 == int32(32)) == int32(0) {
		goto L574
	} else {
		goto L575
	}
L572:
	;
	v2232 = int32(_a_F_dispell_init_35)
	goto L581
L573:
	;
	goto L572
L574:
	;
	if v2219 != 0 {
		goto L573
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	v2229 = F_pg_mblen_cstr(m, v2196)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L1
	} else {
		goto L578
	}
L577:
	;
	goto L419
L578:
	;
	v2196 = v2229 + v2196
	goto L571
L579:
	;
	if v2270-v2271 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L581:
	;
	goto L582
L582:
	;
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	if v2239 != 0 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v2240 = v2196
	v2241 = v2232
	v2242 = int32(4)
	v2243 = v2239
	goto L587
L584:
	;
	v2266 = v2232
	v2270 = int32(0)
	goto L585
L585:
	;
	v2271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2266))))
	goto L579
L586:
	;
	v2266 = v2261
	v2270 = v2263
	goto L585
L587:
	;
	v2245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241))))
	if base.B2i32(v2243 != v2245)|base.B2i32(v2245 == int32(0)) != 0 {
		v2261 = v2241
		v2263 = v2243
		goto L586
	} else {
		goto L589
	}
L588:
	;
	v2261 = v2255
	v2263 = int32(0)
	goto L586
L589:
	;
	v2251 = v2242 - int32(1)
	if v2251 == int32(0) {
		v2261 = v2241
		v2263 = v2243
		goto L586
	} else {
		goto L590
	}
L590:
	;
	v2254 = int32(1)
	v2255 = v2241 + v2254
	v2256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2240)+1)))
	if v2256 != 0 {
		v2240 = v2240 + v2254
		v2241 = v2255
		v2242 = v2251
		v2243 = v2256
		goto L587
	} else {
		goto L591
	}
L591:
	;
	goto L588
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = int32(1)
	goto L419
L593:
	;
	goto L594
L594:
	;
	if v2219 != int32(110) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v2293 = int32(_a_F_dispell_init_36)
	goto L601
L596:
	;
	v2285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196)+1)))
	if v2285 != int32(117) {
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196)+2)))
	if v2288 != int32(109) {
		goto L595
	} else {
		goto L598
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = int32(2)
	goto L419
L599:
	;
	if v2331-v2332 == int32(0) {
		goto L419
	} else {
		goto L612
	}
L601:
	;
	goto L602
L602:
	;
	v2300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	if v2300 != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v2301 = v2196
	v2302 = v2293
	v2303 = int32(7)
	v2304 = v2300
	goto L607
L604:
	;
	v2327 = v2293
	v2331 = int32(0)
	goto L605
L605:
	;
	v2332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2327))))
	goto L599
L606:
	;
	v2327 = v2322
	v2331 = v2324
	goto L605
L607:
	;
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302))))
	if base.B2i32(v2304 != v2306)|base.B2i32(v2306 == int32(0)) != 0 {
		v2322 = v2302
		v2324 = v2304
		goto L606
	} else {
		goto L609
	}
L608:
	;
	v2322 = v2316
	v2324 = int32(0)
	goto L606
L609:
	;
	v2312 = v2303 - int32(1)
	if v2312 == int32(0) {
		v2322 = v2302
		v2324 = v2304
		goto L606
	} else {
		goto L610
	}
L610:
	;
	v2315 = int32(1)
	v2316 = v2302 + v2315
	v2317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2301)+1)))
	if v2317 != 0 {
		v2301 = v2301 + v2315
		v2302 = v2316
		v2303 = v2312
		v2304 = v2317
		goto L607
	} else {
		goto L611
	}
L611:
	;
	goto L608
L612:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	F_errmsg(m, int32(_a_F_dispell_init_37), int32(0))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1349), int32(_a_F_dispell_init_38))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L617:
	;
	v2389 = F_tsearch_readline(m, v649+int32(_a_F_dispell_init_26))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	if v2389 != 0 {
		v1684 = v2389
		goto L417
	} else {
		goto L619
	}
L619:
	;
	goto L418
L620:
	;
	F_finalizeCompoundAffixFlags(m, v57)
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	if int32(2) <= v2424 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	F_pg_qsort(m, v2427, v2424, int32(12), int32(1266))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L1
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	v2433 = v649 + int32(_a_F_dispell_init_26)
	v2434 = F_tsearch_readline_begin(m, v2433, v640)
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L1
	} else {
		goto L627
	}
L625:
	;
	goto L624
L626:
	;
	v2458 = int32(0)
	v2470 = v2458
	v2471 = v2436
	v2474 = v2458
	v2475 = v2458
	v2477 = v642
	v2478 = v2458
	goto L638
L627:
	;
	if v2434 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v2436 = F_tsearch_readline(m, v2433)
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L1
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L634
	}
L631:
	;
	if v2436 != 0 {
		goto L626
	} else {
		goto L632
	}
L632:
	;
	F_tsearch_readline_end(m, v2433)
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	goto L160
L634:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+96)) = v640
	F_errmsg(m, int32(_a_F_dispell_init_24), v649+int32(96))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1368), int32(_a_F_dispell_init_38))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L638:
	;
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2471))))
	switch v2489 {
	case 0, 9, 10, 11, 12, 13, 32, 35:
		v3398 = v2470
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	default:
		goto L641
	}
L639:
	;
	F_tsearch_readline_end(m, v3420)
	mBase = m.M
	v3424 = m.ExcPending
	if v3424 != 0 {
		goto L1
	} else {
		goto L894
	}
L640:
	;
	F_pfree(m, v2471)
	mBase = m.M
	v3418 = m.ExcPending
	if v3418 != 0 {
		goto L1
	} else {
		goto L891
	}
L641:
	;
	v2490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[5]))) = uint8(v2490)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[6]))) = uint8(v2490)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[7]))) = uint8(v2490)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[8]))) = uint8(v2490)
	*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[9]))) = uint8(v2490)
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2471))))
	if v2502 == v2490 {
		v2969 = v2490
		goto L642
	} else {
		goto L643
	}
L642:
	;
	if v2470 != 0 {
		goto L756
	} else {
		goto L757
	}
L643:
	;
	v2506 = v2471
	v2510 = int32(6)
	v2525 = v2490
	goto L644
L644:
	;
	v2542 = int32(1024)
	v2543 = int32(0)
	switch v2510 {
	case 0:
		goto L646
	default:
		goto L155
	case 2:
		goto L649
	case 4:
		goto L648
	case 6:
		goto L651
	case 7:
		goto L650
	}
L645:
	;
	v2875 = v2506
	v2878 = v2542
	v2879 = v2510
	v2891 = v649 + int32(_a_F_dispell_init_39)
	goto L735
L646:
	;
	goto L645
L647:
	;
	v2872 = v2525 + int32(1)
	v2873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2845))))
	if v2873 != 0 {
		v2506 = v2845
		v2510 = v2849
		v2525 = v2872
		goto L644
	} else {
		goto L732
	}
L648:
	;
	v2770 = v2506
	v2773 = v2542
	v2775 = v2543
	v2790 = v649 + int32(_a_F_dispell_init_40)
	goto L712
L649:
	;
	v2695 = v2506
	v2698 = v2542
	v2704 = v2543
	v2705 = v649 + int32(_a_F_dispell_init_41)
	goto L692
L650:
	;
	v2620 = v2506
	v2623 = v2542
	v2625 = v2543
	v2638 = v649 + int32(_a_F_dispell_init_42)
	goto L672
L651:
	;
	v2547 = v2506
	v2549 = v2543
	v2550 = v2542
	v2568 = v649 + int32(_a_F_dispell_init_43)
	goto L652
L652:
	;
	v2573 = F_pg_mblen_cstr(m, v2547)
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L1
	} else {
		goto L654
	}
L653:
	;
	v2614 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2611))) = uint8(v2614)
	if v2609 == v2614 {
		v2969 = v2525
		goto L642
	} else {
		goto L671
	}
L654:
	;
	v2575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547))))
	if v2549 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L655:
	;
	v2612 = v2547 + v2573
	v2613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2612))))
	if v2613 != 0 {
		v2547 = v2612
		v2549 = v2609
		v2550 = v2610
		v2568 = v2611
		goto L652
	} else {
		goto L670
	}
L656:
	;
	if v2573 != 0 {
		goto L667
	} else {
		goto L668
	}
L657:
	;
	v2578 = int32(0)
	if base.Ui32(v2575-int32(9)) < base.Ui32(int32(5)) {
		v2609 = v2578
		v2610 = v2550
		v2611 = v2568
		goto L655
	} else {
		goto L660
	}
L658:
	;
	goto L659
L659:
	;
	v2588 = v2575 - int32(9)
	v2595 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v2588))|base.B2i32(int32(1)<<(uint(v2588)%32)&int32(_a_F_dispell_init_44) == v2595) == v2595 {
		goto L663
	} else {
		goto L664
	}
L660:
	;
	switch v2575 - int32(32) {
	case 0:
		v2609 = v2578
		v2610 = v2550
		v2611 = v2568
		goto L655
	default:
		goto L661
	case 3:
		v2969 = v2525
		goto L642
	}
L661:
	;
	v2585 = int32(1)
	if v2573 < v2550 {
		v2605 = v2585
		goto L656
	} else {
		goto L662
	}
L662:
	;
	v2609 = v2585
	v2610 = v2550
	v2611 = v2568
	goto L655
L663:
	;
	v2601 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2568))) = uint8(v2601)
	v2845 = v2547
	v2849 = int32(7)
	goto L647
L664:
	;
	goto L665
L665:
	;
	v2603 = int32(1)
	if v2550 <= v2573 {
		v2609 = v2603
		v2610 = v2550
		v2611 = v2568
		goto L655
	} else {
		goto L666
	}
L666:
	;
	v2605 = v2603
	goto L656
L667:
	;
	base.MemoryCopy(m, v2568, v2547, v2573)
	goto L669
L668:
	;
	goto L669
L669:
	;
	v2609 = v2605
	v2610 = v2550 - v2573
	v2611 = v2573 + v2568
	goto L655
L670:
	;
	goto L653
L671:
	;
	v2845 = v2612
	v2849 = int32(7)
	goto L647
L672:
	;
	v2646 = F_pg_mblen_cstr(m, v2620)
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L1
	} else {
		goto L674
	}
L673:
	;
	v2689 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2686))) = uint8(v2689)
	if v2685 == v2689 {
		v2969 = v2525
		goto L642
	} else {
		goto L691
	}
L674:
	;
	v2648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620))))
	if v2625 == int32(0) {
		goto L677
	} else {
		goto L678
	}
L675:
	;
	v2687 = v2620 + v2646
	v2688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2687))))
	if v2688 != 0 {
		v2620 = v2687
		v2623 = v2684
		v2625 = v2685
		v2638 = v2686
		goto L672
	} else {
		goto L690
	}
L676:
	;
	if v2646 != 0 {
		goto L687
	} else {
		goto L688
	}
L677:
	;
	v2651 = int32(0)
	if base.Ui32(v2648-int32(9)) < base.Ui32(int32(5)) {
		v2684 = v2623
		v2685 = v2651
		v2686 = v2638
		goto L675
	} else {
		goto L680
	}
L678:
	;
	goto L679
L679:
	;
	v2661 = v2648 - int32(9)
	v2668 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v2661))|base.B2i32(int32(1)<<(uint(v2661)%32)&int32(_a_F_dispell_init_44) == v2668) == v2668 {
		goto L683
	} else {
		goto L684
	}
L680:
	;
	switch v2648 - int32(32) {
	case 0:
		v2684 = v2623
		v2685 = v2651
		v2686 = v2638
		goto L675
	default:
		goto L681
	case 3:
		v2969 = v2525
		goto L642
	}
L681:
	;
	v2658 = int32(1)
	if v2646 < v2623 {
		v2679 = v2658
		goto L676
	} else {
		goto L682
	}
L682:
	;
	v2684 = v2623
	v2685 = v2658
	v2686 = v2638
	goto L675
L683:
	;
	v2674 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2638))) = uint8(v2674)
	v2845 = v2620
	v2849 = int32(2)
	goto L647
L684:
	;
	goto L685
L685:
	;
	v2676 = int32(1)
	if v2623 <= v2646 {
		v2684 = v2623
		v2685 = v2676
		v2686 = v2638
		goto L675
	} else {
		goto L686
	}
L686:
	;
	v2679 = v2676
	goto L676
L687:
	;
	base.MemoryCopy(m, v2638, v2620, v2646)
	goto L689
L688:
	;
	goto L689
L689:
	;
	v2684 = v2623 - v2646
	v2685 = v2679
	v2686 = v2646 + v2638
	goto L675
L690:
	;
	goto L673
L691:
	;
	v2845 = v2687
	v2849 = int32(2)
	goto L647
L692:
	;
	v2721 = F_pg_mblen_cstr(m, v2695)
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L1
	} else {
		goto L694
	}
L693:
	;
	v2764 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2761))) = uint8(v2764)
	if v2760 == v2764 {
		v2969 = v2525
		goto L642
	} else {
		goto L711
	}
L694:
	;
	v2723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2695))))
	if v2704 == int32(0) {
		goto L697
	} else {
		goto L698
	}
L695:
	;
	v2762 = v2695 + v2721
	v2763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2762))))
	if v2763 != 0 {
		v2695 = v2762
		v2698 = v2759
		v2704 = v2760
		v2705 = v2761
		goto L692
	} else {
		goto L710
	}
L696:
	;
	if v2721 != 0 {
		goto L707
	} else {
		goto L708
	}
L697:
	;
	v2726 = int32(0)
	if base.Ui32(v2723-int32(9)) < base.Ui32(int32(5)) {
		v2759 = v2698
		v2760 = v2726
		v2761 = v2705
		goto L695
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	v2736 = v2723 - int32(9)
	v2743 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v2736))|base.B2i32(int32(1)<<(uint(v2736)%32)&int32(_a_F_dispell_init_44) == v2743) == v2743 {
		goto L703
	} else {
		goto L704
	}
L700:
	;
	switch v2723 - int32(32) {
	case 0:
		v2759 = v2698
		v2760 = v2726
		v2761 = v2705
		goto L695
	default:
		goto L701
	case 3:
		v2969 = v2525
		goto L642
	}
L701:
	;
	v2733 = int32(1)
	if v2721 < v2698 {
		v2754 = v2733
		goto L696
	} else {
		goto L702
	}
L702:
	;
	v2759 = v2698
	v2760 = v2733
	v2761 = v2705
	goto L695
L703:
	;
	v2749 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2705))) = uint8(v2749)
	v2845 = v2695
	v2849 = int32(4)
	goto L647
L704:
	;
	goto L705
L705:
	;
	v2751 = int32(1)
	if v2698 <= v2721 {
		v2759 = v2698
		v2760 = v2751
		v2761 = v2705
		goto L695
	} else {
		goto L706
	}
L706:
	;
	v2754 = v2751
	goto L696
L707:
	;
	base.MemoryCopy(m, v2705, v2695, v2721)
	goto L709
L708:
	;
	goto L709
L709:
	;
	v2759 = v2698 - v2721
	v2760 = v2754
	v2761 = v2721 + v2705
	goto L695
L710:
	;
	goto L693
L711:
	;
	v2845 = v2762
	v2849 = int32(4)
	goto L647
L712:
	;
	v2796 = F_pg_mblen_cstr(m, v2770)
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		goto L1
	} else {
		goto L714
	}
L713:
	;
	v2839 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2836))) = uint8(v2839)
	if v2835 == v2839 {
		v2969 = v2525
		goto L642
	} else {
		goto L731
	}
L714:
	;
	v2798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2770))))
	if v2775 == int32(0) {
		goto L717
	} else {
		goto L718
	}
L715:
	;
	v2837 = v2770 + v2796
	v2838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2837))))
	if v2838 != 0 {
		v2770 = v2837
		v2773 = v2834
		v2775 = v2835
		v2790 = v2836
		goto L712
	} else {
		goto L730
	}
L716:
	;
	if v2796 != 0 {
		goto L727
	} else {
		goto L728
	}
L717:
	;
	v2801 = int32(0)
	if base.Ui32(v2798-int32(9)) < base.Ui32(int32(5)) {
		v2834 = v2773
		v2835 = v2801
		v2836 = v2790
		goto L715
	} else {
		goto L720
	}
L718:
	;
	goto L719
L719:
	;
	v2811 = v2798 - int32(9)
	v2818 = int32(0)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v2811))|base.B2i32(int32(1)<<(uint(v2811)%32)&int32(_a_F_dispell_init_44) == v2818) == v2818 {
		goto L723
	} else {
		goto L724
	}
L720:
	;
	switch v2798 - int32(32) {
	case 0:
		v2834 = v2773
		v2835 = v2801
		v2836 = v2790
		goto L715
	default:
		goto L721
	case 3:
		v2969 = v2525
		goto L642
	}
L721:
	;
	v2808 = int32(1)
	if v2796 < v2773 {
		v2829 = v2808
		goto L716
	} else {
		goto L722
	}
L722:
	;
	v2834 = v2773
	v2835 = v2808
	v2836 = v2790
	goto L715
L723:
	;
	v2823 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2790))) = uint8(v2823)
	v2845 = v2770
	v2849 = v2823
	goto L647
L724:
	;
	goto L725
L725:
	;
	v2826 = int32(1)
	if v2773 <= v2796 {
		v2834 = v2773
		v2835 = v2826
		v2836 = v2790
		goto L715
	} else {
		goto L726
	}
L726:
	;
	v2829 = v2826
	goto L716
L727:
	;
	base.MemoryCopy(m, v2790, v2770, v2796)
	goto L729
L728:
	;
	goto L729
L729:
	;
	v2834 = v2773 - v2796
	v2835 = v2829
	v2836 = v2796 + v2790
	goto L715
L730:
	;
	goto L713
L731:
	;
	v2845 = v2837
	v2849 = v2839
	goto L647
L732:
	;
	v2969 = v2872
	goto L642
L733:
	;
	v2969 = v2525 + int32(1)
	goto L642
L734:
	;
	v2940 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2891))) = uint8(v2940)
	goto L733
L735:
	;
	v2901 = F_pg_mblen_cstr(m, v2875)
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L1
	} else {
		goto L737
	}
L736:
	;
	v2938 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2935))) = uint8(v2938)
	if v2934 != 0 {
		goto L733
	} else {
		goto L755
	}
L737:
	;
	v2903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2875))))
	if v2879 == int32(0) {
		goto L740
	} else {
		goto L741
	}
L738:
	;
	v2936 = v2875 + v2901
	v2937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2936))))
	if v2937 != 0 {
		v2875 = v2936
		v2878 = v2933
		v2879 = v2934
		v2891 = v2935
		goto L735
	} else {
		goto L754
	}
L739:
	;
	if v2901 != 0 {
		goto L751
	} else {
		goto L752
	}
L740:
	;
	v2906 = int32(0)
	if base.Ui32(v2903-int32(9)) < base.Ui32(int32(5)) {
		v2933 = v2878
		v2934 = v2906
		v2935 = v2891
		goto L738
	} else {
		goto L743
	}
L741:
	;
	goto L742
L742:
	;
	v2916 = v2903 - int32(9)
	if int32(1)<<(uint(v2916)%32)&int32(_a_F_dispell_init_44) != 0 {
		goto L746
	} else {
		goto L747
	}
L743:
	;
	switch v2903 - int32(32) {
	case 0:
		v2933 = v2878
		v2934 = v2906
		v2935 = v2891
		goto L738
	default:
		goto L744
	case 3:
		v2969 = v2525
		goto L642
	}
L744:
	;
	v2913 = int32(1)
	if v2901 < v2878 {
		v2928 = v2913
		goto L739
	} else {
		goto L745
	}
L745:
	;
	v2933 = v2878
	v2934 = v2913
	v2935 = v2891
	goto L738
L746:
	;
	v2924 = base.B2i32(base.Ui32(v2916) <= base.Ui32(int32(23)))
	goto L748
L747:
	;
	v2924 = int32(0)
	goto L748
L748:
	;
	if v2924 != 0 {
		goto L734
	} else {
		goto L749
	}
L749:
	;
	v2925 = int32(1)
	if v2878 <= v2901 {
		v2933 = v2878
		v2934 = v2925
		v2935 = v2891
		goto L738
	} else {
		goto L750
	}
L750:
	;
	v2928 = v2925
	goto L739
L751:
	;
	base.MemoryCopy(m, v2891, v2875, v2901)
	goto L753
L752:
	;
	goto L753
L753:
	;
	v2933 = v2878 - v2901
	v2934 = v2928
	v2935 = v2901 + v2891
	goto L738
L754:
	;
	goto L736
L755:
	;
	v2969 = v2525
	goto L642
L756:
	;
	F_pfree(m, v2470)
	mBase = m.M
	v2977 = m.ExcPending
	if v2977 != 0 {
		goto L1
	} else {
		goto L759
	}
L757:
	;
	goto L758
L758:
	;
	v2978 = int32(_a_F_dispell_init_7)
	v2979 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3]))
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v2981
	v2984 = v649 + int32(_a_F_dispell_init_43)
	v2985 = F_strlen(m, v2984)
	mBase = m.M
	v2987 = F_str_tolower(m, v2984, v2985, int32(100))
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L1
	} else {
		goto L760
	}
L759:
	;
	goto L758
L760:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v2979
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987))))
	if v2991 == int32(97) {
		goto L761
	} else {
		goto L762
	}
L761:
	;
	v2994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+1)))
	if v2994 != int32(102) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L764
	}
L762:
	;
	goto L763
L763:
	;
	if v2969 < int32(4) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L824
	}
L764:
	;
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v2997 == int32(0) {
		goto L765
	} else {
		goto L766
	}
L765:
	;
	v3000 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)) = uint8(v3000)
	v3007 = v649 + int32(_a_F_dispell_init_42)
	goto L769
L766:
	;
	goto L767
L767:
	;
	if v2474 < v2475 {
		goto L786
	} else {
		goto L787
	}
L768:
	;
	if v3051 <= int32(0) {
		goto L159
	} else {
		goto L784
	}
L769:
	;
	v3012 = v3007 + int32(1)
	v3013 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3007))))
	v3014 = F___isspace(m, v3013)
	mBase = m.M
	if v3014 != 0 {
		v3007 = v3012
		goto L769
	} else {
		goto L771
	}
L770:
	;
	v3015 = int32(1)
	switch v3013&int32(255) - int32(43) {
	case 0:
		v3021 = v3015
		goto L773
	default:
		v3023 = v3013
		v3024 = v3007
		v3025 = v3015
		goto L772
	case 2:
		goto L774
	}
L771:
	;
	goto L770
L772:
	;
	v3026 = int32(0)
	v3028 = v3023 - int32(48)
	if base.Ui32(v3028) <= base.Ui32(int32(9)) {
		goto L775
	} else {
		goto L776
	}
L773:
	;
	v3022 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3012))))
	v3023 = v3022
	v3024 = v3012
	v3025 = v3021
	goto L772
L774:
	;
	v3021 = int32(0)
	goto L773
L775:
	;
	v3031 = v3026
	v3032 = v3028
	v3033 = v3024
	goto L778
L776:
	;
	v3045 = v3026
	goto L777
L777:
	;
	if v3025 != 0 {
		goto L781
	} else {
		goto L782
	}
L778:
	;
	v3035 = int32(10)
	v3037 = v3031*v3035 - v3032
	v3038 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3033)+1)))
	v3042 = v3038 - int32(48)
	if base.Ui32(v3042) < base.Ui32(v3035) {
		v3031 = v3037
		v3032 = v3042
		v3033 = v3033 + int32(1)
		goto L778
	} else {
		goto L780
	}
L779:
	;
	v3045 = v3037
	goto L777
L780:
	;
	goto L779
L781:
	;
	v3051 = int32(0) - v3045
	goto L783
L782:
	;
	v3051 = v3045
	goto L783
L783:
	;
	goto L768
L784:
	;
	v3056 = v3051 + int32(1)
	v3057 = F_palloc0_mul(m, int32(4), v3056)
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L1
	} else {
		goto L785
	}
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v3056
	*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v3057
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v3056
	*(*int32)(unsafe.Add(mBase, uint32(v3057+v2474<<(uint(int32(2))%32)))) = int32(_a_F_dispell_init_6)
	v3398 = v2987
	v3402 = v2474 + int32(1)
	v3403 = v3056
	v3405 = v2477
	v3406 = v2478
	goto L640
L786:
	;
	v3072 = F_strlen(m, v649+int32(_a_F_dispell_init_42))
	mBase = m.M
	v3074 = v3072 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v3074) {
		goto L790
	} else {
		goto L791
	}
L787:
	;
	goto L788
L788:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L1
	} else {
		goto L820
	}
L789:
	;
	v3100 = v649 + int32(_a_F_dispell_init_42)
	if (v3100^v3096)&int32(3) != 0 {
		goto L802
	} else {
		goto L803
	}
L790:
	;
	v3077 = F_palloc0(m, v3074)
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L1
	} else {
		goto L793
	}
L791:
	;
	goto L792
L792:
	;
	v3082 = (v3072 + int32(8)) & int32(4088)
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	if base.Ui32(v3082) <= base.Ui32(v3083) {
		goto L795
	} else {
		goto L796
	}
L793:
	;
	v3096 = v3077
	goto L789
L794:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+84)) = v3090 - v3082
	*(*int32)(unsafe.Add(mBase, uint32(v57)+80)) = v3091 + v3082
	v3096 = v3091
	goto L789
L795:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	v3090 = v3083
	v3091 = v3085
	goto L794
L796:
	;
	goto L797
L797:
	;
	v3086 = int32(_a_F_dispell_init_1)
	v3088 = F_palloc0(m, v3086)
	mBase = m.M
	v3089 = m.ExcPending
	if v3089 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	v3090 = v3086
	v3091 = v3088
	goto L794
L799:
	;
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3175+v2474<<(uint(int32(2))%32)))) = v3096
	v3398 = v2987
	v3402 = v2474 + int32(1)
	v3403 = v2475
	v3405 = v2477
	v3406 = v2478
	goto L640
L800:
	;
	goto L799
L801:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3155))) = uint8(v3154)
	if v3154&int32(255) == int32(0) {
		goto L800
	} else {
		goto L816
	}
L802:
	;
	v3106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[8]))))
	v3153 = v3100
	v3154 = v3106
	v3155 = v3096
	goto L801
L803:
	;
	goto L804
L804:
	;
	if v3100&int32(3) != 0 {
		goto L805
	} else {
		goto L806
	}
L805:
	;
	v3110 = v3100
	v3112 = v3096
	goto L808
L806:
	;
	v3124 = v3100
	v3126 = v3096
	goto L807
L807:
	;
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3124)))
	v3131 = int32(-2139062144)
	if (int32(16843008)-v3128|v3128)&v3131 != v3131 {
		v3153 = v3124
		v3154 = v3128
		v3155 = v3126
		goto L801
	} else {
		goto L812
	}
L808:
	;
	v3113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3112))) = uint8(v3113)
	if v3113 == int32(0) {
		goto L800
	} else {
		goto L810
	}
L809:
	;
	v3124 = v3120
	v3126 = v3118
	goto L807
L810:
	;
	v3117 = int32(1)
	v3118 = v3112 + v3117
	v3120 = v3110 + v3117
	if v3120&int32(3) != 0 {
		v3110 = v3120
		v3112 = v3118
		goto L808
	} else {
		goto L811
	}
L811:
	;
	goto L809
L812:
	;
	v3136 = v3124
	v3137 = v3128
	v3138 = v3126
	goto L813
L813:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3138))) = v3137
	v3140 = int32(4)
	v3141 = v3138 + v3140
	v3143 = v3136 + v3140
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+4))
	v3148 = int32(-2139062144)
	if (int32(16843008)-v3145|v3145)&v3148 == v3148 {
		v3136 = v3143
		v3137 = v3145
		v3138 = v3141
		goto L813
	} else {
		goto L815
	}
L814:
	;
	v3153 = v3143
	v3154 = v3145
	v3155 = v3141
	goto L801
L815:
	;
	goto L814
L816:
	;
	v3162 = v3153
	v3164 = v3155
	goto L817
L817:
	;
	v3165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3162)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3164)+1)) = uint8(v3165)
	v3167 = int32(1)
	if v3165 != 0 {
		v3162 = v3162 + v3167
		v3164 = v3164 + v3167
		goto L817
	} else {
		goto L819
	}
L818:
	;
	goto L800
L819:
	;
	goto L818
L820:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3188 = m.ExcPending
	if v3188 != 0 {
		goto L1
	} else {
		goto L821
	}
L821:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+16)) = v2475 - int32(1)
	F_errmsg(m, int32(_a_F_dispell_init_45), v649+int32(16))
	mBase = m.M
	v3196 = m.ExcPending
	if v3196 != 0 {
		goto L1
	} else {
		goto L822
	}
L822:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1418), int32(_a_F_dispell_init_38))
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L1
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
	switch v2991 - int32(112) {
	case 0:
		goto L826
	default:
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	case 3:
		goto L827
	}
L825:
	;
	v3223 = F_strlen(m, v649+int32(_a_F_dispell_init_42))
	mBase = m.M
	if v3223 == int32(0) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L832
	}
L826:
	;
	v3213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+1)))
	if v3213 != int32(102) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L830
	}
L827:
	;
	v3206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+1)))
	if v3206 != int32(102) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L828
	}
L828:
	;
	v3209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+2)))
	if v3209 != int32(120) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L829
	}
L829:
	;
	v3220 = int32(1)
	goto L825
L830:
	;
	v3216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987)+2)))
	if v3216 != int32(120) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L831
	}
L831:
	;
	v3220 = int32(0)
	goto L825
L832:
	;
	if v3223 < int32(2) {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	if v2969 == int32(4) {
		goto L838
	} else {
		goto L839
	}
L834:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v3228 == int32(0) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L835
	}
L835:
	;
	if v3223 == int32(2) {
		goto L833
	} else {
		goto L836
	}
L836:
	;
	if v3228 == int32(1) {
		v3398 = v2987
		v3402 = v2474
		v3403 = v2475
		v3405 = v2477
		v3406 = v2478
		goto L640
	} else {
		goto L837
	}
L837:
	;
	goto L833
L838:
	;
	v3240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[7]))))
	if v3240&int32(223) == int32(89) {
		goto L841
	} else {
		goto L842
	}
L839:
	;
	goto L840
L840:
	;
	v3248 = int32(47)
	v3249 = F___strchrnul(m, v649+int32(_a_F_dispell_init_40), v3248)
	mBase = m.M
	v3251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3249))))
	if v3251 == v3248 {
		goto L845
	} else {
		goto L846
	}
L841:
	;
	v3245 = int32(64)
	goto L843
L842:
	;
	v3245 = int32(0)
	goto L843
L843:
	;
	v3398 = v2987
	v3402 = v2474
	v3403 = v2475
	v3405 = v3220
	v3406 = v3245
	goto L640
L844:
	;
	if v3255 != 0 {
		goto L848
	} else {
		goto L849
	}
L845:
	;
	v3255 = v3249
	goto L847
L846:
	;
	v3255 = int32(0)
	goto L847
L847:
	;
	goto L844
L848:
	;
	v3256 = int32(1)
	v3257 = v3255 + v3256
	v3258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v3258 != v3256 {
		goto L852
	} else {
		goto L853
	}
L849:
	;
	v3319 = v2979
	v3322 = v2478
	goto L850
L850:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v3324
	v3327 = v649 + int32(_a_F_dispell_init_40)
	v3328 = F_strlen(m, v3327)
	mBase = m.M
	v3330 = F_str_tolower(m, v3327, v3328, int32(100))
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L1
	} else {
		goto L871
	}
L851:
	;
	v3314 = F_getCompoundAffixFlagValue(m, v57, v3312)
	mBase = m.M
	v3315 = m.ExcPending
	if v3315 != 0 {
		goto L1
	} else {
		goto L870
	}
L852:
	;
	v3312 = v3257
	goto L851
L853:
	;
	goto L854
L854:
	;
	v3261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3257))))
	if v3261 == int32(0) {
		goto L855
	} else {
		goto L856
	}
L855:
	;
	v3312 = v3257
	goto L851
L856:
	;
	goto L857
L857:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[10])) = int32(0)
	v3271 = F_strtox_2(m, v3257, v649+int32(_a_F_dispell_init_46), int32(10), int64(2147483648))
	mBase = m.M
	v3272 = base.I32_wrap_i64(v3271)
	goto L858
L858:
	;
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[11])))
	if v3257 == v3273 {
		goto L158
	} else {
		goto L859
	}
L859:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[10]))
	if v3276 == int32(68) {
		goto L158
	} else {
		goto L860
	}
L860:
	;
	v3279 = int32(0)
	v3281 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if base.B2i32(v3272 <= v3279)|base.B2i32(v3281 <= v3272) == v3279 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v3286+v3272<<(uint(int32(2))%32))))
	if v3290 != 0 {
		v3312 = v3290
		goto L851
	} else {
		goto L864
	}
L862:
	;
	goto L863
L863:
	;
	if v3281 < v3272 {
		goto L157
	} else {
		goto L869
	}
L864:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+48)) = v3257
	F_errmsg(m, int32(_a_F_dispell_init_47), v649+int32(48))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1241), int32(_a_F_dispell_init_48))
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L868
	}
L868:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L869:
	;
	v3312 = int32(_a_F_dispell_init_6)
	goto L851
L870:
	;
	v3318 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3]))
	v3319 = v3318
	v3322 = v3314 | v2478
	goto L850
L871:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v3319
	v3334 = int32(47)
	v3335 = F___strchrnul(m, v3330, v3334)
	mBase = m.M
	v3337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3335))))
	if v3337 == v3334 {
		goto L873
	} else {
		goto L874
	}
L872:
	;
	if v3341 != 0 {
		goto L876
	} else {
		goto L877
	}
L873:
	;
	v3341 = v3335
	goto L875
L874:
	;
	v3341 = int32(0)
	goto L875
L875:
	;
	goto L872
L876:
	;
	v3342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3341))) = uint8(v3342)
	v3345 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3]))
	v3346 = v3345
	goto L878
L877:
	;
	v3346 = v3319
	goto L878
L878:
	;
	v3348 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v3348
	v3351 = v649 + int32(_a_F_dispell_init_41)
	v3352 = F_strlen(m, v3351)
	mBase = m.M
	v3354 = F_str_tolower(m, v3351, v3352, int32(100))
	mBase = m.M
	v3355 = m.ExcPending
	if v3355 != 0 {
		goto L1
	} else {
		goto L879
	}
L879:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v3357
	v3360 = v649 + int32(_a_F_dispell_init_39)
	v3361 = F_strlen(m, v3360)
	mBase = m.M
	v3363 = F_str_tolower(m, v3360, v3361, int32(100))
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L1
	} else {
		goto L880
	}
L880:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[3])) = v3346
	v3367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[7]))))
	if v3367 == int32(48) {
		goto L881
	} else {
		goto L882
	}
L881:
	;
	v3370 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3354))) = uint8(v3370)
	goto L883
L882:
	;
	goto L883
L883:
	;
	v3372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_c_F_dispell_init[5]))))
	if v3372 == int32(48) {
		goto L884
	} else {
		goto L885
	}
L884:
	;
	v3375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3330))) = uint8(v3375)
	goto L886
L885:
	;
	goto L886
L886:
	;
	F_NIAddAffix(m, v57, v649+int32(_a_F_dispell_init_42), base.I32_extend8_s(v3322), v3363, v3354, v3330, v2477&int32(1))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	F_pfree(m, v3330)
	mBase = m.M
	v3385 = m.ExcPending
	if v3385 != 0 {
		goto L1
	} else {
		goto L888
	}
L888:
	;
	F_pfree(m, v3354)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L1
	} else {
		goto L889
	}
L889:
	;
	F_pfree(m, v3363)
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L1
	} else {
		goto L890
	}
L890:
	;
	v3398 = v2987
	v3402 = v2474
	v3403 = v2475
	v3405 = v2477
	v3406 = v2478
	goto L640
L891:
	;
	v3420 = v649 + int32(_a_F_dispell_init_26)
	v3421 = F_tsearch_readline(m, v3420)
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	if v3421 != 0 {
		v2470 = v3398
		v2471 = v3421
		v2474 = v3402
		v2475 = v3403
		v2477 = v3405
		v2478 = v3406
		goto L638
	} else {
		goto L893
	}
L893:
	;
	goto L639
L894:
	;
	if v3398 != 0 {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	F_pfree(m, v3398)
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L898
	}
L896:
	;
	goto L897
L897:
	;
	v3427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v3427 != int32(1) {
		goto L160
	} else {
		goto L899
	}
L898:
	;
	goto L897
L899:
	;
	if v3402 != v3403 {
		goto L156
	} else {
		goto L900
	}
L900:
	;
	goto L160
L901:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	F_errmsg(m, int32(_a_F_dispell_init_49), int32(0))
	mBase = m.M
	v3471 = m.ExcPending
	if v3471 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1394), int32(_a_F_dispell_init_38))
	mBase = m.M
	v3476 = m.ExcPending
	if v3476 != 0 {
		goto L1
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
	F_errcode(m, int32(22))
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+32)) = v3257
	F_errmsg(m, int32(_a_F_dispell_init_47), v649+int32(32))
	mBase = m.M
	v3489 = m.ExcPending
	if v3489 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1234), int32(_a_F_dispell_init_48))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L1
	} else {
		goto L908
	}
L908:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L909:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3501 = m.ExcPending
	if v3501 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+64)) = v3257
	F_errmsg(m, int32(_a_F_dispell_init_47), v649-int32(-64))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L1
	} else {
		goto L911
	}
L911:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1252), int32(_a_F_dispell_init_48))
	mBase = m.M
	v3512 = m.ExcPending
	if v3512 != 0 {
		goto L1
	} else {
		goto L912
	}
L912:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L913:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L1
	} else {
		goto L914
	}
L914:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649))) = v3403 - int32(1)
	F_errmsg(m, int32(_a_F_dispell_init_50), v649)
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1492), int32(_a_F_dispell_init_38))
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L917:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+80)) = v2510
	F_errmsg_internal(m, int32(_a_F_dispell_init_51), v649+int32(80))
	mBase = m.M
	v3540 = m.ExcPending
	if v3540 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(893), int32(_a_F_dispell_init_52))
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L920:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3552 = m.ExcPending
	if v3552 != 0 {
		goto L1
	} else {
		goto L921
	}
L921:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+112)) = v640
	F_errmsg(m, int32(_a_F_dispell_init_24), v649+int32(112))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1294), int32(_a_F_dispell_init_38))
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L924:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v3570 = m.ExcPending
	if v3570 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	F_errmsg(m, int32(_a_F_dispell_init_53), int32(0))
	mBase = m.M
	v3574 = m.ExcPending
	if v3574 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1646), int32(_a_F_dispell_init_25))
	mBase = m.M
	v3579 = m.ExcPending
	if v3579 != 0 {
		goto L1
	} else {
		goto L927
	}
L927:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L928:
	;
	v3630 = int32(1)
	v3638 = v78
	v3640 = v80
	goto L14
L929:
	;
	if v3607-v3608 != 0 {
		goto L8
	} else {
		goto L936
	}
L930:
	;
	goto L929
L931:
	;
	v3592 = v87
	v3593 = v3583
	goto L932
L932:
	;
	v3596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3593)+1)))
	v3597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3592)+1)))
	if v3597 == int32(0) {
		v3607 = v3597
		v3608 = v3596
		goto L930
	} else {
		goto L934
	}
L933:
	;
	v3607 = v3597
	v3608 = v3596
	goto L930
L934:
	;
	v3600 = int32(1)
	if v3597 == v3596 {
		v3592 = v3592 + v3600
		v3593 = v3593 + v3600
		goto L932
	} else {
		goto L935
	}
L935:
	;
	goto L933
L936:
	;
	if v80 != 0 {
		goto L9
	} else {
		goto L937
	}
L937:
	;
	v3610 = F_defGetString(m, v86)
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L1
	} else {
		goto L938
	}
L938:
	;
	F_readstoplist(m, v3610, v81)
	mBase = m.M
	v3613 = m.ExcPending
	if v3613 != 0 {
		goto L1
	} else {
		goto L939
	}
L939:
	;
	v3630 = v70
	v3638 = v78
	v3640 = int32(1)
	goto L14
L940:
	;
	v3713 = v57
	v3726 = v3630
	v3734 = v3638
	v3735 = v79
	v3737 = v81
	goto L6
L941:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3652 = m.ExcPending
	if v3652 != 0 {
		goto L1
	} else {
		goto L942
	}
L942:
	;
	F_errmsg(m, int32(_a_F_dispell_init_54), int32(0))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L1
	} else {
		goto L943
	}
L943:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(55), int32(_a_F_dispell_init_56))
	mBase = m.M
	v3661 = m.ExcPending
	if v3661 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L945:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L1
	} else {
		goto L946
	}
L946:
	;
	F_errmsg(m, int32(_a_F_dispell_init_57), int32(0))
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L1
	} else {
		goto L947
	}
L947:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(69), int32(_a_F_dispell_init_56))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L1
	} else {
		goto L948
	}
L948:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L949:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	F_errmsg(m, int32(_a_F_dispell_init_58), int32(0))
	mBase = m.M
	v3688 = m.ExcPending
	if v3688 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(81), int32(_a_F_dispell_init_56))
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L953:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3700 = m.ExcPending
	if v3700 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v3701
	F_errmsg(m, int32(_a_F_dispell_init_59), v79)
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(90), int32(_a_F_dispell_init_56))
	mBase = m.M
	v3710 = m.ExcPending
	if v3710 != 0 {
		goto L1
	} else {
		goto L956
	}
L956:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L957:
	;
	v3745 = int32(0)
	v3747 = m.G0
	v3749 = v3747 - int32(48)
	m.G0 = v3749
	v3751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3713)+36)))
	if v3751 == int32(1) {
		goto L965
	} else {
		goto L966
	}
L958:
	;
	goto L959
L959:
	;
	if v3726 == int32(0) {
		goto L4
	} else {
		goto L1152
	}
L960:
	;
	v4423 = m.G0
	v4425 = v4423 - int32(1040)
	m.G0 = v4425
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+4))
	if v4427 != 0 {
		goto L1088
	} else {
		goto L1089
	}
L961:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L1
	} else {
		goto L1084
	}
L962:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4377 = m.ExcPending
	if v4377 != 0 {
		goto L1
	} else {
		goto L1080
	}
L963:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L1
	} else {
		goto L1076
	}
L964:
	;
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	F_pg_qsort(m, v4338, v4311, int32(4), int32(1268))
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L1
	} else {
		goto L1074
	}
L965:
	;
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v3754 <= int32(0) {
		v4311 = v3754
		goto L964
	} else {
		goto L968
	}
L966:
	;
	goto L967
L967:
	;
	v3853 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3854 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	F_pg_qsort(m, v3853, v3854, int32(4), int32(1267))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L981
	}
L968:
	;
	v3758 = v3745
	goto L969
L969:
	;
	v3784 = int32(0)
	v3786 = v3758 << (uint(int32(2)) % 32)
	v3787 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3789 = *(*int32)(unsafe.Add(mBase, uint32(v3786+v3787)))
	v3790 = *(*int32)(unsafe.Add(mBase, uint32(v3789)))
	v3791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3790))))
	if v3791 == v3784 {
		v3838 = v3784
		v3840 = v3789
		goto L971
	} else {
		goto L972
	}
L970:
	;
	v4311 = v3851
	goto L964
L971:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3840))) = v3838
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3842+v3786)))
	v3847 = F_strlen(m, v3844+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3844)+4)) = v3847
	v3850 = v3758 + int32(1)
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v3850 < v3851 {
		v3758 = v3850
		goto L969
	} else {
		goto L980
	}
L972:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispell_init[10])) = int32(0)
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3799 = *(*int32)(unsafe.Add(mBase, uint32(v3797+v3786)))
	v3800 = *(*int32)(unsafe.Add(mBase, uint32(v3799)))
	v3805 = F_strtox_2(m, v3800, v3749+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	v3806 = base.I32_wrap_i64(v3805)
	goto L973
L973:
	;
	v3807 = *(*int32)(unsafe.Add(mBase, uint32(v3749)+44))
	v3808 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3810 = *(*int32)(unsafe.Add(mBase, uint32(v3808+v3786)))
	v3811 = *(*int32)(unsafe.Add(mBase, uint32(v3810)))
	if v3807 == v3811 {
		goto L963
	} else {
		goto L974
	}
L974:
	;
	v3814 = *(*int32)(unsafe.Add(mBase, _c_F_dispell_init[10]))
	if v3814 == int32(68) {
		goto L963
	} else {
		goto L975
	}
L975:
	;
	if v3806 < int32(0) {
		goto L962
	} else {
		goto L976
	}
L976:
	;
	v3819 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+32))
	if v3819 <= v3806 {
		goto L962
	} else {
		goto L977
	}
L977:
	;
	v3821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3807))))
	if base.B2i32(v3821 == int32(0))|base.B2i32(base.Ui32((v3821-int32(48))&int32(255)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v3821-int32(9)) < base.Ui32(int32(5))) != 0 {
		v3838 = v3806
		v3840 = v3810
		goto L971
	} else {
		goto L978
	}
L978:
	;
	if v3821 != int32(32) {
		goto L961
	} else {
		goto L979
	}
L979:
	;
	v3838 = v3806
	v3840 = v3810
	goto L971
L980:
	;
	goto L970
L981:
	;
	v3859 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v3859 <= int32(0) {
		v3939 = v3745
		goto L982
	} else {
		goto L983
	}
L982:
	;
	v3963 = F_palloc0_mul(m, int32(4), v3939)
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L1
	} else {
		goto L995
	}
L983:
	;
	v3862 = int32(1)
	if v3859 == v3862 {
		v3939 = v3862
		goto L982
	} else {
		goto L984
	}
L984:
	;
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3867 = int32(1)
	v3871 = v3862
	goto L985
L985:
	;
	v3896 = v3865 + v3867<<(uint(int32(2))%32)
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3896)))
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3897)))
	v3901 = *(*int32)(unsafe.Add(mBase, uint32(v3896-int32(4))))
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3901)))
	v3905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3898))))
	v3908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3902))))
	if base.B2i32(v3905 == int32(0))|base.B2i32(v3905 != v3908) != 0 {
		v3926 = v3905
		v3927 = v3908
		goto L988
	} else {
		goto L989
	}
L986:
	;
	v3939 = v3931
	goto L982
L987:
	;
	v3931 = v3871 + base.B2i32(v3926-v3927 != int32(0))
	v3933 = v3867 + int32(1)
	if v3933 != v3859 {
		v3867 = v3933
		v3871 = v3931
		goto L985
	} else {
		goto L994
	}
L988:
	;
	goto L987
L989:
	;
	v3911 = v3898
	v3912 = v3902
	goto L990
L990:
	;
	v3915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3912)+1)))
	v3916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3911)+1)))
	if v3916 == int32(0) {
		v3926 = v3916
		v3927 = v3915
		goto L988
	} else {
		goto L992
	}
L991:
	;
	v3926 = v3916
	v3927 = v3915
	goto L988
L992:
	;
	v3919 = int32(1)
	if v3916 == v3915 {
		v3911 = v3911 + v3919
		v3912 = v3912 + v3919
		goto L990
	} else {
		goto L993
	}
L993:
	;
	goto L991
L994:
	;
	goto L986
L995:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+24)) = v3963
	v3966 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v3966 <= int32(0) {
		v4282 = v3966
		goto L996
	} else {
		goto L997
	}
L996:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+28)) = v3939
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+32)) = v3939
	v4311 = v4282
	goto L964
L997:
	;
	v3969 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3969)))
	v3971 = *(*int32)(unsafe.Add(mBase, uint32(v3970)))
	v3972 = F_strlen(m, v3971)
	mBase = m.M
	v3974 = v3972 + int32(1)
	if base.Ui32(v3974) <= base.Ui32(int32(1024)) {
		goto L999
	} else {
		goto L1000
	}
L998:
	;
	if (v3971^v3996)&int32(3) != 0 {
		goto L1011
	} else {
		goto L1012
	}
L999:
	;
	v3980 = (v3972 + int32(8)) & int32(4088)
	v3981 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+84))
	if base.Ui32(v3980) <= base.Ui32(v3981) {
		goto L1003
	} else {
		goto L1004
	}
L1000:
	;
	goto L1001
L1001:
	;
	v3994 = F_palloc0(m, v3974)
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+84)) = v3988 - v3980
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+80)) = v3989 + v3980
	v3996 = v3989
	goto L998
L1003:
	;
	v3983 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+80))
	v3988 = v3981
	v3989 = v3983
	goto L1002
L1004:
	;
	goto L1005
L1005:
	;
	v3984 = int32(_a_F_dispell_init_1)
	v3986 = F_palloc0(m, v3984)
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1006:
	;
	v3988 = v3984
	v3989 = v3986
	goto L1002
L1007:
	;
	v3996 = v3994
	goto L998
L1008:
	;
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4073))) = v3996
	v4075 = int32(0)
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v4076)))
	*(*int32)(unsafe.Add(mBase, uint32(v4077))) = v4075
	v4080 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4081 = *(*int32)(unsafe.Add(mBase, uint32(v4080)))
	v4084 = F_strlen(m, v4081+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4081)+4)) = v4084
	v4086 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v4086 < int32(2) {
		v4282 = v4086
		goto L996
	} else {
		goto L1029
	}
L1009:
	;
	goto L1008
L1010:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v4053))) = uint8(v4052)
	if v4052&int32(255) == int32(0) {
		goto L1009
	} else {
		goto L1025
	}
L1011:
	;
	v4004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3971))))
	v4051 = v3971
	v4052 = v4004
	v4053 = v3996
	goto L1010
L1012:
	;
	goto L1013
L1013:
	;
	if v3971&int32(3) != 0 {
		goto L1014
	} else {
		goto L1015
	}
L1014:
	;
	v4008 = v3971
	v4010 = v3996
	goto L1017
L1015:
	;
	v4022 = v3971
	v4024 = v3996
	goto L1016
L1016:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v4022)))
	v4029 = int32(-2139062144)
	if (int32(16843008)-v4026|v4026)&v4029 != v4029 {
		v4051 = v4022
		v4052 = v4026
		v4053 = v4024
		goto L1010
	} else {
		goto L1021
	}
L1017:
	;
	v4011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4008))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4010))) = uint8(v4011)
	if v4011 == int32(0) {
		goto L1009
	} else {
		goto L1019
	}
L1018:
	;
	v4022 = v4018
	v4024 = v4016
	goto L1016
L1019:
	;
	v4015 = int32(1)
	v4016 = v4010 + v4015
	v4018 = v4008 + v4015
	if v4018&int32(3) != 0 {
		v4008 = v4018
		v4010 = v4016
		goto L1017
	} else {
		goto L1020
	}
L1020:
	;
	goto L1018
L1021:
	;
	v4034 = v4022
	v4035 = v4026
	v4036 = v4024
	goto L1022
L1022:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4036))) = v4035
	v4038 = int32(4)
	v4039 = v4036 + v4038
	v4041 = v4034 + v4038
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v4034)+4))
	v4046 = int32(-2139062144)
	if (int32(16843008)-v4043|v4043)&v4046 == v4046 {
		v4034 = v4041
		v4035 = v4043
		v4036 = v4039
		goto L1022
	} else {
		goto L1024
	}
L1023:
	;
	v4051 = v4041
	v4052 = v4043
	v4053 = v4039
	goto L1010
L1024:
	;
	goto L1023
L1025:
	;
	v4060 = v4051
	v4062 = v4053
	goto L1026
L1026:
	;
	v4063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4060)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4062)+1)) = uint8(v4063)
	v4065 = int32(1)
	if v4063 != 0 {
		v4060 = v4060 + v4065
		v4062 = v4062 + v4065
		goto L1026
	} else {
		goto L1028
	}
L1027:
	;
	goto L1009
L1028:
	;
	goto L1027
L1029:
	;
	v4091 = int32(1)
	v4093 = v4075
	goto L1030
L1030:
	;
	v4117 = int32(2)
	v4118 = v4091 << (uint(v4117) % 32)
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v4118+v4119)))
	v4122 = *(*int32)(unsafe.Add(mBase, uint32(v4121)))
	v4123 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+24))
	v4127 = *(*int32)(unsafe.Add(mBase, uint32(v4123+v4093<<(uint(v4117)%32))))
	v4130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4122))))
	v4133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4127))))
	if base.B2i32(v4130 == int32(0))|base.B2i32(v4130 != v4133) != 0 {
		v4151 = v4130
		v4152 = v4133
		goto L1033
	} else {
		goto L1034
	}
L1031:
	;
	v4282 = v4280
	goto L996
L1032:
	;
	if v4151-v4152 != 0 {
		goto L1039
	} else {
		goto L1040
	}
L1033:
	;
	goto L1032
L1034:
	;
	v4136 = v4122
	v4137 = v4127
	goto L1035
L1035:
	;
	v4140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4137)+1)))
	v4141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4136)+1)))
	if v4141 == int32(0) {
		v4151 = v4141
		v4152 = v4140
		goto L1033
	} else {
		goto L1037
	}
L1036:
	;
	v4151 = v4141
	v4152 = v4140
	goto L1033
L1037:
	;
	v4144 = int32(1)
	if v4141 == v4140 {
		v4136 = v4136 + v4144
		v4137 = v4137 + v4144
		goto L1035
	} else {
		goto L1038
	}
L1038:
	;
	goto L1036
L1039:
	;
	v4154 = F_strlen(m, v4122)
	mBase = m.M
	v4156 = v4154 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v4156) {
		goto L1043
	} else {
		goto L1044
	}
L1040:
	;
	v4266 = v4093
	v4269 = v4121
	goto L1041
L1041:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4269))) = v4266
	v4271 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4273 = *(*int32)(unsafe.Add(mBase, uint32(v4271+v4118)))
	v4276 = F_strlen(m, v4273+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4273)+4)) = v4276
	v4279 = v4091 + int32(1)
	v4280 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	if v4279 < v4280 {
		v4091 = v4279
		v4093 = v4266
		goto L1030
	} else {
		goto L1073
	}
L1042:
	;
	if (v4122^v4179)&int32(3) != 0 {
		goto L1055
	} else {
		goto L1056
	}
L1043:
	;
	v4159 = F_palloc0(m, v4156)
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		goto L1
	} else {
		goto L1046
	}
L1044:
	;
	goto L1045
L1045:
	;
	v4164 = (v4154 + int32(8)) & int32(4088)
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+84))
	if base.Ui32(v4164) <= base.Ui32(v4165) {
		goto L1048
	} else {
		goto L1049
	}
L1046:
	;
	v4179 = v4159
	goto L1042
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+84)) = v4172 - v4164
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+80)) = v4164 + v4173
	v4179 = v4173
	goto L1042
L1048:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+80))
	v4172 = v4165
	v4173 = v4167
	goto L1047
L1049:
	;
	goto L1050
L1050:
	;
	v4168 = int32(_a_F_dispell_init_1)
	v4170 = F_palloc0(m, v4168)
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1051:
	;
	v4172 = v4168
	v4173 = v4170
	goto L1047
L1052:
	;
	v4255 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+24))
	v4257 = v4093 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4255+v4257<<(uint(int32(2))%32)))) = v4179
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4264 = *(*int32)(unsafe.Add(mBase, uint32(v4262+v4118)))
	v4266 = v4257
	v4269 = v4264
	goto L1041
L1053:
	;
	goto L1052
L1054:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v4235))) = uint8(v4234)
	if v4234&int32(255) == int32(0) {
		goto L1053
	} else {
		goto L1069
	}
L1055:
	;
	v4186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4122))))
	v4233 = v4122
	v4234 = v4186
	v4235 = v4179
	goto L1054
L1056:
	;
	goto L1057
L1057:
	;
	if v4122&int32(3) != 0 {
		goto L1058
	} else {
		goto L1059
	}
L1058:
	;
	v4190 = v4122
	v4192 = v4179
	goto L1061
L1059:
	;
	v4204 = v4122
	v4206 = v4179
	goto L1060
L1060:
	;
	v4208 = *(*int32)(unsafe.Add(mBase, uint32(v4204)))
	v4211 = int32(-2139062144)
	if (int32(16843008)-v4208|v4208)&v4211 != v4211 {
		v4233 = v4204
		v4234 = v4208
		v4235 = v4206
		goto L1054
	} else {
		goto L1065
	}
L1061:
	;
	v4193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4190))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4192))) = uint8(v4193)
	if v4193 == int32(0) {
		goto L1053
	} else {
		goto L1063
	}
L1062:
	;
	v4204 = v4200
	v4206 = v4198
	goto L1060
L1063:
	;
	v4197 = int32(1)
	v4198 = v4192 + v4197
	v4200 = v4190 + v4197
	if v4200&int32(3) != 0 {
		v4190 = v4200
		v4192 = v4198
		goto L1061
	} else {
		goto L1064
	}
L1064:
	;
	goto L1062
L1065:
	;
	v4216 = v4204
	v4217 = v4208
	v4218 = v4206
	goto L1066
L1066:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4218))) = v4217
	v4220 = int32(4)
	v4221 = v4218 + v4220
	v4223 = v4216 + v4220
	v4225 = *(*int32)(unsafe.Add(mBase, uint32(v4216)+4))
	v4228 = int32(-2139062144)
	if (int32(16843008)-v4225|v4225)&v4228 == v4228 {
		v4216 = v4223
		v4217 = v4225
		v4218 = v4221
		goto L1066
	} else {
		goto L1068
	}
L1067:
	;
	v4233 = v4223
	v4234 = v4225
	v4235 = v4221
	goto L1054
L1068:
	;
	goto L1067
L1069:
	;
	v4242 = v4233
	v4244 = v4235
	goto L1070
L1070:
	;
	v4245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4242)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4244)+1)) = uint8(v4245)
	v4247 = int32(1)
	if v4245 != 0 {
		v4242 = v4242 + v4247
		v4244 = v4244 + v4247
		goto L1070
	} else {
		goto L1072
	}
L1071:
	;
	goto L1053
L1072:
	;
	goto L1071
L1073:
	;
	goto L1031
L1074:
	;
	v4343 = int32(0)
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+72))
	v4346 = F_mkSPNode(m, v3713, v4343, v4344, v4343)
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1075:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+20)) = v4346
	m.G0 = v3749 + int32(48)
	goto L960
L1076:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	v4359 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4363 = *(*int32)(unsafe.Add(mBase, uint32(v4359+v3758<<(uint(int32(2))%32))))
	v4364 = *(*int32)(unsafe.Add(mBase, uint32(v4363)))
	*(*int32)(unsafe.Add(mBase, uint32(v3749))) = v4364
	F_errmsg(m, int32(_a_F_dispell_init_47), v3749)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1835), int32(_a_F_dispell_init_60))
	mBase = m.M
	v4373 = m.ExcPending
	if v4373 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1080:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v4380 = m.ExcPending
	if v4380 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	v4381 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4385 = *(*int32)(unsafe.Add(mBase, uint32(v4381+v3758<<(uint(int32(2))%32))))
	v4386 = *(*int32)(unsafe.Add(mBase, uint32(v4385)))
	*(*int32)(unsafe.Add(mBase, uint32(v3749)+16)) = v4386
	F_errmsg(m, int32(_a_F_dispell_init_47), v3749+int32(16))
	mBase = m.M
	v4392 = m.ExcPending
	if v4392 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1082:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1840), int32(_a_F_dispell_init_60))
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1084:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v4404 = m.ExcPending
	if v4404 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1085:
	;
	v4405 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(v4405+v3758<<(uint(int32(2))%32))))
	v4410 = *(*int32)(unsafe.Add(mBase, uint32(v4409)))
	*(*int32)(unsafe.Add(mBase, uint32(v3749)+32)) = v4410
	F_errmsg(m, int32(_a_F_dispell_init_47), v3749+int32(32))
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	F_errfinish(m, int32(_a_F_dispell_init_11), int32(1845), int32(_a_F_dispell_init_60))
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1088:
	;
	if int32(2) <= v4427 {
		goto L1091
	} else {
		goto L1092
	}
L1089:
	;
	goto L1090
L1090:
	;
	m.G0 = v4425 + int32(1040)
	v4878 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+64))
	F_MemoryContextDelete(m, v4878)
	mBase = m.M
	v4880 = m.ExcPending
	if v4880 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1091:
	;
	v4431 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+8))
	F_pg_qsort(m, v4431, v4427, int32(24), int32(1269))
	mBase = m.M
	v4435 = m.ExcPending
	if v4435 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1092:
	;
	v4437 = v4427
	goto L1093
L1093:
	;
	v4440 = F_palloc_mul(m, int32(12), v4437+int32(1))
	mBase = m.M
	v4441 = m.ExcPending
	if v4441 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1094:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+4))
	v4437 = v4436
	goto L1093
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+40)) = v4440
	*(*int32)(unsafe.Add(mBase, uint32(v4440))) = int32(0)
	v4445 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+4))
	if v4445 != 0 {
		goto L1096
	} else {
		goto L1097
	}
L1096:
	;
	v4449 = v4427
	v4450 = v4440
	v4461 = int32(0)
	goto L1099
L1097:
	;
	v4797 = v4427
	v4798 = v4440
	goto L1098
L1098:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4798))) = int32(0)
	v4823 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+40))
	v4827 = F_repalloc(m, v4823, v4798-v4823+int32(12))
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		goto L1
	} else {
		goto L1146
	}
L1099:
	;
	if base.Ui32(v4461) < base.Ui32(v4449) {
		goto L1101
	} else {
		goto L1102
	}
L1100:
	;
	v4797 = v4789
	v4798 = v4766
	goto L1098
L1101:
	;
	v4474 = v4461
	goto L1103
L1102:
	;
	v4474 = v4449
	goto L1103
L1103:
	;
	v4475 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+8))
	v4478 = v4475 + v4461*int32(24)
	v4479 = *(*int32)(unsafe.Add(mBase, uint32(v4478)+4))
	v4484 = int32(0)
	if base.B2i32(v4479&int32(28) == v4484)|base.B2i32(v4479&int32(16776192) == v4484) != 0 {
		v4766 = v4450
		goto L1104
	} else {
		goto L1105
	}
L1104:
	;
	if v4479&int32(1) != 0 {
		goto L1142
	} else {
		goto L1143
	}
L1105:
	;
	v4491 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+32))
	if v4491 <= int32(0) {
		v4766 = v4450
		goto L1104
	} else {
		goto L1106
	}
L1106:
	;
	v4494 = *(*int32)(unsafe.Add(mBase, uint32(v4478)))
	v4501 = int32(0)
	goto L1109
L1107:
	;
	v4751 = *(*int32)(unsafe.Add(mBase, uint32(v4478)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v4450))) = v4751
	v4753 = *(*int32)(unsafe.Add(mBase, uint32(v4478)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v4450)+8)) = uint8(v4620)
	*(*int32)(unsafe.Add(mBase, uint32(v4450)+4)) = int32(base.Ui32(v4753)>>(uint(int32(10))%32)) & int32(_a_F_dispell_init_61)
	v4766 = v4450 + int32(12)
	goto L1104
L1108:
	;
	if base.B2i32(v4694 == int32(0))|base.B2i32(v4698 == v4699) != 0 {
		v4766 = v4450
		goto L1104
	} else {
		goto L1141
	}
L1109:
	;
	v4523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4494))))
	if v4523 != 0 {
		goto L1113
	} else {
		goto L1114
	}
L1110:
	;
	if v4632 < int32(0) {
		goto L1132
	} else {
		goto L1133
	}
L1111:
	;
	goto L1110
L1112:
	;
	v4643 = v4501 + int32(1)
	v4644 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+32))
	if v4643 < v4644 {
		v4501 = v4643
		goto L1109
	} else {
		goto L1131
	}
L1113:
	;
	v4524 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+24))
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v4524+v4501<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4425)+1036)) = v4528
	goto L1116
L1114:
	;
	goto L1115
L1115:
	;
	v4618 = *(*int32)(unsafe.Add(mBase, uint32(v4478)+4))
	v4620 = v4618 & int32(1)
	v4621 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+40))
	if v4450 == v4621 {
		goto L1107
	} else {
		goto L1128
	}
L1116:
	;
	v4557 = *(*int32)(unsafe.Add(mBase, uint32(v4425)+1036))
	v4558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4557))))
	if v4558 == int32(0) {
		goto L1112
	} else {
		goto L1118
	}
L1117:
	;
	goto L1115
L1118:
	;
	F_getNextFlagFromString(m, v3713, v4425+int32(1036), v4425)
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1119:
	;
	v4567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4425))))
	v4570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4494))))
	if base.B2i32(v4567 == int32(0))|base.B2i32(v4567 != v4570) != 0 {
		v4588 = v4567
		v4589 = v4570
		goto L1121
	} else {
		goto L1122
	}
L1120:
	;
	if v4588-v4589 != 0 {
		goto L1116
	} else {
		goto L1127
	}
L1121:
	;
	goto L1120
L1122:
	;
	v4573 = v4425
	v4574 = v4494
	goto L1123
L1123:
	;
	v4577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4574)+1)))
	v4578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4573)+1)))
	if v4578 == int32(0) {
		v4588 = v4578
		v4589 = v4577
		goto L1121
	} else {
		goto L1125
	}
L1124:
	;
	v4588 = v4578
	v4589 = v4577
	goto L1121
L1125:
	;
	v4581 = int32(1)
	if v4578 == v4577 {
		v4573 = v4573 + v4581
		v4574 = v4574 + v4581
		goto L1123
	} else {
		goto L1126
	}
L1126:
	;
	goto L1124
L1127:
	;
	goto L1117
L1128:
	;
	v4625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4450-int32(4)))))
	if v4620 != v4625 {
		goto L1107
	} else {
		goto L1129
	}
L1129:
	;
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v4450-int32(12))))
	v4630 = F_strlen(m, v4629)
	mBase = m.M
	v4631 = int32(1)
	v4632 = v4630 - v4631
	v4633 = *(*int32)(unsafe.Add(mBase, uint32(v4478)+12))
	v4634 = F_strlen(m, v4633)
	mBase = m.M
	v4636 = v4634 - v4631
	v4639 = *(*int32)(unsafe.Add(mBase, uint32(v4450-int32(8))))
	if int32(0) < v4639 {
		goto L1111
	} else {
		goto L1130
	}
L1130:
	;
	v4694 = v4639
	v4698 = v4636
	v4699 = v4632
	goto L1108
L1131:
	;
	v4766 = v4450
	goto L1104
L1132:
	;
	v4694 = v4639
	v4698 = v4636
	v4699 = v4632
	goto L1108
L1133:
	;
	goto L1134
L1134:
	;
	if v4636 < int32(0) {
		v4694 = v4639
		v4698 = v4636
		v4699 = v4632
		goto L1108
	} else {
		goto L1135
	}
L1135:
	;
	v4650 = v4639
	v4655 = v4636
	v4656 = v4632
	goto L1136
L1136:
	;
	v4678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4656+v4629))))
	v4680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4655+v4633))))
	if v4678 != v4680 {
		goto L1107
	} else {
		goto L1138
	}
L1137:
	;
	v4694 = v4683
	v4698 = v4685
	v4699 = v4687
	goto L1108
L1138:
	;
	v4682 = int32(1)
	v4683 = v4650 - v4682
	v4685 = v4655 - v4682
	v4687 = v4656 - v4682
	if v4685|v4687 < int32(0) {
		v4694 = v4683
		v4698 = v4685
		v4699 = v4687
		goto L1108
	} else {
		goto L1139
	}
L1139:
	;
	if int32(1) < v4650 {
		v4650 = v4683
		v4655 = v4685
		v4656 = v4687
		goto L1136
	} else {
		goto L1140
	}
L1140:
	;
	goto L1137
L1141:
	;
	goto L1107
L1142:
	;
	v4789 = v4474
	goto L1144
L1143:
	;
	v4789 = v4449
	goto L1144
L1144:
	;
	v4791 = v4461 + int32(1)
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+4))
	if base.Ui32(v4791) < base.Ui32(v4792) {
		v4449 = v4789
		v4450 = v4766
		v4461 = v4791
		goto L1099
	} else {
		goto L1145
	}
L1145:
	;
	goto L1100
L1146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+40)) = v4827
	v4830 = int32(0)
	v4833 = F_mkANode(m, v3713, v4830, v4797, v4830, v4830)
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+16)) = v4833
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+4))
	v4839 = F_mkANode(m, v3713, v4797, v4836, int32(0), int32(1))
	mBase = m.M
	v4840 = m.ExcPending
	if v4840 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+12)) = v4839
	F_mkVoidAffix(m, v3713, int32(1), v4797)
	mBase = m.M
	v4844 = m.ExcPending
	if v4844 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	F_mkVoidAffix(m, v3713, int32(0), v4797)
	mBase = m.M
	v4847 = m.ExcPending
	if v4847 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	goto L1090
L1151:
	;
	v4881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+80)) = v4881
	*(*int64)(unsafe.Add(mBase, uint32(v3713)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3713)+52)) = v4881
	m.G0 = v3735 + int32(16)
	return base.I64_extend_i32_u(v3737)
L1152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1153:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4900 = m.ExcPending
	if v4900 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1154:
	;
	F_errmsg(m, int32(_a_F_dispell_init_62), int32(0))
	mBase = m.M
	v4904 = m.ExcPending
	if v4904 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1155:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(109), int32(_a_F_dispell_init_56))
	mBase = m.M
	v4909 = m.ExcPending
	if v4909 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1157:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4943 = m.ExcPending
	if v4943 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1158:
	;
	F_errmsg(m, int32(_a_F_dispell_init_63), int32(0))
	mBase = m.M
	v4947 = m.ExcPending
	if v4947 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	F_errfinish(m, int32(_a_F_dispell_init_55), int32(103), int32(_a_F_dispell_init_56))
	mBase = m.M
	v4952 = m.ExcPending
	if v4952 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
