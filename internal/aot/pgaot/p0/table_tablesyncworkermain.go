package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_TableSyncWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v464 int64
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int64
	_ = v702
	var v705 int32
	_ = v705
	var v708 int64
	_ = v708
	var v711 int64
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v806 int32
	_ = v806
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int64
	_ = v885
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int64
	_ = v918
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1065 int32
	_ = v1065
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1080 int64
	_ = v1080
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1108 int32
	_ = v1108
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1332 int32
	_ = v1332
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1391 int32
	_ = v1391
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1424 int32
	_ = v1424
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1444 int32
	_ = v1444
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1589 int64
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1605 int64
	_ = v1605
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int64
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1650 int32
	_ = v1650
	var v1661 int32
	_ = v1661
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1748 int32
	_ = v1748
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1791 int32
	_ = v1791
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1841 int32
	_ = v1841
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1873 int32
	_ = v1873
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2018 int32
	_ = v2018
	var v2024 int32
	_ = v2024
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2126 int32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2171 int32
	_ = v2171
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2238 int32
	_ = v2238
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2253 int32
	_ = v2253
	var v2259 int32
	_ = v2259
	var v2297 int32
	_ = v2297
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2318 int32
	_ = v2318
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2466 int32
	_ = v2466
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2544 int32
	_ = v2544
	var v2555 int32
	_ = v2555
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2597 int32
	_ = v2597
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2659 int32
	_ = v2659
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2672 int64
	_ = v2672
	var v2679 int32
	_ = v2679
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2709 int32
	_ = v2709
	var v2715 int32
	_ = v2715
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2776 int32
	_ = v2776
	var v2780 int32
	_ = v2780
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2830 int32
	_ = v2830
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2840 int64
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2847 int32
	_ = v2847
	var v2854 int32
	_ = v2854
	var v2860 int32
	_ = v2860
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2869 int32
	_ = v2869
	var v2871 int32
	_ = v2871
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2885 int32
	_ = v2885
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2903 int32
	_ = v2903
	var v2912 int32
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2933 int32
	_ = v2933
	var v2939 int32
	_ = v2939
	var v2949 int32
	_ = v2949
	var v2956 int32
	_ = v2956
	var v2962 int32
	_ = v2962
	var v2964 int32
	_ = v2964
	var v2965 int64
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2975 int32
	_ = v2975
	var v2978 int32
	_ = v2978
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v3010 int32
	_ = v3010
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int64
	_ = v3019
	var v3026 int64
	_ = v3026
	var v3035 int32
	_ = v3035
	var v3044 int32
	_ = v3044
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3064 int64
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3099 int32
	_ = v3099
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3116 int32
	_ = v3116
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3130 int32
	_ = v3130
	var v3138 int32
	_ = v3138
	var v3142 int32
	_ = v3142
	var v3150 int32
	_ = v3150
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3163 int32
	_ = v3163
	var v3165 int32
	_ = v3165
	var v3174 int32
	_ = v3174
	var v3181 int32
	_ = v3181
	var v3186 int32
	_ = v3186
	var v3191 int32
	_ = v3191
	var v3200 int32
	_ = v3200
	var v3204 int32
	_ = v3204
	var v3210 int32
	_ = v3210
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3231 int32
	_ = v3231
	var v3237 int32
	_ = v3237
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3258 int32
	_ = v3258
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3282 int32
	_ = v3282
	var v3289 int32
	_ = v3289
	var v3319 int32
	_ = v3319
	var v3320 int64
	_ = v3320
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3365 int64
	_ = v3365
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3383 int32
	_ = v3383
	var v3386 int32
	_ = v3386
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3397 int32
	_ = v3397
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3404 int32
	_ = v3404
	var v3407 int32
	_ = v3407
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3427 int32
	_ = v3427
	var v3429 int32
	_ = v3429
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3436 int32
	_ = v3436
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3444 int64
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3451 int32
	_ = v3451
	v2 = int32(0)
	F_SetupApplyOrSyncWorker(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v33 = m.G0
	v35 = v33 - int32(128)
	m.G0 = v35
	*(*int64)(unsafe.Add(mBase, uint32(v35)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = int32(0)
	v42 = v35 + int32(56)
	v45 = m.G0
	v47 = v45 - int32(784)
	m.G0 = v47
	v52 = v2
	v54 = v2
	v55 = v2
	v56 = v2
	v57 = v2
	v58 = int32(-1)
	v72 = v2
	v73 = v2
	goto L3
L3:
	;
	if v58 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v47 + int32(784)
	v3346 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	v3349 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(v3349)+36))
	v3352 = v35 - int32(-64)
	F_ReplicationOriginNameForLogicalRep(m, v3347, v3350, v3352)
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L1
	} else {
		goto L567
	}
L5:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[2]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[3]))
	v87 = v47 + int32(384)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v47 + int32(380)
	goto L8
L6:
	;
	v93 = v57
	v94 = v72
	v95 = v73
	goto L7
L7:
	;
	goto L9
L8:
	;
	v93 = int32(0)
	v94 = v83
	v95 = v85
	goto L7
L9:
	;
	if v93 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	goto L4
L11:
	;
	v3319 = int32(m.ExcTag)
	v3320 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3319 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[2])) = v94
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[3])) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v3244
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v3246
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v3247
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v3248)
	v3280 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[4]))
	v3281 = F_MemoryContextStrdup(m, v3280, v3258)
	mBase = m.M
	v3282 = m.ExcPending
	if v3282 != 0 {
		goto L11
	} else {
		goto L555
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L11
	} else {
		goto L511
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2579 = v47 + int32(648)
	F_appendStringInfoString(m, v2579, v2555)
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L11
	} else {
		goto L436
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v47+int32(648), int32(_a_F_TableSyncWorkerMain_0))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L11
	} else {
		goto L406
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoChar(m, v47+int32(648), int32(41))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L11
	} else {
		goto L405
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[3])) = v47 + int32(384)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_StartTransactionCommand(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[2])) = v94
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[3])) = v95
	v2228 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v2229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2228)+37)))
	if v2229 == int32(1) {
		goto L398
	} else {
		goto L399
	}
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+36))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v118 = F_GetSubscriptionRelState(m, v111, v110, v47+int32(632))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+38)))
	if v129 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+32)))
	v135 = v132 ^ int32(1)
	goto L25
L24:
	;
	v135 = int32(0)
	goto L25
