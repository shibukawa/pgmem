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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
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
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int64
	_ = v821
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v897 int32
	_ = v897
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v961 float64
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 float64
	_ = v967
	var v969 float64
	_ = v969
	var v973 int64
	_ = v973
	var v975 int64
	_ = v975
	var v978 int64
	_ = v978
	var v983 int64
	_ = v983
	var v985 int64
	_ = v985
	var v990 int64
	_ = v990
	var v992 int64
	_ = v992
	var v996 int64
	_ = v996
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1059 int32
	_ = v1059
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1219 int32
	_ = v1219
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1320 int32
	_ = v1320
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1437 int32
	_ = v1437
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1480 int32
	_ = v1480
	var v1500 int32
	_ = v1500
	var v1516 int32
	_ = v1516
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1552 int32
	_ = v1552
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1658 int32
	_ = v1658
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1784 int32
	_ = v1784
	var v1792 int32
	_ = v1792
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1815 int32
	_ = v1815
	var v1823 int32
	_ = v1823
	var v1832 int32
	_ = v1832
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1913 int32
	_ = v1913
	var v1939 int32
	_ = v1939
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2025 int32
	_ = v2025
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
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2070 int32
	_ = v2070
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2162 int32
	_ = v2162
	var v2172 int32
	_ = v2172
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2203 int32
	_ = v2203
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2227 int32
	_ = v2227
	var v2230 int32
	_ = v2230
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2246 int32
	_ = v2246
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2300 int32
	_ = v2300
	var v2308 int32
	_ = v2308
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
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
	m.G0 = v2329 + int32(112)
	return v2328
L5:
	;
	v62 = l4 + int32(20)
	v64 = l4 + int32(12)
	v66 = l4 + int32(16)
	v78 = v9
	v79 = v9
	v86 = v42
	v88 = v9
	v89 = v9
	v93 = v9
	v94 = v9
	v95 = v9
	v98 = v9
	v99 = v9
	goto L8
L6:
	;
	v2015 = v38
	v2018 = v9
	v2025 = v42
	v2028 = v9
	v2032 = v9
	v2034 = v9
	v2037 = v9
	v2038 = v9
	goto L7
L7:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v2041 != int32(2) {
		goto L449
	} else {
		goto L450
	}
L8:
	;
	v102 = int32(-1)
	v104 = v88 << (uint(int32(2)) % 32)
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
		v2328 = v102
		v2329 = v38
		goto L4
	}
L9:
	;
	if v1978 != 0 {
		v2328 = int32(-1)
		v2329 = v38
		goto L4
	} else {
		goto L448
	}
L10:
	;
	v2003 = v88 + int32(1)
	if v2003 != l2 {
		v78 = v1978
		v79 = v1979
		v86 = v1986
		v88 = v2003
		v89 = v1989
		v93 = v1993
		v94 = v1994
		v95 = v1995
		v98 = v1998
		v99 = v1999
		goto L8
	} else {
		goto L447
	}
L11:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v1963&v79 != 0 {
		goto L444
	} else {
		goto L445
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1939 = v78
	v1947 = v86
	v1950 = v1913
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L13:
	;
	v1256 = l0 + v104
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	v1264 = F_DecodeTimezoneAbbrev(m, v88, v1257, v38-int32(-64), v38+int32(60), v38+int32(52), l7)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L122
	} else {
		goto L329
	}
L14:
	;
	if v78 != 0 {
		goto L176
	} else {
		goto L177
	}
L15:
	;
	if l6 == int32(0) {
		v2328 = v102
		v2329 = v38
		goto L4
	} else {
		goto L148
	}
L16:
	;
	switch v78 {
	case 0, 3:
		goto L133
	default:
		v2328 = v102
		v2329 = v38
		goto L4
	}
L17:
	;
	if v78 == int32(31) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if l6 == int32(0) {
		v2328 = v102
		v2329 = v38
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v301 = int32(0)
	v303 = int32(10)
	if base.B2i32(v78 == v301)&base.B2i32(v79&v303 != v303) == v301 {
		goto L56
	} else {
		goto L57
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v119 = F_strtol(m, v115, v38+int32(72), int32(10))
	mBase = m.M
	goto L22
L22:
	;
	v120 = int32(-2)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v122 == int32(68) {
		v2328 = v120
		v2329 = v38
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v119 < int32(0) {
		v2328 = v120
		v2329 = v38
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v128 = v119 + int32(32044)
	v129 = int32(146097)
	v130 = base.I32_div_u_s(v128, v129)
	v131 = int32(3)
	v137 = int32(2)
	v142 = base.I32_div_u_s((v130*int32(1073595727)+v128)<<(uint(v137)%32)|v131, v129)
	v145 = v119 + v130*v131 + v142 + int32(32104)
	v146 = int32(1461)
	v147 = base.I32_div_u_s(v145, v146)
	v150 = v147*int32(-1461) + v145
	v152 = v150 << (uint(v137) % 32)
	if base.Ui32(v146) <= base.Ui32(v152) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v165 = base.I32_div_u_s(v152, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v165 + v147<<(uint(int32(2))%32) - int32(4800)
	v173 = v163 + int32(123)
	v176 = int32(16)
	v177 = int32(base.Ui32(v173*int32(2141)) >> (uint(v176) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v173 - int32(base.Ui32(v177*int32(7834))>>(uint(int32(8))%32))
	v187 = base.I32_rem_u_s(v177+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v187 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v192 = int32(0)
	v198 = m.G0
	v200 = v198 - v176
	m.G0 = v200
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	switch v203 - int32(43) {
	case 0, 2:
		goto L31
	default:
		v290 = int32(-1)
		goto L30
	}
L26:
	;
	v158 = base.I32_rem_u_s(v150+int32(305), int32(365))
	v163 = v158
	goto L25
L27:
	;
	goto L28
L28:
	;
	v162 = base.I32_rem_u_s(v150+int32(306), int32(366))
	v163 = v162
	goto L25
L29:
	;
	if v290 != 0 {
		v2328 = v290
		v2329 = v38
		goto L4
	} else {
		goto L55
	}
L30:
	;
	m.G0 = v200 + int32(16)
	goto L29
L31:
	;
	v206 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v213 = F_strtoint(m, v191+int32(1), v200+int32(12))
	mBase = m.M
	v214 = int32(-5)
	v216 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v216 == int32(68) {
		v290 = v214
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v220 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if base.Ui32(int32(15)) < base.Ui32(v262) {
		v290 = v214
		goto L30
	} else {
		goto L46
	}
L34:
	;
	if v220 != int32(58) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v252 = F_strlen(m, v191)
	mBase = m.M
	if base.Ui32(v252) < base.Ui32(int32(4)) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v261 = int32(0)
	v262 = v213
	v264 = v192
	goto L33
L38:
	;
	goto L39
L39:
	;
	v224 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v231 = F_strtoint(m, v219+int32(1), v200+int32(12))
	mBase = m.M
	v233 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v233 == int32(68) {
		v290 = v214
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v237 != int32(58) {
		v261 = v231
		v262 = v213
		v264 = v192
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v240 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v247 = F_strtoint(m, v236+int32(1), v200+int32(12))
	mBase = m.M
	v249 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v249 != int32(68) {
		v261 = v231
		v262 = v213
		v264 = v247
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v290 = v214
	goto L30
L43:
	;
	v261 = int32(0)
	v262 = v213
	v264 = v192
	goto L33
L44:
	;
	goto L45
L45:
	;
	v256 = int32(100)
	v257 = base.I32_div_s(v213, v256)
	v261 = v213 - v257*v256
	v262 = v257
	v264 = v192
	goto L33
L46:
	;
	if base.Ui32(int32(59)) < base.Ui32(v261) {
		v290 = v214
		goto L30
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(59)) < base.Ui32(v264) {
		v290 = v214
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v271 = int32(60)
	v276 = (v262*v271+v261)*v271 + v264
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v279 == int32(45) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v282 = v276
	goto L51
L50:
	;
	v282 = int32(0) - v276
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v282
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v287 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v288 = int32(-1)
	goto L54
L53:
	;
	v288 = int32(0)
	goto L54
L54:
	;
	v290 = v288
	goto L30
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = int32(1)
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L56:
	;
	if l6 == int32(0) {
		v2328 = v102
		v2329 = v38
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v542 = F_DecodeDate(m, v537, v79, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L122
	} else {
		goto L131
	}
L59:
	;
	v312 = l0 + v104
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v78 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v526 = F_pg_tzset(m, v313)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L122
	} else {
		goto L127
	}
L61:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if base.Ui32(int32(9)) < base.Ui32((v316-int32(48))&int32(255)) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if base.B2i32(v78 != int32(3))&base.B2i32(v78 != int32(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L66:
	;
	goto L67
L67:
	;
	v329 = int32(31744)
	if v79&v329 == v329 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L69:
	;
	goto L70
L70:
	;
	v334 = int32(45)
	v335 = F___strchrnul(m, v313, v334)
	mBase = m.M
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v337 == v334 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v341 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v341 = v335
	goto L74
L73:
	;
	v341 = int32(0)
	goto L74
L74:
	;
	goto L71
L75:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L76:
	;
	goto L77
L77:
	;
	v345 = int32(0)
	v351 = m.G0
	v353 = v351 - int32(16)
	m.G0 = v353
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	switch v356 - int32(43) {
	case 0, 2:
		goto L80
	default:
		v443 = int32(-1)
		goto L79
	}
L78:
	;
	if v443 != 0 {
		v2328 = v443
		v2329 = v38
		goto L4
	} else {
		goto L104
	}
L79:
	;
	m.G0 = v353 + int32(16)
	goto L78
L80:
	;
	v359 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v366 = F_strtoint(m, v341+int32(1), v353+int32(12))
	mBase = m.M
	v367 = int32(-5)
	v369 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v369 == int32(68) {
		v443 = v367
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v373 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v415) {
		v443 = v367
		goto L79
	} else {
		goto L95
	}
L83:
	;
	if v373 != int32(58) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v405 = F_strlen(m, v341)
	mBase = m.M
	if base.Ui32(v405) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v414 = int32(0)
	v415 = v366
	v417 = v345
	goto L82
L87:
	;
	goto L88
L88:
	;
	v377 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v384 = F_strtoint(m, v372+int32(1), v353+int32(12))
	mBase = m.M
	v386 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v386 == int32(68) {
		v443 = v367
		goto L79
	} else {
		goto L89
	}
L89:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v390 != int32(58) {
		v414 = v384
		v415 = v366
		v417 = v345
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v393 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v400 = F_strtoint(m, v389+int32(1), v353+int32(12))
	mBase = m.M
	v402 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v402 != int32(68) {
		v414 = v384
		v415 = v366
		v417 = v400
		goto L82
	} else {
		goto L91
	}
L91:
	;
	v443 = v367
	goto L79
L92:
	;
	v414 = int32(0)
	v415 = v366
	v417 = v345
	goto L82
L93:
	;
	goto L94
L94:
	;
	v409 = int32(100)
	v410 = base.I32_div_s(v366, v409)
	v414 = v366 - v410*v409
	v415 = v410
	v417 = v345
	goto L82
L95:
	;
	if base.Ui32(int32(59)) < base.Ui32(v414) {
		v443 = v367
		goto L79
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(int32(59)) < base.Ui32(v417) {
		v443 = v367
		goto L79
	} else {
		goto L97
	}
L97:
	;
	v424 = int32(60)
	v429 = (v415*v424+v414)*v424 + v417
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v432 == int32(45) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v435 = v429
	goto L100
L99:
	;
	v435 = int32(0) - v429
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v435
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	if v440 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v441 = int32(-1)
	goto L103
L102:
	;
	v441 = int32(0)
	goto L103
L103:
	;
	v443 = v441
	goto L79
L104:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v450)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v452&int32(3) == v450 {
		v476 = v452
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v514 = F_DecodeNumberField(m, v509, v452, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L122
	} else {
		goto L123
	}
L106:
	;
	v509 = v501 - v452
	goto L105
L107:
	;
	v480 = v476
	goto L116
L108:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v460 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v509 = int32(0)
	goto L105
L110:
	;
	goto L111
L111:
	;
	v465 = v452
	goto L112
L112:
	;
	v469 = v465 + int32(1)
	if v469&int32(3) == int32(0) {
		v476 = v469
		goto L107
	} else {
		goto L114
	}
L113:
	;
	v501 = v469
	goto L106
L114:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	if v474 != 0 {
		v465 = v469
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v489 = int32(-2139062144)
	if (int32(16843008)-v486|v486)&v489 == v489 {
		v480 = v480 + int32(4)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v495 = v480
	goto L119
L118:
	;
	goto L117
L119:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	if v499 != 0 {
		v495 = v495 + int32(1)
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v501 = v495
	goto L106
L121:
	;
	goto L120
L122:
	;
	return int32(0)
L123:
	;
	if v514 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L125:
	;
	goto L126
L126:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v521 | int32(32)
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L127:
	;
	if v526 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v530
	v2328 = int32(-6)
	v2329 = v38
	goto L4
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1939 = int32(0)
	v1947 = v86
	v1950 = v526
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L131:
	;
	if v542 != 0 {
		v2328 = v542
		v2329 = v38
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L133:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v552 = F_DecodeTimeCommon(m, v546, int32(32767), v38+int32(68), v38+int32(72))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L122
	} else {
		goto L134
	}
L134:
	;
	if v552 != 0 {
		v2328 = v552
		v2329 = v38
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v38)+88))
	if int64(2147483648) <= v554 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v2328 = int32(-2)
	v2329 = v38
	goto L4
L137:
	;
	goto L138
L138:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v554)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v559
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v563
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v569 = int32(1)
	if base.Ui32(int32(24)) < base.Ui32(v565) {
		v591 = v569
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v591 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	goto L139
L141:
	;
	if base.Ui32(int32(59)) < base.Ui32(v566) {
		v591 = v569
		goto L140
	} else {
		goto L142
	}
L142:
	;
	if base.Ui32(int32(60)) < base.Ui32(v567) {
		v591 = v569
		goto L140
	} else {
		goto L143
	}
L143:
	;
	if base.Ui32(int32(1000000)) < base.Ui32(v563) {
		v591 = v569
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v579 = int32(60)
	v591 = base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v563)+base.I64_extend_i32_u((v565*v579+v566)*v579+v567)*int64(1000000)))
	goto L140
L145:
	;
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L146:
	;
	goto L147
L147:
	;
	v2328 = int32(-2)
	v2329 = v38
	goto L4
L148:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v602 = int32(0)
	v608 = m.G0
	v610 = v608 - int32(16)
	m.G0 = v610
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	switch v613 - int32(43) {
	case 0, 2:
		goto L151
	default:
		v700 = int32(-1)
		goto L150
	}
L149:
	;
	if v700 != 0 {
		v2328 = v700
		v2329 = v38
		goto L4
	} else {
		goto L175
	}
L150:
	;
	m.G0 = v610 + int32(16)
	goto L149
L151:
	;
	v616 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v623 = F_strtoint(m, v599+int32(1), v610+int32(12))
	mBase = m.M
	v624 = int32(-5)
	v626 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v626 == int32(68) {
		v700 = v624
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	if v630 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if base.Ui32(int32(15)) < base.Ui32(v672) {
		v700 = v624
		goto L150
	} else {
		goto L166
	}
L154:
	;
	if v630 != int32(58) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v662 = F_strlen(m, v599)
	mBase = m.M
	if base.Ui32(v662) < base.Ui32(int32(4)) {
		goto L163
	} else {
		goto L164
	}
L157:
	;
	v671 = int32(0)
	v672 = v623
	v674 = v602
	goto L153
L158:
	;
	goto L159
L159:
	;
	v634 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v641 = F_strtoint(m, v629+int32(1), v610+int32(12))
	mBase = m.M
	v643 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v643 == int32(68) {
		v700 = v624
		goto L150
	} else {
		goto L160
	}
L160:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	if v647 != int32(58) {
		v671 = v641
		v672 = v623
		v674 = v602
		goto L153
	} else {
		goto L161
	}
L161:
	;
	v650 = int32(4613548)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v657 = F_strtoint(m, v646+int32(1), v610+int32(12))
	mBase = m.M
	v659 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v659 != int32(68) {
		v671 = v641
		v672 = v623
		v674 = v657
		goto L153
	} else {
		goto L162
	}
L162:
	;
	v700 = v624
	goto L150
L163:
	;
	v671 = int32(0)
	v672 = v623
	v674 = v602
	goto L153
L164:
	;
	goto L165
L165:
	;
	v666 = int32(100)
	v667 = base.I32_div_s(v623, v666)
	v671 = v623 - v667*v666
	v672 = v667
	v674 = v602
	goto L153
L166:
	;
	if base.Ui32(int32(59)) < base.Ui32(v671) {
		v700 = v624
		goto L150
	} else {
		goto L167
	}
L167:
	;
	if base.Ui32(int32(59)) < base.Ui32(v674) {
		v700 = v624
		goto L150
	} else {
		goto L168
	}
L168:
	;
	v681 = int32(60)
	v686 = (v672*v681+v671)*v681 + v674
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	if v689 == int32(45) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v692 = v686
	goto L171
L170:
	;
	v692 = int32(0) - v686
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(72)))) = v692
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	if v697 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v698 = int32(-1)
	goto L174
L173:
	;
	v698 = int32(0)
	goto L174
L174:
	;
	v700 = v698
	goto L150
L175:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v707
	v1913 = v89
	goto L12
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v712 = l0 + v104
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	v717 = F_strtol(m, v713, v38+int32(4), int32(10))
	mBase = m.M
	goto L179
L177:
	;
	goto L178
L178:
	;
	v1084 = v79 & int32(14)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	if v1086&int32(3) == int32(0) {
		v1110 = v1086
		goto L273
	} else {
		goto L274
	}
L179:
	;
	v718 = int32(-2)
	v720 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v720 == int32(68) {
		v2328 = v718
		v2329 = v38
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723))))
	if v724 == int32(46) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v78 != int32(3) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	if v724 == int32(0) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v1078
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L185:
	;
	if v78 != int32(31) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L187
L187:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	if v1002&int32(3) == int32(0) {
		v1026 = v1002
		goto L251
	} else {
		goto L252
	}
L188:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L189:
	;
	goto L190
L190:
	;
	if v717 < int32(0) {
		v2328 = v718
		v2329 = v38
		goto L4
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	v740 = v717 + int32(32044)
	v741 = int32(146097)
	v742 = base.I32_div_u_s(v740, v741)
	v743 = int32(3)
	v749 = int32(2)
	v754 = base.I32_div_u_s((v742*int32(1073595727)+v740)<<(uint(v749)%32)|v743, v741)
	v757 = v717 + v742*v743 + v754 + int32(32104)
	v758 = int32(1461)
	v759 = base.I32_div_u_s(v757, v758)
	v762 = v759*int32(-1461) + v757
	v764 = v762 << (uint(v749) % 32)
	if base.Ui32(v758) <= base.Ui32(v764) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v777 = base.I32_div_u_s(v764, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v777 + v759<<(uint(int32(2))%32) - int32(4800)
	v785 = v775 + int32(123)
	v789 = int32(base.Ui32(v785*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v785 - int32(base.Ui32(v789*int32(7834))>>(uint(int32(8))%32))
	v796 = int32(1)
	v800 = base.I32_rem_u_s(v789+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v800 + v796
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723))))
	if v804 != int32(46) {
		v1078 = v796
		goto L184
	} else {
		goto L196
	}
L193:
	;
	v770 = base.I32_rem_u_s(v762+int32(305), int32(365))
	v775 = v770
	goto L192
L194:
	;
	goto L195
L195:
	;
	v774 = base.I32_rem_u_s(v762+int32(306), int32(366))
	v775 = v774
	goto L192
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v723
	v809 = v723 + int32(1)
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	if v810 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L198:
	;
	v969 = base.F64_mul(v967, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v969), float64(9.223372036854776e+18)) != 0 {
		goto L245
	} else {
		goto L246
	}
L199:
	;
	v967 = float64(0)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v814 = int32(527523)
	v818 = m.G0
	v820 = v818 - int32(32)
	v821 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+24)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v820)+16)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v820)+8)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v820))) = v821
	v829 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v829 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v809&int32(3) == int32(0) {
		v921 = v809
		goto L225
	} else {
		goto L226
	}
L203:
	;
	v897 = int32(0)
	goto L202
L204:
	;
	goto L205
L205:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v833 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v837 = v809
	goto L209
L207:
	;
	goto L208
L208:
	;
	v847 = v814
	v848 = v829
	goto L212
L209:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if v843 == v829 {
		v837 = v837 + int32(1)
		goto L209
	} else {
		goto L211
	}
L210:
	;
	v897 = v837 - v809
	goto L202
L211:
	;
	goto L210
L212:
	;
	v855 = v820 + int32(base.Ui32(v848)>>(uint(int32(3))%32))&int32(28)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	v857 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v855))) = v856 | v857<<(uint(v848)%32)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+1)))
	if v861 != 0 {
		v847 = v847 + v857
		v848 = v861
		goto L212
	} else {
		goto L214
	}
