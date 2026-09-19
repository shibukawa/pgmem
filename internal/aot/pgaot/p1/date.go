package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int64
	_ = v776
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v860 float64
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 float64
	_ = v867
	var v870 int64
	_ = v870
	var v873 int64
	_ = v873
	var v878 int64
	_ = v878
	var v880 int64
	_ = v880
	var v885 int64
	_ = v885
	var v887 int64
	_ = v887
	var v891 int64
	_ = v891
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1052 int32
	_ = v1052
	var v1063 int32
	_ = v1063
	var v1079 int32
	_ = v1079
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1171 int32
	_ = v1171
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1227 int32
	_ = v1227
	var v1235 int32
	_ = v1235
	var v1249 int32
	_ = v1249
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1287 int32
	_ = v1287
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1393 int32
	_ = v1393
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1519 int32
	_ = v1519
	var v1527 int32
	_ = v1527
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1559 int32
	_ = v1559
	var v1568 int32
	_ = v1568
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1651 int32
	_ = v1651
	var v1675 int32
	_ = v1675
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1788 int32
	_ = v1788
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1839 int32
	_ = v1839
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1900 int32
	_ = v1900
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1982 int32
	_ = v1982
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1999 int32
	_ = v1999
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2054 int64
	_ = v2054
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2086 int32
	_ = v2086
	var v2088 int64
	_ = v2088
	var v2089 int64
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2093 int64
	_ = v2093
	var v2098 int32
	_ = v2098
	var v2108 int32
	_ = v2108
	var v2114 int32
	_ = v2114
	var v2119 int32
	_ = v2119
	var v2129 int32
	_ = v2129
	var v2142 int32
	_ = v2142
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2208 int32
	_ = v2208
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2256 int32
	_ = v2256
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2270 int32
	_ = v2270
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2285 int32
	_ = v2285
	var v2289 int64
	_ = v2289
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2323 int64
	_ = v2323
	var v2324 int64
	_ = v2324
	var v2326 int32
	_ = v2326
	var v2328 int64
	_ = v2328
	var v2333 int32
	_ = v2333
	var v2343 int32
	_ = v2343
	var v2349 int32
	_ = v2349
	var v2354 int32
	_ = v2354
	var v2364 int32
	_ = v2364
	var v2377 int32
	_ = v2377
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	v9 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(112)
	m.G0 = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+59)) = uint8(v9)
	v42 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(-1)
	if l6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v56 = l4 + int32(8)
	if int32(0) < l2 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v2391 + int32(112)
	return v2390
L5:
	;
	v62 = l4 + int32(20)
	v64 = l4 + int32(12)
	v66 = l4 + int32(16)
	v77 = v9
	v80 = v9
	v87 = v42
	v89 = v9
	v90 = v9
	v92 = v9
	v93 = v9
	v95 = v9
	v98 = v9
	v99 = v9
	goto L8
L6:
	;
	v1752 = v38
	v1756 = v9
	v1763 = v42
	v1766 = v9
	v1768 = v9
	v1771 = v9
	v1774 = v9
	v1775 = v9
	goto L7
L7:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1778 != int32(2) {
		goto L336
	} else {
		goto L337
	}
L8:
	;
	v102 = int32(-1)
	v104 = v89 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1+v104)))
	switch v106 {
	case 0:
		goto L14
	case 1, 6:
		goto L13
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v2390 = v102
		v2391 = v38
		goto L4
	}
L9:
	;
	if v1714 != 0 {
		v2390 = int32(-1)
		v2391 = v38
		goto L4
	} else {
		goto L335
	}
L10:
	;
	v1740 = v89 + int32(1)
	if v1740 != l2 {
		v77 = v1714
		v80 = v1717
		v87 = v1724
		v89 = v1740
		v90 = v1727
		v92 = v1729
		v93 = v1730
		v95 = v1732
		v98 = v1735
		v99 = v1736
		goto L8
	} else {
		goto L334
	}
L11:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v1700&v80 != 0 {
		goto L331
	} else {
		goto L332
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1675 = v77
	v1685 = v87
	v1688 = v1651
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L13:
	;
	v989 = l0 + v104
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v997 = F_DecodeTimezoneAbbrev(m, v89, v990, v38-int32(-64), v38+int32(60), v38+int32(52), l7)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L94
	} else {
		goto L220
	}
L14:
	;
	if v77 != 0 {
		goto L141
	} else {
		goto L142
	}
L15:
	;
	if l6 == int32(0) {
		v2390 = v102
		v2391 = v38
		goto L4
	} else {
		goto L115
	}
L16:
	;
	switch v77 {
	case 0, 3:
		goto L105
	default:
		v2390 = v102
		v2391 = v38
		goto L4
	}
L17:
	;
	if v77 == int32(31) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if l6 == int32(0) {
		v2390 = v102
		v2391 = v38
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v306 = int32(0)
	v308 = int32(10)
	if base.B2i32(v77 == v306)&base.B2i32(v80&v308 != v308) == v306 {
		goto L53
	} else {
		goto L54
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v119 = F_strtol(m, v115, v38+int32(72), int32(10))
	mBase = m.M
	goto L22
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if base.B2i32(v122 == int32(68))|base.B2i32(v119 < int32(0)) != 0 {
		v2390 = int32(-2)
		v2391 = v38
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v129 = v119 + int32(_a_F_DecodeDateTime_0)
	v130 = int32(_a_F_DecodeDateTime_1)
	v131 = base.I32_div_u_s(v129, v130)
	v132 = int32(3)
	v138 = int32(2)
	v143 = base.I32_div_u_s((v131*int32(1073595727)+v129)<<(uint(v138)%32)|v132, v130)
	v146 = v119 + v131*v132 + v143 + int32(_a_F_DecodeDateTime_2)
	v147 = int32(1461)
	v148 = base.I32_div_u_s(v146, v147)
	v151 = v148*int32(-1461) + v146
	v153 = v151 << (uint(v138) % 32)
	if base.Ui32(v147) <= base.Ui32(v153) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v166 = base.I32_div_u_s(v153, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v166 + v148<<(uint(int32(2))%32) - int32(_a_F_DecodeDateTime_3)
	v174 = v164 + int32(123)
	v177 = int32(16)
	v178 = int32(base.Ui32(v174*int32(2141)) >> (uint(v177) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v174 - int32(base.Ui32(v178*int32(_a_F_DecodeDateTime_4))>>(uint(int32(8))%32))
	v188 = base.I32_rem_u_s(v178+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v188 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v193 = int32(0)
	v200 = m.G0
	v202 = v200 - v177
	m.G0 = v202
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	switch v205 - int32(43) {
	case 0, 2:
		goto L30
	default:
		v295 = int32(-1)
		goto L29
	}
L25:
	;
	v159 = base.I32_rem_u_s(v151+int32(305), int32(365))
	v164 = v159
	goto L24
L26:
	;
	goto L27
L27:
	;
	v163 = base.I32_rem_u_s(v151+int32(306), int32(366))
	v164 = v163
	goto L24
L28:
	;
	if v295 != 0 {
		v2390 = v295
		v2391 = v38
		goto L4
	} else {
		goto L52
	}
L29:
	;
	m.G0 = v202 + int32(16)
	goto L28
L30:
	;
	v208 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v215 = F_strtoint(m, v192+int32(1), v202+int32(12))
	mBase = m.M
	v216 = int32(-5)
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v218 == int32(68) {
		v295 = v216
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v222 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v268 = int32(59)
	if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v262))|base.B2i32(base.Ui32(v268) < base.Ui32(v261))|base.B2i32(base.Ui32(v268) < base.Ui32(v264)) != 0 {
		v295 = v216
		goto L29
	} else {
		goto L45
	}
L33:
	;
	if v222 != int32(58) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v252 = F_strlen(m, v192)
	mBase = m.M
	if base.Ui32(v252) < base.Ui32(int32(4)) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	v261 = int32(0)
	v262 = v215
	v264 = v193
	goto L32
L37:
	;
	goto L38
L38:
	;
	v226 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v232 = v202 + int32(12)
	v233 = F_strtoint(m, v221+int32(1), v232)
	mBase = m.M
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v235 == int32(68) {
		v295 = v216
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v239 != int32(58) {
		v261 = v233
		v262 = v215
		v264 = v193
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v242 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v247 = F_strtoint(m, v238+int32(1), v232)
	mBase = m.M
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v249 != int32(68) {
		v261 = v233
		v262 = v215
		v264 = v247
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v295 = v216
	goto L29
L42:
	;
	v261 = int32(0)
	v262 = v215
	v264 = v193
	goto L32
L43:
	;
	goto L44
L44:
	;
	v256 = int32(100)
	v257 = base.I32_div_s(v215, v256)
	v261 = v215 - v257*v256
	v262 = v257
	v264 = v193
	goto L32
L45:
	;
	v274 = int32(60)
	v279 = (v262*v274+v261)*v274 + v264
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v282 == int32(45) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v285 = v279
	goto L48
L47:
	;
	v285 = int32(0) - v279
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v285
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v290 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v291 = int32(-1)
	goto L51
L50:
	;
	v291 = int32(0)
	goto L51
L51:
	;
	v295 = v291
	goto L29
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(_a_F_DecodeDateTime_6)
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = int32(1)
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L53:
	;
	if l6 == int32(0) {
		v2390 = v102
		v2391 = v38
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v487 = F_DecodeDate(m, v482, v80, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L94
	} else {
		goto L103
	}
L56:
	;
	v317 = l0 + v104
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v77 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v471 = F_pg_tzset(m, v318)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L94
	} else {
		goto L99
	}
L58:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if base.Ui32(int32(9)) < base.Ui32((v321-int32(48))&int32(255)) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	switch v77 {
	case 0, 3:
		goto L62
	default:
		v2390 = v102
		v2391 = v38
		goto L4
	}
L61:
	;
	goto L60
L62:
	;
	v328 = int32(_a_F_DecodeDateTime_7)
	if v80&v328 == v328 {
		v2390 = v102
		v2391 = v38
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v332 = int32(45)
	v333 = F___strchrnul(m, v318, v332)
	mBase = m.M
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v335 == v332 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v339 == int32(0) {
		v2390 = v102
		v2391 = v38
		goto L4
	} else {
		goto L68
	}
L65:
	;
	v339 = v333
	goto L67
L66:
	;
	v339 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v342 = int32(0)
	v349 = m.G0
	v351 = v349 - int32(16)
	m.G0 = v351
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	switch v354 - int32(43) {
	case 0, 2:
		goto L71
	default:
		v444 = int32(-1)
		goto L70
	}
L69:
	;
	if v444 != 0 {
		v2390 = v444
		v2391 = v38
		goto L4
	} else {
		goto L93
	}
L70:
	;
	m.G0 = v351 + int32(16)
	goto L69
L71:
	;
	v357 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v364 = F_strtoint(m, v339+int32(1), v351+int32(12))
	mBase = m.M
	v365 = int32(-5)
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v367 == int32(68) {
		v444 = v365
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if v371 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v417 = int32(59)
	if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v411))|base.B2i32(base.Ui32(v417) < base.Ui32(v410))|base.B2i32(base.Ui32(v417) < base.Ui32(v413)) != 0 {
		v444 = v365
		goto L70
	} else {
		goto L86
	}
L74:
	;
	if v371 != int32(58) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v401 = F_strlen(m, v339)
	mBase = m.M
	if base.Ui32(v401) < base.Ui32(int32(4)) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v410 = int32(0)
	v411 = v364
	v413 = v342
	goto L73
L78:
	;
	goto L79
L79:
	;
	v375 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v381 = v351 + int32(12)
	v382 = F_strtoint(m, v370+int32(1), v381)
	mBase = m.M
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v384 == int32(68) {
		v444 = v365
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v388 != int32(58) {
		v410 = v382
		v411 = v364
		v413 = v342
		goto L73
	} else {
		goto L81
	}
L81:
	;
	v391 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v396 = F_strtoint(m, v387+int32(1), v381)
	mBase = m.M
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v398 != int32(68) {
		v410 = v382
		v411 = v364
		v413 = v396
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v444 = v365
	goto L70
L83:
	;
	v410 = int32(0)
	v411 = v364
	v413 = v342
	goto L73
L84:
	;
	goto L85
L85:
	;
	v405 = int32(100)
	v406 = base.I32_div_s(v364, v405)
	v410 = v364 - v406*v405
	v411 = v406
	v413 = v342
	goto L73
L86:
	;
	v423 = int32(60)
	v428 = (v411*v423+v410)*v423 + v413
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v431 == int32(45) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v434 = v428
	goto L89
L88:
	;
	v434 = int32(0) - v428
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v434
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v439 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v440 = int32(-1)
	goto L92
L91:
	;
	v440 = int32(0)
	goto L92
L92:
	;
	v444 = v440
	goto L70
L93:
	;
	v451 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v451)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v454 = F_strlen(m, v453)
	mBase = m.M
	v459 = F_DecodeNumberField(m, v454, v453, v80, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	return int32(0)
L95:
	;
	if v459 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L97:
	;
	goto L98
L98:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v466 | int32(32)
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L99:
	;
	if v471 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v475
	v2390 = int32(-6)
	v2391 = v38
	goto L4
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1675 = int32(0)
	v1685 = v87
	v1688 = v471
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L103:
	;
	if v487 != 0 {
		v2390 = v487
		v2391 = v38
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L105:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v497 = F_DecodeTimeCommon(m, v491, int32(_a_F_DecodeDateTime_8), v38+int32(68), v38+int32(72))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L94
	} else {
		goto L106
	}
L106:
	;
	if v497 != 0 {
		v2390 = v497
		v2391 = v38
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v38)+88))
	if int64(2147483648) <= v499 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v2390 = int32(-2)
	v2391 = v38
	goto L4
L109:
	;
	goto L110
L110:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v499)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v520 = int32(60)
	goto L111
L111:
	;
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v510))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(v511))|(base.B2i32(base.Ui32(v520) < base.Ui32(v512))|base.B2i32(base.Ui32(int32(_a_F_DecodeDateTime_9)) < base.Ui32(v508)))|base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v508)+base.I64_extend_i32_u((v510*v520+v511)*v520+v512)*int64(1000000))) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L113:
	;
	goto L114
L114:
	;
	v2390 = int32(-2)
	v2391 = v38
	goto L4
L115:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v550 = int32(0)
	v557 = m.G0
	v559 = v557 - int32(16)
	m.G0 = v559
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	switch v562 - int32(43) {
	case 0, 2:
		goto L118
	default:
		v652 = int32(-1)
		goto L117
	}
L116:
	;
	if v652 != 0 {
		v2390 = v652
		v2391 = v38
		goto L4
	} else {
		goto L140
	}
L117:
	;
	m.G0 = v559 + int32(16)
	goto L116
L118:
	;
	v565 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v572 = F_strtoint(m, v547+int32(1), v559+int32(12))
	mBase = m.M
	v573 = int32(-5)
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v575 == int32(68) {
		v652 = v573
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	if v579 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v625 = int32(59)
	if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v619))|base.B2i32(base.Ui32(v625) < base.Ui32(v618))|base.B2i32(base.Ui32(v625) < base.Ui32(v621)) != 0 {
		v652 = v573
		goto L117
	} else {
		goto L133
	}
L121:
	;
	if v579 != int32(58) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v609 = F_strlen(m, v547)
	mBase = m.M
	if base.Ui32(v609) < base.Ui32(int32(4)) {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v618 = int32(0)
	v619 = v572
	v621 = v550
	goto L120
L125:
	;
	goto L126
L126:
	;
	v583 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v589 = v559 + int32(12)
	v590 = F_strtoint(m, v578+int32(1), v589)
	mBase = m.M
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v592 == int32(68) {
		v652 = v573
		goto L117
	} else {
		goto L127
	}
L127:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	if v596 != int32(58) {
		v618 = v590
		v619 = v572
		v621 = v550
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v599 = int32(_a_F_DecodeDateTime_5)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v604 = F_strtoint(m, v595+int32(1), v589)
	mBase = m.M
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v606 != int32(68) {
		v618 = v590
		v619 = v572
		v621 = v604
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v652 = v573
	goto L117
L130:
	;
	v618 = int32(0)
	v619 = v572
	v621 = v550
	goto L120
L131:
	;
	goto L132
L132:
	;
	v613 = int32(100)
	v614 = base.I32_div_s(v572, v613)
	v618 = v572 - v614*v613
	v619 = v614
	v621 = v550
	goto L120
L133:
	;
	v631 = int32(60)
	v636 = (v619*v631+v618)*v631 + v621
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	if v639 == int32(45) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v642 = v636
	goto L136
L135:
	;
	v642 = int32(0) - v636
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(72)))) = v642
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	if v647 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v648 = int32(-1)
	goto L139
L138:
	;
	v648 = int32(0)
	goto L139
L139:
	;
	v652 = v648
	goto L117
L140:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v659
	v1651 = v90
	goto L12
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v664 = l0 + v104
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v669 = F_strtol(m, v665, v38+int32(4), int32(10))
	mBase = m.M
	goto L144
L142:
	;
	goto L143
L143:
	;
	v923 = v80 & int32(14)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v926 = F_strlen(m, v925)
	mBase = m.M
	v927 = int32(46)
	v928 = F___strchrnul(m, v925, v927)
	mBase = m.M
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928))))
	if v930 == v927 {
		goto L197
	} else {
		goto L198
	}