L25:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v140 = base.AtomicRmwXchg32(m, v137, int32(56), int32(1))
	if v140 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_s_lock(m, v137+int32(56), int32(_a_F_TableSyncWorkerMain_1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+40)) = uint8(v118)
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v47)+632))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+48)) = v153
	v155 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v151)+56)), uint32(v155))
	v159 = v118 & int32(255)
	if v159 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v165 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v159-int32(114)))
	goto L32
L31:
	;
	v165 = v155
	goto L32
L32:
	;
	if v165 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_FinishSyncWorker(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v179 = F_palloc(m, int32(64))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L37
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+36))
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[5]))
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	*(*int64)(unsafe.Add(mBase, uint32(v47)+360)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v47)+356)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v47)+352)) = v186
	v205 = F_pg_snprintf(m, v179, int32(64), int32(_a_F_TableSyncWorkerMain_2), v47+int32(352))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[7]))
	v217 = int32(1)
	v223 = m.T0[v209].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v216, v217, v217, v135&v217, v179, v47+int32(640))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8])) = v223
	if v223 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_ReplicationOriginNameForLogicalRep(m, v270, v267, v47+int32(560))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L48
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v47)+640))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v245
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_3), v47)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1320), int32(_a_F_TableSyncWorkerMain_5))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+40)))
	switch v281 - int32(100) {
	case 0:
		goto L52
	default:
		v295 = v280
		goto L51
	case 2:
		goto L50
	}
L49:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	if v467 != 0 {
		goto L79
	} else {
		goto L80
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_StartTransactionCommand(m)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L11
	} else {
		goto L75
	}
L51:
	;
	v298 = base.AtomicRmwXchg32(m, v295, int32(56), int32(1))
	if v298 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	F_ReplicationSlotDropAtPubNode(m, v289, v179, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v295 = v294
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_s_lock(m, v295+int32(56), int32(_a_F_TableSyncWorkerMain_1))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L11
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+48)) = int64(0)
	v312 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v309)+40)) = uint8(v312)
	v314 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v309)+56)), uint32(v314))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_StartTransactionCommand(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L11
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v324)+48))
	v326 = int32(*(*int8)(unsafe.Add(mBase, uint32(v324)+40)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v324)+36))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_UpdateSubscriptionRelState(m, v328, v327, v326, v325, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v341 = v47 + int32(560)
	v343 = F_replorigin_by_name(m, v341, int32(1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	if v343 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v351 = F_replorigin_create(m, v341)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	v353 = v56
	v354 = v343
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L65
	}
L64:
	;
	v353 = v351
	v354 = v351
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v366 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_StartTransactionCommand(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v382 = F_table_open(m, v376, int32(3))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v394 = int32(0)
	v396 = m.T0[v386].(func(*base.Module, int32, int32, int32, int32) int32)(m, v392, int32(_a_F_TableSyncWorkerMain_6), v394, v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	if v398 == int32(1) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+336)) = v416
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_7), v47+int32(336))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1420), int32(_a_F_TableSyncWorkerMain_5))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v449 = F_replorigin_by_name(m, v47+int32(560), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_replorigin_session_setup(m, v449, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[9])) = uint16(v449)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	v464 = F_replorigin_session_get_progress(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v464
	v2978 = v52
	v2980 = v54
	v2981 = v55
	v2982 = v56
	goto L13
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v467)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L11
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	if v474 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v474)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
	if v481 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v481)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v396)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L11
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+40)))
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v506 = int32(0)
	v509 = m.T0[v499].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v505, v179, v506, v506, v496, int32(2), v42)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_LockRelationOid(m, int32(_a_F_TableSyncWorkerMain_8), int32(3))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v525 = int32(1)
	F_replorigin_advance(m, v354, v519, int64(0), v525, v525)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_UnlockRelationOid(m, int32(_a_F_TableSyncWorkerMain_8), int32(3))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_replorigin_session_setup(m, v354, int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[9])) = uint16(v354)
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+39)))
	if v548 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_SwitchToUntrustedUser(m, v552, v47+int32(548))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L11
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v382)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v573 = F_pg_class_aclcheck(m, v561, v567, int64(1))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L11
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	if v573 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v576 = int32(*(*int8)(unsafe.Add(mBase, uint32(v575)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	switch v576 - int32(73) {
	case 0, 32:
		v590 = int32(20)
		goto L106
	default:
		goto L107
	case 10:
		goto L111
	case 29:
		goto L108
	case 36:
		goto L109
	case 45:
		goto L110
	}
L103:
	;
	goto L104
L104:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v382)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v609 = int32(0)
	v611 = F_check_enable_rls(m, v604, v609, v609)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L11
	} else {
		goto L113
	}
L105:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_aclcheck_error(m, v573, v592, v593+int32(4))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L11
	} else {
		goto L112
	}
L106:
	;
	v592 = v590
	goto L105
L107:
	;
	v590 = int32(42)
	goto L106
L108:
	;
	v592 = int32(18)
	goto L105
L109:
	;
	v592 = int32(23)
	goto L105
L110:
	;
	v592 = int32(52)
	goto L105
L111:
	;
	v592 = int32(38)
	goto L105
L112:
	;
	goto L104
L113:
	;
	if v611 == int32(2) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L11
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v670 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L11
	} else {
		goto L122
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(1088))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v641 = F_GetUserNameFromId(m, v635, int32(1))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v643 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v641
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_9), v47+int32(16))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1481), int32(_a_F_TableSyncWorkerMain_5))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_PushActiveSnapshot(m, v670)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v687 = m.T0[v680].(func(*base.Module, int32) int32)(m, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v695 = F_get_namespace_name(m, v690)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+744)) = v699
	v702 = *(*int64)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+736)) = v702
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+720)) = v705
	v708 = *(*int64)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+712)) = v708
	v711 = *(*int64)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+704)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v47)+700)) = int32(25)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v724 = m.T0[v717].(func(*base.Module, int32) int32)(m, v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	v727 = v697 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+672)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+668)) = v695
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v735 = v47 + int32(752)
	F_initStringInfo(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v742 = F_quote_literal_cstr(m, v695)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v748 = F_quote_literal_cstr(m, v727)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+324)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v47)+320)) = v742
	F_appendStringInfo(m, v735, int32(_a_F_TableSyncWorkerMain_10), v47+int32(320))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v762)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v769 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v47)+752))
	v774 = m.T0[v763].(func(*base.Module, int32, int32, int32, int32) int32)(m, v769, v770, int32(3), v47+int32(736))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	if v776 != int32(2) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L11
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v774)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v822 = F_MakeSingleTupleTableSlot(m, v816, int32(_a_F_TableSyncWorkerMain_11))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L11
	} else {
		goto L139
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v774)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+312)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v47)+308)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+304)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_12), v47+int32(304))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L11
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(760), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L11
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v774)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v831 = F_tuplestore_gettupleslot(m, v824, int32(1), int32(0), v822)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	if v831 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L11
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v870 = int32(*(*int16)(unsafe.Add(mBase, uint32(v822)+6)))
	if v870 <= int32(0) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+292)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+288)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_14), v47+int32(288))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(767), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v822)+8))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v874].(func(*base.Module, int32, int32))(m, v822, int32(1))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L11
	} else {
		goto L151
	}