L213:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	if v864 == int32(0) {
		v889 = v809
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L213
L215:
	;
	v897 = v889 - v809
	goto L202
L216:
	;
	v868 = v809
	v869 = v864
	goto L217
L217:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v820+int32(base.Ui32(v869)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v877)>>(uint(v869)%32))&int32(1) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v889 = v885
	goto L215
L219:
	;
	v889 = v868
	goto L215
L220:
	;
	goto L221
L221:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868)+1)))
	v885 = v868 + int32(1)
	if v883 != 0 {
		v868 = v885
		v869 = v883
		goto L217
	} else {
		goto L222
	}
L222:
	;
	goto L218
L223:
	;
	if v897 != v954 {
		goto L197
	} else {
		goto L240
	}
L224:
	;
	v954 = v946 - v809
	goto L223
L225:
	;
	v925 = v921
	goto L234
L226:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	if v905 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v954 = int32(0)
	goto L223
L228:
	;
	goto L229
L229:
	;
	v910 = v809
	goto L230
L230:
	;
	v914 = v910 + int32(1)
	if v914&int32(3) == int32(0) {
		v921 = v914
		goto L225
	} else {
		goto L232
	}
L231:
	;
	v946 = v914
	goto L224
L232:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	if v919 != 0 {
		v910 = v914
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	v934 = int32(-2139062144)
	if (int32(16843008)-v931|v931)&v934 == v934 {
		v925 = v925 + int32(4)
		goto L234
	} else {
		goto L236
	}
L235:
	;
	v940 = v925
	goto L237
L236:
	;
	goto L235
L237:
	;
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940))))
	if v944 != 0 {
		v940 = v940 + int32(1)
		goto L237
	} else {
		goto L239
	}
L238:
	;
	v946 = v940
	goto L224
L239:
	;
	goto L238
L240:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v961 = F_strtod(m, v723, v38+int32(72))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L122
	} else {
		goto L241
	}
L241:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963))))
	if v964 != 0 {
		goto L197
	} else {
		goto L242
	}
L242:
	;
	v966 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v966 != 0 {
		goto L197
	} else {
		goto L243
	}
L243:
	;
	v967 = v961
	goto L198
L244:
	;
	v978 = base.I64_div_s(v975, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v56))) = uint32(v978)
	v983 = base.I64_extend32_s(v978)*int64(-3600000000) + v975
	v985 = base.I64_div_s(v983, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+int32(4)))) = uint32(v985)
	v990 = base.I64_extend32_s(v985)*int64(-60000000) + v983
	v992 = base.I64_div_s(v990, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v992)
	v996 = v992*int64(4293967296) + v990
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v996)
	goto L248
L245:
	;
	v973 = base.I64_trunc_f64_s(v969)
	v975 = v973
	goto L244
L246:
	;
	goto L247
L247:
	;
	v975 = int64(-9223372036854775807 - 1)
	goto L244
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31758)
	v1078 = v796
	goto L184
L249:
	;
	v1066 = F_DecodeNumberField(m, v1059, v1002, v79|int32(14), v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L122
	} else {
		goto L266
	}
L250:
	;
	v1059 = v1051 - v1002
	goto L249
L251:
	;
	v1030 = v1026
	goto L260
L252:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1010 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1059 = int32(0)
	goto L249
L254:
	;
	goto L255
L255:
	;
	v1015 = v1002
	goto L256
L256:
	;
	v1019 = v1015 + int32(1)
	if v1019&int32(3) == int32(0) {
		v1026 = v1019
		goto L251
	} else {
		goto L258
	}
L257:
	;
	v1051 = v1019
	goto L250
L258:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019))))
	if v1024 != 0 {
		v1015 = v1019
		goto L256
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1030)))
	v1039 = int32(-2139062144)
	if (int32(16843008)-v1036|v1036)&v1039 == v1039 {
		v1030 = v1030 + int32(4)
		goto L260
	} else {
		goto L262
	}
L261:
	;
	v1045 = v1030
	goto L263
L262:
	;
	goto L261
L263:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
	if v1049 != 0 {
		v1045 = v1045 + int32(1)
		goto L263
	} else {
		goto L265
	}
L264:
	;
	v1051 = v1045
	goto L250
L265:
	;
	goto L264
L266:
	;
	if v1066 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L268:
	;
	goto L269
