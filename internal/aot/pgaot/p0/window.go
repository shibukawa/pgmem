package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecWindowAgg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v129 int64
	_ = v129
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
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int64
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
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
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v631 int32
	_ = v631
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v666 int64
	_ = v666
	var v669 int32
	_ = v669
	var v677 int64
	_ = v677
	var v679 int64
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int64
	_ = v712
	var v715 int64
	_ = v715
	var v716 int64
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int64
	_ = v724
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v762 int64
	_ = v762
	var v763 int64
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v967 int32
	_ = v967
	var v970 int64
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int64
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1112 int32
	_ = v1112
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1168 int64
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1259 int64
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1291 int32
	_ = v1291
	var v1296 int64
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int64
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1321 int32
	_ = v1321
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int64
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1508 int32
	_ = v1508
	var v1509 int64
	_ = v1509
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int64
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1647 int32
	_ = v1647
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int64
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1803 int32
	_ = v1803
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1996 int32
	_ = v1996
	var v2003 int32
	_ = v2003
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2335 int32
	_ = v2335
	var v2344 int32
	_ = v2344
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2375 int32
	_ = v2375
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2399 int32
	_ = v2399
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2411 int32
	_ = v2411
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2456 int32
	_ = v2456
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2469 int32
	_ = v2469
	var v2494 int32
	_ = v2494
	var v2503 int32
	_ = v2503
	var v2522 int32
	_ = v2522
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2568 float64
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	v24 = m.G0
	v26 = v24 - int32(832)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[0]))
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v34 == int32(0) {
		v2576 = int32(0)
		v2577 = v26
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
	m.G0 = v2577 + int32(832)
	return v2576
L7:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+376)))
	if v37 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = m.G0
	v42 = v40 - int32(16)
	m.G0 = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v45&int32(_a_F_ExecWindowAgg_0) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L10
L10:
	;
	v211 = l0
	v215 = v26
	v227 = v26 + int32(32)
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
	if v45&int32(_a_F_ExecWindowAgg_1) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v50 = int32(_a_F_ExecWindowAgg_2)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v59 = m.T0[v58].(func(*base.Module, int32, int32, int32) int32)(m, v52, v44, v42+int32(15))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v51
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+15)))
	if v63 == int32(1) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v68 = F_exprType(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_get_typlenbyval(m, v68, v42+int32(12), v42+int32(11))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+11)))
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+12)))
	v78 = F_datumCopy(m, v59, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v78
	if v45&int32(12) == int32(0) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	if v85 < int64(0) {
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+376)) = uint8(v135)
	m.G0 = v42 + int32(16)
	goto L11
L26:
	;
	v94 = int32(_a_F_ExecWindowAgg_2)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v98
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v103 = m.T0[v102].(func(*base.Module, int32, int32, int32) int32)(m, v96, v44, v42+int32(15))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v95
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+15)))
	if v107 == int32(1) {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+24))
	v112 = F_exprType(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_get_typlenbyval(m, v112, v42+int32(12), v42+int32(11))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+11)))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+12)))
	v122 = F_datumCopy(m, v103, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v122
	if v45&int32(12) == int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	if v129 < int64(0) {
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
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2106), int32(_a_F_ExecWindowAgg_5))
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
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2120), int32(_a_F_ExecWindowAgg_5))
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
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2133), int32(_a_F_ExecWindowAgg_5))
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
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2147), int32(_a_F_ExecWindowAgg_5))
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
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+378)))
	if v234 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v2576 = v2277
	v2577 = v215
	goto L6
L52:
	;
	F_spool_tuples(m, v211, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L57
	}
L53:
	;
	F_begin_partition(m, v211)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v240 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v211)+380)) = uint16(v240)
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	v244 = v242 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v211)+176)) = v244
	v246 = v244
	goto L52
L56:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	v246 = v239
	goto L52
L57:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+377)))
	if v249 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v2572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+311)))
	if v2572 != 0 {
		goto L50
	} else {
		goto L450
	}
L59:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	if v2550 == int32(0) {
		v2576 = v2277
		v2577 = v215
		goto L6
	} else {
		goto L446
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+224)) = int32(2)
	goto L59
L61:
	;
	v2522 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+224)) = v2522
	v2576 = v2522
	v2577 = v2503
	goto L6
L62:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v211)+64))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	F_MemoryContextReset(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L68
	}
L63:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v211)+168))
	if v252 < v253 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_release_partition(m, v211)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+379)))
	if v257 != int32(1) {
		v2503 = v215
		goto L61
	} else {
		goto L66
	}
L66:
	;
	F_begin_partition(m, v211)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+224)) = int32(1)
	goto L62
L68:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v211)+144))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v211)+148))
	F_tuplestore_select_read_pointer(m, v268, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v211)+228))
	if v272&int32(_a_F_ExecWindowAgg_9) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v211)+152))
	if int32(0) <= v2255 {
		goto L413
	} else {
		goto L414
	}
L71:
	;
	v2128 = int32(0)
	if v656 != int32(1) {
		goto L406
	} else {
		goto L407
	}
L72:
	;
	v1803 = int32(0)
	goto L352
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L4
	} else {
		goto L349
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L4
	} else {
		goto L346
	}
L75:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v211)+224))
	if v351 != int32(1) {
		goto L70
	} else {
		goto L94
	}
L76:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v211)+144))
	v341 = int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v211)+112))
	v344 = F_tuplestore_gettupleslot(m, v340, v341, v341, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L92
	}
L77:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	if v277 <= int64(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v211)+404))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v211)+112))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+32))
	m.T0[v283].(func(*base.Module, int32, int32))(m, v280, v281)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v211)+144))
	v287 = int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v211)+112))
	v290 = F_tuplestore_gettupleslot(m, v286, v287, v287, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	if v290 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L81:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+96))
	if v295 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v211)+404))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	m.T0[v337].(func(*base.Module, int32))(m, v335)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L91
	}
