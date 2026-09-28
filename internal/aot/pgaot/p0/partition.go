package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v639 int32
	_ = v639
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v737 int32
	_ = v737
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int64
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1010 int32
	_ = v1010
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1076 int32
	_ = v1076
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1281 int32
	_ = v1281
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1418 int32
	_ = v1418
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1489 int32
	_ = v1489
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1571 int32
	_ = v1571
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1670 int32
	_ = v1670
	var v1679 int32
	_ = v1679
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1788 int32
	_ = v1788
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1900 int32
	_ = v1900
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
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
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2123 int32
	_ = v2123
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2194 int32
	_ = v2194
	var v2212 int32
	_ = v2212
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2279 int32
	_ = v2279
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
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
	var v2445 int32
	_ = v2445
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2464 int32
	_ = v2464
	var v2469 int32
	_ = v2469
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	v6 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(384)
	m.G0 = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v42 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v60 = int32(_a_F_ExecFindPartition_0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+131)))
	if v67 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v56 = v42
	v58 = v42 + int32(4)
	v59 = v45
	goto L1
L3:
	;
	goto L4
L4:
	;
	v46 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v51 = v46 + int32(4)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v53 != 0 {
		v56 = v53
		v58 = v51
		v59 = v52
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v54 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v56 = v54
	v58 = v51
	v59 = v52
	goto L1
L9:
	;
	v71 = F_ExecPartitionCheck(m, l1, l3, l4, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v73 == int32(0) {
		v2464 = v6
		v2469 = v39
		v2485 = v58
		v2487 = v59
		v2488 = v61
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2485))) = v2487
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v2488
	m.G0 = v2469 + int32(384)
	return v2464
L14:
	;
	v76 = l0
	v77 = l1
	v78 = l2
	v79 = l3
	v80 = l4
	v85 = v73
	v86 = v39
	v90 = l3
	v98 = v6
	v102 = v58
	v103 = v41
	v104 = v59
	v105 = v61
	goto L15
L15:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[1]))
	if v113 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v2419 == int32(0) {
		v2464 = v2402
		v2469 = v2407
		v2485 = v2423
		v2487 = v2425
		v2488 = v2426
		goto L13
	} else {
		goto L409
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v90
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v121 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	if int32(0) < v132 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v130 = v128
	goto L21
L23:
	;
	if v119 != 0 {
		v127 = v119
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v125 = v119
	goto L25
L25:
	;
	if v125 != 0 {
		v127 = v125
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v122 = F_ExecPrepareExprList(m, v121, v80)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v122
	v125 = v122
	goto L25
L28:
	;
	v130 = int32(0)
	goto L21
L29:
	;
	v1106 = v1076 << (uint(int32(2)) % 32)
	v1107 = v85 + v1106
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+24))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109+v1076))))
	if v1111 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+24)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v285)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v285)+20)) = v556
	v1076 = v558
	goto L29
L31:
	;
	v706 = F_RelationGetPartitionKey(m, v117)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L5
	} else {
		goto L113
	}
L32:
	;
	if int32(0) <= v639 {
		v1076 = v639
		goto L29
	} else {
		goto L112
	}
L33:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	if v556 != v626 {
		goto L30
	} else {
		goto L111
	}
L34:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v286)+32))
	v639 = v625
	goto L32
L35:
	;
	if int32(0) <= v558 {
		goto L33
	} else {
		goto L110
	}
L36:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v285)+28))
	if v485 < int32(16) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	v409 = v339
	goto L93
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L90
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L87
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L84
	}
L41:
	;
	v141 = v131
	v142 = v130
	v143 = int32(0)
	goto L44
L42:
	;
	v251 = v131
	v252 = v130
	v254 = v132
	goto L43
L43:
	;
	if v252 != 0 {
		goto L39
	} else {
		goto L64
	}
L44:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v172+v143<<(uint(int32(1))%32)))))
	if v176 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v251 = v243
	v252 = v225
	v254 = v244
	goto L43
L46:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86-int32(-64)+v143))) = uint8(v233)
	*(*int64)(unsafe.Add(mBase, uint32(v86+int32(96)+v143<<(uint(int32(3))%32)))) = v229
	v242 = v143 + int32(1)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+4)))
	if v242 < v244 {
		v141 = v243
		v142 = v225
		v143 = v242
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+6)))
	if v177 < v176 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v142 == int32(0) {
		goto L40
	} else {
		goto L54
	}
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	m.T0[v180].(func(*base.Module, int32, int32))(m, v90, v176)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v184 = v176 - int32(1)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v187+v184<<(uint(int32(3))%32))))
	v224 = v184 + v185
	v225 = v142
	v229 = v191
	goto L46
L53:
	;
	goto L52
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v80)+152))
	if v195 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v198 = F_MakePerTupleExprContext(m, v80)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	v200 = v195
	goto L57
L57:
	;
	v201 = int32(_a_F_ExecFindPartition_0)
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v204
	v207 = v86 + int32(368)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	v209 = m.T0[v208].(func(*base.Module, int32, int32, int32) int64)(m, v194, v200, v207)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v200 = v198
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v202
	v214 = v142 + int32(4)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if base.Ui32(v214) < base.Ui32(v217+v218<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v223 = v214
	goto L62
L61:
	;
	v223 = int32(0)
	goto L62
L62:
	;
	v224 = v207
	v225 = v223
	v229 = v209
	goto L46
L63:
	;
	goto L45
L64:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v282 == int32(0) {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	switch v287 - int32(104) {
	case 0:
		goto L66
	default:
		goto L38
	case 4:
		goto L68
	case 10:
		goto L67
	}
L66:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v350 = F_compute_partition_hash_value(m, v254, v344, v345, v86+int32(96), v86-int32(-64))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L83
	}
L67:
	;
	v339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+368)) = uint8(v339)
	if v339 < v254 {
		goto L37
	} else {
		goto L82
	}
