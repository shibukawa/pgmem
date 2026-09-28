package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformFromClauseItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
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
	var v105 int32
	_ = v105
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v205 int32
	_ = v205
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 float64
	_ = v704
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v860 int32
	_ = v860
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v914 int32
	_ = v914
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v1029 int32
	_ = v1029
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1103 int32
	_ = v1103
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
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
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1344 int32
	_ = v1344
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1403 int32
	_ = v1403
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1532 int32
	_ = v1532
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1584 int32
	_ = v1584
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1682 int32
	_ = v1682
	var v1691 int32
	_ = v1691
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1815 int32
	_ = v1815
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1910 int32
	_ = v1910
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v2003 int32
	_ = v2003
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2041 int32
	_ = v2041
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2150 int32
	_ = v2150
	var v2172 int32
	_ = v2172
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2232 int64
	_ = v2232
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int64
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2334 int64
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2368 int32
	_ = v2368
	var v2381 int32
	_ = v2381
	var v2389 int32
	_ = v2389
	var v2399 int32
	_ = v2399
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2448 int32
	_ = v2448
	var v2474 int32
	_ = v2474
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2490 int32
	_ = v2490
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2504 int32
	_ = v2504
	var v2508 int32
	_ = v2508
	var v2514 int32
	_ = v2514
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2525 int32
	_ = v2525
	var v2560 int32
	_ = v2560
	var v2584 int32
	_ = v2584
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2624 int32
	_ = v2624
	var v2628 int32
	_ = v2628
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2669 int32
	_ = v2669
	var v2672 int32
	_ = v2672
	var v2678 int32
	_ = v2678
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2692 int32
	_ = v2692
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2713 int32
	_ = v2713
	var v2720 int32
	_ = v2720
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
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
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2802 int32
	_ = v2802
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2853 int32
	_ = v2853
	var v2857 int32
	_ = v2857
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2873 int32
	_ = v2873
	var v2876 int32
	_ = v2876
	var v2878 int32
	_ = v2878
	var v2881 int32
	_ = v2881
	var v2884 int32
	_ = v2884
	var v2887 int32
	_ = v2887
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2898 int32
	_ = v2898
	var v2901 int32
	_ = v2901
	var v2907 int32
	_ = v2907
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2921 int32
	_ = v2921
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2935 int32
	_ = v2935
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v3004 int32
	_ = v3004
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3071 int32
	_ = v3071
	var v3073 int32
	_ = v3073
	var v3077 int32
	_ = v3077
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3157 int32
	_ = v3157
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3178 int32
	_ = v3178
	var v3182 int32
	_ = v3182
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3192 int32
	_ = v3192
	var v3198 int32
	_ = v3198
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3226 int32
	_ = v3226
	var v3232 int32
	_ = v3232
	var v3238 int32
	_ = v3238
	var v3240 int32
	_ = v3240
	var v3246 int32
	_ = v3246
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3261 int32
	_ = v3261
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3316 int32
	_ = v3316
	var v3346 int32
	_ = v3346
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3367 int32
	_ = v3367
	var v3369 int32
	_ = v3369
	var v3377 int32
	_ = v3377
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3416 int32
	_ = v3416
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3434 int32
	_ = v3434
	var v3439 int32
	_ = v3439
	var v3443 int32
	_ = v3443
	var v3446 int32
	_ = v3446
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3453 int32
	_ = v3453
	var v3458 int32
	_ = v3458
	var v3462 int32
	_ = v3462
	var v3465 int32
	_ = v3465
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3480 int32
	_ = v3480
	var v3484 int32
	_ = v3484
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3500 int32
	_ = v3500
	var v3505 int32
	_ = v3505
	var v3509 int32
	_ = v3509
	var v3512 int32
	_ = v3512
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3525 int32
	_ = v3525
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3546 int32
	_ = v3546
	var v3553 int32
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3592 int32
	_ = v3592
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3606 int32
	_ = v3606
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3619 int32
	_ = v3619
	var v3620 int32
	_ = v3620
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3636 int32
	_ = v3636
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3650 int32
	_ = v3650
	var v3651 int32
	_ = v3651
	var v3652 int32
	_ = v3652
	var v3653 int32
	_ = v3653
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3658 int32
	_ = v3658
	var v3662 int32
	_ = v3662
	var v3664 int32
	_ = v3664
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3673 int32
	_ = v3673
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3683 int32
	_ = v3683
	var v3686 int32
	_ = v3686
	var v3693 int32
	_ = v3693
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3720 int32
	_ = v3720
	var v3721 int32
	_ = v3721
	var v3724 int32
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3739 int32
	_ = v3739
	var v3740 int32
	_ = v3740
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3748 int32
	_ = v3748
	var v3761 int32
	_ = v3761
	var v3785 int32
	_ = v3785
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3794 int32
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3805 int32
	_ = v3805
	var v3809 int32
	_ = v3809
	var v3810 int32
	_ = v3810
	var v3818 int32
	_ = v3818
	var v3819 int32
	_ = v3819
	var v3821 int32
	_ = v3821
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3829 int32
	_ = v3829
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
	var v3838 int32
	_ = v3838
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3846 int32
	_ = v3846
	var v3847 int32
	_ = v3847
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3858 int32
	_ = v3858
	var v3859 int32
	_ = v3859
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3871 int32
	_ = v3871
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3886 int32
	_ = v3886
	var v3919 int32
	_ = v3919
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3936 int32
	_ = v3936
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3952 int32
	_ = v3952
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3965 int32
	_ = v3965
	var v3970 int32
	_ = v3970
	var v3974 int32
	_ = v3974
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3984 int32
	_ = v3984
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3996 int32
	_ = v3996
	var v3997 int32
	_ = v3997
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4006 int32
	_ = v4006
	var v4011 int32
	_ = v4011
	var v4049 int32
	_ = v4049
	var v4050 int32
	_ = v4050
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4060 int32
	_ = v4060
	var v4065 int32
	_ = v4065
	var v4069 int32
	_ = v4069
	var v4072 int32
	_ = v4072
	var v4076 int32
	_ = v4076
	var v4080 int32
	_ = v4080
	var v4081 int32
	_ = v4081
	var v4082 int32
	_ = v4082
	var v4084 int32
	_ = v4084
	var v4089 int32
	_ = v4089
	var v4093 int32
	_ = v4093
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4101 int32
	_ = v4101
	var v4106 int32
	_ = v4106
	var v4110 int32
	_ = v4110
	var v4113 int32
	_ = v4113
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4121 int32
	_ = v4121
	var v4126 int32
	_ = v4126
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4138 int32
	_ = v4138
	var v4139 int32
	_ = v4139
	var v4141 int32
	_ = v4141
	var v4146 int32
	_ = v4146
	var v4150 int32
	_ = v4150
	var v4153 int32
	_ = v4153
	var v4157 int32
	_ = v4157
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4161 int32
	_ = v4161
	var v4166 int32
	_ = v4166
	var v4170 int32
	_ = v4170
	var v4174 int32
	_ = v4174
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4194 int32
	_ = v4194
	var v4202 int32
	_ = v4202
	var v4232 int32
	_ = v4232
	var v4236 int32
	_ = v4236
	var v4238 int32
	_ = v4238
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4285 int32
	_ = v4285
	var v4290 int32
	_ = v4290
	var v4291 int32
	_ = v4291
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4301 int32
	_ = v4301
	var v4304 int32
	_ = v4304
	var v4306 int32
	_ = v4306
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4316 int32
	_ = v4316
	var v4317 int32
	_ = v4317
	var v4318 int32
	_ = v4318
	var v4319 int32
	_ = v4319
	var v4324 int32
	_ = v4324
	var v4333 int32
	_ = v4333
	var v4338 int32
	_ = v4338
	var v4361 int32
	_ = v4361
	var v4365 int32
	_ = v4365
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4373 int32
	_ = v4373
	var v4376 int32
	_ = v4376
	var v4379 int32
	_ = v4379
	var v4380 int32
	_ = v4380
	var v4386 int32
	_ = v4386
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4423 int32
	_ = v4423
	var v4426 int32
	_ = v4426
	var v4429 int32
	_ = v4429
	var v4430 int32
	_ = v4430
	var v4433 int32
	_ = v4433
	var v4434 int32
	_ = v4434
	var v4437 int32
	_ = v4437
	var v4444 int32
	_ = v4444
	var v4445 int32
	_ = v4445
	var v4448 int32
	_ = v4448
	var v4450 int32
	_ = v4450
	var v4451 int32
	_ = v4451
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4467 int32
	_ = v4467
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4505 int32
	_ = v4505
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4566 int32
	_ = v4566
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4578 int32
	_ = v4578
	var v4580 int32
	_ = v4580
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4587 int32
	_ = v4587
	var v4599 int32
	_ = v4599
	var v4608 int32
	_ = v4608
	var v4613 int32
	_ = v4613
	var v4615 int32
	_ = v4615
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4624 int32
	_ = v4624
	var v4627 int32
	_ = v4627
	var v4628 int32
	_ = v4628
	var v4629 int32
	_ = v4629
	var v4632 int32
	_ = v4632
	var v4635 int32
	_ = v4635
	var v4638 int32
	_ = v4638
	var v4639 int32
	_ = v4639
	var v4645 int32
	_ = v4645
	var v4678 int32
	_ = v4678
	var v4679 int32
	_ = v4679
	var v4682 int32
	_ = v4682
	var v4685 int32
	_ = v4685
	var v4688 int32
	_ = v4688
	var v4689 int32
	_ = v4689
	var v4692 int32
	_ = v4692
	var v4693 int32
	_ = v4693
	var v4696 int32
	_ = v4696
	var v4703 int32
	_ = v4703
	var v4704 int32
	_ = v4704
	var v4707 int32
	_ = v4707
	var v4712 int32
	_ = v4712
	var v4715 int32
	_ = v4715
	var v4721 int32
	_ = v4721
	var v4726 int32
	_ = v4726
	var v4763 int32
	_ = v4763
	var v4767 int32
	_ = v4767
	var v4770 int32
	_ = v4770
	var v4771 int32
	_ = v4771
	var v4778 int32
	_ = v4778
	var v4780 int32
	_ = v4780
	var v4811 int32
	_ = v4811
	var v4812 int32
	_ = v4812
	var v4815 int32
	_ = v4815
	var v4818 int32
	_ = v4818
	var v4821 int32
	_ = v4821
	var v4822 int32
	_ = v4822
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4829 int32
	_ = v4829
	var v4836 int32
	_ = v4836
	var v4837 int32
	_ = v4837
	var v4841 int32
	_ = v4841
	var v4845 int32
	_ = v4845
	var v4847 int32
	_ = v4847
	var v4855 int32
	_ = v4855
	var v4887 int32
	_ = v4887
	var v4888 int32
	_ = v4888
	var v4891 int32
	_ = v4891
	var v4895 int32
	_ = v4895
	var v4898 int32
	_ = v4898
	var v4899 int32
	_ = v4899
	var v4906 int32
	_ = v4906
	var v4913 int32
	_ = v4913
	var v4939 int32
	_ = v4939
	var v4940 int32
	_ = v4940
	var v4943 int32
	_ = v4943
	var v4946 int32
	_ = v4946
	var v4949 int32
	_ = v4949
	var v4950 int32
	_ = v4950
	var v4953 int32
	_ = v4953
	var v4954 int32
	_ = v4954
	var v4957 int32
	_ = v4957
	var v4964 int32
	_ = v4964
	var v4965 int32
	_ = v4965
	var v4969 int32
	_ = v4969
	var v4973 int32
	_ = v4973
	var v4975 int32
	_ = v4975
	var v4988 int32
	_ = v4988
	var v5015 int32
	_ = v5015
	var v5016 int32
	_ = v5016
	var v5019 int32
	_ = v5019
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5028 int32
	_ = v5028
	var v5030 int32
	_ = v5030
	var v5032 int32
	_ = v5032
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5037 int32
	_ = v5037
	var v5040 int32
	_ = v5040
	var v5041 int32
	_ = v5041
	var v5042 int32
	_ = v5042
	var v5043 int32
	_ = v5043
	var v5044 int32
	_ = v5044
	var v5045 int32
	_ = v5045
	var v5047 int32
	_ = v5047
	var v5048 int32
	_ = v5048
	var v5049 int32
	_ = v5049
	var v5051 int32
	_ = v5051
	var v5053 int32
	_ = v5053
	var v5056 int32
	_ = v5056
	var v5057 int32
	_ = v5057
	var v5058 int32
	_ = v5058
	var v5059 int32
	_ = v5059
	var v5060 int32
	_ = v5060
	var v5061 int32
	_ = v5061
	var v5063 int32
	_ = v5063
	var v5064 int32
	_ = v5064
	var v5066 int32
	_ = v5066
	var v5069 int32
	_ = v5069
	var v5076 int32
	_ = v5076
	var v5106 int32
	_ = v5106
	var v5110 int32
	_ = v5110
	var v5111 int32
	_ = v5111
	var v5114 int32
	_ = v5114
	var v5115 int32
	_ = v5115
	var v5151 int32
	_ = v5151
	var v5154 int32
	_ = v5154
	var v5156 int32
	_ = v5156
	var v5157 int32
	_ = v5157
	var v5159 int32
	_ = v5159
	var v5160 int32
	_ = v5160
	var v5161 int32
	_ = v5161
	var v5201 int32
	_ = v5201
	var v5204 int32
	_ = v5204
	var v5210 int32
	_ = v5210
	var v5215 int32
	_ = v5215
	var v5253 int32
	_ = v5253
	var v5256 int32
	_ = v5256
	var v5262 int32
	_ = v5262
	var v5267 int32
	_ = v5267
	var v5271 int32
	_ = v5271
	var v5274 int32
	_ = v5274
	var v5280 int32
	_ = v5280
	var v5285 int32
	_ = v5285
	var v5323 int32
	_ = v5323
	var v5326 int32
	_ = v5326
	var v5332 int32
	_ = v5332
	var v5337 int32
	_ = v5337
	var v5347 int32
	_ = v5347
	var v5356 int32
	_ = v5356
	var v5361 int32
	_ = v5361
	var v5363 int32
	_ = v5363
	var v5366 int32
	_ = v5366
	var v5375 int32
	_ = v5375
	var v5381 int32
	_ = v5381
	var v5382 int32
	_ = v5382
	var v5411 int32
	_ = v5411
	var v5416 int32
	_ = v5416
	var v5418 int32
	_ = v5418
	var v5422 int32
	_ = v5422
	var v5425 int32
	_ = v5425
	var v5427 int32
	_ = v5427
	var v5432 int32
	_ = v5432
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5443 int32
	_ = v5443
	var v5444 int32
	_ = v5444
	var v5448 int32
	_ = v5448
	var v5449 int32
	_ = v5449
	var v5451 int32
	_ = v5451
	var v5453 int32
	_ = v5453
	var v5458 int32
	_ = v5458
	var v5459 int32
	_ = v5459
	var v5460 int32
	_ = v5460
	var v5461 int32
	_ = v5461
	var v5463 int32
	_ = v5463
	var v5464 int32
	_ = v5464
	var v5465 int32
	_ = v5465
	var v5466 int32
	_ = v5466
	var v5467 int32
	_ = v5467
	var v5468 int32
	_ = v5468
	var v5471 int32
	_ = v5471
	var v5473 int32
	_ = v5473
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5477 int32
	_ = v5477
	var v5488 int32
	_ = v5488
	var v5507 int32
	_ = v5507
	var v5513 int32
	_ = v5513
	var v5514 int32
	_ = v5514
	var v5518 int32
	_ = v5518
	var v5520 int32
	_ = v5520
	var v5521 int32
	_ = v5521
	var v5523 int32
	_ = v5523
	var v5524 int32
	_ = v5524
	var v5525 int32
	_ = v5525
	var v5527 int32
	_ = v5527
	var v5528 int32
	_ = v5528
	var v5530 int32
	_ = v5530
	var v5534 int32
	_ = v5534
	var v5535 int32
	_ = v5535
	var v5541 int32
	_ = v5541
	var v5546 int32
	_ = v5546
	var v5547 int32
	_ = v5547
	var v5549 int32
	_ = v5549
	var v5550 int32
	_ = v5550
	var v5554 int32
	_ = v5554
	var v5561 int32
	_ = v5561
	var v5581 int32
	_ = v5581
	var v5590 int32
	_ = v5590
	var v5594 int32
	_ = v5594
	var v5596 int32
	_ = v5596
	var v5600 int32
	_ = v5600
	var v5603 int32
	_ = v5603
	var v5607 int32
	_ = v5607
	var v5610 int32
	_ = v5610
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5620 int32
	_ = v5620
	var v5622 int32
	_ = v5622
	var v5623 int32
	_ = v5623
	var v5626 int32
	_ = v5626
	var v5629 int32
	_ = v5629
	var v5632 int32
	_ = v5632
	var v5635 int32
	_ = v5635
	var v5637 int32
	_ = v5637
	var v5638 int32
	_ = v5638
	var v5641 int32
	_ = v5641
	var v5645 int32
	_ = v5645
	var v5649 int32
	_ = v5649
	var v5652 int32
	_ = v5652
	var v5655 int32
	_ = v5655
	var v5657 int32
	_ = v5657
	var v5658 int32
	_ = v5658
	var v5661 int32
	_ = v5661
	var v5664 int32
	_ = v5664
	var v5667 int32
	_ = v5667
	var v5670 int32
	_ = v5670
	var v5672 int32
	_ = v5672
	var v5673 int32
	_ = v5673
	var v5676 int32
	_ = v5676
	var v5680 int32
	_ = v5680
	var v5684 int32
	_ = v5684
	var v5687 int32
	_ = v5687
	var v5688 int32
	_ = v5688
	var v5697 int32
	_ = v5697
	var v5698 int32
	_ = v5698
	var v5701 int32
	_ = v5701
	var v5702 int32
	_ = v5702
	var v5711 int32
	_ = v5711
	var v5712 int32
	_ = v5712
	var v5713 int32
	_ = v5713
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5720 int32
	_ = v5720
	var v5721 int32
	_ = v5721
	var v5722 int32
	_ = v5722
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5734 int32
	_ = v5734
	var v5735 int32
	_ = v5735
	var v5736 int32
	_ = v5736
	var v5740 int32
	_ = v5740
	var v5741 int32
	_ = v5741
	var v5742 int32
	_ = v5742
	var v5743 int32
	_ = v5743
	var v5746 int32
	_ = v5746
	var v5749 int32
	_ = v5749
	var v5751 int32
	_ = v5751
	var v5752 int32
	_ = v5752
	var v5764 int32
	_ = v5764
	var v5765 int32
	_ = v5765
	var v5772 int32
	_ = v5772
	var v5778 int32
	_ = v5778
	var v5783 int32
	_ = v5783
	var v5786 int32
	_ = v5786
	var v5788 int32
	_ = v5788
	var v5790 int32
	_ = v5790
	var v5793 int32
	_ = v5793
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5797 int64
	_ = v5797
	var v5799 int64
	_ = v5799
	var v5801 int64
	_ = v5801
	var v5803 int64
	_ = v5803
	var v5806 int64
	_ = v5806
	var v5808 int64
	_ = v5808
	var v5810 int64
	_ = v5810
	var v5812 int64
	_ = v5812
	var v5814 int32
	_ = v5814
	var v5817 int32
	_ = v5817
	var v5818 int32
	_ = v5818
	var v5820 int32
	_ = v5820
	var v5821 int32
	_ = v5821
	var v5823 int32
	_ = v5823
	var v5824 int32
	_ = v5824
	var v5826 int32
	_ = v5826
	var v5829 int32
	_ = v5829
	var v5831 int32
	_ = v5831
	var v5838 int32
	_ = v5838
	var v5870 int32
	_ = v5870
	var v5872 int32
	_ = v5872
	var v5878 int32
	_ = v5878
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5884 int32
	_ = v5884
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5889 int32
	_ = v5889
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5894 int32
	_ = v5894
	var v5909 int32
	_ = v5909
	var v5912 int32
	_ = v5912
	var v5939 int32
	_ = v5939
	var v5940 int32
	_ = v5940
	var v5942 int32
	_ = v5942
	var v5944 int32
	_ = v5944
	var v5947 int32
	_ = v5947
	var v5950 int32
	_ = v5950
	var v5952 int32
	_ = v5952
	var v5955 int32
	_ = v5955
	var v5958 int32
	_ = v5958
	var v5960 int32
	_ = v5960
	var v5963 int32
	_ = v5963
	var v5966 int32
	_ = v5966
	var v5967 int32
	_ = v5967
	var v5968 int32
	_ = v5968
	var v5972 int32
	_ = v5972
	var v5980 int32
	_ = v5980
	var v6014 int32
	_ = v6014
	var v6020 int32
	_ = v6020
	var v6044 int32
	_ = v6044
	var v6047 int32
	_ = v6047
	var v6048 int32
	_ = v6048
	var v6049 int32
	_ = v6049
	var v6053 int32
	_ = v6053
	var v6090 int32
	_ = v6090
	var v6091 int32
	_ = v6091
	var v6092 int32
	_ = v6092
	var v6093 int32
	_ = v6093
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6097 int32
	_ = v6097
	var v6098 int32
	_ = v6098
	var v6099 int32
	_ = v6099
	var v6100 int32
	_ = v6100
	var v6102 int32
	_ = v6102
	var v6103 int32
	_ = v6103
	var v6104 int32
	_ = v6104
	var v6105 int32
	_ = v6105
	var v6108 int32
	_ = v6108
	var v6109 int32
	_ = v6109
	var v6115 int32
	_ = v6115
	var v6121 int32
	_ = v6121
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6157 int32
	_ = v6157
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6190 int32
	_ = v6190
	var v6192 int32
	_ = v6192
	var v6193 int32
	_ = v6193
	var v6194 int32
	_ = v6194
	var v6196 int32
	_ = v6196
	var v6198 int32
	_ = v6198
	var v6210 int32
	_ = v6210
	var v6211 int32
	_ = v6211
	var v6213 int32
	_ = v6213
	var v6214 int32
	_ = v6214
	var v6215 int32
	_ = v6215
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6221 int32
	_ = v6221
	var v6222 int32
	_ = v6222
	var v6225 int32
	_ = v6225
	var v6226 int32
	_ = v6226
	var v6233 int32
	_ = v6233
	var v6263 int32
	_ = v6263
	var v6267 int32
	_ = v6267
	var v6268 int32
	_ = v6268
	var v6271 int32
	_ = v6271
	var v6272 int32
	_ = v6272
	var v6274 int32
	_ = v6274
	var v6312 int32
	_ = v6312
	var v6323 int32
	_ = v6323
	var v6347 int32
	_ = v6347
	var v6348 int32
	_ = v6348
	var v6354 int32
	_ = v6354
	var v6355 int32
	_ = v6355
	var v6357 int32
	_ = v6357
	var v6358 int32
	_ = v6358
	var v6359 int32
	_ = v6359
	var v6360 int32
	_ = v6360
	var v6361 int32
	_ = v6361
	var v6362 int32
	_ = v6362
	var v6363 int32
	_ = v6363
	var v6365 int32
	_ = v6365
	var v6377 int32
	_ = v6377
	var v6378 int32
	_ = v6378
	var v6382 int32
	_ = v6382
	var v6383 int32
	_ = v6383
	var v6384 int32
	_ = v6384
	var v6385 int32
	_ = v6385
	var v6388 int32
	_ = v6388
	var v6389 int32
	_ = v6389
	var v6391 int32
	_ = v6391
	var v6392 int32
	_ = v6392
	var v6396 int32
	_ = v6396
	var v6397 int32
	_ = v6397
	var v6398 int32
	_ = v6398
	var v6399 int32
	_ = v6399
	var v6400 int32
	_ = v6400
	var v6401 int32
	_ = v6401
	var v6403 int32
	_ = v6403
	var v6409 int32
	_ = v6409
	var v6411 int32
	_ = v6411
	var v6439 int32
	_ = v6439
	var v6443 int32
	_ = v6443
	var v6445 int32
	_ = v6445
	var v6449 int32
	_ = v6449
	var v6455 int32
	_ = v6455
	var v6458 int32
	_ = v6458
	var v6460 int32
	_ = v6460
	var v6462 int32
	_ = v6462
	var v6463 int32
	_ = v6463
	var v6467 int32
	_ = v6467
	var v6468 int32
	_ = v6468
	var v6471 int32
	_ = v6471
	var v6472 int32
	_ = v6472
	var v6474 int32
	_ = v6474
	var v6475 int32
	_ = v6475
	var v6481 int32
	_ = v6481
	var v6482 int32
	_ = v6482
	var v6484 int32
	_ = v6484
	var v6485 int32
	_ = v6485
	var v6487 int32
	_ = v6487
	var v6488 int32
	_ = v6488
	var v6490 int32
	_ = v6490
	var v6493 int32
	_ = v6493
	var v6494 int32
	_ = v6494
	var v6498 int32
	_ = v6498
	var v6501 int32
	_ = v6501
	var v6502 int32
	_ = v6502
	var v6503 int32
	_ = v6503
	var v6504 int32
	_ = v6504
	var v6510 int32
	_ = v6510
	var v6511 int32
	_ = v6511
	var v6513 int32
	_ = v6513
	var v6518 int32
	_ = v6518
	var v6522 int32
	_ = v6522
	var v6525 int32
	_ = v6525
	var v6526 int32
	_ = v6526
	var v6527 int32
	_ = v6527
	var v6528 int32
	_ = v6528
	var v6536 int32
	_ = v6536
	var v6537 int32
	_ = v6537
	var v6539 int32
	_ = v6539
	var v6544 int32
	_ = v6544
	var v6548 int32
	_ = v6548
	var v6551 int32
	_ = v6551
	var v6552 int32
	_ = v6552
	var v6553 int32
	_ = v6553
	var v6554 int32
	_ = v6554
	var v6560 int32
	_ = v6560
	var v6561 int32
	_ = v6561
	var v6563 int32
	_ = v6563
	var v6568 int32
	_ = v6568
	var v6573 int32
	_ = v6573
	var v6576 int32
	_ = v6576
	var v6580 int32
	_ = v6580
	var v6581 int32
	_ = v6581
	var v6582 int32
	_ = v6582
	var v6584 int32
	_ = v6584
	var v6589 int32
	_ = v6589
	var v6593 int32
	_ = v6593
	var v6596 int32
	_ = v6596
	var v6598 int32
	_ = v6598
	var v6599 int32
	_ = v6599
	var v6600 int32
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6602 int32
	_ = v6602
	var v6603 int32
	_ = v6603
	var v6604 int32
	_ = v6604
	var v6605 int32
	_ = v6605
	var v6606 int32
	_ = v6606
	var v6607 int32
	_ = v6607
	var v6608 int32
	_ = v6608
	var v6610 int32
	_ = v6610
	var v6619 int32
	_ = v6619
	var v6620 int32
	_ = v6620
	var v6622 int32
	_ = v6622
	var v6627 int32
	_ = v6627
	var v6663 int32
	_ = v6663
	var v6664 int32
	_ = v6664
	var v6665 int32
	_ = v6665
	var v6674 int32
	_ = v6674
	var v6676 int32
	_ = v6676
	var v6690 int32
	_ = v6690
	var v6696 int32
	_ = v6696
	var v6703 int32
	_ = v6703
	var v6707 int32
	_ = v6707
	var v6708 int32
	_ = v6708
	var v6710 int32
	_ = v6710
	var v6711 int32
	_ = v6711
	var v6714 int32
	_ = v6714
	var v6715 int32
	_ = v6715
	var v6717 int32
	_ = v6717
	var v6718 int32
	_ = v6718
	var v6719 int32
	_ = v6719
	var v6720 int32
	_ = v6720
	var v6723 int32
	_ = v6723
	var v6726 int32
	_ = v6726
	var v6732 int32
	_ = v6732
	var v6765 int32
	_ = v6765
	var v6766 int32
	_ = v6766
	var v6769 int32
	_ = v6769
	var v6772 int32
	_ = v6772
	var v6775 int32
	_ = v6775
	var v6776 int32
	_ = v6776
	var v6779 int32
	_ = v6779
	var v6780 int32
	_ = v6780
	var v6783 int32
	_ = v6783
	var v6790 int32
	_ = v6790
	var v6791 int32
	_ = v6791
	var v6796 int32
	_ = v6796
	var v6832 int32
	_ = v6832
	var v6833 int32
	_ = v6833
	var v6834 int32
	_ = v6834
	var v6841 int32
	_ = v6841
	var v6844 int32
	_ = v6844
	var v6848 int32
	_ = v6848
	var v6849 int32
	_ = v6849
	var v6851 int32
	_ = v6851
	var v6856 int32
	_ = v6856
	var v6860 int32
	_ = v6860
	var v6863 int32
	_ = v6863
	var v6864 int32
	_ = v6864
	var v6870 int32
	_ = v6870
	var v6871 int32
	_ = v6871
	var v6873 int32
	_ = v6873
	var v6878 int32
	_ = v6878
	var v6883 int32
	_ = v6883
	var v6884 int32
	_ = v6884
	var v6913 int32
	_ = v6913
	var v6914 int32
	_ = v6914
	var v6916 int32
	_ = v6916
	var v6917 int32
	_ = v6917
	var v6940 int32
	_ = v6940
	var v6946 int32
	_ = v6946
	var v6989 int32
	_ = v6989
	var v6991 int32
	_ = v6991
	var v6994 int32
	_ = v6994
	var v6998 int32
	_ = v6998
	var v6999 int32
	_ = v6999
	var v7000 int32
	_ = v7000
	var v7001 int32
	_ = v7001
	var v7002 int32
	_ = v7002
	var v7003 int32
	_ = v7003
	var v7038 int32
	_ = v7038
	var v7045 int32
	_ = v7045
	var v7046 int32
	_ = v7046
	var v7049 int32
	_ = v7049
	var v7050 int32
	_ = v7050
	var v7053 int32
	_ = v7053
	var v7056 int32
	_ = v7056
	v5 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(336)
	m.G0 = v37
	F_check_stack_depth(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v43 - int32(85) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L15
	case 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38:
		goto L14
	case 4:
		goto L6
	case 39:
		goto L16
	default:
		goto L19
	}