L149:
	;
	v883 = v870
	goto L150
L150:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v822)+16))
	v885 = *(*int64)(unsafe.Add(mBase, uint32(v884)))
	*(*uint32)(unsafe.Add(mBase, uint32(v47)+664)) = uint32(v885)
	if base.I32_extend16_s(v883) <= int32(1) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v822)+6)))
	v883 = v882
	goto L150
L152:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v822)+8))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v891].(func(*base.Module, int32, int32))(m, v822, int32(2))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L11
	} else {
		goto L155
	}
L153:
	;
	v901 = v884
	goto L154
L154:
	;
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v901)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+688)) = uint8(v902)
	v904 = int32(*(*int16)(unsafe.Add(mBase, uint32(v822)+6)))
	if v904 <= int32(2) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v822)+16))
	v901 = v899
	goto L154
L156:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v822)+8))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v908].(func(*base.Module, int32, int32))(m, v822, int32(3))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L11
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v822)+16))
	v918 = *(*int64)(unsafe.Add(mBase, uint32(v917)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+689)) = uint8(v918)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_ExecDropSingleTupleTableSlot(m, v822)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L11
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v774)+8))
	if v926 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v926)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L11
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v774)+12))
	if v933 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v933)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L11
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v774)+16))
	if v940 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L167
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v940)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L11
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v774)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L11
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v953 = int32(0)
	v955 = base.B2i32(v724 < int32(_a_F_TableSyncWorkerMain_15))
	if v724 < int32(_a_F_TableSyncWorkerMain_15) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1349 = v47 + int32(752)
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	v1351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1350))) = uint8(v1351)
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+12)) = v1351
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+4)) = v1351
	goto L237
L175:
	;
	v1320 = v55
	v1324 = v953
	v1332 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+696)) = int32(22)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v963 = F_makeStringInfo(m)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	v966 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_GetPublicationsStr(m, v967, v963, int32(1))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L11
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v980 = v47 + int32(752)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)))
	v982 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v981))) = uint8(v982)
	*(*int32)(unsafe.Add(mBase, uint32(v980)+12)) = v982
	*(*int32)(unsafe.Add(mBase, uint32(v980)+4)) = v982
	goto L180
L180:
	;
	if base.Ui32(int32(_a_F_TableSyncWorkerMain_16)) <= base.Ui32(v724) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v47)+752))
	v1033 = m.T0[v1022].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1028, v1029, int32(1), v47+int32(696))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L11
	} else {
		goto L187
	}
L182:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+260)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v47)+256)) = v990
	F_appendStringInfo(m, v980, int32(_a_F_TableSyncWorkerMain_17), v47+int32(256))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L11
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+276)) = v1003
	*(*int32)(unsafe.Add(mBase, uint32(v47)+272)) = v1008
	F_appendStringInfo(m, v47+int32(752), int32(_a_F_TableSyncWorkerMain_18), v47+int32(272))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L11
	} else {
		goto L186
	}
L185:
	;
	goto L181
L186:
	;
	goto L181
L187:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	if v1035 != int32(2) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L11
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1080 = *(*int64)(unsafe.Add(mBase, uint32(v1075)+40))
	if int64(2) <= v1080 {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+248)) = v1053
	*(*int32)(unsafe.Add(mBase, uint32(v47)+244)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+240)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_19), v47+int32(240))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(838), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L11
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L11
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1124 = F_MakeSingleTupleTableSlot(m, v1118, int32(_a_F_TableSyncWorkerMain_11))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L11
	} else {
		goto L202
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(1088))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L11
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_20), v47+int32(32))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L11
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(852), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L11
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1133 = F_tuplestore_gettupleslot(m, v1126, int32(1), int32(0), v1124)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L11
	} else {
		goto L203
	}
L203:
	;
	if v1133 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1124)+6)))
	if v1135 <= int32(0) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v1258 = v55
	v1262 = v953
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_ExecDropSingleTupleTableSlot(m, v1124)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L11
	} else {
		goto L223
	}
L207:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+8))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1139].(func(*base.Module, int32, int32))(m, v1124, int32(1))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L11
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+20))
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148))))
	if v1149 != 0 {
		v1221 = v55
		v1225 = v953
		goto L211
	} else {
		goto L212
	}
L210:
	;
	goto L209
L211:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+8))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1221
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1246].(func(*base.Module, int32))(m, v1124)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L11
	} else {
		goto L222
	}
L212:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+16))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1156 = F_pg_detoast_datum(m, v1151)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L11
	} else {
		goto L213
	}
L213:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+16))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+8))
	if v1159 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+4))
	v1169 = (v1162<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L216
L215:
	;
	v1169 = v1159
	goto L216
L216:
	;
	if v1158 <= int32(0) {
		v1221 = v55
		v1225 = v953
		goto L211
	} else {
		goto L217
	}
L217:
	;
	v1179 = v55
	v1181 = int32(0)
	v1183 = v953
	goto L218
L218:
	;
	v1206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1156+v1169+v1181<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1179
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1211 = F_bms_add_member(m, v1183, v1206)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L11
	} else {
		goto L220
	}
L219:
	;
	v1221 = v1211
	v1225 = v1211
	goto L211