L83:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v211)+404))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v211)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v299)+8)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v299)+12)) = v298
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v211)+140))
	if v303 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	F_MemoryContextReset(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v309 = int32(_a_F_ExecWindowAgg_2)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v312
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	v317 = m.T0[v316].(func(*base.Module, int32, int32, int32) int32)(m, v303, v299, v215+int32(12))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L82
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v310
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	F_MemoryContextReset(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v317 != 0 {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+382)) = uint8(v324)
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+344)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v211)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+320)) = v328 + int64(1)
	goto L82
L91:
	;
	goto L75
L92:
	;
	if v344 == int32(0) {
		goto L74
	} else {
		goto L93
	}
L93:
	;
	goto L75
L94:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v211)+120))
	if int32(0) < v354 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v365 = int32(0)
	goto L98
L96:
	;
	goto L97
L97:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v211)+124))
	if v656 <= int32(0) {
		goto L70
	} else {
		goto L123
	}
L98:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v211)+128))
	v384 = v381 + v365*int32(56)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+47)))
	if v385 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L97
L100:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v392 = int32(_a_F_ExecWindowAgg_2)
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v211)+64))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v396
	v399 = v384 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v399
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v384)+52))
	v402 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v401
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v384)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)) = uint8(v402)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v405
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+30)) = uint16(v409)
	if v402 < v409 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v631 = v365 + int32(1)
	if v631 != v354 {
		v365 = v631
		goto L98
	} else {
		goto L122
	}
L103:
	;
	v414 = v409 & int32(7)
	v415 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v409) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v560 = v399
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = int32(0)
	v583 = v388 + v391<<(uint(int32(2))%32)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v587 = m.T0[v586].(func(*base.Module, int32) int32)(m, v215+int32(12))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L117
	}
L106:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v560 = v554
	goto L105
L107:
	;
	v425 = v415
	v432 = int32(0)
	goto L110
L108:
	;
	v478 = v415
	goto L109
L109:
	;
	v501 = v478
	v503 = v415
	goto L114
L110:
	;
	v449 = v215 + int32(12) + v425<<(uint(int32(3))%32)
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+80)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+72)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449-int32(-64)))) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+56)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+48)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+40)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+32)) = uint8(v450)
	*(*uint8)(unsafe.Add(mBase, uint32(v449)+24)) = uint8(v450)
	v468 = int32(8)
	v469 = v425 + v468
	v471 = v432 + v468
	if v471 != v409&int32(2147483640) {
		v425 = v469
		v432 = v471
		goto L110
	} else {
		goto L112
	}
L111:
	;
	if v414 == int32(0) {
		goto L106
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v478 = v469
	goto L109
L114:
	;
	v524 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v501<<(uint(int32(3))%32)+v215)+36)) = uint8(v524)
	v529 = v503 + v524
	if v529 != v414 {
		v501 = v501 + v524
		v503 = v529
		goto L114
	} else {
		goto L116
	}
L115:
	;
	goto L106
L116:
	;
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v587
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v391+v389))) = uint8(v590)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+46)))
	if (v590|v592)&int32(1) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v393
	goto L102
L119:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v211)+120))
	if v596 < int32(2) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	v601 = int32(*(*int16)(unsafe.Add(mBase, uint32(v384)+44)))
	v602 = F_datumCopy(m, v599, int32(0), v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v602
	goto L118
L122:
	;
	goto L99
L123:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v211)+400))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v211)+396))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v211)+200))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v211)+64))
	F_update_frameheadpos(m, v211)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v211)+184))
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v211)+208))
	if v666 <= v665 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if v665 != v666 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L4
	} else {
		goto L343
	}
L128:
	;
	v683 = int32(0)
	v686 = v683
	v691 = v683
	goto L133
L129:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v211)+228))
	if base.B2i32(v669&int32(1280) == int32(0))|v669&int32(_a_F_ExecWindowAgg_10) != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	if v677 < v665 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	if v677 < v679 {
		goto L71
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v711 = v708 + v686*int32(160)
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v211)+176))
	if v712 == int64(0) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	if v656 <= v733 {
		v1151 = v733
		goto L148
	} else {
		goto L149
	}
L135:
	;
	v736 = v686 + int32(1)
	if v736 != v656 {
		v686 = v736
		v691 = v733
		goto L133
	} else {
		goto L145
	}
L136:
	;
	v731 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v711)+152)) = uint8(v731)
	v733 = v691
	goto L135
L137:
	;
	v727 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v711)+152)) = uint8(v727)
	v733 = v691 + v727
	goto L135
L138:
	;
	v715 = *(*int64)(unsafe.Add(mBase, uint32(v211)+184))
	v716 = *(*int64)(unsafe.Add(mBase, uint32(v211)+208))
	if v715 != v716 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	if v718 == int32(0) {
		goto L137
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+229)))
	if v721&int32(896) != 0 {
		goto L137
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v724 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	if v715 < v724 {
		goto L136
	} else {
		goto L144
	}
L144:
	;
	goto L137
L145:
	;
	goto L134
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L4
	} else {
		goto L340
	}
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L4
	} else {
		goto L337
	}
L148:
	;
	v1168 = *(*int64)(unsafe.Add(mBase, uint32(v211)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+208)) = v1168
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v661)+16))
	if int32(0) <= v1170 {
		goto L219
	} else {
		goto L220
	}
L149:
	;
	v745 = v733
	goto L150