L3:
	;
	m.G0 = v37 + int32(336)
	return v7056
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v7038
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v7038
	*(*int32)(unsafe.Add(mBase, uint32(v37)+296)) = v7038
	v7045 = F_list_make1_impl(m, int32(1), v37+int32(220))
	mBase = m.M
	v7046 = m.ExcPending
	if v7046 != 0 {
		goto L1
	} else {
		goto L1249
	}
L5:
	;
	F_pfree(m, v3743)
	mBase = m.M
	v6663 = m.ExcPending
	if v6663 != 0 {
		goto L1
	} else {
		goto L1194
	}
L6:
	;
	v6357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6358 = F_transformFromClauseItem(m, l0, v6357, l2, l3)
	mBase = m.M
	v6359 = m.ExcPending
	if v6359 != 0 {
		goto L1
	} else {
		goto L1110
	}
L7:
	;
	v4180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4185 = F_transformFromClauseItem(m, l0, v4180, v37+int32(292), v37+int32(284))
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L1
	} else {
		goto L778
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4170 = m.ExcPending
	if v4170 != 0 {
		goto L1
	} else {
		goto L775
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L1
	} else {
		goto L770
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4130 = m.ExcPending
	if v4130 != 0 {
		goto L1
	} else {
		goto L765
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4110 = m.ExcPending
	if v4110 != 0 {
		goto L1
	} else {
		goto L760
	}
L12:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_0), int32(0))
	mBase = m.M
	v4093 = m.ExcPending
	if v4093 != 0 {
		goto L1
	} else {
		goto L756
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4069 = m.ExcPending
	if v4069 != 0 {
		goto L1
	} else {
		goto L750
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L1
	} else {
		goto L747
	}
L15:
	;
	v3708 = F_palloc0(m, int32(72))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L1
	} else {
		goto L667
	}
L16:
	;
	v3563 = m.G0
	v3565 = v3563 - int32(96)
	m.G0 = v3565
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+44)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v3565)+48)) = int64(0)
	v3571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3571 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L17:
	;
	v1786 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1786)
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1788 == int32(0) {
		v2026 = v5
		v2031 = v5
		v2041 = v5
		goto L298
	} else {
		goto L299
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(4)
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1505)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1509 != 0 {
		goto L262
	} else {
		goto L263
	}
L19:
	;
	if v43 == int32(64) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v43 != int32(3) {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1457
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v1457
	*(*int32)(unsafe.Add(mBase, uint32(v37)+308)) = v1457
	v1493 = F_list_make1_impl(m, int32(1), v37+int32(12))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L259
	}
L23:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v1068 = F_palloc0(m, int32(136))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L205
	}
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if l0 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v1029 != 0 {
		v1457 = v1029
		goto L22
	} else {
		goto L204
	}
L26:
	;
	if v275 != 0 {
		goto L52
	} else {
		goto L53
	}
L27:
	;
	v60 = l0
	v65 = v5
	goto L30
L28:
	;
	goto L29
L29:
	;
	v275 = int32(0)
	goto L26
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)+36))
	if v88 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v205 != 0 {
		v60 = v205
		v65 = v65 + int32(1)
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v91 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v94 = int32(0)
	if v94 < v91 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v97 = v91
	goto L37
L36:
	;
	v97 = v94
	goto L37
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v105 = int32(0)
	goto L38
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v98+v105<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.B2i32(v141 == int32(0))|base.B2i32(v141 != v144) != 0 {
		v162 = v141
		v163 = v144
		goto L41
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(332)))) = v65
	v275 = v137
	goto L26
L40:
	;
	if v162-v163 != 0 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v147 = v138
	v148 = v51
	goto L43
L43:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v152 == int32(0) {
		v162 = v152
		v163 = v151
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v162 = v152
	v163 = v151
	goto L41
L45:
	;
	v155 = int32(1)
	if v152 == v151 {
		v147 = v147 + v155
		v148 = v148 + v155
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v166 = v105 + int32(1)
	if v97 != v166 {
		v105 = v166
		goto L38
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L39
L50:
	;
	goto L32
L51:
	;
	goto L31
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v37)+332))
	v277 = m.G0
	v279 = v277 - int32(32)
	m.G0 = v279
	v282 = F_palloc0(m, int32(136))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v629 = F_get_visible_ENR_metadata(m, v628, v627)
	mBase = m.M
	goto L143
L55:
	;
	v1029 = v532
	goto L25
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = int32(101)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v286 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v287 = v286
	goto L59
L58:
	;
	v287 = v275
	goto L59
L59:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = int32(6)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+88)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v282)+84)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v297 = base.B2i32(v295 != int32(67))
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+92)) = uint8(v297)
	if v297 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+36)) = v301 + int32(1)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v306 != int32(67) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L139
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L134
	}
L65:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v275)+44))
	v316 = F_list_copy(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v309 == int32(1) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v305)+96))
	if v312 == int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+96)) = v316
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v275)+48))
	v320 = F_list_copy(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+100)) = v320
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v275)+52))
	v324 = F_list_copy(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v282)+104)) = v324
	if v286 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v334 = int32(0)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	if v336 != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v328 = F_copyObjectImpl(m, v286)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v331 = F_makeAlias(m, v288, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v333 = v328
	goto L72
L77:
	;
	v333 = v331
	goto L72
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v338 = v337
	goto L80
L79:
	;
	v338 = v334
	goto L80
L80:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v275)+40))
	if v339 == int32(0) {
		v404 = v334
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v404 < v338 {
		goto L63
	} else {
		goto L91
	}
L82:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v342 <= int32(0) {
		v404 = v334
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v356 = v334
	v368 = v336
	goto L84
L84:
	;
	if v338 <= v356 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v404 = v390
	goto L81
L86:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380+v356<<(uint(int32(2))%32))))
	v385 = F_lappend(m, v368, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	v388 = v368
	goto L88
L88:
	;
	v390 = v356 + int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v390 < v391 {
		v356 = v390
		v368 = v388
		goto L84
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v385
	v388 = v385
	goto L88
L90:
	;
	goto L85
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v333
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	if v429 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	v432 = F_makeString(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v460 = int32(0)
	goto L94
L94:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	if v461 != 0 {
		goto L103
	} else {
		goto L104
	}
L95:
	;
	v434 = F_lappend(m, v430, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = v434
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v282)+96))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+8)))
	if v442 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v443 = int32(2249)
	goto L99
L98:
	;
	v443 = int32(2287)
	goto L99
L99:
	;
	v444 = F_lappend_oid(m, v438, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+96)) = v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v282)+100))
	v449 = F_lappend_int(m, v447, int32(-1))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+100)) = v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v282)+104))
	v454 = F_lappend_oid(m, v452, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+104)) = v454
	v460 = int32(1)
	goto L94
L103:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v465 = F_makeString(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	v517 = v460
	goto L105
L105:
	;
	v518 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+125)) = uint8(v518)
	v520 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+124)) = uint8(v520)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v523 = F_lappend(m, v522, v282)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L116
	}
L106:
	;
	v467 = F_lappend(m, v463, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+8)) = v467
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v282)+96))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+28))
	v474 = F_lappend_oid(m, v471, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+96)) = v474
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v282)+100))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+32))
	v480 = F_lappend_int(m, v477, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+100)) = v480
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v282)+104))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+36))
	v486 = F_lappend_oid(m, v483, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+104)) = v486
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)+8))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+20))
	v493 = F_makeString(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v495 = F_lappend(m, v490, v493)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v497)+8)) = v495
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v282)+96))
	v501 = F_lappend_oid(m, v499, int32(2287))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+96)) = v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v282)+100))
	v506 = F_lappend_int(m, v504, int32(-1))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+100)) = v506
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v282)+104))
	v511 = F_lappend_oid(m, v509, int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+104)) = v511
	v517 = v460 | int32(2)
	goto L105
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v523
	if v523 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v528 = v526
	goto L119
L118:
	;
	v528 = int32(0)
	goto L119
L119:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v282)+96))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v282)+100))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v282)+104))
	v532 = F_buildNSItemFromLists(m, v282, v528, v529, v530, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v282)+88))
	v535 = int32(0)
	if base.B2i32(v534 == v535)|base.B2i32(v517 == v535) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	m.G0 = v279 + int32(32)
	goto L55
L122:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	if v542 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v545 = v543
	goto L125
L124:
	;
	v545 = int32(0)
	goto L125
L125:
	;
	v551 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v540+v545<<(uint(int32(5))%32)-int32(2)))) = uint8(v551)
	if v517 == v551 {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	if v557 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	v560 = v558
	goto L129
L128:
	;
	v560 = int32(0)
	goto L129
L129:
	;
	v566 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v555+v560<<(uint(int32(5))%32)-int32(34)))) = uint8(v566)
	if v517 == int32(2) {
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	if v572 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v575 = v573
	goto L133
L132:
	;
	v575 = int32(0)
	goto L133
L133:
	;
	v581 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v570+v575<<(uint(int32(5))%32)-int32(66)))) = uint8(v581)
	goto L121
L134:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+16)) = v594
	F_errmsg(m, int32(_a_F_transformFromClauseItem_1), v279+int32(16))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_parser_errposition(m, l0, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2399), int32(_a_F_transformFromClauseItem_3))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
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
	F_errcode(m, int32(_a_F_transformFromClauseItem_4))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+8)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v288
	F_errmsg(m, int32(_a_F_transformFromClauseItem_5), v279)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2425), int32(_a_F_transformFromClauseItem_3))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	if base.B2i32(v629 != int32(0)) == int32(0) {
		goto L23
	} else {
		goto L144
	}
L144:
	;
	v634 = m.G0
	v636 = v634 - int32(32)
	m.G0 = v636
	v639 = F_palloc0(m, int32(136))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L146
	}
L145:
	;
	v1029 = v951
	goto L25
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = int32(101)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v647 = l1 + int32(12)
	if v643 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v648 = v643 + int32(4)
	goto L149
L148:
	;
	v648 = v647
	goto L149
L149:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v652 = int32(0)
	if v650 == v652 {
		v684 = v652
		goto L152
	} else {
		goto L153
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L201
	}
L151:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	if v687 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L152:
	;
	goto L151
L153:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v650)))
	if v657 == int32(0) {
		v684 = v652
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	if v660 <= int32(0) {
		v684 = v652
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v657)+12))
	v665 = int32(0)
	goto L156
L156:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v663+v665<<(uint(int32(2))%32))))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	v675 = F_strcmp(m, v674, v651)
	mBase = m.M
	if v675 == int32(0) {
		v684 = v673
		goto L152
	} else {
		goto L158
	}
L157:
	;
	v684 = int32(0)
	goto L152
L158:
	;
	v679 = v665 + int32(1)
	if v660 != v679 {
		v665 = v679
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+12)) = int32(7)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+16)) = v692
	v694 = F_ENRMetadataGetTupDesc(m, v684)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L198
	}
L163:
	;
	v697 = F_makeAlias(m, v649, int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+8)) = v697
	F_buildRelationAliases(m, v694, v643, v697)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+108)) = v702
	v704 = *(*float64)(unsafe.Add(mBase, uint32(v684)+16))
	v705 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v639)+104)) = v705
	*(*int64)(unsafe.Add(mBase, uint32(v639)+96)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v639)+112)) = v704
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	if v705 < v711 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v721 = v711
	v725 = int32(1)
	goto L169
L167:
	;
	goto L168
L168:
	;
	v831 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v639)+125)) = uint8(v831)
	v833 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v639)+124)) = uint8(v833)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v836 = F_lappend(m, v835, v639)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L183
	}
L169:
	;
	v753 = v694 + v721<<(uint(int32(3))%32) + v725*int32(100)
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+19)))
	if v754 == int32(1) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L168
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+104)) = v791
	v794 = v725 + int32(1)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	if v794 <= v795 {
		v721 = v795
		v725 = v794
		goto L169
	} else {
		goto L182
	}
L172:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v639)+96))
	v759 = F_lappend_oid(m, v757, int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v772 = v753 - int32(72)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+68))
	if v773 == int32(0) {
		goto L150
	} else {
		goto L178
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+96)) = v759
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v639)+100))
	v764 = F_lappend_int(m, v762, int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+100)) = v764
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v639)+104))
	v769 = F_lappend_oid(m, v767, int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v791 = v769
	goto L171
L178:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v639)+96))
	v777 = F_lappend_oid(m, v776, v773)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+96)) = v777
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v639)+100))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v772)+76))
	v782 = F_lappend_int(m, v780, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+100)) = v782
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v639)+104))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v772)+96))
	v787 = F_lappend_oid(m, v785, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v791 = v787
	goto L171
L182:
	;
	goto L170
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v836
	if v836 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	v840 = v839
	goto L186
L185:
	;
	v840 = v5
	goto L186
L186:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v844 = F_palloc0(m, v841<<(uint(int32(5))%32))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if int32(0) < v841 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v860 = int32(0)
	goto L191
L189:
	;
	goto L190
L190:
	;
	v951 = F_palloc(m, int32(28))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L197
	}