L220:
	;
	v1214 = v1181 + int32(1)
	if v1214 != v1158 {
		v1179 = v1211
		v1181 = v1214
		v1183 = v1211
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v1258 = v1221
	v1262 = v1225
	goto L206
L223:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+8))
	if v1288 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1288)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L11
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+12))
	if v1295 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L226
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v1295)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L11
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+16))
	if v1302 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v1302)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L11
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1033)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L11
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	v1320 = v1258
	v1324 = v1262
	v1332 = v963
	goto L174
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v1349, int32(_a_F_TableSyncWorkerMain_21))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L11
	} else {
		goto L238
	}
L238:
	;
	v1365 = base.B2i32(v724 < int32(_a_F_TableSyncWorkerMain_22))
	if v724 < int32(_a_F_TableSyncWorkerMain_22) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1377 = int32(4)
	goto L241
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v47+int32(752), int32(_a_F_TableSyncWorkerMain_23))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L11
	} else {
		goto L242
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+232)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v47)+224)) = v1382
	if base.Ui32(v724-int32(_a_F_TableSyncWorkerMain_24)) < base.Ui32(int32(_a_F_TableSyncWorkerMain_25)) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1377 = int32(5)
	goto L241
L243:
	;
	v1391 = int32(_a_F_TableSyncWorkerMain_26)
	goto L245
L244:
	;
	v1391 = int32(_a_F_TableSyncWorkerMain_27)
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+228)) = v1391
	F_appendStringInfo(m, v47+int32(752), int32(_a_F_TableSyncWorkerMain_28), v47+int32(224))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L11
	} else {
		goto L246
	}
L246:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1408 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v47)+752))
	v1412 = m.T0[v1402].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1408, v1409, v1377, v47+int32(704))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1412)))
	if v1414 != int32(2) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L11
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1460 = F_palloc0_mul(m, int32(4), int32(1664))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L11
	} else {
		goto L255
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L11
	} else {
		goto L252
	}
L252:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+216)) = v1432
	*(*int32)(unsafe.Add(mBase, uint32(v47)+212)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+208)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_12), v47+int32(208))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L11
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(919), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L11
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+680)) = v1460
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1469 = F_palloc0_mul(m, int32(4), int32(1664))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L11
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+692)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+684)) = v1469
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1480 = F_MakeSingleTupleTableSlot(m, v1474, int32(_a_F_TableSyncWorkerMain_11))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L11
	} else {
		goto L257
	}
L257:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1489 = F_tuplestore_gettupleslot(m, v1482, int32(1), int32(0), v1480)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L11
	} else {
		goto L258
	}
L258:
	;
	v1491 = int32(0)
	if v1489 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1503 = v1491
	v1505 = v1491
	goto L262
L260:
	;
	v1700 = v1491
	v1702 = v1491
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_ExecDropSingleTupleTableSlot(m, v1480)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L11
	} else {
		goto L305
	}
L262:
	;
	v1524 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480)+6)))
	if v1524 <= int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1700 = v1671
	v1702 = v1672
	goto L261
L264:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1528].(func(*base.Module, int32, int32))(m, v1480, int32(1))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L11
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	if v1324 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L266
L268:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1676].(func(*base.Module, int32))(m, v1480)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L11
	} else {
		goto L302
	}
L269:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+16))
	v1538 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1537))))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1543 = F_bms_is_member(m, v1538, v1324)
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L11
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480)+6)))
	if v1547 <= int32(1) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	if v1543 == int32(0) {
		v1671 = v1503
		v1672 = v1505
		goto L268
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1551].(func(*base.Module, int32, int32))(m, v1480, int32(2))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L11
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+16))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1566 = F_text_to_cstring(m, v1561)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L11
	} else {
		goto L278
	}
L277:
	;
	goto L276
L278:
	;
	v1568 = int32(2)
	v1569 = v1503 << (uint(v1568) % 32)
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v47)+680))
	*(*int32)(unsafe.Add(mBase, uint32(v1569+v1570))) = v1566
	v1573 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480)+6)))
	if v1573 <= v1568 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1577].(func(*base.Module, int32, int32))(m, v1480, int32(3))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L11
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v47)+684))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+16))
	v1589 = *(*int64)(unsafe.Add(mBase, uint32(v1588)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v1586+v1569))) = uint32(v1589)
	v1591 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480)+6)))
	if v1591 <= int32(3) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L281
L283:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1595].(func(*base.Module, int32, int32))(m, v1480, int32(4))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L11
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+16))
	v1605 = *(*int64)(unsafe.Add(mBase, uint32(v1604)+24))
	if v1605 != int64(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L285
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v47)+692))
	v1613 = F_bms_add_member(m, v1612, v1503)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L11
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	if (v1505|v1365)&int32(1) != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+692)) = v1613
	goto L289
L291:
	;
	v1638 = v1505 | base.B2i32(int32(_a_F_TableSyncWorkerMain_29) < v724)
	goto L293
L292:
	;
	v1620 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480)+6)))
	if v1620 <= int32(4) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1640 = v1503 + int32(1)
	if v1640 < int32(1664) {
		v1671 = v1640
		v1672 = v1638
		goto L268
	} else {
		goto L298
	}
L294:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+8))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1624].(func(*base.Module, int32, int32))(m, v1480, int32(5))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L11
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1480)+16))
	v1634 = *(*int64)(unsafe.Add(mBase, uint32(v1633)+32))
	v1638 = base.B2i32(v1634 != int64(0))
	goto L293
L297:
	;
	goto L296
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L11
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v695
	F_errmsg_internal(m, int32(_a_F_TableSyncWorkerMain_30), v47+int32(48))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L11
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(967), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L11
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
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1690 = F_tuplestore_gettupleslot(m, v1683, int32(1), int32(0), v1480)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L11
	} else {
		goto L303
	}
L303:
	;
	if v1690 != 0 {
		v1503 = v1671
		v1505 = v1672
		goto L262
	} else {
		goto L304
	}
L304:
	;
	goto L263
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+676)) = v1700
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+8))
	if v1728 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1728)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L11
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+12))
	if v1735 != 0 {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	goto L308
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v1735)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L11
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1412)+16))
	if v1742 != 0 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	goto L312
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v1742)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L11
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1412)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L11
	} else {
		goto L318
	}
L317:
	;
	goto L316
L318:
	;
	v1755 = int32(0)
	if v955 == v1755 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1763 = v47 + int32(752)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1763)))
	v1765 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1764))) = uint8(v1765)
	*(*int32)(unsafe.Add(mBase, uint32(v1763)+12)) = v1765
	*(*int32)(unsafe.Add(mBase, uint32(v1763)+4)) = v1765
	goto L322