L144:
	;
	v670 = int32(-2)
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v672 == int32(68) {
		v2390 = v670
		v2391 = v38
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	v679 = int32(0)
	if base.B2i32(v676 == int32(46))|base.B2i32(v676 == v679) == v679 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L147:
	;
	goto L148
L148:
	;
	if v77 != int32(3) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v917
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L150:
	;
	if v77 != int32(31) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v898 = F_strlen(m, v897)
	mBase = m.M
	v905 = F_DecodeNumberField(m, v898, v897, v80|int32(14), v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L94
	} else {
		goto L191
	}
L153:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L154:
	;
	goto L155
L155:
	;
	if v669 < int32(0) {
		v2390 = v670
		v2391 = v38
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	v695 = v669 + int32(_a_F_DecodeDateTime_0)
	v696 = int32(_a_F_DecodeDateTime_1)
	v697 = base.I32_div_u_s(v695, v696)
	v698 = int32(3)
	v704 = int32(2)
	v709 = base.I32_div_u_s((v697*int32(1073595727)+v695)<<(uint(v704)%32)|v698, v696)
	v712 = v669 + v697*v698 + v709 + int32(_a_F_DecodeDateTime_2)
	v713 = int32(1461)
	v714 = base.I32_div_u_s(v712, v713)
	v717 = v714*int32(-1461) + v712
	v719 = v717 << (uint(v704) % 32)
	if base.Ui32(v713) <= base.Ui32(v719) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v732 = base.I32_div_u_s(v719, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v732 + v714<<(uint(int32(2))%32) - int32(_a_F_DecodeDateTime_3)
	v740 = v730 + int32(123)
	v744 = int32(base.Ui32(v740*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v740 - int32(base.Ui32(v744*int32(_a_F_DecodeDateTime_4))>>(uint(int32(8))%32))
	v751 = int32(1)
	v755 = base.I32_rem_u_s(v744+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v755 + v751
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v759 != int32(46) {
		v917 = v751
		goto L149
	} else {
		goto L161
	}
L158:
	;
	v725 = base.I32_rem_u_s(v717+int32(305), int32(365))
	v730 = v725
	goto L157
L159:
	;
	goto L160
L160:
	;
	v729 = base.I32_rem_u_s(v717+int32(306), int32(366))
	v730 = v729
	goto L157
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v675
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+1)))
	if v763 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L163:
	;
	v870 = base.I64_trunc_sat_f64_s(base.F64_mul(v867, float64(8.64e+10)))
	v873 = base.I64_div_s(v870, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v56))) = uint32(v873)
	v878 = base.I64_extend32_s(v873)*int64(-3600000000) + v870
	v880 = base.I64_div_s(v878, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+int32(4)))) = uint32(v880)
	v885 = base.I64_extend32_s(v880)*int64(-60000000) + v878
	v887 = base.I64_div_s(v885, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v887)
	v891 = v887*int64(4293967296) + v885
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v891)
	goto L190
L164:
	;
	v867 = float64(0)
	goto L163
L165:
	;
	goto L166
L166:
	;
	v768 = v675 + int32(1)
	v769 = int32(_a_F_DecodeDateTime_10)
	v773 = m.G0
	v775 = v773 - int32(32)
	v776 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v775)+24)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v775)+16)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v775)+8)) = v776
	*(*int64)(unsafe.Add(mBase, uint32(v775))) = v776
	v784 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeDateTime[1])))
	if v784 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v853 = F_strlen(m, v768)
	mBase = m.M
	if v852 != v853 {
		goto L162
	} else {
		goto L186
	}
L168:
	;
	v852 = int32(0)
	goto L167
L169:
	;
	goto L170
L170:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeDateTime[2])))
	if v788 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v792 = v768
	goto L174
L172:
	;
	goto L173
L173:
	;
	v802 = v769
	v803 = v784
	goto L177
L174:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792))))
	if v798 == v784 {
		v792 = v792 + int32(1)
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v852 = v792 - v768
	goto L167
L176:
	;
	goto L175
L177:
	;
	v810 = v775 + int32(base.Ui32(v803)>>(uint(int32(3))%32))&int32(28)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	v812 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v811 | v812<<(uint(v803)%32)
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802)+1)))
	if v816 != 0 {
		v802 = v802 + v812
		v803 = v816
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768))))
	if v819 == int32(0) {
		v842 = v768
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v852 = v842 - v768
	goto L167
L181:
	;
	v823 = v768
	v824 = v819
	goto L182
L182:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(base.Ui32(v824)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v832)>>(uint(v824)%32))&int32(1) == int32(0) {
		v842 = v823
		goto L180
	} else {
		goto L184
	}
L183:
	;
	v842 = v840
	goto L180
L184:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823)+1)))
	v840 = v823 + int32(1)
	if v838 != 0 {
		v823 = v840
		v824 = v838
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0])) = int32(0)
	v860 = F_strtod(m, v675, v38+int32(72))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L94
	} else {
		goto L187
	}
L187:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	if v863 != 0 {
		goto L162
	} else {
		goto L188
	}
L188:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[0]))
	if v865 != 0 {
		goto L162
	} else {
		goto L189
	}
L189:
	;
	v867 = v860
	goto L163
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(_a_F_DecodeDateTime_11)
	v917 = v751
	goto L149
L191:
	;
	if v905 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L193:
	;
	goto L194
L194:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v911 != int32(_a_F_DecodeDateTime_7) {
		v2390 = int32(-1)
		v2391 = v38
		goto L4
	} else {
		goto L195
	}
L195:
	;
	v917 = v92
	goto L149
L196:
	;
	v935 = int32(0)
	if base.B2i32(v934 == v935)|v923 == v935 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	v934 = v928
	goto L199
L198:
	;
	v934 = int32(0)
	goto L199
L199:
	;
	goto L196
L200:
	;
	v944 = F_DecodeDate(m, v925, v80, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L94
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if v934 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	if v944 != 0 {
		v2390 = v944
		v2391 = v38
		goto L4
	} else {
		goto L204
	}
L204:
	;
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L205:
	;
	if v80&int32(_a_F_DecodeDateTime_7) != 0 {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v949 = F_strlen(m, v934)
	mBase = m.M
	if base.Ui32(v926-v949) < base.Ui32(int32(3)) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v958 = F_DecodeNumberField(m, v926, v925, v80, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L94
	} else {
		goto L208
	}
L208:
	;
	if int32(0) <= v958 {
		v1675 = int32(0)
		v1685 = v87
		v1688 = v90
		v1690 = v92
		v1691 = v93
		v1693 = v95
		v1696 = v98
		v1697 = v99
		goto L11
	} else {
		goto L209
	}
L209:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L210:
	;
	v968 = v923
	goto L212
L211:
	;
	v968 = int32(0)
	goto L212
L212:
	;
	if base.B2i32(v926 < int32(6))|v968 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v977 = F_DecodeNumberField(m, v926, v925, v80, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L94
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v986 = F_DecodeNumber(m, v926, v925, v93, v80, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L94
	} else {
		goto L218
	}
L216:
	;
	if int32(0) <= v977 {
		v1675 = int32(0)
		v1685 = v87
		v1688 = v90
		v1690 = v92
		v1691 = v93
		v1693 = v95
		v1696 = v98
		v1697 = v99
		goto L11
	} else {
		goto L217
	}
L217:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L218:
	;
	if v986 != 0 {
		v2390 = v986
		v2391 = v38
		goto L4
	} else {
		goto L219
	}
L219:
	;
	v1675 = int32(0)
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L220:
	;
	if v997 != 0 {
		v2390 = v997
		v2391 = v38
		goto L4
	} else {
		goto L221
	}
L221:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v999 == int32(31) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_DecodeDateTime[3])))
	if v1003 != 0 {
		goto L227
	} else {
		goto L228
	}
L223:
	;
	v1249 = v999
	goto L224
L224:
	;
	if v1249 == int32(8) {
		v1714 = v77
		v1717 = v80
		v1724 = v87
		v1727 = v90
		v1729 = v92
		v1730 = v93
		v1732 = v95
		v1735 = v98
		v1736 = v99
		goto L10
	} else {
		goto L270
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1235
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1227
	v1249 = v1235
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_DecodeDateTime[3]))) = v1171
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+12))
	v1199 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1171)+11)))
	v1227 = v1198
	v1235 = v1199
	goto L225
L227:
	;
	goto L232
L228:
	;
	goto L229
L229:
	;
	v1052 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1002))))
	v1063 = int32(_a_F_DecodeDateTime_12)
	v1079 = int32(_a_F_DecodeDateTime_13)
	goto L244
L230:
	;
	if v1041-v1042 == int32(0) {
		v1171 = v1003
		goto L226
	} else {
		goto L243
	}
L232:
	;
	goto L233
L233:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1010 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1011 = v1002
	v1012 = v1003
	v1013 = int32(10)
	v1014 = v1010
	goto L238
L235:
	;
	v1037 = v1003
	v1041 = int32(0)
	goto L236
L236:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037))))
	goto L230
L237:
	;
	v1037 = v1032
	v1041 = v1034
	goto L236
L238:
	;
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012))))
	if base.B2i32(v1014 != v1016)|base.B2i32(v1016 == int32(0)) != 0 {
		v1032 = v1012
		v1034 = v1014
		goto L237
	} else {
		goto L240
	}
L239:
	;
	v1032 = v1026
	v1034 = int32(0)
	goto L237
L240:
	;
	v1022 = v1013 - int32(1)
	if v1022 == int32(0) {
		v1032 = v1012
		v1034 = v1014
		goto L237
	} else {
		goto L241
	}