L191:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v889 = v694 + v883<<(uint(int32(3))%32) + v860*int32(100)
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889)+119)))
	if v890 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L190
L193:
	;
	v895 = v844 + v860<<(uint(int32(5))%32)
	v897 = v860 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v895)+4)) = uint16(v897)
	*(*int32)(unsafe.Add(mBase, uint32(v895))) = v840
	v901 = v889 + int32(28)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v901)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v895)+8)) = v902
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v901)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v895)+12)) = v904
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v901)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v895)+28)) = uint16(v897)
	*(*int32)(unsafe.Add(mBase, uint32(v895)+24)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v895)+16)) = v906
	goto L195
L194:
	;
	goto L195
L195:
	;
	v914 = v860 + int32(1)
	if v914 != v841 {
		v860 = v914
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v639)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v951)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+16)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v951)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+8)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v951)+4)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v951))) = v953
	m.G0 = v636 + int32(32)
	goto L145
L198:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+16)) = v969
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_6), v636+int32(16))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2528), int32(_a_F_transformFromClauseItem_7))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = v985
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_8), v636)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2568), int32(_a_F_transformFromClauseItem_7))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	goto L23
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068))) = int32(101)
	if v1064 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1076 = v1064 + int32(4)
	goto L208
L207:
	;
	v1076 = l1 + int32(12)
	goto L208
L208:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+4)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+12)) = int32(0)
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v1081 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1283 = F_parserOpenTable(m, l0, l1, v1254)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L235
	}
L210:
	;
	v1254 = int32(2)
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1083 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1254 = int32(1)
	goto L209
L214:
	;
	goto L215
L215:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+4))
	if v1088 <= int32(0) {
		v1254 = int32(1)
		goto L209
	} else {
		goto L216
	}
L216:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+12))
	v1103 = int32(0)
	goto L217
L217:
	;
	v1126 = int32(2)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1091+v1103<<(uint(v1126)%32))))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+4))
	if v1131 == int32(0) {
		v1254 = v1126
		goto L209
	} else {
		goto L219
	}
L218:
	;
	v1254 = v1245
	goto L209
L219:
	;
	if v1077 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1245 = int32(1)
	v1247 = v1103 + v1245
	if v1088 != v1247 {
		v1103 = v1247
		goto L217
	} else {
		goto L234
	}
L221:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+4))
	if v1136 <= int32(0) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+12))
	v1147 = int32(0)
	goto L223
L223:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1139+v1147<<(uint(int32(2))%32))))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+12))
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077))))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179))))
	if base.B2i32(v1182 == int32(0))|base.B2i32(v1182 != v1185) != 0 {
		v1203 = v1182
		v1204 = v1185
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L220
L225:
	;
	if v1203-v1204 == int32(0) {
		v1254 = v1126
		goto L209
	} else {
		goto L232
	}
L226:
	;
	goto L225
L227:
	;
	v1188 = v1077
	v1189 = v1179
	goto L228
L228:
	;
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+1)))
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1188)+1)))
	if v1193 == int32(0) {
		v1203 = v1193
		v1204 = v1192
		goto L226
	} else {
		goto L230
	}
L229:
	;
	v1203 = v1193
	v1204 = v1192
	goto L226
L230:
	;
	v1196 = int32(1)
	if v1193 == v1192 {
		v1188 = v1188 + v1196
		v1189 = v1189 + v1196
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1209 = v1147 + int32(1)
	if v1209 != v1136 {
		v1147 = v1209
		goto L223
	} else {
		goto L233
	}
L233:
	;
	goto L224
L234:
	;
	goto L218
L235:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v1068)+20)) = uint8(v1065)
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+16)) = v1285
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+48))
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+24)) = v1254
	*(*uint8)(unsafe.Add(mBase, uint32(v1068)+21)) = uint8(v1289)
	v1293 = F_makeAlias(m, v1077, int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+8)) = v1293
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+52))
	F_buildRelationAliases(m, v1296, v1064, v1293)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1068)+125)) = uint8(v1299)
	v1301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1068)+124)) = uint8(v1301)
	v1304 = F_palloc0(m, int32(40))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1304))) = int32(102)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+4)) = v1308
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304)+8)) = uint8(v1310)
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1313 = F_lappend(m, v1312, v1304)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1313
	if v1313 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	v1318 = v1316
	goto L242
L241:
	;
	v1318 = int32(0)
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+28)) = v1318
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+16)) = int64(2)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1323 = F_lappend(m, v1322, v1068)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1323
	v1326 = int32(0)
	if v1323 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+4))
	v1329 = v1328
	goto L246
L245:
	;
	v1329 = v1326
	goto L246
L246:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+52))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1334 = F_palloc0(m, v1331<<(uint(int32(5))%32))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	if int32(0) < v1331 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1344 = v1326
	goto L251
L249:
	;
	goto L250
L250:
	;
	v1440 = F_palloc(m, int32(28))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L257
	}
L251:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1378 = v1330 + v1372<<(uint(int32(3))%32) + v1344*int32(100)
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378)+119)))
	if v1379 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L250
L253:
	;
	v1384 = v1334 + v1344<<(uint(int32(5))%32)
	v1386 = v1344 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1384)+4)) = uint16(v1386)
	*(*int32)(unsafe.Add(mBase, uint32(v1384))) = v1329
	v1390 = v1378 + int32(28)
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1384)+8)) = v1391
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1384)+12)) = v1393
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v1384)+28)) = uint16(v1386)
	*(*int32)(unsafe.Add(mBase, uint32(v1384)+24)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v1384)+16)) = v1395
	goto L255
L254:
	;
	goto L255
L255:
	;
	v1403 = v1344 + int32(1)
	if v1403 != v1331 {
		v1344 = v1403
		goto L251
	} else {
		goto L256
	}
L256:
	;
	goto L252
L257:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1440)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+16)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+12)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+8)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v1440)+4)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = v1442
	F_relation_close(m, v1283, int32(0))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1457 = v1440
	goto L22
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1493
	v1497 = F_palloc0(m, int32(8))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1497))) = int32(63)
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1497)+4)) = v1501
	v7056 = v1497
	goto L3
L261:
	;
	v1753 = F_parse_sub_analyze(m, v1507, l0, int32(0), v1752)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L1
	} else {
		goto L292
	}
L262:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+4))
	v1512 = v1510
	goto L264
L263:
	;
	v1512 = int32(0)
	goto L264
L264:
	;
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v1515 != 0 {
		v1752 = int32(1)
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1516 = int32(0)
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1517 == v1516 {
		v1752 = v1516
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+4))
	if v1520 <= int32(0) {
		v1691 = v5
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1752 = v1691
	goto L261
L268:
	;
	v1523 = int32(0)
	if v1523 < v1520 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1526 = v1520
	goto L271
L270:
	;
	v1526 = v1523
	goto L271
L271:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+12))
	v1532 = int32(0)
	goto L272
L272:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1527+v1532<<(uint(int32(2))%32))))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	v1568 = base.B2i32(v1566 == int32(0))
	if v1566 == int32(0) {
		v1691 = v1568
		goto L267
	} else {
		goto L274
	}
L273:
	;
	v1691 = v1568
	goto L267
L274:
	;
	if v1512 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1682 = v1532 + int32(1)
	if v1682 != v1526 {
		v1532 = v1682
		goto L272
	} else {
		goto L291
	}
L276:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	if v1573 <= int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+12))
	v1584 = int32(0)
	goto L278
L278:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1576+v1584<<(uint(int32(2))%32))))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1615)+12))
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512))))
	v1622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1616))))
	if base.B2i32(v1619 == int32(0))|base.B2i32(v1619 != v1622) != 0 {
		v1640 = v1619
		v1641 = v1622
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1752 = int32(1)
	goto L261
L280:
	;
	if v1640-v1641 != 0 {
		goto L287
	} else {
		goto L288
	}
L281:
	;
	goto L280
L282:
	;
	v1625 = v1512
	v1626 = v1616
	goto L283
L283:
	;
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+1)))
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625)+1)))
	if v1630 == int32(0) {
		v1640 = v1630
		v1641 = v1629
		goto L281
	} else {
		goto L285
	}
L284:
	;
	v1640 = v1630
	v1641 = v1629
	goto L281
L285:
	;
	v1633 = int32(1)
	if v1630 == v1629 {
		v1625 = v1625 + v1633
		v1626 = v1626 + v1633
		goto L283
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	v1644 = v1584 + int32(1)
	if v1573 != v1644 {
		v1584 = v1644
		goto L278
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	goto L279
L290:
	;
	goto L275
L291:
	;
	goto L273
L292:
	;
	v1755 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1755
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v1755)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1753)))
	if v1759 != int32(67) {
		goto L8
	} else {
		goto L293
	}
L293:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1753)+4))
	if v1762 != int32(1) {
		goto L8
	} else {
		goto L294
	}
L294:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v1768 = F_addRangeTableEntryForSubquery(m, l0, v1753, v1765, v1766, int32(1))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1768
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v1768
	*(*int32)(unsafe.Add(mBase, uint32(v37)+304)) = v1768
	v1776 = F_list_make1_impl(m, int32(1), v37+int32(16))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1776
	v1780 = F_palloc0(m, int32(8))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1780))) = int32(63)
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1780)+4)) = v1784
	v7056 = v1780
	goto L3
L298:
	;
	v2054 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v2054)
	F_assign_list_collations(m, l0, v2026)
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L347
	}
L299:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	if v1791 <= int32(0) {
		v2026 = v5
		v2031 = v5
		v2041 = v5
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1800 = v5
	v1804 = v5
	v1805 = v5
	v1815 = v5
	goto L301
L301:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+12))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1828+v1804<<(uint(int32(2))%32))))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1832)+12))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1833)+4))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1835)))
	if v1836 != int32(76) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v2026 = v1988
	v2031 = v1993
	v2041 = v2003
	goto L298
L303:
	;
	v2017 = v1804 + int32(1)
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	if v2017 < v2018 {
		v1800 = v1988
		v1804 = v2017
		v1805 = v1993
		v1815 = v2003
		goto L301
	} else {
		goto L346
	}
L304:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1967 = F_transformExpr(m, l0, v1835, int32(5))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L336
	}
L305:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+4))
	if v1839 == int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+4))
	if v1842 != int32(1) {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+12))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1845)))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+4))
	v1848 = int32(_a_F_transformFromClauseItem_9)
	v1851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1847))))
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformFromClauseItem[0])))
	if base.B2i32(v1851 == int32(0))|base.B2i32(v1851 != v1854) != 0 {
		v1872 = v1851
		v1873 = v1854
		goto L309
	} else {
		goto L310
	}
L308:
	;
	if v1872-v1873 != 0 {
		goto L304
	} else {
		goto L315
	}
L309:
	;
	goto L308
L310:
	;
	v1857 = v1847
	v1858 = v1848
	goto L311
L311:
	;
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1858)+1)))
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1857)+1)))
	if v1862 == int32(0) {
		v1872 = v1862
		v1873 = v1861
		goto L309
	} else {
		goto L313
	}
L312:
	;
	v1872 = v1862
	v1873 = v1861
	goto L309
L313:
	;
	v1865 = int32(1)
	if v1862 == v1861 {
		v1857 = v1857 + v1865
		v1858 = v1858 + v1865
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+8))
	if v1875 == int32(0) {
		goto L304
	} else {
		goto L316
	}
L316:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1878 < int32(2) {
		goto L304
	} else {
		goto L317
	}
L317:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+12))
	if v1881 != 0 {
		goto L304
	} else {
		goto L318
	}
L318:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+16))
	if v1882 != 0 {
		goto L304
	} else {
		goto L319
	}
L319:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+20))
	if v1883 != 0 {
		goto L304
	} else {
		goto L320
	}
L320:
	;
	v1884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1835)+29)))
	if v1884 != 0 {
		goto L304
	} else {
		goto L321
	}
L321:
	;
	v1885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1835)+30)))
	if v1885 != 0 {
		goto L304
	} else {
		goto L322
	}
L322:
	;
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1835)+31)))
	if v1886|v1834 != 0 {
		goto L304
	} else {
		goto L323
	}
L323:
	;
	v1895 = v1800
	v1896 = int32(0)
	v1900 = v1805
	v1910 = v1815
	goto L324
L324:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+12))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1923+v1896<<(uint(int32(2))%32))))
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1930 = F_SystemFuncName(m, int32(_a_F_transformFromClauseItem_9))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L1
	} else {
		goto L326
	}
L325:
	;
	v1988 = v1950
	v1993 = v1957
	v2003 = v1954
	goto L303
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v1927
	*(*int32)(unsafe.Add(mBase, uint32(v37)+332)) = v1927
	v1937 = F_list_make1_impl(m, int32(1), v37+int32(28))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+36))
	v1941 = F_makeFuncCall(m, v1930, v1937, int32(0), v1940)
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v1944 = F_transformExpr(m, l0, v1941, int32(5))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if base.B2i32(v1946 != v1928)&base.B2i32(v1946 != v1944) != 0 {
		goto L9
	} else {
		goto L330
	}
L330:
	;
	v1950 = F_lappend(m, v1895, v1944)
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1952 = F_FigureColname(m, v1941)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1954 = F_lappend(m, v1910, v1952)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v1957 = F_lappend(m, v1900, int32(0))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1960 = v1896 + int32(1)
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	if v1960 < v1961 {
		v1895 = v1950
		v1896 = v1960
		v1900 = v1957
		v1910 = v1954
		goto L324
	} else {
		goto L335
	}
L335:
	;
	goto L325
L336:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if base.B2i32(v1965 != v1969)&base.B2i32(v1967 != v1969) != 0 {
		goto L10
	} else {
		goto L337
	}
L337:
	;
	v1973 = F_lappend(m, v1800, v1967)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v1975 = F_FigureColname(m, v1835)
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v1977 = F_lappend(m, v1815, v1975)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	if v1834 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1979 != 0 {
		goto L11
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1980 = F_lappend(m, v1805, v1834)
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L1
	} else {
		goto L345
	}
L344:
	;
	goto L343
L345:
	;
	v1988 = v1973
	v1993 = v1980
	v2003 = v1977
	goto L303
L346:
	;
	goto L302
L347:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v2058 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if v2026 != 0 {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v2099 = v2031
	goto L350
L350:
	;
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v2100 != 0 {
		goto L366
	} else {
		goto L367
	}
L351:
	;
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v2089 == int32(1) {
		goto L13
	} else {
		goto L363
	}
L352:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	if v2059 == int32(1) {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L1
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	if v2062 == int32(1) {
		goto L12
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_10), int32(0))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	F_errhint(m, int32(_a_F_transformFromClauseItem_11), int32(0))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2081 = F_exprLocation(m, v2080)
	mBase = m.M
	F_parser_errposition(m, l0, v2081)
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(652), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v2058
	*(*int32)(unsafe.Add(mBase, uint32(v37)+328)) = v2058
	v2097 = F_list_make1_impl(m, int32(1), v37+int32(24))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v2099 = v2097
	goto L350
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v3414
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v3414
	*(*int32)(unsafe.Add(mBase, uint32(v37)+300)) = v3414
	v3553 = F_list_make1_impl(m, int32(1), v37+int32(20))
	mBase = m.M
	v3554 = m.ExcPending
	if v3554 != 0 {
		goto L1
	} else {
		goto L635
	}
L366:
	;
	v2105 = int32(1)
	goto L368
L367:
	;
	v2103 = F_contain_vars_of_level(m, v2026, int32(0))
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L1
	} else {
		goto L369
	}
L368:
	;
	v2107 = m.G0
	v2109 = v2107 - int32(80)
	m.G0 = v2109
	v2112 = F_palloc0(m, int32(136))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L1
	} else {
		goto L370
	}
L369:
	;
	v2105 = v2103
	goto L368
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2112))) = int32(101)
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2026 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	v2118 = v2117
	goto L373
L372:
	;
	v2118 = v5
	goto L373
L373:
	;
	v2119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+68)) = v2119
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+36)) = v2119
	*(*int64)(unsafe.Add(mBase, uint32(v2112)+12)) = int64(3)
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+4)) = v2116
	*(*uint8)(unsafe.Add(mBase, uint32(v2112)+72)) = uint8(v2125)
	if v2116 != 0 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v2131 = v2116 + int32(4)
	goto L376
L375:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+12))
	v2131 = v2130
	goto L376
L376:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2131)))
	v2134 = F_makeAlias(m, v2132, int32(0))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+8)) = v2134
	v2141 = base.B2i32(v2116 == int32(0)) | base.B2i32(v2118 != int32(1))
	v2143 = F_palloc_mul(m, int32(4), v2118)
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v2150 = int32(0)
	v2172 = v5
	goto L386
L379:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3529 = m.ExcPending
	if v3529 != 0 {
		goto L1
	} else {
		goto L630
	}
L380:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L1
	} else {
		goto L625
	}
L381:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L1
	} else {
		goto L619
	}
L382:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L1
	} else {
		goto L614
	}
L383:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L1
	} else {
		goto L609
	}
L384:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_14), int32(0))
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L1
	} else {
		goto L606
	}
L385:
	;
	F_buildRelationAliases(m, v3261, v2116, v2134)
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L1
	} else {
		goto L590
	}
L386:
	;
	v2179 = int32(0)
	if v2026 == v2179 {
		v2189 = v2179
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v2984 = v2214 + v2215
	if int32(1665) <= v2984 {
		goto L380
	} else {
		goto L552
	}
L388:
	;
	v2190 = int32(0)
	if v2041 == v2190 {
		v2201 = v2190
		goto L391
	} else {
		goto L392
	}
L389:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	if v2183 <= v2172 {
		v2189 = int32(0)
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+12))
	v2189 = v2185 + v2172<<(uint(int32(2))%32)
	goto L388
L391:
	;
	if v2099 != 0 {
		goto L396
	} else {
		goto L397
	}
L392:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+4))
	if v2195 <= v2172 {
		v2201 = int32(0)
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+12))
	v2201 = v2197 + v2172<<(uint(int32(2))%32)
	goto L391
L394:
	;
	goto L387
L395:
	;
	v2222 = v2172 << (uint(int32(2)) % 32)
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2212+v2222)))
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2201)))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2189)))
	v2228 = F_palloc0(m, int32(32))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L1
	} else {
		goto L404
	}
L396:
	;
	v2202 = int32(0)
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+4))
	if base.B2i32(v2201 == v2202)|(base.B2i32(v2189 == v2202)|base.B2i32(v2206 <= v2172)) == v2202 {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	v2214 = v2190
	goto L398
L398:
	;
	v2215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v2215|base.B2i32(int32(1) < v2118) != 0 {
		goto L394
	} else {
		goto L403
	}
L399:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+12))
	if v2212 != 0 {
		goto L395
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v2214 = v2150
	goto L398
L402:
	;
	goto L401
L403:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+76)) = v2219
	v3261 = v2219
	goto L385
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2228))) = int32(103)
	v2232 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2228)+12)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+4)) = v2226
	*(*int64)(unsafe.Add(mBase, uint32(v2228)+20)) = v2232
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+28)) = int32(0)
	v2243 = F_get_expr_result_type(m, v2226, v2109+int32(72), v2109+int32(76))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	if v2224 != 0 {
		goto L408
	} else {
		goto L409
	}
L406:
	;
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2970)))
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+8)) = v2971
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+68))
	v2974 = F_lappend(m, v2973, v2228)
	mBase = m.M
	v2975 = m.ExcPending
	if v2975 != 0 {
		goto L1
	} else {
		goto L551
	}
L407:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+4))
	if int32(1601) <= v2702 {
		goto L382
	} else {
		goto L512
	}
L408:
	;
	if v2243 == int32(3) {
		goto L407
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	if v2243 == int32(3) {
		goto L383
	} else {
		goto L427
	}
L411:
	;
	v2247 = int32(1)
	if base.Ui32(v2243-v2247) <= base.Ui32(v2247) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2251 = F_exprType(m, v2226)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L1
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L1
	} else {
		goto L422
	}
L415:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	if v2251 == int32(2249) {
		goto L384
	} else {
		goto L418
	}
L418:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_15), int32(0))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	v2266 = F_exprLocation(m, v2224)
	mBase = m.M
	F_parser_errposition(m, l0, v2266)
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1878), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_17), int32(0))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v2285 = F_exprLocation(m, v2224)
	mBase = m.M
	F_parser_errposition(m, l0, v2285)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1885), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	if base.Ui32(v2243-int32(1)) < base.Ui32(int32(2)) {
		goto L406
	} else {
		goto L428
	}
L428:
	;
	if v2243 != 0 {
		goto L381
	} else {
		goto L429
	}
L429:
	;
	v2300 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+76)) = v2300
	if v2226 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L431:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+72))
	v2597 = F_exprTypmod(m, v2226)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L489
	}
L432:
	;
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2116)+4))
	v2584 = v2560
	goto L431
L433:
	;
	if v2141 != 0 {
		v2584 = v2225
		goto L431
	} else {
		goto L488
	}
L434:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2226)))
	if v2305 != int32(15) {
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+4))
	v2309 = int32(0)
	v2312 = m.G0
	v2314 = v2312 - int32(32)
	m.G0 = v2314
	v2318 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v2308))
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L1
	} else {
		goto L439
	}
L436:
	;
	v2520 = int32(0)
	if v2141|base.B2i32(v2448 != v2520) == v2520 {
		goto L432
	} else {
		goto L484
	}