L68:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+64)))
	if v290 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v286)+28))
	if v293 == int32(-1) {
		goto L34
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v285)+28))
	if int32(16) <= v296 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v639 = v293
	goto L32
L73:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v336+v304)))
	v639 = v338
	goto L32
L74:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v304 = v302 << (uint(int32(2)) % 32)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304+v305)))
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v86)+96))
	v310 = F_FunctionCall2Coll(m, v299, v301, v308, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v316 = int32(-1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v86)+96))
	v322 = F_partition_list_bsearch(m, v317, v318, v286, v319, v86+int32(368))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	if base.I32_wrap_i64(v310) == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v322 < int32(0) {
		v556 = v322
		v558 = v316
		goto L35
	} else {
		goto L80
	}
L80:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+368)))
	if v326&int32(1) == int32(0) {
		v556 = v322
		v558 = v316
		goto L35
	} else {
		goto L81
	}
L81:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331+v322<<(uint(int32(2))%32))))
	v556 = v322
	v558 = v335
	goto L35
L82:
	;
	goto L36
L83:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v353 = int64(*(*int32)(unsafe.Add(mBase, uint32(v286)+20)))
	v354 = base.I64_rem_u_s(v350, v353)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v352+base.I32_wrap_i64(v354)<<(uint(int32(2))%32))))
	v639 = v359
	goto L32
L84:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1445), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1456), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = v390
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_4), v86+int32(48))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1679), int32(_a_F_ExecFindPartition_5))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86-int32(-64)+v409))))
	if v441 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v447 = int32(-1)
	v556 = v447
	v558 = v447
	goto L35
L95:
	;
	v445 = v409 + int32(1)
	if v254 != v445 {
		v409 = v445
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L36
L99:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v539 = int32(*(*int16)(unsafe.Add(mBase, uint32(v251)+4)))
	v544 = F_partition_range_datum_bsearch(m, v537, v538, v286, v539, v86+int32(96), v86+int32(368))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L5
	} else {
		goto L109
	}
L100:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v492 = v490 << (uint(int32(2)) % 32)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v492+v493)))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v496+v492)))
	v501 = F_partition_rbound_datum_cmp(m, v488, v489, v495, v498, v86+int32(96), v254)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v501 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v505+v492)+4))
	v639 = v507
	goto L32
L103:
	;
	goto L104
L104:
	;
	if int32(0) <= v501 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v511 = v490 + int32(1)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v512 <= v511 {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v517 = v511 << (uint(int32(2)) % 32)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v517+v518)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v521+v517)))
	v526 = int32(*(*int16)(unsafe.Add(mBase, uint32(v251)+4)))
	v527 = F_partition_rbound_datum_cmp(m, v514, v515, v520, v523, v86+int32(96), v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v527 <= int32(0) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v531+v517)))
	v639 = v533
	goto L32
L109:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546+v544<<(uint(int32(2))%32))+4))
	v556 = v544
	v558 = v550
	goto L35
L110:
	;
	goto L34
L111:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v285)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+28)) = v628 + int32(1)
	v1076 = v558
	goto L29
L112:
	;
	goto L31
L113:
	;
	v708 = int32(*(*int16)(unsafe.Add(mBase, uint32(v706)+4)))
	v709 = int32(0)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v117)+56))
	v713 = F_check_enable_rls(m, v710, v709, int32(1))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L5
	} else {
		goto L115
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L5
	} else {
		goto L164
	}
L115:
	;
	if v713 == int32(2) {
		v1010 = v709
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v720 = F_pg_class_aclcheck(m, v710, v718, int64(2))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v722 = int32(0)
	if base.B2i32(v720 == v722)|base.B2i32(v708 <= v722) == v722 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v737 = int32(0)
	goto L121
L119:
	;
	goto L120
L120:
	;
	v818 = v86 + int32(368)
	F_initStringInfo(m, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L5
	} else {
		goto L127
	}
L121:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v706)+8))
	v770 = int32(*(*int16)(unsafe.Add(mBase, uint32(v766+v737<<(uint(int32(1))%32)))))
	if v770 == int32(0) {
		v1010 = v709
		goto L114
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v776 = F_pg_attribute_aclcheck(m, v710, v770, v774, int64(2))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v776 != 0 {
		v1010 = v709
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v779 = v737 + int32(1)
	if v779 != v708 {
		v737 = v779
		goto L121
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v824 = F_pg_get_partkeydef_worker(m, v710, int32(7), int32(1), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v824
	F_appendStringInfo(m, v818, int32(_a_F_ExecFindPartition_6), v86+int32(32))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v708 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_appendStringInfoChar(m, v86+int32(368), int32(41))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L163
	}
L131:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+64)))
	if v835 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v706)+32))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	F_getTypeOutputInfo(m, v839, v86+int32(364), v86+int32(363))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	v850 = int32(_a_F_ExecFindPartition_7)
	goto L134