L241:
	;
	v1025 = int32(1)
	v1026 = v1012 + v1025
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011)+1)))
	if v1027 != 0 {
		v1011 = v1011 + v1025
		v1012 = v1026
		v1013 = v1022
		v1014 = v1027
		goto L238
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	goto L229
L244:
	;
	v1095 = v1063 + (v1079-v1063)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v1096 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1095))))
	v1097 = v1052 - v1096
	if v1097 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1227 = v1149
	v1235 = int32(31)
	goto L225
L246:
	;
	goto L251
L247:
	;
	v1148 = v1097
	goto L248
L248:
	;
	v1149 = int32(0)
	v1153 = base.B2i32(v1148 < v1149)
	if v1148 < v1149 {
		goto L263
	} else {
		goto L264
	}
L249:
	;
	if v1139 == int32(0) {
		v1171 = v1095
		goto L226
	} else {
		goto L262
	}
L251:
	;
	goto L252
L252:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1106 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1107 = v1002
	v1108 = v1095
	v1109 = int32(10)
	v1110 = v1106
	goto L257
L254:
	;
	v1133 = v1095
	v1137 = int32(0)
	goto L255
L255:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133))))
	v1139 = v1137 - v1138
	goto L249
L256:
	;
	v1133 = v1128
	v1137 = v1130
	goto L255
L257:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	if base.B2i32(v1110 != v1112)|base.B2i32(v1112 == int32(0)) != 0 {
		v1128 = v1108
		v1130 = v1110
		goto L256
	} else {
		goto L259
	}
L258:
	;
	v1128 = v1122
	v1130 = int32(0)
	goto L256
L259:
	;
	v1118 = v1109 - int32(1)
	if v1118 == int32(0) {
		v1128 = v1108
		v1130 = v1110
		goto L256
	} else {
		goto L260
	}
L260:
	;
	v1121 = int32(1)
	v1122 = v1108 + v1121
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107)+1)))
	if v1123 != 0 {
		v1107 = v1107 + v1121
		v1108 = v1122
		v1109 = v1118
		v1110 = v1123
		goto L257
	} else {
		goto L261
	}
L261:
	;
	goto L258
L262:
	;
	v1148 = v1139
	goto L248
L263:
	;
	v1154 = v1095 - int32(16)
	goto L265
L264:
	;
	v1154 = v1079
	goto L265
L265:
	;
	if v1148 < v1149 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1157 = v1063
	goto L268
L267:
	;
	v1157 = v1095 + int32(16)
	goto L268
L268:
	;
	if base.Ui32(v1157) <= base.Ui32(v1154) {
		v1063 = v1157
		v1079 = v1154
		goto L244
	} else {
		goto L269
	}
L269:
	;
	goto L245
L270:
	;
	v1276 = int32(1) << (uint(v1249) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1276
	v1278 = int32(-1)
	switch v1249 {
	case 0:
		goto L282
	case 1:
		goto L281
	default:
		v2390 = v1278
		v2391 = v38
		goto L4
	case 5:
		goto L278
	case 6:
		goto L279
	case 7:
		goto L277
	case 9:
		goto L276
	case 16:
		goto L274
	case 17:
		goto L273
	case 18:
		goto L275
	case 23:
		goto L272
	case 28:
		goto L280
	case 31:
		goto L271
	}
L271:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1626 = F_pg_tzset(m, v1625)
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L94
	} else {
		goto L329
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	v1619 = int32(14)
	if base.B2i32(v80&v1619 != v1619)|v77 != 0 {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L328
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	if v77 != 0 {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L327
	}
L274:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+24)) = v1612
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L275:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = base.B2i32(v1609 == int32(1))
	goto L11
L276:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1675 = v77
	v1685 = v1608
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1276 | int32(32)
	if l6 == int32(0) {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L326
	}
L278:
	;
	v1593 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1593
	if l6 == v1593 {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L325
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1276 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L324
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1276 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2390 = v1278
		v2391 = v38
		goto L4
	} else {
		goto L323
	}
L281:
	;
	if v80&int32(8)|(v93|base.B2i32(v80&int32(2) == int32(0))) != 0 {
		goto L320
	} else {
		goto L321
	}
L282:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	switch v1279 - int32(9) {
	case 0, 1, 2:
		goto L284
	case 3:
		goto L289
	case 4:
		goto L288
	case 5:
		goto L287
	case 6:
		goto L286
	case 7:
		goto L285
	default:
		goto L283
	}
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L94
	} else {
		goto L317
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(_a_F_DecodeDateTime_6)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1279
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(_a_F_DecodeDateTime_14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1527 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1527
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	if l6 == v1527 {
		v1675 = v77
		v1685 = v87
		v1688 = v90
		v1690 = v92
		v1691 = v93
		v1693 = v95
		v1696 = v98
		v1697 = v99
		goto L11
	} else {
		goto L316
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L94
	} else {
		goto L304
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L94
	} else {
		goto L303
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L94
	} else {
		goto L291
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(_a_F_DecodeDateTime_6)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, l4, l5, l6)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L94
	} else {
		goto L290
	}
L290:
	;
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L291:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1305 = base.B2i32(int32(2) < v1303)
	if int32(2) < v1303 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1306 = int32(_a_F_DecodeDateTime_3)
	goto L294
L293:
	;
	v1306 = int32(_a_F_DecodeDateTime_15)
	goto L294
L294:
	;
	v1307 = v1300 + v1306
	v1312 = base.I32_div_s(v1307, int32(4))
	v1315 = base.I32_div_s(v1307, int32(-100))
	v1318 = base.I32_div_s(v1307, int32(400))
	if int32(2) < v1303 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1322 = int32(1)
	goto L297
L296:
	;
	v1322 = int32(13)
	goto L297
L297:
	;
	v1327 = base.I32_div_s((v1322+v1303)*int32(_a_F_DecodeDateTime_4), int32(256))
	v1330 = v1299 + v1307*int32(365) + v1312 + v1315 + v1318 + v1327 - int32(_a_F_DecodeDateTime_16)
	v1334 = v1330 + int32(_a_F_DecodeDateTime_0)
	v1335 = int32(_a_F_DecodeDateTime_1)
	v1336 = base.I32_div_u_s(v1334, v1335)
	v1337 = int32(3)
	v1343 = int32(2)
	v1348 = base.I32_div_u_s((v1336*int32(1073595727)+v1334)<<(uint(v1343)%32)|v1337, v1335)
	v1351 = v1330 + v1336*v1337 + v1348 + int32(_a_F_DecodeDateTime_2)
	v1352 = int32(1461)
	v1353 = base.I32_div_u_s(v1351, v1352)
	v1356 = v1353*int32(-1461) + v1351
	v1358 = v1356 << (uint(v1343) % 32)
	if base.Ui32(v1352) <= base.Ui32(v1358) {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L299:
	;
	v1371 = base.I32_div_u_s(v1358, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1371 + v1353<<(uint(int32(2))%32) - int32(_a_F_DecodeDateTime_3)
	v1379 = v1369 + int32(123)
	v1383 = int32(base.Ui32(v1379*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1379 - int32(base.Ui32(v1383*int32(_a_F_DecodeDateTime_4))>>(uint(int32(8))%32))
	v1393 = base.I32_rem_u_s(v1383+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1393 + int32(1)
	goto L298
L300:
	;
	v1364 = base.I32_rem_u_s(v1356+int32(305), int32(365))
	v1369 = v1364
	goto L299
L301:
	;
	goto L302
L302:
	;
	v1368 = base.I32_rem_u_s(v1356+int32(306), int32(366))
	v1369 = v1368
	goto L299
L303:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1408
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1410
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1412
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L304:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1431 = base.B2i32(int32(2) < v1429)
	if int32(2) < v1429 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1432 = int32(_a_F_DecodeDateTime_3)
	goto L307
L306:
	;
	v1432 = int32(_a_F_DecodeDateTime_15)
	goto L307
L307:
	;
	v1433 = v1426 + v1432
	v1438 = base.I32_div_s(v1433, int32(4))
	v1441 = base.I32_div_s(v1433, int32(-100))
	v1444 = base.I32_div_s(v1433, int32(400))
	if int32(2) < v1429 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1448 = int32(1)
	goto L310
L309:
	;
	v1448 = int32(13)
	goto L310
L310:
	;
	v1453 = base.I32_div_s((v1448+v1429)*int32(_a_F_DecodeDateTime_4), int32(256))
	v1456 = v1425 + v1433*int32(365) + v1438 + v1441 + v1444 + v1453 - int32(_a_F_DecodeDateTime_17)
	v1460 = v1456 + int32(_a_F_DecodeDateTime_0)
	v1461 = int32(_a_F_DecodeDateTime_1)
	v1462 = base.I32_div_u_s(v1460, v1461)
	v1463 = int32(3)
	v1469 = int32(2)
	v1474 = base.I32_div_u_s((v1462*int32(1073595727)+v1460)<<(uint(v1469)%32)|v1463, v1461)
	v1477 = v1456 + v1462*v1463 + v1474 + int32(_a_F_DecodeDateTime_2)
	v1478 = int32(1461)
	v1479 = base.I32_div_u_s(v1477, v1478)
	v1482 = v1479*int32(-1461) + v1477
	v1484 = v1482 << (uint(v1469) % 32)
	if base.Ui32(v1478) <= base.Ui32(v1484) {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L312:
	;
	v1497 = base.I32_div_u_s(v1484, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1497 + v1479<<(uint(int32(2))%32) - int32(_a_F_DecodeDateTime_3)
	v1505 = v1495 + int32(123)
	v1509 = int32(base.Ui32(v1505*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1505 - int32(base.Ui32(v1509*int32(_a_F_DecodeDateTime_4))>>(uint(int32(8))%32))
	v1519 = base.I32_rem_u_s(v1509+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1519 + int32(1)
	goto L311
L313:
	;
	v1490 = base.I32_rem_u_s(v1482+int32(305), int32(365))
	v1495 = v1490
	goto L312
L314:
	;
	goto L315
L315:
	;
	v1494 = base.I32_rem_u_s(v1482+int32(306), int32(366))
	v1495 = v1494
	goto L312
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1279
	F_errmsg_internal(m, int32(_a_F_DecodeDateTime_18), v38)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L94
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_DecodeDateTime_19), int32(1390), int32(_a_F_DecodeDateTime_20))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L94
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1568
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = int32(1)
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L321:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if base.Ui32(int32(30)) < base.Ui32(v1559-int32(1)) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1559
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(8)
	goto L320
L323:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1578 - v1579
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L324:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1590
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L325:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1598
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L326:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v1675 = v77
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v1607
	v1696 = v1606
	v1697 = v99
	goto L11
L327:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1675 = v1616
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L328:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1675 = v1624
	v1685 = v87
	v1688 = v90
	v1690 = v92
	v1691 = v93
	v1693 = v95
	v1696 = v98
	v1697 = v99
	goto L11
L329:
	;
	if v1626 != 0 {
		v1651 = v1626
		goto L12
	} else {
		goto L330
	}
L330:
	;
	v2390 = v1278
	v2391 = v38
	goto L4
L331:
	;
	v2390 = int32(-1)
	v2391 = v38
	goto L4
L332:
	;
	goto L333
L333:
	;
	v1714 = v1675
	v1717 = v1700 | v80
	v1724 = v1685
	v1727 = v1688
	v1729 = v1690
	v1730 = v1691
	v1732 = v1693
	v1735 = v1696
	v1736 = v1697
	goto L10
L334:
	;
	goto L9
L335:
	;
	v1752 = v38
	v1756 = v1717
	v1763 = v1724
	v1766 = v1727
	v1768 = v1729
	v1771 = v1732
	v1774 = v1735
	v1775 = v1736
	goto L7
L336:
	;
	v2390 = int32(0)
	v2391 = v1752
	goto L4
L337:
	;
	goto L338
L338:
	;
	v1782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752)+59)))
	if v1768|base.B2i32(v1756&int32(4) == int32(0)) != 0 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	if v1950 != 0 {
		v2390 = v1950
		v2391 = v1752
		goto L4
	} else {
		goto L376
	}
L340:
	;
	if v1756&int32(_a_F_DecodeDateTime_21) != 0 {
		goto L357
	} else {
		goto L358
	}
L341:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1775 != 0 {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1808
	goto L340
L343:
	;
	v1808 = int32(1) - v1788
	goto L342
L344:
	;
	if int32(0) < v1788 {
		goto L343
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	if v1782 != 0 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1950 = int32(-2)
	goto L339
L348:
	;
	if v1788 < int32(0) {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	goto L350
L350:
	;
	if int32(0) < v1788 {
		goto L340
	} else {
		goto L356
	}
L351:
	;
	v1950 = int32(-2)
	goto L339
L352:
	;
	goto L353
L353:
	;
	if base.Ui32(v1788) <= base.Ui32(int32(69)) {
		v1808 = v1788 + int32(2000)
		goto L342
	} else {
		goto L354
	}
L354:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1788) {
		goto L340
	} else {
		goto L355
	}
L355:
	;
	v1808 = v1788 + int32(1900)
	goto L342
L356:
	;
	v1950 = int32(-2)
	goto L339
L357:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v1819 = v1814 + int32(_a_F_DecodeDateTime_15)
	v1821 = base.I32_div_s(v1819, int32(4))
	v1824 = base.I32_div_s(v1819, int32(-100))
	v1827 = base.I32_div_s(v1819, int32(400))
	v1828 = v1813 + v1814*int32(365) + v1821 + v1824 + v1827
	v1830 = v1828 + int32(_a_F_DecodeDateTime_22)
	v1831 = int32(_a_F_DecodeDateTime_1)
	v1832 = base.I32_div_u_s(v1830, v1831)
	v1833 = int32(3)
	v1839 = int32(2)
	v1844 = base.I32_div_u_s((v1832*int32(1073595727)+v1830)<<(uint(v1839)%32)|v1833, v1831)
	v1847 = v1828 + v1832*v1833 + v1844 + int32(_a_F_DecodeDateTime_23)
	v1848 = int32(1461)
	v1849 = base.I32_div_u_s(v1847, v1848)
	v1852 = v1849*int32(-1461) + v1847
	v1854 = v1852 << (uint(v1839) % 32)
	if base.Ui32(v1848) <= base.Ui32(v1854) {
		goto L361
	} else {
		goto L362
	}
L358:
	;
	goto L359
L359:
	;
	if v1756&int32(2) == int32(0) {
		goto L364
	} else {
		goto L365
	}
L360:
	;
	v1867 = base.I32_div_u_s(v1854, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1867 + v1849<<(uint(int32(2))%32) - int32(_a_F_DecodeDateTime_3)
	v1875 = v1865 + int32(123)
	v1879 = int32(base.Ui32(v1875*int32(2141)) >> (uint(int32(16)) % 32))
	v1883 = base.I32_rem_u_s(v1879+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v1883 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v1875 - int32(base.Ui32(v1879*int32(_a_F_DecodeDateTime_4))>>(uint(int32(8))%32))
	goto L359
L361:
	;
	v1860 = base.I32_rem_u_s(v1852+int32(305), int32(365))
	v1865 = v1860
	goto L360
L362:
	;
	goto L363
L363:
	;
	v1864 = base.I32_rem_u_s(v1852+int32(306), int32(366))
	v1865 = v1864
	goto L360
L364:
	;
	if v1756&int32(8) == int32(0) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v1900-int32(13)) {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1950 = int32(-3)
	goto L339
L367:
	;
	v1916 = int32(14)
	if v1756&v1916 != v1916 {
		goto L370
	} else {
		goto L371
	}
L368:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v1910-int32(32)) {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1950 = int32(-3)
	goto L339
L370:
	;
	v1950 = int32(0)
	goto L339
L371:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1922&int32(3) != 0 {
		v1932 = int32(0)
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1932*int32(52)+v1935<<(uint(int32(2))%32))+uint32(_c_F_DecodeDateTime[4])))
	if v1920 <= v1941 {
		goto L370
	} else {
		goto L375
	}
L373:
	;
	v1927 = base.I32_rem_s(v1922, int32(100))
	if v1927 != 0 {
		v1932 = int32(1)
		goto L372
	} else {
		goto L374
	}
L374:
	;
	v1929 = base.I32_rem_s(v1922, int32(400))
	v1932 = base.B2i32(v1929 == int32(0))
	goto L372
L375:
	;
	v1950 = int32(-2)
	goto L339
L376:
	;
	if v1763 == int32(2) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1967 = int32(14)
	if v1756&v1967 != v1967 {
		goto L387
	} else {
		goto L388
	}
L378:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if int32(12) < v1953 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v2390 = int32(-2)
	v2391 = v1752
	goto L4
L380:
	;
	goto L381
L381:
	;
	switch v1763 {
	case 0:
		goto L384
	case 1:
		goto L383
	default:
		goto L377
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v1964
	goto L377
L383:
	;
	if v1953 == int32(12) {
		goto L377
	} else {
		goto L386
	}
L384:
	;
	if v1953 == int32(12) {
		v1964 = int32(0)
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L377
L386:
	;
	v1964 = v1953 + int32(12)
	goto L382
L387:
	;
	v1973 = int32(_a_F_DecodeDateTime_7)
	if v1756&v1973 == v1973 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	goto L389
L389:
	;
	if v1766 != 0 {
		goto L393
	} else {
		goto L394
	}
L390:
	;
	v1977 = int32(1)
	goto L392
L391:
	;
	v1977 = int32(-1)
	goto L392
L392:
	;
	v2390 = v1977
	v2391 = v1752
	goto L4
L393:
	;
	if v1756&int32(268435456) != 0 {
		goto L396
	} else {
		goto L397
	}
L394:
	;
	goto L395
L395:
	;
	if v1771 != 0 {
		goto L433
	} else {
		goto L434
	}
L396:
	;
	v2390 = int32(-1)
	v2391 = v1752
	goto L4
L397:
	;
	goto L398
L398:
	;
	v1982 = v1752 + int32(72)
	v1990 = m.G0
	v1992 = v1990 - int32(32)
	m.G0 = v1992
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1994 <= int32(-4713) {
		goto L403
	} else {
		goto L404
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2142
	goto L395
L400:
	;
	m.G0 = v1992 + int32(32)
	goto L399
L401:
	;
	v2129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2129
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = int64(0)
	v2142 = v2129
	goto L400
L402:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v2014 = int32(60)
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2025 = base.B2i32(int32(2) < v2010)
	if int32(2) < v2010 {
		goto L413
	} else {
		goto L414
	}
L403:
	;
	if v1994 != int32(-4713) {
		goto L401
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	if v1994 <= int32(_a_F_DecodeDateTime_24) {
		goto L408
	} else {
		goto L409
	}
L406:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(10) < v1999 {
		v2010 = v1999
		goto L402
	} else {
		goto L407
	}
L407:
	;
	goto L401
L408:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v2010 = v2004
	goto L402
L409:
	;
	goto L410
L410:
	;
	if v1994 != int32(_a_F_DecodeDateTime_25) {
		goto L401
	} else {
		goto L411
	}
L411:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(5) < v2007 {
		goto L401
	} else {
		goto L412
	}
L412:
	;
	v2010 = v2007
	goto L402
L413:
	;
	v2026 = int32(_a_F_DecodeDateTime_3)
	goto L415
L414:
	;
	v2026 = int32(_a_F_DecodeDateTime_15)
	goto L415
L415:
	;
	v2027 = v2026 + v1994
	v2035 = base.I32_div_u_s(v2027, int32(100))
	v2038 = base.I32_div_u_s(v2027, int32(400))
	if int32(2) < v2010 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v2042 = int32(1)
	goto L418
L417:
	;
	v2042 = int32(13)
	goto L418
L418:
	;
	v2047 = base.I32_div_s((v2042+v2010)*int32(_a_F_DecodeDateTime_4), int32(256))
	v2050 = v2021 + v2027*int32(365) + int32(base.Ui32(v2027)>>(uint(int32(2))%32)) - v2035 + v2038 + v2047 - int32(_a_F_DecodeDateTime_26)
	v2054 = base.I64_extend_i32_s(v2011+(v2012+v2013*v2014)*v2014) + base.I64_extend_i32_s(v2050)*int64(86400)
	if base.B2i32(int32(0) < v2050)&base.B2i32(v2054 < int64(0)) != 0 {
		goto L401
	} else {
		goto L419
	}
L419:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1992)+24)) = v2054 - int64(86400)
	v2073 = F_pg_next_dst_boundary(m, v1992+int32(24), v1992+int32(12), v1992+int32(4), v1992+int32(16), v1992+int32(8), v1992, v1766)
	mBase = m.M
	if v2073 < int32(0) {
		goto L401
	} else {
		goto L420
	}
L420:
	;
	if v2073 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2078
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = v2054 - base.I64_extend_i32_s(v2080)
	v2142 = int32(0) - v2080
	goto L400
L422:
	;
	goto L423
L423:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+12))
	v2088 = v2054 - base.I64_extend_i32_s(v2086)
	v2089 = *(*int64)(unsafe.Add(mBase, uint32(v1992)+16))
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+8))
	v2093 = v2054 - base.I64_extend_i32_s(v2091)
	if base.B2i32(v2089 <= v2088)|base.B2i32(v2089 <= v2093) == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2098
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = v2088
	v2142 = int32(0) - v2086
	goto L400
L425:
	;
	goto L426
L426:
	;
	if base.B2i32(v2093 < v2089)|base.B2i32(v2088 <= v2089) == int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v1992)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2108
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = v2093
	v2142 = int32(0) - v2091
	goto L400
L428:
	;
	goto L429
L429:
	;
	if v2086 < v2091 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v1992)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2114
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = v2088
	v2142 = int32(0) - v2086
	goto L400
L431:
	;
	goto L432
L432:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v1992)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2119
	*(*int64)(unsafe.Add(mBase, uint32(v1982))) = v2093
	v2142 = int32(0) - v2091
	goto L400
L433:
	;
	if v1756&int32(268435456) != 0 {
		goto L436
	} else {
		goto L437
	}
L434:
	;
	goto L435
L435:
	;
	if l6 == int32(0) {
		goto L449
	} else {
		goto L450
	}
L436:
	;
	v2390 = int32(-1)
	v2391 = v1752
	goto L4
L437:
	;
	goto L438
L438:
	;
	v2154 = m.G0
	v2156 = v2154 - int32(288)
	m.G0 = v2156
	v2160 = F_DetermineTimeZoneOffsetInternal(m, l4, v1771, v2156+int32(280))
	mBase = m.M
	v2162 = v2156 + int32(16)
	v2164 = F_strlcpy(m, v2162, v1774, int32(256))
	mBase = m.M
	v2165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156)+16)))
	if v2165 != 0 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2200
	goto L435