L150:
	;
	v762 = *(*int64)(unsafe.Add(mBase, uint32(v211)+208))
	v763 = *(*int64)(unsafe.Add(mBase, uint32(v211)+184))
	if v763 <= v762 {
		v1151 = v745
		goto L148
	} else {
		goto L152
	}
L151:
	;
	v1151 = v1112
	goto L148
L152:
	;
	v765 = F_window_gettupleslot(m, v661, v762, v659)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	if v765 == int32(0) {
		goto L146
	} else {
		goto L154
	}
L154:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v769)+12)) = v659
	v778 = v745
	v784 = int32(0)
	goto L155
L155:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v798 = v795 + v784*int32(160)
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+152)))
	if v799 != 0 {
		v1112 = v778
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+20))
	F_MemoryContextReset(m, v1133)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L4
	} else {
		goto L216
	}
L157:
	;
	v1130 = v784 + int32(1)
	if v1130 != v656 {
		v778 = v1112
		v784 = v1130
		goto L155
	} else {
		goto L215
	}
L158:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v211)+128))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798)+124))
	v804 = v800 + v801*int32(56)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v804)+8))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v804)))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v806)+12))
	v808 = int32(_a_F_ExecWindowAgg_2)
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v812
	if v807 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v806)+8))
	if v829 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v807)+20))
	v819 = m.T0[v818].(func(*base.Module, int32, int32, int32) int32)(m, v807, v811, v215+int32(11))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+11)))
	if v819 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v823 = v821
	goto L164
L163:
	;
	v823 = int32(1)
	goto L164
L164:
	;
	if v823 == int32(0) {
		goto L159
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v809
	v1112 = v778
	goto L157
L166:
	;
	v903 = int32(1)
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+50)))
	if base.B2i32(v904 != v903)|base.B2i32(v805 <= int32(0)) != 0 {
		goto L173
	} else {
		goto L174
	}
L167:
	;
	v833 = int32(0)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	if v834 <= v833 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v838 = int32(1)
	v840 = v833
	goto L169
L169:
	;
	v862 = v227 + v838<<(uint(int32(3))%32)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v829)+12))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v863+v840<<(uint(int32(2))%32))))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v867)+20))
	v871 = m.T0[v870].(func(*base.Module, int32, int32, int32) int32)(m, v867, v811, v862+int32(4))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L171
	}
L170:
	;
	goto L166
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v871
	v874 = int32(1)
	v877 = v840 + v874
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	if v877 < v878 {
		v838 = v838 + v874
		v840 = v877
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+136)))
	if v967 == int32(1) {
		goto L147
	} else {
		goto L181
	}
L174:
	;
	v911 = v903
	goto L175
L175:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911<<(uint(int32(3))%32)+v215)+36)))
	if v936 != int32(1) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v809
	v1112 = v778
	goto L157
L177:
	;
	v940 = v911 + int32(1)
	if v940 <= v805 {
		v911 = v940
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	goto L176
L180:
	;
	goto L173
L181:
	;
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v798)+144))
	if v970 == int64(1) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v809
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v798)+128))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v211)+364))
	if v975 != v976 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L184
L184:
	;
	v1009 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v798 + int32(40)
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v804)+40))
	v1017 = v805 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+30)) = uint16(v1017)
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)) = uint8(v1009)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v1015
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v798)+132))
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+36)) = uint8(v1009)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v1022
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v798)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = v1026
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v798)+40))
	v1031 = m.T0[v1030].(func(*base.Module, int32) int32)(m, v215+int32(12))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L4
	} else {
		goto L194
	}
L185:
	;
	F_MemoryContextReset(m, v975)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L4
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+104)))
	if v980 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L187
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v798)+144)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+136)) = uint8(v1000)
	*(*int32)(unsafe.Add(mBase, uint32(v798)+132)) = v998
	v1005 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+112)) = uint8(v1005)
	*(*int32)(unsafe.Add(mBase, uint32(v798)+108)) = int32(0)
	v1112 = v778
	goto L157
L190:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v798)+100))
	v998 = v983
	v1000 = int32(1)
	goto L189
L191:
	;
	goto L192
L192:
	;
	v985 = int32(_a_F_ExecWindowAgg_2)
	v986 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v798)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v988
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v798)+100))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+122)))
	v992 = int32(*(*int16)(unsafe.Add(mBase, uint32(v798)+118)))
	v993 = F_datumCopy(m, v990, v991, v992)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v986
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+104)))
	v998 = v993
	v1000 = v997
	goto L189
L194:
	;
	v1033 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = v1033
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	if v1035 == v1033 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1038 = *(*int64)(unsafe.Add(mBase, uint32(v798)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v798)+144)) = v1038 - int64(1)
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+122)))
	if v1042 != 0 {
		v1092 = v1031
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v809
	v1102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+152)) = uint8(v1102)
	v1112 = v778 + v1102
	goto L157
L198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v798)+132)) = v1092
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+136)) = uint8(v1098)
	v1112 = v778
	goto L157
L199:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v798)+132))
	if v1031 == v1043 {
		v1092 = v1031
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1045 = int32(0)
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v798)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1047
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+118)))
	if v1049 != int32(_a_F_ExecWindowAgg_11) {
		v1067 = v1049
		v1068 = v1045
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+136)))
	if v1077 != 0 {
		v1092 = v1074
		goto L198
	} else {
		goto L208
	}
L202:
	;
	v1072 = F_datumCopy(m, v1031, v1068&int32(1), base.I32_extend16_s(v1067))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L207
	}
L203:
	;
	v1052 = int32(_a_F_ExecWindowAgg_11)
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	if v1053 != int32(1) {
		v1067 = v1052
		v1068 = v1045
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031)+1)))
	if v1056 != int32(3) {
		v1067 = v1052
		v1068 = v1045
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+2))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+8))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+16))
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	if v1061 == v1063 {
		v1074 = v1031
		goto L201
	} else {
		goto L206
	}