L134:
	;
	v851 = F_strlen(m, v850)
	mBase = m.M
	if int32(65) <= v851 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v86)+364))
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v86)+96))
	v848 = F_OidOutputFunctionCall(m, v846, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v850 = v848
	goto L134
L137:
	;
	v869 = int32(1)
	if v708 == v869 {
		goto L130
	} else {
		goto L145
	}
L138:
	;
	v855 = v86 + int32(368)
	v857 = F_pg_mbcliplen(m, v850, v851, int32(64))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_appendBinaryStringInfo(m, v86+int32(368), v850, v851)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L5
	} else {
		goto L144
	}
L141:
	;
	F_appendBinaryStringInfo(m, v855, v850, v857)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	F_appendStringInfoString(m, v855, int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	goto L137
L145:
	;
	v879 = v869
	goto L146
L146:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86-int32(-64)+v879))))
	if v912 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L130
L148:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v706)+32))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v915+v879<<(uint(int32(2))%32))))
	F_getTypeOutputInfo(m, v919, v86+int32(364), v86+int32(363))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	v935 = int32(_a_F_ExecFindPartition_7)
	goto L150
L150:
	;
	v937 = v86 + int32(368)
	F_appendStringInfoString(m, v937, int32(_a_F_ExecFindPartition_9))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v86)+364))
	v932 = *(*int64)(unsafe.Add(mBase, uint32(v86+int32(96)+v879<<(uint(int32(3))%32))))
	v933 = F_OidOutputFunctionCall(m, v926, v932)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v935 = v933
	goto L150
L153:
	;
	v941 = F_strlen(m, v935)
	mBase = m.M
	if v941 <= int32(64) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v958 = v879 + int32(1)
	if v958 != v708 {
		v879 = v958
		goto L146
	} else {
		goto L162
	}
L155:
	;
	F_appendBinaryStringInfo(m, v937, v935, v941)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L5
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v947 = v86 + int32(368)
	v949 = F_pg_mbcliplen(m, v935, v941, int32(64))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L5
	} else {
		goto L159
	}
L158:
	;
	goto L154
L159:
	;
	F_appendBinaryStringInfo(m, v947, v935, v949)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	F_appendStringInfoString(m, v947, int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	goto L154
L162:
	;
	goto L147
L163:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v86)+368))
	v1010 = v1001
	goto L114
L164:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v1045 + int32(4)
	F_errmsg(m, int32(_a_F_ExecFindPartition_10), v86+int32(16))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	if v1010 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v1010
	v1056 = F_errdetail(m, int32(_a_F_ExecFindPartition_11), v86)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L5
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F_errtable(m, v117)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L5
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(338), int32(_a_F_ExecFindPartition_12))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+32))
	if v2434 == v1076 {
		goto L399
	} else {
		goto L400
	}
L174:
	;
	if int32(0) <= v1108 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if int32(0) <= v1108 {
		goto L389
	} else {
		goto L390
	}
L177:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1116+v1108<<(uint(int32(2))%32))))
	v2397 = v76
	v2398 = v77
	v2399 = v78
	v2400 = v79
	v2401 = v80
	v2402 = v1120
	v2406 = int32(0)
	v2407 = v86
	v2411 = v90
	v2419 = v98
	v2423 = v102
	v2424 = v103
	v2425 = v104
	v2426 = v105
	goto L173
L178:
	;
	goto L179
L179:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1122+v1106)))
	v1127 = F_ExecLookupResultRelByOid(m, v76, v1124, int32(1), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	if v1127 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v1130 != 0 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	v1141 = int32(0)
	v1146 = m.G0
	v1148 = v1146 - int32(16)
	m.G0 = v1148
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+8))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1151+v1076<<(uint(int32(2))%32))))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v76)+116))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+8))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+4))
	v1160 = int32(_a_F_ExecFindPartition_0)
	v1161 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v78)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1163
	v1166 = F_table_open(m, v1155, int32(3))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L5
	} else {
		goto L190
	}
L184:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+128))
	v1133 = v1131
	goto L186
L185:
	;
	v1133 = int32(0)
	goto L186
L186:
	;
	F_CheckValidResultRel(m, v1127, int32(3), v1133, int32(0))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L5
	} else {
		goto L187
	}
L187:
	;
	F_ExecInitRoutingInfo(m, v76, v80, v78, v85, v1127, v1076, int32(1))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v2397 = v76
	v2398 = v77
	v2399 = v78
	v2400 = v79
	v2401 = v80
	v2402 = v1127
	v2406 = int32(0)
	v2407 = v86
	v2411 = v90
	v2419 = v98
	v2423 = v102
	v2424 = v103
	v2425 = v104
	v2426 = v105
	goto L173
L189:
	;
	v2397 = v76
	v2398 = v77
	v2399 = v78
	v2400 = v79
	v2401 = v80
	v2402 = v1169
	v2406 = int32(0)
	v2407 = v86
	v2411 = v90
	v2419 = v98
	v2423 = v102
	v2424 = v103
	v2425 = v104
	v2426 = v105
	goto L173
L190:
	;
	v1169 = F_palloc0(m, int32(216))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169))) = int32(394)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v80)+132))
	F_InitResultRelInfo(m, v1169, v1166, int32(0), v77, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if v1156 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+128))
	v1180 = v1178
	goto L195
L194:
	;
	v1180 = int32(0)
	goto L195
L195:
	;
	F_CheckValidResultRel(m, v1169, int32(3), v1180, int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184)+116)))
	if v1185 != int32(1) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if v1156 != 0 {
		goto L212
	} else {
		goto L213
	}