L269:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v1072 != int32(31744) {
		v2328 = int32(-1)
		v2329 = v38
		goto L4
	} else {
		goto L270
	}
L270:
	;
	v1078 = v93
	goto L184
L271:
	;
	v1144 = int32(46)
	v1145 = F___strchrnul(m, v1086, v1144)
	mBase = m.M
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145))))
	if v1147 == v1144 {
		goto L290
	} else {
		goto L291
	}
L272:
	;
	v1143 = v1135 - v1086
	goto L271
L273:
	;
	v1114 = v1110
	goto L282
L274:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086))))
	if v1094 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1143 = int32(0)
	goto L271
L276:
	;
	goto L277
L277:
	;
	v1099 = v1086
	goto L278
L278:
	;
	v1103 = v1099 + int32(1)
	if v1103&int32(3) == int32(0) {
		v1110 = v1103
		goto L273
	} else {
		goto L280
	}
L279:
	;
	v1135 = v1103
	goto L272
L280:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103))))
	if v1108 != 0 {
		v1099 = v1103
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1114)))
	v1123 = int32(-2139062144)
	if (int32(16843008)-v1120|v1120)&v1123 == v1123 {
		v1114 = v1114 + int32(4)
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v1129 = v1114
	goto L285
L284:
	;
	goto L283
L285:
	;
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	if v1133 != 0 {
		v1129 = v1129 + int32(1)
		goto L285
	} else {
		goto L287
	}
L286:
	;
	v1135 = v1129
	goto L272
L287:
	;
	goto L286
L288:
	;
	if v1151 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L289:
	;
	if v1151 == int32(0) {
		goto L288
	} else {
		goto L293
	}
L290:
	;
	v1151 = v1145
	goto L292
L291:
	;
	v1151 = int32(0)
	goto L292
L292:
	;
	goto L289
L293:
	;
	if v1084 != 0 {
		goto L288
	} else {
		goto L294
	}
L294:
	;
	v1158 = F_DecodeDate(m, v1086, v79, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L122
	} else {
		goto L295
	}
L295:
	;
	if v1158 != 0 {
		v2328 = v1158
		v2329 = v38
		goto L4
	} else {
		goto L296
	}
L296:
	;
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L297:
	;
	if v1143 < int32(6) {
		goto L319
	} else {
		goto L320
	}
L298:
	;
	if v1151&int32(3) == int32(0) {
		v1186 = v1151
		goto L301
	} else {
		goto L302
	}
L299:
	;
	if base.Ui32(v1143-v1219) < base.Ui32(int32(3)) {
		goto L297
	} else {
		goto L316
	}
L300:
	;
	v1219 = v1211 - v1151
	goto L299
L301:
	;
	v1190 = v1186
	goto L310
L302:
	;
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	if v1170 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1219 = int32(0)
	goto L299
L304:
	;
	goto L305
L305:
	;
	v1175 = v1151
	goto L306
L306:
	;
	v1179 = v1175 + int32(1)
	if v1179&int32(3) == int32(0) {
		v1186 = v1179
		goto L301
	} else {
		goto L308
	}
L307:
	;
	v1211 = v1179
	goto L300
L308:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179))))
	if v1184 != 0 {
		v1175 = v1179
		goto L306
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1199 = int32(-2139062144)
	if (int32(16843008)-v1196|v1196)&v1199 == v1199 {
		v1190 = v1190 + int32(4)
		goto L310
	} else {
		goto L312
	}
L311:
	;
	v1205 = v1190
	goto L313
L312:
	;
	goto L311
L313:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205))))
	if v1209 != 0 {
		v1205 = v1205 + int32(1)
		goto L313
	} else {
		goto L315
	}
L314:
	;
	v1211 = v1205
	goto L300
L315:
	;
	goto L314
L316:
	;
	v1228 = F_DecodeNumberField(m, v1143, v1086, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L122
	} else {
		goto L317
	}
L317:
	;
	if int32(0) <= v1228 {
		v1939 = int32(0)
		v1947 = v86
		v1950 = v89
		v1954 = v93
		v1955 = v94
		v1956 = v95
		v1959 = v98
		v1960 = v99
		goto L11
	} else {
		goto L318
	}
L318:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L319:
	;
	v1253 = F_DecodeNumber(m, v1143, v1086, v94, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L122
	} else {
		goto L327
	}
L320:
	;
	if v79&int32(31744) != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1238 = v1084
	goto L323
L322:
	;
	v1238 = int32(0)
	goto L323
L323:
	;
	if v1238 != 0 {
		goto L319
	} else {
		goto L324
	}
L324:
	;
	v1244 = F_DecodeNumberField(m, v1143, v1086, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L122
	} else {
		goto L325
	}
L325:
	;
	if int32(0) <= v1244 {
		v1939 = int32(0)
		v1947 = v86
		v1950 = v89
		v1954 = v93
		v1955 = v94
		v1956 = v95
		v1959 = v98
		v1960 = v99
		goto L11
	} else {
		goto L326
	}
L326:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L327:
	;
	if v1253 != 0 {
		v2328 = v1253
		v2329 = v38
		goto L4
	} else {
		goto L328
	}
L328:
	;
	v1939 = int32(0)
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L329:
	;
	if v1264 != 0 {
		v2328 = v1264
		v2329 = v38
		goto L4
	} else {
		goto L330
	}
L330:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v1266 == int32(31) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1068])))
	if v1272 != 0 {
		goto L336
	} else {
		goto L337
	}
L332:
	;
	v1516 = v1266
	goto L333
L333:
	;
	if v1516 == int32(8) {
		v1978 = v78
		v1979 = v79
		v1986 = v86
		v1989 = v89
		v1993 = v93
		v1994 = v94
		v1995 = v95
		v1998 = v98
		v1999 = v99
		goto L10
	} else {
		goto L381
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1480
	v1516 = v1500
	goto L333
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1068]))) = v1437
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+12))
	v1464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1437)+11)))
	v1480 = v1463
	v1500 = v1464
	goto L334
L336:
	;
	goto L341
L337:
	;
	goto L338
L338:
	;
	v1320 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1269))))
	v1331 = int32(1621536)
	v1336 = int32(1622672)
	goto L354
L339:
	;
	if v1309-v1310 == int32(0) {
		v1437 = v1272
		goto L335
	} else {
		goto L353
	}
L341:
	;
	goto L342
L342:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269))))
	if v1279 != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1280 = v1269
	v1281 = v1272
	v1282 = int32(10)
	v1283 = v1279
	goto L347
L344:
	;
	v1305 = v1272
	v1309 = int32(0)
	goto L345
L345:
	;
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305))))
	goto L339
L346:
	;
	v1305 = v1300
	v1309 = v1302
	goto L345
L347:
	;
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281))))
	if v1283 != v1285 {
		v1300 = v1281
		v1302 = v1283
		goto L346
	} else {
		goto L349
	}
L348:
	;
	v1300 = v1294
	v1302 = int32(0)
	goto L346
L349:
	;
	if v1285 == int32(0) {
		v1300 = v1281
		v1302 = v1283
		goto L346
	} else {
		goto L350
	}
L350:
	;
	v1290 = v1282 - int32(1)
	if v1290 == int32(0) {
		v1300 = v1281
		v1302 = v1283
		goto L346
	} else {
		goto L351
	}
L351:
	;
	v1293 = int32(1)
	v1294 = v1281 + v1293
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+1)))
	if v1295 != 0 {
		v1280 = v1280 + v1293
		v1281 = v1294
		v1282 = v1290
		v1283 = v1295
		goto L347
	} else {
		goto L352
	}
L352:
	;
	goto L348
L353:
	;
	goto L338
L354:
	;
	v1363 = v1331 + (v1336-v1331)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v1364 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1363))))
	v1365 = v1320 - v1364
	if v1365 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	v1480 = v1416
	v1500 = int32(31)
	goto L334
L356:
	;
	goto L361
L357:
	;
	v1415 = v1365
	goto L358
L358:
	;
	v1416 = int32(0)
	v1420 = base.B2i32(v1415 < v1416)
	if v1415 < v1416 {
		goto L374
	} else {
		goto L375
	}
L359:
	;
	if v1406 == int32(0) {
		v1437 = v1363
		goto L335
	} else {
		goto L373
	}
L361:
	;
	goto L362
L362:
	;
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269))))
	if v1374 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1375 = v1269
	v1376 = v1363
	v1377 = int32(10)
	v1378 = v1374
	goto L367
L364:
	;
	v1400 = v1363
	v1404 = int32(0)
	goto L365
L365:
	;
	v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400))))
	v1406 = v1404 - v1405
	goto L359
L366:
	;
	v1400 = v1395
	v1404 = v1397
	goto L365
L367:
	;
	v1380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376))))
	if v1378 != v1380 {
		v1395 = v1376
		v1397 = v1378
		goto L366
	} else {
		goto L369
	}
L368:
	;
	v1395 = v1389
	v1397 = int32(0)
	goto L366
L369:
	;
	if v1380 == int32(0) {
		v1395 = v1376
		v1397 = v1378
		goto L366
	} else {
		goto L370
	}
L370:
	;
	v1385 = v1377 - int32(1)
	if v1385 == int32(0) {
		v1395 = v1376
		v1397 = v1378
		goto L366
	} else {
		goto L371
	}
L371:
	;
	v1388 = int32(1)
	v1389 = v1376 + v1388
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375)+1)))
	if v1390 != 0 {
		v1375 = v1375 + v1388
		v1376 = v1389
		v1377 = v1385
		v1378 = v1390
		goto L367
	} else {
		goto L372
	}
L372:
	;
	goto L368
L373:
	;
	v1415 = v1406
	goto L358
L374:
	;
	v1421 = v1363 - int32(16)
	goto L376
L375:
	;
	v1421 = v1336
	goto L376
L376:
	;
	if v1415 < v1416 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1424 = v1331
	goto L379
L378:
	;
	v1424 = v1363 + int32(16)
	goto L379
L379:
	;
	if base.Ui32(v1424) <= base.Ui32(v1421) {
		v1331 = v1424
		v1336 = v1421
		goto L354
	} else {
		goto L380
	}
L380:
	;
	goto L355
L381:
	;
	v1541 = int32(1) << (uint(v1516) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1541
	v1543 = int32(-1)
	switch v1516 {
	case 0:
		goto L393
	case 1:
		goto L392
	default:
		v2328 = v1543
		v2329 = v38
		goto L4
	case 5:
		goto L389
	case 6:
		goto L390
	case 7:
		goto L388
	case 9:
		goto L387
	case 16:
		goto L385
	case 17:
		goto L384
	case 18:
		goto L386
	case 23:
		goto L383
	case 28:
		goto L391
	case 31:
		goto L382
	}
L382:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	v1889 = F_pg_tzset(m, v1888)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L122
	} else {
		goto L442
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	v1883 = int32(14)
	if v79&v1883 != v1883 {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L440
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	if v78 != 0 {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L439
	}
L385:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+24)) = v1876
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L386:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = base.B2i32(v1873 == int32(1))
	goto L11
L387:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1939 = v78
	v1947 = v1872
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1541 | int32(32)
	if l6 == int32(0) {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L438
	}
L389:
	;
	v1857 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1857
	if l6 == v1857 {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L437
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1541 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L436
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1541 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L435
	}
L392:
	;
	if v94|base.B2i32(v79&int32(2) == int32(0)) != 0 {
		goto L431
	} else {
		goto L432
	}
L393:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	switch v1544 - int32(9) {
	case 0, 1, 2:
		goto L395
	case 3:
		goto L400
	case 4:
		goto L399
	case 5:
		goto L398
	case 6:
		goto L397
	case 7:
		goto L396
	default:
		goto L394
	}
L394:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L122
	} else {
		goto L428
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1544
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31776)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1792 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1792
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	if l6 == v1792 {
		v1939 = v78
		v1947 = v86
		v1950 = v89
		v1954 = v93
		v1955 = v94
		v1956 = v95
		v1959 = v98
		v1960 = v99
		goto L11
	} else {
		goto L427
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L122
	} else {
		goto L415
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L122
	} else {
		goto L414
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L122
	} else {
		goto L402
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, l4, l5, l6)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L122
	} else {
		goto L401
	}
L401:
	;
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L402:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1570 = base.B2i32(int32(2) < v1568)
	if int32(2) < v1568 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1571 = int32(4800)
	goto L405
L404:
	;
	v1571 = int32(4799)
	goto L405
L405:
	;
	v1572 = v1565 + v1571
	v1577 = base.I32_div_s(v1572, int32(4))
	v1580 = base.I32_div_s(v1572, int32(-100))
	v1583 = base.I32_div_s(v1572, int32(400))
	if int32(2) < v1568 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1587 = int32(1)
	goto L408