L440:
	;
	v2167 = v2162
	v2171 = v2165
	goto L443
L441:
	;
	goto L442
L442:
	;
	v2193 = F_pg_interpret_timezone_abbrev(m, v2156+int32(16), v2156+int32(280), v2156+int32(12), v2156+int32(8), v1771)
	mBase = m.M
	if v2193 != 0 {
		goto L446
	} else {
		goto L447
	}
L443:
	;
	v2173 = F_pg_toupper(m, v2171)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v2167))) = uint8(v2173)
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2167)+1)))
	if v2175 != 0 {
		v2167 = v2167 + int32(1)
		v2171 = v2175
		goto L443
	} else {
		goto L445
	}
L444:
	;
	goto L442
L445:
	;
	goto L444
L446:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2156)+12))
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2156)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2195
	v2200 = int32(0) - v2194
	goto L448
L447:
	;
	v2200 = v2160
	goto L448
L448:
	;
	m.G0 = v2156 + int32(288)
	goto L439
L449:
	;
	v2390 = int32(0)
	v2391 = v1752
	goto L4
L450:
	;
	goto L451
L451:
	;
	v2208 = int32(0)
	if v1756&int32(32) != 0 {
		v2390 = v2208
		v2391 = v1752
		goto L4
	} else {
		goto L452
	}
L452:
	;
	if v1756&int32(268435456) != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v2390 = int32(-1)
	v2391 = v1752
	goto L4
L454:
	;
	goto L455
L455:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeDateTime[5]))
	v2217 = v1752 + int32(72)
	v2225 = m.G0
	v2227 = v2225 - int32(32)
	m.G0 = v2227
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v2229 <= int32(-4713) {
		goto L460
	} else {
		goto L461
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2377
	v2390 = v2208
	v2391 = v1752
	goto L4
L457:
	;
	m.G0 = v2227 + int32(32)
	goto L456
L458:
	;
	v2364 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2364
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = int64(0)
	v2377 = v2364
	goto L457
L459:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v2249 = int32(60)
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2260 = base.B2i32(int32(2) < v2245)
	if int32(2) < v2245 {
		goto L470
	} else {
		goto L471
	}
L460:
	;
	if v2229 != int32(-4713) {
		goto L458
	} else {
		goto L463
	}
L461:
	;
	goto L462
L462:
	;
	if v2229 <= int32(_a_F_DecodeDateTime_24) {
		goto L465
	} else {
		goto L466
	}
L463:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(10) < v2234 {
		v2245 = v2234
		goto L459
	} else {
		goto L464
	}
L464:
	;
	goto L458
L465:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v2245 = v2239
	goto L459
L466:
	;
	goto L467
L467:
	;
	if v2229 != int32(_a_F_DecodeDateTime_25) {
		goto L458
	} else {
		goto L468
	}
L468:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(5) < v2242 {
		goto L458
	} else {
		goto L469
	}
L469:
	;
	v2245 = v2242
	goto L459
L470:
	;
	v2261 = int32(_a_F_DecodeDateTime_3)
	goto L472
L471:
	;
	v2261 = int32(_a_F_DecodeDateTime_15)
	goto L472
L472:
	;
	v2262 = v2261 + v2229
	v2270 = base.I32_div_u_s(v2262, int32(100))
	v2273 = base.I32_div_u_s(v2262, int32(400))
	if int32(2) < v2245 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v2277 = int32(1)
	goto L475
L474:
	;
	v2277 = int32(13)
	goto L475
L475:
	;
	v2282 = base.I32_div_s((v2277+v2245)*int32(_a_F_DecodeDateTime_4), int32(256))
	v2285 = v2256 + v2262*int32(365) + int32(base.Ui32(v2262)>>(uint(int32(2))%32)) - v2270 + v2273 + v2282 - int32(_a_F_DecodeDateTime_26)
	v2289 = base.I64_extend_i32_s(v2246+(v2247+v2248*v2249)*v2249) + base.I64_extend_i32_s(v2285)*int64(86400)
	if base.B2i32(int32(0) < v2285)&base.B2i32(v2289 < int64(0)) != 0 {
		goto L458
	} else {
		goto L476
	}
L476:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2227)+24)) = v2289 - int64(86400)
	v2308 = F_pg_next_dst_boundary(m, v2227+int32(24), v2227+int32(12), v2227+int32(4), v2227+int32(16), v2227+int32(8), v2227, v2215)
	mBase = m.M
	if v2308 < int32(0) {
		goto L458
	} else {
		goto L477
	}