L198:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+16))
	if v1188 != 0 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	if v1156 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+128))
	v1193 = base.B2i32(v1189 != int32(0))
	goto L202
L201:
	;
	v1193 = int32(0)
	goto L202
L202:
	;
	F_ExecOpenIndices(m, v1169, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	goto L197
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L5
	} else {
		goto L385
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1161
	m.G0 = v1148 + int32(16)
	goto L189
L206:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v80)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v2149
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	v2152 = F_lappend(m, v2151, v1169)
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L5
	} else {
		goto L349
	}
L207:
	;
	if v1995-v1974 != v1977 {
		goto L204
	} else {
		goto L317
	}
L208:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+4))
	v1974 = v1890
	v1975 = v1891
	v1977 = v1919
	v1995 = v1958
	goto L207
L209:
	;
	v1956 = int32(0)
	v1974 = v1935
	v1975 = v1956
	v1977 = v1938
	v1995 = v1956
	goto L207
L210:
	;
	F_list_free(m, v1884)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L5
	} else {
		goto L310
	}
L211:
	;
	v1706 = int32(0)
	if base.B2i32(v1663 == v1706)|base.B2i32(v1670 == v1706) != 0 {
		v1883 = v1662
		v1884 = v1663
		v1890 = v1706
		v1891 = v1670
		v1900 = v1679
		goto L210
	} else {
		goto L294
	}
L212:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+96))
	if v1196 != 0 {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	goto L214
L214:
	;
	F_ExecInitRoutingInfo(m, v76, v80, v78, v85, v1169, v1076, int32(0))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L5
	} else {
		goto L292
	}
L215:
	;
	F_ExecInitRoutingInfo(m, v76, v80, v78, v85, v1169, v1076, int32(0))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L5
	} else {
		goto L237
	}
L216:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+72))
	v1398 = F_map_variable_attnos(m, v1371, v1159, v1369, v1395, v1148+int32(15))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L5
	} else {
		goto L235
	}
L217:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v1356 = F_build_attrmap_by_name(m, v1353, v1354, int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L5
	} else {
		goto L234
	}
L218:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+12))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1197)))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v1202 = F_build_attrmap_by_name(m, v1199, v1200, int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L5
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+108))
	if v1311 == int32(0) {
		v1418 = v1141
		goto L215
	} else {
		goto L233
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+120)) = v1281
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+116)) = v1208
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+108))
	if v1304 == int32(0) {
		v1418 = v1202
		goto L215
	} else {
		goto L231
	}
L222:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+72))
	v1208 = F_map_variable_attnos(m, v1198, v1159, v1202, v1205, v1148+int32(15))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	if v1208 == int32(0) {
		v1281 = v1141
		goto L221
	} else {
		goto L224
	}
L224:
	;
	v1212 = int32(0)
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1213 <= v1212 {
		v1281 = v1141
		goto L221
	} else {
		goto L225
	}
L225:
	;
	v1229 = v1212
	v1231 = v1141
	goto L226
L226:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+12))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1252+v1229<<(uint(int32(2))%32))))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+16))
	v1258 = F_ExecInitQual(m, v1257, v76)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L5
	} else {
		goto L228
	}
L227:
	;
	v1281 = v1260
	goto L221
L228:
	;
	v1260 = F_lappend(m, v1231, v1258)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	v1263 = v1229 + int32(1)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1263 < v1264 {
		v1229 = v1263
		v1231 = v1260
		goto L226
	} else {
		goto L230
	}
L230:
	;
	goto L227
L231:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+12))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	if v1202 == int32(0) {
		v1352 = v1308
		goto L217
	} else {
		goto L232
	}
L232:
	;
	v1369 = v1202
	v1371 = v1308
	goto L216
L233:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+12))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1314)))
	v1352 = v1315
	goto L217
L234:
	;
	v1369 = v1356
	v1371 = v1352
	goto L216
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+148)) = v1398
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v76)+64))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v76)+60))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v1404 = F_ExecBuildProjectionInfo(m, v1398, v1401, v1402, v76, v1403)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+152)) = v1404
	v1418 = v1369
	goto L215
L237:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+128))
	if v1446 == int32(0) {
		v2123 = v1418
		goto L206
	} else {
		goto L238
	}
L238:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v76)+64))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v77)+156))
	if v1451 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1935 = int32(0)
	v1938 = v1141
	goto L209
L240:
	;
	goto L241
L241:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	if v1455 <= int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1458 = int32(0)
	v1883 = v1458
	v1884 = v1458
	v1890 = v1458
	v1891 = v1141
	v1900 = v1141
	goto L210
L243:
	;
	goto L244
L244:
	;
	v1461 = int32(0)
	v1472 = v1461
	v1473 = v1461
	v1479 = v1461
	v1480 = v1141
	v1489 = v1141
	goto L245
L245:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+16))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1500+v1479<<(uint(int32(2))%32))))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+56))
	v1506 = F_get_partition_ancestors(m, v1505)
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L5
	} else {
		goto L249
	}
L246:
	;
	goto L211
L247:
	;
	F_list_free(m, v1506)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L5
	} else {
		goto L290
	}
L248:
	;
	v1652 = F_lappend_int(m, v1473, v1479)
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L5
	} else {
		goto L289
	}