L206:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+122)))
	v1066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+118)))
	v1067 = v1066
	v1068 = v1065
	goto L202
L207:
	;
	v1074 = v1072
	goto L201
L208:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v798)+132))
	v1079 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+118)))
	if v1079 != int32(_a_F_ExecWindowAgg_11) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_pfree(m, v1078)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L214
	}
L210:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078))))
	if v1082 != int32(1) {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078)+1)))
	if v1085 != int32(3) {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	F_DeleteExpandedObject(m, v1078)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v1092 = v1074
	goto L198
L214:
	;
	v1092 = v1074
	goto L198
L215:
	;
	goto L156
L216:
	;
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v211)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+208)) = v1136 + int64(1)
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v659)+8))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+12))
	m.T0[v1141].(func(*base.Module, int32))(m, v659)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	if v1112 < v656 {
		v745 = v1112
		goto L150
	} else {
		goto L218
	}
L218:
	;
	goto L151
L219:
	;
	F_WinSetMarkPosition(m, v661, v1168)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v1175 = int32(0)
	v1176 = base.B2i32(v1151 <= v1175)
	if v1176 == v1175 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L221
L223:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v211)+364))
	F_MemoryContextReset(m, v1179)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1186 = int32(0)
	goto L227
L226:
	;
	goto L225
L227:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v1209 = v1206 + v1186*int32(160)
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+152)))
	if v1210 == int32(1) {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v1259 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	if v1151 <= v1175 {
		goto L247
	} else {
		goto L248
	}
L229:
	;
	v1257 = v1186 + int32(1)
	if v1257 != v656 {
		v1186 = v1257
		goto L227
	} else {
		goto L246
	}
L230:
	;
	v1250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1209)+112)) = uint8(v1250)
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+108)) = int32(0)
	goto L229
L231:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+128))
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v211)+364))
	if v1213 != v1214 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+112)))
	if v1243 != 0 {
		goto L229
	} else {
		goto L243
	}
L234:
	;
	F_MemoryContextReset(m, v1213)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L4
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+104)))
	if v1218 == int32(1) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L236
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1209)+144)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1209)+136)) = uint8(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+132)) = v1236
	goto L230
L239:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+100))
	v1236 = v1221
	v1238 = int32(1)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1223 = int32(_a_F_ExecWindowAgg_2)
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1226
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+100))
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+122)))
	v1230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1209)+118)))
	v1231 = F_datumCopy(m, v1228, v1229, v1230)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1224
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+104)))
	v1236 = v1231
	v1238 = v1235
	goto L238
L243:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+121)))
	if v1244 != 0 {
		goto L230
	} else {
		goto L244
	}
L244:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+108))
	F_pfree(m, v1245)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
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
	v1260 = *(*int64)(unsafe.Add(mBase, uint32(v211)+184))
	if v1259 == v1260 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v211)+216)) = v1260
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+12))
	m.T0[v1264].(func(*base.Module, int32))(m, v660)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	if v660 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1301 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	v1302 = F_row_is_in_frame(m, v211, v1301, v660)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L4
	} else {
		goto L260
	}
L254:
	;
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+4)))
	if v1291&int32(2) == int32(0) {
		goto L253
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1296 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	v1297 = F_window_gettupleslot(m, v661, v1296, v660)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L4
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	if v1297 == int32(0) {
		goto L72
	} else {
		goto L259
	}
L259:
	;
	goto L253
L260:
	;
	if v1302 < int32(0) {
		goto L72
	} else {
		goto L261
	}
L261:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	if v1302 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1306)+12)) = v660
	v1321 = int32(0)
	goto L265
L263:
	;
	v1717 = v1306
	goto L264
L264:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+20))
	F_MemoryContextReset(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L4
	} else {
		goto L335
	}
L265:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v1335 = v1332 + v1321*int32(160)
	v1336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+152)))
	if v1336 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	v1717 = v1693
	goto L264
L267:
	;
	v1691 = v1321 + int32(1)
	if v1691 != v656 {
		v1321 = v1691
		goto L265
	} else {
		goto L334
	}
L268:
	;
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	if v1339 < v1259 {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v211)+128))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+124))
	v1345 = v1341 + v1342*int32(56)
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+8))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1345)))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+12))
	v1349 = int32(_a_F_ExecWindowAgg_2)
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v211)+372))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1353
	if v1348 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L270
L272:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+8))
	if v1370 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L273:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+20))
	v1360 = m.T0[v1359].(func(*base.Module, int32, int32, int32) int32)(m, v1348, v1352, v215+int32(11))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+11)))
	if v1360 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1364 = v1362
	goto L277
L276:
	;
	v1364 = int32(1)
	goto L277
L277:
	;
	if v1364 == int32(0) {
		goto L272
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1350
	goto L267
L279:
	;
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+22)))
	if v1444 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L280:
	;
	v1374 = int32(0)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+4))
	if v1375 <= v1374 {
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v1379 = int32(1)
	v1381 = v1374
	goto L282
L282:
	;
	v1403 = v227 + v1379<<(uint(int32(3))%32)
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+12))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1404+v1381<<(uint(int32(2))%32))))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+20))
	v1412 = m.T0[v1411].(func(*base.Module, int32, int32, int32) int32)(m, v1408, v1352, v1403+int32(4))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L4
	} else {
		goto L284
	}