L477:
	;
	if v2308 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2313
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = v2289 - base.I64_extend_i32_s(v2315)
	v2377 = int32(0) - v2315
	goto L457
L479:
	;
	goto L480
L480:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+12))
	v2323 = v2289 - base.I64_extend_i32_s(v2321)
	v2324 = *(*int64)(unsafe.Add(mBase, uint32(v2227)+16))
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+8))
	v2328 = v2289 - base.I64_extend_i32_s(v2326)
	if base.B2i32(v2324 <= v2323)|base.B2i32(v2324 <= v2328) == int32(0) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2333
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = v2323
	v2377 = int32(0) - v2321
	goto L457
L482:
	;
	goto L483
L483:
	;
	if base.B2i32(v2328 < v2324)|base.B2i32(v2323 <= v2324) == int32(0) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2343
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = v2328
	v2377 = int32(0) - v2326
	goto L457
L485:
	;
	goto L486
L486:
	;
	if v2321 < v2326 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2349
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = v2323
	v2377 = int32(0) - v2321
	goto L457
L488:
	;
	goto L489
L489:
	;
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2354
	*(*int64)(unsafe.Add(mBase, uint32(v2217))) = v2328
	v2377 = int32(0) - v2326
	goto L457
}
func F_EncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v19 = l2 & base.B2i32(int32(0) <= v16)
	switch l5 - int32(1) {
	case 0, 3:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v22 {
			v27 = v22
		} else {
			v27 = int32(1) - v22
		}
		v28 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v27)) == int32(0) {
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v39)
			v54 = l6 + int32(2)
		} else {
			v43 = F_pg_ultoa_n(m, v27, l6)
			mBase = m.M
			if v28 <= v43 {
				v54 = l6 + v43
			} else {
				v46 = l6 + v28
				if v43 != 0 {
					base.MemoryCopy(m, v46-v43, l6, v43)
				} else {
				}
				v49 = v28 - v43
				if v49 != 0 {
					base.MemoryFill(m, l6, int32(48), v49)
				} else {
				}
				v54 = v46
			}
		}
		v55 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
		v58 = v54 + int32(1)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v60 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v59)) == int32(0) {
			v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v58))) = uint16(v71)
			v86 = v54 + int32(3)
		} else {
			v75 = F_pg_ultoa_n(m, v59, v58)
			mBase = m.M
			if v60 <= v75 {
				v86 = v58 + v75
			} else {
				v78 = v54 + int32(3)
				if v75 != 0 {
					base.MemoryCopy(m, v78-v75, v58, v75)
				} else {
				}
				v81 = v60 - v75
				if v81 != 0 {
					base.MemoryFill(m, v58, int32(48), v81)
				} else {
				}
				v86 = v78
			}
		}
		v87 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v87)
		v90 = v86 + int32(1)
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v92 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v91)) == int32(0) {
			v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v90))) = uint16(v103)
			v118 = v86 + int32(3)
		} else {
			v107 = F_pg_ultoa_n(m, v91, v90)
			mBase = m.M
			if v92 <= v107 {
				v118 = v90 + v107
			} else {
				v110 = v86 + int32(3)
				if v107 != 0 {
					base.MemoryCopy(m, v110-v107, v90, v107)
				} else {
				}
				v113 = v92 - v107
				if v113 != 0 {
					base.MemoryFill(m, v90, int32(48), v113)
				} else {
				}
				v118 = v110
			}
		}
		if l5 == int32(1) {
			v123 = int32(32)
		} else {
			v123 = int32(84)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v123)
		v126 = v118 + int32(1)
		v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v128 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v127)) == int32(0) {
			v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v126))) = uint16(v139)
			v154 = v118 + int32(3)
		} else {
			v143 = F_pg_ultoa_n(m, v127, v126)
			mBase = m.M
			if v128 <= v143 {
				v154 = v126 + v143
			} else {
				v146 = v118 + int32(3)
				if v143 != 0 {
					base.MemoryCopy(m, v146-v143, v126, v143)
				} else {
				}
				v149 = v128 - v143
				if v149 != 0 {
					base.MemoryFill(m, v126, int32(48), v149)
				} else {
				}
				v154 = v146
			}
		}
		v155 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v155)
		v158 = v154 + int32(1)
		v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v160 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v159)) == int32(0) {
			v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v158))) = uint16(v171)
			v186 = v154 + int32(3)
		} else {
			v175 = F_pg_ultoa_n(m, v159, v158)
			mBase = m.M
			if v160 <= v175 {
				v186 = v158 + v175
			} else {
				v178 = v154 + int32(3)
				if v175 != 0 {
					base.MemoryCopy(m, v178-v175, v158, v175)
				} else {
				}
				v181 = v160 - v175
				if v181 != 0 {
					base.MemoryFill(m, v158, int32(48), v181)
				} else {
				}
				v186 = v178
			}
		}
		v187 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v187)
		v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v196 = v191 >> (uint(int32(31)) % 32)
		v200 = F_pg_ultostr_zeropad(m, v186+int32(1), v191^v196-v196, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v307 = v200
		} else {
			v205 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v205)
			v208 = l1 >> (uint(int32(31)) % 32)
			v210 = l1 ^ v208 - v208
			v212 = base.I32_div_s(v210, int32(10))
			v215 = v212*int32(-10) + v210
			if v215 != 0 {
				v217 = v215 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+6)) = uint8(v217)
				v223 = v200 + int32(7)
			} else {
				v223 = v200 + int32(6)
			}
			v225 = base.I32_div_s(v210, int32(100))
			v228 = v225*int32(-10) + v212
			v229 = v215 | v228
			if v229 == int32(0) {
				v237 = v200 + int32(5)
			} else {
				v235 = v228 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+5)) = uint8(v235)
				v237 = v223
			}
			v239 = base.I32_div_s(v210, int32(1000))
			v242 = v225 + v239*int32(-10)
			v243 = v229 | v242
			if v243 == int32(0) {
				v251 = v200 + int32(4)
			} else {
				v249 = v242 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+4)) = uint8(v249)
				v251 = v237
			}
			v253 = base.I32_div_s(v210, int32(_a_F_EncodeDateTime_0))
			v256 = v239 + v253*int32(-10)
			v257 = v243 | v256
			if v257 == int32(0) {
				v265 = v200 + int32(3)
			} else {
				v263 = v256 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+3)) = uint8(v263)
				v265 = v251
			}
			v267 = base.I32_div_s(v210, int32(_a_F_EncodeDateTime_1))
			v270 = v253 + v267*int32(-10)
			v271 = v257 | v270
			if v271 == int32(0) {
				v279 = v200 + int32(2)
			} else {
				v277 = v270 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+2)) = uint8(v277)
				v279 = v265
			}
			v281 = base.I32_div_s(v210, int32(_a_F_EncodeDateTime_2))
			v284 = v281*int32(-10) + v267
			if v271|v284 == int32(0) {
				v293 = v200 + int32(1)
			} else {
				v291 = v284 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)) = uint8(v291)
				v293 = v279
			}
			if base.Ui32(int32(19)) <= base.Ui32(v267+int32(9)) {
				v300 = F_pg_ultostr(m, v200+int32(1), v210)
				mBase = m.M
				v301 = v300
			} else {
				v301 = v293
			}
			v307 = v301
		}
		if v19 == int32(0) {
			v1684 = v307
		} else {
			if l3 <= int32(0) {
				v318 = int32(43)
			} else {
				v318 = int32(45)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v318)
			v321 = l3 >> (uint(int32(31)) % 32)
			v323 = l3 ^ v321 - v321
			v325 = base.I32_div_s(v323, int32(3600))
			v326 = int32(-60)
			v329 = base.I32_div_s(v323, int32(60))
			v330 = v325*v326 + v329
			v332 = v307 + int32(1)
			v335 = v329*v326 + v323
			if v335 != 0 {
				v336 = int32(2)
				v337 = F_pg_ultostr_zeropad(m, v332, v325, v336)
				mBase = m.M
				v338 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v337))) = uint8(v338)
				v343 = F_pg_ultostr_zeropad(m, v337+int32(1), v330, v336)
				mBase = m.M
				v350 = v343
				v351 = v335
				v352 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v352)
				v357 = F_pg_ultostr_zeropad(m, v350+int32(1), v351, int32(2))
				mBase = m.M
				v358 = v357
			} else {
				v345 = F_pg_ultostr_zeropad(m, v332, v325, int32(2))
				mBase = m.M
				if l5 == int32(4) {
					v350 = v345
					v351 = v330
					v352 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v352)
					v357 = F_pg_ultostr_zeropad(m, v350+int32(1), v351, int32(2))
					mBase = m.M
					v358 = v357
				} else {
					if v330 == int32(0) {
						v358 = v345
					} else {
						v350 = v345
						v351 = v330
						v352 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v352)
						v357 = F_pg_ultostr_zeropad(m, v350+int32(1), v351, int32(2))
						mBase = m.M
						v358 = v357
					}
				}
			}
			v1684 = v358
		}
		v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v1685 <= int32(0) {
			v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
			*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
			v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
			*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
			v1696 = v1684 + int32(3)
		} else {
			v1696 = v1684
		}
		v1697 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
		m.G0 = v14 + int32(48)
		return
	case 1:
		v363 = *(*int32)(unsafe.Add(mBase, _c_F_EncodeDateTime[3]))
		v365 = base.B2i32(v363 == int32(1))
		if v363 == int32(1) {
			v366 = int32(12)
		} else {
			v366 = int32(16)
		}
		v368 = *(*int32)(unsafe.Add(mBase, uint32(l0+v366)))
		v369 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v368)) == int32(0) {
			v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v368<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v380)
			v395 = l6 + int32(2)
		} else {
			v384 = F_pg_ultoa_n(m, v368, l6)
			mBase = m.M
			if v369 <= v384 {
				v395 = l6 + v384
			} else {
				v387 = l6 + v369
				if v384 != 0 {
					base.MemoryCopy(m, v387-v384, l6, v384)
				} else {
				}
				v390 = v369 - v384
				if v390 != 0 {
					base.MemoryFill(m, l6, int32(48), v390)
				} else {
				}
				v395 = v387
			}
		}
		v396 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v395))) = uint8(v396)
		v399 = v395 + int32(1)
		if v363 == int32(1) {
			v402 = int32(16)
		} else {
			v402 = int32(12)
		}
		v404 = *(*int32)(unsafe.Add(mBase, uint32(l0+v402)))
		v405 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v404)) == int32(0) {
			v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v399))) = uint16(v416)
			v431 = v395 + int32(3)
		} else {
			v420 = F_pg_ultoa_n(m, v404, v399)
			mBase = m.M
			if v405 <= v420 {
				v431 = v399 + v420
			} else {
				v423 = v395 + int32(3)
				if v420 != 0 {
					base.MemoryCopy(m, v423-v420, v399, v420)
				} else {
				}
				v426 = v405 - v420
				if v426 != 0 {
					base.MemoryFill(m, v399, int32(48), v426)
				} else {
				}
				v431 = v423
			}
		}
		v432 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v431))) = uint8(v432)
		v434 = int32(1)
		v435 = v431 + v434
		v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v436 {
			v441 = v436
		} else {
			v441 = v434 - v436
		}
		v442 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v441)) == int32(0) {
			v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v435))) = uint16(v453)
			v468 = v431 + int32(3)
		} else {
			v457 = F_pg_ultoa_n(m, v441, v435)
			mBase = m.M
			if v442 <= v457 {
				v468 = v435 + v457
			} else {
				v460 = v431 + int32(5)
				if v457 != 0 {
					base.MemoryCopy(m, v460-v457, v435, v457)
				} else {
				}
				v463 = v442 - v457
				if v463 != 0 {
					base.MemoryFill(m, v435, int32(48), v463)
				} else {
				}
				v468 = v460
			}
		}
		v469 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v469)
		v472 = v468 + int32(1)
		v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v474 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v473)) == int32(0) {
			v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v472))) = uint16(v485)
			v500 = v468 + int32(3)
		} else {
			v489 = F_pg_ultoa_n(m, v473, v472)
			mBase = m.M
			if v474 <= v489 {
				v500 = v472 + v489
			} else {
				v492 = v468 + int32(3)
				if v489 != 0 {
					base.MemoryCopy(m, v492-v489, v472, v489)
				} else {
				}
				v495 = v474 - v489
				if v495 != 0 {
					base.MemoryFill(m, v472, int32(48), v495)
				} else {
				}
				v500 = v492
			}
		}
		v501 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v501)
		v504 = v500 + int32(1)
		v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v506 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v505)) == int32(0) {
			v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v505<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v504))) = uint16(v517)
			v532 = v500 + int32(3)
		} else {
			v521 = F_pg_ultoa_n(m, v505, v504)
			mBase = m.M
			if v506 <= v521 {
				v532 = v504 + v521
			} else {
				v524 = v500 + int32(3)
				if v521 != 0 {
					base.MemoryCopy(m, v524-v521, v504, v521)
				} else {
				}
				v527 = v506 - v521
				if v527 != 0 {
					base.MemoryFill(m, v504, int32(48), v527)
				} else {
				}
				v532 = v524
			}
		}
		v533 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v533)
		v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v542 = v537 >> (uint(int32(31)) % 32)
		v546 = F_pg_ultostr_zeropad(m, v532+int32(1), v537^v542-v542, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v653 = v546
		} else {
			v551 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v546))) = uint8(v551)
			v554 = l1 >> (uint(int32(31)) % 32)
			v556 = l1 ^ v554 - v554
			v558 = base.I32_div_s(v556, int32(10))
			v561 = v558*int32(-10) + v556
			if v561 != 0 {
				v563 = v561 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+6)) = uint8(v563)
				v569 = v546 + int32(7)
			} else {
				v569 = v546 + int32(6)
			}
			v571 = base.I32_div_s(v556, int32(100))
			v574 = v571*int32(-10) + v558
			v575 = v561 | v574
			if v575 == int32(0) {
				v583 = v546 + int32(5)
			} else {
				v581 = v574 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+5)) = uint8(v581)
				v583 = v569
			}
			v585 = base.I32_div_s(v556, int32(1000))
			v588 = v571 + v585*int32(-10)
			v589 = v575 | v588
			if v589 == int32(0) {
				v597 = v546 + int32(4)
			} else {
				v595 = v588 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+4)) = uint8(v595)
				v597 = v583
			}
			v599 = base.I32_div_s(v556, int32(_a_F_EncodeDateTime_0))
			v602 = v585 + v599*int32(-10)
			v603 = v589 | v602
			if v603 == int32(0) {
				v611 = v546 + int32(3)
			} else {
				v609 = v602 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+3)) = uint8(v609)
				v611 = v597
			}
			v613 = base.I32_div_s(v556, int32(_a_F_EncodeDateTime_1))
			v616 = v599 + v613*int32(-10)
			v617 = v603 | v616
			if v617 == int32(0) {
				v625 = v546 + int32(2)
			} else {
				v623 = v616 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+2)) = uint8(v623)
				v625 = v611
			}
			v627 = base.I32_div_s(v556, int32(_a_F_EncodeDateTime_2))
			v630 = v627*int32(-10) + v613
			if v617|v630 == int32(0) {
				v639 = v546 + int32(1)
			} else {
				v637 = v630 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v546)+1)) = uint8(v637)
				v639 = v625
			}
			if base.Ui32(int32(19)) <= base.Ui32(v613+int32(9)) {
				v646 = F_pg_ultostr(m, v546+int32(1), v556)
				mBase = m.M
				v647 = v646
			} else {
				v647 = v639
			}
			v653 = v647
		}
		if v19 == int32(0) {
			v1684 = v653
			v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1685 <= int32(0) {
				v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
				v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
				v1696 = v1684 + int32(3)
			} else {
				v1696 = v1684
			}
			v1697 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(10)
				v662 = F_pg_sprintf(m, v653, int32(_a_F_EncodeDateTime_3), v14+int32(16))
				mBase = m.M
				v663 = m.ExcPending
				if v663 != 0 {
					return
				} else {
					v664 = F_strlen(m, v653)
					mBase = m.M
					v1684 = v664 + v653
					v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1685 <= int32(0) {
						v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
						v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
						v1696 = v1684 + int32(3)
					} else {
						v1696 = v1684
					}
					v1697 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				if l3 <= int32(0) {
					v670 = int32(43)
				} else {
					v670 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v653))) = uint8(v670)
				v673 = l3 >> (uint(int32(31)) % 32)
				v675 = l3 ^ v673 - v673
				v677 = base.I32_div_s(v675, int32(3600))
				v678 = int32(-60)
				v681 = base.I32_div_s(v675, int32(60))
				v682 = v677*v678 + v681
				v684 = v653 + int32(1)
				v687 = v681*v678 + v675
				if v687 != 0 {
					v688 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v677)) == int32(0) {
						v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v677<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v684))) = uint16(v699)
						v714 = v653 + int32(3)
					} else {
						v703 = F_pg_ultoa_n(m, v677, v684)
						mBase = m.M
						if v688 <= v703 {
							v714 = v684 + v703
						} else {
							v706 = v653 + int32(3)
							if v703 != 0 {
								base.MemoryCopy(m, v706-v703, v684, v703)
							} else {
							}
							v709 = v688 - v703
							if v709 != 0 {
								base.MemoryFill(m, v684, int32(48), v709)
							} else {
							}
							v714 = v706
						}
					}
					v715 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v714))) = uint8(v715)
					v718 = v714 + int32(1)
					v719 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v682)) == int32(0) {
						v730 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v682<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v718))) = uint16(v730)
						v745 = v714 + int32(3)
					} else {
						v734 = F_pg_ultoa_n(m, v682, v718)
						mBase = m.M
						if v719 <= v734 {
							v745 = v718 + v734
						} else {
							v737 = v714 + int32(3)
							if v734 != 0 {
								base.MemoryCopy(m, v737-v734, v718, v734)
							} else {
							}
							v740 = v719 - v734
							if v740 != 0 {
								base.MemoryFill(m, v718, int32(48), v740)
							} else {
							}
							v745 = v737
						}
					}
					v776 = v687
					v777 = v745
					v778 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v778)
					v781 = v777 + int32(1)
					v782 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v776)) == int32(0) {
						v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v776<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v781))) = uint16(v793)
						v808 = v777 + int32(3)
					} else {
						v797 = F_pg_ultoa_n(m, v776, v781)
						mBase = m.M
						if v782 <= v797 {
							v808 = v781 + v797
						} else {
							v800 = v777 + int32(3)
							if v797 != 0 {
								base.MemoryCopy(m, v800-v797, v781, v797)
							} else {
							}
							v803 = v782 - v797
							if v803 != 0 {
								base.MemoryFill(m, v781, int32(48), v803)
							} else {
							}
							v808 = v800
						}
					}
					v1684 = v808
				} else {
					v746 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v677)) == int32(0) {
						v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v677<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v684))) = uint16(v757)
						v772 = v653 + int32(3)
					} else {
						v761 = F_pg_ultoa_n(m, v677, v684)
						mBase = m.M
						if v746 <= v761 {
							v772 = v684 + v761
						} else {
							v764 = v653 + int32(3)
							if v761 != 0 {
								base.MemoryCopy(m, v764-v761, v684, v761)
							} else {
							}
							v767 = v746 - v761
							if v767 != 0 {
								base.MemoryFill(m, v684, int32(48), v767)
							} else {
							}
							v772 = v764
						}
					}
					if v682 == int32(0) {
						v1684 = v772
					} else {
						v776 = v682
						v777 = v772
						v778 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v778)
						v781 = v777 + int32(1)
						v782 = int32(2)
						if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v776)) == int32(0) {
							v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v776<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v781))) = uint16(v793)
							v808 = v777 + int32(3)
						} else {
							v797 = F_pg_ultoa_n(m, v776, v781)
							mBase = m.M
							if v782 <= v797 {
								v808 = v781 + v797
							} else {
								v800 = v777 + int32(3)
								if v797 != 0 {
									base.MemoryCopy(m, v800-v797, v781, v797)
								} else {
								}
								v803 = v782 - v797
								if v803 != 0 {
									base.MemoryFill(m, v781, int32(48), v803)
								} else {
								}
								v808 = v800
							}
						}
						v1684 = v808
					}
				}
				v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1685 <= int32(0) {
					v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
					v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
					v1696 = v1684 + int32(3)
				} else {
					v1696 = v1684
				}
				v1697 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
				m.G0 = v14 + int32(48)
				return
			}
		}
	case 2:
		v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v810 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v809)) == int32(0) {
			v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v809<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v821)
			v836 = l6 + int32(2)
		} else {
			v825 = F_pg_ultoa_n(m, v809, l6)
			mBase = m.M
			if v810 <= v825 {
				v836 = l6 + v825
			} else {
				v828 = l6 + v810
				if v825 != 0 {
					base.MemoryCopy(m, v828-v825, l6, v825)
				} else {
				}
				v831 = v810 - v825
				if v831 != 0 {
					base.MemoryFill(m, l6, int32(48), v831)
				} else {
				}
				v836 = v828
			}
		}
		v837 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v836))) = uint8(v837)
		v840 = v836 + int32(1)
		v841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v842 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v841)) == int32(0) {
			v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v841<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v840))) = uint16(v853)
			v868 = v836 + int32(3)
		} else {
			v857 = F_pg_ultoa_n(m, v841, v840)
			mBase = m.M
			if v842 <= v857 {
				v868 = v840 + v857
			} else {
				v860 = v836 + int32(3)
				if v857 != 0 {
					base.MemoryCopy(m, v860-v857, v840, v857)
				} else {
				}
				v863 = v842 - v857
				if v863 != 0 {
					base.MemoryFill(m, v840, int32(48), v863)
				} else {
				}
				v868 = v860
			}
		}
		v869 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v868))) = uint8(v869)
		v871 = int32(1)
		v872 = v868 + v871
		v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v873 {
			v878 = v873
		} else {
			v878 = v871 - v873
		}
		v879 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v878)) == int32(0) {
			v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v878<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v872))) = uint16(v890)
			v905 = v868 + int32(3)
		} else {
			v894 = F_pg_ultoa_n(m, v878, v872)
			mBase = m.M
			if v879 <= v894 {
				v905 = v872 + v894
			} else {
				v897 = v868 + int32(5)
				if v894 != 0 {
					base.MemoryCopy(m, v897-v894, v872, v894)
				} else {
				}
				v900 = v879 - v894
				if v900 != 0 {
					base.MemoryFill(m, v872, int32(48), v900)
				} else {
				}
				v905 = v897
			}
		}
		v906 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v905))) = uint8(v906)
		v909 = v905 + int32(1)
		v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v911 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v910)) == int32(0) {
			v922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v910<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v909))) = uint16(v922)
			v937 = v905 + int32(3)
		} else {
			v926 = F_pg_ultoa_n(m, v910, v909)
			mBase = m.M
			if v911 <= v926 {
				v937 = v909 + v926
			} else {
				v929 = v905 + int32(3)
				if v926 != 0 {
					base.MemoryCopy(m, v929-v926, v909, v926)
				} else {
				}
				v932 = v911 - v926
				if v932 != 0 {
					base.MemoryFill(m, v909, int32(48), v932)
				} else {
				}
				v937 = v929
			}
		}
		v938 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v938)
		v941 = v937 + int32(1)
		v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v943 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v942)) == int32(0) {
			v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v942<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v941))) = uint16(v954)
			v969 = v937 + int32(3)
		} else {
			v958 = F_pg_ultoa_n(m, v942, v941)
			mBase = m.M
			if v943 <= v958 {
				v969 = v941 + v958
			} else {
				v961 = v937 + int32(3)
				if v958 != 0 {
					base.MemoryCopy(m, v961-v958, v941, v958)
				} else {
				}
				v964 = v943 - v958
				if v964 != 0 {
					base.MemoryFill(m, v941, int32(48), v964)
				} else {
				}
				v969 = v961
			}
		}
		v970 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v969))) = uint8(v970)
		v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v979 = v974 >> (uint(int32(31)) % 32)
		v983 = F_pg_ultostr_zeropad(m, v969+int32(1), v974^v979-v979, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v1090 = v983
		} else {
			v988 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v983))) = uint8(v988)
			v991 = l1 >> (uint(int32(31)) % 32)
			v993 = l1 ^ v991 - v991
			v995 = base.I32_div_s(v993, int32(10))
			v998 = v995*int32(-10) + v993
			if v998 != 0 {
				v1000 = v998 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+6)) = uint8(v1000)
				v1006 = v983 + int32(7)
			} else {
				v1006 = v983 + int32(6)
			}
			v1008 = base.I32_div_s(v993, int32(100))
			v1011 = v1008*int32(-10) + v995
			v1012 = v998 | v1011
			if v1012 == int32(0) {
				v1020 = v983 + int32(5)
			} else {
				v1018 = v1011 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+5)) = uint8(v1018)
				v1020 = v1006
			}
			v1022 = base.I32_div_s(v993, int32(1000))
			v1025 = v1008 + v1022*int32(-10)
			v1026 = v1012 | v1025
			if v1026 == int32(0) {
				v1034 = v983 + int32(4)
			} else {
				v1032 = v1025 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+4)) = uint8(v1032)
				v1034 = v1020
			}
			v1036 = base.I32_div_s(v993, int32(_a_F_EncodeDateTime_0))
			v1039 = v1022 + v1036*int32(-10)
			v1040 = v1026 | v1039
			if v1040 == int32(0) {
				v1048 = v983 + int32(3)
			} else {
				v1046 = v1039 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+3)) = uint8(v1046)
				v1048 = v1034
			}
			v1050 = base.I32_div_s(v993, int32(_a_F_EncodeDateTime_1))
			v1053 = v1036 + v1050*int32(-10)
			v1054 = v1040 | v1053
			if v1054 == int32(0) {
				v1062 = v983 + int32(2)
			} else {
				v1060 = v1053 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+2)) = uint8(v1060)
				v1062 = v1048
			}
			v1064 = base.I32_div_s(v993, int32(_a_F_EncodeDateTime_2))
			v1067 = v1064*int32(-10) + v1050
			if v1054|v1067 == int32(0) {
				v1076 = v983 + int32(1)
			} else {
				v1074 = v1067 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v983)+1)) = uint8(v1074)
				v1076 = v1062
			}
			if base.Ui32(int32(19)) <= base.Ui32(v1050+int32(9)) {
				v1083 = F_pg_ultostr(m, v983+int32(1), v993)
				mBase = m.M
				v1084 = v1083
			} else {
				v1084 = v1076
			}
			v1090 = v1084
		}
		if v19 == int32(0) {
			v1684 = v1090
			v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1685 <= int32(0) {
				v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
				v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
				v1696 = v1684 + int32(3)
			} else {
				v1696 = v1684
			}
			v1697 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(10)
				v1099 = F_pg_sprintf(m, v1090, int32(_a_F_EncodeDateTime_3), v14+int32(32))
				mBase = m.M
				v1100 = m.ExcPending
				if v1100 != 0 {
					return
				} else {
					v1101 = F_strlen(m, v1090)
					mBase = m.M
					v1684 = v1101 + v1090
					v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1685 <= int32(0) {
						v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
						v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
						v1696 = v1684 + int32(3)
					} else {
						v1696 = v1684
					}
					v1697 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				if l3 <= int32(0) {
					v1107 = int32(43)
				} else {
					v1107 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v1090))) = uint8(v1107)
				v1110 = l3 >> (uint(int32(31)) % 32)
				v1112 = l3 ^ v1110 - v1110
				v1114 = base.I32_div_s(v1112, int32(3600))
				v1115 = int32(-60)
				v1118 = base.I32_div_s(v1112, int32(60))
				v1119 = v1114*v1115 + v1118
				v1121 = v1090 + int32(1)
				v1124 = v1118*v1115 + v1112
				if v1124 != 0 {
					v1125 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1114)) == int32(0) {
						v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1121))) = uint16(v1136)
						v1151 = v1090 + int32(3)
					} else {
						v1140 = F_pg_ultoa_n(m, v1114, v1121)
						mBase = m.M
						if v1125 <= v1140 {
							v1151 = v1121 + v1140
						} else {
							v1143 = v1090 + int32(3)
							if v1140 != 0 {
								base.MemoryCopy(m, v1143-v1140, v1121, v1140)
							} else {
							}
							v1146 = v1125 - v1140
							if v1146 != 0 {
								base.MemoryFill(m, v1121, int32(48), v1146)
							} else {
							}
							v1151 = v1143
						}
					}
					v1152 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1151))) = uint8(v1152)
					v1155 = v1151 + int32(1)
					v1156 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1119)) == int32(0) {
						v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1119<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1155))) = uint16(v1167)
						v1182 = v1151 + int32(3)
					} else {
						v1171 = F_pg_ultoa_n(m, v1119, v1155)
						mBase = m.M
						if v1156 <= v1171 {
							v1182 = v1155 + v1171
						} else {
							v1174 = v1151 + int32(3)
							if v1171 != 0 {
								base.MemoryCopy(m, v1174-v1171, v1155, v1171)
							} else {
							}
							v1177 = v1156 - v1171
							if v1177 != 0 {
								base.MemoryFill(m, v1155, int32(48), v1177)
							} else {
							}
							v1182 = v1174
						}
					}
					v1213 = v1124
					v1214 = v1182
					v1215 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1214))) = uint8(v1215)
					v1218 = v1214 + int32(1)
					v1219 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1213)) == int32(0) {
						v1230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1218))) = uint16(v1230)
						v1245 = v1214 + int32(3)
					} else {
						v1234 = F_pg_ultoa_n(m, v1213, v1218)
						mBase = m.M
						if v1219 <= v1234 {
							v1245 = v1218 + v1234
						} else {
							v1237 = v1214 + int32(3)
							if v1234 != 0 {
								base.MemoryCopy(m, v1237-v1234, v1218, v1234)
							} else {
							}
							v1240 = v1219 - v1234
							if v1240 != 0 {
								base.MemoryFill(m, v1218, int32(48), v1240)
							} else {
							}
							v1245 = v1237
						}
					}
					v1684 = v1245
				} else {
					v1183 = int32(2)
					if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1114)) == int32(0) {
						v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1121))) = uint16(v1194)
						v1209 = v1090 + int32(3)
					} else {
						v1198 = F_pg_ultoa_n(m, v1114, v1121)
						mBase = m.M
						if v1183 <= v1198 {
							v1209 = v1121 + v1198
						} else {
							v1201 = v1090 + int32(3)
							if v1198 != 0 {
								base.MemoryCopy(m, v1201-v1198, v1121, v1198)
							} else {
							}
							v1204 = v1183 - v1198
							if v1204 != 0 {
								base.MemoryFill(m, v1121, int32(48), v1204)
							} else {
							}
							v1209 = v1201
						}
					}
					if v1119 == int32(0) {
						v1684 = v1209
					} else {
						v1213 = v1119
						v1214 = v1209
						v1215 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v1214))) = uint8(v1215)
						v1218 = v1214 + int32(1)
						v1219 = int32(2)
						if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1213)) == int32(0) {
							v1230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v1218))) = uint16(v1230)
							v1245 = v1214 + int32(3)
						} else {
							v1234 = F_pg_ultoa_n(m, v1213, v1218)
							mBase = m.M
							if v1219 <= v1234 {
								v1245 = v1218 + v1234
							} else {
								v1237 = v1214 + int32(3)
								if v1234 != 0 {
									base.MemoryCopy(m, v1237-v1234, v1218, v1234)
								} else {
								}
								v1240 = v1219 - v1234
								if v1240 != 0 {
									base.MemoryFill(m, v1218, int32(48), v1240)
								} else {
								}
								v1245 = v1237
							}
						}
						v1684 = v1245
					}
				}
				v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1685 <= int32(0) {
					v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
					v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
					v1696 = v1684 + int32(3)
				} else {
					v1696 = v1684
				}
				v1697 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
				m.G0 = v14 + int32(48)
				return
			}
		}
	default:
		v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v1252 = base.B2i32(int32(2) < v1250)
		if int32(2) < v1250 {
			v1253 = int32(_a_F_EncodeDateTime_4)
		} else {
			v1253 = int32(_a_F_EncodeDateTime_5)
		}
		v1254 = v1247 + v1253
		v1259 = base.I32_div_s(v1254, int32(4))
		v1262 = base.I32_div_s(v1254, int32(-100))
		v1265 = base.I32_div_s(v1254, int32(400))
		if int32(2) < v1250 {
			v1269 = int32(1)
		} else {
			v1269 = int32(13)
		}
		v1274 = base.I32_div_s((v1269+v1250)*int32(_a_F_EncodeDateTime_6), int32(256))
		v1278 = int32(7)
		v1279 = base.I32_rem_s(v1246+v1254*int32(365)+v1259+v1262+v1265+v1274-int32(_a_F_EncodeDateTime_7), v1278)
		if v1279 < int32(0) {
			v1284 = v1279 + v1278
		} else {
			v1284 = v1279
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1284
		v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1284<<(uint(int32(2))%32))+uint32(_c_F_EncodeDateTime[4])))
		v1289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1288))))
		v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288)+2)))
		v1291 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)) = uint8(v1291)
		*(*uint8)(unsafe.Add(mBase, uint32(l6)+2)) = uint8(v1290)
		*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v1289)
		v1296 = l6 + int32(4)
		v1298 = *(*int32)(unsafe.Add(mBase, _c_F_EncodeDateTime[3]))
		if v1298 == int32(1) {
			v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v1302 = int32(2)
			if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1301)) == int32(0) {
				v1313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1301<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v1296))) = uint16(v1313)
				v1328 = l6 + int32(6)
			} else {
				v1317 = F_pg_ultoa_n(m, v1301, v1296)
				mBase = m.M
				if v1302 <= v1317 {
					v1328 = v1296 + v1317
				} else {
					v1320 = l6 + int32(6)
					if v1317 != 0 {
						base.MemoryCopy(m, v1320-v1317, v1296, v1317)
					} else {
					}
					v1323 = v1302 - v1317
					if v1323 != 0 {
						base.MemoryFill(m, v1296, int32(48), v1323)
					} else {
					}
					v1328 = v1320
				}
			}
			v1329 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v1328))) = uint8(v1329)
			v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1331<<(uint(int32(2))%32))+uint32(_c_F_EncodeDateTime[5])))
			v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336)+2)))
			*(*uint8)(unsafe.Add(mBase, uint32(v1328)+3)) = uint8(v1337)
			v1339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1336))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1328)+1)) = uint16(v1339)
			v1388 = v1328 + int32(4)
		} else {
			v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v1344 = int32(2)
			v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1343<<(uint(v1344)%32))+uint32(_c_F_EncodeDateTime[5])))
			v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1348)+2)))
			*(*uint8)(unsafe.Add(mBase, uint32(v1296)+2)) = uint8(v1349)
			v1351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1348))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1296))) = uint16(v1351)
			v1353 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(l6)+7)) = uint8(v1353)
			v1356 = l6 + int32(8)
			v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1357)) == int32(0) {
				v1369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v1356))) = uint16(v1369)
				v1384 = l6 + int32(10)
			} else {
				v1373 = F_pg_ultoa_n(m, v1357, v1356)
				mBase = m.M
				if v1344 <= v1373 {
					v1384 = v1356 + v1373
				} else {
					v1376 = l6 + int32(10)
					if v1373 != 0 {
						base.MemoryCopy(m, v1376-v1373, v1356, v1373)
					} else {
					}
					v1379 = v1344 - v1373
					if v1379 != 0 {
						base.MemoryFill(m, v1356, int32(48), v1379)
					} else {
					}
					v1384 = v1376
				}
			}
			v1388 = v1384
		}
		v1389 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v1388))) = uint8(v1389)
		v1392 = v1388 + int32(1)
		v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v1394 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1393)) == int32(0) {
			v1405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1393<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1392))) = uint16(v1405)
			v1420 = v1388 + int32(3)
		} else {
			v1409 = F_pg_ultoa_n(m, v1393, v1392)
			mBase = m.M
			if v1394 <= v1409 {
				v1420 = v1392 + v1409
			} else {
				v1412 = v1388 + int32(3)
				if v1409 != 0 {
					base.MemoryCopy(m, v1412-v1409, v1392, v1409)
				} else {
				}
				v1415 = v1394 - v1409
				if v1415 != 0 {
					base.MemoryFill(m, v1392, int32(48), v1415)
				} else {
				}
				v1420 = v1412
			}
		}
		v1421 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v1420))) = uint8(v1421)
		v1424 = v1420 + int32(1)
		v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v1426 = int32(2)
		if int32(0)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1425)) == int32(0) {
			v1437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1425<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1424))) = uint16(v1437)
			v1452 = v1420 + int32(3)
		} else {
			v1441 = F_pg_ultoa_n(m, v1425, v1424)
			mBase = m.M
			if v1426 <= v1441 {
				v1452 = v1424 + v1441
			} else {
				v1444 = v1420 + int32(3)
				if v1441 != 0 {
					base.MemoryCopy(m, v1444-v1441, v1424, v1441)
				} else {
				}
				v1447 = v1426 - v1441
				if v1447 != 0 {
					base.MemoryFill(m, v1424, int32(48), v1447)
				} else {
				}
				v1452 = v1444
			}
		}
		v1453 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v1452))) = uint8(v1453)
		v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v1462 = v1457 >> (uint(int32(31)) % 32)
		v1466 = F_pg_ultostr_zeropad(m, v1452+int32(1), v1457^v1462-v1462, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v1573 = v1466
		} else {
			v1471 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v1466))) = uint8(v1471)
			v1474 = l1 >> (uint(int32(31)) % 32)
			v1476 = l1 ^ v1474 - v1474
			v1478 = base.I32_div_s(v1476, int32(10))
			v1481 = v1478*int32(-10) + v1476
			if v1481 != 0 {
				v1483 = v1481 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+6)) = uint8(v1483)
				v1489 = v1466 + int32(7)
			} else {
				v1489 = v1466 + int32(6)
			}
			v1491 = base.I32_div_s(v1476, int32(100))
			v1494 = v1491*int32(-10) + v1478
			v1495 = v1481 | v1494
			if v1495 == int32(0) {
				v1503 = v1466 + int32(5)
			} else {
				v1501 = v1494 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+5)) = uint8(v1501)
				v1503 = v1489
			}
			v1505 = base.I32_div_s(v1476, int32(1000))
			v1508 = v1491 + v1505*int32(-10)
			v1509 = v1495 | v1508
			if v1509 == int32(0) {
				v1517 = v1466 + int32(4)
			} else {
				v1515 = v1508 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+4)) = uint8(v1515)
				v1517 = v1503
			}
			v1519 = base.I32_div_s(v1476, int32(_a_F_EncodeDateTime_0))
			v1522 = v1505 + v1519*int32(-10)
			v1523 = v1509 | v1522
			if v1523 == int32(0) {
				v1531 = v1466 + int32(3)
			} else {
				v1529 = v1522 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+3)) = uint8(v1529)
				v1531 = v1517
			}
			v1533 = base.I32_div_s(v1476, int32(_a_F_EncodeDateTime_1))
			v1536 = v1519 + v1533*int32(-10)
			v1537 = v1523 | v1536
			if v1537 == int32(0) {
				v1545 = v1466 + int32(2)
			} else {
				v1543 = v1536 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+2)) = uint8(v1543)
				v1545 = v1531
			}
			v1547 = base.I32_div_s(v1476, int32(_a_F_EncodeDateTime_2))
			v1550 = v1547*int32(-10) + v1533
			if v1537|v1550 == int32(0) {
				v1559 = v1466 + int32(1)
			} else {
				v1557 = v1550 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1466)+1)) = uint8(v1557)
				v1559 = v1545
			}
			if base.Ui32(int32(19)) <= base.Ui32(v1533+int32(9)) {
				v1566 = F_pg_ultostr(m, v1466+int32(1), v1476)
				mBase = m.M
				v1567 = v1566
			} else {
				v1567 = v1559
			}
			v1573 = v1567
		}
		v1574 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v1573))) = uint8(v1574)
		v1576 = int32(1)
		v1577 = v1573 + v1576
		v1578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v1578 {
			v1583 = v1578
		} else {
			v1583 = v1576 - v1578
		}
		v1584 = int32(4)
		if int32(1)|base.B2i32(base.Ui32(int32(99)) < base.Ui32(v1583)) == int32(0) {
			v1595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1583<<(uint(int32(1))%32))+uint32(_c_F_EncodeDateTime[0]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1577))) = uint16(v1595)
			v1610 = v1573 + int32(3)
		} else {
			v1599 = F_pg_ultoa_n(m, v1583, v1577)
			mBase = m.M
			if v1584 <= v1599 {
				v1610 = v1577 + v1599
			} else {
				v1602 = v1573 + int32(5)
				if v1599 != 0 {
					base.MemoryCopy(m, v1602-v1599, v1577, v1599)
				} else {
				}
				v1605 = v1584 - v1599
				if v1605 != 0 {
					base.MemoryFill(m, v1577, int32(48), v1605)
				} else {
				}
				v1610 = v1602
			}
		}
		if v19 == int32(0) {
			v1684 = v1610
			v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1685 <= int32(0) {
				v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
				v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
				v1696 = v1684 + int32(3)
			} else {
				v1696 = v1684
			}
			v1697 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(10)
				v1617 = F_pg_sprintf(m, v1610, int32(_a_F_EncodeDateTime_3), v14)
				mBase = m.M
				v1618 = m.ExcPending
				if v1618 != 0 {
					return
				} else {
					v1619 = F_strlen(m, v1610)
					mBase = m.M
					v1684 = v1619 + v1610
					v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1685 <= int32(0) {
						v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
						v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
						v1696 = v1684 + int32(3)
					} else {
						v1696 = v1684
					}
					v1697 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				v1621 = int32(32)
				*(*uint8)(unsafe.Add(mBase, uint32(v1610))) = uint8(v1621)
				if l3 <= int32(0) {
					v1633 = int32(43)
				} else {
					v1633 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v1610+int32(1)))) = uint8(v1633)
				v1636 = l3 >> (uint(int32(31)) % 32)
				v1638 = l3 ^ v1636 - v1636
				v1640 = base.I32_div_s(v1638, int32(3600))
				v1641 = int32(-60)
				v1644 = base.I32_div_s(v1638, int32(60))
				v1645 = v1640*v1641 + v1644
				v1647 = v1610 + int32(2)
				v1650 = v1644*v1641 + v1638
				if v1650 != 0 {
					v1651 = int32(2)
					v1652 = F_pg_ultostr_zeropad(m, v1647, v1640, v1651)
					mBase = m.M
					v1653 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1652))) = uint8(v1653)
					v1658 = F_pg_ultostr_zeropad(m, v1652+int32(1), v1645, v1651)
					mBase = m.M
					v1665 = v1658
					v1666 = v1650
					v1667 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1665))) = uint8(v1667)
					v1672 = F_pg_ultostr_zeropad(m, v1665+int32(1), v1666, int32(2))
					mBase = m.M
					v1673 = v1672
				} else {
					v1660 = F_pg_ultostr_zeropad(m, v1647, v1640, int32(2))
					mBase = m.M
					if l5 == int32(4) {
						v1665 = v1660
						v1666 = v1645
						v1667 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v1665))) = uint8(v1667)
						v1672 = F_pg_ultostr_zeropad(m, v1665+int32(1), v1666, int32(2))
						mBase = m.M
						v1673 = v1672
					} else {
						if v1645 == int32(0) {
							v1673 = v1660
						} else {
							v1665 = v1660
							v1666 = v1645
							v1667 = int32(58)
							*(*uint8)(unsafe.Add(mBase, uint32(v1665))) = uint8(v1667)
							v1672 = F_pg_ultostr_zeropad(m, v1665+int32(1), v1666, int32(2))
							mBase = m.M
							v1673 = v1672
						}
					}
				}
				v1684 = v1673
				v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1685 <= int32(0) {
					v1689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EncodeDateTime[1])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1684)+2)) = uint8(v1689)
					v1692 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_EncodeDateTime[2])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1684))) = uint16(v1692)
					v1696 = v1684 + int32(3)
				} else {
					v1696 = v1684
				}
				v1697 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1696))) = uint8(v1697)
				m.G0 = v14 + int32(48)
				return
			}
		}
	}
}
func F_date_gt_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_date_gt_timestamptz_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_date_gt_timestamptz[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(int32(0) < v65)
}
func F_date_le_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
			}
		}
	}
}
func F_date_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2)) <= base.Ui32(v3-int32(2147483647)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if base.Ui32(int32(1)) < base.Ui32(v8-int32(2147483647)) {
			return v3 - v8
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_date_mi_0), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_date_mi_1), int32(560), int32(_a_F_date_mi_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_date_mi_0), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_date_mi_1), int32(560), int32(_a_F_date_mi_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_date_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	v3 = m.G0
	v5 = v3 - int32(176)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v7-int32(2147483647)) <= base.Ui32(int32(1)) {
		if v7 == int32(-2147483648) {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_date_out[0])))
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+8)) = uint16(v15)
			v18 = *(*int64)(unsafe.Add(mBase, _c_F_date_out[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v18
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_date_out[2])))
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v21)
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_date_out[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v24
		}
	} else {
		v37 = v7 + int32(_a_F_date_out_0)
		v38 = int32(_a_F_date_out_1)
		v39 = base.I32_div_u_s(v37, v38)
		v40 = int32(3)
		v46 = int32(2)
		v51 = base.I32_div_u_s((v39*int32(1073595727)+v37)<<(uint(v46)%32)|v40, v38)
		v54 = v7 + int32(_a_F_date_out_2) + v39*v40 + v51 + int32(_a_F_date_out_3)
		v55 = int32(1461)
		v56 = base.I32_div_u_s(v54, v55)
		v59 = v56*int32(-1461) + v54
		v61 = v59 << (uint(v46) % 32)
		if base.Ui32(v55) <= base.Ui32(v61) {
			v67 = base.I32_rem_u_s(v59+int32(305), int32(365))
			v72 = v67
		} else {
			v71 = base.I32_rem_u_s(v59+int32(306), int32(366))
			v72 = v71
		}
		v74 = base.I32_div_u_s(v61, int32(1461))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(152)))) = v74 + v56<<(uint(int32(2))%32) - int32(_a_F_date_out_4)
		v82 = v72 + int32(123)
		v86 = int32(base.Ui32(v82*int32(2141)) >> (uint(int32(16)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(144)))) = v82 - int32(base.Ui32(v86*int32(_a_F_date_out_5))>>(uint(int32(8))%32))
		v96 = base.I32_rem_u_s(v86+int32(10), int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(148)))) = v96 + int32(1)
		v101 = v5 + int32(132)
		v103 = *(*int32)(unsafe.Add(mBase, _c_F_date_out[4]))
		switch v103 - int32(1) {
		case 0, 3:
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v106 {
				v111 = v106
			} else {
				v111 = int32(1) - v106
			}
			v113 = F_pg_ultostr_zeropad(m, v5, v111, int32(4))
			mBase = m.M
			v114 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v114)
			v116 = int32(1)
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
			v119 = int32(2)
			v120 = F_pg_ultostr_zeropad(m, v113+v116, v118, v119)
			mBase = m.M
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v114)
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
			v127 = F_pg_ultostr_zeropad(m, v120+v116, v125, v119)
			mBase = m.M
			v220 = v127
		case 1:
			v131 = *(*int32)(unsafe.Add(mBase, _c_F_date_out[5]))
			v133 = base.B2i32(v131 == int32(1))
			if v131 == int32(1) {
				v134 = int32(12)
			} else {
				v134 = int32(16)
			}
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v101+v134)))
			v138 = F_pg_ultostr_zeropad(m, v5, v136, int32(2))
			mBase = m.M
			v139 = int32(47)
			*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v139)
			if v131 == int32(1) {
				v145 = int32(16)
			} else {
				v145 = int32(12)
			}
			v147 = *(*int32)(unsafe.Add(mBase, uint32(v101+v145)))
			v149 = F_pg_ultostr_zeropad(m, v138+int32(1), v147, int32(2))
			mBase = m.M
			v150 = int32(47)
			*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v150)
			v152 = int32(1)
			v154 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v154 {
				v159 = v154
			} else {
				v159 = v152 - v154
			}
			v161 = F_pg_ultostr_zeropad(m, v149+v152, v159, int32(4))
			mBase = m.M
			v220 = v161
		case 2:
			v162 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
			v163 = int32(2)
			v164 = F_pg_ultostr_zeropad(m, v5, v162, v163)
			mBase = m.M
			v165 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
			v167 = int32(1)
			v169 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
			v171 = F_pg_ultostr_zeropad(m, v164+v167, v169, v163)
			mBase = m.M
			*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v165)
			v176 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v176 {
				v181 = v176
			} else {
				v181 = v167 - v176
			}
			v183 = F_pg_ultostr_zeropad(m, v171+v167, v181, int32(4))
			mBase = m.M
			v220 = v183
		default:
			v187 = *(*int32)(unsafe.Add(mBase, _c_F_date_out[5]))
			v189 = base.B2i32(v187 == int32(1))
			if v187 == int32(1) {
				v190 = int32(12)
			} else {
				v190 = int32(16)
			}
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v101+v190)))
			v194 = F_pg_ultostr_zeropad(m, v5, v192, int32(2))
			mBase = m.M
			v195 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v195)
			if v187 == int32(1) {
				v201 = int32(16)
			} else {
				v201 = int32(12)
			}
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v101+v201)))
			v205 = F_pg_ultostr_zeropad(m, v194+int32(1), v203, int32(2))
			mBase = m.M
			v206 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v206)
			v208 = int32(1)
			v210 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v210 {
				v215 = v210
			} else {
				v215 = v208 - v210
			}
			v217 = F_pg_ultostr_zeropad(m, v205+v208, v215, int32(4))
			mBase = m.M
			v220 = v217
		}
		v221 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
		if v221 <= int32(0) {
			v225 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_date_out[6])))
			*(*uint8)(unsafe.Add(mBase, uint32(v220)+2)) = uint8(v225)
			v228 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_date_out[7])))
			*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v228)
			v232 = v220 + int32(3)
		} else {
			v232 = v220
		}
		v233 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v233)
	}
	v235 = F_pstrdup(m, v5)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(176)
		return v235
	}
}
func F_date_pl_interval(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13889(m, l0, int32(1267))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_date_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(1264)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(1265)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9223372034707292160)
	return int32(0)
}
func F_date_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_date_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(-2147483648) {
		v6 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		if v2 == int32(2147483647) {
			v14 = F_Int64GetDatum(m, int64(9223372036854775807))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		} else {
			if v2 < int32(106751983) {
				v22 = F_Int64GetDatum(m, base.I64_extend_i32_s(v2)*int64(86400000000))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v22
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_date_timestamp_0), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_date_timestamp_1), int32(658), int32(_a_F_date_timestamp_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_date_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v133 int64
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if v4 == int32(-2147483648) {
		v155 = int64(-9223372036854775807 - 1)
		m.G0 = v7 + int32(48)
		v159 = F_Int64GetDatum(m, v155)
		mBase = m.M
		v160 = m.ExcPending
		if v160 != 0 {
			return int32(0)
		} else {
			return v159
		}
	} else {
		if v4 != int32(2147483647) {
			if int32(106751983) <= v4 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_date_timestamptz_0), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_date_timestamptz_1), int32(721), int32(_a_F_date_timestamptz_2))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v45 = v4 + int32(_a_F_date_timestamptz_3)
				v46 = int32(_a_F_date_timestamptz_4)
				v47 = base.I32_div_u_s(v45, v46)
				v48 = int32(3)
				v54 = int32(2)
				v59 = base.I32_div_u_s((v47*int32(1073595727)+v45)<<(uint(v54)%32)|v48, v46)
				v62 = v4 + int32(_a_F_date_timestamptz_5) + v47*v48 + v59 + int32(_a_F_date_timestamptz_6)
				v63 = int32(1461)
				v64 = base.I32_div_u_s(v62, v63)
				v67 = v64*int32(-1461) + v62
				v69 = v67 << (uint(v54) % 32)
				if base.Ui32(v63) <= base.Ui32(v69) {
					v75 = base.I32_rem_u_s(v67+int32(305), int32(365))
					v80 = v75
				} else {
					v79 = base.I32_rem_u_s(v67+int32(306), int32(366))
					v80 = v79
				}
				v82 = base.I32_div_u_s(v69, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v82 + v64<<(uint(int32(2))%32) - int32(_a_F_date_timestamptz_7)
				v90 = v80 + int32(123)
				v94 = int32(base.Ui32(v90*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v90 - int32(base.Ui32(v94*int32(_a_F_date_timestamptz_8))>>(uint(int32(8))%32))
				v104 = base.I32_rem_u_s(v94+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v104 + int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v115 = *(*int32)(unsafe.Add(mBase, _c_F_date_timestamptz[0]))
				v117 = m.G0
				v118 = int32(16)
				v119 = v117 - v118
				m.G0 = v119
				v123 = F_DetermineTimeZoneOffsetInternal(m, v7+int32(4), v115, v119+int32(8))
				mBase = m.M
				m.G0 = v119 + v118
				v133 = base.I64_extend_i32_s(v123)*int64(1000000) + base.I64_extend_i32_s(v4)*int64(86400000000)
				if base.Ui64(v133+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v155 = v133
					m.G0 = v7 + int32(48)
					v159 = F_Int64GetDatum(m, v155)
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						return v159
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_date_timestamptz_0), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_date_timestamptz_1), int32(757), int32(_a_F_date_timestamptz_2))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			v155 = int64(9223372036854775807)
			m.G0 = v7 + int32(48)
			v159 = F_Int64GetDatum(m, v155)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				return v159
			}
		}
	}
}