L249:
	;
	if v1506 == int32(0) {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+12))
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)))
	v1512 = int32(0)
	if v1489 == v1512 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	if v1550 != 0 {
		goto L248
	} else {
		goto L264
	}
L252:
	;
	v1550 = int32(0)
	goto L251
L253:
	;
	goto L254
L254:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+4))
	if v1518 <= int32(0) {
		v1544 = v1512
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1550 = v1544
	goto L251
L256:
	;
	v1521 = int32(0)
	if v1521 < v1518 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1524 = v1518
	goto L259
L258:
	;
	v1524 = v1521
	goto L259
L259:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+12))
	v1527 = int32(0)
	goto L260
L260:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1525+v1527<<(uint(int32(2))%32))))
	v1536 = base.B2i32(v1535 == v1511)
	if v1535 == v1511 {
		v1544 = v1536
		goto L255
	} else {
		goto L262
	}
L261:
	;
	v1544 = v1536
	goto L255
L262:
	;
	v1538 = v1527 + int32(1)
	if v1538 != v1524 {
		v1527 = v1538
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v77)+156))
	if v1551 == int32(0) {
		v1662 = v1472
		v1663 = v1473
		v1670 = v1480
		v1679 = v1489
		goto L247
	} else {
		goto L265
	}
L265:
	;
	v1554 = int32(0)
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+4))
	if v1555 <= v1554 {
		v1662 = v1472
		v1663 = v1473
		v1670 = v1480
		v1679 = v1489
		goto L247
	} else {
		goto L266
	}
L266:
	;
	v1571 = v1554
	goto L267
L267:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+12))
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1594+v1571<<(uint(int32(2))%32))))
	v1599 = int32(0)
	if v1506 == v1599 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+12))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)))
	v1646 = F_lappend_oid(m, v1489, v1645)
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L5
	} else {
		goto L286
	}
L269:
	;
	if v1637 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L270:
	;
	v1637 = int32(0)
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+4))
	if v1605 <= int32(0) {
		v1631 = v1599
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1637 = v1631
	goto L269
L274:
	;
	v1608 = int32(0)
	if v1608 < v1605 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1611 = v1605
	goto L277
L276:
	;
	v1611 = v1608
	goto L277
L277:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+12))
	v1614 = int32(0)
	goto L278
L278:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1612+v1614<<(uint(int32(2))%32))))
	v1623 = base.B2i32(v1622 == v1598)
	if v1622 == v1598 {
		v1631 = v1623
		goto L273
	} else {
		goto L280
	}
L279:
	;
	v1631 = v1623
	goto L273
L280:
	;
	v1625 = v1614 + int32(1)
	if v1625 != v1611 {
		v1614 = v1625
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v1641 = v1571 + int32(1)
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+4))
	if v1641 < v1642 {
		v1571 = v1641
		goto L267
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	goto L268
L285:
	;
	v1662 = v1472
	v1663 = v1473
	v1670 = v1480
	v1679 = v1489
	goto L247
L286:
	;
	v1648 = F_lappend_oid(m, v1480, v1505)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v1650 = F_lappend_int(m, v1472, v1479)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	v1662 = v1650
	v1663 = v1473
	v1670 = v1648
	v1679 = v1646
	goto L247
L289:
	;
	v1662 = v1472
	v1663 = v1652
	v1670 = v1480
	v1679 = v1489
	goto L247
L290:
	;
	v1693 = v1479 + int32(1)
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+12))
	if v1693 < v1694 {
		v1472 = v1662
		v1473 = v1663
		v1479 = v1693
		v1480 = v1670
		v1489 = v1679
		goto L245
	} else {
		goto L291
	}
L291:
	;
	goto L246
L292:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v80)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1700
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	v1703 = F_lappend(m, v1702, v1169)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+80)) = v1703
	goto L205
L294:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+4))
	if v1712 <= int32(0) {
		v1883 = v1662
		v1884 = v1663
		v1890 = v1706
		v1891 = v1670
		v1900 = v1679
		goto L210
	} else {
		goto L295
	}
L295:
	;
	v1731 = v1706
	v1732 = v1670
	v1733 = int32(0)
	goto L296
L296:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+12))
	v1755 = int32(2)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1754+v1733<<(uint(v1755)%32))))
	v1760 = v1758 << (uint(v1755) % 32)
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+20))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1760+v1761)))
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1763)+118)))
	if base.B2i32(v1662 == int32(0))|base.B2i32(v1764 != int32(1)) != 0 {
		v1850 = v1731
		v1851 = v1732
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1883 = v1662
	v1884 = v1663
	v1890 = v1850
	v1891 = v1851
	v1900 = v1679
	goto L210
L298:
	;
	v1872 = v1733 + int32(1)
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1663)+4))
	if v1872 < v1873 {
		v1731 = v1850
		v1732 = v1851
		v1733 = v1872
		goto L296
	} else {
		goto L309
	}
L299:
	;
	v1768 = int32(0)
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+4))
	if v1769 <= v1768 {
		v1850 = v1731
		v1851 = v1732
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+16))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1760)))
	v1788 = v1768
	goto L301
L301:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+16))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+12))
	v1813 = int32(2)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1788<<(uint(v1813)%32))))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1811+v1816<<(uint(v1813)%32))))
	v1821 = F_IsIndexCompatibleAsArbiter(m, v1820, v1774)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L5
	} else {
		goto L303
	}
L302:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+192))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	v1833 = F_lappend_oid(m, v1732, v1832)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L5
	} else {
		goto L308
	}