L283:
	;
	goto L279
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403))) = v1412
	v1415 = int32(1)
	v1418 = v1381 + v1415
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+4))
	if v1418 < v1419 {
		v1379 = v1379 + v1415
		v1381 = v1418
		goto L282
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L4
	} else {
		goto L330
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1350
	goto L267
L288:
	;
	v1558 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v1558
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v211
	v1561 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v1335 + v1561
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+40))
	v1566 = v1346 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+30)) = uint16(v1566)
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)) = uint8(v1558)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v1564
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+132))
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+36)) = uint8(v1557)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v1571
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = v1574
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+12))
	v1579 = m.T0[v1578].(func(*base.Module, int32) int32)(m, v215+v1561)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L4
	} else {
		goto L307
	}
L289:
	;
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+136)))
	v1557 = v1447
	goto L288
L290:
	;
	goto L291
L291:
	;
	if v1346 <= int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+136)))
	v1509 = *(*int64)(unsafe.Add(mBase, uint32(v1335)+144))
	if v1509 == int64(0) {
		goto L301
	} else {
		goto L302
	}
L293:
	;
	v1452 = int32(1)
	goto L294
L294:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1452<<(uint(int32(3))%32)+v215)+36)))
	if v1477 != int32(1) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1350
	goto L267
L296:
	;
	v1481 = v1452 + int32(1)
	if v1481 <= v1346 {
		v1452 = v1481
		goto L294
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	goto L295
L299:
	;
	goto L292
L300:
	;
	v1557 = int32(0)
	goto L288
L301:
	;
	if v1508&int32(1) == int32(0) {
		goto L300
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	if v1508&int32(1) != 0 {
		goto L287
	} else {
		goto L306
	}
L304:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1517
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v215)+40))
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+122)))
	v1521 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1335)+118)))
	v1522 = F_datumCopy(m, v1519, v1520, v1521)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L4
	} else {
		goto L305
	}
L305:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1335)+144)) = int64(1)
	v1526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+136)) = uint8(v1526)
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+132)) = v1522
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1350
	goto L267
L306:
	;
	goto L300
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = int32(0)
	v1583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	if v1583 == int32(1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+4))
	if v1586 != 0 {
		goto L286
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v1587 = *(*int64)(unsafe.Add(mBase, uint32(v1335)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v1335)+144)) = v1587 + int64(1)
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+122)))
	if v1591 != 0 {
		v1641 = v1579
		goto L312
	} else {
		goto L313
	}
L311:
	;
	goto L310
L312:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+132)) = v1641
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+136)) = uint8(v1647)
	goto L267
L313:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+132))
	if v1579 == v1592 {
		v1641 = v1579
		goto L312
	} else {
		goto L314
	}
L314:
	;
	if v1583 != 0 {
		v1623 = v1579
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+136)))
	if v1626 != 0 {
		v1641 = v1623
		goto L312
	} else {
		goto L323
	}
L316:
	;
	v1594 = int32(0)
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1596
	v1598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1335)+118)))
	if v1598 != int32(_a_F_ExecWindowAgg_11) {
		v1616 = v1598
		v1617 = v1594
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1621 = F_datumCopy(m, v1579, v1617&int32(1), base.I32_extend16_s(v1616))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L4
	} else {
		goto L322
	}
L318:
	;
	v1601 = int32(_a_F_ExecWindowAgg_11)
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579))))
	if v1602 != int32(1) {
		v1616 = v1601
		v1617 = v1594
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579)+1)))
	if v1605 != int32(3) {
		v1616 = v1601
		v1617 = v1594
		goto L317
	} else {
		goto L320
	}
L320:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+2))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+8))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1609)+16))
	v1612 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	if v1610 == v1612 {
		v1623 = v1579
		goto L315
	} else {
		goto L321
	}
L321:
	;
	v1614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335)+122)))
	v1615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1335)+118)))
	v1616 = v1615
	v1617 = v1614
	goto L317
L322:
	;
	v1623 = v1621
	goto L315
L323:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+132))
	v1628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1335)+118)))
	if v1628 != int32(_a_F_ExecWindowAgg_11) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	F_pfree(m, v1627)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L4
	} else {
		goto L329
	}
L325:
	;
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627))))
	if v1631 != int32(1) {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627)+1)))
	if v1634 != int32(3) {
		goto L324
	} else {
		goto L327
	}
L327:
	;
	F_DeleteExpandedObject(m, v1627)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L4
	} else {
		goto L328
	}
L328:
	;
	v1641 = v1623
	goto L312
L329:
	;
	v1641 = v1623
	goto L312
L330:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L4
	} else {
		goto L331
	}
L331:
	;
	F_errmsg(m, int32(_a_F_ExecWindowAgg_12), int32(0))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L4
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(356), int32(_a_F_ExecWindowAgg_13))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	goto L266
L335:
	;
	v1721 = *(*int64)(unsafe.Add(mBase, uint32(v211)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+216)) = v1721 + int64(1)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1725)+12))
	m.T0[v1726].(func(*base.Module, int32))(m, v660)
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	goto L251
L337:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_14), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L4
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(488), int32(_a_F_ExecWindowAgg_15))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L4
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_16), int32(0))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(816), int32(_a_F_ExecWindowAgg_17))
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_18), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L4
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(738), int32(_a_F_ExecWindowAgg_17))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_19), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2275), int32(_a_F_ExecWindowAgg_20))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecWindowAgg_19), int32(0))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_ExecWindowAgg_4), int32(2261), int32(_a_F_ExecWindowAgg_20))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v1822 = v1819 + v1803*int32(160)
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+124))
	v1824 = int32(_a_F_ExecWindowAgg_2)
	v1825 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v211)+128))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v662)+32))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v662)+36))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v211)+64))
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1830)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1831
	v1833 = v1823 + v1828
	v1836 = v1827 + v1823<<(uint(int32(2))%32)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+8))
	if v1837 != 0 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	goto L70
L354:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1825
	v2108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+121)))
	if v2108 != 0 {
		goto L401
	} else {
		goto L402
	}