L407:
	;
	v1587 = int32(13)
	goto L408
L408:
	;
	v1592 = base.I32_div_s((v1587+v1568)*int32(7834), int32(256))
	v1595 = v1564 + v1572*int32(365) + v1577 + v1580 + v1583 + v1592 - int32(32168)
	v1599 = v1595 + int32(32044)
	v1600 = int32(146097)
	v1601 = base.I32_div_u_s(v1599, v1600)
	v1602 = int32(3)
	v1608 = int32(2)
	v1613 = base.I32_div_u_s((v1601*int32(1073595727)+v1599)<<(uint(v1608)%32)|v1602, v1600)
	v1616 = v1595 + v1601*v1602 + v1613 + int32(32104)
	v1617 = int32(1461)
	v1618 = base.I32_div_u_s(v1616, v1617)
	v1621 = v1618*int32(-1461) + v1616
	v1623 = v1621 << (uint(v1608) % 32)
	if base.Ui32(v1617) <= base.Ui32(v1623) {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L410:
	;
	v1636 = base.I32_div_u_s(v1623, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1636 + v1618<<(uint(int32(2))%32) - int32(4800)
	v1644 = v1634 + int32(123)
	v1648 = int32(base.Ui32(v1644*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1644 - int32(base.Ui32(v1648*int32(7834))>>(uint(int32(8))%32))
	v1658 = base.I32_rem_u_s(v1648+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1658 + int32(1)
	goto L409
L411:
	;
	v1629 = base.I32_rem_u_s(v1621+int32(305), int32(365))
	v1634 = v1629
	goto L410
L412:
	;
	goto L413
L413:
	;
	v1633 = base.I32_rem_u_s(v1621+int32(306), int32(366))
	v1634 = v1633
	goto L410
L414:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1673
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1675
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1677
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L415:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1696 = base.B2i32(int32(2) < v1694)
	if int32(2) < v1694 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1697 = int32(4800)
	goto L418
L417:
	;
	v1697 = int32(4799)
	goto L418
L418:
	;
	v1698 = v1691 + v1697
	v1703 = base.I32_div_s(v1698, int32(4))
	v1706 = base.I32_div_s(v1698, int32(-100))
	v1709 = base.I32_div_s(v1698, int32(400))
	if int32(2) < v1694 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1713 = int32(1)
	goto L421
L420:
	;
	v1713 = int32(13)
	goto L421
L421:
	;
	v1718 = base.I32_div_s((v1713+v1694)*int32(7834), int32(256))
	v1721 = v1690 + v1698*int32(365) + v1703 + v1706 + v1709 + v1718 - int32(32166)
	v1725 = v1721 + int32(32044)
	v1726 = int32(146097)
	v1727 = base.I32_div_u_s(v1725, v1726)
	v1728 = int32(3)
	v1734 = int32(2)
	v1739 = base.I32_div_u_s((v1727*int32(1073595727)+v1725)<<(uint(v1734)%32)|v1728, v1726)
	v1742 = v1721 + v1727*v1728 + v1739 + int32(32104)
	v1743 = int32(1461)
	v1744 = base.I32_div_u_s(v1742, v1743)
	v1747 = v1744*int32(-1461) + v1742
	v1749 = v1747 << (uint(v1734) % 32)
	if base.Ui32(v1743) <= base.Ui32(v1749) {
		goto L424
	} else {
		goto L425
	}
L422:
	;
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L423:
	;
	v1762 = base.I32_div_u_s(v1749, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1762 + v1744<<(uint(int32(2))%32) - int32(4800)
	v1770 = v1760 + int32(123)
	v1774 = int32(base.Ui32(v1770*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1770 - int32(base.Ui32(v1774*int32(7834))>>(uint(int32(8))%32))
	v1784 = base.I32_rem_u_s(v1774+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1784 + int32(1)
	goto L422
L424:
	;
	v1755 = base.I32_rem_u_s(v1747+int32(305), int32(365))
	v1760 = v1755
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1759 = base.I32_rem_u_s(v1747+int32(306), int32(366))
	v1760 = v1759
	goto L423
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1544
	F_errmsg_internal(m, int32(464540), v38)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L122
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(479605), int32(1390), int32(361301))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L122
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1832
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = int32(1)
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L432:
	;
	if v79&int32(8) != 0 {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if base.Ui32(int32(30)) < base.Ui32(v1823-int32(1)) {
		goto L431
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(8)
	goto L431
L435:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1842 - v1843
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L436:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1854
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L437:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1862
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L438:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v1939 = v78
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v1871
	v1959 = v1870
	v1960 = v99
	goto L11
L439:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1939 = v1880
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L440:
	;
	if v78 != 0 {
		v2328 = v1543
		v2329 = v38
		goto L4
	} else {
		goto L441
	}
L441:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1939 = v1887
	v1947 = v86
	v1950 = v89
	v1954 = v93
	v1955 = v94
	v1956 = v95
	v1959 = v98
	v1960 = v99
	goto L11
L442:
	;
	if v1889 != 0 {
		v1913 = v1889
		goto L12
	} else {
		goto L443
	}
L443:
	;
	v2328 = v1543
	v2329 = v38
	goto L4
L444:
	;
	v2328 = int32(-1)
	v2329 = v38
	goto L4
L445:
	;
	goto L446
L446:
	;
	v1978 = v1939
	v1979 = v1963 | v79
	v1986 = v1947
	v1989 = v1950
	v1993 = v1954
	v1994 = v1955
	v1995 = v1956
	v1998 = v1959
	v1999 = v1960
	goto L10
L447:
	;
	goto L9
L448:
	;
	v2015 = v38
	v2018 = v1979
	v2025 = v1986
	v2028 = v1989
	v2032 = v1993
	v2034 = v1995
	v2037 = v1998
	v2038 = v1999
	goto L7
L449:
	;
	v2328 = int32(0)
	v2329 = v2015
	goto L4
L450:
	;
	goto L451
L451:
	;
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2015)+59)))
	if v2032 != 0 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	if v2213 != 0 {
		v2328 = v2213
		v2329 = v2015
		goto L4
	} else {
		goto L490
	}
L453:
	;
	if v2018&int32(32768) != 0 {
		goto L471
	} else {
		goto L472
	}
L454:
	;
	if v2018&int32(4) == int32(0) {
		goto L453
	} else {
		goto L455
	}
L455:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v2038 != 0 {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v2070
	goto L453
L457:
	;
	v2070 = int32(1) - v2050
	goto L456
L458:
	;
	if int32(0) < v2050 {
		goto L457
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	if v2045 != 0 {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	v2213 = int32(-2)
	goto L452
L462:
	;
	if v2050 < int32(0) {
		goto L465
	} else {
		goto L466
	}
L463:
	;
	goto L464
L464:
	;
	if int32(0) < v2050 {
		goto L453
	} else {
		goto L470
	}
L465:
	;
	v2213 = int32(-2)
	goto L452
L466:
	;
	goto L467
L467:
	;
	if base.Ui32(v2050) <= base.Ui32(int32(69)) {
		v2070 = v2050 + int32(2000)
		goto L456
	} else {
		goto L468
	}
L468:
	;
	if base.Ui32(int32(99)) < base.Ui32(v2050) {
		goto L453
	} else {
		goto L469
	}
L469:
	;
	v2070 = v2050 + int32(1900)
	goto L456
L470:
	;
	v2213 = int32(-2)
	goto L452
L471:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v2081 = v2076 + int32(4799)
	v2083 = base.I32_div_s(v2081, int32(4))
	v2086 = base.I32_div_s(v2081, int32(-100))
	v2089 = base.I32_div_s(v2081, int32(400))
	v2090 = v2075 + v2076*int32(365) + v2083 + v2086 + v2089
	v2092 = v2090 + int32(1751940)
	v2093 = int32(146097)
	v2094 = base.I32_div_u_s(v2092, v2093)
	v2095 = int32(3)
	v2101 = int32(2)
	v2106 = base.I32_div_u_s((v2094*int32(1073595727)+v2092)<<(uint(v2101)%32)|v2095, v2093)
	v2109 = v2090 + v2094*v2095 + v2106 + int32(1752000)
	v2110 = int32(1461)
	v2111 = base.I32_div_u_s(v2109, v2110)
	v2114 = v2111*int32(-1461) + v2109
	v2116 = v2114 << (uint(v2101) % 32)
	if base.Ui32(v2110) <= base.Ui32(v2116) {
		goto L475
	} else {
		goto L476
	}
L472:
	;
	goto L473
L473:
	;
	if v2018&int32(2) == int32(0) {
		goto L478
	} else {
		goto L479
	}
L474:
	;
	v2129 = base.I32_div_u_s(v2116, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v2129 + v2111<<(uint(int32(2))%32) - int32(4800)
	v2137 = v2127 + int32(123)
	v2141 = int32(base.Ui32(v2137*int32(2141)) >> (uint(int32(16)) % 32))
	v2145 = base.I32_rem_u_s(v2141+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v2145 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v2137 - int32(base.Ui32(v2141*int32(7834))>>(uint(int32(8))%32))
	goto L473
L475:
	;
	v2122 = base.I32_rem_u_s(v2114+int32(305), int32(365))
	v2127 = v2122
	goto L474
L476:
	;
	goto L477
L477:
	;
	v2126 = base.I32_rem_u_s(v2114+int32(306), int32(366))
	v2127 = v2126
	goto L474
L478:
	;
	if v2018&int32(8) == int32(0) {
		goto L481
	} else {
		goto L482
	}
L479:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v2162-int32(13)) {
		goto L478
	} else {
		goto L480
	}
L480:
	;
	v2213 = int32(-3)
	goto L452
L481:
	;
	v2178 = int32(14)
	if v2018&v2178 != v2178 {
		goto L484
	} else {
		goto L485
	}
L482:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v2172-int32(32)) {
		goto L481
	} else {
		goto L483
	}
L483:
	;
	v2213 = int32(-3)
	goto L452
L484:
	;
	v2213 = int32(0)
	goto L452
L485:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v2184&int32(3) != 0 {
		v2194 = int32(0)
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2194*int32(52)+v2197<<(uint(int32(2))%32))+uint32(_consts[1071])))
	if v2182 <= v2203 {
		goto L484
	} else {
		goto L489
	}
L487:
	;
	v2189 = base.I32_rem_s(v2184, int32(100))
	if v2189 != 0 {
		v2194 = int32(1)
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v2191 = base.I32_rem_s(v2184, int32(400))
	v2194 = base.B2i32(v2191 == int32(0))
	goto L486
L489:
	;
	v2213 = int32(-2)
	goto L452
L490:
	;
	if v2025 == int32(2) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v2230 = int32(14)
	if v2018&v2230 != v2230 {
		goto L501
	} else {
		goto L502
	}
L492:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if int32(12) < v2216 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v2328 = int32(-2)
	v2329 = v2015
	goto L4
L494:
	;
	goto L495
L495:
	;
	switch v2025 {
	case 0:
		goto L498
	case 1:
		goto L497
	default:
		goto L491
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v2227
	goto L491
L497:
	;
	if v2216 == int32(12) {
		goto L491
	} else {
		goto L500
	}
L498:
	;
	if v2216 == int32(12) {
		v2227 = int32(0)
		goto L496
	} else {
		goto L499
	}
L499:
	;
	goto L491
L500:
	;
	v2227 = v2216 + int32(12)
	goto L496
L501:
	;
	v2236 = int32(31744)
	if v2018&v2236 == v2236 {
		goto L504
	} else {
		goto L505
	}
L502:
	;
	goto L503
L503:
	;
	if v2028 != 0 {
		goto L507
	} else {
		goto L508
	}
L504:
	;
	v2240 = int32(1)
	goto L506
L505:
	;
	v2240 = int32(-1)
	goto L506
L506:
	;
	v2328 = v2240
	v2329 = v2015
	goto L4
L507:
	;
	if v2018&int32(268435456) != 0 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	goto L509
L509:
	;
	if v2034 != 0 {
		goto L513
	} else {
		goto L514
	}
L510:
	;
	v2328 = int32(-1)
	v2329 = v2015
	goto L4
L511:
	;
	goto L512
L512:
	;
	v2246 = F_DetermineTimeZoneOffsetInternal(m, l4, v2028, v2015+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2246
	goto L509
L513:
	;
	if v2018&int32(268435456) != 0 {
		goto L516
	} else {
		goto L517
	}
L514:
	;
	goto L515
L515:
	;
	if l6 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L516:
	;
	v2328 = int32(-1)
	v2329 = v2015
	goto L4
L517:
	;
	goto L518
L518:
	;
	v2254 = m.G0
	v2256 = v2254 - int32(288)
	m.G0 = v2256
	v2260 = F_DetermineTimeZoneOffsetInternal(m, l4, v2034, v2256+int32(280))
	mBase = m.M
	v2264 = F_strlcpy(m, v2256+int32(16), v2037, int32(256))
	mBase = m.M
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2256)+16)))
	if v2265 != 0 {
		goto L520
	} else {
		goto L521
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2300
	goto L515
L520:
	;
	v2269 = v2256 + int32(16)
	v2273 = v2265
	goto L523
L521:
	;
	goto L522
L522:
	;
	v2293 = F_pg_interpret_timezone_abbrev(m, v2256+int32(16), v2256+int32(280), v2256+int32(12), v2256+int32(8), v2034)
	mBase = m.M
	if v2293 != 0 {
		goto L526
	} else {
		goto L527
	}
L523:
	;
	v2274 = F_pg_toupper(m, v2273)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v2269))) = uint8(v2274)
	v2277 = v2269 + int32(1)
	v2278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2277))))
	if v2278 != 0 {
		v2269 = v2277
		v2273 = v2278
		goto L523
	} else {
		goto L525
	}