L303:
	;
	if v1821 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1826 = v1788 + int32(1)
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1662)+4))
	if v1826 < v1827 {
		v1788 = v1826
		goto L301
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	goto L302
L307:
	;
	v1850 = v1731
	v1851 = v1732
	goto L298
L308:
	;
	v1850 = v1731 + int32(1)
	v1851 = v1833
	goto L298
L309:
	;
	goto L297
L310:
	;
	F_list_free(m, v1883)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	F_list_free(m, v1900)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v77)+156))
	if v1917 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+4))
	v1919 = v1918
	goto L315
L314:
	;
	v1919 = v1141
	goto L315
L315:
	;
	if v1891 != 0 {
		goto L208
	} else {
		goto L316
	}
L316:
	;
	v1935 = v1890
	v1938 = v1919
	goto L209
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+156)) = v1975
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+128))
	if v1999&int32(-2) != int32(2) {
		v2123 = v1418
		goto L206
	} else {
		goto L318
	}
L318:
	;
	v2005 = F_palloc0(m, int32(24))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005))) = int32(392)
	v2009 = F_ExecGetRootToChildMap(m, v1169, v80)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+160)) = v2005
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v77)+160))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+16)) = v2013
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+8))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v2019 = F_table_slot_create(m, v2015, v2016+int32(104))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+4)) = v2019
	if v2009 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+20)) = v2110
	v2123 = v2106
	goto L206
L323:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v77)+160))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+8)) = v2025
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v77)+160))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2027)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+12)) = v2028
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v77)+160))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+20))
	v2106 = v1418
	v2110 = v2031
	goto L322
L324:
	;
	goto L325
L325:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+128))
	if v2032 == int32(2) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+140))
	v2036 = F_copyObjectImpl(m, v2035)
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L5
	} else {
		goto L329
	}
L327:
	;
	v2074 = v1418
	goto L328
L328:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+148))
	if v2078 == int32(0) {
		v2123 = v2074
		goto L206
	} else {
		goto L340
	}
L329:
	;
	if v1418 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v2044 = F_build_attrmap_by_name(m, v2041, v2042, int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L5
	} else {
		goto L333
	}
L331:
	;
	v2046 = v1418
	goto L332
L332:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+72))
	v2050 = v1148 + int32(15)
	v2051 = F_map_variable_attnos(m, v2036, int32(-1), v2046, v2048, v2050)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L5
	} else {
		goto L334
	}
L333:
	;
	v2046 = v2044
	goto L332
L334:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+72))
	v2055 = F_map_variable_attnos(m, v2051, v1159, v2046, v2054, v2050)
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+144))
	v2058 = F_ExecGetChildToRootMap(m, v1169)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2058)+8))
	v2061 = F_adjust_partition_colnos_using_map(m, v2057, v2060)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v2066 = F_table_slot_create(m, v1166, v2063+int32(104))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+8)) = v2066
	v2070 = F_ExecBuildUpdateProjection(m, v2055, int32(1), v2061, v1450, v1449, v2066, v76)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+12)) = v2070
	v2074 = v2046
	goto L328
L340:
	;
	if v2074 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2087 = v2074
	v2088 = v2078
	goto L343
L342:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v2084 = F_build_attrmap_by_name(m, v2081, v2082, int32(0))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L5
	} else {
		goto L344
	}
L343:
	;
	v2089 = F_copyObjectImpl(m, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L5
	} else {
		goto L345
	}
L344:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+148))
	v2087 = v2084
	v2088 = v2086
	goto L343
L345:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+72))
	v2095 = v1148 + int32(15)
	v2096 = F_map_variable_attnos(m, v2089, int32(-1), v2087, v2093, v2095)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L5
	} else {
		goto L346
	}
L346:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+72))
	v2100 = F_map_variable_attnos(m, v2096, v1159, v2087, v2099, v2095)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L5
	} else {
		goto L347
	}
L347:
	;
	v2102 = F_ExecInitQual(m, v2100, v76)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	v2106 = v2087
	v2110 = v2102
	goto L322
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+80)) = v2152
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+72))
	if v2155 != int32(5) {
		goto L205
	} else {
		goto L350
	}
L350:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v76)+64))
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+160))
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2159)+12))
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2160)))
	if v2123 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v2167 = F_build_attrmap_by_name(m, v2164, v2165, int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L5
	} else {
		goto L354
	}
L352:
	;
	v2169 = v2123
	goto L353
L353:
	;
	v2170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169)+48)))
	if v2170 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v2169 = v2167
	goto L353
L355:
	;
	F_ExecInitMergeTupleSlots(m, v76, v1169)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L5
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+164))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+12))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2177)))
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2179)+72))
	v2183 = F_map_variable_attnos(m, v2178, v1159, v2169, v2180, v1148+int32(15))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L5
	} else {
		goto L359
	}
L358:
	;
	goto L357
L359:
	;
	v2185 = F_ExecInitQual(m, v2183, v76)
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1169)+176)) = v2185
	if v2161 == int32(0) {
		goto L205
	} else {
		goto L361
	}
L361:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+4))
	if v2190 <= int32(0) {
		goto L205
	} else {
		goto L362
	}
L362:
	;
	v2194 = v1169 + int32(164)
	v2212 = int32(0)
	goto L363
L363:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+12))
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2231+v2212<<(uint(int32(2))%32))))
	v2236 = F_copyObjectImpl(m, v2235)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L5
	} else {
		goto L365
	}