L437:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L1
	} else {
		goto L481
	}
L438:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L1
	} else {
		goto L478
	}
L439:
	;
	if v2318 != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v2322 = F_heap_attisnull(m, v2318, int32(22), int32(0))
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L1
	} else {
		goto L444
	}
L441:
	;
	goto L442
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L1
	} else {
		goto L475
	}
L443:
	;
	F_ReleaseCatCache(m, v2318)
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L1
	} else {
		goto L474
	}
L444:
	;
	if v2322 != 0 {
		v2448 = v2309
		goto L443
	} else {
		goto L445
	}
L445:
	;
	v2326 = F_heap_attisnull(m, v2318, int32(23), int32(0))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	if v2326 != 0 {
		v2448 = v2309
		goto L443
	} else {
		goto L447
	}
L447:
	;
	v2330 = F_SysCacheGetAttrNotNull(m, int32(47), v2318, int32(22))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v2334 = F_SysCacheGetAttrNotNull(m, int32(47), v2318, int32(23))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v2337 = F_pg_detoast_datum(m, base.I32_wrap_i64(v2330))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+4))
	if v2339 != int32(1) {
		goto L438
	} else {
		goto L451
	}
L451:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+16))
	if v2342 < int32(0) {
		goto L438
	} else {
		goto L452
	}
L452:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+8))
	if v2345 != 0 {
		goto L438
	} else {
		goto L453
	}
L453:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+12))
	if v2346 != int32(18) {
		goto L438
	} else {
		goto L454
	}
L454:
	;
	v2350 = F_pg_detoast_datum(m, base.I32_wrap_i64(v2334))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2350)+4))
	if v2352 != int32(1) {
		goto L437
	} else {
		goto L456
	}
L456:
	;
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2350)+16))
	if v2355 != v2342 {
		goto L437
	} else {
		goto L457
	}
L457:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2350)+8))
	if v2357 != 0 {
		goto L437
	} else {
		goto L458
	}
L458:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2350)+12))
	if v2358 != int32(25) {
		goto L437
	} else {
		goto L459
	}
L459:
	;
	F_deconstruct_array_builtin(m, v2350, int32(25), v2314+int32(28), int32(0), v2314+int32(24))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	if v2342 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2448 = v2309
	goto L443
L462:
	;
	goto L463
L463:
	;
	v2381 = v2309
	v2389 = v2309
	v2399 = int32(0)
	goto L464
L464:
	;
	v2409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337+int32(24)+v2399))))
	v2411 = v2409 - int32(105)
	v2412 = int32(0)
	if base.B2i32(v2411 == v2412)|base.B2i32(v2411 == int32(13)) == v2412 {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v2448 = v2433
	goto L443
L466:
	;
	v2419 = int32(0)
	if v2389 != 0 {
		v2448 = v2419
		goto L443
	} else {
		goto L469
	}
L467:
	;
	v2433 = v2381
	v2435 = v2389
	goto L468
L468:
	;
	v2437 = v2399 + int32(1)
	if v2437 != v2342 {
		v2381 = v2433
		v2389 = v2435
		v2399 = v2437
		goto L464
	} else {
		goto L473
	}
L469:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+28))
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v2420+v2399<<(uint(int32(3))%32))))
	v2425 = F_text_to_cstring(m, v2424)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	if v2425 == int32(0) {
		v2448 = v2419
		goto L443
	} else {
		goto L471
	}
L471:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2425))))
	if v2430 == int32(0) {
		v2448 = v2419
		goto L443
	} else {
		goto L472
	}
L472:
	;
	v2433 = v2425
	v2435 = int32(1)
	goto L468
L473:
	;
	goto L465
L474:
	;
	m.G0 = v2314 + int32(32)
	goto L436
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314))) = v2308
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_18), v2314)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1627), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_21), int32(0))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1653), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+16)) = v2342
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_22), v2314+int32(16))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_19), int32(1661), int32(_a_F_transformFromClauseItem_20))
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	if v2448 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2525 = v2448
	goto L487
L486:
	;
	v2525 = v2225
	goto L487
L487:
	;
	v2584 = v2525
	goto L431
L488:
	;
	goto L432
L489:
	;
	F_TupleDescInitEntry(m, v2300, int32(1), v2584, v2596, v2597, int32(0))
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v2602 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2604 = F_exprCollation(m, v2226)
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2602)))
	*(*int32)(unsafe.Add(mBase, uint32(v2602+v2606<<(uint(int32(3))%32)+int32(100))+24)) = v2604
	goto L492
L492:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2615 = int32(0)
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v2614)))
	if v2615 < v2624 {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	goto L406
L494:
	;
	v2628 = v2614 + int32(28)
	v2635 = v2615
	v2636 = v2624
	v2638 = v2615
	goto L498
L495:
	;
	v2692 = v2615
	v2699 = v2624
	goto L496
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2614)+20)) = v2699
	*(*int32)(unsafe.Add(mBase, uint32(v2614)+16)) = v2692
	goto L493
L497:
	;
	v2692 = v2686
	v2699 = v2665
	goto L496
L498:
	;
	v2644 = v2628 + v2624<<(uint(int32(3))%32) + v2635*int32(100)
	v2647 = v2628 + v2635<<(uint(int32(3))%32)
	if v2624 != v2636 {
		v2665 = v2636
		goto L500
	} else {
		goto L501
	}
L499:
	;
	v2686 = v2624
	goto L497
L500:
	;
	v2666 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2647)+2)))
	if v2666 <= int32(0) {
		v2686 = v2635
		goto L497
	} else {
		goto L508
	}
L501:
	;
	v2649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+7)))
	if v2649 != int32(118) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2665 = v2635
	goto L500
L503:
	;
	v2652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+4)))
	if v2652 != int32(1) {
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v2655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+6)))
	if v2655&int32(6) != 0 {
		goto L502
	} else {
		goto L505
	}
L505:
	;
	v2658 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2647)+2)))
	if v2658 <= int32(0) {
		goto L502
	} else {
		goto L506
	}
L506:
	;
	v2661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2644)+90)))
	if v2661 != int32(118) {
		v2665 = v2624
		goto L500
	} else {
		goto L507
	}
L507:
	;
	goto L502
L508:
	;
	v2669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2644)+90)))
	if v2669 == int32(118) {
		v2686 = v2635
		goto L497
	} else {
		goto L509
	}
L509:
	;
	v2672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+5)))
	v2678 = (v2638 + v2672 - int32(1)) & (int32(0) - v2672)
	if int32(_a_F_transformFromClauseItem_23) < v2678 {
		v2686 = v2635
		goto L497
	} else {
		goto L510
	}
L510:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2647))) = uint16(v2678)
	v2684 = v2635 + int32(1)
	if v2684 != v2624 {
		v2635 = v2684
		v2636 = v2665
		v2638 = v2678 + v2666
		goto L498
	} else {
		goto L511
	}
L511:
	;
	goto L499
L512:
	;
	v2705 = F_CreateTemplateTupleDesc(m, v2702)
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+76)) = v2705
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+4))
	if int32(0) < v2708 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2713 = int32(1)
	v2720 = int32(0)
	goto L517
L515:
	;
	v2843 = v2705
	goto L516
L516:
	;
	v2844 = int32(0)
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2843)))
	if v2844 < v2853 {
		goto L532
	} else {
		goto L533
	}
L517:
	;
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+12))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2747+v2720<<(uint(int32(2))%32))))
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+4))
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+8))
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2753)+12)))
	if v2754 != 0 {
		goto L379
	} else {
		goto L519
	}
L518:
	;
	v2808 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2843 = v2808
	goto L516
L519:
	;
	F_typenameTypeIdAndMod(m, l0, v2753, v2109+int32(68), v2109-int32(-64))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+68))
	v2762 = F_GetColumnDefCollation(m, l0, v2751, v2761)
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2765 = base.I32_extend16_s(v2713)
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+68))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+64))
	F_TupleDescInitEntry(m, v2764, v2765, v2752, v2766, v2767, int32(0))
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2771)))
	*(*int32)(unsafe.Add(mBase, uint32(v2771+v2772<<(uint(int32(3))%32)+v2765*int32(100))+24)) = v2762
	goto L523
L523:
	;
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+12))
	v2781 = F_pstrdup(m, v2752)
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	v2783 = F_makeString(m, v2781)
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v2785 = F_lappend(m, v2780, v2783)
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+12)) = v2785
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+16))
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+68))
	v2790 = F_lappend_oid(m, v2788, v2789)
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+16)) = v2790
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+20))
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+64))
	v2795 = F_lappend_int(m, v2793, v2794)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+20)) = v2795
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+24))
	v2799 = F_lappend_oid(m, v2798, v2762)
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2228)+24)) = v2799
	v2802 = int32(1)
	v2805 = v2720 + v2802
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+4))
	if v2805 < v2806 {
		v2713 = v2713 + v2802
		v2720 = v2805
		goto L517
	} else {
		goto L530
	}
L530:
	;
	goto L518
L531:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	F_CheckAttributeNamesTypes(m, v2931, int32(99), int32(2))
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L550
	}
L532:
	;
	v2857 = v2843 + int32(28)
	v2864 = v2844
	v2865 = v2853
	v2867 = v2844
	goto L536
L533:
	;
	v2921 = v2844
	v2928 = v2853
	goto L534
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2843)+20)) = v2928
	*(*int32)(unsafe.Add(mBase, uint32(v2843)+16)) = v2921
	goto L531
L535:
	;
	v2921 = v2915
	v2928 = v2894
	goto L534
L536:
	;
	v2873 = v2857 + v2853<<(uint(int32(3))%32) + v2864*int32(100)
	v2876 = v2857 + v2864<<(uint(int32(3))%32)
	if v2853 != v2865 {
		v2894 = v2865
		goto L538
	} else {
		goto L539
	}
L537:
	;
	v2915 = v2853
	goto L535
L538:
	;
	v2895 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2876)+2)))
	if v2895 <= int32(0) {
		v2915 = v2864
		goto L535
	} else {
		goto L546
	}
L539:
	;
	v2878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2876)+7)))
	if v2878 != int32(118) {
		goto L540
	} else {
		goto L541
	}
L540:
	;
	v2894 = v2864
	goto L538
L541:
	;
	v2881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2876)+4)))
	if v2881 != int32(1) {
		goto L540
	} else {
		goto L542
	}
L542:
	;
	v2884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2876)+6)))
	if v2884&int32(6) != 0 {
		goto L540
	} else {
		goto L543
	}
L543:
	;
	v2887 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2876)+2)))
	if v2887 <= int32(0) {
		goto L540
	} else {
		goto L544
	}
L544:
	;
	v2890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2873)+90)))
	if v2890 != int32(118) {
		v2894 = v2853
		goto L538
	} else {
		goto L545
	}
L545:
	;
	goto L540
L546:
	;
	v2898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2873)+90)))
	if v2898 == int32(118) {
		v2915 = v2864
		goto L535
	} else {
		goto L547
	}
L547:
	;
	v2901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2876)+5)))
	v2907 = (v2867 + v2901 - int32(1)) & (int32(0) - v2901)
	if int32(_a_F_transformFromClauseItem_23) < v2907 {
		v2915 = v2864
		goto L535
	} else {
		goto L548
	}
L548:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2876))) = uint16(v2907)
	v2913 = v2864 + int32(1)
	if v2913 != v2853 {
		v2864 = v2913
		v2865 = v2894
		v2867 = v2907 + v2895
		goto L536
	} else {
		goto L549
	}
L549:
	;
	goto L537
L550:
	;
	goto L406
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+68)) = v2974
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2222+v2143))) = v2978
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2978)))
	v2150 = v2982 + v2150
	v2172 = v2172 + int32(1)
	goto L386
L552:
	;
	v2987 = F_CreateTemplateTupleDesc(m, v2984)
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+76)) = v2987
	if int32(0) < v2118 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2993 = int32(0)
	v2995 = v2993
	v3004 = v2993
	goto L557
L555:
	;
	v3127 = v2987
	v3130 = int32(1)
	goto L556
L556:
	;
	v3157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v3157 == int32(1) {
		goto L567
	} else {
		goto L568
	}
L557:
	;
	v3032 = v2143 + v3004<<(uint(int32(2))%32)
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3032)))
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3033)))
	if int32(0) < v3034 {
		goto L559
	} else {
		goto L560
	}
L558:
	;
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v3127 = v3122
	v3130 = v3083 + int32(1)
	goto L556
L559:
	;
	v3037 = v2995
	v3041 = int32(1)
	v3044 = v3033
	goto L562
L560:
	;
	v3083 = v2995
	goto L561
L561:
	;
	v3118 = v3004 + int32(1)
	if v3118 != v2118 {
		v2995 = v3083
		v3004 = v3118
		goto L557
	} else {
		goto L566
	}
L562:
	;
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v3073 = v3037 + int32(1)
	F_TupleDescCopyEntry(m, v3071, base.I32_extend16_s(v3073), v3044, base.I32_extend16_s(v3041))
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L1
	} else {
		goto L564
	}
L563:
	;
	v3083 = v3073
	goto L561
L564:
	;
	v3079 = v3041 + int32(1)
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(v3032)))
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v3080)))
	if v3079 <= v3081 {
		v3037 = v3073
		v3041 = v3079
		v3044 = v3080
		goto L562
	} else {
		goto L565
	}
L565:
	;
	goto L563
L566:
	;
	goto L558
L567:
	;
	F_TupleDescInitEntry(m, v3127, base.I32_extend16_s(v3130), int32(_a_F_transformFromClauseItem_24), int32(20), int32(-1), int32(0))
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		goto L1
	} else {
		goto L570
	}
L568:
	;
	v3168 = v3127
	goto L569
L569:
	;
	v3169 = int32(0)
	v3178 = *(*int32)(unsafe.Add(mBase, uint32(v3168)))
	if v3169 < v3178 {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v3168 = v3167
	goto L569
L571:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v3261 = v3256
	goto L385
L572:
	;
	v3182 = v3168 + int32(28)
	v3189 = v3169
	v3190 = v3178
	v3192 = v3169
	goto L576
L573:
	;
	v3246 = v3169
	v3253 = v3178
	goto L574
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+20)) = v3253
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+16)) = v3246
	goto L571
L575:
	;
	v3246 = v3240
	v3253 = v3219
	goto L574
L576:
	;
	v3198 = v3182 + v3178<<(uint(int32(3))%32) + v3189*int32(100)
	v3201 = v3182 + v3189<<(uint(int32(3))%32)
	if v3178 != v3190 {
		v3219 = v3190
		goto L578
	} else {
		goto L579
	}
L577:
	;
	v3240 = v3178
	goto L575
L578:
	;
	v3220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3201)+2)))
	if v3220 <= int32(0) {
		v3240 = v3189
		goto L575
	} else {
		goto L586
	}
L579:
	;
	v3203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3201)+7)))
	if v3203 != int32(118) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v3219 = v3189
	goto L578
L581:
	;
	v3206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3201)+4)))
	if v3206 != int32(1) {
		goto L580
	} else {
		goto L582
	}
L582:
	;
	v3209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3201)+6)))
	if v3209&int32(6) != 0 {
		goto L580
	} else {
		goto L583
	}
L583:
	;
	v3212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3201)+2)))
	if v3212 <= int32(0) {
		goto L580
	} else {
		goto L584
	}
L584:
	;
	v3215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3198)+90)))
	if v3215 != int32(118) {
		v3219 = v3178
		goto L578
	} else {
		goto L585
	}
L585:
	;
	goto L580
L586:
	;
	v3223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3198)+90)))
	if v3223 == int32(118) {
		v3240 = v3189
		goto L575
	} else {
		goto L587
	}
L587:
	;
	v3226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3201)+5)))
	v3232 = (v3192 + v3226 - int32(1)) & (int32(0) - v3226)
	if int32(_a_F_transformFromClauseItem_23) < v3232 {
		v3240 = v3189
		goto L575
	} else {
		goto L588
	}
L588:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3201))) = uint16(v3232)
	v3238 = v3189 + int32(1)
	if v3238 != v3178 {
		v3189 = v3238
		v3190 = v3219
		v3192 = v3232 + v3220
		goto L576
	} else {
		goto L589
	}
L589:
	;
	goto L577
L590:
	;
	v3293 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2112)+125)) = uint8(v3293)
	*(*uint8)(unsafe.Add(mBase, uint32(v2112)+124)) = uint8(v2105)
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3297 = F_lappend(m, v3296, v2112)
	mBase = m.M
	v3298 = m.ExcPending
	if v3298 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v3297
	v3300 = int32(0)
	if v3297 != 0 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v3297)+4))
	v3303 = v3302
	goto L594
L593:
	;
	v3303 = v3300
	goto L594
L594:
	;
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+76))
	v3305 = *(*int32)(unsafe.Add(mBase, uint32(v3304)))
	v3308 = F_palloc0(m, v3305<<(uint(int32(5))%32))
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	if int32(0) < v3305 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v3316 = v3300
	goto L599
L597:
	;
	goto L598
L598:
	;
	v3414 = F_palloc(m, int32(28))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L1
	} else {
		goto L605
	}
L599:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(v3304)))
	v3352 = v3304 + v3346<<(uint(int32(3))%32) + v3316*int32(100)
	v3353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3352)+119)))
	if v3353 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L600:
	;
	goto L598
L601:
	;
	v3358 = v3308 + v3316<<(uint(int32(5))%32)
	v3360 = v3316 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3358)+4)) = uint16(v3360)
	*(*int32)(unsafe.Add(mBase, uint32(v3358))) = v3303
	v3364 = v3352 + int32(28)
	v3365 = *(*int32)(unsafe.Add(mBase, uint32(v3364)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v3358)+8)) = v3365
	v3367 = *(*int32)(unsafe.Add(mBase, uint32(v3364)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3358)+12)) = v3367
	v3369 = *(*int32)(unsafe.Add(mBase, uint32(v3364)+96))
	*(*uint16)(unsafe.Add(mBase, uint32(v3358)+28)) = uint16(v3360)
	*(*int32)(unsafe.Add(mBase, uint32(v3358)+24)) = v3303
	*(*int32)(unsafe.Add(mBase, uint32(v3358)+16)) = v3369
	goto L603
L602:
	;
	goto L603
L603:
	;
	v3377 = v3316 + int32(1)
	if v3377 != v3305 {
		v3316 = v3377
		goto L599
	} else {
		goto L604
	}
L604:
	;
	goto L600
L605:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3414)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v3414)+16)) = v3308
	*(*int32)(unsafe.Add(mBase, uint32(v3414)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3414)+8)) = v3303
	*(*int32)(unsafe.Add(mBase, uint32(v3414)+4)) = v2112
	*(*int32)(unsafe.Add(mBase, uint32(v3414))) = v3416
	m.G0 = v2109 + int32(80)
	goto L365
L606:
	;
	v3432 = F_exprLocation(m, v2224)
	mBase = m.M
	F_parser_errposition(m, l0, v3432)
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1872), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L609:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3446 = m.ExcPending
	if v3446 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_25), int32(0))
	mBase = m.M
	v3450 = m.ExcPending
	if v3450 != 0 {
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v3451 = F_exprLocation(m, v2226)
	mBase = m.M
	F_parser_errposition(m, l0, v3451)
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1895), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3458 = m.ExcPending
	if v3458 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L614:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+32)) = int32(1600)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_26), v2109+int32(32))
	mBase = m.M
	v3472 = m.ExcPending
	if v3472 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v3473 = F_exprLocation(m, v2224)
	mBase = m.M
	F_parser_errposition(m, l0, v3473)
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1935), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L619:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	v3488 = *(*int32)(unsafe.Add(mBase, uint32(v2109)+72))
	v3489 = F_format_type_be(m, v3488)
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+20)) = v3489
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+16)) = v2225
	F_errmsg(m, int32(_a_F_transformFromClauseItem_27), v2109+int32(16))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	v3498 = F_exprLocation(m, v2226)
	mBase = m.M
	F_parser_errposition(m, l0, v3498)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1995), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L625:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v3512 = m.ExcPending
	if v3512 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109))) = int32(1664)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_28), v2109)
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	v3518 = F_exprLocation(m, v2026)
	mBase = m.M
	F_parser_errposition(m, l0, v3518)
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(2023), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3525 = m.ExcPending
	if v3525 != 0 {
		goto L1
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2109)+48)) = v2752
	F_errmsg(m, int32(_a_F_transformFromClauseItem_29), v2109+int32(48))
	mBase = m.M
	v3538 = m.ExcPending
	if v3538 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+64))
	F_parser_errposition(m, l0, v3539)
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_2), int32(1952), int32(_a_F_transformFromClauseItem_16))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v3553
	v3557 = F_palloc0(m, int32(8))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3557))) = int32(63)
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3414)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3557)+4)) = v3561
	v7056 = v3557
	goto L3
L637:
	;
	v7038 = v3675
	goto L4
L638:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3683 = m.ExcPending
	if v3683 != 0 {
		goto L1
	} else {
		goto L661
	}
L639:
	;
	v3582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+60)) = v3582
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v3567)+8))
	if v3584 == v3582 {
		goto L643
	} else {
		goto L644
	}
L640:
	;
	v3574 = *(*int32)(unsafe.Add(mBase, uint32(v3571)+4))
	if base.Ui32(v3574-int32(1)) < base.Ui32(int32(2)) {
		goto L639
	} else {
		goto L641
	}
L641:
	;
	if v3574 != int32(6) {
		goto L638
	} else {
		goto L642
	}