L524:
	;
	goto L522
L525:
	;
	goto L524
L526:
	;
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+12))
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2295
	v2300 = int32(0) - v2294
	goto L528
L527:
	;
	v2300 = v2260
	goto L528
L528:
	;
	m.G0 = v2256 + int32(288)
	goto L519
L529:
	;
	v2328 = int32(0)
	v2329 = v2015
	goto L4
L530:
	;
	goto L531
L531:
	;
	v2308 = int32(0)
	if v2018&int32(32) != 0 {
		v2328 = v2308
		v2329 = v2015
		goto L4
	} else {
		goto L532
	}
L532:
	;
	if v2018&int32(268435456) != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2328 = int32(-1)
	v2329 = v2015
	goto L4
L534:
	;
	goto L535
L535:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
	v2318 = F_DetermineTimeZoneOffsetInternal(m, l4, v2315, v2015+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2318
	v2328 = v2308
	v2329 = v2015
	goto L4
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1717 int32
	_ = v1717
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v19 = l2 & base.B2i32(int32(0) <= v16)
	switch l5 - int32(1) {
	case 0, 3:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L3
	default:
		goto L2
	}
L1:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1831 <= int32(0) {
		goto L482
	} else {
		goto L483
	}
L2:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1344 = base.B2i32(int32(2) < v1342)
	if int32(2) < v1342 {
		goto L363
	} else {
		goto L364
	}
L3:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v854 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v853) {
		goto L233
	} else {
		goto L234
	}
L4:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
	v361 = base.B2i32(v359 == int32(1))
	if v359 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = v22
	goto L8
L7:
	;
	v27 = int32(1) - v22
	goto L8
L8:
	;
	v28 = int32(4)
	if base.Ui32(int32(99)) < base.Ui32(v27) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v54 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	v57 = v53 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v59 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v58) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v42 = F_pg_ultoa_n(m, v27, l6)
	mBase = m.M
	if v28 <= v42 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L10
L13:
	;
	v53 = l6 + v42
	goto L9
L14:
	;
	goto L15
L15:
	;
	v45 = l6 + v28
	v47 = F_memmove(m, v45-v42, l6, v42)
	mBase = m.M
	v50 = F___memset(m, l6, int32(48), v28-v42)
	mBase = m.M
	v53 = v45
	goto L9
L16:
	;
	v85 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v85)
	v88 = v84 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v89) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v73 = F_pg_ultoa_n(m, v58, v57)
	mBase = m.M
	if v59 <= v73 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v69)
	v84 = v53 + int32(3)
	goto L16
L20:
	;
	v84 = v57 + v73
	goto L16
L21:
	;
	goto L22
L22:
	;
	v76 = v53 + int32(3)
	v78 = F_memmove(m, v76-v73, v57, v73)
	mBase = m.M
	v81 = F___memset(m, v57, int32(48), v59-v73)
	mBase = m.M
	v84 = v76
	goto L16
L23:
	;
	if l5 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v104 = F_pg_ultoa_n(m, v89, v88)
	mBase = m.M
	if v90 <= v104 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v88))) = uint16(v100)
	v115 = v84 + int32(3)
	goto L23
L27:
	;
	v115 = v88 + v104
	goto L23
L28:
	;
	goto L29
L29:
	;
	v107 = v84 + int32(3)
	v109 = F_memmove(m, v107-v104, v88, v104)
	mBase = m.M
	v112 = F___memset(m, v88, int32(48), v90-v104)
	mBase = m.M
	v115 = v107
	goto L23
L30:
	;
	v120 = int32(32)
	goto L32
L31:
	;
	v120 = int32(84)
	goto L32
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v120)
	v123 = v115 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v125 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v124) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v151 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v151)
	v154 = v150 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v156 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v155) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v139 = F_pg_ultoa_n(m, v124, v123)
	mBase = m.M
	if v125 <= v139 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v135)
	v150 = v115 + int32(3)
	goto L33
L37:
	;
	v150 = v123 + v139
	goto L33
L38:
	;
	goto L39
L39:
	;
	v142 = v115 + int32(3)
	v144 = F_memmove(m, v142-v139, v123, v139)
	mBase = m.M
	v147 = F___memset(m, v123, int32(48), v125-v139)
	mBase = m.M
	v150 = v142
	goto L33
L40:
	;
	v182 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v182)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = v186 >> (uint(int32(31)) % 32)
	goto L49
L41:
	;
	v170 = F_pg_ultoa_n(m, v155, v154)
	mBase = m.M
	if v156 <= v170 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v166)
	v181 = v150 + int32(3)
	goto L40
L44:
	;
	v181 = v154 + v170
	goto L40
L45:
	;
	goto L46
L46:
	;
	v173 = v150 + int32(3)
	v175 = F_memmove(m, v173-v170, v154, v170)
	mBase = m.M
	v178 = F___memset(m, v154, int32(48), v156-v170)
	mBase = m.M
	v181 = v173
	goto L40
L47:
	;
	if v19 == int32(0) {
		v1830 = v304
		goto L1
	} else {
		goto L82
	}
L48:
	;
	if l1 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v196 = F_pg_ultostr_zeropad(m, v181+int32(1), v186^v192-v192, int32(2))
	mBase = m.M
	goto L48
L52:
	;
	v304 = v196
	goto L47
L53:
	;
	goto L54
L54:
	;
	v201 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v201)
	v204 = l1 >> (uint(int32(31)) % 32)
	v206 = l1 ^ v204 - v204
	v208 = base.I32_div_s(v206, int32(10))
	v211 = v208*int32(-10) + v206
	if v211 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v221 = base.I32_div_s(v206, int32(100))
	v224 = v221*int32(-10) + v208
	v225 = v211 | v224
	if v225 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v213 = v211 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+6)) = uint8(v213)
	v219 = v196 + int32(7)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v219 = v196 + int32(6)
	goto L55
L59:
	;
	v235 = base.I32_div_s(v206, int32(1000))
	v238 = v235*int32(-10) + v221
	v239 = v225 | v238
	if v239 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v233 = v196 + int32(5)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v231 = v224 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+5)) = uint8(v231)
	v233 = v219
	goto L59
L63:
	;
	v249 = base.I32_div_s(v206, int32(10000))
	v252 = v249*int32(-10) + v235
	v253 = v239 | v252
	if v253 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v247 = v196 + int32(4)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v245 = v238 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)) = uint8(v245)
	v247 = v233
	goto L63
L67:
	;
	v263 = base.I32_div_s(v206, int32(100000))
	v266 = v263*int32(-10) + v249
	v267 = v253 | v266
	if v267 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v261 = v196 + int32(3)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v259 = v252 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+3)) = uint8(v259)
	v261 = v247
	goto L67
L71:
	;
	v277 = base.I32_div_s(v206, int32(1000000))
	v280 = v277*int32(-10) + v263
	if v267|v280 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v275 = v196 + int32(2)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v273 = v266 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+2)) = uint8(v273)
	v275 = v261
	goto L71
L75:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v263+int32(9)) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v289 = v196 + int32(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v287 = v280 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)) = uint8(v287)
	v289 = v275
	goto L75
L79:
	;
	v296 = F_pg_ultostr(m, v196+int32(1), v206)
	mBase = m.M
	v297 = v296
	goto L81
L80:
	;
	v297 = v289
	goto L81
L81:
	;
	v304 = v297
	goto L47
L82:
	;
	if l3 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v1830 = v354
	goto L1
L84:
	;
	v314 = int32(43)
	goto L86
L85:
	;
	v314 = int32(45)
	goto L86
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v314)
	v317 = l3 >> (uint(int32(31)) % 32)
	v319 = l3 ^ v317 - v317
	v321 = base.I32_div_s(v319, int32(3600))
	v322 = int32(-60)
	v325 = base.I32_div_s(v319, int32(60))
	v326 = v321*v322 + v325
	v328 = v304 + int32(1)
	v331 = v325*v322 + v319
	if v331 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L83
L88:
	;
	v348 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v346))) = uint8(v348)
	v353 = F_pg_ultostr_zeropad(m, v346+int32(1), v347, int32(2))
	mBase = m.M
	v354 = v353
	goto L87
L89:
	;
	v332 = int32(2)
	v333 = F_pg_ultostr_zeropad(m, v328, v321, v332)
	mBase = m.M
	v334 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v334)
	v339 = F_pg_ultostr_zeropad(m, v333+int32(1), v326, v332)
	mBase = m.M
	v346 = v339
	v347 = v331
	goto L88
L90:
	;
	goto L91
L91:
	;
	v341 = F_pg_ultostr_zeropad(m, v328, v321, int32(2))
	mBase = m.M
	if l5 == int32(4) {
		v346 = v341
		v347 = v326
		goto L88
	} else {
		goto L92
	}
L92:
	;
	if v326 == int32(0) {
		v354 = v341
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v346 = v341
	v347 = v326
	goto L88
L94:
	;
	v362 = int32(12)
	goto L96
L95:
	;
	v362 = int32(16)
	goto L96
L96:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0+v362)))
	v365 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v364) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v391 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v391)
	v394 = v390 + int32(1)
	if v359 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v379 = F_pg_ultoa_n(m, v364, l6)
	mBase = m.M
	if v365 <= v379 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v375)
	v390 = l6 + int32(2)
	goto L97
L101:
	;
	v390 = l6 + v379
	goto L97
L102:
	;
	goto L103
L103:
	;
	v382 = l6 + v365
	v384 = F_memmove(m, v382-v379, l6, v379)
	mBase = m.M
	v387 = F___memset(m, l6, int32(48), v365-v379)
	mBase = m.M
	v390 = v382
	goto L97
L104:
	;
	v397 = int32(16)
	goto L106
L105:
	;
	v397 = int32(12)
	goto L106
L106:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0+v397)))
	v400 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v399) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v426 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v426)
	v428 = int32(1)
	v429 = v425 + v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v430 {
		goto L114
	} else {
		goto L115
	}
L108:
	;
	v414 = F_pg_ultoa_n(m, v399, v394)
	mBase = m.M
	if v400 <= v414 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v394))) = uint16(v410)
	v425 = v390 + int32(3)
	goto L107
L111:
	;
	v425 = v394 + v414
	goto L107
L112:
	;
	goto L113
L113:
	;
	v417 = v390 + int32(3)
	v419 = F_memmove(m, v417-v414, v394, v414)
	mBase = m.M
	v422 = F___memset(m, v394, int32(48), v400-v414)
	mBase = m.M
	v425 = v417
	goto L107
L114:
	;
	v435 = v430
	goto L116
L115:
	;
	v435 = v428 - v430
	goto L116
L116:
	;
	v436 = int32(4)
	if base.Ui32(int32(99)) < base.Ui32(v435) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v462 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v462)
	v465 = v461 + int32(1)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v467 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v466) {
		goto L125
	} else {
		goto L126
	}
L118:
	;
	v450 = F_pg_ultoa_n(m, v435, v429)
	mBase = m.M
	if v436 <= v450 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L118
L121:
	;
	v461 = v429 + v450
	goto L117
L122:
	;
	goto L123
L123:
	;
	v453 = v425 + int32(5)
	v455 = F_memmove(m, v453-v450, v429, v450)
	mBase = m.M
	v458 = F___memset(m, v429, int32(48), v436-v450)
	mBase = m.M
	v461 = v453
	goto L117
L124:
	;
	v493 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v492))) = uint8(v493)
	v496 = v492 + int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v497) {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	v481 = F_pg_ultoa_n(m, v466, v465)
	mBase = m.M
	if v467 <= v481 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v465))) = uint16(v477)
	v492 = v461 + int32(3)
	goto L124
L128:
	;
	v492 = v465 + v481
	goto L124