L320:
	;
	v2034 = v1755
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v47)+752))
	F_pfree(m, v2064)
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L11
	} else {
		goto L374
	}
L322:
	;
	if base.Ui32(int32(_a_F_TableSyncWorkerMain_16)) <= base.Ui32(v724) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1811 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v47)+752))
	v1816 = m.T0[v1805].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1811, v1812, int32(1), v47+int32(700))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L11
	} else {
		goto L329
	}
L324:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+180)) = v1778
	*(*int32)(unsafe.Add(mBase, uint32(v47)+176)) = v1773
	F_appendStringInfo(m, v1763, int32(_a_F_TableSyncWorkerMain_31), v47+int32(176))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L11
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+196)) = v1786
	*(*int32)(unsafe.Add(mBase, uint32(v47)+192)) = v1791
	F_appendStringInfo(m, v47+int32(752), int32(_a_F_TableSyncWorkerMain_32), v47+int32(192))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L11
	} else {
		goto L328
	}
L327:
	;
	goto L323
L328:
	;
	goto L323
L329:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1816)))
	if v1818 != int32(2) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L11
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1857 = F_MakeSingleTupleTableSlot(m, v1851, int32(_a_F_TableSyncWorkerMain_11))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L11
	} else {
		goto L336
	}
L333:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+168)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v47)+164)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v47)+160)) = v695
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_33), v47+int32(160))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L11
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1031), int32(_a_F_TableSyncWorkerMain_13))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L11
	} else {
		goto L335
	}
L335:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1866 = F_tuplestore_gettupleslot(m, v1859, int32(1), int32(0), v1857)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L11
	} else {
		goto L338
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_ExecDropSingleTupleTableSlot(m, v1857)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L11
	} else {
		goto L359
	}
L338:
	;
	if v1866 == int32(0) {
		v1966 = v1755
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1873 = v1755
	goto L340
L340:
	;
	v1899 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1857)+6)))
	if v1899 <= int32(0) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	v1966 = v1944
	goto L337
L342:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+8))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1902)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1903].(func(*base.Module, int32, int32))(m, v1857, int32(1))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L11
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+20))
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912))))
	if v1913 == int32(1) {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	goto L344
L346:
	;
	if v1873 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	goto L348
L348:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+16))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1932 = F_text_to_cstring(m, v1927)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L11
	} else {
		goto L353
	}
L349:
	;
	v1966 = int32(0)
	goto L337
L350:
	;
	goto L351
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_list_free_deep(m, v1873)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L11
	} else {
		goto L352
	}
L352:
	;
	v1966 = int32(0)
	goto L337
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1938 = F_makeString(m, v1932)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L11
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1944 = F_lappend(m, v1873, v1938)
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L11
	} else {
		goto L355
	}
L355:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+8))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	m.T0[v1947].(func(*base.Module, int32))(m, v1857)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L11
	} else {
		goto L356
	}
L356:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v1961 = F_tuplestore_gettupleslot(m, v1954, int32(1), int32(0), v1857)
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L11
	} else {
		goto L357
	}
L357:
	;
	if v1961 != 0 {
		v1873 = v1944
		goto L340
	} else {
		goto L358
	}
L358:
	;
	goto L341
L359:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	if v1998 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1998)
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L11
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+12))
	if v2005 != 0 {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	goto L362
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v2005)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L11
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+16))
	if v2012 != 0 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	goto L366
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v2012)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L11
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v1816)
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L11
	} else {
		goto L372
	}
L371:
	;
	goto L370
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_free_attrmap(m, v1332)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L11
	} else {
		goto L373
	}
L373:
	;
	v2034 = v1966
	goto L321
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_logicalrep_relmap_update(m, v47+int32(664))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L11
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v47)+664))
	v2081 = F_logicalrep_rel_open(m, v2079, int32(0))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L11
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_initStringInfo(m, v47+int32(648))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L11
	} else {
		goto L377
	}
L377:
	;
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+689)))
	if v2091 != int32(114) {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v47)+668))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v47)+672))
	v2118 = F_quote_qualified_identifier(m, v2116, v2117)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L11
	} else {
		goto L384
	}
L379:
	;
	v2096 = int32(0)
	if (base.B2i32(v2091 != int32(112))|base.B2i32(v2034 != v2096)|base.B2i32(v687 < int32(_a_F_TableSyncWorkerMain_16))|v1702)&int32(1) == v2096 {
		goto L378
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	if (base.B2i32(v2034 != int32(0))|v1702)&int32(1) != 0 {
		goto L15
	} else {
		goto L383
	}
L382:
	;
	goto L15
L383:
	;
	goto L378
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+144)) = v2118
	v2126 = v47 + int32(648)
	F_appendStringInfo(m, v2126, int32(_a_F_TableSyncWorkerMain_34), v47+int32(144))
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L11
	} else {
		goto L385
	}
L385:
	;
	v2132 = int32(_a_F_TableSyncWorkerMain_35)
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2133 == int32(0) {
		v2555 = v2132
		goto L14
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2126, int32(_a_F_TableSyncWorkerMain_36))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L11
	} else {
		goto L387
	}
L387:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2143 <= int32(0) {
		goto L16
	} else {
		goto L388
	}
L388:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v47)+680))
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2146)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2152 = F_quote_identifier(m, v2147)
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2126, v2152)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L11
	} else {
		goto L390
	}
L390:
	;
	v2160 = int32(1)
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2161 <= v2160 {
		goto L16
	} else {
		goto L391
	}
L391:
	;
	v2171 = v2160
	goto L392
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2198 = v47 + int32(648)
	F_appendStringInfoString(m, v2198, int32(_a_F_TableSyncWorkerMain_37))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L11
	} else {
		goto L394
	}
L393:
	;
	goto L16
L394:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v47)+680))
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2202+v2171<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2211 = F_quote_identifier(m, v2206)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L11
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2198, v2211)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L11
	} else {
		goto L396
	}
L396:
	;
	v2220 = v2171 + int32(1)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2220 < v2221 {
		v2171 = v2220
		goto L392
	} else {
		goto L397
	}
L397:
	;
	goto L393
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L11
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L11
	} else {
		goto L402
	}
L401:
	;
	v3244 = v52
	v3246 = v54
	v3247 = v55
	v3248 = v56
	v3258 = int32(0)
	goto L12