L642:
	;
	goto L639
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+60)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+16)) = int32(0)
	v3592 = v3565 - int32(-64)
	v3597 = F_pg_snprintf(m, v3592, int32(32), int32(_a_F_transformFromClauseItem_30), v3565+int32(16))
	mBase = m.M
	v3598 = m.ExcPending
	if v3598 != 0 {
		goto L1
	} else {
		goto L646
	}
L644:
	;
	v3606 = v3584
	goto L645
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+12)) = v3606
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+40)) = v3606
	v3612 = F_list_make1_impl(m, int32(1), v3565+int32(12))
	mBase = m.M
	v3613 = m.ExcPending
	if v3613 != 0 {
		goto L1
	} else {
		goto L649
	}
L646:
	;
	v3600 = F_pstrdup(m, v3592)
	mBase = m.M
	v3601 = m.ExcPending
	if v3601 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	v3602 = F_lappend(m, int32(0), v3600)
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3567)+8)) = v3600
	v3606 = v3600
	goto L645
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+56)) = v3612
	v3616 = v3565 + int32(44)
	v3617 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_CheckDuplicateColumnOrPathNames(m, v3616, v3617)
	mBase = m.M
	v3619 = m.ExcPending
	if v3619 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	v3620 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3620)
	v3623 = F_palloc0(m, int32(72))
	mBase = m.M
	v3624 = m.ExcPending
	if v3624 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3623))) = int64(4294967300)
	v3628 = F_palloc0(m, int32(48))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L652
	}
L652:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3628))) = int64(12884902010)
	v3632 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+12)) = v3632
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3567)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+16)) = v3634
	v3636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+20)) = v3636
	v3640 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+32)) = v3640
	v3642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3628)+44)) = v3642
	v3645 = F_transformExpr(m, l0, v3628, int32(5))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3623)+16)) = v3645
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+52)) = v3623
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+48)) = l1
	v3650 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3652 = F_transformJsonTableColumns(m, v3616, v3650, v3651, v3567)
	mBase = m.M
	v3653 = m.ExcPending
	if v3653 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3623)+60)) = v3652
	v3655 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+16))
	v3656 = *(*int32)(unsafe.Add(mBase, uint32(v3655)+32))
	v3657 = F_copyObjectImpl(m, v3656)
	mBase = m.M
	v3658 = m.ExcPending
	if v3658 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3623)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3623)+52)) = v3657
	v3662 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3623)+68)) = v3662
	v3664 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3664)
	v3667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v3667 == v3664 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3671 = F_contain_vars_of_level(m, v3623, int32(0))
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L1
	} else {
		goto L659
	}
L657:
	;
	v3673 = int32(1)
	goto L658
L658:
	;
	v3674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3675 = F_addRangeTableEntryForTableFunc(m, l0, v3623, v3674, v3673)
	mBase = m.M
	v3676 = m.ExcPending
	if v3676 != 0 {
		goto L1
	} else {
		goto L660
	}
L659:
	;
	v3673 = v3671
	goto L658
L660:
	;
	m.G0 = v3565 + int32(96)
	goto L637
L661:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+32)) = int32(_a_F_transformFromClauseItem_31)
	F_errmsg(m, int32(_a_F_transformFromClauseItem_32), v3565+int32(32))
	mBase = m.M
	v3693 = m.ExcPending
	if v3693 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v3696 = F_errdetail(m, int32(_a_F_transformFromClauseItem_33), int32(0))
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3699 = *(*int32)(unsafe.Add(mBase, uint32(v3698)+16))
	F_parser_errposition(m, l0, v3699)
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_34), int32(94), int32(_a_F_transformFromClauseItem_35))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L1
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
	*(*int64)(unsafe.Add(mBase, uint32(v3708))) = int64(4)
	v3712 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v3712)
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3716 = F_transformExpr(m, l0, v3714, int32(5))
	mBase = m.M
	v3717 = m.ExcPending
	if v3717 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	v3720 = F_coerce_to_specific_type(m, l0, v3716, int32(25), int32(_a_F_transformFromClauseItem_36))
	mBase = m.M
	v3721 = m.ExcPending
	if v3721 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+20)) = v3720
	F_assign_expr_collations(m, l0, v3720)
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	v3725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v3727 = F_transformExpr(m, l0, v3725, int32(5))
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	v3731 = F_coerce_to_specific_type(m, l0, v3727, int32(142), int32(_a_F_transformFromClauseItem_36))
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+16)) = v3731
	F_assign_expr_collations(m, l0, v3731)
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+64)) = int32(-1)
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3739 != 0 {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v3740 = *(*int32)(unsafe.Add(mBase, uint32(v3739)+4))
	v3742 = v3740
	goto L676
L675:
	;
	v3742 = int32(0)
	goto L676
L676:
	;
	v3743 = F_palloc_mul(m, int32(4), v3742)
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L1
	} else {
		goto L677
	}
L677:
	;
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v3745 == int32(0) {
		goto L5
	} else {
		goto L678
	}
L678:
	;
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3745)+4))
	if v3748 <= int32(0) {
		goto L5
	} else {
		goto L679
	}
L679:
	;
	v3761 = v5
	goto L680
L680:
	;
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+24))
	v3787 = v3761 << (uint(int32(2)) % 32)
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v3745)+12))
	v3790 = *(*int32)(unsafe.Add(mBase, uint32(v3787+v3788)))
	v3791 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+4))
	v3792 = F_pstrdup(m, v3791)
	mBase = m.M
	v3793 = m.ExcPending
	if v3793 != 0 {
		goto L1
	} else {
		goto L682
	}
L681:
	;
	goto L5
L682:
	;
	v3794 = F_makeString(m, v3792)
	mBase = m.M
	v3795 = m.ExcPending
	if v3795 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	v3796 = F_lappend(m, v3785, v3794)
	mBase = m.M
	v3797 = m.ExcPending
	if v3797 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+24)) = v3796
	v3799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3790)+12)))
	if v3799 != 0 {
		goto L689
	} else {
		goto L690
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3787+v3743))) = v3878
	v4049 = v3761 + int32(1)
	v4050 = *(*int32)(unsafe.Add(mBase, uint32(v3745)+4))
	if v4049 < v4050 {
		v3761 = v4049
		goto L680
	} else {
		goto L746
	}
L686:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L1
	} else {
		goto L741
	}
L687:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3974 = m.ExcPending
	if v3974 != 0 {
		goto L1
	} else {
		goto L736
	}
L688:
	;
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+28))
	v3823 = F_lappend_oid(m, v3822, v3821)
	mBase = m.M
	v3824 = m.ExcPending
	if v3824 != 0 {
		goto L1
	} else {
		goto L695
	}
L689:
	;
	v3800 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+64))
	if v3800 != int32(-1) {
		goto L687
	} else {
		goto L692
	}
L690:
	;
	goto L691
L691:
	;
	v3809 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+8))
	v3810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3809)+12)))
	if v3810 == int32(1) {
		goto L686
	} else {
		goto L693
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+328)) = int32(-1)
	v3805 = int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+332)) = v3805
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+64)) = v3761
	v3821 = v3805
	goto L688
L693:
	;
	F_typenameTypeIdAndMod(m, l0, v3809, v37+int32(332), v37+int32(328))
	mBase = m.M
	v3818 = m.ExcPending
	if v3818 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v3819 = *(*int32)(unsafe.Add(mBase, uint32(v37)+332))
	v3821 = v3819
	goto L688
L695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+28)) = v3823
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+32))
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v37)+328))
	v3828 = F_lappend_int(m, v3826, v3827)
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L1
	} else {
		goto L696
	}
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+32)) = v3828
	v3831 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+36))
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v37)+332))
	v3833 = F_get_typcollation(m, v3832)
	mBase = m.M
	v3834 = m.ExcPending
	if v3834 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	v3835 = F_lappend_oid(m, v3831, v3833)
	mBase = m.M
	v3836 = m.ExcPending
	if v3836 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+36)) = v3835
	v3838 = int32(0)
	v3840 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+16))
	if v3840 != 0 {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v3842 = F_transformExpr(m, l0, v3840, int32(5))
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L1
	} else {
		goto L702
	}
L700:
	;
	v3850 = v3838
	goto L701
L701:
	;
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+20))
	if v3851 != 0 {
		goto L705
	} else {
		goto L706
	}
L702:
	;
	v3846 = F_coerce_to_specific_type(m, l0, v3842, int32(25), int32(_a_F_transformFromClauseItem_36))
	mBase = m.M
	v3847 = m.ExcPending
	if v3847 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	F_assign_expr_collations(m, l0, v3846)
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	v3850 = v3846
	goto L701
L705:
	;
	v3853 = F_transformExpr(m, l0, v3851, int32(5))
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		goto L1
	} else {
		goto L708
	}
L706:
	;
	v3862 = v3838
	goto L707
L707:
	;
	v3863 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+40))
	v3864 = F_lappend(m, v3863, v3850)
	mBase = m.M
	v3865 = m.ExcPending
	if v3865 != 0 {
		goto L1
	} else {
		goto L711
	}
L708:
	;
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v37)+332))
	v3856 = *(*int32)(unsafe.Add(mBase, uint32(v37)+328))
	v3858 = F_coerce_to_specific_type_typmod(m, l0, v3853, v3855, v3856, int32(_a_F_transformFromClauseItem_36))
	mBase = m.M
	v3859 = m.ExcPending
	if v3859 != 0 {
		goto L1
	} else {
		goto L709
	}
L709:
	;
	F_assign_expr_collations(m, l0, v3858)
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	v3862 = v3858
	goto L707
L711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+40)) = v3864
	v3867 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+44))
	v3868 = F_lappend(m, v3867, v3862)
	mBase = m.M
	v3869 = m.ExcPending
	if v3869 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+44)) = v3868
	v3871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3790)+13)))
	if v3871 == int32(1) {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3708)+56))
	v3875 = F_bms_add_member(m, v3874, v3761)
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L1
	} else {
		goto L716
	}
L714:
	;
	goto L715
L715:
	;
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+4))
	v3879 = int32(0)
	if v3761 == v3879 {
		goto L685
	} else {
		goto L717
	}
L716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+56)) = v3875
	goto L715
L717:
	;
	v3886 = v3879
	goto L718
L718:
	;
	v3919 = *(*int32)(unsafe.Add(mBase, uint32(v3743+v3886<<(uint(int32(2))%32))))
	v3922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3919))))
	v3925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3878))))
	if base.B2i32(v3922 == int32(0))|base.B2i32(v3922 != v3925) != 0 {
		v3943 = v3922
		v3944 = v3925
		goto L721
	} else {
		goto L722
	}
L719:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		goto L1
	} else {
		goto L731
	}
L720:
	;
	if v3943-v3944 != 0 {
		goto L727
	} else {
		goto L728
	}
L721:
	;
	goto L720
L722:
	;
	v3928 = v3919
	v3929 = v3878
	goto L723
L723:
	;
	v3932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3929)+1)))
	v3933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3928)+1)))
	if v3933 == int32(0) {
		v3943 = v3933
		v3944 = v3932
		goto L721
	} else {
		goto L725
	}
L724:
	;
	v3943 = v3933
	v3944 = v3932
	goto L721
L725:
	;
	v3936 = int32(1)
	if v3933 == v3932 {
		v3928 = v3928 + v3936
		v3929 = v3929 + v3936
		goto L723
	} else {
		goto L726
	}
L726:
	;
	goto L724
L727:
	;
	v3947 = v3886 + int32(1)
	if v3761 != v3947 {
		v3886 = v3947
		goto L718
	} else {
		goto L730
	}
L728:
	;
	goto L729
L729:
	;
	goto L719
L730:
	;
	goto L685
L731:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3955 = m.ExcPending
	if v3955 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v3956
	F_errmsg(m, int32(_a_F_transformFromClauseItem_37), v37+int32(240))
	mBase = m.M
	v3962 = m.ExcPending
	if v3962 != 0 {
		goto L1
	} else {
		goto L733
	}
L733:
	;
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+24))
	F_parser_errposition(m, l0, v3963)
	mBase = m.M
	v3965 = m.ExcPending
	if v3965 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(825), int32(_a_F_transformFromClauseItem_38))
	mBase = m.M
	v3970 = m.ExcPending
	if v3970 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L736:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_39), int32(0))
	mBase = m.M
	v3981 = m.ExcPending
	if v3981 != 0 {
		goto L1
	} else {
		goto L738
	}
L738:
	;
	v3982 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+24))
	F_parser_errposition(m, l0, v3982)
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(763), int32(_a_F_transformFromClauseItem_38))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L741:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v3996 = m.ExcPending
	if v3996 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v3997 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+256)) = v3997
	F_errmsg(m, int32(_a_F_transformFromClauseItem_29), v37+int32(256))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	v4004 = *(*int32)(unsafe.Add(mBase, uint32(v3790)+24))
	F_parser_errposition(m, l0, v4004)
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(776), int32(_a_F_transformFromClauseItem_38))
	mBase = m.M
	v4011 = m.ExcPending
	if v4011 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L746:
	;
	goto L681
L747:
	;
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v4056
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_40), v37)
	mBase = m.M
	v4060 = m.ExcPending
	if v4060 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1627), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L750:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		goto L1
	} else {
		goto L751
	}
L751:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_42), int32(0))
	mBase = m.M
	v4076 = m.ExcPending
	if v4076 != 0 {
		goto L1
	} else {
		goto L752
	}
L752:
	;
	F_errhint(m, int32(_a_F_transformFromClauseItem_43), int32(0))
	mBase = m.M
	v4080 = m.ExcPending
	if v4080 != 0 {
		goto L1
	} else {
		goto L753
	}
L753:
	;
	v4081 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4082 = F_exprLocation(m, v4081)
	mBase = m.M
	F_parser_errposition(m, l0, v4082)
	mBase = m.M
	v4084 = m.ExcPending
	if v4084 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(660), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v4089 = m.ExcPending
	if v4089 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L756:
	;
	F_errhint(m, int32(_a_F_transformFromClauseItem_44), int32(0))
	mBase = m.M
	v4097 = m.ExcPending
	if v4097 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	v4098 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4099 = F_exprLocation(m, v4098)
	mBase = m.M
	F_parser_errposition(m, l0, v4099)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(645), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L1
	} else {
		goto L759
	}
L759:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L760:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4113 = m.ExcPending
	if v4113 != 0 {
		goto L1
	} else {
		goto L761
	}
L761:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_45), int32(0))
	mBase = m.M
	v4117 = m.ExcPending
	if v4117 != 0 {
		goto L1
	} else {
		goto L762
	}
L762:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4119 = F_exprLocation(m, v4118)
	mBase = m.M
	F_parser_errposition(m, l0, v4119)
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L1
	} else {
		goto L763
	}
L763:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(611), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v4126 = m.ExcPending
	if v4126 != 0 {
		goto L1
	} else {
		goto L764
	}
L764:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L765:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4133 = m.ExcPending
	if v4133 != 0 {
		goto L1
	} else {
		goto L766
	}
L766:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_46), int32(0))
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L1
	} else {
		goto L767
	}
L767:
	;
	v4138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v4139 = F_exprLocation(m, v4138)
	mBase = m.M
	F_parser_errposition(m, l0, v4139)
	mBase = m.M
	v4141 = m.ExcPending
	if v4141 != 0 {
		goto L1
	} else {
		goto L768
	}
L768:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(599), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v4146 = m.ExcPending
	if v4146 != 0 {
		goto L1
	} else {
		goto L769
	}
L769:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L770:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4153 = m.ExcPending
	if v4153 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_46), int32(0))
	mBase = m.M
	v4157 = m.ExcPending
	if v4157 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	v4158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v4159 = F_exprLocation(m, v4158)
	mBase = m.M
	F_parser_errposition(m, l0, v4159)
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(571), int32(_a_F_transformFromClauseItem_13))
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L775:
	;
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_47), int32(0))
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(448), int32(_a_F_transformFromClauseItem_48))
	mBase = m.M
	v4179 = m.ExcPending
	if v4179 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L778:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v4185
	v4188 = *(*int32)(unsafe.Add(mBase, uint32(v37)+284))
	if v4188 == int32(0) {
		goto L779
	} else {
		goto L780
	}
L779:
	;
	v4278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4278 != 0 {
		goto L785
	} else {
		goto L786
	}
L780:
	;
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+4))
	if v4191 <= int32(0) {
		goto L779
	} else {
		goto L781
	}
L781:
	;
	v4194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v4202 = int32(0)
	goto L782
L782:
	;
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+12))
	v4236 = *(*int32)(unsafe.Add(mBase, uint32(v4232+v4202<<(uint(int32(2))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4236)+23)) = uint8(base.B2i32(base.Ui32(v4194) < base.Ui32(int32(2))))
	v4238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4236)+22)) = uint8(v4238)
	v4241 = v4202 + v4238
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+4))
	if v4241 < v4242 {
		v4202 = v4241
		goto L782
	} else {
		goto L784
	}
L783:
	;
	goto L779
L784:
	;
	goto L783
L785:
	;
	v4279 = *(*int32)(unsafe.Add(mBase, uint32(v4278)+4))
	v4281 = v4279
	goto L787
L786:
	;
	v4281 = int32(0)
	goto L787
L787:
	;
	v4282 = F_list_concat(m, v4278, v4188)
	mBase = m.M
	v4283 = m.ExcPending
	if v4283 != 0 {
		goto L1
	} else {
		goto L788
	}
L788:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4282
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4290 = F_transformFromClauseItem(m, l0, v4285, v37+int32(288), v37+int32(280))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		goto L1
	} else {
		goto L789
	}
L789:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v4290
	v4293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4294 = int32(0)
	if base.B2i32(v4293 == v4294)|base.B2i32(v4281 <= v4294) != 0 {
		goto L791
	} else {
		goto L792
	}
L790:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4304
	v4306 = *(*int32)(unsafe.Add(mBase, uint32(v37)+280))
	F_checkNameSpaceConflicts(m, v4188, v4306)
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L797
	}
L791:
	;
	v4304 = int32(0)
	goto L793
L792:
	;
	v4301 = *(*int32)(unsafe.Add(mBase, uint32(v4293)+4))
	if v4281 < v4301 {
		goto L794
	} else {
		goto L795
	}
L793:
	;
	goto L790
L794:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4293)+4)) = v4281
	goto L796
L795:
	;
	goto L796
L796:
	;
	v4304 = v4293
	goto L793
L797:
	;
	v4309 = F_list_concat(m, v4188, v4306)
	mBase = m.M
	v4310 = m.ExcPending
	if v4310 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	v4311 = *(*int32)(unsafe.Add(mBase, uint32(v37)+288))
	v4312 = *(*int32)(unsafe.Add(mBase, uint32(v4311)+16))
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v37)+292))
	v4314 = *(*int32)(unsafe.Add(mBase, uint32(v4313)+16))
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4311)))
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4315)+8))
	v4317 = *(*int32)(unsafe.Add(mBase, uint32(v4313)))
	v4318 = *(*int32)(unsafe.Add(mBase, uint32(v4317)+8))
	v4319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v4319 == int32(1) {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	if v4318 == int32(0) {
		v4505 = v5
		goto L802
	} else {
		goto L803
	}
L800:
	;
	goto L801
L801:
	;
	v4563 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v4563 != 0 {
		goto L830
	} else {
		goto L831
	}
L802:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4505
	goto L801
L803:
	;
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+4))
	if v4324 <= int32(0) {
		v4505 = v5
		goto L802
	} else {
		goto L804
	}
L804:
	;
	v4333 = v5
	v4338 = v5
	goto L805
L805:
	;
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+12))
	v4365 = *(*int32)(unsafe.Add(mBase, uint32(v4361+v4333<<(uint(int32(2))%32))))
	v4366 = *(*int32)(unsafe.Add(mBase, uint32(v4365)+4))
	v4367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4366))))
	v4368 = int32(0)
	if base.B2i32(v4367 == v4368)|base.B2i32(v4316 == v4368) != 0 {
		v4467 = v4338
		goto L807
	} else {
		goto L808
	}
L806:
	;
	v4505 = v4467
	goto L802
L807:
	;
	v4491 = v4333 + int32(1)
	v4492 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+4))
	if v4491 < v4492 {
		v4333 = v4491
		v4338 = v4467
		goto L805
	} else {
		goto L829
	}
L808:
	;
	v4373 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v4373 <= int32(0) {
		v4467 = v4338
		goto L807
	} else {
		goto L809
	}
L809:
	;
	v4376 = int32(0)
	if v4376 < v4373 {
		goto L810
	} else {
		goto L811
	}
L810:
	;
	v4379 = v4373
	goto L812
L811:
	;
	v4379 = v4376
	goto L812
L812:
	;
	v4380 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v4386 = int32(0)
	goto L813
L813:
	;
	v4419 = *(*int32)(unsafe.Add(mBase, uint32(v4380+v4386<<(uint(int32(2))%32))))
	v4420 = *(*int32)(unsafe.Add(mBase, uint32(v4419)+4))
	v4423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4366))))
	v4426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4420))))
	if base.B2i32(v4423 == int32(0))|base.B2i32(v4423 != v4426) != 0 {
		v4444 = v4423
		v4445 = v4426
		goto L816
	} else {
		goto L817
	}
L814:
	;
	v4450 = F_makeString(m, v4366)
	mBase = m.M
	v4451 = m.ExcPending
	if v4451 != 0 {
		goto L1
	} else {
		goto L826
	}
L815:
	;
	if v4444-v4445 != 0 {
		goto L822
	} else {
		goto L823
	}