L364:
	;
	goto L205
L365:
	;
	v2239 = F_palloc0(m, int32(16))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L5
	} else {
		goto L366
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2239)+4)) = v2236
	*(*int32)(unsafe.Add(mBase, uint32(v2239))) = int32(393)
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+4))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v2244<<(uint(int32(2))%32))))
	v2249 = F_lappend(m, v2248, v2239)
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L5
	} else {
		goto L367
	}
L367:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+4))
	v2252 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2194+v2251<<(uint(v2252)%32)))) = v2249
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+8))
	switch v2256 - v2252 {
	case 0:
		goto L372
	case 1:
		goto L370
	case 2, 5:
		goto L368
	default:
		goto L371
	}
L368:
	;
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+16))
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+48))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+72))
	v2299 = F_map_variable_attnos(m, v2294, v1159, v2169, v2296, v1148+int32(15))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L5
	} else {
		goto L382
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2239)+8)) = v2291
	goto L368
L370:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+20))
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+40))
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+52))
	v2288 = F_ExecBuildProjectionInfo(m, v2285, v2158, v2286, v76, v2287)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L5
	} else {
		goto L381
	}
L371:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L5
	} else {
		goto L378
	}
L372:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+24))
	if v2169 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v2260 = F_adjust_partition_colnos_using_map(m, v2259, v2169)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L5
	} else {
		goto L376
	}
L374:
	;
	v2263 = v2259
	goto L375
L375:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2236)+20))
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+8))
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2266)+52))
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+40))
	v2270 = F_ExecBuildUpdateProjection(m, v2264, int32(1), v2263, v2267, v2158, v2268, int32(0))
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L5
	} else {
		goto L377
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2236)+24)) = v2260
	v2263 = v2260
	goto L375
L377:
	;
	v2291 = v2270
	goto L369
L378:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_13), int32(0))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L5
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1080), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L5
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	v2291 = v2288
	goto L369
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2236)+16)) = v2299
	v2302 = F_ExecInitQual(m, v2299, v76)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2239)+12)) = v2302
	v2306 = v2212 + int32(1)
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v2161)+4))
	if v2306 < v2307 {
		v2212 = v2306
		goto L363
	} else {
		goto L384
	}
L384:
	;
	goto L364
L385:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_15), int32(0))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L5
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(820), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L5
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2383)))
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+16))
	if v2387 == int32(0) {
		v2397 = v76
		v2398 = v77
		v2399 = v78
		v2400 = v79
		v2401 = v80
		v2402 = v2386
		v2406 = v2385
		v2407 = v86
		v2411 = v90
		v2419 = v98
		v2423 = v102
		v2424 = v103
		v2425 = v104
		v2426 = v105
		goto L173
	} else {
		goto L393
	}
L389:
	;
	v2367 = v1108 << (uint(int32(2)) % 32)
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2367+v103)))
	v2383 = v2367 + v2368
	v2385 = v2371
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(v2372+v1106)))
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v76)+120))
	v2376 = F_ExecInitPartitionDispatchInfo(m, v80, v78, v2374, v85, v1076, v2375)
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L5
	} else {
		goto L392
	}
L392:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+24))
	v2383 = v2378 + v2379<<(uint(int32(2))%32)
	v2385 = v2376
	goto L388
L393:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+20))
	v2391 = F_execute_attr_map_slot(m, v2390, v90, v2387)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L5
	} else {
		goto L394
	}
L394:
	;
	if v98 != 0 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+12))
	m.T0[v2394].(func(*base.Module, int32))(m, v98)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L5
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v2397 = v76
	v2398 = v77
	v2399 = v78
	v2400 = v79
	v2401 = v80
	v2402 = v2386
	v2406 = v2385
	v2407 = v86
	v2411 = v2391
	v2419 = v2387
	v2423 = v102
	v2424 = v103
	v2425 = v104
	v2426 = v105
	goto L173
L398:
	;
	goto L397
L399:
	;
	if v1111 == int32(0) {
		v2447 = v2411
		goto L402
	} else {
		goto L403
	}
L400:
	;
	v2452 = v2411
	goto L401
L401:
	;
	if v2406 != 0 {
		v76 = v2397
		v77 = v2398
		v78 = v2399
		v79 = v2400
		v80 = v2401
		v85 = v2406
		v86 = v2407
		v90 = v2452
		v98 = v2419
		v102 = v2423
		v103 = v2424
		v104 = v2425
		v105 = v2426
		goto L15
	} else {
		goto L408
	}
L402:
	;
	v2449 = F_ExecPartitionCheck(m, v2402, v2447, v2401, int32(1))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L5
	} else {
		goto L407
	}
L403:
	;
	v2438 = F_ExecGetRootToChildMap(m, v2402, v2401)
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L5
	} else {
		goto L404
	}
L404:
	;
	if v2438 == int32(0) {
		v2447 = v2400
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2438)+8))
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2402)+204))
	v2444 = F_execute_attr_map_slot(m, v2442, v2400, v2443)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	v2447 = v2444
	goto L402
L407:
	;
	v2452 = v2447
	goto L401
L408:
	;
	goto L16
L409:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2419)+8))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2455)+12))
	m.T0[v2456].(func(*base.Module, int32))(m, v2419)
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L5
	} else {
		goto L410
	}