L129:
	;
	goto L130
L130:
	;
	v484 = v461 + int32(3)
	v486 = F_memmove(m, v484-v481, v465, v481)
	mBase = m.M
	v489 = F___memset(m, v465, int32(48), v467-v481)
	mBase = m.M
	v492 = v484
	goto L124
L131:
	;
	v524 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v523))) = uint8(v524)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v534 = v528 >> (uint(int32(31)) % 32)
	goto L140
L132:
	;
	v512 = F_pg_ultoa_n(m, v497, v496)
	mBase = m.M
	if v498 <= v512 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v496))) = uint16(v508)
	v523 = v492 + int32(3)
	goto L131
L135:
	;
	v523 = v496 + v512
	goto L131
L136:
	;
	goto L137
L137:
	;
	v515 = v492 + int32(3)
	v517 = F_memmove(m, v515-v512, v496, v512)
	mBase = m.M
	v520 = F___memset(m, v496, int32(48), v498-v512)
	mBase = m.M
	v523 = v515
	goto L131
L138:
	;
	if v19 == int32(0) {
		v1830 = v646
		goto L1
	} else {
		goto L173
	}
L139:
	;
	if l1 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v538 = F_pg_ultostr_zeropad(m, v523+int32(1), v528^v534-v534, int32(2))
	mBase = m.M
	goto L139
L143:
	;
	v646 = v538
	goto L138
L144:
	;
	goto L145
L145:
	;
	v543 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v538))) = uint8(v543)
	v546 = l1 >> (uint(int32(31)) % 32)
	v548 = l1 ^ v546 - v546
	v550 = base.I32_div_s(v548, int32(10))
	v553 = v550*int32(-10) + v548
	if v553 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v563 = base.I32_div_s(v548, int32(100))
	v566 = v563*int32(-10) + v550
	v567 = v553 | v566
	if v567 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v555 = v553 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+6)) = uint8(v555)
	v561 = v538 + int32(7)
	goto L146
L148:
	;
	goto L149
L149:
	;
	v561 = v538 + int32(6)
	goto L146
L150:
	;
	v577 = base.I32_div_s(v548, int32(1000))
	v580 = v577*int32(-10) + v563
	v581 = v567 | v580
	if v581 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v575 = v538 + int32(5)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v573 = v566 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+5)) = uint8(v573)
	v575 = v561
	goto L150
L154:
	;
	v591 = base.I32_div_s(v548, int32(10000))
	v594 = v591*int32(-10) + v577
	v595 = v581 | v594
	if v595 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v589 = v538 + int32(4)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v587 = v580 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+4)) = uint8(v587)
	v589 = v575
	goto L154
L158:
	;
	v605 = base.I32_div_s(v548, int32(100000))
	v608 = v605*int32(-10) + v591
	v609 = v595 | v608
	if v609 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L159:
	;
	v603 = v538 + int32(3)
	goto L158
L160:
	;
	goto L161
L161:
	;
	v601 = v594 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+3)) = uint8(v601)
	v603 = v589
	goto L158
L162:
	;
	v619 = base.I32_div_s(v548, int32(1000000))
	v622 = v619*int32(-10) + v605
	if v609|v622 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	v617 = v538 + int32(2)
	goto L162
L164:
	;
	goto L165
L165:
	;
	v615 = v608 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+2)) = uint8(v615)
	v617 = v603
	goto L162
L166:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v605+int32(9)) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v631 = v538 + int32(1)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v629 = v622 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+1)) = uint8(v629)
	v631 = v617
	goto L166
L170:
	;
	v638 = F_pg_ultostr(m, v538+int32(1), v548)
	mBase = m.M
	v639 = v638
	goto L172
L171:
	;
	v639 = v631
	goto L172
L172:
	;
	v646 = v639
	goto L138
L173:
	;
	if l4 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(10)
	v655 = F_pg_sprintf(m, v646, int32(166845), v14+int32(16))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if l3 <= int32(0) {
		goto L196
	} else {
		goto L197
	}
L177:
	;
	return
L178:
	;
	if v646&int32(3) == int32(0) {
		v680 = v646
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v1830 = v713 + v646
	goto L1
L180:
	;
	v713 = v705 - v646
	goto L179
L181:
	;
	v684 = v680
	goto L190
L182:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	if v664 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v713 = int32(0)
	goto L179
L184:
	;
	goto L185
L185:
	;
	v669 = v646
	goto L186
L186:
	;
	v673 = v669 + int32(1)
	if v673&int32(3) == int32(0) {
		v680 = v673
		goto L181
	} else {
		goto L188
	}
L187:
	;
	v705 = v673
	goto L180
L188:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	if v678 != 0 {
		v669 = v673
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	v693 = int32(-2139062144)
	if (int32(16843008)-v690|v690)&v693 == v693 {
		v684 = v684 + int32(4)
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v699 = v684
	goto L193
L192:
	;
	goto L191
L193:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v703 != 0 {
		v699 = v699 + int32(1)
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v705 = v699
	goto L180
L195:
	;
	goto L194
L196:
	;
	v719 = int32(43)
	goto L198
L197:
	;
	v719 = int32(45)
	goto L198
L198:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v646))) = uint8(v719)
	v722 = l3 >> (uint(int32(31)) % 32)
	v724 = l3 ^ v722 - v722
	v726 = base.I32_div_s(v724, int32(3600))
	v727 = int32(-60)
	v730 = base.I32_div_s(v724, int32(60))
	v731 = v726*v727 + v730
	v733 = v646 + int32(1)
	v736 = v730*v727 + v724
	if v736 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v823 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v822))) = uint8(v823)
	v826 = v822 + int32(1)
	v827 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v821) {
		goto L226
	} else {
		goto L227
	}
L200:
	;
	v737 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v726) {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	goto L202
L202:
	;
	v793 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v726) {
		goto L218
	} else {
		goto L219
	}
L203:
	;
	v763 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v762))) = uint8(v763)
	v766 = v762 + int32(1)
	v767 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v731) {
		goto L211
	} else {
		goto L212
	}
L204:
	;
	v751 = F_pg_ultoa_n(m, v726, v733)
	mBase = m.M
	if v737 <= v751 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v733))) = uint16(v747)
	v762 = v646 + int32(3)
	goto L203
L207:
	;
	v762 = v733 + v751
	goto L203
L208:
	;
	goto L209
L209:
	;
	v754 = v646 + int32(3)
	v756 = F_memmove(m, v754-v751, v733, v751)
	mBase = m.M
	v759 = F___memset(m, v733, int32(48), v737-v751)
	mBase = m.M
	v762 = v754
	goto L203
L210:
	;
	v821 = v736
	v822 = v792
	goto L199
L211:
	;
	v781 = F_pg_ultoa_n(m, v731, v766)
	mBase = m.M
	if v767 <= v781 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v777 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v766))) = uint16(v777)
	v792 = v762 + int32(3)
	goto L210
L214:
	;
	v792 = v766 + v781
	goto L210
L215:
	;
	goto L216
L216:
	;
	v784 = v762 + int32(3)
	v786 = F_memmove(m, v784-v781, v766, v781)
	mBase = m.M
	v789 = F___memset(m, v766, int32(48), v767-v781)
	mBase = m.M
	v792 = v784
	goto L210
L217:
	;
	if v731 == int32(0) {
		v1830 = v818
		goto L1
	} else {
		goto L224
	}
L218:
	;
	v807 = F_pg_ultoa_n(m, v726, v733)
	mBase = m.M
	if v793 <= v807 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v733))) = uint16(v803)
	v818 = v646 + int32(3)
	goto L217
L221:
	;
	v818 = v733 + v807
	goto L217
L222:
	;
	goto L223
L223:
	;
	v810 = v646 + int32(3)
	v812 = F_memmove(m, v810-v807, v733, v807)
	mBase = m.M
	v815 = F___memset(m, v733, int32(48), v793-v807)
	mBase = m.M
	v818 = v810
	goto L217
L224:
	;
	v821 = v731
	v822 = v818
	goto L199
L225:
	;
	v1830 = v852
	goto L1
L226:
	;
	v841 = F_pg_ultoa_n(m, v821, v826)
	mBase = m.M
	if v827 <= v841 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v821<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v826))) = uint16(v837)
	v852 = v822 + int32(3)
	goto L225
L229:
	;
	v852 = v826 + v841
	goto L225
L230:
	;
	goto L231
L231:
	;
	v844 = v822 + int32(3)
	v846 = F_memmove(m, v844-v841, v826, v841)
	mBase = m.M
	v849 = F___memset(m, v826, int32(48), v827-v841)
	mBase = m.M
	v852 = v844
	goto L225
L232:
	;
	v880 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v879))) = uint8(v880)
	v883 = v879 + int32(1)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v885 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v884) {
		goto L240
	} else {
		goto L241
	}
L233:
	;
	v868 = F_pg_ultoa_n(m, v853, l6)
	mBase = m.M
	if v854 <= v868 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v853<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v864)
	v879 = l6 + int32(2)
	goto L232
L236:
	;
	v879 = l6 + v868
	goto L232
L237:
	;
	goto L238
L238:
	;
	v871 = l6 + v854
	v873 = F_memmove(m, v871-v868, l6, v868)
	mBase = m.M
	v876 = F___memset(m, l6, int32(48), v854-v868)
	mBase = m.M
	v879 = v871
	goto L232
L239:
	;
	v911 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v910))) = uint8(v911)
	v913 = int32(1)
	v914 = v910 + v913
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v915 {
		goto L246
	} else {
		goto L247
	}
L240:
	;
	v899 = F_pg_ultoa_n(m, v884, v883)
	mBase = m.M
	if v885 <= v899 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v883))) = uint16(v895)
	v910 = v879 + int32(3)
	goto L239
L243:
	;
	v910 = v883 + v899
	goto L239
L244:
	;
	goto L245
L245:
	;
	v902 = v879 + int32(3)
	v904 = F_memmove(m, v902-v899, v883, v899)
	mBase = m.M
	v907 = F___memset(m, v883, int32(48), v885-v899)
	mBase = m.M
	v910 = v902
	goto L239
L246:
	;
	v920 = v915
	goto L248
L247:
	;
	v920 = v913 - v915
	goto L248
L248:
	;
	v921 = int32(4)
	if base.Ui32(int32(99)) < base.Ui32(v920) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v947 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v946))) = uint8(v947)
	v950 = v946 + int32(1)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v952 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v951) {
		goto L257
	} else {
		goto L258
	}
L250:
	;
	v935 = F_pg_ultoa_n(m, v920, v914)
	mBase = m.M
	if v921 <= v935 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L250
L253:
	;
	v946 = v914 + v935
	goto L249
L254:
	;
	goto L255
L255:
	;
	v938 = v910 + int32(5)
	v940 = F_memmove(m, v938-v935, v914, v935)
	mBase = m.M
	v943 = F___memset(m, v914, int32(48), v921-v935)
	mBase = m.M
	v946 = v938
	goto L249
L256:
	;
	v978 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v977))) = uint8(v978)
	v981 = v977 + int32(1)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v983 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v982) {
		goto L264
	} else {
		goto L265
	}
L257:
	;
	v966 = F_pg_ultoa_n(m, v951, v950)
	mBase = m.M
	if v952 <= v966 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v951<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v950))) = uint16(v962)
	v977 = v946 + int32(3)
	goto L256
L260:
	;
	v977 = v950 + v966
	goto L256
L261:
	;
	goto L262
L262:
	;
	v969 = v946 + int32(3)
	v971 = F_memmove(m, v969-v966, v950, v966)
	mBase = m.M
	v974 = F___memset(m, v950, int32(48), v952-v966)
	mBase = m.M
	v977 = v969
	goto L256
L263:
	;
	v1009 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1008))) = uint8(v1009)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1019 = v1013 >> (uint(int32(31)) % 32)
	goto L272
L264:
	;
	v997 = F_pg_ultoa_n(m, v982, v981)
	mBase = m.M
	if v983 <= v997 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	v993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v982<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v981))) = uint16(v993)
	v1008 = v977 + int32(3)
	goto L263
L267:
	;
	v1008 = v981 + v997
	goto L263
L268:
	;
	goto L269
L269:
	;
	v1000 = v977 + int32(3)
	v1002 = F_memmove(m, v1000-v997, v981, v997)
	mBase = m.M
	v1005 = F___memset(m, v981, int32(48), v983-v997)
	mBase = m.M
	v1008 = v1000
	goto L263
L270:
	;
	if v19 == int32(0) {
		v1830 = v1131
		goto L1
	} else {
		goto L305
	}
L271:
	;
	if l1 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v1023 = F_pg_ultostr_zeropad(m, v1008+int32(1), v1013^v1019-v1019, int32(2))
	mBase = m.M
	goto L271