L402:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v2246)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_pgstat_report_subscription_error(m, v2247)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L11
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v56)
	F_pg_re_throw(m)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L11
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
	v2555 = v2132
	goto L14
L406:
	;
	v2307 = int32(0)
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2307 < v2308 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v2318 = v2307
	goto L410
L408:
	;
	goto L409
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2409 = v47 + int32(648)
	F_appendStringInfoString(m, v2409, int32(_a_F_TableSyncWorkerMain_38))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L11
	} else {
		goto L419
	}
L410:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v47)+680))
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2340+v2318<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2349 = F_quote_identifier(m, v2344)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L11
	} else {
		goto L412
	}
L411:
	;
	goto L409
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2356 = v47 + int32(648)
	F_appendStringInfoString(m, v2356, v2349)
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L11
	} else {
		goto L413
	}
L413:
	;
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	if v2318 < v2359-int32(1) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2356, int32(_a_F_TableSyncWorkerMain_37))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L11
	} else {
		goto L417
	}
L415:
	;
	v2371 = v2359
	goto L416
L416:
	;
	v2373 = v2318 + int32(1)
	if v2373 < v2371 {
		v2318 = v2373
		goto L410
	} else {
		goto L418
	}
L417:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v47)+676))
	v2371 = v2370
	goto L416
L418:
	;
	goto L411
L419:
	;
	v2413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+689)))
	if v2413 == int32(114) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2409, int32(_a_F_TableSyncWorkerMain_39))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L11
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v47)+668))
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v47)+672))
	v2429 = F_quote_qualified_identifier(m, v2427, v2428)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L11
	} else {
		goto L424
	}
L423:
	;
	goto L422
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2436 = v47 + int32(648)
	F_appendStringInfoString(m, v2436, v2429)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L11
	} else {
		goto L425
	}
L425:
	;
	v2439 = int32(_a_F_TableSyncWorkerMain_40)
	if v2034 == int32(0) {
		v2555 = v2439
		goto L14
	} else {
		goto L426
	}
L426:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+12))
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2442)))
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2443)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v2444
	F_appendStringInfo(m, v2436, int32(_a_F_TableSyncWorkerMain_41), v47+int32(128))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L11
	} else {
		goto L427
	}
L427:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+4))
	if int32(2) <= v2456 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v2466 = int32(1)
	goto L431
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_list_free_deep(m, v2034)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L11
	} else {
		goto L435
	}
L431:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+12))
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2488+v2466<<(uint(int32(2))%32))))
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2492)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+112)) = v2493
	F_appendStringInfo(m, v47+int32(648), int32(_a_F_TableSyncWorkerMain_42), v47+int32(112))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L11
	} else {
		goto L433
	}
L432:
	;
	goto L430
L433:
	;
	v2507 = v2466 + int32(1)
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+4))
	if v2507 < v2508 {
		v2466 = v2507
		goto L431
	} else {
		goto L434
	}
L434:
	;
	goto L432
L435:
	;
	v2555 = v2439
	goto L14
L436:
	;
	if v687 < int32(_a_F_TableSyncWorkerMain_43) {
		v2626 = v54
		v2627 = int32(0)
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2629)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2636 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v47)+648))
	v2638 = int32(0)
	v2640 = m.T0[v2630].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2636, v2637, v2638, v2638)
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L11
	} else {
		goto L444
	}
L438:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v2588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2587)+34)))
	if v2588 != int32(1) {
		v2626 = v54
		v2627 = int32(0)
		goto L437
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_appendStringInfoString(m, v2579, int32(_a_F_TableSyncWorkerMain_44))
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L11
	} else {
		goto L440
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2603 = F_makeString(m, int32(_a_F_TableSyncWorkerMain_45))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L11
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2611 = F_makeDefElem(m, int32(_a_F_TableSyncWorkerMain_46), v2603, int32(-1))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L11
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+644)) = v2611
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+108)) = v2611
	v2623 = F_list_make1_impl(m, int32(1), v47+int32(108))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L11
	} else {
		goto L443
	}
L443:
	;
	v2626 = v2623
	v2627 = v2623
	goto L437
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v47)+648))
	F_pfree(m, v2646)
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L11
	} else {
		goto L445
	}
L445:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2640)))
	if v2649 != int32(4) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2659 = m.ExcPending
	if v2659 != 0 {
		goto L11
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+8))
	if v2689 != 0 {
		goto L453
	} else {
		goto L454
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L11
	} else {
		goto L450
	}
L450:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2672 = *(*int64)(unsafe.Add(mBase, uint32(v47)+668))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+104)) = v2667
	*(*int64)(unsafe.Add(mBase, uint32(v47)+96)) = v2672
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_47), v47+int32(96))
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L11
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1204), int32(_a_F_TableSyncWorkerMain_48))
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L11
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v2689)
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L11
	} else {
		goto L456
	}
L454:
	;
	goto L455
L455:
	;
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+12))
	if v2696 != 0 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	goto L455
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v2696)
	mBase = m.M
	v2702 = m.ExcPending
	if v2702 != 0 {
		goto L11
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+16))
	if v2703 != 0 {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	goto L459
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v2703)
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		goto L11
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v2640)
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L11
	} else {
		goto L465
	}
L464:
	;
	goto L463
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2721 = F_makeStringInfo(m)
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L11
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[16])) = v2721
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2729 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L11
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2736 = int32(0)
	v2739 = F_addRangeTableEntryForRelation(m, v2729, v382, int32(1), v2736, v2736, v2736)
	mBase = m.M
	v2740 = m.ExcPending
	if v2740 != 0 {
		goto L11
	} else {
		goto L468
	}
L468:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+12))
	if v2741 <= int32(0) {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2830 = int32(0)
	v2834 = F_BeginCopyFrom(m, v2729, v382, v2830, v2830, v2830, int32(1095), v2800, v2627)
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L11
	} else {
		goto L478
	}
L470:
	;
	v2799 = v52
	v2800 = int32(0)
	goto L469
L471:
	;
	goto L472
L472:
	;
	v2745 = int32(0)
	v2749 = v52
	v2754 = v2745
	v2756 = v2745
	goto L473
L473:
	;
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+16))
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2776+v2754<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2749
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2785 = F_makeString(m, v2780)
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L11
	} else {
		goto L475
	}
L474:
	;
	v2799 = v2791
	v2800 = v2791
	goto L469
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2749
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2791 = F_lappend(m, v2756, v2785)
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L11
	} else {
		goto L476
	}