L410:
	;
	v2464 = v2402
	v2469 = v2407
	v2485 = v2423
	v2487 = v2425
	v2488 = v2426
	goto L13
}
func F_ExecPartitionCheckEmitError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
		v20 = F_build_attrmap_by_name_if_req(m, v17, v18, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 == int32(0) {
				v33 = l1
				v34 = v13
				v35 = v18
				v36 = v15
				v38 = F_ExecGetInsertedCols(m, v34, l2)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = F_ExecGetUpdatedCols(m, v34, l2)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = F_bms_union(m, v38, v40)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = F_ExecBuildSlotValueDescription(m, v36, v33, v35, v42)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_errcode(m, int32(67391682))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v54 + int32(4)
										F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											if v44 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
												v65 = F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													F_errtable(m, v67)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_errtable(m, v67)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return
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
					}
				}
			} else {
				v26 = F_MakeTupleTableSlot(m, v18, int32(_a_F_ExecPartitionCheckEmitError_4), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = F_execute_attr_map_slot(m, v20, l1, v26)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v33 = v28
						v34 = v13
						v35 = v18
						v36 = v15
						v38 = F_ExecGetInsertedCols(m, v34, l2)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = F_ExecGetUpdatedCols(m, v34, l2)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = F_bms_union(m, v38, v40)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									v44 = F_ExecBuildSlotValueDescription(m, v36, v33, v35, v42)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											F_errcode(m, int32(67391682))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v54 + int32(4)
												F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													if v44 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
														v65 = F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															F_errtable(m, v67)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
																mBase = m.M
																v74 = m.ExcPending
																if v74 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														F_errtable(m, v67)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return
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
							}
						}
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
		v33 = l1
		v34 = l0
		v35 = v31
		v36 = v32
		v38 = F_ExecGetInsertedCols(m, v34, l2)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = F_ExecGetUpdatedCols(m, v34, l2)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				v42 = F_bms_union(m, v38, v40)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = F_ExecBuildSlotValueDescription(m, v36, v33, v35, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errcode(m, int32(67391682))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v54 + int32(4)
								F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									if v44 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
										v65 = F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_errtable(m, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_errtable(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1996), int32(_a_F_ExecPartitionCheckEmitError_3))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
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
			}
		}
	}
}
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F_get_partition_parent_worker(m, v12, l0, v8+int32(31))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				if l1 == int32(0) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v22&int32(1) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							F_errmsg_internal(m, int32(_a_F_get_partition_parent_0), v8+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(69), int32(_a_F_get_partition_parent_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_relation_close(m, v12, int32(1))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return v18
						}
					}
				} else {
					F_relation_close(m, v12, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return v18
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_get_partition_parent_3), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(65), int32(_a_F_get_partition_parent_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v4 {
		v136 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v136
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v136 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = v27
	goto L8
L7:
	;
	v28 = v4
	goto L8
L8:
	;
	if v25 <= int32(0) {
		v136 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = v28
	v35 = v4
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v35<<(uint(int32(1))%32)))))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v136 = int32(0)
	goto L1
L12:
	;
	v129 = v35 + int32(1)
	if v129 != v25 {
		v34 = v125
		v35 = v129
		goto L10
	} else {
		goto L42
	}
L13:
	;
	v48 = F_bms_is_member(m, v45+int32(7), l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	F_pull_varattnos(m, v57, int32(1), v13+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L19
	}
L16:
	;
	if v48 == int32(0) {
		v125 = v34
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(1)
	if l2 == int32(0) {
		v136 = v52
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	v136 = v52
	goto L1
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v68 = int32(0)
	if base.B2i32(l1 == v68)|base.B2i32(v67 == v68) != 0 {
		v113 = v68
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v113 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	goto L20
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v78 < v79 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = v78
	goto L25
L24:
	;
	v81 = v79
	goto L25
L25:
	;
	if v81 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = int32(1)
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v85 = int32(8)
	v90 = int32(0)
	goto L29
L29:
	;
	v97 = v90 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67+v85+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85+v97)))
	v102 = v99 & v101
	v104 = base.B2i32(v102 != int32(0))
	if v102 != 0 {
		v113 = v104
		goto L21
	} else {
		goto L31
	}
L30:
	;
	v113 = v104
	goto L21
L31:
	;
	v106 = v90 + int32(1)
	if v106 != v84 {
		v90 = v106
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v118 = v34 + int32(4)
	if base.Ui32(v118) < base.Ui32(v66+v65<<(uint(int32(2))%32)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v114)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v136 = int32(1)
	goto L1
L39:
	;
	v124 = v118
	goto L41
L40:
	;
	v124 = int32(0)
	goto L41
L41:
	;
	v125 = v124
	goto L12
L42:
	;
	goto L11
}
func F_partition_list_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v13 = v11 - int32(1)
	if v13 < int32(0) {
		v57 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v57
L2:
	;
	v21 = v10
	v22 = v13
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v30 = int32(2)
	v31 = base.I32_div_s(v21+v22+int32(1), v30)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(v30)%32))))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v37 = F_FunctionCall2Coll(m, l0, v25, v36, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v57 = v49
	goto L1
L5:
	;
	if v49 < v50 {
		v21 = v49
		v22 = v50
		goto L3
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	v41 = base.I32_wrap_i64(v37)
	if v41 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(base.B2i32(v41 == int32(0)))
	if v41 != 0 {
		v49 = v31
		v50 = v22
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v49 = v21
	v50 = v31 - int32(1)
	goto L5
L11:
	;
	v57 = v31
	goto L1
L12:
	;
	goto L4
}