L816:
	;
	goto L815
L817:
	;
	v4429 = v4366
	v4430 = v4420
	goto L818
L818:
	;
	v4433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4430)+1)))
	v4434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4429)+1)))
	if v4434 == int32(0) {
		v4444 = v4434
		v4445 = v4433
		goto L816
	} else {
		goto L820
	}
L819:
	;
	v4444 = v4434
	v4445 = v4433
	goto L816
L820:
	;
	v4437 = int32(1)
	if v4434 == v4433 {
		v4429 = v4429 + v4437
		v4430 = v4430 + v4437
		goto L818
	} else {
		goto L821
	}
L821:
	;
	goto L819
L822:
	;
	v4448 = v4386 + int32(1)
	if v4379 != v4448 {
		v4386 = v4448
		goto L813
	} else {
		goto L825
	}
L823:
	;
	goto L824
L824:
	;
	goto L814
L825:
	;
	v4467 = v4338
	goto L807
L826:
	;
	if v4450 == int32(0) {
		v4467 = v4338
		goto L807
	} else {
		goto L827
	}
L827:
	;
	v4454 = F_lappend(m, v4338, v4450)
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	v4467 = v4454
	goto L807
L829:
	;
	goto L806
L830:
	;
	v4564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4563)+8)) = v4564
	goto L832
L831:
	;
	goto L832
L832:
	;
	v4566 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+268)) = v4566
	*(*int32)(unsafe.Add(mBase, uint32(v37)+272)) = v4566
	*(*int32)(unsafe.Add(mBase, uint32(v37)+276)) = v4566
	*(*int32)(unsafe.Add(mBase, uint32(v37)+264)) = v4566
	if v4318 != 0 {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	v4576 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+4))
	v4577 = v4576
	goto L835
L834:
	;
	v4577 = v4566
	goto L835
L835:
	;
	if v4316 != 0 {
		goto L836
	} else {
		goto L837
	}
L836:
	;
	v4578 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	v4580 = v4578
	goto L838
L837:
	;
	v4580 = int32(0)
	goto L838
L838:
	;
	v4584 = F_palloc0(m, (v4580+v4577)<<(uint(int32(5))%32))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L1
	} else {
		goto L839
	}
L839:
	;
	v4586 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4586 != 0 {
		goto L846
	} else {
		goto L847
	}
L840:
	;
	v5513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5513 != 0 {
		goto L986
	} else {
		goto L987
	}
L841:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+272)) = v5366
	*(*int32)(unsafe.Add(mBase, uint32(v37)+276)) = v5361
	*(*int32)(unsafe.Add(mBase, uint32(v37)+268)) = v5347
	v5375 = int32(0)
	v5381 = v5375
	v5382 = v5375
	goto L962
L842:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5323 = m.ExcPending
	if v5323 != 0 {
		goto L1
	} else {
		goto L957
	}
L843:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5271 = m.ExcPending
	if v5271 != 0 {
		goto L1
	} else {
		goto L953
	}
L844:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5253 = m.ExcPending
	if v5253 != 0 {
		goto L1
	} else {
		goto L949
	}
L845:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5201 = m.ExcPending
	if v5201 != 0 {
		goto L1
	} else {
		goto L945
	}
L846:
	;
	v4587 = *(*int32)(unsafe.Add(mBase, uint32(v4586)+4))
	if v4587 <= int32(0) {
		v5347 = v5
		v5356 = v5
		v5361 = v5
		v5363 = v5
		v5366 = v5
		goto L841
	} else {
		goto L849
	}
L847:
	;
	goto L848
L848:
	;
	v5066 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v5066 != 0 {
		goto L931
	} else {
		goto L932
	}
L849:
	;
	v4599 = v5
	v4608 = v5
	v4613 = v5
	v4615 = v5
	v4618 = v5
	v4619 = v5
	goto L850
L850:
	;
	v4624 = *(*int32)(unsafe.Add(mBase, uint32(v4586)+12))
	v4627 = v4624 + v4619<<(uint(int32(2))%32)
	v4628 = *(*int32)(unsafe.Add(mBase, uint32(v4627)))
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v4628)+4))
	if v4613 == int32(0) {
		goto L852
	} else {
		goto L853
	}
L851:
	;
	v5347 = v5015
	v5356 = v5036
	v5361 = v5060
	v5363 = v5057
	v5366 = v4887
	goto L841
L852:
	;
	if v4318 == int32(0) {
		goto L844
	} else {
		goto L875
	}
L853:
	;
	v4632 = *(*int32)(unsafe.Add(mBase, uint32(v4613)+4))
	if v4632 <= int32(0) {
		goto L852
	} else {
		goto L854
	}
L854:
	;
	v4635 = int32(0)
	if v4635 < v4632 {
		goto L855
	} else {
		goto L856
	}
L855:
	;
	v4638 = v4632
	goto L857
L856:
	;
	v4638 = v4635
	goto L857
L857:
	;
	v4639 = *(*int32)(unsafe.Add(mBase, uint32(v4613)+12))
	v4645 = int32(0)
	goto L858
L858:
	;
	v4678 = *(*int32)(unsafe.Add(mBase, uint32(v4639+v4645<<(uint(int32(2))%32))))
	v4679 = *(*int32)(unsafe.Add(mBase, uint32(v4678)+4))
	v4682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4679))))
	v4685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4629))))
	if base.B2i32(v4682 == int32(0))|base.B2i32(v4682 != v4685) != 0 {
		v4703 = v4682
		v4704 = v4685
		goto L861
	} else {
		goto L862
	}
L859:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4712 = m.ExcPending
	if v4712 != 0 {
		goto L1
	} else {
		goto L871
	}
L860:
	;
	if v4703-v4704 != 0 {
		goto L867
	} else {
		goto L868
	}
L861:
	;
	goto L860
L862:
	;
	v4688 = v4679
	v4689 = v4629
	goto L863
L863:
	;
	v4692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4689)+1)))
	v4693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4688)+1)))
	if v4693 == int32(0) {
		v4703 = v4693
		v4704 = v4692
		goto L861
	} else {
		goto L865
	}
L864:
	;
	v4703 = v4693
	v4704 = v4692
	goto L861
L865:
	;
	v4696 = int32(1)
	if v4693 == v4692 {
		v4688 = v4688 + v4696
		v4689 = v4689 + v4696
		goto L863
	} else {
		goto L866
	}
L866:
	;
	goto L864
L867:
	;
	v4707 = v4645 + int32(1)
	if v4638 != v4707 {
		v4645 = v4707
		goto L858
	} else {
		goto L870
	}
L868:
	;
	goto L869
L869:
	;
	goto L859
L870:
	;
	goto L852
L871:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v4715 = m.ExcPending
	if v4715 != 0 {
		goto L1
	} else {
		goto L872
	}
L872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+144)) = v4629
	F_errmsg(m, int32(_a_F_transformFromClauseItem_49), v37+int32(144))
	mBase = m.M
	v4721 = m.ExcPending
	if v4721 != 0 {
		goto L1
	} else {
		goto L873
	}
L873:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1332), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v4726 = m.ExcPending
	if v4726 != 0 {
		goto L1
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
	v4763 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+4))
	if v4763 <= int32(0) {
		goto L877
	} else {
		goto L878
	}
L876:
	;
	if v4855 < int32(0) {
		goto L844
	} else {
		goto L897
	}
L877:
	;
	v4855 = int32(-1)
	goto L876
L878:
	;
	goto L879
L879:
	;
	v4767 = int32(0)
	if v4767 < v4763 {
		goto L880
	} else {
		goto L881
	}
L880:
	;
	v4770 = v4763
	goto L882
L881:
	;
	v4770 = v4767
	goto L882
L882:
	;
	v4771 = *(*int32)(unsafe.Add(mBase, uint32(v4318)+12))
	v4778 = int32(0)
	v4780 = int32(-1)
	goto L883
L883:
	;
	v4811 = *(*int32)(unsafe.Add(mBase, uint32(v4771+v4778<<(uint(int32(2))%32))))
	v4812 = *(*int32)(unsafe.Add(mBase, uint32(v4811)+4))
	v4815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4812))))
	v4818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4629))))
	if base.B2i32(v4815 == int32(0))|base.B2i32(v4815 != v4818) != 0 {
		v4836 = v4815
		v4837 = v4818
		goto L886
	} else {
		goto L887
	}
L884:
	;
	v4855 = v4845
	goto L876
L885:
	;
	if v4836-v4837 == int32(0) {
		goto L892
	} else {
		goto L893
	}
L886:
	;
	goto L885
L887:
	;
	v4821 = v4812
	v4822 = v4629
	goto L888
L888:
	;
	v4825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4822)+1)))
	v4826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4821)+1)))
	if v4826 == int32(0) {
		v4836 = v4826
		v4837 = v4825
		goto L886
	} else {
		goto L890
	}
L889:
	;
	v4836 = v4826
	v4837 = v4825
	goto L886
L890:
	;
	v4829 = int32(1)
	if v4826 == v4825 {
		v4821 = v4821 + v4829
		v4822 = v4822 + v4829
		goto L888
	} else {
		goto L891
	}
L891:
	;
	goto L889
L892:
	;
	v4841 = int32(0)
	if base.B2i32(v4780 < v4841) == v4841 {
		goto L845
	} else {
		goto L895
	}
L893:
	;
	v4845 = v4780
	goto L894
L894:
	;
	v4847 = v4778 + int32(1)
	if v4847 != v4770 {
		v4778 = v4847
		v4780 = v4845
		goto L883
	} else {
		goto L896
	}
L895:
	;
	v4845 = v4778
	goto L894
L896:
	;
	goto L884
L897:
	;
	v4887 = F_lappend_int(m, v4618, v4855+int32(1))
	mBase = m.M
	v4888 = m.ExcPending
	if v4888 != 0 {
		goto L1
	} else {
		goto L898
	}
L898:
	;
	if v4316 == int32(0) {
		goto L842
	} else {
		goto L899
	}
L899:
	;
	v4891 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+4))
	if v4891 <= int32(0) {
		goto L901
	} else {
		goto L902
	}
L900:
	;
	if v4988 < int32(0) {
		goto L842
	} else {
		goto L921
	}
L901:
	;
	v4988 = int32(-1)
	goto L900
L902:
	;
	goto L903
L903:
	;
	v4895 = int32(0)
	if v4895 < v4891 {
		goto L904
	} else {
		goto L905
	}
L904:
	;
	v4898 = v4891
	goto L906
L905:
	;
	v4898 = v4895
	goto L906
L906:
	;
	v4899 = *(*int32)(unsafe.Add(mBase, uint32(v4316)+12))
	v4906 = int32(0)
	v4913 = int32(-1)
	goto L907
L907:
	;
	v4939 = *(*int32)(unsafe.Add(mBase, uint32(v4899+v4906<<(uint(int32(2))%32))))
	v4940 = *(*int32)(unsafe.Add(mBase, uint32(v4939)+4))
	v4943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4940))))
	v4946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4629))))
	if base.B2i32(v4943 == int32(0))|base.B2i32(v4943 != v4946) != 0 {
		v4964 = v4943
		v4965 = v4946
		goto L910
	} else {
		goto L911
	}
L908:
	;
	v4988 = v4973
	goto L900
L909:
	;
	if v4964-v4965 == int32(0) {
		goto L916
	} else {
		goto L917
	}
L910:
	;
	goto L909
L911:
	;
	v4949 = v4940
	v4950 = v4629
	goto L912
L912:
	;
	v4953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4950)+1)))
	v4954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4949)+1)))
	if v4954 == int32(0) {
		v4964 = v4954
		v4965 = v4953
		goto L910
	} else {
		goto L914
	}
L913:
	;
	v4964 = v4954
	v4965 = v4953
	goto L910
L914:
	;
	v4957 = int32(1)
	if v4954 == v4953 {
		v4949 = v4949 + v4957
		v4950 = v4950 + v4957
		goto L912
	} else {
		goto L915
	}
L915:
	;
	goto L913
L916:
	;
	v4969 = int32(0)
	if base.B2i32(v4913 < v4969) == v4969 {
		goto L843
	} else {
		goto L919
	}
L917:
	;
	v4973 = v4913
	goto L918
L918:
	;
	v4975 = v4906 + int32(1)
	if v4975 != v4898 {
		v4906 = v4975
		v4913 = v4973
		goto L907
	} else {
		goto L920
	}
L919:
	;
	v4973 = v4906
	goto L918
L920:
	;
	goto L908
L921:
	;
	v5015 = F_lappend_int(m, v4599, v4988+int32(1))
	mBase = m.M
	v5016 = m.ExcPending
	if v5016 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	v5019 = v4314 + v4855<<(uint(int32(5))%32)
	v5020 = *(*int32)(unsafe.Add(mBase, uint32(v5019)))
	v5021 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5019)+4)))
	v5022 = *(*int32)(unsafe.Add(mBase, uint32(v5019)+8))
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v5019)+12))
	v5024 = *(*int32)(unsafe.Add(mBase, uint32(v5019)+16))
	v5026 = F_makeVar(m, v5020, v5021, v5022, v5023, v5024, int32(0))
	mBase = m.M
	v5027 = m.ExcPending
	if v5027 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v5028 = *(*int32)(unsafe.Add(mBase, uint32(v5019)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5026)+32)) = v5028
	v5030 = *(*int32)(unsafe.Add(mBase, uint32(v5019)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v5026)+36)) = v5030
	v5032 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5019)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5026)+40)) = uint16(v5032)
	F_markNullableIfNeeded(m, l0, v5026)
	mBase = m.M
	v5035 = m.ExcPending
	if v5035 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v5036 = F_lappend(m, v4608, v5026)
	mBase = m.M
	v5037 = m.ExcPending
	if v5037 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	v5040 = v4312 + v4988<<(uint(int32(5))%32)
	v5041 = *(*int32)(unsafe.Add(mBase, uint32(v5040)))
	v5042 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5040)+4)))
	v5043 = *(*int32)(unsafe.Add(mBase, uint32(v5040)+8))
	v5044 = *(*int32)(unsafe.Add(mBase, uint32(v5040)+12))
	v5045 = *(*int32)(unsafe.Add(mBase, uint32(v5040)+16))
	v5047 = F_makeVar(m, v5041, v5042, v5043, v5044, v5045, int32(0))
	mBase = m.M
	v5048 = m.ExcPending
	if v5048 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(v5040)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5047)+32)) = v5049
	v5051 = *(*int32)(unsafe.Add(mBase, uint32(v5040)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v5047)+36)) = v5051
	v5053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5040)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5047)+40)) = uint16(v5053)
	F_markNullableIfNeeded(m, l0, v5047)
	mBase = m.M
	v5056 = m.ExcPending
	if v5056 != 0 {
		goto L1
	} else {
		goto L927
	}
L927:
	;
	v5057 = F_lappend(m, v4615, v5047)
	mBase = m.M
	v5058 = m.ExcPending
	if v5058 != 0 {
		goto L1
	} else {
		goto L928
	}
L928:
	;
	v5059 = *(*int32)(unsafe.Add(mBase, uint32(v4627)))
	v5060 = F_lappend(m, v4613, v5059)
	mBase = m.M
	v5061 = m.ExcPending
	if v5061 != 0 {
		goto L1
	} else {
		goto L929
	}
L929:
	;
	v5063 = v4619 + int32(1)
	v5064 = *(*int32)(unsafe.Add(mBase, uint32(v4586)+4))
	if v5063 < v5064 {
		v4599 = v5015
		v4608 = v5036
		v4613 = v5060
		v4615 = v5057
		v4618 = v4887
		v4619 = v5063
		goto L850
	} else {
		goto L930
	}
L930:
	;
	goto L851
L931:
	;
	if v4309 == int32(0) {
		goto L934
	} else {
		goto L935
	}
L932:
	;
	goto L933
L933:
	;
	v5488 = v5
	v5507 = v5
	goto L840
L934:
	;
	v5151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v4309
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v5154 != 0 {
		goto L940
	} else {
		goto L941
	}
L935:
	;
	v5069 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+4))
	if v5069 <= int32(0) {
		goto L934
	} else {
		goto L936
	}
L936:
	;
	v5076 = v4566
	goto L937
L937:
	;
	v5106 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+12))
	v5110 = *(*int32)(unsafe.Add(mBase, uint32(v5106+v5076<<(uint(int32(2))%32))))
	v5111 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v5110)+22)) = uint16(v5111)
	v5114 = v5076 + int32(1)
	v5115 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+4))
	if v5114 < v5115 {
		v5076 = v5114
		goto L937
	} else {
		goto L939
	}
L938:
	;
	goto L934
L939:
	;
	goto L938
L940:
	;
	v5156 = F_transformExpr(m, l0, v5154, int32(2))
	mBase = m.M
	v5157 = m.ExcPending
	if v5157 != 0 {
		goto L1
	} else {
		goto L943
	}
L941:
	;
	v5161 = int32(0)
	goto L942
L942:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5151
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v5161
	goto L933
L943:
	;
	v5159 = F_coerce_to_boolean(m, l0, v5156, int32(_a_F_transformFromClauseItem_50))
	mBase = m.M
	v5160 = m.ExcPending
	if v5160 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	v5161 = v5159
	goto L942
L945:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v5204 = m.ExcPending
	if v5204 != 0 {
		goto L1
	} else {
		goto L946
	}
L946:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+128)) = v4629
	F_errmsg(m, int32(_a_F_transformFromClauseItem_51), v37+int32(128))
	mBase = m.M
	v5210 = m.ExcPending
	if v5210 != 0 {
		goto L1
	} else {
		goto L947
	}
L947:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1347), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v5215 = m.ExcPending
	if v5215 != 0 {
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v5256 = m.ExcPending
	if v5256 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+80)) = v4629
	F_errmsg(m, int32(_a_F_transformFromClauseItem_52), v37+int32(80))
	mBase = m.M
	v5262 = m.ExcPending
	if v5262 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1356), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v5267 = m.ExcPending
	if v5267 != 0 {
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
	F_errcode(m, int32(33583236))
	mBase = m.M
	v5274 = m.ExcPending
	if v5274 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+112)) = v4629
	F_errmsg(m, int32(_a_F_transformFromClauseItem_53), v37+int32(112))
	mBase = m.M
	v5280 = m.ExcPending
	if v5280 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1371), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v5285 = m.ExcPending
	if v5285 != 0 {
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v5326 = m.ExcPending
	if v5326 != 0 {
		goto L1
	} else {
		goto L958
	}
L958:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = v4629
	F_errmsg(m, int32(_a_F_transformFromClauseItem_54), v37+int32(96))
	mBase = m.M
	v5332 = m.ExcPending
	if v5332 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1380), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v5337 = m.ExcPending
	if v5337 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L961:
	;
	v5473 = F_transformExpr(m, l0, v5471, int32(3))
	mBase = m.M
	v5474 = m.ExcPending
	if v5474 != 0 {
		goto L1
	} else {
		goto L984
	}
L962:
	;
	v5411 = int32(0)
	if v5356 == v5411 {
		v5422 = v5411
		goto L964
	} else {
		goto L965
	}
L963:
	;
	v5467 = *(*int32)(unsafe.Add(mBase, uint32(v5382)+12))
	v5468 = *(*int32)(unsafe.Add(mBase, uint32(v5467)))
	v5471 = v5468
	goto L961
L964:
	;
	if v5363 == int32(0) {
		v5439 = v5411
		goto L969
	} else {
		goto L970
	}
L965:
	;
	v5416 = *(*int32)(unsafe.Add(mBase, uint32(v5356)+4))
	if v5416 <= v5381 {
		v5422 = int32(0)
		goto L964
	} else {
		goto L966
	}
L966:
	;
	v5418 = *(*int32)(unsafe.Add(mBase, uint32(v5356)+12))
	v5422 = v5418 + v5381<<(uint(int32(2))%32)
	goto L964
L967:
	;
	goto L963
L968:
	;
	v5448 = *(*int32)(unsafe.Add(mBase, uint32(v5432+v5381<<(uint(int32(2))%32))))
	v5449 = *(*int32)(unsafe.Add(mBase, uint32(v5422)))
	F_markVarForSelectPriv(m, l0, v5449)
	mBase = m.M
	v5451 = m.ExcPending
	if v5451 != 0 {
		goto L1
	} else {
		goto L978
	}
L969:
	;
	v5443 = F_makeBoolExpr(m, int32(0), v5439, int32(-1))
	mBase = m.M
	v5444 = m.ExcPending
	if v5444 != 0 {
		goto L1
	} else {
		goto L977
	}
L970:
	;
	v5425 = int32(0)
	v5427 = *(*int32)(unsafe.Add(mBase, uint32(v5363)+4))
	if base.B2i32(v5422 == v5425)|base.B2i32(v5427 <= v5381) == v5425 {
		goto L971
	} else {
		goto L972
	}
L971:
	;
	v5432 = *(*int32)(unsafe.Add(mBase, uint32(v5363)+12))
	if v5432 != 0 {
		goto L968
	} else {
		goto L974
	}
L972:
	;
	goto L973
L973:
	;
	if v5382 == int32(0) {
		v5439 = v5411
		goto L969
	} else {
		goto L975
	}
L974:
	;
	goto L973
L975:
	;
	v5436 = *(*int32)(unsafe.Add(mBase, uint32(v5382)+4))
	if v5436 == int32(1) {
		goto L967
	} else {
		goto L976
	}
L976:
	;
	v5439 = v5382
	goto L969
L977:
	;
	v5471 = v5443
	goto L961