L476:
	;
	v2794 = v2754 + int32(1)
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+12))
	if v2794 < v2795 {
		v2749 = v2791
		v2754 = v2794
		v2756 = v2791
		goto L473
	} else {
		goto L477
	}
L477:
	;
	goto L474
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2840 = F_CopyFrom(m, v2834)
	mBase = m.M
	v2841 = m.ExcPending
	if v2841 != 0 {
		goto L11
	} else {
		goto L479
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_EndCopyFrom(m, v2834)
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L11
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_logicalrep_rel_close(m, v2081, int32(0))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L11
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L11
	} else {
		goto L482
	}
L482:
	;
	v2862 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2862)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	v2869 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v2871 = int32(0)
	v2873 = m.T0[v2863].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2869, int32(_a_F_TableSyncWorkerMain_49), v2871, v2871)
	mBase = m.M
	v2874 = m.ExcPending
	if v2874 != 0 {
		goto L11
	} else {
		goto L483
	}
L483:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v2873)))
	if v2875 != int32(1) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2885 = m.ExcPending
	if v2885 != 0 {
		goto L11
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+8))
	if v2913 != 0 {
		goto L491
	} else {
		goto L492
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L11
	} else {
		goto L488
	}
L488:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = v2893
	F_errmsg(m, int32(_a_F_TableSyncWorkerMain_50), v47+int32(80))
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L11
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1493), int32(_a_F_TableSyncWorkerMain_5))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L11
	} else {
		goto L490
	}
L490:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v2913)
	mBase = m.M
	v2919 = m.ExcPending
	if v2919 != 0 {
		goto L11
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+12))
	if v2920 != 0 {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	goto L493
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_tuplestore_end(m, v2920)
	mBase = m.M
	v2926 = m.ExcPending
	if v2926 != 0 {
		goto L11
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+16))
	if v2927 != 0 {
		goto L499
	} else {
		goto L500
	}
L498:
	;
	goto L497
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_FreeTupleDesc(m, v2927)
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L11
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_pfree(m, v2873)
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L11
	} else {
		goto L503
	}
L502:
	;
	goto L501
L503:
	;
	if v548 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_RestoreUserContext(m, v47+int32(548))
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L11
	} else {
		goto L507
	}
L505:
	;
	goto L506
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_relation_close(m, v382, int32(0))
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L11
	} else {
		goto L508
	}
L507:
	;
	goto L506
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L11
	} else {
		goto L509
	}
L509:
	;
	v2964 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v2965 = *(*int64)(unsafe.Add(mBase, uint32(v2964)+48))
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v2964)+36))
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2964)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2799
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2626
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v1320
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v353)
	F_UpdateSubscriptionRelState(m, v2967, v2966, int32(102), v2965, int32(0))
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		goto L11
	} else {
		goto L510
	}
L510:
	;
	v2978 = v2799
	v2980 = v2626
	v2981 = v1320
	v2982 = v353
	goto L13
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3017 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v3018 = m.ExcPending
	if v3018 != 0 {
		goto L11
	} else {
		goto L512
	}
L512:
	;
	if v3017 != 0 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v3019 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	*(*uint32)(unsafe.Add(mBase, uint32(v47)+72)) = uint32(v3019)
	v3026 = int64(base.Ui64(v3019) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v47)+68)) = uint32(v3026)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v47 + int32(560)
	F_errmsg_internal(m, int32(_a_F_TableSyncWorkerMain_51), v47-int32(-64))
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L11
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v3050 = base.AtomicRmwXchg32(m, v3047, int32(56), int32(1))
	if v3050 != 0 {
		goto L518
	} else {
		goto L519
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	F_errfinish(m, int32(_a_F_TableSyncWorkerMain_4), int32(1520), int32(_a_F_TableSyncWorkerMain_5))
	mBase = m.M
	v3044 = m.ExcPending
	if v3044 != 0 {
		goto L11
	} else {
		goto L517
	}
L517:
	;
	goto L515
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	F_s_lock(m, v3047+int32(56), int32(_a_F_TableSyncWorkerMain_1))
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L11
	} else {
		goto L521
	}
L519:
	;
	goto L520
L520:
	;
	v3061 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v3062 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v3061)+40)) = uint8(v3062)
	v3064 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(v3061)+48)) = v3064
	v3066 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3061)+56)), uint32(v3066))
	goto L522
L521:
	;
	goto L520
L522:
	;
	v3099 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[17]))
	if v3099 != 0 {
		goto L524
	} else {
		goto L525
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3237 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[18]))
	F_LWLockRelease(m, v3237+int32(_a_F_TableSyncWorkerMain_52))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		goto L11
	} else {
		goto L554
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	F_ProcessInterrupts(m)
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L11
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v3107 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v3108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3107)+40)))
	if v3108 == int32(99) {
		v3244 = v2978
		v3246 = v2980
		v3247 = v2981
		v3248 = v2982
		v3258 = v179
		goto L12
	} else {
		goto L528
	}
L527:
	;
	goto L526
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3116 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[18]))
	v3120 = F_LWLockAcquire(m, v3116+int32(_a_F_TableSyncWorkerMain_52), int32(1))
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L11
	} else {
		goto L529
	}
L529:
	;
	v3123 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3123)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3130 = int32(0)
	v3138 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[19]))
	if v3138 <= v3130 {
		v3181 = v3130
		goto L531
	} else {
		goto L532
	}
L530:
	;
	if v3181 != 0 {
		goto L543
	} else {
		goto L544
	}
L531:
	;
	goto L530
L532:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[20]))
	v3150 = v3130
	goto L533
L533:
	;
	v3156 = v3142 + int32(16) + v3150<<(uint(int32(7))%32)
	v3157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3156)+16)))
	if v3157 != int32(1) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	v3181 = int32(0)
	goto L531
L535:
	;
	v3174 = v3150 + int32(1)
	if v3174 != v3138 {
		v3150 = v3174
		goto L533
	} else {
		goto L542
	}
L536:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v3156)))
	if v3160 == int32(4) {
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(v3156)+32))
	if v3163 != v3124 {
		goto L535
	} else {
		goto L538
	}
L538:
	;
	v3165 = *(*int32)(unsafe.Add(mBase, uint32(v3156)+36))
	if base.B2i32(v3165 != v3130)|base.B2i32(int32(3) != v3160) != 0 {
		goto L535
	} else {
		goto L539
	}