L275:
	;
	v1131 = v1023
	goto L270
L276:
	;
	goto L277
L277:
	;
	v1028 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023))) = uint8(v1028)
	v1031 = l1 >> (uint(int32(31)) % 32)
	v1033 = l1 ^ v1031 - v1031
	v1035 = base.I32_div_s(v1033, int32(10))
	v1038 = v1035*int32(-10) + v1033
	if v1038 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1048 = base.I32_div_s(v1033, int32(100))
	v1051 = v1048*int32(-10) + v1035
	v1052 = v1038 | v1051
	if v1052 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	v1040 = v1038 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+6)) = uint8(v1040)
	v1046 = v1023 + int32(7)
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1046 = v1023 + int32(6)
	goto L278
L282:
	;
	v1062 = base.I32_div_s(v1033, int32(1000))
	v1065 = v1062*int32(-10) + v1048
	v1066 = v1052 | v1065
	if v1066 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	v1060 = v1023 + int32(5)
	goto L282
L284:
	;
	goto L285
L285:
	;
	v1058 = v1051 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+5)) = uint8(v1058)
	v1060 = v1046
	goto L282
L286:
	;
	v1076 = base.I32_div_s(v1033, int32(10000))
	v1079 = v1076*int32(-10) + v1062
	v1080 = v1066 | v1079
	if v1080 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L287:
	;
	v1074 = v1023 + int32(4)
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1072 = v1065 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+4)) = uint8(v1072)
	v1074 = v1060
	goto L286
L290:
	;
	v1090 = base.I32_div_s(v1033, int32(100000))
	v1093 = v1090*int32(-10) + v1076
	v1094 = v1080 | v1093
	if v1094 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L291:
	;
	v1088 = v1023 + int32(3)
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1086 = v1079 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+3)) = uint8(v1086)
	v1088 = v1074
	goto L290
L294:
	;
	v1104 = base.I32_div_s(v1033, int32(1000000))
	v1107 = v1104*int32(-10) + v1090
	if v1094|v1107 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	v1102 = v1023 + int32(2)
	goto L294
L296:
	;
	goto L297
L297:
	;
	v1100 = v1093 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+2)) = uint8(v1100)
	v1102 = v1088
	goto L294
L298:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1090+int32(9)) {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	v1116 = v1023 + int32(1)
	goto L298
L300:
	;
	goto L301
L301:
	;
	v1114 = v1107 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+1)) = uint8(v1114)
	v1116 = v1102
	goto L298
L302:
	;
	v1123 = F_pg_ultostr(m, v1023+int32(1), v1033)
	mBase = m.M
	v1124 = v1123
	goto L304
L303:
	;
	v1124 = v1116
	goto L304
L304:
	;
	v1131 = v1124
	goto L270
L305:
	;
	if l4 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(10)
	v1140 = F_pg_sprintf(m, v1131, int32(166845), v14+int32(32))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L177
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	if l3 <= int32(0) {
		goto L327
	} else {
		goto L328
	}
L309:
	;
	if v1131&int32(3) == int32(0) {
		v1165 = v1131
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1830 = v1198 + v1131
	goto L1
L311:
	;
	v1198 = v1190 - v1131
	goto L310
L312:
	;
	v1169 = v1165
	goto L321
L313:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131))))
	if v1149 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1198 = int32(0)
	goto L310
L315:
	;
	goto L316
L316:
	;
	v1154 = v1131
	goto L317
L317:
	;
	v1158 = v1154 + int32(1)
	if v1158&int32(3) == int32(0) {
		v1165 = v1158
		goto L312
	} else {
		goto L319
	}
L318:
	;
	v1190 = v1158
	goto L311
L319:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158))))
	if v1163 != 0 {
		v1154 = v1158
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	v1178 = int32(-2139062144)
	if (int32(16843008)-v1175|v1175)&v1178 == v1178 {
		v1169 = v1169 + int32(4)
		goto L321
	} else {
		goto L323
	}
L322:
	;
	v1184 = v1169
	goto L324
L323:
	;
	goto L322
L324:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184))))
	if v1188 != 0 {
		v1184 = v1184 + int32(1)
		goto L324
	} else {
		goto L326
	}
L325:
	;
	v1190 = v1184
	goto L311
L326:
	;
	goto L325
L327:
	;
	v1204 = int32(43)
	goto L329
L328:
	;
	v1204 = int32(45)
	goto L329
L329:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1131))) = uint8(v1204)
	v1207 = l3 >> (uint(int32(31)) % 32)
	v1209 = l3 ^ v1207 - v1207
	v1211 = base.I32_div_s(v1209, int32(3600))
	v1212 = int32(-60)
	v1215 = base.I32_div_s(v1209, int32(60))
	v1216 = v1211*v1212 + v1215
	v1218 = v1131 + int32(1)
	v1221 = v1215*v1212 + v1209
	if v1221 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1308 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1307))) = uint8(v1308)
	v1311 = v1307 + int32(1)
	v1312 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1306) {
		goto L357
	} else {
		goto L358
	}
L331:
	;
	v1222 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1211) {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	goto L333
L333:
	;
	v1278 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1211) {
		goto L349
	} else {
		goto L350
	}
L334:
	;
	v1248 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1247))) = uint8(v1248)
	v1251 = v1247 + int32(1)
	v1252 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1216) {
		goto L342
	} else {
		goto L343
	}
L335:
	;
	v1236 = F_pg_ultoa_n(m, v1211, v1218)
	mBase = m.M
	if v1222 <= v1236 {
		goto L338
	} else {
		goto L339
	}
L336:
	;
	goto L337
L337:
	;
	v1232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1218))) = uint16(v1232)
	v1247 = v1131 + int32(3)
	goto L334
L338:
	;
	v1247 = v1218 + v1236
	goto L334
L339:
	;
	goto L340
L340:
	;
	v1239 = v1131 + int32(3)
	v1241 = F_memmove(m, v1239-v1236, v1218, v1236)
	mBase = m.M
	v1244 = F___memset(m, v1218, int32(48), v1222-v1236)
	mBase = m.M
	v1247 = v1239
	goto L334
L341:
	;
	v1306 = v1221
	v1307 = v1277
	goto L330
L342:
	;
	v1266 = F_pg_ultoa_n(m, v1216, v1251)
	mBase = m.M
	if v1252 <= v1266 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	goto L344
L344:
	;
	v1262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1216<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1251))) = uint16(v1262)
	v1277 = v1247 + int32(3)
	goto L341
L345:
	;
	v1277 = v1251 + v1266
	goto L341
L346:
	;
	goto L347
L347:
	;
	v1269 = v1247 + int32(3)
	v1271 = F_memmove(m, v1269-v1266, v1251, v1266)
	mBase = m.M
	v1274 = F___memset(m, v1251, int32(48), v1252-v1266)
	mBase = m.M
	v1277 = v1269
	goto L341
L348:
	;
	if v1216 == int32(0) {
		v1830 = v1303
		goto L1
	} else {
		goto L355
	}
L349:
	;
	v1292 = F_pg_ultoa_n(m, v1211, v1218)
	mBase = m.M
	if v1278 <= v1292 {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	goto L351
L351:
	;
	v1288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1211<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1218))) = uint16(v1288)
	v1303 = v1131 + int32(3)
	goto L348
L352:
	;
	v1303 = v1218 + v1292
	goto L348
L353:
	;
	goto L354
L354:
	;
	v1295 = v1131 + int32(3)
	v1297 = F_memmove(m, v1295-v1292, v1218, v1292)
	mBase = m.M
	v1300 = F___memset(m, v1218, int32(48), v1278-v1292)
	mBase = m.M
	v1303 = v1295
	goto L348
L355:
	;
	v1306 = v1216
	v1307 = v1303
	goto L330
L356:
	;
	v1830 = v1337
	goto L1
L357:
	;
	v1326 = F_pg_ultoa_n(m, v1306, v1311)
	mBase = m.M
	if v1312 <= v1326 {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	goto L359
L359:
	;
	v1322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1306<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1311))) = uint16(v1322)
	v1337 = v1307 + int32(3)
	goto L356
L360:
	;
	v1337 = v1311 + v1326
	goto L356
L361:
	;
	goto L362
L362:
	;
	v1329 = v1307 + int32(3)
	v1331 = F_memmove(m, v1329-v1326, v1311, v1326)
	mBase = m.M
	v1334 = F___memset(m, v1311, int32(48), v1312-v1326)
	mBase = m.M
	v1337 = v1329
	goto L356
L363:
	;
	v1345 = int32(4800)
	goto L365
L364:
	;
	v1345 = int32(4799)
	goto L365
L365:
	;
	v1346 = v1339 + v1345
	v1351 = base.I32_div_s(v1346, int32(4))
	v1354 = base.I32_div_s(v1346, int32(-100))
	v1357 = base.I32_div_s(v1346, int32(400))
	if int32(2) < v1342 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1361 = int32(1)
	goto L368
L367:
	;
	v1361 = int32(13)
	goto L368
L368:
	;
	v1366 = base.I32_div_s((v1361+v1342)*int32(7834), int32(256))
	v1370 = int32(7)
	v1371 = base.I32_rem_s(v1338+v1346*int32(365)+v1351+v1354+v1357+v1366-int32(32166), v1370)
	if v1371 < int32(0) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1376 = v1371 + v1370
	goto L371
L370:
	;
	v1376 = v1371
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1376
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1376<<(uint(int32(2))%32))+uint32(_consts[1074])))
	v1383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1382))))
	*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v1383)
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382)+2)))
	v1386 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)) = uint8(v1386)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+2)) = uint8(v1385)
	v1390 = l6 + int32(4)
	v1392 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
	if v1392 == int32(1) {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	v1481 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1480))) = uint8(v1481)
	v1484 = v1480 + int32(1)
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1486 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1485) {
		goto L391
	} else {
		goto L392
	}
L373:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1396 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1395) {
		goto L377
	} else {
		goto L378
	}
L374:
	;
	goto L375
L375:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1437 = int32(2)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1436<<(uint(v1437)%32))+uint32(_consts[1075])))
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1441))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1390))) = uint16(v1442)
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1390)+2)) = uint8(v1444)
	v1446 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+7)) = uint8(v1446)
	v1449 = l6 + int32(8)
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(int32(99)) < base.Ui32(v1450) {
		goto L384
	} else {
		goto L385
	}
L376:
	;
	v1422 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1421))) = uint8(v1422)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1424<<(uint(int32(2))%32))+uint32(_consts[1075])))
	v1430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1429))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1421)+1)) = uint16(v1430)
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1429)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1421)+3)) = uint8(v1432)
	v1480 = v1421 + int32(4)
	goto L372
L377:
	;
	v1410 = F_pg_ultoa_n(m, v1395, v1390)
	mBase = m.M
	if v1396 <= v1410 {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	goto L379
L379:
	;
	v1406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1395<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1390))) = uint16(v1406)
	v1421 = l6 + int32(6)
	goto L376
L380:
	;
	v1421 = v1390 + v1410
	goto L376
L381:
	;
	goto L382
L382:
	;
	v1413 = l6 + int32(6)
	v1415 = F_memmove(m, v1413-v1410, v1390, v1410)
	mBase = m.M
	v1418 = F___memset(m, v1390, int32(48), v1396-v1410)
	mBase = m.M
	v1421 = v1413
	goto L376
L383:
	;
	v1480 = v1476
	goto L372
L384:
	;
	v1465 = F_pg_ultoa_n(m, v1450, v1449)
	mBase = m.M
	if v1437 <= v1465 {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	goto L386
L386:
	;
	v1461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1450<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1449))) = uint16(v1461)
	v1476 = l6 + int32(10)
	goto L383
L387:
	;
	v1476 = v1449 + v1465
	goto L383
L388:
	;
	goto L389
L389:
	;
	v1468 = l6 + int32(10)
	v1470 = F_memmove(m, v1468-v1465, v1449, v1465)
	mBase = m.M
	v1473 = F___memset(m, v1449, int32(48), v1437-v1465)
	mBase = m.M
	v1476 = v1468
	goto L383
L390:
	;
	v1512 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1511))) = uint8(v1512)
	v1515 = v1511 + int32(1)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1517 = int32(2)
	if base.Ui32(int32(99)) < base.Ui32(v1516) {
		goto L398
	} else {
		goto L399
	}
L391:
	;
	v1500 = F_pg_ultoa_n(m, v1485, v1484)
	mBase = m.M
	if v1486 <= v1500 {
		goto L394
	} else {
		goto L395
	}
L392:
	;
	goto L393
L393:
	;
	v1496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1485<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1484))) = uint16(v1496)
	v1511 = v1480 + int32(3)
	goto L390
L394:
	;
	v1511 = v1484 + v1500
	goto L390