L355:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+96))
	v1839 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v1839
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v1822 + int32(68)
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1826+v1823*int32(56))+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+30)) = uint16(v1838)
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)) = uint8(v1839)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v1848
	v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+136)))
	if v1853 == v1839 {
		goto L360
	} else {
		goto L361
	}
L356:
	;
	goto L357
L357:
	;
	v2061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+136)))
	if v2061 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L358:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+36)) = uint8(v1873)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v1872
	if v1838 < int32(2) {
		v2003 = v1873
		goto L368
	} else {
		goto L369
	}
L359:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+132))
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860))))
	if v1861 != int32(1) {
		v1870 = v1860
		goto L365
	} else {
		goto L366
	}
L360:
	;
	v1856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1822)+118)))
	if v1856 == int32(_a_F_ExecWindowAgg_11) {
		goto L359
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+132))
	v1872 = v1859
	v1873 = v1853
	goto L358
L363:
	;
	goto L362
L364:
	;
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+136)))
	v1872 = v1870
	v1873 = v1871
	goto L358
L365:
	;
	goto L364
L366:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860)+1)))
	if v1864 != int32(3) {
		v1870 = v1860
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1860)+2))
	v1870 = v1867 + int32(18)
	goto L365
L368:
	;
	v2021 = int32(1)
	v2023 = int32(0)
	v2025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+78)))
	if base.B2i32(v2003&v2021 == v2023)|base.B2i32(v2025 != v2021) == v2023 {
		goto L379
	} else {
		goto L380
	}
L369:
	;
	v1878 = int32(1)
	v1879 = v1838 - v1878
	v1880 = int32(3)
	v1881 = v1879 & v1880
	if base.Ui32(v1838-int32(2)) < base.Ui32(v1880) {
		v1941 = v1878
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1963 = int32(0)
	v1965 = v1941
	goto L376
L371:
	;
	v1893 = v1878
	v1895 = int32(0)
	goto L372
L372:
	;
	v1915 = v227 + v1893<<(uint(int32(3))%32)
	v1916 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1915)+28)) = uint8(v1916)
	v1918 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1915)+24)) = v1918
	*(*uint8)(unsafe.Add(mBase, uint32(v1915)+20)) = uint8(v1916)
	*(*int32)(unsafe.Add(mBase, uint32(v1915)+16)) = v1918
	*(*uint8)(unsafe.Add(mBase, uint32(v1915)+12)) = uint8(v1916)
	*(*int32)(unsafe.Add(mBase, uint32(v1915)+8)) = v1918
	*(*uint8)(unsafe.Add(mBase, uint32(v1915)+4)) = uint8(v1916)
	*(*int32)(unsafe.Add(mBase, uint32(v1915))) = v1918
	v1932 = int32(4)
	v1933 = v1893 + v1932
	v1935 = v1895 + v1932
	if v1935 != v1879&int32(-4) {
		v1893 = v1933
		v1895 = v1935
		goto L372
	} else {
		goto L374
	}
L373:
	;
	if v1881 != 0 {
		v1941 = v1933
		goto L370
	} else {
		goto L375
	}
L374:
	;
	goto L373
L375:
	;
	v2003 = int32(1)
	goto L368
L376:
	;
	v1985 = int32(1)
	v1988 = v227 + v1965<<(uint(int32(3))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1988)+4)) = uint8(v1985)
	*(*int32)(unsafe.Add(mBase, uint32(v1988))) = int32(0)
	v1996 = v1963 + v1985
	if v1996 != v1881 {
		v1963 = v1996
		v1965 = v1965 + v1985
		goto L376
	} else {
		goto L378
	}
L377:
	;
	v2003 = v1985
	goto L368
L378:
	;
	goto L377
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1836))) = int32(0)
	v2033 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1833))) = uint8(v2033)
	goto L354
L380:
	;
	goto L381
L381:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = v2035
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+68))
	v2040 = m.T0[v2039].(func(*base.Module, int32) int32)(m, v215+int32(12))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L4
	} else {
		goto L382
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+368)) = int32(0)
	v2044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1833))) = uint8(v2044)
	if v2044 != 0 {
		v2059 = v2040
		goto L383
	} else {
		goto L384
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1836))) = v2059
	goto L354
L384:
	;
	v2046 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1822)+116)))
	if v2046 != int32(_a_F_ExecWindowAgg_11) {
		v2059 = v2040
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v2049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2040))))
	if v2049 != int32(1) {
		v2058 = v2040
		goto L387
	} else {
		goto L388
	}
L386:
	;
	v2059 = v2058
	goto L383
L387:
	;
	goto L386
L388:
	;
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2040)+1)))
	if v2052 != int32(3) {
		v2058 = v2040
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2040)+2))
	v2058 = v2055 + int32(18)
	goto L387
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1836))) = v2079
	v2081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+136)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1833))) = uint8(v2081)
	goto L354
L391:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+132))
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068))))
	if v2069 != int32(1) {
		v2078 = v2068
		goto L397
	} else {
		goto L398
	}
L392:
	;
	v2064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1822)+118)))
	if v2064 == int32(_a_F_ExecWindowAgg_11) {
		goto L391
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+132))
	v2079 = v2067
	goto L390
L395:
	;
	goto L394
L396:
	;
	v2079 = v2078
	goto L390
L397:
	;
	goto L396
L398:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068)+1)))
	if v2072 != int32(3) {
		v2078 = v2068
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+2))
	v2078 = v2075 + int32(18)
	goto L397
L400:
	;
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1822)+112)) = uint8(v2123)
	v2126 = v1803 + int32(1)
	if v2126 != v656 {
		v1803 = v2126
		goto L352
	} else {
		goto L405
	}