L539:
	;
	v3181 = v3156
	goto L531
L542:
	;
	goto L534
L543:
	;
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3181)+20))
	if v3186 != 0 {
		goto L546
	} else {
		goto L547
	}
L544:
	;
	goto L545
L545:
	;
	goto L523
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v3181)+20))
	F_SetLatch(m, v3191+int32(316))
	mBase = m.M
	goto L549
L547:
	;
	goto L548
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3200 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[18]))
	F_LWLockRelease(m, v3200+int32(_a_F_TableSyncWorkerMain_52))
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		goto L11
	} else {
		goto L550
	}
L549:
	;
	goto L548
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3210 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[21]))
	v3214 = F_WaitLatch(m, v3210, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L11
	} else {
		goto L551
	}
L551:
	;
	if v3214&int32(1) == int32(0) {
		goto L522
	} else {
		goto L552
	}
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v2978
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v2980
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v2981
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v2982)
	v3225 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[21]))
	v3226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3225))) = v3226
	v3231 = base.AtomicRmwOr32(m, v3226, int32(_a_F_TableSyncWorkerMain_53), v3226)
	goto L553
L553:
	;
	goto L522
L554:
	;
	v3244 = v2978
	v3246 = v2980
	v3247 = v2981
	v3248 = v2982
	v3258 = v179
	goto L12
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35+int32(52)))) = v3281
	*(*int32)(unsafe.Add(mBase, uint32(v47)+772)) = v3246
	*(*int32)(unsafe.Add(mBase, uint32(v47)+768)) = v3244
	*(*int32)(unsafe.Add(mBase, uint32(v47)+776)) = v3247
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)) = uint16(v3248)
	F_pfree(m, v3258)
	mBase = m.M
	v3289 = m.ExcPending
	if v3289 != 0 {
		goto L11
	} else {
		goto L556
	}
L556:
	;
	goto L10
L557:
	;
	v3324 = int32(v3320)
	m.G0 = v47
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(v3324)+4))
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v3324)))
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3327)))
	if v47+int32(380) == v3330 {
		goto L560
	} else {
		goto L561
	}
L558:
	;
	m.ExcPending = 1
	goto L1
L559:
	;
	if v3334 != 0 {
		goto L563
	} else {
		goto L564
	}
L560:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(v3327)+4))
	v3334 = v3332
	goto L562
L561:
	;
	v3334 = int32(0)
	goto L562
L562:
	;
	goto L559
L563:
	;
	v3335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+782)))
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(v47)+776))
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v47)+772))
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v47)+768))
	v52 = v3338
	v54 = v3337
	v55 = v3336
	v56 = v3335
	v57 = v3326
	v58 = v3334
	v72 = v94
	v73 = v95
	goto L3
L564:
	;
	goto L565
L565:
	;
	F___wasm_longjmp(m, v3327, v3326)
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L567:
	;
	F_set_apply_error_context_origin(m, v3352)
	mBase = m.M
	v3356 = m.ExcPending
	if v3356 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v3358 = int32(1)
	v3360 = v35 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v3360))) = uint8(v3358)
	v3365 = *(*int64)(unsafe.Add(mBase, uint32(v35+int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+4)) = v3357
	*(*int64)(unsafe.Add(mBase, uint32(v3360)+8)) = v3365
	v3370 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v3372 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v3372)+24))
	v3374 = m.T0[v3373].(func(*base.Module, int32) int32)(m, v3370)
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+28)) = v3424
	v3427 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3427)+68)) = uint8(v3423)
	v3429 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3360)+32)) = uint8(v3429)
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3421)+68))
	v3432 = F_pstrdup(m, v3431)
	mBase = m.M
	v3433 = m.ExcPending
	if v3433 != 0 {
		goto L1
	} else {
		goto L588
	}
L570:
	;
	v3413 = int32(0)
	if v3411&int32(255) != int32(102) {
		goto L585
	} else {
		goto L586
	}
L571:
	;
	if v3374 <= int32(_a_F_TableSyncWorkerMain_54) {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	if int32(_a_F_TableSyncWorkerMain_55) < v3374 {
		goto L575
	} else {
		goto L576
	}
L573:
	;
	goto L574
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+16)) = int32(4)
	v3401 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v3401)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+20)) = v3402
	v3404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3401)+34)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3360)+24)) = uint8(v3404)
	v3407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3401)+35)))
	if v3407 == int32(112) {
		v3421 = v3401
		v3423 = v3358
		v3424 = int32(_a_F_TableSyncWorkerMain_56)
		goto L569
	} else {
		goto L584
	}
L575:
	;
	v3383 = int32(2)
	goto L577
L576:
	;
	v3383 = int32(1)
	goto L577
L577:
	;
	if int32(_a_F_TableSyncWorkerMain_57) < v3374 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v3386 = int32(3)
	goto L580
L579:
	;
	v3386 = v3383
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+16)) = v3386
	v3389 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[0]))
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v3389)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+20)) = v3390
	v3392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3389)+34)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3360)+24)) = uint8(v3392)
	if v3374 < int32(_a_F_TableSyncWorkerMain_58) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v3421 = v3389
	v3423 = int32(0)
	v3424 = int32(0)
	goto L569
L582:
	;
	goto L583
L583:
	;
	v3397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3389)+35)))
	v3410 = v3389
	v3411 = v3397
	goto L570
L584:
	;
	v3410 = v3401
	v3411 = v3407
	goto L570
L585:
	;
	v3420 = int32(_a_F_TableSyncWorkerMain_59)
	goto L587
L586:
	;
	v3420 = v3413
	goto L587
L587:
	;
	v3421 = v3410
	v3423 = v3413
	v3424 = v3420
	goto L569
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+36)) = v3432
	v3436 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[8]))
	v3440 = *(*int32)(unsafe.Add(mBase, _c_F_TableSyncWorkerMain[6]))
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3440)+32))
	v3442 = m.T0[v3441].(func(*base.Module, int32, int32) int32)(m, v3436, v35+int32(8))
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v3444 = *(*int64)(unsafe.Add(mBase, uint32(v35)+56))
	F_start_apply(m, v3444)
	mBase = m.M
	v3446 = m.ExcPending
	if v3446 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	m.G0 = v35 + int32(128)
	F_FinishSyncWorker(m)
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