L395:
	;
	goto L396
L396:
	;
	v1503 = v1480 + int32(3)
	v1505 = F_memmove(m, v1503-v1500, v1484, v1500)
	mBase = m.M
	v1508 = F___memset(m, v1484, int32(48), v1486-v1500)
	mBase = m.M
	v1511 = v1503
	goto L390
L397:
	;
	v1543 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1542))) = uint8(v1543)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1553 = v1547 >> (uint(int32(31)) % 32)
	goto L406
L398:
	;
	v1531 = F_pg_ultoa_n(m, v1516, v1515)
	mBase = m.M
	if v1517 <= v1531 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	goto L400
L400:
	;
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516<<(uint(int32(1))%32))+uint32(_consts[1073]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1515))) = uint16(v1527)
	v1542 = v1511 + int32(3)
	goto L397
L401:
	;
	v1542 = v1515 + v1531
	goto L397
L402:
	;
	goto L403
L403:
	;
	v1534 = v1511 + int32(3)
	v1536 = F_memmove(m, v1534-v1531, v1515, v1531)
	mBase = m.M
	v1539 = F___memset(m, v1515, int32(48), v1517-v1531)
	mBase = m.M
	v1542 = v1534
	goto L397
L404:
	;
	v1666 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1665))) = uint8(v1666)
	v1668 = int32(1)
	v1669 = v1665 + v1668
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v1670 {
		goto L439
	} else {
		goto L440
	}
L405:
	;
	if l1 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L406:
	;
	v1557 = F_pg_ultostr_zeropad(m, v1542+int32(1), v1547^v1553-v1553, int32(2))
	mBase = m.M
	goto L405
L409:
	;
	v1665 = v1557
	goto L404
L410:
	;
	goto L411
L411:
	;
	v1562 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557))) = uint8(v1562)
	v1565 = l1 >> (uint(int32(31)) % 32)
	v1567 = l1 ^ v1565 - v1565
	v1569 = base.I32_div_s(v1567, int32(10))
	v1572 = v1569*int32(-10) + v1567
	if v1572 != 0 {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	v1582 = base.I32_div_s(v1567, int32(100))
	v1585 = v1582*int32(-10) + v1569
	v1586 = v1572 | v1585
	if v1586 == int32(0) {
		goto L417
	} else {
		goto L418
	}
L413:
	;
	v1574 = v1572 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+6)) = uint8(v1574)
	v1580 = v1557 + int32(7)
	goto L412
L414:
	;
	goto L415
L415:
	;
	v1580 = v1557 + int32(6)
	goto L412
L416:
	;
	v1596 = base.I32_div_s(v1567, int32(1000))
	v1599 = v1596*int32(-10) + v1582
	v1600 = v1586 | v1599
	if v1600 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L417:
	;
	v1594 = v1557 + int32(5)
	goto L416
L418:
	;
	goto L419
L419:
	;
	v1592 = v1585 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+5)) = uint8(v1592)
	v1594 = v1580
	goto L416
L420:
	;
	v1610 = base.I32_div_s(v1567, int32(10000))
	v1613 = v1610*int32(-10) + v1596
	v1614 = v1600 | v1613
	if v1614 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L421:
	;
	v1608 = v1557 + int32(4)
	goto L420
L422:
	;
	goto L423
L423:
	;
	v1606 = v1599 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+4)) = uint8(v1606)
	v1608 = v1594
	goto L420
L424:
	;
	v1624 = base.I32_div_s(v1567, int32(100000))
	v1627 = v1624*int32(-10) + v1610
	v1628 = v1614 | v1627
	if v1628 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L425:
	;
	v1622 = v1557 + int32(3)
	goto L424
L426:
	;
	goto L427
L427:
	;
	v1620 = v1613 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+3)) = uint8(v1620)
	v1622 = v1608
	goto L424
L428:
	;
	v1638 = base.I32_div_s(v1567, int32(1000000))
	v1641 = v1638*int32(-10) + v1624
	if v1628|v1641 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L429:
	;
	v1636 = v1557 + int32(2)
	goto L428
L430:
	;
	goto L431
L431:
	;
	v1634 = v1627 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+2)) = uint8(v1634)
	v1636 = v1622
	goto L428
L432:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1624+int32(9)) {
		goto L436
	} else {
		goto L437
	}
L433:
	;
	v1650 = v1557 + int32(1)
	goto L432
L434:
	;
	goto L435
L435:
	;
	v1648 = v1641 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1557)+1)) = uint8(v1648)
	v1650 = v1636
	goto L432
L436:
	;
	v1657 = F_pg_ultostr(m, v1557+int32(1), v1567)
	mBase = m.M
	v1658 = v1657
	goto L438
L437:
	;
	v1658 = v1650
	goto L438
L438:
	;
	v1665 = v1658
	goto L404
L439:
	;
	v1675 = v1670
	goto L441
L440:
	;
	v1675 = v1668 - v1670
	goto L441
L441:
	;
	v1676 = int32(4)
	if base.Ui32(int32(99)) < base.Ui32(v1675) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	if v19 == int32(0) {
		v1830 = v1701
		goto L1
	} else {
		goto L449
	}
L443:
	;
	v1690 = F_pg_ultoa_n(m, v1675, v1669)
	mBase = m.M
	if v1676 <= v1690 {
		goto L446
	} else {
		goto L447
	}
L444:
	;
	goto L443
L446:
	;
	v1701 = v1669 + v1690
	goto L442
L447:
	;
	goto L448
L448:
	;
	v1693 = v1665 + int32(5)
	v1695 = F_memmove(m, v1693-v1690, v1669, v1690)
	mBase = m.M
	v1698 = F___memset(m, v1669, int32(48), v1676-v1690)
	mBase = m.M
	v1701 = v1693
	goto L442
L449:
	;
	if l4 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(10)
	v1708 = F_pg_sprintf(m, v1701, int32(166845), v14)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L177
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v1768 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1701))) = uint8(v1768)
	if l3 <= int32(0) {
		goto L472
	} else {
		goto L473
	}
L453:
	;
	if v1701&int32(3) == int32(0) {
		v1733 = v1701
		goto L456
	} else {
		goto L457
	}
L454:
	;
	v1830 = v1766 + v1701
	goto L1
L455:
	;
	v1766 = v1758 - v1701
	goto L454
L456:
	;
	v1737 = v1733
	goto L465
L457:
	;
	v1717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1701))))
	if v1717 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v1766 = int32(0)
	goto L454
L459:
	;
	goto L460
L460:
	;
	v1722 = v1701
	goto L461
L461:
	;
	v1726 = v1722 + int32(1)
	if v1726&int32(3) == int32(0) {
		v1733 = v1726
		goto L456
	} else {
		goto L463
	}
L462:
	;
	v1758 = v1726
	goto L455
L463:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726))))
	if v1731 != 0 {
		v1722 = v1726
		goto L461
	} else {
		goto L464
	}
L464:
	;
	goto L462
L465:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1737)))
	v1746 = int32(-2139062144)
	if (int32(16843008)-v1743|v1743)&v1746 == v1746 {
		v1737 = v1737 + int32(4)
		goto L465
	} else {
		goto L467
	}
L466:
	;
	v1752 = v1737
	goto L468
L467:
	;
	goto L466
L468:
	;
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752))))
	if v1756 != 0 {
		v1752 = v1752 + int32(1)
		goto L468
	} else {
		goto L470
	}
L469:
	;
	v1758 = v1752
	goto L455
L470:
	;
	goto L469
L471:
	;
	v1830 = v1819
	goto L1
L472:
	;
	v1779 = int32(43)
	goto L474
L473:
	;
	v1779 = int32(45)
	goto L474
L474:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1701+int32(1)))) = uint8(v1779)
	v1782 = l3 >> (uint(int32(31)) % 32)
	v1784 = l3 ^ v1782 - v1782
	v1786 = base.I32_div_s(v1784, int32(3600))
	v1787 = int32(-60)
	v1790 = base.I32_div_s(v1784, int32(60))
	v1791 = v1786*v1787 + v1790
	v1793 = v1701 + int32(2)
	v1796 = v1790*v1787 + v1784
	if v1796 != 0 {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L471
L476:
	;
	v1813 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1811))) = uint8(v1813)
	v1818 = F_pg_ultostr_zeropad(m, v1811+int32(1), v1812, int32(2))
	mBase = m.M
	v1819 = v1818
	goto L475
L477:
	;
	v1797 = int32(2)
	v1798 = F_pg_ultostr_zeropad(m, v1793, v1786, v1797)
	mBase = m.M
	v1799 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v1798))) = uint8(v1799)
	v1804 = F_pg_ultostr_zeropad(m, v1798+int32(1), v1791, v1797)
	mBase = m.M
	v1811 = v1804
	v1812 = v1796
	goto L476
L478:
	;
	goto L479
L479:
	;
	v1806 = F_pg_ultostr_zeropad(m, v1793, v1786, int32(2))
	mBase = m.M
	if l5 == int32(4) {
		v1811 = v1806
		v1812 = v1791
		goto L476
	} else {
		goto L480
	}
L480:
	;
	if v1791 == int32(0) {
		v1819 = v1806
		goto L475
	} else {
		goto L481
	}
L481:
	;
	v1811 = v1806
	v1812 = v1791
	goto L476
L482:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1830)+2)) = uint8(v1835)
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1830))) = uint16(v1838)
	v1842 = v1830 + int32(3)
	goto L484
L483:
	;
	v1842 = v1830
	goto L484
L484:
	;
	v1843 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1842))) = uint8(v1843)
	m.G0 = v14 + int32(48)
	return
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
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
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
					F_errmsg(m, int32(152288), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(479368), int32(560), int32(306977))
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
				F_errmsg(m, int32(152288), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(479368), int32(560), int32(306977))
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
			v15 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1056])))
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+8)) = uint16(v15)
			v18 = *(*int64)(unsafe.Add(mBase, _consts[1057]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v18
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1058])))
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v21)
			v24 = *(*int64)(unsafe.Add(mBase, _consts[1059]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v24
		}
	} else {
		v37 = v7 + int32(2483589)
		v38 = int32(146097)
		v39 = base.I32_div_u_s(v37, v38)
		v40 = int32(3)
		v46 = int32(2)
		v51 = base.I32_div_u_s((v39*int32(1073595727)+v37)<<(uint(v46)%32)|v40, v38)
		v54 = v7 + int32(2451545) + v39*v40 + v51 + int32(32104)
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
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(152)))) = v74 + v56<<(uint(int32(2))%32) - int32(4800)
		v82 = v72 + int32(123)
		v86 = int32(base.Ui32(v82*int32(2141)) >> (uint(int32(16)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(144)))) = v82 - int32(base.Ui32(v86*int32(7834))>>(uint(int32(8))%32))
		v96 = base.I32_rem_u_s(v86+int32(10), int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(148)))) = v96 + int32(1)
		v101 = v5 + int32(132)
		v103 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
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
			v131 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
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
			v187 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
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
			v225 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
			*(*uint8)(unsafe.Add(mBase, uint32(v220)+2)) = uint8(v225)
			v228 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v19 = F_Int64GetDatum(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_DirectFunctionCall2Coll(m, int32(1282), int32(0), v19, v3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		}
	} else {
		if v7 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v19 = F_Int64GetDatum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_DirectFunctionCall2Coll(m, int32(1282), int32(0), v19, v3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		} else {
			if int32(106751983) <= v7 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(227099), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(479368), int32(658), int32(29371))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				v18 = base.I64_extend_i32_s(v7) * int64(86400000000)
				v19 = F_Int64GetDatum(m, v18)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_DirectFunctionCall2Coll(m, int32(1282), int32(0), v19, v3)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
	}
}
func F_date_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(1279)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(1280)
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
						F_errmsg(m, int32(227099), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(479368), int32(658), int32(29371))
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
						F_errmsg(m, int32(227099), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(479368), int32(721), int32(29341))
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
				v45 = v4 + int32(2483589)
				v46 = int32(146097)
				v47 = base.I32_div_u_s(v45, v46)
				v48 = int32(3)
				v54 = int32(2)
				v59 = base.I32_div_u_s((v47*int32(1073595727)+v45)<<(uint(v54)%32)|v48, v46)
				v62 = v4 + int32(2451545) + v47*v48 + v59 + int32(32104)
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
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v82 + v64<<(uint(int32(2))%32) - int32(4800)
				v90 = v80 + int32(123)
				v94 = int32(base.Ui32(v90*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v90 - int32(base.Ui32(v94*int32(7834))>>(uint(int32(8))%32))
				v104 = base.I32_rem_u_s(v94+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v104 + int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
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
							F_errmsg(m, int32(227099), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(479368), int32(757), int32(29341))
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