L401:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v1836)))
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+108)) = v2121
	goto L400
L402:
	;
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833))))
	if v2109 != 0 {
		goto L401
	} else {
		goto L403
	}
L403:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2111
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v1836)))
	v2114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822)+121)))
	v2115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1822)+116)))
	v2116 = F_datumCopy(m, v2113, v2114, v2115)
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L4
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1822)+108)) = v2116
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v1825
	goto L400
L405:
	;
	goto L353
L406:
	;
	v2137 = v2128
	v2141 = int32(0)
	goto L409
L407:
	;
	v2195 = v2128
	goto L408
L408:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v662)+32))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v2221 = v2218 + v2195*int32(160)
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2221)+124))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2221)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v2217+v2222<<(uint(int32(2))%32)))) = v2226
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v662)+36))
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2221)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2222+v2228))) = uint8(v2230)
	goto L70
L409:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v662)+32))
	v2161 = v2137 * int32(160)
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v2163 = v2161 + v2162
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v2163)+124))
	v2165 = int32(2)
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2163)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v2159+v2164<<(uint(v2165)%32)))) = v2168
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v662)+36))
	v2172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2163)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2164+v2170))) = uint8(v2172)
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v662)+32))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v211)+132))
	v2176 = v2175 + v2161
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+284))
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v2174+v2177<<(uint(v2165)%32)))) = v2181
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v662)+36))
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2176)+272)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2177+v2183))) = uint8(v2185)
	v2188 = v2137 + v2165
	v2190 = v2141 + v2165
	if v2190 != v656&int32(2147483646) {
		v2137 = v2188
		v2141 = v2190
		goto L409
	} else {
		goto L411
	}
L410:
	;
	if v656&int32(1) == int32(0) {
		goto L70
	} else {
		goto L412
	}
L411:
	;
	goto L410
L412:
	;
	v2195 = v2188
	goto L408
L413:
	;
	F_update_frameheadpos(m, v211)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L4
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v211)+156))
	if int32(0) <= v2260 {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	goto L415
L417:
	;
	F_update_frametailpos(m, v211)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L4
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v211)+160))
	if int32(0) <= v2265 {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	goto L419
L421:
	;
	F_update_grouptailpos(m, v211)
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L4
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v211)+144))
	F_tuplestore_trim(m, v2270)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L4
	} else {
		goto L425
	}
L424:
	;
	goto L423
L425:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v211)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = v2273
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v211)+68))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+72))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+16))
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2277)+8))
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2278)+12))
	m.T0[v2279].(func(*base.Module, int32))(m, v2277)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L4
	} else {
		goto L426
	}
L426:
	;
	v2282 = int32(_a_F_ExecWindowAgg_2)
	v2283 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2285
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+24))
	v2291 = m.T0[v2290].(func(*base.Module, int32, int32, int32) int32)(m, v2275+int32(4), v2276, int32(0))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L4
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2283
	v2295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2277)+4)))
	v2297 = v2295 & int32(_a_F_ExecWindowAgg_21)
	*(*uint16)(unsafe.Add(mBase, uint32(v2277)+4)) = uint16(v2297)
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v2277)+12))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2299)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2277)+6)) = uint16(v2300)
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v211)+224))
	if v2302 != int32(1) {
		goto L58
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v2277
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v211)+312))
	if v2306 == int32(0) {
		goto L59
	} else {
		goto L429
	}
L429:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2310
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2306)+20))
	v2315 = m.T0[v2314].(func(*base.Module, int32, int32, int32) int32)(m, v2306, v264, v215+int32(12))
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L4
	} else {
		goto L430
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2283
	if v2315 != 0 {
		goto L59
	} else {
		goto L431
	}
L431:
	;
	v2319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+310)))
	if v2319 != int32(1) {
		v2503 = v215
		goto L61
	} else {
		goto L432
	}
L432:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v211)+120))
	if v2322 <= int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+311)))
	if v2494 != int32(1) {
		goto L60
	} else {
		goto L445
	}
L434:
	;
	v2326 = v2322 & int32(3)
	v2327 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2322) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2335 = v2327
	v2344 = int32(0)
	goto L438
L436:
	;
	v2411 = v2327
	goto L437
L437:
	;
	v2434 = v2411
	v2438 = v2327
	goto L442
L438:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v2358 = int32(2)
	v2361 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2357+v2335<<(uint(v2358)%32)))) = v2361
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	v2365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2363+v2335))) = uint8(v2365)
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v2369 = v2335 | v2365
	*(*int32)(unsafe.Add(mBase, uint32(v2367+v2369<<(uint(v2358)%32)))) = v2361
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2375+v2369))) = uint8(v2365)
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v2381 = v2335 | v2358
	*(*int32)(unsafe.Add(mBase, uint32(v2379+v2381<<(uint(v2358)%32)))) = v2361
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2387+v2381))) = uint8(v2365)
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v2393 = v2335 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v2391+v2393<<(uint(v2358)%32)))) = v2361
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2399+v2393))) = uint8(v2365)
	v2403 = int32(4)
	v2404 = v2335 + v2403
	v2406 = v2344 + v2403
	if v2406 != v2322&int32(2147483644) {
		v2335 = v2404
		v2344 = v2406
		goto L438
	} else {
		goto L440
	}
L439:
	;
	if v2326 == int32(0) {
		goto L433
	} else {
		goto L441
	}
L440:
	;
	goto L439
L441:
	;
	v2411 = v2404
	goto L437
L442:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2456+v2434<<(uint(int32(2))%32)))) = int32(0)
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	v2464 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2462+v2434))) = uint8(v2464)
	v2469 = v2438 + v2464
	if v2469 != v2326 {
		v2434 = v2434 + v2464
		v2438 = v2469
		goto L442
	} else {
		goto L444
	}