L978:
	;
	F_markVarForSelectPriv(m, l0, v5448)
	mBase = m.M
	v5453 = m.ExcPending
	if v5453 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	v5458 = F_copyObjectImpl(m, v5449)
	mBase = m.M
	v5459 = m.ExcPending
	if v5459 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	v5460 = F_copyObjectImpl(m, v5448)
	mBase = m.M
	v5461 = m.ExcPending
	if v5461 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	v5463 = F_makeSimpleA_Expr(m, int32(0), int32(_a_F_transformFromClauseItem_55), v5458, v5460, int32(-1))
	mBase = m.M
	v5464 = m.ExcPending
	if v5464 != 0 {
		goto L1
	} else {
		goto L982
	}
L982:
	;
	v5465 = F_lappend(m, v5382, v5463)
	mBase = m.M
	v5466 = m.ExcPending
	if v5466 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	v5381 = v5381 + int32(1)
	v5382 = v5465
	goto L962
L984:
	;
	v5476 = F_coerce_to_boolean(m, l0, v5473, int32(_a_F_transformFromClauseItem_56))
	mBase = m.M
	v5477 = m.ExcPending
	if v5477 != 0 {
		goto L1
	} else {
		goto L985
	}
L985:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v5476
	v5488 = v5347
	v5507 = v5366
	goto L840
L986:
	;
	v5514 = *(*int32)(unsafe.Add(mBase, uint32(v5513)+4))
	v5518 = v5514 + int32(1)
	goto L988
L987:
	;
	v5518 = int32(1)
	goto L988
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v5518
	v5520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v5520 {
	case 0:
		goto L989
	case 1:
		goto L990
	case 2:
		goto L993
	case 3:
		goto L992
	default:
		goto L991
	}
L989:
	;
	v5550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5550 == int32(0) {
		goto L1002
	} else {
		goto L1003
	}
L990:
	;
	v5547 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_markRelsAsNulledBy(m, l0, v5547, v5518)
	mBase = m.M
	v5549 = m.ExcPending
	if v5549 != 0 {
		goto L1
	} else {
		goto L1000
	}
L991:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5534 = m.ExcPending
	if v5534 != 0 {
		goto L1
	} else {
		goto L997
	}
L992:
	;
	v5528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v5528, v5518)
	mBase = m.M
	v5530 = m.ExcPending
	if v5530 != 0 {
		goto L1
	} else {
		goto L996
	}
L993:
	;
	v5521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_markRelsAsNulledBy(m, l0, v5521, v5518)
	mBase = m.M
	v5523 = m.ExcPending
	if v5523 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	v5524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v5525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_markRelsAsNulledBy(m, l0, v5524, v5525)
	mBase = m.M
	v5527 = m.ExcPending
	if v5527 != 0 {
		goto L1
	} else {
		goto L995
	}
L995:
	;
	goto L989
L996:
	;
	goto L989
L997:
	;
	v5535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v5535
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_57), v37+int32(32))
	mBase = m.M
	v5541 = m.ExcPending
	if v5541 != 0 {
		goto L1
	} else {
		goto L998
	}
L998:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1449), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v5546 = m.ExcPending
	if v5546 != 0 {
		goto L1
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
	goto L989
L1001:
	;
	v5870 = v37 + int32(276)
	v5872 = v37 + int32(264)
	v5878 = F_extractRemainingColumns(m, l0, v4314, v4318, v37+int32(272), v5870, v5872, v4584+v5838<<(uint(int32(5))%32))
	mBase = m.M
	v5879 = m.ExcPending
	if v5879 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1002:
	;
	v5838 = int32(0)
	goto L1001
L1003:
	;
	goto L1004
L1004:
	;
	v5554 = int32(0)
	v5561 = v5554
	v5581 = v5554
	goto L1005
L1005:
	;
	v5590 = int32(0)
	if v5507 == v5590 {
		v5600 = v5590
		goto L1007
	} else {
		goto L1008
	}
L1006:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+264)) = v5831
	v5838 = v5829
	goto L1001
L1007:
	;
	if v5488 == int32(0) {
		goto L1011
	} else {
		goto L1012
	}
L1008:
	;
	v5594 = *(*int32)(unsafe.Add(mBase, uint32(v5507)+4))
	if v5594 <= v5561 {
		v5600 = int32(0)
		goto L1007
	} else {
		goto L1009
	}
L1009:
	;
	v5596 = *(*int32)(unsafe.Add(mBase, uint32(v5507)+12))
	v5600 = v5596 + v5561<<(uint(int32(2))%32)
	goto L1007
L1010:
	;
	goto L1006
L1011:
	;
	v5603 = int32(0)
	v5829 = v5603
	v5831 = v5603
	goto L1010
L1012:
	;
	goto L1013
L1013:
	;
	v5607 = *(*int32)(unsafe.Add(mBase, uint32(v5488)+4))
	if base.B2i32(v5600 == int32(0))|base.B2i32(v5607 <= v5561) != 0 {
		v5829 = v5561
		v5831 = v5581
		goto L1010
	} else {
		goto L1014
	}
L1014:
	;
	v5610 = *(*int32)(unsafe.Add(mBase, uint32(v5488)+12))
	if v5610 == int32(0) {
		v5829 = v5561
		v5831 = v5581
		goto L1010
	} else {
		goto L1015
	}
L1015:
	;
	v5616 = *(*int32)(unsafe.Add(mBase, uint32(v5610+v5561<<(uint(int32(2))%32))))
	v5617 = *(*int32)(unsafe.Add(mBase, uint32(v5600)))
	v5620 = v4314 + v5617<<(uint(int32(5))%32)
	v5622 = v5620 - int32(32)
	v5623 = *(*int32)(unsafe.Add(mBase, uint32(v5622)))
	v5626 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5620-int32(28)))))
	v5629 = *(*int32)(unsafe.Add(mBase, uint32(v5620-int32(24))))
	v5632 = *(*int32)(unsafe.Add(mBase, uint32(v5620-int32(20))))
	v5635 = *(*int32)(unsafe.Add(mBase, uint32(v5620-int32(16))))
	v5637 = F_makeVar(m, v5623, v5626, v5629, v5632, v5635, int32(0))
	mBase = m.M
	v5638 = m.ExcPending
	if v5638 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	v5641 = *(*int32)(unsafe.Add(mBase, uint32(v5620-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v5637)+32)) = v5641
	v5645 = *(*int32)(unsafe.Add(mBase, uint32(v5620-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v5637)+36)) = v5645
	v5649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5620-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5637)+40)) = uint16(v5649)
	F_markNullableIfNeeded(m, l0, v5637)
	mBase = m.M
	v5652 = m.ExcPending
	if v5652 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	v5655 = v4312 + v5616<<(uint(int32(5))%32)
	v5657 = v5655 - int32(32)
	v5658 = *(*int32)(unsafe.Add(mBase, uint32(v5657)))
	v5661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5655-int32(28)))))
	v5664 = *(*int32)(unsafe.Add(mBase, uint32(v5655-int32(24))))
	v5667 = *(*int32)(unsafe.Add(mBase, uint32(v5655-int32(20))))
	v5670 = *(*int32)(unsafe.Add(mBase, uint32(v5655-int32(16))))
	v5672 = F_makeVar(m, v5658, v5661, v5664, v5667, v5670, int32(0))
	mBase = m.M
	v5673 = m.ExcPending
	if v5673 != 0 {
		goto L1
	} else {
		goto L1018
	}
L1018:
	;
	v5676 = *(*int32)(unsafe.Add(mBase, uint32(v5655-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v5672)+32)) = v5676
	v5680 = *(*int32)(unsafe.Add(mBase, uint32(v5655-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v5672)+36)) = v5680
	v5684 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5655-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5672)+40)) = uint16(v5684)
	F_markNullableIfNeeded(m, l0, v5672)
	mBase = m.M
	v5687 = m.ExcPending
	if v5687 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1019:
	;
	v5688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+76)) = v5637
	*(*int32)(unsafe.Add(mBase, uint32(v37)+72)) = v5672
	*(*int32)(unsafe.Add(mBase, uint32(v37)+332)) = v5637
	*(*int32)(unsafe.Add(mBase, uint32(v37)+328)) = v5672
	v5697 = F_list_make2_impl(m, v37+int32(76), v37+int32(72))
	mBase = m.M
	v5698 = m.ExcPending
	if v5698 != 0 {
		goto L1
	} else {
		goto L1020
	}
L1020:
	;
	v5701 = F_select_common_type(m, l0, v5697, int32(_a_F_transformFromClauseItem_56), int32(0))
	mBase = m.M
	v5702 = m.ExcPending
	if v5702 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+320)) = v5672
	*(*int32)(unsafe.Add(mBase, uint32(v37)+324)) = v5637
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v5637
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v5672
	v5711 = F_list_make2_impl(m, v37+int32(68), v37-int32(-64))
	mBase = m.M
	v5712 = m.ExcPending
	if v5712 != 0 {
		goto L1
	} else {
		goto L1022
	}
L1022:
	;
	v5713 = F_select_common_typmod(m, v5711, v5701)
	mBase = m.M
	v5714 = m.ExcPending
	if v5714 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	v5715 = *(*int32)(unsafe.Add(mBase, uint32(v5637)+12))
	if v5715 != v5701 {
		goto L1025
	} else {
		goto L1026
	}
L1024:
	;
	v5729 = *(*int32)(unsafe.Add(mBase, uint32(v5672)+12))
	if v5701 != v5729 {
		goto L1032
	} else {
		goto L1033
	}
L1025:
	;
	v5720 = F_coerce_type(m, l0, v5637, v5715, v5701, v5713, int32(0), int32(2), int32(-1))
	mBase = m.M
	v5721 = m.ExcPending
	if v5721 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1026:
	;
	goto L1027
L1027:
	;
	v5722 = *(*int32)(unsafe.Add(mBase, uint32(v5637)+16))
	if v5722 == v5713 {
		v5728 = v5637
		goto L1024
	} else {
		goto L1029
	}
L1028:
	;
	v5728 = v5720
	goto L1024
L1029:
	;
	v5726 = F_makeRelabelType(m, v5637, v5701, v5713, int32(0), int32(2))
	mBase = m.M
	v5727 = m.ExcPending
	if v5727 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	v5728 = v5726
	goto L1024
L1031:
	;
	switch v5688 {
	case 0:
		goto L1042
	case 1:
		v5786 = v5728
		goto L1038
	case 2:
		goto L1041
	case 3:
		goto L1039
	default:
		goto L1040
	}
L1032:
	;
	v5734 = F_coerce_type(m, l0, v5672, v5729, v5701, v5713, int32(0), int32(2), int32(-1))
	mBase = m.M
	v5735 = m.ExcPending
	if v5735 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1033:
	;
	goto L1034
L1034:
	;
	v5736 = *(*int32)(unsafe.Add(mBase, uint32(v5672)+16))
	if v5736 == v5713 {
		v5742 = v5672
		goto L1031
	} else {
		goto L1036
	}
L1035:
	;
	v5742 = v5734
	goto L1031
L1036:
	;
	v5740 = F_makeRelabelType(m, v5672, v5701, v5713, int32(0), int32(2))
	mBase = m.M
	v5741 = m.ExcPending
	if v5741 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	v5742 = v5740
	goto L1031
L1038:
	;
	F_assign_expr_collations(m, l0, v5786)
	mBase = m.M
	v5788 = m.ExcPending
	if v5788 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1039:
	;
	v5786 = v5742
	goto L1038
L1040:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1041:
	;
	v5751 = F_palloc0(m, int32(20))
	mBase = m.M
	v5752 = m.ExcPending
	if v5752 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1042:
	;
	v5743 = *(*int32)(unsafe.Add(mBase, uint32(v5728)))
	if v5743 == int32(6) {
		v5786 = v5728
		goto L1038
	} else {
		goto L1043
	}
L1043:
	;
	v5746 = *(*int32)(unsafe.Add(mBase, uint32(v5742)))
	if v5746 == int32(6) {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	v5749 = v5742
	goto L1046
L1045:
	;
	v5749 = v5728
	goto L1046
L1046:
	;
	v5786 = v5749
	goto L1038
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5751)+4)) = v5701
	*(*int32)(unsafe.Add(mBase, uint32(v5751))) = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+312)) = v5742
	*(*int32)(unsafe.Add(mBase, uint32(v37)+316)) = v5728
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v5728
	*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v5742
	v5764 = F_list_make2_impl(m, v37+int32(60), v37+int32(56))
	mBase = m.M
	v5765 = m.ExcPending
	if v5765 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5751)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5751)+12)) = v5764
	v5786 = v5751
	goto L1038
L1049:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v5688
	F_errmsg_internal(m, int32(_a_F_transformFromClauseItem_57), v37+int32(48))
	mBase = m.M
	v5778 = m.ExcPending
	if v5778 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1050:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1756), int32(_a_F_transformFromClauseItem_58))
	mBase = m.M
	v5783 = m.ExcPending
	if v5783 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1051:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1052:
	;
	v5790 = v5561 + int32(1)
	v5793 = v4584 + v5561<<(uint(int32(5))%32)
	v5794 = F_lappend(m, v5581, v5786)
	mBase = m.M
	v5795 = m.ExcPending
	if v5795 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1053:
	;
	if v5786 == v5637 {
		goto L1054
	} else {
		goto L1055
	}
L1054:
	;
	v5797 = *(*int64)(unsafe.Add(mBase, uint32(v5622)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+24)) = v5797
	v5799 = *(*int64)(unsafe.Add(mBase, uint32(v5622)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+16)) = v5799
	v5801 = *(*int64)(unsafe.Add(mBase, uint32(v5622)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+8)) = v5801
	v5803 = *(*int64)(unsafe.Add(mBase, uint32(v5622)))
	*(*int64)(unsafe.Add(mBase, uint32(v5793))) = v5803
	v5561 = v5790
	v5581 = v5794
	goto L1005
L1055:
	;
	goto L1056
L1056:
	;
	if v5786 == v5672 {
		goto L1057
	} else {
		goto L1058
	}
L1057:
	;
	v5806 = *(*int64)(unsafe.Add(mBase, uint32(v5657)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+24)) = v5806
	v5808 = *(*int64)(unsafe.Add(mBase, uint32(v5657)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+16)) = v5808
	v5810 = *(*int64)(unsafe.Add(mBase, uint32(v5657)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5793)+8)) = v5810
	v5812 = *(*int64)(unsafe.Add(mBase, uint32(v5657)))
	*(*int64)(unsafe.Add(mBase, uint32(v5793))) = v5812
	v5561 = v5790
	v5581 = v5794
	goto L1005
L1058:
	;
	goto L1059
L1059:
	;
	v5814 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v5793)+4)) = uint16(v5790)
	*(*int32)(unsafe.Add(mBase, uint32(v5793))) = v5814
	v5817 = F_exprType(m, v5786)
	mBase = m.M
	v5818 = m.ExcPending
	if v5818 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5793)+8)) = v5817
	v5820 = F_exprTypmod(m, v5786)
	mBase = m.M
	v5821 = m.ExcPending
	if v5821 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5793)+12)) = v5820
	v5823 = F_exprCollation(m, v5786)
	mBase = m.M
	v5824 = m.ExcPending
	if v5824 != 0 {
		goto L1
	} else {
		goto L1062
	}
L1062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5793)+16)) = v5823
	v5826 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v5793)+28)) = uint16(v5790)
	*(*int32)(unsafe.Add(mBase, uint32(v5793)+24)) = v5826
	v5561 = v5790
	v5581 = v5794
	goto L1005
L1063:
	;
	v5880 = v5878 + v5838
	v5884 = F_extractRemainingColumns(m, l0, v4312, v4316, v37+int32(268), v5870, v5872, v4584+v5880<<(uint(int32(5))%32))
	mBase = m.M
	v5885 = m.ExcPending
	if v5885 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1064:
	;
	v5886 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v5886 == int32(0) {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	v6090 = *(*int32)(unsafe.Add(mBase, uint32(v37)+276))
	v6091 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6092 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v6092 != 0 {
		goto L1078
	} else {
		goto L1079
	}
L1066:
	;
	v5889 = v5880 + v5884
	if v5889 <= int32(0) {
		goto L1065
	} else {
		goto L1067
	}
L1067:
	;
	v5892 = int32(3)
	v5893 = v5889 & v5892
	v5894 = int32(0)
	if base.Ui32(v5892) <= base.Ui32(v5878+v5884+v5838-int32(1)) {
		goto L1068
	} else {
		goto L1069
	}
L1068:
	;
	v5909 = v5894
	v5912 = int32(0)
	goto L1071
L1069:
	;
	v5980 = v5894
	goto L1070
L1070:
	;
	v6014 = v5980
	v6020 = v5894
	goto L1075
L1071:
	;
	v5939 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5940 = int32(5)
	v5942 = v4584 + v5909<<(uint(v5940)%32)
	v5944 = v5909 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5942)+28)) = uint16(v5944)
	*(*int32)(unsafe.Add(mBase, uint32(v5942)+24)) = v5939
	v5947 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5950 = v4584 + v5944<<(uint(v5940)%32)
	v5952 = v5909 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v5950)+28)) = uint16(v5952)
	*(*int32)(unsafe.Add(mBase, uint32(v5950)+24)) = v5947
	v5955 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5958 = v4584 + v5952<<(uint(v5940)%32)
	v5960 = v5909 | int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v5958)+28)) = uint16(v5960)
	*(*int32)(unsafe.Add(mBase, uint32(v5958)+24)) = v5955
	v5963 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5966 = v4584 + v5960<<(uint(v5940)%32)
	v5967 = int32(4)
	v5968 = v5909 + v5967
	*(*uint16)(unsafe.Add(mBase, uint32(v5966)+28)) = uint16(v5968)
	*(*int32)(unsafe.Add(mBase, uint32(v5966)+24)) = v5963
	v5972 = v5912 + v5967
	if v5972 != v5889&int32(2147483644) {
		v5909 = v5968
		v5912 = v5972
		goto L1071
	} else {
		goto L1073
	}
L1072:
	;
	if v5893 == int32(0) {
		goto L1065
	} else {
		goto L1074
	}
L1073:
	;
	goto L1072
L1074:
	;
	v5980 = v5968
	goto L1070
L1075:
	;
	v6044 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v6047 = v4584 + v6014<<(uint(int32(5))%32)
	v6048 = int32(1)
	v6049 = v6014 + v6048
	*(*uint16)(unsafe.Add(mBase, uint32(v6047)+28)) = uint16(v6049)
	*(*int32)(unsafe.Add(mBase, uint32(v6047)+24)) = v6044
	v6053 = v6020 + v6048
	if v6053 != v5893 {
		v6014 = v6049
		v6020 = v6053
		goto L1075
	} else {
		goto L1077
	}
L1076:
	;
	goto L1065
L1077:
	;
	goto L1076
L1078:
	;
	v6093 = *(*int32)(unsafe.Add(mBase, uint32(v6092)+4))
	v6095 = v6093
	goto L1080
L1079:
	;
	v6095 = int32(0)
	goto L1080
L1080:
	;
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(v37)+264))
	v6097 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
	v6098 = *(*int32)(unsafe.Add(mBase, uint32(v37)+268))
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v6100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v6102 = F_addRangeTableEntryForJoin(m, l0, v6090, v4584, v6091, v6095, v6096, v6097, v6098, v6099, v6100, int32(1))
	mBase = m.M
	v6103 = m.ExcPending
	if v6103 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	v6104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6104 != 0 {
		goto L1082
	} else {
		goto L1083
	}
L1082:
	;
	v6105 = *(*int32)(unsafe.Add(mBase, uint32(v6104)+4))
	v6108 = v6105 + int32(1)
	goto L1084
L1083:
	;
	v6108 = int32(1)
	goto L1084
L1084:
	;
	v6109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v6108 < v6109 {
		goto L1085
	} else {
		goto L1086
	}
L1085:
	;
	v6115 = v6104
	v6121 = v6108
	goto L1088
L1086:
	;
	v6157 = v6104
	goto L1087
L1087:
	;
	v6187 = F_lappend(m, v6157, l1)
	mBase = m.M
	v6188 = m.ExcPending
	if v6188 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1088:
	;
	v6146 = F_lappend(m, v6115, int32(0))
	mBase = m.M
	v6147 = m.ExcPending
	if v6147 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1089:
	;
	v6157 = v6146
	goto L1087
L1090:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6146
	v6150 = v6121 + int32(1)
	v6151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v6150 < v6151 {
		v6115 = v6146
		v6121 = v6150
		goto L1088
	} else {
		goto L1091
	}
L1091:
	;
	goto L1089
L1092:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6187
	v6190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v6190 != 0 {
		goto L1093
	} else {
		goto L1094
	}
L1093:
	;
	v6192 = F_palloc(m, int32(28))
	mBase = m.M
	v6193 = m.ExcPending
	if v6193 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1094:
	;
	v6218 = v4309
	goto L1095
L1095:
	;
	v6219 = int32(0)
	v6221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v6221 != 0 {
		v6323 = v6219
		v6347 = int32(1)
		goto L1100
	} else {
		goto L1101
	}