L443:
	;
	goto L433
L444:
	;
	goto L443
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+224)) = int32(3)
	goto L50
L446:
	;
	v2553 = int32(_a_F_ExecWindowAgg_2)
	v2554 = *(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1]))
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2556
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2550)+20))
	v2561 = m.T0[v2560].(func(*base.Module, int32, int32, int32) int32)(m, v2550, v264, v215+int32(12))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L4
	} else {
		goto L447
	}
L447:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecWindowAgg[1])) = v2554
	if v2561 != 0 {
		v2576 = v2277
		v2577 = v215
		goto L6
	} else {
		goto L448
	}
L448:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	if v2565 == int32(0) {
		goto L50
	} else {
		goto L449
	}
L449:
	;
	v2568 = *(*float64)(unsafe.Add(mBase, uint32(v2565)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v2565)+240)) = base.F64_add(v2568, float64(1))
	goto L50
L450:
	;
	goto L51
}
func F_window_lag_with_offset_and_default(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v8 + int32(15)
	v14 = F_WinGetFuncArgCurrent(m, v10, int32(1), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v18 == int32(0) {
			v21 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v23 == v21 {
				v71 = v21
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
				if v29 == int32(0) {
					v71 = v21
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v34 = v32 - int32(11)
					if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v34))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v34)%32))&int32(1) == int32(0))|int32(0) != 0 {
						v71 = v21
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_window_lag_with_offset_and_default[0])))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v29+v49)))
						if v51 == int32(0) {
							v71 = v21
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							if v54 <= int32(1) {
								v71 = v21
							} else {
								v56 = int32(1)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(4))))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								switch v62 - int32(7) {
								case 0:
									v71 = v56
								case 1:
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
									if v65 == int32(0) {
										v71 = v56
									} else {
										v71 = int32(0)
									}
								default:
									v71 = int32(0)
								}
							}
						}
					}
				}
			}
			v74 = F_WinGetFuncArgInPartition(m, v10, v21-v14, v71, v13, v8+int32(14))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
				if v76 == int32(1) {
					v80 = F_WinGetFuncArgCurrent(m, v10, int32(2), v13)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = v80
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v83 != int32(1) {
							v90 = v82
						} else {
							v87 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
							v90 = int32(0)
						}
						m.G0 = v8 + int32(16)
						return v90
					}
				} else {
					v82 = v74
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v83 != int32(1) {
						v90 = v82
					} else {
						v87 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
						v90 = int32(0)
					}
					m.G0 = v8 + int32(16)
					return v90
				}
			}
		} else {
			v87 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
			v90 = int32(0)
			m.G0 = v8 + int32(16)
			return v90
		}
	}
}
func F_window_lead_with_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v8 + int32(15)
	v14 = F_WinGetFuncArgCurrent(m, v10, int32(1), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = int32(0)
			if v21 == v23 {
				v69 = v23
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
				if v27 == int32(0) {
					v69 = v23
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v32 = v30 - int32(11)
					if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v32))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v32)%32))&int32(1) == int32(0))|int32(0) != 0 {
						v69 = v23
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_window_lead_with_offset[0])))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v27+v47)))
						if v49 == int32(0) {
							v69 = v23
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							if v52 <= int32(1) {
								v69 = v23
							} else {
								v54 = int32(1)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(4))))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								switch v60 - int32(7) {
								case 0:
									v69 = v54
								case 1:
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
									if v63 == int32(0) {
										v69 = v54
									} else {
										v69 = int32(0)
									}
								default:
									v69 = int32(0)
								}
							}
						}
					}
				}
			}
			v72 = F_WinGetFuncArgInPartition(m, v10, v14, v69, v13, v8+int32(14))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v74 != int32(1) {
					v81 = v72
				} else {
					v78 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
					v81 = int32(0)
				}
				m.G0 = v8 + int32(16)
				return v81
			}
		} else {
			v78 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v78)
			v81 = int32(0)
			m.G0 = v8 + int32(16)
			return v81
		}
	}
}
func F_window_lead_with_offset_and_default(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v8 + int32(15)
	v14 = F_WinGetFuncArgCurrent(m, v10, int32(1), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = int32(0)
			if v21 == v23 {
				v69 = v23
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
				if v27 == int32(0) {
					v69 = v23
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v32 = v30 - int32(11)
					if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v32))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v32)%32))&int32(1) == int32(0))|int32(0) != 0 {
						v69 = v23
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_window_lead_with_offset_and_default[0])))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v27+v47)))
						if v49 == int32(0) {
							v69 = v23
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							if v52 <= int32(1) {
								v69 = v23
							} else {
								v54 = int32(1)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(4))))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								switch v60 - int32(7) {
								case 0:
									v69 = v54
								case 1:
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
									if v63 == int32(0) {
										v69 = v54
									} else {
										v69 = int32(0)
									}
								default:
									v69 = int32(0)
								}
							}
						}
					}
				}
			}
			v72 = F_WinGetFuncArgInPartition(m, v10, v14, v69, v13, v8+int32(14))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
				if v74 == int32(1) {
					v78 = F_WinGetFuncArgCurrent(m, v10, int32(2), v13)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v80 = v78
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v81 != int32(1) {
							v88 = v80
						} else {
							v85 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
							v88 = int32(0)
						}
						m.G0 = v8 + int32(16)
						return v88
					}
				} else {
					v80 = v72
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v81 != int32(1) {
						v88 = v80
					} else {
						v85 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
						v88 = int32(0)
					}
					m.G0 = v8 + int32(16)
					return v88
				}
			}
		} else {
			v85 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
			v88 = int32(0)
			m.G0 = v8 + int32(16)
			return v88
		}
	}
}