L1096:
	;
	v6194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6192))) = v6194
	v6196 = *(*int32)(unsafe.Add(mBase, uint32(v6102)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6192)+4)) = v6196
	v6198 = *(*int32)(unsafe.Add(mBase, uint32(v6102)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v6192)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v6192)+16)) = v4584
	*(*int32)(unsafe.Add(mBase, uint32(v6192)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6192)+8)) = v6198
	*(*int32)(unsafe.Add(mBase, uint32(v37)+44)) = v6192
	*(*int32)(unsafe.Add(mBase, uint32(v37)+260)) = v6192
	v6210 = F_list_make1_impl(m, int32(1), v37+int32(44))
	mBase = m.M
	v6211 = m.ExcPending
	if v6211 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	F_checkNameSpaceConflicts(m, v6210, v4309)
	mBase = m.M
	v6213 = m.ExcPending
	if v6213 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	v6214 = F_lappend(m, v4309, v6192)
	mBase = m.M
	v6215 = m.ExcPending
	if v6215 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1099:
	;
	v6218 = v6214
	goto L1095
L1100:
	;
	v6348 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6102)+23)) = uint8(v6348)
	*(*uint16)(unsafe.Add(mBase, uint32(v6102)+21)) = uint16(v6348)
	*(*uint8)(unsafe.Add(mBase, uint32(v6102)+20)) = uint8(v6347)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v6102
	v6354 = F_lappend(m, v6323, v6102)
	mBase = m.M
	v6355 = m.ExcPending
	if v6355 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1101:
	;
	v6222 = int32(0)
	if v6218 == v6222 {
		v6323 = v6219
		v6347 = v6222
		goto L1100
	} else {
		goto L1102
	}
L1102:
	;
	v6225 = int32(0)
	v6226 = *(*int32)(unsafe.Add(mBase, uint32(v6218)+4))
	if v6225 < v6226 {
		goto L1103
	} else {
		goto L1104
	}
L1103:
	;
	v6233 = v6225
	goto L1106
L1104:
	;
	v6312 = int32(0)
	goto L1105
L1105:
	;
	v6323 = v6218
	v6347 = v6312
	goto L1100
L1106:
	;
	v6263 = *(*int32)(unsafe.Add(mBase, uint32(v6218)+12))
	v6267 = *(*int32)(unsafe.Add(mBase, uint32(v6263+v6233<<(uint(int32(2))%32))))
	v6268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6267)+21)) = uint8(v6268)
	v6271 = v6233 + int32(1)
	v6272 = *(*int32)(unsafe.Add(mBase, uint32(v6218)+4))
	if v6271 < v6272 {
		v6233 = v6271
		goto L1106
	} else {
		goto L1108
	}
L1107:
	;
	v6274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v6312 = base.B2i32(v6274 != int32(0))
	goto L1105
L1108:
	;
	goto L1107
L1109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v6354
	v7056 = l1
	goto L3
L1110:
	;
	v6360 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v6361 = *(*int32)(unsafe.Add(mBase, uint32(v6360)+4))
	v6362 = *(*int32)(unsafe.Add(mBase, uint32(v6361)+12))
	if v6362 != 0 {
		goto L1112
	} else {
		goto L1113
	}
L1111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6593 = m.ExcPending
	if v6593 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6573 = m.ExcPending
	if v6573 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1113:
	;
	v6363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6361)+21)))
	v6365 = v6363 - int32(109)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v6365))|base.B2i32(int32(1)<<(uint(v6365)%32)&int32(41) == int32(0)) != 0 {
		goto L1112
	} else {
		goto L1114
	}
L1114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+332)) = int32(2281)
	v6377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6378 = int32(1)
	v6382 = F_LookupFuncName(m, v6377, v6378, v37+int32(332), v6378)
	mBase = m.M
	v6383 = m.ExcPending
	if v6383 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	if v6382 != 0 {
		goto L1116
	} else {
		goto L1117
	}
L1116:
	;
	v6384 = F_get_func_rettype(m, v6382)
	mBase = m.M
	v6385 = m.ExcPending
	if v6385 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1117:
	;
	goto L1118
L1118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6548 = m.ExcPending
	if v6548 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1119:
	;
	if v6384 == int32(3310) {
		goto L1120
	} else {
		goto L1121
	}
L1120:
	;
	v6388 = F_GetTsmRoutine(m, v6382)
	mBase = m.M
	v6389 = m.ExcPending
	if v6389 != 0 {
		goto L1
	} else {
		goto L1123
	}
L1121:
	;
	goto L1122
L1122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6522 = m.ExcPending
	if v6522 != 0 {
		goto L1
	} else {
		goto L1162
	}
L1123:
	;
	v6391 = F_palloc0(m, int32(16))
	mBase = m.M
	v6392 = m.ExcPending
	if v6392 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6391)+4)) = v6382
	*(*int32)(unsafe.Add(mBase, uint32(v6391))) = int32(104)
	v6396 = int32(0)
	v6397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v6397 != 0 {
		goto L1125
	} else {
		goto L1126
	}
L1125:
	;
	v6398 = *(*int32)(unsafe.Add(mBase, uint32(v6397)+4))
	v6399 = v6398
	goto L1127
L1126:
	;
	v6399 = v5
	goto L1127
L1127:
	;
	v6400 = *(*int32)(unsafe.Add(mBase, uint32(v6388)+4))
	if v6400 != 0 {
		goto L1128
	} else {
		goto L1129
	}
L1128:
	;
	v6401 = *(*int32)(unsafe.Add(mBase, uint32(v6400)+4))
	v6403 = v6401
	goto L1130
L1129:
	;
	v6403 = int32(0)
	goto L1130
L1130:
	;
	if v6403 != v6399 {
		goto L1111
	} else {
		goto L1131
	}
L1131:
	;
	v6409 = v6396
	v6411 = v5
	goto L1132
L1132:
	;
	v6439 = int32(0)
	if v6397 == v6439 {
		v6449 = v6439
		goto L1134
	} else {
		goto L1135
	}
L1133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6498 = m.ExcPending
	if v6498 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1134:
	;
	if v6400 == int32(0) {
		goto L1140
	} else {
		goto L1141
	}
L1135:
	;
	v6443 = *(*int32)(unsafe.Add(mBase, uint32(v6397)+4))
	if v6443 <= v6409 {
		v6449 = int32(0)
		goto L1134
	} else {
		goto L1136
	}
L1136:
	;
	v6445 = *(*int32)(unsafe.Add(mBase, uint32(v6397)+12))
	v6449 = v6445 + v6409<<(uint(int32(2))%32)
	goto L1134
L1137:
	;
	goto L1133
L1138:
	;
	v6481 = *(*int32)(unsafe.Add(mBase, uint32(v6458+v6409<<(uint(int32(2))%32))))
	v6482 = *(*int32)(unsafe.Add(mBase, uint32(v6449)))
	v6484 = F_transformExpr(m, l0, v6482, int32(5))
	mBase = m.M
	v6485 = m.ExcPending
	if v6485 != 0 {
		goto L1
	} else {
		goto L1152
	}
L1139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6391)+8)) = v6460
	v6462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6462 != 0 {
		goto L1145
	} else {
		goto L1146
	}
L1140:
	;
	v6460 = int32(0)
	goto L1139
L1141:
	;
	goto L1142
L1142:
	;
	v6455 = *(*int32)(unsafe.Add(mBase, uint32(v6400)+4))
	if base.B2i32(v6449 == int32(0))|base.B2i32(v6455 <= v6409) != 0 {
		v6460 = v6411
		goto L1139
	} else {
		goto L1143
	}
L1143:
	;
	v6458 = *(*int32)(unsafe.Add(mBase, uint32(v6400)+12))
	if v6458 != 0 {
		goto L1138
	} else {
		goto L1144
	}
L1144:
	;
	v6460 = v6411
	goto L1139
L1145:
	;
	v6463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6388)+8)))
	if v6463 == int32(0) {
		goto L1137
	} else {
		goto L1148
	}
L1146:
	;
	v6475 = v5
	goto L1147
L1147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6391)+12)) = v6475
	*(*int32)(unsafe.Add(mBase, uint32(v6361)+32)) = v6391
	v7056 = v6358
	goto L3
L1148:
	;
	v6467 = F_transformExpr(m, l0, v6462, int32(5))
	mBase = m.M
	v6468 = m.ExcPending
	if v6468 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	v6471 = F_coerce_to_specific_type(m, l0, v6467, int32(701), int32(_a_F_transformFromClauseItem_59))
	mBase = m.M
	v6472 = m.ExcPending
	if v6472 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	F_assign_expr_collations(m, l0, v6471)
	mBase = m.M
	v6474 = m.ExcPending
	if v6474 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1151:
	;
	v6475 = v6471
	goto L1147
L1152:
	;
	v6487 = F_coerce_to_specific_type(m, l0, v6484, v6481, int32(_a_F_transformFromClauseItem_60))
	mBase = m.M
	v6488 = m.ExcPending
	if v6488 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1153:
	;
	F_assign_expr_collations(m, l0, v6487)
	mBase = m.M
	v6490 = m.ExcPending
	if v6490 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1154:
	;
	v6493 = F_lappend(m, v6411, v6487)
	mBase = m.M
	v6494 = m.ExcPending
	if v6494 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1155:
	;
	v6409 = v6409 + int32(1)
	v6411 = v6493
	goto L1132
L1156:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6501 = m.ExcPending
	if v6501 != 0 {
		goto L1
	} else {
		goto L1157
	}
L1157:
	;
	v6502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6503 = F_NameListToString(m, v6502)
	mBase = m.M
	v6504 = m.ExcPending
	if v6504 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+176)) = v6503
	F_errmsg(m, int32(_a_F_transformFromClauseItem_61), v37+int32(176))
	mBase = m.M
	v6510 = m.ExcPending
	if v6510 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	v6511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6511)
	mBase = m.M
	v6513 = m.ExcPending
	if v6513 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1160:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(993), int32(_a_F_transformFromClauseItem_62))
	mBase = m.M
	v6518 = m.ExcPending
	if v6518 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1162:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1163:
	;
	v6526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6527 = F_NameListToString(m, v6526)
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+212)) = int32(_a_F_transformFromClauseItem_63)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v6527
	F_errmsg(m, int32(_a_F_transformFromClauseItem_64), v37+int32(208))
	mBase = m.M
	v6536 = m.ExcPending
	if v6536 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	v6537 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6537)
	mBase = m.M
	v6539 = m.ExcPending
	if v6539 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(945), int32(_a_F_transformFromClauseItem_62))
	mBase = m.M
	v6544 = m.ExcPending
	if v6544 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1168:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v6551 = m.ExcPending
	if v6551 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1169:
	;
	v6552 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6553 = F_NameListToString(m, v6552)
	mBase = m.M
	v6554 = m.ExcPending
	if v6554 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+160)) = v6553
	F_errmsg(m, int32(_a_F_transformFromClauseItem_65), v37+int32(160))
	mBase = m.M
	v6560 = m.ExcPending
	if v6560 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	v6561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6561)
	mBase = m.M
	v6563 = m.ExcPending
	if v6563 != 0 {
		goto L1
	} else {
		goto L1172
	}
L1172:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(937), int32(_a_F_transformFromClauseItem_62))
	mBase = m.M
	v6568 = m.ExcPending
	if v6568 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1174:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v6576 = m.ExcPending
	if v6576 != 0 {
		goto L1
	} else {
		goto L1175
	}
L1175:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_66), int32(0))
	mBase = m.M
	v6580 = m.ExcPending
	if v6580 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	v6581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6582 = F_exprLocation(m, v6581)
	mBase = m.M
	F_parser_errposition(m, l0, v6582)
	mBase = m.M
	v6584 = m.ExcPending
	if v6584 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(1145), int32(_a_F_transformFromClauseItem_41))
	mBase = m.M
	v6589 = m.ExcPending
	if v6589 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1179:
	;
	F_errcode(m, int32(403177602))
	mBase = m.M
	v6596 = m.ExcPending
	if v6596 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	v6598 = *(*int32)(unsafe.Add(mBase, uint32(v6388)+4))
	if v6598 != 0 {
		goto L1181
	} else {
		goto L1182
	}
L1181:
	;
	v6599 = *(*int32)(unsafe.Add(mBase, uint32(v6598)+4))
	v6600 = v6599
	goto L1183
L1182:
	;
	v6600 = int32(0)
	goto L1183
L1183:
	;
	v6601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6602 = F_NameListToString(m, v6601)
	mBase = m.M
	v6603 = m.ExcPending
	if v6603 != 0 {
		goto L1
	} else {
		goto L1184
	}
L1184:
	;
	v6604 = *(*int32)(unsafe.Add(mBase, uint32(v6388)+4))
	if v6604 != 0 {
		goto L1185
	} else {
		goto L1186
	}
L1185:
	;
	v6605 = *(*int32)(unsafe.Add(mBase, uint32(v6604)+4))
	v6606 = v6605
	goto L1187
L1186:
	;
	v6606 = v6396
	goto L1187
L1187:
	;
	v6607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v6607 != 0 {
		goto L1188
	} else {
		goto L1189
	}
L1188:
	;
	v6608 = *(*int32)(unsafe.Add(mBase, uint32(v6607)+4))
	v6610 = v6608
	goto L1190
L1189:
	;
	v6610 = int32(0)
	goto L1190
L1190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+200)) = v6610
	*(*int32)(unsafe.Add(mBase, uint32(v37)+196)) = v6606
	*(*int32)(unsafe.Add(mBase, uint32(v37)+192)) = v6602
	F_errmsg_plural(m, int32(_a_F_transformFromClauseItem_67), int32(_a_F_transformFromClauseItem_68), v6600, v37+int32(192))
	mBase = m.M
	v6619 = m.ExcPending
	if v6619 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	v6620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_parser_errposition(m, l0, v6620)
	mBase = m.M
	v6622 = m.ExcPending
	if v6622 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(963), int32(_a_F_transformFromClauseItem_62))
	mBase = m.M
	v6627 = m.ExcPending
	if v6627 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1194:
	;
	v6664 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6664 != 0 {
		goto L1195
	} else {
		goto L1196
	}
L1195:
	;
	v6665 = *(*int32)(unsafe.Add(mBase, uint32(v6664)+4))
	if v6665 <= int32(0) {
		goto L1199
	} else {
		goto L1200
	}
L1196:
	;
	goto L1197
L1197:
	;
	v6989 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+68)) = v6989
	v6991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v6991)
	v6994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v6994 == v6991 {
		goto L1244
	} else {
		goto L1245
	}
L1198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+12)) = v6940
	*(*int32)(unsafe.Add(mBase, uint32(v3708)+8)) = v6946
	goto L1197
L1199:
	;
	v6940 = v5
	v6946 = v5
	goto L1198
L1200:
	;
	goto L1201
L1201:
	;
	v6674 = v5
	v6676 = int32(0)
	v6690 = v5
	v6696 = v5
	goto L1202
L1202:
	;
	v6703 = *(*int32)(unsafe.Add(mBase, uint32(v6664)+12))
	v6707 = *(*int32)(unsafe.Add(mBase, uint32(v6703+v6676<<(uint(int32(2))%32))))
	v6708 = *(*int32)(unsafe.Add(mBase, uint32(v6707)+12))
	v6710 = F_transformExpr(m, l0, v6708, int32(5))
	mBase = m.M
	v6711 = m.ExcPending
	if v6711 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1203:
	;
	v6940 = v6913
	v6946 = v6718
	goto L1198
L1204:
	;
	v6714 = F_coerce_to_specific_type(m, l0, v6710, int32(25), int32(_a_F_transformFromClauseItem_36))
	mBase = m.M
	v6715 = m.ExcPending
	if v6715 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	F_assign_expr_collations(m, l0, v6714)
	mBase = m.M
	v6717 = m.ExcPending
	if v6717 != 0 {
		goto L1
	} else {
		goto L1206
	}
L1206:
	;
	v6718 = F_lappend(m, v6696, v6714)
	mBase = m.M
	v6719 = m.ExcPending
	if v6719 != 0 {
		goto L1
	} else {
		goto L1207
	}
L1207:
	;
	v6720 = *(*int32)(unsafe.Add(mBase, uint32(v6707)+4))
	if v6720 != 0 {
		goto L1210
	} else {
		goto L1211
	}
L1208:
	;
	v6913 = F_lappend(m, v6690, v6883)
	mBase = m.M
	v6914 = m.ExcPending
	if v6914 != 0 {
		goto L1
	} else {
		goto L1242
	}
L1209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6860 = m.ExcPending
	if v6860 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1210:
	;
	if v6690 == int32(0) {
		goto L1213
	} else {
		goto L1214
	}
L1211:
	;
	goto L1212
L1212:
	;
	v6834 = int32(0)
	if v6674 == v6834 {
		v6883 = v6834
		v6884 = int32(1)
		goto L1208
	} else {
		goto L1231
	}
L1213:
	;
	v6832 = F_makeString(m, v6720)
	mBase = m.M
	v6833 = m.ExcPending
	if v6833 != 0 {
		goto L1
	} else {
		goto L1230
	}
L1214:
	;
	v6723 = *(*int32)(unsafe.Add(mBase, uint32(v6690)+4))
	if v6723 <= int32(0) {
		goto L1213
	} else {
		goto L1215
	}
L1215:
	;
	v6726 = *(*int32)(unsafe.Add(mBase, uint32(v6690)+12))
	v6732 = int32(0)
	goto L1216
L1216:
	;
	v6765 = *(*int32)(unsafe.Add(mBase, uint32(v6726+v6732<<(uint(int32(2))%32))))
	if v6765 != 0 {
		goto L1218
	} else {
		goto L1219
	}
L1217:
	;
	goto L1213
L1218:
	;
	v6766 = *(*int32)(unsafe.Add(mBase, uint32(v6765)+4))
	v6769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6766))))
	v6772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6720))))
	if base.B2i32(v6769 == int32(0))|base.B2i32(v6769 != v6772) != 0 {
		v6790 = v6769
		v6791 = v6772
		goto L1222
	} else {
		goto L1223
	}
L1219:
	;
	goto L1220
L1220:
	;
	v6796 = v6732 + int32(1)
	if v6723 != v6796 {
		v6732 = v6796
		goto L1216
	} else {
		goto L1229
	}
L1221:
	;
	if v6790-v6791 == int32(0) {
		goto L1209
	} else {
		goto L1228
	}
L1222:
	;
	goto L1221
L1223:
	;
	v6775 = v6766
	v6776 = v6720
	goto L1224
L1224:
	;
	v6779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6776)+1)))
	v6780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6775)+1)))
	if v6780 == int32(0) {
		v6790 = v6780
		v6791 = v6779
		goto L1222
	} else {
		goto L1226
	}
L1225:
	;
	v6790 = v6780
	v6791 = v6779
	goto L1222
L1226:
	;
	v6783 = int32(1)
	if v6780 == v6779 {
		v6775 = v6775 + v6783
		v6776 = v6776 + v6783
		goto L1224
	} else {
		goto L1227
	}
L1227:
	;
	goto L1225
L1228:
	;
	goto L1220
L1229:
	;
	goto L1217
L1230:
	;
	v6883 = v6832
	v6884 = v6674
	goto L1208
L1231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6841 = m.ExcPending
	if v6841 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1232:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6844 = m.ExcPending
	if v6844 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1233:
	;
	F_errmsg(m, int32(_a_F_transformFromClauseItem_69), int32(0))
	mBase = m.M
	v6848 = m.ExcPending
	if v6848 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1234:
	;
	v6849 = *(*int32)(unsafe.Add(mBase, uint32(v6707)+16))
	F_parser_errposition(m, l0, v6849)
	mBase = m.M
	v6851 = m.ExcPending
	if v6851 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(876), int32(_a_F_transformFromClauseItem_38))
	mBase = m.M
	v6856 = m.ExcPending
	if v6856 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1237:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6863 = m.ExcPending
	if v6863 != 0 {
		goto L1
	} else {
		goto L1238
	}
L1238:
	;
	v6864 = *(*int32)(unsafe.Add(mBase, uint32(v6707)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v6864
	F_errmsg(m, int32(_a_F_transformFromClauseItem_70), v37+int32(224))
	mBase = m.M
	v6870 = m.ExcPending
	if v6870 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	v6871 = *(*int32)(unsafe.Add(mBase, uint32(v6707)+16))
	F_parser_errposition(m, l0, v6871)
	mBase = m.M
	v6873 = m.ExcPending
	if v6873 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	F_errfinish(m, int32(_a_F_transformFromClauseItem_12), int32(867), int32(_a_F_transformFromClauseItem_38))
	mBase = m.M
	v6878 = m.ExcPending
	if v6878 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1242:
	;
	v6916 = v6676 + int32(1)
	v6917 = *(*int32)(unsafe.Add(mBase, uint32(v6664)+4))
	if v6916 < v6917 {
		v6674 = v6884
		v6676 = v6916
		v6690 = v6913
		v6696 = v6718
		goto L1202
	} else {
		goto L1243
	}
L1243:
	;
	goto L1203
L1244:
	;
	v6998 = F_contain_vars_of_level(m, v3708, int32(0))
	mBase = m.M
	v6999 = m.ExcPending
	if v6999 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1245:
	;
	v7000 = int32(1)
	goto L1246
L1246:
	;
	v7001 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v7002 = F_addRangeTableEntryForTableFunc(m, l0, v3708, v7001, v7000)
	mBase = m.M
	v7003 = m.ExcPending
	if v7003 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1247:
	;
	v7000 = v6998
	goto L1246
L1248:
	;
	v7038 = v7002
	goto L4
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7045
	v7049 = F_palloc0(m, int32(8))
	mBase = m.M
	v7050 = m.ExcPending
	if v7050 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7049))) = int32(63)
	v7053 = *(*int32)(unsafe.Add(mBase, uint32(v7038)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7049)+4)) = v7053
	v7056 = v7049
	goto L3
}
